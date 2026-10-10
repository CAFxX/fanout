; config_t3_2089 target=rvv twin=baseline topo=raid1()/enc op=encode U=1
; code_size=160 twins=[] uops_per_byte=0.375
     0: 63140600                 bnez a2, 8
     4: 6f004009                 j 0x98
     8: 970f0000                 auipc t6, 0
    12: 938f8f09                 addi t6, t6, 0x98
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f000004                 j 0x68
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07800302                 vle8.v v0, (t2)
    76: 83b30500                 ld t2, 0(a1)
    80: b3836300                 add t2, t2, t1
    84: 27800302                 vse8.v v0, (t2)
    88: 3303d300                 add t1, t1, a3
    92: 1307f7ff                 addi a4, a4, -1
    96: e31007fe                 bnez a4, 0x40
   100: 33066640                 sub a2, a2, t1
   104: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   108: 63940200                 bnez t0, 0x74
   112: 6f008002                 j 0x98
   116: 83330500                 ld t2, 0(a0)
   120: b3836300                 add t2, t2, t1
   124: 07800302                 vle8.v v0, (t2)
   128: 83b30500                 ld t2, 0(a1)
   132: b3836300                 add t2, t2, t1
   136: 27800302                 vse8.v v0, (t2)
   140: 33035300                 add t1, t1, t0
   144: 33065640                 sub a2, a2, t0
   148: 6ff05ffd                 j 0x68
   152: 67800000                 ret 
   156: 0000                     unimp 
   158: 0000                     unimp 
