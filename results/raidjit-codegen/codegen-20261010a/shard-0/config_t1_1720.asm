; config_t1_1720 target=rvv-zvbc32e twin=notail topo=raid5(4)/enc op=encode U=1
; code_size=160 twins=["notail"] uops_per_byte=0.5625
     0: 63140600                 bnez a2, 8
     4: 6f000009                 j 0x94
     8: 970f0000                 auipc t6, 0
    12: 938f8f09                 addi t6, t6, 0x98
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f00c006                 j 0x94
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07800302                 vle8.v v0, (t2)
    76: 83338500                 ld t2, 8(a0)
    80: b3836300                 add t2, t2, t1
    84: 87800302                 vle8.v v1, (t2)
    88: 5700102e                 vxor.vv v0, v1, v0
    92: 83330501                 ld t2, 0x10(a0)
    96: b3836300                 add t2, t2, t1
   100: 87800302                 vle8.v v1, (t2)
   104: 5700102e                 vxor.vv v0, v1, v0
   108: 83338501                 ld t2, 0x18(a0)
   112: b3836300                 add t2, t2, t1
   116: 87800302                 vle8.v v1, (t2)
   120: 5700102e                 vxor.vv v0, v1, v0
   124: 83b30500                 ld t2, 0(a1)
   128: b3836300                 add t2, t2, t1
   132: 27800302                 vse8.v v0, (t2)
   136: 3303d300                 add t1, t1, a3
   140: 1307f7ff                 addi a4, a4, -1
   144: e31807fa                 bnez a4, 0x40
   148: 67800000                 ret 
   152: 0000                     unimp 
   154: 0000                     unimp 
   156: 0000                     unimp 
   158: 0000                     unimp 
