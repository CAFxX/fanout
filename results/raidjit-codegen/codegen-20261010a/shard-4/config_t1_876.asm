; config_t1_876 target=sve2 twin=predtail topo=raid5(4)/enc op=encode U=8
; code_size=448 twins=["predtail"] uops_per_byte=0.5
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: 8df17dd3                 lsl x13, x12, #3
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: ce0900b4                 cbz x14, 0x154
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
   288: 2a0140f9                 ldr x10, [x9]
   292: 4a010b8b                 add x10, x10, x11
   296: 50e100e4                 st1b { z16.b }, p0, [x10]
   300: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
   304: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
   308: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
   312: 54e104e4                 st1b { z20.b }, p0, [x10, #4, mul vl]
   316: 55e105e4                 st1b { z21.b }, p0, [x10, #5, mul vl]
   320: 56e106e4                 st1b { z22.b }, p0, [x10, #6, mul vl]
   324: 57e107e4                 st1b { z23.b }, p0, [x10, #7, mul vl]
   328: 6b010d8b                 add x11, x11, x13
   332: ce0500f1                 subs x14, x14, #1
   336: 81f6ff54                 b.ne 0x20
   340: 7f0102eb                 cmp x11, x2
   344: 02030054                 b.hs 0x1b8
   348: 4d000bcb                 sub x13, x2, x11
   352: e1172d25                 whilelt p1.b, xzr, x13
   356: 0a0140f9                 ldr x10, [x8]
   360: 4a010b8b                 add x10, x10, x11
   364: 50a500a4                 ld1b { z16.b }, p1/z, [x10]
   368: 0a0540f9                 ldr x10, [x8, #8]
   372: 4a010b8b                 add x10, x10, x11
   376: 58a500a4                 ld1b { z24.b }, p1/z, [x10]
   380: 10071904                 eor z16.b, p1/m, z16.b, z24.b
   384: 0a0940f9                 ldr x10, [x8, #0x10]
   388: 4a010b8b                 add x10, x10, x11
   392: 58a500a4                 ld1b { z24.b }, p1/z, [x10]
   396: 10071904                 eor z16.b, p1/m, z16.b, z24.b
   400: 0a0d40f9                 ldr x10, [x8, #0x18]
   404: 4a010b8b                 add x10, x10, x11
   408: 58a500a4                 ld1b { z24.b }, p1/z, [x10]
   412: 10071904                 eor z16.b, p1/m, z16.b, z24.b
   416: 2a0140f9                 ldr x10, [x9]
   420: 4a010b8b                 add x10, x10, x11
   424: 50e500e4                 st1b { z16.b }, p1, [x10]
   428: 6b010c8b                 add x11, x11, x12
   432: ad010ceb                 subs x13, x13, x12
   436: 6cfdff54                 b.gt 0x160
   440: c0035fd6                 ret 
   444: 00000000                 udf #0
