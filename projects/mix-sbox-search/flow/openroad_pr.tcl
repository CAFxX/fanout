# Tiny-macro Sky130 P&R for a single 4-bit S-box (combinational).
#
# Runs inside the PINNED PUBLIC image:
#   efabless/openlane:1.0.0-amd64@sha256:2a161827bc615796d60b5e3c7ce5b67ec463ea4ef4afb08ead7ee92fab972cff
# (Docker Hub, verified 2026-10-05; free to pull, costs no GHCR storage.)
# The project image is NEVER used for P&R.
#
# Env inputs:
#   DESIGN_V   verilog file (single 4-bit S-box module, from gen_sbox --lanes4)
#   TOP        top module name
#   PDK_ROOT   volare PDK root (contains sky130A/)
#   OUT_DIR    where to write routed.v, design.spef
#   DIE_UM     die size in um (default 40)
#
# Outputs (stdout, parsed by driver/run_pr_pick.py):
#   PR_MARKER_BEGIN_CHECKS / END_CHECKS : report_checks -path_delay max
#   PR_MARKER_BEGIN_POWER  / END_POWER  : report_power
#   PR_MARKER_DIE_AREA <x1 y1 x2 y2> <dbu_per_micron>
#   PR_MARKER_BEGIN_PORTS  / END_PORTS  : input port list
# Plus files: $OUT_DIR/routed.v, $OUT_DIR/design.spef
#
# DOCUMENTED SCOPE: estimates only (post-route delay, routed area, power).
# No PDN mesh is built and DRC/LVS-cleanliness is NOT a goal — the P&R
# stage exists to measure wire effects on the critical path for the
# README calibration plan, not to tape out the macro.
# NOT YET RUN (2026-10-05): first P&R dispatch will validate this script
# end-to-end; treat as draft until then.

set design_v $env(DESIGN_V)
set top      $env(TOP)
set pdk_root $env(PDK_ROOT)
set out_dir  $env(OUT_DIR)
set die_um   [expr {[info exists env(DIE_UM)] ? $env(DIE_UM) : 40}]

set std     "$pdk_root/sky130A/libs.ref/sky130_fd_sc_hd"
set tlef    "$std/lef/sky130_fd_sc_hd.tlef"
set lef     "$std/lef/sky130_fd_sc_hd.lef"
set lib     "$std/lib/sky130_fd_sc_hd__tt_100C_1v80.lib"
foreach f [list $tlef $lef $lib $design_v] {
    if {![file exists $f]} {
        puts "PR_FATAL missing file: $f"
        exit 2
    }
}
file mkdir $out_dir

read_lef $tlef
read_lef $lef
read_liberty $lib
read_verilog $design_v
link_design $top

# floorplan: tiny die, core margin for pins
set dd [expr {$die_um / 2.0}]
initialize_floorplan -die_area "0 0 $die_um $die_um" \
    -core_area "3 3 [expr {$die_um - 3}] [expr {$die_um - 3}]"

add_global_connection -net VDD -inst_pattern .* -pin_pattern {^VDD$} -power
add_global_connection -net VSS -inst_pattern .* -pin_pattern {^VSS$} -ground
global_connect

tapcell -distance 14 \
    -tapcell_master sky130_fd_sc_hd__tapvpwrvg_1 \
    -endcap_master sky130_fd_sc_hd__decap_3

global_placement
detailed_placement

# NOTE: no PDN mesh (estimates only; see header). Signal routing:
global_route
detailed_route

write_verilog "$out_dir/routed.v"
write_spef "$out_dir/design.spef"

# ---- reports (parsed by the driver) ----
puts "PR_MARKER_BEGIN_PORTS"
foreach port [get_ports *] {
    if {[get_property $port direction] eq "input"} {
        puts "PORT_INPUT $port"
    }
}
puts "PR_MARKER_END_PORTS"

puts "PR_MARKER_DIE_AREA [ord::get_die_area] [ord::get_dbu_per_micron]"

puts "PR_MARKER_BEGIN_CHECKS"
report_checks -path_delay max -format full_clock_expanded -digits 4
puts "PR_MARKER_END_CHECKS"

puts "PR_MARKER_BEGIN_POWER"
report_power
puts "PR_MARKER_END_POWER"

puts "PR_DONE"
exit 0
