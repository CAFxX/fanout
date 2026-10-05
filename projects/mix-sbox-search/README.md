# mix-sbox-search — tenant of the fanout repo

RV64 `mix` instruction S-box search: search, validate and measure 4-bit
S-box quality on all relevant dimensions, on GitHub Actions runners. The
VM keeps the in-context candidate batteries (S0→S4 + sac_z5 + BIC).

This project is the first tenant of `fanout/` (see `../../README.md` for
the generic layer, including the approval-economy principle and hygiene
rules). All mix-specific logic lives here.

Two compute tracks share this tenant:

1. **S-box search** (`driver/`, `flow/`, `pools/`) — S-box-level
   validate → unit-delay → liberty-delay screens + P&R on picks
   (documented below).
2. **Full battery campaign** (`battery/`) — the complete staged S0→S4
   battery (unit-delay synth, statistical screen, differential profiler,
   PractRand, TestU01 SmallCrush) run per candidate with early stop,
   one campaign = one approval (documented in § Battery campaign).

---

## Battery campaign

`battery/` ports `~/workspace/mix/exploration/fast_battery/` bit-exactly
to GHA. One campaign = one job spec = one approval; each matrix shard
runs the FULL S0→S4 pipeline for its candidate(s) with early stop at the
first failing stage (`battery/campaign.py`).

```
battery/
  s0.py s1.py s2.py s3.py s4.py   stage drivers (shard-aware, --only filter)
  campaign.py                     per-candidate S0->S4 with early stop
  aggregate.py                    agent-side: _campaign.json -> campaign_summary.json
  common.py                       env contract, xcheck, canary gate, sharding
  s4_golden_check.c               fast S4 pass-canary helper
  vendor/                         verbatim fast_battery sources + env-override
                                  patches only (see VENDOR_NOTES.md)
  canaries/                       golden canaries (self-contained bundles)
  candidates/                     candidate bundles synced from the VM
                                  (committed; GHA checks out the repo)
  third_party/                    PractRand + TestU01 sources (see LICENSES.md)
  ref/                            s2_kill_repeats.txt (= 10, from the VM)
```

### Bit-exactness

Stage logic is vendored verbatim from `fast_battery/src/`; the only
patches are environment-variable overrides (documented per-file in
`vendor/VENDOR_NOTES.md`). Critical flow fact: S0 uses
`abc -genlib mini.genlib` with the SAT-proven **no-dretime** script
(`abc_nodretime.script`), NOT plain `abc -genlib` — the PARETO S0
numbers are the no-dretime flow.

### impl.c contract

Every candidate bundle provides `impl.c` with
`uint64_t mix_hash(uint64_t v, uint64_t k)` and a `--hash` mode
(`--hash <vecs>`: read `v k` hex pairs on stdin, print `mix_hash` hex on
stdout). The drivers gate every C stage on a 10k-vector C-vs-Python
xcheck (`model.py` numpy reference). `model.py` also provides
`mix_np(vs, ks)` vectorized; `META["name"]` must match the bundle name.

### Golden canaries (red-green)

Each stage runs its canary before its candidates; a mismatch quarantines
the stage (the driver refuses to run):

| stage | pass canary | kill canary |
|-------|-------------|-------------|
| s0 | r23 4-round SPN → PASS (~28u) | mul64 → FAIL (ratio 1.0 > 0.9) |
| s1 | r23 → CAL_PASS | r23 1-round → FLAG (LIN+DIFF) |
| s2 | r23 → PASS | identity mix → KILL (repeats) |
| s3 | r23 @16MB+64MB, no FAIL | identity → FAIL @16MB |
| s4 | r23 golden + xcheck (fast) | identity → SmallCrush FAIL |

Canary bundles are self-contained (no VM absolute paths); bit-exactness
proofs in `canaries/PROOF.md` (100k/100k C-vs-numpy on r23).

### Third-party tools

| tool | version | license | redistribution |
|------|---------|---------|----------------|
| PractRand | upstream (Chris Doty-Humphrey) | CC0 1.0 Universal | ✅ allowed |
| TestU01 | TestU01-2009-master | Apache 2.0 | ✅ allowed (unmodified, LICENSE vendored) |

Built by `third_party/build_third_party.sh`; the install dir is cached
via `actions/cache` (key `fanout-mix-sbox-search-third-party`) — the
10-20 min TestU01 build is paid once. Full texts in
`third_party/LICENSES.md`.

### S4 verdict parsing (validated)

Per the 2026-10-05 lesson: FAIL iff `/p-values? outside/` matches
(TestU01 prints the plural "gave p-values outside"); PASS iff "All tests
were passed" is present (mutually exclusive); else INFRA_FAIL. Validated
against the real `fast_battery` r23 `s4.out` (pass) and the identity
canary (fail).

### Campaign usage

