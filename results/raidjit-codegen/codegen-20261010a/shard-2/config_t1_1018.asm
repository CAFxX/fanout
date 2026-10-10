; config_t1_1018 target=sve2 twin=notail topo=rs(6,3)/rec{0} op=recover U=1
; code_size=192 twins=["notail"] uops_per_byte=0.75
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: ed030caa                 mov x13, x12
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: ce0300b4                 cbz x14, 0x94
    32: 0a0140f9                 ldr x10, [x8]
    36: 4a010b8b                 add x10, x10, x11
    40: 50a100a4                 ld1b { z16.b }, p0/z, [x10]
    44: 0a0540f9                 ldr x10, [x8, #8]
    48: 4a010b8b                 add x10, x10, x11
    52: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
    56: 30021904                 eor z16.b, p0/m, z16.b, z17.b
    60: 0a0940f9                 ldr x10, [x8, #0x10]
    64: 4a010b8b                 add x10, x10, x11
    68: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
    72: 30021904                 eor z16.b, p0/m, z16.b, z17.b
    76: 0a0d40f9                 ldr x10, [x8, #0x18]
    80: 4a010b8b                 add x10, x10, x11
    84: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
    88: 30021904                 eor z16.b, p0/m, z16.b, z17.b
    92: 0a1140f9                 ldr x10, [x8, #0x20]
    96: 4a010b8b                 add x10, x10, x11
   100: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
   104: 30021904                 eor z16.b, p0/m, z16.b, z17.b
   108: 0a1540f9                 ldr x10, [x8, #0x28]
   112: 4a010b8b                 add x10, x10, x11
   116: 51a100a4                 ld1b { z17.b }, p0/z, [x10]
   120: 30021904                 eor z16.b, p0/m, z16.b, z17.b
   124: 2a0140f9                 ldr x10, [x9]
   128: 4a010b8b                 add x10, x10, x11
   132: 50e100e4                 st1b { z16.b }, p0, [x10]
   136: 6b010d8b                 add x11, x11, x13
   140: ce0500f1                 subs x14, x14, #1
   144: 81fcff54                 b.ne 0x20
   148: c0035fd6                 ret 
   152: 00000000                 udf #0
   156: 00000000                 udf #0
   160: 00000000                 udf #0
   164: 00000000                 udf #0
   168: 00000000                 udf #0
   172: 00000000                 udf #0
   176: 00000000                 udf #0
   180: 00000000                 udf #0
   184: 00000000                 udf #0
   188: 00000000                 udf #0
