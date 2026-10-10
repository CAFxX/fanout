; config_t1_1181 target=neon twin=pipe topo=raid6(8)/enc op=encode U=1
; code_size=960 twins=["pipe"] uops_per_byte=4.8125
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: e80300aa                 mov x8, x0
    12: e90301aa                 mov x9, x1
    16: 901d0010                 adr x16, 0x3c0
    20: eb031faa                 mov x11, xzr
    24: 0e0280d2                 mov x14, #0x10
    28: 4d08ce9a                 udiv x13, x2, x14
    32: ad0d00b4                 cbz x13, 0x1d4
    36: eae5004f                 movi v10.16b, #0xf
    40: 0a0140f9                 ldr x10, [x8]
    44: 4a010b8b                 add x10, x10, x11
    48: 4301c03d                 ldr q3, [x10]
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 601ca34e                 mov v0.16b, v3.16b
    68: 0802c03d                 ldr q8, [x16]
    72: 0912c03d                 ldr q9, [x16, #0x40]
    76: 641c2a4e                 and v4.16b, v3.16b, v10.16b
    80: 65040c6f                 ushr v5.16b, v3.16b, #4
    84: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
    88: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
    92: c11ca64e                 mov v1.16b, v6.16b
    96: 211c276e                 eor v1.16b, v1.16b, v7.16b
   100: 0a0540f9                 ldr x10, [x8, #8]
   104: 4a010b8b                 add x10, x10, x11
   108: 4201c03d                 ldr q2, [x10]
   112: 001c226e                 eor v0.16b, v0.16b, v2.16b
   116: 0822c03d                 ldr q8, [x16, #0x80]
   120: 0932c03d                 ldr q9, [x16, #0xc0]
   124: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   128: 45040c6f                 ushr v5.16b, v2.16b, #4
   132: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   136: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   140: 211c266e                 eor v1.16b, v1.16b, v6.16b
   144: 211c276e                 eor v1.16b, v1.16b, v7.16b
   148: 0a0940f9                 ldr x10, [x8, #0x10]
   152: 4a010b8b                 add x10, x10, x11
   156: 4201c03d                 ldr q2, [x10]
   160: 001c226e                 eor v0.16b, v0.16b, v2.16b
   164: 0842c03d                 ldr q8, [x16, #0x100]
   168: 0952c03d                 ldr q9, [x16, #0x140]
   172: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   176: 45040c6f                 ushr v5.16b, v2.16b, #4
   180: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   184: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   188: 211c266e                 eor v1.16b, v1.16b, v6.16b
   192: 211c276e                 eor v1.16b, v1.16b, v7.16b
   196: 0a0d40f9                 ldr x10, [x8, #0x18]
   200: 4a010b8b                 add x10, x10, x11
   204: 4201c03d                 ldr q2, [x10]
   208: 001c226e                 eor v0.16b, v0.16b, v2.16b
   212: 0862c03d                 ldr q8, [x16, #0x180]
   216: 0972c03d                 ldr q9, [x16, #0x1c0]
   220: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   224: 45040c6f                 ushr v5.16b, v2.16b, #4
   228: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   232: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   236: 211c266e                 eor v1.16b, v1.16b, v6.16b
   240: 211c276e                 eor v1.16b, v1.16b, v7.16b
   244: 0a1140f9                 ldr x10, [x8, #0x20]
   248: 4a010b8b                 add x10, x10, x11
   252: 4201c03d                 ldr q2, [x10]
   256: 001c226e                 eor v0.16b, v0.16b, v2.16b
   260: 0882c03d                 ldr q8, [x16, #0x200]
   264: 0992c03d                 ldr q9, [x16, #0x240]
   268: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   272: 45040c6f                 ushr v5.16b, v2.16b, #4
   276: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   280: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   284: 211c266e                 eor v1.16b, v1.16b, v6.16b
   288: 211c276e                 eor v1.16b, v1.16b, v7.16b
   292: 0a1540f9                 ldr x10, [x8, #0x28]
   296: 4a010b8b                 add x10, x10, x11
   300: 4201c03d                 ldr q2, [x10]
   304: 001c226e                 eor v0.16b, v0.16b, v2.16b
   308: 08a2c03d                 ldr q8, [x16, #0x280]
   312: 09b2c03d                 ldr q9, [x16, #0x2c0]
   316: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   320: 45040c6f                 ushr v5.16b, v2.16b, #4
   324: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   328: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   332: 211c266e                 eor v1.16b, v1.16b, v6.16b
   336: 211c276e                 eor v1.16b, v1.16b, v7.16b
   340: 0a1940f9                 ldr x10, [x8, #0x30]
   344: 4a010b8b                 add x10, x10, x11
   348: 4201c03d                 ldr q2, [x10]
   352: 001c226e                 eor v0.16b, v0.16b, v2.16b
   356: 08c2c03d                 ldr q8, [x16, #0x300]
   360: 09d2c03d                 ldr q9, [x16, #0x340]
   364: 441c2a4e                 and v4.16b, v2.16b, v10.16b
   368: 45040c6f                 ushr v5.16b, v2.16b, #4
   372: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   376: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   380: 211c266e                 eor v1.16b, v1.16b, v6.16b
   384: 211c276e                 eor v1.16b, v1.16b, v7.16b
   388: 0a1d40f9                 ldr x10, [x8, #0x38]
   392: 4a010b8b                 add x10, x10, x11
   396: 4201c03d                 ldr q2, [x10]
   400: 001c226e                 eor v0.16b, v0.16b, v2.16b
   404: 211c226e                 eor v1.16b, v1.16b, v2.16b
   408: bf0500f1                 cmp x13, #1
   412: ad000054                 b.le 0x1b0
   416: 0a0140f9                 ldr x10, [x8]
   420: 4a010b8b                 add x10, x10, x11
   424: 4a410091                 add x10, x10, #0x10
   428: 4301c03d                 ldr q3, [x10]
   432: 2a0140f9                 ldr x10, [x9]
   436: 4a010b8b                 add x10, x10, x11
   440: 4001803d                 str q0, [x10]
   444: 2a0540f9                 ldr x10, [x9, #8]
   448: 4a010b8b                 add x10, x10, x11
   452: 4101803d                 str q1, [x10]
   456: 6b410091                 add x11, x11, #0x10
   460: ad0500f1                 subs x13, x13, #1
   464: 21f3ff54                 b.ne 0x34
   468: 7f0102eb                 cmp x11, x2
   472: 420e0054                 b.hs 0x3a0
   476: 0f020e91                 add x15, x16, #0x380
   480: 11021291                 add x17, x16, #0x480
   484: 2a0140f9                 ldr x10, [x9]
   488: ed030baa                 mov x13, x11
   492: 0c008052                 mov w12, #0
   496: 0e0140f9                 ldr x14, [x8]
   500: ce696d38                 ldrb w14, [x14, x13]
   504: 8c010e4a                 eor w12, w12, w14
   508: 0e0540f9                 ldr x14, [x8, #8]
   512: ce696d38                 ldrb w14, [x14, x13]
   516: 8c010e4a                 eor w12, w12, w14
   520: 0e0940f9                 ldr x14, [x8, #0x10]
   524: ce696d38                 ldrb w14, [x14, x13]
   528: 8c010e4a                 eor w12, w12, w14
   532: 0e0d40f9                 ldr x14, [x8, #0x18]
   536: ce696d38                 ldrb w14, [x14, x13]
   540: 8c010e4a                 eor w12, w12, w14
   544: 0e1140f9                 ldr x14, [x8, #0x20]
   548: ce696d38                 ldrb w14, [x14, x13]
   552: 8c010e4a                 eor w12, w12, w14
   556: 0e1540f9                 ldr x14, [x8, #0x28]
   560: ce696d38                 ldrb w14, [x14, x13]
   564: 8c010e4a                 eor w12, w12, w14
   568: 0e1940f9                 ldr x14, [x8, #0x30]
   572: ce696d38                 ldrb w14, [x14, x13]
   576: 8c010e4a                 eor w12, w12, w14
   580: 0e1d40f9                 ldr x14, [x8, #0x38]
   584: ce696d38                 ldrb w14, [x14, x13]
   588: 8c010e4a                 eor w12, w12, w14
   592: 4c692d38                 strb w12, [x10, x13]
   596: ad050091                 add x13, x13, #1
   600: bf0102eb                 cmp x13, x2
   604: 83fcff54                 b.lo 0x1ec
   608: 2a0540f9                 ldr x10, [x9, #8]
   612: ed030baa                 mov x13, x11
   616: 0c008052                 mov w12, #0
   620: 0e0140f9                 ldr x14, [x8]
   624: ce696d38                 ldrb w14, [x14, x13]
   628: 0e010034                 cbz w14, 0x294
   632: ee696e38                 ldrb w14, [x15, x14]
   636: ce990311                 add w14, w14, #0xe6
   640: dffd0371                 cmp w14, #0xff
   644: 43000054                 b.lo 0x28c
   648: cefd0351                 sub w14, w14, #0xff
   652: 2e6a6e38                 ldrb w14, [x17, x14]
   656: 8c010e4a                 eor w12, w12, w14
   660: 0e0540f9                 ldr x14, [x8, #8]
   664: ce696d38                 ldrb w14, [x14, x13]
   668: 0e010034                 cbz w14, 0x2bc
   672: ee696e38                 ldrb w14, [x15, x14]
   676: ce350311                 add w14, w14, #0xcd
   680: dffd0371                 cmp w14, #0xff
   684: 43000054                 b.lo 0x2b4
   688: cefd0351                 sub w14, w14, #0xff
   692: 2e6a6e38                 ldrb w14, [x17, x14]
   696: 8c010e4a                 eor w12, w12, w14
   700: 0e0940f9                 ldr x14, [x8, #0x10]
   704: ce696d38                 ldrb w14, [x14, x13]
   708: 0e010034                 cbz w14, 0x2e4
   712: ee696e38                 ldrb w14, [x15, x14]
   716: ce9d0311                 add w14, w14, #0xe7
   720: dffd0371                 cmp w14, #0xff
   724: 43000054                 b.lo 0x2dc
   728: cefd0351                 sub w14, w14, #0xff
   732: 2e6a6e38                 ldrb w14, [x17, x14]
   736: 8c010e4a                 eor w12, w12, w14
   740: 0e0d40f9                 ldr x14, [x8, #0x18]
   744: ce696d38                 ldrb w14, [x14, x13]
   748: 0e010034                 cbz w14, 0x30c
   752: ee696e38                 ldrb w14, [x15, x14]
   756: ce390311                 add w14, w14, #0xce
   760: dffd0371                 cmp w14, #0xff
   764: 43000054                 b.lo 0x304
   768: cefd0351                 sub w14, w14, #0xff
   772: 2e6a6e38                 ldrb w14, [x17, x14]
   776: 8c010e4a                 eor w12, w12, w14
   780: 0e1140f9                 ldr x14, [x8, #0x20]
   784: ce696d38                 ldrb w14, [x14, x13]
   788: 0e010034                 cbz w14, 0x334
   792: ee696e38                 ldrb w14, [x15, x14]
   796: ce2d0111                 add w14, w14, #0x4b
   800: dffd0371                 cmp w14, #0xff
   804: 43000054                 b.lo 0x32c
   808: cefd0351                 sub w14, w14, #0xff
   812: 2e6a6e38                 ldrb w14, [x17, x14]
   816: 8c010e4a                 eor w12, w12, w14
   820: 0e1540f9                 ldr x14, [x8, #0x28]
   824: ce696d38                 ldrb w14, [x14, x13]
   828: 0e010034                 cbz w14, 0x35c
   832: ee696e38                 ldrb w14, [x15, x14]
   836: cec90011                 add w14, w14, #0x32
   840: dffd0371                 cmp w14, #0xff
   844: 43000054                 b.lo 0x354
   848: cefd0351                 sub w14, w14, #0xff
   852: 2e6a6e38                 ldrb w14, [x17, x14]
   856: 8c010e4a                 eor w12, w12, w14
   860: 0e1940f9                 ldr x14, [x8, #0x30]
   864: ce696d38                 ldrb w14, [x14, x13]
   868: 0e010034                 cbz w14, 0x384
   872: ee696e38                 ldrb w14, [x15, x14]
   876: ce650011                 add w14, w14, #0x19
   880: dffd0371                 cmp w14, #0xff
   884: 43000054                 b.lo 0x37c
   888: cefd0351                 sub w14, w14, #0xff
   892: 2e6a6e38                 ldrb w14, [x17, x14]
   896: 8c010e4a                 eor w12, w12, w14
   900: 0e1d40f9                 ldr x14, [x8, #0x38]
   904: ce696d38                 ldrb w14, [x14, x13]
   908: 8c010e4a                 eor w12, w12, w14
   912: 4c692d38                 strb w12, [x10, x13]
   916: ad050091                 add x13, x13, #1
   920: bf0102eb                 cmp x13, x2
   924: 63f6ff54                 b.lo 0x268
   928: ea2fc1ac                 ldp q10, q11, [sp], #0x20
   932: e827c1ac                 ldp q8, q9, [sp], #0x20
   936: c0035fd6                 ret 
   940: 00000000                 udf #0
   944: 00000000                 udf #0
   948: 00000000                 udf #0
   952: 00000000                 udf #0
   956: 00000000                 udf #0
