; config_t1_175 target=avx512-gfni twin=maskedtail topo=rs(6,3)/rec{0} op=recover U=8
; code_size=704 twins=["maskedtail"] uops_per_byte=0.111328125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da7020000           leaq 0x2a7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e800020000           subq $0x200, %rax
    57: 0f82d2010000             jb 0x211
    63: 90                       nop 
    64: 62b17d486f041e           vmovdqa32 (%rsi, %r11), %zmm0
    71: 62b17d486f4c1e01         vmovdqa32 0x40(%rsi, %r11), %zmm1
    79: 62b17d486f541e02         vmovdqa32 0x80(%rsi, %r11), %zmm2
    87: 62b17d486f5c1e03         vmovdqa32 0xc0(%rsi, %r11), %zmm3
    95: 62b17d486f641e04         vmovdqa32 0x100(%rsi, %r11), %zmm4
   103: 62b17d486f6c1e05         vmovdqa32 0x140(%rsi, %r11), %zmm5
   111: 62b17d486f741e06         vmovdqa32 0x180(%rsi, %r11), %zmm6
   119: 62b17d486f7c1e07         vmovdqa32 0x1c0(%rsi, %r11), %zmm7
   127: 62b17d48ef041f           vpxord (%rdi, %r11), %zmm0, %zmm0
   134: 62b17548ef4c1f01         vpxord 0x40(%rdi, %r11), %zmm1, %zmm1
   142: 62b16d48ef541f02         vpxord 0x80(%rdi, %r11), %zmm2, %zmm2
   150: 62b16548ef5c1f03         vpxord 0xc0(%rdi, %r11), %zmm3, %zmm3
   158: 62b15d48ef641f04         vpxord 0x100(%rdi, %r11), %zmm4, %zmm4
   166: 62b15548ef6c1f05         vpxord 0x140(%rdi, %r11), %zmm5, %zmm5
   174: 62b14d48ef741f06         vpxord 0x180(%rdi, %r11), %zmm6, %zmm6
   182: 62b14548ef7c1f07         vpxord 0x1c0(%rdi, %r11), %zmm7, %zmm7
   190: 62917d48ef041c           vpxord (%r12, %r11), %zmm0, %zmm0
   197: 62917548ef4c1c01         vpxord 0x40(%r12, %r11), %zmm1, %zmm1
   205: 62916d48ef541c02         vpxord 0x80(%r12, %r11), %zmm2, %zmm2
   213: 62916548ef5c1c03         vpxord 0xc0(%r12, %r11), %zmm3, %zmm3
   221: 62915d48ef641c04         vpxord 0x100(%r12, %r11), %zmm4, %zmm4
   229: 62915548ef6c1c05         vpxord 0x140(%r12, %r11), %zmm5, %zmm5
   237: 62914d48ef741c06         vpxord 0x180(%r12, %r11), %zmm6, %zmm6
   245: 62914548ef7c1c07         vpxord 0x1c0(%r12, %r11), %zmm7, %zmm7
   253: 62917d48ef441d00         vpxord (%r13, %r11), %zmm0, %zmm0
   261: 62917548ef4c1d01         vpxord 0x40(%r13, %r11), %zmm1, %zmm1
   269: 62916d48ef541d02         vpxord 0x80(%r13, %r11), %zmm2, %zmm2
   277: 62916548ef5c1d03         vpxord 0xc0(%r13, %r11), %zmm3, %zmm3
   285: 62915d48ef641d04         vpxord 0x100(%r13, %r11), %zmm4, %zmm4
   293: 62915548ef6c1d05         vpxord 0x140(%r13, %r11), %zmm5, %zmm5
   301: 62914d48ef741d06         vpxord 0x180(%r13, %r11), %zmm6, %zmm6
   309: 62914548ef7c1d07         vpxord 0x1c0(%r13, %r11), %zmm7, %zmm7
   317: 62917d48ef041e           vpxord (%r14, %r11), %zmm0, %zmm0
   324: 62917548ef4c1e01         vpxord 0x40(%r14, %r11), %zmm1, %zmm1
   332: 62916d48ef541e02         vpxord 0x80(%r14, %r11), %zmm2, %zmm2
   340: 62916548ef5c1e03         vpxord 0xc0(%r14, %r11), %zmm3, %zmm3
   348: 62915d48ef641e04         vpxord 0x100(%r14, %r11), %zmm4, %zmm4
   356: 62915548ef6c1e05         vpxord 0x140(%r14, %r11), %zmm5, %zmm5
   364: 62914d48ef741e06         vpxord 0x180(%r14, %r11), %zmm6, %zmm6
   372: 62914548ef7c1e07         vpxord 0x1c0(%r14, %r11), %zmm7, %zmm7
   380: 4d8b5028                 movq 0x28(%r8), %r10
   384: 62917d48ef041a           vpxord (%r10, %r11), %zmm0, %zmm0
   391: 62917548ef4c1a01         vpxord 0x40(%r10, %r11), %zmm1, %zmm1
   399: 62916d48ef541a02         vpxord 0x80(%r10, %r11), %zmm2, %zmm2
   407: 62916548ef5c1a03         vpxord 0xc0(%r10, %r11), %zmm3, %zmm3
   415: 62915d48ef641a04         vpxord 0x100(%r10, %r11), %zmm4, %zmm4
   423: 62915548ef6c1a05         vpxord 0x140(%r10, %r11), %zmm5, %zmm5
   431: 62914d48ef741a06         vpxord 0x180(%r10, %r11), %zmm6, %zmm6
   439: 62914548ef7c1a07         vpxord 0x1c0(%r10, %r11), %zmm7, %zmm7
   447: 4d8b11                   movq (%r9), %r10
   450: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   457: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
   465: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
   473: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
   481: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
   489: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
   497: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
   505: 62917d487f7c1a07         vmovdqa32 %zmm7, 0x1c0(%r10, %r11)
   513: 4981c300020000           addq $0x200, %r11
   520: 4c3bd8                   cmpq %rax, %r11
   523: 0f862ffeffff             jbe 0x40
   529: 4c3bd9                   cmpq %rcx, %r11
   532: 0f838f000000             jae 0x2a9
   538: 4c8bd1                   movq %rcx, %r10
   541: 4d2bd3                   subq %r11, %r10
   544: 498bd2                   movq %r10, %rdx
   547: 4883fa40                 cmpq $0x40, %rdx
   551: 0f8607000000             jbe 0x234
   557: 48c7c240000000           movq $0x40, %rdx
   564: c4c1f8900cd7             kmovq (%r15, %rdx, 8), %k1
   570: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
   576: 62317fc96f041e           vmovdqu8 (%rsi, %r11), %zmm8 {%k1} {z}
   583: 62d17d48efc0             vpxord %zmm8, %zmm0, %zmm0
   589: 62317fc96f041f           vmovdqu8 (%rdi, %r11), %zmm8 {%k1} {z}
   596: 62d17d48efc0             vpxord %zmm8, %zmm0, %zmm0
   602: 62117fc96f041c           vmovdqu8 (%r12, %r11), %zmm8 {%k1} {z}
   609: 62d17d48efc0             vpxord %zmm8, %zmm0, %zmm0
   615: 62117fc96f441d00         vmovdqu8 (%r13, %r11), %zmm8 {%k1} {z}
   623: 62d17d48efc0             vpxord %zmm8, %zmm0, %zmm0
   629: 62117fc96f041e           vmovdqu8 (%r14, %r11), %zmm8 {%k1} {z}
   636: 62d17d48efc0             vpxord %zmm8, %zmm0, %zmm0
   642: 498b4028                 movq 0x28(%r8), %rax
   646: 62317fc96f0418           vmovdqu8 (%rax, %r11), %zmm8 {%k1} {z}
   653: 62d17d48efc0             vpxord %zmm8, %zmm0, %zmm0
   659: 498b01                   movq (%r9), %rax
   662: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
   669: 4c03da                   addq %rdx, %r11
   672: 4c2bd2                   subq %rdx, %r10
   675: 0f8577ffffff             jne 0x220
   681: 415e                     popq %r14
   683: 415d                     popq %r13
   685: 415c                     popq %r12
   687: 415f                     popq %r15
   689: 5b                       popq %rbx
   690: c5f877                   vzeroupper 
   693: c3                       retq 
   694: 0f0b                     ud2 
   696: 0f0b                     ud2 
   698: 0f0b                     ud2 
   700: 0f0b                     ud2 
   702: 0f0b                     ud2 
