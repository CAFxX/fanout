; config_t1_1097 target=sve2 twin=predtail topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=512 twins=["predtail"] uops_per_byte=2.1875
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d00f0010                 adr x16, 0x200
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
   276: 7f0102eb                 cmp x11, x2
   280: 02070054                 b.hs 0x1f8
   284: 4d000bcb                 sub x13, x2, x11
   288: e1172d25                 whilelt p1.b, xzr, x13
   292: 0a0940f9                 ldr x10, [x8, #0x10]
   296: 4a010b8b                 add x10, x10, x11
   300: 50a500a4                 ld1b { z16.b }, p1/z, [x10]
   304: 0a0d40f9                 ldr x10, [x8, #0x18]
   308: 4a010b8b                 add x10, x10, x11
   312: 51a500a4                 ld1b { z17.b }, p1/z, [x10]
   316: 0a1140f9                 ldr x10, [x8, #0x20]
   320: 4a010b8b                 add x10, x10, x11
   324: 52a500a4                 ld1b { z18.b }, p1/z, [x10]
   328: 0a0140f9                 ldr x10, [x8]
   332: 4a010b8b                 add x10, x10, x11
   336: 030540f9                 ldr x3, [x8, #8]
   340: 63000b8b                 add x3, x3, x11
   344: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   348: 76a400a4                 ld1b { z22.b }, p1/z, [x3]
   352: d3061904                 eor z19.b, p1/m, z19.b, z22.b
   356: 70061904                 eor z16.b, p1/m, z16.b, z19.b
   360: 0f020091                 add x15, x16, #0
   364: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   368: 0f020191                 add x15, x16, #0x40
   372: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   376: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   380: 76a400a4                 ld1b { z22.b }, p1/z, [x3]
   384: d3061904                 eor z19.b, p1/m, z19.b, z22.b
   388: 74323a04                 and z20.d, z19.d, z26.d
   392: 75962c04                 lsr z21.b, z19.b, #4
   396: 16333405                 tbl z22.b, { z24.b }, z20.b
   400: 37333505                 tbl z23.b, { z25.b }, z21.b
   404: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   408: 0f020291                 add x15, x16, #0x80
   412: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   416: 0f020391                 add x15, x16, #0xc0
   420: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   424: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   428: 76a400a4                 ld1b { z22.b }, p1/z, [x3]
   432: d3061904                 eor z19.b, p1/m, z19.b, z22.b
   436: 74323a04                 and z20.d, z19.d, z26.d
   440: 75962c04                 lsr z21.b, z19.b, #4
   444: 16333405                 tbl z22.b, { z24.b }, z20.b
   448: 37333505                 tbl z23.b, { z25.b }, z21.b
   452: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   456: 2a0140f9                 ldr x10, [x9]
   460: 4a010b8b                 add x10, x10, x11
   464: 50e500e4                 st1b { z16.b }, p1, [x10]
   468: 2a0540f9                 ldr x10, [x9, #8]
   472: 4a010b8b                 add x10, x10, x11
   476: 51e500e4                 st1b { z17.b }, p1, [x10]
   480: 2a0940f9                 ldr x10, [x9, #0x10]
   484: 4a010b8b                 add x10, x10, x11
   488: 52e500e4                 st1b { z18.b }, p1, [x10]
   492: 6b010c8b                 add x11, x11, x12
   496: ad010ceb                 subs x13, x13, x12
   500: 6cf9ff54                 b.gt 0x120
   504: c0035fd6                 ret 
   508: 00000000                 udf #0
