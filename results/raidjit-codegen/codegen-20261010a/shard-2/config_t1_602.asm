; config_t1_602 target=sse twin=notail topo=raid6(8)/enc op=encode U=1
; code_size=704 twins=["notail"] uops_per_byte=7.0625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da7020000           leaq 0x2a7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66450f6f07               movdqa (%r15), %xmm8
    52: 488bc1                   movq %rcx, %rax
    55: 4883e810                 subq $0x10, %rax
    59: 0f8272020000             jb 0x2b3
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   102: 660f6fda                 movdqa %xmm2, %xmm3
   106: 66410fdbd8               pand %xmm8, %xmm3
   111: 660f6fe2                 movdqa %xmm2, %xmm4
   115: 660f71d404               psrlw $4, %xmm4
   120: 66410fdbe0               pand %xmm8, %xmm4
   125: 660f6fc2                 movdqa %xmm2, %xmm0
   129: 66410f6f7f40             movdqa 0x40(%r15), %xmm7
   135: 660f6fef                 movdqa %xmm7, %xmm5
   139: 660f3800eb               pshufb %xmm3, %xmm5
   144: 66410f6fbf80000000       movdqa 0x80(%r15), %xmm7
   153: 660f6ff7                 movdqa %xmm7, %xmm6
   157: 660f3800f4               pshufb %xmm4, %xmm6
   162: 660f6fcd                 movdqa %xmm5, %xmm1
   166: 660fefce                 pxor %xmm6, %xmm1
   170: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
   176: 660f6fda                 movdqa %xmm2, %xmm3
   180: 66410fdbd8               pand %xmm8, %xmm3
   185: 660f6fe2                 movdqa %xmm2, %xmm4
   189: 660f71d404               psrlw $4, %xmm4
   194: 66410fdbe0               pand %xmm8, %xmm4
   199: 660fefc2                 pxor %xmm2, %xmm0
   203: 66410f6fbfc0000000       movdqa 0xc0(%r15), %xmm7
   212: 660f6fef                 movdqa %xmm7, %xmm5
   216: 660f3800eb               pshufb %xmm3, %xmm5
   221: 66410f6fbf00010000       movdqa 0x100(%r15), %xmm7
   230: 660f6ff7                 movdqa %xmm7, %xmm6
   234: 660f3800f4               pshufb %xmm4, %xmm6
   239: 660fefcd                 pxor %xmm5, %xmm1
   243: 660fefce                 pxor %xmm6, %xmm1
   247: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   253: 660f6fda                 movdqa %xmm2, %xmm3
   257: 66410fdbd8               pand %xmm8, %xmm3
   262: 660f6fe2                 movdqa %xmm2, %xmm4
   266: 660f71d404               psrlw $4, %xmm4
   271: 66410fdbe0               pand %xmm8, %xmm4
   276: 660fefc2                 pxor %xmm2, %xmm0
   280: 66410f6fbf40010000       movdqa 0x140(%r15), %xmm7
   289: 660f6fef                 movdqa %xmm7, %xmm5
   293: 660f3800eb               pshufb %xmm3, %xmm5
   298: 66410f6fbf80010000       movdqa 0x180(%r15), %xmm7
   307: 660f6ff7                 movdqa %xmm7, %xmm6
   311: 660f3800f4               pshufb %xmm4, %xmm6
   316: 660fefcd                 pxor %xmm5, %xmm1
   320: 660fefce                 pxor %xmm6, %xmm1
   324: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
   331: 660f6fda                 movdqa %xmm2, %xmm3
   335: 66410fdbd8               pand %xmm8, %xmm3
   340: 660f6fe2                 movdqa %xmm2, %xmm4
   344: 660f71d404               psrlw $4, %xmm4
   349: 66410fdbe0               pand %xmm8, %xmm4
   354: 660fefc2                 pxor %xmm2, %xmm0
   358: 66410f6fbfc0010000       movdqa 0x1c0(%r15), %xmm7
   367: 660f6fef                 movdqa %xmm7, %xmm5
   371: 660f3800eb               pshufb %xmm3, %xmm5
   376: 66410f6fbf00020000       movdqa 0x200(%r15), %xmm7
   385: 660f6ff7                 movdqa %xmm7, %xmm6
   389: 660f3800f4               pshufb %xmm4, %xmm6
   394: 660fefcd                 pxor %xmm5, %xmm1
   398: 660fefce                 pxor %xmm6, %xmm1
   402: 66430f6f141e             movdqa (%r14, %r11), %xmm2
   408: 660f6fda                 movdqa %xmm2, %xmm3
   412: 66410fdbd8               pand %xmm8, %xmm3
   417: 660f6fe2                 movdqa %xmm2, %xmm4
   421: 660f71d404               psrlw $4, %xmm4
   426: 66410fdbe0               pand %xmm8, %xmm4
   431: 660fefc2                 pxor %xmm2, %xmm0
   435: 66410f6fbf40020000       movdqa 0x240(%r15), %xmm7
   444: 660f6fef                 movdqa %xmm7, %xmm5
   448: 660f3800eb               pshufb %xmm3, %xmm5
   453: 66410f6fbf80020000       movdqa 0x280(%r15), %xmm7
   462: 660f6ff7                 movdqa %xmm7, %xmm6
   466: 660f3800f4               pshufb %xmm4, %xmm6
   471: 660fefcd                 pxor %xmm5, %xmm1
   475: 660fefce                 pxor %xmm6, %xmm1
   479: 4d8b5028                 movq 0x28(%r8), %r10
   483: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   489: 660f6fda                 movdqa %xmm2, %xmm3
   493: 66410fdbd8               pand %xmm8, %xmm3
   498: 660f6fe2                 movdqa %xmm2, %xmm4
   502: 660f71d404               psrlw $4, %xmm4
   507: 66410fdbe0               pand %xmm8, %xmm4
   512: 660fefc2                 pxor %xmm2, %xmm0
   516: 66410f6fbfc0020000       movdqa 0x2c0(%r15), %xmm7
   525: 660f6fef                 movdqa %xmm7, %xmm5
   529: 660f3800eb               pshufb %xmm3, %xmm5
   534: 66410f6fbf00030000       movdqa 0x300(%r15), %xmm7
   543: 660f6ff7                 movdqa %xmm7, %xmm6
   547: 660f3800f4               pshufb %xmm4, %xmm6
   552: 660fefcd                 pxor %xmm5, %xmm1
   556: 660fefce                 pxor %xmm6, %xmm1
   560: 4d8b5030                 movq 0x30(%r8), %r10
   564: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   570: 660f6fda                 movdqa %xmm2, %xmm3
   574: 66410fdbd8               pand %xmm8, %xmm3
   579: 660f6fe2                 movdqa %xmm2, %xmm4
   583: 660f71d404               psrlw $4, %xmm4
   588: 66410fdbe0               pand %xmm8, %xmm4
   593: 660fefc2                 pxor %xmm2, %xmm0
   597: 66410f6fbf40030000       movdqa 0x340(%r15), %xmm7
   606: 660f6fef                 movdqa %xmm7, %xmm5
   610: 660f3800eb               pshufb %xmm3, %xmm5
   615: 66410f6fbf80030000       movdqa 0x380(%r15), %xmm7
   624: 660f6ff7                 movdqa %xmm7, %xmm6
   628: 660f3800f4               pshufb %xmm4, %xmm6
   633: 660fefcd                 pxor %xmm5, %xmm1
   637: 660fefce                 pxor %xmm6, %xmm1
   641: 4d8b5038                 movq 0x38(%r8), %r10
   645: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   651: 660fefc2                 pxor %xmm2, %xmm0
   655: 660fefca                 pxor %xmm2, %xmm1
   659: 4d8b11                   movq (%r9), %r10
   662: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   668: 4d8b5108                 movq 8(%r9), %r10
   672: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
   678: 4983c310                 addq $0x10, %r11
   682: 4c3bd8                   cmpq %rax, %r11
   685: 0f86adfdffff             jbe 0x60
   691: 415e                     popq %r14
   693: 415d                     popq %r13
   695: 415c                     popq %r12
   697: 415f                     popq %r15
   699: 5b                       popq %rbx
   700: c3                       retq 
   701: 0f0b                     ud2 
   703: cc                       int3 
