; config_t1_188 target=avx512-gfni twin=maskedtail topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=768 twins=["maskedtail"] uops_per_byte=0.6875
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
    54: 0f823d010000             jb 0x179
    60: 0f1f4000                 nopl (%rax)
    64: 62b17d486f141e           vmovdqa32 (%rsi, %r11), %zmm2
    71: 62d17d486f27             vmovdqa32 (%r15), %zmm4
    77: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
    84: 62f17d486fc3             vmovdqa32 %zmm3, %zmm0
    90: 62d17d486f6701           vmovdqa32 0x40(%r15), %zmm4
    97: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   104: 62f17d486fcb             vmovdqa32 %zmm3, %zmm1
   110: 62b17d486f141f           vmovdqa32 (%rdi, %r11), %zmm2
   117: 62d17d486f6702           vmovdqa32 0x80(%r15), %zmm4
   124: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   131: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   137: 62d17d486f6703           vmovdqa32 0xc0(%r15), %zmm4
   144: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   151: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   157: 62917d486f141c           vmovdqa32 (%r12, %r11), %zmm2
   164: 62d17d486f6704           vmovdqa32 0x100(%r15), %zmm4
   171: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   178: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   184: 62d17d486f6705           vmovdqa32 0x140(%r15), %zmm4
   191: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   198: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   204: 62917d486f541d00         vmovdqa32 (%r13, %r11), %zmm2
   212: 62d17d486f6706           vmovdqa32 0x180(%r15), %zmm4
   219: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   226: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   232: 62d17d486f6707           vmovdqa32 0x1c0(%r15), %zmm4
   239: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   246: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   252: 62917d486f141e           vmovdqa32 (%r14, %r11), %zmm2
   259: 62d17d486f6708           vmovdqa32 0x200(%r15), %zmm4
   266: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   273: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   279: 62d17d486f6709           vmovdqa32 0x240(%r15), %zmm4
   286: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   293: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   299: 4d8b5028                 movq 0x28(%r8), %r10
   303: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   310: 62d17d486f670a           vmovdqa32 0x280(%r15), %zmm4
   317: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   324: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   330: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   337: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   343: 4d8b11                   movq (%r9), %r10
   346: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   353: 4d8b5108                 movq 8(%r9), %r10
   357: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
   364: 4983c340                 addq $0x40, %r11
   368: 4c3bd8                   cmpq %rax, %r11
   371: 0f86c7feffff             jbe 0x40
   377: 4c3bd9                   cmpq %rcx, %r11
   380: 0f8368010000             jae 0x2ea
   386: 4c8bd1                   movq %rcx, %r10
   389: 4d2bd3                   subq %r11, %r10
   392: 498bd2                   movq %r10, %rdx
   395: 4883fa40                 cmpq $0x40, %rdx
   399: 0f8607000000             jbe 0x19c
   405: 48c7c240000000           movq $0x40, %rdx
   412: c4c1f8908cd7c0020000     kmovq 0x2c0(%r15, %rdx, 8), %k1
   422: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
   428: 62f17548efc9             vpxord %zmm1, %zmm1, %zmm1
   434: 62b17fc96f141e           vmovdqu8 (%rsi, %r11), %zmm2 {%k1} {z}
   441: 62d17d486f27             vmovdqa32 (%r15), %zmm4
   447: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   454: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   460: 62d17d486f6701           vmovdqa32 0x40(%r15), %zmm4
   467: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   474: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   480: 62b17fc96f141f           vmovdqu8 (%rdi, %r11), %zmm2 {%k1} {z}
   487: 62d17d486f6702           vmovdqa32 0x80(%r15), %zmm4
   494: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   501: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   507: 62d17d486f6703           vmovdqa32 0xc0(%r15), %zmm4
   514: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   521: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   527: 62917fc96f141c           vmovdqu8 (%r12, %r11), %zmm2 {%k1} {z}
   534: 62d17d486f6704           vmovdqa32 0x100(%r15), %zmm4
   541: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   548: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   554: 62d17d486f6705           vmovdqa32 0x140(%r15), %zmm4
   561: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   568: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   574: 62917fc96f541d00         vmovdqu8 (%r13, %r11), %zmm2 {%k1} {z}
   582: 62d17d486f6706           vmovdqa32 0x180(%r15), %zmm4
   589: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   596: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   602: 62d17d486f6707           vmovdqa32 0x1c0(%r15), %zmm4
   609: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   616: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   622: 62917fc96f141e           vmovdqu8 (%r14, %r11), %zmm2 {%k1} {z}
   629: 62d17d486f6708           vmovdqa32 0x200(%r15), %zmm4
   636: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   643: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   649: 62d17d486f6709           vmovdqa32 0x240(%r15), %zmm4
   656: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   663: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   669: 498b4028                 movq 0x28(%r8), %rax
   673: 62b17fc96f1418           vmovdqu8 (%rax, %r11), %zmm2 {%k1} {z}
   680: 62d17d486f670a           vmovdqa32 0x280(%r15), %zmm4
   687: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   694: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   700: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   707: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   713: 498b01                   movq (%r9), %rax
   716: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
   723: 498b4108                 movq 8(%r9), %rax
   727: 62b17f497f0c18           vmovdqu8 %zmm1, (%rax, %r11) {%k1}
   734: 4c03da                   addq %rdx, %r11
   737: 4c2bd2                   subq %rdx, %r10
   740: 0f859efeffff             jne 0x188
   746: 415e                     popq %r14
   748: 415d                     popq %r13
   750: 415c                     popq %r12
   752: 415f                     popq %r15
   754: 5b                       popq %rbx
   755: c5f877                   vzeroupper 
   758: c3                       retq 
   759: 0f0b                     ud2 
   761: 0f0b                     ud2 
   763: 0f0b                     ud2 
   765: 0f0b                     ud2 
   767: cc                       int3 
