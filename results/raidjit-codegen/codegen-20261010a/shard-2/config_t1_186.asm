; config_t1_186 target=avx512-gfni twin=notail topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=448 twins=["notail"] uops_per_byte=0.6875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da7010000           leaq 0x1a7(%rip), %r15
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
   377: 415e                     popq %r14
   379: 415d                     popq %r13
   381: 415c                     popq %r12
   383: 415f                     popq %r15
   385: 5b                       popq %rbx
   386: c5f877                   vzeroupper 
   389: c3                       retq 
   390: 0f0b                     ud2 
   392: 0f0b                     ud2 
   394: 0f0b                     ud2 
   396: 0f0b                     ud2 
   398: 0f0b                     ud2 
   400: 0f0b                     ud2 
   402: 0f0b                     ud2 
   404: 0f0b                     ud2 
   406: 0f0b                     ud2 
   408: 0f0b                     ud2 
   410: 0f0b                     ud2 
   412: 0f0b                     ud2 
   414: 0f0b                     ud2 
   416: 0f0b                     ud2 
   418: 0f0b                     ud2 
   420: 0f0b                     ud2 
   422: 0f0b                     ud2 
   424: 0f0b                     ud2 
   426: 0f0b                     ud2 
   428: 0f0b                     ud2 
   430: 0f0b                     ud2 
   432: 0f0b                     ud2 
   434: 0f0b                     ud2 
   436: 0f0b                     ud2 
   438: 0f0b                     ud2 
   440: 0f0b                     ud2 
   442: 0f0b                     ud2 
   444: 0f0b                     ud2 
   446: 0f0b                     ud2 
