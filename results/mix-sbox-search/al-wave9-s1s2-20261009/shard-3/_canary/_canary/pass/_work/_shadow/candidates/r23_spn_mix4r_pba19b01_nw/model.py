"""Self-contained numpy model for the r23 SPN canary (no VM paths).

Construction (from fast_battery/candidates/r23_spn_mix4r_pba19b01_nw/impl.c):
  SBOX S1 = [0,1,2,4,3,8,15,12,9,5,11,6,7,14,13,10] (per nibble)
  P(i) = (19*i+1) mod 64
  LIN(x) = x ^ rotl64(x,3) ^ rotl64(x,11)
  round i (0-based): x = LIN(P(SBOX(x ^ rk_i)))
  rk_i = rotl64(key, 13*(i+1)) ^ RC[i]
  RC = [0x9E3779B97F4A7C15, 0x3C6EF372FE94F82A,
        0xDAA66D2C7DDF743F, 0x78DDE6E5FD29F054]
ROUNDS is set per canary bundle (4 = pass, 1 = S1 kill).
Bit-exactness vs the VM model was proven on 100k random vectors
(see battery/canaries/PROOF.md).
"""
import numpy as np

SBOX = np.array([0, 1, 2, 4, 3, 8, 15, 12, 9, 5, 11, 6, 7, 14, 13, 10],
                dtype=np.uint64)
PBOX = np.array([(19 * i + 1) % 64 for i in range(64)], dtype=np.uint64)
RC = np.array([0x9E3779B97F4A7C15, 0x3C6EF372FE94F82A,
               0xDAA66D2C7DDF743F, 0x78DDE6E5FD29F054], dtype=np.uint64)
M64 = np.uint64(0xFFFFFFFFFFFFFFFF)

ROUNDS = 4  # overridden per bundle


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
    return (x ^ _rotl(x, 3) ^ _rotl(x, 11)) & M64


def _round(x, key, i):
    rk = _rotl(key, 13 * (i + 1)) ^ RC[i]
    return _lin(_pbox64(_sbox64(x ^ rk)))


def mix_np(V, K):
    x = np.asarray(V, dtype=np.uint64)
    k = np.asarray(K, dtype=np.uint64)
    for i in range(ROUNDS):
        x = _round(x, k, i)
    return x & M64


def mix(v, k):
    return int(mix_np(np.uint64(v), np.uint64(k)))
