; config_t1_171 target=avx512-gfni twin=ternlog topo=rs(6,3)/rec{0} op=recover U=8
; code_size=768 twins=["ternlog"] uops_per_byte=0.126953125
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
    31: 4d8b7020                 movq 0x20(%r8), %r14
    35: 4d33db                   xorq %r11, %r11
    38: 488bc1                   movq %rcx, %rax
    41: 4881e800020000           subq $0x200, %rax
    48: 0f8223020000             jb 0x259
    54: 66660f1f840000000000     nopw (%rax, %rax)
    64: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
    70: 62f17548efc9             vpxord %zmm1, %zmm1, %zmm1
    76: 62f16d48efd2             vpxord %zmm2, %zmm2, %zmm2
    82: 62f16548efdb             vpxord %zmm3, %zmm3, %zmm3
    88: 62f15d48efe4             vpxord %zmm4, %zmm4, %zmm4
    94: 62f15548efed             vpxord %zmm5, %zmm5, %zmm5
   100: 62f14d48eff6             vpxord %zmm6, %zmm6, %zmm6
   106: 62f14548efff             vpxord %zmm7, %zmm7, %zmm7
   112: 62317d486f041e           vmovdqa32 (%rsi, %r11), %zmm8
   119: 62b33d4825041f96         vpternlogd $0x96, (%rdi, %r11), %zmm8, %zmm0
   127: 62317d486f441e01         vmovdqa32 0x40(%rsi, %r11), %zmm8
   135: 62b33d48254c1f0196       vpternlogd $0x96, 0x40(%rdi, %r11), %zmm8, %zmm1
   144: 62317d486f441e02         vmovdqa32 0x80(%rsi, %r11), %zmm8
   152: 62b33d4825541f0296       vpternlogd $0x96, 0x80(%rdi, %r11), %zmm8, %zmm2
   161: 62317d486f441e03         vmovdqa32 0xc0(%rsi, %r11), %zmm8
   169: 62b33d48255c1f0396       vpternlogd $0x96, 0xc0(%rdi, %r11), %zmm8, %zmm3
   178: 62317d486f441e04         vmovdqa32 0x100(%rsi, %r11), %zmm8
   186: 62b33d4825641f0496       vpternlogd $0x96, 0x100(%rdi, %r11), %zmm8, %zmm4
   195: 62317d486f441e05         vmovdqa32 0x140(%rsi, %r11), %zmm8
   203: 62b33d48256c1f0596       vpternlogd $0x96, 0x140(%rdi, %r11), %zmm8, %zmm5
   212: 62317d486f441e06         vmovdqa32 0x180(%rsi, %r11), %zmm8
   220: 62b33d4825741f0696       vpternlogd $0x96, 0x180(%rdi, %r11), %zmm8, %zmm6
   229: 62317d486f441e07         vmovdqa32 0x1c0(%rsi, %r11), %zmm8
   237: 62b33d48257c1f0796       vpternlogd $0x96, 0x1c0(%rdi, %r11), %zmm8, %zmm7
   246: 62117d486f041c           vmovdqa32 (%r12, %r11), %zmm8
   253: 62933d4825441d0096       vpternlogd $0x96, (%r13, %r11), %zmm8, %zmm0
   262: 62117d486f441c01         vmovdqa32 0x40(%r12, %r11), %zmm8
   270: 62933d48254c1d0196       vpternlogd $0x96, 0x40(%r13, %r11), %zmm8, %zmm1
   279: 62117d486f441c02         vmovdqa32 0x80(%r12, %r11), %zmm8
   287: 62933d4825541d0296       vpternlogd $0x96, 0x80(%r13, %r11), %zmm8, %zmm2
   296: 62117d486f441c03         vmovdqa32 0xc0(%r12, %r11), %zmm8
   304: 62933d48255c1d0396       vpternlogd $0x96, 0xc0(%r13, %r11), %zmm8, %zmm3
   313: 62117d486f441c04         vmovdqa32 0x100(%r12, %r11), %zmm8
   321: 62933d4825641d0496       vpternlogd $0x96, 0x100(%r13, %r11), %zmm8, %zmm4
   330: 62117d486f441c05         vmovdqa32 0x140(%r12, %r11), %zmm8
   338: 62933d48256c1d0596       vpternlogd $0x96, 0x140(%r13, %r11), %zmm8, %zmm5
   347: 62117d486f441c06         vmovdqa32 0x180(%r12, %r11), %zmm8
   355: 62933d4825741d0696       vpternlogd $0x96, 0x180(%r13, %r11), %zmm8, %zmm6
   364: 62117d486f441c07         vmovdqa32 0x1c0(%r12, %r11), %zmm8
   372: 62933d48257c1d0796       vpternlogd $0x96, 0x1c0(%r13, %r11), %zmm8, %zmm7
   381: 498b5028                 movq 0x28(%r8), %rdx
   385: 62117d486f041e           vmovdqa32 (%r14, %r11), %zmm8
   392: 62b33d4825041a96         vpternlogd $0x96, (%rdx, %r11), %zmm8, %zmm0
   400: 62117d486f441e01         vmovdqa32 0x40(%r14, %r11), %zmm8
   408: 62b33d48254c1a0196       vpternlogd $0x96, 0x40(%rdx, %r11), %zmm8, %zmm1
   417: 62117d486f441e02         vmovdqa32 0x80(%r14, %r11), %zmm8
   425: 62b33d4825541a0296       vpternlogd $0x96, 0x80(%rdx, %r11), %zmm8, %zmm2
   434: 62117d486f441e03         vmovdqa32 0xc0(%r14, %r11), %zmm8
   442: 62b33d48255c1a0396       vpternlogd $0x96, 0xc0(%rdx, %r11), %zmm8, %zmm3
   451: 62117d486f441e04         vmovdqa32 0x100(%r14, %r11), %zmm8
   459: 62b33d4825641a0496       vpternlogd $0x96, 0x100(%rdx, %r11), %zmm8, %zmm4
   468: 62117d486f441e05         vmovdqa32 0x140(%r14, %r11), %zmm8
   476: 62b33d48256c1a0596       vpternlogd $0x96, 0x140(%rdx, %r11), %zmm8, %zmm5
   485: 62117d486f441e06         vmovdqa32 0x180(%r14, %r11), %zmm8
   493: 62b33d4825741a0696       vpternlogd $0x96, 0x180(%rdx, %r11), %zmm8, %zmm6
   502: 62117d486f441e07         vmovdqa32 0x1c0(%r14, %r11), %zmm8
   510: 62b33d48257c1a0796       vpternlogd $0x96, 0x1c0(%rdx, %r11), %zmm8, %zmm7
   519: 4d8b11                   movq (%r9), %r10
   522: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   529: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
   537: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
   545: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
   553: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
   561: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
   569: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
   577: 62917d487f7c1a07         vmovdqa32 %zmm7, 0x1c0(%r10, %r11)
   585: 4981c300020000           addq $0x200, %r11
   592: 4c3bd8                   cmpq %rax, %r11
   595: 0f86e7fdffff             jbe 0x40
   601: 4c3bd9                   cmpq %rcx, %r11
   604: 0f835a000000             jae 0x2bc
   610: 4153                     pushq %r11
   612: 4d8b11                   movq (%r9), %r10
   615: 4c8b1c24                 movq (%rsp), %r11
   619: 498b00                   movq (%r8), %rax
   622: 428a1418                 movb (%rax, %r11), %dl
   626: 8ada                     movb %dl, %bl
   628: 498b4008                 movq 8(%r8), %rax
   632: 428a1418                 movb (%rax, %r11), %dl
   636: 32da                     xorb %dl, %bl
   638: 498b4010                 movq 0x10(%r8), %rax
   642: 428a1418                 movb (%rax, %r11), %dl
   646: 32da                     xorb %dl, %bl
   648: 498b4018                 movq 0x18(%r8), %rax
   652: 428a1418                 movb (%rax, %r11), %dl
   656: 32da                     xorb %dl, %bl
   658: 498b4020                 movq 0x20(%r8), %rax
   662: 428a1418                 movb (%rax, %r11), %dl
   666: 32da                     xorb %dl, %bl
   668: 498b4028                 movq 0x28(%r8), %rax
   672: 428a1418                 movb (%rax, %r11), %dl
   676: 32da                     xorb %dl, %bl
   678: 43881c1a                 movb %bl, (%r10, %r11)
   682: 33db                     xorl %ebx, %ebx
   684: 49ffc3                   incq %r11
   687: 4c3bd9                   cmpq %rcx, %r11
   690: 0f82b3ffffff             jb 0x26b
   696: 4883c408                 addq $8, %rsp
   700: 415e                     popq %r14
   702: 415d                     popq %r13
   704: 415c                     popq %r12
   706: 5b                       popq %rbx
   707: c5f877                   vzeroupper 
   710: c3                       retq 
   711: 0f0b                     ud2 
   713: 0f0b                     ud2 
   715: 0f0b                     ud2 
   717: 0f0b                     ud2 
   719: 0f0b                     ud2 
   721: 0f0b                     ud2 
   723: 0f0b                     ud2 
   725: 0f0b                     ud2 
   727: 0f0b                     ud2 
   729: 0f0b                     ud2 
   731: 0f0b                     ud2 
   733: 0f0b                     ud2 
   735: 0f0b                     ud2 
   737: 0f0b                     ud2 
   739: 0f0b                     ud2 
   741: 0f0b                     ud2 
   743: 0f0b                     ud2 
   745: 0f0b                     ud2 
   747: 0f0b                     ud2 
   749: 0f0b                     ud2 
   751: 0f0b                     ud2 
   753: 0f0b                     ud2 
   755: 0f0b                     ud2 
   757: 0f0b                     ud2 
   759: 0f0b                     ud2 
   761: 0f0b                     ud2 
   763: 0f0b                     ud2 
   765: 0f0b                     ud2 
   767: cc                       int3 
