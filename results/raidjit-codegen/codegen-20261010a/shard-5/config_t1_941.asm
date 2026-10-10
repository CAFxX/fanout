; config_t1_941 target=sve2 twin=predtail topo=rs(6,3)/enc op=encode U=1
; code_size=1216 twins=["predtail"] uops_per_byte=7.5
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0250010                 adr x16, 0x4c0
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ed030caa                 mov x13, x12
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: fac13825                 mov z26.b, #0xf
    36: 6e1200b4                 cbz x14, 0x270
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
    76: 70327304                 mov z16.d, z19.d
    80: 0f020091                 add x15, x16, #0
    84: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
    88: 0f020191                 add x15, x16, #0x40
    92: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
    96: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   100: 74323a04                 and z20.d, z19.d, z26.d
   104: 75962c04                 lsr z21.b, z19.b, #4
   108: 16333405                 tbl z22.b, { z24.b }, z20.b
   112: 37333505                 tbl z23.b, { z25.b }, z21.b
   116: d132b704                 eor z17.d, z22.d, z23.d
   120: 0f020291                 add x15, x16, #0x80
   124: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   128: 0f020391                 add x15, x16, #0xc0
   132: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   136: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   140: 74323a04                 and z20.d, z19.d, z26.d
   144: 75962c04                 lsr z21.b, z19.b, #4
   148: 16333405                 tbl z22.b, { z24.b }, z20.b
   152: 37333505                 tbl z23.b, { z25.b }, z21.b
   156: d232b704                 eor z18.d, z22.d, z23.d
   160: 0a0540f9                 ldr x10, [x8, #8]
   164: 4a010b8b                 add x10, x10, x11
   168: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   172: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   176: 0f020291                 add x15, x16, #0x80
   180: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   184: 0f020391                 add x15, x16, #0xc0
   188: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   192: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   196: 74323a04                 and z20.d, z19.d, z26.d
   200: 75962c04                 lsr z21.b, z19.b, #4
   204: 16333405                 tbl z22.b, { z24.b }, z20.b
   208: 37333505                 tbl z23.b, { z25.b }, z21.b
   212: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   216: 0f020491                 add x15, x16, #0x100
   220: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   224: 0f020591                 add x15, x16, #0x140
   228: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   232: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   236: 74323a04                 and z20.d, z19.d, z26.d
   240: 75962c04                 lsr z21.b, z19.b, #4
   244: 16333405                 tbl z22.b, { z24.b }, z20.b
   248: 37333505                 tbl z23.b, { z25.b }, z21.b
   252: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   256: 0a0940f9                 ldr x10, [x8, #0x10]
   260: 4a010b8b                 add x10, x10, x11
   264: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   268: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   272: 0f020491                 add x15, x16, #0x100
   276: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   280: 0f020591                 add x15, x16, #0x140
   284: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   288: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   292: 74323a04                 and z20.d, z19.d, z26.d
   296: 75962c04                 lsr z21.b, z19.b, #4
   300: 16333405                 tbl z22.b, { z24.b }, z20.b
   304: 37333505                 tbl z23.b, { z25.b }, z21.b
   308: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   312: 0f020691                 add x15, x16, #0x180
   316: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   320: 0f020791                 add x15, x16, #0x1c0
   324: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   328: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   332: 74323a04                 and z20.d, z19.d, z26.d
   336: 75962c04                 lsr z21.b, z19.b, #4
   340: 16333405                 tbl z22.b, { z24.b }, z20.b
   344: 37333505                 tbl z23.b, { z25.b }, z21.b
   348: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   352: 0a0d40f9                 ldr x10, [x8, #0x18]
   356: 4a010b8b                 add x10, x10, x11
   360: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   364: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   368: 0f020891                 add x15, x16, #0x200
   372: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   376: 0f020991                 add x15, x16, #0x240
   380: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   384: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   388: 74323a04                 and z20.d, z19.d, z26.d
   392: 75962c04                 lsr z21.b, z19.b, #4
   396: 16333405                 tbl z22.b, { z24.b }, z20.b
   400: 37333505                 tbl z23.b, { z25.b }, z21.b
   404: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   408: 0f020091                 add x15, x16, #0
   412: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   416: 0f020191                 add x15, x16, #0x40
   420: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   424: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   428: 74323a04                 and z20.d, z19.d, z26.d
   432: 75962c04                 lsr z21.b, z19.b, #4
   436: 16333405                 tbl z22.b, { z24.b }, z20.b
   440: 37333505                 tbl z23.b, { z25.b }, z21.b
   444: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   448: 0a1140f9                 ldr x10, [x8, #0x20]
   452: 4a010b8b                 add x10, x10, x11
   456: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   460: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   464: 0f020691                 add x15, x16, #0x180
   468: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   472: 0f020791                 add x15, x16, #0x1c0
   476: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   480: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   484: 74323a04                 and z20.d, z19.d, z26.d
   488: 75962c04                 lsr z21.b, z19.b, #4
   492: 16333405                 tbl z22.b, { z24.b }, z20.b
   496: 37333505                 tbl z23.b, { z25.b }, z21.b
   500: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   504: 0f020891                 add x15, x16, #0x200
   508: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   512: 0f020991                 add x15, x16, #0x240
   516: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   520: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   524: 74323a04                 and z20.d, z19.d, z26.d
   528: 75962c04                 lsr z21.b, z19.b, #4
   532: 16333405                 tbl z22.b, { z24.b }, z20.b
   536: 37333505                 tbl z23.b, { z25.b }, z21.b
   540: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   544: 0a1540f9                 ldr x10, [x8, #0x28]
   548: 4a010b8b                 add x10, x10, x11
   552: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   556: 70021904                 eor z16.b, p0/m, z16.b, z19.b
   560: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   564: 71021904                 eor z17.b, p0/m, z17.b, z19.b
   568: 53a100a4                 ld1b { z19.b }, p0/z, [x10]
   572: 72021904                 eor z18.b, p0/m, z18.b, z19.b
   576: 2a0140f9                 ldr x10, [x9]
   580: 4a010b8b                 add x10, x10, x11
   584: 50e100e4                 st1b { z16.b }, p0, [x10]
   588: 2a0540f9                 ldr x10, [x9, #8]
   592: 4a010b8b                 add x10, x10, x11
   596: 51e100e4                 st1b { z17.b }, p0, [x10]
   600: 2a0940f9                 ldr x10, [x9, #0x10]
   604: 4a010b8b                 add x10, x10, x11
   608: 52e100e4                 st1b { z18.b }, p0, [x10]
   612: 6b010d8b                 add x11, x11, x13
   616: ce0500f1                 subs x14, x14, #1
   620: e1edff54                 b.ne 0x28
   624: 7f0102eb                 cmp x11, x2
   628: e2110054                 b.hs 0x4b0
   632: 4d000bcb                 sub x13, x2, x11
   636: e1172d25                 whilelt p1.b, xzr, x13
   640: 0a0140f9                 ldr x10, [x8]
   644: 4a010b8b                 add x10, x10, x11
   648: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   652: 70327304                 mov z16.d, z19.d
   656: 0f020091                 add x15, x16, #0
   660: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   664: 0f020191                 add x15, x16, #0x40
   668: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   672: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   676: 74323a04                 and z20.d, z19.d, z26.d
   680: 75962c04                 lsr z21.b, z19.b, #4
   684: 16333405                 tbl z22.b, { z24.b }, z20.b
   688: 37333505                 tbl z23.b, { z25.b }, z21.b
   692: d132b704                 eor z17.d, z22.d, z23.d
   696: 0f020291                 add x15, x16, #0x80
   700: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   704: 0f020391                 add x15, x16, #0xc0
   708: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   712: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   716: 74323a04                 and z20.d, z19.d, z26.d
   720: 75962c04                 lsr z21.b, z19.b, #4
   724: 16333405                 tbl z22.b, { z24.b }, z20.b
   728: 37333505                 tbl z23.b, { z25.b }, z21.b
   732: d232b704                 eor z18.d, z22.d, z23.d
   736: 0a0540f9                 ldr x10, [x8, #8]
   740: 4a010b8b                 add x10, x10, x11
   744: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   748: 70061904                 eor z16.b, p1/m, z16.b, z19.b
   752: 0f020291                 add x15, x16, #0x80
   756: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   760: 0f020391                 add x15, x16, #0xc0
   764: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   768: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   772: 74323a04                 and z20.d, z19.d, z26.d
   776: 75962c04                 lsr z21.b, z19.b, #4
   780: 16333405                 tbl z22.b, { z24.b }, z20.b
   784: 37333505                 tbl z23.b, { z25.b }, z21.b
   788: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   792: 0f020491                 add x15, x16, #0x100
   796: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   800: 0f020591                 add x15, x16, #0x140
   804: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   808: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   812: 74323a04                 and z20.d, z19.d, z26.d
   816: 75962c04                 lsr z21.b, z19.b, #4
   820: 16333405                 tbl z22.b, { z24.b }, z20.b
   824: 37333505                 tbl z23.b, { z25.b }, z21.b
   828: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   832: 0a0940f9                 ldr x10, [x8, #0x10]
   836: 4a010b8b                 add x10, x10, x11
   840: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   844: 70061904                 eor z16.b, p1/m, z16.b, z19.b
   848: 0f020491                 add x15, x16, #0x100
   852: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   856: 0f020591                 add x15, x16, #0x140
   860: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   864: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   868: 74323a04                 and z20.d, z19.d, z26.d
   872: 75962c04                 lsr z21.b, z19.b, #4
   876: 16333405                 tbl z22.b, { z24.b }, z20.b
   880: 37333505                 tbl z23.b, { z25.b }, z21.b
   884: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   888: 0f020691                 add x15, x16, #0x180
   892: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   896: 0f020791                 add x15, x16, #0x1c0
   900: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   904: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   908: 74323a04                 and z20.d, z19.d, z26.d
   912: 75962c04                 lsr z21.b, z19.b, #4
   916: 16333405                 tbl z22.b, { z24.b }, z20.b
   920: 37333505                 tbl z23.b, { z25.b }, z21.b
   924: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
   928: 0a0d40f9                 ldr x10, [x8, #0x18]
   932: 4a010b8b                 add x10, x10, x11
   936: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   940: 70061904                 eor z16.b, p1/m, z16.b, z19.b
   944: 0f020891                 add x15, x16, #0x200
   948: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   952: 0f020991                 add x15, x16, #0x240
   956: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   960: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
   964: 74323a04                 and z20.d, z19.d, z26.d
   968: 75962c04                 lsr z21.b, z19.b, #4
   972: 16333405                 tbl z22.b, { z24.b }, z20.b
   976: 37333505                 tbl z23.b, { z25.b }, z21.b
   980: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
   984: 0f020091                 add x15, x16, #0
   988: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
   992: 0f020191                 add x15, x16, #0x40
   996: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
  1000: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
  1004: 74323a04                 and z20.d, z19.d, z26.d
  1008: 75962c04                 lsr z21.b, z19.b, #4
  1012: 16333405                 tbl z22.b, { z24.b }, z20.b
  1016: 37333505                 tbl z23.b, { z25.b }, z21.b
  1020: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
  1024: 0a1140f9                 ldr x10, [x8, #0x20]
  1028: 4a010b8b                 add x10, x10, x11
  1032: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
  1036: 70061904                 eor z16.b, p1/m, z16.b, z19.b
  1040: 0f020691                 add x15, x16, #0x180
  1044: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
  1048: 0f020791                 add x15, x16, #0x1c0
  1052: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
  1056: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
  1060: 74323a04                 and z20.d, z19.d, z26.d
  1064: 75962c04                 lsr z21.b, z19.b, #4
  1068: 16333405                 tbl z22.b, { z24.b }, z20.b
  1072: 37333505                 tbl z23.b, { z25.b }, z21.b
  1076: f13a3604                 eor3 z17.d, z17.d, z22.d, z23.d
  1080: 0f020891                 add x15, x16, #0x200
  1084: f82100a4                 ld1rqb { z24.b }, p0/z, [x15]
  1088: 0f020991                 add x15, x16, #0x240
  1092: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
  1096: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
  1100: 74323a04                 and z20.d, z19.d, z26.d
  1104: 75962c04                 lsr z21.b, z19.b, #4
  1108: 16333405                 tbl z22.b, { z24.b }, z20.b
  1112: 37333505                 tbl z23.b, { z25.b }, z21.b
  1116: f23a3604                 eor3 z18.d, z18.d, z22.d, z23.d
  1120: 0a1540f9                 ldr x10, [x8, #0x28]
  1124: 4a010b8b                 add x10, x10, x11
  1128: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
  1132: 70061904                 eor z16.b, p1/m, z16.b, z19.b
  1136: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
  1140: 71061904                 eor z17.b, p1/m, z17.b, z19.b
  1144: 53a500a4                 ld1b { z19.b }, p1/z, [x10]
  1148: 72061904                 eor z18.b, p1/m, z18.b, z19.b
  1152: 2a0140f9                 ldr x10, [x9]
  1156: 4a010b8b                 add x10, x10, x11
  1160: 50e500e4                 st1b { z16.b }, p1, [x10]
  1164: 2a0540f9                 ldr x10, [x9, #8]
  1168: 4a010b8b                 add x10, x10, x11
  1172: 51e500e4                 st1b { z17.b }, p1, [x10]
  1176: 2a0940f9                 ldr x10, [x9, #0x10]
  1180: 4a010b8b                 add x10, x10, x11
  1184: 52e500e4                 st1b { z18.b }, p1, [x10]
  1188: 6b010c8b                 add x11, x11, x12
  1192: ad010ceb                 subs x13, x13, x12
  1196: 8ceeff54                 b.gt 0x27c
  1200: c0035fd6                 ret 
  1204: 00000000                 udf #0
  1208: 00000000                 udf #0
  1212: 00000000                 udf #0
