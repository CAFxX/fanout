; config_t1_1876 target=rvv-zvbc32e twin=notail topo=rs(6,3)/rec{0} op=recover U=1
; code_size=192 twins=["notail"] uops_per_byte=0.8125
     0: 63140600                 bnez a2, 8
     4: 6f00000b                 j 0xb4
     8: 970f0000                 auipc t6, 0
    12: 938f8f0b                 addi t6, t6, 0xb8
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f00c008                 j 0xb4
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07800302                 vle8.v v0, (t2)
    76: 83338500                 ld t2, 8(a0)
    80: b3836300                 add t2, t2, t1
    84: 87800302                 vle8.v v1, (t2)
    88: 5700102e                 vxor.vv v0, v1, v0
    92: 83330501                 ld t2, 0x10(a0)
    96: b3836300                 add t2, t2, t1
   100: 87800302                 vle8.v v1, (t2)
   104: 5700102e                 vxor.vv v0, v1, v0
   108: 83338501                 ld t2, 0x18(a0)
   112: b3836300                 add t2, t2, t1
   116: 87800302                 vle8.v v1, (t2)
   120: 5700102e                 vxor.vv v0, v1, v0
   124: 83330502                 ld t2, 0x20(a0)
   128: b3836300                 add t2, t2, t1
   132: 87800302                 vle8.v v1, (t2)
   136: 5700102e                 vxor.vv v0, v1, v0
   140: 83338502                 ld t2, 0x28(a0)
   144: b3836300                 add t2, t2, t1
   148: 87800302                 vle8.v v1, (t2)
   152: 5700102e                 vxor.vv v0, v1, v0
   156: 83b30500                 ld t2, 0(a1)
   160: b3836300                 add t2, t2, t1
   164: 27800302                 vse8.v v0, (t2)
   168: 3303d300                 add t1, t1, a3
   172: 1307f7ff                 addi a4, a4, -1
   176: e31807f8                 bnez a4, 0x40
   180: 67800000                 ret 
   184: 0000                     unimp 
   186: 0000                     unimp 
   188: 0000                     unimp 
   190: 0000                     unimp 
