; config_t1_550 target=avx2 twin=notail topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=320 twins=["notail"] uops_per_byte=0.875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3d2d010000           leaq 0x12d(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: c4417d6f07               vmovdqa (%r15), %ymm8
    27: 488bc1                   movq %rcx, %rax
    30: 4883e820                 subq $0x20, %rax
    34: 0f82de000000             jb 0x106
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5020                 movq 0x20(%r8), %r10
    68: c4817d6f041a             vmovdqa (%r10, %r11), %ymm0
    74: 4d8b5028                 movq 0x28(%r8), %r10
    78: c4817d6f0c1a             vmovdqa (%r10, %r11), %ymm1
    84: 4d8b10                   movq (%r8), %r10
    87: 498b5010                 movq 0x10(%r8), %rdx
    91: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
    97: c4a16def141a             vpxor (%rdx, %r11), %ymm2, %ymm2
   103: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   108: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   113: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   118: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   122: c4c17d6f7f40             vmovdqa 0x40(%r15), %ymm7
   128: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   133: c4c17d6fbf80000000       vmovdqa 0x80(%r15), %ymm7
   142: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   147: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   151: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   155: 4d8b5008                 movq 8(%r8), %r10
   159: 498b5018                 movq 0x18(%r8), %rdx
   163: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   169: c4a16def141a             vpxor (%rdx, %r11), %ymm2, %ymm2
   175: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   180: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   185: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   190: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   194: c4c17d6fbfc0000000       vmovdqa 0xc0(%r15), %ymm7
   203: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   208: c4c17d6fbf00010000       vmovdqa 0x100(%r15), %ymm7
   217: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   222: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   226: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   230: 4d8b11                   movq (%r9), %r10
   233: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   239: 4d8b5108                 movq 8(%r9), %r10
   243: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
   249: 4983c320                 addq $0x20, %r11
   253: 4c3bd8                   cmpq %rax, %r11
   256: 0f863affffff             jbe 0x40
   262: 415f                     popq %r15
   264: 5b                       popq %rbx
   265: c5f877                   vzeroupper 
   268: c3                       retq 
   269: 0f0b                     ud2 
   271: 0f0b                     ud2 
   273: 0f0b                     ud2 
   275: 0f0b                     ud2 
   277: 0f0b                     ud2 
   279: 0f0b                     ud2 
   281: 0f0b                     ud2 
   283: 0f0b                     ud2 
   285: 0f0b                     ud2 
   287: 0f0b                     ud2 
   289: 0f0b                     ud2 
   291: 0f0b                     ud2 
   293: 0f0b                     ud2 
   295: 0f0b                     ud2 
   297: 0f0b                     ud2 
   299: 0f0b                     ud2 
   301: 0f0b                     ud2 
   303: 0f0b                     ud2 
   305: 0f0b                     ud2 
   307: 0f0b                     ud2 
   309: 0f0b                     ud2 
   311: 0f0b                     ud2 
   313: 0f0b                     ud2 
   315: 0f0b                     ud2 
   317: 0f0b                     ud2 
   319: cc                       int3 
