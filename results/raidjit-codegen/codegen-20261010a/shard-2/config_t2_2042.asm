; config_t2_2042 target=rvv-zvbc32e twin=nt+pipe topo=rs(6,3)/enc op=encode U=4
; code_size=2448 twins=["nt","pipe"] uops_per_byte=7.765625
     0: 63140600                 bnez a2, 8
     4: 6f005018                 j 0x988
     8: 971f0000                 auipc t6, 1
    12: 938f8f98                 addi t6, t6, -0x678
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574d0e5e                 vmv.v.x v26, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93962200                 slli a3, t0, 2
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00006a                 j 0x6d0
    52: 83330500                 ld t2, 0(a0)
    56: b3836300                 add t2, t2, t1
    60: 07880302                 vle8.v v16, (t2)
    64: b3835300                 add t2, t2, t0
    68: 87880302                 vle8.v v17, (t2)
    72: b3835300                 add t2, t2, t0
    76: 07890302                 vle8.v v18, (t2)
    80: b3835300                 add t2, t2, t0
    84: 87890302                 vle8.v v19, (t2)
    88: 13000000                 nop 
    92: 13000000                 nop 
    96: 5700085e                 vmv.v.v v0, v16
   100: d780085e                 vmv.v.v v1, v17
   104: 5701095e                 vmv.v.v v2, v18
   108: d781095e                 vmv.v.v v3, v19
   112: 138e0f00                 mv t3, t6
   116: 070c0e02                 vle8.v v24, (t3)
   120: 138e0f01                 addi t3, t6, 0x10
   124: 870c0e02                 vle8.v v25, (t3)
   128: 570aa827                 vand.vv v20, v26, v16
   132: d73a02a3                 vsrl.vi v21, v16, 4
   136: 570b8a33                 vrgather.vv v22, v24, v20
   140: d78b9a33                 vrgather.vv v23, v25, v21
   144: 570b7b2f                 vxor.vv v22, v23, v22
   148: 57020b5e                 vmv.v.v v4, v22
   152: 578aa827                 vand.vv v20, v26, v17
   156: d73a12a3                 vsrl.vi v21, v17, 4
   160: 570b8a33                 vrgather.vv v22, v24, v20
   164: d78b9a33                 vrgather.vv v23, v25, v21
   168: 570b7b2f                 vxor.vv v22, v23, v22
   172: d7020b5e                 vmv.v.v v5, v22
   176: 570aa927                 vand.vv v20, v26, v18
   180: d73a22a3                 vsrl.vi v21, v18, 4
   184: 570b8a33                 vrgather.vv v22, v24, v20
   188: d78b9a33                 vrgather.vv v23, v25, v21
   192: 570b7b2f                 vxor.vv v22, v23, v22
   196: 57030b5e                 vmv.v.v v6, v22
   200: 578aa927                 vand.vv v20, v26, v19
   204: d73a32a3                 vsrl.vi v21, v19, 4
   208: 570b8a33                 vrgather.vv v22, v24, v20
   212: d78b9a33                 vrgather.vv v23, v25, v21
   216: 570b7b2f                 vxor.vv v22, v23, v22
   220: d7030b5e                 vmv.v.v v7, v22
   224: 138e0f02                 addi t3, t6, 0x20
   228: 070c0e02                 vle8.v v24, (t3)
   232: 138e0f03                 addi t3, t6, 0x30
   236: 870c0e02                 vle8.v v25, (t3)
   240: 570aa827                 vand.vv v20, v26, v16
   244: d73a02a3                 vsrl.vi v21, v16, 4
   248: 570b8a33                 vrgather.vv v22, v24, v20
   252: d78b9a33                 vrgather.vv v23, v25, v21
   256: 570b7b2f                 vxor.vv v22, v23, v22
   260: 57040b5e                 vmv.v.v v8, v22
   264: 578aa827                 vand.vv v20, v26, v17
   268: d73a12a3                 vsrl.vi v21, v17, 4
   272: 570b8a33                 vrgather.vv v22, v24, v20
   276: d78b9a33                 vrgather.vv v23, v25, v21
   280: 570b7b2f                 vxor.vv v22, v23, v22
   284: d7040b5e                 vmv.v.v v9, v22
   288: 570aa927                 vand.vv v20, v26, v18
   292: d73a22a3                 vsrl.vi v21, v18, 4
   296: 570b8a33                 vrgather.vv v22, v24, v20
   300: d78b9a33                 vrgather.vv v23, v25, v21
   304: 570b7b2f                 vxor.vv v22, v23, v22
   308: 57050b5e                 vmv.v.v v10, v22
   312: 578aa927                 vand.vv v20, v26, v19
   316: d73a32a3                 vsrl.vi v21, v19, 4
   320: 570b8a33                 vrgather.vv v22, v24, v20
   324: d78b9a33                 vrgather.vv v23, v25, v21
   328: 570b7b2f                 vxor.vv v22, v23, v22
   332: d7050b5e                 vmv.v.v v11, v22
   336: 83338500                 ld t2, 8(a0)
   340: b3836300                 add t2, t2, t1
   344: 07860302                 vle8.v v12, (t2)
   348: b3835300                 add t2, t2, t0
   352: 87860302                 vle8.v v13, (t2)
   356: b3835300                 add t2, t2, t0
   360: 07870302                 vle8.v v14, (t2)
   364: b3835300                 add t2, t2, t0
   368: 87870302                 vle8.v v15, (t2)
   372: 5700c02e                 vxor.vv v0, v12, v0
   376: d780d02e                 vxor.vv v1, v13, v1
   380: 5701e12e                 vxor.vv v2, v14, v2
   384: d781f12e                 vxor.vv v3, v15, v3
   388: 138e0f02                 addi t3, t6, 0x20
   392: 070c0e02                 vle8.v v24, (t3)
   396: 138e0f03                 addi t3, t6, 0x30
   400: 870c0e02                 vle8.v v25, (t3)
   404: 570aa627                 vand.vv v20, v26, v12
   408: d73ac2a2                 vsrl.vi v21, v12, 4
   412: 570b8a33                 vrgather.vv v22, v24, v20
   416: d78b9a33                 vrgather.vv v23, v25, v21
   420: 570b7b2f                 vxor.vv v22, v23, v22
   424: 5702622f                 vxor.vv v4, v22, v4
   428: 578aa627                 vand.vv v20, v26, v13
   432: d73ad2a2                 vsrl.vi v21, v13, 4
   436: 570b8a33                 vrgather.vv v22, v24, v20
   440: d78b9a33                 vrgather.vv v23, v25, v21
   444: 570b7b2f                 vxor.vv v22, v23, v22
   448: d782622f                 vxor.vv v5, v22, v5
   452: 570aa727                 vand.vv v20, v26, v14
   456: d73ae2a2                 vsrl.vi v21, v14, 4
   460: 570b8a33                 vrgather.vv v22, v24, v20
   464: d78b9a33                 vrgather.vv v23, v25, v21
   468: 570b7b2f                 vxor.vv v22, v23, v22
   472: 5703632f                 vxor.vv v6, v22, v6
   476: 578aa727                 vand.vv v20, v26, v15
   480: d73af2a2                 vsrl.vi v21, v15, 4
   484: 570b8a33                 vrgather.vv v22, v24, v20
   488: d78b9a33                 vrgather.vv v23, v25, v21
   492: 570b7b2f                 vxor.vv v22, v23, v22
   496: d783632f                 vxor.vv v7, v22, v7
   500: 138e0f04                 addi t3, t6, 0x40
   504: 070c0e02                 vle8.v v24, (t3)
   508: 138e0f05                 addi t3, t6, 0x50
   512: 870c0e02                 vle8.v v25, (t3)
   516: 570aa627                 vand.vv v20, v26, v12
   520: d73ac2a2                 vsrl.vi v21, v12, 4
   524: 570b8a33                 vrgather.vv v22, v24, v20
   528: d78b9a33                 vrgather.vv v23, v25, v21
   532: 570b7b2f                 vxor.vv v22, v23, v22
   536: 5704642f                 vxor.vv v8, v22, v8
   540: 578aa627                 vand.vv v20, v26, v13
   544: d73ad2a2                 vsrl.vi v21, v13, 4
   548: 570b8a33                 vrgather.vv v22, v24, v20
   552: d78b9a33                 vrgather.vv v23, v25, v21
   556: 570b7b2f                 vxor.vv v22, v23, v22
   560: d784642f                 vxor.vv v9, v22, v9
   564: 570aa727                 vand.vv v20, v26, v14
   568: d73ae2a2                 vsrl.vi v21, v14, 4
   572: 570b8a33                 vrgather.vv v22, v24, v20
   576: d78b9a33                 vrgather.vv v23, v25, v21
   580: 570b7b2f                 vxor.vv v22, v23, v22
   584: 5705652f                 vxor.vv v10, v22, v10
   588: 578aa727                 vand.vv v20, v26, v15
   592: d73af2a2                 vsrl.vi v21, v15, 4
   596: 570b8a33                 vrgather.vv v22, v24, v20
   600: d78b9a33                 vrgather.vv v23, v25, v21
   604: 570b7b2f                 vxor.vv v22, v23, v22
   608: d785652f                 vxor.vv v11, v22, v11
   612: 83330501                 ld t2, 0x10(a0)
   616: b3836300                 add t2, t2, t1
   620: 07860302                 vle8.v v12, (t2)
   624: b3835300                 add t2, t2, t0
   628: 87860302                 vle8.v v13, (t2)
   632: b3835300                 add t2, t2, t0
   636: 07870302                 vle8.v v14, (t2)
   640: b3835300                 add t2, t2, t0
   644: 87870302                 vle8.v v15, (t2)
   648: 5700c02e                 vxor.vv v0, v12, v0
   652: d780d02e                 vxor.vv v1, v13, v1
   656: 5701e12e                 vxor.vv v2, v14, v2
   660: d781f12e                 vxor.vv v3, v15, v3
   664: 138e0f04                 addi t3, t6, 0x40
   668: 070c0e02                 vle8.v v24, (t3)
   672: 138e0f05                 addi t3, t6, 0x50
   676: 870c0e02                 vle8.v v25, (t3)
   680: 570aa627                 vand.vv v20, v26, v12
   684: d73ac2a2                 vsrl.vi v21, v12, 4
   688: 570b8a33                 vrgather.vv v22, v24, v20
   692: d78b9a33                 vrgather.vv v23, v25, v21
   696: 570b7b2f                 vxor.vv v22, v23, v22
   700: 5702622f                 vxor.vv v4, v22, v4
   704: 578aa627                 vand.vv v20, v26, v13
   708: d73ad2a2                 vsrl.vi v21, v13, 4
   712: 570b8a33                 vrgather.vv v22, v24, v20
   716: d78b9a33                 vrgather.vv v23, v25, v21
   720: 570b7b2f                 vxor.vv v22, v23, v22
   724: d782622f                 vxor.vv v5, v22, v5
   728: 570aa727                 vand.vv v20, v26, v14
   732: d73ae2a2                 vsrl.vi v21, v14, 4
   736: 570b8a33                 vrgather.vv v22, v24, v20
   740: d78b9a33                 vrgather.vv v23, v25, v21
   744: 570b7b2f                 vxor.vv v22, v23, v22
   748: 5703632f                 vxor.vv v6, v22, v6
   752: 578aa727                 vand.vv v20, v26, v15
   756: d73af2a2                 vsrl.vi v21, v15, 4
   760: 570b8a33                 vrgather.vv v22, v24, v20
   764: d78b9a33                 vrgather.vv v23, v25, v21
   768: 570b7b2f                 vxor.vv v22, v23, v22
   772: d783632f                 vxor.vv v7, v22, v7
   776: 138e0f06                 addi t3, t6, 0x60
   780: 070c0e02                 vle8.v v24, (t3)
   784: 138e0f07                 addi t3, t6, 0x70
   788: 870c0e02                 vle8.v v25, (t3)
   792: 570aa627                 vand.vv v20, v26, v12
   796: d73ac2a2                 vsrl.vi v21, v12, 4
   800: 570b8a33                 vrgather.vv v22, v24, v20
   804: d78b9a33                 vrgather.vv v23, v25, v21
   808: 570b7b2f                 vxor.vv v22, v23, v22
   812: 5704642f                 vxor.vv v8, v22, v8
   816: 578aa627                 vand.vv v20, v26, v13
   820: d73ad2a2                 vsrl.vi v21, v13, 4
   824: 570b8a33                 vrgather.vv v22, v24, v20
   828: d78b9a33                 vrgather.vv v23, v25, v21
   832: 570b7b2f                 vxor.vv v22, v23, v22
   836: d784642f                 vxor.vv v9, v22, v9
   840: 570aa727                 vand.vv v20, v26, v14
   844: d73ae2a2                 vsrl.vi v21, v14, 4
   848: 570b8a33                 vrgather.vv v22, v24, v20
   852: d78b9a33                 vrgather.vv v23, v25, v21
   856: 570b7b2f                 vxor.vv v22, v23, v22
   860: 5705652f                 vxor.vv v10, v22, v10
   864: 578aa727                 vand.vv v20, v26, v15
   868: d73af2a2                 vsrl.vi v21, v15, 4
   872: 570b8a33                 vrgather.vv v22, v24, v20
   876: d78b9a33                 vrgather.vv v23, v25, v21
   880: 570b7b2f                 vxor.vv v22, v23, v22
   884: d785652f                 vxor.vv v11, v22, v11
   888: 83338501                 ld t2, 0x18(a0)
   892: b3836300                 add t2, t2, t1
   896: 07860302                 vle8.v v12, (t2)
   900: b3835300                 add t2, t2, t0
   904: 87860302                 vle8.v v13, (t2)
   908: b3835300                 add t2, t2, t0
   912: 07870302                 vle8.v v14, (t2)
   916: b3835300                 add t2, t2, t0
   920: 87870302                 vle8.v v15, (t2)
   924: 5700c02e                 vxor.vv v0, v12, v0
   928: d780d02e                 vxor.vv v1, v13, v1
   932: 5701e12e                 vxor.vv v2, v14, v2
   936: d781f12e                 vxor.vv v3, v15, v3
   940: 138e0f08                 addi t3, t6, 0x80
   944: 070c0e02                 vle8.v v24, (t3)
   948: 138e0f09                 addi t3, t6, 0x90
   952: 870c0e02                 vle8.v v25, (t3)
   956: 570aa627                 vand.vv v20, v26, v12
   960: d73ac2a2                 vsrl.vi v21, v12, 4
   964: 570b8a33                 vrgather.vv v22, v24, v20
   968: d78b9a33                 vrgather.vv v23, v25, v21
   972: 570b7b2f                 vxor.vv v22, v23, v22
   976: 5702622f                 vxor.vv v4, v22, v4
   980: 578aa627                 vand.vv v20, v26, v13
   984: d73ad2a2                 vsrl.vi v21, v13, 4
   988: 570b8a33                 vrgather.vv v22, v24, v20
   992: d78b9a33                 vrgather.vv v23, v25, v21
   996: 570b7b2f                 vxor.vv v22, v23, v22
  1000: d782622f                 vxor.vv v5, v22, v5
  1004: 570aa727                 vand.vv v20, v26, v14
  1008: d73ae2a2                 vsrl.vi v21, v14, 4
  1012: 570b8a33                 vrgather.vv v22, v24, v20
  1016: d78b9a33                 vrgather.vv v23, v25, v21
  1020: 570b7b2f                 vxor.vv v22, v23, v22
  1024: 5703632f                 vxor.vv v6, v22, v6
  1028: 578aa727                 vand.vv v20, v26, v15
  1032: d73af2a2                 vsrl.vi v21, v15, 4
  1036: 570b8a33                 vrgather.vv v22, v24, v20
  1040: d78b9a33                 vrgather.vv v23, v25, v21
  1044: 570b7b2f                 vxor.vv v22, v23, v22
  1048: d783632f                 vxor.vv v7, v22, v7
  1052: 138e0f00                 mv t3, t6
  1056: 070c0e02                 vle8.v v24, (t3)
  1060: 138e0f01                 addi t3, t6, 0x10
  1064: 870c0e02                 vle8.v v25, (t3)
  1068: 570aa627                 vand.vv v20, v26, v12
  1072: d73ac2a2                 vsrl.vi v21, v12, 4
  1076: 570b8a33                 vrgather.vv v22, v24, v20
  1080: d78b9a33                 vrgather.vv v23, v25, v21
  1084: 570b7b2f                 vxor.vv v22, v23, v22
  1088: 5704642f                 vxor.vv v8, v22, v8
  1092: 578aa627                 vand.vv v20, v26, v13
  1096: d73ad2a2                 vsrl.vi v21, v13, 4
  1100: 570b8a33                 vrgather.vv v22, v24, v20
  1104: d78b9a33                 vrgather.vv v23, v25, v21
  1108: 570b7b2f                 vxor.vv v22, v23, v22
  1112: d784642f                 vxor.vv v9, v22, v9
  1116: 570aa727                 vand.vv v20, v26, v14
  1120: d73ae2a2                 vsrl.vi v21, v14, 4
  1124: 570b8a33                 vrgather.vv v22, v24, v20
  1128: d78b9a33                 vrgather.vv v23, v25, v21
  1132: 570b7b2f                 vxor.vv v22, v23, v22
  1136: 5705652f                 vxor.vv v10, v22, v10
  1140: 578aa727                 vand.vv v20, v26, v15
  1144: d73af2a2                 vsrl.vi v21, v15, 4
  1148: 570b8a33                 vrgather.vv v22, v24, v20
  1152: d78b9a33                 vrgather.vv v23, v25, v21
  1156: 570b7b2f                 vxor.vv v22, v23, v22
  1160: d785652f                 vxor.vv v11, v22, v11
  1164: 83330502                 ld t2, 0x20(a0)
  1168: b3836300                 add t2, t2, t1
  1172: 07860302                 vle8.v v12, (t2)
  1176: b3835300                 add t2, t2, t0
  1180: 87860302                 vle8.v v13, (t2)
  1184: b3835300                 add t2, t2, t0
  1188: 07870302                 vle8.v v14, (t2)
  1192: b3835300                 add t2, t2, t0
  1196: 87870302                 vle8.v v15, (t2)
  1200: 5700c02e                 vxor.vv v0, v12, v0
  1204: d780d02e                 vxor.vv v1, v13, v1
  1208: 5701e12e                 vxor.vv v2, v14, v2
  1212: d781f12e                 vxor.vv v3, v15, v3
  1216: 138e0f06                 addi t3, t6, 0x60
  1220: 070c0e02                 vle8.v v24, (t3)
  1224: 138e0f07                 addi t3, t6, 0x70
  1228: 870c0e02                 vle8.v v25, (t3)
  1232: 570aa627                 vand.vv v20, v26, v12
  1236: d73ac2a2                 vsrl.vi v21, v12, 4
  1240: 570b8a33                 vrgather.vv v22, v24, v20
  1244: d78b9a33                 vrgather.vv v23, v25, v21
  1248: 570b7b2f                 vxor.vv v22, v23, v22
  1252: 5702622f                 vxor.vv v4, v22, v4
  1256: 578aa627                 vand.vv v20, v26, v13
  1260: d73ad2a2                 vsrl.vi v21, v13, 4
  1264: 570b8a33                 vrgather.vv v22, v24, v20
  1268: d78b9a33                 vrgather.vv v23, v25, v21
  1272: 570b7b2f                 vxor.vv v22, v23, v22
  1276: d782622f                 vxor.vv v5, v22, v5
  1280: 570aa727                 vand.vv v20, v26, v14
  1284: d73ae2a2                 vsrl.vi v21, v14, 4
  1288: 570b8a33                 vrgather.vv v22, v24, v20
  1292: d78b9a33                 vrgather.vv v23, v25, v21
  1296: 570b7b2f                 vxor.vv v22, v23, v22
  1300: 5703632f                 vxor.vv v6, v22, v6
  1304: 578aa727                 vand.vv v20, v26, v15
  1308: d73af2a2                 vsrl.vi v21, v15, 4
  1312: 570b8a33                 vrgather.vv v22, v24, v20
  1316: d78b9a33                 vrgather.vv v23, v25, v21
  1320: 570b7b2f                 vxor.vv v22, v23, v22
  1324: d783632f                 vxor.vv v7, v22, v7
  1328: 138e0f08                 addi t3, t6, 0x80
  1332: 070c0e02                 vle8.v v24, (t3)
  1336: 138e0f09                 addi t3, t6, 0x90
  1340: 870c0e02                 vle8.v v25, (t3)
  1344: 570aa627                 vand.vv v20, v26, v12
  1348: d73ac2a2                 vsrl.vi v21, v12, 4
  1352: 570b8a33                 vrgather.vv v22, v24, v20
  1356: d78b9a33                 vrgather.vv v23, v25, v21
  1360: 570b7b2f                 vxor.vv v22, v23, v22
  1364: 5704642f                 vxor.vv v8, v22, v8
  1368: 578aa627                 vand.vv v20, v26, v13
  1372: d73ad2a2                 vsrl.vi v21, v13, 4
  1376: 570b8a33                 vrgather.vv v22, v24, v20
  1380: d78b9a33                 vrgather.vv v23, v25, v21
  1384: 570b7b2f                 vxor.vv v22, v23, v22
  1388: d784642f                 vxor.vv v9, v22, v9
  1392: 570aa727                 vand.vv v20, v26, v14
  1396: d73ae2a2                 vsrl.vi v21, v14, 4
  1400: 570b8a33                 vrgather.vv v22, v24, v20
  1404: d78b9a33                 vrgather.vv v23, v25, v21
  1408: 570b7b2f                 vxor.vv v22, v23, v22
  1412: 5705652f                 vxor.vv v10, v22, v10
  1416: 578aa727                 vand.vv v20, v26, v15
  1420: d73af2a2                 vsrl.vi v21, v15, 4
  1424: 570b8a33                 vrgather.vv v22, v24, v20
  1428: d78b9a33                 vrgather.vv v23, v25, v21
  1432: 570b7b2f                 vxor.vv v22, v23, v22
  1436: d785652f                 vxor.vv v11, v22, v11
  1440: 83338502                 ld t2, 0x28(a0)
  1444: b3836300                 add t2, t2, t1
  1448: 07860302                 vle8.v v12, (t2)
  1452: b3835300                 add t2, t2, t0
  1456: 87860302                 vle8.v v13, (t2)
  1460: b3835300                 add t2, t2, t0
  1464: 07870302                 vle8.v v14, (t2)
  1468: b3835300                 add t2, t2, t0
  1472: 87870302                 vle8.v v15, (t2)
  1476: 5700c02e                 vxor.vv v0, v12, v0
  1480: d780d02e                 vxor.vv v1, v13, v1
  1484: 5701e12e                 vxor.vv v2, v14, v2
  1488: d781f12e                 vxor.vv v3, v15, v3
  1492: 5702c22e                 vxor.vv v4, v12, v4
  1496: d782d22e                 vxor.vv v5, v13, v5
  1500: 5703e32e                 vxor.vv v6, v14, v6
  1504: d783f32e                 vxor.vv v7, v15, v7
  1508: 5704c42e                 vxor.vv v8, v12, v8
  1512: d784d42e                 vxor.vv v9, v13, v9
  1516: 5705e52e                 vxor.vv v10, v14, v10
  1520: d785f52e                 vxor.vv v11, v15, v11
  1524: 130ff7ff                 addi t5, a4, -1
  1528: 63060f02                 beqz t5, 0x624
  1532: 330fd300                 add t5, t1, a3
  1536: 83330500                 ld t2, 0(a0)
  1540: b383e301                 add t2, t2, t5
  1544: 07880302                 vle8.v v16, (t2)
  1548: b3835300                 add t2, t2, t0
  1552: 87880302                 vle8.v v17, (t2)
  1556: b3835300                 add t2, t2, t0
  1560: 07890302                 vle8.v v18, (t2)
  1564: b3835300                 add t2, t2, t0
  1568: 87890302                 vle8.v v19, (t2)
  1572: 83b30500                 ld t2, 0(a1)
  1576: b3836300                 add t2, t2, t1
  1580: 33005000                 ntl.all 
  1584: 27800302                 vse8.v v0, (t2)
  1588: b3835300                 add t2, t2, t0
  1592: 33005000                 ntl.all 
  1596: a7800302                 vse8.v v1, (t2)
  1600: b3835300                 add t2, t2, t0
  1604: 33005000                 ntl.all 
  1608: 27810302                 vse8.v v2, (t2)
  1612: b3835300                 add t2, t2, t0
  1616: 33005000                 ntl.all 
  1620: a7810302                 vse8.v v3, (t2)
  1624: 83b38500                 ld t2, 8(a1)
  1628: b3836300                 add t2, t2, t1
  1632: 33005000                 ntl.all 
  1636: 27820302                 vse8.v v4, (t2)
  1640: b3835300                 add t2, t2, t0
  1644: 33005000                 ntl.all 
  1648: a7820302                 vse8.v v5, (t2)
  1652: b3835300                 add t2, t2, t0
  1656: 33005000                 ntl.all 
  1660: 27830302                 vse8.v v6, (t2)
  1664: b3835300                 add t2, t2, t0
  1668: 33005000                 ntl.all 
  1672: a7830302                 vse8.v v7, (t2)
  1676: 83b30501                 ld t2, 0x10(a1)
  1680: b3836300                 add t2, t2, t1
  1684: 33005000                 ntl.all 
  1688: 27840302                 vse8.v v8, (t2)
  1692: b3835300                 add t2, t2, t0
  1696: 33005000                 ntl.all 
  1700: a7840302                 vse8.v v9, (t2)
  1704: b3835300                 add t2, t2, t0
  1708: 33005000                 ntl.all 
  1712: 27850302                 vse8.v v10, (t2)
  1716: b3835300                 add t2, t2, t0
  1720: 33005000                 ntl.all 
  1724: a7850302                 vse8.v v11, (t2)
  1728: 3303d300                 add t1, t1, a3
  1732: 1307f7ff                 addi a4, a4, -1
  1736: e31c0798                 bnez a4, 0x60
  1740: 33066640                 sub a2, a2, t1
  1744: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
  1748: 63940200                 bnez t0, 0x6dc
  1752: 6f00002b                 j 0x988
  1756: 83330500                 ld t2, 0(a0)
  1760: b3836300                 add t2, t2, t1
  1764: 07860302                 vle8.v v12, (t2)
  1768: 5700065e                 vmv.v.v v0, v12
  1772: 138f0200                 mv t5, t0
  1776: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1780: 138e0f00                 mv t3, t6
  1784: 070c0e02                 vle8.v v24, (t3)
  1788: 138e0f01                 addi t3, t6, 0x10
  1792: 870c0e02                 vle8.v v25, (t3)
  1796: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1800: 570aa627                 vand.vv v20, v26, v12
  1804: d73ac2a2                 vsrl.vi v21, v12, 4
  1808: 570b8a33                 vrgather.vv v22, v24, v20
  1812: d78b9a33                 vrgather.vv v23, v25, v21
  1816: 570b7b2f                 vxor.vv v22, v23, v22
  1820: 57020b5e                 vmv.v.v v4, v22
  1824: 138f0200                 mv t5, t0
  1828: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1832: 138e0f02                 addi t3, t6, 0x20
  1836: 070c0e02                 vle8.v v24, (t3)
  1840: 138e0f03                 addi t3, t6, 0x30
  1844: 870c0e02                 vle8.v v25, (t3)
  1848: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1852: 570aa627                 vand.vv v20, v26, v12
  1856: d73ac2a2                 vsrl.vi v21, v12, 4
  1860: 570b8a33                 vrgather.vv v22, v24, v20
  1864: d78b9a33                 vrgather.vv v23, v25, v21
  1868: 570b7b2f                 vxor.vv v22, v23, v22
  1872: 57040b5e                 vmv.v.v v8, v22
  1876: 83338500                 ld t2, 8(a0)
  1880: b3836300                 add t2, t2, t1
  1884: 07860302                 vle8.v v12, (t2)
  1888: 5700c02e                 vxor.vv v0, v12, v0
  1892: 138f0200                 mv t5, t0
  1896: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1900: 138e0f02                 addi t3, t6, 0x20
  1904: 070c0e02                 vle8.v v24, (t3)
  1908: 138e0f03                 addi t3, t6, 0x30
  1912: 870c0e02                 vle8.v v25, (t3)
  1916: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1920: 570aa627                 vand.vv v20, v26, v12
  1924: d73ac2a2                 vsrl.vi v21, v12, 4
  1928: 570b8a33                 vrgather.vv v22, v24, v20
  1932: d78b9a33                 vrgather.vv v23, v25, v21
  1936: 570b7b2f                 vxor.vv v22, v23, v22
  1940: 5702622f                 vxor.vv v4, v22, v4
  1944: 138f0200                 mv t5, t0
  1948: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1952: 138e0f04                 addi t3, t6, 0x40
  1956: 070c0e02                 vle8.v v24, (t3)
  1960: 138e0f05                 addi t3, t6, 0x50
  1964: 870c0e02                 vle8.v v25, (t3)
  1968: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1972: 570aa627                 vand.vv v20, v26, v12
  1976: d73ac2a2                 vsrl.vi v21, v12, 4
  1980: 570b8a33                 vrgather.vv v22, v24, v20
  1984: d78b9a33                 vrgather.vv v23, v25, v21
  1988: 570b7b2f                 vxor.vv v22, v23, v22
  1992: 5704642f                 vxor.vv v8, v22, v8
  1996: 83330501                 ld t2, 0x10(a0)
  2000: b3836300                 add t2, t2, t1
  2004: 07860302                 vle8.v v12, (t2)
  2008: 5700c02e                 vxor.vv v0, v12, v0
  2012: 138f0200                 mv t5, t0
  2016: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2020: 138e0f04                 addi t3, t6, 0x40
  2024: 070c0e02                 vle8.v v24, (t3)
  2028: 138e0f05                 addi t3, t6, 0x50
  2032: 870c0e02                 vle8.v v25, (t3)
  2036: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2040: 570aa627                 vand.vv v20, v26, v12
  2044: d73ac2a2                 vsrl.vi v21, v12, 4
  2048: 570b8a33                 vrgather.vv v22, v24, v20
  2052: d78b9a33                 vrgather.vv v23, v25, v21
  2056: 570b7b2f                 vxor.vv v22, v23, v22
  2060: 5702622f                 vxor.vv v4, v22, v4
  2064: 138f0200                 mv t5, t0
  2068: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2072: 138e0f06                 addi t3, t6, 0x60
  2076: 070c0e02                 vle8.v v24, (t3)
  2080: 138e0f07                 addi t3, t6, 0x70
  2084: 870c0e02                 vle8.v v25, (t3)
  2088: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2092: 570aa627                 vand.vv v20, v26, v12
  2096: d73ac2a2                 vsrl.vi v21, v12, 4
  2100: 570b8a33                 vrgather.vv v22, v24, v20
  2104: d78b9a33                 vrgather.vv v23, v25, v21
  2108: 570b7b2f                 vxor.vv v22, v23, v22
  2112: 5704642f                 vxor.vv v8, v22, v8
  2116: 83338501                 ld t2, 0x18(a0)
  2120: b3836300                 add t2, t2, t1
  2124: 07860302                 vle8.v v12, (t2)
  2128: 5700c02e                 vxor.vv v0, v12, v0
  2132: 138f0200                 mv t5, t0
  2136: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2140: 138e0f08                 addi t3, t6, 0x80
  2144: 070c0e02                 vle8.v v24, (t3)
  2148: 138e0f09                 addi t3, t6, 0x90
  2152: 870c0e02                 vle8.v v25, (t3)
  2156: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2160: 570aa627                 vand.vv v20, v26, v12
  2164: d73ac2a2                 vsrl.vi v21, v12, 4
  2168: 570b8a33                 vrgather.vv v22, v24, v20
  2172: d78b9a33                 vrgather.vv v23, v25, v21
  2176: 570b7b2f                 vxor.vv v22, v23, v22
  2180: 5702622f                 vxor.vv v4, v22, v4
  2184: 138f0200                 mv t5, t0
  2188: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2192: 138e0f00                 mv t3, t6
  2196: 070c0e02                 vle8.v v24, (t3)
  2200: 138e0f01                 addi t3, t6, 0x10
  2204: 870c0e02                 vle8.v v25, (t3)
  2208: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2212: 570aa627                 vand.vv v20, v26, v12
  2216: d73ac2a2                 vsrl.vi v21, v12, 4
  2220: 570b8a33                 vrgather.vv v22, v24, v20
  2224: d78b9a33                 vrgather.vv v23, v25, v21
  2228: 570b7b2f                 vxor.vv v22, v23, v22
  2232: 5704642f                 vxor.vv v8, v22, v8
  2236: 83330502                 ld t2, 0x20(a0)
  2240: b3836300                 add t2, t2, t1
  2244: 07860302                 vle8.v v12, (t2)
  2248: 5700c02e                 vxor.vv v0, v12, v0
  2252: 138f0200                 mv t5, t0
  2256: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2260: 138e0f06                 addi t3, t6, 0x60
  2264: 070c0e02                 vle8.v v24, (t3)
  2268: 138e0f07                 addi t3, t6, 0x70
  2272: 870c0e02                 vle8.v v25, (t3)
  2276: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2280: 570aa627                 vand.vv v20, v26, v12
  2284: d73ac2a2                 vsrl.vi v21, v12, 4
  2288: 570b8a33                 vrgather.vv v22, v24, v20
  2292: d78b9a33                 vrgather.vv v23, v25, v21
  2296: 570b7b2f                 vxor.vv v22, v23, v22
  2300: 5702622f                 vxor.vv v4, v22, v4
  2304: 138f0200                 mv t5, t0
  2308: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  2312: 138e0f08                 addi t3, t6, 0x80
  2316: 070c0e02                 vle8.v v24, (t3)
  2320: 138e0f09                 addi t3, t6, 0x90
  2324: 870c0e02                 vle8.v v25, (t3)
  2328: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  2332: 570aa627                 vand.vv v20, v26, v12
  2336: d73ac2a2                 vsrl.vi v21, v12, 4
  2340: 570b8a33                 vrgather.vv v22, v24, v20
  2344: d78b9a33                 vrgather.vv v23, v25, v21
  2348: 570b7b2f                 vxor.vv v22, v23, v22
  2352: 5704642f                 vxor.vv v8, v22, v8
  2356: 83338502                 ld t2, 0x28(a0)
  2360: b3836300                 add t2, t2, t1
  2364: 07860302                 vle8.v v12, (t2)
  2368: 5700c02e                 vxor.vv v0, v12, v0
  2372: 5702c22e                 vxor.vv v4, v12, v4
  2376: 5704c42e                 vxor.vv v8, v12, v8
  2380: 83b30500                 ld t2, 0(a1)
  2384: b3836300                 add t2, t2, t1
  2388: 33005000                 ntl.all 
  2392: 27800302                 vse8.v v0, (t2)
  2396: 83b38500                 ld t2, 8(a1)
  2400: b3836300                 add t2, t2, t1
  2404: 33005000                 ntl.all 
  2408: 27820302                 vse8.v v4, (t2)
  2412: 83b30501                 ld t2, 0x10(a1)
  2416: b3836300                 add t2, t2, t1
  2420: 33005000                 ntl.all 
  2424: 27840302                 vse8.v v8, (t2)
  2428: 33035300                 add t1, t1, t0
  2432: 33065640                 sub a2, a2, t0
  2436: 6ff0dfd4                 j 0x6d0
  2440: 67800000                 ret 
  2444: 0000                     unimp 
  2446: 0000                     unimp 
