; config_t1_1204 target=neon twin=xorsched topo=rdp(7)/enc op=encode U=1
; code_size=2816 twins=["xorsched"] uops_per_byte=9.6875
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: ec37bfad                 stp q12, q13, [sp, #-0x20]!
    12: e80300aa                 mov x8, x0
    16: e90301aa                 mov x9, x1
    20: eb031faa                 mov x11, xzr
    24: 0e0280d2                 mov x14, #0x10
    28: 4d08ce9a                 udiv x13, x2, x14
    32: 0d2700b4                 cbz x13, 0x500
    36: 1f2003d5                 nop 
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0540f9                 ldr x10, [x8, #8]
    68: 4a010b8b                 add x10, x10, x11
    72: 031d40f9                 ldr x3, [x8, #0x38]
    76: 63000b8b                 add x3, x3, x11
    80: 4c01c03d                 ldr q12, [x10]
    84: 6d00c03d                 ldr q13, [x3]
    88: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
    92: 811dac4e                 mov v1.16b, v12.16b
    96: 861dac4e                 mov v6.16b, v12.16b
   100: 0a0940f9                 ldr x10, [x8, #0x10]
   104: 4a010b8b                 add x10, x10, x11
   108: 032140f9                 ldr x3, [x8, #0x40]
   112: 63000b8b                 add x3, x3, x11
   116: 4c01c03d                 ldr q12, [x10]
   120: 6d00c03d                 ldr q13, [x3]
   124: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   128: 821dac4e                 mov v2.16b, v12.16b
   132: 871dac4e                 mov v7.16b, v12.16b
   136: 0a0d40f9                 ldr x10, [x8, #0x18]
   140: 4a010b8b                 add x10, x10, x11
   144: 032540f9                 ldr x3, [x8, #0x48]
   148: 63000b8b                 add x3, x3, x11
   152: 4c01c03d                 ldr q12, [x10]
   156: 6d00c03d                 ldr q13, [x3]
   160: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   164: 831dac4e                 mov v3.16b, v12.16b
   168: 881dac4e                 mov v8.16b, v12.16b
   172: 0a1140f9                 ldr x10, [x8, #0x20]
   176: 4a010b8b                 add x10, x10, x11
   180: 032940f9                 ldr x3, [x8, #0x50]
   184: 63000b8b                 add x3, x3, x11
   188: 4c01c03d                 ldr q12, [x10]
   192: 6d00c03d                 ldr q13, [x3]
   196: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   200: 841dac4e                 mov v4.16b, v12.16b
   204: 891dac4e                 mov v9.16b, v12.16b
   208: 0a1540f9                 ldr x10, [x8, #0x28]
   212: 4a010b8b                 add x10, x10, x11
   216: 032d40f9                 ldr x3, [x8, #0x58]
   220: 63000b8b                 add x3, x3, x11
   224: 4c01c03d                 ldr q12, [x10]
   228: 6d00c03d                 ldr q13, [x3]
   232: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   236: 851dac4e                 mov v5.16b, v12.16b
   240: 8a1dac4e                 mov v10.16b, v12.16b
   244: 0a3540f9                 ldr x10, [x8, #0x68]
   248: 4a010b8b                 add x10, x10, x11
   252: 034d40f9                 ldr x3, [x8, #0x98]
   256: 63000b8b                 add x3, x3, x11
   260: 4c01c03d                 ldr q12, [x10]
   264: 6d00c03d                 ldr q13, [x3]
   268: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   272: 211c2c6e                 eor v1.16b, v1.16b, v12.16b
   276: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   280: 0a3940f9                 ldr x10, [x8, #0x70]
   284: 4a010b8b                 add x10, x10, x11
   288: 035140f9                 ldr x3, [x8, #0xa0]
   292: 63000b8b                 add x3, x3, x11
   296: 4c01c03d                 ldr q12, [x10]
   300: 6d00c03d                 ldr q13, [x3]
   304: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   308: 421c2c6e                 eor v2.16b, v2.16b, v12.16b
   312: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   316: 0a3d40f9                 ldr x10, [x8, #0x78]
   320: 4a010b8b                 add x10, x10, x11
   324: 035540f9                 ldr x3, [x8, #0xa8]
   328: 63000b8b                 add x3, x3, x11
   332: 4c01c03d                 ldr q12, [x10]
   336: 6d00c03d                 ldr q13, [x3]
   340: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   344: 631c2c6e                 eor v3.16b, v3.16b, v12.16b
   348: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   352: 0a4140f9                 ldr x10, [x8, #0x80]
   356: 4a010b8b                 add x10, x10, x11
   360: 035940f9                 ldr x3, [x8, #0xb0]
   364: 63000b8b                 add x3, x3, x11
   368: 4c01c03d                 ldr q12, [x10]
   372: 6d00c03d                 ldr q13, [x3]
   376: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   380: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   384: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   388: 0a4540f9                 ldr x10, [x8, #0x88]
   392: 4a010b8b                 add x10, x10, x11
   396: 035d40f9                 ldr x3, [x8, #0xb8]
   400: 63000b8b                 add x3, x3, x11
   404: 4c01c03d                 ldr q12, [x10]
   408: 6d00c03d                 ldr q13, [x3]
   412: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   416: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   420: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   424: 0a6540f9                 ldr x10, [x8, #0xc8]
   428: 4a010b8b                 add x10, x10, x11
   432: 037d40f9                 ldr x3, [x8, #0xf8]
   436: 63000b8b                 add x3, x3, x11
   440: 4c01c03d                 ldr q12, [x10]
   444: 6d00c03d                 ldr q13, [x3]
   448: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   452: 211c2c6e                 eor v1.16b, v1.16b, v12.16b
   456: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   460: 0a6940f9                 ldr x10, [x8, #0xd0]
   464: 4a010b8b                 add x10, x10, x11
   468: 038140f9                 ldr x3, [x8, #0x100]
   472: 63000b8b                 add x3, x3, x11
   476: 4c01c03d                 ldr q12, [x10]
   480: 6d00c03d                 ldr q13, [x3]
   484: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   488: 421c2c6e                 eor v2.16b, v2.16b, v12.16b
   492: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   496: 0a6d40f9                 ldr x10, [x8, #0xd8]
   500: 4a010b8b                 add x10, x10, x11
   504: 038540f9                 ldr x3, [x8, #0x108]
   508: 63000b8b                 add x3, x3, x11
   512: 4c01c03d                 ldr q12, [x10]
   516: 6d00c03d                 ldr q13, [x3]
   520: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   524: 631c2c6e                 eor v3.16b, v3.16b, v12.16b
   528: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   532: 0a7140f9                 ldr x10, [x8, #0xe0]
   536: 4a010b8b                 add x10, x10, x11
   540: 038940f9                 ldr x3, [x8, #0x110]
   544: 63000b8b                 add x3, x3, x11
   548: 4c01c03d                 ldr q12, [x10]
   552: 6d00c03d                 ldr q13, [x3]
   556: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   560: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   564: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   568: 0a7540f9                 ldr x10, [x8, #0xe8]
   572: 4a010b8b                 add x10, x10, x11
   576: 038d40f9                 ldr x3, [x8, #0x118]
   580: 63000b8b                 add x3, x3, x11
   584: 4c01c03d                 ldr q12, [x10]
   588: 6d00c03d                 ldr q13, [x3]
   592: 8c1d2d6e                 eor v12.16b, v12.16b, v13.16b
   596: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   600: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   604: 0a0140f9                 ldr x10, [x8]
   608: 4a010b8b                 add x10, x10, x11
   612: 4c01c03d                 ldr q12, [x10]
   616: 801dac4e                 mov v0.16b, v12.16b
   620: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   624: 0a0540f9                 ldr x10, [x8, #8]
   628: 4a010b8b                 add x10, x10, x11
   632: 4c01c03d                 ldr q12, [x10]
   636: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   640: 0a0940f9                 ldr x10, [x8, #0x10]
   644: 4a010b8b                 add x10, x10, x11
   648: 4c01c03d                 ldr q12, [x10]
   652: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   656: 0a0d40f9                 ldr x10, [x8, #0x18]
   660: 4a010b8b                 add x10, x10, x11
   664: 4c01c03d                 ldr q12, [x10]
   668: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   672: 0a1140f9                 ldr x10, [x8, #0x20]
   676: 4a010b8b                 add x10, x10, x11
   680: 4c01c03d                 ldr q12, [x10]
   684: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   688: 0a1540f9                 ldr x10, [x8, #0x28]
   692: 4a010b8b                 add x10, x10, x11
   696: 4c01c03d                 ldr q12, [x10]
   700: 8b1dac4e                 mov v11.16b, v12.16b
   704: 0a1940f9                 ldr x10, [x8, #0x30]
   708: 4a010b8b                 add x10, x10, x11
   712: 4c01c03d                 ldr q12, [x10]
   716: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
   720: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   724: 0a1d40f9                 ldr x10, [x8, #0x38]
   728: 4a010b8b                 add x10, x10, x11
   732: 4c01c03d                 ldr q12, [x10]
   736: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   740: 0a2140f9                 ldr x10, [x8, #0x40]
   744: 4a010b8b                 add x10, x10, x11
   748: 4c01c03d                 ldr q12, [x10]
   752: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   756: 0a2540f9                 ldr x10, [x8, #0x48]
   760: 4a010b8b                 add x10, x10, x11
   764: 4c01c03d                 ldr q12, [x10]
   768: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   772: 0a2940f9                 ldr x10, [x8, #0x50]
   776: 4a010b8b                 add x10, x10, x11
   780: 4c01c03d                 ldr q12, [x10]
   784: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   788: 0a3140f9                 ldr x10, [x8, #0x60]
   792: 4a010b8b                 add x10, x10, x11
   796: 4c01c03d                 ldr q12, [x10]
   800: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
   804: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   808: 0a3540f9                 ldr x10, [x8, #0x68]
   812: 4a010b8b                 add x10, x10, x11
   816: 4c01c03d                 ldr q12, [x10]
   820: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   824: 0a3940f9                 ldr x10, [x8, #0x70]
   828: 4a010b8b                 add x10, x10, x11
   832: 4c01c03d                 ldr q12, [x10]
   836: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   840: 0a3d40f9                 ldr x10, [x8, #0x78]
   844: 4a010b8b                 add x10, x10, x11
   848: 4c01c03d                 ldr q12, [x10]
   852: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   856: 0a4540f9                 ldr x10, [x8, #0x88]
   860: 4a010b8b                 add x10, x10, x11
   864: 4c01c03d                 ldr q12, [x10]
   868: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   872: 0a4940f9                 ldr x10, [x8, #0x90]
   876: 4a010b8b                 add x10, x10, x11
   880: 4c01c03d                 ldr q12, [x10]
   884: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
   888: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   892: 0a4d40f9                 ldr x10, [x8, #0x98]
   896: 4a010b8b                 add x10, x10, x11
   900: 4c01c03d                 ldr q12, [x10]
   904: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   908: 0a5140f9                 ldr x10, [x8, #0xa0]
   912: 4a010b8b                 add x10, x10, x11
   916: 4c01c03d                 ldr q12, [x10]
   920: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   924: 0a5940f9                 ldr x10, [x8, #0xb0]
   928: 4a010b8b                 add x10, x10, x11
   932: 4c01c03d                 ldr q12, [x10]
   936: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   940: 0a5d40f9                 ldr x10, [x8, #0xb8]
   944: 4a010b8b                 add x10, x10, x11
   948: 4c01c03d                 ldr q12, [x10]
   952: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   956: 0a6140f9                 ldr x10, [x8, #0xc0]
   960: 4a010b8b                 add x10, x10, x11
   964: 4c01c03d                 ldr q12, [x10]
   968: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
   972: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   976: 0a6540f9                 ldr x10, [x8, #0xc8]
   980: 4a010b8b                 add x10, x10, x11
   984: 4c01c03d                 ldr q12, [x10]
   988: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   992: 0a6d40f9                 ldr x10, [x8, #0xd8]
   996: 4a010b8b                 add x10, x10, x11
  1000: 4c01c03d                 ldr q12, [x10]
  1004: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
  1008: 0a7140f9                 ldr x10, [x8, #0xe0]
  1012: 4a010b8b                 add x10, x10, x11
  1016: 4c01c03d                 ldr q12, [x10]
  1020: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
  1024: 0a7540f9                 ldr x10, [x8, #0xe8]
  1028: 4a010b8b                 add x10, x10, x11
  1032: 4c01c03d                 ldr q12, [x10]
  1036: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
  1040: 0a7940f9                 ldr x10, [x8, #0xf0]
  1044: 4a010b8b                 add x10, x10, x11
  1048: 4c01c03d                 ldr q12, [x10]
  1052: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
  1056: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
  1060: 0a8140f9                 ldr x10, [x8, #0x100]
  1064: 4a010b8b                 add x10, x10, x11
  1068: 4c01c03d                 ldr q12, [x10]
  1072: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
  1076: 0a8540f9                 ldr x10, [x8, #0x108]
  1080: 4a010b8b                 add x10, x10, x11
  1084: 4c01c03d                 ldr q12, [x10]
  1088: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
  1092: 0a8940f9                 ldr x10, [x8, #0x110]
  1096: 4a010b8b                 add x10, x10, x11
  1100: 4c01c03d                 ldr q12, [x10]
  1104: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
  1108: 0a8d40f9                 ldr x10, [x8, #0x118]
  1112: 4a010b8b                 add x10, x10, x11
  1116: 4c01c03d                 ldr q12, [x10]
  1120: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
  1124: 2a0140f9                 ldr x10, [x9]
  1128: 4a010b8b                 add x10, x10, x11
  1132: 4001803d                 str q0, [x10]
  1136: 2a0540f9                 ldr x10, [x9, #8]
  1140: 4a010b8b                 add x10, x10, x11
  1144: 4101803d                 str q1, [x10]
  1148: 2a0940f9                 ldr x10, [x9, #0x10]
  1152: 4a010b8b                 add x10, x10, x11
  1156: 4201803d                 str q2, [x10]
  1160: 2a0d40f9                 ldr x10, [x9, #0x18]
  1164: 4a010b8b                 add x10, x10, x11
  1168: 4301803d                 str q3, [x10]
  1172: 2a1140f9                 ldr x10, [x9, #0x20]
  1176: 4a010b8b                 add x10, x10, x11
  1180: 4401803d                 str q4, [x10]
  1184: 2a1540f9                 ldr x10, [x9, #0x28]
  1188: 4a010b8b                 add x10, x10, x11
  1192: 4501803d                 str q5, [x10]
  1196: 2a1940f9                 ldr x10, [x9, #0x30]
  1200: 4a010b8b                 add x10, x10, x11
  1204: 4601803d                 str q6, [x10]
  1208: 2a1d40f9                 ldr x10, [x9, #0x38]
  1212: 4a010b8b                 add x10, x10, x11
  1216: 4701803d                 str q7, [x10]
  1220: 2a2140f9                 ldr x10, [x9, #0x40]
  1224: 4a010b8b                 add x10, x10, x11
  1228: 4801803d                 str q8, [x10]
  1232: 2a2540f9                 ldr x10, [x9, #0x48]
  1236: 4a010b8b                 add x10, x10, x11
  1240: 4901803d                 str q9, [x10]
  1244: 2a2940f9                 ldr x10, [x9, #0x50]
  1248: 4a010b8b                 add x10, x10, x11
  1252: 4a01803d                 str q10, [x10]
  1256: 2a2d40f9                 ldr x10, [x9, #0x58]
  1260: 4a010b8b                 add x10, x10, x11
  1264: 4b01803d                 str q11, [x10]
  1268: 6b410091                 add x11, x11, #0x10
  1272: ad0500f1                 subs x13, x13, #1
  1276: 41d9ff54                 b.ne 0x24
  1280: 7f0102eb                 cmp x11, x2
  1284: 822d0054                 b.hs 0xab4
  1288: 2a0140f9                 ldr x10, [x9]
  1292: ed030baa                 mov x13, x11
  1296: 0e0140f9                 ldr x14, [x8]
  1300: ce696d38                 ldrb w14, [x14, x13]
  1304: ec030e2a                 mov w12, w14
  1308: 0e1940f9                 ldr x14, [x8, #0x30]
  1312: ce696d38                 ldrb w14, [x14, x13]
  1316: 8c010e4a                 eor w12, w12, w14
  1320: 0e3140f9                 ldr x14, [x8, #0x60]
  1324: ce696d38                 ldrb w14, [x14, x13]
  1328: 8c010e4a                 eor w12, w12, w14
  1332: 0e4940f9                 ldr x14, [x8, #0x90]
  1336: ce696d38                 ldrb w14, [x14, x13]
  1340: 8c010e4a                 eor w12, w12, w14
  1344: 0e6140f9                 ldr x14, [x8, #0xc0]
  1348: ce696d38                 ldrb w14, [x14, x13]
  1352: 8c010e4a                 eor w12, w12, w14
  1356: 0e7940f9                 ldr x14, [x8, #0xf0]
  1360: ce696d38                 ldrb w14, [x14, x13]
  1364: 8c010e4a                 eor w12, w12, w14
  1368: 4c692d38                 strb w12, [x10, x13]
  1372: ad050091                 add x13, x13, #1
  1376: bf0102eb                 cmp x13, x2
  1380: 63fdff54                 b.lo 0x510
  1384: 2a0540f9                 ldr x10, [x9, #8]
  1388: ed030baa                 mov x13, x11
  1392: 0e0540f9                 ldr x14, [x8, #8]
  1396: ce696d38                 ldrb w14, [x14, x13]
  1400: ec030e2a                 mov w12, w14
  1404: 0e1d40f9                 ldr x14, [x8, #0x38]
  1408: ce696d38                 ldrb w14, [x14, x13]
  1412: 8c010e4a                 eor w12, w12, w14
  1416: 0e3540f9                 ldr x14, [x8, #0x68]
  1420: ce696d38                 ldrb w14, [x14, x13]
  1424: 8c010e4a                 eor w12, w12, w14
  1428: 0e4d40f9                 ldr x14, [x8, #0x98]
  1432: ce696d38                 ldrb w14, [x14, x13]
  1436: 8c010e4a                 eor w12, w12, w14
  1440: 0e6540f9                 ldr x14, [x8, #0xc8]
  1444: ce696d38                 ldrb w14, [x14, x13]
  1448: 8c010e4a                 eor w12, w12, w14
  1452: 0e7d40f9                 ldr x14, [x8, #0xf8]
  1456: ce696d38                 ldrb w14, [x14, x13]
  1460: 8c010e4a                 eor w12, w12, w14
  1464: 4c692d38                 strb w12, [x10, x13]
  1468: ad050091                 add x13, x13, #1
  1472: bf0102eb                 cmp x13, x2
  1476: 63fdff54                 b.lo 0x570
  1480: 2a0940f9                 ldr x10, [x9, #0x10]
  1484: ed030baa                 mov x13, x11
  1488: 0e0940f9                 ldr x14, [x8, #0x10]
  1492: ce696d38                 ldrb w14, [x14, x13]
  1496: ec030e2a                 mov w12, w14
  1500: 0e2140f9                 ldr x14, [x8, #0x40]
  1504: ce696d38                 ldrb w14, [x14, x13]
  1508: 8c010e4a                 eor w12, w12, w14
  1512: 0e3940f9                 ldr x14, [x8, #0x70]
  1516: ce696d38                 ldrb w14, [x14, x13]
  1520: 8c010e4a                 eor w12, w12, w14
  1524: 0e5140f9                 ldr x14, [x8, #0xa0]
  1528: ce696d38                 ldrb w14, [x14, x13]
  1532: 8c010e4a                 eor w12, w12, w14
  1536: 0e6940f9                 ldr x14, [x8, #0xd0]
  1540: ce696d38                 ldrb w14, [x14, x13]
  1544: 8c010e4a                 eor w12, w12, w14
  1548: 0e8140f9                 ldr x14, [x8, #0x100]
  1552: ce696d38                 ldrb w14, [x14, x13]
  1556: 8c010e4a                 eor w12, w12, w14
  1560: 4c692d38                 strb w12, [x10, x13]
  1564: ad050091                 add x13, x13, #1
  1568: bf0102eb                 cmp x13, x2
  1572: 63fdff54                 b.lo 0x5d0
  1576: 2a0d40f9                 ldr x10, [x9, #0x18]
  1580: ed030baa                 mov x13, x11
  1584: 0e0d40f9                 ldr x14, [x8, #0x18]
  1588: ce696d38                 ldrb w14, [x14, x13]
  1592: ec030e2a                 mov w12, w14
  1596: 0e2540f9                 ldr x14, [x8, #0x48]
  1600: ce696d38                 ldrb w14, [x14, x13]
  1604: 8c010e4a                 eor w12, w12, w14
  1608: 0e3d40f9                 ldr x14, [x8, #0x78]
  1612: ce696d38                 ldrb w14, [x14, x13]
  1616: 8c010e4a                 eor w12, w12, w14
  1620: 0e5540f9                 ldr x14, [x8, #0xa8]
  1624: ce696d38                 ldrb w14, [x14, x13]
  1628: 8c010e4a                 eor w12, w12, w14
  1632: 0e6d40f9                 ldr x14, [x8, #0xd8]
  1636: ce696d38                 ldrb w14, [x14, x13]
  1640: 8c010e4a                 eor w12, w12, w14
  1644: 0e8540f9                 ldr x14, [x8, #0x108]
  1648: ce696d38                 ldrb w14, [x14, x13]
  1652: 8c010e4a                 eor w12, w12, w14
  1656: 4c692d38                 strb w12, [x10, x13]
  1660: ad050091                 add x13, x13, #1
  1664: bf0102eb                 cmp x13, x2
  1668: 63fdff54                 b.lo 0x630
  1672: 2a1140f9                 ldr x10, [x9, #0x20]
  1676: ed030baa                 mov x13, x11
  1680: 0e1140f9                 ldr x14, [x8, #0x20]
  1684: ce696d38                 ldrb w14, [x14, x13]
  1688: ec030e2a                 mov w12, w14
  1692: 0e2940f9                 ldr x14, [x8, #0x50]
  1696: ce696d38                 ldrb w14, [x14, x13]
  1700: 8c010e4a                 eor w12, w12, w14
  1704: 0e4140f9                 ldr x14, [x8, #0x80]
  1708: ce696d38                 ldrb w14, [x14, x13]
  1712: 8c010e4a                 eor w12, w12, w14
  1716: 0e5940f9                 ldr x14, [x8, #0xb0]
  1720: ce696d38                 ldrb w14, [x14, x13]
  1724: 8c010e4a                 eor w12, w12, w14
  1728: 0e7140f9                 ldr x14, [x8, #0xe0]
  1732: ce696d38                 ldrb w14, [x14, x13]
  1736: 8c010e4a                 eor w12, w12, w14
  1740: 0e8940f9                 ldr x14, [x8, #0x110]
  1744: ce696d38                 ldrb w14, [x14, x13]
  1748: 8c010e4a                 eor w12, w12, w14
  1752: 4c692d38                 strb w12, [x10, x13]
  1756: ad050091                 add x13, x13, #1
  1760: bf0102eb                 cmp x13, x2
  1764: 63fdff54                 b.lo 0x690
  1768: 2a1540f9                 ldr x10, [x9, #0x28]
  1772: ed030baa                 mov x13, x11
  1776: 0e1540f9                 ldr x14, [x8, #0x28]
  1780: ce696d38                 ldrb w14, [x14, x13]
  1784: ec030e2a                 mov w12, w14
  1788: 0e2d40f9                 ldr x14, [x8, #0x58]
  1792: ce696d38                 ldrb w14, [x14, x13]
  1796: 8c010e4a                 eor w12, w12, w14
  1800: 0e4540f9                 ldr x14, [x8, #0x88]
  1804: ce696d38                 ldrb w14, [x14, x13]
  1808: 8c010e4a                 eor w12, w12, w14
  1812: 0e5d40f9                 ldr x14, [x8, #0xb8]
  1816: ce696d38                 ldrb w14, [x14, x13]
  1820: 8c010e4a                 eor w12, w12, w14
  1824: 0e7540f9                 ldr x14, [x8, #0xe8]
  1828: ce696d38                 ldrb w14, [x14, x13]
  1832: 8c010e4a                 eor w12, w12, w14
  1836: 0e8d40f9                 ldr x14, [x8, #0x118]
  1840: ce696d38                 ldrb w14, [x14, x13]
  1844: 8c010e4a                 eor w12, w12, w14
  1848: 4c692d38                 strb w12, [x10, x13]
  1852: ad050091                 add x13, x13, #1
  1856: bf0102eb                 cmp x13, x2
  1860: 63fdff54                 b.lo 0x6f0
  1864: 2a1940f9                 ldr x10, [x9, #0x30]
  1868: ed030baa                 mov x13, x11
  1872: 0e0140f9                 ldr x14, [x8]
  1876: ce696d38                 ldrb w14, [x14, x13]
  1880: ec030e2a                 mov w12, w14
  1884: 0e0540f9                 ldr x14, [x8, #8]
  1888: ce696d38                 ldrb w14, [x14, x13]
  1892: 8c010e4a                 eor w12, w12, w14
  1896: 0e1d40f9                 ldr x14, [x8, #0x38]
  1900: ce696d38                 ldrb w14, [x14, x13]
  1904: 8c010e4a                 eor w12, w12, w14
  1908: 0e3540f9                 ldr x14, [x8, #0x68]
  1912: ce696d38                 ldrb w14, [x14, x13]
  1916: 8c010e4a                 eor w12, w12, w14
  1920: 0e4540f9                 ldr x14, [x8, #0x88]
  1924: ce696d38                 ldrb w14, [x14, x13]
  1928: 8c010e4a                 eor w12, w12, w14
  1932: 0e4d40f9                 ldr x14, [x8, #0x98]
  1936: ce696d38                 ldrb w14, [x14, x13]
  1940: 8c010e4a                 eor w12, w12, w14
  1944: 0e5940f9                 ldr x14, [x8, #0xb0]
  1948: ce696d38                 ldrb w14, [x14, x13]
  1952: 8c010e4a                 eor w12, w12, w14
  1956: 0e6540f9                 ldr x14, [x8, #0xc8]
  1960: ce696d38                 ldrb w14, [x14, x13]
  1964: 8c010e4a                 eor w12, w12, w14
  1968: 0e6d40f9                 ldr x14, [x8, #0xd8]
  1972: ce696d38                 ldrb w14, [x14, x13]
  1976: 8c010e4a                 eor w12, w12, w14
  1980: 0e7d40f9                 ldr x14, [x8, #0xf8]
  1984: ce696d38                 ldrb w14, [x14, x13]
  1988: 8c010e4a                 eor w12, w12, w14
  1992: 0e8140f9                 ldr x14, [x8, #0x100]
  1996: ce696d38                 ldrb w14, [x14, x13]
  2000: 8c010e4a                 eor w12, w12, w14
  2004: 4c692d38                 strb w12, [x10, x13]
  2008: ad050091                 add x13, x13, #1
  2012: bf0102eb                 cmp x13, x2
  2016: 83fbff54                 b.lo 0x750
  2020: 2a1d40f9                 ldr x10, [x9, #0x38]
  2024: ed030baa                 mov x13, x11
  2028: 0e0540f9                 ldr x14, [x8, #8]
  2032: ce696d38                 ldrb w14, [x14, x13]
  2036: ec030e2a                 mov w12, w14
  2040: 0e0940f9                 ldr x14, [x8, #0x10]
  2044: ce696d38                 ldrb w14, [x14, x13]
  2048: 8c010e4a                 eor w12, w12, w14
  2052: 0e1940f9                 ldr x14, [x8, #0x30]
  2056: ce696d38                 ldrb w14, [x14, x13]
  2060: 8c010e4a                 eor w12, w12, w14
  2064: 0e2140f9                 ldr x14, [x8, #0x40]
  2068: ce696d38                 ldrb w14, [x14, x13]
  2072: 8c010e4a                 eor w12, w12, w14
  2076: 0e3940f9                 ldr x14, [x8, #0x70]
  2080: ce696d38                 ldrb w14, [x14, x13]
  2084: 8c010e4a                 eor w12, w12, w14
  2088: 0e5140f9                 ldr x14, [x8, #0xa0]
  2092: ce696d38                 ldrb w14, [x14, x13]
  2096: 8c010e4a                 eor w12, w12, w14
  2100: 0e5d40f9                 ldr x14, [x8, #0xb8]
  2104: ce696d38                 ldrb w14, [x14, x13]
  2108: 8c010e4a                 eor w12, w12, w14
  2112: 0e6940f9                 ldr x14, [x8, #0xd0]
  2116: ce696d38                 ldrb w14, [x14, x13]
  2120: 8c010e4a                 eor w12, w12, w14
  2124: 0e7140f9                 ldr x14, [x8, #0xe0]
  2128: ce696d38                 ldrb w14, [x14, x13]
  2132: 8c010e4a                 eor w12, w12, w14
  2136: 0e8140f9                 ldr x14, [x8, #0x100]
  2140: ce696d38                 ldrb w14, [x14, x13]
  2144: 8c010e4a                 eor w12, w12, w14
  2148: 0e8540f9                 ldr x14, [x8, #0x108]
  2152: ce696d38                 ldrb w14, [x14, x13]
  2156: 8c010e4a                 eor w12, w12, w14
  2160: 4c692d38                 strb w12, [x10, x13]
  2164: ad050091                 add x13, x13, #1
  2168: bf0102eb                 cmp x13, x2
  2172: 83fbff54                 b.lo 0x7ec
  2176: 2a2140f9                 ldr x10, [x9, #0x40]
  2180: ed030baa                 mov x13, x11
  2184: 0e0940f9                 ldr x14, [x8, #0x10]
  2188: ce696d38                 ldrb w14, [x14, x13]
  2192: ec030e2a                 mov w12, w14
  2196: 0e0d40f9                 ldr x14, [x8, #0x18]
  2200: ce696d38                 ldrb w14, [x14, x13]
  2204: 8c010e4a                 eor w12, w12, w14
  2208: 0e1d40f9                 ldr x14, [x8, #0x38]
  2212: ce696d38                 ldrb w14, [x14, x13]
  2216: 8c010e4a                 eor w12, w12, w14
  2220: 0e2540f9                 ldr x14, [x8, #0x48]
  2224: ce696d38                 ldrb w14, [x14, x13]
  2228: 8c010e4a                 eor w12, w12, w14
  2232: 0e3140f9                 ldr x14, [x8, #0x60]
  2236: ce696d38                 ldrb w14, [x14, x13]
  2240: 8c010e4a                 eor w12, w12, w14
  2244: 0e3d40f9                 ldr x14, [x8, #0x78]
  2248: ce696d38                 ldrb w14, [x14, x13]
  2252: 8c010e4a                 eor w12, w12, w14
  2256: 0e5540f9                 ldr x14, [x8, #0xa8]
  2260: ce696d38                 ldrb w14, [x14, x13]
  2264: 8c010e4a                 eor w12, w12, w14
  2268: 0e6d40f9                 ldr x14, [x8, #0xd8]
  2272: ce696d38                 ldrb w14, [x14, x13]
  2276: 8c010e4a                 eor w12, w12, w14
  2280: 0e7540f9                 ldr x14, [x8, #0xe8]
  2284: ce696d38                 ldrb w14, [x14, x13]
  2288: 8c010e4a                 eor w12, w12, w14
  2292: 0e8540f9                 ldr x14, [x8, #0x108]
  2296: ce696d38                 ldrb w14, [x14, x13]
  2300: 8c010e4a                 eor w12, w12, w14
  2304: 0e8940f9                 ldr x14, [x8, #0x110]
  2308: ce696d38                 ldrb w14, [x14, x13]
  2312: 8c010e4a                 eor w12, w12, w14
  2316: 4c692d38                 strb w12, [x10, x13]
  2320: ad050091                 add x13, x13, #1
  2324: bf0102eb                 cmp x13, x2
  2328: 83fbff54                 b.lo 0x888
  2332: 2a2540f9                 ldr x10, [x9, #0x48]
  2336: ed030baa                 mov x13, x11
  2340: 0e0d40f9                 ldr x14, [x8, #0x18]
  2344: ce696d38                 ldrb w14, [x14, x13]
  2348: ec030e2a                 mov w12, w14
  2352: 0e1140f9                 ldr x14, [x8, #0x20]
  2356: ce696d38                 ldrb w14, [x14, x13]
  2360: 8c010e4a                 eor w12, w12, w14
  2364: 0e2140f9                 ldr x14, [x8, #0x40]
  2368: ce696d38                 ldrb w14, [x14, x13]
  2372: 8c010e4a                 eor w12, w12, w14
  2376: 0e2940f9                 ldr x14, [x8, #0x50]
  2380: ce696d38                 ldrb w14, [x14, x13]
  2384: 8c010e4a                 eor w12, w12, w14
  2388: 0e3540f9                 ldr x14, [x8, #0x68]
  2392: ce696d38                 ldrb w14, [x14, x13]
  2396: 8c010e4a                 eor w12, w12, w14
  2400: 0e4140f9                 ldr x14, [x8, #0x80]
  2404: ce696d38                 ldrb w14, [x14, x13]
  2408: 8c010e4a                 eor w12, w12, w14
  2412: 0e4940f9                 ldr x14, [x8, #0x90]
  2416: ce696d38                 ldrb w14, [x14, x13]
  2420: 8c010e4a                 eor w12, w12, w14
  2424: 0e5940f9                 ldr x14, [x8, #0xb0]
  2428: ce696d38                 ldrb w14, [x14, x13]
  2432: 8c010e4a                 eor w12, w12, w14
  2436: 0e7140f9                 ldr x14, [x8, #0xe0]
  2440: ce696d38                 ldrb w14, [x14, x13]
  2444: 8c010e4a                 eor w12, w12, w14
  2448: 0e8940f9                 ldr x14, [x8, #0x110]
  2452: ce696d38                 ldrb w14, [x14, x13]
  2456: 8c010e4a                 eor w12, w12, w14
  2460: 0e8d40f9                 ldr x14, [x8, #0x118]
  2464: ce696d38                 ldrb w14, [x14, x13]
  2468: 8c010e4a                 eor w12, w12, w14
  2472: 4c692d38                 strb w12, [x10, x13]
  2476: ad050091                 add x13, x13, #1
  2480: bf0102eb                 cmp x13, x2
  2484: 83fbff54                 b.lo 0x924
  2488: 2a2940f9                 ldr x10, [x9, #0x50]
  2492: ed030baa                 mov x13, x11
  2496: 0e1140f9                 ldr x14, [x8, #0x20]
  2500: ce696d38                 ldrb w14, [x14, x13]
  2504: ec030e2a                 mov w12, w14
  2508: 0e1540f9                 ldr x14, [x8, #0x28]
  2512: ce696d38                 ldrb w14, [x14, x13]
  2516: 8c010e4a                 eor w12, w12, w14
  2520: 0e2540f9                 ldr x14, [x8, #0x48]
  2524: ce696d38                 ldrb w14, [x14, x13]
  2528: 8c010e4a                 eor w12, w12, w14
  2532: 0e2d40f9                 ldr x14, [x8, #0x58]
  2536: ce696d38                 ldrb w14, [x14, x13]
  2540: 8c010e4a                 eor w12, w12, w14
  2544: 0e3940f9                 ldr x14, [x8, #0x70]
  2548: ce696d38                 ldrb w14, [x14, x13]
  2552: 8c010e4a                 eor w12, w12, w14
  2556: 0e4540f9                 ldr x14, [x8, #0x88]
  2560: ce696d38                 ldrb w14, [x14, x13]
  2564: 8c010e4a                 eor w12, w12, w14
  2568: 0e4d40f9                 ldr x14, [x8, #0x98]
  2572: ce696d38                 ldrb w14, [x14, x13]
  2576: 8c010e4a                 eor w12, w12, w14
  2580: 0e5d40f9                 ldr x14, [x8, #0xb8]
  2584: ce696d38                 ldrb w14, [x14, x13]
  2588: 8c010e4a                 eor w12, w12, w14
  2592: 0e6140f9                 ldr x14, [x8, #0xc0]
  2596: ce696d38                 ldrb w14, [x14, x13]
  2600: 8c010e4a                 eor w12, w12, w14
  2604: 0e7540f9                 ldr x14, [x8, #0xe8]
  2608: ce696d38                 ldrb w14, [x14, x13]
  2612: 8c010e4a                 eor w12, w12, w14
  2616: 0e8d40f9                 ldr x14, [x8, #0x118]
  2620: ce696d38                 ldrb w14, [x14, x13]
  2624: 8c010e4a                 eor w12, w12, w14
  2628: 4c692d38                 strb w12, [x10, x13]
  2632: ad050091                 add x13, x13, #1
  2636: bf0102eb                 cmp x13, x2
  2640: 83fbff54                 b.lo 0x9c0
  2644: 2a2d40f9                 ldr x10, [x9, #0x58]
  2648: ed030baa                 mov x13, x11
  2652: 0e1540f9                 ldr x14, [x8, #0x28]
  2656: ce696d38                 ldrb w14, [x14, x13]
  2660: ec030e2a                 mov w12, w14
  2664: 0e2940f9                 ldr x14, [x8, #0x50]
  2668: ce696d38                 ldrb w14, [x14, x13]
  2672: 8c010e4a                 eor w12, w12, w14
  2676: 0e3d40f9                 ldr x14, [x8, #0x78]
  2680: ce696d38                 ldrb w14, [x14, x13]
  2684: 8c010e4a                 eor w12, w12, w14
  2688: 0e5140f9                 ldr x14, [x8, #0xa0]
  2692: ce696d38                 ldrb w14, [x14, x13]
  2696: 8c010e4a                 eor w12, w12, w14
  2700: 0e6540f9                 ldr x14, [x8, #0xc8]
  2704: ce696d38                 ldrb w14, [x14, x13]
  2708: 8c010e4a                 eor w12, w12, w14
  2712: 0e7940f9                 ldr x14, [x8, #0xf0]
  2716: ce696d38                 ldrb w14, [x14, x13]
  2720: 8c010e4a                 eor w12, w12, w14
  2724: 4c692d38                 strb w12, [x10, x13]
  2728: ad050091                 add x13, x13, #1
  2732: bf0102eb                 cmp x13, x2
  2736: 63fdff54                 b.lo 0xa5c
  2740: ec37c1ac                 ldp q12, q13, [sp], #0x20
  2744: ea2fc1ac                 ldp q10, q11, [sp], #0x20
  2748: e827c1ac                 ldp q8, q9, [sp], #0x20
  2752: c0035fd6                 ret 
  2756: 00000000                 udf #0
  2760: 00000000                 udf #0
  2764: 00000000                 udf #0
  2768: 00000000                 udf #0
  2772: 00000000                 udf #0
  2776: 00000000                 udf #0
  2780: 00000000                 udf #0
  2784: 00000000                 udf #0
  2788: 00000000                 udf #0
  2792: 00000000                 udf #0
  2796: 00000000                 udf #0
  2800: 00000000                 udf #0
  2804: 00000000                 udf #0
  2808: 00000000                 udf #0
  2812: 00000000                 udf #0
