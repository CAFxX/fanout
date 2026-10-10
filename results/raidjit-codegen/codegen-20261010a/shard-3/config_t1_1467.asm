; config_t1_1467 target=rvv twin=pipe topo=raid6(8)/enc op=encode U=1
; code_size=1088 twins=["pipe"] uops_per_byte=12.5625
     0: 63140600                 bnez a2, 8
     4: 6f00c042                 j 0x430
     8: 970f0000                 auipc t6, 0
    12: 938f8f43                 addi t6, t6, 0x438
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 57450e5e                 vmv.v.x v10, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00001e                 j 0x210
    52: 83330500                 ld t2, 0(a0)
    56: b3836300                 add t2, t2, t1
    60: 87810302                 vle8.v v3, (t2)
    64: 5780015e                 vmv.v.v v0, v3
    68: 138e0f00                 mv t3, t6
    72: 07040e02                 vle8.v v8, (t3)
    76: 138e0f01                 addi t3, t6, 0x10
    80: 87040e02                 vle8.v v9, (t3)
    84: 5782a126                 vand.vv v4, v10, v3
    88: d73232a2                 vsrl.vi v5, v3, 4
    92: 57038232                 vrgather.vv v6, v8, v4
    96: d7839232                 vrgather.vv v7, v9, v5
   100: 5703732e                 vxor.vv v6, v7, v6
   104: d700035e                 vmv.v.v v1, v6
   108: 83338500                 ld t2, 8(a0)
   112: b3836300                 add t2, t2, t1
   116: 07810302                 vle8.v v2, (t2)
   120: 5700202e                 vxor.vv v0, v2, v0
   124: 138e0f02                 addi t3, t6, 0x20
   128: 07040e02                 vle8.v v8, (t3)
   132: 138e0f03                 addi t3, t6, 0x30
   136: 87040e02                 vle8.v v9, (t3)
   140: 5702a126                 vand.vv v4, v10, v2
   144: d73222a2                 vsrl.vi v5, v2, 4
   148: 57038232                 vrgather.vv v6, v8, v4
   152: d7839232                 vrgather.vv v7, v9, v5
   156: 5703732e                 vxor.vv v6, v7, v6
   160: d780602e                 vxor.vv v1, v6, v1
   164: 83330501                 ld t2, 0x10(a0)
   168: b3836300                 add t2, t2, t1
   172: 07810302                 vle8.v v2, (t2)
   176: 5700202e                 vxor.vv v0, v2, v0
   180: 138e0f04                 addi t3, t6, 0x40
   184: 07040e02                 vle8.v v8, (t3)
   188: 138e0f05                 addi t3, t6, 0x50
   192: 87040e02                 vle8.v v9, (t3)
   196: 5702a126                 vand.vv v4, v10, v2
   200: d73222a2                 vsrl.vi v5, v2, 4
   204: 57038232                 vrgather.vv v6, v8, v4
   208: d7839232                 vrgather.vv v7, v9, v5
   212: 5703732e                 vxor.vv v6, v7, v6
   216: d780602e                 vxor.vv v1, v6, v1
   220: 83338501                 ld t2, 0x18(a0)
   224: b3836300                 add t2, t2, t1
   228: 07810302                 vle8.v v2, (t2)
   232: 5700202e                 vxor.vv v0, v2, v0
   236: 138e0f06                 addi t3, t6, 0x60
   240: 07040e02                 vle8.v v8, (t3)
   244: 138e0f07                 addi t3, t6, 0x70
   248: 87040e02                 vle8.v v9, (t3)
   252: 5702a126                 vand.vv v4, v10, v2
   256: d73222a2                 vsrl.vi v5, v2, 4
   260: 57038232                 vrgather.vv v6, v8, v4
   264: d7839232                 vrgather.vv v7, v9, v5
   268: 5703732e                 vxor.vv v6, v7, v6
   272: d780602e                 vxor.vv v1, v6, v1
   276: 83330502                 ld t2, 0x20(a0)
   280: b3836300                 add t2, t2, t1
   284: 07810302                 vle8.v v2, (t2)
   288: 5700202e                 vxor.vv v0, v2, v0
   292: 138e0f08                 addi t3, t6, 0x80
   296: 07040e02                 vle8.v v8, (t3)
   300: 138e0f09                 addi t3, t6, 0x90
   304: 87040e02                 vle8.v v9, (t3)
   308: 5702a126                 vand.vv v4, v10, v2
   312: d73222a2                 vsrl.vi v5, v2, 4
   316: 57038232                 vrgather.vv v6, v8, v4
   320: d7839232                 vrgather.vv v7, v9, v5
   324: 5703732e                 vxor.vv v6, v7, v6
   328: d780602e                 vxor.vv v1, v6, v1
   332: 83338502                 ld t2, 0x28(a0)
   336: b3836300                 add t2, t2, t1
   340: 07810302                 vle8.v v2, (t2)
   344: 5700202e                 vxor.vv v0, v2, v0
   348: 138e0f0a                 addi t3, t6, 0xa0
   352: 07040e02                 vle8.v v8, (t3)
   356: 138e0f0b                 addi t3, t6, 0xb0
   360: 87040e02                 vle8.v v9, (t3)
   364: 5702a126                 vand.vv v4, v10, v2
   368: d73222a2                 vsrl.vi v5, v2, 4
   372: 57038232                 vrgather.vv v6, v8, v4
   376: d7839232                 vrgather.vv v7, v9, v5
   380: 5703732e                 vxor.vv v6, v7, v6
   384: d780602e                 vxor.vv v1, v6, v1
   388: 83330503                 ld t2, 0x30(a0)
   392: b3836300                 add t2, t2, t1
   396: 07810302                 vle8.v v2, (t2)
   400: 5700202e                 vxor.vv v0, v2, v0
   404: 138e0f0c                 addi t3, t6, 0xc0
   408: 07040e02                 vle8.v v8, (t3)
   412: 138e0f0d                 addi t3, t6, 0xd0
   416: 87040e02                 vle8.v v9, (t3)
   420: 5702a126                 vand.vv v4, v10, v2
   424: d73222a2                 vsrl.vi v5, v2, 4
   428: 57038232                 vrgather.vv v6, v8, v4
   432: d7839232                 vrgather.vv v7, v9, v5
   436: 5703732e                 vxor.vv v6, v7, v6
   440: d780602e                 vxor.vv v1, v6, v1
   444: 83338503                 ld t2, 0x38(a0)
   448: b3836300                 add t2, t2, t1
   452: 07810302                 vle8.v v2, (t2)
   456: 5700202e                 vxor.vv v0, v2, v0
   460: d780202e                 vxor.vv v1, v2, v1
   464: 130ff7ff                 addi t5, a4, -1
   468: 630a0f00                 beqz t5, 0x1e8
   472: 330fd300                 add t5, t1, a3
   476: 83330500                 ld t2, 0(a0)
   480: b383e301                 add t2, t2, t5
   484: 87810302                 vle8.v v3, (t2)
   488: 83b30500                 ld t2, 0(a1)
   492: b3836300                 add t2, t2, t1
   496: 27800302                 vse8.v v0, (t2)
   500: 83b38500                 ld t2, 8(a1)
   504: b3836300                 add t2, t2, t1
   508: a7800302                 vse8.v v1, (t2)
   512: 3303d300                 add t1, t1, a3
   516: 1307f7ff                 addi a4, a4, -1
   520: e31c07e2                 bnez a4, 0x40
   524: 33066640                 sub a2, a2, t1
   528: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   532: 63940200                 bnez t0, 0x21c
   536: 6f008021                 j 0x430
   540: 83330500                 ld t2, 0(a0)
   544: b3836300                 add t2, t2, t1
   548: 07810302                 vle8.v v2, (t2)
   552: 5700015e                 vmv.v.v v0, v2
   556: 138f0200                 mv t5, t0
   560: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   564: 138e0f00                 mv t3, t6
   568: 07040e02                 vle8.v v8, (t3)
   572: 138e0f01                 addi t3, t6, 0x10
   576: 87040e02                 vle8.v v9, (t3)
   580: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   584: 5702a126                 vand.vv v4, v10, v2
   588: d73222a2                 vsrl.vi v5, v2, 4
   592: 57038232                 vrgather.vv v6, v8, v4
   596: d7839232                 vrgather.vv v7, v9, v5
   600: 5703732e                 vxor.vv v6, v7, v6
   604: d700035e                 vmv.v.v v1, v6
   608: 83338500                 ld t2, 8(a0)
   612: b3836300                 add t2, t2, t1
   616: 07810302                 vle8.v v2, (t2)
   620: 5700202e                 vxor.vv v0, v2, v0
   624: 138f0200                 mv t5, t0
   628: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   632: 138e0f02                 addi t3, t6, 0x20
   636: 07040e02                 vle8.v v8, (t3)
   640: 138e0f03                 addi t3, t6, 0x30
   644: 87040e02                 vle8.v v9, (t3)
   648: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   652: 5702a126                 vand.vv v4, v10, v2
   656: d73222a2                 vsrl.vi v5, v2, 4
   660: 57038232                 vrgather.vv v6, v8, v4
   664: d7839232                 vrgather.vv v7, v9, v5
   668: 5703732e                 vxor.vv v6, v7, v6
   672: d780602e                 vxor.vv v1, v6, v1
   676: 83330501                 ld t2, 0x10(a0)
   680: b3836300                 add t2, t2, t1
   684: 07810302                 vle8.v v2, (t2)
   688: 5700202e                 vxor.vv v0, v2, v0
   692: 138f0200                 mv t5, t0
   696: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   700: 138e0f04                 addi t3, t6, 0x40
   704: 07040e02                 vle8.v v8, (t3)
   708: 138e0f05                 addi t3, t6, 0x50
   712: 87040e02                 vle8.v v9, (t3)
   716: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   720: 5702a126                 vand.vv v4, v10, v2
   724: d73222a2                 vsrl.vi v5, v2, 4
   728: 57038232                 vrgather.vv v6, v8, v4
   732: d7839232                 vrgather.vv v7, v9, v5
   736: 5703732e                 vxor.vv v6, v7, v6
   740: d780602e                 vxor.vv v1, v6, v1
   744: 83338501                 ld t2, 0x18(a0)
   748: b3836300                 add t2, t2, t1
   752: 07810302                 vle8.v v2, (t2)
   756: 5700202e                 vxor.vv v0, v2, v0
   760: 138f0200                 mv t5, t0
   764: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   768: 138e0f06                 addi t3, t6, 0x60
   772: 07040e02                 vle8.v v8, (t3)
   776: 138e0f07                 addi t3, t6, 0x70
   780: 87040e02                 vle8.v v9, (t3)
   784: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   788: 5702a126                 vand.vv v4, v10, v2
   792: d73222a2                 vsrl.vi v5, v2, 4
   796: 57038232                 vrgather.vv v6, v8, v4
   800: d7839232                 vrgather.vv v7, v9, v5
   804: 5703732e                 vxor.vv v6, v7, v6
   808: d780602e                 vxor.vv v1, v6, v1
   812: 83330502                 ld t2, 0x20(a0)
   816: b3836300                 add t2, t2, t1
   820: 07810302                 vle8.v v2, (t2)
   824: 5700202e                 vxor.vv v0, v2, v0
   828: 138f0200                 mv t5, t0
   832: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   836: 138e0f08                 addi t3, t6, 0x80
   840: 07040e02                 vle8.v v8, (t3)
   844: 138e0f09                 addi t3, t6, 0x90
   848: 87040e02                 vle8.v v9, (t3)
   852: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   856: 5702a126                 vand.vv v4, v10, v2
   860: d73222a2                 vsrl.vi v5, v2, 4
   864: 57038232                 vrgather.vv v6, v8, v4
   868: d7839232                 vrgather.vv v7, v9, v5
   872: 5703732e                 vxor.vv v6, v7, v6
   876: d780602e                 vxor.vv v1, v6, v1
   880: 83338502                 ld t2, 0x28(a0)
   884: b3836300                 add t2, t2, t1
   888: 07810302                 vle8.v v2, (t2)
   892: 5700202e                 vxor.vv v0, v2, v0
   896: 138f0200                 mv t5, t0
   900: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   904: 138e0f0a                 addi t3, t6, 0xa0
   908: 07040e02                 vle8.v v8, (t3)
   912: 138e0f0b                 addi t3, t6, 0xb0
   916: 87040e02                 vle8.v v9, (t3)
   920: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   924: 5702a126                 vand.vv v4, v10, v2
   928: d73222a2                 vsrl.vi v5, v2, 4
   932: 57038232                 vrgather.vv v6, v8, v4
   936: d7839232                 vrgather.vv v7, v9, v5
   940: 5703732e                 vxor.vv v6, v7, v6
   944: d780602e                 vxor.vv v1, v6, v1
   948: 83330503                 ld t2, 0x30(a0)
   952: b3836300                 add t2, t2, t1
   956: 07810302                 vle8.v v2, (t2)
   960: 5700202e                 vxor.vv v0, v2, v0
   964: 138f0200                 mv t5, t0
   968: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   972: 138e0f0c                 addi t3, t6, 0xc0
   976: 07040e02                 vle8.v v8, (t3)
   980: 138e0f0d                 addi t3, t6, 0xd0
   984: 87040e02                 vle8.v v9, (t3)
   988: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   992: 5702a126                 vand.vv v4, v10, v2
   996: d73222a2                 vsrl.vi v5, v2, 4
  1000: 57038232                 vrgather.vv v6, v8, v4
  1004: d7839232                 vrgather.vv v7, v9, v5
  1008: 5703732e                 vxor.vv v6, v7, v6
  1012: d780602e                 vxor.vv v1, v6, v1
  1016: 83338503                 ld t2, 0x38(a0)
  1020: b3836300                 add t2, t2, t1
  1024: 07810302                 vle8.v v2, (t2)
  1028: 5700202e                 vxor.vv v0, v2, v0
  1032: d780202e                 vxor.vv v1, v2, v1
  1036: 83b30500                 ld t2, 0(a1)
  1040: b3836300                 add t2, t2, t1
  1044: 27800302                 vse8.v v0, (t2)
  1048: 83b38500                 ld t2, 8(a1)
  1052: b3836300                 add t2, t2, t1
  1056: a7800302                 vse8.v v1, (t2)
  1060: 33035300                 add t1, t1, t0
  1064: 33065640                 sub a2, a2, t0
  1068: 6ff05fde                 j 0x210
  1072: 67800000                 ret 
  1076: 0000                     unimp 
  1078: 0000                     unimp 
  1080: 0000                     unimp 
  1082: 0000                     unimp 
  1084: 0000                     unimp 
  1086: 0000                     unimp 
