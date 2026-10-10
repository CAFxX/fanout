; config_t3_2062 target=avx512-gfni twin=baseline topo=rs(3,1)/enc op=encode U=8
; code_size=448 twins=[] uops_per_byte=0.0625
     0: 53                       pushq %rbx
     1: 4154                     pushq %r12
     3: 4155                     pushq %r13
     5: 4156                     pushq %r14
     7: 4c8bc7                   movq %rdi, %r8
    10: 4c8bce                   movq %rsi, %r9
    13: 488bca                   movq %rdx, %rcx
    16: 498b30                   movq (%r8), %rsi
    19: 498b7808                 movq 8(%r8), %rdi
    23: 4d8b6010                 movq 0x10(%r8), %r12
    27: 4d33db                   xorq %r11, %r11
    30: 488bc1                   movq %rcx, %rax
    33: 4881e800020000           subq $0x200, %rax
    40: 0f8221010000             jb 0x14f
    46: 666666666666660f1f840000000000 nopw (%rax, %rax)
    61: 0f1f00                   nopl (%rax)
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
   253: 4d8b11                   movq (%r9), %r10
   256: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   263: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
   271: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
   279: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
   287: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
   295: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
   303: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
   311: 62917d487f7c1a07         vmovdqa32 %zmm7, 0x1c0(%r10, %r11)
   319: 4981c300020000           addq $0x200, %r11
   326: 4c3bd8                   cmpq %rax, %r11
   329: 0f86f1feffff             jbe 0x40
   335: 4c3bd9                   cmpq %rcx, %r11
   338: 0f833c000000             jae 0x194
   344: 4153                     pushq %r11
   346: 4d8b11                   movq (%r9), %r10
   349: 4c8b1c24                 movq (%rsp), %r11
   353: 498b00                   movq (%r8), %rax
   356: 428a1418                 movb (%rax, %r11), %dl
   360: 8ada                     movb %dl, %bl
   362: 498b4008                 movq 8(%r8), %rax
   366: 428a1418                 movb (%rax, %r11), %dl
   370: 32da                     xorb %dl, %bl
   372: 498b4010                 movq 0x10(%r8), %rax
   376: 428a1418                 movb (%rax, %r11), %dl
   380: 32da                     xorb %dl, %bl
   382: 43881c1a                 movb %bl, (%r10, %r11)
   386: 33db                     xorl %ebx, %ebx
   388: 49ffc3                   incq %r11
   391: 4c3bd9                   cmpq %rcx, %r11
   394: 0f82d1ffffff             jb 0x161
   400: 4883c408                 addq $8, %rsp
   404: 415e                     popq %r14
   406: 415d                     popq %r13
   408: 415c                     popq %r12
   410: 5b                       popq %rbx
   411: c5f877                   vzeroupper 
   414: c3                       retq 
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
