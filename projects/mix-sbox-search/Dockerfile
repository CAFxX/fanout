# Single small image for the mix-sbox-search staged pipeline.
# Base: mambaorg/micromamba (conda-forge yosys, python).
FROM mambaorg/micromamba:latest

# Install yosys + python + battery toolchain.
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

WORKDIR /work
CMD ["python3", "--version"]
