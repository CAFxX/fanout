; config_t1_1742 target=rvv-zvbc32e twin=baseline topo=raid6(8)/enc op=encode U=1
; code_size=992 twins=[] uops_per_byte=2.8125
     0: 63140600                 bnez a2, 8
     4: 6f00c03c                 j 0x3d0
     8: 970f0000                 auipc t6, 0
    12: 938f8f3d                 addi t6, t6, 0x3d8
    16: 577f200c                 vsetvli t5, zero, e8, m4, ta, ma
    20: 130eb001                 li t3, 0x1b
    24: 574e0e5e                 vmv.v.x v28, t3
    28: 13030000                 li t1, 0
    32: d772200c                 vsetvli t0, zero, e8, m4, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00401d                 j 0x204
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07840302                 vle8.v v8, (t2)
    76: 5700045e                 vmv.v.v v0, v8
    80: 130ed008                 li t3, 0x8d
    84: 574c0e5e                 vmv.v.x v24, t3
    88: 57268c32                 vclmul.vv v12, v8, v24
    92: 57288c36                 vclmulh.vv v16, v8, v24
    96: 572a0e33                 vclmul.vv v20, v16, v28
   100: 57280e37                 vclmulh.vv v16, v16, v28
   104: 57280e33                 vclmul.vv v16, v16, v28
   108: 5706462f                 vxor.vv v12, v20, v12
   112: 5706062f                 vxor.vv v12, v16, v12
   116: 5702065e                 vmv.v.v v4, v12
   120: 83338500                 ld t2, 8(a0)
   124: b3836300                 add t2, t2, t1
   128: 07840302                 vle8.v v8, (t2)
   132: 5700802e                 vxor.vv v0, v8, v0
   136: 130eb00c                 li t3, 0xcb
   140: 574c0e5e                 vmv.v.x v24, t3
   144: 57268c32                 vclmul.vv v12, v8, v24
   148: 57288c36                 vclmulh.vv v16, v8, v24
   152: 572a0e33                 vclmul.vv v20, v16, v28
   156: 57280e37                 vclmulh.vv v16, v16, v28
   160: 57280e33                 vclmul.vv v16, v16, v28
   164: 5706462f                 vxor.vv v12, v20, v12
   168: 5706062f                 vxor.vv v12, v16, v12
   172: 5702c22e                 vxor.vv v4, v12, v4
   176: 83330501                 ld t2, 0x10(a0)
   180: b3836300                 add t2, t2, t1
   184: 07840302                 vle8.v v8, (t2)
   188: 5700802e                 vxor.vv v0, v8, v0
   192: 130ec008                 li t3, 0x8c
   196: 574c0e5e                 vmv.v.x v24, t3
   200: 57268c32                 vclmul.vv v12, v8, v24
   204: 57288c36                 vclmulh.vv v16, v8, v24
   208: 572a0e33                 vclmul.vv v20, v16, v28
   212: 57280e37                 vclmulh.vv v16, v16, v28
   216: 57280e33                 vclmul.vv v16, v16, v28
   220: 5706462f                 vxor.vv v12, v20, v12
   224: 5706062f                 vxor.vv v12, v16, v12
   228: 5702c22e                 vxor.vv v4, v12, v4
   232: 83338501                 ld t2, 0x18(a0)
   236: b3836300                 add t2, t2, t1
   240: 07840302                 vle8.v v8, (t2)
   244: 5700802e                 vxor.vv v0, v8, v0
   248: 130e6004                 li t3, 0x46
   252: 574c0e5e                 vmv.v.x v24, t3
   256: 57268c32                 vclmul.vv v12, v8, v24
   260: 57288c36                 vclmulh.vv v16, v8, v24
   264: 572a0e33                 vclmul.vv v20, v16, v28
   268: 57280e37                 vclmulh.vv v16, v16, v28
   272: 57280e33                 vclmul.vv v16, v16, v28
   276: 5706462f                 vxor.vv v12, v20, v12
   280: 5706062f                 vxor.vv v12, v16, v12
   284: 5702c22e                 vxor.vv v4, v12, v4
   288: 83330502                 ld t2, 0x20(a0)
   292: b3836300                 add t2, t2, t1
   296: 07840302                 vle8.v v8, (t2)
   300: 5700802e                 vxor.vv v0, v8, v0
   304: 130e8000                 li t3, 8
   308: 574c0e5e                 vmv.v.x v24, t3
   312: 57268c32                 vclmul.vv v12, v8, v24
   316: 57288c36                 vclmulh.vv v16, v8, v24
   320: 572a0e33                 vclmul.vv v20, v16, v28
   324: 57280e37                 vclmulh.vv v16, v16, v28
   328: 57280e33                 vclmul.vv v16, v16, v28
   332: 5706462f                 vxor.vv v12, v20, v12
   336: 5706062f                 vxor.vv v12, v16, v12
   340: 5702c22e                 vxor.vv v4, v12, v4
   344: 83338502                 ld t2, 0x28(a0)
   348: b3836300                 add t2, t2, t1
   352: 07840302                 vle8.v v8, (t2)
   356: 5700802e                 vxor.vv v0, v8, v0
   360: 130e4000                 li t3, 4
   364: 574c0e5e                 vmv.v.x v24, t3
   368: 57268c32                 vclmul.vv v12, v8, v24
   372: 57288c36                 vclmulh.vv v16, v8, v24
   376: 572a0e33                 vclmul.vv v20, v16, v28
   380: 57280e37                 vclmulh.vv v16, v16, v28
   384: 57280e33                 vclmul.vv v16, v16, v28
   388: 5706462f                 vxor.vv v12, v20, v12
   392: 5706062f                 vxor.vv v12, v16, v12
   396: 5702c22e                 vxor.vv v4, v12, v4
   400: 83330503                 ld t2, 0x30(a0)
   404: b3836300                 add t2, t2, t1
   408: 07840302                 vle8.v v8, (t2)
   412: 5700802e                 vxor.vv v0, v8, v0
   416: 130e2000                 li t3, 2
   420: 574c0e5e                 vmv.v.x v24, t3
   424: 57268c32                 vclmul.vv v12, v8, v24
   428: 57288c36                 vclmulh.vv v16, v8, v24
   432: 572a0e33                 vclmul.vv v20, v16, v28
   436: 57280e37                 vclmulh.vv v16, v16, v28
   440: 57280e33                 vclmul.vv v16, v16, v28
   444: 5706462f                 vxor.vv v12, v20, v12
   448: 5706062f                 vxor.vv v12, v16, v12
   452: 5702c22e                 vxor.vv v4, v12, v4
   456: 83338503                 ld t2, 0x38(a0)
   460: b3836300                 add t2, t2, t1
   464: 07840302                 vle8.v v8, (t2)
   468: 5700802e                 vxor.vv v0, v8, v0
   472: 5702822e                 vxor.vv v4, v8, v4
   476: 83b30500                 ld t2, 0(a1)
   480: b3836300                 add t2, t2, t1
   484: 27800302                 vse8.v v0, (t2)
   488: 83b38500                 ld t2, 8(a1)
   492: b3836300                 add t2, t2, t1
   496: 27820302                 vse8.v v4, (t2)
   500: 3303d300                 add t1, t1, a3
   504: 1307f7ff                 addi a4, a4, -1
   508: e31207e4                 bnez a4, 0x40
   512: 33066640                 sub a2, a2, t1
   516: d772260c                 vsetvli t0, a2, e8, m4, ta, ma
   520: 63940200                 bnez t0, 0x210
   524: 6f00401c                 j 0x3d0
   528: 83330500                 ld t2, 0(a0)
   532: b3836300                 add t2, t2, t1
   536: 07840302                 vle8.v v8, (t2)
   540: 5700045e                 vmv.v.v v0, v8
   544: 130ed008                 li t3, 0x8d
   548: 574c0e5e                 vmv.v.x v24, t3
   552: 57268c32                 vclmul.vv v12, v8, v24
   556: 57288c36                 vclmulh.vv v16, v8, v24
   560: 572a0e33                 vclmul.vv v20, v16, v28
   564: 57280e37                 vclmulh.vv v16, v16, v28
   568: 57280e33                 vclmul.vv v16, v16, v28
   572: 5706462f                 vxor.vv v12, v20, v12
   576: 5706062f                 vxor.vv v12, v16, v12
   580: 5702065e                 vmv.v.v v4, v12
   584: 83338500                 ld t2, 8(a0)
   588: b3836300                 add t2, t2, t1
   592: 07840302                 vle8.v v8, (t2)
   596: 5700802e                 vxor.vv v0, v8, v0
   600: 130eb00c                 li t3, 0xcb
   604: 574c0e5e                 vmv.v.x v24, t3
   608: 57268c32                 vclmul.vv v12, v8, v24
   612: 57288c36                 vclmulh.vv v16, v8, v24
   616: 572a0e33                 vclmul.vv v20, v16, v28
   620: 57280e37                 vclmulh.vv v16, v16, v28
   624: 57280e33                 vclmul.vv v16, v16, v28
   628: 5706462f                 vxor.vv v12, v20, v12
   632: 5706062f                 vxor.vv v12, v16, v12
   636: 5702c22e                 vxor.vv v4, v12, v4
   640: 83330501                 ld t2, 0x10(a0)
   644: b3836300                 add t2, t2, t1
   648: 07840302                 vle8.v v8, (t2)
   652: 5700802e                 vxor.vv v0, v8, v0
   656: 130ec008                 li t3, 0x8c
   660: 574c0e5e                 vmv.v.x v24, t3
   664: 57268c32                 vclmul.vv v12, v8, v24
   668: 57288c36                 vclmulh.vv v16, v8, v24
   672: 572a0e33                 vclmul.vv v20, v16, v28
   676: 57280e37                 vclmulh.vv v16, v16, v28
   680: 57280e33                 vclmul.vv v16, v16, v28
   684: 5706462f                 vxor.vv v12, v20, v12
   688: 5706062f                 vxor.vv v12, v16, v12
   692: 5702c22e                 vxor.vv v4, v12, v4
   696: 83338501                 ld t2, 0x18(a0)
   700: b3836300                 add t2, t2, t1
   704: 07840302                 vle8.v v8, (t2)
   708: 5700802e                 vxor.vv v0, v8, v0
   712: 130e6004                 li t3, 0x46
   716: 574c0e5e                 vmv.v.x v24, t3
   720: 57268c32                 vclmul.vv v12, v8, v24
   724: 57288c36                 vclmulh.vv v16, v8, v24
   728: 572a0e33                 vclmul.vv v20, v16, v28
   732: 57280e37                 vclmulh.vv v16, v16, v28
   736: 57280e33                 vclmul.vv v16, v16, v28
   740: 5706462f                 vxor.vv v12, v20, v12
   744: 5706062f                 vxor.vv v12, v16, v12
   748: 5702c22e                 vxor.vv v4, v12, v4
   752: 83330502                 ld t2, 0x20(a0)
   756: b3836300                 add t2, t2, t1
   760: 07840302                 vle8.v v8, (t2)
   764: 5700802e                 vxor.vv v0, v8, v0
   768: 130e8000                 li t3, 8
   772: 574c0e5e                 vmv.v.x v24, t3
   776: 57268c32                 vclmul.vv v12, v8, v24
   780: 57288c36                 vclmulh.vv v16, v8, v24
   784: 572a0e33                 vclmul.vv v20, v16, v28
   788: 57280e37                 vclmulh.vv v16, v16, v28
   792: 57280e33                 vclmul.vv v16, v16, v28
   796: 5706462f                 vxor.vv v12, v20, v12
   800: 5706062f                 vxor.vv v12, v16, v12
   804: 5702c22e                 vxor.vv v4, v12, v4
   808: 83338502                 ld t2, 0x28(a0)
   812: b3836300                 add t2, t2, t1
   816: 07840302                 vle8.v v8, (t2)
   820: 5700802e                 vxor.vv v0, v8, v0
   824: 130e4000                 li t3, 4
   828: 574c0e5e                 vmv.v.x v24, t3
   832: 57268c32                 vclmul.vv v12, v8, v24
   836: 57288c36                 vclmulh.vv v16, v8, v24
   840: 572a0e33                 vclmul.vv v20, v16, v28
   844: 57280e37                 vclmulh.vv v16, v16, v28
   848: 57280e33                 vclmul.vv v16, v16, v28
   852: 5706462f                 vxor.vv v12, v20, v12
   856: 5706062f                 vxor.vv v12, v16, v12
   860: 5702c22e                 vxor.vv v4, v12, v4
   864: 83330503                 ld t2, 0x30(a0)
   868: b3836300                 add t2, t2, t1
   872: 07840302                 vle8.v v8, (t2)
   876: 5700802e                 vxor.vv v0, v8, v0
   880: 130e2000                 li t3, 2
   884: 574c0e5e                 vmv.v.x v24, t3
   888: 57268c32                 vclmul.vv v12, v8, v24
   892: 57288c36                 vclmulh.vv v16, v8, v24
   896: 572a0e33                 vclmul.vv v20, v16, v28
   900: 57280e37                 vclmulh.vv v16, v16, v28
   904: 57280e33                 vclmul.vv v16, v16, v28
   908: 5706462f                 vxor.vv v12, v20, v12
   912: 5706062f                 vxor.vv v12, v16, v12
   916: 5702c22e                 vxor.vv v4, v12, v4
   920: 83338503                 ld t2, 0x38(a0)
   924: b3836300                 add t2, t2, t1
   928: 07840302                 vle8.v v8, (t2)
   932: 5700802e                 vxor.vv v0, v8, v0
   936: 5702822e                 vxor.vv v4, v8, v4
   940: 83b30500                 ld t2, 0(a1)
   944: b3836300                 add t2, t2, t1
   948: 27800302                 vse8.v v0, (t2)
   952: 83b38500                 ld t2, 8(a1)
   956: b3836300                 add t2, t2, t1
   960: 27820302                 vse8.v v4, (t2)
   964: 33035300                 add t1, t1, t0
   968: 33065640                 sub a2, a2, t0
   972: 6ff09fe3                 j 0x204
   976: 67800000                 ret 
   980: 0000                     unimp 
   982: 0000                     unimp 
   984: 0000                     unimp 
   986: 0000                     unimp 
   988: 0000                     unimp 
   990: 0000                     unimp 
