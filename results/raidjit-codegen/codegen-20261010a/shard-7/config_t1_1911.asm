; config_t1_1911 target=rvv-zvbc32e twin=baseline topo=rs(6,3)/rec{0,1} op=recover U=8
; code_size=1024 twins=[] uops_per_byte=2.84375
     0: 63140600                 bnez a2, 8
     4: 6f00403f                 j 0x3f8
     8: 970f0000                 auipc t6, 0
    12: 938f8f3f                 addi t6, t6, 0x3f8
    16: 577f200c                 vsetvli t5, zero, e8, m4, ta, ma
    20: 130eb001                 li t3, 0x1b
    24: 574e0e5e                 vmv.v.x v28, t3
    28: 13030000                 li t1, 0
    32: d772200c                 vsetvli t0, zero, e8, m4, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00801e                 j 0x218
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07840302                 vle8.v v8, (t2)
    76: 5700045e                 vmv.v.v v0, v8
    80: 83338500                 ld t2, 8(a0)
    84: b3836300                 add t2, t2, t1
    88: 07840302                 vle8.v v8, (t2)
    92: 5700802e                 vxor.vv v0, v8, v0
    96: 83330501                 ld t2, 0x10(a0)
   100: b3836300                 add t2, t2, t1
   104: 07840302                 vle8.v v8, (t2)
   108: 5700802e                 vxor.vv v0, v8, v0
   112: 83338501                 ld t2, 0x18(a0)
   116: b3836300                 add t2, t2, t1
   120: 07840302                 vle8.v v8, (t2)
   124: 5700802e                 vxor.vv v0, v8, v0
   128: 83330502                 ld t2, 0x20(a0)
   132: b3836300                 add t2, t2, t1
   136: 07840302                 vle8.v v8, (t2)
   140: 5700802e                 vxor.vv v0, v8, v0
   144: 83330500                 ld t2, 0(a0)
   148: b3836300                 add t2, t2, t1
   152: 130e8000                 li t3, 8
   156: 574c0e5e                 vmv.v.x v24, t3
   160: 07840302                 vle8.v v8, (t2)
   164: 57268c32                 vclmul.vv v12, v8, v24
   168: 57288c36                 vclmulh.vv v16, v8, v24
   172: 572a0e33                 vclmul.vv v20, v16, v28
   176: 57280e37                 vclmulh.vv v16, v16, v28
   180: 57280e33                 vclmul.vv v16, v16, v28
   184: 5706462f                 vxor.vv v12, v20, v12
   188: 5706062f                 vxor.vv v12, v16, v12
   192: 5702065e                 vmv.v.v v4, v12
   196: 83338500                 ld t2, 8(a0)
   200: b3836300                 add t2, t2, t1
   204: 130e4000                 li t3, 4
   208: 574c0e5e                 vmv.v.x v24, t3
   212: 07840302                 vle8.v v8, (t2)
   216: 57268c32                 vclmul.vv v12, v8, v24
   220: 57288c36                 vclmulh.vv v16, v8, v24
   224: 572a0e33                 vclmul.vv v20, v16, v28
   228: 57280e37                 vclmulh.vv v16, v16, v28
   232: 57280e33                 vclmul.vv v16, v16, v28
   236: 5706462f                 vxor.vv v12, v20, v12
   240: 5706062f                 vxor.vv v12, v16, v12
   244: 5702c22e                 vxor.vv v4, v12, v4
   248: 83330501                 ld t2, 0x10(a0)
   252: b3836300                 add t2, t2, t1
   256: 130e2000                 li t3, 2
   260: 574c0e5e                 vmv.v.x v24, t3
   264: 07840302                 vle8.v v8, (t2)
   268: 57268c32                 vclmul.vv v12, v8, v24
   272: 57288c36                 vclmulh.vv v16, v8, v24
   276: 572a0e33                 vclmul.vv v20, v16, v28
   280: 57280e37                 vclmulh.vv v16, v16, v28
   284: 57280e33                 vclmul.vv v16, v16, v28
   288: 5706462f                 vxor.vv v12, v20, v12
   292: 5706062f                 vxor.vv v12, v16, v12
   296: 5702c22e                 vxor.vv v4, v12, v4
   300: 83338501                 ld t2, 0x18(a0)
   304: b3836300                 add t2, t2, t1
   308: 07840302                 vle8.v v8, (t2)
   312: 5702822e                 vxor.vv v4, v8, v4
   316: 83338502                 ld t2, 0x28(a0)
   320: b3836300                 add t2, t2, t1
   324: 07840302                 vle8.v v8, (t2)
   328: 5702822e                 vxor.vv v4, v8, v4
   332: 130eb00c                 li t3, 0xcb
   336: 574c0e5e                 vmv.v.x v24, t3
   340: 57260c32                 vclmul.vv v12, v0, v24
   344: 57280c36                 vclmulh.vv v16, v0, v24
   348: 572a0e33                 vclmul.vv v20, v16, v28
   352: 57280e37                 vclmulh.vv v16, v16, v28
   356: 57280e33                 vclmul.vv v16, v16, v28
   360: 5706462f                 vxor.vv v12, v20, v12
   364: 5706062f                 vxor.vv v12, v16, v12
   368: 5704065e                 vmv.v.v v8, v12
   372: 5704442e                 vxor.vv v8, v4, v8
   376: 130e500f                 li t3, 0xf5
   380: 574c0e5e                 vmv.v.x v24, t3
   384: 57268c32                 vclmul.vv v12, v8, v24
   388: 57288c36                 vclmulh.vv v16, v8, v24
   392: 572a0e33                 vclmul.vv v20, v16, v28
   396: 57280e37                 vclmulh.vv v16, v16, v28
   400: 57280e33                 vclmul.vv v16, v16, v28
   404: 5706462f                 vxor.vv v12, v20, v12
   408: 5706062f                 vxor.vv v12, v16, v12
   412: 5704065e                 vmv.v.v v8, v12
   416: 130ed008                 li t3, 0x8d
   420: 574c0e5e                 vmv.v.x v24, t3
   424: 57260c32                 vclmul.vv v12, v0, v24
   428: 57280c36                 vclmulh.vv v16, v0, v24
   432: 572a0e33                 vclmul.vv v20, v16, v28
   436: 57280e37                 vclmulh.vv v16, v16, v28
   440: 57280e33                 vclmul.vv v16, v16, v28
   444: 5706462f                 vxor.vv v12, v20, v12
   448: 5706062f                 vxor.vv v12, v16, v12
   452: 5702c22e                 vxor.vv v4, v12, v4
   456: 130e500f                 li t3, 0xf5
   460: 574c0e5e                 vmv.v.x v24, t3
   464: 57264c32                 vclmul.vv v12, v4, v24
   468: 57284c36                 vclmulh.vv v16, v4, v24
   472: 572a0e33                 vclmul.vv v20, v16, v28
   476: 57280e37                 vclmulh.vv v16, v16, v28
   480: 57280e33                 vclmul.vv v16, v16, v28
   484: 5706462f                 vxor.vv v12, v20, v12
   488: 5706062f                 vxor.vv v12, v16, v12
   492: 5702065e                 vmv.v.v v4, v12
   496: 83b30500                 ld t2, 0(a1)
   500: b3836300                 add t2, t2, t1
   504: 27840302                 vse8.v v8, (t2)
   508: 83b38500                 ld t2, 8(a1)
   512: b3836300                 add t2, t2, t1
   516: 27820302                 vse8.v v4, (t2)
   520: 3303d300                 add t1, t1, a3
   524: 1307f7ff                 addi a4, a4, -1
   528: e31807e2                 bnez a4, 0x40
   532: 33066640                 sub a2, a2, t1
   536: d772260c                 vsetvli t0, a2, e8, m4, ta, ma
   540: 63940200                 bnez t0, 0x224
   544: 6f00801d                 j 0x3f8
   548: 83330500                 ld t2, 0(a0)
   552: b3836300                 add t2, t2, t1
   556: 07840302                 vle8.v v8, (t2)
   560: 5700045e                 vmv.v.v v0, v8
   564: 83338500                 ld t2, 8(a0)
   568: b3836300                 add t2, t2, t1
   572: 07840302                 vle8.v v8, (t2)
   576: 5700802e                 vxor.vv v0, v8, v0
   580: 83330501                 ld t2, 0x10(a0)
   584: b3836300                 add t2, t2, t1
   588: 07840302                 vle8.v v8, (t2)
   592: 5700802e                 vxor.vv v0, v8, v0
   596: 83338501                 ld t2, 0x18(a0)
   600: b3836300                 add t2, t2, t1
   604: 07840302                 vle8.v v8, (t2)
   608: 5700802e                 vxor.vv v0, v8, v0
   612: 83330502                 ld t2, 0x20(a0)
   616: b3836300                 add t2, t2, t1
   620: 07840302                 vle8.v v8, (t2)
   624: 5700802e                 vxor.vv v0, v8, v0
   628: 83330500                 ld t2, 0(a0)
   632: b3836300                 add t2, t2, t1
   636: 130e8000                 li t3, 8
   640: 574c0e5e                 vmv.v.x v24, t3
   644: 07840302                 vle8.v v8, (t2)
   648: 57268c32                 vclmul.vv v12, v8, v24
   652: 57288c36                 vclmulh.vv v16, v8, v24
   656: 572a0e33                 vclmul.vv v20, v16, v28
   660: 57280e37                 vclmulh.vv v16, v16, v28
   664: 57280e33                 vclmul.vv v16, v16, v28
   668: 5706462f                 vxor.vv v12, v20, v12
   672: 5706062f                 vxor.vv v12, v16, v12
   676: 5702065e                 vmv.v.v v4, v12
   680: 83338500                 ld t2, 8(a0)
   684: b3836300                 add t2, t2, t1
   688: 130e4000                 li t3, 4
   692: 574c0e5e                 vmv.v.x v24, t3
   696: 07840302                 vle8.v v8, (t2)
   700: 57268c32                 vclmul.vv v12, v8, v24
   704: 57288c36                 vclmulh.vv v16, v8, v24
   708: 572a0e33                 vclmul.vv v20, v16, v28
   712: 57280e37                 vclmulh.vv v16, v16, v28
   716: 57280e33                 vclmul.vv v16, v16, v28
   720: 5706462f                 vxor.vv v12, v20, v12
   724: 5706062f                 vxor.vv v12, v16, v12
   728: 5702c22e                 vxor.vv v4, v12, v4
   732: 83330501                 ld t2, 0x10(a0)
   736: b3836300                 add t2, t2, t1
   740: 130e2000                 li t3, 2
   744: 574c0e5e                 vmv.v.x v24, t3
   748: 07840302                 vle8.v v8, (t2)
   752: 57268c32                 vclmul.vv v12, v8, v24
   756: 57288c36                 vclmulh.vv v16, v8, v24
   760: 572a0e33                 vclmul.vv v20, v16, v28
   764: 57280e37                 vclmulh.vv v16, v16, v28
   768: 57280e33                 vclmul.vv v16, v16, v28
   772: 5706462f                 vxor.vv v12, v20, v12
   776: 5706062f                 vxor.vv v12, v16, v12
   780: 5702c22e                 vxor.vv v4, v12, v4
   784: 83338501                 ld t2, 0x18(a0)
   788: b3836300                 add t2, t2, t1
   792: 07840302                 vle8.v v8, (t2)
   796: 5702822e                 vxor.vv v4, v8, v4
   800: 83338502                 ld t2, 0x28(a0)
   804: b3836300                 add t2, t2, t1
   808: 07840302                 vle8.v v8, (t2)
   812: 5702822e                 vxor.vv v4, v8, v4
   816: 130eb00c                 li t3, 0xcb
   820: 574c0e5e                 vmv.v.x v24, t3
   824: 57260c32                 vclmul.vv v12, v0, v24
   828: 57280c36                 vclmulh.vv v16, v0, v24
   832: 572a0e33                 vclmul.vv v20, v16, v28
   836: 57280e37                 vclmulh.vv v16, v16, v28
   840: 57280e33                 vclmul.vv v16, v16, v28
   844: 5706462f                 vxor.vv v12, v20, v12
   848: 5706062f                 vxor.vv v12, v16, v12
   852: 5704065e                 vmv.v.v v8, v12
   856: 5704442e                 vxor.vv v8, v4, v8
   860: 130e500f                 li t3, 0xf5
   864: 574c0e5e                 vmv.v.x v24, t3
   868: 57268c32                 vclmul.vv v12, v8, v24
   872: 57288c36                 vclmulh.vv v16, v8, v24
   876: 572a0e33                 vclmul.vv v20, v16, v28
   880: 57280e37                 vclmulh.vv v16, v16, v28
   884: 57280e33                 vclmul.vv v16, v16, v28
   888: 5706462f                 vxor.vv v12, v20, v12
   892: 5706062f                 vxor.vv v12, v16, v12
   896: 5704065e                 vmv.v.v v8, v12
   900: 130ed008                 li t3, 0x8d
   904: 574c0e5e                 vmv.v.x v24, t3
   908: 57260c32                 vclmul.vv v12, v0, v24
   912: 57280c36                 vclmulh.vv v16, v0, v24
   916: 572a0e33                 vclmul.vv v20, v16, v28
   920: 57280e37                 vclmulh.vv v16, v16, v28
   924: 57280e33                 vclmul.vv v16, v16, v28
   928: 5706462f                 vxor.vv v12, v20, v12
   932: 5706062f                 vxor.vv v12, v16, v12
   936: 5702c22e                 vxor.vv v4, v12, v4
   940: 130e500f                 li t3, 0xf5
   944: 574c0e5e                 vmv.v.x v24, t3
   948: 57264c32                 vclmul.vv v12, v4, v24
   952: 57284c36                 vclmulh.vv v16, v4, v24
   956: 572a0e33                 vclmul.vv v20, v16, v28
   960: 57280e37                 vclmulh.vv v16, v16, v28
   964: 57280e33                 vclmul.vv v16, v16, v28
   968: 5706462f                 vxor.vv v12, v20, v12
   972: 5706062f                 vxor.vv v12, v16, v12
   976: 5702065e                 vmv.v.v v4, v12
   980: 83b30500                 ld t2, 0(a1)
   984: b3836300                 add t2, t2, t1
   988: 27840302                 vse8.v v8, (t2)
   992: 83b38500                 ld t2, 8(a1)
   996: b3836300                 add t2, t2, t1
  1000: 27820302                 vse8.v v4, (t2)
  1004: 33035300                 add t1, t1, t0
  1008: 33065640                 sub a2, a2, t0
  1012: 6ff05fe2                 j 0x218
  1016: 67800000                 ret 
  1020: 0000                     unimp 
  1022: 0000                     unimp 
