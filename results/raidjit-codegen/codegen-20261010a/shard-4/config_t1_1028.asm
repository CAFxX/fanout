; config_t1_1028 target=sve2 twin=nt topo=rs(6,3)/rec{0} op=recover U=8
; code_size=640 twins=["nt"] uops_per_byte=0.75
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: 8df17dd3                 lsl x13, x12, #3
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: 4e0e00b4                 cbz x14, 0x1e4
    32: 0a0140f9                 ldr x10, [x8]
    36: 4a010b8b                 add x10, x10, x11
    40: 50a100a4                 ld1b { z16.b }, p0/z, [x10]
    44: 51a101a4                 ld1b { z17.b }, p0/z, [x10, #1, mul vl]
    48: 52a102a4                 ld1b { z18.b }, p0/z, [x10, #2, mul vl]
    52: 53a103a4                 ld1b { z19.b }, p0/z, [x10, #3, mul vl]
    56: 54a104a4                 ld1b { z20.b }, p0/z, [x10, #4, mul vl]
    60: 55a105a4                 ld1b { z21.b }, p0/z, [x10, #5, mul vl]
    64: 56a106a4                 ld1b { z22.b }, p0/z, [x10, #6, mul vl]
    68: 57a107a4                 ld1b { z23.b }, p0/z, [x10, #7, mul vl]
    72: 0a0540f9                 ldr x10, [x8, #8]
    76: 4a010b8b                 add x10, x10, x11
    80: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
    84: 10031904                 eor z16.b, p0/m, z16.b, z24.b
    88: 58a101a4                 ld1b { z24.b }, p0/z, [x10, #1, mul vl]
    92: 11031904                 eor z17.b, p0/m, z17.b, z24.b
    96: 58a102a4                 ld1b { z24.b }, p0/z, [x10, #2, mul vl]
   100: 12031904                 eor z18.b, p0/m, z18.b, z24.b
   104: 58a103a4                 ld1b { z24.b }, p0/z, [x10, #3, mul vl]
   108: 13031904                 eor z19.b, p0/m, z19.b, z24.b
   112: 58a104a4                 ld1b { z24.b }, p0/z, [x10, #4, mul vl]
   116: 14031904                 eor z20.b, p0/m, z20.b, z24.b
   120: 58a105a4                 ld1b { z24.b }, p0/z, [x10, #5, mul vl]
   124: 15031904                 eor z21.b, p0/m, z21.b, z24.b
   128: 58a106a4                 ld1b { z24.b }, p0/z, [x10, #6, mul vl]
   132: 16031904                 eor z22.b, p0/m, z22.b, z24.b
   136: 58a107a4                 ld1b { z24.b }, p0/z, [x10, #7, mul vl]
   140: 17031904                 eor z23.b, p0/m, z23.b, z24.b
   144: 0a0940f9                 ldr x10, [x8, #0x10]
   148: 4a010b8b                 add x10, x10, x11
   152: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   156: 10031904                 eor z16.b, p0/m, z16.b, z24.b
   160: 58a101a4                 ld1b { z24.b }, p0/z, [x10, #1, mul vl]
   164: 11031904                 eor z17.b, p0/m, z17.b, z24.b
   168: 58a102a4                 ld1b { z24.b }, p0/z, [x10, #2, mul vl]
   172: 12031904                 eor z18.b, p0/m, z18.b, z24.b
   176: 58a103a4                 ld1b { z24.b }, p0/z, [x10, #3, mul vl]
   180: 13031904                 eor z19.b, p0/m, z19.b, z24.b
   184: 58a104a4                 ld1b { z24.b }, p0/z, [x10, #4, mul vl]
   188: 14031904                 eor z20.b, p0/m, z20.b, z24.b
   192: 58a105a4                 ld1b { z24.b }, p0/z, [x10, #5, mul vl]
   196: 15031904                 eor z21.b, p0/m, z21.b, z24.b
   200: 58a106a4                 ld1b { z24.b }, p0/z, [x10, #6, mul vl]
   204: 16031904                 eor z22.b, p0/m, z22.b, z24.b
   208: 58a107a4                 ld1b { z24.b }, p0/z, [x10, #7, mul vl]
   212: 17031904                 eor z23.b, p0/m, z23.b, z24.b
   216: 0a0d40f9                 ldr x10, [x8, #0x18]
   220: 4a010b8b                 add x10, x10, x11
   224: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   228: 10031904                 eor z16.b, p0/m, z16.b, z24.b
   232: 58a101a4                 ld1b { z24.b }, p0/z, [x10, #1, mul vl]
   236: 11031904                 eor z17.b, p0/m, z17.b, z24.b
   240: 58a102a4                 ld1b { z24.b }, p0/z, [x10, #2, mul vl]
   244: 12031904                 eor z18.b, p0/m, z18.b, z24.b
   248: 58a103a4                 ld1b { z24.b }, p0/z, [x10, #3, mul vl]
   252: 13031904                 eor z19.b, p0/m, z19.b, z24.b
   256: 58a104a4                 ld1b { z24.b }, p0/z, [x10, #4, mul vl]
   260: 14031904                 eor z20.b, p0/m, z20.b, z24.b
   264: 58a105a4                 ld1b { z24.b }, p0/z, [x10, #5, mul vl]
   268: 15031904                 eor z21.b, p0/m, z21.b, z24.b
   272: 58a106a4                 ld1b { z24.b }, p0/z, [x10, #6, mul vl]
   276: 16031904                 eor z22.b, p0/m, z22.b, z24.b
   280: 58a107a4                 ld1b { z24.b }, p0/z, [x10, #7, mul vl]
   284: 17031904                 eor z23.b, p0/m, z23.b, z24.b
   288: 0a1140f9                 ldr x10, [x8, #0x20]
   292: 4a010b8b                 add x10, x10, x11
   296: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   300: 10031904                 eor z16.b, p0/m, z16.b, z24.b
   304: 58a101a4                 ld1b { z24.b }, p0/z, [x10, #1, mul vl]
   308: 11031904                 eor z17.b, p0/m, z17.b, z24.b
   312: 58a102a4                 ld1b { z24.b }, p0/z, [x10, #2, mul vl]
   316: 12031904                 eor z18.b, p0/m, z18.b, z24.b
   320: 58a103a4                 ld1b { z24.b }, p0/z, [x10, #3, mul vl]
   324: 13031904                 eor z19.b, p0/m, z19.b, z24.b
   328: 58a104a4                 ld1b { z24.b }, p0/z, [x10, #4, mul vl]
   332: 14031904                 eor z20.b, p0/m, z20.b, z24.b
   336: 58a105a4                 ld1b { z24.b }, p0/z, [x10, #5, mul vl]
   340: 15031904                 eor z21.b, p0/m, z21.b, z24.b
   344: 58a106a4                 ld1b { z24.b }, p0/z, [x10, #6, mul vl]
   348: 16031904                 eor z22.b, p0/m, z22.b, z24.b
   352: 58a107a4                 ld1b { z24.b }, p0/z, [x10, #7, mul vl]
   356: 17031904                 eor z23.b, p0/m, z23.b, z24.b
   360: 0a1540f9                 ldr x10, [x8, #0x28]
   364: 4a010b8b                 add x10, x10, x11
   368: 58a100a4                 ld1b { z24.b }, p0/z, [x10]
   372: 10031904                 eor z16.b, p0/m, z16.b, z24.b
   376: 58a101a4                 ld1b { z24.b }, p0/z, [x10, #1, mul vl]
   380: 11031904                 eor z17.b, p0/m, z17.b, z24.b
   384: 58a102a4                 ld1b { z24.b }, p0/z, [x10, #2, mul vl]
   388: 12031904                 eor z18.b, p0/m, z18.b, z24.b
   392: 58a103a4                 ld1b { z24.b }, p0/z, [x10, #3, mul vl]
   396: 13031904                 eor z19.b, p0/m, z19.b, z24.b
   400: 58a104a4                 ld1b { z24.b }, p0/z, [x10, #4, mul vl]
   404: 14031904                 eor z20.b, p0/m, z20.b, z24.b
   408: 58a105a4                 ld1b { z24.b }, p0/z, [x10, #5, mul vl]
   412: 15031904                 eor z21.b, p0/m, z21.b, z24.b
   416: 58a106a4                 ld1b { z24.b }, p0/z, [x10, #6, mul vl]
   420: 16031904                 eor z22.b, p0/m, z22.b, z24.b
   424: 58a107a4                 ld1b { z24.b }, p0/z, [x10, #7, mul vl]
   428: 17031904                 eor z23.b, p0/m, z23.b, z24.b
   432: 2a0140f9                 ldr x10, [x9]
   436: 4a010b8b                 add x10, x10, x11
   440: 50e110e4                 stnt1b { z16.b }, p0, [x10]
   444: 51e111e4                 stnt1b { z17.b }, p0, [x10, #1, mul vl]
   448: 52e112e4                 stnt1b { z18.b }, p0, [x10, #2, mul vl]
   452: 53e113e4                 stnt1b { z19.b }, p0, [x10, #3, mul vl]
   456: 54e114e4                 stnt1b { z20.b }, p0, [x10, #4, mul vl]
   460: 55e115e4                 stnt1b { z21.b }, p0, [x10, #5, mul vl]
   464: 56e116e4                 stnt1b { z22.b }, p0, [x10, #6, mul vl]
   468: 57e117e4                 stnt1b { z23.b }, p0, [x10, #7, mul vl]
   472: 6b010d8b                 add x11, x11, x13
   476: ce0500f1                 subs x14, x14, #1
   480: 01f2ff54                 b.ne 0x20
   484: 7f0102eb                 cmp x11, x2
   488: 22030054                 b.hs 0x24c
   492: 2a0140f9                 ldr x10, [x9]
   496: ed030baa                 mov x13, x11
   500: 0e0140f9                 ldr x14, [x8]
   504: ce696d38                 ldrb w14, [x14, x13]
   508: ec030e2a                 mov w12, w14
   512: 0e0540f9                 ldr x14, [x8, #8]
   516: ce696d38                 ldrb w14, [x14, x13]
   520: 8c010e4a                 eor w12, w12, w14
   524: 0e0940f9                 ldr x14, [x8, #0x10]
   528: ce696d38                 ldrb w14, [x14, x13]
   532: 8c010e4a                 eor w12, w12, w14
   536: 0e0d40f9                 ldr x14, [x8, #0x18]
   540: ce696d38                 ldrb w14, [x14, x13]
   544: 8c010e4a                 eor w12, w12, w14
   548: 0e1140f9                 ldr x14, [x8, #0x20]
   552: ce696d38                 ldrb w14, [x14, x13]
   556: 8c010e4a                 eor w12, w12, w14
   560: 0e1540f9                 ldr x14, [x8, #0x28]
   564: ce696d38                 ldrb w14, [x14, x13]
   568: 8c010e4a                 eor w12, w12, w14
   572: 4c692d38                 strb w12, [x10, x13]
   576: ad050091                 add x13, x13, #1
   580: bf0102eb                 cmp x13, x2
   584: 63fdff54                 b.lo 0x1f4
   588: c0035fd6                 ret 
   592: 00000000                 udf #0
   596: 00000000                 udf #0
   600: 00000000                 udf #0
   604: 00000000                 udf #0
   608: 00000000                 udf #0
   612: 00000000                 udf #0
   616: 00000000                 udf #0
   620: 00000000                 udf #0
   624: 00000000                 udf #0
   628: 00000000                 udf #0
   632: 00000000                 udf #0
   636: 00000000                 udf #0
