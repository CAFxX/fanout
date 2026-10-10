; config_t1_1995 target=rvv-zvbc32e twin=maskedtail topo=raid6(8)/upd{0,1} op=parity_update U=8
; code_size=3040 twins=["maskedtail"] uops_per_byte=4.0859375
     0: 63140600                 bnez a2, 8
     4: 6f00503d                 j 0xbd8
     8: 971f0000                 auipc t6, 1
    12: 938f8fbd                 addi t6, t6, -0x428
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d74f0e5e                 vmv.v.x v31, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93068000                 li a3, 8
    40: b386d202                 mul a3, t0, a3
    44: 3357d602                 divu a4, a2, a3
    48: 63140700                 bnez a4, 0x38
    52: 6f00c045                 j 0x490
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330502                 ld t2, 0x20(a0)
    68: b3836300                 add t2, t2, t1
    72: 87800302                 vle8.v v1, (t2)
    76: b3835300                 add t2, t2, t0
    80: 07810302                 vle8.v v2, (t2)
    84: b3835300                 add t2, t2, t0
    88: 87810302                 vle8.v v3, (t2)
    92: b3835300                 add t2, t2, t0
    96: 07820302                 vle8.v v4, (t2)
   100: b3835300                 add t2, t2, t0
   104: 87820302                 vle8.v v5, (t2)
   108: b3835300                 add t2, t2, t0
   112: 07830302                 vle8.v v6, (t2)
   116: b3835300                 add t2, t2, t0
   120: 87830302                 vle8.v v7, (t2)
   124: b3835300                 add t2, t2, t0
   128: 07840302                 vle8.v v8, (t2)
   132: 83338502                 ld t2, 0x28(a0)
   136: b3836300                 add t2, t2, t1
   140: 87840302                 vle8.v v9, (t2)
   144: b3835300                 add t2, t2, t0
   148: 07850302                 vle8.v v10, (t2)
   152: b3835300                 add t2, t2, t0
   156: 87850302                 vle8.v v11, (t2)
   160: b3835300                 add t2, t2, t0
   164: 07860302                 vle8.v v12, (t2)
   168: b3835300                 add t2, t2, t0
   172: 87860302                 vle8.v v13, (t2)
   176: b3835300                 add t2, t2, t0
   180: 07870302                 vle8.v v14, (t2)
   184: b3835300                 add t2, t2, t0
   188: 87870302                 vle8.v v15, (t2)
   192: b3835300                 add t2, t2, t0
   196: 07880302                 vle8.v v16, (t2)
   200: 83330500                 ld t2, 0(a0)
   204: b3836300                 add t2, t2, t1
   208: 833e0501                 ld t4, 0x10(a0)
   212: b38e6e00                 add t4, t4, t1
   216: 878c0302                 vle8.v v25, (t2)
   220: 078d0e02                 vle8.v v26, (t4)
   224: d788ac2f                 vxor.vv v17, v26, v25
   228: b3835300                 add t2, t2, t0
   232: b38e5e00                 add t4, t4, t0
   236: 878c0302                 vle8.v v25, (t2)
   240: 078d0e02                 vle8.v v26, (t4)
   244: 5789ac2f                 vxor.vv v18, v26, v25
   248: b3835300                 add t2, t2, t0
   252: b38e5e00                 add t4, t4, t0
   256: 878c0302                 vle8.v v25, (t2)
   260: 078d0e02                 vle8.v v26, (t4)
   264: d789ac2f                 vxor.vv v19, v26, v25
   268: b3835300                 add t2, t2, t0
   272: b38e5e00                 add t4, t4, t0
   276: 878c0302                 vle8.v v25, (t2)
   280: 078d0e02                 vle8.v v26, (t4)
   284: 578aac2f                 vxor.vv v20, v26, v25
   288: b3835300                 add t2, t2, t0
   292: b38e5e00                 add t4, t4, t0
   296: 878c0302                 vle8.v v25, (t2)
   300: 078d0e02                 vle8.v v26, (t4)
   304: d78aac2f                 vxor.vv v21, v26, v25
   308: b3835300                 add t2, t2, t0
   312: b38e5e00                 add t4, t4, t0
   316: 878c0302                 vle8.v v25, (t2)
   320: 078d0e02                 vle8.v v26, (t4)
   324: 578bac2f                 vxor.vv v22, v26, v25
   328: b3835300                 add t2, t2, t0
   332: b38e5e00                 add t4, t4, t0
   336: 878c0302                 vle8.v v25, (t2)
   340: 078d0e02                 vle8.v v26, (t4)
   344: d78bac2f                 vxor.vv v23, v26, v25
   348: b3835300                 add t2, t2, t0
   352: b38e5e00                 add t4, t4, t0
   356: 878c0302                 vle8.v v25, (t2)
   360: 078d0e02                 vle8.v v26, (t4)
   364: 578cac2f                 vxor.vv v24, v26, v25
   368: d780102f                 vxor.vv v1, v17, v1
   372: 5701212f                 vxor.vv v2, v18, v2
   376: d781312f                 vxor.vv v3, v19, v3
   380: 5702422f                 vxor.vv v4, v20, v4
   384: d782522f                 vxor.vv v5, v21, v5
   388: 5703632f                 vxor.vv v6, v22, v6
   392: d783732f                 vxor.vv v7, v23, v7
   396: 5704842f                 vxor.vv v8, v24, v8
   400: 138e0f00                 mv t3, t6
   404: 870e0e02                 vle8.v v29, (t3)
   408: 138e0f01                 addi t3, t6, 0x10
   412: 070f0e02                 vle8.v v30, (t3)
   416: d78cf827                 vand.vv v25, v31, v17
   420: 573d12a3                 vsrl.vi v26, v17, 4
   424: d78ddc33                 vrgather.vv v27, v29, v25
   428: 570eed33                 vrgather.vv v28, v30, v26
   432: d78dcd2f                 vxor.vv v27, v28, v27
   436: d784b42f                 vxor.vv v9, v27, v9
   440: d70cf927                 vand.vv v25, v31, v18
   444: 573d22a3                 vsrl.vi v26, v18, 4
   448: d78ddc33                 vrgather.vv v27, v29, v25
   452: 570eed33                 vrgather.vv v28, v30, v26
   456: d78dcd2f                 vxor.vv v27, v28, v27
   460: 5705b52f                 vxor.vv v10, v27, v10
   464: d78cf927                 vand.vv v25, v31, v19
   468: 573d32a3                 vsrl.vi v26, v19, 4
   472: d78ddc33                 vrgather.vv v27, v29, v25
   476: 570eed33                 vrgather.vv v28, v30, v26
   480: d78dcd2f                 vxor.vv v27, v28, v27
   484: d785b52f                 vxor.vv v11, v27, v11
   488: d70cfa27                 vand.vv v25, v31, v20
   492: 573d42a3                 vsrl.vi v26, v20, 4
   496: d78ddc33                 vrgather.vv v27, v29, v25
   500: 570eed33                 vrgather.vv v28, v30, v26
   504: d78dcd2f                 vxor.vv v27, v28, v27
   508: 5706b62f                 vxor.vv v12, v27, v12
   512: d78cfa27                 vand.vv v25, v31, v21
   516: 573d52a3                 vsrl.vi v26, v21, 4
   520: d78ddc33                 vrgather.vv v27, v29, v25
   524: 570eed33                 vrgather.vv v28, v30, v26
   528: d78dcd2f                 vxor.vv v27, v28, v27
   532: d786b62f                 vxor.vv v13, v27, v13
   536: d70cfb27                 vand.vv v25, v31, v22
   540: 573d62a3                 vsrl.vi v26, v22, 4
   544: d78ddc33                 vrgather.vv v27, v29, v25
   548: 570eed33                 vrgather.vv v28, v30, v26
   552: d78dcd2f                 vxor.vv v27, v28, v27
   556: 5707b72f                 vxor.vv v14, v27, v14
   560: d78cfb27                 vand.vv v25, v31, v23
   564: 573d72a3                 vsrl.vi v26, v23, 4
   568: d78ddc33                 vrgather.vv v27, v29, v25
   572: 570eed33                 vrgather.vv v28, v30, v26
   576: d78dcd2f                 vxor.vv v27, v28, v27
   580: d787b72f                 vxor.vv v15, v27, v15
   584: d70cfc27                 vand.vv v25, v31, v24
   588: 573d82a3                 vsrl.vi v26, v24, 4
   592: d78ddc33                 vrgather.vv v27, v29, v25
   596: 570eed33                 vrgather.vv v28, v30, v26
   600: d78dcd2f                 vxor.vv v27, v28, v27
   604: 5708b82f                 vxor.vv v16, v27, v16
   608: 83338500                 ld t2, 8(a0)
   612: b3836300                 add t2, t2, t1
   616: 833e8501                 ld t4, 0x18(a0)
   620: b38e6e00                 add t4, t4, t1
   624: 878c0302                 vle8.v v25, (t2)
   628: 078d0e02                 vle8.v v26, (t4)
   632: d788ac2f                 vxor.vv v17, v26, v25
   636: b3835300                 add t2, t2, t0
   640: b38e5e00                 add t4, t4, t0
   644: 878c0302                 vle8.v v25, (t2)
   648: 078d0e02                 vle8.v v26, (t4)
   652: 5789ac2f                 vxor.vv v18, v26, v25
   656: b3835300                 add t2, t2, t0
   660: b38e5e00                 add t4, t4, t0
   664: 878c0302                 vle8.v v25, (t2)
   668: 078d0e02                 vle8.v v26, (t4)
   672: d789ac2f                 vxor.vv v19, v26, v25
   676: b3835300                 add t2, t2, t0
   680: b38e5e00                 add t4, t4, t0
   684: 878c0302                 vle8.v v25, (t2)
   688: 078d0e02                 vle8.v v26, (t4)
   692: 578aac2f                 vxor.vv v20, v26, v25
   696: b3835300                 add t2, t2, t0
   700: b38e5e00                 add t4, t4, t0
   704: 878c0302                 vle8.v v25, (t2)
   708: 078d0e02                 vle8.v v26, (t4)
   712: d78aac2f                 vxor.vv v21, v26, v25
   716: b3835300                 add t2, t2, t0
   720: b38e5e00                 add t4, t4, t0
   724: 878c0302                 vle8.v v25, (t2)
   728: 078d0e02                 vle8.v v26, (t4)
   732: 578bac2f                 vxor.vv v22, v26, v25
   736: b3835300                 add t2, t2, t0
   740: b38e5e00                 add t4, t4, t0
   744: 878c0302                 vle8.v v25, (t2)
   748: 078d0e02                 vle8.v v26, (t4)
   752: d78bac2f                 vxor.vv v23, v26, v25
   756: b3835300                 add t2, t2, t0
   760: b38e5e00                 add t4, t4, t0
   764: 878c0302                 vle8.v v25, (t2)
   768: 078d0e02                 vle8.v v26, (t4)
   772: 578cac2f                 vxor.vv v24, v26, v25
   776: d780102f                 vxor.vv v1, v17, v1
   780: 5701212f                 vxor.vv v2, v18, v2
   784: d781312f                 vxor.vv v3, v19, v3
   788: 5702422f                 vxor.vv v4, v20, v4
   792: d782522f                 vxor.vv v5, v21, v5
   796: 5703632f                 vxor.vv v6, v22, v6
   800: d783732f                 vxor.vv v7, v23, v7
   804: 5704842f                 vxor.vv v8, v24, v8
   808: 138e0f02                 addi t3, t6, 0x20
   812: 870e0e02                 vle8.v v29, (t3)
   816: 138e0f03                 addi t3, t6, 0x30
   820: 070f0e02                 vle8.v v30, (t3)
   824: d78cf827                 vand.vv v25, v31, v17
   828: 573d12a3                 vsrl.vi v26, v17, 4
   832: d78ddc33                 vrgather.vv v27, v29, v25
   836: 570eed33                 vrgather.vv v28, v30, v26
   840: d78dcd2f                 vxor.vv v27, v28, v27
   844: d784b42f                 vxor.vv v9, v27, v9
   848: d70cf927                 vand.vv v25, v31, v18
   852: 573d22a3                 vsrl.vi v26, v18, 4
   856: d78ddc33                 vrgather.vv v27, v29, v25
   860: 570eed33                 vrgather.vv v28, v30, v26
   864: d78dcd2f                 vxor.vv v27, v28, v27
   868: 5705b52f                 vxor.vv v10, v27, v10
   872: d78cf927                 vand.vv v25, v31, v19
   876: 573d32a3                 vsrl.vi v26, v19, 4
   880: d78ddc33                 vrgather.vv v27, v29, v25
   884: 570eed33                 vrgather.vv v28, v30, v26
   888: d78dcd2f                 vxor.vv v27, v28, v27
   892: d785b52f                 vxor.vv v11, v27, v11
   896: d70cfa27                 vand.vv v25, v31, v20
   900: 573d42a3                 vsrl.vi v26, v20, 4
   904: d78ddc33                 vrgather.vv v27, v29, v25
   908: 570eed33                 vrgather.vv v28, v30, v26
   912: d78dcd2f                 vxor.vv v27, v28, v27
   916: 5706b62f                 vxor.vv v12, v27, v12
   920: d78cfa27                 vand.vv v25, v31, v21
   924: 573d52a3                 vsrl.vi v26, v21, 4
   928: d78ddc33                 vrgather.vv v27, v29, v25
   932: 570eed33                 vrgather.vv v28, v30, v26
   936: d78dcd2f                 vxor.vv v27, v28, v27
   940: d786b62f                 vxor.vv v13, v27, v13
   944: d70cfb27                 vand.vv v25, v31, v22
   948: 573d62a3                 vsrl.vi v26, v22, 4
   952: d78ddc33                 vrgather.vv v27, v29, v25
   956: 570eed33                 vrgather.vv v28, v30, v26
   960: d78dcd2f                 vxor.vv v27, v28, v27
   964: 5707b72f                 vxor.vv v14, v27, v14
   968: d78cfb27                 vand.vv v25, v31, v23
   972: 573d72a3                 vsrl.vi v26, v23, 4
   976: d78ddc33                 vrgather.vv v27, v29, v25
   980: 570eed33                 vrgather.vv v28, v30, v26
   984: d78dcd2f                 vxor.vv v27, v28, v27
   988: d787b72f                 vxor.vv v15, v27, v15
   992: d70cfc27                 vand.vv v25, v31, v24
   996: 573d82a3                 vsrl.vi v26, v24, 4
  1000: d78ddc33                 vrgather.vv v27, v29, v25
  1004: 570eed33                 vrgather.vv v28, v30, v26
  1008: d78dcd2f                 vxor.vv v27, v28, v27
  1012: 5708b82f                 vxor.vv v16, v27, v16
  1016: 83b30500                 ld t2, 0(a1)
  1020: b3836300                 add t2, t2, t1
  1024: a7800302                 vse8.v v1, (t2)
  1028: b3835300                 add t2, t2, t0
  1032: 27810302                 vse8.v v2, (t2)
  1036: b3835300                 add t2, t2, t0
  1040: a7810302                 vse8.v v3, (t2)
  1044: b3835300                 add t2, t2, t0
  1048: 27820302                 vse8.v v4, (t2)
  1052: b3835300                 add t2, t2, t0
  1056: a7820302                 vse8.v v5, (t2)
  1060: b3835300                 add t2, t2, t0
  1064: 27830302                 vse8.v v6, (t2)
  1068: b3835300                 add t2, t2, t0
  1072: a7830302                 vse8.v v7, (t2)
  1076: b3835300                 add t2, t2, t0
  1080: 27840302                 vse8.v v8, (t2)
  1084: 83b38500                 ld t2, 8(a1)
  1088: b3836300                 add t2, t2, t1
  1092: a7840302                 vse8.v v9, (t2)
  1096: b3835300                 add t2, t2, t0
  1100: 27850302                 vse8.v v10, (t2)
  1104: b3835300                 add t2, t2, t0
  1108: a7850302                 vse8.v v11, (t2)
  1112: b3835300                 add t2, t2, t0
  1116: 27860302                 vse8.v v12, (t2)
  1120: b3835300                 add t2, t2, t0
  1124: a7860302                 vse8.v v13, (t2)
  1128: b3835300                 add t2, t2, t0
  1132: 27870302                 vse8.v v14, (t2)
  1136: b3835300                 add t2, t2, t0
  1140: a7870302                 vse8.v v15, (t2)
  1144: b3835300                 add t2, t2, t0
  1148: 27880302                 vse8.v v16, (t2)
  1152: 3303d300                 add t1, t1, a3
  1156: 1307f7ff                 addi a4, a4, -1
  1160: e31c07ba                 bnez a4, 0x40
  1164: 33066640                 sub a2, a2, t1
  1168: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
  1172: 93060600                 mv a3, a2
  1176: 6344d000                 bgtz a3, 0x4a0
  1180: 6f00c00d                 j 0x578
  1184: b3835640                 sub t2, a3, t0
  1188: 13def343                 srai t3, t2, 0x3f
  1192: b3f3c301                 and t2, t2, t3
  1196: b3837200                 add t2, t0, t2
  1200: d7a80852                 vid.v v17
  1204: 57c0136b                 vmsltu.vx v0, v17, t2
  1208: 83330502                 ld t2, 0x20(a0)
  1212: b3836300                 add t2, t2, t1
  1216: 87800300                 vle8.v v1, (t2), v0.t
  1220: 83338502                 ld t2, 0x28(a0)
  1224: b3836300                 add t2, t2, t1
  1228: 87840300                 vle8.v v9, (t2), v0.t
  1232: 83330500                 ld t2, 0(a0)
  1236: b3836300                 add t2, t2, t1
  1240: 833e0501                 ld t4, 0x10(a0)
  1244: b38e6e00                 add t4, t4, t1
  1248: 878c0300                 vle8.v v25, (t2), v0.t
  1252: 078d0e00                 vle8.v v26, (t4), v0.t
  1256: d788ac2f                 vxor.vv v17, v26, v25
  1260: d780102f                 vxor.vv v1, v17, v1
  1264: 138e0f00                 mv t3, t6
  1268: 870e0e02                 vle8.v v29, (t3)
  1272: 138e0f01                 addi t3, t6, 0x10
  1276: 070f0e02                 vle8.v v30, (t3)
  1280: d78cf827                 vand.vv v25, v31, v17
  1284: 573d12a3                 vsrl.vi v26, v17, 4
  1288: d78ddc33                 vrgather.vv v27, v29, v25
  1292: 570eed33                 vrgather.vv v28, v30, v26
  1296: d78dcd2f                 vxor.vv v27, v28, v27
  1300: d784b42f                 vxor.vv v9, v27, v9
  1304: 83338500                 ld t2, 8(a0)
  1308: b3836300                 add t2, t2, t1
  1312: 833e8501                 ld t4, 0x18(a0)
  1316: b38e6e00                 add t4, t4, t1
  1320: 878c0300                 vle8.v v25, (t2), v0.t
  1324: 078d0e00                 vle8.v v26, (t4), v0.t
  1328: d788ac2f                 vxor.vv v17, v26, v25
  1332: d780102f                 vxor.vv v1, v17, v1
  1336: 138e0f02                 addi t3, t6, 0x20
  1340: 870e0e02                 vle8.v v29, (t3)
  1344: 138e0f03                 addi t3, t6, 0x30
  1348: 070f0e02                 vle8.v v30, (t3)
  1352: d78cf827                 vand.vv v25, v31, v17
  1356: 573d12a3                 vsrl.vi v26, v17, 4
  1360: d78ddc33                 vrgather.vv v27, v29, v25
  1364: 570eed33                 vrgather.vv v28, v30, v26
  1368: d78dcd2f                 vxor.vv v27, v28, v27
  1372: d784b42f                 vxor.vv v9, v27, v9
  1376: 83b30500                 ld t2, 0(a1)
  1380: b3836300                 add t2, t2, t1
  1384: a7800300                 vse8.v v1, (t2), v0.t
  1388: 83b38500                 ld t2, 8(a1)
  1392: b3836300                 add t2, t2, t1
  1396: a7840300                 vse8.v v9, (t2), v0.t
  1400: 33035300                 add t1, t1, t0
  1404: b3865640                 sub a3, a3, t0
  1408: 6344d000                 bgtz a3, 0x588
  1412: 6f00c00d                 j 0x660
  1416: b3835640                 sub t2, a3, t0
  1420: 13def343                 srai t3, t2, 0x3f
  1424: b3f3c301                 and t2, t2, t3
  1428: b3837200                 add t2, t0, t2
  1432: d7a80852                 vid.v v17
  1436: 57c0136b                 vmsltu.vx v0, v17, t2
  1440: 83330502                 ld t2, 0x20(a0)
  1444: b3836300                 add t2, t2, t1
  1448: 87800300                 vle8.v v1, (t2), v0.t
  1452: 83338502                 ld t2, 0x28(a0)
  1456: b3836300                 add t2, t2, t1
  1460: 87840300                 vle8.v v9, (t2), v0.t
  1464: 83330500                 ld t2, 0(a0)
  1468: b3836300                 add t2, t2, t1
  1472: 833e0501                 ld t4, 0x10(a0)
  1476: b38e6e00                 add t4, t4, t1
  1480: 878c0300                 vle8.v v25, (t2), v0.t
  1484: 078d0e00                 vle8.v v26, (t4), v0.t
  1488: d788ac2f                 vxor.vv v17, v26, v25
  1492: d780102f                 vxor.vv v1, v17, v1
  1496: 138e0f00                 mv t3, t6
  1500: 870e0e02                 vle8.v v29, (t3)
  1504: 138e0f01                 addi t3, t6, 0x10
  1508: 070f0e02                 vle8.v v30, (t3)
  1512: d78cf827                 vand.vv v25, v31, v17
  1516: 573d12a3                 vsrl.vi v26, v17, 4
  1520: d78ddc33                 vrgather.vv v27, v29, v25
  1524: 570eed33                 vrgather.vv v28, v30, v26
  1528: d78dcd2f                 vxor.vv v27, v28, v27
  1532: d784b42f                 vxor.vv v9, v27, v9
  1536: 83338500                 ld t2, 8(a0)
  1540: b3836300                 add t2, t2, t1
  1544: 833e8501                 ld t4, 0x18(a0)
  1548: b38e6e00                 add t4, t4, t1
  1552: 878c0300                 vle8.v v25, (t2), v0.t
  1556: 078d0e00                 vle8.v v26, (t4), v0.t
  1560: d788ac2f                 vxor.vv v17, v26, v25
  1564: d780102f                 vxor.vv v1, v17, v1
  1568: 138e0f02                 addi t3, t6, 0x20
  1572: 870e0e02                 vle8.v v29, (t3)
  1576: 138e0f03                 addi t3, t6, 0x30
  1580: 070f0e02                 vle8.v v30, (t3)
  1584: d78cf827                 vand.vv v25, v31, v17
  1588: 573d12a3                 vsrl.vi v26, v17, 4
  1592: d78ddc33                 vrgather.vv v27, v29, v25
  1596: 570eed33                 vrgather.vv v28, v30, v26
  1600: d78dcd2f                 vxor.vv v27, v28, v27
  1604: d784b42f                 vxor.vv v9, v27, v9
  1608: 83b30500                 ld t2, 0(a1)
  1612: b3836300                 add t2, t2, t1
  1616: a7800300                 vse8.v v1, (t2), v0.t
  1620: 83b38500                 ld t2, 8(a1)
  1624: b3836300                 add t2, t2, t1
  1628: a7840300                 vse8.v v9, (t2), v0.t
  1632: 33035300                 add t1, t1, t0
  1636: b3865640                 sub a3, a3, t0
  1640: 6344d000                 bgtz a3, 0x670
  1644: 6f00c00d                 j 0x748
  1648: b3835640                 sub t2, a3, t0
  1652: 13def343                 srai t3, t2, 0x3f
  1656: b3f3c301                 and t2, t2, t3
  1660: b3837200                 add t2, t0, t2
  1664: d7a80852                 vid.v v17
  1668: 57c0136b                 vmsltu.vx v0, v17, t2
  1672: 83330502                 ld t2, 0x20(a0)
  1676: b3836300                 add t2, t2, t1
  1680: 87800300                 vle8.v v1, (t2), v0.t
  1684: 83338502                 ld t2, 0x28(a0)
  1688: b3836300                 add t2, t2, t1
  1692: 87840300                 vle8.v v9, (t2), v0.t
  1696: 83330500                 ld t2, 0(a0)
  1700: b3836300                 add t2, t2, t1
  1704: 833e0501                 ld t4, 0x10(a0)
  1708: b38e6e00                 add t4, t4, t1
  1712: 878c0300                 vle8.v v25, (t2), v0.t
  1716: 078d0e00                 vle8.v v26, (t4), v0.t
  1720: d788ac2f                 vxor.vv v17, v26, v25
  1724: d780102f                 vxor.vv v1, v17, v1
  1728: 138e0f00                 mv t3, t6
  1732: 870e0e02                 vle8.v v29, (t3)
  1736: 138e0f01                 addi t3, t6, 0x10
  1740: 070f0e02                 vle8.v v30, (t3)
  1744: d78cf827                 vand.vv v25, v31, v17
  1748: 573d12a3                 vsrl.vi v26, v17, 4
  1752: d78ddc33                 vrgather.vv v27, v29, v25
  1756: 570eed33                 vrgather.vv v28, v30, v26
  1760: d78dcd2f                 vxor.vv v27, v28, v27
  1764: d784b42f                 vxor.vv v9, v27, v9
  1768: 83338500                 ld t2, 8(a0)
  1772: b3836300                 add t2, t2, t1
  1776: 833e8501                 ld t4, 0x18(a0)
  1780: b38e6e00                 add t4, t4, t1
  1784: 878c0300                 vle8.v v25, (t2), v0.t
  1788: 078d0e00                 vle8.v v26, (t4), v0.t
  1792: d788ac2f                 vxor.vv v17, v26, v25
  1796: d780102f                 vxor.vv v1, v17, v1
  1800: 138e0f02                 addi t3, t6, 0x20
  1804: 870e0e02                 vle8.v v29, (t3)
  1808: 138e0f03                 addi t3, t6, 0x30
  1812: 070f0e02                 vle8.v v30, (t3)
  1816: d78cf827                 vand.vv v25, v31, v17
  1820: 573d12a3                 vsrl.vi v26, v17, 4
  1824: d78ddc33                 vrgather.vv v27, v29, v25
  1828: 570eed33                 vrgather.vv v28, v30, v26
  1832: d78dcd2f                 vxor.vv v27, v28, v27
  1836: d784b42f                 vxor.vv v9, v27, v9
  1840: 83b30500                 ld t2, 0(a1)
  1844: b3836300                 add t2, t2, t1
  1848: a7800300                 vse8.v v1, (t2), v0.t
  1852: 83b38500                 ld t2, 8(a1)
  1856: b3836300                 add t2, t2, t1
  1860: a7840300                 vse8.v v9, (t2), v0.t
  1864: 33035300                 add t1, t1, t0
  1868: b3865640                 sub a3, a3, t0
  1872: 6344d000                 bgtz a3, 0x758
  1876: 6f00c00d                 j 0x830
  1880: b3835640                 sub t2, a3, t0
  1884: 13def343                 srai t3, t2, 0x3f
  1888: b3f3c301                 and t2, t2, t3
  1892: b3837200                 add t2, t0, t2
  1896: d7a80852                 vid.v v17
  1900: 57c0136b                 vmsltu.vx v0, v17, t2
  1904: 83330502                 ld t2, 0x20(a0)
  1908: b3836300                 add t2, t2, t1
  1912: 87800300                 vle8.v v1, (t2), v0.t
  1916: 83338502                 ld t2, 0x28(a0)
  1920: b3836300                 add t2, t2, t1
  1924: 87840300                 vle8.v v9, (t2), v0.t
  1928: 83330500                 ld t2, 0(a0)
  1932: b3836300                 add t2, t2, t1
  1936: 833e0501                 ld t4, 0x10(a0)
  1940: b38e6e00                 add t4, t4, t1
  1944: 878c0300                 vle8.v v25, (t2), v0.t
  1948: 078d0e00                 vle8.v v26, (t4), v0.t
  1952: d788ac2f                 vxor.vv v17, v26, v25
  1956: d780102f                 vxor.vv v1, v17, v1
  1960: 138e0f00                 mv t3, t6
  1964: 870e0e02                 vle8.v v29, (t3)
  1968: 138e0f01                 addi t3, t6, 0x10
  1972: 070f0e02                 vle8.v v30, (t3)
  1976: d78cf827                 vand.vv v25, v31, v17
  1980: 573d12a3                 vsrl.vi v26, v17, 4
  1984: d78ddc33                 vrgather.vv v27, v29, v25
  1988: 570eed33                 vrgather.vv v28, v30, v26
  1992: d78dcd2f                 vxor.vv v27, v28, v27
  1996: d784b42f                 vxor.vv v9, v27, v9
  2000: 83338500                 ld t2, 8(a0)
  2004: b3836300                 add t2, t2, t1
  2008: 833e8501                 ld t4, 0x18(a0)
  2012: b38e6e00                 add t4, t4, t1
  2016: 878c0300                 vle8.v v25, (t2), v0.t
  2020: 078d0e00                 vle8.v v26, (t4), v0.t
  2024: d788ac2f                 vxor.vv v17, v26, v25
  2028: d780102f                 vxor.vv v1, v17, v1
  2032: 138e0f02                 addi t3, t6, 0x20
  2036: 870e0e02                 vle8.v v29, (t3)
  2040: 138e0f03                 addi t3, t6, 0x30
  2044: 070f0e02                 vle8.v v30, (t3)
  2048: d78cf827                 vand.vv v25, v31, v17
  2052: 573d12a3                 vsrl.vi v26, v17, 4
  2056: d78ddc33                 vrgather.vv v27, v29, v25
  2060: 570eed33                 vrgather.vv v28, v30, v26
  2064: d78dcd2f                 vxor.vv v27, v28, v27
  2068: d784b42f                 vxor.vv v9, v27, v9
  2072: 83b30500                 ld t2, 0(a1)
  2076: b3836300                 add t2, t2, t1
  2080: a7800300                 vse8.v v1, (t2), v0.t
  2084: 83b38500                 ld t2, 8(a1)
  2088: b3836300                 add t2, t2, t1
  2092: a7840300                 vse8.v v9, (t2), v0.t
  2096: 33035300                 add t1, t1, t0
  2100: b3865640                 sub a3, a3, t0
  2104: 6344d000                 bgtz a3, 0x840
  2108: 6f00c00d                 j 0x918
  2112: b3835640                 sub t2, a3, t0
  2116: 13def343                 srai t3, t2, 0x3f
  2120: b3f3c301                 and t2, t2, t3
  2124: b3837200                 add t2, t0, t2
  2128: d7a80852                 vid.v v17
  2132: 57c0136b                 vmsltu.vx v0, v17, t2
  2136: 83330502                 ld t2, 0x20(a0)
  2140: b3836300                 add t2, t2, t1
  2144: 87800300                 vle8.v v1, (t2), v0.t
  2148: 83338502                 ld t2, 0x28(a0)
  2152: b3836300                 add t2, t2, t1
  2156: 87840300                 vle8.v v9, (t2), v0.t
  2160: 83330500                 ld t2, 0(a0)
  2164: b3836300                 add t2, t2, t1
  2168: 833e0501                 ld t4, 0x10(a0)
  2172: b38e6e00                 add t4, t4, t1
  2176: 878c0300                 vle8.v v25, (t2), v0.t
  2180: 078d0e00                 vle8.v v26, (t4), v0.t
  2184: d788ac2f                 vxor.vv v17, v26, v25
  2188: d780102f                 vxor.vv v1, v17, v1
  2192: 138e0f00                 mv t3, t6
  2196: 870e0e02                 vle8.v v29, (t3)
  2200: 138e0f01                 addi t3, t6, 0x10
  2204: 070f0e02                 vle8.v v30, (t3)
  2208: d78cf827                 vand.vv v25, v31, v17
  2212: 573d12a3                 vsrl.vi v26, v17, 4
  2216: d78ddc33                 vrgather.vv v27, v29, v25
  2220: 570eed33                 vrgather.vv v28, v30, v26
  2224: d78dcd2f                 vxor.vv v27, v28, v27
  2228: d784b42f                 vxor.vv v9, v27, v9
  2232: 83338500                 ld t2, 8(a0)
  2236: b3836300                 add t2, t2, t1
  2240: 833e8501                 ld t4, 0x18(a0)
  2244: b38e6e00                 add t4, t4, t1
  2248: 878c0300                 vle8.v v25, (t2), v0.t
  2252: 078d0e00                 vle8.v v26, (t4), v0.t
  2256: d788ac2f                 vxor.vv v17, v26, v25
  2260: d780102f                 vxor.vv v1, v17, v1
  2264: 138e0f02                 addi t3, t6, 0x20
  2268: 870e0e02                 vle8.v v29, (t3)
  2272: 138e0f03                 addi t3, t6, 0x30
  2276: 070f0e02                 vle8.v v30, (t3)
  2280: d78cf827                 vand.vv v25, v31, v17
  2284: 573d12a3                 vsrl.vi v26, v17, 4
  2288: d78ddc33                 vrgather.vv v27, v29, v25
  2292: 570eed33                 vrgather.vv v28, v30, v26
  2296: d78dcd2f                 vxor.vv v27, v28, v27
  2300: d784b42f                 vxor.vv v9, v27, v9
  2304: 83b30500                 ld t2, 0(a1)
  2308: b3836300                 add t2, t2, t1
  2312: a7800300                 vse8.v v1, (t2), v0.t
  2316: 83b38500                 ld t2, 8(a1)
  2320: b3836300                 add t2, t2, t1
  2324: a7840300                 vse8.v v9, (t2), v0.t
  2328: 33035300                 add t1, t1, t0
  2332: b3865640                 sub a3, a3, t0
  2336: 6344d000                 bgtz a3, 0x928
  2340: 6f00c00d                 j 0xa00
  2344: b3835640                 sub t2, a3, t0
  2348: 13def343                 srai t3, t2, 0x3f
  2352: b3f3c301                 and t2, t2, t3
  2356: b3837200                 add t2, t0, t2
  2360: d7a80852                 vid.v v17
  2364: 57c0136b                 vmsltu.vx v0, v17, t2
  2368: 83330502                 ld t2, 0x20(a0)
  2372: b3836300                 add t2, t2, t1
  2376: 87800300                 vle8.v v1, (t2), v0.t
  2380: 83338502                 ld t2, 0x28(a0)
  2384: b3836300                 add t2, t2, t1
  2388: 87840300                 vle8.v v9, (t2), v0.t
  2392: 83330500                 ld t2, 0(a0)
  2396: b3836300                 add t2, t2, t1
  2400: 833e0501                 ld t4, 0x10(a0)
  2404: b38e6e00                 add t4, t4, t1
  2408: 878c0300                 vle8.v v25, (t2), v0.t
  2412: 078d0e00                 vle8.v v26, (t4), v0.t
  2416: d788ac2f                 vxor.vv v17, v26, v25
  2420: d780102f                 vxor.vv v1, v17, v1
  2424: 138e0f00                 mv t3, t6
  2428: 870e0e02                 vle8.v v29, (t3)
  2432: 138e0f01                 addi t3, t6, 0x10
  2436: 070f0e02                 vle8.v v30, (t3)
  2440: d78cf827                 vand.vv v25, v31, v17
  2444: 573d12a3                 vsrl.vi v26, v17, 4
  2448: d78ddc33                 vrgather.vv v27, v29, v25
  2452: 570eed33                 vrgather.vv v28, v30, v26
  2456: d78dcd2f                 vxor.vv v27, v28, v27
  2460: d784b42f                 vxor.vv v9, v27, v9
  2464: 83338500                 ld t2, 8(a0)
  2468: b3836300                 add t2, t2, t1
  2472: 833e8501                 ld t4, 0x18(a0)
  2476: b38e6e00                 add t4, t4, t1
  2480: 878c0300                 vle8.v v25, (t2), v0.t
  2484: 078d0e00                 vle8.v v26, (t4), v0.t
  2488: d788ac2f                 vxor.vv v17, v26, v25
  2492: d780102f                 vxor.vv v1, v17, v1
  2496: 138e0f02                 addi t3, t6, 0x20
  2500: 870e0e02                 vle8.v v29, (t3)
  2504: 138e0f03                 addi t3, t6, 0x30
  2508: 070f0e02                 vle8.v v30, (t3)
  2512: d78cf827                 vand.vv v25, v31, v17
  2516: 573d12a3                 vsrl.vi v26, v17, 4
  2520: d78ddc33                 vrgather.vv v27, v29, v25
  2524: 570eed33                 vrgather.vv v28, v30, v26
  2528: d78dcd2f                 vxor.vv v27, v28, v27
  2532: d784b42f                 vxor.vv v9, v27, v9
  2536: 83b30500                 ld t2, 0(a1)
  2540: b3836300                 add t2, t2, t1
  2544: a7800300                 vse8.v v1, (t2), v0.t
  2548: 83b38500                 ld t2, 8(a1)
  2552: b3836300                 add t2, t2, t1
  2556: a7840300                 vse8.v v9, (t2), v0.t
  2560: 33035300                 add t1, t1, t0
  2564: b3865640                 sub a3, a3, t0
  2568: 6344d000                 bgtz a3, 0xa10
  2572: 6f00c00d                 j 0xae8
  2576: b3835640                 sub t2, a3, t0
  2580: 13def343                 srai t3, t2, 0x3f
  2584: b3f3c301                 and t2, t2, t3
  2588: b3837200                 add t2, t0, t2
  2592: d7a80852                 vid.v v17
  2596: 57c0136b                 vmsltu.vx v0, v17, t2
  2600: 83330502                 ld t2, 0x20(a0)
  2604: b3836300                 add t2, t2, t1
  2608: 87800300                 vle8.v v1, (t2), v0.t
  2612: 83338502                 ld t2, 0x28(a0)
  2616: b3836300                 add t2, t2, t1
  2620: 87840300                 vle8.v v9, (t2), v0.t
  2624: 83330500                 ld t2, 0(a0)
  2628: b3836300                 add t2, t2, t1
  2632: 833e0501                 ld t4, 0x10(a0)
  2636: b38e6e00                 add t4, t4, t1
  2640: 878c0300                 vle8.v v25, (t2), v0.t
  2644: 078d0e00                 vle8.v v26, (t4), v0.t
  2648: d788ac2f                 vxor.vv v17, v26, v25
  2652: d780102f                 vxor.vv v1, v17, v1
  2656: 138e0f00                 mv t3, t6
  2660: 870e0e02                 vle8.v v29, (t3)
  2664: 138e0f01                 addi t3, t6, 0x10
  2668: 070f0e02                 vle8.v v30, (t3)
  2672: d78cf827                 vand.vv v25, v31, v17
  2676: 573d12a3                 vsrl.vi v26, v17, 4
  2680: d78ddc33                 vrgather.vv v27, v29, v25
  2684: 570eed33                 vrgather.vv v28, v30, v26
  2688: d78dcd2f                 vxor.vv v27, v28, v27
  2692: d784b42f                 vxor.vv v9, v27, v9
  2696: 83338500                 ld t2, 8(a0)
  2700: b3836300                 add t2, t2, t1
  2704: 833e8501                 ld t4, 0x18(a0)
  2708: b38e6e00                 add t4, t4, t1
  2712: 878c0300                 vle8.v v25, (t2), v0.t
  2716: 078d0e00                 vle8.v v26, (t4), v0.t
  2720: d788ac2f                 vxor.vv v17, v26, v25
  2724: d780102f                 vxor.vv v1, v17, v1
  2728: 138e0f02                 addi t3, t6, 0x20
  2732: 870e0e02                 vle8.v v29, (t3)
  2736: 138e0f03                 addi t3, t6, 0x30
  2740: 070f0e02                 vle8.v v30, (t3)
  2744: d78cf827                 vand.vv v25, v31, v17
  2748: 573d12a3                 vsrl.vi v26, v17, 4
  2752: d78ddc33                 vrgather.vv v27, v29, v25
  2756: 570eed33                 vrgather.vv v28, v30, v26
  2760: d78dcd2f                 vxor.vv v27, v28, v27
  2764: d784b42f                 vxor.vv v9, v27, v9
  2768: 83b30500                 ld t2, 0(a1)
  2772: b3836300                 add t2, t2, t1
  2776: a7800300                 vse8.v v1, (t2), v0.t
  2780: 83b38500                 ld t2, 8(a1)
  2784: b3836300                 add t2, t2, t1
  2788: a7840300                 vse8.v v9, (t2), v0.t
  2792: 33035300                 add t1, t1, t0
  2796: b3865640                 sub a3, a3, t0
  2800: 6344d000                 bgtz a3, 0xaf8
  2804: 6f00c00d                 j 0xbd0
  2808: b3835640                 sub t2, a3, t0
  2812: 13def343                 srai t3, t2, 0x3f
  2816: b3f3c301                 and t2, t2, t3
  2820: b3837200                 add t2, t0, t2
  2824: d7a80852                 vid.v v17
  2828: 57c0136b                 vmsltu.vx v0, v17, t2
  2832: 83330502                 ld t2, 0x20(a0)
  2836: b3836300                 add t2, t2, t1
  2840: 87800300                 vle8.v v1, (t2), v0.t
  2844: 83338502                 ld t2, 0x28(a0)
  2848: b3836300                 add t2, t2, t1
  2852: 87840300                 vle8.v v9, (t2), v0.t
  2856: 83330500                 ld t2, 0(a0)
  2860: b3836300                 add t2, t2, t1
  2864: 833e0501                 ld t4, 0x10(a0)
  2868: b38e6e00                 add t4, t4, t1
  2872: 878c0300                 vle8.v v25, (t2), v0.t
  2876: 078d0e00                 vle8.v v26, (t4), v0.t
  2880: d788ac2f                 vxor.vv v17, v26, v25
  2884: d780102f                 vxor.vv v1, v17, v1
  2888: 138e0f00                 mv t3, t6
  2892: 870e0e02                 vle8.v v29, (t3)
  2896: 138e0f01                 addi t3, t6, 0x10
  2900: 070f0e02                 vle8.v v30, (t3)
  2904: d78cf827                 vand.vv v25, v31, v17
  2908: 573d12a3                 vsrl.vi v26, v17, 4
  2912: d78ddc33                 vrgather.vv v27, v29, v25
  2916: 570eed33                 vrgather.vv v28, v30, v26
  2920: d78dcd2f                 vxor.vv v27, v28, v27
  2924: d784b42f                 vxor.vv v9, v27, v9
  2928: 83338500                 ld t2, 8(a0)
  2932: b3836300                 add t2, t2, t1
  2936: 833e8501                 ld t4, 0x18(a0)
  2940: b38e6e00                 add t4, t4, t1
  2944: 878c0300                 vle8.v v25, (t2), v0.t
  2948: 078d0e00                 vle8.v v26, (t4), v0.t
  2952: d788ac2f                 vxor.vv v17, v26, v25
  2956: d780102f                 vxor.vv v1, v17, v1
  2960: 138e0f02                 addi t3, t6, 0x20
  2964: 870e0e02                 vle8.v v29, (t3)
  2968: 138e0f03                 addi t3, t6, 0x30
  2972: 070f0e02                 vle8.v v30, (t3)
  2976: d78cf827                 vand.vv v25, v31, v17
  2980: 573d12a3                 vsrl.vi v26, v17, 4
  2984: d78ddc33                 vrgather.vv v27, v29, v25
  2988: 570eed33                 vrgather.vv v28, v30, v26
  2992: d78dcd2f                 vxor.vv v27, v28, v27
  2996: d784b42f                 vxor.vv v9, v27, v9
  3000: 83b30500                 ld t2, 0(a1)
  3004: b3836300                 add t2, t2, t1
  3008: a7800300                 vse8.v v1, (t2), v0.t
  3012: 83b38500                 ld t2, 8(a1)
  3016: b3836300                 add t2, t2, t1
  3020: a7840300                 vse8.v v9, (t2), v0.t
  3024: 33035300                 add t1, t1, t0
  3028: b3865640                 sub a3, a3, t0
  3032: 67800000                 ret 
  3036: 0000                     unimp 
  3038: 0000                     unimp 
