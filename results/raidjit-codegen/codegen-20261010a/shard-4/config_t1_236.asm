; config_t1_236 target=avx512-gfni twin=ternlog topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=512 twins=["ternlog"] uops_per_byte=0.234375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3ded010000           leaq 0x1ed(%rip), %r15
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
   208: 4c3bd9                   cmpq %rcx, %r11
   211: 0f8307010000             jae 0x1e0
   217: 4153                     pushq %r11
   219: 498dbf80000000           leaq 0x80(%r15), %rdi
   226: 498db780010000           leaq 0x180(%r15), %rsi
   233: 4d8b11                   movq (%r9), %r10
   236: 4c8b1c24                 movq (%rsp), %r11
   240: 498b4010                 movq 0x10(%r8), %rax
   244: 420fb61c18               movzbl (%rax, %r11), %ebx
   249: 498b00                   movq (%r8), %rax
   252: 420fb61418               movzbl (%rax, %r11), %edx
   257: 498b4008                 movq 8(%r8), %rax
   261: 420fb60418               movzbl (%rax, %r11), %eax
   266: 32d0                     xorb %al, %dl
   268: 32da                     xorb %dl, %bl
   270: 43881c1a                 movb %bl, (%r10, %r11)
   274: 49ffc3                   incq %r11
   277: 4c3bd9                   cmpq %rcx, %r11
   280: 0f82d2ffffff             jb 0xf0
   286: 4d8b5108                 movq 8(%r9), %r10
   290: 4c8b1c24                 movq (%rsp), %r11
   294: 498b4018                 movq 0x18(%r8), %rax
   298: 420fb61c18               movzbl (%rax, %r11), %ebx
   303: 498b00                   movq (%r8), %rax
   306: 420fb61418               movzbl (%rax, %r11), %edx
   311: 498b4008                 movq 8(%r8), %rax
   315: 420fb60418               movzbl (%rax, %r11), %eax
   320: 32d0                     xorb %al, %dl
   322: 4885d2                   testq %rdx, %rdx
   325: 0f8422000000             je 0x16d
   331: 0fb60417                 movzbl (%rdi, %rdx), %eax
   335: 81c0e6000000             addl $0xe6, %eax
   341: 81f8ff000000             cmpl $0xff, %eax
   347: 0f8206000000             jb 0x167
   353: 81c001ffffff             addl $0xffffff01, %eax
   359: 0fb60406                 movzbl (%rsi, %rax), %eax
   363: 32d8                     xorb %al, %bl
   365: 43881c1a                 movb %bl, (%r10, %r11)
   369: 49ffc3                   incq %r11
   372: 4c3bd9                   cmpq %rcx, %r11
   375: 0f82a9ffffff             jb 0x126
   381: 4d8b5110                 movq 0x10(%r9), %r10
   385: 4c8b1c24                 movq (%rsp), %r11
   389: 498b4020                 movq 0x20(%r8), %rax
   393: 420fb61c18               movzbl (%rax, %r11), %ebx
   398: 498b00                   movq (%r8), %rax
   401: 420fb61418               movzbl (%rax, %r11), %edx
   406: 498b4008                 movq 8(%r8), %rax
   410: 420fb60418               movzbl (%rax, %r11), %eax
   415: 32d0                     xorb %al, %dl
   417: 4885d2                   testq %rdx, %rdx
   420: 0f8422000000             je 0x1cc
   426: 0fb60417                 movzbl (%rdi, %rdx), %eax
   430: 81c0cd000000             addl $0xcd, %eax
   436: 81f8ff000000             cmpl $0xff, %eax
   442: 0f8206000000             jb 0x1c6
   448: 81c001ffffff             addl $0xffffff01, %eax
   454: 0fb60406                 movzbl (%rsi, %rax), %eax
   458: 32d8                     xorb %al, %bl
   460: 43881c1a                 movb %bl, (%r10, %r11)
   464: 49ffc3                   incq %r11
   467: 4c3bd9                   cmpq %rcx, %r11
   470: 0f82a9ffffff             jb 0x185
   476: 4883c408                 addq $8, %rsp
   480: 415f                     popq %r15
   482: 5b                       popq %rbx
   483: c5f877                   vzeroupper 
   486: c3                       retq 
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
