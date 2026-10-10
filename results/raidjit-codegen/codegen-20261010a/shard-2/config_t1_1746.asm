; config_t1_1746 target=rvv-zvbc32e twin=notail topo=raid6(8)/enc op=encode U=1
; code_size=528 twins=["notail"] uops_per_byte=5.625
     0: 63140600                 bnez a2, 8
     4: 6f00c01f                 j 0x200
     8: 970f0000                 auipc t6, 0
    12: 938f8f20                 addi t6, t6, 0x208
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: d7440e5e                 vmv.v.x v9, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93960200                 slli a3, t0, 0
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f00001d                 j 0x200
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
   484: 27800302                 vse8.v v0, (t2)
   488: 83b38500                 ld t2, 8(a1)
   492: b3836300                 add t2, t2, t1
   496: a7800302                 vse8.v v1, (t2)
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
