; config_t1_1873 target=rvv-zvbc32e twin=nt topo=rs(6,3)/rec{0} op=recover U=1
; code_size=336 twins=["nt"] uops_per_byte=1.625
     0: 63140600                 bnez a2, 8
     4: 6f00c013                 j 0x140
     8: 970f0000                 auipc t6, 0
    12: 938f8f14                 addi t6, t6, 0x148
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f004009                 j 0xbc
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
   164: 33005000                 ntl.all 
   168: 27800302                 vse8.v v0, (t2)
   172: 3303d300                 add t1, t1, a3
   176: 1307f7ff                 addi a4, a4, -1
   180: e31607f8                 bnez a4, 0x40
   184: 33066640                 sub a2, a2, t1
   188: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   192: 63940200                 bnez t0, 0xc8
   196: 6f00c007                 j 0x140
   200: 83330500                 ld t2, 0(a0)
   204: b3836300                 add t2, t2, t1
   208: 07800302                 vle8.v v0, (t2)
   212: 83338500                 ld t2, 8(a0)
   216: b3836300                 add t2, t2, t1
   220: 87800302                 vle8.v v1, (t2)
   224: 5700102e                 vxor.vv v0, v1, v0
   228: 83330501                 ld t2, 0x10(a0)
   232: b3836300                 add t2, t2, t1
   236: 87800302                 vle8.v v1, (t2)
   240: 5700102e                 vxor.vv v0, v1, v0
   244: 83338501                 ld t2, 0x18(a0)
   248: b3836300                 add t2, t2, t1
   252: 87800302                 vle8.v v1, (t2)
   256: 5700102e                 vxor.vv v0, v1, v0
   260: 83330502                 ld t2, 0x20(a0)
   264: b3836300                 add t2, t2, t1
   268: 87800302                 vle8.v v1, (t2)
   272: 5700102e                 vxor.vv v0, v1, v0
   276: 83338502                 ld t2, 0x28(a0)
   280: b3836300                 add t2, t2, t1
   284: 87800302                 vle8.v v1, (t2)
   288: 5700102e                 vxor.vv v0, v1, v0
   292: 83b30500                 ld t2, 0(a1)
   296: b3836300                 add t2, t2, t1
   300: 33005000                 ntl.all 
   304: 27800302                 vse8.v v0, (t2)
   308: 33035300                 add t1, t1, t0
   312: 33065640                 sub a2, a2, t0
   316: 6ff01ff8                 j 0xbc
   320: 67800000                 ret 
   324: 0000                     unimp 
   326: 0000                     unimp 
   328: 0000                     unimp 
   330: 0000                     unimp 
   332: 0000                     unimp 
   334: 0000                     unimp 
