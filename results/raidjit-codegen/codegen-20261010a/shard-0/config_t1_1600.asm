; config_t1_1600 target=rvv twin=nt topo=rs(6,3)/rec{0} op=recover U=8
; code_size=896 twins=["nt"] uops_per_byte=0.859375
     0: 63140600                 bnez a2, 8
     4: 6f00c036                 j 0x370
     8: 970f0000                 auipc t6, 0
    12: 938f8f37                 addi t6, t6, 0x378
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93068000                 li a3, 8
    32: b386d202                 mul a3, t0, a3
    36: 3357d602                 divu a4, a2, a3
    40: 63140700                 bnez a4, 0x30
    44: 6f00002c                 j 0x2ec
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07800302                 vle8.v v0, (t2)
    76: b3835300                 add t2, t2, t0
    80: 87800302                 vle8.v v1, (t2)
    84: b3835300                 add t2, t2, t0
    88: 07810302                 vle8.v v2, (t2)
    92: b3835300                 add t2, t2, t0
    96: 87810302                 vle8.v v3, (t2)
   100: b3835300                 add t2, t2, t0
   104: 07820302                 vle8.v v4, (t2)
   108: b3835300                 add t2, t2, t0
   112: 87820302                 vle8.v v5, (t2)
   116: b3835300                 add t2, t2, t0
   120: 07830302                 vle8.v v6, (t2)
   124: b3835300                 add t2, t2, t0
   128: 87830302                 vle8.v v7, (t2)
   132: 83338500                 ld t2, 8(a0)
   136: b3836300                 add t2, t2, t1
   140: 07840302                 vle8.v v8, (t2)
   144: 5700802e                 vxor.vv v0, v8, v0
   148: b3835300                 add t2, t2, t0
   152: 07840302                 vle8.v v8, (t2)
   156: d780802e                 vxor.vv v1, v8, v1
   160: b3835300                 add t2, t2, t0
   164: 07840302                 vle8.v v8, (t2)
   168: 5701812e                 vxor.vv v2, v8, v2
   172: b3835300                 add t2, t2, t0
   176: 07840302                 vle8.v v8, (t2)
   180: d781812e                 vxor.vv v3, v8, v3
   184: b3835300                 add t2, t2, t0
   188: 07840302                 vle8.v v8, (t2)
   192: 5702822e                 vxor.vv v4, v8, v4
   196: b3835300                 add t2, t2, t0
   200: 07840302                 vle8.v v8, (t2)
   204: d782822e                 vxor.vv v5, v8, v5
   208: b3835300                 add t2, t2, t0
   212: 07840302                 vle8.v v8, (t2)
   216: 5703832e                 vxor.vv v6, v8, v6
   220: b3835300                 add t2, t2, t0
   224: 07840302                 vle8.v v8, (t2)
   228: d783832e                 vxor.vv v7, v8, v7
   232: 83330501                 ld t2, 0x10(a0)
   236: b3836300                 add t2, t2, t1
   240: 07840302                 vle8.v v8, (t2)
   244: 5700802e                 vxor.vv v0, v8, v0
   248: b3835300                 add t2, t2, t0
   252: 07840302                 vle8.v v8, (t2)
   256: d780802e                 vxor.vv v1, v8, v1
   260: b3835300                 add t2, t2, t0
   264: 07840302                 vle8.v v8, (t2)
   268: 5701812e                 vxor.vv v2, v8, v2
   272: b3835300                 add t2, t2, t0
   276: 07840302                 vle8.v v8, (t2)
   280: d781812e                 vxor.vv v3, v8, v3
   284: b3835300                 add t2, t2, t0
   288: 07840302                 vle8.v v8, (t2)
   292: 5702822e                 vxor.vv v4, v8, v4
   296: b3835300                 add t2, t2, t0
   300: 07840302                 vle8.v v8, (t2)
   304: d782822e                 vxor.vv v5, v8, v5
   308: b3835300                 add t2, t2, t0
   312: 07840302                 vle8.v v8, (t2)
   316: 5703832e                 vxor.vv v6, v8, v6
   320: b3835300                 add t2, t2, t0
   324: 07840302                 vle8.v v8, (t2)
   328: d783832e                 vxor.vv v7, v8, v7
   332: 83338501                 ld t2, 0x18(a0)
   336: b3836300                 add t2, t2, t1
   340: 07840302                 vle8.v v8, (t2)
   344: 5700802e                 vxor.vv v0, v8, v0
   348: b3835300                 add t2, t2, t0
   352: 07840302                 vle8.v v8, (t2)
   356: d780802e                 vxor.vv v1, v8, v1
   360: b3835300                 add t2, t2, t0
   364: 07840302                 vle8.v v8, (t2)
   368: 5701812e                 vxor.vv v2, v8, v2
   372: b3835300                 add t2, t2, t0
   376: 07840302                 vle8.v v8, (t2)
   380: d781812e                 vxor.vv v3, v8, v3
   384: b3835300                 add t2, t2, t0
   388: 07840302                 vle8.v v8, (t2)
   392: 5702822e                 vxor.vv v4, v8, v4
   396: b3835300                 add t2, t2, t0
   400: 07840302                 vle8.v v8, (t2)
   404: d782822e                 vxor.vv v5, v8, v5
   408: b3835300                 add t2, t2, t0
   412: 07840302                 vle8.v v8, (t2)
   416: 5703832e                 vxor.vv v6, v8, v6
   420: b3835300                 add t2, t2, t0
   424: 07840302                 vle8.v v8, (t2)
   428: d783832e                 vxor.vv v7, v8, v7
   432: 83330502                 ld t2, 0x20(a0)
   436: b3836300                 add t2, t2, t1
   440: 07840302                 vle8.v v8, (t2)
   444: 5700802e                 vxor.vv v0, v8, v0
   448: b3835300                 add t2, t2, t0
   452: 07840302                 vle8.v v8, (t2)
   456: d780802e                 vxor.vv v1, v8, v1
   460: b3835300                 add t2, t2, t0
   464: 07840302                 vle8.v v8, (t2)
   468: 5701812e                 vxor.vv v2, v8, v2
   472: b3835300                 add t2, t2, t0
   476: 07840302                 vle8.v v8, (t2)
   480: d781812e                 vxor.vv v3, v8, v3
   484: b3835300                 add t2, t2, t0
   488: 07840302                 vle8.v v8, (t2)
   492: 5702822e                 vxor.vv v4, v8, v4
   496: b3835300                 add t2, t2, t0
   500: 07840302                 vle8.v v8, (t2)
   504: d782822e                 vxor.vv v5, v8, v5
   508: b3835300                 add t2, t2, t0
   512: 07840302                 vle8.v v8, (t2)
   516: 5703832e                 vxor.vv v6, v8, v6
   520: b3835300                 add t2, t2, t0
   524: 07840302                 vle8.v v8, (t2)
   528: d783832e                 vxor.vv v7, v8, v7
   532: 83338502                 ld t2, 0x28(a0)
   536: b3836300                 add t2, t2, t1
   540: 07840302                 vle8.v v8, (t2)
   544: 5700802e                 vxor.vv v0, v8, v0
   548: b3835300                 add t2, t2, t0
   552: 07840302                 vle8.v v8, (t2)
   556: d780802e                 vxor.vv v1, v8, v1
   560: b3835300                 add t2, t2, t0
   564: 07840302                 vle8.v v8, (t2)
   568: 5701812e                 vxor.vv v2, v8, v2
   572: b3835300                 add t2, t2, t0
   576: 07840302                 vle8.v v8, (t2)
   580: d781812e                 vxor.vv v3, v8, v3
   584: b3835300                 add t2, t2, t0
   588: 07840302                 vle8.v v8, (t2)
   592: 5702822e                 vxor.vv v4, v8, v4
   596: b3835300                 add t2, t2, t0
   600: 07840302                 vle8.v v8, (t2)
   604: d782822e                 vxor.vv v5, v8, v5
   608: b3835300                 add t2, t2, t0
   612: 07840302                 vle8.v v8, (t2)
   616: 5703832e                 vxor.vv v6, v8, v6
   620: b3835300                 add t2, t2, t0
   624: 07840302                 vle8.v v8, (t2)
   628: d783832e                 vxor.vv v7, v8, v7
   632: 83b30500                 ld t2, 0(a1)
   636: b3836300                 add t2, t2, t1
   640: 33005000                 ntl.all 
   644: 27800302                 vse8.v v0, (t2)
   648: b3835300                 add t2, t2, t0
   652: 33005000                 ntl.all 
   656: a7800302                 vse8.v v1, (t2)
   660: b3835300                 add t2, t2, t0
   664: 33005000                 ntl.all 
   668: 27810302                 vse8.v v2, (t2)
   672: b3835300                 add t2, t2, t0
   676: 33005000                 ntl.all 
   680: a7810302                 vse8.v v3, (t2)
   684: b3835300                 add t2, t2, t0
   688: 33005000                 ntl.all 
   692: 27820302                 vse8.v v4, (t2)
   696: b3835300                 add t2, t2, t0
   700: 33005000                 ntl.all 
   704: a7820302                 vse8.v v5, (t2)
   708: b3835300                 add t2, t2, t0
   712: 33005000                 ntl.all 
   716: 27830302                 vse8.v v6, (t2)
   720: b3835300                 add t2, t2, t0
   724: 33005000                 ntl.all 
   728: a7830302                 vse8.v v7, (t2)
   732: 3303d300                 add t1, t1, a3
   736: 1307f7ff                 addi a4, a4, -1
   740: e31e07d4                 bnez a4, 0x40
   744: 33066640                 sub a2, a2, t1
   748: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   752: 63940200                 bnez t0, 0x2f8
   756: 6f00c007                 j 0x370
   760: 83330500                 ld t2, 0(a0)
   764: b3836300                 add t2, t2, t1
   768: 07800302                 vle8.v v0, (t2)
   772: 83338500                 ld t2, 8(a0)
   776: b3836300                 add t2, t2, t1
   780: 07840302                 vle8.v v8, (t2)
   784: 5700802e                 vxor.vv v0, v8, v0
   788: 83330501                 ld t2, 0x10(a0)
   792: b3836300                 add t2, t2, t1
   796: 07840302                 vle8.v v8, (t2)
   800: 5700802e                 vxor.vv v0, v8, v0
   804: 83338501                 ld t2, 0x18(a0)
   808: b3836300                 add t2, t2, t1
   812: 07840302                 vle8.v v8, (t2)
   816: 5700802e                 vxor.vv v0, v8, v0
   820: 83330502                 ld t2, 0x20(a0)
   824: b3836300                 add t2, t2, t1
   828: 07840302                 vle8.v v8, (t2)
   832: 5700802e                 vxor.vv v0, v8, v0
   836: 83338502                 ld t2, 0x28(a0)
   840: b3836300                 add t2, t2, t1
   844: 07840302                 vle8.v v8, (t2)
   848: 5700802e                 vxor.vv v0, v8, v0
   852: 83b30500                 ld t2, 0(a1)
   856: b3836300                 add t2, t2, t1
   860: 33005000                 ntl.all 
   864: 27800302                 vse8.v v0, (t2)
   868: 33035300                 add t1, t1, t0
   872: 33065640                 sub a2, a2, t0
   876: 6ff01ff8                 j 0x2ec
   880: 67800000                 ret 
   884: 0000                     unimp 
   886: 0000                     unimp 
   888: 0000                     unimp 
   890: 0000                     unimp 
   892: 0000                     unimp 
   894: 0000                     unimp 
