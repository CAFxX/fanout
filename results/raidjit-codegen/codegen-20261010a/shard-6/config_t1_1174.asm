; config_t1_1174 target=neon twin=notail topo=raid6(8)/enc op=encode U=1
; code_size=512 twins=["notail"] uops_per_byte=5.25
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: e80300aa                 mov x8, x0
     8: e90301aa                 mov x9, x1
    12: b00f0010                 adr x16, 0x200
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
   488: e827c1ac                 ldp q8, q9, [sp], #0x20
   492: c0035fd6                 ret 
   496: 00000000                 udf #0
   500: 00000000                 udf #0
   504: 00000000                 udf #0
   508: 00000000                 udf #0
