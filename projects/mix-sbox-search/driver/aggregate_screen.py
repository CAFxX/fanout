#!/usr/bin/env python3
"""Aggregate per-shard screen JSONs into a ranked results file.

QUARANTINE RULE: every shard carries a golden canary (identity variant vi=0
synthesized before its slice; must equal GOLDEN_CRIT[box] exactly). If ANY
shard reports canary.ok == false (or a missing canary), the aggregate FAILS
LOUDLY (exit 1) and writes NO ranking — the flow diverged from the pinned
reference and all numbers are untrusted.

On success writes:
  results/ranked_<box>_<YYYYMMDD>.json  (all records, ranked by crit, area)
  results/top32_<box>_<YYYYMMDD>.json   (top-32 measured records)

Usage: python3 aggregate_screen.py --box midori_sb0 --shards-dir shards/
         --out-dir results/
"""
import argparse
import datetime
import glob
import json
import os
import sys


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--box", required=True)
    ap.add_argument("--shards-dir", required=True)
    ap.add_argument("--out-dir", required=True)
    a = ap.parse_args()

    files = sorted(glob.glob(os.path.join(a.shards_dir, "*.json")))
    if not files:
        sys.exit("aggregate_screen: no shard JSONs found — failing")
    shards = [json.load(open(f)) for f in files]

    bad = [s for s in shards
           if not (s.get("canary") or {}).get("ok")]
    if bad:
        idxs = [s.get("shard_idx") for s in bad]
        sys.exit(f"QUARANTINE: canary failed in shards {idxs} for box "
                 f"{a.box} — flow diverged from pinned reference; "
                 f"no ranking written")

    recs = []
    for s in shards:
        recs.extend(s.get("records", []))
    measured = [r for r in recs if r.get("status") == "measured"]
    measured.sort(key=lambda r: (r["crit"], r.get("area") or 0))

    date = datetime.date.today().isoformat().replace("-", "")
    os.makedirs(a.out_dir, exist_ok=True)
    ranked_p = os.path.join(a.out_dir, f"ranked_{a.box}_{date}.json")
    top_p = os.path.join(a.out_dir, f"top32_{a.box}_{date}.json")
    with open(ranked_p, "w") as f:
        json.dump({"box": a.box, "date": date, "canary_ok": True,
                   "n_shards": len(shards), "n_records": len(recs),
                   "n_measured": len(measured),
                   "records": sorted(recs,
                                     key=lambda r: ((r.get("crit") if r.get("crit") is not None else 1e9),
                                                    r.get("area") or 0))}, f)
    with open(top_p, "w") as f:
        json.dump({"box": a.box, "date": date, "records": measured[:32]}, f,
                  indent=1)
    best = measured[0] if measured else None
    print(f"aggregate_screen: {len(measured)}/{len(recs)} measured, "
          f"canary OK on {len(shards)} shards")
    if best:
        print(f"  best: {best['table']} crit={best['crit']}u "
              f"cells={best['cells']} area={best['area']:.0f}")
    print(f"  wrote {ranked_p}\n  wrote {top_p}")


if __name__ == "__main__":
    main()
