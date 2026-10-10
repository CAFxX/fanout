; config_t1_914 target=sve2 twin=notail topo=rdp(7)/enc op=encode U=1
; code_size=1024 twins=["notail"] uops_per_byte=9.0625
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
   860: 50e100e4                 st1b { z16.b }, p0, [x10]
   864: 2a0540f9                 ldr x10, [x9, #8]
   868: 4a010b8b                 add x10, x10, x11
   872: 51e100e4                 st1b { z17.b }, p0, [x10]
   876: 2a0940f9                 ldr x10, [x9, #0x10]
   880: 4a010b8b                 add x10, x10, x11
   884: 52e100e4                 st1b { z18.b }, p0, [x10]
   888: 2a0d40f9                 ldr x10, [x9, #0x18]
   892: 4a010b8b                 add x10, x10, x11
   896: 53e100e4                 st1b { z19.b }, p0, [x10]
   900: 2a1140f9                 ldr x10, [x9, #0x20]
   904: 4a010b8b                 add x10, x10, x11
   908: 54e100e4                 st1b { z20.b }, p0, [x10]
   912: 2a1540f9                 ldr x10, [x9, #0x28]
   916: 4a010b8b                 add x10, x10, x11
   920: 55e100e4                 st1b { z21.b }, p0, [x10]
   924: 2a1940f9                 ldr x10, [x9, #0x30]
   928: 4a010b8b                 add x10, x10, x11
   932: 56e100e4                 st1b { z22.b }, p0, [x10]
   936: 2a1d40f9                 ldr x10, [x9, #0x38]
   940: 4a010b8b                 add x10, x10, x11
   944: 57e100e4                 st1b { z23.b }, p0, [x10]
   948: 2a2140f9                 ldr x10, [x9, #0x40]
   952: 4a010b8b                 add x10, x10, x11
   956: 58e100e4                 st1b { z24.b }, p0, [x10]
   960: 2a2540f9                 ldr x10, [x9, #0x48]
   964: 4a010b8b                 add x10, x10, x11
   968: 59e100e4                 st1b { z25.b }, p0, [x10]
   972: 2a2940f9                 ldr x10, [x9, #0x50]
   976: 4a010b8b                 add x10, x10, x11
   980: 5ae100e4                 st1b { z26.b }, p0, [x10]
   984: 2a2d40f9                 ldr x10, [x9, #0x58]
   988: 4a010b8b                 add x10, x10, x11
   992: 5be100e4                 st1b { z27.b }, p0, [x10]
   996: 6b010d8b                 add x11, x11, x13
  1000: ce0500f1                 subs x14, x14, #1
  1004: a1e1ff54                 b.ne 0x20
  1008: c0035fd6                 ret 
  1012: 00000000                 udf #0
  1016: 00000000                 udf #0
  1020: 00000000                 udf #0
