#!/bin/bash
# Build third-party tools for the GHA battery.
#   build_third_party.sh [install_dir] [which]
# which: practrand | testu01 | all (default all)
# Idempotent: skips a product whose install marker already exists.
# Designed for actions/cache on $install_dir (key: hash of third_party/
# sources); a cache hit skips the 10-20 min TestU01 build.
#
# Sources are downloaded at build time (not vendored) and verified by
# content hash before use.
#   PractRand 0.96 (CC0), via GitHub mirrors of the SourceForge release:
#     https://github.com/csc-lab/practrand (primary)
#     https://github.com/robang74/pracrand (fallback)
#   TestU01 1.2.3 (per its LICENSE), via GitHub mirrors:
#     https://github.com/umontreal-simul/TestU01-2009 (primary)
#     https://github.com/sysfce2/testu01 (fallback)
#
# 2026-10-06 incident: the old upstream URLs
# (pracrand.sourceforge.net, simul.iro.umontreal.ca) went dead and
# SourceForge mirrors returned truncated transfers, killing shards with
# rc=1 and no output. Fix: working mirror URLs, curl --fail --retry,
# and sha256 content verification (pinned 2026-10-06, cross-checked
# byte-identical across two independent mirrors per product).
set -u
HERE="$(cd "$(dirname "$0")" && pwd)"
INSTALL="${1:-$HERE/install}"
WHICH="${2:-all}"
mkdir -p "$INSTALL/bin" "$INSTALL/lib" "$INSTALL/include"

# Pinned content hashes (sha256 over the canonical source file list,
# null-delimited, LC_ALL=C sorted). Recompute with tree_hash() below.
# PractRand scope: exactly the build-consumed surface (src, include,
# tools/RNG_test.cpp) -- verified byte-identical across two GitHub
# mirrors AND the original upstream download (three-way match).
PRACTRAND_SHA256="b5449606816eee324031fced37b13f4ad377c90834f01e2aaa265765b6b14073"
PRACTRAND_SCOPE="src include tools/RNG_test.cpp"
PRACTRAND_URLS=(
  "https://github.com/csc-lab/practrand/archive/refs/heads/master.zip"
  "https://github.com/robang74/pracrand/archive/refs/heads/master.zip"
)
# TestU01 scope: full tree (both mirrors byte-identical).
TESTU01_SHA256="d42ed1bb398da0b17a32450423cfae2984bcc91a9fcc6ac8609c7237d73233f6"
TESTU01_SCOPE="."
TESTU01_URLS=(
  "https://github.com/umontreal-simul/TestU01-2009/archive/refs/heads/master.zip"
  "https://github.com/sysfce2/testu01/archive/refs/heads/master.zip"
)

# sha256 of a source tree's canonical file list.
# Args: $1 = dir, $2... = scope paths relative to dir.
tree_hash() {
  local dir="$1"; shift
  ( cd "$dir" && find "$@" -type f -print0 | LC_ALL=C sort -z \
      | xargs -0 sha256sum ) | sha256sum | cut -d' ' -f1
}

