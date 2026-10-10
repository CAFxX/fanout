; config_t3_2080 target=sve2 twin=baseline topo=rs(3,1)/enc op=encode U=8
; code_size=384 twins=[] uops_per_byte=0.375
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: 8df17dd3                 lsl x13, x12, #3
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: 8e0700b4                 cbz x14, 0x10c
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
   216: 2a0140f9                 ldr x10, [x9]
   220: 4a010b8b                 add x10, x10, x11
   224: 50e100e4                 st1b { z16.b }, p0, [x10]
   228: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
   232: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
   236: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
   240: 54e104e4                 st1b { z20.b }, p0, [x10, #4, mul vl]
   244: 55e105e4                 st1b { z21.b }, p0, [x10, #5, mul vl]
   248: 56e106e4                 st1b { z22.b }, p0, [x10, #6, mul vl]
   252: 57e107e4                 st1b { z23.b }, p0, [x10, #7, mul vl]
   256: 6b010d8b                 add x11, x11, x13
   260: ce0500f1                 subs x14, x14, #1
   264: c1f8ff54                 b.ne 0x20
   268: 7f0102eb                 cmp x11, x2
   272: 02020054                 b.hs 0x150
   276: 2a0140f9                 ldr x10, [x9]
   280: ed030baa                 mov x13, x11
   284: 0e0140f9                 ldr x14, [x8]
   288: ce696d38                 ldrb w14, [x14, x13]
   292: ec030e2a                 mov w12, w14
   296: 0e0540f9                 ldr x14, [x8, #8]
   300: ce696d38                 ldrb w14, [x14, x13]
   304: 8c010e4a                 eor w12, w12, w14
   308: 0e0940f9                 ldr x14, [x8, #0x10]
   312: ce696d38                 ldrb w14, [x14, x13]
   316: 8c010e4a                 eor w12, w12, w14
   320: 4c692d38                 strb w12, [x10, x13]
   324: ad050091                 add x13, x13, #1
   328: bf0102eb                 cmp x13, x2
   332: 83feff54                 b.lo 0x11c
   336: c0035fd6                 ret 
   340: 00000000                 udf #0
   344: 00000000                 udf #0
   348: 00000000                 udf #0
   352: 00000000                 udf #0
   356: 00000000                 udf #0
   360: 00000000                 udf #0
   364: 00000000                 udf #0
   368: 00000000                 udf #0
   372: 00000000                 udf #0
   376: 00000000                 udf #0
   380: 00000000                 udf #0
