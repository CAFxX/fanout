; config_t1_1566 target=rvv twin=maskedtail topo=liberation(8)/enc op=encode U=1
; code_size=4240 twins=["maskedtail"] uops_per_byte=37.3125
     0: 63140600                 bnez a2, 8
     4: 6f104008                 j 0x1088
     8: 971f0000                 auipc t6, 1
    12: 938f8f08                 addi t6, t6, 0x88
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f00d002                 j 0x854
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 878b0302                 vle8.v v23, (t2)
    76: d7800b5e                 vmv.v.v v1, v23
    80: 57860b5e                 vmv.v.v v12, v23
    84: 83338500                 ld t2, 8(a0)
    88: b3836300                 add t2, t2, t1
    92: 878b0302                 vle8.v v23, (t2)
    96: 57810b5e                 vmv.v.v v2, v23
   100: d7860b5e                 vmv.v.v v13, v23
   104: 83330501                 ld t2, 0x10(a0)
   108: b3836300                 add t2, t2, t1
   112: 878b0302                 vle8.v v23, (t2)
   116: d7810b5e                 vmv.v.v v3, v23
   120: 57870b5e                 vmv.v.v v14, v23
   124: 83338501                 ld t2, 0x18(a0)
   128: b3836300                 add t2, t2, t1
   132: 878b0302                 vle8.v v23, (t2)
   136: 57820b5e                 vmv.v.v v4, v23
   140: d7870b5e                 vmv.v.v v15, v23
   144: 83330502                 ld t2, 0x20(a0)
   148: b3836300                 add t2, t2, t1
   152: 878b0302                 vle8.v v23, (t2)
   156: d7820b5e                 vmv.v.v v5, v23
   160: 57880b5e                 vmv.v.v v16, v23
   164: 83338502                 ld t2, 0x28(a0)
   168: b3836300                 add t2, t2, t1
   172: 878b0302                 vle8.v v23, (t2)
   176: 57830b5e                 vmv.v.v v6, v23
   180: d7880b5e                 vmv.v.v v17, v23
   184: 83330503                 ld t2, 0x30(a0)
   188: b3836300                 add t2, t2, t1
   192: 878b0302                 vle8.v v23, (t2)
   196: d7830b5e                 vmv.v.v v7, v23
   200: 57890b5e                 vmv.v.v v18, v23
   204: 83338503                 ld t2, 0x38(a0)
   208: b3836300                 add t2, t2, t1
   212: 878b0302                 vle8.v v23, (t2)
   216: 57840b5e                 vmv.v.v v8, v23
   220: d7890b5e                 vmv.v.v v19, v23
   224: 83330504                 ld t2, 0x40(a0)
   228: b3836300                 add t2, t2, t1
   232: 878b0302                 vle8.v v23, (t2)
   236: d7840b5e                 vmv.v.v v9, v23
   240: 578a0b5e                 vmv.v.v v20, v23
   244: 83338504                 ld t2, 0x48(a0)
   248: b3836300                 add t2, t2, t1
   252: 878b0302                 vle8.v v23, (t2)
   256: 57850b5e                 vmv.v.v v10, v23
   260: d78a0b5e                 vmv.v.v v21, v23
   264: 83330505                 ld t2, 0x50(a0)
   268: b3836300                 add t2, t2, t1
   272: 878b0302                 vle8.v v23, (t2)
   276: d7850b5e                 vmv.v.v v11, v23
   280: 578b0b5e                 vmv.v.v v22, v23
   284: 83338505                 ld t2, 0x58(a0)
   288: b3836300                 add t2, t2, t1
   292: 878b0302                 vle8.v v23, (t2)
   296: d780702f                 vxor.vv v1, v23, v1
   300: 570b7b2f                 vxor.vv v22, v23, v22
   304: 83330506                 ld t2, 0x60(a0)
   308: b3836300                 add t2, t2, t1
   312: 878b0302                 vle8.v v23, (t2)
   316: 5701712f                 vxor.vv v2, v23, v2
   320: 5706762f                 vxor.vv v12, v23, v12
   324: 83338506                 ld t2, 0x68(a0)
   328: b3836300                 add t2, t2, t1
   332: 878b0302                 vle8.v v23, (t2)
   336: d781712f                 vxor.vv v3, v23, v3
   340: d786762f                 vxor.vv v13, v23, v13
   344: 83330507                 ld t2, 0x70(a0)
   348: b3836300                 add t2, t2, t1
   352: 878b0302                 vle8.v v23, (t2)
   356: 5702722f                 vxor.vv v4, v23, v4
   360: 5707772f                 vxor.vv v14, v23, v14
   364: 83338507                 ld t2, 0x78(a0)
   368: b3836300                 add t2, t2, t1
   372: 878b0302                 vle8.v v23, (t2)
   376: d782722f                 vxor.vv v5, v23, v5
   380: d787772f                 vxor.vv v15, v23, v15
   384: 83330508                 ld t2, 0x80(a0)
   388: b3836300                 add t2, t2, t1
   392: 878b0302                 vle8.v v23, (t2)
   396: 5703732f                 vxor.vv v6, v23, v6
   400: 5708782f                 vxor.vv v16, v23, v16
   404: d788782f                 vxor.vv v17, v23, v17
   408: 83338508                 ld t2, 0x88(a0)
   412: b3836300                 add t2, t2, t1
   416: 878b0302                 vle8.v v23, (t2)
   420: d783732f                 vxor.vv v7, v23, v7
   424: d788782f                 vxor.vv v17, v23, v17
   428: 83330509                 ld t2, 0x90(a0)
   432: b3836300                 add t2, t2, t1
   436: 878b0302                 vle8.v v23, (t2)
   440: 5704742f                 vxor.vv v8, v23, v8
   444: 5709792f                 vxor.vv v18, v23, v18
   448: 83338509                 ld t2, 0x98(a0)
   452: b3836300                 add t2, t2, t1
   456: 878b0302                 vle8.v v23, (t2)
   460: d784742f                 vxor.vv v9, v23, v9
   464: d789792f                 vxor.vv v19, v23, v19
   468: 8333050a                 ld t2, 0xa0(a0)
   472: b3836300                 add t2, t2, t1
   476: 878b0302                 vle8.v v23, (t2)
   480: 5705752f                 vxor.vv v10, v23, v10
   484: 570a7a2f                 vxor.vv v20, v23, v20
   488: 8333850a                 ld t2, 0xa8(a0)
   492: b3836300                 add t2, t2, t1
   496: 878b0302                 vle8.v v23, (t2)
   500: d785752f                 vxor.vv v11, v23, v11
   504: d78a7a2f                 vxor.vv v21, v23, v21
   508: 8333050b                 ld t2, 0xb0(a0)
   512: b3836300                 add t2, t2, t1
   516: 878b0302                 vle8.v v23, (t2)
   520: d780702f                 vxor.vv v1, v23, v1
   524: d78a7a2f                 vxor.vv v21, v23, v21
   528: 570b7b2f                 vxor.vv v22, v23, v22
   532: 8333850b                 ld t2, 0xb8(a0)
   536: b3836300                 add t2, t2, t1
   540: 878b0302                 vle8.v v23, (t2)
   544: 5701712f                 vxor.vv v2, v23, v2
   548: 570b7b2f                 vxor.vv v22, v23, v22
   552: 8333050c                 ld t2, 0xc0(a0)
   556: b3836300                 add t2, t2, t1
   560: 878b0302                 vle8.v v23, (t2)
   564: d781712f                 vxor.vv v3, v23, v3
   568: 5706762f                 vxor.vv v12, v23, v12
   572: 8333850c                 ld t2, 0xc8(a0)
   576: b3836300                 add t2, t2, t1
   580: 878b0302                 vle8.v v23, (t2)
   584: 5702722f                 vxor.vv v4, v23, v4
   588: d786762f                 vxor.vv v13, v23, v13
   592: 8333050d                 ld t2, 0xd0(a0)
   596: b3836300                 add t2, t2, t1
   600: 878b0302                 vle8.v v23, (t2)
   604: d782722f                 vxor.vv v5, v23, v5
   608: 5707772f                 vxor.vv v14, v23, v14
   612: 8333850d                 ld t2, 0xd8(a0)
   616: b3836300                 add t2, t2, t1
   620: 878b0302                 vle8.v v23, (t2)
   624: 5703732f                 vxor.vv v6, v23, v6
   628: d787772f                 vxor.vv v15, v23, v15
   632: 8333050e                 ld t2, 0xe0(a0)
   636: b3836300                 add t2, t2, t1
   640: 878b0302                 vle8.v v23, (t2)
   644: d783732f                 vxor.vv v7, v23, v7
   648: 5708782f                 vxor.vv v16, v23, v16
   652: 8333850e                 ld t2, 0xe8(a0)
   656: b3836300                 add t2, t2, t1
   660: 878b0302                 vle8.v v23, (t2)
   664: 5704742f                 vxor.vv v8, v23, v8
   668: d788782f                 vxor.vv v17, v23, v17
   672: 8333050f                 ld t2, 0xf0(a0)
   676: b3836300                 add t2, t2, t1
   680: 878b0302                 vle8.v v23, (t2)
   684: d784742f                 vxor.vv v9, v23, v9
   688: 5709792f                 vxor.vv v18, v23, v18
   692: 8333850f                 ld t2, 0xf8(a0)
   696: b3836300                 add t2, t2, t1
   700: 878b0302                 vle8.v v23, (t2)
   704: 5705752f                 vxor.vv v10, v23, v10
   708: d789792f                 vxor.vv v19, v23, v19
   712: 83330510                 ld t2, 0x100(a0)
   716: b3836300                 add t2, t2, t1
   720: 878b0302                 vle8.v v23, (t2)
   724: d785752f                 vxor.vv v11, v23, v11
   728: 570a7a2f                 vxor.vv v20, v23, v20
   732: 83338510                 ld t2, 0x108(a0)
   736: b3836300                 add t2, t2, t1
   740: 878b0302                 vle8.v v23, (t2)
   744: d780702f                 vxor.vv v1, v23, v1
   748: 570a7a2f                 vxor.vv v20, v23, v20
   752: 83330511                 ld t2, 0x110(a0)
   756: b3836300                 add t2, t2, t1
   760: 878b0302                 vle8.v v23, (t2)
   764: 5701712f                 vxor.vv v2, v23, v2
   768: d78a7a2f                 vxor.vv v21, v23, v21
   772: 83338511                 ld t2, 0x118(a0)
   776: b3836300                 add t2, t2, t1
   780: 878b0302                 vle8.v v23, (t2)
   784: d781712f                 vxor.vv v3, v23, v3
   788: 570b7b2f                 vxor.vv v22, v23, v22
   792: 83330512                 ld t2, 0x120(a0)
   796: b3836300                 add t2, t2, t1
   800: 878b0302                 vle8.v v23, (t2)
   804: 5702722f                 vxor.vv v4, v23, v4
   808: 5706762f                 vxor.vv v12, v23, v12
   812: 83338512                 ld t2, 0x128(a0)
   816: b3836300                 add t2, t2, t1
   820: 878b0302                 vle8.v v23, (t2)
   824: d782722f                 vxor.vv v5, v23, v5
   828: d786762f                 vxor.vv v13, v23, v13
   832: 83330513                 ld t2, 0x130(a0)
   836: b3836300                 add t2, t2, t1
   840: 878b0302                 vle8.v v23, (t2)
   844: 5703732f                 vxor.vv v6, v23, v6
   848: 5707772f                 vxor.vv v14, v23, v14
   852: 83338513                 ld t2, 0x138(a0)
   856: b3836300                 add t2, t2, t1
   860: 878b0302                 vle8.v v23, (t2)
   864: d783732f                 vxor.vv v7, v23, v7
   868: d787772f                 vxor.vv v15, v23, v15
   872: 5708782f                 vxor.vv v16, v23, v16
   876: 83330514                 ld t2, 0x140(a0)
   880: b3836300                 add t2, t2, t1
   884: 878b0302                 vle8.v v23, (t2)
   888: 5704742f                 vxor.vv v8, v23, v8
   892: 5708782f                 vxor.vv v16, v23, v16
   896: 83338514                 ld t2, 0x148(a0)
   900: b3836300                 add t2, t2, t1
   904: 878b0302                 vle8.v v23, (t2)
   908: d784742f                 vxor.vv v9, v23, v9
   912: d788782f                 vxor.vv v17, v23, v17
   916: 83330515                 ld t2, 0x150(a0)
   920: b3836300                 add t2, t2, t1
   924: 878b0302                 vle8.v v23, (t2)
   928: 5705752f                 vxor.vv v10, v23, v10
   932: 5709792f                 vxor.vv v18, v23, v18
   936: 83338515                 ld t2, 0x158(a0)
   940: b3836300                 add t2, t2, t1
   944: 878b0302                 vle8.v v23, (t2)
   948: d785752f                 vxor.vv v11, v23, v11
   952: d789792f                 vxor.vv v19, v23, v19
   956: 83330516                 ld t2, 0x160(a0)
   960: b3836300                 add t2, t2, t1
   964: 878b0302                 vle8.v v23, (t2)
   968: d780702f                 vxor.vv v1, v23, v1
   972: d789792f                 vxor.vv v19, v23, v19
   976: 83338516                 ld t2, 0x168(a0)
   980: b3836300                 add t2, t2, t1
   984: 878b0302                 vle8.v v23, (t2)
   988: 5701712f                 vxor.vv v2, v23, v2
   992: 570a7a2f                 vxor.vv v20, v23, v20
   996: d78a7a2f                 vxor.vv v21, v23, v21
  1000: 83330517                 ld t2, 0x170(a0)
  1004: b3836300                 add t2, t2, t1
  1008: 878b0302                 vle8.v v23, (t2)
  1012: d781712f                 vxor.vv v3, v23, v3
  1016: d78a7a2f                 vxor.vv v21, v23, v21
  1020: 83338517                 ld t2, 0x178(a0)
  1024: b3836300                 add t2, t2, t1
  1028: 878b0302                 vle8.v v23, (t2)
  1032: 5702722f                 vxor.vv v4, v23, v4
  1036: 570b7b2f                 vxor.vv v22, v23, v22
  1040: 83330518                 ld t2, 0x180(a0)
  1044: b3836300                 add t2, t2, t1
  1048: 878b0302                 vle8.v v23, (t2)
  1052: d782722f                 vxor.vv v5, v23, v5
  1056: 5706762f                 vxor.vv v12, v23, v12
  1060: 83338518                 ld t2, 0x188(a0)
  1064: b3836300                 add t2, t2, t1
  1068: 878b0302                 vle8.v v23, (t2)
  1072: 5703732f                 vxor.vv v6, v23, v6
  1076: d786762f                 vxor.vv v13, v23, v13
  1080: 83330519                 ld t2, 0x190(a0)
  1084: b3836300                 add t2, t2, t1
  1088: 878b0302                 vle8.v v23, (t2)
  1092: d783732f                 vxor.vv v7, v23, v7
  1096: 5707772f                 vxor.vv v14, v23, v14
  1100: 83338519                 ld t2, 0x198(a0)
  1104: b3836300                 add t2, t2, t1
  1108: 878b0302                 vle8.v v23, (t2)
  1112: 5704742f                 vxor.vv v8, v23, v8
  1116: d787772f                 vxor.vv v15, v23, v15
  1120: 8333051a                 ld t2, 0x1a0(a0)
  1124: b3836300                 add t2, t2, t1
  1128: 878b0302                 vle8.v v23, (t2)
  1132: d784742f                 vxor.vv v9, v23, v9
  1136: 5708782f                 vxor.vv v16, v23, v16
  1140: 8333851a                 ld t2, 0x1a8(a0)
  1144: b3836300                 add t2, t2, t1
  1148: 878b0302                 vle8.v v23, (t2)
  1152: 5705752f                 vxor.vv v10, v23, v10
  1156: d788782f                 vxor.vv v17, v23, v17
  1160: 8333051b                 ld t2, 0x1b0(a0)
  1164: b3836300                 add t2, t2, t1
  1168: 878b0302                 vle8.v v23, (t2)
  1172: d785752f                 vxor.vv v11, v23, v11
  1176: 5709792f                 vxor.vv v18, v23, v18
  1180: 8333851b                 ld t2, 0x1b8(a0)
  1184: b3836300                 add t2, t2, t1
  1188: 878b0302                 vle8.v v23, (t2)
  1192: d780702f                 vxor.vv v1, v23, v1
  1196: 5709792f                 vxor.vv v18, v23, v18
  1200: 8333051c                 ld t2, 0x1c0(a0)
  1204: b3836300                 add t2, t2, t1
  1208: 878b0302                 vle8.v v23, (t2)
  1212: 5701712f                 vxor.vv v2, v23, v2
  1216: d789792f                 vxor.vv v19, v23, v19
  1220: 8333851c                 ld t2, 0x1c8(a0)
  1224: b3836300                 add t2, t2, t1
  1228: 878b0302                 vle8.v v23, (t2)
  1232: d781712f                 vxor.vv v3, v23, v3
  1236: 570a7a2f                 vxor.vv v20, v23, v20
  1240: 8333051d                 ld t2, 0x1d0(a0)
  1244: b3836300                 add t2, t2, t1
  1248: 878b0302                 vle8.v v23, (t2)
  1252: 5702722f                 vxor.vv v4, v23, v4
  1256: d78a7a2f                 vxor.vv v21, v23, v21
  1260: 8333851d                 ld t2, 0x1d8(a0)
  1264: b3836300                 add t2, t2, t1
  1268: 878b0302                 vle8.v v23, (t2)
  1272: d782722f                 vxor.vv v5, v23, v5
  1276: 570b7b2f                 vxor.vv v22, v23, v22
  1280: 8333051e                 ld t2, 0x1e0(a0)
  1284: b3836300                 add t2, t2, t1
  1288: 878b0302                 vle8.v v23, (t2)
  1292: 5703732f                 vxor.vv v6, v23, v6
  1296: 5706762f                 vxor.vv v12, v23, v12
  1300: 8333851e                 ld t2, 0x1e8(a0)
  1304: b3836300                 add t2, t2, t1
  1308: 878b0302                 vle8.v v23, (t2)
  1312: d783732f                 vxor.vv v7, v23, v7
  1316: d786762f                 vxor.vv v13, v23, v13
  1320: 8333051f                 ld t2, 0x1f0(a0)
  1324: b3836300                 add t2, t2, t1
  1328: 878b0302                 vle8.v v23, (t2)
  1332: 5704742f                 vxor.vv v8, v23, v8
  1336: 5707772f                 vxor.vv v14, v23, v14
  1340: d787772f                 vxor.vv v15, v23, v15
  1344: 8333851f                 ld t2, 0x1f8(a0)
  1348: b3836300                 add t2, t2, t1
  1352: 878b0302                 vle8.v v23, (t2)
  1356: d784742f                 vxor.vv v9, v23, v9
  1360: d787772f                 vxor.vv v15, v23, v15
  1364: 83330520                 ld t2, 0x200(a0)
  1368: b3836300                 add t2, t2, t1
  1372: 878b0302                 vle8.v v23, (t2)
  1376: 5705752f                 vxor.vv v10, v23, v10
  1380: 5708782f                 vxor.vv v16, v23, v16
  1384: 83338520                 ld t2, 0x208(a0)
  1388: b3836300                 add t2, t2, t1
  1392: 878b0302                 vle8.v v23, (t2)
  1396: d785752f                 vxor.vv v11, v23, v11
  1400: d788782f                 vxor.vv v17, v23, v17
  1404: 83330521                 ld t2, 0x210(a0)
  1408: b3836300                 add t2, t2, t1
  1412: 878b0302                 vle8.v v23, (t2)
  1416: d780702f                 vxor.vv v1, v23, v1
  1420: d788782f                 vxor.vv v17, v23, v17
  1424: 83338521                 ld t2, 0x218(a0)
  1428: b3836300                 add t2, t2, t1
  1432: 878b0302                 vle8.v v23, (t2)
  1436: 5701712f                 vxor.vv v2, v23, v2
  1440: 5709792f                 vxor.vv v18, v23, v18
  1444: 83330522                 ld t2, 0x220(a0)
  1448: b3836300                 add t2, t2, t1
  1452: 878b0302                 vle8.v v23, (t2)
  1456: d781712f                 vxor.vv v3, v23, v3
  1460: d789792f                 vxor.vv v19, v23, v19
  1464: 570a7a2f                 vxor.vv v20, v23, v20
  1468: 83338522                 ld t2, 0x228(a0)
  1472: b3836300                 add t2, t2, t1
  1476: 878b0302                 vle8.v v23, (t2)
  1480: 5702722f                 vxor.vv v4, v23, v4
  1484: 570a7a2f                 vxor.vv v20, v23, v20
  1488: 83330523                 ld t2, 0x230(a0)
  1492: b3836300                 add t2, t2, t1
  1496: 878b0302                 vle8.v v23, (t2)
  1500: d782722f                 vxor.vv v5, v23, v5
  1504: d78a7a2f                 vxor.vv v21, v23, v21
  1508: 83338523                 ld t2, 0x238(a0)
  1512: b3836300                 add t2, t2, t1
  1516: 878b0302                 vle8.v v23, (t2)
  1520: 5703732f                 vxor.vv v6, v23, v6
  1524: 570b7b2f                 vxor.vv v22, v23, v22
  1528: 83330524                 ld t2, 0x240(a0)
  1532: b3836300                 add t2, t2, t1
  1536: 878b0302                 vle8.v v23, (t2)
  1540: d783732f                 vxor.vv v7, v23, v7
  1544: 5706762f                 vxor.vv v12, v23, v12
  1548: 83338524                 ld t2, 0x248(a0)
  1552: b3836300                 add t2, t2, t1
  1556: 878b0302                 vle8.v v23, (t2)
  1560: 5704742f                 vxor.vv v8, v23, v8
  1564: d786762f                 vxor.vv v13, v23, v13
  1568: 83330525                 ld t2, 0x250(a0)
  1572: b3836300                 add t2, t2, t1
  1576: 878b0302                 vle8.v v23, (t2)
  1580: d784742f                 vxor.vv v9, v23, v9
  1584: 5707772f                 vxor.vv v14, v23, v14
  1588: 83338525                 ld t2, 0x258(a0)
  1592: b3836300                 add t2, t2, t1
  1596: 878b0302                 vle8.v v23, (t2)
  1600: 5705752f                 vxor.vv v10, v23, v10
  1604: d787772f                 vxor.vv v15, v23, v15
  1608: 83330526                 ld t2, 0x260(a0)
  1612: b3836300                 add t2, t2, t1
  1616: 878b0302                 vle8.v v23, (t2)
  1620: d785752f                 vxor.vv v11, v23, v11
  1624: 5708782f                 vxor.vv v16, v23, v16
  1628: 83338526                 ld t2, 0x268(a0)
  1632: b3836300                 add t2, t2, t1
  1636: 878b0302                 vle8.v v23, (t2)
  1640: d780702f                 vxor.vv v1, v23, v1
  1644: 5708782f                 vxor.vv v16, v23, v16
  1648: 83330527                 ld t2, 0x270(a0)
  1652: b3836300                 add t2, t2, t1
  1656: 878b0302                 vle8.v v23, (t2)
  1660: 5701712f                 vxor.vv v2, v23, v2
  1664: d788782f                 vxor.vv v17, v23, v17
  1668: 83338527                 ld t2, 0x278(a0)
  1672: b3836300                 add t2, t2, t1
  1676: 878b0302                 vle8.v v23, (t2)
  1680: d781712f                 vxor.vv v3, v23, v3
  1684: 5709792f                 vxor.vv v18, v23, v18
  1688: 83330528                 ld t2, 0x280(a0)
  1692: b3836300                 add t2, t2, t1
  1696: 878b0302                 vle8.v v23, (t2)
  1700: 5702722f                 vxor.vv v4, v23, v4
  1704: d789792f                 vxor.vv v19, v23, v19
  1708: 83338528                 ld t2, 0x288(a0)
  1712: b3836300                 add t2, t2, t1
  1716: 878b0302                 vle8.v v23, (t2)
  1720: d782722f                 vxor.vv v5, v23, v5
  1724: 570a7a2f                 vxor.vv v20, v23, v20
  1728: 83330529                 ld t2, 0x290(a0)
  1732: b3836300                 add t2, t2, t1
  1736: 878b0302                 vle8.v v23, (t2)
  1740: 5703732f                 vxor.vv v6, v23, v6
  1744: d78a7a2f                 vxor.vv v21, v23, v21
  1748: 83338529                 ld t2, 0x298(a0)
  1752: b3836300                 add t2, t2, t1
  1756: 878b0302                 vle8.v v23, (t2)
  1760: d783732f                 vxor.vv v7, v23, v7
  1764: 570b7b2f                 vxor.vv v22, v23, v22
  1768: 8333052a                 ld t2, 0x2a0(a0)
  1772: b3836300                 add t2, t2, t1
  1776: 878b0302                 vle8.v v23, (t2)
  1780: 5704742f                 vxor.vv v8, v23, v8
  1784: 5706762f                 vxor.vv v12, v23, v12
  1788: 8333852a                 ld t2, 0x2a8(a0)
  1792: b3836300                 add t2, t2, t1
  1796: 878b0302                 vle8.v v23, (t2)
  1800: d784742f                 vxor.vv v9, v23, v9
  1804: d786762f                 vxor.vv v13, v23, v13
  1808: 5707772f                 vxor.vv v14, v23, v14
  1812: 8333052b                 ld t2, 0x2b0(a0)
  1816: b3836300                 add t2, t2, t1
  1820: 878b0302                 vle8.v v23, (t2)
  1824: 5705752f                 vxor.vv v10, v23, v10
  1828: 5707772f                 vxor.vv v14, v23, v14
  1832: 8333852b                 ld t2, 0x2b8(a0)
  1836: b3836300                 add t2, t2, t1
  1840: 878b0302                 vle8.v v23, (t2)
  1844: d785752f                 vxor.vv v11, v23, v11
  1848: d787772f                 vxor.vv v15, v23, v15
  1852: 83b30500                 ld t2, 0(a1)
  1856: b3836300                 add t2, t2, t1
  1860: a7800302                 vse8.v v1, (t2)
  1864: 83b38500                 ld t2, 8(a1)
  1868: b3836300                 add t2, t2, t1
  1872: 27810302                 vse8.v v2, (t2)
  1876: 83b30501                 ld t2, 0x10(a1)
  1880: b3836300                 add t2, t2, t1
  1884: a7810302                 vse8.v v3, (t2)
  1888: 83b38501                 ld t2, 0x18(a1)
  1892: b3836300                 add t2, t2, t1
  1896: 27820302                 vse8.v v4, (t2)
  1900: 83b30502                 ld t2, 0x20(a1)
  1904: b3836300                 add t2, t2, t1
  1908: a7820302                 vse8.v v5, (t2)
  1912: 83b38502                 ld t2, 0x28(a1)
  1916: b3836300                 add t2, t2, t1
  1920: 27830302                 vse8.v v6, (t2)
  1924: 83b30503                 ld t2, 0x30(a1)
  1928: b3836300                 add t2, t2, t1
  1932: a7830302                 vse8.v v7, (t2)
  1936: 83b38503                 ld t2, 0x38(a1)
  1940: b3836300                 add t2, t2, t1
  1944: 27840302                 vse8.v v8, (t2)
  1948: 83b30504                 ld t2, 0x40(a1)
  1952: b3836300                 add t2, t2, t1
  1956: a7840302                 vse8.v v9, (t2)
  1960: 83b38504                 ld t2, 0x48(a1)
  1964: b3836300                 add t2, t2, t1
  1968: 27850302                 vse8.v v10, (t2)
  1972: 83b30505                 ld t2, 0x50(a1)
  1976: b3836300                 add t2, t2, t1
  1980: a7850302                 vse8.v v11, (t2)
  1984: 83b38505                 ld t2, 0x58(a1)
  1988: b3836300                 add t2, t2, t1
  1992: 27860302                 vse8.v v12, (t2)
  1996: 83b30506                 ld t2, 0x60(a1)
  2000: b3836300                 add t2, t2, t1
  2004: a7860302                 vse8.v v13, (t2)
  2008: 83b38506                 ld t2, 0x68(a1)
  2012: b3836300                 add t2, t2, t1
  2016: 27870302                 vse8.v v14, (t2)
  2020: 83b30507                 ld t2, 0x70(a1)
  2024: b3836300                 add t2, t2, t1
  2028: a7870302                 vse8.v v15, (t2)
  2032: 83b38507                 ld t2, 0x78(a1)
  2036: b3836300                 add t2, t2, t1
  2040: 27880302                 vse8.v v16, (t2)
  2044: 83b30508                 ld t2, 0x80(a1)
  2048: b3836300                 add t2, t2, t1
  2052: a7880302                 vse8.v v17, (t2)
  2056: 83b38508                 ld t2, 0x88(a1)
  2060: b3836300                 add t2, t2, t1
  2064: 27890302                 vse8.v v18, (t2)
  2068: 83b30509                 ld t2, 0x90(a1)
  2072: b3836300                 add t2, t2, t1
  2076: a7890302                 vse8.v v19, (t2)
  2080: 83b38509                 ld t2, 0x98(a1)
  2084: b3836300                 add t2, t2, t1
  2088: 278a0302                 vse8.v v20, (t2)
  2092: 83b3050a                 ld t2, 0xa0(a1)
  2096: b3836300                 add t2, t2, t1
  2100: a78a0302                 vse8.v v21, (t2)
  2104: 83b3850a                 ld t2, 0xa8(a1)
  2108: b3836300                 add t2, t2, t1
  2112: 278b0302                 vse8.v v22, (t2)
  2116: 3303d300                 add t1, t1, a3
  2120: 1307f7ff                 addi a4, a4, -1
  2124: 631a07fe                 bnez a4, 0x40
  2128: 33066640                 sub a2, a2, t1
  2132: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
  2136: 93060600                 mv a3, a2
  2140: 6344d000                 bgtz a3, 0x864
  2144: 6f001002                 j 0x1080
  2148: b3835640                 sub t2, a3, t0
  2152: 13def343                 srai t3, t2, 0x3f
  2156: b3f3c301                 and t2, t2, t3
  2160: b3837200                 add t2, t0, t2
  2164: d7ab0852                 vid.v v23
  2168: 57c0736b                 vmsltu.vx v0, v23, t2
  2172: 83330500                 ld t2, 0(a0)
  2176: b3836300                 add t2, t2, t1
  2180: 878b0300                 vle8.v v23, (t2), v0.t
  2184: d7800b5e                 vmv.v.v v1, v23
  2188: 57860b5e                 vmv.v.v v12, v23
  2192: 83338500                 ld t2, 8(a0)
  2196: b3836300                 add t2, t2, t1
  2200: 878b0300                 vle8.v v23, (t2), v0.t
  2204: 57810b5e                 vmv.v.v v2, v23
  2208: d7860b5e                 vmv.v.v v13, v23
  2212: 83330501                 ld t2, 0x10(a0)
  2216: b3836300                 add t2, t2, t1
  2220: 878b0300                 vle8.v v23, (t2), v0.t
  2224: d7810b5e                 vmv.v.v v3, v23
  2228: 57870b5e                 vmv.v.v v14, v23
  2232: 83338501                 ld t2, 0x18(a0)
  2236: b3836300                 add t2, t2, t1
  2240: 878b0300                 vle8.v v23, (t2), v0.t
  2244: 57820b5e                 vmv.v.v v4, v23
  2248: d7870b5e                 vmv.v.v v15, v23
  2252: 83330502                 ld t2, 0x20(a0)
  2256: b3836300                 add t2, t2, t1
  2260: 878b0300                 vle8.v v23, (t2), v0.t
  2264: d7820b5e                 vmv.v.v v5, v23
  2268: 57880b5e                 vmv.v.v v16, v23
  2272: 83338502                 ld t2, 0x28(a0)
  2276: b3836300                 add t2, t2, t1
  2280: 878b0300                 vle8.v v23, (t2), v0.t
  2284: 57830b5e                 vmv.v.v v6, v23
  2288: d7880b5e                 vmv.v.v v17, v23
  2292: 83330503                 ld t2, 0x30(a0)
  2296: b3836300                 add t2, t2, t1
  2300: 878b0300                 vle8.v v23, (t2), v0.t
  2304: d7830b5e                 vmv.v.v v7, v23
  2308: 57890b5e                 vmv.v.v v18, v23
  2312: 83338503                 ld t2, 0x38(a0)
  2316: b3836300                 add t2, t2, t1
  2320: 878b0300                 vle8.v v23, (t2), v0.t
  2324: 57840b5e                 vmv.v.v v8, v23
  2328: d7890b5e                 vmv.v.v v19, v23
  2332: 83330504                 ld t2, 0x40(a0)
  2336: b3836300                 add t2, t2, t1
  2340: 878b0300                 vle8.v v23, (t2), v0.t
  2344: d7840b5e                 vmv.v.v v9, v23
  2348: 578a0b5e                 vmv.v.v v20, v23
  2352: 83338504                 ld t2, 0x48(a0)
  2356: b3836300                 add t2, t2, t1
  2360: 878b0300                 vle8.v v23, (t2), v0.t
  2364: 57850b5e                 vmv.v.v v10, v23
  2368: d78a0b5e                 vmv.v.v v21, v23
  2372: 83330505                 ld t2, 0x50(a0)
  2376: b3836300                 add t2, t2, t1
  2380: 878b0300                 vle8.v v23, (t2), v0.t
  2384: d7850b5e                 vmv.v.v v11, v23
  2388: 578b0b5e                 vmv.v.v v22, v23
  2392: 83338505                 ld t2, 0x58(a0)
  2396: b3836300                 add t2, t2, t1
  2400: 878b0300                 vle8.v v23, (t2), v0.t
  2404: d780702f                 vxor.vv v1, v23, v1
  2408: 570b7b2f                 vxor.vv v22, v23, v22
  2412: 83330506                 ld t2, 0x60(a0)
  2416: b3836300                 add t2, t2, t1
  2420: 878b0300                 vle8.v v23, (t2), v0.t
  2424: 5701712f                 vxor.vv v2, v23, v2
  2428: 5706762f                 vxor.vv v12, v23, v12
  2432: 83338506                 ld t2, 0x68(a0)
  2436: b3836300                 add t2, t2, t1
  2440: 878b0300                 vle8.v v23, (t2), v0.t
  2444: d781712f                 vxor.vv v3, v23, v3
  2448: d786762f                 vxor.vv v13, v23, v13
  2452: 83330507                 ld t2, 0x70(a0)
  2456: b3836300                 add t2, t2, t1
  2460: 878b0300                 vle8.v v23, (t2), v0.t
  2464: 5702722f                 vxor.vv v4, v23, v4
  2468: 5707772f                 vxor.vv v14, v23, v14
  2472: 83338507                 ld t2, 0x78(a0)
  2476: b3836300                 add t2, t2, t1
  2480: 878b0300                 vle8.v v23, (t2), v0.t
  2484: d782722f                 vxor.vv v5, v23, v5
  2488: d787772f                 vxor.vv v15, v23, v15
  2492: 83330508                 ld t2, 0x80(a0)
  2496: b3836300                 add t2, t2, t1
  2500: 878b0300                 vle8.v v23, (t2), v0.t
  2504: 5703732f                 vxor.vv v6, v23, v6
  2508: 5708782f                 vxor.vv v16, v23, v16
  2512: d788782f                 vxor.vv v17, v23, v17
  2516: 83338508                 ld t2, 0x88(a0)
  2520: b3836300                 add t2, t2, t1
  2524: 878b0300                 vle8.v v23, (t2), v0.t
  2528: d783732f                 vxor.vv v7, v23, v7
  2532: d788782f                 vxor.vv v17, v23, v17
  2536: 83330509                 ld t2, 0x90(a0)
  2540: b3836300                 add t2, t2, t1
  2544: 878b0300                 vle8.v v23, (t2), v0.t
  2548: 5704742f                 vxor.vv v8, v23, v8
  2552: 5709792f                 vxor.vv v18, v23, v18
  2556: 83338509                 ld t2, 0x98(a0)
  2560: b3836300                 add t2, t2, t1
  2564: 878b0300                 vle8.v v23, (t2), v0.t
  2568: d784742f                 vxor.vv v9, v23, v9
  2572: d789792f                 vxor.vv v19, v23, v19
  2576: 8333050a                 ld t2, 0xa0(a0)
  2580: b3836300                 add t2, t2, t1
  2584: 878b0300                 vle8.v v23, (t2), v0.t
  2588: 5705752f                 vxor.vv v10, v23, v10
  2592: 570a7a2f                 vxor.vv v20, v23, v20
  2596: 8333850a                 ld t2, 0xa8(a0)
  2600: b3836300                 add t2, t2, t1
  2604: 878b0300                 vle8.v v23, (t2), v0.t
  2608: d785752f                 vxor.vv v11, v23, v11
  2612: d78a7a2f                 vxor.vv v21, v23, v21
  2616: 8333050b                 ld t2, 0xb0(a0)
  2620: b3836300                 add t2, t2, t1
  2624: 878b0300                 vle8.v v23, (t2), v0.t
  2628: d780702f                 vxor.vv v1, v23, v1
  2632: d78a7a2f                 vxor.vv v21, v23, v21
  2636: 570b7b2f                 vxor.vv v22, v23, v22
  2640: 8333850b                 ld t2, 0xb8(a0)
  2644: b3836300                 add t2, t2, t1
  2648: 878b0300                 vle8.v v23, (t2), v0.t
  2652: 5701712f                 vxor.vv v2, v23, v2
  2656: 570b7b2f                 vxor.vv v22, v23, v22
  2660: 8333050c                 ld t2, 0xc0(a0)
  2664: b3836300                 add t2, t2, t1
  2668: 878b0300                 vle8.v v23, (t2), v0.t
  2672: d781712f                 vxor.vv v3, v23, v3
  2676: 5706762f                 vxor.vv v12, v23, v12
  2680: 8333850c                 ld t2, 0xc8(a0)
  2684: b3836300                 add t2, t2, t1
  2688: 878b0300                 vle8.v v23, (t2), v0.t
  2692: 5702722f                 vxor.vv v4, v23, v4
  2696: d786762f                 vxor.vv v13, v23, v13
  2700: 8333050d                 ld t2, 0xd0(a0)
  2704: b3836300                 add t2, t2, t1
  2708: 878b0300                 vle8.v v23, (t2), v0.t
  2712: d782722f                 vxor.vv v5, v23, v5
  2716: 5707772f                 vxor.vv v14, v23, v14
  2720: 8333850d                 ld t2, 0xd8(a0)
  2724: b3836300                 add t2, t2, t1
  2728: 878b0300                 vle8.v v23, (t2), v0.t
  2732: 5703732f                 vxor.vv v6, v23, v6
  2736: d787772f                 vxor.vv v15, v23, v15
  2740: 8333050e                 ld t2, 0xe0(a0)
  2744: b3836300                 add t2, t2, t1
  2748: 878b0300                 vle8.v v23, (t2), v0.t
  2752: d783732f                 vxor.vv v7, v23, v7
  2756: 5708782f                 vxor.vv v16, v23, v16
  2760: 8333850e                 ld t2, 0xe8(a0)
  2764: b3836300                 add t2, t2, t1
  2768: 878b0300                 vle8.v v23, (t2), v0.t
  2772: 5704742f                 vxor.vv v8, v23, v8
  2776: d788782f                 vxor.vv v17, v23, v17
  2780: 8333050f                 ld t2, 0xf0(a0)
  2784: b3836300                 add t2, t2, t1
  2788: 878b0300                 vle8.v v23, (t2), v0.t
  2792: d784742f                 vxor.vv v9, v23, v9
  2796: 5709792f                 vxor.vv v18, v23, v18
  2800: 8333850f                 ld t2, 0xf8(a0)
  2804: b3836300                 add t2, t2, t1
  2808: 878b0300                 vle8.v v23, (t2), v0.t
  2812: 5705752f                 vxor.vv v10, v23, v10
  2816: d789792f                 vxor.vv v19, v23, v19
  2820: 83330510                 ld t2, 0x100(a0)
  2824: b3836300                 add t2, t2, t1
  2828: 878b0300                 vle8.v v23, (t2), v0.t
  2832: d785752f                 vxor.vv v11, v23, v11
  2836: 570a7a2f                 vxor.vv v20, v23, v20
  2840: 83338510                 ld t2, 0x108(a0)
  2844: b3836300                 add t2, t2, t1
  2848: 878b0300                 vle8.v v23, (t2), v0.t
  2852: d780702f                 vxor.vv v1, v23, v1
  2856: 570a7a2f                 vxor.vv v20, v23, v20
  2860: 83330511                 ld t2, 0x110(a0)
  2864: b3836300                 add t2, t2, t1
  2868: 878b0300                 vle8.v v23, (t2), v0.t
  2872: 5701712f                 vxor.vv v2, v23, v2
  2876: d78a7a2f                 vxor.vv v21, v23, v21
  2880: 83338511                 ld t2, 0x118(a0)
  2884: b3836300                 add t2, t2, t1
  2888: 878b0300                 vle8.v v23, (t2), v0.t
  2892: d781712f                 vxor.vv v3, v23, v3
  2896: 570b7b2f                 vxor.vv v22, v23, v22
  2900: 83330512                 ld t2, 0x120(a0)
  2904: b3836300                 add t2, t2, t1
  2908: 878b0300                 vle8.v v23, (t2), v0.t
  2912: 5702722f                 vxor.vv v4, v23, v4
  2916: 5706762f                 vxor.vv v12, v23, v12
  2920: 83338512                 ld t2, 0x128(a0)
  2924: b3836300                 add t2, t2, t1
  2928: 878b0300                 vle8.v v23, (t2), v0.t
  2932: d782722f                 vxor.vv v5, v23, v5
  2936: d786762f                 vxor.vv v13, v23, v13
  2940: 83330513                 ld t2, 0x130(a0)
  2944: b3836300                 add t2, t2, t1
  2948: 878b0300                 vle8.v v23, (t2), v0.t
  2952: 5703732f                 vxor.vv v6, v23, v6
  2956: 5707772f                 vxor.vv v14, v23, v14
  2960: 83338513                 ld t2, 0x138(a0)
  2964: b3836300                 add t2, t2, t1
  2968: 878b0300                 vle8.v v23, (t2), v0.t
  2972: d783732f                 vxor.vv v7, v23, v7
  2976: d787772f                 vxor.vv v15, v23, v15
  2980: 5708782f                 vxor.vv v16, v23, v16
  2984: 83330514                 ld t2, 0x140(a0)
  2988: b3836300                 add t2, t2, t1
  2992: 878b0300                 vle8.v v23, (t2), v0.t
  2996: 5704742f                 vxor.vv v8, v23, v8
  3000: 5708782f                 vxor.vv v16, v23, v16
  3004: 83338514                 ld t2, 0x148(a0)
  3008: b3836300                 add t2, t2, t1
  3012: 878b0300                 vle8.v v23, (t2), v0.t
  3016: d784742f                 vxor.vv v9, v23, v9
  3020: d788782f                 vxor.vv v17, v23, v17
  3024: 83330515                 ld t2, 0x150(a0)
  3028: b3836300                 add t2, t2, t1
  3032: 878b0300                 vle8.v v23, (t2), v0.t
  3036: 5705752f                 vxor.vv v10, v23, v10
  3040: 5709792f                 vxor.vv v18, v23, v18
  3044: 83338515                 ld t2, 0x158(a0)
  3048: b3836300                 add t2, t2, t1
  3052: 878b0300                 vle8.v v23, (t2), v0.t
  3056: d785752f                 vxor.vv v11, v23, v11
  3060: d789792f                 vxor.vv v19, v23, v19
  3064: 83330516                 ld t2, 0x160(a0)
  3068: b3836300                 add t2, t2, t1
  3072: 878b0300                 vle8.v v23, (t2), v0.t
  3076: d780702f                 vxor.vv v1, v23, v1
  3080: d789792f                 vxor.vv v19, v23, v19
  3084: 83338516                 ld t2, 0x168(a0)
  3088: b3836300                 add t2, t2, t1
  3092: 878b0300                 vle8.v v23, (t2), v0.t
  3096: 5701712f                 vxor.vv v2, v23, v2
  3100: 570a7a2f                 vxor.vv v20, v23, v20
  3104: d78a7a2f                 vxor.vv v21, v23, v21
  3108: 83330517                 ld t2, 0x170(a0)
  3112: b3836300                 add t2, t2, t1
  3116: 878b0300                 vle8.v v23, (t2), v0.t
  3120: d781712f                 vxor.vv v3, v23, v3
  3124: d78a7a2f                 vxor.vv v21, v23, v21
  3128: 83338517                 ld t2, 0x178(a0)
  3132: b3836300                 add t2, t2, t1
  3136: 878b0300                 vle8.v v23, (t2), v0.t
  3140: 5702722f                 vxor.vv v4, v23, v4
  3144: 570b7b2f                 vxor.vv v22, v23, v22
  3148: 83330518                 ld t2, 0x180(a0)
  3152: b3836300                 add t2, t2, t1
  3156: 878b0300                 vle8.v v23, (t2), v0.t
  3160: d782722f                 vxor.vv v5, v23, v5
  3164: 5706762f                 vxor.vv v12, v23, v12
  3168: 83338518                 ld t2, 0x188(a0)
  3172: b3836300                 add t2, t2, t1
  3176: 878b0300                 vle8.v v23, (t2), v0.t
  3180: 5703732f                 vxor.vv v6, v23, v6
  3184: d786762f                 vxor.vv v13, v23, v13
  3188: 83330519                 ld t2, 0x190(a0)
  3192: b3836300                 add t2, t2, t1
  3196: 878b0300                 vle8.v v23, (t2), v0.t
  3200: d783732f                 vxor.vv v7, v23, v7
  3204: 5707772f                 vxor.vv v14, v23, v14
  3208: 83338519                 ld t2, 0x198(a0)
  3212: b3836300                 add t2, t2, t1
  3216: 878b0300                 vle8.v v23, (t2), v0.t
  3220: 5704742f                 vxor.vv v8, v23, v8
  3224: d787772f                 vxor.vv v15, v23, v15
  3228: 8333051a                 ld t2, 0x1a0(a0)
  3232: b3836300                 add t2, t2, t1
  3236: 878b0300                 vle8.v v23, (t2), v0.t
  3240: d784742f                 vxor.vv v9, v23, v9
  3244: 5708782f                 vxor.vv v16, v23, v16
  3248: 8333851a                 ld t2, 0x1a8(a0)
  3252: b3836300                 add t2, t2, t1
  3256: 878b0300                 vle8.v v23, (t2), v0.t
  3260: 5705752f                 vxor.vv v10, v23, v10
  3264: d788782f                 vxor.vv v17, v23, v17
  3268: 8333051b                 ld t2, 0x1b0(a0)
  3272: b3836300                 add t2, t2, t1
  3276: 878b0300                 vle8.v v23, (t2), v0.t
  3280: d785752f                 vxor.vv v11, v23, v11
  3284: 5709792f                 vxor.vv v18, v23, v18
  3288: 8333851b                 ld t2, 0x1b8(a0)
  3292: b3836300                 add t2, t2, t1
  3296: 878b0300                 vle8.v v23, (t2), v0.t
  3300: d780702f                 vxor.vv v1, v23, v1
  3304: 5709792f                 vxor.vv v18, v23, v18
  3308: 8333051c                 ld t2, 0x1c0(a0)
  3312: b3836300                 add t2, t2, t1
  3316: 878b0300                 vle8.v v23, (t2), v0.t
  3320: 5701712f                 vxor.vv v2, v23, v2
  3324: d789792f                 vxor.vv v19, v23, v19
  3328: 8333851c                 ld t2, 0x1c8(a0)
  3332: b3836300                 add t2, t2, t1
  3336: 878b0300                 vle8.v v23, (t2), v0.t
  3340: d781712f                 vxor.vv v3, v23, v3
  3344: 570a7a2f                 vxor.vv v20, v23, v20
  3348: 8333051d                 ld t2, 0x1d0(a0)
  3352: b3836300                 add t2, t2, t1
  3356: 878b0300                 vle8.v v23, (t2), v0.t
  3360: 5702722f                 vxor.vv v4, v23, v4
  3364: d78a7a2f                 vxor.vv v21, v23, v21
  3368: 8333851d                 ld t2, 0x1d8(a0)
  3372: b3836300                 add t2, t2, t1
  3376: 878b0300                 vle8.v v23, (t2), v0.t
  3380: d782722f                 vxor.vv v5, v23, v5
  3384: 570b7b2f                 vxor.vv v22, v23, v22
  3388: 8333051e                 ld t2, 0x1e0(a0)
  3392: b3836300                 add t2, t2, t1
  3396: 878b0300                 vle8.v v23, (t2), v0.t
  3400: 5703732f                 vxor.vv v6, v23, v6
  3404: 5706762f                 vxor.vv v12, v23, v12
  3408: 8333851e                 ld t2, 0x1e8(a0)
  3412: b3836300                 add t2, t2, t1
  3416: 878b0300                 vle8.v v23, (t2), v0.t
  3420: d783732f                 vxor.vv v7, v23, v7
  3424: d786762f                 vxor.vv v13, v23, v13
  3428: 8333051f                 ld t2, 0x1f0(a0)
  3432: b3836300                 add t2, t2, t1
  3436: 878b0300                 vle8.v v23, (t2), v0.t
  3440: 5704742f                 vxor.vv v8, v23, v8
  3444: 5707772f                 vxor.vv v14, v23, v14
  3448: d787772f                 vxor.vv v15, v23, v15
  3452: 8333851f                 ld t2, 0x1f8(a0)
  3456: b3836300                 add t2, t2, t1
  3460: 878b0300                 vle8.v v23, (t2), v0.t
  3464: d784742f                 vxor.vv v9, v23, v9
  3468: d787772f                 vxor.vv v15, v23, v15
  3472: 83330520                 ld t2, 0x200(a0)
  3476: b3836300                 add t2, t2, t1
  3480: 878b0300                 vle8.v v23, (t2), v0.t
  3484: 5705752f                 vxor.vv v10, v23, v10
  3488: 5708782f                 vxor.vv v16, v23, v16
  3492: 83338520                 ld t2, 0x208(a0)
  3496: b3836300                 add t2, t2, t1
  3500: 878b0300                 vle8.v v23, (t2), v0.t
  3504: d785752f                 vxor.vv v11, v23, v11
  3508: d788782f                 vxor.vv v17, v23, v17
  3512: 83330521                 ld t2, 0x210(a0)
  3516: b3836300                 add t2, t2, t1
  3520: 878b0300                 vle8.v v23, (t2), v0.t
  3524: d780702f                 vxor.vv v1, v23, v1
  3528: d788782f                 vxor.vv v17, v23, v17
  3532: 83338521                 ld t2, 0x218(a0)
  3536: b3836300                 add t2, t2, t1
  3540: 878b0300                 vle8.v v23, (t2), v0.t
  3544: 5701712f                 vxor.vv v2, v23, v2
  3548: 5709792f                 vxor.vv v18, v23, v18
  3552: 83330522                 ld t2, 0x220(a0)
  3556: b3836300                 add t2, t2, t1
  3560: 878b0300                 vle8.v v23, (t2), v0.t
  3564: d781712f                 vxor.vv v3, v23, v3
  3568: d789792f                 vxor.vv v19, v23, v19
  3572: 570a7a2f                 vxor.vv v20, v23, v20
  3576: 83338522                 ld t2, 0x228(a0)
  3580: b3836300                 add t2, t2, t1
  3584: 878b0300                 vle8.v v23, (t2), v0.t
  3588: 5702722f                 vxor.vv v4, v23, v4
  3592: 570a7a2f                 vxor.vv v20, v23, v20
  3596: 83330523                 ld t2, 0x230(a0)
  3600: b3836300                 add t2, t2, t1
  3604: 878b0300                 vle8.v v23, (t2), v0.t
  3608: d782722f                 vxor.vv v5, v23, v5
  3612: d78a7a2f                 vxor.vv v21, v23, v21
  3616: 83338523                 ld t2, 0x238(a0)
  3620: b3836300                 add t2, t2, t1
  3624: 878b0300                 vle8.v v23, (t2), v0.t
  3628: 5703732f                 vxor.vv v6, v23, v6
  3632: 570b7b2f                 vxor.vv v22, v23, v22
  3636: 83330524                 ld t2, 0x240(a0)
  3640: b3836300                 add t2, t2, t1
  3644: 878b0300                 vle8.v v23, (t2), v0.t
  3648: d783732f                 vxor.vv v7, v23, v7
  3652: 5706762f                 vxor.vv v12, v23, v12
  3656: 83338524                 ld t2, 0x248(a0)
  3660: b3836300                 add t2, t2, t1
  3664: 878b0300                 vle8.v v23, (t2), v0.t
  3668: 5704742f                 vxor.vv v8, v23, v8
  3672: d786762f                 vxor.vv v13, v23, v13
  3676: 83330525                 ld t2, 0x250(a0)
  3680: b3836300                 add t2, t2, t1
  3684: 878b0300                 vle8.v v23, (t2), v0.t
  3688: d784742f                 vxor.vv v9, v23, v9
  3692: 5707772f                 vxor.vv v14, v23, v14
  3696: 83338525                 ld t2, 0x258(a0)
  3700: b3836300                 add t2, t2, t1
  3704: 878b0300                 vle8.v v23, (t2), v0.t
  3708: 5705752f                 vxor.vv v10, v23, v10
  3712: d787772f                 vxor.vv v15, v23, v15
  3716: 83330526                 ld t2, 0x260(a0)
  3720: b3836300                 add t2, t2, t1
  3724: 878b0300                 vle8.v v23, (t2), v0.t
  3728: d785752f                 vxor.vv v11, v23, v11
  3732: 5708782f                 vxor.vv v16, v23, v16
  3736: 83338526                 ld t2, 0x268(a0)
  3740: b3836300                 add t2, t2, t1
  3744: 878b0300                 vle8.v v23, (t2), v0.t
  3748: d780702f                 vxor.vv v1, v23, v1
  3752: 5708782f                 vxor.vv v16, v23, v16
  3756: 83330527                 ld t2, 0x270(a0)
  3760: b3836300                 add t2, t2, t1
  3764: 878b0300                 vle8.v v23, (t2), v0.t
  3768: 5701712f                 vxor.vv v2, v23, v2
  3772: d788782f                 vxor.vv v17, v23, v17
  3776: 83338527                 ld t2, 0x278(a0)
  3780: b3836300                 add t2, t2, t1
  3784: 878b0300                 vle8.v v23, (t2), v0.t
  3788: d781712f                 vxor.vv v3, v23, v3
  3792: 5709792f                 vxor.vv v18, v23, v18
  3796: 83330528                 ld t2, 0x280(a0)
  3800: b3836300                 add t2, t2, t1
  3804: 878b0300                 vle8.v v23, (t2), v0.t
  3808: 5702722f                 vxor.vv v4, v23, v4
  3812: d789792f                 vxor.vv v19, v23, v19
  3816: 83338528                 ld t2, 0x288(a0)
  3820: b3836300                 add t2, t2, t1
  3824: 878b0300                 vle8.v v23, (t2), v0.t
  3828: d782722f                 vxor.vv v5, v23, v5
  3832: 570a7a2f                 vxor.vv v20, v23, v20
  3836: 83330529                 ld t2, 0x290(a0)
  3840: b3836300                 add t2, t2, t1
  3844: 878b0300                 vle8.v v23, (t2), v0.t
  3848: 5703732f                 vxor.vv v6, v23, v6
  3852: d78a7a2f                 vxor.vv v21, v23, v21
  3856: 83338529                 ld t2, 0x298(a0)
  3860: b3836300                 add t2, t2, t1
  3864: 878b0300                 vle8.v v23, (t2), v0.t
  3868: d783732f                 vxor.vv v7, v23, v7
  3872: 570b7b2f                 vxor.vv v22, v23, v22
  3876: 8333052a                 ld t2, 0x2a0(a0)
  3880: b3836300                 add t2, t2, t1
  3884: 878b0300                 vle8.v v23, (t2), v0.t
  3888: 5704742f                 vxor.vv v8, v23, v8
  3892: 5706762f                 vxor.vv v12, v23, v12
  3896: 8333852a                 ld t2, 0x2a8(a0)
  3900: b3836300                 add t2, t2, t1
  3904: 878b0300                 vle8.v v23, (t2), v0.t
  3908: d784742f                 vxor.vv v9, v23, v9
  3912: d786762f                 vxor.vv v13, v23, v13
  3916: 5707772f                 vxor.vv v14, v23, v14
  3920: 8333052b                 ld t2, 0x2b0(a0)
  3924: b3836300                 add t2, t2, t1
  3928: 878b0300                 vle8.v v23, (t2), v0.t
  3932: 5705752f                 vxor.vv v10, v23, v10
  3936: 5707772f                 vxor.vv v14, v23, v14
  3940: 8333852b                 ld t2, 0x2b8(a0)
  3944: b3836300                 add t2, t2, t1
  3948: 878b0300                 vle8.v v23, (t2), v0.t
  3952: d785752f                 vxor.vv v11, v23, v11
  3956: d787772f                 vxor.vv v15, v23, v15
  3960: 83b30500                 ld t2, 0(a1)
  3964: b3836300                 add t2, t2, t1
  3968: a7800300                 vse8.v v1, (t2), v0.t
  3972: 83b38500                 ld t2, 8(a1)
  3976: b3836300                 add t2, t2, t1
  3980: 27810300                 vse8.v v2, (t2), v0.t
  3984: 83b30501                 ld t2, 0x10(a1)
  3988: b3836300                 add t2, t2, t1
  3992: a7810300                 vse8.v v3, (t2), v0.t
  3996: 83b38501                 ld t2, 0x18(a1)
  4000: b3836300                 add t2, t2, t1
  4004: 27820300                 vse8.v v4, (t2), v0.t
  4008: 83b30502                 ld t2, 0x20(a1)
  4012: b3836300                 add t2, t2, t1
  4016: a7820300                 vse8.v v5, (t2), v0.t
  4020: 83b38502                 ld t2, 0x28(a1)
  4024: b3836300                 add t2, t2, t1
  4028: 27830300                 vse8.v v6, (t2), v0.t
  4032: 83b30503                 ld t2, 0x30(a1)
  4036: b3836300                 add t2, t2, t1
  4040: a7830300                 vse8.v v7, (t2), v0.t
  4044: 83b38503                 ld t2, 0x38(a1)
  4048: b3836300                 add t2, t2, t1
  4052: 27840300                 vse8.v v8, (t2), v0.t
  4056: 83b30504                 ld t2, 0x40(a1)
  4060: b3836300                 add t2, t2, t1
  4064: a7840300                 vse8.v v9, (t2), v0.t
  4068: 83b38504                 ld t2, 0x48(a1)
  4072: b3836300                 add t2, t2, t1
  4076: 27850300                 vse8.v v10, (t2), v0.t
  4080: 83b30505                 ld t2, 0x50(a1)
  4084: b3836300                 add t2, t2, t1
  4088: a7850300                 vse8.v v11, (t2), v0.t
  4092: 83b38505                 ld t2, 0x58(a1)
  4096: b3836300                 add t2, t2, t1
  4100: 27860300                 vse8.v v12, (t2), v0.t
  4104: 83b30506                 ld t2, 0x60(a1)
  4108: b3836300                 add t2, t2, t1
  4112: a7860300                 vse8.v v13, (t2), v0.t
  4116: 83b38506                 ld t2, 0x68(a1)
  4120: b3836300                 add t2, t2, t1
  4124: 27870300                 vse8.v v14, (t2), v0.t
  4128: 83b30507                 ld t2, 0x70(a1)
  4132: b3836300                 add t2, t2, t1
  4136: a7870300                 vse8.v v15, (t2), v0.t
  4140: 83b38507                 ld t2, 0x78(a1)
  4144: b3836300                 add t2, t2, t1
  4148: 27880300                 vse8.v v16, (t2), v0.t
  4152: 83b30508                 ld t2, 0x80(a1)
  4156: b3836300                 add t2, t2, t1
  4160: a7880300                 vse8.v v17, (t2), v0.t
  4164: 83b38508                 ld t2, 0x88(a1)
  4168: b3836300                 add t2, t2, t1
  4172: 27890300                 vse8.v v18, (t2), v0.t
  4176: 83b30509                 ld t2, 0x90(a1)
  4180: b3836300                 add t2, t2, t1
  4184: a7890300                 vse8.v v19, (t2), v0.t
  4188: 83b38509                 ld t2, 0x98(a1)
  4192: b3836300                 add t2, t2, t1
  4196: 278a0300                 vse8.v v20, (t2), v0.t
  4200: 83b3050a                 ld t2, 0xa0(a1)
  4204: b3836300                 add t2, t2, t1
  4208: a78a0300                 vse8.v v21, (t2), v0.t
  4212: 83b3850a                 ld t2, 0xa8(a1)
  4216: b3836300                 add t2, t2, t1
  4220: 278b0300                 vse8.v v22, (t2), v0.t
  4224: 33035300                 add t1, t1, t0
  4228: b3865640                 sub a3, a3, t0
  4232: 67800000                 ret 
  4236: 0000                     unimp 
  4238: 0000                     unimp 
