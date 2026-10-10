; config_t3_2124 target=rvv-zvbc32e twin=baseline topo=lrc_azure()/rec{0} op=recover U=4
; code_size=2560 twins=[] uops_per_byte=8.171875
     0: 63140600                 bnez a2, 8
     4: 6f00501f                 j 0x9f8
     8: 971f0000                 auipc t6, 1
    12: 938f8f9f                 addi t6, t6, -0x608
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 57470e5e                 vmv.v.x v14, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93962200                 slli a3, t0, 2
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00406d                 j 0x704
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07820302                 vle8.v v4, (t2)
    76: b3835300                 add t2, t2, t0
    80: 87820302                 vle8.v v5, (t2)
    84: b3835300                 add t2, t2, t0
    88: 07830302                 vle8.v v6, (t2)
    92: b3835300                 add t2, t2, t0
    96: 87830302                 vle8.v v7, (t2)
   100: 138e0f00                 mv t3, t6
   104: 07060e02                 vle8.v v12, (t3)
   108: 138e0f01                 addi t3, t6, 0x10
   112: 87060e02                 vle8.v v13, (t3)
   116: 5704e226                 vand.vv v8, v14, v4
   120: d73442a2                 vsrl.vi v9, v4, 4
   124: 5705c432                 vrgather.vv v10, v12, v8
   128: d785d432                 vrgather.vv v11, v13, v9
   132: 5705b52e                 vxor.vv v10, v11, v10
   136: 5700055e                 vmv.v.v v0, v10
   140: 5784e226                 vand.vv v8, v14, v5
   144: d73452a2                 vsrl.vi v9, v5, 4
   148: 5705c432                 vrgather.vv v10, v12, v8
   152: d785d432                 vrgather.vv v11, v13, v9
   156: 5705b52e                 vxor.vv v10, v11, v10
   160: d700055e                 vmv.v.v v1, v10
   164: 5704e326                 vand.vv v8, v14, v6
   168: d73462a2                 vsrl.vi v9, v6, 4
   172: 5705c432                 vrgather.vv v10, v12, v8
   176: d785d432                 vrgather.vv v11, v13, v9
   180: 5705b52e                 vxor.vv v10, v11, v10
   184: 5701055e                 vmv.v.v v2, v10
   188: 5784e326                 vand.vv v8, v14, v7
   192: d73472a2                 vsrl.vi v9, v7, 4
   196: 5705c432                 vrgather.vv v10, v12, v8
   200: d785d432                 vrgather.vv v11, v13, v9
   204: 5705b52e                 vxor.vv v10, v11, v10
   208: d701055e                 vmv.v.v v3, v10
   212: 83338500                 ld t2, 8(a0)
   216: b3836300                 add t2, t2, t1
   220: 07820302                 vle8.v v4, (t2)
   224: b3835300                 add t2, t2, t0
   228: 87820302                 vle8.v v5, (t2)
   232: b3835300                 add t2, t2, t0
   236: 07830302                 vle8.v v6, (t2)
   240: b3835300                 add t2, t2, t0
   244: 87830302                 vle8.v v7, (t2)
   248: 138e0f02                 addi t3, t6, 0x20
   252: 07060e02                 vle8.v v12, (t3)
   256: 138e0f03                 addi t3, t6, 0x30
   260: 87060e02                 vle8.v v13, (t3)
   264: 5704e226                 vand.vv v8, v14, v4
   268: d73442a2                 vsrl.vi v9, v4, 4
   272: 5705c432                 vrgather.vv v10, v12, v8
   276: d785d432                 vrgather.vv v11, v13, v9
   280: 5705b52e                 vxor.vv v10, v11, v10
   284: 5700a02e                 vxor.vv v0, v10, v0
   288: 5784e226                 vand.vv v8, v14, v5
   292: d73452a2                 vsrl.vi v9, v5, 4
   296: 5705c432                 vrgather.vv v10, v12, v8
   300: d785d432                 vrgather.vv v11, v13, v9
   304: 5705b52e                 vxor.vv v10, v11, v10
   308: d780a02e                 vxor.vv v1, v10, v1
   312: 5704e326                 vand.vv v8, v14, v6
   316: d73462a2                 vsrl.vi v9, v6, 4
   320: 5705c432                 vrgather.vv v10, v12, v8
   324: d785d432                 vrgather.vv v11, v13, v9
   328: 5705b52e                 vxor.vv v10, v11, v10
   332: 5701a12e                 vxor.vv v2, v10, v2
   336: 5784e326                 vand.vv v8, v14, v7
   340: d73472a2                 vsrl.vi v9, v7, 4
   344: 5705c432                 vrgather.vv v10, v12, v8
   348: d785d432                 vrgather.vv v11, v13, v9
   352: 5705b52e                 vxor.vv v10, v11, v10
   356: d781a12e                 vxor.vv v3, v10, v3
   360: 83330501                 ld t2, 0x10(a0)
   364: b3836300                 add t2, t2, t1
   368: 07820302                 vle8.v v4, (t2)
   372: b3835300                 add t2, t2, t0
   376: 87820302                 vle8.v v5, (t2)
   380: b3835300                 add t2, t2, t0
   384: 07830302                 vle8.v v6, (t2)
   388: b3835300                 add t2, t2, t0
   392: 87830302                 vle8.v v7, (t2)
   396: 138e0f04                 addi t3, t6, 0x40
   400: 07060e02                 vle8.v v12, (t3)
   404: 138e0f05                 addi t3, t6, 0x50
   408: 87060e02                 vle8.v v13, (t3)
   412: 5704e226                 vand.vv v8, v14, v4
   416: d73442a2                 vsrl.vi v9, v4, 4
   420: 5705c432                 vrgather.vv v10, v12, v8
   424: d785d432                 vrgather.vv v11, v13, v9
   428: 5705b52e                 vxor.vv v10, v11, v10
   432: 5700a02e                 vxor.vv v0, v10, v0
   436: 5784e226                 vand.vv v8, v14, v5
   440: d73452a2                 vsrl.vi v9, v5, 4
   444: 5705c432                 vrgather.vv v10, v12, v8
   448: d785d432                 vrgather.vv v11, v13, v9
   452: 5705b52e                 vxor.vv v10, v11, v10
   456: d780a02e                 vxor.vv v1, v10, v1
   460: 5704e326                 vand.vv v8, v14, v6
   464: d73462a2                 vsrl.vi v9, v6, 4
   468: 5705c432                 vrgather.vv v10, v12, v8
   472: d785d432                 vrgather.vv v11, v13, v9
   476: 5705b52e                 vxor.vv v10, v11, v10
   480: 5701a12e                 vxor.vv v2, v10, v2
   484: 5784e326                 vand.vv v8, v14, v7
   488: d73472a2                 vsrl.vi v9, v7, 4
   492: 5705c432                 vrgather.vv v10, v12, v8
   496: d785d432                 vrgather.vv v11, v13, v9
   500: 5705b52e                 vxor.vv v10, v11, v10
   504: d781a12e                 vxor.vv v3, v10, v3
   508: 83338501                 ld t2, 0x18(a0)
   512: b3836300                 add t2, t2, t1
   516: 07820302                 vle8.v v4, (t2)
   520: b3835300                 add t2, t2, t0
   524: 87820302                 vle8.v v5, (t2)
   528: b3835300                 add t2, t2, t0
   532: 07830302                 vle8.v v6, (t2)
   536: b3835300                 add t2, t2, t0
   540: 87830302                 vle8.v v7, (t2)
   544: 138e0f06                 addi t3, t6, 0x60
   548: 07060e02                 vle8.v v12, (t3)
   552: 138e0f07                 addi t3, t6, 0x70
   556: 87060e02                 vle8.v v13, (t3)
   560: 5704e226                 vand.vv v8, v14, v4
   564: d73442a2                 vsrl.vi v9, v4, 4
   568: 5705c432                 vrgather.vv v10, v12, v8
   572: d785d432                 vrgather.vv v11, v13, v9
   576: 5705b52e                 vxor.vv v10, v11, v10
   580: 5700a02e                 vxor.vv v0, v10, v0
   584: 5784e226                 vand.vv v8, v14, v5
   588: d73452a2                 vsrl.vi v9, v5, 4
   592: 5705c432                 vrgather.vv v10, v12, v8
   596: d785d432                 vrgather.vv v11, v13, v9
   600: 5705b52e                 vxor.vv v10, v11, v10
   604: d780a02e                 vxor.vv v1, v10, v1
   608: 5704e326                 vand.vv v8, v14, v6
   612: d73462a2                 vsrl.vi v9, v6, 4
   616: 5705c432                 vrgather.vv v10, v12, v8
   620: d785d432                 vrgather.vv v11, v13, v9
   624: 5705b52e                 vxor.vv v10, v11, v10
   628: 5701a12e                 vxor.vv v2, v10, v2
   632: 5784e326                 vand.vv v8, v14, v7
   636: d73472a2                 vsrl.vi v9, v7, 4
   640: 5705c432                 vrgather.vv v10, v12, v8
   644: d785d432                 vrgather.vv v11, v13, v9
   648: 5705b52e                 vxor.vv v10, v11, v10
   652: d781a12e                 vxor.vv v3, v10, v3
   656: 83330502                 ld t2, 0x20(a0)
   660: b3836300                 add t2, t2, t1
   664: 07820302                 vle8.v v4, (t2)
   668: b3835300                 add t2, t2, t0
   672: 87820302                 vle8.v v5, (t2)
   676: b3835300                 add t2, t2, t0
   680: 07830302                 vle8.v v6, (t2)
   684: b3835300                 add t2, t2, t0
   688: 87830302                 vle8.v v7, (t2)
   692: 138e0f08                 addi t3, t6, 0x80
   696: 07060e02                 vle8.v v12, (t3)
   700: 138e0f09                 addi t3, t6, 0x90
   704: 87060e02                 vle8.v v13, (t3)
   708: 5704e226                 vand.vv v8, v14, v4
   712: d73442a2                 vsrl.vi v9, v4, 4
   716: 5705c432                 vrgather.vv v10, v12, v8
   720: d785d432                 vrgather.vv v11, v13, v9
   724: 5705b52e                 vxor.vv v10, v11, v10
   728: 5700a02e                 vxor.vv v0, v10, v0
   732: 5784e226                 vand.vv v8, v14, v5
   736: d73452a2                 vsrl.vi v9, v5, 4
   740: 5705c432                 vrgather.vv v10, v12, v8
   744: d785d432                 vrgather.vv v11, v13, v9
   748: 5705b52e                 vxor.vv v10, v11, v10
   752: d780a02e                 vxor.vv v1, v10, v1
   756: 5704e326                 vand.vv v8, v14, v6
   760: d73462a2                 vsrl.vi v9, v6, 4
   764: 5705c432                 vrgather.vv v10, v12, v8
   768: d785d432                 vrgather.vv v11, v13, v9
   772: 5705b52e                 vxor.vv v10, v11, v10
   776: 5701a12e                 vxor.vv v2, v10, v2
   780: 5784e326                 vand.vv v8, v14, v7
   784: d73472a2                 vsrl.vi v9, v7, 4
   788: 5705c432                 vrgather.vv v10, v12, v8
   792: d785d432                 vrgather.vv v11, v13, v9
   796: 5705b52e                 vxor.vv v10, v11, v10
   800: d781a12e                 vxor.vv v3, v10, v3
   804: 83338502                 ld t2, 0x28(a0)
   808: b3836300                 add t2, t2, t1
   812: 07820302                 vle8.v v4, (t2)
   816: b3835300                 add t2, t2, t0
   820: 87820302                 vle8.v v5, (t2)
   824: b3835300                 add t2, t2, t0
   828: 07830302                 vle8.v v6, (t2)
   832: b3835300                 add t2, t2, t0
   836: 87830302                 vle8.v v7, (t2)
   840: 138e0f0a                 addi t3, t6, 0xa0
   844: 07060e02                 vle8.v v12, (t3)
   848: 138e0f0b                 addi t3, t6, 0xb0
   852: 87060e02                 vle8.v v13, (t3)
   856: 5704e226                 vand.vv v8, v14, v4
   860: d73442a2                 vsrl.vi v9, v4, 4
   864: 5705c432                 vrgather.vv v10, v12, v8
   868: d785d432                 vrgather.vv v11, v13, v9
   872: 5705b52e                 vxor.vv v10, v11, v10
   876: 5700a02e                 vxor.vv v0, v10, v0
   880: 5784e226                 vand.vv v8, v14, v5
   884: d73452a2                 vsrl.vi v9, v5, 4
   888: 5705c432                 vrgather.vv v10, v12, v8
   892: d785d432                 vrgather.vv v11, v13, v9
   896: 5705b52e                 vxor.vv v10, v11, v10
   900: d780a02e                 vxor.vv v1, v10, v1
   904: 5704e326                 vand.vv v8, v14, v6
   908: d73462a2                 vsrl.vi v9, v6, 4
   912: 5705c432                 vrgather.vv v10, v12, v8
   916: d785d432                 vrgather.vv v11, v13, v9
   920: 5705b52e                 vxor.vv v10, v11, v10
   924: 5701a12e                 vxor.vv v2, v10, v2
   928: 5784e326                 vand.vv v8, v14, v7
   932: d73472a2                 vsrl.vi v9, v7, 4
   936: 5705c432                 vrgather.vv v10, v12, v8
   940: d785d432                 vrgather.vv v11, v13, v9
   944: 5705b52e                 vxor.vv v10, v11, v10
   948: d781a12e                 vxor.vv v3, v10, v3
   952: 83330503                 ld t2, 0x30(a0)
   956: b3836300                 add t2, t2, t1
   960: 07820302                 vle8.v v4, (t2)
   964: b3835300                 add t2, t2, t0
   968: 87820302                 vle8.v v5, (t2)
   972: b3835300                 add t2, t2, t0
   976: 07830302                 vle8.v v6, (t2)
   980: b3835300                 add t2, t2, t0
   984: 87830302                 vle8.v v7, (t2)
   988: 138e0f0c                 addi t3, t6, 0xc0
   992: 07060e02                 vle8.v v12, (t3)
   996: 138e0f0d                 addi t3, t6, 0xd0
  1000: 87060e02                 vle8.v v13, (t3)
  1004: 5704e226                 vand.vv v8, v14, v4
  1008: d73442a2                 vsrl.vi v9, v4, 4
  1012: 5705c432                 vrgather.vv v10, v12, v8
  1016: d785d432                 vrgather.vv v11, v13, v9
  1020: 5705b52e                 vxor.vv v10, v11, v10
  1024: 5700a02e                 vxor.vv v0, v10, v0
  1028: 5784e226                 vand.vv v8, v14, v5
  1032: d73452a2                 vsrl.vi v9, v5, 4
  1036: 5705c432                 vrgather.vv v10, v12, v8
  1040: d785d432                 vrgather.vv v11, v13, v9
  1044: 5705b52e                 vxor.vv v10, v11, v10
  1048: d780a02e                 vxor.vv v1, v10, v1
  1052: 5704e326                 vand.vv v8, v14, v6
  1056: d73462a2                 vsrl.vi v9, v6, 4
  1060: 5705c432                 vrgather.vv v10, v12, v8
  1064: d785d432                 vrgather.vv v11, v13, v9
  1068: 5705b52e                 vxor.vv v10, v11, v10
  1072: 5701a12e                 vxor.vv v2, v10, v2
  1076: 5784e326                 vand.vv v8, v14, v7
  1080: d73472a2                 vsrl.vi v9, v7, 4
  1084: 5705c432                 vrgather.vv v10, v12, v8
  1088: d785d432                 vrgather.vv v11, v13, v9
  1092: 5705b52e                 vxor.vv v10, v11, v10
  1096: d781a12e                 vxor.vv v3, v10, v3
  1100: 83338503                 ld t2, 0x38(a0)
  1104: b3836300                 add t2, t2, t1
  1108: 07820302                 vle8.v v4, (t2)
  1112: b3835300                 add t2, t2, t0
  1116: 87820302                 vle8.v v5, (t2)
  1120: b3835300                 add t2, t2, t0
  1124: 07830302                 vle8.v v6, (t2)
  1128: b3835300                 add t2, t2, t0
  1132: 87830302                 vle8.v v7, (t2)
  1136: 138e0f0e                 addi t3, t6, 0xe0
  1140: 07060e02                 vle8.v v12, (t3)
  1144: 138e0f0f                 addi t3, t6, 0xf0
  1148: 87060e02                 vle8.v v13, (t3)
  1152: 5704e226                 vand.vv v8, v14, v4
  1156: d73442a2                 vsrl.vi v9, v4, 4
  1160: 5705c432                 vrgather.vv v10, v12, v8
  1164: d785d432                 vrgather.vv v11, v13, v9
  1168: 5705b52e                 vxor.vv v10, v11, v10
  1172: 5700a02e                 vxor.vv v0, v10, v0
  1176: 5784e226                 vand.vv v8, v14, v5
  1180: d73452a2                 vsrl.vi v9, v5, 4
  1184: 5705c432                 vrgather.vv v10, v12, v8
  1188: d785d432                 vrgather.vv v11, v13, v9
  1192: 5705b52e                 vxor.vv v10, v11, v10
  1196: d780a02e                 vxor.vv v1, v10, v1
  1200: 5704e326                 vand.vv v8, v14, v6
  1204: d73462a2                 vsrl.vi v9, v6, 4
  1208: 5705c432                 vrgather.vv v10, v12, v8
  1212: d785d432                 vrgather.vv v11, v13, v9
  1216: 5705b52e                 vxor.vv v10, v11, v10
  1220: 5701a12e                 vxor.vv v2, v10, v2
  1224: 5784e326                 vand.vv v8, v14, v7
  1228: d73472a2                 vsrl.vi v9, v7, 4
  1232: 5705c432                 vrgather.vv v10, v12, v8
  1236: d785d432                 vrgather.vv v11, v13, v9
  1240: 5705b52e                 vxor.vv v10, v11, v10
  1244: d781a12e                 vxor.vv v3, v10, v3
  1248: 83330504                 ld t2, 0x40(a0)
  1252: b3836300                 add t2, t2, t1
  1256: 07820302                 vle8.v v4, (t2)
  1260: b3835300                 add t2, t2, t0
  1264: 87820302                 vle8.v v5, (t2)
  1268: b3835300                 add t2, t2, t0
  1272: 07830302                 vle8.v v6, (t2)
  1276: b3835300                 add t2, t2, t0
  1280: 87830302                 vle8.v v7, (t2)
  1284: 138e0f10                 addi t3, t6, 0x100
  1288: 07060e02                 vle8.v v12, (t3)
  1292: 138e0f11                 addi t3, t6, 0x110
  1296: 87060e02                 vle8.v v13, (t3)
  1300: 5704e226                 vand.vv v8, v14, v4
  1304: d73442a2                 vsrl.vi v9, v4, 4
  1308: 5705c432                 vrgather.vv v10, v12, v8
  1312: d785d432                 vrgather.vv v11, v13, v9
  1316: 5705b52e                 vxor.vv v10, v11, v10
  1320: 5700a02e                 vxor.vv v0, v10, v0
  1324: 5784e226                 vand.vv v8, v14, v5
  1328: d73452a2                 vsrl.vi v9, v5, 4
  1332: 5705c432                 vrgather.vv v10, v12, v8
  1336: d785d432                 vrgather.vv v11, v13, v9
  1340: 5705b52e                 vxor.vv v10, v11, v10
  1344: d780a02e                 vxor.vv v1, v10, v1
  1348: 5704e326                 vand.vv v8, v14, v6
  1352: d73462a2                 vsrl.vi v9, v6, 4
  1356: 5705c432                 vrgather.vv v10, v12, v8
  1360: d785d432                 vrgather.vv v11, v13, v9
  1364: 5705b52e                 vxor.vv v10, v11, v10
  1368: 5701a12e                 vxor.vv v2, v10, v2
  1372: 5784e326                 vand.vv v8, v14, v7
  1376: d73472a2                 vsrl.vi v9, v7, 4
  1380: 5705c432                 vrgather.vv v10, v12, v8
  1384: d785d432                 vrgather.vv v11, v13, v9
  1388: 5705b52e                 vxor.vv v10, v11, v10
  1392: d781a12e                 vxor.vv v3, v10, v3
  1396: 83338504                 ld t2, 0x48(a0)
  1400: b3836300                 add t2, t2, t1
  1404: 07820302                 vle8.v v4, (t2)
  1408: b3835300                 add t2, t2, t0
  1412: 87820302                 vle8.v v5, (t2)
  1416: b3835300                 add t2, t2, t0
  1420: 07830302                 vle8.v v6, (t2)
  1424: b3835300                 add t2, t2, t0
  1428: 87830302                 vle8.v v7, (t2)
  1432: 138e0f12                 addi t3, t6, 0x120
  1436: 07060e02                 vle8.v v12, (t3)
  1440: 138e0f13                 addi t3, t6, 0x130
  1444: 87060e02                 vle8.v v13, (t3)
  1448: 5704e226                 vand.vv v8, v14, v4
  1452: d73442a2                 vsrl.vi v9, v4, 4
  1456: 5705c432                 vrgather.vv v10, v12, v8
  1460: d785d432                 vrgather.vv v11, v13, v9
  1464: 5705b52e                 vxor.vv v10, v11, v10
  1468: 5700a02e                 vxor.vv v0, v10, v0
  1472: 5784e226                 vand.vv v8, v14, v5
  1476: d73452a2                 vsrl.vi v9, v5, 4
  1480: 5705c432                 vrgather.vv v10, v12, v8
  1484: d785d432                 vrgather.vv v11, v13, v9
  1488: 5705b52e                 vxor.vv v10, v11, v10
  1492: d780a02e                 vxor.vv v1, v10, v1
  1496: 5704e326                 vand.vv v8, v14, v6
  1500: d73462a2                 vsrl.vi v9, v6, 4
  1504: 5705c432                 vrgather.vv v10, v12, v8
  1508: d785d432                 vrgather.vv v11, v13, v9
  1512: 5705b52e                 vxor.vv v10, v11, v10
  1516: 5701a12e                 vxor.vv v2, v10, v2
  1520: 5784e326                 vand.vv v8, v14, v7
  1524: d73472a2                 vsrl.vi v9, v7, 4
  1528: 5705c432                 vrgather.vv v10, v12, v8
  1532: d785d432                 vrgather.vv v11, v13, v9
  1536: 5705b52e                 vxor.vv v10, v11, v10
  1540: d781a12e                 vxor.vv v3, v10, v3
  1544: 83330505                 ld t2, 0x50(a0)
  1548: b3836300                 add t2, t2, t1
  1552: 07820302                 vle8.v v4, (t2)
  1556: b3835300                 add t2, t2, t0
  1560: 87820302                 vle8.v v5, (t2)
  1564: b3835300                 add t2, t2, t0
  1568: 07830302                 vle8.v v6, (t2)
  1572: b3835300                 add t2, t2, t0
  1576: 87830302                 vle8.v v7, (t2)
  1580: 138e0f14                 addi t3, t6, 0x140
  1584: 07060e02                 vle8.v v12, (t3)
  1588: 138e0f15                 addi t3, t6, 0x150
  1592: 87060e02                 vle8.v v13, (t3)
  1596: 5704e226                 vand.vv v8, v14, v4
  1600: d73442a2                 vsrl.vi v9, v4, 4
  1604: 5705c432                 vrgather.vv v10, v12, v8
  1608: d785d432                 vrgather.vv v11, v13, v9
  1612: 5705b52e                 vxor.vv v10, v11, v10
  1616: 5700a02e                 vxor.vv v0, v10, v0
  1620: 5784e226                 vand.vv v8, v14, v5
  1624: d73452a2                 vsrl.vi v9, v5, 4
  1628: 5705c432                 vrgather.vv v10, v12, v8
  1632: d785d432                 vrgather.vv v11, v13, v9
  1636: 5705b52e                 vxor.vv v10, v11, v10
  1640: d780a02e                 vxor.vv v1, v10, v1
  1644: 5704e326                 vand.vv v8, v14, v6
  1648: d73462a2                 vsrl.vi v9, v6, 4
  1652: 5705c432                 vrgather.vv v10, v12, v8
  1656: d785d432                 vrgather.vv v11, v13, v9
  1660: 5705b52e                 vxor.vv v10, v11, v10
  1664: 5701a12e                 vxor.vv v2, v10, v2
  1668: 5784e326                 vand.vv v8, v14, v7
  1672: d73472a2                 vsrl.vi v9, v7, 4
  1676: 5705c432                 vrgather.vv v10, v12, v8
  1680: d785d432                 vrgather.vv v11, v13, v9
  1684: 5705b52e                 vxor.vv v10, v11, v10
  1688: d781a12e                 vxor.vv v3, v10, v3
  1692: 83338505                 ld t2, 0x58(a0)
  1696: b3836300                 add t2, t2, t1
  1700: 07820302                 vle8.v v4, (t2)
  1704: b3835300                 add t2, t2, t0
  1708: 87820302                 vle8.v v5, (t2)
  1712: b3835300                 add t2, t2, t0
  1716: 07830302                 vle8.v v6, (t2)
  1720: b3835300                 add t2, t2, t0
  1724: 87830302                 vle8.v v7, (t2)
  1728: 5700402e                 vxor.vv v0, v4, v0
  1732: d780502e                 vxor.vv v1, v5, v1
  1736: 5701612e                 vxor.vv v2, v6, v2
  1740: d781712e                 vxor.vv v3, v7, v3
  1744: 83b30500                 ld t2, 0(a1)
  1748: b3836300                 add t2, t2, t1
  1752: 27800302                 vse8.v v0, (t2)
  1756: b3835300                 add t2, t2, t0
  1760: a7800302                 vse8.v v1, (t2)
  1764: b3835300                 add t2, t2, t0
  1768: 27810302                 vse8.v v2, (t2)
  1772: b3835300                 add t2, t2, t0
  1776: a7810302                 vse8.v v3, (t2)
  1780: 3303d300                 add t1, t1, a3
  1784: 1307f7ff                 addi a4, a4, -1
  1788: e3120794                 bnez a4, 0x40
  1792: 33066640                 sub a2, a2, t1
  1796: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
  1800: 63940200                 bnez t0, 0x710
  1804: 6f00c02e                 j 0x9f8
  1808: 83330500                 ld t2, 0(a0)
  1812: b3836300                 add t2, t2, t1
  1816: 07820302                 vle8.v v4, (t2)
  1820: 138f0200                 mv t5, t0
  1824: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1828: 138e0f00                 mv t3, t6
  1832: 07060e02                 vle8.v v12, (t3)
  1836: 138e0f01                 addi t3, t6, 0x10
  1840: 87060e02                 vle8.v v13, (t3)
  1844: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1848: 5704e226                 vand.vv v8, v14, v4
  1852: d73442a2                 vsrl.vi v9, v4, 4
  1856: 5705c432                 vrgather.vv v10, v12, v8
  1860: d785d432                 vrgather.vv v11, v13, v9
  1864: 5705b52e                 vxor.vv v10, v11, v10
  1868: 5700055e                 vmv.v.v v0, v10
  1872: 83338500                 ld t2, 8(a0)
  1876: b3836300                 add t2, t2, t1
  1880: 07820302                 vle8.v v4, (t2)
  1884: 138f0200                 mv t5, t0
  1888: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1892: 138e0f02                 addi t3, t6, 0x20
  1896: 07060e02                 vle8.v v12, (t3)
  1900: 138e0f03                 addi t3, t6, 0x30
  1904: 87060e02                 vle8.v v13, (t3)
  1908: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1912: 5704e226                 vand.vv v8, v14, v4
  1916: d73442a2                 vsrl.vi v9, v4, 4
  1920: 5705c432                 vrgather.vv v10, v12, v8
  1924: d785d432                 vrgather.vv v11, v13, v9
  1928: 5705b52e                 vxor.vv v10, v11, v10
  1932: 5700a02e                 vxor.vv v0, v10, v0
  1936: 83330501                 ld t2, 0x10(a0)
  1940: b3836300                 add t2, t2, t1
  1944: 07820302                 vle8.v v4, (t2)
  1948: 138f0200                 mv t5, t0
  1952: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1956: 138e0f04                 addi t3, t6, 0x40
  1960: 07060e02                 vle8.v v12, (t3)
  1964: 138e0f05                 addi t3, t6, 0x50
  1968: 87060e02                 vle8.v v13, (t3)
  1972: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1976: 5704e226                 vand.vv v8, v14, v4
  1980: d73442a2                 vsrl.vi v9, v4, 4
  1984: 5705c432                 vrgather.vv v10, v12, v8
  1988: d785d432                 vrgather.vv v11, v13, v9
  1992: 5705b52e                 vxor.vv v10, v11, v10
  1996: 5700a02e                 vxor.vv v0, v10, v0
  2000: 83338501                 ld t2, 0x18(a0)
  2004: b3836300                 add t2, t2, t1
  2008: 07820302                 vle8.v v4, (t2)
  2012: 138f0200                 mv t5, t0
  2016: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2020: 138e0f06                 addi t3, t6, 0x60
  2024: 07060e02                 vle8.v v12, (t3)
  2028: 138e0f07                 addi t3, t6, 0x70
  2032: 87060e02                 vle8.v v13, (t3)
  2036: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2040: 5704e226                 vand.vv v8, v14, v4
  2044: d73442a2                 vsrl.vi v9, v4, 4
  2048: 5705c432                 vrgather.vv v10, v12, v8
  2052: d785d432                 vrgather.vv v11, v13, v9
  2056: 5705b52e                 vxor.vv v10, v11, v10
  2060: 5700a02e                 vxor.vv v0, v10, v0
  2064: 83330502                 ld t2, 0x20(a0)
  2068: b3836300                 add t2, t2, t1
  2072: 07820302                 vle8.v v4, (t2)
  2076: 138f0200                 mv t5, t0
  2080: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2084: 138e0f08                 addi t3, t6, 0x80
  2088: 07060e02                 vle8.v v12, (t3)
  2092: 138e0f09                 addi t3, t6, 0x90
  2096: 87060e02                 vle8.v v13, (t3)
  2100: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2104: 5704e226                 vand.vv v8, v14, v4
  2108: d73442a2                 vsrl.vi v9, v4, 4
  2112: 5705c432                 vrgather.vv v10, v12, v8
  2116: d785d432                 vrgather.vv v11, v13, v9
  2120: 5705b52e                 vxor.vv v10, v11, v10
  2124: 5700a02e                 vxor.vv v0, v10, v0
  2128: 83338502                 ld t2, 0x28(a0)
  2132: b3836300                 add t2, t2, t1
  2136: 07820302                 vle8.v v4, (t2)
  2140: 138f0200                 mv t5, t0
  2144: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2148: 138e0f0a                 addi t3, t6, 0xa0
  2152: 07060e02                 vle8.v v12, (t3)
  2156: 138e0f0b                 addi t3, t6, 0xb0
  2160: 87060e02                 vle8.v v13, (t3)
  2164: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2168: 5704e226                 vand.vv v8, v14, v4
  2172: d73442a2                 vsrl.vi v9, v4, 4
  2176: 5705c432                 vrgather.vv v10, v12, v8
  2180: d785d432                 vrgather.vv v11, v13, v9
  2184: 5705b52e                 vxor.vv v10, v11, v10
  2188: 5700a02e                 vxor.vv v0, v10, v0
  2192: 83330503                 ld t2, 0x30(a0)
  2196: b3836300                 add t2, t2, t1
  2200: 07820302                 vle8.v v4, (t2)
  2204: 138f0200                 mv t5, t0
  2208: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2212: 138e0f0c                 addi t3, t6, 0xc0
  2216: 07060e02                 vle8.v v12, (t3)
  2220: 138e0f0d                 addi t3, t6, 0xd0
  2224: 87060e02                 vle8.v v13, (t3)
  2228: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2232: 5704e226                 vand.vv v8, v14, v4
  2236: d73442a2                 vsrl.vi v9, v4, 4
  2240: 5705c432                 vrgather.vv v10, v12, v8
  2244: d785d432                 vrgather.vv v11, v13, v9
  2248: 5705b52e                 vxor.vv v10, v11, v10
  2252: 5700a02e                 vxor.vv v0, v10, v0
  2256: 83338503                 ld t2, 0x38(a0)
  2260: b3836300                 add t2, t2, t1
  2264: 07820302                 vle8.v v4, (t2)
  2268: 138f0200                 mv t5, t0
  2272: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2276: 138e0f0e                 addi t3, t6, 0xe0
  2280: 07060e02                 vle8.v v12, (t3)
  2284: 138e0f0f                 addi t3, t6, 0xf0
  2288: 87060e02                 vle8.v v13, (t3)
  2292: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2296: 5704e226                 vand.vv v8, v14, v4
  2300: d73442a2                 vsrl.vi v9, v4, 4
  2304: 5705c432                 vrgather.vv v10, v12, v8
  2308: d785d432                 vrgather.vv v11, v13, v9
  2312: 5705b52e                 vxor.vv v10, v11, v10
  2316: 5700a02e                 vxor.vv v0, v10, v0
  2320: 83330504                 ld t2, 0x40(a0)
  2324: b3836300                 add t2, t2, t1
  2328: 07820302                 vle8.v v4, (t2)
  2332: 138f0200                 mv t5, t0
  2336: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2340: 138e0f10                 addi t3, t6, 0x100
  2344: 07060e02                 vle8.v v12, (t3)
  2348: 138e0f11                 addi t3, t6, 0x110
  2352: 87060e02                 vle8.v v13, (t3)
  2356: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2360: 5704e226                 vand.vv v8, v14, v4
  2364: d73442a2                 vsrl.vi v9, v4, 4
  2368: 5705c432                 vrgather.vv v10, v12, v8
  2372: d785d432                 vrgather.vv v11, v13, v9
  2376: 5705b52e                 vxor.vv v10, v11, v10
  2380: 5700a02e                 vxor.vv v0, v10, v0
  2384: 83338504                 ld t2, 0x48(a0)
  2388: b3836300                 add t2, t2, t1
  2392: 07820302                 vle8.v v4, (t2)
  2396: 138f0200                 mv t5, t0
  2400: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2404: 138e0f12                 addi t3, t6, 0x120
  2408: 07060e02                 vle8.v v12, (t3)
  2412: 138e0f13                 addi t3, t6, 0x130
  2416: 87060e02                 vle8.v v13, (t3)
  2420: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2424: 5704e226                 vand.vv v8, v14, v4
  2428: d73442a2                 vsrl.vi v9, v4, 4
  2432: 5705c432                 vrgather.vv v10, v12, v8
  2436: d785d432                 vrgather.vv v11, v13, v9
  2440: 5705b52e                 vxor.vv v10, v11, v10
  2444: 5700a02e                 vxor.vv v0, v10, v0
  2448: 83330505                 ld t2, 0x50(a0)
  2452: b3836300                 add t2, t2, t1
  2456: 07820302                 vle8.v v4, (t2)
  2460: 138f0200                 mv t5, t0
  2464: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2468: 138e0f14                 addi t3, t6, 0x140
  2472: 07060e02                 vle8.v v12, (t3)
  2476: 138e0f15                 addi t3, t6, 0x150
  2480: 87060e02                 vle8.v v13, (t3)
  2484: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2488: 5704e226                 vand.vv v8, v14, v4
  2492: d73442a2                 vsrl.vi v9, v4, 4
  2496: 5705c432                 vrgather.vv v10, v12, v8
  2500: d785d432                 vrgather.vv v11, v13, v9
  2504: 5705b52e                 vxor.vv v10, v11, v10
  2508: 5700a02e                 vxor.vv v0, v10, v0
  2512: 83338505                 ld t2, 0x58(a0)
  2516: b3836300                 add t2, t2, t1
  2520: 07820302                 vle8.v v4, (t2)
  2524: 5700402e                 vxor.vv v0, v4, v0
  2528: 83b30500                 ld t2, 0(a1)
  2532: b3836300                 add t2, t2, t1
  2536: 27800302                 vse8.v v0, (t2)
  2540: 33035300                 add t1, t1, t0
  2544: 33065640                 sub a2, a2, t0
  2548: 6ff01fd1                 j 0x704
  2552: 67800000                 ret 
  2556: 0000                     unimp 
  2558: 0000                     unimp 
