; config_t1_262 target=avx512-gfni twin=ternlog topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=512 twins=["ternlog"] uops_per_byte=0.25
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3ded010000           leaq 0x1ed(%rip), %r15
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
   214: 4c3bd9                   cmpq %rcx, %r11
   217: 0f83fd000000             jae 0x1dc
   223: 4153                     pushq %r11
   225: 498dbf80000000           leaq 0x80(%r15), %rdi
   232: 498db780010000           leaq 0x180(%r15), %rsi
   239: 4d8b11                   movq (%r9), %r10
   242: 4c8b1c24                 movq (%rsp), %r11
   246: 498b4020                 movq 0x20(%r8), %rax
   250: 420fb61c18               movzbl (%rax, %r11), %ebx
   255: 498b00                   movq (%r8), %rax
   258: 420fb61418               movzbl (%rax, %r11), %edx
   263: 498b4010                 movq 0x10(%r8), %rax
   267: 420fb60418               movzbl (%rax, %r11), %eax
   272: 32d0                     xorb %al, %dl
   274: 32da                     xorb %dl, %bl
   276: 498b4008                 movq 8(%r8), %rax
   280: 420fb61418               movzbl (%rax, %r11), %edx
   285: 498b4018                 movq 0x18(%r8), %rax
   289: 420fb60418               movzbl (%rax, %r11), %eax
   294: 32d0                     xorb %al, %dl
   296: 32da                     xorb %dl, %bl
   298: 43881c1a                 movb %bl, (%r10, %r11)
   302: 49ffc3                   incq %r11
   305: 4c3bd9                   cmpq %rcx, %r11
   308: 0f82bcffffff             jb 0xf6
   314: 4d8b5108                 movq 8(%r9), %r10
   318: 4c8b1c24                 movq (%rsp), %r11
   322: 498b4028                 movq 0x28(%r8), %rax
   326: 420fb61c18               movzbl (%rax, %r11), %ebx
   331: 498b00                   movq (%r8), %rax
   334: 420fb61418               movzbl (%rax, %r11), %edx
   339: 498b4010                 movq 0x10(%r8), %rax
   343: 420fb60418               movzbl (%rax, %r11), %eax
   348: 32d0                     xorb %al, %dl
   350: 4885d2                   testq %rdx, %rdx
   353: 0f8422000000             je 0x189
   359: 0fb60417                 movzbl (%rdi, %rdx), %eax
   363: 81c0e6000000             addl $0xe6, %eax
   369: 81f8ff000000             cmpl $0xff, %eax
   375: 0f8206000000             jb 0x183
   381: 81c001ffffff             addl $0xffffff01, %eax
   387: 0fb60406                 movzbl (%rsi, %rax), %eax
   391: 32d8                     xorb %al, %bl
   393: 498b4008                 movq 8(%r8), %rax
   397: 420fb61418               movzbl (%rax, %r11), %edx
   402: 498b4018                 movq 0x18(%r8), %rax
   406: 420fb60418               movzbl (%rax, %r11), %eax
   411: 32d0                     xorb %al, %dl
   413: 4885d2                   testq %rdx, %rdx
   416: 0f8422000000             je 0x1c8
   422: 0fb60417                 movzbl (%rdi, %rdx), %eax
   426: 81c0cd000000             addl $0xcd, %eax
   432: 81f8ff000000             cmpl $0xff, %eax
   438: 0f8206000000             jb 0x1c2
   444: 81c001ffffff             addl $0xffffff01, %eax
   450: 0fb60406                 movzbl (%rsi, %rax), %eax
   454: 32d8                     xorb %al, %bl
   456: 43881c1a                 movb %bl, (%r10, %r11)
   460: 49ffc3                   incq %r11
   463: 4c3bd9                   cmpq %rcx, %r11
   466: 0f826affffff             jb 0x142
   472: 4883c408                 addq $8, %rsp
   476: 415f                     popq %r15
   478: 5b                       popq %rbx
   479: c5f877                   vzeroupper 
   482: c3                       retq 
   483: 0f0b                     ud2 
   485: 0f0b                     ud2 
   487: 0f0b                     ud2 
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
