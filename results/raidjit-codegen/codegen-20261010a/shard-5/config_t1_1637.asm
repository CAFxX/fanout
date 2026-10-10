; config_t1_1637 target=rvv twin=rec_cf topo=rs(6,3)/rec{0,1} op=recover U=8
; code_size=3184 twins=[] uops_per_byte=4.9765625
     0: 63140600                 bnez a2, 8
     4: 6f001046                 j 0xc64
     8: 971f0000                 auipc t6, 1
    12: 938f8fc6                 addi t6, t6, -0x398
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574f0e5e                 vmv.v.x v30, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93068000                 li a3, 8
    40: b386d202                 mul a3, t0, a3
    44: 3357d602                 divu a4, a2, a3
    48: 63140700                 bnez a4, 0x38
    52: 6f00d01f                 j 0xa30
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07880302                 vle8.v v16, (t2)
    76: b3835300                 add t2, t2, t0
    80: 5700085e                 vmv.v.v v0, v16
    84: 87880302                 vle8.v v17, (t2)
    88: b3835300                 add t2, t2, t0
    92: d780085e                 vmv.v.v v1, v17
    96: 07890302                 vle8.v v18, (t2)
   100: b3835300                 add t2, t2, t0
   104: 5701095e                 vmv.v.v v2, v18
   108: 87890302                 vle8.v v19, (t2)
   112: b3835300                 add t2, t2, t0
   116: d781095e                 vmv.v.v v3, v19
   120: 078a0302                 vle8.v v20, (t2)
   124: b3835300                 add t2, t2, t0
   128: 57020a5e                 vmv.v.v v4, v20
   132: 878a0302                 vle8.v v21, (t2)
   136: b3835300                 add t2, t2, t0
   140: d7820a5e                 vmv.v.v v5, v21
   144: 078b0302                 vle8.v v22, (t2)
   148: b3835300                 add t2, t2, t0
   152: 57030b5e                 vmv.v.v v6, v22
   156: 878b0302                 vle8.v v23, (t2)
   160: d7830b5e                 vmv.v.v v7, v23
   164: 83338500                 ld t2, 8(a0)
   168: b3836300                 add t2, t2, t1
   172: 07880302                 vle8.v v16, (t2)
   176: b3835300                 add t2, t2, t0
   180: 5700002f                 vxor.vv v0, v16, v0
   184: 87880302                 vle8.v v17, (t2)
   188: b3835300                 add t2, t2, t0
   192: d780102f                 vxor.vv v1, v17, v1
   196: 07890302                 vle8.v v18, (t2)
   200: b3835300                 add t2, t2, t0
   204: 5701212f                 vxor.vv v2, v18, v2
   208: 87890302                 vle8.v v19, (t2)
   212: b3835300                 add t2, t2, t0
   216: d781312f                 vxor.vv v3, v19, v3
   220: 078a0302                 vle8.v v20, (t2)
   224: b3835300                 add t2, t2, t0
   228: 5702422f                 vxor.vv v4, v20, v4
   232: 878a0302                 vle8.v v21, (t2)
   236: b3835300                 add t2, t2, t0
   240: d782522f                 vxor.vv v5, v21, v5
   244: 078b0302                 vle8.v v22, (t2)
   248: b3835300                 add t2, t2, t0
   252: 5703632f                 vxor.vv v6, v22, v6
   256: 878b0302                 vle8.v v23, (t2)
   260: d783732f                 vxor.vv v7, v23, v7
   264: 83330501                 ld t2, 0x10(a0)
   268: b3836300                 add t2, t2, t1
   272: 07880302                 vle8.v v16, (t2)
   276: b3835300                 add t2, t2, t0
   280: 5700002f                 vxor.vv v0, v16, v0
   284: 87880302                 vle8.v v17, (t2)
   288: b3835300                 add t2, t2, t0
   292: d780102f                 vxor.vv v1, v17, v1
   296: 07890302                 vle8.v v18, (t2)
   300: b3835300                 add t2, t2, t0
   304: 5701212f                 vxor.vv v2, v18, v2
   308: 87890302                 vle8.v v19, (t2)
   312: b3835300                 add t2, t2, t0
   316: d781312f                 vxor.vv v3, v19, v3
   320: 078a0302                 vle8.v v20, (t2)
   324: b3835300                 add t2, t2, t0
   328: 5702422f                 vxor.vv v4, v20, v4
   332: 878a0302                 vle8.v v21, (t2)
   336: b3835300                 add t2, t2, t0
   340: d782522f                 vxor.vv v5, v21, v5
   344: 078b0302                 vle8.v v22, (t2)
   348: b3835300                 add t2, t2, t0
   352: 5703632f                 vxor.vv v6, v22, v6
   356: 878b0302                 vle8.v v23, (t2)
   360: d783732f                 vxor.vv v7, v23, v7
   364: 83338501                 ld t2, 0x18(a0)
   368: b3836300                 add t2, t2, t1
   372: 07880302                 vle8.v v16, (t2)
   376: b3835300                 add t2, t2, t0
   380: 5700002f                 vxor.vv v0, v16, v0
   384: 87880302                 vle8.v v17, (t2)
   388: b3835300                 add t2, t2, t0
   392: d780102f                 vxor.vv v1, v17, v1
   396: 07890302                 vle8.v v18, (t2)
   400: b3835300                 add t2, t2, t0
   404: 5701212f                 vxor.vv v2, v18, v2
   408: 87890302                 vle8.v v19, (t2)
   412: b3835300                 add t2, t2, t0
   416: d781312f                 vxor.vv v3, v19, v3
   420: 078a0302                 vle8.v v20, (t2)
   424: b3835300                 add t2, t2, t0
   428: 5702422f                 vxor.vv v4, v20, v4
   432: 878a0302                 vle8.v v21, (t2)
   436: b3835300                 add t2, t2, t0
   440: d782522f                 vxor.vv v5, v21, v5
   444: 078b0302                 vle8.v v22, (t2)
   448: b3835300                 add t2, t2, t0
   452: 5703632f                 vxor.vv v6, v22, v6
   456: 878b0302                 vle8.v v23, (t2)
   460: d783732f                 vxor.vv v7, v23, v7
   464: 83330502                 ld t2, 0x20(a0)
   468: b3836300                 add t2, t2, t1
   472: 07880302                 vle8.v v16, (t2)
   476: b3835300                 add t2, t2, t0
   480: 5700002f                 vxor.vv v0, v16, v0
   484: 87880302                 vle8.v v17, (t2)
   488: b3835300                 add t2, t2, t0
   492: d780102f                 vxor.vv v1, v17, v1
   496: 07890302                 vle8.v v18, (t2)
   500: b3835300                 add t2, t2, t0
   504: 5701212f                 vxor.vv v2, v18, v2
   508: 87890302                 vle8.v v19, (t2)
   512: b3835300                 add t2, t2, t0
   516: d781312f                 vxor.vv v3, v19, v3
   520: 078a0302                 vle8.v v20, (t2)
   524: b3835300                 add t2, t2, t0
   528: 5702422f                 vxor.vv v4, v20, v4
   532: 878a0302                 vle8.v v21, (t2)
   536: b3835300                 add t2, t2, t0
   540: d782522f                 vxor.vv v5, v21, v5
   544: 078b0302                 vle8.v v22, (t2)
   548: b3835300                 add t2, t2, t0
   552: 5703632f                 vxor.vv v6, v22, v6
   556: 878b0302                 vle8.v v23, (t2)
   560: d783732f                 vxor.vv v7, v23, v7
   564: 83330500                 ld t2, 0(a0)
   568: b3836300                 add t2, t2, t1
   572: 138e0f00                 mv t3, t6
   576: 070e0e02                 vle8.v v28, (t3)
   580: 138e0f01                 addi t3, t6, 0x10
   584: 870e0e02                 vle8.v v29, (t3)
   588: 07880302                 vle8.v v16, (t2)
   592: b3835300                 add t2, t2, t0
   596: 570ce827                 vand.vv v24, v30, v16
   600: d73c02a3                 vsrl.vi v25, v16, 4
   604: 570dcc33                 vrgather.vv v26, v28, v24
   608: d78ddc33                 vrgather.vv v27, v29, v25
   612: 570dbd2f                 vxor.vv v26, v27, v26
   616: 57040d5e                 vmv.v.v v8, v26
   620: 87880302                 vle8.v v17, (t2)
   624: b3835300                 add t2, t2, t0
   628: 578ce827                 vand.vv v24, v30, v17
   632: d73c12a3                 vsrl.vi v25, v17, 4
   636: 570dcc33                 vrgather.vv v26, v28, v24
   640: d78ddc33                 vrgather.vv v27, v29, v25
   644: 570dbd2f                 vxor.vv v26, v27, v26
   648: d7040d5e                 vmv.v.v v9, v26
   652: 07890302                 vle8.v v18, (t2)
   656: b3835300                 add t2, t2, t0
   660: 570ce927                 vand.vv v24, v30, v18
   664: d73c22a3                 vsrl.vi v25, v18, 4
   668: 570dcc33                 vrgather.vv v26, v28, v24
   672: d78ddc33                 vrgather.vv v27, v29, v25
   676: 570dbd2f                 vxor.vv v26, v27, v26
   680: 57050d5e                 vmv.v.v v10, v26
   684: 87890302                 vle8.v v19, (t2)
   688: b3835300                 add t2, t2, t0
   692: 578ce927                 vand.vv v24, v30, v19
   696: d73c32a3                 vsrl.vi v25, v19, 4
   700: 570dcc33                 vrgather.vv v26, v28, v24
   704: d78ddc33                 vrgather.vv v27, v29, v25
   708: 570dbd2f                 vxor.vv v26, v27, v26
   712: d7050d5e                 vmv.v.v v11, v26
   716: 078a0302                 vle8.v v20, (t2)
   720: b3835300                 add t2, t2, t0
   724: 570cea27                 vand.vv v24, v30, v20
   728: d73c42a3                 vsrl.vi v25, v20, 4
   732: 570dcc33                 vrgather.vv v26, v28, v24
   736: d78ddc33                 vrgather.vv v27, v29, v25
   740: 570dbd2f                 vxor.vv v26, v27, v26
   744: 57060d5e                 vmv.v.v v12, v26
   748: 878a0302                 vle8.v v21, (t2)
   752: b3835300                 add t2, t2, t0
   756: 578cea27                 vand.vv v24, v30, v21
   760: d73c52a3                 vsrl.vi v25, v21, 4
   764: 570dcc33                 vrgather.vv v26, v28, v24
   768: d78ddc33                 vrgather.vv v27, v29, v25
   772: 570dbd2f                 vxor.vv v26, v27, v26
   776: d7060d5e                 vmv.v.v v13, v26
   780: 078b0302                 vle8.v v22, (t2)
   784: b3835300                 add t2, t2, t0
   788: 570ceb27                 vand.vv v24, v30, v22
   792: d73c62a3                 vsrl.vi v25, v22, 4
   796: 570dcc33                 vrgather.vv v26, v28, v24
   800: d78ddc33                 vrgather.vv v27, v29, v25
   804: 570dbd2f                 vxor.vv v26, v27, v26
   808: 57070d5e                 vmv.v.v v14, v26
   812: 878b0302                 vle8.v v23, (t2)
   816: 578ceb27                 vand.vv v24, v30, v23
   820: d73c72a3                 vsrl.vi v25, v23, 4
   824: 570dcc33                 vrgather.vv v26, v28, v24
   828: d78ddc33                 vrgather.vv v27, v29, v25
   832: 570dbd2f                 vxor.vv v26, v27, v26
   836: d7070d5e                 vmv.v.v v15, v26
   840: 83338500                 ld t2, 8(a0)
   844: b3836300                 add t2, t2, t1
   848: 138e0f02                 addi t3, t6, 0x20
   852: 070e0e02                 vle8.v v28, (t3)
   856: 138e0f03                 addi t3, t6, 0x30
   860: 870e0e02                 vle8.v v29, (t3)
   864: 07880302                 vle8.v v16, (t2)
   868: b3835300                 add t2, t2, t0
   872: 570ce827                 vand.vv v24, v30, v16
   876: d73c02a3                 vsrl.vi v25, v16, 4
   880: 570dcc33                 vrgather.vv v26, v28, v24
   884: d78ddc33                 vrgather.vv v27, v29, v25
   888: 570dbd2f                 vxor.vv v26, v27, v26
   892: 5704a42f                 vxor.vv v8, v26, v8
   896: 87880302                 vle8.v v17, (t2)
   900: b3835300                 add t2, t2, t0
   904: 578ce827                 vand.vv v24, v30, v17
   908: d73c12a3                 vsrl.vi v25, v17, 4
   912: 570dcc33                 vrgather.vv v26, v28, v24
   916: d78ddc33                 vrgather.vv v27, v29, v25
   920: 570dbd2f                 vxor.vv v26, v27, v26
   924: d784a42f                 vxor.vv v9, v26, v9
   928: 07890302                 vle8.v v18, (t2)
   932: b3835300                 add t2, t2, t0
   936: 570ce927                 vand.vv v24, v30, v18
   940: d73c22a3                 vsrl.vi v25, v18, 4
   944: 570dcc33                 vrgather.vv v26, v28, v24
   948: d78ddc33                 vrgather.vv v27, v29, v25
   952: 570dbd2f                 vxor.vv v26, v27, v26
   956: 5705a52f                 vxor.vv v10, v26, v10
   960: 87890302                 vle8.v v19, (t2)
   964: b3835300                 add t2, t2, t0
   968: 578ce927                 vand.vv v24, v30, v19
   972: d73c32a3                 vsrl.vi v25, v19, 4
   976: 570dcc33                 vrgather.vv v26, v28, v24
   980: d78ddc33                 vrgather.vv v27, v29, v25
   984: 570dbd2f                 vxor.vv v26, v27, v26
   988: d785a52f                 vxor.vv v11, v26, v11
   992: 078a0302                 vle8.v v20, (t2)
   996: b3835300                 add t2, t2, t0
  1000: 570cea27                 vand.vv v24, v30, v20
  1004: d73c42a3                 vsrl.vi v25, v20, 4
  1008: 570dcc33                 vrgather.vv v26, v28, v24
  1012: d78ddc33                 vrgather.vv v27, v29, v25
  1016: 570dbd2f                 vxor.vv v26, v27, v26
  1020: 5706a62f                 vxor.vv v12, v26, v12
  1024: 878a0302                 vle8.v v21, (t2)
  1028: b3835300                 add t2, t2, t0
  1032: 578cea27                 vand.vv v24, v30, v21
  1036: d73c52a3                 vsrl.vi v25, v21, 4
  1040: 570dcc33                 vrgather.vv v26, v28, v24
  1044: d78ddc33                 vrgather.vv v27, v29, v25
  1048: 570dbd2f                 vxor.vv v26, v27, v26
  1052: d786a62f                 vxor.vv v13, v26, v13
  1056: 078b0302                 vle8.v v22, (t2)
  1060: b3835300                 add t2, t2, t0
  1064: 570ceb27                 vand.vv v24, v30, v22
  1068: d73c62a3                 vsrl.vi v25, v22, 4
  1072: 570dcc33                 vrgather.vv v26, v28, v24
  1076: d78ddc33                 vrgather.vv v27, v29, v25
  1080: 570dbd2f                 vxor.vv v26, v27, v26
  1084: 5707a72f                 vxor.vv v14, v26, v14
  1088: 878b0302                 vle8.v v23, (t2)
  1092: 578ceb27                 vand.vv v24, v30, v23
  1096: d73c72a3                 vsrl.vi v25, v23, 4
  1100: 570dcc33                 vrgather.vv v26, v28, v24
  1104: d78ddc33                 vrgather.vv v27, v29, v25
  1108: 570dbd2f                 vxor.vv v26, v27, v26
  1112: d787a72f                 vxor.vv v15, v26, v15
  1116: 83330501                 ld t2, 0x10(a0)
  1120: b3836300                 add t2, t2, t1
  1124: 138e0f04                 addi t3, t6, 0x40
  1128: 070e0e02                 vle8.v v28, (t3)
  1132: 138e0f05                 addi t3, t6, 0x50
  1136: 870e0e02                 vle8.v v29, (t3)
  1140: 07880302                 vle8.v v16, (t2)
  1144: b3835300                 add t2, t2, t0
  1148: 570ce827                 vand.vv v24, v30, v16
  1152: d73c02a3                 vsrl.vi v25, v16, 4
  1156: 570dcc33                 vrgather.vv v26, v28, v24
  1160: d78ddc33                 vrgather.vv v27, v29, v25
  1164: 570dbd2f                 vxor.vv v26, v27, v26
  1168: 5704a42f                 vxor.vv v8, v26, v8
  1172: 87880302                 vle8.v v17, (t2)
  1176: b3835300                 add t2, t2, t0
  1180: 578ce827                 vand.vv v24, v30, v17
  1184: d73c12a3                 vsrl.vi v25, v17, 4
  1188: 570dcc33                 vrgather.vv v26, v28, v24
  1192: d78ddc33                 vrgather.vv v27, v29, v25
  1196: 570dbd2f                 vxor.vv v26, v27, v26
  1200: d784a42f                 vxor.vv v9, v26, v9
  1204: 07890302                 vle8.v v18, (t2)
  1208: b3835300                 add t2, t2, t0
  1212: 570ce927                 vand.vv v24, v30, v18
  1216: d73c22a3                 vsrl.vi v25, v18, 4
  1220: 570dcc33                 vrgather.vv v26, v28, v24
  1224: d78ddc33                 vrgather.vv v27, v29, v25
  1228: 570dbd2f                 vxor.vv v26, v27, v26
  1232: 5705a52f                 vxor.vv v10, v26, v10
  1236: 87890302                 vle8.v v19, (t2)
  1240: b3835300                 add t2, t2, t0
  1244: 578ce927                 vand.vv v24, v30, v19
  1248: d73c32a3                 vsrl.vi v25, v19, 4
  1252: 570dcc33                 vrgather.vv v26, v28, v24
  1256: d78ddc33                 vrgather.vv v27, v29, v25
  1260: 570dbd2f                 vxor.vv v26, v27, v26
  1264: d785a52f                 vxor.vv v11, v26, v11
  1268: 078a0302                 vle8.v v20, (t2)
  1272: b3835300                 add t2, t2, t0
  1276: 570cea27                 vand.vv v24, v30, v20
  1280: d73c42a3                 vsrl.vi v25, v20, 4
  1284: 570dcc33                 vrgather.vv v26, v28, v24
  1288: d78ddc33                 vrgather.vv v27, v29, v25
  1292: 570dbd2f                 vxor.vv v26, v27, v26
  1296: 5706a62f                 vxor.vv v12, v26, v12
  1300: 878a0302                 vle8.v v21, (t2)
  1304: b3835300                 add t2, t2, t0
  1308: 578cea27                 vand.vv v24, v30, v21
  1312: d73c52a3                 vsrl.vi v25, v21, 4
  1316: 570dcc33                 vrgather.vv v26, v28, v24
  1320: d78ddc33                 vrgather.vv v27, v29, v25
  1324: 570dbd2f                 vxor.vv v26, v27, v26
  1328: d786a62f                 vxor.vv v13, v26, v13
  1332: 078b0302                 vle8.v v22, (t2)
  1336: b3835300                 add t2, t2, t0
  1340: 570ceb27                 vand.vv v24, v30, v22
  1344: d73c62a3                 vsrl.vi v25, v22, 4
  1348: 570dcc33                 vrgather.vv v26, v28, v24
  1352: d78ddc33                 vrgather.vv v27, v29, v25
  1356: 570dbd2f                 vxor.vv v26, v27, v26
  1360: 5707a72f                 vxor.vv v14, v26, v14
  1364: 878b0302                 vle8.v v23, (t2)
  1368: 578ceb27                 vand.vv v24, v30, v23
  1372: d73c72a3                 vsrl.vi v25, v23, 4
  1376: 570dcc33                 vrgather.vv v26, v28, v24
  1380: d78ddc33                 vrgather.vv v27, v29, v25
  1384: 570dbd2f                 vxor.vv v26, v27, v26
  1388: d787a72f                 vxor.vv v15, v26, v15
  1392: 83338501                 ld t2, 0x18(a0)
  1396: b3836300                 add t2, t2, t1
  1400: 07880302                 vle8.v v16, (t2)
  1404: b3835300                 add t2, t2, t0
  1408: 5704042f                 vxor.vv v8, v16, v8
  1412: 87880302                 vle8.v v17, (t2)
  1416: b3835300                 add t2, t2, t0
  1420: d784142f                 vxor.vv v9, v17, v9
  1424: 07890302                 vle8.v v18, (t2)
  1428: b3835300                 add t2, t2, t0
  1432: 5705252f                 vxor.vv v10, v18, v10
  1436: 87890302                 vle8.v v19, (t2)
  1440: b3835300                 add t2, t2, t0
  1444: d785352f                 vxor.vv v11, v19, v11
  1448: 078a0302                 vle8.v v20, (t2)
  1452: b3835300                 add t2, t2, t0
  1456: 5706462f                 vxor.vv v12, v20, v12
  1460: 878a0302                 vle8.v v21, (t2)
  1464: b3835300                 add t2, t2, t0
  1468: d786562f                 vxor.vv v13, v21, v13
  1472: 078b0302                 vle8.v v22, (t2)
  1476: b3835300                 add t2, t2, t0
  1480: 5707672f                 vxor.vv v14, v22, v14
  1484: 878b0302                 vle8.v v23, (t2)
  1488: d787772f                 vxor.vv v15, v23, v15
  1492: 83338502                 ld t2, 0x28(a0)
  1496: b3836300                 add t2, t2, t1
  1500: 07880302                 vle8.v v16, (t2)
  1504: b3835300                 add t2, t2, t0
  1508: 5704042f                 vxor.vv v8, v16, v8
  1512: 87880302                 vle8.v v17, (t2)
  1516: b3835300                 add t2, t2, t0
  1520: d784142f                 vxor.vv v9, v17, v9
  1524: 07890302                 vle8.v v18, (t2)
  1528: b3835300                 add t2, t2, t0
  1532: 5705252f                 vxor.vv v10, v18, v10
  1536: 87890302                 vle8.v v19, (t2)
  1540: b3835300                 add t2, t2, t0
  1544: d785352f                 vxor.vv v11, v19, v11
  1548: 078a0302                 vle8.v v20, (t2)
  1552: b3835300                 add t2, t2, t0
  1556: 5706462f                 vxor.vv v12, v20, v12
  1560: 878a0302                 vle8.v v21, (t2)
  1564: b3835300                 add t2, t2, t0
  1568: d786562f                 vxor.vv v13, v21, v13
  1572: 078b0302                 vle8.v v22, (t2)
  1576: b3835300                 add t2, t2, t0
  1580: 5707672f                 vxor.vv v14, v22, v14
  1584: 878b0302                 vle8.v v23, (t2)
  1588: d787772f                 vxor.vv v15, v23, v15
  1592: 138e0f06                 addi t3, t6, 0x60
  1596: 070e0e02                 vle8.v v28, (t3)
  1600: 138e0f07                 addi t3, t6, 0x70
  1604: 870e0e02                 vle8.v v29, (t3)
  1608: 570ce027                 vand.vv v24, v30, v0
  1612: d73c02a2                 vsrl.vi v25, v0, 4
  1616: 570dcc33                 vrgather.vv v26, v28, v24
  1620: d78ddc33                 vrgather.vv v27, v29, v25
  1624: 570dbd2f                 vxor.vv v26, v27, v26
  1628: 57080d5e                 vmv.v.v v16, v26
  1632: 578ce027                 vand.vv v24, v30, v1
  1636: d73c12a2                 vsrl.vi v25, v1, 4
  1640: 570dcc33                 vrgather.vv v26, v28, v24
  1644: d78ddc33                 vrgather.vv v27, v29, v25
  1648: 570dbd2f                 vxor.vv v26, v27, v26
  1652: d7080d5e                 vmv.v.v v17, v26
  1656: 570ce127                 vand.vv v24, v30, v2
  1660: d73c22a2                 vsrl.vi v25, v2, 4
  1664: 570dcc33                 vrgather.vv v26, v28, v24
  1668: d78ddc33                 vrgather.vv v27, v29, v25
  1672: 570dbd2f                 vxor.vv v26, v27, v26
  1676: 57090d5e                 vmv.v.v v18, v26
  1680: 578ce127                 vand.vv v24, v30, v3
  1684: d73c32a2                 vsrl.vi v25, v3, 4
  1688: 570dcc33                 vrgather.vv v26, v28, v24
  1692: d78ddc33                 vrgather.vv v27, v29, v25
  1696: 570dbd2f                 vxor.vv v26, v27, v26
  1700: d7090d5e                 vmv.v.v v19, v26
  1704: 570ce227                 vand.vv v24, v30, v4
  1708: d73c42a2                 vsrl.vi v25, v4, 4
  1712: 570dcc33                 vrgather.vv v26, v28, v24
  1716: d78ddc33                 vrgather.vv v27, v29, v25
  1720: 570dbd2f                 vxor.vv v26, v27, v26
  1724: 570a0d5e                 vmv.v.v v20, v26
  1728: 578ce227                 vand.vv v24, v30, v5
  1732: d73c52a2                 vsrl.vi v25, v5, 4
  1736: 570dcc33                 vrgather.vv v26, v28, v24
  1740: d78ddc33                 vrgather.vv v27, v29, v25
  1744: 570dbd2f                 vxor.vv v26, v27, v26
  1748: d70a0d5e                 vmv.v.v v21, v26
  1752: 570ce327                 vand.vv v24, v30, v6
  1756: d73c62a2                 vsrl.vi v25, v6, 4
  1760: 570dcc33                 vrgather.vv v26, v28, v24
  1764: d78ddc33                 vrgather.vv v27, v29, v25
  1768: 570dbd2f                 vxor.vv v26, v27, v26
  1772: 570b0d5e                 vmv.v.v v22, v26
  1776: 578ce327                 vand.vv v24, v30, v7
  1780: d73c72a2                 vsrl.vi v25, v7, 4
  1784: 570dcc33                 vrgather.vv v26, v28, v24
  1788: d78ddc33                 vrgather.vv v27, v29, v25
  1792: 570dbd2f                 vxor.vv v26, v27, v26
  1796: d70b0d5e                 vmv.v.v v23, v26
  1800: 5708882e                 vxor.vv v16, v8, v16
  1804: d788982e                 vxor.vv v17, v9, v17
  1808: 5709a92e                 vxor.vv v18, v10, v18
  1812: d789b92e                 vxor.vv v19, v11, v19
  1816: 570aca2e                 vxor.vv v20, v12, v20
  1820: d78ada2e                 vxor.vv v21, v13, v21
  1824: 570beb2e                 vxor.vv v22, v14, v22
  1828: d78bfb2e                 vxor.vv v23, v15, v23
  1832: 138e0f08                 addi t3, t6, 0x80
  1836: 070e0e02                 vle8.v v28, (t3)
  1840: 138e0f09                 addi t3, t6, 0x90
  1844: 870e0e02                 vle8.v v29, (t3)
  1848: 570ce827                 vand.vv v24, v30, v16
  1852: d73c02a3                 vsrl.vi v25, v16, 4
  1856: 570dcc33                 vrgather.vv v26, v28, v24
  1860: d78ddc33                 vrgather.vv v27, v29, v25
  1864: 570dbd2f                 vxor.vv v26, v27, v26
  1868: 57080d5e                 vmv.v.v v16, v26
  1872: 578ce827                 vand.vv v24, v30, v17
  1876: d73c12a3                 vsrl.vi v25, v17, 4
  1880: 570dcc33                 vrgather.vv v26, v28, v24
  1884: d78ddc33                 vrgather.vv v27, v29, v25
  1888: 570dbd2f                 vxor.vv v26, v27, v26
  1892: d7080d5e                 vmv.v.v v17, v26
  1896: 570ce927                 vand.vv v24, v30, v18
  1900: d73c22a3                 vsrl.vi v25, v18, 4
  1904: 570dcc33                 vrgather.vv v26, v28, v24
  1908: d78ddc33                 vrgather.vv v27, v29, v25
  1912: 570dbd2f                 vxor.vv v26, v27, v26
  1916: 57090d5e                 vmv.v.v v18, v26
  1920: 578ce927                 vand.vv v24, v30, v19
  1924: d73c32a3                 vsrl.vi v25, v19, 4
  1928: 570dcc33                 vrgather.vv v26, v28, v24
  1932: d78ddc33                 vrgather.vv v27, v29, v25
  1936: 570dbd2f                 vxor.vv v26, v27, v26
  1940: d7090d5e                 vmv.v.v v19, v26
  1944: 570cea27                 vand.vv v24, v30, v20
  1948: d73c42a3                 vsrl.vi v25, v20, 4
  1952: 570dcc33                 vrgather.vv v26, v28, v24
  1956: d78ddc33                 vrgather.vv v27, v29, v25
  1960: 570dbd2f                 vxor.vv v26, v27, v26
  1964: 570a0d5e                 vmv.v.v v20, v26
  1968: 578cea27                 vand.vv v24, v30, v21
  1972: d73c52a3                 vsrl.vi v25, v21, 4
  1976: 570dcc33                 vrgather.vv v26, v28, v24
  1980: d78ddc33                 vrgather.vv v27, v29, v25
  1984: 570dbd2f                 vxor.vv v26, v27, v26
  1988: d70a0d5e                 vmv.v.v v21, v26
  1992: 570ceb27                 vand.vv v24, v30, v22
  1996: d73c62a3                 vsrl.vi v25, v22, 4
  2000: 570dcc33                 vrgather.vv v26, v28, v24
  2004: d78ddc33                 vrgather.vv v27, v29, v25
  2008: 570dbd2f                 vxor.vv v26, v27, v26
  2012: 570b0d5e                 vmv.v.v v22, v26
  2016: 578ceb27                 vand.vv v24, v30, v23
  2020: d73c72a3                 vsrl.vi v25, v23, 4
  2024: 570dcc33                 vrgather.vv v26, v28, v24
  2028: d78ddc33                 vrgather.vv v27, v29, v25
  2032: 570dbd2f                 vxor.vv v26, v27, v26
  2036: d70b0d5e                 vmv.v.v v23, v26
  2040: 138e0f0a                 addi t3, t6, 0xa0
  2044: 070e0e02                 vle8.v v28, (t3)
  2048: 138e0f0b                 addi t3, t6, 0xb0
  2052: 870e0e02                 vle8.v v29, (t3)
  2056: 570ce027                 vand.vv v24, v30, v0
  2060: d73c02a2                 vsrl.vi v25, v0, 4
  2064: 570dcc33                 vrgather.vv v26, v28, v24
  2068: d78ddc33                 vrgather.vv v27, v29, v25
  2072: 570dbd2f                 vxor.vv v26, v27, v26
  2076: 5704a42f                 vxor.vv v8, v26, v8
  2080: 578ce027                 vand.vv v24, v30, v1
  2084: d73c12a2                 vsrl.vi v25, v1, 4
  2088: 570dcc33                 vrgather.vv v26, v28, v24
  2092: d78ddc33                 vrgather.vv v27, v29, v25
  2096: 570dbd2f                 vxor.vv v26, v27, v26
  2100: d784a42f                 vxor.vv v9, v26, v9
  2104: 570ce127                 vand.vv v24, v30, v2
  2108: d73c22a2                 vsrl.vi v25, v2, 4
  2112: 570dcc33                 vrgather.vv v26, v28, v24
  2116: d78ddc33                 vrgather.vv v27, v29, v25
  2120: 570dbd2f                 vxor.vv v26, v27, v26
  2124: 5705a52f                 vxor.vv v10, v26, v10
  2128: 578ce127                 vand.vv v24, v30, v3
  2132: d73c32a2                 vsrl.vi v25, v3, 4
  2136: 570dcc33                 vrgather.vv v26, v28, v24
  2140: d78ddc33                 vrgather.vv v27, v29, v25
  2144: 570dbd2f                 vxor.vv v26, v27, v26
  2148: d785a52f                 vxor.vv v11, v26, v11
  2152: 570ce227                 vand.vv v24, v30, v4
  2156: d73c42a2                 vsrl.vi v25, v4, 4
  2160: 570dcc33                 vrgather.vv v26, v28, v24
  2164: d78ddc33                 vrgather.vv v27, v29, v25
  2168: 570dbd2f                 vxor.vv v26, v27, v26
  2172: 5706a62f                 vxor.vv v12, v26, v12
  2176: 578ce227                 vand.vv v24, v30, v5
  2180: d73c52a2                 vsrl.vi v25, v5, 4
  2184: 570dcc33                 vrgather.vv v26, v28, v24
  2188: d78ddc33                 vrgather.vv v27, v29, v25
  2192: 570dbd2f                 vxor.vv v26, v27, v26
  2196: d786a62f                 vxor.vv v13, v26, v13
  2200: 570ce327                 vand.vv v24, v30, v6
  2204: d73c62a2                 vsrl.vi v25, v6, 4
  2208: 570dcc33                 vrgather.vv v26, v28, v24
  2212: d78ddc33                 vrgather.vv v27, v29, v25
  2216: 570dbd2f                 vxor.vv v26, v27, v26
  2220: 5707a72f                 vxor.vv v14, v26, v14
  2224: 578ce327                 vand.vv v24, v30, v7
  2228: d73c72a2                 vsrl.vi v25, v7, 4
  2232: 570dcc33                 vrgather.vv v26, v28, v24
  2236: d78ddc33                 vrgather.vv v27, v29, v25
  2240: 570dbd2f                 vxor.vv v26, v27, v26
  2244: d787a72f                 vxor.vv v15, v26, v15
  2248: 138e0f08                 addi t3, t6, 0x80
  2252: 070e0e02                 vle8.v v28, (t3)
  2256: 138e0f09                 addi t3, t6, 0x90
  2260: 870e0e02                 vle8.v v29, (t3)
  2264: 570ce427                 vand.vv v24, v30, v8
  2268: d73c82a2                 vsrl.vi v25, v8, 4
  2272: 570dcc33                 vrgather.vv v26, v28, v24
  2276: d78ddc33                 vrgather.vv v27, v29, v25
  2280: 570dbd2f                 vxor.vv v26, v27, v26
  2284: 57040d5e                 vmv.v.v v8, v26
  2288: 578ce427                 vand.vv v24, v30, v9
  2292: d73c92a2                 vsrl.vi v25, v9, 4
  2296: 570dcc33                 vrgather.vv v26, v28, v24
  2300: d78ddc33                 vrgather.vv v27, v29, v25
  2304: 570dbd2f                 vxor.vv v26, v27, v26
  2308: d7040d5e                 vmv.v.v v9, v26
  2312: 570ce527                 vand.vv v24, v30, v10
  2316: d73ca2a2                 vsrl.vi v25, v10, 4
  2320: 570dcc33                 vrgather.vv v26, v28, v24
  2324: d78ddc33                 vrgather.vv v27, v29, v25
  2328: 570dbd2f                 vxor.vv v26, v27, v26
  2332: 57050d5e                 vmv.v.v v10, v26
  2336: 578ce527                 vand.vv v24, v30, v11
  2340: d73cb2a2                 vsrl.vi v25, v11, 4
  2344: 570dcc33                 vrgather.vv v26, v28, v24
  2348: d78ddc33                 vrgather.vv v27, v29, v25
  2352: 570dbd2f                 vxor.vv v26, v27, v26
  2356: d7050d5e                 vmv.v.v v11, v26
  2360: 570ce627                 vand.vv v24, v30, v12
  2364: d73cc2a2                 vsrl.vi v25, v12, 4
  2368: 570dcc33                 vrgather.vv v26, v28, v24
  2372: d78ddc33                 vrgather.vv v27, v29, v25
  2376: 570dbd2f                 vxor.vv v26, v27, v26
  2380: 57060d5e                 vmv.v.v v12, v26
  2384: 578ce627                 vand.vv v24, v30, v13
  2388: d73cd2a2                 vsrl.vi v25, v13, 4
  2392: 570dcc33                 vrgather.vv v26, v28, v24
  2396: d78ddc33                 vrgather.vv v27, v29, v25
  2400: 570dbd2f                 vxor.vv v26, v27, v26
  2404: d7060d5e                 vmv.v.v v13, v26
  2408: 570ce727                 vand.vv v24, v30, v14
  2412: d73ce2a2                 vsrl.vi v25, v14, 4
  2416: 570dcc33                 vrgather.vv v26, v28, v24
  2420: d78ddc33                 vrgather.vv v27, v29, v25
  2424: 570dbd2f                 vxor.vv v26, v27, v26
  2428: 57070d5e                 vmv.v.v v14, v26
  2432: 578ce727                 vand.vv v24, v30, v15
  2436: d73cf2a2                 vsrl.vi v25, v15, 4
  2440: 570dcc33                 vrgather.vv v26, v28, v24
  2444: d78ddc33                 vrgather.vv v27, v29, v25
  2448: 570dbd2f                 vxor.vv v26, v27, v26
  2452: d7070d5e                 vmv.v.v v15, v26
  2456: 83b30500                 ld t2, 0(a1)
  2460: b3836300                 add t2, t2, t1
  2464: 27880302                 vse8.v v16, (t2)
  2468: b3835300                 add t2, t2, t0
  2472: a7880302                 vse8.v v17, (t2)
  2476: b3835300                 add t2, t2, t0
  2480: 27890302                 vse8.v v18, (t2)
  2484: b3835300                 add t2, t2, t0
  2488: a7890302                 vse8.v v19, (t2)
  2492: b3835300                 add t2, t2, t0
  2496: 278a0302                 vse8.v v20, (t2)
  2500: b3835300                 add t2, t2, t0
  2504: a78a0302                 vse8.v v21, (t2)
  2508: b3835300                 add t2, t2, t0
  2512: 278b0302                 vse8.v v22, (t2)
  2516: b3835300                 add t2, t2, t0
  2520: a78b0302                 vse8.v v23, (t2)
  2524: 83b38500                 ld t2, 8(a1)
  2528: b3836300                 add t2, t2, t1
  2532: 27840302                 vse8.v v8, (t2)
  2536: b3835300                 add t2, t2, t0
  2540: a7840302                 vse8.v v9, (t2)
  2544: b3835300                 add t2, t2, t0
  2548: 27850302                 vse8.v v10, (t2)
  2552: b3835300                 add t2, t2, t0
  2556: a7850302                 vse8.v v11, (t2)
  2560: b3835300                 add t2, t2, t0
  2564: 27860302                 vse8.v v12, (t2)
  2568: b3835300                 add t2, t2, t0
  2572: a7860302                 vse8.v v13, (t2)
  2576: b3835300                 add t2, t2, t0
  2580: 27870302                 vse8.v v14, (t2)
  2584: b3835300                 add t2, t2, t0
  2588: a7870302                 vse8.v v15, (t2)
  2592: 3303d300                 add t1, t1, a3
  2596: 1307f7ff                 addi a4, a4, -1
  2600: 631c07e0                 bnez a4, 0x40
  2604: 33066640                 sub a2, a2, t1
  2608: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
  2612: 63940200                 bnez t0, 0xa3c
  2616: 6f00c022                 j 0xc64
  2620: 83330500                 ld t2, 0(a0)
  2624: b3836300                 add t2, t2, t1
  2628: 07880302                 vle8.v v16, (t2)
  2632: 5700085e                 vmv.v.v v0, v16
  2636: 83338500                 ld t2, 8(a0)
  2640: b3836300                 add t2, t2, t1
  2644: 07880302                 vle8.v v16, (t2)
  2648: 5700002f                 vxor.vv v0, v16, v0
  2652: 83330501                 ld t2, 0x10(a0)
  2656: b3836300                 add t2, t2, t1
  2660: 07880302                 vle8.v v16, (t2)
  2664: 5700002f                 vxor.vv v0, v16, v0
  2668: 83338501                 ld t2, 0x18(a0)
  2672: b3836300                 add t2, t2, t1
  2676: 07880302                 vle8.v v16, (t2)
  2680: 5700002f                 vxor.vv v0, v16, v0
  2684: 83330502                 ld t2, 0x20(a0)
  2688: b3836300                 add t2, t2, t1
  2692: 07880302                 vle8.v v16, (t2)
  2696: 5700002f                 vxor.vv v0, v16, v0
  2700: 83330500                 ld t2, 0(a0)
  2704: b3836300                 add t2, t2, t1
  2708: 138f0200                 mv t5, t0
  2712: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2716: 138e0f00                 mv t3, t6
  2720: 070e0e02                 vle8.v v28, (t3)
  2724: 138e0f01                 addi t3, t6, 0x10
  2728: 870e0e02                 vle8.v v29, (t3)
  2732: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2736: 07880302                 vle8.v v16, (t2)
  2740: 570ce827                 vand.vv v24, v30, v16
  2744: d73c02a3                 vsrl.vi v25, v16, 4
  2748: 570dcc33                 vrgather.vv v26, v28, v24
  2752: d78ddc33                 vrgather.vv v27, v29, v25
  2756: 570dbd2f                 vxor.vv v26, v27, v26
  2760: 57040d5e                 vmv.v.v v8, v26
  2764: 83338500                 ld t2, 8(a0)
  2768: b3836300                 add t2, t2, t1
  2772: 138f0200                 mv t5, t0
  2776: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2780: 138e0f02                 addi t3, t6, 0x20
  2784: 070e0e02                 vle8.v v28, (t3)
  2788: 138e0f03                 addi t3, t6, 0x30
  2792: 870e0e02                 vle8.v v29, (t3)
  2796: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2800: 07880302                 vle8.v v16, (t2)
  2804: 570ce827                 vand.vv v24, v30, v16
  2808: d73c02a3                 vsrl.vi v25, v16, 4
  2812: 570dcc33                 vrgather.vv v26, v28, v24
  2816: d78ddc33                 vrgather.vv v27, v29, v25
  2820: 570dbd2f                 vxor.vv v26, v27, v26
  2824: 5704a42f                 vxor.vv v8, v26, v8
  2828: 83330501                 ld t2, 0x10(a0)
  2832: b3836300                 add t2, t2, t1
  2836: 138f0200                 mv t5, t0
  2840: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2844: 138e0f04                 addi t3, t6, 0x40
  2848: 070e0e02                 vle8.v v28, (t3)
  2852: 138e0f05                 addi t3, t6, 0x50
  2856: 870e0e02                 vle8.v v29, (t3)
  2860: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2864: 07880302                 vle8.v v16, (t2)
  2868: 570ce827                 vand.vv v24, v30, v16
  2872: d73c02a3                 vsrl.vi v25, v16, 4
  2876: 570dcc33                 vrgather.vv v26, v28, v24
  2880: d78ddc33                 vrgather.vv v27, v29, v25
  2884: 570dbd2f                 vxor.vv v26, v27, v26
  2888: 5704a42f                 vxor.vv v8, v26, v8
  2892: 83338501                 ld t2, 0x18(a0)
  2896: b3836300                 add t2, t2, t1
  2900: 07880302                 vle8.v v16, (t2)
  2904: 5704042f                 vxor.vv v8, v16, v8
  2908: 83338502                 ld t2, 0x28(a0)
  2912: b3836300                 add t2, t2, t1
  2916: 07880302                 vle8.v v16, (t2)
  2920: 5704042f                 vxor.vv v8, v16, v8
  2924: 138f0200                 mv t5, t0
  2928: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2932: 138e0f06                 addi t3, t6, 0x60
  2936: 070e0e02                 vle8.v v28, (t3)
  2940: 138e0f07                 addi t3, t6, 0x70
  2944: 870e0e02                 vle8.v v29, (t3)
  2948: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2952: 570ce027                 vand.vv v24, v30, v0
  2956: d73c02a2                 vsrl.vi v25, v0, 4
  2960: 570dcc33                 vrgather.vv v26, v28, v24
  2964: d78ddc33                 vrgather.vv v27, v29, v25
  2968: 570dbd2f                 vxor.vv v26, v27, v26
  2972: 57080d5e                 vmv.v.v v16, v26
  2976: 5708882e                 vxor.vv v16, v8, v16
  2980: 138f0200                 mv t5, t0
  2984: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2988: 138e0f08                 addi t3, t6, 0x80
  2992: 070e0e02                 vle8.v v28, (t3)
  2996: 138e0f09                 addi t3, t6, 0x90
  3000: 870e0e02                 vle8.v v29, (t3)
  3004: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3008: 570ce827                 vand.vv v24, v30, v16
  3012: d73c02a3                 vsrl.vi v25, v16, 4
  3016: 570dcc33                 vrgather.vv v26, v28, v24
  3020: d78ddc33                 vrgather.vv v27, v29, v25
  3024: 570dbd2f                 vxor.vv v26, v27, v26
  3028: 57080d5e                 vmv.v.v v16, v26
  3032: 138f0200                 mv t5, t0
  3036: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3040: 138e0f0a                 addi t3, t6, 0xa0
  3044: 070e0e02                 vle8.v v28, (t3)
  3048: 138e0f0b                 addi t3, t6, 0xb0
  3052: 870e0e02                 vle8.v v29, (t3)
  3056: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3060: 570ce027                 vand.vv v24, v30, v0
  3064: d73c02a2                 vsrl.vi v25, v0, 4
  3068: 570dcc33                 vrgather.vv v26, v28, v24
  3072: d78ddc33                 vrgather.vv v27, v29, v25
  3076: 570dbd2f                 vxor.vv v26, v27, v26
  3080: 5704a42f                 vxor.vv v8, v26, v8
  3084: 138f0200                 mv t5, t0
  3088: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3092: 138e0f08                 addi t3, t6, 0x80
  3096: 070e0e02                 vle8.v v28, (t3)
  3100: 138e0f09                 addi t3, t6, 0x90
  3104: 870e0e02                 vle8.v v29, (t3)
  3108: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3112: 570ce427                 vand.vv v24, v30, v8
  3116: d73c82a2                 vsrl.vi v25, v8, 4
  3120: 570dcc33                 vrgather.vv v26, v28, v24
  3124: d78ddc33                 vrgather.vv v27, v29, v25
  3128: 570dbd2f                 vxor.vv v26, v27, v26
  3132: 57040d5e                 vmv.v.v v8, v26
  3136: 83b30500                 ld t2, 0(a1)
  3140: b3836300                 add t2, t2, t1
  3144: 27880302                 vse8.v v16, (t2)
  3148: 83b38500                 ld t2, 8(a1)
  3152: b3836300                 add t2, t2, t1
  3156: 27840302                 vse8.v v8, (t2)
  3160: 33035300                 add t1, t1, t0
  3164: 33065640                 sub a2, a2, t0
  3168: 6ff01fdd                 j 0xa30
  3172: 67800000                 ret 
  3176: 0000                     unimp 
  3178: 0000                     unimp 
  3180: 0000                     unimp 
  3182: 0000                     unimp 
