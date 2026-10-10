; config_t1_1899 target=rvv-zvbc32e twin=nt topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1392 twins=["nt"] uops_per_byte=17.6875
     0: 63140600                 bnez a2, 8
     4: 6f008055                 j 0x55c
     8: 970f0000                 auipc t6, 0
    12: 938f8f56                 addi t6, t6, 0x568
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d7440e5e                 vmv.v.x v9, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f008025                 j 0x288
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07810302                 vle8.v v2, (t2)
    76: 138e0f00                 mv t3, t6
    80: 87030e02                 vle8.v v7, (t3)
    84: 138e0f01                 addi t3, t6, 0x10
    88: 07040e02                 vle8.v v8, (t3)
    92: d7019126                 vand.vv v3, v9, v2
    96: 573222a2                 vsrl.vi v4, v2, 4
   100: d7827132                 vrgather.vv v5, v7, v3
   104: 57038232                 vrgather.vv v6, v8, v4
   108: d782622e                 vxor.vv v5, v6, v5
   112: 5780025e                 vmv.v.v v0, v5
   116: 138e0f02                 addi t3, t6, 0x20
   120: 87030e02                 vle8.v v7, (t3)
   124: 138e0f03                 addi t3, t6, 0x30
   128: 07040e02                 vle8.v v8, (t3)
   132: d7019126                 vand.vv v3, v9, v2
   136: 573222a2                 vsrl.vi v4, v2, 4
   140: d7827132                 vrgather.vv v5, v7, v3
   144: 57038232                 vrgather.vv v6, v8, v4
   148: d782622e                 vxor.vv v5, v6, v5
   152: d780025e                 vmv.v.v v1, v5
   156: 83338500                 ld t2, 8(a0)
   160: b3836300                 add t2, t2, t1
   164: 07810302                 vle8.v v2, (t2)
   168: 138e0f04                 addi t3, t6, 0x40
   172: 87030e02                 vle8.v v7, (t3)
   176: 138e0f05                 addi t3, t6, 0x50
   180: 07040e02                 vle8.v v8, (t3)
   184: d7019126                 vand.vv v3, v9, v2
   188: 573222a2                 vsrl.vi v4, v2, 4
   192: d7827132                 vrgather.vv v5, v7, v3
   196: 57038232                 vrgather.vv v6, v8, v4
   200: d782622e                 vxor.vv v5, v6, v5
   204: 5700502e                 vxor.vv v0, v5, v0
   208: 138e0f06                 addi t3, t6, 0x60
   212: 87030e02                 vle8.v v7, (t3)
   216: 138e0f07                 addi t3, t6, 0x70
   220: 07040e02                 vle8.v v8, (t3)
   224: d7019126                 vand.vv v3, v9, v2
   228: 573222a2                 vsrl.vi v4, v2, 4
   232: d7827132                 vrgather.vv v5, v7, v3
   236: 57038232                 vrgather.vv v6, v8, v4
   240: d782622e                 vxor.vv v5, v6, v5
   244: d780502e                 vxor.vv v1, v5, v1
   248: 83330501                 ld t2, 0x10(a0)
   252: b3836300                 add t2, t2, t1
   256: 07810302                 vle8.v v2, (t2)
   260: 138e0f08                 addi t3, t6, 0x80
   264: 87030e02                 vle8.v v7, (t3)
   268: 138e0f09                 addi t3, t6, 0x90
   272: 07040e02                 vle8.v v8, (t3)
   276: d7019126                 vand.vv v3, v9, v2
   280: 573222a2                 vsrl.vi v4, v2, 4
   284: d7827132                 vrgather.vv v5, v7, v3
   288: 57038232                 vrgather.vv v6, v8, v4
   292: d782622e                 vxor.vv v5, v6, v5
   296: 5700502e                 vxor.vv v0, v5, v0
   300: 138e0f0a                 addi t3, t6, 0xa0
   304: 87030e02                 vle8.v v7, (t3)
   308: 138e0f0b                 addi t3, t6, 0xb0
   312: 07040e02                 vle8.v v8, (t3)
   316: d7019126                 vand.vv v3, v9, v2
   320: 573222a2                 vsrl.vi v4, v2, 4
   324: d7827132                 vrgather.vv v5, v7, v3
   328: 57038232                 vrgather.vv v6, v8, v4
   332: d782622e                 vxor.vv v5, v6, v5
   336: d780502e                 vxor.vv v1, v5, v1
   340: 83338501                 ld t2, 0x18(a0)
   344: b3836300                 add t2, t2, t1
   348: 07810302                 vle8.v v2, (t2)
   352: 138e0f0c                 addi t3, t6, 0xc0
   356: 87030e02                 vle8.v v7, (t3)
   360: 138e0f0d                 addi t3, t6, 0xd0
   364: 07040e02                 vle8.v v8, (t3)
   368: d7019126                 vand.vv v3, v9, v2
   372: 573222a2                 vsrl.vi v4, v2, 4
   376: d7827132                 vrgather.vv v5, v7, v3
   380: 57038232                 vrgather.vv v6, v8, v4
   384: d782622e                 vxor.vv v5, v6, v5
   388: 5700502e                 vxor.vv v0, v5, v0
   392: 138e0f0e                 addi t3, t6, 0xe0
   396: 87030e02                 vle8.v v7, (t3)
   400: 138e0f0f                 addi t3, t6, 0xf0
   404: 07040e02                 vle8.v v8, (t3)
   408: d7019126                 vand.vv v3, v9, v2
   412: 573222a2                 vsrl.vi v4, v2, 4
   416: d7827132                 vrgather.vv v5, v7, v3
   420: 57038232                 vrgather.vv v6, v8, v4
   424: d782622e                 vxor.vv v5, v6, v5
   428: d780502e                 vxor.vv v1, v5, v1
   432: 83330502                 ld t2, 0x20(a0)
   436: b3836300                 add t2, t2, t1
   440: 07810302                 vle8.v v2, (t2)
   444: 138e0f10                 addi t3, t6, 0x100
   448: 87030e02                 vle8.v v7, (t3)
   452: 138e0f11                 addi t3, t6, 0x110
   456: 07040e02                 vle8.v v8, (t3)
   460: d7019126                 vand.vv v3, v9, v2
   464: 573222a2                 vsrl.vi v4, v2, 4
   468: d7827132                 vrgather.vv v5, v7, v3
   472: 57038232                 vrgather.vv v6, v8, v4
   476: d782622e                 vxor.vv v5, v6, v5
   480: 5700502e                 vxor.vv v0, v5, v0
   484: 138e0f12                 addi t3, t6, 0x120
   488: 87030e02                 vle8.v v7, (t3)
   492: 138e0f13                 addi t3, t6, 0x130
   496: 07040e02                 vle8.v v8, (t3)
   500: d7019126                 vand.vv v3, v9, v2
   504: 573222a2                 vsrl.vi v4, v2, 4
   508: d7827132                 vrgather.vv v5, v7, v3
   512: 57038232                 vrgather.vv v6, v8, v4
   516: d782622e                 vxor.vv v5, v6, v5
   520: d780502e                 vxor.vv v1, v5, v1
   524: 83338502                 ld t2, 0x28(a0)
   528: b3836300                 add t2, t2, t1
   532: 07810302                 vle8.v v2, (t2)
   536: 138e0f14                 addi t3, t6, 0x140
   540: 87030e02                 vle8.v v7, (t3)
   544: 138e0f15                 addi t3, t6, 0x150
   548: 07040e02                 vle8.v v8, (t3)
   552: d7019126                 vand.vv v3, v9, v2
   556: 573222a2                 vsrl.vi v4, v2, 4
   560: d7827132                 vrgather.vv v5, v7, v3
   564: 57038232                 vrgather.vv v6, v8, v4
   568: d782622e                 vxor.vv v5, v6, v5
   572: 5700502e                 vxor.vv v0, v5, v0
   576: d7019126                 vand.vv v3, v9, v2
   580: 573222a2                 vsrl.vi v4, v2, 4
   584: d7827132                 vrgather.vv v5, v7, v3
   588: 57038232                 vrgather.vv v6, v8, v4
   592: d782622e                 vxor.vv v5, v6, v5
   596: d780502e                 vxor.vv v1, v5, v1
   600: 83b30500                 ld t2, 0(a1)
   604: b3836300                 add t2, t2, t1
   608: 33005000                 ntl.all 
   612: 27800302                 vse8.v v0, (t2)
   616: 83b38500                 ld t2, 8(a1)
   620: b3836300                 add t2, t2, t1
   624: 33005000                 ntl.all 
   628: a7800302                 vse8.v v1, (t2)
   632: 3303d300                 add t1, t1, a3
   636: 1307f7ff                 addi a4, a4, -1
   640: e31007dc                 bnez a4, 0x40
   644: 33066640                 sub a2, a2, t1
   648: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   652: 63940200                 bnez t0, 0x294
   656: 6f00c02c                 j 0x55c
   660: 83330500                 ld t2, 0(a0)
   664: b3836300                 add t2, t2, t1
   668: 07810302                 vle8.v v2, (t2)
   672: 138f0200                 mv t5, t0
   676: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   680: 138e0f00                 mv t3, t6
   684: 87030e02                 vle8.v v7, (t3)
   688: 138e0f01                 addi t3, t6, 0x10
   692: 07040e02                 vle8.v v8, (t3)
   696: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   700: d7019126                 vand.vv v3, v9, v2
   704: 573222a2                 vsrl.vi v4, v2, 4
   708: d7827132                 vrgather.vv v5, v7, v3
   712: 57038232                 vrgather.vv v6, v8, v4
   716: d782622e                 vxor.vv v5, v6, v5
   720: 5780025e                 vmv.v.v v0, v5
   724: 138f0200                 mv t5, t0
   728: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   732: 138e0f02                 addi t3, t6, 0x20
   736: 87030e02                 vle8.v v7, (t3)
   740: 138e0f03                 addi t3, t6, 0x30
   744: 07040e02                 vle8.v v8, (t3)
   748: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   752: d7019126                 vand.vv v3, v9, v2
   756: 573222a2                 vsrl.vi v4, v2, 4
   760: d7827132                 vrgather.vv v5, v7, v3
   764: 57038232                 vrgather.vv v6, v8, v4
   768: d782622e                 vxor.vv v5, v6, v5
   772: d780025e                 vmv.v.v v1, v5
   776: 83338500                 ld t2, 8(a0)
   780: b3836300                 add t2, t2, t1
   784: 07810302                 vle8.v v2, (t2)
   788: 138f0200                 mv t5, t0
   792: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   796: 138e0f04                 addi t3, t6, 0x40
   800: 87030e02                 vle8.v v7, (t3)
   804: 138e0f05                 addi t3, t6, 0x50
   808: 07040e02                 vle8.v v8, (t3)
   812: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   816: d7019126                 vand.vv v3, v9, v2
   820: 573222a2                 vsrl.vi v4, v2, 4
   824: d7827132                 vrgather.vv v5, v7, v3
   828: 57038232                 vrgather.vv v6, v8, v4
   832: d782622e                 vxor.vv v5, v6, v5
   836: 5700502e                 vxor.vv v0, v5, v0
   840: 138f0200                 mv t5, t0
   844: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   848: 138e0f06                 addi t3, t6, 0x60
   852: 87030e02                 vle8.v v7, (t3)
   856: 138e0f07                 addi t3, t6, 0x70
   860: 07040e02                 vle8.v v8, (t3)
   864: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   868: d7019126                 vand.vv v3, v9, v2
   872: 573222a2                 vsrl.vi v4, v2, 4
   876: d7827132                 vrgather.vv v5, v7, v3
   880: 57038232                 vrgather.vv v6, v8, v4
   884: d782622e                 vxor.vv v5, v6, v5
   888: d780502e                 vxor.vv v1, v5, v1
   892: 83330501                 ld t2, 0x10(a0)
   896: b3836300                 add t2, t2, t1
   900: 07810302                 vle8.v v2, (t2)
   904: 138f0200                 mv t5, t0
   908: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   912: 138e0f08                 addi t3, t6, 0x80
   916: 87030e02                 vle8.v v7, (t3)
   920: 138e0f09                 addi t3, t6, 0x90
   924: 07040e02                 vle8.v v8, (t3)
   928: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   932: d7019126                 vand.vv v3, v9, v2
   936: 573222a2                 vsrl.vi v4, v2, 4
   940: d7827132                 vrgather.vv v5, v7, v3
   944: 57038232                 vrgather.vv v6, v8, v4
   948: d782622e                 vxor.vv v5, v6, v5
   952: 5700502e                 vxor.vv v0, v5, v0
   956: 138f0200                 mv t5, t0
   960: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   964: 138e0f0a                 addi t3, t6, 0xa0
   968: 87030e02                 vle8.v v7, (t3)
   972: 138e0f0b                 addi t3, t6, 0xb0
   976: 07040e02                 vle8.v v8, (t3)
   980: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   984: d7019126                 vand.vv v3, v9, v2
   988: 573222a2                 vsrl.vi v4, v2, 4
   992: d7827132                 vrgather.vv v5, v7, v3
   996: 57038232                 vrgather.vv v6, v8, v4
  1000: d782622e                 vxor.vv v5, v6, v5
  1004: d780502e                 vxor.vv v1, v5, v1
  1008: 83338501                 ld t2, 0x18(a0)
  1012: b3836300                 add t2, t2, t1
  1016: 07810302                 vle8.v v2, (t2)
  1020: 138f0200                 mv t5, t0
  1024: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1028: 138e0f0c                 addi t3, t6, 0xc0
  1032: 87030e02                 vle8.v v7, (t3)
  1036: 138e0f0d                 addi t3, t6, 0xd0
  1040: 07040e02                 vle8.v v8, (t3)
  1044: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1048: d7019126                 vand.vv v3, v9, v2
  1052: 573222a2                 vsrl.vi v4, v2, 4
  1056: d7827132                 vrgather.vv v5, v7, v3
  1060: 57038232                 vrgather.vv v6, v8, v4
  1064: d782622e                 vxor.vv v5, v6, v5
  1068: 5700502e                 vxor.vv v0, v5, v0
  1072: 138f0200                 mv t5, t0
  1076: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1080: 138e0f0e                 addi t3, t6, 0xe0
  1084: 87030e02                 vle8.v v7, (t3)
  1088: 138e0f0f                 addi t3, t6, 0xf0
  1092: 07040e02                 vle8.v v8, (t3)
  1096: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1100: d7019126                 vand.vv v3, v9, v2
  1104: 573222a2                 vsrl.vi v4, v2, 4
  1108: d7827132                 vrgather.vv v5, v7, v3
  1112: 57038232                 vrgather.vv v6, v8, v4
  1116: d782622e                 vxor.vv v5, v6, v5
  1120: d780502e                 vxor.vv v1, v5, v1
  1124: 83330502                 ld t2, 0x20(a0)
  1128: b3836300                 add t2, t2, t1
  1132: 07810302                 vle8.v v2, (t2)
  1136: 138f0200                 mv t5, t0
  1140: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1144: 138e0f10                 addi t3, t6, 0x100
  1148: 87030e02                 vle8.v v7, (t3)
  1152: 138e0f11                 addi t3, t6, 0x110
  1156: 07040e02                 vle8.v v8, (t3)
  1160: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1164: d7019126                 vand.vv v3, v9, v2
  1168: 573222a2                 vsrl.vi v4, v2, 4
  1172: d7827132                 vrgather.vv v5, v7, v3
  1176: 57038232                 vrgather.vv v6, v8, v4
  1180: d782622e                 vxor.vv v5, v6, v5
  1184: 5700502e                 vxor.vv v0, v5, v0
  1188: 138f0200                 mv t5, t0
  1192: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1196: 138e0f12                 addi t3, t6, 0x120
  1200: 87030e02                 vle8.v v7, (t3)
  1204: 138e0f13                 addi t3, t6, 0x130
  1208: 07040e02                 vle8.v v8, (t3)
  1212: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1216: d7019126                 vand.vv v3, v9, v2
  1220: 573222a2                 vsrl.vi v4, v2, 4
  1224: d7827132                 vrgather.vv v5, v7, v3
  1228: 57038232                 vrgather.vv v6, v8, v4
  1232: d782622e                 vxor.vv v5, v6, v5
  1236: d780502e                 vxor.vv v1, v5, v1
  1240: 83338502                 ld t2, 0x28(a0)
  1244: b3836300                 add t2, t2, t1
  1248: 07810302                 vle8.v v2, (t2)
  1252: 138f0200                 mv t5, t0
  1256: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1260: 138e0f14                 addi t3, t6, 0x140
  1264: 87030e02                 vle8.v v7, (t3)
  1268: 138e0f15                 addi t3, t6, 0x150
  1272: 07040e02                 vle8.v v8, (t3)
  1276: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1280: d7019126                 vand.vv v3, v9, v2
  1284: 573222a2                 vsrl.vi v4, v2, 4
  1288: d7827132                 vrgather.vv v5, v7, v3
  1292: 57038232                 vrgather.vv v6, v8, v4
  1296: d782622e                 vxor.vv v5, v6, v5
  1300: 5700502e                 vxor.vv v0, v5, v0
  1304: d7019126                 vand.vv v3, v9, v2
  1308: 573222a2                 vsrl.vi v4, v2, 4
  1312: d7827132                 vrgather.vv v5, v7, v3
  1316: 57038232                 vrgather.vv v6, v8, v4
  1320: d782622e                 vxor.vv v5, v6, v5
  1324: d780502e                 vxor.vv v1, v5, v1
  1328: 83b30500                 ld t2, 0(a1)
  1332: b3836300                 add t2, t2, t1
  1336: 33005000                 ntl.all 
  1340: 27800302                 vse8.v v0, (t2)
  1344: 83b38500                 ld t2, 8(a1)
  1348: b3836300                 add t2, t2, t1
  1352: 33005000                 ntl.all 
  1356: a7800302                 vse8.v v1, (t2)
  1360: 33035300                 add t1, t1, t0
  1364: 33065640                 sub a2, a2, t0
  1368: 6ff01fd3                 j 0x288
  1372: 67800000                 ret 
  1376: 0000                     unimp 
  1378: 0000                     unimp 
  1380: 0000                     unimp 
  1382: 0000                     unimp 
  1384: 0000                     unimp 
  1386: 0000                     unimp 
  1388: 0000                     unimp 
  1390: 0000                     unimp 
