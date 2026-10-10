; config_t1_1707 target=rvv twin=notail topo=raid6(8)/upd{0,1} op=parity_update U=8
; code_size=1184 twins=["notail"] uops_per_byte=1.5703125
     0: 63140600                 bnez a2, 8
     4: 6f008048                 j 0x48c
     8: 970f0000                 auipc t6, 0
    12: 938f8f49                 addi t6, t6, 0x498
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 130ef000                 li t3, 0xf
    24: 574f0e5e                 vmv.v.x v30, t3
    28: 13030000                 li t1, 0
    32: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    36: 93068000                 li a3, 8
    40: b386d202                 mul a3, t0, a3
    44: 3357d602                 divu a4, a2, a3
    48: 63140700                 bnez a4, 0x38
    52: 6f008045                 j 0x48c
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330502                 ld t2, 0x20(a0)
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
   132: 83338502                 ld t2, 0x28(a0)
   136: b3836300                 add t2, t2, t1
   140: 07840302                 vle8.v v8, (t2)
   144: b3835300                 add t2, t2, t0
   148: 87840302                 vle8.v v9, (t2)
   152: b3835300                 add t2, t2, t0
   156: 07850302                 vle8.v v10, (t2)
   160: b3835300                 add t2, t2, t0
   164: 87850302                 vle8.v v11, (t2)
   168: b3835300                 add t2, t2, t0
   172: 07860302                 vle8.v v12, (t2)
   176: b3835300                 add t2, t2, t0
   180: 87860302                 vle8.v v13, (t2)
   184: b3835300                 add t2, t2, t0
   188: 07870302                 vle8.v v14, (t2)
   192: b3835300                 add t2, t2, t0
   196: 87870302                 vle8.v v15, (t2)
   200: 83330500                 ld t2, 0(a0)
   204: b3836300                 add t2, t2, t1
   208: 833e0501                 ld t4, 0x10(a0)
   212: b38e6e00                 add t4, t4, t1
   216: 078c0302                 vle8.v v24, (t2)
   220: 878c0e02                 vle8.v v25, (t4)
   224: 57089c2f                 vxor.vv v16, v25, v24
   228: b3835300                 add t2, t2, t0
   232: b38e5e00                 add t4, t4, t0
   236: 078c0302                 vle8.v v24, (t2)
   240: 878c0e02                 vle8.v v25, (t4)
   244: d7089c2f                 vxor.vv v17, v25, v24
   248: b3835300                 add t2, t2, t0
   252: b38e5e00                 add t4, t4, t0
   256: 078c0302                 vle8.v v24, (t2)
   260: 878c0e02                 vle8.v v25, (t4)
   264: 57099c2f                 vxor.vv v18, v25, v24
   268: b3835300                 add t2, t2, t0
   272: b38e5e00                 add t4, t4, t0
   276: 078c0302                 vle8.v v24, (t2)
   280: 878c0e02                 vle8.v v25, (t4)
   284: d7099c2f                 vxor.vv v19, v25, v24
   288: b3835300                 add t2, t2, t0
   292: b38e5e00                 add t4, t4, t0
   296: 078c0302                 vle8.v v24, (t2)
   300: 878c0e02                 vle8.v v25, (t4)
   304: 570a9c2f                 vxor.vv v20, v25, v24
   308: b3835300                 add t2, t2, t0
   312: b38e5e00                 add t4, t4, t0
   316: 078c0302                 vle8.v v24, (t2)
   320: 878c0e02                 vle8.v v25, (t4)
   324: d70a9c2f                 vxor.vv v21, v25, v24
   328: b3835300                 add t2, t2, t0
   332: b38e5e00                 add t4, t4, t0
   336: 078c0302                 vle8.v v24, (t2)
   340: 878c0e02                 vle8.v v25, (t4)
   344: 570b9c2f                 vxor.vv v22, v25, v24
   348: b3835300                 add t2, t2, t0
   352: b38e5e00                 add t4, t4, t0
   356: 078c0302                 vle8.v v24, (t2)
   360: 878c0e02                 vle8.v v25, (t4)
   364: d70b9c2f                 vxor.vv v23, v25, v24
   368: 5700002f                 vxor.vv v0, v16, v0
   372: d780102f                 vxor.vv v1, v17, v1
   376: 5701212f                 vxor.vv v2, v18, v2
   380: d781312f                 vxor.vv v3, v19, v3
   384: 5702422f                 vxor.vv v4, v20, v4
   388: d782522f                 vxor.vv v5, v21, v5
   392: 5703632f                 vxor.vv v6, v22, v6
   396: d783732f                 vxor.vv v7, v23, v7
   400: 138e0f00                 mv t3, t6
   404: 070e0e02                 vle8.v v28, (t3)
   408: 138e0f01                 addi t3, t6, 0x10
   412: 870e0e02                 vle8.v v29, (t3)
   416: 570ce827                 vand.vv v24, v30, v16
   420: d73c02a3                 vsrl.vi v25, v16, 4
   424: 570dcc33                 vrgather.vv v26, v28, v24
   428: d78ddc33                 vrgather.vv v27, v29, v25
   432: 570dbd2f                 vxor.vv v26, v27, v26
   436: 5704a42f                 vxor.vv v8, v26, v8
   440: 578ce827                 vand.vv v24, v30, v17
   444: d73c12a3                 vsrl.vi v25, v17, 4
   448: 570dcc33                 vrgather.vv v26, v28, v24
   452: d78ddc33                 vrgather.vv v27, v29, v25
   456: 570dbd2f                 vxor.vv v26, v27, v26
   460: d784a42f                 vxor.vv v9, v26, v9
   464: 570ce927                 vand.vv v24, v30, v18
   468: d73c22a3                 vsrl.vi v25, v18, 4
   472: 570dcc33                 vrgather.vv v26, v28, v24
   476: d78ddc33                 vrgather.vv v27, v29, v25
   480: 570dbd2f                 vxor.vv v26, v27, v26
   484: 5705a52f                 vxor.vv v10, v26, v10
   488: 578ce927                 vand.vv v24, v30, v19
   492: d73c32a3                 vsrl.vi v25, v19, 4
   496: 570dcc33                 vrgather.vv v26, v28, v24
   500: d78ddc33                 vrgather.vv v27, v29, v25
   504: 570dbd2f                 vxor.vv v26, v27, v26
   508: d785a52f                 vxor.vv v11, v26, v11
   512: 570cea27                 vand.vv v24, v30, v20
   516: d73c42a3                 vsrl.vi v25, v20, 4
   520: 570dcc33                 vrgather.vv v26, v28, v24
   524: d78ddc33                 vrgather.vv v27, v29, v25
   528: 570dbd2f                 vxor.vv v26, v27, v26
   532: 5706a62f                 vxor.vv v12, v26, v12
   536: 578cea27                 vand.vv v24, v30, v21
   540: d73c52a3                 vsrl.vi v25, v21, 4
   544: 570dcc33                 vrgather.vv v26, v28, v24
   548: d78ddc33                 vrgather.vv v27, v29, v25
   552: 570dbd2f                 vxor.vv v26, v27, v26
   556: d786a62f                 vxor.vv v13, v26, v13
   560: 570ceb27                 vand.vv v24, v30, v22
   564: d73c62a3                 vsrl.vi v25, v22, 4
   568: 570dcc33                 vrgather.vv v26, v28, v24
   572: d78ddc33                 vrgather.vv v27, v29, v25
   576: 570dbd2f                 vxor.vv v26, v27, v26
   580: 5707a72f                 vxor.vv v14, v26, v14
   584: 578ceb27                 vand.vv v24, v30, v23
   588: d73c72a3                 vsrl.vi v25, v23, 4
   592: 570dcc33                 vrgather.vv v26, v28, v24
   596: d78ddc33                 vrgather.vv v27, v29, v25
   600: 570dbd2f                 vxor.vv v26, v27, v26
   604: d787a72f                 vxor.vv v15, v26, v15
   608: 83338500                 ld t2, 8(a0)
   612: b3836300                 add t2, t2, t1
   616: 833e8501                 ld t4, 0x18(a0)
   620: b38e6e00                 add t4, t4, t1
   624: 078c0302                 vle8.v v24, (t2)
   628: 878c0e02                 vle8.v v25, (t4)
   632: 57089c2f                 vxor.vv v16, v25, v24
   636: b3835300                 add t2, t2, t0
   640: b38e5e00                 add t4, t4, t0
   644: 078c0302                 vle8.v v24, (t2)
   648: 878c0e02                 vle8.v v25, (t4)
   652: d7089c2f                 vxor.vv v17, v25, v24
   656: b3835300                 add t2, t2, t0
   660: b38e5e00                 add t4, t4, t0
   664: 078c0302                 vle8.v v24, (t2)
   668: 878c0e02                 vle8.v v25, (t4)
   672: 57099c2f                 vxor.vv v18, v25, v24
   676: b3835300                 add t2, t2, t0
   680: b38e5e00                 add t4, t4, t0
   684: 078c0302                 vle8.v v24, (t2)
   688: 878c0e02                 vle8.v v25, (t4)
   692: d7099c2f                 vxor.vv v19, v25, v24
   696: b3835300                 add t2, t2, t0
   700: b38e5e00                 add t4, t4, t0
   704: 078c0302                 vle8.v v24, (t2)
   708: 878c0e02                 vle8.v v25, (t4)
   712: 570a9c2f                 vxor.vv v20, v25, v24
   716: b3835300                 add t2, t2, t0
   720: b38e5e00                 add t4, t4, t0
   724: 078c0302                 vle8.v v24, (t2)
   728: 878c0e02                 vle8.v v25, (t4)
   732: d70a9c2f                 vxor.vv v21, v25, v24
   736: b3835300                 add t2, t2, t0
   740: b38e5e00                 add t4, t4, t0
   744: 078c0302                 vle8.v v24, (t2)
   748: 878c0e02                 vle8.v v25, (t4)
   752: 570b9c2f                 vxor.vv v22, v25, v24
   756: b3835300                 add t2, t2, t0
   760: b38e5e00                 add t4, t4, t0
   764: 078c0302                 vle8.v v24, (t2)
   768: 878c0e02                 vle8.v v25, (t4)
   772: d70b9c2f                 vxor.vv v23, v25, v24
   776: 5700002f                 vxor.vv v0, v16, v0
   780: d780102f                 vxor.vv v1, v17, v1
   784: 5701212f                 vxor.vv v2, v18, v2
   788: d781312f                 vxor.vv v3, v19, v3
   792: 5702422f                 vxor.vv v4, v20, v4
   796: d782522f                 vxor.vv v5, v21, v5
   800: 5703632f                 vxor.vv v6, v22, v6
   804: d783732f                 vxor.vv v7, v23, v7
   808: 138e0f02                 addi t3, t6, 0x20
   812: 070e0e02                 vle8.v v28, (t3)
   816: 138e0f03                 addi t3, t6, 0x30
   820: 870e0e02                 vle8.v v29, (t3)
   824: 570ce827                 vand.vv v24, v30, v16
   828: d73c02a3                 vsrl.vi v25, v16, 4
   832: 570dcc33                 vrgather.vv v26, v28, v24
   836: d78ddc33                 vrgather.vv v27, v29, v25
   840: 570dbd2f                 vxor.vv v26, v27, v26
   844: 5704a42f                 vxor.vv v8, v26, v8
   848: 578ce827                 vand.vv v24, v30, v17
   852: d73c12a3                 vsrl.vi v25, v17, 4
   856: 570dcc33                 vrgather.vv v26, v28, v24
   860: d78ddc33                 vrgather.vv v27, v29, v25
   864: 570dbd2f                 vxor.vv v26, v27, v26
   868: d784a42f                 vxor.vv v9, v26, v9
   872: 570ce927                 vand.vv v24, v30, v18
   876: d73c22a3                 vsrl.vi v25, v18, 4
   880: 570dcc33                 vrgather.vv v26, v28, v24
   884: d78ddc33                 vrgather.vv v27, v29, v25
   888: 570dbd2f                 vxor.vv v26, v27, v26
   892: 5705a52f                 vxor.vv v10, v26, v10
   896: 578ce927                 vand.vv v24, v30, v19
   900: d73c32a3                 vsrl.vi v25, v19, 4
   904: 570dcc33                 vrgather.vv v26, v28, v24
   908: d78ddc33                 vrgather.vv v27, v29, v25
   912: 570dbd2f                 vxor.vv v26, v27, v26
   916: d785a52f                 vxor.vv v11, v26, v11
   920: 570cea27                 vand.vv v24, v30, v20
   924: d73c42a3                 vsrl.vi v25, v20, 4
   928: 570dcc33                 vrgather.vv v26, v28, v24
   932: d78ddc33                 vrgather.vv v27, v29, v25
   936: 570dbd2f                 vxor.vv v26, v27, v26
   940: 5706a62f                 vxor.vv v12, v26, v12
   944: 578cea27                 vand.vv v24, v30, v21
   948: d73c52a3                 vsrl.vi v25, v21, 4
   952: 570dcc33                 vrgather.vv v26, v28, v24
   956: d78ddc33                 vrgather.vv v27, v29, v25
   960: 570dbd2f                 vxor.vv v26, v27, v26
   964: d786a62f                 vxor.vv v13, v26, v13
   968: 570ceb27                 vand.vv v24, v30, v22
   972: d73c62a3                 vsrl.vi v25, v22, 4
   976: 570dcc33                 vrgather.vv v26, v28, v24
   980: d78ddc33                 vrgather.vv v27, v29, v25
   984: 570dbd2f                 vxor.vv v26, v27, v26
   988: 5707a72f                 vxor.vv v14, v26, v14
   992: 578ceb27                 vand.vv v24, v30, v23
   996: d73c72a3                 vsrl.vi v25, v23, 4
  1000: 570dcc33                 vrgather.vv v26, v28, v24
  1004: d78ddc33                 vrgather.vv v27, v29, v25
  1008: 570dbd2f                 vxor.vv v26, v27, v26
  1012: d787a72f                 vxor.vv v15, v26, v15
  1016: 83b30500                 ld t2, 0(a1)
  1020: b3836300                 add t2, t2, t1
  1024: 27800302                 vse8.v v0, (t2)
  1028: b3835300                 add t2, t2, t0
  1032: a7800302                 vse8.v v1, (t2)
  1036: b3835300                 add t2, t2, t0
  1040: 27810302                 vse8.v v2, (t2)
  1044: b3835300                 add t2, t2, t0
  1048: a7810302                 vse8.v v3, (t2)
  1052: b3835300                 add t2, t2, t0
  1056: 27820302                 vse8.v v4, (t2)
  1060: b3835300                 add t2, t2, t0
  1064: a7820302                 vse8.v v5, (t2)
  1068: b3835300                 add t2, t2, t0
  1072: 27830302                 vse8.v v6, (t2)
  1076: b3835300                 add t2, t2, t0
  1080: a7830302                 vse8.v v7, (t2)
  1084: 83b38500                 ld t2, 8(a1)
  1088: b3836300                 add t2, t2, t1
  1092: 27840302                 vse8.v v8, (t2)
  1096: b3835300                 add t2, t2, t0
  1100: a7840302                 vse8.v v9, (t2)
  1104: b3835300                 add t2, t2, t0
  1108: 27850302                 vse8.v v10, (t2)
  1112: b3835300                 add t2, t2, t0
  1116: a7850302                 vse8.v v11, (t2)
  1120: b3835300                 add t2, t2, t0
  1124: 27860302                 vse8.v v12, (t2)
  1128: b3835300                 add t2, t2, t0
  1132: a7860302                 vse8.v v13, (t2)
  1136: b3835300                 add t2, t2, t0
  1140: 27870302                 vse8.v v14, (t2)
  1144: b3835300                 add t2, t2, t0
  1148: a7870302                 vse8.v v15, (t2)
  1152: 3303d300                 add t1, t1, a3
  1156: 1307f7ff                 addi a4, a4, -1
  1160: e31c07ba                 bnez a4, 0x40
  1164: 67800000                 ret 
  1168: 0000                     unimp 
  1170: 0000                     unimp 
  1172: 0000                     unimp 
  1174: 0000                     unimp 
  1176: 0000                     unimp 
  1178: 0000                     unimp 
  1180: 0000                     unimp 
  1182: 0000                     unimp 
