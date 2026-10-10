; config_t1_1053 target=sve2 twin=baseline topo=rs(6,3)/rec{0,1} op=recover U=8
; code_size=3136 twins=[] uops_per_byte=4.6015625
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0610010                 adr x16, 0xc40
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: 8df17dd3                 lsl x13, x12, #3
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: e7c13825                 mov z7.b, #0xf
    36: ce4d00b4                 cbz x14, 0x9dc
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
    76: 10306004                 mov z16.d, z0.d
    80: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
    84: 11306004                 mov z17.d, z0.d
    88: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
    92: 12306004                 mov z18.d, z0.d
    96: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   100: 13306004                 mov z19.d, z0.d
   104: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   108: 14306004                 mov z20.d, z0.d
   112: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   116: 15306004                 mov z21.d, z0.d
   120: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   124: 16306004                 mov z22.d, z0.d
   128: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   132: 17306004                 mov z23.d, z0.d
   136: 0a0540f9                 ldr x10, [x8, #8]
   140: 4a010b8b                 add x10, x10, x11
   144: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   148: 10001904                 eor z16.b, p0/m, z16.b, z0.b
   152: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   156: 11001904                 eor z17.b, p0/m, z17.b, z0.b
   160: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   164: 12001904                 eor z18.b, p0/m, z18.b, z0.b
   168: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   172: 13001904                 eor z19.b, p0/m, z19.b, z0.b
   176: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   180: 14001904                 eor z20.b, p0/m, z20.b, z0.b
   184: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   188: 15001904                 eor z21.b, p0/m, z21.b, z0.b
   192: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   196: 16001904                 eor z22.b, p0/m, z22.b, z0.b
   200: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   204: 17001904                 eor z23.b, p0/m, z23.b, z0.b
   208: 0a0940f9                 ldr x10, [x8, #0x10]
   212: 4a010b8b                 add x10, x10, x11
   216: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   220: 10001904                 eor z16.b, p0/m, z16.b, z0.b
   224: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   228: 11001904                 eor z17.b, p0/m, z17.b, z0.b
   232: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   236: 12001904                 eor z18.b, p0/m, z18.b, z0.b
   240: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   244: 13001904                 eor z19.b, p0/m, z19.b, z0.b
   248: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   252: 14001904                 eor z20.b, p0/m, z20.b, z0.b
   256: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   260: 15001904                 eor z21.b, p0/m, z21.b, z0.b
   264: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   268: 16001904                 eor z22.b, p0/m, z22.b, z0.b
   272: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   276: 17001904                 eor z23.b, p0/m, z23.b, z0.b
   280: 0a0d40f9                 ldr x10, [x8, #0x18]
   284: 4a010b8b                 add x10, x10, x11
   288: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   292: 10001904                 eor z16.b, p0/m, z16.b, z0.b
   296: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   300: 11001904                 eor z17.b, p0/m, z17.b, z0.b
   304: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   308: 12001904                 eor z18.b, p0/m, z18.b, z0.b
   312: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   316: 13001904                 eor z19.b, p0/m, z19.b, z0.b
   320: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   324: 14001904                 eor z20.b, p0/m, z20.b, z0.b
   328: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   332: 15001904                 eor z21.b, p0/m, z21.b, z0.b
   336: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   340: 16001904                 eor z22.b, p0/m, z22.b, z0.b
   344: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   348: 17001904                 eor z23.b, p0/m, z23.b, z0.b
   352: 0a1140f9                 ldr x10, [x8, #0x20]
   356: 4a010b8b                 add x10, x10, x11
   360: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   364: 10001904                 eor z16.b, p0/m, z16.b, z0.b
   368: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   372: 11001904                 eor z17.b, p0/m, z17.b, z0.b
   376: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   380: 12001904                 eor z18.b, p0/m, z18.b, z0.b
   384: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   388: 13001904                 eor z19.b, p0/m, z19.b, z0.b
   392: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   396: 14001904                 eor z20.b, p0/m, z20.b, z0.b
   400: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   404: 15001904                 eor z21.b, p0/m, z21.b, z0.b
   408: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   412: 16001904                 eor z22.b, p0/m, z22.b, z0.b
   416: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   420: 17001904                 eor z23.b, p0/m, z23.b, z0.b
   424: 0f020091                 add x15, x16, #0
   428: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   432: 0f020191                 add x15, x16, #0x40
   436: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   440: 0a0140f9                 ldr x10, [x8]
   444: 4a010b8b                 add x10, x10, x11
   448: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   452: 01302704                 and z1.d, z0.d, z7.d
   456: 02942c04                 lsr z2.b, z0.b, #4
   460: a3302105                 tbl z3.b, { z5.b }, z1.b
   464: c4302205                 tbl z4.b, { z6.b }, z2.b
   468: 7830a404                 eor z24.d, z3.d, z4.d
   472: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   476: 01302704                 and z1.d, z0.d, z7.d
   480: 02942c04                 lsr z2.b, z0.b, #4
   484: a3302105                 tbl z3.b, { z5.b }, z1.b
   488: c4302205                 tbl z4.b, { z6.b }, z2.b
   492: 7930a404                 eor z25.d, z3.d, z4.d
   496: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   500: 01302704                 and z1.d, z0.d, z7.d
   504: 02942c04                 lsr z2.b, z0.b, #4
   508: a3302105                 tbl z3.b, { z5.b }, z1.b
   512: c4302205                 tbl z4.b, { z6.b }, z2.b
   516: 7a30a404                 eor z26.d, z3.d, z4.d
   520: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   524: 01302704                 and z1.d, z0.d, z7.d
   528: 02942c04                 lsr z2.b, z0.b, #4
   532: a3302105                 tbl z3.b, { z5.b }, z1.b
   536: c4302205                 tbl z4.b, { z6.b }, z2.b
   540: 7b30a404                 eor z27.d, z3.d, z4.d
   544: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   548: 01302704                 and z1.d, z0.d, z7.d
   552: 02942c04                 lsr z2.b, z0.b, #4
   556: a3302105                 tbl z3.b, { z5.b }, z1.b
   560: c4302205                 tbl z4.b, { z6.b }, z2.b
   564: 7c30a404                 eor z28.d, z3.d, z4.d
   568: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   572: 01302704                 and z1.d, z0.d, z7.d
   576: 02942c04                 lsr z2.b, z0.b, #4
   580: a3302105                 tbl z3.b, { z5.b }, z1.b
   584: c4302205                 tbl z4.b, { z6.b }, z2.b
   588: 7d30a404                 eor z29.d, z3.d, z4.d
   592: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   596: 01302704                 and z1.d, z0.d, z7.d
   600: 02942c04                 lsr z2.b, z0.b, #4
   604: a3302105                 tbl z3.b, { z5.b }, z1.b
   608: c4302205                 tbl z4.b, { z6.b }, z2.b
   612: 7e30a404                 eor z30.d, z3.d, z4.d
   616: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   620: 01302704                 and z1.d, z0.d, z7.d
   624: 02942c04                 lsr z2.b, z0.b, #4
   628: a3302105                 tbl z3.b, { z5.b }, z1.b
   632: c4302205                 tbl z4.b, { z6.b }, z2.b
   636: 7f30a404                 eor z31.d, z3.d, z4.d
   640: 0f020291                 add x15, x16, #0x80
   644: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   648: 0f020391                 add x15, x16, #0xc0
   652: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   656: 0a0540f9                 ldr x10, [x8, #8]
   660: 4a010b8b                 add x10, x10, x11
   664: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   668: 01302704                 and z1.d, z0.d, z7.d
   672: 02942c04                 lsr z2.b, z0.b, #4
   676: a3302105                 tbl z3.b, { z5.b }, z1.b
   680: c4302205                 tbl z4.b, { z6.b }, z2.b
   684: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
   688: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   692: 01302704                 and z1.d, z0.d, z7.d
   696: 02942c04                 lsr z2.b, z0.b, #4
   700: a3302105                 tbl z3.b, { z5.b }, z1.b
   704: c4302205                 tbl z4.b, { z6.b }, z2.b
   708: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
   712: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   716: 01302704                 and z1.d, z0.d, z7.d
   720: 02942c04                 lsr z2.b, z0.b, #4
   724: a3302105                 tbl z3.b, { z5.b }, z1.b
   728: c4302205                 tbl z4.b, { z6.b }, z2.b
   732: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
   736: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   740: 01302704                 and z1.d, z0.d, z7.d
   744: 02942c04                 lsr z2.b, z0.b, #4
   748: a3302105                 tbl z3.b, { z5.b }, z1.b
   752: c4302205                 tbl z4.b, { z6.b }, z2.b
   756: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
   760: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   764: 01302704                 and z1.d, z0.d, z7.d
   768: 02942c04                 lsr z2.b, z0.b, #4
   772: a3302105                 tbl z3.b, { z5.b }, z1.b
   776: c4302205                 tbl z4.b, { z6.b }, z2.b
   780: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
   784: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   788: 01302704                 and z1.d, z0.d, z7.d
   792: 02942c04                 lsr z2.b, z0.b, #4
   796: a3302105                 tbl z3.b, { z5.b }, z1.b
   800: c4302205                 tbl z4.b, { z6.b }, z2.b
   804: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
   808: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   812: 01302704                 and z1.d, z0.d, z7.d
   816: 02942c04                 lsr z2.b, z0.b, #4
   820: a3302105                 tbl z3.b, { z5.b }, z1.b
   824: c4302205                 tbl z4.b, { z6.b }, z2.b
   828: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
   832: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   836: 01302704                 and z1.d, z0.d, z7.d
   840: 02942c04                 lsr z2.b, z0.b, #4
   844: a3302105                 tbl z3.b, { z5.b }, z1.b
   848: c4302205                 tbl z4.b, { z6.b }, z2.b
   852: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
   856: 0f020491                 add x15, x16, #0x100
   860: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   864: 0f020591                 add x15, x16, #0x140
   868: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   872: 0a0940f9                 ldr x10, [x8, #0x10]
   876: 4a010b8b                 add x10, x10, x11
   880: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   884: 01302704                 and z1.d, z0.d, z7.d
   888: 02942c04                 lsr z2.b, z0.b, #4
   892: a3302105                 tbl z3.b, { z5.b }, z1.b
   896: c4302205                 tbl z4.b, { z6.b }, z2.b
   900: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
   904: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   908: 01302704                 and z1.d, z0.d, z7.d
   912: 02942c04                 lsr z2.b, z0.b, #4
   916: a3302105                 tbl z3.b, { z5.b }, z1.b
   920: c4302205                 tbl z4.b, { z6.b }, z2.b
   924: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
   928: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   932: 01302704                 and z1.d, z0.d, z7.d
   936: 02942c04                 lsr z2.b, z0.b, #4
   940: a3302105                 tbl z3.b, { z5.b }, z1.b
   944: c4302205                 tbl z4.b, { z6.b }, z2.b
   948: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
   952: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   956: 01302704                 and z1.d, z0.d, z7.d
   960: 02942c04                 lsr z2.b, z0.b, #4
   964: a3302105                 tbl z3.b, { z5.b }, z1.b
   968: c4302205                 tbl z4.b, { z6.b }, z2.b
   972: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
   976: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   980: 01302704                 and z1.d, z0.d, z7.d
   984: 02942c04                 lsr z2.b, z0.b, #4
   988: a3302105                 tbl z3.b, { z5.b }, z1.b
   992: c4302205                 tbl z4.b, { z6.b }, z2.b
   996: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
  1000: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1004: 01302704                 and z1.d, z0.d, z7.d
  1008: 02942c04                 lsr z2.b, z0.b, #4
  1012: a3302105                 tbl z3.b, { z5.b }, z1.b
  1016: c4302205                 tbl z4.b, { z6.b }, z2.b
  1020: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
  1024: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1028: 01302704                 and z1.d, z0.d, z7.d
  1032: 02942c04                 lsr z2.b, z0.b, #4
  1036: a3302105                 tbl z3.b, { z5.b }, z1.b
  1040: c4302205                 tbl z4.b, { z6.b }, z2.b
  1044: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
  1048: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1052: 01302704                 and z1.d, z0.d, z7.d
  1056: 02942c04                 lsr z2.b, z0.b, #4
  1060: a3302105                 tbl z3.b, { z5.b }, z1.b
  1064: c4302205                 tbl z4.b, { z6.b }, z2.b
  1068: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
  1072: 0a0d40f9                 ldr x10, [x8, #0x18]
  1076: 4a010b8b                 add x10, x10, x11
  1080: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1084: 18001904                 eor z24.b, p0/m, z24.b, z0.b
  1088: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1092: 19001904                 eor z25.b, p0/m, z25.b, z0.b
  1096: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1100: 1a001904                 eor z26.b, p0/m, z26.b, z0.b
  1104: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1108: 1b001904                 eor z27.b, p0/m, z27.b, z0.b
  1112: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1116: 1c001904                 eor z28.b, p0/m, z28.b, z0.b
  1120: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1124: 1d001904                 eor z29.b, p0/m, z29.b, z0.b
  1128: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1132: 1e001904                 eor z30.b, p0/m, z30.b, z0.b
  1136: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1140: 1f001904                 eor z31.b, p0/m, z31.b, z0.b
  1144: 0a1540f9                 ldr x10, [x8, #0x28]
  1148: 4a010b8b                 add x10, x10, x11
  1152: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1156: 18001904                 eor z24.b, p0/m, z24.b, z0.b
  1160: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1164: 19001904                 eor z25.b, p0/m, z25.b, z0.b
  1168: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1172: 1a001904                 eor z26.b, p0/m, z26.b, z0.b
  1176: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1180: 1b001904                 eor z27.b, p0/m, z27.b, z0.b
  1184: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1188: 1c001904                 eor z28.b, p0/m, z28.b, z0.b
  1192: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1196: 1d001904                 eor z29.b, p0/m, z29.b, z0.b
  1200: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1204: 1e001904                 eor z30.b, p0/m, z30.b, z0.b
  1208: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1212: 1f001904                 eor z31.b, p0/m, z31.b, z0.b
  1216: 0f020691                 add x15, x16, #0x180
  1220: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1224: 0f020791                 add x15, x16, #0x1c0
  1228: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1232: 01322704                 and z1.d, z16.d, z7.d
  1236: 02962c04                 lsr z2.b, z16.b, #4
  1240: a3302105                 tbl z3.b, { z5.b }, z1.b
  1244: c4302205                 tbl z4.b, { z6.b }, z2.b
  1248: 6030a404                 eor z0.d, z3.d, z4.d
  1252: 00031904                 eor z0.b, p0/m, z0.b, z24.b
  1256: 0f020891                 add x15, x16, #0x200
  1260: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1264: 0f020991                 add x15, x16, #0x240
  1268: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1272: 01322704                 and z1.d, z16.d, z7.d
  1276: 02962c04                 lsr z2.b, z16.b, #4
  1280: a3302105                 tbl z3.b, { z5.b }, z1.b
  1284: c4302205                 tbl z4.b, { z6.b }, z2.b
  1288: 6330a404                 eor z3.d, z3.d, z4.d
  1292: 03031904                 eor z3.b, p0/m, z3.b, z24.b
  1296: 0f020a91                 add x15, x16, #0x280
  1300: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1304: 0f020b91                 add x15, x16, #0x2c0
  1308: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1312: 61302704                 and z1.d, z3.d, z7.d
  1316: 62942c04                 lsr z2.b, z3.b, #4
  1320: a3302105                 tbl z3.b, { z5.b }, z1.b
  1324: c4302205                 tbl z4.b, { z6.b }, z2.b
  1328: 7830a404                 eor z24.d, z3.d, z4.d
  1332: 0f020a91                 add x15, x16, #0x280
  1336: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1340: 0f020b91                 add x15, x16, #0x2c0
  1344: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1348: 01302704                 and z1.d, z0.d, z7.d
  1352: 02942c04                 lsr z2.b, z0.b, #4
  1356: a3302105                 tbl z3.b, { z5.b }, z1.b
  1360: c4302205                 tbl z4.b, { z6.b }, z2.b
  1364: 7030a404                 eor z16.d, z3.d, z4.d
  1368: 0f020691                 add x15, x16, #0x180
  1372: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1376: 0f020791                 add x15, x16, #0x1c0
  1380: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1384: 21322704                 and z1.d, z17.d, z7.d
  1388: 22962c04                 lsr z2.b, z17.b, #4
  1392: a3302105                 tbl z3.b, { z5.b }, z1.b
  1396: c4302205                 tbl z4.b, { z6.b }, z2.b
  1400: 6030a404                 eor z0.d, z3.d, z4.d
  1404: 20031904                 eor z0.b, p0/m, z0.b, z25.b
  1408: 0f020891                 add x15, x16, #0x200
  1412: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1416: 0f020991                 add x15, x16, #0x240
  1420: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1424: 21322704                 and z1.d, z17.d, z7.d
  1428: 22962c04                 lsr z2.b, z17.b, #4
  1432: a3302105                 tbl z3.b, { z5.b }, z1.b
  1436: c4302205                 tbl z4.b, { z6.b }, z2.b
  1440: 6330a404                 eor z3.d, z3.d, z4.d
  1444: 23031904                 eor z3.b, p0/m, z3.b, z25.b
  1448: 0f020a91                 add x15, x16, #0x280
  1452: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1456: 0f020b91                 add x15, x16, #0x2c0
  1460: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1464: 61302704                 and z1.d, z3.d, z7.d
  1468: 62942c04                 lsr z2.b, z3.b, #4
  1472: a3302105                 tbl z3.b, { z5.b }, z1.b
  1476: c4302205                 tbl z4.b, { z6.b }, z2.b
  1480: 7930a404                 eor z25.d, z3.d, z4.d
  1484: 0f020a91                 add x15, x16, #0x280
  1488: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1492: 0f020b91                 add x15, x16, #0x2c0
  1496: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1500: 01302704                 and z1.d, z0.d, z7.d
  1504: 02942c04                 lsr z2.b, z0.b, #4
  1508: a3302105                 tbl z3.b, { z5.b }, z1.b
  1512: c4302205                 tbl z4.b, { z6.b }, z2.b
  1516: 7130a404                 eor z17.d, z3.d, z4.d
  1520: 0f020691                 add x15, x16, #0x180
  1524: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1528: 0f020791                 add x15, x16, #0x1c0
  1532: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1536: 41322704                 and z1.d, z18.d, z7.d
  1540: 42962c04                 lsr z2.b, z18.b, #4
  1544: a3302105                 tbl z3.b, { z5.b }, z1.b
  1548: c4302205                 tbl z4.b, { z6.b }, z2.b
  1552: 6030a404                 eor z0.d, z3.d, z4.d
  1556: 40031904                 eor z0.b, p0/m, z0.b, z26.b
  1560: 0f020891                 add x15, x16, #0x200
  1564: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1568: 0f020991                 add x15, x16, #0x240
  1572: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1576: 41322704                 and z1.d, z18.d, z7.d
  1580: 42962c04                 lsr z2.b, z18.b, #4
  1584: a3302105                 tbl z3.b, { z5.b }, z1.b
  1588: c4302205                 tbl z4.b, { z6.b }, z2.b
  1592: 6330a404                 eor z3.d, z3.d, z4.d
  1596: 43031904                 eor z3.b, p0/m, z3.b, z26.b
  1600: 0f020a91                 add x15, x16, #0x280
  1604: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1608: 0f020b91                 add x15, x16, #0x2c0
  1612: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1616: 61302704                 and z1.d, z3.d, z7.d
  1620: 62942c04                 lsr z2.b, z3.b, #4
  1624: a3302105                 tbl z3.b, { z5.b }, z1.b
  1628: c4302205                 tbl z4.b, { z6.b }, z2.b
  1632: 7a30a404                 eor z26.d, z3.d, z4.d
  1636: 0f020a91                 add x15, x16, #0x280
  1640: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1644: 0f020b91                 add x15, x16, #0x2c0
  1648: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1652: 01302704                 and z1.d, z0.d, z7.d
  1656: 02942c04                 lsr z2.b, z0.b, #4
  1660: a3302105                 tbl z3.b, { z5.b }, z1.b
  1664: c4302205                 tbl z4.b, { z6.b }, z2.b
  1668: 7230a404                 eor z18.d, z3.d, z4.d
  1672: 0f020691                 add x15, x16, #0x180
  1676: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1680: 0f020791                 add x15, x16, #0x1c0
  1684: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1688: 61322704                 and z1.d, z19.d, z7.d
  1692: 62962c04                 lsr z2.b, z19.b, #4
  1696: a3302105                 tbl z3.b, { z5.b }, z1.b
  1700: c4302205                 tbl z4.b, { z6.b }, z2.b
  1704: 6030a404                 eor z0.d, z3.d, z4.d
  1708: 60031904                 eor z0.b, p0/m, z0.b, z27.b
  1712: 0f020891                 add x15, x16, #0x200
  1716: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1720: 0f020991                 add x15, x16, #0x240
  1724: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1728: 61322704                 and z1.d, z19.d, z7.d
  1732: 62962c04                 lsr z2.b, z19.b, #4
  1736: a3302105                 tbl z3.b, { z5.b }, z1.b
  1740: c4302205                 tbl z4.b, { z6.b }, z2.b
  1744: 6330a404                 eor z3.d, z3.d, z4.d
  1748: 63031904                 eor z3.b, p0/m, z3.b, z27.b
  1752: 0f020a91                 add x15, x16, #0x280
  1756: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1760: 0f020b91                 add x15, x16, #0x2c0
  1764: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1768: 61302704                 and z1.d, z3.d, z7.d
  1772: 62942c04                 lsr z2.b, z3.b, #4
  1776: a3302105                 tbl z3.b, { z5.b }, z1.b
  1780: c4302205                 tbl z4.b, { z6.b }, z2.b
  1784: 7b30a404                 eor z27.d, z3.d, z4.d
  1788: 0f020a91                 add x15, x16, #0x280
  1792: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1796: 0f020b91                 add x15, x16, #0x2c0
  1800: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1804: 01302704                 and z1.d, z0.d, z7.d
  1808: 02942c04                 lsr z2.b, z0.b, #4
  1812: a3302105                 tbl z3.b, { z5.b }, z1.b
  1816: c4302205                 tbl z4.b, { z6.b }, z2.b
  1820: 7330a404                 eor z19.d, z3.d, z4.d
  1824: 0f020691                 add x15, x16, #0x180
  1828: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1832: 0f020791                 add x15, x16, #0x1c0
  1836: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1840: 81322704                 and z1.d, z20.d, z7.d
  1844: 82962c04                 lsr z2.b, z20.b, #4
  1848: a3302105                 tbl z3.b, { z5.b }, z1.b
  1852: c4302205                 tbl z4.b, { z6.b }, z2.b
  1856: 6030a404                 eor z0.d, z3.d, z4.d
  1860: 80031904                 eor z0.b, p0/m, z0.b, z28.b
  1864: 0f020891                 add x15, x16, #0x200
  1868: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1872: 0f020991                 add x15, x16, #0x240
  1876: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1880: 81322704                 and z1.d, z20.d, z7.d
  1884: 82962c04                 lsr z2.b, z20.b, #4
  1888: a3302105                 tbl z3.b, { z5.b }, z1.b
  1892: c4302205                 tbl z4.b, { z6.b }, z2.b
  1896: 6330a404                 eor z3.d, z3.d, z4.d
  1900: 83031904                 eor z3.b, p0/m, z3.b, z28.b
  1904: 0f020a91                 add x15, x16, #0x280
  1908: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1912: 0f020b91                 add x15, x16, #0x2c0
  1916: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1920: 61302704                 and z1.d, z3.d, z7.d
  1924: 62942c04                 lsr z2.b, z3.b, #4
  1928: a3302105                 tbl z3.b, { z5.b }, z1.b
  1932: c4302205                 tbl z4.b, { z6.b }, z2.b
  1936: 7c30a404                 eor z28.d, z3.d, z4.d
  1940: 0f020a91                 add x15, x16, #0x280
  1944: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1948: 0f020b91                 add x15, x16, #0x2c0
  1952: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1956: 01302704                 and z1.d, z0.d, z7.d
  1960: 02942c04                 lsr z2.b, z0.b, #4
  1964: a3302105                 tbl z3.b, { z5.b }, z1.b
  1968: c4302205                 tbl z4.b, { z6.b }, z2.b
  1972: 7430a404                 eor z20.d, z3.d, z4.d
  1976: 0f020691                 add x15, x16, #0x180
  1980: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1984: 0f020791                 add x15, x16, #0x1c0
  1988: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1992: a1322704                 and z1.d, z21.d, z7.d
  1996: a2962c04                 lsr z2.b, z21.b, #4
  2000: a3302105                 tbl z3.b, { z5.b }, z1.b
  2004: c4302205                 tbl z4.b, { z6.b }, z2.b
  2008: 6030a404                 eor z0.d, z3.d, z4.d
  2012: a0031904                 eor z0.b, p0/m, z0.b, z29.b
  2016: 0f020891                 add x15, x16, #0x200
  2020: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2024: 0f020991                 add x15, x16, #0x240
  2028: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2032: a1322704                 and z1.d, z21.d, z7.d
  2036: a2962c04                 lsr z2.b, z21.b, #4
  2040: a3302105                 tbl z3.b, { z5.b }, z1.b
  2044: c4302205                 tbl z4.b, { z6.b }, z2.b
  2048: 6330a404                 eor z3.d, z3.d, z4.d
  2052: a3031904                 eor z3.b, p0/m, z3.b, z29.b
  2056: 0f020a91                 add x15, x16, #0x280
  2060: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2064: 0f020b91                 add x15, x16, #0x2c0
  2068: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2072: 61302704                 and z1.d, z3.d, z7.d
  2076: 62942c04                 lsr z2.b, z3.b, #4
  2080: a3302105                 tbl z3.b, { z5.b }, z1.b
  2084: c4302205                 tbl z4.b, { z6.b }, z2.b
  2088: 7d30a404                 eor z29.d, z3.d, z4.d
  2092: 0f020a91                 add x15, x16, #0x280
  2096: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2100: 0f020b91                 add x15, x16, #0x2c0
  2104: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2108: 01302704                 and z1.d, z0.d, z7.d
  2112: 02942c04                 lsr z2.b, z0.b, #4
  2116: a3302105                 tbl z3.b, { z5.b }, z1.b
  2120: c4302205                 tbl z4.b, { z6.b }, z2.b
  2124: 7530a404                 eor z21.d, z3.d, z4.d
  2128: 0f020691                 add x15, x16, #0x180
  2132: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2136: 0f020791                 add x15, x16, #0x1c0
  2140: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2144: c1322704                 and z1.d, z22.d, z7.d
  2148: c2962c04                 lsr z2.b, z22.b, #4
  2152: a3302105                 tbl z3.b, { z5.b }, z1.b
  2156: c4302205                 tbl z4.b, { z6.b }, z2.b
  2160: 6030a404                 eor z0.d, z3.d, z4.d
  2164: c0031904                 eor z0.b, p0/m, z0.b, z30.b
  2168: 0f020891                 add x15, x16, #0x200
  2172: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2176: 0f020991                 add x15, x16, #0x240
  2180: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2184: c1322704                 and z1.d, z22.d, z7.d
  2188: c2962c04                 lsr z2.b, z22.b, #4
  2192: a3302105                 tbl z3.b, { z5.b }, z1.b
  2196: c4302205                 tbl z4.b, { z6.b }, z2.b
  2200: 6330a404                 eor z3.d, z3.d, z4.d
  2204: c3031904                 eor z3.b, p0/m, z3.b, z30.b
  2208: 0f020a91                 add x15, x16, #0x280
  2212: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2216: 0f020b91                 add x15, x16, #0x2c0
  2220: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2224: 61302704                 and z1.d, z3.d, z7.d
  2228: 62942c04                 lsr z2.b, z3.b, #4
  2232: a3302105                 tbl z3.b, { z5.b }, z1.b
  2236: c4302205                 tbl z4.b, { z6.b }, z2.b
  2240: 7e30a404                 eor z30.d, z3.d, z4.d
  2244: 0f020a91                 add x15, x16, #0x280
  2248: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2252: 0f020b91                 add x15, x16, #0x2c0
  2256: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2260: 01302704                 and z1.d, z0.d, z7.d
  2264: 02942c04                 lsr z2.b, z0.b, #4
  2268: a3302105                 tbl z3.b, { z5.b }, z1.b
  2272: c4302205                 tbl z4.b, { z6.b }, z2.b
  2276: 7630a404                 eor z22.d, z3.d, z4.d
  2280: 0f020691                 add x15, x16, #0x180
  2284: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2288: 0f020791                 add x15, x16, #0x1c0
  2292: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2296: e1322704                 and z1.d, z23.d, z7.d
  2300: e2962c04                 lsr z2.b, z23.b, #4
  2304: a3302105                 tbl z3.b, { z5.b }, z1.b
  2308: c4302205                 tbl z4.b, { z6.b }, z2.b
  2312: 6030a404                 eor z0.d, z3.d, z4.d
  2316: e0031904                 eor z0.b, p0/m, z0.b, z31.b
  2320: 0f020891                 add x15, x16, #0x200
  2324: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2328: 0f020991                 add x15, x16, #0x240
  2332: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2336: e1322704                 and z1.d, z23.d, z7.d
  2340: e2962c04                 lsr z2.b, z23.b, #4
  2344: a3302105                 tbl z3.b, { z5.b }, z1.b
  2348: c4302205                 tbl z4.b, { z6.b }, z2.b
  2352: 6330a404                 eor z3.d, z3.d, z4.d
  2356: e3031904                 eor z3.b, p0/m, z3.b, z31.b
  2360: 0f020a91                 add x15, x16, #0x280
  2364: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2368: 0f020b91                 add x15, x16, #0x2c0
  2372: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2376: 61302704                 and z1.d, z3.d, z7.d
  2380: 62942c04                 lsr z2.b, z3.b, #4
  2384: a3302105                 tbl z3.b, { z5.b }, z1.b
  2388: c4302205                 tbl z4.b, { z6.b }, z2.b
  2392: 7f30a404                 eor z31.d, z3.d, z4.d
  2396: 0f020a91                 add x15, x16, #0x280
  2400: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2404: 0f020b91                 add x15, x16, #0x2c0
  2408: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2412: 01302704                 and z1.d, z0.d, z7.d
  2416: 02942c04                 lsr z2.b, z0.b, #4
  2420: a3302105                 tbl z3.b, { z5.b }, z1.b
  2424: c4302205                 tbl z4.b, { z6.b }, z2.b
  2428: 7730a404                 eor z23.d, z3.d, z4.d
  2432: 2a0140f9                 ldr x10, [x9]
  2436: 4a010b8b                 add x10, x10, x11
  2440: 50e100e4                 st1b { z16.b }, p0, [x10]
  2444: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
  2448: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
  2452: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
  2456: 54e104e4                 st1b { z20.b }, p0, [x10, #4, mul vl]
  2460: 55e105e4                 st1b { z21.b }, p0, [x10, #5, mul vl]
  2464: 56e106e4                 st1b { z22.b }, p0, [x10, #6, mul vl]
  2468: 57e107e4                 st1b { z23.b }, p0, [x10, #7, mul vl]
  2472: 2a0540f9                 ldr x10, [x9, #8]
  2476: 4a010b8b                 add x10, x10, x11
  2480: 58e100e4                 st1b { z24.b }, p0, [x10]
  2484: 59e101e4                 st1b { z25.b }, p0, [x10, #1, mul vl]
  2488: 5ae102e4                 st1b { z26.b }, p0, [x10, #2, mul vl]
  2492: 5be103e4                 st1b { z27.b }, p0, [x10, #3, mul vl]
  2496: 5ce104e4                 st1b { z28.b }, p0, [x10, #4, mul vl]
  2500: 5de105e4                 st1b { z29.b }, p0, [x10, #5, mul vl]
  2504: 5ee106e4                 st1b { z30.b }, p0, [x10, #6, mul vl]
  2508: 5fe107e4                 st1b { z31.b }, p0, [x10, #7, mul vl]
  2512: 6b010d8b                 add x11, x11, x13
  2516: ce0500f1                 subs x14, x14, #1
  2520: 81b2ff54                 b.ne 0x28
  2524: 7f0102eb                 cmp x11, x2
  2528: 22110054                 b.hs 0xc04
  2532: 0f020c91                 add x15, x16, #0x300
  2536: 11021091                 add x17, x16, #0x400
  2540: 2a0140f9                 ldr x10, [x9]
  2544: ed030baa                 mov x13, x11
  2548: 0c008052                 mov w12, #0
  2552: 0e0140f9                 ldr x14, [x8]
  2556: ce696d38                 ldrb w14, [x14, x13]
  2560: 0e010034                 cbz w14, 0xa20
  2564: ee696e38                 ldrb w14, [x15, x14]
  2568: ce050311                 add w14, w14, #0xc1
  2572: dffd0371                 cmp w14, #0xff
  2576: 43000054                 b.lo 0xa18
  2580: cefd0351                 sub w14, w14, #0xff
  2584: 2e6a6e38                 ldrb w14, [x17, x14]
  2588: 8c010e4a                 eor w12, w12, w14
  2592: 0e0540f9                 ldr x14, [x8, #8]
  2596: ce696d38                 ldrb w14, [x14, x13]
  2600: 0e010034                 cbz w14, 0xa48
  2604: ee696e38                 ldrb w14, [x15, x14]
  2608: ce0d0011                 add w14, w14, #3
  2612: dffd0371                 cmp w14, #0xff
  2616: 43000054                 b.lo 0xa40
  2620: cefd0351                 sub w14, w14, #0xff
  2624: 2e6a6e38                 ldrb w14, [x17, x14]
  2628: 8c010e4a                 eor w12, w12, w14
  2632: 0e0940f9                 ldr x14, [x8, #0x10]
  2636: ce696d38                 ldrb w14, [x14, x13]
  2640: 0e010034                 cbz w14, 0xa70
  2644: ee696e38                 ldrb w14, [x15, x14]
  2648: ce190311                 add w14, w14, #0xc6
  2652: dffd0371                 cmp w14, #0xff
  2656: 43000054                 b.lo 0xa68
  2660: cefd0351                 sub w14, w14, #0xff
  2664: 2e6a6e38                 ldrb w14, [x17, x14]
  2668: 8c010e4a                 eor w12, w12, w14
  2672: 0e0d40f9                 ldr x14, [x8, #0x18]
  2676: ce696d38                 ldrb w14, [x14, x13]
  2680: 0e010034                 cbz w14, 0xa98
  2684: ee696e38                 ldrb w14, [x15, x14]
  2688: ce050011                 add w14, w14, #1
  2692: dffd0371                 cmp w14, #0xff
  2696: 43000054                 b.lo 0xa90
  2700: cefd0351                 sub w14, w14, #0xff
  2704: 2e6a6e38                 ldrb w14, [x17, x14]
  2708: 8c010e4a                 eor w12, w12, w14
  2712: 0e1140f9                 ldr x14, [x8, #0x20]
  2716: ce696d38                 ldrb w14, [x14, x13]
  2720: 0e010034                 cbz w14, 0xac0
  2724: ee696e38                 ldrb w14, [x15, x14]
  2728: cef90311                 add w14, w14, #0xfe
  2732: dffd0371                 cmp w14, #0xff
  2736: 43000054                 b.lo 0xab8
  2740: cefd0351                 sub w14, w14, #0xff
  2744: 2e6a6e38                 ldrb w14, [x17, x14]
  2748: 8c010e4a                 eor w12, w12, w14
  2752: 0e1540f9                 ldr x14, [x8, #0x28]
  2756: ce696d38                 ldrb w14, [x14, x13]
  2760: 0e010034                 cbz w14, 0xae8
  2764: ee696e38                 ldrb w14, [x15, x14]
  2768: cec50011                 add w14, w14, #0x31
  2772: dffd0371                 cmp w14, #0xff
  2776: 43000054                 b.lo 0xae0
  2780: cefd0351                 sub w14, w14, #0xff
  2784: 2e6a6e38                 ldrb w14, [x17, x14]
  2788: 8c010e4a                 eor w12, w12, w14
  2792: 4c692d38                 strb w12, [x10, x13]
  2796: ad050091                 add x13, x13, #1
  2800: bf0102eb                 cmp x13, x2
  2804: 03f8ff54                 b.lo 0x9f4
  2808: 2a0540f9                 ldr x10, [x9, #8]
  2812: ed030baa                 mov x13, x11
  2816: 0c008052                 mov w12, #0
  2820: 0e0140f9                 ldr x14, [x8]
  2824: ce696d38                 ldrb w14, [x14, x13]
  2828: 0e010034                 cbz w14, 0xb2c
  2832: ee696e38                 ldrb w14, [x15, x14]
  2836: ce710011                 add w14, w14, #0x1c
  2840: dffd0371                 cmp w14, #0xff
  2844: 43000054                 b.lo 0xb24
  2848: cefd0351                 sub w14, w14, #0xff
  2852: 2e6a6e38                 ldrb w14, [x17, x14]
  2856: 8c010e4a                 eor w12, w12, w14
  2860: 0e0540f9                 ldr x14, [x8, #8]
  2864: ce696d38                 ldrb w14, [x14, x13]
  2868: 0e010034                 cbz w14, 0xb54
  2872: ee696e38                 ldrb w14, [x15, x14]
  2876: ce7d0311                 add w14, w14, #0xdf
  2880: dffd0371                 cmp w14, #0xff
  2884: 43000054                 b.lo 0xb4c
  2888: cefd0351                 sub w14, w14, #0xff
  2892: 2e6a6e38                 ldrb w14, [x17, x14]
  2896: 8c010e4a                 eor w12, w12, w14
  2900: 0e0940f9                 ldr x14, [x8, #0x10]
  2904: ce696d38                 ldrb w14, [x14, x13]
  2908: 0e010034                 cbz w14, 0xb7c
  2912: ee696e38                 ldrb w14, [x15, x14]
  2916: ce690011                 add w14, w14, #0x1a
  2920: dffd0371                 cmp w14, #0xff
  2924: 43000054                 b.lo 0xb74
  2928: cefd0351                 sub w14, w14, #0xff
  2932: 2e6a6e38                 ldrb w14, [x17, x14]
  2936: 8c010e4a                 eor w12, w12, w14
  2940: 0e0d40f9                 ldr x14, [x8, #0x18]
  2944: ce696d38                 ldrb w14, [x14, x13]
  2948: 0e010034                 cbz w14, 0xba4
  2952: ee696e38                 ldrb w14, [x15, x14]
  2956: ce650011                 add w14, w14, #0x19
  2960: dffd0371                 cmp w14, #0xff
  2964: 43000054                 b.lo 0xb9c
  2968: cefd0351                 sub w14, w14, #0xff
  2972: 2e6a6e38                 ldrb w14, [x17, x14]
  2976: 8c010e4a                 eor w12, w12, w14
  2980: 0e1140f9                 ldr x14, [x8, #0x20]
  2984: ce696d38                 ldrb w14, [x14, x13]
  2988: 0e010034                 cbz w14, 0xbcc
  2992: ee696e38                 ldrb w14, [x15, x14]
  2996: ce610011                 add w14, w14, #0x18
  3000: dffd0371                 cmp w14, #0xff
  3004: 43000054                 b.lo 0xbc4
  3008: cefd0351                 sub w14, w14, #0xff
  3012: 2e6a6e38                 ldrb w14, [x17, x14]
  3016: 8c010e4a                 eor w12, w12, w14
  3020: 0e1540f9                 ldr x14, [x8, #0x28]
  3024: ce696d38                 ldrb w14, [x14, x13]
  3028: 0e010034                 cbz w14, 0xbf4
  3032: ee696e38                 ldrb w14, [x15, x14]
  3036: cec50011                 add w14, w14, #0x31
  3040: dffd0371                 cmp w14, #0xff
  3044: 43000054                 b.lo 0xbec
  3048: cefd0351                 sub w14, w14, #0xff
  3052: 2e6a6e38                 ldrb w14, [x17, x14]
  3056: 8c010e4a                 eor w12, w12, w14
  3060: 4c692d38                 strb w12, [x10, x13]
  3064: ad050091                 add x13, x13, #1
  3068: bf0102eb                 cmp x13, x2
  3072: 03f8ff54                 b.lo 0xb00
  3076: c0035fd6                 ret 
  3080: 00000000                 udf #0
  3084: 00000000                 udf #0
  3088: 00000000                 udf #0
  3092: 00000000                 udf #0
  3096: 00000000                 udf #0
  3100: 00000000                 udf #0
  3104: 00000000                 udf #0
  3108: 00000000                 udf #0
  3112: 00000000                 udf #0
  3116: 00000000                 udf #0
  3120: 00000000                 udf #0
  3124: 00000000                 udf #0
  3128: 00000000                 udf #0
  3132: 00000000                 udf #0
