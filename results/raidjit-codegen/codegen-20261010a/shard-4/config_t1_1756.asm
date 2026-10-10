; config_t1_1756 target=rvv-zvbc32e twin=nt topo=raid6(8)/enc op=encode U=8
; code_size=3136 twins=["nt"] uops_per_byte=4.90625
     0: 63140600                 bnez a2, 8
     4: 6f00d042                 j 0xc30
     8: 971f0000                 auipc t6, 1
    12: 938f8fc3                 addi t6, t6, -0x3c8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574f0e5e                 vmv.v.x v30, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93068000                 li a3, 8
    40: b386d202                 mul a3, t0, a3
    44: 3357d602                 divu a4, a2, a3
    48: 63140700                 bnez a4, 0x38
    52: 6f00501d                 j 0xa08
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07880302                 vle8.v v16, (t2)
    76: b3835300                 add t2, t2, t0
    80: 87880302                 vle8.v v17, (t2)
    84: b3835300                 add t2, t2, t0
    88: 07890302                 vle8.v v18, (t2)
    92: b3835300                 add t2, t2, t0
    96: 87890302                 vle8.v v19, (t2)
   100: b3835300                 add t2, t2, t0
   104: 078a0302                 vle8.v v20, (t2)
   108: b3835300                 add t2, t2, t0
   112: 878a0302                 vle8.v v21, (t2)
   116: b3835300                 add t2, t2, t0
   120: 078b0302                 vle8.v v22, (t2)
   124: b3835300                 add t2, t2, t0
   128: 878b0302                 vle8.v v23, (t2)
   132: 5700085e                 vmv.v.v v0, v16
   136: d780085e                 vmv.v.v v1, v17
   140: 5701095e                 vmv.v.v v2, v18
   144: d781095e                 vmv.v.v v3, v19
   148: 57020a5e                 vmv.v.v v4, v20
   152: d7820a5e                 vmv.v.v v5, v21
   156: 57030b5e                 vmv.v.v v6, v22
   160: d7830b5e                 vmv.v.v v7, v23
   164: 138e0f00                 mv t3, t6
   168: 070e0e02                 vle8.v v28, (t3)
   172: 138e0f01                 addi t3, t6, 0x10
   176: 870e0e02                 vle8.v v29, (t3)
   180: 570ce827                 vand.vv v24, v30, v16
   184: d73c02a3                 vsrl.vi v25, v16, 4
   188: 570dcc33                 vrgather.vv v26, v28, v24
   192: d78ddc33                 vrgather.vv v27, v29, v25
   196: 570dbd2f                 vxor.vv v26, v27, v26
   200: 57040d5e                 vmv.v.v v8, v26
   204: 578ce827                 vand.vv v24, v30, v17
   208: d73c12a3                 vsrl.vi v25, v17, 4
   212: 570dcc33                 vrgather.vv v26, v28, v24
   216: d78ddc33                 vrgather.vv v27, v29, v25
   220: 570dbd2f                 vxor.vv v26, v27, v26
   224: d7040d5e                 vmv.v.v v9, v26
   228: 570ce927                 vand.vv v24, v30, v18
   232: d73c22a3                 vsrl.vi v25, v18, 4
   236: 570dcc33                 vrgather.vv v26, v28, v24
   240: d78ddc33                 vrgather.vv v27, v29, v25
   244: 570dbd2f                 vxor.vv v26, v27, v26
   248: 57050d5e                 vmv.v.v v10, v26
   252: 578ce927                 vand.vv v24, v30, v19
   256: d73c32a3                 vsrl.vi v25, v19, 4
   260: 570dcc33                 vrgather.vv v26, v28, v24
   264: d78ddc33                 vrgather.vv v27, v29, v25
   268: 570dbd2f                 vxor.vv v26, v27, v26
   272: d7050d5e                 vmv.v.v v11, v26
   276: 570cea27                 vand.vv v24, v30, v20
   280: d73c42a3                 vsrl.vi v25, v20, 4
   284: 570dcc33                 vrgather.vv v26, v28, v24
   288: d78ddc33                 vrgather.vv v27, v29, v25
   292: 570dbd2f                 vxor.vv v26, v27, v26
   296: 57060d5e                 vmv.v.v v12, v26
   300: 578cea27                 vand.vv v24, v30, v21
   304: d73c52a3                 vsrl.vi v25, v21, 4
   308: 570dcc33                 vrgather.vv v26, v28, v24
   312: d78ddc33                 vrgather.vv v27, v29, v25
   316: 570dbd2f                 vxor.vv v26, v27, v26
   320: d7060d5e                 vmv.v.v v13, v26
   324: 570ceb27                 vand.vv v24, v30, v22
   328: d73c62a3                 vsrl.vi v25, v22, 4
   332: 570dcc33                 vrgather.vv v26, v28, v24
   336: d78ddc33                 vrgather.vv v27, v29, v25
   340: 570dbd2f                 vxor.vv v26, v27, v26
   344: 57070d5e                 vmv.v.v v14, v26
   348: 578ceb27                 vand.vv v24, v30, v23
   352: d73c72a3                 vsrl.vi v25, v23, 4
   356: 570dcc33                 vrgather.vv v26, v28, v24
   360: d78ddc33                 vrgather.vv v27, v29, v25
   364: 570dbd2f                 vxor.vv v26, v27, v26
   368: d7070d5e                 vmv.v.v v15, v26
   372: 83338500                 ld t2, 8(a0)
   376: b3836300                 add t2, t2, t1
   380: 07880302                 vle8.v v16, (t2)
   384: b3835300                 add t2, t2, t0
   388: 87880302                 vle8.v v17, (t2)
   392: b3835300                 add t2, t2, t0
   396: 07890302                 vle8.v v18, (t2)
   400: b3835300                 add t2, t2, t0
   404: 87890302                 vle8.v v19, (t2)
   408: b3835300                 add t2, t2, t0
   412: 078a0302                 vle8.v v20, (t2)
   416: b3835300                 add t2, t2, t0
   420: 878a0302                 vle8.v v21, (t2)
   424: b3835300                 add t2, t2, t0
   428: 078b0302                 vle8.v v22, (t2)
   432: b3835300                 add t2, t2, t0
   436: 878b0302                 vle8.v v23, (t2)
   440: 5700002f                 vxor.vv v0, v16, v0
   444: d780102f                 vxor.vv v1, v17, v1
   448: 5701212f                 vxor.vv v2, v18, v2
   452: d781312f                 vxor.vv v3, v19, v3
   456: 5702422f                 vxor.vv v4, v20, v4
   460: d782522f                 vxor.vv v5, v21, v5
   464: 5703632f                 vxor.vv v6, v22, v6
   468: d783732f                 vxor.vv v7, v23, v7
   472: 138e0f02                 addi t3, t6, 0x20
   476: 070e0e02                 vle8.v v28, (t3)
   480: 138e0f03                 addi t3, t6, 0x30
   484: 870e0e02                 vle8.v v29, (t3)
   488: 570ce827                 vand.vv v24, v30, v16
   492: d73c02a3                 vsrl.vi v25, v16, 4
   496: 570dcc33                 vrgather.vv v26, v28, v24
   500: d78ddc33                 vrgather.vv v27, v29, v25
   504: 570dbd2f                 vxor.vv v26, v27, v26
   508: 5704a42f                 vxor.vv v8, v26, v8
   512: 578ce827                 vand.vv v24, v30, v17
   516: d73c12a3                 vsrl.vi v25, v17, 4
   520: 570dcc33                 vrgather.vv v26, v28, v24
   524: d78ddc33                 vrgather.vv v27, v29, v25
   528: 570dbd2f                 vxor.vv v26, v27, v26
   532: d784a42f                 vxor.vv v9, v26, v9
   536: 570ce927                 vand.vv v24, v30, v18
   540: d73c22a3                 vsrl.vi v25, v18, 4
   544: 570dcc33                 vrgather.vv v26, v28, v24
   548: d78ddc33                 vrgather.vv v27, v29, v25
   552: 570dbd2f                 vxor.vv v26, v27, v26
   556: 5705a52f                 vxor.vv v10, v26, v10
   560: 578ce927                 vand.vv v24, v30, v19
   564: d73c32a3                 vsrl.vi v25, v19, 4
   568: 570dcc33                 vrgather.vv v26, v28, v24
   572: d78ddc33                 vrgather.vv v27, v29, v25
   576: 570dbd2f                 vxor.vv v26, v27, v26
   580: d785a52f                 vxor.vv v11, v26, v11
   584: 570cea27                 vand.vv v24, v30, v20
   588: d73c42a3                 vsrl.vi v25, v20, 4
   592: 570dcc33                 vrgather.vv v26, v28, v24
   596: d78ddc33                 vrgather.vv v27, v29, v25
   600: 570dbd2f                 vxor.vv v26, v27, v26
   604: 5706a62f                 vxor.vv v12, v26, v12
   608: 578cea27                 vand.vv v24, v30, v21
   612: d73c52a3                 vsrl.vi v25, v21, 4
   616: 570dcc33                 vrgather.vv v26, v28, v24
   620: d78ddc33                 vrgather.vv v27, v29, v25
   624: 570dbd2f                 vxor.vv v26, v27, v26
   628: d786a62f                 vxor.vv v13, v26, v13
   632: 570ceb27                 vand.vv v24, v30, v22
   636: d73c62a3                 vsrl.vi v25, v22, 4
   640: 570dcc33                 vrgather.vv v26, v28, v24
   644: d78ddc33                 vrgather.vv v27, v29, v25
   648: 570dbd2f                 vxor.vv v26, v27, v26
   652: 5707a72f                 vxor.vv v14, v26, v14
   656: 578ceb27                 vand.vv v24, v30, v23
   660: d73c72a3                 vsrl.vi v25, v23, 4
   664: 570dcc33                 vrgather.vv v26, v28, v24
   668: d78ddc33                 vrgather.vv v27, v29, v25
   672: 570dbd2f                 vxor.vv v26, v27, v26
   676: d787a72f                 vxor.vv v15, v26, v15
   680: 83330501                 ld t2, 0x10(a0)
   684: b3836300                 add t2, t2, t1
   688: 07880302                 vle8.v v16, (t2)
   692: b3835300                 add t2, t2, t0
   696: 87880302                 vle8.v v17, (t2)
   700: b3835300                 add t2, t2, t0
   704: 07890302                 vle8.v v18, (t2)
   708: b3835300                 add t2, t2, t0
   712: 87890302                 vle8.v v19, (t2)
   716: b3835300                 add t2, t2, t0
   720: 078a0302                 vle8.v v20, (t2)
   724: b3835300                 add t2, t2, t0
   728: 878a0302                 vle8.v v21, (t2)
   732: b3835300                 add t2, t2, t0
   736: 078b0302                 vle8.v v22, (t2)
   740: b3835300                 add t2, t2, t0
   744: 878b0302                 vle8.v v23, (t2)
   748: 5700002f                 vxor.vv v0, v16, v0
   752: d780102f                 vxor.vv v1, v17, v1
   756: 5701212f                 vxor.vv v2, v18, v2
   760: d781312f                 vxor.vv v3, v19, v3
   764: 5702422f                 vxor.vv v4, v20, v4
   768: d782522f                 vxor.vv v5, v21, v5
   772: 5703632f                 vxor.vv v6, v22, v6
   776: d783732f                 vxor.vv v7, v23, v7
   780: 138e0f04                 addi t3, t6, 0x40
   784: 070e0e02                 vle8.v v28, (t3)
   788: 138e0f05                 addi t3, t6, 0x50
   792: 870e0e02                 vle8.v v29, (t3)
   796: 570ce827                 vand.vv v24, v30, v16
   800: d73c02a3                 vsrl.vi v25, v16, 4
   804: 570dcc33                 vrgather.vv v26, v28, v24
   808: d78ddc33                 vrgather.vv v27, v29, v25
   812: 570dbd2f                 vxor.vv v26, v27, v26
   816: 5704a42f                 vxor.vv v8, v26, v8
   820: 578ce827                 vand.vv v24, v30, v17
   824: d73c12a3                 vsrl.vi v25, v17, 4
   828: 570dcc33                 vrgather.vv v26, v28, v24
   832: d78ddc33                 vrgather.vv v27, v29, v25
   836: 570dbd2f                 vxor.vv v26, v27, v26
   840: d784a42f                 vxor.vv v9, v26, v9
   844: 570ce927                 vand.vv v24, v30, v18
   848: d73c22a3                 vsrl.vi v25, v18, 4
   852: 570dcc33                 vrgather.vv v26, v28, v24
   856: d78ddc33                 vrgather.vv v27, v29, v25
   860: 570dbd2f                 vxor.vv v26, v27, v26
   864: 5705a52f                 vxor.vv v10, v26, v10
   868: 578ce927                 vand.vv v24, v30, v19
   872: d73c32a3                 vsrl.vi v25, v19, 4
   876: 570dcc33                 vrgather.vv v26, v28, v24
   880: d78ddc33                 vrgather.vv v27, v29, v25
   884: 570dbd2f                 vxor.vv v26, v27, v26
   888: d785a52f                 vxor.vv v11, v26, v11
   892: 570cea27                 vand.vv v24, v30, v20
   896: d73c42a3                 vsrl.vi v25, v20, 4
   900: 570dcc33                 vrgather.vv v26, v28, v24
   904: d78ddc33                 vrgather.vv v27, v29, v25
   908: 570dbd2f                 vxor.vv v26, v27, v26
   912: 5706a62f                 vxor.vv v12, v26, v12
   916: 578cea27                 vand.vv v24, v30, v21
   920: d73c52a3                 vsrl.vi v25, v21, 4
   924: 570dcc33                 vrgather.vv v26, v28, v24
   928: d78ddc33                 vrgather.vv v27, v29, v25
   932: 570dbd2f                 vxor.vv v26, v27, v26
   936: d786a62f                 vxor.vv v13, v26, v13
   940: 570ceb27                 vand.vv v24, v30, v22
   944: d73c62a3                 vsrl.vi v25, v22, 4
   948: 570dcc33                 vrgather.vv v26, v28, v24
   952: d78ddc33                 vrgather.vv v27, v29, v25
   956: 570dbd2f                 vxor.vv v26, v27, v26
   960: 5707a72f                 vxor.vv v14, v26, v14
   964: 578ceb27                 vand.vv v24, v30, v23
   968: d73c72a3                 vsrl.vi v25, v23, 4
   972: 570dcc33                 vrgather.vv v26, v28, v24
   976: d78ddc33                 vrgather.vv v27, v29, v25
   980: 570dbd2f                 vxor.vv v26, v27, v26
   984: d787a72f                 vxor.vv v15, v26, v15
   988: 83338501                 ld t2, 0x18(a0)
   992: b3836300                 add t2, t2, t1
   996: 07880302                 vle8.v v16, (t2)
  1000: b3835300                 add t2, t2, t0
  1004: 87880302                 vle8.v v17, (t2)
  1008: b3835300                 add t2, t2, t0
  1012: 07890302                 vle8.v v18, (t2)
  1016: b3835300                 add t2, t2, t0
  1020: 87890302                 vle8.v v19, (t2)
  1024: b3835300                 add t2, t2, t0
  1028: 078a0302                 vle8.v v20, (t2)
  1032: b3835300                 add t2, t2, t0
  1036: 878a0302                 vle8.v v21, (t2)
  1040: b3835300                 add t2, t2, t0
  1044: 078b0302                 vle8.v v22, (t2)
  1048: b3835300                 add t2, t2, t0
  1052: 878b0302                 vle8.v v23, (t2)
  1056: 5700002f                 vxor.vv v0, v16, v0
  1060: d780102f                 vxor.vv v1, v17, v1
  1064: 5701212f                 vxor.vv v2, v18, v2
  1068: d781312f                 vxor.vv v3, v19, v3
  1072: 5702422f                 vxor.vv v4, v20, v4
  1076: d782522f                 vxor.vv v5, v21, v5
  1080: 5703632f                 vxor.vv v6, v22, v6
  1084: d783732f                 vxor.vv v7, v23, v7
  1088: 138e0f06                 addi t3, t6, 0x60
  1092: 070e0e02                 vle8.v v28, (t3)
  1096: 138e0f07                 addi t3, t6, 0x70
  1100: 870e0e02                 vle8.v v29, (t3)
  1104: 570ce827                 vand.vv v24, v30, v16
  1108: d73c02a3                 vsrl.vi v25, v16, 4
  1112: 570dcc33                 vrgather.vv v26, v28, v24
  1116: d78ddc33                 vrgather.vv v27, v29, v25
  1120: 570dbd2f                 vxor.vv v26, v27, v26
  1124: 5704a42f                 vxor.vv v8, v26, v8
  1128: 578ce827                 vand.vv v24, v30, v17
  1132: d73c12a3                 vsrl.vi v25, v17, 4
  1136: 570dcc33                 vrgather.vv v26, v28, v24
  1140: d78ddc33                 vrgather.vv v27, v29, v25
  1144: 570dbd2f                 vxor.vv v26, v27, v26
  1148: d784a42f                 vxor.vv v9, v26, v9
  1152: 570ce927                 vand.vv v24, v30, v18
  1156: d73c22a3                 vsrl.vi v25, v18, 4
  1160: 570dcc33                 vrgather.vv v26, v28, v24
  1164: d78ddc33                 vrgather.vv v27, v29, v25
  1168: 570dbd2f                 vxor.vv v26, v27, v26
  1172: 5705a52f                 vxor.vv v10, v26, v10
  1176: 578ce927                 vand.vv v24, v30, v19
  1180: d73c32a3                 vsrl.vi v25, v19, 4
  1184: 570dcc33                 vrgather.vv v26, v28, v24
  1188: d78ddc33                 vrgather.vv v27, v29, v25
  1192: 570dbd2f                 vxor.vv v26, v27, v26
  1196: d785a52f                 vxor.vv v11, v26, v11
  1200: 570cea27                 vand.vv v24, v30, v20
  1204: d73c42a3                 vsrl.vi v25, v20, 4
  1208: 570dcc33                 vrgather.vv v26, v28, v24
  1212: d78ddc33                 vrgather.vv v27, v29, v25
  1216: 570dbd2f                 vxor.vv v26, v27, v26
  1220: 5706a62f                 vxor.vv v12, v26, v12
  1224: 578cea27                 vand.vv v24, v30, v21
  1228: d73c52a3                 vsrl.vi v25, v21, 4
  1232: 570dcc33                 vrgather.vv v26, v28, v24
  1236: d78ddc33                 vrgather.vv v27, v29, v25
  1240: 570dbd2f                 vxor.vv v26, v27, v26
  1244: d786a62f                 vxor.vv v13, v26, v13
  1248: 570ceb27                 vand.vv v24, v30, v22
  1252: d73c62a3                 vsrl.vi v25, v22, 4
  1256: 570dcc33                 vrgather.vv v26, v28, v24
  1260: d78ddc33                 vrgather.vv v27, v29, v25
  1264: 570dbd2f                 vxor.vv v26, v27, v26
  1268: 5707a72f                 vxor.vv v14, v26, v14
  1272: 578ceb27                 vand.vv v24, v30, v23
  1276: d73c72a3                 vsrl.vi v25, v23, 4
  1280: 570dcc33                 vrgather.vv v26, v28, v24
  1284: d78ddc33                 vrgather.vv v27, v29, v25
  1288: 570dbd2f                 vxor.vv v26, v27, v26
  1292: d787a72f                 vxor.vv v15, v26, v15
  1296: 83330502                 ld t2, 0x20(a0)
  1300: b3836300                 add t2, t2, t1
  1304: 07880302                 vle8.v v16, (t2)
  1308: b3835300                 add t2, t2, t0
  1312: 87880302                 vle8.v v17, (t2)
  1316: b3835300                 add t2, t2, t0
  1320: 07890302                 vle8.v v18, (t2)
  1324: b3835300                 add t2, t2, t0
  1328: 87890302                 vle8.v v19, (t2)
  1332: b3835300                 add t2, t2, t0
  1336: 078a0302                 vle8.v v20, (t2)
  1340: b3835300                 add t2, t2, t0
  1344: 878a0302                 vle8.v v21, (t2)
  1348: b3835300                 add t2, t2, t0
  1352: 078b0302                 vle8.v v22, (t2)
  1356: b3835300                 add t2, t2, t0
  1360: 878b0302                 vle8.v v23, (t2)
  1364: 5700002f                 vxor.vv v0, v16, v0
  1368: d780102f                 vxor.vv v1, v17, v1
  1372: 5701212f                 vxor.vv v2, v18, v2
  1376: d781312f                 vxor.vv v3, v19, v3
  1380: 5702422f                 vxor.vv v4, v20, v4
  1384: d782522f                 vxor.vv v5, v21, v5
  1388: 5703632f                 vxor.vv v6, v22, v6
  1392: d783732f                 vxor.vv v7, v23, v7
  1396: 138e0f08                 addi t3, t6, 0x80
  1400: 070e0e02                 vle8.v v28, (t3)
  1404: 138e0f09                 addi t3, t6, 0x90
  1408: 870e0e02                 vle8.v v29, (t3)
  1412: 570ce827                 vand.vv v24, v30, v16
  1416: d73c02a3                 vsrl.vi v25, v16, 4
  1420: 570dcc33                 vrgather.vv v26, v28, v24
  1424: d78ddc33                 vrgather.vv v27, v29, v25
  1428: 570dbd2f                 vxor.vv v26, v27, v26
  1432: 5704a42f                 vxor.vv v8, v26, v8
  1436: 578ce827                 vand.vv v24, v30, v17
  1440: d73c12a3                 vsrl.vi v25, v17, 4
  1444: 570dcc33                 vrgather.vv v26, v28, v24
  1448: d78ddc33                 vrgather.vv v27, v29, v25
  1452: 570dbd2f                 vxor.vv v26, v27, v26
  1456: d784a42f                 vxor.vv v9, v26, v9
  1460: 570ce927                 vand.vv v24, v30, v18
  1464: d73c22a3                 vsrl.vi v25, v18, 4
  1468: 570dcc33                 vrgather.vv v26, v28, v24
  1472: d78ddc33                 vrgather.vv v27, v29, v25
  1476: 570dbd2f                 vxor.vv v26, v27, v26
  1480: 5705a52f                 vxor.vv v10, v26, v10
  1484: 578ce927                 vand.vv v24, v30, v19
  1488: d73c32a3                 vsrl.vi v25, v19, 4
  1492: 570dcc33                 vrgather.vv v26, v28, v24
  1496: d78ddc33                 vrgather.vv v27, v29, v25
  1500: 570dbd2f                 vxor.vv v26, v27, v26
  1504: d785a52f                 vxor.vv v11, v26, v11
  1508: 570cea27                 vand.vv v24, v30, v20
  1512: d73c42a3                 vsrl.vi v25, v20, 4
  1516: 570dcc33                 vrgather.vv v26, v28, v24
  1520: d78ddc33                 vrgather.vv v27, v29, v25
  1524: 570dbd2f                 vxor.vv v26, v27, v26
  1528: 5706a62f                 vxor.vv v12, v26, v12
  1532: 578cea27                 vand.vv v24, v30, v21
  1536: d73c52a3                 vsrl.vi v25, v21, 4
  1540: 570dcc33                 vrgather.vv v26, v28, v24
  1544: d78ddc33                 vrgather.vv v27, v29, v25
  1548: 570dbd2f                 vxor.vv v26, v27, v26
  1552: d786a62f                 vxor.vv v13, v26, v13
  1556: 570ceb27                 vand.vv v24, v30, v22
  1560: d73c62a3                 vsrl.vi v25, v22, 4
  1564: 570dcc33                 vrgather.vv v26, v28, v24
  1568: d78ddc33                 vrgather.vv v27, v29, v25
  1572: 570dbd2f                 vxor.vv v26, v27, v26
  1576: 5707a72f                 vxor.vv v14, v26, v14
  1580: 578ceb27                 vand.vv v24, v30, v23
  1584: d73c72a3                 vsrl.vi v25, v23, 4
  1588: 570dcc33                 vrgather.vv v26, v28, v24
  1592: d78ddc33                 vrgather.vv v27, v29, v25
  1596: 570dbd2f                 vxor.vv v26, v27, v26
  1600: d787a72f                 vxor.vv v15, v26, v15
  1604: 83338502                 ld t2, 0x28(a0)
  1608: b3836300                 add t2, t2, t1
  1612: 07880302                 vle8.v v16, (t2)
  1616: b3835300                 add t2, t2, t0
  1620: 87880302                 vle8.v v17, (t2)
  1624: b3835300                 add t2, t2, t0
  1628: 07890302                 vle8.v v18, (t2)
  1632: b3835300                 add t2, t2, t0
  1636: 87890302                 vle8.v v19, (t2)
  1640: b3835300                 add t2, t2, t0
  1644: 078a0302                 vle8.v v20, (t2)
  1648: b3835300                 add t2, t2, t0
  1652: 878a0302                 vle8.v v21, (t2)
  1656: b3835300                 add t2, t2, t0
  1660: 078b0302                 vle8.v v22, (t2)
  1664: b3835300                 add t2, t2, t0
  1668: 878b0302                 vle8.v v23, (t2)
  1672: 5700002f                 vxor.vv v0, v16, v0
  1676: d780102f                 vxor.vv v1, v17, v1
  1680: 5701212f                 vxor.vv v2, v18, v2
  1684: d781312f                 vxor.vv v3, v19, v3
  1688: 5702422f                 vxor.vv v4, v20, v4
  1692: d782522f                 vxor.vv v5, v21, v5
  1696: 5703632f                 vxor.vv v6, v22, v6
  1700: d783732f                 vxor.vv v7, v23, v7
  1704: 138e0f0a                 addi t3, t6, 0xa0
  1708: 070e0e02                 vle8.v v28, (t3)
  1712: 138e0f0b                 addi t3, t6, 0xb0
  1716: 870e0e02                 vle8.v v29, (t3)
  1720: 570ce827                 vand.vv v24, v30, v16
  1724: d73c02a3                 vsrl.vi v25, v16, 4
  1728: 570dcc33                 vrgather.vv v26, v28, v24
  1732: d78ddc33                 vrgather.vv v27, v29, v25
  1736: 570dbd2f                 vxor.vv v26, v27, v26
  1740: 5704a42f                 vxor.vv v8, v26, v8
  1744: 578ce827                 vand.vv v24, v30, v17
  1748: d73c12a3                 vsrl.vi v25, v17, 4
  1752: 570dcc33                 vrgather.vv v26, v28, v24
  1756: d78ddc33                 vrgather.vv v27, v29, v25
  1760: 570dbd2f                 vxor.vv v26, v27, v26
  1764: d784a42f                 vxor.vv v9, v26, v9
  1768: 570ce927                 vand.vv v24, v30, v18
  1772: d73c22a3                 vsrl.vi v25, v18, 4
  1776: 570dcc33                 vrgather.vv v26, v28, v24
  1780: d78ddc33                 vrgather.vv v27, v29, v25
  1784: 570dbd2f                 vxor.vv v26, v27, v26
  1788: 5705a52f                 vxor.vv v10, v26, v10
  1792: 578ce927                 vand.vv v24, v30, v19
  1796: d73c32a3                 vsrl.vi v25, v19, 4
  1800: 570dcc33                 vrgather.vv v26, v28, v24
  1804: d78ddc33                 vrgather.vv v27, v29, v25
  1808: 570dbd2f                 vxor.vv v26, v27, v26
  1812: d785a52f                 vxor.vv v11, v26, v11
  1816: 570cea27                 vand.vv v24, v30, v20
  1820: d73c42a3                 vsrl.vi v25, v20, 4
  1824: 570dcc33                 vrgather.vv v26, v28, v24
  1828: d78ddc33                 vrgather.vv v27, v29, v25
  1832: 570dbd2f                 vxor.vv v26, v27, v26
  1836: 5706a62f                 vxor.vv v12, v26, v12
  1840: 578cea27                 vand.vv v24, v30, v21
  1844: d73c52a3                 vsrl.vi v25, v21, 4
  1848: 570dcc33                 vrgather.vv v26, v28, v24
  1852: d78ddc33                 vrgather.vv v27, v29, v25
  1856: 570dbd2f                 vxor.vv v26, v27, v26
  1860: d786a62f                 vxor.vv v13, v26, v13
  1864: 570ceb27                 vand.vv v24, v30, v22
  1868: d73c62a3                 vsrl.vi v25, v22, 4
  1872: 570dcc33                 vrgather.vv v26, v28, v24
  1876: d78ddc33                 vrgather.vv v27, v29, v25
  1880: 570dbd2f                 vxor.vv v26, v27, v26
  1884: 5707a72f                 vxor.vv v14, v26, v14
  1888: 578ceb27                 vand.vv v24, v30, v23
  1892: d73c72a3                 vsrl.vi v25, v23, 4
  1896: 570dcc33                 vrgather.vv v26, v28, v24
  1900: d78ddc33                 vrgather.vv v27, v29, v25
  1904: 570dbd2f                 vxor.vv v26, v27, v26
  1908: d787a72f                 vxor.vv v15, v26, v15
  1912: 83330503                 ld t2, 0x30(a0)
  1916: b3836300                 add t2, t2, t1
  1920: 07880302                 vle8.v v16, (t2)
  1924: b3835300                 add t2, t2, t0
  1928: 87880302                 vle8.v v17, (t2)
  1932: b3835300                 add t2, t2, t0
  1936: 07890302                 vle8.v v18, (t2)
  1940: b3835300                 add t2, t2, t0
  1944: 87890302                 vle8.v v19, (t2)
  1948: b3835300                 add t2, t2, t0
  1952: 078a0302                 vle8.v v20, (t2)
  1956: b3835300                 add t2, t2, t0
  1960: 878a0302                 vle8.v v21, (t2)
  1964: b3835300                 add t2, t2, t0
  1968: 078b0302                 vle8.v v22, (t2)
  1972: b3835300                 add t2, t2, t0
  1976: 878b0302                 vle8.v v23, (t2)
  1980: 5700002f                 vxor.vv v0, v16, v0
  1984: d780102f                 vxor.vv v1, v17, v1
  1988: 5701212f                 vxor.vv v2, v18, v2
  1992: d781312f                 vxor.vv v3, v19, v3
  1996: 5702422f                 vxor.vv v4, v20, v4
  2000: d782522f                 vxor.vv v5, v21, v5
  2004: 5703632f                 vxor.vv v6, v22, v6
  2008: d783732f                 vxor.vv v7, v23, v7
  2012: 138e0f0c                 addi t3, t6, 0xc0
  2016: 070e0e02                 vle8.v v28, (t3)
  2020: 138e0f0d                 addi t3, t6, 0xd0
  2024: 870e0e02                 vle8.v v29, (t3)
  2028: 570ce827                 vand.vv v24, v30, v16
  2032: d73c02a3                 vsrl.vi v25, v16, 4
  2036: 570dcc33                 vrgather.vv v26, v28, v24
  2040: d78ddc33                 vrgather.vv v27, v29, v25
  2044: 570dbd2f                 vxor.vv v26, v27, v26
  2048: 5704a42f                 vxor.vv v8, v26, v8
  2052: 578ce827                 vand.vv v24, v30, v17
  2056: d73c12a3                 vsrl.vi v25, v17, 4
  2060: 570dcc33                 vrgather.vv v26, v28, v24
  2064: d78ddc33                 vrgather.vv v27, v29, v25
  2068: 570dbd2f                 vxor.vv v26, v27, v26
  2072: d784a42f                 vxor.vv v9, v26, v9
  2076: 570ce927                 vand.vv v24, v30, v18
  2080: d73c22a3                 vsrl.vi v25, v18, 4
  2084: 570dcc33                 vrgather.vv v26, v28, v24
  2088: d78ddc33                 vrgather.vv v27, v29, v25
  2092: 570dbd2f                 vxor.vv v26, v27, v26
  2096: 5705a52f                 vxor.vv v10, v26, v10
  2100: 578ce927                 vand.vv v24, v30, v19
  2104: d73c32a3                 vsrl.vi v25, v19, 4
  2108: 570dcc33                 vrgather.vv v26, v28, v24
  2112: d78ddc33                 vrgather.vv v27, v29, v25
  2116: 570dbd2f                 vxor.vv v26, v27, v26
  2120: d785a52f                 vxor.vv v11, v26, v11
  2124: 570cea27                 vand.vv v24, v30, v20
  2128: d73c42a3                 vsrl.vi v25, v20, 4
  2132: 570dcc33                 vrgather.vv v26, v28, v24
  2136: d78ddc33                 vrgather.vv v27, v29, v25
  2140: 570dbd2f                 vxor.vv v26, v27, v26
  2144: 5706a62f                 vxor.vv v12, v26, v12
  2148: 578cea27                 vand.vv v24, v30, v21
  2152: d73c52a3                 vsrl.vi v25, v21, 4
  2156: 570dcc33                 vrgather.vv v26, v28, v24
  2160: d78ddc33                 vrgather.vv v27, v29, v25
  2164: 570dbd2f                 vxor.vv v26, v27, v26
  2168: d786a62f                 vxor.vv v13, v26, v13
  2172: 570ceb27                 vand.vv v24, v30, v22
  2176: d73c62a3                 vsrl.vi v25, v22, 4
  2180: 570dcc33                 vrgather.vv v26, v28, v24
  2184: d78ddc33                 vrgather.vv v27, v29, v25
  2188: 570dbd2f                 vxor.vv v26, v27, v26
  2192: 5707a72f                 vxor.vv v14, v26, v14
  2196: 578ceb27                 vand.vv v24, v30, v23
  2200: d73c72a3                 vsrl.vi v25, v23, 4
  2204: 570dcc33                 vrgather.vv v26, v28, v24
  2208: d78ddc33                 vrgather.vv v27, v29, v25
  2212: 570dbd2f                 vxor.vv v26, v27, v26
  2216: d787a72f                 vxor.vv v15, v26, v15
  2220: 83338503                 ld t2, 0x38(a0)
  2224: b3836300                 add t2, t2, t1
  2228: 07880302                 vle8.v v16, (t2)
  2232: b3835300                 add t2, t2, t0
  2236: 87880302                 vle8.v v17, (t2)
  2240: b3835300                 add t2, t2, t0
  2244: 07890302                 vle8.v v18, (t2)
  2248: b3835300                 add t2, t2, t0
  2252: 87890302                 vle8.v v19, (t2)
  2256: b3835300                 add t2, t2, t0
  2260: 078a0302                 vle8.v v20, (t2)
  2264: b3835300                 add t2, t2, t0
  2268: 878a0302                 vle8.v v21, (t2)
  2272: b3835300                 add t2, t2, t0
  2276: 078b0302                 vle8.v v22, (t2)
  2280: b3835300                 add t2, t2, t0
  2284: 878b0302                 vle8.v v23, (t2)
  2288: 5700002f                 vxor.vv v0, v16, v0
  2292: d780102f                 vxor.vv v1, v17, v1
  2296: 5701212f                 vxor.vv v2, v18, v2
  2300: d781312f                 vxor.vv v3, v19, v3
  2304: 5702422f                 vxor.vv v4, v20, v4
  2308: d782522f                 vxor.vv v5, v21, v5
  2312: 5703632f                 vxor.vv v6, v22, v6
  2316: d783732f                 vxor.vv v7, v23, v7
  2320: 5704042f                 vxor.vv v8, v16, v8
  2324: d784142f                 vxor.vv v9, v17, v9
  2328: 5705252f                 vxor.vv v10, v18, v10
  2332: d785352f                 vxor.vv v11, v19, v11
  2336: 5706462f                 vxor.vv v12, v20, v12
  2340: d786562f                 vxor.vv v13, v21, v13
  2344: 5707672f                 vxor.vv v14, v22, v14
  2348: d787772f                 vxor.vv v15, v23, v15
  2352: 83b30500                 ld t2, 0(a1)
  2356: b3836300                 add t2, t2, t1
  2360: 33005000                 ntl.all 
  2364: 27800302                 vse8.v v0, (t2)
  2368: b3835300                 add t2, t2, t0
  2372: 33005000                 ntl.all 
  2376: a7800302                 vse8.v v1, (t2)
  2380: b3835300                 add t2, t2, t0
  2384: 33005000                 ntl.all 
  2388: 27810302                 vse8.v v2, (t2)
  2392: b3835300                 add t2, t2, t0
  2396: 33005000                 ntl.all 
  2400: a7810302                 vse8.v v3, (t2)
  2404: b3835300                 add t2, t2, t0
  2408: 33005000                 ntl.all 
  2412: 27820302                 vse8.v v4, (t2)
  2416: b3835300                 add t2, t2, t0
  2420: 33005000                 ntl.all 
  2424: a7820302                 vse8.v v5, (t2)
  2428: b3835300                 add t2, t2, t0
  2432: 33005000                 ntl.all 
  2436: 27830302                 vse8.v v6, (t2)
  2440: b3835300                 add t2, t2, t0
  2444: 33005000                 ntl.all 
  2448: a7830302                 vse8.v v7, (t2)
  2452: 83b38500                 ld t2, 8(a1)
  2456: b3836300                 add t2, t2, t1
  2460: 33005000                 ntl.all 
  2464: 27840302                 vse8.v v8, (t2)
  2468: b3835300                 add t2, t2, t0
  2472: 33005000                 ntl.all 
  2476: a7840302                 vse8.v v9, (t2)
  2480: b3835300                 add t2, t2, t0
  2484: 33005000                 ntl.all 
  2488: 27850302                 vse8.v v10, (t2)
  2492: b3835300                 add t2, t2, t0
  2496: 33005000                 ntl.all 
  2500: a7850302                 vse8.v v11, (t2)
  2504: b3835300                 add t2, t2, t0
  2508: 33005000                 ntl.all 
  2512: 27860302                 vse8.v v12, (t2)
  2516: b3835300                 add t2, t2, t0
  2520: 33005000                 ntl.all 
  2524: a7860302                 vse8.v v13, (t2)
  2528: b3835300                 add t2, t2, t0
  2532: 33005000                 ntl.all 
  2536: 27870302                 vse8.v v14, (t2)
  2540: b3835300                 add t2, t2, t0
  2544: 33005000                 ntl.all 
  2548: a7870302                 vse8.v v15, (t2)
  2552: 3303d300                 add t1, t1, a3
  2556: 1307f7ff                 addi a4, a4, -1
  2560: 631007e4                 bnez a4, 0x40
  2564: 33066640                 sub a2, a2, t1
  2568: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
  2572: 63940200                 bnez t0, 0xa14
  2576: 6f000022                 j 0xc30
  2580: 83330500                 ld t2, 0(a0)
  2584: b3836300                 add t2, t2, t1
  2588: 07880302                 vle8.v v16, (t2)
  2592: 5700085e                 vmv.v.v v0, v16
  2596: 138f0200                 mv t5, t0
  2600: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2604: 138e0f00                 mv t3, t6
  2608: 070e0e02                 vle8.v v28, (t3)
  2612: 138e0f01                 addi t3, t6, 0x10
  2616: 870e0e02                 vle8.v v29, (t3)
  2620: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2624: 570ce827                 vand.vv v24, v30, v16
  2628: d73c02a3                 vsrl.vi v25, v16, 4
  2632: 570dcc33                 vrgather.vv v26, v28, v24
  2636: d78ddc33                 vrgather.vv v27, v29, v25
  2640: 570dbd2f                 vxor.vv v26, v27, v26
  2644: 57040d5e                 vmv.v.v v8, v26
  2648: 83338500                 ld t2, 8(a0)
  2652: b3836300                 add t2, t2, t1
  2656: 07880302                 vle8.v v16, (t2)
  2660: 5700002f                 vxor.vv v0, v16, v0
  2664: 138f0200                 mv t5, t0
  2668: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2672: 138e0f02                 addi t3, t6, 0x20
  2676: 070e0e02                 vle8.v v28, (t3)
  2680: 138e0f03                 addi t3, t6, 0x30
  2684: 870e0e02                 vle8.v v29, (t3)
  2688: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2692: 570ce827                 vand.vv v24, v30, v16
  2696: d73c02a3                 vsrl.vi v25, v16, 4
  2700: 570dcc33                 vrgather.vv v26, v28, v24
  2704: d78ddc33                 vrgather.vv v27, v29, v25
  2708: 570dbd2f                 vxor.vv v26, v27, v26
  2712: 5704a42f                 vxor.vv v8, v26, v8
  2716: 83330501                 ld t2, 0x10(a0)
  2720: b3836300                 add t2, t2, t1
  2724: 07880302                 vle8.v v16, (t2)
  2728: 5700002f                 vxor.vv v0, v16, v0
  2732: 138f0200                 mv t5, t0
  2736: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2740: 138e0f04                 addi t3, t6, 0x40
  2744: 070e0e02                 vle8.v v28, (t3)
  2748: 138e0f05                 addi t3, t6, 0x50
  2752: 870e0e02                 vle8.v v29, (t3)
  2756: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2760: 570ce827                 vand.vv v24, v30, v16
  2764: d73c02a3                 vsrl.vi v25, v16, 4
  2768: 570dcc33                 vrgather.vv v26, v28, v24
  2772: d78ddc33                 vrgather.vv v27, v29, v25
  2776: 570dbd2f                 vxor.vv v26, v27, v26
  2780: 5704a42f                 vxor.vv v8, v26, v8
  2784: 83338501                 ld t2, 0x18(a0)
  2788: b3836300                 add t2, t2, t1
  2792: 07880302                 vle8.v v16, (t2)
  2796: 5700002f                 vxor.vv v0, v16, v0
  2800: 138f0200                 mv t5, t0
  2804: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2808: 138e0f06                 addi t3, t6, 0x60
  2812: 070e0e02                 vle8.v v28, (t3)
  2816: 138e0f07                 addi t3, t6, 0x70
  2820: 870e0e02                 vle8.v v29, (t3)
  2824: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2828: 570ce827                 vand.vv v24, v30, v16
  2832: d73c02a3                 vsrl.vi v25, v16, 4
  2836: 570dcc33                 vrgather.vv v26, v28, v24
  2840: d78ddc33                 vrgather.vv v27, v29, v25
  2844: 570dbd2f                 vxor.vv v26, v27, v26
  2848: 5704a42f                 vxor.vv v8, v26, v8
  2852: 83330502                 ld t2, 0x20(a0)
  2856: b3836300                 add t2, t2, t1
  2860: 07880302                 vle8.v v16, (t2)
  2864: 5700002f                 vxor.vv v0, v16, v0
  2868: 138f0200                 mv t5, t0
  2872: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2876: 138e0f08                 addi t3, t6, 0x80
  2880: 070e0e02                 vle8.v v28, (t3)
  2884: 138e0f09                 addi t3, t6, 0x90
  2888: 870e0e02                 vle8.v v29, (t3)
  2892: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2896: 570ce827                 vand.vv v24, v30, v16
  2900: d73c02a3                 vsrl.vi v25, v16, 4
  2904: 570dcc33                 vrgather.vv v26, v28, v24
  2908: d78ddc33                 vrgather.vv v27, v29, v25
  2912: 570dbd2f                 vxor.vv v26, v27, v26
  2916: 5704a42f                 vxor.vv v8, v26, v8
  2920: 83338502                 ld t2, 0x28(a0)
  2924: b3836300                 add t2, t2, t1
  2928: 07880302                 vle8.v v16, (t2)
  2932: 5700002f                 vxor.vv v0, v16, v0
  2936: 138f0200                 mv t5, t0
  2940: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2944: 138e0f0a                 addi t3, t6, 0xa0
  2948: 070e0e02                 vle8.v v28, (t3)
  2952: 138e0f0b                 addi t3, t6, 0xb0
  2956: 870e0e02                 vle8.v v29, (t3)
  2960: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2964: 570ce827                 vand.vv v24, v30, v16
  2968: d73c02a3                 vsrl.vi v25, v16, 4
  2972: 570dcc33                 vrgather.vv v26, v28, v24
  2976: d78ddc33                 vrgather.vv v27, v29, v25
  2980: 570dbd2f                 vxor.vv v26, v27, v26
  2984: 5704a42f                 vxor.vv v8, v26, v8
  2988: 83330503                 ld t2, 0x30(a0)
  2992: b3836300                 add t2, t2, t1
  2996: 07880302                 vle8.v v16, (t2)
  3000: 5700002f                 vxor.vv v0, v16, v0
  3004: 138f0200                 mv t5, t0
  3008: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3012: 138e0f0c                 addi t3, t6, 0xc0
  3016: 070e0e02                 vle8.v v28, (t3)
  3020: 138e0f0d                 addi t3, t6, 0xd0
  3024: 870e0e02                 vle8.v v29, (t3)
  3028: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3032: 570ce827                 vand.vv v24, v30, v16
  3036: d73c02a3                 vsrl.vi v25, v16, 4
  3040: 570dcc33                 vrgather.vv v26, v28, v24
  3044: d78ddc33                 vrgather.vv v27, v29, v25
  3048: 570dbd2f                 vxor.vv v26, v27, v26
  3052: 5704a42f                 vxor.vv v8, v26, v8
  3056: 83338503                 ld t2, 0x38(a0)
  3060: b3836300                 add t2, t2, t1
  3064: 07880302                 vle8.v v16, (t2)
  3068: 5700002f                 vxor.vv v0, v16, v0
  3072: 5704042f                 vxor.vv v8, v16, v8
  3076: 83b30500                 ld t2, 0(a1)
  3080: b3836300                 add t2, t2, t1
  3084: 33005000                 ntl.all 
  3088: 27800302                 vse8.v v0, (t2)
  3092: 83b38500                 ld t2, 8(a1)
  3096: b3836300                 add t2, t2, t1
  3100: 33005000                 ntl.all 
  3104: 27840302                 vse8.v v8, (t2)
  3108: 33035300                 add t1, t1, t0
  3112: 33065640                 sub a2, a2, t0
  3116: 6ff0dfdd                 j 0xa08
  3120: 67800000                 ret 
  3124: 0000                     unimp 
  3126: 0000                     unimp 
  3128: 0000                     unimp 
  3130: 0000                     unimp 
  3132: 0000                     unimp 
  3134: 0000                     unimp 
