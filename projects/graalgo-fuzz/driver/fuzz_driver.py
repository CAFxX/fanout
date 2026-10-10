#!/usr/bin/env python3
"""GHA shard driver for graalgo-fuzz.

Shard contract: honors $SHARD_IDX / $SHARD_COUNT / $OUT_DIR.
Each shard generates N programs from a disjoint seed range, runs the
differential, and writes $OUT_DIR/results.jsonl.

Env:
  SHARD_IDX, SHARD_COUNT : shard identity
  OUT_DIR                : output directory
  SEEDS_PER_SHARD        : programs per shard (default 50)
  GRAALGO_HOME           : /opt/graalgo (image default)

Golden canary: seed 999999 is a known-pass program; if it doesn't PASS,
the shard writes CANARY_FAIL and exits nonzero (batch quarantined).
"""
import json
import os
import subprocess
import sys
from concurrent.futures import ThreadPoolExecutor

GRAALGO_HOME = os.environ.get("GRAALGO_HOME", "/opt/graalgo")
FUZZ = os.path.join(GRAALGO_HOME, "phases/phase3/fuzzing")
OUT_DIR = os.environ.get("OUT_DIR", "/out")
SHARD_IDX = int(os.environ.get("SHARD_IDX", "0"))
SHARD_COUNT = int(os.environ.get("SHARD_COUNT", "1"))
SEEDS_PER_SHARD = int(os.environ.get("SEEDS_PER_SHARD", "50"))

# Canary: must PASS. (A tiny fixed program exercising int arithmetic.)
CANARY_SEED = 999999


def sh(cmd, timeout=300):
    return subprocess.run(cmd, shell=True, capture_output=True, text=True,
                          timeout=timeout)


def diff_one(gofile, workdir, gmp=1):
    r = sh(f"bash {FUZZ}/harness/diff_one.sh {gofile} {workdir} {gmp}",
           timeout=300)
    try:
        return json.loads(r.stdout.strip().split("\n")[-1])
    except Exception:
        return {"file": os.path.basename(gofile), "result": "HARNESS_ERROR",
                "detail": (r.stdout + r.stderr)[-500:]}


def main():
    os.makedirs(OUT_DIR, exist_ok=True)
    # Container env for the differential harness (diff_one.sh honors
    # GR/JAVA/TMPDIR when set; local defaults are $HOME-based).
    # Image layout: Go at /usr/local/go/bin (already on PATH via the
    # Dockerfile ENV), GraalVM at /opt/graalvm, sources at GRAALGO_HOME.
    os.environ["GR"] = GRAALGO_HOME
    os.environ["JAVA"] = "/opt/graalvm/bin/java"
    tmpdir = os.path.join(OUT_DIR, "tmp")
    os.makedirs(tmpdir, exist_ok=True)
    os.environ["TMPDIR"] = tmpdir
    # The container runs as the runner's uid/gid, which has no passwd entry
    # and no writable $HOME/.cache; keep the Go build/module caches on the
    # mounted $OUT_DIR (runner-owned, writable) so `go build` in diff_one.sh
    # cannot fail on GOCACHE permissions.
    go_cache = os.path.join(OUT_DIR, "gocache")
    os.makedirs(go_cache, exist_ok=True)
    os.environ["GOCACHE"] = go_cache
    go_path = os.path.join(OUT_DIR, "gopath")
    os.makedirs(go_path, exist_ok=True)
    os.environ["GOPATH"] = go_path
    # Pre-flight diagnostics (goes to the GHA job log).
    print(f"GRAALGO_HOME={GRAALGO_HOME}", flush=True)
    print(f"gen exists: {os.path.exists(os.path.join(FUZZ, 'gen/generator.py'))}",
          flush=True)
    print(f"harness exists: {os.path.exists(os.path.join(FUZZ, 'harness/diff_one.sh'))}",
          flush=True)
    pv = sh("command -v python3 && python3 --version")
    print(f"python3: rc={pv.returncode} {pv.stdout.strip()} {pv.stderr.strip()}",
          flush=True)
    gv = sh("command -v go && go version")
    print(f"go: rc={gv.returncode} {gv.stdout.strip()} {gv.stderr.strip()}",
          flush=True)
    jv = sh("/opt/graalvm/bin/java -version")
    print(f"java: rc={jv.returncode} {(jv.stdout + jv.stderr).strip().splitlines()[0] if (jv.stdout + jv.stderr).strip() else ''}",
          flush=True)
    # Shard i owns seeds [i*N, (i+1)*N) offset by a base to avoid overlap
    # with local runs (local uses 0..; GHA uses 1_000_000..).
    base = 1_000_000 + SHARD_IDX * SEEDS_PER_SHARD
    seeds = list(range(base, base + SEEDS_PER_SHARD))
    progdir = os.path.join(OUT_DIR, "programs")
    os.makedirs(progdir, exist_ok=True)

    print(f"shard {SHARD_IDX}/{SHARD_COUNT}: seeds {base}..{base + SEEDS_PER_SHARD - 1}",
          flush=True)

    # Canary first.
    canary_go = os.path.join(progdir, "canary.go")
    r = sh(f"python3 {FUZZ}/gen/generator.py --seed {CANARY_SEED} --out {canary_go}")
    if r.returncode != 0:
        print("CANARY_FAIL: generator error")
        print(f"--- generator stdout (tail) ---\n{r.stdout[-2000:]}", flush=True)
        print(f"--- generator stderr (tail) ---\n{r.stderr[-2000:]}", flush=True)
        sys.exit(2)
    crec = diff_one(canary_go, os.path.join(OUT_DIR, "work_canary"))
    print(f"canary: {crec}", flush=True)
    if crec["result"] != "PASS":
        with open(os.path.join(OUT_DIR, "CANARY_FAIL"), "w") as f:
            f.write(json.dumps(crec))
        print("CANARY_FAIL: known-pass program did not PASS — quarantining shard")
        sys.exit(2)

    # Main batch.
    for s in seeds:
        out = os.path.join(progdir, f"fuzz_{s}.go")
        r = sh(f"python3 {FUZZ}/gen/generator.py --seed {s} --out {out}")
        if r.returncode != 0:
            print(f"generator failed for seed {s}")
            print(f"--- stdout (tail) ---\n{r.stdout[-1000:]}", flush=True)
            print(f"--- stderr (tail) ---\n{r.stderr[-1000:]}", flush=True)
            sys.exit(1)

    results_path = os.path.join(OUT_DIR, "results.jsonl")
    with open(results_path, "w") as rf:
        def run_one(s):
            workdir = os.path.join(OUT_DIR, "work", f"seed_{s}")
            os.makedirs(workdir, exist_ok=True)
            gofile = os.path.join(progdir, f"fuzz_{s}.go")
            rec = diff_one(gofile, workdir)
            rec["seed"] = s
            rec["shard"] = SHARD_IDX
            return rec

        with ThreadPoolExecutor(max_workers=4) as ex:
            for i, rec in enumerate(ex.map(run_one, seeds)):
                rf.write(json.dumps(rec) + "\n")
                rf.flush()
                if (i + 1) % 10 == 0:
                    print(f"  {i+1}/{len(seeds)}", flush=True)

    print("shard done", flush=True)


if __name__ == "__main__":
    main()
