#!/usr/bin/env python3
"""Per-shard S-box screen driver.

Modes:
  --box {midori_sb0,ulbc_s1,thf_blink_s0} --shard-idx I --num-shards N
      Enumerate variant slice [lo,hi) of vi EXACTLY per variant_lib
      (same scheme as the on-VM tick.py sweep).
  --pool-file PATH --shard-idx I --num-shards N
      Shard over pool entries instead of variants.

Per candidate: validate (dimension B); skip synthesis on gate failure and
record the reason; survivors are synthesized via the vendored pinned flow
(flow/run_synth.py) and max fanout is read from the mapped netlist.

CANARY (quarantine gate): before its slice, every shard synthesizes the
identity variant vi=0 of its box (== the base table) and requires
crit == GOLDEN_CRIT[box] EXACTLY (3.0u midori_sb0 / 3.4u ulbc_s1 /
3.4u thf_blink_s0). A mismatch means the flow diverged from the pinned
reference: the shard JSON records canary_ok=false and the driver exits
nonzero; the aggregate job then quarantines the whole box (no ranking).

Resumable/idempotent: --out shard JSON; variants already present in it are
skipped. State is flushed to <out>.tmp + atomic rename every 25 variants.

Output schema (per variant record):
  {vi|name, ip, op, c, table (hex16), validate: {...}, status,
   crit, cells, area, fanout, err}
Shard file: {box, mode, shard_idx, num_shards, slice:[lo,hi),
             canary:{expected, measured, ok}, records:[...]}
"""
import argparse
import json
import os
import re
import shutil
import subprocess
import sys
import tempfile

DRIVER = os.path.dirname(os.path.abspath(__file__))
REPO = os.path.dirname(DRIVER)
FLOW = os.path.join(REPO, "flow")
sys.path.insert(0, DRIVER)
sys.path.insert(0, FLOW)

import variant_lib as vl
import validate_sbox as vs
import gen_sbox
import liberty_screen

RUN_SYNTH = os.path.join(FLOW, "run_synth.py")
SYNTH_TIMEOUT = 600


def find_liberty(explicit=None):
    """Locate the sky130 tt liberty file (never baked into the image)."""
    cands = []
    if explicit:
        cands.append(explicit)
    if os.environ.get("LIBERTY_LIB"):
        cands.append(os.environ["LIBERTY_LIB"])
    pdk = os.environ.get("PDK_ROOT", "./pdk")
    cands.append(os.path.join(pdk, "sky130A", "libs.ref",
                              "sky130_fd_sc_hd", "lib",
                              "sky130_fd_sc_hd__tt_100C_1v80.lib"))
    home_volare = os.path.expanduser("~/.volare")
    for c in cands:
        if os.path.exists(c):
            return c
    # last resort: search under PDK_ROOT for the exact file (case-insensitive)
    for root, _, files in os.walk(pdk):
        for fn in files:
            if fn.lower() == "sky130_fd_sc_hd__tt_100c_1v80.lib":
                return os.path.join(root, fn)
    return None


def parse_result(out):
    for line in out.split("\n"):
        if line.startswith("RESULT"):
            crit = cells = area = None
            for tok in line.split():
                if tok.startswith("crit="):
                    crit = float(tok.split("=", 1)[1])
                elif tok.startswith("cells="):
                    cells = int(tok.split("=", 1)[1])
                elif tok.startswith("area="):
                    area = float(tok.split("=", 1)[1])
            if crit is not None and cells is not None and area is not None:
                return crit, cells, area
    return None, None, None


def max_fanout(mapped_v):
    sinks = {}
    with open(mapped_v) as f:
        for line in f:
            m = re.match(r"\s*\.(\w+)\(([^)]+)\)", line)
            if m:
                port, net = m.group(1), m.group(2)
                if port != "Y":
                    sinks[net] = sinks.get(net, 0) + 1
    return max(sinks.values()) if sinks else 0


