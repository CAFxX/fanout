; config_t1_1052 target=sve2 twin=rec_cf topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1088 twins=[] uops_per_byte=5.3125
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0210010                 adr x16, 0x440
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ed030caa                 mov x13, x12
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: f9c13825                 mov z25.b, #0xf
    36: ce0e00b4                 cbz x14, 0x1fc
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
    76: 50327204                 mov z16.d, z18.d
    80: 0a0540f9                 ldr x10, [x8, #8]
    84: 4a010b8b                 add x10, x10, x11
    88: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
    92: 50021904                 eor z16.b, p0/m, z16.b, z18.b
    96: 0a0940f9                 ldr x10, [x8, #0x10]
   100: 4a010b8b                 add x10, x10, x11
   104: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   108: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   112: 0a0d40f9                 ldr x10, [x8, #0x18]
   116: 4a010b8b                 add x10, x10, x11
   120: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   124: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   128: 0a1140f9                 ldr x10, [x8, #0x20]
   132: 4a010b8b                 add x10, x10, x11
   136: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   140: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   144: 0f020091                 add x15, x16, #0
   148: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   152: 0f020191                 add x15, x16, #0x40
   156: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   160: 0a0140f9                 ldr x10, [x8]
   164: 4a010b8b                 add x10, x10, x11
   168: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   172: 53323904                 and z19.d, z18.d, z25.d
   176: 54962c04                 lsr z20.b, z18.b, #4
   180: f5323305                 tbl z21.b, { z23.b }, z19.b
   184: 16333405                 tbl z22.b, { z24.b }, z20.b
   188: b132b604                 eor z17.d, z21.d, z22.d
   192: 0f020291                 add x15, x16, #0x80
   196: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   200: 0f020391                 add x15, x16, #0xc0
   204: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   208: 0a0540f9                 ldr x10, [x8, #8]
   212: 4a010b8b                 add x10, x10, x11
   216: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   220: 53323904                 and z19.d, z18.d, z25.d
   224: 54962c04                 lsr z20.b, z18.b, #4
   228: f5323305                 tbl z21.b, { z23.b }, z19.b
   232: 16333405                 tbl z22.b, { z24.b }, z20.b
   236: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   240: 0f020491                 add x15, x16, #0x100
   244: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   248: 0f020591                 add x15, x16, #0x140
   252: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   256: 0a0940f9                 ldr x10, [x8, #0x10]
   260: 4a010b8b                 add x10, x10, x11
   264: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   268: 53323904                 and z19.d, z18.d, z25.d
   272: 54962c04                 lsr z20.b, z18.b, #4
   276: f5323305                 tbl z21.b, { z23.b }, z19.b
   280: 16333405                 tbl z22.b, { z24.b }, z20.b
   284: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   288: 0a0d40f9                 ldr x10, [x8, #0x18]
   292: 4a010b8b                 add x10, x10, x11
   296: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   300: 51021904                 eor z17.b, p0/m, z17.b, z18.b
   304: 0a1540f9                 ldr x10, [x8, #0x28]
   308: 4a010b8b                 add x10, x10, x11
   312: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   316: 51021904                 eor z17.b, p0/m, z17.b, z18.b
   320: 0f020691                 add x15, x16, #0x180
   324: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   328: 0f020791                 add x15, x16, #0x1c0
   332: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   336: 13323904                 and z19.d, z16.d, z25.d
   340: 14962c04                 lsr z20.b, z16.b, #4
   344: f5323305                 tbl z21.b, { z23.b }, z19.b
   348: 16333405                 tbl z22.b, { z24.b }, z20.b
   352: b232b604                 eor z18.d, z21.d, z22.d
   356: 32021904                 eor z18.b, p0/m, z18.b, z17.b
   360: 0f020891                 add x15, x16, #0x200
   364: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   368: 0f020991                 add x15, x16, #0x240
   372: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   376: 13323904                 and z19.d, z16.d, z25.d
   380: 14962c04                 lsr z20.b, z16.b, #4
   384: f5323305                 tbl z21.b, { z23.b }, z19.b
   388: 16333405                 tbl z22.b, { z24.b }, z20.b
   392: b532b604                 eor z21.d, z21.d, z22.d
   396: 35021904                 eor z21.b, p0/m, z21.b, z17.b
   400: 0f020a91                 add x15, x16, #0x280
   404: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   408: 0f020b91                 add x15, x16, #0x2c0
   412: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   416: b3323904                 and z19.d, z21.d, z25.d
   420: b4962c04                 lsr z20.b, z21.b, #4
   424: f5323305                 tbl z21.b, { z23.b }, z19.b
   428: 16333405                 tbl z22.b, { z24.b }, z20.b
   432: b132b604                 eor z17.d, z21.d, z22.d
   436: 0f020a91                 add x15, x16, #0x280
   440: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   444: 0f020b91                 add x15, x16, #0x2c0
   448: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   452: 53323904                 and z19.d, z18.d, z25.d
   456: 54962c04                 lsr z20.b, z18.b, #4
   460: f5323305                 tbl z21.b, { z23.b }, z19.b
   464: 16333405                 tbl z22.b, { z24.b }, z20.b
   468: b032b604                 eor z16.d, z21.d, z22.d
   472: 2a0140f9                 ldr x10, [x9]
   476: 4a010b8b                 add x10, x10, x11
   480: 50e100e4                 st1b { z16.b }, p0, [x10]
   484: 2a0540f9                 ldr x10, [x9, #8]
   488: 4a010b8b                 add x10, x10, x11
   492: 51e100e4                 st1b { z17.b }, p0, [x10]
   496: 6b010d8b                 add x11, x11, x13
   500: ce0500f1                 subs x14, x14, #1
   504: 81f1ff54                 b.ne 0x28
   508: 7f0102eb                 cmp x11, x2
   512: 22110054                 b.hs 0x424
   516: 0f020c91                 add x15, x16, #0x300
   520: 11021091                 add x17, x16, #0x400
   524: 2a0140f9                 ldr x10, [x9]
   528: ed030baa                 mov x13, x11
   532: 0c008052                 mov w12, #0
   536: 0e0140f9                 ldr x14, [x8]
   540: ce696d38                 ldrb w14, [x14, x13]
   544: 0e010034                 cbz w14, 0x240
   548: ee696e38                 ldrb w14, [x15, x14]
   552: ce050311                 add w14, w14, #0xc1
   556: dffd0371                 cmp w14, #0xff
   560: 43000054                 b.lo 0x238
   564: cefd0351                 sub w14, w14, #0xff
   568: 2e6a6e38                 ldrb w14, [x17, x14]
   572: 8c010e4a                 eor w12, w12, w14
   576: 0e0540f9                 ldr x14, [x8, #8]
   580: ce696d38                 ldrb w14, [x14, x13]
   584: 0e010034                 cbz w14, 0x268
   588: ee696e38                 ldrb w14, [x15, x14]
   592: ce0d0011                 add w14, w14, #3
   596: dffd0371                 cmp w14, #0xff
   600: 43000054                 b.lo 0x260
   604: cefd0351                 sub w14, w14, #0xff
   608: 2e6a6e38                 ldrb w14, [x17, x14]
   612: 8c010e4a                 eor w12, w12, w14
   616: 0e0940f9                 ldr x14, [x8, #0x10]
   620: ce696d38                 ldrb w14, [x14, x13]
   624: 0e010034                 cbz w14, 0x290
   628: ee696e38                 ldrb w14, [x15, x14]
   632: ce190311                 add w14, w14, #0xc6
   636: dffd0371                 cmp w14, #0xff
   640: 43000054                 b.lo 0x288
   644: cefd0351                 sub w14, w14, #0xff
   648: 2e6a6e38                 ldrb w14, [x17, x14]
   652: 8c010e4a                 eor w12, w12, w14
   656: 0e0d40f9                 ldr x14, [x8, #0x18]
   660: ce696d38                 ldrb w14, [x14, x13]
   664: 0e010034                 cbz w14, 0x2b8
   668: ee696e38                 ldrb w14, [x15, x14]
   672: ce050011                 add w14, w14, #1
   676: dffd0371                 cmp w14, #0xff
   680: 43000054                 b.lo 0x2b0
   684: cefd0351                 sub w14, w14, #0xff
   688: 2e6a6e38                 ldrb w14, [x17, x14]
   692: 8c010e4a                 eor w12, w12, w14
   696: 0e1140f9                 ldr x14, [x8, #0x20]
   700: ce696d38                 ldrb w14, [x14, x13]
   704: 0e010034                 cbz w14, 0x2e0
   708: ee696e38                 ldrb w14, [x15, x14]
   712: cef90311                 add w14, w14, #0xfe
   716: dffd0371                 cmp w14, #0xff
   720: 43000054                 b.lo 0x2d8
   724: cefd0351                 sub w14, w14, #0xff
   728: 2e6a6e38                 ldrb w14, [x17, x14]
   732: 8c010e4a                 eor w12, w12, w14
   736: 0e1540f9                 ldr x14, [x8, #0x28]
   740: ce696d38                 ldrb w14, [x14, x13]
   744: 0e010034                 cbz w14, 0x308
   748: ee696e38                 ldrb w14, [x15, x14]
   752: cec50011                 add w14, w14, #0x31
   756: dffd0371                 cmp w14, #0xff
   760: 43000054                 b.lo 0x300
   764: cefd0351                 sub w14, w14, #0xff
   768: 2e6a6e38                 ldrb w14, [x17, x14]
   772: 8c010e4a                 eor w12, w12, w14
   776: 4c692d38                 strb w12, [x10, x13]
   780: ad050091                 add x13, x13, #1
   784: bf0102eb                 cmp x13, x2
   788: 03f8ff54                 b.lo 0x214
   792: 2a0540f9                 ldr x10, [x9, #8]
   796: ed030baa                 mov x13, x11
   800: 0c008052                 mov w12, #0
   804: 0e0140f9                 ldr x14, [x8]
   808: ce696d38                 ldrb w14, [x14, x13]
   812: 0e010034                 cbz w14, 0x34c
   816: ee696e38                 ldrb w14, [x15, x14]
   820: ce710011                 add w14, w14, #0x1c
   824: dffd0371                 cmp w14, #0xff
   828: 43000054                 b.lo 0x344
   832: cefd0351                 sub w14, w14, #0xff
   836: 2e6a6e38                 ldrb w14, [x17, x14]
   840: 8c010e4a                 eor w12, w12, w14
   844: 0e0540f9                 ldr x14, [x8, #8]
   848: ce696d38                 ldrb w14, [x14, x13]
   852: 0e010034                 cbz w14, 0x374
   856: ee696e38                 ldrb w14, [x15, x14]
   860: ce7d0311                 add w14, w14, #0xdf
   864: dffd0371                 cmp w14, #0xff
   868: 43000054                 b.lo 0x36c
   872: cefd0351                 sub w14, w14, #0xff
   876: 2e6a6e38                 ldrb w14, [x17, x14]
   880: 8c010e4a                 eor w12, w12, w14
   884: 0e0940f9                 ldr x14, [x8, #0x10]
   888: ce696d38                 ldrb w14, [x14, x13]
   892: 0e010034                 cbz w14, 0x39c
   896: ee696e38                 ldrb w14, [x15, x14]
   900: ce690011                 add w14, w14, #0x1a
   904: dffd0371                 cmp w14, #0xff
   908: 43000054                 b.lo 0x394
   912: cefd0351                 sub w14, w14, #0xff
   916: 2e6a6e38                 ldrb w14, [x17, x14]
   920: 8c010e4a                 eor w12, w12, w14
   924: 0e0d40f9                 ldr x14, [x8, #0x18]
   928: ce696d38                 ldrb w14, [x14, x13]
   932: 0e010034                 cbz w14, 0x3c4
   936: ee696e38                 ldrb w14, [x15, x14]
   940: ce650011                 add w14, w14, #0x19
   944: dffd0371                 cmp w14, #0xff
   948: 43000054                 b.lo 0x3bc
   952: cefd0351                 sub w14, w14, #0xff
   956: 2e6a6e38                 ldrb w14, [x17, x14]
   960: 8c010e4a                 eor w12, w12, w14
   964: 0e1140f9                 ldr x14, [x8, #0x20]
   968: ce696d38                 ldrb w14, [x14, x13]
   972: 0e010034                 cbz w14, 0x3ec
   976: ee696e38                 ldrb w14, [x15, x14]
   980: ce610011                 add w14, w14, #0x18
   984: dffd0371                 cmp w14, #0xff
   988: 43000054                 b.lo 0x3e4
   992: cefd0351                 sub w14, w14, #0xff
   996: 2e6a6e38                 ldrb w14, [x17, x14]
  1000: 8c010e4a                 eor w12, w12, w14
  1004: 0e1540f9                 ldr x14, [x8, #0x28]
  1008: ce696d38                 ldrb w14, [x14, x13]
  1012: 0e010034                 cbz w14, 0x414
  1016: ee696e38                 ldrb w14, [x15, x14]
  1020: cec50011                 add w14, w14, #0x31
  1024: dffd0371                 cmp w14, #0xff
  1028: 43000054                 b.lo 0x40c
  1032: cefd0351                 sub w14, w14, #0xff
  1036: 2e6a6e38                 ldrb w14, [x17, x14]
  1040: 8c010e4a                 eor w12, w12, w14
  1044: 4c692d38                 strb w12, [x10, x13]
  1048: ad050091                 add x13, x13, #1
  1052: bf0102eb                 cmp x13, x2
  1056: 03f8ff54                 b.lo 0x320
  1060: c0035fd6                 ret 
  1064: 00000000                 udf #0
  1068: 00000000                 udf #0
  1072: 00000000                 udf #0
  1076: 00000000                 udf #0
  1080: 00000000                 udf #0
  1084: 00000000                 udf #0
