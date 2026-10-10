; config_t3_2111 target=sve2 twin=baseline topo=lrc_azure()/enc op=encode U=4
; code_size=4416 twins=[] uops_per_byte=11.640625
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0890010                 adr x16, 0x1140
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: 8df57ed3                 lsl x13, x12, #2
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: e7c13825                 mov z7.b, #0xf
    36: 4e6200b4                 cbz x14, 0xc6c
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
   104: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   108: 14306004                 mov z20.d, z0.d
   112: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   116: 15306004                 mov z21.d, z0.d
   120: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   124: 16306004                 mov z22.d, z0.d
   128: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   132: 17306004                 mov z23.d, z0.d
   136: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   140: 18306004                 mov z24.d, z0.d
   144: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   148: 19306004                 mov z25.d, z0.d
   152: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   156: 1a306004                 mov z26.d, z0.d
   160: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   164: 1b306004                 mov z27.d, z0.d
   168: 0a0540f9                 ldr x10, [x8, #8]
   172: 4a010b8b                 add x10, x10, x11
   176: 0f020091                 add x15, x16, #0
   180: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   184: 0f020191                 add x15, x16, #0x40
   188: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   192: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   196: 01302704                 and z1.d, z0.d, z7.d
   200: 02942c04                 lsr z2.b, z0.b, #4
   204: a3302105                 tbl z3.b, { z5.b }, z1.b
   208: c4302205                 tbl z4.b, { z6.b }, z2.b
   212: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
   216: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   220: 01302704                 and z1.d, z0.d, z7.d
   224: 02942c04                 lsr z2.b, z0.b, #4
   228: a3302105                 tbl z3.b, { z5.b }, z1.b
   232: c4302205                 tbl z4.b, { z6.b }, z2.b
   236: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
   240: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   244: 01302704                 and z1.d, z0.d, z7.d
   248: 02942c04                 lsr z2.b, z0.b, #4
   252: a3302105                 tbl z3.b, { z5.b }, z1.b
   256: c4302205                 tbl z4.b, { z6.b }, z2.b
   260: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
   264: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   268: 01302704                 and z1.d, z0.d, z7.d
   272: 02942c04                 lsr z2.b, z0.b, #4
   276: a3302105                 tbl z3.b, { z5.b }, z1.b
   280: c4302205                 tbl z4.b, { z6.b }, z2.b
   284: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
   288: 0f020291                 add x15, x16, #0x80
   292: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   296: 0f020391                 add x15, x16, #0xc0
   300: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   304: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   308: 01302704                 and z1.d, z0.d, z7.d
   312: 02942c04                 lsr z2.b, z0.b, #4
   316: a3302105                 tbl z3.b, { z5.b }, z1.b
   320: c4302205                 tbl z4.b, { z6.b }, z2.b
   324: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
   328: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   332: 01302704                 and z1.d, z0.d, z7.d
   336: 02942c04                 lsr z2.b, z0.b, #4
   340: a3302105                 tbl z3.b, { z5.b }, z1.b
   344: c4302205                 tbl z4.b, { z6.b }, z2.b
   348: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
   352: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   356: 01302704                 and z1.d, z0.d, z7.d
   360: 02942c04                 lsr z2.b, z0.b, #4
   364: a3302105                 tbl z3.b, { z5.b }, z1.b
   368: c4302205                 tbl z4.b, { z6.b }, z2.b
   372: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
   376: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   380: 01302704                 and z1.d, z0.d, z7.d
   384: 02942c04                 lsr z2.b, z0.b, #4
   388: a3302105                 tbl z3.b, { z5.b }, z1.b
   392: c4302205                 tbl z4.b, { z6.b }, z2.b
   396: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
   400: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   404: 18001904                 eor z24.b, p0/m, z24.b, z0.b
   408: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   412: 19001904                 eor z25.b, p0/m, z25.b, z0.b
   416: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   420: 1a001904                 eor z26.b, p0/m, z26.b, z0.b
   424: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   428: 1b001904                 eor z27.b, p0/m, z27.b, z0.b
   432: 0a0940f9                 ldr x10, [x8, #0x10]
   436: 4a010b8b                 add x10, x10, x11
   440: 0f020291                 add x15, x16, #0x80
   444: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   448: 0f020391                 add x15, x16, #0xc0
   452: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   456: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   460: 01302704                 and z1.d, z0.d, z7.d
   464: 02942c04                 lsr z2.b, z0.b, #4
   468: a3302105                 tbl z3.b, { z5.b }, z1.b
   472: c4302205                 tbl z4.b, { z6.b }, z2.b
   476: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
   480: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   484: 01302704                 and z1.d, z0.d, z7.d
   488: 02942c04                 lsr z2.b, z0.b, #4
   492: a3302105                 tbl z3.b, { z5.b }, z1.b
   496: c4302205                 tbl z4.b, { z6.b }, z2.b
   500: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
   504: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   508: 01302704                 and z1.d, z0.d, z7.d
   512: 02942c04                 lsr z2.b, z0.b, #4
   516: a3302105                 tbl z3.b, { z5.b }, z1.b
   520: c4302205                 tbl z4.b, { z6.b }, z2.b
   524: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
   528: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   532: 01302704                 and z1.d, z0.d, z7.d
   536: 02942c04                 lsr z2.b, z0.b, #4
   540: a3302105                 tbl z3.b, { z5.b }, z1.b
   544: c4302205                 tbl z4.b, { z6.b }, z2.b
   548: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
   552: 0f020491                 add x15, x16, #0x100
   556: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   560: 0f020591                 add x15, x16, #0x140
   564: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   568: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   572: 01302704                 and z1.d, z0.d, z7.d
   576: 02942c04                 lsr z2.b, z0.b, #4
   580: a3302105                 tbl z3.b, { z5.b }, z1.b
   584: c4302205                 tbl z4.b, { z6.b }, z2.b
   588: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
   592: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   596: 01302704                 and z1.d, z0.d, z7.d
   600: 02942c04                 lsr z2.b, z0.b, #4
   604: a3302105                 tbl z3.b, { z5.b }, z1.b
   608: c4302205                 tbl z4.b, { z6.b }, z2.b
   612: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
   616: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   620: 01302704                 and z1.d, z0.d, z7.d
   624: 02942c04                 lsr z2.b, z0.b, #4
   628: a3302105                 tbl z3.b, { z5.b }, z1.b
   632: c4302205                 tbl z4.b, { z6.b }, z2.b
   636: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
   640: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   644: 01302704                 and z1.d, z0.d, z7.d
   648: 02942c04                 lsr z2.b, z0.b, #4
   652: a3302105                 tbl z3.b, { z5.b }, z1.b
   656: c4302205                 tbl z4.b, { z6.b }, z2.b
   660: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
   664: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   668: 18001904                 eor z24.b, p0/m, z24.b, z0.b
   672: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   676: 19001904                 eor z25.b, p0/m, z25.b, z0.b
   680: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   684: 1a001904                 eor z26.b, p0/m, z26.b, z0.b
   688: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   692: 1b001904                 eor z27.b, p0/m, z27.b, z0.b
   696: 0a0d40f9                 ldr x10, [x8, #0x18]
   700: 4a010b8b                 add x10, x10, x11
   704: 0f020691                 add x15, x16, #0x180
   708: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   712: 0f020791                 add x15, x16, #0x1c0
   716: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   720: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   724: 01302704                 and z1.d, z0.d, z7.d
   728: 02942c04                 lsr z2.b, z0.b, #4
   732: a3302105                 tbl z3.b, { z5.b }, z1.b
   736: c4302205                 tbl z4.b, { z6.b }, z2.b
   740: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
   744: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   748: 01302704                 and z1.d, z0.d, z7.d
   752: 02942c04                 lsr z2.b, z0.b, #4
   756: a3302105                 tbl z3.b, { z5.b }, z1.b
   760: c4302205                 tbl z4.b, { z6.b }, z2.b
   764: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
   768: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   772: 01302704                 and z1.d, z0.d, z7.d
   776: 02942c04                 lsr z2.b, z0.b, #4
   780: a3302105                 tbl z3.b, { z5.b }, z1.b
   784: c4302205                 tbl z4.b, { z6.b }, z2.b
   788: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
   792: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   796: 01302704                 and z1.d, z0.d, z7.d
   800: 02942c04                 lsr z2.b, z0.b, #4
   804: a3302105                 tbl z3.b, { z5.b }, z1.b
   808: c4302205                 tbl z4.b, { z6.b }, z2.b
   812: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
   816: 0f020891                 add x15, x16, #0x200
   820: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   824: 0f020991                 add x15, x16, #0x240
   828: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   832: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   836: 01302704                 and z1.d, z0.d, z7.d
   840: 02942c04                 lsr z2.b, z0.b, #4
   844: a3302105                 tbl z3.b, { z5.b }, z1.b
   848: c4302205                 tbl z4.b, { z6.b }, z2.b
   852: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
   856: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   860: 01302704                 and z1.d, z0.d, z7.d
   864: 02942c04                 lsr z2.b, z0.b, #4
   868: a3302105                 tbl z3.b, { z5.b }, z1.b
   872: c4302205                 tbl z4.b, { z6.b }, z2.b
   876: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
   880: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   884: 01302704                 and z1.d, z0.d, z7.d
   888: 02942c04                 lsr z2.b, z0.b, #4
   892: a3302105                 tbl z3.b, { z5.b }, z1.b
   896: c4302205                 tbl z4.b, { z6.b }, z2.b
   900: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
   904: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   908: 01302704                 and z1.d, z0.d, z7.d
   912: 02942c04                 lsr z2.b, z0.b, #4
   916: a3302105                 tbl z3.b, { z5.b }, z1.b
   920: c4302205                 tbl z4.b, { z6.b }, z2.b
   924: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
   928: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   932: 18001904                 eor z24.b, p0/m, z24.b, z0.b
   936: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   940: 19001904                 eor z25.b, p0/m, z25.b, z0.b
   944: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   948: 1a001904                 eor z26.b, p0/m, z26.b, z0.b
   952: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   956: 1b001904                 eor z27.b, p0/m, z27.b, z0.b
   960: 0a1140f9                 ldr x10, [x8, #0x20]
   964: 4a010b8b                 add x10, x10, x11
   968: 0f020491                 add x15, x16, #0x100
   972: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   976: 0f020591                 add x15, x16, #0x140
   980: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   984: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   988: 01302704                 and z1.d, z0.d, z7.d
   992: 02942c04                 lsr z2.b, z0.b, #4
   996: a3302105                 tbl z3.b, { z5.b }, z1.b
  1000: c4302205                 tbl z4.b, { z6.b }, z2.b
  1004: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  1008: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1012: 01302704                 and z1.d, z0.d, z7.d
  1016: 02942c04                 lsr z2.b, z0.b, #4
  1020: a3302105                 tbl z3.b, { z5.b }, z1.b
  1024: c4302205                 tbl z4.b, { z6.b }, z2.b
  1028: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  1032: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1036: 01302704                 and z1.d, z0.d, z7.d
  1040: 02942c04                 lsr z2.b, z0.b, #4
  1044: a3302105                 tbl z3.b, { z5.b }, z1.b
  1048: c4302205                 tbl z4.b, { z6.b }, z2.b
  1052: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  1056: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1060: 01302704                 and z1.d, z0.d, z7.d
  1064: 02942c04                 lsr z2.b, z0.b, #4
  1068: a3302105                 tbl z3.b, { z5.b }, z1.b
  1072: c4302205                 tbl z4.b, { z6.b }, z2.b
  1076: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  1080: 0f020a91                 add x15, x16, #0x280
  1084: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1088: 0f020b91                 add x15, x16, #0x2c0
  1092: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1096: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1100: 01302704                 and z1.d, z0.d, z7.d
  1104: 02942c04                 lsr z2.b, z0.b, #4
  1108: a3302105                 tbl z3.b, { z5.b }, z1.b
  1112: c4302205                 tbl z4.b, { z6.b }, z2.b
  1116: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  1120: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1124: 01302704                 and z1.d, z0.d, z7.d
  1128: 02942c04                 lsr z2.b, z0.b, #4
  1132: a3302105                 tbl z3.b, { z5.b }, z1.b
  1136: c4302205                 tbl z4.b, { z6.b }, z2.b
  1140: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  1144: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1148: 01302704                 and z1.d, z0.d, z7.d
  1152: 02942c04                 lsr z2.b, z0.b, #4
  1156: a3302105                 tbl z3.b, { z5.b }, z1.b
  1160: c4302205                 tbl z4.b, { z6.b }, z2.b
  1164: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  1168: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1172: 01302704                 and z1.d, z0.d, z7.d
  1176: 02942c04                 lsr z2.b, z0.b, #4
  1180: a3302105                 tbl z3.b, { z5.b }, z1.b
  1184: c4302205                 tbl z4.b, { z6.b }, z2.b
  1188: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  1192: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1196: 18001904                 eor z24.b, p0/m, z24.b, z0.b
  1200: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1204: 19001904                 eor z25.b, p0/m, z25.b, z0.b
  1208: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1212: 1a001904                 eor z26.b, p0/m, z26.b, z0.b
  1216: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1220: 1b001904                 eor z27.b, p0/m, z27.b, z0.b
  1224: 0a1540f9                 ldr x10, [x8, #0x28]
  1228: 4a010b8b                 add x10, x10, x11
  1232: 0f020c91                 add x15, x16, #0x300
  1236: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1240: 0f020d91                 add x15, x16, #0x340
  1244: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1248: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1252: 01302704                 and z1.d, z0.d, z7.d
  1256: 02942c04                 lsr z2.b, z0.b, #4
  1260: a3302105                 tbl z3.b, { z5.b }, z1.b
  1264: c4302205                 tbl z4.b, { z6.b }, z2.b
  1268: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  1272: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1276: 01302704                 and z1.d, z0.d, z7.d
  1280: 02942c04                 lsr z2.b, z0.b, #4
  1284: a3302105                 tbl z3.b, { z5.b }, z1.b
  1288: c4302205                 tbl z4.b, { z6.b }, z2.b
  1292: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  1296: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1300: 01302704                 and z1.d, z0.d, z7.d
  1304: 02942c04                 lsr z2.b, z0.b, #4
  1308: a3302105                 tbl z3.b, { z5.b }, z1.b
  1312: c4302205                 tbl z4.b, { z6.b }, z2.b
  1316: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  1320: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1324: 01302704                 and z1.d, z0.d, z7.d
  1328: 02942c04                 lsr z2.b, z0.b, #4
  1332: a3302105                 tbl z3.b, { z5.b }, z1.b
  1336: c4302205                 tbl z4.b, { z6.b }, z2.b
  1340: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  1344: 0f020e91                 add x15, x16, #0x380
  1348: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1352: 0f020f91                 add x15, x16, #0x3c0
  1356: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1360: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1364: 01302704                 and z1.d, z0.d, z7.d
  1368: 02942c04                 lsr z2.b, z0.b, #4
  1372: a3302105                 tbl z3.b, { z5.b }, z1.b
  1376: c4302205                 tbl z4.b, { z6.b }, z2.b
  1380: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  1384: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1388: 01302704                 and z1.d, z0.d, z7.d
  1392: 02942c04                 lsr z2.b, z0.b, #4
  1396: a3302105                 tbl z3.b, { z5.b }, z1.b
  1400: c4302205                 tbl z4.b, { z6.b }, z2.b
  1404: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  1408: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1412: 01302704                 and z1.d, z0.d, z7.d
  1416: 02942c04                 lsr z2.b, z0.b, #4
  1420: a3302105                 tbl z3.b, { z5.b }, z1.b
  1424: c4302205                 tbl z4.b, { z6.b }, z2.b
  1428: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  1432: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1436: 01302704                 and z1.d, z0.d, z7.d
  1440: 02942c04                 lsr z2.b, z0.b, #4
  1444: a3302105                 tbl z3.b, { z5.b }, z1.b
  1448: c4302205                 tbl z4.b, { z6.b }, z2.b
  1452: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  1456: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1460: 18001904                 eor z24.b, p0/m, z24.b, z0.b
  1464: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1468: 19001904                 eor z25.b, p0/m, z25.b, z0.b
  1472: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1476: 1a001904                 eor z26.b, p0/m, z26.b, z0.b
  1480: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1484: 1b001904                 eor z27.b, p0/m, z27.b, z0.b
  1488: 0a1940f9                 ldr x10, [x8, #0x30]
  1492: 4a010b8b                 add x10, x10, x11
  1496: 0f020891                 add x15, x16, #0x200
  1500: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1504: 0f020991                 add x15, x16, #0x240
  1508: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1512: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1516: 01302704                 and z1.d, z0.d, z7.d
  1520: 02942c04                 lsr z2.b, z0.b, #4
  1524: a3302105                 tbl z3.b, { z5.b }, z1.b
  1528: c4302205                 tbl z4.b, { z6.b }, z2.b
  1532: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  1536: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1540: 01302704                 and z1.d, z0.d, z7.d
  1544: 02942c04                 lsr z2.b, z0.b, #4
  1548: a3302105                 tbl z3.b, { z5.b }, z1.b
  1552: c4302205                 tbl z4.b, { z6.b }, z2.b
  1556: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  1560: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1564: 01302704                 and z1.d, z0.d, z7.d
  1568: 02942c04                 lsr z2.b, z0.b, #4
  1572: a3302105                 tbl z3.b, { z5.b }, z1.b
  1576: c4302205                 tbl z4.b, { z6.b }, z2.b
  1580: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  1584: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1588: 01302704                 and z1.d, z0.d, z7.d
  1592: 02942c04                 lsr z2.b, z0.b, #4
  1596: a3302105                 tbl z3.b, { z5.b }, z1.b
  1600: c4302205                 tbl z4.b, { z6.b }, z2.b
  1604: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  1608: 0f021091                 add x15, x16, #0x400
  1612: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1616: 0f021191                 add x15, x16, #0x440
  1620: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1624: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1628: 01302704                 and z1.d, z0.d, z7.d
  1632: 02942c04                 lsr z2.b, z0.b, #4
  1636: a3302105                 tbl z3.b, { z5.b }, z1.b
  1640: c4302205                 tbl z4.b, { z6.b }, z2.b
  1644: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  1648: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1652: 01302704                 and z1.d, z0.d, z7.d
  1656: 02942c04                 lsr z2.b, z0.b, #4
  1660: a3302105                 tbl z3.b, { z5.b }, z1.b
  1664: c4302205                 tbl z4.b, { z6.b }, z2.b
  1668: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  1672: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1676: 01302704                 and z1.d, z0.d, z7.d
  1680: 02942c04                 lsr z2.b, z0.b, #4
  1684: a3302105                 tbl z3.b, { z5.b }, z1.b
  1688: c4302205                 tbl z4.b, { z6.b }, z2.b
  1692: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  1696: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1700: 01302704                 and z1.d, z0.d, z7.d
  1704: 02942c04                 lsr z2.b, z0.b, #4
  1708: a3302105                 tbl z3.b, { z5.b }, z1.b
  1712: c4302205                 tbl z4.b, { z6.b }, z2.b
  1716: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  1720: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1724: 1c306004                 mov z28.d, z0.d
  1728: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1732: 1d306004                 mov z29.d, z0.d
  1736: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1740: 1e306004                 mov z30.d, z0.d
  1744: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1748: 1f306004                 mov z31.d, z0.d
  1752: 0a1d40f9                 ldr x10, [x8, #0x38]
  1756: 4a010b8b                 add x10, x10, x11
  1760: 0f021291                 add x15, x16, #0x480
  1764: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1768: 0f021391                 add x15, x16, #0x4c0
  1772: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1776: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1780: 01302704                 and z1.d, z0.d, z7.d
  1784: 02942c04                 lsr z2.b, z0.b, #4
  1788: a3302105                 tbl z3.b, { z5.b }, z1.b
  1792: c4302205                 tbl z4.b, { z6.b }, z2.b
  1796: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  1800: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1804: 01302704                 and z1.d, z0.d, z7.d
  1808: 02942c04                 lsr z2.b, z0.b, #4
  1812: a3302105                 tbl z3.b, { z5.b }, z1.b
  1816: c4302205                 tbl z4.b, { z6.b }, z2.b
  1820: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  1824: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1828: 01302704                 and z1.d, z0.d, z7.d
  1832: 02942c04                 lsr z2.b, z0.b, #4
  1836: a3302105                 tbl z3.b, { z5.b }, z1.b
  1840: c4302205                 tbl z4.b, { z6.b }, z2.b
  1844: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  1848: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1852: 01302704                 and z1.d, z0.d, z7.d
  1856: 02942c04                 lsr z2.b, z0.b, #4
  1860: a3302105                 tbl z3.b, { z5.b }, z1.b
  1864: c4302205                 tbl z4.b, { z6.b }, z2.b
  1868: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  1872: 0f021491                 add x15, x16, #0x500
  1876: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1880: 0f021591                 add x15, x16, #0x540
  1884: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1888: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1892: 01302704                 and z1.d, z0.d, z7.d
  1896: 02942c04                 lsr z2.b, z0.b, #4
  1900: a3302105                 tbl z3.b, { z5.b }, z1.b
  1904: c4302205                 tbl z4.b, { z6.b }, z2.b
  1908: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  1912: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1916: 01302704                 and z1.d, z0.d, z7.d
  1920: 02942c04                 lsr z2.b, z0.b, #4
  1924: a3302105                 tbl z3.b, { z5.b }, z1.b
  1928: c4302205                 tbl z4.b, { z6.b }, z2.b
  1932: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  1936: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1940: 01302704                 and z1.d, z0.d, z7.d
  1944: 02942c04                 lsr z2.b, z0.b, #4
  1948: a3302105                 tbl z3.b, { z5.b }, z1.b
  1952: c4302205                 tbl z4.b, { z6.b }, z2.b
  1956: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  1960: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1964: 01302704                 and z1.d, z0.d, z7.d
  1968: 02942c04                 lsr z2.b, z0.b, #4
  1972: a3302105                 tbl z3.b, { z5.b }, z1.b
  1976: c4302205                 tbl z4.b, { z6.b }, z2.b
  1980: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  1984: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1988: 1c001904                 eor z28.b, p0/m, z28.b, z0.b
  1992: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1996: 1d001904                 eor z29.b, p0/m, z29.b, z0.b
  2000: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2004: 1e001904                 eor z30.b, p0/m, z30.b, z0.b
  2008: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2012: 1f001904                 eor z31.b, p0/m, z31.b, z0.b
  2016: 0a2140f9                 ldr x10, [x8, #0x40]
  2020: 4a010b8b                 add x10, x10, x11
  2024: 0f020a91                 add x15, x16, #0x280
  2028: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2032: 0f020b91                 add x15, x16, #0x2c0
  2036: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2040: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2044: 01302704                 and z1.d, z0.d, z7.d
  2048: 02942c04                 lsr z2.b, z0.b, #4
  2052: a3302105                 tbl z3.b, { z5.b }, z1.b
  2056: c4302205                 tbl z4.b, { z6.b }, z2.b
  2060: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  2064: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2068: 01302704                 and z1.d, z0.d, z7.d
  2072: 02942c04                 lsr z2.b, z0.b, #4
  2076: a3302105                 tbl z3.b, { z5.b }, z1.b
  2080: c4302205                 tbl z4.b, { z6.b }, z2.b
  2084: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  2088: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2092: 01302704                 and z1.d, z0.d, z7.d
  2096: 02942c04                 lsr z2.b, z0.b, #4
  2100: a3302105                 tbl z3.b, { z5.b }, z1.b
  2104: c4302205                 tbl z4.b, { z6.b }, z2.b
  2108: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  2112: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2116: 01302704                 and z1.d, z0.d, z7.d
  2120: 02942c04                 lsr z2.b, z0.b, #4
  2124: a3302105                 tbl z3.b, { z5.b }, z1.b
  2128: c4302205                 tbl z4.b, { z6.b }, z2.b
  2132: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  2136: 0f021691                 add x15, x16, #0x580
  2140: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2144: 0f021791                 add x15, x16, #0x5c0
  2148: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2152: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2156: 01302704                 and z1.d, z0.d, z7.d
  2160: 02942c04                 lsr z2.b, z0.b, #4
  2164: a3302105                 tbl z3.b, { z5.b }, z1.b
  2168: c4302205                 tbl z4.b, { z6.b }, z2.b
  2172: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  2176: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2180: 01302704                 and z1.d, z0.d, z7.d
  2184: 02942c04                 lsr z2.b, z0.b, #4
  2188: a3302105                 tbl z3.b, { z5.b }, z1.b
  2192: c4302205                 tbl z4.b, { z6.b }, z2.b
  2196: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  2200: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2204: 01302704                 and z1.d, z0.d, z7.d
  2208: 02942c04                 lsr z2.b, z0.b, #4
  2212: a3302105                 tbl z3.b, { z5.b }, z1.b
  2216: c4302205                 tbl z4.b, { z6.b }, z2.b
  2220: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  2224: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2228: 01302704                 and z1.d, z0.d, z7.d
  2232: 02942c04                 lsr z2.b, z0.b, #4
  2236: a3302105                 tbl z3.b, { z5.b }, z1.b
  2240: c4302205                 tbl z4.b, { z6.b }, z2.b
  2244: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  2248: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2252: 1c001904                 eor z28.b, p0/m, z28.b, z0.b
  2256: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2260: 1d001904                 eor z29.b, p0/m, z29.b, z0.b
  2264: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2268: 1e001904                 eor z30.b, p0/m, z30.b, z0.b
  2272: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2276: 1f001904                 eor z31.b, p0/m, z31.b, z0.b
  2280: 0a2540f9                 ldr x10, [x8, #0x48]
  2284: 4a010b8b                 add x10, x10, x11
  2288: 0f021891                 add x15, x16, #0x600
  2292: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2296: 0f021991                 add x15, x16, #0x640
  2300: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2304: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2308: 01302704                 and z1.d, z0.d, z7.d
  2312: 02942c04                 lsr z2.b, z0.b, #4
  2316: a3302105                 tbl z3.b, { z5.b }, z1.b
  2320: c4302205                 tbl z4.b, { z6.b }, z2.b
  2324: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  2328: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2332: 01302704                 and z1.d, z0.d, z7.d
  2336: 02942c04                 lsr z2.b, z0.b, #4
  2340: a3302105                 tbl z3.b, { z5.b }, z1.b
  2344: c4302205                 tbl z4.b, { z6.b }, z2.b
  2348: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  2352: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2356: 01302704                 and z1.d, z0.d, z7.d
  2360: 02942c04                 lsr z2.b, z0.b, #4
  2364: a3302105                 tbl z3.b, { z5.b }, z1.b
  2368: c4302205                 tbl z4.b, { z6.b }, z2.b
  2372: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  2376: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2380: 01302704                 and z1.d, z0.d, z7.d
  2384: 02942c04                 lsr z2.b, z0.b, #4
  2388: a3302105                 tbl z3.b, { z5.b }, z1.b
  2392: c4302205                 tbl z4.b, { z6.b }, z2.b
  2396: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  2400: 0f021a91                 add x15, x16, #0x680
  2404: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2408: 0f021b91                 add x15, x16, #0x6c0
  2412: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2416: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2420: 01302704                 and z1.d, z0.d, z7.d
  2424: 02942c04                 lsr z2.b, z0.b, #4
  2428: a3302105                 tbl z3.b, { z5.b }, z1.b
  2432: c4302205                 tbl z4.b, { z6.b }, z2.b
  2436: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  2440: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2444: 01302704                 and z1.d, z0.d, z7.d
  2448: 02942c04                 lsr z2.b, z0.b, #4
  2452: a3302105                 tbl z3.b, { z5.b }, z1.b
  2456: c4302205                 tbl z4.b, { z6.b }, z2.b
  2460: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  2464: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2468: 01302704                 and z1.d, z0.d, z7.d
  2472: 02942c04                 lsr z2.b, z0.b, #4
  2476: a3302105                 tbl z3.b, { z5.b }, z1.b
  2480: c4302205                 tbl z4.b, { z6.b }, z2.b
  2484: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  2488: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2492: 01302704                 and z1.d, z0.d, z7.d
  2496: 02942c04                 lsr z2.b, z0.b, #4
  2500: a3302105                 tbl z3.b, { z5.b }, z1.b
  2504: c4302205                 tbl z4.b, { z6.b }, z2.b
  2508: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  2512: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2516: 1c001904                 eor z28.b, p0/m, z28.b, z0.b
  2520: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2524: 1d001904                 eor z29.b, p0/m, z29.b, z0.b
  2528: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2532: 1e001904                 eor z30.b, p0/m, z30.b, z0.b
  2536: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2540: 1f001904                 eor z31.b, p0/m, z31.b, z0.b
  2544: 0a2940f9                 ldr x10, [x8, #0x50]
  2548: 4a010b8b                 add x10, x10, x11
  2552: 0f020e91                 add x15, x16, #0x380
  2556: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2560: 0f020f91                 add x15, x16, #0x3c0
  2564: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2568: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2572: 01302704                 and z1.d, z0.d, z7.d
  2576: 02942c04                 lsr z2.b, z0.b, #4
  2580: a3302105                 tbl z3.b, { z5.b }, z1.b
  2584: c4302205                 tbl z4.b, { z6.b }, z2.b
  2588: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  2592: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2596: 01302704                 and z1.d, z0.d, z7.d
  2600: 02942c04                 lsr z2.b, z0.b, #4
  2604: a3302105                 tbl z3.b, { z5.b }, z1.b
  2608: c4302205                 tbl z4.b, { z6.b }, z2.b
  2612: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  2616: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2620: 01302704                 and z1.d, z0.d, z7.d
  2624: 02942c04                 lsr z2.b, z0.b, #4
  2628: a3302105                 tbl z3.b, { z5.b }, z1.b
  2632: c4302205                 tbl z4.b, { z6.b }, z2.b
  2636: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  2640: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2644: 01302704                 and z1.d, z0.d, z7.d
  2648: 02942c04                 lsr z2.b, z0.b, #4
  2652: a3302105                 tbl z3.b, { z5.b }, z1.b
  2656: c4302205                 tbl z4.b, { z6.b }, z2.b
  2660: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  2664: 0f021c91                 add x15, x16, #0x700
  2668: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2672: 0f021d91                 add x15, x16, #0x740
  2676: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2680: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2684: 01302704                 and z1.d, z0.d, z7.d
  2688: 02942c04                 lsr z2.b, z0.b, #4
  2692: a3302105                 tbl z3.b, { z5.b }, z1.b
  2696: c4302205                 tbl z4.b, { z6.b }, z2.b
  2700: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  2704: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2708: 01302704                 and z1.d, z0.d, z7.d
  2712: 02942c04                 lsr z2.b, z0.b, #4
  2716: a3302105                 tbl z3.b, { z5.b }, z1.b
  2720: c4302205                 tbl z4.b, { z6.b }, z2.b
  2724: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  2728: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2732: 01302704                 and z1.d, z0.d, z7.d
  2736: 02942c04                 lsr z2.b, z0.b, #4
  2740: a3302105                 tbl z3.b, { z5.b }, z1.b
  2744: c4302205                 tbl z4.b, { z6.b }, z2.b
  2748: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  2752: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2756: 01302704                 and z1.d, z0.d, z7.d
  2760: 02942c04                 lsr z2.b, z0.b, #4
  2764: a3302105                 tbl z3.b, { z5.b }, z1.b
  2768: c4302205                 tbl z4.b, { z6.b }, z2.b
  2772: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  2776: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2780: 1c001904                 eor z28.b, p0/m, z28.b, z0.b
  2784: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2788: 1d001904                 eor z29.b, p0/m, z29.b, z0.b
  2792: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2796: 1e001904                 eor z30.b, p0/m, z30.b, z0.b
  2800: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2804: 1f001904                 eor z31.b, p0/m, z31.b, z0.b
  2808: 0a2d40f9                 ldr x10, [x8, #0x58]
  2812: 4a010b8b                 add x10, x10, x11
  2816: 0f021e91                 add x15, x16, #0x780
  2820: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2824: 0f021f91                 add x15, x16, #0x7c0
  2828: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2832: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2836: 01302704                 and z1.d, z0.d, z7.d
  2840: 02942c04                 lsr z2.b, z0.b, #4
  2844: a3302105                 tbl z3.b, { z5.b }, z1.b
  2848: c4302205                 tbl z4.b, { z6.b }, z2.b
  2852: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  2856: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2860: 01302704                 and z1.d, z0.d, z7.d
  2864: 02942c04                 lsr z2.b, z0.b, #4
  2868: a3302105                 tbl z3.b, { z5.b }, z1.b
  2872: c4302205                 tbl z4.b, { z6.b }, z2.b
  2876: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  2880: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2884: 01302704                 and z1.d, z0.d, z7.d
  2888: 02942c04                 lsr z2.b, z0.b, #4
  2892: a3302105                 tbl z3.b, { z5.b }, z1.b
  2896: c4302205                 tbl z4.b, { z6.b }, z2.b
  2900: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  2904: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2908: 01302704                 and z1.d, z0.d, z7.d
  2912: 02942c04                 lsr z2.b, z0.b, #4
  2916: a3302105                 tbl z3.b, { z5.b }, z1.b
  2920: c4302205                 tbl z4.b, { z6.b }, z2.b
  2924: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  2928: 0f022091                 add x15, x16, #0x800
  2932: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2936: 0f022191                 add x15, x16, #0x840
  2940: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2944: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2948: 01302704                 and z1.d, z0.d, z7.d
  2952: 02942c04                 lsr z2.b, z0.b, #4
  2956: a3302105                 tbl z3.b, { z5.b }, z1.b
  2960: c4302205                 tbl z4.b, { z6.b }, z2.b
  2964: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  2968: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2972: 01302704                 and z1.d, z0.d, z7.d
  2976: 02942c04                 lsr z2.b, z0.b, #4
  2980: a3302105                 tbl z3.b, { z5.b }, z1.b
  2984: c4302205                 tbl z4.b, { z6.b }, z2.b
  2988: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  2992: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2996: 01302704                 and z1.d, z0.d, z7.d
  3000: 02942c04                 lsr z2.b, z0.b, #4
  3004: a3302105                 tbl z3.b, { z5.b }, z1.b
  3008: c4302205                 tbl z4.b, { z6.b }, z2.b
  3012: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  3016: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  3020: 01302704                 and z1.d, z0.d, z7.d
  3024: 02942c04                 lsr z2.b, z0.b, #4
  3028: a3302105                 tbl z3.b, { z5.b }, z1.b
  3032: c4302205                 tbl z4.b, { z6.b }, z2.b
  3036: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  3040: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  3044: 1c001904                 eor z28.b, p0/m, z28.b, z0.b
  3048: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  3052: 1d001904                 eor z29.b, p0/m, z29.b, z0.b
  3056: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  3060: 1e001904                 eor z30.b, p0/m, z30.b, z0.b
  3064: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  3068: 1f001904                 eor z31.b, p0/m, z31.b, z0.b
  3072: 2a0140f9                 ldr x10, [x9]
  3076: 4a010b8b                 add x10, x10, x11
  3080: 50e100e4                 st1b { z16.b }, p0, [x10]
  3084: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
  3088: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
  3092: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
  3096: 2a0540f9                 ldr x10, [x9, #8]
  3100: 4a010b8b                 add x10, x10, x11
  3104: 54e100e4                 st1b { z20.b }, p0, [x10]
  3108: 55e101e4                 st1b { z21.b }, p0, [x10, #1, mul vl]
  3112: 56e102e4                 st1b { z22.b }, p0, [x10, #2, mul vl]
  3116: 57e103e4                 st1b { z23.b }, p0, [x10, #3, mul vl]
  3120: 2a0940f9                 ldr x10, [x9, #0x10]
  3124: 4a010b8b                 add x10, x10, x11
  3128: 58e100e4                 st1b { z24.b }, p0, [x10]
  3132: 59e101e4                 st1b { z25.b }, p0, [x10, #1, mul vl]
  3136: 5ae102e4                 st1b { z26.b }, p0, [x10, #2, mul vl]
  3140: 5be103e4                 st1b { z27.b }, p0, [x10, #3, mul vl]
  3144: 2a0d40f9                 ldr x10, [x9, #0x18]
  3148: 4a010b8b                 add x10, x10, x11
  3152: 5ce100e4                 st1b { z28.b }, p0, [x10]
  3156: 5de101e4                 st1b { z29.b }, p0, [x10, #1, mul vl]
  3160: 5ee102e4                 st1b { z30.b }, p0, [x10, #2, mul vl]
  3164: 5fe103e4                 st1b { z31.b }, p0, [x10, #3, mul vl]
  3168: 6b010d8b                 add x11, x11, x13
  3172: ce0500f1                 subs x14, x14, #1
  3176: 019eff54                 b.ne 0x28
  3180: 7f0102eb                 cmp x11, x2
  3184: a2240054                 b.hs 0x1104
  3188: 0f022291                 add x15, x16, #0x880
  3192: 11022691                 add x17, x16, #0x980
  3196: 2a0140f9                 ldr x10, [x9]
  3200: ed030baa                 mov x13, x11
  3204: 0c008052                 mov w12, #0
  3208: 0e0140f9                 ldr x14, [x8]
  3212: ce696d38                 ldrb w14, [x14, x13]
  3216: 8c010e4a                 eor w12, w12, w14
  3220: 0e0540f9                 ldr x14, [x8, #8]
  3224: ce696d38                 ldrb w14, [x14, x13]
  3228: 0e010034                 cbz w14, 0xcbc
  3232: ee696e38                 ldrb w14, [x15, x14]
  3236: ce650011                 add w14, w14, #0x19
  3240: dffd0371                 cmp w14, #0xff
  3244: 43000054                 b.lo 0xcb4
  3248: cefd0351                 sub w14, w14, #0xff
  3252: 2e6a6e38                 ldrb w14, [x17, x14]
  3256: 8c010e4a                 eor w12, w12, w14
  3260: 0e0940f9                 ldr x14, [x8, #0x10]
  3264: ce696d38                 ldrb w14, [x14, x13]
  3268: 0e010034                 cbz w14, 0xce4
  3272: ee696e38                 ldrb w14, [x15, x14]
  3276: cec90011                 add w14, w14, #0x32
  3280: dffd0371                 cmp w14, #0xff
  3284: 43000054                 b.lo 0xcdc
  3288: cefd0351                 sub w14, w14, #0xff
  3292: 2e6a6e38                 ldrb w14, [x17, x14]
  3296: 8c010e4a                 eor w12, w12, w14
  3300: 0e0d40f9                 ldr x14, [x8, #0x18]
  3304: ce696d38                 ldrb w14, [x14, x13]
  3308: 0e010034                 cbz w14, 0xd0c
  3312: ee696e38                 ldrb w14, [x15, x14]
  3316: ce2d0111                 add w14, w14, #0x4b
  3320: dffd0371                 cmp w14, #0xff
  3324: 43000054                 b.lo 0xd04
  3328: cefd0351                 sub w14, w14, #0xff
  3332: 2e6a6e38                 ldrb w14, [x17, x14]
  3336: 8c010e4a                 eor w12, w12, w14
  3340: 0e1140f9                 ldr x14, [x8, #0x20]
  3344: ce696d38                 ldrb w14, [x14, x13]
  3348: 0e010034                 cbz w14, 0xd34
  3352: ee696e38                 ldrb w14, [x15, x14]
  3356: ce910111                 add w14, w14, #0x64
  3360: dffd0371                 cmp w14, #0xff
  3364: 43000054                 b.lo 0xd2c
  3368: cefd0351                 sub w14, w14, #0xff
  3372: 2e6a6e38                 ldrb w14, [x17, x14]
  3376: 8c010e4a                 eor w12, w12, w14
  3380: 0e1540f9                 ldr x14, [x8, #0x28]
  3384: ce696d38                 ldrb w14, [x14, x13]
  3388: 0e010034                 cbz w14, 0xd5c
  3392: ee696e38                 ldrb w14, [x15, x14]
  3396: cef50111                 add w14, w14, #0x7d
  3400: dffd0371                 cmp w14, #0xff
  3404: 43000054                 b.lo 0xd54
  3408: cefd0351                 sub w14, w14, #0xff
  3412: 2e6a6e38                 ldrb w14, [x17, x14]
  3416: 8c010e4a                 eor w12, w12, w14
  3420: 0e1940f9                 ldr x14, [x8, #0x30]
  3424: ce696d38                 ldrb w14, [x14, x13]
  3428: 0e010034                 cbz w14, 0xd84
  3432: ee696e38                 ldrb w14, [x15, x14]
  3436: ce590211                 add w14, w14, #0x96
  3440: dffd0371                 cmp w14, #0xff
  3444: 43000054                 b.lo 0xd7c
  3448: cefd0351                 sub w14, w14, #0xff
  3452: 2e6a6e38                 ldrb w14, [x17, x14]
  3456: 8c010e4a                 eor w12, w12, w14
  3460: 0e1d40f9                 ldr x14, [x8, #0x38]
  3464: ce696d38                 ldrb w14, [x14, x13]
  3468: 0e010034                 cbz w14, 0xdac
  3472: ee696e38                 ldrb w14, [x15, x14]
  3476: cebd0211                 add w14, w14, #0xaf
  3480: dffd0371                 cmp w14, #0xff
  3484: 43000054                 b.lo 0xda4
  3488: cefd0351                 sub w14, w14, #0xff
  3492: 2e6a6e38                 ldrb w14, [x17, x14]
  3496: 8c010e4a                 eor w12, w12, w14
  3500: 0e2140f9                 ldr x14, [x8, #0x40]
  3504: ce696d38                 ldrb w14, [x14, x13]
  3508: 0e010034                 cbz w14, 0xdd4
  3512: ee696e38                 ldrb w14, [x15, x14]
  3516: ce210311                 add w14, w14, #0xc8
  3520: dffd0371                 cmp w14, #0xff
  3524: 43000054                 b.lo 0xdcc
  3528: cefd0351                 sub w14, w14, #0xff
  3532: 2e6a6e38                 ldrb w14, [x17, x14]
  3536: 8c010e4a                 eor w12, w12, w14
  3540: 0e2540f9                 ldr x14, [x8, #0x48]
  3544: ce696d38                 ldrb w14, [x14, x13]
  3548: 0e010034                 cbz w14, 0xdfc
  3552: ee696e38                 ldrb w14, [x15, x14]
  3556: ce850311                 add w14, w14, #0xe1
  3560: dffd0371                 cmp w14, #0xff
  3564: 43000054                 b.lo 0xdf4
  3568: cefd0351                 sub w14, w14, #0xff
  3572: 2e6a6e38                 ldrb w14, [x17, x14]
  3576: 8c010e4a                 eor w12, w12, w14
  3580: 0e2940f9                 ldr x14, [x8, #0x50]
  3584: ce696d38                 ldrb w14, [x14, x13]
  3588: 0e010034                 cbz w14, 0xe24
  3592: ee696e38                 ldrb w14, [x15, x14]
  3596: cee90311                 add w14, w14, #0xfa
  3600: dffd0371                 cmp w14, #0xff
  3604: 43000054                 b.lo 0xe1c
  3608: cefd0351                 sub w14, w14, #0xff
  3612: 2e6a6e38                 ldrb w14, [x17, x14]
  3616: 8c010e4a                 eor w12, w12, w14
  3620: 0e2d40f9                 ldr x14, [x8, #0x58]
  3624: ce696d38                 ldrb w14, [x14, x13]
  3628: 0e010034                 cbz w14, 0xe4c
  3632: ee696e38                 ldrb w14, [x15, x14]
  3636: ce510011                 add w14, w14, #0x14
  3640: dffd0371                 cmp w14, #0xff
  3644: 43000054                 b.lo 0xe44
  3648: cefd0351                 sub w14, w14, #0xff
  3652: 2e6a6e38                 ldrb w14, [x17, x14]
  3656: 8c010e4a                 eor w12, w12, w14
  3660: 4c692d38                 strb w12, [x10, x13]
  3664: ad050091                 add x13, x13, #1
  3668: bf0102eb                 cmp x13, x2
  3672: 63f1ff54                 b.lo 0xc84
  3676: 2a0540f9                 ldr x10, [x9, #8]
  3680: ed030baa                 mov x13, x11
  3684: 0c008052                 mov w12, #0
  3688: 0e0140f9                 ldr x14, [x8]
  3692: ce696d38                 ldrb w14, [x14, x13]
  3696: 8c010e4a                 eor w12, w12, w14
  3700: 0e0540f9                 ldr x14, [x8, #8]
  3704: ce696d38                 ldrb w14, [x14, x13]
  3708: 0e010034                 cbz w14, 0xe9c
  3712: ee696e38                 ldrb w14, [x15, x14]
  3716: cec90011                 add w14, w14, #0x32
  3720: dffd0371                 cmp w14, #0xff
  3724: 43000054                 b.lo 0xe94
  3728: cefd0351                 sub w14, w14, #0xff
  3732: 2e6a6e38                 ldrb w14, [x17, x14]
  3736: 8c010e4a                 eor w12, w12, w14
  3740: 0e0940f9                 ldr x14, [x8, #0x10]
  3744: ce696d38                 ldrb w14, [x14, x13]
  3748: 0e010034                 cbz w14, 0xec4
  3752: ee696e38                 ldrb w14, [x15, x14]
  3756: ce910111                 add w14, w14, #0x64
  3760: dffd0371                 cmp w14, #0xff
  3764: 43000054                 b.lo 0xebc
  3768: cefd0351                 sub w14, w14, #0xff
  3772: 2e6a6e38                 ldrb w14, [x17, x14]
  3776: 8c010e4a                 eor w12, w12, w14
  3780: 0e0d40f9                 ldr x14, [x8, #0x18]
  3784: ce696d38                 ldrb w14, [x14, x13]
  3788: 0e010034                 cbz w14, 0xeec
  3792: ee696e38                 ldrb w14, [x15, x14]
  3796: ce590211                 add w14, w14, #0x96
  3800: dffd0371                 cmp w14, #0xff
  3804: 43000054                 b.lo 0xee4
  3808: cefd0351                 sub w14, w14, #0xff
  3812: 2e6a6e38                 ldrb w14, [x17, x14]
  3816: 8c010e4a                 eor w12, w12, w14
  3820: 0e1140f9                 ldr x14, [x8, #0x20]
  3824: ce696d38                 ldrb w14, [x14, x13]
  3828: 0e010034                 cbz w14, 0xf14
  3832: ee696e38                 ldrb w14, [x15, x14]
  3836: ce210311                 add w14, w14, #0xc8
  3840: dffd0371                 cmp w14, #0xff
  3844: 43000054                 b.lo 0xf0c
  3848: cefd0351                 sub w14, w14, #0xff
  3852: 2e6a6e38                 ldrb w14, [x17, x14]
  3856: 8c010e4a                 eor w12, w12, w14
  3860: 0e1540f9                 ldr x14, [x8, #0x28]
  3864: ce696d38                 ldrb w14, [x14, x13]
  3868: 0e010034                 cbz w14, 0xf3c
  3872: ee696e38                 ldrb w14, [x15, x14]
  3876: cee90311                 add w14, w14, #0xfa
  3880: dffd0371                 cmp w14, #0xff
  3884: 43000054                 b.lo 0xf34
  3888: cefd0351                 sub w14, w14, #0xff
  3892: 2e6a6e38                 ldrb w14, [x17, x14]
  3896: 8c010e4a                 eor w12, w12, w14
  3900: 0e1940f9                 ldr x14, [x8, #0x30]
  3904: ce696d38                 ldrb w14, [x14, x13]
  3908: 0e010034                 cbz w14, 0xf64
  3912: ee696e38                 ldrb w14, [x15, x14]
  3916: ceb50011                 add w14, w14, #0x2d
  3920: dffd0371                 cmp w14, #0xff
  3924: 43000054                 b.lo 0xf5c
  3928: cefd0351                 sub w14, w14, #0xff
  3932: 2e6a6e38                 ldrb w14, [x17, x14]
  3936: 8c010e4a                 eor w12, w12, w14
  3940: 0e1d40f9                 ldr x14, [x8, #0x38]
  3944: ce696d38                 ldrb w14, [x14, x13]
  3948: 0e010034                 cbz w14, 0xf8c
  3952: ee696e38                 ldrb w14, [x15, x14]
  3956: ce7d0111                 add w14, w14, #0x5f
  3960: dffd0371                 cmp w14, #0xff
  3964: 43000054                 b.lo 0xf84
  3968: cefd0351                 sub w14, w14, #0xff
  3972: 2e6a6e38                 ldrb w14, [x17, x14]
  3976: 8c010e4a                 eor w12, w12, w14
  3980: 0e2140f9                 ldr x14, [x8, #0x40]
  3984: ce696d38                 ldrb w14, [x14, x13]
  3988: 0e010034                 cbz w14, 0xfb4
  3992: ee696e38                 ldrb w14, [x15, x14]
  3996: ce450211                 add w14, w14, #0x91
  4000: dffd0371                 cmp w14, #0xff
  4004: 43000054                 b.lo 0xfac
  4008: cefd0351                 sub w14, w14, #0xff
  4012: 2e6a6e38                 ldrb w14, [x17, x14]
  4016: 8c010e4a                 eor w12, w12, w14
  4020: 0e2540f9                 ldr x14, [x8, #0x48]
  4024: ce696d38                 ldrb w14, [x14, x13]
  4028: 0e010034                 cbz w14, 0xfdc
  4032: ee696e38                 ldrb w14, [x15, x14]
  4036: ce0d0311                 add w14, w14, #0xc3
  4040: dffd0371                 cmp w14, #0xff
  4044: 43000054                 b.lo 0xfd4
  4048: cefd0351                 sub w14, w14, #0xff
  4052: 2e6a6e38                 ldrb w14, [x17, x14]
  4056: 8c010e4a                 eor w12, w12, w14
  4060: 0e2940f9                 ldr x14, [x8, #0x50]
  4064: ce696d38                 ldrb w14, [x14, x13]
  4068: 0e010034                 cbz w14, 0x1004
  4072: ee696e38                 ldrb w14, [x15, x14]
  4076: ced50311                 add w14, w14, #0xf5
  4080: dffd0371                 cmp w14, #0xff
  4084: 43000054                 b.lo 0xffc
  4088: cefd0351                 sub w14, w14, #0xff
  4092: 2e6a6e38                 ldrb w14, [x17, x14]
  4096: 8c010e4a                 eor w12, w12, w14
  4100: 0e2d40f9                 ldr x14, [x8, #0x58]
  4104: ce696d38                 ldrb w14, [x14, x13]
  4108: 0e010034                 cbz w14, 0x102c
  4112: ee696e38                 ldrb w14, [x15, x14]
  4116: cea10011                 add w14, w14, #0x28
  4120: dffd0371                 cmp w14, #0xff
  4124: 43000054                 b.lo 0x1024
  4128: cefd0351                 sub w14, w14, #0xff
  4132: 2e6a6e38                 ldrb w14, [x17, x14]
  4136: 8c010e4a                 eor w12, w12, w14
  4140: 4c692d38                 strb w12, [x10, x13]
  4144: ad050091                 add x13, x13, #1
  4148: bf0102eb                 cmp x13, x2
  4152: 63f1ff54                 b.lo 0xe64
  4156: 2a0940f9                 ldr x10, [x9, #0x10]
  4160: ed030baa                 mov x13, x11
  4164: 0c008052                 mov w12, #0
  4168: 0e0140f9                 ldr x14, [x8]
  4172: ce696d38                 ldrb w14, [x14, x13]
  4176: 8c010e4a                 eor w12, w12, w14
  4180: 0e0540f9                 ldr x14, [x8, #8]
  4184: ce696d38                 ldrb w14, [x14, x13]
  4188: 8c010e4a                 eor w12, w12, w14
  4192: 0e0940f9                 ldr x14, [x8, #0x10]
  4196: ce696d38                 ldrb w14, [x14, x13]
  4200: 8c010e4a                 eor w12, w12, w14
  4204: 0e0d40f9                 ldr x14, [x8, #0x18]
  4208: ce696d38                 ldrb w14, [x14, x13]
  4212: 8c010e4a                 eor w12, w12, w14
  4216: 0e1140f9                 ldr x14, [x8, #0x20]
  4220: ce696d38                 ldrb w14, [x14, x13]
  4224: 8c010e4a                 eor w12, w12, w14
  4228: 0e1540f9                 ldr x14, [x8, #0x28]
  4232: ce696d38                 ldrb w14, [x14, x13]
  4236: 8c010e4a                 eor w12, w12, w14
  4240: 4c692d38                 strb w12, [x10, x13]
  4244: ad050091                 add x13, x13, #1
  4248: bf0102eb                 cmp x13, x2
  4252: 43fdff54                 b.lo 0x1044
  4256: 2a0d40f9                 ldr x10, [x9, #0x18]
  4260: ed030baa                 mov x13, x11
  4264: 0c008052                 mov w12, #0
  4268: 0e1940f9                 ldr x14, [x8, #0x30]
  4272: ce696d38                 ldrb w14, [x14, x13]
  4276: 8c010e4a                 eor w12, w12, w14
  4280: 0e1d40f9                 ldr x14, [x8, #0x38]
  4284: ce696d38                 ldrb w14, [x14, x13]
  4288: 8c010e4a                 eor w12, w12, w14
  4292: 0e2140f9                 ldr x14, [x8, #0x40]
  4296: ce696d38                 ldrb w14, [x14, x13]
  4300: 8c010e4a                 eor w12, w12, w14
  4304: 0e2540f9                 ldr x14, [x8, #0x48]
  4308: ce696d38                 ldrb w14, [x14, x13]
  4312: 8c010e4a                 eor w12, w12, w14
  4316: 0e2940f9                 ldr x14, [x8, #0x50]
  4320: ce696d38                 ldrb w14, [x14, x13]
  4324: 8c010e4a                 eor w12, w12, w14
  4328: 0e2d40f9                 ldr x14, [x8, #0x58]
  4332: ce696d38                 ldrb w14, [x14, x13]
  4336: 8c010e4a                 eor w12, w12, w14
  4340: 4c692d38                 strb w12, [x10, x13]
  4344: ad050091                 add x13, x13, #1
  4348: bf0102eb                 cmp x13, x2
  4352: 43fdff54                 b.lo 0x10a8
  4356: c0035fd6                 ret 
  4360: 00000000                 udf #0
  4364: 00000000                 udf #0
  4368: 00000000                 udf #0
  4372: 00000000                 udf #0
  4376: 00000000                 udf #0
  4380: 00000000                 udf #0
  4384: 00000000                 udf #0
  4388: 00000000                 udf #0
  4392: 00000000                 udf #0
  4396: 00000000                 udf #0
  4400: 00000000                 udf #0
  4404: 00000000                 udf #0
  4408: 00000000                 udf #0
  4412: 00000000                 udf #0
