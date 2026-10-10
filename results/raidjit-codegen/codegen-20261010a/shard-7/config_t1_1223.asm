; config_t1_1223 target=neon twin=nt topo=rs(6,3)/enc op=encode U=1
; code_size=1216 twins=["nt"] uops_per_byte=6.875
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: e80300aa                 mov x8, x0
    12: e90301aa                 mov x9, x1
    16: 90250010                 adr x16, 0x4c0
    20: eb031faa                 mov x11, xzr
    24: 0e0280d2                 mov x14, #0x10
    28: 4d08ce9a                 udiv x13, x2, x14
    32: 4d1100b4                 cbz x13, 0x248
    36: eae5004f                 movi v10.16b, #0xf
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 4301c03d                 ldr q3, [x10]
    76: 601ca34e                 mov v0.16b, v3.16b
    80: 0802c03d                 ldr q8, [x16]
    84: 0912c03d                 ldr q9, [x16, #0x40]
    88: 4301c03d                 ldr q3, [x10]
    92: 641c2a4e                 and v4.16b, v3.16b, v10.16b
    96: 65040c6f                 ushr v5.16b, v3.16b, #4
   100: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   104: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   108: c11ca64e                 mov v1.16b, v6.16b
   112: 211c276e                 eor v1.16b, v1.16b, v7.16b
   116: 0822c03d                 ldr q8, [x16, #0x80]
   120: 0932c03d                 ldr q9, [x16, #0xc0]
   124: 4301c03d                 ldr q3, [x10]
   128: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   132: 65040c6f                 ushr v5.16b, v3.16b, #4
   136: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   140: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   144: c21ca64e                 mov v2.16b, v6.16b
   148: 421c276e                 eor v2.16b, v2.16b, v7.16b
   152: 0a0540f9                 ldr x10, [x8, #8]
   156: 4a010b8b                 add x10, x10, x11
   160: 4301c03d                 ldr q3, [x10]
   164: 001c236e                 eor v0.16b, v0.16b, v3.16b
   168: 0822c03d                 ldr q8, [x16, #0x80]
   172: 0932c03d                 ldr q9, [x16, #0xc0]
   176: 4301c03d                 ldr q3, [x10]
   180: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   184: 65040c6f                 ushr v5.16b, v3.16b, #4
   188: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   192: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   196: 211c266e                 eor v1.16b, v1.16b, v6.16b
   200: 211c276e                 eor v1.16b, v1.16b, v7.16b
   204: 0842c03d                 ldr q8, [x16, #0x100]
   208: 0952c03d                 ldr q9, [x16, #0x140]
   212: 4301c03d                 ldr q3, [x10]
   216: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   220: 65040c6f                 ushr v5.16b, v3.16b, #4
   224: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   228: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   232: 421c266e                 eor v2.16b, v2.16b, v6.16b
   236: 421c276e                 eor v2.16b, v2.16b, v7.16b
   240: 0a0940f9                 ldr x10, [x8, #0x10]
   244: 4a010b8b                 add x10, x10, x11
   248: 4301c03d                 ldr q3, [x10]
   252: 001c236e                 eor v0.16b, v0.16b, v3.16b
   256: 0842c03d                 ldr q8, [x16, #0x100]
   260: 0952c03d                 ldr q9, [x16, #0x140]
   264: 4301c03d                 ldr q3, [x10]
   268: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   272: 65040c6f                 ushr v5.16b, v3.16b, #4
   276: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   280: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   284: 211c266e                 eor v1.16b, v1.16b, v6.16b
   288: 211c276e                 eor v1.16b, v1.16b, v7.16b
   292: 0862c03d                 ldr q8, [x16, #0x180]
   296: 0972c03d                 ldr q9, [x16, #0x1c0]
   300: 4301c03d                 ldr q3, [x10]
   304: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   308: 65040c6f                 ushr v5.16b, v3.16b, #4
   312: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   316: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   320: 421c266e                 eor v2.16b, v2.16b, v6.16b
   324: 421c276e                 eor v2.16b, v2.16b, v7.16b
   328: 0a0d40f9                 ldr x10, [x8, #0x18]
   332: 4a010b8b                 add x10, x10, x11
   336: 4301c03d                 ldr q3, [x10]
   340: 001c236e                 eor v0.16b, v0.16b, v3.16b
   344: 0882c03d                 ldr q8, [x16, #0x200]
   348: 0992c03d                 ldr q9, [x16, #0x240]
   352: 4301c03d                 ldr q3, [x10]
   356: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   360: 65040c6f                 ushr v5.16b, v3.16b, #4
   364: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   368: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   372: 211c266e                 eor v1.16b, v1.16b, v6.16b
   376: 211c276e                 eor v1.16b, v1.16b, v7.16b
   380: 0802c03d                 ldr q8, [x16]
   384: 0912c03d                 ldr q9, [x16, #0x40]
   388: 4301c03d                 ldr q3, [x10]
   392: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   396: 65040c6f                 ushr v5.16b, v3.16b, #4
   400: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   404: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   408: 421c266e                 eor v2.16b, v2.16b, v6.16b
   412: 421c276e                 eor v2.16b, v2.16b, v7.16b
   416: 0a1140f9                 ldr x10, [x8, #0x20]
   420: 4a010b8b                 add x10, x10, x11
   424: 4301c03d                 ldr q3, [x10]
   428: 001c236e                 eor v0.16b, v0.16b, v3.16b
   432: 0862c03d                 ldr q8, [x16, #0x180]
   436: 0972c03d                 ldr q9, [x16, #0x1c0]
   440: 4301c03d                 ldr q3, [x10]
   444: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   448: 65040c6f                 ushr v5.16b, v3.16b, #4
   452: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   456: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   460: 211c266e                 eor v1.16b, v1.16b, v6.16b
   464: 211c276e                 eor v1.16b, v1.16b, v7.16b
   468: 0882c03d                 ldr q8, [x16, #0x200]
   472: 0992c03d                 ldr q9, [x16, #0x240]
   476: 4301c03d                 ldr q3, [x10]
   480: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   484: 65040c6f                 ushr v5.16b, v3.16b, #4
   488: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   492: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   496: 421c266e                 eor v2.16b, v2.16b, v6.16b
   500: 421c276e                 eor v2.16b, v2.16b, v7.16b
   504: 0a1540f9                 ldr x10, [x8, #0x28]
   508: 4a010b8b                 add x10, x10, x11
   512: 4301c03d                 ldr q3, [x10]
   516: 001c236e                 eor v0.16b, v0.16b, v3.16b
   520: 4301c03d                 ldr q3, [x10]
   524: 211c236e                 eor v1.16b, v1.16b, v3.16b
   528: 4301c03d                 ldr q3, [x10]
   532: 421c236e                 eor v2.16b, v2.16b, v3.16b
   536: 2a0140f9                 ldr x10, [x9]
   540: 4a010b8b                 add x10, x10, x11
   544: 4001803d                 str q0, [x10]
   548: 2a0540f9                 ldr x10, [x9, #8]
   552: 4a010b8b                 add x10, x10, x11
   556: 4101803d                 str q1, [x10]
   560: 2a0940f9                 ldr x10, [x9, #0x10]
   564: 4a010b8b                 add x10, x10, x11
   568: 4201803d                 str q2, [x10]
   572: 6b410091                 add x11, x11, #0x10
   576: ad0500f1                 subs x13, x13, #1
   580: 21efff54                 b.ne 0x28
   584: 7f0102eb                 cmp x11, x2
   588: 82120054                 b.hs 0x49c
   592: 0f020a91                 add x15, x16, #0x280
   596: 11020e91                 add x17, x16, #0x380
   600: 2a0140f9                 ldr x10, [x9]
   604: ed030baa                 mov x13, x11
   608: 0c008052                 mov w12, #0
   612: 0e0140f9                 ldr x14, [x8]
   616: ce696d38                 ldrb w14, [x14, x13]
   620: 8c010e4a                 eor w12, w12, w14
   624: 0e0540f9                 ldr x14, [x8, #8]
   628: ce696d38                 ldrb w14, [x14, x13]
   632: 8c010e4a                 eor w12, w12, w14
   636: 0e0940f9                 ldr x14, [x8, #0x10]
   640: ce696d38                 ldrb w14, [x14, x13]
   644: 8c010e4a                 eor w12, w12, w14
   648: 0e0d40f9                 ldr x14, [x8, #0x18]
   652: ce696d38                 ldrb w14, [x14, x13]
   656: 8c010e4a                 eor w12, w12, w14
   660: 0e1140f9                 ldr x14, [x8, #0x20]
   664: ce696d38                 ldrb w14, [x14, x13]
   668: 8c010e4a                 eor w12, w12, w14
   672: 0e1540f9                 ldr x14, [x8, #0x28]
   676: ce696d38                 ldrb w14, [x14, x13]
   680: 8c010e4a                 eor w12, w12, w14
   684: 4c692d38                 strb w12, [x10, x13]
   688: ad050091                 add x13, x13, #1
   692: bf0102eb                 cmp x13, x2
   696: 43fdff54                 b.lo 0x260
   700: 2a0540f9                 ldr x10, [x9, #8]
   704: ed030baa                 mov x13, x11
   708: 0c008052                 mov w12, #0
   712: 0e0140f9                 ldr x14, [x8]
   716: ce696d38                 ldrb w14, [x14, x13]
   720: 0e010034                 cbz w14, 0x2f0
   724: ee696e38                 ldrb w14, [x15, x14]
   728: ce990311                 add w14, w14, #0xe6
   732: dffd0371                 cmp w14, #0xff
   736: 43000054                 b.lo 0x2e8
   740: cefd0351                 sub w14, w14, #0xff
   744: 2e6a6e38                 ldrb w14, [x17, x14]
   748: 8c010e4a                 eor w12, w12, w14
   752: 0e0540f9                 ldr x14, [x8, #8]
   756: ce696d38                 ldrb w14, [x14, x13]
   760: 0e010034                 cbz w14, 0x318
   764: ee696e38                 ldrb w14, [x15, x14]
   768: ce350311                 add w14, w14, #0xcd
   772: dffd0371                 cmp w14, #0xff
   776: 43000054                 b.lo 0x310
   780: cefd0351                 sub w14, w14, #0xff
   784: 2e6a6e38                 ldrb w14, [x17, x14]
   788: 8c010e4a                 eor w12, w12, w14
   792: 0e0940f9                 ldr x14, [x8, #0x10]
   796: ce696d38                 ldrb w14, [x14, x13]
   800: 0e010034                 cbz w14, 0x340
   804: ee696e38                 ldrb w14, [x15, x14]
   808: ce2d0111                 add w14, w14, #0x4b
   812: dffd0371                 cmp w14, #0xff
   816: 43000054                 b.lo 0x338
   820: cefd0351                 sub w14, w14, #0xff
   824: 2e6a6e38                 ldrb w14, [x17, x14]
   828: 8c010e4a                 eor w12, w12, w14
   832: 0e0d40f9                 ldr x14, [x8, #0x18]
   836: ce696d38                 ldrb w14, [x14, x13]
   840: 0e010034                 cbz w14, 0x368
   844: ee696e38                 ldrb w14, [x15, x14]
   848: cec90011                 add w14, w14, #0x32
   852: dffd0371                 cmp w14, #0xff
   856: 43000054                 b.lo 0x360
   860: cefd0351                 sub w14, w14, #0xff
   864: 2e6a6e38                 ldrb w14, [x17, x14]
   868: 8c010e4a                 eor w12, w12, w14
   872: 0e1140f9                 ldr x14, [x8, #0x20]
   876: ce696d38                 ldrb w14, [x14, x13]
   880: 0e010034                 cbz w14, 0x390
   884: ee696e38                 ldrb w14, [x15, x14]
   888: ce650011                 add w14, w14, #0x19
   892: dffd0371                 cmp w14, #0xff
   896: 43000054                 b.lo 0x388
   900: cefd0351                 sub w14, w14, #0xff
   904: 2e6a6e38                 ldrb w14, [x17, x14]
   908: 8c010e4a                 eor w12, w12, w14
   912: 0e1540f9                 ldr x14, [x8, #0x28]
   916: ce696d38                 ldrb w14, [x14, x13]
   920: 8c010e4a                 eor w12, w12, w14
   924: 4c692d38                 strb w12, [x10, x13]
   928: ad050091                 add x13, x13, #1
   932: bf0102eb                 cmp x13, x2
   936: e3f8ff54                 b.lo 0x2c4
   940: 2a0940f9                 ldr x10, [x9, #0x10]
   944: ed030baa                 mov x13, x11
   948: 0c008052                 mov w12, #0
   952: 0e0140f9                 ldr x14, [x8]
   956: ce696d38                 ldrb w14, [x14, x13]
   960: 0e010034                 cbz w14, 0x3e0
   964: ee696e38                 ldrb w14, [x15, x14]
   968: ce350311                 add w14, w14, #0xcd
   972: dffd0371                 cmp w14, #0xff
   976: 43000054                 b.lo 0x3d8
   980: cefd0351                 sub w14, w14, #0xff
   984: 2e6a6e38                 ldrb w14, [x17, x14]
   988: 8c010e4a                 eor w12, w12, w14
   992: 0e0540f9                 ldr x14, [x8, #8]
   996: ce696d38                 ldrb w14, [x14, x13]
  1000: 0e010034                 cbz w14, 0x408
  1004: ee696e38                 ldrb w14, [x15, x14]
  1008: ce2d0111                 add w14, w14, #0x4b
  1012: dffd0371                 cmp w14, #0xff
  1016: 43000054                 b.lo 0x400
  1020: cefd0351                 sub w14, w14, #0xff
  1024: 2e6a6e38                 ldrb w14, [x17, x14]
  1028: 8c010e4a                 eor w12, w12, w14
  1032: 0e0940f9                 ldr x14, [x8, #0x10]
  1036: ce696d38                 ldrb w14, [x14, x13]
  1040: 0e010034                 cbz w14, 0x430
  1044: ee696e38                 ldrb w14, [x15, x14]
  1048: ce650011                 add w14, w14, #0x19
  1052: dffd0371                 cmp w14, #0xff
  1056: 43000054                 b.lo 0x428
  1060: cefd0351                 sub w14, w14, #0xff
  1064: 2e6a6e38                 ldrb w14, [x17, x14]
  1068: 8c010e4a                 eor w12, w12, w14
  1072: 0e0d40f9                 ldr x14, [x8, #0x18]
  1076: ce696d38                 ldrb w14, [x14, x13]
  1080: 0e010034                 cbz w14, 0x458
  1084: ee696e38                 ldrb w14, [x15, x14]
  1088: ce990311                 add w14, w14, #0xe6
  1092: dffd0371                 cmp w14, #0xff
  1096: 43000054                 b.lo 0x450
  1100: cefd0351                 sub w14, w14, #0xff
  1104: 2e6a6e38                 ldrb w14, [x17, x14]
  1108: 8c010e4a                 eor w12, w12, w14
  1112: 0e1140f9                 ldr x14, [x8, #0x20]
  1116: ce696d38                 ldrb w14, [x14, x13]
  1120: 0e010034                 cbz w14, 0x480
  1124: ee696e38                 ldrb w14, [x15, x14]
  1128: cec90011                 add w14, w14, #0x32
  1132: dffd0371                 cmp w14, #0xff
  1136: 43000054                 b.lo 0x478
  1140: cefd0351                 sub w14, w14, #0xff
  1144: 2e6a6e38                 ldrb w14, [x17, x14]
  1148: 8c010e4a                 eor w12, w12, w14
  1152: 0e1540f9                 ldr x14, [x8, #0x28]
  1156: ce696d38                 ldrb w14, [x14, x13]
  1160: 8c010e4a                 eor w12, w12, w14
  1164: 4c692d38                 strb w12, [x10, x13]
  1168: ad050091                 add x13, x13, #1
  1172: bf0102eb                 cmp x13, x2
  1176: e3f8ff54                 b.lo 0x3b4
  1180: ea2fc1ac                 ldp q10, q11, [sp], #0x20
  1184: e827c1ac                 ldp q8, q9, [sp], #0x20
  1188: c0035fd6                 ret 
  1192: 00000000                 udf #0
  1196: 00000000                 udf #0
  1200: 00000000                 udf #0
  1204: 00000000                 udf #0
  1208: 00000000                 udf #0
  1212: 00000000                 udf #0
