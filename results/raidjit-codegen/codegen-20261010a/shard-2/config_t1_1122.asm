; config_t1_1122 target=sve2 twin=notail topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=320 twins=["notail"] uops_per_byte=2.3125
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0090010                 adr x16, 0x140
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ed030caa                 mov x13, x12
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: f9c13825                 mov z25.b, #0xf
    36: ce0700b4                 cbz x14, 0x11c
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a1140f9                 ldr x10, [x8, #0x20]
    68: 4a010b8b                 add x10, x10, x11
    72: 50a100a4                 ld1b { z16.b }, p0/z, [x10]
    76: 0a1540f9                 ldr x10, [x8, #0x28]
    80: 4a010b8b                 add x10, x10, x11
    84: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
    88: 0a0140f9                 ldr x10, [x8]
    92: 4a010b8b                 add x10, x10, x11
    96: 030940f9                 ldr x3, [x8, #0x10]
   100: 63000b8b                 add x3, x3, x11
   104: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   108: 75a000a4                 ld1b { z21.b }, p0/z, [x3]
   112: b2021904                 eor z18.b, p0/m, z18.b, z21.b
   116: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   120: 0f020091                 add x15, x16, #0
   124: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   128: 0f020191                 add x15, x16, #0x40
   132: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   136: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   140: 75a000a4                 ld1b { z21.b }, p0/z, [x3]
   144: b2021904                 eor z18.b, p0/m, z18.b, z21.b
   148: 53323904                 and z19.d, z18.d, z25.d
   152: 54962c04                 lsr z20.b, z18.b, #4
   156: f5323305                 tbl z21.b, { z23.b }, z19.b
   160: 16333405                 tbl z22.b, { z24.b }, z20.b
   164: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   168: 0a0540f9                 ldr x10, [x8, #8]
   172: 4a010b8b                 add x10, x10, x11
   176: 030d40f9                 ldr x3, [x8, #0x18]
   180: 63000b8b                 add x3, x3, x11
   184: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   188: 75a000a4                 ld1b { z21.b }, p0/z, [x3]
   192: b2021904                 eor z18.b, p0/m, z18.b, z21.b
   196: 50021904                 eor z16.b, p0/m, z16.b, z18.b
   200: 0f020291                 add x15, x16, #0x80
   204: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   208: 0f020391                 add x15, x16, #0xc0
   212: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   216: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   220: 75a000a4                 ld1b { z21.b }, p0/z, [x3]
   224: b2021904                 eor z18.b, p0/m, z18.b, z21.b
   228: 53323904                 and z19.d, z18.d, z25.d
   232: 54962c04                 lsr z20.b, z18.b, #4
   236: f5323305                 tbl z21.b, { z23.b }, z19.b
   240: 16333405                 tbl z22.b, { z24.b }, z20.b
   244: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   248: 2a0140f9                 ldr x10, [x9]
   252: 4a010b8b                 add x10, x10, x11
   256: 50e100e4                 st1b { z16.b }, p0, [x10]
   260: 2a0540f9                 ldr x10, [x9, #8]
   264: 4a010b8b                 add x10, x10, x11
   268: 51e100e4                 st1b { z17.b }, p0, [x10]
   272: 6b010d8b                 add x11, x11, x13
   276: ce0500f1                 subs x14, x14, #1
   280: 81f8ff54                 b.ne 0x28
   284: c0035fd6                 ret 
   288: 00000000                 udf #0
   292: 00000000                 udf #0
   296: 00000000                 udf #0
   300: 00000000                 udf #0
   304: 00000000                 udf #0
   308: 00000000                 udf #0
   312: 00000000                 udf #0
   316: 00000000                 udf #0
