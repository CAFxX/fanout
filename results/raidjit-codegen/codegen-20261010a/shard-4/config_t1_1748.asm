; config_t1_1748 target=rvv-zvbc32e twin=maskedtail topo=raid6(8)/enc op=encode U=1
; code_size=1008 twins=["maskedtail"] uops_per_byte=11.8125
     0: 63140600                 bnez a2, 8
     4: 6f00403e                 j 0x3e8
     8: 970f0000                 auipc t6, 0
    12: 938f8f3e                 addi t6, t6, 0x3e8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 57450e5e                 vmv.v.x v10, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00401d                 j 0x204
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 87810302                 vle8.v v3, (t2)
    76: d780015e                 vmv.v.v v1, v3
    80: 138e0f00                 mv t3, t6
    84: 07040e02                 vle8.v v8, (t3)
    88: 138e0f01                 addi t3, t6, 0x10
    92: 87040e02                 vle8.v v9, (t3)
    96: 5782a126                 vand.vv v4, v10, v3
   100: d73232a2                 vsrl.vi v5, v3, 4
   104: 57038232                 vrgather.vv v6, v8, v4
   108: d7839232                 vrgather.vv v7, v9, v5
   112: 5703732e                 vxor.vv v6, v7, v6
   116: 5701035e                 vmv.v.v v2, v6
   120: 83338500                 ld t2, 8(a0)
   124: b3836300                 add t2, t2, t1
   128: 87810302                 vle8.v v3, (t2)
   132: d780302e                 vxor.vv v1, v3, v1
   136: 138e0f02                 addi t3, t6, 0x20
   140: 07040e02                 vle8.v v8, (t3)
   144: 138e0f03                 addi t3, t6, 0x30
   148: 87040e02                 vle8.v v9, (t3)
   152: 5782a126                 vand.vv v4, v10, v3
   156: d73232a2                 vsrl.vi v5, v3, 4
   160: 57038232                 vrgather.vv v6, v8, v4
   164: d7839232                 vrgather.vv v7, v9, v5
   168: 5703732e                 vxor.vv v6, v7, v6
   172: 5701612e                 vxor.vv v2, v6, v2
   176: 83330501                 ld t2, 0x10(a0)
   180: b3836300                 add t2, t2, t1
   184: 87810302                 vle8.v v3, (t2)
   188: d780302e                 vxor.vv v1, v3, v1
   192: 138e0f04                 addi t3, t6, 0x40
   196: 07040e02                 vle8.v v8, (t3)
   200: 138e0f05                 addi t3, t6, 0x50
   204: 87040e02                 vle8.v v9, (t3)
   208: 5782a126                 vand.vv v4, v10, v3
   212: d73232a2                 vsrl.vi v5, v3, 4
   216: 57038232                 vrgather.vv v6, v8, v4
   220: d7839232                 vrgather.vv v7, v9, v5
   224: 5703732e                 vxor.vv v6, v7, v6
   228: 5701612e                 vxor.vv v2, v6, v2
   232: 83338501                 ld t2, 0x18(a0)
   236: b3836300                 add t2, t2, t1
   240: 87810302                 vle8.v v3, (t2)
   244: d780302e                 vxor.vv v1, v3, v1
   248: 138e0f06                 addi t3, t6, 0x60
   252: 07040e02                 vle8.v v8, (t3)
   256: 138e0f07                 addi t3, t6, 0x70
   260: 87040e02                 vle8.v v9, (t3)
   264: 5782a126                 vand.vv v4, v10, v3
   268: d73232a2                 vsrl.vi v5, v3, 4
   272: 57038232                 vrgather.vv v6, v8, v4
   276: d7839232                 vrgather.vv v7, v9, v5
   280: 5703732e                 vxor.vv v6, v7, v6
   284: 5701612e                 vxor.vv v2, v6, v2
   288: 83330502                 ld t2, 0x20(a0)
   292: b3836300                 add t2, t2, t1
   296: 87810302                 vle8.v v3, (t2)
   300: d780302e                 vxor.vv v1, v3, v1
   304: 138e0f08                 addi t3, t6, 0x80
   308: 07040e02                 vle8.v v8, (t3)
   312: 138e0f09                 addi t3, t6, 0x90
   316: 87040e02                 vle8.v v9, (t3)
   320: 5782a126                 vand.vv v4, v10, v3
   324: d73232a2                 vsrl.vi v5, v3, 4
   328: 57038232                 vrgather.vv v6, v8, v4
   332: d7839232                 vrgather.vv v7, v9, v5
   336: 5703732e                 vxor.vv v6, v7, v6
   340: 5701612e                 vxor.vv v2, v6, v2
   344: 83338502                 ld t2, 0x28(a0)
   348: b3836300                 add t2, t2, t1
   352: 87810302                 vle8.v v3, (t2)
   356: d780302e                 vxor.vv v1, v3, v1
   360: 138e0f0a                 addi t3, t6, 0xa0
   364: 07040e02                 vle8.v v8, (t3)
   368: 138e0f0b                 addi t3, t6, 0xb0
   372: 87040e02                 vle8.v v9, (t3)
   376: 5782a126                 vand.vv v4, v10, v3
   380: d73232a2                 vsrl.vi v5, v3, 4
   384: 57038232                 vrgather.vv v6, v8, v4
   388: d7839232                 vrgather.vv v7, v9, v5
   392: 5703732e                 vxor.vv v6, v7, v6
   396: 5701612e                 vxor.vv v2, v6, v2
   400: 83330503                 ld t2, 0x30(a0)
   404: b3836300                 add t2, t2, t1
   408: 87810302                 vle8.v v3, (t2)
   412: d780302e                 vxor.vv v1, v3, v1
   416: 138e0f0c                 addi t3, t6, 0xc0
   420: 07040e02                 vle8.v v8, (t3)
   424: 138e0f0d                 addi t3, t6, 0xd0
   428: 87040e02                 vle8.v v9, (t3)
   432: 5782a126                 vand.vv v4, v10, v3
   436: d73232a2                 vsrl.vi v5, v3, 4
   440: 57038232                 vrgather.vv v6, v8, v4
   444: d7839232                 vrgather.vv v7, v9, v5
   448: 5703732e                 vxor.vv v6, v7, v6
   452: 5701612e                 vxor.vv v2, v6, v2
   456: 83338503                 ld t2, 0x38(a0)
   460: b3836300                 add t2, t2, t1
   464: 87810302                 vle8.v v3, (t2)
   468: d780302e                 vxor.vv v1, v3, v1
   472: 5701312e                 vxor.vv v2, v3, v2
   476: 83b30500                 ld t2, 0(a1)
   480: b3836300                 add t2, t2, t1
   484: a7800302                 vse8.v v1, (t2)
   488: 83b38500                 ld t2, 8(a1)
   492: b3836300                 add t2, t2, t1
   496: 27810302                 vse8.v v2, (t2)
   500: 3303d300                 add t1, t1, a3
   504: 1307f7ff                 addi a4, a4, -1
   508: e31207e4                 bnez a4, 0x40
   512: 33066640                 sub a2, a2, t1
   516: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
   520: 93060600                 mv a3, a2
   524: 6344d000                 bgtz a3, 0x214
   528: 6f00001d                 j 0x3e0
   532: b3835640                 sub t2, a3, t0
   536: 13def343                 srai t3, t2, 0x3f
   540: b3f3c301                 and t2, t2, t3
   544: b3837200                 add t2, t0, t2
   548: d7a10852                 vid.v v3
   552: 57c0336a                 vmsltu.vx v0, v3, t2
   556: 83330500                 ld t2, 0(a0)
   560: b3836300                 add t2, t2, t1
   564: 87810300                 vle8.v v3, (t2), v0.t
   568: d780015e                 vmv.v.v v1, v3
   572: 138e0f00                 mv t3, t6
   576: 07040e02                 vle8.v v8, (t3)
   580: 138e0f01                 addi t3, t6, 0x10
   584: 87040e02                 vle8.v v9, (t3)
   588: 5782a126                 vand.vv v4, v10, v3
   592: d73232a2                 vsrl.vi v5, v3, 4
   596: 57038232                 vrgather.vv v6, v8, v4
   600: d7839232                 vrgather.vv v7, v9, v5
   604: 5703732e                 vxor.vv v6, v7, v6
   608: 5701035e                 vmv.v.v v2, v6
   612: 83338500                 ld t2, 8(a0)
   616: b3836300                 add t2, t2, t1
   620: 87810300                 vle8.v v3, (t2), v0.t
   624: d780302e                 vxor.vv v1, v3, v1
   628: 138e0f02                 addi t3, t6, 0x20
   632: 07040e02                 vle8.v v8, (t3)
   636: 138e0f03                 addi t3, t6, 0x30
   640: 87040e02                 vle8.v v9, (t3)
   644: 5782a126                 vand.vv v4, v10, v3
   648: d73232a2                 vsrl.vi v5, v3, 4
   652: 57038232                 vrgather.vv v6, v8, v4
   656: d7839232                 vrgather.vv v7, v9, v5
   660: 5703732e                 vxor.vv v6, v7, v6
   664: 5701612e                 vxor.vv v2, v6, v2
   668: 83330501                 ld t2, 0x10(a0)
   672: b3836300                 add t2, t2, t1
   676: 87810300                 vle8.v v3, (t2), v0.t
   680: d780302e                 vxor.vv v1, v3, v1
   684: 138e0f04                 addi t3, t6, 0x40
   688: 07040e02                 vle8.v v8, (t3)
   692: 138e0f05                 addi t3, t6, 0x50
   696: 87040e02                 vle8.v v9, (t3)
   700: 5782a126                 vand.vv v4, v10, v3
   704: d73232a2                 vsrl.vi v5, v3, 4
   708: 57038232                 vrgather.vv v6, v8, v4
   712: d7839232                 vrgather.vv v7, v9, v5
   716: 5703732e                 vxor.vv v6, v7, v6
   720: 5701612e                 vxor.vv v2, v6, v2
   724: 83338501                 ld t2, 0x18(a0)
   728: b3836300                 add t2, t2, t1
   732: 87810300                 vle8.v v3, (t2), v0.t
   736: d780302e                 vxor.vv v1, v3, v1
   740: 138e0f06                 addi t3, t6, 0x60
   744: 07040e02                 vle8.v v8, (t3)
   748: 138e0f07                 addi t3, t6, 0x70
   752: 87040e02                 vle8.v v9, (t3)
   756: 5782a126                 vand.vv v4, v10, v3
   760: d73232a2                 vsrl.vi v5, v3, 4
   764: 57038232                 vrgather.vv v6, v8, v4
   768: d7839232                 vrgather.vv v7, v9, v5
   772: 5703732e                 vxor.vv v6, v7, v6
   776: 5701612e                 vxor.vv v2, v6, v2
   780: 83330502                 ld t2, 0x20(a0)
   784: b3836300                 add t2, t2, t1
   788: 87810300                 vle8.v v3, (t2), v0.t
   792: d780302e                 vxor.vv v1, v3, v1
   796: 138e0f08                 addi t3, t6, 0x80
   800: 07040e02                 vle8.v v8, (t3)
   804: 138e0f09                 addi t3, t6, 0x90
   808: 87040e02                 vle8.v v9, (t3)
   812: 5782a126                 vand.vv v4, v10, v3
   816: d73232a2                 vsrl.vi v5, v3, 4
   820: 57038232                 vrgather.vv v6, v8, v4
   824: d7839232                 vrgather.vv v7, v9, v5
   828: 5703732e                 vxor.vv v6, v7, v6
   832: 5701612e                 vxor.vv v2, v6, v2
   836: 83338502                 ld t2, 0x28(a0)
   840: b3836300                 add t2, t2, t1
   844: 87810300                 vle8.v v3, (t2), v0.t
   848: d780302e                 vxor.vv v1, v3, v1
   852: 138e0f0a                 addi t3, t6, 0xa0
   856: 07040e02                 vle8.v v8, (t3)
   860: 138e0f0b                 addi t3, t6, 0xb0
   864: 87040e02                 vle8.v v9, (t3)
   868: 5782a126                 vand.vv v4, v10, v3
   872: d73232a2                 vsrl.vi v5, v3, 4
   876: 57038232                 vrgather.vv v6, v8, v4
   880: d7839232                 vrgather.vv v7, v9, v5
   884: 5703732e                 vxor.vv v6, v7, v6
   888: 5701612e                 vxor.vv v2, v6, v2
   892: 83330503                 ld t2, 0x30(a0)
   896: b3836300                 add t2, t2, t1
   900: 87810300                 vle8.v v3, (t2), v0.t
   904: d780302e                 vxor.vv v1, v3, v1
   908: 138e0f0c                 addi t3, t6, 0xc0
   912: 07040e02                 vle8.v v8, (t3)
   916: 138e0f0d                 addi t3, t6, 0xd0
   920: 87040e02                 vle8.v v9, (t3)
   924: 5782a126                 vand.vv v4, v10, v3
   928: d73232a2                 vsrl.vi v5, v3, 4
   932: 57038232                 vrgather.vv v6, v8, v4
   936: d7839232                 vrgather.vv v7, v9, v5
   940: 5703732e                 vxor.vv v6, v7, v6
   944: 5701612e                 vxor.vv v2, v6, v2
   948: 83338503                 ld t2, 0x38(a0)
   952: b3836300                 add t2, t2, t1
   956: 87810300                 vle8.v v3, (t2), v0.t
   960: d780302e                 vxor.vv v1, v3, v1
   964: 5701312e                 vxor.vv v2, v3, v2
   968: 83b30500                 ld t2, 0(a1)
   972: b3836300                 add t2, t2, t1
   976: a7800300                 vse8.v v1, (t2), v0.t
   980: 83b38500                 ld t2, 8(a1)
   984: b3836300                 add t2, t2, t1
   988: 27810300                 vse8.v v2, (t2), v0.t
   992: 33035300                 add t1, t1, t0
   996: b3865640                 sub a3, a3, t0
  1000: 67800000                 ret 
  1004: 0000                     unimp 
  1006: 0000                     unimp 
