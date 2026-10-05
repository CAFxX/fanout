# fanout — general-purpose GitHub-Actions compute fanout

Run CPU-heavy, embarrassingly-parallel batch work on GitHub Actions
runners instead of the dev VM. The generic layer knows nothing about any
project's domain; each project is a tenant under `projects/`.

## Approval economy (read first)

Every write through the GitHub App connector forces a **manual user
review**. Maximize offloaded work per approval:

1. **Batch aggressively.** `fanout.py submit-batch` pushes MULTIPLE job
   specs in ONE `push_files` call — one approval covers many jobs. The
   dispatch workflow fans out to every spec in the push.
2. **Chain stages inside jobs.** A campaign is ONE job spec whose matrix
   shards each run the FULL pipeline with early stop (e.g.
   `battery/campaign.py` runs S0->S4 per candidate). Per-stage dispatches
   (one approval per stage) are banned.
3. **Trigger builds on push.** Pushing `projects/<project>/Dockerfile`
   auto-rebuilds the image via `build-image.yml` — zero extra approvals.
4. **Initial repo push = 1 approval.** Push the entire tree with a single
   `push_files` call (many files), never per-file.

The hygiene rules below still apply.

## Layout

```
fanout/
  README.md                      this file
  lib/
    fanout.py                    agent-side helper (submit/poll/fetch/aggregate/budget)
    pipeline.py                  campaign helper (one spec -> full pipeline -> verdicts)
    budget.py                    hard-cap ledger + priority queue
  .github/workflows/
    build-image.yml              generic: docker build+push -> ghcr.io/<owner>/fanout-<project>
                                 (push trigger on projects/**/Dockerfile + manual)
    dispatch.yml                 generic: matrix over (spec x shard), docker-run per
                                 shard, per-spec collect commits results/
  projects/
    mix-sbox-search/             first tenant: RV64 `mix` S-box search (see its README)
```

## The loop (push/commit — no gh CLI, no PAT)

The authenticated `github` CLI exposes NO Actions APIs (no dispatch, no
run polling, no artifact download), and `gh` is not logged in. The loop
is push-triggered with results committed back:

```
submit:  agent writes jobs/<project>/<job-id>.json via
         `github call-tool --name push_files`
         -> dispatch.yml triggers on push (paths: ['jobs/**']),
            runs the (spec x shard) matrix; per-spec collect jobs commit
            results/<project>/<job-id>/ (shard-<i>/ + STATUS.json +
            FILES.json) via GITHUB_TOKEN
poll:    agent reads results/<project>/<job-id>/STATUS.json via
         `github call-read-tool --name get_file_contents` until
         done|partial|failed
fetch:   agent reads result files via get_file_contents into dest/
aggregate: fetch + run the project's aggregate script locally
```

`lib/fanout.py` wraps this. `lib/pipeline.py` builds campaign specs
(one spec = full pipeline over many candidates).

## Agent workflow

```bash
# 1. submit a campaign — PRINTS the estimate and REFUSES without --confirm
python3 lib/pipeline.py run --project mix-sbox-search \
  --candidates c1,c2,c3 \
  --image ghcr.io/<owner>/fanout-mix-sbox-search:latest \
  --priority new-leads --confirm
# this submit = 1 approval covering 1 job / ~180 estimated minutes

# 2. batch several independent jobs in ONE approval
python3 lib/fanout.py submit-batch --spec-files job1.json job2.json \
  --confirm
# this submit = 1 approval covering 2 jobs / ~Y estimated minutes

# 3. poll until done
python3 lib/fanout.py poll --project mix-sbox-search --job-id <id>

# 4. download + aggregate locally
python3 lib/pipeline.py aggregate --project mix-sbox-search \
  --job-id <id> --dest /tmp/campaign

# 5. budget
python3 lib/fanout.py budget
```

Every confirmed submit appends its estimate to `~/.fanout/ledger.json`.
Estimates are not metered actuals — check repo Settings > Billing for
ground truth.

## Shared budget

**2000 Linux-minutes/month (free plan) across ALL projects**, hard cap
**1800** (default; `FANOUT_CAP` env or `--cap` overrides). `submit`
REFUSES when the ledger projects over-cap. Priorities, highest first:

1. `frontier-challenger` — re-testing a live frontier claim;
2. `new-leads` — fresh candidates from directed search;
3. `evolution` — GA/tweaking fitness batches;
4. `fuzz` — stochastic fuzz.

`fanout.py queue` lists pending submits in priority order.

## Minute math (mix-sbox-search battery campaign, LABELED estimates)

Per-candidate unit costs (measured on the VM 2026-10-05 unless noted):

| stage | what | unit cost |
|-------|------|-----------|
| S0 | unit-delay synth (Yosys) + 1000-vec equiv | ~3-4 min |
| S1 | LIN/DIFF + marginal 468-cell screen | ~5-10 min |
| S2 | differential profiler, N=2^22 | ~3-6 min |
| S3 | PractRand 16MB->64MB->256MB->1GB, early stop | minutes if killed early; ~1-2h full 1GB |
| S4 | TestU01 SmallCrush | ~45 min |
| canary | golden canaries per shard (all stages) | ~20-30 min/shard |

