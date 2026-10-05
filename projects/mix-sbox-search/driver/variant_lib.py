"""Shared variant-enumeration library. EXACT port of the scheme in
exploration/round25/hidden_files/variant_sweep/tick.py — the vi -> (ip, op, c)
mapping, the 16-sampled permutation sets, and the base tables are identical,
so a variant with index vi here is the same S-box the on-VM sweep measures.

  S'(x) = P_out( S( P_in(x) ) ) XOR c
  P_in : 16 evenly-sampled of the 24 input bit-permutations
  P_out: 16 evenly-sampled of the 24 output bit-permutations
  c    : all 16 4-bit output XOR constants
=> 4096 variants per base box (vi in [0, 4096)).

NOTE on visit order: tick.py walks variants in a STRIDE=7919 permutation
order (an early-stop optimization for the local cron). The GH sharder uses
contiguous vi slices; the vi -> table mapping is what must match, and it
does. Every variant is still visited exactly once.
"""
import itertools
import os

# Base boxes (exact tables from tick.py CANDS)
CANDS = {
    "midori_sb0": [0xc, 0xa, 0xd, 0x3, 0xe, 0xb, 0xf, 0x7,
                   0x8, 0x9, 0x1, 0x5, 0x0, 0x2, 0x4, 0x6],
    "ulbc_s1": [0x8, 0x0, 0x1, 0x5, 0xc, 0x7, 0x4, 0x6,
                0x2, 0xa, 0x3, 0xd, 0xe, 0xf, 0xb, 0x9],
    "thf_blink_s0": [0x1, 0x0, 0x9, 0x3, 0x8, 0x5, 0xe, 0x7,
                     0x4, 0x2, 0xc, 0xb, 0xa, 0xf, 0x6, 0xd],
}
ORDER = ["midori_sb0", "ulbc_s1", "thf_blink_s0"]

# Golden unit-delay values from ROUND25_TABLE.md (pinned flow, 16-lane form).
# The per-shard canary synthesizes vi=0 (identity variant == base table) and
# must reproduce these EXACTLY, else the shard is quarantined.
GOLDEN_CRIT = {
    "midori_sb0": 3.0,
    "ulbc_s1": 3.4,
    "thf_blink_s0": 3.4,
}

# S1 baseline (3.7u golden) for the validation self-test / pool canary.
S1 = [0, 1, 2, 4, 3, 8, 15, 12, 9, 5, 11, 6, 7, 14, 13, 10]

ALL_PERMS = sorted(itertools.permutations(range(4)))          # 24
SAMPLE_IDX = [round(i * 23 / 15) for i in range(16)]            # 16 evenly spread
assert len(set(SAMPLE_IDX)) == 16 and SAMPLE_IDX[0] == 0 and SAMPLE_IDX[-1] == 23
PERMS = [ALL_PERMS[i] for i in SAMPLE_IDX]                     # p: newbit[p[i]] = oldbit[i]

N_VARIANTS = 4096


def apply_perm(x, p):
    y = 0
    for i in range(4):
        if (x >> i) & 1:
            y |= 1 << p[i]
    return y


def variant_spec(vi):
    return vi // 256, (vi // 16) % 16, vi % 16


def build_table(base, ip, op, c):
    pin, pout = PERMS[ip], PERMS[op]
    return [apply_perm(base[apply_perm(x, pin)], pout) ^ c for x in range(16)]


def variant_table(box, vi):
    base = CANDS[box]
    ip, op, c = variant_spec(vi)
    return build_table(base, ip, op, c)


def table_hex(tab):
    return "".join(f"{v:x}" for v in tab)


def shard_range(shard_idx, num_shards):
    """Contiguous vi slice [lo, hi) for this shard."""
    lo = (shard_idx * N_VARIANTS) // num_shards
    hi = ((shard_idx + 1) * N_VARIANTS) // num_shards
    return lo, hi


def load_pool(path):
    """Parse a pool file. Two formats:
    'NAME 0c 0a ...' rows, or bare '0x0,0x1,...,' CSV rows (one box/line)."""
    boxes = []
    with open(path) as f:
        for line in f:
            line = line.strip()
            if not line or line.startswith("#"):
                continue
            if "," in line:
                vals = [int(t, 16) for t in line.replace(",", " ").split()]
                if len(vals) == 16:
                    boxes.append((f"pool_{len(boxes)}", vals))
            else:
                toks = line.split()
                vals = [int(t, 16) for t in toks[1:17]]
                if len(vals) == 16:
                    boxes.append((toks[0], vals))
    return boxes


def pool_tables(pools_dir):
    out = []
    for fn in sorted(os.listdir(pools_dir)):
        if fn.endswith(".txt"):
            for name, tab in load_pool(os.path.join(pools_dir, fn)):
                out.append((f"{fn}:{name}", tab))
    return out
