; config_t1_895 target=sve2 twin=pipe topo=raid6(8)/enc op=encode U=1
; code_size=1024 twins=["pipe"] uops_per_byte=5.25
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d01f0010                 adr x16, 0x400
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ed030caa                 mov x13, x12
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: fac13825                 mov z26.b, #0xf
    36: 6e0e00b4                 cbz x14, 0x1f0
    40: 0a0140f9                 ldr x10, [x8]
    44: 4a010b8b                 add x10, x10, x11
    48: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 70327304                 mov z16.d, z19.d
    68: 0f020091                 add x15, x16, #0
    72: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
    76: 0f020191                 add x15, x16, #0x40
    80: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
    84: 74323a04                 and z20.d, z19.d, z26.d
    88: 75962c04                 lsr z21.b, z19.b, #4
    92: 16333405                 tbl z22.b, { z24.b }, z20.b
    96: 37333505                 tbl z23.b, { z25.b }, z21.b
   100: d132b704                 eor z17.d, z22.d, z23.d
   104: 0a0540f9                 ldr x10, [x8, #8]
   108: 4a010b8b                 add x10, x10, x11
   112: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   116: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   120: 0f020291                 add x15, x16, #0x80
   124: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   128: 0f020391                 add x15, x16, #0xc0
   132: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   136: 54323a04                 and z20.d, z18.d, z26.d
   140: 55962c04                 lsr z21.b, z18.b, #4
   144: 16333405                 tbl z22.b, { z24.b }, z20.b
   148: 37333505                 tbl z23.b, { z25.b }, z21.b
   152: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   156: 0a0940f9                 ldr x10, [x8, #0x10]
   160: 4a010b8b                 add x10, x10, x11
   164: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   168: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   172: 0f020491                 add x15, x16, #0x100
   176: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   180: 0f020591                 add x15, x16, #0x140
   184: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   188: 54323a04                 and z20.d, z18.d, z26.d
   192: 55962c04                 lsr z21.b, z18.b, #4
   196: 16333405                 tbl z22.b, { z24.b }, z20.b
   200: 37333505                 tbl z23.b, { z25.b }, z21.b
   204: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   208: 0a0d40f9                 ldr x10, [x8, #0x18]
   212: 4a010b8b                 add x10, x10, x11
   216: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   220: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   224: 0f020691                 add x15, x16, #0x180
   228: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   232: 0f020791                 add x15, x16, #0x1c0
   236: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   240: 54323a04                 and z20.d, z18.d, z26.d
   244: 55962c04                 lsr z21.b, z18.b, #4
   248: 16333405                 tbl z22.b, { z24.b }, z20.b
   252: 37333505                 tbl z23.b, { z25.b }, z21.b
   256: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   260: 0a1140f9                 ldr x10, [x8, #0x20]
   264: 4a010b8b                 add x10, x10, x11
   268: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   272: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   276: 0f020891                 add x15, x16, #0x200
   280: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   284: 0f020991                 add x15, x16, #0x240
   288: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   292: 54323a04                 and z20.d, z18.d, z26.d
   296: 55962c04                 lsr z21.b, z18.b, #4
   300: 16333405                 tbl z22.b, { z24.b }, z20.b
   304: 37333505                 tbl z23.b, { z25.b }, z21.b
   308: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   312: 0a1540f9                 ldr x10, [x8, #0x28]
   316: 4a010b8b                 add x10, x10, x11
   320: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   324: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   328: 0f020a91                 add x15, x16, #0x280
   332: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   336: 0f020b91                 add x15, x16, #0x2c0
   340: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   344: 54323a04                 and z20.d, z18.d, z26.d
   348: 55962c04                 lsr z21.b, z18.b, #4
   352: 16333405                 tbl z22.b, { z24.b }, z20.b
   356: 37333505                 tbl z23.b, { z25.b }, z21.b
   360: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   364: 0a1940f9                 ldr x10, [x8, #0x30]
   368: 4a010b8b                 add x10, x10, x11
   372: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   376: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   380: 0f020c91                 add x15, x16, #0x300
   384: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   388: 0f020d91                 add x15, x16, #0x340
   392: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   396: 54323a04                 and z20.d, z18.d, z26.d
   400: 55962c04                 lsr z21.b, z18.b, #4
   404: 16333405                 tbl z22.b, { z24.b }, z20.b
   408: 37333505                 tbl z23.b, { z25.b }, z21.b
   412: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   416: 0a1d40f9                 ldr x10, [x8, #0x38]
   420: 4a010b8b                 add x10, x10, x11
   424: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   428: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   432: 51021904                 eor z17.b, p0/m, z17.b, z18.b
   436: df0500f1                 cmp x14, #1
   440: ad000054                 b.le 0x1cc
   444: 0a0140f9                 ldr x10, [x8]
   448: 4a010b8b                 add x10, x10, x11
   452: 4a410091                 add x10, x10, #0x10
   456: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   460: 2a0140f9                 ldr x10, [x9]
   464: 4a010b8b                 add x10, x10, x11
   468: 50e100e4                 st1b { z16.b }, p0, [x10]
   472: 2a0540f9                 ldr x10, [x9, #8]
   476: 4a010b8b                 add x10, x10, x11
   480: 51e100e4                 st1b { z17.b }, p0, [x10]
   484: 6b010d8b                 add x11, x11, x13
   488: ce0500f1                 subs x14, x14, #1
   492: 41f2ff54                 b.ne 0x34
   496: 7f0102eb                 cmp x11, x2
   500: 420e0054                 b.hs 0x3bc
   504: 0f020e91                 add x15, x16, #0x380
   508: 11021291                 add x17, x16, #0x480
   512: 2a0140f9                 ldr x10, [x9]
   516: ed030baa                 mov x13, x11
   520: 0c008052                 mov w12, #0
   524: 0e0140f9                 ldr x14, [x8]
   528: ce696d38                 ldrb w14, [x14, x13]
   532: 8c010e4a                 eor w12, w12, w14
   536: 0e0540f9                 ldr x14, [x8, #8]
   540: ce696d38                 ldrb w14, [x14, x13]
   544: 8c010e4a                 eor w12, w12, w14
   548: 0e0940f9                 ldr x14, [x8, #0x10]
   552: ce696d38                 ldrb w14, [x14, x13]
   556: 8c010e4a                 eor w12, w12, w14
   560: 0e0d40f9                 ldr x14, [x8, #0x18]
   564: ce696d38                 ldrb w14, [x14, x13]
   568: 8c010e4a                 eor w12, w12, w14
   572: 0e1140f9                 ldr x14, [x8, #0x20]
   576: ce696d38                 ldrb w14, [x14, x13]
   580: 8c010e4a                 eor w12, w12, w14
   584: 0e1540f9                 ldr x14, [x8, #0x28]
   588: ce696d38                 ldrb w14, [x14, x13]
   592: 8c010e4a                 eor w12, w12, w14
   596: 0e1940f9                 ldr x14, [x8, #0x30]
   600: ce696d38                 ldrb w14, [x14, x13]
   604: 8c010e4a                 eor w12, w12, w14
   608: 0e1d40f9                 ldr x14, [x8, #0x38]
   612: ce696d38                 ldrb w14, [x14, x13]
   616: 8c010e4a                 eor w12, w12, w14
   620: 4c692d38                 strb w12, [x10, x13]
   624: ad050091                 add x13, x13, #1
   628: bf0102eb                 cmp x13, x2
   632: 83fcff54                 b.lo 0x208
   636: 2a0540f9                 ldr x10, [x9, #8]
   640: ed030baa                 mov x13, x11
   644: 0c008052                 mov w12, #0
   648: 0e0140f9                 ldr x14, [x8]
   652: ce696d38                 ldrb w14, [x14, x13]
   656: 0e010034                 cbz w14, 0x2b0
   660: ee696e38                 ldrb w14, [x15, x14]
   664: ce990311                 add w14, w14, #0xe6
   668: dffd0371                 cmp w14, #0xff
   672: 43000054                 b.lo 0x2a8
   676: cefd0351                 sub w14, w14, #0xff
   680: 2e6a6e38                 ldrb w14, [x17, x14]
   684: 8c010e4a                 eor w12, w12, w14
   688: 0e0540f9                 ldr x14, [x8, #8]
   692: ce696d38                 ldrb w14, [x14, x13]
   696: 0e010034                 cbz w14, 0x2d8
   700: ee696e38                 ldrb w14, [x15, x14]
   704: ce350311                 add w14, w14, #0xcd
   708: dffd0371                 cmp w14, #0xff
   712: 43000054                 b.lo 0x2d0
   716: cefd0351                 sub w14, w14, #0xff
   720: 2e6a6e38                 ldrb w14, [x17, x14]
   724: 8c010e4a                 eor w12, w12, w14
   728: 0e0940f9                 ldr x14, [x8, #0x10]
   732: ce696d38                 ldrb w14, [x14, x13]
   736: 0e010034                 cbz w14, 0x300
   740: ee696e38                 ldrb w14, [x15, x14]
   744: ce9d0311                 add w14, w14, #0xe7
   748: dffd0371                 cmp w14, #0xff
   752: 43000054                 b.lo 0x2f8
   756: cefd0351                 sub w14, w14, #0xff
   760: 2e6a6e38                 ldrb w14, [x17, x14]
   764: 8c010e4a                 eor w12, w12, w14
   768: 0e0d40f9                 ldr x14, [x8, #0x18]
   772: ce696d38                 ldrb w14, [x14, x13]
   776: 0e010034                 cbz w14, 0x328
   780: ee696e38                 ldrb w14, [x15, x14]
   784: ce390311                 add w14, w14, #0xce
   788: dffd0371                 cmp w14, #0xff
   792: 43000054                 b.lo 0x320
   796: cefd0351                 sub w14, w14, #0xff
   800: 2e6a6e38                 ldrb w14, [x17, x14]
   804: 8c010e4a                 eor w12, w12, w14
   808: 0e1140f9                 ldr x14, [x8, #0x20]
   812: ce696d38                 ldrb w14, [x14, x13]
   816: 0e010034                 cbz w14, 0x350
   820: ee696e38                 ldrb w14, [x15, x14]
   824: ce2d0111                 add w14, w14, #0x4b
   828: dffd0371                 cmp w14, #0xff
   832: 43000054                 b.lo 0x348
   836: cefd0351                 sub w14, w14, #0xff
   840: 2e6a6e38                 ldrb w14, [x17, x14]
   844: 8c010e4a                 eor w12, w12, w14
   848: 0e1540f9                 ldr x14, [x8, #0x28]
   852: ce696d38                 ldrb w14, [x14, x13]
   856: 0e010034                 cbz w14, 0x378
   860: ee696e38                 ldrb w14, [x15, x14]
   864: cec90011                 add w14, w14, #0x32
   868: dffd0371                 cmp w14, #0xff
   872: 43000054                 b.lo 0x370
   876: cefd0351                 sub w14, w14, #0xff
   880: 2e6a6e38                 ldrb w14, [x17, x14]
   884: 8c010e4a                 eor w12, w12, w14
   888: 0e1940f9                 ldr x14, [x8, #0x30]
   892: ce696d38                 ldrb w14, [x14, x13]
   896: 0e010034                 cbz w14, 0x3a0
   900: ee696e38                 ldrb w14, [x15, x14]
   904: ce650011                 add w14, w14, #0x19
   908: dffd0371                 cmp w14, #0xff
   912: 43000054                 b.lo 0x398
   916: cefd0351                 sub w14, w14, #0xff
   920: 2e6a6e38                 ldrb w14, [x17, x14]
   924: 8c010e4a                 eor w12, w12, w14
   928: 0e1d40f9                 ldr x14, [x8, #0x38]
   932: ce696d38                 ldrb w14, [x14, x13]
   936: 8c010e4a                 eor w12, w12, w14
   940: 4c692d38                 strb w12, [x10, x13]
   944: ad050091                 add x13, x13, #1
   948: bf0102eb                 cmp x13, x2
   952: 63f6ff54                 b.lo 0x284
   956: c0035fd6                 ret 
   960: 00000000                 udf #0
   964: 00000000                 udf #0
   968: 00000000                 udf #0
   972: 00000000                 udf #0
   976: 00000000                 udf #0
   980: 00000000                 udf #0
   984: 00000000                 udf #0
   988: 00000000                 udf #0
   992: 00000000                 udf #0
   996: 00000000                 udf #0
  1000: 00000000                 udf #0
  1004: 00000000                 udf #0
  1008: 00000000                 udf #0
  1012: 00000000                 udf #0
  1016: 00000000                 udf #0
  1020: 00000000                 udf #0
