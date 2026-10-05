#!/usr/bin/env python3
"""S3 driver: PractRand incremental with early kill (GHA battery).

Bit-exact port of the fast_battery S3 stage (src/s3_gen.c + s3_practrand.sh):
compiles vendor/s3_gen.c + bundle impl.c, gates on the C-vs-Python xcheck
(10k vectors), then pipes the generator DIRECTLY into RNG_test stdin64
(no intermediate files) at 16MB -> 64MB -> 256MB -> 1GB, stopping at the
first FAIL. Plus the strided run (2nd practrand key, golden-ratio stride).

RNG_test comes from $PRACTRAND_BIN, else battery/third_party/install/bin
(built on demand by third_party/build_third_party.sh).

GHA rule: S3 on GHA stops at 1GB (--max-bytes default 1073741824).
Breaking-point runs (>1GB) stay on the VM.

Golden canaries (run first; stage refused on mismatch):
  pass: r23_spn_mix4r_pba19b01_nw -> PASS at --canary-bytes (16MB+64MB)
  kill: identity                     -> KILL (FAIL at 16MB)

Usage:
  s3.py --candidates-dir DIR --out OUT_DIR [--shard-idx I --shard-count N]
        [--candidate NAME] [--skip-canary] [--max-bytes B]
        [--canary-bytes B] [--third-party-dir DIR]
"""
import argparse
import os
import re
import subprocess
import sys
import time

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import common

XCHECK_VECS = 10000
STAGES = [(16777216, "16MB"), (67108864, "64MB"),
          (268435456, "256MB"), (1073741824, "1GB")]
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
        if r.returncode != 0:
            raise RuntimeError(
                f"practrand build failed:\n{r.stderr[-2000:]}")
    return inst


def run_stage(gen_bin, rng_bin, key, nbytes, stride, c_out, tag):
    """One PractRand stage; returns (killed: bool, log excerpt, infra_ok).

    Silent-verdict guard: a PASS requires evidence that PractRand actually
    ran tests ("length=" + "no anomalies" lines). If RNG_test dies on
    startup ("error reading from file" with no test output — a pipe/load
    flake), the stage is INFRA_FAIL, never PASS.
    """
    log = os.path.join(c_out, f"s3_practrand_{tag}.log")
    err = os.path.join(c_out, f"s3_gen_{tag}.err")
    with open(err, "w") as ef:
        gen = subprocess.Popen(
            [gen_bin, key, str(nbytes), str(stride)],
            stdout=subprocess.PIPE, stderr=ef)
        with open(log, "w") as lf:
            rng = subprocess.run(
                [rng_bin, "stdin64"], stdin=gen.stdout,
                stdout=lf, stderr=subprocess.STDOUT, timeout=7200)
        gen.stdout.close()
        gen.wait(timeout=60)
    text = open(log).read()
    fails = [l for l in text.splitlines() if "FAIL" in l][:8]
    ran_tests = ("length=" in text and
                 ("no anomalies in" in text or "FAIL" in text))
    infra_ok = ran_tests or len(fails) > 0
    return (len(fails) > 0, "\n".join(fails), text[-500:], infra_ok)


def drive_one(bundle, c_out, rng_bin, max_bytes, canary_scope=False):
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

    stages = [(nb, nm) for nb, nm in STAGES if nb <= max_bytes]
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
        rec = drive_one(b, os.path.join(out_dir, n), rng_bin, a.max_bytes)
        print(f"[s3 {n}] {rec['verdict']}: "
              f"{rec['detail'].get('signal', '').splitlines()[0]}",
              flush=True)
    print("[s3] shard done", flush=True)


if __name__ == "__main__":
    main()
