; config_t1_1338 target=neon twin=rec_cf topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1088 twins=[] uops_per_byte=4.875
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: e80300aa                 mov x8, x0
     8: e90301aa                 mov x9, x1
    12: b0210010                 adr x16, 0x440
    16: eb031faa                 mov x11, xzr
    20: 0e0280d2                 mov x14, #0x10
    24: 4d08ce9a                 udiv x13, x2, x14
    28: 2d0e00b4                 cbz x13, 0x1e0
    32: e9e5004f                 movi v9.16b, #0xf
    36: 1f2003d5                 nop 
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 4201c03d                 ldr q2, [x10]
    76: 401ca24e                 mov v0.16b, v2.16b
    80: 0a0540f9                 ldr x10, [x8, #8]
    84: 4a010b8b                 add x10, x10, x11
    88: 4201c03d                 ldr q2, [x10]
    92: 001c226e                 eor v0.16b, v0.16b, v2.16b
    96: 0a0940f9                 ldr x10, [x8, #0x10]
   100: 4a010b8b                 add x10, x10, x11
   104: 4201c03d                 ldr q2, [x10]
   108: 001c226e                 eor v0.16b, v0.16b, v2.16b
   112: 0a0d40f9                 ldr x10, [x8, #0x18]
   116: 4a010b8b                 add x10, x10, x11
   120: 4201c03d                 ldr q2, [x10]
   124: 001c226e                 eor v0.16b, v0.16b, v2.16b
   128: 0a1140f9                 ldr x10, [x8, #0x20]
   132: 4a010b8b                 add x10, x10, x11
   136: 4201c03d                 ldr q2, [x10]
   140: 001c226e                 eor v0.16b, v0.16b, v2.16b
   144: 0702c03d                 ldr q7, [x16]
   148: 0812c03d                 ldr q8, [x16, #0x40]
   152: 0a0140f9                 ldr x10, [x8]
   156: 4a010b8b                 add x10, x10, x11
   160: 4201c03d                 ldr q2, [x10]
   164: 431c294e                 and v3.16b, v2.16b, v9.16b
   168: 44040c6f                 ushr v4.16b, v2.16b, #4
   172: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   176: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   180: a11ca54e                 mov v1.16b, v5.16b
   184: 211c266e                 eor v1.16b, v1.16b, v6.16b
   188: 0722c03d                 ldr q7, [x16, #0x80]
   192: 0832c03d                 ldr q8, [x16, #0xc0]
   196: 0a0540f9                 ldr x10, [x8, #8]
   200: 4a010b8b                 add x10, x10, x11
   204: 4201c03d                 ldr q2, [x10]
   208: 431c294e                 and v3.16b, v2.16b, v9.16b
   212: 44040c6f                 ushr v4.16b, v2.16b, #4
   216: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   220: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   224: 211c256e                 eor v1.16b, v1.16b, v5.16b
   228: 211c266e                 eor v1.16b, v1.16b, v6.16b
   232: 0742c03d                 ldr q7, [x16, #0x100]
   236: 0852c03d                 ldr q8, [x16, #0x140]
   240: 0a0940f9                 ldr x10, [x8, #0x10]
   244: 4a010b8b                 add x10, x10, x11
   248: 4201c03d                 ldr q2, [x10]
   252: 431c294e                 and v3.16b, v2.16b, v9.16b
   256: 44040c6f                 ushr v4.16b, v2.16b, #4
   260: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   264: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   268: 211c256e                 eor v1.16b, v1.16b, v5.16b
   272: 211c266e                 eor v1.16b, v1.16b, v6.16b
   276: 0a0d40f9                 ldr x10, [x8, #0x18]
   280: 4a010b8b                 add x10, x10, x11
   284: 4201c03d                 ldr q2, [x10]
   288: 211c226e                 eor v1.16b, v1.16b, v2.16b
   292: 0a1540f9                 ldr x10, [x8, #0x28]
   296: 4a010b8b                 add x10, x10, x11
   300: 4201c03d                 ldr q2, [x10]
   304: 211c226e                 eor v1.16b, v1.16b, v2.16b
   308: 0762c03d                 ldr q7, [x16, #0x180]
   312: 0872c03d                 ldr q8, [x16, #0x1c0]
   316: 031c294e                 and v3.16b, v0.16b, v9.16b
   320: 04040c6f                 ushr v4.16b, v0.16b, #4
   324: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   328: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   332: a21ca54e                 mov v2.16b, v5.16b
   336: 421c266e                 eor v2.16b, v2.16b, v6.16b
   340: 421c216e                 eor v2.16b, v2.16b, v1.16b
   344: 0782c03d                 ldr q7, [x16, #0x200]
   348: 0892c03d                 ldr q8, [x16, #0x240]
   352: 031c294e                 and v3.16b, v0.16b, v9.16b
   356: 04040c6f                 ushr v4.16b, v0.16b, #4
   360: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   364: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   368: a51ca54e                 mov v5.16b, v5.16b
   372: a51c266e                 eor v5.16b, v5.16b, v6.16b
   376: a51c216e                 eor v5.16b, v5.16b, v1.16b
   380: 07a2c03d                 ldr q7, [x16, #0x280]
   384: 08b2c03d                 ldr q8, [x16, #0x2c0]
   388: a31c294e                 and v3.16b, v5.16b, v9.16b
   392: a4040c6f                 ushr v4.16b, v5.16b, #4
   396: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   400: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   404: a11ca54e                 mov v1.16b, v5.16b
   408: 211c266e                 eor v1.16b, v1.16b, v6.16b
   412: 07a2c03d                 ldr q7, [x16, #0x280]
   416: 08b2c03d                 ldr q8, [x16, #0x2c0]
   420: 431c294e                 and v3.16b, v2.16b, v9.16b
   424: 44040c6f                 ushr v4.16b, v2.16b, #4
   428: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   432: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   436: a01ca54e                 mov v0.16b, v5.16b
   440: 001c266e                 eor v0.16b, v0.16b, v6.16b
   444: 2a0140f9                 ldr x10, [x9]
   448: 4a010b8b                 add x10, x10, x11
   452: 4001803d                 str q0, [x10]
   456: 2a0540f9                 ldr x10, [x9, #8]
   460: 4a010b8b                 add x10, x10, x11
   464: 4101803d                 str q1, [x10]
   468: 6b410091                 add x11, x11, #0x10
   472: ad0500f1                 subs x13, x13, #1
   476: 41f2ff54                 b.ne 0x24
   480: 7f0102eb                 cmp x11, x2
   484: 22110054                 b.hs 0x408
   488: 0f020c91                 add x15, x16, #0x300
   492: 11021091                 add x17, x16, #0x400
   496: 2a0140f9                 ldr x10, [x9]
   500: ed030baa                 mov x13, x11
   504: 0c008052                 mov w12, #0
   508: 0e0140f9                 ldr x14, [x8]
   512: ce696d38                 ldrb w14, [x14, x13]
   516: 0e010034                 cbz w14, 0x224
   520: ee696e38                 ldrb w14, [x15, x14]
   524: ce050311                 add w14, w14, #0xc1
   528: dffd0371                 cmp w14, #0xff
   532: 43000054                 b.lo 0x21c
   536: cefd0351                 sub w14, w14, #0xff
   540: 2e6a6e38                 ldrb w14, [x17, x14]
   544: 8c010e4a                 eor w12, w12, w14
   548: 0e0540f9                 ldr x14, [x8, #8]
   552: ce696d38                 ldrb w14, [x14, x13]
   556: 0e010034                 cbz w14, 0x24c
   560: ee696e38                 ldrb w14, [x15, x14]
   564: ce0d0011                 add w14, w14, #3
   568: dffd0371                 cmp w14, #0xff
   572: 43000054                 b.lo 0x244
   576: cefd0351                 sub w14, w14, #0xff
   580: 2e6a6e38                 ldrb w14, [x17, x14]
   584: 8c010e4a                 eor w12, w12, w14
   588: 0e0940f9                 ldr x14, [x8, #0x10]
   592: ce696d38                 ldrb w14, [x14, x13]
   596: 0e010034                 cbz w14, 0x274
   600: ee696e38                 ldrb w14, [x15, x14]
   604: ce190311                 add w14, w14, #0xc6
   608: dffd0371                 cmp w14, #0xff
   612: 43000054                 b.lo 0x26c
   616: cefd0351                 sub w14, w14, #0xff
   620: 2e6a6e38                 ldrb w14, [x17, x14]
   624: 8c010e4a                 eor w12, w12, w14
   628: 0e0d40f9                 ldr x14, [x8, #0x18]
   632: ce696d38                 ldrb w14, [x14, x13]
   636: 0e010034                 cbz w14, 0x29c
   640: ee696e38                 ldrb w14, [x15, x14]
   644: ce050011                 add w14, w14, #1
   648: dffd0371                 cmp w14, #0xff
   652: 43000054                 b.lo 0x294
   656: cefd0351                 sub w14, w14, #0xff
   660: 2e6a6e38                 ldrb w14, [x17, x14]
   664: 8c010e4a                 eor w12, w12, w14
   668: 0e1140f9                 ldr x14, [x8, #0x20]
   672: ce696d38                 ldrb w14, [x14, x13]
   676: 0e010034                 cbz w14, 0x2c4
   680: ee696e38                 ldrb w14, [x15, x14]
   684: cef90311                 add w14, w14, #0xfe
   688: dffd0371                 cmp w14, #0xff
   692: 43000054                 b.lo 0x2bc
   696: cefd0351                 sub w14, w14, #0xff
   700: 2e6a6e38                 ldrb w14, [x17, x14]
   704: 8c010e4a                 eor w12, w12, w14
   708: 0e1540f9                 ldr x14, [x8, #0x28]
   712: ce696d38                 ldrb w14, [x14, x13]
   716: 0e010034                 cbz w14, 0x2ec
   720: ee696e38                 ldrb w14, [x15, x14]
   724: cec50011                 add w14, w14, #0x31
   728: dffd0371                 cmp w14, #0xff
   732: 43000054                 b.lo 0x2e4
   736: cefd0351                 sub w14, w14, #0xff
   740: 2e6a6e38                 ldrb w14, [x17, x14]
   744: 8c010e4a                 eor w12, w12, w14
   748: 4c692d38                 strb w12, [x10, x13]
   752: ad050091                 add x13, x13, #1
   756: bf0102eb                 cmp x13, x2
   760: 03f8ff54                 b.lo 0x1f8
   764: 2a0540f9                 ldr x10, [x9, #8]
   768: ed030baa                 mov x13, x11
   772: 0c008052                 mov w12, #0
   776: 0e0140f9                 ldr x14, [x8]
   780: ce696d38                 ldrb w14, [x14, x13]
   784: 0e010034                 cbz w14, 0x330
   788: ee696e38                 ldrb w14, [x15, x14]
   792: ce710011                 add w14, w14, #0x1c
   796: dffd0371                 cmp w14, #0xff
   800: 43000054                 b.lo 0x328
   804: cefd0351                 sub w14, w14, #0xff
   808: 2e6a6e38                 ldrb w14, [x17, x14]
   812: 8c010e4a                 eor w12, w12, w14
   816: 0e0540f9                 ldr x14, [x8, #8]
   820: ce696d38                 ldrb w14, [x14, x13]
   824: 0e010034                 cbz w14, 0x358
   828: ee696e38                 ldrb w14, [x15, x14]
   832: ce7d0311                 add w14, w14, #0xdf
   836: dffd0371                 cmp w14, #0xff
   840: 43000054                 b.lo 0x350
   844: cefd0351                 sub w14, w14, #0xff
   848: 2e6a6e38                 ldrb w14, [x17, x14]
   852: 8c010e4a                 eor w12, w12, w14
   856: 0e0940f9                 ldr x14, [x8, #0x10]
   860: ce696d38                 ldrb w14, [x14, x13]
   864: 0e010034                 cbz w14, 0x380
   868: ee696e38                 ldrb w14, [x15, x14]
   872: ce690011                 add w14, w14, #0x1a
   876: dffd0371                 cmp w14, #0xff
   880: 43000054                 b.lo 0x378
   884: cefd0351                 sub w14, w14, #0xff
   888: 2e6a6e38                 ldrb w14, [x17, x14]
   892: 8c010e4a                 eor w12, w12, w14
   896: 0e0d40f9                 ldr x14, [x8, #0x18]
   900: ce696d38                 ldrb w14, [x14, x13]
   904: 0e010034                 cbz w14, 0x3a8
   908: ee696e38                 ldrb w14, [x15, x14]
   912: ce650011                 add w14, w14, #0x19
   916: dffd0371                 cmp w14, #0xff
   920: 43000054                 b.lo 0x3a0
   924: cefd0351                 sub w14, w14, #0xff
   928: 2e6a6e38                 ldrb w14, [x17, x14]
   932: 8c010e4a                 eor w12, w12, w14
   936: 0e1140f9                 ldr x14, [x8, #0x20]
   940: ce696d38                 ldrb w14, [x14, x13]
   944: 0e010034                 cbz w14, 0x3d0
   948: ee696e38                 ldrb w14, [x15, x14]
   952: ce610011                 add w14, w14, #0x18
   956: dffd0371                 cmp w14, #0xff
   960: 43000054                 b.lo 0x3c8
   964: cefd0351                 sub w14, w14, #0xff
   968: 2e6a6e38                 ldrb w14, [x17, x14]
   972: 8c010e4a                 eor w12, w12, w14
   976: 0e1540f9                 ldr x14, [x8, #0x28]
   980: ce696d38                 ldrb w14, [x14, x13]
   984: 0e010034                 cbz w14, 0x3f8
   988: ee696e38                 ldrb w14, [x15, x14]
   992: cec50011                 add w14, w14, #0x31
   996: dffd0371                 cmp w14, #0xff
  1000: 43000054                 b.lo 0x3f0
  1004: cefd0351                 sub w14, w14, #0xff
  1008: 2e6a6e38                 ldrb w14, [x17, x14]
  1012: 8c010e4a                 eor w12, w12, w14
  1016: 4c692d38                 strb w12, [x10, x13]
  1020: ad050091                 add x13, x13, #1
  1024: bf0102eb                 cmp x13, x2
  1028: 03f8ff54                 b.lo 0x304
  1032: e827c1ac                 ldp q8, q9, [sp], #0x20
  1036: c0035fd6                 ret 
  1040: 00000000                 udf #0
  1044: 00000000                 udf #0
  1048: 00000000                 udf #0
  1052: 00000000                 udf #0
  1056: 00000000                 udf #0
  1060: 00000000                 udf #0
  1064: 00000000                 udf #0
  1068: 00000000                 udf #0
  1072: 00000000                 udf #0
  1076: 00000000                 udf #0
  1080: 00000000                 udf #0
  1084: 00000000                 udf #0
