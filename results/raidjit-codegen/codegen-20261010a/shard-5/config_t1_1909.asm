; config_t1_1909 target=rvv-zvbc32e twin=pipe topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1376 twins=["pipe"] uops_per_byte=17.6875
     0: 63140600                 bnez a2, 8
     4: 6f004055                 j 0x558
     8: 970f0000                 auipc t6, 0
    12: 938f8f55                 addi t6, t6, 0x558
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 57450e5e                 vmv.v.x v10, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00c025                 j 0x28c
    52: 83330500                 ld t2, 0(a0)
    56: b3836300                 add t2, t2, t1
    60: 87810302                 vle8.v v3, (t2)
    64: 138e0f00                 mv t3, t6
    68: 07040e02                 vle8.v v8, (t3)
    72: 138e0f01                 addi t3, t6, 0x10
    76: 87040e02                 vle8.v v9, (t3)
    80: 5782a126                 vand.vv v4, v10, v3
    84: d73232a2                 vsrl.vi v5, v3, 4
    88: 57038232                 vrgather.vv v6, v8, v4
    92: d7839232                 vrgather.vv v7, v9, v5
    96: 5703732e                 vxor.vv v6, v7, v6
   100: 5700035e                 vmv.v.v v0, v6
   104: 138e0f02                 addi t3, t6, 0x20
   108: 07040e02                 vle8.v v8, (t3)
   112: 138e0f03                 addi t3, t6, 0x30
   116: 87040e02                 vle8.v v9, (t3)
   120: 5782a126                 vand.vv v4, v10, v3
   124: d73232a2                 vsrl.vi v5, v3, 4
   128: 57038232                 vrgather.vv v6, v8, v4
   132: d7839232                 vrgather.vv v7, v9, v5
   136: 5703732e                 vxor.vv v6, v7, v6
   140: d700035e                 vmv.v.v v1, v6
   144: 83338500                 ld t2, 8(a0)
   148: b3836300                 add t2, t2, t1
   152: 07810302                 vle8.v v2, (t2)
   156: 138e0f04                 addi t3, t6, 0x40
   160: 07040e02                 vle8.v v8, (t3)
   164: 138e0f05                 addi t3, t6, 0x50
   168: 87040e02                 vle8.v v9, (t3)
   172: 5702a126                 vand.vv v4, v10, v2
   176: d73222a2                 vsrl.vi v5, v2, 4
   180: 57038232                 vrgather.vv v6, v8, v4
   184: d7839232                 vrgather.vv v7, v9, v5
   188: 5703732e                 vxor.vv v6, v7, v6
   192: 5700602e                 vxor.vv v0, v6, v0
   196: 138e0f06                 addi t3, t6, 0x60
   200: 07040e02                 vle8.v v8, (t3)
   204: 138e0f07                 addi t3, t6, 0x70
   208: 87040e02                 vle8.v v9, (t3)
   212: 5702a126                 vand.vv v4, v10, v2
   216: d73222a2                 vsrl.vi v5, v2, 4
   220: 57038232                 vrgather.vv v6, v8, v4
   224: d7839232                 vrgather.vv v7, v9, v5
   228: 5703732e                 vxor.vv v6, v7, v6
   232: d780602e                 vxor.vv v1, v6, v1
   236: 83330501                 ld t2, 0x10(a0)
   240: b3836300                 add t2, t2, t1
   244: 07810302                 vle8.v v2, (t2)
   248: 138e0f08                 addi t3, t6, 0x80
   252: 07040e02                 vle8.v v8, (t3)
   256: 138e0f09                 addi t3, t6, 0x90
   260: 87040e02                 vle8.v v9, (t3)
   264: 5702a126                 vand.vv v4, v10, v2
   268: d73222a2                 vsrl.vi v5, v2, 4
   272: 57038232                 vrgather.vv v6, v8, v4
   276: d7839232                 vrgather.vv v7, v9, v5
   280: 5703732e                 vxor.vv v6, v7, v6
   284: 5700602e                 vxor.vv v0, v6, v0
   288: 138e0f0a                 addi t3, t6, 0xa0
   292: 07040e02                 vle8.v v8, (t3)
   296: 138e0f0b                 addi t3, t6, 0xb0
   300: 87040e02                 vle8.v v9, (t3)
   304: 5702a126                 vand.vv v4, v10, v2
   308: d73222a2                 vsrl.vi v5, v2, 4
   312: 57038232                 vrgather.vv v6, v8, v4
   316: d7839232                 vrgather.vv v7, v9, v5
   320: 5703732e                 vxor.vv v6, v7, v6
   324: d780602e                 vxor.vv v1, v6, v1
   328: 83338501                 ld t2, 0x18(a0)
   332: b3836300                 add t2, t2, t1
   336: 07810302                 vle8.v v2, (t2)
   340: 138e0f0c                 addi t3, t6, 0xc0
   344: 07040e02                 vle8.v v8, (t3)
   348: 138e0f0d                 addi t3, t6, 0xd0
   352: 87040e02                 vle8.v v9, (t3)
   356: 5702a126                 vand.vv v4, v10, v2
   360: d73222a2                 vsrl.vi v5, v2, 4
   364: 57038232                 vrgather.vv v6, v8, v4
   368: d7839232                 vrgather.vv v7, v9, v5
   372: 5703732e                 vxor.vv v6, v7, v6
   376: 5700602e                 vxor.vv v0, v6, v0
   380: 138e0f0e                 addi t3, t6, 0xe0
   384: 07040e02                 vle8.v v8, (t3)
   388: 138e0f0f                 addi t3, t6, 0xf0
   392: 87040e02                 vle8.v v9, (t3)
   396: 5702a126                 vand.vv v4, v10, v2
   400: d73222a2                 vsrl.vi v5, v2, 4
   404: 57038232                 vrgather.vv v6, v8, v4
   408: d7839232                 vrgather.vv v7, v9, v5
   412: 5703732e                 vxor.vv v6, v7, v6
   416: d780602e                 vxor.vv v1, v6, v1
   420: 83330502                 ld t2, 0x20(a0)
   424: b3836300                 add t2, t2, t1
   428: 07810302                 vle8.v v2, (t2)
   432: 138e0f10                 addi t3, t6, 0x100
   436: 07040e02                 vle8.v v8, (t3)
   440: 138e0f11                 addi t3, t6, 0x110
   444: 87040e02                 vle8.v v9, (t3)
   448: 5702a126                 vand.vv v4, v10, v2
   452: d73222a2                 vsrl.vi v5, v2, 4
   456: 57038232                 vrgather.vv v6, v8, v4
   460: d7839232                 vrgather.vv v7, v9, v5
   464: 5703732e                 vxor.vv v6, v7, v6
   468: 5700602e                 vxor.vv v0, v6, v0
   472: 138e0f12                 addi t3, t6, 0x120
   476: 07040e02                 vle8.v v8, (t3)
   480: 138e0f13                 addi t3, t6, 0x130
   484: 87040e02                 vle8.v v9, (t3)
   488: 5702a126                 vand.vv v4, v10, v2
   492: d73222a2                 vsrl.vi v5, v2, 4
   496: 57038232                 vrgather.vv v6, v8, v4
   500: d7839232                 vrgather.vv v7, v9, v5
   504: 5703732e                 vxor.vv v6, v7, v6
   508: d780602e                 vxor.vv v1, v6, v1
   512: 83338502                 ld t2, 0x28(a0)
   516: b3836300                 add t2, t2, t1
   520: 07810302                 vle8.v v2, (t2)
   524: 138e0f14                 addi t3, t6, 0x140
   528: 07040e02                 vle8.v v8, (t3)
   532: 138e0f15                 addi t3, t6, 0x150
   536: 87040e02                 vle8.v v9, (t3)
   540: 5702a126                 vand.vv v4, v10, v2
   544: d73222a2                 vsrl.vi v5, v2, 4
   548: 57038232                 vrgather.vv v6, v8, v4
   552: d7839232                 vrgather.vv v7, v9, v5
   556: 5703732e                 vxor.vv v6, v7, v6
   560: 5700602e                 vxor.vv v0, v6, v0
   564: 5702a126                 vand.vv v4, v10, v2
   568: d73222a2                 vsrl.vi v5, v2, 4
   572: 57038232                 vrgather.vv v6, v8, v4
   576: d7839232                 vrgather.vv v7, v9, v5
   580: 5703732e                 vxor.vv v6, v7, v6
   584: d780602e                 vxor.vv v1, v6, v1
   588: 130ff7ff                 addi t5, a4, -1
   592: 630a0f00                 beqz t5, 0x264
   596: 330fd300                 add t5, t1, a3
   600: 83330500                 ld t2, 0(a0)
   604: b383e301                 add t2, t2, t5
   608: 87810302                 vle8.v v3, (t2)
   612: 83b30500                 ld t2, 0(a1)
   616: b3836300                 add t2, t2, t1
   620: 27800302                 vse8.v v0, (t2)
   624: 83b38500                 ld t2, 8(a1)
   628: b3836300                 add t2, t2, t1
   632: a7800302                 vse8.v v1, (t2)
   636: 3303d300                 add t1, t1, a3
   640: 1307f7ff                 addi a4, a4, -1
   644: e31e07da                 bnez a4, 0x40
   648: 33066640                 sub a2, a2, t1
   652: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   656: 63940200                 bnez t0, 0x298
   660: 6f00402c                 j 0x558
   664: 83330500                 ld t2, 0(a0)
   668: b3836300                 add t2, t2, t1
   672: 07810302                 vle8.v v2, (t2)
   676: 138f0200                 mv t5, t0
   680: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   684: 138e0f00                 mv t3, t6
   688: 07040e02                 vle8.v v8, (t3)
   692: 138e0f01                 addi t3, t6, 0x10
   696: 87040e02                 vle8.v v9, (t3)
   700: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   704: 5702a126                 vand.vv v4, v10, v2
   708: d73222a2                 vsrl.vi v5, v2, 4
   712: 57038232                 vrgather.vv v6, v8, v4
   716: d7839232                 vrgather.vv v7, v9, v5
   720: 5703732e                 vxor.vv v6, v7, v6
   724: 5700035e                 vmv.v.v v0, v6
   728: 138f0200                 mv t5, t0
   732: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   736: 138e0f02                 addi t3, t6, 0x20
   740: 07040e02                 vle8.v v8, (t3)
   744: 138e0f03                 addi t3, t6, 0x30
   748: 87040e02                 vle8.v v9, (t3)
   752: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   756: 5702a126                 vand.vv v4, v10, v2
   760: d73222a2                 vsrl.vi v5, v2, 4
   764: 57038232                 vrgather.vv v6, v8, v4
   768: d7839232                 vrgather.vv v7, v9, v5
   772: 5703732e                 vxor.vv v6, v7, v6
   776: d700035e                 vmv.v.v v1, v6
   780: 83338500                 ld t2, 8(a0)
   784: b3836300                 add t2, t2, t1
   788: 07810302                 vle8.v v2, (t2)
   792: 138f0200                 mv t5, t0
   796: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   800: 138e0f04                 addi t3, t6, 0x40
   804: 07040e02                 vle8.v v8, (t3)
   808: 138e0f05                 addi t3, t6, 0x50
   812: 87040e02                 vle8.v v9, (t3)
   816: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   820: 5702a126                 vand.vv v4, v10, v2
   824: d73222a2                 vsrl.vi v5, v2, 4
   828: 57038232                 vrgather.vv v6, v8, v4
   832: d7839232                 vrgather.vv v7, v9, v5
   836: 5703732e                 vxor.vv v6, v7, v6
   840: 5700602e                 vxor.vv v0, v6, v0
   844: 138f0200                 mv t5, t0
   848: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   852: 138e0f06                 addi t3, t6, 0x60
   856: 07040e02                 vle8.v v8, (t3)
   860: 138e0f07                 addi t3, t6, 0x70
   864: 87040e02                 vle8.v v9, (t3)
   868: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   872: 5702a126                 vand.vv v4, v10, v2
   876: d73222a2                 vsrl.vi v5, v2, 4
   880: 57038232                 vrgather.vv v6, v8, v4
   884: d7839232                 vrgather.vv v7, v9, v5
   888: 5703732e                 vxor.vv v6, v7, v6
   892: d780602e                 vxor.vv v1, v6, v1
   896: 83330501                 ld t2, 0x10(a0)
   900: b3836300                 add t2, t2, t1
   904: 07810302                 vle8.v v2, (t2)
   908: 138f0200                 mv t5, t0
   912: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   916: 138e0f08                 addi t3, t6, 0x80
   920: 07040e02                 vle8.v v8, (t3)
   924: 138e0f09                 addi t3, t6, 0x90
   928: 87040e02                 vle8.v v9, (t3)
   932: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   936: 5702a126                 vand.vv v4, v10, v2
   940: d73222a2                 vsrl.vi v5, v2, 4
   944: 57038232                 vrgather.vv v6, v8, v4
   948: d7839232                 vrgather.vv v7, v9, v5
   952: 5703732e                 vxor.vv v6, v7, v6
   956: 5700602e                 vxor.vv v0, v6, v0
   960: 138f0200                 mv t5, t0
   964: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   968: 138e0f0a                 addi t3, t6, 0xa0
   972: 07040e02                 vle8.v v8, (t3)
   976: 138e0f0b                 addi t3, t6, 0xb0
   980: 87040e02                 vle8.v v9, (t3)
   984: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   988: 5702a126                 vand.vv v4, v10, v2
   992: d73222a2                 vsrl.vi v5, v2, 4
   996: 57038232                 vrgather.vv v6, v8, v4
  1000: d7839232                 vrgather.vv v7, v9, v5
  1004: 5703732e                 vxor.vv v6, v7, v6
  1008: d780602e                 vxor.vv v1, v6, v1
  1012: 83338501                 ld t2, 0x18(a0)
  1016: b3836300                 add t2, t2, t1
  1020: 07810302                 vle8.v v2, (t2)
  1024: 138f0200                 mv t5, t0
  1028: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1032: 138e0f0c                 addi t3, t6, 0xc0
  1036: 07040e02                 vle8.v v8, (t3)
  1040: 138e0f0d                 addi t3, t6, 0xd0
  1044: 87040e02                 vle8.v v9, (t3)
  1048: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1052: 5702a126                 vand.vv v4, v10, v2
  1056: d73222a2                 vsrl.vi v5, v2, 4
  1060: 57038232                 vrgather.vv v6, v8, v4
  1064: d7839232                 vrgather.vv v7, v9, v5
  1068: 5703732e                 vxor.vv v6, v7, v6
  1072: 5700602e                 vxor.vv v0, v6, v0
  1076: 138f0200                 mv t5, t0
  1080: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1084: 138e0f0e                 addi t3, t6, 0xe0
  1088: 07040e02                 vle8.v v8, (t3)
  1092: 138e0f0f                 addi t3, t6, 0xf0
  1096: 87040e02                 vle8.v v9, (t3)
  1100: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1104: 5702a126                 vand.vv v4, v10, v2
  1108: d73222a2                 vsrl.vi v5, v2, 4
  1112: 57038232                 vrgather.vv v6, v8, v4
  1116: d7839232                 vrgather.vv v7, v9, v5
  1120: 5703732e                 vxor.vv v6, v7, v6
  1124: d780602e                 vxor.vv v1, v6, v1
  1128: 83330502                 ld t2, 0x20(a0)
  1132: b3836300                 add t2, t2, t1
  1136: 07810302                 vle8.v v2, (t2)
  1140: 138f0200                 mv t5, t0
  1144: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1148: 138e0f10                 addi t3, t6, 0x100
  1152: 07040e02                 vle8.v v8, (t3)
  1156: 138e0f11                 addi t3, t6, 0x110
  1160: 87040e02                 vle8.v v9, (t3)
  1164: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1168: 5702a126                 vand.vv v4, v10, v2
  1172: d73222a2                 vsrl.vi v5, v2, 4
  1176: 57038232                 vrgather.vv v6, v8, v4
  1180: d7839232                 vrgather.vv v7, v9, v5
  1184: 5703732e                 vxor.vv v6, v7, v6
  1188: 5700602e                 vxor.vv v0, v6, v0
  1192: 138f0200                 mv t5, t0
  1196: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1200: 138e0f12                 addi t3, t6, 0x120
  1204: 07040e02                 vle8.v v8, (t3)
  1208: 138e0f13                 addi t3, t6, 0x130
  1212: 87040e02                 vle8.v v9, (t3)
  1216: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1220: 5702a126                 vand.vv v4, v10, v2
  1224: d73222a2                 vsrl.vi v5, v2, 4
  1228: 57038232                 vrgather.vv v6, v8, v4
  1232: d7839232                 vrgather.vv v7, v9, v5
  1236: 5703732e                 vxor.vv v6, v7, v6
  1240: d780602e                 vxor.vv v1, v6, v1
  1244: 83338502                 ld t2, 0x28(a0)
  1248: b3836300                 add t2, t2, t1
  1252: 07810302                 vle8.v v2, (t2)
  1256: 138f0200                 mv t5, t0
  1260: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1264: 138e0f14                 addi t3, t6, 0x140
  1268: 07040e02                 vle8.v v8, (t3)
  1272: 138e0f15                 addi t3, t6, 0x150
  1276: 87040e02                 vle8.v v9, (t3)
  1280: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1284: 5702a126                 vand.vv v4, v10, v2
  1288: d73222a2                 vsrl.vi v5, v2, 4
  1292: 57038232                 vrgather.vv v6, v8, v4
  1296: d7839232                 vrgather.vv v7, v9, v5
  1300: 5703732e                 vxor.vv v6, v7, v6
  1304: 5700602e                 vxor.vv v0, v6, v0
  1308: 5702a126                 vand.vv v4, v10, v2
  1312: d73222a2                 vsrl.vi v5, v2, 4
  1316: 57038232                 vrgather.vv v6, v8, v4
  1320: d7839232                 vrgather.vv v7, v9, v5
  1324: 5703732e                 vxor.vv v6, v7, v6
  1328: d780602e                 vxor.vv v1, v6, v1
  1332: 83b30500                 ld t2, 0(a1)
  1336: b3836300                 add t2, t2, t1
  1340: 27800302                 vse8.v v0, (t2)
  1344: 83b38500                 ld t2, 8(a1)
  1348: b3836300                 add t2, t2, t1
  1352: a7800302                 vse8.v v1, (t2)
  1356: 33035300                 add t1, t1, t0
  1360: 33065640                 sub a2, a2, t0
  1364: 6ff09fd3                 j 0x28c
  1368: 67800000                 ret 
  1372: 0000                     unimp 
  1374: 0000                     unimp 
