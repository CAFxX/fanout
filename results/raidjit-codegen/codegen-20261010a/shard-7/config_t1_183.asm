; config_t1_183 target=avx512-gfni twin=nt topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1088 twins=["nt"] uops_per_byte=0.6875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d27040000           leaq 0x427(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4883e840                 subq $0x40, %rax
    54: 0f823d010000             jb 0x179
    60: 0f1f4000                 nopl (%rax)
    64: 62b17d486f141e           vmovdqa32 (%rsi, %r11), %zmm2
    71: 62d17d486f27             vmovdqa32 (%r15), %zmm4
    77: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
    84: 62f17d486fc3             vmovdqa32 %zmm3, %zmm0
    90: 62d17d486f6701           vmovdqa32 0x40(%r15), %zmm4
    97: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   104: 62f17d486fcb             vmovdqa32 %zmm3, %zmm1
   110: 62b17d486f141f           vmovdqa32 (%rdi, %r11), %zmm2
   117: 62d17d486f6702           vmovdqa32 0x80(%r15), %zmm4
   124: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   131: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   137: 62d17d486f6703           vmovdqa32 0xc0(%r15), %zmm4
   144: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   151: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   157: 62917d486f141c           vmovdqa32 (%r12, %r11), %zmm2
   164: 62d17d486f6704           vmovdqa32 0x100(%r15), %zmm4
   171: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   178: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   184: 62d17d486f6705           vmovdqa32 0x140(%r15), %zmm4
   191: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   198: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   204: 62917d486f541d00         vmovdqa32 (%r13, %r11), %zmm2
   212: 62d17d486f6706           vmovdqa32 0x180(%r15), %zmm4
   219: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   226: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   232: 62d17d486f6707           vmovdqa32 0x1c0(%r15), %zmm4
   239: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   246: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   252: 62917d486f141e           vmovdqa32 (%r14, %r11), %zmm2
   259: 62d17d486f6708           vmovdqa32 0x200(%r15), %zmm4
   266: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   273: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   279: 62d17d486f6709           vmovdqa32 0x240(%r15), %zmm4
   286: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   293: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   299: 4d8b5028                 movq 0x28(%r8), %r10
   303: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   310: 62d17d486f670a           vmovdqa32 0x280(%r15), %zmm4
   317: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   324: 62f17d48efc3             vpxord %zmm3, %zmm0, %zmm0
   330: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   337: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   343: 4d8b11                   movq (%r9), %r10
   346: 62917d48e7041a           vmovntdq %zmm0, (%r10, %r11)
   353: 4d8b5108                 movq 8(%r9), %r10
   357: 62917d48e70c1a           vmovntdq %zmm1, (%r10, %r11)
   364: 4983c340                 addq $0x40, %r11
   368: 4c3bd8                   cmpq %rax, %r11
   371: 0f86c7feffff             jbe 0x40
   377: 4c3bd9                   cmpq %rcx, %r11
   380: 0f83a1020000             jae 0x423
   386: 4153                     pushq %r11
   388: 498dbfc0020000           leaq 0x2c0(%r15), %rdi
   395: 498db7c0030000           leaq 0x3c0(%r15), %rsi
   402: 4d8b11                   movq (%r9), %r10
   405: 4c8b1c24                 movq (%rsp), %r11
   409: 33db                     xorl %ebx, %ebx
   411: 498b00                   movq (%r8), %rax
   414: 420fb61418               movzbl (%rax, %r11), %edx
   419: 4885d2                   testq %rdx, %rdx
   422: 0f8422000000             je 0x1ce
   428: 0fb60417                 movzbl (%rdi, %rdx), %eax
   432: 81c0c1000000             addl $0xc1, %eax
   438: 81f8ff000000             cmpl $0xff, %eax
   444: 0f8206000000             jb 0x1c8
   450: 81c001ffffff             addl $0xffffff01, %eax
   456: 0fb60406                 movzbl (%rsi, %rax), %eax
   460: 32d8                     xorb %al, %bl
   462: 498b4008                 movq 8(%r8), %rax
   466: 420fb61418               movzbl (%rax, %r11), %edx
   471: 4885d2                   testq %rdx, %rdx
   474: 0f841f000000             je 0x1ff
   480: 0fb60417                 movzbl (%rdi, %rdx), %eax
   484: 83c003                   addl $3, %eax
   487: 81f8ff000000             cmpl $0xff, %eax
   493: 0f8206000000             jb 0x1f9
   499: 81c001ffffff             addl $0xffffff01, %eax
   505: 0fb60406                 movzbl (%rsi, %rax), %eax
   509: 32d8                     xorb %al, %bl
   511: 498b4010                 movq 0x10(%r8), %rax
   515: 420fb61418               movzbl (%rax, %r11), %edx
   520: 4885d2                   testq %rdx, %rdx
   523: 0f8422000000             je 0x233
   529: 0fb60417                 movzbl (%rdi, %rdx), %eax
   533: 81c0c6000000             addl $0xc6, %eax
   539: 81f8ff000000             cmpl $0xff, %eax
   545: 0f8206000000             jb 0x22d
   551: 81c001ffffff             addl $0xffffff01, %eax
   557: 0fb60406                 movzbl (%rsi, %rax), %eax
   561: 32d8                     xorb %al, %bl
   563: 498b4018                 movq 0x18(%r8), %rax
   567: 420fb61418               movzbl (%rax, %r11), %edx
   572: 4885d2                   testq %rdx, %rdx
   575: 0f841f000000             je 0x264
   581: 0fb60417                 movzbl (%rdi, %rdx), %eax
   585: 83c001                   addl $1, %eax
   588: 81f8ff000000             cmpl $0xff, %eax
   594: 0f8206000000             jb 0x25e
   600: 81c001ffffff             addl $0xffffff01, %eax
   606: 0fb60406                 movzbl (%rsi, %rax), %eax
   610: 32d8                     xorb %al, %bl
   612: 498b4020                 movq 0x20(%r8), %rax
   616: 420fb61418               movzbl (%rax, %r11), %edx
   621: 4885d2                   testq %rdx, %rdx
   624: 0f8422000000             je 0x298
   630: 0fb60417                 movzbl (%rdi, %rdx), %eax
   634: 81c0fe000000             addl $0xfe, %eax
   640: 81f8ff000000             cmpl $0xff, %eax
   646: 0f8206000000             jb 0x292
   652: 81c001ffffff             addl $0xffffff01, %eax
   658: 0fb60406                 movzbl (%rsi, %rax), %eax
   662: 32d8                     xorb %al, %bl
   664: 498b4028                 movq 0x28(%r8), %rax
   668: 420fb61418               movzbl (%rax, %r11), %edx
   673: 4885d2                   testq %rdx, %rdx
   676: 0f841f000000             je 0x2c9
   682: 0fb60417                 movzbl (%rdi, %rdx), %eax
   686: 83c031                   addl $0x31, %eax
   689: 81f8ff000000             cmpl $0xff, %eax
   695: 0f8206000000             jb 0x2c3
   701: 81c001ffffff             addl $0xffffff01, %eax
   707: 0fb60406                 movzbl (%rsi, %rax), %eax
   711: 32d8                     xorb %al, %bl
   713: 43881c1a                 movb %bl, (%r10, %r11)
   717: 33db                     xorl %ebx, %ebx
   719: 49ffc3                   incq %r11
   722: 4c3bd9                   cmpq %rcx, %r11
   725: 0f82c0feffff             jb 0x19b
   731: 4d8b5108                 movq 8(%r9), %r10
   735: 4c8b1c24                 movq (%rsp), %r11
   739: 33db                     xorl %ebx, %ebx
   741: 498b00                   movq (%r8), %rax
   744: 420fb61418               movzbl (%rax, %r11), %edx
   749: 4885d2                   testq %rdx, %rdx
   752: 0f841f000000             je 0x315
   758: 0fb60417                 movzbl (%rdi, %rdx), %eax
   762: 83c01c                   addl $0x1c, %eax
   765: 81f8ff000000             cmpl $0xff, %eax
   771: 0f8206000000             jb 0x30f
   777: 81c001ffffff             addl $0xffffff01, %eax
   783: 0fb60406                 movzbl (%rsi, %rax), %eax
   787: 32d8                     xorb %al, %bl
   789: 498b4008                 movq 8(%r8), %rax
   793: 420fb61418               movzbl (%rax, %r11), %edx
   798: 4885d2                   testq %rdx, %rdx
   801: 0f8422000000             je 0x349
   807: 0fb60417                 movzbl (%rdi, %rdx), %eax
   811: 81c0df000000             addl $0xdf, %eax
   817: 81f8ff000000             cmpl $0xff, %eax
   823: 0f8206000000             jb 0x343
   829: 81c001ffffff             addl $0xffffff01, %eax
   835: 0fb60406                 movzbl (%rsi, %rax), %eax
   839: 32d8                     xorb %al, %bl
   841: 498b4010                 movq 0x10(%r8), %rax
   845: 420fb61418               movzbl (%rax, %r11), %edx
   850: 4885d2                   testq %rdx, %rdx
   853: 0f841f000000             je 0x37a
   859: 0fb60417                 movzbl (%rdi, %rdx), %eax
   863: 83c01a                   addl $0x1a, %eax
   866: 81f8ff000000             cmpl $0xff, %eax
   872: 0f8206000000             jb 0x374
   878: 81c001ffffff             addl $0xffffff01, %eax
   884: 0fb60406                 movzbl (%rsi, %rax), %eax
   888: 32d8                     xorb %al, %bl
   890: 498b4018                 movq 0x18(%r8), %rax
   894: 420fb61418               movzbl (%rax, %r11), %edx
   899: 4885d2                   testq %rdx, %rdx
   902: 0f841f000000             je 0x3ab
   908: 0fb60417                 movzbl (%rdi, %rdx), %eax
   912: 83c019                   addl $0x19, %eax
   915: 81f8ff000000             cmpl $0xff, %eax
   921: 0f8206000000             jb 0x3a5
   927: 81c001ffffff             addl $0xffffff01, %eax
   933: 0fb60406                 movzbl (%rsi, %rax), %eax
   937: 32d8                     xorb %al, %bl
   939: 498b4020                 movq 0x20(%r8), %rax
   943: 420fb61418               movzbl (%rax, %r11), %edx
   948: 4885d2                   testq %rdx, %rdx
   951: 0f841f000000             je 0x3dc
   957: 0fb60417                 movzbl (%rdi, %rdx), %eax
   961: 83c018                   addl $0x18, %eax
   964: 81f8ff000000             cmpl $0xff, %eax
   970: 0f8206000000             jb 0x3d6
   976: 81c001ffffff             addl $0xffffff01, %eax
   982: 0fb60406                 movzbl (%rsi, %rax), %eax
   986: 32d8                     xorb %al, %bl
   988: 498b4028                 movq 0x28(%r8), %rax
   992: 420fb61418               movzbl (%rax, %r11), %edx
   997: 4885d2                   testq %rdx, %rdx
  1000: 0f841f000000             je 0x40d
  1006: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1010: 83c031                   addl $0x31, %eax
  1013: 81f8ff000000             cmpl $0xff, %eax
  1019: 0f8206000000             jb 0x407
  1025: 81c001ffffff             addl $0xffffff01, %eax
  1031: 0fb60406                 movzbl (%rsi, %rax), %eax
  1035: 32d8                     xorb %al, %bl
  1037: 43881c1a                 movb %bl, (%r10, %r11)
  1041: 33db                     xorl %ebx, %ebx
  1043: 49ffc3                   incq %r11
  1046: 4c3bd9                   cmpq %rcx, %r11
  1049: 0f82c6feffff             jb 0x2e5
  1055: 4883c408                 addq $8, %rsp
  1059: 415e                     popq %r14
  1061: 415d                     popq %r13
  1063: 415c                     popq %r12
  1065: 415f                     popq %r15
  1067: 5b                       popq %rbx
  1068: c5f877                   vzeroupper 
  1071: c3                       retq 
  1072: 0f0b                     ud2 
  1074: 0f0b                     ud2 
  1076: 0f0b                     ud2 
  1078: 0f0b                     ud2 
  1080: 0f0b                     ud2 
  1082: 0f0b                     ud2 
  1084: 0f0b                     ud2 
  1086: 0f0b                     ud2 
