; config_t1_533 target=avx2 twin=baseline topo=rs(6,3)/upd{0} op=parity_update U=8
; code_size=832 twins=[] uops_per_byte=0.75
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3d2d030000           leaq 0x32d(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: c4417d6f3f               vmovdqa (%r15), %ymm15
    27: 488bc1                   movq %rcx, %rax
    30: 4883e860                 subq $0x60, %rax
    34: 0f82fd010000             jb 0x225
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5010                 movq 0x10(%r8), %r10
    68: c4817d6f041a             vmovdqa (%r10, %r11), %ymm0
    74: c4817d6f4c1a20           vmovdqa 0x20(%r10, %r11), %ymm1
    81: c4817d6f541a40           vmovdqa 0x40(%r10, %r11), %ymm2
    88: 4d8b5018                 movq 0x18(%r8), %r10
    92: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
    98: c4817d6f641a20           vmovdqa 0x20(%r10, %r11), %ymm4
   105: c4817d6f6c1a40           vmovdqa 0x40(%r10, %r11), %ymm5
   112: 4d8b5020                 movq 0x20(%r8), %r10
   116: c4817d6f341a             vmovdqa (%r10, %r11), %ymm6
   122: c4817d6f7c1a20           vmovdqa 0x20(%r10, %r11), %ymm7
   129: c4017d6f441a40           vmovdqa 0x40(%r10, %r11), %ymm8
   136: 4d8b10                   movq (%r8), %r10
   139: 498b5008                 movq 8(%r8), %rdx
   143: c4017d6f0c1a             vmovdqa (%r10, %r11), %ymm9
   149: c42135ef0c1a             vpxor (%rdx, %r11), %ymm9, %ymm9
   155: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   160: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   166: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   171: c4c17defc1               vpxor %ymm9, %ymm0, %ymm0
   176: c4417d6f7740             vmovdqa 0x40(%r15), %ymm14
   182: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   187: c4417d6fb780000000       vmovdqa 0x80(%r15), %ymm14
   196: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   201: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   206: c4c165efdd               vpxor %ymm13, %ymm3, %ymm3
   211: c4417d6fb7c0000000       vmovdqa 0xc0(%r15), %ymm14
   220: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   225: c4417d6fb700010000       vmovdqa 0x100(%r15), %ymm14
   234: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   239: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   244: c4c14deff5               vpxor %ymm13, %ymm6, %ymm6
   249: c4017d6f4c1a20           vmovdqa 0x20(%r10, %r11), %ymm9
   256: c42135ef4c1a20           vpxor 0x20(%rdx, %r11), %ymm9, %ymm9
   263: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   268: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   274: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   279: c4c175efc9               vpxor %ymm9, %ymm1, %ymm1
   284: c4417d6f7740             vmovdqa 0x40(%r15), %ymm14
   290: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   295: c4417d6fb780000000       vmovdqa 0x80(%r15), %ymm14
   304: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   309: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   314: c4c15defe5               vpxor %ymm13, %ymm4, %ymm4
   319: c4417d6fb7c0000000       vmovdqa 0xc0(%r15), %ymm14
   328: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   333: c4417d6fb700010000       vmovdqa 0x100(%r15), %ymm14
   342: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   347: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   352: c4c145effd               vpxor %ymm13, %ymm7, %ymm7
   357: c4017d6f4c1a40           vmovdqa 0x40(%r10, %r11), %ymm9
   364: c42135ef4c1a40           vpxor 0x40(%rdx, %r11), %ymm9, %ymm9
   371: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   376: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   382: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   387: c4c16defd1               vpxor %ymm9, %ymm2, %ymm2
   392: c4417d6f7740             vmovdqa 0x40(%r15), %ymm14
   398: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   403: c4417d6fb780000000       vmovdqa 0x80(%r15), %ymm14
   412: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   417: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   422: c4c155efed               vpxor %ymm13, %ymm5, %ymm5
   427: c4417d6fb7c0000000       vmovdqa 0xc0(%r15), %ymm14
   436: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   441: c4417d6fb700010000       vmovdqa 0x100(%r15), %ymm14
   450: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   455: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   460: c4413defc5               vpxor %ymm13, %ymm8, %ymm8
   465: 4d8b11                   movq (%r9), %r10
   468: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   474: c4817d7f4c1a20           vmovdqa %ymm1, 0x20(%r10, %r11)
   481: c4817d7f541a40           vmovdqa %ymm2, 0x40(%r10, %r11)
   488: 4d8b5108                 movq 8(%r9), %r10
   492: c4817d7f1c1a             vmovdqa %ymm3, (%r10, %r11)
   498: c4817d7f641a20           vmovdqa %ymm4, 0x20(%r10, %r11)
   505: c4817d7f6c1a40           vmovdqa %ymm5, 0x40(%r10, %r11)
   512: 4d8b5110                 movq 0x10(%r9), %r10
   516: c4817d7f341a             vmovdqa %ymm6, (%r10, %r11)
   522: c4817d7f7c1a20           vmovdqa %ymm7, 0x20(%r10, %r11)
   529: c4017d7f441a40           vmovdqa %ymm8, 0x40(%r10, %r11)
   536: 4983c360                 addq $0x60, %r11
   540: 4c3bd8                   cmpq %rax, %r11
   543: 0f861bfeffff             jbe 0x40
   549: 4c3bd9                   cmpq %rcx, %r11
   552: 0f8307010000             jae 0x335
   558: 4153                     pushq %r11
   560: 498dbf40010000           leaq 0x140(%r15), %rdi
   567: 498db740020000           leaq 0x240(%r15), %rsi
   574: 4d8b11                   movq (%r9), %r10
   577: 4c8b1c24                 movq (%rsp), %r11
   581: 498b4010                 movq 0x10(%r8), %rax
   585: 420fb61c18               movzbl (%rax, %r11), %ebx
   590: 498b00                   movq (%r8), %rax
   593: 420fb61418               movzbl (%rax, %r11), %edx
   598: 498b4008                 movq 8(%r8), %rax
   602: 420fb60418               movzbl (%rax, %r11), %eax
   607: 32d0                     xorb %al, %dl
   609: 32da                     xorb %dl, %bl
   611: 43881c1a                 movb %bl, (%r10, %r11)
   615: 49ffc3                   incq %r11
   618: 4c3bd9                   cmpq %rcx, %r11
   621: 0f82d2ffffff             jb 0x245
   627: 4d8b5108                 movq 8(%r9), %r10
   631: 4c8b1c24                 movq (%rsp), %r11
   635: 498b4018                 movq 0x18(%r8), %rax
   639: 420fb61c18               movzbl (%rax, %r11), %ebx
   644: 498b00                   movq (%r8), %rax
   647: 420fb61418               movzbl (%rax, %r11), %edx
   652: 498b4008                 movq 8(%r8), %rax
   656: 420fb60418               movzbl (%rax, %r11), %eax
   661: 32d0                     xorb %al, %dl
   663: 4885d2                   testq %rdx, %rdx
   666: 0f8422000000             je 0x2c2
   672: 0fb60417                 movzbl (%rdi, %rdx), %eax
   676: 81c0e6000000             addl $0xe6, %eax
   682: 81f8ff000000             cmpl $0xff, %eax
   688: 0f8206000000             jb 0x2bc
   694: 81c001ffffff             addl $0xffffff01, %eax
   700: 0fb60406                 movzbl (%rsi, %rax), %eax
   704: 32d8                     xorb %al, %bl
   706: 43881c1a                 movb %bl, (%r10, %r11)
   710: 49ffc3                   incq %r11
   713: 4c3bd9                   cmpq %rcx, %r11
   716: 0f82a9ffffff             jb 0x27b
   722: 4d8b5110                 movq 0x10(%r9), %r10
   726: 4c8b1c24                 movq (%rsp), %r11
   730: 498b4020                 movq 0x20(%r8), %rax
   734: 420fb61c18               movzbl (%rax, %r11), %ebx
   739: 498b00                   movq (%r8), %rax
   742: 420fb61418               movzbl (%rax, %r11), %edx
   747: 498b4008                 movq 8(%r8), %rax
   751: 420fb60418               movzbl (%rax, %r11), %eax
   756: 32d0                     xorb %al, %dl
   758: 4885d2                   testq %rdx, %rdx
   761: 0f8422000000             je 0x321
   767: 0fb60417                 movzbl (%rdi, %rdx), %eax
   771: 81c0cd000000             addl $0xcd, %eax
   777: 81f8ff000000             cmpl $0xff, %eax
   783: 0f8206000000             jb 0x31b
   789: 81c001ffffff             addl $0xffffff01, %eax
   795: 0fb60406                 movzbl (%rsi, %rax), %eax
   799: 32d8                     xorb %al, %bl
   801: 43881c1a                 movb %bl, (%r10, %r11)
   805: 49ffc3                   incq %r11
   808: 4c3bd9                   cmpq %rcx, %r11
   811: 0f82a9ffffff             jb 0x2da
   817: 4883c408                 addq $8, %rsp
   821: 415f                     popq %r15
   823: 5b                       popq %rbx
   824: c5f877                   vzeroupper 
   827: c3                       retq 
   828: 0f0b                     ud2 
   830: 0f0b                     ud2 
