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
