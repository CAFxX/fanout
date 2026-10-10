; config_t1_194 target=avx512-gfni twin=rec_cf topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1024 twins=[] uops_per_byte=0.53125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de7030000           leaq 0x3e7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4883e840                 subq $0x40, %rax
    54: 0f82f7000000             jb 0x133
    60: 0f1f4000                 nopl (%rax)
    64: 62b17d486f141e           vmovdqa32 (%rsi, %r11), %zmm2
    71: 62f17d486fc2             vmovdqa32 %zmm2, %zmm0
    77: 62d17d486f27             vmovdqa32 (%r15), %zmm4
    83: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
    90: 62f17d486fcb             vmovdqa32 %zmm3, %zmm1
    96: 62b17d486f141f           vmovdqa32 (%rdi, %r11), %zmm2
   103: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   109: 62d17d486f6701           vmovdqa32 0x40(%r15), %zmm4
   116: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   123: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   129: 62917d486f141c           vmovdqa32 (%r12, %r11), %zmm2
   136: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   142: 62d17d486f6702           vmovdqa32 0x80(%r15), %zmm4
   149: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   156: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   162: 62917d486f541d00         vmovdqa32 (%r13, %r11), %zmm2
   170: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   176: 62f17548efca             vpxord %zmm2, %zmm1, %zmm1
   182: 62917d486f141e           vmovdqa32 (%r14, %r11), %zmm2
   189: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   195: 4d8b5028                 movq 0x28(%r8), %r10
   199: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   206: 62f17548efca             vpxord %zmm2, %zmm1, %zmm1
   212: 62d17d486f6703           vmovdqa32 0xc0(%r15), %zmm4
   219: 62f3fd48cedc00           vgf2p8affineqb $0, %zmm4, %zmm0, %zmm3
   226: 62f16548efd9             vpxord %zmm1, %zmm3, %zmm3
   232: 62d17d486f6704           vmovdqa32 0x100(%r15), %zmm4
   239: 62f3fd48ced400           vgf2p8affineqb $0, %zmm4, %zmm0, %zmm2
   246: 62f16d48efd1             vpxord %zmm1, %zmm2, %zmm2
   252: 62d17d486f6705           vmovdqa32 0x140(%r15), %zmm4
   259: 62f3e548cecc00           vgf2p8affineqb $0, %zmm4, %zmm3, %zmm1
   266: 62f3ed48cec400           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm0
   273: 4d8b11                   movq (%r9), %r10
   276: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   283: 4d8b5108                 movq 8(%r9), %r10
   287: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
   294: 4983c340                 addq $0x40, %r11
   298: 4c3bd8                   cmpq %rax, %r11
   301: 0f860dffffff             jbe 0x40
   307: 4c3bd9                   cmpq %rcx, %r11
   310: 0f83a1020000             jae 0x3dd
   316: 4153                     pushq %r11
   318: 498dbf80010000           leaq 0x180(%r15), %rdi
   325: 498db780020000           leaq 0x280(%r15), %rsi
   332: 4d8b11                   movq (%r9), %r10
   335: 4c8b1c24                 movq (%rsp), %r11
   339: 33db                     xorl %ebx, %ebx
   341: 498b00                   movq (%r8), %rax
   344: 420fb61418               movzbl (%rax, %r11), %edx
   349: 4885d2                   testq %rdx, %rdx
   352: 0f8422000000             je 0x188
   358: 0fb60417                 movzbl (%rdi, %rdx), %eax
   362: 81c0c1000000             addl $0xc1, %eax
   368: 81f8ff000000             cmpl $0xff, %eax
   374: 0f8206000000             jb 0x182
   380: 81c001ffffff             addl $0xffffff01, %eax
   386: 0fb60406                 movzbl (%rsi, %rax), %eax
   390: 32d8                     xorb %al, %bl
   392: 498b4008                 movq 8(%r8), %rax
   396: 420fb61418               movzbl (%rax, %r11), %edx
   401: 4885d2                   testq %rdx, %rdx
   404: 0f841f000000             je 0x1b9
   410: 0fb60417                 movzbl (%rdi, %rdx), %eax
   414: 83c003                   addl $3, %eax
   417: 81f8ff000000             cmpl $0xff, %eax
   423: 0f8206000000             jb 0x1b3
   429: 81c001ffffff             addl $0xffffff01, %eax
   435: 0fb60406                 movzbl (%rsi, %rax), %eax
   439: 32d8                     xorb %al, %bl
   441: 498b4010                 movq 0x10(%r8), %rax
   445: 420fb61418               movzbl (%rax, %r11), %edx
   450: 4885d2                   testq %rdx, %rdx
   453: 0f8422000000             je 0x1ed
   459: 0fb60417                 movzbl (%rdi, %rdx), %eax
   463: 81c0c6000000             addl $0xc6, %eax
   469: 81f8ff000000             cmpl $0xff, %eax
   475: 0f8206000000             jb 0x1e7
   481: 81c001ffffff             addl $0xffffff01, %eax
   487: 0fb60406                 movzbl (%rsi, %rax), %eax
   491: 32d8                     xorb %al, %bl
   493: 498b4018                 movq 0x18(%r8), %rax
   497: 420fb61418               movzbl (%rax, %r11), %edx
   502: 4885d2                   testq %rdx, %rdx
   505: 0f841f000000             je 0x21e
   511: 0fb60417                 movzbl (%rdi, %rdx), %eax
   515: 83c001                   addl $1, %eax
   518: 81f8ff000000             cmpl $0xff, %eax
   524: 0f8206000000             jb 0x218
   530: 81c001ffffff             addl $0xffffff01, %eax
   536: 0fb60406                 movzbl (%rsi, %rax), %eax
   540: 32d8                     xorb %al, %bl
   542: 498b4020                 movq 0x20(%r8), %rax
   546: 420fb61418               movzbl (%rax, %r11), %edx
   551: 4885d2                   testq %rdx, %rdx
   554: 0f8422000000             je 0x252
   560: 0fb60417                 movzbl (%rdi, %rdx), %eax
   564: 81c0fe000000             addl $0xfe, %eax
   570: 81f8ff000000             cmpl $0xff, %eax
   576: 0f8206000000             jb 0x24c
   582: 81c001ffffff             addl $0xffffff01, %eax
   588: 0fb60406                 movzbl (%rsi, %rax), %eax
   592: 32d8                     xorb %al, %bl
   594: 498b4028                 movq 0x28(%r8), %rax
   598: 420fb61418               movzbl (%rax, %r11), %edx
   603: 4885d2                   testq %rdx, %rdx
   606: 0f841f000000             je 0x283
   612: 0fb60417                 movzbl (%rdi, %rdx), %eax
   616: 83c031                   addl $0x31, %eax
   619: 81f8ff000000             cmpl $0xff, %eax
   625: 0f8206000000             jb 0x27d
   631: 81c001ffffff             addl $0xffffff01, %eax
   637: 0fb60406                 movzbl (%rsi, %rax), %eax
   641: 32d8                     xorb %al, %bl
   643: 43881c1a                 movb %bl, (%r10, %r11)
   647: 33db                     xorl %ebx, %ebx
   649: 49ffc3                   incq %r11
   652: 4c3bd9                   cmpq %rcx, %r11
   655: 0f82c0feffff             jb 0x155
   661: 4d8b5108                 movq 8(%r9), %r10
   665: 4c8b1c24                 movq (%rsp), %r11
   669: 33db                     xorl %ebx, %ebx
   671: 498b00                   movq (%r8), %rax
   674: 420fb61418               movzbl (%rax, %r11), %edx
   679: 4885d2                   testq %rdx, %rdx
   682: 0f841f000000             je 0x2cf
   688: 0fb60417                 movzbl (%rdi, %rdx), %eax
   692: 83c01c                   addl $0x1c, %eax
   695: 81f8ff000000             cmpl $0xff, %eax
   701: 0f8206000000             jb 0x2c9
   707: 81c001ffffff             addl $0xffffff01, %eax
   713: 0fb60406                 movzbl (%rsi, %rax), %eax
   717: 32d8                     xorb %al, %bl
   719: 498b4008                 movq 8(%r8), %rax
   723: 420fb61418               movzbl (%rax, %r11), %edx
   728: 4885d2                   testq %rdx, %rdx
   731: 0f8422000000             je 0x303
   737: 0fb60417                 movzbl (%rdi, %rdx), %eax
   741: 81c0df000000             addl $0xdf, %eax
   747: 81f8ff000000             cmpl $0xff, %eax
   753: 0f8206000000             jb 0x2fd
   759: 81c001ffffff             addl $0xffffff01, %eax
   765: 0fb60406                 movzbl (%rsi, %rax), %eax
   769: 32d8                     xorb %al, %bl
   771: 498b4010                 movq 0x10(%r8), %rax
   775: 420fb61418               movzbl (%rax, %r11), %edx
   780: 4885d2                   testq %rdx, %rdx
   783: 0f841f000000             je 0x334
   789: 0fb60417                 movzbl (%rdi, %rdx), %eax
   793: 83c01a                   addl $0x1a, %eax
   796: 81f8ff000000             cmpl $0xff, %eax
   802: 0f8206000000             jb 0x32e
   808: 81c001ffffff             addl $0xffffff01, %eax
   814: 0fb60406                 movzbl (%rsi, %rax), %eax
   818: 32d8                     xorb %al, %bl
   820: 498b4018                 movq 0x18(%r8), %rax
   824: 420fb61418               movzbl (%rax, %r11), %edx
   829: 4885d2                   testq %rdx, %rdx
   832: 0f841f000000             je 0x365
   838: 0fb60417                 movzbl (%rdi, %rdx), %eax
   842: 83c019                   addl $0x19, %eax
   845: 81f8ff000000             cmpl $0xff, %eax
   851: 0f8206000000             jb 0x35f
   857: 81c001ffffff             addl $0xffffff01, %eax
   863: 0fb60406                 movzbl (%rsi, %rax), %eax
   867: 32d8                     xorb %al, %bl
   869: 498b4020                 movq 0x20(%r8), %rax
   873: 420fb61418               movzbl (%rax, %r11), %edx
   878: 4885d2                   testq %rdx, %rdx
   881: 0f841f000000             je 0x396
   887: 0fb60417                 movzbl (%rdi, %rdx), %eax
   891: 83c018                   addl $0x18, %eax
   894: 81f8ff000000             cmpl $0xff, %eax
   900: 0f8206000000             jb 0x390
   906: 81c001ffffff             addl $0xffffff01, %eax
   912: 0fb60406                 movzbl (%rsi, %rax), %eax
   916: 32d8                     xorb %al, %bl
   918: 498b4028                 movq 0x28(%r8), %rax
   922: 420fb61418               movzbl (%rax, %r11), %edx
   927: 4885d2                   testq %rdx, %rdx
   930: 0f841f000000             je 0x3c7
   936: 0fb60417                 movzbl (%rdi, %rdx), %eax
   940: 83c031                   addl $0x31, %eax
   943: 81f8ff000000             cmpl $0xff, %eax
   949: 0f8206000000             jb 0x3c1
   955: 81c001ffffff             addl $0xffffff01, %eax
   961: 0fb60406                 movzbl (%rsi, %rax), %eax
   965: 32d8                     xorb %al, %bl
   967: 43881c1a                 movb %bl, (%r10, %r11)
   971: 33db                     xorl %ebx, %ebx
   973: 49ffc3                   incq %r11
   976: 4c3bd9                   cmpq %rcx, %r11
   979: 0f82c6feffff             jb 0x29f
   985: 4883c408                 addq $8, %rsp
   989: 415e                     popq %r14
   991: 415d                     popq %r13
   993: 415c                     popq %r12
   995: 415f                     popq %r15
   997: 5b                       popq %rbx
   998: c5f877                   vzeroupper 
  1001: c3                       retq 
  1002: 0f0b                     ud2 
  1004: 0f0b                     ud2 
  1006: 0f0b                     ud2 
  1008: 0f0b                     ud2 
  1010: 0f0b                     ud2 
  1012: 0f0b                     ud2 
  1014: 0f0b                     ud2 
  1016: 0f0b                     ud2 
  1018: 0f0b                     ud2 
  1020: 0f0b                     ud2 
  1022: 0f0b                     ud2 
