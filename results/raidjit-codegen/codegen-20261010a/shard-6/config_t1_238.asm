; config_t1_238 target=avx512-gfni twin=notail topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=256 twins=["notail"] uops_per_byte=0.234375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3ded000000           leaq 0xed(%rip), %r15
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
   208: 415f                     popq %r15
   210: 5b                       popq %rbx
   211: c5f877                   vzeroupper 
   214: c3                       retq 
   215: 0f0b                     ud2 
   217: 0f0b                     ud2 
   219: 0f0b                     ud2 
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
