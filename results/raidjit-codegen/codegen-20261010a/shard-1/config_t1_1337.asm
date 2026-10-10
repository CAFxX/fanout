; config_t1_1337 target=neon twin=pipe topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1152 twins=["pipe"] uops_per_byte=6.625
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: e80300aa                 mov x8, x0
    12: e90301aa                 mov x9, x1
    16: 90230010                 adr x16, 0x480
    20: eb031faa                 mov x11, xzr
    24: 0e0280d2                 mov x14, #0x10
    28: 4d08ce9a                 udiv x13, x2, x14
    32: cd1000b4                 cbz x13, 0x238
    36: eae5004f                 movi v10.16b, #0xf
    40: 0a0140f9                 ldr x10, [x8]
    44: 4a010b8b                 add x10, x10, x11
    48: 4301c03d                 ldr q3, [x10]
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0802c03d                 ldr q8, [x16]
    68: 0912c03d                 ldr q9, [x16, #0x40]
    72: 641c2a4e                 and v4.16b, v3.16b, v10.16b
    76: 65040c6f                 ushr v5.16b, v3.16b, #4
    80: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
    84: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
    88: c01ca64e                 mov v0.16b, v6.16b
    92: 001c276e                 eor v0.16b, v0.16b, v7.16b
    96: 0822c03d                 ldr q8, [x16, #0x80]
   100: 0932c03d                 ldr q9, [x16, #0xc0]
   104: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   108: 65040c6f                 ushr v5.16b, v3.16b, #4
   112: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   116: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   120: c11ca64e                 mov v1.16b, v6.16b
   124: 211c276e                 eor v1.16b, v1.16b, v7.16b
   128: 0a0540f9                 ldr x10, [x8, #8]
   132: 4a010b8b                 add x10, x10, x11
   136: 4201c03d                 ldr q2, [x10]
   140: 0842c03d                 ldr q8, [x16, #0x100]
   144: 0952c03d                 ldr q9, [x16, #0x140]
   148: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   152: 45040c6f                 ushr v5.16b, v2.16b, #4
   156: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   160: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   164: 001c266e                 eor v0.16b, v0.16b, v6.16b
   168: 001c276e                 eor v0.16b, v0.16b, v7.16b
   172: 0862c03d                 ldr q8, [x16, #0x180]
   176: 0972c03d                 ldr q9, [x16, #0x1c0]
   180: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   184: 45040c6f                 ushr v5.16b, v2.16b, #4
   188: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   192: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   196: 211c266e                 eor v1.16b, v1.16b, v6.16b
   200: 211c276e                 eor v1.16b, v1.16b, v7.16b
   204: 0a0940f9                 ldr x10, [x8, #0x10]
   208: 4a010b8b                 add x10, x10, x11
   212: 4201c03d                 ldr q2, [x10]
   216: 0882c03d                 ldr q8, [x16, #0x200]
   220: 0992c03d                 ldr q9, [x16, #0x240]
   224: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   228: 45040c6f                 ushr v5.16b, v2.16b, #4
   232: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   236: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   240: 001c266e                 eor v0.16b, v0.16b, v6.16b
   244: 001c276e                 eor v0.16b, v0.16b, v7.16b
   248: 08a2c03d                 ldr q8, [x16, #0x280]
   252: 09b2c03d                 ldr q9, [x16, #0x2c0]
   256: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   260: 45040c6f                 ushr v5.16b, v2.16b, #4
   264: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   268: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   272: 211c266e                 eor v1.16b, v1.16b, v6.16b
   276: 211c276e                 eor v1.16b, v1.16b, v7.16b
   280: 0a0d40f9                 ldr x10, [x8, #0x18]
   284: 4a010b8b                 add x10, x10, x11
   288: 4201c03d                 ldr q2, [x10]
   292: 08c2c03d                 ldr q8, [x16, #0x300]
   296: 09d2c03d                 ldr q9, [x16, #0x340]
   300: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   304: 45040c6f                 ushr v5.16b, v2.16b, #4
   308: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   312: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   316: 001c266e                 eor v0.16b, v0.16b, v6.16b
   320: 001c276e                 eor v0.16b, v0.16b, v7.16b
   324: 08e2c03d                 ldr q8, [x16, #0x380]
   328: 09f2c03d                 ldr q9, [x16, #0x3c0]
   332: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   336: 45040c6f                 ushr v5.16b, v2.16b, #4
   340: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   344: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   348: 211c266e                 eor v1.16b, v1.16b, v6.16b
   352: 211c276e                 eor v1.16b, v1.16b, v7.16b
   356: 0a1140f9                 ldr x10, [x8, #0x20]
   360: 4a010b8b                 add x10, x10, x11
   364: 4201c03d                 ldr q2, [x10]
   368: 0802c13d                 ldr q8, [x16, #0x400]
   372: 0912c13d                 ldr q9, [x16, #0x440]
   376: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   380: 45040c6f                 ushr v5.16b, v2.16b, #4
   384: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   388: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   392: 001c266e                 eor v0.16b, v0.16b, v6.16b
   396: 001c276e                 eor v0.16b, v0.16b, v7.16b
   400: 0822c13d                 ldr q8, [x16, #0x480]
   404: 0932c13d                 ldr q9, [x16, #0x4c0]
   408: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   412: 45040c6f                 ushr v5.16b, v2.16b, #4
   416: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   420: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   424: 211c266e                 eor v1.16b, v1.16b, v6.16b
   428: 211c276e                 eor v1.16b, v1.16b, v7.16b
   432: 0a1540f9                 ldr x10, [x8, #0x28]
   436: 4a010b8b                 add x10, x10, x11
   440: 4201c03d                 ldr q2, [x10]
   444: 0842c13d                 ldr q8, [x16, #0x500]
   448: 0952c13d                 ldr q9, [x16, #0x540]
   452: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   456: 45040c6f                 ushr v5.16b, v2.16b, #4
   460: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   464: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   468: 001c266e                 eor v0.16b, v0.16b, v6.16b
   472: 001c276e                 eor v0.16b, v0.16b, v7.16b
   476: 0842c13d                 ldr q8, [x16, #0x500]
   480: 0952c13d                 ldr q9, [x16, #0x540]
   484: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   488: 45040c6f                 ushr v5.16b, v2.16b, #4
   492: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   496: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   500: 211c266e                 eor v1.16b, v1.16b, v6.16b
   504: 211c276e                 eor v1.16b, v1.16b, v7.16b
   508: bf0500f1                 cmp x13, #1
   512: ad000054                 b.le 0x214
   516: 0a0140f9                 ldr x10, [x8]
   520: 4a010b8b                 add x10, x10, x11
   524: 4a410091                 add x10, x10, #0x10
   528: 4301c03d                 ldr q3, [x10]
   532: 2a0140f9                 ldr x10, [x9]
   536: 4a010b8b                 add x10, x10, x11
   540: 4001803d                 str q0, [x10]
   544: 2a0540f9                 ldr x10, [x9, #8]
   548: 4a010b8b                 add x10, x10, x11
   552: 4101803d                 str q1, [x10]
   556: 6b410091                 add x11, x11, #0x10
   560: ad0500f1                 subs x13, x13, #1
   564: 01f0ff54                 b.ne 0x34
   568: 7f0102eb                 cmp x11, x2
   572: 22110054                 b.hs 0x460
   576: 0f021691                 add x15, x16, #0x580
   580: 11021a91                 add x17, x16, #0x680
   584: 2a0140f9                 ldr x10, [x9]
   588: ed030baa                 mov x13, x11
   592: 0c008052                 mov w12, #0
   596: 0e0140f9                 ldr x14, [x8]
   600: ce696d38                 ldrb w14, [x14, x13]
   604: 0e010034                 cbz w14, 0x27c
   608: ee696e38                 ldrb w14, [x15, x14]
   612: ce050311                 add w14, w14, #0xc1
   616: dffd0371                 cmp w14, #0xff
   620: 43000054                 b.lo 0x274
   624: cefd0351                 sub w14, w14, #0xff
   628: 2e6a6e38                 ldrb w14, [x17, x14]
   632: 8c010e4a                 eor w12, w12, w14
   636: 0e0540f9                 ldr x14, [x8, #8]
   640: ce696d38                 ldrb w14, [x14, x13]
   644: 0e010034                 cbz w14, 0x2a4
   648: ee696e38                 ldrb w14, [x15, x14]
   652: ce0d0011                 add w14, w14, #3
   656: dffd0371                 cmp w14, #0xff
   660: 43000054                 b.lo 0x29c
   664: cefd0351                 sub w14, w14, #0xff
   668: 2e6a6e38                 ldrb w14, [x17, x14]
   672: 8c010e4a                 eor w12, w12, w14
   676: 0e0940f9                 ldr x14, [x8, #0x10]
   680: ce696d38                 ldrb w14, [x14, x13]
   684: 0e010034                 cbz w14, 0x2cc
   688: ee696e38                 ldrb w14, [x15, x14]
   692: ce190311                 add w14, w14, #0xc6
   696: dffd0371                 cmp w14, #0xff
   700: 43000054                 b.lo 0x2c4
   704: cefd0351                 sub w14, w14, #0xff
   708: 2e6a6e38                 ldrb w14, [x17, x14]
   712: 8c010e4a                 eor w12, w12, w14
   716: 0e0d40f9                 ldr x14, [x8, #0x18]
   720: ce696d38                 ldrb w14, [x14, x13]
   724: 0e010034                 cbz w14, 0x2f4
   728: ee696e38                 ldrb w14, [x15, x14]
   732: ce050011                 add w14, w14, #1
   736: dffd0371                 cmp w14, #0xff
   740: 43000054                 b.lo 0x2ec
   744: cefd0351                 sub w14, w14, #0xff
   748: 2e6a6e38                 ldrb w14, [x17, x14]
   752: 8c010e4a                 eor w12, w12, w14
   756: 0e1140f9                 ldr x14, [x8, #0x20]
   760: ce696d38                 ldrb w14, [x14, x13]
   764: 0e010034                 cbz w14, 0x31c
   768: ee696e38                 ldrb w14, [x15, x14]
   772: cef90311                 add w14, w14, #0xfe
   776: dffd0371                 cmp w14, #0xff
   780: 43000054                 b.lo 0x314
   784: cefd0351                 sub w14, w14, #0xff
   788: 2e6a6e38                 ldrb w14, [x17, x14]
   792: 8c010e4a                 eor w12, w12, w14
   796: 0e1540f9                 ldr x14, [x8, #0x28]
   800: ce696d38                 ldrb w14, [x14, x13]
   804: 0e010034                 cbz w14, 0x344
   808: ee696e38                 ldrb w14, [x15, x14]
   812: cec50011                 add w14, w14, #0x31
   816: dffd0371                 cmp w14, #0xff
   820: 43000054                 b.lo 0x33c
   824: cefd0351                 sub w14, w14, #0xff
   828: 2e6a6e38                 ldrb w14, [x17, x14]
   832: 8c010e4a                 eor w12, w12, w14
   836: 4c692d38                 strb w12, [x10, x13]
   840: ad050091                 add x13, x13, #1
   844: bf0102eb                 cmp x13, x2
   848: 03f8ff54                 b.lo 0x250
   852: 2a0540f9                 ldr x10, [x9, #8]
   856: ed030baa                 mov x13, x11
   860: 0c008052                 mov w12, #0
   864: 0e0140f9                 ldr x14, [x8]
   868: ce696d38                 ldrb w14, [x14, x13]
   872: 0e010034                 cbz w14, 0x388
   876: ee696e38                 ldrb w14, [x15, x14]
   880: ce710011                 add w14, w14, #0x1c
   884: dffd0371                 cmp w14, #0xff
   888: 43000054                 b.lo 0x380
   892: cefd0351                 sub w14, w14, #0xff
   896: 2e6a6e38                 ldrb w14, [x17, x14]
   900: 8c010e4a                 eor w12, w12, w14
   904: 0e0540f9                 ldr x14, [x8, #8]
   908: ce696d38                 ldrb w14, [x14, x13]
   912: 0e010034                 cbz w14, 0x3b0
   916: ee696e38                 ldrb w14, [x15, x14]
   920: ce7d0311                 add w14, w14, #0xdf
   924: dffd0371                 cmp w14, #0xff
   928: 43000054                 b.lo 0x3a8
   932: cefd0351                 sub w14, w14, #0xff
   936: 2e6a6e38                 ldrb w14, [x17, x14]
   940: 8c010e4a                 eor w12, w12, w14
   944: 0e0940f9                 ldr x14, [x8, #0x10]
   948: ce696d38                 ldrb w14, [x14, x13]
   952: 0e010034                 cbz w14, 0x3d8
   956: ee696e38                 ldrb w14, [x15, x14]
   960: ce690011                 add w14, w14, #0x1a
   964: dffd0371                 cmp w14, #0xff
   968: 43000054                 b.lo 0x3d0
   972: cefd0351                 sub w14, w14, #0xff
   976: 2e6a6e38                 ldrb w14, [x17, x14]
   980: 8c010e4a                 eor w12, w12, w14
   984: 0e0d40f9                 ldr x14, [x8, #0x18]
   988: ce696d38                 ldrb w14, [x14, x13]
   992: 0e010034                 cbz w14, 0x400
   996: ee696e38                 ldrb w14, [x15, x14]
  1000: ce650011                 add w14, w14, #0x19
  1004: dffd0371                 cmp w14, #0xff
  1008: 43000054                 b.lo 0x3f8
  1012: cefd0351                 sub w14, w14, #0xff
  1016: 2e6a6e38                 ldrb w14, [x17, x14]
  1020: 8c010e4a                 eor w12, w12, w14
  1024: 0e1140f9                 ldr x14, [x8, #0x20]
  1028: ce696d38                 ldrb w14, [x14, x13]
  1032: 0e010034                 cbz w14, 0x428
  1036: ee696e38                 ldrb w14, [x15, x14]
  1040: ce610011                 add w14, w14, #0x18
  1044: dffd0371                 cmp w14, #0xff
  1048: 43000054                 b.lo 0x420
  1052: cefd0351                 sub w14, w14, #0xff
  1056: 2e6a6e38                 ldrb w14, [x17, x14]
  1060: 8c010e4a                 eor w12, w12, w14
  1064: 0e1540f9                 ldr x14, [x8, #0x28]
  1068: ce696d38                 ldrb w14, [x14, x13]
  1072: 0e010034                 cbz w14, 0x450
  1076: ee696e38                 ldrb w14, [x15, x14]
  1080: cec50011                 add w14, w14, #0x31
  1084: dffd0371                 cmp w14, #0xff
  1088: 43000054                 b.lo 0x448
  1092: cefd0351                 sub w14, w14, #0xff
  1096: 2e6a6e38                 ldrb w14, [x17, x14]
  1100: 8c010e4a                 eor w12, w12, w14
  1104: 4c692d38                 strb w12, [x10, x13]
  1108: ad050091                 add x13, x13, #1
  1112: bf0102eb                 cmp x13, x2
  1116: 03f8ff54                 b.lo 0x35c
  1120: ea2fc1ac                 ldp q10, q11, [sp], #0x20
  1124: e827c1ac                 ldp q8, q9, [sp], #0x20
  1128: c0035fd6                 ret 
  1132: 00000000                 udf #0
  1136: 00000000                 udf #0
  1140: 00000000                 udf #0
  1144: 00000000                 udf #0
  1148: 00000000                 udf #0
