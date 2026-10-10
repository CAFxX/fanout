; config_t1_1631 target=rvv twin=maskedtail topo=rs(6,3)/rec{0,1} op=recover U=8
; code_size=7920 twins=["maskedtail"] uops_per_byte=13.6171875
     0: 63140600                 bnez a2, 8
     4: 6f10506e                 j 0x1ee8
     8: 972f0000                 auipc t6, 2
    12: 938f8fee                 addi t6, t6, -0x118
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d74f0e5e                 vmv.v.x v31, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93068000                 li a3, 8
    40: b386d202                 mul a3, t0, a3
    44: 3357d602                 divu a4, a2, a3
    48: 63140700                 bnez a4, 0x38
    52: 6f00d03e                 j 0xc20
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
   132: 138e0f00                 mv t3, t6
   136: 870e0e02                 vle8.v v29, (t3)
   140: 138e0f01                 addi t3, t6, 0x10
   144: 070f0e02                 vle8.v v30, (t3)
   148: d78cf827                 vand.vv v25, v31, v17
   152: 573d12a3                 vsrl.vi v26, v17, 4
   156: d78ddc33                 vrgather.vv v27, v29, v25
   160: 570eed33                 vrgather.vv v28, v30, v26
   164: d78dcd2f                 vxor.vv v27, v28, v27
   168: d7800d5e                 vmv.v.v v1, v27
   172: d70cf927                 vand.vv v25, v31, v18
   176: 573d22a3                 vsrl.vi v26, v18, 4
   180: d78ddc33                 vrgather.vv v27, v29, v25
   184: 570eed33                 vrgather.vv v28, v30, v26
   188: d78dcd2f                 vxor.vv v27, v28, v27
   192: 57810d5e                 vmv.v.v v2, v27
   196: d78cf927                 vand.vv v25, v31, v19
   200: 573d32a3                 vsrl.vi v26, v19, 4
   204: d78ddc33                 vrgather.vv v27, v29, v25
   208: 570eed33                 vrgather.vv v28, v30, v26
   212: d78dcd2f                 vxor.vv v27, v28, v27
   216: d7810d5e                 vmv.v.v v3, v27
   220: d70cfa27                 vand.vv v25, v31, v20
   224: 573d42a3                 vsrl.vi v26, v20, 4
   228: d78ddc33                 vrgather.vv v27, v29, v25
   232: 570eed33                 vrgather.vv v28, v30, v26
   236: d78dcd2f                 vxor.vv v27, v28, v27
   240: 57820d5e                 vmv.v.v v4, v27
   244: d78cfa27                 vand.vv v25, v31, v21
   248: 573d52a3                 vsrl.vi v26, v21, 4
   252: d78ddc33                 vrgather.vv v27, v29, v25
   256: 570eed33                 vrgather.vv v28, v30, v26
   260: d78dcd2f                 vxor.vv v27, v28, v27
   264: d7820d5e                 vmv.v.v v5, v27
   268: d70cfb27                 vand.vv v25, v31, v22
   272: 573d62a3                 vsrl.vi v26, v22, 4
   276: d78ddc33                 vrgather.vv v27, v29, v25
   280: 570eed33                 vrgather.vv v28, v30, v26
   284: d78dcd2f                 vxor.vv v27, v28, v27
   288: 57830d5e                 vmv.v.v v6, v27
   292: d78cfb27                 vand.vv v25, v31, v23
   296: 573d72a3                 vsrl.vi v26, v23, 4
   300: d78ddc33                 vrgather.vv v27, v29, v25
   304: 570eed33                 vrgather.vv v28, v30, v26
   308: d78dcd2f                 vxor.vv v27, v28, v27
   312: d7830d5e                 vmv.v.v v7, v27
   316: d70cfc27                 vand.vv v25, v31, v24
   320: 573d82a3                 vsrl.vi v26, v24, 4
   324: d78ddc33                 vrgather.vv v27, v29, v25
   328: 570eed33                 vrgather.vv v28, v30, v26
   332: d78dcd2f                 vxor.vv v27, v28, v27
   336: 57840d5e                 vmv.v.v v8, v27
   340: 138e0f02                 addi t3, t6, 0x20
   344: 870e0e02                 vle8.v v29, (t3)
   348: 138e0f03                 addi t3, t6, 0x30
   352: 070f0e02                 vle8.v v30, (t3)
   356: d78cf827                 vand.vv v25, v31, v17
   360: 573d12a3                 vsrl.vi v26, v17, 4
   364: d78ddc33                 vrgather.vv v27, v29, v25
   368: 570eed33                 vrgather.vv v28, v30, v26
   372: d78dcd2f                 vxor.vv v27, v28, v27
   376: d7840d5e                 vmv.v.v v9, v27
   380: d70cf927                 vand.vv v25, v31, v18
   384: 573d22a3                 vsrl.vi v26, v18, 4
   388: d78ddc33                 vrgather.vv v27, v29, v25
   392: 570eed33                 vrgather.vv v28, v30, v26
   396: d78dcd2f                 vxor.vv v27, v28, v27
   400: 57850d5e                 vmv.v.v v10, v27
   404: d78cf927                 vand.vv v25, v31, v19
   408: 573d32a3                 vsrl.vi v26, v19, 4
   412: d78ddc33                 vrgather.vv v27, v29, v25
   416: 570eed33                 vrgather.vv v28, v30, v26
   420: d78dcd2f                 vxor.vv v27, v28, v27
   424: d7850d5e                 vmv.v.v v11, v27
   428: d70cfa27                 vand.vv v25, v31, v20
   432: 573d42a3                 vsrl.vi v26, v20, 4
   436: d78ddc33                 vrgather.vv v27, v29, v25
   440: 570eed33                 vrgather.vv v28, v30, v26
   444: d78dcd2f                 vxor.vv v27, v28, v27
   448: 57860d5e                 vmv.v.v v12, v27
   452: d78cfa27                 vand.vv v25, v31, v21
   456: 573d52a3                 vsrl.vi v26, v21, 4
   460: d78ddc33                 vrgather.vv v27, v29, v25
   464: 570eed33                 vrgather.vv v28, v30, v26
   468: d78dcd2f                 vxor.vv v27, v28, v27
   472: d7860d5e                 vmv.v.v v13, v27
   476: d70cfb27                 vand.vv v25, v31, v22
   480: 573d62a3                 vsrl.vi v26, v22, 4
   484: d78ddc33                 vrgather.vv v27, v29, v25
   488: 570eed33                 vrgather.vv v28, v30, v26
   492: d78dcd2f                 vxor.vv v27, v28, v27
   496: 57870d5e                 vmv.v.v v14, v27
   500: d78cfb27                 vand.vv v25, v31, v23
   504: 573d72a3                 vsrl.vi v26, v23, 4
   508: d78ddc33                 vrgather.vv v27, v29, v25
   512: 570eed33                 vrgather.vv v28, v30, v26
   516: d78dcd2f                 vxor.vv v27, v28, v27
   520: d7870d5e                 vmv.v.v v15, v27
   524: d70cfc27                 vand.vv v25, v31, v24
   528: 573d82a3                 vsrl.vi v26, v24, 4
   532: d78ddc33                 vrgather.vv v27, v29, v25
   536: 570eed33                 vrgather.vv v28, v30, v26
   540: d78dcd2f                 vxor.vv v27, v28, v27
   544: 57880d5e                 vmv.v.v v16, v27
   548: 83338500                 ld t2, 8(a0)
   552: b3836300                 add t2, t2, t1
   556: 87880302                 vle8.v v17, (t2)
   560: b3835300                 add t2, t2, t0
   564: 07890302                 vle8.v v18, (t2)
   568: b3835300                 add t2, t2, t0
   572: 87890302                 vle8.v v19, (t2)
   576: b3835300                 add t2, t2, t0
   580: 078a0302                 vle8.v v20, (t2)
   584: b3835300                 add t2, t2, t0
   588: 878a0302                 vle8.v v21, (t2)
   592: b3835300                 add t2, t2, t0
   596: 078b0302                 vle8.v v22, (t2)
   600: b3835300                 add t2, t2, t0
   604: 878b0302                 vle8.v v23, (t2)
   608: b3835300                 add t2, t2, t0
   612: 078c0302                 vle8.v v24, (t2)
   616: 138e0f04                 addi t3, t6, 0x40
   620: 870e0e02                 vle8.v v29, (t3)
   624: 138e0f05                 addi t3, t6, 0x50
   628: 070f0e02                 vle8.v v30, (t3)
   632: d78cf827                 vand.vv v25, v31, v17
   636: 573d12a3                 vsrl.vi v26, v17, 4
   640: d78ddc33                 vrgather.vv v27, v29, v25
   644: 570eed33                 vrgather.vv v28, v30, v26
   648: d78dcd2f                 vxor.vv v27, v28, v27
   652: d780b02f                 vxor.vv v1, v27, v1
   656: d70cf927                 vand.vv v25, v31, v18
   660: 573d22a3                 vsrl.vi v26, v18, 4
   664: d78ddc33                 vrgather.vv v27, v29, v25
   668: 570eed33                 vrgather.vv v28, v30, v26
   672: d78dcd2f                 vxor.vv v27, v28, v27
   676: 5701b12f                 vxor.vv v2, v27, v2
   680: d78cf927                 vand.vv v25, v31, v19
   684: 573d32a3                 vsrl.vi v26, v19, 4
   688: d78ddc33                 vrgather.vv v27, v29, v25
   692: 570eed33                 vrgather.vv v28, v30, v26
   696: d78dcd2f                 vxor.vv v27, v28, v27
   700: d781b12f                 vxor.vv v3, v27, v3
   704: d70cfa27                 vand.vv v25, v31, v20
   708: 573d42a3                 vsrl.vi v26, v20, 4
   712: d78ddc33                 vrgather.vv v27, v29, v25
   716: 570eed33                 vrgather.vv v28, v30, v26
   720: d78dcd2f                 vxor.vv v27, v28, v27
   724: 5702b22f                 vxor.vv v4, v27, v4
   728: d78cfa27                 vand.vv v25, v31, v21
   732: 573d52a3                 vsrl.vi v26, v21, 4
   736: d78ddc33                 vrgather.vv v27, v29, v25
   740: 570eed33                 vrgather.vv v28, v30, v26
   744: d78dcd2f                 vxor.vv v27, v28, v27
   748: d782b22f                 vxor.vv v5, v27, v5
   752: d70cfb27                 vand.vv v25, v31, v22
   756: 573d62a3                 vsrl.vi v26, v22, 4
   760: d78ddc33                 vrgather.vv v27, v29, v25
   764: 570eed33                 vrgather.vv v28, v30, v26
   768: d78dcd2f                 vxor.vv v27, v28, v27
   772: 5703b32f                 vxor.vv v6, v27, v6
   776: d78cfb27                 vand.vv v25, v31, v23
   780: 573d72a3                 vsrl.vi v26, v23, 4
   784: d78ddc33                 vrgather.vv v27, v29, v25
   788: 570eed33                 vrgather.vv v28, v30, v26
   792: d78dcd2f                 vxor.vv v27, v28, v27
   796: d783b32f                 vxor.vv v7, v27, v7
   800: d70cfc27                 vand.vv v25, v31, v24
   804: 573d82a3                 vsrl.vi v26, v24, 4
   808: d78ddc33                 vrgather.vv v27, v29, v25
   812: 570eed33                 vrgather.vv v28, v30, v26
   816: d78dcd2f                 vxor.vv v27, v28, v27
   820: 5704b42f                 vxor.vv v8, v27, v8
   824: 138e0f06                 addi t3, t6, 0x60
   828: 870e0e02                 vle8.v v29, (t3)
   832: 138e0f07                 addi t3, t6, 0x70
   836: 070f0e02                 vle8.v v30, (t3)
   840: d78cf827                 vand.vv v25, v31, v17
   844: 573d12a3                 vsrl.vi v26, v17, 4
   848: d78ddc33                 vrgather.vv v27, v29, v25
   852: 570eed33                 vrgather.vv v28, v30, v26
   856: d78dcd2f                 vxor.vv v27, v28, v27
   860: d784b42f                 vxor.vv v9, v27, v9
   864: d70cf927                 vand.vv v25, v31, v18
   868: 573d22a3                 vsrl.vi v26, v18, 4
   872: d78ddc33                 vrgather.vv v27, v29, v25
   876: 570eed33                 vrgather.vv v28, v30, v26
   880: d78dcd2f                 vxor.vv v27, v28, v27
   884: 5705b52f                 vxor.vv v10, v27, v10
   888: d78cf927                 vand.vv v25, v31, v19
   892: 573d32a3                 vsrl.vi v26, v19, 4
   896: d78ddc33                 vrgather.vv v27, v29, v25
   900: 570eed33                 vrgather.vv v28, v30, v26
   904: d78dcd2f                 vxor.vv v27, v28, v27
   908: d785b52f                 vxor.vv v11, v27, v11
   912: d70cfa27                 vand.vv v25, v31, v20
   916: 573d42a3                 vsrl.vi v26, v20, 4
   920: d78ddc33                 vrgather.vv v27, v29, v25
   924: 570eed33                 vrgather.vv v28, v30, v26
   928: d78dcd2f                 vxor.vv v27, v28, v27
   932: 5706b62f                 vxor.vv v12, v27, v12
   936: d78cfa27                 vand.vv v25, v31, v21
   940: 573d52a3                 vsrl.vi v26, v21, 4
   944: d78ddc33                 vrgather.vv v27, v29, v25
   948: 570eed33                 vrgather.vv v28, v30, v26
   952: d78dcd2f                 vxor.vv v27, v28, v27
   956: d786b62f                 vxor.vv v13, v27, v13
   960: d70cfb27                 vand.vv v25, v31, v22
   964: 573d62a3                 vsrl.vi v26, v22, 4
   968: d78ddc33                 vrgather.vv v27, v29, v25
   972: 570eed33                 vrgather.vv v28, v30, v26
   976: d78dcd2f                 vxor.vv v27, v28, v27
   980: 5707b72f                 vxor.vv v14, v27, v14
   984: d78cfb27                 vand.vv v25, v31, v23
   988: 573d72a3                 vsrl.vi v26, v23, 4
   992: d78ddc33                 vrgather.vv v27, v29, v25
   996: 570eed33                 vrgather.vv v28, v30, v26
  1000: d78dcd2f                 vxor.vv v27, v28, v27
  1004: d787b72f                 vxor.vv v15, v27, v15
  1008: d70cfc27                 vand.vv v25, v31, v24
  1012: 573d82a3                 vsrl.vi v26, v24, 4
  1016: d78ddc33                 vrgather.vv v27, v29, v25
  1020: 570eed33                 vrgather.vv v28, v30, v26
  1024: d78dcd2f                 vxor.vv v27, v28, v27
  1028: 5708b82f                 vxor.vv v16, v27, v16
  1032: 83330501                 ld t2, 0x10(a0)
  1036: b3836300                 add t2, t2, t1
  1040: 87880302                 vle8.v v17, (t2)
  1044: b3835300                 add t2, t2, t0
  1048: 07890302                 vle8.v v18, (t2)
  1052: b3835300                 add t2, t2, t0
  1056: 87890302                 vle8.v v19, (t2)
  1060: b3835300                 add t2, t2, t0
  1064: 078a0302                 vle8.v v20, (t2)
  1068: b3835300                 add t2, t2, t0
  1072: 878a0302                 vle8.v v21, (t2)
  1076: b3835300                 add t2, t2, t0
  1080: 078b0302                 vle8.v v22, (t2)
  1084: b3835300                 add t2, t2, t0
  1088: 878b0302                 vle8.v v23, (t2)
  1092: b3835300                 add t2, t2, t0
  1096: 078c0302                 vle8.v v24, (t2)
  1100: 138e0f08                 addi t3, t6, 0x80
  1104: 870e0e02                 vle8.v v29, (t3)
  1108: 138e0f09                 addi t3, t6, 0x90
  1112: 070f0e02                 vle8.v v30, (t3)
  1116: d78cf827                 vand.vv v25, v31, v17
  1120: 573d12a3                 vsrl.vi v26, v17, 4
  1124: d78ddc33                 vrgather.vv v27, v29, v25
  1128: 570eed33                 vrgather.vv v28, v30, v26
  1132: d78dcd2f                 vxor.vv v27, v28, v27
  1136: d780b02f                 vxor.vv v1, v27, v1
  1140: d70cf927                 vand.vv v25, v31, v18
  1144: 573d22a3                 vsrl.vi v26, v18, 4
  1148: d78ddc33                 vrgather.vv v27, v29, v25
  1152: 570eed33                 vrgather.vv v28, v30, v26
  1156: d78dcd2f                 vxor.vv v27, v28, v27
  1160: 5701b12f                 vxor.vv v2, v27, v2
  1164: d78cf927                 vand.vv v25, v31, v19
  1168: 573d32a3                 vsrl.vi v26, v19, 4
  1172: d78ddc33                 vrgather.vv v27, v29, v25
  1176: 570eed33                 vrgather.vv v28, v30, v26
  1180: d78dcd2f                 vxor.vv v27, v28, v27
  1184: d781b12f                 vxor.vv v3, v27, v3
  1188: d70cfa27                 vand.vv v25, v31, v20
  1192: 573d42a3                 vsrl.vi v26, v20, 4
  1196: d78ddc33                 vrgather.vv v27, v29, v25
  1200: 570eed33                 vrgather.vv v28, v30, v26
  1204: d78dcd2f                 vxor.vv v27, v28, v27
  1208: 5702b22f                 vxor.vv v4, v27, v4
  1212: d78cfa27                 vand.vv v25, v31, v21
  1216: 573d52a3                 vsrl.vi v26, v21, 4
  1220: d78ddc33                 vrgather.vv v27, v29, v25
  1224: 570eed33                 vrgather.vv v28, v30, v26
  1228: d78dcd2f                 vxor.vv v27, v28, v27
  1232: d782b22f                 vxor.vv v5, v27, v5
  1236: d70cfb27                 vand.vv v25, v31, v22
  1240: 573d62a3                 vsrl.vi v26, v22, 4
  1244: d78ddc33                 vrgather.vv v27, v29, v25
  1248: 570eed33                 vrgather.vv v28, v30, v26
  1252: d78dcd2f                 vxor.vv v27, v28, v27
  1256: 5703b32f                 vxor.vv v6, v27, v6
  1260: d78cfb27                 vand.vv v25, v31, v23
  1264: 573d72a3                 vsrl.vi v26, v23, 4
  1268: d78ddc33                 vrgather.vv v27, v29, v25
  1272: 570eed33                 vrgather.vv v28, v30, v26
  1276: d78dcd2f                 vxor.vv v27, v28, v27
  1280: d783b32f                 vxor.vv v7, v27, v7
  1284: d70cfc27                 vand.vv v25, v31, v24
  1288: 573d82a3                 vsrl.vi v26, v24, 4
  1292: d78ddc33                 vrgather.vv v27, v29, v25
  1296: 570eed33                 vrgather.vv v28, v30, v26
  1300: d78dcd2f                 vxor.vv v27, v28, v27
  1304: 5704b42f                 vxor.vv v8, v27, v8
  1308: 138e0f0a                 addi t3, t6, 0xa0
  1312: 870e0e02                 vle8.v v29, (t3)
  1316: 138e0f0b                 addi t3, t6, 0xb0
  1320: 070f0e02                 vle8.v v30, (t3)
  1324: d78cf827                 vand.vv v25, v31, v17
  1328: 573d12a3                 vsrl.vi v26, v17, 4
  1332: d78ddc33                 vrgather.vv v27, v29, v25
  1336: 570eed33                 vrgather.vv v28, v30, v26
  1340: d78dcd2f                 vxor.vv v27, v28, v27
  1344: d784b42f                 vxor.vv v9, v27, v9
  1348: d70cf927                 vand.vv v25, v31, v18
  1352: 573d22a3                 vsrl.vi v26, v18, 4
  1356: d78ddc33                 vrgather.vv v27, v29, v25
  1360: 570eed33                 vrgather.vv v28, v30, v26
  1364: d78dcd2f                 vxor.vv v27, v28, v27
  1368: 5705b52f                 vxor.vv v10, v27, v10
  1372: d78cf927                 vand.vv v25, v31, v19
  1376: 573d32a3                 vsrl.vi v26, v19, 4
  1380: d78ddc33                 vrgather.vv v27, v29, v25
  1384: 570eed33                 vrgather.vv v28, v30, v26
  1388: d78dcd2f                 vxor.vv v27, v28, v27
  1392: d785b52f                 vxor.vv v11, v27, v11
  1396: d70cfa27                 vand.vv v25, v31, v20
  1400: 573d42a3                 vsrl.vi v26, v20, 4
  1404: d78ddc33                 vrgather.vv v27, v29, v25
  1408: 570eed33                 vrgather.vv v28, v30, v26
  1412: d78dcd2f                 vxor.vv v27, v28, v27
  1416: 5706b62f                 vxor.vv v12, v27, v12
  1420: d78cfa27                 vand.vv v25, v31, v21
  1424: 573d52a3                 vsrl.vi v26, v21, 4
  1428: d78ddc33                 vrgather.vv v27, v29, v25
  1432: 570eed33                 vrgather.vv v28, v30, v26
  1436: d78dcd2f                 vxor.vv v27, v28, v27
  1440: d786b62f                 vxor.vv v13, v27, v13
  1444: d70cfb27                 vand.vv v25, v31, v22
  1448: 573d62a3                 vsrl.vi v26, v22, 4
  1452: d78ddc33                 vrgather.vv v27, v29, v25
  1456: 570eed33                 vrgather.vv v28, v30, v26
  1460: d78dcd2f                 vxor.vv v27, v28, v27
  1464: 5707b72f                 vxor.vv v14, v27, v14
  1468: d78cfb27                 vand.vv v25, v31, v23
  1472: 573d72a3                 vsrl.vi v26, v23, 4
  1476: d78ddc33                 vrgather.vv v27, v29, v25
  1480: 570eed33                 vrgather.vv v28, v30, v26
  1484: d78dcd2f                 vxor.vv v27, v28, v27
  1488: d787b72f                 vxor.vv v15, v27, v15
  1492: d70cfc27                 vand.vv v25, v31, v24
  1496: 573d82a3                 vsrl.vi v26, v24, 4
  1500: d78ddc33                 vrgather.vv v27, v29, v25
  1504: 570eed33                 vrgather.vv v28, v30, v26
  1508: d78dcd2f                 vxor.vv v27, v28, v27
  1512: 5708b82f                 vxor.vv v16, v27, v16
  1516: 83338501                 ld t2, 0x18(a0)
  1520: b3836300                 add t2, t2, t1
  1524: 87880302                 vle8.v v17, (t2)
  1528: b3835300                 add t2, t2, t0
  1532: 07890302                 vle8.v v18, (t2)
  1536: b3835300                 add t2, t2, t0
  1540: 87890302                 vle8.v v19, (t2)
  1544: b3835300                 add t2, t2, t0
  1548: 078a0302                 vle8.v v20, (t2)
  1552: b3835300                 add t2, t2, t0
  1556: 878a0302                 vle8.v v21, (t2)
  1560: b3835300                 add t2, t2, t0
  1564: 078b0302                 vle8.v v22, (t2)
  1568: b3835300                 add t2, t2, t0
  1572: 878b0302                 vle8.v v23, (t2)
  1576: b3835300                 add t2, t2, t0
  1580: 078c0302                 vle8.v v24, (t2)
  1584: 138e0f0c                 addi t3, t6, 0xc0
  1588: 870e0e02                 vle8.v v29, (t3)
  1592: 138e0f0d                 addi t3, t6, 0xd0
  1596: 070f0e02                 vle8.v v30, (t3)
  1600: d78cf827                 vand.vv v25, v31, v17
  1604: 573d12a3                 vsrl.vi v26, v17, 4
  1608: d78ddc33                 vrgather.vv v27, v29, v25
  1612: 570eed33                 vrgather.vv v28, v30, v26
  1616: d78dcd2f                 vxor.vv v27, v28, v27
  1620: d780b02f                 vxor.vv v1, v27, v1
  1624: d70cf927                 vand.vv v25, v31, v18
  1628: 573d22a3                 vsrl.vi v26, v18, 4
  1632: d78ddc33                 vrgather.vv v27, v29, v25
  1636: 570eed33                 vrgather.vv v28, v30, v26
  1640: d78dcd2f                 vxor.vv v27, v28, v27
  1644: 5701b12f                 vxor.vv v2, v27, v2
  1648: d78cf927                 vand.vv v25, v31, v19
  1652: 573d32a3                 vsrl.vi v26, v19, 4
  1656: d78ddc33                 vrgather.vv v27, v29, v25
  1660: 570eed33                 vrgather.vv v28, v30, v26
  1664: d78dcd2f                 vxor.vv v27, v28, v27
  1668: d781b12f                 vxor.vv v3, v27, v3
  1672: d70cfa27                 vand.vv v25, v31, v20
  1676: 573d42a3                 vsrl.vi v26, v20, 4
  1680: d78ddc33                 vrgather.vv v27, v29, v25
  1684: 570eed33                 vrgather.vv v28, v30, v26
  1688: d78dcd2f                 vxor.vv v27, v28, v27
  1692: 5702b22f                 vxor.vv v4, v27, v4
  1696: d78cfa27                 vand.vv v25, v31, v21
  1700: 573d52a3                 vsrl.vi v26, v21, 4
  1704: d78ddc33                 vrgather.vv v27, v29, v25
  1708: 570eed33                 vrgather.vv v28, v30, v26
  1712: d78dcd2f                 vxor.vv v27, v28, v27
  1716: d782b22f                 vxor.vv v5, v27, v5
  1720: d70cfb27                 vand.vv v25, v31, v22
  1724: 573d62a3                 vsrl.vi v26, v22, 4
  1728: d78ddc33                 vrgather.vv v27, v29, v25
  1732: 570eed33                 vrgather.vv v28, v30, v26
  1736: d78dcd2f                 vxor.vv v27, v28, v27
  1740: 5703b32f                 vxor.vv v6, v27, v6
  1744: d78cfb27                 vand.vv v25, v31, v23
  1748: 573d72a3                 vsrl.vi v26, v23, 4
  1752: d78ddc33                 vrgather.vv v27, v29, v25
  1756: 570eed33                 vrgather.vv v28, v30, v26
  1760: d78dcd2f                 vxor.vv v27, v28, v27
  1764: d783b32f                 vxor.vv v7, v27, v7
  1768: d70cfc27                 vand.vv v25, v31, v24
  1772: 573d82a3                 vsrl.vi v26, v24, 4
  1776: d78ddc33                 vrgather.vv v27, v29, v25
  1780: 570eed33                 vrgather.vv v28, v30, v26
  1784: d78dcd2f                 vxor.vv v27, v28, v27
  1788: 5704b42f                 vxor.vv v8, v27, v8
  1792: 138e0f0e                 addi t3, t6, 0xe0
  1796: 870e0e02                 vle8.v v29, (t3)
  1800: 138e0f0f                 addi t3, t6, 0xf0
  1804: 070f0e02                 vle8.v v30, (t3)
  1808: d78cf827                 vand.vv v25, v31, v17
  1812: 573d12a3                 vsrl.vi v26, v17, 4
  1816: d78ddc33                 vrgather.vv v27, v29, v25
  1820: 570eed33                 vrgather.vv v28, v30, v26
  1824: d78dcd2f                 vxor.vv v27, v28, v27
  1828: d784b42f                 vxor.vv v9, v27, v9
  1832: d70cf927                 vand.vv v25, v31, v18
  1836: 573d22a3                 vsrl.vi v26, v18, 4
  1840: d78ddc33                 vrgather.vv v27, v29, v25
  1844: 570eed33                 vrgather.vv v28, v30, v26
  1848: d78dcd2f                 vxor.vv v27, v28, v27
  1852: 5705b52f                 vxor.vv v10, v27, v10
  1856: d78cf927                 vand.vv v25, v31, v19
  1860: 573d32a3                 vsrl.vi v26, v19, 4
  1864: d78ddc33                 vrgather.vv v27, v29, v25
  1868: 570eed33                 vrgather.vv v28, v30, v26
  1872: d78dcd2f                 vxor.vv v27, v28, v27
  1876: d785b52f                 vxor.vv v11, v27, v11
  1880: d70cfa27                 vand.vv v25, v31, v20
  1884: 573d42a3                 vsrl.vi v26, v20, 4
  1888: d78ddc33                 vrgather.vv v27, v29, v25
  1892: 570eed33                 vrgather.vv v28, v30, v26
  1896: d78dcd2f                 vxor.vv v27, v28, v27
  1900: 5706b62f                 vxor.vv v12, v27, v12
  1904: d78cfa27                 vand.vv v25, v31, v21
  1908: 573d52a3                 vsrl.vi v26, v21, 4
  1912: d78ddc33                 vrgather.vv v27, v29, v25
  1916: 570eed33                 vrgather.vv v28, v30, v26
  1920: d78dcd2f                 vxor.vv v27, v28, v27
  1924: d786b62f                 vxor.vv v13, v27, v13
  1928: d70cfb27                 vand.vv v25, v31, v22
  1932: 573d62a3                 vsrl.vi v26, v22, 4
  1936: d78ddc33                 vrgather.vv v27, v29, v25
  1940: 570eed33                 vrgather.vv v28, v30, v26
  1944: d78dcd2f                 vxor.vv v27, v28, v27
  1948: 5707b72f                 vxor.vv v14, v27, v14
  1952: d78cfb27                 vand.vv v25, v31, v23
  1956: 573d72a3                 vsrl.vi v26, v23, 4
  1960: d78ddc33                 vrgather.vv v27, v29, v25
  1964: 570eed33                 vrgather.vv v28, v30, v26
  1968: d78dcd2f                 vxor.vv v27, v28, v27
  1972: d787b72f                 vxor.vv v15, v27, v15
  1976: d70cfc27                 vand.vv v25, v31, v24
  1980: 573d82a3                 vsrl.vi v26, v24, 4
  1984: d78ddc33                 vrgather.vv v27, v29, v25
  1988: 570eed33                 vrgather.vv v28, v30, v26
  1992: d78dcd2f                 vxor.vv v27, v28, v27
  1996: 5708b82f                 vxor.vv v16, v27, v16
  2000: 83330502                 ld t2, 0x20(a0)
  2004: b3836300                 add t2, t2, t1
  2008: 87880302                 vle8.v v17, (t2)
  2012: b3835300                 add t2, t2, t0
  2016: 07890302                 vle8.v v18, (t2)
  2020: b3835300                 add t2, t2, t0
  2024: 87890302                 vle8.v v19, (t2)
  2028: b3835300                 add t2, t2, t0
  2032: 078a0302                 vle8.v v20, (t2)
  2036: b3835300                 add t2, t2, t0
  2040: 878a0302                 vle8.v v21, (t2)
  2044: b3835300                 add t2, t2, t0
  2048: 078b0302                 vle8.v v22, (t2)
  2052: b3835300                 add t2, t2, t0
  2056: 878b0302                 vle8.v v23, (t2)
  2060: b3835300                 add t2, t2, t0
  2064: 078c0302                 vle8.v v24, (t2)
  2068: 138e0f10                 addi t3, t6, 0x100
  2072: 870e0e02                 vle8.v v29, (t3)
  2076: 138e0f11                 addi t3, t6, 0x110
  2080: 070f0e02                 vle8.v v30, (t3)
  2084: d78cf827                 vand.vv v25, v31, v17
  2088: 573d12a3                 vsrl.vi v26, v17, 4
  2092: d78ddc33                 vrgather.vv v27, v29, v25
  2096: 570eed33                 vrgather.vv v28, v30, v26
  2100: d78dcd2f                 vxor.vv v27, v28, v27
  2104: d780b02f                 vxor.vv v1, v27, v1
  2108: d70cf927                 vand.vv v25, v31, v18
  2112: 573d22a3                 vsrl.vi v26, v18, 4
  2116: d78ddc33                 vrgather.vv v27, v29, v25
  2120: 570eed33                 vrgather.vv v28, v30, v26
  2124: d78dcd2f                 vxor.vv v27, v28, v27
  2128: 5701b12f                 vxor.vv v2, v27, v2
  2132: d78cf927                 vand.vv v25, v31, v19
  2136: 573d32a3                 vsrl.vi v26, v19, 4
  2140: d78ddc33                 vrgather.vv v27, v29, v25
  2144: 570eed33                 vrgather.vv v28, v30, v26
  2148: d78dcd2f                 vxor.vv v27, v28, v27
  2152: d781b12f                 vxor.vv v3, v27, v3
  2156: d70cfa27                 vand.vv v25, v31, v20
  2160: 573d42a3                 vsrl.vi v26, v20, 4
  2164: d78ddc33                 vrgather.vv v27, v29, v25
  2168: 570eed33                 vrgather.vv v28, v30, v26
  2172: d78dcd2f                 vxor.vv v27, v28, v27
  2176: 5702b22f                 vxor.vv v4, v27, v4
  2180: d78cfa27                 vand.vv v25, v31, v21
  2184: 573d52a3                 vsrl.vi v26, v21, 4
  2188: d78ddc33                 vrgather.vv v27, v29, v25
  2192: 570eed33                 vrgather.vv v28, v30, v26
  2196: d78dcd2f                 vxor.vv v27, v28, v27
  2200: d782b22f                 vxor.vv v5, v27, v5
  2204: d70cfb27                 vand.vv v25, v31, v22
  2208: 573d62a3                 vsrl.vi v26, v22, 4
  2212: d78ddc33                 vrgather.vv v27, v29, v25
  2216: 570eed33                 vrgather.vv v28, v30, v26
  2220: d78dcd2f                 vxor.vv v27, v28, v27
  2224: 5703b32f                 vxor.vv v6, v27, v6
  2228: d78cfb27                 vand.vv v25, v31, v23
  2232: 573d72a3                 vsrl.vi v26, v23, 4
  2236: d78ddc33                 vrgather.vv v27, v29, v25
  2240: 570eed33                 vrgather.vv v28, v30, v26
  2244: d78dcd2f                 vxor.vv v27, v28, v27
  2248: d783b32f                 vxor.vv v7, v27, v7
  2252: d70cfc27                 vand.vv v25, v31, v24
  2256: 573d82a3                 vsrl.vi v26, v24, 4
  2260: d78ddc33                 vrgather.vv v27, v29, v25
  2264: 570eed33                 vrgather.vv v28, v30, v26
  2268: d78dcd2f                 vxor.vv v27, v28, v27
  2272: 5704b42f                 vxor.vv v8, v27, v8
  2276: 138e0f12                 addi t3, t6, 0x120
  2280: 870e0e02                 vle8.v v29, (t3)
  2284: 138e0f13                 addi t3, t6, 0x130
  2288: 070f0e02                 vle8.v v30, (t3)
  2292: d78cf827                 vand.vv v25, v31, v17
  2296: 573d12a3                 vsrl.vi v26, v17, 4
  2300: d78ddc33                 vrgather.vv v27, v29, v25
  2304: 570eed33                 vrgather.vv v28, v30, v26
  2308: d78dcd2f                 vxor.vv v27, v28, v27
  2312: d784b42f                 vxor.vv v9, v27, v9
  2316: d70cf927                 vand.vv v25, v31, v18
  2320: 573d22a3                 vsrl.vi v26, v18, 4
  2324: d78ddc33                 vrgather.vv v27, v29, v25
  2328: 570eed33                 vrgather.vv v28, v30, v26
  2332: d78dcd2f                 vxor.vv v27, v28, v27
  2336: 5705b52f                 vxor.vv v10, v27, v10
  2340: d78cf927                 vand.vv v25, v31, v19
  2344: 573d32a3                 vsrl.vi v26, v19, 4
  2348: d78ddc33                 vrgather.vv v27, v29, v25
  2352: 570eed33                 vrgather.vv v28, v30, v26
  2356: d78dcd2f                 vxor.vv v27, v28, v27
  2360: d785b52f                 vxor.vv v11, v27, v11
  2364: d70cfa27                 vand.vv v25, v31, v20
  2368: 573d42a3                 vsrl.vi v26, v20, 4
  2372: d78ddc33                 vrgather.vv v27, v29, v25
  2376: 570eed33                 vrgather.vv v28, v30, v26
  2380: d78dcd2f                 vxor.vv v27, v28, v27
  2384: 5706b62f                 vxor.vv v12, v27, v12
  2388: d78cfa27                 vand.vv v25, v31, v21
  2392: 573d52a3                 vsrl.vi v26, v21, 4
  2396: d78ddc33                 vrgather.vv v27, v29, v25
  2400: 570eed33                 vrgather.vv v28, v30, v26
  2404: d78dcd2f                 vxor.vv v27, v28, v27
  2408: d786b62f                 vxor.vv v13, v27, v13
  2412: d70cfb27                 vand.vv v25, v31, v22
  2416: 573d62a3                 vsrl.vi v26, v22, 4
  2420: d78ddc33                 vrgather.vv v27, v29, v25
  2424: 570eed33                 vrgather.vv v28, v30, v26
  2428: d78dcd2f                 vxor.vv v27, v28, v27
  2432: 5707b72f                 vxor.vv v14, v27, v14
  2436: d78cfb27                 vand.vv v25, v31, v23
  2440: 573d72a3                 vsrl.vi v26, v23, 4
  2444: d78ddc33                 vrgather.vv v27, v29, v25
  2448: 570eed33                 vrgather.vv v28, v30, v26
  2452: d78dcd2f                 vxor.vv v27, v28, v27
  2456: d787b72f                 vxor.vv v15, v27, v15
  2460: d70cfc27                 vand.vv v25, v31, v24
  2464: 573d82a3                 vsrl.vi v26, v24, 4
  2468: d78ddc33                 vrgather.vv v27, v29, v25
  2472: 570eed33                 vrgather.vv v28, v30, v26
  2476: d78dcd2f                 vxor.vv v27, v28, v27
  2480: 5708b82f                 vxor.vv v16, v27, v16
  2484: 83338502                 ld t2, 0x28(a0)
  2488: b3836300                 add t2, t2, t1
  2492: 87880302                 vle8.v v17, (t2)
  2496: b3835300                 add t2, t2, t0
  2500: 07890302                 vle8.v v18, (t2)
  2504: b3835300                 add t2, t2, t0
  2508: 87890302                 vle8.v v19, (t2)
  2512: b3835300                 add t2, t2, t0
  2516: 078a0302                 vle8.v v20, (t2)
  2520: b3835300                 add t2, t2, t0
  2524: 878a0302                 vle8.v v21, (t2)
  2528: b3835300                 add t2, t2, t0
  2532: 078b0302                 vle8.v v22, (t2)
  2536: b3835300                 add t2, t2, t0
  2540: 878b0302                 vle8.v v23, (t2)
  2544: b3835300                 add t2, t2, t0
  2548: 078c0302                 vle8.v v24, (t2)
  2552: 138e0f14                 addi t3, t6, 0x140
  2556: 870e0e02                 vle8.v v29, (t3)
  2560: 138e0f15                 addi t3, t6, 0x150
  2564: 070f0e02                 vle8.v v30, (t3)
  2568: d78cf827                 vand.vv v25, v31, v17
  2572: 573d12a3                 vsrl.vi v26, v17, 4
  2576: d78ddc33                 vrgather.vv v27, v29, v25
  2580: 570eed33                 vrgather.vv v28, v30, v26
  2584: d78dcd2f                 vxor.vv v27, v28, v27
  2588: d780b02f                 vxor.vv v1, v27, v1
  2592: d70cf927                 vand.vv v25, v31, v18
  2596: 573d22a3                 vsrl.vi v26, v18, 4
  2600: d78ddc33                 vrgather.vv v27, v29, v25
  2604: 570eed33                 vrgather.vv v28, v30, v26
  2608: d78dcd2f                 vxor.vv v27, v28, v27
  2612: 5701b12f                 vxor.vv v2, v27, v2
  2616: d78cf927                 vand.vv v25, v31, v19
  2620: 573d32a3                 vsrl.vi v26, v19, 4
  2624: d78ddc33                 vrgather.vv v27, v29, v25
  2628: 570eed33                 vrgather.vv v28, v30, v26
  2632: d78dcd2f                 vxor.vv v27, v28, v27
  2636: d781b12f                 vxor.vv v3, v27, v3
  2640: d70cfa27                 vand.vv v25, v31, v20
  2644: 573d42a3                 vsrl.vi v26, v20, 4
  2648: d78ddc33                 vrgather.vv v27, v29, v25
  2652: 570eed33                 vrgather.vv v28, v30, v26
  2656: d78dcd2f                 vxor.vv v27, v28, v27
  2660: 5702b22f                 vxor.vv v4, v27, v4
  2664: d78cfa27                 vand.vv v25, v31, v21
  2668: 573d52a3                 vsrl.vi v26, v21, 4
  2672: d78ddc33                 vrgather.vv v27, v29, v25
  2676: 570eed33                 vrgather.vv v28, v30, v26
  2680: d78dcd2f                 vxor.vv v27, v28, v27
  2684: d782b22f                 vxor.vv v5, v27, v5
  2688: d70cfb27                 vand.vv v25, v31, v22
  2692: 573d62a3                 vsrl.vi v26, v22, 4
  2696: d78ddc33                 vrgather.vv v27, v29, v25
  2700: 570eed33                 vrgather.vv v28, v30, v26
  2704: d78dcd2f                 vxor.vv v27, v28, v27
  2708: 5703b32f                 vxor.vv v6, v27, v6
  2712: d78cfb27                 vand.vv v25, v31, v23
  2716: 573d72a3                 vsrl.vi v26, v23, 4
  2720: d78ddc33                 vrgather.vv v27, v29, v25
  2724: 570eed33                 vrgather.vv v28, v30, v26
  2728: d78dcd2f                 vxor.vv v27, v28, v27
  2732: d783b32f                 vxor.vv v7, v27, v7
  2736: d70cfc27                 vand.vv v25, v31, v24
  2740: 573d82a3                 vsrl.vi v26, v24, 4
  2744: d78ddc33                 vrgather.vv v27, v29, v25
  2748: 570eed33                 vrgather.vv v28, v30, v26
  2752: d78dcd2f                 vxor.vv v27, v28, v27
  2756: 5704b42f                 vxor.vv v8, v27, v8
  2760: d78cf827                 vand.vv v25, v31, v17
  2764: 573d12a3                 vsrl.vi v26, v17, 4
  2768: d78ddc33                 vrgather.vv v27, v29, v25
  2772: 570eed33                 vrgather.vv v28, v30, v26
  2776: d78dcd2f                 vxor.vv v27, v28, v27
  2780: d784b42f                 vxor.vv v9, v27, v9
  2784: d70cf927                 vand.vv v25, v31, v18
  2788: 573d22a3                 vsrl.vi v26, v18, 4
  2792: d78ddc33                 vrgather.vv v27, v29, v25
  2796: 570eed33                 vrgather.vv v28, v30, v26
  2800: d78dcd2f                 vxor.vv v27, v28, v27
  2804: 5705b52f                 vxor.vv v10, v27, v10
  2808: d78cf927                 vand.vv v25, v31, v19
  2812: 573d32a3                 vsrl.vi v26, v19, 4
  2816: d78ddc33                 vrgather.vv v27, v29, v25
  2820: 570eed33                 vrgather.vv v28, v30, v26
  2824: d78dcd2f                 vxor.vv v27, v28, v27
  2828: d785b52f                 vxor.vv v11, v27, v11
  2832: d70cfa27                 vand.vv v25, v31, v20
  2836: 573d42a3                 vsrl.vi v26, v20, 4
  2840: d78ddc33                 vrgather.vv v27, v29, v25
  2844: 570eed33                 vrgather.vv v28, v30, v26
  2848: d78dcd2f                 vxor.vv v27, v28, v27
  2852: 5706b62f                 vxor.vv v12, v27, v12
  2856: d78cfa27                 vand.vv v25, v31, v21
  2860: 573d52a3                 vsrl.vi v26, v21, 4
  2864: d78ddc33                 vrgather.vv v27, v29, v25
  2868: 570eed33                 vrgather.vv v28, v30, v26
  2872: d78dcd2f                 vxor.vv v27, v28, v27
  2876: d786b62f                 vxor.vv v13, v27, v13
  2880: d70cfb27                 vand.vv v25, v31, v22
  2884: 573d62a3                 vsrl.vi v26, v22, 4
  2888: d78ddc33                 vrgather.vv v27, v29, v25
  2892: 570eed33                 vrgather.vv v28, v30, v26
  2896: d78dcd2f                 vxor.vv v27, v28, v27
  2900: 5707b72f                 vxor.vv v14, v27, v14
  2904: d78cfb27                 vand.vv v25, v31, v23
  2908: 573d72a3                 vsrl.vi v26, v23, 4
  2912: d78ddc33                 vrgather.vv v27, v29, v25
  2916: 570eed33                 vrgather.vv v28, v30, v26
  2920: d78dcd2f                 vxor.vv v27, v28, v27
  2924: d787b72f                 vxor.vv v15, v27, v15
  2928: d70cfc27                 vand.vv v25, v31, v24
  2932: 573d82a3                 vsrl.vi v26, v24, 4
  2936: d78ddc33                 vrgather.vv v27, v29, v25
  2940: 570eed33                 vrgather.vv v28, v30, v26
  2944: d78dcd2f                 vxor.vv v27, v28, v27
  2948: 5708b82f                 vxor.vv v16, v27, v16
  2952: 83b30500                 ld t2, 0(a1)
  2956: b3836300                 add t2, t2, t1
  2960: a7800302                 vse8.v v1, (t2)
  2964: b3835300                 add t2, t2, t0
  2968: 27810302                 vse8.v v2, (t2)
  2972: b3835300                 add t2, t2, t0
  2976: a7810302                 vse8.v v3, (t2)
  2980: b3835300                 add t2, t2, t0
  2984: 27820302                 vse8.v v4, (t2)
  2988: b3835300                 add t2, t2, t0
  2992: a7820302                 vse8.v v5, (t2)
  2996: b3835300                 add t2, t2, t0
  3000: 27830302                 vse8.v v6, (t2)
  3004: b3835300                 add t2, t2, t0
  3008: a7830302                 vse8.v v7, (t2)
  3012: b3835300                 add t2, t2, t0
  3016: 27840302                 vse8.v v8, (t2)
  3020: 83b38500                 ld t2, 8(a1)
  3024: b3836300                 add t2, t2, t1
  3028: a7840302                 vse8.v v9, (t2)
  3032: b3835300                 add t2, t2, t0
  3036: 27850302                 vse8.v v10, (t2)
  3040: b3835300                 add t2, t2, t0
  3044: a7850302                 vse8.v v11, (t2)
  3048: b3835300                 add t2, t2, t0
  3052: 27860302                 vse8.v v12, (t2)
  3056: b3835300                 add t2, t2, t0
  3060: a7860302                 vse8.v v13, (t2)
  3064: b3835300                 add t2, t2, t0
  3068: 27870302                 vse8.v v14, (t2)
  3072: b3835300                 add t2, t2, t0
  3076: a7870302                 vse8.v v15, (t2)
  3080: b3835300                 add t2, t2, t0
  3084: 27880302                 vse8.v v16, (t2)
  3088: 3303d300                 add t1, t1, a3
  3092: 1307f7ff                 addi a4, a4, -1
  3096: 631407c2                 bnez a4, 0x40
  3100: 33066640                 sub a2, a2, t1
  3104: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
  3108: 93060600                 mv a3, a2
  3112: 6344d000                 bgtz a3, 0xc30
  3116: 6f00c024                 j 0xe78
  3120: b3835640                 sub t2, a3, t0
  3124: 13def343                 srai t3, t2, 0x3f
  3128: b3f3c301                 and t2, t2, t3
  3132: b3837200                 add t2, t0, t2
  3136: d7a80852                 vid.v v17
  3140: 57c0136b                 vmsltu.vx v0, v17, t2
  3144: 83330500                 ld t2, 0(a0)
  3148: b3836300                 add t2, t2, t1
  3152: 87880300                 vle8.v v17, (t2), v0.t
  3156: 138e0f00                 mv t3, t6
  3160: 870e0e02                 vle8.v v29, (t3)
  3164: 138e0f01                 addi t3, t6, 0x10
  3168: 070f0e02                 vle8.v v30, (t3)
  3172: d78cf827                 vand.vv v25, v31, v17
  3176: 573d12a3                 vsrl.vi v26, v17, 4
  3180: d78ddc33                 vrgather.vv v27, v29, v25
  3184: 570eed33                 vrgather.vv v28, v30, v26
  3188: d78dcd2f                 vxor.vv v27, v28, v27
  3192: d7800d5e                 vmv.v.v v1, v27
  3196: 138e0f02                 addi t3, t6, 0x20
  3200: 870e0e02                 vle8.v v29, (t3)
  3204: 138e0f03                 addi t3, t6, 0x30
  3208: 070f0e02                 vle8.v v30, (t3)
  3212: d78cf827                 vand.vv v25, v31, v17
  3216: 573d12a3                 vsrl.vi v26, v17, 4
  3220: d78ddc33                 vrgather.vv v27, v29, v25
  3224: 570eed33                 vrgather.vv v28, v30, v26
  3228: d78dcd2f                 vxor.vv v27, v28, v27
  3232: d7840d5e                 vmv.v.v v9, v27
  3236: 83338500                 ld t2, 8(a0)
  3240: b3836300                 add t2, t2, t1
  3244: 87880300                 vle8.v v17, (t2), v0.t
  3248: 138e0f04                 addi t3, t6, 0x40
  3252: 870e0e02                 vle8.v v29, (t3)
  3256: 138e0f05                 addi t3, t6, 0x50
  3260: 070f0e02                 vle8.v v30, (t3)
  3264: d78cf827                 vand.vv v25, v31, v17
  3268: 573d12a3                 vsrl.vi v26, v17, 4
  3272: d78ddc33                 vrgather.vv v27, v29, v25
  3276: 570eed33                 vrgather.vv v28, v30, v26
  3280: d78dcd2f                 vxor.vv v27, v28, v27
  3284: d780b02f                 vxor.vv v1, v27, v1
  3288: 138e0f06                 addi t3, t6, 0x60
  3292: 870e0e02                 vle8.v v29, (t3)
  3296: 138e0f07                 addi t3, t6, 0x70
  3300: 070f0e02                 vle8.v v30, (t3)
  3304: d78cf827                 vand.vv v25, v31, v17
  3308: 573d12a3                 vsrl.vi v26, v17, 4
  3312: d78ddc33                 vrgather.vv v27, v29, v25
  3316: 570eed33                 vrgather.vv v28, v30, v26
  3320: d78dcd2f                 vxor.vv v27, v28, v27
  3324: d784b42f                 vxor.vv v9, v27, v9
  3328: 83330501                 ld t2, 0x10(a0)
  3332: b3836300                 add t2, t2, t1
  3336: 87880300                 vle8.v v17, (t2), v0.t
  3340: 138e0f08                 addi t3, t6, 0x80
  3344: 870e0e02                 vle8.v v29, (t3)
  3348: 138e0f09                 addi t3, t6, 0x90
  3352: 070f0e02                 vle8.v v30, (t3)
  3356: d78cf827                 vand.vv v25, v31, v17
  3360: 573d12a3                 vsrl.vi v26, v17, 4
  3364: d78ddc33                 vrgather.vv v27, v29, v25
  3368: 570eed33                 vrgather.vv v28, v30, v26
  3372: d78dcd2f                 vxor.vv v27, v28, v27
  3376: d780b02f                 vxor.vv v1, v27, v1
  3380: 138e0f0a                 addi t3, t6, 0xa0
  3384: 870e0e02                 vle8.v v29, (t3)
  3388: 138e0f0b                 addi t3, t6, 0xb0
  3392: 070f0e02                 vle8.v v30, (t3)
  3396: d78cf827                 vand.vv v25, v31, v17
  3400: 573d12a3                 vsrl.vi v26, v17, 4
  3404: d78ddc33                 vrgather.vv v27, v29, v25
  3408: 570eed33                 vrgather.vv v28, v30, v26
  3412: d78dcd2f                 vxor.vv v27, v28, v27
  3416: d784b42f                 vxor.vv v9, v27, v9
  3420: 83338501                 ld t2, 0x18(a0)
  3424: b3836300                 add t2, t2, t1
  3428: 87880300                 vle8.v v17, (t2), v0.t
  3432: 138e0f0c                 addi t3, t6, 0xc0
  3436: 870e0e02                 vle8.v v29, (t3)
  3440: 138e0f0d                 addi t3, t6, 0xd0
  3444: 070f0e02                 vle8.v v30, (t3)
  3448: d78cf827                 vand.vv v25, v31, v17
  3452: 573d12a3                 vsrl.vi v26, v17, 4
  3456: d78ddc33                 vrgather.vv v27, v29, v25
  3460: 570eed33                 vrgather.vv v28, v30, v26
  3464: d78dcd2f                 vxor.vv v27, v28, v27
  3468: d780b02f                 vxor.vv v1, v27, v1
  3472: 138e0f0e                 addi t3, t6, 0xe0
  3476: 870e0e02                 vle8.v v29, (t3)
  3480: 138e0f0f                 addi t3, t6, 0xf0
  3484: 070f0e02                 vle8.v v30, (t3)
  3488: d78cf827                 vand.vv v25, v31, v17
  3492: 573d12a3                 vsrl.vi v26, v17, 4
  3496: d78ddc33                 vrgather.vv v27, v29, v25
  3500: 570eed33                 vrgather.vv v28, v30, v26
  3504: d78dcd2f                 vxor.vv v27, v28, v27
  3508: d784b42f                 vxor.vv v9, v27, v9
  3512: 83330502                 ld t2, 0x20(a0)
  3516: b3836300                 add t2, t2, t1
  3520: 87880300                 vle8.v v17, (t2), v0.t
  3524: 138e0f10                 addi t3, t6, 0x100
  3528: 870e0e02                 vle8.v v29, (t3)
  3532: 138e0f11                 addi t3, t6, 0x110
  3536: 070f0e02                 vle8.v v30, (t3)
  3540: d78cf827                 vand.vv v25, v31, v17
  3544: 573d12a3                 vsrl.vi v26, v17, 4
  3548: d78ddc33                 vrgather.vv v27, v29, v25
  3552: 570eed33                 vrgather.vv v28, v30, v26
  3556: d78dcd2f                 vxor.vv v27, v28, v27
  3560: d780b02f                 vxor.vv v1, v27, v1
  3564: 138e0f12                 addi t3, t6, 0x120
  3568: 870e0e02                 vle8.v v29, (t3)
  3572: 138e0f13                 addi t3, t6, 0x130
  3576: 070f0e02                 vle8.v v30, (t3)
  3580: d78cf827                 vand.vv v25, v31, v17
  3584: 573d12a3                 vsrl.vi v26, v17, 4
  3588: d78ddc33                 vrgather.vv v27, v29, v25
  3592: 570eed33                 vrgather.vv v28, v30, v26
  3596: d78dcd2f                 vxor.vv v27, v28, v27
  3600: d784b42f                 vxor.vv v9, v27, v9
  3604: 83338502                 ld t2, 0x28(a0)
  3608: b3836300                 add t2, t2, t1
  3612: 87880300                 vle8.v v17, (t2), v0.t
  3616: 138e0f14                 addi t3, t6, 0x140
  3620: 870e0e02                 vle8.v v29, (t3)
  3624: 138e0f15                 addi t3, t6, 0x150
  3628: 070f0e02                 vle8.v v30, (t3)
  3632: d78cf827                 vand.vv v25, v31, v17
  3636: 573d12a3                 vsrl.vi v26, v17, 4
  3640: d78ddc33                 vrgather.vv v27, v29, v25
  3644: 570eed33                 vrgather.vv v28, v30, v26
  3648: d78dcd2f                 vxor.vv v27, v28, v27
  3652: d780b02f                 vxor.vv v1, v27, v1
  3656: d78cf827                 vand.vv v25, v31, v17
  3660: 573d12a3                 vsrl.vi v26, v17, 4
  3664: d78ddc33                 vrgather.vv v27, v29, v25
  3668: 570eed33                 vrgather.vv v28, v30, v26
  3672: d78dcd2f                 vxor.vv v27, v28, v27
  3676: d784b42f                 vxor.vv v9, v27, v9
  3680: 83b30500                 ld t2, 0(a1)
  3684: b3836300                 add t2, t2, t1
  3688: a7800300                 vse8.v v1, (t2), v0.t
  3692: 83b38500                 ld t2, 8(a1)
  3696: b3836300                 add t2, t2, t1
  3700: a7840300                 vse8.v v9, (t2), v0.t
  3704: 33035300                 add t1, t1, t0
  3708: b3865640                 sub a3, a3, t0
  3712: 6344d000                 bgtz a3, 0xe88
  3716: 6f00c024                 j 0x10d0
  3720: b3835640                 sub t2, a3, t0
  3724: 13def343                 srai t3, t2, 0x3f
  3728: b3f3c301                 and t2, t2, t3
  3732: b3837200                 add t2, t0, t2
  3736: d7a80852                 vid.v v17
  3740: 57c0136b                 vmsltu.vx v0, v17, t2
  3744: 83330500                 ld t2, 0(a0)
  3748: b3836300                 add t2, t2, t1
  3752: 87880300                 vle8.v v17, (t2), v0.t
  3756: 138e0f00                 mv t3, t6
  3760: 870e0e02                 vle8.v v29, (t3)
  3764: 138e0f01                 addi t3, t6, 0x10
  3768: 070f0e02                 vle8.v v30, (t3)
  3772: d78cf827                 vand.vv v25, v31, v17
  3776: 573d12a3                 vsrl.vi v26, v17, 4
  3780: d78ddc33                 vrgather.vv v27, v29, v25
  3784: 570eed33                 vrgather.vv v28, v30, v26
  3788: d78dcd2f                 vxor.vv v27, v28, v27
  3792: d7800d5e                 vmv.v.v v1, v27
  3796: 138e0f02                 addi t3, t6, 0x20
  3800: 870e0e02                 vle8.v v29, (t3)
  3804: 138e0f03                 addi t3, t6, 0x30
  3808: 070f0e02                 vle8.v v30, (t3)
  3812: d78cf827                 vand.vv v25, v31, v17
  3816: 573d12a3                 vsrl.vi v26, v17, 4
  3820: d78ddc33                 vrgather.vv v27, v29, v25
  3824: 570eed33                 vrgather.vv v28, v30, v26
  3828: d78dcd2f                 vxor.vv v27, v28, v27
  3832: d7840d5e                 vmv.v.v v9, v27
  3836: 83338500                 ld t2, 8(a0)
  3840: b3836300                 add t2, t2, t1
  3844: 87880300                 vle8.v v17, (t2), v0.t
  3848: 138e0f04                 addi t3, t6, 0x40
  3852: 870e0e02                 vle8.v v29, (t3)
  3856: 138e0f05                 addi t3, t6, 0x50
  3860: 070f0e02                 vle8.v v30, (t3)
  3864: d78cf827                 vand.vv v25, v31, v17
  3868: 573d12a3                 vsrl.vi v26, v17, 4
  3872: d78ddc33                 vrgather.vv v27, v29, v25
  3876: 570eed33                 vrgather.vv v28, v30, v26
  3880: d78dcd2f                 vxor.vv v27, v28, v27
  3884: d780b02f                 vxor.vv v1, v27, v1
  3888: 138e0f06                 addi t3, t6, 0x60
  3892: 870e0e02                 vle8.v v29, (t3)
  3896: 138e0f07                 addi t3, t6, 0x70
  3900: 070f0e02                 vle8.v v30, (t3)
  3904: d78cf827                 vand.vv v25, v31, v17
  3908: 573d12a3                 vsrl.vi v26, v17, 4
  3912: d78ddc33                 vrgather.vv v27, v29, v25
  3916: 570eed33                 vrgather.vv v28, v30, v26
  3920: d78dcd2f                 vxor.vv v27, v28, v27
  3924: d784b42f                 vxor.vv v9, v27, v9
  3928: 83330501                 ld t2, 0x10(a0)
  3932: b3836300                 add t2, t2, t1
  3936: 87880300                 vle8.v v17, (t2), v0.t
  3940: 138e0f08                 addi t3, t6, 0x80
  3944: 870e0e02                 vle8.v v29, (t3)
  3948: 138e0f09                 addi t3, t6, 0x90
  3952: 070f0e02                 vle8.v v30, (t3)
  3956: d78cf827                 vand.vv v25, v31, v17
  3960: 573d12a3                 vsrl.vi v26, v17, 4
  3964: d78ddc33                 vrgather.vv v27, v29, v25
  3968: 570eed33                 vrgather.vv v28, v30, v26
  3972: d78dcd2f                 vxor.vv v27, v28, v27
  3976: d780b02f                 vxor.vv v1, v27, v1
  3980: 138e0f0a                 addi t3, t6, 0xa0
  3984: 870e0e02                 vle8.v v29, (t3)
  3988: 138e0f0b                 addi t3, t6, 0xb0
  3992: 070f0e02                 vle8.v v30, (t3)
  3996: d78cf827                 vand.vv v25, v31, v17
  4000: 573d12a3                 vsrl.vi v26, v17, 4
  4004: d78ddc33                 vrgather.vv v27, v29, v25
  4008: 570eed33                 vrgather.vv v28, v30, v26
  4012: d78dcd2f                 vxor.vv v27, v28, v27
  4016: d784b42f                 vxor.vv v9, v27, v9
  4020: 83338501                 ld t2, 0x18(a0)
  4024: b3836300                 add t2, t2, t1
  4028: 87880300                 vle8.v v17, (t2), v0.t
  4032: 138e0f0c                 addi t3, t6, 0xc0
  4036: 870e0e02                 vle8.v v29, (t3)
  4040: 138e0f0d                 addi t3, t6, 0xd0
  4044: 070f0e02                 vle8.v v30, (t3)
  4048: d78cf827                 vand.vv v25, v31, v17
  4052: 573d12a3                 vsrl.vi v26, v17, 4
  4056: d78ddc33                 vrgather.vv v27, v29, v25
  4060: 570eed33                 vrgather.vv v28, v30, v26
  4064: d78dcd2f                 vxor.vv v27, v28, v27
  4068: d780b02f                 vxor.vv v1, v27, v1
  4072: 138e0f0e                 addi t3, t6, 0xe0
  4076: 870e0e02                 vle8.v v29, (t3)
  4080: 138e0f0f                 addi t3, t6, 0xf0
  4084: 070f0e02                 vle8.v v30, (t3)
  4088: d78cf827                 vand.vv v25, v31, v17
  4092: 573d12a3                 vsrl.vi v26, v17, 4
  4096: d78ddc33                 vrgather.vv v27, v29, v25
  4100: 570eed33                 vrgather.vv v28, v30, v26
  4104: d78dcd2f                 vxor.vv v27, v28, v27
  4108: d784b42f                 vxor.vv v9, v27, v9
  4112: 83330502                 ld t2, 0x20(a0)
  4116: b3836300                 add t2, t2, t1
  4120: 87880300                 vle8.v v17, (t2), v0.t
  4124: 138e0f10                 addi t3, t6, 0x100
  4128: 870e0e02                 vle8.v v29, (t3)
  4132: 138e0f11                 addi t3, t6, 0x110
  4136: 070f0e02                 vle8.v v30, (t3)
  4140: d78cf827                 vand.vv v25, v31, v17
  4144: 573d12a3                 vsrl.vi v26, v17, 4
  4148: d78ddc33                 vrgather.vv v27, v29, v25
  4152: 570eed33                 vrgather.vv v28, v30, v26
  4156: d78dcd2f                 vxor.vv v27, v28, v27
  4160: d780b02f                 vxor.vv v1, v27, v1
  4164: 138e0f12                 addi t3, t6, 0x120
  4168: 870e0e02                 vle8.v v29, (t3)
  4172: 138e0f13                 addi t3, t6, 0x130
  4176: 070f0e02                 vle8.v v30, (t3)
  4180: d78cf827                 vand.vv v25, v31, v17
  4184: 573d12a3                 vsrl.vi v26, v17, 4
  4188: d78ddc33                 vrgather.vv v27, v29, v25
  4192: 570eed33                 vrgather.vv v28, v30, v26
  4196: d78dcd2f                 vxor.vv v27, v28, v27
  4200: d784b42f                 vxor.vv v9, v27, v9
  4204: 83338502                 ld t2, 0x28(a0)
  4208: b3836300                 add t2, t2, t1
  4212: 87880300                 vle8.v v17, (t2), v0.t
  4216: 138e0f14                 addi t3, t6, 0x140
  4220: 870e0e02                 vle8.v v29, (t3)
  4224: 138e0f15                 addi t3, t6, 0x150
  4228: 070f0e02                 vle8.v v30, (t3)
  4232: d78cf827                 vand.vv v25, v31, v17
  4236: 573d12a3                 vsrl.vi v26, v17, 4
  4240: d78ddc33                 vrgather.vv v27, v29, v25
  4244: 570eed33                 vrgather.vv v28, v30, v26
  4248: d78dcd2f                 vxor.vv v27, v28, v27
  4252: d780b02f                 vxor.vv v1, v27, v1
  4256: d78cf827                 vand.vv v25, v31, v17
  4260: 573d12a3                 vsrl.vi v26, v17, 4
  4264: d78ddc33                 vrgather.vv v27, v29, v25
  4268: 570eed33                 vrgather.vv v28, v30, v26
  4272: d78dcd2f                 vxor.vv v27, v28, v27
  4276: d784b42f                 vxor.vv v9, v27, v9
  4280: 83b30500                 ld t2, 0(a1)
  4284: b3836300                 add t2, t2, t1
  4288: a7800300                 vse8.v v1, (t2), v0.t
  4292: 83b38500                 ld t2, 8(a1)
  4296: b3836300                 add t2, t2, t1
  4300: a7840300                 vse8.v v9, (t2), v0.t
  4304: 33035300                 add t1, t1, t0
  4308: b3865640                 sub a3, a3, t0
  4312: 6344d000                 bgtz a3, 0x10e0
  4316: 6f00c024                 j 0x1328
  4320: b3835640                 sub t2, a3, t0
  4324: 13def343                 srai t3, t2, 0x3f
  4328: b3f3c301                 and t2, t2, t3
  4332: b3837200                 add t2, t0, t2
  4336: d7a80852                 vid.v v17
  4340: 57c0136b                 vmsltu.vx v0, v17, t2
  4344: 83330500                 ld t2, 0(a0)
  4348: b3836300                 add t2, t2, t1
  4352: 87880300                 vle8.v v17, (t2), v0.t
  4356: 138e0f00                 mv t3, t6
  4360: 870e0e02                 vle8.v v29, (t3)
  4364: 138e0f01                 addi t3, t6, 0x10
  4368: 070f0e02                 vle8.v v30, (t3)
  4372: d78cf827                 vand.vv v25, v31, v17
  4376: 573d12a3                 vsrl.vi v26, v17, 4
  4380: d78ddc33                 vrgather.vv v27, v29, v25
  4384: 570eed33                 vrgather.vv v28, v30, v26
  4388: d78dcd2f                 vxor.vv v27, v28, v27
  4392: d7800d5e                 vmv.v.v v1, v27
  4396: 138e0f02                 addi t3, t6, 0x20
  4400: 870e0e02                 vle8.v v29, (t3)
  4404: 138e0f03                 addi t3, t6, 0x30
  4408: 070f0e02                 vle8.v v30, (t3)
  4412: d78cf827                 vand.vv v25, v31, v17
  4416: 573d12a3                 vsrl.vi v26, v17, 4
  4420: d78ddc33                 vrgather.vv v27, v29, v25
  4424: 570eed33                 vrgather.vv v28, v30, v26
  4428: d78dcd2f                 vxor.vv v27, v28, v27
  4432: d7840d5e                 vmv.v.v v9, v27
  4436: 83338500                 ld t2, 8(a0)
  4440: b3836300                 add t2, t2, t1
  4444: 87880300                 vle8.v v17, (t2), v0.t
  4448: 138e0f04                 addi t3, t6, 0x40
  4452: 870e0e02                 vle8.v v29, (t3)
  4456: 138e0f05                 addi t3, t6, 0x50
  4460: 070f0e02                 vle8.v v30, (t3)
  4464: d78cf827                 vand.vv v25, v31, v17
  4468: 573d12a3                 vsrl.vi v26, v17, 4
  4472: d78ddc33                 vrgather.vv v27, v29, v25
  4476: 570eed33                 vrgather.vv v28, v30, v26
  4480: d78dcd2f                 vxor.vv v27, v28, v27
  4484: d780b02f                 vxor.vv v1, v27, v1
  4488: 138e0f06                 addi t3, t6, 0x60
  4492: 870e0e02                 vle8.v v29, (t3)
  4496: 138e0f07                 addi t3, t6, 0x70
  4500: 070f0e02                 vle8.v v30, (t3)
  4504: d78cf827                 vand.vv v25, v31, v17
  4508: 573d12a3                 vsrl.vi v26, v17, 4
  4512: d78ddc33                 vrgather.vv v27, v29, v25
  4516: 570eed33                 vrgather.vv v28, v30, v26
  4520: d78dcd2f                 vxor.vv v27, v28, v27
  4524: d784b42f                 vxor.vv v9, v27, v9
  4528: 83330501                 ld t2, 0x10(a0)
  4532: b3836300                 add t2, t2, t1
  4536: 87880300                 vle8.v v17, (t2), v0.t
  4540: 138e0f08                 addi t3, t6, 0x80
  4544: 870e0e02                 vle8.v v29, (t3)
  4548: 138e0f09                 addi t3, t6, 0x90
  4552: 070f0e02                 vle8.v v30, (t3)
  4556: d78cf827                 vand.vv v25, v31, v17
  4560: 573d12a3                 vsrl.vi v26, v17, 4
  4564: d78ddc33                 vrgather.vv v27, v29, v25
  4568: 570eed33                 vrgather.vv v28, v30, v26
  4572: d78dcd2f                 vxor.vv v27, v28, v27
  4576: d780b02f                 vxor.vv v1, v27, v1
  4580: 138e0f0a                 addi t3, t6, 0xa0
  4584: 870e0e02                 vle8.v v29, (t3)
  4588: 138e0f0b                 addi t3, t6, 0xb0
  4592: 070f0e02                 vle8.v v30, (t3)
  4596: d78cf827                 vand.vv v25, v31, v17
  4600: 573d12a3                 vsrl.vi v26, v17, 4
  4604: d78ddc33                 vrgather.vv v27, v29, v25
  4608: 570eed33                 vrgather.vv v28, v30, v26
  4612: d78dcd2f                 vxor.vv v27, v28, v27
  4616: d784b42f                 vxor.vv v9, v27, v9
  4620: 83338501                 ld t2, 0x18(a0)
  4624: b3836300                 add t2, t2, t1
  4628: 87880300                 vle8.v v17, (t2), v0.t
  4632: 138e0f0c                 addi t3, t6, 0xc0
  4636: 870e0e02                 vle8.v v29, (t3)
  4640: 138e0f0d                 addi t3, t6, 0xd0
  4644: 070f0e02                 vle8.v v30, (t3)
  4648: d78cf827                 vand.vv v25, v31, v17
  4652: 573d12a3                 vsrl.vi v26, v17, 4
  4656: d78ddc33                 vrgather.vv v27, v29, v25
  4660: 570eed33                 vrgather.vv v28, v30, v26
  4664: d78dcd2f                 vxor.vv v27, v28, v27
  4668: d780b02f                 vxor.vv v1, v27, v1
  4672: 138e0f0e                 addi t3, t6, 0xe0
  4676: 870e0e02                 vle8.v v29, (t3)
  4680: 138e0f0f                 addi t3, t6, 0xf0
  4684: 070f0e02                 vle8.v v30, (t3)
  4688: d78cf827                 vand.vv v25, v31, v17
  4692: 573d12a3                 vsrl.vi v26, v17, 4
  4696: d78ddc33                 vrgather.vv v27, v29, v25
  4700: 570eed33                 vrgather.vv v28, v30, v26
  4704: d78dcd2f                 vxor.vv v27, v28, v27
  4708: d784b42f                 vxor.vv v9, v27, v9
  4712: 83330502                 ld t2, 0x20(a0)
  4716: b3836300                 add t2, t2, t1
  4720: 87880300                 vle8.v v17, (t2), v0.t
  4724: 138e0f10                 addi t3, t6, 0x100
  4728: 870e0e02                 vle8.v v29, (t3)
  4732: 138e0f11                 addi t3, t6, 0x110
  4736: 070f0e02                 vle8.v v30, (t3)
  4740: d78cf827                 vand.vv v25, v31, v17
  4744: 573d12a3                 vsrl.vi v26, v17, 4
  4748: d78ddc33                 vrgather.vv v27, v29, v25
  4752: 570eed33                 vrgather.vv v28, v30, v26
  4756: d78dcd2f                 vxor.vv v27, v28, v27
  4760: d780b02f                 vxor.vv v1, v27, v1
  4764: 138e0f12                 addi t3, t6, 0x120
  4768: 870e0e02                 vle8.v v29, (t3)
  4772: 138e0f13                 addi t3, t6, 0x130
  4776: 070f0e02                 vle8.v v30, (t3)
  4780: d78cf827                 vand.vv v25, v31, v17
  4784: 573d12a3                 vsrl.vi v26, v17, 4
  4788: d78ddc33                 vrgather.vv v27, v29, v25
  4792: 570eed33                 vrgather.vv v28, v30, v26
  4796: d78dcd2f                 vxor.vv v27, v28, v27
  4800: d784b42f                 vxor.vv v9, v27, v9
  4804: 83338502                 ld t2, 0x28(a0)
  4808: b3836300                 add t2, t2, t1
  4812: 87880300                 vle8.v v17, (t2), v0.t
  4816: 138e0f14                 addi t3, t6, 0x140
  4820: 870e0e02                 vle8.v v29, (t3)
  4824: 138e0f15                 addi t3, t6, 0x150
  4828: 070f0e02                 vle8.v v30, (t3)
  4832: d78cf827                 vand.vv v25, v31, v17
  4836: 573d12a3                 vsrl.vi v26, v17, 4
  4840: d78ddc33                 vrgather.vv v27, v29, v25
  4844: 570eed33                 vrgather.vv v28, v30, v26
  4848: d78dcd2f                 vxor.vv v27, v28, v27
  4852: d780b02f                 vxor.vv v1, v27, v1
  4856: d78cf827                 vand.vv v25, v31, v17
  4860: 573d12a3                 vsrl.vi v26, v17, 4
  4864: d78ddc33                 vrgather.vv v27, v29, v25
  4868: 570eed33                 vrgather.vv v28, v30, v26
  4872: d78dcd2f                 vxor.vv v27, v28, v27
  4876: d784b42f                 vxor.vv v9, v27, v9
  4880: 83b30500                 ld t2, 0(a1)
  4884: b3836300                 add t2, t2, t1
  4888: a7800300                 vse8.v v1, (t2), v0.t
  4892: 83b38500                 ld t2, 8(a1)
  4896: b3836300                 add t2, t2, t1
  4900: a7840300                 vse8.v v9, (t2), v0.t
  4904: 33035300                 add t1, t1, t0
  4908: b3865640                 sub a3, a3, t0
  4912: 6344d000                 bgtz a3, 0x1338
  4916: 6f00c024                 j 0x1580
  4920: b3835640                 sub t2, a3, t0
  4924: 13def343                 srai t3, t2, 0x3f
  4928: b3f3c301                 and t2, t2, t3
  4932: b3837200                 add t2, t0, t2
  4936: d7a80852                 vid.v v17
  4940: 57c0136b                 vmsltu.vx v0, v17, t2
  4944: 83330500                 ld t2, 0(a0)
  4948: b3836300                 add t2, t2, t1
  4952: 87880300                 vle8.v v17, (t2), v0.t
  4956: 138e0f00                 mv t3, t6
  4960: 870e0e02                 vle8.v v29, (t3)
  4964: 138e0f01                 addi t3, t6, 0x10
  4968: 070f0e02                 vle8.v v30, (t3)
  4972: d78cf827                 vand.vv v25, v31, v17
  4976: 573d12a3                 vsrl.vi v26, v17, 4
  4980: d78ddc33                 vrgather.vv v27, v29, v25
  4984: 570eed33                 vrgather.vv v28, v30, v26
  4988: d78dcd2f                 vxor.vv v27, v28, v27
  4992: d7800d5e                 vmv.v.v v1, v27
  4996: 138e0f02                 addi t3, t6, 0x20
  5000: 870e0e02                 vle8.v v29, (t3)
  5004: 138e0f03                 addi t3, t6, 0x30
  5008: 070f0e02                 vle8.v v30, (t3)
  5012: d78cf827                 vand.vv v25, v31, v17
  5016: 573d12a3                 vsrl.vi v26, v17, 4
  5020: d78ddc33                 vrgather.vv v27, v29, v25
  5024: 570eed33                 vrgather.vv v28, v30, v26
  5028: d78dcd2f                 vxor.vv v27, v28, v27
  5032: d7840d5e                 vmv.v.v v9, v27
  5036: 83338500                 ld t2, 8(a0)
  5040: b3836300                 add t2, t2, t1
  5044: 87880300                 vle8.v v17, (t2), v0.t
  5048: 138e0f04                 addi t3, t6, 0x40
  5052: 870e0e02                 vle8.v v29, (t3)
  5056: 138e0f05                 addi t3, t6, 0x50
  5060: 070f0e02                 vle8.v v30, (t3)
  5064: d78cf827                 vand.vv v25, v31, v17
  5068: 573d12a3                 vsrl.vi v26, v17, 4
  5072: d78ddc33                 vrgather.vv v27, v29, v25
  5076: 570eed33                 vrgather.vv v28, v30, v26
  5080: d78dcd2f                 vxor.vv v27, v28, v27
  5084: d780b02f                 vxor.vv v1, v27, v1
  5088: 138e0f06                 addi t3, t6, 0x60
  5092: 870e0e02                 vle8.v v29, (t3)
  5096: 138e0f07                 addi t3, t6, 0x70
  5100: 070f0e02                 vle8.v v30, (t3)
  5104: d78cf827                 vand.vv v25, v31, v17
  5108: 573d12a3                 vsrl.vi v26, v17, 4
  5112: d78ddc33                 vrgather.vv v27, v29, v25
  5116: 570eed33                 vrgather.vv v28, v30, v26
  5120: d78dcd2f                 vxor.vv v27, v28, v27
  5124: d784b42f                 vxor.vv v9, v27, v9
  5128: 83330501                 ld t2, 0x10(a0)
  5132: b3836300                 add t2, t2, t1
  5136: 87880300                 vle8.v v17, (t2), v0.t
  5140: 138e0f08                 addi t3, t6, 0x80
  5144: 870e0e02                 vle8.v v29, (t3)
  5148: 138e0f09                 addi t3, t6, 0x90
  5152: 070f0e02                 vle8.v v30, (t3)
  5156: d78cf827                 vand.vv v25, v31, v17
  5160: 573d12a3                 vsrl.vi v26, v17, 4
  5164: d78ddc33                 vrgather.vv v27, v29, v25
  5168: 570eed33                 vrgather.vv v28, v30, v26
  5172: d78dcd2f                 vxor.vv v27, v28, v27
  5176: d780b02f                 vxor.vv v1, v27, v1
  5180: 138e0f0a                 addi t3, t6, 0xa0
  5184: 870e0e02                 vle8.v v29, (t3)
  5188: 138e0f0b                 addi t3, t6, 0xb0
  5192: 070f0e02                 vle8.v v30, (t3)
  5196: d78cf827                 vand.vv v25, v31, v17
  5200: 573d12a3                 vsrl.vi v26, v17, 4
  5204: d78ddc33                 vrgather.vv v27, v29, v25
  5208: 570eed33                 vrgather.vv v28, v30, v26
  5212: d78dcd2f                 vxor.vv v27, v28, v27
  5216: d784b42f                 vxor.vv v9, v27, v9
  5220: 83338501                 ld t2, 0x18(a0)
  5224: b3836300                 add t2, t2, t1
  5228: 87880300                 vle8.v v17, (t2), v0.t
  5232: 138e0f0c                 addi t3, t6, 0xc0
  5236: 870e0e02                 vle8.v v29, (t3)
  5240: 138e0f0d                 addi t3, t6, 0xd0
  5244: 070f0e02                 vle8.v v30, (t3)
  5248: d78cf827                 vand.vv v25, v31, v17
  5252: 573d12a3                 vsrl.vi v26, v17, 4
  5256: d78ddc33                 vrgather.vv v27, v29, v25
  5260: 570eed33                 vrgather.vv v28, v30, v26
  5264: d78dcd2f                 vxor.vv v27, v28, v27
  5268: d780b02f                 vxor.vv v1, v27, v1
  5272: 138e0f0e                 addi t3, t6, 0xe0
  5276: 870e0e02                 vle8.v v29, (t3)
  5280: 138e0f0f                 addi t3, t6, 0xf0
  5284: 070f0e02                 vle8.v v30, (t3)
  5288: d78cf827                 vand.vv v25, v31, v17
  5292: 573d12a3                 vsrl.vi v26, v17, 4
  5296: d78ddc33                 vrgather.vv v27, v29, v25
  5300: 570eed33                 vrgather.vv v28, v30, v26
  5304: d78dcd2f                 vxor.vv v27, v28, v27
  5308: d784b42f                 vxor.vv v9, v27, v9
  5312: 83330502                 ld t2, 0x20(a0)
  5316: b3836300                 add t2, t2, t1
  5320: 87880300                 vle8.v v17, (t2), v0.t
  5324: 138e0f10                 addi t3, t6, 0x100
  5328: 870e0e02                 vle8.v v29, (t3)
  5332: 138e0f11                 addi t3, t6, 0x110
  5336: 070f0e02                 vle8.v v30, (t3)
  5340: d78cf827                 vand.vv v25, v31, v17
  5344: 573d12a3                 vsrl.vi v26, v17, 4
  5348: d78ddc33                 vrgather.vv v27, v29, v25
  5352: 570eed33                 vrgather.vv v28, v30, v26
  5356: d78dcd2f                 vxor.vv v27, v28, v27
  5360: d780b02f                 vxor.vv v1, v27, v1
  5364: 138e0f12                 addi t3, t6, 0x120
  5368: 870e0e02                 vle8.v v29, (t3)
  5372: 138e0f13                 addi t3, t6, 0x130
  5376: 070f0e02                 vle8.v v30, (t3)
  5380: d78cf827                 vand.vv v25, v31, v17
  5384: 573d12a3                 vsrl.vi v26, v17, 4
  5388: d78ddc33                 vrgather.vv v27, v29, v25
  5392: 570eed33                 vrgather.vv v28, v30, v26
  5396: d78dcd2f                 vxor.vv v27, v28, v27
  5400: d784b42f                 vxor.vv v9, v27, v9
  5404: 83338502                 ld t2, 0x28(a0)
  5408: b3836300                 add t2, t2, t1
  5412: 87880300                 vle8.v v17, (t2), v0.t
  5416: 138e0f14                 addi t3, t6, 0x140
  5420: 870e0e02                 vle8.v v29, (t3)
  5424: 138e0f15                 addi t3, t6, 0x150
  5428: 070f0e02                 vle8.v v30, (t3)
  5432: d78cf827                 vand.vv v25, v31, v17
  5436: 573d12a3                 vsrl.vi v26, v17, 4
  5440: d78ddc33                 vrgather.vv v27, v29, v25
  5444: 570eed33                 vrgather.vv v28, v30, v26
  5448: d78dcd2f                 vxor.vv v27, v28, v27
  5452: d780b02f                 vxor.vv v1, v27, v1
  5456: d78cf827                 vand.vv v25, v31, v17
  5460: 573d12a3                 vsrl.vi v26, v17, 4
  5464: d78ddc33                 vrgather.vv v27, v29, v25
  5468: 570eed33                 vrgather.vv v28, v30, v26
  5472: d78dcd2f                 vxor.vv v27, v28, v27
  5476: d784b42f                 vxor.vv v9, v27, v9
  5480: 83b30500                 ld t2, 0(a1)
  5484: b3836300                 add t2, t2, t1
  5488: a7800300                 vse8.v v1, (t2), v0.t
  5492: 83b38500                 ld t2, 8(a1)
  5496: b3836300                 add t2, t2, t1
  5500: a7840300                 vse8.v v9, (t2), v0.t
  5504: 33035300                 add t1, t1, t0
  5508: b3865640                 sub a3, a3, t0
  5512: 6344d000                 bgtz a3, 0x1590
  5516: 6f00c024                 j 0x17d8
  5520: b3835640                 sub t2, a3, t0
  5524: 13def343                 srai t3, t2, 0x3f
  5528: b3f3c301                 and t2, t2, t3
  5532: b3837200                 add t2, t0, t2
  5536: d7a80852                 vid.v v17
  5540: 57c0136b                 vmsltu.vx v0, v17, t2
  5544: 83330500                 ld t2, 0(a0)
  5548: b3836300                 add t2, t2, t1
  5552: 87880300                 vle8.v v17, (t2), v0.t
  5556: 138e0f00                 mv t3, t6
  5560: 870e0e02                 vle8.v v29, (t3)
  5564: 138e0f01                 addi t3, t6, 0x10
  5568: 070f0e02                 vle8.v v30, (t3)
  5572: d78cf827                 vand.vv v25, v31, v17
  5576: 573d12a3                 vsrl.vi v26, v17, 4
  5580: d78ddc33                 vrgather.vv v27, v29, v25
  5584: 570eed33                 vrgather.vv v28, v30, v26
  5588: d78dcd2f                 vxor.vv v27, v28, v27
  5592: d7800d5e                 vmv.v.v v1, v27
  5596: 138e0f02                 addi t3, t6, 0x20
  5600: 870e0e02                 vle8.v v29, (t3)
  5604: 138e0f03                 addi t3, t6, 0x30
  5608: 070f0e02                 vle8.v v30, (t3)
  5612: d78cf827                 vand.vv v25, v31, v17
  5616: 573d12a3                 vsrl.vi v26, v17, 4
  5620: d78ddc33                 vrgather.vv v27, v29, v25
  5624: 570eed33                 vrgather.vv v28, v30, v26
  5628: d78dcd2f                 vxor.vv v27, v28, v27
  5632: d7840d5e                 vmv.v.v v9, v27
  5636: 83338500                 ld t2, 8(a0)
  5640: b3836300                 add t2, t2, t1
  5644: 87880300                 vle8.v v17, (t2), v0.t
  5648: 138e0f04                 addi t3, t6, 0x40
  5652: 870e0e02                 vle8.v v29, (t3)
  5656: 138e0f05                 addi t3, t6, 0x50
  5660: 070f0e02                 vle8.v v30, (t3)
  5664: d78cf827                 vand.vv v25, v31, v17
  5668: 573d12a3                 vsrl.vi v26, v17, 4
  5672: d78ddc33                 vrgather.vv v27, v29, v25
  5676: 570eed33                 vrgather.vv v28, v30, v26
  5680: d78dcd2f                 vxor.vv v27, v28, v27
  5684: d780b02f                 vxor.vv v1, v27, v1
  5688: 138e0f06                 addi t3, t6, 0x60
  5692: 870e0e02                 vle8.v v29, (t3)
  5696: 138e0f07                 addi t3, t6, 0x70
  5700: 070f0e02                 vle8.v v30, (t3)
  5704: d78cf827                 vand.vv v25, v31, v17
  5708: 573d12a3                 vsrl.vi v26, v17, 4
  5712: d78ddc33                 vrgather.vv v27, v29, v25
  5716: 570eed33                 vrgather.vv v28, v30, v26
  5720: d78dcd2f                 vxor.vv v27, v28, v27
  5724: d784b42f                 vxor.vv v9, v27, v9
  5728: 83330501                 ld t2, 0x10(a0)
  5732: b3836300                 add t2, t2, t1
  5736: 87880300                 vle8.v v17, (t2), v0.t
  5740: 138e0f08                 addi t3, t6, 0x80
  5744: 870e0e02                 vle8.v v29, (t3)
  5748: 138e0f09                 addi t3, t6, 0x90
  5752: 070f0e02                 vle8.v v30, (t3)
  5756: d78cf827                 vand.vv v25, v31, v17
  5760: 573d12a3                 vsrl.vi v26, v17, 4
  5764: d78ddc33                 vrgather.vv v27, v29, v25
  5768: 570eed33                 vrgather.vv v28, v30, v26
  5772: d78dcd2f                 vxor.vv v27, v28, v27
  5776: d780b02f                 vxor.vv v1, v27, v1
  5780: 138e0f0a                 addi t3, t6, 0xa0
  5784: 870e0e02                 vle8.v v29, (t3)
  5788: 138e0f0b                 addi t3, t6, 0xb0
  5792: 070f0e02                 vle8.v v30, (t3)
  5796: d78cf827                 vand.vv v25, v31, v17
  5800: 573d12a3                 vsrl.vi v26, v17, 4
  5804: d78ddc33                 vrgather.vv v27, v29, v25
  5808: 570eed33                 vrgather.vv v28, v30, v26
  5812: d78dcd2f                 vxor.vv v27, v28, v27
  5816: d784b42f                 vxor.vv v9, v27, v9
  5820: 83338501                 ld t2, 0x18(a0)
  5824: b3836300                 add t2, t2, t1
  5828: 87880300                 vle8.v v17, (t2), v0.t
  5832: 138e0f0c                 addi t3, t6, 0xc0
  5836: 870e0e02                 vle8.v v29, (t3)
  5840: 138e0f0d                 addi t3, t6, 0xd0
  5844: 070f0e02                 vle8.v v30, (t3)
  5848: d78cf827                 vand.vv v25, v31, v17
  5852: 573d12a3                 vsrl.vi v26, v17, 4
  5856: d78ddc33                 vrgather.vv v27, v29, v25
  5860: 570eed33                 vrgather.vv v28, v30, v26
  5864: d78dcd2f                 vxor.vv v27, v28, v27
  5868: d780b02f                 vxor.vv v1, v27, v1
  5872: 138e0f0e                 addi t3, t6, 0xe0
  5876: 870e0e02                 vle8.v v29, (t3)
  5880: 138e0f0f                 addi t3, t6, 0xf0
  5884: 070f0e02                 vle8.v v30, (t3)
  5888: d78cf827                 vand.vv v25, v31, v17
  5892: 573d12a3                 vsrl.vi v26, v17, 4
  5896: d78ddc33                 vrgather.vv v27, v29, v25
  5900: 570eed33                 vrgather.vv v28, v30, v26
  5904: d78dcd2f                 vxor.vv v27, v28, v27
  5908: d784b42f                 vxor.vv v9, v27, v9
  5912: 83330502                 ld t2, 0x20(a0)
  5916: b3836300                 add t2, t2, t1
  5920: 87880300                 vle8.v v17, (t2), v0.t
  5924: 138e0f10                 addi t3, t6, 0x100
  5928: 870e0e02                 vle8.v v29, (t3)
  5932: 138e0f11                 addi t3, t6, 0x110
  5936: 070f0e02                 vle8.v v30, (t3)
  5940: d78cf827                 vand.vv v25, v31, v17
  5944: 573d12a3                 vsrl.vi v26, v17, 4
  5948: d78ddc33                 vrgather.vv v27, v29, v25
  5952: 570eed33                 vrgather.vv v28, v30, v26
  5956: d78dcd2f                 vxor.vv v27, v28, v27
  5960: d780b02f                 vxor.vv v1, v27, v1
  5964: 138e0f12                 addi t3, t6, 0x120
  5968: 870e0e02                 vle8.v v29, (t3)
  5972: 138e0f13                 addi t3, t6, 0x130
  5976: 070f0e02                 vle8.v v30, (t3)
  5980: d78cf827                 vand.vv v25, v31, v17
  5984: 573d12a3                 vsrl.vi v26, v17, 4
  5988: d78ddc33                 vrgather.vv v27, v29, v25
  5992: 570eed33                 vrgather.vv v28, v30, v26
  5996: d78dcd2f                 vxor.vv v27, v28, v27
  6000: d784b42f                 vxor.vv v9, v27, v9
  6004: 83338502                 ld t2, 0x28(a0)
  6008: b3836300                 add t2, t2, t1
  6012: 87880300                 vle8.v v17, (t2), v0.t
  6016: 138e0f14                 addi t3, t6, 0x140
  6020: 870e0e02                 vle8.v v29, (t3)
  6024: 138e0f15                 addi t3, t6, 0x150
  6028: 070f0e02                 vle8.v v30, (t3)
  6032: d78cf827                 vand.vv v25, v31, v17
  6036: 573d12a3                 vsrl.vi v26, v17, 4
  6040: d78ddc33                 vrgather.vv v27, v29, v25
  6044: 570eed33                 vrgather.vv v28, v30, v26
  6048: d78dcd2f                 vxor.vv v27, v28, v27
  6052: d780b02f                 vxor.vv v1, v27, v1
  6056: d78cf827                 vand.vv v25, v31, v17
  6060: 573d12a3                 vsrl.vi v26, v17, 4
  6064: d78ddc33                 vrgather.vv v27, v29, v25
  6068: 570eed33                 vrgather.vv v28, v30, v26
  6072: d78dcd2f                 vxor.vv v27, v28, v27
  6076: d784b42f                 vxor.vv v9, v27, v9
  6080: 83b30500                 ld t2, 0(a1)
  6084: b3836300                 add t2, t2, t1
  6088: a7800300                 vse8.v v1, (t2), v0.t
  6092: 83b38500                 ld t2, 8(a1)
  6096: b3836300                 add t2, t2, t1
  6100: a7840300                 vse8.v v9, (t2), v0.t
  6104: 33035300                 add t1, t1, t0
  6108: b3865640                 sub a3, a3, t0
  6112: 6344d000                 bgtz a3, 0x17e8
  6116: 6f00c024                 j 0x1a30
  6120: b3835640                 sub t2, a3, t0
  6124: 13def343                 srai t3, t2, 0x3f
  6128: b3f3c301                 and t2, t2, t3
  6132: b3837200                 add t2, t0, t2
  6136: d7a80852                 vid.v v17
  6140: 57c0136b                 vmsltu.vx v0, v17, t2
  6144: 83330500                 ld t2, 0(a0)
  6148: b3836300                 add t2, t2, t1
  6152: 87880300                 vle8.v v17, (t2), v0.t
  6156: 138e0f00                 mv t3, t6
  6160: 870e0e02                 vle8.v v29, (t3)
  6164: 138e0f01                 addi t3, t6, 0x10
  6168: 070f0e02                 vle8.v v30, (t3)
  6172: d78cf827                 vand.vv v25, v31, v17
  6176: 573d12a3                 vsrl.vi v26, v17, 4
  6180: d78ddc33                 vrgather.vv v27, v29, v25
  6184: 570eed33                 vrgather.vv v28, v30, v26
  6188: d78dcd2f                 vxor.vv v27, v28, v27
  6192: d7800d5e                 vmv.v.v v1, v27
  6196: 138e0f02                 addi t3, t6, 0x20
  6200: 870e0e02                 vle8.v v29, (t3)
  6204: 138e0f03                 addi t3, t6, 0x30
  6208: 070f0e02                 vle8.v v30, (t3)
  6212: d78cf827                 vand.vv v25, v31, v17
  6216: 573d12a3                 vsrl.vi v26, v17, 4
  6220: d78ddc33                 vrgather.vv v27, v29, v25
  6224: 570eed33                 vrgather.vv v28, v30, v26
  6228: d78dcd2f                 vxor.vv v27, v28, v27
  6232: d7840d5e                 vmv.v.v v9, v27
  6236: 83338500                 ld t2, 8(a0)
  6240: b3836300                 add t2, t2, t1
  6244: 87880300                 vle8.v v17, (t2), v0.t
  6248: 138e0f04                 addi t3, t6, 0x40
  6252: 870e0e02                 vle8.v v29, (t3)
  6256: 138e0f05                 addi t3, t6, 0x50
  6260: 070f0e02                 vle8.v v30, (t3)
  6264: d78cf827                 vand.vv v25, v31, v17
  6268: 573d12a3                 vsrl.vi v26, v17, 4
  6272: d78ddc33                 vrgather.vv v27, v29, v25
  6276: 570eed33                 vrgather.vv v28, v30, v26
  6280: d78dcd2f                 vxor.vv v27, v28, v27
  6284: d780b02f                 vxor.vv v1, v27, v1
  6288: 138e0f06                 addi t3, t6, 0x60
  6292: 870e0e02                 vle8.v v29, (t3)
  6296: 138e0f07                 addi t3, t6, 0x70
  6300: 070f0e02                 vle8.v v30, (t3)
  6304: d78cf827                 vand.vv v25, v31, v17
  6308: 573d12a3                 vsrl.vi v26, v17, 4
  6312: d78ddc33                 vrgather.vv v27, v29, v25
  6316: 570eed33                 vrgather.vv v28, v30, v26
  6320: d78dcd2f                 vxor.vv v27, v28, v27
  6324: d784b42f                 vxor.vv v9, v27, v9
  6328: 83330501                 ld t2, 0x10(a0)
  6332: b3836300                 add t2, t2, t1
  6336: 87880300                 vle8.v v17, (t2), v0.t
  6340: 138e0f08                 addi t3, t6, 0x80
  6344: 870e0e02                 vle8.v v29, (t3)
  6348: 138e0f09                 addi t3, t6, 0x90
  6352: 070f0e02                 vle8.v v30, (t3)
  6356: d78cf827                 vand.vv v25, v31, v17
  6360: 573d12a3                 vsrl.vi v26, v17, 4
  6364: d78ddc33                 vrgather.vv v27, v29, v25
  6368: 570eed33                 vrgather.vv v28, v30, v26
  6372: d78dcd2f                 vxor.vv v27, v28, v27
  6376: d780b02f                 vxor.vv v1, v27, v1
  6380: 138e0f0a                 addi t3, t6, 0xa0
  6384: 870e0e02                 vle8.v v29, (t3)
  6388: 138e0f0b                 addi t3, t6, 0xb0
  6392: 070f0e02                 vle8.v v30, (t3)
  6396: d78cf827                 vand.vv v25, v31, v17
  6400: 573d12a3                 vsrl.vi v26, v17, 4
  6404: d78ddc33                 vrgather.vv v27, v29, v25
  6408: 570eed33                 vrgather.vv v28, v30, v26
  6412: d78dcd2f                 vxor.vv v27, v28, v27
  6416: d784b42f                 vxor.vv v9, v27, v9
  6420: 83338501                 ld t2, 0x18(a0)
  6424: b3836300                 add t2, t2, t1
  6428: 87880300                 vle8.v v17, (t2), v0.t
  6432: 138e0f0c                 addi t3, t6, 0xc0
  6436: 870e0e02                 vle8.v v29, (t3)
  6440: 138e0f0d                 addi t3, t6, 0xd0
  6444: 070f0e02                 vle8.v v30, (t3)
  6448: d78cf827                 vand.vv v25, v31, v17
  6452: 573d12a3                 vsrl.vi v26, v17, 4
  6456: d78ddc33                 vrgather.vv v27, v29, v25
  6460: 570eed33                 vrgather.vv v28, v30, v26
  6464: d78dcd2f                 vxor.vv v27, v28, v27
  6468: d780b02f                 vxor.vv v1, v27, v1
  6472: 138e0f0e                 addi t3, t6, 0xe0
  6476: 870e0e02                 vle8.v v29, (t3)
  6480: 138e0f0f                 addi t3, t6, 0xf0
  6484: 070f0e02                 vle8.v v30, (t3)
  6488: d78cf827                 vand.vv v25, v31, v17
  6492: 573d12a3                 vsrl.vi v26, v17, 4
  6496: d78ddc33                 vrgather.vv v27, v29, v25
  6500: 570eed33                 vrgather.vv v28, v30, v26
  6504: d78dcd2f                 vxor.vv v27, v28, v27
  6508: d784b42f                 vxor.vv v9, v27, v9
  6512: 83330502                 ld t2, 0x20(a0)
  6516: b3836300                 add t2, t2, t1
  6520: 87880300                 vle8.v v17, (t2), v0.t
  6524: 138e0f10                 addi t3, t6, 0x100
  6528: 870e0e02                 vle8.v v29, (t3)
  6532: 138e0f11                 addi t3, t6, 0x110
  6536: 070f0e02                 vle8.v v30, (t3)
  6540: d78cf827                 vand.vv v25, v31, v17
  6544: 573d12a3                 vsrl.vi v26, v17, 4
  6548: d78ddc33                 vrgather.vv v27, v29, v25
  6552: 570eed33                 vrgather.vv v28, v30, v26
  6556: d78dcd2f                 vxor.vv v27, v28, v27
  6560: d780b02f                 vxor.vv v1, v27, v1
  6564: 138e0f12                 addi t3, t6, 0x120
  6568: 870e0e02                 vle8.v v29, (t3)
  6572: 138e0f13                 addi t3, t6, 0x130
  6576: 070f0e02                 vle8.v v30, (t3)
  6580: d78cf827                 vand.vv v25, v31, v17
  6584: 573d12a3                 vsrl.vi v26, v17, 4
  6588: d78ddc33                 vrgather.vv v27, v29, v25
  6592: 570eed33                 vrgather.vv v28, v30, v26
  6596: d78dcd2f                 vxor.vv v27, v28, v27
  6600: d784b42f                 vxor.vv v9, v27, v9
  6604: 83338502                 ld t2, 0x28(a0)
  6608: b3836300                 add t2, t2, t1
  6612: 87880300                 vle8.v v17, (t2), v0.t
  6616: 138e0f14                 addi t3, t6, 0x140
  6620: 870e0e02                 vle8.v v29, (t3)
  6624: 138e0f15                 addi t3, t6, 0x150
  6628: 070f0e02                 vle8.v v30, (t3)
  6632: d78cf827                 vand.vv v25, v31, v17
  6636: 573d12a3                 vsrl.vi v26, v17, 4
  6640: d78ddc33                 vrgather.vv v27, v29, v25
  6644: 570eed33                 vrgather.vv v28, v30, v26
  6648: d78dcd2f                 vxor.vv v27, v28, v27
  6652: d780b02f                 vxor.vv v1, v27, v1
  6656: d78cf827                 vand.vv v25, v31, v17
  6660: 573d12a3                 vsrl.vi v26, v17, 4
  6664: d78ddc33                 vrgather.vv v27, v29, v25
  6668: 570eed33                 vrgather.vv v28, v30, v26
  6672: d78dcd2f                 vxor.vv v27, v28, v27
  6676: d784b42f                 vxor.vv v9, v27, v9
  6680: 83b30500                 ld t2, 0(a1)
  6684: b3836300                 add t2, t2, t1
  6688: a7800300                 vse8.v v1, (t2), v0.t
  6692: 83b38500                 ld t2, 8(a1)
  6696: b3836300                 add t2, t2, t1
  6700: a7840300                 vse8.v v9, (t2), v0.t
  6704: 33035300                 add t1, t1, t0
  6708: b3865640                 sub a3, a3, t0
  6712: 6344d000                 bgtz a3, 0x1a40
  6716: 6f00c024                 j 0x1c88
  6720: b3835640                 sub t2, a3, t0
  6724: 13def343                 srai t3, t2, 0x3f
  6728: b3f3c301                 and t2, t2, t3
  6732: b3837200                 add t2, t0, t2
  6736: d7a80852                 vid.v v17
  6740: 57c0136b                 vmsltu.vx v0, v17, t2
  6744: 83330500                 ld t2, 0(a0)
  6748: b3836300                 add t2, t2, t1
  6752: 87880300                 vle8.v v17, (t2), v0.t
  6756: 138e0f00                 mv t3, t6
  6760: 870e0e02                 vle8.v v29, (t3)
  6764: 138e0f01                 addi t3, t6, 0x10
  6768: 070f0e02                 vle8.v v30, (t3)
  6772: d78cf827                 vand.vv v25, v31, v17
  6776: 573d12a3                 vsrl.vi v26, v17, 4
  6780: d78ddc33                 vrgather.vv v27, v29, v25
  6784: 570eed33                 vrgather.vv v28, v30, v26
  6788: d78dcd2f                 vxor.vv v27, v28, v27
  6792: d7800d5e                 vmv.v.v v1, v27
  6796: 138e0f02                 addi t3, t6, 0x20
  6800: 870e0e02                 vle8.v v29, (t3)
  6804: 138e0f03                 addi t3, t6, 0x30
  6808: 070f0e02                 vle8.v v30, (t3)
  6812: d78cf827                 vand.vv v25, v31, v17
  6816: 573d12a3                 vsrl.vi v26, v17, 4
  6820: d78ddc33                 vrgather.vv v27, v29, v25
  6824: 570eed33                 vrgather.vv v28, v30, v26
  6828: d78dcd2f                 vxor.vv v27, v28, v27
  6832: d7840d5e                 vmv.v.v v9, v27
  6836: 83338500                 ld t2, 8(a0)
  6840: b3836300                 add t2, t2, t1
  6844: 87880300                 vle8.v v17, (t2), v0.t
  6848: 138e0f04                 addi t3, t6, 0x40
  6852: 870e0e02                 vle8.v v29, (t3)
  6856: 138e0f05                 addi t3, t6, 0x50
  6860: 070f0e02                 vle8.v v30, (t3)
  6864: d78cf827                 vand.vv v25, v31, v17
  6868: 573d12a3                 vsrl.vi v26, v17, 4
  6872: d78ddc33                 vrgather.vv v27, v29, v25
  6876: 570eed33                 vrgather.vv v28, v30, v26
  6880: d78dcd2f                 vxor.vv v27, v28, v27
  6884: d780b02f                 vxor.vv v1, v27, v1
  6888: 138e0f06                 addi t3, t6, 0x60
  6892: 870e0e02                 vle8.v v29, (t3)
  6896: 138e0f07                 addi t3, t6, 0x70
  6900: 070f0e02                 vle8.v v30, (t3)
  6904: d78cf827                 vand.vv v25, v31, v17
  6908: 573d12a3                 vsrl.vi v26, v17, 4
  6912: d78ddc33                 vrgather.vv v27, v29, v25
  6916: 570eed33                 vrgather.vv v28, v30, v26
  6920: d78dcd2f                 vxor.vv v27, v28, v27
  6924: d784b42f                 vxor.vv v9, v27, v9
  6928: 83330501                 ld t2, 0x10(a0)
  6932: b3836300                 add t2, t2, t1
  6936: 87880300                 vle8.v v17, (t2), v0.t
  6940: 138e0f08                 addi t3, t6, 0x80
  6944: 870e0e02                 vle8.v v29, (t3)
  6948: 138e0f09                 addi t3, t6, 0x90
  6952: 070f0e02                 vle8.v v30, (t3)
  6956: d78cf827                 vand.vv v25, v31, v17
  6960: 573d12a3                 vsrl.vi v26, v17, 4
  6964: d78ddc33                 vrgather.vv v27, v29, v25
  6968: 570eed33                 vrgather.vv v28, v30, v26
  6972: d78dcd2f                 vxor.vv v27, v28, v27
  6976: d780b02f                 vxor.vv v1, v27, v1
  6980: 138e0f0a                 addi t3, t6, 0xa0
  6984: 870e0e02                 vle8.v v29, (t3)
  6988: 138e0f0b                 addi t3, t6, 0xb0
  6992: 070f0e02                 vle8.v v30, (t3)
  6996: d78cf827                 vand.vv v25, v31, v17
  7000: 573d12a3                 vsrl.vi v26, v17, 4
  7004: d78ddc33                 vrgather.vv v27, v29, v25
  7008: 570eed33                 vrgather.vv v28, v30, v26
  7012: d78dcd2f                 vxor.vv v27, v28, v27
  7016: d784b42f                 vxor.vv v9, v27, v9
  7020: 83338501                 ld t2, 0x18(a0)
  7024: b3836300                 add t2, t2, t1
  7028: 87880300                 vle8.v v17, (t2), v0.t
  7032: 138e0f0c                 addi t3, t6, 0xc0
  7036: 870e0e02                 vle8.v v29, (t3)
  7040: 138e0f0d                 addi t3, t6, 0xd0
  7044: 070f0e02                 vle8.v v30, (t3)
  7048: d78cf827                 vand.vv v25, v31, v17
  7052: 573d12a3                 vsrl.vi v26, v17, 4
  7056: d78ddc33                 vrgather.vv v27, v29, v25
  7060: 570eed33                 vrgather.vv v28, v30, v26
  7064: d78dcd2f                 vxor.vv v27, v28, v27
  7068: d780b02f                 vxor.vv v1, v27, v1
  7072: 138e0f0e                 addi t3, t6, 0xe0
  7076: 870e0e02                 vle8.v v29, (t3)
  7080: 138e0f0f                 addi t3, t6, 0xf0
  7084: 070f0e02                 vle8.v v30, (t3)
  7088: d78cf827                 vand.vv v25, v31, v17
  7092: 573d12a3                 vsrl.vi v26, v17, 4
  7096: d78ddc33                 vrgather.vv v27, v29, v25
  7100: 570eed33                 vrgather.vv v28, v30, v26
  7104: d78dcd2f                 vxor.vv v27, v28, v27
  7108: d784b42f                 vxor.vv v9, v27, v9
  7112: 83330502                 ld t2, 0x20(a0)
  7116: b3836300                 add t2, t2, t1
  7120: 87880300                 vle8.v v17, (t2), v0.t
  7124: 138e0f10                 addi t3, t6, 0x100
  7128: 870e0e02                 vle8.v v29, (t3)
  7132: 138e0f11                 addi t3, t6, 0x110
  7136: 070f0e02                 vle8.v v30, (t3)
  7140: d78cf827                 vand.vv v25, v31, v17
  7144: 573d12a3                 vsrl.vi v26, v17, 4
  7148: d78ddc33                 vrgather.vv v27, v29, v25
  7152: 570eed33                 vrgather.vv v28, v30, v26
  7156: d78dcd2f                 vxor.vv v27, v28, v27
  7160: d780b02f                 vxor.vv v1, v27, v1
  7164: 138e0f12                 addi t3, t6, 0x120
  7168: 870e0e02                 vle8.v v29, (t3)
  7172: 138e0f13                 addi t3, t6, 0x130
  7176: 070f0e02                 vle8.v v30, (t3)
  7180: d78cf827                 vand.vv v25, v31, v17
  7184: 573d12a3                 vsrl.vi v26, v17, 4
  7188: d78ddc33                 vrgather.vv v27, v29, v25
  7192: 570eed33                 vrgather.vv v28, v30, v26
  7196: d78dcd2f                 vxor.vv v27, v28, v27
  7200: d784b42f                 vxor.vv v9, v27, v9
  7204: 83338502                 ld t2, 0x28(a0)
  7208: b3836300                 add t2, t2, t1
  7212: 87880300                 vle8.v v17, (t2), v0.t
  7216: 138e0f14                 addi t3, t6, 0x140
  7220: 870e0e02                 vle8.v v29, (t3)
  7224: 138e0f15                 addi t3, t6, 0x150
  7228: 070f0e02                 vle8.v v30, (t3)
  7232: d78cf827                 vand.vv v25, v31, v17
  7236: 573d12a3                 vsrl.vi v26, v17, 4
  7240: d78ddc33                 vrgather.vv v27, v29, v25
  7244: 570eed33                 vrgather.vv v28, v30, v26
  7248: d78dcd2f                 vxor.vv v27, v28, v27
  7252: d780b02f                 vxor.vv v1, v27, v1
  7256: d78cf827                 vand.vv v25, v31, v17
  7260: 573d12a3                 vsrl.vi v26, v17, 4
  7264: d78ddc33                 vrgather.vv v27, v29, v25
  7268: 570eed33                 vrgather.vv v28, v30, v26
  7272: d78dcd2f                 vxor.vv v27, v28, v27
  7276: d784b42f                 vxor.vv v9, v27, v9
  7280: 83b30500                 ld t2, 0(a1)
  7284: b3836300                 add t2, t2, t1
  7288: a7800300                 vse8.v v1, (t2), v0.t
  7292: 83b38500                 ld t2, 8(a1)
  7296: b3836300                 add t2, t2, t1
  7300: a7840300                 vse8.v v9, (t2), v0.t
  7304: 33035300                 add t1, t1, t0
  7308: b3865640                 sub a3, a3, t0
  7312: 6344d000                 bgtz a3, 0x1c98
  7316: 6f00c024                 j 0x1ee0
  7320: b3835640                 sub t2, a3, t0
  7324: 13def343                 srai t3, t2, 0x3f
  7328: b3f3c301                 and t2, t2, t3
  7332: b3837200                 add t2, t0, t2
  7336: d7a80852                 vid.v v17
  7340: 57c0136b                 vmsltu.vx v0, v17, t2
  7344: 83330500                 ld t2, 0(a0)
  7348: b3836300                 add t2, t2, t1
  7352: 87880300                 vle8.v v17, (t2), v0.t
  7356: 138e0f00                 mv t3, t6
  7360: 870e0e02                 vle8.v v29, (t3)
  7364: 138e0f01                 addi t3, t6, 0x10
  7368: 070f0e02                 vle8.v v30, (t3)
  7372: d78cf827                 vand.vv v25, v31, v17
  7376: 573d12a3                 vsrl.vi v26, v17, 4
  7380: d78ddc33                 vrgather.vv v27, v29, v25
  7384: 570eed33                 vrgather.vv v28, v30, v26
  7388: d78dcd2f                 vxor.vv v27, v28, v27
  7392: d7800d5e                 vmv.v.v v1, v27
  7396: 138e0f02                 addi t3, t6, 0x20
  7400: 870e0e02                 vle8.v v29, (t3)
  7404: 138e0f03                 addi t3, t6, 0x30
  7408: 070f0e02                 vle8.v v30, (t3)
  7412: d78cf827                 vand.vv v25, v31, v17
  7416: 573d12a3                 vsrl.vi v26, v17, 4
  7420: d78ddc33                 vrgather.vv v27, v29, v25
  7424: 570eed33                 vrgather.vv v28, v30, v26
  7428: d78dcd2f                 vxor.vv v27, v28, v27
  7432: d7840d5e                 vmv.v.v v9, v27
  7436: 83338500                 ld t2, 8(a0)
  7440: b3836300                 add t2, t2, t1
  7444: 87880300                 vle8.v v17, (t2), v0.t
  7448: 138e0f04                 addi t3, t6, 0x40
  7452: 870e0e02                 vle8.v v29, (t3)
  7456: 138e0f05                 addi t3, t6, 0x50
  7460: 070f0e02                 vle8.v v30, (t3)
  7464: d78cf827                 vand.vv v25, v31, v17
  7468: 573d12a3                 vsrl.vi v26, v17, 4
  7472: d78ddc33                 vrgather.vv v27, v29, v25
  7476: 570eed33                 vrgather.vv v28, v30, v26
  7480: d78dcd2f                 vxor.vv v27, v28, v27
  7484: d780b02f                 vxor.vv v1, v27, v1
  7488: 138e0f06                 addi t3, t6, 0x60
  7492: 870e0e02                 vle8.v v29, (t3)
  7496: 138e0f07                 addi t3, t6, 0x70
  7500: 070f0e02                 vle8.v v30, (t3)
  7504: d78cf827                 vand.vv v25, v31, v17
  7508: 573d12a3                 vsrl.vi v26, v17, 4
  7512: d78ddc33                 vrgather.vv v27, v29, v25
  7516: 570eed33                 vrgather.vv v28, v30, v26
  7520: d78dcd2f                 vxor.vv v27, v28, v27
  7524: d784b42f                 vxor.vv v9, v27, v9
  7528: 83330501                 ld t2, 0x10(a0)
  7532: b3836300                 add t2, t2, t1
  7536: 87880300                 vle8.v v17, (t2), v0.t
  7540: 138e0f08                 addi t3, t6, 0x80
  7544: 870e0e02                 vle8.v v29, (t3)
  7548: 138e0f09                 addi t3, t6, 0x90
  7552: 070f0e02                 vle8.v v30, (t3)
  7556: d78cf827                 vand.vv v25, v31, v17
  7560: 573d12a3                 vsrl.vi v26, v17, 4
  7564: d78ddc33                 vrgather.vv v27, v29, v25
  7568: 570eed33                 vrgather.vv v28, v30, v26
  7572: d78dcd2f                 vxor.vv v27, v28, v27
  7576: d780b02f                 vxor.vv v1, v27, v1
  7580: 138e0f0a                 addi t3, t6, 0xa0
  7584: 870e0e02                 vle8.v v29, (t3)
  7588: 138e0f0b                 addi t3, t6, 0xb0
  7592: 070f0e02                 vle8.v v30, (t3)
  7596: d78cf827                 vand.vv v25, v31, v17
  7600: 573d12a3                 vsrl.vi v26, v17, 4
  7604: d78ddc33                 vrgather.vv v27, v29, v25
  7608: 570eed33                 vrgather.vv v28, v30, v26
  7612: d78dcd2f                 vxor.vv v27, v28, v27
  7616: d784b42f                 vxor.vv v9, v27, v9
  7620: 83338501                 ld t2, 0x18(a0)
  7624: b3836300                 add t2, t2, t1
  7628: 87880300                 vle8.v v17, (t2), v0.t
  7632: 138e0f0c                 addi t3, t6, 0xc0
  7636: 870e0e02                 vle8.v v29, (t3)
  7640: 138e0f0d                 addi t3, t6, 0xd0
  7644: 070f0e02                 vle8.v v30, (t3)
  7648: d78cf827                 vand.vv v25, v31, v17
  7652: 573d12a3                 vsrl.vi v26, v17, 4
  7656: d78ddc33                 vrgather.vv v27, v29, v25
  7660: 570eed33                 vrgather.vv v28, v30, v26
  7664: d78dcd2f                 vxor.vv v27, v28, v27
  7668: d780b02f                 vxor.vv v1, v27, v1
  7672: 138e0f0e                 addi t3, t6, 0xe0
  7676: 870e0e02                 vle8.v v29, (t3)
  7680: 138e0f0f                 addi t3, t6, 0xf0
  7684: 070f0e02                 vle8.v v30, (t3)
  7688: d78cf827                 vand.vv v25, v31, v17
  7692: 573d12a3                 vsrl.vi v26, v17, 4
  7696: d78ddc33                 vrgather.vv v27, v29, v25
  7700: 570eed33                 vrgather.vv v28, v30, v26
  7704: d78dcd2f                 vxor.vv v27, v28, v27
  7708: d784b42f                 vxor.vv v9, v27, v9
  7712: 83330502                 ld t2, 0x20(a0)
  7716: b3836300                 add t2, t2, t1
  7720: 87880300                 vle8.v v17, (t2), v0.t
  7724: 138e0f10                 addi t3, t6, 0x100
  7728: 870e0e02                 vle8.v v29, (t3)
  7732: 138e0f11                 addi t3, t6, 0x110
  7736: 070f0e02                 vle8.v v30, (t3)
  7740: d78cf827                 vand.vv v25, v31, v17
  7744: 573d12a3                 vsrl.vi v26, v17, 4
  7748: d78ddc33                 vrgather.vv v27, v29, v25
  7752: 570eed33                 vrgather.vv v28, v30, v26
  7756: d78dcd2f                 vxor.vv v27, v28, v27
  7760: d780b02f                 vxor.vv v1, v27, v1
  7764: 138e0f12                 addi t3, t6, 0x120
  7768: 870e0e02                 vle8.v v29, (t3)
  7772: 138e0f13                 addi t3, t6, 0x130
  7776: 070f0e02                 vle8.v v30, (t3)
  7780: d78cf827                 vand.vv v25, v31, v17
  7784: 573d12a3                 vsrl.vi v26, v17, 4
  7788: d78ddc33                 vrgather.vv v27, v29, v25
  7792: 570eed33                 vrgather.vv v28, v30, v26
  7796: d78dcd2f                 vxor.vv v27, v28, v27
  7800: d784b42f                 vxor.vv v9, v27, v9
  7804: 83338502                 ld t2, 0x28(a0)
  7808: b3836300                 add t2, t2, t1
  7812: 87880300                 vle8.v v17, (t2), v0.t
  7816: 138e0f14                 addi t3, t6, 0x140
  7820: 870e0e02                 vle8.v v29, (t3)
  7824: 138e0f15                 addi t3, t6, 0x150
  7828: 070f0e02                 vle8.v v30, (t3)
  7832: d78cf827                 vand.vv v25, v31, v17
  7836: 573d12a3                 vsrl.vi v26, v17, 4
  7840: d78ddc33                 vrgather.vv v27, v29, v25
  7844: 570eed33                 vrgather.vv v28, v30, v26
  7848: d78dcd2f                 vxor.vv v27, v28, v27
  7852: d780b02f                 vxor.vv v1, v27, v1
  7856: d78cf827                 vand.vv v25, v31, v17
  7860: 573d12a3                 vsrl.vi v26, v17, 4
  7864: d78ddc33                 vrgather.vv v27, v29, v25
  7868: 570eed33                 vrgather.vv v28, v30, v26
  7872: d78dcd2f                 vxor.vv v27, v28, v27
  7876: d784b42f                 vxor.vv v9, v27, v9
  7880: 83b30500                 ld t2, 0(a1)
  7884: b3836300                 add t2, t2, t1
  7888: a7800300                 vse8.v v1, (t2), v0.t
  7892: 83b38500                 ld t2, 8(a1)
  7896: b3836300                 add t2, t2, t1
  7900: a7840300                 vse8.v v9, (t2), v0.t
  7904: 33035300                 add t1, t1, t0
  7908: b3865640                 sub a3, a3, t0
  7912: 67800000                 ret 
  7916: 0000                     unimp 
  7918: 0000                     unimp 
