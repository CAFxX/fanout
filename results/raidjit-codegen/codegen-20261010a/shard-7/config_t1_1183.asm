; config_t1_1183 target=neon twin=baseline topo=raid6(8)/enc op=encode U=8
; code_size=1216 twins=["pipe"] uops_per_byte=4.3125
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: ec37bfad                 stp q12, q13, [sp, #-0x20]!
    12: ee3fbfad                 stp q14, q15, [sp, #-0x20]!
    16: e80300aa                 mov x8, x0
    20: e90301aa                 mov x9, x1
    24: 50250010                 adr x16, 0x4c0
    28: eb031faa                 mov x11, xzr
    32: 0e0480d2                 mov x14, #0x20
    36: 4d08ce9a                 udiv x13, x2, x14
    40: 0d1500b4                 cbz x13, 0x2c8
    44: eee5004f                 movi v14.16b, #0xf
    48: 0a0140f9                 ldr x10, [x8]
    52: 4a010b8b                 add x10, x10, x11
    56: 4601c03d                 ldr q6, [x10]
    60: 4705c03d                 ldr q7, [x10, #0x10]
    64: c01ca64e                 mov v0.16b, v6.16b
    68: e11ca74e                 mov v1.16b, v7.16b
    72: 0c02c03d                 ldr q12, [x16]
    76: 0d12c03d                 ldr q13, [x16, #0x40]
    80: c81c2e4e                 and v8.16b, v6.16b, v14.16b
    84: c9040c6f                 ushr v9.16b, v6.16b, #4
    88: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
    92: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
    96: 421daa4e                 mov v2.16b, v10.16b
   100: 421c2b6e                 eor v2.16b, v2.16b, v11.16b
   104: e81c2e4e                 and v8.16b, v7.16b, v14.16b
   108: e9040c6f                 ushr v9.16b, v7.16b, #4
   112: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   116: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   120: 431daa4e                 mov v3.16b, v10.16b
   124: 631c2b6e                 eor v3.16b, v3.16b, v11.16b
   128: 0a0540f9                 ldr x10, [x8, #8]
   132: 4a010b8b                 add x10, x10, x11
   136: 4401c03d                 ldr q4, [x10]
   140: 4505c03d                 ldr q5, [x10, #0x10]
   144: 001c246e                 eor v0.16b, v0.16b, v4.16b
   148: 211c256e                 eor v1.16b, v1.16b, v5.16b
   152: 0c22c03d                 ldr q12, [x16, #0x80]
   156: 0d32c03d                 ldr q13, [x16, #0xc0]
   160: 881c2e4e                 and v8.16b, v4.16b, v14.16b
   164: 89040c6f                 ushr v9.16b, v4.16b, #4
   168: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   172: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   176: 421c2a6e                 eor v2.16b, v2.16b, v10.16b
   180: 421c2b6e                 eor v2.16b, v2.16b, v11.16b
   184: a81c2e4e                 and v8.16b, v5.16b, v14.16b
   188: a9040c6f                 ushr v9.16b, v5.16b, #4
   192: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   196: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   200: 631c2a6e                 eor v3.16b, v3.16b, v10.16b
   204: 631c2b6e                 eor v3.16b, v3.16b, v11.16b
   208: 0a0940f9                 ldr x10, [x8, #0x10]
   212: 4a010b8b                 add x10, x10, x11
   216: 4401c03d                 ldr q4, [x10]
   220: 4505c03d                 ldr q5, [x10, #0x10]
   224: 001c246e                 eor v0.16b, v0.16b, v4.16b
   228: 211c256e                 eor v1.16b, v1.16b, v5.16b
   232: 0c42c03d                 ldr q12, [x16, #0x100]
   236: 0d52c03d                 ldr q13, [x16, #0x140]
   240: 881c2e4e                 and v8.16b, v4.16b, v14.16b
   244: 89040c6f                 ushr v9.16b, v4.16b, #4
   248: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   252: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   256: 421c2a6e                 eor v2.16b, v2.16b, v10.16b
   260: 421c2b6e                 eor v2.16b, v2.16b, v11.16b
   264: a81c2e4e                 and v8.16b, v5.16b, v14.16b
   268: a9040c6f                 ushr v9.16b, v5.16b, #4
   272: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   276: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   280: 631c2a6e                 eor v3.16b, v3.16b, v10.16b
   284: 631c2b6e                 eor v3.16b, v3.16b, v11.16b
   288: 0a0d40f9                 ldr x10, [x8, #0x18]
   292: 4a010b8b                 add x10, x10, x11
   296: 4401c03d                 ldr q4, [x10]
   300: 4505c03d                 ldr q5, [x10, #0x10]
   304: 001c246e                 eor v0.16b, v0.16b, v4.16b
   308: 211c256e                 eor v1.16b, v1.16b, v5.16b
   312: 0c62c03d                 ldr q12, [x16, #0x180]
   316: 0d72c03d                 ldr q13, [x16, #0x1c0]
   320: 881c2e4e                 and v8.16b, v4.16b, v14.16b
   324: 89040c6f                 ushr v9.16b, v4.16b, #4
   328: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   332: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   336: 421c2a6e                 eor v2.16b, v2.16b, v10.16b
   340: 421c2b6e                 eor v2.16b, v2.16b, v11.16b
   344: a81c2e4e                 and v8.16b, v5.16b, v14.16b
   348: a9040c6f                 ushr v9.16b, v5.16b, #4
   352: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   356: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   360: 631c2a6e                 eor v3.16b, v3.16b, v10.16b
   364: 631c2b6e                 eor v3.16b, v3.16b, v11.16b
   368: 0a1140f9                 ldr x10, [x8, #0x20]
   372: 4a010b8b                 add x10, x10, x11
   376: 4401c03d                 ldr q4, [x10]
   380: 4505c03d                 ldr q5, [x10, #0x10]
   384: 001c246e                 eor v0.16b, v0.16b, v4.16b
   388: 211c256e                 eor v1.16b, v1.16b, v5.16b
   392: 0c82c03d                 ldr q12, [x16, #0x200]
   396: 0d92c03d                 ldr q13, [x16, #0x240]
   400: 881c2e4e                 and v8.16b, v4.16b, v14.16b
   404: 89040c6f                 ushr v9.16b, v4.16b, #4
   408: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   412: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   416: 421c2a6e                 eor v2.16b, v2.16b, v10.16b
   420: 421c2b6e                 eor v2.16b, v2.16b, v11.16b
   424: a81c2e4e                 and v8.16b, v5.16b, v14.16b
   428: a9040c6f                 ushr v9.16b, v5.16b, #4
   432: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   436: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   440: 631c2a6e                 eor v3.16b, v3.16b, v10.16b
   444: 631c2b6e                 eor v3.16b, v3.16b, v11.16b
   448: 0a1540f9                 ldr x10, [x8, #0x28]
   452: 4a010b8b                 add x10, x10, x11
   456: 4401c03d                 ldr q4, [x10]
   460: 4505c03d                 ldr q5, [x10, #0x10]
   464: 001c246e                 eor v0.16b, v0.16b, v4.16b
   468: 211c256e                 eor v1.16b, v1.16b, v5.16b
   472: 0ca2c03d                 ldr q12, [x16, #0x280]
   476: 0db2c03d                 ldr q13, [x16, #0x2c0]
   480: 881c2e4e                 and v8.16b, v4.16b, v14.16b
   484: 89040c6f                 ushr v9.16b, v4.16b, #4
   488: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   492: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   496: 421c2a6e                 eor v2.16b, v2.16b, v10.16b
   500: 421c2b6e                 eor v2.16b, v2.16b, v11.16b
   504: a81c2e4e                 and v8.16b, v5.16b, v14.16b
   508: a9040c6f                 ushr v9.16b, v5.16b, #4
   512: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   516: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   520: 631c2a6e                 eor v3.16b, v3.16b, v10.16b
   524: 631c2b6e                 eor v3.16b, v3.16b, v11.16b
   528: 0a1940f9                 ldr x10, [x8, #0x30]
   532: 4a010b8b                 add x10, x10, x11
   536: 4401c03d                 ldr q4, [x10]
   540: 4505c03d                 ldr q5, [x10, #0x10]
   544: 001c246e                 eor v0.16b, v0.16b, v4.16b
   548: 211c256e                 eor v1.16b, v1.16b, v5.16b
   552: 0cc2c03d                 ldr q12, [x16, #0x300]
   556: 0dd2c03d                 ldr q13, [x16, #0x340]
   560: 881c2e4e                 and v8.16b, v4.16b, v14.16b
   564: 89040c6f                 ushr v9.16b, v4.16b, #4
   568: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   572: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   576: 421c2a6e                 eor v2.16b, v2.16b, v10.16b
   580: 421c2b6e                 eor v2.16b, v2.16b, v11.16b
   584: a81c2e4e                 and v8.16b, v5.16b, v14.16b
   588: a9040c6f                 ushr v9.16b, v5.16b, #4
   592: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   596: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   600: 631c2a6e                 eor v3.16b, v3.16b, v10.16b
   604: 631c2b6e                 eor v3.16b, v3.16b, v11.16b
   608: 0a1d40f9                 ldr x10, [x8, #0x38]
   612: 4a010b8b                 add x10, x10, x11
   616: 4401c03d                 ldr q4, [x10]
   620: 4505c03d                 ldr q5, [x10, #0x10]
   624: 001c246e                 eor v0.16b, v0.16b, v4.16b
   628: 211c256e                 eor v1.16b, v1.16b, v5.16b
   632: 421c246e                 eor v2.16b, v2.16b, v4.16b
   636: 631c256e                 eor v3.16b, v3.16b, v5.16b
   640: bf0500f1                 cmp x13, #1
   644: cd000054                 b.le 0x29c
   648: 0a0140f9                 ldr x10, [x8]
   652: 4a010b8b                 add x10, x10, x11
   656: 4a810091                 add x10, x10, #0x20
   660: 4601c03d                 ldr q6, [x10]
   664: 4705c03d                 ldr q7, [x10, #0x10]
   668: 2a0140f9                 ldr x10, [x9]
   672: 4a010b8b                 add x10, x10, x11
   676: 4001803d                 str q0, [x10]
   680: 4105803d                 str q1, [x10, #0x10]
   684: 2a0540f9                 ldr x10, [x9, #8]
   688: 4a010b8b                 add x10, x10, x11
   692: 4201803d                 str q2, [x10]
   696: 4305803d                 str q3, [x10, #0x10]
   700: 6b810091                 add x11, x11, #0x20
   704: ad0500f1                 subs x13, x13, #1
   708: e1ebff54                 b.ne 0x40
   712: 7f0102eb                 cmp x11, x2
   716: 420e0054                 b.hs 0x494
   720: 0f020e91                 add x15, x16, #0x380
   724: 11021291                 add x17, x16, #0x480
   728: 2a0140f9                 ldr x10, [x9]
   732: ed030baa                 mov x13, x11
   736: 0c008052                 mov w12, #0
   740: 0e0140f9                 ldr x14, [x8]
   744: ce696d38                 ldrb w14, [x14, x13]
   748: 8c010e4a                 eor w12, w12, w14
   752: 0e0540f9                 ldr x14, [x8, #8]
   756: ce696d38                 ldrb w14, [x14, x13]
   760: 8c010e4a                 eor w12, w12, w14
   764: 0e0940f9                 ldr x14, [x8, #0x10]
   768: ce696d38                 ldrb w14, [x14, x13]
   772: 8c010e4a                 eor w12, w12, w14
   776: 0e0d40f9                 ldr x14, [x8, #0x18]
   780: ce696d38                 ldrb w14, [x14, x13]
   784: 8c010e4a                 eor w12, w12, w14
   788: 0e1140f9                 ldr x14, [x8, #0x20]
   792: ce696d38                 ldrb w14, [x14, x13]
   796: 8c010e4a                 eor w12, w12, w14
   800: 0e1540f9                 ldr x14, [x8, #0x28]
   804: ce696d38                 ldrb w14, [x14, x13]
   808: 8c010e4a                 eor w12, w12, w14
   812: 0e1940f9                 ldr x14, [x8, #0x30]
   816: ce696d38                 ldrb w14, [x14, x13]
   820: 8c010e4a                 eor w12, w12, w14
   824: 0e1d40f9                 ldr x14, [x8, #0x38]
   828: ce696d38                 ldrb w14, [x14, x13]
   832: 8c010e4a                 eor w12, w12, w14
   836: 4c692d38                 strb w12, [x10, x13]
   840: ad050091                 add x13, x13, #1
   844: bf0102eb                 cmp x13, x2
   848: 83fcff54                 b.lo 0x2e0
   852: 2a0540f9                 ldr x10, [x9, #8]
   856: ed030baa                 mov x13, x11
   860: 0c008052                 mov w12, #0
   864: 0e0140f9                 ldr x14, [x8]
   868: ce696d38                 ldrb w14, [x14, x13]
   872: 0e010034                 cbz w14, 0x388
   876: ee696e38                 ldrb w14, [x15, x14]
   880: ce990311                 add w14, w14, #0xe6
   884: dffd0371                 cmp w14, #0xff
   888: 43000054                 b.lo 0x380
   892: cefd0351                 sub w14, w14, #0xff
   896: 2e6a6e38                 ldrb w14, [x17, x14]
   900: 8c010e4a                 eor w12, w12, w14
   904: 0e0540f9                 ldr x14, [x8, #8]
   908: ce696d38                 ldrb w14, [x14, x13]
   912: 0e010034                 cbz w14, 0x3b0
   916: ee696e38                 ldrb w14, [x15, x14]
   920: ce350311                 add w14, w14, #0xcd
   924: dffd0371                 cmp w14, #0xff
   928: 43000054                 b.lo 0x3a8
   932: cefd0351                 sub w14, w14, #0xff
   936: 2e6a6e38                 ldrb w14, [x17, x14]
   940: 8c010e4a                 eor w12, w12, w14
   944: 0e0940f9                 ldr x14, [x8, #0x10]
   948: ce696d38                 ldrb w14, [x14, x13]
   952: 0e010034                 cbz w14, 0x3d8
   956: ee696e38                 ldrb w14, [x15, x14]
   960: ce9d0311                 add w14, w14, #0xe7
   964: dffd0371                 cmp w14, #0xff
   968: 43000054                 b.lo 0x3d0
   972: cefd0351                 sub w14, w14, #0xff
   976: 2e6a6e38                 ldrb w14, [x17, x14]
   980: 8c010e4a                 eor w12, w12, w14
   984: 0e0d40f9                 ldr x14, [x8, #0x18]
   988: ce696d38                 ldrb w14, [x14, x13]
   992: 0e010034                 cbz w14, 0x400
   996: ee696e38                 ldrb w14, [x15, x14]
  1000: ce390311                 add w14, w14, #0xce
  1004: dffd0371                 cmp w14, #0xff
  1008: 43000054                 b.lo 0x3f8
  1012: cefd0351                 sub w14, w14, #0xff
  1016: 2e6a6e38                 ldrb w14, [x17, x14]
  1020: 8c010e4a                 eor w12, w12, w14
  1024: 0e1140f9                 ldr x14, [x8, #0x20]
  1028: ce696d38                 ldrb w14, [x14, x13]
  1032: 0e010034                 cbz w14, 0x428
  1036: ee696e38                 ldrb w14, [x15, x14]
  1040: ce2d0111                 add w14, w14, #0x4b
  1044: dffd0371                 cmp w14, #0xff
  1048: 43000054                 b.lo 0x420
  1052: cefd0351                 sub w14, w14, #0xff
  1056: 2e6a6e38                 ldrb w14, [x17, x14]
  1060: 8c010e4a                 eor w12, w12, w14
  1064: 0e1540f9                 ldr x14, [x8, #0x28]
  1068: ce696d38                 ldrb w14, [x14, x13]
  1072: 0e010034                 cbz w14, 0x450
  1076: ee696e38                 ldrb w14, [x15, x14]
  1080: cec90011                 add w14, w14, #0x32
  1084: dffd0371                 cmp w14, #0xff
  1088: 43000054                 b.lo 0x448
  1092: cefd0351                 sub w14, w14, #0xff
  1096: 2e6a6e38                 ldrb w14, [x17, x14]
  1100: 8c010e4a                 eor w12, w12, w14
  1104: 0e1940f9                 ldr x14, [x8, #0x30]
  1108: ce696d38                 ldrb w14, [x14, x13]
  1112: 0e010034                 cbz w14, 0x478
  1116: ee696e38                 ldrb w14, [x15, x14]
  1120: ce650011                 add w14, w14, #0x19
  1124: dffd0371                 cmp w14, #0xff
  1128: 43000054                 b.lo 0x470
  1132: cefd0351                 sub w14, w14, #0xff
  1136: 2e6a6e38                 ldrb w14, [x17, x14]
  1140: 8c010e4a                 eor w12, w12, w14
  1144: 0e1d40f9                 ldr x14, [x8, #0x38]
  1148: ce696d38                 ldrb w14, [x14, x13]
  1152: 8c010e4a                 eor w12, w12, w14
  1156: 4c692d38                 strb w12, [x10, x13]
  1160: ad050091                 add x13, x13, #1
  1164: bf0102eb                 cmp x13, x2
  1168: 63f6ff54                 b.lo 0x35c
  1172: ee3fc1ac                 ldp q14, q15, [sp], #0x20
  1176: ec37c1ac                 ldp q12, q13, [sp], #0x20
  1180: ea2fc1ac                 ldp q10, q11, [sp], #0x20
  1184: e827c1ac                 ldp q8, q9, [sp], #0x20
  1188: c0035fd6                 ret 
  1192: 00000000                 udf #0
  1196: 00000000                 udf #0
  1200: 00000000                 udf #0
  1204: 00000000                 udf #0
  1208: 00000000                 udf #0
  1212: 00000000                 udf #0
