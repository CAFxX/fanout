; config_t1_885 target=sve2 twin=nt topo=raid6(8)/enc op=encode U=1
; code_size=1024 twins=["nt"] uops_per_byte=5.6875
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d01f0010                 adr x16, 0x400
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ed030caa                 mov x13, x12
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: f9c13825                 mov z25.b, #0xf
    36: 0e0f00b4                 cbz x14, 0x204
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
    80: 0f020091                 add x15, x16, #0
    84: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
    88: 0f020191                 add x15, x16, #0x40
    92: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
    96: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   100: 53323904                 and z19.d, z18.d, z25.d
   104: 54962c04                 lsr z20.b, z18.b, #4
   108: f5323305                 tbl z21.b, { z23.b }, z19.b
   112: 16333405                 tbl z22.b, { z24.b }, z20.b
   116: b132b604                 eor z17.d, z21.d, z22.d
   120: 0a0540f9                 ldr x10, [x8, #8]
   124: 4a010b8b                 add x10, x10, x11
   128: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   132: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   136: 0f020291                 add x15, x16, #0x80
   140: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   144: 0f020391                 add x15, x16, #0xc0
   148: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   152: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   156: 53323904                 and z19.d, z18.d, z25.d
   160: 54962c04                 lsr z20.b, z18.b, #4
   164: f5323305                 tbl z21.b, { z23.b }, z19.b
   168: 16333405                 tbl z22.b, { z24.b }, z20.b
   172: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   176: 0a0940f9                 ldr x10, [x8, #0x10]
   180: 4a010b8b                 add x10, x10, x11
   184: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   188: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   192: 0f020491                 add x15, x16, #0x100
   196: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   200: 0f020591                 add x15, x16, #0x140
   204: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   208: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   212: 53323904                 and z19.d, z18.d, z25.d
   216: 54962c04                 lsr z20.b, z18.b, #4
   220: f5323305                 tbl z21.b, { z23.b }, z19.b
   224: 16333405                 tbl z22.b, { z24.b }, z20.b
   228: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   232: 0a0d40f9                 ldr x10, [x8, #0x18]
   236: 4a010b8b                 add x10, x10, x11
   240: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   244: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   248: 0f020691                 add x15, x16, #0x180
   252: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   256: 0f020791                 add x15, x16, #0x1c0
   260: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   264: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   268: 53323904                 and z19.d, z18.d, z25.d
   272: 54962c04                 lsr z20.b, z18.b, #4
   276: f5323305                 tbl z21.b, { z23.b }, z19.b
   280: 16333405                 tbl z22.b, { z24.b }, z20.b
   284: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   288: 0a1140f9                 ldr x10, [x8, #0x20]
   292: 4a010b8b                 add x10, x10, x11
   296: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   300: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   304: 0f020891                 add x15, x16, #0x200
   308: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   312: 0f020991                 add x15, x16, #0x240
   316: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   320: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   324: 53323904                 and z19.d, z18.d, z25.d
   328: 54962c04                 lsr z20.b, z18.b, #4
   332: f5323305                 tbl z21.b, { z23.b }, z19.b
   336: 16333405                 tbl z22.b, { z24.b }, z20.b
   340: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   344: 0a1540f9                 ldr x10, [x8, #0x28]
   348: 4a010b8b                 add x10, x10, x11
   352: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   356: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   360: 0f020a91                 add x15, x16, #0x280
   364: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   368: 0f020b91                 add x15, x16, #0x2c0
   372: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   376: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   380: 53323904                 and z19.d, z18.d, z25.d
   384: 54962c04                 lsr z20.b, z18.b, #4
   388: f5323305                 tbl z21.b, { z23.b }, z19.b
   392: 16333405                 tbl z22.b, { z24.b }, z20.b
   396: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   400: 0a1940f9                 ldr x10, [x8, #0x30]
   404: 4a010b8b                 add x10, x10, x11
   408: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   412: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   416: 0f020c91                 add x15, x16, #0x300
   420: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   424: 0f020d91                 add x15, x16, #0x340
   428: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   432: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   436: 53323904                 and z19.d, z18.d, z25.d
   440: 54962c04                 lsr z20.b, z18.b, #4
   444: f5323305                 tbl z21.b, { z23.b }, z19.b
   448: 16333405                 tbl z22.b, { z24.b }, z20.b
   452: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   456: 0a1d40f9                 ldr x10, [x8, #0x38]
   460: 4a010b8b                 add x10, x10, x11
   464: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   468: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   472: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   476: 51021904                 eor z17.b, p0/m, z17.b, z18.b
   480: 2a0140f9                 ldr x10, [x9]
   484: 4a010b8b                 add x10, x10, x11
   488: 50e110e4                 stnt1b { z16.b }, p0, [x10]
   492: 2a0540f9                 ldr x10, [x9, #8]
   496: 4a010b8b                 add x10, x10, x11
   500: 51e110e4                 stnt1b { z17.b }, p0, [x10]
   504: 6b010d8b                 add x11, x11, x13
   508: ce0500f1                 subs x14, x14, #1
   512: 41f1ff54                 b.ne 0x28
   516: 7f0102eb                 cmp x11, x2
   520: 420e0054                 b.hs 0x3d0
   524: 0f020e91                 add x15, x16, #0x380
   528: 11021291                 add x17, x16, #0x480
   532: 2a0140f9                 ldr x10, [x9]
   536: ed030baa                 mov x13, x11
   540: 0c008052                 mov w12, #0
   544: 0e0140f9                 ldr x14, [x8]
   548: ce696d38                 ldrb w14, [x14, x13]
   552: 8c010e4a                 eor w12, w12, w14
   556: 0e0540f9                 ldr x14, [x8, #8]
   560: ce696d38                 ldrb w14, [x14, x13]
   564: 8c010e4a                 eor w12, w12, w14
   568: 0e0940f9                 ldr x14, [x8, #0x10]
   572: ce696d38                 ldrb w14, [x14, x13]
   576: 8c010e4a                 eor w12, w12, w14
   580: 0e0d40f9                 ldr x14, [x8, #0x18]
   584: ce696d38                 ldrb w14, [x14, x13]
   588: 8c010e4a                 eor w12, w12, w14
   592: 0e1140f9                 ldr x14, [x8, #0x20]
   596: ce696d38                 ldrb w14, [x14, x13]
   600: 8c010e4a                 eor w12, w12, w14
   604: 0e1540f9                 ldr x14, [x8, #0x28]
   608: ce696d38                 ldrb w14, [x14, x13]
   612: 8c010e4a                 eor w12, w12, w14
   616: 0e1940f9                 ldr x14, [x8, #0x30]
   620: ce696d38                 ldrb w14, [x14, x13]
   624: 8c010e4a                 eor w12, w12, w14
   628: 0e1d40f9                 ldr x14, [x8, #0x38]
   632: ce696d38                 ldrb w14, [x14, x13]
   636: 8c010e4a                 eor w12, w12, w14
   640: 4c692d38                 strb w12, [x10, x13]
   644: ad050091                 add x13, x13, #1
   648: bf0102eb                 cmp x13, x2
   652: 83fcff54                 b.lo 0x21c
   656: 2a0540f9                 ldr x10, [x9, #8]
   660: ed030baa                 mov x13, x11
   664: 0c008052                 mov w12, #0
   668: 0e0140f9                 ldr x14, [x8]
   672: ce696d38                 ldrb w14, [x14, x13]
   676: 0e010034                 cbz w14, 0x2c4
   680: ee696e38                 ldrb w14, [x15, x14]
   684: ce990311                 add w14, w14, #0xe6
   688: dffd0371                 cmp w14, #0xff
   692: 43000054                 b.lo 0x2bc
   696: cefd0351                 sub w14, w14, #0xff
   700: 2e6a6e38                 ldrb w14, [x17, x14]
   704: 8c010e4a                 eor w12, w12, w14
   708: 0e0540f9                 ldr x14, [x8, #8]
   712: ce696d38                 ldrb w14, [x14, x13]
   716: 0e010034                 cbz w14, 0x2ec
   720: ee696e38                 ldrb w14, [x15, x14]
   724: ce350311                 add w14, w14, #0xcd
   728: dffd0371                 cmp w14, #0xff
   732: 43000054                 b.lo 0x2e4
   736: cefd0351                 sub w14, w14, #0xff
   740: 2e6a6e38                 ldrb w14, [x17, x14]
   744: 8c010e4a                 eor w12, w12, w14
   748: 0e0940f9                 ldr x14, [x8, #0x10]
   752: ce696d38                 ldrb w14, [x14, x13]
   756: 0e010034                 cbz w14, 0x314
   760: ee696e38                 ldrb w14, [x15, x14]
   764: ce9d0311                 add w14, w14, #0xe7
   768: dffd0371                 cmp w14, #0xff
   772: 43000054                 b.lo 0x30c
   776: cefd0351                 sub w14, w14, #0xff
   780: 2e6a6e38                 ldrb w14, [x17, x14]
   784: 8c010e4a                 eor w12, w12, w14
   788: 0e0d40f9                 ldr x14, [x8, #0x18]
   792: ce696d38                 ldrb w14, [x14, x13]
   796: 0e010034                 cbz w14, 0x33c
   800: ee696e38                 ldrb w14, [x15, x14]
   804: ce390311                 add w14, w14, #0xce
   808: dffd0371                 cmp w14, #0xff
   812: 43000054                 b.lo 0x334
   816: cefd0351                 sub w14, w14, #0xff
   820: 2e6a6e38                 ldrb w14, [x17, x14]
   824: 8c010e4a                 eor w12, w12, w14
   828: 0e1140f9                 ldr x14, [x8, #0x20]
   832: ce696d38                 ldrb w14, [x14, x13]
   836: 0e010034                 cbz w14, 0x364
   840: ee696e38                 ldrb w14, [x15, x14]
   844: ce2d0111                 add w14, w14, #0x4b
   848: dffd0371                 cmp w14, #0xff
   852: 43000054                 b.lo 0x35c
   856: cefd0351                 sub w14, w14, #0xff
   860: 2e6a6e38                 ldrb w14, [x17, x14]
   864: 8c010e4a                 eor w12, w12, w14
   868: 0e1540f9                 ldr x14, [x8, #0x28]
   872: ce696d38                 ldrb w14, [x14, x13]
   876: 0e010034                 cbz w14, 0x38c
   880: ee696e38                 ldrb w14, [x15, x14]
   884: cec90011                 add w14, w14, #0x32
   888: dffd0371                 cmp w14, #0xff
   892: 43000054                 b.lo 0x384
   896: cefd0351                 sub w14, w14, #0xff
   900: 2e6a6e38                 ldrb w14, [x17, x14]
   904: 8c010e4a                 eor w12, w12, w14
   908: 0e1940f9                 ldr x14, [x8, #0x30]
   912: ce696d38                 ldrb w14, [x14, x13]
   916: 0e010034                 cbz w14, 0x3b4
   920: ee696e38                 ldrb w14, [x15, x14]
   924: ce650011                 add w14, w14, #0x19
   928: dffd0371                 cmp w14, #0xff
   932: 43000054                 b.lo 0x3ac
   936: cefd0351                 sub w14, w14, #0xff
   940: 2e6a6e38                 ldrb w14, [x17, x14]
   944: 8c010e4a                 eor w12, w12, w14
   948: 0e1d40f9                 ldr x14, [x8, #0x38]
   952: ce696d38                 ldrb w14, [x14, x13]
   956: 8c010e4a                 eor w12, w12, w14
   960: 4c692d38                 strb w12, [x10, x13]
   964: ad050091                 add x13, x13, #1
   968: bf0102eb                 cmp x13, x2
   972: 63f6ff54                 b.lo 0x298
   976: c0035fd6                 ret 
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
