; config_t1_758 target=sse twin=notail topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=896 twins=["notail"] uops_per_byte=8.4375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d67030000           leaq 0x367(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66450f6f07               movdqa (%r15), %xmm8
    52: 488bc1                   movq %rcx, %rax
    55: 4883e810                 subq $0x10, %rax
    59: 0f82ff020000             jb 0x340
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   102: 660f6fda                 movdqa %xmm2, %xmm3
   106: 66410fdbd8               pand %xmm8, %xmm3
   111: 660f6fe2                 movdqa %xmm2, %xmm4
   115: 660f71d404               psrlw $4, %xmm4
   120: 66410fdbe0               pand %xmm8, %xmm4
   125: 66410f6f7f40             movdqa 0x40(%r15), %xmm7
   131: 660f6fef                 movdqa %xmm7, %xmm5
   135: 660f3800eb               pshufb %xmm3, %xmm5
   140: 66410f6fbf80000000       movdqa 0x80(%r15), %xmm7
   149: 660f6ff7                 movdqa %xmm7, %xmm6
   153: 660f3800f4               pshufb %xmm4, %xmm6
   158: 660f6fc5                 movdqa %xmm5, %xmm0
   162: 660fefc6                 pxor %xmm6, %xmm0
   166: 66410f6fbfc0000000       movdqa 0xc0(%r15), %xmm7
   175: 660f6fef                 movdqa %xmm7, %xmm5
   179: 660f3800eb               pshufb %xmm3, %xmm5
   184: 66410f6fbf00010000       movdqa 0x100(%r15), %xmm7
   193: 660f6ff7                 movdqa %xmm7, %xmm6
   197: 660f3800f4               pshufb %xmm4, %xmm6
   202: 660f6fcd                 movdqa %xmm5, %xmm1
   206: 660fefce                 pxor %xmm6, %xmm1
   210: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
   216: 660f6fda                 movdqa %xmm2, %xmm3
   220: 66410fdbd8               pand %xmm8, %xmm3
   225: 660f6fe2                 movdqa %xmm2, %xmm4
   229: 660f71d404               psrlw $4, %xmm4
   234: 66410fdbe0               pand %xmm8, %xmm4
   239: 66410f6fbf40010000       movdqa 0x140(%r15), %xmm7
   248: 660f6fef                 movdqa %xmm7, %xmm5
   252: 660f3800eb               pshufb %xmm3, %xmm5
   257: 66410f6fbf80010000       movdqa 0x180(%r15), %xmm7
   266: 660f6ff7                 movdqa %xmm7, %xmm6
   270: 660f3800f4               pshufb %xmm4, %xmm6
   275: 660fefc5                 pxor %xmm5, %xmm0
   279: 660fefc6                 pxor %xmm6, %xmm0
   283: 66410f6fbfc0010000       movdqa 0x1c0(%r15), %xmm7
   292: 660f6fef                 movdqa %xmm7, %xmm5
   296: 660f3800eb               pshufb %xmm3, %xmm5
   301: 66410f6fbf00020000       movdqa 0x200(%r15), %xmm7
   310: 660f6ff7                 movdqa %xmm7, %xmm6
   314: 660f3800f4               pshufb %xmm4, %xmm6
   319: 660fefcd                 pxor %xmm5, %xmm1
   323: 660fefce                 pxor %xmm6, %xmm1
   327: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   333: 660f6fda                 movdqa %xmm2, %xmm3
   337: 66410fdbd8               pand %xmm8, %xmm3
   342: 660f6fe2                 movdqa %xmm2, %xmm4
   346: 660f71d404               psrlw $4, %xmm4
   351: 66410fdbe0               pand %xmm8, %xmm4
   356: 66410f6fbf40020000       movdqa 0x240(%r15), %xmm7
   365: 660f6fef                 movdqa %xmm7, %xmm5
   369: 660f3800eb               pshufb %xmm3, %xmm5
   374: 66410f6fbf80020000       movdqa 0x280(%r15), %xmm7
   383: 660f6ff7                 movdqa %xmm7, %xmm6
   387: 660f3800f4               pshufb %xmm4, %xmm6
   392: 660fefc5                 pxor %xmm5, %xmm0
   396: 660fefc6                 pxor %xmm6, %xmm0
   400: 66410f6fbfc0020000       movdqa 0x2c0(%r15), %xmm7
   409: 660f6fef                 movdqa %xmm7, %xmm5
   413: 660f3800eb               pshufb %xmm3, %xmm5
   418: 66410f6fbf00030000       movdqa 0x300(%r15), %xmm7
   427: 660f6ff7                 movdqa %xmm7, %xmm6
   431: 660f3800f4               pshufb %xmm4, %xmm6
   436: 660fefcd                 pxor %xmm5, %xmm1
   440: 660fefce                 pxor %xmm6, %xmm1
   444: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
   451: 660f6fda                 movdqa %xmm2, %xmm3
   455: 66410fdbd8               pand %xmm8, %xmm3
   460: 660f6fe2                 movdqa %xmm2, %xmm4
   464: 660f71d404               psrlw $4, %xmm4
   469: 66410fdbe0               pand %xmm8, %xmm4
   474: 66410f6fbf40030000       movdqa 0x340(%r15), %xmm7
   483: 660f6fef                 movdqa %xmm7, %xmm5
   487: 660f3800eb               pshufb %xmm3, %xmm5
   492: 66410f6fbf80030000       movdqa 0x380(%r15), %xmm7
   501: 660f6ff7                 movdqa %xmm7, %xmm6
   505: 660f3800f4               pshufb %xmm4, %xmm6
   510: 660fefc5                 pxor %xmm5, %xmm0
   514: 660fefc6                 pxor %xmm6, %xmm0
   518: 66410f6fbfc0030000       movdqa 0x3c0(%r15), %xmm7
   527: 660f6fef                 movdqa %xmm7, %xmm5
   531: 660f3800eb               pshufb %xmm3, %xmm5
   536: 66410f6fbf00040000       movdqa 0x400(%r15), %xmm7
   545: 660f6ff7                 movdqa %xmm7, %xmm6
   549: 660f3800f4               pshufb %xmm4, %xmm6
   554: 660fefcd                 pxor %xmm5, %xmm1
   558: 660fefce                 pxor %xmm6, %xmm1
   562: 66430f6f141e             movdqa (%r14, %r11), %xmm2
   568: 660f6fda                 movdqa %xmm2, %xmm3
   572: 66410fdbd8               pand %xmm8, %xmm3
   577: 660f6fe2                 movdqa %xmm2, %xmm4
   581: 660f71d404               psrlw $4, %xmm4
   586: 66410fdbe0               pand %xmm8, %xmm4
   591: 66410f6fbf40040000       movdqa 0x440(%r15), %xmm7
   600: 660f6fef                 movdqa %xmm7, %xmm5
   604: 660f3800eb               pshufb %xmm3, %xmm5
   609: 66410f6fbf80040000       movdqa 0x480(%r15), %xmm7
   618: 660f6ff7                 movdqa %xmm7, %xmm6
   622: 660f3800f4               pshufb %xmm4, %xmm6
   627: 660fefc5                 pxor %xmm5, %xmm0
   631: 660fefc6                 pxor %xmm6, %xmm0
   635: 66410f6fbfc0040000       movdqa 0x4c0(%r15), %xmm7
   644: 660f6fef                 movdqa %xmm7, %xmm5
   648: 660f3800eb               pshufb %xmm3, %xmm5
   653: 66410f6fbf00050000       movdqa 0x500(%r15), %xmm7
   662: 660f6ff7                 movdqa %xmm7, %xmm6
   666: 660f3800f4               pshufb %xmm4, %xmm6
   671: 660fefcd                 pxor %xmm5, %xmm1
   675: 660fefce                 pxor %xmm6, %xmm1
   679: 4d8b5028                 movq 0x28(%r8), %r10
   683: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   689: 660f6fda                 movdqa %xmm2, %xmm3
   693: 66410fdbd8               pand %xmm8, %xmm3
   698: 660f6fe2                 movdqa %xmm2, %xmm4
   702: 660f71d404               psrlw $4, %xmm4
   707: 66410fdbe0               pand %xmm8, %xmm4
   712: 66410f6fbf40050000       movdqa 0x540(%r15), %xmm7
   721: 660f6fef                 movdqa %xmm7, %xmm5
   725: 660f3800eb               pshufb %xmm3, %xmm5
   730: 66410f6fbf80050000       movdqa 0x580(%r15), %xmm7
   739: 660f6ff7                 movdqa %xmm7, %xmm6
   743: 660f3800f4               pshufb %xmm4, %xmm6
   748: 660fefc5                 pxor %xmm5, %xmm0
   752: 660fefc6                 pxor %xmm6, %xmm0
   756: 66410f6fbf40050000       movdqa 0x540(%r15), %xmm7
   765: 660f6fef                 movdqa %xmm7, %xmm5
   769: 660f3800eb               pshufb %xmm3, %xmm5
   774: 66410f6fbf80050000       movdqa 0x580(%r15), %xmm7
   783: 660f6ff7                 movdqa %xmm7, %xmm6
   787: 660f3800f4               pshufb %xmm4, %xmm6
   792: 660fefcd                 pxor %xmm5, %xmm1
   796: 660fefce                 pxor %xmm6, %xmm1
   800: 4d8b11                   movq (%r9), %r10
   803: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   809: 4d8b5108                 movq 8(%r9), %r10
   813: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
   819: 4983c310                 addq $0x10, %r11
   823: 4c3bd8                   cmpq %rax, %r11
   826: 0f8620fdffff             jbe 0x60
   832: 415e                     popq %r14
   834: 415d                     popq %r13
   836: 415c                     popq %r12
   838: 415f                     popq %r15
   840: 5b                       popq %rbx
   841: c3                       retq 
   842: 0f0b                     ud2 
   844: 0f0b                     ud2 
   846: 0f0b                     ud2 
   848: 0f0b                     ud2 
   850: 0f0b                     ud2 
   852: 0f0b                     ud2 
   854: 0f0b                     ud2 
   856: 0f0b                     ud2 
   858: 0f0b                     ud2 
   860: 0f0b                     ud2 
   862: 0f0b                     ud2 
   864: 0f0b                     ud2 
   866: 0f0b                     ud2 
   868: 0f0b                     ud2 
   870: 0f0b                     ud2 
   872: 0f0b                     ud2 
   874: 0f0b                     ud2 
   876: 0f0b                     ud2 
   878: 0f0b                     ud2 
   880: 0f0b                     ud2 
   882: 0f0b                     ud2 
   884: 0f0b                     ud2 
   886: 0f0b                     ud2 
   888: 0f0b                     ud2 
   890: 0f0b                     ud2 
   892: 0f0b                     ud2 
   894: 0f0b                     ud2 
