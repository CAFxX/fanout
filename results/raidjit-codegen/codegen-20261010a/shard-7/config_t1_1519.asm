; config_t1_1519 target=rvv twin=pipe topo=rs(6,3)/enc op=encode U=1
; code_size=1328 twins=["pipe"] uops_per_byte=16.625
     0: 63140600                 bnez a2, 8
     4: 6f000052                 j 0x524
     8: 970f0000                 auipc t6, 0
    12: 938f8f52                 addi t6, t6, 0x528
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d7450e5e                 vmv.v.x v11, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f008024                 j 0x278
    52: 83330500                 ld t2, 0(a0)
    56: b3836300                 add t2, t2, t1
    60: 07820302                 vle8.v v4, (t2)
    64: 5700025e                 vmv.v.v v0, v4
    68: 138e0f00                 mv t3, t6
    72: 87040e02                 vle8.v v9, (t3)
    76: 138e0f01                 addi t3, t6, 0x10
    80: 07050e02                 vle8.v v10, (t3)
    84: d702b226                 vand.vv v5, v11, v4
    88: 573342a2                 vsrl.vi v6, v4, 4
    92: d7839232                 vrgather.vv v7, v9, v5
    96: 5704a332                 vrgather.vv v8, v10, v6
   100: d783832e                 vxor.vv v7, v8, v7
   104: d780035e                 vmv.v.v v1, v7
   108: 138e0f02                 addi t3, t6, 0x20
   112: 87040e02                 vle8.v v9, (t3)
   116: 138e0f03                 addi t3, t6, 0x30
   120: 07050e02                 vle8.v v10, (t3)
   124: d702b226                 vand.vv v5, v11, v4
   128: 573342a2                 vsrl.vi v6, v4, 4
   132: d7839232                 vrgather.vv v7, v9, v5
   136: 5704a332                 vrgather.vv v8, v10, v6
   140: d783832e                 vxor.vv v7, v8, v7
   144: 5781035e                 vmv.v.v v2, v7
   148: 83338500                 ld t2, 8(a0)
   152: b3836300                 add t2, t2, t1
   156: 87810302                 vle8.v v3, (t2)
   160: 5700302e                 vxor.vv v0, v3, v0
   164: 138e0f02                 addi t3, t6, 0x20
   168: 87040e02                 vle8.v v9, (t3)
   172: 138e0f03                 addi t3, t6, 0x30
   176: 07050e02                 vle8.v v10, (t3)
   180: d782b126                 vand.vv v5, v11, v3
   184: 573332a2                 vsrl.vi v6, v3, 4
   188: d7839232                 vrgather.vv v7, v9, v5
   192: 5704a332                 vrgather.vv v8, v10, v6
   196: d783832e                 vxor.vv v7, v8, v7
   200: d780702e                 vxor.vv v1, v7, v1
   204: 138e0f04                 addi t3, t6, 0x40
   208: 87040e02                 vle8.v v9, (t3)
   212: 138e0f05                 addi t3, t6, 0x50
   216: 07050e02                 vle8.v v10, (t3)
   220: d782b126                 vand.vv v5, v11, v3
   224: 573332a2                 vsrl.vi v6, v3, 4
   228: d7839232                 vrgather.vv v7, v9, v5
   232: 5704a332                 vrgather.vv v8, v10, v6
   236: d783832e                 vxor.vv v7, v8, v7
   240: 5701712e                 vxor.vv v2, v7, v2
   244: 83330501                 ld t2, 0x10(a0)
   248: b3836300                 add t2, t2, t1
   252: 87810302                 vle8.v v3, (t2)
   256: 5700302e                 vxor.vv v0, v3, v0
   260: 138e0f04                 addi t3, t6, 0x40
   264: 87040e02                 vle8.v v9, (t3)
   268: 138e0f05                 addi t3, t6, 0x50
   272: 07050e02                 vle8.v v10, (t3)
   276: d782b126                 vand.vv v5, v11, v3
   280: 573332a2                 vsrl.vi v6, v3, 4
   284: d7839232                 vrgather.vv v7, v9, v5
   288: 5704a332                 vrgather.vv v8, v10, v6
   292: d783832e                 vxor.vv v7, v8, v7
   296: d780702e                 vxor.vv v1, v7, v1
   300: 138e0f06                 addi t3, t6, 0x60
   304: 87040e02                 vle8.v v9, (t3)
   308: 138e0f07                 addi t3, t6, 0x70
   312: 07050e02                 vle8.v v10, (t3)
   316: d782b126                 vand.vv v5, v11, v3
   320: 573332a2                 vsrl.vi v6, v3, 4
   324: d7839232                 vrgather.vv v7, v9, v5
   328: 5704a332                 vrgather.vv v8, v10, v6
   332: d783832e                 vxor.vv v7, v8, v7
   336: 5701712e                 vxor.vv v2, v7, v2
   340: 83338501                 ld t2, 0x18(a0)
   344: b3836300                 add t2, t2, t1
   348: 87810302                 vle8.v v3, (t2)
   352: 5700302e                 vxor.vv v0, v3, v0
   356: 138e0f08                 addi t3, t6, 0x80
   360: 87040e02                 vle8.v v9, (t3)
   364: 138e0f09                 addi t3, t6, 0x90
   368: 07050e02                 vle8.v v10, (t3)
   372: d782b126                 vand.vv v5, v11, v3
   376: 573332a2                 vsrl.vi v6, v3, 4
   380: d7839232                 vrgather.vv v7, v9, v5
   384: 5704a332                 vrgather.vv v8, v10, v6
   388: d783832e                 vxor.vv v7, v8, v7
   392: d780702e                 vxor.vv v1, v7, v1
   396: 138e0f00                 mv t3, t6
   400: 87040e02                 vle8.v v9, (t3)
   404: 138e0f01                 addi t3, t6, 0x10
   408: 07050e02                 vle8.v v10, (t3)
   412: d782b126                 vand.vv v5, v11, v3
   416: 573332a2                 vsrl.vi v6, v3, 4
   420: d7839232                 vrgather.vv v7, v9, v5
   424: 5704a332                 vrgather.vv v8, v10, v6
   428: d783832e                 vxor.vv v7, v8, v7
   432: 5701712e                 vxor.vv v2, v7, v2
   436: 83330502                 ld t2, 0x20(a0)
   440: b3836300                 add t2, t2, t1
   444: 87810302                 vle8.v v3, (t2)
   448: 5700302e                 vxor.vv v0, v3, v0
   452: 138e0f06                 addi t3, t6, 0x60
   456: 87040e02                 vle8.v v9, (t3)
   460: 138e0f07                 addi t3, t6, 0x70
   464: 07050e02                 vle8.v v10, (t3)
   468: d782b126                 vand.vv v5, v11, v3
   472: 573332a2                 vsrl.vi v6, v3, 4
   476: d7839232                 vrgather.vv v7, v9, v5
   480: 5704a332                 vrgather.vv v8, v10, v6
   484: d783832e                 vxor.vv v7, v8, v7
   488: d780702e                 vxor.vv v1, v7, v1
   492: 138e0f08                 addi t3, t6, 0x80
   496: 87040e02                 vle8.v v9, (t3)
   500: 138e0f09                 addi t3, t6, 0x90
   504: 07050e02                 vle8.v v10, (t3)
   508: d782b126                 vand.vv v5, v11, v3
   512: 573332a2                 vsrl.vi v6, v3, 4
   516: d7839232                 vrgather.vv v7, v9, v5
   520: 5704a332                 vrgather.vv v8, v10, v6
   524: d783832e                 vxor.vv v7, v8, v7
   528: 5701712e                 vxor.vv v2, v7, v2
   532: 83338502                 ld t2, 0x28(a0)
   536: b3836300                 add t2, t2, t1
   540: 87810302                 vle8.v v3, (t2)
   544: 5700302e                 vxor.vv v0, v3, v0
   548: d780302e                 vxor.vv v1, v3, v1
   552: 5701312e                 vxor.vv v2, v3, v2
   556: 130ff7ff                 addi t5, a4, -1
   560: 630a0f00                 beqz t5, 0x244
   564: 330fd300                 add t5, t1, a3
   568: 83330500                 ld t2, 0(a0)
   572: b383e301                 add t2, t2, t5
   576: 07820302                 vle8.v v4, (t2)
   580: 83b30500                 ld t2, 0(a1)
   584: b3836300                 add t2, t2, t1
   588: 27800302                 vse8.v v0, (t2)
   592: 83b38500                 ld t2, 8(a1)
   596: b3836300                 add t2, t2, t1
   600: a7800302                 vse8.v v1, (t2)
   604: 83b30501                 ld t2, 0x10(a1)
   608: b3836300                 add t2, t2, t1
   612: 27810302                 vse8.v v2, (t2)
   616: 3303d300                 add t1, t1, a3
   620: 1307f7ff                 addi a4, a4, -1
   624: e31807dc                 bnez a4, 0x40
   628: 33066640                 sub a2, a2, t1
   632: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   636: 63940200                 bnez t0, 0x284
   640: 6f00402a                 j 0x524
   644: 83330500                 ld t2, 0(a0)
   648: b3836300                 add t2, t2, t1
   652: 87810302                 vle8.v v3, (t2)
   656: 5780015e                 vmv.v.v v0, v3
   660: 138f0200                 mv t5, t0
   664: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   668: 138e0f00                 mv t3, t6
   672: 87040e02                 vle8.v v9, (t3)
   676: 138e0f01                 addi t3, t6, 0x10
   680: 07050e02                 vle8.v v10, (t3)
   684: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   688: d782b126                 vand.vv v5, v11, v3
   692: 573332a2                 vsrl.vi v6, v3, 4
   696: d7839232                 vrgather.vv v7, v9, v5
   700: 5704a332                 vrgather.vv v8, v10, v6
   704: d783832e                 vxor.vv v7, v8, v7
   708: d780035e                 vmv.v.v v1, v7
   712: 138f0200                 mv t5, t0
   716: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   720: 138e0f02                 addi t3, t6, 0x20
   724: 87040e02                 vle8.v v9, (t3)
   728: 138e0f03                 addi t3, t6, 0x30
   732: 07050e02                 vle8.v v10, (t3)
   736: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   740: d782b126                 vand.vv v5, v11, v3
   744: 573332a2                 vsrl.vi v6, v3, 4
   748: d7839232                 vrgather.vv v7, v9, v5
   752: 5704a332                 vrgather.vv v8, v10, v6
   756: d783832e                 vxor.vv v7, v8, v7
   760: 5781035e                 vmv.v.v v2, v7
   764: 83338500                 ld t2, 8(a0)
   768: b3836300                 add t2, t2, t1
   772: 87810302                 vle8.v v3, (t2)
   776: 5700302e                 vxor.vv v0, v3, v0
   780: 138f0200                 mv t5, t0
   784: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   788: 138e0f02                 addi t3, t6, 0x20
   792: 87040e02                 vle8.v v9, (t3)
   796: 138e0f03                 addi t3, t6, 0x30
   800: 07050e02                 vle8.v v10, (t3)
   804: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   808: d782b126                 vand.vv v5, v11, v3
   812: 573332a2                 vsrl.vi v6, v3, 4
   816: d7839232                 vrgather.vv v7, v9, v5
   820: 5704a332                 vrgather.vv v8, v10, v6
   824: d783832e                 vxor.vv v7, v8, v7
   828: d780702e                 vxor.vv v1, v7, v1
   832: 138f0200                 mv t5, t0
   836: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   840: 138e0f04                 addi t3, t6, 0x40
   844: 87040e02                 vle8.v v9, (t3)
   848: 138e0f05                 addi t3, t6, 0x50
   852: 07050e02                 vle8.v v10, (t3)
   856: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   860: d782b126                 vand.vv v5, v11, v3
   864: 573332a2                 vsrl.vi v6, v3, 4
   868: d7839232                 vrgather.vv v7, v9, v5
   872: 5704a332                 vrgather.vv v8, v10, v6
   876: d783832e                 vxor.vv v7, v8, v7
   880: 5701712e                 vxor.vv v2, v7, v2
   884: 83330501                 ld t2, 0x10(a0)
   888: b3836300                 add t2, t2, t1
   892: 87810302                 vle8.v v3, (t2)
   896: 5700302e                 vxor.vv v0, v3, v0
   900: 138f0200                 mv t5, t0
   904: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   908: 138e0f04                 addi t3, t6, 0x40
   912: 87040e02                 vle8.v v9, (t3)
   916: 138e0f05                 addi t3, t6, 0x50
   920: 07050e02                 vle8.v v10, (t3)
   924: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   928: d782b126                 vand.vv v5, v11, v3
   932: 573332a2                 vsrl.vi v6, v3, 4
   936: d7839232                 vrgather.vv v7, v9, v5
   940: 5704a332                 vrgather.vv v8, v10, v6
   944: d783832e                 vxor.vv v7, v8, v7
   948: d780702e                 vxor.vv v1, v7, v1
   952: 138f0200                 mv t5, t0
   956: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   960: 138e0f06                 addi t3, t6, 0x60
   964: 87040e02                 vle8.v v9, (t3)
   968: 138e0f07                 addi t3, t6, 0x70
   972: 07050e02                 vle8.v v10, (t3)
   976: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   980: d782b126                 vand.vv v5, v11, v3
   984: 573332a2                 vsrl.vi v6, v3, 4
   988: d7839232                 vrgather.vv v7, v9, v5
   992: 5704a332                 vrgather.vv v8, v10, v6
   996: d783832e                 vxor.vv v7, v8, v7
  1000: 5701712e                 vxor.vv v2, v7, v2
  1004: 83338501                 ld t2, 0x18(a0)
  1008: b3836300                 add t2, t2, t1
  1012: 87810302                 vle8.v v3, (t2)
  1016: 5700302e                 vxor.vv v0, v3, v0
  1020: 138f0200                 mv t5, t0
  1024: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1028: 138e0f08                 addi t3, t6, 0x80
  1032: 87040e02                 vle8.v v9, (t3)
  1036: 138e0f09                 addi t3, t6, 0x90
  1040: 07050e02                 vle8.v v10, (t3)
  1044: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1048: d782b126                 vand.vv v5, v11, v3
  1052: 573332a2                 vsrl.vi v6, v3, 4
  1056: d7839232                 vrgather.vv v7, v9, v5
  1060: 5704a332                 vrgather.vv v8, v10, v6
  1064: d783832e                 vxor.vv v7, v8, v7
  1068: d780702e                 vxor.vv v1, v7, v1
  1072: 138f0200                 mv t5, t0
  1076: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1080: 138e0f00                 mv t3, t6
  1084: 87040e02                 vle8.v v9, (t3)
  1088: 138e0f01                 addi t3, t6, 0x10
  1092: 07050e02                 vle8.v v10, (t3)
  1096: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1100: d782b126                 vand.vv v5, v11, v3
  1104: 573332a2                 vsrl.vi v6, v3, 4
  1108: d7839232                 vrgather.vv v7, v9, v5
  1112: 5704a332                 vrgather.vv v8, v10, v6
  1116: d783832e                 vxor.vv v7, v8, v7
  1120: 5701712e                 vxor.vv v2, v7, v2
  1124: 83330502                 ld t2, 0x20(a0)
  1128: b3836300                 add t2, t2, t1
  1132: 87810302                 vle8.v v3, (t2)
  1136: 5700302e                 vxor.vv v0, v3, v0
  1140: 138f0200                 mv t5, t0
  1144: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1148: 138e0f06                 addi t3, t6, 0x60
  1152: 87040e02                 vle8.v v9, (t3)
  1156: 138e0f07                 addi t3, t6, 0x70
  1160: 07050e02                 vle8.v v10, (t3)
  1164: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1168: d782b126                 vand.vv v5, v11, v3
  1172: 573332a2                 vsrl.vi v6, v3, 4
  1176: d7839232                 vrgather.vv v7, v9, v5
  1180: 5704a332                 vrgather.vv v8, v10, v6
  1184: d783832e                 vxor.vv v7, v8, v7
  1188: d780702e                 vxor.vv v1, v7, v1
  1192: 138f0200                 mv t5, t0
  1196: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1200: 138e0f08                 addi t3, t6, 0x80
  1204: 87040e02                 vle8.v v9, (t3)
  1208: 138e0f09                 addi t3, t6, 0x90
  1212: 07050e02                 vle8.v v10, (t3)
  1216: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1220: d782b126                 vand.vv v5, v11, v3
  1224: 573332a2                 vsrl.vi v6, v3, 4
  1228: d7839232                 vrgather.vv v7, v9, v5
  1232: 5704a332                 vrgather.vv v8, v10, v6
  1236: d783832e                 vxor.vv v7, v8, v7
  1240: 5701712e                 vxor.vv v2, v7, v2
  1244: 83338502                 ld t2, 0x28(a0)
  1248: b3836300                 add t2, t2, t1
  1252: 87810302                 vle8.v v3, (t2)
  1256: 5700302e                 vxor.vv v0, v3, v0
  1260: d780302e                 vxor.vv v1, v3, v1
  1264: 5701312e                 vxor.vv v2, v3, v2
  1268: 83b30500                 ld t2, 0(a1)
  1272: b3836300                 add t2, t2, t1
  1276: 27800302                 vse8.v v0, (t2)
  1280: 83b38500                 ld t2, 8(a1)
  1284: b3836300                 add t2, t2, t1
  1288: a7800302                 vse8.v v1, (t2)
  1292: 83b30501                 ld t2, 0x10(a1)
  1296: b3836300                 add t2, t2, t1
  1300: 27810302                 vse8.v v2, (t2)
  1304: 33035300                 add t1, t1, t0
  1308: 33065640                 sub a2, a2, t0
  1312: 6ff09fd5                 j 0x278
  1316: 67800000                 ret 
  1320: 0000                     unimp 
  1322: 0000                     unimp 
  1324: 0000                     unimp 
  1326: 0000                     unimp 
