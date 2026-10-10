; config_t1_1612 target=rvv twin=baseline topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1712 twins=[] uops_per_byte=3.40625
     0: 63140600                 bnez a2, 8
     4: 6f008069                 j 0x69c
     8: 970f0000                 auipc t6, 0
    12: 938f8f6a                 addi t6, t6, 0x6a8
    16: 577f100c                 vsetvli t5, zero, e8, m2, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574f0e5e                 vmv.v.x v30, t3
    28: 13030000                 li t1, 0
    32: d772100c                 vsetvli t0, zero, e8, m2, ta, ma
    36: 93063000                 li a3, 3
    40: b386d202                 mul a3, t0, a3
    44: 3357d602                 divu a4, a2, a3
    48: 63140700                 bnez a4, 0x38
    52: 6f004043                 j 0x468
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07860302                 vle8.v v12, (t2)
    76: b3835300                 add t2, t2, t0
    80: 5700065e                 vmv.v.v v0, v12
    84: 07870302                 vle8.v v14, (t2)
    88: b3835300                 add t2, t2, t0
    92: 5701075e                 vmv.v.v v2, v14
    96: 07880302                 vle8.v v16, (t2)
   100: 5702085e                 vmv.v.v v4, v16
   104: 83338500                 ld t2, 8(a0)
   108: b3836300                 add t2, t2, t1
   112: 07860302                 vle8.v v12, (t2)
   116: b3835300                 add t2, t2, t0
   120: 5700c02e                 vxor.vv v0, v12, v0
   124: 07870302                 vle8.v v14, (t2)
   128: b3835300                 add t2, t2, t0
   132: 5701e12e                 vxor.vv v2, v14, v2
   136: 07880302                 vle8.v v16, (t2)
   140: 5702022f                 vxor.vv v4, v16, v4
   144: 83330501                 ld t2, 0x10(a0)
   148: b3836300                 add t2, t2, t1
   152: 07860302                 vle8.v v12, (t2)
   156: b3835300                 add t2, t2, t0
   160: 5700c02e                 vxor.vv v0, v12, v0
   164: 07870302                 vle8.v v14, (t2)
   168: b3835300                 add t2, t2, t0
   172: 5701e12e                 vxor.vv v2, v14, v2
   176: 07880302                 vle8.v v16, (t2)
   180: 5702022f                 vxor.vv v4, v16, v4
   184: 83338501                 ld t2, 0x18(a0)
   188: b3836300                 add t2, t2, t1
   192: 07860302                 vle8.v v12, (t2)
   196: b3835300                 add t2, t2, t0
   200: 5700c02e                 vxor.vv v0, v12, v0
   204: 07870302                 vle8.v v14, (t2)
   208: b3835300                 add t2, t2, t0
   212: 5701e12e                 vxor.vv v2, v14, v2
   216: 07880302                 vle8.v v16, (t2)
   220: 5702022f                 vxor.vv v4, v16, v4
   224: 83330502                 ld t2, 0x20(a0)
   228: b3836300                 add t2, t2, t1
   232: 07860302                 vle8.v v12, (t2)
   236: b3835300                 add t2, t2, t0
   240: 5700c02e                 vxor.vv v0, v12, v0
   244: 07870302                 vle8.v v14, (t2)
   248: b3835300                 add t2, t2, t0
   252: 5701e12e                 vxor.vv v2, v14, v2
   256: 07880302                 vle8.v v16, (t2)
   260: 5702022f                 vxor.vv v4, v16, v4
   264: 83330500                 ld t2, 0(a0)
   268: b3836300                 add t2, t2, t1
   272: 138e0f00                 mv t3, t6
   276: 070d0e02                 vle8.v v26, (t3)
   280: 138e0f02                 addi t3, t6, 0x20
   284: 070e0e02                 vle8.v v28, (t3)
   288: 07860302                 vle8.v v12, (t2)
   292: b3835300                 add t2, t2, t0
   296: 5709e627                 vand.vv v18, v30, v12
   300: 573ac2a2                 vsrl.vi v20, v12, 4
   304: 570ba933                 vrgather.vv v22, v26, v18
   308: 570cca33                 vrgather.vv v24, v28, v20
   312: 570b8b2f                 vxor.vv v22, v24, v22
   316: 57030b5e                 vmv.v.v v6, v22
   320: 07870302                 vle8.v v14, (t2)
   324: b3835300                 add t2, t2, t0
   328: 5709e727                 vand.vv v18, v30, v14
   332: 573ae2a2                 vsrl.vi v20, v14, 4
   336: 570ba933                 vrgather.vv v22, v26, v18
   340: 570cca33                 vrgather.vv v24, v28, v20
   344: 570b8b2f                 vxor.vv v22, v24, v22
   348: 57040b5e                 vmv.v.v v8, v22
   352: 07880302                 vle8.v v16, (t2)
   356: 5709e827                 vand.vv v18, v30, v16
   360: 573a02a3                 vsrl.vi v20, v16, 4
   364: 570ba933                 vrgather.vv v22, v26, v18
   368: 570cca33                 vrgather.vv v24, v28, v20
   372: 570b8b2f                 vxor.vv v22, v24, v22
   376: 57050b5e                 vmv.v.v v10, v22
   380: 83338500                 ld t2, 8(a0)
   384: b3836300                 add t2, t2, t1
   388: 138e0f04                 addi t3, t6, 0x40
   392: 070d0e02                 vle8.v v26, (t3)
   396: 138e0f06                 addi t3, t6, 0x60
   400: 070e0e02                 vle8.v v28, (t3)
   404: 07860302                 vle8.v v12, (t2)
   408: b3835300                 add t2, t2, t0
   412: 5709e627                 vand.vv v18, v30, v12
   416: 573ac2a2                 vsrl.vi v20, v12, 4
   420: 570ba933                 vrgather.vv v22, v26, v18
   424: 570cca33                 vrgather.vv v24, v28, v20
   428: 570b8b2f                 vxor.vv v22, v24, v22
   432: 5703632f                 vxor.vv v6, v22, v6
   436: 07870302                 vle8.v v14, (t2)
   440: b3835300                 add t2, t2, t0
   444: 5709e727                 vand.vv v18, v30, v14
   448: 573ae2a2                 vsrl.vi v20, v14, 4
   452: 570ba933                 vrgather.vv v22, v26, v18
   456: 570cca33                 vrgather.vv v24, v28, v20
   460: 570b8b2f                 vxor.vv v22, v24, v22
   464: 5704642f                 vxor.vv v8, v22, v8
   468: 07880302                 vle8.v v16, (t2)
   472: 5709e827                 vand.vv v18, v30, v16
   476: 573a02a3                 vsrl.vi v20, v16, 4
   480: 570ba933                 vrgather.vv v22, v26, v18
   484: 570cca33                 vrgather.vv v24, v28, v20
   488: 570b8b2f                 vxor.vv v22, v24, v22
   492: 5705652f                 vxor.vv v10, v22, v10
   496: 83330501                 ld t2, 0x10(a0)
   500: b3836300                 add t2, t2, t1
   504: 138e0f08                 addi t3, t6, 0x80
   508: 070d0e02                 vle8.v v26, (t3)
   512: 138e0f0a                 addi t3, t6, 0xa0
   516: 070e0e02                 vle8.v v28, (t3)
   520: 07860302                 vle8.v v12, (t2)
   524: b3835300                 add t2, t2, t0
   528: 5709e627                 vand.vv v18, v30, v12
   532: 573ac2a2                 vsrl.vi v20, v12, 4
   536: 570ba933                 vrgather.vv v22, v26, v18
   540: 570cca33                 vrgather.vv v24, v28, v20
   544: 570b8b2f                 vxor.vv v22, v24, v22
   548: 5703632f                 vxor.vv v6, v22, v6
   552: 07870302                 vle8.v v14, (t2)
   556: b3835300                 add t2, t2, t0
   560: 5709e727                 vand.vv v18, v30, v14
   564: 573ae2a2                 vsrl.vi v20, v14, 4
   568: 570ba933                 vrgather.vv v22, v26, v18
   572: 570cca33                 vrgather.vv v24, v28, v20
   576: 570b8b2f                 vxor.vv v22, v24, v22
   580: 5704642f                 vxor.vv v8, v22, v8
   584: 07880302                 vle8.v v16, (t2)
   588: 5709e827                 vand.vv v18, v30, v16
   592: 573a02a3                 vsrl.vi v20, v16, 4
   596: 570ba933                 vrgather.vv v22, v26, v18
   600: 570cca33                 vrgather.vv v24, v28, v20
   604: 570b8b2f                 vxor.vv v22, v24, v22
   608: 5705652f                 vxor.vv v10, v22, v10
   612: 83338501                 ld t2, 0x18(a0)
   616: b3836300                 add t2, t2, t1
   620: 07860302                 vle8.v v12, (t2)
   624: b3835300                 add t2, t2, t0
   628: 5703c32e                 vxor.vv v6, v12, v6
   632: 07870302                 vle8.v v14, (t2)
   636: b3835300                 add t2, t2, t0
   640: 5704e42e                 vxor.vv v8, v14, v8
   644: 07880302                 vle8.v v16, (t2)
   648: 5705052f                 vxor.vv v10, v16, v10
   652: 83338502                 ld t2, 0x28(a0)
   656: b3836300                 add t2, t2, t1
   660: 07860302                 vle8.v v12, (t2)
   664: b3835300                 add t2, t2, t0
   668: 5703c32e                 vxor.vv v6, v12, v6
   672: 07870302                 vle8.v v14, (t2)
   676: b3835300                 add t2, t2, t0
   680: 5704e42e                 vxor.vv v8, v14, v8
   684: 07880302                 vle8.v v16, (t2)
   688: 5705052f                 vxor.vv v10, v16, v10
   692: 138e0f0c                 addi t3, t6, 0xc0
   696: 070d0e02                 vle8.v v26, (t3)
   700: 138e0f0e                 addi t3, t6, 0xe0
   704: 070e0e02                 vle8.v v28, (t3)
   708: 5709e027                 vand.vv v18, v30, v0
   712: 573a02a2                 vsrl.vi v20, v0, 4
   716: 570ba933                 vrgather.vv v22, v26, v18
   720: 570cca33                 vrgather.vv v24, v28, v20
   724: 570b8b2f                 vxor.vv v22, v24, v22
   728: 57060b5e                 vmv.v.v v12, v22
   732: 5709e127                 vand.vv v18, v30, v2
   736: 573a22a2                 vsrl.vi v20, v2, 4
   740: 570ba933                 vrgather.vv v22, v26, v18
   744: 570cca33                 vrgather.vv v24, v28, v20
   748: 570b8b2f                 vxor.vv v22, v24, v22
   752: 57070b5e                 vmv.v.v v14, v22
   756: 5709e227                 vand.vv v18, v30, v4
   760: 573a42a2                 vsrl.vi v20, v4, 4
   764: 570ba933                 vrgather.vv v22, v26, v18
   768: 570cca33                 vrgather.vv v24, v28, v20
   772: 570b8b2f                 vxor.vv v22, v24, v22
   776: 57080b5e                 vmv.v.v v16, v22
   780: 5706662e                 vxor.vv v12, v6, v12
   784: 5707872e                 vxor.vv v14, v8, v14
   788: 5708a82e                 vxor.vv v16, v10, v16
   792: 138e0f10                 addi t3, t6, 0x100
   796: 070d0e02                 vle8.v v26, (t3)
   800: 138e0f12                 addi t3, t6, 0x120
   804: 070e0e02                 vle8.v v28, (t3)
   808: 5709e627                 vand.vv v18, v30, v12
   812: 573ac2a2                 vsrl.vi v20, v12, 4
   816: 570ba933                 vrgather.vv v22, v26, v18
   820: 570cca33                 vrgather.vv v24, v28, v20
   824: 570b8b2f                 vxor.vv v22, v24, v22
   828: 57060b5e                 vmv.v.v v12, v22
   832: 5709e727                 vand.vv v18, v30, v14
   836: 573ae2a2                 vsrl.vi v20, v14, 4
   840: 570ba933                 vrgather.vv v22, v26, v18
   844: 570cca33                 vrgather.vv v24, v28, v20
   848: 570b8b2f                 vxor.vv v22, v24, v22
   852: 57070b5e                 vmv.v.v v14, v22
   856: 5709e827                 vand.vv v18, v30, v16
   860: 573a02a3                 vsrl.vi v20, v16, 4
   864: 570ba933                 vrgather.vv v22, v26, v18
   868: 570cca33                 vrgather.vv v24, v28, v20
   872: 570b8b2f                 vxor.vv v22, v24, v22
   876: 57080b5e                 vmv.v.v v16, v22
   880: 138e0f14                 addi t3, t6, 0x140
   884: 070d0e02                 vle8.v v26, (t3)
   888: 138e0f16                 addi t3, t6, 0x160
   892: 070e0e02                 vle8.v v28, (t3)
   896: 5709e027                 vand.vv v18, v30, v0
   900: 573a02a2                 vsrl.vi v20, v0, 4
   904: 570ba933                 vrgather.vv v22, v26, v18
   908: 570cca33                 vrgather.vv v24, v28, v20
   912: 570b8b2f                 vxor.vv v22, v24, v22
   916: 5703632f                 vxor.vv v6, v22, v6
   920: 5709e127                 vand.vv v18, v30, v2
   924: 573a22a2                 vsrl.vi v20, v2, 4
   928: 570ba933                 vrgather.vv v22, v26, v18
   932: 570cca33                 vrgather.vv v24, v28, v20
   936: 570b8b2f                 vxor.vv v22, v24, v22
   940: 5704642f                 vxor.vv v8, v22, v8
   944: 5709e227                 vand.vv v18, v30, v4
   948: 573a42a2                 vsrl.vi v20, v4, 4
   952: 570ba933                 vrgather.vv v22, v26, v18
   956: 570cca33                 vrgather.vv v24, v28, v20
   960: 570b8b2f                 vxor.vv v22, v24, v22
   964: 5705652f                 vxor.vv v10, v22, v10
   968: 138e0f10                 addi t3, t6, 0x100
   972: 070d0e02                 vle8.v v26, (t3)
   976: 138e0f12                 addi t3, t6, 0x120
   980: 070e0e02                 vle8.v v28, (t3)
   984: 5709e327                 vand.vv v18, v30, v6
   988: 573a62a2                 vsrl.vi v20, v6, 4
   992: 570ba933                 vrgather.vv v22, v26, v18
   996: 570cca33                 vrgather.vv v24, v28, v20
  1000: 570b8b2f                 vxor.vv v22, v24, v22
  1004: 57030b5e                 vmv.v.v v6, v22
  1008: 5709e427                 vand.vv v18, v30, v8
  1012: 573a82a2                 vsrl.vi v20, v8, 4
  1016: 570ba933                 vrgather.vv v22, v26, v18
  1020: 570cca33                 vrgather.vv v24, v28, v20
  1024: 570b8b2f                 vxor.vv v22, v24, v22
  1028: 57040b5e                 vmv.v.v v8, v22
  1032: 5709e527                 vand.vv v18, v30, v10
  1036: 573aa2a2                 vsrl.vi v20, v10, 4
  1040: 570ba933                 vrgather.vv v22, v26, v18
  1044: 570cca33                 vrgather.vv v24, v28, v20
  1048: 570b8b2f                 vxor.vv v22, v24, v22
  1052: 57050b5e                 vmv.v.v v10, v22
  1056: 83b30500                 ld t2, 0(a1)
  1060: b3836300                 add t2, t2, t1
  1064: 27860302                 vse8.v v12, (t2)
  1068: b3835300                 add t2, t2, t0
  1072: 27870302                 vse8.v v14, (t2)
  1076: b3835300                 add t2, t2, t0
  1080: 27880302                 vse8.v v16, (t2)
  1084: 83b38500                 ld t2, 8(a1)
  1088: b3836300                 add t2, t2, t1
  1092: 27830302                 vse8.v v6, (t2)
  1096: b3835300                 add t2, t2, t0
  1100: 27840302                 vse8.v v8, (t2)
  1104: b3835300                 add t2, t2, t0
  1108: 27850302                 vse8.v v10, (t2)
  1112: 3303d300                 add t1, t1, a3
  1116: 1307f7ff                 addi a4, a4, -1
  1120: e31007be                 bnez a4, 0x40
  1124: 33066640                 sub a2, a2, t1
  1128: d772160c                 vsetvli t0, a2, e8, m2, ta, ma
  1132: 63940200                 bnez t0, 0x474
  1136: 6f00c022                 j 0x69c
  1140: 83330500                 ld t2, 0(a0)
  1144: b3836300                 add t2, t2, t1
  1148: 07860302                 vle8.v v12, (t2)
  1152: 5700065e                 vmv.v.v v0, v12
  1156: 83338500                 ld t2, 8(a0)
  1160: b3836300                 add t2, t2, t1
  1164: 07860302                 vle8.v v12, (t2)
  1168: 5700c02e                 vxor.vv v0, v12, v0
  1172: 83330501                 ld t2, 0x10(a0)
  1176: b3836300                 add t2, t2, t1
  1180: 07860302                 vle8.v v12, (t2)
  1184: 5700c02e                 vxor.vv v0, v12, v0
  1188: 83338501                 ld t2, 0x18(a0)
  1192: b3836300                 add t2, t2, t1
  1196: 07860302                 vle8.v v12, (t2)
  1200: 5700c02e                 vxor.vv v0, v12, v0
  1204: 83330502                 ld t2, 0x20(a0)
  1208: b3836300                 add t2, t2, t1
  1212: 07860302                 vle8.v v12, (t2)
  1216: 5700c02e                 vxor.vv v0, v12, v0
  1220: 83330500                 ld t2, 0(a0)
  1224: b3836300                 add t2, t2, t1
  1228: 138f0200                 mv t5, t0
  1232: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1236: 138e0f00                 mv t3, t6
  1240: 070d0e02                 vle8.v v26, (t3)
  1244: 138e0f02                 addi t3, t6, 0x20
  1248: 070e0e02                 vle8.v v28, (t3)
  1252: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1256: 07860302                 vle8.v v12, (t2)
  1260: 5709e627                 vand.vv v18, v30, v12
  1264: 573ac2a2                 vsrl.vi v20, v12, 4
  1268: 570ba933                 vrgather.vv v22, v26, v18
  1272: 570cca33                 vrgather.vv v24, v28, v20
  1276: 570b8b2f                 vxor.vv v22, v24, v22
  1280: 57030b5e                 vmv.v.v v6, v22
  1284: 83338500                 ld t2, 8(a0)
  1288: b3836300                 add t2, t2, t1
  1292: 138f0200                 mv t5, t0
  1296: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1300: 138e0f04                 addi t3, t6, 0x40
  1304: 070d0e02                 vle8.v v26, (t3)
  1308: 138e0f06                 addi t3, t6, 0x60
  1312: 070e0e02                 vle8.v v28, (t3)
  1316: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1320: 07860302                 vle8.v v12, (t2)
  1324: 5709e627                 vand.vv v18, v30, v12
  1328: 573ac2a2                 vsrl.vi v20, v12, 4
  1332: 570ba933                 vrgather.vv v22, v26, v18
  1336: 570cca33                 vrgather.vv v24, v28, v20
  1340: 570b8b2f                 vxor.vv v22, v24, v22
  1344: 5703632f                 vxor.vv v6, v22, v6
  1348: 83330501                 ld t2, 0x10(a0)
  1352: b3836300                 add t2, t2, t1
  1356: 138f0200                 mv t5, t0
  1360: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1364: 138e0f08                 addi t3, t6, 0x80
  1368: 070d0e02                 vle8.v v26, (t3)
  1372: 138e0f0a                 addi t3, t6, 0xa0
  1376: 070e0e02                 vle8.v v28, (t3)
  1380: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1384: 07860302                 vle8.v v12, (t2)
  1388: 5709e627                 vand.vv v18, v30, v12
  1392: 573ac2a2                 vsrl.vi v20, v12, 4
  1396: 570ba933                 vrgather.vv v22, v26, v18
  1400: 570cca33                 vrgather.vv v24, v28, v20
  1404: 570b8b2f                 vxor.vv v22, v24, v22
  1408: 5703632f                 vxor.vv v6, v22, v6
  1412: 83338501                 ld t2, 0x18(a0)
  1416: b3836300                 add t2, t2, t1
  1420: 07860302                 vle8.v v12, (t2)
  1424: 5703c32e                 vxor.vv v6, v12, v6
  1428: 83338502                 ld t2, 0x28(a0)
  1432: b3836300                 add t2, t2, t1
  1436: 07860302                 vle8.v v12, (t2)
  1440: 5703c32e                 vxor.vv v6, v12, v6
  1444: 138f0200                 mv t5, t0
  1448: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1452: 138e0f0c                 addi t3, t6, 0xc0
  1456: 070d0e02                 vle8.v v26, (t3)
  1460: 138e0f0e                 addi t3, t6, 0xe0
  1464: 070e0e02                 vle8.v v28, (t3)
  1468: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1472: 5709e027                 vand.vv v18, v30, v0
  1476: 573a02a2                 vsrl.vi v20, v0, 4
  1480: 570ba933                 vrgather.vv v22, v26, v18
  1484: 570cca33                 vrgather.vv v24, v28, v20
  1488: 570b8b2f                 vxor.vv v22, v24, v22
  1492: 57060b5e                 vmv.v.v v12, v22
  1496: 5706662e                 vxor.vv v12, v6, v12
  1500: 138f0200                 mv t5, t0
  1504: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1508: 138e0f10                 addi t3, t6, 0x100
  1512: 070d0e02                 vle8.v v26, (t3)
  1516: 138e0f12                 addi t3, t6, 0x120
  1520: 070e0e02                 vle8.v v28, (t3)
  1524: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1528: 5709e627                 vand.vv v18, v30, v12
  1532: 573ac2a2                 vsrl.vi v20, v12, 4
  1536: 570ba933                 vrgather.vv v22, v26, v18
  1540: 570cca33                 vrgather.vv v24, v28, v20
  1544: 570b8b2f                 vxor.vv v22, v24, v22
  1548: 57060b5e                 vmv.v.v v12, v22
  1552: 138f0200                 mv t5, t0
  1556: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1560: 138e0f14                 addi t3, t6, 0x140
  1564: 070d0e02                 vle8.v v26, (t3)
  1568: 138e0f16                 addi t3, t6, 0x160
  1572: 070e0e02                 vle8.v v28, (t3)
  1576: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1580: 5709e027                 vand.vv v18, v30, v0
  1584: 573a02a2                 vsrl.vi v20, v0, 4
  1588: 570ba933                 vrgather.vv v22, v26, v18
  1592: 570cca33                 vrgather.vv v24, v28, v20
  1596: 570b8b2f                 vxor.vv v22, v24, v22
  1600: 5703632f                 vxor.vv v6, v22, v6
  1604: 138f0200                 mv t5, t0
  1608: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
  1612: 138e0f10                 addi t3, t6, 0x100
  1616: 070d0e02                 vle8.v v26, (t3)
  1620: 138e0f12                 addi t3, t6, 0x120
  1624: 070e0e02                 vle8.v v28, (t3)
  1628: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
  1632: 5709e327                 vand.vv v18, v30, v6
  1636: 573a62a2                 vsrl.vi v20, v6, 4
  1640: 570ba933                 vrgather.vv v22, v26, v18
  1644: 570cca33                 vrgather.vv v24, v28, v20
  1648: 570b8b2f                 vxor.vv v22, v24, v22
  1652: 57030b5e                 vmv.v.v v6, v22
  1656: 83b30500                 ld t2, 0(a1)
  1660: b3836300                 add t2, t2, t1
  1664: 27860302                 vse8.v v12, (t2)
  1668: 83b38500                 ld t2, 8(a1)
  1672: b3836300                 add t2, t2, t1
  1676: 27830302                 vse8.v v6, (t2)
  1680: 33035300                 add t1, t1, t0
  1684: 33065640                 sub a2, a2, t0
  1688: 6ff01fdd                 j 0x468
  1692: 67800000                 ret 
  1696: 0000                     unimp 
  1698: 0000                     unimp 
  1700: 0000                     unimp 
  1702: 0000                     unimp 
  1704: 0000                     unimp 
  1706: 0000                     unimp 
  1708: 0000                     unimp 
  1710: 0000                     unimp 
