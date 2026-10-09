#!/bin/bash
# Builds raidjit WASM modules on GHA. Shard contract: SHARD_IDX selects
# the build target (0 = disasm, 1 = aarch64). Outputs land under $OUT_DIR.
# Source tree arrives as projects/wasm-build/src/raidjit-wasm-src.tar.gz
# (slim tree: raidjit src+include+tools/wasm + Capstone subset for the
# disasm target). Reproducible: pinned emsdk image (see Dockerfile).
# Full build log is teed to $OUT_DIR/build-shard-$SHARD_IDX.log so
# failures are diagnosable from the committed results (no log API needed).
set -e
set -o pipefail
cd /work/projects/wasm-build/src
rm -rf raidjit-wasm
tar xzf raidjit-wasm-src.tar.gz
cd raidjit-wasm
source /emsdk/emsdk_env.sh 2>/dev/null || true
export PATH=/emsdk/upstream/emscripten:$PATH
export TMPDIR=/tmp
mkdir -p "$OUT_DIR"
IDX="${SHARD_IDX:-0}"
LOG="$OUT_DIR/build-shard-$IDX.log"
{
echo "=== shard $IDX starting $(date -u +%FT%TZ) ==="
emcc --version | head -1
if [ "$IDX" = "0" ]; then
  make -f tools/wasm/Makefile.wasm disasm
  cp tools/wasm/out/disasm.js "$OUT_DIR/"
  echo "shard 0: disasm built"
else
  make -f tools/wasm/Makefile.wasm aarch64
  cp tools/wasm/out/raidjit_aarch64.js "$OUT_DIR/"
  echo "shard $IDX: aarch64 built"
fi
echo "=== shard $IDX done ==="
ls -lh "$OUT_DIR/"
} 2>&1 | tee "$LOG"
