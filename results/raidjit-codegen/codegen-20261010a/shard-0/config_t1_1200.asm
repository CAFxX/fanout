; config_t1_1200 target=neon twin=notail topo=rdp(7)/enc op=encode U=1
; code_size=1088 twins=["notail"] uops_per_byte=9.0625
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
  1040: ec37c1ac                 ldp q12, q13, [sp], #0x20
  1044: ea2fc1ac                 ldp q10, q11, [sp], #0x20
  1048: e827c1ac                 ldp q8, q9, [sp], #0x20
  1052: c0035fd6                 ret 
  1056: 00000000                 udf #0
  1060: 00000000                 udf #0
  1064: 00000000                 udf #0
  1068: 00000000                 udf #0
  1072: 00000000                 udf #0
  1076: 00000000                 udf #0
  1080: 00000000                 udf #0
  1084: 00000000                 udf #0
