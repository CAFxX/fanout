; config_t1_1694 target=rvv twin=notail topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=288 twins=["notail"] uops_per_byte=2.0625
     0: 63140600                 bnez a2, 8
     4: 6f008010                 j 0x10c
     8: 970f0000                 auipc t6, 0
    12: 938f8f11                 addi t6, t6, 0x118
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d7440e5e                 vmv.v.x v9, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00c00d                 j 0x10c
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330502                 ld t2, 0x20(a0)
    68: b3836300                 add t2, t2, t1
    72: 07800302                 vle8.v v0, (t2)
    76: 83338502                 ld t2, 0x28(a0)
    80: b3836300                 add t2, t2, t1
    84: 87800302                 vle8.v v1, (t2)
    88: 83330500                 ld t2, 0(a0)
    92: b3836300                 add t2, t2, t1
    96: 833e0501                 ld t4, 0x10(a0)
   100: b38e6e00                 add t4, t4, t1
   104: 87810302                 vle8.v v3, (t2)
   108: 07820e02                 vle8.v v4, (t4)
   112: 5781412e                 vxor.vv v2, v4, v3
   116: 5700202e                 vxor.vv v0, v2, v0
   120: 138e0f00                 mv t3, t6
   124: 87030e02                 vle8.v v7, (t3)
   128: 138e0f01                 addi t3, t6, 0x10
   132: 07040e02                 vle8.v v8, (t3)
   136: d7019126                 vand.vv v3, v9, v2
   140: 573222a2                 vsrl.vi v4, v2, 4
   144: d7827132                 vrgather.vv v5, v7, v3
   148: 57038232                 vrgather.vv v6, v8, v4
   152: d782622e                 vxor.vv v5, v6, v5
   156: d780502e                 vxor.vv v1, v5, v1
   160: 83338500                 ld t2, 8(a0)
   164: b3836300                 add t2, t2, t1
   168: 833e8501                 ld t4, 0x18(a0)
   172: b38e6e00                 add t4, t4, t1
   176: 87810302                 vle8.v v3, (t2)
   180: 07820e02                 vle8.v v4, (t4)
   184: 5781412e                 vxor.vv v2, v4, v3
   188: 5700202e                 vxor.vv v0, v2, v0
   192: 138e0f02                 addi t3, t6, 0x20
   196: 87030e02                 vle8.v v7, (t3)
   200: 138e0f03                 addi t3, t6, 0x30
   204: 07040e02                 vle8.v v8, (t3)
   208: d7019126                 vand.vv v3, v9, v2
   212: 573222a2                 vsrl.vi v4, v2, 4
   216: d7827132                 vrgather.vv v5, v7, v3
   220: 57038232                 vrgather.vv v6, v8, v4
   224: d782622e                 vxor.vv v5, v6, v5
   228: d780502e                 vxor.vv v1, v5, v1
   232: 83b30500                 ld t2, 0(a1)
   236: b3836300                 add t2, t2, t1
   240: 27800302                 vse8.v v0, (t2)
   244: 83b38500                 ld t2, 8(a1)
   248: b3836300                 add t2, t2, t1
   252: a7800302                 vse8.v v1, (t2)
   256: 3303d300                 add t1, t1, a3
   260: 1307f7ff                 addi a4, a4, -1
   264: e31c07f2                 bnez a4, 0x40
   268: 67800000                 ret 
   272: 0000                     unimp 
   274: 0000                     unimp 
   276: 0000                     unimp 
   278: 0000                     unimp 
   280: 0000                     unimp 
   282: 0000                     unimp 
   284: 0000                     unimp 
   286: 0000                     unimp 
