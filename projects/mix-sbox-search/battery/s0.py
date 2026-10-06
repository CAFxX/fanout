#!/usr/bin/env python3
"""S0 driver: unit-delay synthesis probe (GHA battery).

Bit-exact port of fast_battery/src/s0_delay.py (vendored in ../vendor/):
Yosys 0.69, memory_map before techmap, abc -genlib mini.genlib with the
SAT-proven no-dretime script, dual-analyzer cross-check, 1000-vector
iverilog mapped-netlist equivalence check. Bar: delay_ratio_vs_mult <= 0.9.

Golden canaries (run first; stage refused on mismatch):
  pass: r23_spn_mix4r_pba19b01_nw -> PASS (~28.1u)
  kill: mul64                        -> FAIL (ratio 1.0 > 0.9)

Usage:
  s0.py --candidates-dir DIR --out OUT_DIR [--shard-idx I --shard-count N]
        [--flow-dir DIR] [--candidate NAME] [--skip-canary]
Env: BATTERY_HOME, FLOW_DIR, YOSYS_BIN, IVERILOG_BIN, VVP_BIN (see common.py).
"""
import argparse
import importlib.util
import json
import os
import shutil
import sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import common


def parse_args():
    ap = argparse.ArgumentParser()
    ap.add_argument("--candidates-dir", required=True)
    ap.add_argument("--out", required=True)
    ap.add_argument("--shard-idx", type=int, default=int(
        os.environ.get("SHARD_IDX", 0)))
    ap.add_argument("--shard-count", type=int, default=int(
        os.environ.get("SHARD_COUNT", 1)))
    ap.add_argument("--flow-dir", default=os.environ.get("FLOW_DIR", ""))
    ap.add_argument("--candidate", default="")
    ap.add_argument("--only", default="",
                        help="comma-separated candidate allow-list")
    ap.add_argument("--skip-canary", action="store_true")
    return ap.parse_args()


def main():
    a = parse_args()
    # Absolute paths: vendored modules may os.chdir();
    # relative paths would silently break mid-run.
    a.candidates_dir = os.path.abspath(a.candidates_dir)
    a.out = os.path.abspath(a.out)
    flow_dir = os.path.abspath(a.flow_dir or os.path.join(
        os.path.dirname(common.BATTERY_HOME), "flow"))
    common.setup_env(flow_dir)
    out_dir = os.environ.get("OUT_DIR", a.out)

    spec = importlib.util.spec_from_file_location(
        "vendored_s0_delay",
        os.path.join(common.VENDOR, "s0_delay.py"))
    s0 = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(s0)

    def stage_in(bundle):
        d = os.path.join(common.BATTERY_HOME, "candidates", bundle["name"])
        if os.path.abspath(d) != os.path.abspath(bundle["dir"]):
            os.makedirs(d, exist_ok=True)
            for f in ("impl.c", "model.py", "meta.json", "rtl.v"):
                s = os.path.join(bundle["dir"], f)
                if os.path.exists(s):
                    shutil.copy(s, os.path.join(d, f))
        return bundle["name"]

    def drive_one(bundle, c_out):
        name = stage_in(bundle)
        os.makedirs(c_out, exist_ok=True)
        sys.argv = ["s0_delay", name, c_out]
        rc = s0.main()
        res = json.load(open(os.path.join(c_out, "s0_result.json")))
        verdict = res["verdict"]  # PASS | FAIL
        return common.write_result(
            out_dir, bundle["name"], "s0", verdict,
            {"delay_u": res["delay_u"], "cells": res["cells"],
             "area": res["area"],
             "delay_ratio_vs_mult": res["delay_ratio_vs_mult"],
             "equiv": res["equiv_check"],
             "native_result": "s0_result.json"})

    if not a.skip_canary:
        can_root = os.path.join(common.BATTERY_HOME, "canaries")
        common.run_canaries("s0", drive_one, can_root,
                            os.path.join(out_dir, "_canary"),
                            {"pass": "pass/r23_spn_mix4r_pba19b01_nw",
                             "kill": "kill_s0/mul64"},
                            need_rtl=True)
        print("[s0] canaries OK", flush=True)

    if a.candidate:
        names = [a.candidate]
    else:
        names = common.list_candidates(a.candidates_dir, a.only)
    for n in common.shard_slice(names, a.shard_idx, a.shard_count):
        b = common.load_bundle(a.candidates_dir, n, need_rtl=True)
        rec = drive_one(b, os.path.join(out_dir, n))
        print(f"[s0 {n}] {rec['verdict']} "
              f"{rec['detail'].get('delay_u')}u", flush=True)
    print("[s0] shard done", flush=True)


if __name__ == "__main__":
    main()
