; config_t1_1408 target=neon twin=notail topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=320 twins=["notail"] uops_per_byte=2.1875
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: e80300aa                 mov x8, x0
     8: e90301aa                 mov x9, x1
    12: b0090010                 adr x16, 0x140
    16: eb031faa                 mov x11, xzr
    20: 0e0280d2                 mov x14, #0x10
    24: 4d08ce9a                 udiv x13, x2, x14
    28: cd0700b4                 cbz x13, 0x114
    32: e9e5004f                 movi v9.16b, #0xf
    36: 1f2003d5                 nop 
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a1140f9                 ldr x10, [x8, #0x20]
    68: 4a010b8b                 add x10, x10, x11
    72: 4001c03d                 ldr q0, [x10]
    76: 0a1540f9                 ldr x10, [x8, #0x28]
    80: 4a010b8b                 add x10, x10, x11
    84: 4101c03d                 ldr q1, [x10]
    88: 0a0140f9                 ldr x10, [x8]
    92: 4a010b8b                 add x10, x10, x11
    96: 030940f9                 ldr x3, [x8, #0x10]
   100: 63000b8b                 add x3, x3, x11
   104: 4201c03d                 ldr q2, [x10]
   108: 6500c03d                 ldr q5, [x3]
   112: 421c256e                 eor v2.16b, v2.16b, v5.16b
   116: 001c226e                 eor v0.16b, v0.16b, v2.16b
   120: 0702c03d                 ldr q7, [x16]
   124: 0812c03d                 ldr q8, [x16, #0x40]
   128: 4201c03d                 ldr q2, [x10]
   132: 6500c03d                 ldr q5, [x3]
   136: 421c256e                 eor v2.16b, v2.16b, v5.16b
   140: 431c294e                 and v3.16b, v2.16b, v9.16b
   144: 44040c6f                 ushr v4.16b, v2.16b, #4
   148: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   152: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   156: 211c256e                 eor v1.16b, v1.16b, v5.16b
   160: 211c266e                 eor v1.16b, v1.16b, v6.16b
   164: 0a0540f9                 ldr x10, [x8, #8]
   168: 4a010b8b                 add x10, x10, x11
   172: 030d40f9                 ldr x3, [x8, #0x18]
   176: 63000b8b                 add x3, x3, x11
   180: 4201c03d                 ldr q2, [x10]
   184: 6500c03d                 ldr q5, [x3]
   188: 421c256e                 eor v2.16b, v2.16b, v5.16b
   192: 001c226e                 eor v0.16b, v0.16b, v2.16b
   196: 0722c03d                 ldr q7, [x16, #0x80]
   200: 0832c03d                 ldr q8, [x16, #0xc0]
   204: 4201c03d                 ldr q2, [x10]
   208: 6500c03d                 ldr q5, [x3]
   212: 421c256e                 eor v2.16b, v2.16b, v5.16b
   216: 431c294e                 and v3.16b, v2.16b, v9.16b
   220: 44040c6f                 ushr v4.16b, v2.16b, #4
   224: e500034e                 tbl v5.16b, { v7.16b }, v3.16b
   228: 0601044e                 tbl v6.16b, { v8.16b }, v4.16b
   232: 211c256e                 eor v1.16b, v1.16b, v5.16b
   236: 211c266e                 eor v1.16b, v1.16b, v6.16b
   240: 2a0140f9                 ldr x10, [x9]
   244: 4a010b8b                 add x10, x10, x11
   248: 4001803d                 str q0, [x10]
   252: 2a0540f9                 ldr x10, [x9, #8]
   256: 4a010b8b                 add x10, x10, x11
   260: 4101803d                 str q1, [x10]
   264: 6b410091                 add x11, x11, #0x10
   268: ad0500f1                 subs x13, x13, #1
   272: a1f8ff54                 b.ne 0x24
   276: e827c1ac                 ldp q8, q9, [sp], #0x20
   280: c0035fd6                 ret 
   284: 00000000                 udf #0
   288: 00000000                 udf #0
   292: 00000000                 udf #0
   296: 00000000                 udf #0
   300: 00000000                 udf #0
   304: 00000000                 udf #0
   308: 00000000                 udf #0
   312: 00000000                 udf #0
   316: 00000000                 udf #0
