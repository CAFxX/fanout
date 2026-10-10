; config_t1_339 target=avx2 twin=nt topo=rdp(7)/enc op=encode U=1
; code_size=2560 twins=["nt"] uops_per_byte=5.5
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
    45: 0f8207040000             jb 0x43a
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: c4217d6f241e             vmovdqa (%rsi, %r11), %ymm12
    70: c4c17d6fc4               vmovdqa %ymm12, %ymm0
    75: c4c17d6ff4               vmovdqa %ymm12, %ymm6
    80: c4217d6f241f             vmovdqa (%rdi, %r11), %ymm12
    86: c4c17d6fcc               vmovdqa %ymm12, %ymm1
    91: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
    96: c4c17d6ffc               vmovdqa %ymm12, %ymm7
   101: c4017d6f241c             vmovdqa (%r12, %r11), %ymm12
   107: c4c17d6fd4               vmovdqa %ymm12, %ymm2
   112: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   117: c4417d6fc4               vmovdqa %ymm12, %ymm8
   122: c4017d6f641d00           vmovdqa (%r13, %r11), %ymm12
   129: c4c17d6fdc               vmovdqa %ymm12, %ymm3
   134: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   139: c4417d6fcc               vmovdqa %ymm12, %ymm9
   144: c4017d6f241e             vmovdqa (%r14, %r11), %ymm12
   150: c4c17d6fe4               vmovdqa %ymm12, %ymm4
   155: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   160: c4417d6fd4               vmovdqa %ymm12, %ymm10
   165: 4d8b5028                 movq 0x28(%r8), %r10
   169: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   175: c4c17d6fec               vmovdqa %ymm12, %ymm5
   180: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   185: c4417d6fdc               vmovdqa %ymm12, %ymm11
   190: 4d8b5030                 movq 0x30(%r8), %r10
   194: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   200: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   205: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   210: 4d8b5038                 movq 0x38(%r8), %r10
   214: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   220: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   225: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   230: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   235: 4d8b5040                 movq 0x40(%r8), %r10
   239: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   245: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   250: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   255: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   260: 4d8b5048                 movq 0x48(%r8), %r10
   264: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   270: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   275: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   280: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   285: 4d8b5050                 movq 0x50(%r8), %r10
   289: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   295: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   300: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   305: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   310: 4d8b5058                 movq 0x58(%r8), %r10
   314: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   320: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   325: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   330: 4d8b5060                 movq 0x60(%r8), %r10
   334: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   340: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   345: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   350: 4d8b5068                 movq 0x68(%r8), %r10
   354: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   360: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   365: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   370: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   375: 4d8b5070                 movq 0x70(%r8), %r10
   379: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   385: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   390: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   395: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   400: 4d8b5078                 movq 0x78(%r8), %r10
   404: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   410: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   415: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   420: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   425: 4d8b9080000000           movq 0x80(%r8), %r10
   432: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   438: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   443: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   448: 4d8b9088000000           movq 0x88(%r8), %r10
   455: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   461: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   466: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   471: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   476: 4d8b9090000000           movq 0x90(%r8), %r10
   483: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   489: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   494: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   499: 4d8b9098000000           movq 0x98(%r8), %r10
   506: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   512: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   517: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   522: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   527: 4d8b90a0000000           movq 0xa0(%r8), %r10
   534: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   540: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   545: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   550: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   555: 4d8b90a8000000           movq 0xa8(%r8), %r10
   562: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   568: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   573: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   578: 4d8b90b0000000           movq 0xb0(%r8), %r10
   585: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   591: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   596: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   601: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   606: 4d8b90b8000000           movq 0xb8(%r8), %r10
   613: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   619: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   624: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   629: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   634: 4d8b90c0000000           movq 0xc0(%r8), %r10
   641: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   647: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   652: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   657: 4d8b90c8000000           movq 0xc8(%r8), %r10
   664: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   670: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   675: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   680: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   685: 4d8b90d0000000           movq 0xd0(%r8), %r10
   692: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   698: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   703: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   708: 4d8b90d8000000           movq 0xd8(%r8), %r10
   715: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   721: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   726: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   731: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   736: 4d8b90e0000000           movq 0xe0(%r8), %r10
   743: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   749: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   754: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   759: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   764: 4d8b90e8000000           movq 0xe8(%r8), %r10
   771: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   777: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   782: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   787: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   792: 4d8b90f0000000           movq 0xf0(%r8), %r10
   799: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   805: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   810: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   815: 4d8b90f8000000           movq 0xf8(%r8), %r10
   822: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   828: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   833: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   838: 4d8b9000010000           movq 0x100(%r8), %r10
   845: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   851: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   856: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   861: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   866: 4d8b9008010000           movq 0x108(%r8), %r10
   873: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   879: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   884: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   889: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   894: 4d8b9010010000           movq 0x110(%r8), %r10
   901: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   907: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   912: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   917: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   922: 4d8b9018010000           movq 0x118(%r8), %r10
   929: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   935: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   940: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   945: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   950: 4d8b11                   movq (%r9), %r10
   953: c4817de7041a             vmovntdq %ymm0, (%r10, %r11)
   959: 4d8b5108                 movq 8(%r9), %r10
   963: c4817de70c1a             vmovntdq %ymm1, (%r10, %r11)
   969: 4d8b5110                 movq 0x10(%r9), %r10
   973: c4817de7141a             vmovntdq %ymm2, (%r10, %r11)
   979: 4d8b5118                 movq 0x18(%r9), %r10
   983: c4817de71c1a             vmovntdq %ymm3, (%r10, %r11)
   989: 4d8b5120                 movq 0x20(%r9), %r10
   993: c4817de7241a             vmovntdq %ymm4, (%r10, %r11)
   999: 4d8b5128                 movq 0x28(%r9), %r10
  1003: c4817de72c1a             vmovntdq %ymm5, (%r10, %r11)
  1009: 4d8b5130                 movq 0x30(%r9), %r10
  1013: c4817de7341a             vmovntdq %ymm6, (%r10, %r11)
  1019: 4d8b5138                 movq 0x38(%r9), %r10
  1023: c4817de73c1a             vmovntdq %ymm7, (%r10, %r11)
  1029: 4d8b5140                 movq 0x40(%r9), %r10
  1033: c4017de7041a             vmovntdq %ymm8, (%r10, %r11)
  1039: 4d8b5148                 movq 0x48(%r9), %r10
  1043: c4017de70c1a             vmovntdq %ymm9, (%r10, %r11)
  1049: 4d8b5150                 movq 0x50(%r9), %r10
  1053: c4017de7141a             vmovntdq %ymm10, (%r10, %r11)
  1059: 4d8b5158                 movq 0x58(%r9), %r10
  1063: c4017de71c1a             vmovntdq %ymm11, (%r10, %r11)
  1069: 4983c320                 addq $0x20, %r11
  1073: 4c3bd8                   cmpq %rax, %r11
  1076: 0f8606fcffff             jbe 0x40
  1082: 4c3bd9                   cmpq %rcx, %r11
  1085: 0f83a4050000             jae 0x9e7
  1091: 4153                     pushq %r11
  1093: 4d8b11                   movq (%r9), %r10
  1096: 4c8b1c24                 movq (%rsp), %r11
  1100: 498b00                   movq (%r8), %rax
  1103: 428a1418                 movb (%rax, %r11), %dl
  1107: 8ada                     movb %dl, %bl
  1109: 498b4030                 movq 0x30(%r8), %rax
  1113: 428a1418                 movb (%rax, %r11), %dl
  1117: 32da                     xorb %dl, %bl
  1119: 498b4060                 movq 0x60(%r8), %rax
  1123: 428a1418                 movb (%rax, %r11), %dl
  1127: 32da                     xorb %dl, %bl
  1129: 498b8090000000           movq 0x90(%r8), %rax
  1136: 428a1418                 movb (%rax, %r11), %dl
  1140: 32da                     xorb %dl, %bl
  1142: 498b80c0000000           movq 0xc0(%r8), %rax
  1149: 428a1418                 movb (%rax, %r11), %dl
  1153: 32da                     xorb %dl, %bl
  1155: 498b80f0000000           movq 0xf0(%r8), %rax
  1162: 428a1418                 movb (%rax, %r11), %dl
  1166: 32da                     xorb %dl, %bl
  1168: 43881c1a                 movb %bl, (%r10, %r11)
  1172: 33db                     xorl %ebx, %ebx
  1174: 49ffc3                   incq %r11
  1177: 4c3bd9                   cmpq %rcx, %r11
  1180: 0f82aaffffff             jb 0x44c
  1186: 4d8b5108                 movq 8(%r9), %r10
  1190: 4c8b1c24                 movq (%rsp), %r11
  1194: 498b4008                 movq 8(%r8), %rax
  1198: 428a1418                 movb (%rax, %r11), %dl
  1202: 8ada                     movb %dl, %bl
  1204: 498b4038                 movq 0x38(%r8), %rax
  1208: 428a1418                 movb (%rax, %r11), %dl
  1212: 32da                     xorb %dl, %bl
  1214: 498b4068                 movq 0x68(%r8), %rax
  1218: 428a1418                 movb (%rax, %r11), %dl
  1222: 32da                     xorb %dl, %bl
  1224: 498b8098000000           movq 0x98(%r8), %rax
  1231: 428a1418                 movb (%rax, %r11), %dl
  1235: 32da                     xorb %dl, %bl
  1237: 498b80c8000000           movq 0xc8(%r8), %rax
  1244: 428a1418                 movb (%rax, %r11), %dl
  1248: 32da                     xorb %dl, %bl
  1250: 498b80f8000000           movq 0xf8(%r8), %rax
  1257: 428a1418                 movb (%rax, %r11), %dl
  1261: 32da                     xorb %dl, %bl
  1263: 43881c1a                 movb %bl, (%r10, %r11)
  1267: 33db                     xorl %ebx, %ebx
  1269: 49ffc3                   incq %r11
  1272: 4c3bd9                   cmpq %rcx, %r11
  1275: 0f82a9ffffff             jb 0x4aa
  1281: 4d8b5110                 movq 0x10(%r9), %r10
  1285: 4c8b1c24                 movq (%rsp), %r11
  1289: 498b4010                 movq 0x10(%r8), %rax
  1293: 428a1418                 movb (%rax, %r11), %dl
  1297: 8ada                     movb %dl, %bl
  1299: 498b4040                 movq 0x40(%r8), %rax
  1303: 428a1418                 movb (%rax, %r11), %dl
  1307: 32da                     xorb %dl, %bl
  1309: 498b4070                 movq 0x70(%r8), %rax
  1313: 428a1418                 movb (%rax, %r11), %dl
  1317: 32da                     xorb %dl, %bl
  1319: 498b80a0000000           movq 0xa0(%r8), %rax
  1326: 428a1418                 movb (%rax, %r11), %dl
  1330: 32da                     xorb %dl, %bl
  1332: 498b80d0000000           movq 0xd0(%r8), %rax
  1339: 428a1418                 movb (%rax, %r11), %dl
  1343: 32da                     xorb %dl, %bl
  1345: 498b8000010000           movq 0x100(%r8), %rax
  1352: 428a1418                 movb (%rax, %r11), %dl
  1356: 32da                     xorb %dl, %bl
  1358: 43881c1a                 movb %bl, (%r10, %r11)
  1362: 33db                     xorl %ebx, %ebx
  1364: 49ffc3                   incq %r11
  1367: 4c3bd9                   cmpq %rcx, %r11
  1370: 0f82a9ffffff             jb 0x509
  1376: 4d8b5118                 movq 0x18(%r9), %r10
  1380: 4c8b1c24                 movq (%rsp), %r11
  1384: 498b4018                 movq 0x18(%r8), %rax
  1388: 428a1418                 movb (%rax, %r11), %dl
  1392: 8ada                     movb %dl, %bl
  1394: 498b4048                 movq 0x48(%r8), %rax
  1398: 428a1418                 movb (%rax, %r11), %dl
  1402: 32da                     xorb %dl, %bl
  1404: 498b4078                 movq 0x78(%r8), %rax
  1408: 428a1418                 movb (%rax, %r11), %dl
  1412: 32da                     xorb %dl, %bl
  1414: 498b80a8000000           movq 0xa8(%r8), %rax
  1421: 428a1418                 movb (%rax, %r11), %dl
  1425: 32da                     xorb %dl, %bl
  1427: 498b80d8000000           movq 0xd8(%r8), %rax
  1434: 428a1418                 movb (%rax, %r11), %dl
  1438: 32da                     xorb %dl, %bl
  1440: 498b8008010000           movq 0x108(%r8), %rax
  1447: 428a1418                 movb (%rax, %r11), %dl
  1451: 32da                     xorb %dl, %bl
  1453: 43881c1a                 movb %bl, (%r10, %r11)
  1457: 33db                     xorl %ebx, %ebx
  1459: 49ffc3                   incq %r11
  1462: 4c3bd9                   cmpq %rcx, %r11
  1465: 0f82a9ffffff             jb 0x568
  1471: 4d8b5120                 movq 0x20(%r9), %r10
  1475: 4c8b1c24                 movq (%rsp), %r11
  1479: 498b4020                 movq 0x20(%r8), %rax
  1483: 428a1418                 movb (%rax, %r11), %dl
  1487: 8ada                     movb %dl, %bl
  1489: 498b4050                 movq 0x50(%r8), %rax
  1493: 428a1418                 movb (%rax, %r11), %dl
  1497: 32da                     xorb %dl, %bl
  1499: 498b8080000000           movq 0x80(%r8), %rax
  1506: 428a1418                 movb (%rax, %r11), %dl
  1510: 32da                     xorb %dl, %bl
  1512: 498b80b0000000           movq 0xb0(%r8), %rax
  1519: 428a1418                 movb (%rax, %r11), %dl
  1523: 32da                     xorb %dl, %bl
  1525: 498b80e0000000           movq 0xe0(%r8), %rax
  1532: 428a1418                 movb (%rax, %r11), %dl
  1536: 32da                     xorb %dl, %bl
  1538: 498b8010010000           movq 0x110(%r8), %rax
  1545: 428a1418                 movb (%rax, %r11), %dl
  1549: 32da                     xorb %dl, %bl
  1551: 43881c1a                 movb %bl, (%r10, %r11)
  1555: 33db                     xorl %ebx, %ebx
  1557: 49ffc3                   incq %r11
  1560: 4c3bd9                   cmpq %rcx, %r11
  1563: 0f82a6ffffff             jb 0x5c7
  1569: 4d8b5128                 movq 0x28(%r9), %r10
  1573: 4c8b1c24                 movq (%rsp), %r11
  1577: 498b4028                 movq 0x28(%r8), %rax
  1581: 428a1418                 movb (%rax, %r11), %dl
  1585: 8ada                     movb %dl, %bl
  1587: 498b4058                 movq 0x58(%r8), %rax
  1591: 428a1418                 movb (%rax, %r11), %dl
  1595: 32da                     xorb %dl, %bl
  1597: 498b8088000000           movq 0x88(%r8), %rax
  1604: 428a1418                 movb (%rax, %r11), %dl
  1608: 32da                     xorb %dl, %bl
  1610: 498b80b8000000           movq 0xb8(%r8), %rax
  1617: 428a1418                 movb (%rax, %r11), %dl
  1621: 32da                     xorb %dl, %bl
  1623: 498b80e8000000           movq 0xe8(%r8), %rax
  1630: 428a1418                 movb (%rax, %r11), %dl
  1634: 32da                     xorb %dl, %bl
  1636: 498b8018010000           movq 0x118(%r8), %rax
  1643: 428a1418                 movb (%rax, %r11), %dl
  1647: 32da                     xorb %dl, %bl
  1649: 43881c1a                 movb %bl, (%r10, %r11)
  1653: 33db                     xorl %ebx, %ebx
  1655: 49ffc3                   incq %r11
  1658: 4c3bd9                   cmpq %rcx, %r11
  1661: 0f82a6ffffff             jb 0x629
  1667: 4d8b5130                 movq 0x30(%r9), %r10
  1671: 4c8b1c24                 movq (%rsp), %r11
  1675: 498b00                   movq (%r8), %rax
  1678: 428a1418                 movb (%rax, %r11), %dl
  1682: 8ada                     movb %dl, %bl
  1684: 498b4008                 movq 8(%r8), %rax
  1688: 428a1418                 movb (%rax, %r11), %dl
  1692: 32da                     xorb %dl, %bl
  1694: 498b4038                 movq 0x38(%r8), %rax
  1698: 428a1418                 movb (%rax, %r11), %dl
  1702: 32da                     xorb %dl, %bl
  1704: 498b4068                 movq 0x68(%r8), %rax
  1708: 428a1418                 movb (%rax, %r11), %dl
  1712: 32da                     xorb %dl, %bl
  1714: 498b8088000000           movq 0x88(%r8), %rax
  1721: 428a1418                 movb (%rax, %r11), %dl
  1725: 32da                     xorb %dl, %bl
  1727: 498b8098000000           movq 0x98(%r8), %rax
  1734: 428a1418                 movb (%rax, %r11), %dl
  1738: 32da                     xorb %dl, %bl
  1740: 498b80b0000000           movq 0xb0(%r8), %rax
  1747: 428a1418                 movb (%rax, %r11), %dl
  1751: 32da                     xorb %dl, %bl
  1753: 498b80c8000000           movq 0xc8(%r8), %rax
  1760: 428a1418                 movb (%rax, %r11), %dl
  1764: 32da                     xorb %dl, %bl
  1766: 498b80d8000000           movq 0xd8(%r8), %rax
  1773: 428a1418                 movb (%rax, %r11), %dl
  1777: 32da                     xorb %dl, %bl
  1779: 498b80f8000000           movq 0xf8(%r8), %rax
  1786: 428a1418                 movb (%rax, %r11), %dl
  1790: 32da                     xorb %dl, %bl
  1792: 498b8000010000           movq 0x100(%r8), %rax
  1799: 428a1418                 movb (%rax, %r11), %dl
  1803: 32da                     xorb %dl, %bl
  1805: 43881c1a                 movb %bl, (%r10, %r11)
  1809: 33db                     xorl %ebx, %ebx
  1811: 49ffc3                   incq %r11
  1814: 4c3bd9                   cmpq %rcx, %r11
  1817: 0f826cffffff             jb 0x68b
  1823: 4d8b5138                 movq 0x38(%r9), %r10
  1827: 4c8b1c24                 movq (%rsp), %r11
  1831: 498b4008                 movq 8(%r8), %rax
  1835: 428a1418                 movb (%rax, %r11), %dl
  1839: 8ada                     movb %dl, %bl
  1841: 498b4010                 movq 0x10(%r8), %rax
  1845: 428a1418                 movb (%rax, %r11), %dl
  1849: 32da                     xorb %dl, %bl
  1851: 498b4030                 movq 0x30(%r8), %rax
  1855: 428a1418                 movb (%rax, %r11), %dl
  1859: 32da                     xorb %dl, %bl
  1861: 498b4040                 movq 0x40(%r8), %rax
  1865: 428a1418                 movb (%rax, %r11), %dl
  1869: 32da                     xorb %dl, %bl
  1871: 498b4070                 movq 0x70(%r8), %rax
  1875: 428a1418                 movb (%rax, %r11), %dl
  1879: 32da                     xorb %dl, %bl
  1881: 498b80a0000000           movq 0xa0(%r8), %rax
  1888: 428a1418                 movb (%rax, %r11), %dl
  1892: 32da                     xorb %dl, %bl
  1894: 498b80b8000000           movq 0xb8(%r8), %rax
  1901: 428a1418                 movb (%rax, %r11), %dl
  1905: 32da                     xorb %dl, %bl
  1907: 498b80d0000000           movq 0xd0(%r8), %rax
  1914: 428a1418                 movb (%rax, %r11), %dl
  1918: 32da                     xorb %dl, %bl
  1920: 498b80e0000000           movq 0xe0(%r8), %rax
  1927: 428a1418                 movb (%rax, %r11), %dl
  1931: 32da                     xorb %dl, %bl
  1933: 498b8000010000           movq 0x100(%r8), %rax
  1940: 428a1418                 movb (%rax, %r11), %dl
  1944: 32da                     xorb %dl, %bl
  1946: 498b8008010000           movq 0x108(%r8), %rax
  1953: 428a1418                 movb (%rax, %r11), %dl
  1957: 32da                     xorb %dl, %bl
  1959: 43881c1a                 movb %bl, (%r10, %r11)
  1963: 33db                     xorl %ebx, %ebx
  1965: 49ffc3                   incq %r11
  1968: 4c3bd9                   cmpq %rcx, %r11
  1971: 0f826effffff             jb 0x727
  1977: 4d8b5140                 movq 0x40(%r9), %r10
  1981: 4c8b1c24                 movq (%rsp), %r11
  1985: 498b4010                 movq 0x10(%r8), %rax
  1989: 428a1418                 movb (%rax, %r11), %dl
  1993: 8ada                     movb %dl, %bl
  1995: 498b4018                 movq 0x18(%r8), %rax
  1999: 428a1418                 movb (%rax, %r11), %dl
  2003: 32da                     xorb %dl, %bl
  2005: 498b4038                 movq 0x38(%r8), %rax
  2009: 428a1418                 movb (%rax, %r11), %dl
  2013: 32da                     xorb %dl, %bl
  2015: 498b4048                 movq 0x48(%r8), %rax
  2019: 428a1418                 movb (%rax, %r11), %dl
  2023: 32da                     xorb %dl, %bl
  2025: 498b4060                 movq 0x60(%r8), %rax
  2029: 428a1418                 movb (%rax, %r11), %dl
  2033: 32da                     xorb %dl, %bl
  2035: 498b4078                 movq 0x78(%r8), %rax
  2039: 428a1418                 movb (%rax, %r11), %dl
  2043: 32da                     xorb %dl, %bl
  2045: 498b80a8000000           movq 0xa8(%r8), %rax
  2052: 428a1418                 movb (%rax, %r11), %dl
  2056: 32da                     xorb %dl, %bl
  2058: 498b80d8000000           movq 0xd8(%r8), %rax
  2065: 428a1418                 movb (%rax, %r11), %dl
  2069: 32da                     xorb %dl, %bl
  2071: 498b80e8000000           movq 0xe8(%r8), %rax
  2078: 428a1418                 movb (%rax, %r11), %dl
  2082: 32da                     xorb %dl, %bl
  2084: 498b8008010000           movq 0x108(%r8), %rax
  2091: 428a1418                 movb (%rax, %r11), %dl
  2095: 32da                     xorb %dl, %bl
  2097: 498b8010010000           movq 0x110(%r8), %rax
  2104: 428a1418                 movb (%rax, %r11), %dl
  2108: 32da                     xorb %dl, %bl
  2110: 43881c1a                 movb %bl, (%r10, %r11)
  2114: 33db                     xorl %ebx, %ebx
  2116: 49ffc3                   incq %r11
  2119: 4c3bd9                   cmpq %rcx, %r11
  2122: 0f8271ffffff             jb 0x7c1
  2128: 4d8b5148                 movq 0x48(%r9), %r10
  2132: 4c8b1c24                 movq (%rsp), %r11
  2136: 498b4018                 movq 0x18(%r8), %rax
  2140: 428a1418                 movb (%rax, %r11), %dl
  2144: 8ada                     movb %dl, %bl
  2146: 498b4020                 movq 0x20(%r8), %rax
  2150: 428a1418                 movb (%rax, %r11), %dl
  2154: 32da                     xorb %dl, %bl
  2156: 498b4040                 movq 0x40(%r8), %rax
  2160: 428a1418                 movb (%rax, %r11), %dl
  2164: 32da                     xorb %dl, %bl
  2166: 498b4050                 movq 0x50(%r8), %rax
  2170: 428a1418                 movb (%rax, %r11), %dl
  2174: 32da                     xorb %dl, %bl
  2176: 498b4068                 movq 0x68(%r8), %rax
  2180: 428a1418                 movb (%rax, %r11), %dl
  2184: 32da                     xorb %dl, %bl
  2186: 498b8080000000           movq 0x80(%r8), %rax
  2193: 428a1418                 movb (%rax, %r11), %dl
  2197: 32da                     xorb %dl, %bl
  2199: 498b8090000000           movq 0x90(%r8), %rax
  2206: 428a1418                 movb (%rax, %r11), %dl
  2210: 32da                     xorb %dl, %bl
  2212: 498b80b0000000           movq 0xb0(%r8), %rax
  2219: 428a1418                 movb (%rax, %r11), %dl
  2223: 32da                     xorb %dl, %bl
  2225: 498b80e0000000           movq 0xe0(%r8), %rax
  2232: 428a1418                 movb (%rax, %r11), %dl
  2236: 32da                     xorb %dl, %bl
  2238: 498b8010010000           movq 0x110(%r8), %rax
  2245: 428a1418                 movb (%rax, %r11), %dl
  2249: 32da                     xorb %dl, %bl
  2251: 498b8018010000           movq 0x118(%r8), %rax
  2258: 428a1418                 movb (%rax, %r11), %dl
  2262: 32da                     xorb %dl, %bl
  2264: 43881c1a                 movb %bl, (%r10, %r11)
  2268: 33db                     xorl %ebx, %ebx
  2270: 49ffc3                   incq %r11
  2273: 4c3bd9                   cmpq %rcx, %r11
  2276: 0f826effffff             jb 0x858
  2282: 4d8b5150                 movq 0x50(%r9), %r10
  2286: 4c8b1c24                 movq (%rsp), %r11
  2290: 498b4020                 movq 0x20(%r8), %rax
  2294: 428a1418                 movb (%rax, %r11), %dl
  2298: 8ada                     movb %dl, %bl
  2300: 498b4028                 movq 0x28(%r8), %rax
  2304: 428a1418                 movb (%rax, %r11), %dl
  2308: 32da                     xorb %dl, %bl
  2310: 498b4048                 movq 0x48(%r8), %rax
  2314: 428a1418                 movb (%rax, %r11), %dl
  2318: 32da                     xorb %dl, %bl
  2320: 498b4058                 movq 0x58(%r8), %rax
  2324: 428a1418                 movb (%rax, %r11), %dl
  2328: 32da                     xorb %dl, %bl
  2330: 498b4070                 movq 0x70(%r8), %rax
  2334: 428a1418                 movb (%rax, %r11), %dl
  2338: 32da                     xorb %dl, %bl
  2340: 498b8088000000           movq 0x88(%r8), %rax
  2347: 428a1418                 movb (%rax, %r11), %dl
  2351: 32da                     xorb %dl, %bl
  2353: 498b8098000000           movq 0x98(%r8), %rax
  2360: 428a1418                 movb (%rax, %r11), %dl
  2364: 32da                     xorb %dl, %bl
  2366: 498b80b8000000           movq 0xb8(%r8), %rax
  2373: 428a1418                 movb (%rax, %r11), %dl
  2377: 32da                     xorb %dl, %bl
  2379: 498b80c0000000           movq 0xc0(%r8), %rax
  2386: 428a1418                 movb (%rax, %r11), %dl
  2390: 32da                     xorb %dl, %bl
  2392: 498b80e8000000           movq 0xe8(%r8), %rax
  2399: 428a1418                 movb (%rax, %r11), %dl
  2403: 32da                     xorb %dl, %bl
  2405: 498b8018010000           movq 0x118(%r8), %rax
  2412: 428a1418                 movb (%rax, %r11), %dl
  2416: 32da                     xorb %dl, %bl
  2418: 43881c1a                 movb %bl, (%r10, %r11)
  2422: 33db                     xorl %ebx, %ebx
  2424: 49ffc3                   incq %r11
  2427: 4c3bd9                   cmpq %rcx, %r11
  2430: 0f826effffff             jb 0x8f2
  2436: 4d8b5158                 movq 0x58(%r9), %r10
  2440: 4c8b1c24                 movq (%rsp), %r11
  2444: 498b4028                 movq 0x28(%r8), %rax
  2448: 428a1418                 movb (%rax, %r11), %dl
  2452: 8ada                     movb %dl, %bl
  2454: 498b4050                 movq 0x50(%r8), %rax
  2458: 428a1418                 movb (%rax, %r11), %dl
  2462: 32da                     xorb %dl, %bl
  2464: 498b4078                 movq 0x78(%r8), %rax
  2468: 428a1418                 movb (%rax, %r11), %dl
  2472: 32da                     xorb %dl, %bl
  2474: 498b80a0000000           movq 0xa0(%r8), %rax
  2481: 428a1418                 movb (%rax, %r11), %dl
  2485: 32da                     xorb %dl, %bl
  2487: 498b80c8000000           movq 0xc8(%r8), %rax
  2494: 428a1418                 movb (%rax, %r11), %dl
  2498: 32da                     xorb %dl, %bl
  2500: 498b80f0000000           movq 0xf0(%r8), %rax
  2507: 428a1418                 movb (%rax, %r11), %dl
  2511: 32da                     xorb %dl, %bl
  2513: 43881c1a                 movb %bl, (%r10, %r11)
  2517: 33db                     xorl %ebx, %ebx
  2519: 49ffc3                   incq %r11
  2522: 4c3bd9                   cmpq %rcx, %r11
  2525: 0f82a9ffffff             jb 0x98c
  2531: 4883c408                 addq $8, %rsp
  2535: 415e                     popq %r14
  2537: 415d                     popq %r13
  2539: 415c                     popq %r12
  2541: 5b                       popq %rbx
  2542: c5f877                   vzeroupper 
  2545: c3                       retq 
  2546: 0f0b                     ud2 
  2548: 0f0b                     ud2 
  2550: 0f0b                     ud2 
  2552: 0f0b                     ud2 
  2554: 0f0b                     ud2 
  2556: 0f0b                     ud2 
  2558: 0f0b                     ud2 
