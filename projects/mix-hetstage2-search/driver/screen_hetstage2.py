#!/usr/bin/env python3
"""Per-shard hetstage2 screen driver (GHA).

Evaluates a slice of pools/hetstage2_pool.jsonl with the verbatim eval
port (driver/hetstage2_eval.py): inner-DIFF-SEQ PRIMARY triage, fisher
triage second. KEEP iff inner lin==PASS and diff==PASS and
fisher n_reject==0.

CANARY (quarantine gate): before its slice, every shard evaluates
pools/canary.jsonl — the 9 pilot keeps must reproduce inner-PASS +
fisher-PASS with |z - expect_z| < 0.5, and the try53 replay must
inner-FLAG. On mismatch: canary_ok=false, driver exits nonzero; the
aggregate quarantines the campaign (no ranking).

Shard contract: $SHARD_IDX / $SHARD_COUNT (0-based, contiguous slice);
all outputs under $OUT_DIR (from env, never a literal).
Resumable/idempotent: records already present in the out file are
skipped; state flushed to <out>.tmp + atomic rename every 25 draws.

Output schema: {pool, job, shard_idx, num_shards, slice:[lo,hi),
  canary:{ok, detail}, records:[{pool_idx, kind, template, stages,
  inner_lin, inner_diff, n_reject, z, min_fisher, keep}]}
"""
import argparse
import json
import os
import sys

DRIVER = os.path.dirname(os.path.abspath(__file__))
TENANT = os.path.dirname(DRIVER)
sys.path.insert(0, DRIVER)
import hetstage2_eval as E

CANARY_Z_TOL = 0.5


def load_jsonl(path):
    with open(path) as f:
        return [json.loads(l) for l in f if l.strip()]


def run_canary():
    """Returns (ok, detail). Quarantine gate for the shard."""
    can = load_jsonl(os.path.join(TENANT, "pools", "canary.jsonl"))
    detail = []
    ok = True
    for rec in can:
        stages = [E.deser(s) for s in rec["stages"]]
        r = E.evaluate(stages)
        name = rec["name"]
        if rec["expect"] == "KEEP":
            good = (r["keep"] and
                    abs(r["z"] - rec["expect_z"]) < CANARY_Z_TOL)
            detail.append("%s: keep=%s z=%.2f expect_z=%.2f %s"
                          % (name, r["keep"], r["z"], rec["expect_z"],
                             "OK" if good else "MISMATCH"))
            ok = ok and good
        elif rec["expect"] == "INNER-FLAG":
            good = (r["inner_diff"] != "PASS")
            detail.append("%s: inner_diff=%s %s"
                          % (name, r["inner_diff"],
                             "OK" if good else "MISMATCH"))
            ok = ok and good
        else:
            raise AssertionError("canary: unknown expect %r" % rec["expect"])
    return ok, detail


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--pool-file", required=True)
    ap.add_argument("--shard-idx", type=int, required=False,
                    default=None)
    ap.add_argument("--num-shards", type=int, required=False,
                    default=None)
    ap.add_argument("--out", required=True)
    a = ap.parse_args()
    if a.shard_idx is None:
        a.shard_idx = int(os.environ["SHARD_IDX"])
    if a.num_shards is None:
        a.num_shards = int(os.environ["SHARD_COUNT"])
    out_dir = os.path.dirname(os.path.abspath(a.out))
    os.makedirs(out_dir, exist_ok=True)

    pool = load_jsonl(a.pool_file)
    n = len(pool)
    lo = (n * a.shard_idx) // a.num_shards
    hi = (n * (a.shard_idx + 1)) // a.num_shards

    done = set()
    records = []
    if os.path.exists(a.out):
        try:
            with open(a.out) as f:
                prev = json.load(f)
            records = prev.get("records", [])
            for r in records:
                done.add(r["pool_idx"])
        except Exception:
            records = []
            done = set()

    can_ok, can_detail = run_canary()
    for d in can_detail:
        print("CANARY %s" % d, flush=True)
    if not can_ok:
        print("CANARY FAIL — quarantining shard", flush=True)
        res = {"pool": a.pool_file, "shard_idx": a.shard_idx,
               "num_shards": a.num_shards, "slice": [lo, hi],
               "canary": {"ok": False, "detail": can_detail}, "records": []}
        tmp = a.out + ".tmp"
        with open(tmp, "w") as f:
            json.dump(res, f)
        os.replace(tmp, a.out)
        sys.exit(2)

    records = []
    done = set()
    if os.path.exists(a.out):
        try:
            with open(a.out) as f:
                prev = json.load(f)
            records = prev.get("records", [])
            for r in records:
                done.add(r["pool_idx"])
        except Exception:
            records = []
            done = set()
    for idx in range(lo, hi):
        if idx in done:
            continue
        rec = pool[idx]
        stages = [E.deser(s) for s in rec["stages"]]
        r = E.evaluate(stages)
        records.append({"pool_idx": idx, "kind": rec["kind"],
                        "template": rec["template"], "stages": rec["stages"],
                        "inner_lin": r["inner_lin"],
                        "inner_diff": r["inner_diff"],
                        "n_reject": r["n_reject"], "z": r["z"],
                        "worst_key": r["worst_key"],
                        "min_fisher": r["min_fisher"], "keep": r["keep"]})
        done.add(idx)
        print("draw %d/%d kind=%s inner=%s/%s fisher_nrej=%d z=%.1f %s"
              % (idx - lo + 1, hi - lo, rec["kind"], r["inner_lin"],
                 r["inner_diff"], r["n_reject"], r["z"],
                 "KEEP" if r["keep"] else "kill"), flush=True)
        if len(records) % 25 == 0:
            tmp = a.out + ".tmp"
            with open(tmp, "w") as f:
                json.dump({"pool": a.pool_file, "shard_idx": a.shard_idx,
                           "num_shards": a.num_shards, "slice": [lo, hi],
                           "canary": {"ok": True, "detail": can_detail},
                           "records": records}, f)
            os.replace(tmp, a.out)
    res = {"pool": a.pool_file, "shard_idx": a.shard_idx,
           "num_shards": a.num_shards, "slice": [lo, hi],
           "canary": {"ok": True, "detail": can_detail}, "records": records}
    tmp = a.out + ".tmp"
    with open(tmp, "w") as f:
        json.dump(res, f)
    os.replace(tmp, a.out)
    keeps = sum(1 for r in records if r["keep"])
    print("SHARD %d DONE: %d/%d KEEP canary_ok=true" %
          (a.shard_idx, keeps, len(records)), flush=True)


if __name__ == "__main__":
    main()
