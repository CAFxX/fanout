#!/usr/bin/env python3
"""S2 driver: differential early-kill profiler (GHA battery).

Bit-exact port of the fast_battery S2 stage: compiles
vendor/s2_diffprof.c + bundle impl.c, gates on the C-vs-Python xcheck
(10k vectors), then runs the differential sweep
(N=2^22, 115 deltas, kill_repeats from battery/ref/s2_kill_repeats.txt,
default 10). Verdict PASS iff no delta exceeds kill_repeats.

Golden canaries (run first; stage refused on mismatch):
  pass: r23_spn_mix4r_pba19b01_nw -> PASS
  kill: identity                     -> KILL (repeats)

Usage:
  s2.py --candidates-dir DIR --out OUT_DIR [--shard-idx I --shard-count N]
        [--candidate NAME] [--skip-canary] [--log2n 22] [--kill-repeats R]
"""
import argparse
import json
import os
import re
import subprocess
import sys
import time

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import common

XCHECK_VECS = 10000


def parse_args():
    ap = argparse.ArgumentParser()
    ap.add_argument("--candidates-dir", required=True)
    ap.add_argument("--out", required=True)
    ap.add_argument("--shard-idx", type=int, default=int(
        os.environ.get("SHARD_IDX", 0)))
    ap.add_argument("--shard-count", type=int, default=int(
        os.environ.get("SHARD_COUNT", 1)))
    ap.add_argument("--candidate", default="")
    ap.add_argument("--only", default="",
                        help="comma-separated candidate allow-list")
    ap.add_argument("--skip-canary", action="store_true")
    ap.add_argument("--log2n", type=int, default=22)
    ap.add_argument("--kill-repeats", type=int, default=0,
                    help="0 = read battery/ref/s2_kill_repeats.txt")
    return ap.parse_args()


def kill_repeats_default():
    p = os.path.join(common.BATTERY_HOME, "ref", "s2_kill_repeats.txt")
    try:
        return int(open(p).read().strip())
    except Exception:
        return 10


def drive_one(bundle, c_out, log2n, kill_repeats):
    os.makedirs(c_out, exist_ok=True)
    work = os.path.join(c_out, "_work")
    os.makedirs(work, exist_ok=True)
    t0 = time.time()
    binary = os.path.join(work, "diffprof")
    try:
        common.compile_c(
            [os.path.join(common.VENDOR, "s2_diffprof.c"), bundle["impl_c"]],
            binary, extra=["-lm"])
    except RuntimeError as e:
        return common.write_result(
            os.path.dirname(c_out), bundle["name"], "s2", "INFRA_FAIL",
            {"signal": f"compile failed: {e}"})
    ok, msg = common.xcheck_with_bundle_dir(
        bundle, binary, XCHECK_VECS, work)
    if not ok:
        return common.write_result(
            os.path.dirname(c_out), bundle["name"], "s2", "INFRA_FAIL",
            {"signal": f"xcheck C-vs-Python mismatch: {msg}"})
    key = bundle["meta"]["practrand_keys"][0]
    r = subprocess.run(
        [binary, key, str(log2n), str(kill_repeats)],
        capture_output=True, text=True, timeout=3600, cwd=work)
    out = r.stdout + r.stderr
    open(os.path.join(c_out, "s2.out"), "w").write(out)
    verdict, signal = "PASS", "no delta exceeded kill_repeats"
    m = re.search(r"KILL delta=(0x[0-9a-f]+) repeats=(\d+)", out)
    if r.returncode == 2 or m:
        verdict = "KILL"
        signal = (f"KILL delta={m.group(1)} repeats={m.group(2)}"
                  if m else f"rc=2: {out[-300:]}")
    elif r.returncode != 0:
        verdict = "INFRA_FAIL"
        signal = f"diffprof rc={r.returncode}: {out[-500:]}"
    else:
        pm = re.search(r"PASS max_repeats=(\d+) max\|z\|=([\d.]+)", out)
        if pm:
            signal = (f"PASS max_repeats={pm.group(1)} "
                      f"max|z|={pm.group(2)}")
    return common.write_result(
        os.path.dirname(c_out), bundle["name"], "s2", verdict,
        {"signal": signal, "log2n": log2n, "kill_repeats": kill_repeats,
         "key": key, "xcheck_vecs": XCHECK_VECS,
         "wall_s": round(time.time() - t0, 1),
         "native_result": "s2.out"})


def main():
    a = parse_args()
    # Absolute paths: vendored modules may os.chdir();
    # relative paths would silently break mid-run.
    a.candidates_dir = os.path.abspath(a.candidates_dir)
    a.out = os.path.abspath(a.out)
    common.setup_env()
    out_dir = os.environ.get("OUT_DIR", a.out)
    kr = a.kill_repeats or kill_repeats_default()

    if not a.skip_canary:
        can_root = os.path.join(common.BATTERY_HOME, "canaries")
        common.run_canaries(
            "s2",
            lambda b, o: drive_one(b, o, a.log2n, kr),
            can_root, os.path.join(out_dir, "_canary"),
            {"pass": "pass/r23_spn_mix4r_pba19b01_nw",
             "kill": "kill_s234/identity"})
        print("[s2] canaries OK", flush=True)

    if a.candidate:
        names = [a.candidate]
    else:
        names = common.list_candidates(a.candidates_dir, a.only)
    for n in common.shard_slice(names, a.shard_idx, a.shard_count):
        b = common.load_bundle(a.candidates_dir, n, need_rtl=False)
        rec = drive_one(b, os.path.join(out_dir, n), a.log2n, kr)
        print(f"[s2 {n}] {rec['verdict']}: "
              f"{rec['detail'].get('signal')}", flush=True)
    print("[s2] shard done", flush=True)


if __name__ == "__main__":
    main()
