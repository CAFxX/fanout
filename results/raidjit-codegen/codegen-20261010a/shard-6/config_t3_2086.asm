; config_t3_2086 target=neon twin=baseline topo=rs(3,1)/enc op=encode U=8
; code_size=384 twins=[] uops_per_byte=0.375
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: e80300aa                 mov x8, x0
     8: e90301aa                 mov x9, x1
    12: eb031faa                 mov x11, xzr
    16: 0e1080d2                 mov x14, #0x80
    20: 4d08ce9a                 udiv x13, x2, x14
    24: ad0700b4                 cbz x13, 0x10c
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
   216: 2a0140f9                 ldr x10, [x9]
   220: 4a010b8b                 add x10, x10, x11
   224: 4001803d                 str q0, [x10]
   228: 4105803d                 str q1, [x10, #0x10]
   232: 4209803d                 str q2, [x10, #0x20]
   236: 430d803d                 str q3, [x10, #0x30]
   240: 4411803d                 str q4, [x10, #0x40]
   244: 4515803d                 str q5, [x10, #0x50]
   248: 4619803d                 str q6, [x10, #0x60]
   252: 471d803d                 str q7, [x10, #0x70]
   256: 6b010291                 add x11, x11, #0x80
   260: ad0500f1                 subs x13, x13, #1
   264: a1f8ff54                 b.ne 0x1c
   268: 7f0102eb                 cmp x11, x2
   272: 02020054                 b.hs 0x150
   276: 2a0140f9                 ldr x10, [x9]
   280: ed030baa                 mov x13, x11
   284: 0e0140f9                 ldr x14, [x8]
   288: ce696d38                 ldrb w14, [x14, x13]
   292: ec030e2a                 mov w12, w14
   296: 0e0540f9                 ldr x14, [x8, #8]
   300: ce696d38                 ldrb w14, [x14, x13]
   304: 8c010e4a                 eor w12, w12, w14
   308: 0e0940f9                 ldr x14, [x8, #0x10]
   312: ce696d38                 ldrb w14, [x14, x13]
   316: 8c010e4a                 eor w12, w12, w14
   320: 4c692d38                 strb w12, [x10, x13]
   324: ad050091                 add x13, x13, #1
   328: bf0102eb                 cmp x13, x2
   332: 83feff54                 b.lo 0x11c
   336: e827c1ac                 ldp q8, q9, [sp], #0x20
   340: c0035fd6                 ret 
   344: 00000000                 udf #0
   348: 00000000                 udf #0
   352: 00000000                 udf #0
   356: 00000000                 udf #0
   360: 00000000                 udf #0
   364: 00000000                 udf #0
   368: 00000000                 udf #0
   372: 00000000                 udf #0
   376: 00000000                 udf #0
   380: 00000000                 udf #0
