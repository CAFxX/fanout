; config_t1_30 target=avx512-gfni twin=notail topo=raid6(8)/enc op=encode U=1
; code_size=384 twins=["notail"] uops_per_byte=0.671875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d67010000           leaq 0x167(%rip), %r15
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
   360: 415e                     popq %r14
   362: 415d                     popq %r13
   364: 415c                     popq %r12
   366: 415f                     popq %r15
   368: 5b                       popq %rbx
   369: c5f877                   vzeroupper 
   372: c3                       retq 
   373: 0f0b                     ud2 
   375: 0f0b                     ud2 
   377: 0f0b                     ud2 
   379: 0f0b                     ud2 
   381: 0f0b                     ud2 
   383: cc                       int3 
