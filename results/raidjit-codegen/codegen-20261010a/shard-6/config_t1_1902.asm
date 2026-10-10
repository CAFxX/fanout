; config_t1_1902 target=rvv-zvbc32e twin=notail topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=656 twins=["notail"] uops_per_byte=7.8125
     0: 63140600                 bnez a2, 8
     4: 6f008027                 j 0x27c
     8: 970f0000                 auipc t6, 0
    12: 938f8f28                 addi t6, t6, 0x288
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d7440e5e                 vmv.v.x v9, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00c024                 j 0x27c
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07810302                 vle8.v v2, (t2)
    76: 138e0f00                 mv t3, t6
    80: 87030e02                 vle8.v v7, (t3)
    84: 138e0f01                 addi t3, t6, 0x10
    88: 07040e02                 vle8.v v8, (t3)
    92: d7019126                 vand.vv v3, v9, v2
    96: 573222a2                 vsrl.vi v4, v2, 4
   100: d7827132                 vrgather.vv v5, v7, v3
   104: 57038232                 vrgather.vv v6, v8, v4
   108: d782622e                 vxor.vv v5, v6, v5
   112: 5780025e                 vmv.v.v v0, v5
   116: 138e0f02                 addi t3, t6, 0x20
   120: 87030e02                 vle8.v v7, (t3)
   124: 138e0f03                 addi t3, t6, 0x30
   128: 07040e02                 vle8.v v8, (t3)
   132: d7019126                 vand.vv v3, v9, v2
   136: 573222a2                 vsrl.vi v4, v2, 4
   140: d7827132                 vrgather.vv v5, v7, v3
   144: 57038232                 vrgather.vv v6, v8, v4
   148: d782622e                 vxor.vv v5, v6, v5
   152: d780025e                 vmv.v.v v1, v5
   156: 83338500                 ld t2, 8(a0)
   160: b3836300                 add t2, t2, t1
   164: 07810302                 vle8.v v2, (t2)
   168: 138e0f04                 addi t3, t6, 0x40
   172: 87030e02                 vle8.v v7, (t3)
   176: 138e0f05                 addi t3, t6, 0x50
   180: 07040e02                 vle8.v v8, (t3)
   184: d7019126                 vand.vv v3, v9, v2
   188: 573222a2                 vsrl.vi v4, v2, 4
   192: d7827132                 vrgather.vv v5, v7, v3
   196: 57038232                 vrgather.vv v6, v8, v4
   200: d782622e                 vxor.vv v5, v6, v5
   204: 5700502e                 vxor.vv v0, v5, v0
   208: 138e0f06                 addi t3, t6, 0x60
   212: 87030e02                 vle8.v v7, (t3)
   216: 138e0f07                 addi t3, t6, 0x70
   220: 07040e02                 vle8.v v8, (t3)
   224: d7019126                 vand.vv v3, v9, v2
   228: 573222a2                 vsrl.vi v4, v2, 4
   232: d7827132                 vrgather.vv v5, v7, v3
   236: 57038232                 vrgather.vv v6, v8, v4
   240: d782622e                 vxor.vv v5, v6, v5
   244: d780502e                 vxor.vv v1, v5, v1
   248: 83330501                 ld t2, 0x10(a0)
   252: b3836300                 add t2, t2, t1
   256: 07810302                 vle8.v v2, (t2)
   260: 138e0f08                 addi t3, t6, 0x80
   264: 87030e02                 vle8.v v7, (t3)
   268: 138e0f09                 addi t3, t6, 0x90
   272: 07040e02                 vle8.v v8, (t3)
   276: d7019126                 vand.vv v3, v9, v2
   280: 573222a2                 vsrl.vi v4, v2, 4
   284: d7827132                 vrgather.vv v5, v7, v3
   288: 57038232                 vrgather.vv v6, v8, v4
   292: d782622e                 vxor.vv v5, v6, v5
   296: 5700502e                 vxor.vv v0, v5, v0
   300: 138e0f0a                 addi t3, t6, 0xa0
   304: 87030e02                 vle8.v v7, (t3)
   308: 138e0f0b                 addi t3, t6, 0xb0
   312: 07040e02                 vle8.v v8, (t3)
   316: d7019126                 vand.vv v3, v9, v2
   320: 573222a2                 vsrl.vi v4, v2, 4
   324: d7827132                 vrgather.vv v5, v7, v3
   328: 57038232                 vrgather.vv v6, v8, v4
   332: d782622e                 vxor.vv v5, v6, v5
   336: d780502e                 vxor.vv v1, v5, v1
   340: 83338501                 ld t2, 0x18(a0)
   344: b3836300                 add t2, t2, t1
   348: 07810302                 vle8.v v2, (t2)
   352: 138e0f0c                 addi t3, t6, 0xc0
   356: 87030e02                 vle8.v v7, (t3)
   360: 138e0f0d                 addi t3, t6, 0xd0
   364: 07040e02                 vle8.v v8, (t3)
   368: d7019126                 vand.vv v3, v9, v2
   372: 573222a2                 vsrl.vi v4, v2, 4
   376: d7827132                 vrgather.vv v5, v7, v3
   380: 57038232                 vrgather.vv v6, v8, v4
   384: d782622e                 vxor.vv v5, v6, v5
   388: 5700502e                 vxor.vv v0, v5, v0
   392: 138e0f0e                 addi t3, t6, 0xe0
   396: 87030e02                 vle8.v v7, (t3)
   400: 138e0f0f                 addi t3, t6, 0xf0
   404: 07040e02                 vle8.v v8, (t3)
   408: d7019126                 vand.vv v3, v9, v2
   412: 573222a2                 vsrl.vi v4, v2, 4
   416: d7827132                 vrgather.vv v5, v7, v3
   420: 57038232                 vrgather.vv v6, v8, v4
   424: d782622e                 vxor.vv v5, v6, v5
   428: d780502e                 vxor.vv v1, v5, v1
   432: 83330502                 ld t2, 0x20(a0)
   436: b3836300                 add t2, t2, t1
   440: 07810302                 vle8.v v2, (t2)
   444: 138e0f10                 addi t3, t6, 0x100
   448: 87030e02                 vle8.v v7, (t3)
   452: 138e0f11                 addi t3, t6, 0x110
   456: 07040e02                 vle8.v v8, (t3)
   460: d7019126                 vand.vv v3, v9, v2
   464: 573222a2                 vsrl.vi v4, v2, 4
   468: d7827132                 vrgather.vv v5, v7, v3
   472: 57038232                 vrgather.vv v6, v8, v4
   476: d782622e                 vxor.vv v5, v6, v5
   480: 5700502e                 vxor.vv v0, v5, v0
   484: 138e0f12                 addi t3, t6, 0x120
   488: 87030e02                 vle8.v v7, (t3)
   492: 138e0f13                 addi t3, t6, 0x130
   496: 07040e02                 vle8.v v8, (t3)
   500: d7019126                 vand.vv v3, v9, v2
   504: 573222a2                 vsrl.vi v4, v2, 4
   508: d7827132                 vrgather.vv v5, v7, v3
   512: 57038232                 vrgather.vv v6, v8, v4
   516: d782622e                 vxor.vv v5, v6, v5
   520: d780502e                 vxor.vv v1, v5, v1
   524: 83338502                 ld t2, 0x28(a0)
   528: b3836300                 add t2, t2, t1
   532: 07810302                 vle8.v v2, (t2)
   536: 138e0f14                 addi t3, t6, 0x140
   540: 87030e02                 vle8.v v7, (t3)
   544: 138e0f15                 addi t3, t6, 0x150
   548: 07040e02                 vle8.v v8, (t3)
   552: d7019126                 vand.vv v3, v9, v2
   556: 573222a2                 vsrl.vi v4, v2, 4
   560: d7827132                 vrgather.vv v5, v7, v3
   564: 57038232                 vrgather.vv v6, v8, v4
   568: d782622e                 vxor.vv v5, v6, v5
   572: 5700502e                 vxor.vv v0, v5, v0
   576: d7019126                 vand.vv v3, v9, v2
   580: 573222a2                 vsrl.vi v4, v2, 4
   584: d7827132                 vrgather.vv v5, v7, v3
   588: 57038232                 vrgather.vv v6, v8, v4
   592: d782622e                 vxor.vv v5, v6, v5
   596: d780502e                 vxor.vv v1, v5, v1
   600: 83b30500                 ld t2, 0(a1)
   604: b3836300                 add t2, t2, t1
   608: 27800302                 vse8.v v0, (t2)
   612: 83b38500                 ld t2, 8(a1)
   616: b3836300                 add t2, t2, t1
   620: a7800302                 vse8.v v1, (t2)
   624: 3303d300                 add t1, t1, a3
   628: 1307f7ff                 addi a4, a4, -1
   632: e31407dc                 bnez a4, 0x40
   636: 67800000                 ret 
   640: 0000                     unimp 
   642: 0000                     unimp 
   644: 0000                     unimp 
   646: 0000                     unimp 
   648: 0000                     unimp 
   650: 0000                     unimp 
   652: 0000                     unimp 
   654: 0000                     unimp 
