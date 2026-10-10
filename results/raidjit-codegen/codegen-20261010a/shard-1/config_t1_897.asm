; config_t1_897 target=sve2 twin=baseline topo=raid6(8)/enc op=encode U=8
; code_size=1664 twins=["pipe"] uops_per_byte=3.84375
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0330010                 adr x16, 0x680
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: 8df57ed3                 lsl x13, x12, #2
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: e6c13825                 mov z6.b, #0xf
    36: ae2200b4                 cbz x14, 0x478
    40: 0a0140f9                 ldr x10, [x8]
    44: 4a010b8b                 add x10, x10, x11
    48: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
    52: 5da101a4                 ld1b { z29.b }, p0/z, [x10, #1, mul vl]
    56: 5ea102a4                 ld1b { z30.b }, p0/z, [x10, #2, mul vl]
    60: 5fa103a4                 ld1b { z31.b }, p0/z, [x10, #3, mul vl]
    64: 90337c04                 mov z16.d, z28.d
    68: b1337d04                 mov z17.d, z29.d
    72: d2337e04                 mov z18.d, z30.d
    76: f3337f04                 mov z19.d, z31.d
    80: 0f020091                 add x15, x16, #0
    84: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
    88: 0f020191                 add x15, x16, #0x40
    92: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
    96: 80332604                 and z0.d, z28.d, z6.d
   100: 81972c04                 lsr z1.b, z28.b, #4
   104: 82302005                 tbl z2.b, { z4.b }, z0.b
   108: a3302105                 tbl z3.b, { z5.b }, z1.b
   112: 5430a304                 eor z20.d, z2.d, z3.d
   116: a0332604                 and z0.d, z29.d, z6.d
   120: a1972c04                 lsr z1.b, z29.b, #4
   124: 82302005                 tbl z2.b, { z4.b }, z0.b
   128: a3302105                 tbl z3.b, { z5.b }, z1.b
   132: 5530a304                 eor z21.d, z2.d, z3.d
   136: c0332604                 and z0.d, z30.d, z6.d
   140: c1972c04                 lsr z1.b, z30.b, #4
   144: 82302005                 tbl z2.b, { z4.b }, z0.b
   148: a3302105                 tbl z3.b, { z5.b }, z1.b
   152: 5630a304                 eor z22.d, z2.d, z3.d
   156: e0332604                 and z0.d, z31.d, z6.d
   160: e1972c04                 lsr z1.b, z31.b, #4
   164: 82302005                 tbl z2.b, { z4.b }, z0.b
   168: a3302105                 tbl z3.b, { z5.b }, z1.b
   172: 5730a304                 eor z23.d, z2.d, z3.d
   176: 0a0540f9                 ldr x10, [x8, #8]
   180: 4a010b8b                 add x10, x10, x11
   184: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   188: 59a101a4                 ld1b { z25.b }, p0/z, [x10, #1, mul vl]
   192: 5aa102a4                 ld1b { z26.b }, p0/z, [x10, #2, mul vl]
   196: 5ba103a4                 ld1b { z27.b }, p0/z, [x10, #3, mul vl]
   200: 10031904                 eor z16.b, p0/m, z16.b, z24.b
   204: 31031904                 eor z17.b, p0/m, z17.b, z25.b
   208: 52031904                 eor z18.b, p0/m, z18.b, z26.b
   212: 73031904                 eor z19.b, p0/m, z19.b, z27.b
   216: 0f020291                 add x15, x16, #0x80
   220: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   224: 0f020391                 add x15, x16, #0xc0
   228: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   232: 00332604                 and z0.d, z24.d, z6.d
   236: 01972c04                 lsr z1.b, z24.b, #4
   240: 82302005                 tbl z2.b, { z4.b }, z0.b
   244: a3302105                 tbl z3.b, { z5.b }, z1.b
   248: 74382204                 eor3 z20.d, z20.d, z2.d, z3.d
   252: 20332604                 and z0.d, z25.d, z6.d
   256: 21972c04                 lsr z1.b, z25.b, #4
   260: 82302005                 tbl z2.b, { z4.b }, z0.b
   264: a3302105                 tbl z3.b, { z5.b }, z1.b
   268: 75382204                 eor3 z21.d, z21.d, z2.d, z3.d
   272: 40332604                 and z0.d, z26.d, z6.d
   276: 41972c04                 lsr z1.b, z26.b, #4
   280: 82302005                 tbl z2.b, { z4.b }, z0.b
   284: a3302105                 tbl z3.b, { z5.b }, z1.b
   288: 76382204                 eor3 z22.d, z22.d, z2.d, z3.d
   292: 60332604                 and z0.d, z27.d, z6.d
   296: 61972c04                 lsr z1.b, z27.b, #4
   300: 82302005                 tbl z2.b, { z4.b }, z0.b
   304: a3302105                 tbl z3.b, { z5.b }, z1.b
   308: 77382204                 eor3 z23.d, z23.d, z2.d, z3.d
   312: 0a0940f9                 ldr x10, [x8, #0x10]
   316: 4a010b8b                 add x10, x10, x11
   320: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   324: 59a101a4                 ld1b { z25.b }, p0/z, [x10, #1, mul vl]
   328: 5aa102a4                 ld1b { z26.b }, p0/z, [x10, #2, mul vl]
   332: 5ba103a4                 ld1b { z27.b }, p0/z, [x10, #3, mul vl]
   336: 10031904                 eor z16.b, p0/m, z16.b, z24.b
   340: 31031904                 eor z17.b, p0/m, z17.b, z25.b
   344: 52031904                 eor z18.b, p0/m, z18.b, z26.b
   348: 73031904                 eor z19.b, p0/m, z19.b, z27.b
   352: 0f020491                 add x15, x16, #0x100
   356: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   360: 0f020591                 add x15, x16, #0x140
   364: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   368: 00332604                 and z0.d, z24.d, z6.d
   372: 01972c04                 lsr z1.b, z24.b, #4
   376: 82302005                 tbl z2.b, { z4.b }, z0.b
   380: a3302105                 tbl z3.b, { z5.b }, z1.b
   384: 74382204                 eor3 z20.d, z20.d, z2.d, z3.d
   388: 20332604                 and z0.d, z25.d, z6.d
   392: 21972c04                 lsr z1.b, z25.b, #4
   396: 82302005                 tbl z2.b, { z4.b }, z0.b
   400: a3302105                 tbl z3.b, { z5.b }, z1.b
   404: 75382204                 eor3 z21.d, z21.d, z2.d, z3.d
   408: 40332604                 and z0.d, z26.d, z6.d
   412: 41972c04                 lsr z1.b, z26.b, #4
   416: 82302005                 tbl z2.b, { z4.b }, z0.b
   420: a3302105                 tbl z3.b, { z5.b }, z1.b
   424: 76382204                 eor3 z22.d, z22.d, z2.d, z3.d
   428: 60332604                 and z0.d, z27.d, z6.d
   432: 61972c04                 lsr z1.b, z27.b, #4
   436: 82302005                 tbl z2.b, { z4.b }, z0.b
   440: a3302105                 tbl z3.b, { z5.b }, z1.b
   444: 77382204                 eor3 z23.d, z23.d, z2.d, z3.d
   448: 0a0d40f9                 ldr x10, [x8, #0x18]
   452: 4a010b8b                 add x10, x10, x11
   456: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   460: 59a101a4                 ld1b { z25.b }, p0/z, [x10, #1, mul vl]
   464: 5aa102a4                 ld1b { z26.b }, p0/z, [x10, #2, mul vl]
   468: 5ba103a4                 ld1b { z27.b }, p0/z, [x10, #3, mul vl]
   472: 10031904                 eor z16.b, p0/m, z16.b, z24.b
   476: 31031904                 eor z17.b, p0/m, z17.b, z25.b
   480: 52031904                 eor z18.b, p0/m, z18.b, z26.b
   484: 73031904                 eor z19.b, p0/m, z19.b, z27.b
   488: 0f020691                 add x15, x16, #0x180
   492: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   496: 0f020791                 add x15, x16, #0x1c0
   500: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   504: 00332604                 and z0.d, z24.d, z6.d
   508: 01972c04                 lsr z1.b, z24.b, #4
   512: 82302005                 tbl z2.b, { z4.b }, z0.b
   516: a3302105                 tbl z3.b, { z5.b }, z1.b
   520: 74382204                 eor3 z20.d, z20.d, z2.d, z3.d
   524: 20332604                 and z0.d, z25.d, z6.d
   528: 21972c04                 lsr z1.b, z25.b, #4
   532: 82302005                 tbl z2.b, { z4.b }, z0.b
   536: a3302105                 tbl z3.b, { z5.b }, z1.b
   540: 75382204                 eor3 z21.d, z21.d, z2.d, z3.d
   544: 40332604                 and z0.d, z26.d, z6.d
   548: 41972c04                 lsr z1.b, z26.b, #4
   552: 82302005                 tbl z2.b, { z4.b }, z0.b
   556: a3302105                 tbl z3.b, { z5.b }, z1.b
   560: 76382204                 eor3 z22.d, z22.d, z2.d, z3.d
   564: 60332604                 and z0.d, z27.d, z6.d
   568: 61972c04                 lsr z1.b, z27.b, #4
   572: 82302005                 tbl z2.b, { z4.b }, z0.b
   576: a3302105                 tbl z3.b, { z5.b }, z1.b
   580: 77382204                 eor3 z23.d, z23.d, z2.d, z3.d
   584: 0a1140f9                 ldr x10, [x8, #0x20]
   588: 4a010b8b                 add x10, x10, x11
   592: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   596: 59a101a4                 ld1b { z25.b }, p0/z, [x10, #1, mul vl]
   600: 5aa102a4                 ld1b { z26.b }, p0/z, [x10, #2, mul vl]
   604: 5ba103a4                 ld1b { z27.b }, p0/z, [x10, #3, mul vl]
   608: 10031904                 eor z16.b, p0/m, z16.b, z24.b
   612: 31031904                 eor z17.b, p0/m, z17.b, z25.b
   616: 52031904                 eor z18.b, p0/m, z18.b, z26.b
   620: 73031904                 eor z19.b, p0/m, z19.b, z27.b
   624: 0f020891                 add x15, x16, #0x200
   628: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   632: 0f020991                 add x15, x16, #0x240
   636: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   640: 00332604                 and z0.d, z24.d, z6.d
   644: 01972c04                 lsr z1.b, z24.b, #4
   648: 82302005                 tbl z2.b, { z4.b }, z0.b
   652: a3302105                 tbl z3.b, { z5.b }, z1.b
   656: 74382204                 eor3 z20.d, z20.d, z2.d, z3.d
   660: 20332604                 and z0.d, z25.d, z6.d
   664: 21972c04                 lsr z1.b, z25.b, #4
   668: 82302005                 tbl z2.b, { z4.b }, z0.b
   672: a3302105                 tbl z3.b, { z5.b }, z1.b
   676: 75382204                 eor3 z21.d, z21.d, z2.d, z3.d
   680: 40332604                 and z0.d, z26.d, z6.d
   684: 41972c04                 lsr z1.b, z26.b, #4
   688: 82302005                 tbl z2.b, { z4.b }, z0.b
   692: a3302105                 tbl z3.b, { z5.b }, z1.b
   696: 76382204                 eor3 z22.d, z22.d, z2.d, z3.d
   700: 60332604                 and z0.d, z27.d, z6.d
   704: 61972c04                 lsr z1.b, z27.b, #4
   708: 82302005                 tbl z2.b, { z4.b }, z0.b
   712: a3302105                 tbl z3.b, { z5.b }, z1.b
   716: 77382204                 eor3 z23.d, z23.d, z2.d, z3.d
   720: 0a1540f9                 ldr x10, [x8, #0x28]
   724: 4a010b8b                 add x10, x10, x11
   728: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   732: 59a101a4                 ld1b { z25.b }, p0/z, [x10, #1, mul vl]
   736: 5aa102a4                 ld1b { z26.b }, p0/z, [x10, #2, mul vl]
   740: 5ba103a4                 ld1b { z27.b }, p0/z, [x10, #3, mul vl]
   744: 10031904                 eor z16.b, p0/m, z16.b, z24.b
   748: 31031904                 eor z17.b, p0/m, z17.b, z25.b
   752: 52031904                 eor z18.b, p0/m, z18.b, z26.b
   756: 73031904                 eor z19.b, p0/m, z19.b, z27.b
   760: 0f020a91                 add x15, x16, #0x280
   764: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   768: 0f020b91                 add x15, x16, #0x2c0
   772: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   776: 00332604                 and z0.d, z24.d, z6.d
   780: 01972c04                 lsr z1.b, z24.b, #4
   784: 82302005                 tbl z2.b, { z4.b }, z0.b
   788: a3302105                 tbl z3.b, { z5.b }, z1.b
   792: 74382204                 eor3 z20.d, z20.d, z2.d, z3.d
   796: 20332604                 and z0.d, z25.d, z6.d
   800: 21972c04                 lsr z1.b, z25.b, #4
   804: 82302005                 tbl z2.b, { z4.b }, z0.b
   808: a3302105                 tbl z3.b, { z5.b }, z1.b
   812: 75382204                 eor3 z21.d, z21.d, z2.d, z3.d
   816: 40332604                 and z0.d, z26.d, z6.d
   820: 41972c04                 lsr z1.b, z26.b, #4
   824: 82302005                 tbl z2.b, { z4.b }, z0.b
   828: a3302105                 tbl z3.b, { z5.b }, z1.b
   832: 76382204                 eor3 z22.d, z22.d, z2.d, z3.d
   836: 60332604                 and z0.d, z27.d, z6.d
   840: 61972c04                 lsr z1.b, z27.b, #4
   844: 82302005                 tbl z2.b, { z4.b }, z0.b
   848: a3302105                 tbl z3.b, { z5.b }, z1.b
   852: 77382204                 eor3 z23.d, z23.d, z2.d, z3.d
   856: 0a1940f9                 ldr x10, [x8, #0x30]
   860: 4a010b8b                 add x10, x10, x11
   864: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   868: 59a101a4                 ld1b { z25.b }, p0/z, [x10, #1, mul vl]
   872: 5aa102a4                 ld1b { z26.b }, p0/z, [x10, #2, mul vl]
   876: 5ba103a4                 ld1b { z27.b }, p0/z, [x10, #3, mul vl]
   880: 10031904                 eor z16.b, p0/m, z16.b, z24.b
   884: 31031904                 eor z17.b, p0/m, z17.b, z25.b
   888: 52031904                 eor z18.b, p0/m, z18.b, z26.b
   892: 73031904                 eor z19.b, p0/m, z19.b, z27.b
   896: 0f020c91                 add x15, x16, #0x300
   900: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   904: 0f020d91                 add x15, x16, #0x340
   908: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   912: 00332604                 and z0.d, z24.d, z6.d
   916: 01972c04                 lsr z1.b, z24.b, #4
   920: 82302005                 tbl z2.b, { z4.b }, z0.b
   924: a3302105                 tbl z3.b, { z5.b }, z1.b
   928: 74382204                 eor3 z20.d, z20.d, z2.d, z3.d
   932: 20332604                 and z0.d, z25.d, z6.d
   936: 21972c04                 lsr z1.b, z25.b, #4
   940: 82302005                 tbl z2.b, { z4.b }, z0.b
   944: a3302105                 tbl z3.b, { z5.b }, z1.b
   948: 75382204                 eor3 z21.d, z21.d, z2.d, z3.d
   952: 40332604                 and z0.d, z26.d, z6.d
   956: 41972c04                 lsr z1.b, z26.b, #4
   960: 82302005                 tbl z2.b, { z4.b }, z0.b
   964: a3302105                 tbl z3.b, { z5.b }, z1.b
   968: 76382204                 eor3 z22.d, z22.d, z2.d, z3.d
   972: 60332604                 and z0.d, z27.d, z6.d
   976: 61972c04                 lsr z1.b, z27.b, #4
   980: 82302005                 tbl z2.b, { z4.b }, z0.b
   984: a3302105                 tbl z3.b, { z5.b }, z1.b
   988: 77382204                 eor3 z23.d, z23.d, z2.d, z3.d
   992: 0a1d40f9                 ldr x10, [x8, #0x38]
   996: 4a010b8b                 add x10, x10, x11
  1000: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
  1004: 59a101a4                 ld1b { z25.b }, p0/z, [x10, #1, mul vl]
  1008: 5aa102a4                 ld1b { z26.b }, p0/z, [x10, #2, mul vl]
  1012: 5ba103a4                 ld1b { z27.b }, p0/z, [x10, #3, mul vl]
  1016: 10031904                 eor z16.b, p0/m, z16.b, z24.b
  1020: 31031904                 eor z17.b, p0/m, z17.b, z25.b
  1024: 52031904                 eor z18.b, p0/m, z18.b, z26.b
  1028: 73031904                 eor z19.b, p0/m, z19.b, z27.b
  1032: 14031904                 eor z20.b, p0/m, z20.b, z24.b
  1036: 35031904                 eor z21.b, p0/m, z21.b, z25.b
  1040: 56031904                 eor z22.b, p0/m, z22.b, z26.b
  1044: 77031904                 eor z23.b, p0/m, z23.b, z27.b
  1048: df0500f1                 cmp x14, #1
  1052: 0d010054                 b.le 0x43c
  1056: 0a0140f9                 ldr x10, [x8]
  1060: 4a010b8b                 add x10, x10, x11
  1064: 4a010191                 add x10, x10, #0x40
  1068: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
  1072: 5da101a4                 ld1b { z29.b }, p0/z, [x10, #1, mul vl]
  1076: 5ea102a4                 ld1b { z30.b }, p0/z, [x10, #2, mul vl]
  1080: 5fa103a4                 ld1b { z31.b }, p0/z, [x10, #3, mul vl]
  1084: 2a0140f9                 ldr x10, [x9]
  1088: 4a010b8b                 add x10, x10, x11
  1092: 50e100e4                 st1b { z16.b }, p0, [x10]
  1096: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
  1100: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
  1104: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
  1108: 2a0540f9                 ldr x10, [x9, #8]
  1112: 4a010b8b                 add x10, x10, x11
  1116: 54e100e4                 st1b { z20.b }, p0, [x10]
  1120: 55e101e4                 st1b { z21.b }, p0, [x10, #1, mul vl]
  1124: 56e102e4                 st1b { z22.b }, p0, [x10, #2, mul vl]
  1128: 57e103e4                 st1b { z23.b }, p0, [x10, #3, mul vl]
  1132: 6b010d8b                 add x11, x11, x13
  1136: ce0500f1                 subs x14, x14, #1
  1140: 61deff54                 b.ne 0x40
  1144: 7f0102eb                 cmp x11, x2
  1148: 420e0054                 b.hs 0x644
  1152: 0f020e91                 add x15, x16, #0x380
  1156: 11021291                 add x17, x16, #0x480
  1160: 2a0140f9                 ldr x10, [x9]
  1164: ed030baa                 mov x13, x11
  1168: 0c008052                 mov w12, #0
  1172: 0e0140f9                 ldr x14, [x8]
  1176: ce696d38                 ldrb w14, [x14, x13]
  1180: 8c010e4a                 eor w12, w12, w14
  1184: 0e0540f9                 ldr x14, [x8, #8]
  1188: ce696d38                 ldrb w14, [x14, x13]
  1192: 8c010e4a                 eor w12, w12, w14
  1196: 0e0940f9                 ldr x14, [x8, #0x10]
  1200: ce696d38                 ldrb w14, [x14, x13]
  1204: 8c010e4a                 eor w12, w12, w14
  1208: 0e0d40f9                 ldr x14, [x8, #0x18]
  1212: ce696d38                 ldrb w14, [x14, x13]
  1216: 8c010e4a                 eor w12, w12, w14
  1220: 0e1140f9                 ldr x14, [x8, #0x20]
  1224: ce696d38                 ldrb w14, [x14, x13]
  1228: 8c010e4a                 eor w12, w12, w14
  1232: 0e1540f9                 ldr x14, [x8, #0x28]
  1236: ce696d38                 ldrb w14, [x14, x13]
  1240: 8c010e4a                 eor w12, w12, w14
  1244: 0e1940f9                 ldr x14, [x8, #0x30]
  1248: ce696d38                 ldrb w14, [x14, x13]
  1252: 8c010e4a                 eor w12, w12, w14
  1256: 0e1d40f9                 ldr x14, [x8, #0x38]
  1260: ce696d38                 ldrb w14, [x14, x13]
  1264: 8c010e4a                 eor w12, w12, w14
  1268: 4c692d38                 strb w12, [x10, x13]
  1272: ad050091                 add x13, x13, #1
  1276: bf0102eb                 cmp x13, x2
  1280: 83fcff54                 b.lo 0x490
  1284: 2a0540f9                 ldr x10, [x9, #8]
  1288: ed030baa                 mov x13, x11
  1292: 0c008052                 mov w12, #0
  1296: 0e0140f9                 ldr x14, [x8]
  1300: ce696d38                 ldrb w14, [x14, x13]
  1304: 0e010034                 cbz w14, 0x538
  1308: ee696e38                 ldrb w14, [x15, x14]
  1312: ce990311                 add w14, w14, #0xe6
  1316: dffd0371                 cmp w14, #0xff
  1320: 43000054                 b.lo 0x530
  1324: cefd0351                 sub w14, w14, #0xff
  1328: 2e6a6e38                 ldrb w14, [x17, x14]
  1332: 8c010e4a                 eor w12, w12, w14
  1336: 0e0540f9                 ldr x14, [x8, #8]
  1340: ce696d38                 ldrb w14, [x14, x13]
  1344: 0e010034                 cbz w14, 0x560
  1348: ee696e38                 ldrb w14, [x15, x14]
  1352: ce350311                 add w14, w14, #0xcd
  1356: dffd0371                 cmp w14, #0xff
  1360: 43000054                 b.lo 0x558
  1364: cefd0351                 sub w14, w14, #0xff
  1368: 2e6a6e38                 ldrb w14, [x17, x14]
  1372: 8c010e4a                 eor w12, w12, w14
  1376: 0e0940f9                 ldr x14, [x8, #0x10]
  1380: ce696d38                 ldrb w14, [x14, x13]
  1384: 0e010034                 cbz w14, 0x588
  1388: ee696e38                 ldrb w14, [x15, x14]
  1392: ce9d0311                 add w14, w14, #0xe7
  1396: dffd0371                 cmp w14, #0xff
  1400: 43000054                 b.lo 0x580
  1404: cefd0351                 sub w14, w14, #0xff
  1408: 2e6a6e38                 ldrb w14, [x17, x14]
  1412: 8c010e4a                 eor w12, w12, w14
  1416: 0e0d40f9                 ldr x14, [x8, #0x18]
  1420: ce696d38                 ldrb w14, [x14, x13]
  1424: 0e010034                 cbz w14, 0x5b0
  1428: ee696e38                 ldrb w14, [x15, x14]
  1432: ce390311                 add w14, w14, #0xce
  1436: dffd0371                 cmp w14, #0xff
  1440: 43000054                 b.lo 0x5a8
  1444: cefd0351                 sub w14, w14, #0xff
  1448: 2e6a6e38                 ldrb w14, [x17, x14]
  1452: 8c010e4a                 eor w12, w12, w14
  1456: 0e1140f9                 ldr x14, [x8, #0x20]
  1460: ce696d38                 ldrb w14, [x14, x13]
  1464: 0e010034                 cbz w14, 0x5d8
  1468: ee696e38                 ldrb w14, [x15, x14]
  1472: ce2d0111                 add w14, w14, #0x4b
  1476: dffd0371                 cmp w14, #0xff
  1480: 43000054                 b.lo 0x5d0
  1484: cefd0351                 sub w14, w14, #0xff
  1488: 2e6a6e38                 ldrb w14, [x17, x14]
  1492: 8c010e4a                 eor w12, w12, w14
  1496: 0e1540f9                 ldr x14, [x8, #0x28]
  1500: ce696d38                 ldrb w14, [x14, x13]
  1504: 0e010034                 cbz w14, 0x600
  1508: ee696e38                 ldrb w14, [x15, x14]
  1512: cec90011                 add w14, w14, #0x32
  1516: dffd0371                 cmp w14, #0xff
  1520: 43000054                 b.lo 0x5f8
  1524: cefd0351                 sub w14, w14, #0xff
  1528: 2e6a6e38                 ldrb w14, [x17, x14]
  1532: 8c010e4a                 eor w12, w12, w14
  1536: 0e1940f9                 ldr x14, [x8, #0x30]
  1540: ce696d38                 ldrb w14, [x14, x13]
  1544: 0e010034                 cbz w14, 0x628
  1548: ee696e38                 ldrb w14, [x15, x14]
  1552: ce650011                 add w14, w14, #0x19
  1556: dffd0371                 cmp w14, #0xff
  1560: 43000054                 b.lo 0x620
  1564: cefd0351                 sub w14, w14, #0xff
  1568: 2e6a6e38                 ldrb w14, [x17, x14]
  1572: 8c010e4a                 eor w12, w12, w14
  1576: 0e1d40f9                 ldr x14, [x8, #0x38]
  1580: ce696d38                 ldrb w14, [x14, x13]
  1584: 8c010e4a                 eor w12, w12, w14
  1588: 4c692d38                 strb w12, [x10, x13]
  1592: ad050091                 add x13, x13, #1
  1596: bf0102eb                 cmp x13, x2
  1600: 63f6ff54                 b.lo 0x50c
  1604: c0035fd6                 ret 
  1608: 00000000                 udf #0
  1612: 00000000                 udf #0
  1616: 00000000                 udf #0
  1620: 00000000                 udf #0
  1624: 00000000                 udf #0
  1628: 00000000                 udf #0
  1632: 00000000                 udf #0
  1636: 00000000                 udf #0
  1640: 00000000                 udf #0
  1644: 00000000                 udf #0
  1648: 00000000                 udf #0
  1652: 00000000                 udf #0
  1656: 00000000                 udf #0
  1660: 00000000                 udf #0