def synth_table(tab, name, workroot):
    """Synthesize one table via the pinned flow. Returns
    (crit, cells, area, fanout, err)."""
    vp = os.path.join(workroot, f"{name}.v")
    with open(vp, "w") as f:
        f.write(gen_sbox.write_sbox64(tab, name))
    wd = os.path.join(workroot, f"synth_{name}")
    os.makedirs(wd, exist_ok=True)
    crit = cells = area = fanout = None
    err = None
    try:
        r = subprocess.run(
            [sys.executable, RUN_SYNTH, vp, name, wd],
            capture_output=True, text=True, timeout=SYNTH_TIMEOUT)
        out = r.stdout + r.stderr
        crit, cells, area = parse_result(out)
        if crit is None:
            err = f"no RESULT in output: {out[-500:]}"
        else:
            mapped = os.path.join(wd, "_dut", f"mapped_{name}.v")
            if os.path.exists(mapped):
                try:
                    fanout = max_fanout(mapped)
                except Exception as e:
                    err = f"fanout parse failed: {e}"
            else:
                err = "mapped netlist missing"
    except Exception as e:
        err = f"{type(e).__name__}: {e}"
    finally:
        shutil.rmtree(wd, ignore_errors=True)
        try:
            os.remove(vp)
        except OSError:
            pass
    return crit, cells, area, fanout, err


def load_state(out):
    if os.path.exists(out):
        try:
            return json.load(open(out))
        except Exception:
            pass
    return None


# ---------------- P&R stage (PLACEHOLDER — backend direction ON HOLD) -----
# Contract (stable regardless of which of the three P&R paths is chosen):
#   * eligibility: hard gates pass AND unit-delay crit < --pr-threshold
#   * eligible records will carry pr = {postroute_crit, routed_area_um2,
#     power_w, max_input_pin_cap_ff, corner, backend, err}
# Until the backend lands, eligible records get status "backend_tbd" and NO
# P&R tooling is invoked. Candidates failing earlier hurdles never pay P&R.
PR_THRESHOLD_DEFAULT = 3.4
LIBERTY_THRESHOLD_DEFAULT_PS = 333.0  # measured: thf_blink_s0 = 333.2 ps
# (3.4u tier). Strictly-better-than-tier, like the unit < 3.4u rule.


def pr_stage_placeholder(rec, pr_threshold=PR_THRESHOLD_DEFAULT,
                         liberty_threshold=LIBERTY_THRESHOLD_DEFAULT_PS):
    v = rec.get("validate") or {}
    g = v.get("gates") or {}
    gates_ok = all(g.get(k) for k in vs.HARD_GATES)
    unit_ok = (rec.get("status") == "measured"
               and rec.get("crit") is not None
               and rec["crit"] < pr_threshold)
    lib_ps = rec.get("liberty_crit_ps")
    lib_ok = (lib_ps is None) or (lib_ps < liberty_threshold)
    eligible = gates_ok and unit_ok and lib_ok
    if not eligible:
        rec["pr"] = {"eligible": False,
                     "reason": "hurdles not cleared "
                               f"(gates={gates_ok} unit<{pr_threshold}={unit_ok} "
                               f"liberty<{liberty_threshold}ps={lib_ok})"}
    else:
        rec["pr"] = {"eligible": True, "status": "backend_tbd",
                     "reason": "P&R runs on promote.py picks via the public "
                               "OpenLane image (see sbox-pr dispatch)"}


