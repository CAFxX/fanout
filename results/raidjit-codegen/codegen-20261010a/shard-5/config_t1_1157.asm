; config_t1_1157 target=neon twin=baseline topo=raid5(4)/enc op=encode U=8
; code_size=448 twins=[] uops_per_byte=0.5
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: e80300aa                 mov x8, x0
     8: e90301aa                 mov x9, x1
    12: eb031faa                 mov x11, xzr
    16: 0e1080d2                 mov x14, #0x80
    20: 4d08ce9a                 udiv x13, x2, x14
    24: ed0900b4                 cbz x13, 0x154
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
   288: 2a0140f9                 ldr x10, [x9]
   292: 4a010b8b                 add x10, x10, x11
   296: 4001803d                 str q0, [x10]
   300: 4105803d                 str q1, [x10, #0x10]
   304: 4209803d                 str q2, [x10, #0x20]
   308: 430d803d                 str q3, [x10, #0x30]
   312: 4411803d                 str q4, [x10, #0x40]
   316: 4515803d                 str q5, [x10, #0x50]
   320: 4619803d                 str q6, [x10, #0x60]
   324: 471d803d                 str q7, [x10, #0x70]
   328: 6b010291                 add x11, x11, #0x80
   332: ad0500f1                 subs x13, x13, #1
   336: 61f6ff54                 b.ne 0x1c
   340: 7f0102eb                 cmp x11, x2
   344: 62020054                 b.hs 0x1a4
   348: 2a0140f9                 ldr x10, [x9]
   352: ed030baa                 mov x13, x11
   356: 0e0140f9                 ldr x14, [x8]
   360: ce696d38                 ldrb w14, [x14, x13]
   364: ec030e2a                 mov w12, w14
   368: 0e0540f9                 ldr x14, [x8, #8]
   372: ce696d38                 ldrb w14, [x14, x13]
   376: 8c010e4a                 eor w12, w12, w14
   380: 0e0940f9                 ldr x14, [x8, #0x10]
   384: ce696d38                 ldrb w14, [x14, x13]
   388: 8c010e4a                 eor w12, w12, w14
   392: 0e0d40f9                 ldr x14, [x8, #0x18]
   396: ce696d38                 ldrb w14, [x14, x13]
   400: 8c010e4a                 eor w12, w12, w14
   404: 4c692d38                 strb w12, [x10, x13]
   408: ad050091                 add x13, x13, #1
   412: bf0102eb                 cmp x13, x2
   416: 23feff54                 b.lo 0x164
   420: e827c1ac                 ldp q8, q9, [sp], #0x20
   424: c0035fd6                 ret 
   428: 00000000                 udf #0
   432: 00000000                 udf #0
   436: 00000000                 udf #0
   440: 00000000                 udf #0
   444: 00000000                 udf #0
