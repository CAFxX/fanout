import numpy as np
def mix(v, k): return int(v) & 0xFFFFFFFFFFFFFFFF
def mix_np(V, K): return np.asarray(V, dtype=np.uint64)
