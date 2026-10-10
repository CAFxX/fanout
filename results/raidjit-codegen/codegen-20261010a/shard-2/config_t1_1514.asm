; config_t1_1514 target=rvv twin=maskedtail topo=rs(6,3)/enc op=encode U=1
; code_size=1216 twins=["maskedtail"] uops_per_byte=15.3125
     0: 63140600                 bnez a2, 8
     4: 6f00404b                 j 0x4b8
     8: 970f0000                 auipc t6, 0
    12: 938f8f4b                 addi t6, t6, 0x4b8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d7450e5e                 vmv.v.x v11, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00c023                 j 0x26c
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07820302                 vle8.v v4, (t2)
    76: d700025e                 vmv.v.v v1, v4
    80: 138e0f00                 mv t3, t6
    84: 87040e02                 vle8.v v9, (t3)
    88: 138e0f01                 addi t3, t6, 0x10
    92: 07050e02                 vle8.v v10, (t3)
    96: d702b226                 vand.vv v5, v11, v4
   100: 573342a2                 vsrl.vi v6, v4, 4
   104: d7839232                 vrgather.vv v7, v9, v5
   108: 5704a332                 vrgather.vv v8, v10, v6
   112: d783832e                 vxor.vv v7, v8, v7
   116: 5781035e                 vmv.v.v v2, v7
   120: 138e0f02                 addi t3, t6, 0x20
   124: 87040e02                 vle8.v v9, (t3)
   128: 138e0f03                 addi t3, t6, 0x30
   132: 07050e02                 vle8.v v10, (t3)
   136: d702b226                 vand.vv v5, v11, v4
   140: 573342a2                 vsrl.vi v6, v4, 4
   144: d7839232                 vrgather.vv v7, v9, v5
   148: 5704a332                 vrgather.vv v8, v10, v6
   152: d783832e                 vxor.vv v7, v8, v7
   156: d781035e                 vmv.v.v v3, v7
   160: 83338500                 ld t2, 8(a0)
   164: b3836300                 add t2, t2, t1
   168: 07820302                 vle8.v v4, (t2)
   172: d780402e                 vxor.vv v1, v4, v1
   176: 138e0f02                 addi t3, t6, 0x20
   180: 87040e02                 vle8.v v9, (t3)
   184: 138e0f03                 addi t3, t6, 0x30
   188: 07050e02                 vle8.v v10, (t3)
   192: d702b226                 vand.vv v5, v11, v4
   196: 573342a2                 vsrl.vi v6, v4, 4
   200: d7839232                 vrgather.vv v7, v9, v5
   204: 5704a332                 vrgather.vv v8, v10, v6
   208: d783832e                 vxor.vv v7, v8, v7
   212: 5701712e                 vxor.vv v2, v7, v2
   216: 138e0f04                 addi t3, t6, 0x40
   220: 87040e02                 vle8.v v9, (t3)
   224: 138e0f05                 addi t3, t6, 0x50
   228: 07050e02                 vle8.v v10, (t3)
   232: d702b226                 vand.vv v5, v11, v4
   236: 573342a2                 vsrl.vi v6, v4, 4
   240: d7839232                 vrgather.vv v7, v9, v5
   244: 5704a332                 vrgather.vv v8, v10, v6
   248: d783832e                 vxor.vv v7, v8, v7
   252: d781712e                 vxor.vv v3, v7, v3
   256: 83330501                 ld t2, 0x10(a0)
   260: b3836300                 add t2, t2, t1
   264: 07820302                 vle8.v v4, (t2)
   268: d780402e                 vxor.vv v1, v4, v1
   272: 138e0f04                 addi t3, t6, 0x40
   276: 87040e02                 vle8.v v9, (t3)
   280: 138e0f05                 addi t3, t6, 0x50
   284: 07050e02                 vle8.v v10, (t3)
   288: d702b226                 vand.vv v5, v11, v4
   292: 573342a2                 vsrl.vi v6, v4, 4
   296: d7839232                 vrgather.vv v7, v9, v5
   300: 5704a332                 vrgather.vv v8, v10, v6
   304: d783832e                 vxor.vv v7, v8, v7
   308: 5701712e                 vxor.vv v2, v7, v2
   312: 138e0f06                 addi t3, t6, 0x60
   316: 87040e02                 vle8.v v9, (t3)
   320: 138e0f07                 addi t3, t6, 0x70
   324: 07050e02                 vle8.v v10, (t3)
   328: d702b226                 vand.vv v5, v11, v4
   332: 573342a2                 vsrl.vi v6, v4, 4
   336: d7839232                 vrgather.vv v7, v9, v5
   340: 5704a332                 vrgather.vv v8, v10, v6
   344: d783832e                 vxor.vv v7, v8, v7
   348: d781712e                 vxor.vv v3, v7, v3
   352: 83338501                 ld t2, 0x18(a0)
   356: b3836300                 add t2, t2, t1
   360: 07820302                 vle8.v v4, (t2)
   364: d780402e                 vxor.vv v1, v4, v1
   368: 138e0f08                 addi t3, t6, 0x80
   372: 87040e02                 vle8.v v9, (t3)
   376: 138e0f09                 addi t3, t6, 0x90
   380: 07050e02                 vle8.v v10, (t3)
   384: d702b226                 vand.vv v5, v11, v4
   388: 573342a2                 vsrl.vi v6, v4, 4
   392: d7839232                 vrgather.vv v7, v9, v5
   396: 5704a332                 vrgather.vv v8, v10, v6
   400: d783832e                 vxor.vv v7, v8, v7
   404: 5701712e                 vxor.vv v2, v7, v2
   408: 138e0f00                 mv t3, t6
   412: 87040e02                 vle8.v v9, (t3)
   416: 138e0f01                 addi t3, t6, 0x10
   420: 07050e02                 vle8.v v10, (t3)
   424: d702b226                 vand.vv v5, v11, v4
   428: 573342a2                 vsrl.vi v6, v4, 4
   432: d7839232                 vrgather.vv v7, v9, v5
   436: 5704a332                 vrgather.vv v8, v10, v6
   440: d783832e                 vxor.vv v7, v8, v7
   444: d781712e                 vxor.vv v3, v7, v3
   448: 83330502                 ld t2, 0x20(a0)
   452: b3836300                 add t2, t2, t1
   456: 07820302                 vle8.v v4, (t2)
   460: d780402e                 vxor.vv v1, v4, v1
   464: 138e0f06                 addi t3, t6, 0x60
   468: 87040e02                 vle8.v v9, (t3)
   472: 138e0f07                 addi t3, t6, 0x70
   476: 07050e02                 vle8.v v10, (t3)
   480: d702b226                 vand.vv v5, v11, v4
   484: 573342a2                 vsrl.vi v6, v4, 4
   488: d7839232                 vrgather.vv v7, v9, v5
   492: 5704a332                 vrgather.vv v8, v10, v6
   496: d783832e                 vxor.vv v7, v8, v7
   500: 5701712e                 vxor.vv v2, v7, v2
   504: 138e0f08                 addi t3, t6, 0x80
   508: 87040e02                 vle8.v v9, (t3)
   512: 138e0f09                 addi t3, t6, 0x90
   516: 07050e02                 vle8.v v10, (t3)
   520: d702b226                 vand.vv v5, v11, v4
   524: 573342a2                 vsrl.vi v6, v4, 4
   528: d7839232                 vrgather.vv v7, v9, v5
   532: 5704a332                 vrgather.vv v8, v10, v6
   536: d783832e                 vxor.vv v7, v8, v7
   540: d781712e                 vxor.vv v3, v7, v3
   544: 83338502                 ld t2, 0x28(a0)
   548: b3836300                 add t2, t2, t1
   552: 07820302                 vle8.v v4, (t2)
   556: d780402e                 vxor.vv v1, v4, v1
   560: 5701412e                 vxor.vv v2, v4, v2
   564: d781412e                 vxor.vv v3, v4, v3
   568: 83b30500                 ld t2, 0(a1)
   572: b3836300                 add t2, t2, t1
   576: a7800302                 vse8.v v1, (t2)
   580: 83b38500                 ld t2, 8(a1)
   584: b3836300                 add t2, t2, t1
   588: 27810302                 vse8.v v2, (t2)
   592: 83b30501                 ld t2, 0x10(a1)
   596: b3836300                 add t2, t2, t1
   600: a7810302                 vse8.v v3, (t2)
   604: 3303d300                 add t1, t1, a3
   608: 1307f7ff                 addi a4, a4, -1
   612: e31e07dc                 bnez a4, 0x40
   616: 33066640                 sub a2, a2, t1
   620: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
   624: 93060600                 mv a3, a2
   628: 6344d000                 bgtz a3, 0x27c
   632: 6f008023                 j 0x4b0
   636: b3835640                 sub t2, a3, t0
   640: 13def343                 srai t3, t2, 0x3f
   644: b3f3c301                 and t2, t2, t3
   648: b3837200                 add t2, t0, t2
   652: 57a20852                 vid.v v4
   656: 57c0436a                 vmsltu.vx v0, v4, t2
   660: 83330500                 ld t2, 0(a0)
   664: b3836300                 add t2, t2, t1
   668: 07820300                 vle8.v v4, (t2), v0.t
   672: d700025e                 vmv.v.v v1, v4
   676: 138e0f00                 mv t3, t6
   680: 87040e02                 vle8.v v9, (t3)
   684: 138e0f01                 addi t3, t6, 0x10
   688: 07050e02                 vle8.v v10, (t3)
   692: d702b226                 vand.vv v5, v11, v4
   696: 573342a2                 vsrl.vi v6, v4, 4
   700: d7839232                 vrgather.vv v7, v9, v5
   704: 5704a332                 vrgather.vv v8, v10, v6
   708: d783832e                 vxor.vv v7, v8, v7
   712: 5781035e                 vmv.v.v v2, v7
   716: 138e0f02                 addi t3, t6, 0x20
   720: 87040e02                 vle8.v v9, (t3)
   724: 138e0f03                 addi t3, t6, 0x30
   728: 07050e02                 vle8.v v10, (t3)
   732: d702b226                 vand.vv v5, v11, v4
   736: 573342a2                 vsrl.vi v6, v4, 4
   740: d7839232                 vrgather.vv v7, v9, v5
   744: 5704a332                 vrgather.vv v8, v10, v6
   748: d783832e                 vxor.vv v7, v8, v7
   752: d781035e                 vmv.v.v v3, v7
   756: 83338500                 ld t2, 8(a0)
   760: b3836300                 add t2, t2, t1
   764: 07820300                 vle8.v v4, (t2), v0.t
   768: d780402e                 vxor.vv v1, v4, v1
   772: 138e0f02                 addi t3, t6, 0x20
   776: 87040e02                 vle8.v v9, (t3)
   780: 138e0f03                 addi t3, t6, 0x30
   784: 07050e02                 vle8.v v10, (t3)
   788: d702b226                 vand.vv v5, v11, v4
   792: 573342a2                 vsrl.vi v6, v4, 4
   796: d7839232                 vrgather.vv v7, v9, v5
   800: 5704a332                 vrgather.vv v8, v10, v6
   804: d783832e                 vxor.vv v7, v8, v7
   808: 5701712e                 vxor.vv v2, v7, v2
   812: 138e0f04                 addi t3, t6, 0x40
   816: 87040e02                 vle8.v v9, (t3)
   820: 138e0f05                 addi t3, t6, 0x50
   824: 07050e02                 vle8.v v10, (t3)
   828: d702b226                 vand.vv v5, v11, v4
   832: 573342a2                 vsrl.vi v6, v4, 4
   836: d7839232                 vrgather.vv v7, v9, v5
   840: 5704a332                 vrgather.vv v8, v10, v6
   844: d783832e                 vxor.vv v7, v8, v7
   848: d781712e                 vxor.vv v3, v7, v3
   852: 83330501                 ld t2, 0x10(a0)
   856: b3836300                 add t2, t2, t1
   860: 07820300                 vle8.v v4, (t2), v0.t
   864: d780402e                 vxor.vv v1, v4, v1
   868: 138e0f04                 addi t3, t6, 0x40
   872: 87040e02                 vle8.v v9, (t3)
   876: 138e0f05                 addi t3, t6, 0x50
   880: 07050e02                 vle8.v v10, (t3)
   884: d702b226                 vand.vv v5, v11, v4
   888: 573342a2                 vsrl.vi v6, v4, 4
   892: d7839232                 vrgather.vv v7, v9, v5
   896: 5704a332                 vrgather.vv v8, v10, v6
   900: d783832e                 vxor.vv v7, v8, v7
   904: 5701712e                 vxor.vv v2, v7, v2
   908: 138e0f06                 addi t3, t6, 0x60
   912: 87040e02                 vle8.v v9, (t3)
   916: 138e0f07                 addi t3, t6, 0x70
   920: 07050e02                 vle8.v v10, (t3)
   924: d702b226                 vand.vv v5, v11, v4
   928: 573342a2                 vsrl.vi v6, v4, 4
   932: d7839232                 vrgather.vv v7, v9, v5
   936: 5704a332                 vrgather.vv v8, v10, v6
   940: d783832e                 vxor.vv v7, v8, v7
   944: d781712e                 vxor.vv v3, v7, v3
   948: 83338501                 ld t2, 0x18(a0)
   952: b3836300                 add t2, t2, t1
   956: 07820300                 vle8.v v4, (t2), v0.t
   960: d780402e                 vxor.vv v1, v4, v1
   964: 138e0f08                 addi t3, t6, 0x80
   968: 87040e02                 vle8.v v9, (t3)
   972: 138e0f09                 addi t3, t6, 0x90
   976: 07050e02                 vle8.v v10, (t3)
   980: d702b226                 vand.vv v5, v11, v4
   984: 573342a2                 vsrl.vi v6, v4, 4
   988: d7839232                 vrgather.vv v7, v9, v5
   992: 5704a332                 vrgather.vv v8, v10, v6
   996: d783832e                 vxor.vv v7, v8, v7
  1000: 5701712e                 vxor.vv v2, v7, v2
  1004: 138e0f00                 mv t3, t6
  1008: 87040e02                 vle8.v v9, (t3)
  1012: 138e0f01                 addi t3, t6, 0x10
  1016: 07050e02                 vle8.v v10, (t3)
  1020: d702b226                 vand.vv v5, v11, v4
  1024: 573342a2                 vsrl.vi v6, v4, 4
  1028: d7839232                 vrgather.vv v7, v9, v5
  1032: 5704a332                 vrgather.vv v8, v10, v6
  1036: d783832e                 vxor.vv v7, v8, v7
  1040: d781712e                 vxor.vv v3, v7, v3
  1044: 83330502                 ld t2, 0x20(a0)
  1048: b3836300                 add t2, t2, t1
  1052: 07820300                 vle8.v v4, (t2), v0.t
  1056: d780402e                 vxor.vv v1, v4, v1
  1060: 138e0f06                 addi t3, t6, 0x60
  1064: 87040e02                 vle8.v v9, (t3)
  1068: 138e0f07                 addi t3, t6, 0x70
  1072: 07050e02                 vle8.v v10, (t3)
  1076: d702b226                 vand.vv v5, v11, v4
  1080: 573342a2                 vsrl.vi v6, v4, 4
  1084: d7839232                 vrgather.vv v7, v9, v5
  1088: 5704a332                 vrgather.vv v8, v10, v6
  1092: d783832e                 vxor.vv v7, v8, v7
  1096: 5701712e                 vxor.vv v2, v7, v2
  1100: 138e0f08                 addi t3, t6, 0x80
  1104: 87040e02                 vle8.v v9, (t3)
  1108: 138e0f09                 addi t3, t6, 0x90
  1112: 07050e02                 vle8.v v10, (t3)
  1116: d702b226                 vand.vv v5, v11, v4
  1120: 573342a2                 vsrl.vi v6, v4, 4
  1124: d7839232                 vrgather.vv v7, v9, v5
  1128: 5704a332                 vrgather.vv v8, v10, v6
  1132: d783832e                 vxor.vv v7, v8, v7
  1136: d781712e                 vxor.vv v3, v7, v3
  1140: 83338502                 ld t2, 0x28(a0)
  1144: b3836300                 add t2, t2, t1
  1148: 07820300                 vle8.v v4, (t2), v0.t
  1152: d780402e                 vxor.vv v1, v4, v1
  1156: 5701412e                 vxor.vv v2, v4, v2
  1160: d781412e                 vxor.vv v3, v4, v3
  1164: 83b30500                 ld t2, 0(a1)
  1168: b3836300                 add t2, t2, t1
  1172: a7800300                 vse8.v v1, (t2), v0.t
  1176: 83b38500                 ld t2, 8(a1)
  1180: b3836300                 add t2, t2, t1
  1184: 27810300                 vse8.v v2, (t2), v0.t
  1188: 83b30501                 ld t2, 0x10(a1)
  1192: b3836300                 add t2, t2, t1
  1196: a7810300                 vse8.v v3, (t2), v0.t
  1200: 33035300                 add t1, t1, t0
  1204: b3865640                 sub a3, a3, t0
  1208: 67800000                 ret 
  1212: 0000                     unimp 
  1214: 0000                     unimp 
