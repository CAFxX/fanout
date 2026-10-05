#!/usr/bin/env python3
"""Shared driver library for the GHA full-battery offload.

Candidate bundle layout (battery/candidates/<name>/):
    impl.c   - uint64_t mix_hash(uint64_t val, uint64_t key)
    model.py - mix(v,k) -> int, mix_np(V,K) -> np array
    meta.json- {rtl_top, screen_keys, practrand_keys, goldens[{val,key,out}]}
    rtl.v    - synthesizable RTL, top = meta.rtl_top, ports (val,key,out)

Env contract (set by the driver before importing vendored modules):
    BATTERY_HOME  battery/ root (vendored FB replacement)
    VENDOR_DIR    battery/vendor/ (vendored src/ replacement)
    FLOW_DIR      flow/ (mini.genlib, analyze.py, analyze2.py, mul64.v)
    MIX_EXPLORATION = FLOW_DIR (vendored analyze2.py lookup)
    YOSYS_BIN, IVERILOG_BIN, VVP_BIN, CC (default gcc)
    ABC_NODRETIME_SCRIPT, CELLS_SIM_V
    PRACTRAND_BIN (RNG_test), TESTU01_HOME (include/lib)
"""
import json
import os
import subprocess
import sys
import time

BATTERY_HOME = os.environ.get(
    "BATTERY_HOME",
    os.path.dirname(os.path.abspath(__file__)))
VENDOR = os.path.join(BATTERY_HOME, "vendor")
CANDIDATES = os.path.join(BATTERY_HOME, "candidates")


def setup_env(flow_dir=None):
    """Point every vendored path override at the battery/flow layout."""
    os.environ["BATTERY_HOME"] = BATTERY_HOME
    os.environ["VENDOR_DIR"] = VENDOR
    if flow_dir:
        os.environ["FLOW_DIR"] = flow_dir
        os.environ["MIX_EXPLORATION"] = flow_dir
    os.environ.setdefault("ABC_NODRETIME_SCRIPT",
                          os.path.join(VENDOR, "abc_nodretime.script"))
    os.environ.setdefault("CELLS_SIM_V",
                          os.path.join(VENDOR, "cells_sim.v"))
    if "YOSYS_BIN" not in os.environ:
        os.environ["YOSYS_BIN"] = \
            "/home/hatch/tools/micromamba-root/envs/sky130/bin/yosys"
    os.environ.setdefault(
        "IVERILOG_BIN", "/home/hatch/tools/micromamba-root/bin/iverilog")
    os.environ.setdefault(
        "VVP_BIN", "/home/hatch/tools/micromamba-root/bin/vvp")
    sys.path.insert(0, VENDOR)


def load_bundle(candidates_dir, name, need_rtl=True):
    d = os.path.join(candidates_dir, name)
    for f in ("impl.c", "model.py", "meta.json"):
        if not os.path.exists(os.path.join(d, f)):
            raise FileNotFoundError(f"bundle {name}: missing {f}")
    rtl_v = os.path.join(d, "rtl.v")
    if need_rtl and not os.path.exists(rtl_v):
        raise FileNotFoundError(f"bundle {name}: missing rtl.v")
    meta = json.load(open(os.path.join(d, "meta.json")))
    return {"name": name, "dir": d,
            "impl_c": os.path.join(d, "impl.c"),
            "model_py": os.path.join(d, "model.py"),
            "meta": meta,
            "rtl_v": rtl_v if os.path.exists(rtl_v) else None}


def compile_c(sources, out, extra=()):
    """Compile C sources with $CC -O2. Returns binary path (raises on fail)."""
    cc = os.environ.get("CC", "gcc")
    cmd = [cc, "-O2", "-o", out] + list(sources) + list(extra)
    r = subprocess.run(cmd, capture_output=True, text=True, timeout=600)
    if r.returncode != 0:
        raise RuntimeError(f"compile failed:\n{r.stderr[-3000:]}")
    return out


def xcheck_with_bundle_dir(bundle, binary, nvecs, workdir):
    """xcheck with BATTERY_HOME shadowed so model.py resolves from bundle."""
    import shutil
    shadow = os.path.join(workdir, "_shadow")
    fb = os.path.join(shadow, "candidates", bundle["name"])
    os.makedirs(fb, exist_ok=True)
    for f in ("model.py", "meta.json"):
        shutil.copy(os.path.join(bundle["dir"], f), os.path.join(fb, f))
    vecfile = os.path.join(workdir, "_xcheck_vecs.txt")
    env = dict(os.environ, BATTERY_HOME=shadow)
    r = subprocess.run(
        [sys.executable, os.path.join(VENDOR, "xcheck.py"),
         bundle["name"], binary, str(nvecs), vecfile],
        capture_output=True, text=True, timeout=600, env=env)
    ok = r.returncode == 0
    return ok, (r.stdout + r.stderr)[-1500:]


def apply_only(names, only):
    """Filter a candidate name list to a comma-separated allow-list."""
    if not only:
        return names
    want = {w.strip() for w in only.split(",") if w.strip()}
    return [n for n in names if n in want]


def list_candidates(candidates_dir, only=""):
    names = sorted(n for n in os.listdir(candidates_dir)
                   if os.path.isdir(os.path.join(candidates_dir, n)))
    return apply_only(names, only)


def shard_slice(items, idx, count):
    items = sorted(items)
    return [x for i, x in enumerate(items) if i % count == idx]


def write_result(out_dir, candidate, stage, verdict, detail):
    """Per-candidate result JSON: {stage, verdict, exact failing signal}."""
    os.makedirs(out_dir, exist_ok=True)
    rec = {"candidate": candidate, "stage": stage, "verdict": verdict,
           "detail": detail,
           "ts": time.strftime("%Y-%m-%dT%H:%M:%SZ", time.gmtime())}
    p = os.path.join(out_dir, f"{candidate}_{stage}.json")
    json.dump(rec, open(p, "w"), indent=1, default=str)
    return rec


class CanaryFail(Exception):
    pass


def run_canaries(stage, driver_fn, canaries_root, out_dir, expect,
                 need_rtl=False, kill_verdicts=("KILL", "FAIL")):
    """Golden canaries: expect = {"pass": "pass/<bundle>", "kill": "kill_s0/<bundle>"}.

    Paths are relative to canaries_root. Refuses the stage on mismatch.
    Pass verdicts: PASS (S0/S2/S3/S4) or CAL_PASS (S1 calibrated screen).
    Kill verdicts: KILL/FAIL by default; S1 passes kill_verdicts=("FLAG",)
    since its kill signal is FLAG.
    need_rtl: True only for S0 (synthesis needs the Verilog)."""
    for role, rel in expect.items():
        want = ("PASS", "CAL_PASS") if role == "pass" else kill_verdicts
        b = load_bundle(os.path.join(canaries_root, os.path.dirname(rel)),
                        os.path.basename(rel), need_rtl=need_rtl)
        rec = driver_fn(b, os.path.join(out_dir, "_canary", role))
        got = rec["verdict"]
        ok = (got in ("PASS", "CAL_PASS")) if role == "pass" else (got in kill_verdicts)
        print(f"[canary {stage}/{role}={rel}] want={want} got={got} "
              f"-> {'OK' if ok else 'MISMATCH'}", flush=True)
        if not ok:
            raise CanaryFail(
                f"canary {stage}/{role} ({rel}): want {want}, "
                f"got {got}; detail={rec['detail']}; STAGE QUARANTINED")
