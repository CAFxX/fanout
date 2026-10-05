#!/usr/bin/env python3
"""Ensure the Sky130 PDK (for the liberty-delay screen) is present.

The liberty file is NEVER baked into the image. This script fetches it at
job start via volare, with the download dir cached by actions/cache
(see README dispatch recipes).

Resolution order for the liberty file:
  1. $LIBERTY_LIB (explicit path)
  2. $PDK_ROOT/sky130A/libs.ref/sky130_fd_sc_hd/lib/sky130_fd_sc_hd__tt_100C_1v80.lib
  3. volare fetch: `pip install volare` (if needed), then
     `volare enable --pdk sky130A` into $PDK_ROOT (default ./pdk)

Prints the liberty path to stdout and exports it for callers that source
the output. Exits 0 even if the fetch fails (callers record liberty_err
instead of failing the shard); exits 2 only on unexpected errors.

NOTE: volare CLI usage below follows efabless/volare docs; the first real
GH dispatch will validate it end-to-end (not yet run as of 2026-10-05).
"""
import glob
import os
import subprocess
import sys

LIB_NAME = "sky130_fd_sc_hd__tt_100C_1v80.lib"
PDK = os.environ.get("VOLARE_PDK", "sky130A")


def find_liberty():
    if os.environ.get("LIBERTY_LIB") and os.path.exists(
            os.environ["LIBERTY_LIB"]):
        return os.environ["LIBERTY_LIB"]
    pdk_root = os.environ.get("PDK_ROOT", "./pdk")
    hits = glob.glob(os.path.join(
        pdk_root, "**", "sky130_fd_sc_hd", "lib", LIB_NAME), recursive=True)
    if hits:
        return hits[0]
    # case-insensitive fallback
    for root, _, files in os.walk(pdk_root):
        for fn in files:
            if fn.lower() == LIB_NAME.lower():
                return os.path.join(root, fn)
    return None


def ensure_volare(pdk_root):
    try:
        subprocess.run([sys.executable, "-m", "volare", "--help"],
                       capture_output=True, timeout=60, check=True)
    except Exception:
        print("installing volare via pip...", flush=True)
        subprocess.run([sys.executable, "-m", "pip", "install", "--quiet",
                        "volare"], check=True, timeout=600)
    env = dict(os.environ, PDK_ROOT=os.path.abspath(pdk_root))
    print(f"volare enable --pdk {PDK} (PDK_ROOT={env['PDK_ROOT']})...",
          flush=True)
    subprocess.run([sys.executable, "-m", "volare", "enable",
                    "--pdk", PDK], env=env, check=True, timeout=3600)


def main():
    try:
        lib = find_liberty()
        if lib is None:
            pdk_root = os.environ.get("PDK_ROOT", "./pdk")
            try:
                ensure_volare(pdk_root)
            except Exception as e:
                print(f"PDK_FETCH_FAILED: {type(e).__name__}: {e}")
                return 0
            lib = find_liberty()
        if lib is None:
            print("PDK_FETCH_FAILED: liberty file not found after volare run")
            return 0
        print(f"LIBERTY_LIB={os.path.abspath(lib)}")
        return 0
    except Exception as e:
        print(f"ensure_pdk unexpected error: {type(e).__name__}: {e}",
              file=sys.stderr)
        return 2


if __name__ == "__main__":
    sys.exit(main())
