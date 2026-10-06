#!/usr/bin/env python3
"""Campaign driver: full S0->S4 battery per candidate with early stop.

ONE job spec fans out to a matrix over candidates; EACH shard runs the
FULL staged pipeline (s0 -> s1 -> s2 -> s3 -> s4) for its candidate(s),
stopping each candidate at its first failing stage. Worst case ~3h per
candidate, under the 5h GHA cap.

Approval economy: a single push (one approval) covers dozens of
candidates end-to-end. Per-stage dispatches (one approval per stage)
are banned — chain the stages inside the job instead.

For each candidate the shard writes <name>_campaign.json:
  {candidate, stages: {s0: <record>, ...}, final_stage, final_verdict,
   survived}
where each stage record is the driver's <name>_<stage>.json
{stage, verdict, signal, artifacts}.

Golden canaries: the first candidate's run of each stage executes that
stage's canary (driver default); subsequent candidates pass
--skip-canary. A canary mismatch aborts the whole shard (red-green).

Usage:
  campaign.py --candidates-dir DIR --out OUT_DIR
      [--shard-idx I --shard-count N] [--only c1,c2] [--skip-canary]
      [--stages s0,s1,s2,s3,s4] [--max-bytes B]
Env: BATTERY_HOME (see common.py). Honors SHARD_IDX/SHARD_COUNT.
"""
import argparse
import json
import os
import subprocess
import sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import common

STAGES = ["s0", "s1", "s2", "s3", "s4"]
# stage -> pass verdicts (must match battery/aggregate.py)
PASS = {"s0": {"PASS"}, "s1": {"CAL_PASS"}, "s2": {"PASS"},
        "s3": {"PASS"}, "s4": {"PASS"}}


def parse_args():
    ap = argparse.ArgumentParser()
    ap.add_argument("--candidates-dir", required=True)
    ap.add_argument("--out", required=True)
    ap.add_argument("--shard-idx", type=int, default=int(
        os.environ.get("SHARD_IDX", 0)))
    ap.add_argument("--shard-count", type=int, default=int(
        os.environ.get("SHARD_COUNT", 1)))
    ap.add_argument("--only", default="")
    ap.add_argument("--skip-canary", action="store_true")
    ap.add_argument("--stages", default=",".join(STAGES))
    ap.add_argument("--max-bytes", type=int, default=1073741824,
                    help="S3 cap on GHA (breaking-point runs stay on VM)")
    return ap.parse_args()


def run_stage_driver(stage, candidates_dir, out_dir, candidate,
                     skip_canary, max_bytes):
    """Run one stage driver for one candidate. Returns the record dict."""
    here = os.path.dirname(os.path.abspath(__file__))
    cmd = [sys.executable, os.path.join(here, f"{stage}.py"),
           "--candidates-dir", candidates_dir,
           "--out", out_dir,
           "--only", candidate,
           # campaign.py already sharded; the driver must NOT re-shard
           # via SHARD_IDX/SHARD_COUNT env (would filter to empty).
           "--shard-idx", "0", "--shard-count", "1"]
    if skip_canary:
        cmd.append("--skip-canary")
    if stage == "s3":
        cmd += ["--max-bytes", str(max_bytes)]
    r = subprocess.run(cmd, capture_output=True, text=True, timeout=12000)
    sys.stdout.write(r.stdout)
    sys.stderr.write(r.stderr)
    if r.returncode != 0:
        # driver refused (canary mismatch) or crashed: stage verdict
        # is INFRA_FAIL; the campaign records it and stops the candidate.
        return {"candidate": candidate, "stage": stage,
                "verdict": "INFRA_FAIL",
                "detail": {"signal": f"driver rc={r.returncode}: "
                                     f"{r.stderr[-500:]}"}}
    rec_path = os.path.join(out_dir, f"{candidate}_{stage}.json")
    try:
        return json.load(open(rec_path))
    except Exception as e:
        return {"candidate": candidate, "stage": stage,
                "verdict": "INFRA_FAIL",
                "detail": {"signal": f"unreadable record: {e}"}}


def main():
    a = parse_args()
    # Absolute paths: vendored modules may os.chdir();
    # relative paths would silently break mid-run.
    a.candidates_dir = os.path.abspath(a.candidates_dir)
    a.out = os.path.abspath(a.out)
    common.setup_env()
    out_dir = os.environ.get("OUT_DIR", a.out)
    os.makedirs(out_dir, exist_ok=True)
    stages = [s.strip() for s in a.stages.split(",") if s.strip()]
    for s in stages:
        if s not in STAGES:
            sys.exit(f"unknown stage {s}")

    names = common.list_candidates(a.candidates_dir, a.only)
    mine = common.shard_slice(names, a.shard_idx, a.shard_count)
    print(f"[campaign] shard {a.shard_idx}/{a.shard_count}: "
          f"{len(mine)} candidates {mine}", flush=True)

    first_cand = not a.skip_canary
    for cand in mine:
        camp = {"candidate": cand, "stages": {}, "final_stage": None,
                "final_verdict": None, "survived": False}
        for stage in stages:
            # the first candidate's run of EACH stage executes that
            # stage's golden canary; later candidates skip it.
            rec = run_stage_driver(stage, a.candidates_dir, out_dir,
                                   cand, skip_canary=not first_cand,
                                   max_bytes=a.max_bytes)
            camp["stages"][stage] = {
                "verdict": rec["verdict"],
                "signal": rec["detail"].get("signal", ""),
                "artifacts": rec["detail"].get("native_result", ""),
            }
            camp["final_stage"] = stage
            camp["final_verdict"] = rec["verdict"]
            print(f"[campaign {cand}] {stage}: {rec['verdict']}",
                  flush=True)
            if rec["verdict"] not in PASS[stage]:
                break  # early stop for this candidate
        else:
            camp["survived"] = True
        json.dump(camp, open(os.path.join(
            out_dir, f"{cand}_campaign.json"), "w"), indent=1)
        print(f"[campaign {cand}] done: final={camp['final_stage']}/"
              f"{camp['final_verdict']} survived={camp['survived']}",
              flush=True)
        first_cand = False
    print("[campaign] shard done", flush=True)


if __name__ == "__main__":
    main()
