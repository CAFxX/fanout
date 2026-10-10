; config_t1_368 target=avx2 twin=notail topo=rs(6,3)/enc op=encode U=1
; code_size=704 twins=["notail"] uops_per_byte=2.84375
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
    47: c4417d6f0f               vmovdqa (%r15), %ymm9
    52: 488bc1                   movq %rcx, %rax
    55: 4883e820                 subq $0x20, %rax
    59: 0f8237020000             jb 0x278
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4a17d6f1c1e             vmovdqa (%rsi, %r11), %ymm3
   102: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   107: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   112: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   117: c5fd6fc3                 vmovdqa %ymm3, %ymm0
   121: c4417d6f4740             vmovdqa 0x40(%r15), %ymm8
   127: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   132: c4417d6f8780000000       vmovdqa 0x80(%r15), %ymm8
   141: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   146: c5cdefcf                 vpxor %ymm7, %ymm6, %ymm1
   150: c4417d6f87c0000000       vmovdqa 0xc0(%r15), %ymm8
   159: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   164: c4417d6f8700010000       vmovdqa 0x100(%r15), %ymm8
   173: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   178: c5cdefd7                 vpxor %ymm7, %ymm6, %ymm2
   182: c4a17d6f1c1f             vmovdqa (%rdi, %r11), %ymm3
   188: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   193: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   198: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   203: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   207: c4417d6f87c0000000       vmovdqa 0xc0(%r15), %ymm8
   216: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   221: c4417d6f8700010000       vmovdqa 0x100(%r15), %ymm8
   230: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   235: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   239: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   243: c4417d6f8740010000       vmovdqa 0x140(%r15), %ymm8
   252: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   257: c4417d6f8780010000       vmovdqa 0x180(%r15), %ymm8
   266: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   271: c5edefd6                 vpxor %ymm6, %ymm2, %ymm2
   275: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   279: c4817d6f1c1c             vmovdqa (%r12, %r11), %ymm3
   285: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   290: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   295: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   300: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   304: c4417d6f8740010000       vmovdqa 0x140(%r15), %ymm8
   313: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   318: c4417d6f8780010000       vmovdqa 0x180(%r15), %ymm8
   327: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   332: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   336: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   340: c4417d6f87c0010000       vmovdqa 0x1c0(%r15), %ymm8
   349: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   354: c4417d6f8700020000       vmovdqa 0x200(%r15), %ymm8
   363: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   368: c5edefd6                 vpxor %ymm6, %ymm2, %ymm2
   372: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   376: c4817d6f5c1d00           vmovdqa (%r13, %r11), %ymm3
   383: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   388: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   393: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   398: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   402: c4417d6f8740020000       vmovdqa 0x240(%r15), %ymm8
   411: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   416: c4417d6f8780020000       vmovdqa 0x280(%r15), %ymm8
   425: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   430: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   434: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   438: c4417d6f4740             vmovdqa 0x40(%r15), %ymm8
   444: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   449: c4417d6f8780000000       vmovdqa 0x80(%r15), %ymm8
   458: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   463: c5edefd6                 vpxor %ymm6, %ymm2, %ymm2
   467: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   471: c4817d6f1c1e             vmovdqa (%r14, %r11), %ymm3
   477: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   482: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   487: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   492: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   496: c4417d6f87c0010000       vmovdqa 0x1c0(%r15), %ymm8
   505: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   510: c4417d6f8700020000       vmovdqa 0x200(%r15), %ymm8
   519: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   524: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   528: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   532: c4417d6f8740020000       vmovdqa 0x240(%r15), %ymm8
   541: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   546: c4417d6f8780020000       vmovdqa 0x280(%r15), %ymm8
   555: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   560: c5edefd6                 vpxor %ymm6, %ymm2, %ymm2
   564: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   568: 4d8b5028                 movq 0x28(%r8), %r10
   572: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   578: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   582: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   586: c5edefd3                 vpxor %ymm3, %ymm2, %ymm2
   590: 4d8b11                   movq (%r9), %r10
   593: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   599: 4d8b5108                 movq 8(%r9), %r10
   603: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
   609: 4d8b5110                 movq 0x10(%r9), %r10
   613: c4817d7f141a             vmovdqa %ymm2, (%r10, %r11)
   619: 4983c320                 addq $0x20, %r11
   623: 4c3bd8                   cmpq %rax, %r11
   626: 0f86e8fdffff             jbe 0x60
   632: 415e                     popq %r14
   634: 415d                     popq %r13
   636: 415c                     popq %r12
   638: 415f                     popq %r15
   640: 5b                       popq %rbx
   641: c5f877                   vzeroupper 
   644: c3                       retq 
   645: 0f0b                     ud2 
   647: 0f0b                     ud2 
   649: 0f0b                     ud2 
   651: 0f0b                     ud2 
   653: 0f0b                     ud2 
   655: 0f0b                     ud2 
   657: 0f0b                     ud2 
   659: 0f0b                     ud2 
   661: 0f0b                     ud2 
   663: 0f0b                     ud2 
   665: 0f0b                     ud2 
   667: 0f0b                     ud2 
   669: 0f0b                     ud2 
   671: 0f0b                     ud2 
   673: 0f0b                     ud2 
   675: 0f0b                     ud2 
   677: 0f0b                     ud2 
   679: 0f0b                     ud2 
   681: 0f0b                     ud2 
   683: 0f0b                     ud2 
   685: 0f0b                     ud2 
   687: 0f0b                     ud2 
   689: 0f0b                     ud2 
   691: 0f0b                     ud2 
   693: 0f0b                     ud2 
   695: 0f0b                     ud2 
   697: 0f0b                     ud2 
   699: 0f0b                     ud2 
   701: 0f0b                     ud2 
   703: cc                       int3 
