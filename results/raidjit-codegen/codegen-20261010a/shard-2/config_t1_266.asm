; config_t1_266 target=avx512-gfni twin=maskedtail topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=448 twins=["maskedtail"] uops_per_byte=0.25
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3dad010000           leaq 0x1ad(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: 488bc1                   movq %rcx, %rax
    25: 4883e840                 subq $0x40, %rax
    29: 0f82b3000000             jb 0xd6
    35: 666666666666660f1f840000000000 nopw (%rax, %rax)
    50: 6666666666660f1f840000000000 nopw (%rax, %rax)
    64: 4d8b5020                 movq 0x20(%r8), %r10
    68: 62917d486f041a           vmovdqa32 (%r10, %r11), %zmm0
    75: 4d8b5028                 movq 0x28(%r8), %r10
    79: 62917d486f0c1a           vmovdqa32 (%r10, %r11), %zmm1
    86: 4d8b10                   movq (%r8), %r10
    89: 498b5010                 movq 0x10(%r8), %rdx
    93: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   100: 62b16d48ef141a           vpxord (%rdx, %r11), %zmm2, %zmm2
   107: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   113: 62d17d486f27             vmovdqa32 (%r15), %zmm4
   119: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   126: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   132: 4d8b5008                 movq 8(%r8), %r10
   136: 498b5018                 movq 0x18(%r8), %rdx
   140: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   147: 62b16d48ef141a           vpxord (%rdx, %r11), %zmm2, %zmm2
   154: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   160: 62d17d486f6701           vmovdqa32 0x40(%r15), %zmm4
   167: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   174: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   180: 4d8b11                   movq (%r9), %r10
   183: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   190: 4d8b5108                 movq 8(%r9), %r10
   194: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
   201: 4983c340                 addq $0x40, %r11
   205: 4c3bd8                   cmpq %rax, %r11
   208: 0f866affffff             jbe 0x40
   214: 4c3bd9                   cmpq %rcx, %r11
   217: 0f83c5000000             jae 0x1a4
   223: 4c8bd1                   movq %rcx, %r10
   226: 4d2bd3                   subq %r11, %r10
   229: 498bd2                   movq %r10, %rdx
   232: 4883fa40                 cmpq $0x40, %rdx
   236: 0f8607000000             jbe 0xf9
   242: 48c7c240000000           movq $0x40, %rdx
   249: c4c1f8908cd780000000     kmovq 0x80(%r15, %rdx, 8), %k1
   259: 498b4020                 movq 0x20(%r8), %rax
   263: 62b17fc96f0418           vmovdqu8 (%rax, %r11), %zmm0 {%k1} {z}
   270: 498b4028                 movq 0x28(%r8), %rax
   274: 62b17fc96f0c18           vmovdqu8 (%rax, %r11), %zmm1 {%k1} {z}
   281: 498b00                   movq (%r8), %rax
   284: 62b17fc96f1418           vmovdqu8 (%rax, %r11), %zmm2 {%k1} {z}
   291: 498b4010                 movq 0x10(%r8), %rax
   295: 62b17fc96f1c18           vmovdqu8 (%rax, %r11), %zmm3 {%k1} {z}
   302: 62f16d48efd3             vpxord %zmm3, %zmm2, %zmm2
   308: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   314: 62d17d486f27             vmovdqa32 (%r15), %zmm4
   320: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   327: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   333: 498b4008                 movq 8(%r8), %rax
   337: 62b17fc96f1418           vmovdqu8 (%rax, %r11), %zmm2 {%k1} {z}
   344: 498b4018                 movq 0x18(%r8), %rax
   348: 62b17fc96f1c18           vmovdqu8 (%rax, %r11), %zmm3 {%k1} {z}
   355: 62f16d48efd3             vpxord %zmm3, %zmm2, %zmm2
   361: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   367: 62d17d486f6701           vmovdqa32 0x40(%r15), %zmm4
   374: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   381: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   387: 498b01                   movq (%r9), %rax
   390: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
   397: 498b4108                 movq 8(%r9), %rax
   401: 62b17f497f0c18           vmovdqu8 %zmm1, (%rax, %r11) {%k1}
   408: 4c03da                   addq %rdx, %r11
   411: 4c2bd2                   subq %rdx, %r10
   414: 0f8541ffffff             jne 0xe5
   420: 415f                     popq %r15
   422: 5b                       popq %rbx
   423: c5f877                   vzeroupper 
   426: c3                       retq 
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
