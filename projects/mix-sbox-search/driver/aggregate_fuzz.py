#!/usr/bin/env python3
"""Aggregate per-seed fuzz JSONs into results/fuzz_finds.json.

QUARANTINE RULE: same as aggregate_screen — any seed whose canary
(MIDORI_Sb0 == 3.0u) failed aborts the run with exit 1 and no merge.

Merge policy: keep every gate-passing find with crit < --crit-max that is
not already present (dedup by table hex), preserving earlier entries.
Writes results/fuzz_finds.json = {"updated": <date>, "finds": [...]}.

Usage: python3 aggregate_fuzz.py --shards-dir fuzz_shards/ --out-dir results/
         [--crit-max 3.4]
"""
import argparse
import datetime
import glob
import json
import os
import sys


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--shards-dir", required=True)
    ap.add_argument("--out-dir", required=True)
    ap.add_argument("--crit-max", type=float, default=3.4)
    a = ap.parse_args()

    files = sorted(glob.glob(os.path.join(a.shards_dir, "*.json")))
    if not files:
        sys.exit("aggregate_fuzz: no seed JSONs found — failing")
    seeds = [json.load(open(f)) for f in files]
    bad = [s for s in seeds if not (s.get("canary") or {}).get("ok")]
    if bad:
        sys.exit(f"QUARANTINE: canary failed in seeds "
                 f"{[s.get('seed') for s in bad]} — no merge")

    outp = os.path.join(a.out_dir, "fuzz_finds.json")
    prev = []
    if os.path.exists(outp):
        try:
            prev = json.load(open(outp)).get("finds", [])
        except Exception:
            pass
    seen = {f["table"] for f in prev}
    new = 0
    for s in seeds:
        for fnd in s.get("finds", []):
            if (fnd.get("status") == "measured"
                    and fnd.get("crit") is not None
                    and fnd["crit"] < a.crit_max
                    and fnd["table"] not in seen):
                seen.add(fnd["table"])
                fnd["seed"] = s.get("seed")
                prev.append(fnd)
                new += 1
    prev.sort(key=lambda r: (r["crit"], r.get("area") or 0))
    os.makedirs(a.out_dir, exist_ok=True)
    with open(outp, "w") as f:
        json.dump({"updated": datetime.date.today().isoformat(),
                   "crit_max": a.crit_max,
                   "finds": prev}, f, indent=1)
    print(f"aggregate_fuzz: {new} new sub-{a.crit_max}u finds, "
          f"{len(prev)} total -> {outp}")


if __name__ == "__main__":
    main()
