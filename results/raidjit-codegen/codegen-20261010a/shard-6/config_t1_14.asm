; config_t1_14 target=avx512-gfni twin=nt topo=raid5(4)/enc op=encode U=8
; code_size=512 twins=["nt"] uops_per_byte=0.078125
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
    27: 4d8b6818                 movq 0x18(%r8), %r13
    31: 4d33db                   xorq %r11, %r11
    34: 488bc1                   movq %rcx, %rax
    37: 4881e800020000           subq $0x200, %rax
    44: 0f825d010000             jb 0x18f
    50: 6666666666660f1f840000000000 nopw (%rax, %rax)
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
   320: 62917d48e7041a           vmovntdq %zmm0, (%r10, %r11)
   327: 62917d48e74c1a01         vmovntdq %zmm1, 0x40(%r10, %r11)
   335: 62917d48e7541a02         vmovntdq %zmm2, 0x80(%r10, %r11)
   343: 62917d48e75c1a03         vmovntdq %zmm3, 0xc0(%r10, %r11)
   351: 62917d48e7641a04         vmovntdq %zmm4, 0x100(%r10, %r11)
   359: 62917d48e76c1a05         vmovntdq %zmm5, 0x140(%r10, %r11)
   367: 62917d48e7741a06         vmovntdq %zmm6, 0x180(%r10, %r11)
   375: 62917d48e77c1a07         vmovntdq %zmm7, 0x1c0(%r10, %r11)
   383: 4981c300020000           addq $0x200, %r11
   390: 4c3bd8                   cmpq %rax, %r11
   393: 0f86b1feffff             jbe 0x40
   399: 4c3bd9                   cmpq %rcx, %r11
   402: 0f8346000000             jae 0x1de
   408: 4153                     pushq %r11
   410: 4d8b11                   movq (%r9), %r10
   413: 4c8b1c24                 movq (%rsp), %r11
   417: 498b00                   movq (%r8), %rax
   420: 428a1418                 movb (%rax, %r11), %dl
   424: 8ada                     movb %dl, %bl
   426: 498b4008                 movq 8(%r8), %rax
   430: 428a1418                 movb (%rax, %r11), %dl
   434: 32da                     xorb %dl, %bl
   436: 498b4010                 movq 0x10(%r8), %rax
   440: 428a1418                 movb (%rax, %r11), %dl
   444: 32da                     xorb %dl, %bl
   446: 498b4018                 movq 0x18(%r8), %rax
   450: 428a1418                 movb (%rax, %r11), %dl
   454: 32da                     xorb %dl, %bl
   456: 43881c1a                 movb %bl, (%r10, %r11)
   460: 33db                     xorl %ebx, %ebx
   462: 49ffc3                   incq %r11
   465: 4c3bd9                   cmpq %rcx, %r11
   468: 0f82c7ffffff             jb 0x1a1
   474: 4883c408                 addq $8, %rsp
   478: 415e                     popq %r14
   480: 415d                     popq %r13
   482: 415c                     popq %r12
   484: 5b                       popq %rbx
   485: c5f877                   vzeroupper 
   488: c3                       retq 
   489: 0f0b                     ud2 
   491: 0f0b                     ud2 
   493: 0f0b                     ud2 
   495: 0f0b                     ud2 
   497: 0f0b                     ud2 
   499: 0f0b                     ud2 
   501: 0f0b                     ud2 
   503: 0f0b                     ud2 
   505: 0f0b                     ud2 
   507: 0f0b                     ud2 
   509: 0f0b                     ud2 
   511: cc                       int3 
