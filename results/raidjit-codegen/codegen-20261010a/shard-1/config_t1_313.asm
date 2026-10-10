; config_t1_313 target=avx2 twin=nt topo=raid6(8)/enc op=encode U=1
; code_size=1152 twins=["nt"] uops_per_byte=2.625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d67040000           leaq 0x467(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: c4417d6f07               vmovdqa (%r15), %ymm8
    52: 488bc1                   movq %rcx, %rax
    55: 4883e820                 subq $0x20, %rax
    59: 0f82fe010000             jb 0x23f
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   102: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   107: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   112: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   117: c5fd6fc2                 vmovdqa %ymm2, %ymm0
   121: c4c17d6f7f40             vmovdqa 0x40(%r15), %ymm7
   127: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   132: c4c17d6fbf80000000       vmovdqa 0x80(%r15), %ymm7
   141: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   146: c5d5efce                 vpxor %ymm6, %ymm5, %ymm1
   150: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   156: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   161: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   166: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   171: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   175: c4c17d6fbfc0000000       vmovdqa 0xc0(%r15), %ymm7
   184: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   189: c4c17d6fbf00010000       vmovdqa 0x100(%r15), %ymm7
   198: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   203: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   207: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   211: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   217: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   222: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   227: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   232: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   236: c4c17d6fbf40010000       vmovdqa 0x140(%r15), %ymm7
   245: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   250: c4c17d6fbf80010000       vmovdqa 0x180(%r15), %ymm7
   259: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   264: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   268: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   272: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   279: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   284: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   289: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   294: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   298: c4c17d6fbfc0010000       vmovdqa 0x1c0(%r15), %ymm7
   307: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   312: c4c17d6fbf00020000       vmovdqa 0x200(%r15), %ymm7
   321: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   326: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   330: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   334: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   340: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   345: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   350: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   355: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   359: c4c17d6fbf40020000       vmovdqa 0x240(%r15), %ymm7
   368: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   373: c4c17d6fbf80020000       vmovdqa 0x280(%r15), %ymm7
   382: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   387: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   391: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   395: 4d8b5028                 movq 0x28(%r8), %r10
   399: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   405: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   410: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   415: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   420: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   424: c4c17d6fbfc0020000       vmovdqa 0x2c0(%r15), %ymm7
   433: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   438: c4c17d6fbf00030000       vmovdqa 0x300(%r15), %ymm7
   447: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   452: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   456: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   460: 4d8b5030                 movq 0x30(%r8), %r10
   464: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   470: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   475: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   480: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   485: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   489: c4c17d6fbf40030000       vmovdqa 0x340(%r15), %ymm7
   498: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   503: c4c17d6fbf80030000       vmovdqa 0x380(%r15), %ymm7
   512: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   517: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   521: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   525: 4d8b5038                 movq 0x38(%r8), %r10
   529: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   535: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   539: c5f5efca                 vpxor %ymm2, %ymm1, %ymm1
   543: 4d8b11                   movq (%r9), %r10
   546: c4817de7041a             vmovntdq %ymm0, (%r10, %r11)
   552: 4d8b5108                 movq 8(%r9), %r10
   556: c4817de70c1a             vmovntdq %ymm1, (%r10, %r11)
   562: 4983c320                 addq $0x20, %r11
   566: 4c3bd8                   cmpq %rax, %r11
   569: 0f8621feffff             jbe 0x60
   575: 4c3bd9                   cmpq %rcx, %r11
   578: 0f830f020000             jae 0x457
   584: 4153                     pushq %r11
   586: 498dbfc0030000           leaq 0x3c0(%r15), %rdi
   593: 498db7c0040000           leaq 0x4c0(%r15), %rsi
   600: 4d8b11                   movq (%r9), %r10
   603: 4c8b1c24                 movq (%rsp), %r11
   607: 33db                     xorl %ebx, %ebx
   609: 498b00                   movq (%r8), %rax
   612: 420fb61418               movzbl (%rax, %r11), %edx
   617: 32da                     xorb %dl, %bl
   619: 498b4008                 movq 8(%r8), %rax
   623: 420fb61418               movzbl (%rax, %r11), %edx
   628: 32da                     xorb %dl, %bl
   630: 498b4010                 movq 0x10(%r8), %rax
   634: 420fb61418               movzbl (%rax, %r11), %edx
   639: 32da                     xorb %dl, %bl
   641: 498b4018                 movq 0x18(%r8), %rax
   645: 420fb61418               movzbl (%rax, %r11), %edx
   650: 32da                     xorb %dl, %bl
   652: 498b4020                 movq 0x20(%r8), %rax
   656: 420fb61418               movzbl (%rax, %r11), %edx
   661: 32da                     xorb %dl, %bl
   663: 498b4028                 movq 0x28(%r8), %rax
   667: 420fb61418               movzbl (%rax, %r11), %edx
   672: 32da                     xorb %dl, %bl
   674: 498b4030                 movq 0x30(%r8), %rax
   678: 420fb61418               movzbl (%rax, %r11), %edx
   683: 32da                     xorb %dl, %bl
   685: 498b4038                 movq 0x38(%r8), %rax
   689: 420fb61418               movzbl (%rax, %r11), %edx
   694: 32da                     xorb %dl, %bl
   696: 43881c1a                 movb %bl, (%r10, %r11)
   700: 33db                     xorl %ebx, %ebx
   702: 49ffc3                   incq %r11
   705: 4c3bd9                   cmpq %rcx, %r11
   708: 0f8297ffffff             jb 0x261
   714: 4d8b5108                 movq 8(%r9), %r10
   718: 4c8b1c24                 movq (%rsp), %r11
   722: 33db                     xorl %ebx, %ebx
   724: 498b00                   movq (%r8), %rax
   727: 420fb61418               movzbl (%rax, %r11), %edx
   732: 4885d2                   testq %rdx, %rdx
   735: 0f8422000000             je 0x307
   741: 0fb60417                 movzbl (%rdi, %rdx), %eax
   745: 81c0e6000000             addl $0xe6, %eax
   751: 81f8ff000000             cmpl $0xff, %eax
   757: 0f8206000000             jb 0x301
   763: 81c001ffffff             addl $0xffffff01, %eax
   769: 0fb60406                 movzbl (%rsi, %rax), %eax
   773: 32d8                     xorb %al, %bl
   775: 498b4008                 movq 8(%r8), %rax
   779: 420fb61418               movzbl (%rax, %r11), %edx
   784: 4885d2                   testq %rdx, %rdx
   787: 0f8422000000             je 0x33b
   793: 0fb60417                 movzbl (%rdi, %rdx), %eax
   797: 81c0cd000000             addl $0xcd, %eax
   803: 81f8ff000000             cmpl $0xff, %eax
   809: 0f8206000000             jb 0x335
   815: 81c001ffffff             addl $0xffffff01, %eax
   821: 0fb60406                 movzbl (%rsi, %rax), %eax
   825: 32d8                     xorb %al, %bl
   827: 498b4010                 movq 0x10(%r8), %rax
   831: 420fb61418               movzbl (%rax, %r11), %edx
   836: 4885d2                   testq %rdx, %rdx
   839: 0f8422000000             je 0x36f
   845: 0fb60417                 movzbl (%rdi, %rdx), %eax
   849: 81c0e7000000             addl $0xe7, %eax
   855: 81f8ff000000             cmpl $0xff, %eax
   861: 0f8206000000             jb 0x369
   867: 81c001ffffff             addl $0xffffff01, %eax
   873: 0fb60406                 movzbl (%rsi, %rax), %eax
   877: 32d8                     xorb %al, %bl
   879: 498b4018                 movq 0x18(%r8), %rax
   883: 420fb61418               movzbl (%rax, %r11), %edx
   888: 4885d2                   testq %rdx, %rdx
   891: 0f8422000000             je 0x3a3
   897: 0fb60417                 movzbl (%rdi, %rdx), %eax
   901: 81c0ce000000             addl $0xce, %eax
   907: 81f8ff000000             cmpl $0xff, %eax
   913: 0f8206000000             jb 0x39d
   919: 81c001ffffff             addl $0xffffff01, %eax
   925: 0fb60406                 movzbl (%rsi, %rax), %eax
   929: 32d8                     xorb %al, %bl
   931: 498b4020                 movq 0x20(%r8), %rax
   935: 420fb61418               movzbl (%rax, %r11), %edx
   940: 4885d2                   testq %rdx, %rdx
   943: 0f841f000000             je 0x3d4
   949: 0fb60417                 movzbl (%rdi, %rdx), %eax
   953: 83c04b                   addl $0x4b, %eax
   956: 81f8ff000000             cmpl $0xff, %eax
   962: 0f8206000000             jb 0x3ce
   968: 81c001ffffff             addl $0xffffff01, %eax
   974: 0fb60406                 movzbl (%rsi, %rax), %eax
   978: 32d8                     xorb %al, %bl
   980: 498b4028                 movq 0x28(%r8), %rax
   984: 420fb61418               movzbl (%rax, %r11), %edx
   989: 4885d2                   testq %rdx, %rdx
   992: 0f841f000000             je 0x405
   998: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1002: 83c032                   addl $0x32, %eax
  1005: 81f8ff000000             cmpl $0xff, %eax
  1011: 0f8206000000             jb 0x3ff
  1017: 81c001ffffff             addl $0xffffff01, %eax
  1023: 0fb60406                 movzbl (%rsi, %rax), %eax
  1027: 32d8                     xorb %al, %bl
  1029: 498b4030                 movq 0x30(%r8), %rax
  1033: 420fb61418               movzbl (%rax, %r11), %edx
  1038: 4885d2                   testq %rdx, %rdx
  1041: 0f841f000000             je 0x436
  1047: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1051: 83c019                   addl $0x19, %eax
  1054: 81f8ff000000             cmpl $0xff, %eax
  1060: 0f8206000000             jb 0x430
  1066: 81c001ffffff             addl $0xffffff01, %eax
  1072: 0fb60406                 movzbl (%rsi, %rax), %eax
  1076: 32d8                     xorb %al, %bl
  1078: 498b4038                 movq 0x38(%r8), %rax
  1082: 420fb61418               movzbl (%rax, %r11), %edx
  1087: 32da                     xorb %dl, %bl
  1089: 43881c1a                 movb %bl, (%r10, %r11)
  1093: 33db                     xorl %ebx, %ebx
  1095: 49ffc3                   incq %r11
  1098: 4c3bd9                   cmpq %rcx, %r11
  1101: 0f8281feffff             jb 0x2d4
  1107: 4883c408                 addq $8, %rsp
  1111: 415e                     popq %r14
  1113: 415d                     popq %r13
  1115: 415c                     popq %r12
  1117: 415f                     popq %r15
  1119: 5b                       popq %rbx
  1120: c5f877                   vzeroupper 
  1123: c3                       retq 
  1124: 0f0b                     ud2 
  1126: 0f0b                     ud2 
  1128: 0f0b                     ud2 
  1130: 0f0b                     ud2 
  1132: 0f0b                     ud2 
  1134: 0f0b                     ud2 
  1136: 0f0b                     ud2 
  1138: 0f0b                     ud2 
  1140: 0f0b                     ud2 
  1142: 0f0b                     ud2 
  1144: 0f0b                     ud2 
  1146: 0f0b                     ud2 
  1148: 0f0b                     ud2 
  1150: 0f0b                     ud2 
