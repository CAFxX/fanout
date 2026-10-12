#!/usr/bin/env python3
"""Generate the hetstage2 GHA pool (runs LOCALLY, deterministic).

Pool (300 draws):
  - 150 dense-GF: pc stratified over {8,10,12,14,16,18,20,24,28,32,36,40},
    templates H1/H3 (H3 weighted: the pilot's keeps were H3-only).
  - 150 R5: [CIRC2, MUXD x4].

Draw operators are imported from the VM pilot
(t8_hetstage2_screen._draw_gfmul/_draw_circ2/TEMPLATES, het8._draw_muxd) —
the GHA side evaluates pre-drawn stage specs, so no RNG runs on GHA.

Also writes canary.jsonl: the 9 pilot keeps (must reproduce inner-PASS +
fisher-PASS with z within tolerance) + the hetstage try53 replay
(must inner-FLAG).
"""
import json
import os
import sys

import numpy as np

HERE = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, "/home/hatch/workspace/mix/exploration/sbox_screen/hetstage2")
sys.path.insert(0, "/home/hatch/workspace/mix/exploration/sbox_screen/hetstage")
sys.path.insert(0, "/home/hatch/workspace/mix/exploration/sbox_screen/het8")
sys.path.insert(0, "/home/hatch/workspace/mix/exploration")
import t8_hetstage2_screen as S2

POOL_SEED = 0xB804
N_GF = 150
N_R5 = 150
GF_PCS = S2.POOL_PCS


def ser(st):
    if st[0] == "GFMUL":
        return ["GFMUL", hex(st[1])]
    return list(st)


def main():
    rng = np.random.default_rng(POOL_SEED)
    pool = []
    # Dense-GF: H3 weighted 2:1 over H1 (pilot: keeps were H3-only).
    for i in range(N_GF):
        pc = GF_PCS[i % len(GF_PCS)]
        tname = "H3" if (i % 3) != 2 else "H1"
        h = S2._draw_gfmul(rng, pc)
        stages = S2.TEMPLATES[tname](rng, h)
        pool.append({"pool_idx": len(pool), "kind": "GF",
                     "template": tname, "pc": pc,
                     "stages": [ser(s) for s in stages],
                     "seed": POOL_SEED})
    # R5: [CIRC2, MUXD x4].
    for _ in range(N_R5):
        h = S2._draw_circ2(rng)
        stages = S2.TEMPLATES["R5"](rng, h)
        pool.append({"pool_idx": len(pool), "kind": "R5",
                     "template": "R5", "stages": [ser(s) for s in stages],
                     "seed": POOL_SEED})
    with open(os.path.join(HERE, "hetstage2_pool.jsonl"), "w") as f:
        for rec in pool:
            f.write(json.dumps(rec) + "\n")
    # Canary: 9 pilot keeps + try53 kill replay.
    keeps = [json.loads(l) for l in open(
        "/home/hatch/workspace/mix/exploration/sbox_screen/hetstage2/"
        "pilot_keeps.jsonl")]
    rep = S2._replay_hetstage_tries()
    can = []
    for rec in keeps:
        can.append({"name": "keep_z%.2f_%s_%s" % (rec["z"], rec["type"],
                                                 rec["template"]),
                    "expect": "KEEP", "expect_z": rec["z"],
                    "stages": rec["stages"]})
    can.append({"name": "kill_try53_replay", "expect": "INNER-FLAG",
                "stages": [ser(s) for s in rep[53]]})
    with open(os.path.join(HERE, "canary.jsonl"), "w") as f:
        for rec in can:
            f.write(json.dumps(rec) + "\n")
    print("pool: %d draws (%d GF + %d R5), canary: %d entries" %
          (len(pool), N_GF, N_R5, len(can)), flush=True)


if __name__ == "__main__":
    main()
