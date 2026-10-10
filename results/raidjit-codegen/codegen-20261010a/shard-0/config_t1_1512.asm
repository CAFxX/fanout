; config_t1_1512 target=rvv twin=notail topo=rs(6,3)/enc op=encode U=1
; code_size=624 twins=["notail"] uops_per_byte=7.375
     0: 63140600                 bnez a2, 8
     4: 6f004026                 j 0x268
     8: 970f0000                 auipc t6, 0
    12: 938f8f26                 addi t6, t6, 0x268
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 57450e5e                 vmv.v.x v10, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f008023                 j 0x268
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 87810302                 vle8.v v3, (t2)
    76: 5780015e                 vmv.v.v v0, v3
    80: 138e0f00                 mv t3, t6
    84: 07040e02                 vle8.v v8, (t3)
    88: 138e0f01                 addi t3, t6, 0x10
    92: 87040e02                 vle8.v v9, (t3)
    96: 5782a126                 vand.vv v4, v10, v3
   100: d73232a2                 vsrl.vi v5, v3, 4
   104: 57038232                 vrgather.vv v6, v8, v4
   108: d7839232                 vrgather.vv v7, v9, v5
   112: 5703732e                 vxor.vv v6, v7, v6
   116: d700035e                 vmv.v.v v1, v6
   120: 138e0f02                 addi t3, t6, 0x20
   124: 07040e02                 vle8.v v8, (t3)
   128: 138e0f03                 addi t3, t6, 0x30
   132: 87040e02                 vle8.v v9, (t3)
   136: 5782a126                 vand.vv v4, v10, v3
   140: d73232a2                 vsrl.vi v5, v3, 4
   144: 57038232                 vrgather.vv v6, v8, v4
   148: d7839232                 vrgather.vv v7, v9, v5
   152: 5703732e                 vxor.vv v6, v7, v6
   156: 5701035e                 vmv.v.v v2, v6
   160: 83338500                 ld t2, 8(a0)
   164: b3836300                 add t2, t2, t1
   168: 87810302                 vle8.v v3, (t2)
   172: 5700302e                 vxor.vv v0, v3, v0
   176: 138e0f02                 addi t3, t6, 0x20
   180: 07040e02                 vle8.v v8, (t3)
   184: 138e0f03                 addi t3, t6, 0x30
   188: 87040e02                 vle8.v v9, (t3)
   192: 5782a126                 vand.vv v4, v10, v3
   196: d73232a2                 vsrl.vi v5, v3, 4
   200: 57038232                 vrgather.vv v6, v8, v4
   204: d7839232                 vrgather.vv v7, v9, v5
   208: 5703732e                 vxor.vv v6, v7, v6
   212: d780602e                 vxor.vv v1, v6, v1
   216: 138e0f04                 addi t3, t6, 0x40
   220: 07040e02                 vle8.v v8, (t3)
   224: 138e0f05                 addi t3, t6, 0x50
   228: 87040e02                 vle8.v v9, (t3)
   232: 5782a126                 vand.vv v4, v10, v3
   236: d73232a2                 vsrl.vi v5, v3, 4
   240: 57038232                 vrgather.vv v6, v8, v4
   244: d7839232                 vrgather.vv v7, v9, v5
   248: 5703732e                 vxor.vv v6, v7, v6
   252: 5701612e                 vxor.vv v2, v6, v2
   256: 83330501                 ld t2, 0x10(a0)
   260: b3836300                 add t2, t2, t1
   264: 87810302                 vle8.v v3, (t2)
   268: 5700302e                 vxor.vv v0, v3, v0
   272: 138e0f04                 addi t3, t6, 0x40
   276: 07040e02                 vle8.v v8, (t3)
   280: 138e0f05                 addi t3, t6, 0x50
   284: 87040e02                 vle8.v v9, (t3)
   288: 5782a126                 vand.vv v4, v10, v3
   292: d73232a2                 vsrl.vi v5, v3, 4
   296: 57038232                 vrgather.vv v6, v8, v4
   300: d7839232                 vrgather.vv v7, v9, v5
   304: 5703732e                 vxor.vv v6, v7, v6
   308: d780602e                 vxor.vv v1, v6, v1
   312: 138e0f06                 addi t3, t6, 0x60
   316: 07040e02                 vle8.v v8, (t3)
   320: 138e0f07                 addi t3, t6, 0x70
   324: 87040e02                 vle8.v v9, (t3)
   328: 5782a126                 vand.vv v4, v10, v3
   332: d73232a2                 vsrl.vi v5, v3, 4
   336: 57038232                 vrgather.vv v6, v8, v4
   340: d7839232                 vrgather.vv v7, v9, v5
   344: 5703732e                 vxor.vv v6, v7, v6
   348: 5701612e                 vxor.vv v2, v6, v2
   352: 83338501                 ld t2, 0x18(a0)
   356: b3836300                 add t2, t2, t1
   360: 87810302                 vle8.v v3, (t2)
   364: 5700302e                 vxor.vv v0, v3, v0
   368: 138e0f08                 addi t3, t6, 0x80
   372: 07040e02                 vle8.v v8, (t3)
   376: 138e0f09                 addi t3, t6, 0x90
   380: 87040e02                 vle8.v v9, (t3)
   384: 5782a126                 vand.vv v4, v10, v3
   388: d73232a2                 vsrl.vi v5, v3, 4
   392: 57038232                 vrgather.vv v6, v8, v4
   396: d7839232                 vrgather.vv v7, v9, v5
   400: 5703732e                 vxor.vv v6, v7, v6
   404: d780602e                 vxor.vv v1, v6, v1
   408: 138e0f00                 mv t3, t6
   412: 07040e02                 vle8.v v8, (t3)
   416: 138e0f01                 addi t3, t6, 0x10
   420: 87040e02                 vle8.v v9, (t3)
   424: 5782a126                 vand.vv v4, v10, v3
   428: d73232a2                 vsrl.vi v5, v3, 4
   432: 57038232                 vrgather.vv v6, v8, v4
   436: d7839232                 vrgather.vv v7, v9, v5
   440: 5703732e                 vxor.vv v6, v7, v6
   444: 5701612e                 vxor.vv v2, v6, v2
   448: 83330502                 ld t2, 0x20(a0)
   452: b3836300                 add t2, t2, t1
   456: 87810302                 vle8.v v3, (t2)
   460: 5700302e                 vxor.vv v0, v3, v0
   464: 138e0f06                 addi t3, t6, 0x60
   468: 07040e02                 vle8.v v8, (t3)
   472: 138e0f07                 addi t3, t6, 0x70
   476: 87040e02                 vle8.v v9, (t3)
   480: 5782a126                 vand.vv v4, v10, v3
   484: d73232a2                 vsrl.vi v5, v3, 4
   488: 57038232                 vrgather.vv v6, v8, v4
   492: d7839232                 vrgather.vv v7, v9, v5
   496: 5703732e                 vxor.vv v6, v7, v6
   500: d780602e                 vxor.vv v1, v6, v1
   504: 138e0f08                 addi t3, t6, 0x80
   508: 07040e02                 vle8.v v8, (t3)
   512: 138e0f09                 addi t3, t6, 0x90
   516: 87040e02                 vle8.v v9, (t3)
   520: 5782a126                 vand.vv v4, v10, v3
   524: d73232a2                 vsrl.vi v5, v3, 4
   528: 57038232                 vrgather.vv v6, v8, v4
   532: d7839232                 vrgather.vv v7, v9, v5
   536: 5703732e                 vxor.vv v6, v7, v6
   540: 5701612e                 vxor.vv v2, v6, v2
   544: 83338502                 ld t2, 0x28(a0)
   548: b3836300                 add t2, t2, t1
   552: 87810302                 vle8.v v3, (t2)
   556: 5700302e                 vxor.vv v0, v3, v0
   560: d780302e                 vxor.vv v1, v3, v1
   564: 5701312e                 vxor.vv v2, v3, v2
   568: 83b30500                 ld t2, 0(a1)
   572: b3836300                 add t2, t2, t1
   576: 27800302                 vse8.v v0, (t2)
   580: 83b38500                 ld t2, 8(a1)
   584: b3836300                 add t2, t2, t1
   588: a7800302                 vse8.v v1, (t2)
   592: 83b30501                 ld t2, 0x10(a1)
   596: b3836300                 add t2, t2, t1
   600: 27810302                 vse8.v v2, (t2)
   604: 3303d300                 add t1, t1, a3
   608: 1307f7ff                 addi a4, a4, -1
   612: e31e07dc                 bnez a4, 0x40
   616: 67800000                 ret 
   620: 0000                     unimp 
   622: 0000                     unimp 
