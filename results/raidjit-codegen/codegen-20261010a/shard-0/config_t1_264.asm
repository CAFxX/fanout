; config_t1_264 target=avx512-gfni twin=notail topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=256 twins=["notail"] uops_per_byte=0.25
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3ded000000           leaq 0xed(%rip), %r15
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
   214: 415f                     popq %r15
   216: 5b                       popq %rbx
   217: c5f877                   vzeroupper 
   220: c3                       retq 
   221: 0f0b                     ud2 
   223: 0f0b                     ud2 
   225: 0f0b                     ud2 
   227: 0f0b                     ud2 
   229: 0f0b                     ud2 
   231: 0f0b                     ud2 
   233: 0f0b                     ud2 
   235: 0f0b                     ud2 
   237: 0f0b                     ud2 
   239: 0f0b                     ud2 
   241: 0f0b                     ud2 
   243: 0f0b                     ud2 
   245: 0f0b                     ud2 
   247: 0f0b                     ud2 
   249: 0f0b                     ud2 
   251: 0f0b                     ud2 
   253: 0f0b                     ud2 
   255: cc                       int3 
