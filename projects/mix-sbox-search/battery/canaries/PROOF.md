# Canary bit-exactness proofs (2026-10-05)

Canary bundles must be self-contained (no VM absolute paths) for GHA, so
their `model.py` files are fresh numpy implementations. Each was proven
bit-exact against the VM ground truth:

## Stage-specific kill verdicts (2026-10-05 fix)

`common.run_canaries` accepts a `kill_verdicts` parameter because stages
use different kill verdicts:
- S0: FAIL (delay kill) | S1: **FLAG** (not KILL/FAIL) | S2: KILL | S3: KILL | S4: FAIL.
- The S1 kill canary (r23 1-round) emits FLAG; the driver passes `kill_verdicts=("FLAG",)`.
- Bug caught 2026-10-05: the original hardcoded `("KILL","FAIL")` would have quarantined a correct S1 kill canary as a mismatch.

## S3 positive-evidence requirement (2026-10-05 fix)

PractRand can emit only `error reading from file` (pipe/load flake) and exit 0.
The S3 driver requires positive evidence: actual `length=` test output PLUS
either "no anomalies" or `FAIL`. No test output → INFRA_FAIL (retried 3×), never PASS.
Bug caught 2026-10-05: an identity-kill run once returned PASS with zero test output.

## r23 4-round (pass canary)

- Source: `r23_model_template.py` with `ROUNDS = 4`.
- Proof 1: 100,000 random (v,k) vectors — template `mix_np` vs the C
  `impl.c` (compiled from the VM's
  `fast_battery/candidates/r23_spn_mix4r_pba19b01_nw/impl.c`, `--hash`
  mode): **100000/100000 match**.
- Proof 2: `meta.json` goldens — `mix(0, 0x123456789abcdef0) =
  0xa6ed8cf09fdb3fa0` and `mix(0x123456789abcdef0, 0) =
  0xf7c1e1644334bdc4`: **both OK**.
- (The VM's own `model.py` could not be used as reference — its
  `specs25.json` lookup no longer contains this candidate; the C impl is
  the stronger ground truth anyway, being what xcheck validates.)

## r23 1-round (S1 kill canary)

- `impl.c`: hand-written 1-round variant (same S-box/P-box/LIN/key
  schedule, loop bound 1; pbox emitted as a loop over
  `p=(19*i+1)&63`, verified against the unrolled VM form by the proof).
- Proof: 20,000 random vectors, template (`ROUNDS = 1`) vs C `--hash`:
  **20000/20000 match**.
- Kill proof (VM, 2026-10-05): temporary `fast_battery/candidates/`
  entry run through the VM's own `src/s1_screen.py` →
  **FLAG, LIN 10/10 + DIFF 10/10 rejected** (worst|z| ~ 1.5M).
  Temp entry removed afterwards.

## mul64 (S0 delay-kill canary)

- `mix(v,k) = v*k`; `rtl.v` is the multiplier with `(val,key,out)` ports.
- Expected S0 verdict: FAIL — delay ratio exactly 1.0 > 0.9 bar.
- C/Python trivially identical by construction.

## identity (S2/S3/S4 kill canary)

- `mix(v,k) = v`.
- S2: every delta's output differences repeat massively → KILL by the
  repeats gate (verified in smoke test).
- S3: PractRand FAIL at 16MB (verified in smoke test).
- S4: SmallCrush fails immediately (verified in smoke test).
