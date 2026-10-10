; config_t1_1904 target=rvv-zvbc32e twin=maskedtail topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1264 twins=["maskedtail"] uops_per_byte=16.1875
     0: 63140600                 bnez a2, 8
     4: 6f00c04d                 j 0x4e0
     8: 970f0000                 auipc t6, 0
    12: 938f8f4e                 addi t6, t6, 0x4e8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 57450e5e                 vmv.v.x v10, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f000025                 j 0x280
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 87810302                 vle8.v v3, (t2)
    76: 138e0f00                 mv t3, t6
    80: 07040e02                 vle8.v v8, (t3)
    84: 138e0f01                 addi t3, t6, 0x10
    88: 87040e02                 vle8.v v9, (t3)
    92: 5782a126                 vand.vv v4, v10, v3
    96: d73232a2                 vsrl.vi v5, v3, 4
   100: 57038232                 vrgather.vv v6, v8, v4
   104: d7839232                 vrgather.vv v7, v9, v5
   108: 5703732e                 vxor.vv v6, v7, v6
   112: d700035e                 vmv.v.v v1, v6
   116: 138e0f02                 addi t3, t6, 0x20
   120: 07040e02                 vle8.v v8, (t3)
   124: 138e0f03                 addi t3, t6, 0x30
   128: 87040e02                 vle8.v v9, (t3)
   132: 5782a126                 vand.vv v4, v10, v3
   136: d73232a2                 vsrl.vi v5, v3, 4
   140: 57038232                 vrgather.vv v6, v8, v4
   144: d7839232                 vrgather.vv v7, v9, v5
   148: 5703732e                 vxor.vv v6, v7, v6
   152: 5701035e                 vmv.v.v v2, v6
   156: 83338500                 ld t2, 8(a0)
   160: b3836300                 add t2, t2, t1
   164: 87810302                 vle8.v v3, (t2)
   168: 138e0f04                 addi t3, t6, 0x40
   172: 07040e02                 vle8.v v8, (t3)
   176: 138e0f05                 addi t3, t6, 0x50
   180: 87040e02                 vle8.v v9, (t3)
   184: 5782a126                 vand.vv v4, v10, v3
   188: d73232a2                 vsrl.vi v5, v3, 4
   192: 57038232                 vrgather.vv v6, v8, v4
   196: d7839232                 vrgather.vv v7, v9, v5
   200: 5703732e                 vxor.vv v6, v7, v6
   204: d780602e                 vxor.vv v1, v6, v1
   208: 138e0f06                 addi t3, t6, 0x60
   212: 07040e02                 vle8.v v8, (t3)
   216: 138e0f07                 addi t3, t6, 0x70
   220: 87040e02                 vle8.v v9, (t3)
   224: 5782a126                 vand.vv v4, v10, v3
   228: d73232a2                 vsrl.vi v5, v3, 4
   232: 57038232                 vrgather.vv v6, v8, v4
   236: d7839232                 vrgather.vv v7, v9, v5
   240: 5703732e                 vxor.vv v6, v7, v6
   244: 5701612e                 vxor.vv v2, v6, v2
   248: 83330501                 ld t2, 0x10(a0)
   252: b3836300                 add t2, t2, t1
   256: 87810302                 vle8.v v3, (t2)
   260: 138e0f08                 addi t3, t6, 0x80
   264: 07040e02                 vle8.v v8, (t3)
   268: 138e0f09                 addi t3, t6, 0x90
   272: 87040e02                 vle8.v v9, (t3)
   276: 5782a126                 vand.vv v4, v10, v3
   280: d73232a2                 vsrl.vi v5, v3, 4
   284: 57038232                 vrgather.vv v6, v8, v4
   288: d7839232                 vrgather.vv v7, v9, v5
   292: 5703732e                 vxor.vv v6, v7, v6
   296: d780602e                 vxor.vv v1, v6, v1
   300: 138e0f0a                 addi t3, t6, 0xa0
   304: 07040e02                 vle8.v v8, (t3)
   308: 138e0f0b                 addi t3, t6, 0xb0
   312: 87040e02                 vle8.v v9, (t3)
   316: 5782a126                 vand.vv v4, v10, v3
   320: d73232a2                 vsrl.vi v5, v3, 4
   324: 57038232                 vrgather.vv v6, v8, v4
   328: d7839232                 vrgather.vv v7, v9, v5
   332: 5703732e                 vxor.vv v6, v7, v6
   336: 5701612e                 vxor.vv v2, v6, v2
   340: 83338501                 ld t2, 0x18(a0)
   344: b3836300                 add t2, t2, t1
   348: 87810302                 vle8.v v3, (t2)
   352: 138e0f0c                 addi t3, t6, 0xc0
   356: 07040e02                 vle8.v v8, (t3)
   360: 138e0f0d                 addi t3, t6, 0xd0
   364: 87040e02                 vle8.v v9, (t3)
   368: 5782a126                 vand.vv v4, v10, v3
   372: d73232a2                 vsrl.vi v5, v3, 4
   376: 57038232                 vrgather.vv v6, v8, v4
   380: d7839232                 vrgather.vv v7, v9, v5
   384: 5703732e                 vxor.vv v6, v7, v6
   388: d780602e                 vxor.vv v1, v6, v1
   392: 138e0f0e                 addi t3, t6, 0xe0
   396: 07040e02                 vle8.v v8, (t3)
   400: 138e0f0f                 addi t3, t6, 0xf0
   404: 87040e02                 vle8.v v9, (t3)
   408: 5782a126                 vand.vv v4, v10, v3
   412: d73232a2                 vsrl.vi v5, v3, 4
   416: 57038232                 vrgather.vv v6, v8, v4
   420: d7839232                 vrgather.vv v7, v9, v5
   424: 5703732e                 vxor.vv v6, v7, v6
   428: 5701612e                 vxor.vv v2, v6, v2
   432: 83330502                 ld t2, 0x20(a0)
   436: b3836300                 add t2, t2, t1
   440: 87810302                 vle8.v v3, (t2)
   444: 138e0f10                 addi t3, t6, 0x100
   448: 07040e02                 vle8.v v8, (t3)
   452: 138e0f11                 addi t3, t6, 0x110
   456: 87040e02                 vle8.v v9, (t3)
   460: 5782a126                 vand.vv v4, v10, v3
   464: d73232a2                 vsrl.vi v5, v3, 4
   468: 57038232                 vrgather.vv v6, v8, v4
   472: d7839232                 vrgather.vv v7, v9, v5
   476: 5703732e                 vxor.vv v6, v7, v6
   480: d780602e                 vxor.vv v1, v6, v1
   484: 138e0f12                 addi t3, t6, 0x120
   488: 07040e02                 vle8.v v8, (t3)
   492: 138e0f13                 addi t3, t6, 0x130
   496: 87040e02                 vle8.v v9, (t3)
   500: 5782a126                 vand.vv v4, v10, v3
   504: d73232a2                 vsrl.vi v5, v3, 4
   508: 57038232                 vrgather.vv v6, v8, v4
   512: d7839232                 vrgather.vv v7, v9, v5
   516: 5703732e                 vxor.vv v6, v7, v6
   520: 5701612e                 vxor.vv v2, v6, v2
   524: 83338502                 ld t2, 0x28(a0)
   528: b3836300                 add t2, t2, t1
   532: 87810302                 vle8.v v3, (t2)
   536: 138e0f14                 addi t3, t6, 0x140
   540: 07040e02                 vle8.v v8, (t3)
   544: 138e0f15                 addi t3, t6, 0x150
   548: 87040e02                 vle8.v v9, (t3)
   552: 5782a126                 vand.vv v4, v10, v3
   556: d73232a2                 vsrl.vi v5, v3, 4
   560: 57038232                 vrgather.vv v6, v8, v4
   564: d7839232                 vrgather.vv v7, v9, v5
   568: 5703732e                 vxor.vv v6, v7, v6
   572: d780602e                 vxor.vv v1, v6, v1
   576: 5782a126                 vand.vv v4, v10, v3
   580: d73232a2                 vsrl.vi v5, v3, 4
   584: 57038232                 vrgather.vv v6, v8, v4
   588: d7839232                 vrgather.vv v7, v9, v5
   592: 5703732e                 vxor.vv v6, v7, v6
   596: 5701612e                 vxor.vv v2, v6, v2
   600: 83b30500                 ld t2, 0(a1)
   604: b3836300                 add t2, t2, t1
   608: a7800302                 vse8.v v1, (t2)
   612: 83b38500                 ld t2, 8(a1)
   616: b3836300                 add t2, t2, t1
   620: 27810302                 vse8.v v2, (t2)
   624: 3303d300                 add t1, t1, a3
   628: 1307f7ff                 addi a4, a4, -1
   632: e31407dc                 bnez a4, 0x40
   636: 33066640                 sub a2, a2, t1
   640: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
   644: 93060600                 mv a3, a2
   648: 6344d000                 bgtz a3, 0x290
   652: 6f00c024                 j 0x4d8
   656: b3835640                 sub t2, a3, t0
   660: 13def343                 srai t3, t2, 0x3f
   664: b3f3c301                 and t2, t2, t3
   668: b3837200                 add t2, t0, t2
   672: d7a10852                 vid.v v3
   676: 57c0336a                 vmsltu.vx v0, v3, t2
   680: 83330500                 ld t2, 0(a0)
   684: b3836300                 add t2, t2, t1
   688: 87810300                 vle8.v v3, (t2), v0.t
   692: 138e0f00                 mv t3, t6
   696: 07040e02                 vle8.v v8, (t3)
   700: 138e0f01                 addi t3, t6, 0x10
   704: 87040e02                 vle8.v v9, (t3)
   708: 5782a126                 vand.vv v4, v10, v3
   712: d73232a2                 vsrl.vi v5, v3, 4
   716: 57038232                 vrgather.vv v6, v8, v4
   720: d7839232                 vrgather.vv v7, v9, v5
   724: 5703732e                 vxor.vv v6, v7, v6
   728: d700035e                 vmv.v.v v1, v6
   732: 138e0f02                 addi t3, t6, 0x20
   736: 07040e02                 vle8.v v8, (t3)
   740: 138e0f03                 addi t3, t6, 0x30
   744: 87040e02                 vle8.v v9, (t3)
   748: 5782a126                 vand.vv v4, v10, v3
   752: d73232a2                 vsrl.vi v5, v3, 4
   756: 57038232                 vrgather.vv v6, v8, v4
   760: d7839232                 vrgather.vv v7, v9, v5
   764: 5703732e                 vxor.vv v6, v7, v6
   768: 5701035e                 vmv.v.v v2, v6
   772: 83338500                 ld t2, 8(a0)
   776: b3836300                 add t2, t2, t1
   780: 87810300                 vle8.v v3, (t2), v0.t
   784: 138e0f04                 addi t3, t6, 0x40
   788: 07040e02                 vle8.v v8, (t3)
   792: 138e0f05                 addi t3, t6, 0x50
   796: 87040e02                 vle8.v v9, (t3)
   800: 5782a126                 vand.vv v4, v10, v3
   804: d73232a2                 vsrl.vi v5, v3, 4
   808: 57038232                 vrgather.vv v6, v8, v4
   812: d7839232                 vrgather.vv v7, v9, v5
   816: 5703732e                 vxor.vv v6, v7, v6
   820: d780602e                 vxor.vv v1, v6, v1
   824: 138e0f06                 addi t3, t6, 0x60
   828: 07040e02                 vle8.v v8, (t3)
   832: 138e0f07                 addi t3, t6, 0x70
   836: 87040e02                 vle8.v v9, (t3)
   840: 5782a126                 vand.vv v4, v10, v3
   844: d73232a2                 vsrl.vi v5, v3, 4
   848: 57038232                 vrgather.vv v6, v8, v4
   852: d7839232                 vrgather.vv v7, v9, v5
   856: 5703732e                 vxor.vv v6, v7, v6
   860: 5701612e                 vxor.vv v2, v6, v2
   864: 83330501                 ld t2, 0x10(a0)
   868: b3836300                 add t2, t2, t1
   872: 87810300                 vle8.v v3, (t2), v0.t
   876: 138e0f08                 addi t3, t6, 0x80
   880: 07040e02                 vle8.v v8, (t3)
   884: 138e0f09                 addi t3, t6, 0x90
   888: 87040e02                 vle8.v v9, (t3)
   892: 5782a126                 vand.vv v4, v10, v3
   896: d73232a2                 vsrl.vi v5, v3, 4
   900: 57038232                 vrgather.vv v6, v8, v4
   904: d7839232                 vrgather.vv v7, v9, v5
   908: 5703732e                 vxor.vv v6, v7, v6
   912: d780602e                 vxor.vv v1, v6, v1
   916: 138e0f0a                 addi t3, t6, 0xa0
   920: 07040e02                 vle8.v v8, (t3)
   924: 138e0f0b                 addi t3, t6, 0xb0
   928: 87040e02                 vle8.v v9, (t3)
   932: 5782a126                 vand.vv v4, v10, v3
   936: d73232a2                 vsrl.vi v5, v3, 4
   940: 57038232                 vrgather.vv v6, v8, v4
   944: d7839232                 vrgather.vv v7, v9, v5
   948: 5703732e                 vxor.vv v6, v7, v6
   952: 5701612e                 vxor.vv v2, v6, v2
   956: 83338501                 ld t2, 0x18(a0)
   960: b3836300                 add t2, t2, t1
   964: 87810300                 vle8.v v3, (t2), v0.t
   968: 138e0f0c                 addi t3, t6, 0xc0
   972: 07040e02                 vle8.v v8, (t3)
   976: 138e0f0d                 addi t3, t6, 0xd0
   980: 87040e02                 vle8.v v9, (t3)
   984: 5782a126                 vand.vv v4, v10, v3
   988: d73232a2                 vsrl.vi v5, v3, 4
   992: 57038232                 vrgather.vv v6, v8, v4
   996: d7839232                 vrgather.vv v7, v9, v5
  1000: 5703732e                 vxor.vv v6, v7, v6
  1004: d780602e                 vxor.vv v1, v6, v1
  1008: 138e0f0e                 addi t3, t6, 0xe0
  1012: 07040e02                 vle8.v v8, (t3)
  1016: 138e0f0f                 addi t3, t6, 0xf0
  1020: 87040e02                 vle8.v v9, (t3)
  1024: 5782a126                 vand.vv v4, v10, v3
  1028: d73232a2                 vsrl.vi v5, v3, 4
  1032: 57038232                 vrgather.vv v6, v8, v4
  1036: d7839232                 vrgather.vv v7, v9, v5
  1040: 5703732e                 vxor.vv v6, v7, v6
  1044: 5701612e                 vxor.vv v2, v6, v2
  1048: 83330502                 ld t2, 0x20(a0)
  1052: b3836300                 add t2, t2, t1
  1056: 87810300                 vle8.v v3, (t2), v0.t
  1060: 138e0f10                 addi t3, t6, 0x100
  1064: 07040e02                 vle8.v v8, (t3)
  1068: 138e0f11                 addi t3, t6, 0x110
  1072: 87040e02                 vle8.v v9, (t3)
  1076: 5782a126                 vand.vv v4, v10, v3
  1080: d73232a2                 vsrl.vi v5, v3, 4
  1084: 57038232                 vrgather.vv v6, v8, v4
  1088: d7839232                 vrgather.vv v7, v9, v5
  1092: 5703732e                 vxor.vv v6, v7, v6
  1096: d780602e                 vxor.vv v1, v6, v1
  1100: 138e0f12                 addi t3, t6, 0x120
  1104: 07040e02                 vle8.v v8, (t3)
  1108: 138e0f13                 addi t3, t6, 0x130
  1112: 87040e02                 vle8.v v9, (t3)
  1116: 5782a126                 vand.vv v4, v10, v3
  1120: d73232a2                 vsrl.vi v5, v3, 4
  1124: 57038232                 vrgather.vv v6, v8, v4
  1128: d7839232                 vrgather.vv v7, v9, v5
  1132: 5703732e                 vxor.vv v6, v7, v6
  1136: 5701612e                 vxor.vv v2, v6, v2
  1140: 83338502                 ld t2, 0x28(a0)
  1144: b3836300                 add t2, t2, t1
  1148: 87810300                 vle8.v v3, (t2), v0.t
  1152: 138e0f14                 addi t3, t6, 0x140
  1156: 07040e02                 vle8.v v8, (t3)
  1160: 138e0f15                 addi t3, t6, 0x150
  1164: 87040e02                 vle8.v v9, (t3)
  1168: 5782a126                 vand.vv v4, v10, v3
  1172: d73232a2                 vsrl.vi v5, v3, 4
  1176: 57038232                 vrgather.vv v6, v8, v4
  1180: d7839232                 vrgather.vv v7, v9, v5
  1184: 5703732e                 vxor.vv v6, v7, v6
  1188: d780602e                 vxor.vv v1, v6, v1
  1192: 5782a126                 vand.vv v4, v10, v3
  1196: d73232a2                 vsrl.vi v5, v3, 4
  1200: 57038232                 vrgather.vv v6, v8, v4
  1204: d7839232                 vrgather.vv v7, v9, v5
  1208: 5703732e                 vxor.vv v6, v7, v6
  1212: 5701612e                 vxor.vv v2, v6, v2
  1216: 83b30500                 ld t2, 0(a1)
  1220: b3836300                 add t2, t2, t1
  1224: a7800300                 vse8.v v1, (t2), v0.t
  1228: 83b38500                 ld t2, 8(a1)
  1232: b3836300                 add t2, t2, t1
  1236: 27810300                 vse8.v v2, (t2), v0.t
  1240: 33035300                 add t1, t1, t0
  1244: b3865640                 sub a3, a3, t0
  1248: 67800000                 ret 
  1252: 0000                     unimp 
  1254: 0000                     unimp 
  1256: 0000                     unimp 
  1258: 0000                     unimp 
  1260: 0000                     unimp 
  1262: 0000                     unimp 
