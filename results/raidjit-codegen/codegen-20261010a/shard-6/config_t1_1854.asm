; config_t1_1854 target=rvv-zvbc32e twin=xorsched topo=liberation(8)/enc op=encode U=1
; code_size=4336 twins=["xorsched"] uops_per_byte=36.75
     0: 63140600                 bnez a2, 8
     4: 6f10c00d                 j 0x10e0
     8: 971f0000                 auipc t6, 1
    12: 938f8f0e                 addi t6, t6, 0xe8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f005006                 j 0x88c
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 033e8502                 ld t3, 0x28(a0)
    68: 330e6e00                 add t3, t3, t1
    72: 833e0508                 ld t4, 0x80(a0)
    76: b38e6e00                 add t4, t4, t1
    80: 070b0e02                 vle8.v v22, (t3)
    84: 878b0e02                 vle8.v v23, (t4)
    88: 570b7b2f                 vxor.vv v22, v23, v22
    92: d7020b5e                 vmv.v.v v5, v22
    96: 57080b5e                 vmv.v.v v16, v22
   100: 033e8505                 ld t3, 0x58(a0)
   104: 330e6e00                 add t3, t3, t1
   108: 833e050b                 ld t4, 0xb0(a0)
   112: b38e6e00                 add t4, t4, t1
   116: 070b0e02                 vle8.v v22, (t3)
   120: 878b0e02                 vle8.v v23, (t4)
   124: 570b7b2f                 vxor.vv v22, v23, v22
   128: 57000b5e                 vmv.v.v v0, v22
   132: d70a0b5e                 vmv.v.v v21, v22
   136: 033e050e                 ld t3, 0xe0(a0)
   140: 330e6e00                 add t3, t3, t1
   144: 833e8513                 ld t4, 0x138(a0)
   148: b38e6e00                 add t4, t4, t1
   152: 070b0e02                 vle8.v v22, (t3)
   156: 878b0e02                 vle8.v v23, (t4)
   160: 570b7b2f                 vxor.vv v22, v23, v22
   164: 57030b5e                 vmv.v.v v6, v22
   168: d7070b5e                 vmv.v.v v15, v22
   172: 033e0511                 ld t3, 0x110(a0)
   176: 330e6e00                 add t3, t3, t1
   180: 833e8516                 ld t4, 0x168(a0)
   184: b38e6e00                 add t4, t4, t1
   188: 070b0e02                 vle8.v v22, (t3)
   192: 878b0e02                 vle8.v v23, (t4)
   196: 570b7b2f                 vxor.vv v22, v23, v22
   200: d7000b5e                 vmv.v.v v1, v22
   204: 570a0b5e                 vmv.v.v v20, v22
   208: 033e8519                 ld t3, 0x198(a0)
   212: 330e6e00                 add t3, t3, t1
   216: 833e051f                 ld t4, 0x1f0(a0)
   220: b38e6e00                 add t4, t4, t1
   224: 070b0e02                 vle8.v v22, (t3)
   228: 878b0e02                 vle8.v v23, (t4)
   232: 570b7b2f                 vxor.vv v22, v23, v22
   236: d7030b5e                 vmv.v.v v7, v22
   240: 57070b5e                 vmv.v.v v14, v22
   244: 033e851c                 ld t3, 0x1c8(a0)
   248: 330e6e00                 add t3, t3, t1
   252: 833e0522                 ld t4, 0x220(a0)
   256: b38e6e00                 add t4, t4, t1
   260: 070b0e02                 vle8.v v22, (t3)
   264: 878b0e02                 vle8.v v23, (t4)
   268: 570b7b2f                 vxor.vv v22, v23, v22
   272: 57010b5e                 vmv.v.v v2, v22
   276: d7090b5e                 vmv.v.v v19, v22
   280: 033e0525                 ld t3, 0x250(a0)
   284: 330e6e00                 add t3, t3, t1
   288: 833e852a                 ld t4, 0x2a8(a0)
   292: b38e6e00                 add t4, t4, t1
   296: 070b0e02                 vle8.v v22, (t3)
   300: 878b0e02                 vle8.v v23, (t4)
   304: 570b7b2f                 vxor.vv v22, v23, v22
   308: 57040b5e                 vmv.v.v v8, v22
   312: d7060b5e                 vmv.v.v v13, v22
   316: 83330500                 ld t2, 0(a0)
   320: b3836300                 add t2, t2, t1
   324: 078b0302                 vle8.v v22, (t2)
   328: 5700602f                 vxor.vv v0, v22, v0
   332: d7050b5e                 vmv.v.v v11, v22
   336: 83338500                 ld t2, 8(a0)
   340: b3836300                 add t2, t2, t1
   344: 078b0302                 vle8.v v22, (t2)
   348: d780602f                 vxor.vv v1, v22, v1
   352: 57060b5e                 vmv.v.v v12, v22
   356: 83330501                 ld t2, 0x10(a0)
   360: b3836300                 add t2, t2, t1
   364: 078b0302                 vle8.v v22, (t2)
   368: 5701612f                 vxor.vv v2, v22, v2
   372: d786662f                 vxor.vv v13, v22, v13
   376: 83338501                 ld t2, 0x18(a0)
   380: b3836300                 add t2, t2, t1
   384: 078b0302                 vle8.v v22, (t2)
   388: d7010b5e                 vmv.v.v v3, v22
   392: 5707672f                 vxor.vv v14, v22, v14
   396: 83330502                 ld t2, 0x20(a0)
   400: b3836300                 add t2, t2, t1
   404: 078b0302                 vle8.v v22, (t2)
   408: 57020b5e                 vmv.v.v v4, v22
   412: d787672f                 vxor.vv v15, v22, v15
   416: 83330503                 ld t2, 0x30(a0)
   420: b3836300                 add t2, t2, t1
   424: 078b0302                 vle8.v v22, (t2)
   428: 5703632f                 vxor.vv v6, v22, v6
   432: d7080b5e                 vmv.v.v v17, v22
   436: 83338503                 ld t2, 0x38(a0)
   440: b3836300                 add t2, t2, t1
   444: 078b0302                 vle8.v v22, (t2)
   448: d783632f                 vxor.vv v7, v22, v7
   452: 57090b5e                 vmv.v.v v18, v22
   456: 83330504                 ld t2, 0x40(a0)
   460: b3836300                 add t2, t2, t1
   464: 078b0302                 vle8.v v22, (t2)
   468: 5704642f                 vxor.vv v8, v22, v8
   472: d789692f                 vxor.vv v19, v22, v19
   476: 83338504                 ld t2, 0x48(a0)
   480: b3836300                 add t2, t2, t1
   484: 078b0302                 vle8.v v22, (t2)
   488: d7040b5e                 vmv.v.v v9, v22
   492: 570a6a2f                 vxor.vv v20, v22, v20
   496: 83330505                 ld t2, 0x50(a0)
   500: b3836300                 add t2, t2, t1
   504: 078b0302                 vle8.v v22, (t2)
   508: 57050b5e                 vmv.v.v v10, v22
   512: d78a6a2f                 vxor.vv v21, v22, v21
   516: 83330506                 ld t2, 0x60(a0)
   520: b3836300                 add t2, t2, t1
   524: 078b0302                 vle8.v v22, (t2)
   528: d780602f                 vxor.vv v1, v22, v1
   532: d785652f                 vxor.vv v11, v22, v11
   536: 83338506                 ld t2, 0x68(a0)
   540: b3836300                 add t2, t2, t1
   544: 078b0302                 vle8.v v22, (t2)
   548: 5701612f                 vxor.vv v2, v22, v2
   552: 5706662f                 vxor.vv v12, v22, v12
   556: 83330507                 ld t2, 0x70(a0)
   560: b3836300                 add t2, t2, t1
   564: 078b0302                 vle8.v v22, (t2)
   568: d781612f                 vxor.vv v3, v22, v3
   572: d786662f                 vxor.vv v13, v22, v13
   576: 83338507                 ld t2, 0x78(a0)
   580: b3836300                 add t2, t2, t1
   584: 078b0302                 vle8.v v22, (t2)
   588: 5702622f                 vxor.vv v4, v22, v4
   592: 5707672f                 vxor.vv v14, v22, v14
   596: 83330508                 ld t2, 0x80(a0)
   600: b3836300                 add t2, t2, t1
   604: 078b0302                 vle8.v v22, (t2)
   608: d787672f                 vxor.vv v15, v22, v15
   612: 83338508                 ld t2, 0x88(a0)
   616: b3836300                 add t2, t2, t1
   620: 078b0302                 vle8.v v22, (t2)
   624: 5703632f                 vxor.vv v6, v22, v6
   628: 5708682f                 vxor.vv v16, v22, v16
   632: 83330509                 ld t2, 0x90(a0)
   636: b3836300                 add t2, t2, t1
   640: 078b0302                 vle8.v v22, (t2)
   644: d783632f                 vxor.vv v7, v22, v7
   648: d788682f                 vxor.vv v17, v22, v17
   652: 83338509                 ld t2, 0x98(a0)
   656: b3836300                 add t2, t2, t1
   660: 078b0302                 vle8.v v22, (t2)
   664: 5704642f                 vxor.vv v8, v22, v8
   668: 5709692f                 vxor.vv v18, v22, v18
   672: 8333050a                 ld t2, 0xa0(a0)
   676: b3836300                 add t2, t2, t1
   680: 078b0302                 vle8.v v22, (t2)
   684: d784642f                 vxor.vv v9, v22, v9
   688: d789692f                 vxor.vv v19, v22, v19
   692: 8333850a                 ld t2, 0xa8(a0)
   696: b3836300                 add t2, t2, t1
   700: 078b0302                 vle8.v v22, (t2)
   704: 5705652f                 vxor.vv v10, v22, v10
   708: 570a6a2f                 vxor.vv v20, v22, v20
   712: 8333050b                 ld t2, 0xb0(a0)
   716: b3836300                 add t2, t2, t1
   720: 078b0302                 vle8.v v22, (t2)
   724: 570a6a2f                 vxor.vv v20, v22, v20
   728: 8333850b                 ld t2, 0xb8(a0)
   732: b3836300                 add t2, t2, t1
   736: 078b0302                 vle8.v v22, (t2)
   740: d780602f                 vxor.vv v1, v22, v1
   744: d78a6a2f                 vxor.vv v21, v22, v21
   748: 8333050c                 ld t2, 0xc0(a0)
   752: b3836300                 add t2, t2, t1
   756: 078b0302                 vle8.v v22, (t2)
   760: 5701612f                 vxor.vv v2, v22, v2
   764: d785652f                 vxor.vv v11, v22, v11
   768: 8333850c                 ld t2, 0xc8(a0)
   772: b3836300                 add t2, t2, t1
   776: 078b0302                 vle8.v v22, (t2)
   780: d781612f                 vxor.vv v3, v22, v3
   784: 5706662f                 vxor.vv v12, v22, v12
   788: 8333050d                 ld t2, 0xd0(a0)
   792: b3836300                 add t2, t2, t1
   796: 078b0302                 vle8.v v22, (t2)
   800: 5702622f                 vxor.vv v4, v22, v4
   804: d786662f                 vxor.vv v13, v22, v13
   808: 8333850d                 ld t2, 0xd8(a0)
   812: b3836300                 add t2, t2, t1
   816: 078b0302                 vle8.v v22, (t2)
   820: d782622f                 vxor.vv v5, v22, v5
   824: 5707672f                 vxor.vv v14, v22, v14
   828: 8333850e                 ld t2, 0xe8(a0)
   832: b3836300                 add t2, t2, t1
   836: 078b0302                 vle8.v v22, (t2)
   840: d783632f                 vxor.vv v7, v22, v7
   844: 5708682f                 vxor.vv v16, v22, v16
   848: 8333050f                 ld t2, 0xf0(a0)
   852: b3836300                 add t2, t2, t1
   856: 078b0302                 vle8.v v22, (t2)
   860: 5704642f                 vxor.vv v8, v22, v8
   864: d788682f                 vxor.vv v17, v22, v17
   868: 8333850f                 ld t2, 0xf8(a0)
   872: b3836300                 add t2, t2, t1
   876: 078b0302                 vle8.v v22, (t2)
   880: d784642f                 vxor.vv v9, v22, v9
   884: 5709692f                 vxor.vv v18, v22, v18
   888: 83330510                 ld t2, 0x100(a0)
   892: b3836300                 add t2, t2, t1
   896: 078b0302                 vle8.v v22, (t2)
   900: 5705652f                 vxor.vv v10, v22, v10
   904: d789692f                 vxor.vv v19, v22, v19
   908: 83338510                 ld t2, 0x108(a0)
   912: b3836300                 add t2, t2, t1
   916: 078b0302                 vle8.v v22, (t2)
   920: 5700602f                 vxor.vv v0, v22, v0
   924: d789692f                 vxor.vv v19, v22, v19
   928: 83338511                 ld t2, 0x118(a0)
   932: b3836300                 add t2, t2, t1
   936: 078b0302                 vle8.v v22, (t2)
   940: 5701612f                 vxor.vv v2, v22, v2
   944: d78a6a2f                 vxor.vv v21, v22, v21
   948: 83330512                 ld t2, 0x120(a0)
   952: b3836300                 add t2, t2, t1
   956: 078b0302                 vle8.v v22, (t2)
   960: d781612f                 vxor.vv v3, v22, v3
   964: d785652f                 vxor.vv v11, v22, v11
   968: 83338512                 ld t2, 0x128(a0)
   972: b3836300                 add t2, t2, t1
   976: 078b0302                 vle8.v v22, (t2)
   980: 5702622f                 vxor.vv v4, v22, v4
   984: 5706662f                 vxor.vv v12, v22, v12
   988: 83330513                 ld t2, 0x130(a0)
   992: b3836300                 add t2, t2, t1
   996: 078b0302                 vle8.v v22, (t2)
  1000: d782622f                 vxor.vv v5, v22, v5
  1004: d786662f                 vxor.vv v13, v22, v13
  1008: 83338513                 ld t2, 0x138(a0)
  1012: b3836300                 add t2, t2, t1
  1016: 078b0302                 vle8.v v22, (t2)
  1020: 5707672f                 vxor.vv v14, v22, v14
  1024: 83330514                 ld t2, 0x140(a0)
  1028: b3836300                 add t2, t2, t1
  1032: 078b0302                 vle8.v v22, (t2)
  1036: d783632f                 vxor.vv v7, v22, v7
  1040: d787672f                 vxor.vv v15, v22, v15
  1044: 83338514                 ld t2, 0x148(a0)
  1048: b3836300                 add t2, t2, t1
  1052: 078b0302                 vle8.v v22, (t2)
  1056: 5704642f                 vxor.vv v8, v22, v8
  1060: 5708682f                 vxor.vv v16, v22, v16
  1064: 83330515                 ld t2, 0x150(a0)
  1068: b3836300                 add t2, t2, t1
  1072: 078b0302                 vle8.v v22, (t2)
  1076: d784642f                 vxor.vv v9, v22, v9
  1080: d788682f                 vxor.vv v17, v22, v17
  1084: 83338515                 ld t2, 0x158(a0)
  1088: b3836300                 add t2, t2, t1
  1092: 078b0302                 vle8.v v22, (t2)
  1096: 5705652f                 vxor.vv v10, v22, v10
  1100: 5709692f                 vxor.vv v18, v22, v18
  1104: 83330516                 ld t2, 0x160(a0)
  1108: b3836300                 add t2, t2, t1
  1112: 078b0302                 vle8.v v22, (t2)
  1116: 5700602f                 vxor.vv v0, v22, v0
  1120: 5709692f                 vxor.vv v18, v22, v18
  1124: 83338516                 ld t2, 0x168(a0)
  1128: b3836300                 add t2, t2, t1
  1132: 078b0302                 vle8.v v22, (t2)
  1136: d789692f                 vxor.vv v19, v22, v19
  1140: 83330517                 ld t2, 0x170(a0)
  1144: b3836300                 add t2, t2, t1
  1148: 078b0302                 vle8.v v22, (t2)
  1152: 5701612f                 vxor.vv v2, v22, v2
  1156: 570a6a2f                 vxor.vv v20, v22, v20
  1160: 83338517                 ld t2, 0x178(a0)
  1164: b3836300                 add t2, t2, t1
  1168: 078b0302                 vle8.v v22, (t2)
  1172: d781612f                 vxor.vv v3, v22, v3
  1176: d78a6a2f                 vxor.vv v21, v22, v21
  1180: 83330518                 ld t2, 0x180(a0)
  1184: b3836300                 add t2, t2, t1
  1188: 078b0302                 vle8.v v22, (t2)
  1192: 5702622f                 vxor.vv v4, v22, v4
  1196: d785652f                 vxor.vv v11, v22, v11
  1200: 83338518                 ld t2, 0x188(a0)
  1204: b3836300                 add t2, t2, t1
  1208: 078b0302                 vle8.v v22, (t2)
  1212: d782622f                 vxor.vv v5, v22, v5
  1216: 5706662f                 vxor.vv v12, v22, v12
  1220: 83330519                 ld t2, 0x190(a0)
  1224: b3836300                 add t2, t2, t1
  1228: 078b0302                 vle8.v v22, (t2)
  1232: 5703632f                 vxor.vv v6, v22, v6
  1236: d786662f                 vxor.vv v13, v22, v13
  1240: 8333051a                 ld t2, 0x1a0(a0)
  1244: b3836300                 add t2, t2, t1
  1248: 078b0302                 vle8.v v22, (t2)
  1252: 5704642f                 vxor.vv v8, v22, v8
  1256: d787672f                 vxor.vv v15, v22, v15
  1260: 8333851a                 ld t2, 0x1a8(a0)
  1264: b3836300                 add t2, t2, t1
  1268: 078b0302                 vle8.v v22, (t2)
  1272: d784642f                 vxor.vv v9, v22, v9
  1276: 5708682f                 vxor.vv v16, v22, v16
  1280: 8333051b                 ld t2, 0x1b0(a0)
  1284: b3836300                 add t2, t2, t1
  1288: 078b0302                 vle8.v v22, (t2)
  1292: 5705652f                 vxor.vv v10, v22, v10
  1296: d788682f                 vxor.vv v17, v22, v17
  1300: 8333851b                 ld t2, 0x1b8(a0)
  1304: b3836300                 add t2, t2, t1
  1308: 078b0302                 vle8.v v22, (t2)
  1312: 5700602f                 vxor.vv v0, v22, v0
  1316: d788682f                 vxor.vv v17, v22, v17
  1320: 8333051c                 ld t2, 0x1c0(a0)
  1324: b3836300                 add t2, t2, t1
  1328: 078b0302                 vle8.v v22, (t2)
  1332: d780602f                 vxor.vv v1, v22, v1
  1336: 5709692f                 vxor.vv v18, v22, v18
  1340: 8333051d                 ld t2, 0x1d0(a0)
  1344: b3836300                 add t2, t2, t1
  1348: 078b0302                 vle8.v v22, (t2)
  1352: d781612f                 vxor.vv v3, v22, v3
  1356: 570a6a2f                 vxor.vv v20, v22, v20
  1360: 8333851d                 ld t2, 0x1d8(a0)
  1364: b3836300                 add t2, t2, t1
  1368: 078b0302                 vle8.v v22, (t2)
  1372: 5702622f                 vxor.vv v4, v22, v4
  1376: d78a6a2f                 vxor.vv v21, v22, v21
  1380: 8333051e                 ld t2, 0x1e0(a0)
  1384: b3836300                 add t2, t2, t1
  1388: 078b0302                 vle8.v v22, (t2)
  1392: d782622f                 vxor.vv v5, v22, v5
  1396: d785652f                 vxor.vv v11, v22, v11
  1400: 8333851e                 ld t2, 0x1e8(a0)
  1404: b3836300                 add t2, t2, t1
  1408: 078b0302                 vle8.v v22, (t2)
  1412: 5703632f                 vxor.vv v6, v22, v6
  1416: 5706662f                 vxor.vv v12, v22, v12
  1420: 8333051f                 ld t2, 0x1f0(a0)
  1424: b3836300                 add t2, t2, t1
  1428: 078b0302                 vle8.v v22, (t2)
  1432: d786662f                 vxor.vv v13, v22, v13
  1436: 8333851f                 ld t2, 0x1f8(a0)
  1440: b3836300                 add t2, t2, t1
  1444: 078b0302                 vle8.v v22, (t2)
  1448: 5704642f                 vxor.vv v8, v22, v8
  1452: 5707672f                 vxor.vv v14, v22, v14
  1456: 83330520                 ld t2, 0x200(a0)
  1460: b3836300                 add t2, t2, t1
  1464: 078b0302                 vle8.v v22, (t2)
  1468: d784642f                 vxor.vv v9, v22, v9
  1472: d787672f                 vxor.vv v15, v22, v15
  1476: 83338520                 ld t2, 0x208(a0)
  1480: b3836300                 add t2, t2, t1
  1484: 078b0302                 vle8.v v22, (t2)
  1488: 5705652f                 vxor.vv v10, v22, v10
  1492: 5708682f                 vxor.vv v16, v22, v16
  1496: 83330521                 ld t2, 0x210(a0)
  1500: b3836300                 add t2, t2, t1
  1504: 078b0302                 vle8.v v22, (t2)
  1508: 5700602f                 vxor.vv v0, v22, v0
  1512: 5708682f                 vxor.vv v16, v22, v16
  1516: 83338521                 ld t2, 0x218(a0)
  1520: b3836300                 add t2, t2, t1
  1524: 078b0302                 vle8.v v22, (t2)
  1528: d780602f                 vxor.vv v1, v22, v1
  1532: d788682f                 vxor.vv v17, v22, v17
  1536: 83330522                 ld t2, 0x220(a0)
  1540: b3836300                 add t2, t2, t1
  1544: 078b0302                 vle8.v v22, (t2)
  1548: 5709692f                 vxor.vv v18, v22, v18
  1552: 83338522                 ld t2, 0x228(a0)
  1556: b3836300                 add t2, t2, t1
  1560: 078b0302                 vle8.v v22, (t2)
  1564: d781612f                 vxor.vv v3, v22, v3
  1568: d789692f                 vxor.vv v19, v22, v19
  1572: 83330523                 ld t2, 0x230(a0)
  1576: b3836300                 add t2, t2, t1
  1580: 078b0302                 vle8.v v22, (t2)
  1584: 5702622f                 vxor.vv v4, v22, v4
  1588: 570a6a2f                 vxor.vv v20, v22, v20
  1592: 83338523                 ld t2, 0x238(a0)
  1596: b3836300                 add t2, t2, t1
  1600: 078b0302                 vle8.v v22, (t2)
  1604: d782622f                 vxor.vv v5, v22, v5
  1608: d78a6a2f                 vxor.vv v21, v22, v21
  1612: 83330524                 ld t2, 0x240(a0)
  1616: b3836300                 add t2, t2, t1
  1620: 078b0302                 vle8.v v22, (t2)
  1624: 5703632f                 vxor.vv v6, v22, v6
  1628: d785652f                 vxor.vv v11, v22, v11
  1632: 83338524                 ld t2, 0x248(a0)
  1636: b3836300                 add t2, t2, t1
  1640: 078b0302                 vle8.v v22, (t2)
  1644: d783632f                 vxor.vv v7, v22, v7
  1648: 5706662f                 vxor.vv v12, v22, v12
  1652: 83338525                 ld t2, 0x258(a0)
  1656: b3836300                 add t2, t2, t1
  1660: 078b0302                 vle8.v v22, (t2)
  1664: d784642f                 vxor.vv v9, v22, v9
  1668: 5707672f                 vxor.vv v14, v22, v14
  1672: 83330526                 ld t2, 0x260(a0)
  1676: b3836300                 add t2, t2, t1
  1680: 078b0302                 vle8.v v22, (t2)
  1684: 5705652f                 vxor.vv v10, v22, v10
  1688: d787672f                 vxor.vv v15, v22, v15
  1692: 83338526                 ld t2, 0x268(a0)
  1696: b3836300                 add t2, t2, t1
  1700: 078b0302                 vle8.v v22, (t2)
  1704: 5700602f                 vxor.vv v0, v22, v0
  1708: d787672f                 vxor.vv v15, v22, v15
  1712: 83330527                 ld t2, 0x270(a0)
  1716: b3836300                 add t2, t2, t1
  1720: 078b0302                 vle8.v v22, (t2)
  1724: d780602f                 vxor.vv v1, v22, v1
  1728: 5708682f                 vxor.vv v16, v22, v16
  1732: 83338527                 ld t2, 0x278(a0)
  1736: b3836300                 add t2, t2, t1
  1740: 078b0302                 vle8.v v22, (t2)
  1744: 5701612f                 vxor.vv v2, v22, v2
  1748: d788682f                 vxor.vv v17, v22, v17
  1752: 83330528                 ld t2, 0x280(a0)
  1756: b3836300                 add t2, t2, t1
  1760: 078b0302                 vle8.v v22, (t2)
  1764: d781612f                 vxor.vv v3, v22, v3
  1768: 5709692f                 vxor.vv v18, v22, v18
  1772: 83338528                 ld t2, 0x288(a0)
  1776: b3836300                 add t2, t2, t1
  1780: 078b0302                 vle8.v v22, (t2)
  1784: 5702622f                 vxor.vv v4, v22, v4
  1788: d789692f                 vxor.vv v19, v22, v19
  1792: 83330529                 ld t2, 0x290(a0)
  1796: b3836300                 add t2, t2, t1
  1800: 078b0302                 vle8.v v22, (t2)
  1804: d782622f                 vxor.vv v5, v22, v5
  1808: 570a6a2f                 vxor.vv v20, v22, v20
  1812: 83338529                 ld t2, 0x298(a0)
  1816: b3836300                 add t2, t2, t1
  1820: 078b0302                 vle8.v v22, (t2)
  1824: 5703632f                 vxor.vv v6, v22, v6
  1828: d78a6a2f                 vxor.vv v21, v22, v21
  1832: 8333052a                 ld t2, 0x2a0(a0)
  1836: b3836300                 add t2, t2, t1
  1840: 078b0302                 vle8.v v22, (t2)
  1844: d783632f                 vxor.vv v7, v22, v7
  1848: d785652f                 vxor.vv v11, v22, v11
  1852: 8333852a                 ld t2, 0x2a8(a0)
  1856: b3836300                 add t2, t2, t1
  1860: 078b0302                 vle8.v v22, (t2)
  1864: 5706662f                 vxor.vv v12, v22, v12
  1868: 8333052b                 ld t2, 0x2b0(a0)
  1872: b3836300                 add t2, t2, t1
  1876: 078b0302                 vle8.v v22, (t2)
  1880: d784642f                 vxor.vv v9, v22, v9
  1884: d786662f                 vxor.vv v13, v22, v13
  1888: 8333852b                 ld t2, 0x2b8(a0)
  1892: b3836300                 add t2, t2, t1
  1896: 078b0302                 vle8.v v22, (t2)
  1900: 5705652f                 vxor.vv v10, v22, v10
  1904: 5707672f                 vxor.vv v14, v22, v14
  1908: 83b30500                 ld t2, 0(a1)
  1912: b3836300                 add t2, t2, t1
  1916: 27800302                 vse8.v v0, (t2)
  1920: 83b38500                 ld t2, 8(a1)
  1924: b3836300                 add t2, t2, t1
  1928: a7800302                 vse8.v v1, (t2)
  1932: 83b30501                 ld t2, 0x10(a1)
  1936: b3836300                 add t2, t2, t1
  1940: 27810302                 vse8.v v2, (t2)
  1944: 83b38501                 ld t2, 0x18(a1)
  1948: b3836300                 add t2, t2, t1
  1952: a7810302                 vse8.v v3, (t2)
  1956: 83b30502                 ld t2, 0x20(a1)
  1960: b3836300                 add t2, t2, t1
  1964: 27820302                 vse8.v v4, (t2)
  1968: 83b38502                 ld t2, 0x28(a1)
  1972: b3836300                 add t2, t2, t1
  1976: a7820302                 vse8.v v5, (t2)
  1980: 83b30503                 ld t2, 0x30(a1)
  1984: b3836300                 add t2, t2, t1
  1988: 27830302                 vse8.v v6, (t2)
  1992: 83b38503                 ld t2, 0x38(a1)
  1996: b3836300                 add t2, t2, t1
  2000: a7830302                 vse8.v v7, (t2)
  2004: 83b30504                 ld t2, 0x40(a1)
  2008: b3836300                 add t2, t2, t1
  2012: 27840302                 vse8.v v8, (t2)
  2016: 83b38504                 ld t2, 0x48(a1)
  2020: b3836300                 add t2, t2, t1
  2024: a7840302                 vse8.v v9, (t2)
  2028: 83b30505                 ld t2, 0x50(a1)
  2032: b3836300                 add t2, t2, t1
  2036: 27850302                 vse8.v v10, (t2)
  2040: 83b38505                 ld t2, 0x58(a1)
  2044: b3836300                 add t2, t2, t1
  2048: a7850302                 vse8.v v11, (t2)
  2052: 83b30506                 ld t2, 0x60(a1)
  2056: b3836300                 add t2, t2, t1
  2060: 27860302                 vse8.v v12, (t2)
  2064: 83b38506                 ld t2, 0x68(a1)
  2068: b3836300                 add t2, t2, t1
  2072: a7860302                 vse8.v v13, (t2)
  2076: 83b30507                 ld t2, 0x70(a1)
  2080: b3836300                 add t2, t2, t1
  2084: 27870302                 vse8.v v14, (t2)
  2088: 83b38507                 ld t2, 0x78(a1)
  2092: b3836300                 add t2, t2, t1
  2096: a7870302                 vse8.v v15, (t2)
  2100: 83b30508                 ld t2, 0x80(a1)
  2104: b3836300                 add t2, t2, t1
  2108: 27880302                 vse8.v v16, (t2)
  2112: 83b38508                 ld t2, 0x88(a1)
  2116: b3836300                 add t2, t2, t1
  2120: a7880302                 vse8.v v17, (t2)
  2124: 83b30509                 ld t2, 0x90(a1)
  2128: b3836300                 add t2, t2, t1
  2132: 27890302                 vse8.v v18, (t2)
  2136: 83b38509                 ld t2, 0x98(a1)
  2140: b3836300                 add t2, t2, t1
  2144: a7890302                 vse8.v v19, (t2)
  2148: 83b3050a                 ld t2, 0xa0(a1)
  2152: b3836300                 add t2, t2, t1
  2156: 278a0302                 vse8.v v20, (t2)
  2160: 83b3850a                 ld t2, 0xa8(a1)
  2164: b3836300                 add t2, t2, t1
  2168: a78a0302                 vse8.v v21, (t2)
  2172: 3303d300                 add t1, t1, a3
  2176: 1307f7ff                 addi a4, a4, -1
  2180: 631e07fa                 bnez a4, 0x40
  2184: 33066640                 sub a2, a2, t1
  2188: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
  2192: 63940200                 bnez t0, 0x898
  2196: 6f00d004                 j 0x10e0
  2200: 033e8502                 ld t3, 0x28(a0)
  2204: 330e6e00                 add t3, t3, t1
  2208: 833e0508                 ld t4, 0x80(a0)
  2212: b38e6e00                 add t4, t4, t1
  2216: 070b0e02                 vle8.v v22, (t3)
  2220: 878b0e02                 vle8.v v23, (t4)
  2224: 570b7b2f                 vxor.vv v22, v23, v22
  2228: d7020b5e                 vmv.v.v v5, v22
  2232: 57080b5e                 vmv.v.v v16, v22
  2236: 033e8505                 ld t3, 0x58(a0)
  2240: 330e6e00                 add t3, t3, t1
  2244: 833e050b                 ld t4, 0xb0(a0)
  2248: b38e6e00                 add t4, t4, t1
  2252: 070b0e02                 vle8.v v22, (t3)
  2256: 878b0e02                 vle8.v v23, (t4)
  2260: 570b7b2f                 vxor.vv v22, v23, v22
  2264: 57000b5e                 vmv.v.v v0, v22
  2268: d70a0b5e                 vmv.v.v v21, v22
  2272: 033e050e                 ld t3, 0xe0(a0)
  2276: 330e6e00                 add t3, t3, t1
  2280: 833e8513                 ld t4, 0x138(a0)
  2284: b38e6e00                 add t4, t4, t1
  2288: 070b0e02                 vle8.v v22, (t3)
  2292: 878b0e02                 vle8.v v23, (t4)
  2296: 570b7b2f                 vxor.vv v22, v23, v22
  2300: 57030b5e                 vmv.v.v v6, v22
  2304: d7070b5e                 vmv.v.v v15, v22
  2308: 033e0511                 ld t3, 0x110(a0)
  2312: 330e6e00                 add t3, t3, t1
  2316: 833e8516                 ld t4, 0x168(a0)
  2320: b38e6e00                 add t4, t4, t1
  2324: 070b0e02                 vle8.v v22, (t3)
  2328: 878b0e02                 vle8.v v23, (t4)
  2332: 570b7b2f                 vxor.vv v22, v23, v22
  2336: d7000b5e                 vmv.v.v v1, v22
  2340: 570a0b5e                 vmv.v.v v20, v22
  2344: 033e8519                 ld t3, 0x198(a0)
  2348: 330e6e00                 add t3, t3, t1
  2352: 833e051f                 ld t4, 0x1f0(a0)
  2356: b38e6e00                 add t4, t4, t1
  2360: 070b0e02                 vle8.v v22, (t3)
  2364: 878b0e02                 vle8.v v23, (t4)
  2368: 570b7b2f                 vxor.vv v22, v23, v22
  2372: d7030b5e                 vmv.v.v v7, v22
  2376: 57070b5e                 vmv.v.v v14, v22
  2380: 033e851c                 ld t3, 0x1c8(a0)
  2384: 330e6e00                 add t3, t3, t1
  2388: 833e0522                 ld t4, 0x220(a0)
  2392: b38e6e00                 add t4, t4, t1
  2396: 070b0e02                 vle8.v v22, (t3)
  2400: 878b0e02                 vle8.v v23, (t4)
  2404: 570b7b2f                 vxor.vv v22, v23, v22
  2408: 57010b5e                 vmv.v.v v2, v22
  2412: d7090b5e                 vmv.v.v v19, v22
  2416: 033e0525                 ld t3, 0x250(a0)
  2420: 330e6e00                 add t3, t3, t1
  2424: 833e852a                 ld t4, 0x2a8(a0)
  2428: b38e6e00                 add t4, t4, t1
  2432: 070b0e02                 vle8.v v22, (t3)
  2436: 878b0e02                 vle8.v v23, (t4)
  2440: 570b7b2f                 vxor.vv v22, v23, v22
  2444: 57040b5e                 vmv.v.v v8, v22
  2448: d7060b5e                 vmv.v.v v13, v22
  2452: 83330500                 ld t2, 0(a0)
  2456: b3836300                 add t2, t2, t1
  2460: 078b0302                 vle8.v v22, (t2)
  2464: 5700602f                 vxor.vv v0, v22, v0
  2468: d7050b5e                 vmv.v.v v11, v22
  2472: 83338500                 ld t2, 8(a0)
  2476: b3836300                 add t2, t2, t1
  2480: 078b0302                 vle8.v v22, (t2)
  2484: d780602f                 vxor.vv v1, v22, v1
  2488: 57060b5e                 vmv.v.v v12, v22
  2492: 83330501                 ld t2, 0x10(a0)
  2496: b3836300                 add t2, t2, t1
  2500: 078b0302                 vle8.v v22, (t2)
  2504: 5701612f                 vxor.vv v2, v22, v2
  2508: d786662f                 vxor.vv v13, v22, v13
  2512: 83338501                 ld t2, 0x18(a0)
  2516: b3836300                 add t2, t2, t1
  2520: 078b0302                 vle8.v v22, (t2)
  2524: d7010b5e                 vmv.v.v v3, v22
  2528: 5707672f                 vxor.vv v14, v22, v14
  2532: 83330502                 ld t2, 0x20(a0)
  2536: b3836300                 add t2, t2, t1
  2540: 078b0302                 vle8.v v22, (t2)
  2544: 57020b5e                 vmv.v.v v4, v22
  2548: d787672f                 vxor.vv v15, v22, v15
  2552: 83330503                 ld t2, 0x30(a0)
  2556: b3836300                 add t2, t2, t1
  2560: 078b0302                 vle8.v v22, (t2)
  2564: 5703632f                 vxor.vv v6, v22, v6
  2568: d7080b5e                 vmv.v.v v17, v22
  2572: 83338503                 ld t2, 0x38(a0)
  2576: b3836300                 add t2, t2, t1
  2580: 078b0302                 vle8.v v22, (t2)
  2584: d783632f                 vxor.vv v7, v22, v7
  2588: 57090b5e                 vmv.v.v v18, v22
  2592: 83330504                 ld t2, 0x40(a0)
  2596: b3836300                 add t2, t2, t1
  2600: 078b0302                 vle8.v v22, (t2)
  2604: 5704642f                 vxor.vv v8, v22, v8
  2608: d789692f                 vxor.vv v19, v22, v19
  2612: 83338504                 ld t2, 0x48(a0)
  2616: b3836300                 add t2, t2, t1
  2620: 078b0302                 vle8.v v22, (t2)
  2624: d7040b5e                 vmv.v.v v9, v22
  2628: 570a6a2f                 vxor.vv v20, v22, v20
  2632: 83330505                 ld t2, 0x50(a0)
  2636: b3836300                 add t2, t2, t1
  2640: 078b0302                 vle8.v v22, (t2)
  2644: 57050b5e                 vmv.v.v v10, v22
  2648: d78a6a2f                 vxor.vv v21, v22, v21
  2652: 83330506                 ld t2, 0x60(a0)
  2656: b3836300                 add t2, t2, t1
  2660: 078b0302                 vle8.v v22, (t2)
  2664: d780602f                 vxor.vv v1, v22, v1
  2668: d785652f                 vxor.vv v11, v22, v11
  2672: 83338506                 ld t2, 0x68(a0)
  2676: b3836300                 add t2, t2, t1
  2680: 078b0302                 vle8.v v22, (t2)
  2684: 5701612f                 vxor.vv v2, v22, v2
  2688: 5706662f                 vxor.vv v12, v22, v12
  2692: 83330507                 ld t2, 0x70(a0)
  2696: b3836300                 add t2, t2, t1
  2700: 078b0302                 vle8.v v22, (t2)
  2704: d781612f                 vxor.vv v3, v22, v3
  2708: d786662f                 vxor.vv v13, v22, v13
  2712: 83338507                 ld t2, 0x78(a0)
  2716: b3836300                 add t2, t2, t1
  2720: 078b0302                 vle8.v v22, (t2)
  2724: 5702622f                 vxor.vv v4, v22, v4
  2728: 5707672f                 vxor.vv v14, v22, v14
  2732: 83330508                 ld t2, 0x80(a0)
  2736: b3836300                 add t2, t2, t1
  2740: 078b0302                 vle8.v v22, (t2)
  2744: d787672f                 vxor.vv v15, v22, v15
  2748: 83338508                 ld t2, 0x88(a0)
  2752: b3836300                 add t2, t2, t1
  2756: 078b0302                 vle8.v v22, (t2)
  2760: 5703632f                 vxor.vv v6, v22, v6
  2764: 5708682f                 vxor.vv v16, v22, v16
  2768: 83330509                 ld t2, 0x90(a0)
  2772: b3836300                 add t2, t2, t1
  2776: 078b0302                 vle8.v v22, (t2)
  2780: d783632f                 vxor.vv v7, v22, v7
  2784: d788682f                 vxor.vv v17, v22, v17
  2788: 83338509                 ld t2, 0x98(a0)
  2792: b3836300                 add t2, t2, t1
  2796: 078b0302                 vle8.v v22, (t2)
  2800: 5704642f                 vxor.vv v8, v22, v8
  2804: 5709692f                 vxor.vv v18, v22, v18
  2808: 8333050a                 ld t2, 0xa0(a0)
  2812: b3836300                 add t2, t2, t1
  2816: 078b0302                 vle8.v v22, (t2)
  2820: d784642f                 vxor.vv v9, v22, v9
  2824: d789692f                 vxor.vv v19, v22, v19
  2828: 8333850a                 ld t2, 0xa8(a0)
  2832: b3836300                 add t2, t2, t1
  2836: 078b0302                 vle8.v v22, (t2)
  2840: 5705652f                 vxor.vv v10, v22, v10
  2844: 570a6a2f                 vxor.vv v20, v22, v20
  2848: 8333050b                 ld t2, 0xb0(a0)
  2852: b3836300                 add t2, t2, t1
  2856: 078b0302                 vle8.v v22, (t2)
  2860: 570a6a2f                 vxor.vv v20, v22, v20
  2864: 8333850b                 ld t2, 0xb8(a0)
  2868: b3836300                 add t2, t2, t1
  2872: 078b0302                 vle8.v v22, (t2)
  2876: d780602f                 vxor.vv v1, v22, v1
  2880: d78a6a2f                 vxor.vv v21, v22, v21
  2884: 8333050c                 ld t2, 0xc0(a0)
  2888: b3836300                 add t2, t2, t1
  2892: 078b0302                 vle8.v v22, (t2)
  2896: 5701612f                 vxor.vv v2, v22, v2
  2900: d785652f                 vxor.vv v11, v22, v11
  2904: 8333850c                 ld t2, 0xc8(a0)
  2908: b3836300                 add t2, t2, t1
  2912: 078b0302                 vle8.v v22, (t2)
  2916: d781612f                 vxor.vv v3, v22, v3
  2920: 5706662f                 vxor.vv v12, v22, v12
  2924: 8333050d                 ld t2, 0xd0(a0)
  2928: b3836300                 add t2, t2, t1
  2932: 078b0302                 vle8.v v22, (t2)
  2936: 5702622f                 vxor.vv v4, v22, v4
  2940: d786662f                 vxor.vv v13, v22, v13
  2944: 8333850d                 ld t2, 0xd8(a0)
  2948: b3836300                 add t2, t2, t1
  2952: 078b0302                 vle8.v v22, (t2)
  2956: d782622f                 vxor.vv v5, v22, v5
  2960: 5707672f                 vxor.vv v14, v22, v14
  2964: 8333850e                 ld t2, 0xe8(a0)
  2968: b3836300                 add t2, t2, t1
  2972: 078b0302                 vle8.v v22, (t2)
  2976: d783632f                 vxor.vv v7, v22, v7
  2980: 5708682f                 vxor.vv v16, v22, v16
  2984: 8333050f                 ld t2, 0xf0(a0)
  2988: b3836300                 add t2, t2, t1
  2992: 078b0302                 vle8.v v22, (t2)
  2996: 5704642f                 vxor.vv v8, v22, v8
  3000: d788682f                 vxor.vv v17, v22, v17
  3004: 8333850f                 ld t2, 0xf8(a0)
  3008: b3836300                 add t2, t2, t1
  3012: 078b0302                 vle8.v v22, (t2)
  3016: d784642f                 vxor.vv v9, v22, v9
  3020: 5709692f                 vxor.vv v18, v22, v18
  3024: 83330510                 ld t2, 0x100(a0)
  3028: b3836300                 add t2, t2, t1
  3032: 078b0302                 vle8.v v22, (t2)
  3036: 5705652f                 vxor.vv v10, v22, v10
  3040: d789692f                 vxor.vv v19, v22, v19
  3044: 83338510                 ld t2, 0x108(a0)
  3048: b3836300                 add t2, t2, t1
  3052: 078b0302                 vle8.v v22, (t2)
  3056: 5700602f                 vxor.vv v0, v22, v0
  3060: d789692f                 vxor.vv v19, v22, v19
  3064: 83338511                 ld t2, 0x118(a0)
  3068: b3836300                 add t2, t2, t1
  3072: 078b0302                 vle8.v v22, (t2)
  3076: 5701612f                 vxor.vv v2, v22, v2
  3080: d78a6a2f                 vxor.vv v21, v22, v21
  3084: 83330512                 ld t2, 0x120(a0)
  3088: b3836300                 add t2, t2, t1
  3092: 078b0302                 vle8.v v22, (t2)
  3096: d781612f                 vxor.vv v3, v22, v3
  3100: d785652f                 vxor.vv v11, v22, v11
  3104: 83338512                 ld t2, 0x128(a0)
  3108: b3836300                 add t2, t2, t1
  3112: 078b0302                 vle8.v v22, (t2)
  3116: 5702622f                 vxor.vv v4, v22, v4
  3120: 5706662f                 vxor.vv v12, v22, v12
  3124: 83330513                 ld t2, 0x130(a0)
  3128: b3836300                 add t2, t2, t1
  3132: 078b0302                 vle8.v v22, (t2)
  3136: d782622f                 vxor.vv v5, v22, v5
  3140: d786662f                 vxor.vv v13, v22, v13
  3144: 83338513                 ld t2, 0x138(a0)
  3148: b3836300                 add t2, t2, t1
  3152: 078b0302                 vle8.v v22, (t2)
  3156: 5707672f                 vxor.vv v14, v22, v14
  3160: 83330514                 ld t2, 0x140(a0)
  3164: b3836300                 add t2, t2, t1
  3168: 078b0302                 vle8.v v22, (t2)
  3172: d783632f                 vxor.vv v7, v22, v7
  3176: d787672f                 vxor.vv v15, v22, v15
  3180: 83338514                 ld t2, 0x148(a0)
  3184: b3836300                 add t2, t2, t1
  3188: 078b0302                 vle8.v v22, (t2)
  3192: 5704642f                 vxor.vv v8, v22, v8
  3196: 5708682f                 vxor.vv v16, v22, v16
  3200: 83330515                 ld t2, 0x150(a0)
  3204: b3836300                 add t2, t2, t1
  3208: 078b0302                 vle8.v v22, (t2)
  3212: d784642f                 vxor.vv v9, v22, v9
  3216: d788682f                 vxor.vv v17, v22, v17
  3220: 83338515                 ld t2, 0x158(a0)
  3224: b3836300                 add t2, t2, t1
  3228: 078b0302                 vle8.v v22, (t2)
  3232: 5705652f                 vxor.vv v10, v22, v10
  3236: 5709692f                 vxor.vv v18, v22, v18
  3240: 83330516                 ld t2, 0x160(a0)
  3244: b3836300                 add t2, t2, t1
  3248: 078b0302                 vle8.v v22, (t2)
  3252: 5700602f                 vxor.vv v0, v22, v0
  3256: 5709692f                 vxor.vv v18, v22, v18
  3260: 83338516                 ld t2, 0x168(a0)
  3264: b3836300                 add t2, t2, t1
  3268: 078b0302                 vle8.v v22, (t2)
  3272: d789692f                 vxor.vv v19, v22, v19
  3276: 83330517                 ld t2, 0x170(a0)
  3280: b3836300                 add t2, t2, t1
  3284: 078b0302                 vle8.v v22, (t2)
  3288: 5701612f                 vxor.vv v2, v22, v2
  3292: 570a6a2f                 vxor.vv v20, v22, v20
  3296: 83338517                 ld t2, 0x178(a0)
  3300: b3836300                 add t2, t2, t1
  3304: 078b0302                 vle8.v v22, (t2)
  3308: d781612f                 vxor.vv v3, v22, v3
  3312: d78a6a2f                 vxor.vv v21, v22, v21
  3316: 83330518                 ld t2, 0x180(a0)
  3320: b3836300                 add t2, t2, t1
  3324: 078b0302                 vle8.v v22, (t2)
  3328: 5702622f                 vxor.vv v4, v22, v4
  3332: d785652f                 vxor.vv v11, v22, v11
  3336: 83338518                 ld t2, 0x188(a0)
  3340: b3836300                 add t2, t2, t1
  3344: 078b0302                 vle8.v v22, (t2)
  3348: d782622f                 vxor.vv v5, v22, v5
  3352: 5706662f                 vxor.vv v12, v22, v12
  3356: 83330519                 ld t2, 0x190(a0)
  3360: b3836300                 add t2, t2, t1
  3364: 078b0302                 vle8.v v22, (t2)
  3368: 5703632f                 vxor.vv v6, v22, v6
  3372: d786662f                 vxor.vv v13, v22, v13
  3376: 8333051a                 ld t2, 0x1a0(a0)
  3380: b3836300                 add t2, t2, t1
  3384: 078b0302                 vle8.v v22, (t2)
  3388: 5704642f                 vxor.vv v8, v22, v8
  3392: d787672f                 vxor.vv v15, v22, v15
  3396: 8333851a                 ld t2, 0x1a8(a0)
  3400: b3836300                 add t2, t2, t1
  3404: 078b0302                 vle8.v v22, (t2)
  3408: d784642f                 vxor.vv v9, v22, v9
  3412: 5708682f                 vxor.vv v16, v22, v16
  3416: 8333051b                 ld t2, 0x1b0(a0)
  3420: b3836300                 add t2, t2, t1
  3424: 078b0302                 vle8.v v22, (t2)
  3428: 5705652f                 vxor.vv v10, v22, v10
  3432: d788682f                 vxor.vv v17, v22, v17
  3436: 8333851b                 ld t2, 0x1b8(a0)
  3440: b3836300                 add t2, t2, t1
  3444: 078b0302                 vle8.v v22, (t2)
  3448: 5700602f                 vxor.vv v0, v22, v0
  3452: d788682f                 vxor.vv v17, v22, v17
  3456: 8333051c                 ld t2, 0x1c0(a0)
  3460: b3836300                 add t2, t2, t1
  3464: 078b0302                 vle8.v v22, (t2)
  3468: d780602f                 vxor.vv v1, v22, v1
  3472: 5709692f                 vxor.vv v18, v22, v18
  3476: 8333051d                 ld t2, 0x1d0(a0)
  3480: b3836300                 add t2, t2, t1
  3484: 078b0302                 vle8.v v22, (t2)
  3488: d781612f                 vxor.vv v3, v22, v3
  3492: 570a6a2f                 vxor.vv v20, v22, v20
  3496: 8333851d                 ld t2, 0x1d8(a0)
  3500: b3836300                 add t2, t2, t1
  3504: 078b0302                 vle8.v v22, (t2)
  3508: 5702622f                 vxor.vv v4, v22, v4
  3512: d78a6a2f                 vxor.vv v21, v22, v21
  3516: 8333051e                 ld t2, 0x1e0(a0)
  3520: b3836300                 add t2, t2, t1
  3524: 078b0302                 vle8.v v22, (t2)
  3528: d782622f                 vxor.vv v5, v22, v5
  3532: d785652f                 vxor.vv v11, v22, v11
  3536: 8333851e                 ld t2, 0x1e8(a0)
  3540: b3836300                 add t2, t2, t1
  3544: 078b0302                 vle8.v v22, (t2)
  3548: 5703632f                 vxor.vv v6, v22, v6
  3552: 5706662f                 vxor.vv v12, v22, v12
  3556: 8333051f                 ld t2, 0x1f0(a0)
  3560: b3836300                 add t2, t2, t1
  3564: 078b0302                 vle8.v v22, (t2)
  3568: d786662f                 vxor.vv v13, v22, v13
  3572: 8333851f                 ld t2, 0x1f8(a0)
  3576: b3836300                 add t2, t2, t1
  3580: 078b0302                 vle8.v v22, (t2)
  3584: 5704642f                 vxor.vv v8, v22, v8
  3588: 5707672f                 vxor.vv v14, v22, v14
  3592: 83330520                 ld t2, 0x200(a0)
  3596: b3836300                 add t2, t2, t1
  3600: 078b0302                 vle8.v v22, (t2)
  3604: d784642f                 vxor.vv v9, v22, v9
  3608: d787672f                 vxor.vv v15, v22, v15
  3612: 83338520                 ld t2, 0x208(a0)
  3616: b3836300                 add t2, t2, t1
  3620: 078b0302                 vle8.v v22, (t2)
  3624: 5705652f                 vxor.vv v10, v22, v10
  3628: 5708682f                 vxor.vv v16, v22, v16
  3632: 83330521                 ld t2, 0x210(a0)
  3636: b3836300                 add t2, t2, t1
  3640: 078b0302                 vle8.v v22, (t2)
  3644: 5700602f                 vxor.vv v0, v22, v0
  3648: 5708682f                 vxor.vv v16, v22, v16
  3652: 83338521                 ld t2, 0x218(a0)
  3656: b3836300                 add t2, t2, t1
  3660: 078b0302                 vle8.v v22, (t2)
  3664: d780602f                 vxor.vv v1, v22, v1
  3668: d788682f                 vxor.vv v17, v22, v17
  3672: 83330522                 ld t2, 0x220(a0)
  3676: b3836300                 add t2, t2, t1
  3680: 078b0302                 vle8.v v22, (t2)
  3684: 5709692f                 vxor.vv v18, v22, v18
  3688: 83338522                 ld t2, 0x228(a0)
  3692: b3836300                 add t2, t2, t1
  3696: 078b0302                 vle8.v v22, (t2)
  3700: d781612f                 vxor.vv v3, v22, v3
  3704: d789692f                 vxor.vv v19, v22, v19
  3708: 83330523                 ld t2, 0x230(a0)
  3712: b3836300                 add t2, t2, t1
  3716: 078b0302                 vle8.v v22, (t2)
  3720: 5702622f                 vxor.vv v4, v22, v4
  3724: 570a6a2f                 vxor.vv v20, v22, v20
  3728: 83338523                 ld t2, 0x238(a0)
  3732: b3836300                 add t2, t2, t1
  3736: 078b0302                 vle8.v v22, (t2)
  3740: d782622f                 vxor.vv v5, v22, v5
  3744: d78a6a2f                 vxor.vv v21, v22, v21
  3748: 83330524                 ld t2, 0x240(a0)
  3752: b3836300                 add t2, t2, t1
  3756: 078b0302                 vle8.v v22, (t2)
  3760: 5703632f                 vxor.vv v6, v22, v6
  3764: d785652f                 vxor.vv v11, v22, v11
  3768: 83338524                 ld t2, 0x248(a0)
  3772: b3836300                 add t2, t2, t1
  3776: 078b0302                 vle8.v v22, (t2)
  3780: d783632f                 vxor.vv v7, v22, v7
  3784: 5706662f                 vxor.vv v12, v22, v12
  3788: 83338525                 ld t2, 0x258(a0)
  3792: b3836300                 add t2, t2, t1
  3796: 078b0302                 vle8.v v22, (t2)
  3800: d784642f                 vxor.vv v9, v22, v9
  3804: 5707672f                 vxor.vv v14, v22, v14
  3808: 83330526                 ld t2, 0x260(a0)
  3812: b3836300                 add t2, t2, t1
  3816: 078b0302                 vle8.v v22, (t2)
  3820: 5705652f                 vxor.vv v10, v22, v10
  3824: d787672f                 vxor.vv v15, v22, v15
  3828: 83338526                 ld t2, 0x268(a0)
  3832: b3836300                 add t2, t2, t1
  3836: 078b0302                 vle8.v v22, (t2)
  3840: 5700602f                 vxor.vv v0, v22, v0
  3844: d787672f                 vxor.vv v15, v22, v15
  3848: 83330527                 ld t2, 0x270(a0)
  3852: b3836300                 add t2, t2, t1
  3856: 078b0302                 vle8.v v22, (t2)
  3860: d780602f                 vxor.vv v1, v22, v1
  3864: 5708682f                 vxor.vv v16, v22, v16
  3868: 83338527                 ld t2, 0x278(a0)
  3872: b3836300                 add t2, t2, t1
  3876: 078b0302                 vle8.v v22, (t2)
  3880: 5701612f                 vxor.vv v2, v22, v2
  3884: d788682f                 vxor.vv v17, v22, v17
  3888: 83330528                 ld t2, 0x280(a0)
  3892: b3836300                 add t2, t2, t1
  3896: 078b0302                 vle8.v v22, (t2)
  3900: d781612f                 vxor.vv v3, v22, v3
  3904: 5709692f                 vxor.vv v18, v22, v18
  3908: 83338528                 ld t2, 0x288(a0)
  3912: b3836300                 add t2, t2, t1
  3916: 078b0302                 vle8.v v22, (t2)
  3920: 5702622f                 vxor.vv v4, v22, v4
  3924: d789692f                 vxor.vv v19, v22, v19
  3928: 83330529                 ld t2, 0x290(a0)
  3932: b3836300                 add t2, t2, t1
  3936: 078b0302                 vle8.v v22, (t2)
  3940: d782622f                 vxor.vv v5, v22, v5
  3944: 570a6a2f                 vxor.vv v20, v22, v20
  3948: 83338529                 ld t2, 0x298(a0)
  3952: b3836300                 add t2, t2, t1
  3956: 078b0302                 vle8.v v22, (t2)
  3960: 5703632f                 vxor.vv v6, v22, v6
  3964: d78a6a2f                 vxor.vv v21, v22, v21
  3968: 8333052a                 ld t2, 0x2a0(a0)
  3972: b3836300                 add t2, t2, t1
  3976: 078b0302                 vle8.v v22, (t2)
  3980: d783632f                 vxor.vv v7, v22, v7
  3984: d785652f                 vxor.vv v11, v22, v11
  3988: 8333852a                 ld t2, 0x2a8(a0)
  3992: b3836300                 add t2, t2, t1
  3996: 078b0302                 vle8.v v22, (t2)
  4000: 5706662f                 vxor.vv v12, v22, v12
  4004: 8333052b                 ld t2, 0x2b0(a0)
  4008: b3836300                 add t2, t2, t1
  4012: 078b0302                 vle8.v v22, (t2)
  4016: d784642f                 vxor.vv v9, v22, v9
  4020: d786662f                 vxor.vv v13, v22, v13
  4024: 8333852b                 ld t2, 0x2b8(a0)
  4028: b3836300                 add t2, t2, t1
  4032: 078b0302                 vle8.v v22, (t2)
  4036: 5705652f                 vxor.vv v10, v22, v10
  4040: 5707672f                 vxor.vv v14, v22, v14
  4044: 83b30500                 ld t2, 0(a1)
  4048: b3836300                 add t2, t2, t1
  4052: 27800302                 vse8.v v0, (t2)
  4056: 83b38500                 ld t2, 8(a1)
  4060: b3836300                 add t2, t2, t1
  4064: a7800302                 vse8.v v1, (t2)
  4068: 83b30501                 ld t2, 0x10(a1)
  4072: b3836300                 add t2, t2, t1
  4076: 27810302                 vse8.v v2, (t2)
  4080: 83b38501                 ld t2, 0x18(a1)
  4084: b3836300                 add t2, t2, t1
  4088: a7810302                 vse8.v v3, (t2)
  4092: 83b30502                 ld t2, 0x20(a1)
  4096: b3836300                 add t2, t2, t1
  4100: 27820302                 vse8.v v4, (t2)
  4104: 83b38502                 ld t2, 0x28(a1)
  4108: b3836300                 add t2, t2, t1
  4112: a7820302                 vse8.v v5, (t2)
  4116: 83b30503                 ld t2, 0x30(a1)
  4120: b3836300                 add t2, t2, t1
  4124: 27830302                 vse8.v v6, (t2)
  4128: 83b38503                 ld t2, 0x38(a1)
  4132: b3836300                 add t2, t2, t1
  4136: a7830302                 vse8.v v7, (t2)
  4140: 83b30504                 ld t2, 0x40(a1)
  4144: b3836300                 add t2, t2, t1
  4148: 27840302                 vse8.v v8, (t2)
  4152: 83b38504                 ld t2, 0x48(a1)
  4156: b3836300                 add t2, t2, t1
  4160: a7840302                 vse8.v v9, (t2)
  4164: 83b30505                 ld t2, 0x50(a1)
  4168: b3836300                 add t2, t2, t1
  4172: 27850302                 vse8.v v10, (t2)
  4176: 83b38505                 ld t2, 0x58(a1)
  4180: b3836300                 add t2, t2, t1
  4184: a7850302                 vse8.v v11, (t2)
  4188: 83b30506                 ld t2, 0x60(a1)
  4192: b3836300                 add t2, t2, t1
  4196: 27860302                 vse8.v v12, (t2)
  4200: 83b38506                 ld t2, 0x68(a1)
  4204: b3836300                 add t2, t2, t1
  4208: a7860302                 vse8.v v13, (t2)
  4212: 83b30507                 ld t2, 0x70(a1)
  4216: b3836300                 add t2, t2, t1
  4220: 27870302                 vse8.v v14, (t2)
  4224: 83b38507                 ld t2, 0x78(a1)
  4228: b3836300                 add t2, t2, t1
  4232: a7870302                 vse8.v v15, (t2)
  4236: 83b30508                 ld t2, 0x80(a1)
  4240: b3836300                 add t2, t2, t1
  4244: 27880302                 vse8.v v16, (t2)
  4248: 83b38508                 ld t2, 0x88(a1)
  4252: b3836300                 add t2, t2, t1
  4256: a7880302                 vse8.v v17, (t2)
  4260: 83b30509                 ld t2, 0x90(a1)
  4264: b3836300                 add t2, t2, t1
  4268: 27890302                 vse8.v v18, (t2)
  4272: 83b38509                 ld t2, 0x98(a1)
  4276: b3836300                 add t2, t2, t1
  4280: a7890302                 vse8.v v19, (t2)
  4284: 83b3050a                 ld t2, 0xa0(a1)
  4288: b3836300                 add t2, t2, t1
  4292: 278a0302                 vse8.v v20, (t2)
  4296: 83b3850a                 ld t2, 0xa8(a1)
  4300: b3836300                 add t2, t2, t1
  4304: a78a0302                 vse8.v v21, (t2)
  4308: 33035300                 add t1, t1, t0
  4312: 33065640                 sub a2, a2, t0
  4316: 6ff00ffb                 j 0x88c
  4320: 67800000                 ret 
  4324: 0000                     unimp 
  4326: 0000                     unimp 
  4328: 0000                     unimp 
  4330: 0000                     unimp 
  4332: 0000                     unimp 
  4334: 0000                     unimp 
