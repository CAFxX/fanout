#!/usr/bin/env python3
"""hetstage2 evaluation core — VERBATIM port of the VM pilot's eval path.

Copied (not reimplemented) from:
  mix/exploration/sbox_screen/hetstage2/t8_hetstage2_screen.py
    (_parity, gf_masks, apply_stage, eval_inner, eval_sched, deser)
  mix/exploration/sbox_screen/het8/t8_het8_screen.py
    (rotl, mux, MUXD stage formula)
  mix/exploration/sbox_screen/hetstage/stage_cost.py
    (gf_mat — the transpose-fixed version)

Vendored instruments: lin_diff.screen_lin_diff / diff_cell_logps /
holm_reject, screen_lib.KEYS / N.
"""
import os
import sys

import numpy as np

sys.path.insert(0, os.path.join(os.path.dirname(os.path.abspath(__file__)),
                                "..", "vendored"))
from lin_diff import diff_cell_logps, holm_reject, screen_lin_diff  # noqa: E402
from screen_lib import KEYS, N  # noqa: E402

M64 = np.uint64(0xFFFFFFFFFFFFFFFF)
KEYS10 = [int(k) for k in KEYS]
INNER_KEYS = [0xFFFFFFFFFFFFFFFF, 0xAAAAAAAAAAAAAAAA]


# --- het8/t8_het8_screen.py verbatim ---
def rotl(x, r):
    r %= 64
    return ((x << np.uint64(r)) | (x >> np.uint64(64 - r))) & M64 if r else x


def mux(s, a, b):
    return ((s & a) | ((~s) & b)) & M64


# --- hetstage/stage_cost.py gf_mat verbatim (transpose-fixed) ---
def gf_mat(C):
    """Return list of 64 int masks: outbit j = parity(x & mask[j])."""
    P = (1 << 64) | (1 << 4) | (1 << 3) | (1 << 1) | 1  # degree-64 poly bits
    colvec = []
    for i in range(64):
        prod = 0
        cc = C
        b = 0
        while cc:
            if cc & 1:
                prod ^= 1 << (i + b)
            cc >>= 1
            b += 1
        # reduce bits 127..64
        for k in range(127, 63, -1):
            if (prod >> k) & 1:
                prod ^= (P << (k - 64))
        colvec.append(prod & ((1 << 64) - 1))
    masks = []
    for j in range(64):
        m = 0
        for i in range(64):
            if (colvec[i] >> j) & 1:
                m |= 1 << i
        masks.append(m)
    return masks


# --- hetstage2/t8_hetstage2_screen.py verbatim ---
_GF_MATS = {}


def gf_masks(C):
    if C not in _GF_MATS:
        _GF_MATS[C] = tuple(gf_mat(C))
    return _GF_MATS[C]


def _parity(v):
    v = v ^ (v >> np.uint64(32))
    v = v ^ (v >> np.uint64(16))
    v = v ^ (v >> np.uint64(8))
    v = v ^ (v >> np.uint64(4))
    v = v ^ (v >> np.uint64(2))
    v = v ^ (v >> np.uint64(1))
    return v & np.uint64(1)


def apply_stage(x, st):
    tag = st[0]
    if tag == "CIRC2":
        _, a, b = st
        return (x ^ rotl(x, a) ^ rotl(x, b)) & M64
    if tag == "GFMUL":
        _, C = st
        y = np.zeros(x.shape, dtype=np.uint64)
        for j, m in enumerate(gf_masks(C)):
            y |= _parity(x & np.uint64(m)) << np.uint64(j)
        return y
    if tag == "MUXD":
        _, s, a, b, d = st
        assert len({s, a, b}) == 3, "degenerate mux draw"
        return (x ^ mux(rotl(x, s), rotl(x, a), rotl(x, b))
                ^ rotl(x, d)) & M64
    raise AssertionError("unknown stage %r" % (st,))


def eval_inner(stages):
    """PRIMARY triage: screen_lin_diff on N=2^18 sequential outputs,
    INNER_KEYS. Verbatim instrument (hetstage/batch-62)."""
    seq = {}
    V = np.arange(N, dtype=np.uint64)
    for key in INNER_KEYS:
        K = np.full(N, np.uint64(key), dtype=np.uint64)
        x = (V ^ K) & M64
        for st in stages:
            x = apply_stage(x, st)
        seq[key] = ((x ^ rotl(K, 13)) & M64).astype(np.uint64)
    return screen_lin_diff(seq, verbose=False)


def eval_sched(stages):
    """Fisher triage, gate rule verbatim: KEEP iff n_reject == 0."""
    V = np.arange(N, dtype=np.uint64)
    worst = 0.0
    worst_key = None
    min_flp = 0.0
    cells = []
    for key in KEYS10:
        K = np.full(N, np.uint64(key), dtype=np.uint64)
        x = (V ^ K) & M64
        for st in stages:
            x = apply_stage(x, st)
        S = ((x ^ rotl(K, 13)) & M64).astype(np.uint64)
        _, wz, flp, _ = diff_cell_logps(S, key)
        cells.append((("DIFF", key), flp))
        if wz > worst:
            worst, worst_key = wz, key
        min_flp = min(min_flp, flp)
    n_reject = len(holm_reject(cells))
    return n_reject, worst, worst_key, min_flp


def deser(st):
    if st[0] == "GFMUL":
        return ("GFMUL", int(st[1], 16))
    return tuple(st[0:1]) + tuple(int(v) for v in st[1:])


def evaluate(stages):
    """Full triage. Returns dict with inner + fisher verdicts and KEEP."""
    inner = eval_inner(stages)
    ilin, idiff = inner["lin_verdict"], inner["diff_verdict"]
    nrej, wz, wkey, flp = eval_sched(stages)
    keep = (ilin == "PASS" and idiff == "PASS" and nrej == 0)
    return {"inner_lin": ilin, "inner_diff": idiff, "n_reject": nrej,
            "z": wz, "worst_key": wkey, "min_fisher": flp, "keep": keep}
