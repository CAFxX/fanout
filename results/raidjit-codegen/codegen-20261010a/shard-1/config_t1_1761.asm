; config_t1_1761 target=rvv-zvbc32e twin=maskedtail topo=raid6(8)/enc op=encode U=8
; code_size=6336 twins=["maskedtail"] uops_per_byte=10.1171875
     0: 63140600                 bnez a2, 8
     4: 6f10d00a                 j 0x18b0
     8: 972f0000                 auipc t6, 2
    12: 938f8f8b                 addi t6, t6, -0x748
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d74f0e5e                 vmv.v.x v31, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93068000                 li a3, 8
    40: b386d202                 mul a3, t0, a3
    44: 3357d602                 divu a4, a2, a3
    48: 63140700                 bnez a4, 0x38
    52: 6f005019                 j 0x9c8
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 87880302                 vle8.v v17, (t2)
    76: b3835300                 add t2, t2, t0
    80: 07890302                 vle8.v v18, (t2)
    84: b3835300                 add t2, t2, t0
    88: 87890302                 vle8.v v19, (t2)
    92: b3835300                 add t2, t2, t0
    96: 078a0302                 vle8.v v20, (t2)
   100: b3835300                 add t2, t2, t0
   104: 878a0302                 vle8.v v21, (t2)
   108: b3835300                 add t2, t2, t0
   112: 078b0302                 vle8.v v22, (t2)
   116: b3835300                 add t2, t2, t0
   120: 878b0302                 vle8.v v23, (t2)
   124: b3835300                 add t2, t2, t0
   128: 078c0302                 vle8.v v24, (t2)
   132: d780085e                 vmv.v.v v1, v17
   136: 5701095e                 vmv.v.v v2, v18
   140: d781095e                 vmv.v.v v3, v19
   144: 57020a5e                 vmv.v.v v4, v20
   148: d7820a5e                 vmv.v.v v5, v21
   152: 57030b5e                 vmv.v.v v6, v22
   156: d7830b5e                 vmv.v.v v7, v23
   160: 57040c5e                 vmv.v.v v8, v24
   164: 138e0f00                 mv t3, t6
   168: 870e0e02                 vle8.v v29, (t3)
   172: 138e0f01                 addi t3, t6, 0x10
   176: 070f0e02                 vle8.v v30, (t3)
   180: d78cf827                 vand.vv v25, v31, v17
   184: 573d12a3                 vsrl.vi v26, v17, 4
   188: d78ddc33                 vrgather.vv v27, v29, v25
   192: 570eed33                 vrgather.vv v28, v30, v26
   196: d78dcd2f                 vxor.vv v27, v28, v27
   200: d7840d5e                 vmv.v.v v9, v27
   204: d70cf927                 vand.vv v25, v31, v18
   208: 573d22a3                 vsrl.vi v26, v18, 4
   212: d78ddc33                 vrgather.vv v27, v29, v25
   216: 570eed33                 vrgather.vv v28, v30, v26
   220: d78dcd2f                 vxor.vv v27, v28, v27
   224: 57850d5e                 vmv.v.v v10, v27
   228: d78cf927                 vand.vv v25, v31, v19
   232: 573d32a3                 vsrl.vi v26, v19, 4
   236: d78ddc33                 vrgather.vv v27, v29, v25
   240: 570eed33                 vrgather.vv v28, v30, v26
   244: d78dcd2f                 vxor.vv v27, v28, v27
   248: d7850d5e                 vmv.v.v v11, v27
   252: d70cfa27                 vand.vv v25, v31, v20
   256: 573d42a3                 vsrl.vi v26, v20, 4
   260: d78ddc33                 vrgather.vv v27, v29, v25
   264: 570eed33                 vrgather.vv v28, v30, v26
   268: d78dcd2f                 vxor.vv v27, v28, v27
   272: 57860d5e                 vmv.v.v v12, v27
   276: d78cfa27                 vand.vv v25, v31, v21
   280: 573d52a3                 vsrl.vi v26, v21, 4
   284: d78ddc33                 vrgather.vv v27, v29, v25
   288: 570eed33                 vrgather.vv v28, v30, v26
   292: d78dcd2f                 vxor.vv v27, v28, v27
   296: d7860d5e                 vmv.v.v v13, v27
   300: d70cfb27                 vand.vv v25, v31, v22
   304: 573d62a3                 vsrl.vi v26, v22, 4
   308: d78ddc33                 vrgather.vv v27, v29, v25
   312: 570eed33                 vrgather.vv v28, v30, v26
   316: d78dcd2f                 vxor.vv v27, v28, v27
   320: 57870d5e                 vmv.v.v v14, v27
   324: d78cfb27                 vand.vv v25, v31, v23
   328: 573d72a3                 vsrl.vi v26, v23, 4
   332: d78ddc33                 vrgather.vv v27, v29, v25
   336: 570eed33                 vrgather.vv v28, v30, v26
   340: d78dcd2f                 vxor.vv v27, v28, v27
   344: d7870d5e                 vmv.v.v v15, v27
   348: d70cfc27                 vand.vv v25, v31, v24
   352: 573d82a3                 vsrl.vi v26, v24, 4
   356: d78ddc33                 vrgather.vv v27, v29, v25
   360: 570eed33                 vrgather.vv v28, v30, v26
   364: d78dcd2f                 vxor.vv v27, v28, v27
   368: 57880d5e                 vmv.v.v v16, v27
   372: 83338500                 ld t2, 8(a0)
   376: b3836300                 add t2, t2, t1
   380: 87880302                 vle8.v v17, (t2)
   384: b3835300                 add t2, t2, t0
   388: 07890302                 vle8.v v18, (t2)
   392: b3835300                 add t2, t2, t0
   396: 87890302                 vle8.v v19, (t2)
   400: b3835300                 add t2, t2, t0
   404: 078a0302                 vle8.v v20, (t2)
   408: b3835300                 add t2, t2, t0
   412: 878a0302                 vle8.v v21, (t2)
   416: b3835300                 add t2, t2, t0
   420: 078b0302                 vle8.v v22, (t2)
   424: b3835300                 add t2, t2, t0
   428: 878b0302                 vle8.v v23, (t2)
   432: b3835300                 add t2, t2, t0
   436: 078c0302                 vle8.v v24, (t2)
   440: d780102f                 vxor.vv v1, v17, v1
   444: 5701212f                 vxor.vv v2, v18, v2
   448: d781312f                 vxor.vv v3, v19, v3
   452: 5702422f                 vxor.vv v4, v20, v4
   456: d782522f                 vxor.vv v5, v21, v5
   460: 5703632f                 vxor.vv v6, v22, v6
   464: d783732f                 vxor.vv v7, v23, v7
   468: 5704842f                 vxor.vv v8, v24, v8
   472: 138e0f02                 addi t3, t6, 0x20
   476: 870e0e02                 vle8.v v29, (t3)
   480: 138e0f03                 addi t3, t6, 0x30
   484: 070f0e02                 vle8.v v30, (t3)
   488: d78cf827                 vand.vv v25, v31, v17
   492: 573d12a3                 vsrl.vi v26, v17, 4
   496: d78ddc33                 vrgather.vv v27, v29, v25
   500: 570eed33                 vrgather.vv v28, v30, v26
   504: d78dcd2f                 vxor.vv v27, v28, v27
   508: d784b42f                 vxor.vv v9, v27, v9
   512: d70cf927                 vand.vv v25, v31, v18
   516: 573d22a3                 vsrl.vi v26, v18, 4
   520: d78ddc33                 vrgather.vv v27, v29, v25
   524: 570eed33                 vrgather.vv v28, v30, v26
   528: d78dcd2f                 vxor.vv v27, v28, v27
   532: 5705b52f                 vxor.vv v10, v27, v10
   536: d78cf927                 vand.vv v25, v31, v19
   540: 573d32a3                 vsrl.vi v26, v19, 4
   544: d78ddc33                 vrgather.vv v27, v29, v25
   548: 570eed33                 vrgather.vv v28, v30, v26
   552: d78dcd2f                 vxor.vv v27, v28, v27
   556: d785b52f                 vxor.vv v11, v27, v11
   560: d70cfa27                 vand.vv v25, v31, v20
   564: 573d42a3                 vsrl.vi v26, v20, 4
   568: d78ddc33                 vrgather.vv v27, v29, v25
   572: 570eed33                 vrgather.vv v28, v30, v26
   576: d78dcd2f                 vxor.vv v27, v28, v27
   580: 5706b62f                 vxor.vv v12, v27, v12
   584: d78cfa27                 vand.vv v25, v31, v21
   588: 573d52a3                 vsrl.vi v26, v21, 4
   592: d78ddc33                 vrgather.vv v27, v29, v25
   596: 570eed33                 vrgather.vv v28, v30, v26
   600: d78dcd2f                 vxor.vv v27, v28, v27
   604: d786b62f                 vxor.vv v13, v27, v13
   608: d70cfb27                 vand.vv v25, v31, v22
   612: 573d62a3                 vsrl.vi v26, v22, 4
   616: d78ddc33                 vrgather.vv v27, v29, v25
   620: 570eed33                 vrgather.vv v28, v30, v26
   624: d78dcd2f                 vxor.vv v27, v28, v27
   628: 5707b72f                 vxor.vv v14, v27, v14
   632: d78cfb27                 vand.vv v25, v31, v23
   636: 573d72a3                 vsrl.vi v26, v23, 4
   640: d78ddc33                 vrgather.vv v27, v29, v25
   644: 570eed33                 vrgather.vv v28, v30, v26
   648: d78dcd2f                 vxor.vv v27, v28, v27
   652: d787b72f                 vxor.vv v15, v27, v15
   656: d70cfc27                 vand.vv v25, v31, v24
   660: 573d82a3                 vsrl.vi v26, v24, 4
   664: d78ddc33                 vrgather.vv v27, v29, v25
   668: 570eed33                 vrgather.vv v28, v30, v26
   672: d78dcd2f                 vxor.vv v27, v28, v27
   676: 5708b82f                 vxor.vv v16, v27, v16
   680: 83330501                 ld t2, 0x10(a0)
   684: b3836300                 add t2, t2, t1
   688: 87880302                 vle8.v v17, (t2)
   692: b3835300                 add t2, t2, t0
   696: 07890302                 vle8.v v18, (t2)
   700: b3835300                 add t2, t2, t0
   704: 87890302                 vle8.v v19, (t2)
   708: b3835300                 add t2, t2, t0
   712: 078a0302                 vle8.v v20, (t2)
   716: b3835300                 add t2, t2, t0
   720: 878a0302                 vle8.v v21, (t2)
   724: b3835300                 add t2, t2, t0
   728: 078b0302                 vle8.v v22, (t2)
   732: b3835300                 add t2, t2, t0
   736: 878b0302                 vle8.v v23, (t2)
   740: b3835300                 add t2, t2, t0
   744: 078c0302                 vle8.v v24, (t2)
   748: d780102f                 vxor.vv v1, v17, v1
   752: 5701212f                 vxor.vv v2, v18, v2
   756: d781312f                 vxor.vv v3, v19, v3
   760: 5702422f                 vxor.vv v4, v20, v4
   764: d782522f                 vxor.vv v5, v21, v5
   768: 5703632f                 vxor.vv v6, v22, v6
   772: d783732f                 vxor.vv v7, v23, v7
   776: 5704842f                 vxor.vv v8, v24, v8
   780: 138e0f04                 addi t3, t6, 0x40
   784: 870e0e02                 vle8.v v29, (t3)
   788: 138e0f05                 addi t3, t6, 0x50
   792: 070f0e02                 vle8.v v30, (t3)
   796: d78cf827                 vand.vv v25, v31, v17
   800: 573d12a3                 vsrl.vi v26, v17, 4
   804: d78ddc33                 vrgather.vv v27, v29, v25
   808: 570eed33                 vrgather.vv v28, v30, v26
   812: d78dcd2f                 vxor.vv v27, v28, v27
   816: d784b42f                 vxor.vv v9, v27, v9
   820: d70cf927                 vand.vv v25, v31, v18
   824: 573d22a3                 vsrl.vi v26, v18, 4
   828: d78ddc33                 vrgather.vv v27, v29, v25
   832: 570eed33                 vrgather.vv v28, v30, v26
   836: d78dcd2f                 vxor.vv v27, v28, v27
   840: 5705b52f                 vxor.vv v10, v27, v10
   844: d78cf927                 vand.vv v25, v31, v19
   848: 573d32a3                 vsrl.vi v26, v19, 4
   852: d78ddc33                 vrgather.vv v27, v29, v25
   856: 570eed33                 vrgather.vv v28, v30, v26
   860: d78dcd2f                 vxor.vv v27, v28, v27
   864: d785b52f                 vxor.vv v11, v27, v11
   868: d70cfa27                 vand.vv v25, v31, v20
   872: 573d42a3                 vsrl.vi v26, v20, 4
   876: d78ddc33                 vrgather.vv v27, v29, v25
   880: 570eed33                 vrgather.vv v28, v30, v26
   884: d78dcd2f                 vxor.vv v27, v28, v27
   888: 5706b62f                 vxor.vv v12, v27, v12
   892: d78cfa27                 vand.vv v25, v31, v21
   896: 573d52a3                 vsrl.vi v26, v21, 4
   900: d78ddc33                 vrgather.vv v27, v29, v25
   904: 570eed33                 vrgather.vv v28, v30, v26
   908: d78dcd2f                 vxor.vv v27, v28, v27
   912: d786b62f                 vxor.vv v13, v27, v13
   916: d70cfb27                 vand.vv v25, v31, v22
   920: 573d62a3                 vsrl.vi v26, v22, 4
   924: d78ddc33                 vrgather.vv v27, v29, v25
   928: 570eed33                 vrgather.vv v28, v30, v26
   932: d78dcd2f                 vxor.vv v27, v28, v27
   936: 5707b72f                 vxor.vv v14, v27, v14
   940: d78cfb27                 vand.vv v25, v31, v23
   944: 573d72a3                 vsrl.vi v26, v23, 4
   948: d78ddc33                 vrgather.vv v27, v29, v25
   952: 570eed33                 vrgather.vv v28, v30, v26
   956: d78dcd2f                 vxor.vv v27, v28, v27
   960: d787b72f                 vxor.vv v15, v27, v15
   964: d70cfc27                 vand.vv v25, v31, v24
   968: 573d82a3                 vsrl.vi v26, v24, 4
   972: d78ddc33                 vrgather.vv v27, v29, v25
   976: 570eed33                 vrgather.vv v28, v30, v26
   980: d78dcd2f                 vxor.vv v27, v28, v27
   984: 5708b82f                 vxor.vv v16, v27, v16
   988: 83338501                 ld t2, 0x18(a0)
   992: b3836300                 add t2, t2, t1
   996: 87880302                 vle8.v v17, (t2)
  1000: b3835300                 add t2, t2, t0
  1004: 07890302                 vle8.v v18, (t2)
  1008: b3835300                 add t2, t2, t0
  1012: 87890302                 vle8.v v19, (t2)
  1016: b3835300                 add t2, t2, t0
  1020: 078a0302                 vle8.v v20, (t2)
  1024: b3835300                 add t2, t2, t0
  1028: 878a0302                 vle8.v v21, (t2)
  1032: b3835300                 add t2, t2, t0
  1036: 078b0302                 vle8.v v22, (t2)
  1040: b3835300                 add t2, t2, t0
  1044: 878b0302                 vle8.v v23, (t2)
  1048: b3835300                 add t2, t2, t0
  1052: 078c0302                 vle8.v v24, (t2)
  1056: d780102f                 vxor.vv v1, v17, v1
  1060: 5701212f                 vxor.vv v2, v18, v2
  1064: d781312f                 vxor.vv v3, v19, v3
  1068: 5702422f                 vxor.vv v4, v20, v4
  1072: d782522f                 vxor.vv v5, v21, v5
  1076: 5703632f                 vxor.vv v6, v22, v6
  1080: d783732f                 vxor.vv v7, v23, v7
  1084: 5704842f                 vxor.vv v8, v24, v8
  1088: 138e0f06                 addi t3, t6, 0x60
  1092: 870e0e02                 vle8.v v29, (t3)
  1096: 138e0f07                 addi t3, t6, 0x70
  1100: 070f0e02                 vle8.v v30, (t3)
  1104: d78cf827                 vand.vv v25, v31, v17
  1108: 573d12a3                 vsrl.vi v26, v17, 4
  1112: d78ddc33                 vrgather.vv v27, v29, v25
  1116: 570eed33                 vrgather.vv v28, v30, v26
  1120: d78dcd2f                 vxor.vv v27, v28, v27
  1124: d784b42f                 vxor.vv v9, v27, v9
  1128: d70cf927                 vand.vv v25, v31, v18
  1132: 573d22a3                 vsrl.vi v26, v18, 4
  1136: d78ddc33                 vrgather.vv v27, v29, v25
  1140: 570eed33                 vrgather.vv v28, v30, v26
  1144: d78dcd2f                 vxor.vv v27, v28, v27
  1148: 5705b52f                 vxor.vv v10, v27, v10
  1152: d78cf927                 vand.vv v25, v31, v19
  1156: 573d32a3                 vsrl.vi v26, v19, 4
  1160: d78ddc33                 vrgather.vv v27, v29, v25
  1164: 570eed33                 vrgather.vv v28, v30, v26
  1168: d78dcd2f                 vxor.vv v27, v28, v27
  1172: d785b52f                 vxor.vv v11, v27, v11
  1176: d70cfa27                 vand.vv v25, v31, v20
  1180: 573d42a3                 vsrl.vi v26, v20, 4
  1184: d78ddc33                 vrgather.vv v27, v29, v25
  1188: 570eed33                 vrgather.vv v28, v30, v26
  1192: d78dcd2f                 vxor.vv v27, v28, v27
  1196: 5706b62f                 vxor.vv v12, v27, v12
  1200: d78cfa27                 vand.vv v25, v31, v21
  1204: 573d52a3                 vsrl.vi v26, v21, 4
  1208: d78ddc33                 vrgather.vv v27, v29, v25
  1212: 570eed33                 vrgather.vv v28, v30, v26
  1216: d78dcd2f                 vxor.vv v27, v28, v27
  1220: d786b62f                 vxor.vv v13, v27, v13
  1224: d70cfb27                 vand.vv v25, v31, v22
  1228: 573d62a3                 vsrl.vi v26, v22, 4
  1232: d78ddc33                 vrgather.vv v27, v29, v25
  1236: 570eed33                 vrgather.vv v28, v30, v26
  1240: d78dcd2f                 vxor.vv v27, v28, v27
  1244: 5707b72f                 vxor.vv v14, v27, v14
  1248: d78cfb27                 vand.vv v25, v31, v23
  1252: 573d72a3                 vsrl.vi v26, v23, 4
  1256: d78ddc33                 vrgather.vv v27, v29, v25
  1260: 570eed33                 vrgather.vv v28, v30, v26
  1264: d78dcd2f                 vxor.vv v27, v28, v27
  1268: d787b72f                 vxor.vv v15, v27, v15
  1272: d70cfc27                 vand.vv v25, v31, v24
  1276: 573d82a3                 vsrl.vi v26, v24, 4
  1280: d78ddc33                 vrgather.vv v27, v29, v25
  1284: 570eed33                 vrgather.vv v28, v30, v26
  1288: d78dcd2f                 vxor.vv v27, v28, v27
  1292: 5708b82f                 vxor.vv v16, v27, v16
  1296: 83330502                 ld t2, 0x20(a0)
  1300: b3836300                 add t2, t2, t1
  1304: 87880302                 vle8.v v17, (t2)
  1308: b3835300                 add t2, t2, t0
  1312: 07890302                 vle8.v v18, (t2)
  1316: b3835300                 add t2, t2, t0
  1320: 87890302                 vle8.v v19, (t2)
  1324: b3835300                 add t2, t2, t0
  1328: 078a0302                 vle8.v v20, (t2)
  1332: b3835300                 add t2, t2, t0
  1336: 878a0302                 vle8.v v21, (t2)
  1340: b3835300                 add t2, t2, t0
  1344: 078b0302                 vle8.v v22, (t2)
  1348: b3835300                 add t2, t2, t0
  1352: 878b0302                 vle8.v v23, (t2)
  1356: b3835300                 add t2, t2, t0
  1360: 078c0302                 vle8.v v24, (t2)
  1364: d780102f                 vxor.vv v1, v17, v1
  1368: 5701212f                 vxor.vv v2, v18, v2
  1372: d781312f                 vxor.vv v3, v19, v3
  1376: 5702422f                 vxor.vv v4, v20, v4
  1380: d782522f                 vxor.vv v5, v21, v5
  1384: 5703632f                 vxor.vv v6, v22, v6
  1388: d783732f                 vxor.vv v7, v23, v7
  1392: 5704842f                 vxor.vv v8, v24, v8
  1396: 138e0f08                 addi t3, t6, 0x80
  1400: 870e0e02                 vle8.v v29, (t3)
  1404: 138e0f09                 addi t3, t6, 0x90
  1408: 070f0e02                 vle8.v v30, (t3)
  1412: d78cf827                 vand.vv v25, v31, v17
  1416: 573d12a3                 vsrl.vi v26, v17, 4
  1420: d78ddc33                 vrgather.vv v27, v29, v25
  1424: 570eed33                 vrgather.vv v28, v30, v26
  1428: d78dcd2f                 vxor.vv v27, v28, v27
  1432: d784b42f                 vxor.vv v9, v27, v9
  1436: d70cf927                 vand.vv v25, v31, v18
  1440: 573d22a3                 vsrl.vi v26, v18, 4
  1444: d78ddc33                 vrgather.vv v27, v29, v25
  1448: 570eed33                 vrgather.vv v28, v30, v26
  1452: d78dcd2f                 vxor.vv v27, v28, v27
  1456: 5705b52f                 vxor.vv v10, v27, v10
  1460: d78cf927                 vand.vv v25, v31, v19
  1464: 573d32a3                 vsrl.vi v26, v19, 4
  1468: d78ddc33                 vrgather.vv v27, v29, v25
  1472: 570eed33                 vrgather.vv v28, v30, v26
  1476: d78dcd2f                 vxor.vv v27, v28, v27
  1480: d785b52f                 vxor.vv v11, v27, v11
  1484: d70cfa27                 vand.vv v25, v31, v20
  1488: 573d42a3                 vsrl.vi v26, v20, 4
  1492: d78ddc33                 vrgather.vv v27, v29, v25
  1496: 570eed33                 vrgather.vv v28, v30, v26
  1500: d78dcd2f                 vxor.vv v27, v28, v27
  1504: 5706b62f                 vxor.vv v12, v27, v12
  1508: d78cfa27                 vand.vv v25, v31, v21
  1512: 573d52a3                 vsrl.vi v26, v21, 4
  1516: d78ddc33                 vrgather.vv v27, v29, v25
  1520: 570eed33                 vrgather.vv v28, v30, v26
  1524: d78dcd2f                 vxor.vv v27, v28, v27
  1528: d786b62f                 vxor.vv v13, v27, v13
  1532: d70cfb27                 vand.vv v25, v31, v22
  1536: 573d62a3                 vsrl.vi v26, v22, 4
  1540: d78ddc33                 vrgather.vv v27, v29, v25
  1544: 570eed33                 vrgather.vv v28, v30, v26
  1548: d78dcd2f                 vxor.vv v27, v28, v27
  1552: 5707b72f                 vxor.vv v14, v27, v14
  1556: d78cfb27                 vand.vv v25, v31, v23
  1560: 573d72a3                 vsrl.vi v26, v23, 4
  1564: d78ddc33                 vrgather.vv v27, v29, v25
  1568: 570eed33                 vrgather.vv v28, v30, v26
  1572: d78dcd2f                 vxor.vv v27, v28, v27
  1576: d787b72f                 vxor.vv v15, v27, v15
  1580: d70cfc27                 vand.vv v25, v31, v24
  1584: 573d82a3                 vsrl.vi v26, v24, 4
  1588: d78ddc33                 vrgather.vv v27, v29, v25
  1592: 570eed33                 vrgather.vv v28, v30, v26
  1596: d78dcd2f                 vxor.vv v27, v28, v27
  1600: 5708b82f                 vxor.vv v16, v27, v16
  1604: 83338502                 ld t2, 0x28(a0)
  1608: b3836300                 add t2, t2, t1
  1612: 87880302                 vle8.v v17, (t2)
  1616: b3835300                 add t2, t2, t0
  1620: 07890302                 vle8.v v18, (t2)
  1624: b3835300                 add t2, t2, t0
  1628: 87890302                 vle8.v v19, (t2)
  1632: b3835300                 add t2, t2, t0
  1636: 078a0302                 vle8.v v20, (t2)
  1640: b3835300                 add t2, t2, t0
  1644: 878a0302                 vle8.v v21, (t2)
  1648: b3835300                 add t2, t2, t0
  1652: 078b0302                 vle8.v v22, (t2)
  1656: b3835300                 add t2, t2, t0
  1660: 878b0302                 vle8.v v23, (t2)
  1664: b3835300                 add t2, t2, t0
  1668: 078c0302                 vle8.v v24, (t2)
  1672: d780102f                 vxor.vv v1, v17, v1
  1676: 5701212f                 vxor.vv v2, v18, v2
  1680: d781312f                 vxor.vv v3, v19, v3
  1684: 5702422f                 vxor.vv v4, v20, v4
  1688: d782522f                 vxor.vv v5, v21, v5
  1692: 5703632f                 vxor.vv v6, v22, v6
  1696: d783732f                 vxor.vv v7, v23, v7
  1700: 5704842f                 vxor.vv v8, v24, v8
  1704: 138e0f0a                 addi t3, t6, 0xa0
  1708: 870e0e02                 vle8.v v29, (t3)
  1712: 138e0f0b                 addi t3, t6, 0xb0
  1716: 070f0e02                 vle8.v v30, (t3)
  1720: d78cf827                 vand.vv v25, v31, v17
  1724: 573d12a3                 vsrl.vi v26, v17, 4
  1728: d78ddc33                 vrgather.vv v27, v29, v25
  1732: 570eed33                 vrgather.vv v28, v30, v26
  1736: d78dcd2f                 vxor.vv v27, v28, v27
  1740: d784b42f                 vxor.vv v9, v27, v9
  1744: d70cf927                 vand.vv v25, v31, v18
  1748: 573d22a3                 vsrl.vi v26, v18, 4
  1752: d78ddc33                 vrgather.vv v27, v29, v25
  1756: 570eed33                 vrgather.vv v28, v30, v26
  1760: d78dcd2f                 vxor.vv v27, v28, v27
  1764: 5705b52f                 vxor.vv v10, v27, v10
  1768: d78cf927                 vand.vv v25, v31, v19
  1772: 573d32a3                 vsrl.vi v26, v19, 4
  1776: d78ddc33                 vrgather.vv v27, v29, v25
  1780: 570eed33                 vrgather.vv v28, v30, v26
  1784: d78dcd2f                 vxor.vv v27, v28, v27
  1788: d785b52f                 vxor.vv v11, v27, v11
  1792: d70cfa27                 vand.vv v25, v31, v20
  1796: 573d42a3                 vsrl.vi v26, v20, 4
  1800: d78ddc33                 vrgather.vv v27, v29, v25
  1804: 570eed33                 vrgather.vv v28, v30, v26
  1808: d78dcd2f                 vxor.vv v27, v28, v27
  1812: 5706b62f                 vxor.vv v12, v27, v12
  1816: d78cfa27                 vand.vv v25, v31, v21
  1820: 573d52a3                 vsrl.vi v26, v21, 4
  1824: d78ddc33                 vrgather.vv v27, v29, v25
  1828: 570eed33                 vrgather.vv v28, v30, v26
  1832: d78dcd2f                 vxor.vv v27, v28, v27
  1836: d786b62f                 vxor.vv v13, v27, v13
  1840: d70cfb27                 vand.vv v25, v31, v22
  1844: 573d62a3                 vsrl.vi v26, v22, 4
  1848: d78ddc33                 vrgather.vv v27, v29, v25
  1852: 570eed33                 vrgather.vv v28, v30, v26
  1856: d78dcd2f                 vxor.vv v27, v28, v27
  1860: 5707b72f                 vxor.vv v14, v27, v14
  1864: d78cfb27                 vand.vv v25, v31, v23
  1868: 573d72a3                 vsrl.vi v26, v23, 4
  1872: d78ddc33                 vrgather.vv v27, v29, v25
  1876: 570eed33                 vrgather.vv v28, v30, v26
  1880: d78dcd2f                 vxor.vv v27, v28, v27
  1884: d787b72f                 vxor.vv v15, v27, v15
  1888: d70cfc27                 vand.vv v25, v31, v24
  1892: 573d82a3                 vsrl.vi v26, v24, 4
  1896: d78ddc33                 vrgather.vv v27, v29, v25
  1900: 570eed33                 vrgather.vv v28, v30, v26
  1904: d78dcd2f                 vxor.vv v27, v28, v27
  1908: 5708b82f                 vxor.vv v16, v27, v16
  1912: 83330503                 ld t2, 0x30(a0)
  1916: b3836300                 add t2, t2, t1
  1920: 87880302                 vle8.v v17, (t2)
  1924: b3835300                 add t2, t2, t0
  1928: 07890302                 vle8.v v18, (t2)
  1932: b3835300                 add t2, t2, t0
  1936: 87890302                 vle8.v v19, (t2)
  1940: b3835300                 add t2, t2, t0
  1944: 078a0302                 vle8.v v20, (t2)
  1948: b3835300                 add t2, t2, t0
  1952: 878a0302                 vle8.v v21, (t2)
  1956: b3835300                 add t2, t2, t0
  1960: 078b0302                 vle8.v v22, (t2)
  1964: b3835300                 add t2, t2, t0
  1968: 878b0302                 vle8.v v23, (t2)
  1972: b3835300                 add t2, t2, t0
  1976: 078c0302                 vle8.v v24, (t2)
  1980: d780102f                 vxor.vv v1, v17, v1
  1984: 5701212f                 vxor.vv v2, v18, v2
  1988: d781312f                 vxor.vv v3, v19, v3
  1992: 5702422f                 vxor.vv v4, v20, v4
  1996: d782522f                 vxor.vv v5, v21, v5
  2000: 5703632f                 vxor.vv v6, v22, v6
  2004: d783732f                 vxor.vv v7, v23, v7
  2008: 5704842f                 vxor.vv v8, v24, v8
  2012: 138e0f0c                 addi t3, t6, 0xc0
  2016: 870e0e02                 vle8.v v29, (t3)
  2020: 138e0f0d                 addi t3, t6, 0xd0
  2024: 070f0e02                 vle8.v v30, (t3)
  2028: d78cf827                 vand.vv v25, v31, v17
  2032: 573d12a3                 vsrl.vi v26, v17, 4
  2036: d78ddc33                 vrgather.vv v27, v29, v25
  2040: 570eed33                 vrgather.vv v28, v30, v26
  2044: d78dcd2f                 vxor.vv v27, v28, v27
  2048: d784b42f                 vxor.vv v9, v27, v9
  2052: d70cf927                 vand.vv v25, v31, v18
  2056: 573d22a3                 vsrl.vi v26, v18, 4
  2060: d78ddc33                 vrgather.vv v27, v29, v25
  2064: 570eed33                 vrgather.vv v28, v30, v26
  2068: d78dcd2f                 vxor.vv v27, v28, v27
  2072: 5705b52f                 vxor.vv v10, v27, v10
  2076: d78cf927                 vand.vv v25, v31, v19
  2080: 573d32a3                 vsrl.vi v26, v19, 4
  2084: d78ddc33                 vrgather.vv v27, v29, v25
  2088: 570eed33                 vrgather.vv v28, v30, v26
  2092: d78dcd2f                 vxor.vv v27, v28, v27
  2096: d785b52f                 vxor.vv v11, v27, v11
  2100: d70cfa27                 vand.vv v25, v31, v20
  2104: 573d42a3                 vsrl.vi v26, v20, 4
  2108: d78ddc33                 vrgather.vv v27, v29, v25
  2112: 570eed33                 vrgather.vv v28, v30, v26
  2116: d78dcd2f                 vxor.vv v27, v28, v27
  2120: 5706b62f                 vxor.vv v12, v27, v12
  2124: d78cfa27                 vand.vv v25, v31, v21
  2128: 573d52a3                 vsrl.vi v26, v21, 4
  2132: d78ddc33                 vrgather.vv v27, v29, v25
  2136: 570eed33                 vrgather.vv v28, v30, v26
  2140: d78dcd2f                 vxor.vv v27, v28, v27
  2144: d786b62f                 vxor.vv v13, v27, v13
  2148: d70cfb27                 vand.vv v25, v31, v22
  2152: 573d62a3                 vsrl.vi v26, v22, 4
  2156: d78ddc33                 vrgather.vv v27, v29, v25
  2160: 570eed33                 vrgather.vv v28, v30, v26
  2164: d78dcd2f                 vxor.vv v27, v28, v27
  2168: 5707b72f                 vxor.vv v14, v27, v14
  2172: d78cfb27                 vand.vv v25, v31, v23
  2176: 573d72a3                 vsrl.vi v26, v23, 4
  2180: d78ddc33                 vrgather.vv v27, v29, v25
  2184: 570eed33                 vrgather.vv v28, v30, v26
  2188: d78dcd2f                 vxor.vv v27, v28, v27
  2192: d787b72f                 vxor.vv v15, v27, v15
  2196: d70cfc27                 vand.vv v25, v31, v24
  2200: 573d82a3                 vsrl.vi v26, v24, 4
  2204: d78ddc33                 vrgather.vv v27, v29, v25
  2208: 570eed33                 vrgather.vv v28, v30, v26
  2212: d78dcd2f                 vxor.vv v27, v28, v27
  2216: 5708b82f                 vxor.vv v16, v27, v16
  2220: 83338503                 ld t2, 0x38(a0)
  2224: b3836300                 add t2, t2, t1
  2228: 87880302                 vle8.v v17, (t2)
  2232: b3835300                 add t2, t2, t0
  2236: 07890302                 vle8.v v18, (t2)
  2240: b3835300                 add t2, t2, t0
  2244: 87890302                 vle8.v v19, (t2)
  2248: b3835300                 add t2, t2, t0
  2252: 078a0302                 vle8.v v20, (t2)
  2256: b3835300                 add t2, t2, t0
  2260: 878a0302                 vle8.v v21, (t2)
  2264: b3835300                 add t2, t2, t0
  2268: 078b0302                 vle8.v v22, (t2)
  2272: b3835300                 add t2, t2, t0
  2276: 878b0302                 vle8.v v23, (t2)
  2280: b3835300                 add t2, t2, t0
  2284: 078c0302                 vle8.v v24, (t2)
  2288: d780102f                 vxor.vv v1, v17, v1
  2292: 5701212f                 vxor.vv v2, v18, v2
  2296: d781312f                 vxor.vv v3, v19, v3
  2300: 5702422f                 vxor.vv v4, v20, v4
  2304: d782522f                 vxor.vv v5, v21, v5
  2308: 5703632f                 vxor.vv v6, v22, v6
  2312: d783732f                 vxor.vv v7, v23, v7
  2316: 5704842f                 vxor.vv v8, v24, v8
  2320: d784142f                 vxor.vv v9, v17, v9
  2324: 5705252f                 vxor.vv v10, v18, v10
  2328: d785352f                 vxor.vv v11, v19, v11
  2332: 5706462f                 vxor.vv v12, v20, v12
  2336: d786562f                 vxor.vv v13, v21, v13
  2340: 5707672f                 vxor.vv v14, v22, v14
  2344: d787772f                 vxor.vv v15, v23, v15
  2348: 5708882f                 vxor.vv v16, v24, v16
  2352: 83b30500                 ld t2, 0(a1)
  2356: b3836300                 add t2, t2, t1
  2360: a7800302                 vse8.v v1, (t2)
  2364: b3835300                 add t2, t2, t0
  2368: 27810302                 vse8.v v2, (t2)
  2372: b3835300                 add t2, t2, t0
  2376: a7810302                 vse8.v v3, (t2)
  2380: b3835300                 add t2, t2, t0
  2384: 27820302                 vse8.v v4, (t2)
  2388: b3835300                 add t2, t2, t0
  2392: a7820302                 vse8.v v5, (t2)
  2396: b3835300                 add t2, t2, t0
  2400: 27830302                 vse8.v v6, (t2)
  2404: b3835300                 add t2, t2, t0
  2408: a7830302                 vse8.v v7, (t2)
  2412: b3835300                 add t2, t2, t0
  2416: 27840302                 vse8.v v8, (t2)
  2420: 83b38500                 ld t2, 8(a1)
  2424: b3836300                 add t2, t2, t1
  2428: a7840302                 vse8.v v9, (t2)
  2432: b3835300                 add t2, t2, t0
  2436: 27850302                 vse8.v v10, (t2)
  2440: b3835300                 add t2, t2, t0
  2444: a7850302                 vse8.v v11, (t2)
  2448: b3835300                 add t2, t2, t0
  2452: 27860302                 vse8.v v12, (t2)
  2456: b3835300                 add t2, t2, t0
  2460: a7860302                 vse8.v v13, (t2)
  2464: b3835300                 add t2, t2, t0
  2468: 27870302                 vse8.v v14, (t2)
  2472: b3835300                 add t2, t2, t0
  2476: a7870302                 vse8.v v15, (t2)
  2480: b3835300                 add t2, t2, t0
  2484: 27880302                 vse8.v v16, (t2)
  2488: 3303d300                 add t1, t1, a3
  2492: 1307f7ff                 addi a4, a4, -1
  2496: 631007e8                 bnez a4, 0x40
  2500: 33066640                 sub a2, a2, t1
  2504: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
  2508: 93060600                 mv a3, a2
  2512: 6344d000                 bgtz a3, 0x9d8
  2516: 6f00001d                 j 0xba4
  2520: b3835640                 sub t2, a3, t0
  2524: 13def343                 srai t3, t2, 0x3f
  2528: b3f3c301                 and t2, t2, t3
  2532: b3837200                 add t2, t0, t2
  2536: d7a80852                 vid.v v17
  2540: 57c0136b                 vmsltu.vx v0, v17, t2
  2544: 83330500                 ld t2, 0(a0)
  2548: b3836300                 add t2, t2, t1
  2552: 87880300                 vle8.v v17, (t2), v0.t
  2556: d780085e                 vmv.v.v v1, v17
  2560: 138e0f00                 mv t3, t6
  2564: 870e0e02                 vle8.v v29, (t3)
  2568: 138e0f01                 addi t3, t6, 0x10
  2572: 070f0e02                 vle8.v v30, (t3)
  2576: d78cf827                 vand.vv v25, v31, v17
  2580: 573d12a3                 vsrl.vi v26, v17, 4
  2584: d78ddc33                 vrgather.vv v27, v29, v25
  2588: 570eed33                 vrgather.vv v28, v30, v26
  2592: d78dcd2f                 vxor.vv v27, v28, v27
  2596: d7840d5e                 vmv.v.v v9, v27
  2600: 83338500                 ld t2, 8(a0)
  2604: b3836300                 add t2, t2, t1
  2608: 87880300                 vle8.v v17, (t2), v0.t
  2612: d780102f                 vxor.vv v1, v17, v1
  2616: 138e0f02                 addi t3, t6, 0x20
  2620: 870e0e02                 vle8.v v29, (t3)
  2624: 138e0f03                 addi t3, t6, 0x30
  2628: 070f0e02                 vle8.v v30, (t3)
  2632: d78cf827                 vand.vv v25, v31, v17
  2636: 573d12a3                 vsrl.vi v26, v17, 4
  2640: d78ddc33                 vrgather.vv v27, v29, v25
  2644: 570eed33                 vrgather.vv v28, v30, v26
  2648: d78dcd2f                 vxor.vv v27, v28, v27
  2652: d784b42f                 vxor.vv v9, v27, v9
  2656: 83330501                 ld t2, 0x10(a0)
  2660: b3836300                 add t2, t2, t1
  2664: 87880300                 vle8.v v17, (t2), v0.t
  2668: d780102f                 vxor.vv v1, v17, v1
  2672: 138e0f04                 addi t3, t6, 0x40
  2676: 870e0e02                 vle8.v v29, (t3)
  2680: 138e0f05                 addi t3, t6, 0x50
  2684: 070f0e02                 vle8.v v30, (t3)
  2688: d78cf827                 vand.vv v25, v31, v17
  2692: 573d12a3                 vsrl.vi v26, v17, 4
  2696: d78ddc33                 vrgather.vv v27, v29, v25
  2700: 570eed33                 vrgather.vv v28, v30, v26
  2704: d78dcd2f                 vxor.vv v27, v28, v27
  2708: d784b42f                 vxor.vv v9, v27, v9
  2712: 83338501                 ld t2, 0x18(a0)
  2716: b3836300                 add t2, t2, t1
  2720: 87880300                 vle8.v v17, (t2), v0.t
  2724: d780102f                 vxor.vv v1, v17, v1
  2728: 138e0f06                 addi t3, t6, 0x60
  2732: 870e0e02                 vle8.v v29, (t3)
  2736: 138e0f07                 addi t3, t6, 0x70
  2740: 070f0e02                 vle8.v v30, (t3)
  2744: d78cf827                 vand.vv v25, v31, v17
  2748: 573d12a3                 vsrl.vi v26, v17, 4
  2752: d78ddc33                 vrgather.vv v27, v29, v25
  2756: 570eed33                 vrgather.vv v28, v30, v26
  2760: d78dcd2f                 vxor.vv v27, v28, v27
  2764: d784b42f                 vxor.vv v9, v27, v9
  2768: 83330502                 ld t2, 0x20(a0)
  2772: b3836300                 add t2, t2, t1
  2776: 87880300                 vle8.v v17, (t2), v0.t
  2780: d780102f                 vxor.vv v1, v17, v1
  2784: 138e0f08                 addi t3, t6, 0x80
  2788: 870e0e02                 vle8.v v29, (t3)
  2792: 138e0f09                 addi t3, t6, 0x90
  2796: 070f0e02                 vle8.v v30, (t3)
  2800: d78cf827                 vand.vv v25, v31, v17
  2804: 573d12a3                 vsrl.vi v26, v17, 4
  2808: d78ddc33                 vrgather.vv v27, v29, v25
  2812: 570eed33                 vrgather.vv v28, v30, v26
  2816: d78dcd2f                 vxor.vv v27, v28, v27
  2820: d784b42f                 vxor.vv v9, v27, v9
  2824: 83338502                 ld t2, 0x28(a0)
  2828: b3836300                 add t2, t2, t1
  2832: 87880300                 vle8.v v17, (t2), v0.t
  2836: d780102f                 vxor.vv v1, v17, v1
  2840: 138e0f0a                 addi t3, t6, 0xa0
  2844: 870e0e02                 vle8.v v29, (t3)
  2848: 138e0f0b                 addi t3, t6, 0xb0
  2852: 070f0e02                 vle8.v v30, (t3)
  2856: d78cf827                 vand.vv v25, v31, v17
  2860: 573d12a3                 vsrl.vi v26, v17, 4
  2864: d78ddc33                 vrgather.vv v27, v29, v25
  2868: 570eed33                 vrgather.vv v28, v30, v26
  2872: d78dcd2f                 vxor.vv v27, v28, v27
  2876: d784b42f                 vxor.vv v9, v27, v9
  2880: 83330503                 ld t2, 0x30(a0)
  2884: b3836300                 add t2, t2, t1
  2888: 87880300                 vle8.v v17, (t2), v0.t
  2892: d780102f                 vxor.vv v1, v17, v1
  2896: 138e0f0c                 addi t3, t6, 0xc0
  2900: 870e0e02                 vle8.v v29, (t3)
  2904: 138e0f0d                 addi t3, t6, 0xd0
  2908: 070f0e02                 vle8.v v30, (t3)
  2912: d78cf827                 vand.vv v25, v31, v17
  2916: 573d12a3                 vsrl.vi v26, v17, 4
  2920: d78ddc33                 vrgather.vv v27, v29, v25
  2924: 570eed33                 vrgather.vv v28, v30, v26
  2928: d78dcd2f                 vxor.vv v27, v28, v27
  2932: d784b42f                 vxor.vv v9, v27, v9
  2936: 83338503                 ld t2, 0x38(a0)
  2940: b3836300                 add t2, t2, t1
  2944: 87880300                 vle8.v v17, (t2), v0.t
  2948: d780102f                 vxor.vv v1, v17, v1
  2952: d784142f                 vxor.vv v9, v17, v9
  2956: 83b30500                 ld t2, 0(a1)
  2960: b3836300                 add t2, t2, t1
  2964: a7800300                 vse8.v v1, (t2), v0.t
  2968: 83b38500                 ld t2, 8(a1)
  2972: b3836300                 add t2, t2, t1
  2976: a7840300                 vse8.v v9, (t2), v0.t
  2980: 33035300                 add t1, t1, t0
  2984: b3865640                 sub a3, a3, t0
  2988: 6344d000                 bgtz a3, 0xbb4
  2992: 6f00001d                 j 0xd80
  2996: b3835640                 sub t2, a3, t0
  3000: 13def343                 srai t3, t2, 0x3f
  3004: b3f3c301                 and t2, t2, t3
  3008: b3837200                 add t2, t0, t2
  3012: d7a80852                 vid.v v17
  3016: 57c0136b                 vmsltu.vx v0, v17, t2
  3020: 83330500                 ld t2, 0(a0)
  3024: b3836300                 add t2, t2, t1
  3028: 87880300                 vle8.v v17, (t2), v0.t
  3032: d780085e                 vmv.v.v v1, v17
  3036: 138e0f00                 mv t3, t6
  3040: 870e0e02                 vle8.v v29, (t3)
  3044: 138e0f01                 addi t3, t6, 0x10
  3048: 070f0e02                 vle8.v v30, (t3)
  3052: d78cf827                 vand.vv v25, v31, v17
  3056: 573d12a3                 vsrl.vi v26, v17, 4
  3060: d78ddc33                 vrgather.vv v27, v29, v25
  3064: 570eed33                 vrgather.vv v28, v30, v26
  3068: d78dcd2f                 vxor.vv v27, v28, v27
  3072: d7840d5e                 vmv.v.v v9, v27
  3076: 83338500                 ld t2, 8(a0)
  3080: b3836300                 add t2, t2, t1
  3084: 87880300                 vle8.v v17, (t2), v0.t
  3088: d780102f                 vxor.vv v1, v17, v1
  3092: 138e0f02                 addi t3, t6, 0x20
  3096: 870e0e02                 vle8.v v29, (t3)
  3100: 138e0f03                 addi t3, t6, 0x30
  3104: 070f0e02                 vle8.v v30, (t3)
  3108: d78cf827                 vand.vv v25, v31, v17
  3112: 573d12a3                 vsrl.vi v26, v17, 4
  3116: d78ddc33                 vrgather.vv v27, v29, v25
  3120: 570eed33                 vrgather.vv v28, v30, v26
  3124: d78dcd2f                 vxor.vv v27, v28, v27
  3128: d784b42f                 vxor.vv v9, v27, v9
  3132: 83330501                 ld t2, 0x10(a0)
  3136: b3836300                 add t2, t2, t1
  3140: 87880300                 vle8.v v17, (t2), v0.t
  3144: d780102f                 vxor.vv v1, v17, v1
  3148: 138e0f04                 addi t3, t6, 0x40
  3152: 870e0e02                 vle8.v v29, (t3)
  3156: 138e0f05                 addi t3, t6, 0x50
  3160: 070f0e02                 vle8.v v30, (t3)
  3164: d78cf827                 vand.vv v25, v31, v17
  3168: 573d12a3                 vsrl.vi v26, v17, 4
  3172: d78ddc33                 vrgather.vv v27, v29, v25
  3176: 570eed33                 vrgather.vv v28, v30, v26
  3180: d78dcd2f                 vxor.vv v27, v28, v27
  3184: d784b42f                 vxor.vv v9, v27, v9
  3188: 83338501                 ld t2, 0x18(a0)
  3192: b3836300                 add t2, t2, t1
  3196: 87880300                 vle8.v v17, (t2), v0.t
  3200: d780102f                 vxor.vv v1, v17, v1
  3204: 138e0f06                 addi t3, t6, 0x60
  3208: 870e0e02                 vle8.v v29, (t3)
  3212: 138e0f07                 addi t3, t6, 0x70
  3216: 070f0e02                 vle8.v v30, (t3)
  3220: d78cf827                 vand.vv v25, v31, v17
  3224: 573d12a3                 vsrl.vi v26, v17, 4
  3228: d78ddc33                 vrgather.vv v27, v29, v25
  3232: 570eed33                 vrgather.vv v28, v30, v26
  3236: d78dcd2f                 vxor.vv v27, v28, v27
  3240: d784b42f                 vxor.vv v9, v27, v9
  3244: 83330502                 ld t2, 0x20(a0)
  3248: b3836300                 add t2, t2, t1
  3252: 87880300                 vle8.v v17, (t2), v0.t
  3256: d780102f                 vxor.vv v1, v17, v1
  3260: 138e0f08                 addi t3, t6, 0x80
  3264: 870e0e02                 vle8.v v29, (t3)
  3268: 138e0f09                 addi t3, t6, 0x90
  3272: 070f0e02                 vle8.v v30, (t3)
  3276: d78cf827                 vand.vv v25, v31, v17
  3280: 573d12a3                 vsrl.vi v26, v17, 4
  3284: d78ddc33                 vrgather.vv v27, v29, v25
  3288: 570eed33                 vrgather.vv v28, v30, v26
  3292: d78dcd2f                 vxor.vv v27, v28, v27
  3296: d784b42f                 vxor.vv v9, v27, v9
  3300: 83338502                 ld t2, 0x28(a0)
  3304: b3836300                 add t2, t2, t1
  3308: 87880300                 vle8.v v17, (t2), v0.t
  3312: d780102f                 vxor.vv v1, v17, v1
  3316: 138e0f0a                 addi t3, t6, 0xa0
  3320: 870e0e02                 vle8.v v29, (t3)
  3324: 138e0f0b                 addi t3, t6, 0xb0
  3328: 070f0e02                 vle8.v v30, (t3)
  3332: d78cf827                 vand.vv v25, v31, v17
  3336: 573d12a3                 vsrl.vi v26, v17, 4
  3340: d78ddc33                 vrgather.vv v27, v29, v25
  3344: 570eed33                 vrgather.vv v28, v30, v26
  3348: d78dcd2f                 vxor.vv v27, v28, v27
  3352: d784b42f                 vxor.vv v9, v27, v9
  3356: 83330503                 ld t2, 0x30(a0)
  3360: b3836300                 add t2, t2, t1
  3364: 87880300                 vle8.v v17, (t2), v0.t
  3368: d780102f                 vxor.vv v1, v17, v1
  3372: 138e0f0c                 addi t3, t6, 0xc0
  3376: 870e0e02                 vle8.v v29, (t3)
  3380: 138e0f0d                 addi t3, t6, 0xd0
  3384: 070f0e02                 vle8.v v30, (t3)
  3388: d78cf827                 vand.vv v25, v31, v17
  3392: 573d12a3                 vsrl.vi v26, v17, 4
  3396: d78ddc33                 vrgather.vv v27, v29, v25
  3400: 570eed33                 vrgather.vv v28, v30, v26
  3404: d78dcd2f                 vxor.vv v27, v28, v27
  3408: d784b42f                 vxor.vv v9, v27, v9
  3412: 83338503                 ld t2, 0x38(a0)
  3416: b3836300                 add t2, t2, t1
  3420: 87880300                 vle8.v v17, (t2), v0.t
  3424: d780102f                 vxor.vv v1, v17, v1
  3428: d784142f                 vxor.vv v9, v17, v9
  3432: 83b30500                 ld t2, 0(a1)
  3436: b3836300                 add t2, t2, t1
  3440: a7800300                 vse8.v v1, (t2), v0.t
  3444: 83b38500                 ld t2, 8(a1)
  3448: b3836300                 add t2, t2, t1
  3452: a7840300                 vse8.v v9, (t2), v0.t
  3456: 33035300                 add t1, t1, t0
  3460: b3865640                 sub a3, a3, t0
  3464: 6344d000                 bgtz a3, 0xd90
  3468: 6f00001d                 j 0xf5c
  3472: b3835640                 sub t2, a3, t0
  3476: 13def343                 srai t3, t2, 0x3f
  3480: b3f3c301                 and t2, t2, t3
  3484: b3837200                 add t2, t0, t2
  3488: d7a80852                 vid.v v17
  3492: 57c0136b                 vmsltu.vx v0, v17, t2
  3496: 83330500                 ld t2, 0(a0)
  3500: b3836300                 add t2, t2, t1
  3504: 87880300                 vle8.v v17, (t2), v0.t
  3508: d780085e                 vmv.v.v v1, v17
  3512: 138e0f00                 mv t3, t6
  3516: 870e0e02                 vle8.v v29, (t3)
  3520: 138e0f01                 addi t3, t6, 0x10
  3524: 070f0e02                 vle8.v v30, (t3)
  3528: d78cf827                 vand.vv v25, v31, v17
  3532: 573d12a3                 vsrl.vi v26, v17, 4
  3536: d78ddc33                 vrgather.vv v27, v29, v25
  3540: 570eed33                 vrgather.vv v28, v30, v26
  3544: d78dcd2f                 vxor.vv v27, v28, v27
  3548: d7840d5e                 vmv.v.v v9, v27
  3552: 83338500                 ld t2, 8(a0)
  3556: b3836300                 add t2, t2, t1
  3560: 87880300                 vle8.v v17, (t2), v0.t
  3564: d780102f                 vxor.vv v1, v17, v1
  3568: 138e0f02                 addi t3, t6, 0x20
  3572: 870e0e02                 vle8.v v29, (t3)
  3576: 138e0f03                 addi t3, t6, 0x30
  3580: 070f0e02                 vle8.v v30, (t3)
  3584: d78cf827                 vand.vv v25, v31, v17
  3588: 573d12a3                 vsrl.vi v26, v17, 4
  3592: d78ddc33                 vrgather.vv v27, v29, v25
  3596: 570eed33                 vrgather.vv v28, v30, v26
  3600: d78dcd2f                 vxor.vv v27, v28, v27
  3604: d784b42f                 vxor.vv v9, v27, v9
  3608: 83330501                 ld t2, 0x10(a0)
  3612: b3836300                 add t2, t2, t1
  3616: 87880300                 vle8.v v17, (t2), v0.t
  3620: d780102f                 vxor.vv v1, v17, v1
  3624: 138e0f04                 addi t3, t6, 0x40
  3628: 870e0e02                 vle8.v v29, (t3)
  3632: 138e0f05                 addi t3, t6, 0x50
  3636: 070f0e02                 vle8.v v30, (t3)
  3640: d78cf827                 vand.vv v25, v31, v17
  3644: 573d12a3                 vsrl.vi v26, v17, 4
  3648: d78ddc33                 vrgather.vv v27, v29, v25
  3652: 570eed33                 vrgather.vv v28, v30, v26
  3656: d78dcd2f                 vxor.vv v27, v28, v27
  3660: d784b42f                 vxor.vv v9, v27, v9
  3664: 83338501                 ld t2, 0x18(a0)
  3668: b3836300                 add t2, t2, t1
  3672: 87880300                 vle8.v v17, (t2), v0.t
  3676: d780102f                 vxor.vv v1, v17, v1
  3680: 138e0f06                 addi t3, t6, 0x60
  3684: 870e0e02                 vle8.v v29, (t3)
  3688: 138e0f07                 addi t3, t6, 0x70
  3692: 070f0e02                 vle8.v v30, (t3)
  3696: d78cf827                 vand.vv v25, v31, v17
  3700: 573d12a3                 vsrl.vi v26, v17, 4
  3704: d78ddc33                 vrgather.vv v27, v29, v25
  3708: 570eed33                 vrgather.vv v28, v30, v26
  3712: d78dcd2f                 vxor.vv v27, v28, v27
  3716: d784b42f                 vxor.vv v9, v27, v9
  3720: 83330502                 ld t2, 0x20(a0)
  3724: b3836300                 add t2, t2, t1
  3728: 87880300                 vle8.v v17, (t2), v0.t
  3732: d780102f                 vxor.vv v1, v17, v1
  3736: 138e0f08                 addi t3, t6, 0x80
  3740: 870e0e02                 vle8.v v29, (t3)
  3744: 138e0f09                 addi t3, t6, 0x90
  3748: 070f0e02                 vle8.v v30, (t3)
  3752: d78cf827                 vand.vv v25, v31, v17
  3756: 573d12a3                 vsrl.vi v26, v17, 4
  3760: d78ddc33                 vrgather.vv v27, v29, v25
  3764: 570eed33                 vrgather.vv v28, v30, v26
  3768: d78dcd2f                 vxor.vv v27, v28, v27
  3772: d784b42f                 vxor.vv v9, v27, v9
  3776: 83338502                 ld t2, 0x28(a0)
  3780: b3836300                 add t2, t2, t1
  3784: 87880300                 vle8.v v17, (t2), v0.t
  3788: d780102f                 vxor.vv v1, v17, v1
  3792: 138e0f0a                 addi t3, t6, 0xa0
  3796: 870e0e02                 vle8.v v29, (t3)
  3800: 138e0f0b                 addi t3, t6, 0xb0
  3804: 070f0e02                 vle8.v v30, (t3)
  3808: d78cf827                 vand.vv v25, v31, v17
  3812: 573d12a3                 vsrl.vi v26, v17, 4
  3816: d78ddc33                 vrgather.vv v27, v29, v25
  3820: 570eed33                 vrgather.vv v28, v30, v26
  3824: d78dcd2f                 vxor.vv v27, v28, v27
  3828: d784b42f                 vxor.vv v9, v27, v9
  3832: 83330503                 ld t2, 0x30(a0)
  3836: b3836300                 add t2, t2, t1
  3840: 87880300                 vle8.v v17, (t2), v0.t
  3844: d780102f                 vxor.vv v1, v17, v1
  3848: 138e0f0c                 addi t3, t6, 0xc0
  3852: 870e0e02                 vle8.v v29, (t3)
  3856: 138e0f0d                 addi t3, t6, 0xd0
  3860: 070f0e02                 vle8.v v30, (t3)
  3864: d78cf827                 vand.vv v25, v31, v17
  3868: 573d12a3                 vsrl.vi v26, v17, 4
  3872: d78ddc33                 vrgather.vv v27, v29, v25
  3876: 570eed33                 vrgather.vv v28, v30, v26
  3880: d78dcd2f                 vxor.vv v27, v28, v27
  3884: d784b42f                 vxor.vv v9, v27, v9
  3888: 83338503                 ld t2, 0x38(a0)
  3892: b3836300                 add t2, t2, t1
  3896: 87880300                 vle8.v v17, (t2), v0.t
  3900: d780102f                 vxor.vv v1, v17, v1
  3904: d784142f                 vxor.vv v9, v17, v9
  3908: 83b30500                 ld t2, 0(a1)
  3912: b3836300                 add t2, t2, t1
  3916: a7800300                 vse8.v v1, (t2), v0.t
  3920: 83b38500                 ld t2, 8(a1)
  3924: b3836300                 add t2, t2, t1
  3928: a7840300                 vse8.v v9, (t2), v0.t
  3932: 33035300                 add t1, t1, t0
  3936: b3865640                 sub a3, a3, t0
  3940: 6344d000                 bgtz a3, 0xf6c
  3944: 6f00001d                 j 0x1138
  3948: b3835640                 sub t2, a3, t0
  3952: 13def343                 srai t3, t2, 0x3f
  3956: b3f3c301                 and t2, t2, t3
  3960: b3837200                 add t2, t0, t2
  3964: d7a80852                 vid.v v17
  3968: 57c0136b                 vmsltu.vx v0, v17, t2
  3972: 83330500                 ld t2, 0(a0)
  3976: b3836300                 add t2, t2, t1
  3980: 87880300                 vle8.v v17, (t2), v0.t
  3984: d780085e                 vmv.v.v v1, v17
  3988: 138e0f00                 mv t3, t6
  3992: 870e0e02                 vle8.v v29, (t3)
  3996: 138e0f01                 addi t3, t6, 0x10
  4000: 070f0e02                 vle8.v v30, (t3)
  4004: d78cf827                 vand.vv v25, v31, v17
  4008: 573d12a3                 vsrl.vi v26, v17, 4
  4012: d78ddc33                 vrgather.vv v27, v29, v25
  4016: 570eed33                 vrgather.vv v28, v30, v26
  4020: d78dcd2f                 vxor.vv v27, v28, v27
  4024: d7840d5e                 vmv.v.v v9, v27
  4028: 83338500                 ld t2, 8(a0)
  4032: b3836300                 add t2, t2, t1
  4036: 87880300                 vle8.v v17, (t2), v0.t
  4040: d780102f                 vxor.vv v1, v17, v1
  4044: 138e0f02                 addi t3, t6, 0x20
  4048: 870e0e02                 vle8.v v29, (t3)
  4052: 138e0f03                 addi t3, t6, 0x30
  4056: 070f0e02                 vle8.v v30, (t3)
  4060: d78cf827                 vand.vv v25, v31, v17
  4064: 573d12a3                 vsrl.vi v26, v17, 4
  4068: d78ddc33                 vrgather.vv v27, v29, v25
  4072: 570eed33                 vrgather.vv v28, v30, v26
  4076: d78dcd2f                 vxor.vv v27, v28, v27
  4080: d784b42f                 vxor.vv v9, v27, v9
  4084: 83330501                 ld t2, 0x10(a0)
  4088: b3836300                 add t2, t2, t1
  4092: 87880300                 vle8.v v17, (t2), v0.t
  4096: d780102f                 vxor.vv v1, v17, v1
  4100: 138e0f04                 addi t3, t6, 0x40
  4104: 870e0e02                 vle8.v v29, (t3)
  4108: 138e0f05                 addi t3, t6, 0x50
  4112: 070f0e02                 vle8.v v30, (t3)
  4116: d78cf827                 vand.vv v25, v31, v17
  4120: 573d12a3                 vsrl.vi v26, v17, 4
  4124: d78ddc33                 vrgather.vv v27, v29, v25
  4128: 570eed33                 vrgather.vv v28, v30, v26
  4132: d78dcd2f                 vxor.vv v27, v28, v27
  4136: d784b42f                 vxor.vv v9, v27, v9
  4140: 83338501                 ld t2, 0x18(a0)
  4144: b3836300                 add t2, t2, t1
  4148: 87880300                 vle8.v v17, (t2), v0.t
  4152: d780102f                 vxor.vv v1, v17, v1
  4156: 138e0f06                 addi t3, t6, 0x60
  4160: 870e0e02                 vle8.v v29, (t3)
  4164: 138e0f07                 addi t3, t6, 0x70
  4168: 070f0e02                 vle8.v v30, (t3)
  4172: d78cf827                 vand.vv v25, v31, v17
  4176: 573d12a3                 vsrl.vi v26, v17, 4
  4180: d78ddc33                 vrgather.vv v27, v29, v25
  4184: 570eed33                 vrgather.vv v28, v30, v26
  4188: d78dcd2f                 vxor.vv v27, v28, v27
  4192: d784b42f                 vxor.vv v9, v27, v9
  4196: 83330502                 ld t2, 0x20(a0)
  4200: b3836300                 add t2, t2, t1
  4204: 87880300                 vle8.v v17, (t2), v0.t
  4208: d780102f                 vxor.vv v1, v17, v1
  4212: 138e0f08                 addi t3, t6, 0x80
  4216: 870e0e02                 vle8.v v29, (t3)
  4220: 138e0f09                 addi t3, t6, 0x90
  4224: 070f0e02                 vle8.v v30, (t3)
  4228: d78cf827                 vand.vv v25, v31, v17
  4232: 573d12a3                 vsrl.vi v26, v17, 4
  4236: d78ddc33                 vrgather.vv v27, v29, v25
  4240: 570eed33                 vrgather.vv v28, v30, v26
  4244: d78dcd2f                 vxor.vv v27, v28, v27
  4248: d784b42f                 vxor.vv v9, v27, v9
  4252: 83338502                 ld t2, 0x28(a0)
  4256: b3836300                 add t2, t2, t1
  4260: 87880300                 vle8.v v17, (t2), v0.t
  4264: d780102f                 vxor.vv v1, v17, v1
  4268: 138e0f0a                 addi t3, t6, 0xa0
  4272: 870e0e02                 vle8.v v29, (t3)
  4276: 138e0f0b                 addi t3, t6, 0xb0
  4280: 070f0e02                 vle8.v v30, (t3)
  4284: d78cf827                 vand.vv v25, v31, v17
  4288: 573d12a3                 vsrl.vi v26, v17, 4
  4292: d78ddc33                 vrgather.vv v27, v29, v25
  4296: 570eed33                 vrgather.vv v28, v30, v26
  4300: d78dcd2f                 vxor.vv v27, v28, v27
  4304: d784b42f                 vxor.vv v9, v27, v9
  4308: 83330503                 ld t2, 0x30(a0)
  4312: b3836300                 add t2, t2, t1
  4316: 87880300                 vle8.v v17, (t2), v0.t
  4320: d780102f                 vxor.vv v1, v17, v1
  4324: 138e0f0c                 addi t3, t6, 0xc0
  4328: 870e0e02                 vle8.v v29, (t3)
  4332: 138e0f0d                 addi t3, t6, 0xd0
  4336: 070f0e02                 vle8.v v30, (t3)
  4340: d78cf827                 vand.vv v25, v31, v17
  4344: 573d12a3                 vsrl.vi v26, v17, 4
  4348: d78ddc33                 vrgather.vv v27, v29, v25
  4352: 570eed33                 vrgather.vv v28, v30, v26
  4356: d78dcd2f                 vxor.vv v27, v28, v27
  4360: d784b42f                 vxor.vv v9, v27, v9
  4364: 83338503                 ld t2, 0x38(a0)
  4368: b3836300                 add t2, t2, t1
  4372: 87880300                 vle8.v v17, (t2), v0.t
  4376: d780102f                 vxor.vv v1, v17, v1
  4380: d784142f                 vxor.vv v9, v17, v9
  4384: 83b30500                 ld t2, 0(a1)
  4388: b3836300                 add t2, t2, t1
  4392: a7800300                 vse8.v v1, (t2), v0.t
  4396: 83b38500                 ld t2, 8(a1)
  4400: b3836300                 add t2, t2, t1
  4404: a7840300                 vse8.v v9, (t2), v0.t
  4408: 33035300                 add t1, t1, t0
  4412: b3865640                 sub a3, a3, t0
  4416: 6344d000                 bgtz a3, 0x1148
  4420: 6f00001d                 j 0x1314
  4424: b3835640                 sub t2, a3, t0
  4428: 13def343                 srai t3, t2, 0x3f
  4432: b3f3c301                 and t2, t2, t3
  4436: b3837200                 add t2, t0, t2
  4440: d7a80852                 vid.v v17
  4444: 57c0136b                 vmsltu.vx v0, v17, t2
  4448: 83330500                 ld t2, 0(a0)
  4452: b3836300                 add t2, t2, t1
  4456: 87880300                 vle8.v v17, (t2), v0.t
  4460: d780085e                 vmv.v.v v1, v17
  4464: 138e0f00                 mv t3, t6
  4468: 870e0e02                 vle8.v v29, (t3)
  4472: 138e0f01                 addi t3, t6, 0x10
  4476: 070f0e02                 vle8.v v30, (t3)
  4480: d78cf827                 vand.vv v25, v31, v17
  4484: 573d12a3                 vsrl.vi v26, v17, 4
  4488: d78ddc33                 vrgather.vv v27, v29, v25
  4492: 570eed33                 vrgather.vv v28, v30, v26
  4496: d78dcd2f                 vxor.vv v27, v28, v27
  4500: d7840d5e                 vmv.v.v v9, v27
  4504: 83338500                 ld t2, 8(a0)
  4508: b3836300                 add t2, t2, t1
  4512: 87880300                 vle8.v v17, (t2), v0.t
  4516: d780102f                 vxor.vv v1, v17, v1
  4520: 138e0f02                 addi t3, t6, 0x20
  4524: 870e0e02                 vle8.v v29, (t3)
  4528: 138e0f03                 addi t3, t6, 0x30
  4532: 070f0e02                 vle8.v v30, (t3)
  4536: d78cf827                 vand.vv v25, v31, v17
  4540: 573d12a3                 vsrl.vi v26, v17, 4
  4544: d78ddc33                 vrgather.vv v27, v29, v25
  4548: 570eed33                 vrgather.vv v28, v30, v26
  4552: d78dcd2f                 vxor.vv v27, v28, v27
  4556: d784b42f                 vxor.vv v9, v27, v9
  4560: 83330501                 ld t2, 0x10(a0)
  4564: b3836300                 add t2, t2, t1
  4568: 87880300                 vle8.v v17, (t2), v0.t
  4572: d780102f                 vxor.vv v1, v17, v1
  4576: 138e0f04                 addi t3, t6, 0x40
  4580: 870e0e02                 vle8.v v29, (t3)
  4584: 138e0f05                 addi t3, t6, 0x50
  4588: 070f0e02                 vle8.v v30, (t3)
  4592: d78cf827                 vand.vv v25, v31, v17
  4596: 573d12a3                 vsrl.vi v26, v17, 4
  4600: d78ddc33                 vrgather.vv v27, v29, v25
  4604: 570eed33                 vrgather.vv v28, v30, v26
  4608: d78dcd2f                 vxor.vv v27, v28, v27
  4612: d784b42f                 vxor.vv v9, v27, v9
  4616: 83338501                 ld t2, 0x18(a0)
  4620: b3836300                 add t2, t2, t1
  4624: 87880300                 vle8.v v17, (t2), v0.t
  4628: d780102f                 vxor.vv v1, v17, v1
  4632: 138e0f06                 addi t3, t6, 0x60
  4636: 870e0e02                 vle8.v v29, (t3)
  4640: 138e0f07                 addi t3, t6, 0x70
  4644: 070f0e02                 vle8.v v30, (t3)
  4648: d78cf827                 vand.vv v25, v31, v17
  4652: 573d12a3                 vsrl.vi v26, v17, 4
  4656: d78ddc33                 vrgather.vv v27, v29, v25
  4660: 570eed33                 vrgather.vv v28, v30, v26
  4664: d78dcd2f                 vxor.vv v27, v28, v27
  4668: d784b42f                 vxor.vv v9, v27, v9
  4672: 83330502                 ld t2, 0x20(a0)
  4676: b3836300                 add t2, t2, t1
  4680: 87880300                 vle8.v v17, (t2), v0.t
  4684: d780102f                 vxor.vv v1, v17, v1
  4688: 138e0f08                 addi t3, t6, 0x80
  4692: 870e0e02                 vle8.v v29, (t3)
  4696: 138e0f09                 addi t3, t6, 0x90
  4700: 070f0e02                 vle8.v v30, (t3)
  4704: d78cf827                 vand.vv v25, v31, v17
  4708: 573d12a3                 vsrl.vi v26, v17, 4
  4712: d78ddc33                 vrgather.vv v27, v29, v25
  4716: 570eed33                 vrgather.vv v28, v30, v26
  4720: d78dcd2f                 vxor.vv v27, v28, v27
  4724: d784b42f                 vxor.vv v9, v27, v9
  4728: 83338502                 ld t2, 0x28(a0)
  4732: b3836300                 add t2, t2, t1
  4736: 87880300                 vle8.v v17, (t2), v0.t
  4740: d780102f                 vxor.vv v1, v17, v1
  4744: 138e0f0a                 addi t3, t6, 0xa0
  4748: 870e0e02                 vle8.v v29, (t3)
  4752: 138e0f0b                 addi t3, t6, 0xb0
  4756: 070f0e02                 vle8.v v30, (t3)
  4760: d78cf827                 vand.vv v25, v31, v17
  4764: 573d12a3                 vsrl.vi v26, v17, 4
  4768: d78ddc33                 vrgather.vv v27, v29, v25
  4772: 570eed33                 vrgather.vv v28, v30, v26
  4776: d78dcd2f                 vxor.vv v27, v28, v27
  4780: d784b42f                 vxor.vv v9, v27, v9
  4784: 83330503                 ld t2, 0x30(a0)
  4788: b3836300                 add t2, t2, t1
  4792: 87880300                 vle8.v v17, (t2), v0.t
  4796: d780102f                 vxor.vv v1, v17, v1
  4800: 138e0f0c                 addi t3, t6, 0xc0
  4804: 870e0e02                 vle8.v v29, (t3)
  4808: 138e0f0d                 addi t3, t6, 0xd0
  4812: 070f0e02                 vle8.v v30, (t3)
  4816: d78cf827                 vand.vv v25, v31, v17
  4820: 573d12a3                 vsrl.vi v26, v17, 4
  4824: d78ddc33                 vrgather.vv v27, v29, v25
  4828: 570eed33                 vrgather.vv v28, v30, v26
  4832: d78dcd2f                 vxor.vv v27, v28, v27
  4836: d784b42f                 vxor.vv v9, v27, v9
  4840: 83338503                 ld t2, 0x38(a0)
  4844: b3836300                 add t2, t2, t1
  4848: 87880300                 vle8.v v17, (t2), v0.t
  4852: d780102f                 vxor.vv v1, v17, v1
  4856: d784142f                 vxor.vv v9, v17, v9
  4860: 83b30500                 ld t2, 0(a1)
  4864: b3836300                 add t2, t2, t1
  4868: a7800300                 vse8.v v1, (t2), v0.t
  4872: 83b38500                 ld t2, 8(a1)
  4876: b3836300                 add t2, t2, t1
  4880: a7840300                 vse8.v v9, (t2), v0.t
  4884: 33035300                 add t1, t1, t0
  4888: b3865640                 sub a3, a3, t0
  4892: 6344d000                 bgtz a3, 0x1324
  4896: 6f00001d                 j 0x14f0
  4900: b3835640                 sub t2, a3, t0
  4904: 13def343                 srai t3, t2, 0x3f
  4908: b3f3c301                 and t2, t2, t3
  4912: b3837200                 add t2, t0, t2
  4916: d7a80852                 vid.v v17
  4920: 57c0136b                 vmsltu.vx v0, v17, t2
  4924: 83330500                 ld t2, 0(a0)
  4928: b3836300                 add t2, t2, t1
  4932: 87880300                 vle8.v v17, (t2), v0.t
  4936: d780085e                 vmv.v.v v1, v17
  4940: 138e0f00                 mv t3, t6
  4944: 870e0e02                 vle8.v v29, (t3)
  4948: 138e0f01                 addi t3, t6, 0x10
  4952: 070f0e02                 vle8.v v30, (t3)
  4956: d78cf827                 vand.vv v25, v31, v17
  4960: 573d12a3                 vsrl.vi v26, v17, 4
  4964: d78ddc33                 vrgather.vv v27, v29, v25
  4968: 570eed33                 vrgather.vv v28, v30, v26
  4972: d78dcd2f                 vxor.vv v27, v28, v27
  4976: d7840d5e                 vmv.v.v v9, v27
  4980: 83338500                 ld t2, 8(a0)
  4984: b3836300                 add t2, t2, t1
  4988: 87880300                 vle8.v v17, (t2), v0.t
  4992: d780102f                 vxor.vv v1, v17, v1
  4996: 138e0f02                 addi t3, t6, 0x20
  5000: 870e0e02                 vle8.v v29, (t3)
  5004: 138e0f03                 addi t3, t6, 0x30
  5008: 070f0e02                 vle8.v v30, (t3)
  5012: d78cf827                 vand.vv v25, v31, v17
  5016: 573d12a3                 vsrl.vi v26, v17, 4
  5020: d78ddc33                 vrgather.vv v27, v29, v25
  5024: 570eed33                 vrgather.vv v28, v30, v26
  5028: d78dcd2f                 vxor.vv v27, v28, v27
  5032: d784b42f                 vxor.vv v9, v27, v9
  5036: 83330501                 ld t2, 0x10(a0)
  5040: b3836300                 add t2, t2, t1
  5044: 87880300                 vle8.v v17, (t2), v0.t
  5048: d780102f                 vxor.vv v1, v17, v1
  5052: 138e0f04                 addi t3, t6, 0x40
  5056: 870e0e02                 vle8.v v29, (t3)
  5060: 138e0f05                 addi t3, t6, 0x50
  5064: 070f0e02                 vle8.v v30, (t3)
  5068: d78cf827                 vand.vv v25, v31, v17
  5072: 573d12a3                 vsrl.vi v26, v17, 4
  5076: d78ddc33                 vrgather.vv v27, v29, v25
  5080: 570eed33                 vrgather.vv v28, v30, v26
  5084: d78dcd2f                 vxor.vv v27, v28, v27
  5088: d784b42f                 vxor.vv v9, v27, v9
  5092: 83338501                 ld t2, 0x18(a0)
  5096: b3836300                 add t2, t2, t1
  5100: 87880300                 vle8.v v17, (t2), v0.t
  5104: d780102f                 vxor.vv v1, v17, v1
  5108: 138e0f06                 addi t3, t6, 0x60
  5112: 870e0e02                 vle8.v v29, (t3)
  5116: 138e0f07                 addi t3, t6, 0x70
  5120: 070f0e02                 vle8.v v30, (t3)
  5124: d78cf827                 vand.vv v25, v31, v17
  5128: 573d12a3                 vsrl.vi v26, v17, 4
  5132: d78ddc33                 vrgather.vv v27, v29, v25
  5136: 570eed33                 vrgather.vv v28, v30, v26
  5140: d78dcd2f                 vxor.vv v27, v28, v27
  5144: d784b42f                 vxor.vv v9, v27, v9
  5148: 83330502                 ld t2, 0x20(a0)
  5152: b3836300                 add t2, t2, t1
  5156: 87880300                 vle8.v v17, (t2), v0.t
  5160: d780102f                 vxor.vv v1, v17, v1
  5164: 138e0f08                 addi t3, t6, 0x80
  5168: 870e0e02                 vle8.v v29, (t3)
  5172: 138e0f09                 addi t3, t6, 0x90
  5176: 070f0e02                 vle8.v v30, (t3)
  5180: d78cf827                 vand.vv v25, v31, v17
  5184: 573d12a3                 vsrl.vi v26, v17, 4
  5188: d78ddc33                 vrgather.vv v27, v29, v25
  5192: 570eed33                 vrgather.vv v28, v30, v26
  5196: d78dcd2f                 vxor.vv v27, v28, v27
  5200: d784b42f                 vxor.vv v9, v27, v9
  5204: 83338502                 ld t2, 0x28(a0)
  5208: b3836300                 add t2, t2, t1
  5212: 87880300                 vle8.v v17, (t2), v0.t
  5216: d780102f                 vxor.vv v1, v17, v1
  5220: 138e0f0a                 addi t3, t6, 0xa0
  5224: 870e0e02                 vle8.v v29, (t3)
  5228: 138e0f0b                 addi t3, t6, 0xb0
  5232: 070f0e02                 vle8.v v30, (t3)
  5236: d78cf827                 vand.vv v25, v31, v17
  5240: 573d12a3                 vsrl.vi v26, v17, 4
  5244: d78ddc33                 vrgather.vv v27, v29, v25
  5248: 570eed33                 vrgather.vv v28, v30, v26
  5252: d78dcd2f                 vxor.vv v27, v28, v27
  5256: d784b42f                 vxor.vv v9, v27, v9
  5260: 83330503                 ld t2, 0x30(a0)
  5264: b3836300                 add t2, t2, t1
  5268: 87880300                 vle8.v v17, (t2), v0.t
  5272: d780102f                 vxor.vv v1, v17, v1
  5276: 138e0f0c                 addi t3, t6, 0xc0
  5280: 870e0e02                 vle8.v v29, (t3)
  5284: 138e0f0d                 addi t3, t6, 0xd0
  5288: 070f0e02                 vle8.v v30, (t3)
  5292: d78cf827                 vand.vv v25, v31, v17
  5296: 573d12a3                 vsrl.vi v26, v17, 4
  5300: d78ddc33                 vrgather.vv v27, v29, v25
  5304: 570eed33                 vrgather.vv v28, v30, v26
  5308: d78dcd2f                 vxor.vv v27, v28, v27
  5312: d784b42f                 vxor.vv v9, v27, v9
  5316: 83338503                 ld t2, 0x38(a0)
  5320: b3836300                 add t2, t2, t1
  5324: 87880300                 vle8.v v17, (t2), v0.t
  5328: d780102f                 vxor.vv v1, v17, v1
  5332: d784142f                 vxor.vv v9, v17, v9
  5336: 83b30500                 ld t2, 0(a1)
  5340: b3836300                 add t2, t2, t1
  5344: a7800300                 vse8.v v1, (t2), v0.t
  5348: 83b38500                 ld t2, 8(a1)
  5352: b3836300                 add t2, t2, t1
  5356: a7840300                 vse8.v v9, (t2), v0.t
  5360: 33035300                 add t1, t1, t0
  5364: b3865640                 sub a3, a3, t0
  5368: 6344d000                 bgtz a3, 0x1500
  5372: 6f00001d                 j 0x16cc
  5376: b3835640                 sub t2, a3, t0
  5380: 13def343                 srai t3, t2, 0x3f
  5384: b3f3c301                 and t2, t2, t3
  5388: b3837200                 add t2, t0, t2
  5392: d7a80852                 vid.v v17
  5396: 57c0136b                 vmsltu.vx v0, v17, t2
  5400: 83330500                 ld t2, 0(a0)
  5404: b3836300                 add t2, t2, t1
  5408: 87880300                 vle8.v v17, (t2), v0.t
  5412: d780085e                 vmv.v.v v1, v17
  5416: 138e0f00                 mv t3, t6
  5420: 870e0e02                 vle8.v v29, (t3)
  5424: 138e0f01                 addi t3, t6, 0x10
  5428: 070f0e02                 vle8.v v30, (t3)
  5432: d78cf827                 vand.vv v25, v31, v17
  5436: 573d12a3                 vsrl.vi v26, v17, 4
  5440: d78ddc33                 vrgather.vv v27, v29, v25
  5444: 570eed33                 vrgather.vv v28, v30, v26
  5448: d78dcd2f                 vxor.vv v27, v28, v27
  5452: d7840d5e                 vmv.v.v v9, v27
  5456: 83338500                 ld t2, 8(a0)
  5460: b3836300                 add t2, t2, t1
  5464: 87880300                 vle8.v v17, (t2), v0.t
  5468: d780102f                 vxor.vv v1, v17, v1
  5472: 138e0f02                 addi t3, t6, 0x20
  5476: 870e0e02                 vle8.v v29, (t3)
  5480: 138e0f03                 addi t3, t6, 0x30
  5484: 070f0e02                 vle8.v v30, (t3)
  5488: d78cf827                 vand.vv v25, v31, v17
  5492: 573d12a3                 vsrl.vi v26, v17, 4
  5496: d78ddc33                 vrgather.vv v27, v29, v25
  5500: 570eed33                 vrgather.vv v28, v30, v26
  5504: d78dcd2f                 vxor.vv v27, v28, v27
  5508: d784b42f                 vxor.vv v9, v27, v9
  5512: 83330501                 ld t2, 0x10(a0)
  5516: b3836300                 add t2, t2, t1
  5520: 87880300                 vle8.v v17, (t2), v0.t
  5524: d780102f                 vxor.vv v1, v17, v1
  5528: 138e0f04                 addi t3, t6, 0x40
  5532: 870e0e02                 vle8.v v29, (t3)
  5536: 138e0f05                 addi t3, t6, 0x50
  5540: 070f0e02                 vle8.v v30, (t3)
  5544: d78cf827                 vand.vv v25, v31, v17
  5548: 573d12a3                 vsrl.vi v26, v17, 4
  5552: d78ddc33                 vrgather.vv v27, v29, v25
  5556: 570eed33                 vrgather.vv v28, v30, v26
  5560: d78dcd2f                 vxor.vv v27, v28, v27
  5564: d784b42f                 vxor.vv v9, v27, v9
  5568: 83338501                 ld t2, 0x18(a0)
  5572: b3836300                 add t2, t2, t1
  5576: 87880300                 vle8.v v17, (t2), v0.t
  5580: d780102f                 vxor.vv v1, v17, v1
  5584: 138e0f06                 addi t3, t6, 0x60
  5588: 870e0e02                 vle8.v v29, (t3)
  5592: 138e0f07                 addi t3, t6, 0x70
  5596: 070f0e02                 vle8.v v30, (t3)
  5600: d78cf827                 vand.vv v25, v31, v17
  5604: 573d12a3                 vsrl.vi v26, v17, 4
  5608: d78ddc33                 vrgather.vv v27, v29, v25
  5612: 570eed33                 vrgather.vv v28, v30, v26
  5616: d78dcd2f                 vxor.vv v27, v28, v27
  5620: d784b42f                 vxor.vv v9, v27, v9
  5624: 83330502                 ld t2, 0x20(a0)
  5628: b3836300                 add t2, t2, t1
  5632: 87880300                 vle8.v v17, (t2), v0.t
  5636: d780102f                 vxor.vv v1, v17, v1
  5640: 138e0f08                 addi t3, t6, 0x80
  5644: 870e0e02                 vle8.v v29, (t3)
  5648: 138e0f09                 addi t3, t6, 0x90
  5652: 070f0e02                 vle8.v v30, (t3)
  5656: d78cf827                 vand.vv v25, v31, v17
  5660: 573d12a3                 vsrl.vi v26, v17, 4
  5664: d78ddc33                 vrgather.vv v27, v29, v25
  5668: 570eed33                 vrgather.vv v28, v30, v26
  5672: d78dcd2f                 vxor.vv v27, v28, v27
  5676: d784b42f                 vxor.vv v9, v27, v9
  5680: 83338502                 ld t2, 0x28(a0)
  5684: b3836300                 add t2, t2, t1
  5688: 87880300                 vle8.v v17, (t2), v0.t
  5692: d780102f                 vxor.vv v1, v17, v1
  5696: 138e0f0a                 addi t3, t6, 0xa0
  5700: 870e0e02                 vle8.v v29, (t3)
  5704: 138e0f0b                 addi t3, t6, 0xb0
  5708: 070f0e02                 vle8.v v30, (t3)
  5712: d78cf827                 vand.vv v25, v31, v17
  5716: 573d12a3                 vsrl.vi v26, v17, 4
  5720: d78ddc33                 vrgather.vv v27, v29, v25
  5724: 570eed33                 vrgather.vv v28, v30, v26
  5728: d78dcd2f                 vxor.vv v27, v28, v27
  5732: d784b42f                 vxor.vv v9, v27, v9
  5736: 83330503                 ld t2, 0x30(a0)
  5740: b3836300                 add t2, t2, t1
  5744: 87880300                 vle8.v v17, (t2), v0.t
  5748: d780102f                 vxor.vv v1, v17, v1
  5752: 138e0f0c                 addi t3, t6, 0xc0
  5756: 870e0e02                 vle8.v v29, (t3)
  5760: 138e0f0d                 addi t3, t6, 0xd0
  5764: 070f0e02                 vle8.v v30, (t3)
  5768: d78cf827                 vand.vv v25, v31, v17
  5772: 573d12a3                 vsrl.vi v26, v17, 4
  5776: d78ddc33                 vrgather.vv v27, v29, v25
  5780: 570eed33                 vrgather.vv v28, v30, v26
  5784: d78dcd2f                 vxor.vv v27, v28, v27
  5788: d784b42f                 vxor.vv v9, v27, v9
  5792: 83338503                 ld t2, 0x38(a0)
  5796: b3836300                 add t2, t2, t1
  5800: 87880300                 vle8.v v17, (t2), v0.t
  5804: d780102f                 vxor.vv v1, v17, v1
  5808: d784142f                 vxor.vv v9, v17, v9
  5812: 83b30500                 ld t2, 0(a1)
  5816: b3836300                 add t2, t2, t1
  5820: a7800300                 vse8.v v1, (t2), v0.t
  5824: 83b38500                 ld t2, 8(a1)
  5828: b3836300                 add t2, t2, t1
  5832: a7840300                 vse8.v v9, (t2), v0.t
  5836: 33035300                 add t1, t1, t0
  5840: b3865640                 sub a3, a3, t0
  5844: 6344d000                 bgtz a3, 0x16dc
  5848: 6f00001d                 j 0x18a8
  5852: b3835640                 sub t2, a3, t0
  5856: 13def343                 srai t3, t2, 0x3f
  5860: b3f3c301                 and t2, t2, t3
  5864: b3837200                 add t2, t0, t2
  5868: d7a80852                 vid.v v17
  5872: 57c0136b                 vmsltu.vx v0, v17, t2
  5876: 83330500                 ld t2, 0(a0)
  5880: b3836300                 add t2, t2, t1
  5884: 87880300                 vle8.v v17, (t2), v0.t
  5888: d780085e                 vmv.v.v v1, v17
  5892: 138e0f00                 mv t3, t6
  5896: 870e0e02                 vle8.v v29, (t3)
  5900: 138e0f01                 addi t3, t6, 0x10
  5904: 070f0e02                 vle8.v v30, (t3)
  5908: d78cf827                 vand.vv v25, v31, v17
  5912: 573d12a3                 vsrl.vi v26, v17, 4
  5916: d78ddc33                 vrgather.vv v27, v29, v25
  5920: 570eed33                 vrgather.vv v28, v30, v26
  5924: d78dcd2f                 vxor.vv v27, v28, v27
  5928: d7840d5e                 vmv.v.v v9, v27
  5932: 83338500                 ld t2, 8(a0)
  5936: b3836300                 add t2, t2, t1
  5940: 87880300                 vle8.v v17, (t2), v0.t
  5944: d780102f                 vxor.vv v1, v17, v1
  5948: 138e0f02                 addi t3, t6, 0x20
  5952: 870e0e02                 vle8.v v29, (t3)
  5956: 138e0f03                 addi t3, t6, 0x30
  5960: 070f0e02                 vle8.v v30, (t3)
  5964: d78cf827                 vand.vv v25, v31, v17
  5968: 573d12a3                 vsrl.vi v26, v17, 4
  5972: d78ddc33                 vrgather.vv v27, v29, v25
  5976: 570eed33                 vrgather.vv v28, v30, v26
  5980: d78dcd2f                 vxor.vv v27, v28, v27
  5984: d784b42f                 vxor.vv v9, v27, v9
  5988: 83330501                 ld t2, 0x10(a0)
  5992: b3836300                 add t2, t2, t1
  5996: 87880300                 vle8.v v17, (t2), v0.t
  6000: d780102f                 vxor.vv v1, v17, v1
  6004: 138e0f04                 addi t3, t6, 0x40
  6008: 870e0e02                 vle8.v v29, (t3)
  6012: 138e0f05                 addi t3, t6, 0x50
  6016: 070f0e02                 vle8.v v30, (t3)
  6020: d78cf827                 vand.vv v25, v31, v17
  6024: 573d12a3                 vsrl.vi v26, v17, 4
  6028: d78ddc33                 vrgather.vv v27, v29, v25
  6032: 570eed33                 vrgather.vv v28, v30, v26
  6036: d78dcd2f                 vxor.vv v27, v28, v27
  6040: d784b42f                 vxor.vv v9, v27, v9
  6044: 83338501                 ld t2, 0x18(a0)
  6048: b3836300                 add t2, t2, t1
  6052: 87880300                 vle8.v v17, (t2), v0.t
  6056: d780102f                 vxor.vv v1, v17, v1
  6060: 138e0f06                 addi t3, t6, 0x60
  6064: 870e0e02                 vle8.v v29, (t3)
  6068: 138e0f07                 addi t3, t6, 0x70
  6072: 070f0e02                 vle8.v v30, (t3)
  6076: d78cf827                 vand.vv v25, v31, v17
  6080: 573d12a3                 vsrl.vi v26, v17, 4
  6084: d78ddc33                 vrgather.vv v27, v29, v25
  6088: 570eed33                 vrgather.vv v28, v30, v26
  6092: d78dcd2f                 vxor.vv v27, v28, v27
  6096: d784b42f                 vxor.vv v9, v27, v9
  6100: 83330502                 ld t2, 0x20(a0)
  6104: b3836300                 add t2, t2, t1
  6108: 87880300                 vle8.v v17, (t2), v0.t
  6112: d780102f                 vxor.vv v1, v17, v1
  6116: 138e0f08                 addi t3, t6, 0x80
  6120: 870e0e02                 vle8.v v29, (t3)
  6124: 138e0f09                 addi t3, t6, 0x90
  6128: 070f0e02                 vle8.v v30, (t3)
  6132: d78cf827                 vand.vv v25, v31, v17
  6136: 573d12a3                 vsrl.vi v26, v17, 4
  6140: d78ddc33                 vrgather.vv v27, v29, v25
  6144: 570eed33                 vrgather.vv v28, v30, v26
  6148: d78dcd2f                 vxor.vv v27, v28, v27
  6152: d784b42f                 vxor.vv v9, v27, v9
  6156: 83338502                 ld t2, 0x28(a0)
  6160: b3836300                 add t2, t2, t1
  6164: 87880300                 vle8.v v17, (t2), v0.t
  6168: d780102f                 vxor.vv v1, v17, v1
  6172: 138e0f0a                 addi t3, t6, 0xa0
  6176: 870e0e02                 vle8.v v29, (t3)
  6180: 138e0f0b                 addi t3, t6, 0xb0
  6184: 070f0e02                 vle8.v v30, (t3)
  6188: d78cf827                 vand.vv v25, v31, v17
  6192: 573d12a3                 vsrl.vi v26, v17, 4
  6196: d78ddc33                 vrgather.vv v27, v29, v25
  6200: 570eed33                 vrgather.vv v28, v30, v26
  6204: d78dcd2f                 vxor.vv v27, v28, v27
  6208: d784b42f                 vxor.vv v9, v27, v9
  6212: 83330503                 ld t2, 0x30(a0)
  6216: b3836300                 add t2, t2, t1
  6220: 87880300                 vle8.v v17, (t2), v0.t
  6224: d780102f                 vxor.vv v1, v17, v1
  6228: 138e0f0c                 addi t3, t6, 0xc0
  6232: 870e0e02                 vle8.v v29, (t3)
  6236: 138e0f0d                 addi t3, t6, 0xd0
  6240: 070f0e02                 vle8.v v30, (t3)
  6244: d78cf827                 vand.vv v25, v31, v17
  6248: 573d12a3                 vsrl.vi v26, v17, 4
  6252: d78ddc33                 vrgather.vv v27, v29, v25
  6256: 570eed33                 vrgather.vv v28, v30, v26
  6260: d78dcd2f                 vxor.vv v27, v28, v27
  6264: d784b42f                 vxor.vv v9, v27, v9
  6268: 83338503                 ld t2, 0x38(a0)
  6272: b3836300                 add t2, t2, t1
  6276: 87880300                 vle8.v v17, (t2), v0.t
  6280: d780102f                 vxor.vv v1, v17, v1
  6284: d784142f                 vxor.vv v9, v17, v9
  6288: 83b30500                 ld t2, 0(a1)
  6292: b3836300                 add t2, t2, t1
  6296: a7800300                 vse8.v v1, (t2), v0.t
  6300: 83b38500                 ld t2, 8(a1)
  6304: b3836300                 add t2, t2, t1
  6308: a7840300                 vse8.v v9, (t2), v0.t
  6312: 33035300                 add t1, t1, t0
  6316: b3865640                 sub a3, a3, t0
  6320: 67800000                 ret 
  6324: 0000                     unimp 
  6326: 0000                     unimp 
  6328: 0000                     unimp 
  6330: 0000                     unimp 
  6332: 0000                     unimp 
  6334: 0000                     unimp 
