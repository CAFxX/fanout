; config_t1_1735 target=rvv-zvbc32e twin=maskedtail topo=raid5(4)/enc op=encode U=8
; code_size=1440 twins=["maskedtail"] uops_per_byte=1.5234375
     0: 63140600                 bnez a2, 8
     4: 6f008058                 j 0x58c
     8: 970f0000                 auipc t6, 0
    12: 938f8f59                 addi t6, t6, 0x598
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93068000                 li a3, 8
    32: b386d202                 mul a3, t0, a3
    36: 3357d602                 divu a4, a2, a3
    40: 63140700                 bnez a4, 0x30
    44: 6f00801d                 j 0x204
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
   432: 83b30500                 ld t2, 0(a1)
   436: b3836300                 add t2, t2, t1
   440: a7800302                 vse8.v v1, (t2)
   444: b3835300                 add t2, t2, t0
   448: 27810302                 vse8.v v2, (t2)
   452: b3835300                 add t2, t2, t0
   456: a7810302                 vse8.v v3, (t2)
   460: b3835300                 add t2, t2, t0
   464: 27820302                 vse8.v v4, (t2)
   468: b3835300                 add t2, t2, t0
   472: a7820302                 vse8.v v5, (t2)
   476: b3835300                 add t2, t2, t0
   480: 27830302                 vse8.v v6, (t2)
   484: b3835300                 add t2, t2, t0
   488: a7830302                 vse8.v v7, (t2)
   492: b3835300                 add t2, t2, t0
   496: 27840302                 vse8.v v8, (t2)
   500: 3303d300                 add t1, t1, a3
   504: 1307f7ff                 addi a4, a4, -1
   508: e31207e4                 bnez a4, 0x40
   512: 33066640                 sub a2, a2, t1
   516: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
   520: 93060600                 mv a3, a2
   524: 6344d000                 bgtz a3, 0x214
   528: 6f004006                 j 0x274
   532: b3835640                 sub t2, a3, t0
   536: 13def343                 srai t3, t2, 0x3f
   540: b3f3c301                 and t2, t2, t3
   544: b3837200                 add t2, t0, t2
   548: d7a40852                 vid.v v9
   552: 57c0936a                 vmsltu.vx v0, v9, t2
   556: 83330500                 ld t2, 0(a0)
   560: b3836300                 add t2, t2, t1
   564: 87800300                 vle8.v v1, (t2), v0.t
   568: 83338500                 ld t2, 8(a0)
   572: b3836300                 add t2, t2, t1
   576: 87840300                 vle8.v v9, (t2), v0.t
   580: d780902e                 vxor.vv v1, v9, v1
   584: 83330501                 ld t2, 0x10(a0)
   588: b3836300                 add t2, t2, t1
   592: 87840300                 vle8.v v9, (t2), v0.t
   596: d780902e                 vxor.vv v1, v9, v1
   600: 83338501                 ld t2, 0x18(a0)
   604: b3836300                 add t2, t2, t1
   608: 87840300                 vle8.v v9, (t2), v0.t
   612: d780902e                 vxor.vv v1, v9, v1
   616: 83b30500                 ld t2, 0(a1)
   620: b3836300                 add t2, t2, t1
   624: a7800300                 vse8.v v1, (t2), v0.t
   628: 33035300                 add t1, t1, t0
   632: b3865640                 sub a3, a3, t0
   636: 6344d000                 bgtz a3, 0x284
   640: 6f004006                 j 0x2e4
   644: b3835640                 sub t2, a3, t0
   648: 13def343                 srai t3, t2, 0x3f
   652: b3f3c301                 and t2, t2, t3
   656: b3837200                 add t2, t0, t2
   660: d7a40852                 vid.v v9
   664: 57c0936a                 vmsltu.vx v0, v9, t2
   668: 83330500                 ld t2, 0(a0)
   672: b3836300                 add t2, t2, t1
   676: 87800300                 vle8.v v1, (t2), v0.t
   680: 83338500                 ld t2, 8(a0)
   684: b3836300                 add t2, t2, t1
   688: 87840300                 vle8.v v9, (t2), v0.t
   692: d780902e                 vxor.vv v1, v9, v1
   696: 83330501                 ld t2, 0x10(a0)
   700: b3836300                 add t2, t2, t1
   704: 87840300                 vle8.v v9, (t2), v0.t
   708: d780902e                 vxor.vv v1, v9, v1
   712: 83338501                 ld t2, 0x18(a0)
   716: b3836300                 add t2, t2, t1
   720: 87840300                 vle8.v v9, (t2), v0.t
   724: d780902e                 vxor.vv v1, v9, v1
   728: 83b30500                 ld t2, 0(a1)
   732: b3836300                 add t2, t2, t1
   736: a7800300                 vse8.v v1, (t2), v0.t
   740: 33035300                 add t1, t1, t0
   744: b3865640                 sub a3, a3, t0
   748: 6344d000                 bgtz a3, 0x2f4
   752: 6f004006                 j 0x354
   756: b3835640                 sub t2, a3, t0
   760: 13def343                 srai t3, t2, 0x3f
   764: b3f3c301                 and t2, t2, t3
   768: b3837200                 add t2, t0, t2
   772: d7a40852                 vid.v v9
   776: 57c0936a                 vmsltu.vx v0, v9, t2
   780: 83330500                 ld t2, 0(a0)
   784: b3836300                 add t2, t2, t1
   788: 87800300                 vle8.v v1, (t2), v0.t
   792: 83338500                 ld t2, 8(a0)
   796: b3836300                 add t2, t2, t1
   800: 87840300                 vle8.v v9, (t2), v0.t
   804: d780902e                 vxor.vv v1, v9, v1
   808: 83330501                 ld t2, 0x10(a0)
   812: b3836300                 add t2, t2, t1
   816: 87840300                 vle8.v v9, (t2), v0.t
   820: d780902e                 vxor.vv v1, v9, v1
   824: 83338501                 ld t2, 0x18(a0)
   828: b3836300                 add t2, t2, t1
   832: 87840300                 vle8.v v9, (t2), v0.t
   836: d780902e                 vxor.vv v1, v9, v1
   840: 83b30500                 ld t2, 0(a1)
   844: b3836300                 add t2, t2, t1
   848: a7800300                 vse8.v v1, (t2), v0.t
   852: 33035300                 add t1, t1, t0
   856: b3865640                 sub a3, a3, t0
   860: 6344d000                 bgtz a3, 0x364
   864: 6f004006                 j 0x3c4
   868: b3835640                 sub t2, a3, t0
   872: 13def343                 srai t3, t2, 0x3f
   876: b3f3c301                 and t2, t2, t3
   880: b3837200                 add t2, t0, t2
   884: d7a40852                 vid.v v9
   888: 57c0936a                 vmsltu.vx v0, v9, t2
   892: 83330500                 ld t2, 0(a0)
   896: b3836300                 add t2, t2, t1
   900: 87800300                 vle8.v v1, (t2), v0.t
   904: 83338500                 ld t2, 8(a0)
   908: b3836300                 add t2, t2, t1
   912: 87840300                 vle8.v v9, (t2), v0.t
   916: d780902e                 vxor.vv v1, v9, v1
   920: 83330501                 ld t2, 0x10(a0)
   924: b3836300                 add t2, t2, t1
   928: 87840300                 vle8.v v9, (t2), v0.t
   932: d780902e                 vxor.vv v1, v9, v1
   936: 83338501                 ld t2, 0x18(a0)
   940: b3836300                 add t2, t2, t1
   944: 87840300                 vle8.v v9, (t2), v0.t
   948: d780902e                 vxor.vv v1, v9, v1
   952: 83b30500                 ld t2, 0(a1)
   956: b3836300                 add t2, t2, t1
   960: a7800300                 vse8.v v1, (t2), v0.t
   964: 33035300                 add t1, t1, t0
   968: b3865640                 sub a3, a3, t0
   972: 6344d000                 bgtz a3, 0x3d4
   976: 6f004006                 j 0x434
   980: b3835640                 sub t2, a3, t0
   984: 13def343                 srai t3, t2, 0x3f
   988: b3f3c301                 and t2, t2, t3
   992: b3837200                 add t2, t0, t2
   996: d7a40852                 vid.v v9
  1000: 57c0936a                 vmsltu.vx v0, v9, t2
  1004: 83330500                 ld t2, 0(a0)
  1008: b3836300                 add t2, t2, t1
  1012: 87800300                 vle8.v v1, (t2), v0.t
  1016: 83338500                 ld t2, 8(a0)
  1020: b3836300                 add t2, t2, t1
  1024: 87840300                 vle8.v v9, (t2), v0.t
  1028: d780902e                 vxor.vv v1, v9, v1
  1032: 83330501                 ld t2, 0x10(a0)
  1036: b3836300                 add t2, t2, t1
  1040: 87840300                 vle8.v v9, (t2), v0.t
  1044: d780902e                 vxor.vv v1, v9, v1
  1048: 83338501                 ld t2, 0x18(a0)
  1052: b3836300                 add t2, t2, t1
  1056: 87840300                 vle8.v v9, (t2), v0.t
  1060: d780902e                 vxor.vv v1, v9, v1
  1064: 83b30500                 ld t2, 0(a1)
  1068: b3836300                 add t2, t2, t1
  1072: a7800300                 vse8.v v1, (t2), v0.t
  1076: 33035300                 add t1, t1, t0
  1080: b3865640                 sub a3, a3, t0
  1084: 6344d000                 bgtz a3, 0x444
  1088: 6f004006                 j 0x4a4
  1092: b3835640                 sub t2, a3, t0
  1096: 13def343                 srai t3, t2, 0x3f
  1100: b3f3c301                 and t2, t2, t3
  1104: b3837200                 add t2, t0, t2
  1108: d7a40852                 vid.v v9
  1112: 57c0936a                 vmsltu.vx v0, v9, t2
  1116: 83330500                 ld t2, 0(a0)
  1120: b3836300                 add t2, t2, t1
  1124: 87800300                 vle8.v v1, (t2), v0.t
  1128: 83338500                 ld t2, 8(a0)
  1132: b3836300                 add t2, t2, t1
  1136: 87840300                 vle8.v v9, (t2), v0.t
  1140: d780902e                 vxor.vv v1, v9, v1
  1144: 83330501                 ld t2, 0x10(a0)
  1148: b3836300                 add t2, t2, t1
  1152: 87840300                 vle8.v v9, (t2), v0.t
  1156: d780902e                 vxor.vv v1, v9, v1
  1160: 83338501                 ld t2, 0x18(a0)
  1164: b3836300                 add t2, t2, t1
  1168: 87840300                 vle8.v v9, (t2), v0.t
  1172: d780902e                 vxor.vv v1, v9, v1
  1176: 83b30500                 ld t2, 0(a1)
  1180: b3836300                 add t2, t2, t1
  1184: a7800300                 vse8.v v1, (t2), v0.t
  1188: 33035300                 add t1, t1, t0
  1192: b3865640                 sub a3, a3, t0
  1196: 6344d000                 bgtz a3, 0x4b4
  1200: 6f004006                 j 0x514
  1204: b3835640                 sub t2, a3, t0
  1208: 13def343                 srai t3, t2, 0x3f
  1212: b3f3c301                 and t2, t2, t3
  1216: b3837200                 add t2, t0, t2
  1220: d7a40852                 vid.v v9
  1224: 57c0936a                 vmsltu.vx v0, v9, t2
  1228: 83330500                 ld t2, 0(a0)
  1232: b3836300                 add t2, t2, t1
  1236: 87800300                 vle8.v v1, (t2), v0.t
  1240: 83338500                 ld t2, 8(a0)
  1244: b3836300                 add t2, t2, t1
  1248: 87840300                 vle8.v v9, (t2), v0.t
  1252: d780902e                 vxor.vv v1, v9, v1
  1256: 83330501                 ld t2, 0x10(a0)
  1260: b3836300                 add t2, t2, t1
  1264: 87840300                 vle8.v v9, (t2), v0.t
  1268: d780902e                 vxor.vv v1, v9, v1
  1272: 83338501                 ld t2, 0x18(a0)
  1276: b3836300                 add t2, t2, t1
  1280: 87840300                 vle8.v v9, (t2), v0.t
  1284: d780902e                 vxor.vv v1, v9, v1
  1288: 83b30500                 ld t2, 0(a1)
  1292: b3836300                 add t2, t2, t1
  1296: a7800300                 vse8.v v1, (t2), v0.t
  1300: 33035300                 add t1, t1, t0
  1304: b3865640                 sub a3, a3, t0
  1308: 6344d000                 bgtz a3, 0x524
  1312: 6f004006                 j 0x584
  1316: b3835640                 sub t2, a3, t0
  1320: 13def343                 srai t3, t2, 0x3f
  1324: b3f3c301                 and t2, t2, t3
  1328: b3837200                 add t2, t0, t2
  1332: d7a40852                 vid.v v9
  1336: 57c0936a                 vmsltu.vx v0, v9, t2
  1340: 83330500                 ld t2, 0(a0)
  1344: b3836300                 add t2, t2, t1
  1348: 87800300                 vle8.v v1, (t2), v0.t
  1352: 83338500                 ld t2, 8(a0)
  1356: b3836300                 add t2, t2, t1
  1360: 87840300                 vle8.v v9, (t2), v0.t
  1364: d780902e                 vxor.vv v1, v9, v1
  1368: 83330501                 ld t2, 0x10(a0)
  1372: b3836300                 add t2, t2, t1
  1376: 87840300                 vle8.v v9, (t2), v0.t
  1380: d780902e                 vxor.vv v1, v9, v1
  1384: 83338501                 ld t2, 0x18(a0)
  1388: b3836300                 add t2, t2, t1
  1392: 87840300                 vle8.v v9, (t2), v0.t
  1396: d780902e                 vxor.vv v1, v9, v1
  1400: 83b30500                 ld t2, 0(a1)
  1404: b3836300                 add t2, t2, t1
  1408: a7800300                 vse8.v v1, (t2), v0.t
  1412: 33035300                 add t1, t1, t0
  1416: b3865640                 sub a3, a3, t0
  1420: 67800000                 ret 
  1424: 0000                     unimp 
  1426: 0000                     unimp 
  1428: 0000                     unimp 
  1430: 0000                     unimp 
  1432: 0000                     unimp 
  1434: 0000                     unimp 
  1436: 0000                     unimp 
  1438: 0000                     unimp 
