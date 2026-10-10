; config_t1_91 target=avx512-gfni twin=baseline topo=rs(6,3)/enc op=encode U=8
; code_size=1152 twins=[] uops_per_byte=0.71875
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
    47: 488bc1                   movq %rcx, %rax
    50: 4883e840                 subq $0x40, %rax
    54: 0f8248010000             jb 0x184
    60: 0f1f4000                 nopl (%rax)
    64: 62b17d486f1c1e           vmovdqa32 (%rsi, %r11), %zmm3
    71: 62f17d486fc3             vmovdqa32 %zmm3, %zmm0
    77: 62d17d486f2f             vmovdqa32 (%r15), %zmm5
    83: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
    90: 62f17d486fcc             vmovdqa32 %zmm4, %zmm1
    96: 62d17d486f6f01           vmovdqa32 0x40(%r15), %zmm5
   103: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   110: 62f17d486fd4             vmovdqa32 %zmm4, %zmm2
   116: 62b17d486f1c1f           vmovdqa32 (%rdi, %r11), %zmm3
   123: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   129: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   136: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   142: 62d17d486f6f02           vmovdqa32 0x80(%r15), %zmm5
   149: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   156: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   162: 62917d486f1c1c           vmovdqa32 (%r12, %r11), %zmm3
   169: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   175: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   182: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   188: 62d17d486f6f03           vmovdqa32 0xc0(%r15), %zmm5
   195: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   202: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   208: 62917d486f5c1d00         vmovdqa32 (%r13, %r11), %zmm3
   216: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   222: 62d17d486f6f04           vmovdqa32 0x100(%r15), %zmm5
   229: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   236: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   242: 62d17d486f2f             vmovdqa32 (%r15), %zmm5
   248: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   255: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   261: 62917d486f1c1e           vmovdqa32 (%r14, %r11), %zmm3
   268: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   274: 62d17d486f6f03           vmovdqa32 0xc0(%r15), %zmm5
   281: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   288: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   294: 62d17d486f6f04           vmovdqa32 0x100(%r15), %zmm5
   301: 62f3e548cee500           vgf2p8affineqb $0, %zmm5, %zmm3, %zmm4
   308: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   314: 4d8b5028                 movq 0x28(%r8), %r10
   318: 62917d486f1c1a           vmovdqa32 (%r10, %r11), %zmm3
   325: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   331: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   337: 62f16d48efd3             vpxord %zmm3, %zmm2, %zmm2
   343: 4d8b11                   movq (%r9), %r10
   346: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   353: 4d8b5108                 movq 8(%r9), %r10
   357: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
   364: 4d8b5110                 movq 0x10(%r9), %r10
   368: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
   375: 4983c340                 addq $0x40, %r11
   379: 4c3bd8                   cmpq %rax, %r11
   382: 0f86bcfeffff             jbe 0x40
   388: 4c3bd9                   cmpq %rcx, %r11
   391: 0f83b2020000             jae 0x43f
   397: 4153                     pushq %r11
   399: 498dbf40010000           leaq 0x140(%r15), %rdi
   406: 498db740020000           leaq 0x240(%r15), %rsi
   413: 4d8b11                   movq (%r9), %r10
   416: 4c8b1c24                 movq (%rsp), %r11
   420: 33db                     xorl %ebx, %ebx
   422: 498b00                   movq (%r8), %rax
   425: 420fb61418               movzbl (%rax, %r11), %edx
   430: 32da                     xorb %dl, %bl
   432: 498b4008                 movq 8(%r8), %rax
   436: 420fb61418               movzbl (%rax, %r11), %edx
   441: 32da                     xorb %dl, %bl
   443: 498b4010                 movq 0x10(%r8), %rax
   447: 420fb61418               movzbl (%rax, %r11), %edx
   452: 32da                     xorb %dl, %bl
   454: 498b4018                 movq 0x18(%r8), %rax
   458: 420fb61418               movzbl (%rax, %r11), %edx
   463: 32da                     xorb %dl, %bl
   465: 498b4020                 movq 0x20(%r8), %rax
   469: 420fb61418               movzbl (%rax, %r11), %edx
   474: 32da                     xorb %dl, %bl
   476: 498b4028                 movq 0x28(%r8), %rax
   480: 420fb61418               movzbl (%rax, %r11), %edx
   485: 32da                     xorb %dl, %bl
   487: 43881c1a                 movb %bl, (%r10, %r11)
   491: 33db                     xorl %ebx, %ebx
   493: 49ffc3                   incq %r11
   496: 4c3bd9                   cmpq %rcx, %r11
   499: 0f82adffffff             jb 0x1a6
   505: 4d8b5108                 movq 8(%r9), %r10
   509: 4c8b1c24                 movq (%rsp), %r11
   513: 33db                     xorl %ebx, %ebx
   515: 498b00                   movq (%r8), %rax
   518: 420fb61418               movzbl (%rax, %r11), %edx
   523: 4885d2                   testq %rdx, %rdx
   526: 0f8422000000             je 0x236
   532: 0fb60417                 movzbl (%rdi, %rdx), %eax
   536: 81c0e6000000             addl $0xe6, %eax
   542: 81f8ff000000             cmpl $0xff, %eax
   548: 0f8206000000             jb 0x230
   554: 81c001ffffff             addl $0xffffff01, %eax
   560: 0fb60406                 movzbl (%rsi, %rax), %eax
   564: 32d8                     xorb %al, %bl
   566: 498b4008                 movq 8(%r8), %rax
   570: 420fb61418               movzbl (%rax, %r11), %edx
   575: 4885d2                   testq %rdx, %rdx
   578: 0f8422000000             je 0x26a
   584: 0fb60417                 movzbl (%rdi, %rdx), %eax
   588: 81c0cd000000             addl $0xcd, %eax
   594: 81f8ff000000             cmpl $0xff, %eax
   600: 0f8206000000             jb 0x264
   606: 81c001ffffff             addl $0xffffff01, %eax
   612: 0fb60406                 movzbl (%rsi, %rax), %eax
   616: 32d8                     xorb %al, %bl
   618: 498b4010                 movq 0x10(%r8), %rax
   622: 420fb61418               movzbl (%rax, %r11), %edx
   627: 4885d2                   testq %rdx, %rdx
   630: 0f841f000000             je 0x29b
   636: 0fb60417                 movzbl (%rdi, %rdx), %eax
   640: 83c04b                   addl $0x4b, %eax
   643: 81f8ff000000             cmpl $0xff, %eax
   649: 0f8206000000             jb 0x295
   655: 81c001ffffff             addl $0xffffff01, %eax
   661: 0fb60406                 movzbl (%rsi, %rax), %eax
   665: 32d8                     xorb %al, %bl
   667: 498b4018                 movq 0x18(%r8), %rax
   671: 420fb61418               movzbl (%rax, %r11), %edx
   676: 4885d2                   testq %rdx, %rdx
   679: 0f841f000000             je 0x2cc
   685: 0fb60417                 movzbl (%rdi, %rdx), %eax
   689: 83c032                   addl $0x32, %eax
   692: 81f8ff000000             cmpl $0xff, %eax
   698: 0f8206000000             jb 0x2c6
   704: 81c001ffffff             addl $0xffffff01, %eax
   710: 0fb60406                 movzbl (%rsi, %rax), %eax
   714: 32d8                     xorb %al, %bl
   716: 498b4020                 movq 0x20(%r8), %rax
   720: 420fb61418               movzbl (%rax, %r11), %edx
   725: 4885d2                   testq %rdx, %rdx
   728: 0f841f000000             je 0x2fd
   734: 0fb60417                 movzbl (%rdi, %rdx), %eax
   738: 83c019                   addl $0x19, %eax
   741: 81f8ff000000             cmpl $0xff, %eax
   747: 0f8206000000             jb 0x2f7
   753: 81c001ffffff             addl $0xffffff01, %eax
   759: 0fb60406                 movzbl (%rsi, %rax), %eax
   763: 32d8                     xorb %al, %bl
   765: 498b4028                 movq 0x28(%r8), %rax
   769: 420fb61418               movzbl (%rax, %r11), %edx
   774: 32da                     xorb %dl, %bl
   776: 43881c1a                 movb %bl, (%r10, %r11)
   780: 33db                     xorl %ebx, %ebx
   782: 49ffc3                   incq %r11
   785: 4c3bd9                   cmpq %rcx, %r11
   788: 0f82e9feffff             jb 0x203
   794: 4d8b5110                 movq 0x10(%r9), %r10
   798: 4c8b1c24                 movq (%rsp), %r11
   802: 33db                     xorl %ebx, %ebx
   804: 498b00                   movq (%r8), %rax
   807: 420fb61418               movzbl (%rax, %r11), %edx
   812: 4885d2                   testq %rdx, %rdx
   815: 0f8422000000             je 0x357
   821: 0fb60417                 movzbl (%rdi, %rdx), %eax
   825: 81c0cd000000             addl $0xcd, %eax
   831: 81f8ff000000             cmpl $0xff, %eax
   837: 0f8206000000             jb 0x351
   843: 81c001ffffff             addl $0xffffff01, %eax
   849: 0fb60406                 movzbl (%rsi, %rax), %eax
   853: 32d8                     xorb %al, %bl
   855: 498b4008                 movq 8(%r8), %rax
   859: 420fb61418               movzbl (%rax, %r11), %edx
   864: 4885d2                   testq %rdx, %rdx
   867: 0f841f000000             je 0x388
   873: 0fb60417                 movzbl (%rdi, %rdx), %eax
   877: 83c04b                   addl $0x4b, %eax
   880: 81f8ff000000             cmpl $0xff, %eax
   886: 0f8206000000             jb 0x382
   892: 81c001ffffff             addl $0xffffff01, %eax
   898: 0fb60406                 movzbl (%rsi, %rax), %eax
   902: 32d8                     xorb %al, %bl
   904: 498b4010                 movq 0x10(%r8), %rax
   908: 420fb61418               movzbl (%rax, %r11), %edx
   913: 4885d2                   testq %rdx, %rdx
   916: 0f841f000000             je 0x3b9
   922: 0fb60417                 movzbl (%rdi, %rdx), %eax
   926: 83c019                   addl $0x19, %eax
   929: 81f8ff000000             cmpl $0xff, %eax
   935: 0f8206000000             jb 0x3b3
   941: 81c001ffffff             addl $0xffffff01, %eax
   947: 0fb60406                 movzbl (%rsi, %rax), %eax
   951: 32d8                     xorb %al, %bl
   953: 498b4018                 movq 0x18(%r8), %rax
   957: 420fb61418               movzbl (%rax, %r11), %edx
   962: 4885d2                   testq %rdx, %rdx
   965: 0f8422000000             je 0x3ed
   971: 0fb60417                 movzbl (%rdi, %rdx), %eax
   975: 81c0e6000000             addl $0xe6, %eax
   981: 81f8ff000000             cmpl $0xff, %eax
   987: 0f8206000000             jb 0x3e7
   993: 81c001ffffff             addl $0xffffff01, %eax
   999: 0fb60406                 movzbl (%rsi, %rax), %eax
  1003: 32d8                     xorb %al, %bl
  1005: 498b4020                 movq 0x20(%r8), %rax
  1009: 420fb61418               movzbl (%rax, %r11), %edx
  1014: 4885d2                   testq %rdx, %rdx
  1017: 0f841f000000             je 0x41e
  1023: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1027: 83c032                   addl $0x32, %eax
  1030: 81f8ff000000             cmpl $0xff, %eax
  1036: 0f8206000000             jb 0x418
  1042: 81c001ffffff             addl $0xffffff01, %eax
  1048: 0fb60406                 movzbl (%rsi, %rax), %eax
  1052: 32d8                     xorb %al, %bl
  1054: 498b4028                 movq 0x28(%r8), %rax
  1058: 420fb61418               movzbl (%rax, %r11), %edx
  1063: 32da                     xorb %dl, %bl
  1065: 43881c1a                 movb %bl, (%r10, %r11)
  1069: 33db                     xorl %ebx, %ebx
  1071: 49ffc3                   incq %r11
  1074: 4c3bd9                   cmpq %rcx, %r11
  1077: 0f82e9feffff             jb 0x324
  1083: 4883c408                 addq $8, %rsp
  1087: 415e                     popq %r14
  1089: 415d                     popq %r13
  1091: 415c                     popq %r12
  1093: 415f                     popq %r15
  1095: 5b                       popq %rbx
  1096: c5f877                   vzeroupper 
  1099: c3                       retq 
  1100: 0f0b                     ud2 
  1102: 0f0b                     ud2 
  1104: 0f0b                     ud2 
  1106: 0f0b                     ud2 
  1108: 0f0b                     ud2 
  1110: 0f0b                     ud2 
  1112: 0f0b                     ud2 
  1114: 0f0b                     ud2 
  1116: 0f0b                     ud2 
  1118: 0f0b                     ud2 
  1120: 0f0b                     ud2 
  1122: 0f0b                     ud2 
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
