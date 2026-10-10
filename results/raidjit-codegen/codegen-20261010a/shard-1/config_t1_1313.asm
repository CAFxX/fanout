; config_t1_1313 target=neon twin=baseline topo=rs(6,3)/rec{0} op=recover U=8
; code_size=640 twins=[] uops_per_byte=0.75
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: e80300aa                 mov x8, x0
     8: e90301aa                 mov x9, x1
    12: eb031faa                 mov x11, xzr
    16: 0e1080d2                 mov x14, #0x80
    20: 4d08ce9a                 udiv x13, x2, x14
    24: 6d0e00b4                 cbz x13, 0x1e4
    28: 1f2003d5                 nop 
    32: 0a0140f9                 ldr x10, [x8]
    36: 4a010b8b                 add x10, x10, x11
    40: 4001c03d                 ldr q0, [x10]
    44: 4105c03d                 ldr q1, [x10, #0x10]
    48: 4209c03d                 ldr q2, [x10, #0x20]
    52: 430dc03d                 ldr q3, [x10, #0x30]
    56: 4411c03d                 ldr q4, [x10, #0x40]
    60: 4515c03d                 ldr q5, [x10, #0x50]
    64: 4619c03d                 ldr q6, [x10, #0x60]
    68: 471dc03d                 ldr q7, [x10, #0x70]
    72: 0a0540f9                 ldr x10, [x8, #8]
    76: 4a010b8b                 add x10, x10, x11
    80: 4801c03d                 ldr q8, [x10]
    84: 001c286e                 eor v0.16b, v0.16b, v8.16b
    88: 4805c03d                 ldr q8, [x10, #0x10]
    92: 211c286e                 eor v1.16b, v1.16b, v8.16b
    96: 4809c03d                 ldr q8, [x10, #0x20]
   100: 421c286e                 eor v2.16b, v2.16b, v8.16b
   104: 480dc03d                 ldr q8, [x10, #0x30]
   108: 631c286e                 eor v3.16b, v3.16b, v8.16b
   112: 4811c03d                 ldr q8, [x10, #0x40]
   116: 841c286e                 eor v4.16b, v4.16b, v8.16b
   120: 4815c03d                 ldr q8, [x10, #0x50]
   124: a51c286e                 eor v5.16b, v5.16b, v8.16b
   128: 4819c03d                 ldr q8, [x10, #0x60]
   132: c61c286e                 eor v6.16b, v6.16b, v8.16b
   136: 481dc03d                 ldr q8, [x10, #0x70]
   140: e71c286e                 eor v7.16b, v7.16b, v8.16b
   144: 0a0940f9                 ldr x10, [x8, #0x10]
   148: 4a010b8b                 add x10, x10, x11
   152: 4801c03d                 ldr q8, [x10]
   156: 001c286e                 eor v0.16b, v0.16b, v8.16b
   160: 4805c03d                 ldr q8, [x10, #0x10]
   164: 211c286e                 eor v1.16b, v1.16b, v8.16b
   168: 4809c03d                 ldr q8, [x10, #0x20]
   172: 421c286e                 eor v2.16b, v2.16b, v8.16b
   176: 480dc03d                 ldr q8, [x10, #0x30]
   180: 631c286e                 eor v3.16b, v3.16b, v8.16b
   184: 4811c03d                 ldr q8, [x10, #0x40]
   188: 841c286e                 eor v4.16b, v4.16b, v8.16b
   192: 4815c03d                 ldr q8, [x10, #0x50]
   196: a51c286e                 eor v5.16b, v5.16b, v8.16b
   200: 4819c03d                 ldr q8, [x10, #0x60]
   204: c61c286e                 eor v6.16b, v6.16b, v8.16b
   208: 481dc03d                 ldr q8, [x10, #0x70]
   212: e71c286e                 eor v7.16b, v7.16b, v8.16b
   216: 0a0d40f9                 ldr x10, [x8, #0x18]
   220: 4a010b8b                 add x10, x10, x11
   224: 4801c03d                 ldr q8, [x10]
   228: 001c286e                 eor v0.16b, v0.16b, v8.16b
   232: 4805c03d                 ldr q8, [x10, #0x10]
   236: 211c286e                 eor v1.16b, v1.16b, v8.16b
   240: 4809c03d                 ldr q8, [x10, #0x20]
   244: 421c286e                 eor v2.16b, v2.16b, v8.16b
   248: 480dc03d                 ldr q8, [x10, #0x30]
   252: 631c286e                 eor v3.16b, v3.16b, v8.16b
   256: 4811c03d                 ldr q8, [x10, #0x40]
   260: 841c286e                 eor v4.16b, v4.16b, v8.16b
   264: 4815c03d                 ldr q8, [x10, #0x50]
   268: a51c286e                 eor v5.16b, v5.16b, v8.16b
   272: 4819c03d                 ldr q8, [x10, #0x60]
   276: c61c286e                 eor v6.16b, v6.16b, v8.16b
   280: 481dc03d                 ldr q8, [x10, #0x70]
   284: e71c286e                 eor v7.16b, v7.16b, v8.16b
   288: 0a1140f9                 ldr x10, [x8, #0x20]
   292: 4a010b8b                 add x10, x10, x11
   296: 4801c03d                 ldr q8, [x10]
   300: 001c286e                 eor v0.16b, v0.16b, v8.16b
   304: 4805c03d                 ldr q8, [x10, #0x10]
   308: 211c286e                 eor v1.16b, v1.16b, v8.16b
   312: 4809c03d                 ldr q8, [x10, #0x20]
   316: 421c286e                 eor v2.16b, v2.16b, v8.16b
   320: 480dc03d                 ldr q8, [x10, #0x30]
   324: 631c286e                 eor v3.16b, v3.16b, v8.16b
   328: 4811c03d                 ldr q8, [x10, #0x40]
   332: 841c286e                 eor v4.16b, v4.16b, v8.16b
   336: 4815c03d                 ldr q8, [x10, #0x50]
   340: a51c286e                 eor v5.16b, v5.16b, v8.16b
   344: 4819c03d                 ldr q8, [x10, #0x60]
   348: c61c286e                 eor v6.16b, v6.16b, v8.16b
   352: 481dc03d                 ldr q8, [x10, #0x70]
   356: e71c286e                 eor v7.16b, v7.16b, v8.16b
   360: 0a1540f9                 ldr x10, [x8, #0x28]
   364: 4a010b8b                 add x10, x10, x11
   368: 4801c03d                 ldr q8, [x10]
   372: 001c286e                 eor v0.16b, v0.16b, v8.16b
   376: 4805c03d                 ldr q8, [x10, #0x10]
   380: 211c286e                 eor v1.16b, v1.16b, v8.16b
   384: 4809c03d                 ldr q8, [x10, #0x20]
   388: 421c286e                 eor v2.16b, v2.16b, v8.16b
   392: 480dc03d                 ldr q8, [x10, #0x30]
   396: 631c286e                 eor v3.16b, v3.16b, v8.16b
   400: 4811c03d                 ldr q8, [x10, #0x40]
   404: 841c286e                 eor v4.16b, v4.16b, v8.16b
   408: 4815c03d                 ldr q8, [x10, #0x50]
   412: a51c286e                 eor v5.16b, v5.16b, v8.16b
   416: 4819c03d                 ldr q8, [x10, #0x60]
   420: c61c286e                 eor v6.16b, v6.16b, v8.16b
   424: 481dc03d                 ldr q8, [x10, #0x70]
   428: e71c286e                 eor v7.16b, v7.16b, v8.16b
   432: 2a0140f9                 ldr x10, [x9]
   436: 4a010b8b                 add x10, x10, x11
   440: 4001803d                 str q0, [x10]
   444: 4105803d                 str q1, [x10, #0x10]
   448: 4209803d                 str q2, [x10, #0x20]
   452: 430d803d                 str q3, [x10, #0x30]
   456: 4411803d                 str q4, [x10, #0x40]
   460: 4515803d                 str q5, [x10, #0x50]
   464: 4619803d                 str q6, [x10, #0x60]
   468: 471d803d                 str q7, [x10, #0x70]
   472: 6b010291                 add x11, x11, #0x80
   476: ad0500f1                 subs x13, x13, #1
   480: e1f1ff54                 b.ne 0x1c
   484: 7f0102eb                 cmp x11, x2
   488: 22030054                 b.hs 0x24c
   492: 2a0140f9                 ldr x10, [x9]
   496: ed030baa                 mov x13, x11
   500: 0e0140f9                 ldr x14, [x8]
   504: ce696d38                 ldrb w14, [x14, x13]
   508: ec030e2a                 mov w12, w14
   512: 0e0540f9                 ldr x14, [x8, #8]
   516: ce696d38                 ldrb w14, [x14, x13]
   520: 8c010e4a                 eor w12, w12, w14
   524: 0e0940f9                 ldr x14, [x8, #0x10]
   528: ce696d38                 ldrb w14, [x14, x13]
   532: 8c010e4a                 eor w12, w12, w14
   536: 0e0d40f9                 ldr x14, [x8, #0x18]
   540: ce696d38                 ldrb w14, [x14, x13]
   544: 8c010e4a                 eor w12, w12, w14
   548: 0e1140f9                 ldr x14, [x8, #0x20]
   552: ce696d38                 ldrb w14, [x14, x13]
   556: 8c010e4a                 eor w12, w12, w14
   560: 0e1540f9                 ldr x14, [x8, #0x28]
   564: ce696d38                 ldrb w14, [x14, x13]
   568: 8c010e4a                 eor w12, w12, w14
   572: 4c692d38                 strb w12, [x10, x13]
   576: ad050091                 add x13, x13, #1
   580: bf0102eb                 cmp x13, x2
   584: 63fdff54                 b.lo 0x1f4
   588: e827c1ac                 ldp q8, q9, [sp], #0x20
   592: c0035fd6                 ret 
   596: 00000000                 udf #0
   600: 00000000                 udf #0
   604: 00000000                 udf #0
   608: 00000000                 udf #0
   612: 00000000                 udf #0
   616: 00000000                 udf #0
   620: 00000000                 udf #0
   624: 00000000                 udf #0
   628: 00000000                 udf #0
   632: 00000000                 udf #0
   636: 00000000                 udf #0
