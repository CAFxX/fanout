; config_t1_1703 target=rvv twin=baseline topo=raid6(8)/upd{0,1} op=parity_update U=8
; code_size=784 twins=[] uops_per_byte=1.25
     0: 63140600                 bnez a2, 8
     4: 6f00c02f                 j 0x300
     8: 970f0000                 auipc t6, 0
    12: 938f8f30                 addi t6, t6, 0x308
    16: 577f100c                 vsetvli t5, zero, e8, m2, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574f0e5e                 vmv.v.x v30, t3
    28: 13030000                 li t1, 0
    32: d772100c                 vsetvli t0, zero, e8, m2, ta, ma
    36: 93063000                 li a3, 3
    40: b386d202                 mul a3, t0, a3
    44: 3357d602                 divu a4, a2, a3
    48: 63140700                 bnez a4, 0x38
    52: 6f00c01d                 j 0x210
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330502                 ld t2, 0x20(a0)
    68: b3836300                 add t2, t2, t1
    72: 07800302                 vle8.v v0, (t2)
    76: b3835300                 add t2, t2, t0
    80: 07810302                 vle8.v v2, (t2)
    84: b3835300                 add t2, t2, t0
    88: 07820302                 vle8.v v4, (t2)
    92: 83338502                 ld t2, 0x28(a0)
    96: b3836300                 add t2, t2, t1
   100: 07830302                 vle8.v v6, (t2)
   104: b3835300                 add t2, t2, t0
   108: 07840302                 vle8.v v8, (t2)
   112: b3835300                 add t2, t2, t0
   116: 07850302                 vle8.v v10, (t2)
   120: 83330500                 ld t2, 0(a0)
   124: b3836300                 add t2, t2, t1
   128: 833e0501                 ld t4, 0x10(a0)
   132: b38e6e00                 add t4, t4, t1
   136: 07890302                 vle8.v v18, (t2)
   140: 078a0e02                 vle8.v v20, (t4)
   144: 5706492f                 vxor.vv v12, v20, v18
   148: b3835300                 add t2, t2, t0
   152: b38e5e00                 add t4, t4, t0
   156: 07890302                 vle8.v v18, (t2)
   160: 078a0e02                 vle8.v v20, (t4)
   164: 5707492f                 vxor.vv v14, v20, v18
   168: b3835300                 add t2, t2, t0
   172: b38e5e00                 add t4, t4, t0
   176: 07890302                 vle8.v v18, (t2)
   180: 078a0e02                 vle8.v v20, (t4)
   184: 5708492f                 vxor.vv v16, v20, v18
   188: 5700c02e                 vxor.vv v0, v12, v0
   192: 5701e12e                 vxor.vv v2, v14, v2
   196: 5702022f                 vxor.vv v4, v16, v4
   200: 138e0f00                 mv t3, t6
   204: 070d0e02                 vle8.v v26, (t3)
   208: 138e0f02                 addi t3, t6, 0x20
   212: 070e0e02                 vle8.v v28, (t3)
   216: 5709e627                 vand.vv v18, v30, v12
   220: 573ac2a2                 vsrl.vi v20, v12, 4
   224: 570ba933                 vrgather.vv v22, v26, v18
   228: 570cca33                 vrgather.vv v24, v28, v20
   232: 570b8b2f                 vxor.vv v22, v24, v22
   236: 5703632f                 vxor.vv v6, v22, v6
   240: 5709e727                 vand.vv v18, v30, v14
   244: 573ae2a2                 vsrl.vi v20, v14, 4
   248: 570ba933                 vrgather.vv v22, v26, v18
   252: 570cca33                 vrgather.vv v24, v28, v20
   256: 570b8b2f                 vxor.vv v22, v24, v22
   260: 5704642f                 vxor.vv v8, v22, v8
   264: 5709e827                 vand.vv v18, v30, v16
   268: 573a02a3                 vsrl.vi v20, v16, 4
   272: 570ba933                 vrgather.vv v22, v26, v18
   276: 570cca33                 vrgather.vv v24, v28, v20
   280: 570b8b2f                 vxor.vv v22, v24, v22
   284: 5705652f                 vxor.vv v10, v22, v10
   288: 83338500                 ld t2, 8(a0)
   292: b3836300                 add t2, t2, t1
   296: 833e8501                 ld t4, 0x18(a0)
   300: b38e6e00                 add t4, t4, t1
   304: 07890302                 vle8.v v18, (t2)
   308: 078a0e02                 vle8.v v20, (t4)
   312: 5706492f                 vxor.vv v12, v20, v18
   316: b3835300                 add t2, t2, t0
   320: b38e5e00                 add t4, t4, t0
   324: 07890302                 vle8.v v18, (t2)
   328: 078a0e02                 vle8.v v20, (t4)
   332: 5707492f                 vxor.vv v14, v20, v18
   336: b3835300                 add t2, t2, t0
   340: b38e5e00                 add t4, t4, t0
   344: 07890302                 vle8.v v18, (t2)
   348: 078a0e02                 vle8.v v20, (t4)
   352: 5708492f                 vxor.vv v16, v20, v18
   356: 5700c02e                 vxor.vv v0, v12, v0
   360: 5701e12e                 vxor.vv v2, v14, v2
   364: 5702022f                 vxor.vv v4, v16, v4
   368: 138e0f04                 addi t3, t6, 0x40
   372: 070d0e02                 vle8.v v26, (t3)
   376: 138e0f06                 addi t3, t6, 0x60
   380: 070e0e02                 vle8.v v28, (t3)
   384: 5709e627                 vand.vv v18, v30, v12
   388: 573ac2a2                 vsrl.vi v20, v12, 4
   392: 570ba933                 vrgather.vv v22, v26, v18
   396: 570cca33                 vrgather.vv v24, v28, v20
   400: 570b8b2f                 vxor.vv v22, v24, v22
   404: 5703632f                 vxor.vv v6, v22, v6
   408: 5709e727                 vand.vv v18, v30, v14
   412: 573ae2a2                 vsrl.vi v20, v14, 4
   416: 570ba933                 vrgather.vv v22, v26, v18
   420: 570cca33                 vrgather.vv v24, v28, v20
   424: 570b8b2f                 vxor.vv v22, v24, v22
   428: 5704642f                 vxor.vv v8, v22, v8
   432: 5709e827                 vand.vv v18, v30, v16
   436: 573a02a3                 vsrl.vi v20, v16, 4
   440: 570ba933                 vrgather.vv v22, v26, v18
   444: 570cca33                 vrgather.vv v24, v28, v20
   448: 570b8b2f                 vxor.vv v22, v24, v22
   452: 5705652f                 vxor.vv v10, v22, v10
   456: 83b30500                 ld t2, 0(a1)
   460: b3836300                 add t2, t2, t1
   464: 27800302                 vse8.v v0, (t2)
   468: b3835300                 add t2, t2, t0
   472: 27810302                 vse8.v v2, (t2)
   476: b3835300                 add t2, t2, t0
   480: 27820302                 vse8.v v4, (t2)
   484: 83b38500                 ld t2, 8(a1)
   488: b3836300                 add t2, t2, t1
   492: 27830302                 vse8.v v6, (t2)
   496: b3835300                 add t2, t2, t0
   500: 27840302                 vse8.v v8, (t2)
   504: b3835300                 add t2, t2, t0
   508: 27850302                 vse8.v v10, (t2)
   512: 3303d300                 add t1, t1, a3
   516: 1307f7ff                 addi a4, a4, -1
   520: e31c07e2                 bnez a4, 0x40
   524: 33066640                 sub a2, a2, t1
   528: d772160c                 vsetvli t0, a2, e8, m2, ta, ma
   532: 63940200                 bnez t0, 0x21c
   536: 6f00800e                 j 0x300
   540: 83330502                 ld t2, 0x20(a0)
   544: b3836300                 add t2, t2, t1
   548: 07800302                 vle8.v v0, (t2)
   552: 83338502                 ld t2, 0x28(a0)
   556: b3836300                 add t2, t2, t1
   560: 07830302                 vle8.v v6, (t2)
   564: 83330500                 ld t2, 0(a0)
   568: b3836300                 add t2, t2, t1
   572: 833e0501                 ld t4, 0x10(a0)
   576: b38e6e00                 add t4, t4, t1
   580: 07890302                 vle8.v v18, (t2)
   584: 078a0e02                 vle8.v v20, (t4)
   588: 5706492f                 vxor.vv v12, v20, v18
   592: 5700c02e                 vxor.vv v0, v12, v0
   596: 138f0200                 mv t5, t0
   600: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
   604: 138e0f00                 mv t3, t6
   608: 070d0e02                 vle8.v v26, (t3)
   612: 138e0f02                 addi t3, t6, 0x20
   616: 070e0e02                 vle8.v v28, (t3)
   620: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
   624: 5709e627                 vand.vv v18, v30, v12
   628: 573ac2a2                 vsrl.vi v20, v12, 4
   632: 570ba933                 vrgather.vv v22, v26, v18
   636: 570cca33                 vrgather.vv v24, v28, v20
   640: 570b8b2f                 vxor.vv v22, v24, v22
   644: 5703632f                 vxor.vv v6, v22, v6
   648: 83338500                 ld t2, 8(a0)
   652: b3836300                 add t2, t2, t1
   656: 833e8501                 ld t4, 0x18(a0)
   660: b38e6e00                 add t4, t4, t1
   664: 07890302                 vle8.v v18, (t2)
   668: 078a0e02                 vle8.v v20, (t4)
   672: 5706492f                 vxor.vv v12, v20, v18
   676: 5700c02e                 vxor.vv v0, v12, v0
   680: 138f0200                 mv t5, t0
   684: 577018cc                 vsetivli zero, 0x10, e8, m2, ta, ma
   688: 138e0f04                 addi t3, t6, 0x40
   692: 070d0e02                 vle8.v v26, (t3)
   696: 138e0f06                 addi t3, t6, 0x60
   700: 070e0e02                 vle8.v v28, (t3)
   704: 57701f0c                 vsetvli zero, t5, e8, m2, ta, ma
   708: 5709e627                 vand.vv v18, v30, v12
   712: 573ac2a2                 vsrl.vi v20, v12, 4
   716: 570ba933                 vrgather.vv v22, v26, v18
   720: 570cca33                 vrgather.vv v24, v28, v20
   724: 570b8b2f                 vxor.vv v22, v24, v22
   728: 5703632f                 vxor.vv v6, v22, v6
   732: 83b30500                 ld t2, 0(a1)
   736: b3836300                 add t2, t2, t1
   740: 27800302                 vse8.v v0, (t2)
   744: 83b38500                 ld t2, 8(a1)
   748: b3836300                 add t2, t2, t1
   752: 27830302                 vse8.v v6, (t2)
   756: 33035300                 add t1, t1, t0
   760: 33065640                 sub a2, a2, t0
   764: 6ff05ff1                 j 0x210
   768: 67800000                 ret 
   772: 0000                     unimp 
   774: 0000                     unimp 
   776: 0000                     unimp 
   778: 0000                     unimp 
   780: 0000                     unimp 
   782: 0000                     unimp 
