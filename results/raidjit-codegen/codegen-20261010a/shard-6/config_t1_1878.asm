; config_t1_1878 target=rvv-zvbc32e twin=maskedtail topo=rs(6,3)/rec{0} op=recover U=1
; code_size=352 twins=["maskedtail"] uops_per_byte=2.1875
     0: 63140600                 bnez a2, 8
     4: 6f00c014                 j 0x150
     8: 970f0000                 auipc t6, 0
    12: 938f8f15                 addi t6, t6, 0x158
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93960200                 slli a3, t0, 0
    32: 3357d602                 divu a4, a2, a3
    36: 63140700                 bnez a4, 0x2c
    40: 6f000009                 j 0xb8
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
   124: 83330502                 ld t2, 0x20(a0)
   128: b3836300                 add t2, t2, t1
   132: 07810302                 vle8.v v2, (t2)
   136: d780202e                 vxor.vv v1, v2, v1
   140: 83338502                 ld t2, 0x28(a0)
   144: b3836300                 add t2, t2, t1
   148: 07810302                 vle8.v v2, (t2)
   152: d780202e                 vxor.vv v1, v2, v1
   156: 83b30500                 ld t2, 0(a1)
   160: b3836300                 add t2, t2, t1
   164: a7800302                 vse8.v v1, (t2)
   168: 3303d300                 add t1, t1, a3
   172: 1307f7ff                 addi a4, a4, -1
   176: e31807f8                 bnez a4, 0x40
   180: 33066640                 sub a2, a2, t1
   184: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
   188: 93060600                 mv a3, a2
   192: 6344d000                 bgtz a3, 0xc8
   196: 6f004008                 j 0x148
   200: b3835640                 sub t2, a3, t0
   204: 13def343                 srai t3, t2, 0x3f
   208: b3f3c301                 and t2, t2, t3
   212: b3837200                 add t2, t0, t2
   216: 57a10852                 vid.v v2
   220: 57c0236a                 vmsltu.vx v0, v2, t2
   224: 83330500                 ld t2, 0(a0)
   228: b3836300                 add t2, t2, t1
   232: 87800300                 vle8.v v1, (t2), v0.t
   236: 83338500                 ld t2, 8(a0)
   240: b3836300                 add t2, t2, t1
   244: 07810300                 vle8.v v2, (t2), v0.t
   248: d780202e                 vxor.vv v1, v2, v1
   252: 83330501                 ld t2, 0x10(a0)
   256: b3836300                 add t2, t2, t1
   260: 07810300                 vle8.v v2, (t2), v0.t
   264: d780202e                 vxor.vv v1, v2, v1
   268: 83338501                 ld t2, 0x18(a0)
   272: b3836300                 add t2, t2, t1
   276: 07810300                 vle8.v v2, (t2), v0.t
   280: d780202e                 vxor.vv v1, v2, v1
   284: 83330502                 ld t2, 0x20(a0)
   288: b3836300                 add t2, t2, t1
   292: 07810300                 vle8.v v2, (t2), v0.t
   296: d780202e                 vxor.vv v1, v2, v1
   300: 83338502                 ld t2, 0x28(a0)
   304: b3836300                 add t2, t2, t1
   308: 07810300                 vle8.v v2, (t2), v0.t
   312: d780202e                 vxor.vv v1, v2, v1
   316: 83b30500                 ld t2, 0(a1)
   320: b3836300                 add t2, t2, t1
   324: a7800300                 vse8.v v1, (t2), v0.t
   328: 33035300                 add t1, t1, t0
   332: b3865640                 sub a3, a3, t0
   336: 67800000                 ret 
   340: 0000                     unimp 
   342: 0000                     unimp 
   344: 0000                     unimp 
   346: 0000                     unimp 
   348: 0000                     unimp 
   350: 0000                     unimp 
