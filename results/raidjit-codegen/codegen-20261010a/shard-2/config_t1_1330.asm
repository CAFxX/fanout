; config_t1_1330 target=neon twin=notail topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=640 twins=["notail"] uops_per_byte=6.9375
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: e80300aa                 mov x8, x0
     8: e90301aa                 mov x9, x1
    12: b0130010                 adr x16, 0x280
    16: eb031faa                 mov x11, xzr
    20: 0e0280d2                 mov x14, #0x10
    24: 4d08ce9a                 udiv x13, x2, x14
    28: 4d1100b4                 cbz x13, 0x244
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
    72: 0702c03d                 ldr q7, [x16]
    76: 0812c03d                 ldr q8, [x16, #0x40]
    80: 4201c03d                 ldr q2, [x10]
    84: 431c294e                 and v3.16b, v2.16b, v9.16b
    88: 44040c6f                 ushr v4.16b, v2.16b, #4
    92: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
    96: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   100: a01ca54e                 mov v0.16b, v5.16b
   104: 001c266e                 eor v0.16b, v0.16b, v6.16b
   108: 0722c03d                 ldr q7, [x16, #0x80]
   112: 0832c03d                 ldr q8, [x16, #0xc0]
   116: 4201c03d                 ldr q2, [x10]
   120: 431c294e                 and v3.16b, v2.16b, v9.16b
   124: 44040c6f                 ushr v4.16b, v2.16b, #4
   128: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   132: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   136: a11ca54e                 mov v1.16b, v5.16b
   140: 211c266e                 eor v1.16b, v1.16b, v6.16b
   144: 0a0540f9                 ldr x10, [x8, #8]
   148: 4a010b8b                 add x10, x10, x11
   152: 0742c03d                 ldr q7, [x16, #0x100]
   156: 0852c03d                 ldr q8, [x16, #0x140]
   160: 4201c03d                 ldr q2, [x10]
   164: 431c294e                 and v3.16b, v2.16b, v9.16b
   168: 44040c6f                 ushr v4.16b, v2.16b, #4
   172: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   176: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   180: 001c256e                 eor v0.16b, v0.16b, v5.16b
   184: 001c266e                 eor v0.16b, v0.16b, v6.16b
   188: 0762c03d                 ldr q7, [x16, #0x180]
   192: 0872c03d                 ldr q8, [x16, #0x1c0]
   196: 4201c03d                 ldr q2, [x10]
   200: 431c294e                 and v3.16b, v2.16b, v9.16b
   204: 44040c6f                 ushr v4.16b, v2.16b, #4
   208: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   212: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   216: 211c256e                 eor v1.16b, v1.16b, v5.16b
   220: 211c266e                 eor v1.16b, v1.16b, v6.16b
   224: 0a0940f9                 ldr x10, [x8, #0x10]
   228: 4a010b8b                 add x10, x10, x11
   232: 0782c03d                 ldr q7, [x16, #0x200]
   236: 0892c03d                 ldr q8, [x16, #0x240]
   240: 4201c03d                 ldr q2, [x10]
   244: 431c294e                 and v3.16b, v2.16b, v9.16b
   248: 44040c6f                 ushr v4.16b, v2.16b, #4
   252: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   256: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   260: 001c256e                 eor v0.16b, v0.16b, v5.16b
   264: 001c266e                 eor v0.16b, v0.16b, v6.16b
   268: 07a2c03d                 ldr q7, [x16, #0x280]
   272: 08b2c03d                 ldr q8, [x16, #0x2c0]
   276: 4201c03d                 ldr q2, [x10]
   280: 431c294e                 and v3.16b, v2.16b, v9.16b
   284: 44040c6f                 ushr v4.16b, v2.16b, #4
   288: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   292: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   296: 211c256e                 eor v1.16b, v1.16b, v5.16b
   300: 211c266e                 eor v1.16b, v1.16b, v6.16b
   304: 0a0d40f9                 ldr x10, [x8, #0x18]
   308: 4a010b8b                 add x10, x10, x11
   312: 07c2c03d                 ldr q7, [x16, #0x300]
   316: 08d2c03d                 ldr q8, [x16, #0x340]
   320: 4201c03d                 ldr q2, [x10]
   324: 431c294e                 and v3.16b, v2.16b, v9.16b
   328: 44040c6f                 ushr v4.16b, v2.16b, #4
   332: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   336: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   340: 001c256e                 eor v0.16b, v0.16b, v5.16b
   344: 001c266e                 eor v0.16b, v0.16b, v6.16b
   348: 07e2c03d                 ldr q7, [x16, #0x380]
   352: 08f2c03d                 ldr q8, [x16, #0x3c0]
   356: 4201c03d                 ldr q2, [x10]
   360: 431c294e                 and v3.16b, v2.16b, v9.16b
   364: 44040c6f                 ushr v4.16b, v2.16b, #4
   368: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   372: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   376: 211c256e                 eor v1.16b, v1.16b, v5.16b
   380: 211c266e                 eor v1.16b, v1.16b, v6.16b
   384: 0a1140f9                 ldr x10, [x8, #0x20]
   388: 4a010b8b                 add x10, x10, x11
   392: 0702c13d                 ldr q7, [x16, #0x400]
   396: 0812c13d                 ldr q8, [x16, #0x440]
   400: 4201c03d                 ldr q2, [x10]
   404: 431c294e                 and v3.16b, v2.16b, v9.16b
   408: 44040c6f                 ushr v4.16b, v2.16b, #4
   412: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   416: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   420: 001c256e                 eor v0.16b, v0.16b, v5.16b
   424: 001c266e                 eor v0.16b, v0.16b, v6.16b
   428: 0722c13d                 ldr q7, [x16, #0x480]
   432: 0832c13d                 ldr q8, [x16, #0x4c0]
   436: 4201c03d                 ldr q2, [x10]
   440: 431c294e                 and v3.16b, v2.16b, v9.16b
   444: 44040c6f                 ushr v4.16b, v2.16b, #4
   448: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   452: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   456: 211c256e                 eor v1.16b, v1.16b, v5.16b
   460: 211c266e                 eor v1.16b, v1.16b, v6.16b
   464: 0a1540f9                 ldr x10, [x8, #0x28]
   468: 4a010b8b                 add x10, x10, x11
   472: 0742c13d                 ldr q7, [x16, #0x500]
   476: 0852c13d                 ldr q8, [x16, #0x540]
   480: 4201c03d                 ldr q2, [x10]
   484: 431c294e                 and v3.16b, v2.16b, v9.16b
   488: 44040c6f                 ushr v4.16b, v2.16b, #4
   492: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   496: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   500: 001c256e                 eor v0.16b, v0.16b, v5.16b
   504: 001c266e                 eor v0.16b, v0.16b, v6.16b
   508: 0742c13d                 ldr q7, [x16, #0x500]
   512: 0852c13d                 ldr q8, [x16, #0x540]
   516: 4201c03d                 ldr q2, [x10]
   520: 431c294e                 and v3.16b, v2.16b, v9.16b
   524: 44040c6f                 ushr v4.16b, v2.16b, #4
   528: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   532: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   536: 211c256e                 eor v1.16b, v1.16b, v5.16b
   540: 211c266e                 eor v1.16b, v1.16b, v6.16b
   544: 2a0140f9                 ldr x10, [x9]
   548: 4a010b8b                 add x10, x10, x11
   552: 4001803d                 str q0, [x10]
   556: 2a0540f9                 ldr x10, [x9, #8]
   560: 4a010b8b                 add x10, x10, x11
   564: 4101803d                 str q1, [x10]
   568: 6b410091                 add x11, x11, #0x10
   572: ad0500f1                 subs x13, x13, #1
   576: 21efff54                 b.ne 0x24
   580: e827c1ac                 ldp q8, q9, [sp], #0x20
   584: c0035fd6                 ret 
   588: 00000000                 udf #0
   592: 00000000                 udf #0
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
