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

# Install yosys + python + battery toolchain into ${MAMBA_USER}'s base env.
# conda-forge only. Battery needs: gcc (C harnesses), iverilog (S0 equiv
# check), numpy + scipy (S1 screen).
RUN micromamba install -y -n base -c conda-forge \
        yosys \
        python \
        gcc \
        autoconf \
        curl \
        unzip \
        iverilog \
        numpy \
        scipy \
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

# Golden self-test on build: validate_sbox.py must pass its self-test
# inside THIS image, else the image is quarantined (build fails).
# (Full synthesis golden test disabled for now — validate separately.)
RUN python3 /opt/sbox/flow/validate_sbox.py --self-test && \
    echo "golden self-test OK: validate_sbox passes in-image"

WORKDIR /work
CMD ["python3", "--version"]
