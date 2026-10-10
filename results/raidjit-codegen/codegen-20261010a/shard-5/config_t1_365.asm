; config_t1_365 target=avx2 twin=nt topo=rs(6,3)/enc op=encode U=1
; code_size=1408 twins=["nt"] uops_per_byte=2.84375
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
    47: c4417d6f0f               vmovdqa (%r15), %ymm9
    52: 488bc1                   movq %rcx, %rax
    55: 4883e820                 subq $0x20, %rax
    59: 0f8237020000             jb 0x278
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4a17d6f1c1e             vmovdqa (%rsi, %r11), %ymm3
   102: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   107: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   112: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   117: c5fd6fc3                 vmovdqa %ymm3, %ymm0
   121: c4417d6f4740             vmovdqa 0x40(%r15), %ymm8
   127: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   132: c4417d6f8780000000       vmovdqa 0x80(%r15), %ymm8
   141: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   146: c5cdefcf                 vpxor %ymm7, %ymm6, %ymm1
   150: c4417d6f87c0000000       vmovdqa 0xc0(%r15), %ymm8
   159: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   164: c4417d6f8700010000       vmovdqa 0x100(%r15), %ymm8
   173: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   178: c5cdefd7                 vpxor %ymm7, %ymm6, %ymm2
   182: c4a17d6f1c1f             vmovdqa (%rdi, %r11), %ymm3
   188: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   193: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   198: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   203: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   207: c4417d6f87c0000000       vmovdqa 0xc0(%r15), %ymm8
   216: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   221: c4417d6f8700010000       vmovdqa 0x100(%r15), %ymm8
   230: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   235: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   239: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   243: c4417d6f8740010000       vmovdqa 0x140(%r15), %ymm8
   252: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   257: c4417d6f8780010000       vmovdqa 0x180(%r15), %ymm8
   266: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   271: c5edefd6                 vpxor %ymm6, %ymm2, %ymm2
   275: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   279: c4817d6f1c1c             vmovdqa (%r12, %r11), %ymm3
   285: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   290: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   295: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   300: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   304: c4417d6f8740010000       vmovdqa 0x140(%r15), %ymm8
   313: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   318: c4417d6f8780010000       vmovdqa 0x180(%r15), %ymm8
   327: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   332: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   336: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   340: c4417d6f87c0010000       vmovdqa 0x1c0(%r15), %ymm8
   349: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   354: c4417d6f8700020000       vmovdqa 0x200(%r15), %ymm8
   363: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   368: c5edefd6                 vpxor %ymm6, %ymm2, %ymm2
   372: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   376: c4817d6f5c1d00           vmovdqa (%r13, %r11), %ymm3
   383: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   388: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   393: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   398: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   402: c4417d6f8740020000       vmovdqa 0x240(%r15), %ymm8
   411: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   416: c4417d6f8780020000       vmovdqa 0x280(%r15), %ymm8
   425: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   430: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   434: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   438: c4417d6f4740             vmovdqa 0x40(%r15), %ymm8
   444: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   449: c4417d6f8780000000       vmovdqa 0x80(%r15), %ymm8
   458: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   463: c5edefd6                 vpxor %ymm6, %ymm2, %ymm2
   467: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   471: c4817d6f1c1e             vmovdqa (%r14, %r11), %ymm3
   477: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   482: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   487: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   492: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   496: c4417d6f87c0010000       vmovdqa 0x1c0(%r15), %ymm8
   505: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   510: c4417d6f8700020000       vmovdqa 0x200(%r15), %ymm8
   519: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   524: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   528: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   532: c4417d6f8740020000       vmovdqa 0x240(%r15), %ymm8
   541: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   546: c4417d6f8780020000       vmovdqa 0x280(%r15), %ymm8
   555: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   560: c5edefd6                 vpxor %ymm6, %ymm2, %ymm2
   564: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   568: 4d8b5028                 movq 0x28(%r8), %r10
   572: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   578: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   582: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   586: c5edefd3                 vpxor %ymm3, %ymm2, %ymm2
   590: 4d8b11                   movq (%r9), %r10
   593: c4817de7041a             vmovntdq %ymm0, (%r10, %r11)
   599: 4d8b5108                 movq 8(%r9), %r10
   603: c4817de70c1a             vmovntdq %ymm1, (%r10, %r11)
   609: 4d8b5110                 movq 0x10(%r9), %r10
   613: c4817de7141a             vmovntdq %ymm2, (%r10, %r11)
   619: 4983c320                 addq $0x20, %r11
   623: 4c3bd8                   cmpq %rax, %r11
   626: 0f86e8fdffff             jbe 0x60
   632: 4c3bd9                   cmpq %rcx, %r11
   635: 0f83b2020000             jae 0x533
   641: 4153                     pushq %r11
   643: 498dbfc0020000           leaq 0x2c0(%r15), %rdi
   650: 498db7c0030000           leaq 0x3c0(%r15), %rsi
   657: 4d8b11                   movq (%r9), %r10
   660: 4c8b1c24                 movq (%rsp), %r11
   664: 33db                     xorl %ebx, %ebx
   666: 498b00                   movq (%r8), %rax
   669: 420fb61418               movzbl (%rax, %r11), %edx
   674: 32da                     xorb %dl, %bl
   676: 498b4008                 movq 8(%r8), %rax
   680: 420fb61418               movzbl (%rax, %r11), %edx
   685: 32da                     xorb %dl, %bl
   687: 498b4010                 movq 0x10(%r8), %rax
   691: 420fb61418               movzbl (%rax, %r11), %edx
   696: 32da                     xorb %dl, %bl
   698: 498b4018                 movq 0x18(%r8), %rax
   702: 420fb61418               movzbl (%rax, %r11), %edx
   707: 32da                     xorb %dl, %bl
   709: 498b4020                 movq 0x20(%r8), %rax
   713: 420fb61418               movzbl (%rax, %r11), %edx
   718: 32da                     xorb %dl, %bl
   720: 498b4028                 movq 0x28(%r8), %rax
   724: 420fb61418               movzbl (%rax, %r11), %edx
   729: 32da                     xorb %dl, %bl
   731: 43881c1a                 movb %bl, (%r10, %r11)
   735: 33db                     xorl %ebx, %ebx
   737: 49ffc3                   incq %r11
   740: 4c3bd9                   cmpq %rcx, %r11
   743: 0f82adffffff             jb 0x29a
   749: 4d8b5108                 movq 8(%r9), %r10
   753: 4c8b1c24                 movq (%rsp), %r11
   757: 33db                     xorl %ebx, %ebx
   759: 498b00                   movq (%r8), %rax
   762: 420fb61418               movzbl (%rax, %r11), %edx
   767: 4885d2                   testq %rdx, %rdx
   770: 0f8422000000             je 0x32a
   776: 0fb60417                 movzbl (%rdi, %rdx), %eax
   780: 81c0e6000000             addl $0xe6, %eax
   786: 81f8ff000000             cmpl $0xff, %eax
   792: 0f8206000000             jb 0x324
   798: 81c001ffffff             addl $0xffffff01, %eax
   804: 0fb60406                 movzbl (%rsi, %rax), %eax
   808: 32d8                     xorb %al, %bl
   810: 498b4008                 movq 8(%r8), %rax
   814: 420fb61418               movzbl (%rax, %r11), %edx
   819: 4885d2                   testq %rdx, %rdx
   822: 0f8422000000             je 0x35e
   828: 0fb60417                 movzbl (%rdi, %rdx), %eax
   832: 81c0cd000000             addl $0xcd, %eax
   838: 81f8ff000000             cmpl $0xff, %eax
   844: 0f8206000000             jb 0x358
   850: 81c001ffffff             addl $0xffffff01, %eax
   856: 0fb60406                 movzbl (%rsi, %rax), %eax
   860: 32d8                     xorb %al, %bl
   862: 498b4010                 movq 0x10(%r8), %rax
   866: 420fb61418               movzbl (%rax, %r11), %edx
   871: 4885d2                   testq %rdx, %rdx
   874: 0f841f000000             je 0x38f
   880: 0fb60417                 movzbl (%rdi, %rdx), %eax
   884: 83c04b                   addl $0x4b, %eax
   887: 81f8ff000000             cmpl $0xff, %eax
   893: 0f8206000000             jb 0x389
   899: 81c001ffffff             addl $0xffffff01, %eax
   905: 0fb60406                 movzbl (%rsi, %rax), %eax
   909: 32d8                     xorb %al, %bl
   911: 498b4018                 movq 0x18(%r8), %rax
   915: 420fb61418               movzbl (%rax, %r11), %edx
   920: 4885d2                   testq %rdx, %rdx
   923: 0f841f000000             je 0x3c0
   929: 0fb60417                 movzbl (%rdi, %rdx), %eax
   933: 83c032                   addl $0x32, %eax
   936: 81f8ff000000             cmpl $0xff, %eax
   942: 0f8206000000             jb 0x3ba
   948: 81c001ffffff             addl $0xffffff01, %eax
   954: 0fb60406                 movzbl (%rsi, %rax), %eax
   958: 32d8                     xorb %al, %bl
   960: 498b4020                 movq 0x20(%r8), %rax
   964: 420fb61418               movzbl (%rax, %r11), %edx
   969: 4885d2                   testq %rdx, %rdx
   972: 0f841f000000             je 0x3f1
   978: 0fb60417                 movzbl (%rdi, %rdx), %eax
   982: 83c019                   addl $0x19, %eax
   985: 81f8ff000000             cmpl $0xff, %eax
   991: 0f8206000000             jb 0x3eb
   997: 81c001ffffff             addl $0xffffff01, %eax
  1003: 0fb60406                 movzbl (%rsi, %rax), %eax
  1007: 32d8                     xorb %al, %bl
  1009: 498b4028                 movq 0x28(%r8), %rax
  1013: 420fb61418               movzbl (%rax, %r11), %edx
  1018: 32da                     xorb %dl, %bl
  1020: 43881c1a                 movb %bl, (%r10, %r11)
  1024: 33db                     xorl %ebx, %ebx
  1026: 49ffc3                   incq %r11
  1029: 4c3bd9                   cmpq %rcx, %r11
  1032: 0f82e9feffff             jb 0x2f7
  1038: 4d8b5110                 movq 0x10(%r9), %r10
  1042: 4c8b1c24                 movq (%rsp), %r11
  1046: 33db                     xorl %ebx, %ebx
  1048: 498b00                   movq (%r8), %rax
  1051: 420fb61418               movzbl (%rax, %r11), %edx
  1056: 4885d2                   testq %rdx, %rdx
  1059: 0f8422000000             je 0x44b
  1065: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1069: 81c0cd000000             addl $0xcd, %eax
  1075: 81f8ff000000             cmpl $0xff, %eax
  1081: 0f8206000000             jb 0x445
  1087: 81c001ffffff             addl $0xffffff01, %eax
  1093: 0fb60406                 movzbl (%rsi, %rax), %eax
  1097: 32d8                     xorb %al, %bl
  1099: 498b4008                 movq 8(%r8), %rax
  1103: 420fb61418               movzbl (%rax, %r11), %edx
  1108: 4885d2                   testq %rdx, %rdx
  1111: 0f841f000000             je 0x47c
  1117: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1121: 83c04b                   addl $0x4b, %eax
  1124: 81f8ff000000             cmpl $0xff, %eax
  1130: 0f8206000000             jb 0x476
  1136: 81c001ffffff             addl $0xffffff01, %eax
  1142: 0fb60406                 movzbl (%rsi, %rax), %eax
  1146: 32d8                     xorb %al, %bl
  1148: 498b4010                 movq 0x10(%r8), %rax
  1152: 420fb61418               movzbl (%rax, %r11), %edx
  1157: 4885d2                   testq %rdx, %rdx
  1160: 0f841f000000             je 0x4ad
  1166: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1170: 83c019                   addl $0x19, %eax
  1173: 81f8ff000000             cmpl $0xff, %eax
  1179: 0f8206000000             jb 0x4a7
  1185: 81c001ffffff             addl $0xffffff01, %eax
  1191: 0fb60406                 movzbl (%rsi, %rax), %eax
  1195: 32d8                     xorb %al, %bl
  1197: 498b4018                 movq 0x18(%r8), %rax
  1201: 420fb61418               movzbl (%rax, %r11), %edx
  1206: 4885d2                   testq %rdx, %rdx
  1209: 0f8422000000             je 0x4e1
  1215: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1219: 81c0e6000000             addl $0xe6, %eax
  1225: 81f8ff000000             cmpl $0xff, %eax
  1231: 0f8206000000             jb 0x4db
  1237: 81c001ffffff             addl $0xffffff01, %eax
  1243: 0fb60406                 movzbl (%rsi, %rax), %eax
  1247: 32d8                     xorb %al, %bl
  1249: 498b4020                 movq 0x20(%r8), %rax
  1253: 420fb61418               movzbl (%rax, %r11), %edx
  1258: 4885d2                   testq %rdx, %rdx
  1261: 0f841f000000             je 0x512
  1267: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1271: 83c032                   addl $0x32, %eax
  1274: 81f8ff000000             cmpl $0xff, %eax
  1280: 0f8206000000             jb 0x50c
  1286: 81c001ffffff             addl $0xffffff01, %eax
  1292: 0fb60406                 movzbl (%rsi, %rax), %eax
  1296: 32d8                     xorb %al, %bl
  1298: 498b4028                 movq 0x28(%r8), %rax
  1302: 420fb61418               movzbl (%rax, %r11), %edx
  1307: 32da                     xorb %dl, %bl
  1309: 43881c1a                 movb %bl, (%r10, %r11)
  1313: 33db                     xorl %ebx, %ebx
  1315: 49ffc3                   incq %r11
  1318: 4c3bd9                   cmpq %rcx, %r11
  1321: 0f82e9feffff             jb 0x418
  1327: 4883c408                 addq $8, %rsp
  1331: 415e                     popq %r14
  1333: 415d                     popq %r13
  1335: 415c                     popq %r12
  1337: 415f                     popq %r15
  1339: 5b                       popq %rbx
  1340: c5f877                   vzeroupper 
  1343: c3                       retq 
  1344: 0f0b                     ud2 
  1346: 0f0b                     ud2 
  1348: 0f0b                     ud2 
  1350: 0f0b                     ud2 
  1352: 0f0b                     ud2 
  1354: 0f0b                     ud2 
  1356: 0f0b                     ud2 
  1358: 0f0b                     ud2 
  1360: 0f0b                     ud2 
  1362: 0f0b                     ud2 
  1364: 0f0b                     ud2 
  1366: 0f0b                     ud2 
  1368: 0f0b                     ud2 
  1370: 0f0b                     ud2 
  1372: 0f0b                     ud2 
  1374: 0f0b                     ud2 
  1376: 0f0b                     ud2 
  1378: 0f0b                     ud2 
  1380: 0f0b                     ud2 
  1382: 0f0b                     ud2 
  1384: 0f0b                     ud2 
  1386: 0f0b                     ud2 
  1388: 0f0b                     ud2 
  1390: 0f0b                     ud2 
  1392: 0f0b                     ud2 
  1394: 0f0b                     ud2 
  1396: 0f0b                     ud2 
  1398: 0f0b                     ud2 
  1400: 0f0b                     ud2 
  1402: 0f0b                     ud2 
  1404: 0f0b                     ud2 
  1406: 0f0b                     ud2 
