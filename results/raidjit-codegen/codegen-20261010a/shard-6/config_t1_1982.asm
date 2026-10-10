; config_t1_1982 target=rvv-zvbc32e twin=maskedtail topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=528 twins=["maskedtail"] uops_per_byte=4.6875
     0: 63140600                 bnez a2, 8
     4: 6f00c01f                 j 0x200
     8: 970f0000                 auipc t6, 0
    12: 938f8f20                 addi t6, t6, 0x208
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 57450e5e                 vmv.v.x v10, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00000e                 j 0x110
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330502                 ld t2, 0x20(a0)
    68: b3836300                 add t2, t2, t1
    72: 87800302                 vle8.v v1, (t2)
    76: 83338502                 ld t2, 0x28(a0)
    80: b3836300                 add t2, t2, t1
    84: 07810302                 vle8.v v2, (t2)
    88: 83330500                 ld t2, 0(a0)
    92: b3836300                 add t2, t2, t1
    96: 833e0501                 ld t4, 0x10(a0)
   100: b38e6e00                 add t4, t4, t1
   104: 07820302                 vle8.v v4, (t2)
   108: 87820e02                 vle8.v v5, (t4)
   112: d701522e                 vxor.vv v3, v5, v4
   116: d780302e                 vxor.vv v1, v3, v1
   120: 138e0f00                 mv t3, t6
   124: 07040e02                 vle8.v v8, (t3)
   128: 138e0f01                 addi t3, t6, 0x10
   132: 87040e02                 vle8.v v9, (t3)
   136: 5782a126                 vand.vv v4, v10, v3
   140: d73232a2                 vsrl.vi v5, v3, 4
   144: 57038232                 vrgather.vv v6, v8, v4
   148: d7839232                 vrgather.vv v7, v9, v5
   152: 5703732e                 vxor.vv v6, v7, v6
   156: 5701612e                 vxor.vv v2, v6, v2
   160: 83338500                 ld t2, 8(a0)
   164: b3836300                 add t2, t2, t1
   168: 833e8501                 ld t4, 0x18(a0)
   172: b38e6e00                 add t4, t4, t1
   176: 07820302                 vle8.v v4, (t2)
   180: 87820e02                 vle8.v v5, (t4)
   184: d701522e                 vxor.vv v3, v5, v4
   188: d780302e                 vxor.vv v1, v3, v1
   192: 138e0f02                 addi t3, t6, 0x20
   196: 07040e02                 vle8.v v8, (t3)
   200: 138e0f03                 addi t3, t6, 0x30
   204: 87040e02                 vle8.v v9, (t3)
   208: 5782a126                 vand.vv v4, v10, v3
   212: d73232a2                 vsrl.vi v5, v3, 4
   216: 57038232                 vrgather.vv v6, v8, v4
   220: d7839232                 vrgather.vv v7, v9, v5
   224: 5703732e                 vxor.vv v6, v7, v6
   228: 5701612e                 vxor.vv v2, v6, v2
   232: 83b30500                 ld t2, 0(a1)
   236: b3836300                 add t2, t2, t1
   240: a7800302                 vse8.v v1, (t2)
   244: 83b38500                 ld t2, 8(a1)
   248: b3836300                 add t2, t2, t1
   252: 27810302                 vse8.v v2, (t2)
   256: 3303d300                 add t1, t1, a3
   260: 1307f7ff                 addi a4, a4, -1
   264: e31c07f2                 bnez a4, 0x40
   268: 33066640                 sub a2, a2, t1
   272: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
   276: 93060600                 mv a3, a2
   280: 6344d000                 bgtz a3, 0x120
   284: 6f00c00d                 j 0x1f8
   288: b3835640                 sub t2, a3, t0
   292: 13def343                 srai t3, t2, 0x3f
   296: b3f3c301                 and t2, t2, t3
   300: b3837200                 add t2, t0, t2
   304: d7a10852                 vid.v v3
   308: 57c0336a                 vmsltu.vx v0, v3, t2
   312: 83330502                 ld t2, 0x20(a0)
   316: b3836300                 add t2, t2, t1
   320: 87800300                 vle8.v v1, (t2), v0.t
   324: 83338502                 ld t2, 0x28(a0)
   328: b3836300                 add t2, t2, t1
   332: 07810300                 vle8.v v2, (t2), v0.t
   336: 83330500                 ld t2, 0(a0)
   340: b3836300                 add t2, t2, t1
   344: 833e0501                 ld t4, 0x10(a0)
   348: b38e6e00                 add t4, t4, t1
   352: 07820300                 vle8.v v4, (t2), v0.t
   356: 87820e00                 vle8.v v5, (t4), v0.t
   360: d701522e                 vxor.vv v3, v5, v4
   364: d780302e                 vxor.vv v1, v3, v1
   368: 138e0f00                 mv t3, t6
   372: 07040e02                 vle8.v v8, (t3)
   376: 138e0f01                 addi t3, t6, 0x10
   380: 87040e02                 vle8.v v9, (t3)
   384: 5782a126                 vand.vv v4, v10, v3
   388: d73232a2                 vsrl.vi v5, v3, 4
   392: 57038232                 vrgather.vv v6, v8, v4
   396: d7839232                 vrgather.vv v7, v9, v5
   400: 5703732e                 vxor.vv v6, v7, v6
   404: 5701612e                 vxor.vv v2, v6, v2
   408: 83338500                 ld t2, 8(a0)
   412: b3836300                 add t2, t2, t1
   416: 833e8501                 ld t4, 0x18(a0)
   420: b38e6e00                 add t4, t4, t1
   424: 07820300                 vle8.v v4, (t2), v0.t
   428: 87820e00                 vle8.v v5, (t4), v0.t
   432: d701522e                 vxor.vv v3, v5, v4
   436: d780302e                 vxor.vv v1, v3, v1
   440: 138e0f02                 addi t3, t6, 0x20
   444: 07040e02                 vle8.v v8, (t3)
   448: 138e0f03                 addi t3, t6, 0x30
   452: 87040e02                 vle8.v v9, (t3)
   456: 5782a126                 vand.vv v4, v10, v3
   460: d73232a2                 vsrl.vi v5, v3, 4
   464: 57038232                 vrgather.vv v6, v8, v4
   468: d7839232                 vrgather.vv v7, v9, v5
   472: 5703732e                 vxor.vv v6, v7, v6
   476: 5701612e                 vxor.vv v2, v6, v2
   480: 83b30500                 ld t2, 0(a1)
   484: b3836300                 add t2, t2, t1
   488: a7800300                 vse8.v v1, (t2), v0.t
   492: 83b38500                 ld t2, 8(a1)
   496: b3836300                 add t2, t2, t1
   500: 27810300                 vse8.v v2, (t2), v0.t
   504: 33035300                 add t1, t1, t0
   508: b3865640                 sub a3, a3, t0
   512: 67800000                 ret 
   516: 0000                     unimp 
   518: 0000                     unimp 
   520: 0000                     unimp 
   522: 0000                     unimp 
   524: 0000                     unimp 
   526: 0000                     unimp 
