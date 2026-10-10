; config_t3_2085 target=neon twin=baseline topo=rs(3,1)/enc op=encode U=1
; code_size=192 twins=[] uops_per_byte=0.375
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: 0e0280d2                 mov x14, #0x10
    16: 4d08ce9a                 udiv x13, x2, x14
    20: 8d0200b4                 cbz x13, 0x64
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
    76: 2a0140f9                 ldr x10, [x9]
    80: 4a010b8b                 add x10, x10, x11
    84: 4001803d                 str q0, [x10]
    88: 6b410091                 add x11, x11, #0x10
    92: ad0500f1                 subs x13, x13, #1
    96: c1fdff54                 b.ne 0x18
   100: 7f0102eb                 cmp x11, x2
   104: 02020054                 b.hs 0xa8
   108: 2a0140f9                 ldr x10, [x9]
   112: ed030baa                 mov x13, x11
   116: 0e0140f9                 ldr x14, [x8]
   120: ce696d38                 ldrb w14, [x14, x13]
   124: ec030e2a                 mov w12, w14
   128: 0e0540f9                 ldr x14, [x8, #8]
   132: ce696d38                 ldrb w14, [x14, x13]
   136: 8c010e4a                 eor w12, w12, w14
   140: 0e0940f9                 ldr x14, [x8, #0x10]
   144: ce696d38                 ldrb w14, [x14, x13]
   148: 8c010e4a                 eor w12, w12, w14
   152: 4c692d38                 strb w12, [x10, x13]
   156: ad050091                 add x13, x13, #1
   160: bf0102eb                 cmp x13, x2
   164: 83feff54                 b.lo 0x74
   168: c0035fd6                 ret 
   172: 00000000                 udf #0
   176: 00000000                 udf #0
   180: 00000000                 udf #0
   184: 00000000                 udf #0
   188: 00000000                 udf #0
