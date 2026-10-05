#!/usr/bin/env python3
"""Stochastic S-box search driver.

Samples random bijective 4-bit S-boxes (uniform random permutation of 16,
deterministic per --seed), validates (dimension B), synthesizes survivors
that pass the hard gates, and records everything.

Budget-bounded: exactly --n-samples candidates per seed. Deterministic:
same seed => same stream => same results (modulo synthesis, which is
deterministic given the pinned flow).

Output: {"seed":..., "n_samples":..., "canary": {...}, "finds": [...]}
where finds holds one record per gate-passing candidate:
  {sample_idx, table, validate:{...}, crit, cells, area, fanout, err}
The aggregate workflow merges finds with crit < 3.4u into results/fuzz_finds.json.
"""
import argparse
import json
import os
import random
import sys
import tempfile

DRIVER = os.path.dirname(os.path.abspath(__file__))
REPO = os.path.dirname(DRIVER)
FLOW = os.path.join(REPO, "flow")
sys.path.insert(0, DRIVER)
sys.path.insert(0, FLOW)

import variant_lib as vl
import validate_sbox as vs
from screen import synth_table, save_state, pr_stage_placeholder, \
    PR_THRESHOLD_DEFAULT, LIBERTY_THRESHOLD_DEFAULT_PS, find_liberty
import liberty_screen


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--seed", type=int,
                    default=int(os.environ.get("SHARD_IDX", 0)))
    ap.add_argument("--n-samples", type=int, required=True)
    ap.add_argument("--out", required=True)
    ap.add_argument("--workroot", default=None)
    ap.add_argument("--pr-threshold", type=float, default=PR_THRESHOLD_DEFAULT)
    ap.add_argument("--liberty-threshold", type=float,
                    default=LIBERTY_THRESHOLD_DEFAULT_PS)
    ap.add_argument("--liberty", default=None)
    ap.add_argument("--no-liberty", action="store_true")
    a = ap.parse_args()

    workroot = a.workroot or tempfile.mkdtemp(prefix="sbox_fuzz_")
    os.makedirs(workroot, exist_ok=True)
    rng = random.Random(a.seed)

    state = {"seed": a.seed, "n_samples": a.n_samples,
             "canary": None, "finds": [], "done": 0}
    if os.path.exists(a.out):
        try:
            prev = json.load(open(a.out))
            if prev.get("seed") == a.seed and prev.get("n_samples") == a.n_samples:
                state = prev
        except Exception:
            pass

    # canary: pinned flow must reproduce MIDORI_Sb0 = 3.0u
    if state["canary"] is None:
        crit, cells, area, fanout, err = synth_table(
            vl.CANDS["midori_sb0"], f"fuzz_canary_{a.seed}", workroot)
        ok = err is None and crit == vl.GOLDEN_CRIT["midori_sb0"]
        state["canary"] = {"expected": 3.0, "measured": crit, "ok": ok, "err": err}
        save_state(a.out, state)
        print(f"CANARY expected=3.0 measured={crit} ok={ok}", flush=True)
        if not ok:
            sys.exit(f"QUARANTINE: fuzz canary mismatch (got {crit}, err={err})")

    # resume: re-skip already-consumed RNG draws by replaying permutations
    # (cheap: permutation generation is microseconds)
    for i in range(state["done"]):
        rng.sample(range(16), 16)

    for i in range(state["done"], a.n_samples):
        tab = rng.sample(range(16), 16)
        v = vs.validate(tab)
        if vs.hard_gates_pass(v):
            crit, cells, area, fanout, err = synth_table(
                tab, f"fuzz_{a.seed}_{i}", workroot)
            rec = {
                "sample_idx": i, "table": vl.table_hex(tab), "validate": v,
                "crit": crit, "cells": cells, "area": area, "fanout": fanout,
                "err": err,
                "status": "measured" if err is None else "synth_fail"}
            if rec["status"] == "measured" and not a.no_liberty:
                lib = find_liberty(a.liberty)
                if lib is None:
                    rec["liberty_err"] = "sky130 tt liberty not found"
                else:
                    try:
                        lr = liberty_screen.liberty_screen(
                            tab, f"fuzz_{a.seed}_{i}_lib", workroot, lib)
                        rec.update({"liberty_crit_ps": lr["crit_ps"],
                                    "liberty_cells": lr["n_cells"],
                                    "liberty_area_um2": lr["area_um2"],
                                    "liberty_err": None})
                    except Exception as e:
                        rec["liberty_err"] = f"{type(e).__name__}: {e}"
            pr_stage_placeholder(rec, a.pr_threshold, a.liberty_threshold)
            state["finds"].append(rec)
        state["done"] = i + 1
        if (i + 1) % 100 == 0:
            save_state(a.out, state)
            print(f"  ... {i+1}/{a.n_samples} samples (seed {a.seed})", flush=True)
    save_state(a.out, state)
    print(f"FUZZ DONE seed={a.seed}: {len(state['finds'])} gate-passing finds "
          f"out of {a.n_samples} samples", flush=True)


if __name__ == "__main__":
    main()
