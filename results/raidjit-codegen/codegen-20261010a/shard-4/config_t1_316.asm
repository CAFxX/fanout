; config_t1_316 target=avx2 twin=notail topo=raid6(8)/enc op=encode U=1
; code_size=640 twins=["notail"] uops_per_byte=2.625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d67020000           leaq 0x267(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: c4417d6f07               vmovdqa (%r15), %ymm8
    52: 488bc1                   movq %rcx, %rax
    55: 4883e820                 subq $0x20, %rax
    59: 0f82fe010000             jb 0x23f
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   102: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   107: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   112: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   117: c5fd6fc2                 vmovdqa %ymm2, %ymm0
   121: c4c17d6f7f40             vmovdqa 0x40(%r15), %ymm7
   127: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   132: c4c17d6fbf80000000       vmovdqa 0x80(%r15), %ymm7
   141: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   146: c5d5efce                 vpxor %ymm6, %ymm5, %ymm1
   150: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   156: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   161: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   166: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   171: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   175: c4c17d6fbfc0000000       vmovdqa 0xc0(%r15), %ymm7
   184: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   189: c4c17d6fbf00010000       vmovdqa 0x100(%r15), %ymm7
   198: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   203: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   207: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   211: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   217: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   222: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   227: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   232: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   236: c4c17d6fbf40010000       vmovdqa 0x140(%r15), %ymm7
   245: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   250: c4c17d6fbf80010000       vmovdqa 0x180(%r15), %ymm7
   259: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   264: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   268: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   272: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   279: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   284: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   289: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   294: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   298: c4c17d6fbfc0010000       vmovdqa 0x1c0(%r15), %ymm7
   307: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   312: c4c17d6fbf00020000       vmovdqa 0x200(%r15), %ymm7
   321: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   326: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   330: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   334: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   340: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   345: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   350: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   355: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   359: c4c17d6fbf40020000       vmovdqa 0x240(%r15), %ymm7
   368: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   373: c4c17d6fbf80020000       vmovdqa 0x280(%r15), %ymm7
   382: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   387: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   391: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   395: 4d8b5028                 movq 0x28(%r8), %r10
   399: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   405: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   410: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   415: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   420: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   424: c4c17d6fbfc0020000       vmovdqa 0x2c0(%r15), %ymm7
   433: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   438: c4c17d6fbf00030000       vmovdqa 0x300(%r15), %ymm7
   447: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   452: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   456: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   460: 4d8b5030                 movq 0x30(%r8), %r10
   464: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   470: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   475: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   480: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   485: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   489: c4c17d6fbf40030000       vmovdqa 0x340(%r15), %ymm7
   498: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   503: c4c17d6fbf80030000       vmovdqa 0x380(%r15), %ymm7
   512: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   517: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   521: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   525: 4d8b5038                 movq 0x38(%r8), %r10
   529: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   535: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   539: c5f5efca                 vpxor %ymm2, %ymm1, %ymm1
   543: 4d8b11                   movq (%r9), %r10
   546: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   552: 4d8b5108                 movq 8(%r9), %r10
   556: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
   562: 4983c320                 addq $0x20, %r11
   566: 4c3bd8                   cmpq %rax, %r11
   569: 0f8621feffff             jbe 0x60
   575: 415e                     popq %r14
   577: 415d                     popq %r13
   579: 415c                     popq %r12
   581: 415f                     popq %r15
   583: 5b                       popq %rbx
   584: c5f877                   vzeroupper 
   587: c3                       retq 
   588: 0f0b                     ud2 
   590: 0f0b                     ud2 
   592: 0f0b                     ud2 
   594: 0f0b                     ud2 
   596: 0f0b                     ud2 
   598: 0f0b                     ud2 
   600: 0f0b                     ud2 
   602: 0f0b                     ud2 
   604: 0f0b                     ud2 
   606: 0f0b                     ud2 
   608: 0f0b                     ud2 
   610: 0f0b                     ud2 
   612: 0f0b                     ud2 
   614: 0f0b                     ud2 
   616: 0f0b                     ud2 
   618: 0f0b                     ud2 
   620: 0f0b                     ud2 
   622: 0f0b                     ud2 
   624: 0f0b                     ud2 
   626: 0f0b                     ud2 
   628: 0f0b                     ud2 
   630: 0f0b                     ud2 
   632: 0f0b                     ud2 
   634: 0f0b                     ud2 
   636: 0f0b                     ud2 
   638: 0f0b                     ud2 
