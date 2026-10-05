# Third-party tools vendored for the GHA battery

## PractRand — CC0 1.0 Universal (public domain dedication) ✅

- Version: as found at `~/tools/practrand/PractRand` (upstream
  PractRand by Chris Doty-Humphrey).
- License: `practrand/doc/license.txt` (vendored full text) — "All
  PractRand source code and tools are freely usable for any purpose"
  under CC0 1.0 Universal. **Redistribution permitted**, including in
  this private repo and its GHA runners.
- Vendored subset: `src/`, `include/`, `tools/RNG_test.cpp` +
  the `tools/*.h` headers it needs (`RNG_from_name.h`,
  `TestManager.h`, `MultithreadedTestManager.h`,
  `Candidate_RNGs.h`, `multithreading.h`, `dummy_rng.h`,
  `SeedingTester.h`, `measure_RNG_performance.h`) — everything needed
  to build `RNG_test`; the full upstream tree also ships extra
  RNGs/tools we don't use.
- Built by `build_third_party.sh` → `$INSTALL/bin/RNG_test`.

## TestU01 — Apache License 2.0 ✅

- Version: `TestU01-2009-master` as found at `~/tools/testu01/`
  (autotools tree; upstream Pierre L'Ecuyer's TestU01).
- License: `testu01/LICENSE` (vendored) — Apache 2.0.
  **Redistribution permitted** provided: (a) recipients get a copy of
  the License (vendored here), (b) modified files carry change notices,
  (c) copyright/attribution notices are retained. We distribute it
  unmodified.
- Built by `build_third_party.sh` (`configure --prefix=$INSTALL &&
  make -j && make install`, ~10-20 min first build) →
  `$INSTALL/{include,lib}/libtestu01.*`. **Cache
  `$INSTALL` with actions/cache** (key: hash of `third_party/`) —
  never pay the build twice.

## Build

```bash
battery/third_party/build_third_party.sh [install_dir] [practrand|testu01|all]
```

Idempotent (skips products whose install marker exists). On GHA the
battery drivers call it with the actions/cache-restored install dir;
a cache miss builds once, then `actions/cache/save` stores it.
