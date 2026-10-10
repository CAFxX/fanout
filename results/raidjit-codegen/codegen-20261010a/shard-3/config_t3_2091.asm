; config_t3_2091 target=rvv twin=baseline topo=rs(3,1)/enc op=encode U=1
; code_size=224 twins=[] uops_per_byte=0.875
     0: 63140600                 bnez a2, 8
     4: 6f00400d                 j 0xd8
     8: 970f0000                 auipc t6, 0
    12: 938f8f0d                 addi t6, t6, 0xd8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f000006                 j 0x88
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
   108: 83b30500                 ld t2, 0(a1)
   112: b3836300                 add t2, t2, t1
   116: 27800302                 vse8.v v0, (t2)
   120: 3303d300                 add t1, t1, a3
   124: 1307f7ff                 addi a4, a4, -1
   128: e31007fc                 bnez a4, 0x40
   132: 33066640                 sub a2, a2, t1
   136: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   140: 63940200                 bnez t0, 0x94
   144: 6f008004                 j 0xd8
   148: 83330500                 ld t2, 0(a0)
   152: b3836300                 add t2, t2, t1
   156: 07800302                 vle8.v v0, (t2)
   160: 83338500                 ld t2, 8(a0)
   164: b3836300                 add t2, t2, t1
   168: 87800302                 vle8.v v1, (t2)
   172: 5700102e                 vxor.vv v0, v1, v0
   176: 83330501                 ld t2, 0x10(a0)
   180: b3836300                 add t2, t2, t1
   184: 87800302                 vle8.v v1, (t2)
   188: 5700102e                 vxor.vv v0, v1, v0
   192: 83b30500                 ld t2, 0(a1)
   196: b3836300                 add t2, t2, t1
   200: 27800302                 vse8.v v0, (t2)
   204: 33035300                 add t1, t1, t0
   208: 33065640                 sub a2, a2, t0
   212: 6ff05ffb                 j 0x88
   216: 67800000                 ret 
   220: 0000                     unimp 
   222: 0000                     unimp 
