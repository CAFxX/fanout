; config_t1_240 target=avx512-gfni twin=maskedtail topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=448 twins=["maskedtail"] uops_per_byte=0.234375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3dad010000           leaq 0x1ad(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: 488bc1                   movq %rcx, %rax
    25: 4883e840                 subq $0x40, %rax
    29: 0f82ad000000             jb 0xd0
    35: 666666666666660f1f840000000000 nopw (%rax, %rax)
    50: 6666666666660f1f840000000000 nopw (%rax, %rax)
    64: 4d8b5010                 movq 0x10(%r8), %r10
    68: 62917d486f041a           vmovdqa32 (%r10, %r11), %zmm0
    75: 4d8b5018                 movq 0x18(%r8), %r10
    79: 62917d486f0c1a           vmovdqa32 (%r10, %r11), %zmm1
    86: 4d8b5020                 movq 0x20(%r8), %r10
    90: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
    97: 4d8b10                   movq (%r8), %r10
   100: 498b5008                 movq 8(%r8), %rdx
   104: 62917d486f1c1a           vmovdqa32 (%r10, %r11), %zmm3
   111: 62b16548ef1c1a           vpxord (%rdx, %r11), %zmm3, %zmm3
   118: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   124: 62d17d486f2f             vmovdqa32 (%r15), %zmm5
   130: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   137: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   143: 62d17d486f6f01           vmovdqa32 0x40(%r15), %zmm5
   150: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   157: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   163: 4d8b11                   movq (%r9), %r10
   166: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   173: 4d8b5108                 movq 8(%r9), %r10
   177: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
   184: 4d8b5110                 movq 0x10(%r9), %r10
   188: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
   195: 4983c340                 addq $0x40, %r11
   199: 4c3bd8                   cmpq %rax, %r11
   202: 0f8670ffffff             jbe 0x40
   208: 4c3bd9                   cmpq %rcx, %r11
   211: 0f83b9000000             jae 0x192
   217: 4c8bd1                   movq %rcx, %r10
   220: 4d2bd3                   subq %r11, %r10
   223: 498bd2                   movq %r10, %rdx
   226: 4883fa40                 cmpq $0x40, %rdx
   230: 0f8607000000             jbe 0xf3
   236: 48c7c240000000           movq $0x40, %rdx
   243: c4c1f8908cd780000000     kmovq 0x80(%r15, %rdx, 8), %k1
   253: 498b4010                 movq 0x10(%r8), %rax
   257: 62b17fc96f0418           vmovdqu8 (%rax, %r11), %zmm0 {%k1} {z}
   264: 498b4018                 movq 0x18(%r8), %rax
   268: 62b17fc96f0c18           vmovdqu8 (%rax, %r11), %zmm1 {%k1} {z}
   275: 498b4020                 movq 0x20(%r8), %rax
   279: 62b17fc96f1418           vmovdqu8 (%rax, %r11), %zmm2 {%k1} {z}
   286: 498b00                   movq (%r8), %rax
   289: 62b17fc96f1c18           vmovdqu8 (%rax, %r11), %zmm3 {%k1} {z}
   296: 498b4008                 movq 8(%r8), %rax
   300: 62b17fc96f2418           vmovdqu8 (%rax, %r11), %zmm4 {%k1} {z}
   307: 62f16548efdc             vpxord %zmm4, %zmm3, %zmm3
   313: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   319: 62d17d486f2f             vmovdqa32 (%r15), %zmm5
   325: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   332: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   338: 62d17d486f6f01           vmovdqa32 0x40(%r15), %zmm5
   345: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   352: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   358: 498b01                   movq (%r9), %rax
   361: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
   368: 498b4108                 movq 8(%r9), %rax
   372: 62b17f497f0c18           vmovdqu8 %zmm1, (%rax, %r11) {%k1}
   379: 498b4110                 movq 0x10(%r9), %rax
   383: 62b17f497f1418           vmovdqu8 %zmm2, (%rax, %r11) {%k1}
   390: 4c03da                   addq %rdx, %r11
   393: 4c2bd2                   subq %rdx, %r10
   396: 0f854dffffff             jne 0xdf
   402: 415f                     popq %r15
   404: 5b                       popq %rbx
   405: c5f877                   vzeroupper 
   408: c3                       retq 
   409: 0f0b                     ud2 
   411: 0f0b                     ud2 
   413: 0f0b                     ud2 
   415: 0f0b                     ud2 
   417: 0f0b                     ud2 
   419: 0f0b                     ud2 
   421: 0f0b                     ud2 
   423: 0f0b                     ud2 
   425: 0f0b                     ud2 
   427: 0f0b                     ud2 
   429: 0f0b                     ud2 
   431: 0f0b                     ud2 
   433: 0f0b                     ud2 
   435: 0f0b                     ud2 
   437: 0f0b                     ud2 
   439: 0f0b                     ud2 
   441: 0f0b                     ud2 
   443: 0f0b                     ud2 
   445: 0f0b                     ud2 
   447: cc                       int3 
