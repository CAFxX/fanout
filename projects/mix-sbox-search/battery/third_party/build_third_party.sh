#!/bin/bash
# Build third-party tools for the GHA battery.
#   build_third_party.sh [install_dir] [which]
# which: practrand | testu01 | all (default all)
# Idempotent: skips a product whose install marker already exists.
# Designed for actions/cache on $install_dir (key: hash of third_party/
# sources); a cache hit skips the 10-20 min TestU01 build.
#
# Sources are downloaded from upstream if not vendored locally:
#   PractRand 0.96 (CC0): http://pracrand.sourceforge.net/PractRand_0.96.zip
#   TestU01 1.2.3 (Apache 2.0): http://simul.iro.umontreal.ca/testu01/TestU01.zip
set -u
HERE="$(cd "$(dirname "$0")" && pwd)"
INSTALL="${1:-$HERE/install}"
WHICH="${2:-all}"
mkdir -p "$INSTALL/bin" "$INSTALL/lib" "$INSTALL/include"

# Download and extract a source archive if the directory is missing.
# Args: $1 = dest dir name, $2 = URL, $3 = archive filename
ensure_source() {
  local dest="$HERE/$1" url="$2" arc="$3"
  [ -d "$dest" ] && return 0
  echo "[third_party] downloading $1 from upstream ..."
  local tmp="/tmp/third_party_dl_$1"
  mkdir -p "$tmp"
  curl -sSL -o "$tmp/$arc" "$url" || { echo "download failed: $url"; exit 1; }
  case "$arc" in
    *.zip) unzip -q "$tmp/$arc" -d "$tmp" || { echo "unzip failed"; exit 1; } ;;
    *.tgz|*.tar.gz) tar xzf "$tmp/$arc" -C "$tmp" || { echo "untar failed"; exit 1; } ;;
    *) echo "unknown archive type: $arc"; exit 1 ;;
  esac
  # Find the extracted top-level directory (may differ from $1)
  local extracted
  extracted=$(find "$tmp" -maxdepth 1 -type d ! -path "$tmp" | head -1)
  [ -n "$extracted" ] || { echo "no directory extracted"; exit 1; }
  mv "$extracted" "$dest" || { echo "move failed"; exit 1; }
  rm -rf "$tmp"
  echo "[third_party] $1 ready at $dest"
}

build_practrand() {
  [ -f "$INSTALL/bin/RNG_test" ] && \
    { echo "[third_party] RNG_test already built; skipping"; return 0; }
  ensure_source "practrand" \
    "http://pracrand.sourceforge.net/PractRand_0.96.zip" "PractRand_0.96.zip"
  echo "[third_party] building PractRand RNG_test ..."
  local src="$HERE/practrand"
  local b="$HERE/_build_practrand"
  rm -rf "$b"; mkdir -p "$b"; cd "$b"
  g++ -c "$src"/src/*.cpp "$src"/src/RNGs/*.cpp "$src"/src/RNGs/other/*.cpp \
      -O3 -I"$src/include" -Wno-constant-logical-operand \
      || { echo "practrand compile failed"; exit 1; }
  ar rcs PractRand.a ./*.o
  g++ -o "$INSTALL/bin/RNG_test" "$src/tools/RNG_test.cpp" PractRand.a \
      -O3 -I"$src/include" -I"$src/tools" -pthread \
      || { echo "RNG_test link failed"; exit 1; }
  cd "$HERE"; rm -rf "$b"
  "$INSTALL/bin/RNG_test" --help 2>&1 | head -2 || true
  echo "[third_party] RNG_test OK"
}

build_testu01() {
  [ -f "$INSTALL/lib/libtestu01.a" ] && \
    { echo "[third_party] TestU01 already built; skipping"; return 0; }
  ensure_source "testu01" \
    "http://simul.iro.umontreal.ca/testu01/TestU01.zip" "TestU01.zip"
  echo "[third_party] building TestU01 (10-20 min on first build) ..."
  local src="$HERE/testu01"
  local b="$HERE/_build_testu01"
  # The vendored source may carry a stale in-tree configure (from a
  # previous build attempt); distclean it so the out-of-tree build works.
  (cd "$src" && make distclean >/dev/null 2>&1) || true
  rm -rf "$b"; mkdir -p "$b"; cd "$b"
  "$src/configure" --prefix="$INSTALL" \
      || { echo "testu01 configure failed"; exit 1; }
  make -j"$(nproc)" || { echo "testu01 make failed"; exit 1; }
  make install || { echo "testu01 install failed"; exit 1; }
  cd "$HERE"; rm -rf "$b"
  ls "$INSTALL/lib"/libtestu01.* || { echo "libtestu01 missing"; exit 1; }
  echo "[third_party] TestU01 OK"
}

case "$WHICH" in
  practrand) build_practrand ;;
  testu01)   build_testu01 ;;
  all)       build_practrand; build_testu01 ;;
  *) echo "usage: build_third_party.sh [install_dir] [practrand|testu01|all]"; exit 1 ;;
esac
echo "[third_party] done -> $INSTALL"
