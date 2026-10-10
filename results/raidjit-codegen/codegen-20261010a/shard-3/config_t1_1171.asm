; config_t1_1171 target=neon twin=nt topo=raid6(8)/enc op=encode U=1
; code_size=960 twins=["nt"] uops_per_byte=5.25
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: e80300aa                 mov x8, x0
     8: e90301aa                 mov x9, x1
    12: b01d0010                 adr x16, 0x3c0
    16: eb031faa                 mov x11, xzr
    20: 0e0280d2                 mov x14, #0x10
    24: 4d08ce9a                 udiv x13, x2, x14
    28: 6d0e00b4                 cbz x13, 0x1e8
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
    80: 0702c03d                 ldr q7, [x16]
    84: 0812c03d                 ldr q8, [x16, #0x40]
    88: 4201c03d                 ldr q2, [x10]
    92: 431c294e                 and v3.16b, v2.16b, v9.16b
    96: 44040c6f                 ushr v4.16b, v2.16b, #4
   100: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   104: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   108: a11ca54e                 mov v1.16b, v5.16b
   112: 211c266e                 eor v1.16b, v1.16b, v6.16b
   116: 0a0540f9                 ldr x10, [x8, #8]
   120: 4a010b8b                 add x10, x10, x11
   124: 4201c03d                 ldr q2, [x10]
   128: 001c226e                 eor v0.16b, v0.16b, v2.16b
   132: 0722c03d                 ldr q7, [x16, #0x80]
   136: 0832c03d                 ldr q8, [x16, #0xc0]
   140: 4201c03d                 ldr q2, [x10]
   144: 431c294e                 and v3.16b, v2.16b, v9.16b
   148: 44040c6f                 ushr v4.16b, v2.16b, #4
   152: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   156: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   160: 211c256e                 eor v1.16b, v1.16b, v5.16b
   164: 211c266e                 eor v1.16b, v1.16b, v6.16b
   168: 0a0940f9                 ldr x10, [x8, #0x10]
   172: 4a010b8b                 add x10, x10, x11
   176: 4201c03d                 ldr q2, [x10]
   180: 001c226e                 eor v0.16b, v0.16b, v2.16b
   184: 0742c03d                 ldr q7, [x16, #0x100]
   188: 0852c03d                 ldr q8, [x16, #0x140]
   192: 4201c03d                 ldr q2, [x10]
   196: 431c294e                 and v3.16b, v2.16b, v9.16b
   200: 44040c6f                 ushr v4.16b, v2.16b, #4
   204: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   208: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   212: 211c256e                 eor v1.16b, v1.16b, v5.16b
   216: 211c266e                 eor v1.16b, v1.16b, v6.16b
   220: 0a0d40f9                 ldr x10, [x8, #0x18]
   224: 4a010b8b                 add x10, x10, x11
   228: 4201c03d                 ldr q2, [x10]
   232: 001c226e                 eor v0.16b, v0.16b, v2.16b
   236: 0762c03d                 ldr q7, [x16, #0x180]
   240: 0872c03d                 ldr q8, [x16, #0x1c0]
   244: 4201c03d                 ldr q2, [x10]
   248: 431c294e                 and v3.16b, v2.16b, v9.16b
   252: 44040c6f                 ushr v4.16b, v2.16b, #4
   256: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   260: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   264: 211c256e                 eor v1.16b, v1.16b, v5.16b
   268: 211c266e                 eor v1.16b, v1.16b, v6.16b
   272: 0a1140f9                 ldr x10, [x8, #0x20]
   276: 4a010b8b                 add x10, x10, x11
   280: 4201c03d                 ldr q2, [x10]
   284: 001c226e                 eor v0.16b, v0.16b, v2.16b
   288: 0782c03d                 ldr q7, [x16, #0x200]
   292: 0892c03d                 ldr q8, [x16, #0x240]
   296: 4201c03d                 ldr q2, [x10]
   300: 431c294e                 and v3.16b, v2.16b, v9.16b
   304: 44040c6f                 ushr v4.16b, v2.16b, #4
   308: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   312: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   316: 211c256e                 eor v1.16b, v1.16b, v5.16b
   320: 211c266e                 eor v1.16b, v1.16b, v6.16b
   324: 0a1540f9                 ldr x10, [x8, #0x28]
   328: 4a010b8b                 add x10, x10, x11
   332: 4201c03d                 ldr q2, [x10]
   336: 001c226e                 eor v0.16b, v0.16b, v2.16b
   340: 07a2c03d                 ldr q7, [x16, #0x280]
   344: 08b2c03d                 ldr q8, [x16, #0x2c0]
   348: 4201c03d                 ldr q2, [x10]
   352: 431c294e                 and v3.16b, v2.16b, v9.16b
   356: 44040c6f                 ushr v4.16b, v2.16b, #4
   360: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   364: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   368: 211c256e                 eor v1.16b, v1.16b, v5.16b
   372: 211c266e                 eor v1.16b, v1.16b, v6.16b
   376: 0a1940f9                 ldr x10, [x8, #0x30]
   380: 4a010b8b                 add x10, x10, x11
   384: 4201c03d                 ldr q2, [x10]
   388: 001c226e                 eor v0.16b, v0.16b, v2.16b
   392: 07c2c03d                 ldr q7, [x16, #0x300]
   396: 08d2c03d                 ldr q8, [x16, #0x340]
   400: 4201c03d                 ldr q2, [x10]
   404: 431c294e                 and v3.16b, v2.16b, v9.16b
   408: 44040c6f                 ushr v4.16b, v2.16b, #4
   412: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   416: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   420: 211c256e                 eor v1.16b, v1.16b, v5.16b
   424: 211c266e                 eor v1.16b, v1.16b, v6.16b
   428: 0a1d40f9                 ldr x10, [x8, #0x38]
   432: 4a010b8b                 add x10, x10, x11
   436: 4201c03d                 ldr q2, [x10]
   440: 001c226e                 eor v0.16b, v0.16b, v2.16b
   444: 4201c03d                 ldr q2, [x10]
   448: 211c226e                 eor v1.16b, v1.16b, v2.16b
   452: 2a0140f9                 ldr x10, [x9]
   456: 4a010b8b                 add x10, x10, x11
   460: 4001803d                 str q0, [x10]
   464: 2a0540f9                 ldr x10, [x9, #8]
   468: 4a010b8b                 add x10, x10, x11
   472: 4101803d                 str q1, [x10]
   476: 6b410091                 add x11, x11, #0x10
   480: ad0500f1                 subs x13, x13, #1
   484: 01f2ff54                 b.ne 0x24
   488: 7f0102eb                 cmp x11, x2
   492: 420e0054                 b.hs 0x3b4
   496: 0f020e91                 add x15, x16, #0x380
   500: 11021291                 add x17, x16, #0x480
   504: 2a0140f9                 ldr x10, [x9]
   508: ed030baa                 mov x13, x11
   512: 0c008052                 mov w12, #0
   516: 0e0140f9                 ldr x14, [x8]
   520: ce696d38                 ldrb w14, [x14, x13]
   524: 8c010e4a                 eor w12, w12, w14
   528: 0e0540f9                 ldr x14, [x8, #8]
   532: ce696d38                 ldrb w14, [x14, x13]
   536: 8c010e4a                 eor w12, w12, w14
   540: 0e0940f9                 ldr x14, [x8, #0x10]
   544: ce696d38                 ldrb w14, [x14, x13]
   548: 8c010e4a                 eor w12, w12, w14
   552: 0e0d40f9                 ldr x14, [x8, #0x18]
   556: ce696d38                 ldrb w14, [x14, x13]
   560: 8c010e4a                 eor w12, w12, w14
   564: 0e1140f9                 ldr x14, [x8, #0x20]
   568: ce696d38                 ldrb w14, [x14, x13]
   572: 8c010e4a                 eor w12, w12, w14
   576: 0e1540f9                 ldr x14, [x8, #0x28]
   580: ce696d38                 ldrb w14, [x14, x13]
   584: 8c010e4a                 eor w12, w12, w14
   588: 0e1940f9                 ldr x14, [x8, #0x30]
   592: ce696d38                 ldrb w14, [x14, x13]
   596: 8c010e4a                 eor w12, w12, w14
   600: 0e1d40f9                 ldr x14, [x8, #0x38]
   604: ce696d38                 ldrb w14, [x14, x13]
   608: 8c010e4a                 eor w12, w12, w14
   612: 4c692d38                 strb w12, [x10, x13]
   616: ad050091                 add x13, x13, #1
   620: bf0102eb                 cmp x13, x2
   624: 83fcff54                 b.lo 0x200
   628: 2a0540f9                 ldr x10, [x9, #8]
   632: ed030baa                 mov x13, x11
   636: 0c008052                 mov w12, #0
   640: 0e0140f9                 ldr x14, [x8]
   644: ce696d38                 ldrb w14, [x14, x13]
   648: 0e010034                 cbz w14, 0x2a8
   652: ee696e38                 ldrb w14, [x15, x14]
   656: ce990311                 add w14, w14, #0xe6
   660: dffd0371                 cmp w14, #0xff
   664: 43000054                 b.lo 0x2a0
   668: cefd0351                 sub w14, w14, #0xff
   672: 2e6a6e38                 ldrb w14, [x17, x14]
   676: 8c010e4a                 eor w12, w12, w14
   680: 0e0540f9                 ldr x14, [x8, #8]
   684: ce696d38                 ldrb w14, [x14, x13]
   688: 0e010034                 cbz w14, 0x2d0
   692: ee696e38                 ldrb w14, [x15, x14]
   696: ce350311                 add w14, w14, #0xcd
   700: dffd0371                 cmp w14, #0xff
   704: 43000054                 b.lo 0x2c8
   708: cefd0351                 sub w14, w14, #0xff
   712: 2e6a6e38                 ldrb w14, [x17, x14]
   716: 8c010e4a                 eor w12, w12, w14
   720: 0e0940f9                 ldr x14, [x8, #0x10]
   724: ce696d38                 ldrb w14, [x14, x13]
   728: 0e010034                 cbz w14, 0x2f8
   732: ee696e38                 ldrb w14, [x15, x14]
   736: ce9d0311                 add w14, w14, #0xe7
   740: dffd0371                 cmp w14, #0xff
   744: 43000054                 b.lo 0x2f0
   748: cefd0351                 sub w14, w14, #0xff
   752: 2e6a6e38                 ldrb w14, [x17, x14]
   756: 8c010e4a                 eor w12, w12, w14
   760: 0e0d40f9                 ldr x14, [x8, #0x18]
   764: ce696d38                 ldrb w14, [x14, x13]
   768: 0e010034                 cbz w14, 0x320
   772: ee696e38                 ldrb w14, [x15, x14]
   776: ce390311                 add w14, w14, #0xce
   780: dffd0371                 cmp w14, #0xff
   784: 43000054                 b.lo 0x318
   788: cefd0351                 sub w14, w14, #0xff
   792: 2e6a6e38                 ldrb w14, [x17, x14]
   796: 8c010e4a                 eor w12, w12, w14
   800: 0e1140f9                 ldr x14, [x8, #0x20]
   804: ce696d38                 ldrb w14, [x14, x13]
   808: 0e010034                 cbz w14, 0x348
   812: ee696e38                 ldrb w14, [x15, x14]
   816: ce2d0111                 add w14, w14, #0x4b
   820: dffd0371                 cmp w14, #0xff
   824: 43000054                 b.lo 0x340
   828: cefd0351                 sub w14, w14, #0xff
   832: 2e6a6e38                 ldrb w14, [x17, x14]
   836: 8c010e4a                 eor w12, w12, w14
   840: 0e1540f9                 ldr x14, [x8, #0x28]
   844: ce696d38                 ldrb w14, [x14, x13]
   848: 0e010034                 cbz w14, 0x370
   852: ee696e38                 ldrb w14, [x15, x14]
   856: cec90011                 add w14, w14, #0x32
   860: dffd0371                 cmp w14, #0xff
   864: 43000054                 b.lo 0x368
   868: cefd0351                 sub w14, w14, #0xff
   872: 2e6a6e38                 ldrb w14, [x17, x14]
   876: 8c010e4a                 eor w12, w12, w14
   880: 0e1940f9                 ldr x14, [x8, #0x30]
   884: ce696d38                 ldrb w14, [x14, x13]
   888: 0e010034                 cbz w14, 0x398
   892: ee696e38                 ldrb w14, [x15, x14]
   896: ce650011                 add w14, w14, #0x19
   900: dffd0371                 cmp w14, #0xff
   904: 43000054                 b.lo 0x390
   908: cefd0351                 sub w14, w14, #0xff
   912: 2e6a6e38                 ldrb w14, [x17, x14]
   916: 8c010e4a                 eor w12, w12, w14
   920: 0e1d40f9                 ldr x14, [x8, #0x38]
   924: ce696d38                 ldrb w14, [x14, x13]
   928: 8c010e4a                 eor w12, w12, w14
   932: 4c692d38                 strb w12, [x10, x13]
   936: ad050091                 add x13, x13, #1
   940: bf0102eb                 cmp x13, x2
   944: 63f6ff54                 b.lo 0x27c
   948: e827c1ac                 ldp q8, q9, [sp], #0x20
   952: c0035fd6                 ret 
   956: 00000000                 udf #0
