; config_t1_469 target=avx2 twin=nt topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1408 twins=["nt"] uops_per_byte=3.03125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d67050000           leaq 0x567(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: c4417d6f07               vmovdqa (%r15), %ymm8
    52: 488bc1                   movq %rcx, %rax
    55: 4883e820                 subq $0x20, %rax
    59: 0f8267020000             jb 0x2a8
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   102: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   107: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   112: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   117: c4c17d6f7f40             vmovdqa 0x40(%r15), %ymm7
   123: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   128: c4c17d6fbf80000000       vmovdqa 0x80(%r15), %ymm7
   137: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   142: c5d5efc6                 vpxor %ymm6, %ymm5, %ymm0
   146: c4c17d6fbfc0000000       vmovdqa 0xc0(%r15), %ymm7
   155: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   160: c4c17d6fbf00010000       vmovdqa 0x100(%r15), %ymm7
   169: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   174: c5d5efce                 vpxor %ymm6, %ymm5, %ymm1
   178: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   184: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   189: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   194: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   199: c4c17d6fbf40010000       vmovdqa 0x140(%r15), %ymm7
   208: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   213: c4c17d6fbf80010000       vmovdqa 0x180(%r15), %ymm7
   222: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   227: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   231: c5fdefc6                 vpxor %ymm6, %ymm0, %ymm0
   235: c4c17d6fbfc0010000       vmovdqa 0x1c0(%r15), %ymm7
   244: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   249: c4c17d6fbf00020000       vmovdqa 0x200(%r15), %ymm7
   258: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   263: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   267: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   271: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   277: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   282: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   287: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   292: c4c17d6fbf40020000       vmovdqa 0x240(%r15), %ymm7
   301: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   306: c4c17d6fbf80020000       vmovdqa 0x280(%r15), %ymm7
   315: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   320: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   324: c5fdefc6                 vpxor %ymm6, %ymm0, %ymm0
   328: c4c17d6fbfc0020000       vmovdqa 0x2c0(%r15), %ymm7
   337: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   342: c4c17d6fbf00030000       vmovdqa 0x300(%r15), %ymm7
   351: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   356: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   360: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   364: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   371: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   376: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   381: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   386: c4c17d6fbf40030000       vmovdqa 0x340(%r15), %ymm7
   395: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   400: c4c17d6fbf80030000       vmovdqa 0x380(%r15), %ymm7
   409: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   414: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   418: c5fdefc6                 vpxor %ymm6, %ymm0, %ymm0
   422: c4c17d6fbfc0030000       vmovdqa 0x3c0(%r15), %ymm7
   431: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   436: c4c17d6fbf00040000       vmovdqa 0x400(%r15), %ymm7
   445: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   450: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   454: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   458: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   464: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   469: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   474: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   479: c4c17d6fbf40040000       vmovdqa 0x440(%r15), %ymm7
   488: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   493: c4c17d6fbf80040000       vmovdqa 0x480(%r15), %ymm7
   502: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   507: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   511: c5fdefc6                 vpxor %ymm6, %ymm0, %ymm0
   515: c4c17d6fbfc0040000       vmovdqa 0x4c0(%r15), %ymm7
   524: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   529: c4c17d6fbf00050000       vmovdqa 0x500(%r15), %ymm7
   538: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   543: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   547: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   551: 4d8b5028                 movq 0x28(%r8), %r10
   555: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   561: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   566: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   571: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   576: c4c17d6fbf40050000       vmovdqa 0x540(%r15), %ymm7
   585: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   590: c4c17d6fbf80050000       vmovdqa 0x580(%r15), %ymm7
   599: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   604: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   608: c5fdefc6                 vpxor %ymm6, %ymm0, %ymm0
   612: c4c17d6fbf40050000       vmovdqa 0x540(%r15), %ymm7
   621: c4e24500eb               vpshufb %ymm3, %ymm7, %ymm5
   626: c4c17d6fbf80050000       vmovdqa 0x580(%r15), %ymm7
   635: c4e24500f4               vpshufb %ymm4, %ymm7, %ymm6
   640: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   644: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   648: 4d8b11                   movq (%r9), %r10
   651: c4817de7041a             vmovntdq %ymm0, (%r10, %r11)
   657: 4d8b5108                 movq 8(%r9), %r10
   661: c4817de70c1a             vmovntdq %ymm1, (%r10, %r11)
   667: 4983c320                 addq $0x20, %r11
   671: 4c3bd8                   cmpq %rax, %r11
   674: 0f86b8fdffff             jbe 0x60
   680: 4c3bd9                   cmpq %rcx, %r11
   683: 0f83a1020000             jae 0x552
   689: 4153                     pushq %r11
   691: 498dbfc0050000           leaq 0x5c0(%r15), %rdi
   698: 498db7c0060000           leaq 0x6c0(%r15), %rsi
   705: 4d8b11                   movq (%r9), %r10
   708: 4c8b1c24                 movq (%rsp), %r11
   712: 33db                     xorl %ebx, %ebx
   714: 498b00                   movq (%r8), %rax
   717: 420fb61418               movzbl (%rax, %r11), %edx
   722: 4885d2                   testq %rdx, %rdx
   725: 0f8422000000             je 0x2fd
   731: 0fb60417                 movzbl (%rdi, %rdx), %eax
   735: 81c0c1000000             addl $0xc1, %eax
   741: 81f8ff000000             cmpl $0xff, %eax
   747: 0f8206000000             jb 0x2f7
   753: 81c001ffffff             addl $0xffffff01, %eax
   759: 0fb60406                 movzbl (%rsi, %rax), %eax
   763: 32d8                     xorb %al, %bl
   765: 498b4008                 movq 8(%r8), %rax
   769: 420fb61418               movzbl (%rax, %r11), %edx
   774: 4885d2                   testq %rdx, %rdx
   777: 0f841f000000             je 0x32e
   783: 0fb60417                 movzbl (%rdi, %rdx), %eax
   787: 83c003                   addl $3, %eax
   790: 81f8ff000000             cmpl $0xff, %eax
   796: 0f8206000000             jb 0x328
   802: 81c001ffffff             addl $0xffffff01, %eax
   808: 0fb60406                 movzbl (%rsi, %rax), %eax
   812: 32d8                     xorb %al, %bl
   814: 498b4010                 movq 0x10(%r8), %rax
   818: 420fb61418               movzbl (%rax, %r11), %edx
   823: 4885d2                   testq %rdx, %rdx
   826: 0f8422000000             je 0x362
   832: 0fb60417                 movzbl (%rdi, %rdx), %eax
   836: 81c0c6000000             addl $0xc6, %eax
   842: 81f8ff000000             cmpl $0xff, %eax
   848: 0f8206000000             jb 0x35c
   854: 81c001ffffff             addl $0xffffff01, %eax
   860: 0fb60406                 movzbl (%rsi, %rax), %eax
   864: 32d8                     xorb %al, %bl
   866: 498b4018                 movq 0x18(%r8), %rax
   870: 420fb61418               movzbl (%rax, %r11), %edx
   875: 4885d2                   testq %rdx, %rdx
   878: 0f841f000000             je 0x393
   884: 0fb60417                 movzbl (%rdi, %rdx), %eax
   888: 83c001                   addl $1, %eax
   891: 81f8ff000000             cmpl $0xff, %eax
   897: 0f8206000000             jb 0x38d
   903: 81c001ffffff             addl $0xffffff01, %eax
   909: 0fb60406                 movzbl (%rsi, %rax), %eax
   913: 32d8                     xorb %al, %bl
   915: 498b4020                 movq 0x20(%r8), %rax
   919: 420fb61418               movzbl (%rax, %r11), %edx
   924: 4885d2                   testq %rdx, %rdx
   927: 0f8422000000             je 0x3c7
   933: 0fb60417                 movzbl (%rdi, %rdx), %eax
   937: 81c0fe000000             addl $0xfe, %eax
   943: 81f8ff000000             cmpl $0xff, %eax
   949: 0f8206000000             jb 0x3c1
   955: 81c001ffffff             addl $0xffffff01, %eax
   961: 0fb60406                 movzbl (%rsi, %rax), %eax
   965: 32d8                     xorb %al, %bl
   967: 498b4028                 movq 0x28(%r8), %rax
   971: 420fb61418               movzbl (%rax, %r11), %edx
   976: 4885d2                   testq %rdx, %rdx
   979: 0f841f000000             je 0x3f8
   985: 0fb60417                 movzbl (%rdi, %rdx), %eax
   989: 83c031                   addl $0x31, %eax
   992: 81f8ff000000             cmpl $0xff, %eax
   998: 0f8206000000             jb 0x3f2
  1004: 81c001ffffff             addl $0xffffff01, %eax
  1010: 0fb60406                 movzbl (%rsi, %rax), %eax
  1014: 32d8                     xorb %al, %bl
  1016: 43881c1a                 movb %bl, (%r10, %r11)
  1020: 33db                     xorl %ebx, %ebx
  1022: 49ffc3                   incq %r11
  1025: 4c3bd9                   cmpq %rcx, %r11
  1028: 0f82c0feffff             jb 0x2ca
  1034: 4d8b5108                 movq 8(%r9), %r10
  1038: 4c8b1c24                 movq (%rsp), %r11
  1042: 33db                     xorl %ebx, %ebx
  1044: 498b00                   movq (%r8), %rax
  1047: 420fb61418               movzbl (%rax, %r11), %edx
  1052: 4885d2                   testq %rdx, %rdx
  1055: 0f841f000000             je 0x444
  1061: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1065: 83c01c                   addl $0x1c, %eax
  1068: 81f8ff000000             cmpl $0xff, %eax
  1074: 0f8206000000             jb 0x43e
  1080: 81c001ffffff             addl $0xffffff01, %eax
  1086: 0fb60406                 movzbl (%rsi, %rax), %eax
  1090: 32d8                     xorb %al, %bl
  1092: 498b4008                 movq 8(%r8), %rax
  1096: 420fb61418               movzbl (%rax, %r11), %edx
  1101: 4885d2                   testq %rdx, %rdx
  1104: 0f8422000000             je 0x478
  1110: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1114: 81c0df000000             addl $0xdf, %eax
  1120: 81f8ff000000             cmpl $0xff, %eax
  1126: 0f8206000000             jb 0x472
  1132: 81c001ffffff             addl $0xffffff01, %eax
  1138: 0fb60406                 movzbl (%rsi, %rax), %eax
  1142: 32d8                     xorb %al, %bl
  1144: 498b4010                 movq 0x10(%r8), %rax
  1148: 420fb61418               movzbl (%rax, %r11), %edx
  1153: 4885d2                   testq %rdx, %rdx
  1156: 0f841f000000             je 0x4a9
  1162: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1166: 83c01a                   addl $0x1a, %eax
  1169: 81f8ff000000             cmpl $0xff, %eax
  1175: 0f8206000000             jb 0x4a3
  1181: 81c001ffffff             addl $0xffffff01, %eax
  1187: 0fb60406                 movzbl (%rsi, %rax), %eax
  1191: 32d8                     xorb %al, %bl
  1193: 498b4018                 movq 0x18(%r8), %rax
  1197: 420fb61418               movzbl (%rax, %r11), %edx
  1202: 4885d2                   testq %rdx, %rdx
  1205: 0f841f000000             je 0x4da
  1211: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1215: 83c019                   addl $0x19, %eax
  1218: 81f8ff000000             cmpl $0xff, %eax
  1224: 0f8206000000             jb 0x4d4
  1230: 81c001ffffff             addl $0xffffff01, %eax
  1236: 0fb60406                 movzbl (%rsi, %rax), %eax
  1240: 32d8                     xorb %al, %bl
  1242: 498b4020                 movq 0x20(%r8), %rax
  1246: 420fb61418               movzbl (%rax, %r11), %edx
  1251: 4885d2                   testq %rdx, %rdx
  1254: 0f841f000000             je 0x50b
  1260: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1264: 83c018                   addl $0x18, %eax
  1267: 81f8ff000000             cmpl $0xff, %eax
  1273: 0f8206000000             jb 0x505
  1279: 81c001ffffff             addl $0xffffff01, %eax
  1285: 0fb60406                 movzbl (%rsi, %rax), %eax
  1289: 32d8                     xorb %al, %bl
  1291: 498b4028                 movq 0x28(%r8), %rax
  1295: 420fb61418               movzbl (%rax, %r11), %edx
  1300: 4885d2                   testq %rdx, %rdx
  1303: 0f841f000000             je 0x53c
  1309: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1313: 83c031                   addl $0x31, %eax
  1316: 81f8ff000000             cmpl $0xff, %eax
  1322: 0f8206000000             jb 0x536
  1328: 81c001ffffff             addl $0xffffff01, %eax
  1334: 0fb60406                 movzbl (%rsi, %rax), %eax
  1338: 32d8                     xorb %al, %bl
  1340: 43881c1a                 movb %bl, (%r10, %r11)
  1344: 33db                     xorl %ebx, %ebx
  1346: 49ffc3                   incq %r11
  1349: 4c3bd9                   cmpq %rcx, %r11
  1352: 0f82c6feffff             jb 0x414
  1358: 4883c408                 addq $8, %rsp
  1362: 415e                     popq %r14
  1364: 415d                     popq %r13
  1366: 415c                     popq %r12
  1368: 415f                     popq %r15
  1370: 5b                       popq %rbx
  1371: c5f877                   vzeroupper 
  1374: c3                       retq 
  1375: 0f0b                     ud2 
  1377: 0f0b                     ud2 
  1379: 0f0b                     ud2 
  1381: 0f0b                     ud2 
  1383: 0f0b                     ud2 
  1385: 0f0b                     ud2 
  1387: 0f0b                     ud2 
  1389: 0f0b                     ud2 
  1391: 0f0b                     ud2 
  1393: 0f0b                     ud2 
  1395: 0f0b                     ud2 
  1397: 0f0b                     ud2 
  1399: 0f0b                     ud2 
  1401: 0f0b                     ud2 
  1403: 0f0b                     ud2 
  1405: 0f0b                     ud2 
  1407: cc                       int3 
