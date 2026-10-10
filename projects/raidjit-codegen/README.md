# raidjit-codegen — bulk codegen for expert review campaign

Generates disassembly for ~2,137 raidjit kernel configs via the WASM
modules, sharded across GHA runners. Part of the comprehensive
generated-code review campaign (correctness, safety, optimality).

## Layout

- `driver.mjs` — shard-aware Node.js driver. Reads `configs.jsonl`,
  selects `$SHARD_IDX/$SHARD_COUNT` slice, compiles each via the WASM
  `rjw_compile` API, disassembles via Capstone `disasm.js`, writes
  `config_{tier}_{idx}.asm` + `config_{tier}_{idx}.json` to `$OUT_DIR`.
- `configs.jsonl` — 2,137 configs (Tier 1: 2002, Tier 2: 55, Tier 3: 80).
  Each line: `{tier, idx, isa_module, target, doc, meta}` where `doc` is
  the `rjw_compile` input and `meta` describes the pinned twin/topology.
- `wasm/` — prebuilt WASM modules (raidjit_x86.js, raidjit_aarch64.js,
  raidjit_riscv.js, disasm.js) from `~/workspace/raidjit-wasm/tools/wasm/out/`
  (build 2026-10-10a). Bundled so shards don't need emsdk.

## Config tiers

- **Tier 1** (~2,000): baseline + each of 12 twins pinned, for 7 CPU
  targets × 11 ops (encode/recover/parity_update) × U∈{1,8}. Invalid
  pins are gracefully rejected by the backend (recorded as `ok:false`).
- **Tier 2** (~55): twin interactions (nt×body, ternlog×body), LMUL
  sweep (RVV), U sweep.
- **Tier 3** (~80): degenerate topologies (raid0, raid1, rs(3,1)), LRC
  topologies, hand-picked invalid combos (expect rejection).

## Docker image

Uses `node:24-slim` directly (no custom build). Driver runs with
`node driver.mjs` from the tenant dir (`/work/projects/raidjit-codegen`).

## Output

Per config: `config_{tier}_{idx}.json` (always) with
`{tier, idx, isa_module, target, doc, meta, result}` where result is
`{ok, code_size, twins, uops_per_byte, ...}` or `{ok:false, err}`.
On success also `config_{tier}_{idx}.asm` (compact Capstone disassembly).
Plus `_shard_status.json` per shard.

## Review phase (local, after generation)

Automated analysis + expert reviewer subagents on the `.asm`/`.json`
files. See the campaign brief (parent agent) for the review checklist.
