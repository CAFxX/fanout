; config_t1_1488 target=rvv twin=maskedtail topo=rdp(7)/enc op=encode U=1
; code_size=2064 twins=["maskedtail"] uops_per_byte=18.8125
     0: 63140600                 bnez a2, 8
     4: 6f005000                 j 0x808
     8: 971f0000                 auipc t6, 1
    12: 938f8f80                 addi t6, t6, -0x7f8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f00c03e                 j 0x414
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 87860302                 vle8.v v13, (t2)
    76: d780065e                 vmv.v.v v1, v13
    80: d783065e                 vmv.v.v v7, v13
    84: 83338500                 ld t2, 8(a0)
    88: b3836300                 add t2, t2, t1
    92: 87860302                 vle8.v v13, (t2)
    96: 5781065e                 vmv.v.v v2, v13
   100: d783d32e                 vxor.vv v7, v13, v7
   104: 5784065e                 vmv.v.v v8, v13
   108: 83330501                 ld t2, 0x10(a0)
   112: b3836300                 add t2, t2, t1
   116: 87860302                 vle8.v v13, (t2)
   120: d781065e                 vmv.v.v v3, v13
   124: 5704d42e                 vxor.vv v8, v13, v8
   128: d784065e                 vmv.v.v v9, v13
   132: 83338501                 ld t2, 0x18(a0)
   136: b3836300                 add t2, t2, t1
   140: 87860302                 vle8.v v13, (t2)
   144: 5782065e                 vmv.v.v v4, v13
   148: d784d42e                 vxor.vv v9, v13, v9
   152: 5785065e                 vmv.v.v v10, v13
   156: 83330502                 ld t2, 0x20(a0)
   160: b3836300                 add t2, t2, t1
   164: 87860302                 vle8.v v13, (t2)
   168: d782065e                 vmv.v.v v5, v13
   172: 5705d52e                 vxor.vv v10, v13, v10
   176: d785065e                 vmv.v.v v11, v13
   180: 83338502                 ld t2, 0x28(a0)
   184: b3836300                 add t2, t2, t1
   188: 87860302                 vle8.v v13, (t2)
   192: 5783065e                 vmv.v.v v6, v13
   196: d785d52e                 vxor.vv v11, v13, v11
   200: 5786065e                 vmv.v.v v12, v13
   204: 83330503                 ld t2, 0x30(a0)
   208: b3836300                 add t2, t2, t1
   212: 87860302                 vle8.v v13, (t2)
   216: d780d02e                 vxor.vv v1, v13, v1
   220: 5704d42e                 vxor.vv v8, v13, v8
   224: 83338503                 ld t2, 0x38(a0)
   228: b3836300                 add t2, t2, t1
   232: 87860302                 vle8.v v13, (t2)
   236: 5701d12e                 vxor.vv v2, v13, v2
   240: d783d32e                 vxor.vv v7, v13, v7
   244: d784d42e                 vxor.vv v9, v13, v9
   248: 83330504                 ld t2, 0x40(a0)
   252: b3836300                 add t2, t2, t1
   256: 87860302                 vle8.v v13, (t2)
   260: d781d12e                 vxor.vv v3, v13, v3
   264: 5704d42e                 vxor.vv v8, v13, v8
   268: 5705d52e                 vxor.vv v10, v13, v10
   272: 83338504                 ld t2, 0x48(a0)
   276: b3836300                 add t2, t2, t1
   280: 87860302                 vle8.v v13, (t2)
   284: 5702d22e                 vxor.vv v4, v13, v4
   288: d784d42e                 vxor.vv v9, v13, v9
   292: d785d52e                 vxor.vv v11, v13, v11
   296: 83330505                 ld t2, 0x50(a0)
   300: b3836300                 add t2, t2, t1
   304: 87860302                 vle8.v v13, (t2)
   308: d782d22e                 vxor.vv v5, v13, v5
   312: 5705d52e                 vxor.vv v10, v13, v10
   316: 5706d62e                 vxor.vv v12, v13, v12
   320: 83338505                 ld t2, 0x58(a0)
   324: b3836300                 add t2, t2, t1
   328: 87860302                 vle8.v v13, (t2)
   332: 5703d32e                 vxor.vv v6, v13, v6
   336: d785d52e                 vxor.vv v11, v13, v11
   340: 83330506                 ld t2, 0x60(a0)
   344: b3836300                 add t2, t2, t1
   348: 87860302                 vle8.v v13, (t2)
   352: d780d02e                 vxor.vv v1, v13, v1
   356: d784d42e                 vxor.vv v9, v13, v9
   360: 83338506                 ld t2, 0x68(a0)
   364: b3836300                 add t2, t2, t1
   368: 87860302                 vle8.v v13, (t2)
   372: 5701d12e                 vxor.vv v2, v13, v2
   376: d783d32e                 vxor.vv v7, v13, v7
   380: 5705d52e                 vxor.vv v10, v13, v10
   384: 83330507                 ld t2, 0x70(a0)
   388: b3836300                 add t2, t2, t1
   392: 87860302                 vle8.v v13, (t2)
   396: d781d12e                 vxor.vv v3, v13, v3
   400: 5704d42e                 vxor.vv v8, v13, v8
   404: d785d52e                 vxor.vv v11, v13, v11
   408: 83338507                 ld t2, 0x78(a0)
   412: b3836300                 add t2, t2, t1
   416: 87860302                 vle8.v v13, (t2)
   420: 5702d22e                 vxor.vv v4, v13, v4
   424: d784d42e                 vxor.vv v9, v13, v9
   428: 5706d62e                 vxor.vv v12, v13, v12
   432: 83330508                 ld t2, 0x80(a0)
   436: b3836300                 add t2, t2, t1
   440: 87860302                 vle8.v v13, (t2)
   444: d782d22e                 vxor.vv v5, v13, v5
   448: 5705d52e                 vxor.vv v10, v13, v10
   452: 83338508                 ld t2, 0x88(a0)
   456: b3836300                 add t2, t2, t1
   460: 87860302                 vle8.v v13, (t2)
   464: 5703d32e                 vxor.vv v6, v13, v6
   468: d783d32e                 vxor.vv v7, v13, v7
   472: d785d52e                 vxor.vv v11, v13, v11
   476: 83330509                 ld t2, 0x90(a0)
   480: b3836300                 add t2, t2, t1
   484: 87860302                 vle8.v v13, (t2)
   488: d780d02e                 vxor.vv v1, v13, v1
   492: 5705d52e                 vxor.vv v10, v13, v10
   496: 83338509                 ld t2, 0x98(a0)
   500: b3836300                 add t2, t2, t1
   504: 87860302                 vle8.v v13, (t2)
   508: 5701d12e                 vxor.vv v2, v13, v2
   512: d783d32e                 vxor.vv v7, v13, v7
   516: d785d52e                 vxor.vv v11, v13, v11
   520: 8333050a                 ld t2, 0xa0(a0)
   524: b3836300                 add t2, t2, t1
   528: 87860302                 vle8.v v13, (t2)
   532: d781d12e                 vxor.vv v3, v13, v3
   536: 5704d42e                 vxor.vv v8, v13, v8
   540: 5706d62e                 vxor.vv v12, v13, v12
   544: 8333850a                 ld t2, 0xa8(a0)
   548: b3836300                 add t2, t2, t1
   552: 87860302                 vle8.v v13, (t2)
   556: 5702d22e                 vxor.vv v4, v13, v4
   560: d784d42e                 vxor.vv v9, v13, v9
   564: 8333050b                 ld t2, 0xb0(a0)
   568: b3836300                 add t2, t2, t1
   572: 87860302                 vle8.v v13, (t2)
   576: d782d22e                 vxor.vv v5, v13, v5
   580: d783d32e                 vxor.vv v7, v13, v7
   584: 5705d52e                 vxor.vv v10, v13, v10
   588: 8333850b                 ld t2, 0xb8(a0)
   592: b3836300                 add t2, t2, t1
   596: 87860302                 vle8.v v13, (t2)
   600: 5703d32e                 vxor.vv v6, v13, v6
   604: 5704d42e                 vxor.vv v8, v13, v8
   608: d785d52e                 vxor.vv v11, v13, v11
   612: 8333050c                 ld t2, 0xc0(a0)
   616: b3836300                 add t2, t2, t1
   620: 87860302                 vle8.v v13, (t2)
   624: d780d02e                 vxor.vv v1, v13, v1
   628: d785d52e                 vxor.vv v11, v13, v11
   632: 8333850c                 ld t2, 0xc8(a0)
   636: b3836300                 add t2, t2, t1
   640: 87860302                 vle8.v v13, (t2)
   644: 5701d12e                 vxor.vv v2, v13, v2
   648: d783d32e                 vxor.vv v7, v13, v7
   652: 5706d62e                 vxor.vv v12, v13, v12
   656: 8333050d                 ld t2, 0xd0(a0)
   660: b3836300                 add t2, t2, t1
   664: 87860302                 vle8.v v13, (t2)
   668: d781d12e                 vxor.vv v3, v13, v3
   672: 5704d42e                 vxor.vv v8, v13, v8
   676: 8333850d                 ld t2, 0xd8(a0)
   680: b3836300                 add t2, t2, t1
   684: 87860302                 vle8.v v13, (t2)
   688: 5702d22e                 vxor.vv v4, v13, v4
   692: d783d32e                 vxor.vv v7, v13, v7
   696: d784d42e                 vxor.vv v9, v13, v9
   700: 8333050e                 ld t2, 0xe0(a0)
   704: b3836300                 add t2, t2, t1
   708: 87860302                 vle8.v v13, (t2)
   712: d782d22e                 vxor.vv v5, v13, v5
   716: 5704d42e                 vxor.vv v8, v13, v8
   720: 5705d52e                 vxor.vv v10, v13, v10
   724: 8333850e                 ld t2, 0xe8(a0)
   728: b3836300                 add t2, t2, t1
   732: 87860302                 vle8.v v13, (t2)
   736: 5703d32e                 vxor.vv v6, v13, v6
   740: d784d42e                 vxor.vv v9, v13, v9
   744: d785d52e                 vxor.vv v11, v13, v11
   748: 8333050f                 ld t2, 0xf0(a0)
   752: b3836300                 add t2, t2, t1
   756: 87860302                 vle8.v v13, (t2)
   760: d780d02e                 vxor.vv v1, v13, v1
   764: 5706d62e                 vxor.vv v12, v13, v12
   768: 8333850f                 ld t2, 0xf8(a0)
   772: b3836300                 add t2, t2, t1
   776: 87860302                 vle8.v v13, (t2)
   780: 5701d12e                 vxor.vv v2, v13, v2
   784: d783d32e                 vxor.vv v7, v13, v7
   788: 83330510                 ld t2, 0x100(a0)
   792: b3836300                 add t2, t2, t1
   796: 87860302                 vle8.v v13, (t2)
   800: d781d12e                 vxor.vv v3, v13, v3
   804: d783d32e                 vxor.vv v7, v13, v7
   808: 5704d42e                 vxor.vv v8, v13, v8
   812: 83338510                 ld t2, 0x108(a0)
   816: b3836300                 add t2, t2, t1
   820: 87860302                 vle8.v v13, (t2)
   824: 5702d22e                 vxor.vv v4, v13, v4
   828: 5704d42e                 vxor.vv v8, v13, v8
   832: d784d42e                 vxor.vv v9, v13, v9
   836: 83330511                 ld t2, 0x110(a0)
   840: b3836300                 add t2, t2, t1
   844: 87860302                 vle8.v v13, (t2)
   848: d782d22e                 vxor.vv v5, v13, v5
   852: d784d42e                 vxor.vv v9, v13, v9
   856: 5705d52e                 vxor.vv v10, v13, v10
   860: 83338511                 ld t2, 0x118(a0)
   864: b3836300                 add t2, t2, t1
   868: 87860302                 vle8.v v13, (t2)
   872: 5703d32e                 vxor.vv v6, v13, v6
   876: 5705d52e                 vxor.vv v10, v13, v10
   880: d785d52e                 vxor.vv v11, v13, v11
   884: 83b30500                 ld t2, 0(a1)
   888: b3836300                 add t2, t2, t1
   892: a7800302                 vse8.v v1, (t2)
   896: 83b38500                 ld t2, 8(a1)
   900: b3836300                 add t2, t2, t1
   904: 27810302                 vse8.v v2, (t2)
   908: 83b30501                 ld t2, 0x10(a1)
   912: b3836300                 add t2, t2, t1
   916: a7810302                 vse8.v v3, (t2)
   920: 83b38501                 ld t2, 0x18(a1)
   924: b3836300                 add t2, t2, t1
   928: 27820302                 vse8.v v4, (t2)
   932: 83b30502                 ld t2, 0x20(a1)
   936: b3836300                 add t2, t2, t1
   940: a7820302                 vse8.v v5, (t2)
   944: 83b38502                 ld t2, 0x28(a1)
   948: b3836300                 add t2, t2, t1
   952: 27830302                 vse8.v v6, (t2)
   956: 83b30503                 ld t2, 0x30(a1)
   960: b3836300                 add t2, t2, t1
   964: a7830302                 vse8.v v7, (t2)
   968: 83b38503                 ld t2, 0x38(a1)
   972: b3836300                 add t2, t2, t1
   976: 27840302                 vse8.v v8, (t2)
   980: 83b30504                 ld t2, 0x40(a1)
   984: b3836300                 add t2, t2, t1
   988: a7840302                 vse8.v v9, (t2)
   992: 83b38504                 ld t2, 0x48(a1)
   996: b3836300                 add t2, t2, t1
  1000: 27850302                 vse8.v v10, (t2)
  1004: 83b30505                 ld t2, 0x50(a1)
  1008: b3836300                 add t2, t2, t1
  1012: a7850302                 vse8.v v11, (t2)
  1016: 83b38505                 ld t2, 0x58(a1)
  1020: b3836300                 add t2, t2, t1
  1024: 27860302                 vse8.v v12, (t2)
  1028: 3303d300                 add t1, t1, a3
  1032: 1307f7ff                 addi a4, a4, -1
  1036: e31a07c2                 bnez a4, 0x40
  1040: 33066640                 sub a2, a2, t1
  1044: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
  1048: 93060600                 mv a3, a2
  1052: 6344d000                 bgtz a3, 0x424
  1056: 6f00003e                 j 0x800
  1060: b3835640                 sub t2, a3, t0
  1064: 13def343                 srai t3, t2, 0x3f
  1068: b3f3c301                 and t2, t2, t3
  1072: b3837200                 add t2, t0, t2
  1076: d7a60852                 vid.v v13
  1080: 57c0d36a                 vmsltu.vx v0, v13, t2
  1084: 83330500                 ld t2, 0(a0)
  1088: b3836300                 add t2, t2, t1
  1092: 87860300                 vle8.v v13, (t2), v0.t
  1096: d780065e                 vmv.v.v v1, v13
  1100: d783065e                 vmv.v.v v7, v13
  1104: 83338500                 ld t2, 8(a0)
  1108: b3836300                 add t2, t2, t1
  1112: 87860300                 vle8.v v13, (t2), v0.t
  1116: 5781065e                 vmv.v.v v2, v13
  1120: d783d32e                 vxor.vv v7, v13, v7
  1124: 5784065e                 vmv.v.v v8, v13
  1128: 83330501                 ld t2, 0x10(a0)
  1132: b3836300                 add t2, t2, t1
  1136: 87860300                 vle8.v v13, (t2), v0.t
  1140: d781065e                 vmv.v.v v3, v13
  1144: 5704d42e                 vxor.vv v8, v13, v8
  1148: d784065e                 vmv.v.v v9, v13
  1152: 83338501                 ld t2, 0x18(a0)
  1156: b3836300                 add t2, t2, t1
  1160: 87860300                 vle8.v v13, (t2), v0.t
  1164: 5782065e                 vmv.v.v v4, v13
  1168: d784d42e                 vxor.vv v9, v13, v9
  1172: 5785065e                 vmv.v.v v10, v13
  1176: 83330502                 ld t2, 0x20(a0)
  1180: b3836300                 add t2, t2, t1
  1184: 87860300                 vle8.v v13, (t2), v0.t
  1188: d782065e                 vmv.v.v v5, v13
  1192: 5705d52e                 vxor.vv v10, v13, v10
  1196: d785065e                 vmv.v.v v11, v13
  1200: 83338502                 ld t2, 0x28(a0)
  1204: b3836300                 add t2, t2, t1
  1208: 87860300                 vle8.v v13, (t2), v0.t
  1212: 5783065e                 vmv.v.v v6, v13
  1216: d785d52e                 vxor.vv v11, v13, v11
  1220: 5786065e                 vmv.v.v v12, v13
  1224: 83330503                 ld t2, 0x30(a0)
  1228: b3836300                 add t2, t2, t1
  1232: 87860300                 vle8.v v13, (t2), v0.t
  1236: d780d02e                 vxor.vv v1, v13, v1
  1240: 5704d42e                 vxor.vv v8, v13, v8
  1244: 83338503                 ld t2, 0x38(a0)
  1248: b3836300                 add t2, t2, t1
  1252: 87860300                 vle8.v v13, (t2), v0.t
  1256: 5701d12e                 vxor.vv v2, v13, v2
  1260: d783d32e                 vxor.vv v7, v13, v7
  1264: d784d42e                 vxor.vv v9, v13, v9
  1268: 83330504                 ld t2, 0x40(a0)
  1272: b3836300                 add t2, t2, t1
  1276: 87860300                 vle8.v v13, (t2), v0.t
  1280: d781d12e                 vxor.vv v3, v13, v3
  1284: 5704d42e                 vxor.vv v8, v13, v8
  1288: 5705d52e                 vxor.vv v10, v13, v10
  1292: 83338504                 ld t2, 0x48(a0)
  1296: b3836300                 add t2, t2, t1
  1300: 87860300                 vle8.v v13, (t2), v0.t
  1304: 5702d22e                 vxor.vv v4, v13, v4
  1308: d784d42e                 vxor.vv v9, v13, v9
  1312: d785d52e                 vxor.vv v11, v13, v11
  1316: 83330505                 ld t2, 0x50(a0)
  1320: b3836300                 add t2, t2, t1
  1324: 87860300                 vle8.v v13, (t2), v0.t
  1328: d782d22e                 vxor.vv v5, v13, v5
  1332: 5705d52e                 vxor.vv v10, v13, v10
  1336: 5706d62e                 vxor.vv v12, v13, v12
  1340: 83338505                 ld t2, 0x58(a0)
  1344: b3836300                 add t2, t2, t1
  1348: 87860300                 vle8.v v13, (t2), v0.t
  1352: 5703d32e                 vxor.vv v6, v13, v6
  1356: d785d52e                 vxor.vv v11, v13, v11
  1360: 83330506                 ld t2, 0x60(a0)
  1364: b3836300                 add t2, t2, t1
  1368: 87860300                 vle8.v v13, (t2), v0.t
  1372: d780d02e                 vxor.vv v1, v13, v1
  1376: d784d42e                 vxor.vv v9, v13, v9
  1380: 83338506                 ld t2, 0x68(a0)
  1384: b3836300                 add t2, t2, t1
  1388: 87860300                 vle8.v v13, (t2), v0.t
  1392: 5701d12e                 vxor.vv v2, v13, v2
  1396: d783d32e                 vxor.vv v7, v13, v7
  1400: 5705d52e                 vxor.vv v10, v13, v10
  1404: 83330507                 ld t2, 0x70(a0)
  1408: b3836300                 add t2, t2, t1
  1412: 87860300                 vle8.v v13, (t2), v0.t
  1416: d781d12e                 vxor.vv v3, v13, v3
  1420: 5704d42e                 vxor.vv v8, v13, v8
  1424: d785d52e                 vxor.vv v11, v13, v11
  1428: 83338507                 ld t2, 0x78(a0)
  1432: b3836300                 add t2, t2, t1
  1436: 87860300                 vle8.v v13, (t2), v0.t
  1440: 5702d22e                 vxor.vv v4, v13, v4
  1444: d784d42e                 vxor.vv v9, v13, v9
  1448: 5706d62e                 vxor.vv v12, v13, v12
  1452: 83330508                 ld t2, 0x80(a0)
  1456: b3836300                 add t2, t2, t1
  1460: 87860300                 vle8.v v13, (t2), v0.t
  1464: d782d22e                 vxor.vv v5, v13, v5
  1468: 5705d52e                 vxor.vv v10, v13, v10
  1472: 83338508                 ld t2, 0x88(a0)
  1476: b3836300                 add t2, t2, t1
  1480: 87860300                 vle8.v v13, (t2), v0.t
  1484: 5703d32e                 vxor.vv v6, v13, v6
  1488: d783d32e                 vxor.vv v7, v13, v7
  1492: d785d52e                 vxor.vv v11, v13, v11
  1496: 83330509                 ld t2, 0x90(a0)
  1500: b3836300                 add t2, t2, t1
  1504: 87860300                 vle8.v v13, (t2), v0.t
  1508: d780d02e                 vxor.vv v1, v13, v1
  1512: 5705d52e                 vxor.vv v10, v13, v10
  1516: 83338509                 ld t2, 0x98(a0)
  1520: b3836300                 add t2, t2, t1
  1524: 87860300                 vle8.v v13, (t2), v0.t
  1528: 5701d12e                 vxor.vv v2, v13, v2
  1532: d783d32e                 vxor.vv v7, v13, v7
  1536: d785d52e                 vxor.vv v11, v13, v11
  1540: 8333050a                 ld t2, 0xa0(a0)
  1544: b3836300                 add t2, t2, t1
  1548: 87860300                 vle8.v v13, (t2), v0.t
  1552: d781d12e                 vxor.vv v3, v13, v3
  1556: 5704d42e                 vxor.vv v8, v13, v8
  1560: 5706d62e                 vxor.vv v12, v13, v12
  1564: 8333850a                 ld t2, 0xa8(a0)
  1568: b3836300                 add t2, t2, t1
  1572: 87860300                 vle8.v v13, (t2), v0.t
  1576: 5702d22e                 vxor.vv v4, v13, v4
  1580: d784d42e                 vxor.vv v9, v13, v9
  1584: 8333050b                 ld t2, 0xb0(a0)
  1588: b3836300                 add t2, t2, t1
  1592: 87860300                 vle8.v v13, (t2), v0.t
  1596: d782d22e                 vxor.vv v5, v13, v5
  1600: d783d32e                 vxor.vv v7, v13, v7
  1604: 5705d52e                 vxor.vv v10, v13, v10
  1608: 8333850b                 ld t2, 0xb8(a0)
  1612: b3836300                 add t2, t2, t1
  1616: 87860300                 vle8.v v13, (t2), v0.t
  1620: 5703d32e                 vxor.vv v6, v13, v6
  1624: 5704d42e                 vxor.vv v8, v13, v8
  1628: d785d52e                 vxor.vv v11, v13, v11
  1632: 8333050c                 ld t2, 0xc0(a0)
  1636: b3836300                 add t2, t2, t1
  1640: 87860300                 vle8.v v13, (t2), v0.t
  1644: d780d02e                 vxor.vv v1, v13, v1
  1648: d785d52e                 vxor.vv v11, v13, v11
  1652: 8333850c                 ld t2, 0xc8(a0)
  1656: b3836300                 add t2, t2, t1
  1660: 87860300                 vle8.v v13, (t2), v0.t
  1664: 5701d12e                 vxor.vv v2, v13, v2
  1668: d783d32e                 vxor.vv v7, v13, v7
  1672: 5706d62e                 vxor.vv v12, v13, v12
  1676: 8333050d                 ld t2, 0xd0(a0)
  1680: b3836300                 add t2, t2, t1
  1684: 87860300                 vle8.v v13, (t2), v0.t
  1688: d781d12e                 vxor.vv v3, v13, v3
  1692: 5704d42e                 vxor.vv v8, v13, v8
  1696: 8333850d                 ld t2, 0xd8(a0)
  1700: b3836300                 add t2, t2, t1
  1704: 87860300                 vle8.v v13, (t2), v0.t
  1708: 5702d22e                 vxor.vv v4, v13, v4
  1712: d783d32e                 vxor.vv v7, v13, v7
  1716: d784d42e                 vxor.vv v9, v13, v9
  1720: 8333050e                 ld t2, 0xe0(a0)
  1724: b3836300                 add t2, t2, t1
  1728: 87860300                 vle8.v v13, (t2), v0.t
  1732: d782d22e                 vxor.vv v5, v13, v5
  1736: 5704d42e                 vxor.vv v8, v13, v8
  1740: 5705d52e                 vxor.vv v10, v13, v10
  1744: 8333850e                 ld t2, 0xe8(a0)
  1748: b3836300                 add t2, t2, t1
  1752: 87860300                 vle8.v v13, (t2), v0.t
  1756: 5703d32e                 vxor.vv v6, v13, v6
  1760: d784d42e                 vxor.vv v9, v13, v9
  1764: d785d52e                 vxor.vv v11, v13, v11
  1768: 8333050f                 ld t2, 0xf0(a0)
  1772: b3836300                 add t2, t2, t1
  1776: 87860300                 vle8.v v13, (t2), v0.t
  1780: d780d02e                 vxor.vv v1, v13, v1
  1784: 5706d62e                 vxor.vv v12, v13, v12
  1788: 8333850f                 ld t2, 0xf8(a0)
  1792: b3836300                 add t2, t2, t1
  1796: 87860300                 vle8.v v13, (t2), v0.t
  1800: 5701d12e                 vxor.vv v2, v13, v2
  1804: d783d32e                 vxor.vv v7, v13, v7
  1808: 83330510                 ld t2, 0x100(a0)
  1812: b3836300                 add t2, t2, t1
  1816: 87860300                 vle8.v v13, (t2), v0.t
  1820: d781d12e                 vxor.vv v3, v13, v3
  1824: d783d32e                 vxor.vv v7, v13, v7
  1828: 5704d42e                 vxor.vv v8, v13, v8
  1832: 83338510                 ld t2, 0x108(a0)
  1836: b3836300                 add t2, t2, t1
  1840: 87860300                 vle8.v v13, (t2), v0.t
  1844: 5702d22e                 vxor.vv v4, v13, v4
  1848: 5704d42e                 vxor.vv v8, v13, v8
  1852: d784d42e                 vxor.vv v9, v13, v9
  1856: 83330511                 ld t2, 0x110(a0)
  1860: b3836300                 add t2, t2, t1
  1864: 87860300                 vle8.v v13, (t2), v0.t
  1868: d782d22e                 vxor.vv v5, v13, v5
  1872: d784d42e                 vxor.vv v9, v13, v9
  1876: 5705d52e                 vxor.vv v10, v13, v10
  1880: 83338511                 ld t2, 0x118(a0)
  1884: b3836300                 add t2, t2, t1
  1888: 87860300                 vle8.v v13, (t2), v0.t
  1892: 5703d32e                 vxor.vv v6, v13, v6
  1896: 5705d52e                 vxor.vv v10, v13, v10
  1900: d785d52e                 vxor.vv v11, v13, v11
  1904: 83b30500                 ld t2, 0(a1)
  1908: b3836300                 add t2, t2, t1
  1912: a7800300                 vse8.v v1, (t2), v0.t
  1916: 83b38500                 ld t2, 8(a1)
  1920: b3836300                 add t2, t2, t1
  1924: 27810300                 vse8.v v2, (t2), v0.t
  1928: 83b30501                 ld t2, 0x10(a1)
  1932: b3836300                 add t2, t2, t1
  1936: a7810300                 vse8.v v3, (t2), v0.t
  1940: 83b38501                 ld t2, 0x18(a1)
  1944: b3836300                 add t2, t2, t1
  1948: 27820300                 vse8.v v4, (t2), v0.t
  1952: 83b30502                 ld t2, 0x20(a1)
  1956: b3836300                 add t2, t2, t1
  1960: a7820300                 vse8.v v5, (t2), v0.t
  1964: 83b38502                 ld t2, 0x28(a1)
  1968: b3836300                 add t2, t2, t1
  1972: 27830300                 vse8.v v6, (t2), v0.t
  1976: 83b30503                 ld t2, 0x30(a1)
  1980: b3836300                 add t2, t2, t1
  1984: a7830300                 vse8.v v7, (t2), v0.t
  1988: 83b38503                 ld t2, 0x38(a1)
  1992: b3836300                 add t2, t2, t1
  1996: 27840300                 vse8.v v8, (t2), v0.t
  2000: 83b30504                 ld t2, 0x40(a1)
  2004: b3836300                 add t2, t2, t1
  2008: a7840300                 vse8.v v9, (t2), v0.t
  2012: 83b38504                 ld t2, 0x48(a1)
  2016: b3836300                 add t2, t2, t1
  2020: 27850300                 vse8.v v10, (t2), v0.t
  2024: 83b30505                 ld t2, 0x50(a1)
  2028: b3836300                 add t2, t2, t1
  2032: a7850300                 vse8.v v11, (t2), v0.t
  2036: 83b38505                 ld t2, 0x58(a1)
  2040: b3836300                 add t2, t2, t1
  2044: 27860300                 vse8.v v12, (t2), v0.t
  2048: 33035300                 add t1, t1, t0
  2052: b3865640                 sub a3, a3, t0
  2056: 67800000                 ret 
  2060: 0000                     unimp 
  2062: 0000                     unimp 
