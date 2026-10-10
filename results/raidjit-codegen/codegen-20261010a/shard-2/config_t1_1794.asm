; config_t1_1794 target=rvv-zvbc32e twin=baseline topo=rs(6,3)/enc op=encode U=1
; code_size=1616 twins=[] uops_per_byte=5.203125
     0: 63140600                 bnez a2, 8
     4: 6f004064                 j 0x648
     8: 970f0000                 auipc t6, 0
    12: 938f8f64                 addi t6, t6, 0x648
    16: 577f100c                 vsetvli t5, zero, e8, m2, ta, ma
    20: 130eb001                 li t3, 0x1b
    24: 574c0e5e                 vmv.v.x v24, t3
    28: 13030000                 li t1, 0
    32: d772100c                 vsetvli t0, zero, e8, m2, ta, ma
    36: 93961200                 slli a3, t0, 1
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00403e                 j 0x414
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07860302                 vle8.v v12, (t2)
    76: b3835300                 add t2, t2, t0
    80: 07870302                 vle8.v v14, (t2)
    84: 5700065e                 vmv.v.v v0, v12
    88: 5701075e                 vmv.v.v v2, v14
    92: 130ed008                 li t3, 0x8d
    96: 574b0e5e                 vmv.v.x v22, t3
   100: 5728cb32                 vclmul.vv v16, v12, v22
   104: 5729cb36                 vclmulh.vv v18, v12, v22
   108: 572a2c33                 vclmul.vv v20, v18, v24
   112: 57292c37                 vclmulh.vv v18, v18, v24
   116: 57292c33                 vclmul.vv v18, v18, v24
   120: 5708482f                 vxor.vv v16, v20, v16
   124: 5708282f                 vxor.vv v16, v18, v16
   128: 5702085e                 vmv.v.v v4, v16
   132: 5728eb32                 vclmul.vv v16, v14, v22
   136: 5729eb36                 vclmulh.vv v18, v14, v22
   140: 572a2c33                 vclmul.vv v20, v18, v24
   144: 57292c37                 vclmulh.vv v18, v18, v24
   148: 57292c33                 vclmul.vv v18, v18, v24
   152: 5708482f                 vxor.vv v16, v20, v16
   156: 5708282f                 vxor.vv v16, v18, v16
   160: 5703085e                 vmv.v.v v6, v16
   164: 130eb00c                 li t3, 0xcb
   168: 574b0e5e                 vmv.v.x v22, t3
   172: 5728cb32                 vclmul.vv v16, v12, v22
   176: 5729cb36                 vclmulh.vv v18, v12, v22
   180: 572a2c33                 vclmul.vv v20, v18, v24
   184: 57292c37                 vclmulh.vv v18, v18, v24
   188: 57292c33                 vclmul.vv v18, v18, v24
   192: 5708482f                 vxor.vv v16, v20, v16
   196: 5708282f                 vxor.vv v16, v18, v16
   200: 5704085e                 vmv.v.v v8, v16
   204: 5728eb32                 vclmul.vv v16, v14, v22
   208: 5729eb36                 vclmulh.vv v18, v14, v22
   212: 572a2c33                 vclmul.vv v20, v18, v24
   216: 57292c37                 vclmulh.vv v18, v18, v24
   220: 57292c33                 vclmul.vv v18, v18, v24
   224: 5708482f                 vxor.vv v16, v20, v16
   228: 5708282f                 vxor.vv v16, v18, v16
   232: 5705085e                 vmv.v.v v10, v16
   236: 83338500                 ld t2, 8(a0)
   240: b3836300                 add t2, t2, t1
   244: 07860302                 vle8.v v12, (t2)
   248: b3835300                 add t2, t2, t0
   252: 07870302                 vle8.v v14, (t2)
   256: 5700c02e                 vxor.vv v0, v12, v0
   260: 5701e12e                 vxor.vv v2, v14, v2
   264: 130eb00c                 li t3, 0xcb
   268: 574b0e5e                 vmv.v.x v22, t3
   272: 5728cb32                 vclmul.vv v16, v12, v22
   276: 5729cb36                 vclmulh.vv v18, v12, v22
   280: 572a2c33                 vclmul.vv v20, v18, v24
   284: 57292c37                 vclmulh.vv v18, v18, v24
   288: 57292c33                 vclmul.vv v18, v18, v24
   292: 5708482f                 vxor.vv v16, v20, v16
   296: 5708282f                 vxor.vv v16, v18, v16
   300: 5702022f                 vxor.vv v4, v16, v4
   304: 5728eb32                 vclmul.vv v16, v14, v22
   308: 5729eb36                 vclmulh.vv v18, v14, v22
   312: 572a2c33                 vclmul.vv v20, v18, v24
   316: 57292c37                 vclmulh.vv v18, v18, v24
   320: 57292c33                 vclmul.vv v18, v18, v24
   324: 5708482f                 vxor.vv v16, v20, v16
   328: 5708282f                 vxor.vv v16, v18, v16
   332: 5703032f                 vxor.vv v6, v16, v6
   336: 130e8000                 li t3, 8
   340: 574b0e5e                 vmv.v.x v22, t3
   344: 5728cb32                 vclmul.vv v16, v12, v22
   348: 5729cb36                 vclmulh.vv v18, v12, v22
   352: 572a2c33                 vclmul.vv v20, v18, v24
   356: 57292c37                 vclmulh.vv v18, v18, v24
   360: 57292c33                 vclmul.vv v18, v18, v24
   364: 5708482f                 vxor.vv v16, v20, v16
   368: 5708282f                 vxor.vv v16, v18, v16
   372: 5704042f                 vxor.vv v8, v16, v8
   376: 5728eb32                 vclmul.vv v16, v14, v22
   380: 5729eb36                 vclmulh.vv v18, v14, v22
   384: 572a2c33                 vclmul.vv v20, v18, v24
   388: 57292c37                 vclmulh.vv v18, v18, v24
   392: 57292c33                 vclmul.vv v18, v18, v24
   396: 5708482f                 vxor.vv v16, v20, v16
   400: 5708282f                 vxor.vv v16, v18, v16
   404: 5705052f                 vxor.vv v10, v16, v10
   408: 83330501                 ld t2, 0x10(a0)
   412: b3836300                 add t2, t2, t1
   416: 07860302                 vle8.v v12, (t2)
   420: b3835300                 add t2, t2, t0
   424: 07870302                 vle8.v v14, (t2)
   428: 5700c02e                 vxor.vv v0, v12, v0
   432: 5701e12e                 vxor.vv v2, v14, v2
   436: 130e8000                 li t3, 8
   440: 574b0e5e                 vmv.v.x v22, t3
   444: 5728cb32                 vclmul.vv v16, v12, v22
   448: 5729cb36                 vclmulh.vv v18, v12, v22
   452: 572a2c33                 vclmul.vv v20, v18, v24
   456: 57292c37                 vclmulh.vv v18, v18, v24
   460: 57292c33                 vclmul.vv v18, v18, v24
   464: 5708482f                 vxor.vv v16, v20, v16
   468: 5708282f                 vxor.vv v16, v18, v16
   472: 5702022f                 vxor.vv v4, v16, v4
   476: 5728eb32                 vclmul.vv v16, v14, v22
   480: 5729eb36                 vclmulh.vv v18, v14, v22
   484: 572a2c33                 vclmul.vv v20, v18, v24
   488: 57292c37                 vclmulh.vv v18, v18, v24
   492: 57292c33                 vclmul.vv v18, v18, v24
   496: 5708482f                 vxor.vv v16, v20, v16
   500: 5708282f                 vxor.vv v16, v18, v16
   504: 5703032f                 vxor.vv v6, v16, v6
   508: 130e2000                 li t3, 2
   512: 574b0e5e                 vmv.v.x v22, t3
   516: 5728cb32                 vclmul.vv v16, v12, v22
   520: 5729cb36                 vclmulh.vv v18, v12, v22
   524: 572a2c33                 vclmul.vv v20, v18, v24
   528: 57292c37                 vclmulh.vv v18, v18, v24
   532: 57292c33                 vclmul.vv v18, v18, v24
   536: 5708482f                 vxor.vv v16, v20, v16
   540: 5708282f                 vxor.vv v16, v18, v16
   544: 5704042f                 vxor.vv v8, v16, v8
   548: 5728eb32                 vclmul.vv v16, v14, v22
   552: 5729eb36                 vclmulh.vv v18, v14, v22
   556: 572a2c33                 vclmul.vv v20, v18, v24
   560: 57292c37                 vclmulh.vv v18, v18, v24
   564: 57292c33                 vclmul.vv v18, v18, v24
   568: 5708482f                 vxor.vv v16, v20, v16
   572: 5708282f                 vxor.vv v16, v18, v16
   576: 5705052f                 vxor.vv v10, v16, v10
   580: 83338501                 ld t2, 0x18(a0)
   584: b3836300                 add t2, t2, t1
   588: 07860302                 vle8.v v12, (t2)
   592: b3835300                 add t2, t2, t0
   596: 07870302                 vle8.v v14, (t2)
   600: 5700c02e                 vxor.vv v0, v12, v0
   604: 5701e12e                 vxor.vv v2, v14, v2
   608: 130e4000                 li t3, 4
   612: 574b0e5e                 vmv.v.x v22, t3
   616: 5728cb32                 vclmul.vv v16, v12, v22
   620: 5729cb36                 vclmulh.vv v18, v12, v22
   624: 572a2c33                 vclmul.vv v20, v18, v24
   628: 57292c37                 vclmulh.vv v18, v18, v24
   632: 57292c33                 vclmul.vv v18, v18, v24
   636: 5708482f                 vxor.vv v16, v20, v16
   640: 5708282f                 vxor.vv v16, v18, v16
   644: 5702022f                 vxor.vv v4, v16, v4
   648: 5728eb32                 vclmul.vv v16, v14, v22
   652: 5729eb36                 vclmulh.vv v18, v14, v22
   656: 572a2c33                 vclmul.vv v20, v18, v24
   660: 57292c37                 vclmulh.vv v18, v18, v24
   664: 57292c33                 vclmul.vv v18, v18, v24
   668: 5708482f                 vxor.vv v16, v20, v16
   672: 5708282f                 vxor.vv v16, v18, v16
   676: 5703032f                 vxor.vv v6, v16, v6
   680: 130ed008                 li t3, 0x8d
   684: 574b0e5e                 vmv.v.x v22, t3
   688: 5728cb32                 vclmul.vv v16, v12, v22
   692: 5729cb36                 vclmulh.vv v18, v12, v22
   696: 572a2c33                 vclmul.vv v20, v18, v24
   700: 57292c37                 vclmulh.vv v18, v18, v24
   704: 57292c33                 vclmul.vv v18, v18, v24
   708: 5708482f                 vxor.vv v16, v20, v16
   712: 5708282f                 vxor.vv v16, v18, v16
   716: 5704042f                 vxor.vv v8, v16, v8
   720: 5728eb32                 vclmul.vv v16, v14, v22
   724: 5729eb36                 vclmulh.vv v18, v14, v22
   728: 572a2c33                 vclmul.vv v20, v18, v24
   732: 57292c37                 vclmulh.vv v18, v18, v24
   736: 57292c33                 vclmul.vv v18, v18, v24
   740: 5708482f                 vxor.vv v16, v20, v16
   744: 5708282f                 vxor.vv v16, v18, v16
   748: 5705052f                 vxor.vv v10, v16, v10
   752: 83330502                 ld t2, 0x20(a0)
   756: b3836300                 add t2, t2, t1
   760: 07860302                 vle8.v v12, (t2)
   764: b3835300                 add t2, t2, t0
   768: 07870302                 vle8.v v14, (t2)
   772: 5700c02e                 vxor.vv v0, v12, v0
   776: 5701e12e                 vxor.vv v2, v14, v2
   780: 130e2000                 li t3, 2
   784: 574b0e5e                 vmv.v.x v22, t3
   788: 5728cb32                 vclmul.vv v16, v12, v22
   792: 5729cb36                 vclmulh.vv v18, v12, v22
   796: 572a2c33                 vclmul.vv v20, v18, v24
   800: 57292c37                 vclmulh.vv v18, v18, v24
   804: 57292c33                 vclmul.vv v18, v18, v24
   808: 5708482f                 vxor.vv v16, v20, v16
   812: 5708282f                 vxor.vv v16, v18, v16
   816: 5702022f                 vxor.vv v4, v16, v4
   820: 5728eb32                 vclmul.vv v16, v14, v22
   824: 5729eb36                 vclmulh.vv v18, v14, v22
   828: 572a2c33                 vclmul.vv v20, v18, v24
   832: 57292c37                 vclmulh.vv v18, v18, v24
   836: 57292c33                 vclmul.vv v18, v18, v24
   840: 5708482f                 vxor.vv v16, v20, v16
   844: 5708282f                 vxor.vv v16, v18, v16
   848: 5703032f                 vxor.vv v6, v16, v6
   852: 130e4000                 li t3, 4
   856: 574b0e5e                 vmv.v.x v22, t3
   860: 5728cb32                 vclmul.vv v16, v12, v22
   864: 5729cb36                 vclmulh.vv v18, v12, v22
   868: 572a2c33                 vclmul.vv v20, v18, v24
   872: 57292c37                 vclmulh.vv v18, v18, v24
   876: 57292c33                 vclmul.vv v18, v18, v24
   880: 5708482f                 vxor.vv v16, v20, v16
   884: 5708282f                 vxor.vv v16, v18, v16
   888: 5704042f                 vxor.vv v8, v16, v8
   892: 5728eb32                 vclmul.vv v16, v14, v22
   896: 5729eb36                 vclmulh.vv v18, v14, v22
   900: 572a2c33                 vclmul.vv v20, v18, v24
   904: 57292c37                 vclmulh.vv v18, v18, v24
   908: 57292c33                 vclmul.vv v18, v18, v24
   912: 5708482f                 vxor.vv v16, v20, v16
   916: 5708282f                 vxor.vv v16, v18, v16
   920: 5705052f                 vxor.vv v10, v16, v10
   924: 83338502                 ld t2, 0x28(a0)
   928: b3836300                 add t2, t2, t1
   932: 07860302                 vle8.v v12, (t2)
   936: b3835300                 add t2, t2, t0
   940: 07870302                 vle8.v v14, (t2)
   944: 5700c02e                 vxor.vv v0, v12, v0
   948: 5701e12e                 vxor.vv v2, v14, v2
   952: 5702c22e                 vxor.vv v4, v12, v4
   956: 5703e32e                 vxor.vv v6, v14, v6
   960: 5704c42e                 vxor.vv v8, v12, v8
   964: 5705e52e                 vxor.vv v10, v14, v10
   968: 83b30500                 ld t2, 0(a1)
   972: b3836300                 add t2, t2, t1
   976: 27800302                 vse8.v v0, (t2)
   980: b3835300                 add t2, t2, t0
   984: 27810302                 vse8.v v2, (t2)
   988: 83b38500                 ld t2, 8(a1)
   992: b3836300                 add t2, t2, t1
   996: 27820302                 vse8.v v4, (t2)
  1000: b3835300                 add t2, t2, t0
  1004: 27830302                 vse8.v v6, (t2)
  1008: 83b30501                 ld t2, 0x10(a1)
  1012: b3836300                 add t2, t2, t1
  1016: 27840302                 vse8.v v8, (t2)
  1020: b3835300                 add t2, t2, t0
  1024: 27850302                 vse8.v v10, (t2)
  1028: 3303d300                 add t1, t1, a3
  1032: 1307f7ff                 addi a4, a4, -1
  1036: e31a07c2                 bnez a4, 0x40
  1040: 33066640                 sub a2, a2, t1
  1044: d772160c                 vsetvli t0, a2, e8, m2, ta, ma
  1048: 63940200                 bnez t0, 0x420
  1052: 6f00c022                 j 0x648
  1056: 83330500                 ld t2, 0(a0)
  1060: b3836300                 add t2, t2, t1
  1064: 07860302                 vle8.v v12, (t2)
  1068: 5700065e                 vmv.v.v v0, v12
  1072: 130ed008                 li t3, 0x8d
  1076: 574b0e5e                 vmv.v.x v22, t3
  1080: 5728cb32                 vclmul.vv v16, v12, v22
  1084: 5729cb36                 vclmulh.vv v18, v12, v22
  1088: 572a2c33                 vclmul.vv v20, v18, v24
  1092: 57292c37                 vclmulh.vv v18, v18, v24
  1096: 57292c33                 vclmul.vv v18, v18, v24
  1100: 5708482f                 vxor.vv v16, v20, v16
  1104: 5708282f                 vxor.vv v16, v18, v16
  1108: 5702085e                 vmv.v.v v4, v16
  1112: 130eb00c                 li t3, 0xcb
  1116: 574b0e5e                 vmv.v.x v22, t3
  1120: 5728cb32                 vclmul.vv v16, v12, v22
  1124: 5729cb36                 vclmulh.vv v18, v12, v22
  1128: 572a2c33                 vclmul.vv v20, v18, v24
  1132: 57292c37                 vclmulh.vv v18, v18, v24
  1136: 57292c33                 vclmul.vv v18, v18, v24
  1140: 5708482f                 vxor.vv v16, v20, v16
  1144: 5708282f                 vxor.vv v16, v18, v16
  1148: 5704085e                 vmv.v.v v8, v16
  1152: 83338500                 ld t2, 8(a0)
  1156: b3836300                 add t2, t2, t1
  1160: 07860302                 vle8.v v12, (t2)
  1164: 5700c02e                 vxor.vv v0, v12, v0
  1168: 130eb00c                 li t3, 0xcb
  1172: 574b0e5e                 vmv.v.x v22, t3
  1176: 5728cb32                 vclmul.vv v16, v12, v22
  1180: 5729cb36                 vclmulh.vv v18, v12, v22
  1184: 572a2c33                 vclmul.vv v20, v18, v24
  1188: 57292c37                 vclmulh.vv v18, v18, v24
  1192: 57292c33                 vclmul.vv v18, v18, v24
  1196: 5708482f                 vxor.vv v16, v20, v16
  1200: 5708282f                 vxor.vv v16, v18, v16
  1204: 5702022f                 vxor.vv v4, v16, v4
  1208: 130e8000                 li t3, 8
  1212: 574b0e5e                 vmv.v.x v22, t3
  1216: 5728cb32                 vclmul.vv v16, v12, v22
  1220: 5729cb36                 vclmulh.vv v18, v12, v22
  1224: 572a2c33                 vclmul.vv v20, v18, v24
  1228: 57292c37                 vclmulh.vv v18, v18, v24
  1232: 57292c33                 vclmul.vv v18, v18, v24
  1236: 5708482f                 vxor.vv v16, v20, v16
  1240: 5708282f                 vxor.vv v16, v18, v16
  1244: 5704042f                 vxor.vv v8, v16, v8
  1248: 83330501                 ld t2, 0x10(a0)
  1252: b3836300                 add t2, t2, t1
  1256: 07860302                 vle8.v v12, (t2)
  1260: 5700c02e                 vxor.vv v0, v12, v0
  1264: 130e8000                 li t3, 8
  1268: 574b0e5e                 vmv.v.x v22, t3
  1272: 5728cb32                 vclmul.vv v16, v12, v22
  1276: 5729cb36                 vclmulh.vv v18, v12, v22
  1280: 572a2c33                 vclmul.vv v20, v18, v24
  1284: 57292c37                 vclmulh.vv v18, v18, v24
  1288: 57292c33                 vclmul.vv v18, v18, v24
  1292: 5708482f                 vxor.vv v16, v20, v16
  1296: 5708282f                 vxor.vv v16, v18, v16
  1300: 5702022f                 vxor.vv v4, v16, v4
  1304: 130e2000                 li t3, 2
  1308: 574b0e5e                 vmv.v.x v22, t3
  1312: 5728cb32                 vclmul.vv v16, v12, v22
  1316: 5729cb36                 vclmulh.vv v18, v12, v22
  1320: 572a2c33                 vclmul.vv v20, v18, v24
  1324: 57292c37                 vclmulh.vv v18, v18, v24
  1328: 57292c33                 vclmul.vv v18, v18, v24
  1332: 5708482f                 vxor.vv v16, v20, v16
  1336: 5708282f                 vxor.vv v16, v18, v16
  1340: 5704042f                 vxor.vv v8, v16, v8
  1344: 83338501                 ld t2, 0x18(a0)
  1348: b3836300                 add t2, t2, t1
  1352: 07860302                 vle8.v v12, (t2)
  1356: 5700c02e                 vxor.vv v0, v12, v0
  1360: 130e4000                 li t3, 4
  1364: 574b0e5e                 vmv.v.x v22, t3
  1368: 5728cb32                 vclmul.vv v16, v12, v22
  1372: 5729cb36                 vclmulh.vv v18, v12, v22
  1376: 572a2c33                 vclmul.vv v20, v18, v24
  1380: 57292c37                 vclmulh.vv v18, v18, v24
  1384: 57292c33                 vclmul.vv v18, v18, v24
  1388: 5708482f                 vxor.vv v16, v20, v16
  1392: 5708282f                 vxor.vv v16, v18, v16
  1396: 5702022f                 vxor.vv v4, v16, v4
  1400: 130ed008                 li t3, 0x8d
  1404: 574b0e5e                 vmv.v.x v22, t3
  1408: 5728cb32                 vclmul.vv v16, v12, v22
  1412: 5729cb36                 vclmulh.vv v18, v12, v22
  1416: 572a2c33                 vclmul.vv v20, v18, v24
  1420: 57292c37                 vclmulh.vv v18, v18, v24
  1424: 57292c33                 vclmul.vv v18, v18, v24
  1428: 5708482f                 vxor.vv v16, v20, v16
  1432: 5708282f                 vxor.vv v16, v18, v16
  1436: 5704042f                 vxor.vv v8, v16, v8
  1440: 83330502                 ld t2, 0x20(a0)
  1444: b3836300                 add t2, t2, t1
  1448: 07860302                 vle8.v v12, (t2)
  1452: 5700c02e                 vxor.vv v0, v12, v0
  1456: 130e2000                 li t3, 2
  1460: 574b0e5e                 vmv.v.x v22, t3
  1464: 5728cb32                 vclmul.vv v16, v12, v22
  1468: 5729cb36                 vclmulh.vv v18, v12, v22
  1472: 572a2c33                 vclmul.vv v20, v18, v24
  1476: 57292c37                 vclmulh.vv v18, v18, v24
  1480: 57292c33                 vclmul.vv v18, v18, v24
  1484: 5708482f                 vxor.vv v16, v20, v16
  1488: 5708282f                 vxor.vv v16, v18, v16
  1492: 5702022f                 vxor.vv v4, v16, v4
  1496: 130e4000                 li t3, 4
  1500: 574b0e5e                 vmv.v.x v22, t3
  1504: 5728cb32                 vclmul.vv v16, v12, v22
  1508: 5729cb36                 vclmulh.vv v18, v12, v22
  1512: 572a2c33                 vclmul.vv v20, v18, v24
  1516: 57292c37                 vclmulh.vv v18, v18, v24
  1520: 57292c33                 vclmul.vv v18, v18, v24
  1524: 5708482f                 vxor.vv v16, v20, v16
  1528: 5708282f                 vxor.vv v16, v18, v16
  1532: 5704042f                 vxor.vv v8, v16, v8
  1536: 83338502                 ld t2, 0x28(a0)
  1540: b3836300                 add t2, t2, t1
  1544: 07860302                 vle8.v v12, (t2)
  1548: 5700c02e                 vxor.vv v0, v12, v0
  1552: 5702c22e                 vxor.vv v4, v12, v4
  1556: 5704c42e                 vxor.vv v8, v12, v8
  1560: 83b30500                 ld t2, 0(a1)
  1564: b3836300                 add t2, t2, t1
  1568: 27800302                 vse8.v v0, (t2)
  1572: 83b38500                 ld t2, 8(a1)
  1576: b3836300                 add t2, t2, t1
  1580: 27820302                 vse8.v v4, (t2)
  1584: 83b30501                 ld t2, 0x10(a1)
  1588: b3836300                 add t2, t2, t1
  1592: 27840302                 vse8.v v8, (t2)
  1596: 33035300                 add t1, t1, t0
  1600: 33065640                 sub a2, a2, t0
  1604: 6ff01fdd                 j 0x414
  1608: 67800000                 ret 
  1612: 0000                     unimp 
  1614: 0000                     unimp 
