; config_t1_28 target=avx512-gfni twin=ternlog topo=raid6(8)/enc op=encode U=1
; code_size=960 twins=["ternlog"] uops_per_byte=0.671875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da7030000           leaq 0x3a7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4883e840                 subq $0x40, %rax
    54: 0f822c010000             jb 0x168
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
   176: 62d17d486f6703           vmovdqa32 0xc0(%r15), %zmm4
   183: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   190: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   196: 62917d486f141e           vmovdqa32 (%r14, %r11), %zmm2
   203: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   209: 62d17d486f6704           vmovdqa32 0x100(%r15), %zmm4
   216: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   223: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   229: 4d8b5028                 movq 0x28(%r8), %r10
   233: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   240: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   246: 62d17d486f6705           vmovdqa32 0x140(%r15), %zmm4
   253: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   260: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   266: 4d8b5030                 movq 0x30(%r8), %r10
   270: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   277: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   283: 62d17d486f6706           vmovdqa32 0x180(%r15), %zmm4
   290: 62f3ed48cedc00           vgf2p8affineqb $0, %zmm4, %zmm2, %zmm3
   297: 62f17548efcb             vpxord %zmm3, %zmm1, %zmm1
   303: 4d8b5038                 movq 0x38(%r8), %r10
   307: 62917d486f141a           vmovdqa32 (%r10, %r11), %zmm2
   314: 62f17d48efc2             vpxord %zmm2, %zmm0, %zmm0
   320: 62f17548efca             vpxord %zmm2, %zmm1, %zmm1
   326: 4d8b11                   movq (%r9), %r10
   329: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   336: 4d8b5108                 movq 8(%r9), %r10
   340: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
   347: 4983c340                 addq $0x40, %r11
   351: 4c3bd8                   cmpq %rax, %r11
   354: 0f86d8feffff             jbe 0x40
   360: 4c3bd9                   cmpq %rcx, %r11
   363: 0f830f020000             jae 0x380
   369: 4153                     pushq %r11
   371: 498dbfc0010000           leaq 0x1c0(%r15), %rdi
   378: 498db7c0020000           leaq 0x2c0(%r15), %rsi
   385: 4d8b11                   movq (%r9), %r10
   388: 4c8b1c24                 movq (%rsp), %r11
   392: 33db                     xorl %ebx, %ebx
   394: 498b00                   movq (%r8), %rax
   397: 420fb61418               movzbl (%rax, %r11), %edx
   402: 32da                     xorb %dl, %bl
   404: 498b4008                 movq 8(%r8), %rax
   408: 420fb61418               movzbl (%rax, %r11), %edx
   413: 32da                     xorb %dl, %bl
   415: 498b4010                 movq 0x10(%r8), %rax
   419: 420fb61418               movzbl (%rax, %r11), %edx
   424: 32da                     xorb %dl, %bl
   426: 498b4018                 movq 0x18(%r8), %rax
   430: 420fb61418               movzbl (%rax, %r11), %edx
   435: 32da                     xorb %dl, %bl
   437: 498b4020                 movq 0x20(%r8), %rax
   441: 420fb61418               movzbl (%rax, %r11), %edx
   446: 32da                     xorb %dl, %bl
   448: 498b4028                 movq 0x28(%r8), %rax
   452: 420fb61418               movzbl (%rax, %r11), %edx
   457: 32da                     xorb %dl, %bl
   459: 498b4030                 movq 0x30(%r8), %rax
   463: 420fb61418               movzbl (%rax, %r11), %edx
   468: 32da                     xorb %dl, %bl
   470: 498b4038                 movq 0x38(%r8), %rax
   474: 420fb61418               movzbl (%rax, %r11), %edx
   479: 32da                     xorb %dl, %bl
   481: 43881c1a                 movb %bl, (%r10, %r11)
   485: 33db                     xorl %ebx, %ebx
   487: 49ffc3                   incq %r11
   490: 4c3bd9                   cmpq %rcx, %r11
   493: 0f8297ffffff             jb 0x18a
   499: 4d8b5108                 movq 8(%r9), %r10
   503: 4c8b1c24                 movq (%rsp), %r11
   507: 33db                     xorl %ebx, %ebx
   509: 498b00                   movq (%r8), %rax
   512: 420fb61418               movzbl (%rax, %r11), %edx
   517: 4885d2                   testq %rdx, %rdx
   520: 0f8422000000             je 0x230
   526: 0fb60417                 movzbl (%rdi, %rdx), %eax
   530: 81c0e6000000             addl $0xe6, %eax
   536: 81f8ff000000             cmpl $0xff, %eax
   542: 0f8206000000             jb 0x22a
   548: 81c001ffffff             addl $0xffffff01, %eax
   554: 0fb60406                 movzbl (%rsi, %rax), %eax
   558: 32d8                     xorb %al, %bl
   560: 498b4008                 movq 8(%r8), %rax
   564: 420fb61418               movzbl (%rax, %r11), %edx
   569: 4885d2                   testq %rdx, %rdx
   572: 0f8422000000             je 0x264
   578: 0fb60417                 movzbl (%rdi, %rdx), %eax
   582: 81c0cd000000             addl $0xcd, %eax
   588: 81f8ff000000             cmpl $0xff, %eax
   594: 0f8206000000             jb 0x25e
   600: 81c001ffffff             addl $0xffffff01, %eax
   606: 0fb60406                 movzbl (%rsi, %rax), %eax
   610: 32d8                     xorb %al, %bl
   612: 498b4010                 movq 0x10(%r8), %rax
   616: 420fb61418               movzbl (%rax, %r11), %edx
   621: 4885d2                   testq %rdx, %rdx
   624: 0f8422000000             je 0x298
   630: 0fb60417                 movzbl (%rdi, %rdx), %eax
   634: 81c0e7000000             addl $0xe7, %eax
   640: 81f8ff000000             cmpl $0xff, %eax
   646: 0f8206000000             jb 0x292
   652: 81c001ffffff             addl $0xffffff01, %eax
   658: 0fb60406                 movzbl (%rsi, %rax), %eax
   662: 32d8                     xorb %al, %bl
   664: 498b4018                 movq 0x18(%r8), %rax
   668: 420fb61418               movzbl (%rax, %r11), %edx
   673: 4885d2                   testq %rdx, %rdx
   676: 0f8422000000             je 0x2cc
   682: 0fb60417                 movzbl (%rdi, %rdx), %eax
   686: 81c0ce000000             addl $0xce, %eax
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
   738: 83c04b                   addl $0x4b, %eax
   741: 81f8ff000000             cmpl $0xff, %eax
   747: 0f8206000000             jb 0x2f7
   753: 81c001ffffff             addl $0xffffff01, %eax
   759: 0fb60406                 movzbl (%rsi, %rax), %eax
   763: 32d8                     xorb %al, %bl
   765: 498b4028                 movq 0x28(%r8), %rax
   769: 420fb61418               movzbl (%rax, %r11), %edx
   774: 4885d2                   testq %rdx, %rdx
   777: 0f841f000000             je 0x32e
   783: 0fb60417                 movzbl (%rdi, %rdx), %eax
   787: 83c032                   addl $0x32, %eax
   790: 81f8ff000000             cmpl $0xff, %eax
   796: 0f8206000000             jb 0x328
   802: 81c001ffffff             addl $0xffffff01, %eax
   808: 0fb60406                 movzbl (%rsi, %rax), %eax
   812: 32d8                     xorb %al, %bl
   814: 498b4030                 movq 0x30(%r8), %rax
   818: 420fb61418               movzbl (%rax, %r11), %edx
   823: 4885d2                   testq %rdx, %rdx
   826: 0f841f000000             je 0x35f
   832: 0fb60417                 movzbl (%rdi, %rdx), %eax
   836: 83c019                   addl $0x19, %eax
   839: 81f8ff000000             cmpl $0xff, %eax
   845: 0f8206000000             jb 0x359
   851: 81c001ffffff             addl $0xffffff01, %eax
   857: 0fb60406                 movzbl (%rsi, %rax), %eax
   861: 32d8                     xorb %al, %bl
   863: 498b4038                 movq 0x38(%r8), %rax
   867: 420fb61418               movzbl (%rax, %r11), %edx
   872: 32da                     xorb %dl, %bl
   874: 43881c1a                 movb %bl, (%r10, %r11)
   878: 33db                     xorl %ebx, %ebx
   880: 49ffc3                   incq %r11
   883: 4c3bd9                   cmpq %rcx, %r11
   886: 0f8281feffff             jb 0x1fd
   892: 4883c408                 addq $8, %rsp
   896: 415e                     popq %r14
   898: 415d                     popq %r13
   900: 415c                     popq %r12
   902: 415f                     popq %r15
   904: 5b                       popq %rbx
   905: c5f877                   vzeroupper 
   908: c3                       retq 
   909: 0f0b                     ud2 
   911: 0f0b                     ud2 
   913: 0f0b                     ud2 
   915: 0f0b                     ud2 
   917: 0f0b                     ud2 
   919: 0f0b                     ud2 
   921: 0f0b                     ud2 
   923: 0f0b                     ud2 
   925: 0f0b                     ud2 
   927: 0f0b                     ud2 
   929: 0f0b                     ud2 
   931: 0f0b                     ud2 
   933: 0f0b                     ud2 
   935: 0f0b                     ud2 
   937: 0f0b                     ud2 
   939: 0f0b                     ud2 
   941: 0f0b                     ud2 
   943: 0f0b                     ud2 
   945: 0f0b                     ud2 
   947: 0f0b                     ud2 
   949: 0f0b                     ud2 
   951: 0f0b                     ud2 
   953: 0f0b                     ud2 
   955: 0f0b                     ud2 
   957: 0f0b                     ud2 
   959: cc                       int3 
