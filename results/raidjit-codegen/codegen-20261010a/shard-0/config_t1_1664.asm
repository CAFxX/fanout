; config_t1_1664 target=rvv twin=baseline topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=624 twins=[] uops_per_byte=1.40625
     0: 63140600                 bnez a2, 8
     4: 6f004026                 j 0x268
     8: 970f0000                 auipc t6, 0
    12: 938f8f26                 addi t6, t6, 0x268
    16: 577f100c                 vsetvli t5, zero, e8, m2, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574e0e5e                 vmv.v.x v28, t3
    28: 13030000                 li t1, 0
    32: d772100c                 vsetvli t0, zero, e8, m2, ta, ma
    36: 93961200                 slli a3, t0, 1
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f000015                 j 0x180
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330501                 ld t2, 0x10(a0)
    68: b3836300                 add t2, t2, t1
    72: 07800302                 vle8.v v0, (t2)
    76: b3835300                 add t2, t2, t0
    80: 07810302                 vle8.v v2, (t2)
    84: 83338501                 ld t2, 0x18(a0)
    88: b3836300                 add t2, t2, t1
    92: 07820302                 vle8.v v4, (t2)
    96: b3835300                 add t2, t2, t0
   100: 07830302                 vle8.v v6, (t2)
   104: 83330502                 ld t2, 0x20(a0)
   108: b3836300                 add t2, t2, t1
   112: 07840302                 vle8.v v8, (t2)
   116: b3835300                 add t2, t2, t0
   120: 07850302                 vle8.v v10, (t2)
   124: 83330500                 ld t2, 0(a0)
   128: b3836300                 add t2, t2, t1
   132: 833e8500                 ld t4, 8(a0)
   136: b38e6e00                 add t4, t4, t1
   140: 07880302                 vle8.v v16, (t2)
   144: 07890e02                 vle8.v v18, (t4)
   148: 5706282f                 vxor.vv v12, v18, v16
   152: b3835300                 add t2, t2, t0
   156: b38e5e00                 add t4, t4, t0
   160: 07880302                 vle8.v v16, (t2)
   164: 07890e02                 vle8.v v18, (t4)
   168: 5707282f                 vxor.vv v14, v18, v16
   172: 5700c02e                 vxor.vv v0, v12, v0
   176: 5701e12e                 vxor.vv v2, v14, v2
   180: 138e0f00                 mv t3, t6
   184: 070c0e02                 vle8.v v24, (t3)
   188: 138e0f02                 addi t3, t6, 0x20
   192: 070d0e02                 vle8.v v26, (t3)
   196: 5708c627                 vand.vv v16, v28, v12
   200: 5739c2a2                 vsrl.vi v18, v12, 4
   204: 570a8833                 vrgather.vv v20, v24, v16
   208: 570ba933                 vrgather.vv v22, v26, v18
   212: 570a6a2f                 vxor.vv v20, v22, v20
   216: 5702422f                 vxor.vv v4, v20, v4
   220: 5708c727                 vand.vv v16, v28, v14
   224: 5739e2a2                 vsrl.vi v18, v14, 4
   228: 570a8833                 vrgather.vv v20, v24, v16
   232: 570ba933                 vrgather.vv v22, v26, v18
   236: 570a6a2f                 vxor.vv v20, v22, v20
   240: 5703432f                 vxor.vv v6, v20, v6
   244: 138e0f04                 addi t3, t6, 0x40
   248: 070c0e02                 vle8.v v24, (t3)
   252: 138e0f06                 addi t3, t6, 0x60
   256: 070d0e02                 vle8.v v26, (t3)
   260: 5708c627                 vand.vv v16, v28, v12
   264: 5739c2a2                 vsrl.vi v18, v12, 4
   268: 570a8833                 vrgather.vv v20, v24, v16
   272: 570ba933                 vrgather.vv v22, v26, v18
   276: 570a6a2f                 vxor.vv v20, v22, v20
   280: 5704442f                 vxor.vv v8, v20, v8
   284: 5708c727                 vand.vv v16, v28, v14
   288: 5739e2a2                 vsrl.vi v18, v14, 4
   292: 570a8833                 vrgather.vv v20, v24, v16
   296: 570ba933                 vrgather.vv v22, v26, v18
   300: 570a6a2f                 vxor.vv v20, v22, v20
   304: 5705452f                 vxor.vv v10, v20, v10
   308: 83b30500                 ld t2, 0(a1)
   312: b3836300                 add t2, t2, t1
   316: 27800302                 vse8.v v0, (t2)
   320: b3835300                 add t2, t2, t0
   324: 27810302                 vse8.v v2, (t2)
   328: 83b38500                 ld t2, 8(a1)
   332: b3836300                 add t2, t2, t1
   336: 27820302                 vse8.v v4, (t2)
   340: b3835300                 add t2, t2, t0
   344: 27830302                 vse8.v v6, (t2)
   348: 83b30501                 ld t2, 0x10(a1)
   352: b3836300                 add t2, t2, t1
   356: 27840302                 vse8.v v8, (t2)
   360: b3835300                 add t2, t2, t0
   364: 27850302                 vse8.v v10, (t2)
   368: 3303d300                 add t1, t1, a3
   372: 1307f7ff                 addi a4, a4, -1
   376: e31407ec                 bnez a4, 0x40
   380: 33066640                 sub a2, a2, t1
   384: d772160c                 vsetvli t0, a2, e8, m2, ta, ma
   388: 63940200                 bnez t0, 0x18c
   392: 6f00000e                 j 0x268
   396: 83330501                 ld t2, 0x10(a0)
   400: b3836300                 add t2, t2, t1
   404: 07800302                 vle8.v v0, (t2)
   408: 83338501                 ld t2, 0x18(a0)
   412: b3836300                 add t2, t2, t1
   416: 07820302                 vle8.v v4, (t2)
   420: 83330502                 ld t2, 0x20(a0)
   424: b3836300                 add t2, t2, t1
   428: 07840302                 vle8.v v8, (t2)
   432: 83330500                 ld t2, 0(a0)
   436: b3836300                 add t2, t2, t1
   440: 833e8500                 ld t4, 8(a0)
   444: b38e6e00                 add t4, t4, t1
   448: 07880302                 vle8.v v16, (t2)
   452: 07890e02                 vle8.v v18, (t4)
   456: 5706282f                 vxor.vv v12, v18, v16
   460: 5700c02e                 vxor.vv v0, v12, v0
   464: 138f0200                 mv t5, t0
   468: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
   472: 138e0f00                 mv t3, t6
   476: 070c0e02                 vle8.v v24, (t3)
   480: 138e0f02                 addi t3, t6, 0x20
   484: 070d0e02                 vle8.v v26, (t3)
   488: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
   492: 5708c627                 vand.vv v16, v28, v12
   496: 5739c2a2                 vsrl.vi v18, v12, 4
   500: 570a8833                 vrgather.vv v20, v24, v16
   504: 570ba933                 vrgather.vv v22, v26, v18
   508: 570a6a2f                 vxor.vv v20, v22, v20
   512: 5702422f                 vxor.vv v4, v20, v4
   516: 138f0200                 mv t5, t0
   520: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
   524: 138e0f04                 addi t3, t6, 0x40
   528: 070c0e02                 vle8.v v24, (t3)
   532: 138e0f06                 addi t3, t6, 0x60
   536: 070d0e02                 vle8.v v26, (t3)
   540: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
   544: 5708c627                 vand.vv v16, v28, v12
   548: 5739c2a2                 vsrl.vi v18, v12, 4
   552: 570a8833                 vrgather.vv v20, v24, v16
   556: 570ba933                 vrgather.vv v22, v26, v18
   560: 570a6a2f                 vxor.vv v20, v22, v20
   564: 5704442f                 vxor.vv v8, v20, v8
   568: 83b30500                 ld t2, 0(a1)
   572: b3836300                 add t2, t2, t1
   576: 27800302                 vse8.v v0, (t2)
   580: 83b38500                 ld t2, 8(a1)
   584: b3836300                 add t2, t2, t1
   588: 27820302                 vse8.v v4, (t2)
   592: 83b30501                 ld t2, 0x10(a1)
   596: b3836300                 add t2, t2, t1
   600: 27840302                 vse8.v v8, (t2)
   604: 33035300                 add t1, t1, t0
   608: 33065640                 sub a2, a2, t0
   612: 6ff0dff1                 j 0x180
   616: 67800000                 ret 
   620: 0000                     unimp 
   622: 0000                     unimp 
