; config_t1_947 target=sve2 twin=pipe topo=rs(6,3)/enc op=encode U=1
; code_size=1216 twins=["pipe"] uops_per_byte=6.8125
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0250010                 adr x16, 0x4c0
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ed030caa                 mov x13, x12
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: fbc13825                 mov z27.b, #0xf
    36: 4e1100b4                 cbz x14, 0x24c
    40: 0a0140f9                 ldr x10, [x8]
    44: 4a010b8b                 add x10, x10, x11
    48: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 90327404                 mov z16.d, z20.d
    68: 0f020091                 add x15, x16, #0
    72: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
    76: 0f020191                 add x15, x16, #0x40
    80: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
    84: 95323b04                 and z21.d, z20.d, z27.d
    88: 96962c04                 lsr z22.b, z20.b, #4
    92: 37333505                 tbl z23.b, { z25.b }, z21.b
    96: 58333605                 tbl z24.b, { z26.b }, z22.b
   100: f132b804                 eor z17.d, z23.d, z24.d
   104: 0f020291                 add x15, x16, #0x80
   108: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   112: 0f020391                 add x15, x16, #0xc0
   116: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   120: 95323b04                 and z21.d, z20.d, z27.d
   124: 96962c04                 lsr z22.b, z20.b, #4
   128: 37333505                 tbl z23.b, { z25.b }, z21.b
   132: 58333605                 tbl z24.b, { z26.b }, z22.b
   136: f232b804                 eor z18.d, z23.d, z24.d
   140: 0a0540f9                 ldr x10, [x8, #8]
   144: 4a010b8b                 add x10, x10, x11
   148: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   152: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   156: 0f020291                 add x15, x16, #0x80
   160: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   164: 0f020391                 add x15, x16, #0xc0
   168: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   172: 75323b04                 and z21.d, z19.d, z27.d
   176: 76962c04                 lsr z22.b, z19.b, #4
   180: 37333505                 tbl z23.b, { z25.b }, z21.b
   184: 58333605                 tbl z24.b, { z26.b }, z22.b
   188: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   192: 0f020491                 add x15, x16, #0x100
   196: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   200: 0f020591                 add x15, x16, #0x140
   204: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   208: 75323b04                 and z21.d, z19.d, z27.d
   212: 76962c04                 lsr z22.b, z19.b, #4
   216: 37333505                 tbl z23.b, { z25.b }, z21.b
   220: 58333605                 tbl z24.b, { z26.b }, z22.b
   224: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
   228: 0a0940f9                 ldr x10, [x8, #0x10]
   232: 4a010b8b                 add x10, x10, x11
   236: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   240: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   244: 0f020491                 add x15, x16, #0x100
   248: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   252: 0f020591                 add x15, x16, #0x140
   256: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   260: 75323b04                 and z21.d, z19.d, z27.d
   264: 76962c04                 lsr z22.b, z19.b, #4
   268: 37333505                 tbl z23.b, { z25.b }, z21.b
   272: 58333605                 tbl z24.b, { z26.b }, z22.b
   276: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   280: 0f020691                 add x15, x16, #0x180
   284: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   288: 0f020791                 add x15, x16, #0x1c0
   292: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   296: 75323b04                 and z21.d, z19.d, z27.d
   300: 76962c04                 lsr z22.b, z19.b, #4
   304: 37333505                 tbl z23.b, { z25.b }, z21.b
   308: 58333605                 tbl z24.b, { z26.b }, z22.b
   312: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
   316: 0a0d40f9                 ldr x10, [x8, #0x18]
   320: 4a010b8b                 add x10, x10, x11
   324: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   328: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   332: 0f020891                 add x15, x16, #0x200
   336: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   340: 0f020991                 add x15, x16, #0x240
   344: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   348: 75323b04                 and z21.d, z19.d, z27.d
   352: 76962c04                 lsr z22.b, z19.b, #4
   356: 37333505                 tbl z23.b, { z25.b }, z21.b
   360: 58333605                 tbl z24.b, { z26.b }, z22.b
   364: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   368: 0f020091                 add x15, x16, #0
   372: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   376: 0f020191                 add x15, x16, #0x40
   380: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   384: 75323b04                 and z21.d, z19.d, z27.d
   388: 76962c04                 lsr z22.b, z19.b, #4
   392: 37333505                 tbl z23.b, { z25.b }, z21.b
   396: 58333605                 tbl z24.b, { z26.b }, z22.b
   400: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
   404: 0a1140f9                 ldr x10, [x8, #0x20]
   408: 4a010b8b                 add x10, x10, x11
   412: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   416: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   420: 0f020691                 add x15, x16, #0x180
   424: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   428: 0f020791                 add x15, x16, #0x1c0
   432: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   436: 75323b04                 and z21.d, z19.d, z27.d
   440: 76962c04                 lsr z22.b, z19.b, #4
   444: 37333505                 tbl z23.b, { z25.b }, z21.b
   448: 58333605                 tbl z24.b, { z26.b }, z22.b
   452: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   456: 0f020891                 add x15, x16, #0x200
   460: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   464: 0f020991                 add x15, x16, #0x240
   468: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   472: 75323b04                 and z21.d, z19.d, z27.d
   476: 76962c04                 lsr z22.b, z19.b, #4
   480: 37333505                 tbl z23.b, { z25.b }, z21.b
   484: 58333605                 tbl z24.b, { z26.b }, z22.b
   488: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
   492: 0a1540f9                 ldr x10, [x8, #0x28]
   496: 4a010b8b                 add x10, x10, x11
   500: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   504: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   508: 71021904                 eor z17.b, p0/m, z17.b, z19.b
   512: 72021904                 eor z18.b, p0/m, z18.b, z19.b
   516: df0500f1                 cmp x14, #1
   520: ad000054                 b.le 0x21c
   524: 0a0140f9                 ldr x10, [x8]
   528: 4a010b8b                 add x10, x10, x11
   532: 4a410091                 add x10, x10, #0x10
   536: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
   540: 2a0140f9                 ldr x10, [x9]
   544: 4a010b8b                 add x10, x10, x11
   548: 50e100e4                 st1b { z16.b }, p0, [x10]
   552: 2a0540f9                 ldr x10, [x9, #8]
   556: 4a010b8b                 add x10, x10, x11
   560: 51e100e4                 st1b { z17.b }, p0, [x10]
   564: 2a0940f9                 ldr x10, [x9, #0x10]
   568: 4a010b8b                 add x10, x10, x11
   572: 52e100e4                 st1b { z18.b }, p0, [x10]
   576: 6b010d8b                 add x11, x11, x13
   580: ce0500f1                 subs x14, x14, #1
   584: 61efff54                 b.ne 0x34
   588: 7f0102eb                 cmp x11, x2
   592: 82120054                 b.hs 0x4a0
   596: 0f020a91                 add x15, x16, #0x280
   600: 11020e91                 add x17, x16, #0x380
   604: 2a0140f9                 ldr x10, [x9]
   608: ed030baa                 mov x13, x11
   612: 0c008052                 mov w12, #0
   616: 0e0140f9                 ldr x14, [x8]
   620: ce696d38                 ldrb w14, [x14, x13]
   624: 8c010e4a                 eor w12, w12, w14
   628: 0e0540f9                 ldr x14, [x8, #8]
   632: ce696d38                 ldrb w14, [x14, x13]
   636: 8c010e4a                 eor w12, w12, w14
   640: 0e0940f9                 ldr x14, [x8, #0x10]
   644: ce696d38                 ldrb w14, [x14, x13]
   648: 8c010e4a                 eor w12, w12, w14
   652: 0e0d40f9                 ldr x14, [x8, #0x18]
   656: ce696d38                 ldrb w14, [x14, x13]
   660: 8c010e4a                 eor w12, w12, w14
   664: 0e1140f9                 ldr x14, [x8, #0x20]
   668: ce696d38                 ldrb w14, [x14, x13]
   672: 8c010e4a                 eor w12, w12, w14
   676: 0e1540f9                 ldr x14, [x8, #0x28]
   680: ce696d38                 ldrb w14, [x14, x13]
   684: 8c010e4a                 eor w12, w12, w14
   688: 4c692d38                 strb w12, [x10, x13]
   692: ad050091                 add x13, x13, #1
   696: bf0102eb                 cmp x13, x2
   700: 43fdff54                 b.lo 0x264
   704: 2a0540f9                 ldr x10, [x9, #8]
   708: ed030baa                 mov x13, x11
   712: 0c008052                 mov w12, #0
   716: 0e0140f9                 ldr x14, [x8]
   720: ce696d38                 ldrb w14, [x14, x13]
   724: 0e010034                 cbz w14, 0x2f4
   728: ee696e38                 ldrb w14, [x15, x14]
   732: ce990311                 add w14, w14, #0xe6
   736: dffd0371                 cmp w14, #0xff
   740: 43000054                 b.lo 0x2ec
   744: cefd0351                 sub w14, w14, #0xff
   748: 2e6a6e38                 ldrb w14, [x17, x14]
   752: 8c010e4a                 eor w12, w12, w14
   756: 0e0540f9                 ldr x14, [x8, #8]
   760: ce696d38                 ldrb w14, [x14, x13]
   764: 0e010034                 cbz w14, 0x31c
   768: ee696e38                 ldrb w14, [x15, x14]
   772: ce350311                 add w14, w14, #0xcd
   776: dffd0371                 cmp w14, #0xff
   780: 43000054                 b.lo 0x314
   784: cefd0351                 sub w14, w14, #0xff
   788: 2e6a6e38                 ldrb w14, [x17, x14]
   792: 8c010e4a                 eor w12, w12, w14
   796: 0e0940f9                 ldr x14, [x8, #0x10]
   800: ce696d38                 ldrb w14, [x14, x13]
   804: 0e010034                 cbz w14, 0x344
   808: ee696e38                 ldrb w14, [x15, x14]
   812: ce2d0111                 add w14, w14, #0x4b
   816: dffd0371                 cmp w14, #0xff
   820: 43000054                 b.lo 0x33c
   824: cefd0351                 sub w14, w14, #0xff
   828: 2e6a6e38                 ldrb w14, [x17, x14]
   832: 8c010e4a                 eor w12, w12, w14
   836: 0e0d40f9                 ldr x14, [x8, #0x18]
   840: ce696d38                 ldrb w14, [x14, x13]
   844: 0e010034                 cbz w14, 0x36c
   848: ee696e38                 ldrb w14, [x15, x14]
   852: cec90011                 add w14, w14, #0x32
   856: dffd0371                 cmp w14, #0xff
   860: 43000054                 b.lo 0x364
   864: cefd0351                 sub w14, w14, #0xff
   868: 2e6a6e38                 ldrb w14, [x17, x14]
   872: 8c010e4a                 eor w12, w12, w14
   876: 0e1140f9                 ldr x14, [x8, #0x20]
   880: ce696d38                 ldrb w14, [x14, x13]
   884: 0e010034                 cbz w14, 0x394
   888: ee696e38                 ldrb w14, [x15, x14]
   892: ce650011                 add w14, w14, #0x19
   896: dffd0371                 cmp w14, #0xff
   900: 43000054                 b.lo 0x38c
   904: cefd0351                 sub w14, w14, #0xff
   908: 2e6a6e38                 ldrb w14, [x17, x14]
   912: 8c010e4a                 eor w12, w12, w14
   916: 0e1540f9                 ldr x14, [x8, #0x28]
   920: ce696d38                 ldrb w14, [x14, x13]
   924: 8c010e4a                 eor w12, w12, w14
   928: 4c692d38                 strb w12, [x10, x13]
   932: ad050091                 add x13, x13, #1
   936: bf0102eb                 cmp x13, x2
   940: e3f8ff54                 b.lo 0x2c8
   944: 2a0940f9                 ldr x10, [x9, #0x10]
   948: ed030baa                 mov x13, x11
   952: 0c008052                 mov w12, #0
   956: 0e0140f9                 ldr x14, [x8]
   960: ce696d38                 ldrb w14, [x14, x13]
   964: 0e010034                 cbz w14, 0x3e4
   968: ee696e38                 ldrb w14, [x15, x14]
   972: ce350311                 add w14, w14, #0xcd
   976: dffd0371                 cmp w14, #0xff
   980: 43000054                 b.lo 0x3dc
   984: cefd0351                 sub w14, w14, #0xff
   988: 2e6a6e38                 ldrb w14, [x17, x14]
   992: 8c010e4a                 eor w12, w12, w14
   996: 0e0540f9                 ldr x14, [x8, #8]
  1000: ce696d38                 ldrb w14, [x14, x13]
  1004: 0e010034                 cbz w14, 0x40c
  1008: ee696e38                 ldrb w14, [x15, x14]
  1012: ce2d0111                 add w14, w14, #0x4b
  1016: dffd0371                 cmp w14, #0xff
  1020: 43000054                 b.lo 0x404
  1024: cefd0351                 sub w14, w14, #0xff
  1028: 2e6a6e38                 ldrb w14, [x17, x14]
  1032: 8c010e4a                 eor w12, w12, w14
  1036: 0e0940f9                 ldr x14, [x8, #0x10]
  1040: ce696d38                 ldrb w14, [x14, x13]
  1044: 0e010034                 cbz w14, 0x434
  1048: ee696e38                 ldrb w14, [x15, x14]
  1052: ce650011                 add w14, w14, #0x19
  1056: dffd0371                 cmp w14, #0xff
  1060: 43000054                 b.lo 0x42c
  1064: cefd0351                 sub w14, w14, #0xff
  1068: 2e6a6e38                 ldrb w14, [x17, x14]
  1072: 8c010e4a                 eor w12, w12, w14
  1076: 0e0d40f9                 ldr x14, [x8, #0x18]
  1080: ce696d38                 ldrb w14, [x14, x13]
  1084: 0e010034                 cbz w14, 0x45c
  1088: ee696e38                 ldrb w14, [x15, x14]
  1092: ce990311                 add w14, w14, #0xe6
  1096: dffd0371                 cmp w14, #0xff
  1100: 43000054                 b.lo 0x454
  1104: cefd0351                 sub w14, w14, #0xff
  1108: 2e6a6e38                 ldrb w14, [x17, x14]
  1112: 8c010e4a                 eor w12, w12, w14
  1116: 0e1140f9                 ldr x14, [x8, #0x20]
  1120: ce696d38                 ldrb w14, [x14, x13]
  1124: 0e010034                 cbz w14, 0x484
  1128: ee696e38                 ldrb w14, [x15, x14]
  1132: cec90011                 add w14, w14, #0x32
  1136: dffd0371                 cmp w14, #0xff
  1140: 43000054                 b.lo 0x47c
  1144: cefd0351                 sub w14, w14, #0xff
  1148: 2e6a6e38                 ldrb w14, [x17, x14]
  1152: 8c010e4a                 eor w12, w12, w14
  1156: 0e1540f9                 ldr x14, [x8, #0x28]
  1160: ce696d38                 ldrb w14, [x14, x13]
  1164: 8c010e4a                 eor w12, w12, w14
  1168: 4c692d38                 strb w12, [x10, x13]
  1172: ad050091                 add x13, x13, #1
  1176: bf0102eb                 cmp x13, x2
  1180: e3f8ff54                 b.lo 0x3b8
  1184: c0035fd6                 ret 
  1188: 00000000                 udf #0
  1192: 00000000                 udf #0
  1196: 00000000                 udf #0
  1200: 00000000                 udf #0
  1204: 00000000                 udf #0
  1208: 00000000                 udf #0
  1212: 00000000                 udf #0
