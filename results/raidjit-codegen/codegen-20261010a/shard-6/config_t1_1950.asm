; config_t1_1950 target=rvv-zvbc32e twin=baseline topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=624 twins=[] uops_per_byte=1.375
     0: 63140600                 bnez a2, 8
     4: 6f00c025                 j 0x260
     8: 970f0000                 auipc t6, 0
    12: 938f8f26                 addi t6, t6, 0x268
    16: 577f100c                 vsetvli t5, zero, e8, m2, ta, ma
    20: 130eb001                 li t3, 0x1b
    24: 574e0e5e                 vmv.v.x v28, t3
    28: 13030000                 li t1, 0
    32: d772100c                 vsetvli t0, zero, e8, m2, ta, ma
    36: 93961200                 slli a3, t0, 1
    40: 3357d602                 divu a4, a2, a3
    44: 63140700                 bnez a4, 0x34
    48: 6f000016                 j 0x190
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
   180: 130ed008                 li t3, 0x8d
   184: 574d0e5e                 vmv.v.x v26, t3
   188: 572acd32                 vclmul.vv v20, v12, v26
   192: 572bcd36                 vclmulh.vv v22, v12, v26
   196: 572c6e33                 vclmul.vv v24, v22, v28
   200: 572b6e37                 vclmulh.vv v22, v22, v28
   204: 572b6e33                 vclmul.vv v22, v22, v28
   208: 570a8a2f                 vxor.vv v20, v24, v20
   212: 570a6a2f                 vxor.vv v20, v22, v20
   216: 5702422f                 vxor.vv v4, v20, v4
   220: 572aed32                 vclmul.vv v20, v14, v26
   224: 572bed36                 vclmulh.vv v22, v14, v26
   228: 572c6e33                 vclmul.vv v24, v22, v28
   232: 572b6e37                 vclmulh.vv v22, v22, v28
   236: 572b6e33                 vclmul.vv v22, v22, v28
   240: 570a8a2f                 vxor.vv v20, v24, v20
   244: 570a6a2f                 vxor.vv v20, v22, v20
   248: 5703432f                 vxor.vv v6, v20, v6
   252: 130eb00c                 li t3, 0xcb
   256: 574d0e5e                 vmv.v.x v26, t3
   260: 572acd32                 vclmul.vv v20, v12, v26
   264: 572bcd36                 vclmulh.vv v22, v12, v26
   268: 572c6e33                 vclmul.vv v24, v22, v28
   272: 572b6e37                 vclmulh.vv v22, v22, v28
   276: 572b6e33                 vclmul.vv v22, v22, v28
   280: 570a8a2f                 vxor.vv v20, v24, v20
   284: 570a6a2f                 vxor.vv v20, v22, v20
   288: 5704442f                 vxor.vv v8, v20, v8
   292: 572aed32                 vclmul.vv v20, v14, v26
   296: 572bed36                 vclmulh.vv v22, v14, v26
   300: 572c6e33                 vclmul.vv v24, v22, v28
   304: 572b6e37                 vclmulh.vv v22, v22, v28
   308: 572b6e33                 vclmul.vv v22, v22, v28
   312: 570a8a2f                 vxor.vv v20, v24, v20
   316: 570a6a2f                 vxor.vv v20, v22, v20
   320: 5705452f                 vxor.vv v10, v20, v10
   324: 83b30500                 ld t2, 0(a1)
   328: b3836300                 add t2, t2, t1
   332: 27800302                 vse8.v v0, (t2)
   336: b3835300                 add t2, t2, t0
   340: 27810302                 vse8.v v2, (t2)
   344: 83b38500                 ld t2, 8(a1)
   348: b3836300                 add t2, t2, t1
   352: 27820302                 vse8.v v4, (t2)
   356: b3835300                 add t2, t2, t0
   360: 27830302                 vse8.v v6, (t2)
   364: 83b30501                 ld t2, 0x10(a1)
   368: b3836300                 add t2, t2, t1
   372: 27840302                 vse8.v v8, (t2)
   376: b3835300                 add t2, t2, t0
   380: 27850302                 vse8.v v10, (t2)
   384: 3303d300                 add t1, t1, a3
   388: 1307f7ff                 addi a4, a4, -1
   392: e31c07ea                 bnez a4, 0x40
   396: 33066640                 sub a2, a2, t1
   400: d772160c                 vsetvli t0, a2, e8, m2, ta, ma
   404: 63940200                 bnez t0, 0x19c
   408: 6f00800c                 j 0x260
   412: 83330501                 ld t2, 0x10(a0)
   416: b3836300                 add t2, t2, t1
   420: 07800302                 vle8.v v0, (t2)
   424: 83338501                 ld t2, 0x18(a0)
   428: b3836300                 add t2, t2, t1
   432: 07820302                 vle8.v v4, (t2)
   436: 83330502                 ld t2, 0x20(a0)
   440: b3836300                 add t2, t2, t1
   444: 07840302                 vle8.v v8, (t2)
   448: 83330500                 ld t2, 0(a0)
   452: b3836300                 add t2, t2, t1
   456: 833e8500                 ld t4, 8(a0)
   460: b38e6e00                 add t4, t4, t1
   464: 07880302                 vle8.v v16, (t2)
   468: 07890e02                 vle8.v v18, (t4)
   472: 5706282f                 vxor.vv v12, v18, v16
   476: 5700c02e                 vxor.vv v0, v12, v0
   480: 130ed008                 li t3, 0x8d
   484: 574d0e5e                 vmv.v.x v26, t3
   488: 572acd32                 vclmul.vv v20, v12, v26
   492: 572bcd36                 vclmulh.vv v22, v12, v26
   496: 572c6e33                 vclmul.vv v24, v22, v28
   500: 572b6e37                 vclmulh.vv v22, v22, v28
   504: 572b6e33                 vclmul.vv v22, v22, v28
   508: 570a8a2f                 vxor.vv v20, v24, v20
   512: 570a6a2f                 vxor.vv v20, v22, v20
   516: 5702422f                 vxor.vv v4, v20, v4
   520: 130eb00c                 li t3, 0xcb
   524: 574d0e5e                 vmv.v.x v26, t3
   528: 572acd32                 vclmul.vv v20, v12, v26
   532: 572bcd36                 vclmulh.vv v22, v12, v26
   536: 572c6e33                 vclmul.vv v24, v22, v28
   540: 572b6e37                 vclmulh.vv v22, v22, v28
   544: 572b6e33                 vclmul.vv v22, v22, v28
   548: 570a8a2f                 vxor.vv v20, v24, v20
   552: 570a6a2f                 vxor.vv v20, v22, v20
   556: 5704442f                 vxor.vv v8, v20, v8
   560: 83b30500                 ld t2, 0(a1)
   564: b3836300                 add t2, t2, t1
   568: 27800302                 vse8.v v0, (t2)
   572: 83b38500                 ld t2, 8(a1)
   576: b3836300                 add t2, t2, t1
   580: 27820302                 vse8.v v4, (t2)
   584: 83b30501                 ld t2, 0x10(a1)
   588: b3836300                 add t2, t2, t1
   592: 27840302                 vse8.v v8, (t2)
   596: 33035300                 add t1, t1, t0
   600: 33065640                 sub a2, a2, t0
   604: 6ff05ff3                 j 0x190
   608: 67800000                 ret 
   612: 0000                     unimp 
   614: 0000                     unimp 
   616: 0000                     unimp 
   618: 0000                     unimp 
   620: 0000                     unimp 
   622: 0000                     unimp 
