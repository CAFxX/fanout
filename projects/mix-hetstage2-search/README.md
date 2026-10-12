# mix-hetstage2-search

GHA campaign tenant for the track-2 hetstage2 lane (RV64 `mix` search):
dense GF-mul-by-constant hetero stages + [CIRC2, MUXD×4] in the MU5 frame,
with the inner-DIFF-SEQ gate as the **primary** triage (the batch-62 lesson).

Parent: O131 (cheap hetero stages killed at triage; dense-C dose-response
pc 3→5→7: z=5280→250→9.0 with a residual truncated differential).
Pilot: `mix/exploration/sbox_screen/hetstage2/` — 9/36 KEEP (seed 0xB803):
4 dense-GF/H3 (pc 16/24/32/40, z 2.4–3.5) + 5 R5 [CIRC2, MUXD×4] (z 2.3–3.8).

## Layout

- `driver/screen_hetstage2.py` — per-shard driver. Shard contract:
  `$SHARD_IDX`/`$SHARD_COUNT` select a contiguous slice of the pool;
  outputs under `$OUT_DIR` (from env). Resumable (skips records already
  in the out file; atomic rename every 25 draws).
- `driver/hetstage2_eval.py` — evaluation core, **copied verbatim** from
  the VM pilot (`t8_hetstage2_screen.py`, het8 `rotl`/`mux`, hetstage
  `gf_mat`), never reimplemented. `evaluate()` returns inner + fisher
  verdicts; KEEP iff inner lin==PASS and diff==PASS and fisher
  n_reject==0.
- `driver/aggregate_hetstage2.py` — merges shard keeps, ranks by z;
  quarantines the campaign if any shard's canary failed.
- `vendored/` — `lin_diff.py`, `screen_lib.py` copied from the VM
  (instruments, not reimplemented).
- `pools/hetstage2_pool.jsonl` — 300 pre-drawn stage specs (150 dense-GF
  H1/H3, pc stratified 8..40; 150 R5 [CIRC2, MUXD×4]), generated locally
  by `pools/gen_hetstage2_pool.py` (seed 0xB804) with the pilot's draw
  operators. GHA evaluates; it never draws.
- `pools/canary.jsonl` — 9 pilot keeps (must reproduce inner-PASS +
  fisher-PASS, |z − expect_z| < 0.5) + hetstage try53 replay (must
  inner-FLAG). Every shard runs the canary first; a mismatch quarantines
  the shard (driver exits 2).

## Campaign

`hetstage2-screen-20261012`: 8 shards, est ~240 GHA-min, priority
`new-leads`. Image: the existing `ghcr.io/cafxx/fanout-mix-sbox-search`
(python + numpy + scipy) — no new image needed; the driver runs from the
repo checkout. Job spec: `jobs/mix-hetstage2-search/hetstage2-screen-20261012.json`.

## Rules followed

- GHA results are leads, never verdicts: keeps return to the VM for the
  4-gate bar → bundle → S0 → battery → Q1 chain; PARETO/kill judgments
  stay local.
- Bit-exactness: the eval port is copied from VM sources; the canary
  proves the port reproduces the pilot.
- Everything namespaced per project (`jobs/mix-hetstage2-search/`,
  `results/mix-hetstage2-search/`).
