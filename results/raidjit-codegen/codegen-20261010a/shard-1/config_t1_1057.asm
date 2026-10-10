; config_t1_1057 target=sve2 twin=notail topo=rs(6,3)/rec{0,1} op=recover U=8
; code_size=2752 twins=["notail"] uops_per_byte=5.0078125
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0550010                 adr x16, 0xac0
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: 8df17dd3                 lsl x13, x12, #3
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: e7c13825                 mov z7.b, #0xf
    36: 4e5300b4                 cbz x14, 0xa8c
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 0f020091                 add x15, x16, #0
    76: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
    80: 0f020191                 add x15, x16, #0x40
    84: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
    88: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
    92: 01302704                 and z1.d, z0.d, z7.d
    96: 02942c04                 lsr z2.b, z0.b, #4
   100: a3302105                 tbl z3.b, { z5.b }, z1.b
   104: c4302205                 tbl z4.b, { z6.b }, z2.b
   108: 7030a404                 eor z16.d, z3.d, z4.d
   112: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   116: 01302704                 and z1.d, z0.d, z7.d
   120: 02942c04                 lsr z2.b, z0.b, #4
   124: a3302105                 tbl z3.b, { z5.b }, z1.b
   128: c4302205                 tbl z4.b, { z6.b }, z2.b
   132: 7130a404                 eor z17.d, z3.d, z4.d
   136: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   140: 01302704                 and z1.d, z0.d, z7.d
   144: 02942c04                 lsr z2.b, z0.b, #4
   148: a3302105                 tbl z3.b, { z5.b }, z1.b
   152: c4302205                 tbl z4.b, { z6.b }, z2.b
   156: 7230a404                 eor z18.d, z3.d, z4.d
   160: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   164: 01302704                 and z1.d, z0.d, z7.d
   168: 02942c04                 lsr z2.b, z0.b, #4
   172: a3302105                 tbl z3.b, { z5.b }, z1.b
   176: c4302205                 tbl z4.b, { z6.b }, z2.b
   180: 7330a404                 eor z19.d, z3.d, z4.d
   184: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   188: 01302704                 and z1.d, z0.d, z7.d
   192: 02942c04                 lsr z2.b, z0.b, #4
   196: a3302105                 tbl z3.b, { z5.b }, z1.b
   200: c4302205                 tbl z4.b, { z6.b }, z2.b
   204: 7430a404                 eor z20.d, z3.d, z4.d
   208: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   212: 01302704                 and z1.d, z0.d, z7.d
   216: 02942c04                 lsr z2.b, z0.b, #4
   220: a3302105                 tbl z3.b, { z5.b }, z1.b
   224: c4302205                 tbl z4.b, { z6.b }, z2.b
   228: 7530a404                 eor z21.d, z3.d, z4.d
   232: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   236: 01302704                 and z1.d, z0.d, z7.d
   240: 02942c04                 lsr z2.b, z0.b, #4
   244: a3302105                 tbl z3.b, { z5.b }, z1.b
   248: c4302205                 tbl z4.b, { z6.b }, z2.b
   252: 7630a404                 eor z22.d, z3.d, z4.d
   256: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   260: 01302704                 and z1.d, z0.d, z7.d
   264: 02942c04                 lsr z2.b, z0.b, #4
   268: a3302105                 tbl z3.b, { z5.b }, z1.b
   272: c4302205                 tbl z4.b, { z6.b }, z2.b
   276: 7730a404                 eor z23.d, z3.d, z4.d
   280: 0f020291                 add x15, x16, #0x80
   284: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   288: 0f020391                 add x15, x16, #0xc0
   292: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   296: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   300: 01302704                 and z1.d, z0.d, z7.d
   304: 02942c04                 lsr z2.b, z0.b, #4
   308: a3302105                 tbl z3.b, { z5.b }, z1.b
   312: c4302205                 tbl z4.b, { z6.b }, z2.b
   316: 7830a404                 eor z24.d, z3.d, z4.d
   320: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   324: 01302704                 and z1.d, z0.d, z7.d
   328: 02942c04                 lsr z2.b, z0.b, #4
   332: a3302105                 tbl z3.b, { z5.b }, z1.b
   336: c4302205                 tbl z4.b, { z6.b }, z2.b
   340: 7930a404                 eor z25.d, z3.d, z4.d
   344: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   348: 01302704                 and z1.d, z0.d, z7.d
   352: 02942c04                 lsr z2.b, z0.b, #4
   356: a3302105                 tbl z3.b, { z5.b }, z1.b
   360: c4302205                 tbl z4.b, { z6.b }, z2.b
   364: 7a30a404                 eor z26.d, z3.d, z4.d
   368: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   372: 01302704                 and z1.d, z0.d, z7.d
   376: 02942c04                 lsr z2.b, z0.b, #4
   380: a3302105                 tbl z3.b, { z5.b }, z1.b
   384: c4302205                 tbl z4.b, { z6.b }, z2.b
   388: 7b30a404                 eor z27.d, z3.d, z4.d
   392: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   396: 01302704                 and z1.d, z0.d, z7.d
   400: 02942c04                 lsr z2.b, z0.b, #4
   404: a3302105                 tbl z3.b, { z5.b }, z1.b
   408: c4302205                 tbl z4.b, { z6.b }, z2.b
   412: 7c30a404                 eor z28.d, z3.d, z4.d
   416: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   420: 01302704                 and z1.d, z0.d, z7.d
   424: 02942c04                 lsr z2.b, z0.b, #4
   428: a3302105                 tbl z3.b, { z5.b }, z1.b
   432: c4302205                 tbl z4.b, { z6.b }, z2.b
   436: 7d30a404                 eor z29.d, z3.d, z4.d
   440: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   444: 01302704                 and z1.d, z0.d, z7.d
   448: 02942c04                 lsr z2.b, z0.b, #4
   452: a3302105                 tbl z3.b, { z5.b }, z1.b
   456: c4302205                 tbl z4.b, { z6.b }, z2.b
   460: 7e30a404                 eor z30.d, z3.d, z4.d
   464: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   468: 01302704                 and z1.d, z0.d, z7.d
   472: 02942c04                 lsr z2.b, z0.b, #4
   476: a3302105                 tbl z3.b, { z5.b }, z1.b
   480: c4302205                 tbl z4.b, { z6.b }, z2.b
   484: 7f30a404                 eor z31.d, z3.d, z4.d
   488: 0a0540f9                 ldr x10, [x8, #8]
   492: 4a010b8b                 add x10, x10, x11
   496: 0f020491                 add x15, x16, #0x100
   500: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   504: 0f020591                 add x15, x16, #0x140
   508: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   512: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   516: 01302704                 and z1.d, z0.d, z7.d
   520: 02942c04                 lsr z2.b, z0.b, #4
   524: a3302105                 tbl z3.b, { z5.b }, z1.b
   528: c4302205                 tbl z4.b, { z6.b }, z2.b
   532: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
   536: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   540: 01302704                 and z1.d, z0.d, z7.d
   544: 02942c04                 lsr z2.b, z0.b, #4
   548: a3302105                 tbl z3.b, { z5.b }, z1.b
   552: c4302205                 tbl z4.b, { z6.b }, z2.b
   556: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
   560: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   564: 01302704                 and z1.d, z0.d, z7.d
   568: 02942c04                 lsr z2.b, z0.b, #4
   572: a3302105                 tbl z3.b, { z5.b }, z1.b
   576: c4302205                 tbl z4.b, { z6.b }, z2.b
   580: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
   584: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   588: 01302704                 and z1.d, z0.d, z7.d
   592: 02942c04                 lsr z2.b, z0.b, #4
   596: a3302105                 tbl z3.b, { z5.b }, z1.b
   600: c4302205                 tbl z4.b, { z6.b }, z2.b
   604: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
   608: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   612: 01302704                 and z1.d, z0.d, z7.d
   616: 02942c04                 lsr z2.b, z0.b, #4
   620: a3302105                 tbl z3.b, { z5.b }, z1.b
   624: c4302205                 tbl z4.b, { z6.b }, z2.b
   628: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
   632: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   636: 01302704                 and z1.d, z0.d, z7.d
   640: 02942c04                 lsr z2.b, z0.b, #4
   644: a3302105                 tbl z3.b, { z5.b }, z1.b
   648: c4302205                 tbl z4.b, { z6.b }, z2.b
   652: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
   656: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   660: 01302704                 and z1.d, z0.d, z7.d
   664: 02942c04                 lsr z2.b, z0.b, #4
   668: a3302105                 tbl z3.b, { z5.b }, z1.b
   672: c4302205                 tbl z4.b, { z6.b }, z2.b
   676: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
   680: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   684: 01302704                 and z1.d, z0.d, z7.d
   688: 02942c04                 lsr z2.b, z0.b, #4
   692: a3302105                 tbl z3.b, { z5.b }, z1.b
   696: c4302205                 tbl z4.b, { z6.b }, z2.b
   700: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
   704: 0f020691                 add x15, x16, #0x180
   708: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   712: 0f020791                 add x15, x16, #0x1c0
   716: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   720: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   724: 01302704                 and z1.d, z0.d, z7.d
   728: 02942c04                 lsr z2.b, z0.b, #4
   732: a3302105                 tbl z3.b, { z5.b }, z1.b
   736: c4302205                 tbl z4.b, { z6.b }, z2.b
   740: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
   744: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   748: 01302704                 and z1.d, z0.d, z7.d
   752: 02942c04                 lsr z2.b, z0.b, #4
   756: a3302105                 tbl z3.b, { z5.b }, z1.b
   760: c4302205                 tbl z4.b, { z6.b }, z2.b
   764: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
   768: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   772: 01302704                 and z1.d, z0.d, z7.d
   776: 02942c04                 lsr z2.b, z0.b, #4
   780: a3302105                 tbl z3.b, { z5.b }, z1.b
   784: c4302205                 tbl z4.b, { z6.b }, z2.b
   788: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
   792: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
   796: 01302704                 and z1.d, z0.d, z7.d
   800: 02942c04                 lsr z2.b, z0.b, #4
   804: a3302105                 tbl z3.b, { z5.b }, z1.b
   808: c4302205                 tbl z4.b, { z6.b }, z2.b
   812: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
   816: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
   820: 01302704                 and z1.d, z0.d, z7.d
   824: 02942c04                 lsr z2.b, z0.b, #4
   828: a3302105                 tbl z3.b, { z5.b }, z1.b
   832: c4302205                 tbl z4.b, { z6.b }, z2.b
   836: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
   840: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
   844: 01302704                 and z1.d, z0.d, z7.d
   848: 02942c04                 lsr z2.b, z0.b, #4
   852: a3302105                 tbl z3.b, { z5.b }, z1.b
   856: c4302205                 tbl z4.b, { z6.b }, z2.b
   860: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
   864: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
   868: 01302704                 and z1.d, z0.d, z7.d
   872: 02942c04                 lsr z2.b, z0.b, #4
   876: a3302105                 tbl z3.b, { z5.b }, z1.b
   880: c4302205                 tbl z4.b, { z6.b }, z2.b
   884: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
   888: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
   892: 01302704                 and z1.d, z0.d, z7.d
   896: 02942c04                 lsr z2.b, z0.b, #4
   900: a3302105                 tbl z3.b, { z5.b }, z1.b
   904: c4302205                 tbl z4.b, { z6.b }, z2.b
   908: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
   912: 0a0940f9                 ldr x10, [x8, #0x10]
   916: 4a010b8b                 add x10, x10, x11
   920: 0f020891                 add x15, x16, #0x200
   924: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   928: 0f020991                 add x15, x16, #0x240
   932: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
   936: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
   940: 01302704                 and z1.d, z0.d, z7.d
   944: 02942c04                 lsr z2.b, z0.b, #4
   948: a3302105                 tbl z3.b, { z5.b }, z1.b
   952: c4302205                 tbl z4.b, { z6.b }, z2.b
   956: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
   960: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
   964: 01302704                 and z1.d, z0.d, z7.d
   968: 02942c04                 lsr z2.b, z0.b, #4
   972: a3302105                 tbl z3.b, { z5.b }, z1.b
   976: c4302205                 tbl z4.b, { z6.b }, z2.b
   980: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
   984: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
   988: 01302704                 and z1.d, z0.d, z7.d
   992: 02942c04                 lsr z2.b, z0.b, #4
   996: a3302105                 tbl z3.b, { z5.b }, z1.b
  1000: c4302205                 tbl z4.b, { z6.b }, z2.b
  1004: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  1008: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1012: 01302704                 and z1.d, z0.d, z7.d
  1016: 02942c04                 lsr z2.b, z0.b, #4
  1020: a3302105                 tbl z3.b, { z5.b }, z1.b
  1024: c4302205                 tbl z4.b, { z6.b }, z2.b
  1028: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  1032: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1036: 01302704                 and z1.d, z0.d, z7.d
  1040: 02942c04                 lsr z2.b, z0.b, #4
  1044: a3302105                 tbl z3.b, { z5.b }, z1.b
  1048: c4302205                 tbl z4.b, { z6.b }, z2.b
  1052: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  1056: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1060: 01302704                 and z1.d, z0.d, z7.d
  1064: 02942c04                 lsr z2.b, z0.b, #4
  1068: a3302105                 tbl z3.b, { z5.b }, z1.b
  1072: c4302205                 tbl z4.b, { z6.b }, z2.b
  1076: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  1080: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1084: 01302704                 and z1.d, z0.d, z7.d
  1088: 02942c04                 lsr z2.b, z0.b, #4
  1092: a3302105                 tbl z3.b, { z5.b }, z1.b
  1096: c4302205                 tbl z4.b, { z6.b }, z2.b
  1100: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  1104: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1108: 01302704                 and z1.d, z0.d, z7.d
  1112: 02942c04                 lsr z2.b, z0.b, #4
  1116: a3302105                 tbl z3.b, { z5.b }, z1.b
  1120: c4302205                 tbl z4.b, { z6.b }, z2.b
  1124: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  1128: 0f020a91                 add x15, x16, #0x280
  1132: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1136: 0f020b91                 add x15, x16, #0x2c0
  1140: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1144: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1148: 01302704                 and z1.d, z0.d, z7.d
  1152: 02942c04                 lsr z2.b, z0.b, #4
  1156: a3302105                 tbl z3.b, { z5.b }, z1.b
  1160: c4302205                 tbl z4.b, { z6.b }, z2.b
  1164: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
  1168: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1172: 01302704                 and z1.d, z0.d, z7.d
  1176: 02942c04                 lsr z2.b, z0.b, #4
  1180: a3302105                 tbl z3.b, { z5.b }, z1.b
  1184: c4302205                 tbl z4.b, { z6.b }, z2.b
  1188: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
  1192: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1196: 01302704                 and z1.d, z0.d, z7.d
  1200: 02942c04                 lsr z2.b, z0.b, #4
  1204: a3302105                 tbl z3.b, { z5.b }, z1.b
  1208: c4302205                 tbl z4.b, { z6.b }, z2.b
  1212: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
  1216: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1220: 01302704                 and z1.d, z0.d, z7.d
  1224: 02942c04                 lsr z2.b, z0.b, #4
  1228: a3302105                 tbl z3.b, { z5.b }, z1.b
  1232: c4302205                 tbl z4.b, { z6.b }, z2.b
  1236: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
  1240: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1244: 01302704                 and z1.d, z0.d, z7.d
  1248: 02942c04                 lsr z2.b, z0.b, #4
  1252: a3302105                 tbl z3.b, { z5.b }, z1.b
  1256: c4302205                 tbl z4.b, { z6.b }, z2.b
  1260: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
  1264: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1268: 01302704                 and z1.d, z0.d, z7.d
  1272: 02942c04                 lsr z2.b, z0.b, #4
  1276: a3302105                 tbl z3.b, { z5.b }, z1.b
  1280: c4302205                 tbl z4.b, { z6.b }, z2.b
  1284: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
  1288: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1292: 01302704                 and z1.d, z0.d, z7.d
  1296: 02942c04                 lsr z2.b, z0.b, #4
  1300: a3302105                 tbl z3.b, { z5.b }, z1.b
  1304: c4302205                 tbl z4.b, { z6.b }, z2.b
  1308: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
  1312: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1316: 01302704                 and z1.d, z0.d, z7.d
  1320: 02942c04                 lsr z2.b, z0.b, #4
  1324: a3302105                 tbl z3.b, { z5.b }, z1.b
  1328: c4302205                 tbl z4.b, { z6.b }, z2.b
  1332: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
  1336: 0a0d40f9                 ldr x10, [x8, #0x18]
  1340: 4a010b8b                 add x10, x10, x11
  1344: 0f020c91                 add x15, x16, #0x300
  1348: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1352: 0f020d91                 add x15, x16, #0x340
  1356: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1360: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1364: 01302704                 and z1.d, z0.d, z7.d
  1368: 02942c04                 lsr z2.b, z0.b, #4
  1372: a3302105                 tbl z3.b, { z5.b }, z1.b
  1376: c4302205                 tbl z4.b, { z6.b }, z2.b
  1380: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  1384: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1388: 01302704                 and z1.d, z0.d, z7.d
  1392: 02942c04                 lsr z2.b, z0.b, #4
  1396: a3302105                 tbl z3.b, { z5.b }, z1.b
  1400: c4302205                 tbl z4.b, { z6.b }, z2.b
  1404: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  1408: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1412: 01302704                 and z1.d, z0.d, z7.d
  1416: 02942c04                 lsr z2.b, z0.b, #4
  1420: a3302105                 tbl z3.b, { z5.b }, z1.b
  1424: c4302205                 tbl z4.b, { z6.b }, z2.b
  1428: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  1432: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1436: 01302704                 and z1.d, z0.d, z7.d
  1440: 02942c04                 lsr z2.b, z0.b, #4
  1444: a3302105                 tbl z3.b, { z5.b }, z1.b
  1448: c4302205                 tbl z4.b, { z6.b }, z2.b
  1452: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  1456: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1460: 01302704                 and z1.d, z0.d, z7.d
  1464: 02942c04                 lsr z2.b, z0.b, #4
  1468: a3302105                 tbl z3.b, { z5.b }, z1.b
  1472: c4302205                 tbl z4.b, { z6.b }, z2.b
  1476: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  1480: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1484: 01302704                 and z1.d, z0.d, z7.d
  1488: 02942c04                 lsr z2.b, z0.b, #4
  1492: a3302105                 tbl z3.b, { z5.b }, z1.b
  1496: c4302205                 tbl z4.b, { z6.b }, z2.b
  1500: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  1504: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1508: 01302704                 and z1.d, z0.d, z7.d
  1512: 02942c04                 lsr z2.b, z0.b, #4
  1516: a3302105                 tbl z3.b, { z5.b }, z1.b
  1520: c4302205                 tbl z4.b, { z6.b }, z2.b
  1524: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  1528: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1532: 01302704                 and z1.d, z0.d, z7.d
  1536: 02942c04                 lsr z2.b, z0.b, #4
  1540: a3302105                 tbl z3.b, { z5.b }, z1.b
  1544: c4302205                 tbl z4.b, { z6.b }, z2.b
  1548: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  1552: 0f020e91                 add x15, x16, #0x380
  1556: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1560: 0f020f91                 add x15, x16, #0x3c0
  1564: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1568: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1572: 01302704                 and z1.d, z0.d, z7.d
  1576: 02942c04                 lsr z2.b, z0.b, #4
  1580: a3302105                 tbl z3.b, { z5.b }, z1.b
  1584: c4302205                 tbl z4.b, { z6.b }, z2.b
  1588: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
  1592: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1596: 01302704                 and z1.d, z0.d, z7.d
  1600: 02942c04                 lsr z2.b, z0.b, #4
  1604: a3302105                 tbl z3.b, { z5.b }, z1.b
  1608: c4302205                 tbl z4.b, { z6.b }, z2.b
  1612: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
  1616: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1620: 01302704                 and z1.d, z0.d, z7.d
  1624: 02942c04                 lsr z2.b, z0.b, #4
  1628: a3302105                 tbl z3.b, { z5.b }, z1.b
  1632: c4302205                 tbl z4.b, { z6.b }, z2.b
  1636: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
  1640: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1644: 01302704                 and z1.d, z0.d, z7.d
  1648: 02942c04                 lsr z2.b, z0.b, #4
  1652: a3302105                 tbl z3.b, { z5.b }, z1.b
  1656: c4302205                 tbl z4.b, { z6.b }, z2.b
  1660: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
  1664: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1668: 01302704                 and z1.d, z0.d, z7.d
  1672: 02942c04                 lsr z2.b, z0.b, #4
  1676: a3302105                 tbl z3.b, { z5.b }, z1.b
  1680: c4302205                 tbl z4.b, { z6.b }, z2.b
  1684: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
  1688: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1692: 01302704                 and z1.d, z0.d, z7.d
  1696: 02942c04                 lsr z2.b, z0.b, #4
  1700: a3302105                 tbl z3.b, { z5.b }, z1.b
  1704: c4302205                 tbl z4.b, { z6.b }, z2.b
  1708: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
  1712: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1716: 01302704                 and z1.d, z0.d, z7.d
  1720: 02942c04                 lsr z2.b, z0.b, #4
  1724: a3302105                 tbl z3.b, { z5.b }, z1.b
  1728: c4302205                 tbl z4.b, { z6.b }, z2.b
  1732: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
  1736: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1740: 01302704                 and z1.d, z0.d, z7.d
  1744: 02942c04                 lsr z2.b, z0.b, #4
  1748: a3302105                 tbl z3.b, { z5.b }, z1.b
  1752: c4302205                 tbl z4.b, { z6.b }, z2.b
  1756: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
  1760: 0a1140f9                 ldr x10, [x8, #0x20]
  1764: 4a010b8b                 add x10, x10, x11
  1768: 0f021091                 add x15, x16, #0x400
  1772: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1776: 0f021191                 add x15, x16, #0x440
  1780: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1784: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1788: 01302704                 and z1.d, z0.d, z7.d
  1792: 02942c04                 lsr z2.b, z0.b, #4
  1796: a3302105                 tbl z3.b, { z5.b }, z1.b
  1800: c4302205                 tbl z4.b, { z6.b }, z2.b
  1804: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  1808: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  1812: 01302704                 and z1.d, z0.d, z7.d
  1816: 02942c04                 lsr z2.b, z0.b, #4
  1820: a3302105                 tbl z3.b, { z5.b }, z1.b
  1824: c4302205                 tbl z4.b, { z6.b }, z2.b
  1828: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  1832: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  1836: 01302704                 and z1.d, z0.d, z7.d
  1840: 02942c04                 lsr z2.b, z0.b, #4
  1844: a3302105                 tbl z3.b, { z5.b }, z1.b
  1848: c4302205                 tbl z4.b, { z6.b }, z2.b
  1852: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  1856: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  1860: 01302704                 and z1.d, z0.d, z7.d
  1864: 02942c04                 lsr z2.b, z0.b, #4
  1868: a3302105                 tbl z3.b, { z5.b }, z1.b
  1872: c4302205                 tbl z4.b, { z6.b }, z2.b
  1876: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  1880: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  1884: 01302704                 and z1.d, z0.d, z7.d
  1888: 02942c04                 lsr z2.b, z0.b, #4
  1892: a3302105                 tbl z3.b, { z5.b }, z1.b
  1896: c4302205                 tbl z4.b, { z6.b }, z2.b
  1900: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  1904: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  1908: 01302704                 and z1.d, z0.d, z7.d
  1912: 02942c04                 lsr z2.b, z0.b, #4
  1916: a3302105                 tbl z3.b, { z5.b }, z1.b
  1920: c4302205                 tbl z4.b, { z6.b }, z2.b
  1924: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  1928: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  1932: 01302704                 and z1.d, z0.d, z7.d
  1936: 02942c04                 lsr z2.b, z0.b, #4
  1940: a3302105                 tbl z3.b, { z5.b }, z1.b
  1944: c4302205                 tbl z4.b, { z6.b }, z2.b
  1948: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  1952: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  1956: 01302704                 and z1.d, z0.d, z7.d
  1960: 02942c04                 lsr z2.b, z0.b, #4
  1964: a3302105                 tbl z3.b, { z5.b }, z1.b
  1968: c4302205                 tbl z4.b, { z6.b }, z2.b
  1972: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  1976: 0f021291                 add x15, x16, #0x480
  1980: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  1984: 0f021391                 add x15, x16, #0x4c0
  1988: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  1992: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  1996: 01302704                 and z1.d, z0.d, z7.d
  2000: 02942c04                 lsr z2.b, z0.b, #4
  2004: a3302105                 tbl z3.b, { z5.b }, z1.b
  2008: c4302205                 tbl z4.b, { z6.b }, z2.b
  2012: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
  2016: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2020: 01302704                 and z1.d, z0.d, z7.d
  2024: 02942c04                 lsr z2.b, z0.b, #4
  2028: a3302105                 tbl z3.b, { z5.b }, z1.b
  2032: c4302205                 tbl z4.b, { z6.b }, z2.b
  2036: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
  2040: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2044: 01302704                 and z1.d, z0.d, z7.d
  2048: 02942c04                 lsr z2.b, z0.b, #4
  2052: a3302105                 tbl z3.b, { z5.b }, z1.b
  2056: c4302205                 tbl z4.b, { z6.b }, z2.b
  2060: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
  2064: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2068: 01302704                 and z1.d, z0.d, z7.d
  2072: 02942c04                 lsr z2.b, z0.b, #4
  2076: a3302105                 tbl z3.b, { z5.b }, z1.b
  2080: c4302205                 tbl z4.b, { z6.b }, z2.b
  2084: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
  2088: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  2092: 01302704                 and z1.d, z0.d, z7.d
  2096: 02942c04                 lsr z2.b, z0.b, #4
  2100: a3302105                 tbl z3.b, { z5.b }, z1.b
  2104: c4302205                 tbl z4.b, { z6.b }, z2.b
  2108: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
  2112: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  2116: 01302704                 and z1.d, z0.d, z7.d
  2120: 02942c04                 lsr z2.b, z0.b, #4
  2124: a3302105                 tbl z3.b, { z5.b }, z1.b
  2128: c4302205                 tbl z4.b, { z6.b }, z2.b
  2132: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
  2136: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  2140: 01302704                 and z1.d, z0.d, z7.d
  2144: 02942c04                 lsr z2.b, z0.b, #4
  2148: a3302105                 tbl z3.b, { z5.b }, z1.b
  2152: c4302205                 tbl z4.b, { z6.b }, z2.b
  2156: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
  2160: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  2164: 01302704                 and z1.d, z0.d, z7.d
  2168: 02942c04                 lsr z2.b, z0.b, #4
  2172: a3302105                 tbl z3.b, { z5.b }, z1.b
  2176: c4302205                 tbl z4.b, { z6.b }, z2.b
  2180: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
  2184: 0a1540f9                 ldr x10, [x8, #0x28]
  2188: 4a010b8b                 add x10, x10, x11
  2192: 0f021491                 add x15, x16, #0x500
  2196: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2200: 0f021591                 add x15, x16, #0x540
  2204: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2208: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2212: 01302704                 and z1.d, z0.d, z7.d
  2216: 02942c04                 lsr z2.b, z0.b, #4
  2220: a3302105                 tbl z3.b, { z5.b }, z1.b
  2224: c4302205                 tbl z4.b, { z6.b }, z2.b
  2228: 90382304                 eor3 z16.d, z16.d, z3.d, z4.d
  2232: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2236: 01302704                 and z1.d, z0.d, z7.d
  2240: 02942c04                 lsr z2.b, z0.b, #4
  2244: a3302105                 tbl z3.b, { z5.b }, z1.b
  2248: c4302205                 tbl z4.b, { z6.b }, z2.b
  2252: 91382304                 eor3 z17.d, z17.d, z3.d, z4.d
  2256: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2260: 01302704                 and z1.d, z0.d, z7.d
  2264: 02942c04                 lsr z2.b, z0.b, #4
  2268: a3302105                 tbl z3.b, { z5.b }, z1.b
  2272: c4302205                 tbl z4.b, { z6.b }, z2.b
  2276: 92382304                 eor3 z18.d, z18.d, z3.d, z4.d
  2280: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2284: 01302704                 and z1.d, z0.d, z7.d
  2288: 02942c04                 lsr z2.b, z0.b, #4
  2292: a3302105                 tbl z3.b, { z5.b }, z1.b
  2296: c4302205                 tbl z4.b, { z6.b }, z2.b
  2300: 93382304                 eor3 z19.d, z19.d, z3.d, z4.d
  2304: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  2308: 01302704                 and z1.d, z0.d, z7.d
  2312: 02942c04                 lsr z2.b, z0.b, #4
  2316: a3302105                 tbl z3.b, { z5.b }, z1.b
  2320: c4302205                 tbl z4.b, { z6.b }, z2.b
  2324: 94382304                 eor3 z20.d, z20.d, z3.d, z4.d
  2328: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  2332: 01302704                 and z1.d, z0.d, z7.d
  2336: 02942c04                 lsr z2.b, z0.b, #4
  2340: a3302105                 tbl z3.b, { z5.b }, z1.b
  2344: c4302205                 tbl z4.b, { z6.b }, z2.b
  2348: 95382304                 eor3 z21.d, z21.d, z3.d, z4.d
  2352: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  2356: 01302704                 and z1.d, z0.d, z7.d
  2360: 02942c04                 lsr z2.b, z0.b, #4
  2364: a3302105                 tbl z3.b, { z5.b }, z1.b
  2368: c4302205                 tbl z4.b, { z6.b }, z2.b
  2372: 96382304                 eor3 z22.d, z22.d, z3.d, z4.d
  2376: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  2380: 01302704                 and z1.d, z0.d, z7.d
  2384: 02942c04                 lsr z2.b, z0.b, #4
  2388: a3302105                 tbl z3.b, { z5.b }, z1.b
  2392: c4302205                 tbl z4.b, { z6.b }, z2.b
  2396: 97382304                 eor3 z23.d, z23.d, z3.d, z4.d
  2400: 0f021491                 add x15, x16, #0x500
  2404: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
  2408: 0f021591                 add x15, x16, #0x540
  2412: e62100a4                 ld1rqb { z6.b }, p0/z, [x15]
  2416: 40a100a4                 ld1b { z0.b }, p0/z, [x10]
  2420: 01302704                 and z1.d, z0.d, z7.d
  2424: 02942c04                 lsr z2.b, z0.b, #4
  2428: a3302105                 tbl z3.b, { z5.b }, z1.b
  2432: c4302205                 tbl z4.b, { z6.b }, z2.b
  2436: 98382304                 eor3 z24.d, z24.d, z3.d, z4.d
  2440: 40a101a4                 ld1b { z0.b }, p0/z, [x10, #1, mul vl]
  2444: 01302704                 and z1.d, z0.d, z7.d
  2448: 02942c04                 lsr z2.b, z0.b, #4
  2452: a3302105                 tbl z3.b, { z5.b }, z1.b
  2456: c4302205                 tbl z4.b, { z6.b }, z2.b
  2460: 99382304                 eor3 z25.d, z25.d, z3.d, z4.d
  2464: 40a102a4                 ld1b { z0.b }, p0/z, [x10, #2, mul vl]
  2468: 01302704                 and z1.d, z0.d, z7.d
  2472: 02942c04                 lsr z2.b, z0.b, #4
  2476: a3302105                 tbl z3.b, { z5.b }, z1.b
  2480: c4302205                 tbl z4.b, { z6.b }, z2.b
  2484: 9a382304                 eor3 z26.d, z26.d, z3.d, z4.d
  2488: 40a103a4                 ld1b { z0.b }, p0/z, [x10, #3, mul vl]
  2492: 01302704                 and z1.d, z0.d, z7.d
  2496: 02942c04                 lsr z2.b, z0.b, #4
  2500: a3302105                 tbl z3.b, { z5.b }, z1.b
  2504: c4302205                 tbl z4.b, { z6.b }, z2.b
  2508: 9b382304                 eor3 z27.d, z27.d, z3.d, z4.d
  2512: 40a104a4                 ld1b { z0.b }, p0/z, [x10, #4, mul vl]
  2516: 01302704                 and z1.d, z0.d, z7.d
  2520: 02942c04                 lsr z2.b, z0.b, #4
  2524: a3302105                 tbl z3.b, { z5.b }, z1.b
  2528: c4302205                 tbl z4.b, { z6.b }, z2.b
  2532: 9c382304                 eor3 z28.d, z28.d, z3.d, z4.d
  2536: 40a105a4                 ld1b { z0.b }, p0/z, [x10, #5, mul vl]
  2540: 01302704                 and z1.d, z0.d, z7.d
  2544: 02942c04                 lsr z2.b, z0.b, #4
  2548: a3302105                 tbl z3.b, { z5.b }, z1.b
  2552: c4302205                 tbl z4.b, { z6.b }, z2.b
  2556: 9d382304                 eor3 z29.d, z29.d, z3.d, z4.d
  2560: 40a106a4                 ld1b { z0.b }, p0/z, [x10, #6, mul vl]
  2564: 01302704                 and z1.d, z0.d, z7.d
  2568: 02942c04                 lsr z2.b, z0.b, #4
  2572: a3302105                 tbl z3.b, { z5.b }, z1.b
  2576: c4302205                 tbl z4.b, { z6.b }, z2.b
  2580: 9e382304                 eor3 z30.d, z30.d, z3.d, z4.d
  2584: 40a107a4                 ld1b { z0.b }, p0/z, [x10, #7, mul vl]
  2588: 01302704                 and z1.d, z0.d, z7.d
  2592: 02942c04                 lsr z2.b, z0.b, #4
  2596: a3302105                 tbl z3.b, { z5.b }, z1.b
  2600: c4302205                 tbl z4.b, { z6.b }, z2.b
  2604: 9f382304                 eor3 z31.d, z31.d, z3.d, z4.d
  2608: 2a0140f9                 ldr x10, [x9]
  2612: 4a010b8b                 add x10, x10, x11
  2616: 50e100e4                 st1b { z16.b }, p0, [x10]
  2620: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
  2624: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
  2628: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
  2632: 54e104e4                 st1b { z20.b }, p0, [x10, #4, mul vl]
  2636: 55e105e4                 st1b { z21.b }, p0, [x10, #5, mul vl]
  2640: 56e106e4                 st1b { z22.b }, p0, [x10, #6, mul vl]
  2644: 57e107e4                 st1b { z23.b }, p0, [x10, #7, mul vl]
  2648: 2a0540f9                 ldr x10, [x9, #8]
  2652: 4a010b8b                 add x10, x10, x11
  2656: 58e100e4                 st1b { z24.b }, p0, [x10]
  2660: 59e101e4                 st1b { z25.b }, p0, [x10, #1, mul vl]
  2664: 5ae102e4                 st1b { z26.b }, p0, [x10, #2, mul vl]
  2668: 5be103e4                 st1b { z27.b }, p0, [x10, #3, mul vl]
  2672: 5ce104e4                 st1b { z28.b }, p0, [x10, #4, mul vl]
  2676: 5de105e4                 st1b { z29.b }, p0, [x10, #5, mul vl]
  2680: 5ee106e4                 st1b { z30.b }, p0, [x10, #6, mul vl]
  2684: 5fe107e4                 st1b { z31.b }, p0, [x10, #7, mul vl]
  2688: 6b010d8b                 add x11, x11, x13
  2692: ce0500f1                 subs x14, x14, #1
  2696: 01adff54                 b.ne 0x28
  2700: c0035fd6                 ret 
  2704: 00000000                 udf #0
  2708: 00000000                 udf #0
  2712: 00000000                 udf #0
  2716: 00000000                 udf #0
  2720: 00000000                 udf #0
  2724: 00000000                 udf #0
  2728: 00000000                 udf #0
  2732: 00000000                 udf #0
  2736: 00000000                 udf #0
  2740: 00000000                 udf #0
  2744: 00000000                 udf #0
  2748: 00000000                 udf #0
