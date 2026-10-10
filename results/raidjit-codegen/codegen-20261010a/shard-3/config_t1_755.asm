; config_t1_755 target=sse twin=nt topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1536 twins=["nt"] uops_per_byte=8.4375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de7050000           leaq 0x5e7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66450f6f07               movdqa (%r15), %xmm8
    52: 488bc1                   movq %rcx, %rax
    55: 4883e810                 subq $0x10, %rax
    59: 0f82ff020000             jb 0x340
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   102: 660f6fda                 movdqa %xmm2, %xmm3
   106: 66410fdbd8               pand %xmm8, %xmm3
   111: 660f6fe2                 movdqa %xmm2, %xmm4
   115: 660f71d404               psrlw $4, %xmm4
   120: 66410fdbe0               pand %xmm8, %xmm4
   125: 66410f6f7f40             movdqa 0x40(%r15), %xmm7
   131: 660f6fef                 movdqa %xmm7, %xmm5
   135: 660f3800eb               pshufb %xmm3, %xmm5
   140: 66410f6fbf80000000       movdqa 0x80(%r15), %xmm7
   149: 660f6ff7                 movdqa %xmm7, %xmm6
   153: 660f3800f4               pshufb %xmm4, %xmm6
   158: 660f6fc5                 movdqa %xmm5, %xmm0
   162: 660fefc6                 pxor %xmm6, %xmm0
   166: 66410f6fbfc0000000       movdqa 0xc0(%r15), %xmm7
   175: 660f6fef                 movdqa %xmm7, %xmm5
   179: 660f3800eb               pshufb %xmm3, %xmm5
   184: 66410f6fbf00010000       movdqa 0x100(%r15), %xmm7
   193: 660f6ff7                 movdqa %xmm7, %xmm6
   197: 660f3800f4               pshufb %xmm4, %xmm6
   202: 660f6fcd                 movdqa %xmm5, %xmm1
   206: 660fefce                 pxor %xmm6, %xmm1
   210: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
   216: 660f6fda                 movdqa %xmm2, %xmm3
   220: 66410fdbd8               pand %xmm8, %xmm3
   225: 660f6fe2                 movdqa %xmm2, %xmm4
   229: 660f71d404               psrlw $4, %xmm4
   234: 66410fdbe0               pand %xmm8, %xmm4
   239: 66410f6fbf40010000       movdqa 0x140(%r15), %xmm7
   248: 660f6fef                 movdqa %xmm7, %xmm5
   252: 660f3800eb               pshufb %xmm3, %xmm5
   257: 66410f6fbf80010000       movdqa 0x180(%r15), %xmm7
   266: 660f6ff7                 movdqa %xmm7, %xmm6
   270: 660f3800f4               pshufb %xmm4, %xmm6
   275: 660fefc5                 pxor %xmm5, %xmm0
   279: 660fefc6                 pxor %xmm6, %xmm0
   283: 66410f6fbfc0010000       movdqa 0x1c0(%r15), %xmm7
   292: 660f6fef                 movdqa %xmm7, %xmm5
   296: 660f3800eb               pshufb %xmm3, %xmm5
   301: 66410f6fbf00020000       movdqa 0x200(%r15), %xmm7
   310: 660f6ff7                 movdqa %xmm7, %xmm6
   314: 660f3800f4               pshufb %xmm4, %xmm6
   319: 660fefcd                 pxor %xmm5, %xmm1
   323: 660fefce                 pxor %xmm6, %xmm1
   327: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   333: 660f6fda                 movdqa %xmm2, %xmm3
   337: 66410fdbd8               pand %xmm8, %xmm3
   342: 660f6fe2                 movdqa %xmm2, %xmm4
   346: 660f71d404               psrlw $4, %xmm4
   351: 66410fdbe0               pand %xmm8, %xmm4
   356: 66410f6fbf40020000       movdqa 0x240(%r15), %xmm7
   365: 660f6fef                 movdqa %xmm7, %xmm5
   369: 660f3800eb               pshufb %xmm3, %xmm5
   374: 66410f6fbf80020000       movdqa 0x280(%r15), %xmm7
   383: 660f6ff7                 movdqa %xmm7, %xmm6
   387: 660f3800f4               pshufb %xmm4, %xmm6
   392: 660fefc5                 pxor %xmm5, %xmm0
   396: 660fefc6                 pxor %xmm6, %xmm0
   400: 66410f6fbfc0020000       movdqa 0x2c0(%r15), %xmm7
   409: 660f6fef                 movdqa %xmm7, %xmm5
   413: 660f3800eb               pshufb %xmm3, %xmm5
   418: 66410f6fbf00030000       movdqa 0x300(%r15), %xmm7
   427: 660f6ff7                 movdqa %xmm7, %xmm6
   431: 660f3800f4               pshufb %xmm4, %xmm6
   436: 660fefcd                 pxor %xmm5, %xmm1
   440: 660fefce                 pxor %xmm6, %xmm1
   444: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
   451: 660f6fda                 movdqa %xmm2, %xmm3
   455: 66410fdbd8               pand %xmm8, %xmm3
   460: 660f6fe2                 movdqa %xmm2, %xmm4
   464: 660f71d404               psrlw $4, %xmm4
   469: 66410fdbe0               pand %xmm8, %xmm4
   474: 66410f6fbf40030000       movdqa 0x340(%r15), %xmm7
   483: 660f6fef                 movdqa %xmm7, %xmm5
   487: 660f3800eb               pshufb %xmm3, %xmm5
   492: 66410f6fbf80030000       movdqa 0x380(%r15), %xmm7
   501: 660f6ff7                 movdqa %xmm7, %xmm6
   505: 660f3800f4               pshufb %xmm4, %xmm6
   510: 660fefc5                 pxor %xmm5, %xmm0
   514: 660fefc6                 pxor %xmm6, %xmm0
   518: 66410f6fbfc0030000       movdqa 0x3c0(%r15), %xmm7
   527: 660f6fef                 movdqa %xmm7, %xmm5
   531: 660f3800eb               pshufb %xmm3, %xmm5
   536: 66410f6fbf00040000       movdqa 0x400(%r15), %xmm7
   545: 660f6ff7                 movdqa %xmm7, %xmm6
   549: 660f3800f4               pshufb %xmm4, %xmm6
   554: 660fefcd                 pxor %xmm5, %xmm1
   558: 660fefce                 pxor %xmm6, %xmm1
   562: 66430f6f141e             movdqa (%r14, %r11), %xmm2
   568: 660f6fda                 movdqa %xmm2, %xmm3
   572: 66410fdbd8               pand %xmm8, %xmm3
   577: 660f6fe2                 movdqa %xmm2, %xmm4
   581: 660f71d404               psrlw $4, %xmm4
   586: 66410fdbe0               pand %xmm8, %xmm4
   591: 66410f6fbf40040000       movdqa 0x440(%r15), %xmm7
   600: 660f6fef                 movdqa %xmm7, %xmm5
   604: 660f3800eb               pshufb %xmm3, %xmm5
   609: 66410f6fbf80040000       movdqa 0x480(%r15), %xmm7
   618: 660f6ff7                 movdqa %xmm7, %xmm6
   622: 660f3800f4               pshufb %xmm4, %xmm6
   627: 660fefc5                 pxor %xmm5, %xmm0
   631: 660fefc6                 pxor %xmm6, %xmm0
   635: 66410f6fbfc0040000       movdqa 0x4c0(%r15), %xmm7
   644: 660f6fef                 movdqa %xmm7, %xmm5
   648: 660f3800eb               pshufb %xmm3, %xmm5
   653: 66410f6fbf00050000       movdqa 0x500(%r15), %xmm7
   662: 660f6ff7                 movdqa %xmm7, %xmm6
   666: 660f3800f4               pshufb %xmm4, %xmm6
   671: 660fefcd                 pxor %xmm5, %xmm1
   675: 660fefce                 pxor %xmm6, %xmm1
   679: 4d8b5028                 movq 0x28(%r8), %r10
   683: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   689: 660f6fda                 movdqa %xmm2, %xmm3
   693: 66410fdbd8               pand %xmm8, %xmm3
   698: 660f6fe2                 movdqa %xmm2, %xmm4
   702: 660f71d404               psrlw $4, %xmm4
   707: 66410fdbe0               pand %xmm8, %xmm4
   712: 66410f6fbf40050000       movdqa 0x540(%r15), %xmm7
   721: 660f6fef                 movdqa %xmm7, %xmm5
   725: 660f3800eb               pshufb %xmm3, %xmm5
   730: 66410f6fbf80050000       movdqa 0x580(%r15), %xmm7
   739: 660f6ff7                 movdqa %xmm7, %xmm6
   743: 660f3800f4               pshufb %xmm4, %xmm6
   748: 660fefc5                 pxor %xmm5, %xmm0
   752: 660fefc6                 pxor %xmm6, %xmm0
   756: 66410f6fbf40050000       movdqa 0x540(%r15), %xmm7
   765: 660f6fef                 movdqa %xmm7, %xmm5
   769: 660f3800eb               pshufb %xmm3, %xmm5
   774: 66410f6fbf80050000       movdqa 0x580(%r15), %xmm7
   783: 660f6ff7                 movdqa %xmm7, %xmm6
   787: 660f3800f4               pshufb %xmm4, %xmm6
   792: 660fefcd                 pxor %xmm5, %xmm1
   796: 660fefce                 pxor %xmm6, %xmm1
   800: 4d8b11                   movq (%r9), %r10
   803: 66430fe7041a             movntdq %xmm0, (%r10, %r11)
   809: 4d8b5108                 movq 8(%r9), %r10
   813: 66430fe70c1a             movntdq %xmm1, (%r10, %r11)
   819: 4983c310                 addq $0x10, %r11
   823: 4c3bd8                   cmpq %rax, %r11
   826: 0f8620fdffff             jbe 0x60
   832: 4c3bd9                   cmpq %rcx, %r11
   835: 0f83a1020000             jae 0x5ea
   841: 4153                     pushq %r11
   843: 498dbfc0050000           leaq 0x5c0(%r15), %rdi
   850: 498db7c0060000           leaq 0x6c0(%r15), %rsi
   857: 4d8b11                   movq (%r9), %r10
   860: 4c8b1c24                 movq (%rsp), %r11
   864: 33db                     xorl %ebx, %ebx
   866: 498b00                   movq (%r8), %rax
   869: 420fb61418               movzbl (%rax, %r11), %edx
   874: 4885d2                   testq %rdx, %rdx
   877: 0f8422000000             je 0x395
   883: 0fb60417                 movzbl (%rdi, %rdx), %eax
   887: 81c0c1000000             addl $0xc1, %eax
   893: 81f8ff000000             cmpl $0xff, %eax
   899: 0f8206000000             jb 0x38f
   905: 81c001ffffff             addl $0xffffff01, %eax
   911: 0fb60406                 movzbl (%rsi, %rax), %eax
   915: 32d8                     xorb %al, %bl
   917: 498b4008                 movq 8(%r8), %rax
   921: 420fb61418               movzbl (%rax, %r11), %edx
   926: 4885d2                   testq %rdx, %rdx
   929: 0f841f000000             je 0x3c6
   935: 0fb60417                 movzbl (%rdi, %rdx), %eax
   939: 83c003                   addl $3, %eax
   942: 81f8ff000000             cmpl $0xff, %eax
   948: 0f8206000000             jb 0x3c0
   954: 81c001ffffff             addl $0xffffff01, %eax
   960: 0fb60406                 movzbl (%rsi, %rax), %eax
   964: 32d8                     xorb %al, %bl
   966: 498b4010                 movq 0x10(%r8), %rax
   970: 420fb61418               movzbl (%rax, %r11), %edx
   975: 4885d2                   testq %rdx, %rdx
   978: 0f8422000000             je 0x3fa
   984: 0fb60417                 movzbl (%rdi, %rdx), %eax
   988: 81c0c6000000             addl $0xc6, %eax
   994: 81f8ff000000             cmpl $0xff, %eax
  1000: 0f8206000000             jb 0x3f4
  1006: 81c001ffffff             addl $0xffffff01, %eax
  1012: 0fb60406                 movzbl (%rsi, %rax), %eax
  1016: 32d8                     xorb %al, %bl
  1018: 498b4018                 movq 0x18(%r8), %rax
  1022: 420fb61418               movzbl (%rax, %r11), %edx
  1027: 4885d2                   testq %rdx, %rdx
  1030: 0f841f000000             je 0x42b
  1036: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1040: 83c001                   addl $1, %eax
  1043: 81f8ff000000             cmpl $0xff, %eax
  1049: 0f8206000000             jb 0x425
  1055: 81c001ffffff             addl $0xffffff01, %eax
  1061: 0fb60406                 movzbl (%rsi, %rax), %eax
  1065: 32d8                     xorb %al, %bl
  1067: 498b4020                 movq 0x20(%r8), %rax
  1071: 420fb61418               movzbl (%rax, %r11), %edx
  1076: 4885d2                   testq %rdx, %rdx
  1079: 0f8422000000             je 0x45f
  1085: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1089: 81c0fe000000             addl $0xfe, %eax
  1095: 81f8ff000000             cmpl $0xff, %eax
  1101: 0f8206000000             jb 0x459
  1107: 81c001ffffff             addl $0xffffff01, %eax
  1113: 0fb60406                 movzbl (%rsi, %rax), %eax
  1117: 32d8                     xorb %al, %bl
  1119: 498b4028                 movq 0x28(%r8), %rax
  1123: 420fb61418               movzbl (%rax, %r11), %edx
  1128: 4885d2                   testq %rdx, %rdx
  1131: 0f841f000000             je 0x490
  1137: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1141: 83c031                   addl $0x31, %eax
  1144: 81f8ff000000             cmpl $0xff, %eax
  1150: 0f8206000000             jb 0x48a
  1156: 81c001ffffff             addl $0xffffff01, %eax
  1162: 0fb60406                 movzbl (%rsi, %rax), %eax
  1166: 32d8                     xorb %al, %bl
  1168: 43881c1a                 movb %bl, (%r10, %r11)
  1172: 33db                     xorl %ebx, %ebx
  1174: 49ffc3                   incq %r11
  1177: 4c3bd9                   cmpq %rcx, %r11
  1180: 0f82c0feffff             jb 0x362
  1186: 4d8b5108                 movq 8(%r9), %r10
  1190: 4c8b1c24                 movq (%rsp), %r11
  1194: 33db                     xorl %ebx, %ebx
  1196: 498b00                   movq (%r8), %rax
  1199: 420fb61418               movzbl (%rax, %r11), %edx
  1204: 4885d2                   testq %rdx, %rdx
  1207: 0f841f000000             je 0x4dc
  1213: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1217: 83c01c                   addl $0x1c, %eax
  1220: 81f8ff000000             cmpl $0xff, %eax
  1226: 0f8206000000             jb 0x4d6
  1232: 81c001ffffff             addl $0xffffff01, %eax
  1238: 0fb60406                 movzbl (%rsi, %rax), %eax
  1242: 32d8                     xorb %al, %bl
  1244: 498b4008                 movq 8(%r8), %rax
  1248: 420fb61418               movzbl (%rax, %r11), %edx
  1253: 4885d2                   testq %rdx, %rdx
  1256: 0f8422000000             je 0x510
  1262: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1266: 81c0df000000             addl $0xdf, %eax
  1272: 81f8ff000000             cmpl $0xff, %eax
  1278: 0f8206000000             jb 0x50a
  1284: 81c001ffffff             addl $0xffffff01, %eax
  1290: 0fb60406                 movzbl (%rsi, %rax), %eax
  1294: 32d8                     xorb %al, %bl
  1296: 498b4010                 movq 0x10(%r8), %rax
  1300: 420fb61418               movzbl (%rax, %r11), %edx
  1305: 4885d2                   testq %rdx, %rdx
  1308: 0f841f000000             je 0x541
  1314: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1318: 83c01a                   addl $0x1a, %eax
  1321: 81f8ff000000             cmpl $0xff, %eax
  1327: 0f8206000000             jb 0x53b
  1333: 81c001ffffff             addl $0xffffff01, %eax
  1339: 0fb60406                 movzbl (%rsi, %rax), %eax
  1343: 32d8                     xorb %al, %bl
  1345: 498b4018                 movq 0x18(%r8), %rax
  1349: 420fb61418               movzbl (%rax, %r11), %edx
  1354: 4885d2                   testq %rdx, %rdx
  1357: 0f841f000000             je 0x572
  1363: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1367: 83c019                   addl $0x19, %eax
  1370: 81f8ff000000             cmpl $0xff, %eax
  1376: 0f8206000000             jb 0x56c
  1382: 81c001ffffff             addl $0xffffff01, %eax
  1388: 0fb60406                 movzbl (%rsi, %rax), %eax
  1392: 32d8                     xorb %al, %bl
  1394: 498b4020                 movq 0x20(%r8), %rax
  1398: 420fb61418               movzbl (%rax, %r11), %edx
  1403: 4885d2                   testq %rdx, %rdx
  1406: 0f841f000000             je 0x5a3
  1412: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1416: 83c018                   addl $0x18, %eax
  1419: 81f8ff000000             cmpl $0xff, %eax
  1425: 0f8206000000             jb 0x59d
  1431: 81c001ffffff             addl $0xffffff01, %eax
  1437: 0fb60406                 movzbl (%rsi, %rax), %eax
  1441: 32d8                     xorb %al, %bl
  1443: 498b4028                 movq 0x28(%r8), %rax
  1447: 420fb61418               movzbl (%rax, %r11), %edx
  1452: 4885d2                   testq %rdx, %rdx
  1455: 0f841f000000             je 0x5d4
  1461: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1465: 83c031                   addl $0x31, %eax
  1468: 81f8ff000000             cmpl $0xff, %eax
  1474: 0f8206000000             jb 0x5ce
  1480: 81c001ffffff             addl $0xffffff01, %eax
  1486: 0fb60406                 movzbl (%rsi, %rax), %eax
  1490: 32d8                     xorb %al, %bl
  1492: 43881c1a                 movb %bl, (%r10, %r11)
  1496: 33db                     xorl %ebx, %ebx
  1498: 49ffc3                   incq %r11
  1501: 4c3bd9                   cmpq %rcx, %r11
  1504: 0f82c6feffff             jb 0x4ac
  1510: 4883c408                 addq $8, %rsp
  1514: 415e                     popq %r14
  1516: 415d                     popq %r13
  1518: 415c                     popq %r12
  1520: 415f                     popq %r15
  1522: 5b                       popq %rbx
  1523: c3                       retq 
  1524: 0f0b                     ud2 
  1526: 0f0b                     ud2 
  1528: 0f0b                     ud2 
  1530: 0f0b                     ud2 
  1532: 0f0b                     ud2 
  1534: 0f0b                     ud2 
