#!/usr/bin/env python3
"""Emit per-S-box P&R task specs from a ranked results JSON.

Usage: python3 pr_top.py --ranked results/ranked_<box>_<date>.json --n 32 \
           --out pr_tasks.json

Takes the top-N records by (crit, area), and emits a list of task specs the
P&R stage consumes (one per S-box):
  {name, table (16-list), table_hex, crit, cells, area, provenance}

NOTE (2026-10-05): the P&R backend direction is ON HOLD (three candidate
paths under review: single image + source-built OpenROAD, small screen image
+ public P&R image, or single small image + abc -liberty Sky130 mapping
instead of full P&R). This tool defines the stable task-spec CONTRACT the
future backend will consume; it invokes no P&R tooling itself.
"""
import argparse
import json
import os
import sys

DRIVER = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, DRIVER)
import variant_lib as vl


def load_records(ranked):
    d = json.load(open(ranked))
    recs = d.get("records") or d.get("finds") or []
    src = d.get("box") or d.get("source") or "unknown"
    return src, recs


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--ranked", required=True)
    ap.add_argument("--n", type=int, default=32)
    ap.add_argument("--out", required=True)
    a = ap.parse_args()

    src, recs = load_records(a.ranked)
    ok = [r for r in recs
          if r.get("status") == "measured" and r.get("crit") is not None]
    ok.sort(key=lambda r: (r["crit"], r.get("area") or 0))
    tasks = []
    for r in ok[:a.n]:
        tab = [int(c, 16) for c in r["table"]]
        prov = {"source": src}
        if "vi" in r:
            prov.update({"kind": "variant", "vi": r["vi"],
                         "ip": r.get("ip"), "op": r.get("op"), "c": r.get("c")})
        elif "name" in r:
            prov.update({"kind": "pool", "name": r["name"]})
        elif "sample_idx" in r:
            prov.update({"kind": "fuzz", "sample_idx": r["sample_idx"]})
        tasks.append({
            "name": f"pr_{r.get('key', r.get('vi', '?'))}",
            "table": tab,
            "table_hex": r["table"],
            "unit_crit": r["crit"],
            "unit_cells": r.get("cells"),
            "unit_area": r.get("area"),
            "provenance": prov,
        })
    with open(a.out, "w") as f:
        json.dump({"from": a.ranked, "n": a.n, "tasks": tasks}, f, indent=1)
    print(f"pr_top: {len(tasks)} P&R task specs -> {a.out}")


if __name__ == "__main__":
    main()
