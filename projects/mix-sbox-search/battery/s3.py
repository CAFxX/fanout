#!/usr/bin/env python3
"""S3 driver: PractRand incremental with early kill (GHA battery).

Bit-exact port of the fast_battery S3 stage (src/s3_gen.c + s3_practrand.sh):
compiles vendor/s3_gen.c + bundle impl.c, gates on the C-vs-Python xcheck
(10k vectors), then pipes the generator DIRECTLY into RNG_test stdin64
(no intermediate files) at 16MB -> 64MB -> 256MB -> 1GB, stopping at the
first FAIL. Plus the strided run (2nd practrand key, golden-ratio stride).

RNG_test comes from $PRACTRAND_BIN, else battery/third_party/install/bin
(built on demand by third_party/build_third_party.sh).

GHA rule: S3 on GHA supports deep breaking-point runs via --min-bytes/--max-bytes
(up to 32GB). Long PractRand runs (even beyond 1GB) go to the GHA fanout,
not the VM (his directive 2026-10-06).

Golden canaries (run first; stage refused on mismatch):
  pass: r23_spn_mix4r_pba19b01_nw -> PASS at --canary-bytes (16MB+64MB)
  kill: identity                     -> KILL (FAIL at 16MB)

Usage:
  s3.py --candidates-dir DIR --out OUT_DIR [--shard-idx I --shard-count N]
        [--candidate NAME] [--skip-canary] [--max-bytes B]
        [--canary-bytes B] [--third-party-dir DIR]
"""
import argparse
import json
import os
import re
import subprocess
import sys
import time

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import common

XCHECK_VECS = 10000
STAGES = [(16777216, "16MB"), (67108864, "64MB"),
          (268435456, "256MB"), (1073741824, "1GB"),
          (2147483648, "2GB"), (4294967296, "4GB"),
          (8589934592, "8GB"), (17179869184, "16GB"),
          (34359738368, "32GB")]
STRIDE_GOLDEN = "0x9E3779B97F4A7C15"


def parse_args():
    ap = argparse.ArgumentParser()
    ap.add_argument("--candidates-dir", required=True)
    ap.add_argument("--out", required=True)
    ap.add_argument("--shard-idx", type=int, default=int(
        os.environ.get("SHARD_IDX", 0)))
    ap.add_argument("--shard-count", type=int, default=int(
        os.environ.get("SHARD_COUNT", 1)))
    ap.add_argument("--candidate", default="")
    ap.add_argument("--only", default="",
                        help="comma-separated candidate allow-list")
    ap.add_argument("--skip-canary", action="store_true")
    ap.add_argument("--max-bytes", type=int, default=1073741824)
    ap.add_argument("--min-bytes", type=int, default=0,
                    help="skip stages smaller than this (deep runs: start "
                         "where a previous run stopped)")
    ap.add_argument("--canary-bytes", type=int, default=67108864)
    ap.add_argument("--third-party-dir", default="")
    return ap.parse_args()


def rng_test_bin(third_party_dir):
    if os.environ.get("PRACTRAND_BIN"):
        return os.environ["PRACTRAND_BIN"]
    tp = third_party_dir or os.path.join(common.BATTERY_HOME,
                                         "third_party")
    inst = os.path.join(tp, "install", "bin", "RNG_test")
    if not os.path.exists(inst):
        r = subprocess.run(
            ["bash", os.path.join(tp, "build_third_party.sh"),
             os.path.join(tp, "install"), "practrand"],
            capture_output=True, text=True, timeout=1800)
        # Defense in depth (2026-10-06): build_third_party.sh once masked
        # build failures as exit 0, so a zero rc alone does not prove the
        # binary exists. Verify the artifact, not just the return code.
        if r.returncode != 0 or not os.path.exists(inst):
            raise RuntimeError(
                f"practrand build failed (rc={r.returncode}, "
                f"binary_present={os.path.exists(inst)}):\n"
                f"--- stdout ---\n{r.stdout[-3000:]}\n"
                f"--- stderr ---\n{r.stderr[-3000:]}")
    return inst


