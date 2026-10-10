; config_t3_2090 target=rvv twin=baseline topo=raid1()/enc op=encode U=8
; code_size=272 twins=[] uops_per_byte=0.15625
     0: 63140600                 bnez a2, 8
     4: 6f004010                 j 0x108
     8: 970f0000                 auipc t6, 0
    12: 938f8f10                 addi t6, t6, 0x108
    16: 577f000c                 vsetvli t5, zero, e8, m1, ta, ma
    20: 13030000                 li t1, 0
    24: d772000c                 vsetvli t0, zero, e8, m1, ta, ma
    28: 93068000                 li a3, 8
    32: b386d202                 mul a3, t0, a3
    36: 3357d602                 divu a4, a2, a3
    40: 63140700                 bnez a4, 0x30
    44: 6f00c00a                 j 0xd8
    48: 13000000                 nop 
    52: 13000000                 nop 
    56: 13000000                 nop 
    60: 13000000                 nop 
    64: 83330500                 ld t2, 0(a0)
    68: b3836300                 add t2, t2, t1
    72: 07800302                 vle8.v v0, (t2)
    76: b3835300                 add t2, t2, t0
    80: 87800302                 vle8.v v1, (t2)
    84: b3835300                 add t2, t2, t0
    88: 07810302                 vle8.v v2, (t2)
    92: b3835300                 add t2, t2, t0
    96: 87810302                 vle8.v v3, (t2)
   100: b3835300                 add t2, t2, t0
   104: 07820302                 vle8.v v4, (t2)
   108: b3835300                 add t2, t2, t0
   112: 87820302                 vle8.v v5, (t2)
   116: b3835300                 add t2, t2, t0
   120: 07830302                 vle8.v v6, (t2)
   124: b3835300                 add t2, t2, t0
   128: 87830302                 vle8.v v7, (t2)
   132: 83b30500                 ld t2, 0(a1)
   136: b3836300                 add t2, t2, t1
   140: 27800302                 vse8.v v0, (t2)
   144: b3835300                 add t2, t2, t0
   148: a7800302                 vse8.v v1, (t2)
   152: b3835300                 add t2, t2, t0
   156: 27810302                 vse8.v v2, (t2)
   160: b3835300                 add t2, t2, t0
   164: a7810302                 vse8.v v3, (t2)
   168: b3835300                 add t2, t2, t0
   172: 27820302                 vse8.v v4, (t2)
   176: b3835300                 add t2, t2, t0
   180: a7820302                 vse8.v v5, (t2)
   184: b3835300                 add t2, t2, t0
   188: 27830302                 vse8.v v6, (t2)
   192: b3835300                 add t2, t2, t0
   196: a7830302                 vse8.v v7, (t2)
   200: 3303d300                 add t1, t1, a3
   204: 1307f7ff                 addi a4, a4, -1
   208: e31807f6                 bnez a4, 0x40
   212: 33066640                 sub a2, a2, t1
   216: d772060c                 vsetvli t0, a2, e8, m1, ta, ma
   220: 63940200                 bnez t0, 0xe4
   224: 6f008002                 j 0x108
   228: 83330500                 ld t2, 0(a0)
   232: b3836300                 add t2, t2, t1
   236: 07800302                 vle8.v v0, (t2)
   240: 83b30500                 ld t2, 0(a1)
   244: b3836300                 add t2, t2, t1
   248: 27800302                 vse8.v v0, (t2)
   252: 33035300                 add t1, t1, t0
   256: 33065640                 sub a2, a2, t0
   260: 6ff05ffd                 j 0xd8
   264: 67800000                 ret 
   268: 0000                     unimp 
   270: 0000                     unimp 
