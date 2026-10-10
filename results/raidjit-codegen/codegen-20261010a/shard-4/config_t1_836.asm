; config_t1_836 target=sse twin=notail topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=320 twins=["notail"] uops_per_byte=2.375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3d2d010000           leaq 0x12d(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: 66450f6f07               movdqa (%r15), %xmm8
    27: 488bc1                   movq %rcx, %rax
    30: 4883e810                 subq $0x10, %rax
    34: 0f8206010000             jb 0x12e
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5020                 movq 0x20(%r8), %r10
    68: 66430f6f041a             movdqa (%r10, %r11), %xmm0
    74: 4d8b5028                 movq 0x28(%r8), %r10
    78: 66430f6f0c1a             movdqa (%r10, %r11), %xmm1
    84: 4d8b10                   movq (%r8), %r10
    87: 498b5010                 movq 0x10(%r8), %rdx
    91: 66430f6f141a             movdqa (%r10, %r11), %xmm2
    97: 66420f6f2c1a             movdqa (%rdx, %r11), %xmm5
   103: 660fefd5                 pxor %xmm5, %xmm2
   107: 660f6fda                 movdqa %xmm2, %xmm3
   111: 66410fdbd8               pand %xmm8, %xmm3
   116: 660f6fe2                 movdqa %xmm2, %xmm4
   120: 660f71d404               psrlw $4, %xmm4
   125: 66410fdbe0               pand %xmm8, %xmm4
   130: 660fefc2                 pxor %xmm2, %xmm0
   134: 66410f6f7f40             movdqa 0x40(%r15), %xmm7
   140: 660f6fef                 movdqa %xmm7, %xmm5
   144: 660f3800eb               pshufb %xmm3, %xmm5
   149: 66410f6fbf80000000       movdqa 0x80(%r15), %xmm7
   158: 660f6ff7                 movdqa %xmm7, %xmm6
   162: 660f3800f4               pshufb %xmm4, %xmm6
   167: 660fefcd                 pxor %xmm5, %xmm1
   171: 660fefce                 pxor %xmm6, %xmm1
   175: 4d8b5008                 movq 8(%r8), %r10
   179: 498b5018                 movq 0x18(%r8), %rdx
   183: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   189: 66420f6f2c1a             movdqa (%rdx, %r11), %xmm5
   195: 660fefd5                 pxor %xmm5, %xmm2
   199: 660f6fda                 movdqa %xmm2, %xmm3
   203: 66410fdbd8               pand %xmm8, %xmm3
   208: 660f6fe2                 movdqa %xmm2, %xmm4
   212: 660f71d404               psrlw $4, %xmm4
   217: 66410fdbe0               pand %xmm8, %xmm4
   222: 660fefc2                 pxor %xmm2, %xmm0
   226: 66410f6fbfc0000000       movdqa 0xc0(%r15), %xmm7
   235: 660f6fef                 movdqa %xmm7, %xmm5
   239: 660f3800eb               pshufb %xmm3, %xmm5
   244: 66410f6fbf00010000       movdqa 0x100(%r15), %xmm7
   253: 660f6ff7                 movdqa %xmm7, %xmm6
   257: 660f3800f4               pshufb %xmm4, %xmm6
   262: 660fefcd                 pxor %xmm5, %xmm1
   266: 660fefce                 pxor %xmm6, %xmm1
   270: 4d8b11                   movq (%r9), %r10
   273: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   279: 4d8b5108                 movq 8(%r9), %r10
   283: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
   289: 4983c310                 addq $0x10, %r11
   293: 4c3bd8                   cmpq %rax, %r11
   296: 0f8612ffffff             jbe 0x40
   302: 415f                     popq %r15
   304: 5b                       popq %rbx
   305: c3                       retq 
   306: 0f0b                     ud2 
   308: 0f0b                     ud2 
   310: 0f0b                     ud2 
   312: 0f0b                     ud2 
   314: 0f0b                     ud2 
   316: 0f0b                     ud2 
   318: 0f0b                     ud2 
