; config_t1_1850 target=rvv-zvbc32e twin=notail topo=liberation(8)/enc op=encode U=1
; code_size=2144 twins=["notail"] uops_per_byte=18.375
     0: 63140600                 bnez a2, 8
     4: 6f00d004                 j 0x850
     8: 971f0000                 auipc t6, 1
    12: 938f8f85                 addi t6, t6, -0x7a8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f009002                 j 0x850
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 078b0302                 vle8.v v22, (t2)
    76: 57000b5e                 vmv.v.v v0, v22
    80: d7050b5e                 vmv.v.v v11, v22
    84: 83338500                 ld t2, 8(a0)
    88: b3836300                 add t2, t2, t1
    92: 078b0302                 vle8.v v22, (t2)
    96: d7000b5e                 vmv.v.v v1, v22
   100: 57060b5e                 vmv.v.v v12, v22
   104: 83330501                 ld t2, 0x10(a0)
   108: b3836300                 add t2, t2, t1
   112: 078b0302                 vle8.v v22, (t2)
   116: 57010b5e                 vmv.v.v v2, v22
   120: d7060b5e                 vmv.v.v v13, v22
   124: 83338501                 ld t2, 0x18(a0)
   128: b3836300                 add t2, t2, t1
   132: 078b0302                 vle8.v v22, (t2)
   136: d7010b5e                 vmv.v.v v3, v22
   140: 57070b5e                 vmv.v.v v14, v22
   144: 83330502                 ld t2, 0x20(a0)
   148: b3836300                 add t2, t2, t1
   152: 078b0302                 vle8.v v22, (t2)
   156: 57020b5e                 vmv.v.v v4, v22
   160: d7070b5e                 vmv.v.v v15, v22
   164: 83338502                 ld t2, 0x28(a0)
   168: b3836300                 add t2, t2, t1
   172: 078b0302                 vle8.v v22, (t2)
   176: d7020b5e                 vmv.v.v v5, v22
   180: 57080b5e                 vmv.v.v v16, v22
   184: 83330503                 ld t2, 0x30(a0)
   188: b3836300                 add t2, t2, t1
   192: 078b0302                 vle8.v v22, (t2)
   196: 57030b5e                 vmv.v.v v6, v22
   200: d7080b5e                 vmv.v.v v17, v22
   204: 83338503                 ld t2, 0x38(a0)
   208: b3836300                 add t2, t2, t1
   212: 078b0302                 vle8.v v22, (t2)
   216: d7030b5e                 vmv.v.v v7, v22
   220: 57090b5e                 vmv.v.v v18, v22
   224: 83330504                 ld t2, 0x40(a0)
   228: b3836300                 add t2, t2, t1
   232: 078b0302                 vle8.v v22, (t2)
   236: 57040b5e                 vmv.v.v v8, v22
   240: d7090b5e                 vmv.v.v v19, v22
   244: 83338504                 ld t2, 0x48(a0)
   248: b3836300                 add t2, t2, t1
   252: 078b0302                 vle8.v v22, (t2)
   256: d7040b5e                 vmv.v.v v9, v22
   260: 570a0b5e                 vmv.v.v v20, v22
   264: 83330505                 ld t2, 0x50(a0)
   268: b3836300                 add t2, t2, t1
   272: 078b0302                 vle8.v v22, (t2)
   276: 57050b5e                 vmv.v.v v10, v22
   280: d70a0b5e                 vmv.v.v v21, v22
   284: 83338505                 ld t2, 0x58(a0)
   288: b3836300                 add t2, t2, t1
   292: 078b0302                 vle8.v v22, (t2)
   296: 5700602f                 vxor.vv v0, v22, v0
   300: d78a6a2f                 vxor.vv v21, v22, v21
   304: 83330506                 ld t2, 0x60(a0)
   308: b3836300                 add t2, t2, t1
   312: 078b0302                 vle8.v v22, (t2)
   316: d780602f                 vxor.vv v1, v22, v1
   320: d785652f                 vxor.vv v11, v22, v11
   324: 83338506                 ld t2, 0x68(a0)
   328: b3836300                 add t2, t2, t1
   332: 078b0302                 vle8.v v22, (t2)
   336: 5701612f                 vxor.vv v2, v22, v2
   340: 5706662f                 vxor.vv v12, v22, v12
   344: 83330507                 ld t2, 0x70(a0)
   348: b3836300                 add t2, t2, t1
   352: 078b0302                 vle8.v v22, (t2)
   356: d781612f                 vxor.vv v3, v22, v3
   360: d786662f                 vxor.vv v13, v22, v13
   364: 83338507                 ld t2, 0x78(a0)
   368: b3836300                 add t2, t2, t1
   372: 078b0302                 vle8.v v22, (t2)
   376: 5702622f                 vxor.vv v4, v22, v4
   380: 5707672f                 vxor.vv v14, v22, v14
   384: 83330508                 ld t2, 0x80(a0)
   388: b3836300                 add t2, t2, t1
   392: 078b0302                 vle8.v v22, (t2)
   396: d782622f                 vxor.vv v5, v22, v5
   400: d787672f                 vxor.vv v15, v22, v15
   404: 5708682f                 vxor.vv v16, v22, v16
   408: 83338508                 ld t2, 0x88(a0)
   412: b3836300                 add t2, t2, t1
   416: 078b0302                 vle8.v v22, (t2)
   420: 5703632f                 vxor.vv v6, v22, v6
   424: 5708682f                 vxor.vv v16, v22, v16
   428: 83330509                 ld t2, 0x90(a0)
   432: b3836300                 add t2, t2, t1
   436: 078b0302                 vle8.v v22, (t2)
   440: d783632f                 vxor.vv v7, v22, v7
   444: d788682f                 vxor.vv v17, v22, v17
   448: 83338509                 ld t2, 0x98(a0)
   452: b3836300                 add t2, t2, t1
   456: 078b0302                 vle8.v v22, (t2)
   460: 5704642f                 vxor.vv v8, v22, v8
   464: 5709692f                 vxor.vv v18, v22, v18
   468: 8333050a                 ld t2, 0xa0(a0)
   472: b3836300                 add t2, t2, t1
   476: 078b0302                 vle8.v v22, (t2)
   480: d784642f                 vxor.vv v9, v22, v9
   484: d789692f                 vxor.vv v19, v22, v19
   488: 8333850a                 ld t2, 0xa8(a0)
   492: b3836300                 add t2, t2, t1
   496: 078b0302                 vle8.v v22, (t2)
   500: 5705652f                 vxor.vv v10, v22, v10
   504: 570a6a2f                 vxor.vv v20, v22, v20
   508: 8333050b                 ld t2, 0xb0(a0)
   512: b3836300                 add t2, t2, t1
   516: 078b0302                 vle8.v v22, (t2)
   520: 5700602f                 vxor.vv v0, v22, v0
   524: 570a6a2f                 vxor.vv v20, v22, v20
   528: d78a6a2f                 vxor.vv v21, v22, v21
   532: 8333850b                 ld t2, 0xb8(a0)
   536: b3836300                 add t2, t2, t1
   540: 078b0302                 vle8.v v22, (t2)
   544: d780602f                 vxor.vv v1, v22, v1
   548: d78a6a2f                 vxor.vv v21, v22, v21
   552: 8333050c                 ld t2, 0xc0(a0)
   556: b3836300                 add t2, t2, t1
   560: 078b0302                 vle8.v v22, (t2)
   564: 5701612f                 vxor.vv v2, v22, v2
   568: d785652f                 vxor.vv v11, v22, v11
   572: 8333850c                 ld t2, 0xc8(a0)
   576: b3836300                 add t2, t2, t1
   580: 078b0302                 vle8.v v22, (t2)
   584: d781612f                 vxor.vv v3, v22, v3
   588: 5706662f                 vxor.vv v12, v22, v12
   592: 8333050d                 ld t2, 0xd0(a0)
   596: b3836300                 add t2, t2, t1
   600: 078b0302                 vle8.v v22, (t2)
   604: 5702622f                 vxor.vv v4, v22, v4
   608: d786662f                 vxor.vv v13, v22, v13
   612: 8333850d                 ld t2, 0xd8(a0)
   616: b3836300                 add t2, t2, t1
   620: 078b0302                 vle8.v v22, (t2)
   624: d782622f                 vxor.vv v5, v22, v5
   628: 5707672f                 vxor.vv v14, v22, v14
   632: 8333050e                 ld t2, 0xe0(a0)
   636: b3836300                 add t2, t2, t1
   640: 078b0302                 vle8.v v22, (t2)
   644: 5703632f                 vxor.vv v6, v22, v6
   648: d787672f                 vxor.vv v15, v22, v15
   652: 8333850e                 ld t2, 0xe8(a0)
   656: b3836300                 add t2, t2, t1
   660: 078b0302                 vle8.v v22, (t2)
   664: d783632f                 vxor.vv v7, v22, v7
   668: 5708682f                 vxor.vv v16, v22, v16
   672: 8333050f                 ld t2, 0xf0(a0)
   676: b3836300                 add t2, t2, t1
   680: 078b0302                 vle8.v v22, (t2)
   684: 5704642f                 vxor.vv v8, v22, v8
   688: d788682f                 vxor.vv v17, v22, v17
   692: 8333850f                 ld t2, 0xf8(a0)
   696: b3836300                 add t2, t2, t1
   700: 078b0302                 vle8.v v22, (t2)
   704: d784642f                 vxor.vv v9, v22, v9
   708: 5709692f                 vxor.vv v18, v22, v18
   712: 83330510                 ld t2, 0x100(a0)
   716: b3836300                 add t2, t2, t1
   720: 078b0302                 vle8.v v22, (t2)
   724: 5705652f                 vxor.vv v10, v22, v10
   728: d789692f                 vxor.vv v19, v22, v19
   732: 83338510                 ld t2, 0x108(a0)
   736: b3836300                 add t2, t2, t1
   740: 078b0302                 vle8.v v22, (t2)
   744: 5700602f                 vxor.vv v0, v22, v0
   748: d789692f                 vxor.vv v19, v22, v19
   752: 83330511                 ld t2, 0x110(a0)
   756: b3836300                 add t2, t2, t1
   760: 078b0302                 vle8.v v22, (t2)
   764: d780602f                 vxor.vv v1, v22, v1
   768: 570a6a2f                 vxor.vv v20, v22, v20
   772: 83338511                 ld t2, 0x118(a0)
   776: b3836300                 add t2, t2, t1
   780: 078b0302                 vle8.v v22, (t2)
   784: 5701612f                 vxor.vv v2, v22, v2
   788: d78a6a2f                 vxor.vv v21, v22, v21
   792: 83330512                 ld t2, 0x120(a0)
   796: b3836300                 add t2, t2, t1
   800: 078b0302                 vle8.v v22, (t2)
   804: d781612f                 vxor.vv v3, v22, v3
   808: d785652f                 vxor.vv v11, v22, v11
   812: 83338512                 ld t2, 0x128(a0)
   816: b3836300                 add t2, t2, t1
   820: 078b0302                 vle8.v v22, (t2)
   824: 5702622f                 vxor.vv v4, v22, v4
   828: 5706662f                 vxor.vv v12, v22, v12
   832: 83330513                 ld t2, 0x130(a0)
   836: b3836300                 add t2, t2, t1
   840: 078b0302                 vle8.v v22, (t2)
   844: d782622f                 vxor.vv v5, v22, v5
   848: d786662f                 vxor.vv v13, v22, v13
   852: 83338513                 ld t2, 0x138(a0)
   856: b3836300                 add t2, t2, t1
   860: 078b0302                 vle8.v v22, (t2)
   864: 5703632f                 vxor.vv v6, v22, v6
   868: 5707672f                 vxor.vv v14, v22, v14
   872: d787672f                 vxor.vv v15, v22, v15
   876: 83330514                 ld t2, 0x140(a0)
   880: b3836300                 add t2, t2, t1
   884: 078b0302                 vle8.v v22, (t2)
   888: d783632f                 vxor.vv v7, v22, v7
   892: d787672f                 vxor.vv v15, v22, v15
   896: 83338514                 ld t2, 0x148(a0)
   900: b3836300                 add t2, t2, t1
   904: 078b0302                 vle8.v v22, (t2)
   908: 5704642f                 vxor.vv v8, v22, v8
   912: 5708682f                 vxor.vv v16, v22, v16
   916: 83330515                 ld t2, 0x150(a0)
   920: b3836300                 add t2, t2, t1
   924: 078b0302                 vle8.v v22, (t2)
   928: d784642f                 vxor.vv v9, v22, v9
   932: d788682f                 vxor.vv v17, v22, v17
   936: 83338515                 ld t2, 0x158(a0)
   940: b3836300                 add t2, t2, t1
   944: 078b0302                 vle8.v v22, (t2)
   948: 5705652f                 vxor.vv v10, v22, v10
   952: 5709692f                 vxor.vv v18, v22, v18
   956: 83330516                 ld t2, 0x160(a0)
   960: b3836300                 add t2, t2, t1
   964: 078b0302                 vle8.v v22, (t2)
   968: 5700602f                 vxor.vv v0, v22, v0
   972: 5709692f                 vxor.vv v18, v22, v18
   976: 83338516                 ld t2, 0x168(a0)
   980: b3836300                 add t2, t2, t1
   984: 078b0302                 vle8.v v22, (t2)
   988: d780602f                 vxor.vv v1, v22, v1
   992: d789692f                 vxor.vv v19, v22, v19
   996: 570a6a2f                 vxor.vv v20, v22, v20
  1000: 83330517                 ld t2, 0x170(a0)
  1004: b3836300                 add t2, t2, t1
  1008: 078b0302                 vle8.v v22, (t2)
  1012: 5701612f                 vxor.vv v2, v22, v2
  1016: 570a6a2f                 vxor.vv v20, v22, v20
  1020: 83338517                 ld t2, 0x178(a0)
  1024: b3836300                 add t2, t2, t1
  1028: 078b0302                 vle8.v v22, (t2)
  1032: d781612f                 vxor.vv v3, v22, v3
  1036: d78a6a2f                 vxor.vv v21, v22, v21
  1040: 83330518                 ld t2, 0x180(a0)
  1044: b3836300                 add t2, t2, t1
  1048: 078b0302                 vle8.v v22, (t2)
  1052: 5702622f                 vxor.vv v4, v22, v4
  1056: d785652f                 vxor.vv v11, v22, v11
  1060: 83338518                 ld t2, 0x188(a0)
  1064: b3836300                 add t2, t2, t1
  1068: 078b0302                 vle8.v v22, (t2)
  1072: d782622f                 vxor.vv v5, v22, v5
  1076: 5706662f                 vxor.vv v12, v22, v12
  1080: 83330519                 ld t2, 0x190(a0)
  1084: b3836300                 add t2, t2, t1
  1088: 078b0302                 vle8.v v22, (t2)
  1092: 5703632f                 vxor.vv v6, v22, v6
  1096: d786662f                 vxor.vv v13, v22, v13
  1100: 83338519                 ld t2, 0x198(a0)
  1104: b3836300                 add t2, t2, t1
  1108: 078b0302                 vle8.v v22, (t2)
  1112: d783632f                 vxor.vv v7, v22, v7
  1116: 5707672f                 vxor.vv v14, v22, v14
  1120: 8333051a                 ld t2, 0x1a0(a0)
  1124: b3836300                 add t2, t2, t1
  1128: 078b0302                 vle8.v v22, (t2)
  1132: 5704642f                 vxor.vv v8, v22, v8
  1136: d787672f                 vxor.vv v15, v22, v15
  1140: 8333851a                 ld t2, 0x1a8(a0)
  1144: b3836300                 add t2, t2, t1
  1148: 078b0302                 vle8.v v22, (t2)
  1152: d784642f                 vxor.vv v9, v22, v9
  1156: 5708682f                 vxor.vv v16, v22, v16
  1160: 8333051b                 ld t2, 0x1b0(a0)
  1164: b3836300                 add t2, t2, t1
  1168: 078b0302                 vle8.v v22, (t2)
  1172: 5705652f                 vxor.vv v10, v22, v10
  1176: d788682f                 vxor.vv v17, v22, v17
  1180: 8333851b                 ld t2, 0x1b8(a0)
  1184: b3836300                 add t2, t2, t1
  1188: 078b0302                 vle8.v v22, (t2)
  1192: 5700602f                 vxor.vv v0, v22, v0
  1196: d788682f                 vxor.vv v17, v22, v17
  1200: 8333051c                 ld t2, 0x1c0(a0)
  1204: b3836300                 add t2, t2, t1
  1208: 078b0302                 vle8.v v22, (t2)
  1212: d780602f                 vxor.vv v1, v22, v1
  1216: 5709692f                 vxor.vv v18, v22, v18
  1220: 8333851c                 ld t2, 0x1c8(a0)
  1224: b3836300                 add t2, t2, t1
  1228: 078b0302                 vle8.v v22, (t2)
  1232: 5701612f                 vxor.vv v2, v22, v2
  1236: d789692f                 vxor.vv v19, v22, v19
  1240: 8333051d                 ld t2, 0x1d0(a0)
  1244: b3836300                 add t2, t2, t1
  1248: 078b0302                 vle8.v v22, (t2)
  1252: d781612f                 vxor.vv v3, v22, v3
  1256: 570a6a2f                 vxor.vv v20, v22, v20
  1260: 8333851d                 ld t2, 0x1d8(a0)
  1264: b3836300                 add t2, t2, t1
  1268: 078b0302                 vle8.v v22, (t2)
  1272: 5702622f                 vxor.vv v4, v22, v4
  1276: d78a6a2f                 vxor.vv v21, v22, v21
  1280: 8333051e                 ld t2, 0x1e0(a0)
  1284: b3836300                 add t2, t2, t1
  1288: 078b0302                 vle8.v v22, (t2)
  1292: d782622f                 vxor.vv v5, v22, v5
  1296: d785652f                 vxor.vv v11, v22, v11
  1300: 8333851e                 ld t2, 0x1e8(a0)
  1304: b3836300                 add t2, t2, t1
  1308: 078b0302                 vle8.v v22, (t2)
  1312: 5703632f                 vxor.vv v6, v22, v6
  1316: 5706662f                 vxor.vv v12, v22, v12
  1320: 8333051f                 ld t2, 0x1f0(a0)
  1324: b3836300                 add t2, t2, t1
  1328: 078b0302                 vle8.v v22, (t2)
  1332: d783632f                 vxor.vv v7, v22, v7
  1336: d786662f                 vxor.vv v13, v22, v13
  1340: 5707672f                 vxor.vv v14, v22, v14
  1344: 8333851f                 ld t2, 0x1f8(a0)
  1348: b3836300                 add t2, t2, t1
  1352: 078b0302                 vle8.v v22, (t2)
  1356: 5704642f                 vxor.vv v8, v22, v8
  1360: 5707672f                 vxor.vv v14, v22, v14
  1364: 83330520                 ld t2, 0x200(a0)
  1368: b3836300                 add t2, t2, t1
  1372: 078b0302                 vle8.v v22, (t2)
  1376: d784642f                 vxor.vv v9, v22, v9
  1380: d787672f                 vxor.vv v15, v22, v15
  1384: 83338520                 ld t2, 0x208(a0)
  1388: b3836300                 add t2, t2, t1
  1392: 078b0302                 vle8.v v22, (t2)
  1396: 5705652f                 vxor.vv v10, v22, v10
  1400: 5708682f                 vxor.vv v16, v22, v16
  1404: 83330521                 ld t2, 0x210(a0)
  1408: b3836300                 add t2, t2, t1
  1412: 078b0302                 vle8.v v22, (t2)
  1416: 5700602f                 vxor.vv v0, v22, v0
  1420: 5708682f                 vxor.vv v16, v22, v16
  1424: 83338521                 ld t2, 0x218(a0)
  1428: b3836300                 add t2, t2, t1
  1432: 078b0302                 vle8.v v22, (t2)
  1436: d780602f                 vxor.vv v1, v22, v1
  1440: d788682f                 vxor.vv v17, v22, v17
  1444: 83330522                 ld t2, 0x220(a0)
  1448: b3836300                 add t2, t2, t1
  1452: 078b0302                 vle8.v v22, (t2)
  1456: 5701612f                 vxor.vv v2, v22, v2
  1460: 5709692f                 vxor.vv v18, v22, v18
  1464: d789692f                 vxor.vv v19, v22, v19
  1468: 83338522                 ld t2, 0x228(a0)
  1472: b3836300                 add t2, t2, t1
  1476: 078b0302                 vle8.v v22, (t2)
  1480: d781612f                 vxor.vv v3, v22, v3
  1484: d789692f                 vxor.vv v19, v22, v19
  1488: 83330523                 ld t2, 0x230(a0)
  1492: b3836300                 add t2, t2, t1
  1496: 078b0302                 vle8.v v22, (t2)
  1500: 5702622f                 vxor.vv v4, v22, v4
  1504: 570a6a2f                 vxor.vv v20, v22, v20
  1508: 83338523                 ld t2, 0x238(a0)
  1512: b3836300                 add t2, t2, t1
  1516: 078b0302                 vle8.v v22, (t2)
  1520: d782622f                 vxor.vv v5, v22, v5
  1524: d78a6a2f                 vxor.vv v21, v22, v21
  1528: 83330524                 ld t2, 0x240(a0)
  1532: b3836300                 add t2, t2, t1
  1536: 078b0302                 vle8.v v22, (t2)
  1540: 5703632f                 vxor.vv v6, v22, v6
  1544: d785652f                 vxor.vv v11, v22, v11
  1548: 83338524                 ld t2, 0x248(a0)
  1552: b3836300                 add t2, t2, t1
  1556: 078b0302                 vle8.v v22, (t2)
  1560: d783632f                 vxor.vv v7, v22, v7
  1564: 5706662f                 vxor.vv v12, v22, v12
  1568: 83330525                 ld t2, 0x250(a0)
  1572: b3836300                 add t2, t2, t1
  1576: 078b0302                 vle8.v v22, (t2)
  1580: 5704642f                 vxor.vv v8, v22, v8
  1584: d786662f                 vxor.vv v13, v22, v13
  1588: 83338525                 ld t2, 0x258(a0)
  1592: b3836300                 add t2, t2, t1
  1596: 078b0302                 vle8.v v22, (t2)
  1600: d784642f                 vxor.vv v9, v22, v9
  1604: 5707672f                 vxor.vv v14, v22, v14
  1608: 83330526                 ld t2, 0x260(a0)
  1612: b3836300                 add t2, t2, t1
  1616: 078b0302                 vle8.v v22, (t2)
  1620: 5705652f                 vxor.vv v10, v22, v10
  1624: d787672f                 vxor.vv v15, v22, v15
  1628: 83338526                 ld t2, 0x268(a0)
  1632: b3836300                 add t2, t2, t1
  1636: 078b0302                 vle8.v v22, (t2)
  1640: 5700602f                 vxor.vv v0, v22, v0
  1644: d787672f                 vxor.vv v15, v22, v15
  1648: 83330527                 ld t2, 0x270(a0)
  1652: b3836300                 add t2, t2, t1
  1656: 078b0302                 vle8.v v22, (t2)
  1660: d780602f                 vxor.vv v1, v22, v1
  1664: 5708682f                 vxor.vv v16, v22, v16
  1668: 83338527                 ld t2, 0x278(a0)
  1672: b3836300                 add t2, t2, t1
  1676: 078b0302                 vle8.v v22, (t2)
  1680: 5701612f                 vxor.vv v2, v22, v2
  1684: d788682f                 vxor.vv v17, v22, v17
  1688: 83330528                 ld t2, 0x280(a0)
  1692: b3836300                 add t2, t2, t1
  1696: 078b0302                 vle8.v v22, (t2)
  1700: d781612f                 vxor.vv v3, v22, v3
  1704: 5709692f                 vxor.vv v18, v22, v18
  1708: 83338528                 ld t2, 0x288(a0)
  1712: b3836300                 add t2, t2, t1
  1716: 078b0302                 vle8.v v22, (t2)
  1720: 5702622f                 vxor.vv v4, v22, v4
  1724: d789692f                 vxor.vv v19, v22, v19
  1728: 83330529                 ld t2, 0x290(a0)
  1732: b3836300                 add t2, t2, t1
  1736: 078b0302                 vle8.v v22, (t2)
  1740: d782622f                 vxor.vv v5, v22, v5
  1744: 570a6a2f                 vxor.vv v20, v22, v20
  1748: 83338529                 ld t2, 0x298(a0)
  1752: b3836300                 add t2, t2, t1
  1756: 078b0302                 vle8.v v22, (t2)
  1760: 5703632f                 vxor.vv v6, v22, v6
  1764: d78a6a2f                 vxor.vv v21, v22, v21
  1768: 8333052a                 ld t2, 0x2a0(a0)
  1772: b3836300                 add t2, t2, t1
  1776: 078b0302                 vle8.v v22, (t2)
  1780: d783632f                 vxor.vv v7, v22, v7
  1784: d785652f                 vxor.vv v11, v22, v11
  1788: 8333852a                 ld t2, 0x2a8(a0)
  1792: b3836300                 add t2, t2, t1
  1796: 078b0302                 vle8.v v22, (t2)
  1800: 5704642f                 vxor.vv v8, v22, v8
  1804: 5706662f                 vxor.vv v12, v22, v12
  1808: d786662f                 vxor.vv v13, v22, v13
  1812: 8333052b                 ld t2, 0x2b0(a0)
  1816: b3836300                 add t2, t2, t1
  1820: 078b0302                 vle8.v v22, (t2)
  1824: d784642f                 vxor.vv v9, v22, v9
  1828: d786662f                 vxor.vv v13, v22, v13
  1832: 8333852b                 ld t2, 0x2b8(a0)
  1836: b3836300                 add t2, t2, t1
  1840: 078b0302                 vle8.v v22, (t2)
  1844: 5705652f                 vxor.vv v10, v22, v10
  1848: 5707672f                 vxor.vv v14, v22, v14
  1852: 83b30500                 ld t2, 0(a1)
  1856: b3836300                 add t2, t2, t1
  1860: 27800302                 vse8.v v0, (t2)
  1864: 83b38500                 ld t2, 8(a1)
  1868: b3836300                 add t2, t2, t1
  1872: a7800302                 vse8.v v1, (t2)
  1876: 83b30501                 ld t2, 0x10(a1)
  1880: b3836300                 add t2, t2, t1
  1884: 27810302                 vse8.v v2, (t2)
  1888: 83b38501                 ld t2, 0x18(a1)
  1892: b3836300                 add t2, t2, t1
  1896: a7810302                 vse8.v v3, (t2)
  1900: 83b30502                 ld t2, 0x20(a1)
  1904: b3836300                 add t2, t2, t1
  1908: 27820302                 vse8.v v4, (t2)
  1912: 83b38502                 ld t2, 0x28(a1)
  1916: b3836300                 add t2, t2, t1
  1920: a7820302                 vse8.v v5, (t2)
  1924: 83b30503                 ld t2, 0x30(a1)
  1928: b3836300                 add t2, t2, t1
  1932: 27830302                 vse8.v v6, (t2)
  1936: 83b38503                 ld t2, 0x38(a1)
  1940: b3836300                 add t2, t2, t1
  1944: a7830302                 vse8.v v7, (t2)
  1948: 83b30504                 ld t2, 0x40(a1)
  1952: b3836300                 add t2, t2, t1
  1956: 27840302                 vse8.v v8, (t2)
  1960: 83b38504                 ld t2, 0x48(a1)
  1964: b3836300                 add t2, t2, t1
  1968: a7840302                 vse8.v v9, (t2)
  1972: 83b30505                 ld t2, 0x50(a1)
  1976: b3836300                 add t2, t2, t1
  1980: 27850302                 vse8.v v10, (t2)
  1984: 83b38505                 ld t2, 0x58(a1)
  1988: b3836300                 add t2, t2, t1
  1992: a7850302                 vse8.v v11, (t2)
  1996: 83b30506                 ld t2, 0x60(a1)
  2000: b3836300                 add t2, t2, t1
  2004: 27860302                 vse8.v v12, (t2)
  2008: 83b38506                 ld t2, 0x68(a1)
  2012: b3836300                 add t2, t2, t1
  2016: a7860302                 vse8.v v13, (t2)
  2020: 83b30507                 ld t2, 0x70(a1)
  2024: b3836300                 add t2, t2, t1
  2028: 27870302                 vse8.v v14, (t2)
  2032: 83b38507                 ld t2, 0x78(a1)
  2036: b3836300                 add t2, t2, t1
  2040: a7870302                 vse8.v v15, (t2)
  2044: 83b30508                 ld t2, 0x80(a1)
  2048: b3836300                 add t2, t2, t1
  2052: 27880302                 vse8.v v16, (t2)
  2056: 83b38508                 ld t2, 0x88(a1)
  2060: b3836300                 add t2, t2, t1
  2064: a7880302                 vse8.v v17, (t2)
  2068: 83b30509                 ld t2, 0x90(a1)
  2072: b3836300                 add t2, t2, t1
  2076: 27890302                 vse8.v v18, (t2)
  2080: 83b38509                 ld t2, 0x98(a1)
  2084: b3836300                 add t2, t2, t1
  2088: a7890302                 vse8.v v19, (t2)
  2092: 83b3050a                 ld t2, 0xa0(a1)
  2096: b3836300                 add t2, t2, t1
  2100: 278a0302                 vse8.v v20, (t2)
  2104: 83b3850a                 ld t2, 0xa8(a1)
  2108: b3836300                 add t2, t2, t1
  2112: a78a0302                 vse8.v v21, (t2)
  2116: 3303d300                 add t1, t1, a3
  2120: 1307f7ff                 addi a4, a4, -1
  2124: 631a07fe                 bnez a4, 0x40
  2128: 67800000                 ret 
  2132: 0000                     unimp 
  2134: 0000                     unimp 
  2136: 0000                     unimp 
  2138: 0000                     unimp 
  2140: 0000                     unimp 
  2142: 0000                     unimp 
