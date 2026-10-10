; config_t3_2112 target=sve2 twin=baseline topo=lrc_azure()/rec{0} op=recover U=4
; code_size=1984 twins=[] uops_per_byte=5.015625
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d03d0010                 adr x16, 0x7c0
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: 8df57ed3                 lsl x13, x12, #2
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: fbc13825                 mov z27.b, #0xf
    36: 8e2c00b4                 cbz x14, 0x5b4
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 0f020091                 add x15, x16, #0
    76: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
    80: 0f020191                 add x15, x16, #0x40
    84: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
    88: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
    92: 95323b04                 and z21.d, z20.d, z27.d
    96: 96962c04                 lsr z22.b, z20.b, #4
   100: 37333505                 tbl z23.b, { z25.b }, z21.b
   104: 58333605                 tbl z24.b, { z26.b }, z22.b
   108: f032b804                 eor z16.d, z23.d, z24.d
   112: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
   116: 95323b04                 and z21.d, z20.d, z27.d
   120: 96962c04                 lsr z22.b, z20.b, #4
   124: 37333505                 tbl z23.b, { z25.b }, z21.b
   128: 58333605                 tbl z24.b, { z26.b }, z22.b
   132: f132b804                 eor z17.d, z23.d, z24.d
   136: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
   140: 95323b04                 and z21.d, z20.d, z27.d
   144: 96962c04                 lsr z22.b, z20.b, #4
   148: 37333505                 tbl z23.b, { z25.b }, z21.b
   152: 58333605                 tbl z24.b, { z26.b }, z22.b
   156: f232b804                 eor z18.d, z23.d, z24.d
   160: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
   164: 95323b04                 and z21.d, z20.d, z27.d
   168: 96962c04                 lsr z22.b, z20.b, #4
   172: 37333505                 tbl z23.b, { z25.b }, z21.b
   176: 58333605                 tbl z24.b, { z26.b }, z22.b
   180: f332b804                 eor z19.d, z23.d, z24.d
   184: 0a0540f9                 ldr x10, [x8, #8]
   188: 4a010b8b                 add x10, x10, x11
   192: 0f020291                 add x15, x16, #0x80
   196: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   200: 0f020391                 add x15, x16, #0xc0
   204: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   208: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
   212: 95323b04                 and z21.d, z20.d, z27.d
   216: 96962c04                 lsr z22.b, z20.b, #4
   220: 37333505                 tbl z23.b, { z25.b }, z21.b
   224: 58333605                 tbl z24.b, { z26.b }, z22.b
   228: 103b3704                 eor3 z16.d, z16.d, z23.d, z24.d
   232: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
   236: 95323b04                 and z21.d, z20.d, z27.d
   240: 96962c04                 lsr z22.b, z20.b, #4
   244: 37333505                 tbl z23.b, { z25.b }, z21.b
   248: 58333605                 tbl z24.b, { z26.b }, z22.b
   252: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   256: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
   260: 95323b04                 and z21.d, z20.d, z27.d
   264: 96962c04                 lsr z22.b, z20.b, #4
   268: 37333505                 tbl z23.b, { z25.b }, z21.b
   272: 58333605                 tbl z24.b, { z26.b }, z22.b
   276: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
   280: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
   284: 95323b04                 and z21.d, z20.d, z27.d
   288: 96962c04                 lsr z22.b, z20.b, #4
   292: 37333505                 tbl z23.b, { z25.b }, z21.b
   296: 58333605                 tbl z24.b, { z26.b }, z22.b
   300: 133b3704                 eor3 z19.d, z19.d, z23.d, z24.d
   304: 0a0940f9                 ldr x10, [x8, #0x10]
   308: 4a010b8b                 add x10, x10, x11
   312: 0f020491                 add x15, x16, #0x100
   316: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   320: 0f020591                 add x15, x16, #0x140
   324: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   328: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
   332: 95323b04                 and z21.d, z20.d, z27.d
   336: 96962c04                 lsr z22.b, z20.b, #4
   340: 37333505                 tbl z23.b, { z25.b }, z21.b
   344: 58333605                 tbl z24.b, { z26.b }, z22.b
   348: 103b3704                 eor3 z16.d, z16.d, z23.d, z24.d
   352: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
   356: 95323b04                 and z21.d, z20.d, z27.d
   360: 96962c04                 lsr z22.b, z20.b, #4
   364: 37333505                 tbl z23.b, { z25.b }, z21.b
   368: 58333605                 tbl z24.b, { z26.b }, z22.b
   372: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   376: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
   380: 95323b04                 and z21.d, z20.d, z27.d
   384: 96962c04                 lsr z22.b, z20.b, #4
   388: 37333505                 tbl z23.b, { z25.b }, z21.b
   392: 58333605                 tbl z24.b, { z26.b }, z22.b
   396: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
   400: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
   404: 95323b04                 and z21.d, z20.d, z27.d
   408: 96962c04                 lsr z22.b, z20.b, #4
   412: 37333505                 tbl z23.b, { z25.b }, z21.b
   416: 58333605                 tbl z24.b, { z26.b }, z22.b
   420: 133b3704                 eor3 z19.d, z19.d, z23.d, z24.d
   424: 0a0d40f9                 ldr x10, [x8, #0x18]
   428: 4a010b8b                 add x10, x10, x11
   432: 0f020691                 add x15, x16, #0x180
   436: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   440: 0f020791                 add x15, x16, #0x1c0
   444: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   448: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
   452: 95323b04                 and z21.d, z20.d, z27.d
   456: 96962c04                 lsr z22.b, z20.b, #4
   460: 37333505                 tbl z23.b, { z25.b }, z21.b
   464: 58333605                 tbl z24.b, { z26.b }, z22.b
   468: 103b3704                 eor3 z16.d, z16.d, z23.d, z24.d
   472: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
   476: 95323b04                 and z21.d, z20.d, z27.d
   480: 96962c04                 lsr z22.b, z20.b, #4
   484: 37333505                 tbl z23.b, { z25.b }, z21.b
   488: 58333605                 tbl z24.b, { z26.b }, z22.b
   492: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   496: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
   500: 95323b04                 and z21.d, z20.d, z27.d
   504: 96962c04                 lsr z22.b, z20.b, #4
   508: 37333505                 tbl z23.b, { z25.b }, z21.b
   512: 58333605                 tbl z24.b, { z26.b }, z22.b
   516: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
   520: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
   524: 95323b04                 and z21.d, z20.d, z27.d
   528: 96962c04                 lsr z22.b, z20.b, #4
   532: 37333505                 tbl z23.b, { z25.b }, z21.b
   536: 58333605                 tbl z24.b, { z26.b }, z22.b
   540: 133b3704                 eor3 z19.d, z19.d, z23.d, z24.d
   544: 0a1140f9                 ldr x10, [x8, #0x20]
   548: 4a010b8b                 add x10, x10, x11
   552: 0f020891                 add x15, x16, #0x200
   556: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   560: 0f020991                 add x15, x16, #0x240
   564: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   568: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
   572: 95323b04                 and z21.d, z20.d, z27.d
   576: 96962c04                 lsr z22.b, z20.b, #4
   580: 37333505                 tbl z23.b, { z25.b }, z21.b
   584: 58333605                 tbl z24.b, { z26.b }, z22.b
   588: 103b3704                 eor3 z16.d, z16.d, z23.d, z24.d
   592: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
   596: 95323b04                 and z21.d, z20.d, z27.d
   600: 96962c04                 lsr z22.b, z20.b, #4
   604: 37333505                 tbl z23.b, { z25.b }, z21.b
   608: 58333605                 tbl z24.b, { z26.b }, z22.b
   612: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   616: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
   620: 95323b04                 and z21.d, z20.d, z27.d
   624: 96962c04                 lsr z22.b, z20.b, #4
   628: 37333505                 tbl z23.b, { z25.b }, z21.b
   632: 58333605                 tbl z24.b, { z26.b }, z22.b
   636: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
   640: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
   644: 95323b04                 and z21.d, z20.d, z27.d
   648: 96962c04                 lsr z22.b, z20.b, #4
   652: 37333505                 tbl z23.b, { z25.b }, z21.b
   656: 58333605                 tbl z24.b, { z26.b }, z22.b
   660: 133b3704                 eor3 z19.d, z19.d, z23.d, z24.d
   664: 0a1540f9                 ldr x10, [x8, #0x28]
   668: 4a010b8b                 add x10, x10, x11
   672: 0f020a91                 add x15, x16, #0x280
   676: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   680: 0f020b91                 add x15, x16, #0x2c0
   684: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   688: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
   692: 95323b04                 and z21.d, z20.d, z27.d
   696: 96962c04                 lsr z22.b, z20.b, #4
   700: 37333505                 tbl z23.b, { z25.b }, z21.b
   704: 58333605                 tbl z24.b, { z26.b }, z22.b
   708: 103b3704                 eor3 z16.d, z16.d, z23.d, z24.d
   712: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
   716: 95323b04                 and z21.d, z20.d, z27.d
   720: 96962c04                 lsr z22.b, z20.b, #4
   724: 37333505                 tbl z23.b, { z25.b }, z21.b
   728: 58333605                 tbl z24.b, { z26.b }, z22.b
   732: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   736: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
   740: 95323b04                 and z21.d, z20.d, z27.d
   744: 96962c04                 lsr z22.b, z20.b, #4
   748: 37333505                 tbl z23.b, { z25.b }, z21.b
   752: 58333605                 tbl z24.b, { z26.b }, z22.b
   756: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
   760: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
   764: 95323b04                 and z21.d, z20.d, z27.d
   768: 96962c04                 lsr z22.b, z20.b, #4
   772: 37333505                 tbl z23.b, { z25.b }, z21.b
   776: 58333605                 tbl z24.b, { z26.b }, z22.b
   780: 133b3704                 eor3 z19.d, z19.d, z23.d, z24.d
   784: 0a1940f9                 ldr x10, [x8, #0x30]
   788: 4a010b8b                 add x10, x10, x11
   792: 0f020c91                 add x15, x16, #0x300
   796: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   800: 0f020d91                 add x15, x16, #0x340
   804: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   808: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
   812: 95323b04                 and z21.d, z20.d, z27.d
   816: 96962c04                 lsr z22.b, z20.b, #4
   820: 37333505                 tbl z23.b, { z25.b }, z21.b
   824: 58333605                 tbl z24.b, { z26.b }, z22.b
   828: 103b3704                 eor3 z16.d, z16.d, z23.d, z24.d
   832: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
   836: 95323b04                 and z21.d, z20.d, z27.d
   840: 96962c04                 lsr z22.b, z20.b, #4
   844: 37333505                 tbl z23.b, { z25.b }, z21.b
   848: 58333605                 tbl z24.b, { z26.b }, z22.b
   852: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   856: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
   860: 95323b04                 and z21.d, z20.d, z27.d
   864: 96962c04                 lsr z22.b, z20.b, #4
   868: 37333505                 tbl z23.b, { z25.b }, z21.b
   872: 58333605                 tbl z24.b, { z26.b }, z22.b
   876: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
   880: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
   884: 95323b04                 and z21.d, z20.d, z27.d
   888: 96962c04                 lsr z22.b, z20.b, #4
   892: 37333505                 tbl z23.b, { z25.b }, z21.b
   896: 58333605                 tbl z24.b, { z26.b }, z22.b
   900: 133b3704                 eor3 z19.d, z19.d, z23.d, z24.d
   904: 0a1d40f9                 ldr x10, [x8, #0x38]
   908: 4a010b8b                 add x10, x10, x11
   912: 0f020e91                 add x15, x16, #0x380
   916: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
   920: 0f020f91                 add x15, x16, #0x3c0
   924: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
   928: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
   932: 95323b04                 and z21.d, z20.d, z27.d
   936: 96962c04                 lsr z22.b, z20.b, #4
   940: 37333505                 tbl z23.b, { z25.b }, z21.b
   944: 58333605                 tbl z24.b, { z26.b }, z22.b
   948: 103b3704                 eor3 z16.d, z16.d, z23.d, z24.d
   952: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
   956: 95323b04                 and z21.d, z20.d, z27.d
   960: 96962c04                 lsr z22.b, z20.b, #4
   964: 37333505                 tbl z23.b, { z25.b }, z21.b
   968: 58333605                 tbl z24.b, { z26.b }, z22.b
   972: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
   976: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
   980: 95323b04                 and z21.d, z20.d, z27.d
   984: 96962c04                 lsr z22.b, z20.b, #4
   988: 37333505                 tbl z23.b, { z25.b }, z21.b
   992: 58333605                 tbl z24.b, { z26.b }, z22.b
   996: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
  1000: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
  1004: 95323b04                 and z21.d, z20.d, z27.d
  1008: 96962c04                 lsr z22.b, z20.b, #4
  1012: 37333505                 tbl z23.b, { z25.b }, z21.b
  1016: 58333605                 tbl z24.b, { z26.b }, z22.b
  1020: 133b3704                 eor3 z19.d, z19.d, z23.d, z24.d
  1024: 0a2140f9                 ldr x10, [x8, #0x40]
  1028: 4a010b8b                 add x10, x10, x11
  1032: 0f021091                 add x15, x16, #0x400
  1036: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
  1040: 0f021191                 add x15, x16, #0x440
  1044: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
  1048: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
  1052: 95323b04                 and z21.d, z20.d, z27.d
  1056: 96962c04                 lsr z22.b, z20.b, #4
  1060: 37333505                 tbl z23.b, { z25.b }, z21.b
  1064: 58333605                 tbl z24.b, { z26.b }, z22.b
  1068: 103b3704                 eor3 z16.d, z16.d, z23.d, z24.d
  1072: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
  1076: 95323b04                 and z21.d, z20.d, z27.d
  1080: 96962c04                 lsr z22.b, z20.b, #4
  1084: 37333505                 tbl z23.b, { z25.b }, z21.b
  1088: 58333605                 tbl z24.b, { z26.b }, z22.b
  1092: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
  1096: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
  1100: 95323b04                 and z21.d, z20.d, z27.d
  1104: 96962c04                 lsr z22.b, z20.b, #4
  1108: 37333505                 tbl z23.b, { z25.b }, z21.b
  1112: 58333605                 tbl z24.b, { z26.b }, z22.b
  1116: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
  1120: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
  1124: 95323b04                 and z21.d, z20.d, z27.d
  1128: 96962c04                 lsr z22.b, z20.b, #4
  1132: 37333505                 tbl z23.b, { z25.b }, z21.b
  1136: 58333605                 tbl z24.b, { z26.b }, z22.b
  1140: 133b3704                 eor3 z19.d, z19.d, z23.d, z24.d
  1144: 0a2540f9                 ldr x10, [x8, #0x48]
  1148: 4a010b8b                 add x10, x10, x11
  1152: 0f021291                 add x15, x16, #0x480
  1156: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
  1160: 0f021391                 add x15, x16, #0x4c0
  1164: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
  1168: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
  1172: 95323b04                 and z21.d, z20.d, z27.d
  1176: 96962c04                 lsr z22.b, z20.b, #4
  1180: 37333505                 tbl z23.b, { z25.b }, z21.b
  1184: 58333605                 tbl z24.b, { z26.b }, z22.b
  1188: 103b3704                 eor3 z16.d, z16.d, z23.d, z24.d
  1192: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
  1196: 95323b04                 and z21.d, z20.d, z27.d
  1200: 96962c04                 lsr z22.b, z20.b, #4
  1204: 37333505                 tbl z23.b, { z25.b }, z21.b
  1208: 58333605                 tbl z24.b, { z26.b }, z22.b
  1212: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
  1216: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
  1220: 95323b04                 and z21.d, z20.d, z27.d
  1224: 96962c04                 lsr z22.b, z20.b, #4
  1228: 37333505                 tbl z23.b, { z25.b }, z21.b
  1232: 58333605                 tbl z24.b, { z26.b }, z22.b
  1236: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
  1240: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
  1244: 95323b04                 and z21.d, z20.d, z27.d
  1248: 96962c04                 lsr z22.b, z20.b, #4
  1252: 37333505                 tbl z23.b, { z25.b }, z21.b
  1256: 58333605                 tbl z24.b, { z26.b }, z22.b
  1260: 133b3704                 eor3 z19.d, z19.d, z23.d, z24.d
  1264: 0a2940f9                 ldr x10, [x8, #0x50]
  1268: 4a010b8b                 add x10, x10, x11
  1272: 0f021491                 add x15, x16, #0x500
  1276: f92100a4                 ld1rqb { z25.b }, p0/z, [x15]
  1280: 0f021591                 add x15, x16, #0x540
  1284: fa2100a4                 ld1rqb { z26.b }, p0/z, [x15]
  1288: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
  1292: 95323b04                 and z21.d, z20.d, z27.d
  1296: 96962c04                 lsr z22.b, z20.b, #4
  1300: 37333505                 tbl z23.b, { z25.b }, z21.b
  1304: 58333605                 tbl z24.b, { z26.b }, z22.b
  1308: 103b3704                 eor3 z16.d, z16.d, z23.d, z24.d
  1312: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
  1316: 95323b04                 and z21.d, z20.d, z27.d
  1320: 96962c04                 lsr z22.b, z20.b, #4
  1324: 37333505                 tbl z23.b, { z25.b }, z21.b
  1328: 58333605                 tbl z24.b, { z26.b }, z22.b
  1332: 113b3704                 eor3 z17.d, z17.d, z23.d, z24.d
  1336: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
  1340: 95323b04                 and z21.d, z20.d, z27.d
  1344: 96962c04                 lsr z22.b, z20.b, #4
  1348: 37333505                 tbl z23.b, { z25.b }, z21.b
  1352: 58333605                 tbl z24.b, { z26.b }, z22.b
  1356: 123b3704                 eor3 z18.d, z18.d, z23.d, z24.d
  1360: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
  1364: 95323b04                 and z21.d, z20.d, z27.d
  1368: 96962c04                 lsr z22.b, z20.b, #4
  1372: 37333505                 tbl z23.b, { z25.b }, z21.b
  1376: 58333605                 tbl z24.b, { z26.b }, z22.b
  1380: 133b3704                 eor3 z19.d, z19.d, z23.d, z24.d
  1384: 0a2d40f9                 ldr x10, [x8, #0x58]
  1388: 4a010b8b                 add x10, x10, x11
  1392: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
  1396: 90021904                 eor z16.b, p0/m, z16.b, z20.b
  1400: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
  1404: 91021904                 eor z17.b, p0/m, z17.b, z20.b
  1408: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
  1412: 92021904                 eor z18.b, p0/m, z18.b, z20.b
  1416: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
  1420: 93021904                 eor z19.b, p0/m, z19.b, z20.b
  1424: 2a0140f9                 ldr x10, [x9]
  1428: 4a010b8b                 add x10, x10, x11
  1432: 50e100e4                 st1b { z16.b }, p0, [x10]
  1436: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
  1440: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
  1444: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
  1448: 6b010d8b                 add x11, x11, x13
  1452: ce0500f1                 subs x14, x14, #1
  1456: c1d3ff54                 b.ne 0x28
  1460: 7f0102eb                 cmp x11, x2
  1464: 620f0054                 b.hs 0x7a4
  1468: 0f021691                 add x15, x16, #0x580
  1472: 11021a91                 add x17, x16, #0x680
  1476: 2a0140f9                 ldr x10, [x9]
  1480: ed030baa                 mov x13, x11
  1484: 0c008052                 mov w12, #0
  1488: 0e0140f9                 ldr x14, [x8]
  1492: ce696d38                 ldrb w14, [x14, x13]
  1496: 0e010034                 cbz w14, 0x5f8
  1500: ee696e38                 ldrb w14, [x15, x14]
  1504: ce650011                 add w14, w14, #0x19
  1508: dffd0371                 cmp w14, #0xff
  1512: 43000054                 b.lo 0x5f0
  1516: cefd0351                 sub w14, w14, #0xff
  1520: 2e6a6e38                 ldrb w14, [x17, x14]
  1524: 8c010e4a                 eor w12, w12, w14
  1528: 0e0540f9                 ldr x14, [x8, #8]
  1532: ce696d38                 ldrb w14, [x14, x13]
  1536: 0e010034                 cbz w14, 0x620
  1540: ee696e38                 ldrb w14, [x15, x14]
  1544: cec90011                 add w14, w14, #0x32
  1548: dffd0371                 cmp w14, #0xff
  1552: 43000054                 b.lo 0x618
  1556: cefd0351                 sub w14, w14, #0xff
  1560: 2e6a6e38                 ldrb w14, [x17, x14]
  1564: 8c010e4a                 eor w12, w12, w14
  1568: 0e0940f9                 ldr x14, [x8, #0x10]
  1572: ce696d38                 ldrb w14, [x14, x13]
  1576: 0e010034                 cbz w14, 0x648
  1580: ee696e38                 ldrb w14, [x15, x14]
  1584: ce2d0111                 add w14, w14, #0x4b
  1588: dffd0371                 cmp w14, #0xff
  1592: 43000054                 b.lo 0x640
  1596: cefd0351                 sub w14, w14, #0xff
  1600: 2e6a6e38                 ldrb w14, [x17, x14]
  1604: 8c010e4a                 eor w12, w12, w14
  1608: 0e0d40f9                 ldr x14, [x8, #0x18]
  1612: ce696d38                 ldrb w14, [x14, x13]
  1616: 0e010034                 cbz w14, 0x670
  1620: ee696e38                 ldrb w14, [x15, x14]
  1624: ce910111                 add w14, w14, #0x64
  1628: dffd0371                 cmp w14, #0xff
  1632: 43000054                 b.lo 0x668
  1636: cefd0351                 sub w14, w14, #0xff
  1640: 2e6a6e38                 ldrb w14, [x17, x14]
  1644: 8c010e4a                 eor w12, w12, w14
  1648: 0e1140f9                 ldr x14, [x8, #0x20]
  1652: ce696d38                 ldrb w14, [x14, x13]
  1656: 0e010034                 cbz w14, 0x698
  1660: ee696e38                 ldrb w14, [x15, x14]
  1664: cef50111                 add w14, w14, #0x7d
  1668: dffd0371                 cmp w14, #0xff
  1672: 43000054                 b.lo 0x690
  1676: cefd0351                 sub w14, w14, #0xff
  1680: 2e6a6e38                 ldrb w14, [x17, x14]
  1684: 8c010e4a                 eor w12, w12, w14
  1688: 0e1540f9                 ldr x14, [x8, #0x28]
  1692: ce696d38                 ldrb w14, [x14, x13]
  1696: 0e010034                 cbz w14, 0x6c0
  1700: ee696e38                 ldrb w14, [x15, x14]
  1704: ce590211                 add w14, w14, #0x96
  1708: dffd0371                 cmp w14, #0xff
  1712: 43000054                 b.lo 0x6b8
  1716: cefd0351                 sub w14, w14, #0xff
  1720: 2e6a6e38                 ldrb w14, [x17, x14]
  1724: 8c010e4a                 eor w12, w12, w14
  1728: 0e1940f9                 ldr x14, [x8, #0x30]
  1732: ce696d38                 ldrb w14, [x14, x13]
  1736: 0e010034                 cbz w14, 0x6e8
  1740: ee696e38                 ldrb w14, [x15, x14]
  1744: cebd0211                 add w14, w14, #0xaf
  1748: dffd0371                 cmp w14, #0xff
  1752: 43000054                 b.lo 0x6e0
  1756: cefd0351                 sub w14, w14, #0xff
  1760: 2e6a6e38                 ldrb w14, [x17, x14]
  1764: 8c010e4a                 eor w12, w12, w14
  1768: 0e1d40f9                 ldr x14, [x8, #0x38]
  1772: ce696d38                 ldrb w14, [x14, x13]
  1776: 0e010034                 cbz w14, 0x710
  1780: ee696e38                 ldrb w14, [x15, x14]
  1784: ce210311                 add w14, w14, #0xc8
  1788: dffd0371                 cmp w14, #0xff
  1792: 43000054                 b.lo 0x708
  1796: cefd0351                 sub w14, w14, #0xff
  1800: 2e6a6e38                 ldrb w14, [x17, x14]
  1804: 8c010e4a                 eor w12, w12, w14
  1808: 0e2140f9                 ldr x14, [x8, #0x40]
  1812: ce696d38                 ldrb w14, [x14, x13]
  1816: 0e010034                 cbz w14, 0x738
  1820: ee696e38                 ldrb w14, [x15, x14]
  1824: ce850311                 add w14, w14, #0xe1
  1828: dffd0371                 cmp w14, #0xff
  1832: 43000054                 b.lo 0x730
  1836: cefd0351                 sub w14, w14, #0xff
  1840: 2e6a6e38                 ldrb w14, [x17, x14]
  1844: 8c010e4a                 eor w12, w12, w14
  1848: 0e2540f9                 ldr x14, [x8, #0x48]
  1852: ce696d38                 ldrb w14, [x14, x13]
  1856: 0e010034                 cbz w14, 0x760
  1860: ee696e38                 ldrb w14, [x15, x14]
  1864: cee90311                 add w14, w14, #0xfa
  1868: dffd0371                 cmp w14, #0xff
  1872: 43000054                 b.lo 0x758
  1876: cefd0351                 sub w14, w14, #0xff
  1880: 2e6a6e38                 ldrb w14, [x17, x14]
  1884: 8c010e4a                 eor w12, w12, w14
  1888: 0e2940f9                 ldr x14, [x8, #0x50]
  1892: ce696d38                 ldrb w14, [x14, x13]
  1896: 0e010034                 cbz w14, 0x788
  1900: ee696e38                 ldrb w14, [x15, x14]
  1904: ce510011                 add w14, w14, #0x14
  1908: dffd0371                 cmp w14, #0xff
  1912: 43000054                 b.lo 0x780
  1916: cefd0351                 sub w14, w14, #0xff
  1920: 2e6a6e38                 ldrb w14, [x17, x14]
  1924: 8c010e4a                 eor w12, w12, w14
  1928: 0e2d40f9                 ldr x14, [x8, #0x58]
  1932: ce696d38                 ldrb w14, [x14, x13]
  1936: 8c010e4a                 eor w12, w12, w14
  1940: 4c692d38                 strb w12, [x10, x13]
  1944: ad050091                 add x13, x13, #1
  1948: bf0102eb                 cmp x13, x2
  1952: 63f1ff54                 b.lo 0x5cc
  1956: c0035fd6                 ret 
  1960: 00000000                 udf #0
  1964: 00000000                 udf #0
  1968: 00000000                 udf #0
  1972: 00000000                 udf #0
  1976: 00000000                 udf #0
  1980: 00000000                 udf #0
