; config_t3_2118 target=neon twin=baseline topo=lrc(8,2,2)/rec{0} op=recover U=4
; code_size=1408 twins=[] uops_per_byte=3.484375
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: e80300aa                 mov x8, x0
    12: e90301aa                 mov x9, x1
    16: 902b0010                 adr x16, 0x580
    20: eb031faa                 mov x11, xzr
    24: 0e0880d2                 mov x14, #0x40
    28: 4d08ce9a                 udiv x13, x2, x14
    32: 6d1f00b4                 cbz x13, 0x40c
    36: ebe5004f                 movi v11.16b, #0xf
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 0902c03d                 ldr q9, [x16]
    76: 0a12c03d                 ldr q10, [x16, #0x40]
    80: 4401c03d                 ldr q4, [x10]
    84: 851c2b4e                 and v5.16b, v4.16b, v11.16b
    88: 86040c6f                 ushr v6.16b, v4.16b, #4
    92: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
    96: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   100: e01ca74e                 mov v0.16b, v7.16b
   104: 001c286e                 eor v0.16b, v0.16b, v8.16b
   108: 4405c03d                 ldr q4, [x10, #0x10]
   112: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   116: 86040c6f                 ushr v6.16b, v4.16b, #4
   120: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   124: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   128: e11ca74e                 mov v1.16b, v7.16b
   132: 211c286e                 eor v1.16b, v1.16b, v8.16b
   136: 4409c03d                 ldr q4, [x10, #0x20]
   140: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   144: 86040c6f                 ushr v6.16b, v4.16b, #4
   148: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   152: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   156: e21ca74e                 mov v2.16b, v7.16b
   160: 421c286e                 eor v2.16b, v2.16b, v8.16b
   164: 440dc03d                 ldr q4, [x10, #0x30]
   168: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   172: 86040c6f                 ushr v6.16b, v4.16b, #4
   176: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   180: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   184: e31ca74e                 mov v3.16b, v7.16b
   188: 631c286e                 eor v3.16b, v3.16b, v8.16b
   192: 0a0540f9                 ldr x10, [x8, #8]
   196: 4a010b8b                 add x10, x10, x11
   200: 0922c03d                 ldr q9, [x16, #0x80]
   204: 0a32c03d                 ldr q10, [x16, #0xc0]
   208: 4401c03d                 ldr q4, [x10]
   212: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   216: 86040c6f                 ushr v6.16b, v4.16b, #4
   220: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   224: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   228: 001c276e                 eor v0.16b, v0.16b, v7.16b
   232: 001c286e                 eor v0.16b, v0.16b, v8.16b
   236: 4405c03d                 ldr q4, [x10, #0x10]
   240: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   244: 86040c6f                 ushr v6.16b, v4.16b, #4
   248: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   252: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   256: 211c276e                 eor v1.16b, v1.16b, v7.16b
   260: 211c286e                 eor v1.16b, v1.16b, v8.16b
   264: 4409c03d                 ldr q4, [x10, #0x20]
   268: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   272: 86040c6f                 ushr v6.16b, v4.16b, #4
   276: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   280: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   284: 421c276e                 eor v2.16b, v2.16b, v7.16b
   288: 421c286e                 eor v2.16b, v2.16b, v8.16b
   292: 440dc03d                 ldr q4, [x10, #0x30]
   296: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   300: 86040c6f                 ushr v6.16b, v4.16b, #4
   304: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   308: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   312: 631c276e                 eor v3.16b, v3.16b, v7.16b
   316: 631c286e                 eor v3.16b, v3.16b, v8.16b
   320: 0a0940f9                 ldr x10, [x8, #0x10]
   324: 4a010b8b                 add x10, x10, x11
   328: 0942c03d                 ldr q9, [x16, #0x100]
   332: 0a52c03d                 ldr q10, [x16, #0x140]
   336: 4401c03d                 ldr q4, [x10]
   340: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   344: 86040c6f                 ushr v6.16b, v4.16b, #4
   348: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   352: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   356: 001c276e                 eor v0.16b, v0.16b, v7.16b
   360: 001c286e                 eor v0.16b, v0.16b, v8.16b
   364: 4405c03d                 ldr q4, [x10, #0x10]
   368: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   372: 86040c6f                 ushr v6.16b, v4.16b, #4
   376: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   380: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   384: 211c276e                 eor v1.16b, v1.16b, v7.16b
   388: 211c286e                 eor v1.16b, v1.16b, v8.16b
   392: 4409c03d                 ldr q4, [x10, #0x20]
   396: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   400: 86040c6f                 ushr v6.16b, v4.16b, #4
   404: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   408: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   412: 421c276e                 eor v2.16b, v2.16b, v7.16b
   416: 421c286e                 eor v2.16b, v2.16b, v8.16b
   420: 440dc03d                 ldr q4, [x10, #0x30]
   424: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   428: 86040c6f                 ushr v6.16b, v4.16b, #4
   432: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   436: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   440: 631c276e                 eor v3.16b, v3.16b, v7.16b
   444: 631c286e                 eor v3.16b, v3.16b, v8.16b
   448: 0a0d40f9                 ldr x10, [x8, #0x18]
   452: 4a010b8b                 add x10, x10, x11
   456: 0962c03d                 ldr q9, [x16, #0x180]
   460: 0a72c03d                 ldr q10, [x16, #0x1c0]
   464: 4401c03d                 ldr q4, [x10]
   468: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   472: 86040c6f                 ushr v6.16b, v4.16b, #4
   476: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   480: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   484: 001c276e                 eor v0.16b, v0.16b, v7.16b
   488: 001c286e                 eor v0.16b, v0.16b, v8.16b
   492: 4405c03d                 ldr q4, [x10, #0x10]
   496: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   500: 86040c6f                 ushr v6.16b, v4.16b, #4
   504: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   508: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   512: 211c276e                 eor v1.16b, v1.16b, v7.16b
   516: 211c286e                 eor v1.16b, v1.16b, v8.16b
   520: 4409c03d                 ldr q4, [x10, #0x20]
   524: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   528: 86040c6f                 ushr v6.16b, v4.16b, #4
   532: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   536: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   540: 421c276e                 eor v2.16b, v2.16b, v7.16b
   544: 421c286e                 eor v2.16b, v2.16b, v8.16b
   548: 440dc03d                 ldr q4, [x10, #0x30]
   552: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   556: 86040c6f                 ushr v6.16b, v4.16b, #4
   560: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   564: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   568: 631c276e                 eor v3.16b, v3.16b, v7.16b
   572: 631c286e                 eor v3.16b, v3.16b, v8.16b
   576: 0a1140f9                 ldr x10, [x8, #0x20]
   580: 4a010b8b                 add x10, x10, x11
   584: 0982c03d                 ldr q9, [x16, #0x200]
   588: 0a92c03d                 ldr q10, [x16, #0x240]
   592: 4401c03d                 ldr q4, [x10]
   596: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   600: 86040c6f                 ushr v6.16b, v4.16b, #4
   604: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   608: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   612: 001c276e                 eor v0.16b, v0.16b, v7.16b
   616: 001c286e                 eor v0.16b, v0.16b, v8.16b
   620: 4405c03d                 ldr q4, [x10, #0x10]
   624: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   628: 86040c6f                 ushr v6.16b, v4.16b, #4
   632: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   636: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   640: 211c276e                 eor v1.16b, v1.16b, v7.16b
   644: 211c286e                 eor v1.16b, v1.16b, v8.16b
   648: 4409c03d                 ldr q4, [x10, #0x20]
   652: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   656: 86040c6f                 ushr v6.16b, v4.16b, #4
   660: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   664: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   668: 421c276e                 eor v2.16b, v2.16b, v7.16b
   672: 421c286e                 eor v2.16b, v2.16b, v8.16b
   676: 440dc03d                 ldr q4, [x10, #0x30]
   680: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   684: 86040c6f                 ushr v6.16b, v4.16b, #4
   688: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   692: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   696: 631c276e                 eor v3.16b, v3.16b, v7.16b
   700: 631c286e                 eor v3.16b, v3.16b, v8.16b
   704: 0a1540f9                 ldr x10, [x8, #0x28]
   708: 4a010b8b                 add x10, x10, x11
   712: 09a2c03d                 ldr q9, [x16, #0x280]
   716: 0ab2c03d                 ldr q10, [x16, #0x2c0]
   720: 4401c03d                 ldr q4, [x10]
   724: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   728: 86040c6f                 ushr v6.16b, v4.16b, #4
   732: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   736: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   740: 001c276e                 eor v0.16b, v0.16b, v7.16b
   744: 001c286e                 eor v0.16b, v0.16b, v8.16b
   748: 4405c03d                 ldr q4, [x10, #0x10]
   752: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   756: 86040c6f                 ushr v6.16b, v4.16b, #4
   760: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   764: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   768: 211c276e                 eor v1.16b, v1.16b, v7.16b
   772: 211c286e                 eor v1.16b, v1.16b, v8.16b
   776: 4409c03d                 ldr q4, [x10, #0x20]
   780: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   784: 86040c6f                 ushr v6.16b, v4.16b, #4
   788: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   792: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   796: 421c276e                 eor v2.16b, v2.16b, v7.16b
   800: 421c286e                 eor v2.16b, v2.16b, v8.16b
   804: 440dc03d                 ldr q4, [x10, #0x30]
   808: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   812: 86040c6f                 ushr v6.16b, v4.16b, #4
   816: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   820: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   824: 631c276e                 eor v3.16b, v3.16b, v7.16b
   828: 631c286e                 eor v3.16b, v3.16b, v8.16b
   832: 0a1940f9                 ldr x10, [x8, #0x30]
   836: 4a010b8b                 add x10, x10, x11
   840: 09c2c03d                 ldr q9, [x16, #0x300]
   844: 0ad2c03d                 ldr q10, [x16, #0x340]
   848: 4401c03d                 ldr q4, [x10]
   852: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   856: 86040c6f                 ushr v6.16b, v4.16b, #4
   860: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   864: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   868: 001c276e                 eor v0.16b, v0.16b, v7.16b
   872: 001c286e                 eor v0.16b, v0.16b, v8.16b
   876: 4405c03d                 ldr q4, [x10, #0x10]
   880: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   884: 86040c6f                 ushr v6.16b, v4.16b, #4
   888: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   892: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   896: 211c276e                 eor v1.16b, v1.16b, v7.16b
   900: 211c286e                 eor v1.16b, v1.16b, v8.16b
   904: 4409c03d                 ldr q4, [x10, #0x20]
   908: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   912: 86040c6f                 ushr v6.16b, v4.16b, #4
   916: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   920: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   924: 421c276e                 eor v2.16b, v2.16b, v7.16b
   928: 421c286e                 eor v2.16b, v2.16b, v8.16b
   932: 440dc03d                 ldr q4, [x10, #0x30]
   936: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   940: 86040c6f                 ushr v6.16b, v4.16b, #4
   944: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   948: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   952: 631c276e                 eor v3.16b, v3.16b, v7.16b
   956: 631c286e                 eor v3.16b, v3.16b, v8.16b
   960: 0a1d40f9                 ldr x10, [x8, #0x38]
   964: 4a010b8b                 add x10, x10, x11
   968: 4401c03d                 ldr q4, [x10]
   972: 001c246e                 eor v0.16b, v0.16b, v4.16b
   976: 4405c03d                 ldr q4, [x10, #0x10]
   980: 211c246e                 eor v1.16b, v1.16b, v4.16b
   984: 4409c03d                 ldr q4, [x10, #0x20]
   988: 421c246e                 eor v2.16b, v2.16b, v4.16b
   992: 440dc03d                 ldr q4, [x10, #0x30]
   996: 631c246e                 eor v3.16b, v3.16b, v4.16b
  1000: 2a0140f9                 ldr x10, [x9]
  1004: 4a010b8b                 add x10, x10, x11
  1008: 4001803d                 str q0, [x10]
  1012: 4105803d                 str q1, [x10, #0x10]
  1016: 4209803d                 str q2, [x10, #0x20]
  1020: 430d803d                 str q3, [x10, #0x30]
  1024: 6b010191                 add x11, x11, #0x40
  1028: ad0500f1                 subs x13, x13, #1
  1032: 01e1ff54                 b.ne 0x28
  1036: 7f0102eb                 cmp x11, x2
  1040: 620a0054                 b.hs 0x55c
  1044: 0f020e91                 add x15, x16, #0x380
  1048: 11021291                 add x17, x16, #0x480
  1052: 2a0140f9                 ldr x10, [x9]
  1056: ed030baa                 mov x13, x11
  1060: 0c008052                 mov w12, #0
  1064: 0e0140f9                 ldr x14, [x8]
  1068: ce696d38                 ldrb w14, [x14, x13]
  1072: 0e010034                 cbz w14, 0x450
  1076: ee696e38                 ldrb w14, [x15, x14]
  1080: ce650011                 add w14, w14, #0x19
  1084: dffd0371                 cmp w14, #0xff
  1088: 43000054                 b.lo 0x448
  1092: cefd0351                 sub w14, w14, #0xff
  1096: 2e6a6e38                 ldrb w14, [x17, x14]
  1100: 8c010e4a                 eor w12, w12, w14
  1104: 0e0540f9                 ldr x14, [x8, #8]
  1108: ce696d38                 ldrb w14, [x14, x13]
  1112: 0e010034                 cbz w14, 0x478
  1116: ee696e38                 ldrb w14, [x15, x14]
  1120: cec90011                 add w14, w14, #0x32
  1124: dffd0371                 cmp w14, #0xff
  1128: 43000054                 b.lo 0x470
  1132: cefd0351                 sub w14, w14, #0xff
  1136: 2e6a6e38                 ldrb w14, [x17, x14]
  1140: 8c010e4a                 eor w12, w12, w14
  1144: 0e0940f9                 ldr x14, [x8, #0x10]
  1148: ce696d38                 ldrb w14, [x14, x13]
  1152: 0e010034                 cbz w14, 0x4a0
  1156: ee696e38                 ldrb w14, [x15, x14]
  1160: ce2d0111                 add w14, w14, #0x4b
  1164: dffd0371                 cmp w14, #0xff
  1168: 43000054                 b.lo 0x498
  1172: cefd0351                 sub w14, w14, #0xff
  1176: 2e6a6e38                 ldrb w14, [x17, x14]
  1180: 8c010e4a                 eor w12, w12, w14
  1184: 0e0d40f9                 ldr x14, [x8, #0x18]
  1188: ce696d38                 ldrb w14, [x14, x13]
  1192: 0e010034                 cbz w14, 0x4c8
  1196: ee696e38                 ldrb w14, [x15, x14]
  1200: ce910111                 add w14, w14, #0x64
  1204: dffd0371                 cmp w14, #0xff
  1208: 43000054                 b.lo 0x4c0
  1212: cefd0351                 sub w14, w14, #0xff
  1216: 2e6a6e38                 ldrb w14, [x17, x14]
  1220: 8c010e4a                 eor w12, w12, w14
  1224: 0e1140f9                 ldr x14, [x8, #0x20]
  1228: ce696d38                 ldrb w14, [x14, x13]
  1232: 0e010034                 cbz w14, 0x4f0
  1236: ee696e38                 ldrb w14, [x15, x14]
  1240: cef50111                 add w14, w14, #0x7d
  1244: dffd0371                 cmp w14, #0xff
  1248: 43000054                 b.lo 0x4e8
  1252: cefd0351                 sub w14, w14, #0xff
  1256: 2e6a6e38                 ldrb w14, [x17, x14]
  1260: 8c010e4a                 eor w12, w12, w14
  1264: 0e1540f9                 ldr x14, [x8, #0x28]
  1268: ce696d38                 ldrb w14, [x14, x13]
  1272: 0e010034                 cbz w14, 0x518
  1276: ee696e38                 ldrb w14, [x15, x14]
  1280: ce590211                 add w14, w14, #0x96
  1284: dffd0371                 cmp w14, #0xff
  1288: 43000054                 b.lo 0x510
  1292: cefd0351                 sub w14, w14, #0xff
  1296: 2e6a6e38                 ldrb w14, [x17, x14]
  1300: 8c010e4a                 eor w12, w12, w14
  1304: 0e1940f9                 ldr x14, [x8, #0x30]
  1308: ce696d38                 ldrb w14, [x14, x13]
  1312: 0e010034                 cbz w14, 0x540
  1316: ee696e38                 ldrb w14, [x15, x14]
  1320: cebd0211                 add w14, w14, #0xaf
  1324: dffd0371                 cmp w14, #0xff
  1328: 43000054                 b.lo 0x538
  1332: cefd0351                 sub w14, w14, #0xff
  1336: 2e6a6e38                 ldrb w14, [x17, x14]
  1340: 8c010e4a                 eor w12, w12, w14
  1344: 0e1d40f9                 ldr x14, [x8, #0x38]
  1348: ce696d38                 ldrb w14, [x14, x13]
  1352: 8c010e4a                 eor w12, w12, w14
  1356: 4c692d38                 strb w12, [x10, x13]
  1360: ad050091                 add x13, x13, #1
  1364: bf0102eb                 cmp x13, x2
  1368: 63f6ff54                 b.lo 0x424
  1372: ea2fc1ac                 ldp q10, q11, [sp], #0x20
  1376: e827c1ac                 ldp q8, q9, [sp], #0x20
  1380: c0035fd6                 ret 
  1384: 00000000                 udf #0
  1388: 00000000                 udf #0
  1392: 00000000                 udf #0
  1396: 00000000                 udf #0
  1400: 00000000                 udf #0
  1404: 00000000                 udf #0
