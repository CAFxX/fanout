#!/usr/bin/env python3
"""Promote S-box-level GH finds back toward the full design search.

Reads an aggregated results JSON (ranked_<box>_<date>.json from the screen
aggregate, or fuzz_finds.json from the fuzz aggregate), applies promotion
thresholds, and emits promote/picks_<YYYYMMDD>.json.

Default thresholds (all three stages must clear):
  - unconditional hard gates: bijective, DU=4, NL=4
    (degree-3-all and full-dependency are "where required" per the expert
    methodology — the 3.0u champion MIDORI_Sb0 is [3,2,3,3] without full
    dependency, so they are OPT-IN via --require-deg3/--require-full-dep,
    never silently exclusionary by default; always recorded per pick)
  - unit-delay crit < 3.4u  (strictly better than the ULBC_s1/THF_BLINK_S0
    3.4u tier; MIDORI_Sb0's 3.0u is the record to beat)
  - liberty-delay crit < 333.0 ps (typical corner, cell-delay only;
    strictly better than thf_blink_s0's measured 333.2 ps — the best of
    the 3.4u tier. Measured tier: midori_sb0 267.4, thf 333.2,
    ulbc_s1 366.3, s1 366.6 ps. Records without liberty data skip this
    gate with a warning, never silently.)
  - table NOT already present in pools/*.txt and NOT already recorded in
    results/promoted_registry.json  (dedup rule: never re-promote a table
    already in ROUND25_TABLE.md or the pools)

Each pick carries: the exact 16-entry table, ALL measured dimensions
(validate dict + crit/cells/area/fanout + pr record if present), provenance
(pool name / variant box+vi / fuzz seed+sample_idx), and the promotion reason.

Dedup registry: results/promoted_registry.json is a list of
{table_hex, promoted_date, status}. promote.py appends new picks with status
"promoted_pending_ingest". The parent (on the VM) flips entries to
"ingested_round25" when it merges them into ROUND25_TABLE.md, so a later
promote run never re-promotes the same table.

PIPELINE CONTRACT (standing): the GH side searches, validates and measures
at S-box level ONLY. A GH fast rank is NEVER a frontier claim. The parent
pulls promote/picks_*.json, plugs each picked S-box into the in-context
constructions (currently the 4-round pba19b01 SPN), and runs the full
S0->S1->S2->S3->S4 + sac_z5 + BIC battery on the VM — because
construction-specific kills (e.g. the MIDORI Sb0 zero-key S1 DIFF flag) can
only be proven in-context.

Usage: python3 promote.py --input results/ranked_midori_sb0_20261005.json
         [--crit-max 3.4] [--no-require-deg3] [--dedup results/promoted_registry.json]
         [--out-dir promote/]
"""
import argparse
import datetime
import json
import os
import sys

DRIVER = os.path.dirname(os.path.abspath(__file__))
REPO = os.path.dirname(DRIVER)
sys.path.insert(0, DRIVER)
import variant_lib as vl


def load_records(path):
    d = json.load(open(path))
    recs = d.get("records") or d.get("finds") or []
    return d, recs


def pool_tables_hex(pools_dir):
    s = set()
    for name, tab in vl.pool_tables(pools_dir):
        s.add(vl.table_hex(tab))
    return s


def registry_tables(reg_path):
    if not os.path.exists(reg_path):
        return {}
    try:
        return {e["table_hex"]: e for e in json.load(open(reg_path))}
    except Exception:
        return {}


