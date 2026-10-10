; config_t3_2119 target=rvv twin=baseline topo=lrc_azure()/enc op=encode U=4
; code_size=4768 twins=[] uops_per_byte=16.46875
     0: 63140600                 bnez a2, 8
     4: 6f104029                 j 0x1298
     8: 971f0000                 auipc t6, 1
    12: 938f8f29                 addi t6, t6, 0x298
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574d0e5e                 vmv.v.x v26, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93962200                 slli a3, t0, 2
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00104e                 j 0xd10
    52: 13000000                 nop 
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
   100: 5700085e                 vmv.v.v v0, v16
   104: d780085e                 vmv.v.v v1, v17
   108: 5701095e                 vmv.v.v v2, v18
   112: d781095e                 vmv.v.v v3, v19
   116: 5702085e                 vmv.v.v v4, v16
   120: d782085e                 vmv.v.v v5, v17
   124: 5703095e                 vmv.v.v v6, v18
   128: d783095e                 vmv.v.v v7, v19
   132: 5704085e                 vmv.v.v v8, v16
   136: d784085e                 vmv.v.v v9, v17
   140: 5705095e                 vmv.v.v v10, v18
   144: d785095e                 vmv.v.v v11, v19
   148: 83338500                 ld t2, 8(a0)
   152: b3836300                 add t2, t2, t1
   156: 07880302                 vle8.v v16, (t2)
   160: b3835300                 add t2, t2, t0
   164: 87880302                 vle8.v v17, (t2)
   168: b3835300                 add t2, t2, t0
   172: 07890302                 vle8.v v18, (t2)
   176: b3835300                 add t2, t2, t0
   180: 87890302                 vle8.v v19, (t2)
   184: 5704042f                 vxor.vv v8, v16, v8
   188: d784142f                 vxor.vv v9, v17, v9
   192: 5705252f                 vxor.vv v10, v18, v10
   196: d785352f                 vxor.vv v11, v19, v11
   200: 138e0f00                 mv t3, t6
   204: 070c0e02                 vle8.v v24, (t3)
   208: 138e0f01                 addi t3, t6, 0x10
   212: 870c0e02                 vle8.v v25, (t3)
   216: 570aa827                 vand.vv v20, v26, v16
   220: d73a02a3                 vsrl.vi v21, v16, 4
   224: 570b8a33                 vrgather.vv v22, v24, v20
   228: d78b9a33                 vrgather.vv v23, v25, v21
   232: 570b7b2f                 vxor.vv v22, v23, v22
   236: 5700602f                 vxor.vv v0, v22, v0
   240: 578aa827                 vand.vv v20, v26, v17
   244: d73a12a3                 vsrl.vi v21, v17, 4
   248: 570b8a33                 vrgather.vv v22, v24, v20
   252: d78b9a33                 vrgather.vv v23, v25, v21
   256: 570b7b2f                 vxor.vv v22, v23, v22
   260: d780602f                 vxor.vv v1, v22, v1
   264: 570aa927                 vand.vv v20, v26, v18
   268: d73a22a3                 vsrl.vi v21, v18, 4
   272: 570b8a33                 vrgather.vv v22, v24, v20
   276: d78b9a33                 vrgather.vv v23, v25, v21
   280: 570b7b2f                 vxor.vv v22, v23, v22
   284: 5701612f                 vxor.vv v2, v22, v2
   288: 578aa927                 vand.vv v20, v26, v19
   292: d73a32a3                 vsrl.vi v21, v19, 4
   296: 570b8a33                 vrgather.vv v22, v24, v20
   300: d78b9a33                 vrgather.vv v23, v25, v21
   304: 570b7b2f                 vxor.vv v22, v23, v22
   308: d781612f                 vxor.vv v3, v22, v3
   312: 138e0f02                 addi t3, t6, 0x20
   316: 070c0e02                 vle8.v v24, (t3)
   320: 138e0f03                 addi t3, t6, 0x30
   324: 870c0e02                 vle8.v v25, (t3)
   328: 570aa827                 vand.vv v20, v26, v16
   332: d73a02a3                 vsrl.vi v21, v16, 4
   336: 570b8a33                 vrgather.vv v22, v24, v20
   340: d78b9a33                 vrgather.vv v23, v25, v21
   344: 570b7b2f                 vxor.vv v22, v23, v22
   348: 5702622f                 vxor.vv v4, v22, v4
   352: 578aa827                 vand.vv v20, v26, v17
   356: d73a12a3                 vsrl.vi v21, v17, 4
   360: 570b8a33                 vrgather.vv v22, v24, v20
   364: d78b9a33                 vrgather.vv v23, v25, v21
   368: 570b7b2f                 vxor.vv v22, v23, v22
   372: d782622f                 vxor.vv v5, v22, v5
   376: 570aa927                 vand.vv v20, v26, v18
   380: d73a22a3                 vsrl.vi v21, v18, 4
   384: 570b8a33                 vrgather.vv v22, v24, v20
   388: d78b9a33                 vrgather.vv v23, v25, v21
   392: 570b7b2f                 vxor.vv v22, v23, v22
   396: 5703632f                 vxor.vv v6, v22, v6
   400: 578aa927                 vand.vv v20, v26, v19
   404: d73a32a3                 vsrl.vi v21, v19, 4
   408: 570b8a33                 vrgather.vv v22, v24, v20
   412: d78b9a33                 vrgather.vv v23, v25, v21
   416: 570b7b2f                 vxor.vv v22, v23, v22
   420: d783632f                 vxor.vv v7, v22, v7
   424: 83330501                 ld t2, 0x10(a0)
   428: b3836300                 add t2, t2, t1
   432: 07880302                 vle8.v v16, (t2)
   436: b3835300                 add t2, t2, t0
   440: 87880302                 vle8.v v17, (t2)
   444: b3835300                 add t2, t2, t0
   448: 07890302                 vle8.v v18, (t2)
   452: b3835300                 add t2, t2, t0
   456: 87890302                 vle8.v v19, (t2)
   460: 5704042f                 vxor.vv v8, v16, v8
   464: d784142f                 vxor.vv v9, v17, v9
   468: 5705252f                 vxor.vv v10, v18, v10
   472: d785352f                 vxor.vv v11, v19, v11
   476: 138e0f02                 addi t3, t6, 0x20
   480: 070c0e02                 vle8.v v24, (t3)
   484: 138e0f03                 addi t3, t6, 0x30
   488: 870c0e02                 vle8.v v25, (t3)
   492: 570aa827                 vand.vv v20, v26, v16
   496: d73a02a3                 vsrl.vi v21, v16, 4
   500: 570b8a33                 vrgather.vv v22, v24, v20
   504: d78b9a33                 vrgather.vv v23, v25, v21
   508: 570b7b2f                 vxor.vv v22, v23, v22
   512: 5700602f                 vxor.vv v0, v22, v0
   516: 578aa827                 vand.vv v20, v26, v17
   520: d73a12a3                 vsrl.vi v21, v17, 4
   524: 570b8a33                 vrgather.vv v22, v24, v20
   528: d78b9a33                 vrgather.vv v23, v25, v21
   532: 570b7b2f                 vxor.vv v22, v23, v22
   536: d780602f                 vxor.vv v1, v22, v1
   540: 570aa927                 vand.vv v20, v26, v18
   544: d73a22a3                 vsrl.vi v21, v18, 4
   548: 570b8a33                 vrgather.vv v22, v24, v20
   552: d78b9a33                 vrgather.vv v23, v25, v21
   556: 570b7b2f                 vxor.vv v22, v23, v22
   560: 5701612f                 vxor.vv v2, v22, v2
   564: 578aa927                 vand.vv v20, v26, v19
   568: d73a32a3                 vsrl.vi v21, v19, 4
   572: 570b8a33                 vrgather.vv v22, v24, v20
   576: d78b9a33                 vrgather.vv v23, v25, v21
   580: 570b7b2f                 vxor.vv v22, v23, v22
   584: d781612f                 vxor.vv v3, v22, v3
   588: 138e0f04                 addi t3, t6, 0x40
   592: 070c0e02                 vle8.v v24, (t3)
   596: 138e0f05                 addi t3, t6, 0x50
   600: 870c0e02                 vle8.v v25, (t3)
   604: 570aa827                 vand.vv v20, v26, v16
   608: d73a02a3                 vsrl.vi v21, v16, 4
   612: 570b8a33                 vrgather.vv v22, v24, v20
   616: d78b9a33                 vrgather.vv v23, v25, v21
   620: 570b7b2f                 vxor.vv v22, v23, v22
   624: 5702622f                 vxor.vv v4, v22, v4
   628: 578aa827                 vand.vv v20, v26, v17
   632: d73a12a3                 vsrl.vi v21, v17, 4
   636: 570b8a33                 vrgather.vv v22, v24, v20
   640: d78b9a33                 vrgather.vv v23, v25, v21
   644: 570b7b2f                 vxor.vv v22, v23, v22
   648: d782622f                 vxor.vv v5, v22, v5
   652: 570aa927                 vand.vv v20, v26, v18
   656: d73a22a3                 vsrl.vi v21, v18, 4
   660: 570b8a33                 vrgather.vv v22, v24, v20
   664: d78b9a33                 vrgather.vv v23, v25, v21
   668: 570b7b2f                 vxor.vv v22, v23, v22
   672: 5703632f                 vxor.vv v6, v22, v6
   676: 578aa927                 vand.vv v20, v26, v19
   680: d73a32a3                 vsrl.vi v21, v19, 4
   684: 570b8a33                 vrgather.vv v22, v24, v20
   688: d78b9a33                 vrgather.vv v23, v25, v21
   692: 570b7b2f                 vxor.vv v22, v23, v22
   696: d783632f                 vxor.vv v7, v22, v7
   700: 83338501                 ld t2, 0x18(a0)
   704: b3836300                 add t2, t2, t1
   708: 07880302                 vle8.v v16, (t2)
   712: b3835300                 add t2, t2, t0
   716: 87880302                 vle8.v v17, (t2)
   720: b3835300                 add t2, t2, t0
   724: 07890302                 vle8.v v18, (t2)
   728: b3835300                 add t2, t2, t0
   732: 87890302                 vle8.v v19, (t2)
   736: 5704042f                 vxor.vv v8, v16, v8
   740: d784142f                 vxor.vv v9, v17, v9
   744: 5705252f                 vxor.vv v10, v18, v10
   748: d785352f                 vxor.vv v11, v19, v11
   752: 138e0f06                 addi t3, t6, 0x60
   756: 070c0e02                 vle8.v v24, (t3)
   760: 138e0f07                 addi t3, t6, 0x70
   764: 870c0e02                 vle8.v v25, (t3)
   768: 570aa827                 vand.vv v20, v26, v16
   772: d73a02a3                 vsrl.vi v21, v16, 4
   776: 570b8a33                 vrgather.vv v22, v24, v20
   780: d78b9a33                 vrgather.vv v23, v25, v21
   784: 570b7b2f                 vxor.vv v22, v23, v22
   788: 5700602f                 vxor.vv v0, v22, v0
   792: 578aa827                 vand.vv v20, v26, v17
   796: d73a12a3                 vsrl.vi v21, v17, 4
   800: 570b8a33                 vrgather.vv v22, v24, v20
   804: d78b9a33                 vrgather.vv v23, v25, v21
   808: 570b7b2f                 vxor.vv v22, v23, v22
   812: d780602f                 vxor.vv v1, v22, v1
   816: 570aa927                 vand.vv v20, v26, v18
   820: d73a22a3                 vsrl.vi v21, v18, 4
   824: 570b8a33                 vrgather.vv v22, v24, v20
   828: d78b9a33                 vrgather.vv v23, v25, v21
   832: 570b7b2f                 vxor.vv v22, v23, v22
   836: 5701612f                 vxor.vv v2, v22, v2
   840: 578aa927                 vand.vv v20, v26, v19
   844: d73a32a3                 vsrl.vi v21, v19, 4
   848: 570b8a33                 vrgather.vv v22, v24, v20
   852: d78b9a33                 vrgather.vv v23, v25, v21
   856: 570b7b2f                 vxor.vv v22, v23, v22
   860: d781612f                 vxor.vv v3, v22, v3
   864: 138e0f08                 addi t3, t6, 0x80
   868: 070c0e02                 vle8.v v24, (t3)
   872: 138e0f09                 addi t3, t6, 0x90
   876: 870c0e02                 vle8.v v25, (t3)
   880: 570aa827                 vand.vv v20, v26, v16
   884: d73a02a3                 vsrl.vi v21, v16, 4
   888: 570b8a33                 vrgather.vv v22, v24, v20
   892: d78b9a33                 vrgather.vv v23, v25, v21
   896: 570b7b2f                 vxor.vv v22, v23, v22
   900: 5702622f                 vxor.vv v4, v22, v4
   904: 578aa827                 vand.vv v20, v26, v17
   908: d73a12a3                 vsrl.vi v21, v17, 4
   912: 570b8a33                 vrgather.vv v22, v24, v20
   916: d78b9a33                 vrgather.vv v23, v25, v21
   920: 570b7b2f                 vxor.vv v22, v23, v22
   924: d782622f                 vxor.vv v5, v22, v5
   928: 570aa927                 vand.vv v20, v26, v18
   932: d73a22a3                 vsrl.vi v21, v18, 4
   936: 570b8a33                 vrgather.vv v22, v24, v20
   940: d78b9a33                 vrgather.vv v23, v25, v21
   944: 570b7b2f                 vxor.vv v22, v23, v22
   948: 5703632f                 vxor.vv v6, v22, v6
   952: 578aa927                 vand.vv v20, v26, v19
   956: d73a32a3                 vsrl.vi v21, v19, 4
   960: 570b8a33                 vrgather.vv v22, v24, v20
   964: d78b9a33                 vrgather.vv v23, v25, v21
   968: 570b7b2f                 vxor.vv v22, v23, v22
   972: d783632f                 vxor.vv v7, v22, v7
   976: 83330502                 ld t2, 0x20(a0)
   980: b3836300                 add t2, t2, t1
   984: 07880302                 vle8.v v16, (t2)
   988: b3835300                 add t2, t2, t0
   992: 87880302                 vle8.v v17, (t2)
   996: b3835300                 add t2, t2, t0
  1000: 07890302                 vle8.v v18, (t2)
  1004: b3835300                 add t2, t2, t0
  1008: 87890302                 vle8.v v19, (t2)
  1012: 5704042f                 vxor.vv v8, v16, v8
  1016: d784142f                 vxor.vv v9, v17, v9
  1020: 5705252f                 vxor.vv v10, v18, v10
  1024: d785352f                 vxor.vv v11, v19, v11
  1028: 138e0f04                 addi t3, t6, 0x40
  1032: 070c0e02                 vle8.v v24, (t3)
  1036: 138e0f05                 addi t3, t6, 0x50
  1040: 870c0e02                 vle8.v v25, (t3)
  1044: 570aa827                 vand.vv v20, v26, v16
  1048: d73a02a3                 vsrl.vi v21, v16, 4
  1052: 570b8a33                 vrgather.vv v22, v24, v20
  1056: d78b9a33                 vrgather.vv v23, v25, v21
  1060: 570b7b2f                 vxor.vv v22, v23, v22
  1064: 5700602f                 vxor.vv v0, v22, v0
  1068: 578aa827                 vand.vv v20, v26, v17
  1072: d73a12a3                 vsrl.vi v21, v17, 4
  1076: 570b8a33                 vrgather.vv v22, v24, v20
  1080: d78b9a33                 vrgather.vv v23, v25, v21
  1084: 570b7b2f                 vxor.vv v22, v23, v22
  1088: d780602f                 vxor.vv v1, v22, v1
  1092: 570aa927                 vand.vv v20, v26, v18
  1096: d73a22a3                 vsrl.vi v21, v18, 4
  1100: 570b8a33                 vrgather.vv v22, v24, v20
  1104: d78b9a33                 vrgather.vv v23, v25, v21
  1108: 570b7b2f                 vxor.vv v22, v23, v22
  1112: 5701612f                 vxor.vv v2, v22, v2
  1116: 578aa927                 vand.vv v20, v26, v19
  1120: d73a32a3                 vsrl.vi v21, v19, 4
  1124: 570b8a33                 vrgather.vv v22, v24, v20
  1128: d78b9a33                 vrgather.vv v23, v25, v21
  1132: 570b7b2f                 vxor.vv v22, v23, v22
  1136: d781612f                 vxor.vv v3, v22, v3
  1140: 138e0f0a                 addi t3, t6, 0xa0
  1144: 070c0e02                 vle8.v v24, (t3)
  1148: 138e0f0b                 addi t3, t6, 0xb0
  1152: 870c0e02                 vle8.v v25, (t3)
  1156: 570aa827                 vand.vv v20, v26, v16
  1160: d73a02a3                 vsrl.vi v21, v16, 4
  1164: 570b8a33                 vrgather.vv v22, v24, v20
  1168: d78b9a33                 vrgather.vv v23, v25, v21
  1172: 570b7b2f                 vxor.vv v22, v23, v22
  1176: 5702622f                 vxor.vv v4, v22, v4
  1180: 578aa827                 vand.vv v20, v26, v17
  1184: d73a12a3                 vsrl.vi v21, v17, 4
  1188: 570b8a33                 vrgather.vv v22, v24, v20
  1192: d78b9a33                 vrgather.vv v23, v25, v21
  1196: 570b7b2f                 vxor.vv v22, v23, v22
  1200: d782622f                 vxor.vv v5, v22, v5
  1204: 570aa927                 vand.vv v20, v26, v18
  1208: d73a22a3                 vsrl.vi v21, v18, 4
  1212: 570b8a33                 vrgather.vv v22, v24, v20
  1216: d78b9a33                 vrgather.vv v23, v25, v21
  1220: 570b7b2f                 vxor.vv v22, v23, v22
  1224: 5703632f                 vxor.vv v6, v22, v6
  1228: 578aa927                 vand.vv v20, v26, v19
  1232: d73a32a3                 vsrl.vi v21, v19, 4
  1236: 570b8a33                 vrgather.vv v22, v24, v20
  1240: d78b9a33                 vrgather.vv v23, v25, v21
  1244: 570b7b2f                 vxor.vv v22, v23, v22
  1248: d783632f                 vxor.vv v7, v22, v7
  1252: 83338502                 ld t2, 0x28(a0)
  1256: b3836300                 add t2, t2, t1
  1260: 07880302                 vle8.v v16, (t2)
  1264: b3835300                 add t2, t2, t0
  1268: 87880302                 vle8.v v17, (t2)
  1272: b3835300                 add t2, t2, t0
  1276: 07890302                 vle8.v v18, (t2)
  1280: b3835300                 add t2, t2, t0
  1284: 87890302                 vle8.v v19, (t2)
  1288: 5704042f                 vxor.vv v8, v16, v8
  1292: d784142f                 vxor.vv v9, v17, v9
  1296: 5705252f                 vxor.vv v10, v18, v10
  1300: d785352f                 vxor.vv v11, v19, v11
  1304: 138e0f0c                 addi t3, t6, 0xc0
  1308: 070c0e02                 vle8.v v24, (t3)
  1312: 138e0f0d                 addi t3, t6, 0xd0
  1316: 870c0e02                 vle8.v v25, (t3)
  1320: 570aa827                 vand.vv v20, v26, v16
  1324: d73a02a3                 vsrl.vi v21, v16, 4
  1328: 570b8a33                 vrgather.vv v22, v24, v20
  1332: d78b9a33                 vrgather.vv v23, v25, v21
  1336: 570b7b2f                 vxor.vv v22, v23, v22
  1340: 5700602f                 vxor.vv v0, v22, v0
  1344: 578aa827                 vand.vv v20, v26, v17
  1348: d73a12a3                 vsrl.vi v21, v17, 4
  1352: 570b8a33                 vrgather.vv v22, v24, v20
  1356: d78b9a33                 vrgather.vv v23, v25, v21
  1360: 570b7b2f                 vxor.vv v22, v23, v22
  1364: d780602f                 vxor.vv v1, v22, v1
  1368: 570aa927                 vand.vv v20, v26, v18
  1372: d73a22a3                 vsrl.vi v21, v18, 4
  1376: 570b8a33                 vrgather.vv v22, v24, v20
  1380: d78b9a33                 vrgather.vv v23, v25, v21
  1384: 570b7b2f                 vxor.vv v22, v23, v22
  1388: 5701612f                 vxor.vv v2, v22, v2
  1392: 578aa927                 vand.vv v20, v26, v19
  1396: d73a32a3                 vsrl.vi v21, v19, 4
  1400: 570b8a33                 vrgather.vv v22, v24, v20
  1404: d78b9a33                 vrgather.vv v23, v25, v21
  1408: 570b7b2f                 vxor.vv v22, v23, v22
  1412: d781612f                 vxor.vv v3, v22, v3
  1416: 138e0f0e                 addi t3, t6, 0xe0
  1420: 070c0e02                 vle8.v v24, (t3)
  1424: 138e0f0f                 addi t3, t6, 0xf0
  1428: 870c0e02                 vle8.v v25, (t3)
  1432: 570aa827                 vand.vv v20, v26, v16
  1436: d73a02a3                 vsrl.vi v21, v16, 4
  1440: 570b8a33                 vrgather.vv v22, v24, v20
  1444: d78b9a33                 vrgather.vv v23, v25, v21
  1448: 570b7b2f                 vxor.vv v22, v23, v22
  1452: 5702622f                 vxor.vv v4, v22, v4
  1456: 578aa827                 vand.vv v20, v26, v17
  1460: d73a12a3                 vsrl.vi v21, v17, 4
  1464: 570b8a33                 vrgather.vv v22, v24, v20
  1468: d78b9a33                 vrgather.vv v23, v25, v21
  1472: 570b7b2f                 vxor.vv v22, v23, v22
  1476: d782622f                 vxor.vv v5, v22, v5
  1480: 570aa927                 vand.vv v20, v26, v18
  1484: d73a22a3                 vsrl.vi v21, v18, 4
  1488: 570b8a33                 vrgather.vv v22, v24, v20
  1492: d78b9a33                 vrgather.vv v23, v25, v21
  1496: 570b7b2f                 vxor.vv v22, v23, v22
  1500: 5703632f                 vxor.vv v6, v22, v6
  1504: 578aa927                 vand.vv v20, v26, v19
  1508: d73a32a3                 vsrl.vi v21, v19, 4
  1512: 570b8a33                 vrgather.vv v22, v24, v20
  1516: d78b9a33                 vrgather.vv v23, v25, v21
  1520: 570b7b2f                 vxor.vv v22, v23, v22
  1524: d783632f                 vxor.vv v7, v22, v7
  1528: 83330503                 ld t2, 0x30(a0)
  1532: b3836300                 add t2, t2, t1
  1536: 07880302                 vle8.v v16, (t2)
  1540: b3835300                 add t2, t2, t0
  1544: 87880302                 vle8.v v17, (t2)
  1548: b3835300                 add t2, t2, t0
  1552: 07890302                 vle8.v v18, (t2)
  1556: b3835300                 add t2, t2, t0
  1560: 87890302                 vle8.v v19, (t2)
  1564: 5706085e                 vmv.v.v v12, v16
  1568: d786085e                 vmv.v.v v13, v17
  1572: 5707095e                 vmv.v.v v14, v18
  1576: d787095e                 vmv.v.v v15, v19
  1580: 138e0f08                 addi t3, t6, 0x80
  1584: 070c0e02                 vle8.v v24, (t3)
  1588: 138e0f09                 addi t3, t6, 0x90
  1592: 870c0e02                 vle8.v v25, (t3)
  1596: 570aa827                 vand.vv v20, v26, v16
  1600: d73a02a3                 vsrl.vi v21, v16, 4
  1604: 570b8a33                 vrgather.vv v22, v24, v20
  1608: d78b9a33                 vrgather.vv v23, v25, v21
  1612: 570b7b2f                 vxor.vv v22, v23, v22
  1616: 5700602f                 vxor.vv v0, v22, v0
  1620: 578aa827                 vand.vv v20, v26, v17
  1624: d73a12a3                 vsrl.vi v21, v17, 4
  1628: 570b8a33                 vrgather.vv v22, v24, v20
  1632: d78b9a33                 vrgather.vv v23, v25, v21
  1636: 570b7b2f                 vxor.vv v22, v23, v22
  1640: d780602f                 vxor.vv v1, v22, v1
  1644: 570aa927                 vand.vv v20, v26, v18
  1648: d73a22a3                 vsrl.vi v21, v18, 4
  1652: 570b8a33                 vrgather.vv v22, v24, v20
  1656: d78b9a33                 vrgather.vv v23, v25, v21
  1660: 570b7b2f                 vxor.vv v22, v23, v22
  1664: 5701612f                 vxor.vv v2, v22, v2
  1668: 578aa927                 vand.vv v20, v26, v19
  1672: d73a32a3                 vsrl.vi v21, v19, 4
  1676: 570b8a33                 vrgather.vv v22, v24, v20
  1680: d78b9a33                 vrgather.vv v23, v25, v21
  1684: 570b7b2f                 vxor.vv v22, v23, v22
  1688: d781612f                 vxor.vv v3, v22, v3
  1692: 138e0f10                 addi t3, t6, 0x100
  1696: 070c0e02                 vle8.v v24, (t3)
  1700: 138e0f11                 addi t3, t6, 0x110
  1704: 870c0e02                 vle8.v v25, (t3)
  1708: 570aa827                 vand.vv v20, v26, v16
  1712: d73a02a3                 vsrl.vi v21, v16, 4
  1716: 570b8a33                 vrgather.vv v22, v24, v20
  1720: d78b9a33                 vrgather.vv v23, v25, v21
  1724: 570b7b2f                 vxor.vv v22, v23, v22
  1728: 5702622f                 vxor.vv v4, v22, v4
  1732: 578aa827                 vand.vv v20, v26, v17
  1736: d73a12a3                 vsrl.vi v21, v17, 4
  1740: 570b8a33                 vrgather.vv v22, v24, v20
  1744: d78b9a33                 vrgather.vv v23, v25, v21
  1748: 570b7b2f                 vxor.vv v22, v23, v22
  1752: d782622f                 vxor.vv v5, v22, v5
  1756: 570aa927                 vand.vv v20, v26, v18
  1760: d73a22a3                 vsrl.vi v21, v18, 4
  1764: 570b8a33                 vrgather.vv v22, v24, v20
  1768: d78b9a33                 vrgather.vv v23, v25, v21
  1772: 570b7b2f                 vxor.vv v22, v23, v22
  1776: 5703632f                 vxor.vv v6, v22, v6
  1780: 578aa927                 vand.vv v20, v26, v19
  1784: d73a32a3                 vsrl.vi v21, v19, 4
  1788: 570b8a33                 vrgather.vv v22, v24, v20
  1792: d78b9a33                 vrgather.vv v23, v25, v21
  1796: 570b7b2f                 vxor.vv v22, v23, v22
  1800: d783632f                 vxor.vv v7, v22, v7
  1804: 83338503                 ld t2, 0x38(a0)
  1808: b3836300                 add t2, t2, t1
  1812: 07880302                 vle8.v v16, (t2)
  1816: b3835300                 add t2, t2, t0
  1820: 87880302                 vle8.v v17, (t2)
  1824: b3835300                 add t2, t2, t0
  1828: 07890302                 vle8.v v18, (t2)
  1832: b3835300                 add t2, t2, t0
  1836: 87890302                 vle8.v v19, (t2)
  1840: 5706062f                 vxor.vv v12, v16, v12
  1844: d786162f                 vxor.vv v13, v17, v13
  1848: 5707272f                 vxor.vv v14, v18, v14
  1852: d787372f                 vxor.vv v15, v19, v15
  1856: 138e0f12                 addi t3, t6, 0x120
  1860: 070c0e02                 vle8.v v24, (t3)
  1864: 138e0f13                 addi t3, t6, 0x130
  1868: 870c0e02                 vle8.v v25, (t3)
  1872: 570aa827                 vand.vv v20, v26, v16
  1876: d73a02a3                 vsrl.vi v21, v16, 4
  1880: 570b8a33                 vrgather.vv v22, v24, v20
  1884: d78b9a33                 vrgather.vv v23, v25, v21
  1888: 570b7b2f                 vxor.vv v22, v23, v22
  1892: 5700602f                 vxor.vv v0, v22, v0
  1896: 578aa827                 vand.vv v20, v26, v17
  1900: d73a12a3                 vsrl.vi v21, v17, 4
  1904: 570b8a33                 vrgather.vv v22, v24, v20
  1908: d78b9a33                 vrgather.vv v23, v25, v21
  1912: 570b7b2f                 vxor.vv v22, v23, v22
  1916: d780602f                 vxor.vv v1, v22, v1
  1920: 570aa927                 vand.vv v20, v26, v18
  1924: d73a22a3                 vsrl.vi v21, v18, 4
  1928: 570b8a33                 vrgather.vv v22, v24, v20
  1932: d78b9a33                 vrgather.vv v23, v25, v21
  1936: 570b7b2f                 vxor.vv v22, v23, v22
  1940: 5701612f                 vxor.vv v2, v22, v2
  1944: 578aa927                 vand.vv v20, v26, v19
  1948: d73a32a3                 vsrl.vi v21, v19, 4
  1952: 570b8a33                 vrgather.vv v22, v24, v20
  1956: d78b9a33                 vrgather.vv v23, v25, v21
  1960: 570b7b2f                 vxor.vv v22, v23, v22
  1964: d781612f                 vxor.vv v3, v22, v3
  1968: 138e0f14                 addi t3, t6, 0x140
  1972: 070c0e02                 vle8.v v24, (t3)
  1976: 138e0f15                 addi t3, t6, 0x150
  1980: 870c0e02                 vle8.v v25, (t3)
  1984: 570aa827                 vand.vv v20, v26, v16
  1988: d73a02a3                 vsrl.vi v21, v16, 4
  1992: 570b8a33                 vrgather.vv v22, v24, v20
  1996: d78b9a33                 vrgather.vv v23, v25, v21
  2000: 570b7b2f                 vxor.vv v22, v23, v22
  2004: 5702622f                 vxor.vv v4, v22, v4
  2008: 578aa827                 vand.vv v20, v26, v17
  2012: d73a12a3                 vsrl.vi v21, v17, 4
  2016: 570b8a33                 vrgather.vv v22, v24, v20
  2020: d78b9a33                 vrgather.vv v23, v25, v21
  2024: 570b7b2f                 vxor.vv v22, v23, v22
  2028: d782622f                 vxor.vv v5, v22, v5
  2032: 570aa927                 vand.vv v20, v26, v18
  2036: d73a22a3                 vsrl.vi v21, v18, 4
  2040: 570b8a33                 vrgather.vv v22, v24, v20
  2044: d78b9a33                 vrgather.vv v23, v25, v21
  2048: 570b7b2f                 vxor.vv v22, v23, v22
  2052: 5703632f                 vxor.vv v6, v22, v6
  2056: 578aa927                 vand.vv v20, v26, v19
  2060: d73a32a3                 vsrl.vi v21, v19, 4
  2064: 570b8a33                 vrgather.vv v22, v24, v20
  2068: d78b9a33                 vrgather.vv v23, v25, v21
  2072: 570b7b2f                 vxor.vv v22, v23, v22
  2076: d783632f                 vxor.vv v7, v22, v7
  2080: 83330504                 ld t2, 0x40(a0)
  2084: b3836300                 add t2, t2, t1
  2088: 07880302                 vle8.v v16, (t2)
  2092: b3835300                 add t2, t2, t0
  2096: 87880302                 vle8.v v17, (t2)
  2100: b3835300                 add t2, t2, t0
  2104: 07890302                 vle8.v v18, (t2)
  2108: b3835300                 add t2, t2, t0
  2112: 87890302                 vle8.v v19, (t2)
  2116: 5706062f                 vxor.vv v12, v16, v12
  2120: d786162f                 vxor.vv v13, v17, v13
  2124: 5707272f                 vxor.vv v14, v18, v14
  2128: d787372f                 vxor.vv v15, v19, v15
  2132: 138e0f0a                 addi t3, t6, 0xa0
  2136: 070c0e02                 vle8.v v24, (t3)
  2140: 138e0f0b                 addi t3, t6, 0xb0
  2144: 870c0e02                 vle8.v v25, (t3)
  2148: 570aa827                 vand.vv v20, v26, v16
  2152: d73a02a3                 vsrl.vi v21, v16, 4
  2156: 570b8a33                 vrgather.vv v22, v24, v20
  2160: d78b9a33                 vrgather.vv v23, v25, v21
  2164: 570b7b2f                 vxor.vv v22, v23, v22
  2168: 5700602f                 vxor.vv v0, v22, v0
  2172: 578aa827                 vand.vv v20, v26, v17
  2176: d73a12a3                 vsrl.vi v21, v17, 4
  2180: 570b8a33                 vrgather.vv v22, v24, v20
  2184: d78b9a33                 vrgather.vv v23, v25, v21
  2188: 570b7b2f                 vxor.vv v22, v23, v22
  2192: d780602f                 vxor.vv v1, v22, v1
  2196: 570aa927                 vand.vv v20, v26, v18
  2200: d73a22a3                 vsrl.vi v21, v18, 4
  2204: 570b8a33                 vrgather.vv v22, v24, v20
  2208: d78b9a33                 vrgather.vv v23, v25, v21
  2212: 570b7b2f                 vxor.vv v22, v23, v22
  2216: 5701612f                 vxor.vv v2, v22, v2
  2220: 578aa927                 vand.vv v20, v26, v19
  2224: d73a32a3                 vsrl.vi v21, v19, 4
  2228: 570b8a33                 vrgather.vv v22, v24, v20
  2232: d78b9a33                 vrgather.vv v23, v25, v21
  2236: 570b7b2f                 vxor.vv v22, v23, v22
  2240: d781612f                 vxor.vv v3, v22, v3
  2244: 138e0f16                 addi t3, t6, 0x160
  2248: 070c0e02                 vle8.v v24, (t3)
  2252: 138e0f17                 addi t3, t6, 0x170
  2256: 870c0e02                 vle8.v v25, (t3)
  2260: 570aa827                 vand.vv v20, v26, v16
  2264: d73a02a3                 vsrl.vi v21, v16, 4
  2268: 570b8a33                 vrgather.vv v22, v24, v20
  2272: d78b9a33                 vrgather.vv v23, v25, v21
  2276: 570b7b2f                 vxor.vv v22, v23, v22
  2280: 5702622f                 vxor.vv v4, v22, v4
  2284: 578aa827                 vand.vv v20, v26, v17
  2288: d73a12a3                 vsrl.vi v21, v17, 4
  2292: 570b8a33                 vrgather.vv v22, v24, v20
  2296: d78b9a33                 vrgather.vv v23, v25, v21
  2300: 570b7b2f                 vxor.vv v22, v23, v22
  2304: d782622f                 vxor.vv v5, v22, v5
  2308: 570aa927                 vand.vv v20, v26, v18
  2312: d73a22a3                 vsrl.vi v21, v18, 4
  2316: 570b8a33                 vrgather.vv v22, v24, v20
  2320: d78b9a33                 vrgather.vv v23, v25, v21
  2324: 570b7b2f                 vxor.vv v22, v23, v22
  2328: 5703632f                 vxor.vv v6, v22, v6
  2332: 578aa927                 vand.vv v20, v26, v19
  2336: d73a32a3                 vsrl.vi v21, v19, 4
  2340: 570b8a33                 vrgather.vv v22, v24, v20
  2344: d78b9a33                 vrgather.vv v23, v25, v21
  2348: 570b7b2f                 vxor.vv v22, v23, v22
  2352: d783632f                 vxor.vv v7, v22, v7
  2356: 83338504                 ld t2, 0x48(a0)
  2360: b3836300                 add t2, t2, t1
  2364: 07880302                 vle8.v v16, (t2)
  2368: b3835300                 add t2, t2, t0
  2372: 87880302                 vle8.v v17, (t2)
  2376: b3835300                 add t2, t2, t0
  2380: 07890302                 vle8.v v18, (t2)
  2384: b3835300                 add t2, t2, t0
  2388: 87890302                 vle8.v v19, (t2)
  2392: 5706062f                 vxor.vv v12, v16, v12
  2396: d786162f                 vxor.vv v13, v17, v13
  2400: 5707272f                 vxor.vv v14, v18, v14
  2404: d787372f                 vxor.vv v15, v19, v15
  2408: 138e0f18                 addi t3, t6, 0x180
  2412: 070c0e02                 vle8.v v24, (t3)
  2416: 138e0f19                 addi t3, t6, 0x190
  2420: 870c0e02                 vle8.v v25, (t3)
  2424: 570aa827                 vand.vv v20, v26, v16
  2428: d73a02a3                 vsrl.vi v21, v16, 4
  2432: 570b8a33                 vrgather.vv v22, v24, v20
  2436: d78b9a33                 vrgather.vv v23, v25, v21
  2440: 570b7b2f                 vxor.vv v22, v23, v22
  2444: 5700602f                 vxor.vv v0, v22, v0
  2448: 578aa827                 vand.vv v20, v26, v17
  2452: d73a12a3                 vsrl.vi v21, v17, 4
  2456: 570b8a33                 vrgather.vv v22, v24, v20
  2460: d78b9a33                 vrgather.vv v23, v25, v21
  2464: 570b7b2f                 vxor.vv v22, v23, v22
  2468: d780602f                 vxor.vv v1, v22, v1
  2472: 570aa927                 vand.vv v20, v26, v18
  2476: d73a22a3                 vsrl.vi v21, v18, 4
  2480: 570b8a33                 vrgather.vv v22, v24, v20
  2484: d78b9a33                 vrgather.vv v23, v25, v21
  2488: 570b7b2f                 vxor.vv v22, v23, v22
  2492: 5701612f                 vxor.vv v2, v22, v2
  2496: 578aa927                 vand.vv v20, v26, v19
  2500: d73a32a3                 vsrl.vi v21, v19, 4
  2504: 570b8a33                 vrgather.vv v22, v24, v20
  2508: d78b9a33                 vrgather.vv v23, v25, v21
  2512: 570b7b2f                 vxor.vv v22, v23, v22
  2516: d781612f                 vxor.vv v3, v22, v3
  2520: 138e0f1a                 addi t3, t6, 0x1a0
  2524: 070c0e02                 vle8.v v24, (t3)
  2528: 138e0f1b                 addi t3, t6, 0x1b0
  2532: 870c0e02                 vle8.v v25, (t3)
  2536: 570aa827                 vand.vv v20, v26, v16
  2540: d73a02a3                 vsrl.vi v21, v16, 4
  2544: 570b8a33                 vrgather.vv v22, v24, v20
  2548: d78b9a33                 vrgather.vv v23, v25, v21
  2552: 570b7b2f                 vxor.vv v22, v23, v22
  2556: 5702622f                 vxor.vv v4, v22, v4
  2560: 578aa827                 vand.vv v20, v26, v17
  2564: d73a12a3                 vsrl.vi v21, v17, 4
  2568: 570b8a33                 vrgather.vv v22, v24, v20
  2572: d78b9a33                 vrgather.vv v23, v25, v21
  2576: 570b7b2f                 vxor.vv v22, v23, v22
  2580: d782622f                 vxor.vv v5, v22, v5
  2584: 570aa927                 vand.vv v20, v26, v18
  2588: d73a22a3                 vsrl.vi v21, v18, 4
  2592: 570b8a33                 vrgather.vv v22, v24, v20
  2596: d78b9a33                 vrgather.vv v23, v25, v21
  2600: 570b7b2f                 vxor.vv v22, v23, v22
  2604: 5703632f                 vxor.vv v6, v22, v6
  2608: 578aa927                 vand.vv v20, v26, v19
  2612: d73a32a3                 vsrl.vi v21, v19, 4
  2616: 570b8a33                 vrgather.vv v22, v24, v20
  2620: d78b9a33                 vrgather.vv v23, v25, v21
  2624: 570b7b2f                 vxor.vv v22, v23, v22
  2628: d783632f                 vxor.vv v7, v22, v7
  2632: 83330505                 ld t2, 0x50(a0)
  2636: b3836300                 add t2, t2, t1
  2640: 07880302                 vle8.v v16, (t2)
  2644: b3835300                 add t2, t2, t0
  2648: 87880302                 vle8.v v17, (t2)
  2652: b3835300                 add t2, t2, t0
  2656: 07890302                 vle8.v v18, (t2)
  2660: b3835300                 add t2, t2, t0
  2664: 87890302                 vle8.v v19, (t2)
  2668: 5706062f                 vxor.vv v12, v16, v12
  2672: d786162f                 vxor.vv v13, v17, v13
  2676: 5707272f                 vxor.vv v14, v18, v14
  2680: d787372f                 vxor.vv v15, v19, v15
  2684: 138e0f0e                 addi t3, t6, 0xe0
  2688: 070c0e02                 vle8.v v24, (t3)
  2692: 138e0f0f                 addi t3, t6, 0xf0
  2696: 870c0e02                 vle8.v v25, (t3)
  2700: 570aa827                 vand.vv v20, v26, v16
  2704: d73a02a3                 vsrl.vi v21, v16, 4
  2708: 570b8a33                 vrgather.vv v22, v24, v20
  2712: d78b9a33                 vrgather.vv v23, v25, v21
  2716: 570b7b2f                 vxor.vv v22, v23, v22
  2720: 5700602f                 vxor.vv v0, v22, v0
  2724: 578aa827                 vand.vv v20, v26, v17
  2728: d73a12a3                 vsrl.vi v21, v17, 4
  2732: 570b8a33                 vrgather.vv v22, v24, v20
  2736: d78b9a33                 vrgather.vv v23, v25, v21
  2740: 570b7b2f                 vxor.vv v22, v23, v22
  2744: d780602f                 vxor.vv v1, v22, v1
  2748: 570aa927                 vand.vv v20, v26, v18
  2752: d73a22a3                 vsrl.vi v21, v18, 4
  2756: 570b8a33                 vrgather.vv v22, v24, v20
  2760: d78b9a33                 vrgather.vv v23, v25, v21
  2764: 570b7b2f                 vxor.vv v22, v23, v22
  2768: 5701612f                 vxor.vv v2, v22, v2
  2772: 578aa927                 vand.vv v20, v26, v19
  2776: d73a32a3                 vsrl.vi v21, v19, 4
  2780: 570b8a33                 vrgather.vv v22, v24, v20
  2784: d78b9a33                 vrgather.vv v23, v25, v21
  2788: 570b7b2f                 vxor.vv v22, v23, v22
  2792: d781612f                 vxor.vv v3, v22, v3
  2796: 138e0f1c                 addi t3, t6, 0x1c0
  2800: 070c0e02                 vle8.v v24, (t3)
  2804: 138e0f1d                 addi t3, t6, 0x1d0
  2808: 870c0e02                 vle8.v v25, (t3)
  2812: 570aa827                 vand.vv v20, v26, v16
  2816: d73a02a3                 vsrl.vi v21, v16, 4
  2820: 570b8a33                 vrgather.vv v22, v24, v20
  2824: d78b9a33                 vrgather.vv v23, v25, v21
  2828: 570b7b2f                 vxor.vv v22, v23, v22
  2832: 5702622f                 vxor.vv v4, v22, v4
  2836: 578aa827                 vand.vv v20, v26, v17
  2840: d73a12a3                 vsrl.vi v21, v17, 4
  2844: 570b8a33                 vrgather.vv v22, v24, v20
  2848: d78b9a33                 vrgather.vv v23, v25, v21
  2852: 570b7b2f                 vxor.vv v22, v23, v22
  2856: d782622f                 vxor.vv v5, v22, v5
  2860: 570aa927                 vand.vv v20, v26, v18
  2864: d73a22a3                 vsrl.vi v21, v18, 4
  2868: 570b8a33                 vrgather.vv v22, v24, v20
  2872: d78b9a33                 vrgather.vv v23, v25, v21
  2876: 570b7b2f                 vxor.vv v22, v23, v22
  2880: 5703632f                 vxor.vv v6, v22, v6
  2884: 578aa927                 vand.vv v20, v26, v19
  2888: d73a32a3                 vsrl.vi v21, v19, 4
  2892: 570b8a33                 vrgather.vv v22, v24, v20
  2896: d78b9a33                 vrgather.vv v23, v25, v21
  2900: 570b7b2f                 vxor.vv v22, v23, v22
  2904: d783632f                 vxor.vv v7, v22, v7
  2908: 83338505                 ld t2, 0x58(a0)
  2912: b3836300                 add t2, t2, t1
  2916: 07880302                 vle8.v v16, (t2)
  2920: b3835300                 add t2, t2, t0
  2924: 87880302                 vle8.v v17, (t2)
  2928: b3835300                 add t2, t2, t0
  2932: 07890302                 vle8.v v18, (t2)
  2936: b3835300                 add t2, t2, t0
  2940: 87890302                 vle8.v v19, (t2)
  2944: 5706062f                 vxor.vv v12, v16, v12
  2948: d786162f                 vxor.vv v13, v17, v13
  2952: 5707272f                 vxor.vv v14, v18, v14
  2956: d787372f                 vxor.vv v15, v19, v15
  2960: 138e0f1e                 addi t3, t6, 0x1e0
  2964: 070c0e02                 vle8.v v24, (t3)
  2968: 138e0f1f                 addi t3, t6, 0x1f0
  2972: 870c0e02                 vle8.v v25, (t3)
  2976: 570aa827                 vand.vv v20, v26, v16
  2980: d73a02a3                 vsrl.vi v21, v16, 4
  2984: 570b8a33                 vrgather.vv v22, v24, v20
  2988: d78b9a33                 vrgather.vv v23, v25, v21
  2992: 570b7b2f                 vxor.vv v22, v23, v22
  2996: 5700602f                 vxor.vv v0, v22, v0
  3000: 578aa827                 vand.vv v20, v26, v17
  3004: d73a12a3                 vsrl.vi v21, v17, 4
  3008: 570b8a33                 vrgather.vv v22, v24, v20
  3012: d78b9a33                 vrgather.vv v23, v25, v21
  3016: 570b7b2f                 vxor.vv v22, v23, v22
  3020: d780602f                 vxor.vv v1, v22, v1
  3024: 570aa927                 vand.vv v20, v26, v18
  3028: d73a22a3                 vsrl.vi v21, v18, 4
  3032: 570b8a33                 vrgather.vv v22, v24, v20
  3036: d78b9a33                 vrgather.vv v23, v25, v21
  3040: 570b7b2f                 vxor.vv v22, v23, v22
  3044: 5701612f                 vxor.vv v2, v22, v2
  3048: 578aa927                 vand.vv v20, v26, v19
  3052: d73a32a3                 vsrl.vi v21, v19, 4
  3056: 570b8a33                 vrgather.vv v22, v24, v20
  3060: d78b9a33                 vrgather.vv v23, v25, v21
  3064: 570b7b2f                 vxor.vv v22, v23, v22
  3068: d781612f                 vxor.vv v3, v22, v3
  3072: 138e0f20                 addi t3, t6, 0x200
  3076: 070c0e02                 vle8.v v24, (t3)
  3080: 138e0f21                 addi t3, t6, 0x210
  3084: 870c0e02                 vle8.v v25, (t3)
  3088: 570aa827                 vand.vv v20, v26, v16
  3092: d73a02a3                 vsrl.vi v21, v16, 4
  3096: 570b8a33                 vrgather.vv v22, v24, v20
  3100: d78b9a33                 vrgather.vv v23, v25, v21
  3104: 570b7b2f                 vxor.vv v22, v23, v22
  3108: 5702622f                 vxor.vv v4, v22, v4
  3112: 578aa827                 vand.vv v20, v26, v17
  3116: d73a12a3                 vsrl.vi v21, v17, 4
  3120: 570b8a33                 vrgather.vv v22, v24, v20
  3124: d78b9a33                 vrgather.vv v23, v25, v21
  3128: 570b7b2f                 vxor.vv v22, v23, v22
  3132: d782622f                 vxor.vv v5, v22, v5
  3136: 570aa927                 vand.vv v20, v26, v18
  3140: d73a22a3                 vsrl.vi v21, v18, 4
  3144: 570b8a33                 vrgather.vv v22, v24, v20
  3148: d78b9a33                 vrgather.vv v23, v25, v21
  3152: 570b7b2f                 vxor.vv v22, v23, v22
  3156: 5703632f                 vxor.vv v6, v22, v6
  3160: 578aa927                 vand.vv v20, v26, v19
  3164: d73a32a3                 vsrl.vi v21, v19, 4
  3168: 570b8a33                 vrgather.vv v22, v24, v20
  3172: d78b9a33                 vrgather.vv v23, v25, v21
  3176: 570b7b2f                 vxor.vv v22, v23, v22
  3180: d783632f                 vxor.vv v7, v22, v7
  3184: 83b30500                 ld t2, 0(a1)
  3188: b3836300                 add t2, t2, t1
  3192: 27800302                 vse8.v v0, (t2)
  3196: b3835300                 add t2, t2, t0
  3200: a7800302                 vse8.v v1, (t2)
  3204: b3835300                 add t2, t2, t0
  3208: 27810302                 vse8.v v2, (t2)
  3212: b3835300                 add t2, t2, t0
  3216: a7810302                 vse8.v v3, (t2)
  3220: 83b38500                 ld t2, 8(a1)
  3224: b3836300                 add t2, t2, t1
  3228: 27820302                 vse8.v v4, (t2)
  3232: b3835300                 add t2, t2, t0
  3236: a7820302                 vse8.v v5, (t2)
  3240: b3835300                 add t2, t2, t0
  3244: 27830302                 vse8.v v6, (t2)
  3248: b3835300                 add t2, t2, t0
  3252: a7830302                 vse8.v v7, (t2)
  3256: 83b30501                 ld t2, 0x10(a1)
  3260: b3836300                 add t2, t2, t1
  3264: 27840302                 vse8.v v8, (t2)
  3268: b3835300                 add t2, t2, t0
  3272: a7840302                 vse8.v v9, (t2)
  3276: b3835300                 add t2, t2, t0
  3280: 27850302                 vse8.v v10, (t2)
  3284: b3835300                 add t2, t2, t0
  3288: a7850302                 vse8.v v11, (t2)
  3292: 83b38501                 ld t2, 0x18(a1)
  3296: b3836300                 add t2, t2, t1
  3300: 27860302                 vse8.v v12, (t2)
  3304: b3835300                 add t2, t2, t0
  3308: a7860302                 vse8.v v13, (t2)
  3312: b3835300                 add t2, t2, t0
  3316: 27870302                 vse8.v v14, (t2)
  3320: b3835300                 add t2, t2, t0
  3324: a7870302                 vse8.v v15, (t2)
  3328: 3303d300                 add t1, t1, a3
  3332: 1307f7ff                 addi a4, a4, -1
  3336: 631c07b2                 bnez a4, 0x40
  3340: 33066640                 sub a2, a2, t1
  3344: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
  3348: 63940200                 bnez t0, 0xd1c
  3352: 6f000058                 j 0x1298
  3356: 83330500                 ld t2, 0(a0)
  3360: b3836300                 add t2, t2, t1
  3364: 07880302                 vle8.v v16, (t2)
  3368: 5700085e                 vmv.v.v v0, v16
  3372: 5702085e                 vmv.v.v v4, v16
  3376: 5704085e                 vmv.v.v v8, v16
  3380: 83338500                 ld t2, 8(a0)
  3384: b3836300                 add t2, t2, t1
  3388: 07880302                 vle8.v v16, (t2)
  3392: 5704042f                 vxor.vv v8, v16, v8
  3396: 138f0200                 mv t5, t0
  3400: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3404: 138e0f00                 mv t3, t6
  3408: 070c0e02                 vle8.v v24, (t3)
  3412: 138e0f01                 addi t3, t6, 0x10
  3416: 870c0e02                 vle8.v v25, (t3)
  3420: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3424: 570aa827                 vand.vv v20, v26, v16
  3428: d73a02a3                 vsrl.vi v21, v16, 4
  3432: 570b8a33                 vrgather.vv v22, v24, v20
  3436: d78b9a33                 vrgather.vv v23, v25, v21
  3440: 570b7b2f                 vxor.vv v22, v23, v22
  3444: 5700602f                 vxor.vv v0, v22, v0
  3448: 138f0200                 mv t5, t0
  3452: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3456: 138e0f02                 addi t3, t6, 0x20
  3460: 070c0e02                 vle8.v v24, (t3)
  3464: 138e0f03                 addi t3, t6, 0x30
  3468: 870c0e02                 vle8.v v25, (t3)
  3472: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3476: 570aa827                 vand.vv v20, v26, v16
  3480: d73a02a3                 vsrl.vi v21, v16, 4
  3484: 570b8a33                 vrgather.vv v22, v24, v20
  3488: d78b9a33                 vrgather.vv v23, v25, v21
  3492: 570b7b2f                 vxor.vv v22, v23, v22
  3496: 5702622f                 vxor.vv v4, v22, v4
  3500: 83330501                 ld t2, 0x10(a0)
  3504: b3836300                 add t2, t2, t1
  3508: 07880302                 vle8.v v16, (t2)
  3512: 5704042f                 vxor.vv v8, v16, v8
  3516: 138f0200                 mv t5, t0
  3520: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3524: 138e0f02                 addi t3, t6, 0x20
  3528: 070c0e02                 vle8.v v24, (t3)
  3532: 138e0f03                 addi t3, t6, 0x30
  3536: 870c0e02                 vle8.v v25, (t3)
  3540: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3544: 570aa827                 vand.vv v20, v26, v16
  3548: d73a02a3                 vsrl.vi v21, v16, 4
  3552: 570b8a33                 vrgather.vv v22, v24, v20
  3556: d78b9a33                 vrgather.vv v23, v25, v21
  3560: 570b7b2f                 vxor.vv v22, v23, v22
  3564: 5700602f                 vxor.vv v0, v22, v0
  3568: 138f0200                 mv t5, t0
  3572: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3576: 138e0f04                 addi t3, t6, 0x40
  3580: 070c0e02                 vle8.v v24, (t3)
  3584: 138e0f05                 addi t3, t6, 0x50
  3588: 870c0e02                 vle8.v v25, (t3)
  3592: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3596: 570aa827                 vand.vv v20, v26, v16
  3600: d73a02a3                 vsrl.vi v21, v16, 4
  3604: 570b8a33                 vrgather.vv v22, v24, v20
  3608: d78b9a33                 vrgather.vv v23, v25, v21
  3612: 570b7b2f                 vxor.vv v22, v23, v22
  3616: 5702622f                 vxor.vv v4, v22, v4
  3620: 83338501                 ld t2, 0x18(a0)
  3624: b3836300                 add t2, t2, t1
  3628: 07880302                 vle8.v v16, (t2)
  3632: 5704042f                 vxor.vv v8, v16, v8
  3636: 138f0200                 mv t5, t0
  3640: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3644: 138e0f06                 addi t3, t6, 0x60
  3648: 070c0e02                 vle8.v v24, (t3)
  3652: 138e0f07                 addi t3, t6, 0x70
  3656: 870c0e02                 vle8.v v25, (t3)
  3660: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3664: 570aa827                 vand.vv v20, v26, v16
  3668: d73a02a3                 vsrl.vi v21, v16, 4
  3672: 570b8a33                 vrgather.vv v22, v24, v20
  3676: d78b9a33                 vrgather.vv v23, v25, v21
  3680: 570b7b2f                 vxor.vv v22, v23, v22
  3684: 5700602f                 vxor.vv v0, v22, v0
  3688: 138f0200                 mv t5, t0
  3692: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3696: 138e0f08                 addi t3, t6, 0x80
  3700: 070c0e02                 vle8.v v24, (t3)
  3704: 138e0f09                 addi t3, t6, 0x90
  3708: 870c0e02                 vle8.v v25, (t3)
  3712: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3716: 570aa827                 vand.vv v20, v26, v16
  3720: d73a02a3                 vsrl.vi v21, v16, 4
  3724: 570b8a33                 vrgather.vv v22, v24, v20
  3728: d78b9a33                 vrgather.vv v23, v25, v21
  3732: 570b7b2f                 vxor.vv v22, v23, v22
  3736: 5702622f                 vxor.vv v4, v22, v4
  3740: 83330502                 ld t2, 0x20(a0)
  3744: b3836300                 add t2, t2, t1
  3748: 07880302                 vle8.v v16, (t2)
  3752: 5704042f                 vxor.vv v8, v16, v8
  3756: 138f0200                 mv t5, t0
  3760: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3764: 138e0f04                 addi t3, t6, 0x40
  3768: 070c0e02                 vle8.v v24, (t3)
  3772: 138e0f05                 addi t3, t6, 0x50
  3776: 870c0e02                 vle8.v v25, (t3)
  3780: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3784: 570aa827                 vand.vv v20, v26, v16
  3788: d73a02a3                 vsrl.vi v21, v16, 4
  3792: 570b8a33                 vrgather.vv v22, v24, v20
  3796: d78b9a33                 vrgather.vv v23, v25, v21
  3800: 570b7b2f                 vxor.vv v22, v23, v22
  3804: 5700602f                 vxor.vv v0, v22, v0
  3808: 138f0200                 mv t5, t0
  3812: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3816: 138e0f0a                 addi t3, t6, 0xa0
  3820: 070c0e02                 vle8.v v24, (t3)
  3824: 138e0f0b                 addi t3, t6, 0xb0
  3828: 870c0e02                 vle8.v v25, (t3)
  3832: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3836: 570aa827                 vand.vv v20, v26, v16
  3840: d73a02a3                 vsrl.vi v21, v16, 4
  3844: 570b8a33                 vrgather.vv v22, v24, v20
  3848: d78b9a33                 vrgather.vv v23, v25, v21
  3852: 570b7b2f                 vxor.vv v22, v23, v22
  3856: 5702622f                 vxor.vv v4, v22, v4
  3860: 83338502                 ld t2, 0x28(a0)
  3864: b3836300                 add t2, t2, t1
  3868: 07880302                 vle8.v v16, (t2)
  3872: 5704042f                 vxor.vv v8, v16, v8
  3876: 138f0200                 mv t5, t0
  3880: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3884: 138e0f0c                 addi t3, t6, 0xc0
  3888: 070c0e02                 vle8.v v24, (t3)
  3892: 138e0f0d                 addi t3, t6, 0xd0
  3896: 870c0e02                 vle8.v v25, (t3)
  3900: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3904: 570aa827                 vand.vv v20, v26, v16
  3908: d73a02a3                 vsrl.vi v21, v16, 4
  3912: 570b8a33                 vrgather.vv v22, v24, v20
  3916: d78b9a33                 vrgather.vv v23, v25, v21
  3920: 570b7b2f                 vxor.vv v22, v23, v22
  3924: 5700602f                 vxor.vv v0, v22, v0
  3928: 138f0200                 mv t5, t0
  3932: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  3936: 138e0f0e                 addi t3, t6, 0xe0
  3940: 070c0e02                 vle8.v v24, (t3)
  3944: 138e0f0f                 addi t3, t6, 0xf0
  3948: 870c0e02                 vle8.v v25, (t3)
  3952: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  3956: 570aa827                 vand.vv v20, v26, v16
  3960: d73a02a3                 vsrl.vi v21, v16, 4
  3964: 570b8a33                 vrgather.vv v22, v24, v20
  3968: d78b9a33                 vrgather.vv v23, v25, v21
  3972: 570b7b2f                 vxor.vv v22, v23, v22
  3976: 5702622f                 vxor.vv v4, v22, v4
  3980: 83330503                 ld t2, 0x30(a0)
  3984: b3836300                 add t2, t2, t1
  3988: 07880302                 vle8.v v16, (t2)
  3992: 5706085e                 vmv.v.v v12, v16
  3996: 138f0200                 mv t5, t0
  4000: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4004: 138e0f08                 addi t3, t6, 0x80
  4008: 070c0e02                 vle8.v v24, (t3)
  4012: 138e0f09                 addi t3, t6, 0x90
  4016: 870c0e02                 vle8.v v25, (t3)
  4020: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4024: 570aa827                 vand.vv v20, v26, v16
  4028: d73a02a3                 vsrl.vi v21, v16, 4
  4032: 570b8a33                 vrgather.vv v22, v24, v20
  4036: d78b9a33                 vrgather.vv v23, v25, v21
  4040: 570b7b2f                 vxor.vv v22, v23, v22
  4044: 5700602f                 vxor.vv v0, v22, v0
  4048: 138f0200                 mv t5, t0
  4052: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4056: 138e0f10                 addi t3, t6, 0x100
  4060: 070c0e02                 vle8.v v24, (t3)
  4064: 138e0f11                 addi t3, t6, 0x110
  4068: 870c0e02                 vle8.v v25, (t3)
  4072: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4076: 570aa827                 vand.vv v20, v26, v16
  4080: d73a02a3                 vsrl.vi v21, v16, 4
  4084: 570b8a33                 vrgather.vv v22, v24, v20
  4088: d78b9a33                 vrgather.vv v23, v25, v21
  4092: 570b7b2f                 vxor.vv v22, v23, v22
  4096: 5702622f                 vxor.vv v4, v22, v4
  4100: 83338503                 ld t2, 0x38(a0)
  4104: b3836300                 add t2, t2, t1
  4108: 07880302                 vle8.v v16, (t2)
  4112: 5706062f                 vxor.vv v12, v16, v12
  4116: 138f0200                 mv t5, t0
  4120: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4124: 138e0f12                 addi t3, t6, 0x120
  4128: 070c0e02                 vle8.v v24, (t3)
  4132: 138e0f13                 addi t3, t6, 0x130
  4136: 870c0e02                 vle8.v v25, (t3)
  4140: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4144: 570aa827                 vand.vv v20, v26, v16
  4148: d73a02a3                 vsrl.vi v21, v16, 4
  4152: 570b8a33                 vrgather.vv v22, v24, v20
  4156: d78b9a33                 vrgather.vv v23, v25, v21
  4160: 570b7b2f                 vxor.vv v22, v23, v22
  4164: 5700602f                 vxor.vv v0, v22, v0
  4168: 138f0200                 mv t5, t0
  4172: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4176: 138e0f14                 addi t3, t6, 0x140
  4180: 070c0e02                 vle8.v v24, (t3)
  4184: 138e0f15                 addi t3, t6, 0x150
  4188: 870c0e02                 vle8.v v25, (t3)
  4192: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4196: 570aa827                 vand.vv v20, v26, v16
  4200: d73a02a3                 vsrl.vi v21, v16, 4
  4204: 570b8a33                 vrgather.vv v22, v24, v20
  4208: d78b9a33                 vrgather.vv v23, v25, v21
  4212: 570b7b2f                 vxor.vv v22, v23, v22
  4216: 5702622f                 vxor.vv v4, v22, v4
  4220: 83330504                 ld t2, 0x40(a0)
  4224: b3836300                 add t2, t2, t1
  4228: 07880302                 vle8.v v16, (t2)
  4232: 5706062f                 vxor.vv v12, v16, v12
  4236: 138f0200                 mv t5, t0
  4240: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4244: 138e0f0a                 addi t3, t6, 0xa0
  4248: 070c0e02                 vle8.v v24, (t3)
  4252: 138e0f0b                 addi t3, t6, 0xb0
  4256: 870c0e02                 vle8.v v25, (t3)
  4260: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4264: 570aa827                 vand.vv v20, v26, v16
  4268: d73a02a3                 vsrl.vi v21, v16, 4
  4272: 570b8a33                 vrgather.vv v22, v24, v20
  4276: d78b9a33                 vrgather.vv v23, v25, v21
  4280: 570b7b2f                 vxor.vv v22, v23, v22
  4284: 5700602f                 vxor.vv v0, v22, v0
  4288: 138f0200                 mv t5, t0
  4292: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4296: 138e0f16                 addi t3, t6, 0x160
  4300: 070c0e02                 vle8.v v24, (t3)
  4304: 138e0f17                 addi t3, t6, 0x170
  4308: 870c0e02                 vle8.v v25, (t3)
  4312: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4316: 570aa827                 vand.vv v20, v26, v16
  4320: d73a02a3                 vsrl.vi v21, v16, 4
  4324: 570b8a33                 vrgather.vv v22, v24, v20
  4328: d78b9a33                 vrgather.vv v23, v25, v21
  4332: 570b7b2f                 vxor.vv v22, v23, v22
  4336: 5702622f                 vxor.vv v4, v22, v4
  4340: 83338504                 ld t2, 0x48(a0)
  4344: b3836300                 add t2, t2, t1
  4348: 07880302                 vle8.v v16, (t2)
  4352: 5706062f                 vxor.vv v12, v16, v12
  4356: 138f0200                 mv t5, t0
  4360: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4364: 138e0f18                 addi t3, t6, 0x180
  4368: 070c0e02                 vle8.v v24, (t3)
  4372: 138e0f19                 addi t3, t6, 0x190
  4376: 870c0e02                 vle8.v v25, (t3)
  4380: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4384: 570aa827                 vand.vv v20, v26, v16
  4388: d73a02a3                 vsrl.vi v21, v16, 4
  4392: 570b8a33                 vrgather.vv v22, v24, v20
  4396: d78b9a33                 vrgather.vv v23, v25, v21
  4400: 570b7b2f                 vxor.vv v22, v23, v22
  4404: 5700602f                 vxor.vv v0, v22, v0
  4408: 138f0200                 mv t5, t0
  4412: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4416: 138e0f1a                 addi t3, t6, 0x1a0
  4420: 070c0e02                 vle8.v v24, (t3)
  4424: 138e0f1b                 addi t3, t6, 0x1b0
  4428: 870c0e02                 vle8.v v25, (t3)
  4432: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4436: 570aa827                 vand.vv v20, v26, v16
  4440: d73a02a3                 vsrl.vi v21, v16, 4
  4444: 570b8a33                 vrgather.vv v22, v24, v20
  4448: d78b9a33                 vrgather.vv v23, v25, v21
  4452: 570b7b2f                 vxor.vv v22, v23, v22
  4456: 5702622f                 vxor.vv v4, v22, v4
  4460: 83330505                 ld t2, 0x50(a0)
  4464: b3836300                 add t2, t2, t1
  4468: 07880302                 vle8.v v16, (t2)
  4472: 5706062f                 vxor.vv v12, v16, v12
  4476: 138f0200                 mv t5, t0
  4480: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4484: 138e0f0e                 addi t3, t6, 0xe0
  4488: 070c0e02                 vle8.v v24, (t3)
  4492: 138e0f0f                 addi t3, t6, 0xf0
  4496: 870c0e02                 vle8.v v25, (t3)
  4500: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4504: 570aa827                 vand.vv v20, v26, v16
  4508: d73a02a3                 vsrl.vi v21, v16, 4
  4512: 570b8a33                 vrgather.vv v22, v24, v20
  4516: d78b9a33                 vrgather.vv v23, v25, v21
  4520: 570b7b2f                 vxor.vv v22, v23, v22
  4524: 5700602f                 vxor.vv v0, v22, v0
  4528: 138f0200                 mv t5, t0
  4532: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4536: 138e0f1c                 addi t3, t6, 0x1c0
  4540: 070c0e02                 vle8.v v24, (t3)
  4544: 138e0f1d                 addi t3, t6, 0x1d0
  4548: 870c0e02                 vle8.v v25, (t3)
  4552: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4556: 570aa827                 vand.vv v20, v26, v16
  4560: d73a02a3                 vsrl.vi v21, v16, 4
  4564: 570b8a33                 vrgather.vv v22, v24, v20
  4568: d78b9a33                 vrgather.vv v23, v25, v21
  4572: 570b7b2f                 vxor.vv v22, v23, v22
  4576: 5702622f                 vxor.vv v4, v22, v4
  4580: 83338505                 ld t2, 0x58(a0)
  4584: b3836300                 add t2, t2, t1
  4588: 07880302                 vle8.v v16, (t2)
  4592: 5706062f                 vxor.vv v12, v16, v12
  4596: 138f0200                 mv t5, t0
  4600: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4604: 138e0f1e                 addi t3, t6, 0x1e0
  4608: 070c0e02                 vle8.v v24, (t3)
  4612: 138e0f1f                 addi t3, t6, 0x1f0
  4616: 870c0e02                 vle8.v v25, (t3)
  4620: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4624: 570aa827                 vand.vv v20, v26, v16
  4628: d73a02a3                 vsrl.vi v21, v16, 4
  4632: 570b8a33                 vrgather.vv v22, v24, v20
  4636: d78b9a33                 vrgather.vv v23, v25, v21
  4640: 570b7b2f                 vxor.vv v22, v23, v22
  4644: 5700602f                 vxor.vv v0, v22, v0
  4648: 138f0200                 mv t5, t0
  4652: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  4656: 138e0f20                 addi t3, t6, 0x200
  4660: 070c0e02                 vle8.v v24, (t3)
  4664: 138e0f21                 addi t3, t6, 0x210
  4668: 870c0e02                 vle8.v v25, (t3)
  4672: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  4676: 570aa827                 vand.vv v20, v26, v16
  4680: d73a02a3                 vsrl.vi v21, v16, 4
  4684: 570b8a33                 vrgather.vv v22, v24, v20
  4688: d78b9a33                 vrgather.vv v23, v25, v21
  4692: 570b7b2f                 vxor.vv v22, v23, v22
  4696: 5702622f                 vxor.vv v4, v22, v4
  4700: 83b30500                 ld t2, 0(a1)
  4704: b3836300                 add t2, t2, t1
  4708: 27800302                 vse8.v v0, (t2)
  4712: 83b38500                 ld t2, 8(a1)
  4716: b3836300                 add t2, t2, t1
  4720: 27820302                 vse8.v v4, (t2)
  4724: 83b30501                 ld t2, 0x10(a1)
  4728: b3836300                 add t2, t2, t1
  4732: 27840302                 vse8.v v8, (t2)
  4736: 83b38501                 ld t2, 0x18(a1)
  4740: b3836300                 add t2, t2, t1
  4744: 27860302                 vse8.v v12, (t2)
  4748: 33035300                 add t1, t1, t0
  4752: 33065640                 sub a2, a2, t0
  4756: 6ff0dfa7                 j 0xd10
  4760: 67800000                 ret 
  4764: 0000                     unimp 
  4766: 0000                     unimp 
