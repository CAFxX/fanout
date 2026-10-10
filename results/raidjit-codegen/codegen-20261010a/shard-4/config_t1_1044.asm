; config_t1_1044 target=sve2 twin=notail topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=640 twins=["notail"] uops_per_byte=7.6875
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0130010                 adr x16, 0x280
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ed030caa                 mov x13, x12
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: f9c13825                 mov z25.b, #0xf
    36: 8e1200b4                 cbz x14, 0x274
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 0f020091                 add x15, x16, #0
    76: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
    80: 0f020191                 add x15, x16, #0x40
    84: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
    88: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
    92: 53323904                 and z19.d, z18.d, z25.d
    96: 54962c04                 lsr z20.b, z18.b, #4
   100: f5323305                 tbl z21.b, { z23.b }, z19.b
   104: 16333405                 tbl z22.b, { z24.b }, z20.b
   108: b032b604                 eor z16.d, z21.d, z22.d
   112: 0f020291                 add x15, x16, #0x80
   116: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   120: 0f020391                 add x15, x16, #0xc0
   124: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   128: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   132: 53323904                 and z19.d, z18.d, z25.d
   136: 54962c04                 lsr z20.b, z18.b, #4
   140: f5323305                 tbl z21.b, { z23.b }, z19.b
   144: 16333405                 tbl z22.b, { z24.b }, z20.b
   148: b132b604                 eor z17.d, z21.d, z22.d
   152: 0a0540f9                 ldr x10, [x8, #8]
   156: 4a010b8b                 add x10, x10, x11
   160: 0f020491                 add x15, x16, #0x100
   164: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   168: 0f020591                 add x15, x16, #0x140
   172: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   176: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   180: 53323904                 and z19.d, z18.d, z25.d
   184: 54962c04                 lsr z20.b, z18.b, #4
   188: f5323305                 tbl z21.b, { z23.b }, z19.b
   192: 16333405                 tbl z22.b, { z24.b }, z20.b
   196: d03a3504                 eor3 z16.d, z16.d, z21.d, z22.d
   200: 0f020691                 add x15, x16, #0x180
   204: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   208: 0f020791                 add x15, x16, #0x1c0
   212: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   216: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   220: 53323904                 and z19.d, z18.d, z25.d
   224: 54962c04                 lsr z20.b, z18.b, #4
   228: f5323305                 tbl z21.b, { z23.b }, z19.b
   232: 16333405                 tbl z22.b, { z24.b }, z20.b
   236: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   240: 0a0940f9                 ldr x10, [x8, #0x10]
   244: 4a010b8b                 add x10, x10, x11
   248: 0f020891                 add x15, x16, #0x200
   252: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   256: 0f020991                 add x15, x16, #0x240
   260: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   264: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   268: 53323904                 and z19.d, z18.d, z25.d
   272: 54962c04                 lsr z20.b, z18.b, #4
   276: f5323305                 tbl z21.b, { z23.b }, z19.b
   280: 16333405                 tbl z22.b, { z24.b }, z20.b
   284: d03a3504                 eor3 z16.d, z16.d, z21.d, z22.d
   288: 0f020a91                 add x15, x16, #0x280
   292: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   296: 0f020b91                 add x15, x16, #0x2c0
   300: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   304: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   308: 53323904                 and z19.d, z18.d, z25.d
   312: 54962c04                 lsr z20.b, z18.b, #4
   316: f5323305                 tbl z21.b, { z23.b }, z19.b
   320: 16333405                 tbl z22.b, { z24.b }, z20.b
   324: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   328: 0a0d40f9                 ldr x10, [x8, #0x18]
   332: 4a010b8b                 add x10, x10, x11
   336: 0f020c91                 add x15, x16, #0x300
   340: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   344: 0f020d91                 add x15, x16, #0x340
   348: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   352: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   356: 53323904                 and z19.d, z18.d, z25.d
   360: 54962c04                 lsr z20.b, z18.b, #4
   364: f5323305                 tbl z21.b, { z23.b }, z19.b
   368: 16333405                 tbl z22.b, { z24.b }, z20.b
   372: d03a3504                 eor3 z16.d, z16.d, z21.d, z22.d
   376: 0f020e91                 add x15, x16, #0x380
   380: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   384: 0f020f91                 add x15, x16, #0x3c0
   388: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   392: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   396: 53323904                 and z19.d, z18.d, z25.d
   400: 54962c04                 lsr z20.b, z18.b, #4
   404: f5323305                 tbl z21.b, { z23.b }, z19.b
   408: 16333405                 tbl z22.b, { z24.b }, z20.b
   412: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   416: 0a1140f9                 ldr x10, [x8, #0x20]
   420: 4a010b8b                 add x10, x10, x11
   424: 0f021091                 add x15, x16, #0x400
   428: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   432: 0f021191                 add x15, x16, #0x440
   436: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   440: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   444: 53323904                 and z19.d, z18.d, z25.d
   448: 54962c04                 lsr z20.b, z18.b, #4
   452: f5323305                 tbl z21.b, { z23.b }, z19.b
   456: 16333405                 tbl z22.b, { z24.b }, z20.b
   460: d03a3504                 eor3 z16.d, z16.d, z21.d, z22.d
   464: 0f021291                 add x15, x16, #0x480
   468: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   472: 0f021391                 add x15, x16, #0x4c0
   476: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   480: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   484: 53323904                 and z19.d, z18.d, z25.d
   488: 54962c04                 lsr z20.b, z18.b, #4
   492: f5323305                 tbl z21.b, { z23.b }, z19.b
   496: 16333405                 tbl z22.b, { z24.b }, z20.b
   500: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   504: 0a1540f9                 ldr x10, [x8, #0x28]
   508: 4a010b8b                 add x10, x10, x11
   512: 0f021491                 add x15, x16, #0x500
   516: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   520: 0f021591                 add x15, x16, #0x540
   524: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   528: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   532: 53323904                 and z19.d, z18.d, z25.d
   536: 54962c04                 lsr z20.b, z18.b, #4
   540: f5323305                 tbl z21.b, { z23.b }, z19.b
   544: 16333405                 tbl z22.b, { z24.b }, z20.b
   548: d03a3504                 eor3 z16.d, z16.d, z21.d, z22.d
   552: 0f021491                 add x15, x16, #0x500
   556: f72100a4                 ld1rqb { z23.b }, p0/z, [x15]
   560: 0f021591                 add x15, x16, #0x540
   564: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   568: 52a100a4                 ld1b { z18.b }, p0/z, [x10]
   572: 53323904                 and z19.d, z18.d, z25.d
   576: 54962c04                 lsr z20.b, z18.b, #4
   580: f5323305                 tbl z21.b, { z23.b }, z19.b
   584: 16333405                 tbl z22.b, { z24.b }, z20.b
   588: d13a3504                 eor3 z17.d, z17.d, z21.d, z22.d
   592: 2a0140f9                 ldr x10, [x9]
   596: 4a010b8b                 add x10, x10, x11
   600: 50e100e4                 st1b { z16.b }, p0, [x10]
   604: 2a0540f9                 ldr x10, [x9, #8]
   608: 4a010b8b                 add x10, x10, x11
   612: 51e100e4                 st1b { z17.b }, p0, [x10]
   616: 6b010d8b                 add x11, x11, x13
   620: ce0500f1                 subs x14, x14, #1
   624: c1edff54                 b.ne 0x28
   628: c0035fd6                 ret 
   632: 00000000                 udf #0
   636: 00000000                 udf #0
