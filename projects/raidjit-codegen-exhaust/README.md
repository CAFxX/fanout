# raidjit-codegen-exhaust — exhaustive twin-combination codegen (Carlo 2026-10-11)

No-pruning campaign: full Cartesian space over 128 structured twin states x
ISA variants (targets x isa pins x mul x LMUL x VL overrides) x U{1..8} x
14 topologies x encode/recover/parity_update patterns. Exact counts and the
per-cut ledger (every exclusion justified, C1-C6) live in the campaign dir:
~/workspace/raidjit/hidden_files/codegen_exhaust_campaign/inputs/
(counts.json, cuts_ledger.json, phases.json, manifest_schema.json).

driver.mjs: compile/disasm logic bit-exact from raidjit-codegen/driver.mjs;
compact output stage (results.jsonl with code_sha256/asm_sha256 +
asm_bundle.txt.gz per shard). Image: node:24-slim (no custom build).

Current job codegen-exhaust-e1x86-20261011a = phase E tranche for
encode/rs(6,3)/x86 variants only (27,648 configs, 32 shards). Later tranches
(aarch64, riscv, other topos, recover, parity_update) follow phases.json;
tranche ordering is execution order, not a space cut.

## Provenance / re-sync (2026-10-11, evidence-integrity fix)
- STALE (pre-sync) modules were byte-identical to raidjit-wasm out/ built
  from the wasm-explorer snapshot (sync 2026-10-10a), which lags current
  main (~/workspace/raidjit branch crc-fast @ 6dd183e) by the BUG_t1_80
  ternlog-phantom series, BUG_t1_529 xorbit-x86 ban, crc-fold-v2 and more
  (planid.c/plan.c/prune.c/api.c/jit.c/emit.c/ui_meta.c all differ; main
  adds crc_fast.c/h and x86_64/crc_fold.c/h). All pre-sync results are
  evidence about that snapshot ONLY.
- SYNCED modules (built 2026-10-11, Makefile.wasm overlay: wasm-explorer
  tools/wasm + current-main src/include wholesale, emcc 6.0.11):
  raidjit_x86.js sha256 4051132673dd5ee3ef1acbcafb21e70abf170c8a24aad2525fafb494dc6afbaa
  raidjit_aarch64.js sha256 36a00a18f4404ffa1f6835ce362b9fdb4e62a87d67a4c222fd6c5c52d00cf0f4
- raidjit_riscv.js is STALE (18e50d96...) and BLOCKED-EVIDENCE: current
  main does not compile for RISC-V (src/riscv/jit.c: 7 plan_need_regs
  call sites pass 7 args vs 8-arg plan.h:345 declaration). Do not submit
  riscv jobs as current-main evidence until main fixes it and the module
  is rebuilt.
- disasm.js unchanged (d5b556b4...): Capstone-only, embeds no raidjit
  src; a local rebuild hashes differently (toolchain drift), so the
  verified module is kept.
- driver.mjs sha256 93a1b29e...: BUG-COMBO-TRAP-002 fix — on any throw
  the WASM instance is dropped and the same config retried once on a
  fresh instance; only a fresh-instance throw is a trap verdict.
