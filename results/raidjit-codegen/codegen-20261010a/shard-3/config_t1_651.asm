; config_t1_651 target=sse twin=nt topo=rs(6,3)/enc op=encode U=1
; code_size=1536 twins=["nt"] uops_per_byte=7.6875
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
    47: 66450f6f0f               movdqa (%r15), %xmm9
    52: 488bc1                   movq %rcx, %rax
    55: 4883e810                 subq $0x10, %rax
    59: 0f82cb020000             jb 0x30c
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66420f6f1c1e             movdqa (%rsi, %r11), %xmm3
   102: 660f6fe3                 movdqa %xmm3, %xmm4
   106: 66410fdbe1               pand %xmm9, %xmm4
   111: 660f6feb                 movdqa %xmm3, %xmm5
   115: 660f71d504               psrlw $4, %xmm5
   120: 66410fdbe9               pand %xmm9, %xmm5
   125: 660f6fc3                 movdqa %xmm3, %xmm0
   129: 66450f6f4740             movdqa 0x40(%r15), %xmm8
   135: 66410f6ff0               movdqa %xmm8, %xmm6
   140: 660f3800f4               pshufb %xmm4, %xmm6
   145: 66450f6f8780000000       movdqa 0x80(%r15), %xmm8
   154: 66410f6ff8               movdqa %xmm8, %xmm7
   159: 660f3800fd               pshufb %xmm5, %xmm7
   164: 660f6fce                 movdqa %xmm6, %xmm1
   168: 660fefcf                 pxor %xmm7, %xmm1
   172: 66450f6f87c0000000       movdqa 0xc0(%r15), %xmm8
   181: 66410f6ff0               movdqa %xmm8, %xmm6
   186: 660f3800f4               pshufb %xmm4, %xmm6
   191: 66450f6f8700010000       movdqa 0x100(%r15), %xmm8
   200: 66410f6ff8               movdqa %xmm8, %xmm7
   205: 660f3800fd               pshufb %xmm5, %xmm7
   210: 660f6fd6                 movdqa %xmm6, %xmm2
   214: 660fefd7                 pxor %xmm7, %xmm2
   218: 66420f6f1c1f             movdqa (%rdi, %r11), %xmm3
   224: 660f6fe3                 movdqa %xmm3, %xmm4
   228: 66410fdbe1               pand %xmm9, %xmm4
   233: 660f6feb                 movdqa %xmm3, %xmm5
   237: 660f71d504               psrlw $4, %xmm5
   242: 66410fdbe9               pand %xmm9, %xmm5
   247: 660fefc3                 pxor %xmm3, %xmm0
   251: 66450f6f87c0000000       movdqa 0xc0(%r15), %xmm8
   260: 66410f6ff0               movdqa %xmm8, %xmm6
   265: 660f3800f4               pshufb %xmm4, %xmm6
   270: 66450f6f8700010000       movdqa 0x100(%r15), %xmm8
   279: 66410f6ff8               movdqa %xmm8, %xmm7
   284: 660f3800fd               pshufb %xmm5, %xmm7
   289: 660fefce                 pxor %xmm6, %xmm1
   293: 660fefcf                 pxor %xmm7, %xmm1
   297: 66450f6f8740010000       movdqa 0x140(%r15), %xmm8
   306: 66410f6ff0               movdqa %xmm8, %xmm6
   311: 660f3800f4               pshufb %xmm4, %xmm6
   316: 66450f6f8780010000       movdqa 0x180(%r15), %xmm8
   325: 66410f6ff8               movdqa %xmm8, %xmm7
   330: 660f3800fd               pshufb %xmm5, %xmm7
   335: 660fefd6                 pxor %xmm6, %xmm2
   339: 660fefd7                 pxor %xmm7, %xmm2
   343: 66430f6f1c1c             movdqa (%r12, %r11), %xmm3
   349: 660f6fe3                 movdqa %xmm3, %xmm4
   353: 66410fdbe1               pand %xmm9, %xmm4
   358: 660f6feb                 movdqa %xmm3, %xmm5
   362: 660f71d504               psrlw $4, %xmm5
   367: 66410fdbe9               pand %xmm9, %xmm5
   372: 660fefc3                 pxor %xmm3, %xmm0
   376: 66450f6f8740010000       movdqa 0x140(%r15), %xmm8
   385: 66410f6ff0               movdqa %xmm8, %xmm6
   390: 660f3800f4               pshufb %xmm4, %xmm6
   395: 66450f6f8780010000       movdqa 0x180(%r15), %xmm8
   404: 66410f6ff8               movdqa %xmm8, %xmm7
   409: 660f3800fd               pshufb %xmm5, %xmm7
   414: 660fefce                 pxor %xmm6, %xmm1
   418: 660fefcf                 pxor %xmm7, %xmm1
   422: 66450f6f87c0010000       movdqa 0x1c0(%r15), %xmm8
   431: 66410f6ff0               movdqa %xmm8, %xmm6
   436: 660f3800f4               pshufb %xmm4, %xmm6
   441: 66450f6f8700020000       movdqa 0x200(%r15), %xmm8
   450: 66410f6ff8               movdqa %xmm8, %xmm7
   455: 660f3800fd               pshufb %xmm5, %xmm7
   460: 660fefd6                 pxor %xmm6, %xmm2
   464: 660fefd7                 pxor %xmm7, %xmm2
   468: 66430f6f5c1d00           movdqa (%r13, %r11), %xmm3
   475: 660f6fe3                 movdqa %xmm3, %xmm4
   479: 66410fdbe1               pand %xmm9, %xmm4
   484: 660f6feb                 movdqa %xmm3, %xmm5
   488: 660f71d504               psrlw $4, %xmm5
   493: 66410fdbe9               pand %xmm9, %xmm5
   498: 660fefc3                 pxor %xmm3, %xmm0
   502: 66450f6f8740020000       movdqa 0x240(%r15), %xmm8
   511: 66410f6ff0               movdqa %xmm8, %xmm6
   516: 660f3800f4               pshufb %xmm4, %xmm6
   521: 66450f6f8780020000       movdqa 0x280(%r15), %xmm8
   530: 66410f6ff8               movdqa %xmm8, %xmm7
   535: 660f3800fd               pshufb %xmm5, %xmm7
   540: 660fefce                 pxor %xmm6, %xmm1
   544: 660fefcf                 pxor %xmm7, %xmm1
   548: 66450f6f4740             movdqa 0x40(%r15), %xmm8
   554: 66410f6ff0               movdqa %xmm8, %xmm6
   559: 660f3800f4               pshufb %xmm4, %xmm6
   564: 66450f6f8780000000       movdqa 0x80(%r15), %xmm8
   573: 66410f6ff8               movdqa %xmm8, %xmm7
   578: 660f3800fd               pshufb %xmm5, %xmm7
   583: 660fefd6                 pxor %xmm6, %xmm2
   587: 660fefd7                 pxor %xmm7, %xmm2
   591: 66430f6f1c1e             movdqa (%r14, %r11), %xmm3
   597: 660f6fe3                 movdqa %xmm3, %xmm4
   601: 66410fdbe1               pand %xmm9, %xmm4
   606: 660f6feb                 movdqa %xmm3, %xmm5
   610: 660f71d504               psrlw $4, %xmm5
   615: 66410fdbe9               pand %xmm9, %xmm5
   620: 660fefc3                 pxor %xmm3, %xmm0
   624: 66450f6f87c0010000       movdqa 0x1c0(%r15), %xmm8
   633: 66410f6ff0               movdqa %xmm8, %xmm6
   638: 660f3800f4               pshufb %xmm4, %xmm6
   643: 66450f6f8700020000       movdqa 0x200(%r15), %xmm8
   652: 66410f6ff8               movdqa %xmm8, %xmm7
   657: 660f3800fd               pshufb %xmm5, %xmm7
   662: 660fefce                 pxor %xmm6, %xmm1
   666: 660fefcf                 pxor %xmm7, %xmm1
   670: 66450f6f8740020000       movdqa 0x240(%r15), %xmm8
   679: 66410f6ff0               movdqa %xmm8, %xmm6
   684: 660f3800f4               pshufb %xmm4, %xmm6
   689: 66450f6f8780020000       movdqa 0x280(%r15), %xmm8
   698: 66410f6ff8               movdqa %xmm8, %xmm7
   703: 660f3800fd               pshufb %xmm5, %xmm7
   708: 660fefd6                 pxor %xmm6, %xmm2
   712: 660fefd7                 pxor %xmm7, %xmm2
   716: 4d8b5028                 movq 0x28(%r8), %r10
   720: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   726: 660fefc3                 pxor %xmm3, %xmm0
   730: 660fefcb                 pxor %xmm3, %xmm1
   734: 660fefd3                 pxor %xmm3, %xmm2
   738: 4d8b11                   movq (%r9), %r10
   741: 66430fe7041a             movntdq %xmm0, (%r10, %r11)
   747: 4d8b5108                 movq 8(%r9), %r10
   751: 66430fe70c1a             movntdq %xmm1, (%r10, %r11)
   757: 4d8b5110                 movq 0x10(%r9), %r10
   761: 66430fe7141a             movntdq %xmm2, (%r10, %r11)
   767: 4983c310                 addq $0x10, %r11
   771: 4c3bd8                   cmpq %rax, %r11
   774: 0f8654fdffff             jbe 0x60
   780: 4c3bd9                   cmpq %rcx, %r11
   783: 0f83b2020000             jae 0x5c7
   789: 4153                     pushq %r11
   791: 498dbfc0020000           leaq 0x2c0(%r15), %rdi
   798: 498db7c0030000           leaq 0x3c0(%r15), %rsi
   805: 4d8b11                   movq (%r9), %r10
   808: 4c8b1c24                 movq (%rsp), %r11
   812: 33db                     xorl %ebx, %ebx
   814: 498b00                   movq (%r8), %rax
   817: 420fb61418               movzbl (%rax, %r11), %edx
   822: 32da                     xorb %dl, %bl
   824: 498b4008                 movq 8(%r8), %rax
   828: 420fb61418               movzbl (%rax, %r11), %edx
   833: 32da                     xorb %dl, %bl
   835: 498b4010                 movq 0x10(%r8), %rax
   839: 420fb61418               movzbl (%rax, %r11), %edx
   844: 32da                     xorb %dl, %bl
   846: 498b4018                 movq 0x18(%r8), %rax
   850: 420fb61418               movzbl (%rax, %r11), %edx
   855: 32da                     xorb %dl, %bl
   857: 498b4020                 movq 0x20(%r8), %rax
   861: 420fb61418               movzbl (%rax, %r11), %edx
   866: 32da                     xorb %dl, %bl
   868: 498b4028                 movq 0x28(%r8), %rax
   872: 420fb61418               movzbl (%rax, %r11), %edx
   877: 32da                     xorb %dl, %bl
   879: 43881c1a                 movb %bl, (%r10, %r11)
   883: 33db                     xorl %ebx, %ebx
   885: 49ffc3                   incq %r11
   888: 4c3bd9                   cmpq %rcx, %r11
   891: 0f82adffffff             jb 0x32e
   897: 4d8b5108                 movq 8(%r9), %r10
   901: 4c8b1c24                 movq (%rsp), %r11
   905: 33db                     xorl %ebx, %ebx
   907: 498b00                   movq (%r8), %rax
   910: 420fb61418               movzbl (%rax, %r11), %edx
   915: 4885d2                   testq %rdx, %rdx
   918: 0f8422000000             je 0x3be
   924: 0fb60417                 movzbl (%rdi, %rdx), %eax
   928: 81c0e6000000             addl $0xe6, %eax
   934: 81f8ff000000             cmpl $0xff, %eax
   940: 0f8206000000             jb 0x3b8
   946: 81c001ffffff             addl $0xffffff01, %eax
   952: 0fb60406                 movzbl (%rsi, %rax), %eax
   956: 32d8                     xorb %al, %bl
   958: 498b4008                 movq 8(%r8), %rax
   962: 420fb61418               movzbl (%rax, %r11), %edx
   967: 4885d2                   testq %rdx, %rdx
   970: 0f8422000000             je 0x3f2
   976: 0fb60417                 movzbl (%rdi, %rdx), %eax
   980: 81c0cd000000             addl $0xcd, %eax
   986: 81f8ff000000             cmpl $0xff, %eax
   992: 0f8206000000             jb 0x3ec
   998: 81c001ffffff             addl $0xffffff01, %eax
  1004: 0fb60406                 movzbl (%rsi, %rax), %eax
  1008: 32d8                     xorb %al, %bl
  1010: 498b4010                 movq 0x10(%r8), %rax
  1014: 420fb61418               movzbl (%rax, %r11), %edx
  1019: 4885d2                   testq %rdx, %rdx
  1022: 0f841f000000             je 0x423
  1028: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1032: 83c04b                   addl $0x4b, %eax
  1035: 81f8ff000000             cmpl $0xff, %eax
  1041: 0f8206000000             jb 0x41d
  1047: 81c001ffffff             addl $0xffffff01, %eax
  1053: 0fb60406                 movzbl (%rsi, %rax), %eax
  1057: 32d8                     xorb %al, %bl
  1059: 498b4018                 movq 0x18(%r8), %rax
  1063: 420fb61418               movzbl (%rax, %r11), %edx
  1068: 4885d2                   testq %rdx, %rdx
  1071: 0f841f000000             je 0x454
  1077: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1081: 83c032                   addl $0x32, %eax
  1084: 81f8ff000000             cmpl $0xff, %eax
  1090: 0f8206000000             jb 0x44e
  1096: 81c001ffffff             addl $0xffffff01, %eax
  1102: 0fb60406                 movzbl (%rsi, %rax), %eax
  1106: 32d8                     xorb %al, %bl
  1108: 498b4020                 movq 0x20(%r8), %rax
  1112: 420fb61418               movzbl (%rax, %r11), %edx
  1117: 4885d2                   testq %rdx, %rdx
  1120: 0f841f000000             je 0x485
  1126: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1130: 83c019                   addl $0x19, %eax
  1133: 81f8ff000000             cmpl $0xff, %eax
  1139: 0f8206000000             jb 0x47f
  1145: 81c001ffffff             addl $0xffffff01, %eax
  1151: 0fb60406                 movzbl (%rsi, %rax), %eax
  1155: 32d8                     xorb %al, %bl
  1157: 498b4028                 movq 0x28(%r8), %rax
  1161: 420fb61418               movzbl (%rax, %r11), %edx
  1166: 32da                     xorb %dl, %bl
  1168: 43881c1a                 movb %bl, (%r10, %r11)
  1172: 33db                     xorl %ebx, %ebx
  1174: 49ffc3                   incq %r11
  1177: 4c3bd9                   cmpq %rcx, %r11
  1180: 0f82e9feffff             jb 0x38b
  1186: 4d8b5110                 movq 0x10(%r9), %r10
  1190: 4c8b1c24                 movq (%rsp), %r11
  1194: 33db                     xorl %ebx, %ebx
  1196: 498b00                   movq (%r8), %rax
  1199: 420fb61418               movzbl (%rax, %r11), %edx
  1204: 4885d2                   testq %rdx, %rdx
  1207: 0f8422000000             je 0x4df
  1213: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1217: 81c0cd000000             addl $0xcd, %eax
  1223: 81f8ff000000             cmpl $0xff, %eax
  1229: 0f8206000000             jb 0x4d9
  1235: 81c001ffffff             addl $0xffffff01, %eax
  1241: 0fb60406                 movzbl (%rsi, %rax), %eax
  1245: 32d8                     xorb %al, %bl
  1247: 498b4008                 movq 8(%r8), %rax
  1251: 420fb61418               movzbl (%rax, %r11), %edx
  1256: 4885d2                   testq %rdx, %rdx
  1259: 0f841f000000             je 0x510
  1265: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1269: 83c04b                   addl $0x4b, %eax
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
  1318: 83c019                   addl $0x19, %eax
  1321: 81f8ff000000             cmpl $0xff, %eax
  1327: 0f8206000000             jb 0x53b
  1333: 81c001ffffff             addl $0xffffff01, %eax
  1339: 0fb60406                 movzbl (%rsi, %rax), %eax
  1343: 32d8                     xorb %al, %bl
  1345: 498b4018                 movq 0x18(%r8), %rax
  1349: 420fb61418               movzbl (%rax, %r11), %edx
  1354: 4885d2                   testq %rdx, %rdx
  1357: 0f8422000000             je 0x575
  1363: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1367: 81c0e6000000             addl $0xe6, %eax
  1373: 81f8ff000000             cmpl $0xff, %eax
  1379: 0f8206000000             jb 0x56f
  1385: 81c001ffffff             addl $0xffffff01, %eax
  1391: 0fb60406                 movzbl (%rsi, %rax), %eax
  1395: 32d8                     xorb %al, %bl
  1397: 498b4020                 movq 0x20(%r8), %rax
  1401: 420fb61418               movzbl (%rax, %r11), %edx
  1406: 4885d2                   testq %rdx, %rdx
  1409: 0f841f000000             je 0x5a6
  1415: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1419: 83c032                   addl $0x32, %eax
  1422: 81f8ff000000             cmpl $0xff, %eax
  1428: 0f8206000000             jb 0x5a0
  1434: 81c001ffffff             addl $0xffffff01, %eax
  1440: 0fb60406                 movzbl (%rsi, %rax), %eax
  1444: 32d8                     xorb %al, %bl
  1446: 498b4028                 movq 0x28(%r8), %rax
  1450: 420fb61418               movzbl (%rax, %r11), %edx
  1455: 32da                     xorb %dl, %bl
  1457: 43881c1a                 movb %bl, (%r10, %r11)
  1461: 33db                     xorl %ebx, %ebx
  1463: 49ffc3                   incq %r11
  1466: 4c3bd9                   cmpq %rcx, %r11
  1469: 0f82e9feffff             jb 0x4ac
  1475: 4883c408                 addq $8, %rsp
  1479: 415e                     popq %r14
  1481: 415d                     popq %r13
  1483: 415c                     popq %r12
  1485: 415f                     popq %r15
  1487: 5b                       popq %rbx
  1488: c3                       retq 
  1489: 0f0b                     ud2 
  1491: 0f0b                     ud2 
  1493: 0f0b                     ud2 
  1495: 0f0b                     ud2 
  1497: 0f0b                     ud2 
  1499: 0f0b                     ud2 
  1501: 0f0b                     ud2 
  1503: 0f0b                     ud2 
  1505: 0f0b                     ud2 
  1507: 0f0b                     ud2 
  1509: 0f0b                     ud2 
  1511: 0f0b                     ud2 
  1513: 0f0b                     ud2 
  1515: 0f0b                     ud2 
  1517: 0f0b                     ud2 
  1519: 0f0b                     ud2 
  1521: 0f0b                     ud2 
  1523: 0f0b                     ud2 
  1525: 0f0b                     ud2 
  1527: 0f0b                     ud2 
  1529: 0f0b                     ud2 
  1531: 0f0b                     ud2 
  1533: 0f0b                     ud2 
  1535: cc                       int3 
