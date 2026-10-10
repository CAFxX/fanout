; config_t1_859 target=sve2 twin=nt topo=raid5(4)/enc op=encode U=1
; code_size=256 twins=["nt"] uops_per_byte=0.5
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: ed030caa                 mov x13, x12
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: ce0200b4                 cbz x14, 0x74
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
    76: 0a0d40f9                 ldr x10, [x8, #0x18]
    80: 4a010b8b                 add x10, x10, x11
    84: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
    88: 30021904                 eor z16.b, p0/m, z16.b, z17.b
    92: 2a0140f9                 ldr x10, [x9]
    96: 4a010b8b                 add x10, x10, x11
   100: 50e110e4                 stnt1b { z16.b }, p0, [x10]
   104: 6b010d8b                 add x11, x11, x13
   108: ce0500f1                 subs x14, x14, #1
   112: 81fdff54                 b.ne 0x20
   116: 7f0102eb                 cmp x11, x2
   120: 62020054                 b.hs 0xc4
   124: 2a0140f9                 ldr x10, [x9]
   128: ed030baa                 mov x13, x11
   132: 0e0140f9                 ldr x14, [x8]
   136: ce696d38                 ldrb w14, [x14, x13]
   140: ec030e2a                 mov w12, w14
   144: 0e0540f9                 ldr x14, [x8, #8]
   148: ce696d38                 ldrb w14, [x14, x13]
   152: 8c010e4a                 eor w12, w12, w14
   156: 0e0940f9                 ldr x14, [x8, #0x10]
   160: ce696d38                 ldrb w14, [x14, x13]
   164: 8c010e4a                 eor w12, w12, w14
   168: 0e0d40f9                 ldr x14, [x8, #0x18]
   172: ce696d38                 ldrb w14, [x14, x13]
   176: 8c010e4a                 eor w12, w12, w14
   180: 4c692d38                 strb w12, [x10, x13]
   184: ad050091                 add x13, x13, #1
   188: bf0102eb                 cmp x13, x2
   192: 23feff54                 b.lo 0x84
   196: c0035fd6                 ret 
   200: 00000000                 udf #0
   204: 00000000                 udf #0
   208: 00000000                 udf #0
   212: 00000000                 udf #0
   216: 00000000                 udf #0
   220: 00000000                 udf #0
   224: 00000000                 udf #0
   228: 00000000                 udf #0
   232: 00000000                 udf #0
   236: 00000000                 udf #0
   240: 00000000                 udf #0
   244: 00000000                 udf #0
   248: 00000000                 udf #0
   252: 00000000                 udf #0