def run_stage(gen_bin, rng_bin, key, nbytes, stride, c_out, tag):
    """One PractRand stage; returns (killed: bool, log excerpt, infra_ok).

    Silent-verdict guard: a PASS requires evidence that PractRand actually
    ran tests ("length=" + "no anomalies" lines). If RNG_test dies on
    startup ("error reading from file" with no test output — a pipe/load
    flake), the stage is INFRA_FAIL, never PASS.

    Timeout scales with size: ~10 min per GB (measured ~3.5 min/GB at 1GB;
    PractRand's suite grows with input, so 3x headroom), min 2h.

    Implementation (2026-10-06): pipe via bash, not Python subprocess.
    Python's subprocess.PIPE stdin redirection silently delivers EOF to
    RNG_test in the GHA job container (works at image build time, fails
    at job runtime). Bash pipes work in both contexts.
    """
    # Sanity check: verify the generator produces output.
    try:
        t = subprocess.run([gen_bin, key, "1024", stride],
                           capture_output=True, timeout=30)
        if len(t.stdout) != 1024 or t.returncode != 0:
            raise RuntimeError(
                f"gen sanity check failed: expected 1024 bytes, got "
                f"{len(t.stdout)}, rc={t.returncode}, "
                f"stderr={t.stderr[:500]!r}")
    except RuntimeError:
        raise
    except Exception as e:
        raise RuntimeError(f"gen sanity check exception: {e}")
    timeout_s = max(7200, int(nbytes / (1024**3) * 600))
    log = os.path.join(c_out, f"s3_practrand_{tag}.log")
    err = os.path.join(c_out, f"s3_gen_{tag}.err")
    # Do the entire stage in bash (2026-10-06): Python subprocess pipe
    # redirection silently delivers EOF to RNG_test in the GHA job
    # container (docker run --user $RUID). Bash pipes work at image build
    # time; test if they work at job runtime.
    import shlex
    tmpf = f"/tmp/_gen_{tag}.bin"
    # Step 1: gen to file via bash.
    cmd1 = (f"{shlex.quote(gen_bin)} {shlex.quote(key)} {nbytes} "
            f"{shlex.quote(stride)} > {shlex.quote(tmpf)} "
            f"2>{shlex.quote(err)}; echo \"gen_exit:$?\"; ls -la {shlex.quote(tmpf)}")
    r1 = subprocess.run(["bash", "-c", cmd1], capture_output=True, text=True,
                        timeout=timeout_s)
    print(f"[s3_diag] step1: {r1.stdout.strip()}", flush=True)
    # Step 2: RNG_test on file via bash.
    cmd2 = (f"{shlex.quote(rng_bin)} {shlex.quote(f'file64({tmpf})')} "
            f"> {shlex.quote(log)} 2>&1")
    r2 = subprocess.run(["bash", "-c", cmd2], timeout=timeout_s)
    # Cleanup.
    subprocess.run(["rm", "-f", tmpf], timeout=30)
    text = open(log).read()
    fails = [l for l in text.splitlines() if "FAIL" in l][:8]
    ran_tests = ("length=" in text and
                 ("no anomalies in" in text or "FAIL" in text))
    infra_ok = ran_tests or len(fails) > 0
    return (len(fails) > 0, "\n".join(fails), text[-500:], infra_ok)


