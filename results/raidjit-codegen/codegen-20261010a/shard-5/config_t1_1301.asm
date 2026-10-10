; config_t1_1301 target=neon twin=nt topo=rs(6,3)/rec{0} op=recover U=1
; code_size=320 twins=["nt"] uops_per_byte=0.75
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: 0e0280d2                 mov x14, #0x10
    16: 4d08ce9a                 udiv x13, x2, x14
    20: 0d0400b4                 cbz x13, 0x94
    24: 1f2003d5                 nop 
    28: 1f2003d5                 nop 
    32: 0a0140f9                 ldr x10, [x8]
    36: 4a010b8b                 add x10, x10, x11
    40: 4001c03d                 ldr q0, [x10]
    44: 0a0540f9                 ldr x10, [x8, #8]
    48: 4a010b8b                 add x10, x10, x11
    52: 4101c03d                 ldr q1, [x10]
    56: 001c216e                 eor v0.16b, v0.16b, v1.16b
    60: 0a0940f9                 ldr x10, [x8, #0x10]
    64: 4a010b8b                 add x10, x10, x11
    68: 4101c03d                 ldr q1, [x10]
    72: 001c216e                 eor v0.16b, v0.16b, v1.16b
    76: 0a0d40f9                 ldr x10, [x8, #0x18]
    80: 4a010b8b                 add x10, x10, x11
    84: 4101c03d                 ldr q1, [x10]
    88: 001c216e                 eor v0.16b, v0.16b, v1.16b
    92: 0a1140f9                 ldr x10, [x8, #0x20]
    96: 4a010b8b                 add x10, x10, x11
   100: 4101c03d                 ldr q1, [x10]
   104: 001c216e                 eor v0.16b, v0.16b, v1.16b
   108: 0a1540f9                 ldr x10, [x8, #0x28]
   112: 4a010b8b                 add x10, x10, x11
   116: 4101c03d                 ldr q1, [x10]
   120: 001c216e                 eor v0.16b, v0.16b, v1.16b
   124: 2a0140f9                 ldr x10, [x9]
   128: 4a010b8b                 add x10, x10, x11
   132: 4001803d                 str q0, [x10]
   136: 6b410091                 add x11, x11, #0x10
   140: ad0500f1                 subs x13, x13, #1
   144: 41fcff54                 b.ne 0x18
   148: 7f0102eb                 cmp x11, x2
   152: 22030054                 b.hs 0xfc
   156: 2a0140f9                 ldr x10, [x9]
   160: ed030baa                 mov x13, x11
   164: 0e0140f9                 ldr x14, [x8]
   168: ce696d38                 ldrb w14, [x14, x13]
   172: ec030e2a                 mov w12, w14
   176: 0e0540f9                 ldr x14, [x8, #8]
   180: ce696d38                 ldrb w14, [x14, x13]
   184: 8c010e4a                 eor w12, w12, w14
   188: 0e0940f9                 ldr x14, [x8, #0x10]
   192: ce696d38                 ldrb w14, [x14, x13]
   196: 8c010e4a                 eor w12, w12, w14
   200: 0e0d40f9                 ldr x14, [x8, #0x18]
   204: ce696d38                 ldrb w14, [x14, x13]
   208: 8c010e4a                 eor w12, w12, w14
   212: 0e1140f9                 ldr x14, [x8, #0x20]
   216: ce696d38                 ldrb w14, [x14, x13]
   220: 8c010e4a                 eor w12, w12, w14
   224: 0e1540f9                 ldr x14, [x8, #0x28]
   228: ce696d38                 ldrb w14, [x14, x13]
   232: 8c010e4a                 eor w12, w12, w14
   236: 4c692d38                 strb w12, [x10, x13]
   240: ad050091                 add x13, x13, #1
   244: bf0102eb                 cmp x13, x2
   248: 63fdff54                 b.lo 0xa4
   252: c0035fd6                 ret 
   256: 00000000                 udf #0
   260: 00000000                 udf #0
   264: 00000000                 udf #0
   268: 00000000                 udf #0
   272: 00000000                 udf #0
   276: 00000000                 udf #0
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
