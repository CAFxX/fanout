; config_t1_1051 target=sve2 twin=pipe topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1216 twins=["pipe"] uops_per_byte=7.375
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0250010                 adr x16, 0x4c0
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ed030caa                 mov x13, x12
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: fac13825                 mov z26.b, #0xf
    36: 2e1200b4                 cbz x14, 0x268
    40: 0a0140f9                 ldr x10, [x8]
    44: 4a010b8b                 add x10, x10, x11
    48: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0f020091                 add x15, x16, #0
    68: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
    72: 0f020191                 add x15, x16, #0x40
    76: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
    80: 74323a04                 and z20.d, z19.d, z26.d
    84: 75962c04                 lsr z21.b, z19.b, #4
    88: 16333405                 tbl z22.b, { z24.b }, z20.b
    92: 37333505                 tbl z23.b, { z25.b }, z21.b
    96: d032b704                 eor z16.d, z22.d, z23.d
   100: 0f020291                 add x15, x16, #0x80
   104: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   108: 0f020391                 add x15, x16, #0xc0
   112: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   116: 74323a04                 and z20.d, z19.d, z26.d
   120: 75962c04                 lsr z21.b, z19.b, #4
   124: 16333405                 tbl z22.b, { z24.b }, z20.b
   128: 37333505                 tbl z23.b, { z25.b }, z21.b
   132: d132b704                 eor z17.d, z22.d, z23.d
   136: 0a0540f9                 ldr x10, [x8, #8]
   140: 4a010b8b                 add x10, x10, x11
   144: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   148: 0f020491                 add x15, x16, #0x100
   152: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   156: 0f020591                 add x15, x16, #0x140
   160: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   164: 54323a04                 and z20.d, z18.d, z26.d
   168: 55962c04                 lsr z21.b, z18.b, #4
   172: 16333405                 tbl z22.b, { z24.b }, z20.b
   176: 37333505                 tbl z23.b, { z25.b }, z21.b
   180: f03a3604                 eor3 z16.d, z16.d, z22.d, z23.d
   184: 0f020691                 add x15, x16, #0x180
   188: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   192: 0f020791                 add x15, x16, #0x1c0
   196: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   200: 54323a04                 and z20.d, z18.d, z26.d
   204: 55962c04                 lsr z21.b, z18.b, #4
   208: 16333405                 tbl z22.b, { z24.b }, z20.b
   212: 37333505                 tbl z23.b, { z25.b }, z21.b
   216: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   220: 0a0940f9                 ldr x10, [x8, #0x10]
   224: 4a010b8b                 add x10, x10, x11
   228: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   232: 0f020891                 add x15, x16, #0x200
   236: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   240: 0f020991                 add x15, x16, #0x240
   244: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   248: 54323a04                 and z20.d, z18.d, z26.d
   252: 55962c04                 lsr z21.b, z18.b, #4
   256: 16333405                 tbl z22.b, { z24.b }, z20.b
   260: 37333505                 tbl z23.b, { z25.b }, z21.b
   264: f03a3604                 eor3 z16.d, z16.d, z22.d, z23.d
   268: 0f020a91                 add x15, x16, #0x280
   272: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   276: 0f020b91                 add x15, x16, #0x2c0
   280: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   284: 54323a04                 and z20.d, z18.d, z26.d
   288: 55962c04                 lsr z21.b, z18.b, #4
   292: 16333405                 tbl z22.b, { z24.b }, z20.b
   296: 37333505                 tbl z23.b, { z25.b }, z21.b
   300: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   304: 0a0d40f9                 ldr x10, [x8, #0x18]
   308: 4a010b8b                 add x10, x10, x11
   312: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   316: 0f020c91                 add x15, x16, #0x300
   320: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   324: 0f020d91                 add x15, x16, #0x340
   328: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   332: 54323a04                 and z20.d, z18.d, z26.d
   336: 55962c04                 lsr z21.b, z18.b, #4
   340: 16333405                 tbl z22.b, { z24.b }, z20.b
   344: 37333505                 tbl z23.b, { z25.b }, z21.b
   348: f03a3604                 eor3 z16.d, z16.d, z22.d, z23.d
   352: 0f020e91                 add x15, x16, #0x380
   356: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   360: 0f020f91                 add x15, x16, #0x3c0
   364: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   368: 54323a04                 and z20.d, z18.d, z26.d
   372: 55962c04                 lsr z21.b, z18.b, #4
   376: 16333405                 tbl z22.b, { z24.b }, z20.b
   380: 37333505                 tbl z23.b, { z25.b }, z21.b
   384: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   388: 0a1140f9                 ldr x10, [x8, #0x20]
   392: 4a010b8b                 add x10, x10, x11
   396: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   400: 0f021091                 add x15, x16, #0x400
   404: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   408: 0f021191                 add x15, x16, #0x440
   412: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   416: 54323a04                 and z20.d, z18.d, z26.d
   420: 55962c04                 lsr z21.b, z18.b, #4
   424: 16333405                 tbl z22.b, { z24.b }, z20.b
   428: 37333505                 tbl z23.b, { z25.b }, z21.b
   432: f03a3604                 eor3 z16.d, z16.d, z22.d, z23.d
   436: 0f021291                 add x15, x16, #0x480
   440: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   444: 0f021391                 add x15, x16, #0x4c0
   448: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   452: 54323a04                 and z20.d, z18.d, z26.d
   456: 55962c04                 lsr z21.b, z18.b, #4
   460: 16333405                 tbl z22.b, { z24.b }, z20.b
   464: 37333505                 tbl z23.b, { z25.b }, z21.b
   468: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   472: 0a1540f9                 ldr x10, [x8, #0x28]
   476: 4a010b8b                 add x10, x10, x11
   480: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   484: 0f021491                 add x15, x16, #0x500
   488: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   492: 0f021591                 add x15, x16, #0x540
   496: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   500: 54323a04                 and z20.d, z18.d, z26.d
   504: 55962c04                 lsr z21.b, z18.b, #4
   508: 16333405                 tbl z22.b, { z24.b }, z20.b
   512: 37333505                 tbl z23.b, { z25.b }, z21.b
   516: f03a3604                 eor3 z16.d, z16.d, z22.d, z23.d
   520: 0f021491                 add x15, x16, #0x500
   524: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   528: 0f021591                 add x15, x16, #0x540
   532: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   536: 54323a04                 and z20.d, z18.d, z26.d
   540: 55962c04                 lsr z21.b, z18.b, #4
   544: 16333405                 tbl z22.b, { z24.b }, z20.b
   548: 37333505                 tbl z23.b, { z25.b }, z21.b
   552: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   556: df0500f1                 cmp x14, #1
   560: ad000054                 b.le 0x244
   564: 0a0140f9                 ldr x10, [x8]
   568: 4a010b8b                 add x10, x10, x11
   572: 4a410091                 add x10, x10, #0x10
   576: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   580: 2a0140f9                 ldr x10, [x9]
   584: 4a010b8b                 add x10, x10, x11
   588: 50e100e4                 st1b { z16.b }, p0, [x10]
   592: 2a0540f9                 ldr x10, [x9, #8]
   596: 4a010b8b                 add x10, x10, x11
   600: 51e100e4                 st1b { z17.b }, p0, [x10]
   604: 6b010d8b                 add x11, x11, x13
   608: ce0500f1                 subs x14, x14, #1
   612: 81eeff54                 b.ne 0x34
   616: 7f0102eb                 cmp x11, x2
   620: 22110054                 b.hs 0x490
   624: 0f021691                 add x15, x16, #0x580
   628: 11021a91                 add x17, x16, #0x680
   632: 2a0140f9                 ldr x10, [x9]
   636: ed030baa                 mov x13, x11
   640: 0c008052                 mov w12, #0
   644: 0e0140f9                 ldr x14, [x8]
   648: ce696d38                 ldrb w14, [x14, x13]
   652: 0e010034                 cbz w14, 0x2ac
   656: ee696e38                 ldrb w14, [x15, x14]
   660: ce050311                 add w14, w14, #0xc1
   664: dffd0371                 cmp w14, #0xff
   668: 43000054                 b.lo 0x2a4
   672: cefd0351                 sub w14, w14, #0xff
   676: 2e6a6e38                 ldrb w14, [x17, x14]
   680: 8c010e4a                 eor w12, w12, w14
   684: 0e0540f9                 ldr x14, [x8, #8]
   688: ce696d38                 ldrb w14, [x14, x13]
   692: 0e010034                 cbz w14, 0x2d4
   696: ee696e38                 ldrb w14, [x15, x14]
   700: ce0d0011                 add w14, w14, #3
   704: dffd0371                 cmp w14, #0xff
   708: 43000054                 b.lo 0x2cc
   712: cefd0351                 sub w14, w14, #0xff
   716: 2e6a6e38                 ldrb w14, [x17, x14]
   720: 8c010e4a                 eor w12, w12, w14
   724: 0e0940f9                 ldr x14, [x8, #0x10]
   728: ce696d38                 ldrb w14, [x14, x13]
   732: 0e010034                 cbz w14, 0x2fc
   736: ee696e38                 ldrb w14, [x15, x14]
   740: ce190311                 add w14, w14, #0xc6
   744: dffd0371                 cmp w14, #0xff
   748: 43000054                 b.lo 0x2f4
   752: cefd0351                 sub w14, w14, #0xff
   756: 2e6a6e38                 ldrb w14, [x17, x14]
   760: 8c010e4a                 eor w12, w12, w14
   764: 0e0d40f9                 ldr x14, [x8, #0x18]
   768: ce696d38                 ldrb w14, [x14, x13]
   772: 0e010034                 cbz w14, 0x324
   776: ee696e38                 ldrb w14, [x15, x14]
   780: ce050011                 add w14, w14, #1
   784: dffd0371                 cmp w14, #0xff
   788: 43000054                 b.lo 0x31c
   792: cefd0351                 sub w14, w14, #0xff
   796: 2e6a6e38                 ldrb w14, [x17, x14]
   800: 8c010e4a                 eor w12, w12, w14
   804: 0e1140f9                 ldr x14, [x8, #0x20]
   808: ce696d38                 ldrb w14, [x14, x13]
   812: 0e010034                 cbz w14, 0x34c
   816: ee696e38                 ldrb w14, [x15, x14]
   820: cef90311                 add w14, w14, #0xfe
   824: dffd0371                 cmp w14, #0xff
   828: 43000054                 b.lo 0x344
   832: cefd0351                 sub w14, w14, #0xff
   836: 2e6a6e38                 ldrb w14, [x17, x14]
   840: 8c010e4a                 eor w12, w12, w14
   844: 0e1540f9                 ldr x14, [x8, #0x28]
   848: ce696d38                 ldrb w14, [x14, x13]
   852: 0e010034                 cbz w14, 0x374
   856: ee696e38                 ldrb w14, [x15, x14]
   860: cec50011                 add w14, w14, #0x31
   864: dffd0371                 cmp w14, #0xff
   868: 43000054                 b.lo 0x36c
   872: cefd0351                 sub w14, w14, #0xff
   876: 2e6a6e38                 ldrb w14, [x17, x14]
   880: 8c010e4a                 eor w12, w12, w14
   884: 4c692d38                 strb w12, [x10, x13]
   888: ad050091                 add x13, x13, #1
   892: bf0102eb                 cmp x13, x2
   896: 03f8ff54                 b.lo 0x280
   900: 2a0540f9                 ldr x10, [x9, #8]
   904: ed030baa                 mov x13, x11
   908: 0c008052                 mov w12, #0
   912: 0e0140f9                 ldr x14, [x8]
   916: ce696d38                 ldrb w14, [x14, x13]
   920: 0e010034                 cbz w14, 0x3b8
   924: ee696e38                 ldrb w14, [x15, x14]
   928: ce710011                 add w14, w14, #0x1c
   932: dffd0371                 cmp w14, #0xff
   936: 43000054                 b.lo 0x3b0
   940: cefd0351                 sub w14, w14, #0xff
   944: 2e6a6e38                 ldrb w14, [x17, x14]
   948: 8c010e4a                 eor w12, w12, w14
   952: 0e0540f9                 ldr x14, [x8, #8]
   956: ce696d38                 ldrb w14, [x14, x13]
   960: 0e010034                 cbz w14, 0x3e0
   964: ee696e38                 ldrb w14, [x15, x14]
   968: ce7d0311                 add w14, w14, #0xdf
   972: dffd0371                 cmp w14, #0xff
   976: 43000054                 b.lo 0x3d8
   980: cefd0351                 sub w14, w14, #0xff
   984: 2e6a6e38                 ldrb w14, [x17, x14]
   988: 8c010e4a                 eor w12, w12, w14
   992: 0e0940f9                 ldr x14, [x8, #0x10]
   996: ce696d38                 ldrb w14, [x14, x13]
  1000: 0e010034                 cbz w14, 0x408
  1004: ee696e38                 ldrb w14, [x15, x14]
  1008: ce690011                 add w14, w14, #0x1a
  1012: dffd0371                 cmp w14, #0xff
  1016: 43000054                 b.lo 0x400
  1020: cefd0351                 sub w14, w14, #0xff
  1024: 2e6a6e38                 ldrb w14, [x17, x14]
  1028: 8c010e4a                 eor w12, w12, w14
  1032: 0e0d40f9                 ldr x14, [x8, #0x18]
  1036: ce696d38                 ldrb w14, [x14, x13]
  1040: 0e010034                 cbz w14, 0x430
  1044: ee696e38                 ldrb w14, [x15, x14]
  1048: ce650011                 add w14, w14, #0x19
  1052: dffd0371                 cmp w14, #0xff
  1056: 43000054                 b.lo 0x428
  1060: cefd0351                 sub w14, w14, #0xff
  1064: 2e6a6e38                 ldrb w14, [x17, x14]
  1068: 8c010e4a                 eor w12, w12, w14
  1072: 0e1140f9                 ldr x14, [x8, #0x20]
  1076: ce696d38                 ldrb w14, [x14, x13]
  1080: 0e010034                 cbz w14, 0x458
  1084: ee696e38                 ldrb w14, [x15, x14]
  1088: ce610011                 add w14, w14, #0x18
  1092: dffd0371                 cmp w14, #0xff
  1096: 43000054                 b.lo 0x450
  1100: cefd0351                 sub w14, w14, #0xff
  1104: 2e6a6e38                 ldrb w14, [x17, x14]
  1108: 8c010e4a                 eor w12, w12, w14
  1112: 0e1540f9                 ldr x14, [x8, #0x28]
  1116: ce696d38                 ldrb w14, [x14, x13]
  1120: 0e010034                 cbz w14, 0x480
  1124: ee696e38                 ldrb w14, [x15, x14]
  1128: cec50011                 add w14, w14, #0x31
  1132: dffd0371                 cmp w14, #0xff
  1136: 43000054                 b.lo 0x478
  1140: cefd0351                 sub w14, w14, #0xff
  1144: 2e6a6e38                 ldrb w14, [x17, x14]
  1148: 8c010e4a                 eor w12, w12, w14
  1152: 4c692d38                 strb w12, [x10, x13]
  1156: ad050091                 add x13, x13, #1
  1160: bf0102eb                 cmp x13, x2
  1164: 03f8ff54                 b.lo 0x38c
  1168: c0035fd6                 ret 
  1172: 00000000                 udf #0
  1176: 00000000                 udf #0
  1180: 00000000                 udf #0
  1184: 00000000                 udf #0
  1188: 00000000                 udf #0
  1192: 00000000                 udf #0
  1196: 00000000                 udf #0
  1200: 00000000                 udf #0
  1204: 00000000                 udf #0
  1208: 00000000                 udf #0
  1212: 00000000                 udf #0
