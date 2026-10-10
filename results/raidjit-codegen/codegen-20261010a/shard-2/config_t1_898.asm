; config_t1_898 target=sve2 twin=nt topo=raid6(8)/enc op=encode U=8
; code_size=2752 twins=["nt"] uops_per_byte=4.1015625
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0550010                 adr x16, 0xac0
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: 8df17dd3                 lsl x13, x12, #3
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: e7c13825                 mov z7.b, #0xf
    36: 4e4500b4                 cbz x14, 0x8cc
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
   136: 0f020091                 add x15, x16, #0
   140: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   144: 0f020191                 add x15, x16, #0x40
   148: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   152: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   156: 01302704                 and z1.d, z0.d, z7.d
   160: 02942c04                 lsr z2.b, z0.b, #4
   164: a3302105                 tbl z3.b, { z5.b }, z1.b
   168: c4302205                 tbl z4.b, { z6.b }, z2.b
   172: 7830a404                 eor z24.d, z3.d, z4.d
   176: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   180: 01302704                 and z1.d, z0.d, z7.d
   184: 02942c04                 lsr z2.b, z0.b, #4
   188: a3302105                 tbl z3.b, { z5.b }, z1.b
   192: c4302205                 tbl z4.b, { z6.b }, z2.b
   196: 7930a404                 eor z25.d, z3.d, z4.d
   200: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   204: 01302704                 and z1.d, z0.d, z7.d
   208: 02942c04                 lsr z2.b, z0.b, #4
   212: a3302105                 tbl z3.b, { z5.b }, z1.b
   216: c4302205                 tbl z4.b, { z6.b }, z2.b
   220: 7a30a404                 eor z26.d, z3.d, z4.d
   224: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   228: 01302704                 and z1.d, z0.d, z7.d
   232: 02942c04                 lsr z2.b, z0.b, #4
   236: a3302105                 tbl z3.b, { z5.b }, z1.b
   240: c4302205                 tbl z4.b, { z6.b }, z2.b
   244: 7b30a404                 eor z27.d, z3.d, z4.d
   248: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   252: 01302704                 and z1.d, z0.d, z7.d
   256: 02942c04                 lsr z2.b, z0.b, #4
   260: a3302105                 tbl z3.b, { z5.b }, z1.b
   264: c4302205                 tbl z4.b, { z6.b }, z2.b
   268: 7c30a404                 eor z28.d, z3.d, z4.d
   272: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   276: 01302704                 and z1.d, z0.d, z7.d
   280: 02942c04                 lsr z2.b, z0.b, #4
   284: a3302105                 tbl z3.b, { z5.b }, z1.b
   288: c4302205                 tbl z4.b, { z6.b }, z2.b
   292: 7d30a404                 eor z29.d, z3.d, z4.d
   296: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   300: 01302704                 and z1.d, z0.d, z7.d
   304: 02942c04                 lsr z2.b, z0.b, #4
   308: a3302105                 tbl z3.b, { z5.b }, z1.b
   312: c4302205                 tbl z4.b, { z6.b }, z2.b
   316: 7e30a404                 eor z30.d, z3.d, z4.d
   320: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   324: 01302704                 and z1.d, z0.d, z7.d
   328: 02942c04                 lsr z2.b, z0.b, #4
   332: a3302105                 tbl z3.b, { z5.b }, z1.b
   336: c4302205                 tbl z4.b, { z6.b }, z2.b
   340: 7f30a404                 eor z31.d, z3.d, z4.d
   344: 0a0540f9                 ldr x10, [x8, #8]
   348: 4a010b8b                 add x10, x10, x11
   352: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   356: 10001904                 eor z16.b, p0/m, z16.b, z0.b
   360: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   364: 11001904                 eor z17.b, p0/m, z17.b, z0.b
   368: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   372: 12001904                 eor z18.b, p0/m, z18.b, z0.b
   376: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   380: 13001904                 eor z19.b, p0/m, z19.b, z0.b
   384: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   388: 14001904                 eor z20.b, p0/m, z20.b, z0.b
   392: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   396: 15001904                 eor z21.b, p0/m, z21.b, z0.b
   400: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   404: 16001904                 eor z22.b, p0/m, z22.b, z0.b
   408: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   412: 17001904                 eor z23.b, p0/m, z23.b, z0.b
   416: 0f020291                 add x15, x16, #0x80
   420: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   424: 0f020391                 add x15, x16, #0xc0
   428: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   432: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   436: 01302704                 and z1.d, z0.d, z7.d
   440: 02942c04                 lsr z2.b, z0.b, #4
   444: a3302105                 tbl z3.b, { z5.b }, z1.b
   448: c4302205                 tbl z4.b, { z6.b }, z2.b
   452: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
   456: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   460: 01302704                 and z1.d, z0.d, z7.d
   464: 02942c04                 lsr z2.b, z0.b, #4
   468: a3302105                 tbl z3.b, { z5.b }, z1.b
   472: c4302205                 tbl z4.b, { z6.b }, z2.b
   476: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
   480: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   484: 01302704                 and z1.d, z0.d, z7.d
   488: 02942c04                 lsr z2.b, z0.b, #4
   492: a3302105                 tbl z3.b, { z5.b }, z1.b
   496: c4302205                 tbl z4.b, { z6.b }, z2.b
   500: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
   504: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   508: 01302704                 and z1.d, z0.d, z7.d
   512: 02942c04                 lsr z2.b, z0.b, #4
   516: a3302105                 tbl z3.b, { z5.b }, z1.b
   520: c4302205                 tbl z4.b, { z6.b }, z2.b
   524: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
   528: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   532: 01302704                 and z1.d, z0.d, z7.d
   536: 02942c04                 lsr z2.b, z0.b, #4
   540: a3302105                 tbl z3.b, { z5.b }, z1.b
   544: c4302205                 tbl z4.b, { z6.b }, z2.b
   548: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
   552: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   556: 01302704                 and z1.d, z0.d, z7.d
   560: 02942c04                 lsr z2.b, z0.b, #4
   564: a3302105                 tbl z3.b, { z5.b }, z1.b
   568: c4302205                 tbl z4.b, { z6.b }, z2.b
   572: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
   576: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   580: 01302704                 and z1.d, z0.d, z7.d
   584: 02942c04                 lsr z2.b, z0.b, #4
   588: a3302105                 tbl z3.b, { z5.b }, z1.b
   592: c4302205                 tbl z4.b, { z6.b }, z2.b
   596: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
   600: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   604: 01302704                 and z1.d, z0.d, z7.d
   608: 02942c04                 lsr z2.b, z0.b, #4
   612: a3302105                 tbl z3.b, { z5.b }, z1.b
   616: c4302205                 tbl z4.b, { z6.b }, z2.b
   620: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
   624: 0a0940f9                 ldr x10, [x8, #0x10]
   628: 4a010b8b                 add x10, x10, x11
   632: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   636: 10001904                 eor z16.b, p0/m, z16.b, z0.b
   640: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   644: 11001904                 eor z17.b, p0/m, z17.b, z0.b
   648: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   652: 12001904                 eor z18.b, p0/m, z18.b, z0.b
   656: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   660: 13001904                 eor z19.b, p0/m, z19.b, z0.b
   664: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   668: 14001904                 eor z20.b, p0/m, z20.b, z0.b
   672: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   676: 15001904                 eor z21.b, p0/m, z21.b, z0.b
   680: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   684: 16001904                 eor z22.b, p0/m, z22.b, z0.b
   688: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   692: 17001904                 eor z23.b, p0/m, z23.b, z0.b
   696: 0f020491                 add x15, x16, #0x100
   700: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   704: 0f020591                 add x15, x16, #0x140
   708: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   712: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   716: 01302704                 and z1.d, z0.d, z7.d
   720: 02942c04                 lsr z2.b, z0.b, #4
   724: a3302105                 tbl z3.b, { z5.b }, z1.b
   728: c4302205                 tbl z4.b, { z6.b }, z2.b
   732: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
   736: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   740: 01302704                 and z1.d, z0.d, z7.d
   744: 02942c04                 lsr z2.b, z0.b, #4
   748: a3302105                 tbl z3.b, { z5.b }, z1.b
   752: c4302205                 tbl z4.b, { z6.b }, z2.b
   756: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
   760: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   764: 01302704                 and z1.d, z0.d, z7.d
   768: 02942c04                 lsr z2.b, z0.b, #4
   772: a3302105                 tbl z3.b, { z5.b }, z1.b
   776: c4302205                 tbl z4.b, { z6.b }, z2.b
   780: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
   784: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   788: 01302704                 and z1.d, z0.d, z7.d
   792: 02942c04                 lsr z2.b, z0.b, #4
   796: a3302105                 tbl z3.b, { z5.b }, z1.b
   800: c4302205                 tbl z4.b, { z6.b }, z2.b
   804: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
   808: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   812: 01302704                 and z1.d, z0.d, z7.d
   816: 02942c04                 lsr z2.b, z0.b, #4
   820: a3302105                 tbl z3.b, { z5.b }, z1.b
   824: c4302205                 tbl z4.b, { z6.b }, z2.b
   828: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
   832: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   836: 01302704                 and z1.d, z0.d, z7.d
   840: 02942c04                 lsr z2.b, z0.b, #4
   844: a3302105                 tbl z3.b, { z5.b }, z1.b
   848: c4302205                 tbl z4.b, { z6.b }, z2.b
   852: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
   856: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   860: 01302704                 and z1.d, z0.d, z7.d
   864: 02942c04                 lsr z2.b, z0.b, #4
   868: a3302105                 tbl z3.b, { z5.b }, z1.b
   872: c4302205                 tbl z4.b, { z6.b }, z2.b
   876: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
   880: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   884: 01302704                 and z1.d, z0.d, z7.d
   888: 02942c04                 lsr z2.b, z0.b, #4
   892: a3302105                 tbl z3.b, { z5.b }, z1.b
   896: c4302205                 tbl z4.b, { z6.b }, z2.b
   900: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
   904: 0a0d40f9                 ldr x10, [x8, #0x18]
   908: 4a010b8b                 add x10, x10, x11
   912: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   916: 10001904                 eor z16.b, p0/m, z16.b, z0.b
   920: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   924: 11001904                 eor z17.b, p0/m, z17.b, z0.b
   928: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   932: 12001904                 eor z18.b, p0/m, z18.b, z0.b
   936: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   940: 13001904                 eor z19.b, p0/m, z19.b, z0.b
   944: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   948: 14001904                 eor z20.b, p0/m, z20.b, z0.b
   952: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   956: 15001904                 eor z21.b, p0/m, z21.b, z0.b
   960: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   964: 16001904                 eor z22.b, p0/m, z22.b, z0.b
   968: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   972: 17001904                 eor z23.b, p0/m, z23.b, z0.b
   976: 0f020691                 add x15, x16, #0x180
   980: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   984: 0f020791                 add x15, x16, #0x1c0
   988: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   992: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   996: 01302704                 and z1.d, z0.d, z7.d
  1000: 02942c04                 lsr z2.b, z0.b, #4
  1004: a3302105                 tbl z3.b, { z5.b }, z1.b
  1008: c4302205                 tbl z4.b, { z6.b }, z2.b
  1012: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
  1016: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1020: 01302704                 and z1.d, z0.d, z7.d
  1024: 02942c04                 lsr z2.b, z0.b, #4
  1028: a3302105                 tbl z3.b, { z5.b }, z1.b
  1032: c4302205                 tbl z4.b, { z6.b }, z2.b
  1036: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
  1040: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1044: 01302704                 and z1.d, z0.d, z7.d
  1048: 02942c04                 lsr z2.b, z0.b, #4
  1052: a3302105                 tbl z3.b, { z5.b }, z1.b
  1056: c4302205                 tbl z4.b, { z6.b }, z2.b
  1060: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
  1064: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1068: 01302704                 and z1.d, z0.d, z7.d
  1072: 02942c04                 lsr z2.b, z0.b, #4
  1076: a3302105                 tbl z3.b, { z5.b }, z1.b
  1080: c4302205                 tbl z4.b, { z6.b }, z2.b
  1084: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
  1088: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1092: 01302704                 and z1.d, z0.d, z7.d
  1096: 02942c04                 lsr z2.b, z0.b, #4
  1100: a3302105                 tbl z3.b, { z5.b }, z1.b
  1104: c4302205                 tbl z4.b, { z6.b }, z2.b
  1108: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
  1112: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1116: 01302704                 and z1.d, z0.d, z7.d
  1120: 02942c04                 lsr z2.b, z0.b, #4
  1124: a3302105                 tbl z3.b, { z5.b }, z1.b
  1128: c4302205                 tbl z4.b, { z6.b }, z2.b
  1132: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
  1136: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1140: 01302704                 and z1.d, z0.d, z7.d
  1144: 02942c04                 lsr z2.b, z0.b, #4
  1148: a3302105                 tbl z3.b, { z5.b }, z1.b
  1152: c4302205                 tbl z4.b, { z6.b }, z2.b
  1156: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
  1160: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1164: 01302704                 and z1.d, z0.d, z7.d
  1168: 02942c04                 lsr z2.b, z0.b, #4
  1172: a3302105                 tbl z3.b, { z5.b }, z1.b
  1176: c4302205                 tbl z4.b, { z6.b }, z2.b
  1180: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
  1184: 0a1140f9                 ldr x10, [x8, #0x20]
  1188: 4a010b8b                 add x10, x10, x11
  1192: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1196: 10001904                 eor z16.b, p0/m, z16.b, z0.b
  1200: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1204: 11001904                 eor z17.b, p0/m, z17.b, z0.b
  1208: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1212: 12001904                 eor z18.b, p0/m, z18.b, z0.b
  1216: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1220: 13001904                 eor z19.b, p0/m, z19.b, z0.b
  1224: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1228: 14001904                 eor z20.b, p0/m, z20.b, z0.b
  1232: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1236: 15001904                 eor z21.b, p0/m, z21.b, z0.b
  1240: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1244: 16001904                 eor z22.b, p0/m, z22.b, z0.b
  1248: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1252: 17001904                 eor z23.b, p0/m, z23.b, z0.b
  1256: 0f020891                 add x15, x16, #0x200
  1260: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1264: 0f020991                 add x15, x16, #0x240
  1268: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1272: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1276: 01302704                 and z1.d, z0.d, z7.d
  1280: 02942c04                 lsr z2.b, z0.b, #4
  1284: a3302105                 tbl z3.b, { z5.b }, z1.b
  1288: c4302205                 tbl z4.b, { z6.b }, z2.b
  1292: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
  1296: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1300: 01302704                 and z1.d, z0.d, z7.d
  1304: 02942c04                 lsr z2.b, z0.b, #4
  1308: a3302105                 tbl z3.b, { z5.b }, z1.b
  1312: c4302205                 tbl z4.b, { z6.b }, z2.b
  1316: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
  1320: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1324: 01302704                 and z1.d, z0.d, z7.d
  1328: 02942c04                 lsr z2.b, z0.b, #4
  1332: a3302105                 tbl z3.b, { z5.b }, z1.b
  1336: c4302205                 tbl z4.b, { z6.b }, z2.b
  1340: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
  1344: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1348: 01302704                 and z1.d, z0.d, z7.d
  1352: 02942c04                 lsr z2.b, z0.b, #4
  1356: a3302105                 tbl z3.b, { z5.b }, z1.b
  1360: c4302205                 tbl z4.b, { z6.b }, z2.b
  1364: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
  1368: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1372: 01302704                 and z1.d, z0.d, z7.d
  1376: 02942c04                 lsr z2.b, z0.b, #4
  1380: a3302105                 tbl z3.b, { z5.b }, z1.b
  1384: c4302205                 tbl z4.b, { z6.b }, z2.b
  1388: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
  1392: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1396: 01302704                 and z1.d, z0.d, z7.d
  1400: 02942c04                 lsr z2.b, z0.b, #4
  1404: a3302105                 tbl z3.b, { z5.b }, z1.b
  1408: c4302205                 tbl z4.b, { z6.b }, z2.b
  1412: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
  1416: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1420: 01302704                 and z1.d, z0.d, z7.d
  1424: 02942c04                 lsr z2.b, z0.b, #4
  1428: a3302105                 tbl z3.b, { z5.b }, z1.b
  1432: c4302205                 tbl z4.b, { z6.b }, z2.b
  1436: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
  1440: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1444: 01302704                 and z1.d, z0.d, z7.d
  1448: 02942c04                 lsr z2.b, z0.b, #4
  1452: a3302105                 tbl z3.b, { z5.b }, z1.b
  1456: c4302205                 tbl z4.b, { z6.b }, z2.b
  1460: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
  1464: 0a1540f9                 ldr x10, [x8, #0x28]
  1468: 4a010b8b                 add x10, x10, x11
  1472: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1476: 10001904                 eor z16.b, p0/m, z16.b, z0.b
  1480: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1484: 11001904                 eor z17.b, p0/m, z17.b, z0.b
  1488: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1492: 12001904                 eor z18.b, p0/m, z18.b, z0.b
  1496: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1500: 13001904                 eor z19.b, p0/m, z19.b, z0.b
  1504: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1508: 14001904                 eor z20.b, p0/m, z20.b, z0.b
  1512: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1516: 15001904                 eor z21.b, p0/m, z21.b, z0.b
  1520: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1524: 16001904                 eor z22.b, p0/m, z22.b, z0.b
  1528: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1532: 17001904                 eor z23.b, p0/m, z23.b, z0.b
  1536: 0f020a91                 add x15, x16, #0x280
  1540: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1544: 0f020b91                 add x15, x16, #0x2c0
  1548: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1552: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1556: 01302704                 and z1.d, z0.d, z7.d
  1560: 02942c04                 lsr z2.b, z0.b, #4
  1564: a3302105                 tbl z3.b, { z5.b }, z1.b
  1568: c4302205                 tbl z4.b, { z6.b }, z2.b
  1572: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
  1576: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1580: 01302704                 and z1.d, z0.d, z7.d
  1584: 02942c04                 lsr z2.b, z0.b, #4
  1588: a3302105                 tbl z3.b, { z5.b }, z1.b
  1592: c4302205                 tbl z4.b, { z6.b }, z2.b
  1596: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
  1600: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1604: 01302704                 and z1.d, z0.d, z7.d
  1608: 02942c04                 lsr z2.b, z0.b, #4
  1612: a3302105                 tbl z3.b, { z5.b }, z1.b
  1616: c4302205                 tbl z4.b, { z6.b }, z2.b
  1620: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
  1624: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1628: 01302704                 and z1.d, z0.d, z7.d
  1632: 02942c04                 lsr z2.b, z0.b, #4
  1636: a3302105                 tbl z3.b, { z5.b }, z1.b
  1640: c4302205                 tbl z4.b, { z6.b }, z2.b
  1644: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
  1648: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1652: 01302704                 and z1.d, z0.d, z7.d
  1656: 02942c04                 lsr z2.b, z0.b, #4
  1660: a3302105                 tbl z3.b, { z5.b }, z1.b
  1664: c4302205                 tbl z4.b, { z6.b }, z2.b
  1668: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
  1672: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1676: 01302704                 and z1.d, z0.d, z7.d
  1680: 02942c04                 lsr z2.b, z0.b, #4
  1684: a3302105                 tbl z3.b, { z5.b }, z1.b
  1688: c4302205                 tbl z4.b, { z6.b }, z2.b
  1692: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
  1696: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1700: 01302704                 and z1.d, z0.d, z7.d
  1704: 02942c04                 lsr z2.b, z0.b, #4
  1708: a3302105                 tbl z3.b, { z5.b }, z1.b
  1712: c4302205                 tbl z4.b, { z6.b }, z2.b
  1716: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
  1720: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1724: 01302704                 and z1.d, z0.d, z7.d
  1728: 02942c04                 lsr z2.b, z0.b, #4
  1732: a3302105                 tbl z3.b, { z5.b }, z1.b
  1736: c4302205                 tbl z4.b, { z6.b }, z2.b
  1740: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
  1744: 0a1940f9                 ldr x10, [x8, #0x30]
  1748: 4a010b8b                 add x10, x10, x11
  1752: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1756: 10001904                 eor z16.b, p0/m, z16.b, z0.b
  1760: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1764: 11001904                 eor z17.b, p0/m, z17.b, z0.b
  1768: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1772: 12001904                 eor z18.b, p0/m, z18.b, z0.b
  1776: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1780: 13001904                 eor z19.b, p0/m, z19.b, z0.b
  1784: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1788: 14001904                 eor z20.b, p0/m, z20.b, z0.b
  1792: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1796: 15001904                 eor z21.b, p0/m, z21.b, z0.b
  1800: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1804: 16001904                 eor z22.b, p0/m, z22.b, z0.b
  1808: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1812: 17001904                 eor z23.b, p0/m, z23.b, z0.b
  1816: 0f020c91                 add x15, x16, #0x300
  1820: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1824: 0f020d91                 add x15, x16, #0x340
  1828: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1832: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1836: 01302704                 and z1.d, z0.d, z7.d
  1840: 02942c04                 lsr z2.b, z0.b, #4
  1844: a3302105                 tbl z3.b, { z5.b }, z1.b
  1848: c4302205                 tbl z4.b, { z6.b }, z2.b
  1852: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
  1856: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1860: 01302704                 and z1.d, z0.d, z7.d
  1864: 02942c04                 lsr z2.b, z0.b, #4
  1868: a3302105                 tbl z3.b, { z5.b }, z1.b
  1872: c4302205                 tbl z4.b, { z6.b }, z2.b
  1876: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
  1880: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1884: 01302704                 and z1.d, z0.d, z7.d
  1888: 02942c04                 lsr z2.b, z0.b, #4
  1892: a3302105                 tbl z3.b, { z5.b }, z1.b
  1896: c4302205                 tbl z4.b, { z6.b }, z2.b
  1900: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
  1904: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1908: 01302704                 and z1.d, z0.d, z7.d
  1912: 02942c04                 lsr z2.b, z0.b, #4
  1916: a3302105                 tbl z3.b, { z5.b }, z1.b
  1920: c4302205                 tbl z4.b, { z6.b }, z2.b
  1924: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
  1928: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1932: 01302704                 and z1.d, z0.d, z7.d
  1936: 02942c04                 lsr z2.b, z0.b, #4
  1940: a3302105                 tbl z3.b, { z5.b }, z1.b
  1944: c4302205                 tbl z4.b, { z6.b }, z2.b
  1948: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
  1952: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1956: 01302704                 and z1.d, z0.d, z7.d
  1960: 02942c04                 lsr z2.b, z0.b, #4
  1964: a3302105                 tbl z3.b, { z5.b }, z1.b
  1968: c4302205                 tbl z4.b, { z6.b }, z2.b
  1972: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
  1976: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1980: 01302704                 and z1.d, z0.d, z7.d
  1984: 02942c04                 lsr z2.b, z0.b, #4
  1988: a3302105                 tbl z3.b, { z5.b }, z1.b
  1992: c4302205                 tbl z4.b, { z6.b }, z2.b
  1996: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
  2000: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  2004: 01302704                 and z1.d, z0.d, z7.d
  2008: 02942c04                 lsr z2.b, z0.b, #4
  2012: a3302105                 tbl z3.b, { z5.b }, z1.b
  2016: c4302205                 tbl z4.b, { z6.b }, z2.b
  2020: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
  2024: 0a1d40f9                 ldr x10, [x8, #0x38]
  2028: 4a010b8b                 add x10, x10, x11
  2032: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2036: 10001904                 eor z16.b, p0/m, z16.b, z0.b
  2040: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2044: 11001904                 eor z17.b, p0/m, z17.b, z0.b
  2048: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2052: 12001904                 eor z18.b, p0/m, z18.b, z0.b
  2056: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2060: 13001904                 eor z19.b, p0/m, z19.b, z0.b
  2064: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  2068: 14001904                 eor z20.b, p0/m, z20.b, z0.b
  2072: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  2076: 15001904                 eor z21.b, p0/m, z21.b, z0.b
  2080: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  2084: 16001904                 eor z22.b, p0/m, z22.b, z0.b
  2088: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  2092: 17001904                 eor z23.b, p0/m, z23.b, z0.b
  2096: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2100: 18001904                 eor z24.b, p0/m, z24.b, z0.b
  2104: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2108: 19001904                 eor z25.b, p0/m, z25.b, z0.b
  2112: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2116: 1a001904                 eor z26.b, p0/m, z26.b, z0.b
  2120: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2124: 1b001904                 eor z27.b, p0/m, z27.b, z0.b
  2128: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  2132: 1c001904                 eor z28.b, p0/m, z28.b, z0.b
  2136: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  2140: 1d001904                 eor z29.b, p0/m, z29.b, z0.b
  2144: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  2148: 1e001904                 eor z30.b, p0/m, z30.b, z0.b
  2152: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  2156: 1f001904                 eor z31.b, p0/m, z31.b, z0.b
  2160: 2a0140f9                 ldr x10, [x9]
  2164: 4a010b8b                 add x10, x10, x11
  2168: 50e110e4                 stnt1b { z16.b }, p0, [x10]
  2172: 51e111e4                 stnt1b { z17.b }, p0, [x10, #1, mul vl]
  2176: 52e112e4                 stnt1b { z18.b }, p0, [x10, #2, mul vl]
  2180: 53e113e4                 stnt1b { z19.b }, p0, [x10, #3, mul vl]
  2184: 54e114e4                 stnt1b { z20.b }, p0, [x10, #4, mul vl]
  2188: 55e115e4                 stnt1b { z21.b }, p0, [x10, #5, mul vl]
  2192: 56e116e4                 stnt1b { z22.b }, p0, [x10, #6, mul vl]
  2196: 57e117e4                 stnt1b { z23.b }, p0, [x10, #7, mul vl]
  2200: 2a0540f9                 ldr x10, [x9, #8]
  2204: 4a010b8b                 add x10, x10, x11
  2208: 58e110e4                 stnt1b { z24.b }, p0, [x10]
  2212: 59e111e4                 stnt1b { z25.b }, p0, [x10, #1, mul vl]
  2216: 5ae112e4                 stnt1b { z26.b }, p0, [x10, #2, mul vl]
  2220: 5be113e4                 stnt1b { z27.b }, p0, [x10, #3, mul vl]
  2224: 5ce114e4                 stnt1b { z28.b }, p0, [x10, #4, mul vl]
  2228: 5de115e4                 stnt1b { z29.b }, p0, [x10, #5, mul vl]
  2232: 5ee116e4                 stnt1b { z30.b }, p0, [x10, #6, mul vl]
  2236: 5fe117e4                 stnt1b { z31.b }, p0, [x10, #7, mul vl]
  2240: 6b010d8b                 add x11, x11, x13
  2244: ce0500f1                 subs x14, x14, #1
  2248: 01bbff54                 b.ne 0x28
  2252: 7f0102eb                 cmp x11, x2
  2256: 420e0054                 b.hs 0xa98
  2260: 0f020e91                 add x15, x16, #0x380
  2264: 11021291                 add x17, x16, #0x480
  2268: 2a0140f9                 ldr x10, [x9]
  2272: ed030baa                 mov x13, x11
  2276: 0c008052                 mov w12, #0
  2280: 0e0140f9                 ldr x14, [x8]
  2284: ce696d38                 ldrb w14, [x14, x13]
  2288: 8c010e4a                 eor w12, w12, w14
  2292: 0e0540f9                 ldr x14, [x8, #8]
  2296: ce696d38                 ldrb w14, [x14, x13]
  2300: 8c010e4a                 eor w12, w12, w14
  2304: 0e0940f9                 ldr x14, [x8, #0x10]
  2308: ce696d38                 ldrb w14, [x14, x13]
  2312: 8c010e4a                 eor w12, w12, w14
  2316: 0e0d40f9                 ldr x14, [x8, #0x18]
  2320: ce696d38                 ldrb w14, [x14, x13]
  2324: 8c010e4a                 eor w12, w12, w14
  2328: 0e1140f9                 ldr x14, [x8, #0x20]
  2332: ce696d38                 ldrb w14, [x14, x13]
  2336: 8c010e4a                 eor w12, w12, w14
  2340: 0e1540f9                 ldr x14, [x8, #0x28]
  2344: ce696d38                 ldrb w14, [x14, x13]
  2348: 8c010e4a                 eor w12, w12, w14
  2352: 0e1940f9                 ldr x14, [x8, #0x30]
  2356: ce696d38                 ldrb w14, [x14, x13]
  2360: 8c010e4a                 eor w12, w12, w14
  2364: 0e1d40f9                 ldr x14, [x8, #0x38]
  2368: ce696d38                 ldrb w14, [x14, x13]
  2372: 8c010e4a                 eor w12, w12, w14
  2376: 4c692d38                 strb w12, [x10, x13]
  2380: ad050091                 add x13, x13, #1
  2384: bf0102eb                 cmp x13, x2
  2388: 83fcff54                 b.lo 0x8e4
  2392: 2a0540f9                 ldr x10, [x9, #8]
  2396: ed030baa                 mov x13, x11
  2400: 0c008052                 mov w12, #0
  2404: 0e0140f9                 ldr x14, [x8]
  2408: ce696d38                 ldrb w14, [x14, x13]
  2412: 0e010034                 cbz w14, 0x98c
  2416: ee696e38                 ldrb w14, [x15, x14]
  2420: ce990311                 add w14, w14, #0xe6
  2424: dffd0371                 cmp w14, #0xff
  2428: 43000054                 b.lo 0x984
  2432: cefd0351                 sub w14, w14, #0xff
  2436: 2e6a6e38                 ldrb w14, [x17, x14]
  2440: 8c010e4a                 eor w12, w12, w14
  2444: 0e0540f9                 ldr x14, [x8, #8]
  2448: ce696d38                 ldrb w14, [x14, x13]
  2452: 0e010034                 cbz w14, 0x9b4
  2456: ee696e38                 ldrb w14, [x15, x14]
  2460: ce350311                 add w14, w14, #0xcd
  2464: dffd0371                 cmp w14, #0xff
  2468: 43000054                 b.lo 0x9ac
  2472: cefd0351                 sub w14, w14, #0xff
  2476: 2e6a6e38                 ldrb w14, [x17, x14]
  2480: 8c010e4a                 eor w12, w12, w14
  2484: 0e0940f9                 ldr x14, [x8, #0x10]
  2488: ce696d38                 ldrb w14, [x14, x13]
  2492: 0e010034                 cbz w14, 0x9dc
  2496: ee696e38                 ldrb w14, [x15, x14]
  2500: ce9d0311                 add w14, w14, #0xe7
  2504: dffd0371                 cmp w14, #0xff
  2508: 43000054                 b.lo 0x9d4
  2512: cefd0351                 sub w14, w14, #0xff
  2516: 2e6a6e38                 ldrb w14, [x17, x14]
  2520: 8c010e4a                 eor w12, w12, w14
  2524: 0e0d40f9                 ldr x14, [x8, #0x18]
  2528: ce696d38                 ldrb w14, [x14, x13]
  2532: 0e010034                 cbz w14, 0xa04
  2536: ee696e38                 ldrb w14, [x15, x14]
  2540: ce390311                 add w14, w14, #0xce
  2544: dffd0371                 cmp w14, #0xff
  2548: 43000054                 b.lo 0x9fc
  2552: cefd0351                 sub w14, w14, #0xff
  2556: 2e6a6e38                 ldrb w14, [x17, x14]
  2560: 8c010e4a                 eor w12, w12, w14
  2564: 0e1140f9                 ldr x14, [x8, #0x20]
  2568: ce696d38                 ldrb w14, [x14, x13]
  2572: 0e010034                 cbz w14, 0xa2c
  2576: ee696e38                 ldrb w14, [x15, x14]
  2580: ce2d0111                 add w14, w14, #0x4b
  2584: dffd0371                 cmp w14, #0xff
  2588: 43000054                 b.lo 0xa24
  2592: cefd0351                 sub w14, w14, #0xff
  2596: 2e6a6e38                 ldrb w14, [x17, x14]
  2600: 8c010e4a                 eor w12, w12, w14
  2604: 0e1540f9                 ldr x14, [x8, #0x28]
  2608: ce696d38                 ldrb w14, [x14, x13]
  2612: 0e010034                 cbz w14, 0xa54
  2616: ee696e38                 ldrb w14, [x15, x14]
  2620: cec90011                 add w14, w14, #0x32
  2624: dffd0371                 cmp w14, #0xff
  2628: 43000054                 b.lo 0xa4c
  2632: cefd0351                 sub w14, w14, #0xff
  2636: 2e6a6e38                 ldrb w14, [x17, x14]
  2640: 8c010e4a                 eor w12, w12, w14
  2644: 0e1940f9                 ldr x14, [x8, #0x30]
  2648: ce696d38                 ldrb w14, [x14, x13]
  2652: 0e010034                 cbz w14, 0xa7c
  2656: ee696e38                 ldrb w14, [x15, x14]
  2660: ce650011                 add w14, w14, #0x19
  2664: dffd0371                 cmp w14, #0xff
  2668: 43000054                 b.lo 0xa74
  2672: cefd0351                 sub w14, w14, #0xff
  2676: 2e6a6e38                 ldrb w14, [x17, x14]
  2680: 8c010e4a                 eor w12, w12, w14
  2684: 0e1d40f9                 ldr x14, [x8, #0x38]
  2688: ce696d38                 ldrb w14, [x14, x13]
  2692: 8c010e4a                 eor w12, w12, w14
  2696: 4c692d38                 strb w12, [x10, x13]
  2700: ad050091                 add x13, x13, #1
  2704: bf0102eb                 cmp x13, x2
  2708: 63f6ff54                 b.lo 0x960
  2712: c0035fd6                 ret 
  2716: 00000000                 udf #0
  2720: 00000000                 udf #0
  2724: 00000000                 udf #0
  2728: 00000000                 udf #0
  2732: 00000000                 udf #0
  2736: 00000000                 udf #0
  2740: 00000000                 udf #0
  2744: 00000000                 udf #0
  2748: 00000000                 udf #0
