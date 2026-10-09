# Emscripten for raidjit WASM builds.
# MUST match the local verified toolchain (~/tools/emsdk = 6.0.11):
# Makefile.wasm pins -sSINGLE_FILE_BINARY_ENCODE=0, which only exists in
# emscripten >= 4.0.18 (SPEC.md). 4.0.11 fails with "non-existent setting".
FROM emscripten/emsdk:6.0.11
RUN apt-get update && apt-get install -y git python3 make && rm -rf /var/lib/apt/lists/
WORKDIR /work
