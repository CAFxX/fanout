#!/usr/bin/env python3
"""Full S-box validation suite (dimension B: validate on all relevant axes).

Pure Python, fast at 4 bits. Per candidate table (16-list) returns a JSON-able
dict. NO inverse implementation metrics anywhere (dropped by methodology);
S^-1 is used only as a mathematical tool for boomerang uniformity, which is
a property of the forward function, not an inverse-implementation cost.

Dimensions measured:
  bijective, du (differential uniformity from DDT),
  nl (nonlinearity from LAT), anf_degrees (per-coordinate algebraic degree),
  full_dependency (every output bit depends on all 4 input bits),
  fixed_points, cycle_structure,
  diff_branch_number (from DDT), lin_branch_number (from LAT),
  boomerang_uniformity (from DDT + S^-1 as a math tool).

Hard-gate policy used by drivers (flags, not verdicts here):
  bijective, du == 4, nl == 4, all anf_degrees == 3, full_dependency.
"""
import sys

S1 = [0, 1, 2, 4, 3, 8, 15, 12, 9, 5, 11, 6, 7, 14, 13, 10]
MIDORI_SB0 = [0xc, 0xa, 0xd, 0x3, 0xe, 0xb, 0xf, 0x7,
              0x8, 0x9, 0x1, 0x5, 0x0, 0x2, 0x4, 0x6]


def _wt(x):
    return bin(x).count("1")


def ddt(tab):
    t = [[0] * 16 for _ in range(16)]
    for a in range(16):
        for x in range(16):
            t[a][tab[x] ^ tab[x ^ a]] += 1
    return t


def lat(tab):
    # Walsh-Hadamard: LAT[a][b] = sum_x (-1)^{<a,x>+<b,S(x)>}
    w = [[0] * 16 for _ in range(16)]
    for a in range(16):
        for b in range(16):
            s = 0
            for x in range(16):
                if (_wt(a & x) + _wt(b & tab[x])) % 2 == 0:
                    s += 1
                else:
                    s -= 1
            w[a][b] = s
    return w


def anf_degrees(tab):
    """Algebraic degree of each of the 4 coordinate functions via Moebius."""
    degs = []
    for bit in range(4):
        f = [(tab[x] >> bit) & 1 for x in range(16)]
        # Moebius transform over GF(2)
        a = f[:]
        for i in range(4):
            for mask in range(16):
                if mask & (1 << i):
                    a[mask] ^= a[mask ^ (1 << i)]
        deg = max([_wt(m) for m in range(16) if a[m]] or [0])
        degs.append(deg)
    return degs


def full_dependency(tab):
    """True iff every output bit depends on every input bit."""
    for ob in range(4):
        for ib in range(4):
            if not any((((tab[x] >> ob) & 1) ^ ((tab[x ^ (1 << ib)] >> ob) & 1))
                       for x in range(16)):
                return False
    return True


def cycle_structure(tab):
    seen = [False] * 16
    lens = []
    for s in range(16):
        if not seen[s]:
            c, x = 0, s
            while not seen[x]:
                seen[x] = True
                x = tab[x]
                c += 1
            lens.append(c)
    return sorted(lens)


def branch_numbers(tab, dd=None, ll=None):
    dd = dd or ddt(tab)
    ll = ll or lat(tab)
    dbn = min(_wt(a) + _wt(b)
              for a in range(1, 16) for b in range(16) if dd[a][b] > 0)
    lbn = min(_wt(a) + _wt(b)
              for a in range(1, 16) for b in range(16) if ll[a][b] != 0)
    return dbn, lbn


def boomerang_uniformity(tab):
    inv = [0] * 16
    for x, y in enumerate(tab):
        inv[y] = x
    bu = 0
    for a in range(1, 16):
        for b in range(1, 16):
            c = sum(1 for x in range(16)
                    if (inv[tab[x] ^ b] ^ inv[tab[x ^ a] ^ b]) == a)
            bu = max(bu, c)
    return bu


def validate(tab):
    tab = list(tab)
    dd = ddt(tab)
    ll = lat(tab)
    du = max(dd[a][b] for a in range(1, 16) for b in range(16))
    wmax = max(abs(ll[a][b]) for a in range(1, 16) for b in range(1, 16))
    nl = 8 - wmax // 2
    degs = anf_degrees(tab)
    dep = full_dependency(tab)
    fp = sum(1 for x in range(16) if tab[x] == x)
    bij = len(set(tab)) == 16
    if bij:
        dbn, lbn = branch_numbers(tab, dd, ll)
        bu = boomerang_uniformity(tab)
        cycles = cycle_structure(tab)
    else:
        # degenerate (non-bijective) input: structural metrics undefined
        dbn = lbn = bu = None
        cycles = None
    return {
        "bijective": bij,
        "du": du,
        "nl": nl,
        "anf_degrees": degs,
        "full_dependency": dep,
        "fixed_points": fp,
        "cycle_structure": cycles,
        "diff_branch_number": dbn,
        "lin_branch_number": lbn,
        "boomerang_uniformity": bu,
        # hard-gate summary. UNCONDITIONAL hards (expert methodology):
        # bijective, DU=4, NL=4. Degree-3-all and full-dependency are
        # "where required" (the 3.0u champion MIDORI_Sb0 is [3,2,3,3]
        # without full dependency) — recorded as flags, opt-in at
        # promotion time, never silently exclusionary by default.
        "gates": {
            "bijective": bij,
            "du4": du == 4,
            "nl4": nl == 4,
            "deg3_all": all(d == 3 for d in degs),
            "full_dependency": dep,
        },
    }


# Unconditional hard gates (synthesis gate in drivers, default promotion
# gate). Degree/full-dependency are conditional per the expert methodology.
HARD_GATES = ("bijective", "du4", "nl4")


def hard_gates_pass(v):
    g = v["gates"]
    return all(g[k] for k in HARD_GATES)


def gates_pass(v):
    """Legacy alias: all recorded flags (strict). Prefer hard_gates_pass
    unless the caller explicitly wants the conditional gates too."""
    g = v["gates"]
    return all(g.values())


def self_test():
    v = validate(S1)
    assert v["bijective"], "S1 not bijective"
    assert v["du"] == 4, f"S1 DU={v['du']}"
    assert v["nl"] == 4, f"S1 NL={v['nl']}"
    assert v["anf_degrees"] == [3, 3, 3, 3], f"S1 degrees={v['anf_degrees']}"
    assert v["full_dependency"], "S1 not fully dependent"
    assert hard_gates_pass(v), "S1 hard gates should pass"
    assert gates_pass(v), "S1 all-flags should pass"
    m = validate(MIDORI_SB0)
    assert m["du"] == 4 and m["nl"] == 4, "MIDORI_Sb0 DU/NL"
    assert m["fixed_points"] == 4, f"MIDORI_Sb0 fp={m['fixed_points']}"
    # non-bijective input must fail the bijective gate, not crash
    bad = validate([0] * 16)
    assert not bad["gates"]["bijective"]
    print("validate_sbox self-test OK: "
          f"S1 DU=4 NL=4 deg={v['anf_degrees']} fp={v['fixed_points']} "
          f"dbn={v['diff_branch_number']} lbn={v['lin_branch_number']} "
          f"bu={v['boomerang_uniformity']}; "
          f"MIDORI_Sb0 DU=4 NL=4 fp={m['fixed_points']}")


if __name__ == "__main__":
    if "--self-test" in sys.argv:
        self_test()
    else:
        import json
        tab = [int(c, 16) for c in sys.argv[1]]
        print(json.dumps(validate(tab), indent=1))
