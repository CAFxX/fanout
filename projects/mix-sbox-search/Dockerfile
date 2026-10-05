# Single small image for the mix-sbox-search staged pipeline.
# Base: mambaorg/micromamba (conda-forge yosys 0.69, pinned — same feedstock
# as the VM's pinned toolchain). The Sky130 PDK / liberty file is NEVER
# baked in: fetched at job start via driver/ensure_pdk.py (volare +
# actions/cache).
#
# Build (one-time, needs docker):
#   docker buildx build --push -t ghcr.io/<owner>/fanout-mix-sbox-search:latest \
#     -f projects/mix-sbox-search/Dockerfile .
# Or via the generic workflow: Actions -> build-image -> Run
#   (project=mix-sbox-search, dockerfile_path=projects/mix-sbox-search/Dockerfile)
#
# NOTE (2026-10-05): NOT BUILT YET — no docker on the dev VM
# (`command -v docker` empty). Sizes below are measured component data
# (registry/package APIs), not a `docker images` reading:
#   mambaorg/micromamba:2 base ............ 37.7 MB (Docker Hub API, full_size)
#   yosys 0.69 conda tarball .............. 22.5 MB (anaconda API, linux-64)
#   python 3.12 + deps (tcl/tk/graphviz/...) ~45 MB tarballs (conda-forge norms)
#   => installed toolchain ~170-200 MB, ~80-100 MB compressed in image
#   ESTIMATED IMAGE TOTAL ~130 MB -> fits the 500 MB GHCR free-tier quota
#   with wide margin. No PDK baked in (fetched at job start).
FROM mambaorg/micromamba:2 AS base

# Install yosys 0.69 (pinned) + python + battery toolchain into
# ${MAMBA_USER}'s base env.
# conda-forge only: the VM's yosys is the conda-forge 0.69 build; the
# version assert in flow/run_synth.py ("0.69" in `yosys -V`) guards drift.
# Battery needs: gcc (C harnesses), iverilog (S0 equiv check),
# numpy + scipy (S1 screen; pinned to the VM's versions).
RUN micromamba install -y -n base -c conda-forge \
        yosys=0.69 \
        python=3.12 \
        gcc \
        autoconf \
        curl \
        unzip \
        iverilog \
        numpy==1.26.4 \
        scipy==1.11.4 \
    && micromamba clean --all --yes

ENV YOSYS_BIN=/opt/conda/bin/yosys \
    IVERILOG_BIN=/opt/conda/bin/iverilog \
    VVP_BIN=/opt/conda/bin/vvp \
    PATH=/opt/conda/bin:$PATH

# Bake the pinned flow + drivers (source of truth for the staged pipeline).
COPY projects/mix-sbox-search/flow /opt/sbox/flow
COPY projects/mix-sbox-search/driver /opt/sbox/driver
COPY projects/mix-sbox-search/pools /opt/sbox/pools
# Bake the battery (vendored stage logic + canaries + third-party sources).
# Third-party tools (PractRand RNG_test, TestU01) are built at job start
# into battery/third_party/install and cached via actions/cache — not baked,
# so the image stays slim and the 10-20 min TestU01 build is paid once.
COPY projects/mix-sbox-search/battery /opt/sbox/battery

# Golden self-test on build: the pinned flow must reproduce MIDORI_Sb0 = 3.0u
# inside THIS image, else the image is quarantined (build fails).
RUN python3 /opt/sbox/flow/validate_sbox.py --self-test && \
    python3 - <<'EOF'
import json, subprocess, sys
sys.path.insert(0, "/opt/sbox/driver")
sys.path.insert(0, "/opt/sbox/flow")
import variant_lib as vl
import gen_sbox
tab = vl.CANDS["midori_sb0"]
open("/tmp/canary.v", "w").write(gen_sbox.write_sbox64(tab, "canary"))
r = subprocess.run(
    [sys.executable, "/opt/sbox/flow/run_synth.py",
     "/tmp/canary.v", "canary", "/tmp/canary_wd"],
    capture_output=True, text=True, timeout=900)
ok = False
for line in (r.stdout + r.stderr).split("\n"):
    if line.startswith("RESULT"):
        for tok in line.split():
            if tok.startswith("crit="):
                ok = (float(tok.split("=", 1)[1]) == 3.0)
if not ok:
    sys.exit(f"GOLDEN SELF-TEST FAILED:\n{r.stdout[-2000:]}\n{r.stderr[-2000:]}")
print("golden self-test OK: MIDORI_Sb0 == 3.0u in-image")
EOF

WORKDIR /work
CMD ["python3", "--version"]
