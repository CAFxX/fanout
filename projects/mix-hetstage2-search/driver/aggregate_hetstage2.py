#!/usr/bin/env python3
"""Aggregate hetstage2 campaign shards.

Reads results/<project>/<job-id>/shard_*.json (or out/*.json), verifies
every shard's canary_ok, merges keeps, ranks by fisher z. Quarantines the
campaign (no ranking) if any shard's canary failed.

Usage: aggregate_hetstage2.py <results-dir> [--out keeps.jsonl]
"""
import glob
import json
import os
import sys


def main():
    d = sys.argv[1]
    out = sys.argv[2] if len(sys.argv) > 2 else None
    files = sorted(glob.glob(os.path.join(d, "*.json")))
    keeps = []
    nrec = 0
    bad = []
    for fp in files:
        if os.path.basename(fp).startswith("_"):
            continue
        with open(fp) as f:
            res = json.load(f)
        if not res.get("canary", {}).get("ok", False):
            bad.append(fp)
            continue
        for r in res.get("records", []):
            nrec += 1
            if r.get("keep"):
                keeps.append(r)
    if bad:
        print("QUARANTINE: canary failed on %d shard(s): %s — no ranking"
              % (len(bad), bad), flush=True)
        sys.exit(2)
    keeps.sort(key=lambda r: r["z"])
    print("aggregate: %d records, %d KEEP, %d shard files" %
          (nrec, len(keeps), len(files)), flush=True)
    for r in keeps[:20]:
        print("  keep z=%.2f kind=%s/%s inner=%s/%s nrej=%d" %
              (r["z"], r["kind"], r["template"], r["inner_lin"],
               r["inner_diff"], r["n_reject"]), flush=True)
    if out:
        with open(out, "w") as f:
            for r in keeps:
                f.write(json.dumps(r) + "\n")
        print("wrote %s" % out, flush=True)


if __name__ == "__main__":
    main()
