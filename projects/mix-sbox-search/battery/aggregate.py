#!/usr/bin/env python3
"""Aggregate campaign results (agent-side).

Scans <shards-dir>/shard-*/<candidate>_campaign.json (written by
battery/campaign.py) and produces:
  * stdout: one line per candidate: <candidate> <final_stage>
    <final_verdict> survived=<bool>
  * campaign_summary.json: per-candidate records + survivors list +
    per-stage kill counts

A candidate SURVIVES iff it passes all five stages (s0..s4).

Usage: aggregate.py --shards-dir DIR [--out DIR]
"""
import argparse
import glob
import json
import os
import sys

STAGES = ["s0", "s1", "s2", "s3", "s4"]


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--shards-dir", required=True)
    ap.add_argument("--out", default="")
    a = ap.parse_args()

    camps = {}
    pat = os.path.join(a.shards_dir, "shard-*", "*_campaign.json")
    for f in sorted(glob.glob(pat)):
        try:
            r = json.load(open(f))
        except Exception as e:
            print(f"WARNING: unreadable {f}: {e}", file=sys.stderr)
            continue
        camps[r["candidate"]] = r  # one record per candidate wins

    kill_at = {s: 0 for s in STAGES}
    kill_at["infra"] = 0
    survivors, killed = [], []
    for name in sorted(camps):
        c = camps[name]
        sig = c["stages"].get(c["final_stage"], {}).get("signal", "")
        if isinstance(sig, str):
            sig = sig.splitlines()[0] if sig else ""
        print(f"{name} {c['final_stage']} {c['final_verdict']} "
              f"survived={c['survived']} {sig}", flush=True)
        if c["survived"]:
            survivors.append(name)
        else:
            killed.append(name)
            if c["final_verdict"] == "INFRA_FAIL":
                kill_at["infra"] += 1
            elif c["final_stage"] in kill_at:
                kill_at[c["final_stage"]] += 1

    summary = {"n_total": len(camps), "n_survivors": len(survivors),
               "n_killed": len(killed), "kill_at_stage": kill_at,
               "survivors": survivors, "killed": killed,
               "campaigns": camps}
    out = a.out or a.shards_dir
    sp = os.path.join(out, "campaign_summary.json")
    json.dump(summary, open(sp, "w"), indent=1)
    print(f"survivors ({len(survivors)}/{len(camps)}): "
          f"{','.join(survivors)}", flush=True)
    print(f"kill histogram: {kill_at}", flush=True)
    print(f"wrote {sp}")


if __name__ == "__main__":
    main()
