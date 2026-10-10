; config_t1_1015 target=sve2 twin=nt topo=rs(6,3)/rec{0} op=recover U=1
; code_size=320 twins=["nt"] uops_per_byte=0.75
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: ed030caa                 mov x13, x12
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: ce0300b4                 cbz x14, 0x94
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
    92: 0a1140f9                 ldr x10, [x8, #0x20]
    96: 4a010b8b                 add x10, x10, x11
   100: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
   104: 30021904                 eor z16.b, p0/m, z16.b, z17.b
   108: 0a1540f9                 ldr x10, [x8, #0x28]
   112: 4a010b8b                 add x10, x10, x11
   116: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
   120: 30021904                 eor z16.b, p0/m, z16.b, z17.b
   124: 2a0140f9                 ldr x10, [x9]
   128: 4a010b8b                 add x10, x10, x11
   132: 50e110e4                 stnt1b { z16.b }, p0, [x10]
   136: 6b010d8b                 add x11, x11, x13
   140: ce0500f1                 subs x14, x14, #1
   144: 81fcff54                 b.ne 0x20
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
