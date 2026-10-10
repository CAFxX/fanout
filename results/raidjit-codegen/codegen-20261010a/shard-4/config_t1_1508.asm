; config_t1_1508 target=rvv twin=baseline topo=rs(6,3)/enc op=encode U=1
; code_size=1664 twins=[] uops_per_byte=5.359375
     0: 63140600                 bnez a2, 8
     4: 6f00c066                 j 0x670
     8: 970f0000                 auipc t6, 0
    12: 938f8f67                 addi t6, t6, 0x678
    16: 577f100c                 vsetvli t5, zero, e8, m2, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574e0e5e                 vmv.v.x v28, t3
    28: 13030000                 li t1, 0
    32: d772100c                 vsetvli t0, zero, e8, m2, ta, ma
    36: 93961200                 slli a3, t0, 1
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f004039                 j 0x3c4
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07860302                 vle8.v v12, (t2)
    76: b3835300                 add t2, t2, t0
    80: 07870302                 vle8.v v14, (t2)
    84: 5700065e                 vmv.v.v v0, v12
    88: 5701075e                 vmv.v.v v2, v14
    92: 138e0f00                 mv t3, t6
    96: 070c0e02                 vle8.v v24, (t3)
   100: 138e0f02                 addi t3, t6, 0x20
   104: 070d0e02                 vle8.v v26, (t3)
   108: 5708c627                 vand.vv v16, v28, v12
   112: 5739c2a2                 vsrl.vi v18, v12, 4
   116: 570a8833                 vrgather.vv v20, v24, v16
   120: 570ba933                 vrgather.vv v22, v26, v18
   124: 570a6a2f                 vxor.vv v20, v22, v20
   128: 57020a5e                 vmv.v.v v4, v20
   132: 5708c727                 vand.vv v16, v28, v14
   136: 5739e2a2                 vsrl.vi v18, v14, 4
   140: 570a8833                 vrgather.vv v20, v24, v16
   144: 570ba933                 vrgather.vv v22, v26, v18
   148: 570a6a2f                 vxor.vv v20, v22, v20
   152: 57030a5e                 vmv.v.v v6, v20
   156: 138e0f04                 addi t3, t6, 0x40
   160: 070c0e02                 vle8.v v24, (t3)
   164: 138e0f06                 addi t3, t6, 0x60
   168: 070d0e02                 vle8.v v26, (t3)
   172: 5708c627                 vand.vv v16, v28, v12
   176: 5739c2a2                 vsrl.vi v18, v12, 4
   180: 570a8833                 vrgather.vv v20, v24, v16
   184: 570ba933                 vrgather.vv v22, v26, v18
   188: 570a6a2f                 vxor.vv v20, v22, v20
   192: 57040a5e                 vmv.v.v v8, v20
   196: 5708c727                 vand.vv v16, v28, v14
   200: 5739e2a2                 vsrl.vi v18, v14, 4
   204: 570a8833                 vrgather.vv v20, v24, v16
   208: 570ba933                 vrgather.vv v22, v26, v18
   212: 570a6a2f                 vxor.vv v20, v22, v20
   216: 57050a5e                 vmv.v.v v10, v20
   220: 83338500                 ld t2, 8(a0)
   224: b3836300                 add t2, t2, t1
   228: 07860302                 vle8.v v12, (t2)
   232: b3835300                 add t2, t2, t0
   236: 07870302                 vle8.v v14, (t2)
   240: 5700c02e                 vxor.vv v0, v12, v0
   244: 5701e12e                 vxor.vv v2, v14, v2
   248: 138e0f04                 addi t3, t6, 0x40
   252: 070c0e02                 vle8.v v24, (t3)
   256: 138e0f06                 addi t3, t6, 0x60
   260: 070d0e02                 vle8.v v26, (t3)
   264: 5708c627                 vand.vv v16, v28, v12
   268: 5739c2a2                 vsrl.vi v18, v12, 4
   272: 570a8833                 vrgather.vv v20, v24, v16
   276: 570ba933                 vrgather.vv v22, v26, v18
   280: 570a6a2f                 vxor.vv v20, v22, v20
   284: 5702422f                 vxor.vv v4, v20, v4
   288: 5708c727                 vand.vv v16, v28, v14
   292: 5739e2a2                 vsrl.vi v18, v14, 4
   296: 570a8833                 vrgather.vv v20, v24, v16
   300: 570ba933                 vrgather.vv v22, v26, v18
   304: 570a6a2f                 vxor.vv v20, v22, v20
   308: 5703432f                 vxor.vv v6, v20, v6
   312: 138e0f08                 addi t3, t6, 0x80
   316: 070c0e02                 vle8.v v24, (t3)
   320: 138e0f0a                 addi t3, t6, 0xa0
   324: 070d0e02                 vle8.v v26, (t3)
   328: 5708c627                 vand.vv v16, v28, v12
   332: 5739c2a2                 vsrl.vi v18, v12, 4
   336: 570a8833                 vrgather.vv v20, v24, v16
   340: 570ba933                 vrgather.vv v22, v26, v18
   344: 570a6a2f                 vxor.vv v20, v22, v20
   348: 5704442f                 vxor.vv v8, v20, v8
   352: 5708c727                 vand.vv v16, v28, v14
   356: 5739e2a2                 vsrl.vi v18, v14, 4
   360: 570a8833                 vrgather.vv v20, v24, v16
   364: 570ba933                 vrgather.vv v22, v26, v18
   368: 570a6a2f                 vxor.vv v20, v22, v20
   372: 5705452f                 vxor.vv v10, v20, v10
   376: 83330501                 ld t2, 0x10(a0)
   380: b3836300                 add t2, t2, t1
   384: 07860302                 vle8.v v12, (t2)
   388: b3835300                 add t2, t2, t0
   392: 07870302                 vle8.v v14, (t2)
   396: 5700c02e                 vxor.vv v0, v12, v0
   400: 5701e12e                 vxor.vv v2, v14, v2
   404: 138e0f08                 addi t3, t6, 0x80
   408: 070c0e02                 vle8.v v24, (t3)
   412: 138e0f0a                 addi t3, t6, 0xa0
   416: 070d0e02                 vle8.v v26, (t3)
   420: 5708c627                 vand.vv v16, v28, v12
   424: 5739c2a2                 vsrl.vi v18, v12, 4
   428: 570a8833                 vrgather.vv v20, v24, v16
   432: 570ba933                 vrgather.vv v22, v26, v18
   436: 570a6a2f                 vxor.vv v20, v22, v20
   440: 5702422f                 vxor.vv v4, v20, v4
   444: 5708c727                 vand.vv v16, v28, v14
   448: 5739e2a2                 vsrl.vi v18, v14, 4
   452: 570a8833                 vrgather.vv v20, v24, v16
   456: 570ba933                 vrgather.vv v22, v26, v18
   460: 570a6a2f                 vxor.vv v20, v22, v20
   464: 5703432f                 vxor.vv v6, v20, v6
   468: 138e0f0c                 addi t3, t6, 0xc0
   472: 070c0e02                 vle8.v v24, (t3)
   476: 138e0f0e                 addi t3, t6, 0xe0
   480: 070d0e02                 vle8.v v26, (t3)
   484: 5708c627                 vand.vv v16, v28, v12
   488: 5739c2a2                 vsrl.vi v18, v12, 4
   492: 570a8833                 vrgather.vv v20, v24, v16
   496: 570ba933                 vrgather.vv v22, v26, v18
   500: 570a6a2f                 vxor.vv v20, v22, v20
   504: 5704442f                 vxor.vv v8, v20, v8
   508: 5708c727                 vand.vv v16, v28, v14
   512: 5739e2a2                 vsrl.vi v18, v14, 4
   516: 570a8833                 vrgather.vv v20, v24, v16
   520: 570ba933                 vrgather.vv v22, v26, v18
   524: 570a6a2f                 vxor.vv v20, v22, v20
   528: 5705452f                 vxor.vv v10, v20, v10
   532: 83338501                 ld t2, 0x18(a0)
   536: b3836300                 add t2, t2, t1
   540: 07860302                 vle8.v v12, (t2)
   544: b3835300                 add t2, t2, t0
   548: 07870302                 vle8.v v14, (t2)
   552: 5700c02e                 vxor.vv v0, v12, v0
   556: 5701e12e                 vxor.vv v2, v14, v2
   560: 138e0f10                 addi t3, t6, 0x100
   564: 070c0e02                 vle8.v v24, (t3)
   568: 138e0f12                 addi t3, t6, 0x120
   572: 070d0e02                 vle8.v v26, (t3)
   576: 5708c627                 vand.vv v16, v28, v12
   580: 5739c2a2                 vsrl.vi v18, v12, 4
   584: 570a8833                 vrgather.vv v20, v24, v16
   588: 570ba933                 vrgather.vv v22, v26, v18
   592: 570a6a2f                 vxor.vv v20, v22, v20
   596: 5702422f                 vxor.vv v4, v20, v4
   600: 5708c727                 vand.vv v16, v28, v14
   604: 5739e2a2                 vsrl.vi v18, v14, 4
   608: 570a8833                 vrgather.vv v20, v24, v16
   612: 570ba933                 vrgather.vv v22, v26, v18
   616: 570a6a2f                 vxor.vv v20, v22, v20
   620: 5703432f                 vxor.vv v6, v20, v6
   624: 138e0f00                 mv t3, t6
   628: 070c0e02                 vle8.v v24, (t3)
   632: 138e0f02                 addi t3, t6, 0x20
   636: 070d0e02                 vle8.v v26, (t3)
   640: 5708c627                 vand.vv v16, v28, v12
   644: 5739c2a2                 vsrl.vi v18, v12, 4
   648: 570a8833                 vrgather.vv v20, v24, v16
   652: 570ba933                 vrgather.vv v22, v26, v18
   656: 570a6a2f                 vxor.vv v20, v22, v20
   660: 5704442f                 vxor.vv v8, v20, v8
   664: 5708c727                 vand.vv v16, v28, v14
   668: 5739e2a2                 vsrl.vi v18, v14, 4
   672: 570a8833                 vrgather.vv v20, v24, v16
   676: 570ba933                 vrgather.vv v22, v26, v18
   680: 570a6a2f                 vxor.vv v20, v22, v20
   684: 5705452f                 vxor.vv v10, v20, v10
   688: 83330502                 ld t2, 0x20(a0)
   692: b3836300                 add t2, t2, t1
   696: 07860302                 vle8.v v12, (t2)
   700: b3835300                 add t2, t2, t0
   704: 07870302                 vle8.v v14, (t2)
   708: 5700c02e                 vxor.vv v0, v12, v0
   712: 5701e12e                 vxor.vv v2, v14, v2
   716: 138e0f0c                 addi t3, t6, 0xc0
   720: 070c0e02                 vle8.v v24, (t3)
   724: 138e0f0e                 addi t3, t6, 0xe0
   728: 070d0e02                 vle8.v v26, (t3)
   732: 5708c627                 vand.vv v16, v28, v12
   736: 5739c2a2                 vsrl.vi v18, v12, 4
   740: 570a8833                 vrgather.vv v20, v24, v16
   744: 570ba933                 vrgather.vv v22, v26, v18
   748: 570a6a2f                 vxor.vv v20, v22, v20
   752: 5702422f                 vxor.vv v4, v20, v4
   756: 5708c727                 vand.vv v16, v28, v14
   760: 5739e2a2                 vsrl.vi v18, v14, 4
   764: 570a8833                 vrgather.vv v20, v24, v16
   768: 570ba933                 vrgather.vv v22, v26, v18
   772: 570a6a2f                 vxor.vv v20, v22, v20
   776: 5703432f                 vxor.vv v6, v20, v6
   780: 138e0f10                 addi t3, t6, 0x100
   784: 070c0e02                 vle8.v v24, (t3)
   788: 138e0f12                 addi t3, t6, 0x120
   792: 070d0e02                 vle8.v v26, (t3)
   796: 5708c627                 vand.vv v16, v28, v12
   800: 5739c2a2                 vsrl.vi v18, v12, 4
   804: 570a8833                 vrgather.vv v20, v24, v16
   808: 570ba933                 vrgather.vv v22, v26, v18
   812: 570a6a2f                 vxor.vv v20, v22, v20
   816: 5704442f                 vxor.vv v8, v20, v8
   820: 5708c727                 vand.vv v16, v28, v14
   824: 5739e2a2                 vsrl.vi v18, v14, 4
   828: 570a8833                 vrgather.vv v20, v24, v16
   832: 570ba933                 vrgather.vv v22, v26, v18
   836: 570a6a2f                 vxor.vv v20, v22, v20
   840: 5705452f                 vxor.vv v10, v20, v10
   844: 83338502                 ld t2, 0x28(a0)
   848: b3836300                 add t2, t2, t1
   852: 07860302                 vle8.v v12, (t2)
   856: b3835300                 add t2, t2, t0
   860: 07870302                 vle8.v v14, (t2)
   864: 5700c02e                 vxor.vv v0, v12, v0
   868: 5701e12e                 vxor.vv v2, v14, v2
   872: 5702c22e                 vxor.vv v4, v12, v4
   876: 5703e32e                 vxor.vv v6, v14, v6
   880: 5704c42e                 vxor.vv v8, v12, v8
   884: 5705e52e                 vxor.vv v10, v14, v10
   888: 83b30500                 ld t2, 0(a1)
   892: b3836300                 add t2, t2, t1
   896: 27800302                 vse8.v v0, (t2)
   900: b3835300                 add t2, t2, t0
   904: 27810302                 vse8.v v2, (t2)
   908: 83b38500                 ld t2, 8(a1)
   912: b3836300                 add t2, t2, t1
   916: 27820302                 vse8.v v4, (t2)
   920: b3835300                 add t2, t2, t0
   924: 27830302                 vse8.v v6, (t2)
   928: 83b30501                 ld t2, 0x10(a1)
   932: b3836300                 add t2, t2, t1
   936: 27840302                 vse8.v v8, (t2)
   940: b3835300                 add t2, t2, t0
   944: 27850302                 vse8.v v10, (t2)
   948: 3303d300                 add t1, t1, a3
   952: 1307f7ff                 addi a4, a4, -1
   956: e31207c8                 bnez a4, 0x40
   960: 33066640                 sub a2, a2, t1
   964: d772160c                 vsetvli t0, a2, e8, m2, ta, ma
   968: 63940200                 bnez t0, 0x3d0
   972: 6f00402a                 j 0x670
   976: 83330500                 ld t2, 0(a0)
   980: b3836300                 add t2, t2, t1
   984: 07860302                 vle8.v v12, (t2)
   988: 5700065e                 vmv.v.v v0, v12
   992: 138f0200                 mv t5, t0
   996: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1000: 138e0f00                 mv t3, t6
  1004: 070c0e02                 vle8.v v24, (t3)
  1008: 138e0f02                 addi t3, t6, 0x20
  1012: 070d0e02                 vle8.v v26, (t3)
  1016: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1020: 5708c627                 vand.vv v16, v28, v12
  1024: 5739c2a2                 vsrl.vi v18, v12, 4
  1028: 570a8833                 vrgather.vv v20, v24, v16
  1032: 570ba933                 vrgather.vv v22, v26, v18
  1036: 570a6a2f                 vxor.vv v20, v22, v20
  1040: 57020a5e                 vmv.v.v v4, v20
  1044: 138f0200                 mv t5, t0
  1048: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1052: 138e0f04                 addi t3, t6, 0x40
  1056: 070c0e02                 vle8.v v24, (t3)
  1060: 138e0f06                 addi t3, t6, 0x60
  1064: 070d0e02                 vle8.v v26, (t3)
  1068: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1072: 5708c627                 vand.vv v16, v28, v12
  1076: 5739c2a2                 vsrl.vi v18, v12, 4
  1080: 570a8833                 vrgather.vv v20, v24, v16
  1084: 570ba933                 vrgather.vv v22, v26, v18
  1088: 570a6a2f                 vxor.vv v20, v22, v20
  1092: 57040a5e                 vmv.v.v v8, v20
  1096: 83338500                 ld t2, 8(a0)
  1100: b3836300                 add t2, t2, t1
  1104: 07860302                 vle8.v v12, (t2)
  1108: 5700c02e                 vxor.vv v0, v12, v0
  1112: 138f0200                 mv t5, t0
  1116: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1120: 138e0f04                 addi t3, t6, 0x40
  1124: 070c0e02                 vle8.v v24, (t3)
  1128: 138e0f06                 addi t3, t6, 0x60
  1132: 070d0e02                 vle8.v v26, (t3)
  1136: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1140: 5708c627                 vand.vv v16, v28, v12
  1144: 5739c2a2                 vsrl.vi v18, v12, 4
  1148: 570a8833                 vrgather.vv v20, v24, v16
  1152: 570ba933                 vrgather.vv v22, v26, v18
  1156: 570a6a2f                 vxor.vv v20, v22, v20
  1160: 5702422f                 vxor.vv v4, v20, v4
  1164: 138f0200                 mv t5, t0
  1168: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1172: 138e0f08                 addi t3, t6, 0x80
  1176: 070c0e02                 vle8.v v24, (t3)
  1180: 138e0f0a                 addi t3, t6, 0xa0
  1184: 070d0e02                 vle8.v v26, (t3)
  1188: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1192: 5708c627                 vand.vv v16, v28, v12
  1196: 5739c2a2                 vsrl.vi v18, v12, 4
  1200: 570a8833                 vrgather.vv v20, v24, v16
  1204: 570ba933                 vrgather.vv v22, v26, v18
  1208: 570a6a2f                 vxor.vv v20, v22, v20
  1212: 5704442f                 vxor.vv v8, v20, v8
  1216: 83330501                 ld t2, 0x10(a0)
  1220: b3836300                 add t2, t2, t1
  1224: 07860302                 vle8.v v12, (t2)
  1228: 5700c02e                 vxor.vv v0, v12, v0
  1232: 138f0200                 mv t5, t0
  1236: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1240: 138e0f08                 addi t3, t6, 0x80
  1244: 070c0e02                 vle8.v v24, (t3)
  1248: 138e0f0a                 addi t3, t6, 0xa0
  1252: 070d0e02                 vle8.v v26, (t3)
  1256: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1260: 5708c627                 vand.vv v16, v28, v12
  1264: 5739c2a2                 vsrl.vi v18, v12, 4
  1268: 570a8833                 vrgather.vv v20, v24, v16
  1272: 570ba933                 vrgather.vv v22, v26, v18
  1276: 570a6a2f                 vxor.vv v20, v22, v20
  1280: 5702422f                 vxor.vv v4, v20, v4
  1284: 138f0200                 mv t5, t0
  1288: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1292: 138e0f0c                 addi t3, t6, 0xc0
  1296: 070c0e02                 vle8.v v24, (t3)
  1300: 138e0f0e                 addi t3, t6, 0xe0
  1304: 070d0e02                 vle8.v v26, (t3)
  1308: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1312: 5708c627                 vand.vv v16, v28, v12
  1316: 5739c2a2                 vsrl.vi v18, v12, 4
  1320: 570a8833                 vrgather.vv v20, v24, v16
  1324: 570ba933                 vrgather.vv v22, v26, v18
  1328: 570a6a2f                 vxor.vv v20, v22, v20
  1332: 5704442f                 vxor.vv v8, v20, v8
  1336: 83338501                 ld t2, 0x18(a0)
  1340: b3836300                 add t2, t2, t1
  1344: 07860302                 vle8.v v12, (t2)
  1348: 5700c02e                 vxor.vv v0, v12, v0
  1352: 138f0200                 mv t5, t0
  1356: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1360: 138e0f10                 addi t3, t6, 0x100
  1364: 070c0e02                 vle8.v v24, (t3)
  1368: 138e0f12                 addi t3, t6, 0x120
  1372: 070d0e02                 vle8.v v26, (t3)
  1376: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1380: 5708c627                 vand.vv v16, v28, v12
  1384: 5739c2a2                 vsrl.vi v18, v12, 4
  1388: 570a8833                 vrgather.vv v20, v24, v16
  1392: 570ba933                 vrgather.vv v22, v26, v18
  1396: 570a6a2f                 vxor.vv v20, v22, v20
  1400: 5702422f                 vxor.vv v4, v20, v4
  1404: 138f0200                 mv t5, t0
  1408: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1412: 138e0f00                 mv t3, t6
  1416: 070c0e02                 vle8.v v24, (t3)
  1420: 138e0f02                 addi t3, t6, 0x20
  1424: 070d0e02                 vle8.v v26, (t3)
  1428: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1432: 5708c627                 vand.vv v16, v28, v12
  1436: 5739c2a2                 vsrl.vi v18, v12, 4
  1440: 570a8833                 vrgather.vv v20, v24, v16
  1444: 570ba933                 vrgather.vv v22, v26, v18
  1448: 570a6a2f                 vxor.vv v20, v22, v20
  1452: 5704442f                 vxor.vv v8, v20, v8
  1456: 83330502                 ld t2, 0x20(a0)
  1460: b3836300                 add t2, t2, t1
  1464: 07860302                 vle8.v v12, (t2)
  1468: 5700c02e                 vxor.vv v0, v12, v0
  1472: 138f0200                 mv t5, t0
  1476: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1480: 138e0f0c                 addi t3, t6, 0xc0
  1484: 070c0e02                 vle8.v v24, (t3)
  1488: 138e0f0e                 addi t3, t6, 0xe0
  1492: 070d0e02                 vle8.v v26, (t3)
  1496: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1500: 5708c627                 vand.vv v16, v28, v12
  1504: 5739c2a2                 vsrl.vi v18, v12, 4
  1508: 570a8833                 vrgather.vv v20, v24, v16
  1512: 570ba933                 vrgather.vv v22, v26, v18
  1516: 570a6a2f                 vxor.vv v20, v22, v20
  1520: 5702422f                 vxor.vv v4, v20, v4
  1524: 138f0200                 mv t5, t0
  1528: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1532: 138e0f10                 addi t3, t6, 0x100
  1536: 070c0e02                 vle8.v v24, (t3)
  1540: 138e0f12                 addi t3, t6, 0x120
  1544: 070d0e02                 vle8.v v26, (t3)
  1548: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1552: 5708c627                 vand.vv v16, v28, v12
  1556: 5739c2a2                 vsrl.vi v18, v12, 4
  1560: 570a8833                 vrgather.vv v20, v24, v16
  1564: 570ba933                 vrgather.vv v22, v26, v18
  1568: 570a6a2f                 vxor.vv v20, v22, v20
  1572: 5704442f                 vxor.vv v8, v20, v8
  1576: 83338502                 ld t2, 0x28(a0)
  1580: b3836300                 add t2, t2, t1
  1584: 07860302                 vle8.v v12, (t2)
  1588: 5700c02e                 vxor.vv v0, v12, v0
  1592: 5702c22e                 vxor.vv v4, v12, v4
  1596: 5704c42e                 vxor.vv v8, v12, v8
  1600: 83b30500                 ld t2, 0(a1)
  1604: b3836300                 add t2, t2, t1
  1608: 27800302                 vse8.v v0, (t2)
  1612: 83b38500                 ld t2, 8(a1)
  1616: b3836300                 add t2, t2, t1
  1620: 27820302                 vse8.v v4, (t2)
  1624: 83b30501                 ld t2, 0x10(a1)
  1628: b3836300                 add t2, t2, t1
  1632: 27840302                 vse8.v v8, (t2)
  1636: 33035300                 add t1, t1, t0
  1640: 33065640                 sub a2, a2, t0
  1644: 6ff09fd5                 j 0x3c4
  1648: 67800000                 ret 
  1652: 0000                     unimp 
  1654: 0000                     unimp 
  1656: 0000                     unimp 
  1658: 0000                     unimp 
  1660: 0000                     unimp 
  1662: 0000                     unimp 
