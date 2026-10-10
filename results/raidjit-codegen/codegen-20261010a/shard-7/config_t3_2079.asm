; config_t3_2079 target=sve2 twin=baseline topo=rs(3,1)/enc op=encode U=1
; code_size=192 twins=[] uops_per_byte=0.375
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: ed030caa                 mov x13, x12
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: 4e0200b4                 cbz x14, 0x64
    32: 0a0140f9                 ldr x10, [x8]
    36: 4a010b8b                 add x10, x10, x11
    40: 50a100a4                 ld1b { z16.b }, p0/z, [x10]
    44: 0a0540f9                 ldr x10, [x8, #8]
    48: 4a010b8b                 add x10, x10, x11
    52: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
    56: 30021904                 eor z16.b, p0/m, z16.b, z17.b
    60: 0a0940f9                 ldr x10, [x8, #0x10]
    64: 4a010b8b                 add x10, x10, x11
    68: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
    72: 30021904                 eor z16.b, p0/m, z16.b, z17.b
    76: 2a0140f9                 ldr x10, [x9]
    80: 4a010b8b                 add x10, x10, x11
    84: 50e100e4                 st1b { z16.b }, p0, [x10]
    88: 6b010d8b                 add x11, x11, x13
    92: ce0500f1                 subs x14, x14, #1
    96: 01feff54                 b.ne 0x20
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
