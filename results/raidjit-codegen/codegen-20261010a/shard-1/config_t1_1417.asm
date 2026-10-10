; config_t1_1417 target=neon twin=baseline topo=raid6(8)/upd{0,1} op=parity_update U=8
; code_size=960 twins=[] uops_per_byte=1.953125
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: ec37bfad                 stp q12, q13, [sp, #-0x20]!
    12: ee3fbfad                 stp q14, q15, [sp, #-0x20]!
    16: e80300aa                 mov x8, x0
    20: e90301aa                 mov x9, x1
    24: 501d0010                 adr x16, 0x3c0
    28: eb031faa                 mov x11, xzr
    32: 0e0880d2                 mov x14, #0x40
    36: 4d08ce9a                 udiv x13, x2, x14
    40: ad1200b4                 cbz x13, 0x27c
    44: efe5004f                 movi v15.16b, #0xf
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a1140f9                 ldr x10, [x8, #0x20]
    68: 4a010b8b                 add x10, x10, x11
    72: 4001c03d                 ldr q0, [x10]
    76: 4105c03d                 ldr q1, [x10, #0x10]
    80: 4209c03d                 ldr q2, [x10, #0x20]
    84: 430dc03d                 ldr q3, [x10, #0x30]
    88: 0a1540f9                 ldr x10, [x8, #0x28]
    92: 4a010b8b                 add x10, x10, x11
    96: 4401c03d                 ldr q4, [x10]
   100: 4505c03d                 ldr q5, [x10, #0x10]
   104: 4609c03d                 ldr q6, [x10, #0x20]
   108: 470dc03d                 ldr q7, [x10, #0x30]
   112: 0a0140f9                 ldr x10, [x8]
   116: 4a010b8b                 add x10, x10, x11
   120: 030940f9                 ldr x3, [x8, #0x10]
   124: 63000b8b                 add x3, x3, x11
   128: 4801c03d                 ldr q8, [x10]
   132: 6b00c03d                 ldr q11, [x3]
   136: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   140: 001c286e                 eor v0.16b, v0.16b, v8.16b
   144: 4805c03d                 ldr q8, [x10, #0x10]
   148: 6b04c03d                 ldr q11, [x3, #0x10]
   152: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   156: 211c286e                 eor v1.16b, v1.16b, v8.16b
   160: 4809c03d                 ldr q8, [x10, #0x20]
   164: 6b08c03d                 ldr q11, [x3, #0x20]
   168: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   172: 421c286e                 eor v2.16b, v2.16b, v8.16b
   176: 480dc03d                 ldr q8, [x10, #0x30]
   180: 6b0cc03d                 ldr q11, [x3, #0x30]
   184: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   188: 631c286e                 eor v3.16b, v3.16b, v8.16b
   192: 0d02c03d                 ldr q13, [x16]
   196: 0e12c03d                 ldr q14, [x16, #0x40]
   200: 4801c03d                 ldr q8, [x10]
   204: 6b00c03d                 ldr q11, [x3]
   208: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   212: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   216: 0a050c6f                 ushr v10.16b, v8.16b, #4
   220: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   224: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   228: 841c2b6e                 eor v4.16b, v4.16b, v11.16b
   232: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   236: 4805c03d                 ldr q8, [x10, #0x10]
   240: 6b04c03d                 ldr q11, [x3, #0x10]
   244: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   248: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   252: 0a050c6f                 ushr v10.16b, v8.16b, #4
   256: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   260: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   264: a51c2b6e                 eor v5.16b, v5.16b, v11.16b
   268: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   272: 4809c03d                 ldr q8, [x10, #0x20]
   276: 6b08c03d                 ldr q11, [x3, #0x20]
   280: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   284: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   288: 0a050c6f                 ushr v10.16b, v8.16b, #4
   292: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   296: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   300: c61c2b6e                 eor v6.16b, v6.16b, v11.16b
   304: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   308: 480dc03d                 ldr q8, [x10, #0x30]
   312: 6b0cc03d                 ldr q11, [x3, #0x30]
   316: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   320: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   324: 0a050c6f                 ushr v10.16b, v8.16b, #4
   328: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   332: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   336: e71c2b6e                 eor v7.16b, v7.16b, v11.16b
   340: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   344: 0a0540f9                 ldr x10, [x8, #8]
   348: 4a010b8b                 add x10, x10, x11
   352: 030d40f9                 ldr x3, [x8, #0x18]
   356: 63000b8b                 add x3, x3, x11
   360: 4801c03d                 ldr q8, [x10]
   364: 6b00c03d                 ldr q11, [x3]
   368: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   372: 001c286e                 eor v0.16b, v0.16b, v8.16b
   376: 4805c03d                 ldr q8, [x10, #0x10]
   380: 6b04c03d                 ldr q11, [x3, #0x10]
   384: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   388: 211c286e                 eor v1.16b, v1.16b, v8.16b
   392: 4809c03d                 ldr q8, [x10, #0x20]
   396: 6b08c03d                 ldr q11, [x3, #0x20]
   400: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   404: 421c286e                 eor v2.16b, v2.16b, v8.16b
   408: 480dc03d                 ldr q8, [x10, #0x30]
   412: 6b0cc03d                 ldr q11, [x3, #0x30]
   416: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   420: 631c286e                 eor v3.16b, v3.16b, v8.16b
   424: 0d22c03d                 ldr q13, [x16, #0x80]
   428: 0e32c03d                 ldr q14, [x16, #0xc0]
   432: 4801c03d                 ldr q8, [x10]
   436: 6b00c03d                 ldr q11, [x3]
   440: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   444: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   448: 0a050c6f                 ushr v10.16b, v8.16b, #4
   452: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   456: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   460: 841c2b6e                 eor v4.16b, v4.16b, v11.16b
   464: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   468: 4805c03d                 ldr q8, [x10, #0x10]
   472: 6b04c03d                 ldr q11, [x3, #0x10]
   476: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   480: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   484: 0a050c6f                 ushr v10.16b, v8.16b, #4
   488: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   492: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   496: a51c2b6e                 eor v5.16b, v5.16b, v11.16b
   500: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   504: 4809c03d                 ldr q8, [x10, #0x20]
   508: 6b08c03d                 ldr q11, [x3, #0x20]
   512: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   516: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   520: 0a050c6f                 ushr v10.16b, v8.16b, #4
   524: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   528: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   532: c61c2b6e                 eor v6.16b, v6.16b, v11.16b
   536: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   540: 480dc03d                 ldr q8, [x10, #0x30]
   544: 6b0cc03d                 ldr q11, [x3, #0x30]
   548: 081d2b6e                 eor v8.16b, v8.16b, v11.16b
   552: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   556: 0a050c6f                 ushr v10.16b, v8.16b, #4
   560: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   564: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   568: e71c2b6e                 eor v7.16b, v7.16b, v11.16b
   572: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   576: 2a0140f9                 ldr x10, [x9]
   580: 4a010b8b                 add x10, x10, x11
   584: 4001803d                 str q0, [x10]
   588: 4105803d                 str q1, [x10, #0x10]
   592: 4209803d                 str q2, [x10, #0x20]
   596: 430d803d                 str q3, [x10, #0x30]
   600: 2a0540f9                 ldr x10, [x9, #8]
   604: 4a010b8b                 add x10, x10, x11
   608: 4401803d                 str q4, [x10]
   612: 4505803d                 str q5, [x10, #0x10]
   616: 4609803d                 str q6, [x10, #0x20]
   620: 470d803d                 str q7, [x10, #0x30]
   624: 6b010191                 add x11, x11, #0x40
   628: ad0500f1                 subs x13, x13, #1
   632: c1edff54                 b.ne 0x30
   636: 7f0102eb                 cmp x11, x2
   640: e2080054                 b.hs 0x39c
   644: 0f020491                 add x15, x16, #0x100
   648: 11020891                 add x17, x16, #0x200
   652: 2a0140f9                 ldr x10, [x9]
   656: ed030baa                 mov x13, x11
   660: 0e1140f9                 ldr x14, [x8, #0x20]
   664: cc696d38                 ldrb w12, [x14, x13]
   668: 0e0140f9                 ldr x14, [x8]
   672: ce696d38                 ldrb w14, [x14, x13]
   676: 8c010e4a                 eor w12, w12, w14
   680: 0e0940f9                 ldr x14, [x8, #0x10]
   684: ce696d38                 ldrb w14, [x14, x13]
   688: 8c010e4a                 eor w12, w12, w14
   692: 0e0540f9                 ldr x14, [x8, #8]
   696: ce696d38                 ldrb w14, [x14, x13]
   700: 8c010e4a                 eor w12, w12, w14
   704: 0e0d40f9                 ldr x14, [x8, #0x18]
   708: ce696d38                 ldrb w14, [x14, x13]
   712: 8c010e4a                 eor w12, w12, w14
   716: 4c692d38                 strb w12, [x10, x13]
   720: ad050091                 add x13, x13, #1
   724: bf0102eb                 cmp x13, x2
   728: e3fdff54                 b.lo 0x294
   732: 2a0540f9                 ldr x10, [x9, #8]
   736: ed030baa                 mov x13, x11
   740: 0e1540f9                 ldr x14, [x8, #0x28]
   744: cc696d38                 ldrb w12, [x14, x13]
   748: 0e0140f9                 ldr x14, [x8]
   752: ce696d38                 ldrb w14, [x14, x13]
   756: 0e010034                 cbz w14, 0x314
   760: ee696e38                 ldrb w14, [x15, x14]
   764: ce990311                 add w14, w14, #0xe6
   768: dffd0371                 cmp w14, #0xff
   772: 43000054                 b.lo 0x30c
   776: cefd0351                 sub w14, w14, #0xff
   780: 2e6a6e38                 ldrb w14, [x17, x14]
   784: 8c010e4a                 eor w12, w12, w14
   788: 0e0940f9                 ldr x14, [x8, #0x10]
   792: ce696d38                 ldrb w14, [x14, x13]
   796: 0e010034                 cbz w14, 0x33c
   800: ee696e38                 ldrb w14, [x15, x14]
   804: ce990311                 add w14, w14, #0xe6
   808: dffd0371                 cmp w14, #0xff
   812: 43000054                 b.lo 0x334
   816: cefd0351                 sub w14, w14, #0xff
   820: 2e6a6e38                 ldrb w14, [x17, x14]
   824: 8c010e4a                 eor w12, w12, w14
   828: 0e0540f9                 ldr x14, [x8, #8]
   832: ce696d38                 ldrb w14, [x14, x13]
   836: 0e010034                 cbz w14, 0x364
   840: ee696e38                 ldrb w14, [x15, x14]
   844: ce350311                 add w14, w14, #0xcd
   848: dffd0371                 cmp w14, #0xff
   852: 43000054                 b.lo 0x35c
   856: cefd0351                 sub w14, w14, #0xff
   860: 2e6a6e38                 ldrb w14, [x17, x14]
   864: 8c010e4a                 eor w12, w12, w14
   868: 0e0d40f9                 ldr x14, [x8, #0x18]
   872: ce696d38                 ldrb w14, [x14, x13]
   876: 0e010034                 cbz w14, 0x38c
   880: ee696e38                 ldrb w14, [x15, x14]
   884: ce350311                 add w14, w14, #0xcd
   888: dffd0371                 cmp w14, #0xff
   892: 43000054                 b.lo 0x384
   896: cefd0351                 sub w14, w14, #0xff
   900: 2e6a6e38                 ldrb w14, [x17, x14]
   904: 8c010e4a                 eor w12, w12, w14
   908: 4c692d38                 strb w12, [x10, x13]
   912: ad050091                 add x13, x13, #1
   916: bf0102eb                 cmp x13, x2
   920: 63faff54                 b.lo 0x2e4
   924: ee3fc1ac                 ldp q14, q15, [sp], #0x20
   928: ec37c1ac                 ldp q12, q13, [sp], #0x20
   932: ea2fc1ac                 ldp q10, q11, [sp], #0x20
   936: e827c1ac                 ldp q8, q9, [sp], #0x20
   940: c0035fd6                 ret 
   944: 00000000                 udf #0
   948: 00000000                 udf #0
   952: 00000000                 udf #0
   956: 00000000                 udf #0
