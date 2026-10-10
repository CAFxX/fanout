; config_t1_1092 target=sve2 twin=baseline topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=1024 twins=[] uops_per_byte=1.7375
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: d01f0010                 adr x16, 0x400
    12: eb031faa                 mov x11, xzr
    16: ece32004                 cntb x12
    20: ad0080d2                 mov x13, #5
    24: 8d7d0d9b                 mul x13, x12, x13
    28: 4e08cd9a                 udiv x14, x2, x13
    32: e01f2c25                 whilelo p0.b, xzr, x12
    36: e6c13825                 mov z6.b, #0xf
    40: 6e1400b4                 cbz x14, 0x2b4
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0940f9                 ldr x10, [x8, #0x10]
    68: 4a010b8b                 add x10, x10, x11
    72: 50a100a4                 ld1b { z16.b }, p0/z, [x10]
    76: 51a101a4                 ld1b { z17.b }, p0/z, [x10, #1, mul vl]
    80: 52a102a4                 ld1b { z18.b }, p0/z, [x10, #2, mul vl]
    84: 53a103a4                 ld1b { z19.b }, p0/z, [x10, #3, mul vl]
    88: 54a104a4                 ld1b { z20.b }, p0/z, [x10, #4, mul vl]
    92: 0a0d40f9                 ldr x10, [x8, #0x18]
    96: 4a010b8b                 add x10, x10, x11
   100: 55a100a4                 ld1b { z21.b }, p0/z, [x10]
   104: 56a101a4                 ld1b { z22.b }, p0/z, [x10, #1, mul vl]
   108: 57a102a4                 ld1b { z23.b }, p0/z, [x10, #2, mul vl]
   112: 58a103a4                 ld1b { z24.b }, p0/z, [x10, #3, mul vl]
   116: 59a104a4                 ld1b { z25.b }, p0/z, [x10, #4, mul vl]
   120: 0a1140f9                 ldr x10, [x8, #0x20]
   124: 4a010b8b                 add x10, x10, x11
   128: 5aa100a4                 ld1b { z26.b }, p0/z, [x10]
   132: 5ba101a4                 ld1b { z27.b }, p0/z, [x10, #1, mul vl]
   136: 5ca102a4                 ld1b { z28.b }, p0/z, [x10, #2, mul vl]
   140: 5da103a4                 ld1b { z29.b }, p0/z, [x10, #3, mul vl]
   144: 5ea104a4                 ld1b { z30.b }, p0/z, [x10, #4, mul vl]
   148: 0a0140f9                 ldr x10, [x8]
   152: 4a010b8b                 add x10, x10, x11
   156: 030540f9                 ldr x3, [x8, #8]
   160: 63000b8b                 add x3, x3, x11
   164: 5fa100a4                 ld1b { z31.b }, p0/z, [x10]
   168: 62a000a4                 ld1b { z2.b }, p0/z, [x3]
   172: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   176: f0031904                 eor z16.b, p0/m, z16.b, z31.b
   180: 5fa101a4                 ld1b { z31.b }, p0/z, [x10, #1, mul vl]
   184: 62a001a4                 ld1b { z2.b }, p0/z, [x3, #1, mul vl]
   188: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   192: f1031904                 eor z17.b, p0/m, z17.b, z31.b
   196: 5fa102a4                 ld1b { z31.b }, p0/z, [x10, #2, mul vl]
   200: 62a002a4                 ld1b { z2.b }, p0/z, [x3, #2, mul vl]
   204: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   208: f2031904                 eor z18.b, p0/m, z18.b, z31.b
   212: 5fa103a4                 ld1b { z31.b }, p0/z, [x10, #3, mul vl]
   216: 62a003a4                 ld1b { z2.b }, p0/z, [x3, #3, mul vl]
   220: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   224: f3031904                 eor z19.b, p0/m, z19.b, z31.b
   228: 5fa104a4                 ld1b { z31.b }, p0/z, [x10, #4, mul vl]
   232: 62a004a4                 ld1b { z2.b }, p0/z, [x3, #4, mul vl]
   236: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   240: f4031904                 eor z20.b, p0/m, z20.b, z31.b
   244: 0f020091                 add x15, x16, #0
   248: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   252: 0f020191                 add x15, x16, #0x40
   256: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   260: 5fa100a4                 ld1b { z31.b }, p0/z, [x10]
   264: 62a000a4                 ld1b { z2.b }, p0/z, [x3]
   268: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   272: e0332604                 and z0.d, z31.d, z6.d
   276: e1972c04                 lsr z1.b, z31.b, #4
   280: 82302005                 tbl z2.b, { z4.b }, z0.b
   284: a3302105                 tbl z3.b, { z5.b }, z1.b
   288: 75382204                 eor3 z21.d, z21.d, z2.d, z3.d
   292: 5fa101a4                 ld1b { z31.b }, p0/z, [x10, #1, mul vl]
   296: 62a001a4                 ld1b { z2.b }, p0/z, [x3, #1, mul vl]
   300: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   304: e0332604                 and z0.d, z31.d, z6.d
   308: e1972c04                 lsr z1.b, z31.b, #4
   312: 82302005                 tbl z2.b, { z4.b }, z0.b
   316: a3302105                 tbl z3.b, { z5.b }, z1.b
   320: 76382204                 eor3 z22.d, z22.d, z2.d, z3.d
   324: 5fa102a4                 ld1b { z31.b }, p0/z, [x10, #2, mul vl]
   328: 62a002a4                 ld1b { z2.b }, p0/z, [x3, #2, mul vl]
   332: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   336: e0332604                 and z0.d, z31.d, z6.d
   340: e1972c04                 lsr z1.b, z31.b, #4
   344: 82302005                 tbl z2.b, { z4.b }, z0.b
   348: a3302105                 tbl z3.b, { z5.b }, z1.b
   352: 77382204                 eor3 z23.d, z23.d, z2.d, z3.d
   356: 5fa103a4                 ld1b { z31.b }, p0/z, [x10, #3, mul vl]
   360: 62a003a4                 ld1b { z2.b }, p0/z, [x3, #3, mul vl]
   364: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   368: e0332604                 and z0.d, z31.d, z6.d
   372: e1972c04                 lsr z1.b, z31.b, #4
   376: 82302005                 tbl z2.b, { z4.b }, z0.b
   380: a3302105                 tbl z3.b, { z5.b }, z1.b
   384: 78382204                 eor3 z24.d, z24.d, z2.d, z3.d
   388: 5fa104a4                 ld1b { z31.b }, p0/z, [x10, #4, mul vl]
   392: 62a004a4                 ld1b { z2.b }, p0/z, [x3, #4, mul vl]
   396: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   400: e0332604                 and z0.d, z31.d, z6.d
   404: e1972c04                 lsr z1.b, z31.b, #4
   408: 82302005                 tbl z2.b, { z4.b }, z0.b
   412: a3302105                 tbl z3.b, { z5.b }, z1.b
   416: 79382204                 eor3 z25.d, z25.d, z2.d, z3.d
   420: 0f020291                 add x15, x16, #0x80
   424: e42100a4                 ld1rqb { z4.b }, p0/z, [x15]
   428: 0f020391                 add x15, x16, #0xc0
   432: e52100a4                 ld1rqb { z5.b }, p0/z, [x15]
   436: 5fa100a4                 ld1b { z31.b }, p0/z, [x10]
   440: 62a000a4                 ld1b { z2.b }, p0/z, [x3]
   444: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   448: e0332604                 and z0.d, z31.d, z6.d
   452: e1972c04                 lsr z1.b, z31.b, #4
   456: 82302005                 tbl z2.b, { z4.b }, z0.b
   460: a3302105                 tbl z3.b, { z5.b }, z1.b
   464: 7a382204                 eor3 z26.d, z26.d, z2.d, z3.d
   468: 5fa101a4                 ld1b { z31.b }, p0/z, [x10, #1, mul vl]
   472: 62a001a4                 ld1b { z2.b }, p0/z, [x3, #1, mul vl]
   476: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   480: e0332604                 and z0.d, z31.d, z6.d
   484: e1972c04                 lsr z1.b, z31.b, #4
   488: 82302005                 tbl z2.b, { z4.b }, z0.b
   492: a3302105                 tbl z3.b, { z5.b }, z1.b
   496: 7b382204                 eor3 z27.d, z27.d, z2.d, z3.d
   500: 5fa102a4                 ld1b { z31.b }, p0/z, [x10, #2, mul vl]
   504: 62a002a4                 ld1b { z2.b }, p0/z, [x3, #2, mul vl]
   508: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   512: e0332604                 and z0.d, z31.d, z6.d
   516: e1972c04                 lsr z1.b, z31.b, #4
   520: 82302005                 tbl z2.b, { z4.b }, z0.b
   524: a3302105                 tbl z3.b, { z5.b }, z1.b
   528: 7c382204                 eor3 z28.d, z28.d, z2.d, z3.d
   532: 5fa103a4                 ld1b { z31.b }, p0/z, [x10, #3, mul vl]
   536: 62a003a4                 ld1b { z2.b }, p0/z, [x3, #3, mul vl]
   540: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   544: e0332604                 and z0.d, z31.d, z6.d
   548: e1972c04                 lsr z1.b, z31.b, #4
   552: 82302005                 tbl z2.b, { z4.b }, z0.b
   556: a3302105                 tbl z3.b, { z5.b }, z1.b
   560: 7d382204                 eor3 z29.d, z29.d, z2.d, z3.d
   564: 5fa104a4                 ld1b { z31.b }, p0/z, [x10, #4, mul vl]
   568: 62a004a4                 ld1b { z2.b }, p0/z, [x3, #4, mul vl]
   572: 5f001904                 eor z31.b, p0/m, z31.b, z2.b
   576: e0332604                 and z0.d, z31.d, z6.d
   580: e1972c04                 lsr z1.b, z31.b, #4
   584: 82302005                 tbl z2.b, { z4.b }, z0.b
   588: a3302105                 tbl z3.b, { z5.b }, z1.b
   592: 7e382204                 eor3 z30.d, z30.d, z2.d, z3.d
   596: 2a0140f9                 ldr x10, [x9]
   600: 4a010b8b                 add x10, x10, x11
   604: 50e100e4                 st1b { z16.b }, p0, [x10]
   608: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
   612: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
   616: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
   620: 54e104e4                 st1b { z20.b }, p0, [x10, #4, mul vl]
   624: 2a0540f9                 ldr x10, [x9, #8]
   628: 4a010b8b                 add x10, x10, x11
   632: 55e100e4                 st1b { z21.b }, p0, [x10]
   636: 56e101e4                 st1b { z22.b }, p0, [x10, #1, mul vl]
   640: 57e102e4                 st1b { z23.b }, p0, [x10, #2, mul vl]
   644: 58e103e4                 st1b { z24.b }, p0, [x10, #3, mul vl]
   648: 59e104e4                 st1b { z25.b }, p0, [x10, #4, mul vl]
   652: 2a0940f9                 ldr x10, [x9, #0x10]
   656: 4a010b8b                 add x10, x10, x11
   660: 5ae100e4                 st1b { z26.b }, p0, [x10]
   664: 5be101e4                 st1b { z27.b }, p0, [x10, #1, mul vl]
   668: 5ce102e4                 st1b { z28.b }, p0, [x10, #2, mul vl]
   672: 5de103e4                 st1b { z29.b }, p0, [x10, #3, mul vl]
   676: 5ee104e4                 st1b { z30.b }, p0, [x10, #4, mul vl]
   680: 6b010d8b                 add x11, x11, x13
   684: ce0500f1                 subs x14, x14, #1
   688: e1ebff54                 b.ne 0x2c
   692: 7f0102eb                 cmp x11, x2
   696: 22090054                 b.hs 0x3dc
   700: 0f020491                 add x15, x16, #0x100
   704: 11020891                 add x17, x16, #0x200
   708: 2a0140f9                 ldr x10, [x9]
   712: ed030baa                 mov x13, x11
   716: 0e0940f9                 ldr x14, [x8, #0x10]
   720: cc696d38                 ldrb w12, [x14, x13]
   724: 0e0140f9                 ldr x14, [x8]
   728: ce696d38                 ldrb w14, [x14, x13]
   732: 8c010e4a                 eor w12, w12, w14
   736: 0e0540f9                 ldr x14, [x8, #8]
   740: ce696d38                 ldrb w14, [x14, x13]
   744: 8c010e4a                 eor w12, w12, w14
   748: 4c692d38                 strb w12, [x10, x13]
   752: ad050091                 add x13, x13, #1
   756: bf0102eb                 cmp x13, x2
   760: a3feff54                 b.lo 0x2cc
   764: 2a0540f9                 ldr x10, [x9, #8]
   768: ed030baa                 mov x13, x11
   772: 0e0d40f9                 ldr x14, [x8, #0x18]
   776: cc696d38                 ldrb w12, [x14, x13]
   780: 0e0140f9                 ldr x14, [x8]
   784: ce696d38                 ldrb w14, [x14, x13]
   788: 0e010034                 cbz w14, 0x334
   792: ee696e38                 ldrb w14, [x15, x14]
   796: ce990311                 add w14, w14, #0xe6
   800: dffd0371                 cmp w14, #0xff
   804: 43000054                 b.lo 0x32c
   808: cefd0351                 sub w14, w14, #0xff
   812: 2e6a6e38                 ldrb w14, [x17, x14]
   816: 8c010e4a                 eor w12, w12, w14
   820: 0e0540f9                 ldr x14, [x8, #8]
   824: ce696d38                 ldrb w14, [x14, x13]
   828: 0e010034                 cbz w14, 0x35c
   832: ee696e38                 ldrb w14, [x15, x14]
   836: ce990311                 add w14, w14, #0xe6
   840: dffd0371                 cmp w14, #0xff
   844: 43000054                 b.lo 0x354
   848: cefd0351                 sub w14, w14, #0xff
   852: 2e6a6e38                 ldrb w14, [x17, x14]
   856: 8c010e4a                 eor w12, w12, w14
   860: 4c692d38                 strb w12, [x10, x13]
   864: ad050091                 add x13, x13, #1
   868: bf0102eb                 cmp x13, x2
   872: e3fcff54                 b.lo 0x304
   876: 2a0940f9                 ldr x10, [x9, #0x10]
   880: ed030baa                 mov x13, x11
   884: 0e1140f9                 ldr x14, [x8, #0x20]
   888: cc696d38                 ldrb w12, [x14, x13]
   892: 0e0140f9                 ldr x14, [x8]
   896: ce696d38                 ldrb w14, [x14, x13]
   900: 0e010034                 cbz w14, 0x3a4
   904: ee696e38                 ldrb w14, [x15, x14]
   908: ce350311                 add w14, w14, #0xcd
   912: dffd0371                 cmp w14, #0xff
   916: 43000054                 b.lo 0x39c
   920: cefd0351                 sub w14, w14, #0xff
   924: 2e6a6e38                 ldrb w14, [x17, x14]
   928: 8c010e4a                 eor w12, w12, w14
   932: 0e0540f9                 ldr x14, [x8, #8]
   936: ce696d38                 ldrb w14, [x14, x13]
   940: 0e010034                 cbz w14, 0x3cc
   944: ee696e38                 ldrb w14, [x15, x14]
   948: ce350311                 add w14, w14, #0xcd
   952: dffd0371                 cmp w14, #0xff
   956: 43000054                 b.lo 0x3c4
   960: cefd0351                 sub w14, w14, #0xff
   964: 2e6a6e38                 ldrb w14, [x17, x14]
   968: 8c010e4a                 eor w12, w12, w14
   972: 4c692d38                 strb w12, [x10, x13]
   976: ad050091                 add x13, x13, #1
   980: bf0102eb                 cmp x13, x2
   984: e3fcff54                 b.lo 0x374
   988: c0035fd6                 ret 
   992: 00000000                 udf #0
   996: 00000000                 udf #0
  1000: 00000000                 udf #0
  1004: 00000000                 udf #0
  1008: 00000000                 udf #0
  1012: 00000000                 udf #0
  1016: 00000000                 udf #0
  1020: 00000000                 udf #0
