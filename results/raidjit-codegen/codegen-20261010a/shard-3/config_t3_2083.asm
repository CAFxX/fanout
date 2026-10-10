; config_t3_2083 target=neon twin=baseline topo=raid1()/enc op=encode U=1
; code_size=128 twins=[] uops_per_byte=0.125
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: 0e0280d2                 mov x14, #0x10
    16: 4d08ce9a                 udiv x13, x2, x14
    20: 8d0100b4                 cbz x13, 0x44
    24: 1f2003d5                 nop 
    28: 1f2003d5                 nop 
    32: 0a0140f9                 ldr x10, [x8]
    36: 4a010b8b                 add x10, x10, x11
    40: 4001c03d                 ldr q0, [x10]
    44: 2a0140f9                 ldr x10, [x9]
    48: 4a010b8b                 add x10, x10, x11
    52: 4001803d                 str q0, [x10]
    56: 6b410091                 add x11, x11, #0x10
    60: ad0500f1                 subs x13, x13, #1
    64: c1feff54                 b.ne 0x18
    68: 7f0102eb                 cmp x11, x2
    72: 42010054                 b.hs 0x70
    76: 2a0140f9                 ldr x10, [x9]
    80: ed030baa                 mov x13, x11
    84: 0e0140f9                 ldr x14, [x8]
    88: ce696d38                 ldrb w14, [x14, x13]
    92: ec030e2a                 mov w12, w14
    96: 4c692d38                 strb w12, [x10, x13]
   100: ad050091                 add x13, x13, #1
   104: bf0102eb                 cmp x13, x2
   108: 43ffff54                 b.lo 0x54
   112: c0035fd6                 ret 
   116: 00000000                 udf #0
   120: 00000000                 udf #0
   124: 00000000                 udf #0
