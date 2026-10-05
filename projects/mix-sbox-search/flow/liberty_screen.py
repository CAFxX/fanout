#!/usr/bin/env python3
"""Liberty-delay screen: second opinion on S-box delay in REAL cell delay.

Runs the pinned flow shape but maps with `abc -liberty <sky130 tt lib>`
(typical corner) instead of `abc -genlib mini.genlib`, then computes the
critical path with flow/sta_liberty.py (simplified liberty STA, ps).

This is the staged pipeline's stage (c): validate -> unit-delay screen ->
liberty-delay screen. The liberty file is NEVER baked into the image; it is
fetched at job start (driver/ensure_pdk.py, volare + actions/cache) and
passed via --liberty / $LIBERTY_LIB.

Returns dict {crit_ps, n_cells, area_um2} or raises.
"""
import os
import subprocess
import sys

FLOW = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, FLOW)
import gen_sbox
import sta_liberty

YOSYS = os.environ.get(
    "YOSYS_BIN", "/home/hatch/tools/micromamba-root/envs/sky130/bin/yosys")
YOSYS_WANT = "0.69"

YS_TMPL = """# liberty-delay screen: pinned flow shape, abc -liberty mapping
read_verilog {src}
hierarchy -check -top {top}
proc; opt; flatten; opt
memory_map; opt
techmap; opt
abc -liberty {lib}
opt; clean
stat
write_verilog -noattr {mapped}
"""


def liberty_screen(tab, name, workdir, liberty, timeout=900):
    """tab: 16-list. -> {crit_ps, n_cells, area_um2} (raises on failure)."""
    workdir = os.path.abspath(workdir)
    os.makedirs(workdir, exist_ok=True)
    vp = os.path.join(workdir, f"{name}.v")
    with open(vp, "w") as f:
        f.write(gen_sbox.write_sbox64(tab, name))
    mapped = os.path.join(workdir, f"mapped_lib_{name}.v")
    ys = os.path.join(workdir, f"s_lib_{name}.ys")
    with open(ys, "w") as f:
        f.write(YS_TMPL.format(src=vp, top=name, lib=liberty, mapped=mapped))
    vv = subprocess.run([YOSYS, "-V"], capture_output=True, text=True,
                        timeout=60)
    if YOSYS_WANT not in (vv.stdout + vv.stderr):
        raise RuntimeError(f"wrong yosys at {YOSYS}")
    r = subprocess.run([YOSYS, "-s", ys], capture_output=True, text=True,
                       cwd=workdir, timeout=timeout)
    if r.returncode != 0 or not os.path.exists(mapped):
        raise RuntimeError(f"yosys liberty pass failed for {name}:\n"
                           f"{r.stdout[-2000:]}\n{r.stderr[-2000:]}")
    return sta_liberty.crit_ps(mapped, liberty)


def main():
    import json
    tab = [int(c, 16) for c in sys.argv[1]]
    name, workdir, liberty = sys.argv[2], sys.argv[3], sys.argv[4]
    print(json.dumps(liberty_screen(tab, name, workdir, liberty), indent=1))


if __name__ == "__main__":
    main()
