; config_t1_1096 target=sve2 twin=notail topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=320 twins=["notail"] uops_per_byte=2.1875
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0090010                 adr x16, 0x140
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ed030caa                 mov x13, x12
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: fac13825                 mov z26.b, #0xf
    36: 8e0700b4                 cbz x14, 0x114
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0940f9                 ldr x10, [x8, #0x10]
    68: 4a010b8b                 add x10, x10, x11
    72: 50a100a4                 ld1b { z16.b }, p0/z, [x10]
    76: 0a0d40f9                 ldr x10, [x8, #0x18]
    80: 4a010b8b                 add x10, x10, x11
    84: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
    88: 0a1140f9                 ldr x10, [x8, #0x20]
    92: 4a010b8b                 add x10, x10, x11
    96: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   100: 0a0140f9                 ldr x10, [x8]
   104: 4a010b8b                 add x10, x10, x11
   108: 030540f9                 ldr x3, [x8, #8]
   112: 63000b8b                 add x3, x3, x11
   116: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   120: 76a000a4                 ld1b { z22.b }, p0/z, [x3]
   124: d3021904                 eor z19.b, p0/m, z19.b, z22.b
   128: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   132: 0f020091                 add x15, x16, #0
   136: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   140: 0f020191                 add x15, x16, #0x40
   144: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   148: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   152: 76a000a4                 ld1b { z22.b }, p0/z, [x3]
   156: d3021904                 eor z19.b, p0/m, z19.b, z22.b
   160: 74323a04                 and z20.d, z19.d, z26.d
   164: 75962c04                 lsr z21.b, z19.b, #4
   168: 16333405                 tbl z22.b, { z24.b }, z20.b
   172: 37333505                 tbl z23.b, { z25.b }, z21.b
   176: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   180: 0f020291                 add x15, x16, #0x80
   184: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   188: 0f020391                 add x15, x16, #0xc0
   192: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   196: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   200: 76a000a4                 ld1b { z22.b }, p0/z, [x3]
   204: d3021904                 eor z19.b, p0/m, z19.b, z22.b
   208: 74323a04                 and z20.d, z19.d, z26.d
   212: 75962c04                 lsr z21.b, z19.b, #4
   216: 16333405                 tbl z22.b, { z24.b }, z20.b
   220: 37333505                 tbl z23.b, { z25.b }, z21.b
   224: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   228: 2a0140f9                 ldr x10, [x9]
   232: 4a010b8b                 add x10, x10, x11
   236: 50e100e4                 st1b { z16.b }, p0, [x10]
   240: 2a0540f9                 ldr x10, [x9, #8]
   244: 4a010b8b                 add x10, x10, x11
   248: 51e100e4                 st1b { z17.b }, p0, [x10]
   252: 2a0940f9                 ldr x10, [x9, #0x10]
   256: 4a010b8b                 add x10, x10, x11
   260: 52e100e4                 st1b { z18.b }, p0, [x10]
   264: 6b010d8b                 add x11, x11, x13
   268: ce0500f1                 subs x14, x14, #1
   272: c1f8ff54                 b.ne 0x28
   276: c0035fd6                 ret 
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
