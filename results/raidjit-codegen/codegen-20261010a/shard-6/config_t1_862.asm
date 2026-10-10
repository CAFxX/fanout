; config_t1_862 target=sve2 twin=notail topo=raid5(4)/enc op=encode U=1
; code_size=128 twins=["notail"] uops_per_byte=0.5
     0: e80300aa                 mov x8, x0
     4: e90301aa                 mov x9, x1
     8: eb031faa                 mov x11, xzr
    12: ece32004                 cntb x12
    16: ed030caa                 mov x13, x12
    20: 4e08cd9a                 udiv x14, x2, x13
    24: e01f2c25                 whilelo p0.b, xzr, x12
    28: ce0200b4                 cbz x14, 0x74
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
    92: 2a0140f9                 ldr x10, [x9]
    96: 4a010b8b                 add x10, x10, x11
   100: 50e100e4                 st1b { z16.b }, p0, [x10]
   104: 6b010d8b                 add x11, x11, x13
   108: ce0500f1                 subs x14, x14, #1
   112: 81fdff54                 b.ne 0x20
   116: c0035fd6                 ret 
   120: 00000000                 udf #0
   124: 00000000                 udf #0
