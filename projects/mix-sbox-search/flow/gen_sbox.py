#!/usr/bin/env python3
"""S-box table -> synthesizable Verilog.

Extracted from exploration/round25/gen_r25.py (Quine-McCluskey DNF emission).
The 64-bit 16-lane form is BYTE-IDENTICAL in structure to what the on-VM
variant sweep (tick.py synth_variant) feeds run_synth.py, so unit-delay
numbers reproduce exactly (MIDORI_Sb0 -> 3.0u etc.).

Two emitters:
  write_sbox64(tab, name) -> module {name}(input [63:0] x, output [63:0] y):
      16 parallel lanes of the 4-bit S-box, direct boolean assigns inside
      one always @(*) block. This is the form the golden numbers use.
  write_sbox4(tab, name)  -> module {name}(input [3:0] x, output [3:0] y):
      single nibble. Used for the OpenROAD P&R macro (critical path of one
      lane is what matters; all lanes are identical).

CLI: python3 gen_sbox.py <hex16> <name> [--lanes4] > out.v
  <hex16>: 16 hex chars, e.g. cad30ebf7891050246
"""
import sys


# ---------------- Quine-McCluskey (4 vars), from gen_r25 ----------------
def _qm_primes(minterms):
    cubes = [(0, m) for m in sorted(minterms)]
    primes = set()
    while True:
        used = [False] * len(cubes)
        nxt = []
        seen = set()
        for i in range(len(cubes)):
            for j in range(i + 1, len(cubes)):
                m1, b1 = cubes[i]
                m2, b2 = cubes[j]
                if m1 != m2:
                    continue
                d = (b1 ^ b2) & ~m1 & 0xF
                if d and (d & (d - 1)) == 0:
                    nc = (m1 | d, (b1 & ~d) & 0xF)
                    if nc not in seen:
                        seen.add(nc)
                        nxt.append(nc)
                    used[i] = used[j] = True
        for i, c in enumerate(cubes):
            if not used[i]:
                primes.add(c)
        if not nxt:
            break
        cubes = nxt
    return primes


def _qm_cover(minterms, primes):
    uncovered = set(minterms)
    chosen = []
    primes = list(primes)
    while uncovered:
        def covers(p):
            m, b = p
            return [v for v in uncovered if (v & ~m) == (b & ~m)]
        best = max(primes, key=lambda p: len(covers(p)))
        cov = covers(best)
        assert cov, "prime cover failed"
        chosen.append(best)
        uncovered -= set(cov)
    return chosen


_dnf_cache = {}


def dnf_for_sbox_bit(sbox, bit):
    key = (tuple(sbox), bit)
    if key not in _dnf_cache:
        mt = [v for v in range(16) if (sbox[v] >> bit) & 1]
        if not mt:
            cubes = []
        elif len(mt) == 16:
            cubes = [(0xF, 0)]
        else:
            cubes = _qm_cover(mt, _qm_primes(mt))
        _dnf_cache[key] = cubes
    return _dnf_cache[key]


def _v_sbox16_stmts(sbox, inp, outp, lane_tables=None):
    """Emit S-box DNF assigns. sbox: 16-list (all lanes) or lane_tables:
    list of 16 16-lists (lane n uses lane_tables[n])."""
    lines = []
    for n in range(16):
        tab = lane_tables[n] if lane_tables is not None else sbox
        for bit in range(4):
            cubes = dnf_for_sbox_bit(tab, bit)
            if not cubes:
                lines.append(f"        {outp}[{4*n+bit}] = 1'b0;")
                continue
            terms = []
            for m, b in cubes:
                lits = []
                for k in range(4):
                    if not ((m >> k) & 1):
                        lits.append(f"{'' if (b >> k) & 1 else '~'}{inp}[{4*n+k}]")
                terms.append("(" + " & ".join(lits) + ")" if lits else "1'b1")
            lines.append(f"        {outp}[{4*n+bit}] = " + " | ".join(terms) + ";")
    return lines


def _v_sbox4_stmts(sbox, inp, outp):
    lines = []
    for bit in range(4):
        cubes = dnf_for_sbox_bit(sbox, bit)
        if not cubes:
            lines.append(f"        {outp}[{bit}] = 1'b0;")
            continue
        terms = []
        for m, b in cubes:
            lits = []
            for k in range(4):
                if not ((m >> k) & 1):
                    lits.append(f"{'' if (b >> k) & 1 else '~'}{inp}[{k}]")
            terms.append("(" + " & ".join(lits) + ")" if lits else "1'b1")
        lines.append(f"        {outp}[{bit}] = " + " | ".join(terms) + ";")
    return lines


def write_sbox64(tab, name):
    """16-lane 64-bit form. MUST match tick.py synth_variant exactly."""
    lines = [f"module {name}(input [63:0] x, output [63:0] y);",
             "reg [63:0] y;", "always @(*) begin"]
    lines += ["    " + l.strip() for l in _v_sbox16_stmts(tab, "x", "y")]
    lines += ["end", "endmodule"]
    return "\n".join(lines) + "\n"


def write_sbox4(tab, name):
    """Single-nibble form for the P&R macro."""
    lines = [f"module {name}(input [3:0] x, output [3:0] y);",
             "reg [3:0] y;", "always @(*) begin"]
    lines += ["    " + l.strip() for l in _v_sbox4_stmts(tab, "x", "y")]
    lines += ["end", "endmodule"]
    return "\n".join(lines) + "\n"


def main():
    hex16, name = sys.argv[1], sys.argv[2]
    tab = [int(c, 16) for c in hex16]
    assert len(tab) == 16 and len(set(tab)) == 16, "need 16 distinct hex nibbles"
    if "--lanes4" in sys.argv[3:]:
        sys.stdout.write(write_sbox4(tab, name))
    else:
        sys.stdout.write(write_sbox64(tab, name))


if __name__ == "__main__":
    main()
