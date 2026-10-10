; config_t1_1891 target=rvv-zvbc32e twin=maskedtail topo=rs(6,3)/rec{0} op=recover U=8
; code_size=1888 twins=["maskedtail"] uops_per_byte=2.0234375
     0: 63140600                 bnez a2, 8
     4: 6f000075                 j 0x754
     8: 970f0000                 auipc t6, 0
    12: 938f8f75                 addi t6, t6, 0x758
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93068000                 li a3, 8
    32: b386d202                 mul a3, t0, a3
    36: 3357d602                 divu a4, a2, a3
    40: 63140700                 bnez a4, 0x30
    44: 6f00002a                 j 0x2cc
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 87800302                 vle8.v v1, (t2)
    76: b3835300                 add t2, t2, t0
    80: 07810302                 vle8.v v2, (t2)
    84: b3835300                 add t2, t2, t0
    88: 87810302                 vle8.v v3, (t2)
    92: b3835300                 add t2, t2, t0
    96: 07820302                 vle8.v v4, (t2)
   100: b3835300                 add t2, t2, t0
   104: 87820302                 vle8.v v5, (t2)
   108: b3835300                 add t2, t2, t0
   112: 07830302                 vle8.v v6, (t2)
   116: b3835300                 add t2, t2, t0
   120: 87830302                 vle8.v v7, (t2)
   124: b3835300                 add t2, t2, t0
   128: 07840302                 vle8.v v8, (t2)
   132: 83338500                 ld t2, 8(a0)
   136: b3836300                 add t2, t2, t1
   140: 87840302                 vle8.v v9, (t2)
   144: d780902e                 vxor.vv v1, v9, v1
   148: b3835300                 add t2, t2, t0
   152: 87840302                 vle8.v v9, (t2)
   156: 5701912e                 vxor.vv v2, v9, v2
   160: b3835300                 add t2, t2, t0
   164: 87840302                 vle8.v v9, (t2)
   168: d781912e                 vxor.vv v3, v9, v3
   172: b3835300                 add t2, t2, t0
   176: 87840302                 vle8.v v9, (t2)
   180: 5702922e                 vxor.vv v4, v9, v4
   184: b3835300                 add t2, t2, t0
   188: 87840302                 vle8.v v9, (t2)
   192: d782922e                 vxor.vv v5, v9, v5
   196: b3835300                 add t2, t2, t0
   200: 87840302                 vle8.v v9, (t2)
   204: 5703932e                 vxor.vv v6, v9, v6
   208: b3835300                 add t2, t2, t0
   212: 87840302                 vle8.v v9, (t2)
   216: d783932e                 vxor.vv v7, v9, v7
   220: b3835300                 add t2, t2, t0
   224: 87840302                 vle8.v v9, (t2)
   228: 5704942e                 vxor.vv v8, v9, v8
   232: 83330501                 ld t2, 0x10(a0)
   236: b3836300                 add t2, t2, t1
   240: 87840302                 vle8.v v9, (t2)
   244: d780902e                 vxor.vv v1, v9, v1
   248: b3835300                 add t2, t2, t0
   252: 87840302                 vle8.v v9, (t2)
   256: 5701912e                 vxor.vv v2, v9, v2
   260: b3835300                 add t2, t2, t0
   264: 87840302                 vle8.v v9, (t2)
   268: d781912e                 vxor.vv v3, v9, v3
   272: b3835300                 add t2, t2, t0
   276: 87840302                 vle8.v v9, (t2)
   280: 5702922e                 vxor.vv v4, v9, v4
   284: b3835300                 add t2, t2, t0
   288: 87840302                 vle8.v v9, (t2)
   292: d782922e                 vxor.vv v5, v9, v5
   296: b3835300                 add t2, t2, t0
   300: 87840302                 vle8.v v9, (t2)
   304: 5703932e                 vxor.vv v6, v9, v6
   308: b3835300                 add t2, t2, t0
   312: 87840302                 vle8.v v9, (t2)
   316: d783932e                 vxor.vv v7, v9, v7
   320: b3835300                 add t2, t2, t0
   324: 87840302                 vle8.v v9, (t2)
   328: 5704942e                 vxor.vv v8, v9, v8
   332: 83338501                 ld t2, 0x18(a0)
   336: b3836300                 add t2, t2, t1
   340: 87840302                 vle8.v v9, (t2)
   344: d780902e                 vxor.vv v1, v9, v1
   348: b3835300                 add t2, t2, t0
   352: 87840302                 vle8.v v9, (t2)
   356: 5701912e                 vxor.vv v2, v9, v2
   360: b3835300                 add t2, t2, t0
   364: 87840302                 vle8.v v9, (t2)
   368: d781912e                 vxor.vv v3, v9, v3
   372: b3835300                 add t2, t2, t0
   376: 87840302                 vle8.v v9, (t2)
   380: 5702922e                 vxor.vv v4, v9, v4
   384: b3835300                 add t2, t2, t0
   388: 87840302                 vle8.v v9, (t2)
   392: d782922e                 vxor.vv v5, v9, v5
   396: b3835300                 add t2, t2, t0
   400: 87840302                 vle8.v v9, (t2)
   404: 5703932e                 vxor.vv v6, v9, v6
   408: b3835300                 add t2, t2, t0
   412: 87840302                 vle8.v v9, (t2)
   416: d783932e                 vxor.vv v7, v9, v7
   420: b3835300                 add t2, t2, t0
   424: 87840302                 vle8.v v9, (t2)
   428: 5704942e                 vxor.vv v8, v9, v8
   432: 83330502                 ld t2, 0x20(a0)
   436: b3836300                 add t2, t2, t1
   440: 87840302                 vle8.v v9, (t2)
   444: d780902e                 vxor.vv v1, v9, v1
   448: b3835300                 add t2, t2, t0
   452: 87840302                 vle8.v v9, (t2)
   456: 5701912e                 vxor.vv v2, v9, v2
   460: b3835300                 add t2, t2, t0
   464: 87840302                 vle8.v v9, (t2)
   468: d781912e                 vxor.vv v3, v9, v3
   472: b3835300                 add t2, t2, t0
   476: 87840302                 vle8.v v9, (t2)
   480: 5702922e                 vxor.vv v4, v9, v4
   484: b3835300                 add t2, t2, t0
   488: 87840302                 vle8.v v9, (t2)
   492: d782922e                 vxor.vv v5, v9, v5
   496: b3835300                 add t2, t2, t0
   500: 87840302                 vle8.v v9, (t2)
   504: 5703932e                 vxor.vv v6, v9, v6
   508: b3835300                 add t2, t2, t0
   512: 87840302                 vle8.v v9, (t2)
   516: d783932e                 vxor.vv v7, v9, v7
   520: b3835300                 add t2, t2, t0
   524: 87840302                 vle8.v v9, (t2)
   528: 5704942e                 vxor.vv v8, v9, v8
   532: 83338502                 ld t2, 0x28(a0)
   536: b3836300                 add t2, t2, t1
   540: 87840302                 vle8.v v9, (t2)
   544: d780902e                 vxor.vv v1, v9, v1
   548: b3835300                 add t2, t2, t0
   552: 87840302                 vle8.v v9, (t2)
   556: 5701912e                 vxor.vv v2, v9, v2
   560: b3835300                 add t2, t2, t0
   564: 87840302                 vle8.v v9, (t2)
   568: d781912e                 vxor.vv v3, v9, v3
   572: b3835300                 add t2, t2, t0
   576: 87840302                 vle8.v v9, (t2)
   580: 5702922e                 vxor.vv v4, v9, v4
   584: b3835300                 add t2, t2, t0
   588: 87840302                 vle8.v v9, (t2)
   592: d782922e                 vxor.vv v5, v9, v5
   596: b3835300                 add t2, t2, t0
   600: 87840302                 vle8.v v9, (t2)
   604: 5703932e                 vxor.vv v6, v9, v6
   608: b3835300                 add t2, t2, t0
   612: 87840302                 vle8.v v9, (t2)
   616: d783932e                 vxor.vv v7, v9, v7
   620: b3835300                 add t2, t2, t0
   624: 87840302                 vle8.v v9, (t2)
   628: 5704942e                 vxor.vv v8, v9, v8
   632: 83b30500                 ld t2, 0(a1)
   636: b3836300                 add t2, t2, t1
   640: a7800302                 vse8.v v1, (t2)
   644: b3835300                 add t2, t2, t0
   648: 27810302                 vse8.v v2, (t2)
   652: b3835300                 add t2, t2, t0
   656: a7810302                 vse8.v v3, (t2)
   660: b3835300                 add t2, t2, t0
   664: 27820302                 vse8.v v4, (t2)
   668: b3835300                 add t2, t2, t0
   672: a7820302                 vse8.v v5, (t2)
   676: b3835300                 add t2, t2, t0
   680: 27830302                 vse8.v v6, (t2)
   684: b3835300                 add t2, t2, t0
   688: a7830302                 vse8.v v7, (t2)
   692: b3835300                 add t2, t2, t0
   696: 27840302                 vse8.v v8, (t2)
   700: 3303d300                 add t1, t1, a3
   704: 1307f7ff                 addi a4, a4, -1
   708: e31e07d6                 bnez a4, 0x40
   712: 33066640                 sub a2, a2, t1
   716: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
   720: 93060600                 mv a3, a2
   724: 6344d000                 bgtz a3, 0x2dc
   728: 6f004008                 j 0x35c
   732: b3835640                 sub t2, a3, t0
   736: 13def343                 srai t3, t2, 0x3f
   740: b3f3c301                 and t2, t2, t3
   744: b3837200                 add t2, t0, t2
   748: d7a40852                 vid.v v9
   752: 57c0936a                 vmsltu.vx v0, v9, t2
   756: 83330500                 ld t2, 0(a0)
   760: b3836300                 add t2, t2, t1
   764: 87800300                 vle8.v v1, (t2), v0.t
   768: 83338500                 ld t2, 8(a0)
   772: b3836300                 add t2, t2, t1
   776: 87840300                 vle8.v v9, (t2), v0.t
   780: d780902e                 vxor.vv v1, v9, v1
   784: 83330501                 ld t2, 0x10(a0)
   788: b3836300                 add t2, t2, t1
   792: 87840300                 vle8.v v9, (t2), v0.t
   796: d780902e                 vxor.vv v1, v9, v1
   800: 83338501                 ld t2, 0x18(a0)
   804: b3836300                 add t2, t2, t1
   808: 87840300                 vle8.v v9, (t2), v0.t
   812: d780902e                 vxor.vv v1, v9, v1
   816: 83330502                 ld t2, 0x20(a0)
   820: b3836300                 add t2, t2, t1
   824: 87840300                 vle8.v v9, (t2), v0.t
   828: d780902e                 vxor.vv v1, v9, v1
   832: 83338502                 ld t2, 0x28(a0)
   836: b3836300                 add t2, t2, t1
   840: 87840300                 vle8.v v9, (t2), v0.t
   844: d780902e                 vxor.vv v1, v9, v1
   848: 83b30500                 ld t2, 0(a1)
   852: b3836300                 add t2, t2, t1
   856: a7800300                 vse8.v v1, (t2), v0.t
   860: 33035300                 add t1, t1, t0
   864: b3865640                 sub a3, a3, t0
   868: 6344d000                 bgtz a3, 0x36c
   872: 6f004008                 j 0x3ec
   876: b3835640                 sub t2, a3, t0
   880: 13def343                 srai t3, t2, 0x3f
   884: b3f3c301                 and t2, t2, t3
   888: b3837200                 add t2, t0, t2
   892: d7a40852                 vid.v v9
   896: 57c0936a                 vmsltu.vx v0, v9, t2
   900: 83330500                 ld t2, 0(a0)
   904: b3836300                 add t2, t2, t1
   908: 87800300                 vle8.v v1, (t2), v0.t
   912: 83338500                 ld t2, 8(a0)
   916: b3836300                 add t2, t2, t1
   920: 87840300                 vle8.v v9, (t2), v0.t
   924: d780902e                 vxor.vv v1, v9, v1
   928: 83330501                 ld t2, 0x10(a0)
   932: b3836300                 add t2, t2, t1
   936: 87840300                 vle8.v v9, (t2), v0.t
   940: d780902e                 vxor.vv v1, v9, v1
   944: 83338501                 ld t2, 0x18(a0)
   948: b3836300                 add t2, t2, t1
   952: 87840300                 vle8.v v9, (t2), v0.t
   956: d780902e                 vxor.vv v1, v9, v1
   960: 83330502                 ld t2, 0x20(a0)
   964: b3836300                 add t2, t2, t1
   968: 87840300                 vle8.v v9, (t2), v0.t
   972: d780902e                 vxor.vv v1, v9, v1
   976: 83338502                 ld t2, 0x28(a0)
   980: b3836300                 add t2, t2, t1
   984: 87840300                 vle8.v v9, (t2), v0.t
   988: d780902e                 vxor.vv v1, v9, v1
   992: 83b30500                 ld t2, 0(a1)
   996: b3836300                 add t2, t2, t1
  1000: a7800300                 vse8.v v1, (t2), v0.t
  1004: 33035300                 add t1, t1, t0
  1008: b3865640                 sub a3, a3, t0
  1012: 6344d000                 bgtz a3, 0x3fc
  1016: 6f004008                 j 0x47c
  1020: b3835640                 sub t2, a3, t0
  1024: 13def343                 srai t3, t2, 0x3f
  1028: b3f3c301                 and t2, t2, t3
  1032: b3837200                 add t2, t0, t2
  1036: d7a40852                 vid.v v9
  1040: 57c0936a                 vmsltu.vx v0, v9, t2
  1044: 83330500                 ld t2, 0(a0)
  1048: b3836300                 add t2, t2, t1
  1052: 87800300                 vle8.v v1, (t2), v0.t
  1056: 83338500                 ld t2, 8(a0)
  1060: b3836300                 add t2, t2, t1
  1064: 87840300                 vle8.v v9, (t2), v0.t
  1068: d780902e                 vxor.vv v1, v9, v1
  1072: 83330501                 ld t2, 0x10(a0)
  1076: b3836300                 add t2, t2, t1
  1080: 87840300                 vle8.v v9, (t2), v0.t
  1084: d780902e                 vxor.vv v1, v9, v1
  1088: 83338501                 ld t2, 0x18(a0)
  1092: b3836300                 add t2, t2, t1
  1096: 87840300                 vle8.v v9, (t2), v0.t
  1100: d780902e                 vxor.vv v1, v9, v1
  1104: 83330502                 ld t2, 0x20(a0)
  1108: b3836300                 add t2, t2, t1
  1112: 87840300                 vle8.v v9, (t2), v0.t
  1116: d780902e                 vxor.vv v1, v9, v1
  1120: 83338502                 ld t2, 0x28(a0)
  1124: b3836300                 add t2, t2, t1
  1128: 87840300                 vle8.v v9, (t2), v0.t
  1132: d780902e                 vxor.vv v1, v9, v1
  1136: 83b30500                 ld t2, 0(a1)
  1140: b3836300                 add t2, t2, t1
  1144: a7800300                 vse8.v v1, (t2), v0.t
  1148: 33035300                 add t1, t1, t0
  1152: b3865640                 sub a3, a3, t0
  1156: 6344d000                 bgtz a3, 0x48c
  1160: 6f004008                 j 0x50c
  1164: b3835640                 sub t2, a3, t0
  1168: 13def343                 srai t3, t2, 0x3f
  1172: b3f3c301                 and t2, t2, t3
  1176: b3837200                 add t2, t0, t2
  1180: d7a40852                 vid.v v9
  1184: 57c0936a                 vmsltu.vx v0, v9, t2
  1188: 83330500                 ld t2, 0(a0)
  1192: b3836300                 add t2, t2, t1
  1196: 87800300                 vle8.v v1, (t2), v0.t
  1200: 83338500                 ld t2, 8(a0)
  1204: b3836300                 add t2, t2, t1
  1208: 87840300                 vle8.v v9, (t2), v0.t
  1212: d780902e                 vxor.vv v1, v9, v1
  1216: 83330501                 ld t2, 0x10(a0)
  1220: b3836300                 add t2, t2, t1
  1224: 87840300                 vle8.v v9, (t2), v0.t
  1228: d780902e                 vxor.vv v1, v9, v1
  1232: 83338501                 ld t2, 0x18(a0)
  1236: b3836300                 add t2, t2, t1
  1240: 87840300                 vle8.v v9, (t2), v0.t
  1244: d780902e                 vxor.vv v1, v9, v1
  1248: 83330502                 ld t2, 0x20(a0)
  1252: b3836300                 add t2, t2, t1
  1256: 87840300                 vle8.v v9, (t2), v0.t
  1260: d780902e                 vxor.vv v1, v9, v1
  1264: 83338502                 ld t2, 0x28(a0)
  1268: b3836300                 add t2, t2, t1
  1272: 87840300                 vle8.v v9, (t2), v0.t
  1276: d780902e                 vxor.vv v1, v9, v1
  1280: 83b30500                 ld t2, 0(a1)
  1284: b3836300                 add t2, t2, t1
  1288: a7800300                 vse8.v v1, (t2), v0.t
  1292: 33035300                 add t1, t1, t0
  1296: b3865640                 sub a3, a3, t0
  1300: 6344d000                 bgtz a3, 0x51c
  1304: 6f004008                 j 0x59c
  1308: b3835640                 sub t2, a3, t0
  1312: 13def343                 srai t3, t2, 0x3f
  1316: b3f3c301                 and t2, t2, t3
  1320: b3837200                 add t2, t0, t2
  1324: d7a40852                 vid.v v9
  1328: 57c0936a                 vmsltu.vx v0, v9, t2
  1332: 83330500                 ld t2, 0(a0)
  1336: b3836300                 add t2, t2, t1
  1340: 87800300                 vle8.v v1, (t2), v0.t
  1344: 83338500                 ld t2, 8(a0)
  1348: b3836300                 add t2, t2, t1
  1352: 87840300                 vle8.v v9, (t2), v0.t
  1356: d780902e                 vxor.vv v1, v9, v1
  1360: 83330501                 ld t2, 0x10(a0)
  1364: b3836300                 add t2, t2, t1
  1368: 87840300                 vle8.v v9, (t2), v0.t
  1372: d780902e                 vxor.vv v1, v9, v1
  1376: 83338501                 ld t2, 0x18(a0)
  1380: b3836300                 add t2, t2, t1
  1384: 87840300                 vle8.v v9, (t2), v0.t
  1388: d780902e                 vxor.vv v1, v9, v1
  1392: 83330502                 ld t2, 0x20(a0)
  1396: b3836300                 add t2, t2, t1
  1400: 87840300                 vle8.v v9, (t2), v0.t
  1404: d780902e                 vxor.vv v1, v9, v1
  1408: 83338502                 ld t2, 0x28(a0)
  1412: b3836300                 add t2, t2, t1
  1416: 87840300                 vle8.v v9, (t2), v0.t
  1420: d780902e                 vxor.vv v1, v9, v1
  1424: 83b30500                 ld t2, 0(a1)
  1428: b3836300                 add t2, t2, t1
  1432: a7800300                 vse8.v v1, (t2), v0.t
  1436: 33035300                 add t1, t1, t0
  1440: b3865640                 sub a3, a3, t0
  1444: 6344d000                 bgtz a3, 0x5ac
  1448: 6f004008                 j 0x62c
  1452: b3835640                 sub t2, a3, t0
  1456: 13def343                 srai t3, t2, 0x3f
  1460: b3f3c301                 and t2, t2, t3
  1464: b3837200                 add t2, t0, t2
  1468: d7a40852                 vid.v v9
  1472: 57c0936a                 vmsltu.vx v0, v9, t2
  1476: 83330500                 ld t2, 0(a0)
  1480: b3836300                 add t2, t2, t1
  1484: 87800300                 vle8.v v1, (t2), v0.t
  1488: 83338500                 ld t2, 8(a0)
  1492: b3836300                 add t2, t2, t1
  1496: 87840300                 vle8.v v9, (t2), v0.t
  1500: d780902e                 vxor.vv v1, v9, v1
  1504: 83330501                 ld t2, 0x10(a0)
  1508: b3836300                 add t2, t2, t1
  1512: 87840300                 vle8.v v9, (t2), v0.t
  1516: d780902e                 vxor.vv v1, v9, v1
  1520: 83338501                 ld t2, 0x18(a0)
  1524: b3836300                 add t2, t2, t1
  1528: 87840300                 vle8.v v9, (t2), v0.t
  1532: d780902e                 vxor.vv v1, v9, v1
  1536: 83330502                 ld t2, 0x20(a0)
  1540: b3836300                 add t2, t2, t1
  1544: 87840300                 vle8.v v9, (t2), v0.t
  1548: d780902e                 vxor.vv v1, v9, v1
  1552: 83338502                 ld t2, 0x28(a0)
  1556: b3836300                 add t2, t2, t1
  1560: 87840300                 vle8.v v9, (t2), v0.t
  1564: d780902e                 vxor.vv v1, v9, v1
  1568: 83b30500                 ld t2, 0(a1)
  1572: b3836300                 add t2, t2, t1
  1576: a7800300                 vse8.v v1, (t2), v0.t
  1580: 33035300                 add t1, t1, t0
  1584: b3865640                 sub a3, a3, t0
  1588: 6344d000                 bgtz a3, 0x63c
  1592: 6f004008                 j 0x6bc
  1596: b3835640                 sub t2, a3, t0
  1600: 13def343                 srai t3, t2, 0x3f
  1604: b3f3c301                 and t2, t2, t3
  1608: b3837200                 add t2, t0, t2
  1612: d7a40852                 vid.v v9
  1616: 57c0936a                 vmsltu.vx v0, v9, t2
  1620: 83330500                 ld t2, 0(a0)
  1624: b3836300                 add t2, t2, t1
  1628: 87800300                 vle8.v v1, (t2), v0.t
  1632: 83338500                 ld t2, 8(a0)
  1636: b3836300                 add t2, t2, t1
  1640: 87840300                 vle8.v v9, (t2), v0.t
  1644: d780902e                 vxor.vv v1, v9, v1
  1648: 83330501                 ld t2, 0x10(a0)
  1652: b3836300                 add t2, t2, t1
  1656: 87840300                 vle8.v v9, (t2), v0.t
  1660: d780902e                 vxor.vv v1, v9, v1
  1664: 83338501                 ld t2, 0x18(a0)
  1668: b3836300                 add t2, t2, t1
  1672: 87840300                 vle8.v v9, (t2), v0.t
  1676: d780902e                 vxor.vv v1, v9, v1
  1680: 83330502                 ld t2, 0x20(a0)
  1684: b3836300                 add t2, t2, t1
  1688: 87840300                 vle8.v v9, (t2), v0.t
  1692: d780902e                 vxor.vv v1, v9, v1
  1696: 83338502                 ld t2, 0x28(a0)
  1700: b3836300                 add t2, t2, t1
  1704: 87840300                 vle8.v v9, (t2), v0.t
  1708: d780902e                 vxor.vv v1, v9, v1
  1712: 83b30500                 ld t2, 0(a1)
  1716: b3836300                 add t2, t2, t1
  1720: a7800300                 vse8.v v1, (t2), v0.t
  1724: 33035300                 add t1, t1, t0
  1728: b3865640                 sub a3, a3, t0
  1732: 6344d000                 bgtz a3, 0x6cc
  1736: 6f004008                 j 0x74c
  1740: b3835640                 sub t2, a3, t0
  1744: 13def343                 srai t3, t2, 0x3f
  1748: b3f3c301                 and t2, t2, t3
  1752: b3837200                 add t2, t0, t2
  1756: d7a40852                 vid.v v9
  1760: 57c0936a                 vmsltu.vx v0, v9, t2
  1764: 83330500                 ld t2, 0(a0)
  1768: b3836300                 add t2, t2, t1
  1772: 87800300                 vle8.v v1, (t2), v0.t
  1776: 83338500                 ld t2, 8(a0)
  1780: b3836300                 add t2, t2, t1
  1784: 87840300                 vle8.v v9, (t2), v0.t
  1788: d780902e                 vxor.vv v1, v9, v1
  1792: 83330501                 ld t2, 0x10(a0)
  1796: b3836300                 add t2, t2, t1
  1800: 87840300                 vle8.v v9, (t2), v0.t
  1804: d780902e                 vxor.vv v1, v9, v1
  1808: 83338501                 ld t2, 0x18(a0)
  1812: b3836300                 add t2, t2, t1
  1816: 87840300                 vle8.v v9, (t2), v0.t
  1820: d780902e                 vxor.vv v1, v9, v1
  1824: 83330502                 ld t2, 0x20(a0)
  1828: b3836300                 add t2, t2, t1
  1832: 87840300                 vle8.v v9, (t2), v0.t
  1836: d780902e                 vxor.vv v1, v9, v1
  1840: 83338502                 ld t2, 0x28(a0)
  1844: b3836300                 add t2, t2, t1
  1848: 87840300                 vle8.v v9, (t2), v0.t
  1852: d780902e                 vxor.vv v1, v9, v1
  1856: 83b30500                 ld t2, 0(a1)
  1860: b3836300                 add t2, t2, t1
  1864: a7800300                 vse8.v v1, (t2), v0.t
  1868: 33035300                 add t1, t1, t0
  1872: b3865640                 sub a3, a3, t0
  1876: 67800000                 ret 
  1880: 0000                     unimp 
  1882: 0000                     unimp 
  1884: 0000                     unimp 
  1886: 0000                     unimp 