# Download and extract a verified source tree.
# Args: $1 = dest dir name, $2 = expected sha256, $3 = hash scope,
#       $4... = mirror URLs (tried in order).
# Verifies the existing dir too, so a poisoned actions/cache entry
# self-heals instead of failing the shard.
ensure_source() {
  local name="$1" want="$2" scope="$3"; shift 3
  local dest="$HERE/$name"
  # shellcheck disable=SC2086
  if [ -d "$dest" ] && [ "$(tree_hash "$dest" $scope)" = "$want" ]; then
    echo "[third_party] $name present and hash-verified; skipping download"
    return 0
  fi
  if [ -d "$dest" ]; then
    echo "[third_party] $name hash MISMATCH (stale/corrupt); re-downloading"
    rm -rf "$dest"
  fi
  local url tmp arc extracted
  tmp="/tmp/third_party_dl_$name"
  for url in "$@"; do
    echo "[third_party] downloading $name from $url ..."
    rm -rf "$tmp"; mkdir -p "$tmp"
    arc="${url##*/}"
    if curl -sSL --fail --retry 5 --retry-all-errors --retry-delay 3 \
            --connect-timeout 20 --max-time 900 \
            -o "$tmp/$arc" "$url"; then
      case "$arc" in
        *.zip) unzip -q "$tmp/$arc" -d "$tmp" \
                 || { echo "[third_party] unzip failed for $url"; continue; } ;;
        *.tgz|*.tar.gz) tar xzf "$tmp/$arc" -C "$tmp" \
                 || { echo "[third_party] untar failed for $url"; continue; } ;;
        *) echo "[third_party] unknown archive type: $arc"; continue ;;
      esac
      extracted=$(find "$tmp" -maxdepth 1 -type d ! -path "$tmp" | head -1)
      if [ -z "$extracted" ]; then
        echo "[third_party] no directory extracted from $url"
        continue
      fi
      rm -rf "$dest"; mv "$extracted" "$dest"
      # shellcheck disable=SC2086
      if [ "$(tree_hash "$dest" $scope)" = "$want" ]; then
        echo "[third_party] $name verified (sha256 match)"
        rm -rf "$tmp"
        return 0
      fi
      echo "[third_party] sha256 MISMATCH for $url; trying next mirror"
      rm -rf "$dest"
    else
      echo "[third_party] download failed: $url"
    fi
  done
  rm -rf "$tmp"
  echo "[third_party] FATAL: no mirror yielded a verified $name tree" >&2
  return 1
}

build_practrand() {
  [ -f "$INSTALL/bin/RNG_test" ] && \
    { echo "[third_party] RNG_test already built; skipping"; return 0; }
  ensure_source "practrand" "$PRACTRAND_SHA256" "$PRACTRAND_SCOPE" \
    "${PRACTRAND_URLS[@]}" || return 1
  echo "[third_party] building PractRand RNG_test ..."
  local src="$HERE/practrand"
  local b="$HERE/_build_practrand"
  rm -rf "$b"; mkdir -p "$b"; cd "$b"
  g++ -c "$src"/src/*.cpp "$src"/src/RNGs/*.cpp "$src"/src/RNGs/other/*.cpp \
      -O3 -I"$src/include" -Wno-constant-logical-operand \
      || { echo "practrand compile failed"; return 1; }
  ar rcs PractRand.a ./*.o
  g++ -o "$INSTALL/bin/RNG_test" "$src/tools/RNG_test.cpp" PractRand.a \
      -O3 -I"$src/include" -I"$src/tools" -pthread \
      || { echo "RNG_test link failed"; return 1; }
  cd "$HERE"; rm -rf "$b"
  "$INSTALL/bin/RNG_test" --help 2>&1 | head -2 || true
  echo "[third_party] RNG_test OK"
}

build_testu01() {
  [ -f "$INSTALL/lib/libtestu01.a" ] && \
    { echo "[third_party] TestU01 already built; skipping"; return 0; }
  ensure_source "testu01" "$TESTU01_SHA256" "$TESTU01_SCOPE" \
    "${TESTU01_URLS[@]}" || return 1
  echo "[third_party] building TestU01 (10-20 min on first build) ..."
  local src="$HERE/testu01"
  local b="$HERE/_build_testu01"
  # The vendored source may carry a stale in-tree configure (from a
  # previous build attempt); distclean it so the out-of-tree build works.
  (cd "$src" && make distclean >/dev/null 2>&1) || true
  rm -rf "$b"; mkdir -p "$b"; cd "$b"
  "$src/configure" --prefix="$INSTALL" \
      || { echo "testu01 configure failed"; return 1; }
  make -j"$(nproc)" || { echo "testu01 make failed"; return 1; }
  make install || { echo "testu01 install failed"; return 1; }
  cd "$HERE"; rm -rf "$b"
  ls "$INSTALL/lib"/libtestu01.* || { echo "libtestu01 missing"; return 1; }
  echo "[third_party] TestU01 OK"
}

case "$WHICH" in
  practrand) build_practrand ;;
  testu01)   build_testu01 ;;
  all)       build_practrand && build_testu01 ;;
  *) echo "usage: build_third_party.sh [install_dir] [practrand|testu01|all]"; exit 1 ;;
esac
echo "[third_party] done -> $INSTALL"
