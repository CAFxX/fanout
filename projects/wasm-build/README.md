# wasm-build — raidjit WASM module builds on GHA

Offloads the Emscripten builds of the raidjit WASM live-kernel-explorer
modules to GitHub Actions. Used when the VM is too loaded for local `emcc`
runs (emcc module builds hang at load ~37).

## Layout

- `Dockerfile` — `emscripten/emsdk:4.0.11` + git/python3/make. Pushing a
  change auto-rebuilds `ghcr.io/cafxx/fanout-wasm-build:latest`
  (build-image.yml).
- `driver.sh` — shard-aware build driver. Shard 0 builds `disasm`,
  shard 1 builds `raidjit_aarch64`. Honours `$SHARD_IDX`/`$OUT_DIR`.
- `src/raidjit-wasm-src.tar.gz` — slim source tree (generated from the
  `wasm-explorer` branch of the raidjit-wasm worktree):
  `src/` (all backends), `include/`, `tools/wasm/` (api sources +
  Makefile.wasm), and the Capstone subset compiled by the `disasm`
  target (`arch/{X86,AArch64,RISCV}`, `include/`, root `*.c`/`*.h`).
  Regenerate with the commands in `RESUME_STATE.md` of the raidjit-wasm
  worktree; the `.c` file list must stay identical to what
  `tools/wasm/Makefile.wasm` compiles (see the diff check in the
  generation notes).

## Campaign recipe

1. Push tenant files (Dockerfile/driver/README/src tarball) via PAT —
   triggers the image build.
2. Wait for `ghcr.io/cafxx/fanout-wasm-build:latest` to be fresh.
3. Push `jobs/wasm-build/<job-id>.json` (2 shards, 120 min timeout) —
   triggers dispatch.
4. Poll `results/wasm-build/<job-id>/STATUS.json` until `done`;
   download `shard-0/disasm.js` and `shard-1/raidjit_aarch64.js`.
5. Parity verdicts stay LOCAL: re-run the aarch64 parity harness and the
   disasm smoke test on the VM against the downloaded modules.
   GHA outputs are leads, never verdicts.

## Status

- 2026-10-09: tenant created, first build pending — job
  `wasm-rebuild-20261009a` (disasm.js, raidjit_aarch64.js) to be
  submitted after the image build completes.
