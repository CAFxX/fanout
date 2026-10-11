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
