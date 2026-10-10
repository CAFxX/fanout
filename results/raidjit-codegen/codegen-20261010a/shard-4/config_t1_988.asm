; config_t1_988 target=sve2 twin=baseline topo=liberation(8)/enc op=encode U=1
; code_size=4864 twins=[] uops_per_byte=18.3125
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: ed030caa                 mov x13, x12
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: ae4000b4                 cbz x14, 0x830
    32: 0a0140f9                 ldr x10, [x8]
    36: 4a010b8b                 add x10, x10, x11
    40: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
    44: d0306604                 mov z16.d, z6.d
    48: db306604                 mov z27.d, z6.d
    52: 0a0540f9                 ldr x10, [x8, #8]
    56: 4a010b8b                 add x10, x10, x11
    60: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
    64: d1306604                 mov z17.d, z6.d
    68: dc306604                 mov z28.d, z6.d
    72: 0a0940f9                 ldr x10, [x8, #0x10]
    76: 4a010b8b                 add x10, x10, x11
    80: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
    84: d2306604                 mov z18.d, z6.d
    88: dd306604                 mov z29.d, z6.d
    92: 0a0d40f9                 ldr x10, [x8, #0x18]
    96: 4a010b8b                 add x10, x10, x11
   100: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   104: d3306604                 mov z19.d, z6.d
   108: de306604                 mov z30.d, z6.d
   112: 0a1140f9                 ldr x10, [x8, #0x20]
   116: 4a010b8b                 add x10, x10, x11
   120: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   124: d4306604                 mov z20.d, z6.d
   128: df306604                 mov z31.d, z6.d
   132: 0a1540f9                 ldr x10, [x8, #0x28]
   136: 4a010b8b                 add x10, x10, x11
   140: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   144: d5306604                 mov z21.d, z6.d
   148: c0306604                 mov z0.d, z6.d
   152: 0a1940f9                 ldr x10, [x8, #0x30]
   156: 4a010b8b                 add x10, x10, x11
   160: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   164: d6306604                 mov z22.d, z6.d
   168: c1306604                 mov z1.d, z6.d
   172: 0a1d40f9                 ldr x10, [x8, #0x38]
   176: 4a010b8b                 add x10, x10, x11
   180: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   184: d7306604                 mov z23.d, z6.d
   188: c2306604                 mov z2.d, z6.d
   192: 0a2140f9                 ldr x10, [x8, #0x40]
   196: 4a010b8b                 add x10, x10, x11
   200: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   204: d8306604                 mov z24.d, z6.d
   208: c3306604                 mov z3.d, z6.d
   212: 0a2540f9                 ldr x10, [x8, #0x48]
   216: 4a010b8b                 add x10, x10, x11
   220: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   224: d9306604                 mov z25.d, z6.d
   228: c4306604                 mov z4.d, z6.d
   232: 0a2940f9                 ldr x10, [x8, #0x50]
   236: 4a010b8b                 add x10, x10, x11
   240: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   244: da306604                 mov z26.d, z6.d
   248: c5306604                 mov z5.d, z6.d
   252: 0a2d40f9                 ldr x10, [x8, #0x58]
   256: 4a010b8b                 add x10, x10, x11
   260: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   264: d0001904                 eor z16.b, p0/m, z16.b, z6.b
   268: c5001904                 eor z5.b, p0/m, z5.b, z6.b
   272: 0a3140f9                 ldr x10, [x8, #0x60]
   276: 4a010b8b                 add x10, x10, x11
   280: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   284: d1001904                 eor z17.b, p0/m, z17.b, z6.b
   288: db001904                 eor z27.b, p0/m, z27.b, z6.b
   292: 0a3540f9                 ldr x10, [x8, #0x68]
   296: 4a010b8b                 add x10, x10, x11
   300: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   304: d2001904                 eor z18.b, p0/m, z18.b, z6.b
   308: dc001904                 eor z28.b, p0/m, z28.b, z6.b
   312: 0a3940f9                 ldr x10, [x8, #0x70]
   316: 4a010b8b                 add x10, x10, x11
   320: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   324: d3001904                 eor z19.b, p0/m, z19.b, z6.b
   328: dd001904                 eor z29.b, p0/m, z29.b, z6.b
   332: 0a3d40f9                 ldr x10, [x8, #0x78]
   336: 4a010b8b                 add x10, x10, x11
   340: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   344: d4001904                 eor z20.b, p0/m, z20.b, z6.b
   348: de001904                 eor z30.b, p0/m, z30.b, z6.b
   352: 0a4140f9                 ldr x10, [x8, #0x80]
   356: 4a010b8b                 add x10, x10, x11
   360: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   364: d5001904                 eor z21.b, p0/m, z21.b, z6.b
   368: df001904                 eor z31.b, p0/m, z31.b, z6.b
   372: c0001904                 eor z0.b, p0/m, z0.b, z6.b
   376: 0a4540f9                 ldr x10, [x8, #0x88]
   380: 4a010b8b                 add x10, x10, x11
   384: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   388: d6001904                 eor z22.b, p0/m, z22.b, z6.b
   392: c0001904                 eor z0.b, p0/m, z0.b, z6.b
   396: 0a4940f9                 ldr x10, [x8, #0x90]
   400: 4a010b8b                 add x10, x10, x11
   404: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   408: d7001904                 eor z23.b, p0/m, z23.b, z6.b
   412: c1001904                 eor z1.b, p0/m, z1.b, z6.b
   416: 0a4d40f9                 ldr x10, [x8, #0x98]
   420: 4a010b8b                 add x10, x10, x11
   424: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   428: d8001904                 eor z24.b, p0/m, z24.b, z6.b
   432: c2001904                 eor z2.b, p0/m, z2.b, z6.b
   436: 0a5140f9                 ldr x10, [x8, #0xa0]
   440: 4a010b8b                 add x10, x10, x11
   444: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   448: d9001904                 eor z25.b, p0/m, z25.b, z6.b
   452: c3001904                 eor z3.b, p0/m, z3.b, z6.b
   456: 0a5540f9                 ldr x10, [x8, #0xa8]
   460: 4a010b8b                 add x10, x10, x11
   464: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   468: da001904                 eor z26.b, p0/m, z26.b, z6.b
   472: c4001904                 eor z4.b, p0/m, z4.b, z6.b
   476: 0a5940f9                 ldr x10, [x8, #0xb0]
   480: 4a010b8b                 add x10, x10, x11
   484: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   488: d0001904                 eor z16.b, p0/m, z16.b, z6.b
   492: c4001904                 eor z4.b, p0/m, z4.b, z6.b
   496: c5001904                 eor z5.b, p0/m, z5.b, z6.b
   500: 0a5d40f9                 ldr x10, [x8, #0xb8]
   504: 4a010b8b                 add x10, x10, x11
   508: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   512: d1001904                 eor z17.b, p0/m, z17.b, z6.b
   516: c5001904                 eor z5.b, p0/m, z5.b, z6.b
   520: 0a6140f9                 ldr x10, [x8, #0xc0]
   524: 4a010b8b                 add x10, x10, x11
   528: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   532: d2001904                 eor z18.b, p0/m, z18.b, z6.b
   536: db001904                 eor z27.b, p0/m, z27.b, z6.b
   540: 0a6540f9                 ldr x10, [x8, #0xc8]
   544: 4a010b8b                 add x10, x10, x11
   548: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   552: d3001904                 eor z19.b, p0/m, z19.b, z6.b
   556: dc001904                 eor z28.b, p0/m, z28.b, z6.b
   560: 0a6940f9                 ldr x10, [x8, #0xd0]
   564: 4a010b8b                 add x10, x10, x11
   568: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   572: d4001904                 eor z20.b, p0/m, z20.b, z6.b
   576: dd001904                 eor z29.b, p0/m, z29.b, z6.b
   580: 0a6d40f9                 ldr x10, [x8, #0xd8]
   584: 4a010b8b                 add x10, x10, x11
   588: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   592: d5001904                 eor z21.b, p0/m, z21.b, z6.b
   596: de001904                 eor z30.b, p0/m, z30.b, z6.b
   600: 0a7140f9                 ldr x10, [x8, #0xe0]
   604: 4a010b8b                 add x10, x10, x11
   608: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   612: d6001904                 eor z22.b, p0/m, z22.b, z6.b
   616: df001904                 eor z31.b, p0/m, z31.b, z6.b
   620: 0a7540f9                 ldr x10, [x8, #0xe8]
   624: 4a010b8b                 add x10, x10, x11
   628: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   632: d7001904                 eor z23.b, p0/m, z23.b, z6.b
   636: c0001904                 eor z0.b, p0/m, z0.b, z6.b
   640: 0a7940f9                 ldr x10, [x8, #0xf0]
   644: 4a010b8b                 add x10, x10, x11
   648: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   652: d8001904                 eor z24.b, p0/m, z24.b, z6.b
   656: c1001904                 eor z1.b, p0/m, z1.b, z6.b
   660: 0a7d40f9                 ldr x10, [x8, #0xf8]
   664: 4a010b8b                 add x10, x10, x11
   668: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   672: d9001904                 eor z25.b, p0/m, z25.b, z6.b
   676: c2001904                 eor z2.b, p0/m, z2.b, z6.b
   680: 0a8140f9                 ldr x10, [x8, #0x100]
   684: 4a010b8b                 add x10, x10, x11
   688: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   692: da001904                 eor z26.b, p0/m, z26.b, z6.b
   696: c3001904                 eor z3.b, p0/m, z3.b, z6.b
   700: 0a8540f9                 ldr x10, [x8, #0x108]
   704: 4a010b8b                 add x10, x10, x11
   708: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   712: d0001904                 eor z16.b, p0/m, z16.b, z6.b
   716: c3001904                 eor z3.b, p0/m, z3.b, z6.b
   720: 0a8940f9                 ldr x10, [x8, #0x110]
   724: 4a010b8b                 add x10, x10, x11
   728: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   732: d1001904                 eor z17.b, p0/m, z17.b, z6.b
   736: c4001904                 eor z4.b, p0/m, z4.b, z6.b
   740: 0a8d40f9                 ldr x10, [x8, #0x118]
   744: 4a010b8b                 add x10, x10, x11
   748: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   752: d2001904                 eor z18.b, p0/m, z18.b, z6.b
   756: c5001904                 eor z5.b, p0/m, z5.b, z6.b
   760: 0a9140f9                 ldr x10, [x8, #0x120]
   764: 4a010b8b                 add x10, x10, x11
   768: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   772: d3001904                 eor z19.b, p0/m, z19.b, z6.b
   776: db001904                 eor z27.b, p0/m, z27.b, z6.b
   780: 0a9540f9                 ldr x10, [x8, #0x128]
   784: 4a010b8b                 add x10, x10, x11
   788: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   792: d4001904                 eor z20.b, p0/m, z20.b, z6.b
   796: dc001904                 eor z28.b, p0/m, z28.b, z6.b
   800: 0a9940f9                 ldr x10, [x8, #0x130]
   804: 4a010b8b                 add x10, x10, x11
   808: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   812: d5001904                 eor z21.b, p0/m, z21.b, z6.b
   816: dd001904                 eor z29.b, p0/m, z29.b, z6.b
   820: 0a9d40f9                 ldr x10, [x8, #0x138]
   824: 4a010b8b                 add x10, x10, x11
   828: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   832: d6001904                 eor z22.b, p0/m, z22.b, z6.b
   836: de001904                 eor z30.b, p0/m, z30.b, z6.b
   840: df001904                 eor z31.b, p0/m, z31.b, z6.b
   844: 0aa140f9                 ldr x10, [x8, #0x140]
   848: 4a010b8b                 add x10, x10, x11
   852: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   856: d7001904                 eor z23.b, p0/m, z23.b, z6.b
   860: df001904                 eor z31.b, p0/m, z31.b, z6.b
   864: 0aa540f9                 ldr x10, [x8, #0x148]
   868: 4a010b8b                 add x10, x10, x11
   872: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   876: d8001904                 eor z24.b, p0/m, z24.b, z6.b
   880: c0001904                 eor z0.b, p0/m, z0.b, z6.b
   884: 0aa940f9                 ldr x10, [x8, #0x150]
   888: 4a010b8b                 add x10, x10, x11
   892: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   896: d9001904                 eor z25.b, p0/m, z25.b, z6.b
   900: c1001904                 eor z1.b, p0/m, z1.b, z6.b
   904: 0aad40f9                 ldr x10, [x8, #0x158]
   908: 4a010b8b                 add x10, x10, x11
   912: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   916: da001904                 eor z26.b, p0/m, z26.b, z6.b
   920: c2001904                 eor z2.b, p0/m, z2.b, z6.b
   924: 0ab140f9                 ldr x10, [x8, #0x160]
   928: 4a010b8b                 add x10, x10, x11
   932: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   936: d0001904                 eor z16.b, p0/m, z16.b, z6.b
   940: c2001904                 eor z2.b, p0/m, z2.b, z6.b
   944: 0ab540f9                 ldr x10, [x8, #0x168]
   948: 4a010b8b                 add x10, x10, x11
   952: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   956: d1001904                 eor z17.b, p0/m, z17.b, z6.b
   960: c3001904                 eor z3.b, p0/m, z3.b, z6.b
   964: c4001904                 eor z4.b, p0/m, z4.b, z6.b
   968: 0ab940f9                 ldr x10, [x8, #0x170]
   972: 4a010b8b                 add x10, x10, x11
   976: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
   980: d2001904                 eor z18.b, p0/m, z18.b, z6.b
   984: c4001904                 eor z4.b, p0/m, z4.b, z6.b
   988: 0abd40f9                 ldr x10, [x8, #0x178]
   992: 4a010b8b                 add x10, x10, x11
   996: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1000: d3001904                 eor z19.b, p0/m, z19.b, z6.b
  1004: c5001904                 eor z5.b, p0/m, z5.b, z6.b
  1008: 0ac140f9                 ldr x10, [x8, #0x180]
  1012: 4a010b8b                 add x10, x10, x11
  1016: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1020: d4001904                 eor z20.b, p0/m, z20.b, z6.b
  1024: db001904                 eor z27.b, p0/m, z27.b, z6.b
  1028: 0ac540f9                 ldr x10, [x8, #0x188]
  1032: 4a010b8b                 add x10, x10, x11
  1036: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1040: d5001904                 eor z21.b, p0/m, z21.b, z6.b
  1044: dc001904                 eor z28.b, p0/m, z28.b, z6.b
  1048: 0ac940f9                 ldr x10, [x8, #0x190]
  1052: 4a010b8b                 add x10, x10, x11
  1056: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1060: d6001904                 eor z22.b, p0/m, z22.b, z6.b
  1064: dd001904                 eor z29.b, p0/m, z29.b, z6.b
  1068: 0acd40f9                 ldr x10, [x8, #0x198]
  1072: 4a010b8b                 add x10, x10, x11
  1076: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1080: d7001904                 eor z23.b, p0/m, z23.b, z6.b
  1084: de001904                 eor z30.b, p0/m, z30.b, z6.b
  1088: 0ad140f9                 ldr x10, [x8, #0x1a0]
  1092: 4a010b8b                 add x10, x10, x11
  1096: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1100: d8001904                 eor z24.b, p0/m, z24.b, z6.b
  1104: df001904                 eor z31.b, p0/m, z31.b, z6.b
  1108: 0ad540f9                 ldr x10, [x8, #0x1a8]
  1112: 4a010b8b                 add x10, x10, x11
  1116: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1120: d9001904                 eor z25.b, p0/m, z25.b, z6.b
  1124: c0001904                 eor z0.b, p0/m, z0.b, z6.b
  1128: 0ad940f9                 ldr x10, [x8, #0x1b0]
  1132: 4a010b8b                 add x10, x10, x11
  1136: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1140: da001904                 eor z26.b, p0/m, z26.b, z6.b
  1144: c1001904                 eor z1.b, p0/m, z1.b, z6.b
  1148: 0add40f9                 ldr x10, [x8, #0x1b8]
  1152: 4a010b8b                 add x10, x10, x11
  1156: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1160: d0001904                 eor z16.b, p0/m, z16.b, z6.b
  1164: c1001904                 eor z1.b, p0/m, z1.b, z6.b
  1168: 0ae140f9                 ldr x10, [x8, #0x1c0]
  1172: 4a010b8b                 add x10, x10, x11
  1176: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1180: d1001904                 eor z17.b, p0/m, z17.b, z6.b
  1184: c2001904                 eor z2.b, p0/m, z2.b, z6.b
  1188: 0ae540f9                 ldr x10, [x8, #0x1c8]
  1192: 4a010b8b                 add x10, x10, x11
  1196: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1200: d2001904                 eor z18.b, p0/m, z18.b, z6.b
  1204: c3001904                 eor z3.b, p0/m, z3.b, z6.b
  1208: 0ae940f9                 ldr x10, [x8, #0x1d0]
  1212: 4a010b8b                 add x10, x10, x11
  1216: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1220: d3001904                 eor z19.b, p0/m, z19.b, z6.b
  1224: c4001904                 eor z4.b, p0/m, z4.b, z6.b
  1228: 0aed40f9                 ldr x10, [x8, #0x1d8]
  1232: 4a010b8b                 add x10, x10, x11
  1236: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1240: d4001904                 eor z20.b, p0/m, z20.b, z6.b
  1244: c5001904                 eor z5.b, p0/m, z5.b, z6.b
  1248: 0af140f9                 ldr x10, [x8, #0x1e0]
  1252: 4a010b8b                 add x10, x10, x11
  1256: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1260: d5001904                 eor z21.b, p0/m, z21.b, z6.b
  1264: db001904                 eor z27.b, p0/m, z27.b, z6.b
  1268: 0af540f9                 ldr x10, [x8, #0x1e8]
  1272: 4a010b8b                 add x10, x10, x11
  1276: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1280: d6001904                 eor z22.b, p0/m, z22.b, z6.b
  1284: dc001904                 eor z28.b, p0/m, z28.b, z6.b
  1288: 0af940f9                 ldr x10, [x8, #0x1f0]
  1292: 4a010b8b                 add x10, x10, x11
  1296: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1300: d7001904                 eor z23.b, p0/m, z23.b, z6.b
  1304: dd001904                 eor z29.b, p0/m, z29.b, z6.b
  1308: de001904                 eor z30.b, p0/m, z30.b, z6.b
  1312: 0afd40f9                 ldr x10, [x8, #0x1f8]
  1316: 4a010b8b                 add x10, x10, x11
  1320: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1324: d8001904                 eor z24.b, p0/m, z24.b, z6.b
  1328: de001904                 eor z30.b, p0/m, z30.b, z6.b
  1332: 0a0141f9                 ldr x10, [x8, #0x200]
  1336: 4a010b8b                 add x10, x10, x11
  1340: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1344: d9001904                 eor z25.b, p0/m, z25.b, z6.b
  1348: df001904                 eor z31.b, p0/m, z31.b, z6.b
  1352: 0a0541f9                 ldr x10, [x8, #0x208]
  1356: 4a010b8b                 add x10, x10, x11
  1360: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1364: da001904                 eor z26.b, p0/m, z26.b, z6.b
  1368: c0001904                 eor z0.b, p0/m, z0.b, z6.b
  1372: 0a0941f9                 ldr x10, [x8, #0x210]
  1376: 4a010b8b                 add x10, x10, x11
  1380: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1384: d0001904                 eor z16.b, p0/m, z16.b, z6.b
  1388: c0001904                 eor z0.b, p0/m, z0.b, z6.b
  1392: 0a0d41f9                 ldr x10, [x8, #0x218]
  1396: 4a010b8b                 add x10, x10, x11
  1400: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1404: d1001904                 eor z17.b, p0/m, z17.b, z6.b
  1408: c1001904                 eor z1.b, p0/m, z1.b, z6.b
  1412: 0a1141f9                 ldr x10, [x8, #0x220]
  1416: 4a010b8b                 add x10, x10, x11
  1420: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1424: d2001904                 eor z18.b, p0/m, z18.b, z6.b
  1428: c2001904                 eor z2.b, p0/m, z2.b, z6.b
  1432: c3001904                 eor z3.b, p0/m, z3.b, z6.b
  1436: 0a1541f9                 ldr x10, [x8, #0x228]
  1440: 4a010b8b                 add x10, x10, x11
  1444: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1448: d3001904                 eor z19.b, p0/m, z19.b, z6.b
  1452: c3001904                 eor z3.b, p0/m, z3.b, z6.b
  1456: 0a1941f9                 ldr x10, [x8, #0x230]
  1460: 4a010b8b                 add x10, x10, x11
  1464: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1468: d4001904                 eor z20.b, p0/m, z20.b, z6.b
  1472: c4001904                 eor z4.b, p0/m, z4.b, z6.b
  1476: 0a1d41f9                 ldr x10, [x8, #0x238]
  1480: 4a010b8b                 add x10, x10, x11
  1484: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1488: d5001904                 eor z21.b, p0/m, z21.b, z6.b
  1492: c5001904                 eor z5.b, p0/m, z5.b, z6.b
  1496: 0a2141f9                 ldr x10, [x8, #0x240]
  1500: 4a010b8b                 add x10, x10, x11
  1504: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1508: d6001904                 eor z22.b, p0/m, z22.b, z6.b
  1512: db001904                 eor z27.b, p0/m, z27.b, z6.b
  1516: 0a2541f9                 ldr x10, [x8, #0x248]
  1520: 4a010b8b                 add x10, x10, x11
  1524: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1528: d7001904                 eor z23.b, p0/m, z23.b, z6.b
  1532: dc001904                 eor z28.b, p0/m, z28.b, z6.b
  1536: 0a2941f9                 ldr x10, [x8, #0x250]
  1540: 4a010b8b                 add x10, x10, x11
  1544: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1548: d8001904                 eor z24.b, p0/m, z24.b, z6.b
  1552: dd001904                 eor z29.b, p0/m, z29.b, z6.b
  1556: 0a2d41f9                 ldr x10, [x8, #0x258]
  1560: 4a010b8b                 add x10, x10, x11
  1564: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1568: d9001904                 eor z25.b, p0/m, z25.b, z6.b
  1572: de001904                 eor z30.b, p0/m, z30.b, z6.b
  1576: 0a3141f9                 ldr x10, [x8, #0x260]
  1580: 4a010b8b                 add x10, x10, x11
  1584: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1588: da001904                 eor z26.b, p0/m, z26.b, z6.b
  1592: df001904                 eor z31.b, p0/m, z31.b, z6.b
  1596: 0a3541f9                 ldr x10, [x8, #0x268]
  1600: 4a010b8b                 add x10, x10, x11
  1604: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1608: d0001904                 eor z16.b, p0/m, z16.b, z6.b
  1612: df001904                 eor z31.b, p0/m, z31.b, z6.b
  1616: 0a3941f9                 ldr x10, [x8, #0x270]
  1620: 4a010b8b                 add x10, x10, x11
  1624: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1628: d1001904                 eor z17.b, p0/m, z17.b, z6.b
  1632: c0001904                 eor z0.b, p0/m, z0.b, z6.b
  1636: 0a3d41f9                 ldr x10, [x8, #0x278]
  1640: 4a010b8b                 add x10, x10, x11
  1644: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1648: d2001904                 eor z18.b, p0/m, z18.b, z6.b
  1652: c1001904                 eor z1.b, p0/m, z1.b, z6.b
  1656: 0a4141f9                 ldr x10, [x8, #0x280]
  1660: 4a010b8b                 add x10, x10, x11
  1664: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1668: d3001904                 eor z19.b, p0/m, z19.b, z6.b
  1672: c2001904                 eor z2.b, p0/m, z2.b, z6.b
  1676: 0a4541f9                 ldr x10, [x8, #0x288]
  1680: 4a010b8b                 add x10, x10, x11
  1684: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1688: d4001904                 eor z20.b, p0/m, z20.b, z6.b
  1692: c3001904                 eor z3.b, p0/m, z3.b, z6.b
  1696: 0a4941f9                 ldr x10, [x8, #0x290]
  1700: 4a010b8b                 add x10, x10, x11
  1704: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1708: d5001904                 eor z21.b, p0/m, z21.b, z6.b
  1712: c4001904                 eor z4.b, p0/m, z4.b, z6.b
  1716: 0a4d41f9                 ldr x10, [x8, #0x298]
  1720: 4a010b8b                 add x10, x10, x11
  1724: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1728: d6001904                 eor z22.b, p0/m, z22.b, z6.b
  1732: c5001904                 eor z5.b, p0/m, z5.b, z6.b
  1736: 0a5141f9                 ldr x10, [x8, #0x2a0]
  1740: 4a010b8b                 add x10, x10, x11
  1744: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1748: d7001904                 eor z23.b, p0/m, z23.b, z6.b
  1752: db001904                 eor z27.b, p0/m, z27.b, z6.b
  1756: 0a5541f9                 ldr x10, [x8, #0x2a8]
  1760: 4a010b8b                 add x10, x10, x11
  1764: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1768: d8001904                 eor z24.b, p0/m, z24.b, z6.b
  1772: dc001904                 eor z28.b, p0/m, z28.b, z6.b
  1776: dd001904                 eor z29.b, p0/m, z29.b, z6.b
  1780: 0a5941f9                 ldr x10, [x8, #0x2b0]
  1784: 4a010b8b                 add x10, x10, x11
  1788: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1792: d9001904                 eor z25.b, p0/m, z25.b, z6.b
  1796: dd001904                 eor z29.b, p0/m, z29.b, z6.b
  1800: 0a5d41f9                 ldr x10, [x8, #0x2b8]
  1804: 4a010b8b                 add x10, x10, x11
  1808: 46a100a4                 ld1b { z6.b }, p0/z, [x10]
  1812: da001904                 eor z26.b, p0/m, z26.b, z6.b
  1816: de001904                 eor z30.b, p0/m, z30.b, z6.b
  1820: 2a0140f9                 ldr x10, [x9]
  1824: 4a010b8b                 add x10, x10, x11
  1828: 50e100e4                 st1b { z16.b }, p0, [x10]
  1832: 2a0540f9                 ldr x10, [x9, #8]
  1836: 4a010b8b                 add x10, x10, x11
  1840: 51e100e4                 st1b { z17.b }, p0, [x10]
  1844: 2a0940f9                 ldr x10, [x9, #0x10]
  1848: 4a010b8b                 add x10, x10, x11
  1852: 52e100e4                 st1b { z18.b }, p0, [x10]
  1856: 2a0d40f9                 ldr x10, [x9, #0x18]
  1860: 4a010b8b                 add x10, x10, x11
  1864: 53e100e4                 st1b { z19.b }, p0, [x10]
  1868: 2a1140f9                 ldr x10, [x9, #0x20]
  1872: 4a010b8b                 add x10, x10, x11
  1876: 54e100e4                 st1b { z20.b }, p0, [x10]
  1880: 2a1540f9                 ldr x10, [x9, #0x28]
  1884: 4a010b8b                 add x10, x10, x11
  1888: 55e100e4                 st1b { z21.b }, p0, [x10]
  1892: 2a1940f9                 ldr x10, [x9, #0x30]
  1896: 4a010b8b                 add x10, x10, x11
  1900: 56e100e4                 st1b { z22.b }, p0, [x10]
  1904: 2a1d40f9                 ldr x10, [x9, #0x38]
  1908: 4a010b8b                 add x10, x10, x11
  1912: 57e100e4                 st1b { z23.b }, p0, [x10]
  1916: 2a2140f9                 ldr x10, [x9, #0x40]
  1920: 4a010b8b                 add x10, x10, x11
  1924: 58e100e4                 st1b { z24.b }, p0, [x10]
  1928: 2a2540f9                 ldr x10, [x9, #0x48]
  1932: 4a010b8b                 add x10, x10, x11
  1936: 59e100e4                 st1b { z25.b }, p0, [x10]
  1940: 2a2940f9                 ldr x10, [x9, #0x50]
  1944: 4a010b8b                 add x10, x10, x11
  1948: 5ae100e4                 st1b { z26.b }, p0, [x10]
  1952: 2a2d40f9                 ldr x10, [x9, #0x58]
  1956: 4a010b8b                 add x10, x10, x11
  1960: 5be100e4                 st1b { z27.b }, p0, [x10]
  1964: 2a3140f9                 ldr x10, [x9, #0x60]
  1968: 4a010b8b                 add x10, x10, x11
  1972: 5ce100e4                 st1b { z28.b }, p0, [x10]
  1976: 2a3540f9                 ldr x10, [x9, #0x68]
  1980: 4a010b8b                 add x10, x10, x11
  1984: 5de100e4                 st1b { z29.b }, p0, [x10]
  1988: 2a3940f9                 ldr x10, [x9, #0x70]
  1992: 4a010b8b                 add x10, x10, x11
  1996: 5ee100e4                 st1b { z30.b }, p0, [x10]
  2000: 2a3d40f9                 ldr x10, [x9, #0x78]
  2004: 4a010b8b                 add x10, x10, x11
  2008: 5fe100e4                 st1b { z31.b }, p0, [x10]
  2012: 2a4140f9                 ldr x10, [x9, #0x80]
  2016: 4a010b8b                 add x10, x10, x11
  2020: 40e100e4                 st1b { z0.b }, p0, [x10]
  2024: 2a4540f9                 ldr x10, [x9, #0x88]
  2028: 4a010b8b                 add x10, x10, x11
  2032: 41e100e4                 st1b { z1.b }, p0, [x10]
  2036: 2a4940f9                 ldr x10, [x9, #0x90]
  2040: 4a010b8b                 add x10, x10, x11
  2044: 42e100e4                 st1b { z2.b }, p0, [x10]
  2048: 2a4d40f9                 ldr x10, [x9, #0x98]
  2052: 4a010b8b                 add x10, x10, x11
  2056: 43e100e4                 st1b { z3.b }, p0, [x10]
  2060: 2a5140f9                 ldr x10, [x9, #0xa0]
  2064: 4a010b8b                 add x10, x10, x11
  2068: 44e100e4                 st1b { z4.b }, p0, [x10]
  2072: 2a5540f9                 ldr x10, [x9, #0xa8]
  2076: 4a010b8b                 add x10, x10, x11
  2080: 45e100e4                 st1b { z5.b }, p0, [x10]
  2084: 6b010d8b                 add x11, x11, x13
  2088: ce0500f1                 subs x14, x14, #1
  2092: a1bfff54                 b.ne 0x20
  2096: 7f0102eb                 cmp x11, x2
  2100: 42550054                 b.hs 0x12dc
  2104: 2a0140f9                 ldr x10, [x9]
  2108: ed030baa                 mov x13, x11
  2112: 0e0140f9                 ldr x14, [x8]
  2116: ce696d38                 ldrb w14, [x14, x13]
  2120: ec030e2a                 mov w12, w14
  2124: 0e2d40f9                 ldr x14, [x8, #0x58]
  2128: ce696d38                 ldrb w14, [x14, x13]
  2132: 8c010e4a                 eor w12, w12, w14
  2136: 0e5940f9                 ldr x14, [x8, #0xb0]
  2140: ce696d38                 ldrb w14, [x14, x13]
  2144: 8c010e4a                 eor w12, w12, w14
  2148: 0e8540f9                 ldr x14, [x8, #0x108]
  2152: ce696d38                 ldrb w14, [x14, x13]
  2156: 8c010e4a                 eor w12, w12, w14
  2160: 0eb140f9                 ldr x14, [x8, #0x160]
  2164: ce696d38                 ldrb w14, [x14, x13]
  2168: 8c010e4a                 eor w12, w12, w14
  2172: 0edd40f9                 ldr x14, [x8, #0x1b8]
  2176: ce696d38                 ldrb w14, [x14, x13]
  2180: 8c010e4a                 eor w12, w12, w14
  2184: 0e0941f9                 ldr x14, [x8, #0x210]
  2188: ce696d38                 ldrb w14, [x14, x13]
  2192: 8c010e4a                 eor w12, w12, w14
  2196: 0e3541f9                 ldr x14, [x8, #0x268]
  2200: ce696d38                 ldrb w14, [x14, x13]
  2204: 8c010e4a                 eor w12, w12, w14
  2208: 4c692d38                 strb w12, [x10, x13]
  2212: ad050091                 add x13, x13, #1
  2216: bf0102eb                 cmp x13, x2
  2220: a3fcff54                 b.lo 0x840
  2224: 2a0540f9                 ldr x10, [x9, #8]
  2228: ed030baa                 mov x13, x11
  2232: 0e0540f9                 ldr x14, [x8, #8]
  2236: ce696d38                 ldrb w14, [x14, x13]
  2240: ec030e2a                 mov w12, w14
  2244: 0e3140f9                 ldr x14, [x8, #0x60]
  2248: ce696d38                 ldrb w14, [x14, x13]
  2252: 8c010e4a                 eor w12, w12, w14
  2256: 0e5d40f9                 ldr x14, [x8, #0xb8]
  2260: ce696d38                 ldrb w14, [x14, x13]
  2264: 8c010e4a                 eor w12, w12, w14
  2268: 0e8940f9                 ldr x14, [x8, #0x110]
  2272: ce696d38                 ldrb w14, [x14, x13]
  2276: 8c010e4a                 eor w12, w12, w14
  2280: 0eb540f9                 ldr x14, [x8, #0x168]
  2284: ce696d38                 ldrb w14, [x14, x13]
  2288: 8c010e4a                 eor w12, w12, w14
  2292: 0ee140f9                 ldr x14, [x8, #0x1c0]
  2296: ce696d38                 ldrb w14, [x14, x13]
  2300: 8c010e4a                 eor w12, w12, w14
  2304: 0e0d41f9                 ldr x14, [x8, #0x218]
  2308: ce696d38                 ldrb w14, [x14, x13]
  2312: 8c010e4a                 eor w12, w12, w14
  2316: 0e3941f9                 ldr x14, [x8, #0x270]
  2320: ce696d38                 ldrb w14, [x14, x13]
  2324: 8c010e4a                 eor w12, w12, w14
  2328: 4c692d38                 strb w12, [x10, x13]
  2332: ad050091                 add x13, x13, #1
  2336: bf0102eb                 cmp x13, x2
  2340: a3fcff54                 b.lo 0x8b8
  2344: 2a0940f9                 ldr x10, [x9, #0x10]
  2348: ed030baa                 mov x13, x11
  2352: 0e0940f9                 ldr x14, [x8, #0x10]
  2356: ce696d38                 ldrb w14, [x14, x13]
  2360: ec030e2a                 mov w12, w14
  2364: 0e3540f9                 ldr x14, [x8, #0x68]
  2368: ce696d38                 ldrb w14, [x14, x13]
  2372: 8c010e4a                 eor w12, w12, w14
  2376: 0e6140f9                 ldr x14, [x8, #0xc0]
  2380: ce696d38                 ldrb w14, [x14, x13]
  2384: 8c010e4a                 eor w12, w12, w14
  2388: 0e8d40f9                 ldr x14, [x8, #0x118]
  2392: ce696d38                 ldrb w14, [x14, x13]
  2396: 8c010e4a                 eor w12, w12, w14
  2400: 0eb940f9                 ldr x14, [x8, #0x170]
  2404: ce696d38                 ldrb w14, [x14, x13]
  2408: 8c010e4a                 eor w12, w12, w14
  2412: 0ee540f9                 ldr x14, [x8, #0x1c8]
  2416: ce696d38                 ldrb w14, [x14, x13]
  2420: 8c010e4a                 eor w12, w12, w14
  2424: 0e1141f9                 ldr x14, [x8, #0x220]
  2428: ce696d38                 ldrb w14, [x14, x13]
  2432: 8c010e4a                 eor w12, w12, w14
  2436: 0e3d41f9                 ldr x14, [x8, #0x278]
  2440: ce696d38                 ldrb w14, [x14, x13]
  2444: 8c010e4a                 eor w12, w12, w14
  2448: 4c692d38                 strb w12, [x10, x13]
  2452: ad050091                 add x13, x13, #1
  2456: bf0102eb                 cmp x13, x2
  2460: a3fcff54                 b.lo 0x930
  2464: 2a0d40f9                 ldr x10, [x9, #0x18]
  2468: ed030baa                 mov x13, x11
  2472: 0e0d40f9                 ldr x14, [x8, #0x18]
  2476: ce696d38                 ldrb w14, [x14, x13]
  2480: ec030e2a                 mov w12, w14
  2484: 0e3940f9                 ldr x14, [x8, #0x70]
  2488: ce696d38                 ldrb w14, [x14, x13]
  2492: 8c010e4a                 eor w12, w12, w14
  2496: 0e6540f9                 ldr x14, [x8, #0xc8]
  2500: ce696d38                 ldrb w14, [x14, x13]
  2504: 8c010e4a                 eor w12, w12, w14
  2508: 0e9140f9                 ldr x14, [x8, #0x120]
  2512: ce696d38                 ldrb w14, [x14, x13]
  2516: 8c010e4a                 eor w12, w12, w14
  2520: 0ebd40f9                 ldr x14, [x8, #0x178]
  2524: ce696d38                 ldrb w14, [x14, x13]
  2528: 8c010e4a                 eor w12, w12, w14
  2532: 0ee940f9                 ldr x14, [x8, #0x1d0]
  2536: ce696d38                 ldrb w14, [x14, x13]
  2540: 8c010e4a                 eor w12, w12, w14
  2544: 0e1541f9                 ldr x14, [x8, #0x228]
  2548: ce696d38                 ldrb w14, [x14, x13]
  2552: 8c010e4a                 eor w12, w12, w14
  2556: 0e4141f9                 ldr x14, [x8, #0x280]
  2560: ce696d38                 ldrb w14, [x14, x13]
  2564: 8c010e4a                 eor w12, w12, w14
  2568: 4c692d38                 strb w12, [x10, x13]
  2572: ad050091                 add x13, x13, #1
  2576: bf0102eb                 cmp x13, x2
  2580: a3fcff54                 b.lo 0x9a8
  2584: 2a1140f9                 ldr x10, [x9, #0x20]
  2588: ed030baa                 mov x13, x11
  2592: 0e1140f9                 ldr x14, [x8, #0x20]
  2596: ce696d38                 ldrb w14, [x14, x13]
  2600: ec030e2a                 mov w12, w14
  2604: 0e3d40f9                 ldr x14, [x8, #0x78]
  2608: ce696d38                 ldrb w14, [x14, x13]
  2612: 8c010e4a                 eor w12, w12, w14
  2616: 0e6940f9                 ldr x14, [x8, #0xd0]
  2620: ce696d38                 ldrb w14, [x14, x13]
  2624: 8c010e4a                 eor w12, w12, w14
  2628: 0e9540f9                 ldr x14, [x8, #0x128]
  2632: ce696d38                 ldrb w14, [x14, x13]
  2636: 8c010e4a                 eor w12, w12, w14
  2640: 0ec140f9                 ldr x14, [x8, #0x180]
  2644: ce696d38                 ldrb w14, [x14, x13]
  2648: 8c010e4a                 eor w12, w12, w14
  2652: 0eed40f9                 ldr x14, [x8, #0x1d8]
  2656: ce696d38                 ldrb w14, [x14, x13]
  2660: 8c010e4a                 eor w12, w12, w14
  2664: 0e1941f9                 ldr x14, [x8, #0x230]
  2668: ce696d38                 ldrb w14, [x14, x13]
  2672: 8c010e4a                 eor w12, w12, w14
  2676: 0e4541f9                 ldr x14, [x8, #0x288]
  2680: ce696d38                 ldrb w14, [x14, x13]
  2684: 8c010e4a                 eor w12, w12, w14
  2688: 4c692d38                 strb w12, [x10, x13]
  2692: ad050091                 add x13, x13, #1
  2696: bf0102eb                 cmp x13, x2
  2700: a3fcff54                 b.lo 0xa20
  2704: 2a1540f9                 ldr x10, [x9, #0x28]
  2708: ed030baa                 mov x13, x11
  2712: 0e1540f9                 ldr x14, [x8, #0x28]
  2716: ce696d38                 ldrb w14, [x14, x13]
  2720: ec030e2a                 mov w12, w14
  2724: 0e4140f9                 ldr x14, [x8, #0x80]
  2728: ce696d38                 ldrb w14, [x14, x13]
  2732: 8c010e4a                 eor w12, w12, w14
  2736: 0e6d40f9                 ldr x14, [x8, #0xd8]
  2740: ce696d38                 ldrb w14, [x14, x13]
  2744: 8c010e4a                 eor w12, w12, w14
  2748: 0e9940f9                 ldr x14, [x8, #0x130]
  2752: ce696d38                 ldrb w14, [x14, x13]
  2756: 8c010e4a                 eor w12, w12, w14
  2760: 0ec540f9                 ldr x14, [x8, #0x188]
  2764: ce696d38                 ldrb w14, [x14, x13]
  2768: 8c010e4a                 eor w12, w12, w14
  2772: 0ef140f9                 ldr x14, [x8, #0x1e0]
  2776: ce696d38                 ldrb w14, [x14, x13]
  2780: 8c010e4a                 eor w12, w12, w14
  2784: 0e1d41f9                 ldr x14, [x8, #0x238]
  2788: ce696d38                 ldrb w14, [x14, x13]
  2792: 8c010e4a                 eor w12, w12, w14
  2796: 0e4941f9                 ldr x14, [x8, #0x290]
  2800: ce696d38                 ldrb w14, [x14, x13]
  2804: 8c010e4a                 eor w12, w12, w14
  2808: 4c692d38                 strb w12, [x10, x13]
  2812: ad050091                 add x13, x13, #1
  2816: bf0102eb                 cmp x13, x2
  2820: a3fcff54                 b.lo 0xa98
  2824: 2a1940f9                 ldr x10, [x9, #0x30]
  2828: ed030baa                 mov x13, x11
  2832: 0e1940f9                 ldr x14, [x8, #0x30]
  2836: ce696d38                 ldrb w14, [x14, x13]
  2840: ec030e2a                 mov w12, w14
  2844: 0e4540f9                 ldr x14, [x8, #0x88]
  2848: ce696d38                 ldrb w14, [x14, x13]
  2852: 8c010e4a                 eor w12, w12, w14
  2856: 0e7140f9                 ldr x14, [x8, #0xe0]
  2860: ce696d38                 ldrb w14, [x14, x13]
  2864: 8c010e4a                 eor w12, w12, w14
  2868: 0e9d40f9                 ldr x14, [x8, #0x138]
  2872: ce696d38                 ldrb w14, [x14, x13]
  2876: 8c010e4a                 eor w12, w12, w14
  2880: 0ec940f9                 ldr x14, [x8, #0x190]
  2884: ce696d38                 ldrb w14, [x14, x13]
  2888: 8c010e4a                 eor w12, w12, w14
  2892: 0ef540f9                 ldr x14, [x8, #0x1e8]
  2896: ce696d38                 ldrb w14, [x14, x13]
  2900: 8c010e4a                 eor w12, w12, w14
  2904: 0e2141f9                 ldr x14, [x8, #0x240]
  2908: ce696d38                 ldrb w14, [x14, x13]
  2912: 8c010e4a                 eor w12, w12, w14
  2916: 0e4d41f9                 ldr x14, [x8, #0x298]
  2920: ce696d38                 ldrb w14, [x14, x13]
  2924: 8c010e4a                 eor w12, w12, w14
  2928: 4c692d38                 strb w12, [x10, x13]
  2932: ad050091                 add x13, x13, #1
  2936: bf0102eb                 cmp x13, x2
  2940: a3fcff54                 b.lo 0xb10
  2944: 2a1d40f9                 ldr x10, [x9, #0x38]
  2948: ed030baa                 mov x13, x11
  2952: 0e1d40f9                 ldr x14, [x8, #0x38]
  2956: ce696d38                 ldrb w14, [x14, x13]
  2960: ec030e2a                 mov w12, w14
  2964: 0e4940f9                 ldr x14, [x8, #0x90]
  2968: ce696d38                 ldrb w14, [x14, x13]
  2972: 8c010e4a                 eor w12, w12, w14
  2976: 0e7540f9                 ldr x14, [x8, #0xe8]
  2980: ce696d38                 ldrb w14, [x14, x13]
  2984: 8c010e4a                 eor w12, w12, w14
  2988: 0ea140f9                 ldr x14, [x8, #0x140]
  2992: ce696d38                 ldrb w14, [x14, x13]
  2996: 8c010e4a                 eor w12, w12, w14
  3000: 0ecd40f9                 ldr x14, [x8, #0x198]
  3004: ce696d38                 ldrb w14, [x14, x13]
  3008: 8c010e4a                 eor w12, w12, w14
  3012: 0ef940f9                 ldr x14, [x8, #0x1f0]
  3016: ce696d38                 ldrb w14, [x14, x13]
  3020: 8c010e4a                 eor w12, w12, w14
  3024: 0e2541f9                 ldr x14, [x8, #0x248]
  3028: ce696d38                 ldrb w14, [x14, x13]
  3032: 8c010e4a                 eor w12, w12, w14
  3036: 0e5141f9                 ldr x14, [x8, #0x2a0]
  3040: ce696d38                 ldrb w14, [x14, x13]
  3044: 8c010e4a                 eor w12, w12, w14
  3048: 4c692d38                 strb w12, [x10, x13]
  3052: ad050091                 add x13, x13, #1
  3056: bf0102eb                 cmp x13, x2
  3060: a3fcff54                 b.lo 0xb88
  3064: 2a2140f9                 ldr x10, [x9, #0x40]
  3068: ed030baa                 mov x13, x11
  3072: 0e2140f9                 ldr x14, [x8, #0x40]
  3076: ce696d38                 ldrb w14, [x14, x13]
  3080: ec030e2a                 mov w12, w14
  3084: 0e4d40f9                 ldr x14, [x8, #0x98]
  3088: ce696d38                 ldrb w14, [x14, x13]
  3092: 8c010e4a                 eor w12, w12, w14
  3096: 0e7940f9                 ldr x14, [x8, #0xf0]
  3100: ce696d38                 ldrb w14, [x14, x13]
  3104: 8c010e4a                 eor w12, w12, w14
  3108: 0ea540f9                 ldr x14, [x8, #0x148]
  3112: ce696d38                 ldrb w14, [x14, x13]
  3116: 8c010e4a                 eor w12, w12, w14
  3120: 0ed140f9                 ldr x14, [x8, #0x1a0]
  3124: ce696d38                 ldrb w14, [x14, x13]
  3128: 8c010e4a                 eor w12, w12, w14
  3132: 0efd40f9                 ldr x14, [x8, #0x1f8]
  3136: ce696d38                 ldrb w14, [x14, x13]
  3140: 8c010e4a                 eor w12, w12, w14
  3144: 0e2941f9                 ldr x14, [x8, #0x250]
  3148: ce696d38                 ldrb w14, [x14, x13]
  3152: 8c010e4a                 eor w12, w12, w14
  3156: 0e5541f9                 ldr x14, [x8, #0x2a8]
  3160: ce696d38                 ldrb w14, [x14, x13]
  3164: 8c010e4a                 eor w12, w12, w14
  3168: 4c692d38                 strb w12, [x10, x13]
  3172: ad050091                 add x13, x13, #1
  3176: bf0102eb                 cmp x13, x2
  3180: a3fcff54                 b.lo 0xc00
  3184: 2a2540f9                 ldr x10, [x9, #0x48]
  3188: ed030baa                 mov x13, x11
  3192: 0e2540f9                 ldr x14, [x8, #0x48]
  3196: ce696d38                 ldrb w14, [x14, x13]
  3200: ec030e2a                 mov w12, w14
  3204: 0e5140f9                 ldr x14, [x8, #0xa0]
  3208: ce696d38                 ldrb w14, [x14, x13]
  3212: 8c010e4a                 eor w12, w12, w14
  3216: 0e7d40f9                 ldr x14, [x8, #0xf8]
  3220: ce696d38                 ldrb w14, [x14, x13]
  3224: 8c010e4a                 eor w12, w12, w14
  3228: 0ea940f9                 ldr x14, [x8, #0x150]
  3232: ce696d38                 ldrb w14, [x14, x13]
  3236: 8c010e4a                 eor w12, w12, w14
  3240: 0ed540f9                 ldr x14, [x8, #0x1a8]
  3244: ce696d38                 ldrb w14, [x14, x13]
  3248: 8c010e4a                 eor w12, w12, w14
  3252: 0e0141f9                 ldr x14, [x8, #0x200]
  3256: ce696d38                 ldrb w14, [x14, x13]
  3260: 8c010e4a                 eor w12, w12, w14
  3264: 0e2d41f9                 ldr x14, [x8, #0x258]
  3268: ce696d38                 ldrb w14, [x14, x13]
  3272: 8c010e4a                 eor w12, w12, w14
  3276: 0e5941f9                 ldr x14, [x8, #0x2b0]
  3280: ce696d38                 ldrb w14, [x14, x13]
  3284: 8c010e4a                 eor w12, w12, w14
  3288: 4c692d38                 strb w12, [x10, x13]
  3292: ad050091                 add x13, x13, #1
  3296: bf0102eb                 cmp x13, x2
  3300: a3fcff54                 b.lo 0xc78
  3304: 2a2940f9                 ldr x10, [x9, #0x50]
  3308: ed030baa                 mov x13, x11
  3312: 0e2940f9                 ldr x14, [x8, #0x50]
  3316: ce696d38                 ldrb w14, [x14, x13]
  3320: ec030e2a                 mov w12, w14
  3324: 0e5540f9                 ldr x14, [x8, #0xa8]
  3328: ce696d38                 ldrb w14, [x14, x13]
  3332: 8c010e4a                 eor w12, w12, w14
  3336: 0e8140f9                 ldr x14, [x8, #0x100]
  3340: ce696d38                 ldrb w14, [x14, x13]
  3344: 8c010e4a                 eor w12, w12, w14
  3348: 0ead40f9                 ldr x14, [x8, #0x158]
  3352: ce696d38                 ldrb w14, [x14, x13]
  3356: 8c010e4a                 eor w12, w12, w14
  3360: 0ed940f9                 ldr x14, [x8, #0x1b0]
  3364: ce696d38                 ldrb w14, [x14, x13]
  3368: 8c010e4a                 eor w12, w12, w14
  3372: 0e0541f9                 ldr x14, [x8, #0x208]
  3376: ce696d38                 ldrb w14, [x14, x13]
  3380: 8c010e4a                 eor w12, w12, w14
  3384: 0e3141f9                 ldr x14, [x8, #0x260]
  3388: ce696d38                 ldrb w14, [x14, x13]
  3392: 8c010e4a                 eor w12, w12, w14
  3396: 0e5d41f9                 ldr x14, [x8, #0x2b8]
  3400: ce696d38                 ldrb w14, [x14, x13]
  3404: 8c010e4a                 eor w12, w12, w14
  3408: 4c692d38                 strb w12, [x10, x13]
  3412: ad050091                 add x13, x13, #1
  3416: bf0102eb                 cmp x13, x2
  3420: a3fcff54                 b.lo 0xcf0
  3424: 2a2d40f9                 ldr x10, [x9, #0x58]
  3428: ed030baa                 mov x13, x11
  3432: 0e0140f9                 ldr x14, [x8]
  3436: ce696d38                 ldrb w14, [x14, x13]
  3440: ec030e2a                 mov w12, w14
  3444: 0e3140f9                 ldr x14, [x8, #0x60]
  3448: ce696d38                 ldrb w14, [x14, x13]
  3452: 8c010e4a                 eor w12, w12, w14
  3456: 0e6140f9                 ldr x14, [x8, #0xc0]
  3460: ce696d38                 ldrb w14, [x14, x13]
  3464: 8c010e4a                 eor w12, w12, w14
  3468: 0e9140f9                 ldr x14, [x8, #0x120]
  3472: ce696d38                 ldrb w14, [x14, x13]
  3476: 8c010e4a                 eor w12, w12, w14
  3480: 0ec140f9                 ldr x14, [x8, #0x180]
  3484: ce696d38                 ldrb w14, [x14, x13]
  3488: 8c010e4a                 eor w12, w12, w14
  3492: 0ef140f9                 ldr x14, [x8, #0x1e0]
  3496: ce696d38                 ldrb w14, [x14, x13]
  3500: 8c010e4a                 eor w12, w12, w14
  3504: 0e2141f9                 ldr x14, [x8, #0x240]
  3508: ce696d38                 ldrb w14, [x14, x13]
  3512: 8c010e4a                 eor w12, w12, w14
  3516: 0e5141f9                 ldr x14, [x8, #0x2a0]
  3520: ce696d38                 ldrb w14, [x14, x13]
  3524: 8c010e4a                 eor w12, w12, w14
  3528: 4c692d38                 strb w12, [x10, x13]
  3532: ad050091                 add x13, x13, #1
  3536: bf0102eb                 cmp x13, x2
  3540: a3fcff54                 b.lo 0xd68
  3544: 2a3140f9                 ldr x10, [x9, #0x60]
  3548: ed030baa                 mov x13, x11
  3552: 0e0540f9                 ldr x14, [x8, #8]
  3556: ce696d38                 ldrb w14, [x14, x13]
  3560: ec030e2a                 mov w12, w14
  3564: 0e3540f9                 ldr x14, [x8, #0x68]
  3568: ce696d38                 ldrb w14, [x14, x13]
  3572: 8c010e4a                 eor w12, w12, w14
  3576: 0e6540f9                 ldr x14, [x8, #0xc8]
  3580: ce696d38                 ldrb w14, [x14, x13]
  3584: 8c010e4a                 eor w12, w12, w14
  3588: 0e9540f9                 ldr x14, [x8, #0x128]
  3592: ce696d38                 ldrb w14, [x14, x13]
  3596: 8c010e4a                 eor w12, w12, w14
  3600: 0ec540f9                 ldr x14, [x8, #0x188]
  3604: ce696d38                 ldrb w14, [x14, x13]
  3608: 8c010e4a                 eor w12, w12, w14
  3612: 0ef540f9                 ldr x14, [x8, #0x1e8]
  3616: ce696d38                 ldrb w14, [x14, x13]
  3620: 8c010e4a                 eor w12, w12, w14
  3624: 0e2541f9                 ldr x14, [x8, #0x248]
  3628: ce696d38                 ldrb w14, [x14, x13]
  3632: 8c010e4a                 eor w12, w12, w14
  3636: 0e5541f9                 ldr x14, [x8, #0x2a8]
  3640: ce696d38                 ldrb w14, [x14, x13]
  3644: 8c010e4a                 eor w12, w12, w14
  3648: 4c692d38                 strb w12, [x10, x13]
  3652: ad050091                 add x13, x13, #1
  3656: bf0102eb                 cmp x13, x2
  3660: a3fcff54                 b.lo 0xde0
  3664: 2a3540f9                 ldr x10, [x9, #0x68]
  3668: ed030baa                 mov x13, x11
  3672: 0e0940f9                 ldr x14, [x8, #0x10]
  3676: ce696d38                 ldrb w14, [x14, x13]
  3680: ec030e2a                 mov w12, w14
  3684: 0e3940f9                 ldr x14, [x8, #0x70]
  3688: ce696d38                 ldrb w14, [x14, x13]
  3692: 8c010e4a                 eor w12, w12, w14
  3696: 0e6940f9                 ldr x14, [x8, #0xd0]
  3700: ce696d38                 ldrb w14, [x14, x13]
  3704: 8c010e4a                 eor w12, w12, w14
  3708: 0e9940f9                 ldr x14, [x8, #0x130]
  3712: ce696d38                 ldrb w14, [x14, x13]
  3716: 8c010e4a                 eor w12, w12, w14
  3720: 0ec940f9                 ldr x14, [x8, #0x190]
  3724: ce696d38                 ldrb w14, [x14, x13]
  3728: 8c010e4a                 eor w12, w12, w14
  3732: 0ef940f9                 ldr x14, [x8, #0x1f0]
  3736: ce696d38                 ldrb w14, [x14, x13]
  3740: 8c010e4a                 eor w12, w12, w14
  3744: 0e2941f9                 ldr x14, [x8, #0x250]
  3748: ce696d38                 ldrb w14, [x14, x13]
  3752: 8c010e4a                 eor w12, w12, w14
  3756: 0e5541f9                 ldr x14, [x8, #0x2a8]
  3760: ce696d38                 ldrb w14, [x14, x13]
  3764: 8c010e4a                 eor w12, w12, w14
  3768: 0e5941f9                 ldr x14, [x8, #0x2b0]
  3772: ce696d38                 ldrb w14, [x14, x13]
  3776: 8c010e4a                 eor w12, w12, w14
  3780: 4c692d38                 strb w12, [x10, x13]
  3784: ad050091                 add x13, x13, #1
  3788: bf0102eb                 cmp x13, x2
  3792: 43fcff54                 b.lo 0xe58
  3796: 2a3940f9                 ldr x10, [x9, #0x70]
  3800: ed030baa                 mov x13, x11
  3804: 0e0d40f9                 ldr x14, [x8, #0x18]
  3808: ce696d38                 ldrb w14, [x14, x13]
  3812: ec030e2a                 mov w12, w14
  3816: 0e3d40f9                 ldr x14, [x8, #0x78]
  3820: ce696d38                 ldrb w14, [x14, x13]
  3824: 8c010e4a                 eor w12, w12, w14
  3828: 0e6d40f9                 ldr x14, [x8, #0xd8]
  3832: ce696d38                 ldrb w14, [x14, x13]
  3836: 8c010e4a                 eor w12, w12, w14
  3840: 0e9d40f9                 ldr x14, [x8, #0x138]
  3844: ce696d38                 ldrb w14, [x14, x13]
  3848: 8c010e4a                 eor w12, w12, w14
  3852: 0ecd40f9                 ldr x14, [x8, #0x198]
  3856: ce696d38                 ldrb w14, [x14, x13]
  3860: 8c010e4a                 eor w12, w12, w14
  3864: 0ef940f9                 ldr x14, [x8, #0x1f0]
  3868: ce696d38                 ldrb w14, [x14, x13]
  3872: 8c010e4a                 eor w12, w12, w14
  3876: 0efd40f9                 ldr x14, [x8, #0x1f8]
  3880: ce696d38                 ldrb w14, [x14, x13]
  3884: 8c010e4a                 eor w12, w12, w14
  3888: 0e2d41f9                 ldr x14, [x8, #0x258]
  3892: ce696d38                 ldrb w14, [x14, x13]
  3896: 8c010e4a                 eor w12, w12, w14
  3900: 0e5d41f9                 ldr x14, [x8, #0x2b8]
  3904: ce696d38                 ldrb w14, [x14, x13]
  3908: 8c010e4a                 eor w12, w12, w14
  3912: 4c692d38                 strb w12, [x10, x13]
  3916: ad050091                 add x13, x13, #1
  3920: bf0102eb                 cmp x13, x2
  3924: 43fcff54                 b.lo 0xedc
  3928: 2a3d40f9                 ldr x10, [x9, #0x78]
  3932: ed030baa                 mov x13, x11
  3936: 0e1140f9                 ldr x14, [x8, #0x20]
  3940: ce696d38                 ldrb w14, [x14, x13]
  3944: ec030e2a                 mov w12, w14
  3948: 0e4140f9                 ldr x14, [x8, #0x80]
  3952: ce696d38                 ldrb w14, [x14, x13]
  3956: 8c010e4a                 eor w12, w12, w14
  3960: 0e7140f9                 ldr x14, [x8, #0xe0]
  3964: ce696d38                 ldrb w14, [x14, x13]
  3968: 8c010e4a                 eor w12, w12, w14
  3972: 0e9d40f9                 ldr x14, [x8, #0x138]
  3976: ce696d38                 ldrb w14, [x14, x13]
  3980: 8c010e4a                 eor w12, w12, w14
  3984: 0ea140f9                 ldr x14, [x8, #0x140]
  3988: ce696d38                 ldrb w14, [x14, x13]
  3992: 8c010e4a                 eor w12, w12, w14
  3996: 0ed140f9                 ldr x14, [x8, #0x1a0]
  4000: ce696d38                 ldrb w14, [x14, x13]
  4004: 8c010e4a                 eor w12, w12, w14
  4008: 0e0141f9                 ldr x14, [x8, #0x200]
  4012: ce696d38                 ldrb w14, [x14, x13]
  4016: 8c010e4a                 eor w12, w12, w14
  4020: 0e3141f9                 ldr x14, [x8, #0x260]
  4024: ce696d38                 ldrb w14, [x14, x13]
  4028: 8c010e4a                 eor w12, w12, w14
  4032: 0e3541f9                 ldr x14, [x8, #0x268]
  4036: ce696d38                 ldrb w14, [x14, x13]
  4040: 8c010e4a                 eor w12, w12, w14
  4044: 4c692d38                 strb w12, [x10, x13]
  4048: ad050091                 add x13, x13, #1
  4052: bf0102eb                 cmp x13, x2
  4056: 43fcff54                 b.lo 0xf60
  4060: 2a4140f9                 ldr x10, [x9, #0x80]
  4064: ed030baa                 mov x13, x11
  4068: 0e1540f9                 ldr x14, [x8, #0x28]
  4072: ce696d38                 ldrb w14, [x14, x13]
  4076: ec030e2a                 mov w12, w14
  4080: 0e4140f9                 ldr x14, [x8, #0x80]
  4084: ce696d38                 ldrb w14, [x14, x13]
  4088: 8c010e4a                 eor w12, w12, w14
  4092: 0e4540f9                 ldr x14, [x8, #0x88]
  4096: ce696d38                 ldrb w14, [x14, x13]
  4100: 8c010e4a                 eor w12, w12, w14
  4104: 0e7540f9                 ldr x14, [x8, #0xe8]
  4108: ce696d38                 ldrb w14, [x14, x13]
  4112: 8c010e4a                 eor w12, w12, w14
  4116: 0ea540f9                 ldr x14, [x8, #0x148]
  4120: ce696d38                 ldrb w14, [x14, x13]
  4124: 8c010e4a                 eor w12, w12, w14
  4128: 0ed540f9                 ldr x14, [x8, #0x1a8]
  4132: ce696d38                 ldrb w14, [x14, x13]
  4136: 8c010e4a                 eor w12, w12, w14
  4140: 0e0541f9                 ldr x14, [x8, #0x208]
  4144: ce696d38                 ldrb w14, [x14, x13]
  4148: 8c010e4a                 eor w12, w12, w14
  4152: 0e0941f9                 ldr x14, [x8, #0x210]
  4156: ce696d38                 ldrb w14, [x14, x13]
  4160: 8c010e4a                 eor w12, w12, w14
  4164: 0e3941f9                 ldr x14, [x8, #0x270]
  4168: ce696d38                 ldrb w14, [x14, x13]
  4172: 8c010e4a                 eor w12, w12, w14
  4176: 4c692d38                 strb w12, [x10, x13]
  4180: ad050091                 add x13, x13, #1
  4184: bf0102eb                 cmp x13, x2
  4188: 43fcff54                 b.lo 0xfe4
  4192: 2a4540f9                 ldr x10, [x9, #0x88]
  4196: ed030baa                 mov x13, x11
  4200: 0e1940f9                 ldr x14, [x8, #0x30]
  4204: ce696d38                 ldrb w14, [x14, x13]
  4208: ec030e2a                 mov w12, w14
  4212: 0e4940f9                 ldr x14, [x8, #0x90]
  4216: ce696d38                 ldrb w14, [x14, x13]
  4220: 8c010e4a                 eor w12, w12, w14
  4224: 0e7940f9                 ldr x14, [x8, #0xf0]
  4228: ce696d38                 ldrb w14, [x14, x13]
  4232: 8c010e4a                 eor w12, w12, w14
  4236: 0ea940f9                 ldr x14, [x8, #0x150]
  4240: ce696d38                 ldrb w14, [x14, x13]
  4244: 8c010e4a                 eor w12, w12, w14
  4248: 0ed940f9                 ldr x14, [x8, #0x1b0]
  4252: ce696d38                 ldrb w14, [x14, x13]
  4256: 8c010e4a                 eor w12, w12, w14
  4260: 0edd40f9                 ldr x14, [x8, #0x1b8]
  4264: ce696d38                 ldrb w14, [x14, x13]
  4268: 8c010e4a                 eor w12, w12, w14
  4272: 0e0d41f9                 ldr x14, [x8, #0x218]
  4276: ce696d38                 ldrb w14, [x14, x13]
  4280: 8c010e4a                 eor w12, w12, w14
  4284: 0e3d41f9                 ldr x14, [x8, #0x278]
  4288: ce696d38                 ldrb w14, [x14, x13]
  4292: 8c010e4a                 eor w12, w12, w14
  4296: 4c692d38                 strb w12, [x10, x13]
  4300: ad050091                 add x13, x13, #1
  4304: bf0102eb                 cmp x13, x2
  4308: a3fcff54                 b.lo 0x1068
  4312: 2a4940f9                 ldr x10, [x9, #0x90]
  4316: ed030baa                 mov x13, x11
  4320: 0e1d40f9                 ldr x14, [x8, #0x38]
  4324: ce696d38                 ldrb w14, [x14, x13]
  4328: ec030e2a                 mov w12, w14
  4332: 0e4d40f9                 ldr x14, [x8, #0x98]
  4336: ce696d38                 ldrb w14, [x14, x13]
  4340: 8c010e4a                 eor w12, w12, w14
  4344: 0e7d40f9                 ldr x14, [x8, #0xf8]
  4348: ce696d38                 ldrb w14, [x14, x13]
  4352: 8c010e4a                 eor w12, w12, w14
  4356: 0ead40f9                 ldr x14, [x8, #0x158]
  4360: ce696d38                 ldrb w14, [x14, x13]
  4364: 8c010e4a                 eor w12, w12, w14
  4368: 0eb140f9                 ldr x14, [x8, #0x160]
  4372: ce696d38                 ldrb w14, [x14, x13]
  4376: 8c010e4a                 eor w12, w12, w14
  4380: 0ee140f9                 ldr x14, [x8, #0x1c0]
  4384: ce696d38                 ldrb w14, [x14, x13]
  4388: 8c010e4a                 eor w12, w12, w14
  4392: 0e1141f9                 ldr x14, [x8, #0x220]
  4396: ce696d38                 ldrb w14, [x14, x13]
  4400: 8c010e4a                 eor w12, w12, w14
  4404: 0e4141f9                 ldr x14, [x8, #0x280]
  4408: ce696d38                 ldrb w14, [x14, x13]
  4412: 8c010e4a                 eor w12, w12, w14
  4416: 4c692d38                 strb w12, [x10, x13]
  4420: ad050091                 add x13, x13, #1
  4424: bf0102eb                 cmp x13, x2
  4428: a3fcff54                 b.lo 0x10e0
  4432: 2a4d40f9                 ldr x10, [x9, #0x98]
  4436: ed030baa                 mov x13, x11
  4440: 0e2140f9                 ldr x14, [x8, #0x40]
  4444: ce696d38                 ldrb w14, [x14, x13]
  4448: ec030e2a                 mov w12, w14
  4452: 0e5140f9                 ldr x14, [x8, #0xa0]
  4456: ce696d38                 ldrb w14, [x14, x13]
  4460: 8c010e4a                 eor w12, w12, w14
  4464: 0e8140f9                 ldr x14, [x8, #0x100]
  4468: ce696d38                 ldrb w14, [x14, x13]
  4472: 8c010e4a                 eor w12, w12, w14
  4476: 0e8540f9                 ldr x14, [x8, #0x108]
  4480: ce696d38                 ldrb w14, [x14, x13]
  4484: 8c010e4a                 eor w12, w12, w14
  4488: 0eb540f9                 ldr x14, [x8, #0x168]
  4492: ce696d38                 ldrb w14, [x14, x13]
  4496: 8c010e4a                 eor w12, w12, w14
  4500: 0ee540f9                 ldr x14, [x8, #0x1c8]
  4504: ce696d38                 ldrb w14, [x14, x13]
  4508: 8c010e4a                 eor w12, w12, w14
  4512: 0e1141f9                 ldr x14, [x8, #0x220]
  4516: ce696d38                 ldrb w14, [x14, x13]
  4520: 8c010e4a                 eor w12, w12, w14
  4524: 0e1541f9                 ldr x14, [x8, #0x228]
  4528: ce696d38                 ldrb w14, [x14, x13]
  4532: 8c010e4a                 eor w12, w12, w14
  4536: 0e4541f9                 ldr x14, [x8, #0x288]
  4540: ce696d38                 ldrb w14, [x14, x13]
  4544: 8c010e4a                 eor w12, w12, w14
  4548: 4c692d38                 strb w12, [x10, x13]
  4552: ad050091                 add x13, x13, #1
  4556: bf0102eb                 cmp x13, x2
  4560: 43fcff54                 b.lo 0x1158
  4564: 2a5140f9                 ldr x10, [x9, #0xa0]
  4568: ed030baa                 mov x13, x11
  4572: 0e2540f9                 ldr x14, [x8, #0x48]
  4576: ce696d38                 ldrb w14, [x14, x13]
  4580: ec030e2a                 mov w12, w14
  4584: 0e5540f9                 ldr x14, [x8, #0xa8]
  4588: ce696d38                 ldrb w14, [x14, x13]
  4592: 8c010e4a                 eor w12, w12, w14
  4596: 0e5940f9                 ldr x14, [x8, #0xb0]
  4600: ce696d38                 ldrb w14, [x14, x13]
  4604: 8c010e4a                 eor w12, w12, w14
  4608: 0e8940f9                 ldr x14, [x8, #0x110]
  4612: ce696d38                 ldrb w14, [x14, x13]
  4616: 8c010e4a                 eor w12, w12, w14
  4620: 0eb540f9                 ldr x14, [x8, #0x168]
  4624: ce696d38                 ldrb w14, [x14, x13]
  4628: 8c010e4a                 eor w12, w12, w14
  4632: 0eb940f9                 ldr x14, [x8, #0x170]
  4636: ce696d38                 ldrb w14, [x14, x13]
  4640: 8c010e4a                 eor w12, w12, w14
  4644: 0ee940f9                 ldr x14, [x8, #0x1d0]
  4648: ce696d38                 ldrb w14, [x14, x13]
  4652: 8c010e4a                 eor w12, w12, w14
  4656: 0e1941f9                 ldr x14, [x8, #0x230]
  4660: ce696d38                 ldrb w14, [x14, x13]
  4664: 8c010e4a                 eor w12, w12, w14
  4668: 0e4941f9                 ldr x14, [x8, #0x290]
  4672: ce696d38                 ldrb w14, [x14, x13]
  4676: 8c010e4a                 eor w12, w12, w14
  4680: 4c692d38                 strb w12, [x10, x13]
  4684: ad050091                 add x13, x13, #1
  4688: bf0102eb                 cmp x13, x2
  4692: 43fcff54                 b.lo 0x11dc
  4696: 2a5540f9                 ldr x10, [x9, #0xa8]
  4700: ed030baa                 mov x13, x11
  4704: 0e2940f9                 ldr x14, [x8, #0x50]
  4708: ce696d38                 ldrb w14, [x14, x13]
  4712: ec030e2a                 mov w12, w14
  4716: 0e2d40f9                 ldr x14, [x8, #0x58]
  4720: ce696d38                 ldrb w14, [x14, x13]
  4724: 8c010e4a                 eor w12, w12, w14
  4728: 0e5940f9                 ldr x14, [x8, #0xb0]
  4732: ce696d38                 ldrb w14, [x14, x13]
  4736: 8c010e4a                 eor w12, w12, w14
  4740: 0e5d40f9                 ldr x14, [x8, #0xb8]
  4744: ce696d38                 ldrb w14, [x14, x13]
  4748: 8c010e4a                 eor w12, w12, w14
  4752: 0e8d40f9                 ldr x14, [x8, #0x118]
  4756: ce696d38                 ldrb w14, [x14, x13]
  4760: 8c010e4a                 eor w12, w12, w14
  4764: 0ebd40f9                 ldr x14, [x8, #0x178]
  4768: ce696d38                 ldrb w14, [x14, x13]
  4772: 8c010e4a                 eor w12, w12, w14
  4776: 0eed40f9                 ldr x14, [x8, #0x1d8]
  4780: ce696d38                 ldrb w14, [x14, x13]
  4784: 8c010e4a                 eor w12, w12, w14
  4788: 0e1d41f9                 ldr x14, [x8, #0x238]
  4792: ce696d38                 ldrb w14, [x14, x13]
  4796: 8c010e4a                 eor w12, w12, w14
  4800: 0e4d41f9                 ldr x14, [x8, #0x298]
  4804: ce696d38                 ldrb w14, [x14, x13]
  4808: 8c010e4a                 eor w12, w12, w14
  4812: 4c692d38                 strb w12, [x10, x13]
  4816: ad050091                 add x13, x13, #1
  4820: bf0102eb                 cmp x13, x2
  4824: 43fcff54                 b.lo 0x1260
  4828: c0035fd6                 ret 
  4832: 00000000                 udf #0
  4836: 00000000                 udf #0
  4840: 00000000                 udf #0
  4844: 00000000                 udf #0
  4848: 00000000                 udf #0
  4852: 00000000                 udf #0
  4856: 00000000                 udf #0
  4860: 00000000                 udf #0
