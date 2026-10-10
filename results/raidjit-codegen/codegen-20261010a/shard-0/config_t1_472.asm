; config_t1_472 target=avx2 twin=notail topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=704 twins=["notail"] uops_per_byte=3.03125
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
    47: c4417d6f07               vmovdqa (%r15), %ymm8
    52: 488bc1                   movq %rcx, %rax
    55: 4883e820                 subq $0x20, %rax
    59: 0f8267020000             jb 0x2a8
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   102: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   107: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   112: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   117: c4c17d6f7f40             vmovdqa 0x40(%r15), %ymm7
   123: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   128: c4c17d6fbf80000000       vmovdqa 0x80(%r15), %ymm7
   137: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   142: c5d5efc6                 vpxor %ymm6, %ymm5, %ymm0
   146: c4c17d6fbfc0000000       vmovdqa 0xc0(%r15), %ymm7
   155: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   160: c4c17d6fbf00010000       vmovdqa 0x100(%r15), %ymm7
   169: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   174: c5d5efce                 vpxor %ymm6, %ymm5, %ymm1
   178: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   184: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   189: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   194: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   199: c4c17d6fbf40010000       vmovdqa 0x140(%r15), %ymm7
   208: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   213: c4c17d6fbf80010000       vmovdqa 0x180(%r15), %ymm7
   222: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   227: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   231: c5fdefc6                 vpxor %ymm6, %ymm0, %ymm0
   235: c4c17d6fbfc0010000       vmovdqa 0x1c0(%r15), %ymm7
   244: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   249: c4c17d6fbf00020000       vmovdqa 0x200(%r15), %ymm7
   258: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   263: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   267: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   271: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   277: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   282: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   287: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   292: c4c17d6fbf40020000       vmovdqa 0x240(%r15), %ymm7
   301: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   306: c4c17d6fbf80020000       vmovdqa 0x280(%r15), %ymm7
   315: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   320: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   324: c5fdefc6                 vpxor %ymm6, %ymm0, %ymm0
   328: c4c17d6fbfc0020000       vmovdqa 0x2c0(%r15), %ymm7
   337: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   342: c4c17d6fbf00030000       vmovdqa 0x300(%r15), %ymm7
   351: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   356: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   360: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   364: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   371: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   376: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   381: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   386: c4c17d6fbf40030000       vmovdqa 0x340(%r15), %ymm7
   395: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   400: c4c17d6fbf80030000       vmovdqa 0x380(%r15), %ymm7
   409: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   414: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   418: c5fdefc6                 vpxor %ymm6, %ymm0, %ymm0
   422: c4c17d6fbfc0030000       vmovdqa 0x3c0(%r15), %ymm7
   431: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   436: c4c17d6fbf00040000       vmovdqa 0x400(%r15), %ymm7
   445: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   450: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   454: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   458: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   464: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   469: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   474: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   479: c4c17d6fbf40040000       vmovdqa 0x440(%r15), %ymm7
   488: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   493: c4c17d6fbf80040000       vmovdqa 0x480(%r15), %ymm7
   502: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   507: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   511: c5fdefc6                 vpxor %ymm6, %ymm0, %ymm0
   515: c4c17d6fbfc0040000       vmovdqa 0x4c0(%r15), %ymm7
   524: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   529: c4c17d6fbf00050000       vmovdqa 0x500(%r15), %ymm7
   538: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   543: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   547: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   551: 4d8b5028                 movq 0x28(%r8), %r10
   555: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   561: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   566: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   571: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   576: c4c17d6fbf40050000       vmovdqa 0x540(%r15), %ymm7
   585: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   590: c4c17d6fbf80050000       vmovdqa 0x580(%r15), %ymm7
   599: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   604: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   608: c5fdefc6                 vpxor %ymm6, %ymm0, %ymm0
   612: c4c17d6fbf40050000       vmovdqa 0x540(%r15), %ymm7
   621: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   626: c4c17d6fbf80050000       vmovdqa 0x580(%r15), %ymm7
   635: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   640: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   644: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   648: 4d8b11                   movq (%r9), %r10
   651: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   657: 4d8b5108                 movq 8(%r9), %r10
   661: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
   667: 4983c320                 addq $0x20, %r11
   671: 4c3bd8                   cmpq %rax, %r11
   674: 0f86b8fdffff             jbe 0x60
   680: 415e                     popq %r14
   682: 415d                     popq %r13
   684: 415c                     popq %r12
   686: 415f                     popq %r15
   688: 5b                       popq %rbx
   689: c5f877                   vzeroupper 
   692: c3                       retq 
   693: 0f0b                     ud2 
   695: 0f0b                     ud2 
   697: 0f0b                     ud2 
   699: 0f0b                     ud2 
   701: 0f0b                     ud2 
   703: cc                       int3 
