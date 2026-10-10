; config_t1_82 target=avx512-gfni twin=notail topo=rs(6,3)/enc op=encode U=1
; code_size=448 twins=["notail"] uops_per_byte=0.71875
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
    54: 0f8248010000             jb 0x184
    60: 0f1f4000                 nopl (%rax)
    64: 62b17d486f1c1e           vmovdqa32 (%rsi, %r11), %zmm3
    71: 62f17d486fc3             vmovdqa32 %zmm3, %zmm0
    77: 62d17d486f2f             vmovdqa32 (%r15), %zmm5
    83: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
    90: 62f17d486fcc             vmovdqa32 %zmm4, %zmm1
    96: 62d17d486f6f01           vmovdqa32 0x40(%r15), %zmm5
   103: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   110: 62f17d486fd4             vmovdqa32 %zmm4, %zmm2
   116: 62b17d486f1c1f           vmovdqa32 (%rdi, %r11), %zmm3
   123: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   129: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   136: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   142: 62d17d486f6f02           vmovdqa32 0x80(%r15), %zmm5
   149: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   156: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   162: 62917d486f1c1c           vmovdqa32 (%r12, %r11), %zmm3
   169: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   175: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   182: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   188: 62d17d486f6f03           vmovdqa32 0xc0(%r15), %zmm5
   195: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   202: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   208: 62917d486f5c1d00         vmovdqa32 (%r13, %r11), %zmm3
   216: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   222: 62d17d486f6f04           vmovdqa32 0x100(%r15), %zmm5
   229: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   236: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   242: 62d17d486f2f             vmovdqa32 (%r15), %zmm5
   248: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   255: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   261: 62917d486f1c1e           vmovdqa32 (%r14, %r11), %zmm3
   268: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   274: 62d17d486f6f03           vmovdqa32 0xc0(%r15), %zmm5
   281: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   288: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   294: 62d17d486f6f04           vmovdqa32 0x100(%r15), %zmm5
   301: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   308: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   314: 4d8b5028                 movq 0x28(%r8), %r10
   318: 62917d486f1c1a           vmovdqa32 (%r10, %r11), %zmm3
   325: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   331: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   337: 62f16d48efd3             vpxord %zmm3, %zmm2, %zmm2
   343: 4d8b11                   movq (%r9), %r10
   346: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   353: 4d8b5108                 movq 8(%r9), %r10
   357: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
   364: 4d8b5110                 movq 0x10(%r9), %r10
   368: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
   375: 4983c340                 addq $0x40, %r11
   379: 4c3bd8                   cmpq %rax, %r11
   382: 0f86bcfeffff             jbe 0x40
   388: 415e                     popq %r14
   390: 415d                     popq %r13
   392: 415c                     popq %r12
   394: 415f                     popq %r15
   396: 5b                       popq %rbx
   397: c5f877                   vzeroupper 
   400: c3                       retq 
   401: 0f0b                     ud2 
   403: 0f0b                     ud2 
   405: 0f0b                     ud2 
   407: 0f0b                     ud2 
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
