# GA / tweaking fitness contract (mix-sbox-search)

How the VM-side evolution loops submit individuals and get S0/S1/S2
scores back from the GHA battery.

APPROVAL ECONOMY: fitness is evaluated by a CAMPAIGN — one job spec
(one approval) whose matrix shards each run the full S0->S4 battery
with early stop per individual. For pure fitness (S0/S1/S2), pass
`--stages s0,s1,s2` to `battery/campaign.py` (or set the campaign to
stop after S2) so no shard pays for S3/S4. S3/S4 run only on
fitness-survivors via a full campaign.

## Submit: individual -> candidate bundle

An individual is a directory with exactly:

| file      | contents |
|-----------|----------|
| `impl.c`  | `uint64_t mix_hash(uint64_t v, uint64_t k);` — the round function. Must also answer `--hash` mode (see battery README § "impl.c contract") |
| `model.py`| numpy reference `mix(v, k)` + `mix_np(vs, ks)`; `META` dict with `name` |
| `meta.json`| `{"name","practrand_keys":[hex,hex],"goldens":[{"key":hex,"out":hex}, ...]}` |
| `rtl.v`   | `(val, key, out)` module, needed only for S0; omit for S1/S2-only fitness |

The agent stages bundles with
`pipeline.py sync-candidates --project mix-sbox-search --src <ga_dir>`
(which validates the contract), then commits+pushes them; the campaign's
`--candidates` list selects the generation's individuals.

## Campaign spec (via `pipeline.py run` / `fanout.py submit`)

* project: `mix-sbox-search`
* command: `python3 projects/mix-sbox-search/battery/campaign.py --candidates-dir projects/mix-sbox-search/battery/candidates --out $OUT_DIR --stages s0,s1,s2 --only ind_0042,ind_0043,...`
* env: `BATTERY_HOME=/work/projects/mix-sbox-search/battery`
* shards: `ceil(N / per_shard)` (per_shard=1 default; each shard honors
  `SHARD_IDX`/`SHARD_COUNT` and runs the full stage chain with early
  stop per individual).

## Result: scores back

After the campaign is `done`, `pipeline.py aggregate` writes
`campaign_summary.json`:

```json
{
  "n_total": 2, "n_survivors": 1,
  "survivors": ["ind_0042"], "killed": ["ind_0043"],
  "kill_at_stage": {"s0": 0, "s1": 1, "s2": 0, "s3": 0, "s4": 0},
  "campaigns": {
    "ind_0042": {"candidate": "ind_0042", "final_stage": "s2",
                 "final_verdict": "PASS", "survived": true,
                 "stages": {
                   "s0": {"verdict": "PASS", "signal": "27.9u"},
                   "s1": {"verdict": "CAL_PASS", "signal": ""},
                   "s2": {"verdict": "PASS",
                          "signal": "PASS max_repeats=1 max|z|=14.2"}}},
    "ind_0043": {"candidate": "ind_0043", "final_stage": "s1",
                 "final_verdict": "FLAG", "survived": false,
                 "stages": {
                   "s0": {"verdict": "PASS", "signal": "31.2u"},
                   "s1": {"verdict": "FLAG",
                          "signal": "FLAG LIN 10/10"}}}
  }
}
```

And `per_candidate.json`: `{candidate: {final_stage, final_verdict,
signal, survived, stages, job_id}}`.

## Fitness mapping (VM side)

| stage | pass verdict | fitness signal |
|-------|--------------|----------------|
| s0 | `PASS` | `delay_u` from `stages.s0.signal` (lower is better); `FAIL` → reject |
| s1 | `CAL_PASS` | 0 rejections; `FLAG` → reject. Near-miss ranking: re-run or parse the driver's `<ind>_s1.json` for `pb_rejected + red_rejected` (marginal counts; LIN/DIFF rejections are structural kills) |
| s2 | `PASS` | `max_repeats`/`max|z|` from `stages.s2.signal` (lower is better); `KILL` → reject |

Suggested scalar fitness for selection:
`fitness = delay_u + 1000 * (s1_rejections > 0) + 1000 * (s2_kill)`,
i.e. lexicographic: feasibility (S1/S2 pass) dominates, then S0 delay.
Tune the weights per experiment; the contract only guarantees the fields
above are present.

## Batching guidance

* One fitness campaign per generation (not per individual, not per
  stage): the campaign shards the candidate list; 20 concurrent jobs is
  the account cap. One campaign = one approval.
* Priority for GA batches: `evolution` (below `frontier-challenger` and
  `new-leads`, above `fuzz`) — see root README prioritization.
* The first candidate's run of each stage executes that stage's golden
  canary; a canary mismatch fails the shard (score = unavailable, do NOT
  treat as fitness 0).
* `INFRA_FAIL` verdicts are infra problems (compile/xcheck/toolchain),
  not fitness signals — fix the bundle or the image, don't select on them.