```bash
# stage VM candidates into battery/candidates/ (local)
python3 lib/pipeline.py sync-candidates --project mix-sbox-search \
  --src ~/workspace/mix/exploration/fast_battery/candidates
# commit + push (parent handles the push), then:
python3 lib/pipeline.py run --project mix-sbox-search \
  --candidates r23_spn_mix4r_pba19b01_nw,rasC2_5 \
  --image ghcr.io/<owner>/fanout-mix-sbox-search:latest \
  --priority new-leads --confirm
# poll, fetch, aggregate -> campaign_summary.json + per_candidate.json
```

Minute math, prioritization, and what stays on the VM: see the root
README. GHA S3 stops at 1GB (`--max-bytes`); breaking-point runs stay on
the VM. GA/tweaking fitness: see `FITNESS_CONTRACT.md`.

---

## Staged pipeline (per S-box candidate)

Inside the project image, each candidate goes through three stages:

(a) **validate** (`flow/validate_sbox.py`, pure Python): bijectivity;
    DDT → differential uniformity (gate DU=4); LAT → nonlinearity
    (gate NL=4); per-coordinate algebraic degree via ANF (gate: all 3);
    full coordinate dependency; fixed points; cycle structure;
    differential/linear branch numbers; boomerang uniformity.
    **No inverse metrics anywhere.**

(b) **unit-delay screen** (`flow/run_synth.py` — the vendored pinned flow:
    Yosys 0.69, `memory_map` before `techmap`, `abc -genlib mini.genlib`,
    dual-analyzer cross-check). Golden canary: every shard synthesizes the
    identity variant first and requires **MIDORI_Sb0 == 3.0u exactly**;
    mismatch quarantines the batch (aggregate fails loudly, no ranking).

(c) **liberty-delay screen** (`flow/liberty_screen.py` + `flow/sta_liberty.py`):
    second abc pass with `-liberty sky130_fd_sc_hd__tt_100C_1v80.lib`
    (typical corner), then a simplified liberty STA (bilinear table
    interpolation, slew propagated via transition tables, load = driven
    pin caps, **no wire capacitance**) giving a real cell-delay critical
    path in ps. The liberty file is fetched at job start via
    `driver/ensure_pdk.py` (volare + actions/cache) and is NEVER baked
    into the image.

Measured tier (typical corner, cell-delay only):

| box | unit crit | liberty crit |
|---|---|---|
| midori_sb0 | 3.0u | 267.4 ps |
| thf_blink_s0 | 3.4u | 333.2 ps |
| ulbc_s1 | 3.4u | 366.3 ps |
| s1 | 3.7u | 366.6 ps |

Note the liberty screen discriminates *within* the 3.4u tier (333 vs 366 ps)
— it is not redundant with unit delay.

Per-candidate record: `table`, full `validate` dict, `crit`/`cells`/`area`/
`fanout` (unit), `liberty_crit_ps`/`liberty_cells`/`liberty_area_um2`,
`pr` eligibility.

## P&R stage (path B)

Runs **only on `promote.py` picks** (candidates clearing every hurdle:
hard gates + unit < threshold + liberty < threshold), in a separate
dispatch using the **pinned public image**

```
efabless/openlane:1.0.0-amd64@sha256:2a161827bc615796d60b5e3c7ce5b67ec463ea4ef4afb08ead7ee92fab972cff
```

(Docker Hub, verified 2026-10-05; `1.0.0` is the official semver-tagged
stable release, same digest as `superstable`. Free to pull, costs no GHCR
storage. GHCR has no verifiable public efabless/openlane package.)
Our image is never used for P&R. `flow/openroad_pr.tcl` runs the
tiny-macro flow (floorplan, tapcell, place, route — no PDN mesh,
estimates only, DRC/LVS not a goal) and OpenSTA `report_checks` /
`report_power`; `driver/run_pr_pick.py` parses the markers + SPEF into the
stable `pr` record (`postroute_crit_ps`, `routed_area_um2`, `power_w`,
`max_input_pin_cap_ff`). **The TCL is draft — first P&R dispatch
validates it end-to-end.**

## Calibration plan

The liberty screen excludes wire capacitance by design. The parent will
compare the liberty-delay ranking against the VM's S5 real-P&R ground
truth; **if wire effects ever reorder the ranking, the P&R promotion
threshold tightens** (and the liberty STA gains a wireload term). Until
then, liberty-delay is a cell-delay second opinion, not silicon truth.

## Image (decision B+C)

Single small image: `mambaorg/micromamba:2` + `yosys=0.69` (conda-forge,
pinned — same feedstock as the VM) + `python=3.12` + baked `flow/` +
`driver/` + `pools/`. No PDK baked in. `Dockerfile` (project root);
build via Actions → `build-image` (project=mix-sbox-search) or
`docker buildx build --push -t ghcr.io/<owner>/fanout-mix-sbox-search:latest
-f projects/mix-sbox-search/Dockerfile .`

