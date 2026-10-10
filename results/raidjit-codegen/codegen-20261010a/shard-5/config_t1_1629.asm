; config_t1_1629 target=rvv twin=notail topo=rs(6,3)/rec{0,1} op=recover U=8
; code_size=3120 twins=["notail"] uops_per_byte=5.3515625
     0: 63140600                 bnez a2, 8
     4: 6f009041                 j 0xc1c
     8: 971f0000                 auipc t6, 1
    12: 938f8fc2                 addi t6, t6, -0x3d8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574f0e5e                 vmv.v.x v30, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93068000                 li a3, 8
    40: b386d202                 mul a3, t0, a3
    44: 3357d602                 divu a4, a2, a3
    48: 63140700                 bnez a4, 0x38
    52: 6f00903e                 j 0xc1c
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
   132: 138e0f00                 mv t3, t6
   136: 070e0e02                 vle8.v v28, (t3)
   140: 138e0f01                 addi t3, t6, 0x10
   144: 870e0e02                 vle8.v v29, (t3)
   148: 570ce827                 vand.vv v24, v30, v16
   152: d73c02a3                 vsrl.vi v25, v16, 4
   156: 570dcc33                 vrgather.vv v26, v28, v24
   160: d78ddc33                 vrgather.vv v27, v29, v25
   164: 570dbd2f                 vxor.vv v26, v27, v26
   168: 57000d5e                 vmv.v.v v0, v26
   172: 578ce827                 vand.vv v24, v30, v17
   176: d73c12a3                 vsrl.vi v25, v17, 4
   180: 570dcc33                 vrgather.vv v26, v28, v24
   184: d78ddc33                 vrgather.vv v27, v29, v25
   188: 570dbd2f                 vxor.vv v26, v27, v26
   192: d7000d5e                 vmv.v.v v1, v26
   196: 570ce927                 vand.vv v24, v30, v18
   200: d73c22a3                 vsrl.vi v25, v18, 4
   204: 570dcc33                 vrgather.vv v26, v28, v24
   208: d78ddc33                 vrgather.vv v27, v29, v25
   212: 570dbd2f                 vxor.vv v26, v27, v26
   216: 57010d5e                 vmv.v.v v2, v26
   220: 578ce927                 vand.vv v24, v30, v19
   224: d73c32a3                 vsrl.vi v25, v19, 4
   228: 570dcc33                 vrgather.vv v26, v28, v24
   232: d78ddc33                 vrgather.vv v27, v29, v25
   236: 570dbd2f                 vxor.vv v26, v27, v26
   240: d7010d5e                 vmv.v.v v3, v26
   244: 570cea27                 vand.vv v24, v30, v20
   248: d73c42a3                 vsrl.vi v25, v20, 4
   252: 570dcc33                 vrgather.vv v26, v28, v24
   256: d78ddc33                 vrgather.vv v27, v29, v25
   260: 570dbd2f                 vxor.vv v26, v27, v26
   264: 57020d5e                 vmv.v.v v4, v26
   268: 578cea27                 vand.vv v24, v30, v21
   272: d73c52a3                 vsrl.vi v25, v21, 4
   276: 570dcc33                 vrgather.vv v26, v28, v24
   280: d78ddc33                 vrgather.vv v27, v29, v25
   284: 570dbd2f                 vxor.vv v26, v27, v26
   288: d7020d5e                 vmv.v.v v5, v26
   292: 570ceb27                 vand.vv v24, v30, v22
   296: d73c62a3                 vsrl.vi v25, v22, 4
   300: 570dcc33                 vrgather.vv v26, v28, v24
   304: d78ddc33                 vrgather.vv v27, v29, v25
   308: 570dbd2f                 vxor.vv v26, v27, v26
   312: 57030d5e                 vmv.v.v v6, v26
   316: 578ceb27                 vand.vv v24, v30, v23
   320: d73c72a3                 vsrl.vi v25, v23, 4
   324: 570dcc33                 vrgather.vv v26, v28, v24
   328: d78ddc33                 vrgather.vv v27, v29, v25
   332: 570dbd2f                 vxor.vv v26, v27, v26
   336: d7030d5e                 vmv.v.v v7, v26
   340: 138e0f02                 addi t3, t6, 0x20
   344: 070e0e02                 vle8.v v28, (t3)
   348: 138e0f03                 addi t3, t6, 0x30
   352: 870e0e02                 vle8.v v29, (t3)
   356: 570ce827                 vand.vv v24, v30, v16
   360: d73c02a3                 vsrl.vi v25, v16, 4
   364: 570dcc33                 vrgather.vv v26, v28, v24
   368: d78ddc33                 vrgather.vv v27, v29, v25
   372: 570dbd2f                 vxor.vv v26, v27, v26
   376: 57040d5e                 vmv.v.v v8, v26
   380: 578ce827                 vand.vv v24, v30, v17
   384: d73c12a3                 vsrl.vi v25, v17, 4
   388: 570dcc33                 vrgather.vv v26, v28, v24
   392: d78ddc33                 vrgather.vv v27, v29, v25
   396: 570dbd2f                 vxor.vv v26, v27, v26
   400: d7040d5e                 vmv.v.v v9, v26
   404: 570ce927                 vand.vv v24, v30, v18
   408: d73c22a3                 vsrl.vi v25, v18, 4
   412: 570dcc33                 vrgather.vv v26, v28, v24
   416: d78ddc33                 vrgather.vv v27, v29, v25
   420: 570dbd2f                 vxor.vv v26, v27, v26
   424: 57050d5e                 vmv.v.v v10, v26
   428: 578ce927                 vand.vv v24, v30, v19
   432: d73c32a3                 vsrl.vi v25, v19, 4
   436: 570dcc33                 vrgather.vv v26, v28, v24
   440: d78ddc33                 vrgather.vv v27, v29, v25
   444: 570dbd2f                 vxor.vv v26, v27, v26
   448: d7050d5e                 vmv.v.v v11, v26
   452: 570cea27                 vand.vv v24, v30, v20
   456: d73c42a3                 vsrl.vi v25, v20, 4
   460: 570dcc33                 vrgather.vv v26, v28, v24
   464: d78ddc33                 vrgather.vv v27, v29, v25
   468: 570dbd2f                 vxor.vv v26, v27, v26
   472: 57060d5e                 vmv.v.v v12, v26
   476: 578cea27                 vand.vv v24, v30, v21
   480: d73c52a3                 vsrl.vi v25, v21, 4
   484: 570dcc33                 vrgather.vv v26, v28, v24
   488: d78ddc33                 vrgather.vv v27, v29, v25
   492: 570dbd2f                 vxor.vv v26, v27, v26
   496: d7060d5e                 vmv.v.v v13, v26
   500: 570ceb27                 vand.vv v24, v30, v22
   504: d73c62a3                 vsrl.vi v25, v22, 4
   508: 570dcc33                 vrgather.vv v26, v28, v24
   512: d78ddc33                 vrgather.vv v27, v29, v25
   516: 570dbd2f                 vxor.vv v26, v27, v26
   520: 57070d5e                 vmv.v.v v14, v26
   524: 578ceb27                 vand.vv v24, v30, v23
   528: d73c72a3                 vsrl.vi v25, v23, 4
   532: 570dcc33                 vrgather.vv v26, v28, v24
   536: d78ddc33                 vrgather.vv v27, v29, v25
   540: 570dbd2f                 vxor.vv v26, v27, v26
   544: d7070d5e                 vmv.v.v v15, v26
   548: 83338500                 ld t2, 8(a0)
   552: b3836300                 add t2, t2, t1
   556: 07880302                 vle8.v v16, (t2)
   560: b3835300                 add t2, t2, t0
   564: 87880302                 vle8.v v17, (t2)
   568: b3835300                 add t2, t2, t0
   572: 07890302                 vle8.v v18, (t2)
   576: b3835300                 add t2, t2, t0
   580: 87890302                 vle8.v v19, (t2)
   584: b3835300                 add t2, t2, t0
   588: 078a0302                 vle8.v v20, (t2)
   592: b3835300                 add t2, t2, t0
   596: 878a0302                 vle8.v v21, (t2)
   600: b3835300                 add t2, t2, t0
   604: 078b0302                 vle8.v v22, (t2)
   608: b3835300                 add t2, t2, t0
   612: 878b0302                 vle8.v v23, (t2)
   616: 138e0f04                 addi t3, t6, 0x40
   620: 070e0e02                 vle8.v v28, (t3)
   624: 138e0f05                 addi t3, t6, 0x50
   628: 870e0e02                 vle8.v v29, (t3)
   632: 570ce827                 vand.vv v24, v30, v16
   636: d73c02a3                 vsrl.vi v25, v16, 4
   640: 570dcc33                 vrgather.vv v26, v28, v24
   644: d78ddc33                 vrgather.vv v27, v29, v25
   648: 570dbd2f                 vxor.vv v26, v27, v26
   652: 5700a02f                 vxor.vv v0, v26, v0
   656: 578ce827                 vand.vv v24, v30, v17
   660: d73c12a3                 vsrl.vi v25, v17, 4
   664: 570dcc33                 vrgather.vv v26, v28, v24
   668: d78ddc33                 vrgather.vv v27, v29, v25
   672: 570dbd2f                 vxor.vv v26, v27, v26
   676: d780a02f                 vxor.vv v1, v26, v1
   680: 570ce927                 vand.vv v24, v30, v18
   684: d73c22a3                 vsrl.vi v25, v18, 4
   688: 570dcc33                 vrgather.vv v26, v28, v24
   692: d78ddc33                 vrgather.vv v27, v29, v25
   696: 570dbd2f                 vxor.vv v26, v27, v26
   700: 5701a12f                 vxor.vv v2, v26, v2
   704: 578ce927                 vand.vv v24, v30, v19
   708: d73c32a3                 vsrl.vi v25, v19, 4
   712: 570dcc33                 vrgather.vv v26, v28, v24
   716: d78ddc33                 vrgather.vv v27, v29, v25
   720: 570dbd2f                 vxor.vv v26, v27, v26
   724: d781a12f                 vxor.vv v3, v26, v3
   728: 570cea27                 vand.vv v24, v30, v20
   732: d73c42a3                 vsrl.vi v25, v20, 4
   736: 570dcc33                 vrgather.vv v26, v28, v24
   740: d78ddc33                 vrgather.vv v27, v29, v25
   744: 570dbd2f                 vxor.vv v26, v27, v26
   748: 5702a22f                 vxor.vv v4, v26, v4
   752: 578cea27                 vand.vv v24, v30, v21
   756: d73c52a3                 vsrl.vi v25, v21, 4
   760: 570dcc33                 vrgather.vv v26, v28, v24
   764: d78ddc33                 vrgather.vv v27, v29, v25
   768: 570dbd2f                 vxor.vv v26, v27, v26
   772: d782a22f                 vxor.vv v5, v26, v5
   776: 570ceb27                 vand.vv v24, v30, v22
   780: d73c62a3                 vsrl.vi v25, v22, 4
   784: 570dcc33                 vrgather.vv v26, v28, v24
   788: d78ddc33                 vrgather.vv v27, v29, v25
   792: 570dbd2f                 vxor.vv v26, v27, v26
   796: 5703a32f                 vxor.vv v6, v26, v6
   800: 578ceb27                 vand.vv v24, v30, v23
   804: d73c72a3                 vsrl.vi v25, v23, 4
   808: 570dcc33                 vrgather.vv v26, v28, v24
   812: d78ddc33                 vrgather.vv v27, v29, v25
   816: 570dbd2f                 vxor.vv v26, v27, v26
   820: d783a32f                 vxor.vv v7, v26, v7
   824: 138e0f06                 addi t3, t6, 0x60
   828: 070e0e02                 vle8.v v28, (t3)
   832: 138e0f07                 addi t3, t6, 0x70
   836: 870e0e02                 vle8.v v29, (t3)
   840: 570ce827                 vand.vv v24, v30, v16
   844: d73c02a3                 vsrl.vi v25, v16, 4
   848: 570dcc33                 vrgather.vv v26, v28, v24
   852: d78ddc33                 vrgather.vv v27, v29, v25
   856: 570dbd2f                 vxor.vv v26, v27, v26
   860: 5704a42f                 vxor.vv v8, v26, v8
   864: 578ce827                 vand.vv v24, v30, v17
   868: d73c12a3                 vsrl.vi v25, v17, 4
   872: 570dcc33                 vrgather.vv v26, v28, v24
   876: d78ddc33                 vrgather.vv v27, v29, v25
   880: 570dbd2f                 vxor.vv v26, v27, v26
   884: d784a42f                 vxor.vv v9, v26, v9
   888: 570ce927                 vand.vv v24, v30, v18
   892: d73c22a3                 vsrl.vi v25, v18, 4
   896: 570dcc33                 vrgather.vv v26, v28, v24
   900: d78ddc33                 vrgather.vv v27, v29, v25
   904: 570dbd2f                 vxor.vv v26, v27, v26
   908: 5705a52f                 vxor.vv v10, v26, v10
   912: 578ce927                 vand.vv v24, v30, v19
   916: d73c32a3                 vsrl.vi v25, v19, 4
   920: 570dcc33                 vrgather.vv v26, v28, v24
   924: d78ddc33                 vrgather.vv v27, v29, v25
   928: 570dbd2f                 vxor.vv v26, v27, v26
   932: d785a52f                 vxor.vv v11, v26, v11
   936: 570cea27                 vand.vv v24, v30, v20
   940: d73c42a3                 vsrl.vi v25, v20, 4
   944: 570dcc33                 vrgather.vv v26, v28, v24
   948: d78ddc33                 vrgather.vv v27, v29, v25
   952: 570dbd2f                 vxor.vv v26, v27, v26
   956: 5706a62f                 vxor.vv v12, v26, v12
   960: 578cea27                 vand.vv v24, v30, v21
   964: d73c52a3                 vsrl.vi v25, v21, 4
   968: 570dcc33                 vrgather.vv v26, v28, v24
   972: d78ddc33                 vrgather.vv v27, v29, v25
   976: 570dbd2f                 vxor.vv v26, v27, v26
   980: d786a62f                 vxor.vv v13, v26, v13
   984: 570ceb27                 vand.vv v24, v30, v22
   988: d73c62a3                 vsrl.vi v25, v22, 4
   992: 570dcc33                 vrgather.vv v26, v28, v24
   996: d78ddc33                 vrgather.vv v27, v29, v25
  1000: 570dbd2f                 vxor.vv v26, v27, v26
  1004: 5707a72f                 vxor.vv v14, v26, v14
  1008: 578ceb27                 vand.vv v24, v30, v23
  1012: d73c72a3                 vsrl.vi v25, v23, 4
  1016: 570dcc33                 vrgather.vv v26, v28, v24
  1020: d78ddc33                 vrgather.vv v27, v29, v25
  1024: 570dbd2f                 vxor.vv v26, v27, v26
  1028: d787a72f                 vxor.vv v15, v26, v15
  1032: 83330501                 ld t2, 0x10(a0)
  1036: b3836300                 add t2, t2, t1
  1040: 07880302                 vle8.v v16, (t2)
  1044: b3835300                 add t2, t2, t0
  1048: 87880302                 vle8.v v17, (t2)
  1052: b3835300                 add t2, t2, t0
  1056: 07890302                 vle8.v v18, (t2)
  1060: b3835300                 add t2, t2, t0
  1064: 87890302                 vle8.v v19, (t2)
  1068: b3835300                 add t2, t2, t0
  1072: 078a0302                 vle8.v v20, (t2)
  1076: b3835300                 add t2, t2, t0
  1080: 878a0302                 vle8.v v21, (t2)
  1084: b3835300                 add t2, t2, t0
  1088: 078b0302                 vle8.v v22, (t2)
  1092: b3835300                 add t2, t2, t0
  1096: 878b0302                 vle8.v v23, (t2)
  1100: 138e0f08                 addi t3, t6, 0x80
  1104: 070e0e02                 vle8.v v28, (t3)
  1108: 138e0f09                 addi t3, t6, 0x90
  1112: 870e0e02                 vle8.v v29, (t3)
  1116: 570ce827                 vand.vv v24, v30, v16
  1120: d73c02a3                 vsrl.vi v25, v16, 4
  1124: 570dcc33                 vrgather.vv v26, v28, v24
  1128: d78ddc33                 vrgather.vv v27, v29, v25
  1132: 570dbd2f                 vxor.vv v26, v27, v26
  1136: 5700a02f                 vxor.vv v0, v26, v0
  1140: 578ce827                 vand.vv v24, v30, v17
  1144: d73c12a3                 vsrl.vi v25, v17, 4
  1148: 570dcc33                 vrgather.vv v26, v28, v24
  1152: d78ddc33                 vrgather.vv v27, v29, v25
  1156: 570dbd2f                 vxor.vv v26, v27, v26
  1160: d780a02f                 vxor.vv v1, v26, v1
  1164: 570ce927                 vand.vv v24, v30, v18
  1168: d73c22a3                 vsrl.vi v25, v18, 4
  1172: 570dcc33                 vrgather.vv v26, v28, v24
  1176: d78ddc33                 vrgather.vv v27, v29, v25
  1180: 570dbd2f                 vxor.vv v26, v27, v26
  1184: 5701a12f                 vxor.vv v2, v26, v2
  1188: 578ce927                 vand.vv v24, v30, v19
  1192: d73c32a3                 vsrl.vi v25, v19, 4
  1196: 570dcc33                 vrgather.vv v26, v28, v24
  1200: d78ddc33                 vrgather.vv v27, v29, v25
  1204: 570dbd2f                 vxor.vv v26, v27, v26
  1208: d781a12f                 vxor.vv v3, v26, v3
  1212: 570cea27                 vand.vv v24, v30, v20
  1216: d73c42a3                 vsrl.vi v25, v20, 4
  1220: 570dcc33                 vrgather.vv v26, v28, v24
  1224: d78ddc33                 vrgather.vv v27, v29, v25
  1228: 570dbd2f                 vxor.vv v26, v27, v26
  1232: 5702a22f                 vxor.vv v4, v26, v4
  1236: 578cea27                 vand.vv v24, v30, v21
  1240: d73c52a3                 vsrl.vi v25, v21, 4
  1244: 570dcc33                 vrgather.vv v26, v28, v24
  1248: d78ddc33                 vrgather.vv v27, v29, v25
  1252: 570dbd2f                 vxor.vv v26, v27, v26
  1256: d782a22f                 vxor.vv v5, v26, v5
  1260: 570ceb27                 vand.vv v24, v30, v22
  1264: d73c62a3                 vsrl.vi v25, v22, 4
  1268: 570dcc33                 vrgather.vv v26, v28, v24
  1272: d78ddc33                 vrgather.vv v27, v29, v25
  1276: 570dbd2f                 vxor.vv v26, v27, v26
  1280: 5703a32f                 vxor.vv v6, v26, v6
  1284: 578ceb27                 vand.vv v24, v30, v23
  1288: d73c72a3                 vsrl.vi v25, v23, 4
  1292: 570dcc33                 vrgather.vv v26, v28, v24
  1296: d78ddc33                 vrgather.vv v27, v29, v25
  1300: 570dbd2f                 vxor.vv v26, v27, v26
  1304: d783a32f                 vxor.vv v7, v26, v7
  1308: 138e0f0a                 addi t3, t6, 0xa0
  1312: 070e0e02                 vle8.v v28, (t3)
  1316: 138e0f0b                 addi t3, t6, 0xb0
  1320: 870e0e02                 vle8.v v29, (t3)
  1324: 570ce827                 vand.vv v24, v30, v16
  1328: d73c02a3                 vsrl.vi v25, v16, 4
  1332: 570dcc33                 vrgather.vv v26, v28, v24
  1336: d78ddc33                 vrgather.vv v27, v29, v25
  1340: 570dbd2f                 vxor.vv v26, v27, v26
  1344: 5704a42f                 vxor.vv v8, v26, v8
  1348: 578ce827                 vand.vv v24, v30, v17
  1352: d73c12a3                 vsrl.vi v25, v17, 4
  1356: 570dcc33                 vrgather.vv v26, v28, v24
  1360: d78ddc33                 vrgather.vv v27, v29, v25
  1364: 570dbd2f                 vxor.vv v26, v27, v26
  1368: d784a42f                 vxor.vv v9, v26, v9
  1372: 570ce927                 vand.vv v24, v30, v18
  1376: d73c22a3                 vsrl.vi v25, v18, 4
  1380: 570dcc33                 vrgather.vv v26, v28, v24
  1384: d78ddc33                 vrgather.vv v27, v29, v25
  1388: 570dbd2f                 vxor.vv v26, v27, v26
  1392: 5705a52f                 vxor.vv v10, v26, v10
  1396: 578ce927                 vand.vv v24, v30, v19
  1400: d73c32a3                 vsrl.vi v25, v19, 4
  1404: 570dcc33                 vrgather.vv v26, v28, v24
  1408: d78ddc33                 vrgather.vv v27, v29, v25
  1412: 570dbd2f                 vxor.vv v26, v27, v26
  1416: d785a52f                 vxor.vv v11, v26, v11
  1420: 570cea27                 vand.vv v24, v30, v20
  1424: d73c42a3                 vsrl.vi v25, v20, 4
  1428: 570dcc33                 vrgather.vv v26, v28, v24
  1432: d78ddc33                 vrgather.vv v27, v29, v25
  1436: 570dbd2f                 vxor.vv v26, v27, v26
  1440: 5706a62f                 vxor.vv v12, v26, v12
  1444: 578cea27                 vand.vv v24, v30, v21
  1448: d73c52a3                 vsrl.vi v25, v21, 4
  1452: 570dcc33                 vrgather.vv v26, v28, v24
  1456: d78ddc33                 vrgather.vv v27, v29, v25
  1460: 570dbd2f                 vxor.vv v26, v27, v26
  1464: d786a62f                 vxor.vv v13, v26, v13
  1468: 570ceb27                 vand.vv v24, v30, v22
  1472: d73c62a3                 vsrl.vi v25, v22, 4
  1476: 570dcc33                 vrgather.vv v26, v28, v24
  1480: d78ddc33                 vrgather.vv v27, v29, v25
  1484: 570dbd2f                 vxor.vv v26, v27, v26
  1488: 5707a72f                 vxor.vv v14, v26, v14
  1492: 578ceb27                 vand.vv v24, v30, v23
  1496: d73c72a3                 vsrl.vi v25, v23, 4
  1500: 570dcc33                 vrgather.vv v26, v28, v24
  1504: d78ddc33                 vrgather.vv v27, v29, v25
  1508: 570dbd2f                 vxor.vv v26, v27, v26
  1512: d787a72f                 vxor.vv v15, v26, v15
  1516: 83338501                 ld t2, 0x18(a0)
  1520: b3836300                 add t2, t2, t1
  1524: 07880302                 vle8.v v16, (t2)
  1528: b3835300                 add t2, t2, t0
  1532: 87880302                 vle8.v v17, (t2)
  1536: b3835300                 add t2, t2, t0
  1540: 07890302                 vle8.v v18, (t2)
  1544: b3835300                 add t2, t2, t0
  1548: 87890302                 vle8.v v19, (t2)
  1552: b3835300                 add t2, t2, t0
  1556: 078a0302                 vle8.v v20, (t2)
  1560: b3835300                 add t2, t2, t0
  1564: 878a0302                 vle8.v v21, (t2)
  1568: b3835300                 add t2, t2, t0
  1572: 078b0302                 vle8.v v22, (t2)
  1576: b3835300                 add t2, t2, t0
  1580: 878b0302                 vle8.v v23, (t2)
  1584: 138e0f0c                 addi t3, t6, 0xc0
  1588: 070e0e02                 vle8.v v28, (t3)
  1592: 138e0f0d                 addi t3, t6, 0xd0
  1596: 870e0e02                 vle8.v v29, (t3)
  1600: 570ce827                 vand.vv v24, v30, v16
  1604: d73c02a3                 vsrl.vi v25, v16, 4
  1608: 570dcc33                 vrgather.vv v26, v28, v24
  1612: d78ddc33                 vrgather.vv v27, v29, v25
  1616: 570dbd2f                 vxor.vv v26, v27, v26
  1620: 5700a02f                 vxor.vv v0, v26, v0
  1624: 578ce827                 vand.vv v24, v30, v17
  1628: d73c12a3                 vsrl.vi v25, v17, 4
  1632: 570dcc33                 vrgather.vv v26, v28, v24
  1636: d78ddc33                 vrgather.vv v27, v29, v25
  1640: 570dbd2f                 vxor.vv v26, v27, v26
  1644: d780a02f                 vxor.vv v1, v26, v1
  1648: 570ce927                 vand.vv v24, v30, v18
  1652: d73c22a3                 vsrl.vi v25, v18, 4
  1656: 570dcc33                 vrgather.vv v26, v28, v24
  1660: d78ddc33                 vrgather.vv v27, v29, v25
  1664: 570dbd2f                 vxor.vv v26, v27, v26
  1668: 5701a12f                 vxor.vv v2, v26, v2
  1672: 578ce927                 vand.vv v24, v30, v19
  1676: d73c32a3                 vsrl.vi v25, v19, 4
  1680: 570dcc33                 vrgather.vv v26, v28, v24
  1684: d78ddc33                 vrgather.vv v27, v29, v25
  1688: 570dbd2f                 vxor.vv v26, v27, v26
  1692: d781a12f                 vxor.vv v3, v26, v3
  1696: 570cea27                 vand.vv v24, v30, v20
  1700: d73c42a3                 vsrl.vi v25, v20, 4
  1704: 570dcc33                 vrgather.vv v26, v28, v24
  1708: d78ddc33                 vrgather.vv v27, v29, v25
  1712: 570dbd2f                 vxor.vv v26, v27, v26
  1716: 5702a22f                 vxor.vv v4, v26, v4
  1720: 578cea27                 vand.vv v24, v30, v21
  1724: d73c52a3                 vsrl.vi v25, v21, 4
  1728: 570dcc33                 vrgather.vv v26, v28, v24
  1732: d78ddc33                 vrgather.vv v27, v29, v25
  1736: 570dbd2f                 vxor.vv v26, v27, v26
  1740: d782a22f                 vxor.vv v5, v26, v5
  1744: 570ceb27                 vand.vv v24, v30, v22
  1748: d73c62a3                 vsrl.vi v25, v22, 4
  1752: 570dcc33                 vrgather.vv v26, v28, v24
  1756: d78ddc33                 vrgather.vv v27, v29, v25
  1760: 570dbd2f                 vxor.vv v26, v27, v26
  1764: 5703a32f                 vxor.vv v6, v26, v6
  1768: 578ceb27                 vand.vv v24, v30, v23
  1772: d73c72a3                 vsrl.vi v25, v23, 4
  1776: 570dcc33                 vrgather.vv v26, v28, v24
  1780: d78ddc33                 vrgather.vv v27, v29, v25
  1784: 570dbd2f                 vxor.vv v26, v27, v26
  1788: d783a32f                 vxor.vv v7, v26, v7
  1792: 138e0f0e                 addi t3, t6, 0xe0
  1796: 070e0e02                 vle8.v v28, (t3)
  1800: 138e0f0f                 addi t3, t6, 0xf0
  1804: 870e0e02                 vle8.v v29, (t3)
  1808: 570ce827                 vand.vv v24, v30, v16
  1812: d73c02a3                 vsrl.vi v25, v16, 4
  1816: 570dcc33                 vrgather.vv v26, v28, v24
  1820: d78ddc33                 vrgather.vv v27, v29, v25
  1824: 570dbd2f                 vxor.vv v26, v27, v26
  1828: 5704a42f                 vxor.vv v8, v26, v8
  1832: 578ce827                 vand.vv v24, v30, v17
  1836: d73c12a3                 vsrl.vi v25, v17, 4
  1840: 570dcc33                 vrgather.vv v26, v28, v24
  1844: d78ddc33                 vrgather.vv v27, v29, v25
  1848: 570dbd2f                 vxor.vv v26, v27, v26
  1852: d784a42f                 vxor.vv v9, v26, v9
  1856: 570ce927                 vand.vv v24, v30, v18
  1860: d73c22a3                 vsrl.vi v25, v18, 4
  1864: 570dcc33                 vrgather.vv v26, v28, v24
  1868: d78ddc33                 vrgather.vv v27, v29, v25
  1872: 570dbd2f                 vxor.vv v26, v27, v26
  1876: 5705a52f                 vxor.vv v10, v26, v10
  1880: 578ce927                 vand.vv v24, v30, v19
  1884: d73c32a3                 vsrl.vi v25, v19, 4
  1888: 570dcc33                 vrgather.vv v26, v28, v24
  1892: d78ddc33                 vrgather.vv v27, v29, v25
  1896: 570dbd2f                 vxor.vv v26, v27, v26
  1900: d785a52f                 vxor.vv v11, v26, v11
  1904: 570cea27                 vand.vv v24, v30, v20
  1908: d73c42a3                 vsrl.vi v25, v20, 4
  1912: 570dcc33                 vrgather.vv v26, v28, v24
  1916: d78ddc33                 vrgather.vv v27, v29, v25
  1920: 570dbd2f                 vxor.vv v26, v27, v26
  1924: 5706a62f                 vxor.vv v12, v26, v12
  1928: 578cea27                 vand.vv v24, v30, v21
  1932: d73c52a3                 vsrl.vi v25, v21, 4
  1936: 570dcc33                 vrgather.vv v26, v28, v24
  1940: d78ddc33                 vrgather.vv v27, v29, v25
  1944: 570dbd2f                 vxor.vv v26, v27, v26
  1948: d786a62f                 vxor.vv v13, v26, v13
  1952: 570ceb27                 vand.vv v24, v30, v22
  1956: d73c62a3                 vsrl.vi v25, v22, 4
  1960: 570dcc33                 vrgather.vv v26, v28, v24
  1964: d78ddc33                 vrgather.vv v27, v29, v25
  1968: 570dbd2f                 vxor.vv v26, v27, v26
  1972: 5707a72f                 vxor.vv v14, v26, v14
  1976: 578ceb27                 vand.vv v24, v30, v23
  1980: d73c72a3                 vsrl.vi v25, v23, 4
  1984: 570dcc33                 vrgather.vv v26, v28, v24
  1988: d78ddc33                 vrgather.vv v27, v29, v25
  1992: 570dbd2f                 vxor.vv v26, v27, v26
  1996: d787a72f                 vxor.vv v15, v26, v15
  2000: 83330502                 ld t2, 0x20(a0)
  2004: b3836300                 add t2, t2, t1
  2008: 07880302                 vle8.v v16, (t2)
  2012: b3835300                 add t2, t2, t0
  2016: 87880302                 vle8.v v17, (t2)
  2020: b3835300                 add t2, t2, t0
  2024: 07890302                 vle8.v v18, (t2)
  2028: b3835300                 add t2, t2, t0
  2032: 87890302                 vle8.v v19, (t2)
  2036: b3835300                 add t2, t2, t0
  2040: 078a0302                 vle8.v v20, (t2)
  2044: b3835300                 add t2, t2, t0
  2048: 878a0302                 vle8.v v21, (t2)
  2052: b3835300                 add t2, t2, t0
  2056: 078b0302                 vle8.v v22, (t2)
  2060: b3835300                 add t2, t2, t0
  2064: 878b0302                 vle8.v v23, (t2)
  2068: 138e0f10                 addi t3, t6, 0x100
  2072: 070e0e02                 vle8.v v28, (t3)
  2076: 138e0f11                 addi t3, t6, 0x110
  2080: 870e0e02                 vle8.v v29, (t3)
  2084: 570ce827                 vand.vv v24, v30, v16
  2088: d73c02a3                 vsrl.vi v25, v16, 4
  2092: 570dcc33                 vrgather.vv v26, v28, v24
  2096: d78ddc33                 vrgather.vv v27, v29, v25
  2100: 570dbd2f                 vxor.vv v26, v27, v26
  2104: 5700a02f                 vxor.vv v0, v26, v0
  2108: 578ce827                 vand.vv v24, v30, v17
  2112: d73c12a3                 vsrl.vi v25, v17, 4
  2116: 570dcc33                 vrgather.vv v26, v28, v24
  2120: d78ddc33                 vrgather.vv v27, v29, v25
  2124: 570dbd2f                 vxor.vv v26, v27, v26
  2128: d780a02f                 vxor.vv v1, v26, v1
  2132: 570ce927                 vand.vv v24, v30, v18
  2136: d73c22a3                 vsrl.vi v25, v18, 4
  2140: 570dcc33                 vrgather.vv v26, v28, v24
  2144: d78ddc33                 vrgather.vv v27, v29, v25
  2148: 570dbd2f                 vxor.vv v26, v27, v26
  2152: 5701a12f                 vxor.vv v2, v26, v2
  2156: 578ce927                 vand.vv v24, v30, v19
  2160: d73c32a3                 vsrl.vi v25, v19, 4
  2164: 570dcc33                 vrgather.vv v26, v28, v24
  2168: d78ddc33                 vrgather.vv v27, v29, v25
  2172: 570dbd2f                 vxor.vv v26, v27, v26
  2176: d781a12f                 vxor.vv v3, v26, v3
  2180: 570cea27                 vand.vv v24, v30, v20
  2184: d73c42a3                 vsrl.vi v25, v20, 4
  2188: 570dcc33                 vrgather.vv v26, v28, v24
  2192: d78ddc33                 vrgather.vv v27, v29, v25
  2196: 570dbd2f                 vxor.vv v26, v27, v26
  2200: 5702a22f                 vxor.vv v4, v26, v4
  2204: 578cea27                 vand.vv v24, v30, v21
  2208: d73c52a3                 vsrl.vi v25, v21, 4
  2212: 570dcc33                 vrgather.vv v26, v28, v24
  2216: d78ddc33                 vrgather.vv v27, v29, v25
  2220: 570dbd2f                 vxor.vv v26, v27, v26
  2224: d782a22f                 vxor.vv v5, v26, v5
  2228: 570ceb27                 vand.vv v24, v30, v22
  2232: d73c62a3                 vsrl.vi v25, v22, 4
  2236: 570dcc33                 vrgather.vv v26, v28, v24
  2240: d78ddc33                 vrgather.vv v27, v29, v25
  2244: 570dbd2f                 vxor.vv v26, v27, v26
  2248: 5703a32f                 vxor.vv v6, v26, v6
  2252: 578ceb27                 vand.vv v24, v30, v23
  2256: d73c72a3                 vsrl.vi v25, v23, 4
  2260: 570dcc33                 vrgather.vv v26, v28, v24
  2264: d78ddc33                 vrgather.vv v27, v29, v25
  2268: 570dbd2f                 vxor.vv v26, v27, v26
  2272: d783a32f                 vxor.vv v7, v26, v7
  2276: 138e0f12                 addi t3, t6, 0x120
  2280: 070e0e02                 vle8.v v28, (t3)
  2284: 138e0f13                 addi t3, t6, 0x130
  2288: 870e0e02                 vle8.v v29, (t3)
  2292: 570ce827                 vand.vv v24, v30, v16
  2296: d73c02a3                 vsrl.vi v25, v16, 4
  2300: 570dcc33                 vrgather.vv v26, v28, v24
  2304: d78ddc33                 vrgather.vv v27, v29, v25
  2308: 570dbd2f                 vxor.vv v26, v27, v26
  2312: 5704a42f                 vxor.vv v8, v26, v8
  2316: 578ce827                 vand.vv v24, v30, v17
  2320: d73c12a3                 vsrl.vi v25, v17, 4
  2324: 570dcc33                 vrgather.vv v26, v28, v24
  2328: d78ddc33                 vrgather.vv v27, v29, v25
  2332: 570dbd2f                 vxor.vv v26, v27, v26
  2336: d784a42f                 vxor.vv v9, v26, v9
  2340: 570ce927                 vand.vv v24, v30, v18
  2344: d73c22a3                 vsrl.vi v25, v18, 4
  2348: 570dcc33                 vrgather.vv v26, v28, v24
  2352: d78ddc33                 vrgather.vv v27, v29, v25
  2356: 570dbd2f                 vxor.vv v26, v27, v26
  2360: 5705a52f                 vxor.vv v10, v26, v10
  2364: 578ce927                 vand.vv v24, v30, v19
  2368: d73c32a3                 vsrl.vi v25, v19, 4
  2372: 570dcc33                 vrgather.vv v26, v28, v24
  2376: d78ddc33                 vrgather.vv v27, v29, v25
  2380: 570dbd2f                 vxor.vv v26, v27, v26
  2384: d785a52f                 vxor.vv v11, v26, v11
  2388: 570cea27                 vand.vv v24, v30, v20
  2392: d73c42a3                 vsrl.vi v25, v20, 4
  2396: 570dcc33                 vrgather.vv v26, v28, v24
  2400: d78ddc33                 vrgather.vv v27, v29, v25
  2404: 570dbd2f                 vxor.vv v26, v27, v26
  2408: 5706a62f                 vxor.vv v12, v26, v12
  2412: 578cea27                 vand.vv v24, v30, v21
  2416: d73c52a3                 vsrl.vi v25, v21, 4
  2420: 570dcc33                 vrgather.vv v26, v28, v24
  2424: d78ddc33                 vrgather.vv v27, v29, v25
  2428: 570dbd2f                 vxor.vv v26, v27, v26
  2432: d786a62f                 vxor.vv v13, v26, v13
  2436: 570ceb27                 vand.vv v24, v30, v22
  2440: d73c62a3                 vsrl.vi v25, v22, 4
  2444: 570dcc33                 vrgather.vv v26, v28, v24
  2448: d78ddc33                 vrgather.vv v27, v29, v25
  2452: 570dbd2f                 vxor.vv v26, v27, v26
  2456: 5707a72f                 vxor.vv v14, v26, v14
  2460: 578ceb27                 vand.vv v24, v30, v23
  2464: d73c72a3                 vsrl.vi v25, v23, 4
  2468: 570dcc33                 vrgather.vv v26, v28, v24
  2472: d78ddc33                 vrgather.vv v27, v29, v25
  2476: 570dbd2f                 vxor.vv v26, v27, v26
  2480: d787a72f                 vxor.vv v15, v26, v15
  2484: 83338502                 ld t2, 0x28(a0)
  2488: b3836300                 add t2, t2, t1
  2492: 07880302                 vle8.v v16, (t2)
  2496: b3835300                 add t2, t2, t0
  2500: 87880302                 vle8.v v17, (t2)
  2504: b3835300                 add t2, t2, t0
  2508: 07890302                 vle8.v v18, (t2)
  2512: b3835300                 add t2, t2, t0
  2516: 87890302                 vle8.v v19, (t2)
  2520: b3835300                 add t2, t2, t0
  2524: 078a0302                 vle8.v v20, (t2)
  2528: b3835300                 add t2, t2, t0
  2532: 878a0302                 vle8.v v21, (t2)
  2536: b3835300                 add t2, t2, t0
  2540: 078b0302                 vle8.v v22, (t2)
  2544: b3835300                 add t2, t2, t0
  2548: 878b0302                 vle8.v v23, (t2)
  2552: 138e0f14                 addi t3, t6, 0x140
  2556: 070e0e02                 vle8.v v28, (t3)
  2560: 138e0f15                 addi t3, t6, 0x150
  2564: 870e0e02                 vle8.v v29, (t3)
  2568: 570ce827                 vand.vv v24, v30, v16
  2572: d73c02a3                 vsrl.vi v25, v16, 4
  2576: 570dcc33                 vrgather.vv v26, v28, v24
  2580: d78ddc33                 vrgather.vv v27, v29, v25
  2584: 570dbd2f                 vxor.vv v26, v27, v26
  2588: 5700a02f                 vxor.vv v0, v26, v0
  2592: 578ce827                 vand.vv v24, v30, v17
  2596: d73c12a3                 vsrl.vi v25, v17, 4
  2600: 570dcc33                 vrgather.vv v26, v28, v24
  2604: d78ddc33                 vrgather.vv v27, v29, v25
  2608: 570dbd2f                 vxor.vv v26, v27, v26
  2612: d780a02f                 vxor.vv v1, v26, v1
  2616: 570ce927                 vand.vv v24, v30, v18
  2620: d73c22a3                 vsrl.vi v25, v18, 4
  2624: 570dcc33                 vrgather.vv v26, v28, v24
  2628: d78ddc33                 vrgather.vv v27, v29, v25
  2632: 570dbd2f                 vxor.vv v26, v27, v26
  2636: 5701a12f                 vxor.vv v2, v26, v2
  2640: 578ce927                 vand.vv v24, v30, v19
  2644: d73c32a3                 vsrl.vi v25, v19, 4
  2648: 570dcc33                 vrgather.vv v26, v28, v24
  2652: d78ddc33                 vrgather.vv v27, v29, v25
  2656: 570dbd2f                 vxor.vv v26, v27, v26
  2660: d781a12f                 vxor.vv v3, v26, v3
  2664: 570cea27                 vand.vv v24, v30, v20
  2668: d73c42a3                 vsrl.vi v25, v20, 4
  2672: 570dcc33                 vrgather.vv v26, v28, v24
  2676: d78ddc33                 vrgather.vv v27, v29, v25
  2680: 570dbd2f                 vxor.vv v26, v27, v26
  2684: 5702a22f                 vxor.vv v4, v26, v4
  2688: 578cea27                 vand.vv v24, v30, v21
  2692: d73c52a3                 vsrl.vi v25, v21, 4
  2696: 570dcc33                 vrgather.vv v26, v28, v24
  2700: d78ddc33                 vrgather.vv v27, v29, v25
  2704: 570dbd2f                 vxor.vv v26, v27, v26
  2708: d782a22f                 vxor.vv v5, v26, v5
  2712: 570ceb27                 vand.vv v24, v30, v22
  2716: d73c62a3                 vsrl.vi v25, v22, 4
  2720: 570dcc33                 vrgather.vv v26, v28, v24
  2724: d78ddc33                 vrgather.vv v27, v29, v25
  2728: 570dbd2f                 vxor.vv v26, v27, v26
  2732: 5703a32f                 vxor.vv v6, v26, v6
  2736: 578ceb27                 vand.vv v24, v30, v23
  2740: d73c72a3                 vsrl.vi v25, v23, 4
  2744: 570dcc33                 vrgather.vv v26, v28, v24
  2748: d78ddc33                 vrgather.vv v27, v29, v25
  2752: 570dbd2f                 vxor.vv v26, v27, v26
  2756: d783a32f                 vxor.vv v7, v26, v7
  2760: 570ce827                 vand.vv v24, v30, v16
  2764: d73c02a3                 vsrl.vi v25, v16, 4
  2768: 570dcc33                 vrgather.vv v26, v28, v24
  2772: d78ddc33                 vrgather.vv v27, v29, v25
  2776: 570dbd2f                 vxor.vv v26, v27, v26
  2780: 5704a42f                 vxor.vv v8, v26, v8
  2784: 578ce827                 vand.vv v24, v30, v17
  2788: d73c12a3                 vsrl.vi v25, v17, 4
  2792: 570dcc33                 vrgather.vv v26, v28, v24
  2796: d78ddc33                 vrgather.vv v27, v29, v25
  2800: 570dbd2f                 vxor.vv v26, v27, v26
  2804: d784a42f                 vxor.vv v9, v26, v9
  2808: 570ce927                 vand.vv v24, v30, v18
  2812: d73c22a3                 vsrl.vi v25, v18, 4
  2816: 570dcc33                 vrgather.vv v26, v28, v24
  2820: d78ddc33                 vrgather.vv v27, v29, v25
  2824: 570dbd2f                 vxor.vv v26, v27, v26
  2828: 5705a52f                 vxor.vv v10, v26, v10
  2832: 578ce927                 vand.vv v24, v30, v19
  2836: d73c32a3                 vsrl.vi v25, v19, 4
  2840: 570dcc33                 vrgather.vv v26, v28, v24
  2844: d78ddc33                 vrgather.vv v27, v29, v25
  2848: 570dbd2f                 vxor.vv v26, v27, v26
  2852: d785a52f                 vxor.vv v11, v26, v11
  2856: 570cea27                 vand.vv v24, v30, v20
  2860: d73c42a3                 vsrl.vi v25, v20, 4
  2864: 570dcc33                 vrgather.vv v26, v28, v24
  2868: d78ddc33                 vrgather.vv v27, v29, v25
  2872: 570dbd2f                 vxor.vv v26, v27, v26
  2876: 5706a62f                 vxor.vv v12, v26, v12
  2880: 578cea27                 vand.vv v24, v30, v21
  2884: d73c52a3                 vsrl.vi v25, v21, 4
  2888: 570dcc33                 vrgather.vv v26, v28, v24
  2892: d78ddc33                 vrgather.vv v27, v29, v25
  2896: 570dbd2f                 vxor.vv v26, v27, v26
  2900: d786a62f                 vxor.vv v13, v26, v13
  2904: 570ceb27                 vand.vv v24, v30, v22
  2908: d73c62a3                 vsrl.vi v25, v22, 4
  2912: 570dcc33                 vrgather.vv v26, v28, v24
  2916: d78ddc33                 vrgather.vv v27, v29, v25
  2920: 570dbd2f                 vxor.vv v26, v27, v26
  2924: 5707a72f                 vxor.vv v14, v26, v14
  2928: 578ceb27                 vand.vv v24, v30, v23
  2932: d73c72a3                 vsrl.vi v25, v23, 4
  2936: 570dcc33                 vrgather.vv v26, v28, v24
  2940: d78ddc33                 vrgather.vv v27, v29, v25
  2944: 570dbd2f                 vxor.vv v26, v27, v26
  2948: d787a72f                 vxor.vv v15, v26, v15
  2952: 83b30500                 ld t2, 0(a1)
  2956: b3836300                 add t2, t2, t1
  2960: 27800302                 vse8.v v0, (t2)
  2964: b3835300                 add t2, t2, t0
  2968: a7800302                 vse8.v v1, (t2)
  2972: b3835300                 add t2, t2, t0
  2976: 27810302                 vse8.v v2, (t2)
  2980: b3835300                 add t2, t2, t0
  2984: a7810302                 vse8.v v3, (t2)
  2988: b3835300                 add t2, t2, t0
  2992: 27820302                 vse8.v v4, (t2)
  2996: b3835300                 add t2, t2, t0
  3000: a7820302                 vse8.v v5, (t2)
  3004: b3835300                 add t2, t2, t0
  3008: 27830302                 vse8.v v6, (t2)
  3012: b3835300                 add t2, t2, t0
  3016: a7830302                 vse8.v v7, (t2)
  3020: 83b38500                 ld t2, 8(a1)
  3024: b3836300                 add t2, t2, t1
  3028: 27840302                 vse8.v v8, (t2)
  3032: b3835300                 add t2, t2, t0
  3036: a7840302                 vse8.v v9, (t2)
  3040: b3835300                 add t2, t2, t0
  3044: 27850302                 vse8.v v10, (t2)
  3048: b3835300                 add t2, t2, t0
  3052: a7850302                 vse8.v v11, (t2)
  3056: b3835300                 add t2, t2, t0
  3060: 27860302                 vse8.v v12, (t2)
  3064: b3835300                 add t2, t2, t0
  3068: a7860302                 vse8.v v13, (t2)
  3072: b3835300                 add t2, t2, t0
  3076: 27870302                 vse8.v v14, (t2)
  3080: b3835300                 add t2, t2, t0
  3084: a7870302                 vse8.v v15, (t2)
  3088: 3303d300                 add t1, t1, a3
  3092: 1307f7ff                 addi a4, a4, -1
  3096: 631407c2                 bnez a4, 0x40
  3100: 67800000                 ret 
  3104: 0000                     unimp 
  3106: 0000                     unimp 
  3108: 0000                     unimp 
  3110: 0000                     unimp 
  3112: 0000                     unimp 
  3114: 0000                     unimp 
  3116: 0000                     unimp 
  3118: 0000                     unimp 
