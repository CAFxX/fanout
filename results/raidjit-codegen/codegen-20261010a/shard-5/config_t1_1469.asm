; config_t1_1469 target=rvv twin=baseline topo=raid6(8)/enc op=encode U=8
; code_size=1648 twins=[] uops_per_byte=3.3645833333333335
     0: 63140600                 bnez a2, 8
     4: 6f008065                 j 0x65c
     8: 970f0000                 auipc t6, 0
    12: 938f8f66                 addi t6, t6, 0x668
    16: 577f100c                 vsetvli t5, zero, e8, m2, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574f0e5e                 vmv.v.x v30, t3
    28: 13030000                 li t1, 0
    32: d772100c                 vsetvli t0, zero, e8, m2, ta, ma
    36: 93063000                 li a3, 3
    40: b386d202                 mul a3, t0, a3
    44: 3357d602                 divu a4, a2, a3
    48: 63140700                 bnez a4, 0x38
    52: 6f008040                 j 0x43c
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07860302                 vle8.v v12, (t2)
    76: b3835300                 add t2, t2, t0
    80: 07870302                 vle8.v v14, (t2)
    84: b3835300                 add t2, t2, t0
    88: 07880302                 vle8.v v16, (t2)
    92: 5700065e                 vmv.v.v v0, v12
    96: 5701075e                 vmv.v.v v2, v14
   100: 5702085e                 vmv.v.v v4, v16
   104: 138e0f00                 mv t3, t6
   108: 070d0e02                 vle8.v v26, (t3)
   112: 138e0f02                 addi t3, t6, 0x20
   116: 070e0e02                 vle8.v v28, (t3)
   120: 5709e627                 vand.vv v18, v30, v12
   124: 573ac2a2                 vsrl.vi v20, v12, 4
   128: 570ba933                 vrgather.vv v22, v26, v18
   132: 570cca33                 vrgather.vv v24, v28, v20
   136: 570b8b2f                 vxor.vv v22, v24, v22
   140: 57030b5e                 vmv.v.v v6, v22
   144: 5709e727                 vand.vv v18, v30, v14
   148: 573ae2a2                 vsrl.vi v20, v14, 4
   152: 570ba933                 vrgather.vv v22, v26, v18
   156: 570cca33                 vrgather.vv v24, v28, v20
   160: 570b8b2f                 vxor.vv v22, v24, v22
   164: 57040b5e                 vmv.v.v v8, v22
   168: 5709e827                 vand.vv v18, v30, v16
   172: 573a02a3                 vsrl.vi v20, v16, 4
   176: 570ba933                 vrgather.vv v22, v26, v18
   180: 570cca33                 vrgather.vv v24, v28, v20
   184: 570b8b2f                 vxor.vv v22, v24, v22
   188: 57050b5e                 vmv.v.v v10, v22
   192: 83338500                 ld t2, 8(a0)
   196: b3836300                 add t2, t2, t1
   200: 07860302                 vle8.v v12, (t2)
   204: b3835300                 add t2, t2, t0
   208: 07870302                 vle8.v v14, (t2)
   212: b3835300                 add t2, t2, t0
   216: 07880302                 vle8.v v16, (t2)
   220: 5700c02e                 vxor.vv v0, v12, v0
   224: 5701e12e                 vxor.vv v2, v14, v2
   228: 5702022f                 vxor.vv v4, v16, v4
   232: 138e0f04                 addi t3, t6, 0x40
   236: 070d0e02                 vle8.v v26, (t3)
   240: 138e0f06                 addi t3, t6, 0x60
   244: 070e0e02                 vle8.v v28, (t3)
   248: 5709e627                 vand.vv v18, v30, v12
   252: 573ac2a2                 vsrl.vi v20, v12, 4
   256: 570ba933                 vrgather.vv v22, v26, v18
   260: 570cca33                 vrgather.vv v24, v28, v20
   264: 570b8b2f                 vxor.vv v22, v24, v22
   268: 5703632f                 vxor.vv v6, v22, v6
   272: 5709e727                 vand.vv v18, v30, v14
   276: 573ae2a2                 vsrl.vi v20, v14, 4
   280: 570ba933                 vrgather.vv v22, v26, v18
   284: 570cca33                 vrgather.vv v24, v28, v20
   288: 570b8b2f                 vxor.vv v22, v24, v22
   292: 5704642f                 vxor.vv v8, v22, v8
   296: 5709e827                 vand.vv v18, v30, v16
   300: 573a02a3                 vsrl.vi v20, v16, 4
   304: 570ba933                 vrgather.vv v22, v26, v18
   308: 570cca33                 vrgather.vv v24, v28, v20
   312: 570b8b2f                 vxor.vv v22, v24, v22
   316: 5705652f                 vxor.vv v10, v22, v10
   320: 83330501                 ld t2, 0x10(a0)
   324: b3836300                 add t2, t2, t1
   328: 07860302                 vle8.v v12, (t2)
   332: b3835300                 add t2, t2, t0
   336: 07870302                 vle8.v v14, (t2)
   340: b3835300                 add t2, t2, t0
   344: 07880302                 vle8.v v16, (t2)
   348: 5700c02e                 vxor.vv v0, v12, v0
   352: 5701e12e                 vxor.vv v2, v14, v2
   356: 5702022f                 vxor.vv v4, v16, v4
   360: 138e0f08                 addi t3, t6, 0x80
   364: 070d0e02                 vle8.v v26, (t3)
   368: 138e0f0a                 addi t3, t6, 0xa0
   372: 070e0e02                 vle8.v v28, (t3)
   376: 5709e627                 vand.vv v18, v30, v12
   380: 573ac2a2                 vsrl.vi v20, v12, 4
   384: 570ba933                 vrgather.vv v22, v26, v18
   388: 570cca33                 vrgather.vv v24, v28, v20
   392: 570b8b2f                 vxor.vv v22, v24, v22
   396: 5703632f                 vxor.vv v6, v22, v6
   400: 5709e727                 vand.vv v18, v30, v14
   404: 573ae2a2                 vsrl.vi v20, v14, 4
   408: 570ba933                 vrgather.vv v22, v26, v18
   412: 570cca33                 vrgather.vv v24, v28, v20
   416: 570b8b2f                 vxor.vv v22, v24, v22
   420: 5704642f                 vxor.vv v8, v22, v8
   424: 5709e827                 vand.vv v18, v30, v16
   428: 573a02a3                 vsrl.vi v20, v16, 4
   432: 570ba933                 vrgather.vv v22, v26, v18
   436: 570cca33                 vrgather.vv v24, v28, v20
   440: 570b8b2f                 vxor.vv v22, v24, v22
   444: 5705652f                 vxor.vv v10, v22, v10
   448: 83338501                 ld t2, 0x18(a0)
   452: b3836300                 add t2, t2, t1
   456: 07860302                 vle8.v v12, (t2)
   460: b3835300                 add t2, t2, t0
   464: 07870302                 vle8.v v14, (t2)
   468: b3835300                 add t2, t2, t0
   472: 07880302                 vle8.v v16, (t2)
   476: 5700c02e                 vxor.vv v0, v12, v0
   480: 5701e12e                 vxor.vv v2, v14, v2
   484: 5702022f                 vxor.vv v4, v16, v4
   488: 138e0f0c                 addi t3, t6, 0xc0
   492: 070d0e02                 vle8.v v26, (t3)
   496: 138e0f0e                 addi t3, t6, 0xe0
   500: 070e0e02                 vle8.v v28, (t3)
   504: 5709e627                 vand.vv v18, v30, v12
   508: 573ac2a2                 vsrl.vi v20, v12, 4
   512: 570ba933                 vrgather.vv v22, v26, v18
   516: 570cca33                 vrgather.vv v24, v28, v20
   520: 570b8b2f                 vxor.vv v22, v24, v22
   524: 5703632f                 vxor.vv v6, v22, v6
   528: 5709e727                 vand.vv v18, v30, v14
   532: 573ae2a2                 vsrl.vi v20, v14, 4
   536: 570ba933                 vrgather.vv v22, v26, v18
   540: 570cca33                 vrgather.vv v24, v28, v20
   544: 570b8b2f                 vxor.vv v22, v24, v22
   548: 5704642f                 vxor.vv v8, v22, v8
   552: 5709e827                 vand.vv v18, v30, v16
   556: 573a02a3                 vsrl.vi v20, v16, 4
   560: 570ba933                 vrgather.vv v22, v26, v18
   564: 570cca33                 vrgather.vv v24, v28, v20
   568: 570b8b2f                 vxor.vv v22, v24, v22
   572: 5705652f                 vxor.vv v10, v22, v10
   576: 83330502                 ld t2, 0x20(a0)
   580: b3836300                 add t2, t2, t1
   584: 07860302                 vle8.v v12, (t2)
   588: b3835300                 add t2, t2, t0
   592: 07870302                 vle8.v v14, (t2)
   596: b3835300                 add t2, t2, t0
   600: 07880302                 vle8.v v16, (t2)
   604: 5700c02e                 vxor.vv v0, v12, v0
   608: 5701e12e                 vxor.vv v2, v14, v2
   612: 5702022f                 vxor.vv v4, v16, v4
   616: 138e0f10                 addi t3, t6, 0x100
   620: 070d0e02                 vle8.v v26, (t3)
   624: 138e0f12                 addi t3, t6, 0x120
   628: 070e0e02                 vle8.v v28, (t3)
   632: 5709e627                 vand.vv v18, v30, v12
   636: 573ac2a2                 vsrl.vi v20, v12, 4
   640: 570ba933                 vrgather.vv v22, v26, v18
   644: 570cca33                 vrgather.vv v24, v28, v20
   648: 570b8b2f                 vxor.vv v22, v24, v22
   652: 5703632f                 vxor.vv v6, v22, v6
   656: 5709e727                 vand.vv v18, v30, v14
   660: 573ae2a2                 vsrl.vi v20, v14, 4
   664: 570ba933                 vrgather.vv v22, v26, v18
   668: 570cca33                 vrgather.vv v24, v28, v20
   672: 570b8b2f                 vxor.vv v22, v24, v22
   676: 5704642f                 vxor.vv v8, v22, v8
   680: 5709e827                 vand.vv v18, v30, v16
   684: 573a02a3                 vsrl.vi v20, v16, 4
   688: 570ba933                 vrgather.vv v22, v26, v18
   692: 570cca33                 vrgather.vv v24, v28, v20
   696: 570b8b2f                 vxor.vv v22, v24, v22
   700: 5705652f                 vxor.vv v10, v22, v10
   704: 83338502                 ld t2, 0x28(a0)
   708: b3836300                 add t2, t2, t1
   712: 07860302                 vle8.v v12, (t2)
   716: b3835300                 add t2, t2, t0
   720: 07870302                 vle8.v v14, (t2)
   724: b3835300                 add t2, t2, t0
   728: 07880302                 vle8.v v16, (t2)
   732: 5700c02e                 vxor.vv v0, v12, v0
   736: 5701e12e                 vxor.vv v2, v14, v2
   740: 5702022f                 vxor.vv v4, v16, v4
   744: 138e0f14                 addi t3, t6, 0x140
   748: 070d0e02                 vle8.v v26, (t3)
   752: 138e0f16                 addi t3, t6, 0x160
   756: 070e0e02                 vle8.v v28, (t3)
   760: 5709e627                 vand.vv v18, v30, v12
   764: 573ac2a2                 vsrl.vi v20, v12, 4
   768: 570ba933                 vrgather.vv v22, v26, v18
   772: 570cca33                 vrgather.vv v24, v28, v20
   776: 570b8b2f                 vxor.vv v22, v24, v22
   780: 5703632f                 vxor.vv v6, v22, v6
   784: 5709e727                 vand.vv v18, v30, v14
   788: 573ae2a2                 vsrl.vi v20, v14, 4
   792: 570ba933                 vrgather.vv v22, v26, v18
   796: 570cca33                 vrgather.vv v24, v28, v20
   800: 570b8b2f                 vxor.vv v22, v24, v22
   804: 5704642f                 vxor.vv v8, v22, v8
   808: 5709e827                 vand.vv v18, v30, v16
   812: 573a02a3                 vsrl.vi v20, v16, 4
   816: 570ba933                 vrgather.vv v22, v26, v18
   820: 570cca33                 vrgather.vv v24, v28, v20
   824: 570b8b2f                 vxor.vv v22, v24, v22
   828: 5705652f                 vxor.vv v10, v22, v10
   832: 83330503                 ld t2, 0x30(a0)
   836: b3836300                 add t2, t2, t1
   840: 07860302                 vle8.v v12, (t2)
   844: b3835300                 add t2, t2, t0
   848: 07870302                 vle8.v v14, (t2)
   852: b3835300                 add t2, t2, t0
   856: 07880302                 vle8.v v16, (t2)
   860: 5700c02e                 vxor.vv v0, v12, v0
   864: 5701e12e                 vxor.vv v2, v14, v2
   868: 5702022f                 vxor.vv v4, v16, v4
   872: 138e0f18                 addi t3, t6, 0x180
   876: 070d0e02                 vle8.v v26, (t3)
   880: 138e0f1a                 addi t3, t6, 0x1a0
   884: 070e0e02                 vle8.v v28, (t3)
   888: 5709e627                 vand.vv v18, v30, v12
   892: 573ac2a2                 vsrl.vi v20, v12, 4
   896: 570ba933                 vrgather.vv v22, v26, v18
   900: 570cca33                 vrgather.vv v24, v28, v20
   904: 570b8b2f                 vxor.vv v22, v24, v22
   908: 5703632f                 vxor.vv v6, v22, v6
   912: 5709e727                 vand.vv v18, v30, v14
   916: 573ae2a2                 vsrl.vi v20, v14, 4
   920: 570ba933                 vrgather.vv v22, v26, v18
   924: 570cca33                 vrgather.vv v24, v28, v20
   928: 570b8b2f                 vxor.vv v22, v24, v22
   932: 5704642f                 vxor.vv v8, v22, v8
   936: 5709e827                 vand.vv v18, v30, v16
   940: 573a02a3                 vsrl.vi v20, v16, 4
   944: 570ba933                 vrgather.vv v22, v26, v18
   948: 570cca33                 vrgather.vv v24, v28, v20
   952: 570b8b2f                 vxor.vv v22, v24, v22
   956: 5705652f                 vxor.vv v10, v22, v10
   960: 83338503                 ld t2, 0x38(a0)
   964: b3836300                 add t2, t2, t1
   968: 07860302                 vle8.v v12, (t2)
   972: b3835300                 add t2, t2, t0
   976: 07870302                 vle8.v v14, (t2)
   980: b3835300                 add t2, t2, t0
   984: 07880302                 vle8.v v16, (t2)
   988: 5700c02e                 vxor.vv v0, v12, v0
   992: 5701e12e                 vxor.vv v2, v14, v2
   996: 5702022f                 vxor.vv v4, v16, v4
  1000: 5703c32e                 vxor.vv v6, v12, v6
  1004: 5704e42e                 vxor.vv v8, v14, v8
  1008: 5705052f                 vxor.vv v10, v16, v10
  1012: 83b30500                 ld t2, 0(a1)
  1016: b3836300                 add t2, t2, t1
  1020: 27800302                 vse8.v v0, (t2)
  1024: b3835300                 add t2, t2, t0
  1028: 27810302                 vse8.v v2, (t2)
  1032: b3835300                 add t2, t2, t0
  1036: 27820302                 vse8.v v4, (t2)
  1040: 83b38500                 ld t2, 8(a1)
  1044: b3836300                 add t2, t2, t1
  1048: 27830302                 vse8.v v6, (t2)
  1052: b3835300                 add t2, t2, t0
  1056: 27840302                 vse8.v v8, (t2)
  1060: b3835300                 add t2, t2, t0
  1064: 27850302                 vse8.v v10, (t2)
  1068: 3303d300                 add t1, t1, a3
  1072: 1307f7ff                 addi a4, a4, -1
  1076: e31607c0                 bnez a4, 0x40
  1080: 33066640                 sub a2, a2, t1
  1084: d772160c                 vsetvli t0, a2, e8, m2, ta, ma
  1088: 63940200                 bnez t0, 0x448
  1092: 6f008021                 j 0x65c
  1096: 83330500                 ld t2, 0(a0)
  1100: b3836300                 add t2, t2, t1
  1104: 07860302                 vle8.v v12, (t2)
  1108: 5700065e                 vmv.v.v v0, v12
  1112: 138f0200                 mv t5, t0
  1116: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1120: 138e0f00                 mv t3, t6
  1124: 070d0e02                 vle8.v v26, (t3)
  1128: 138e0f02                 addi t3, t6, 0x20
  1132: 070e0e02                 vle8.v v28, (t3)
  1136: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1140: 5709e627                 vand.vv v18, v30, v12
  1144: 573ac2a2                 vsrl.vi v20, v12, 4
  1148: 570ba933                 vrgather.vv v22, v26, v18
  1152: 570cca33                 vrgather.vv v24, v28, v20
  1156: 570b8b2f                 vxor.vv v22, v24, v22
  1160: 57030b5e                 vmv.v.v v6, v22
  1164: 83338500                 ld t2, 8(a0)
  1168: b3836300                 add t2, t2, t1
  1172: 07860302                 vle8.v v12, (t2)
  1176: 5700c02e                 vxor.vv v0, v12, v0
  1180: 138f0200                 mv t5, t0
  1184: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1188: 138e0f04                 addi t3, t6, 0x40
  1192: 070d0e02                 vle8.v v26, (t3)
  1196: 138e0f06                 addi t3, t6, 0x60
  1200: 070e0e02                 vle8.v v28, (t3)
  1204: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1208: 5709e627                 vand.vv v18, v30, v12
  1212: 573ac2a2                 vsrl.vi v20, v12, 4
  1216: 570ba933                 vrgather.vv v22, v26, v18
  1220: 570cca33                 vrgather.vv v24, v28, v20
  1224: 570b8b2f                 vxor.vv v22, v24, v22
  1228: 5703632f                 vxor.vv v6, v22, v6
  1232: 83330501                 ld t2, 0x10(a0)
  1236: b3836300                 add t2, t2, t1
  1240: 07860302                 vle8.v v12, (t2)
  1244: 5700c02e                 vxor.vv v0, v12, v0
  1248: 138f0200                 mv t5, t0
  1252: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1256: 138e0f08                 addi t3, t6, 0x80
  1260: 070d0e02                 vle8.v v26, (t3)
  1264: 138e0f0a                 addi t3, t6, 0xa0
  1268: 070e0e02                 vle8.v v28, (t3)
  1272: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1276: 5709e627                 vand.vv v18, v30, v12
  1280: 573ac2a2                 vsrl.vi v20, v12, 4
  1284: 570ba933                 vrgather.vv v22, v26, v18
  1288: 570cca33                 vrgather.vv v24, v28, v20
  1292: 570b8b2f                 vxor.vv v22, v24, v22
  1296: 5703632f                 vxor.vv v6, v22, v6
  1300: 83338501                 ld t2, 0x18(a0)
  1304: b3836300                 add t2, t2, t1
  1308: 07860302                 vle8.v v12, (t2)
  1312: 5700c02e                 vxor.vv v0, v12, v0
  1316: 138f0200                 mv t5, t0
  1320: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1324: 138e0f0c                 addi t3, t6, 0xc0
  1328: 070d0e02                 vle8.v v26, (t3)
  1332: 138e0f0e                 addi t3, t6, 0xe0
  1336: 070e0e02                 vle8.v v28, (t3)
  1340: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1344: 5709e627                 vand.vv v18, v30, v12
  1348: 573ac2a2                 vsrl.vi v20, v12, 4
  1352: 570ba933                 vrgather.vv v22, v26, v18
  1356: 570cca33                 vrgather.vv v24, v28, v20
  1360: 570b8b2f                 vxor.vv v22, v24, v22
  1364: 5703632f                 vxor.vv v6, v22, v6
  1368: 83330502                 ld t2, 0x20(a0)
  1372: b3836300                 add t2, t2, t1
  1376: 07860302                 vle8.v v12, (t2)
  1380: 5700c02e                 vxor.vv v0, v12, v0
  1384: 138f0200                 mv t5, t0
  1388: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1392: 138e0f10                 addi t3, t6, 0x100
  1396: 070d0e02                 vle8.v v26, (t3)
  1400: 138e0f12                 addi t3, t6, 0x120
  1404: 070e0e02                 vle8.v v28, (t3)
  1408: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1412: 5709e627                 vand.vv v18, v30, v12
  1416: 573ac2a2                 vsrl.vi v20, v12, 4
  1420: 570ba933                 vrgather.vv v22, v26, v18
  1424: 570cca33                 vrgather.vv v24, v28, v20
  1428: 570b8b2f                 vxor.vv v22, v24, v22
  1432: 5703632f                 vxor.vv v6, v22, v6
  1436: 83338502                 ld t2, 0x28(a0)
  1440: b3836300                 add t2, t2, t1
  1444: 07860302                 vle8.v v12, (t2)
  1448: 5700c02e                 vxor.vv v0, v12, v0
  1452: 138f0200                 mv t5, t0
  1456: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1460: 138e0f14                 addi t3, t6, 0x140
  1464: 070d0e02                 vle8.v v26, (t3)
  1468: 138e0f16                 addi t3, t6, 0x160
  1472: 070e0e02                 vle8.v v28, (t3)
  1476: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1480: 5709e627                 vand.vv v18, v30, v12
  1484: 573ac2a2                 vsrl.vi v20, v12, 4
  1488: 570ba933                 vrgather.vv v22, v26, v18
  1492: 570cca33                 vrgather.vv v24, v28, v20
  1496: 570b8b2f                 vxor.vv v22, v24, v22
  1500: 5703632f                 vxor.vv v6, v22, v6
  1504: 83330503                 ld t2, 0x30(a0)
  1508: b3836300                 add t2, t2, t1
  1512: 07860302                 vle8.v v12, (t2)
  1516: 5700c02e                 vxor.vv v0, v12, v0
  1520: 138f0200                 mv t5, t0
  1524: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1528: 138e0f18                 addi t3, t6, 0x180
  1532: 070d0e02                 vle8.v v26, (t3)
  1536: 138e0f1a                 addi t3, t6, 0x1a0
  1540: 070e0e02                 vle8.v v28, (t3)
  1544: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1548: 5709e627                 vand.vv v18, v30, v12
  1552: 573ac2a2                 vsrl.vi v20, v12, 4
  1556: 570ba933                 vrgather.vv v22, v26, v18
  1560: 570cca33                 vrgather.vv v24, v28, v20
  1564: 570b8b2f                 vxor.vv v22, v24, v22
  1568: 5703632f                 vxor.vv v6, v22, v6
  1572: 83338503                 ld t2, 0x38(a0)
  1576: b3836300                 add t2, t2, t1
  1580: 07860302                 vle8.v v12, (t2)
  1584: 5700c02e                 vxor.vv v0, v12, v0
  1588: 5703c32e                 vxor.vv v6, v12, v6
  1592: 83b30500                 ld t2, 0(a1)
  1596: b3836300                 add t2, t2, t1
  1600: 27800302                 vse8.v v0, (t2)
  1604: 83b38500                 ld t2, 8(a1)
  1608: b3836300                 add t2, t2, t1
  1612: 27830302                 vse8.v v6, (t2)
  1616: 33035300                 add t1, t1, t0
  1620: 33065640                 sub a2, a2, t0
  1624: 6ff05fde                 j 0x43c
  1628: 67800000                 ret 
  1632: 0000                     unimp 
  1634: 0000                     unimp 
  1636: 0000                     unimp 
  1638: 0000                     unimp 
  1640: 0000                     unimp 
  1642: 0000                     unimp 
  1644: 0000                     unimp 
  1646: 0000                     unimp 
