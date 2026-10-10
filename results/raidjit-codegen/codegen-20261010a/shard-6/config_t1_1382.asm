; config_t1_1382 target=neon twin=notail topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=320 twins=["notail"] uops_per_byte=2.0625
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: e80300aa                 mov x8, x0
    12: e90301aa                 mov x9, x1
    16: 90090010                 adr x16, 0x140
    20: eb031faa                 mov x11, xzr
    24: 0e0280d2                 mov x14, #0x10
    28: 4d08ce9a                 udiv x13, x2, x14
    32: 6d0700b4                 cbz x13, 0x10c
    36: eae5004f                 movi v10.16b, #0xf
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0940f9                 ldr x10, [x8, #0x10]
    68: 4a010b8b                 add x10, x10, x11
    72: 4001c03d                 ldr q0, [x10]
    76: 0a0d40f9                 ldr x10, [x8, #0x18]
    80: 4a010b8b                 add x10, x10, x11
    84: 4101c03d                 ldr q1, [x10]
    88: 0a1140f9                 ldr x10, [x8, #0x20]
    92: 4a010b8b                 add x10, x10, x11
    96: 4201c03d                 ldr q2, [x10]
   100: 0a0140f9                 ldr x10, [x8]
   104: 4a010b8b                 add x10, x10, x11
   108: 030540f9                 ldr x3, [x8, #8]
   112: 63000b8b                 add x3, x3, x11
   116: 4301c03d                 ldr q3, [x10]
   120: 6600c03d                 ldr q6, [x3]
   124: 631c266e                 eor v3.16b, v3.16b, v6.16b
   128: 001c236e                 eor v0.16b, v0.16b, v3.16b
   132: 0802c03d                 ldr q8, [x16]
   136: 0912c03d                 ldr q9, [x16, #0x40]
   140: 4301c03d                 ldr q3, [x10]
   144: 6600c03d                 ldr q6, [x3]
   148: 631c266e                 eor v3.16b, v3.16b, v6.16b
   152: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   156: 65040c6f                 ushr v5.16b, v3.16b, #4
   160: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   164: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   168: 211c266e                 eor v1.16b, v1.16b, v6.16b
   172: 211c276e                 eor v1.16b, v1.16b, v7.16b
   176: 0822c03d                 ldr q8, [x16, #0x80]
   180: 0932c03d                 ldr q9, [x16, #0xc0]
   184: 4301c03d                 ldr q3, [x10]
   188: 6600c03d                 ldr q6, [x3]
   192: 631c266e                 eor v3.16b, v3.16b, v6.16b
   196: 641c2a4e                 and v4.16b, v3.16b, v10.16b
   200: 65040c6f                 ushr v5.16b, v3.16b, #4
   204: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   208: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   212: 421c266e                 eor v2.16b, v2.16b, v6.16b
   216: 421c276e                 eor v2.16b, v2.16b, v7.16b
   220: 2a0140f9                 ldr x10, [x9]
   224: 4a010b8b                 add x10, x10, x11
   228: 4001803d                 str q0, [x10]
   232: 2a0540f9                 ldr x10, [x9, #8]
   236: 4a010b8b                 add x10, x10, x11
   240: 4101803d                 str q1, [x10]
   244: 2a0940f9                 ldr x10, [x9, #0x10]
   248: 4a010b8b                 add x10, x10, x11
   252: 4201803d                 str q2, [x10]
   256: 6b410091                 add x11, x11, #0x10
   260: ad0500f1                 subs x13, x13, #1
   264: 01f9ff54                 b.ne 0x28
   268: ea2fc1ac                 ldp q10, q11, [sp], #0x20
   272: e827c1ac                 ldp q8, q9, [sp], #0x20
   276: c0035fd6                 ret 
   280: 00000000                 udf #0
   284: 00000000                 udf #0
   288: 00000000                 udf #0
   292: 00000000                 udf #0
   296: 00000000                 udf #0
   300: 00000000                 udf #0
   304: 00000000                 udf #0
   308: 00000000                 udf #0
   312: 00000000                 udf #0
   316: 00000000                 udf #0
