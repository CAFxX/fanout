; config_t1_32 target=avx512-gfni twin=maskedtail topo=raid6(8)/enc op=encode U=1
; code_size=768 twins=["maskedtail"] uops_per_byte=0.671875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de7020000           leaq 0x2e7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4883e840                 subq $0x40, %rax
    54: 0f822c010000             jb 0x168
    60: 0f1f4000                 nopl (%rax)
    64: 62b17d486f141e           vmovdqa32 (%rsi, %r11), %zmm2
    71: 62f17d486fc2             vmovdqa32 %zmm2, %zmm0
    77: 62d17d486f27             vmovdqa32 (%r15), %zmm4
    83: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
    90: 62f17d486fcb             vmovdqa32 %zmm3, %zmm1
    96: 62b17d486f141f           vmovdqa32 (%rdi, %r11), %zmm2
   103: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   109: 62d17d486f6701           vmovdqa32 0x40(%r15), %zmm4
   116: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   123: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   129: 62917d486f141c           vmovdqa32 (%r12, %r11), %zmm2
   136: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   142: 62d17d486f6702           vmovdqa32 0x80(%r15), %zmm4
   149: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   156: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   162: 62917d486f541d00         vmovdqa32 (%r13, %r11), %zmm2
   170: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   176: 62d17d486f6703           vmovdqa32 0xc0(%r15), %zmm4
   183: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   190: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   196: 62917d486f141e           vmovdqa32 (%r14, %r11), %zmm2
   203: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   209: 62d17d486f6704           vmovdqa32 0x100(%r15), %zmm4
   216: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   223: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   229: 4d8b5028                 movq 0x28(%r8), %r10
   233: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   240: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   246: 62d17d486f6705           vmovdqa32 0x140(%r15), %zmm4
   253: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   260: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   266: 4d8b5030                 movq 0x30(%r8), %r10
   270: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   277: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   283: 62d17d486f6706           vmovdqa32 0x180(%r15), %zmm4
   290: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   297: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   303: 4d8b5038                 movq 0x38(%r8), %r10
   307: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   314: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   320: 62f17548efca             vpxord %zmm2, %zmm1, %zmm1
   326: 4d8b11                   movq (%r9), %r10
   329: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   336: 4d8b5108                 movq 8(%r9), %r10
   340: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
   347: 4983c340                 addq $0x40, %r11
   351: 4c3bd8                   cmpq %rax, %r11
   354: 0f86d8feffff             jbe 0x40
   360: 4c3bd9                   cmpq %rcx, %r11
   363: 0f8357010000             jae 0x2c8
   369: 4c8bd1                   movq %rcx, %r10
   372: 4d2bd3                   subq %r11, %r10
   375: 498bd2                   movq %r10, %rdx
   378: 4883fa40                 cmpq $0x40, %rdx
   382: 0f8607000000             jbe 0x18b
   388: 48c7c240000000           movq $0x40, %rdx
   395: c4c1f8908cd7c0010000     kmovq 0x1c0(%r15, %rdx, 8), %k1
   405: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
   411: 62f17548efc9             vpxord %zmm1, %zmm1, %zmm1
   417: 62b17fc96f141e           vmovdqu8 (%rsi, %r11), %zmm2 {%k1} {z}
   424: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   430: 62d17d486f27             vmovdqa32 (%r15), %zmm4
   436: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   443: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   449: 62b17fc96f141f           vmovdqu8 (%rdi, %r11), %zmm2 {%k1} {z}
   456: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   462: 62d17d486f6701           vmovdqa32 0x40(%r15), %zmm4
   469: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   476: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   482: 62917fc96f141c           vmovdqu8 (%r12, %r11), %zmm2 {%k1} {z}
   489: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   495: 62d17d486f6702           vmovdqa32 0x80(%r15), %zmm4
   502: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   509: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   515: 62917fc96f541d00         vmovdqu8 (%r13, %r11), %zmm2 {%k1} {z}
   523: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   529: 62d17d486f6703           vmovdqa32 0xc0(%r15), %zmm4
   536: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   543: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   549: 62917fc96f141e           vmovdqu8 (%r14, %r11), %zmm2 {%k1} {z}
   556: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   562: 62d17d486f6704           vmovdqa32 0x100(%r15), %zmm4
   569: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   576: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   582: 498b4028                 movq 0x28(%r8), %rax
   586: 62b17fc96f1418           vmovdqu8 (%rax, %r11), %zmm2 {%k1} {z}
   593: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   599: 62d17d486f6705           vmovdqa32 0x140(%r15), %zmm4
   606: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   613: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   619: 498b4030                 movq 0x30(%r8), %rax
   623: 62b17fc96f1418           vmovdqu8 (%rax, %r11), %zmm2 {%k1} {z}
   630: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   636: 62d17d486f6706           vmovdqa32 0x180(%r15), %zmm4
   643: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   650: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   656: 498b4038                 movq 0x38(%r8), %rax
   660: 62b17fc96f1418           vmovdqu8 (%rax, %r11), %zmm2 {%k1} {z}
   667: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   673: 62f17548efca             vpxord %zmm2, %zmm1, %zmm1
   679: 498b01                   movq (%r9), %rax
   682: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
   689: 498b4108                 movq 8(%r9), %rax
   693: 62b17f497f0c18           vmovdqu8 %zmm1, (%rax, %r11) {%k1}
   700: 4c03da                   addq %rdx, %r11
   703: 4c2bd2                   subq %rdx, %r10
   706: 0f85affeffff             jne 0x177
   712: 415e                     popq %r14
   714: 415d                     popq %r13
   716: 415c                     popq %r12
   718: 415f                     popq %r15
   720: 5b                       popq %rbx
   721: c5f877                   vzeroupper 
   724: c3                       retq 
   725: 0f0b                     ud2 
   727: 0f0b                     ud2 
   729: 0f0b                     ud2 
   731: 0f0b                     ud2 
   733: 0f0b                     ud2 
   735: 0f0b                     ud2 
   737: 0f0b                     ud2 
   739: 0f0b                     ud2 
   741: 0f0b                     ud2 
   743: 0f0b                     ud2 
   745: 0f0b                     ud2 
   747: 0f0b                     ud2 
   749: 0f0b                     ud2 
   751: 0f0b                     ud2 
   753: 0f0b                     ud2 
   755: 0f0b                     ud2 
   757: 0f0b                     ud2 
   759: 0f0b                     ud2 
   761: 0f0b                     ud2 
   763: 0f0b                     ud2 
   765: 0f0b                     ud2 
   767: cc                       int3 
