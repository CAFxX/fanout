; config_t1_911 target=sve2 twin=nt topo=rdp(7)/enc op=encode U=1
; code_size=2496 twins=["nt"] uops_per_byte=9.0625
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: ed030caa                 mov x13, x12
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: ae1e00b4                 cbz x14, 0x3f0
    32: 0a0140f9                 ldr x10, [x8]
    36: 4a010b8b                 add x10, x10, x11
    40: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
    44: 90337c04                 mov z16.d, z28.d
    48: 96337c04                 mov z22.d, z28.d
    52: 0a0540f9                 ldr x10, [x8, #8]
    56: 4a010b8b                 add x10, x10, x11
    60: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
    64: 91337c04                 mov z17.d, z28.d
    68: 96031904                 eor z22.b, p0/m, z22.b, z28.b
    72: 97337c04                 mov z23.d, z28.d
    76: 0a0940f9                 ldr x10, [x8, #0x10]
    80: 4a010b8b                 add x10, x10, x11
    84: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
    88: 92337c04                 mov z18.d, z28.d
    92: 97031904                 eor z23.b, p0/m, z23.b, z28.b
    96: 98337c04                 mov z24.d, z28.d
   100: 0a0d40f9                 ldr x10, [x8, #0x18]
   104: 4a010b8b                 add x10, x10, x11
   108: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   112: 93337c04                 mov z19.d, z28.d
   116: 98031904                 eor z24.b, p0/m, z24.b, z28.b
   120: 99337c04                 mov z25.d, z28.d
   124: 0a1140f9                 ldr x10, [x8, #0x20]
   128: 4a010b8b                 add x10, x10, x11
   132: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   136: 94337c04                 mov z20.d, z28.d
   140: 99031904                 eor z25.b, p0/m, z25.b, z28.b
   144: 9a337c04                 mov z26.d, z28.d
   148: 0a1540f9                 ldr x10, [x8, #0x28]
   152: 4a010b8b                 add x10, x10, x11
   156: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   160: 95337c04                 mov z21.d, z28.d
   164: 9a031904                 eor z26.b, p0/m, z26.b, z28.b
   168: 9b337c04                 mov z27.d, z28.d
   172: 0a1940f9                 ldr x10, [x8, #0x30]
   176: 4a010b8b                 add x10, x10, x11
   180: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   184: 90031904                 eor z16.b, p0/m, z16.b, z28.b
   188: 97031904                 eor z23.b, p0/m, z23.b, z28.b
   192: 0a1d40f9                 ldr x10, [x8, #0x38]
   196: 4a010b8b                 add x10, x10, x11
   200: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   204: 91031904                 eor z17.b, p0/m, z17.b, z28.b
   208: 96031904                 eor z22.b, p0/m, z22.b, z28.b
   212: 98031904                 eor z24.b, p0/m, z24.b, z28.b
   216: 0a2140f9                 ldr x10, [x8, #0x40]
   220: 4a010b8b                 add x10, x10, x11
   224: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   228: 92031904                 eor z18.b, p0/m, z18.b, z28.b
   232: 97031904                 eor z23.b, p0/m, z23.b, z28.b
   236: 99031904                 eor z25.b, p0/m, z25.b, z28.b
   240: 0a2540f9                 ldr x10, [x8, #0x48]
   244: 4a010b8b                 add x10, x10, x11
   248: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   252: 93031904                 eor z19.b, p0/m, z19.b, z28.b
   256: 98031904                 eor z24.b, p0/m, z24.b, z28.b
   260: 9a031904                 eor z26.b, p0/m, z26.b, z28.b
   264: 0a2940f9                 ldr x10, [x8, #0x50]
   268: 4a010b8b                 add x10, x10, x11
   272: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   276: 94031904                 eor z20.b, p0/m, z20.b, z28.b
   280: 99031904                 eor z25.b, p0/m, z25.b, z28.b
   284: 9b031904                 eor z27.b, p0/m, z27.b, z28.b
   288: 0a2d40f9                 ldr x10, [x8, #0x58]
   292: 4a010b8b                 add x10, x10, x11
   296: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   300: 95031904                 eor z21.b, p0/m, z21.b, z28.b
   304: 9a031904                 eor z26.b, p0/m, z26.b, z28.b
   308: 0a3140f9                 ldr x10, [x8, #0x60]
   312: 4a010b8b                 add x10, x10, x11
   316: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   320: 90031904                 eor z16.b, p0/m, z16.b, z28.b
   324: 98031904                 eor z24.b, p0/m, z24.b, z28.b
   328: 0a3540f9                 ldr x10, [x8, #0x68]
   332: 4a010b8b                 add x10, x10, x11
   336: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   340: 91031904                 eor z17.b, p0/m, z17.b, z28.b
   344: 96031904                 eor z22.b, p0/m, z22.b, z28.b
   348: 99031904                 eor z25.b, p0/m, z25.b, z28.b
   352: 0a3940f9                 ldr x10, [x8, #0x70]
   356: 4a010b8b                 add x10, x10, x11
   360: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   364: 92031904                 eor z18.b, p0/m, z18.b, z28.b
   368: 97031904                 eor z23.b, p0/m, z23.b, z28.b
   372: 9a031904                 eor z26.b, p0/m, z26.b, z28.b
   376: 0a3d40f9                 ldr x10, [x8, #0x78]
   380: 4a010b8b                 add x10, x10, x11
   384: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   388: 93031904                 eor z19.b, p0/m, z19.b, z28.b
   392: 98031904                 eor z24.b, p0/m, z24.b, z28.b
   396: 9b031904                 eor z27.b, p0/m, z27.b, z28.b
   400: 0a4140f9                 ldr x10, [x8, #0x80]
   404: 4a010b8b                 add x10, x10, x11
   408: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   412: 94031904                 eor z20.b, p0/m, z20.b, z28.b
   416: 99031904                 eor z25.b, p0/m, z25.b, z28.b
   420: 0a4540f9                 ldr x10, [x8, #0x88]
   424: 4a010b8b                 add x10, x10, x11
   428: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   432: 95031904                 eor z21.b, p0/m, z21.b, z28.b
   436: 96031904                 eor z22.b, p0/m, z22.b, z28.b
   440: 9a031904                 eor z26.b, p0/m, z26.b, z28.b
   444: 0a4940f9                 ldr x10, [x8, #0x90]
   448: 4a010b8b                 add x10, x10, x11
   452: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   456: 90031904                 eor z16.b, p0/m, z16.b, z28.b
   460: 99031904                 eor z25.b, p0/m, z25.b, z28.b
   464: 0a4d40f9                 ldr x10, [x8, #0x98]
   468: 4a010b8b                 add x10, x10, x11
   472: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   476: 91031904                 eor z17.b, p0/m, z17.b, z28.b
   480: 96031904                 eor z22.b, p0/m, z22.b, z28.b
   484: 9a031904                 eor z26.b, p0/m, z26.b, z28.b
   488: 0a5140f9                 ldr x10, [x8, #0xa0]
   492: 4a010b8b                 add x10, x10, x11
   496: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   500: 92031904                 eor z18.b, p0/m, z18.b, z28.b
   504: 97031904                 eor z23.b, p0/m, z23.b, z28.b
   508: 9b031904                 eor z27.b, p0/m, z27.b, z28.b
   512: 0a5540f9                 ldr x10, [x8, #0xa8]
   516: 4a010b8b                 add x10, x10, x11
   520: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   524: 93031904                 eor z19.b, p0/m, z19.b, z28.b
   528: 98031904                 eor z24.b, p0/m, z24.b, z28.b
   532: 0a5940f9                 ldr x10, [x8, #0xb0]
   536: 4a010b8b                 add x10, x10, x11
   540: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   544: 94031904                 eor z20.b, p0/m, z20.b, z28.b
   548: 96031904                 eor z22.b, p0/m, z22.b, z28.b
   552: 99031904                 eor z25.b, p0/m, z25.b, z28.b
   556: 0a5d40f9                 ldr x10, [x8, #0xb8]
   560: 4a010b8b                 add x10, x10, x11
   564: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   568: 95031904                 eor z21.b, p0/m, z21.b, z28.b
   572: 97031904                 eor z23.b, p0/m, z23.b, z28.b
   576: 9a031904                 eor z26.b, p0/m, z26.b, z28.b
   580: 0a6140f9                 ldr x10, [x8, #0xc0]
   584: 4a010b8b                 add x10, x10, x11
   588: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   592: 90031904                 eor z16.b, p0/m, z16.b, z28.b
   596: 9a031904                 eor z26.b, p0/m, z26.b, z28.b
   600: 0a6540f9                 ldr x10, [x8, #0xc8]
   604: 4a010b8b                 add x10, x10, x11
   608: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   612: 91031904                 eor z17.b, p0/m, z17.b, z28.b
   616: 96031904                 eor z22.b, p0/m, z22.b, z28.b
   620: 9b031904                 eor z27.b, p0/m, z27.b, z28.b
   624: 0a6940f9                 ldr x10, [x8, #0xd0]
   628: 4a010b8b                 add x10, x10, x11
   632: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   636: 92031904                 eor z18.b, p0/m, z18.b, z28.b
   640: 97031904                 eor z23.b, p0/m, z23.b, z28.b
   644: 0a6d40f9                 ldr x10, [x8, #0xd8]
   648: 4a010b8b                 add x10, x10, x11
   652: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   656: 93031904                 eor z19.b, p0/m, z19.b, z28.b
   660: 96031904                 eor z22.b, p0/m, z22.b, z28.b
   664: 98031904                 eor z24.b, p0/m, z24.b, z28.b
   668: 0a7140f9                 ldr x10, [x8, #0xe0]
   672: 4a010b8b                 add x10, x10, x11
   676: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   680: 94031904                 eor z20.b, p0/m, z20.b, z28.b
   684: 97031904                 eor z23.b, p0/m, z23.b, z28.b
   688: 99031904                 eor z25.b, p0/m, z25.b, z28.b
   692: 0a7540f9                 ldr x10, [x8, #0xe8]
   696: 4a010b8b                 add x10, x10, x11
   700: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   704: 95031904                 eor z21.b, p0/m, z21.b, z28.b
   708: 98031904                 eor z24.b, p0/m, z24.b, z28.b
   712: 9a031904                 eor z26.b, p0/m, z26.b, z28.b
   716: 0a7940f9                 ldr x10, [x8, #0xf0]
   720: 4a010b8b                 add x10, x10, x11
   724: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   728: 90031904                 eor z16.b, p0/m, z16.b, z28.b
   732: 9b031904                 eor z27.b, p0/m, z27.b, z28.b
   736: 0a7d40f9                 ldr x10, [x8, #0xf8]
   740: 4a010b8b                 add x10, x10, x11
   744: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   748: 91031904                 eor z17.b, p0/m, z17.b, z28.b
   752: 96031904                 eor z22.b, p0/m, z22.b, z28.b
   756: 0a8140f9                 ldr x10, [x8, #0x100]
   760: 4a010b8b                 add x10, x10, x11
   764: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   768: 92031904                 eor z18.b, p0/m, z18.b, z28.b
   772: 96031904                 eor z22.b, p0/m, z22.b, z28.b
   776: 97031904                 eor z23.b, p0/m, z23.b, z28.b
   780: 0a8540f9                 ldr x10, [x8, #0x108]
   784: 4a010b8b                 add x10, x10, x11
   788: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   792: 93031904                 eor z19.b, p0/m, z19.b, z28.b
   796: 97031904                 eor z23.b, p0/m, z23.b, z28.b
   800: 98031904                 eor z24.b, p0/m, z24.b, z28.b
   804: 0a8940f9                 ldr x10, [x8, #0x110]
   808: 4a010b8b                 add x10, x10, x11
   812: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   816: 94031904                 eor z20.b, p0/m, z20.b, z28.b
   820: 98031904                 eor z24.b, p0/m, z24.b, z28.b
   824: 99031904                 eor z25.b, p0/m, z25.b, z28.b
   828: 0a8d40f9                 ldr x10, [x8, #0x118]
   832: 4a010b8b                 add x10, x10, x11
   836: 5ca100a4                 ld1b { z28.b }, p0/z, [x10]
   840: 95031904                 eor z21.b, p0/m, z21.b, z28.b
   844: 99031904                 eor z25.b, p0/m, z25.b, z28.b
   848: 9a031904                 eor z26.b, p0/m, z26.b, z28.b
   852: 2a0140f9                 ldr x10, [x9]
   856: 4a010b8b                 add x10, x10, x11
   860: 50e110e4                 stnt1b { z16.b }, p0, [x10]
   864: 2a0540f9                 ldr x10, [x9, #8]
   868: 4a010b8b                 add x10, x10, x11
   872: 51e110e4                 stnt1b { z17.b }, p0, [x10]
   876: 2a0940f9                 ldr x10, [x9, #0x10]
   880: 4a010b8b                 add x10, x10, x11
   884: 52e110e4                 stnt1b { z18.b }, p0, [x10]
   888: 2a0d40f9                 ldr x10, [x9, #0x18]
   892: 4a010b8b                 add x10, x10, x11
   896: 53e110e4                 stnt1b { z19.b }, p0, [x10]
   900: 2a1140f9                 ldr x10, [x9, #0x20]
   904: 4a010b8b                 add x10, x10, x11
   908: 54e110e4                 stnt1b { z20.b }, p0, [x10]
   912: 2a1540f9                 ldr x10, [x9, #0x28]
   916: 4a010b8b                 add x10, x10, x11
   920: 55e110e4                 stnt1b { z21.b }, p0, [x10]
   924: 2a1940f9                 ldr x10, [x9, #0x30]
   928: 4a010b8b                 add x10, x10, x11
   932: 56e110e4                 stnt1b { z22.b }, p0, [x10]
   936: 2a1d40f9                 ldr x10, [x9, #0x38]
   940: 4a010b8b                 add x10, x10, x11
   944: 57e110e4                 stnt1b { z23.b }, p0, [x10]
   948: 2a2140f9                 ldr x10, [x9, #0x40]
   952: 4a010b8b                 add x10, x10, x11
   956: 58e110e4                 stnt1b { z24.b }, p0, [x10]
   960: 2a2540f9                 ldr x10, [x9, #0x48]
   964: 4a010b8b                 add x10, x10, x11
   968: 59e110e4                 stnt1b { z25.b }, p0, [x10]
   972: 2a2940f9                 ldr x10, [x9, #0x50]
   976: 4a010b8b                 add x10, x10, x11
   980: 5ae110e4                 stnt1b { z26.b }, p0, [x10]
   984: 2a2d40f9                 ldr x10, [x9, #0x58]
   988: 4a010b8b                 add x10, x10, x11
   992: 5be110e4                 stnt1b { z27.b }, p0, [x10]
   996: 6b010d8b                 add x11, x11, x13
  1000: ce0500f1                 subs x14, x14, #1
  1004: a1e1ff54                 b.ne 0x20
  1008: 7f0102eb                 cmp x11, x2
  1012: 822d0054                 b.hs 0x9a4
  1016: 2a0140f9                 ldr x10, [x9]
  1020: ed030baa                 mov x13, x11
  1024: 0e0140f9                 ldr x14, [x8]
  1028: ce696d38                 ldrb w14, [x14, x13]
  1032: ec030e2a                 mov w12, w14
  1036: 0e1940f9                 ldr x14, [x8, #0x30]
  1040: ce696d38                 ldrb w14, [x14, x13]
  1044: 8c010e4a                 eor w12, w12, w14
  1048: 0e3140f9                 ldr x14, [x8, #0x60]
  1052: ce696d38                 ldrb w14, [x14, x13]
  1056: 8c010e4a                 eor w12, w12, w14
  1060: 0e4940f9                 ldr x14, [x8, #0x90]
  1064: ce696d38                 ldrb w14, [x14, x13]
  1068: 8c010e4a                 eor w12, w12, w14
  1072: 0e6140f9                 ldr x14, [x8, #0xc0]
  1076: ce696d38                 ldrb w14, [x14, x13]
  1080: 8c010e4a                 eor w12, w12, w14
  1084: 0e7940f9                 ldr x14, [x8, #0xf0]
  1088: ce696d38                 ldrb w14, [x14, x13]
  1092: 8c010e4a                 eor w12, w12, w14
  1096: 4c692d38                 strb w12, [x10, x13]
  1100: ad050091                 add x13, x13, #1
  1104: bf0102eb                 cmp x13, x2
  1108: 63fdff54                 b.lo 0x400
  1112: 2a0540f9                 ldr x10, [x9, #8]
  1116: ed030baa                 mov x13, x11
  1120: 0e0540f9                 ldr x14, [x8, #8]
  1124: ce696d38                 ldrb w14, [x14, x13]
  1128: ec030e2a                 mov w12, w14
  1132: 0e1d40f9                 ldr x14, [x8, #0x38]
  1136: ce696d38                 ldrb w14, [x14, x13]
  1140: 8c010e4a                 eor w12, w12, w14
  1144: 0e3540f9                 ldr x14, [x8, #0x68]
  1148: ce696d38                 ldrb w14, [x14, x13]
  1152: 8c010e4a                 eor w12, w12, w14
  1156: 0e4d40f9                 ldr x14, [x8, #0x98]
  1160: ce696d38                 ldrb w14, [x14, x13]
  1164: 8c010e4a                 eor w12, w12, w14
  1168: 0e6540f9                 ldr x14, [x8, #0xc8]
  1172: ce696d38                 ldrb w14, [x14, x13]
  1176: 8c010e4a                 eor w12, w12, w14
  1180: 0e7d40f9                 ldr x14, [x8, #0xf8]
  1184: ce696d38                 ldrb w14, [x14, x13]
  1188: 8c010e4a                 eor w12, w12, w14
  1192: 4c692d38                 strb w12, [x10, x13]
  1196: ad050091                 add x13, x13, #1
  1200: bf0102eb                 cmp x13, x2
  1204: 63fdff54                 b.lo 0x460
  1208: 2a0940f9                 ldr x10, [x9, #0x10]
  1212: ed030baa                 mov x13, x11
  1216: 0e0940f9                 ldr x14, [x8, #0x10]
  1220: ce696d38                 ldrb w14, [x14, x13]
  1224: ec030e2a                 mov w12, w14
  1228: 0e2140f9                 ldr x14, [x8, #0x40]
  1232: ce696d38                 ldrb w14, [x14, x13]
  1236: 8c010e4a                 eor w12, w12, w14
  1240: 0e3940f9                 ldr x14, [x8, #0x70]
  1244: ce696d38                 ldrb w14, [x14, x13]
  1248: 8c010e4a                 eor w12, w12, w14
  1252: 0e5140f9                 ldr x14, [x8, #0xa0]
  1256: ce696d38                 ldrb w14, [x14, x13]
  1260: 8c010e4a                 eor w12, w12, w14
  1264: 0e6940f9                 ldr x14, [x8, #0xd0]
  1268: ce696d38                 ldrb w14, [x14, x13]
  1272: 8c010e4a                 eor w12, w12, w14
  1276: 0e8140f9                 ldr x14, [x8, #0x100]
  1280: ce696d38                 ldrb w14, [x14, x13]
  1284: 8c010e4a                 eor w12, w12, w14
  1288: 4c692d38                 strb w12, [x10, x13]
  1292: ad050091                 add x13, x13, #1
  1296: bf0102eb                 cmp x13, x2
  1300: 63fdff54                 b.lo 0x4c0
  1304: 2a0d40f9                 ldr x10, [x9, #0x18]
  1308: ed030baa                 mov x13, x11
  1312: 0e0d40f9                 ldr x14, [x8, #0x18]
  1316: ce696d38                 ldrb w14, [x14, x13]
  1320: ec030e2a                 mov w12, w14
  1324: 0e2540f9                 ldr x14, [x8, #0x48]
  1328: ce696d38                 ldrb w14, [x14, x13]
  1332: 8c010e4a                 eor w12, w12, w14
  1336: 0e3d40f9                 ldr x14, [x8, #0x78]
  1340: ce696d38                 ldrb w14, [x14, x13]
  1344: 8c010e4a                 eor w12, w12, w14
  1348: 0e5540f9                 ldr x14, [x8, #0xa8]
  1352: ce696d38                 ldrb w14, [x14, x13]
  1356: 8c010e4a                 eor w12, w12, w14
  1360: 0e6d40f9                 ldr x14, [x8, #0xd8]
  1364: ce696d38                 ldrb w14, [x14, x13]
  1368: 8c010e4a                 eor w12, w12, w14
  1372: 0e8540f9                 ldr x14, [x8, #0x108]
  1376: ce696d38                 ldrb w14, [x14, x13]
  1380: 8c010e4a                 eor w12, w12, w14
  1384: 4c692d38                 strb w12, [x10, x13]
  1388: ad050091                 add x13, x13, #1
  1392: bf0102eb                 cmp x13, x2
  1396: 63fdff54                 b.lo 0x520
  1400: 2a1140f9                 ldr x10, [x9, #0x20]
  1404: ed030baa                 mov x13, x11
  1408: 0e1140f9                 ldr x14, [x8, #0x20]
  1412: ce696d38                 ldrb w14, [x14, x13]
  1416: ec030e2a                 mov w12, w14
  1420: 0e2940f9                 ldr x14, [x8, #0x50]
  1424: ce696d38                 ldrb w14, [x14, x13]
  1428: 8c010e4a                 eor w12, w12, w14
  1432: 0e4140f9                 ldr x14, [x8, #0x80]
  1436: ce696d38                 ldrb w14, [x14, x13]
  1440: 8c010e4a                 eor w12, w12, w14
  1444: 0e5940f9                 ldr x14, [x8, #0xb0]
  1448: ce696d38                 ldrb w14, [x14, x13]
  1452: 8c010e4a                 eor w12, w12, w14
  1456: 0e7140f9                 ldr x14, [x8, #0xe0]
  1460: ce696d38                 ldrb w14, [x14, x13]
  1464: 8c010e4a                 eor w12, w12, w14
  1468: 0e8940f9                 ldr x14, [x8, #0x110]
  1472: ce696d38                 ldrb w14, [x14, x13]
  1476: 8c010e4a                 eor w12, w12, w14
  1480: 4c692d38                 strb w12, [x10, x13]
  1484: ad050091                 add x13, x13, #1
  1488: bf0102eb                 cmp x13, x2
  1492: 63fdff54                 b.lo 0x580
  1496: 2a1540f9                 ldr x10, [x9, #0x28]
  1500: ed030baa                 mov x13, x11
  1504: 0e1540f9                 ldr x14, [x8, #0x28]
  1508: ce696d38                 ldrb w14, [x14, x13]
  1512: ec030e2a                 mov w12, w14
  1516: 0e2d40f9                 ldr x14, [x8, #0x58]
  1520: ce696d38                 ldrb w14, [x14, x13]
  1524: 8c010e4a                 eor w12, w12, w14
  1528: 0e4540f9                 ldr x14, [x8, #0x88]
  1532: ce696d38                 ldrb w14, [x14, x13]
  1536: 8c010e4a                 eor w12, w12, w14
  1540: 0e5d40f9                 ldr x14, [x8, #0xb8]
  1544: ce696d38                 ldrb w14, [x14, x13]
  1548: 8c010e4a                 eor w12, w12, w14
  1552: 0e7540f9                 ldr x14, [x8, #0xe8]
  1556: ce696d38                 ldrb w14, [x14, x13]
  1560: 8c010e4a                 eor w12, w12, w14
  1564: 0e8d40f9                 ldr x14, [x8, #0x118]
  1568: ce696d38                 ldrb w14, [x14, x13]
  1572: 8c010e4a                 eor w12, w12, w14
  1576: 4c692d38                 strb w12, [x10, x13]
  1580: ad050091                 add x13, x13, #1
  1584: bf0102eb                 cmp x13, x2
  1588: 63fdff54                 b.lo 0x5e0
  1592: 2a1940f9                 ldr x10, [x9, #0x30]
  1596: ed030baa                 mov x13, x11
  1600: 0e0140f9                 ldr x14, [x8]
  1604: ce696d38                 ldrb w14, [x14, x13]
  1608: ec030e2a                 mov w12, w14
  1612: 0e0540f9                 ldr x14, [x8, #8]
  1616: ce696d38                 ldrb w14, [x14, x13]
  1620: 8c010e4a                 eor w12, w12, w14
  1624: 0e1d40f9                 ldr x14, [x8, #0x38]
  1628: ce696d38                 ldrb w14, [x14, x13]
  1632: 8c010e4a                 eor w12, w12, w14
  1636: 0e3540f9                 ldr x14, [x8, #0x68]
  1640: ce696d38                 ldrb w14, [x14, x13]
  1644: 8c010e4a                 eor w12, w12, w14
  1648: 0e4540f9                 ldr x14, [x8, #0x88]
  1652: ce696d38                 ldrb w14, [x14, x13]
  1656: 8c010e4a                 eor w12, w12, w14
  1660: 0e4d40f9                 ldr x14, [x8, #0x98]
  1664: ce696d38                 ldrb w14, [x14, x13]
  1668: 8c010e4a                 eor w12, w12, w14
  1672: 0e5940f9                 ldr x14, [x8, #0xb0]
  1676: ce696d38                 ldrb w14, [x14, x13]
  1680: 8c010e4a                 eor w12, w12, w14
  1684: 0e6540f9                 ldr x14, [x8, #0xc8]
  1688: ce696d38                 ldrb w14, [x14, x13]
  1692: 8c010e4a                 eor w12, w12, w14
  1696: 0e6d40f9                 ldr x14, [x8, #0xd8]
  1700: ce696d38                 ldrb w14, [x14, x13]
  1704: 8c010e4a                 eor w12, w12, w14
  1708: 0e7d40f9                 ldr x14, [x8, #0xf8]
  1712: ce696d38                 ldrb w14, [x14, x13]
  1716: 8c010e4a                 eor w12, w12, w14
  1720: 0e8140f9                 ldr x14, [x8, #0x100]
  1724: ce696d38                 ldrb w14, [x14, x13]
  1728: 8c010e4a                 eor w12, w12, w14
  1732: 4c692d38                 strb w12, [x10, x13]
  1736: ad050091                 add x13, x13, #1
  1740: bf0102eb                 cmp x13, x2
  1744: 83fbff54                 b.lo 0x640
  1748: 2a1d40f9                 ldr x10, [x9, #0x38]
  1752: ed030baa                 mov x13, x11
  1756: 0e0540f9                 ldr x14, [x8, #8]
  1760: ce696d38                 ldrb w14, [x14, x13]
  1764: ec030e2a                 mov w12, w14
  1768: 0e0940f9                 ldr x14, [x8, #0x10]
  1772: ce696d38                 ldrb w14, [x14, x13]
  1776: 8c010e4a                 eor w12, w12, w14
  1780: 0e1940f9                 ldr x14, [x8, #0x30]
  1784: ce696d38                 ldrb w14, [x14, x13]
  1788: 8c010e4a                 eor w12, w12, w14
  1792: 0e2140f9                 ldr x14, [x8, #0x40]
  1796: ce696d38                 ldrb w14, [x14, x13]
  1800: 8c010e4a                 eor w12, w12, w14
  1804: 0e3940f9                 ldr x14, [x8, #0x70]
  1808: ce696d38                 ldrb w14, [x14, x13]
  1812: 8c010e4a                 eor w12, w12, w14
  1816: 0e5140f9                 ldr x14, [x8, #0xa0]
  1820: ce696d38                 ldrb w14, [x14, x13]
  1824: 8c010e4a                 eor w12, w12, w14
  1828: 0e5d40f9                 ldr x14, [x8, #0xb8]
  1832: ce696d38                 ldrb w14, [x14, x13]
  1836: 8c010e4a                 eor w12, w12, w14
  1840: 0e6940f9                 ldr x14, [x8, #0xd0]
  1844: ce696d38                 ldrb w14, [x14, x13]
  1848: 8c010e4a                 eor w12, w12, w14
  1852: 0e7140f9                 ldr x14, [x8, #0xe0]
  1856: ce696d38                 ldrb w14, [x14, x13]
  1860: 8c010e4a                 eor w12, w12, w14
  1864: 0e8140f9                 ldr x14, [x8, #0x100]
  1868: ce696d38                 ldrb w14, [x14, x13]
  1872: 8c010e4a                 eor w12, w12, w14
  1876: 0e8540f9                 ldr x14, [x8, #0x108]
  1880: ce696d38                 ldrb w14, [x14, x13]
  1884: 8c010e4a                 eor w12, w12, w14
  1888: 4c692d38                 strb w12, [x10, x13]
  1892: ad050091                 add x13, x13, #1
  1896: bf0102eb                 cmp x13, x2
  1900: 83fbff54                 b.lo 0x6dc
  1904: 2a2140f9                 ldr x10, [x9, #0x40]
  1908: ed030baa                 mov x13, x11
  1912: 0e0940f9                 ldr x14, [x8, #0x10]
  1916: ce696d38                 ldrb w14, [x14, x13]
  1920: ec030e2a                 mov w12, w14
  1924: 0e0d40f9                 ldr x14, [x8, #0x18]
  1928: ce696d38                 ldrb w14, [x14, x13]
  1932: 8c010e4a                 eor w12, w12, w14
  1936: 0e1d40f9                 ldr x14, [x8, #0x38]
  1940: ce696d38                 ldrb w14, [x14, x13]
  1944: 8c010e4a                 eor w12, w12, w14
  1948: 0e2540f9                 ldr x14, [x8, #0x48]
  1952: ce696d38                 ldrb w14, [x14, x13]
  1956: 8c010e4a                 eor w12, w12, w14
  1960: 0e3140f9                 ldr x14, [x8, #0x60]
  1964: ce696d38                 ldrb w14, [x14, x13]
  1968: 8c010e4a                 eor w12, w12, w14
  1972: 0e3d40f9                 ldr x14, [x8, #0x78]
  1976: ce696d38                 ldrb w14, [x14, x13]
  1980: 8c010e4a                 eor w12, w12, w14
  1984: 0e5540f9                 ldr x14, [x8, #0xa8]
  1988: ce696d38                 ldrb w14, [x14, x13]
  1992: 8c010e4a                 eor w12, w12, w14
  1996: 0e6d40f9                 ldr x14, [x8, #0xd8]
  2000: ce696d38                 ldrb w14, [x14, x13]
  2004: 8c010e4a                 eor w12, w12, w14
  2008: 0e7540f9                 ldr x14, [x8, #0xe8]
  2012: ce696d38                 ldrb w14, [x14, x13]
  2016: 8c010e4a                 eor w12, w12, w14
  2020: 0e8540f9                 ldr x14, [x8, #0x108]
  2024: ce696d38                 ldrb w14, [x14, x13]
  2028: 8c010e4a                 eor w12, w12, w14
  2032: 0e8940f9                 ldr x14, [x8, #0x110]
  2036: ce696d38                 ldrb w14, [x14, x13]
  2040: 8c010e4a                 eor w12, w12, w14
  2044: 4c692d38                 strb w12, [x10, x13]
  2048: ad050091                 add x13, x13, #1
  2052: bf0102eb                 cmp x13, x2
  2056: 83fbff54                 b.lo 0x778
  2060: 2a2540f9                 ldr x10, [x9, #0x48]
  2064: ed030baa                 mov x13, x11
  2068: 0e0d40f9                 ldr x14, [x8, #0x18]
  2072: ce696d38                 ldrb w14, [x14, x13]
  2076: ec030e2a                 mov w12, w14
  2080: 0e1140f9                 ldr x14, [x8, #0x20]
  2084: ce696d38                 ldrb w14, [x14, x13]
  2088: 8c010e4a                 eor w12, w12, w14
  2092: 0e2140f9                 ldr x14, [x8, #0x40]
  2096: ce696d38                 ldrb w14, [x14, x13]
  2100: 8c010e4a                 eor w12, w12, w14
  2104: 0e2940f9                 ldr x14, [x8, #0x50]
  2108: ce696d38                 ldrb w14, [x14, x13]
  2112: 8c010e4a                 eor w12, w12, w14
  2116: 0e3540f9                 ldr x14, [x8, #0x68]
  2120: ce696d38                 ldrb w14, [x14, x13]
  2124: 8c010e4a                 eor w12, w12, w14
  2128: 0e4140f9                 ldr x14, [x8, #0x80]
  2132: ce696d38                 ldrb w14, [x14, x13]
  2136: 8c010e4a                 eor w12, w12, w14
  2140: 0e4940f9                 ldr x14, [x8, #0x90]
  2144: ce696d38                 ldrb w14, [x14, x13]
  2148: 8c010e4a                 eor w12, w12, w14
  2152: 0e5940f9                 ldr x14, [x8, #0xb0]
  2156: ce696d38                 ldrb w14, [x14, x13]
  2160: 8c010e4a                 eor w12, w12, w14
  2164: 0e7140f9                 ldr x14, [x8, #0xe0]
  2168: ce696d38                 ldrb w14, [x14, x13]
  2172: 8c010e4a                 eor w12, w12, w14
  2176: 0e8940f9                 ldr x14, [x8, #0x110]
  2180: ce696d38                 ldrb w14, [x14, x13]
  2184: 8c010e4a                 eor w12, w12, w14
  2188: 0e8d40f9                 ldr x14, [x8, #0x118]
  2192: ce696d38                 ldrb w14, [x14, x13]
  2196: 8c010e4a                 eor w12, w12, w14
  2200: 4c692d38                 strb w12, [x10, x13]
  2204: ad050091                 add x13, x13, #1
  2208: bf0102eb                 cmp x13, x2
  2212: 83fbff54                 b.lo 0x814
  2216: 2a2940f9                 ldr x10, [x9, #0x50]
  2220: ed030baa                 mov x13, x11
  2224: 0e1140f9                 ldr x14, [x8, #0x20]
  2228: ce696d38                 ldrb w14, [x14, x13]
  2232: ec030e2a                 mov w12, w14
  2236: 0e1540f9                 ldr x14, [x8, #0x28]
  2240: ce696d38                 ldrb w14, [x14, x13]
  2244: 8c010e4a                 eor w12, w12, w14
  2248: 0e2540f9                 ldr x14, [x8, #0x48]
  2252: ce696d38                 ldrb w14, [x14, x13]
  2256: 8c010e4a                 eor w12, w12, w14
  2260: 0e2d40f9                 ldr x14, [x8, #0x58]
  2264: ce696d38                 ldrb w14, [x14, x13]
  2268: 8c010e4a                 eor w12, w12, w14
  2272: 0e3940f9                 ldr x14, [x8, #0x70]
  2276: ce696d38                 ldrb w14, [x14, x13]
  2280: 8c010e4a                 eor w12, w12, w14
  2284: 0e4540f9                 ldr x14, [x8, #0x88]
  2288: ce696d38                 ldrb w14, [x14, x13]
  2292: 8c010e4a                 eor w12, w12, w14
  2296: 0e4d40f9                 ldr x14, [x8, #0x98]
  2300: ce696d38                 ldrb w14, [x14, x13]
  2304: 8c010e4a                 eor w12, w12, w14
  2308: 0e5d40f9                 ldr x14, [x8, #0xb8]
  2312: ce696d38                 ldrb w14, [x14, x13]
  2316: 8c010e4a                 eor w12, w12, w14
  2320: 0e6140f9                 ldr x14, [x8, #0xc0]
  2324: ce696d38                 ldrb w14, [x14, x13]
  2328: 8c010e4a                 eor w12, w12, w14
  2332: 0e7540f9                 ldr x14, [x8, #0xe8]
  2336: ce696d38                 ldrb w14, [x14, x13]
  2340: 8c010e4a                 eor w12, w12, w14
  2344: 0e8d40f9                 ldr x14, [x8, #0x118]
  2348: ce696d38                 ldrb w14, [x14, x13]
  2352: 8c010e4a                 eor w12, w12, w14
  2356: 4c692d38                 strb w12, [x10, x13]
  2360: ad050091                 add x13, x13, #1
  2364: bf0102eb                 cmp x13, x2
  2368: 83fbff54                 b.lo 0x8b0
  2372: 2a2d40f9                 ldr x10, [x9, #0x58]
  2376: ed030baa                 mov x13, x11
  2380: 0e1540f9                 ldr x14, [x8, #0x28]
  2384: ce696d38                 ldrb w14, [x14, x13]
  2388: ec030e2a                 mov w12, w14
  2392: 0e2940f9                 ldr x14, [x8, #0x50]
  2396: ce696d38                 ldrb w14, [x14, x13]
  2400: 8c010e4a                 eor w12, w12, w14
  2404: 0e3d40f9                 ldr x14, [x8, #0x78]
  2408: ce696d38                 ldrb w14, [x14, x13]
  2412: 8c010e4a                 eor w12, w12, w14
  2416: 0e5140f9                 ldr x14, [x8, #0xa0]
  2420: ce696d38                 ldrb w14, [x14, x13]
  2424: 8c010e4a                 eor w12, w12, w14
  2428: 0e6540f9                 ldr x14, [x8, #0xc8]
  2432: ce696d38                 ldrb w14, [x14, x13]
  2436: 8c010e4a                 eor w12, w12, w14
  2440: 0e7940f9                 ldr x14, [x8, #0xf0]
  2444: ce696d38                 ldrb w14, [x14, x13]
  2448: 8c010e4a                 eor w12, w12, w14
  2452: 4c692d38                 strb w12, [x10, x13]
  2456: ad050091                 add x13, x13, #1
  2460: bf0102eb                 cmp x13, x2
  2464: 63fdff54                 b.lo 0x94c
  2468: c0035fd6                 ret 
  2472: 00000000                 udf #0
  2476: 00000000                 udf #0
  2480: 00000000                 udf #0
  2484: 00000000                 udf #0
  2488: 00000000                 udf #0
  2492: 00000000                 udf #0
