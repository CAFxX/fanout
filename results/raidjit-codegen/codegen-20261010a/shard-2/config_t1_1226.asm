; config_t1_1226 target=neon twin=notail topo=rs(6,3)/enc op=encode U=1
; code_size=640 twins=["notail"] uops_per_byte=6.875
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: e80300aa                 mov x8, x0
    12: e90301aa                 mov x9, x1
    16: 90130010                 adr x16, 0x280
    20: eb031faa                 mov x11, xzr
    24: 0e0280d2                 mov x14, #0x10
    28: 4d08ce9a                 udiv x13, x2, x14
    32: 4d1100b4                 cbz x13, 0x248
    36: eae5004f                 movi v10.16b, #0xf
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 4301c03d                 ldr q3, [x10]
    76: 601ca34e                 mov v0.16b, v3.16b
    80: 0802c03d                 ldr q8, [x16]
    84: 0912c03d                 ldr q9, [x16, #0x40]
    88: 4301c03d                 ldr q3, [x10]
    92: 641c2a4e                 and v4.16b, v3.16b, v10.16b
    96: 65040c6f                 ushr v5.16b, v3.16b, #4
   100: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   104: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   108: c11ca64e                 mov v1.16b, v6.16b
   112: 211c276e                 eor v1.16b, v1.16b, v7.16b
   116: 0822c03d                 ldr q8, [x16, #0x80]
   120: 0932c03d                 ldr q9, [x16, #0xc0]
   124: 4301c03d                 ldr q3, [x10]
   128: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   132: 65040c6f                 ushr v5.16b, v3.16b, #4
   136: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   140: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   144: c21ca64e                 mov v2.16b, v6.16b
   148: 421c276e                 eor v2.16b, v2.16b, v7.16b
   152: 0a0540f9                 ldr x10, [x8, #8]
   156: 4a010b8b                 add x10, x10, x11
   160: 4301c03d                 ldr q3, [x10]
   164: 001c236e                 eor v0.16b, v0.16b, v3.16b
   168: 0822c03d                 ldr q8, [x16, #0x80]
   172: 0932c03d                 ldr q9, [x16, #0xc0]
   176: 4301c03d                 ldr q3, [x10]
   180: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   184: 65040c6f                 ushr v5.16b, v3.16b, #4
   188: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   192: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   196: 211c266e                 eor v1.16b, v1.16b, v6.16b
   200: 211c276e                 eor v1.16b, v1.16b, v7.16b
   204: 0842c03d                 ldr q8, [x16, #0x100]
   208: 0952c03d                 ldr q9, [x16, #0x140]
   212: 4301c03d                 ldr q3, [x10]
   216: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   220: 65040c6f                 ushr v5.16b, v3.16b, #4
   224: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   228: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   232: 421c266e                 eor v2.16b, v2.16b, v6.16b
   236: 421c276e                 eor v2.16b, v2.16b, v7.16b
   240: 0a0940f9                 ldr x10, [x8, #0x10]
   244: 4a010b8b                 add x10, x10, x11
   248: 4301c03d                 ldr q3, [x10]
   252: 001c236e                 eor v0.16b, v0.16b, v3.16b
   256: 0842c03d                 ldr q8, [x16, #0x100]
   260: 0952c03d                 ldr q9, [x16, #0x140]
   264: 4301c03d                 ldr q3, [x10]
   268: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   272: 65040c6f                 ushr v5.16b, v3.16b, #4
   276: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   280: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   284: 211c266e                 eor v1.16b, v1.16b, v6.16b
   288: 211c276e                 eor v1.16b, v1.16b, v7.16b
   292: 0862c03d                 ldr q8, [x16, #0x180]
   296: 0972c03d                 ldr q9, [x16, #0x1c0]
   300: 4301c03d                 ldr q3, [x10]
   304: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   308: 65040c6f                 ushr v5.16b, v3.16b, #4
   312: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   316: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   320: 421c266e                 eor v2.16b, v2.16b, v6.16b
   324: 421c276e                 eor v2.16b, v2.16b, v7.16b
   328: 0a0d40f9                 ldr x10, [x8, #0x18]
   332: 4a010b8b                 add x10, x10, x11
   336: 4301c03d                 ldr q3, [x10]
   340: 001c236e                 eor v0.16b, v0.16b, v3.16b
   344: 0882c03d                 ldr q8, [x16, #0x200]
   348: 0992c03d                 ldr q9, [x16, #0x240]
   352: 4301c03d                 ldr q3, [x10]
   356: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   360: 65040c6f                 ushr v5.16b, v3.16b, #4
   364: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   368: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   372: 211c266e                 eor v1.16b, v1.16b, v6.16b
   376: 211c276e                 eor v1.16b, v1.16b, v7.16b
   380: 0802c03d                 ldr q8, [x16]
   384: 0912c03d                 ldr q9, [x16, #0x40]
   388: 4301c03d                 ldr q3, [x10]
   392: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   396: 65040c6f                 ushr v5.16b, v3.16b, #4
   400: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   404: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   408: 421c266e                 eor v2.16b, v2.16b, v6.16b
   412: 421c276e                 eor v2.16b, v2.16b, v7.16b
   416: 0a1140f9                 ldr x10, [x8, #0x20]
   420: 4a010b8b                 add x10, x10, x11
   424: 4301c03d                 ldr q3, [x10]
   428: 001c236e                 eor v0.16b, v0.16b, v3.16b
   432: 0862c03d                 ldr q8, [x16, #0x180]
   436: 0972c03d                 ldr q9, [x16, #0x1c0]
   440: 4301c03d                 ldr q3, [x10]
   444: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   448: 65040c6f                 ushr v5.16b, v3.16b, #4
   452: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   456: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   460: 211c266e                 eor v1.16b, v1.16b, v6.16b
   464: 211c276e                 eor v1.16b, v1.16b, v7.16b
   468: 0882c03d                 ldr q8, [x16, #0x200]
   472: 0992c03d                 ldr q9, [x16, #0x240]
   476: 4301c03d                 ldr q3, [x10]
   480: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   484: 65040c6f                 ushr v5.16b, v3.16b, #4
   488: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   492: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   496: 421c266e                 eor v2.16b, v2.16b, v6.16b
   500: 421c276e                 eor v2.16b, v2.16b, v7.16b
   504: 0a1540f9                 ldr x10, [x8, #0x28]
   508: 4a010b8b                 add x10, x10, x11
   512: 4301c03d                 ldr q3, [x10]
   516: 001c236e                 eor v0.16b, v0.16b, v3.16b
   520: 4301c03d                 ldr q3, [x10]
   524: 211c236e                 eor v1.16b, v1.16b, v3.16b
   528: 4301c03d                 ldr q3, [x10]
   532: 421c236e                 eor v2.16b, v2.16b, v3.16b
   536: 2a0140f9                 ldr x10, [x9]
   540: 4a010b8b                 add x10, x10, x11
   544: 4001803d                 str q0, [x10]
   548: 2a0540f9                 ldr x10, [x9, #8]
   552: 4a010b8b                 add x10, x10, x11
   556: 4101803d                 str q1, [x10]
   560: 2a0940f9                 ldr x10, [x9, #0x10]
   564: 4a010b8b                 add x10, x10, x11
   568: 4201803d                 str q2, [x10]
   572: 6b410091                 add x11, x11, #0x10
   576: ad0500f1                 subs x13, x13, #1
   580: 21efff54                 b.ne 0x28
   584: ea2fc1ac                 ldp q10, q11, [sp], #0x20
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