Campaign estimate: `shards x per_shard x per_candidate_min`
(default 60 min/candidate: ~30 expected stage cost with early stop +
~30 canary overhead amortized per shard).

Worked funnel — 50 candidates through the full campaign:

| stage | candidates in | unit | stage cost |
|-------|---------------|------|------------|
| S0 | 50 | 3.5 min | ~175 min |
| S1 | 50 | 8 min | ~400 min |
| S2 | ~15 survive | 5 min | ~75 min |
| S3 | ~8 survive | ~30 min avg (early stop) | ~240 min |
| S4 | ~3 survive | 45 min | ~135 min |
| canary | 50 shards | ~25 min | ~1250 min (worst; 1 cand/shard) |

Realistic total ~1100-2300 min depending on shard packing and kill
profile — i.e. **one 50-candidate campaign is roughly one month's
budget**. Pack 2-4 candidates/shard (they usually die early; worst
~3h/candidate stays under the 5h cap) and prioritize ruthlessly:
frontier-challengers first, fuzz last. S3/S4 run only on survivors —
never spend 45-min SmallCrush on a candidate S1 would have killed.

## What STAYS on the VM (never offloaded)

- **PARETO decisions** — GHA results are leads, never verdicts.
- **Raw-log verification** — every kill/claim is verified against raw
  logs locally before it touches PARETO.md or the frontier.
- **S5 P&R calibration ground truth** — P&R stays manual/local (OpenLane
  image via the pinned public image on the VM).
- **Final RTL/docs** — the shipped artifacts are built and reviewed
  locally.
- **Breaking-point runs** — S3 beyond 1GB stays on the VM (GHA S3 stops
  at 1GB by rule).

## HYGIENE (shared repo, many projects — hard rules)

### Namespacing

| thing | pattern |
|-------|---------|
| job specs | `jobs/<project>/<job-id>.json` (never flat) |
| results | `results/<project>/<job-id>/` |
| STATUS | `results/<project>/<job-id>/STATUS.json` |
| cache keys | `fanout-<project>-<key>` |
| artifacts | `shard-<project>-<job-id>-<idx>` |
| images | `ghcr.io/<owner>/fanout-<project>:<tag>` |
| budget ledger | spend tracked per project |

### Trigger discipline

- `dispatch.yml` triggers ONLY on `push: paths: ['jobs/**']` (+ manual
  `workflow_dispatch`). Its collect job commits ONLY under
  `results/<project>/<job-id>/`, so its own push can never retrigger
  dispatch. (GitHub forbids `paths` + `paths-ignore` on one event; the
  positive filter + the collect write-scope IS the guard.)
- `build-image.yml` triggers ONLY on `push: paths:
  ['projects/**/Dockerfile']` (+ manual).
- **Hard rule:** any future per-project workflow MUST carry
  `paths: ['projects/<project>/**']`.

BAD (fires on every push to the repo — project A's push runs B's workflow):
```yaml
on:
  push:   # NO path filter: every push to main triggers this
```

GOOD:
```yaml
on:
  push:
    paths: ['projects/my-project/**']
  workflow_dispatch:
```

### Concurrency

Concurrency group per `(project, job-id)` —
`fanout-<project>-<job-id>` — so two dispatches for the same job never
clobber each other's results. `cancel-in-progress: false` (never kill a
running batch; concurrent runs queue).

### Least privilege

- All dispatch jobs: `contents: read`.
- ONLY the collect job: `contents: write` (its results commit).
- `build-image`: `packages: write` + `contents: read`.

### Invariant

**A push for project A never runs project B's workflows.** Enforced by
the namespaced path filters above; verify it when adding any workflow.

## Images

- Per-project images are built by `build-image.yml` and pushed to
  `ghcr.io/<owner>/fanout-<project>:<tag>` (needs `packages: write`).
  Pushing a Dockerfile change rebuilds automatically. GHCR free tier:
  ~500 MB storage — keep images slim; never bake datasets into them
  (fetch at job start + `actions/cache`, 10 GB quota).
- Pin by digest for reproducibility where it matters.

## Adding a second project

1. New dir `projects/<name>/` with:
   - `Dockerfile` (slim; toolchain only),
   - a shard-aware driver honoring `$SHARD_IDX`/`$SHARD_COUNT`/`$OUT_DIR`
     (resumable, idempotent),
   - an aggregate script (runs agent-side),
   - `README.md` (domain docs + campaign recipes).
2. Push the Dockerfile — the image builds automatically.
3. Campaign: `pipeline.py run --project <name> ...` (or `fanout.py
   submit` / `submit-batch` for bespoke jobs).
4. See `projects/mix-sbox-search/` for the reference implementation.

## Tenants

- `projects/mix-sbox-search/` — RV64 `mix` instruction S-box search:
  campaign driver runs the staged S0->S4 battery (Yosys unit-delay,
  statistical screen, differential profiler, PractRand, TestU01
  SmallCrush) with golden canaries and early stop. Its README has the
  full domain contract.

## Remaining validation

- [ ] True push-triggered dispatch validated on first real submit
      (setup/spec/collect path is verified locally on synthetic data;
      the GitHub-side trigger fires only on a real push).
- [ ] First campaign's canary + verdicts cross-checked against VM
      raw logs before any PARETO decision.
