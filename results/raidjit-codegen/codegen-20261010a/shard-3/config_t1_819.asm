; config_t1_819 target=sse twin=baseline topo=rs(6,3)/upd{0} op=parity_update U=8
; code_size=960 twins=[] uops_per_byte=1.9375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3dad030000           leaq 0x3ad(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: 66450f6f3f               movdqa (%r15), %xmm15
    27: 488bc1                   movq %rcx, %rax
    30: 4883e830                 subq $0x30, %rax
    34: 0f8272020000             jb 0x29a
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5010                 movq 0x10(%r8), %r10
    68: 66430f6f041a             movdqa (%r10, %r11), %xmm0
    74: 66430f6f4c1a10           movdqa 0x10(%r10, %r11), %xmm1
    81: 66430f6f541a20           movdqa 0x20(%r10, %r11), %xmm2
    88: 4d8b5018                 movq 0x18(%r8), %r10
    92: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
    98: 66430f6f641a10           movdqa 0x10(%r10, %r11), %xmm4
   105: 66430f6f6c1a20           movdqa 0x20(%r10, %r11), %xmm5
   112: 4d8b5020                 movq 0x20(%r8), %r10
   116: 66430f6f341a             movdqa (%r10, %r11), %xmm6
   122: 66430f6f7c1a10           movdqa 0x10(%r10, %r11), %xmm7
   129: 66470f6f441a20           movdqa 0x20(%r10, %r11), %xmm8
   136: 4d8b10                   movq (%r8), %r10
   139: 498b5008                 movq 8(%r8), %rdx
   143: 66470f6f0c1a             movdqa (%r10, %r11), %xmm9
   149: 66460f6f241a             movdqa (%rdx, %r11), %xmm12
   155: 66450fefcc               pxor %xmm12, %xmm9
   160: 66450f6fd1               movdqa %xmm9, %xmm10
   165: 66450fdbd7               pand %xmm15, %xmm10
   170: 66450f6fd9               movdqa %xmm9, %xmm11
   175: 66410f71d304             psrlw $4, %xmm11
   181: 66450fdbdf               pand %xmm15, %xmm11
   186: 66410fefc1               pxor %xmm9, %xmm0
   191: 66450f6f7740             movdqa 0x40(%r15), %xmm14
   197: 66450f6fe6               movdqa %xmm14, %xmm12
   202: 66450f3800e2             pshufb %xmm10, %xmm12
   208: 66450f6fb780000000       movdqa 0x80(%r15), %xmm14
   217: 66450f6fee               movdqa %xmm14, %xmm13
   222: 66450f3800eb             pshufb %xmm11, %xmm13
   228: 66410fefdc               pxor %xmm12, %xmm3
   233: 66410fefdd               pxor %xmm13, %xmm3
   238: 66450f6fb7c0000000       movdqa 0xc0(%r15), %xmm14
   247: 66450f6fe6               movdqa %xmm14, %xmm12
   252: 66450f3800e2             pshufb %xmm10, %xmm12
   258: 66450f6fb700010000       movdqa 0x100(%r15), %xmm14
   267: 66450f6fee               movdqa %xmm14, %xmm13
   272: 66450f3800eb             pshufb %xmm11, %xmm13
   278: 66410feff4               pxor %xmm12, %xmm6
   283: 66410feff5               pxor %xmm13, %xmm6
   288: 66470f6f4c1a10           movdqa 0x10(%r10, %r11), %xmm9
   295: 66460f6f641a10           movdqa 0x10(%rdx, %r11), %xmm12
   302: 66450fefcc               pxor %xmm12, %xmm9
   307: 66450f6fd1               movdqa %xmm9, %xmm10
   312: 66450fdbd7               pand %xmm15, %xmm10
   317: 66450f6fd9               movdqa %xmm9, %xmm11
   322: 66410f71d304             psrlw $4, %xmm11
   328: 66450fdbdf               pand %xmm15, %xmm11
   333: 66410fefc9               pxor %xmm9, %xmm1
   338: 66450f6f7740             movdqa 0x40(%r15), %xmm14
   344: 66450f6fe6               movdqa %xmm14, %xmm12
   349: 66450f3800e2             pshufb %xmm10, %xmm12
   355: 66450f6fb780000000       movdqa 0x80(%r15), %xmm14
   364: 66450f6fee               movdqa %xmm14, %xmm13
   369: 66450f3800eb             pshufb %xmm11, %xmm13
   375: 66410fefe4               pxor %xmm12, %xmm4
   380: 66410fefe5               pxor %xmm13, %xmm4
   385: 66450f6fb7c0000000       movdqa 0xc0(%r15), %xmm14
   394: 66450f6fe6               movdqa %xmm14, %xmm12
   399: 66450f3800e2             pshufb %xmm10, %xmm12
   405: 66450f6fb700010000       movdqa 0x100(%r15), %xmm14
   414: 66450f6fee               movdqa %xmm14, %xmm13
   419: 66450f3800eb             pshufb %xmm11, %xmm13
   425: 66410feffc               pxor %xmm12, %xmm7
   430: 66410feffd               pxor %xmm13, %xmm7
   435: 66470f6f4c1a20           movdqa 0x20(%r10, %r11), %xmm9
   442: 66460f6f641a20           movdqa 0x20(%rdx, %r11), %xmm12
   449: 66450fefcc               pxor %xmm12, %xmm9
   454: 66450f6fd1               movdqa %xmm9, %xmm10
   459: 66450fdbd7               pand %xmm15, %xmm10
   464: 66450f6fd9               movdqa %xmm9, %xmm11
   469: 66410f71d304             psrlw $4, %xmm11
   475: 66450fdbdf               pand %xmm15, %xmm11
   480: 66410fefd1               pxor %xmm9, %xmm2
   485: 66450f6f7740             movdqa 0x40(%r15), %xmm14
   491: 66450f6fe6               movdqa %xmm14, %xmm12
   496: 66450f3800e2             pshufb %xmm10, %xmm12
   502: 66450f6fb780000000       movdqa 0x80(%r15), %xmm14
   511: 66450f6fee               movdqa %xmm14, %xmm13
   516: 66450f3800eb             pshufb %xmm11, %xmm13
   522: 66410fefec               pxor %xmm12, %xmm5
   527: 66410fefed               pxor %xmm13, %xmm5
   532: 66450f6fb7c0000000       movdqa 0xc0(%r15), %xmm14
   541: 66450f6fe6               movdqa %xmm14, %xmm12
   546: 66450f3800e2             pshufb %xmm10, %xmm12
   552: 66450f6fb700010000       movdqa 0x100(%r15), %xmm14
   561: 66450f6fee               movdqa %xmm14, %xmm13
   566: 66450f3800eb             pshufb %xmm11, %xmm13
   572: 66450fefc4               pxor %xmm12, %xmm8
   577: 66450fefc5               pxor %xmm13, %xmm8
   582: 4d8b11                   movq (%r9), %r10
   585: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   591: 66430f7f4c1a10           movdqa %xmm1, 0x10(%r10, %r11)
   598: 66430f7f541a20           movdqa %xmm2, 0x20(%r10, %r11)
   605: 4d8b5108                 movq 8(%r9), %r10
   609: 66430f7f1c1a             movdqa %xmm3, (%r10, %r11)
   615: 66430f7f641a10           movdqa %xmm4, 0x10(%r10, %r11)
   622: 66430f7f6c1a20           movdqa %xmm5, 0x20(%r10, %r11)
   629: 4d8b5110                 movq 0x10(%r9), %r10
   633: 66430f7f341a             movdqa %xmm6, (%r10, %r11)
   639: 66430f7f7c1a10           movdqa %xmm7, 0x10(%r10, %r11)
   646: 66470f7f441a20           movdqa %xmm8, 0x20(%r10, %r11)
   653: 4983c330                 addq $0x30, %r11
   657: 4c3bd8                   cmpq %rax, %r11
   660: 0f86a6fdffff             jbe 0x40
   666: 4c3bd9                   cmpq %rcx, %r11
   669: 0f8307010000             jae 0x3aa
   675: 4153                     pushq %r11
   677: 498dbf40010000           leaq 0x140(%r15), %rdi
   684: 498db740020000           leaq 0x240(%r15), %rsi
   691: 4d8b11                   movq (%r9), %r10
   694: 4c8b1c24                 movq (%rsp), %r11
   698: 498b4010                 movq 0x10(%r8), %rax
   702: 420fb61c18               movzbl (%rax, %r11), %ebx
   707: 498b00                   movq (%r8), %rax
   710: 420fb61418               movzbl (%rax, %r11), %edx
   715: 498b4008                 movq 8(%r8), %rax
   719: 420fb60418               movzbl (%rax, %r11), %eax
   724: 32d0                     xorb %al, %dl
   726: 32da                     xorb %dl, %bl
   728: 43881c1a                 movb %bl, (%r10, %r11)
   732: 49ffc3                   incq %r11
   735: 4c3bd9                   cmpq %rcx, %r11
   738: 0f82d2ffffff             jb 0x2ba
   744: 4d8b5108                 movq 8(%r9), %r10
   748: 4c8b1c24                 movq (%rsp), %r11
   752: 498b4018                 movq 0x18(%r8), %rax
   756: 420fb61c18               movzbl (%rax, %r11), %ebx
   761: 498b00                   movq (%r8), %rax
   764: 420fb61418               movzbl (%rax, %r11), %edx
   769: 498b4008                 movq 8(%r8), %rax
   773: 420fb60418               movzbl (%rax, %r11), %eax
   778: 32d0                     xorb %al, %dl
   780: 4885d2                   testq %rdx, %rdx
   783: 0f8422000000             je 0x337
   789: 0fb60417                 movzbl (%rdi, %rdx), %eax
   793: 81c0e6000000             addl $0xe6, %eax
   799: 81f8ff000000             cmpl $0xff, %eax
   805: 0f8206000000             jb 0x331
   811: 81c001ffffff             addl $0xffffff01, %eax
   817: 0fb60406                 movzbl (%rsi, %rax), %eax
   821: 32d8                     xorb %al, %bl
   823: 43881c1a                 movb %bl, (%r10, %r11)
   827: 49ffc3                   incq %r11
   830: 4c3bd9                   cmpq %rcx, %r11
   833: 0f82a9ffffff             jb 0x2f0
   839: 4d8b5110                 movq 0x10(%r9), %r10
   843: 4c8b1c24                 movq (%rsp), %r11
   847: 498b4020                 movq 0x20(%r8), %rax
   851: 420fb61c18               movzbl (%rax, %r11), %ebx
   856: 498b00                   movq (%r8), %rax
   859: 420fb61418               movzbl (%rax, %r11), %edx
   864: 498b4008                 movq 8(%r8), %rax
   868: 420fb60418               movzbl (%rax, %r11), %eax
   873: 32d0                     xorb %al, %dl
   875: 4885d2                   testq %rdx, %rdx
   878: 0f8422000000             je 0x396
   884: 0fb60417                 movzbl (%rdi, %rdx), %eax
   888: 81c0cd000000             addl $0xcd, %eax
   894: 81f8ff000000             cmpl $0xff, %eax
   900: 0f8206000000             jb 0x390
   906: 81c001ffffff             addl $0xffffff01, %eax
   912: 0fb60406                 movzbl (%rsi, %rax), %eax
   916: 32d8                     xorb %al, %bl
   918: 43881c1a                 movb %bl, (%r10, %r11)
   922: 49ffc3                   incq %r11
   925: 4c3bd9                   cmpq %rcx, %r11
   928: 0f82a9ffffff             jb 0x34f
   934: 4883c408                 addq $8, %rsp
   938: 415f                     popq %r15
   940: 5b                       popq %rbx
   941: c3                       retq 
   942: 0f0b                     ud2 
   944: 0f0b                     ud2 
   946: 0f0b                     ud2 
   948: 0f0b                     ud2 
   950: 0f0b                     ud2 
   952: 0f0b                     ud2 
   954: 0f0b                     ud2 
   956: 0f0b                     ud2 
   958: 0f0b                     ud2 