def provenance_of(rec, src):
    if "vi" in rec:
        return {"source": "variant", "box": src,
                "vi": rec["vi"], "ip": rec.get("ip"),
                "op": rec.get("op"), "c": rec.get("c")}
    if rec.get("name"):
        return {"source": "pool", "name": rec["name"]}
    if "sample_idx" in rec:
        return {"source": "fuzz", "seed": rec.get("seed"),
                "sample_idx": rec["sample_idx"]}
    return {"source": src or "unknown"}


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--input", required=True)
    ap.add_argument("--out-dir", default=os.path.join(REPO, "promote"))
    ap.add_argument("--crit-max", type=float, default=3.4,
                    help="unit-delay promotion threshold (u)")
    ap.add_argument("--liberty-max", type=float, default=333.0,
                    help="liberty-delay promotion threshold (ps, typical "
                         "corner, cell-delay only). Default = just under "
                         "thf_blink_s0's measured 333.2 ps (strictly "
                         "better than the 3.4u tier, mirroring --crit-max). "
                         "Measured tier: midori_sb0 267.4, thf_blink_s0 "
                         "333.2, ulbc_s1 366.3, s1 366.6 ps.")
    ap.add_argument("--no-require-deg3", action="store_true",
                    help="(deprecated: deg3 is opt-in now; flag ignored)")
    ap.add_argument("--require-deg3", action="store_true",
                    help="opt-in: require all ANF degrees == 3")
    ap.add_argument("--require-full-dep", action="store_true",
                    help="opt-in: require full coordinate dependency")
    ap.add_argument("--dedup",
                    default=os.path.join(REPO, "results", "promoted_registry.json"))
    ap.add_argument("--pools-dir", default=os.path.join(REPO, "pools"))
    a = ap.parse_args()

    meta, recs = load_records(a.input)
    src = meta.get("box") or meta.get("source") or "ranked"
    seen = pool_tables_hex(a.pools_dir)
    reg = registry_tables(a.dedup)
    seen |= set(reg.keys())

    picks = []
    skipped = {"gate_fail": 0, "crit": 0, "dup": 0, "not_measured": 0}
    for r in recs:
        if r.get("status") != "measured" or r.get("crit") is None:
            skipped["not_measured"] += 1
            continue
        v = r.get("validate") or {}
        g = v.get("gates") or {}
        gates_ok = all(g.get(k) for k in ("bijective", "du4", "nl4"))
        if a.require_deg3:
            gates_ok = gates_ok and g.get("deg3_all")
        if a.require_full_dep:
            gates_ok = gates_ok and g.get("full_dependency")
        if not gates_ok:
            skipped["gate_fail"] += 1
            continue
        if not r["crit"] < a.crit_max:
            skipped["crit"] += 1
            continue
        lib_ps = r.get("liberty_crit_ps")
        if lib_ps is not None and not lib_ps < a.liberty_max:
            skipped["liberty"] = skipped.get("liberty", 0) + 1
            continue
        if r["table"] in seen:
            skipped["dup"] += 1
            continue
        seen.add(r["table"])
        prov = provenance_of(r, src)
        reason = (f"hard gates pass (DU={v['du']},NL={v['nl']}; "
                  f"deg={v['anf_degrees']},full-dep={v['full_dependency']}"
                  f"{' [deg3 required]' if a.require_deg3 else ''}"
                  f"{' [full-dep required]' if a.require_full_dep else ''}), "
                  f"unit {r['crit']:.1f}u < {a.crit_max:.1f}u, "
                  f"liberty {lib_ps:.1f}ps < {a.liberty_max:.1f}ps"
                  if lib_ps is not None else
                  f"hard gates pass (DU={v['du']},NL={v['nl']}), "
                  f"unit {r['crit']:.1f}u < {a.crit_max:.1f}u, "
                  f"liberty n/a, "
                  f"not in pools/registry")
        picks.append({
            "table": [int(c, 16) for c in r["table"]],
            "table_hex": r["table"],
            "validate": v,
            "unit_crit": r["crit"],
            "unit_cells": r.get("cells"),
            "unit_area": r.get("area"),
            "unit_fanout": r.get("fanout"),
            "liberty_crit_ps": r.get("liberty_crit_ps"),
            "liberty_cells": r.get("liberty_cells"),
            "liberty_area_um2": r.get("liberty_area_um2"),
            "pr": r.get("pr"),
            "provenance": prov,
            "reason": reason,
        })

    os.makedirs(a.out_dir, exist_ok=True)
    date = datetime.date.today().isoformat().replace("-", "")
    outp = os.path.join(a.out_dir, f"picks_{date}.json")
    with open(outp, "w") as f:
        json.dump({"date": date, "from": a.input,
                   "thresholds": {"crit_max_u": a.crit_max,
                                  "liberty_max_ps": a.liberty_max,
                                  "require_deg3": a.require_deg3,
                                  "require_full_dep": a.require_full_dep},
                   "picks": picks}, f, indent=1)

    # extend the dedup registry (pending-ingest status)
    reg_list = list(registry_tables(a.dedup).values())
    for p in picks:
        reg_list.append({"table_hex": p["table_hex"],
                         "promoted_date": date,
                         "status": "promoted_pending_ingest",
                         "from": a.input})
    os.makedirs(os.path.dirname(a.dedup) or ".", exist_ok=True)
    with open(a.dedup, "w") as f:
        json.dump(reg_list, f, indent=1)

    print(f"promote: {len(picks)} picks -> {outp}  "
          f"(skipped: {skipped})")


if __name__ == "__main__":
    main()
