; config_t1_1956 target=rvv-zvbc32e twin=maskedtail topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=512 twins=["maskedtail"] uops_per_byte=4.4375
     0: 63140600                 bnez a2, 8
     4: 6f00c01e                 j 0x1f0
     8: 970f0000                 auipc t6, 0
    12: 938f8f1f                 addi t6, t6, 0x1f8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d7450e5e                 vmv.v.x v11, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00800d                 j 0x108
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330501                 ld t2, 0x10(a0)
    68: b3836300                 add t2, t2, t1
    72: 87800302                 vle8.v v1, (t2)
    76: 83338501                 ld t2, 0x18(a0)
    80: b3836300                 add t2, t2, t1
    84: 07810302                 vle8.v v2, (t2)
    88: 83330502                 ld t2, 0x20(a0)
    92: b3836300                 add t2, t2, t1
    96: 87810302                 vle8.v v3, (t2)
   100: 83330500                 ld t2, 0(a0)
   104: b3836300                 add t2, t2, t1
   108: 833e8500                 ld t4, 8(a0)
   112: b38e6e00                 add t4, t4, t1
   116: 87820302                 vle8.v v5, (t2)
   120: 07830e02                 vle8.v v6, (t4)
   124: 5782622e                 vxor.vv v4, v6, v5
   128: d780402e                 vxor.vv v1, v4, v1
   132: 138e0f00                 mv t3, t6
   136: 87040e02                 vle8.v v9, (t3)
   140: 138e0f01                 addi t3, t6, 0x10
   144: 07050e02                 vle8.v v10, (t3)
   148: d702b226                 vand.vv v5, v11, v4
   152: 573342a2                 vsrl.vi v6, v4, 4
   156: d7839232                 vrgather.vv v7, v9, v5
   160: 5704a332                 vrgather.vv v8, v10, v6
   164: d783832e                 vxor.vv v7, v8, v7
   168: 5701712e                 vxor.vv v2, v7, v2
   172: 138e0f02                 addi t3, t6, 0x20
   176: 87040e02                 vle8.v v9, (t3)
   180: 138e0f03                 addi t3, t6, 0x30
   184: 07050e02                 vle8.v v10, (t3)
   188: d702b226                 vand.vv v5, v11, v4
   192: 573342a2                 vsrl.vi v6, v4, 4
   196: d7839232                 vrgather.vv v7, v9, v5
   200: 5704a332                 vrgather.vv v8, v10, v6
   204: d783832e                 vxor.vv v7, v8, v7
   208: d781712e                 vxor.vv v3, v7, v3
   212: 83b30500                 ld t2, 0(a1)
   216: b3836300                 add t2, t2, t1
   220: a7800302                 vse8.v v1, (t2)
   224: 83b38500                 ld t2, 8(a1)
   228: b3836300                 add t2, t2, t1
   232: 27810302                 vse8.v v2, (t2)
   236: 83b30501                 ld t2, 0x10(a1)
   240: b3836300                 add t2, t2, t1
   244: a7810302                 vse8.v v3, (t2)
   248: 3303d300                 add t1, t1, a3
   252: 1307f7ff                 addi a4, a4, -1
   256: e31007f4                 bnez a4, 0x40
   260: 33066640                 sub a2, a2, t1
   264: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
   268: 93060600                 mv a3, a2
   272: 6344d000                 bgtz a3, 0x118
   276: 6f00400d                 j 0x1e8
   280: b3835640                 sub t2, a3, t0
   284: 13def343                 srai t3, t2, 0x3f
   288: b3f3c301                 and t2, t2, t3
   292: b3837200                 add t2, t0, t2
   296: 57a20852                 vid.v v4
   300: 57c0436a                 vmsltu.vx v0, v4, t2
   304: 83330501                 ld t2, 0x10(a0)
   308: b3836300                 add t2, t2, t1
   312: 87800300                 vle8.v v1, (t2), v0.t
   316: 83338501                 ld t2, 0x18(a0)
   320: b3836300                 add t2, t2, t1
   324: 07810300                 vle8.v v2, (t2), v0.t
   328: 83330502                 ld t2, 0x20(a0)
   332: b3836300                 add t2, t2, t1
   336: 87810300                 vle8.v v3, (t2), v0.t
   340: 83330500                 ld t2, 0(a0)
   344: b3836300                 add t2, t2, t1
   348: 833e8500                 ld t4, 8(a0)
   352: b38e6e00                 add t4, t4, t1
   356: 87820300                 vle8.v v5, (t2), v0.t
   360: 07830e00                 vle8.v v6, (t4), v0.t
   364: 5782622e                 vxor.vv v4, v6, v5
   368: d780402e                 vxor.vv v1, v4, v1
   372: 138e0f00                 mv t3, t6
   376: 87040e02                 vle8.v v9, (t3)
   380: 138e0f01                 addi t3, t6, 0x10
   384: 07050e02                 vle8.v v10, (t3)
   388: d702b226                 vand.vv v5, v11, v4
   392: 573342a2                 vsrl.vi v6, v4, 4
   396: d7839232                 vrgather.vv v7, v9, v5
   400: 5704a332                 vrgather.vv v8, v10, v6
   404: d783832e                 vxor.vv v7, v8, v7
   408: 5701712e                 vxor.vv v2, v7, v2
   412: 138e0f02                 addi t3, t6, 0x20
   416: 87040e02                 vle8.v v9, (t3)
   420: 138e0f03                 addi t3, t6, 0x30
   424: 07050e02                 vle8.v v10, (t3)
   428: d702b226                 vand.vv v5, v11, v4
   432: 573342a2                 vsrl.vi v6, v4, 4
   436: d7839232                 vrgather.vv v7, v9, v5
   440: 5704a332                 vrgather.vv v8, v10, v6
   444: d783832e                 vxor.vv v7, v8, v7
   448: d781712e                 vxor.vv v3, v7, v3
   452: 83b30500                 ld t2, 0(a1)
   456: b3836300                 add t2, t2, t1
   460: a7800300                 vse8.v v1, (t2), v0.t
   464: 83b38500                 ld t2, 8(a1)
   468: b3836300                 add t2, t2, t1
   472: 27810300                 vse8.v v2, (t2), v0.t
   476: 83b30501                 ld t2, 0x10(a1)
   480: b3836300                 add t2, t2, t1
   484: a7810300                 vse8.v v3, (t2), v0.t
   488: 33035300                 add t1, t1, t0
   492: b3865640                 sub a3, a3, t0
   496: 67800000                 ret 
   500: 0000                     unimp 
   502: 0000                     unimp 
   504: 0000                     unimp 
   506: 0000                     unimp 
   508: 0000                     unimp 
   510: 0000                     unimp 
