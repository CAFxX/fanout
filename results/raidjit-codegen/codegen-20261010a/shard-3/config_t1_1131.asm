; config_t1_1131 target=sve2 twin=baseline topo=raid6(8)/upd{0,1} op=parity_update U=8
; code_size=1408 twins=[] uops_per_byte=1.8203125
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d02b0010                 adr x16, 0x580
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: 8df17dd3                 lsl x13, x12, #3
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: e7c13825                 mov z7.b, #0xf
    36: 4e2000b4                 cbz x14, 0x42c
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a1140f9                 ldr x10, [x8, #0x20]
    68: 4a010b8b                 add x10, x10, x11
    72: 50a100a4                 ld1b { z16.b }, p0/z, [x10]
    76: 51a101a4                 ld1b { z17.b }, p0/z, [x10, #1, mul vl]
    80: 52a102a4                 ld1b { z18.b }, p0/z, [x10, #2, mul vl]
    84: 53a103a4                 ld1b { z19.b }, p0/z, [x10, #3, mul vl]
    88: 54a104a4                 ld1b { z20.b }, p0/z, [x10, #4, mul vl]
    92: 55a105a4                 ld1b { z21.b }, p0/z, [x10, #5, mul vl]
    96: 56a106a4                 ld1b { z22.b }, p0/z, [x10, #6, mul vl]
   100: 57a107a4                 ld1b { z23.b }, p0/z, [x10, #7, mul vl]
   104: 0a1540f9                 ldr x10, [x8, #0x28]
   108: 4a010b8b                 add x10, x10, x11
   112: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   116: 59a101a4                 ld1b { z25.b }, p0/z, [x10, #1, mul vl]
   120: 5aa102a4                 ld1b { z26.b }, p0/z, [x10, #2, mul vl]
   124: 5ba103a4                 ld1b { z27.b }, p0/z, [x10, #3, mul vl]
   128: 5ca104a4                 ld1b { z28.b }, p0/z, [x10, #4, mul vl]
   132: 5da105a4                 ld1b { z29.b }, p0/z, [x10, #5, mul vl]
   136: 5ea106a4                 ld1b { z30.b }, p0/z, [x10, #6, mul vl]
   140: 5fa107a4                 ld1b { z31.b }, p0/z, [x10, #7, mul vl]
   144: 0a0140f9                 ldr x10, [x8]
   148: 4a010b8b                 add x10, x10, x11
   152: 030940f9                 ldr x3, [x8, #0x10]
   156: 63000b8b                 add x3, x3, x11
   160: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   164: 63a000a4                 ld1b { z3.b }, p0/z, [x3]
   168: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   172: 10001904                 eor z16.b, p0/m, z16.b, z0.b
   176: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   180: 63a001a4                 ld1b { z3.b }, p0/z, [x3, #1, mul vl]
   184: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   188: 11001904                 eor z17.b, p0/m, z17.b, z0.b
   192: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   196: 63a002a4                 ld1b { z3.b }, p0/z, [x3, #2, mul vl]
   200: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   204: 12001904                 eor z18.b, p0/m, z18.b, z0.b
   208: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   212: 63a003a4                 ld1b { z3.b }, p0/z, [x3, #3, mul vl]
   216: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   220: 13001904                 eor z19.b, p0/m, z19.b, z0.b
   224: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   228: 63a004a4                 ld1b { z3.b }, p0/z, [x3, #4, mul vl]
   232: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   236: 14001904                 eor z20.b, p0/m, z20.b, z0.b
   240: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   244: 63a005a4                 ld1b { z3.b }, p0/z, [x3, #5, mul vl]
   248: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   252: 15001904                 eor z21.b, p0/m, z21.b, z0.b
   256: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   260: 63a006a4                 ld1b { z3.b }, p0/z, [x3, #6, mul vl]
   264: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   268: 16001904                 eor z22.b, p0/m, z22.b, z0.b
   272: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   276: 63a007a4                 ld1b { z3.b }, p0/z, [x3, #7, mul vl]
   280: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   284: 17001904                 eor z23.b, p0/m, z23.b, z0.b
   288: 0f020091                 add x15, x16, #0
   292: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   296: 0f020191                 add x15, x16, #0x40
   300: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   304: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   308: 63a000a4                 ld1b { z3.b }, p0/z, [x3]
   312: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   316: 01302704                 and z1.d, z0.d, z7.d
   320: 02942c04                 lsr z2.b, z0.b, #4
   324: a3302105                 tbl z3.b, { z5.b }, z1.b
   328: c4302205                 tbl z4.b, { z6.b }, z2.b
   332: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
   336: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   340: 63a001a4                 ld1b { z3.b }, p0/z, [x3, #1, mul vl]
   344: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   348: 01302704                 and z1.d, z0.d, z7.d
   352: 02942c04                 lsr z2.b, z0.b, #4
   356: a3302105                 tbl z3.b, { z5.b }, z1.b
   360: c4302205                 tbl z4.b, { z6.b }, z2.b
   364: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
   368: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   372: 63a002a4                 ld1b { z3.b }, p0/z, [x3, #2, mul vl]
   376: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   380: 01302704                 and z1.d, z0.d, z7.d
   384: 02942c04                 lsr z2.b, z0.b, #4
   388: a3302105                 tbl z3.b, { z5.b }, z1.b
   392: c4302205                 tbl z4.b, { z6.b }, z2.b
   396: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
   400: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   404: 63a003a4                 ld1b { z3.b }, p0/z, [x3, #3, mul vl]
   408: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   412: 01302704                 and z1.d, z0.d, z7.d
   416: 02942c04                 lsr z2.b, z0.b, #4
   420: a3302105                 tbl z3.b, { z5.b }, z1.b
   424: c4302205                 tbl z4.b, { z6.b }, z2.b
   428: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
   432: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   436: 63a004a4                 ld1b { z3.b }, p0/z, [x3, #4, mul vl]
   440: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   444: 01302704                 and z1.d, z0.d, z7.d
   448: 02942c04                 lsr z2.b, z0.b, #4
   452: a3302105                 tbl z3.b, { z5.b }, z1.b
   456: c4302205                 tbl z4.b, { z6.b }, z2.b
   460: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
   464: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   468: 63a005a4                 ld1b { z3.b }, p0/z, [x3, #5, mul vl]
   472: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   476: 01302704                 and z1.d, z0.d, z7.d
   480: 02942c04                 lsr z2.b, z0.b, #4
   484: a3302105                 tbl z3.b, { z5.b }, z1.b
   488: c4302205                 tbl z4.b, { z6.b }, z2.b
   492: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
   496: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   500: 63a006a4                 ld1b { z3.b }, p0/z, [x3, #6, mul vl]
   504: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   508: 01302704                 and z1.d, z0.d, z7.d
   512: 02942c04                 lsr z2.b, z0.b, #4
   516: a3302105                 tbl z3.b, { z5.b }, z1.b
   520: c4302205                 tbl z4.b, { z6.b }, z2.b
   524: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
   528: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   532: 63a007a4                 ld1b { z3.b }, p0/z, [x3, #7, mul vl]
   536: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   540: 01302704                 and z1.d, z0.d, z7.d
   544: 02942c04                 lsr z2.b, z0.b, #4
   548: a3302105                 tbl z3.b, { z5.b }, z1.b
   552: c4302205                 tbl z4.b, { z6.b }, z2.b
   556: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
   560: 0a0540f9                 ldr x10, [x8, #8]
   564: 4a010b8b                 add x10, x10, x11
   568: 030d40f9                 ldr x3, [x8, #0x18]
   572: 63000b8b                 add x3, x3, x11
   576: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   580: 63a000a4                 ld1b { z3.b }, p0/z, [x3]
   584: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   588: 10001904                 eor z16.b, p0/m, z16.b, z0.b
   592: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   596: 63a001a4                 ld1b { z3.b }, p0/z, [x3, #1, mul vl]
   600: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   604: 11001904                 eor z17.b, p0/m, z17.b, z0.b
   608: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   612: 63a002a4                 ld1b { z3.b }, p0/z, [x3, #2, mul vl]
   616: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   620: 12001904                 eor z18.b, p0/m, z18.b, z0.b
   624: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   628: 63a003a4                 ld1b { z3.b }, p0/z, [x3, #3, mul vl]
   632: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   636: 13001904                 eor z19.b, p0/m, z19.b, z0.b
   640: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   644: 63a004a4                 ld1b { z3.b }, p0/z, [x3, #4, mul vl]
   648: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   652: 14001904                 eor z20.b, p0/m, z20.b, z0.b
   656: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   660: 63a005a4                 ld1b { z3.b }, p0/z, [x3, #5, mul vl]
   664: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   668: 15001904                 eor z21.b, p0/m, z21.b, z0.b
   672: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   676: 63a006a4                 ld1b { z3.b }, p0/z, [x3, #6, mul vl]
   680: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   684: 16001904                 eor z22.b, p0/m, z22.b, z0.b
   688: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   692: 63a007a4                 ld1b { z3.b }, p0/z, [x3, #7, mul vl]
   696: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   700: 17001904                 eor z23.b, p0/m, z23.b, z0.b
   704: 0f020291                 add x15, x16, #0x80
   708: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   712: 0f020391                 add x15, x16, #0xc0
   716: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   720: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   724: 63a000a4                 ld1b { z3.b }, p0/z, [x3]
   728: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   732: 01302704                 and z1.d, z0.d, z7.d
   736: 02942c04                 lsr z2.b, z0.b, #4
   740: a3302105                 tbl z3.b, { z5.b }, z1.b
   744: c4302205                 tbl z4.b, { z6.b }, z2.b
   748: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
   752: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   756: 63a001a4                 ld1b { z3.b }, p0/z, [x3, #1, mul vl]
   760: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   764: 01302704                 and z1.d, z0.d, z7.d
   768: 02942c04                 lsr z2.b, z0.b, #4
   772: a3302105                 tbl z3.b, { z5.b }, z1.b
   776: c4302205                 tbl z4.b, { z6.b }, z2.b
   780: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
   784: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   788: 63a002a4                 ld1b { z3.b }, p0/z, [x3, #2, mul vl]
   792: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   796: 01302704                 and z1.d, z0.d, z7.d
   800: 02942c04                 lsr z2.b, z0.b, #4
   804: a3302105                 tbl z3.b, { z5.b }, z1.b
   808: c4302205                 tbl z4.b, { z6.b }, z2.b
   812: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
   816: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   820: 63a003a4                 ld1b { z3.b }, p0/z, [x3, #3, mul vl]
   824: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   828: 01302704                 and z1.d, z0.d, z7.d
   832: 02942c04                 lsr z2.b, z0.b, #4
   836: a3302105                 tbl z3.b, { z5.b }, z1.b
   840: c4302205                 tbl z4.b, { z6.b }, z2.b
   844: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
   848: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   852: 63a004a4                 ld1b { z3.b }, p0/z, [x3, #4, mul vl]
   856: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   860: 01302704                 and z1.d, z0.d, z7.d
   864: 02942c04                 lsr z2.b, z0.b, #4
   868: a3302105                 tbl z3.b, { z5.b }, z1.b
   872: c4302205                 tbl z4.b, { z6.b }, z2.b
   876: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
   880: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   884: 63a005a4                 ld1b { z3.b }, p0/z, [x3, #5, mul vl]
   888: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   892: 01302704                 and z1.d, z0.d, z7.d
   896: 02942c04                 lsr z2.b, z0.b, #4
   900: a3302105                 tbl z3.b, { z5.b }, z1.b
   904: c4302205                 tbl z4.b, { z6.b }, z2.b
   908: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
   912: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   916: 63a006a4                 ld1b { z3.b }, p0/z, [x3, #6, mul vl]
   920: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   924: 01302704                 and z1.d, z0.d, z7.d
   928: 02942c04                 lsr z2.b, z0.b, #4
   932: a3302105                 tbl z3.b, { z5.b }, z1.b
   936: c4302205                 tbl z4.b, { z6.b }, z2.b
   940: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
   944: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   948: 63a007a4                 ld1b { z3.b }, p0/z, [x3, #7, mul vl]
   952: 60001904                 eor z0.b, p0/m, z0.b, z3.b
   956: 01302704                 and z1.d, z0.d, z7.d
   960: 02942c04                 lsr z2.b, z0.b, #4
   964: a3302105                 tbl z3.b, { z5.b }, z1.b
   968: c4302205                 tbl z4.b, { z6.b }, z2.b
   972: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
   976: 2a0140f9                 ldr x10, [x9]
   980: 4a010b8b                 add x10, x10, x11
   984: 50e100e4                 st1b { z16.b }, p0, [x10]
   988: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
   992: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
   996: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
  1000: 54e104e4                 st1b { z20.b }, p0, [x10, #4, mul vl]
  1004: 55e105e4                 st1b { z21.b }, p0, [x10, #5, mul vl]
  1008: 56e106e4                 st1b { z22.b }, p0, [x10, #6, mul vl]
  1012: 57e107e4                 st1b { z23.b }, p0, [x10, #7, mul vl]
  1016: 2a0540f9                 ldr x10, [x9, #8]
  1020: 4a010b8b                 add x10, x10, x11
  1024: 58e100e4                 st1b { z24.b }, p0, [x10]
  1028: 59e101e4                 st1b { z25.b }, p0, [x10, #1, mul vl]
  1032: 5ae102e4                 st1b { z26.b }, p0, [x10, #2, mul vl]
  1036: 5be103e4                 st1b { z27.b }, p0, [x10, #3, mul vl]
  1040: 5ce104e4                 st1b { z28.b }, p0, [x10, #4, mul vl]
  1044: 5de105e4                 st1b { z29.b }, p0, [x10, #5, mul vl]
  1048: 5ee106e4                 st1b { z30.b }, p0, [x10, #6, mul vl]
  1052: 5fe107e4                 st1b { z31.b }, p0, [x10, #7, mul vl]
  1056: 6b010d8b                 add x11, x11, x13
  1060: ce0500f1                 subs x14, x14, #1
  1064: 01e0ff54                 b.ne 0x28
  1068: 7f0102eb                 cmp x11, x2
  1072: e2080054                 b.hs 0x54c
  1076: 0f020491                 add x15, x16, #0x100
  1080: 11020891                 add x17, x16, #0x200
  1084: 2a0140f9                 ldr x10, [x9]
  1088: ed030baa                 mov x13, x11
  1092: 0e1140f9                 ldr x14, [x8, #0x20]
  1096: cc696d38                 ldrb w12, [x14, x13]
  1100: 0e0140f9                 ldr x14, [x8]
  1104: ce696d38                 ldrb w14, [x14, x13]
  1108: 8c010e4a                 eor w12, w12, w14
  1112: 0e0940f9                 ldr x14, [x8, #0x10]
  1116: ce696d38                 ldrb w14, [x14, x13]
  1120: 8c010e4a                 eor w12, w12, w14
  1124: 0e0540f9                 ldr x14, [x8, #8]
  1128: ce696d38                 ldrb w14, [x14, x13]
  1132: 8c010e4a                 eor w12, w12, w14
  1136: 0e0d40f9                 ldr x14, [x8, #0x18]
  1140: ce696d38                 ldrb w14, [x14, x13]
  1144: 8c010e4a                 eor w12, w12, w14
  1148: 4c692d38                 strb w12, [x10, x13]
  1152: ad050091                 add x13, x13, #1
  1156: bf0102eb                 cmp x13, x2
  1160: e3fdff54                 b.lo 0x444
  1164: 2a0540f9                 ldr x10, [x9, #8]
  1168: ed030baa                 mov x13, x11
  1172: 0e1540f9                 ldr x14, [x8, #0x28]
  1176: cc696d38                 ldrb w12, [x14, x13]
  1180: 0e0140f9                 ldr x14, [x8]
  1184: ce696d38                 ldrb w14, [x14, x13]
  1188: 0e010034                 cbz w14, 0x4c4
  1192: ee696e38                 ldrb w14, [x15, x14]
  1196: ce990311                 add w14, w14, #0xe6
  1200: dffd0371                 cmp w14, #0xff
  1204: 43000054                 b.lo 0x4bc
  1208: cefd0351                 sub w14, w14, #0xff
  1212: 2e6a6e38                 ldrb w14, [x17, x14]
  1216: 8c010e4a                 eor w12, w12, w14
  1220: 0e0940f9                 ldr x14, [x8, #0x10]
  1224: ce696d38                 ldrb w14, [x14, x13]
  1228: 0e010034                 cbz w14, 0x4ec
  1232: ee696e38                 ldrb w14, [x15, x14]
  1236: ce990311                 add w14, w14, #0xe6
  1240: dffd0371                 cmp w14, #0xff
  1244: 43000054                 b.lo 0x4e4
  1248: cefd0351                 sub w14, w14, #0xff
  1252: 2e6a6e38                 ldrb w14, [x17, x14]
  1256: 8c010e4a                 eor w12, w12, w14
  1260: 0e0540f9                 ldr x14, [x8, #8]
  1264: ce696d38                 ldrb w14, [x14, x13]
  1268: 0e010034                 cbz w14, 0x514
  1272: ee696e38                 ldrb w14, [x15, x14]
  1276: ce350311                 add w14, w14, #0xcd
  1280: dffd0371                 cmp w14, #0xff
  1284: 43000054                 b.lo 0x50c
  1288: cefd0351                 sub w14, w14, #0xff
  1292: 2e6a6e38                 ldrb w14, [x17, x14]
  1296: 8c010e4a                 eor w12, w12, w14
  1300: 0e0d40f9                 ldr x14, [x8, #0x18]
  1304: ce696d38                 ldrb w14, [x14, x13]
  1308: 0e010034                 cbz w14, 0x53c
  1312: ee696e38                 ldrb w14, [x15, x14]
  1316: ce350311                 add w14, w14, #0xcd
  1320: dffd0371                 cmp w14, #0xff
  1324: 43000054                 b.lo 0x534
  1328: cefd0351                 sub w14, w14, #0xff
  1332: 2e6a6e38                 ldrb w14, [x17, x14]
  1336: 8c010e4a                 eor w12, w12, w14
  1340: 4c692d38                 strb w12, [x10, x13]
  1344: ad050091                 add x13, x13, #1
  1348: bf0102eb                 cmp x13, x2
  1352: 63faff54                 b.lo 0x494
  1356: c0035fd6                 ret 
  1360: 00000000                 udf #0
  1364: 00000000                 udf #0
  1368: 00000000                 udf #0
  1372: 00000000                 udf #0
  1376: 00000000                 udf #0
  1380: 00000000                 udf #0
  1384: 00000000                 udf #0
  1388: 00000000                 udf #0
  1392: 00000000                 udf #0
  1396: 00000000                 udf #0
  1400: 00000000                 udf #0
  1404: 00000000                 udf #0
