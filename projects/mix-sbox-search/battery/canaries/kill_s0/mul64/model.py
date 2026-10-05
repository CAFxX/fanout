import numpy as np
def mix(v, k): return (int(v) * int(k)) & 0xFFFFFFFFFFFFFFFF
def mix_np(V, K):
    return (np.asarray(V, dtype=np.uint64) * np.asarray(K, dtype=np.uint64)) & np.uint64(0xFFFFFFFFFFFFFFFF)