def drive_one(bundle, c_out, rng_bin, max_bytes, min_bytes=0, canary_scope=False):
    os.makedirs(c_out, exist_ok=True)
    work = os.path.join(c_out, "_work")
    os.makedirs(work, exist_ok=True)
    t0 = time.time()
    out_dir = os.path.dirname(c_out)
    gen_bin = os.path.join(work, "gen")
    try:
        common.compile_c(
            [os.path.join(common.VENDOR, "s3_gen.c"), bundle["impl_c"]],
            gen_bin)
    except RuntimeError as e:
        return common.write_result(
            out_dir, bundle["name"], "s3", "INFRA_FAIL",
            {"signal": f"compile failed: {e}"})
    ok, msg = common.xcheck_with_bundle_dir(
        bundle, gen_bin, XCHECK_VECS, work)
    if not ok:
        return common.write_result(
            out_dir, bundle["name"], "s3", "INFRA_FAIL",
            {"signal": f"xcheck C-vs-Python mismatch: {msg}"})

    stages = [(nb, nm) for nb, nm in STAGES if min_bytes <= nb <= max_bytes]
    keys = bundle["meta"]["practrand_keys"]
    signal, verdict, stage = "all stages: no FAIL", "PASS", None

    def run_stage_retry(nbytes, nm, key, stride, tag):
        """run_stage with up to 3 attempts on INFRA_FAIL (pipe flake)."""
        last = None
        for attempt in range(3):
            last = run_stage(gen_bin, rng_bin, key, nbytes, stride,
                             c_out, f"{tag}_try{attempt}" if attempt else tag)
            killed, fails, _, infra_ok = last
            if infra_ok:
                return last
            time.sleep(5)
        return last

    # main key, sequential
    for nbytes, nm in stages:
        killed, fails, _, infra_ok = run_stage_retry(
            nbytes, nm, keys[0], "1", nm)
        if not infra_ok:
            return common.write_result(
                out_dir, bundle["name"], "s3", "INFRA_FAIL",
                {"signal": f"RNG_test produced no test output at {nm} "
                           f"after 3 attempts — see s3_practrand_{nm}*.log"})
        if killed:
            verdict, stage = "KILL", nm
            signal = f"KILL@{nm} (key={keys[0]} stride=1):\n{fails}"
            break
    # strided class (2nd key, golden-ratio stride)
    if verdict == "PASS" and len(keys) > 1 and not canary_scope:
        for nbytes, nm in stages:
            killed, fails, _, infra_ok = run_stage_retry(
                nbytes, nm, keys[1], STRIDE_GOLDEN, f"strided_{nm}")
            if not infra_ok:
                return common.write_result(
                    out_dir, bundle["name"], "s3", "INFRA_FAIL",
                    {"signal": f"RNG_test produced no test output at "
                               f"strided_{nm} after 3 attempts"})
            if killed:
                verdict, stage = "KILL", f"strided_{nm}"
                signal = (f"KILL@strided_{nm} (key={keys[1]} "
                          f"stride={STRIDE_GOLDEN}):\n{fails}")
                break
    return common.write_result(
        out_dir, bundle["name"], "s3", verdict,
        {"signal": signal, "failed_stage": stage,
         "max_bytes": max_bytes, "xcheck_vecs": XCHECK_VECS,
         "wall_s": round(time.time() - t0, 1),
         "native_result": "s3_practrand_*.log"})


def main():
    a = parse_args()
    # Absolute paths: vendored modules may os.chdir();
    # relative paths would silently break mid-run.
    a.candidates_dir = os.path.abspath(a.candidates_dir)
    a.out = os.path.abspath(a.out)
    common.setup_env()
    out_dir = os.environ.get("OUT_DIR", a.out)
    rng_bin = rng_test_bin(a.third_party_dir)

    if not a.skip_canary:
        can_root = os.path.join(common.BATTERY_HOME, "canaries")
        common.run_canaries(
            "s3",
            lambda b, o: drive_one(b, o, rng_bin, a.canary_bytes,
                                   canary_scope=True),
            can_root, os.path.join(out_dir, "_canary"),
            {"pass": "pass/r23_spn_mix4r_pba19b01_nw",
             "kill": "kill_s234/identity"})
        print("[s3] canaries OK", flush=True)

    if a.candidate:
        names = [a.candidate]
    else:
        names = common.list_candidates(a.candidates_dir, a.only)
    for n in common.shard_slice(names, a.shard_idx, a.shard_count):
        b = common.load_bundle(a.candidates_dir, n, need_rtl=False)
        rec = drive_one(b, os.path.join(out_dir, n), rng_bin,
                        a.max_bytes, a.min_bytes)
        print(f"[s3 {n}] {rec['verdict']}: "
              f"{rec['detail'].get('signal', '').splitlines()[0]}",
              flush=True)
    print("[s3] shard done", flush=True)


def main_wrapped():
    """Wrap main() so ANY crash writes an INFRA_FAIL record instead of
    dying silently with rc=1 and no output (the 2026-10-06 deep-run mystery:
    both shards rc=1, zero files, no diagnosis possible)."""
    import traceback
    try:
        main()
    except Exception as e:
        # Best-effort: write a crash record where the collect job looks.
        try:
            out_dir = os.environ.get("OUT_DIR", "/tmp")
            os.makedirs(out_dir, exist_ok=True)
            crash = {
                "candidate": "_driver",
                "stage": "s3",
                "verdict": "INFRA_FAIL",
                "detail": {
                    "signal": f"driver crash: {type(e).__name__}: {e}\n"
                              f"{traceback.format_exc()[-2000:]}"
                },
            }
            with open(os.path.join(out_dir, "_driver_s3.json"), "w") as f:
                json.dump(crash, f, indent=1)
        except Exception:
            pass
        # Also print to stderr for the workflow log.
        traceback.print_exc()
        sys.exit(1)


if __name__ == "__main__":
    main_wrapped()
