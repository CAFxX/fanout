; config_t1_1490 target=rvv twin=xorsched topo=rdp(7)/enc op=encode U=1
; code_size=2512 twins=["xorsched"] uops_per_byte=19.375
     0: 63140600                 bnez a2, 8
     4: 6f00501c                 j 0x9c8
     8: 971f0000                 auipc t6, 1
    12: 938f8f9c                 addi t6, t6, -0x638
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f00804d                 j 0x500
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 033e8500                 ld t3, 8(a0)
    68: 330e6e00                 add t3, t3, t1
    72: 833e8503                 ld t4, 0x38(a0)
    76: b38e6e00                 add t4, t4, t1
    80: 07060e02                 vle8.v v12, (t3)
    84: 87860e02                 vle8.v v13, (t4)
    88: 5706d62e                 vxor.vv v12, v13, v12
    92: d700065e                 vmv.v.v v1, v12
    96: 5703065e                 vmv.v.v v6, v12
   100: 033e0501                 ld t3, 0x10(a0)
   104: 330e6e00                 add t3, t3, t1
   108: 833e0504                 ld t4, 0x40(a0)
   112: b38e6e00                 add t4, t4, t1
   116: 07060e02                 vle8.v v12, (t3)
   120: 87860e02                 vle8.v v13, (t4)
   124: 5706d62e                 vxor.vv v12, v13, v12
   128: 5701065e                 vmv.v.v v2, v12
   132: d703065e                 vmv.v.v v7, v12
   136: 033e8501                 ld t3, 0x18(a0)
   140: 330e6e00                 add t3, t3, t1
   144: 833e8504                 ld t4, 0x48(a0)
   148: b38e6e00                 add t4, t4, t1
   152: 07060e02                 vle8.v v12, (t3)
   156: 87860e02                 vle8.v v13, (t4)
   160: 5706d62e                 vxor.vv v12, v13, v12
   164: d701065e                 vmv.v.v v3, v12
   168: 5704065e                 vmv.v.v v8, v12
   172: 033e0502                 ld t3, 0x20(a0)
   176: 330e6e00                 add t3, t3, t1
   180: 833e0505                 ld t4, 0x50(a0)
   184: b38e6e00                 add t4, t4, t1
   188: 07060e02                 vle8.v v12, (t3)
   192: 87860e02                 vle8.v v13, (t4)
   196: 5706d62e                 vxor.vv v12, v13, v12
   200: 5702065e                 vmv.v.v v4, v12
   204: d704065e                 vmv.v.v v9, v12
   208: 033e8502                 ld t3, 0x28(a0)
   212: 330e6e00                 add t3, t3, t1
   216: 833e8505                 ld t4, 0x58(a0)
   220: b38e6e00                 add t4, t4, t1
   224: 07060e02                 vle8.v v12, (t3)
   228: 87860e02                 vle8.v v13, (t4)
   232: 5706d62e                 vxor.vv v12, v13, v12
   236: d702065e                 vmv.v.v v5, v12
   240: 5705065e                 vmv.v.v v10, v12
   244: 033e8506                 ld t3, 0x68(a0)
   248: 330e6e00                 add t3, t3, t1
   252: 833e8509                 ld t4, 0x98(a0)
   256: b38e6e00                 add t4, t4, t1
   260: 07060e02                 vle8.v v12, (t3)
   264: 87860e02                 vle8.v v13, (t4)
   268: 5706d62e                 vxor.vv v12, v13, v12
   272: d780c02e                 vxor.vv v1, v12, v1
   276: 5703c32e                 vxor.vv v6, v12, v6
   280: 033e0507                 ld t3, 0x70(a0)
   284: 330e6e00                 add t3, t3, t1
   288: 833e050a                 ld t4, 0xa0(a0)
   292: b38e6e00                 add t4, t4, t1
   296: 07060e02                 vle8.v v12, (t3)
   300: 87860e02                 vle8.v v13, (t4)
   304: 5706d62e                 vxor.vv v12, v13, v12
   308: 5701c12e                 vxor.vv v2, v12, v2
   312: d783c32e                 vxor.vv v7, v12, v7
   316: 033e8507                 ld t3, 0x78(a0)
   320: 330e6e00                 add t3, t3, t1
   324: 833e850a                 ld t4, 0xa8(a0)
   328: b38e6e00                 add t4, t4, t1
   332: 07060e02                 vle8.v v12, (t3)
   336: 87860e02                 vle8.v v13, (t4)
   340: 5706d62e                 vxor.vv v12, v13, v12
   344: d781c12e                 vxor.vv v3, v12, v3
   348: 5704c42e                 vxor.vv v8, v12, v8
   352: 033e0508                 ld t3, 0x80(a0)
   356: 330e6e00                 add t3, t3, t1
   360: 833e050b                 ld t4, 0xb0(a0)
   364: b38e6e00                 add t4, t4, t1
   368: 07060e02                 vle8.v v12, (t3)
   372: 87860e02                 vle8.v v13, (t4)
   376: 5706d62e                 vxor.vv v12, v13, v12
   380: 5702c22e                 vxor.vv v4, v12, v4
   384: d784c42e                 vxor.vv v9, v12, v9
   388: 033e8508                 ld t3, 0x88(a0)
   392: 330e6e00                 add t3, t3, t1
   396: 833e850b                 ld t4, 0xb8(a0)
   400: b38e6e00                 add t4, t4, t1
   404: 07060e02                 vle8.v v12, (t3)
   408: 87860e02                 vle8.v v13, (t4)
   412: 5706d62e                 vxor.vv v12, v13, v12
   416: d782c22e                 vxor.vv v5, v12, v5
   420: 5705c52e                 vxor.vv v10, v12, v10
   424: 033e850c                 ld t3, 0xc8(a0)
   428: 330e6e00                 add t3, t3, t1
   432: 833e850f                 ld t4, 0xf8(a0)
   436: b38e6e00                 add t4, t4, t1
   440: 07060e02                 vle8.v v12, (t3)
   444: 87860e02                 vle8.v v13, (t4)
   448: 5706d62e                 vxor.vv v12, v13, v12
   452: d780c02e                 vxor.vv v1, v12, v1
   456: 5703c32e                 vxor.vv v6, v12, v6
   460: 033e050d                 ld t3, 0xd0(a0)
   464: 330e6e00                 add t3, t3, t1
   468: 833e0510                 ld t4, 0x100(a0)
   472: b38e6e00                 add t4, t4, t1
   476: 07060e02                 vle8.v v12, (t3)
   480: 87860e02                 vle8.v v13, (t4)
   484: 5706d62e                 vxor.vv v12, v13, v12
   488: 5701c12e                 vxor.vv v2, v12, v2
   492: d783c32e                 vxor.vv v7, v12, v7
   496: 033e850d                 ld t3, 0xd8(a0)
   500: 330e6e00                 add t3, t3, t1
   504: 833e8510                 ld t4, 0x108(a0)
   508: b38e6e00                 add t4, t4, t1
   512: 07060e02                 vle8.v v12, (t3)
   516: 87860e02                 vle8.v v13, (t4)
   520: 5706d62e                 vxor.vv v12, v13, v12
   524: d781c12e                 vxor.vv v3, v12, v3
   528: 5704c42e                 vxor.vv v8, v12, v8
   532: 033e050e                 ld t3, 0xe0(a0)
   536: 330e6e00                 add t3, t3, t1
   540: 833e0511                 ld t4, 0x110(a0)
   544: b38e6e00                 add t4, t4, t1
   548: 07060e02                 vle8.v v12, (t3)
   552: 87860e02                 vle8.v v13, (t4)
   556: 5706d62e                 vxor.vv v12, v13, v12
   560: 5702c22e                 vxor.vv v4, v12, v4
   564: d784c42e                 vxor.vv v9, v12, v9
   568: 033e850e                 ld t3, 0xe8(a0)
   572: 330e6e00                 add t3, t3, t1
   576: 833e8511                 ld t4, 0x118(a0)
   580: b38e6e00                 add t4, t4, t1
   584: 07060e02                 vle8.v v12, (t3)
   588: 87860e02                 vle8.v v13, (t4)
   592: 5706d62e                 vxor.vv v12, v13, v12
   596: d782c22e                 vxor.vv v5, v12, v5
   600: 5705c52e                 vxor.vv v10, v12, v10
   604: 83330500                 ld t2, 0(a0)
   608: b3836300                 add t2, t2, t1
   612: 07860302                 vle8.v v12, (t2)
   616: 5700065e                 vmv.v.v v0, v12
   620: 5703c32e                 vxor.vv v6, v12, v6
   624: 83338500                 ld t2, 8(a0)
   628: b3836300                 add t2, t2, t1
   632: 07860302                 vle8.v v12, (t2)
   636: d783c32e                 vxor.vv v7, v12, v7
   640: 83330501                 ld t2, 0x10(a0)
   644: b3836300                 add t2, t2, t1
   648: 07860302                 vle8.v v12, (t2)
   652: 5704c42e                 vxor.vv v8, v12, v8
   656: 83338501                 ld t2, 0x18(a0)
   660: b3836300                 add t2, t2, t1
   664: 07860302                 vle8.v v12, (t2)
   668: d784c42e                 vxor.vv v9, v12, v9
   672: 83330502                 ld t2, 0x20(a0)
   676: b3836300                 add t2, t2, t1
   680: 07860302                 vle8.v v12, (t2)
   684: 5705c52e                 vxor.vv v10, v12, v10
   688: 83338502                 ld t2, 0x28(a0)
   692: b3836300                 add t2, t2, t1
   696: 87850302                 vle8.v v11, (t2)
   700: 83330503                 ld t2, 0x30(a0)
   704: b3836300                 add t2, t2, t1
   708: 07860302                 vle8.v v12, (t2)
   712: 5700c02e                 vxor.vv v0, v12, v0
   716: d783c32e                 vxor.vv v7, v12, v7
   720: 83338503                 ld t2, 0x38(a0)
   724: b3836300                 add t2, t2, t1
   728: 07860302                 vle8.v v12, (t2)
   732: 5704c42e                 vxor.vv v8, v12, v8
   736: 83330504                 ld t2, 0x40(a0)
   740: b3836300                 add t2, t2, t1
   744: 07860302                 vle8.v v12, (t2)
   748: d784c42e                 vxor.vv v9, v12, v9
   752: 83338504                 ld t2, 0x48(a0)
   756: b3836300                 add t2, t2, t1
   760: 07860302                 vle8.v v12, (t2)
   764: 5705c52e                 vxor.vv v10, v12, v10
   768: 83330505                 ld t2, 0x50(a0)
   772: b3836300                 add t2, t2, t1
   776: 07860302                 vle8.v v12, (t2)
   780: d785c52e                 vxor.vv v11, v12, v11
   784: 83330506                 ld t2, 0x60(a0)
   788: b3836300                 add t2, t2, t1
   792: 07860302                 vle8.v v12, (t2)
   796: 5700c02e                 vxor.vv v0, v12, v0
   800: 5704c42e                 vxor.vv v8, v12, v8
   804: 83338506                 ld t2, 0x68(a0)
   808: b3836300                 add t2, t2, t1
   812: 07860302                 vle8.v v12, (t2)
   816: d784c42e                 vxor.vv v9, v12, v9
   820: 83330507                 ld t2, 0x70(a0)
   824: b3836300                 add t2, t2, t1
   828: 07860302                 vle8.v v12, (t2)
   832: 5705c52e                 vxor.vv v10, v12, v10
   836: 83338507                 ld t2, 0x78(a0)
   840: b3836300                 add t2, t2, t1
   844: 07860302                 vle8.v v12, (t2)
   848: d785c52e                 vxor.vv v11, v12, v11
   852: 83338508                 ld t2, 0x88(a0)
   856: b3836300                 add t2, t2, t1
   860: 07860302                 vle8.v v12, (t2)
   864: 5703c32e                 vxor.vv v6, v12, v6
   868: 83330509                 ld t2, 0x90(a0)
   872: b3836300                 add t2, t2, t1
   876: 07860302                 vle8.v v12, (t2)
   880: 5700c02e                 vxor.vv v0, v12, v0
   884: d784c42e                 vxor.vv v9, v12, v9
   888: 83338509                 ld t2, 0x98(a0)
   892: b3836300                 add t2, t2, t1
   896: 07860302                 vle8.v v12, (t2)
   900: 5705c52e                 vxor.vv v10, v12, v10
   904: 8333050a                 ld t2, 0xa0(a0)
   908: b3836300                 add t2, t2, t1
   912: 07860302                 vle8.v v12, (t2)
   916: d785c52e                 vxor.vv v11, v12, v11
   920: 8333050b                 ld t2, 0xb0(a0)
   924: b3836300                 add t2, t2, t1
   928: 07860302                 vle8.v v12, (t2)
   932: 5703c32e                 vxor.vv v6, v12, v6
   936: 8333850b                 ld t2, 0xb8(a0)
   940: b3836300                 add t2, t2, t1
   944: 07860302                 vle8.v v12, (t2)
   948: d783c32e                 vxor.vv v7, v12, v7
   952: 8333050c                 ld t2, 0xc0(a0)
   956: b3836300                 add t2, t2, t1
   960: 07860302                 vle8.v v12, (t2)
   964: 5700c02e                 vxor.vv v0, v12, v0
   968: 5705c52e                 vxor.vv v10, v12, v10
   972: 8333850c                 ld t2, 0xc8(a0)
   976: b3836300                 add t2, t2, t1
   980: 07860302                 vle8.v v12, (t2)
   984: d785c52e                 vxor.vv v11, v12, v11
   988: 8333850d                 ld t2, 0xd8(a0)
   992: b3836300                 add t2, t2, t1
   996: 07860302                 vle8.v v12, (t2)
  1000: 5703c32e                 vxor.vv v6, v12, v6
  1004: 8333050e                 ld t2, 0xe0(a0)
  1008: b3836300                 add t2, t2, t1
  1012: 07860302                 vle8.v v12, (t2)
  1016: d783c32e                 vxor.vv v7, v12, v7
  1020: 8333850e                 ld t2, 0xe8(a0)
  1024: b3836300                 add t2, t2, t1
  1028: 07860302                 vle8.v v12, (t2)
  1032: 5704c42e                 vxor.vv v8, v12, v8
  1036: 8333050f                 ld t2, 0xf0(a0)
  1040: b3836300                 add t2, t2, t1
  1044: 07860302                 vle8.v v12, (t2)
  1048: 5700c02e                 vxor.vv v0, v12, v0
  1052: d785c52e                 vxor.vv v11, v12, v11
  1056: 83330510                 ld t2, 0x100(a0)
  1060: b3836300                 add t2, t2, t1
  1064: 07860302                 vle8.v v12, (t2)
  1068: 5703c32e                 vxor.vv v6, v12, v6
  1072: 83338510                 ld t2, 0x108(a0)
  1076: b3836300                 add t2, t2, t1
  1080: 07860302                 vle8.v v12, (t2)
  1084: d783c32e                 vxor.vv v7, v12, v7
  1088: 83330511                 ld t2, 0x110(a0)
  1092: b3836300                 add t2, t2, t1
  1096: 07860302                 vle8.v v12, (t2)
  1100: 5704c42e                 vxor.vv v8, v12, v8
  1104: 83338511                 ld t2, 0x118(a0)
  1108: b3836300                 add t2, t2, t1
  1112: 07860302                 vle8.v v12, (t2)
  1116: d784c42e                 vxor.vv v9, v12, v9
  1120: 83b30500                 ld t2, 0(a1)
  1124: b3836300                 add t2, t2, t1
  1128: 27800302                 vse8.v v0, (t2)
  1132: 83b38500                 ld t2, 8(a1)
  1136: b3836300                 add t2, t2, t1
  1140: a7800302                 vse8.v v1, (t2)
  1144: 83b30501                 ld t2, 0x10(a1)
  1148: b3836300                 add t2, t2, t1
  1152: 27810302                 vse8.v v2, (t2)
  1156: 83b38501                 ld t2, 0x18(a1)
  1160: b3836300                 add t2, t2, t1
  1164: a7810302                 vse8.v v3, (t2)
  1168: 83b30502                 ld t2, 0x20(a1)
  1172: b3836300                 add t2, t2, t1
  1176: 27820302                 vse8.v v4, (t2)
  1180: 83b38502                 ld t2, 0x28(a1)
  1184: b3836300                 add t2, t2, t1
  1188: a7820302                 vse8.v v5, (t2)
  1192: 83b30503                 ld t2, 0x30(a1)
  1196: b3836300                 add t2, t2, t1
  1200: 27830302                 vse8.v v6, (t2)
  1204: 83b38503                 ld t2, 0x38(a1)
  1208: b3836300                 add t2, t2, t1
  1212: a7830302                 vse8.v v7, (t2)
  1216: 83b30504                 ld t2, 0x40(a1)
  1220: b3836300                 add t2, t2, t1
  1224: 27840302                 vse8.v v8, (t2)
  1228: 83b38504                 ld t2, 0x48(a1)
  1232: b3836300                 add t2, t2, t1
  1236: a7840302                 vse8.v v9, (t2)
  1240: 83b30505                 ld t2, 0x50(a1)
  1244: b3836300                 add t2, t2, t1
  1248: 27850302                 vse8.v v10, (t2)
  1252: 83b38505                 ld t2, 0x58(a1)
  1256: b3836300                 add t2, t2, t1
  1260: a7850302                 vse8.v v11, (t2)
  1264: 3303d300                 add t1, t1, a3
  1268: 1307f7ff                 addi a4, a4, -1
  1272: e31407b4                 bnez a4, 0x40
  1276: 33066640                 sub a2, a2, t1
  1280: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
  1284: 63940200                 bnez t0, 0x50c
  1288: 6f00004c                 j 0x9c8
  1292: 033e8500                 ld t3, 8(a0)
  1296: 330e6e00                 add t3, t3, t1
  1300: 833e8503                 ld t4, 0x38(a0)
  1304: b38e6e00                 add t4, t4, t1
  1308: 07060e02                 vle8.v v12, (t3)
  1312: 87860e02                 vle8.v v13, (t4)
  1316: 5706d62e                 vxor.vv v12, v13, v12
  1320: d700065e                 vmv.v.v v1, v12
  1324: 5703065e                 vmv.v.v v6, v12
  1328: 033e0501                 ld t3, 0x10(a0)
  1332: 330e6e00                 add t3, t3, t1
  1336: 833e0504                 ld t4, 0x40(a0)
  1340: b38e6e00                 add t4, t4, t1
  1344: 07060e02                 vle8.v v12, (t3)
  1348: 87860e02                 vle8.v v13, (t4)
  1352: 5706d62e                 vxor.vv v12, v13, v12
  1356: 5701065e                 vmv.v.v v2, v12
  1360: d703065e                 vmv.v.v v7, v12
  1364: 033e8501                 ld t3, 0x18(a0)
  1368: 330e6e00                 add t3, t3, t1
  1372: 833e8504                 ld t4, 0x48(a0)
  1376: b38e6e00                 add t4, t4, t1
  1380: 07060e02                 vle8.v v12, (t3)
  1384: 87860e02                 vle8.v v13, (t4)
  1388: 5706d62e                 vxor.vv v12, v13, v12
  1392: d701065e                 vmv.v.v v3, v12
  1396: 5704065e                 vmv.v.v v8, v12
  1400: 033e0502                 ld t3, 0x20(a0)
  1404: 330e6e00                 add t3, t3, t1
  1408: 833e0505                 ld t4, 0x50(a0)
  1412: b38e6e00                 add t4, t4, t1
  1416: 07060e02                 vle8.v v12, (t3)
  1420: 87860e02                 vle8.v v13, (t4)
  1424: 5706d62e                 vxor.vv v12, v13, v12
  1428: 5702065e                 vmv.v.v v4, v12
  1432: d704065e                 vmv.v.v v9, v12
  1436: 033e8502                 ld t3, 0x28(a0)
  1440: 330e6e00                 add t3, t3, t1
  1444: 833e8505                 ld t4, 0x58(a0)
  1448: b38e6e00                 add t4, t4, t1
  1452: 07060e02                 vle8.v v12, (t3)
  1456: 87860e02                 vle8.v v13, (t4)
  1460: 5706d62e                 vxor.vv v12, v13, v12
  1464: d702065e                 vmv.v.v v5, v12
  1468: 5705065e                 vmv.v.v v10, v12
  1472: 033e8506                 ld t3, 0x68(a0)
  1476: 330e6e00                 add t3, t3, t1
  1480: 833e8509                 ld t4, 0x98(a0)
  1484: b38e6e00                 add t4, t4, t1
  1488: 07060e02                 vle8.v v12, (t3)
  1492: 87860e02                 vle8.v v13, (t4)
  1496: 5706d62e                 vxor.vv v12, v13, v12
  1500: d780c02e                 vxor.vv v1, v12, v1
  1504: 5703c32e                 vxor.vv v6, v12, v6
  1508: 033e0507                 ld t3, 0x70(a0)
  1512: 330e6e00                 add t3, t3, t1
  1516: 833e050a                 ld t4, 0xa0(a0)
  1520: b38e6e00                 add t4, t4, t1
  1524: 07060e02                 vle8.v v12, (t3)
  1528: 87860e02                 vle8.v v13, (t4)
  1532: 5706d62e                 vxor.vv v12, v13, v12
  1536: 5701c12e                 vxor.vv v2, v12, v2
  1540: d783c32e                 vxor.vv v7, v12, v7
  1544: 033e8507                 ld t3, 0x78(a0)
  1548: 330e6e00                 add t3, t3, t1
  1552: 833e850a                 ld t4, 0xa8(a0)
  1556: b38e6e00                 add t4, t4, t1
  1560: 07060e02                 vle8.v v12, (t3)
  1564: 87860e02                 vle8.v v13, (t4)
  1568: 5706d62e                 vxor.vv v12, v13, v12
  1572: d781c12e                 vxor.vv v3, v12, v3
  1576: 5704c42e                 vxor.vv v8, v12, v8
  1580: 033e0508                 ld t3, 0x80(a0)
  1584: 330e6e00                 add t3, t3, t1
  1588: 833e050b                 ld t4, 0xb0(a0)
  1592: b38e6e00                 add t4, t4, t1
  1596: 07060e02                 vle8.v v12, (t3)
  1600: 87860e02                 vle8.v v13, (t4)
  1604: 5706d62e                 vxor.vv v12, v13, v12
  1608: 5702c22e                 vxor.vv v4, v12, v4
  1612: d784c42e                 vxor.vv v9, v12, v9
  1616: 033e8508                 ld t3, 0x88(a0)
  1620: 330e6e00                 add t3, t3, t1
  1624: 833e850b                 ld t4, 0xb8(a0)
  1628: b38e6e00                 add t4, t4, t1
  1632: 07060e02                 vle8.v v12, (t3)
  1636: 87860e02                 vle8.v v13, (t4)
  1640: 5706d62e                 vxor.vv v12, v13, v12
  1644: d782c22e                 vxor.vv v5, v12, v5
  1648: 5705c52e                 vxor.vv v10, v12, v10
  1652: 033e850c                 ld t3, 0xc8(a0)
  1656: 330e6e00                 add t3, t3, t1
  1660: 833e850f                 ld t4, 0xf8(a0)
  1664: b38e6e00                 add t4, t4, t1
  1668: 07060e02                 vle8.v v12, (t3)
  1672: 87860e02                 vle8.v v13, (t4)
  1676: 5706d62e                 vxor.vv v12, v13, v12
  1680: d780c02e                 vxor.vv v1, v12, v1
  1684: 5703c32e                 vxor.vv v6, v12, v6
  1688: 033e050d                 ld t3, 0xd0(a0)
  1692: 330e6e00                 add t3, t3, t1
  1696: 833e0510                 ld t4, 0x100(a0)
  1700: b38e6e00                 add t4, t4, t1
  1704: 07060e02                 vle8.v v12, (t3)
  1708: 87860e02                 vle8.v v13, (t4)
  1712: 5706d62e                 vxor.vv v12, v13, v12
  1716: 5701c12e                 vxor.vv v2, v12, v2
  1720: d783c32e                 vxor.vv v7, v12, v7
  1724: 033e850d                 ld t3, 0xd8(a0)
  1728: 330e6e00                 add t3, t3, t1
  1732: 833e8510                 ld t4, 0x108(a0)
  1736: b38e6e00                 add t4, t4, t1
  1740: 07060e02                 vle8.v v12, (t3)
  1744: 87860e02                 vle8.v v13, (t4)
  1748: 5706d62e                 vxor.vv v12, v13, v12
  1752: d781c12e                 vxor.vv v3, v12, v3
  1756: 5704c42e                 vxor.vv v8, v12, v8
  1760: 033e050e                 ld t3, 0xe0(a0)
  1764: 330e6e00                 add t3, t3, t1
  1768: 833e0511                 ld t4, 0x110(a0)
  1772: b38e6e00                 add t4, t4, t1
  1776: 07060e02                 vle8.v v12, (t3)
  1780: 87860e02                 vle8.v v13, (t4)
  1784: 5706d62e                 vxor.vv v12, v13, v12
  1788: 5702c22e                 vxor.vv v4, v12, v4
  1792: d784c42e                 vxor.vv v9, v12, v9
  1796: 033e850e                 ld t3, 0xe8(a0)
  1800: 330e6e00                 add t3, t3, t1
  1804: 833e8511                 ld t4, 0x118(a0)
  1808: b38e6e00                 add t4, t4, t1
  1812: 07060e02                 vle8.v v12, (t3)
  1816: 87860e02                 vle8.v v13, (t4)
  1820: 5706d62e                 vxor.vv v12, v13, v12
  1824: d782c22e                 vxor.vv v5, v12, v5
  1828: 5705c52e                 vxor.vv v10, v12, v10
  1832: 83330500                 ld t2, 0(a0)
  1836: b3836300                 add t2, t2, t1
  1840: 07860302                 vle8.v v12, (t2)
  1844: 5700065e                 vmv.v.v v0, v12
  1848: 5703c32e                 vxor.vv v6, v12, v6
  1852: 83338500                 ld t2, 8(a0)
  1856: b3836300                 add t2, t2, t1
  1860: 07860302                 vle8.v v12, (t2)
  1864: d783c32e                 vxor.vv v7, v12, v7
  1868: 83330501                 ld t2, 0x10(a0)
  1872: b3836300                 add t2, t2, t1
  1876: 07860302                 vle8.v v12, (t2)
  1880: 5704c42e                 vxor.vv v8, v12, v8
  1884: 83338501                 ld t2, 0x18(a0)
  1888: b3836300                 add t2, t2, t1
  1892: 07860302                 vle8.v v12, (t2)
  1896: d784c42e                 vxor.vv v9, v12, v9
  1900: 83330502                 ld t2, 0x20(a0)
  1904: b3836300                 add t2, t2, t1
  1908: 07860302                 vle8.v v12, (t2)
  1912: 5705c52e                 vxor.vv v10, v12, v10
  1916: 83338502                 ld t2, 0x28(a0)
  1920: b3836300                 add t2, t2, t1
  1924: 87850302                 vle8.v v11, (t2)
  1928: 83330503                 ld t2, 0x30(a0)
  1932: b3836300                 add t2, t2, t1
  1936: 07860302                 vle8.v v12, (t2)
  1940: 5700c02e                 vxor.vv v0, v12, v0
  1944: d783c32e                 vxor.vv v7, v12, v7
  1948: 83338503                 ld t2, 0x38(a0)
  1952: b3836300                 add t2, t2, t1
  1956: 07860302                 vle8.v v12, (t2)
  1960: 5704c42e                 vxor.vv v8, v12, v8
  1964: 83330504                 ld t2, 0x40(a0)
  1968: b3836300                 add t2, t2, t1
  1972: 07860302                 vle8.v v12, (t2)
  1976: d784c42e                 vxor.vv v9, v12, v9
  1980: 83338504                 ld t2, 0x48(a0)
  1984: b3836300                 add t2, t2, t1
  1988: 07860302                 vle8.v v12, (t2)
  1992: 5705c52e                 vxor.vv v10, v12, v10
  1996: 83330505                 ld t2, 0x50(a0)
  2000: b3836300                 add t2, t2, t1
  2004: 07860302                 vle8.v v12, (t2)
  2008: d785c52e                 vxor.vv v11, v12, v11
  2012: 83330506                 ld t2, 0x60(a0)
  2016: b3836300                 add t2, t2, t1
  2020: 07860302                 vle8.v v12, (t2)
  2024: 5700c02e                 vxor.vv v0, v12, v0
  2028: 5704c42e                 vxor.vv v8, v12, v8
  2032: 83338506                 ld t2, 0x68(a0)
  2036: b3836300                 add t2, t2, t1
  2040: 07860302                 vle8.v v12, (t2)
  2044: d784c42e                 vxor.vv v9, v12, v9
  2048: 83330507                 ld t2, 0x70(a0)
  2052: b3836300                 add t2, t2, t1
  2056: 07860302                 vle8.v v12, (t2)
  2060: 5705c52e                 vxor.vv v10, v12, v10
  2064: 83338507                 ld t2, 0x78(a0)
  2068: b3836300                 add t2, t2, t1
  2072: 07860302                 vle8.v v12, (t2)
  2076: d785c52e                 vxor.vv v11, v12, v11
  2080: 83338508                 ld t2, 0x88(a0)
  2084: b3836300                 add t2, t2, t1
  2088: 07860302                 vle8.v v12, (t2)
  2092: 5703c32e                 vxor.vv v6, v12, v6
  2096: 83330509                 ld t2, 0x90(a0)
  2100: b3836300                 add t2, t2, t1
  2104: 07860302                 vle8.v v12, (t2)
  2108: 5700c02e                 vxor.vv v0, v12, v0
  2112: d784c42e                 vxor.vv v9, v12, v9
  2116: 83338509                 ld t2, 0x98(a0)
  2120: b3836300                 add t2, t2, t1
  2124: 07860302                 vle8.v v12, (t2)
  2128: 5705c52e                 vxor.vv v10, v12, v10
  2132: 8333050a                 ld t2, 0xa0(a0)
  2136: b3836300                 add t2, t2, t1
  2140: 07860302                 vle8.v v12, (t2)
  2144: d785c52e                 vxor.vv v11, v12, v11
  2148: 8333050b                 ld t2, 0xb0(a0)
  2152: b3836300                 add t2, t2, t1
  2156: 07860302                 vle8.v v12, (t2)
  2160: 5703c32e                 vxor.vv v6, v12, v6
  2164: 8333850b                 ld t2, 0xb8(a0)
  2168: b3836300                 add t2, t2, t1
  2172: 07860302                 vle8.v v12, (t2)
  2176: d783c32e                 vxor.vv v7, v12, v7
  2180: 8333050c                 ld t2, 0xc0(a0)
  2184: b3836300                 add t2, t2, t1
  2188: 07860302                 vle8.v v12, (t2)
  2192: 5700c02e                 vxor.vv v0, v12, v0
  2196: 5705c52e                 vxor.vv v10, v12, v10
  2200: 8333850c                 ld t2, 0xc8(a0)
  2204: b3836300                 add t2, t2, t1
  2208: 07860302                 vle8.v v12, (t2)
  2212: d785c52e                 vxor.vv v11, v12, v11
  2216: 8333850d                 ld t2, 0xd8(a0)
  2220: b3836300                 add t2, t2, t1
  2224: 07860302                 vle8.v v12, (t2)
  2228: 5703c32e                 vxor.vv v6, v12, v6
  2232: 8333050e                 ld t2, 0xe0(a0)
  2236: b3836300                 add t2, t2, t1
  2240: 07860302                 vle8.v v12, (t2)
  2244: d783c32e                 vxor.vv v7, v12, v7
  2248: 8333850e                 ld t2, 0xe8(a0)
  2252: b3836300                 add t2, t2, t1
  2256: 07860302                 vle8.v v12, (t2)
  2260: 5704c42e                 vxor.vv v8, v12, v8
  2264: 8333050f                 ld t2, 0xf0(a0)
  2268: b3836300                 add t2, t2, t1
  2272: 07860302                 vle8.v v12, (t2)
  2276: 5700c02e                 vxor.vv v0, v12, v0
  2280: d785c52e                 vxor.vv v11, v12, v11
  2284: 83330510                 ld t2, 0x100(a0)
  2288: b3836300                 add t2, t2, t1
  2292: 07860302                 vle8.v v12, (t2)
  2296: 5703c32e                 vxor.vv v6, v12, v6
  2300: 83338510                 ld t2, 0x108(a0)
  2304: b3836300                 add t2, t2, t1
  2308: 07860302                 vle8.v v12, (t2)
  2312: d783c32e                 vxor.vv v7, v12, v7
  2316: 83330511                 ld t2, 0x110(a0)
  2320: b3836300                 add t2, t2, t1
  2324: 07860302                 vle8.v v12, (t2)
  2328: 5704c42e                 vxor.vv v8, v12, v8
  2332: 83338511                 ld t2, 0x118(a0)
  2336: b3836300                 add t2, t2, t1
  2340: 07860302                 vle8.v v12, (t2)
  2344: d784c42e                 vxor.vv v9, v12, v9
  2348: 83b30500                 ld t2, 0(a1)
  2352: b3836300                 add t2, t2, t1
  2356: 27800302                 vse8.v v0, (t2)
  2360: 83b38500                 ld t2, 8(a1)
  2364: b3836300                 add t2, t2, t1
  2368: a7800302                 vse8.v v1, (t2)
  2372: 83b30501                 ld t2, 0x10(a1)
  2376: b3836300                 add t2, t2, t1
  2380: 27810302                 vse8.v v2, (t2)
  2384: 83b38501                 ld t2, 0x18(a1)
  2388: b3836300                 add t2, t2, t1
  2392: a7810302                 vse8.v v3, (t2)
  2396: 83b30502                 ld t2, 0x20(a1)
  2400: b3836300                 add t2, t2, t1
  2404: 27820302                 vse8.v v4, (t2)
  2408: 83b38502                 ld t2, 0x28(a1)
  2412: b3836300                 add t2, t2, t1
  2416: a7820302                 vse8.v v5, (t2)
  2420: 83b30503                 ld t2, 0x30(a1)
  2424: b3836300                 add t2, t2, t1
  2428: 27830302                 vse8.v v6, (t2)
  2432: 83b38503                 ld t2, 0x38(a1)
  2436: b3836300                 add t2, t2, t1
  2440: a7830302                 vse8.v v7, (t2)
  2444: 83b30504                 ld t2, 0x40(a1)
  2448: b3836300                 add t2, t2, t1
  2452: 27840302                 vse8.v v8, (t2)
  2456: 83b38504                 ld t2, 0x48(a1)
  2460: b3836300                 add t2, t2, t1
  2464: a7840302                 vse8.v v9, (t2)
  2468: 83b30505                 ld t2, 0x50(a1)
  2472: b3836300                 add t2, t2, t1
  2476: 27850302                 vse8.v v10, (t2)
  2480: 83b38505                 ld t2, 0x58(a1)
  2484: b3836300                 add t2, t2, t1
  2488: a7850302                 vse8.v v11, (t2)
  2492: 33035300                 add t1, t1, t0
  2496: 33065640                 sub a2, a2, t0
  2500: 6ff0dfb3                 j 0x500
  2504: 67800000                 ret 
  2508: 0000                     unimp 
  2510: 0000                     unimp 
