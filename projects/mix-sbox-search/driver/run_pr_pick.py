#!/usr/bin/env python3
"""Run P&R for ONE promote.py pick (generic-dispatch shard worker).

Env:
  SHARD_IDX   pick index into the picks JSON (0-based)
  PICKS_JSON  promote/picks_<date>.json path
  OUT_DIR     where to write pr_<idx>.json (default ./out)
  PDK_ROOT    volare PDK root (default ./pdk)
  OPENROAD_BIN (default "openroad" — provided by the pinned public image)

Flow: pick table -> sbox4 verilog (gen_sbox --lanes4) ->
  openroad -no_splash -exit flow/openroad_pr.tcl ->
  parse PR_MARKER_* sections + SPEF *D_NET caps ->
  $OUT_DIR/pr_<idx>.json with the stable pr record contract:
    {name, table_hex, unit_crit, liberty_crit_ps,
     pr: {eligible, status, backend, postroute_crit_ps, routed_area_um2,
          power_w, max_input_pin_cap_ff, corner, die_um, err}}

The public image is efabless/openlane:1.0.0-amd64@sha256:2a1618... (pinned).
NOT YET RUN (2026-10-05): first P&R dispatch validates end-to-end.
"""
import json
import os
import re
import subprocess
import sys
import tempfile

DRIVER = os.path.dirname(os.path.abspath(__file__))
REPO = os.path.dirname(DRIVER)
FLOW = os.path.join(REPO, "flow")
sys.path.insert(0, DRIVER)
sys.path.insert(0, FLOW)
import gen_sbox

PR_IMAGE = ("efabless/openlane:1.0.0-amd64@sha256:"
            "2a161827bc615796d60b5e3c7ce5b67ec463ea4ef4afb08ead7ee92fab972cff")
OPENROAD = os.environ.get("OPENROAD_BIN", "openroad")
TCL = os.path.join(FLOW, "openroad_pr.tcl")


def parse_checks(text):
    m = re.search(r"PR_MARKER_BEGIN_CHECKS(.*?)PR_MARKER_END_CHECKS",
                  text, re.S)
    if not m:
        return None, "checks marker missing"
    body = m.group(1)
    # try several OpenSTA phrasings for the max path delay
    for pat in (r"data arrival time\s+([\d.eE+-]+)",
                r"Data Arrival Time\s+([\d.eE+-]+)",
                r"Path Delay\s*:\s*([\d.eE+-]+)"):
        mm = re.search(pat, body)
        if mm:
            return float(mm.group(1)) * 1000.0, None  # ns -> ps
    return None, f"could not parse path delay from:\n{body[-800:]}"


def parse_power(text):
    m = re.search(r"PR_MARKER_BEGIN_POWER(.*?)PR_MARKER_END_POWER",
                  text, re.S)
    if not m:
        return None, "power marker missing"
    for line in m.group(1).splitlines():
        mm = re.match(r"\s*Total\s+([\d.eE+-]+)", line)
        if mm:
            return float(mm.group(1)), None
    return None, "Total line not found in report_power"


def parse_die_area(text):
    m = re.search(r"PR_MARKER_DIE_AREA\s+\{?([-\d.eE+ ]+)\}?\s+(\d+)",
                  text)
    if not m:
        return None, "die area marker missing"
    vals = [float(x) for x in m.group(1).split()]
    dbu = float(m.group(2))
    x1, y1, x2, y2 = vals[:4]
    return (x2 - x1) / dbu * (y2 - y1) / dbu, None


def parse_ports(text):
    m = re.search(r"PR_MARKER_BEGIN_PORTS(.*?)PR_MARKER_END_PORTS",
                  text, re.S)
    if not m:
        return []
    return re.findall(r"PORT_INPUT\s+(\S+)", m.group(1))


def max_input_cap(spef_path, ports):
    """Max total net capacitance (fF) over the input-port nets in the SPEF."""
    try:
        text = open(spef_path).read()
    except OSError:
        return None
    # *D_NET <net> <totcap>
    caps = {}
    for m in re.finditer(r"\*D_NET\s+(\S+)\s+([\d.eE+-]+)", text):
        caps[m.group(1)] = float(m.group(2))  # pF
    best = 0.0
    for p in ports:
        base = p.split("[")[0]
        for net, cap in caps.items():
            if net == p or net == base or net.startswith(base + "["):
                best = max(best, cap)
    return best * 1000.0 if best else None  # pF -> fF


def main():
    idx = int(os.environ.get("SHARD_IDX", 0))
    picks = json.load(open(os.environ["PICKS_JSON"]))["picks"]
    pick = picks[idx]
    out_dir = os.environ.get("OUT_DIR", "./out")
    os.makedirs(out_dir, exist_ok=True)
    name = f"pr_pick_{idx}"
    tab = pick["table"]

    work = tempfile.mkdtemp(prefix="pr_pick_")
    vpath = os.path.join(work, f"{name}.v")
    with open(vpath, "w") as f:
        f.write(gen_sbox.write_sbox4(tab, name))

    env = dict(os.environ, DESIGN_V=vpath, TOP=name,
               PDK_ROOT=os.environ.get("PDK_ROOT", "./pdk"),
               OUT_DIR=work)
    rec = {"name": name, "table_hex": pick["table_hex"],
           "unit_crit": pick.get("unit_crit"),
           "liberty_crit_ps": pick.get("liberty_crit_ps"),
           "pr": {"eligible": True, "status": "failed", "backend": PR_IMAGE,
                  "corner": "tt_100C_1v80", "postroute_crit_ps": None,
                  "routed_area_um2": None, "power_w": None,
                  "max_input_pin_cap_ff": None, "err": None}}
    try:
        r = subprocess.run([OPENROAD, "-no_splash", "-exit", TCL],
                           capture_output=True, text=True, timeout=3600,
                           env=env, cwd=work)
        out = r.stdout + r.stderr
        if "PR_DONE" not in out:
            raise RuntimeError(f"openroad TCL did not complete:\n{out[-3000:]}")
        crit_ps, e1 = parse_checks(out)
        power_w, e2 = parse_power(out)
        area, e3 = parse_die_area(out)
        ports = parse_ports(out)
        cap_ff = max_input_cap(os.path.join(work, "design.spef"), ports)
        errs = [e for e in (e1, e2, e3) if e]
        if errs:
            raise RuntimeError("; ".join(errs))
        rec["pr"].update({"status": "measured",
                          "postroute_crit_ps": crit_ps,
                          "routed_area_um2": area, "power_w": power_w,
                          "max_input_pin_cap_ff": cap_ff})
    except Exception as e:
        rec["pr"]["err"] = f"{type(e).__name__}: {e}"

    outp = os.path.join(out_dir, f"pr_{idx}.json")
    with open(outp, "w") as f:
        json.dump(rec, f, indent=1)
    st = rec["pr"]["status"]
    print(f"PR_PICK {idx} {pick['table_hex']}: {st} "
          f"crit={rec['pr']['postroute_crit_ps']}ps "
          f"err={rec['pr']['err']}", flush=True)


if __name__ == "__main__":
    main()
