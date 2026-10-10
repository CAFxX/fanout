; config_t2_2047 target=rvv twin=baseline topo=rs(6,3)/enc op=encode U=4
; code_size=2352 twins=[] uops_per_byte=7.765625
     0: 63140600                 bnez a2, 8
     4: 6f00d011                 j 0x920
     8: 971f0000                 auipc t6, 1
    12: 938f8f92                 addi t6, t6, -0x6d8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574b0e5e                 vmv.v.x v22, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93962200                 slli a3, t0, 2
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f004064                 j 0x674
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07860302                 vle8.v v12, (t2)
    76: b3835300                 add t2, t2, t0
    80: 87860302                 vle8.v v13, (t2)
    84: b3835300                 add t2, t2, t0
    88: 07870302                 vle8.v v14, (t2)
    92: b3835300                 add t2, t2, t0
    96: 87870302                 vle8.v v15, (t2)
   100: 5700065e                 vmv.v.v v0, v12
   104: d780065e                 vmv.v.v v1, v13
   108: 5701075e                 vmv.v.v v2, v14
   112: d781075e                 vmv.v.v v3, v15
   116: 138e0f00                 mv t3, t6
   120: 070a0e02                 vle8.v v20, (t3)
   124: 138e0f01                 addi t3, t6, 0x10
   128: 870a0e02                 vle8.v v21, (t3)
   132: 57086627                 vand.vv v16, v22, v12
   136: d738c2a2                 vsrl.vi v17, v12, 4
   140: 57094833                 vrgather.vv v18, v20, v16
   144: d7895833                 vrgather.vv v19, v21, v17
   148: 5709392f                 vxor.vv v18, v19, v18
   152: 5702095e                 vmv.v.v v4, v18
   156: 57886627                 vand.vv v16, v22, v13
   160: d738d2a2                 vsrl.vi v17, v13, 4
   164: 57094833                 vrgather.vv v18, v20, v16
   168: d7895833                 vrgather.vv v19, v21, v17
   172: 5709392f                 vxor.vv v18, v19, v18
   176: d702095e                 vmv.v.v v5, v18
   180: 57086727                 vand.vv v16, v22, v14
   184: d738e2a2                 vsrl.vi v17, v14, 4
   188: 57094833                 vrgather.vv v18, v20, v16
   192: d7895833                 vrgather.vv v19, v21, v17
   196: 5709392f                 vxor.vv v18, v19, v18
   200: 5703095e                 vmv.v.v v6, v18
   204: 57886727                 vand.vv v16, v22, v15
   208: d738f2a2                 vsrl.vi v17, v15, 4
   212: 57094833                 vrgather.vv v18, v20, v16
   216: d7895833                 vrgather.vv v19, v21, v17
   220: 5709392f                 vxor.vv v18, v19, v18
   224: d703095e                 vmv.v.v v7, v18
   228: 138e0f02                 addi t3, t6, 0x20
   232: 070a0e02                 vle8.v v20, (t3)
   236: 138e0f03                 addi t3, t6, 0x30
   240: 870a0e02                 vle8.v v21, (t3)
   244: 57086627                 vand.vv v16, v22, v12
   248: d738c2a2                 vsrl.vi v17, v12, 4
   252: 57094833                 vrgather.vv v18, v20, v16
   256: d7895833                 vrgather.vv v19, v21, v17
   260: 5709392f                 vxor.vv v18, v19, v18
   264: 5704095e                 vmv.v.v v8, v18
   268: 57886627                 vand.vv v16, v22, v13
   272: d738d2a2                 vsrl.vi v17, v13, 4
   276: 57094833                 vrgather.vv v18, v20, v16
   280: d7895833                 vrgather.vv v19, v21, v17
   284: 5709392f                 vxor.vv v18, v19, v18
   288: d704095e                 vmv.v.v v9, v18
   292: 57086727                 vand.vv v16, v22, v14
   296: d738e2a2                 vsrl.vi v17, v14, 4
   300: 57094833                 vrgather.vv v18, v20, v16
   304: d7895833                 vrgather.vv v19, v21, v17
   308: 5709392f                 vxor.vv v18, v19, v18
   312: 5705095e                 vmv.v.v v10, v18
   316: 57886727                 vand.vv v16, v22, v15
   320: d738f2a2                 vsrl.vi v17, v15, 4
   324: 57094833                 vrgather.vv v18, v20, v16
   328: d7895833                 vrgather.vv v19, v21, v17
   332: 5709392f                 vxor.vv v18, v19, v18
   336: d705095e                 vmv.v.v v11, v18
   340: 83338500                 ld t2, 8(a0)
   344: b3836300                 add t2, t2, t1
   348: 07860302                 vle8.v v12, (t2)
   352: b3835300                 add t2, t2, t0
   356: 87860302                 vle8.v v13, (t2)
   360: b3835300                 add t2, t2, t0
   364: 07870302                 vle8.v v14, (t2)
   368: b3835300                 add t2, t2, t0
   372: 87870302                 vle8.v v15, (t2)
   376: 5700c02e                 vxor.vv v0, v12, v0
   380: d780d02e                 vxor.vv v1, v13, v1
   384: 5701e12e                 vxor.vv v2, v14, v2
   388: d781f12e                 vxor.vv v3, v15, v3
   392: 138e0f02                 addi t3, t6, 0x20
   396: 070a0e02                 vle8.v v20, (t3)
   400: 138e0f03                 addi t3, t6, 0x30
   404: 870a0e02                 vle8.v v21, (t3)
   408: 57086627                 vand.vv v16, v22, v12
   412: d738c2a2                 vsrl.vi v17, v12, 4
   416: 57094833                 vrgather.vv v18, v20, v16
   420: d7895833                 vrgather.vv v19, v21, v17
   424: 5709392f                 vxor.vv v18, v19, v18
   428: 5702222f                 vxor.vv v4, v18, v4
   432: 57886627                 vand.vv v16, v22, v13
   436: d738d2a2                 vsrl.vi v17, v13, 4
   440: 57094833                 vrgather.vv v18, v20, v16
   444: d7895833                 vrgather.vv v19, v21, v17
   448: 5709392f                 vxor.vv v18, v19, v18
   452: d782222f                 vxor.vv v5, v18, v5
   456: 57086727                 vand.vv v16, v22, v14
   460: d738e2a2                 vsrl.vi v17, v14, 4
   464: 57094833                 vrgather.vv v18, v20, v16
   468: d7895833                 vrgather.vv v19, v21, v17
   472: 5709392f                 vxor.vv v18, v19, v18
   476: 5703232f                 vxor.vv v6, v18, v6
   480: 57886727                 vand.vv v16, v22, v15
   484: d738f2a2                 vsrl.vi v17, v15, 4
   488: 57094833                 vrgather.vv v18, v20, v16
   492: d7895833                 vrgather.vv v19, v21, v17
   496: 5709392f                 vxor.vv v18, v19, v18
   500: d783232f                 vxor.vv v7, v18, v7
   504: 138e0f04                 addi t3, t6, 0x40
   508: 070a0e02                 vle8.v v20, (t3)
   512: 138e0f05                 addi t3, t6, 0x50
   516: 870a0e02                 vle8.v v21, (t3)
   520: 57086627                 vand.vv v16, v22, v12
   524: d738c2a2                 vsrl.vi v17, v12, 4
   528: 57094833                 vrgather.vv v18, v20, v16
   532: d7895833                 vrgather.vv v19, v21, v17
   536: 5709392f                 vxor.vv v18, v19, v18
   540: 5704242f                 vxor.vv v8, v18, v8
   544: 57886627                 vand.vv v16, v22, v13
   548: d738d2a2                 vsrl.vi v17, v13, 4
   552: 57094833                 vrgather.vv v18, v20, v16
   556: d7895833                 vrgather.vv v19, v21, v17
   560: 5709392f                 vxor.vv v18, v19, v18
   564: d784242f                 vxor.vv v9, v18, v9
   568: 57086727                 vand.vv v16, v22, v14
   572: d738e2a2                 vsrl.vi v17, v14, 4
   576: 57094833                 vrgather.vv v18, v20, v16
   580: d7895833                 vrgather.vv v19, v21, v17
   584: 5709392f                 vxor.vv v18, v19, v18
   588: 5705252f                 vxor.vv v10, v18, v10
   592: 57886727                 vand.vv v16, v22, v15
   596: d738f2a2                 vsrl.vi v17, v15, 4
   600: 57094833                 vrgather.vv v18, v20, v16
   604: d7895833                 vrgather.vv v19, v21, v17
   608: 5709392f                 vxor.vv v18, v19, v18
   612: d785252f                 vxor.vv v11, v18, v11
   616: 83330501                 ld t2, 0x10(a0)
   620: b3836300                 add t2, t2, t1
   624: 07860302                 vle8.v v12, (t2)
   628: b3835300                 add t2, t2, t0
   632: 87860302                 vle8.v v13, (t2)
   636: b3835300                 add t2, t2, t0
   640: 07870302                 vle8.v v14, (t2)
   644: b3835300                 add t2, t2, t0
   648: 87870302                 vle8.v v15, (t2)
   652: 5700c02e                 vxor.vv v0, v12, v0
   656: d780d02e                 vxor.vv v1, v13, v1
   660: 5701e12e                 vxor.vv v2, v14, v2
   664: d781f12e                 vxor.vv v3, v15, v3
   668: 138e0f04                 addi t3, t6, 0x40
   672: 070a0e02                 vle8.v v20, (t3)
   676: 138e0f05                 addi t3, t6, 0x50
   680: 870a0e02                 vle8.v v21, (t3)
   684: 57086627                 vand.vv v16, v22, v12
   688: d738c2a2                 vsrl.vi v17, v12, 4
   692: 57094833                 vrgather.vv v18, v20, v16
   696: d7895833                 vrgather.vv v19, v21, v17
   700: 5709392f                 vxor.vv v18, v19, v18
   704: 5702222f                 vxor.vv v4, v18, v4
   708: 57886627                 vand.vv v16, v22, v13
   712: d738d2a2                 vsrl.vi v17, v13, 4
   716: 57094833                 vrgather.vv v18, v20, v16
   720: d7895833                 vrgather.vv v19, v21, v17
   724: 5709392f                 vxor.vv v18, v19, v18
   728: d782222f                 vxor.vv v5, v18, v5
   732: 57086727                 vand.vv v16, v22, v14
   736: d738e2a2                 vsrl.vi v17, v14, 4
   740: 57094833                 vrgather.vv v18, v20, v16
   744: d7895833                 vrgather.vv v19, v21, v17
   748: 5709392f                 vxor.vv v18, v19, v18
   752: 5703232f                 vxor.vv v6, v18, v6
   756: 57886727                 vand.vv v16, v22, v15
   760: d738f2a2                 vsrl.vi v17, v15, 4
   764: 57094833                 vrgather.vv v18, v20, v16
   768: d7895833                 vrgather.vv v19, v21, v17
   772: 5709392f                 vxor.vv v18, v19, v18
   776: d783232f                 vxor.vv v7, v18, v7
   780: 138e0f06                 addi t3, t6, 0x60
   784: 070a0e02                 vle8.v v20, (t3)
   788: 138e0f07                 addi t3, t6, 0x70
   792: 870a0e02                 vle8.v v21, (t3)
   796: 57086627                 vand.vv v16, v22, v12
   800: d738c2a2                 vsrl.vi v17, v12, 4
   804: 57094833                 vrgather.vv v18, v20, v16
   808: d7895833                 vrgather.vv v19, v21, v17
   812: 5709392f                 vxor.vv v18, v19, v18
   816: 5704242f                 vxor.vv v8, v18, v8
   820: 57886627                 vand.vv v16, v22, v13
   824: d738d2a2                 vsrl.vi v17, v13, 4
   828: 57094833                 vrgather.vv v18, v20, v16
   832: d7895833                 vrgather.vv v19, v21, v17
   836: 5709392f                 vxor.vv v18, v19, v18
   840: d784242f                 vxor.vv v9, v18, v9
   844: 57086727                 vand.vv v16, v22, v14
   848: d738e2a2                 vsrl.vi v17, v14, 4
   852: 57094833                 vrgather.vv v18, v20, v16
   856: d7895833                 vrgather.vv v19, v21, v17
   860: 5709392f                 vxor.vv v18, v19, v18
   864: 5705252f                 vxor.vv v10, v18, v10
   868: 57886727                 vand.vv v16, v22, v15
   872: d738f2a2                 vsrl.vi v17, v15, 4
   876: 57094833                 vrgather.vv v18, v20, v16
   880: d7895833                 vrgather.vv v19, v21, v17
   884: 5709392f                 vxor.vv v18, v19, v18
   888: d785252f                 vxor.vv v11, v18, v11
   892: 83338501                 ld t2, 0x18(a0)
   896: b3836300                 add t2, t2, t1
   900: 07860302                 vle8.v v12, (t2)
   904: b3835300                 add t2, t2, t0
   908: 87860302                 vle8.v v13, (t2)
   912: b3835300                 add t2, t2, t0
   916: 07870302                 vle8.v v14, (t2)
   920: b3835300                 add t2, t2, t0
   924: 87870302                 vle8.v v15, (t2)
   928: 5700c02e                 vxor.vv v0, v12, v0
   932: d780d02e                 vxor.vv v1, v13, v1
   936: 5701e12e                 vxor.vv v2, v14, v2
   940: d781f12e                 vxor.vv v3, v15, v3
   944: 138e0f08                 addi t3, t6, 0x80
   948: 070a0e02                 vle8.v v20, (t3)
   952: 138e0f09                 addi t3, t6, 0x90
   956: 870a0e02                 vle8.v v21, (t3)
   960: 57086627                 vand.vv v16, v22, v12
   964: d738c2a2                 vsrl.vi v17, v12, 4
   968: 57094833                 vrgather.vv v18, v20, v16
   972: d7895833                 vrgather.vv v19, v21, v17
   976: 5709392f                 vxor.vv v18, v19, v18
   980: 5702222f                 vxor.vv v4, v18, v4
   984: 57886627                 vand.vv v16, v22, v13
   988: d738d2a2                 vsrl.vi v17, v13, 4
   992: 57094833                 vrgather.vv v18, v20, v16
   996: d7895833                 vrgather.vv v19, v21, v17
  1000: 5709392f                 vxor.vv v18, v19, v18
  1004: d782222f                 vxor.vv v5, v18, v5
  1008: 57086727                 vand.vv v16, v22, v14
  1012: d738e2a2                 vsrl.vi v17, v14, 4
  1016: 57094833                 vrgather.vv v18, v20, v16
  1020: d7895833                 vrgather.vv v19, v21, v17
  1024: 5709392f                 vxor.vv v18, v19, v18
  1028: 5703232f                 vxor.vv v6, v18, v6
  1032: 57886727                 vand.vv v16, v22, v15
  1036: d738f2a2                 vsrl.vi v17, v15, 4
  1040: 57094833                 vrgather.vv v18, v20, v16
  1044: d7895833                 vrgather.vv v19, v21, v17
  1048: 5709392f                 vxor.vv v18, v19, v18
  1052: d783232f                 vxor.vv v7, v18, v7
  1056: 138e0f00                 mv t3, t6
  1060: 070a0e02                 vle8.v v20, (t3)
  1064: 138e0f01                 addi t3, t6, 0x10
  1068: 870a0e02                 vle8.v v21, (t3)
  1072: 57086627                 vand.vv v16, v22, v12
  1076: d738c2a2                 vsrl.vi v17, v12, 4
  1080: 57094833                 vrgather.vv v18, v20, v16
  1084: d7895833                 vrgather.vv v19, v21, v17
  1088: 5709392f                 vxor.vv v18, v19, v18
  1092: 5704242f                 vxor.vv v8, v18, v8
  1096: 57886627                 vand.vv v16, v22, v13
  1100: d738d2a2                 vsrl.vi v17, v13, 4
  1104: 57094833                 vrgather.vv v18, v20, v16
  1108: d7895833                 vrgather.vv v19, v21, v17
  1112: 5709392f                 vxor.vv v18, v19, v18
  1116: d784242f                 vxor.vv v9, v18, v9
  1120: 57086727                 vand.vv v16, v22, v14
  1124: d738e2a2                 vsrl.vi v17, v14, 4
  1128: 57094833                 vrgather.vv v18, v20, v16
  1132: d7895833                 vrgather.vv v19, v21, v17
  1136: 5709392f                 vxor.vv v18, v19, v18
  1140: 5705252f                 vxor.vv v10, v18, v10
  1144: 57886727                 vand.vv v16, v22, v15
  1148: d738f2a2                 vsrl.vi v17, v15, 4
  1152: 57094833                 vrgather.vv v18, v20, v16
  1156: d7895833                 vrgather.vv v19, v21, v17
  1160: 5709392f                 vxor.vv v18, v19, v18
  1164: d785252f                 vxor.vv v11, v18, v11
  1168: 83330502                 ld t2, 0x20(a0)
  1172: b3836300                 add t2, t2, t1
  1176: 07860302                 vle8.v v12, (t2)
  1180: b3835300                 add t2, t2, t0
  1184: 87860302                 vle8.v v13, (t2)
  1188: b3835300                 add t2, t2, t0
  1192: 07870302                 vle8.v v14, (t2)
  1196: b3835300                 add t2, t2, t0
  1200: 87870302                 vle8.v v15, (t2)
  1204: 5700c02e                 vxor.vv v0, v12, v0
  1208: d780d02e                 vxor.vv v1, v13, v1
  1212: 5701e12e                 vxor.vv v2, v14, v2
  1216: d781f12e                 vxor.vv v3, v15, v3
  1220: 138e0f06                 addi t3, t6, 0x60
  1224: 070a0e02                 vle8.v v20, (t3)
  1228: 138e0f07                 addi t3, t6, 0x70
  1232: 870a0e02                 vle8.v v21, (t3)
  1236: 57086627                 vand.vv v16, v22, v12
  1240: d738c2a2                 vsrl.vi v17, v12, 4
  1244: 57094833                 vrgather.vv v18, v20, v16
  1248: d7895833                 vrgather.vv v19, v21, v17
  1252: 5709392f                 vxor.vv v18, v19, v18
  1256: 5702222f                 vxor.vv v4, v18, v4
  1260: 57886627                 vand.vv v16, v22, v13
  1264: d738d2a2                 vsrl.vi v17, v13, 4
  1268: 57094833                 vrgather.vv v18, v20, v16
  1272: d7895833                 vrgather.vv v19, v21, v17
  1276: 5709392f                 vxor.vv v18, v19, v18
  1280: d782222f                 vxor.vv v5, v18, v5
  1284: 57086727                 vand.vv v16, v22, v14
  1288: d738e2a2                 vsrl.vi v17, v14, 4
  1292: 57094833                 vrgather.vv v18, v20, v16
  1296: d7895833                 vrgather.vv v19, v21, v17
  1300: 5709392f                 vxor.vv v18, v19, v18
  1304: 5703232f                 vxor.vv v6, v18, v6
  1308: 57886727                 vand.vv v16, v22, v15
  1312: d738f2a2                 vsrl.vi v17, v15, 4
  1316: 57094833                 vrgather.vv v18, v20, v16
  1320: d7895833                 vrgather.vv v19, v21, v17
  1324: 5709392f                 vxor.vv v18, v19, v18
  1328: d783232f                 vxor.vv v7, v18, v7
  1332: 138e0f08                 addi t3, t6, 0x80
  1336: 070a0e02                 vle8.v v20, (t3)
  1340: 138e0f09                 addi t3, t6, 0x90
  1344: 870a0e02                 vle8.v v21, (t3)
  1348: 57086627                 vand.vv v16, v22, v12
  1352: d738c2a2                 vsrl.vi v17, v12, 4
  1356: 57094833                 vrgather.vv v18, v20, v16
  1360: d7895833                 vrgather.vv v19, v21, v17
  1364: 5709392f                 vxor.vv v18, v19, v18
  1368: 5704242f                 vxor.vv v8, v18, v8
  1372: 57886627                 vand.vv v16, v22, v13
  1376: d738d2a2                 vsrl.vi v17, v13, 4
  1380: 57094833                 vrgather.vv v18, v20, v16
  1384: d7895833                 vrgather.vv v19, v21, v17
  1388: 5709392f                 vxor.vv v18, v19, v18
  1392: d784242f                 vxor.vv v9, v18, v9
  1396: 57086727                 vand.vv v16, v22, v14
  1400: d738e2a2                 vsrl.vi v17, v14, 4
  1404: 57094833                 vrgather.vv v18, v20, v16
  1408: d7895833                 vrgather.vv v19, v21, v17
  1412: 5709392f                 vxor.vv v18, v19, v18
  1416: 5705252f                 vxor.vv v10, v18, v10
  1420: 57886727                 vand.vv v16, v22, v15
  1424: d738f2a2                 vsrl.vi v17, v15, 4
  1428: 57094833                 vrgather.vv v18, v20, v16
  1432: d7895833                 vrgather.vv v19, v21, v17
  1436: 5709392f                 vxor.vv v18, v19, v18
  1440: d785252f                 vxor.vv v11, v18, v11
  1444: 83338502                 ld t2, 0x28(a0)
  1448: b3836300                 add t2, t2, t1
  1452: 07860302                 vle8.v v12, (t2)
  1456: b3835300                 add t2, t2, t0
  1460: 87860302                 vle8.v v13, (t2)
  1464: b3835300                 add t2, t2, t0
  1468: 07870302                 vle8.v v14, (t2)
  1472: b3835300                 add t2, t2, t0
  1476: 87870302                 vle8.v v15, (t2)
  1480: 5700c02e                 vxor.vv v0, v12, v0
  1484: d780d02e                 vxor.vv v1, v13, v1
  1488: 5701e12e                 vxor.vv v2, v14, v2
  1492: d781f12e                 vxor.vv v3, v15, v3
  1496: 5702c22e                 vxor.vv v4, v12, v4
  1500: d782d22e                 vxor.vv v5, v13, v5
  1504: 5703e32e                 vxor.vv v6, v14, v6
  1508: d783f32e                 vxor.vv v7, v15, v7
  1512: 5704c42e                 vxor.vv v8, v12, v8
  1516: d784d42e                 vxor.vv v9, v13, v9
  1520: 5705e52e                 vxor.vv v10, v14, v10
  1524: d785f52e                 vxor.vv v11, v15, v11
  1528: 83b30500                 ld t2, 0(a1)
  1532: b3836300                 add t2, t2, t1
  1536: 27800302                 vse8.v v0, (t2)
  1540: b3835300                 add t2, t2, t0
  1544: a7800302                 vse8.v v1, (t2)
  1548: b3835300                 add t2, t2, t0
  1552: 27810302                 vse8.v v2, (t2)
  1556: b3835300                 add t2, t2, t0
  1560: a7810302                 vse8.v v3, (t2)
  1564: 83b38500                 ld t2, 8(a1)
  1568: b3836300                 add t2, t2, t1
  1572: 27820302                 vse8.v v4, (t2)
  1576: b3835300                 add t2, t2, t0
  1580: a7820302                 vse8.v v5, (t2)
  1584: b3835300                 add t2, t2, t0
  1588: 27830302                 vse8.v v6, (t2)
  1592: b3835300                 add t2, t2, t0
  1596: a7830302                 vse8.v v7, (t2)
  1600: 83b30501                 ld t2, 0x10(a1)
  1604: b3836300                 add t2, t2, t1
  1608: 27840302                 vse8.v v8, (t2)
  1612: b3835300                 add t2, t2, t0
  1616: a7840302                 vse8.v v9, (t2)
  1620: b3835300                 add t2, t2, t0
  1624: 27850302                 vse8.v v10, (t2)
  1628: b3835300                 add t2, t2, t0
  1632: a7850302                 vse8.v v11, (t2)
  1636: 3303d300                 add t1, t1, a3
  1640: 1307f7ff                 addi a4, a4, -1
  1644: e31a079c                 bnez a4, 0x40
  1648: 33066640                 sub a2, a2, t1
  1652: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
  1656: 63940200                 bnez t0, 0x680
  1660: 6f00402a                 j 0x920
  1664: 83330500                 ld t2, 0(a0)
  1668: b3836300                 add t2, t2, t1
  1672: 07860302                 vle8.v v12, (t2)
  1676: 5700065e                 vmv.v.v v0, v12
  1680: 138f0200                 mv t5, t0
  1684: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1688: 138e0f00                 mv t3, t6
  1692: 070a0e02                 vle8.v v20, (t3)
  1696: 138e0f01                 addi t3, t6, 0x10
  1700: 870a0e02                 vle8.v v21, (t3)
  1704: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1708: 57086627                 vand.vv v16, v22, v12
  1712: d738c2a2                 vsrl.vi v17, v12, 4
  1716: 57094833                 vrgather.vv v18, v20, v16
  1720: d7895833                 vrgather.vv v19, v21, v17
  1724: 5709392f                 vxor.vv v18, v19, v18
  1728: 5702095e                 vmv.v.v v4, v18
  1732: 138f0200                 mv t5, t0
  1736: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1740: 138e0f02                 addi t3, t6, 0x20
  1744: 070a0e02                 vle8.v v20, (t3)
  1748: 138e0f03                 addi t3, t6, 0x30
  1752: 870a0e02                 vle8.v v21, (t3)
  1756: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1760: 57086627                 vand.vv v16, v22, v12
  1764: d738c2a2                 vsrl.vi v17, v12, 4
  1768: 57094833                 vrgather.vv v18, v20, v16
  1772: d7895833                 vrgather.vv v19, v21, v17
  1776: 5709392f                 vxor.vv v18, v19, v18
  1780: 5704095e                 vmv.v.v v8, v18
  1784: 83338500                 ld t2, 8(a0)
  1788: b3836300                 add t2, t2, t1
  1792: 07860302                 vle8.v v12, (t2)
  1796: 5700c02e                 vxor.vv v0, v12, v0
  1800: 138f0200                 mv t5, t0
  1804: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1808: 138e0f02                 addi t3, t6, 0x20
  1812: 070a0e02                 vle8.v v20, (t3)
  1816: 138e0f03                 addi t3, t6, 0x30
  1820: 870a0e02                 vle8.v v21, (t3)
  1824: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1828: 57086627                 vand.vv v16, v22, v12
  1832: d738c2a2                 vsrl.vi v17, v12, 4
  1836: 57094833                 vrgather.vv v18, v20, v16
  1840: d7895833                 vrgather.vv v19, v21, v17
  1844: 5709392f                 vxor.vv v18, v19, v18
  1848: 5702222f                 vxor.vv v4, v18, v4
  1852: 138f0200                 mv t5, t0
  1856: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1860: 138e0f04                 addi t3, t6, 0x40
  1864: 070a0e02                 vle8.v v20, (t3)
  1868: 138e0f05                 addi t3, t6, 0x50
  1872: 870a0e02                 vle8.v v21, (t3)
  1876: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1880: 57086627                 vand.vv v16, v22, v12
  1884: d738c2a2                 vsrl.vi v17, v12, 4
  1888: 57094833                 vrgather.vv v18, v20, v16
  1892: d7895833                 vrgather.vv v19, v21, v17
  1896: 5709392f                 vxor.vv v18, v19, v18
  1900: 5704242f                 vxor.vv v8, v18, v8
  1904: 83330501                 ld t2, 0x10(a0)
  1908: b3836300                 add t2, t2, t1
  1912: 07860302                 vle8.v v12, (t2)
  1916: 5700c02e                 vxor.vv v0, v12, v0
  1920: 138f0200                 mv t5, t0
  1924: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1928: 138e0f04                 addi t3, t6, 0x40
  1932: 070a0e02                 vle8.v v20, (t3)
  1936: 138e0f05                 addi t3, t6, 0x50
  1940: 870a0e02                 vle8.v v21, (t3)
  1944: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1948: 57086627                 vand.vv v16, v22, v12
  1952: d738c2a2                 vsrl.vi v17, v12, 4
  1956: 57094833                 vrgather.vv v18, v20, v16
  1960: d7895833                 vrgather.vv v19, v21, v17
  1964: 5709392f                 vxor.vv v18, v19, v18
  1968: 5702222f                 vxor.vv v4, v18, v4
  1972: 138f0200                 mv t5, t0
  1976: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1980: 138e0f06                 addi t3, t6, 0x60
  1984: 070a0e02                 vle8.v v20, (t3)
  1988: 138e0f07                 addi t3, t6, 0x70
  1992: 870a0e02                 vle8.v v21, (t3)
  1996: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2000: 57086627                 vand.vv v16, v22, v12
  2004: d738c2a2                 vsrl.vi v17, v12, 4
  2008: 57094833                 vrgather.vv v18, v20, v16
  2012: d7895833                 vrgather.vv v19, v21, v17
  2016: 5709392f                 vxor.vv v18, v19, v18
  2020: 5704242f                 vxor.vv v8, v18, v8
  2024: 83338501                 ld t2, 0x18(a0)
  2028: b3836300                 add t2, t2, t1
  2032: 07860302                 vle8.v v12, (t2)
  2036: 5700c02e                 vxor.vv v0, v12, v0
  2040: 138f0200                 mv t5, t0
  2044: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2048: 138e0f08                 addi t3, t6, 0x80
  2052: 070a0e02                 vle8.v v20, (t3)
  2056: 138e0f09                 addi t3, t6, 0x90
  2060: 870a0e02                 vle8.v v21, (t3)
  2064: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2068: 57086627                 vand.vv v16, v22, v12
  2072: d738c2a2                 vsrl.vi v17, v12, 4
  2076: 57094833                 vrgather.vv v18, v20, v16
  2080: d7895833                 vrgather.vv v19, v21, v17
  2084: 5709392f                 vxor.vv v18, v19, v18
  2088: 5702222f                 vxor.vv v4, v18, v4
  2092: 138f0200                 mv t5, t0
  2096: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2100: 138e0f00                 mv t3, t6
  2104: 070a0e02                 vle8.v v20, (t3)
  2108: 138e0f01                 addi t3, t6, 0x10
  2112: 870a0e02                 vle8.v v21, (t3)
  2116: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2120: 57086627                 vand.vv v16, v22, v12
  2124: d738c2a2                 vsrl.vi v17, v12, 4
  2128: 57094833                 vrgather.vv v18, v20, v16
  2132: d7895833                 vrgather.vv v19, v21, v17
  2136: 5709392f                 vxor.vv v18, v19, v18
  2140: 5704242f                 vxor.vv v8, v18, v8
  2144: 83330502                 ld t2, 0x20(a0)
  2148: b3836300                 add t2, t2, t1
  2152: 07860302                 vle8.v v12, (t2)
  2156: 5700c02e                 vxor.vv v0, v12, v0
  2160: 138f0200                 mv t5, t0
  2164: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2168: 138e0f06                 addi t3, t6, 0x60
  2172: 070a0e02                 vle8.v v20, (t3)
  2176: 138e0f07                 addi t3, t6, 0x70
  2180: 870a0e02                 vle8.v v21, (t3)
  2184: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2188: 57086627                 vand.vv v16, v22, v12
  2192: d738c2a2                 vsrl.vi v17, v12, 4
  2196: 57094833                 vrgather.vv v18, v20, v16
  2200: d7895833                 vrgather.vv v19, v21, v17
  2204: 5709392f                 vxor.vv v18, v19, v18
  2208: 5702222f                 vxor.vv v4, v18, v4
  2212: 138f0200                 mv t5, t0
  2216: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2220: 138e0f08                 addi t3, t6, 0x80
  2224: 070a0e02                 vle8.v v20, (t3)
  2228: 138e0f09                 addi t3, t6, 0x90
  2232: 870a0e02                 vle8.v v21, (t3)
  2236: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2240: 57086627                 vand.vv v16, v22, v12
  2244: d738c2a2                 vsrl.vi v17, v12, 4
  2248: 57094833                 vrgather.vv v18, v20, v16
  2252: d7895833                 vrgather.vv v19, v21, v17
  2256: 5709392f                 vxor.vv v18, v19, v18
  2260: 5704242f                 vxor.vv v8, v18, v8
  2264: 83338502                 ld t2, 0x28(a0)
  2268: b3836300                 add t2, t2, t1
  2272: 07860302                 vle8.v v12, (t2)
  2276: 5700c02e                 vxor.vv v0, v12, v0
  2280: 5702c22e                 vxor.vv v4, v12, v4
  2284: 5704c42e                 vxor.vv v8, v12, v8
  2288: 83b30500                 ld t2, 0(a1)
  2292: b3836300                 add t2, t2, t1
  2296: 27800302                 vse8.v v0, (t2)
  2300: 83b38500                 ld t2, 8(a1)
  2304: b3836300                 add t2, t2, t1
  2308: 27820302                 vse8.v v4, (t2)
  2312: 83b30501                 ld t2, 0x10(a1)
  2316: b3836300                 add t2, t2, t1
  2320: 27840302                 vse8.v v8, (t2)
  2324: 33035300                 add t1, t1, t0
  2328: 33065640                 sub a2, a2, t0
  2332: 6ff09fd5                 j 0x674
  2336: 67800000                 ret 
  2340: 0000                     unimp 
  2342: 0000                     unimp 
  2344: 0000                     unimp 
  2346: 0000                     unimp 
  2348: 0000                     unimp 
  2350: 0000                     unimp 
