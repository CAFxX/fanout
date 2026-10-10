; config_t1_1722 target=rvv-zvbc32e twin=maskedtail topo=raid5(4)/enc op=encode U=1
; code_size=288 twins=["maskedtail"] uops_per_byte=1.6875
     0: 63140600                 bnez a2, 8
     4: 6f00c010                 j 0x110
     8: 970f0000                 auipc t6, 0
    12: 938f8f11                 addi t6, t6, 0x118
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f000007                 j 0x98
    44: 13000000                 nop 
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 87800302                 vle8.v v1, (t2)
    76: 83338500                 ld t2, 8(a0)
    80: b3836300                 add t2, t2, t1
    84: 07810302                 vle8.v v2, (t2)
    88: d780202e                 vxor.vv v1, v2, v1
    92: 83330501                 ld t2, 0x10(a0)
    96: b3836300                 add t2, t2, t1
   100: 07810302                 vle8.v v2, (t2)
   104: d780202e                 vxor.vv v1, v2, v1
   108: 83338501                 ld t2, 0x18(a0)
   112: b3836300                 add t2, t2, t1
   116: 07810302                 vle8.v v2, (t2)
   120: d780202e                 vxor.vv v1, v2, v1
   124: 83b30500                 ld t2, 0(a1)
   128: b3836300                 add t2, t2, t1
   132: a7800302                 vse8.v v1, (t2)
   136: 3303d300                 add t1, t1, a3
   140: 1307f7ff                 addi a4, a4, -1
   144: e31807fa                 bnez a4, 0x40
   148: 33066640                 sub a2, a2, t1
   152: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
   156: 93060600                 mv a3, a2
   160: 6344d000                 bgtz a3, 0xa8
   164: 6f004006                 j 0x108
   168: b3835640                 sub t2, a3, t0
   172: 13def343                 srai t3, t2, 0x3f
   176: b3f3c301                 and t2, t2, t3
   180: b3837200                 add t2, t0, t2
   184: 57a10852                 vid.v v2
   188: 57c0236a                 vmsltu.vx v0, v2, t2
   192: 83330500                 ld t2, 0(a0)
   196: b3836300                 add t2, t2, t1
   200: 87800300                 vle8.v v1, (t2), v0.t
   204: 83338500                 ld t2, 8(a0)
   208: b3836300                 add t2, t2, t1
   212: 07810300                 vle8.v v2, (t2), v0.t
   216: d780202e                 vxor.vv v1, v2, v1
   220: 83330501                 ld t2, 0x10(a0)
   224: b3836300                 add t2, t2, t1
   228: 07810300                 vle8.v v2, (t2), v0.t
   232: d780202e                 vxor.vv v1, v2, v1
   236: 83338501                 ld t2, 0x18(a0)
   240: b3836300                 add t2, t2, t1
   244: 07810300                 vle8.v v2, (t2), v0.t
   248: d780202e                 vxor.vv v1, v2, v1
   252: 83b30500                 ld t2, 0(a1)
   256: b3836300                 add t2, t2, t1
   260: a7800300                 vse8.v v1, (t2), v0.t
   264: 33035300                 add t1, t1, t0
   268: b3865640                 sub a3, a3, t0
   272: 67800000                 ret 
   276: 0000                     unimp 
   278: 0000                     unimp 
   280: 0000                     unimp 
   282: 0000                     unimp 
   284: 0000                     unimp 
   286: 0000                     unimp 
