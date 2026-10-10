; config_t3_2114 target=sve2 twin=baseline topo=lrc(8,2,2)/rec{0} op=recover U=4
; code_size=1344 twins=[] uops_per_byte=3.265625
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d0290010                 adr x16, 0x540
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: 8df57ed3                 lsl x13, x12, #2
    24: 4e08cd9a                 udiv x14, x2, x13
    28: e01f2c25                 whilelo p0.b, xzr, x12
    32: fbc13825                 mov z27.b, #0xf
    36: 8e1d00b4                 cbz x14, 0x3d4
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
   912: 54a100a4                 ld1b { z20.b }, p0/z, [x10]
   916: 90021904                 eor z16.b, p0/m, z16.b, z20.b
   920: 54a101a4                 ld1b { z20.b }, p0/z, [x10, #1, mul vl]
   924: 91021904                 eor z17.b, p0/m, z17.b, z20.b
   928: 54a102a4                 ld1b { z20.b }, p0/z, [x10, #2, mul vl]
   932: 92021904                 eor z18.b, p0/m, z18.b, z20.b
   936: 54a103a4                 ld1b { z20.b }, p0/z, [x10, #3, mul vl]
   940: 93021904                 eor z19.b, p0/m, z19.b, z20.b
   944: 2a0140f9                 ldr x10, [x9]
   948: 4a010b8b                 add x10, x10, x11
   952: 50e100e4                 st1b { z16.b }, p0, [x10]
   956: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
   960: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
   964: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
   968: 6b010d8b                 add x11, x11, x13
   972: ce0500f1                 subs x14, x14, #1
   976: c1e2ff54                 b.ne 0x28
   980: 7f0102eb                 cmp x11, x2
   984: 620a0054                 b.hs 0x524
   988: 0f020e91                 add x15, x16, #0x380
   992: 11021291                 add x17, x16, #0x480
   996: 2a0140f9                 ldr x10, [x9]
  1000: ed030baa                 mov x13, x11
  1004: 0c008052                 mov w12, #0
  1008: 0e0140f9                 ldr x14, [x8]
  1012: ce696d38                 ldrb w14, [x14, x13]
  1016: 0e010034                 cbz w14, 0x418
  1020: ee696e38                 ldrb w14, [x15, x14]
  1024: ce650011                 add w14, w14, #0x19
  1028: dffd0371                 cmp w14, #0xff
  1032: 43000054                 b.lo 0x410
  1036: cefd0351                 sub w14, w14, #0xff
  1040: 2e6a6e38                 ldrb w14, [x17, x14]
  1044: 8c010e4a                 eor w12, w12, w14
  1048: 0e0540f9                 ldr x14, [x8, #8]
  1052: ce696d38                 ldrb w14, [x14, x13]
  1056: 0e010034                 cbz w14, 0x440
  1060: ee696e38                 ldrb w14, [x15, x14]
  1064: cec90011                 add w14, w14, #0x32
  1068: dffd0371                 cmp w14, #0xff
  1072: 43000054                 b.lo 0x438
  1076: cefd0351                 sub w14, w14, #0xff
  1080: 2e6a6e38                 ldrb w14, [x17, x14]
  1084: 8c010e4a                 eor w12, w12, w14
  1088: 0e0940f9                 ldr x14, [x8, #0x10]
  1092: ce696d38                 ldrb w14, [x14, x13]
  1096: 0e010034                 cbz w14, 0x468
  1100: ee696e38                 ldrb w14, [x15, x14]
  1104: ce2d0111                 add w14, w14, #0x4b
  1108: dffd0371                 cmp w14, #0xff
  1112: 43000054                 b.lo 0x460
  1116: cefd0351                 sub w14, w14, #0xff
  1120: 2e6a6e38                 ldrb w14, [x17, x14]
  1124: 8c010e4a                 eor w12, w12, w14
  1128: 0e0d40f9                 ldr x14, [x8, #0x18]
  1132: ce696d38                 ldrb w14, [x14, x13]
  1136: 0e010034                 cbz w14, 0x490
  1140: ee696e38                 ldrb w14, [x15, x14]
  1144: ce910111                 add w14, w14, #0x64
  1148: dffd0371                 cmp w14, #0xff
  1152: 43000054                 b.lo 0x488
  1156: cefd0351                 sub w14, w14, #0xff
  1160: 2e6a6e38                 ldrb w14, [x17, x14]
  1164: 8c010e4a                 eor w12, w12, w14
  1168: 0e1140f9                 ldr x14, [x8, #0x20]
  1172: ce696d38                 ldrb w14, [x14, x13]
  1176: 0e010034                 cbz w14, 0x4b8
  1180: ee696e38                 ldrb w14, [x15, x14]
  1184: cef50111                 add w14, w14, #0x7d
  1188: dffd0371                 cmp w14, #0xff
  1192: 43000054                 b.lo 0x4b0
  1196: cefd0351                 sub w14, w14, #0xff
  1200: 2e6a6e38                 ldrb w14, [x17, x14]
  1204: 8c010e4a                 eor w12, w12, w14
  1208: 0e1540f9                 ldr x14, [x8, #0x28]
  1212: ce696d38                 ldrb w14, [x14, x13]
  1216: 0e010034                 cbz w14, 0x4e0
  1220: ee696e38                 ldrb w14, [x15, x14]
  1224: ce590211                 add w14, w14, #0x96
  1228: dffd0371                 cmp w14, #0xff
  1232: 43000054                 b.lo 0x4d8
  1236: cefd0351                 sub w14, w14, #0xff
  1240: 2e6a6e38                 ldrb w14, [x17, x14]
  1244: 8c010e4a                 eor w12, w12, w14
  1248: 0e1940f9                 ldr x14, [x8, #0x30]
  1252: ce696d38                 ldrb w14, [x14, x13]
  1256: 0e010034                 cbz w14, 0x508
  1260: ee696e38                 ldrb w14, [x15, x14]
  1264: cebd0211                 add w14, w14, #0xaf
  1268: dffd0371                 cmp w14, #0xff
  1272: 43000054                 b.lo 0x500
  1276: cefd0351                 sub w14, w14, #0xff
  1280: 2e6a6e38                 ldrb w14, [x17, x14]
  1284: 8c010e4a                 eor w12, w12, w14
  1288: 0e1d40f9                 ldr x14, [x8, #0x38]
  1292: ce696d38                 ldrb w14, [x14, x13]
  1296: 8c010e4a                 eor w12, w12, w14
  1300: 4c692d38                 strb w12, [x10, x13]
  1304: ad050091                 add x13, x13, #1
  1308: bf0102eb                 cmp x13, x2
  1312: 63f6ff54                 b.lo 0x3ec
  1316: c0035fd6                 ret 
  1320: 00000000                 udf #0
  1324: 00000000                 udf #0
  1328: 00000000                 udf #0
  1332: 00000000                 udf #0
  1336: 00000000                 udf #0
  1340: 00000000                 udf #0
