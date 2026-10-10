; config_t3_2122 target=rvv twin=baseline topo=lrc(8,2,2)/rec{0} op=recover U=4
; code_size=1712 twins=[] uops_per_byte=5.296875
     0: 63140600                 bnez a2, 8
     4: 6f00406a                 j 0x6a8
     8: 970f0000                 auipc t6, 0
    12: 938f8f6a                 addi t6, t6, 0x6a8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 57470e5e                 vmv.v.x v14, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93962200                 slli a3, t0, 2
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f004048                 j 0x4b4
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
  1136: 5700402e                 vxor.vv v0, v4, v0
  1140: d780502e                 vxor.vv v1, v5, v1
  1144: 5701612e                 vxor.vv v2, v6, v2
  1148: d781712e                 vxor.vv v3, v7, v3
  1152: 83b30500                 ld t2, 0(a1)
  1156: b3836300                 add t2, t2, t1
  1160: 27800302                 vse8.v v0, (t2)
  1164: b3835300                 add t2, t2, t0
  1168: a7800302                 vse8.v v1, (t2)
  1172: b3835300                 add t2, t2, t0
  1176: 27810302                 vse8.v v2, (t2)
  1180: b3835300                 add t2, t2, t0
  1184: a7810302                 vse8.v v3, (t2)
  1188: 3303d300                 add t1, t1, a3
  1192: 1307f7ff                 addi a4, a4, -1
  1196: e31a07b8                 bnez a4, 0x40
  1200: 33066640                 sub a2, a2, t1
  1204: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
  1208: 63940200                 bnez t0, 0x4c0
  1212: 6f00c01e                 j 0x6a8
  1216: 83330500                 ld t2, 0(a0)
  1220: b3836300                 add t2, t2, t1
  1224: 07820302                 vle8.v v4, (t2)
  1228: 138f0200                 mv t5, t0
  1232: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1236: 138e0f00                 mv t3, t6
  1240: 07060e02                 vle8.v v12, (t3)
  1244: 138e0f01                 addi t3, t6, 0x10
  1248: 87060e02                 vle8.v v13, (t3)
  1252: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1256: 5704e226                 vand.vv v8, v14, v4
  1260: d73442a2                 vsrl.vi v9, v4, 4
  1264: 5705c432                 vrgather.vv v10, v12, v8
  1268: d785d432                 vrgather.vv v11, v13, v9
  1272: 5705b52e                 vxor.vv v10, v11, v10
  1276: 5700055e                 vmv.v.v v0, v10
  1280: 83338500                 ld t2, 8(a0)
  1284: b3836300                 add t2, t2, t1
  1288: 07820302                 vle8.v v4, (t2)
  1292: 138f0200                 mv t5, t0
  1296: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1300: 138e0f02                 addi t3, t6, 0x20
  1304: 07060e02                 vle8.v v12, (t3)
  1308: 138e0f03                 addi t3, t6, 0x30
  1312: 87060e02                 vle8.v v13, (t3)
  1316: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1320: 5704e226                 vand.vv v8, v14, v4
  1324: d73442a2                 vsrl.vi v9, v4, 4
  1328: 5705c432                 vrgather.vv v10, v12, v8
  1332: d785d432                 vrgather.vv v11, v13, v9
  1336: 5705b52e                 vxor.vv v10, v11, v10
  1340: 5700a02e                 vxor.vv v0, v10, v0
  1344: 83330501                 ld t2, 0x10(a0)
  1348: b3836300                 add t2, t2, t1
  1352: 07820302                 vle8.v v4, (t2)
  1356: 138f0200                 mv t5, t0
  1360: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1364: 138e0f04                 addi t3, t6, 0x40
  1368: 07060e02                 vle8.v v12, (t3)
  1372: 138e0f05                 addi t3, t6, 0x50
  1376: 87060e02                 vle8.v v13, (t3)
  1380: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1384: 5704e226                 vand.vv v8, v14, v4
  1388: d73442a2                 vsrl.vi v9, v4, 4
  1392: 5705c432                 vrgather.vv v10, v12, v8
  1396: d785d432                 vrgather.vv v11, v13, v9
  1400: 5705b52e                 vxor.vv v10, v11, v10
  1404: 5700a02e                 vxor.vv v0, v10, v0
  1408: 83338501                 ld t2, 0x18(a0)
  1412: b3836300                 add t2, t2, t1
  1416: 07820302                 vle8.v v4, (t2)
  1420: 138f0200                 mv t5, t0
  1424: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1428: 138e0f06                 addi t3, t6, 0x60
  1432: 07060e02                 vle8.v v12, (t3)
  1436: 138e0f07                 addi t3, t6, 0x70
  1440: 87060e02                 vle8.v v13, (t3)
  1444: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1448: 5704e226                 vand.vv v8, v14, v4
  1452: d73442a2                 vsrl.vi v9, v4, 4
  1456: 5705c432                 vrgather.vv v10, v12, v8
  1460: d785d432                 vrgather.vv v11, v13, v9
  1464: 5705b52e                 vxor.vv v10, v11, v10
  1468: 5700a02e                 vxor.vv v0, v10, v0
  1472: 83330502                 ld t2, 0x20(a0)
  1476: b3836300                 add t2, t2, t1
  1480: 07820302                 vle8.v v4, (t2)
  1484: 138f0200                 mv t5, t0
  1488: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1492: 138e0f08                 addi t3, t6, 0x80
  1496: 07060e02                 vle8.v v12, (t3)
  1500: 138e0f09                 addi t3, t6, 0x90
  1504: 87060e02                 vle8.v v13, (t3)
  1508: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1512: 5704e226                 vand.vv v8, v14, v4
  1516: d73442a2                 vsrl.vi v9, v4, 4
  1520: 5705c432                 vrgather.vv v10, v12, v8
  1524: d785d432                 vrgather.vv v11, v13, v9
  1528: 5705b52e                 vxor.vv v10, v11, v10
  1532: 5700a02e                 vxor.vv v0, v10, v0
  1536: 83338502                 ld t2, 0x28(a0)
  1540: b3836300                 add t2, t2, t1
  1544: 07820302                 vle8.v v4, (t2)
  1548: 138f0200                 mv t5, t0
  1552: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1556: 138e0f0a                 addi t3, t6, 0xa0
  1560: 07060e02                 vle8.v v12, (t3)
  1564: 138e0f0b                 addi t3, t6, 0xb0
  1568: 87060e02                 vle8.v v13, (t3)
  1572: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1576: 5704e226                 vand.vv v8, v14, v4
  1580: d73442a2                 vsrl.vi v9, v4, 4
  1584: 5705c432                 vrgather.vv v10, v12, v8
  1588: d785d432                 vrgather.vv v11, v13, v9
  1592: 5705b52e                 vxor.vv v10, v11, v10
  1596: 5700a02e                 vxor.vv v0, v10, v0
  1600: 83330503                 ld t2, 0x30(a0)
  1604: b3836300                 add t2, t2, t1
  1608: 07820302                 vle8.v v4, (t2)
  1612: 138f0200                 mv t5, t0
  1616: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1620: 138e0f0c                 addi t3, t6, 0xc0
  1624: 07060e02                 vle8.v v12, (t3)
  1628: 138e0f0d                 addi t3, t6, 0xd0
  1632: 87060e02                 vle8.v v13, (t3)
  1636: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1640: 5704e226                 vand.vv v8, v14, v4
  1644: d73442a2                 vsrl.vi v9, v4, 4
  1648: 5705c432                 vrgather.vv v10, v12, v8
  1652: d785d432                 vrgather.vv v11, v13, v9
  1656: 5705b52e                 vxor.vv v10, v11, v10
  1660: 5700a02e                 vxor.vv v0, v10, v0
  1664: 83338503                 ld t2, 0x38(a0)
  1668: b3836300                 add t2, t2, t1
  1672: 07820302                 vle8.v v4, (t2)
  1676: 5700402e                 vxor.vv v0, v4, v0
  1680: 83b30500                 ld t2, 0(a1)
  1684: b3836300                 add t2, t2, t1
  1688: 27800302                 vse8.v v0, (t2)
  1692: 33035300                 add t1, t1, t0
  1696: 33065640                 sub a2, a2, t0
  1700: 6ff01fe1                 j 0x4b4
  1704: 67800000                 ret 
  1708: 0000                     unimp 
  1710: 0000                     unimp 
