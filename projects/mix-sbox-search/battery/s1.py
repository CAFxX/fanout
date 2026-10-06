#!/usr/bin/env python3
"""S1 driver: amended 4-family statistical screen (GHA battery).

Bit-exact port of fast_battery/src/s1_screen.py (vendored in ../vendor/):
LIN+DIFF on SEQ outputs first (10 keys, N=2^18), then the marginal PB+RD
468-cell calibrated screen (skipped on LIN/DIFF kill). Verdict CAL_PASS iff
zero Holm-Bonferroni rejections across all four families.

Golden canaries (run first; stage refused on mismatch):
  pass: r23_spn_mix4r_pba19b01_nw -> CAL_PASS
  kill: r23_1r (1-round)            -> FLAG (LIN+DIFF rejected)

Usage:
  s1.py --candidates-dir DIR --out OUT_DIR [--shard-idx I --shard-count N]
        [--flow-dir DIR] [--candidate NAME] [--skip-canary]
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
    flow_dir = a.flow_dir or os.path.join(
        os.path.dirname(common.BATTERY_HOME), "flow")
    common.setup_env(flow_dir)
    out_dir = os.environ.get("OUT_DIR", a.out)

    spec = importlib.util.spec_from_file_location(
        "vendored_s1_screen",
        os.path.join(common.VENDOR, "s1_screen.py"))
    s1 = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(s1)

    def stage_in(bundle):
        d = os.path.join(common.BATTERY_HOME, "candidates", bundle["name"])
        if os.path.abspath(d) != os.path.abspath(bundle["dir"]):
            os.makedirs(d, exist_ok=True)
            for f in ("model.py", "meta.json"):
                s = os.path.join(bundle["dir"], f)
                if os.path.exists(s):
                    shutil.copy(s, os.path.join(d, f))
        return bundle["name"]

    def drive_one(bundle, c_out):
        name = stage_in(bundle)
        os.makedirs(c_out, exist_ok=True)
        sys.argv = ["s1_screen", name, c_out]
        rc = s1.main()
        res = json.load(open(os.path.join(c_out, "s1_result.json")))
        verdict = res["verdict"]  # CAL_PASS | FLAG
        return common.write_result(
            out_dir, bundle["name"], "s1", verdict,
            {"pb_rejected": res["pb_rejected"],
             "red_rejected": res["red_rejected"],
             "lin_rejected": res["lin_rejected"],
             "diff_rejected": res["diff_rejected"],
             "marginal_skipped": res["marginal_skipped"],
             "time_marginal_s": res["time_marginal_s"],
             "time_lin_diff_s": res["time_lin_diff_s"],
             "native_result": "s1_result.json"})

    if not a.skip_canary:
        can_root = os.path.join(common.BATTERY_HOME, "canaries")
        common.run_canaries("s1", drive_one, can_root,
                            os.path.join(out_dir, "_canary"),
                            {"pass": "pass/r23_spn_mix4r_pba19b01_nw",
                             "kill": "kill_s1/r23_1r"},
                            kill_verdicts=("FLAG",))
        print("[s1] canaries OK", flush=True)

    if a.candidate:
        names = [a.candidate]
    else:
        names = common.list_candidates(a.candidates_dir, a.only)
    for n in common.shard_slice(names, a.shard_idx, a.shard_count):
        b = common.load_bundle(a.candidates_dir, n, need_rtl=False)
        rec = drive_one(b, os.path.join(out_dir, n))
        print(f"[s1 {n}] {rec['verdict']}", flush=True)
    print("[s1] shard done", flush=True)


if __name__ == "__main__":
    main()
