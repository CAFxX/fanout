; config_t1_1954 target=rvv-zvbc32e twin=notail topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=272 twins=["notail"] uops_per_byte=1.9375
     0: 63140600                 bnez a2, 8
     4: 6f000010                 j 0x104
     8: 970f0000                 auipc t6, 0
    12: 938f8f10                 addi t6, t6, 0x108
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 57450e5e                 vmv.v.x v10, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00400d                 j 0x104
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330501                 ld t2, 0x10(a0)
    68: b3836300                 add t2, t2, t1
    72: 07800302                 vle8.v v0, (t2)
    76: 83338501                 ld t2, 0x18(a0)
    80: b3836300                 add t2, t2, t1
    84: 87800302                 vle8.v v1, (t2)
    88: 83330502                 ld t2, 0x20(a0)
    92: b3836300                 add t2, t2, t1
    96: 07810302                 vle8.v v2, (t2)
   100: 83330500                 ld t2, 0(a0)
   104: b3836300                 add t2, t2, t1
   108: 833e8500                 ld t4, 8(a0)
   112: b38e6e00                 add t4, t4, t1
   116: 07820302                 vle8.v v4, (t2)
   120: 87820e02                 vle8.v v5, (t4)
   124: d701522e                 vxor.vv v3, v5, v4
   128: 5700302e                 vxor.vv v0, v3, v0
   132: 138e0f00                 mv t3, t6
   136: 07040e02                 vle8.v v8, (t3)
   140: 138e0f01                 addi t3, t6, 0x10
   144: 87040e02                 vle8.v v9, (t3)
   148: 5782a126                 vand.vv v4, v10, v3
   152: d73232a2                 vsrl.vi v5, v3, 4
   156: 57038232                 vrgather.vv v6, v8, v4
   160: d7839232                 vrgather.vv v7, v9, v5
   164: 5703732e                 vxor.vv v6, v7, v6
   168: d780602e                 vxor.vv v1, v6, v1
   172: 138e0f02                 addi t3, t6, 0x20
   176: 07040e02                 vle8.v v8, (t3)
   180: 138e0f03                 addi t3, t6, 0x30
   184: 87040e02                 vle8.v v9, (t3)
   188: 5782a126                 vand.vv v4, v10, v3
   192: d73232a2                 vsrl.vi v5, v3, 4
   196: 57038232                 vrgather.vv v6, v8, v4
   200: d7839232                 vrgather.vv v7, v9, v5
   204: 5703732e                 vxor.vv v6, v7, v6
   208: 5701612e                 vxor.vv v2, v6, v2
   212: 83b30500                 ld t2, 0(a1)
   216: b3836300                 add t2, t2, t1
   220: 27800302                 vse8.v v0, (t2)
   224: 83b38500                 ld t2, 8(a1)
   228: b3836300                 add t2, t2, t1
   232: a7800302                 vse8.v v1, (t2)
   236: 83b30501                 ld t2, 0x10(a1)
   240: b3836300                 add t2, t2, t1
   244: 27810302                 vse8.v v2, (t2)
   248: 3303d300                 add t1, t1, a3
   252: 1307f7ff                 addi a4, a4, -1
   256: e31007f4                 bnez a4, 0x40
   260: 67800000                 ret 
   264: 0000                     unimp 
   266: 0000                     unimp 
   268: 0000                     unimp 
   270: 0000                     unimp 
