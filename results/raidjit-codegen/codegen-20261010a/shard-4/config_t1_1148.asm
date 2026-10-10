; config_t1_1148 target=neon twin=notail topo=raid5(4)/enc op=encode U=1
; code_size=128 twins=["notail"] uops_per_byte=0.5
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: 0e0280d2                 mov x14, #0x10
    16: 4d08ce9a                 udiv x13, x2, x14
    20: 0d0300b4                 cbz x13, 0x74
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
    92: 2a0140f9                 ldr x10, [x9]
    96: 4a010b8b                 add x10, x10, x11
   100: 4001803d                 str q0, [x10]
   104: 6b410091                 add x11, x11, #0x10
   108: ad0500f1                 subs x13, x13, #1
   112: 41fdff54                 b.ne 0x18
   116: c0035fd6                 ret 
   120: 00000000                 udf #0
   124: 00000000                 udf #0
