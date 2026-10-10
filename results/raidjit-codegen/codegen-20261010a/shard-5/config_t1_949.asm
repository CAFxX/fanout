; config_t1_949 target=sve2 twin=baseline topo=rs(6,3)/enc op=encode U=8
; code_size=1728 twins=["pipe"] uops_per_byte=5.0625
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0350010                 adr x16, 0x6c0
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: 6d0080d2                 mov x13, #3
    24: 8d7d0d9b                 mul x13, x12, x13
    28: 4e08cd9a                 udiv x14, x2, x13
    32: e01f2c25                 whilelo p0.b, xzr, x12
    36: e5c13825                 mov z5.b, #0xf
    40: ee2100b4                 cbz x14, 0x464
    44: 0a0140f9                 ldr x10, [x8]
    48: 4a010b8b                 add x10, x10, x11
    52: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
    56: 5da101a4                 ld1b { z29.b }, p0/z, [x10, #1, mul vl]
    60: 5ea102a4                 ld1b { z30.b }, p0/z, [x10, #2, mul vl]
    64: 90337c04                 mov z16.d, z28.d
    68: b1337d04                 mov z17.d, z29.d
    72: d2337e04                 mov z18.d, z30.d
    76: 0f020091                 add x15, x16, #0
    80: e32100a4                 ld1rqb { z3.b }, p0/z, [x15]
    84: 0f020191                 add x15, x16, #0x40
    88: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
    92: 9f332504                 and z31.d, z28.d, z5.d
    96: 80972c04                 lsr z0.b, z28.b, #4
   100: 61303f05                 tbl z1.b, { z3.b }, z31.b
   104: 82302005                 tbl z2.b, { z4.b }, z0.b
   108: 3330a204                 eor z19.d, z1.d, z2.d
   112: bf332504                 and z31.d, z29.d, z5.d
   116: a0972c04                 lsr z0.b, z29.b, #4
   120: 61303f05                 tbl z1.b, { z3.b }, z31.b
   124: 82302005                 tbl z2.b, { z4.b }, z0.b
   128: 3430a204                 eor z20.d, z1.d, z2.d
   132: df332504                 and z31.d, z30.d, z5.d
   136: c0972c04                 lsr z0.b, z30.b, #4
   140: 61303f05                 tbl z1.b, { z3.b }, z31.b
   144: 82302005                 tbl z2.b, { z4.b }, z0.b
   148: 3530a204                 eor z21.d, z1.d, z2.d
   152: 0f020291                 add x15, x16, #0x80
   156: e32100a4                 ld1rqb { z3.b }, p0/z, [x15]
   160: 0f020391                 add x15, x16, #0xc0
   164: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   168: 9f332504                 and z31.d, z28.d, z5.d
   172: 80972c04                 lsr z0.b, z28.b, #4
   176: 61303f05                 tbl z1.b, { z3.b }, z31.b
   180: 82302005                 tbl z2.b, { z4.b }, z0.b
   184: 3630a204                 eor z22.d, z1.d, z2.d
   188: bf332504                 and z31.d, z29.d, z5.d
   192: a0972c04                 lsr z0.b, z29.b, #4
   196: 61303f05                 tbl z1.b, { z3.b }, z31.b
   200: 82302005                 tbl z2.b, { z4.b }, z0.b
   204: 3730a204                 eor z23.d, z1.d, z2.d
   208: df332504                 and z31.d, z30.d, z5.d
   212: c0972c04                 lsr z0.b, z30.b, #4
   216: 61303f05                 tbl z1.b, { z3.b }, z31.b
   220: 82302005                 tbl z2.b, { z4.b }, z0.b
   224: 3830a204                 eor z24.d, z1.d, z2.d
   228: 0a0540f9                 ldr x10, [x8, #8]
   232: 4a010b8b                 add x10, x10, x11
   236: 59a100a4                 ld1b { z25.b }, p0/z, [x10]
   240: 5aa101a4                 ld1b { z26.b }, p0/z, [x10, #1, mul vl]
   244: 5ba102a4                 ld1b { z27.b }, p0/z, [x10, #2, mul vl]
   248: 30031904                 eor z16.b, p0/m, z16.b, z25.b
   252: 51031904                 eor z17.b, p0/m, z17.b, z26.b
   256: 72031904                 eor z18.b, p0/m, z18.b, z27.b
   260: 0f020291                 add x15, x16, #0x80
   264: e32100a4                 ld1rqb { z3.b }, p0/z, [x15]
   268: 0f020391                 add x15, x16, #0xc0
   272: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   276: 3f332504                 and z31.d, z25.d, z5.d
   280: 20972c04                 lsr z0.b, z25.b, #4
   284: 61303f05                 tbl z1.b, { z3.b }, z31.b
   288: 82302005                 tbl z2.b, { z4.b }, z0.b
   292: 53382104                 eor3 z19.d, z19.d, z1.d, z2.d
   296: 5f332504                 and z31.d, z26.d, z5.d
   300: 40972c04                 lsr z0.b, z26.b, #4
   304: 61303f05                 tbl z1.b, { z3.b }, z31.b
   308: 82302005                 tbl z2.b, { z4.b }, z0.b
   312: 54382104                 eor3 z20.d, z20.d, z1.d, z2.d
   316: 7f332504                 and z31.d, z27.d, z5.d
   320: 60972c04                 lsr z0.b, z27.b, #4
   324: 61303f05                 tbl z1.b, { z3.b }, z31.b
   328: 82302005                 tbl z2.b, { z4.b }, z0.b
   332: 55382104                 eor3 z21.d, z21.d, z1.d, z2.d
   336: 0f020491                 add x15, x16, #0x100
   340: e32100a4                 ld1rqb { z3.b }, p0/z, [x15]
   344: 0f020591                 add x15, x16, #0x140
   348: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   352: 3f332504                 and z31.d, z25.d, z5.d
   356: 20972c04                 lsr z0.b, z25.b, #4
   360: 61303f05                 tbl z1.b, { z3.b }, z31.b
   364: 82302005                 tbl z2.b, { z4.b }, z0.b
   368: 56382104                 eor3 z22.d, z22.d, z1.d, z2.d
   372: 5f332504                 and z31.d, z26.d, z5.d
   376: 40972c04                 lsr z0.b, z26.b, #4
   380: 61303f05                 tbl z1.b, { z3.b }, z31.b
   384: 82302005                 tbl z2.b, { z4.b }, z0.b
   388: 57382104                 eor3 z23.d, z23.d, z1.d, z2.d
   392: 7f332504                 and z31.d, z27.d, z5.d
   396: 60972c04                 lsr z0.b, z27.b, #4
   400: 61303f05                 tbl z1.b, { z3.b }, z31.b
   404: 82302005                 tbl z2.b, { z4.b }, z0.b
   408: 58382104                 eor3 z24.d, z24.d, z1.d, z2.d
   412: 0a0940f9                 ldr x10, [x8, #0x10]
   416: 4a010b8b                 add x10, x10, x11
   420: 59a100a4                 ld1b { z25.b }, p0/z, [x10]
   424: 5aa101a4                 ld1b { z26.b }, p0/z, [x10, #1, mul vl]
   428: 5ba102a4                 ld1b { z27.b }, p0/z, [x10, #2, mul vl]
   432: 30031904                 eor z16.b, p0/m, z16.b, z25.b
   436: 51031904                 eor z17.b, p0/m, z17.b, z26.b
   440: 72031904                 eor z18.b, p0/m, z18.b, z27.b
   444: 0f020491                 add x15, x16, #0x100
   448: e32100a4                 ld1rqb { z3.b }, p0/z, [x15]
   452: 0f020591                 add x15, x16, #0x140
   456: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   460: 3f332504                 and z31.d, z25.d, z5.d
   464: 20972c04                 lsr z0.b, z25.b, #4
   468: 61303f05                 tbl z1.b, { z3.b }, z31.b
   472: 82302005                 tbl z2.b, { z4.b }, z0.b
   476: 53382104                 eor3 z19.d, z19.d, z1.d, z2.d
   480: 5f332504                 and z31.d, z26.d, z5.d
   484: 40972c04                 lsr z0.b, z26.b, #4
   488: 61303f05                 tbl z1.b, { z3.b }, z31.b
   492: 82302005                 tbl z2.b, { z4.b }, z0.b
   496: 54382104                 eor3 z20.d, z20.d, z1.d, z2.d
   500: 7f332504                 and z31.d, z27.d, z5.d
   504: 60972c04                 lsr z0.b, z27.b, #4
   508: 61303f05                 tbl z1.b, { z3.b }, z31.b
   512: 82302005                 tbl z2.b, { z4.b }, z0.b
   516: 55382104                 eor3 z21.d, z21.d, z1.d, z2.d
   520: 0f020691                 add x15, x16, #0x180
   524: e32100a4                 ld1rqb { z3.b }, p0/z, [x15]
   528: 0f020791                 add x15, x16, #0x1c0
   532: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   536: 3f332504                 and z31.d, z25.d, z5.d
   540: 20972c04                 lsr z0.b, z25.b, #4
   544: 61303f05                 tbl z1.b, { z3.b }, z31.b
   548: 82302005                 tbl z2.b, { z4.b }, z0.b
   552: 56382104                 eor3 z22.d, z22.d, z1.d, z2.d
   556: 5f332504                 and z31.d, z26.d, z5.d
   560: 40972c04                 lsr z0.b, z26.b, #4
   564: 61303f05                 tbl z1.b, { z3.b }, z31.b
   568: 82302005                 tbl z2.b, { z4.b }, z0.b
   572: 57382104                 eor3 z23.d, z23.d, z1.d, z2.d
   576: 7f332504                 and z31.d, z27.d, z5.d
   580: 60972c04                 lsr z0.b, z27.b, #4
   584: 61303f05                 tbl z1.b, { z3.b }, z31.b
   588: 82302005                 tbl z2.b, { z4.b }, z0.b
   592: 58382104                 eor3 z24.d, z24.d, z1.d, z2.d
   596: 0a0d40f9                 ldr x10, [x8, #0x18]
   600: 4a010b8b                 add x10, x10, x11
   604: 59a100a4                 ld1b { z25.b }, p0/z, [x10]
   608: 5aa101a4                 ld1b { z26.b }, p0/z, [x10, #1, mul vl]
   612: 5ba102a4                 ld1b { z27.b }, p0/z, [x10, #2, mul vl]
   616: 30031904                 eor z16.b, p0/m, z16.b, z25.b
   620: 51031904                 eor z17.b, p0/m, z17.b, z26.b
   624: 72031904                 eor z18.b, p0/m, z18.b, z27.b
   628: 0f020891                 add x15, x16, #0x200
   632: e32100a4                 ld1rqb { z3.b }, p0/z, [x15]
   636: 0f020991                 add x15, x16, #0x240
   640: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   644: 3f332504                 and z31.d, z25.d, z5.d
   648: 20972c04                 lsr z0.b, z25.b, #4
   652: 61303f05                 tbl z1.b, { z3.b }, z31.b
   656: 82302005                 tbl z2.b, { z4.b }, z0.b
   660: 53382104                 eor3 z19.d, z19.d, z1.d, z2.d
   664: 5f332504                 and z31.d, z26.d, z5.d
   668: 40972c04                 lsr z0.b, z26.b, #4
   672: 61303f05                 tbl z1.b, { z3.b }, z31.b
   676: 82302005                 tbl z2.b, { z4.b }, z0.b
   680: 54382104                 eor3 z20.d, z20.d, z1.d, z2.d
   684: 7f332504                 and z31.d, z27.d, z5.d
   688: 60972c04                 lsr z0.b, z27.b, #4
   692: 61303f05                 tbl z1.b, { z3.b }, z31.b
   696: 82302005                 tbl z2.b, { z4.b }, z0.b
   700: 55382104                 eor3 z21.d, z21.d, z1.d, z2.d
   704: 0f020091                 add x15, x16, #0
   708: e32100a4                 ld1rqb { z3.b }, p0/z, [x15]
   712: 0f020191                 add x15, x16, #0x40
   716: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   720: 3f332504                 and z31.d, z25.d, z5.d
   724: 20972c04                 lsr z0.b, z25.b, #4
   728: 61303f05                 tbl z1.b, { z3.b }, z31.b
   732: 82302005                 tbl z2.b, { z4.b }, z0.b
   736: 56382104                 eor3 z22.d, z22.d, z1.d, z2.d
   740: 5f332504                 and z31.d, z26.d, z5.d
   744: 40972c04                 lsr z0.b, z26.b, #4
   748: 61303f05                 tbl z1.b, { z3.b }, z31.b
   752: 82302005                 tbl z2.b, { z4.b }, z0.b
   756: 57382104                 eor3 z23.d, z23.d, z1.d, z2.d
   760: 7f332504                 and z31.d, z27.d, z5.d
   764: 60972c04                 lsr z0.b, z27.b, #4
   768: 61303f05                 tbl z1.b, { z3.b }, z31.b
   772: 82302005                 tbl z2.b, { z4.b }, z0.b
   776: 58382104                 eor3 z24.d, z24.d, z1.d, z2.d
   780: 0a1140f9                 ldr x10, [x8, #0x20]
   784: 4a010b8b                 add x10, x10, x11
   788: 59a100a4                 ld1b { z25.b }, p0/z, [x10]
   792: 5aa101a4                 ld1b { z26.b }, p0/z, [x10, #1, mul vl]
   796: 5ba102a4                 ld1b { z27.b }, p0/z, [x10, #2, mul vl]
   800: 30031904                 eor z16.b, p0/m, z16.b, z25.b
   804: 51031904                 eor z17.b, p0/m, z17.b, z26.b
   808: 72031904                 eor z18.b, p0/m, z18.b, z27.b
   812: 0f020691                 add x15, x16, #0x180
   816: e32100a4                 ld1rqb { z3.b }, p0/z, [x15]
   820: 0f020791                 add x15, x16, #0x1c0
   824: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   828: 3f332504                 and z31.d, z25.d, z5.d
   832: 20972c04                 lsr z0.b, z25.b, #4
   836: 61303f05                 tbl z1.b, { z3.b }, z31.b
   840: 82302005                 tbl z2.b, { z4.b }, z0.b
   844: 53382104                 eor3 z19.d, z19.d, z1.d, z2.d
   848: 5f332504                 and z31.d, z26.d, z5.d
   852: 40972c04                 lsr z0.b, z26.b, #4
   856: 61303f05                 tbl z1.b, { z3.b }, z31.b
   860: 82302005                 tbl z2.b, { z4.b }, z0.b
   864: 54382104                 eor3 z20.d, z20.d, z1.d, z2.d
   868: 7f332504                 and z31.d, z27.d, z5.d
   872: 60972c04                 lsr z0.b, z27.b, #4
   876: 61303f05                 tbl z1.b, { z3.b }, z31.b
   880: 82302005                 tbl z2.b, { z4.b }, z0.b
   884: 55382104                 eor3 z21.d, z21.d, z1.d, z2.d
   888: 0f020891                 add x15, x16, #0x200
   892: e32100a4                 ld1rqb { z3.b }, p0/z, [x15]
   896: 0f020991                 add x15, x16, #0x240
   900: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   904: 3f332504                 and z31.d, z25.d, z5.d
   908: 20972c04                 lsr z0.b, z25.b, #4
   912: 61303f05                 tbl z1.b, { z3.b }, z31.b
   916: 82302005                 tbl z2.b, { z4.b }, z0.b
   920: 56382104                 eor3 z22.d, z22.d, z1.d, z2.d
   924: 5f332504                 and z31.d, z26.d, z5.d
   928: 40972c04                 lsr z0.b, z26.b, #4
   932: 61303f05                 tbl z1.b, { z3.b }, z31.b
   936: 82302005                 tbl z2.b, { z4.b }, z0.b
   940: 57382104                 eor3 z23.d, z23.d, z1.d, z2.d
   944: 7f332504                 and z31.d, z27.d, z5.d
   948: 60972c04                 lsr z0.b, z27.b, #4
   952: 61303f05                 tbl z1.b, { z3.b }, z31.b
   956: 82302005                 tbl z2.b, { z4.b }, z0.b
   960: 58382104                 eor3 z24.d, z24.d, z1.d, z2.d
   964: 0a1540f9                 ldr x10, [x8, #0x28]
   968: 4a010b8b                 add x10, x10, x11
   972: 59a100a4                 ld1b { z25.b }, p0/z, [x10]
   976: 5aa101a4                 ld1b { z26.b }, p0/z, [x10, #1, mul vl]
   980: 5ba102a4                 ld1b { z27.b }, p0/z, [x10, #2, mul vl]
   984: 30031904                 eor z16.b, p0/m, z16.b, z25.b
   988: 51031904                 eor z17.b, p0/m, z17.b, z26.b
   992: 72031904                 eor z18.b, p0/m, z18.b, z27.b
   996: 33031904                 eor z19.b, p0/m, z19.b, z25.b
  1000: 54031904                 eor z20.b, p0/m, z20.b, z26.b
  1004: 75031904                 eor z21.b, p0/m, z21.b, z27.b
  1008: 36031904                 eor z22.b, p0/m, z22.b, z25.b
  1012: 57031904                 eor z23.b, p0/m, z23.b, z26.b
  1016: 78031904                 eor z24.b, p0/m, z24.b, z27.b
  1020: df0500f1                 cmp x14, #1
  1024: ed000054                 b.le 0x41c
  1028: 0a0140f9                 ldr x10, [x8]
  1032: 4a010b8b                 add x10, x10, x11
  1036: 4ac10091                 add x10, x10, #0x30
  1040: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
  1044: 5da101a4                 ld1b { z29.b }, p0/z, [x10, #1, mul vl]
  1048: 5ea102a4                 ld1b { z30.b }, p0/z, [x10, #2, mul vl]
  1052: 2a0140f9                 ldr x10, [x9]
  1056: 4a010b8b                 add x10, x10, x11
  1060: 50e100e4                 st1b { z16.b }, p0, [x10]
  1064: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
  1068: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
  1072: 2a0540f9                 ldr x10, [x9, #8]
  1076: 4a010b8b                 add x10, x10, x11
  1080: 53e100e4                 st1b { z19.b }, p0, [x10]
  1084: 54e101e4                 st1b { z20.b }, p0, [x10, #1, mul vl]
  1088: 55e102e4                 st1b { z21.b }, p0, [x10, #2, mul vl]
  1092: 2a0940f9                 ldr x10, [x9, #0x10]
  1096: 4a010b8b                 add x10, x10, x11
  1100: 56e100e4                 st1b { z22.b }, p0, [x10]
  1104: 57e101e4                 st1b { z23.b }, p0, [x10, #1, mul vl]
  1108: 58e102e4                 st1b { z24.b }, p0, [x10, #2, mul vl]
  1112: 6b010d8b                 add x11, x11, x13
  1116: ce0500f1                 subs x14, x14, #1
  1120: 01dfff54                 b.ne 0x40
  1124: 7f0102eb                 cmp x11, x2
  1128: 82120054                 b.hs 0x6b8
  1132: 0f020a91                 add x15, x16, #0x280
  1136: 11020e91                 add x17, x16, #0x380
  1140: 2a0140f9                 ldr x10, [x9]
  1144: ed030baa                 mov x13, x11
  1148: 0c008052                 mov w12, #0
  1152: 0e0140f9                 ldr x14, [x8]
  1156: ce696d38                 ldrb w14, [x14, x13]
  1160: 8c010e4a                 eor w12, w12, w14
  1164: 0e0540f9                 ldr x14, [x8, #8]
  1168: ce696d38                 ldrb w14, [x14, x13]
  1172: 8c010e4a                 eor w12, w12, w14
  1176: 0e0940f9                 ldr x14, [x8, #0x10]
  1180: ce696d38                 ldrb w14, [x14, x13]
  1184: 8c010e4a                 eor w12, w12, w14
  1188: 0e0d40f9                 ldr x14, [x8, #0x18]
  1192: ce696d38                 ldrb w14, [x14, x13]
  1196: 8c010e4a                 eor w12, w12, w14
  1200: 0e1140f9                 ldr x14, [x8, #0x20]
  1204: ce696d38                 ldrb w14, [x14, x13]
  1208: 8c010e4a                 eor w12, w12, w14
  1212: 0e1540f9                 ldr x14, [x8, #0x28]
  1216: ce696d38                 ldrb w14, [x14, x13]
  1220: 8c010e4a                 eor w12, w12, w14
  1224: 4c692d38                 strb w12, [x10, x13]
  1228: ad050091                 add x13, x13, #1
  1232: bf0102eb                 cmp x13, x2
  1236: 43fdff54                 b.lo 0x47c
  1240: 2a0540f9                 ldr x10, [x9, #8]
  1244: ed030baa                 mov x13, x11
  1248: 0c008052                 mov w12, #0
  1252: 0e0140f9                 ldr x14, [x8]
  1256: ce696d38                 ldrb w14, [x14, x13]
  1260: 0e010034                 cbz w14, 0x50c
  1264: ee696e38                 ldrb w14, [x15, x14]
  1268: ce990311                 add w14, w14, #0xe6
  1272: dffd0371                 cmp w14, #0xff
  1276: 43000054                 b.lo 0x504
  1280: cefd0351                 sub w14, w14, #0xff
  1284: 2e6a6e38                 ldrb w14, [x17, x14]
  1288: 8c010e4a                 eor w12, w12, w14
  1292: 0e0540f9                 ldr x14, [x8, #8]
  1296: ce696d38                 ldrb w14, [x14, x13]
  1300: 0e010034                 cbz w14, 0x534
  1304: ee696e38                 ldrb w14, [x15, x14]
  1308: ce350311                 add w14, w14, #0xcd
  1312: dffd0371                 cmp w14, #0xff
  1316: 43000054                 b.lo 0x52c
  1320: cefd0351                 sub w14, w14, #0xff
  1324: 2e6a6e38                 ldrb w14, [x17, x14]
  1328: 8c010e4a                 eor w12, w12, w14
  1332: 0e0940f9                 ldr x14, [x8, #0x10]
  1336: ce696d38                 ldrb w14, [x14, x13]
  1340: 0e010034                 cbz w14, 0x55c
  1344: ee696e38                 ldrb w14, [x15, x14]
  1348: ce2d0111                 add w14, w14, #0x4b
  1352: dffd0371                 cmp w14, #0xff
  1356: 43000054                 b.lo 0x554
  1360: cefd0351                 sub w14, w14, #0xff
  1364: 2e6a6e38                 ldrb w14, [x17, x14]
  1368: 8c010e4a                 eor w12, w12, w14
  1372: 0e0d40f9                 ldr x14, [x8, #0x18]
  1376: ce696d38                 ldrb w14, [x14, x13]
  1380: 0e010034                 cbz w14, 0x584
  1384: ee696e38                 ldrb w14, [x15, x14]
  1388: cec90011                 add w14, w14, #0x32
  1392: dffd0371                 cmp w14, #0xff
  1396: 43000054                 b.lo 0x57c
  1400: cefd0351                 sub w14, w14, #0xff
  1404: 2e6a6e38                 ldrb w14, [x17, x14]
  1408: 8c010e4a                 eor w12, w12, w14
  1412: 0e1140f9                 ldr x14, [x8, #0x20]
  1416: ce696d38                 ldrb w14, [x14, x13]
  1420: 0e010034                 cbz w14, 0x5ac
  1424: ee696e38                 ldrb w14, [x15, x14]
  1428: ce650011                 add w14, w14, #0x19
  1432: dffd0371                 cmp w14, #0xff
  1436: 43000054                 b.lo 0x5a4
  1440: cefd0351                 sub w14, w14, #0xff
  1444: 2e6a6e38                 ldrb w14, [x17, x14]
  1448: 8c010e4a                 eor w12, w12, w14
  1452: 0e1540f9                 ldr x14, [x8, #0x28]
  1456: ce696d38                 ldrb w14, [x14, x13]
  1460: 8c010e4a                 eor w12, w12, w14
  1464: 4c692d38                 strb w12, [x10, x13]
  1468: ad050091                 add x13, x13, #1
  1472: bf0102eb                 cmp x13, x2
  1476: e3f8ff54                 b.lo 0x4e0
  1480: 2a0940f9                 ldr x10, [x9, #0x10]
  1484: ed030baa                 mov x13, x11
  1488: 0c008052                 mov w12, #0
  1492: 0e0140f9                 ldr x14, [x8]
  1496: ce696d38                 ldrb w14, [x14, x13]
  1500: 0e010034                 cbz w14, 0x5fc
  1504: ee696e38                 ldrb w14, [x15, x14]
  1508: ce350311                 add w14, w14, #0xcd
  1512: dffd0371                 cmp w14, #0xff
  1516: 43000054                 b.lo 0x5f4
  1520: cefd0351                 sub w14, w14, #0xff
  1524: 2e6a6e38                 ldrb w14, [x17, x14]
  1528: 8c010e4a                 eor w12, w12, w14
  1532: 0e0540f9                 ldr x14, [x8, #8]
  1536: ce696d38                 ldrb w14, [x14, x13]
  1540: 0e010034                 cbz w14, 0x624
  1544: ee696e38                 ldrb w14, [x15, x14]
  1548: ce2d0111                 add w14, w14, #0x4b
  1552: dffd0371                 cmp w14, #0xff
  1556: 43000054                 b.lo 0x61c
  1560: cefd0351                 sub w14, w14, #0xff
  1564: 2e6a6e38                 ldrb w14, [x17, x14]
  1568: 8c010e4a                 eor w12, w12, w14
  1572: 0e0940f9                 ldr x14, [x8, #0x10]
  1576: ce696d38                 ldrb w14, [x14, x13]
  1580: 0e010034                 cbz w14, 0x64c
  1584: ee696e38                 ldrb w14, [x15, x14]
  1588: ce650011                 add w14, w14, #0x19
  1592: dffd0371                 cmp w14, #0xff
  1596: 43000054                 b.lo 0x644
  1600: cefd0351                 sub w14, w14, #0xff
  1604: 2e6a6e38                 ldrb w14, [x17, x14]
  1608: 8c010e4a                 eor w12, w12, w14
  1612: 0e0d40f9                 ldr x14, [x8, #0x18]
  1616: ce696d38                 ldrb w14, [x14, x13]
  1620: 0e010034                 cbz w14, 0x674
  1624: ee696e38                 ldrb w14, [x15, x14]
  1628: ce990311                 add w14, w14, #0xe6
  1632: dffd0371                 cmp w14, #0xff
  1636: 43000054                 b.lo 0x66c
  1640: cefd0351                 sub w14, w14, #0xff
  1644: 2e6a6e38                 ldrb w14, [x17, x14]
  1648: 8c010e4a                 eor w12, w12, w14
  1652: 0e1140f9                 ldr x14, [x8, #0x20]
  1656: ce696d38                 ldrb w14, [x14, x13]
  1660: 0e010034                 cbz w14, 0x69c
  1664: ee696e38                 ldrb w14, [x15, x14]
  1668: cec90011                 add w14, w14, #0x32
  1672: dffd0371                 cmp w14, #0xff
  1676: 43000054                 b.lo 0x694
  1680: cefd0351                 sub w14, w14, #0xff
  1684: 2e6a6e38                 ldrb w14, [x17, x14]
  1688: 8c010e4a                 eor w12, w12, w14
  1692: 0e1540f9                 ldr x14, [x8, #0x28]
  1696: ce696d38                 ldrb w14, [x14, x13]
  1700: 8c010e4a                 eor w12, w12, w14
  1704: 4c692d38                 strb w12, [x10, x13]
  1708: ad050091                 add x13, x13, #1
  1712: bf0102eb                 cmp x13, x2
  1716: e3f8ff54                 b.lo 0x5d0
  1720: c0035fd6                 ret 
  1724: 00000000                 udf #0
