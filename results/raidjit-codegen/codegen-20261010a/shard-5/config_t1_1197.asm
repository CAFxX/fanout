; config_t1_1197 target=neon twin=nt topo=rdp(7)/enc op=encode U=1
; code_size=2560 twins=["nt"] uops_per_byte=9.0625
     0: e827bfad                 stp q8, q9, [sp, #-0x20]!
     4: ea2fbfad                 stp q10, q11, [sp, #-0x20]!
     8: ec37bfad                 stp q12, q13, [sp, #-0x20]!
    12: e80300aa                 mov x8, x0
    16: e90301aa                 mov x9, x1
    20: eb031faa                 mov x11, xzr
    24: 0e0280d2                 mov x14, #0x10
    28: 4d08ce9a                 udiv x13, x2, x14
    32: 8d1f00b4                 cbz x13, 0x410
    36: 1f2003d5                 nop 
    40: 1f2003d5                 nop 
    44: 1f2003d5                 nop 
    48: 1f2003d5                 nop 
    52: 1f2003d5                 nop 
    56: 1f2003d5                 nop 
    60: 1f2003d5                 nop 
    64: 0a0140f9                 ldr x10, [x8]
    68: 4a010b8b                 add x10, x10, x11
    72: 4c01c03d                 ldr q12, [x10]
    76: 801dac4e                 mov v0.16b, v12.16b
    80: 861dac4e                 mov v6.16b, v12.16b
    84: 0a0540f9                 ldr x10, [x8, #8]
    88: 4a010b8b                 add x10, x10, x11
    92: 4c01c03d                 ldr q12, [x10]
    96: 811dac4e                 mov v1.16b, v12.16b
   100: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   104: 871dac4e                 mov v7.16b, v12.16b
   108: 0a0940f9                 ldr x10, [x8, #0x10]
   112: 4a010b8b                 add x10, x10, x11
   116: 4c01c03d                 ldr q12, [x10]
   120: 821dac4e                 mov v2.16b, v12.16b
   124: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   128: 881dac4e                 mov v8.16b, v12.16b
   132: 0a0d40f9                 ldr x10, [x8, #0x18]
   136: 4a010b8b                 add x10, x10, x11
   140: 4c01c03d                 ldr q12, [x10]
   144: 831dac4e                 mov v3.16b, v12.16b
   148: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   152: 891dac4e                 mov v9.16b, v12.16b
   156: 0a1140f9                 ldr x10, [x8, #0x20]
   160: 4a010b8b                 add x10, x10, x11
   164: 4c01c03d                 ldr q12, [x10]
   168: 841dac4e                 mov v4.16b, v12.16b
   172: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   176: 8a1dac4e                 mov v10.16b, v12.16b
   180: 0a1540f9                 ldr x10, [x8, #0x28]
   184: 4a010b8b                 add x10, x10, x11
   188: 4c01c03d                 ldr q12, [x10]
   192: 851dac4e                 mov v5.16b, v12.16b
   196: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   200: 8b1dac4e                 mov v11.16b, v12.16b
   204: 0a1940f9                 ldr x10, [x8, #0x30]
   208: 4a010b8b                 add x10, x10, x11
   212: 4c01c03d                 ldr q12, [x10]
   216: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
   220: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   224: 0a1d40f9                 ldr x10, [x8, #0x38]
   228: 4a010b8b                 add x10, x10, x11
   232: 4c01c03d                 ldr q12, [x10]
   236: 211c2c6e                 eor v1.16b, v1.16b, v12.16b
   240: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   244: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   248: 0a2140f9                 ldr x10, [x8, #0x40]
   252: 4a010b8b                 add x10, x10, x11
   256: 4c01c03d                 ldr q12, [x10]
   260: 421c2c6e                 eor v2.16b, v2.16b, v12.16b
   264: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   268: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   272: 0a2540f9                 ldr x10, [x8, #0x48]
   276: 4a010b8b                 add x10, x10, x11
   280: 4c01c03d                 ldr q12, [x10]
   284: 631c2c6e                 eor v3.16b, v3.16b, v12.16b
   288: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   292: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   296: 0a2940f9                 ldr x10, [x8, #0x50]
   300: 4a010b8b                 add x10, x10, x11
   304: 4c01c03d                 ldr q12, [x10]
   308: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   312: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   316: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   320: 0a2d40f9                 ldr x10, [x8, #0x58]
   324: 4a010b8b                 add x10, x10, x11
   328: 4c01c03d                 ldr q12, [x10]
   332: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   336: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   340: 0a3140f9                 ldr x10, [x8, #0x60]
   344: 4a010b8b                 add x10, x10, x11
   348: 4c01c03d                 ldr q12, [x10]
   352: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
   356: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   360: 0a3540f9                 ldr x10, [x8, #0x68]
   364: 4a010b8b                 add x10, x10, x11
   368: 4c01c03d                 ldr q12, [x10]
   372: 211c2c6e                 eor v1.16b, v1.16b, v12.16b
   376: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   380: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   384: 0a3940f9                 ldr x10, [x8, #0x70]
   388: 4a010b8b                 add x10, x10, x11
   392: 4c01c03d                 ldr q12, [x10]
   396: 421c2c6e                 eor v2.16b, v2.16b, v12.16b
   400: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   404: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   408: 0a3d40f9                 ldr x10, [x8, #0x78]
   412: 4a010b8b                 add x10, x10, x11
   416: 4c01c03d                 ldr q12, [x10]
   420: 631c2c6e                 eor v3.16b, v3.16b, v12.16b
   424: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   428: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   432: 0a4140f9                 ldr x10, [x8, #0x80]
   436: 4a010b8b                 add x10, x10, x11
   440: 4c01c03d                 ldr q12, [x10]
   444: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   448: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   452: 0a4540f9                 ldr x10, [x8, #0x88]
   456: 4a010b8b                 add x10, x10, x11
   460: 4c01c03d                 ldr q12, [x10]
   464: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   468: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   472: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   476: 0a4940f9                 ldr x10, [x8, #0x90]
   480: 4a010b8b                 add x10, x10, x11
   484: 4c01c03d                 ldr q12, [x10]
   488: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
   492: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   496: 0a4d40f9                 ldr x10, [x8, #0x98]
   500: 4a010b8b                 add x10, x10, x11
   504: 4c01c03d                 ldr q12, [x10]
   508: 211c2c6e                 eor v1.16b, v1.16b, v12.16b
   512: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   516: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   520: 0a5140f9                 ldr x10, [x8, #0xa0]
   524: 4a010b8b                 add x10, x10, x11
   528: 4c01c03d                 ldr q12, [x10]
   532: 421c2c6e                 eor v2.16b, v2.16b, v12.16b
   536: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   540: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   544: 0a5540f9                 ldr x10, [x8, #0xa8]
   548: 4a010b8b                 add x10, x10, x11
   552: 4c01c03d                 ldr q12, [x10]
   556: 631c2c6e                 eor v3.16b, v3.16b, v12.16b
   560: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   564: 0a5940f9                 ldr x10, [x8, #0xb0]
   568: 4a010b8b                 add x10, x10, x11
   572: 4c01c03d                 ldr q12, [x10]
   576: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   580: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   584: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   588: 0a5d40f9                 ldr x10, [x8, #0xb8]
   592: 4a010b8b                 add x10, x10, x11
   596: 4c01c03d                 ldr q12, [x10]
   600: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   604: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   608: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   612: 0a6140f9                 ldr x10, [x8, #0xc0]
   616: 4a010b8b                 add x10, x10, x11
   620: 4c01c03d                 ldr q12, [x10]
   624: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
   628: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   632: 0a6540f9                 ldr x10, [x8, #0xc8]
   636: 4a010b8b                 add x10, x10, x11
   640: 4c01c03d                 ldr q12, [x10]
   644: 211c2c6e                 eor v1.16b, v1.16b, v12.16b
   648: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   652: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   656: 0a6940f9                 ldr x10, [x8, #0xd0]
   660: 4a010b8b                 add x10, x10, x11
   664: 4c01c03d                 ldr q12, [x10]
   668: 421c2c6e                 eor v2.16b, v2.16b, v12.16b
   672: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   676: 0a6d40f9                 ldr x10, [x8, #0xd8]
   680: 4a010b8b                 add x10, x10, x11
   684: 4c01c03d                 ldr q12, [x10]
   688: 631c2c6e                 eor v3.16b, v3.16b, v12.16b
   692: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   696: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   700: 0a7140f9                 ldr x10, [x8, #0xe0]
   704: 4a010b8b                 add x10, x10, x11
   708: 4c01c03d                 ldr q12, [x10]
   712: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   716: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   720: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   724: 0a7540f9                 ldr x10, [x8, #0xe8]
   728: 4a010b8b                 add x10, x10, x11
   732: 4c01c03d                 ldr q12, [x10]
   736: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   740: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   744: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   748: 0a7940f9                 ldr x10, [x8, #0xf0]
   752: 4a010b8b                 add x10, x10, x11
   756: 4c01c03d                 ldr q12, [x10]
   760: 001c2c6e                 eor v0.16b, v0.16b, v12.16b
   764: 6b1d2c6e                 eor v11.16b, v11.16b, v12.16b
   768: 0a7d40f9                 ldr x10, [x8, #0xf8]
   772: 4a010b8b                 add x10, x10, x11
   776: 4c01c03d                 ldr q12, [x10]
   780: 211c2c6e                 eor v1.16b, v1.16b, v12.16b
   784: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   788: 0a8140f9                 ldr x10, [x8, #0x100]
   792: 4a010b8b                 add x10, x10, x11
   796: 4c01c03d                 ldr q12, [x10]
   800: 421c2c6e                 eor v2.16b, v2.16b, v12.16b
   804: c61c2c6e                 eor v6.16b, v6.16b, v12.16b
   808: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   812: 0a8540f9                 ldr x10, [x8, #0x108]
   816: 4a010b8b                 add x10, x10, x11
   820: 4c01c03d                 ldr q12, [x10]
   824: 631c2c6e                 eor v3.16b, v3.16b, v12.16b
   828: e71c2c6e                 eor v7.16b, v7.16b, v12.16b
   832: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   836: 0a8940f9                 ldr x10, [x8, #0x110]
   840: 4a010b8b                 add x10, x10, x11
   844: 4c01c03d                 ldr q12, [x10]
   848: 841c2c6e                 eor v4.16b, v4.16b, v12.16b
   852: 081d2c6e                 eor v8.16b, v8.16b, v12.16b
   856: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   860: 0a8d40f9                 ldr x10, [x8, #0x118]
   864: 4a010b8b                 add x10, x10, x11
   868: 4c01c03d                 ldr q12, [x10]
   872: a51c2c6e                 eor v5.16b, v5.16b, v12.16b
   876: 291d2c6e                 eor v9.16b, v9.16b, v12.16b
   880: 4a1d2c6e                 eor v10.16b, v10.16b, v12.16b
   884: 2a0140f9                 ldr x10, [x9]
   888: 4a010b8b                 add x10, x10, x11
   892: 4001803d                 str q0, [x10]
   896: 2a0540f9                 ldr x10, [x9, #8]
   900: 4a010b8b                 add x10, x10, x11
   904: 4101803d                 str q1, [x10]
   908: 2a0940f9                 ldr x10, [x9, #0x10]
   912: 4a010b8b                 add x10, x10, x11
   916: 4201803d                 str q2, [x10]
   920: 2a0d40f9                 ldr x10, [x9, #0x18]
   924: 4a010b8b                 add x10, x10, x11
   928: 4301803d                 str q3, [x10]
   932: 2a1140f9                 ldr x10, [x9, #0x20]
   936: 4a010b8b                 add x10, x10, x11
   940: 4401803d                 str q4, [x10]
   944: 2a1540f9                 ldr x10, [x9, #0x28]
   948: 4a010b8b                 add x10, x10, x11
   952: 4501803d                 str q5, [x10]
   956: 2a1940f9                 ldr x10, [x9, #0x30]
   960: 4a010b8b                 add x10, x10, x11
   964: 4601803d                 str q6, [x10]
   968: 2a1d40f9                 ldr x10, [x9, #0x38]
   972: 4a010b8b                 add x10, x10, x11
   976: 4701803d                 str q7, [x10]
   980: 2a2140f9                 ldr x10, [x9, #0x40]
   984: 4a010b8b                 add x10, x10, x11
   988: 4801803d                 str q8, [x10]
   992: 2a2540f9                 ldr x10, [x9, #0x48]
   996: 4a010b8b                 add x10, x10, x11
  1000: 4901803d                 str q9, [x10]
  1004: 2a2940f9                 ldr x10, [x9, #0x50]
  1008: 4a010b8b                 add x10, x10, x11
  1012: 4a01803d                 str q10, [x10]
  1016: 2a2d40f9                 ldr x10, [x9, #0x58]
  1020: 4a010b8b                 add x10, x10, x11
  1024: 4b01803d                 str q11, [x10]
  1028: 6b410091                 add x11, x11, #0x10
  1032: ad0500f1                 subs x13, x13, #1
  1036: c1e0ff54                 b.ne 0x24
  1040: 7f0102eb                 cmp x11, x2
  1044: 822d0054                 b.hs 0x9c4
  1048: 2a0140f9                 ldr x10, [x9]
  1052: ed030baa                 mov x13, x11
  1056: 0e0140f9                 ldr x14, [x8]
  1060: ce696d38                 ldrb w14, [x14, x13]
  1064: ec030e2a                 mov w12, w14
  1068: 0e1940f9                 ldr x14, [x8, #0x30]
  1072: ce696d38                 ldrb w14, [x14, x13]
  1076: 8c010e4a                 eor w12, w12, w14
  1080: 0e3140f9                 ldr x14, [x8, #0x60]
  1084: ce696d38                 ldrb w14, [x14, x13]
  1088: 8c010e4a                 eor w12, w12, w14
  1092: 0e4940f9                 ldr x14, [x8, #0x90]
  1096: ce696d38                 ldrb w14, [x14, x13]
  1100: 8c010e4a                 eor w12, w12, w14
  1104: 0e6140f9                 ldr x14, [x8, #0xc0]
  1108: ce696d38                 ldrb w14, [x14, x13]
  1112: 8c010e4a                 eor w12, w12, w14
  1116: 0e7940f9                 ldr x14, [x8, #0xf0]
  1120: ce696d38                 ldrb w14, [x14, x13]
  1124: 8c010e4a                 eor w12, w12, w14
  1128: 4c692d38                 strb w12, [x10, x13]
  1132: ad050091                 add x13, x13, #1
  1136: bf0102eb                 cmp x13, x2
  1140: 63fdff54                 b.lo 0x420
  1144: 2a0540f9                 ldr x10, [x9, #8]
  1148: ed030baa                 mov x13, x11
  1152: 0e0540f9                 ldr x14, [x8, #8]
  1156: ce696d38                 ldrb w14, [x14, x13]
  1160: ec030e2a                 mov w12, w14
  1164: 0e1d40f9                 ldr x14, [x8, #0x38]
  1168: ce696d38                 ldrb w14, [x14, x13]
  1172: 8c010e4a                 eor w12, w12, w14
  1176: 0e3540f9                 ldr x14, [x8, #0x68]
  1180: ce696d38                 ldrb w14, [x14, x13]
  1184: 8c010e4a                 eor w12, w12, w14
  1188: 0e4d40f9                 ldr x14, [x8, #0x98]
  1192: ce696d38                 ldrb w14, [x14, x13]
  1196: 8c010e4a                 eor w12, w12, w14
  1200: 0e6540f9                 ldr x14, [x8, #0xc8]
  1204: ce696d38                 ldrb w14, [x14, x13]
  1208: 8c010e4a                 eor w12, w12, w14
  1212: 0e7d40f9                 ldr x14, [x8, #0xf8]
  1216: ce696d38                 ldrb w14, [x14, x13]
  1220: 8c010e4a                 eor w12, w12, w14
  1224: 4c692d38                 strb w12, [x10, x13]
  1228: ad050091                 add x13, x13, #1
  1232: bf0102eb                 cmp x13, x2
  1236: 63fdff54                 b.lo 0x480
  1240: 2a0940f9                 ldr x10, [x9, #0x10]
  1244: ed030baa                 mov x13, x11
  1248: 0e0940f9                 ldr x14, [x8, #0x10]
  1252: ce696d38                 ldrb w14, [x14, x13]
  1256: ec030e2a                 mov w12, w14
  1260: 0e2140f9                 ldr x14, [x8, #0x40]
  1264: ce696d38                 ldrb w14, [x14, x13]
  1268: 8c010e4a                 eor w12, w12, w14
  1272: 0e3940f9                 ldr x14, [x8, #0x70]
  1276: ce696d38                 ldrb w14, [x14, x13]
  1280: 8c010e4a                 eor w12, w12, w14
  1284: 0e5140f9                 ldr x14, [x8, #0xa0]
  1288: ce696d38                 ldrb w14, [x14, x13]
  1292: 8c010e4a                 eor w12, w12, w14
  1296: 0e6940f9                 ldr x14, [x8, #0xd0]
  1300: ce696d38                 ldrb w14, [x14, x13]
  1304: 8c010e4a                 eor w12, w12, w14
  1308: 0e8140f9                 ldr x14, [x8, #0x100]
  1312: ce696d38                 ldrb w14, [x14, x13]
  1316: 8c010e4a                 eor w12, w12, w14
  1320: 4c692d38                 strb w12, [x10, x13]
  1324: ad050091                 add x13, x13, #1
  1328: bf0102eb                 cmp x13, x2
  1332: 63fdff54                 b.lo 0x4e0
  1336: 2a0d40f9                 ldr x10, [x9, #0x18]
  1340: ed030baa                 mov x13, x11
  1344: 0e0d40f9                 ldr x14, [x8, #0x18]
  1348: ce696d38                 ldrb w14, [x14, x13]
  1352: ec030e2a                 mov w12, w14
  1356: 0e2540f9                 ldr x14, [x8, #0x48]
  1360: ce696d38                 ldrb w14, [x14, x13]
  1364: 8c010e4a                 eor w12, w12, w14
  1368: 0e3d40f9                 ldr x14, [x8, #0x78]
  1372: ce696d38                 ldrb w14, [x14, x13]
  1376: 8c010e4a                 eor w12, w12, w14
  1380: 0e5540f9                 ldr x14, [x8, #0xa8]
  1384: ce696d38                 ldrb w14, [x14, x13]
  1388: 8c010e4a                 eor w12, w12, w14
  1392: 0e6d40f9                 ldr x14, [x8, #0xd8]
  1396: ce696d38                 ldrb w14, [x14, x13]
  1400: 8c010e4a                 eor w12, w12, w14
  1404: 0e8540f9                 ldr x14, [x8, #0x108]
  1408: ce696d38                 ldrb w14, [x14, x13]
  1412: 8c010e4a                 eor w12, w12, w14
  1416: 4c692d38                 strb w12, [x10, x13]
  1420: ad050091                 add x13, x13, #1
  1424: bf0102eb                 cmp x13, x2
  1428: 63fdff54                 b.lo 0x540
  1432: 2a1140f9                 ldr x10, [x9, #0x20]
  1436: ed030baa                 mov x13, x11
  1440: 0e1140f9                 ldr x14, [x8, #0x20]
  1444: ce696d38                 ldrb w14, [x14, x13]
  1448: ec030e2a                 mov w12, w14
  1452: 0e2940f9                 ldr x14, [x8, #0x50]
  1456: ce696d38                 ldrb w14, [x14, x13]
  1460: 8c010e4a                 eor w12, w12, w14
  1464: 0e4140f9                 ldr x14, [x8, #0x80]
  1468: ce696d38                 ldrb w14, [x14, x13]
  1472: 8c010e4a                 eor w12, w12, w14
  1476: 0e5940f9                 ldr x14, [x8, #0xb0]
  1480: ce696d38                 ldrb w14, [x14, x13]
  1484: 8c010e4a                 eor w12, w12, w14
  1488: 0e7140f9                 ldr x14, [x8, #0xe0]
  1492: ce696d38                 ldrb w14, [x14, x13]
  1496: 8c010e4a                 eor w12, w12, w14
  1500: 0e8940f9                 ldr x14, [x8, #0x110]
  1504: ce696d38                 ldrb w14, [x14, x13]
  1508: 8c010e4a                 eor w12, w12, w14
  1512: 4c692d38                 strb w12, [x10, x13]
  1516: ad050091                 add x13, x13, #1
  1520: bf0102eb                 cmp x13, x2
  1524: 63fdff54                 b.lo 0x5a0
  1528: 2a1540f9                 ldr x10, [x9, #0x28]
  1532: ed030baa                 mov x13, x11
  1536: 0e1540f9                 ldr x14, [x8, #0x28]
  1540: ce696d38                 ldrb w14, [x14, x13]
  1544: ec030e2a                 mov w12, w14
  1548: 0e2d40f9                 ldr x14, [x8, #0x58]
  1552: ce696d38                 ldrb w14, [x14, x13]
  1556: 8c010e4a                 eor w12, w12, w14
  1560: 0e4540f9                 ldr x14, [x8, #0x88]
  1564: ce696d38                 ldrb w14, [x14, x13]
  1568: 8c010e4a                 eor w12, w12, w14
  1572: 0e5d40f9                 ldr x14, [x8, #0xb8]
  1576: ce696d38                 ldrb w14, [x14, x13]
  1580: 8c010e4a                 eor w12, w12, w14
  1584: 0e7540f9                 ldr x14, [x8, #0xe8]
  1588: ce696d38                 ldrb w14, [x14, x13]
  1592: 8c010e4a                 eor w12, w12, w14
  1596: 0e8d40f9                 ldr x14, [x8, #0x118]
  1600: ce696d38                 ldrb w14, [x14, x13]
  1604: 8c010e4a                 eor w12, w12, w14
  1608: 4c692d38                 strb w12, [x10, x13]
  1612: ad050091                 add x13, x13, #1
  1616: bf0102eb                 cmp x13, x2
  1620: 63fdff54                 b.lo 0x600
  1624: 2a1940f9                 ldr x10, [x9, #0x30]
  1628: ed030baa                 mov x13, x11
  1632: 0e0140f9                 ldr x14, [x8]
  1636: ce696d38                 ldrb w14, [x14, x13]
  1640: ec030e2a                 mov w12, w14
  1644: 0e0540f9                 ldr x14, [x8, #8]
  1648: ce696d38                 ldrb w14, [x14, x13]
  1652: 8c010e4a                 eor w12, w12, w14
  1656: 0e1d40f9                 ldr x14, [x8, #0x38]
  1660: ce696d38                 ldrb w14, [x14, x13]
  1664: 8c010e4a                 eor w12, w12, w14
  1668: 0e3540f9                 ldr x14, [x8, #0x68]
  1672: ce696d38                 ldrb w14, [x14, x13]
  1676: 8c010e4a                 eor w12, w12, w14
  1680: 0e4540f9                 ldr x14, [x8, #0x88]
  1684: ce696d38                 ldrb w14, [x14, x13]
  1688: 8c010e4a                 eor w12, w12, w14
  1692: 0e4d40f9                 ldr x14, [x8, #0x98]
  1696: ce696d38                 ldrb w14, [x14, x13]
  1700: 8c010e4a                 eor w12, w12, w14
  1704: 0e5940f9                 ldr x14, [x8, #0xb0]
  1708: ce696d38                 ldrb w14, [x14, x13]
  1712: 8c010e4a                 eor w12, w12, w14
  1716: 0e6540f9                 ldr x14, [x8, #0xc8]
  1720: ce696d38                 ldrb w14, [x14, x13]
  1724: 8c010e4a                 eor w12, w12, w14
  1728: 0e6d40f9                 ldr x14, [x8, #0xd8]
  1732: ce696d38                 ldrb w14, [x14, x13]
  1736: 8c010e4a                 eor w12, w12, w14
  1740: 0e7d40f9                 ldr x14, [x8, #0xf8]
  1744: ce696d38                 ldrb w14, [x14, x13]
  1748: 8c010e4a                 eor w12, w12, w14
  1752: 0e8140f9                 ldr x14, [x8, #0x100]
  1756: ce696d38                 ldrb w14, [x14, x13]
  1760: 8c010e4a                 eor w12, w12, w14
  1764: 4c692d38                 strb w12, [x10, x13]
  1768: ad050091                 add x13, x13, #1
  1772: bf0102eb                 cmp x13, x2
  1776: 83fbff54                 b.lo 0x660
  1780: 2a1d40f9                 ldr x10, [x9, #0x38]
  1784: ed030baa                 mov x13, x11
  1788: 0e0540f9                 ldr x14, [x8, #8]
  1792: ce696d38                 ldrb w14, [x14, x13]
  1796: ec030e2a                 mov w12, w14
  1800: 0e0940f9                 ldr x14, [x8, #0x10]
  1804: ce696d38                 ldrb w14, [x14, x13]
  1808: 8c010e4a                 eor w12, w12, w14
  1812: 0e1940f9                 ldr x14, [x8, #0x30]
  1816: ce696d38                 ldrb w14, [x14, x13]
  1820: 8c010e4a                 eor w12, w12, w14
  1824: 0e2140f9                 ldr x14, [x8, #0x40]
  1828: ce696d38                 ldrb w14, [x14, x13]
  1832: 8c010e4a                 eor w12, w12, w14
  1836: 0e3940f9                 ldr x14, [x8, #0x70]
  1840: ce696d38                 ldrb w14, [x14, x13]
  1844: 8c010e4a                 eor w12, w12, w14
  1848: 0e5140f9                 ldr x14, [x8, #0xa0]
  1852: ce696d38                 ldrb w14, [x14, x13]
  1856: 8c010e4a                 eor w12, w12, w14
  1860: 0e5d40f9                 ldr x14, [x8, #0xb8]
  1864: ce696d38                 ldrb w14, [x14, x13]
  1868: 8c010e4a                 eor w12, w12, w14
  1872: 0e6940f9                 ldr x14, [x8, #0xd0]
  1876: ce696d38                 ldrb w14, [x14, x13]
  1880: 8c010e4a                 eor w12, w12, w14
  1884: 0e7140f9                 ldr x14, [x8, #0xe0]
  1888: ce696d38                 ldrb w14, [x14, x13]
  1892: 8c010e4a                 eor w12, w12, w14
  1896: 0e8140f9                 ldr x14, [x8, #0x100]
  1900: ce696d38                 ldrb w14, [x14, x13]
  1904: 8c010e4a                 eor w12, w12, w14
  1908: 0e8540f9                 ldr x14, [x8, #0x108]
  1912: ce696d38                 ldrb w14, [x14, x13]
  1916: 8c010e4a                 eor w12, w12, w14
  1920: 4c692d38                 strb w12, [x10, x13]
  1924: ad050091                 add x13, x13, #1
  1928: bf0102eb                 cmp x13, x2
  1932: 83fbff54                 b.lo 0x6fc
  1936: 2a2140f9                 ldr x10, [x9, #0x40]
  1940: ed030baa                 mov x13, x11
  1944: 0e0940f9                 ldr x14, [x8, #0x10]
  1948: ce696d38                 ldrb w14, [x14, x13]
  1952: ec030e2a                 mov w12, w14
  1956: 0e0d40f9                 ldr x14, [x8, #0x18]
  1960: ce696d38                 ldrb w14, [x14, x13]
  1964: 8c010e4a                 eor w12, w12, w14
  1968: 0e1d40f9                 ldr x14, [x8, #0x38]
  1972: ce696d38                 ldrb w14, [x14, x13]
  1976: 8c010e4a                 eor w12, w12, w14
  1980: 0e2540f9                 ldr x14, [x8, #0x48]
  1984: ce696d38                 ldrb w14, [x14, x13]
  1988: 8c010e4a                 eor w12, w12, w14
  1992: 0e3140f9                 ldr x14, [x8, #0x60]
  1996: ce696d38                 ldrb w14, [x14, x13]
  2000: 8c010e4a                 eor w12, w12, w14
  2004: 0e3d40f9                 ldr x14, [x8, #0x78]
  2008: ce696d38                 ldrb w14, [x14, x13]
  2012: 8c010e4a                 eor w12, w12, w14
  2016: 0e5540f9                 ldr x14, [x8, #0xa8]
  2020: ce696d38                 ldrb w14, [x14, x13]
  2024: 8c010e4a                 eor w12, w12, w14
  2028: 0e6d40f9                 ldr x14, [x8, #0xd8]
  2032: ce696d38                 ldrb w14, [x14, x13]
  2036: 8c010e4a                 eor w12, w12, w14
  2040: 0e7540f9                 ldr x14, [x8, #0xe8]
  2044: ce696d38                 ldrb w14, [x14, x13]
  2048: 8c010e4a                 eor w12, w12, w14
  2052: 0e8540f9                 ldr x14, [x8, #0x108]
  2056: ce696d38                 ldrb w14, [x14, x13]
  2060: 8c010e4a                 eor w12, w12, w14
  2064: 0e8940f9                 ldr x14, [x8, #0x110]
  2068: ce696d38                 ldrb w14, [x14, x13]
  2072: 8c010e4a                 eor w12, w12, w14
  2076: 4c692d38                 strb w12, [x10, x13]
  2080: ad050091                 add x13, x13, #1
  2084: bf0102eb                 cmp x13, x2
  2088: 83fbff54                 b.lo 0x798
  2092: 2a2540f9                 ldr x10, [x9, #0x48]
  2096: ed030baa                 mov x13, x11
  2100: 0e0d40f9                 ldr x14, [x8, #0x18]
  2104: ce696d38                 ldrb w14, [x14, x13]
  2108: ec030e2a                 mov w12, w14
  2112: 0e1140f9                 ldr x14, [x8, #0x20]
  2116: ce696d38                 ldrb w14, [x14, x13]
  2120: 8c010e4a                 eor w12, w12, w14
  2124: 0e2140f9                 ldr x14, [x8, #0x40]
  2128: ce696d38                 ldrb w14, [x14, x13]
  2132: 8c010e4a                 eor w12, w12, w14
  2136: 0e2940f9                 ldr x14, [x8, #0x50]
  2140: ce696d38                 ldrb w14, [x14, x13]
  2144: 8c010e4a                 eor w12, w12, w14
  2148: 0e3540f9                 ldr x14, [x8, #0x68]
  2152: ce696d38                 ldrb w14, [x14, x13]
  2156: 8c010e4a                 eor w12, w12, w14
  2160: 0e4140f9                 ldr x14, [x8, #0x80]
  2164: ce696d38                 ldrb w14, [x14, x13]
  2168: 8c010e4a                 eor w12, w12, w14
  2172: 0e4940f9                 ldr x14, [x8, #0x90]
  2176: ce696d38                 ldrb w14, [x14, x13]
  2180: 8c010e4a                 eor w12, w12, w14
  2184: 0e5940f9                 ldr x14, [x8, #0xb0]
  2188: ce696d38                 ldrb w14, [x14, x13]
  2192: 8c010e4a                 eor w12, w12, w14
  2196: 0e7140f9                 ldr x14, [x8, #0xe0]
  2200: ce696d38                 ldrb w14, [x14, x13]
  2204: 8c010e4a                 eor w12, w12, w14
  2208: 0e8940f9                 ldr x14, [x8, #0x110]
  2212: ce696d38                 ldrb w14, [x14, x13]
  2216: 8c010e4a                 eor w12, w12, w14
  2220: 0e8d40f9                 ldr x14, [x8, #0x118]
  2224: ce696d38                 ldrb w14, [x14, x13]
  2228: 8c010e4a                 eor w12, w12, w14
  2232: 4c692d38                 strb w12, [x10, x13]
  2236: ad050091                 add x13, x13, #1
  2240: bf0102eb                 cmp x13, x2
  2244: 83fbff54                 b.lo 0x834
  2248: 2a2940f9                 ldr x10, [x9, #0x50]
  2252: ed030baa                 mov x13, x11
  2256: 0e1140f9                 ldr x14, [x8, #0x20]
  2260: ce696d38                 ldrb w14, [x14, x13]
  2264: ec030e2a                 mov w12, w14
  2268: 0e1540f9                 ldr x14, [x8, #0x28]
  2272: ce696d38                 ldrb w14, [x14, x13]
  2276: 8c010e4a                 eor w12, w12, w14
  2280: 0e2540f9                 ldr x14, [x8, #0x48]
  2284: ce696d38                 ldrb w14, [x14, x13]
  2288: 8c010e4a                 eor w12, w12, w14
  2292: 0e2d40f9                 ldr x14, [x8, #0x58]
  2296: ce696d38                 ldrb w14, [x14, x13]
  2300: 8c010e4a                 eor w12, w12, w14
  2304: 0e3940f9                 ldr x14, [x8, #0x70]
  2308: ce696d38                 ldrb w14, [x14, x13]
  2312: 8c010e4a                 eor w12, w12, w14
  2316: 0e4540f9                 ldr x14, [x8, #0x88]
  2320: ce696d38                 ldrb w14, [x14, x13]
  2324: 8c010e4a                 eor w12, w12, w14
  2328: 0e4d40f9                 ldr x14, [x8, #0x98]
  2332: ce696d38                 ldrb w14, [x14, x13]
  2336: 8c010e4a                 eor w12, w12, w14
  2340: 0e5d40f9                 ldr x14, [x8, #0xb8]
  2344: ce696d38                 ldrb w14, [x14, x13]
  2348: 8c010e4a                 eor w12, w12, w14
  2352: 0e6140f9                 ldr x14, [x8, #0xc0]
  2356: ce696d38                 ldrb w14, [x14, x13]
  2360: 8c010e4a                 eor w12, w12, w14
  2364: 0e7540f9                 ldr x14, [x8, #0xe8]
  2368: ce696d38                 ldrb w14, [x14, x13]
  2372: 8c010e4a                 eor w12, w12, w14
  2376: 0e8d40f9                 ldr x14, [x8, #0x118]
  2380: ce696d38                 ldrb w14, [x14, x13]
  2384: 8c010e4a                 eor w12, w12, w14
  2388: 4c692d38                 strb w12, [x10, x13]
  2392: ad050091                 add x13, x13, #1
  2396: bf0102eb                 cmp x13, x2
  2400: 83fbff54                 b.lo 0x8d0
  2404: 2a2d40f9                 ldr x10, [x9, #0x58]
  2408: ed030baa                 mov x13, x11
  2412: 0e1540f9                 ldr x14, [x8, #0x28]
  2416: ce696d38                 ldrb w14, [x14, x13]
  2420: ec030e2a                 mov w12, w14
  2424: 0e2940f9                 ldr x14, [x8, #0x50]
  2428: ce696d38                 ldrb w14, [x14, x13]
  2432: 8c010e4a                 eor w12, w12, w14
  2436: 0e3d40f9                 ldr x14, [x8, #0x78]
  2440: ce696d38                 ldrb w14, [x14, x13]
  2444: 8c010e4a                 eor w12, w12, w14
  2448: 0e5140f9                 ldr x14, [x8, #0xa0]
  2452: ce696d38                 ldrb w14, [x14, x13]
  2456: 8c010e4a                 eor w12, w12, w14
  2460: 0e6540f9                 ldr x14, [x8, #0xc8]
  2464: ce696d38                 ldrb w14, [x14, x13]
  2468: 8c010e4a                 eor w12, w12, w14
  2472: 0e7940f9                 ldr x14, [x8, #0xf0]
  2476: ce696d38                 ldrb w14, [x14, x13]
  2480: 8c010e4a                 eor w12, w12, w14
  2484: 4c692d38                 strb w12, [x10, x13]
  2488: ad050091                 add x13, x13, #1
  2492: bf0102eb                 cmp x13, x2
  2496: 63fdff54                 b.lo 0x96c
  2500: ec37c1ac                 ldp q12, q13, [sp], #0x20
  2504: ea2fc1ac                 ldp q10, q11, [sp], #0x20
  2508: e827c1ac                 ldp q8, q9, [sp], #0x20
  2512: c0035fd6                 ret 
  2516: 00000000                 udf #0
  2520: 00000000                 udf #0
  2524: 00000000                 udf #0
  2528: 00000000                 udf #0
  2532: 00000000                 udf #0
  2536: 00000000                 udf #0
  2540: 00000000                 udf #0
  2544: 00000000                 udf #0
  2548: 00000000                 udf #0
  2552: 00000000                 udf #0
  2556: 00000000                 udf #0