**NOT BUILT YET** (no docker on the dev VM). Measured component sizes:
base 37.7 MB (Docker Hub API) + yosys 22.5 MB tarball (anaconda API) +
python/deps ~45 MB tarballs → **~130 MB estimated total**, fits the
500 MB GHCR free-tier quota with wide margin. A golden self-test
(MIDORI_Sb0 == 3.0u in-image) runs at build time.

## Dispatch recipes (via `lib/fanout.py`)

```bash
# screen one box: 32 shards x ~20 min (validate + 2 synth passes)
python3 lib/fanout.py submit --project mix-sbox-search \
  --image ghcr.io/<owner>/fanout-mix-sbox-search:latest \
  --command 'eval "$(python3 /opt/sbox/driver/ensure_pdk.py)" && BOX=midori_sb0 python3 /opt/sbox/driver/screen.py --out "$OUT_DIR/shard.json"' \
  --env "BOX=midori_sb0" \
  --shards 32 --per-shard-min 20 --timeout-minutes 45 \
  --cache-key pdk-sky130 --cache-paths pdk \
  --confirm

# fuzz: seeds x samples (deterministic per seed)
python3 lib/fanout.py submit --project mix-sbox-search \
  --image ghcr.io/<owner>/fanout-mix-sbox-search:latest \
  --command 'eval "$(python3 /opt/sbox/driver/ensure_pdk.py)" && python3 /opt/sbox/driver/fuzz.py --n-samples 2000 --out "$OUT_DIR/fuzz.json"' \
  --shards 20 --per-shard-min 10 --timeout-minutes 30 \
  --cache-key pdk-sky130 --cache-paths pdk \
  --confirm

# aggregate agent-side (quarantine enforced here)
python3 lib/fanout.py aggregate <run_id> --dest /tmp/mix_shards \
  --script projects/mix-sbox-search/driver/aggregate_screen.py \
  -- --box midori_sb0 --shards-dir /tmp/mix_shards --out-dir results/

# promote picks for the parent (three-stage thresholds)
python3 projects/mix-sbox-search/driver/promote.py \
  --input results/ranked_midori_sb0_20261005.json
# -> promote/picks_<date>.json + extends results/promoted_registry.json

# P&R on picks (public image; one shard per pick)
python3 lib/fanout.py submit --project mix-sbox-search \
  --image "efabless/openlane:1.0.0-amd64@sha256:2a161827bc615796d60b5e3c7ce5b67ec463ea4ef4afb08ead7ee92fab972cff" \
  --command 'python3 projects/mix-sbox-search/driver/run_pr_pick.py' \
  --env "PICKS_JSON=promote/picks_20261005.json" \
  --shards <n_picks> --per-shard-min 8 --timeout-minutes 30 \
  --cache-key pdk-sky130 --cache-paths pdk \
  --confirm
```

Local smoke tests (real pinned flow on the VM):

```bash
python3 flow/validate_sbox.py --self-test
YOSYS_BIN=/home/hatch/tools/micromamba-root/envs/sky130/bin/yosys \
  LIBERTY_LIB=/home/hatch/.volare/volare/sky130/versions/c6d73a35f524070e85faff4a6a9eef49553ebc2b/sky130A/libs.ref/sky130_fd_sc_hd/lib/sky130_fd_sc_hd__tt_100C_1v80.lib \
  python3 driver/screen.py --box midori_sb0 --shard-idx 0 --num-shards 1024 \
  --out /tmp/smoke.json   # 4-variant slice; canary must read 3.0u
```

## Pipeline contract (standing)

The GH side searches, validates and measures at S-box level ONLY. A GH
fast rank is **never** a frontier claim. The parent pulls
`promote/picks_*.json`, plugs each picked S-box into the in-context
constructions (currently the 4-round pba19b01 SPN), and runs the full
S0→S1→S2→S3→S4 + sac_z5 + BIC battery on the VM — construction-specific
kills (e.g. the MIDORI Sb0 zero-key S1 DIFF flag) can only be proven
in-context. **Dedup rule:** never re-promote a table already in
`ROUND25_TABLE.md` or `pools/` (enforced by `promote.py` against
`pools/*.txt` + `results/promoted_registry.json`).

## Budget math (private repo, free tier = 2000 Linux-min/month, ALL projects)

- Screen: ~8–11 s/variant (validate ms + unit synth + liberty abc pass +
  liberty STA) → 128 variants/shard ≈ 18–24 min → 32 shards ≈ **600–750
  min per box sweep**.
- Fuzz: gate pass rate for random bijective 4-bit boxes is low single-digit
  %; 2000 samples/seed ≈ 6–10 min → 20 seeds ≈ **120–200 min**.
- P&R: ~3–8 min/pick on promoted picks only.
- The `fanout.py submit` estimate + ledger tracks spend across projects.
