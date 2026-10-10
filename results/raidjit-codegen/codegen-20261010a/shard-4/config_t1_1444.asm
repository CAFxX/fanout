; config_t1_1444 target=rvv twin=nt topo=raid5(4)/enc op=encode U=8
; code_size=656 twins=["nt"] uops_per_byte=0.578125
     0: 63140600                 bnez a2, 8
     4: 6f004028                 j 0x288
     8: 970f0000                 auipc t6, 0
    12: 938f8f28                 addi t6, t6, 0x288
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93068000                 li a3, 8
    32: b386d202                 mul a3, t0, a3
    36: 3357d602                 divu a4, a2, a3
    40: 63140700                 bnez a4, 0x30
    44: 6f00801f                 j 0x224
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
   432: 83b30500                 ld t2, 0(a1)
   436: b3836300                 add t2, t2, t1
   440: 33005000                 ntl.all 
   444: 27800302                 vse8.v v0, (t2)
   448: b3835300                 add t2, t2, t0
   452: 33005000                 ntl.all 
   456: a7800302                 vse8.v v1, (t2)
   460: b3835300                 add t2, t2, t0
   464: 33005000                 ntl.all 
   468: 27810302                 vse8.v v2, (t2)
   472: b3835300                 add t2, t2, t0
   476: 33005000                 ntl.all 
   480: a7810302                 vse8.v v3, (t2)
   484: b3835300                 add t2, t2, t0
   488: 33005000                 ntl.all 
   492: 27820302                 vse8.v v4, (t2)
   496: b3835300                 add t2, t2, t0
   500: 33005000                 ntl.all 
   504: a7820302                 vse8.v v5, (t2)
   508: b3835300                 add t2, t2, t0
   512: 33005000                 ntl.all 
   516: 27830302                 vse8.v v6, (t2)
   520: b3835300                 add t2, t2, t0
   524: 33005000                 ntl.all 
   528: a7830302                 vse8.v v7, (t2)
   532: 3303d300                 add t1, t1, a3
   536: 1307f7ff                 addi a4, a4, -1
   540: e31207e2                 bnez a4, 0x40
   544: 33066640                 sub a2, a2, t1
   548: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   552: 63940200                 bnez t0, 0x230
   556: 6f00c005                 j 0x288
   560: 83330500                 ld t2, 0(a0)
   564: b3836300                 add t2, t2, t1
   568: 07800302                 vle8.v v0, (t2)
   572: 83338500                 ld t2, 8(a0)
   576: b3836300                 add t2, t2, t1
   580: 07840302                 vle8.v v8, (t2)
   584: 5700802e                 vxor.vv v0, v8, v0
   588: 83330501                 ld t2, 0x10(a0)
   592: b3836300                 add t2, t2, t1
   596: 07840302                 vle8.v v8, (t2)
   600: 5700802e                 vxor.vv v0, v8, v0
   604: 83338501                 ld t2, 0x18(a0)
   608: b3836300                 add t2, t2, t1
   612: 07840302                 vle8.v v8, (t2)
   616: 5700802e                 vxor.vv v0, v8, v0
   620: 83b30500                 ld t2, 0(a1)
   624: b3836300                 add t2, t2, t1
   628: 33005000                 ntl.all 
   632: 27800302                 vse8.v v0, (t2)
   636: 33035300                 add t1, t1, t0
   640: 33065640                 sub a2, a2, t0
   644: 6ff01ffa                 j 0x224
   648: 67800000                 ret 
   652: 0000                     unimp 
   654: 0000                     unimp 
