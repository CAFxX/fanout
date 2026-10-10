; config_t1_1733 target=rvv-zvbc32e twin=notail topo=raid5(4)/enc op=encode U=8
; code_size=528 twins=["notail"] uops_per_byte=0.5078125
     0: 63140600                 bnez a2, 8
     4: 6f00c01f                 j 0x200
     8: 970f0000                 auipc t6, 0
    12: 938f8f20                 addi t6, t6, 0x208
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93068000                 li a3, 8
    32: b386d202                 mul a3, t0, a3
    36: 3357d602                 divu a4, a2, a3
    40: 63140700                 bnez a4, 0x30
    44: 6f00401d                 j 0x200
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
   440: 27800302                 vse8.v v0, (t2)
   444: b3835300                 add t2, t2, t0
   448: a7800302                 vse8.v v1, (t2)
   452: b3835300                 add t2, t2, t0
   456: 27810302                 vse8.v v2, (t2)
   460: b3835300                 add t2, t2, t0
   464: a7810302                 vse8.v v3, (t2)
   468: b3835300                 add t2, t2, t0
   472: 27820302                 vse8.v v4, (t2)
   476: b3835300                 add t2, t2, t0
   480: a7820302                 vse8.v v5, (t2)
   484: b3835300                 add t2, t2, t0
   488: 27830302                 vse8.v v6, (t2)
   492: b3835300                 add t2, t2, t0
   496: a7830302                 vse8.v v7, (t2)
   500: 3303d300                 add t1, t1, a3
   504: 1307f7ff                 addi a4, a4, -1
   508: e31207e4                 bnez a4, 0x40
   512: 67800000                 ret 
   516: 0000                     unimp 
   518: 0000                     unimp 
   520: 0000                     unimp 
   522: 0000                     unimp 
   524: 0000                     unimp 
   526: 0000                     unimp 
