; config_t1_1391 target=neon twin=baseline topo=rs(6,3)/upd{0} op=parity_update U=8
; code_size=704 twins=[] uops_per_byte=1.90625
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: ec37bfad                 stp q12, q13, [sp, #-0x20]!
    12: e80300aa                 mov x8, x0
    16: e90301aa                 mov x9, x1
    20: 70150010                 adr x16, 0x2c0
    24: eb031faa                 mov x11, xzr
    28: 0e0480d2                 mov x14, #0x20
    32: 4d08ce9a                 udiv x13, x2, x14
    36: cd0a00b4                 cbz x13, 0x17c
    40: ede5004f                 movi v13.16b, #0xf
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0940f9                 ldr x10, [x8, #0x10]
    68: 4a010b8b                 add x10, x10, x11
    72: 4001c03d                 ldr q0, [x10]
    76: 4105c03d                 ldr q1, [x10, #0x10]
    80: 0a0d40f9                 ldr x10, [x8, #0x18]
    84: 4a010b8b                 add x10, x10, x11
    88: 4201c03d                 ldr q2, [x10]
    92: 4305c03d                 ldr q3, [x10, #0x10]
    96: 0a1140f9                 ldr x10, [x8, #0x20]
   100: 4a010b8b                 add x10, x10, x11
   104: 4401c03d                 ldr q4, [x10]
   108: 4505c03d                 ldr q5, [x10, #0x10]
   112: 0a0140f9                 ldr x10, [x8]
   116: 4a010b8b                 add x10, x10, x11
   120: 030540f9                 ldr x3, [x8, #8]
   124: 63000b8b                 add x3, x3, x11
   128: 4601c03d                 ldr q6, [x10]
   132: 6900c03d                 ldr q9, [x3]
   136: c61c296e                 eor v6.16b, v6.16b, v9.16b
   140: 001c266e                 eor v0.16b, v0.16b, v6.16b
   144: 4605c03d                 ldr q6, [x10, #0x10]
   148: 6904c03d                 ldr q9, [x3, #0x10]
   152: c61c296e                 eor v6.16b, v6.16b, v9.16b
   156: 211c266e                 eor v1.16b, v1.16b, v6.16b
   160: 0b02c03d                 ldr q11, [x16]
   164: 0c12c03d                 ldr q12, [x16, #0x40]
   168: 4601c03d                 ldr q6, [x10]
   172: 6900c03d                 ldr q9, [x3]
   176: c61c296e                 eor v6.16b, v6.16b, v9.16b
   180: c71c2d4e                 and v7.16b, v6.16b, v13.16b
   184: c8040c6f                 ushr v8.16b, v6.16b, #4
   188: 6901074e                 tbl v9.16b, { v11.16b }, v7.16b
   192: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   196: 421c296e                 eor v2.16b, v2.16b, v9.16b
   200: 421c2a6e                 eor v2.16b, v2.16b, v10.16b
   204: 4605c03d                 ldr q6, [x10, #0x10]
   208: 6904c03d                 ldr q9, [x3, #0x10]
   212: c61c296e                 eor v6.16b, v6.16b, v9.16b
   216: c71c2d4e                 and v7.16b, v6.16b, v13.16b
   220: c8040c6f                 ushr v8.16b, v6.16b, #4
   224: 6901074e                 tbl v9.16b, { v11.16b }, v7.16b
   228: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   232: 631c296e                 eor v3.16b, v3.16b, v9.16b
   236: 631c2a6e                 eor v3.16b, v3.16b, v10.16b
   240: 0b22c03d                 ldr q11, [x16, #0x80]
   244: 0c32c03d                 ldr q12, [x16, #0xc0]
   248: 4601c03d                 ldr q6, [x10]
   252: 6900c03d                 ldr q9, [x3]
   256: c61c296e                 eor v6.16b, v6.16b, v9.16b
   260: c71c2d4e                 and v7.16b, v6.16b, v13.16b
   264: c8040c6f                 ushr v8.16b, v6.16b, #4
   268: 6901074e                 tbl v9.16b, { v11.16b }, v7.16b
   272: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   276: 841c296e                 eor v4.16b, v4.16b, v9.16b
   280: 841c2a6e                 eor v4.16b, v4.16b, v10.16b
   284: 4605c03d                 ldr q6, [x10, #0x10]
   288: 6904c03d                 ldr q9, [x3, #0x10]
   292: c61c296e                 eor v6.16b, v6.16b, v9.16b
   296: c71c2d4e                 and v7.16b, v6.16b, v13.16b
   300: c8040c6f                 ushr v8.16b, v6.16b, #4
   304: 6901074e                 tbl v9.16b, { v11.16b }, v7.16b
   308: 8a01084e                 tbl v10.16b, { v12.16b }, v8.16b
   312: a51c296e                 eor v5.16b, v5.16b, v9.16b
   316: a51c2a6e                 eor v5.16b, v5.16b, v10.16b
   320: 2a0140f9                 ldr x10, [x9]
   324: 4a010b8b                 add x10, x10, x11
   328: 4001803d                 str q0, [x10]
   332: 4105803d                 str q1, [x10, #0x10]
   336: 2a0540f9                 ldr x10, [x9, #8]
   340: 4a010b8b                 add x10, x10, x11
   344: 4201803d                 str q2, [x10]
   348: 4305803d                 str q3, [x10, #0x10]
   352: 2a0940f9                 ldr x10, [x9, #0x10]
   356: 4a010b8b                 add x10, x10, x11
   360: 4401803d                 str q4, [x10]
   364: 4505803d                 str q5, [x10, #0x10]
   368: 6b810091                 add x11, x11, #0x20
   372: ad0500f1                 subs x13, x13, #1
   376: a1f5ff54                 b.ne 0x2c
   380: 7f0102eb                 cmp x11, x2
   384: 22090054                 b.hs 0x2a4
   388: 0f020491                 add x15, x16, #0x100
   392: 11020891                 add x17, x16, #0x200
   396: 2a0140f9                 ldr x10, [x9]
   400: ed030baa                 mov x13, x11
   404: 0e0940f9                 ldr x14, [x8, #0x10]
   408: cc696d38                 ldrb w12, [x14, x13]
   412: 0e0140f9                 ldr x14, [x8]
   416: ce696d38                 ldrb w14, [x14, x13]
   420: 8c010e4a                 eor w12, w12, w14
   424: 0e0540f9                 ldr x14, [x8, #8]
   428: ce696d38                 ldrb w14, [x14, x13]
   432: 8c010e4a                 eor w12, w12, w14
   436: 4c692d38                 strb w12, [x10, x13]
   440: ad050091                 add x13, x13, #1
   444: bf0102eb                 cmp x13, x2
   448: a3feff54                 b.lo 0x194
   452: 2a0540f9                 ldr x10, [x9, #8]
   456: ed030baa                 mov x13, x11
   460: 0e0d40f9                 ldr x14, [x8, #0x18]
   464: cc696d38                 ldrb w12, [x14, x13]
   468: 0e0140f9                 ldr x14, [x8]
   472: ce696d38                 ldrb w14, [x14, x13]
   476: 0e010034                 cbz w14, 0x1fc
   480: ee696e38                 ldrb w14, [x15, x14]
   484: ce990311                 add w14, w14, #0xe6
   488: dffd0371                 cmp w14, #0xff
   492: 43000054                 b.lo 0x1f4
   496: cefd0351                 sub w14, w14, #0xff
   500: 2e6a6e38                 ldrb w14, [x17, x14]
   504: 8c010e4a                 eor w12, w12, w14
   508: 0e0540f9                 ldr x14, [x8, #8]
   512: ce696d38                 ldrb w14, [x14, x13]
   516: 0e010034                 cbz w14, 0x224
   520: ee696e38                 ldrb w14, [x15, x14]
   524: ce990311                 add w14, w14, #0xe6
   528: dffd0371                 cmp w14, #0xff
   532: 43000054                 b.lo 0x21c
   536: cefd0351                 sub w14, w14, #0xff
   540: 2e6a6e38                 ldrb w14, [x17, x14]
   544: 8c010e4a                 eor w12, w12, w14
   548: 4c692d38                 strb w12, [x10, x13]
   552: ad050091                 add x13, x13, #1
   556: bf0102eb                 cmp x13, x2
   560: e3fcff54                 b.lo 0x1cc
   564: 2a0940f9                 ldr x10, [x9, #0x10]
   568: ed030baa                 mov x13, x11
   572: 0e1140f9                 ldr x14, [x8, #0x20]
   576: cc696d38                 ldrb w12, [x14, x13]
   580: 0e0140f9                 ldr x14, [x8]
   584: ce696d38                 ldrb w14, [x14, x13]
   588: 0e010034                 cbz w14, 0x26c
   592: ee696e38                 ldrb w14, [x15, x14]
   596: ce350311                 add w14, w14, #0xcd
   600: dffd0371                 cmp w14, #0xff
   604: 43000054                 b.lo 0x264
   608: cefd0351                 sub w14, w14, #0xff
   612: 2e6a6e38                 ldrb w14, [x17, x14]
   616: 8c010e4a                 eor w12, w12, w14
   620: 0e0540f9                 ldr x14, [x8, #8]
   624: ce696d38                 ldrb w14, [x14, x13]
   628: 0e010034                 cbz w14, 0x294
   632: ee696e38                 ldrb w14, [x15, x14]
   636: ce350311                 add w14, w14, #0xcd
   640: dffd0371                 cmp w14, #0xff
   644: 43000054                 b.lo 0x28c
   648: cefd0351                 sub w14, w14, #0xff
   652: 2e6a6e38                 ldrb w14, [x17, x14]
   656: 8c010e4a                 eor w12, w12, w14
   660: 4c692d38                 strb w12, [x10, x13]
   664: ad050091                 add x13, x13, #1
   668: bf0102eb                 cmp x13, x2
   672: e3fcff54                 b.lo 0x23c
   676: ec37c1ac                 ldp q12, q13, [sp], #0x20
   680: ea2fc1ac                 ldp q10, q11, [sp], #0x20
   684: e827c1ac                 ldp q8, q9, [sp], #0x20
   688: c0035fd6                 ret 
   692: 00000000                 udf #0
   696: 00000000                 udf #0
   700: 00000000                 udf #0
