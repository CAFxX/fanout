; config_t1_84 target=avx512-gfni twin=maskedtail topo=rs(6,3)/enc op=encode U=1
; code_size=832 twins=["maskedtail"] uops_per_byte=0.71875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d27030000           leaq 0x327(%rip), %r15
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
   388: 4c3bd9                   cmpq %rcx, %r11
   391: 0f8379010000             jae 0x306
   397: 4c8bd1                   movq %rcx, %r10
   400: 4d2bd3                   subq %r11, %r10
   403: 498bd2                   movq %r10, %rdx
   406: 4883fa40                 cmpq $0x40, %rdx
   410: 0f8607000000             jbe 0x1a7
   416: 48c7c240000000           movq $0x40, %rdx
   423: c4c1f8908cd740010000     kmovq 0x140(%r15, %rdx, 8), %k1
   433: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
   439: 62f17548efc9             vpxord %zmm1, %zmm1, %zmm1
   445: 62f16d48efd2             vpxord %zmm2, %zmm2, %zmm2
   451: 62b17fc96f1c1e           vmovdqu8 (%rsi, %r11), %zmm3 {%k1} {z}
   458: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   464: 62d17d486f2f             vmovdqa32 (%r15), %zmm5
   470: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   477: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   483: 62d17d486f6f01           vmovdqa32 0x40(%r15), %zmm5
   490: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   497: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   503: 62b17fc96f1c1f           vmovdqu8 (%rdi, %r11), %zmm3 {%k1} {z}
   510: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   516: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   523: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   529: 62d17d486f6f02           vmovdqa32 0x80(%r15), %zmm5
   536: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   543: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   549: 62917fc96f1c1c           vmovdqu8 (%r12, %r11), %zmm3 {%k1} {z}
   556: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   562: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   569: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   575: 62d17d486f6f03           vmovdqa32 0xc0(%r15), %zmm5
   582: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   589: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   595: 62917fc96f5c1d00         vmovdqu8 (%r13, %r11), %zmm3 {%k1} {z}
   603: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   609: 62d17d486f6f04           vmovdqa32 0x100(%r15), %zmm5
   616: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   623: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   629: 62d17d486f2f             vmovdqa32 (%r15), %zmm5
   635: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   642: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   648: 62917fc96f1c1e           vmovdqu8 (%r14, %r11), %zmm3 {%k1} {z}
   655: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   661: 62d17d486f6f03           vmovdqa32 0xc0(%r15), %zmm5
   668: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   675: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   681: 62d17d486f6f04           vmovdqa32 0x100(%r15), %zmm5
   688: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   695: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   701: 498b4028                 movq 0x28(%r8), %rax
   705: 62b17fc96f1c18           vmovdqu8 (%rax, %r11), %zmm3 {%k1} {z}
   712: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   718: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   724: 62f16d48efd3             vpxord %zmm3, %zmm2, %zmm2
   730: 498b01                   movq (%r9), %rax
   733: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
   740: 498b4108                 movq 8(%r9), %rax
   744: 62b17f497f0c18           vmovdqu8 %zmm1, (%rax, %r11) {%k1}
   751: 498b4110                 movq 0x10(%r9), %rax
   755: 62b17f497f1418           vmovdqu8 %zmm2, (%rax, %r11) {%k1}
   762: 4c03da                   addq %rdx, %r11
   765: 4c2bd2                   subq %rdx, %r10
   768: 0f858dfeffff             jne 0x193
   774: 415e                     popq %r14
   776: 415d                     popq %r13
   778: 415c                     popq %r12
   780: 415f                     popq %r15
   782: 5b                       popq %rbx
   783: c5f877                   vzeroupper 
   786: c3                       retq 
   787: 0f0b                     ud2 
   789: 0f0b                     ud2 
   791: 0f0b                     ud2 
   793: 0f0b                     ud2 
   795: 0f0b                     ud2 
   797: 0f0b                     ud2 
   799: 0f0b                     ud2 
   801: 0f0b                     ud2 
   803: 0f0b                     ud2 
   805: 0f0b                     ud2 
   807: 0f0b                     ud2 
   809: 0f0b                     ud2 
   811: 0f0b                     ud2 
   813: 0f0b                     ud2 
   815: 0f0b                     ud2 
   817: 0f0b                     ud2 
   819: 0f0b                     ud2 
   821: 0f0b                     ud2 
   823: 0f0b                     ud2 
   825: 0f0b                     ud2 
   827: 0f0b                     ud2 
   829: 0f0b                     ud2 
   831: cc                       int3 
