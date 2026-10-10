; config_t1_1326 target=neon twin=baseline topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1920 twins=[] uops_per_byte=4.546875
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: ec37bfad                 stp q12, q13, [sp, #-0x20]!
    12: ee3fbfad                 stp q14, q15, [sp, #-0x20]!
    16: e80300aa                 mov x8, x0
    20: e90301aa                 mov x9, x1
    24: 503b0010                 adr x16, 0x780
    28: eb031faa                 mov x11, xzr
    32: 0e0880d2                 mov x14, #0x40
    36: 4d08ce9a                 udiv x13, x2, x14
    40: 6d2800b4                 cbz x13, 0x534
    44: efe5004f                 movi v15.16b, #0xf
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 4801c03d                 ldr q8, [x10]
    76: 001da84e                 mov v0.16b, v8.16b
    80: 4805c03d                 ldr q8, [x10, #0x10]
    84: 011da84e                 mov v1.16b, v8.16b
    88: 4809c03d                 ldr q8, [x10, #0x20]
    92: 021da84e                 mov v2.16b, v8.16b
    96: 480dc03d                 ldr q8, [x10, #0x30]
   100: 031da84e                 mov v3.16b, v8.16b
   104: 0a0540f9                 ldr x10, [x8, #8]
   108: 4a010b8b                 add x10, x10, x11
   112: 4801c03d                 ldr q8, [x10]
   116: 001c286e                 eor v0.16b, v0.16b, v8.16b
   120: 4805c03d                 ldr q8, [x10, #0x10]
   124: 211c286e                 eor v1.16b, v1.16b, v8.16b
   128: 4809c03d                 ldr q8, [x10, #0x20]
   132: 421c286e                 eor v2.16b, v2.16b, v8.16b
   136: 480dc03d                 ldr q8, [x10, #0x30]
   140: 631c286e                 eor v3.16b, v3.16b, v8.16b
   144: 0a0940f9                 ldr x10, [x8, #0x10]
   148: 4a010b8b                 add x10, x10, x11
   152: 4801c03d                 ldr q8, [x10]
   156: 001c286e                 eor v0.16b, v0.16b, v8.16b
   160: 4805c03d                 ldr q8, [x10, #0x10]
   164: 211c286e                 eor v1.16b, v1.16b, v8.16b
   168: 4809c03d                 ldr q8, [x10, #0x20]
   172: 421c286e                 eor v2.16b, v2.16b, v8.16b
   176: 480dc03d                 ldr q8, [x10, #0x30]
   180: 631c286e                 eor v3.16b, v3.16b, v8.16b
   184: 0a0d40f9                 ldr x10, [x8, #0x18]
   188: 4a010b8b                 add x10, x10, x11
   192: 4801c03d                 ldr q8, [x10]
   196: 001c286e                 eor v0.16b, v0.16b, v8.16b
   200: 4805c03d                 ldr q8, [x10, #0x10]
   204: 211c286e                 eor v1.16b, v1.16b, v8.16b
   208: 4809c03d                 ldr q8, [x10, #0x20]
   212: 421c286e                 eor v2.16b, v2.16b, v8.16b
   216: 480dc03d                 ldr q8, [x10, #0x30]
   220: 631c286e                 eor v3.16b, v3.16b, v8.16b
   224: 0a1140f9                 ldr x10, [x8, #0x20]
   228: 4a010b8b                 add x10, x10, x11
   232: 4801c03d                 ldr q8, [x10]
   236: 001c286e                 eor v0.16b, v0.16b, v8.16b
   240: 4805c03d                 ldr q8, [x10, #0x10]
   244: 211c286e                 eor v1.16b, v1.16b, v8.16b
   248: 4809c03d                 ldr q8, [x10, #0x20]
   252: 421c286e                 eor v2.16b, v2.16b, v8.16b
   256: 480dc03d                 ldr q8, [x10, #0x30]
   260: 631c286e                 eor v3.16b, v3.16b, v8.16b
   264: 0d02c03d                 ldr q13, [x16]
   268: 0e12c03d                 ldr q14, [x16, #0x40]
   272: 0a0140f9                 ldr x10, [x8]
   276: 4a010b8b                 add x10, x10, x11
   280: 4801c03d                 ldr q8, [x10]
   284: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   288: 0a050c6f                 ushr v10.16b, v8.16b, #4
   292: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   296: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   300: 641dab4e                 mov v4.16b, v11.16b
   304: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   308: 4805c03d                 ldr q8, [x10, #0x10]
   312: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   316: 0a050c6f                 ushr v10.16b, v8.16b, #4
   320: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   324: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   328: 651dab4e                 mov v5.16b, v11.16b
   332: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   336: 4809c03d                 ldr q8, [x10, #0x20]
   340: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   344: 0a050c6f                 ushr v10.16b, v8.16b, #4
   348: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   352: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   356: 661dab4e                 mov v6.16b, v11.16b
   360: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   364: 480dc03d                 ldr q8, [x10, #0x30]
   368: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   372: 0a050c6f                 ushr v10.16b, v8.16b, #4
   376: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   380: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   384: 671dab4e                 mov v7.16b, v11.16b
   388: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   392: 0d22c03d                 ldr q13, [x16, #0x80]
   396: 0e32c03d                 ldr q14, [x16, #0xc0]
   400: 0a0540f9                 ldr x10, [x8, #8]
   404: 4a010b8b                 add x10, x10, x11
   408: 4801c03d                 ldr q8, [x10]
   412: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   416: 0a050c6f                 ushr v10.16b, v8.16b, #4
   420: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   424: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   428: 841c2b6e                 eor v4.16b, v4.16b, v11.16b
   432: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   436: 4805c03d                 ldr q8, [x10, #0x10]
   440: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   444: 0a050c6f                 ushr v10.16b, v8.16b, #4
   448: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   452: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   456: a51c2b6e                 eor v5.16b, v5.16b, v11.16b
   460: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   464: 4809c03d                 ldr q8, [x10, #0x20]
   468: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   472: 0a050c6f                 ushr v10.16b, v8.16b, #4
   476: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   480: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   484: c61c2b6e                 eor v6.16b, v6.16b, v11.16b
   488: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   492: 480dc03d                 ldr q8, [x10, #0x30]
   496: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   500: 0a050c6f                 ushr v10.16b, v8.16b, #4
   504: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   508: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   512: e71c2b6e                 eor v7.16b, v7.16b, v11.16b
   516: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   520: 0d42c03d                 ldr q13, [x16, #0x100]
   524: 0e52c03d                 ldr q14, [x16, #0x140]
   528: 0a0940f9                 ldr x10, [x8, #0x10]
   532: 4a010b8b                 add x10, x10, x11
   536: 4801c03d                 ldr q8, [x10]
   540: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   544: 0a050c6f                 ushr v10.16b, v8.16b, #4
   548: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   552: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   556: 841c2b6e                 eor v4.16b, v4.16b, v11.16b
   560: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   564: 4805c03d                 ldr q8, [x10, #0x10]
   568: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   572: 0a050c6f                 ushr v10.16b, v8.16b, #4
   576: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   580: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   584: a51c2b6e                 eor v5.16b, v5.16b, v11.16b
   588: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   592: 4809c03d                 ldr q8, [x10, #0x20]
   596: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   600: 0a050c6f                 ushr v10.16b, v8.16b, #4
   604: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   608: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   612: c61c2b6e                 eor v6.16b, v6.16b, v11.16b
   616: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   620: 480dc03d                 ldr q8, [x10, #0x30]
   624: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   628: 0a050c6f                 ushr v10.16b, v8.16b, #4
   632: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   636: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   640: e71c2b6e                 eor v7.16b, v7.16b, v11.16b
   644: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   648: 0a0d40f9                 ldr x10, [x8, #0x18]
   652: 4a010b8b                 add x10, x10, x11
   656: 4801c03d                 ldr q8, [x10]
   660: 841c286e                 eor v4.16b, v4.16b, v8.16b
   664: 4805c03d                 ldr q8, [x10, #0x10]
   668: a51c286e                 eor v5.16b, v5.16b, v8.16b
   672: 4809c03d                 ldr q8, [x10, #0x20]
   676: c61c286e                 eor v6.16b, v6.16b, v8.16b
   680: 480dc03d                 ldr q8, [x10, #0x30]
   684: e71c286e                 eor v7.16b, v7.16b, v8.16b
   688: 0a1540f9                 ldr x10, [x8, #0x28]
   692: 4a010b8b                 add x10, x10, x11
   696: 4801c03d                 ldr q8, [x10]
   700: 841c286e                 eor v4.16b, v4.16b, v8.16b
   704: 4805c03d                 ldr q8, [x10, #0x10]
   708: a51c286e                 eor v5.16b, v5.16b, v8.16b
   712: 4809c03d                 ldr q8, [x10, #0x20]
   716: c61c286e                 eor v6.16b, v6.16b, v8.16b
   720: 480dc03d                 ldr q8, [x10, #0x30]
   724: e71c286e                 eor v7.16b, v7.16b, v8.16b
   728: 0d62c03d                 ldr q13, [x16, #0x180]
   732: 0e72c03d                 ldr q14, [x16, #0x1c0]
   736: 091c2f4e                 and v9.16b, v0.16b, v15.16b
   740: 0a040c6f                 ushr v10.16b, v0.16b, #4
   744: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   748: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   752: 681dab4e                 mov v8.16b, v11.16b
   756: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   760: 081d246e                 eor v8.16b, v8.16b, v4.16b
   764: 0d82c03d                 ldr q13, [x16, #0x200]
   768: 0e92c03d                 ldr q14, [x16, #0x240]
   772: 091c2f4e                 and v9.16b, v0.16b, v15.16b
   776: 0a040c6f                 ushr v10.16b, v0.16b, #4
   780: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   784: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   788: 6b1dab4e                 mov v11.16b, v11.16b
   792: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   796: 6b1d246e                 eor v11.16b, v11.16b, v4.16b
   800: 0da2c03d                 ldr q13, [x16, #0x280]
   804: 0eb2c03d                 ldr q14, [x16, #0x2c0]
   808: 691d2f4e                 and v9.16b, v11.16b, v15.16b
   812: 6a050c6f                 ushr v10.16b, v11.16b, #4
   816: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   820: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   824: 641dab4e                 mov v4.16b, v11.16b
   828: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   832: 0da2c03d                 ldr q13, [x16, #0x280]
   836: 0eb2c03d                 ldr q14, [x16, #0x2c0]
   840: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   844: 0a050c6f                 ushr v10.16b, v8.16b, #4
   848: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   852: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   856: 601dab4e                 mov v0.16b, v11.16b
   860: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
   864: 0d62c03d                 ldr q13, [x16, #0x180]
   868: 0e72c03d                 ldr q14, [x16, #0x1c0]
   872: 291c2f4e                 and v9.16b, v1.16b, v15.16b
   876: 2a040c6f                 ushr v10.16b, v1.16b, #4
   880: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   884: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   888: 681dab4e                 mov v8.16b, v11.16b
   892: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   896: 081d256e                 eor v8.16b, v8.16b, v5.16b
   900: 0d82c03d                 ldr q13, [x16, #0x200]
   904: 0e92c03d                 ldr q14, [x16, #0x240]
   908: 291c2f4e                 and v9.16b, v1.16b, v15.16b
   912: 2a040c6f                 ushr v10.16b, v1.16b, #4
   916: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   920: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   924: 6b1dab4e                 mov v11.16b, v11.16b
   928: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   932: 6b1d256e                 eor v11.16b, v11.16b, v5.16b
   936: 0da2c03d                 ldr q13, [x16, #0x280]
   940: 0eb2c03d                 ldr q14, [x16, #0x2c0]
   944: 691d2f4e                 and v9.16b, v11.16b, v15.16b
   948: 6a050c6f                 ushr v10.16b, v11.16b, #4
   952: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   956: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   960: 651dab4e                 mov v5.16b, v11.16b
   964: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   968: 0da2c03d                 ldr q13, [x16, #0x280]
   972: 0eb2c03d                 ldr q14, [x16, #0x2c0]
   976: 091d2f4e                 and v9.16b, v8.16b, v15.16b
   980: 0a050c6f                 ushr v10.16b, v8.16b, #4
   984: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
   988: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
   992: 611dab4e                 mov v1.16b, v11.16b
   996: 211c2c6e                 eor v1.16b, v1.16b, v12.16b
  1000: 0d62c03d                 ldr q13, [x16, #0x180]
  1004: 0e72c03d                 ldr q14, [x16, #0x1c0]
  1008: 491c2f4e                 and v9.16b, v2.16b, v15.16b
  1012: 4a040c6f                 ushr v10.16b, v2.16b, #4
  1016: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
  1020: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
  1024: 681dab4e                 mov v8.16b, v11.16b
  1028: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
  1032: 081d266e                 eor v8.16b, v8.16b, v6.16b
  1036: 0d82c03d                 ldr q13, [x16, #0x200]
  1040: 0e92c03d                 ldr q14, [x16, #0x240]
  1044: 491c2f4e                 and v9.16b, v2.16b, v15.16b
  1048: 4a040c6f                 ushr v10.16b, v2.16b, #4
  1052: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
  1056: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
  1060: 6b1dab4e                 mov v11.16b, v11.16b
  1064: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
  1068: 6b1d266e                 eor v11.16b, v11.16b, v6.16b
  1072: 0da2c03d                 ldr q13, [x16, #0x280]
  1076: 0eb2c03d                 ldr q14, [x16, #0x2c0]
  1080: 691d2f4e                 and v9.16b, v11.16b, v15.16b
  1084: 6a050c6f                 ushr v10.16b, v11.16b, #4
  1088: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
  1092: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
  1096: 661dab4e                 mov v6.16b, v11.16b
  1100: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
  1104: 0da2c03d                 ldr q13, [x16, #0x280]
  1108: 0eb2c03d                 ldr q14, [x16, #0x2c0]
  1112: 091d2f4e                 and v9.16b, v8.16b, v15.16b
  1116: 0a050c6f                 ushr v10.16b, v8.16b, #4
  1120: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
  1124: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
  1128: 621dab4e                 mov v2.16b, v11.16b
  1132: 421c2c6e                 eor v2.16b, v2.16b, v12.16b
  1136: 0d62c03d                 ldr q13, [x16, #0x180]
  1140: 0e72c03d                 ldr q14, [x16, #0x1c0]
  1144: 691c2f4e                 and v9.16b, v3.16b, v15.16b
  1148: 6a040c6f                 ushr v10.16b, v3.16b, #4
  1152: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
  1156: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
  1160: 681dab4e                 mov v8.16b, v11.16b
  1164: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
  1168: 081d276e                 eor v8.16b, v8.16b, v7.16b
  1172: 0d82c03d                 ldr q13, [x16, #0x200]
  1176: 0e92c03d                 ldr q14, [x16, #0x240]
  1180: 691c2f4e                 and v9.16b, v3.16b, v15.16b
  1184: 6a040c6f                 ushr v10.16b, v3.16b, #4
  1188: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
  1192: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
  1196: 6b1dab4e                 mov v11.16b, v11.16b
  1200: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
  1204: 6b1d276e                 eor v11.16b, v11.16b, v7.16b
  1208: 0da2c03d                 ldr q13, [x16, #0x280]
  1212: 0eb2c03d                 ldr q14, [x16, #0x2c0]
  1216: 691d2f4e                 and v9.16b, v11.16b, v15.16b
  1220: 6a050c6f                 ushr v10.16b, v11.16b, #4
  1224: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
  1228: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
  1232: 671dab4e                 mov v7.16b, v11.16b
  1236: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
  1240: 0da2c03d                 ldr q13, [x16, #0x280]
  1244: 0eb2c03d                 ldr q14, [x16, #0x2c0]
  1248: 091d2f4e                 and v9.16b, v8.16b, v15.16b
  1252: 0a050c6f                 ushr v10.16b, v8.16b, #4
  1256: ab01094e                 tbl v11.16b, { v13.16b }, v9.16b
  1260: cc010a4e                 tbl v12.16b, { v14.16b }, v10.16b
  1264: 631dab4e                 mov v3.16b, v11.16b
  1268: 631c2c6e                 eor v3.16b, v3.16b, v12.16b
  1272: 2a0140f9                 ldr x10, [x9]
  1276: 4a010b8b                 add x10, x10, x11
  1280: 4001803d                 str q0, [x10]
  1284: 4105803d                 str q1, [x10, #0x10]
  1288: 4209803d                 str q2, [x10, #0x20]
  1292: 430d803d                 str q3, [x10, #0x30]
  1296: 2a0540f9                 ldr x10, [x9, #8]
  1300: 4a010b8b                 add x10, x10, x11
  1304: 4401803d                 str q4, [x10]
  1308: 4505803d                 str q5, [x10, #0x10]
  1312: 4609803d                 str q6, [x10, #0x20]
  1316: 470d803d                 str q7, [x10, #0x30]
  1320: 6b010191                 add x11, x11, #0x40
  1324: ad0500f1                 subs x13, x13, #1
  1328: 01d8ff54                 b.ne 0x30
  1332: 7f0102eb                 cmp x11, x2
  1336: 22110054                 b.hs 0x75c
  1340: 0f020c91                 add x15, x16, #0x300
  1344: 11021091                 add x17, x16, #0x400
  1348: 2a0140f9                 ldr x10, [x9]
  1352: ed030baa                 mov x13, x11
  1356: 0c008052                 mov w12, #0
  1360: 0e0140f9                 ldr x14, [x8]
  1364: ce696d38                 ldrb w14, [x14, x13]
  1368: 0e010034                 cbz w14, 0x578
  1372: ee696e38                 ldrb w14, [x15, x14]
  1376: ce050311                 add w14, w14, #0xc1
  1380: dffd0371                 cmp w14, #0xff
  1384: 43000054                 b.lo 0x570
  1388: cefd0351                 sub w14, w14, #0xff
  1392: 2e6a6e38                 ldrb w14, [x17, x14]
  1396: 8c010e4a                 eor w12, w12, w14
  1400: 0e0540f9                 ldr x14, [x8, #8]
  1404: ce696d38                 ldrb w14, [x14, x13]
  1408: 0e010034                 cbz w14, 0x5a0
  1412: ee696e38                 ldrb w14, [x15, x14]
  1416: ce0d0011                 add w14, w14, #3
  1420: dffd0371                 cmp w14, #0xff
  1424: 43000054                 b.lo 0x598
  1428: cefd0351                 sub w14, w14, #0xff
  1432: 2e6a6e38                 ldrb w14, [x17, x14]
  1436: 8c010e4a                 eor w12, w12, w14
  1440: 0e0940f9                 ldr x14, [x8, #0x10]
  1444: ce696d38                 ldrb w14, [x14, x13]
  1448: 0e010034                 cbz w14, 0x5c8
  1452: ee696e38                 ldrb w14, [x15, x14]
  1456: ce190311                 add w14, w14, #0xc6
  1460: dffd0371                 cmp w14, #0xff
  1464: 43000054                 b.lo 0x5c0
  1468: cefd0351                 sub w14, w14, #0xff
  1472: 2e6a6e38                 ldrb w14, [x17, x14]
  1476: 8c010e4a                 eor w12, w12, w14
  1480: 0e0d40f9                 ldr x14, [x8, #0x18]
  1484: ce696d38                 ldrb w14, [x14, x13]
  1488: 0e010034                 cbz w14, 0x5f0
  1492: ee696e38                 ldrb w14, [x15, x14]
  1496: ce050011                 add w14, w14, #1
  1500: dffd0371                 cmp w14, #0xff
  1504: 43000054                 b.lo 0x5e8
  1508: cefd0351                 sub w14, w14, #0xff
  1512: 2e6a6e38                 ldrb w14, [x17, x14]
  1516: 8c010e4a                 eor w12, w12, w14
  1520: 0e1140f9                 ldr x14, [x8, #0x20]
  1524: ce696d38                 ldrb w14, [x14, x13]
  1528: 0e010034                 cbz w14, 0x618
  1532: ee696e38                 ldrb w14, [x15, x14]
  1536: cef90311                 add w14, w14, #0xfe
  1540: dffd0371                 cmp w14, #0xff
  1544: 43000054                 b.lo 0x610
  1548: cefd0351                 sub w14, w14, #0xff
  1552: 2e6a6e38                 ldrb w14, [x17, x14]
  1556: 8c010e4a                 eor w12, w12, w14
  1560: 0e1540f9                 ldr x14, [x8, #0x28]
  1564: ce696d38                 ldrb w14, [x14, x13]
  1568: 0e010034                 cbz w14, 0x640
  1572: ee696e38                 ldrb w14, [x15, x14]
  1576: cec50011                 add w14, w14, #0x31
  1580: dffd0371                 cmp w14, #0xff
  1584: 43000054                 b.lo 0x638
  1588: cefd0351                 sub w14, w14, #0xff
  1592: 2e6a6e38                 ldrb w14, [x17, x14]
  1596: 8c010e4a                 eor w12, w12, w14
  1600: 4c692d38                 strb w12, [x10, x13]
  1604: ad050091                 add x13, x13, #1
  1608: bf0102eb                 cmp x13, x2
  1612: 03f8ff54                 b.lo 0x54c
  1616: 2a0540f9                 ldr x10, [x9, #8]
  1620: ed030baa                 mov x13, x11
  1624: 0c008052                 mov w12, #0
  1628: 0e0140f9                 ldr x14, [x8]
  1632: ce696d38                 ldrb w14, [x14, x13]
  1636: 0e010034                 cbz w14, 0x684
  1640: ee696e38                 ldrb w14, [x15, x14]
  1644: ce710011                 add w14, w14, #0x1c
  1648: dffd0371                 cmp w14, #0xff
  1652: 43000054                 b.lo 0x67c
  1656: cefd0351                 sub w14, w14, #0xff
  1660: 2e6a6e38                 ldrb w14, [x17, x14]
  1664: 8c010e4a                 eor w12, w12, w14
  1668: 0e0540f9                 ldr x14, [x8, #8]
  1672: ce696d38                 ldrb w14, [x14, x13]
  1676: 0e010034                 cbz w14, 0x6ac
  1680: ee696e38                 ldrb w14, [x15, x14]
  1684: ce7d0311                 add w14, w14, #0xdf
  1688: dffd0371                 cmp w14, #0xff
  1692: 43000054                 b.lo 0x6a4
  1696: cefd0351                 sub w14, w14, #0xff
  1700: 2e6a6e38                 ldrb w14, [x17, x14]
  1704: 8c010e4a                 eor w12, w12, w14
  1708: 0e0940f9                 ldr x14, [x8, #0x10]
  1712: ce696d38                 ldrb w14, [x14, x13]
  1716: 0e010034                 cbz w14, 0x6d4
  1720: ee696e38                 ldrb w14, [x15, x14]
  1724: ce690011                 add w14, w14, #0x1a
  1728: dffd0371                 cmp w14, #0xff
  1732: 43000054                 b.lo 0x6cc
  1736: cefd0351                 sub w14, w14, #0xff
  1740: 2e6a6e38                 ldrb w14, [x17, x14]
  1744: 8c010e4a                 eor w12, w12, w14
  1748: 0e0d40f9                 ldr x14, [x8, #0x18]
  1752: ce696d38                 ldrb w14, [x14, x13]
  1756: 0e010034                 cbz w14, 0x6fc
  1760: ee696e38                 ldrb w14, [x15, x14]
  1764: ce650011                 add w14, w14, #0x19
  1768: dffd0371                 cmp w14, #0xff
  1772: 43000054                 b.lo 0x6f4
  1776: cefd0351                 sub w14, w14, #0xff
  1780: 2e6a6e38                 ldrb w14, [x17, x14]
  1784: 8c010e4a                 eor w12, w12, w14
  1788: 0e1140f9                 ldr x14, [x8, #0x20]
  1792: ce696d38                 ldrb w14, [x14, x13]
  1796: 0e010034                 cbz w14, 0x724
  1800: ee696e38                 ldrb w14, [x15, x14]
  1804: ce610011                 add w14, w14, #0x18
  1808: dffd0371                 cmp w14, #0xff
  1812: 43000054                 b.lo 0x71c
  1816: cefd0351                 sub w14, w14, #0xff
  1820: 2e6a6e38                 ldrb w14, [x17, x14]
  1824: 8c010e4a                 eor w12, w12, w14
  1828: 0e1540f9                 ldr x14, [x8, #0x28]
  1832: ce696d38                 ldrb w14, [x14, x13]
  1836: 0e010034                 cbz w14, 0x74c
  1840: ee696e38                 ldrb w14, [x15, x14]
  1844: cec50011                 add w14, w14, #0x31
  1848: dffd0371                 cmp w14, #0xff
  1852: 43000054                 b.lo 0x744
  1856: cefd0351                 sub w14, w14, #0xff
  1860: 2e6a6e38                 ldrb w14, [x17, x14]
  1864: 8c010e4a                 eor w12, w12, w14
  1868: 4c692d38                 strb w12, [x10, x13]
  1872: ad050091                 add x13, x13, #1
  1876: bf0102eb                 cmp x13, x2
  1880: 03f8ff54                 b.lo 0x658
  1884: ee3fc1ac                 ldp q14, q15, [sp], #0x20
  1888: ec37c1ac                 ldp q12, q13, [sp], #0x20
  1892: ea2fc1ac                 ldp q10, q11, [sp], #0x20
  1896: e827c1ac                 ldp q8, q9, [sp], #0x20
  1900: c0035fd6                 ret 
  1904: 00000000                 udf #0
  1908: 00000000                 udf #0
  1912: 00000000                 udf #0
  1916: 00000000                 udf #0
