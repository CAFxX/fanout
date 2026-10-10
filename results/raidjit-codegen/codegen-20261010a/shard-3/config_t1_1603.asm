; config_t1_1603 target=rvv twin=notail topo=rs(6,3)/rec{0} op=recover U=8
; code_size=720 twins=["notail"] uops_per_byte=0.7578125
     0: 63140600                 bnez a2, 8
     4: 6f00402c                 j 0x2c8
     8: 970f0000                 auipc t6, 0
    12: 938f8f2c                 addi t6, t6, 0x2c8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93068000                 li a3, 8
    32: b386d202                 mul a3, t0, a3
    36: 3357d602                 divu a4, a2, a3
    40: 63140700                 bnez a4, 0x30
    44: 6f00c029                 j 0x2c8
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
   640: 27800302                 vse8.v v0, (t2)
   644: b3835300                 add t2, t2, t0
   648: a7800302                 vse8.v v1, (t2)
   652: b3835300                 add t2, t2, t0
   656: 27810302                 vse8.v v2, (t2)
   660: b3835300                 add t2, t2, t0
   664: a7810302                 vse8.v v3, (t2)
   668: b3835300                 add t2, t2, t0
   672: 27820302                 vse8.v v4, (t2)
   676: b3835300                 add t2, t2, t0
   680: a7820302                 vse8.v v5, (t2)
   684: b3835300                 add t2, t2, t0
   688: 27830302                 vse8.v v6, (t2)
   692: b3835300                 add t2, t2, t0
   696: a7830302                 vse8.v v7, (t2)
   700: 3303d300                 add t1, t1, a3
   704: 1307f7ff                 addi a4, a4, -1
   708: e31e07d6                 bnez a4, 0x40
   712: 67800000                 ret 
   716: 0000                     unimp 
   718: 0000                     unimp 
