; config_t3_2116 target=neon twin=baseline topo=lrc_azure()/rec{0} op=recover U=4
; code_size=2112 twins=[] uops_per_byte=5.359375
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: e80300aa                 mov x8, x0
    12: e90301aa                 mov x9, x1
    16: 90410010                 adr x16, 0x840
    20: eb031faa                 mov x11, xzr
    24: 0e0880d2                 mov x14, #0x40
    28: 4d08ce9a                 udiv x13, x2, x14
    32: 6d2f00b4                 cbz x13, 0x60c
    36: ebe5004f                 movi v11.16b, #0xf
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 0902c03d                 ldr q9, [x16]
    76: 0a12c03d                 ldr q10, [x16, #0x40]
    80: 4401c03d                 ldr q4, [x10]
    84: 851c2b4e                 and v5.16b, v4.16b, v11.16b
    88: 86040c6f                 ushr v6.16b, v4.16b, #4
    92: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
    96: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   100: e01ca74e                 mov v0.16b, v7.16b
   104: 001c286e                 eor v0.16b, v0.16b, v8.16b
   108: 4405c03d                 ldr q4, [x10, #0x10]
   112: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   116: 86040c6f                 ushr v6.16b, v4.16b, #4
   120: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   124: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   128: e11ca74e                 mov v1.16b, v7.16b
   132: 211c286e                 eor v1.16b, v1.16b, v8.16b
   136: 4409c03d                 ldr q4, [x10, #0x20]
   140: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   144: 86040c6f                 ushr v6.16b, v4.16b, #4
   148: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   152: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   156: e21ca74e                 mov v2.16b, v7.16b
   160: 421c286e                 eor v2.16b, v2.16b, v8.16b
   164: 440dc03d                 ldr q4, [x10, #0x30]
   168: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   172: 86040c6f                 ushr v6.16b, v4.16b, #4
   176: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   180: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   184: e31ca74e                 mov v3.16b, v7.16b
   188: 631c286e                 eor v3.16b, v3.16b, v8.16b
   192: 0a0540f9                 ldr x10, [x8, #8]
   196: 4a010b8b                 add x10, x10, x11
   200: 0922c03d                 ldr q9, [x16, #0x80]
   204: 0a32c03d                 ldr q10, [x16, #0xc0]
   208: 4401c03d                 ldr q4, [x10]
   212: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   216: 86040c6f                 ushr v6.16b, v4.16b, #4
   220: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   224: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   228: 001c276e                 eor v0.16b, v0.16b, v7.16b
   232: 001c286e                 eor v0.16b, v0.16b, v8.16b
   236: 4405c03d                 ldr q4, [x10, #0x10]
   240: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   244: 86040c6f                 ushr v6.16b, v4.16b, #4
   248: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   252: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   256: 211c276e                 eor v1.16b, v1.16b, v7.16b
   260: 211c286e                 eor v1.16b, v1.16b, v8.16b
   264: 4409c03d                 ldr q4, [x10, #0x20]
   268: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   272: 86040c6f                 ushr v6.16b, v4.16b, #4
   276: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   280: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   284: 421c276e                 eor v2.16b, v2.16b, v7.16b
   288: 421c286e                 eor v2.16b, v2.16b, v8.16b
   292: 440dc03d                 ldr q4, [x10, #0x30]
   296: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   300: 86040c6f                 ushr v6.16b, v4.16b, #4
   304: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   308: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   312: 631c276e                 eor v3.16b, v3.16b, v7.16b
   316: 631c286e                 eor v3.16b, v3.16b, v8.16b
   320: 0a0940f9                 ldr x10, [x8, #0x10]
   324: 4a010b8b                 add x10, x10, x11
   328: 0942c03d                 ldr q9, [x16, #0x100]
   332: 0a52c03d                 ldr q10, [x16, #0x140]
   336: 4401c03d                 ldr q4, [x10]
   340: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   344: 86040c6f                 ushr v6.16b, v4.16b, #4
   348: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   352: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   356: 001c276e                 eor v0.16b, v0.16b, v7.16b
   360: 001c286e                 eor v0.16b, v0.16b, v8.16b
   364: 4405c03d                 ldr q4, [x10, #0x10]
   368: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   372: 86040c6f                 ushr v6.16b, v4.16b, #4
   376: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   380: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   384: 211c276e                 eor v1.16b, v1.16b, v7.16b
   388: 211c286e                 eor v1.16b, v1.16b, v8.16b
   392: 4409c03d                 ldr q4, [x10, #0x20]
   396: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   400: 86040c6f                 ushr v6.16b, v4.16b, #4
   404: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   408: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   412: 421c276e                 eor v2.16b, v2.16b, v7.16b
   416: 421c286e                 eor v2.16b, v2.16b, v8.16b
   420: 440dc03d                 ldr q4, [x10, #0x30]
   424: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   428: 86040c6f                 ushr v6.16b, v4.16b, #4
   432: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   436: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   440: 631c276e                 eor v3.16b, v3.16b, v7.16b
   444: 631c286e                 eor v3.16b, v3.16b, v8.16b
   448: 0a0d40f9                 ldr x10, [x8, #0x18]
   452: 4a010b8b                 add x10, x10, x11
   456: 0962c03d                 ldr q9, [x16, #0x180]
   460: 0a72c03d                 ldr q10, [x16, #0x1c0]
   464: 4401c03d                 ldr q4, [x10]
   468: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   472: 86040c6f                 ushr v6.16b, v4.16b, #4
   476: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   480: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   484: 001c276e                 eor v0.16b, v0.16b, v7.16b
   488: 001c286e                 eor v0.16b, v0.16b, v8.16b
   492: 4405c03d                 ldr q4, [x10, #0x10]
   496: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   500: 86040c6f                 ushr v6.16b, v4.16b, #4
   504: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   508: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   512: 211c276e                 eor v1.16b, v1.16b, v7.16b
   516: 211c286e                 eor v1.16b, v1.16b, v8.16b
   520: 4409c03d                 ldr q4, [x10, #0x20]
   524: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   528: 86040c6f                 ushr v6.16b, v4.16b, #4
   532: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   536: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   540: 421c276e                 eor v2.16b, v2.16b, v7.16b
   544: 421c286e                 eor v2.16b, v2.16b, v8.16b
   548: 440dc03d                 ldr q4, [x10, #0x30]
   552: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   556: 86040c6f                 ushr v6.16b, v4.16b, #4
   560: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   564: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   568: 631c276e                 eor v3.16b, v3.16b, v7.16b
   572: 631c286e                 eor v3.16b, v3.16b, v8.16b
   576: 0a1140f9                 ldr x10, [x8, #0x20]
   580: 4a010b8b                 add x10, x10, x11
   584: 0982c03d                 ldr q9, [x16, #0x200]
   588: 0a92c03d                 ldr q10, [x16, #0x240]
   592: 4401c03d                 ldr q4, [x10]
   596: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   600: 86040c6f                 ushr v6.16b, v4.16b, #4
   604: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   608: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   612: 001c276e                 eor v0.16b, v0.16b, v7.16b
   616: 001c286e                 eor v0.16b, v0.16b, v8.16b
   620: 4405c03d                 ldr q4, [x10, #0x10]
   624: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   628: 86040c6f                 ushr v6.16b, v4.16b, #4
   632: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   636: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   640: 211c276e                 eor v1.16b, v1.16b, v7.16b
   644: 211c286e                 eor v1.16b, v1.16b, v8.16b
   648: 4409c03d                 ldr q4, [x10, #0x20]
   652: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   656: 86040c6f                 ushr v6.16b, v4.16b, #4
   660: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   664: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   668: 421c276e                 eor v2.16b, v2.16b, v7.16b
   672: 421c286e                 eor v2.16b, v2.16b, v8.16b
   676: 440dc03d                 ldr q4, [x10, #0x30]
   680: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   684: 86040c6f                 ushr v6.16b, v4.16b, #4
   688: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   692: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   696: 631c276e                 eor v3.16b, v3.16b, v7.16b
   700: 631c286e                 eor v3.16b, v3.16b, v8.16b
   704: 0a1540f9                 ldr x10, [x8, #0x28]
   708: 4a010b8b                 add x10, x10, x11
   712: 09a2c03d                 ldr q9, [x16, #0x280]
   716: 0ab2c03d                 ldr q10, [x16, #0x2c0]
   720: 4401c03d                 ldr q4, [x10]
   724: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   728: 86040c6f                 ushr v6.16b, v4.16b, #4
   732: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   736: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   740: 001c276e                 eor v0.16b, v0.16b, v7.16b
   744: 001c286e                 eor v0.16b, v0.16b, v8.16b
   748: 4405c03d                 ldr q4, [x10, #0x10]
   752: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   756: 86040c6f                 ushr v6.16b, v4.16b, #4
   760: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   764: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   768: 211c276e                 eor v1.16b, v1.16b, v7.16b
   772: 211c286e                 eor v1.16b, v1.16b, v8.16b
   776: 4409c03d                 ldr q4, [x10, #0x20]
   780: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   784: 86040c6f                 ushr v6.16b, v4.16b, #4
   788: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   792: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   796: 421c276e                 eor v2.16b, v2.16b, v7.16b
   800: 421c286e                 eor v2.16b, v2.16b, v8.16b
   804: 440dc03d                 ldr q4, [x10, #0x30]
   808: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   812: 86040c6f                 ushr v6.16b, v4.16b, #4
   816: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   820: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   824: 631c276e                 eor v3.16b, v3.16b, v7.16b
   828: 631c286e                 eor v3.16b, v3.16b, v8.16b
   832: 0a1940f9                 ldr x10, [x8, #0x30]
   836: 4a010b8b                 add x10, x10, x11
   840: 09c2c03d                 ldr q9, [x16, #0x300]
   844: 0ad2c03d                 ldr q10, [x16, #0x340]
   848: 4401c03d                 ldr q4, [x10]
   852: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   856: 86040c6f                 ushr v6.16b, v4.16b, #4
   860: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   864: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   868: 001c276e                 eor v0.16b, v0.16b, v7.16b
   872: 001c286e                 eor v0.16b, v0.16b, v8.16b
   876: 4405c03d                 ldr q4, [x10, #0x10]
   880: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   884: 86040c6f                 ushr v6.16b, v4.16b, #4
   888: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   892: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   896: 211c276e                 eor v1.16b, v1.16b, v7.16b
   900: 211c286e                 eor v1.16b, v1.16b, v8.16b
   904: 4409c03d                 ldr q4, [x10, #0x20]
   908: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   912: 86040c6f                 ushr v6.16b, v4.16b, #4
   916: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   920: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   924: 421c276e                 eor v2.16b, v2.16b, v7.16b
   928: 421c286e                 eor v2.16b, v2.16b, v8.16b
   932: 440dc03d                 ldr q4, [x10, #0x30]
   936: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   940: 86040c6f                 ushr v6.16b, v4.16b, #4
   944: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   948: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   952: 631c276e                 eor v3.16b, v3.16b, v7.16b
   956: 631c286e                 eor v3.16b, v3.16b, v8.16b
   960: 0a1d40f9                 ldr x10, [x8, #0x38]
   964: 4a010b8b                 add x10, x10, x11
   968: 09e2c03d                 ldr q9, [x16, #0x380]
   972: 0af2c03d                 ldr q10, [x16, #0x3c0]
   976: 4401c03d                 ldr q4, [x10]
   980: 851c2b4e                 and v5.16b, v4.16b, v11.16b
   984: 86040c6f                 ushr v6.16b, v4.16b, #4
   988: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
   992: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
   996: 001c276e                 eor v0.16b, v0.16b, v7.16b
  1000: 001c286e                 eor v0.16b, v0.16b, v8.16b
  1004: 4405c03d                 ldr q4, [x10, #0x10]
  1008: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1012: 86040c6f                 ushr v6.16b, v4.16b, #4
  1016: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1020: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1024: 211c276e                 eor v1.16b, v1.16b, v7.16b
  1028: 211c286e                 eor v1.16b, v1.16b, v8.16b
  1032: 4409c03d                 ldr q4, [x10, #0x20]
  1036: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1040: 86040c6f                 ushr v6.16b, v4.16b, #4
  1044: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1048: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1052: 421c276e                 eor v2.16b, v2.16b, v7.16b
  1056: 421c286e                 eor v2.16b, v2.16b, v8.16b
  1060: 440dc03d                 ldr q4, [x10, #0x30]
  1064: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1068: 86040c6f                 ushr v6.16b, v4.16b, #4
  1072: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1076: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1080: 631c276e                 eor v3.16b, v3.16b, v7.16b
  1084: 631c286e                 eor v3.16b, v3.16b, v8.16b
  1088: 0a2140f9                 ldr x10, [x8, #0x40]
  1092: 4a010b8b                 add x10, x10, x11
  1096: 0902c13d                 ldr q9, [x16, #0x400]
  1100: 0a12c13d                 ldr q10, [x16, #0x440]
  1104: 4401c03d                 ldr q4, [x10]
  1108: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1112: 86040c6f                 ushr v6.16b, v4.16b, #4
  1116: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1120: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1124: 001c276e                 eor v0.16b, v0.16b, v7.16b
  1128: 001c286e                 eor v0.16b, v0.16b, v8.16b
  1132: 4405c03d                 ldr q4, [x10, #0x10]
  1136: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1140: 86040c6f                 ushr v6.16b, v4.16b, #4
  1144: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1148: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1152: 211c276e                 eor v1.16b, v1.16b, v7.16b
  1156: 211c286e                 eor v1.16b, v1.16b, v8.16b
  1160: 4409c03d                 ldr q4, [x10, #0x20]
  1164: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1168: 86040c6f                 ushr v6.16b, v4.16b, #4
  1172: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1176: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1180: 421c276e                 eor v2.16b, v2.16b, v7.16b
  1184: 421c286e                 eor v2.16b, v2.16b, v8.16b
  1188: 440dc03d                 ldr q4, [x10, #0x30]
  1192: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1196: 86040c6f                 ushr v6.16b, v4.16b, #4
  1200: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1204: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1208: 631c276e                 eor v3.16b, v3.16b, v7.16b
  1212: 631c286e                 eor v3.16b, v3.16b, v8.16b
  1216: 0a2540f9                 ldr x10, [x8, #0x48]
  1220: 4a010b8b                 add x10, x10, x11
  1224: 0922c13d                 ldr q9, [x16, #0x480]
  1228: 0a32c13d                 ldr q10, [x16, #0x4c0]
  1232: 4401c03d                 ldr q4, [x10]
  1236: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1240: 86040c6f                 ushr v6.16b, v4.16b, #4
  1244: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1248: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1252: 001c276e                 eor v0.16b, v0.16b, v7.16b
  1256: 001c286e                 eor v0.16b, v0.16b, v8.16b
  1260: 4405c03d                 ldr q4, [x10, #0x10]
  1264: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1268: 86040c6f                 ushr v6.16b, v4.16b, #4
  1272: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1276: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1280: 211c276e                 eor v1.16b, v1.16b, v7.16b
  1284: 211c286e                 eor v1.16b, v1.16b, v8.16b
  1288: 4409c03d                 ldr q4, [x10, #0x20]
  1292: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1296: 86040c6f                 ushr v6.16b, v4.16b, #4
  1300: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1304: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1308: 421c276e                 eor v2.16b, v2.16b, v7.16b
  1312: 421c286e                 eor v2.16b, v2.16b, v8.16b
  1316: 440dc03d                 ldr q4, [x10, #0x30]
  1320: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1324: 86040c6f                 ushr v6.16b, v4.16b, #4
  1328: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1332: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1336: 631c276e                 eor v3.16b, v3.16b, v7.16b
  1340: 631c286e                 eor v3.16b, v3.16b, v8.16b
  1344: 0a2940f9                 ldr x10, [x8, #0x50]
  1348: 4a010b8b                 add x10, x10, x11
  1352: 0942c13d                 ldr q9, [x16, #0x500]
  1356: 0a52c13d                 ldr q10, [x16, #0x540]
  1360: 4401c03d                 ldr q4, [x10]
  1364: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1368: 86040c6f                 ushr v6.16b, v4.16b, #4
  1372: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1376: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1380: 001c276e                 eor v0.16b, v0.16b, v7.16b
  1384: 001c286e                 eor v0.16b, v0.16b, v8.16b
  1388: 4405c03d                 ldr q4, [x10, #0x10]
  1392: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1396: 86040c6f                 ushr v6.16b, v4.16b, #4
  1400: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1404: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1408: 211c276e                 eor v1.16b, v1.16b, v7.16b
  1412: 211c286e                 eor v1.16b, v1.16b, v8.16b
  1416: 4409c03d                 ldr q4, [x10, #0x20]
  1420: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1424: 86040c6f                 ushr v6.16b, v4.16b, #4
  1428: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1432: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1436: 421c276e                 eor v2.16b, v2.16b, v7.16b
  1440: 421c286e                 eor v2.16b, v2.16b, v8.16b
  1444: 440dc03d                 ldr q4, [x10, #0x30]
  1448: 851c2b4e                 and v5.16b, v4.16b, v11.16b
  1452: 86040c6f                 ushr v6.16b, v4.16b, #4
  1456: 2701054e                 tbl v7.16b, { v9.16b }, v5.16b
  1460: 4801064e                 tbl v8.16b, { v10.16b }, v6.16b
  1464: 631c276e                 eor v3.16b, v3.16b, v7.16b
  1468: 631c286e                 eor v3.16b, v3.16b, v8.16b
  1472: 0a2d40f9                 ldr x10, [x8, #0x58]
  1476: 4a010b8b                 add x10, x10, x11
  1480: 4401c03d                 ldr q4, [x10]
  1484: 001c246e                 eor v0.16b, v0.16b, v4.16b
  1488: 4405c03d                 ldr q4, [x10, #0x10]
  1492: 211c246e                 eor v1.16b, v1.16b, v4.16b
  1496: 4409c03d                 ldr q4, [x10, #0x20]
  1500: 421c246e                 eor v2.16b, v2.16b, v4.16b
  1504: 440dc03d                 ldr q4, [x10, #0x30]
  1508: 631c246e                 eor v3.16b, v3.16b, v4.16b
  1512: 2a0140f9                 ldr x10, [x9]
  1516: 4a010b8b                 add x10, x10, x11
  1520: 4001803d                 str q0, [x10]
  1524: 4105803d                 str q1, [x10, #0x10]
  1528: 4209803d                 str q2, [x10, #0x20]
  1532: 430d803d                 str q3, [x10, #0x30]
  1536: 6b010191                 add x11, x11, #0x40
  1540: ad0500f1                 subs x13, x13, #1
  1544: 01d1ff54                 b.ne 0x28
  1548: 7f0102eb                 cmp x11, x2
  1552: 620f0054                 b.hs 0x7fc
  1556: 0f021691                 add x15, x16, #0x580
  1560: 11021a91                 add x17, x16, #0x680
  1564: 2a0140f9                 ldr x10, [x9]
  1568: ed030baa                 mov x13, x11
  1572: 0c008052                 mov w12, #0
  1576: 0e0140f9                 ldr x14, [x8]
  1580: ce696d38                 ldrb w14, [x14, x13]
  1584: 0e010034                 cbz w14, 0x650
  1588: ee696e38                 ldrb w14, [x15, x14]
  1592: ce650011                 add w14, w14, #0x19
  1596: dffd0371                 cmp w14, #0xff
  1600: 43000054                 b.lo 0x648
  1604: cefd0351                 sub w14, w14, #0xff
  1608: 2e6a6e38                 ldrb w14, [x17, x14]
  1612: 8c010e4a                 eor w12, w12, w14
  1616: 0e0540f9                 ldr x14, [x8, #8]
  1620: ce696d38                 ldrb w14, [x14, x13]
  1624: 0e010034                 cbz w14, 0x678
  1628: ee696e38                 ldrb w14, [x15, x14]
  1632: cec90011                 add w14, w14, #0x32
  1636: dffd0371                 cmp w14, #0xff
  1640: 43000054                 b.lo 0x670
  1644: cefd0351                 sub w14, w14, #0xff
  1648: 2e6a6e38                 ldrb w14, [x17, x14]
  1652: 8c010e4a                 eor w12, w12, w14
  1656: 0e0940f9                 ldr x14, [x8, #0x10]
  1660: ce696d38                 ldrb w14, [x14, x13]
  1664: 0e010034                 cbz w14, 0x6a0
  1668: ee696e38                 ldrb w14, [x15, x14]
  1672: ce2d0111                 add w14, w14, #0x4b
  1676: dffd0371                 cmp w14, #0xff
  1680: 43000054                 b.lo 0x698
  1684: cefd0351                 sub w14, w14, #0xff
  1688: 2e6a6e38                 ldrb w14, [x17, x14]
  1692: 8c010e4a                 eor w12, w12, w14
  1696: 0e0d40f9                 ldr x14, [x8, #0x18]
  1700: ce696d38                 ldrb w14, [x14, x13]
  1704: 0e010034                 cbz w14, 0x6c8
  1708: ee696e38                 ldrb w14, [x15, x14]
  1712: ce910111                 add w14, w14, #0x64
  1716: dffd0371                 cmp w14, #0xff
  1720: 43000054                 b.lo 0x6c0
  1724: cefd0351                 sub w14, w14, #0xff
  1728: 2e6a6e38                 ldrb w14, [x17, x14]
  1732: 8c010e4a                 eor w12, w12, w14
  1736: 0e1140f9                 ldr x14, [x8, #0x20]
  1740: ce696d38                 ldrb w14, [x14, x13]
  1744: 0e010034                 cbz w14, 0x6f0
  1748: ee696e38                 ldrb w14, [x15, x14]
  1752: cef50111                 add w14, w14, #0x7d
  1756: dffd0371                 cmp w14, #0xff
  1760: 43000054                 b.lo 0x6e8
  1764: cefd0351                 sub w14, w14, #0xff
  1768: 2e6a6e38                 ldrb w14, [x17, x14]
  1772: 8c010e4a                 eor w12, w12, w14
  1776: 0e1540f9                 ldr x14, [x8, #0x28]
  1780: ce696d38                 ldrb w14, [x14, x13]
  1784: 0e010034                 cbz w14, 0x718
  1788: ee696e38                 ldrb w14, [x15, x14]
  1792: ce590211                 add w14, w14, #0x96
  1796: dffd0371                 cmp w14, #0xff
  1800: 43000054                 b.lo 0x710
  1804: cefd0351                 sub w14, w14, #0xff
  1808: 2e6a6e38                 ldrb w14, [x17, x14]
  1812: 8c010e4a                 eor w12, w12, w14
  1816: 0e1940f9                 ldr x14, [x8, #0x30]
  1820: ce696d38                 ldrb w14, [x14, x13]
  1824: 0e010034                 cbz w14, 0x740
  1828: ee696e38                 ldrb w14, [x15, x14]
  1832: cebd0211                 add w14, w14, #0xaf
  1836: dffd0371                 cmp w14, #0xff
  1840: 43000054                 b.lo 0x738
  1844: cefd0351                 sub w14, w14, #0xff
  1848: 2e6a6e38                 ldrb w14, [x17, x14]
  1852: 8c010e4a                 eor w12, w12, w14
  1856: 0e1d40f9                 ldr x14, [x8, #0x38]
  1860: ce696d38                 ldrb w14, [x14, x13]
  1864: 0e010034                 cbz w14, 0x768
  1868: ee696e38                 ldrb w14, [x15, x14]
  1872: ce210311                 add w14, w14, #0xc8
  1876: dffd0371                 cmp w14, #0xff
  1880: 43000054                 b.lo 0x760
  1884: cefd0351                 sub w14, w14, #0xff
  1888: 2e6a6e38                 ldrb w14, [x17, x14]
  1892: 8c010e4a                 eor w12, w12, w14
  1896: 0e2140f9                 ldr x14, [x8, #0x40]
  1900: ce696d38                 ldrb w14, [x14, x13]
  1904: 0e010034                 cbz w14, 0x790
  1908: ee696e38                 ldrb w14, [x15, x14]
  1912: ce850311                 add w14, w14, #0xe1
  1916: dffd0371                 cmp w14, #0xff
  1920: 43000054                 b.lo 0x788
  1924: cefd0351                 sub w14, w14, #0xff
  1928: 2e6a6e38                 ldrb w14, [x17, x14]
  1932: 8c010e4a                 eor w12, w12, w14
  1936: 0e2540f9                 ldr x14, [x8, #0x48]
  1940: ce696d38                 ldrb w14, [x14, x13]
  1944: 0e010034                 cbz w14, 0x7b8
  1948: ee696e38                 ldrb w14, [x15, x14]
  1952: cee90311                 add w14, w14, #0xfa
  1956: dffd0371                 cmp w14, #0xff
  1960: 43000054                 b.lo 0x7b0
  1964: cefd0351                 sub w14, w14, #0xff
  1968: 2e6a6e38                 ldrb w14, [x17, x14]
  1972: 8c010e4a                 eor w12, w12, w14
  1976: 0e2940f9                 ldr x14, [x8, #0x50]
  1980: ce696d38                 ldrb w14, [x14, x13]
  1984: 0e010034                 cbz w14, 0x7e0
  1988: ee696e38                 ldrb w14, [x15, x14]
  1992: ce510011                 add w14, w14, #0x14
  1996: dffd0371                 cmp w14, #0xff
  2000: 43000054                 b.lo 0x7d8
  2004: cefd0351                 sub w14, w14, #0xff
  2008: 2e6a6e38                 ldrb w14, [x17, x14]
  2012: 8c010e4a                 eor w12, w12, w14
  2016: 0e2d40f9                 ldr x14, [x8, #0x58]
  2020: ce696d38                 ldrb w14, [x14, x13]
  2024: 8c010e4a                 eor w12, w12, w14
  2028: 4c692d38                 strb w12, [x10, x13]
  2032: ad050091                 add x13, x13, #1
  2036: bf0102eb                 cmp x13, x2
  2040: 63f1ff54                 b.lo 0x624
  2044: ea2fc1ac                 ldp q10, q11, [sp], #0x20
  2048: e827c1ac                 ldp q8, q9, [sp], #0x20
  2052: c0035fd6                 ret 
  2056: 00000000                 udf #0
  2060: 00000000                 udf #0
  2064: 00000000                 udf #0
  2068: 00000000                 udf #0
  2072: 00000000                 udf #0
  2076: 00000000                 udf #0
  2080: 00000000                 udf #0
  2084: 00000000                 udf #0
  2088: 00000000                 udf #0
  2092: 00000000                 udf #0
  2096: 00000000                 udf #0
  2100: 00000000                 udf #0
  2104: 00000000                 udf #0
  2108: 00000000                 udf #0
