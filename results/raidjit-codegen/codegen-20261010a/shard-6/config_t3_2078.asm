; config_t3_2078 target=sve2 twin=baseline topo=raid1()/enc op=encode U=8
; code_size=192 twins=[] uops_per_byte=0.125
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: 8df17dd3                 lsl x13, x12, #3
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: 0e0300b4                 cbz x14, 0x7c
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
    72: 2a0140f9                 ldr x10, [x9]
    76: 4a010b8b                 add x10, x10, x11
    80: 50e100e4                 st1b { z16.b }, p0, [x10]
    84: 51e101e4                 st1b { z17.b }, p0, [x10, #1, mul vl]
    88: 52e102e4                 st1b { z18.b }, p0, [x10, #2, mul vl]
    92: 53e103e4                 st1b { z19.b }, p0, [x10, #3, mul vl]
    96: 54e104e4                 st1b { z20.b }, p0, [x10, #4, mul vl]
   100: 55e105e4                 st1b { z21.b }, p0, [x10, #5, mul vl]
   104: 56e106e4                 st1b { z22.b }, p0, [x10, #6, mul vl]
   108: 57e107e4                 st1b { z23.b }, p0, [x10, #7, mul vl]
   112: 6b010d8b                 add x11, x11, x13
   116: ce0500f1                 subs x14, x14, #1
   120: 41fdff54                 b.ne 0x20
   124: 7f0102eb                 cmp x11, x2
   128: 42010054                 b.hs 0xa8
   132: 2a0140f9                 ldr x10, [x9]
   136: ed030baa                 mov x13, x11
   140: 0e0140f9                 ldr x14, [x8]
   144: ce696d38                 ldrb w14, [x14, x13]
   148: ec030e2a                 mov w12, w14
   152: 4c692d38                 strb w12, [x10, x13]
   156: ad050091                 add x13, x13, #1
   160: bf0102eb                 cmp x13, x2
   164: 43ffff54                 b.lo 0x8c
   168: c0035fd6                 ret 
   172: 00000000                 udf #0
   176: 00000000                 udf #0
   180: 00000000                 udf #0
   184: 00000000                 udf #0
   188: 00000000                 udf #0
