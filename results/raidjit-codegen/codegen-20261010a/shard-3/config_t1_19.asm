; config_t1_19 target=avx512-gfni twin=maskedtail topo=raid5(4)/enc op=encode U=8
; code_size=576 twins=["maskedtail"] uops_per_byte=0.078125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d27020000           leaq 0x227(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d33db                   xorq %r11, %r11
    43: 488bc1                   movq %rcx, %rax
    46: 4881e800020000           subq $0x200, %rax
    53: 0f8254010000             jb 0x18f
    59: 0f1f440000               nopl (%rax, %rax)
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
   317: 4d8b11                   movq (%r9), %r10
   320: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   327: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
   335: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
   343: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
   351: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
   359: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
   367: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
   375: 62917d487f7c1a07         vmovdqa32 %zmm7, 0x1c0(%r10, %r11)
   383: 4981c300020000           addq $0x200, %r11
   390: 4c3bd8                   cmpq %rax, %r11
   393: 0f86b1feffff             jbe 0x40
   399: 4c3bd9                   cmpq %rcx, %r11
   402: 0f8371000000             jae 0x209
   408: 4c8bd1                   movq %rcx, %r10
   411: 4d2bd3                   subq %r11, %r10
   414: 498bd2                   movq %r10, %rdx
   417: 4883fa40                 cmpq $0x40, %rdx
   421: 0f8607000000             jbe 0x1b2
   427: 48c7c240000000           movq $0x40, %rdx
   434: c4c1f8900cd7             kmovq (%r15, %rdx, 8), %k1
   440: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
   446: 62317fc96f041e           vmovdqu8 (%rsi, %r11), %zmm8 {%k1} {z}
   453: 62d17d48efc0             vpxord %zmm8, %zmm0, %zmm0
   459: 62317fc96f041f           vmovdqu8 (%rdi, %r11), %zmm8 {%k1} {z}
   466: 62d17d48efc0             vpxord %zmm8, %zmm0, %zmm0
   472: 62117fc96f041c           vmovdqu8 (%r12, %r11), %zmm8 {%k1} {z}
   479: 62d17d48efc0             vpxord %zmm8, %zmm0, %zmm0
   485: 62117fc96f441d00         vmovdqu8 (%r13, %r11), %zmm8 {%k1} {z}
   493: 62d17d48efc0             vpxord %zmm8, %zmm0, %zmm0
   499: 498b01                   movq (%r9), %rax
   502: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
   509: 4c03da                   addq %rdx, %r11
   512: 4c2bd2                   subq %rdx, %r10
   515: 0f8595ffffff             jne 0x19e
   521: 415e                     popq %r14
   523: 415d                     popq %r13
   525: 415c                     popq %r12
   527: 415f                     popq %r15
   529: 5b                       popq %rbx
   530: c5f877                   vzeroupper 
   533: c3                       retq 
   534: 0f0b                     ud2 
   536: 0f0b                     ud2 
   538: 0f0b                     ud2 
   540: 0f0b                     ud2 
   542: 0f0b                     ud2 
   544: 0f0b                     ud2 
   546: 0f0b                     ud2 
   548: 0f0b                     ud2 
   550: 0f0b                     ud2 
   552: 0f0b                     ud2 
   554: 0f0b                     ud2 
   556: 0f0b                     ud2 
   558: 0f0b                     ud2 
   560: 0f0b                     ud2 
   562: 0f0b                     ud2 
   564: 0f0b                     ud2 
   566: 0f0b                     ud2 
   568: 0f0b                     ud2 
   570: 0f0b                     ud2 
   572: 0f0b                     ud2 
   574: 0f0b                     ud2 
