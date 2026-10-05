#!/usr/bin/env python3
"""S4 driver: TestU01 SmallCrush (GHA battery).

Bit-exact port of the fast_battery S4 stage (src/s4_smallcrush.c):
compiles vendor/s4_smallcrush.c + bundle impl.c against TestU01,
gates on the C-vs-Python xcheck (10k vectors) plus the binary's own
golden-vector self-check, then runs bbattery_SmallCrush.

Verdict parsing (validated 2026-10-05 against real passing AND failing
TestU01 output — see AGENTS.md lesson):
  FAIL if /p-values? outside/ matches (TestU01 1.2.3 prints the plural
       "The following tests gave p-values outside [0.001, 0.9990]:");
  PASS iff the output contains "All tests were passed" (mutually exclusive
       with the failure pattern);
  else INFRA_FAIL (no pass confirmation).

TestU01 comes from $TESTU01_HOME, else
battery/third_party/install (built on demand by
third_party/build_third_party.sh; cache it with actions/cache — the
first build takes 10-20 min).

Golden canaries (run first; stage refused on mismatch):
  pass: xcheck + golden self-check on r23 (fast)
  kill: identity -> SmallCrush FAIL (fast)

Usage:
  s4.py --candidates-dir DIR --out OUT_DIR [--shard-idx I --shard-count N]
        [--candidate NAME] [--skip-canary] [--third-party-dir DIR]
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
FAIL_PAT = re.compile(r"p-values? outside")
PASS_PAT = "All tests were passed"


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
    ap.add_argument("--third-party-dir", default="")
    return ap.parse_args()


def testu01_home(third_party_dir):
    if os.environ.get("TESTU01_HOME"):
        return os.environ["TESTU01_HOME"]
    tp = third_party_dir or os.path.join(common.BATTERY_HOME,
                                         "third_party")
    inst = os.path.join(tp, "install")
    if not os.path.exists(os.path.join(inst, "lib", "libtestu01.a")):
        r = subprocess.run(
            ["bash", os.path.join(tp, "build_third_party.sh"),
             inst, "testu01"],
            capture_output=True, text=True, timeout=2400)
        if r.returncode != 0:
            raise RuntimeError(
                f"testu01 build failed:\n{r.stderr[-2000:]}")
    return inst


def s4_canary(tu_home, can_out):
    """Fast S4 canary (stage refused on mismatch).

    pass: r23 compiles against TestU01 (proves the link), the golden
          vector matches (same assertion s4_smallcrush.c makes before
          SmallCrush), and C==Python on 10k vectors via the s2 hash-mode
          binary. Seconds, not 45 minutes.
    kill: identity runs the real SmallCrush and must FAIL (fails fast).
    The full 45-minute r23 SmallCrush PASS is proven once per image build
    (see battery/README standing proof), not per shard.
    """
    can_root = os.path.join(common.BATTERY_HOME, "canaries")

    def canary_drive(bundle, c_out, fast_pass):
        os.makedirs(c_out, exist_ok=True)
        work = os.path.join(c_out, "_work")
        os.makedirs(work, exist_ok=True)
        golden = bundle["meta"]["goldens"][0]
        if fast_pass:
            gb = os.path.join(work, "golden")
            common.compile_c(
                [os.path.join(common.BATTERY_HOME, "s4_golden_check.c"),
                 bundle["impl_c"]],
                gb, extra=[f"-DGOLDEN_KEY={golden['key']}",
                           f"-DGOLDEN_VAL={golden['val']}",
                           f"-DGOLDEN_OUT={golden['out']}"])
            r = subprocess.run([gb], capture_output=True, text=True)
            if r.returncode != 0:
                return "FAIL", r.stdout.strip()
            hb = os.path.join(work, "hashbin")
            common.compile_c(
                [os.path.join(common.VENDOR, "s2_diffprof.c"),
                 bundle["impl_c"]],
                hb, extra=["-lm"])
            ok, msg = common.xcheck_with_bundle_dir(
                bundle, hb, XCHECK_VECS, work)
            if not ok:
                return "FAIL", f"xcheck: {msg}"
            return "PASS", "golden OK + xcheck 10k OK"
        return None

    b_pass = common.load_bundle(os.path.join(can_root, "pass"),
                                "r23_spn_mix4r_pba19b01_nw", need_rtl=False)
    b_kill = common.load_bundle(os.path.join(can_root, "kill_s234"),
                                "identity", need_rtl=False)
    v, s = canary_drive(b_pass, os.path.join(can_out, "pass"), True)
    print(f"[canary s4/pass] want=PASS got={v} -> "
          f"{'OK' if v == 'PASS' else 'MISMATCH'}", flush=True)
    if v != "PASS":
        raise common.CanaryFail(f"s4 pass canary: {s}; STAGE QUARANTINED")
    rec = drive_one(b_kill, os.path.join(can_out, "kill"), tu_home)
    print(f"[canary s4/kill] want=FAIL got={rec['verdict']} -> "
          f"{'OK' if rec['verdict'] == 'FAIL' else 'MISMATCH'}",
          flush=True)
    if rec["verdict"] != "FAIL":
        raise common.CanaryFail(
            f"s4 kill canary: got {rec['verdict']}; STAGE QUARANTINED")


def drive_one(bundle, c_out, tu_home):
    os.makedirs(c_out, exist_ok=True)
    work = os.path.join(c_out, "_work")
    os.makedirs(work, exist_ok=True)
    t0 = time.time()
    out_dir = os.path.dirname(c_out)
    golden = bundle["meta"]["goldens"][0]
    golden_key = golden["key"]
    golden_out = golden["out"]
    binary = os.path.join(work, "smallcrush")
    try:
        common.compile_c(
            [os.path.join(common.VENDOR, "s4_smallcrush.c"),
             bundle["impl_c"]],
            binary,
            extra=[f"-DGOLDEN_KEY={golden_key}",
                   f"-DGOLDEN_OUT={golden_out}",
                   f"-I{tu_home}/include", f"-L{tu_home}/lib",
                   "-ltestu01", "-lprobdist", "-lmylib", "-lm"])
    except RuntimeError as e:
        return common.write_result(
            out_dir, bundle["name"], "s4", "INFRA_FAIL",
            {"signal": f"compile failed: {e}"})
    # The smallcrush binary has no --hash mode; use a separate hashbin
    # (s2_diffprof.c) for the C-vs-Python cross-check.
    hashbin = os.path.join(work, "hashbin")
    try:
        common.compile_c(
            [os.path.join(common.VENDOR, "s2_diffprof.c"),
             bundle["impl_c"]],
            hashbin, extra=["-lm"])
    except RuntimeError as e:
        return common.write_result(
            out_dir, bundle["name"], "s4", "INFRA_FAIL",
            {"signal": f"hashbin compile failed: {e}"})
    ok, msg = common.xcheck_with_bundle_dir(
        bundle, hashbin, XCHECK_VECS, work)
    if not ok:
        return common.write_result(
            out_dir, bundle["name"], "s4", "INFRA_FAIL",
            {"signal": f"xcheck C-vs-Python mismatch: {msg}"})
    key = bundle["meta"]["practrand_keys"][0]
    env = dict(os.environ,
               LD_LIBRARY_PATH=tu_home + "/lib:" +
               os.environ.get("LD_LIBRARY_PATH", ""))
    log = os.path.join(c_out, "s4.log")
    with open(log, "w") as lf:
        r = subprocess.run([binary, key], stdout=lf,
                           stderr=subprocess.STDOUT,
                           timeout=3600, cwd=work, env=env)
    out = open(log).read()
    fails = FAIL_PAT.findall(out)
    if fails:
        bad = [l.strip() for l in out.splitlines()
               if FAIL_PAT.search(l)][:10]
        verdict, signal = "FAIL", "p-value(s) outside [0.001,0.999]:\n" \
            + "\n".join(bad)
    elif PASS_PAT in out and r.returncode == 0:
        verdict, signal = "PASS", "All tests were passed"
    else:
        verdict, signal = "INFRA_FAIL", \
            f"no pass confirmation (rc={r.returncode}): {out[-500:]}"
    return common.write_result(
        out_dir, bundle["name"], "s4", verdict,
        {"signal": signal, "key": key, "xcheck_vecs": XCHECK_VECS,
         "wall_s": round(time.time() - t0, 1),
         "native_result": "s4.log"})


def main():
    a = parse_args()
    common.setup_env()
    out_dir = os.environ.get("OUT_DIR", a.out)
    tu_home = testu01_home(a.third_party_dir)

    if not a.skip_canary:
        s4_canary(tu_home, os.path.join(out_dir, "_canary"))
        print("[s4] canaries OK", flush=True)

    if a.candidate:
        names = [a.candidate]
    else:
        names = common.list_candidates(a.candidates_dir, a.only)
    for n in common.shard_slice(names, a.shard_idx, a.shard_count):
        b = common.load_bundle(a.candidates_dir, n, need_rtl=False)
        rec = drive_one(b, os.path.join(out_dir, n), tu_home)
        print(f"[s4 {n}] {rec['verdict']}: "
              f"{rec['detail'].get('signal', '').splitlines()[0]}",
              flush=True)
    print("[s4] shard done", flush=True)


if __name__ == "__main__":
    main()