def save_state(out, state):
    tmp = out + ".tmp"
    with open(tmp, "w") as f:
        json.dump(state, f)
    os.replace(tmp, out)


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--box", choices=vl.ORDER, default=os.environ.get("BOX"))
    ap.add_argument("--pool-file")
    ap.add_argument("--shard-idx", type=int,
                    default=int(os.environ.get("SHARD_IDX", 0)))
    ap.add_argument("--num-shards", type=int,
                    default=int(os.environ.get("SHARD_COUNT", 1)))
    ap.add_argument("--out", required=True)
    ap.add_argument("--workroot", default=None)
    ap.add_argument("--pr-threshold", type=float, default=PR_THRESHOLD_DEFAULT,
                    help="P&R eligibility: unit-delay crit below this (u)")
    ap.add_argument("--liberty", default=None,
                    help="path to sky130 tt liberty (default: $LIBERTY_LIB / "
                         "$PDK_ROOT search; missing -> liberty_err recorded)")
    ap.add_argument("--no-liberty", action="store_true",
                    help="skip the liberty-delay screen entirely")
    ap.add_argument("--liberty-threshold", type=float,
                    default=LIBERTY_THRESHOLD_DEFAULT_PS,
                    help="P&R eligibility: liberty crit below this (ps)")
    a = ap.parse_args()
    assert (a.box is None) != (a.pool_file is None), "one of --box/--pool-file"

    workroot = a.workroot or tempfile.mkdtemp(prefix="sbox_screen_")
    os.makedirs(workroot, exist_ok=True)

    if a.box:
        mode = f"box:{a.box}"
        lo, hi = vl.shard_range(a.shard_idx, a.num_shards)
        items = [(vi, vl.variant_table(a.box, vi)) for vi in range(lo, hi)]
        canary_tab = vl.CANDS[a.box]
        canary_expect = vl.GOLDEN_CRIT[a.box]
        canary_name = "canary"
    else:
        mode = f"pool:{os.path.basename(a.pool_file)}"
        entries = vl.load_pool(a.pool_file)
        lo = (a.shard_idx * len(entries)) // a.num_shards
        hi = ((a.shard_idx + 1) * len(entries)) // a.num_shards
        items = [(name, tab) for name, tab in entries[lo:hi]]
        canary_tab = vl.CANDS["midori_sb0"]
        canary_expect = vl.GOLDEN_CRIT["midori_sb0"]
        canary_name = "canary_pool"

    state = load_state(a.out)
    done_keys = set()
    if state and state.get("slice") == [lo, hi] and state.get("mode") == mode:
        for r in state["records"]:
            done_keys.add(r["key"])
    else:
        state = {"box": a.box, "mode": mode, "shard_idx": a.shard_idx,
                 "num_shards": a.num_shards, "slice": [lo, hi],
                 "canary": None, "records": []}

    # ---- canary (quarantine gate) ----
    if state["canary"] is None:
        crit, cells, area, fanout, err = synth_table(
            canary_tab, f"{canary_name}_s{a.shard_idx}", workroot)
        ok = (err is None and crit == canary_expect)
        state["canary"] = {"expected": canary_expect, "measured": crit,
                           "cells": cells, "ok": ok, "err": err}
        save_state(a.out, state)
        print(f"CANARY expected={canary_expect} measured={crit} ok={ok} err={err}",
              flush=True)
        if not ok:
            sys.exit(f"QUARANTINE: canary mismatch for {mode} "
                     f"(expected {canary_expect}, got {crit}, err={err})")

    # ---- slice ----
    n = 0
    for key, tab in items:
        rkey = str(key)
        if rkey in done_keys:
            continue
        v = vs.validate(tab)
        rec = {"key": rkey, "table": vl.table_hex(tab), "validate": v,
               "status": None, "crit": None, "cells": None, "area": None,
               "fanout": None, "err": None}
        if isinstance(key, int):
            ip, op, c = vl.variant_spec(key)
            rec.update({"vi": key, "ip": ip, "op": op, "c": c})
        else:
            rec["name"] = key
        if not vs.hard_gates_pass(v):
            failed = [k for k in vs.HARD_GATES if not v["gates"][k]]
            rec["status"] = "gate_fail"
            rec["err"] = f"hard gates failed: {failed}"
        else:
            name = f"sbox_{a.shard_idx}_{rkey}".replace(":", "_")
            crit, cells, area, fanout, err = synth_table(tab, name, workroot)
            rec.update({"crit": crit, "cells": cells, "area": area,
                        "fanout": fanout, "err": err,
                        "status": "measured" if err is None else "synth_fail"})
            # stage (c): liberty-delay screen (real cell delay, ps)
            if rec["status"] == "measured" and not a.no_liberty:
                lib = find_liberty(a.liberty)
                if lib is None:
                    rec["liberty_err"] = ("sky130 tt liberty not found "
                                          "(see driver/ensure_pdk.py)")
                else:
                    try:
                        lr = liberty_screen.liberty_screen(
                            tab, name + "_lib", workroot, lib)
                        rec.update({"liberty_crit_ps": lr["crit_ps"],
                                    "liberty_cells": lr["n_cells"],
                                    "liberty_area_um2": lr["area_um2"],
                                    "liberty_err": None})
                    except Exception as e:
                        rec["liberty_err"] = f"{type(e).__name__}: {e}"
        pr_stage_placeholder(rec, a.pr_threshold, a.liberty_threshold)
        state["records"].append(rec)
        n += 1
        if n % 25 == 0:
            save_state(a.out, state)
            print(f"  ... {n} variants done in shard {a.shard_idx}", flush=True)
    save_state(a.out, state)
    ok_n = sum(1 for r in state["records"] if r["status"] == "measured")
    print(f"SHARD DONE {mode} [{lo},{hi}): {ok_n}/{len(state['records'])} measured",
          flush=True)


if __name__ == "__main__":
    main()
