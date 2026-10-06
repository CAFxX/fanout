# Single small image for the mix-sbox-search staged pipeline.
# Base: mambaorg/micromamba (conda-forge yosys, python).
FROM mambaorg/micromamba:latest

# Install yosys + python + battery toolchain.
# NOTE (2026-10-06): conda-forge splits C and C++ compilers: `gcc` is C
# only, `gxx` is C++. PractRand (third_party) needs g++; the canary/gen
# C code only needs gcc. Omitting gxx broke every GHA s3 PractRand build.
# Likewise the unprefixed binutils (`ar`, used to archive PractRand.a)
# come from the `binutils` metapackage, NOT from gcc's binutils_impl
# (which only ships prefixed x86_64-conda-linux-gnu-* tools).
RUN micromamba install -y -n base -c conda-forge \
        yosys \
        python \
        gcc \
        gxx \
        binutils \
        autoconf \
        curl \
        unzip \
        iverilog \
        numpy \
        scipy \
    && micromamba clean --all --yes

ENV PATH=/opt/conda/bin:$PATH

# Tool paths for the battery drivers (common.py defaults to the VM's
# absolute paths when these are unset — which do not exist in the image).
ENV YOSYS_BIN=/opt/conda/bin/yosys \
    IVERILOG_BIN=/opt/conda/bin/iverilog \
    VVP_BIN=/opt/conda/bin/vvp

# Bake the pinned flow + drivers.
COPY projects/mix-sbox-search/flow /opt/sbox/flow
COPY projects/mix-sbox-search/driver /opt/sbox/driver
COPY projects/mix-sbox-search/pools /opt/sbox/pools
COPY projects/mix-sbox-search/battery /opt/sbox/battery

# Golden self-test on build.
RUN python3 /opt/sbox/flow/validate_sbox.py --self-test && \
    echo "golden self-test OK"

# Create a user matching the GHA runner UID/GID (2026-10-06): the dispatch
# workflow runs shards via `docker run --user $RUID:$RGID` (1001:1001 on
# github-hosted runners). If the UID does not exist in /etc/passwd, file
# I/O via stdio (fopen/fread) silently fails in the shard container —
# RNG_test gets EOF on valid files. Creating the user fixes it.
RUN if ! id -u 1001 >/dev/null 2>&1; then \
      useradd -u 1001 -g 1001 -m -s /bin/bash runner 2>/dev/null || \
      useradd -u 1001 -m -s /bin/bash runner 2>/dev/null || true; \
    fi

WORKDIR /work
CMD ["python3", "--version"]
