; config_t1_1795 target=rvv-zvbc32e twin=nt topo=rs(6,3)/enc op=encode U=1
; code_size=1344 twins=["nt"] uops_per_byte=16.625
     0: 63140600                 bnez a2, 8
     4: 6f00c052                 j 0x530
     8: 970f0000                 auipc t6, 0
    12: 938f8f53                 addi t6, t6, 0x538
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 57450e5e                 vmv.v.x v10, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f008024                 j 0x278
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
   576: 33005000                 ntl.all 
   580: 27800302                 vse8.v v0, (t2)
   584: 83b38500                 ld t2, 8(a1)
   588: b3836300                 add t2, t2, t1
   592: 33005000                 ntl.all 
   596: a7800302                 vse8.v v1, (t2)
   600: 83b30501                 ld t2, 0x10(a1)
   604: b3836300                 add t2, t2, t1
   608: 33005000                 ntl.all 
   612: 27810302                 vse8.v v2, (t2)
   616: 3303d300                 add t1, t1, a3
   620: 1307f7ff                 addi a4, a4, -1
   624: e31807dc                 bnez a4, 0x40
   628: 33066640                 sub a2, a2, t1
   632: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   636: 63940200                 bnez t0, 0x284
   640: 6f00002b                 j 0x530
   644: 83330500                 ld t2, 0(a0)
   648: b3836300                 add t2, t2, t1
   652: 87810302                 vle8.v v3, (t2)
   656: 5780015e                 vmv.v.v v0, v3
   660: 138f0200                 mv t5, t0
   664: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   668: 138e0f00                 mv t3, t6
   672: 07040e02                 vle8.v v8, (t3)
   676: 138e0f01                 addi t3, t6, 0x10
   680: 87040e02                 vle8.v v9, (t3)
   684: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   688: 5782a126                 vand.vv v4, v10, v3
   692: d73232a2                 vsrl.vi v5, v3, 4
   696: 57038232                 vrgather.vv v6, v8, v4
   700: d7839232                 vrgather.vv v7, v9, v5
   704: 5703732e                 vxor.vv v6, v7, v6
   708: d700035e                 vmv.v.v v1, v6
   712: 138f0200                 mv t5, t0
   716: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   720: 138e0f02                 addi t3, t6, 0x20
   724: 07040e02                 vle8.v v8, (t3)
   728: 138e0f03                 addi t3, t6, 0x30
   732: 87040e02                 vle8.v v9, (t3)
   736: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   740: 5782a126                 vand.vv v4, v10, v3
   744: d73232a2                 vsrl.vi v5, v3, 4
   748: 57038232                 vrgather.vv v6, v8, v4
   752: d7839232                 vrgather.vv v7, v9, v5
   756: 5703732e                 vxor.vv v6, v7, v6
   760: 5701035e                 vmv.v.v v2, v6
   764: 83338500                 ld t2, 8(a0)
   768: b3836300                 add t2, t2, t1
   772: 87810302                 vle8.v v3, (t2)
   776: 5700302e                 vxor.vv v0, v3, v0
   780: 138f0200                 mv t5, t0
   784: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   788: 138e0f02                 addi t3, t6, 0x20
   792: 07040e02                 vle8.v v8, (t3)
   796: 138e0f03                 addi t3, t6, 0x30
   800: 87040e02                 vle8.v v9, (t3)
   804: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   808: 5782a126                 vand.vv v4, v10, v3
   812: d73232a2                 vsrl.vi v5, v3, 4
   816: 57038232                 vrgather.vv v6, v8, v4
   820: d7839232                 vrgather.vv v7, v9, v5
   824: 5703732e                 vxor.vv v6, v7, v6
   828: d780602e                 vxor.vv v1, v6, v1
   832: 138f0200                 mv t5, t0
   836: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   840: 138e0f04                 addi t3, t6, 0x40
   844: 07040e02                 vle8.v v8, (t3)
   848: 138e0f05                 addi t3, t6, 0x50
   852: 87040e02                 vle8.v v9, (t3)
   856: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   860: 5782a126                 vand.vv v4, v10, v3
   864: d73232a2                 vsrl.vi v5, v3, 4
   868: 57038232                 vrgather.vv v6, v8, v4
   872: d7839232                 vrgather.vv v7, v9, v5
   876: 5703732e                 vxor.vv v6, v7, v6
   880: 5701612e                 vxor.vv v2, v6, v2
   884: 83330501                 ld t2, 0x10(a0)
   888: b3836300                 add t2, t2, t1
   892: 87810302                 vle8.v v3, (t2)
   896: 5700302e                 vxor.vv v0, v3, v0
   900: 138f0200                 mv t5, t0
   904: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   908: 138e0f04                 addi t3, t6, 0x40
   912: 07040e02                 vle8.v v8, (t3)
   916: 138e0f05                 addi t3, t6, 0x50
   920: 87040e02                 vle8.v v9, (t3)
   924: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   928: 5782a126                 vand.vv v4, v10, v3
   932: d73232a2                 vsrl.vi v5, v3, 4
   936: 57038232                 vrgather.vv v6, v8, v4
   940: d7839232                 vrgather.vv v7, v9, v5
   944: 5703732e                 vxor.vv v6, v7, v6
   948: d780602e                 vxor.vv v1, v6, v1
   952: 138f0200                 mv t5, t0
   956: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   960: 138e0f06                 addi t3, t6, 0x60
   964: 07040e02                 vle8.v v8, (t3)
   968: 138e0f07                 addi t3, t6, 0x70
   972: 87040e02                 vle8.v v9, (t3)
   976: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   980: 5782a126                 vand.vv v4, v10, v3
   984: d73232a2                 vsrl.vi v5, v3, 4
   988: 57038232                 vrgather.vv v6, v8, v4
   992: d7839232                 vrgather.vv v7, v9, v5
   996: 5703732e                 vxor.vv v6, v7, v6
  1000: 5701612e                 vxor.vv v2, v6, v2
  1004: 83338501                 ld t2, 0x18(a0)
  1008: b3836300                 add t2, t2, t1
  1012: 87810302                 vle8.v v3, (t2)
  1016: 5700302e                 vxor.vv v0, v3, v0
  1020: 138f0200                 mv t5, t0
  1024: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1028: 138e0f08                 addi t3, t6, 0x80
  1032: 07040e02                 vle8.v v8, (t3)
  1036: 138e0f09                 addi t3, t6, 0x90
  1040: 87040e02                 vle8.v v9, (t3)
  1044: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1048: 5782a126                 vand.vv v4, v10, v3
  1052: d73232a2                 vsrl.vi v5, v3, 4
  1056: 57038232                 vrgather.vv v6, v8, v4
  1060: d7839232                 vrgather.vv v7, v9, v5
  1064: 5703732e                 vxor.vv v6, v7, v6
  1068: d780602e                 vxor.vv v1, v6, v1
  1072: 138f0200                 mv t5, t0
  1076: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1080: 138e0f00                 mv t3, t6
  1084: 07040e02                 vle8.v v8, (t3)
  1088: 138e0f01                 addi t3, t6, 0x10
  1092: 87040e02                 vle8.v v9, (t3)
  1096: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1100: 5782a126                 vand.vv v4, v10, v3
  1104: d73232a2                 vsrl.vi v5, v3, 4
  1108: 57038232                 vrgather.vv v6, v8, v4
  1112: d7839232                 vrgather.vv v7, v9, v5
  1116: 5703732e                 vxor.vv v6, v7, v6
  1120: 5701612e                 vxor.vv v2, v6, v2
  1124: 83330502                 ld t2, 0x20(a0)
  1128: b3836300                 add t2, t2, t1
  1132: 87810302                 vle8.v v3, (t2)
  1136: 5700302e                 vxor.vv v0, v3, v0
  1140: 138f0200                 mv t5, t0
  1144: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1148: 138e0f06                 addi t3, t6, 0x60
  1152: 07040e02                 vle8.v v8, (t3)
  1156: 138e0f07                 addi t3, t6, 0x70
  1160: 87040e02                 vle8.v v9, (t3)
  1164: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1168: 5782a126                 vand.vv v4, v10, v3
  1172: d73232a2                 vsrl.vi v5, v3, 4
  1176: 57038232                 vrgather.vv v6, v8, v4
  1180: d7839232                 vrgather.vv v7, v9, v5
  1184: 5703732e                 vxor.vv v6, v7, v6
  1188: d780602e                 vxor.vv v1, v6, v1
  1192: 138f0200                 mv t5, t0
  1196: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1200: 138e0f08                 addi t3, t6, 0x80
  1204: 07040e02                 vle8.v v8, (t3)
  1208: 138e0f09                 addi t3, t6, 0x90
  1212: 87040e02                 vle8.v v9, (t3)
  1216: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1220: 5782a126                 vand.vv v4, v10, v3
  1224: d73232a2                 vsrl.vi v5, v3, 4
  1228: 57038232                 vrgather.vv v6, v8, v4
  1232: d7839232                 vrgather.vv v7, v9, v5
  1236: 5703732e                 vxor.vv v6, v7, v6
  1240: 5701612e                 vxor.vv v2, v6, v2
  1244: 83338502                 ld t2, 0x28(a0)
  1248: b3836300                 add t2, t2, t1
  1252: 87810302                 vle8.v v3, (t2)
  1256: 5700302e                 vxor.vv v0, v3, v0
  1260: d780302e                 vxor.vv v1, v3, v1
  1264: 5701312e                 vxor.vv v2, v3, v2
  1268: 83b30500                 ld t2, 0(a1)
  1272: b3836300                 add t2, t2, t1
  1276: 33005000                 ntl.all 
  1280: 27800302                 vse8.v v0, (t2)
  1284: 83b38500                 ld t2, 8(a1)
  1288: b3836300                 add t2, t2, t1
  1292: 33005000                 ntl.all 
  1296: a7800302                 vse8.v v1, (t2)
  1300: 83b30501                 ld t2, 0x10(a1)
  1304: b3836300                 add t2, t2, t1
  1308: 33005000                 ntl.all 
  1312: 27810302                 vse8.v v2, (t2)
  1316: 33035300                 add t1, t1, t0
  1320: 33065640                 sub a2, a2, t0
  1324: 6ff0dfd4                 j 0x278
  1328: 67800000                 ret 
  1332: 0000                     unimp 
  1334: 0000                     unimp 
  1336: 0000                     unimp 
  1338: 0000                     unimp 
  1340: 0000                     unimp 
  1342: 0000                     unimp 
