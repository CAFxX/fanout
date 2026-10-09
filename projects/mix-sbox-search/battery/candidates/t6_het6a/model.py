"""Self-contained bit-exact Python model for t6_het6a (track 6, batch 6 heterogeneous).

Generated 2026-10-09T09:47:30.750001+00:00 from ~/workspace/mix/exploration/new_constructions/t6_models_b6.py
via inspect.getsource — verbatim copy, no reimplementation (bit-exact by construction).
"""
import numpy as np

M64 = np.uint64(0xFFFFFFFFFFFFFFFF)
RC64 = [np.uint64(0x9E3779B97F4A7C15), np.uint64(0x3C6EF372FE94F82A),
        np.uint64(0xDAA66D2C7DDF743F), np.uint64(0x78DDE6E5FD29F054),
        np.uint64(0x2545F4914F6CDD1D), np.uint64(0x9E3779B97F4A7C15 ^ 0xDEADBEEF),
        np.uint64(0x3C6EF372FE94F82A ^ 0x12345678),
        np.uint64(0xDAA66D2C7DDF743F ^ 0xABCDEF01)]

RC64 = [np.uint64(0x9E3779B97F4A7C15), np.uint64(0x3C6EF372FE94F82A),
        np.uint64(0xDAA66D2C7DDF743F), np.uint64(0x78DDE6E5FD29F054),
        np.uint64(0x2545F4914F6CDD1D), np.uint64(0x9E3779B97F4A7C15 ^ 0xDEADBEEF),
        np.uint64(0x3C6EF372FE94F82A ^ 0x12345678),
        np.uint64(0xDAA66D2C7DDF743F ^ 0xABCDEF01)]

SCHED6_HET = [
    (11, 23, 37, 5, 17, 29),
    (7, 19, 31, 43, 13, 25),
    (3, 15, 27, 39, 51, 9),
    (21, 33, 45, 1, 35, 47),
    (5, 29, 53, 17, 41, 11),
    (13, 37, 61, 23, 47, 19),
]

OP_NROTS = {"maj": 4, "andor": 5, "chi": 3, "majxor": 5, "xor3and": 6}

def _rotl64(x, r):
    r %= 64
    if r == 0:
        return x
    return ((x << np.uint64(r)) | (x >> np.uint64(64 - r))) & M64

def _maj(a, b, c):
    return ((a & b) | (a & c) | (b & c)) & M64

def _chi(a, b, c):
    # Keccak chi: a ^= (b & ~c)
    return (a ^ (b & (~c & M64))) & M64

def _apply_op(op, r):
    """Apply op to rotated values r (list). Returns the nonlinear term."""
    if op == "maj":
        # r[0..3]: maj(r0,r1,r2) ^ r3
        maj = _maj(r[0], r[1], r[2])
        return (maj ^ r[3]) & M64
    elif op == "andor":
        # r[0..4]: ((r0&r1)|(r2&r3)) ^ r4
        return (((r[0] & r[1]) | (r[2] & r[3])) ^ r[4]) & M64
    elif op == "chi":
        # r[0..2]: chi(r0,r1,r2)
        return _chi(r[0], r[1], r[2])
    elif op == "majxor":
        # r[0..4]: maj(r0,r1,r2) ^ (r3&r4)
        maj = _maj(r[0], r[1], r[2])
        return (maj ^ (r[3] & r[4])) & M64
    elif op == "xor3and":
        # r[0..5]: r0^r1^r2 ^ (r3&r4&r5)
        return (r[0] ^ r[1] ^ r[2] ^ (r[3] & r[4] & r[5])) & M64
    else:
        raise ValueError(f"unknown op {op}")

def _hetero_cascade(v, k, seq):
    x = np.uint64(v) ^ np.uint64(k)
    for i, op in enumerate(seq):
        nrots = OP_NROTS[op]
        rots = SCHED6_HET[i][:nrots]
        r = [_rotl64(x, t) for t in rots]
        nl = _apply_op(op, r)
        rk = _rotl64(np.uint64(k), 13 * (i + 1)) ^ RC64[i % 8]
        x = (x ^ nl ^ rk) & M64
    # post-whiten
    x = (x ^ (_rotl64(np.uint64(k), 13) >> np.uint64(32))) & M64
    return x

def _mk_hetero(seq, fname):
    # Named wrapper (never a bare lambda): avoids the '<lambda>' __name__
    # SyntaxError at S1 load time (2026-10-08 lesson).
    def f(v, k):
        return _hetero_cascade(v, k, seq)
    f.__name__ = fname
    return f

def _vec(fn):
    def f(V, K):
        V = np.asarray(V, dtype=np.uint64)
        K = np.asarray(K, dtype=np.uint64)
        return np.array([fn(int(v), int(k)) for v, k in zip(V.flat, K.flat)],
                        dtype=np.uint64).reshape(V.shape)
    return f


_DESC = "heterogeneous cascade A: maj/andor/chi/majxor/xor3and/maj"
SEQ = SEQ_A = ["maj", "andor", "chi", "majxor", "xor3and", "maj"]
SEQ_B = ["chi", "maj", "andor", "xor3and", "majxor", "chi"]
SEQ_C = ["andor", "xor3and", "maj", "chi", "majxor", "andor"]


_mix_scalar = _mk_hetero(SEQ, "hetero_het6a")
_mix_np = _vec(_mix_scalar)

def mix(v, k):
    if np.ndim(v) == 0 and np.ndim(k) == 0:
        return _mix_scalar(int(v), int(k))
    return _mix_np(v, k)

def mix_np(V, K):
    return _mix_np(V, K)

meta = {"desc": _DESC}
