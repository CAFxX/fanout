; config_t1_1743 target=rvv-zvbc32e twin=nt topo=raid6(8)/enc op=encode U=1
; code_size=1088 twins=["nt"] uops_per_byte=12.5625
     0: 63140600                 bnez a2, 8
     4: 6f000043                 j 0x434
     8: 970f0000                 auipc t6, 0
    12: 938f8f43                 addi t6, t6, 0x438
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d7440e5e                 vmv.v.x v9, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00c01d                 j 0x20c
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07810302                 vle8.v v2, (t2)
    76: 5700015e                 vmv.v.v v0, v2
    80: 138e0f00                 mv t3, t6
    84: 87030e02                 vle8.v v7, (t3)
    88: 138e0f01                 addi t3, t6, 0x10
    92: 07040e02                 vle8.v v8, (t3)
    96: d7019126                 vand.vv v3, v9, v2
   100: 573222a2                 vsrl.vi v4, v2, 4
   104: d7827132                 vrgather.vv v5, v7, v3
   108: 57038232                 vrgather.vv v6, v8, v4
   112: d782622e                 vxor.vv v5, v6, v5
   116: d780025e                 vmv.v.v v1, v5
   120: 83338500                 ld t2, 8(a0)
   124: b3836300                 add t2, t2, t1
   128: 07810302                 vle8.v v2, (t2)
   132: 5700202e                 vxor.vv v0, v2, v0
   136: 138e0f02                 addi t3, t6, 0x20
   140: 87030e02                 vle8.v v7, (t3)
   144: 138e0f03                 addi t3, t6, 0x30
   148: 07040e02                 vle8.v v8, (t3)
   152: d7019126                 vand.vv v3, v9, v2
   156: 573222a2                 vsrl.vi v4, v2, 4
   160: d7827132                 vrgather.vv v5, v7, v3
   164: 57038232                 vrgather.vv v6, v8, v4
   168: d782622e                 vxor.vv v5, v6, v5
   172: d780502e                 vxor.vv v1, v5, v1
   176: 83330501                 ld t2, 0x10(a0)
   180: b3836300                 add t2, t2, t1
   184: 07810302                 vle8.v v2, (t2)
   188: 5700202e                 vxor.vv v0, v2, v0
   192: 138e0f04                 addi t3, t6, 0x40
   196: 87030e02                 vle8.v v7, (t3)
   200: 138e0f05                 addi t3, t6, 0x50
   204: 07040e02                 vle8.v v8, (t3)
   208: d7019126                 vand.vv v3, v9, v2
   212: 573222a2                 vsrl.vi v4, v2, 4
   216: d7827132                 vrgather.vv v5, v7, v3
   220: 57038232                 vrgather.vv v6, v8, v4
   224: d782622e                 vxor.vv v5, v6, v5
   228: d780502e                 vxor.vv v1, v5, v1
   232: 83338501                 ld t2, 0x18(a0)
   236: b3836300                 add t2, t2, t1
   240: 07810302                 vle8.v v2, (t2)
   244: 5700202e                 vxor.vv v0, v2, v0
   248: 138e0f06                 addi t3, t6, 0x60
   252: 87030e02                 vle8.v v7, (t3)
   256: 138e0f07                 addi t3, t6, 0x70
   260: 07040e02                 vle8.v v8, (t3)
   264: d7019126                 vand.vv v3, v9, v2
   268: 573222a2                 vsrl.vi v4, v2, 4
   272: d7827132                 vrgather.vv v5, v7, v3
   276: 57038232                 vrgather.vv v6, v8, v4
   280: d782622e                 vxor.vv v5, v6, v5
   284: d780502e                 vxor.vv v1, v5, v1
   288: 83330502                 ld t2, 0x20(a0)
   292: b3836300                 add t2, t2, t1
   296: 07810302                 vle8.v v2, (t2)
   300: 5700202e                 vxor.vv v0, v2, v0
   304: 138e0f08                 addi t3, t6, 0x80
   308: 87030e02                 vle8.v v7, (t3)
   312: 138e0f09                 addi t3, t6, 0x90
   316: 07040e02                 vle8.v v8, (t3)
   320: d7019126                 vand.vv v3, v9, v2
   324: 573222a2                 vsrl.vi v4, v2, 4
   328: d7827132                 vrgather.vv v5, v7, v3
   332: 57038232                 vrgather.vv v6, v8, v4
   336: d782622e                 vxor.vv v5, v6, v5
   340: d780502e                 vxor.vv v1, v5, v1
   344: 83338502                 ld t2, 0x28(a0)
   348: b3836300                 add t2, t2, t1
   352: 07810302                 vle8.v v2, (t2)
   356: 5700202e                 vxor.vv v0, v2, v0
   360: 138e0f0a                 addi t3, t6, 0xa0
   364: 87030e02                 vle8.v v7, (t3)
   368: 138e0f0b                 addi t3, t6, 0xb0
   372: 07040e02                 vle8.v v8, (t3)
   376: d7019126                 vand.vv v3, v9, v2
   380: 573222a2                 vsrl.vi v4, v2, 4
   384: d7827132                 vrgather.vv v5, v7, v3
   388: 57038232                 vrgather.vv v6, v8, v4
   392: d782622e                 vxor.vv v5, v6, v5
   396: d780502e                 vxor.vv v1, v5, v1
   400: 83330503                 ld t2, 0x30(a0)
   404: b3836300                 add t2, t2, t1
   408: 07810302                 vle8.v v2, (t2)
   412: 5700202e                 vxor.vv v0, v2, v0
   416: 138e0f0c                 addi t3, t6, 0xc0
   420: 87030e02                 vle8.v v7, (t3)
   424: 138e0f0d                 addi t3, t6, 0xd0
   428: 07040e02                 vle8.v v8, (t3)
   432: d7019126                 vand.vv v3, v9, v2
   436: 573222a2                 vsrl.vi v4, v2, 4
   440: d7827132                 vrgather.vv v5, v7, v3
   444: 57038232                 vrgather.vv v6, v8, v4
   448: d782622e                 vxor.vv v5, v6, v5
   452: d780502e                 vxor.vv v1, v5, v1
   456: 83338503                 ld t2, 0x38(a0)
   460: b3836300                 add t2, t2, t1
   464: 07810302                 vle8.v v2, (t2)
   468: 5700202e                 vxor.vv v0, v2, v0
   472: d780202e                 vxor.vv v1, v2, v1
   476: 83b30500                 ld t2, 0(a1)
   480: b3836300                 add t2, t2, t1
   484: 33005000                 ntl.all 
   488: 27800302                 vse8.v v0, (t2)
   492: 83b38500                 ld t2, 8(a1)
   496: b3836300                 add t2, t2, t1
   500: 33005000                 ntl.all 
   504: a7800302                 vse8.v v1, (t2)
   508: 3303d300                 add t1, t1, a3
   512: 1307f7ff                 addi a4, a4, -1
   516: e31e07e2                 bnez a4, 0x40
   520: 33066640                 sub a2, a2, t1
   524: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   528: 63940200                 bnez t0, 0x218
   532: 6f000022                 j 0x434
   536: 83330500                 ld t2, 0(a0)
   540: b3836300                 add t2, t2, t1
   544: 07810302                 vle8.v v2, (t2)
   548: 5700015e                 vmv.v.v v0, v2
   552: 138f0200                 mv t5, t0
   556: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   560: 138e0f00                 mv t3, t6
   564: 87030e02                 vle8.v v7, (t3)
   568: 138e0f01                 addi t3, t6, 0x10
   572: 07040e02                 vle8.v v8, (t3)
   576: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   580: d7019126                 vand.vv v3, v9, v2
   584: 573222a2                 vsrl.vi v4, v2, 4
   588: d7827132                 vrgather.vv v5, v7, v3
   592: 57038232                 vrgather.vv v6, v8, v4
   596: d782622e                 vxor.vv v5, v6, v5
   600: d780025e                 vmv.v.v v1, v5
   604: 83338500                 ld t2, 8(a0)
   608: b3836300                 add t2, t2, t1
   612: 07810302                 vle8.v v2, (t2)
   616: 5700202e                 vxor.vv v0, v2, v0
   620: 138f0200                 mv t5, t0
   624: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   628: 138e0f02                 addi t3, t6, 0x20
   632: 87030e02                 vle8.v v7, (t3)
   636: 138e0f03                 addi t3, t6, 0x30
   640: 07040e02                 vle8.v v8, (t3)
   644: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   648: d7019126                 vand.vv v3, v9, v2
   652: 573222a2                 vsrl.vi v4, v2, 4
   656: d7827132                 vrgather.vv v5, v7, v3
   660: 57038232                 vrgather.vv v6, v8, v4
   664: d782622e                 vxor.vv v5, v6, v5
   668: d780502e                 vxor.vv v1, v5, v1
   672: 83330501                 ld t2, 0x10(a0)
   676: b3836300                 add t2, t2, t1
   680: 07810302                 vle8.v v2, (t2)
   684: 5700202e                 vxor.vv v0, v2, v0
   688: 138f0200                 mv t5, t0
   692: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   696: 138e0f04                 addi t3, t6, 0x40
   700: 87030e02                 vle8.v v7, (t3)
   704: 138e0f05                 addi t3, t6, 0x50
   708: 07040e02                 vle8.v v8, (t3)
   712: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   716: d7019126                 vand.vv v3, v9, v2
   720: 573222a2                 vsrl.vi v4, v2, 4
   724: d7827132                 vrgather.vv v5, v7, v3
   728: 57038232                 vrgather.vv v6, v8, v4
   732: d782622e                 vxor.vv v5, v6, v5
   736: d780502e                 vxor.vv v1, v5, v1
   740: 83338501                 ld t2, 0x18(a0)
   744: b3836300                 add t2, t2, t1
   748: 07810302                 vle8.v v2, (t2)
   752: 5700202e                 vxor.vv v0, v2, v0
   756: 138f0200                 mv t5, t0
   760: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   764: 138e0f06                 addi t3, t6, 0x60
   768: 87030e02                 vle8.v v7, (t3)
   772: 138e0f07                 addi t3, t6, 0x70
   776: 07040e02                 vle8.v v8, (t3)
   780: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   784: d7019126                 vand.vv v3, v9, v2
   788: 573222a2                 vsrl.vi v4, v2, 4
   792: d7827132                 vrgather.vv v5, v7, v3
   796: 57038232                 vrgather.vv v6, v8, v4
   800: d782622e                 vxor.vv v5, v6, v5
   804: d780502e                 vxor.vv v1, v5, v1
   808: 83330502                 ld t2, 0x20(a0)
   812: b3836300                 add t2, t2, t1
   816: 07810302                 vle8.v v2, (t2)
   820: 5700202e                 vxor.vv v0, v2, v0
   824: 138f0200                 mv t5, t0
   828: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   832: 138e0f08                 addi t3, t6, 0x80
   836: 87030e02                 vle8.v v7, (t3)
   840: 138e0f09                 addi t3, t6, 0x90
   844: 07040e02                 vle8.v v8, (t3)
   848: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   852: d7019126                 vand.vv v3, v9, v2
   856: 573222a2                 vsrl.vi v4, v2, 4
   860: d7827132                 vrgather.vv v5, v7, v3
   864: 57038232                 vrgather.vv v6, v8, v4
   868: d782622e                 vxor.vv v5, v6, v5
   872: d780502e                 vxor.vv v1, v5, v1
   876: 83338502                 ld t2, 0x28(a0)
   880: b3836300                 add t2, t2, t1
   884: 07810302                 vle8.v v2, (t2)
   888: 5700202e                 vxor.vv v0, v2, v0
   892: 138f0200                 mv t5, t0
   896: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   900: 138e0f0a                 addi t3, t6, 0xa0
   904: 87030e02                 vle8.v v7, (t3)
   908: 138e0f0b                 addi t3, t6, 0xb0
   912: 07040e02                 vle8.v v8, (t3)
   916: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   920: d7019126                 vand.vv v3, v9, v2
   924: 573222a2                 vsrl.vi v4, v2, 4
   928: d7827132                 vrgather.vv v5, v7, v3
   932: 57038232                 vrgather.vv v6, v8, v4
   936: d782622e                 vxor.vv v5, v6, v5
   940: d780502e                 vxor.vv v1, v5, v1
   944: 83330503                 ld t2, 0x30(a0)
   948: b3836300                 add t2, t2, t1
   952: 07810302                 vle8.v v2, (t2)
   956: 5700202e                 vxor.vv v0, v2, v0
   960: 138f0200                 mv t5, t0
   964: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   968: 138e0f0c                 addi t3, t6, 0xc0
   972: 87030e02                 vle8.v v7, (t3)
   976: 138e0f0d                 addi t3, t6, 0xd0
   980: 07040e02                 vle8.v v8, (t3)
   984: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   988: d7019126                 vand.vv v3, v9, v2
   992: 573222a2                 vsrl.vi v4, v2, 4
   996: d7827132                 vrgather.vv v5, v7, v3
  1000: 57038232                 vrgather.vv v6, v8, v4
  1004: d782622e                 vxor.vv v5, v6, v5
  1008: d780502e                 vxor.vv v1, v5, v1
  1012: 83338503                 ld t2, 0x38(a0)
  1016: b3836300                 add t2, t2, t1
  1020: 07810302                 vle8.v v2, (t2)
  1024: 5700202e                 vxor.vv v0, v2, v0
  1028: d780202e                 vxor.vv v1, v2, v1
  1032: 83b30500                 ld t2, 0(a1)
  1036: b3836300                 add t2, t2, t1
  1040: 33005000                 ntl.all 
  1044: 27800302                 vse8.v v0, (t2)
  1048: 83b38500                 ld t2, 8(a1)
  1052: b3836300                 add t2, t2, t1
  1056: 33005000                 ntl.all 
  1060: a7800302                 vse8.v v1, (t2)
  1064: 33035300                 add t1, t1, t0
  1068: 33065640                 sub a2, a2, t0
  1072: 6ff0dfdd                 j 0x20c
  1076: 67800000                 ret 
  1080: 0000                     unimp 
  1082: 0000                     unimp 
  1084: 0000                     unimp 
  1086: 0000                     unimp 
