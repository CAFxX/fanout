; config_t1_1624 target=rvv twin=rec_cf topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1120 twins=[] uops_per_byte=12.6875
     0: 63140600                 bnez a2, 8
     4: 6f008044                 j 0x44c
     8: 970f0000                 auipc t6, 0
    12: 938f8f45                 addi t6, t6, 0x458
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d7440e5e                 vmv.v.x v9, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00801e                 j 0x218
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07810302                 vle8.v v2, (t2)
    76: 5700015e                 vmv.v.v v0, v2
    80: 83338500                 ld t2, 8(a0)
    84: b3836300                 add t2, t2, t1
    88: 07810302                 vle8.v v2, (t2)
    92: 5700202e                 vxor.vv v0, v2, v0
    96: 83330501                 ld t2, 0x10(a0)
   100: b3836300                 add t2, t2, t1
   104: 07810302                 vle8.v v2, (t2)
   108: 5700202e                 vxor.vv v0, v2, v0
   112: 83338501                 ld t2, 0x18(a0)
   116: b3836300                 add t2, t2, t1
   120: 07810302                 vle8.v v2, (t2)
   124: 5700202e                 vxor.vv v0, v2, v0
   128: 83330502                 ld t2, 0x20(a0)
   132: b3836300                 add t2, t2, t1
   136: 07810302                 vle8.v v2, (t2)
   140: 5700202e                 vxor.vv v0, v2, v0
   144: 83330500                 ld t2, 0(a0)
   148: b3836300                 add t2, t2, t1
   152: 138e0f00                 mv t3, t6
   156: 87030e02                 vle8.v v7, (t3)
   160: 138e0f01                 addi t3, t6, 0x10
   164: 07040e02                 vle8.v v8, (t3)
   168: 07810302                 vle8.v v2, (t2)
   172: d7019126                 vand.vv v3, v9, v2
   176: 573222a2                 vsrl.vi v4, v2, 4
   180: d7827132                 vrgather.vv v5, v7, v3
   184: 57038232                 vrgather.vv v6, v8, v4
   188: d782622e                 vxor.vv v5, v6, v5
   192: d780025e                 vmv.v.v v1, v5
   196: 83338500                 ld t2, 8(a0)
   200: b3836300                 add t2, t2, t1
   204: 138e0f02                 addi t3, t6, 0x20
   208: 87030e02                 vle8.v v7, (t3)
   212: 138e0f03                 addi t3, t6, 0x30
   216: 07040e02                 vle8.v v8, (t3)
   220: 07810302                 vle8.v v2, (t2)
   224: d7019126                 vand.vv v3, v9, v2
   228: 573222a2                 vsrl.vi v4, v2, 4
   232: d7827132                 vrgather.vv v5, v7, v3
   236: 57038232                 vrgather.vv v6, v8, v4
   240: d782622e                 vxor.vv v5, v6, v5
   244: d780502e                 vxor.vv v1, v5, v1
   248: 83330501                 ld t2, 0x10(a0)
   252: b3836300                 add t2, t2, t1
   256: 138e0f04                 addi t3, t6, 0x40
   260: 87030e02                 vle8.v v7, (t3)
   264: 138e0f05                 addi t3, t6, 0x50
   268: 07040e02                 vle8.v v8, (t3)
   272: 07810302                 vle8.v v2, (t2)
   276: d7019126                 vand.vv v3, v9, v2
   280: 573222a2                 vsrl.vi v4, v2, 4
   284: d7827132                 vrgather.vv v5, v7, v3
   288: 57038232                 vrgather.vv v6, v8, v4
   292: d782622e                 vxor.vv v5, v6, v5
   296: d780502e                 vxor.vv v1, v5, v1
   300: 83338501                 ld t2, 0x18(a0)
   304: b3836300                 add t2, t2, t1
   308: 07810302                 vle8.v v2, (t2)
   312: d780202e                 vxor.vv v1, v2, v1
   316: 83338502                 ld t2, 0x28(a0)
   320: b3836300                 add t2, t2, t1
   324: 07810302                 vle8.v v2, (t2)
   328: d780202e                 vxor.vv v1, v2, v1
   332: 138e0f06                 addi t3, t6, 0x60
   336: 87030e02                 vle8.v v7, (t3)
   340: 138e0f07                 addi t3, t6, 0x70
   344: 07040e02                 vle8.v v8, (t3)
   348: d7019026                 vand.vv v3, v9, v0
   352: 573202a2                 vsrl.vi v4, v0, 4
   356: d7827132                 vrgather.vv v5, v7, v3
   360: 57038232                 vrgather.vv v6, v8, v4
   364: d782622e                 vxor.vv v5, v6, v5
   368: 5781025e                 vmv.v.v v2, v5
   372: 5701112e                 vxor.vv v2, v1, v2
   376: 138e0f08                 addi t3, t6, 0x80
   380: 87030e02                 vle8.v v7, (t3)
   384: 138e0f09                 addi t3, t6, 0x90
   388: 07040e02                 vle8.v v8, (t3)
   392: d7019126                 vand.vv v3, v9, v2
   396: 573222a2                 vsrl.vi v4, v2, 4
   400: d7827132                 vrgather.vv v5, v7, v3
   404: 57038232                 vrgather.vv v6, v8, v4
   408: d782622e                 vxor.vv v5, v6, v5
   412: 5781025e                 vmv.v.v v2, v5
   416: 138e0f0a                 addi t3, t6, 0xa0
   420: 87030e02                 vle8.v v7, (t3)
   424: 138e0f0b                 addi t3, t6, 0xb0
   428: 07040e02                 vle8.v v8, (t3)
   432: d7019026                 vand.vv v3, v9, v0
   436: 573202a2                 vsrl.vi v4, v0, 4
   440: d7827132                 vrgather.vv v5, v7, v3
   444: 57038232                 vrgather.vv v6, v8, v4
   448: d782622e                 vxor.vv v5, v6, v5
   452: d780502e                 vxor.vv v1, v5, v1
   456: 138e0f08                 addi t3, t6, 0x80
   460: 87030e02                 vle8.v v7, (t3)
   464: 138e0f09                 addi t3, t6, 0x90
   468: 07040e02                 vle8.v v8, (t3)
   472: d7819026                 vand.vv v3, v9, v1
   476: 573212a2                 vsrl.vi v4, v1, 4
   480: d7827132                 vrgather.vv v5, v7, v3
   484: 57038232                 vrgather.vv v6, v8, v4
   488: d782622e                 vxor.vv v5, v6, v5
   492: d780025e                 vmv.v.v v1, v5
   496: 83b30500                 ld t2, 0(a1)
   500: b3836300                 add t2, t2, t1
   504: 27810302                 vse8.v v2, (t2)
   508: 83b38500                 ld t2, 8(a1)
   512: b3836300                 add t2, t2, t1
   516: a7800302                 vse8.v v1, (t2)
   520: 3303d300                 add t1, t1, a3
   524: 1307f7ff                 addi a4, a4, -1
   528: e31807e2                 bnez a4, 0x40
   532: 33066640                 sub a2, a2, t1
   536: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   540: 63940200                 bnez t0, 0x224
   544: 6f00c022                 j 0x44c
   548: 83330500                 ld t2, 0(a0)
   552: b3836300                 add t2, t2, t1
   556: 07810302                 vle8.v v2, (t2)
   560: 5700015e                 vmv.v.v v0, v2
   564: 83338500                 ld t2, 8(a0)
   568: b3836300                 add t2, t2, t1
   572: 07810302                 vle8.v v2, (t2)
   576: 5700202e                 vxor.vv v0, v2, v0
   580: 83330501                 ld t2, 0x10(a0)
   584: b3836300                 add t2, t2, t1
   588: 07810302                 vle8.v v2, (t2)
   592: 5700202e                 vxor.vv v0, v2, v0
   596: 83338501                 ld t2, 0x18(a0)
   600: b3836300                 add t2, t2, t1
   604: 07810302                 vle8.v v2, (t2)
   608: 5700202e                 vxor.vv v0, v2, v0
   612: 83330502                 ld t2, 0x20(a0)
   616: b3836300                 add t2, t2, t1
   620: 07810302                 vle8.v v2, (t2)
   624: 5700202e                 vxor.vv v0, v2, v0
   628: 83330500                 ld t2, 0(a0)
   632: b3836300                 add t2, t2, t1
   636: 138f0200                 mv t5, t0
   640: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   644: 138e0f00                 mv t3, t6
   648: 87030e02                 vle8.v v7, (t3)
   652: 138e0f01                 addi t3, t6, 0x10
   656: 07040e02                 vle8.v v8, (t3)
   660: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   664: 07810302                 vle8.v v2, (t2)
   668: d7019126                 vand.vv v3, v9, v2
   672: 573222a2                 vsrl.vi v4, v2, 4
   676: d7827132                 vrgather.vv v5, v7, v3
   680: 57038232                 vrgather.vv v6, v8, v4
   684: d782622e                 vxor.vv v5, v6, v5
   688: d780025e                 vmv.v.v v1, v5
   692: 83338500                 ld t2, 8(a0)
   696: b3836300                 add t2, t2, t1
   700: 138f0200                 mv t5, t0
   704: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   708: 138e0f02                 addi t3, t6, 0x20
   712: 87030e02                 vle8.v v7, (t3)
   716: 138e0f03                 addi t3, t6, 0x30
   720: 07040e02                 vle8.v v8, (t3)
   724: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   728: 07810302                 vle8.v v2, (t2)
   732: d7019126                 vand.vv v3, v9, v2
   736: 573222a2                 vsrl.vi v4, v2, 4
   740: d7827132                 vrgather.vv v5, v7, v3
   744: 57038232                 vrgather.vv v6, v8, v4
   748: d782622e                 vxor.vv v5, v6, v5
   752: d780502e                 vxor.vv v1, v5, v1
   756: 83330501                 ld t2, 0x10(a0)
   760: b3836300                 add t2, t2, t1
   764: 138f0200                 mv t5, t0
   768: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   772: 138e0f04                 addi t3, t6, 0x40
   776: 87030e02                 vle8.v v7, (t3)
   780: 138e0f05                 addi t3, t6, 0x50
   784: 07040e02                 vle8.v v8, (t3)
   788: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   792: 07810302                 vle8.v v2, (t2)
   796: d7019126                 vand.vv v3, v9, v2
   800: 573222a2                 vsrl.vi v4, v2, 4
   804: d7827132                 vrgather.vv v5, v7, v3
   808: 57038232                 vrgather.vv v6, v8, v4
   812: d782622e                 vxor.vv v5, v6, v5
   816: d780502e                 vxor.vv v1, v5, v1
   820: 83338501                 ld t2, 0x18(a0)
   824: b3836300                 add t2, t2, t1
   828: 07810302                 vle8.v v2, (t2)
   832: d780202e                 vxor.vv v1, v2, v1
   836: 83338502                 ld t2, 0x28(a0)
   840: b3836300                 add t2, t2, t1
   844: 07810302                 vle8.v v2, (t2)
   848: d780202e                 vxor.vv v1, v2, v1
   852: 138f0200                 mv t5, t0
   856: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   860: 138e0f06                 addi t3, t6, 0x60
   864: 87030e02                 vle8.v v7, (t3)
   868: 138e0f07                 addi t3, t6, 0x70
   872: 07040e02                 vle8.v v8, (t3)
   876: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   880: d7019026                 vand.vv v3, v9, v0
   884: 573202a2                 vsrl.vi v4, v0, 4
   888: d7827132                 vrgather.vv v5, v7, v3
   892: 57038232                 vrgather.vv v6, v8, v4
   896: d782622e                 vxor.vv v5, v6, v5
   900: 5781025e                 vmv.v.v v2, v5
   904: 5701112e                 vxor.vv v2, v1, v2
   908: 138f0200                 mv t5, t0
   912: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   916: 138e0f08                 addi t3, t6, 0x80
   920: 87030e02                 vle8.v v7, (t3)
   924: 138e0f09                 addi t3, t6, 0x90
   928: 07040e02                 vle8.v v8, (t3)
   932: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   936: d7019126                 vand.vv v3, v9, v2
   940: 573222a2                 vsrl.vi v4, v2, 4
   944: d7827132                 vrgather.vv v5, v7, v3
   948: 57038232                 vrgather.vv v6, v8, v4
   952: d782622e                 vxor.vv v5, v6, v5
   956: 5781025e                 vmv.v.v v2, v5
   960: 138f0200                 mv t5, t0
   964: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
   968: 138e0f0a                 addi t3, t6, 0xa0
   972: 87030e02                 vle8.v v7, (t3)
   976: 138e0f0b                 addi t3, t6, 0xb0
   980: 07040e02                 vle8.v v8, (t3)
   984: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
   988: d7019026                 vand.vv v3, v9, v0
   992: 573202a2                 vsrl.vi v4, v0, 4
   996: d7827132                 vrgather.vv v5, v7, v3
  1000: 57038232                 vrgather.vv v6, v8, v4
  1004: d782622e                 vxor.vv v5, v6, v5
  1008: d780502e                 vxor.vv v1, v5, v1
  1012: 138f0200                 mv t5, t0
  1016: 577008cc                 vsetivli zero, 0x10, e8, m1, ta, ma
  1020: 138e0f08                 addi t3, t6, 0x80
  1024: 87030e02                 vle8.v v7, (t3)
  1028: 138e0f09                 addi t3, t6, 0x90
  1032: 07040e02                 vle8.v v8, (t3)
  1036: 57700f0c                 vsetvli zero, t5, e8, m1, ta, ma
  1040: d7819026                 vand.vv v3, v9, v1
  1044: 573212a2                 vsrl.vi v4, v1, 4
  1048: d7827132                 vrgather.vv v5, v7, v3
  1052: 57038232                 vrgather.vv v6, v8, v4
  1056: d782622e                 vxor.vv v5, v6, v5
  1060: d780025e                 vmv.v.v v1, v5
  1064: 83b30500                 ld t2, 0(a1)
  1068: b3836300                 add t2, t2, t1
  1072: 27810302                 vse8.v v2, (t2)
  1076: 83b38500                 ld t2, 8(a1)
  1080: b3836300                 add t2, t2, t1
  1084: a7800302                 vse8.v v1, (t2)
  1088: 33035300                 add t1, t1, t0
  1092: 33065640                 sub a2, a2, t0
  1096: 6ff01fdd                 j 0x218
  1100: 67800000                 ret 
  1104: 0000                     unimp 
  1106: 0000                     unimp 
  1108: 0000                     unimp 
  1110: 0000                     unimp 
  1112: 0000                     unimp 
  1114: 0000                     unimp 
  1116: 0000                     unimp 
  1118: 0000                     unimp 
