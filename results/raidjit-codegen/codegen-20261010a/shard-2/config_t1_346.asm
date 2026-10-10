; config_t1_346 target=avx2 twin=xorsched topo=rdp(7)/enc op=encode U=1
; code_size=2560 twins=["xorsched"] uops_per_byte=5.21875
     0: 53                       pushq %rbx
     1: 4154                     pushq %r12
     3: 4155                     pushq %r13
     5: 4156                     pushq %r14
     7: 4c8bc7                   movq %rdi, %r8
    10: 4c8bce                   movq %rsi, %r9
    13: 488bca                   movq %rdx, %rcx
    16: 498b30                   movq (%r8), %rsi
    19: 498b7808                 movq 8(%r8), %rdi
    23: 4d8b6010                 movq 0x10(%r8), %r12
    27: 4d8b6818                 movq 0x18(%r8), %r13
    31: 4d8b7020                 movq 0x20(%r8), %r14
    35: 4d33db                   xorq %r11, %r11
    38: 488bc1                   movq %rcx, %rax
    41: 4883e820                 subq $0x20, %rax
    45: 0f8206040000             jb 0x439
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 498b5038                 movq 0x38(%r8), %rdx
    68: c4217d6f241f             vmovdqa (%rdi, %r11), %ymm12
    74: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
    80: c4c17d6fcc               vmovdqa %ymm12, %ymm1
    85: c4c17d6ff4               vmovdqa %ymm12, %ymm6
    90: 498b5040                 movq 0x40(%r8), %rdx
    94: c4017d6f241c             vmovdqa (%r12, %r11), %ymm12
   100: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   106: c4c17d6fd4               vmovdqa %ymm12, %ymm2
   111: c4c17d6ffc               vmovdqa %ymm12, %ymm7
   116: 498b5048                 movq 0x48(%r8), %rdx
   120: c4017d6f641d00           vmovdqa (%r13, %r11), %ymm12
   127: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   133: c4c17d6fdc               vmovdqa %ymm12, %ymm3
   138: c4417d6fc4               vmovdqa %ymm12, %ymm8
   143: 498b5050                 movq 0x50(%r8), %rdx
   147: c4017d6f241e             vmovdqa (%r14, %r11), %ymm12
   153: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   159: c4c17d6fe4               vmovdqa %ymm12, %ymm4
   164: c4417d6fcc               vmovdqa %ymm12, %ymm9
   169: 4d8b5028                 movq 0x28(%r8), %r10
   173: 498b5058                 movq 0x58(%r8), %rdx
   177: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   183: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   189: c4c17d6fec               vmovdqa %ymm12, %ymm5
   194: c4417d6fd4               vmovdqa %ymm12, %ymm10
   199: 4d8b5068                 movq 0x68(%r8), %r10
   203: 498b9098000000           movq 0x98(%r8), %rdx
   210: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   216: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   222: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   227: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   232: 4d8b5070                 movq 0x70(%r8), %r10
   236: 498b90a0000000           movq 0xa0(%r8), %rdx
   243: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   249: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   255: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   260: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   265: 4d8b5078                 movq 0x78(%r8), %r10
   269: 498b90a8000000           movq 0xa8(%r8), %rdx
   276: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   282: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   288: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   293: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   298: 4d8b9080000000           movq 0x80(%r8), %r10
   305: 498b90b0000000           movq 0xb0(%r8), %rdx
   312: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   318: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   324: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   329: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   334: 4d8b9088000000           movq 0x88(%r8), %r10
   341: 498b90b8000000           movq 0xb8(%r8), %rdx
   348: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   354: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   360: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   365: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   370: 4d8b90c8000000           movq 0xc8(%r8), %r10
   377: 498b90f8000000           movq 0xf8(%r8), %rdx
   384: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   390: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   396: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   401: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   406: 4d8b90d0000000           movq 0xd0(%r8), %r10
   413: 498b9000010000           movq 0x100(%r8), %rdx
   420: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   426: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   432: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   437: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   442: 4d8b90d8000000           movq 0xd8(%r8), %r10
   449: 498b9008010000           movq 0x108(%r8), %rdx
   456: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   462: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   468: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   473: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   478: 4d8b90e0000000           movq 0xe0(%r8), %r10
   485: 498b9010010000           movq 0x110(%r8), %rdx
   492: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   498: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   504: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   509: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   514: 4d8b90e8000000           movq 0xe8(%r8), %r10
   521: 498b9018010000           movq 0x118(%r8), %rdx
   528: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   534: c4211def241a             vpxor (%rdx, %r11), %ymm12, %ymm12
   540: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   545: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   550: c4217d6f241e             vmovdqa (%rsi, %r11), %ymm12
   556: c4c17d6fc4               vmovdqa %ymm12, %ymm0
   561: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   566: c4a145ef3c1f             vpxor (%rdi, %r11), %ymm7, %ymm7
   572: c4013def041c             vpxor (%r12, %r11), %ymm8, %ymm8
   578: c40135ef4c1d00           vpxor (%r13, %r11), %ymm9, %ymm9
   585: c4012def141e             vpxor (%r14, %r11), %ymm10, %ymm10
   591: 4d8b5028                 movq 0x28(%r8), %r10
   595: c4017d6f1c1a             vmovdqa (%r10, %r11), %ymm11
   601: 4d8b5030                 movq 0x30(%r8), %r10
   605: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   611: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   616: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   621: 4d8b5038                 movq 0x38(%r8), %r10
   625: c4013def041a             vpxor (%r10, %r11), %ymm8, %ymm8
   631: 4d8b5040                 movq 0x40(%r8), %r10
   635: c40135ef0c1a             vpxor (%r10, %r11), %ymm9, %ymm9
   641: 4d8b5048                 movq 0x48(%r8), %r10
   645: c4012def141a             vpxor (%r10, %r11), %ymm10, %ymm10
   651: 4d8b5050                 movq 0x50(%r8), %r10
   655: c40125ef1c1a             vpxor (%r10, %r11), %ymm11, %ymm11
   661: 4d8b5060                 movq 0x60(%r8), %r10
   665: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   671: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   676: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   681: 4d8b5068                 movq 0x68(%r8), %r10
   685: c40135ef0c1a             vpxor (%r10, %r11), %ymm9, %ymm9
   691: 4d8b5070                 movq 0x70(%r8), %r10
   695: c4012def141a             vpxor (%r10, %r11), %ymm10, %ymm10
   701: 4d8b5078                 movq 0x78(%r8), %r10
   705: c40125ef1c1a             vpxor (%r10, %r11), %ymm11, %ymm11
   711: 4d8b9088000000           movq 0x88(%r8), %r10
   718: c4814def341a             vpxor (%r10, %r11), %ymm6, %ymm6
   724: 4d8b9090000000           movq 0x90(%r8), %r10
   731: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   737: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   742: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   747: 4d8b9098000000           movq 0x98(%r8), %r10
   754: c4012def141a             vpxor (%r10, %r11), %ymm10, %ymm10
   760: 4d8b90a0000000           movq 0xa0(%r8), %r10
   767: c40125ef1c1a             vpxor (%r10, %r11), %ymm11, %ymm11
   773: 4d8b90b0000000           movq 0xb0(%r8), %r10
   780: c4814def341a             vpxor (%r10, %r11), %ymm6, %ymm6
   786: 4d8b90b8000000           movq 0xb8(%r8), %r10
   793: c48145ef3c1a             vpxor (%r10, %r11), %ymm7, %ymm7
   799: 4d8b90c0000000           movq 0xc0(%r8), %r10
   806: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   812: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   817: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   822: 4d8b90c8000000           movq 0xc8(%r8), %r10
   829: c40125ef1c1a             vpxor (%r10, %r11), %ymm11, %ymm11
   835: 4d8b90d8000000           movq 0xd8(%r8), %r10
   842: c4814def341a             vpxor (%r10, %r11), %ymm6, %ymm6
   848: 4d8b90e0000000           movq 0xe0(%r8), %r10
   855: c48145ef3c1a             vpxor (%r10, %r11), %ymm7, %ymm7
   861: 4d8b90e8000000           movq 0xe8(%r8), %r10
   868: c4013def041a             vpxor (%r10, %r11), %ymm8, %ymm8
   874: 4d8b90f0000000           movq 0xf0(%r8), %r10
   881: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   887: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   892: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   897: 4d8b9000010000           movq 0x100(%r8), %r10
   904: c4814def341a             vpxor (%r10, %r11), %ymm6, %ymm6
   910: 4d8b9008010000           movq 0x108(%r8), %r10
   917: c48145ef3c1a             vpxor (%r10, %r11), %ymm7, %ymm7
   923: 4d8b9010010000           movq 0x110(%r8), %r10
   930: c4013def041a             vpxor (%r10, %r11), %ymm8, %ymm8
   936: 4d8b9018010000           movq 0x118(%r8), %r10
   943: c40135ef0c1a             vpxor (%r10, %r11), %ymm9, %ymm9
   949: 4d8b11                   movq (%r9), %r10
   952: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   958: 4d8b5108                 movq 8(%r9), %r10
   962: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
   968: 4d8b5110                 movq 0x10(%r9), %r10
   972: c4817d7f141a             vmovdqa %ymm2, (%r10, %r11)
   978: 4d8b5118                 movq 0x18(%r9), %r10
   982: c4817d7f1c1a             vmovdqa %ymm3, (%r10, %r11)
   988: 4d8b5120                 movq 0x20(%r9), %r10
   992: c4817d7f241a             vmovdqa %ymm4, (%r10, %r11)
   998: 4d8b5128                 movq 0x28(%r9), %r10
  1002: c4817d7f2c1a             vmovdqa %ymm5, (%r10, %r11)
  1008: 4d8b5130                 movq 0x30(%r9), %r10
  1012: c4817d7f341a             vmovdqa %ymm6, (%r10, %r11)
  1018: 4d8b5138                 movq 0x38(%r9), %r10
  1022: c4817d7f3c1a             vmovdqa %ymm7, (%r10, %r11)
  1028: 4d8b5140                 movq 0x40(%r9), %r10
  1032: c4017d7f041a             vmovdqa %ymm8, (%r10, %r11)
  1038: 4d8b5148                 movq 0x48(%r9), %r10
  1042: c4017d7f0c1a             vmovdqa %ymm9, (%r10, %r11)
  1048: 4d8b5150                 movq 0x50(%r9), %r10
  1052: c4017d7f141a             vmovdqa %ymm10, (%r10, %r11)
  1058: 4d8b5158                 movq 0x58(%r9), %r10
  1062: c4017d7f1c1a             vmovdqa %ymm11, (%r10, %r11)
  1068: 4983c320                 addq $0x20, %r11
  1072: 4c3bd8                   cmpq %rax, %r11
  1075: 0f8607fcffff             jbe 0x40
  1081: 4c3bd9                   cmpq %rcx, %r11
  1084: 0f83a4050000             jae 0x9e6
  1090: 4153                     pushq %r11
  1092: 4d8b11                   movq (%r9), %r10
  1095: 4c8b1c24                 movq (%rsp), %r11
  1099: 498b00                   movq (%r8), %rax
  1102: 428a1418                 movb (%rax, %r11), %dl
  1106: 8ada                     movb %dl, %bl
  1108: 498b4030                 movq 0x30(%r8), %rax
  1112: 428a1418                 movb (%rax, %r11), %dl
  1116: 32da                     xorb %dl, %bl
  1118: 498b4060                 movq 0x60(%r8), %rax
  1122: 428a1418                 movb (%rax, %r11), %dl
  1126: 32da                     xorb %dl, %bl
  1128: 498b8090000000           movq 0x90(%r8), %rax
  1135: 428a1418                 movb (%rax, %r11), %dl
  1139: 32da                     xorb %dl, %bl
  1141: 498b80c0000000           movq 0xc0(%r8), %rax
  1148: 428a1418                 movb (%rax, %r11), %dl
  1152: 32da                     xorb %dl, %bl
  1154: 498b80f0000000           movq 0xf0(%r8), %rax
  1161: 428a1418                 movb (%rax, %r11), %dl
  1165: 32da                     xorb %dl, %bl
  1167: 43881c1a                 movb %bl, (%r10, %r11)
  1171: 33db                     xorl %ebx, %ebx
  1173: 49ffc3                   incq %r11
  1176: 4c3bd9                   cmpq %rcx, %r11
  1179: 0f82aaffffff             jb 0x44b
  1185: 4d8b5108                 movq 8(%r9), %r10
  1189: 4c8b1c24                 movq (%rsp), %r11
  1193: 498b4008                 movq 8(%r8), %rax
  1197: 428a1418                 movb (%rax, %r11), %dl
  1201: 8ada                     movb %dl, %bl
  1203: 498b4038                 movq 0x38(%r8), %rax
  1207: 428a1418                 movb (%rax, %r11), %dl
  1211: 32da                     xorb %dl, %bl
  1213: 498b4068                 movq 0x68(%r8), %rax
  1217: 428a1418                 movb (%rax, %r11), %dl
  1221: 32da                     xorb %dl, %bl
  1223: 498b8098000000           movq 0x98(%r8), %rax
  1230: 428a1418                 movb (%rax, %r11), %dl
  1234: 32da                     xorb %dl, %bl
  1236: 498b80c8000000           movq 0xc8(%r8), %rax
  1243: 428a1418                 movb (%rax, %r11), %dl
  1247: 32da                     xorb %dl, %bl
  1249: 498b80f8000000           movq 0xf8(%r8), %rax
  1256: 428a1418                 movb (%rax, %r11), %dl
  1260: 32da                     xorb %dl, %bl
  1262: 43881c1a                 movb %bl, (%r10, %r11)
  1266: 33db                     xorl %ebx, %ebx
  1268: 49ffc3                   incq %r11
  1271: 4c3bd9                   cmpq %rcx, %r11
  1274: 0f82a9ffffff             jb 0x4a9
  1280: 4d8b5110                 movq 0x10(%r9), %r10
  1284: 4c8b1c24                 movq (%rsp), %r11
  1288: 498b4010                 movq 0x10(%r8), %rax
  1292: 428a1418                 movb (%rax, %r11), %dl
  1296: 8ada                     movb %dl, %bl
  1298: 498b4040                 movq 0x40(%r8), %rax
  1302: 428a1418                 movb (%rax, %r11), %dl
  1306: 32da                     xorb %dl, %bl
  1308: 498b4070                 movq 0x70(%r8), %rax
  1312: 428a1418                 movb (%rax, %r11), %dl
  1316: 32da                     xorb %dl, %bl
  1318: 498b80a0000000           movq 0xa0(%r8), %rax
  1325: 428a1418                 movb (%rax, %r11), %dl
  1329: 32da                     xorb %dl, %bl
  1331: 498b80d0000000           movq 0xd0(%r8), %rax
  1338: 428a1418                 movb (%rax, %r11), %dl
  1342: 32da                     xorb %dl, %bl
  1344: 498b8000010000           movq 0x100(%r8), %rax
  1351: 428a1418                 movb (%rax, %r11), %dl
  1355: 32da                     xorb %dl, %bl
  1357: 43881c1a                 movb %bl, (%r10, %r11)
  1361: 33db                     xorl %ebx, %ebx
  1363: 49ffc3                   incq %r11
  1366: 4c3bd9                   cmpq %rcx, %r11
  1369: 0f82a9ffffff             jb 0x508
  1375: 4d8b5118                 movq 0x18(%r9), %r10
  1379: 4c8b1c24                 movq (%rsp), %r11
  1383: 498b4018                 movq 0x18(%r8), %rax
  1387: 428a1418                 movb (%rax, %r11), %dl
  1391: 8ada                     movb %dl, %bl
  1393: 498b4048                 movq 0x48(%r8), %rax
  1397: 428a1418                 movb (%rax, %r11), %dl
  1401: 32da                     xorb %dl, %bl
  1403: 498b4078                 movq 0x78(%r8), %rax
  1407: 428a1418                 movb (%rax, %r11), %dl
  1411: 32da                     xorb %dl, %bl
  1413: 498b80a8000000           movq 0xa8(%r8), %rax
  1420: 428a1418                 movb (%rax, %r11), %dl
  1424: 32da                     xorb %dl, %bl
  1426: 498b80d8000000           movq 0xd8(%r8), %rax
  1433: 428a1418                 movb (%rax, %r11), %dl
  1437: 32da                     xorb %dl, %bl
  1439: 498b8008010000           movq 0x108(%r8), %rax
  1446: 428a1418                 movb (%rax, %r11), %dl
  1450: 32da                     xorb %dl, %bl
  1452: 43881c1a                 movb %bl, (%r10, %r11)
  1456: 33db                     xorl %ebx, %ebx
  1458: 49ffc3                   incq %r11
  1461: 4c3bd9                   cmpq %rcx, %r11
  1464: 0f82a9ffffff             jb 0x567
  1470: 4d8b5120                 movq 0x20(%r9), %r10
  1474: 4c8b1c24                 movq (%rsp), %r11
  1478: 498b4020                 movq 0x20(%r8), %rax
  1482: 428a1418                 movb (%rax, %r11), %dl
  1486: 8ada                     movb %dl, %bl
  1488: 498b4050                 movq 0x50(%r8), %rax
  1492: 428a1418                 movb (%rax, %r11), %dl
  1496: 32da                     xorb %dl, %bl
  1498: 498b8080000000           movq 0x80(%r8), %rax
  1505: 428a1418                 movb (%rax, %r11), %dl
  1509: 32da                     xorb %dl, %bl
  1511: 498b80b0000000           movq 0xb0(%r8), %rax
  1518: 428a1418                 movb (%rax, %r11), %dl
  1522: 32da                     xorb %dl, %bl
  1524: 498b80e0000000           movq 0xe0(%r8), %rax
  1531: 428a1418                 movb (%rax, %r11), %dl
  1535: 32da                     xorb %dl, %bl
  1537: 498b8010010000           movq 0x110(%r8), %rax
  1544: 428a1418                 movb (%rax, %r11), %dl
  1548: 32da                     xorb %dl, %bl
  1550: 43881c1a                 movb %bl, (%r10, %r11)
  1554: 33db                     xorl %ebx, %ebx
  1556: 49ffc3                   incq %r11
  1559: 4c3bd9                   cmpq %rcx, %r11
  1562: 0f82a6ffffff             jb 0x5c6
  1568: 4d8b5128                 movq 0x28(%r9), %r10
  1572: 4c8b1c24                 movq (%rsp), %r11
  1576: 498b4028                 movq 0x28(%r8), %rax
  1580: 428a1418                 movb (%rax, %r11), %dl
  1584: 8ada                     movb %dl, %bl
  1586: 498b4058                 movq 0x58(%r8), %rax
  1590: 428a1418                 movb (%rax, %r11), %dl
  1594: 32da                     xorb %dl, %bl
  1596: 498b8088000000           movq 0x88(%r8), %rax
  1603: 428a1418                 movb (%rax, %r11), %dl
  1607: 32da                     xorb %dl, %bl
  1609: 498b80b8000000           movq 0xb8(%r8), %rax
  1616: 428a1418                 movb (%rax, %r11), %dl
  1620: 32da                     xorb %dl, %bl
  1622: 498b80e8000000           movq 0xe8(%r8), %rax
  1629: 428a1418                 movb (%rax, %r11), %dl
  1633: 32da                     xorb %dl, %bl
  1635: 498b8018010000           movq 0x118(%r8), %rax
  1642: 428a1418                 movb (%rax, %r11), %dl
  1646: 32da                     xorb %dl, %bl
  1648: 43881c1a                 movb %bl, (%r10, %r11)
  1652: 33db                     xorl %ebx, %ebx
  1654: 49ffc3                   incq %r11
  1657: 4c3bd9                   cmpq %rcx, %r11
  1660: 0f82a6ffffff             jb 0x628
  1666: 4d8b5130                 movq 0x30(%r9), %r10
  1670: 4c8b1c24                 movq (%rsp), %r11
  1674: 498b00                   movq (%r8), %rax
  1677: 428a1418                 movb (%rax, %r11), %dl
  1681: 8ada                     movb %dl, %bl
  1683: 498b4008                 movq 8(%r8), %rax
  1687: 428a1418                 movb (%rax, %r11), %dl
  1691: 32da                     xorb %dl, %bl
  1693: 498b4038                 movq 0x38(%r8), %rax
  1697: 428a1418                 movb (%rax, %r11), %dl
  1701: 32da                     xorb %dl, %bl
  1703: 498b4068                 movq 0x68(%r8), %rax
  1707: 428a1418                 movb (%rax, %r11), %dl
  1711: 32da                     xorb %dl, %bl
  1713: 498b8088000000           movq 0x88(%r8), %rax
  1720: 428a1418                 movb (%rax, %r11), %dl
  1724: 32da                     xorb %dl, %bl
  1726: 498b8098000000           movq 0x98(%r8), %rax
  1733: 428a1418                 movb (%rax, %r11), %dl
  1737: 32da                     xorb %dl, %bl
  1739: 498b80b0000000           movq 0xb0(%r8), %rax
  1746: 428a1418                 movb (%rax, %r11), %dl
  1750: 32da                     xorb %dl, %bl
  1752: 498b80c8000000           movq 0xc8(%r8), %rax
  1759: 428a1418                 movb (%rax, %r11), %dl
  1763: 32da                     xorb %dl, %bl
  1765: 498b80d8000000           movq 0xd8(%r8), %rax
  1772: 428a1418                 movb (%rax, %r11), %dl
  1776: 32da                     xorb %dl, %bl
  1778: 498b80f8000000           movq 0xf8(%r8), %rax
  1785: 428a1418                 movb (%rax, %r11), %dl
  1789: 32da                     xorb %dl, %bl
  1791: 498b8000010000           movq 0x100(%r8), %rax
  1798: 428a1418                 movb (%rax, %r11), %dl
  1802: 32da                     xorb %dl, %bl
  1804: 43881c1a                 movb %bl, (%r10, %r11)
  1808: 33db                     xorl %ebx, %ebx
  1810: 49ffc3                   incq %r11
  1813: 4c3bd9                   cmpq %rcx, %r11
  1816: 0f826cffffff             jb 0x68a
  1822: 4d8b5138                 movq 0x38(%r9), %r10
  1826: 4c8b1c24                 movq (%rsp), %r11
  1830: 498b4008                 movq 8(%r8), %rax
  1834: 428a1418                 movb (%rax, %r11), %dl
  1838: 8ada                     movb %dl, %bl
  1840: 498b4010                 movq 0x10(%r8), %rax
  1844: 428a1418                 movb (%rax, %r11), %dl
  1848: 32da                     xorb %dl, %bl
  1850: 498b4030                 movq 0x30(%r8), %rax
  1854: 428a1418                 movb (%rax, %r11), %dl
  1858: 32da                     xorb %dl, %bl
  1860: 498b4040                 movq 0x40(%r8), %rax
  1864: 428a1418                 movb (%rax, %r11), %dl
  1868: 32da                     xorb %dl, %bl
  1870: 498b4070                 movq 0x70(%r8), %rax
  1874: 428a1418                 movb (%rax, %r11), %dl
  1878: 32da                     xorb %dl, %bl
  1880: 498b80a0000000           movq 0xa0(%r8), %rax
  1887: 428a1418                 movb (%rax, %r11), %dl
  1891: 32da                     xorb %dl, %bl
  1893: 498b80b8000000           movq 0xb8(%r8), %rax
  1900: 428a1418                 movb (%rax, %r11), %dl
  1904: 32da                     xorb %dl, %bl
  1906: 498b80d0000000           movq 0xd0(%r8), %rax
  1913: 428a1418                 movb (%rax, %r11), %dl
  1917: 32da                     xorb %dl, %bl
  1919: 498b80e0000000           movq 0xe0(%r8), %rax
  1926: 428a1418                 movb (%rax, %r11), %dl
  1930: 32da                     xorb %dl, %bl
  1932: 498b8000010000           movq 0x100(%r8), %rax
  1939: 428a1418                 movb (%rax, %r11), %dl
  1943: 32da                     xorb %dl, %bl
  1945: 498b8008010000           movq 0x108(%r8), %rax
  1952: 428a1418                 movb (%rax, %r11), %dl
  1956: 32da                     xorb %dl, %bl
  1958: 43881c1a                 movb %bl, (%r10, %r11)
  1962: 33db                     xorl %ebx, %ebx
  1964: 49ffc3                   incq %r11
  1967: 4c3bd9                   cmpq %rcx, %r11
  1970: 0f826effffff             jb 0x726
  1976: 4d8b5140                 movq 0x40(%r9), %r10
  1980: 4c8b1c24                 movq (%rsp), %r11
  1984: 498b4010                 movq 0x10(%r8), %rax
  1988: 428a1418                 movb (%rax, %r11), %dl
  1992: 8ada                     movb %dl, %bl
  1994: 498b4018                 movq 0x18(%r8), %rax
  1998: 428a1418                 movb (%rax, %r11), %dl
  2002: 32da                     xorb %dl, %bl
  2004: 498b4038                 movq 0x38(%r8), %rax
  2008: 428a1418                 movb (%rax, %r11), %dl
  2012: 32da                     xorb %dl, %bl
  2014: 498b4048                 movq 0x48(%r8), %rax
  2018: 428a1418                 movb (%rax, %r11), %dl
  2022: 32da                     xorb %dl, %bl
  2024: 498b4060                 movq 0x60(%r8), %rax
  2028: 428a1418                 movb (%rax, %r11), %dl
  2032: 32da                     xorb %dl, %bl
  2034: 498b4078                 movq 0x78(%r8), %rax
  2038: 428a1418                 movb (%rax, %r11), %dl
  2042: 32da                     xorb %dl, %bl
  2044: 498b80a8000000           movq 0xa8(%r8), %rax
  2051: 428a1418                 movb (%rax, %r11), %dl
  2055: 32da                     xorb %dl, %bl
  2057: 498b80d8000000           movq 0xd8(%r8), %rax
  2064: 428a1418                 movb (%rax, %r11), %dl
  2068: 32da                     xorb %dl, %bl
  2070: 498b80e8000000           movq 0xe8(%r8), %rax
  2077: 428a1418                 movb (%rax, %r11), %dl
  2081: 32da                     xorb %dl, %bl
  2083: 498b8008010000           movq 0x108(%r8), %rax
  2090: 428a1418                 movb (%rax, %r11), %dl
  2094: 32da                     xorb %dl, %bl
  2096: 498b8010010000           movq 0x110(%r8), %rax
  2103: 428a1418                 movb (%rax, %r11), %dl
  2107: 32da                     xorb %dl, %bl
  2109: 43881c1a                 movb %bl, (%r10, %r11)
  2113: 33db                     xorl %ebx, %ebx
  2115: 49ffc3                   incq %r11
  2118: 4c3bd9                   cmpq %rcx, %r11
  2121: 0f8271ffffff             jb 0x7c0
  2127: 4d8b5148                 movq 0x48(%r9), %r10
  2131: 4c8b1c24                 movq (%rsp), %r11
  2135: 498b4018                 movq 0x18(%r8), %rax
  2139: 428a1418                 movb (%rax, %r11), %dl
  2143: 8ada                     movb %dl, %bl
  2145: 498b4020                 movq 0x20(%r8), %rax
  2149: 428a1418                 movb (%rax, %r11), %dl
  2153: 32da                     xorb %dl, %bl
  2155: 498b4040                 movq 0x40(%r8), %rax
  2159: 428a1418                 movb (%rax, %r11), %dl
  2163: 32da                     xorb %dl, %bl
  2165: 498b4050                 movq 0x50(%r8), %rax
  2169: 428a1418                 movb (%rax, %r11), %dl
  2173: 32da                     xorb %dl, %bl
  2175: 498b4068                 movq 0x68(%r8), %rax
  2179: 428a1418                 movb (%rax, %r11), %dl
  2183: 32da                     xorb %dl, %bl
  2185: 498b8080000000           movq 0x80(%r8), %rax
  2192: 428a1418                 movb (%rax, %r11), %dl
  2196: 32da                     xorb %dl, %bl
  2198: 498b8090000000           movq 0x90(%r8), %rax
  2205: 428a1418                 movb (%rax, %r11), %dl
  2209: 32da                     xorb %dl, %bl
  2211: 498b80b0000000           movq 0xb0(%r8), %rax
  2218: 428a1418                 movb (%rax, %r11), %dl
  2222: 32da                     xorb %dl, %bl
  2224: 498b80e0000000           movq 0xe0(%r8), %rax
  2231: 428a1418                 movb (%rax, %r11), %dl
  2235: 32da                     xorb %dl, %bl
  2237: 498b8010010000           movq 0x110(%r8), %rax
  2244: 428a1418                 movb (%rax, %r11), %dl
  2248: 32da                     xorb %dl, %bl
  2250: 498b8018010000           movq 0x118(%r8), %rax
  2257: 428a1418                 movb (%rax, %r11), %dl
  2261: 32da                     xorb %dl, %bl
  2263: 43881c1a                 movb %bl, (%r10, %r11)
  2267: 33db                     xorl %ebx, %ebx
  2269: 49ffc3                   incq %r11
  2272: 4c3bd9                   cmpq %rcx, %r11
  2275: 0f826effffff             jb 0x857
  2281: 4d8b5150                 movq 0x50(%r9), %r10
  2285: 4c8b1c24                 movq (%rsp), %r11
  2289: 498b4020                 movq 0x20(%r8), %rax
  2293: 428a1418                 movb (%rax, %r11), %dl
  2297: 8ada                     movb %dl, %bl
  2299: 498b4028                 movq 0x28(%r8), %rax
  2303: 428a1418                 movb (%rax, %r11), %dl
  2307: 32da                     xorb %dl, %bl
  2309: 498b4048                 movq 0x48(%r8), %rax
  2313: 428a1418                 movb (%rax, %r11), %dl
  2317: 32da                     xorb %dl, %bl
  2319: 498b4058                 movq 0x58(%r8), %rax
  2323: 428a1418                 movb (%rax, %r11), %dl
  2327: 32da                     xorb %dl, %bl
  2329: 498b4070                 movq 0x70(%r8), %rax
  2333: 428a1418                 movb (%rax, %r11), %dl
  2337: 32da                     xorb %dl, %bl
  2339: 498b8088000000           movq 0x88(%r8), %rax
  2346: 428a1418                 movb (%rax, %r11), %dl
  2350: 32da                     xorb %dl, %bl
  2352: 498b8098000000           movq 0x98(%r8), %rax
  2359: 428a1418                 movb (%rax, %r11), %dl
  2363: 32da                     xorb %dl, %bl
  2365: 498b80b8000000           movq 0xb8(%r8), %rax
  2372: 428a1418                 movb (%rax, %r11), %dl
  2376: 32da                     xorb %dl, %bl
  2378: 498b80c0000000           movq 0xc0(%r8), %rax
  2385: 428a1418                 movb (%rax, %r11), %dl
  2389: 32da                     xorb %dl, %bl
  2391: 498b80e8000000           movq 0xe8(%r8), %rax
  2398: 428a1418                 movb (%rax, %r11), %dl
  2402: 32da                     xorb %dl, %bl
  2404: 498b8018010000           movq 0x118(%r8), %rax
  2411: 428a1418                 movb (%rax, %r11), %dl
  2415: 32da                     xorb %dl, %bl
  2417: 43881c1a                 movb %bl, (%r10, %r11)
  2421: 33db                     xorl %ebx, %ebx
  2423: 49ffc3                   incq %r11
  2426: 4c3bd9                   cmpq %rcx, %r11
  2429: 0f826effffff             jb 0x8f1
  2435: 4d8b5158                 movq 0x58(%r9), %r10
  2439: 4c8b1c24                 movq (%rsp), %r11
  2443: 498b4028                 movq 0x28(%r8), %rax
  2447: 428a1418                 movb (%rax, %r11), %dl
  2451: 8ada                     movb %dl, %bl
  2453: 498b4050                 movq 0x50(%r8), %rax
  2457: 428a1418                 movb (%rax, %r11), %dl
  2461: 32da                     xorb %dl, %bl
  2463: 498b4078                 movq 0x78(%r8), %rax
  2467: 428a1418                 movb (%rax, %r11), %dl
  2471: 32da                     xorb %dl, %bl
  2473: 498b80a0000000           movq 0xa0(%r8), %rax
  2480: 428a1418                 movb (%rax, %r11), %dl
  2484: 32da                     xorb %dl, %bl
  2486: 498b80c8000000           movq 0xc8(%r8), %rax
  2493: 428a1418                 movb (%rax, %r11), %dl
  2497: 32da                     xorb %dl, %bl
  2499: 498b80f0000000           movq 0xf0(%r8), %rax
  2506: 428a1418                 movb (%rax, %r11), %dl
  2510: 32da                     xorb %dl, %bl
  2512: 43881c1a                 movb %bl, (%r10, %r11)
  2516: 33db                     xorl %ebx, %ebx
  2518: 49ffc3                   incq %r11
  2521: 4c3bd9                   cmpq %rcx, %r11
  2524: 0f82a9ffffff             jb 0x98b
  2530: 4883c408                 addq $8, %rsp
  2534: 415e                     popq %r14
  2536: 415d                     popq %r13
  2538: 415c                     popq %r12
  2540: 5b                       popq %rbx
  2541: c5f877                   vzeroupper 
  2544: c3                       retq 
  2545: 0f0b                     ud2 
  2547: 0f0b                     ud2 
  2549: 0f0b                     ud2 
  2551: 0f0b                     ud2 
  2553: 0f0b                     ud2 
  2555: 0f0b                     ud2 
  2557: 0f0b                     ud2 
  2559: cc                       int3 
