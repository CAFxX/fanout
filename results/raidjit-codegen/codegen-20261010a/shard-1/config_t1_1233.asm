; config_t1_1233 target=neon twin=pipe topo=rs(6,3)/enc op=encode U=1
; code_size=1216 twins=["pipe"] uops_per_byte=6.1875
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: e80300aa                 mov x8, x0
    12: e90301aa                 mov x9, x1
    16: 90250010                 adr x16, 0x4c0
    20: eb031faa                 mov x11, xzr
    24: 0e0280d2                 mov x14, #0x10
    28: 4d08ce9a                 udiv x13, x2, x14
    32: 2d1000b4                 cbz x13, 0x224
    36: ebe5004f                 movi v11.16b, #0xf
    40: 0a0140f9                 ldr x10, [x8]
    44: 4a010b8b                 add x10, x10, x11
    48: 4401c03d                 ldr q4, [x10]
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 801ca44e                 mov v0.16b, v4.16b
    68: 0902c03d                 ldr q9, [x16]
    72: 0a12c03d                 ldr q10, [x16, #0x40]
    76: 851c2b4e                 and v5.16b, v4.16b, v11.16b
    80: 86040c6f                 ushr v6.16b, v4.16b, #4
    84: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
    88: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
    92: e11ca74e                 mov v1.16b, v7.16b
    96: 211c286e                 eor v1.16b, v1.16b, v8.16b
   100: 0922c03d                 ldr q9, [x16, #0x80]
   104: 0a32c03d                 ldr q10, [x16, #0xc0]
   108: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   112: 86040c6f                 ushr v6.16b, v4.16b, #4
   116: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   120: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   124: e21ca74e                 mov v2.16b, v7.16b
   128: 421c286e                 eor v2.16b, v2.16b, v8.16b
   132: 0a0540f9                 ldr x10, [x8, #8]
   136: 4a010b8b                 add x10, x10, x11
   140: 4301c03d                 ldr q3, [x10]
   144: 001c236e                 eor v0.16b, v0.16b, v3.16b
   148: 0922c03d                 ldr q9, [x16, #0x80]
   152: 0a32c03d                 ldr q10, [x16, #0xc0]
   156: 651c2b4e                 and v5.16b, v3.16b, v11.16b
   160: 66040c6f                 ushr v6.16b, v3.16b, #4
   164: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   168: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   172: 211c276e                 eor v1.16b, v1.16b, v7.16b
   176: 211c286e                 eor v1.16b, v1.16b, v8.16b
   180: 0942c03d                 ldr q9, [x16, #0x100]
   184: 0a52c03d                 ldr q10, [x16, #0x140]
   188: 651c2b4e                 and v5.16b, v3.16b, v11.16b
   192: 66040c6f                 ushr v6.16b, v3.16b, #4
   196: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   200: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   204: 421c276e                 eor v2.16b, v2.16b, v7.16b
   208: 421c286e                 eor v2.16b, v2.16b, v8.16b
   212: 0a0940f9                 ldr x10, [x8, #0x10]
   216: 4a010b8b                 add x10, x10, x11
   220: 4301c03d                 ldr q3, [x10]
   224: 001c236e                 eor v0.16b, v0.16b, v3.16b
   228: 0942c03d                 ldr q9, [x16, #0x100]
   232: 0a52c03d                 ldr q10, [x16, #0x140]
   236: 651c2b4e                 and v5.16b, v3.16b, v11.16b
   240: 66040c6f                 ushr v6.16b, v3.16b, #4
   244: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   248: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   252: 211c276e                 eor v1.16b, v1.16b, v7.16b
   256: 211c286e                 eor v1.16b, v1.16b, v8.16b
   260: 0962c03d                 ldr q9, [x16, #0x180]
   264: 0a72c03d                 ldr q10, [x16, #0x1c0]
   268: 651c2b4e                 and v5.16b, v3.16b, v11.16b
   272: 66040c6f                 ushr v6.16b, v3.16b, #4
   276: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   280: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   284: 421c276e                 eor v2.16b, v2.16b, v7.16b
   288: 421c286e                 eor v2.16b, v2.16b, v8.16b
   292: 0a0d40f9                 ldr x10, [x8, #0x18]
   296: 4a010b8b                 add x10, x10, x11
   300: 4301c03d                 ldr q3, [x10]
   304: 001c236e                 eor v0.16b, v0.16b, v3.16b
   308: 0982c03d                 ldr q9, [x16, #0x200]
   312: 0a92c03d                 ldr q10, [x16, #0x240]
   316: 651c2b4e                 and v5.16b, v3.16b, v11.16b
   320: 66040c6f                 ushr v6.16b, v3.16b, #4
   324: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   328: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   332: 211c276e                 eor v1.16b, v1.16b, v7.16b
   336: 211c286e                 eor v1.16b, v1.16b, v8.16b
   340: 0902c03d                 ldr q9, [x16]
   344: 0a12c03d                 ldr q10, [x16, #0x40]
   348: 651c2b4e                 and v5.16b, v3.16b, v11.16b
   352: 66040c6f                 ushr v6.16b, v3.16b, #4
   356: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   360: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   364: 421c276e                 eor v2.16b, v2.16b, v7.16b
   368: 421c286e                 eor v2.16b, v2.16b, v8.16b
   372: 0a1140f9                 ldr x10, [x8, #0x20]
   376: 4a010b8b                 add x10, x10, x11
   380: 4301c03d                 ldr q3, [x10]
   384: 001c236e                 eor v0.16b, v0.16b, v3.16b
   388: 0962c03d                 ldr q9, [x16, #0x180]
   392: 0a72c03d                 ldr q10, [x16, #0x1c0]
   396: 651c2b4e                 and v5.16b, v3.16b, v11.16b
   400: 66040c6f                 ushr v6.16b, v3.16b, #4
   404: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   408: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   412: 211c276e                 eor v1.16b, v1.16b, v7.16b
   416: 211c286e                 eor v1.16b, v1.16b, v8.16b
   420: 0982c03d                 ldr q9, [x16, #0x200]
   424: 0a92c03d                 ldr q10, [x16, #0x240]
   428: 651c2b4e                 and v5.16b, v3.16b, v11.16b
   432: 66040c6f                 ushr v6.16b, v3.16b, #4
   436: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   440: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   444: 421c276e                 eor v2.16b, v2.16b, v7.16b
   448: 421c286e                 eor v2.16b, v2.16b, v8.16b
   452: 0a1540f9                 ldr x10, [x8, #0x28]
   456: 4a010b8b                 add x10, x10, x11
   460: 4301c03d                 ldr q3, [x10]
   464: 001c236e                 eor v0.16b, v0.16b, v3.16b
   468: 211c236e                 eor v1.16b, v1.16b, v3.16b
   472: 421c236e                 eor v2.16b, v2.16b, v3.16b
   476: bf0500f1                 cmp x13, #1
   480: ad000054                 b.le 0x1f4
   484: 0a0140f9                 ldr x10, [x8]
   488: 4a010b8b                 add x10, x10, x11
   492: 4a410091                 add x10, x10, #0x10
   496: 4401c03d                 ldr q4, [x10]
   500: 2a0140f9                 ldr x10, [x9]
   504: 4a010b8b                 add x10, x10, x11
   508: 4001803d                 str q0, [x10]
   512: 2a0540f9                 ldr x10, [x9, #8]
   516: 4a010b8b                 add x10, x10, x11
   520: 4101803d                 str q1, [x10]
   524: 2a0940f9                 ldr x10, [x9, #0x10]
   528: 4a010b8b                 add x10, x10, x11
   532: 4201803d                 str q2, [x10]
   536: 6b410091                 add x11, x11, #0x10
   540: ad0500f1                 subs x13, x13, #1
   544: a1f0ff54                 b.ne 0x34
   548: 7f0102eb                 cmp x11, x2
   552: 82120054                 b.hs 0x478
   556: 0f020a91                 add x15, x16, #0x280
   560: 11020e91                 add x17, x16, #0x380
   564: 2a0140f9                 ldr x10, [x9]
   568: ed030baa                 mov x13, x11
   572: 0c008052                 mov w12, #0
   576: 0e0140f9                 ldr x14, [x8]
   580: ce696d38                 ldrb w14, [x14, x13]
   584: 8c010e4a                 eor w12, w12, w14
   588: 0e0540f9                 ldr x14, [x8, #8]
   592: ce696d38                 ldrb w14, [x14, x13]
   596: 8c010e4a                 eor w12, w12, w14
   600: 0e0940f9                 ldr x14, [x8, #0x10]
   604: ce696d38                 ldrb w14, [x14, x13]
   608: 8c010e4a                 eor w12, w12, w14
   612: 0e0d40f9                 ldr x14, [x8, #0x18]
   616: ce696d38                 ldrb w14, [x14, x13]
   620: 8c010e4a                 eor w12, w12, w14
   624: 0e1140f9                 ldr x14, [x8, #0x20]
   628: ce696d38                 ldrb w14, [x14, x13]
   632: 8c010e4a                 eor w12, w12, w14
   636: 0e1540f9                 ldr x14, [x8, #0x28]
   640: ce696d38                 ldrb w14, [x14, x13]
   644: 8c010e4a                 eor w12, w12, w14
   648: 4c692d38                 strb w12, [x10, x13]
   652: ad050091                 add x13, x13, #1
   656: bf0102eb                 cmp x13, x2
   660: 43fdff54                 b.lo 0x23c
   664: 2a0540f9                 ldr x10, [x9, #8]
   668: ed030baa                 mov x13, x11
   672: 0c008052                 mov w12, #0
   676: 0e0140f9                 ldr x14, [x8]
   680: ce696d38                 ldrb w14, [x14, x13]
   684: 0e010034                 cbz w14, 0x2cc
   688: ee696e38                 ldrb w14, [x15, x14]
   692: ce990311                 add w14, w14, #0xe6
   696: dffd0371                 cmp w14, #0xff
   700: 43000054                 b.lo 0x2c4
   704: cefd0351                 sub w14, w14, #0xff
   708: 2e6a6e38                 ldrb w14, [x17, x14]
   712: 8c010e4a                 eor w12, w12, w14
   716: 0e0540f9                 ldr x14, [x8, #8]
   720: ce696d38                 ldrb w14, [x14, x13]
   724: 0e010034                 cbz w14, 0x2f4
   728: ee696e38                 ldrb w14, [x15, x14]
   732: ce350311                 add w14, w14, #0xcd
   736: dffd0371                 cmp w14, #0xff
   740: 43000054                 b.lo 0x2ec
   744: cefd0351                 sub w14, w14, #0xff
   748: 2e6a6e38                 ldrb w14, [x17, x14]
   752: 8c010e4a                 eor w12, w12, w14
   756: 0e0940f9                 ldr x14, [x8, #0x10]
   760: ce696d38                 ldrb w14, [x14, x13]
   764: 0e010034                 cbz w14, 0x31c
   768: ee696e38                 ldrb w14, [x15, x14]
   772: ce2d0111                 add w14, w14, #0x4b
   776: dffd0371                 cmp w14, #0xff
   780: 43000054                 b.lo 0x314
   784: cefd0351                 sub w14, w14, #0xff
   788: 2e6a6e38                 ldrb w14, [x17, x14]
   792: 8c010e4a                 eor w12, w12, w14
   796: 0e0d40f9                 ldr x14, [x8, #0x18]
   800: ce696d38                 ldrb w14, [x14, x13]
   804: 0e010034                 cbz w14, 0x344
   808: ee696e38                 ldrb w14, [x15, x14]
   812: cec90011                 add w14, w14, #0x32
   816: dffd0371                 cmp w14, #0xff
   820: 43000054                 b.lo 0x33c
   824: cefd0351                 sub w14, w14, #0xff
   828: 2e6a6e38                 ldrb w14, [x17, x14]
   832: 8c010e4a                 eor w12, w12, w14
   836: 0e1140f9                 ldr x14, [x8, #0x20]
   840: ce696d38                 ldrb w14, [x14, x13]
   844: 0e010034                 cbz w14, 0x36c
   848: ee696e38                 ldrb w14, [x15, x14]
   852: ce650011                 add w14, w14, #0x19
   856: dffd0371                 cmp w14, #0xff
   860: 43000054                 b.lo 0x364
   864: cefd0351                 sub w14, w14, #0xff
   868: 2e6a6e38                 ldrb w14, [x17, x14]
   872: 8c010e4a                 eor w12, w12, w14
   876: 0e1540f9                 ldr x14, [x8, #0x28]
   880: ce696d38                 ldrb w14, [x14, x13]
   884: 8c010e4a                 eor w12, w12, w14
   888: 4c692d38                 strb w12, [x10, x13]
   892: ad050091                 add x13, x13, #1
   896: bf0102eb                 cmp x13, x2
   900: e3f8ff54                 b.lo 0x2a0
   904: 2a0940f9                 ldr x10, [x9, #0x10]
   908: ed030baa                 mov x13, x11
   912: 0c008052                 mov w12, #0
   916: 0e0140f9                 ldr x14, [x8]
   920: ce696d38                 ldrb w14, [x14, x13]
   924: 0e010034                 cbz w14, 0x3bc
   928: ee696e38                 ldrb w14, [x15, x14]
   932: ce350311                 add w14, w14, #0xcd
   936: dffd0371                 cmp w14, #0xff
   940: 43000054                 b.lo 0x3b4
   944: cefd0351                 sub w14, w14, #0xff
   948: 2e6a6e38                 ldrb w14, [x17, x14]
   952: 8c010e4a                 eor w12, w12, w14
   956: 0e0540f9                 ldr x14, [x8, #8]
   960: ce696d38                 ldrb w14, [x14, x13]
   964: 0e010034                 cbz w14, 0x3e4
   968: ee696e38                 ldrb w14, [x15, x14]
   972: ce2d0111                 add w14, w14, #0x4b
   976: dffd0371                 cmp w14, #0xff
   980: 43000054                 b.lo 0x3dc
   984: cefd0351                 sub w14, w14, #0xff
   988: 2e6a6e38                 ldrb w14, [x17, x14]
   992: 8c010e4a                 eor w12, w12, w14
   996: 0e0940f9                 ldr x14, [x8, #0x10]
  1000: ce696d38                 ldrb w14, [x14, x13]
  1004: 0e010034                 cbz w14, 0x40c
  1008: ee696e38                 ldrb w14, [x15, x14]
  1012: ce650011                 add w14, w14, #0x19
  1016: dffd0371                 cmp w14, #0xff
  1020: 43000054                 b.lo 0x404
  1024: cefd0351                 sub w14, w14, #0xff
  1028: 2e6a6e38                 ldrb w14, [x17, x14]
  1032: 8c010e4a                 eor w12, w12, w14
  1036: 0e0d40f9                 ldr x14, [x8, #0x18]
  1040: ce696d38                 ldrb w14, [x14, x13]
  1044: 0e010034                 cbz w14, 0x434
  1048: ee696e38                 ldrb w14, [x15, x14]
  1052: ce990311                 add w14, w14, #0xe6
  1056: dffd0371                 cmp w14, #0xff
  1060: 43000054                 b.lo 0x42c
  1064: cefd0351                 sub w14, w14, #0xff
  1068: 2e6a6e38                 ldrb w14, [x17, x14]
  1072: 8c010e4a                 eor w12, w12, w14
  1076: 0e1140f9                 ldr x14, [x8, #0x20]
  1080: ce696d38                 ldrb w14, [x14, x13]
  1084: 0e010034                 cbz w14, 0x45c
  1088: ee696e38                 ldrb w14, [x15, x14]
  1092: cec90011                 add w14, w14, #0x32
  1096: dffd0371                 cmp w14, #0xff
  1100: 43000054                 b.lo 0x454
  1104: cefd0351                 sub w14, w14, #0xff
  1108: 2e6a6e38                 ldrb w14, [x17, x14]
  1112: 8c010e4a                 eor w12, w12, w14
  1116: 0e1540f9                 ldr x14, [x8, #0x28]
  1120: ce696d38                 ldrb w14, [x14, x13]
  1124: 8c010e4a                 eor w12, w12, w14
  1128: 4c692d38                 strb w12, [x10, x13]
  1132: ad050091                 add x13, x13, #1
  1136: bf0102eb                 cmp x13, x2
  1140: e3f8ff54                 b.lo 0x390
  1144: ea2fc1ac                 ldp q10, q11, [sp], #0x20
  1148: e827c1ac                 ldp q8, q9, [sp], #0x20
  1152: c0035fd6                 ret 
  1156: 00000000                 udf #0
  1160: 00000000                 udf #0
  1164: 00000000                 udf #0
  1168: 00000000                 udf #0
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
