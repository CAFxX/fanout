#!/usr/bin/env python3
"""8-bit S-box validation suite (track-2 broad screen v2).

Mirrors flow/validate_sbox.py's dimensions but for 256-entry tables,
using numpy for speed (DDT/LAT would be ~minutes in pure Python).
The 4-bit module is PINNED (bit-exactness) and untouched; this is the
separate 8-bit path.

Dimensions: bijective, du (differential uniformity), nl (nonlinearity),
anf_degrees (per-coordinate), full_dependency, fixed_points,
cycle_structure, diff/lin branch numbers, boomerang uniformity.

Hard-gate policy (flags, not verdicts):
  bijective, du <= 8, nl >= 96.
Degree-7-all and full-dependency are recorded flags, opt-in at promotion
(the AES S-box itself is [7]*8 with full dependency, so the flags are
 informative, not exclusionary).
"""
import sys

import numpy as np

N = 256
NB = 8


def _as_tab(tab):
    t = np.asarray(list(tab), dtype=np.uint8)
    assert t.shape == (N,), f"expected 256 entries, got {t.shape}"
    return t


def ddt_du(t):
    """Differential uniformity from the DDT (vectorized over x)."""
    xs = np.arange(N, dtype=np.uint8)
    du = 0
    tx = t[xs]
    for a in range(1, N):
        diffs = tx ^ t[xs ^ np.uint8(a)]
        c = np.bincount(diffs, minlength=N)
        m = int(c.max())
        if m > du:
            du = m
            if du == N:  # can't get worse
                break
    return du


def _fwht(a):
    a = a.astype(np.int32)
    h = 1
    n = len(a)
    while h < n:
        for i in range(0, n, h * 2):
            x = a[i:i + h].copy()
            y = a[i + h:i + 2 * h]
            a[i:i + h] = x + y
            a[i + h:i + 2 * h] = x - y
        h *= 2
    return a


def lat_nl(t):
    """Nonlinearity from the LAT via 256 FWHTs (one per output mask)."""
    worst = 0
    for b in range(1, N):
        fb = t & np.uint8(b)
        # parity of (b & S(x)) for all x
        par = np.unpackbits(fb).reshape(N, 8).sum(axis=1) & 1
        g = np.where(par == 0, 1, -1)
        w = _fwht(g)
        m = int(np.abs(w).max())
        if m > worst:
            worst = m
    return 128 - worst // 2, worst


def anf_degrees(t):
    degs = []
    for bit in range(NB):
        f = ((t >> np.uint8(bit)) & np.uint8(1)).astype(np.uint8)
        a = f.copy()
        for i in range(NB):
            step = 1 << i
            idx = np.arange(N)
            sel = (idx & step) != 0
            a[sel] ^= a[idx[sel] ^ step]
        deg = 0
        nz = np.nonzero(a)[0]
        for mask in nz:
            d = bin(int(mask)).count("1")
            if d > deg:
                deg = d
        degs.append(deg)
    return degs


def full_dependency(t):
    xs = np.arange(N, dtype=np.uint8)
    for ob in range(NB):
        col_o = ((t >> np.uint8(ob)) & np.uint8(1)).astype(np.uint8)
        for ib in range(NB):
            # does flipping input bit ib ever change output bit ob?
            if not np.any(col_o[xs] != col_o[xs ^ np.uint8(1 << ib)]):
                return False
    return True


def diff_branch_number(t):
    # min over a!=0 of wt(a)+wt(b) over DDT-supported (a,b)
    xs = np.arange(N, dtype=np.uint8)
    tx = t[xs]
    dbn = N + 1
    for a in range(1, N):
        diffs = tx ^ t[xs ^ np.uint8(a)]
        seen = np.unique(diffs)
        wa = bin(a).count("1")
        for b in seen:
            w = wa + bin(int(b)).count("1")
            if w < dbn:
                dbn = w
    return dbn


def boomerang_uniformity(t):
    inv = np.zeros(N, dtype=np.uint8)
    inv[t] = np.arange(N, dtype=np.uint8)
    xs = np.arange(N, dtype=np.uint8)
    bu = 0
    for a in range(1, N):
        ia = np.uint8(a)
        tax = t[xs ^ ia]
        tx = t[xs]
        for b in range(1, N):
            ib = np.uint8(b)
            c = np.count_nonzero(
                (inv[tx ^ ib] ^ inv[tax ^ ib]) == ia)
            if c > bu:
                bu = int(c)
    return bu


def cycle_structure(t):
    seen = np.zeros(N, dtype=bool)
    cycles = []
    for s in range(N):
        if seen[s]:
            continue
        cyc = []
        x = s
        while not seen[x]:
            seen[x] = True
            cyc.append(x)
            x = int(t[x])
        cycles.append(len(cyc))
    cycles.sort(reverse=True)
    return cycles


def validate8(tab):
    t = _as_tab(tab)
    bij = len(np.unique(t)) == N
    du = ddt_du(t)
    nl, wmax = lat_nl(t)
    degs = anf_degrees(t)
    dep = full_dependency(t)
    fp = int(np.count_nonzero(t == np.arange(N, dtype=np.uint8)))
    if bij:
        dbn = diff_branch_number(t)
        lbn = None  # linear branch number not computed for 8-bit (cost)
        bu = boomerang_uniformity(t)
        cycles = cycle_structure(t)
    else:
        dbn = lbn = bu = None
        cycles = None
    return {
        "bijective": bool(bij),
        "du": du,
        "nl": nl,
        "lat_wmax": wmax,
        "anf_degrees": degs,
        "full_dependency": bool(dep),
        "fixed_points": fp,
        "cycle_structure": cycles,
        "diff_branch_number": dbn,
        "lin_branch_number": lbn,
        "boomerang_uniformity": bu,
        "gates": {
            "bijective": bool(bij),
            "du_le8": du <= 8,
            "nl_ge96": nl >= 96,
            "deg7_all": all(d == 7 for d in degs),
            "full_dependency": bool(dep),
        },
    }


# Unconditional hard gates for the 8-bit screen (synthesis gate).
HARD_GATES8 = ("bijective", "du_le8", "nl_ge96")


def hard_gates_pass8(v):
    g = v["gates"]
    return all(g[k] for k in HARD_GATES8)


def self_test():
    # AES S-box generated independently by pools/gen_sbox8_pool.py logic
    sys.path.insert(0, os.path.join(os.path.dirname(os.path.abspath(__file__)),
                                    "..", "pools"))
    import gen_sbox8_pool as gp
    aes = gp.aes_sbox()
    v = validate8(aes)
    assert v["bijective"], "AES not bijective"
    assert v["du"] == 4, f"AES DU={v['du']}"
    assert v["nl"] == 112, f"AES NL={v['nl']}"
    assert v["anf_degrees"] == [7] * 8, f"AES degrees={v['anf_degrees']}"
    assert v["full_dependency"], "AES not fully dependent"
    assert hard_gates_pass8(v), "AES hard gates should pass"
    # identity must FAIL the gates (non-vacuity)
    ident = list(range(256))
    vi = validate8(ident)
    assert vi["bijective"], "identity should be bijective"
    assert not hard_gates_pass8(vi), "identity must fail hard gates"
    assert vi["du"] == 256 and vi["nl"] == 0
    print(f"validate_sbox8 self-test OK: AES du=4 nl=112 "
          f"dbn={v['diff_branch_number']} bu={v['boomerang_uniformity']}; "
          f"identity du=256 nl=0 gates_fail OK")


if __name__ == "__main__":
    import os
    if "--self-test" in sys.argv:
        self_test()
    else:
        print("validate_sbox8: use --self-test")
