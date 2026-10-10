; config_t1_937 target=sve2 twin=nt topo=rs(6,3)/enc op=encode U=1
; code_size=1280 twins=["nt"] uops_per_byte=7.5
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0270010                 adr x16, 0x500
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ed030caa                 mov x13, x12
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: fac13825                 mov z26.b, #0xf
    36: 6e1200b4                 cbz x14, 0x270
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
    76: 70327304                 mov z16.d, z19.d
    80: 0f020091                 add x15, x16, #0
    84: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
    88: 0f020191                 add x15, x16, #0x40
    92: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
    96: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   100: 74323a04                 and z20.d, z19.d, z26.d
   104: 75962c04                 lsr z21.b, z19.b, #4
   108: 16333405                 tbl z22.b, { z24.b }, z20.b
   112: 37333505                 tbl z23.b, { z25.b }, z21.b
   116: d132b704                 eor z17.d, z22.d, z23.d
   120: 0f020291                 add x15, x16, #0x80
   124: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   128: 0f020391                 add x15, x16, #0xc0
   132: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   136: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   140: 74323a04                 and z20.d, z19.d, z26.d
   144: 75962c04                 lsr z21.b, z19.b, #4
   148: 16333405                 tbl z22.b, { z24.b }, z20.b
   152: 37333505                 tbl z23.b, { z25.b }, z21.b
   156: d232b704                 eor z18.d, z22.d, z23.d
   160: 0a0540f9                 ldr x10, [x8, #8]
   164: 4a010b8b                 add x10, x10, x11
   168: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   172: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   176: 0f020291                 add x15, x16, #0x80
   180: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   184: 0f020391                 add x15, x16, #0xc0
   188: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   192: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   196: 74323a04                 and z20.d, z19.d, z26.d
   200: 75962c04                 lsr z21.b, z19.b, #4
   204: 16333405                 tbl z22.b, { z24.b }, z20.b
   208: 37333505                 tbl z23.b, { z25.b }, z21.b
   212: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   216: 0f020491                 add x15, x16, #0x100
   220: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   224: 0f020591                 add x15, x16, #0x140
   228: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   232: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   236: 74323a04                 and z20.d, z19.d, z26.d
   240: 75962c04                 lsr z21.b, z19.b, #4
   244: 16333405                 tbl z22.b, { z24.b }, z20.b
   248: 37333505                 tbl z23.b, { z25.b }, z21.b
   252: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   256: 0a0940f9                 ldr x10, [x8, #0x10]
   260: 4a010b8b                 add x10, x10, x11
   264: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   268: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   272: 0f020491                 add x15, x16, #0x100
   276: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   280: 0f020591                 add x15, x16, #0x140
   284: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   288: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   292: 74323a04                 and z20.d, z19.d, z26.d
   296: 75962c04                 lsr z21.b, z19.b, #4
   300: 16333405                 tbl z22.b, { z24.b }, z20.b
   304: 37333505                 tbl z23.b, { z25.b }, z21.b
   308: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   312: 0f020691                 add x15, x16, #0x180
   316: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   320: 0f020791                 add x15, x16, #0x1c0
   324: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   328: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   332: 74323a04                 and z20.d, z19.d, z26.d
   336: 75962c04                 lsr z21.b, z19.b, #4
   340: 16333405                 tbl z22.b, { z24.b }, z20.b
   344: 37333505                 tbl z23.b, { z25.b }, z21.b
   348: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   352: 0a0d40f9                 ldr x10, [x8, #0x18]
   356: 4a010b8b                 add x10, x10, x11
   360: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   364: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   368: 0f020891                 add x15, x16, #0x200
   372: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   376: 0f020991                 add x15, x16, #0x240
   380: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   384: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   388: 74323a04                 and z20.d, z19.d, z26.d
   392: 75962c04                 lsr z21.b, z19.b, #4
   396: 16333405                 tbl z22.b, { z24.b }, z20.b
   400: 37333505                 tbl z23.b, { z25.b }, z21.b
   404: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   408: 0f020091                 add x15, x16, #0
   412: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   416: 0f020191                 add x15, x16, #0x40
   420: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   424: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   428: 74323a04                 and z20.d, z19.d, z26.d
   432: 75962c04                 lsr z21.b, z19.b, #4
   436: 16333405                 tbl z22.b, { z24.b }, z20.b
   440: 37333505                 tbl z23.b, { z25.b }, z21.b
   444: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   448: 0a1140f9                 ldr x10, [x8, #0x20]
   452: 4a010b8b                 add x10, x10, x11
   456: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   460: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   464: 0f020691                 add x15, x16, #0x180
   468: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   472: 0f020791                 add x15, x16, #0x1c0
   476: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   480: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   484: 74323a04                 and z20.d, z19.d, z26.d
   488: 75962c04                 lsr z21.b, z19.b, #4
   492: 16333405                 tbl z22.b, { z24.b }, z20.b
   496: 37333505                 tbl z23.b, { z25.b }, z21.b
   500: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   504: 0f020891                 add x15, x16, #0x200
   508: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   512: 0f020991                 add x15, x16, #0x240
   516: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   520: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   524: 74323a04                 and z20.d, z19.d, z26.d
   528: 75962c04                 lsr z21.b, z19.b, #4
   532: 16333405                 tbl z22.b, { z24.b }, z20.b
   536: 37333505                 tbl z23.b, { z25.b }, z21.b
   540: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   544: 0a1540f9                 ldr x10, [x8, #0x28]
   548: 4a010b8b                 add x10, x10, x11
   552: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   556: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   560: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   564: 71021904                 eor z17.b, p0/m, z17.b, z19.b
   568: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   572: 72021904                 eor z18.b, p0/m, z18.b, z19.b
   576: 2a0140f9                 ldr x10, [x9]
   580: 4a010b8b                 add x10, x10, x11
   584: 50e110e4                 stnt1b { z16.b }, p0, [x10]
   588: 2a0540f9                 ldr x10, [x9, #8]
   592: 4a010b8b                 add x10, x10, x11
   596: 51e110e4                 stnt1b { z17.b }, p0, [x10]
   600: 2a0940f9                 ldr x10, [x9, #0x10]
   604: 4a010b8b                 add x10, x10, x11
   608: 52e110e4                 stnt1b { z18.b }, p0, [x10]
   612: 6b010d8b                 add x11, x11, x13
   616: ce0500f1                 subs x14, x14, #1
   620: e1edff54                 b.ne 0x28
   624: 7f0102eb                 cmp x11, x2
   628: 82120054                 b.hs 0x4c4
   632: 0f020a91                 add x15, x16, #0x280
   636: 11020e91                 add x17, x16, #0x380
   640: 2a0140f9                 ldr x10, [x9]
   644: ed030baa                 mov x13, x11
   648: 0c008052                 mov w12, #0
   652: 0e0140f9                 ldr x14, [x8]
   656: ce696d38                 ldrb w14, [x14, x13]
   660: 8c010e4a                 eor w12, w12, w14
   664: 0e0540f9                 ldr x14, [x8, #8]
   668: ce696d38                 ldrb w14, [x14, x13]
   672: 8c010e4a                 eor w12, w12, w14
   676: 0e0940f9                 ldr x14, [x8, #0x10]
   680: ce696d38                 ldrb w14, [x14, x13]
   684: 8c010e4a                 eor w12, w12, w14
   688: 0e0d40f9                 ldr x14, [x8, #0x18]
   692: ce696d38                 ldrb w14, [x14, x13]
   696: 8c010e4a                 eor w12, w12, w14
   700: 0e1140f9                 ldr x14, [x8, #0x20]
   704: ce696d38                 ldrb w14, [x14, x13]
   708: 8c010e4a                 eor w12, w12, w14
   712: 0e1540f9                 ldr x14, [x8, #0x28]
   716: ce696d38                 ldrb w14, [x14, x13]
   720: 8c010e4a                 eor w12, w12, w14
   724: 4c692d38                 strb w12, [x10, x13]
   728: ad050091                 add x13, x13, #1
   732: bf0102eb                 cmp x13, x2
   736: 43fdff54                 b.lo 0x288
   740: 2a0540f9                 ldr x10, [x9, #8]
   744: ed030baa                 mov x13, x11
   748: 0c008052                 mov w12, #0
   752: 0e0140f9                 ldr x14, [x8]
   756: ce696d38                 ldrb w14, [x14, x13]
   760: 0e010034                 cbz w14, 0x318
   764: ee696e38                 ldrb w14, [x15, x14]
   768: ce990311                 add w14, w14, #0xe6
   772: dffd0371                 cmp w14, #0xff
   776: 43000054                 b.lo 0x310
   780: cefd0351                 sub w14, w14, #0xff
   784: 2e6a6e38                 ldrb w14, [x17, x14]
   788: 8c010e4a                 eor w12, w12, w14
   792: 0e0540f9                 ldr x14, [x8, #8]
   796: ce696d38                 ldrb w14, [x14, x13]
   800: 0e010034                 cbz w14, 0x340
   804: ee696e38                 ldrb w14, [x15, x14]
   808: ce350311                 add w14, w14, #0xcd
   812: dffd0371                 cmp w14, #0xff
   816: 43000054                 b.lo 0x338
   820: cefd0351                 sub w14, w14, #0xff
   824: 2e6a6e38                 ldrb w14, [x17, x14]
   828: 8c010e4a                 eor w12, w12, w14
   832: 0e0940f9                 ldr x14, [x8, #0x10]
   836: ce696d38                 ldrb w14, [x14, x13]
   840: 0e010034                 cbz w14, 0x368
   844: ee696e38                 ldrb w14, [x15, x14]
   848: ce2d0111                 add w14, w14, #0x4b
   852: dffd0371                 cmp w14, #0xff
   856: 43000054                 b.lo 0x360
   860: cefd0351                 sub w14, w14, #0xff
   864: 2e6a6e38                 ldrb w14, [x17, x14]
   868: 8c010e4a                 eor w12, w12, w14
   872: 0e0d40f9                 ldr x14, [x8, #0x18]
   876: ce696d38                 ldrb w14, [x14, x13]
   880: 0e010034                 cbz w14, 0x390
   884: ee696e38                 ldrb w14, [x15, x14]
   888: cec90011                 add w14, w14, #0x32
   892: dffd0371                 cmp w14, #0xff
   896: 43000054                 b.lo 0x388
   900: cefd0351                 sub w14, w14, #0xff
   904: 2e6a6e38                 ldrb w14, [x17, x14]
   908: 8c010e4a                 eor w12, w12, w14
   912: 0e1140f9                 ldr x14, [x8, #0x20]
   916: ce696d38                 ldrb w14, [x14, x13]
   920: 0e010034                 cbz w14, 0x3b8
   924: ee696e38                 ldrb w14, [x15, x14]
   928: ce650011                 add w14, w14, #0x19
   932: dffd0371                 cmp w14, #0xff
   936: 43000054                 b.lo 0x3b0
   940: cefd0351                 sub w14, w14, #0xff
   944: 2e6a6e38                 ldrb w14, [x17, x14]
   948: 8c010e4a                 eor w12, w12, w14
   952: 0e1540f9                 ldr x14, [x8, #0x28]
   956: ce696d38                 ldrb w14, [x14, x13]
   960: 8c010e4a                 eor w12, w12, w14
   964: 4c692d38                 strb w12, [x10, x13]
   968: ad050091                 add x13, x13, #1
   972: bf0102eb                 cmp x13, x2
   976: e3f8ff54                 b.lo 0x2ec
   980: 2a0940f9                 ldr x10, [x9, #0x10]
   984: ed030baa                 mov x13, x11
   988: 0c008052                 mov w12, #0
   992: 0e0140f9                 ldr x14, [x8]
   996: ce696d38                 ldrb w14, [x14, x13]
  1000: 0e010034                 cbz w14, 0x408
  1004: ee696e38                 ldrb w14, [x15, x14]
  1008: ce350311                 add w14, w14, #0xcd
  1012: dffd0371                 cmp w14, #0xff
  1016: 43000054                 b.lo 0x400
  1020: cefd0351                 sub w14, w14, #0xff
  1024: 2e6a6e38                 ldrb w14, [x17, x14]
  1028: 8c010e4a                 eor w12, w12, w14
  1032: 0e0540f9                 ldr x14, [x8, #8]
  1036: ce696d38                 ldrb w14, [x14, x13]
  1040: 0e010034                 cbz w14, 0x430
  1044: ee696e38                 ldrb w14, [x15, x14]
  1048: ce2d0111                 add w14, w14, #0x4b
  1052: dffd0371                 cmp w14, #0xff
  1056: 43000054                 b.lo 0x428
  1060: cefd0351                 sub w14, w14, #0xff
  1064: 2e6a6e38                 ldrb w14, [x17, x14]
  1068: 8c010e4a                 eor w12, w12, w14
  1072: 0e0940f9                 ldr x14, [x8, #0x10]
  1076: ce696d38                 ldrb w14, [x14, x13]
  1080: 0e010034                 cbz w14, 0x458
  1084: ee696e38                 ldrb w14, [x15, x14]
  1088: ce650011                 add w14, w14, #0x19
  1092: dffd0371                 cmp w14, #0xff
  1096: 43000054                 b.lo 0x450
  1100: cefd0351                 sub w14, w14, #0xff
  1104: 2e6a6e38                 ldrb w14, [x17, x14]
  1108: 8c010e4a                 eor w12, w12, w14
  1112: 0e0d40f9                 ldr x14, [x8, #0x18]
  1116: ce696d38                 ldrb w14, [x14, x13]
  1120: 0e010034                 cbz w14, 0x480
  1124: ee696e38                 ldrb w14, [x15, x14]
  1128: ce990311                 add w14, w14, #0xe6
  1132: dffd0371                 cmp w14, #0xff
  1136: 43000054                 b.lo 0x478
  1140: cefd0351                 sub w14, w14, #0xff
  1144: 2e6a6e38                 ldrb w14, [x17, x14]
  1148: 8c010e4a                 eor w12, w12, w14
  1152: 0e1140f9                 ldr x14, [x8, #0x20]
  1156: ce696d38                 ldrb w14, [x14, x13]
  1160: 0e010034                 cbz w14, 0x4a8
  1164: ee696e38                 ldrb w14, [x15, x14]
  1168: cec90011                 add w14, w14, #0x32
  1172: dffd0371                 cmp w14, #0xff
  1176: 43000054                 b.lo 0x4a0
  1180: cefd0351                 sub w14, w14, #0xff
  1184: 2e6a6e38                 ldrb w14, [x17, x14]
  1188: 8c010e4a                 eor w12, w12, w14
  1192: 0e1540f9                 ldr x14, [x8, #0x28]
  1196: ce696d38                 ldrb w14, [x14, x13]
  1200: 8c010e4a                 eor w12, w12, w14
  1204: 4c692d38                 strb w12, [x10, x13]
  1208: ad050091                 add x13, x13, #1
  1212: bf0102eb                 cmp x13, x2
  1216: e3f8ff54                 b.lo 0x3dc
  1220: c0035fd6                 ret 
  1224: 00000000                 udf #0
  1228: 00000000                 udf #0
  1232: 00000000                 udf #0
  1236: 00000000                 udf #0
  1240: 00000000                 udf #0
  1244: 00000000                 udf #0
  1248: 00000000                 udf #0
  1252: 00000000                 udf #0
  1256: 00000000                 udf #0
  1260: 00000000                 udf #0
  1264: 00000000                 udf #0
  1268: 00000000                 udf #0
  1272: 00000000                 udf #0
  1276: 00000000                 udf #0
