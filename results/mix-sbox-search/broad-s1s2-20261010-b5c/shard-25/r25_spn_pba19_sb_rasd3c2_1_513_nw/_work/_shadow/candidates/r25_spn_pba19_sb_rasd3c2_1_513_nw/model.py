"""Self-contained numpy model for r25_spn_pba19_sb_rasd3c2_1_513_nw (no VM paths).

Construction:
  SBOX = [0, 1, 2, 3, 4, 6, 8, 12, 5, 7, 11, 15, 13, 14, 9, 10] (per nibble)
  P(i) = (19*i+1) mod 64   (out[P[i]] = in[i])
  LIN(x) = x ^ rotl64(x,5) ^ rotl64(x,13)
  round i (0-based): x = LIN(P(SBOX(x ^ rk_i)))
  rk_i = rotl64(key, 13*(i+1)) ^ RC[i]
"""
import numpy as np

SBOX = np.array([0, 1, 2, 3, 4, 6, 8, 12, 5, 7, 11, 15, 13, 14, 9, 10], dtype=np.uint64)
PBOX = np.array([1, 20, 39, 58, 13, 32, 51, 6, 25, 44, 63, 18, 37, 56, 11, 30, 49, 4, 23, 42, 61, 16, 35, 54, 9, 28, 47, 2, 21, 40, 59, 14, 33, 52, 7, 26, 45, 0, 19, 38, 57, 12, 31, 50, 5, 24, 43, 62, 17, 36, 55, 10, 29, 48, 3, 22, 41, 60, 15, 34, 53, 8, 27, 46], dtype=np.uint64)
RC = np.array([11400714819323198485, 4354685564936845354, 15755400384260043839, 8709371129873690708], dtype=np.uint64)
M64 = np.uint64(0xFFFFFFFFFFFFFFFF)
ROUNDS = 4
ROTA = 5
ROTB = 13

def _rotl(x, r):
    r %= 64
    if r == 0:
        return x
    return ((x << np.uint64(r)) | (x >> np.uint64(64 - r))) & M64

def _sbox64(x):
    y = np.zeros_like(x)
    for i in range(16):
        y |= SBOX[((x >> np.uint64(4 * i)) & np.uint64(0xF))] \
            << np.uint64(4 * i)
    return y & M64

def _pbox64(x):
    y = np.zeros_like(x)
    for i in range(64):
        y |= ((x >> np.uint64(i)) & np.uint64(1)) << PBOX[i]
    return y & M64

def _lin(x):
    return (x ^ _rotl(x, ROTA) ^ _rotl(x, ROTB)) & M64

def _rk(key, i):
    return _rotl(key, 13 * (i + 1)) ^ RC[i]

def mix(v, k):
    v = np.uint64(v)
    k = np.uint64(k)
    x = v
    for i in range(ROUNDS):
        x = _lin(_pbox64(_sbox64(x ^ _rk(k, i))))
    return int(x & M64)

def mix_np(v, k):
    v = np.asarray(v, dtype=np.uint64)
    k = np.uint64(k)
    x = v.copy()
    for i in range(ROUNDS):
        x = _lin(_pbox64(_sbox64(x ^ _rk(k, i))))
    return x & M64

mix64 = mix_np
