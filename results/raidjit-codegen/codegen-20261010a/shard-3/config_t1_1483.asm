; config_t1_1483 target=rvv twin=nt topo=rdp(7)/enc op=encode U=1
; code_size=2144 twins=["nt"] uops_per_byte=18.25
     0: 63140600                 bnez a2, 8
     4: 6f00d004                 j 0x850
     8: 971f0000                 auipc t6, 1
    12: 938f8f85                 addi t6, t6, -0x7a8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f00c041                 j 0x444
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07860302                 vle8.v v12, (t2)
    76: 5700065e                 vmv.v.v v0, v12
    80: 5703065e                 vmv.v.v v6, v12
    84: 83338500                 ld t2, 8(a0)
    88: b3836300                 add t2, t2, t1
    92: 07860302                 vle8.v v12, (t2)
    96: d700065e                 vmv.v.v v1, v12
   100: 5703c32e                 vxor.vv v6, v12, v6
   104: d703065e                 vmv.v.v v7, v12
   108: 83330501                 ld t2, 0x10(a0)
   112: b3836300                 add t2, t2, t1
   116: 07860302                 vle8.v v12, (t2)
   120: 5701065e                 vmv.v.v v2, v12
   124: d783c32e                 vxor.vv v7, v12, v7
   128: 5704065e                 vmv.v.v v8, v12
   132: 83338501                 ld t2, 0x18(a0)
   136: b3836300                 add t2, t2, t1
   140: 07860302                 vle8.v v12, (t2)
   144: d701065e                 vmv.v.v v3, v12
   148: 5704c42e                 vxor.vv v8, v12, v8
   152: d704065e                 vmv.v.v v9, v12
   156: 83330502                 ld t2, 0x20(a0)
   160: b3836300                 add t2, t2, t1
   164: 07860302                 vle8.v v12, (t2)
   168: 5702065e                 vmv.v.v v4, v12
   172: d784c42e                 vxor.vv v9, v12, v9
   176: 5705065e                 vmv.v.v v10, v12
   180: 83338502                 ld t2, 0x28(a0)
   184: b3836300                 add t2, t2, t1
   188: 07860302                 vle8.v v12, (t2)
   192: d702065e                 vmv.v.v v5, v12
   196: 5705c52e                 vxor.vv v10, v12, v10
   200: d705065e                 vmv.v.v v11, v12
   204: 83330503                 ld t2, 0x30(a0)
   208: b3836300                 add t2, t2, t1
   212: 07860302                 vle8.v v12, (t2)
   216: 5700c02e                 vxor.vv v0, v12, v0
   220: d783c32e                 vxor.vv v7, v12, v7
   224: 83338503                 ld t2, 0x38(a0)
   228: b3836300                 add t2, t2, t1
   232: 07860302                 vle8.v v12, (t2)
   236: d780c02e                 vxor.vv v1, v12, v1
   240: 5703c32e                 vxor.vv v6, v12, v6
   244: 5704c42e                 vxor.vv v8, v12, v8
   248: 83330504                 ld t2, 0x40(a0)
   252: b3836300                 add t2, t2, t1
   256: 07860302                 vle8.v v12, (t2)
   260: 5701c12e                 vxor.vv v2, v12, v2
   264: d783c32e                 vxor.vv v7, v12, v7
   268: d784c42e                 vxor.vv v9, v12, v9
   272: 83338504                 ld t2, 0x48(a0)
   276: b3836300                 add t2, t2, t1
   280: 07860302                 vle8.v v12, (t2)
   284: d781c12e                 vxor.vv v3, v12, v3
   288: 5704c42e                 vxor.vv v8, v12, v8
   292: 5705c52e                 vxor.vv v10, v12, v10
   296: 83330505                 ld t2, 0x50(a0)
   300: b3836300                 add t2, t2, t1
   304: 07860302                 vle8.v v12, (t2)
   308: 5702c22e                 vxor.vv v4, v12, v4
   312: d784c42e                 vxor.vv v9, v12, v9
   316: d785c52e                 vxor.vv v11, v12, v11
   320: 83338505                 ld t2, 0x58(a0)
   324: b3836300                 add t2, t2, t1
   328: 07860302                 vle8.v v12, (t2)
   332: d782c22e                 vxor.vv v5, v12, v5
   336: 5705c52e                 vxor.vv v10, v12, v10
   340: 83330506                 ld t2, 0x60(a0)
   344: b3836300                 add t2, t2, t1
   348: 07860302                 vle8.v v12, (t2)
   352: 5700c02e                 vxor.vv v0, v12, v0
   356: 5704c42e                 vxor.vv v8, v12, v8
   360: 83338506                 ld t2, 0x68(a0)
   364: b3836300                 add t2, t2, t1
   368: 07860302                 vle8.v v12, (t2)
   372: d780c02e                 vxor.vv v1, v12, v1
   376: 5703c32e                 vxor.vv v6, v12, v6
   380: d784c42e                 vxor.vv v9, v12, v9
   384: 83330507                 ld t2, 0x70(a0)
   388: b3836300                 add t2, t2, t1
   392: 07860302                 vle8.v v12, (t2)
   396: 5701c12e                 vxor.vv v2, v12, v2
   400: d783c32e                 vxor.vv v7, v12, v7
   404: 5705c52e                 vxor.vv v10, v12, v10
   408: 83338507                 ld t2, 0x78(a0)
   412: b3836300                 add t2, t2, t1
   416: 07860302                 vle8.v v12, (t2)
   420: d781c12e                 vxor.vv v3, v12, v3
   424: 5704c42e                 vxor.vv v8, v12, v8
   428: d785c52e                 vxor.vv v11, v12, v11
   432: 83330508                 ld t2, 0x80(a0)
   436: b3836300                 add t2, t2, t1
   440: 07860302                 vle8.v v12, (t2)
   444: 5702c22e                 vxor.vv v4, v12, v4
   448: d784c42e                 vxor.vv v9, v12, v9
   452: 83338508                 ld t2, 0x88(a0)
   456: b3836300                 add t2, t2, t1
   460: 07860302                 vle8.v v12, (t2)
   464: d782c22e                 vxor.vv v5, v12, v5
   468: 5703c32e                 vxor.vv v6, v12, v6
   472: 5705c52e                 vxor.vv v10, v12, v10
   476: 83330509                 ld t2, 0x90(a0)
   480: b3836300                 add t2, t2, t1
   484: 07860302                 vle8.v v12, (t2)
   488: 5700c02e                 vxor.vv v0, v12, v0
   492: d784c42e                 vxor.vv v9, v12, v9
   496: 83338509                 ld t2, 0x98(a0)
   500: b3836300                 add t2, t2, t1
   504: 07860302                 vle8.v v12, (t2)
   508: d780c02e                 vxor.vv v1, v12, v1
   512: 5703c32e                 vxor.vv v6, v12, v6
   516: 5705c52e                 vxor.vv v10, v12, v10
   520: 8333050a                 ld t2, 0xa0(a0)
   524: b3836300                 add t2, t2, t1
   528: 07860302                 vle8.v v12, (t2)
   532: 5701c12e                 vxor.vv v2, v12, v2
   536: d783c32e                 vxor.vv v7, v12, v7
   540: d785c52e                 vxor.vv v11, v12, v11
   544: 8333850a                 ld t2, 0xa8(a0)
   548: b3836300                 add t2, t2, t1
   552: 07860302                 vle8.v v12, (t2)
   556: d781c12e                 vxor.vv v3, v12, v3
   560: 5704c42e                 vxor.vv v8, v12, v8
   564: 8333050b                 ld t2, 0xb0(a0)
   568: b3836300                 add t2, t2, t1
   572: 07860302                 vle8.v v12, (t2)
   576: 5702c22e                 vxor.vv v4, v12, v4
   580: 5703c32e                 vxor.vv v6, v12, v6
   584: d784c42e                 vxor.vv v9, v12, v9
   588: 8333850b                 ld t2, 0xb8(a0)
   592: b3836300                 add t2, t2, t1
   596: 07860302                 vle8.v v12, (t2)
   600: d782c22e                 vxor.vv v5, v12, v5
   604: d783c32e                 vxor.vv v7, v12, v7
   608: 5705c52e                 vxor.vv v10, v12, v10
   612: 8333050c                 ld t2, 0xc0(a0)
   616: b3836300                 add t2, t2, t1
   620: 07860302                 vle8.v v12, (t2)
   624: 5700c02e                 vxor.vv v0, v12, v0
   628: 5705c52e                 vxor.vv v10, v12, v10
   632: 8333850c                 ld t2, 0xc8(a0)
   636: b3836300                 add t2, t2, t1
   640: 07860302                 vle8.v v12, (t2)
   644: d780c02e                 vxor.vv v1, v12, v1
   648: 5703c32e                 vxor.vv v6, v12, v6
   652: d785c52e                 vxor.vv v11, v12, v11
   656: 8333050d                 ld t2, 0xd0(a0)
   660: b3836300                 add t2, t2, t1
   664: 07860302                 vle8.v v12, (t2)
   668: 5701c12e                 vxor.vv v2, v12, v2
   672: d783c32e                 vxor.vv v7, v12, v7
   676: 8333850d                 ld t2, 0xd8(a0)
   680: b3836300                 add t2, t2, t1
   684: 07860302                 vle8.v v12, (t2)
   688: d781c12e                 vxor.vv v3, v12, v3
   692: 5703c32e                 vxor.vv v6, v12, v6
   696: 5704c42e                 vxor.vv v8, v12, v8
   700: 8333050e                 ld t2, 0xe0(a0)
   704: b3836300                 add t2, t2, t1
   708: 07860302                 vle8.v v12, (t2)
   712: 5702c22e                 vxor.vv v4, v12, v4
   716: d783c32e                 vxor.vv v7, v12, v7
   720: d784c42e                 vxor.vv v9, v12, v9
   724: 8333850e                 ld t2, 0xe8(a0)
   728: b3836300                 add t2, t2, t1
   732: 07860302                 vle8.v v12, (t2)
   736: d782c22e                 vxor.vv v5, v12, v5
   740: 5704c42e                 vxor.vv v8, v12, v8
   744: 5705c52e                 vxor.vv v10, v12, v10
   748: 8333050f                 ld t2, 0xf0(a0)
   752: b3836300                 add t2, t2, t1
   756: 07860302                 vle8.v v12, (t2)
   760: 5700c02e                 vxor.vv v0, v12, v0
   764: d785c52e                 vxor.vv v11, v12, v11
   768: 8333850f                 ld t2, 0xf8(a0)
   772: b3836300                 add t2, t2, t1
   776: 07860302                 vle8.v v12, (t2)
   780: d780c02e                 vxor.vv v1, v12, v1
   784: 5703c32e                 vxor.vv v6, v12, v6
   788: 83330510                 ld t2, 0x100(a0)
   792: b3836300                 add t2, t2, t1
   796: 07860302                 vle8.v v12, (t2)
   800: 5701c12e                 vxor.vv v2, v12, v2
   804: 5703c32e                 vxor.vv v6, v12, v6
   808: d783c32e                 vxor.vv v7, v12, v7
   812: 83338510                 ld t2, 0x108(a0)
   816: b3836300                 add t2, t2, t1
   820: 07860302                 vle8.v v12, (t2)
   824: d781c12e                 vxor.vv v3, v12, v3
   828: d783c32e                 vxor.vv v7, v12, v7
   832: 5704c42e                 vxor.vv v8, v12, v8
   836: 83330511                 ld t2, 0x110(a0)
   840: b3836300                 add t2, t2, t1
   844: 07860302                 vle8.v v12, (t2)
   848: 5702c22e                 vxor.vv v4, v12, v4
   852: 5704c42e                 vxor.vv v8, v12, v8
   856: d784c42e                 vxor.vv v9, v12, v9
   860: 83338511                 ld t2, 0x118(a0)
   864: b3836300                 add t2, t2, t1
   868: 07860302                 vle8.v v12, (t2)
   872: d782c22e                 vxor.vv v5, v12, v5
   876: d784c42e                 vxor.vv v9, v12, v9
   880: 5705c52e                 vxor.vv v10, v12, v10
   884: 83b30500                 ld t2, 0(a1)
   888: b3836300                 add t2, t2, t1
   892: 33005000                 ntl.all 
   896: 27800302                 vse8.v v0, (t2)
   900: 83b38500                 ld t2, 8(a1)
   904: b3836300                 add t2, t2, t1
   908: 33005000                 ntl.all 
   912: a7800302                 vse8.v v1, (t2)
   916: 83b30501                 ld t2, 0x10(a1)
   920: b3836300                 add t2, t2, t1
   924: 33005000                 ntl.all 
   928: 27810302                 vse8.v v2, (t2)
   932: 83b38501                 ld t2, 0x18(a1)
   936: b3836300                 add t2, t2, t1
   940: 33005000                 ntl.all 
   944: a7810302                 vse8.v v3, (t2)
   948: 83b30502                 ld t2, 0x20(a1)
   952: b3836300                 add t2, t2, t1
   956: 33005000                 ntl.all 
   960: 27820302                 vse8.v v4, (t2)
   964: 83b38502                 ld t2, 0x28(a1)
   968: b3836300                 add t2, t2, t1
   972: 33005000                 ntl.all 
   976: a7820302                 vse8.v v5, (t2)
   980: 83b30503                 ld t2, 0x30(a1)
   984: b3836300                 add t2, t2, t1
   988: 33005000                 ntl.all 
   992: 27830302                 vse8.v v6, (t2)
   996: 83b38503                 ld t2, 0x38(a1)
  1000: b3836300                 add t2, t2, t1
  1004: 33005000                 ntl.all 
  1008: a7830302                 vse8.v v7, (t2)
  1012: 83b30504                 ld t2, 0x40(a1)
  1016: b3836300                 add t2, t2, t1
  1020: 33005000                 ntl.all 
  1024: 27840302                 vse8.v v8, (t2)
  1028: 83b38504                 ld t2, 0x48(a1)
  1032: b3836300                 add t2, t2, t1
  1036: 33005000                 ntl.all 
  1040: a7840302                 vse8.v v9, (t2)
  1044: 83b30505                 ld t2, 0x50(a1)
  1048: b3836300                 add t2, t2, t1
  1052: 33005000                 ntl.all 
  1056: 27850302                 vse8.v v10, (t2)
  1060: 83b38505                 ld t2, 0x58(a1)
  1064: b3836300                 add t2, t2, t1
  1068: 33005000                 ntl.all 
  1072: a7850302                 vse8.v v11, (t2)
  1076: 3303d300                 add t1, t1, a3
  1080: 1307f7ff                 addi a4, a4, -1
  1084: e31207c0                 bnez a4, 0x40
  1088: 33066640                 sub a2, a2, t1
  1092: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
  1096: 63940200                 bnez t0, 0x450
  1100: 6f004040                 j 0x850
  1104: 83330500                 ld t2, 0(a0)
  1108: b3836300                 add t2, t2, t1
  1112: 07860302                 vle8.v v12, (t2)
  1116: 5700065e                 vmv.v.v v0, v12
  1120: 5703065e                 vmv.v.v v6, v12
  1124: 83338500                 ld t2, 8(a0)
  1128: b3836300                 add t2, t2, t1
  1132: 07860302                 vle8.v v12, (t2)
  1136: d700065e                 vmv.v.v v1, v12
  1140: 5703c32e                 vxor.vv v6, v12, v6
  1144: d703065e                 vmv.v.v v7, v12
  1148: 83330501                 ld t2, 0x10(a0)
  1152: b3836300                 add t2, t2, t1
  1156: 07860302                 vle8.v v12, (t2)
  1160: 5701065e                 vmv.v.v v2, v12
  1164: d783c32e                 vxor.vv v7, v12, v7
  1168: 5704065e                 vmv.v.v v8, v12
  1172: 83338501                 ld t2, 0x18(a0)
  1176: b3836300                 add t2, t2, t1
  1180: 07860302                 vle8.v v12, (t2)
  1184: d701065e                 vmv.v.v v3, v12
  1188: 5704c42e                 vxor.vv v8, v12, v8
  1192: d704065e                 vmv.v.v v9, v12
  1196: 83330502                 ld t2, 0x20(a0)
  1200: b3836300                 add t2, t2, t1
  1204: 07860302                 vle8.v v12, (t2)
  1208: 5702065e                 vmv.v.v v4, v12
  1212: d784c42e                 vxor.vv v9, v12, v9
  1216: 5705065e                 vmv.v.v v10, v12
  1220: 83338502                 ld t2, 0x28(a0)
  1224: b3836300                 add t2, t2, t1
  1228: 07860302                 vle8.v v12, (t2)
  1232: d702065e                 vmv.v.v v5, v12
  1236: 5705c52e                 vxor.vv v10, v12, v10
  1240: d705065e                 vmv.v.v v11, v12
  1244: 83330503                 ld t2, 0x30(a0)
  1248: b3836300                 add t2, t2, t1
  1252: 07860302                 vle8.v v12, (t2)
  1256: 5700c02e                 vxor.vv v0, v12, v0
  1260: d783c32e                 vxor.vv v7, v12, v7
  1264: 83338503                 ld t2, 0x38(a0)
  1268: b3836300                 add t2, t2, t1
  1272: 07860302                 vle8.v v12, (t2)
  1276: d780c02e                 vxor.vv v1, v12, v1
  1280: 5703c32e                 vxor.vv v6, v12, v6
  1284: 5704c42e                 vxor.vv v8, v12, v8
  1288: 83330504                 ld t2, 0x40(a0)
  1292: b3836300                 add t2, t2, t1
  1296: 07860302                 vle8.v v12, (t2)
  1300: 5701c12e                 vxor.vv v2, v12, v2
  1304: d783c32e                 vxor.vv v7, v12, v7
  1308: d784c42e                 vxor.vv v9, v12, v9
  1312: 83338504                 ld t2, 0x48(a0)
  1316: b3836300                 add t2, t2, t1
  1320: 07860302                 vle8.v v12, (t2)
  1324: d781c12e                 vxor.vv v3, v12, v3
  1328: 5704c42e                 vxor.vv v8, v12, v8
  1332: 5705c52e                 vxor.vv v10, v12, v10
  1336: 83330505                 ld t2, 0x50(a0)
  1340: b3836300                 add t2, t2, t1
  1344: 07860302                 vle8.v v12, (t2)
  1348: 5702c22e                 vxor.vv v4, v12, v4
  1352: d784c42e                 vxor.vv v9, v12, v9
  1356: d785c52e                 vxor.vv v11, v12, v11
  1360: 83338505                 ld t2, 0x58(a0)
  1364: b3836300                 add t2, t2, t1
  1368: 07860302                 vle8.v v12, (t2)
  1372: d782c22e                 vxor.vv v5, v12, v5
  1376: 5705c52e                 vxor.vv v10, v12, v10
  1380: 83330506                 ld t2, 0x60(a0)
  1384: b3836300                 add t2, t2, t1
  1388: 07860302                 vle8.v v12, (t2)
  1392: 5700c02e                 vxor.vv v0, v12, v0
  1396: 5704c42e                 vxor.vv v8, v12, v8
  1400: 83338506                 ld t2, 0x68(a0)
  1404: b3836300                 add t2, t2, t1
  1408: 07860302                 vle8.v v12, (t2)
  1412: d780c02e                 vxor.vv v1, v12, v1
  1416: 5703c32e                 vxor.vv v6, v12, v6
  1420: d784c42e                 vxor.vv v9, v12, v9
  1424: 83330507                 ld t2, 0x70(a0)
  1428: b3836300                 add t2, t2, t1
  1432: 07860302                 vle8.v v12, (t2)
  1436: 5701c12e                 vxor.vv v2, v12, v2
  1440: d783c32e                 vxor.vv v7, v12, v7
  1444: 5705c52e                 vxor.vv v10, v12, v10
  1448: 83338507                 ld t2, 0x78(a0)
  1452: b3836300                 add t2, t2, t1
  1456: 07860302                 vle8.v v12, (t2)
  1460: d781c12e                 vxor.vv v3, v12, v3
  1464: 5704c42e                 vxor.vv v8, v12, v8
  1468: d785c52e                 vxor.vv v11, v12, v11
  1472: 83330508                 ld t2, 0x80(a0)
  1476: b3836300                 add t2, t2, t1
  1480: 07860302                 vle8.v v12, (t2)
  1484: 5702c22e                 vxor.vv v4, v12, v4
  1488: d784c42e                 vxor.vv v9, v12, v9
  1492: 83338508                 ld t2, 0x88(a0)
  1496: b3836300                 add t2, t2, t1
  1500: 07860302                 vle8.v v12, (t2)
  1504: d782c22e                 vxor.vv v5, v12, v5
  1508: 5703c32e                 vxor.vv v6, v12, v6
  1512: 5705c52e                 vxor.vv v10, v12, v10
  1516: 83330509                 ld t2, 0x90(a0)
  1520: b3836300                 add t2, t2, t1
  1524: 07860302                 vle8.v v12, (t2)
  1528: 5700c02e                 vxor.vv v0, v12, v0
  1532: d784c42e                 vxor.vv v9, v12, v9
  1536: 83338509                 ld t2, 0x98(a0)
  1540: b3836300                 add t2, t2, t1
  1544: 07860302                 vle8.v v12, (t2)
  1548: d780c02e                 vxor.vv v1, v12, v1
  1552: 5703c32e                 vxor.vv v6, v12, v6
  1556: 5705c52e                 vxor.vv v10, v12, v10
  1560: 8333050a                 ld t2, 0xa0(a0)
  1564: b3836300                 add t2, t2, t1
  1568: 07860302                 vle8.v v12, (t2)
  1572: 5701c12e                 vxor.vv v2, v12, v2
  1576: d783c32e                 vxor.vv v7, v12, v7
  1580: d785c52e                 vxor.vv v11, v12, v11
  1584: 8333850a                 ld t2, 0xa8(a0)
  1588: b3836300                 add t2, t2, t1
  1592: 07860302                 vle8.v v12, (t2)
  1596: d781c12e                 vxor.vv v3, v12, v3
  1600: 5704c42e                 vxor.vv v8, v12, v8
  1604: 8333050b                 ld t2, 0xb0(a0)
  1608: b3836300                 add t2, t2, t1
  1612: 07860302                 vle8.v v12, (t2)
  1616: 5702c22e                 vxor.vv v4, v12, v4
  1620: 5703c32e                 vxor.vv v6, v12, v6
  1624: d784c42e                 vxor.vv v9, v12, v9
  1628: 8333850b                 ld t2, 0xb8(a0)
  1632: b3836300                 add t2, t2, t1
  1636: 07860302                 vle8.v v12, (t2)
  1640: d782c22e                 vxor.vv v5, v12, v5
  1644: d783c32e                 vxor.vv v7, v12, v7
  1648: 5705c52e                 vxor.vv v10, v12, v10
  1652: 8333050c                 ld t2, 0xc0(a0)
  1656: b3836300                 add t2, t2, t1
  1660: 07860302                 vle8.v v12, (t2)
  1664: 5700c02e                 vxor.vv v0, v12, v0
  1668: 5705c52e                 vxor.vv v10, v12, v10
  1672: 8333850c                 ld t2, 0xc8(a0)
  1676: b3836300                 add t2, t2, t1
  1680: 07860302                 vle8.v v12, (t2)
  1684: d780c02e                 vxor.vv v1, v12, v1
  1688: 5703c32e                 vxor.vv v6, v12, v6
  1692: d785c52e                 vxor.vv v11, v12, v11
  1696: 8333050d                 ld t2, 0xd0(a0)
  1700: b3836300                 add t2, t2, t1
  1704: 07860302                 vle8.v v12, (t2)
  1708: 5701c12e                 vxor.vv v2, v12, v2
  1712: d783c32e                 vxor.vv v7, v12, v7
  1716: 8333850d                 ld t2, 0xd8(a0)
  1720: b3836300                 add t2, t2, t1
  1724: 07860302                 vle8.v v12, (t2)
  1728: d781c12e                 vxor.vv v3, v12, v3
  1732: 5703c32e                 vxor.vv v6, v12, v6
  1736: 5704c42e                 vxor.vv v8, v12, v8
  1740: 8333050e                 ld t2, 0xe0(a0)
  1744: b3836300                 add t2, t2, t1
  1748: 07860302                 vle8.v v12, (t2)
  1752: 5702c22e                 vxor.vv v4, v12, v4
  1756: d783c32e                 vxor.vv v7, v12, v7
  1760: d784c42e                 vxor.vv v9, v12, v9
  1764: 8333850e                 ld t2, 0xe8(a0)
  1768: b3836300                 add t2, t2, t1
  1772: 07860302                 vle8.v v12, (t2)
  1776: d782c22e                 vxor.vv v5, v12, v5
  1780: 5704c42e                 vxor.vv v8, v12, v8
  1784: 5705c52e                 vxor.vv v10, v12, v10
  1788: 8333050f                 ld t2, 0xf0(a0)
  1792: b3836300                 add t2, t2, t1
  1796: 07860302                 vle8.v v12, (t2)
  1800: 5700c02e                 vxor.vv v0, v12, v0
  1804: d785c52e                 vxor.vv v11, v12, v11
  1808: 8333850f                 ld t2, 0xf8(a0)
  1812: b3836300                 add t2, t2, t1
  1816: 07860302                 vle8.v v12, (t2)
  1820: d780c02e                 vxor.vv v1, v12, v1
  1824: 5703c32e                 vxor.vv v6, v12, v6
  1828: 83330510                 ld t2, 0x100(a0)
  1832: b3836300                 add t2, t2, t1
  1836: 07860302                 vle8.v v12, (t2)
  1840: 5701c12e                 vxor.vv v2, v12, v2
  1844: 5703c32e                 vxor.vv v6, v12, v6
  1848: d783c32e                 vxor.vv v7, v12, v7
  1852: 83338510                 ld t2, 0x108(a0)
  1856: b3836300                 add t2, t2, t1
  1860: 07860302                 vle8.v v12, (t2)
  1864: d781c12e                 vxor.vv v3, v12, v3
  1868: d783c32e                 vxor.vv v7, v12, v7
  1872: 5704c42e                 vxor.vv v8, v12, v8
  1876: 83330511                 ld t2, 0x110(a0)
  1880: b3836300                 add t2, t2, t1
  1884: 07860302                 vle8.v v12, (t2)
  1888: 5702c22e                 vxor.vv v4, v12, v4
  1892: 5704c42e                 vxor.vv v8, v12, v8
  1896: d784c42e                 vxor.vv v9, v12, v9
  1900: 83338511                 ld t2, 0x118(a0)
  1904: b3836300                 add t2, t2, t1
  1908: 07860302                 vle8.v v12, (t2)
  1912: d782c22e                 vxor.vv v5, v12, v5
  1916: d784c42e                 vxor.vv v9, v12, v9
  1920: 5705c52e                 vxor.vv v10, v12, v10
  1924: 83b30500                 ld t2, 0(a1)
  1928: b3836300                 add t2, t2, t1
  1932: 33005000                 ntl.all 
  1936: 27800302                 vse8.v v0, (t2)
  1940: 83b38500                 ld t2, 8(a1)
  1944: b3836300                 add t2, t2, t1
  1948: 33005000                 ntl.all 
  1952: a7800302                 vse8.v v1, (t2)
  1956: 83b30501                 ld t2, 0x10(a1)
  1960: b3836300                 add t2, t2, t1
  1964: 33005000                 ntl.all 
  1968: 27810302                 vse8.v v2, (t2)
  1972: 83b38501                 ld t2, 0x18(a1)
  1976: b3836300                 add t2, t2, t1
  1980: 33005000                 ntl.all 
  1984: a7810302                 vse8.v v3, (t2)
  1988: 83b30502                 ld t2, 0x20(a1)
  1992: b3836300                 add t2, t2, t1
  1996: 33005000                 ntl.all 
  2000: 27820302                 vse8.v v4, (t2)
  2004: 83b38502                 ld t2, 0x28(a1)
  2008: b3836300                 add t2, t2, t1
  2012: 33005000                 ntl.all 
  2016: a7820302                 vse8.v v5, (t2)
  2020: 83b30503                 ld t2, 0x30(a1)
  2024: b3836300                 add t2, t2, t1
  2028: 33005000                 ntl.all 
  2032: 27830302                 vse8.v v6, (t2)
  2036: 83b38503                 ld t2, 0x38(a1)
  2040: b3836300                 add t2, t2, t1
  2044: 33005000                 ntl.all 
  2048: a7830302                 vse8.v v7, (t2)
  2052: 83b30504                 ld t2, 0x40(a1)
  2056: b3836300                 add t2, t2, t1
  2060: 33005000                 ntl.all 
  2064: 27840302                 vse8.v v8, (t2)
  2068: 83b38504                 ld t2, 0x48(a1)
  2072: b3836300                 add t2, t2, t1
  2076: 33005000                 ntl.all 
  2080: a7840302                 vse8.v v9, (t2)
  2084: 83b30505                 ld t2, 0x50(a1)
  2088: b3836300                 add t2, t2, t1
  2092: 33005000                 ntl.all 
  2096: 27850302                 vse8.v v10, (t2)
  2100: 83b38505                 ld t2, 0x58(a1)
  2104: b3836300                 add t2, t2, t1
  2108: 33005000                 ntl.all 
  2112: a7850302                 vse8.v v11, (t2)
  2116: 33035300                 add t1, t1, t0
  2120: 33065640                 sub a2, a2, t0
  2124: 6ff09fbf                 j 0x444
  2128: 67800000                 ret 
  2132: 0000                     unimp 
  2134: 0000                     unimp 
  2136: 0000                     unimp 
  2138: 0000                     unimp 
  2140: 0000                     unimp 
  2142: 0000                     unimp 
