; config_t1_1717 target=rvv-zvbc32e twin=nt topo=raid5(4)/enc op=encode U=1
; code_size=272 twins=["nt"] uops_per_byte=1.125
     0: 63140600                 bnez a2, 8
     4: 6f00c00f                 j 0x100
     8: 970f0000                 auipc t6, 0
    12: 938f8f10                 addi t6, t6, 0x108
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f004007                 j 0x9c
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
   124: 83b30500                 ld t2, 0(a1)
   128: b3836300                 add t2, t2, t1
   132: 33005000                 ntl.all 
   136: 27800302                 vse8.v v0, (t2)
   140: 3303d300                 add t1, t1, a3
   144: 1307f7ff                 addi a4, a4, -1
   148: e31607fa                 bnez a4, 0x40
   152: 33066640                 sub a2, a2, t1
   156: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   160: 63940200                 bnez t0, 0xa8
   164: 6f00c005                 j 0x100
   168: 83330500                 ld t2, 0(a0)
   172: b3836300                 add t2, t2, t1
   176: 07800302                 vle8.v v0, (t2)
   180: 83338500                 ld t2, 8(a0)
   184: b3836300                 add t2, t2, t1
   188: 87800302                 vle8.v v1, (t2)
   192: 5700102e                 vxor.vv v0, v1, v0
   196: 83330501                 ld t2, 0x10(a0)
   200: b3836300                 add t2, t2, t1
   204: 87800302                 vle8.v v1, (t2)
   208: 5700102e                 vxor.vv v0, v1, v0
   212: 83338501                 ld t2, 0x18(a0)
   216: b3836300                 add t2, t2, t1
   220: 87800302                 vle8.v v1, (t2)
   224: 5700102e                 vxor.vv v0, v1, v0
   228: 83b30500                 ld t2, 0(a1)
   232: b3836300                 add t2, t2, t1
   236: 33005000                 ntl.all 
   240: 27800302                 vse8.v v0, (t2)
   244: 33035300                 add t1, t1, t0
   248: 33065640                 sub a2, a2, t0
   252: 6ff01ffa                 j 0x9c
   256: 67800000                 ret 
   260: 0000                     unimp 
   262: 0000                     unimp 
   264: 0000                     unimp 
   266: 0000                     unimp 
   268: 0000                     unimp 
   270: 0000                     unimp 
