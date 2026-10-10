; config_t1_1768 target=rvv-zvbc32e twin=baseline topo=rdp(7)/enc op=encode U=1
; code_size=2048 twins=[] uops_per_byte=9.125
     0: 63140600                 bnez a2, 8
     4: 6f00c07e                 j 0x7f0
     8: 970f0000                 auipc t6, 0
    12: 938f8f7f                 addi t6, t6, 0x7f8
    16: 577f100c                 vsetvli t5, zero, e8, m2, ta, ma
    20: 13030000                 li t1, 0
    24: d772100c                 vsetvli t0, zero, e8, m2, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f00c03e                 j 0x414
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 078c0302                 vle8.v v24, (t2)
    76: 57000c5e                 vmv.v.v v0, v24
    80: 57060c5e                 vmv.v.v v12, v24
    84: 83338500                 ld t2, 8(a0)
    88: b3836300                 add t2, t2, t1
    92: 078c0302                 vle8.v v24, (t2)
    96: 57010c5e                 vmv.v.v v2, v24
   100: 5706862f                 vxor.vv v12, v24, v12
   104: 57070c5e                 vmv.v.v v14, v24
   108: 83330501                 ld t2, 0x10(a0)
   112: b3836300                 add t2, t2, t1
   116: 078c0302                 vle8.v v24, (t2)
   120: 57020c5e                 vmv.v.v v4, v24
   124: 5707872f                 vxor.vv v14, v24, v14
   128: 57080c5e                 vmv.v.v v16, v24
   132: 83338501                 ld t2, 0x18(a0)
   136: b3836300                 add t2, t2, t1
   140: 078c0302                 vle8.v v24, (t2)
   144: 57030c5e                 vmv.v.v v6, v24
   148: 5708882f                 vxor.vv v16, v24, v16
   152: 57090c5e                 vmv.v.v v18, v24
   156: 83330502                 ld t2, 0x20(a0)
   160: b3836300                 add t2, t2, t1
   164: 078c0302                 vle8.v v24, (t2)
   168: 57040c5e                 vmv.v.v v8, v24
   172: 5709892f                 vxor.vv v18, v24, v18
   176: 570a0c5e                 vmv.v.v v20, v24
   180: 83338502                 ld t2, 0x28(a0)
   184: b3836300                 add t2, t2, t1
   188: 078c0302                 vle8.v v24, (t2)
   192: 57050c5e                 vmv.v.v v10, v24
   196: 570a8a2f                 vxor.vv v20, v24, v20
   200: 570b0c5e                 vmv.v.v v22, v24
   204: 83330503                 ld t2, 0x30(a0)
   208: b3836300                 add t2, t2, t1
   212: 078c0302                 vle8.v v24, (t2)
   216: 5700802f                 vxor.vv v0, v24, v0
   220: 5707872f                 vxor.vv v14, v24, v14
   224: 83338503                 ld t2, 0x38(a0)
   228: b3836300                 add t2, t2, t1
   232: 078c0302                 vle8.v v24, (t2)
   236: 5701812f                 vxor.vv v2, v24, v2
   240: 5706862f                 vxor.vv v12, v24, v12
   244: 5708882f                 vxor.vv v16, v24, v16
   248: 83330504                 ld t2, 0x40(a0)
   252: b3836300                 add t2, t2, t1
   256: 078c0302                 vle8.v v24, (t2)
   260: 5702822f                 vxor.vv v4, v24, v4
   264: 5707872f                 vxor.vv v14, v24, v14
   268: 5709892f                 vxor.vv v18, v24, v18
   272: 83338504                 ld t2, 0x48(a0)
   276: b3836300                 add t2, t2, t1
   280: 078c0302                 vle8.v v24, (t2)
   284: 5703832f                 vxor.vv v6, v24, v6
   288: 5708882f                 vxor.vv v16, v24, v16
   292: 570a8a2f                 vxor.vv v20, v24, v20
   296: 83330505                 ld t2, 0x50(a0)
   300: b3836300                 add t2, t2, t1
   304: 078c0302                 vle8.v v24, (t2)
   308: 5704842f                 vxor.vv v8, v24, v8
   312: 5709892f                 vxor.vv v18, v24, v18
   316: 570b8b2f                 vxor.vv v22, v24, v22
   320: 83338505                 ld t2, 0x58(a0)
   324: b3836300                 add t2, t2, t1
   328: 078c0302                 vle8.v v24, (t2)
   332: 5705852f                 vxor.vv v10, v24, v10
   336: 570a8a2f                 vxor.vv v20, v24, v20
   340: 83330506                 ld t2, 0x60(a0)
   344: b3836300                 add t2, t2, t1
   348: 078c0302                 vle8.v v24, (t2)
   352: 5700802f                 vxor.vv v0, v24, v0
   356: 5708882f                 vxor.vv v16, v24, v16
   360: 83338506                 ld t2, 0x68(a0)
   364: b3836300                 add t2, t2, t1
   368: 078c0302                 vle8.v v24, (t2)
   372: 5701812f                 vxor.vv v2, v24, v2
   376: 5706862f                 vxor.vv v12, v24, v12
   380: 5709892f                 vxor.vv v18, v24, v18
   384: 83330507                 ld t2, 0x70(a0)
   388: b3836300                 add t2, t2, t1
   392: 078c0302                 vle8.v v24, (t2)
   396: 5702822f                 vxor.vv v4, v24, v4
   400: 5707872f                 vxor.vv v14, v24, v14
   404: 570a8a2f                 vxor.vv v20, v24, v20
   408: 83338507                 ld t2, 0x78(a0)
   412: b3836300                 add t2, t2, t1
   416: 078c0302                 vle8.v v24, (t2)
   420: 5703832f                 vxor.vv v6, v24, v6
   424: 5708882f                 vxor.vv v16, v24, v16
   428: 570b8b2f                 vxor.vv v22, v24, v22
   432: 83330508                 ld t2, 0x80(a0)
   436: b3836300                 add t2, t2, t1
   440: 078c0302                 vle8.v v24, (t2)
   444: 5704842f                 vxor.vv v8, v24, v8
   448: 5709892f                 vxor.vv v18, v24, v18
   452: 83338508                 ld t2, 0x88(a0)
   456: b3836300                 add t2, t2, t1
   460: 078c0302                 vle8.v v24, (t2)
   464: 5705852f                 vxor.vv v10, v24, v10
   468: 5706862f                 vxor.vv v12, v24, v12
   472: 570a8a2f                 vxor.vv v20, v24, v20
   476: 83330509                 ld t2, 0x90(a0)
   480: b3836300                 add t2, t2, t1
   484: 078c0302                 vle8.v v24, (t2)
   488: 5700802f                 vxor.vv v0, v24, v0
   492: 5709892f                 vxor.vv v18, v24, v18
   496: 83338509                 ld t2, 0x98(a0)
   500: b3836300                 add t2, t2, t1
   504: 078c0302                 vle8.v v24, (t2)
   508: 5701812f                 vxor.vv v2, v24, v2
   512: 5706862f                 vxor.vv v12, v24, v12
   516: 570a8a2f                 vxor.vv v20, v24, v20
   520: 8333050a                 ld t2, 0xa0(a0)
   524: b3836300                 add t2, t2, t1
   528: 078c0302                 vle8.v v24, (t2)
   532: 5702822f                 vxor.vv v4, v24, v4
   536: 5707872f                 vxor.vv v14, v24, v14
   540: 570b8b2f                 vxor.vv v22, v24, v22
   544: 8333850a                 ld t2, 0xa8(a0)
   548: b3836300                 add t2, t2, t1
   552: 078c0302                 vle8.v v24, (t2)
   556: 5703832f                 vxor.vv v6, v24, v6
   560: 5708882f                 vxor.vv v16, v24, v16
   564: 8333050b                 ld t2, 0xb0(a0)
   568: b3836300                 add t2, t2, t1
   572: 078c0302                 vle8.v v24, (t2)
   576: 5704842f                 vxor.vv v8, v24, v8
   580: 5706862f                 vxor.vv v12, v24, v12
   584: 5709892f                 vxor.vv v18, v24, v18
   588: 8333850b                 ld t2, 0xb8(a0)
   592: b3836300                 add t2, t2, t1
   596: 078c0302                 vle8.v v24, (t2)
   600: 5705852f                 vxor.vv v10, v24, v10
   604: 5707872f                 vxor.vv v14, v24, v14
   608: 570a8a2f                 vxor.vv v20, v24, v20
   612: 8333050c                 ld t2, 0xc0(a0)
   616: b3836300                 add t2, t2, t1
   620: 078c0302                 vle8.v v24, (t2)
   624: 5700802f                 vxor.vv v0, v24, v0
   628: 570a8a2f                 vxor.vv v20, v24, v20
   632: 8333850c                 ld t2, 0xc8(a0)
   636: b3836300                 add t2, t2, t1
   640: 078c0302                 vle8.v v24, (t2)
   644: 5701812f                 vxor.vv v2, v24, v2
   648: 5706862f                 vxor.vv v12, v24, v12
   652: 570b8b2f                 vxor.vv v22, v24, v22
   656: 8333050d                 ld t2, 0xd0(a0)
   660: b3836300                 add t2, t2, t1
   664: 078c0302                 vle8.v v24, (t2)
   668: 5702822f                 vxor.vv v4, v24, v4
   672: 5707872f                 vxor.vv v14, v24, v14
   676: 8333850d                 ld t2, 0xd8(a0)
   680: b3836300                 add t2, t2, t1
   684: 078c0302                 vle8.v v24, (t2)
   688: 5703832f                 vxor.vv v6, v24, v6
   692: 5706862f                 vxor.vv v12, v24, v12
   696: 5708882f                 vxor.vv v16, v24, v16
   700: 8333050e                 ld t2, 0xe0(a0)
   704: b3836300                 add t2, t2, t1
   708: 078c0302                 vle8.v v24, (t2)
   712: 5704842f                 vxor.vv v8, v24, v8
   716: 5707872f                 vxor.vv v14, v24, v14
   720: 5709892f                 vxor.vv v18, v24, v18
   724: 8333850e                 ld t2, 0xe8(a0)
   728: b3836300                 add t2, t2, t1
   732: 078c0302                 vle8.v v24, (t2)
   736: 5705852f                 vxor.vv v10, v24, v10
   740: 5708882f                 vxor.vv v16, v24, v16
   744: 570a8a2f                 vxor.vv v20, v24, v20
   748: 8333050f                 ld t2, 0xf0(a0)
   752: b3836300                 add t2, t2, t1
   756: 078c0302                 vle8.v v24, (t2)
   760: 5700802f                 vxor.vv v0, v24, v0
   764: 570b8b2f                 vxor.vv v22, v24, v22
   768: 8333850f                 ld t2, 0xf8(a0)
   772: b3836300                 add t2, t2, t1
   776: 078c0302                 vle8.v v24, (t2)
   780: 5701812f                 vxor.vv v2, v24, v2
   784: 5706862f                 vxor.vv v12, v24, v12
   788: 83330510                 ld t2, 0x100(a0)
   792: b3836300                 add t2, t2, t1
   796: 078c0302                 vle8.v v24, (t2)
   800: 5702822f                 vxor.vv v4, v24, v4
   804: 5706862f                 vxor.vv v12, v24, v12
   808: 5707872f                 vxor.vv v14, v24, v14
   812: 83338510                 ld t2, 0x108(a0)
   816: b3836300                 add t2, t2, t1
   820: 078c0302                 vle8.v v24, (t2)
   824: 5703832f                 vxor.vv v6, v24, v6
   828: 5707872f                 vxor.vv v14, v24, v14
   832: 5708882f                 vxor.vv v16, v24, v16
   836: 83330511                 ld t2, 0x110(a0)
   840: b3836300                 add t2, t2, t1
   844: 078c0302                 vle8.v v24, (t2)
   848: 5704842f                 vxor.vv v8, v24, v8
   852: 5708882f                 vxor.vv v16, v24, v16
   856: 5709892f                 vxor.vv v18, v24, v18
   860: 83338511                 ld t2, 0x118(a0)
   864: b3836300                 add t2, t2, t1
   868: 078c0302                 vle8.v v24, (t2)
   872: 5705852f                 vxor.vv v10, v24, v10
   876: 5709892f                 vxor.vv v18, v24, v18
   880: 570a8a2f                 vxor.vv v20, v24, v20
   884: 83b30500                 ld t2, 0(a1)
   888: b3836300                 add t2, t2, t1
   892: 27800302                 vse8.v v0, (t2)
   896: 83b38500                 ld t2, 8(a1)
   900: b3836300                 add t2, t2, t1
   904: 27810302                 vse8.v v2, (t2)
   908: 83b30501                 ld t2, 0x10(a1)
   912: b3836300                 add t2, t2, t1
   916: 27820302                 vse8.v v4, (t2)
   920: 83b38501                 ld t2, 0x18(a1)
   924: b3836300                 add t2, t2, t1
   928: 27830302                 vse8.v v6, (t2)
   932: 83b30502                 ld t2, 0x20(a1)
   936: b3836300                 add t2, t2, t1
   940: 27840302                 vse8.v v8, (t2)
   944: 83b38502                 ld t2, 0x28(a1)
   948: b3836300                 add t2, t2, t1
   952: 27850302                 vse8.v v10, (t2)
   956: 83b30503                 ld t2, 0x30(a1)
   960: b3836300                 add t2, t2, t1
   964: 27860302                 vse8.v v12, (t2)
   968: 83b38503                 ld t2, 0x38(a1)
   972: b3836300                 add t2, t2, t1
   976: 27870302                 vse8.v v14, (t2)
   980: 83b30504                 ld t2, 0x40(a1)
   984: b3836300                 add t2, t2, t1
   988: 27880302                 vse8.v v16, (t2)
   992: 83b38504                 ld t2, 0x48(a1)
   996: b3836300                 add t2, t2, t1
  1000: 27890302                 vse8.v v18, (t2)
  1004: 83b30505                 ld t2, 0x50(a1)
  1008: b3836300                 add t2, t2, t1
  1012: 278a0302                 vse8.v v20, (t2)
  1016: 83b38505                 ld t2, 0x58(a1)
  1020: b3836300                 add t2, t2, t1
  1024: 278b0302                 vse8.v v22, (t2)
  1028: 3303d300                 add t1, t1, a3
  1032: 1307f7ff                 addi a4, a4, -1
  1036: e31a07c2                 bnez a4, 0x40
  1040: 33066640                 sub a2, a2, t1
  1044: d772160c                 vsetvli t0, a2, e8, m2, ta, ma
  1048: 63940200                 bnez t0, 0x420
  1052: 6f00403d                 j 0x7f0
  1056: 83330500                 ld t2, 0(a0)
  1060: b3836300                 add t2, t2, t1
  1064: 078c0302                 vle8.v v24, (t2)
  1068: 57000c5e                 vmv.v.v v0, v24
  1072: 57060c5e                 vmv.v.v v12, v24
  1076: 83338500                 ld t2, 8(a0)
  1080: b3836300                 add t2, t2, t1
  1084: 078c0302                 vle8.v v24, (t2)
  1088: 57010c5e                 vmv.v.v v2, v24
  1092: 5706862f                 vxor.vv v12, v24, v12
  1096: 57070c5e                 vmv.v.v v14, v24
  1100: 83330501                 ld t2, 0x10(a0)
  1104: b3836300                 add t2, t2, t1
  1108: 078c0302                 vle8.v v24, (t2)
  1112: 57020c5e                 vmv.v.v v4, v24
  1116: 5707872f                 vxor.vv v14, v24, v14
  1120: 57080c5e                 vmv.v.v v16, v24
  1124: 83338501                 ld t2, 0x18(a0)
  1128: b3836300                 add t2, t2, t1
  1132: 078c0302                 vle8.v v24, (t2)
  1136: 57030c5e                 vmv.v.v v6, v24
  1140: 5708882f                 vxor.vv v16, v24, v16
  1144: 57090c5e                 vmv.v.v v18, v24
  1148: 83330502                 ld t2, 0x20(a0)
  1152: b3836300                 add t2, t2, t1
  1156: 078c0302                 vle8.v v24, (t2)
  1160: 57040c5e                 vmv.v.v v8, v24
  1164: 5709892f                 vxor.vv v18, v24, v18
  1168: 570a0c5e                 vmv.v.v v20, v24
  1172: 83338502                 ld t2, 0x28(a0)
  1176: b3836300                 add t2, t2, t1
  1180: 078c0302                 vle8.v v24, (t2)
  1184: 57050c5e                 vmv.v.v v10, v24
  1188: 570a8a2f                 vxor.vv v20, v24, v20
  1192: 570b0c5e                 vmv.v.v v22, v24
  1196: 83330503                 ld t2, 0x30(a0)
  1200: b3836300                 add t2, t2, t1
  1204: 078c0302                 vle8.v v24, (t2)
  1208: 5700802f                 vxor.vv v0, v24, v0
  1212: 5707872f                 vxor.vv v14, v24, v14
  1216: 83338503                 ld t2, 0x38(a0)
  1220: b3836300                 add t2, t2, t1
  1224: 078c0302                 vle8.v v24, (t2)
  1228: 5701812f                 vxor.vv v2, v24, v2
  1232: 5706862f                 vxor.vv v12, v24, v12
  1236: 5708882f                 vxor.vv v16, v24, v16
  1240: 83330504                 ld t2, 0x40(a0)
  1244: b3836300                 add t2, t2, t1
  1248: 078c0302                 vle8.v v24, (t2)
  1252: 5702822f                 vxor.vv v4, v24, v4
  1256: 5707872f                 vxor.vv v14, v24, v14
  1260: 5709892f                 vxor.vv v18, v24, v18
  1264: 83338504                 ld t2, 0x48(a0)
  1268: b3836300                 add t2, t2, t1
  1272: 078c0302                 vle8.v v24, (t2)
  1276: 5703832f                 vxor.vv v6, v24, v6
  1280: 5708882f                 vxor.vv v16, v24, v16
  1284: 570a8a2f                 vxor.vv v20, v24, v20
  1288: 83330505                 ld t2, 0x50(a0)
  1292: b3836300                 add t2, t2, t1
  1296: 078c0302                 vle8.v v24, (t2)
  1300: 5704842f                 vxor.vv v8, v24, v8
  1304: 5709892f                 vxor.vv v18, v24, v18
  1308: 570b8b2f                 vxor.vv v22, v24, v22
  1312: 83338505                 ld t2, 0x58(a0)
  1316: b3836300                 add t2, t2, t1
  1320: 078c0302                 vle8.v v24, (t2)
  1324: 5705852f                 vxor.vv v10, v24, v10
  1328: 570a8a2f                 vxor.vv v20, v24, v20
  1332: 83330506                 ld t2, 0x60(a0)
  1336: b3836300                 add t2, t2, t1
  1340: 078c0302                 vle8.v v24, (t2)
  1344: 5700802f                 vxor.vv v0, v24, v0
  1348: 5708882f                 vxor.vv v16, v24, v16
  1352: 83338506                 ld t2, 0x68(a0)
  1356: b3836300                 add t2, t2, t1
  1360: 078c0302                 vle8.v v24, (t2)
  1364: 5701812f                 vxor.vv v2, v24, v2
  1368: 5706862f                 vxor.vv v12, v24, v12
  1372: 5709892f                 vxor.vv v18, v24, v18
  1376: 83330507                 ld t2, 0x70(a0)
  1380: b3836300                 add t2, t2, t1
  1384: 078c0302                 vle8.v v24, (t2)
  1388: 5702822f                 vxor.vv v4, v24, v4
  1392: 5707872f                 vxor.vv v14, v24, v14
  1396: 570a8a2f                 vxor.vv v20, v24, v20
  1400: 83338507                 ld t2, 0x78(a0)
  1404: b3836300                 add t2, t2, t1
  1408: 078c0302                 vle8.v v24, (t2)
  1412: 5703832f                 vxor.vv v6, v24, v6
  1416: 5708882f                 vxor.vv v16, v24, v16
  1420: 570b8b2f                 vxor.vv v22, v24, v22
  1424: 83330508                 ld t2, 0x80(a0)
  1428: b3836300                 add t2, t2, t1
  1432: 078c0302                 vle8.v v24, (t2)
  1436: 5704842f                 vxor.vv v8, v24, v8
  1440: 5709892f                 vxor.vv v18, v24, v18
  1444: 83338508                 ld t2, 0x88(a0)
  1448: b3836300                 add t2, t2, t1
  1452: 078c0302                 vle8.v v24, (t2)
  1456: 5705852f                 vxor.vv v10, v24, v10
  1460: 5706862f                 vxor.vv v12, v24, v12
  1464: 570a8a2f                 vxor.vv v20, v24, v20
  1468: 83330509                 ld t2, 0x90(a0)
  1472: b3836300                 add t2, t2, t1
  1476: 078c0302                 vle8.v v24, (t2)
  1480: 5700802f                 vxor.vv v0, v24, v0
  1484: 5709892f                 vxor.vv v18, v24, v18
  1488: 83338509                 ld t2, 0x98(a0)
  1492: b3836300                 add t2, t2, t1
  1496: 078c0302                 vle8.v v24, (t2)
  1500: 5701812f                 vxor.vv v2, v24, v2
  1504: 5706862f                 vxor.vv v12, v24, v12
  1508: 570a8a2f                 vxor.vv v20, v24, v20
  1512: 8333050a                 ld t2, 0xa0(a0)
  1516: b3836300                 add t2, t2, t1
  1520: 078c0302                 vle8.v v24, (t2)
  1524: 5702822f                 vxor.vv v4, v24, v4
  1528: 5707872f                 vxor.vv v14, v24, v14
  1532: 570b8b2f                 vxor.vv v22, v24, v22
  1536: 8333850a                 ld t2, 0xa8(a0)
  1540: b3836300                 add t2, t2, t1
  1544: 078c0302                 vle8.v v24, (t2)
  1548: 5703832f                 vxor.vv v6, v24, v6
  1552: 5708882f                 vxor.vv v16, v24, v16
  1556: 8333050b                 ld t2, 0xb0(a0)
  1560: b3836300                 add t2, t2, t1
  1564: 078c0302                 vle8.v v24, (t2)
  1568: 5704842f                 vxor.vv v8, v24, v8
  1572: 5706862f                 vxor.vv v12, v24, v12
  1576: 5709892f                 vxor.vv v18, v24, v18
  1580: 8333850b                 ld t2, 0xb8(a0)
  1584: b3836300                 add t2, t2, t1
  1588: 078c0302                 vle8.v v24, (t2)
  1592: 5705852f                 vxor.vv v10, v24, v10
  1596: 5707872f                 vxor.vv v14, v24, v14
  1600: 570a8a2f                 vxor.vv v20, v24, v20
  1604: 8333050c                 ld t2, 0xc0(a0)
  1608: b3836300                 add t2, t2, t1
  1612: 078c0302                 vle8.v v24, (t2)
  1616: 5700802f                 vxor.vv v0, v24, v0
  1620: 570a8a2f                 vxor.vv v20, v24, v20
  1624: 8333850c                 ld t2, 0xc8(a0)
  1628: b3836300                 add t2, t2, t1
  1632: 078c0302                 vle8.v v24, (t2)
  1636: 5701812f                 vxor.vv v2, v24, v2
  1640: 5706862f                 vxor.vv v12, v24, v12
  1644: 570b8b2f                 vxor.vv v22, v24, v22
  1648: 8333050d                 ld t2, 0xd0(a0)
  1652: b3836300                 add t2, t2, t1
  1656: 078c0302                 vle8.v v24, (t2)
  1660: 5702822f                 vxor.vv v4, v24, v4
  1664: 5707872f                 vxor.vv v14, v24, v14
  1668: 8333850d                 ld t2, 0xd8(a0)
  1672: b3836300                 add t2, t2, t1
  1676: 078c0302                 vle8.v v24, (t2)
  1680: 5703832f                 vxor.vv v6, v24, v6
  1684: 5706862f                 vxor.vv v12, v24, v12
  1688: 5708882f                 vxor.vv v16, v24, v16
  1692: 8333050e                 ld t2, 0xe0(a0)
  1696: b3836300                 add t2, t2, t1
  1700: 078c0302                 vle8.v v24, (t2)
  1704: 5704842f                 vxor.vv v8, v24, v8
  1708: 5707872f                 vxor.vv v14, v24, v14
  1712: 5709892f                 vxor.vv v18, v24, v18
  1716: 8333850e                 ld t2, 0xe8(a0)
  1720: b3836300                 add t2, t2, t1
  1724: 078c0302                 vle8.v v24, (t2)
  1728: 5705852f                 vxor.vv v10, v24, v10
  1732: 5708882f                 vxor.vv v16, v24, v16
  1736: 570a8a2f                 vxor.vv v20, v24, v20
  1740: 8333050f                 ld t2, 0xf0(a0)
  1744: b3836300                 add t2, t2, t1
  1748: 078c0302                 vle8.v v24, (t2)
  1752: 5700802f                 vxor.vv v0, v24, v0
  1756: 570b8b2f                 vxor.vv v22, v24, v22
  1760: 8333850f                 ld t2, 0xf8(a0)
  1764: b3836300                 add t2, t2, t1
  1768: 078c0302                 vle8.v v24, (t2)
  1772: 5701812f                 vxor.vv v2, v24, v2
  1776: 5706862f                 vxor.vv v12, v24, v12
  1780: 83330510                 ld t2, 0x100(a0)
  1784: b3836300                 add t2, t2, t1
  1788: 078c0302                 vle8.v v24, (t2)
  1792: 5702822f                 vxor.vv v4, v24, v4
  1796: 5706862f                 vxor.vv v12, v24, v12
  1800: 5707872f                 vxor.vv v14, v24, v14
  1804: 83338510                 ld t2, 0x108(a0)
  1808: b3836300                 add t2, t2, t1
  1812: 078c0302                 vle8.v v24, (t2)
  1816: 5703832f                 vxor.vv v6, v24, v6
  1820: 5707872f                 vxor.vv v14, v24, v14
  1824: 5708882f                 vxor.vv v16, v24, v16
  1828: 83330511                 ld t2, 0x110(a0)
  1832: b3836300                 add t2, t2, t1
  1836: 078c0302                 vle8.v v24, (t2)
  1840: 5704842f                 vxor.vv v8, v24, v8
  1844: 5708882f                 vxor.vv v16, v24, v16
  1848: 5709892f                 vxor.vv v18, v24, v18
  1852: 83338511                 ld t2, 0x118(a0)
  1856: b3836300                 add t2, t2, t1
  1860: 078c0302                 vle8.v v24, (t2)
  1864: 5705852f                 vxor.vv v10, v24, v10
  1868: 5709892f                 vxor.vv v18, v24, v18
  1872: 570a8a2f                 vxor.vv v20, v24, v20
  1876: 83b30500                 ld t2, 0(a1)
  1880: b3836300                 add t2, t2, t1
  1884: 27800302                 vse8.v v0, (t2)
  1888: 83b38500                 ld t2, 8(a1)
  1892: b3836300                 add t2, t2, t1
  1896: 27810302                 vse8.v v2, (t2)
  1900: 83b30501                 ld t2, 0x10(a1)
  1904: b3836300                 add t2, t2, t1
  1908: 27820302                 vse8.v v4, (t2)
  1912: 83b38501                 ld t2, 0x18(a1)
  1916: b3836300                 add t2, t2, t1
  1920: 27830302                 vse8.v v6, (t2)
  1924: 83b30502                 ld t2, 0x20(a1)
  1928: b3836300                 add t2, t2, t1
  1932: 27840302                 vse8.v v8, (t2)
  1936: 83b38502                 ld t2, 0x28(a1)
  1940: b3836300                 add t2, t2, t1
  1944: 27850302                 vse8.v v10, (t2)
  1948: 83b30503                 ld t2, 0x30(a1)
  1952: b3836300                 add t2, t2, t1
  1956: 27860302                 vse8.v v12, (t2)
  1960: 83b38503                 ld t2, 0x38(a1)
  1964: b3836300                 add t2, t2, t1
  1968: 27870302                 vse8.v v14, (t2)
  1972: 83b30504                 ld t2, 0x40(a1)
  1976: b3836300                 add t2, t2, t1
  1980: 27880302                 vse8.v v16, (t2)
  1984: 83b38504                 ld t2, 0x48(a1)
  1988: b3836300                 add t2, t2, t1
  1992: 27890302                 vse8.v v18, (t2)
  1996: 83b30505                 ld t2, 0x50(a1)
  2000: b3836300                 add t2, t2, t1
  2004: 278a0302                 vse8.v v20, (t2)
  2008: 83b38505                 ld t2, 0x58(a1)
  2012: b3836300                 add t2, t2, t1
  2016: 278b0302                 vse8.v v22, (t2)
  2020: 33035300                 add t1, t1, t0
  2024: 33065640                 sub a2, a2, t0
  2028: 6ff09fc2                 j 0x414
  2032: 67800000                 ret 
  2036: 0000                     unimp 
  2038: 0000                     unimp 
  2040: 0000                     unimp 
  2042: 0000                     unimp 
  2044: 0000                     unimp 
  2046: 0000                     unimp 
