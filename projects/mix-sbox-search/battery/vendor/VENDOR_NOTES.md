# Vendor notes — bit-exact ports of fast_battery logic

Every file under `vendor/` is a copy of the VM original with ONLY the
path/tool patches listed below. No algorithmic line was touched. Verify
with: `diff <(git show …)` — or against the VM originals at
`~/workspace/mix/exploration/fast_battery/src/` and
`~/workspace/mix/exploration/{screen_cal,screen_lib,screen_lindiff}.py`.

## Verbatim copies (zero changes)

- `s2_diffprof.c`, `s3_gen.c`, `s4_smallcrush.c` — pure C, no paths inside
  (candidate `impl.c` is passed at compile time).
- `abc_nodretime.script`, `cells_sim.v` — data files.
- `screen_lib.py`, `lin_diff.py` — no hardcoded paths.

## Patched copies (env overrides only; original literal kept as default)

`s0_delay.py` (9 patches):
- `FB` → `$BATTERY_HOME` (default: original fast_battery path)
- `EXP` → `$MIX_EXPLORATION` (only used for `analyze2.py`)
- `MIX` → `$MIX_ROOT`
- `SYNTH` → `$FLOW_DIR` (mini.genlib, analyze.py, mul64.v)
- `YOSYS` → `$YOSYS_BIN`
- abc script path → `$ABC_NODRETIME_SCRIPT`
- cells_sim.v path → `$CELLS_SIM_V`
- `iverilog`/`vvp` → `$IVERILOG_BIN`/`$VVP_BIN`

`s1_screen.py` (2 patches):
- `FB` → `$BATTERY_HOME`; `sys.path` inserts collapsed to a single
  `$VENDOR_DIR` (default `FB/src`; original inserted EXP then FB/src, so
  FB/src already took precedence — behavior identical).

`screen_cal.py`, `screen_lindiff.py` (1 patch each):
- `sys.path.insert(0, "/home/hatch/workspace/mix/exploration")` →
  `$VENDOR_DIR` (+ added the missing `import os` the new `os.environ`
  call needs; upstream never used `os`).

`xcheck.py` (1 patch):
- `FB` → `$BATTERY_HOME` (+ added the missing `import os` the new
  `os.environ` call needs; upstream never used `os`).

## Flow note (read before touching S0)

The battery S0 is fast_battery's `s0_delay.py`: `abc -genlib mini.genlib
-script abc_nodretime.script` (dretime REMOVED — proven 2026-10-01 via SAT
miter to miscompile `mix_h_xorspn3`; the no-dretime mapping is SAT-proven
equivalent). This is NOT `exploration/run_synth.py`'s plain
`abc -genlib mini.genlib` (which keeps dretime). The PARETO S0 numbers
(27.8u etc.) were measured with the no-dretime flow; the GHA S0 reproduces
that flow exactly, including the 1000-vector iverilog mapped-netlist
equivalence check (`$IVERILOG_BIN`).

## S4 verdict parsing (validated 2026-10-05)

TestU01 1.2.3 prints the plural
"The following tests gave p-values outside [0.001, 0.9990]:" on failure.
`s4.py` greps for the plural form AND requires the mutually-exclusive
"All tests were passed" line for a pass verdict. Both patterns were
validated against real passing output
(`fast_battery/results/r23_spn_mix4r_pba19b01_nw/s4.out`) and real failing
output (identity-mix canary) — see AGENTS.md 2026-10-05 lesson.
