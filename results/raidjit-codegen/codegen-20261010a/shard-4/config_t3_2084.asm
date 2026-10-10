; config_t3_2084 target=neon twin=baseline topo=raid1()/enc op=encode U=8
; code_size=192 twins=[] uops_per_byte=0.125
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: e80300aa                 mov x8, x0
     8: e90301aa                 mov x9, x1
    12: eb031faa                 mov x11, xzr
    16: 0e1080d2                 mov x14, #0x80
    20: 4d08ce9a                 udiv x13, x2, x14
    24: 2d0300b4                 cbz x13, 0x7c
    28: 1f2003d5                 nop 
    32: 0a0140f9                 ldr x10, [x8]
    36: 4a010b8b                 add x10, x10, x11
    40: 4001c03d                 ldr q0, [x10]
    44: 4105c03d                 ldr q1, [x10, #0x10]
    48: 4209c03d                 ldr q2, [x10, #0x20]
    52: 430dc03d                 ldr q3, [x10, #0x30]
    56: 4411c03d                 ldr q4, [x10, #0x40]
    60: 4515c03d                 ldr q5, [x10, #0x50]
    64: 4619c03d                 ldr q6, [x10, #0x60]
    68: 471dc03d                 ldr q7, [x10, #0x70]
    72: 2a0140f9                 ldr x10, [x9]
    76: 4a010b8b                 add x10, x10, x11
    80: 4001803d                 str q0, [x10]
    84: 4105803d                 str q1, [x10, #0x10]
    88: 4209803d                 str q2, [x10, #0x20]
    92: 430d803d                 str q3, [x10, #0x30]
    96: 4411803d                 str q4, [x10, #0x40]
   100: 4515803d                 str q5, [x10, #0x50]
   104: 4619803d                 str q6, [x10, #0x60]
   108: 471d803d                 str q7, [x10, #0x70]
   112: 6b010291                 add x11, x11, #0x80
   116: ad0500f1                 subs x13, x13, #1
   120: 21fdff54                 b.ne 0x1c
   124: 7f0102eb                 cmp x11, x2
   128: 42010054                 b.hs 0xa8
   132: 2a0140f9                 ldr x10, [x9]
   136: ed030baa                 mov x13, x11
   140: 0e0140f9                 ldr x14, [x8]
   144: ce696d38                 ldrb w14, [x14, x13]
   148: ec030e2a                 mov w12, w14
   152: 4c692d38                 strb w12, [x10, x13]
   156: ad050091                 add x13, x13, #1
   160: bf0102eb                 cmp x13, x2
   164: 43ffff54                 b.lo 0x8c
   168: e827c1ac                 ldp q8, q9, [sp], #0x20
   172: c0035fd6                 ret 
   176: 00000000                 udf #0
   180: 00000000                 udf #0
   184: 00000000                 udf #0
   188: 00000000                 udf #0
