; config_t1_1304 target=neon twin=notail topo=rs(6,3)/rec{0} op=recover U=1
; code_size=192 twins=["notail"] uops_per_byte=0.75
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
   148: c0035fd6                 ret 
   152: 00000000                 udf #0
   156: 00000000                 udf #0
   160: 00000000                 udf #0
   164: 00000000                 udf #0
   168: 00000000                 udf #0
   172: 00000000                 udf #0
   176: 00000000                 udf #0
   180: 00000000                 udf #0
   184: 00000000                 udf #0
   188: 00000000                 udf #0
