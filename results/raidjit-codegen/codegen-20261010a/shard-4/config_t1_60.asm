; config_t1_60 target=avx512-gfni twin=xorsched topo=rdp(7)/enc op=encode U=1
; code_size=2688 twins=["xorsched"] uops_per_byte=2.609375
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
    41: 4883e840                 subq $0x40, %rax
    45: 0f8279040000             jb 0x4ac
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 498b5038                 movq 0x38(%r8), %rdx
    68: 62317d486f241f           vmovdqa32 (%rdi, %r11), %zmm12
    75: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
    82: 62d17d486fcc             vmovdqa32 %zmm12, %zmm1
    88: 62d17d486ff4             vmovdqa32 %zmm12, %zmm6
    94: 498b5040                 movq 0x40(%r8), %rdx
    98: 62117d486f241c           vmovdqa32 (%r12, %r11), %zmm12
   105: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   112: 62d17d486fd4             vmovdqa32 %zmm12, %zmm2
   118: 62d17d486ffc             vmovdqa32 %zmm12, %zmm7
   124: 498b5048                 movq 0x48(%r8), %rdx
   128: 62117d486f641d00         vmovdqa32 (%r13, %r11), %zmm12
   136: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   143: 62d17d486fdc             vmovdqa32 %zmm12, %zmm3
   149: 62517d486fc4             vmovdqa32 %zmm12, %zmm8
   155: 498b5050                 movq 0x50(%r8), %rdx
   159: 62117d486f241e           vmovdqa32 (%r14, %r11), %zmm12
   166: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   173: 62d17d486fe4             vmovdqa32 %zmm12, %zmm4
   179: 62517d486fcc             vmovdqa32 %zmm12, %zmm9
   185: 4d8b5028                 movq 0x28(%r8), %r10
   189: 498b5058                 movq 0x58(%r8), %rdx
   193: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   200: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   207: 62d17d486fec             vmovdqa32 %zmm12, %zmm5
   213: 62517d486fd4             vmovdqa32 %zmm12, %zmm10
   219: 4d8b5068                 movq 0x68(%r8), %r10
   223: 498b9098000000           movq 0x98(%r8), %rdx
   230: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   237: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   244: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   250: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   256: 4d8b5070                 movq 0x70(%r8), %r10
   260: 498b90a0000000           movq 0xa0(%r8), %rdx
   267: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   274: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   281: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   287: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   293: 4d8b5078                 movq 0x78(%r8), %r10
   297: 498b90a8000000           movq 0xa8(%r8), %rdx
   304: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   311: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   318: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   324: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   330: 4d8b9080000000           movq 0x80(%r8), %r10
   337: 498b90b0000000           movq 0xb0(%r8), %rdx
   344: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   351: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   358: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
   364: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   370: 4d8b9088000000           movq 0x88(%r8), %r10
   377: 498b90b8000000           movq 0xb8(%r8), %rdx
   384: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   391: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   398: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
   404: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   410: 4d8b90c8000000           movq 0xc8(%r8), %r10
   417: 498b90f8000000           movq 0xf8(%r8), %rdx
   424: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   431: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   438: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   444: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   450: 4d8b90d0000000           movq 0xd0(%r8), %r10
   457: 498b9000010000           movq 0x100(%r8), %rdx
   464: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   471: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   478: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   484: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   490: 4d8b90d8000000           movq 0xd8(%r8), %r10
   497: 498b9008010000           movq 0x108(%r8), %rdx
   504: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   511: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   518: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   524: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   530: 4d8b90e0000000           movq 0xe0(%r8), %r10
   537: 498b9010010000           movq 0x110(%r8), %rdx
   544: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   551: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   558: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
   564: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   570: 4d8b90e8000000           movq 0xe8(%r8), %r10
   577: 498b9018010000           movq 0x118(%r8), %rdx
   584: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   591: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   598: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
   604: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   610: 62317d486f241e           vmovdqa32 (%rsi, %r11), %zmm12
   617: 62d17d486fc4             vmovdqa32 %zmm12, %zmm0
   623: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   629: 62b14548ef3c1f           vpxord (%rdi, %r11), %zmm7, %zmm7
   636: 62113d48ef041c           vpxord (%r12, %r11), %zmm8, %zmm8
   643: 62113548ef4c1d00         vpxord (%r13, %r11), %zmm9, %zmm9
   651: 62112d48ef141e           vpxord (%r14, %r11), %zmm10, %zmm10
   658: 4d8b5028                 movq 0x28(%r8), %r10
   662: 62117d486f1c1a           vmovdqa32 (%r10, %r11), %zmm11
   669: 4d8b5030                 movq 0x30(%r8), %r10
   673: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   680: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   686: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   692: 4d8b5038                 movq 0x38(%r8), %r10
   696: 62113d48ef041a           vpxord (%r10, %r11), %zmm8, %zmm8
   703: 4d8b5040                 movq 0x40(%r8), %r10
   707: 62113548ef0c1a           vpxord (%r10, %r11), %zmm9, %zmm9
   714: 4d8b5048                 movq 0x48(%r8), %r10
   718: 62112d48ef141a           vpxord (%r10, %r11), %zmm10, %zmm10
   725: 4d8b5050                 movq 0x50(%r8), %r10
   729: 62112548ef1c1a           vpxord (%r10, %r11), %zmm11, %zmm11
   736: 4d8b5060                 movq 0x60(%r8), %r10
   740: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   747: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   753: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   759: 4d8b5068                 movq 0x68(%r8), %r10
   763: 62113548ef0c1a           vpxord (%r10, %r11), %zmm9, %zmm9
   770: 4d8b5070                 movq 0x70(%r8), %r10
   774: 62112d48ef141a           vpxord (%r10, %r11), %zmm10, %zmm10
   781: 4d8b5078                 movq 0x78(%r8), %r10
   785: 62112548ef1c1a           vpxord (%r10, %r11), %zmm11, %zmm11
   792: 4d8b9088000000           movq 0x88(%r8), %r10
   799: 62914d48ef341a           vpxord (%r10, %r11), %zmm6, %zmm6
   806: 4d8b9090000000           movq 0x90(%r8), %r10
   813: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   820: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   826: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   832: 4d8b9098000000           movq 0x98(%r8), %r10
   839: 62112d48ef141a           vpxord (%r10, %r11), %zmm10, %zmm10
   846: 4d8b90a0000000           movq 0xa0(%r8), %r10
   853: 62112548ef1c1a           vpxord (%r10, %r11), %zmm11, %zmm11
   860: 4d8b90b0000000           movq 0xb0(%r8), %r10
   867: 62914d48ef341a           vpxord (%r10, %r11), %zmm6, %zmm6
   874: 4d8b90b8000000           movq 0xb8(%r8), %r10
   881: 62914548ef3c1a           vpxord (%r10, %r11), %zmm7, %zmm7
   888: 4d8b90c0000000           movq 0xc0(%r8), %r10
   895: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   902: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   908: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   914: 4d8b90c8000000           movq 0xc8(%r8), %r10
   921: 62112548ef1c1a           vpxord (%r10, %r11), %zmm11, %zmm11
   928: 4d8b90d8000000           movq 0xd8(%r8), %r10
   935: 62914d48ef341a           vpxord (%r10, %r11), %zmm6, %zmm6
   942: 4d8b90e0000000           movq 0xe0(%r8), %r10
   949: 62914548ef3c1a           vpxord (%r10, %r11), %zmm7, %zmm7
   956: 4d8b90e8000000           movq 0xe8(%r8), %r10
   963: 62113d48ef041a           vpxord (%r10, %r11), %zmm8, %zmm8
   970: 4d8b90f0000000           movq 0xf0(%r8), %r10
   977: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   984: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   990: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   996: 4d8b9000010000           movq 0x100(%r8), %r10
  1003: 62914d48ef341a           vpxord (%r10, %r11), %zmm6, %zmm6
  1010: 4d8b9008010000           movq 0x108(%r8), %r10
  1017: 62914548ef3c1a           vpxord (%r10, %r11), %zmm7, %zmm7
  1024: 4d8b9010010000           movq 0x110(%r8), %r10
  1031: 62113d48ef041a           vpxord (%r10, %r11), %zmm8, %zmm8
  1038: 4d8b9018010000           movq 0x118(%r8), %r10
  1045: 62113548ef0c1a           vpxord (%r10, %r11), %zmm9, %zmm9
  1052: 4d8b11                   movq (%r9), %r10
  1055: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  1062: 4d8b5108                 movq 8(%r9), %r10
  1066: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
  1073: 4d8b5110                 movq 0x10(%r9), %r10
  1077: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
  1084: 4d8b5118                 movq 0x18(%r9), %r10
  1088: 62917d487f1c1a           vmovdqa32 %zmm3, (%r10, %r11)
  1095: 4d8b5120                 movq 0x20(%r9), %r10
  1099: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
  1106: 4d8b5128                 movq 0x28(%r9), %r10
  1110: 62917d487f2c1a           vmovdqa32 %zmm5, (%r10, %r11)
  1117: 4d8b5130                 movq 0x30(%r9), %r10
  1121: 62917d487f341a           vmovdqa32 %zmm6, (%r10, %r11)
  1128: 4d8b5138                 movq 0x38(%r9), %r10
  1132: 62917d487f3c1a           vmovdqa32 %zmm7, (%r10, %r11)
  1139: 4d8b5140                 movq 0x40(%r9), %r10
  1143: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  1150: 4d8b5148                 movq 0x48(%r9), %r10
  1154: 62117d487f0c1a           vmovdqa32 %zmm9, (%r10, %r11)
  1161: 4d8b5150                 movq 0x50(%r9), %r10
  1165: 62117d487f141a           vmovdqa32 %zmm10, (%r10, %r11)
  1172: 4d8b5158                 movq 0x58(%r9), %r10
  1176: 62117d487f1c1a           vmovdqa32 %zmm11, (%r10, %r11)
  1183: 4983c340                 addq $0x40, %r11
  1187: 4c3bd8                   cmpq %rax, %r11
  1190: 0f8694fbffff             jbe 0x40
  1196: 4c3bd9                   cmpq %rcx, %r11
  1199: 0f83a4050000             jae 0xa59
  1205: 4153                     pushq %r11
  1207: 4d8b11                   movq (%r9), %r10
  1210: 4c8b1c24                 movq (%rsp), %r11
  1214: 498b00                   movq (%r8), %rax
  1217: 428a1418                 movb (%rax, %r11), %dl
  1221: 8ada                     movb %dl, %bl
  1223: 498b4030                 movq 0x30(%r8), %rax
  1227: 428a1418                 movb (%rax, %r11), %dl
  1231: 32da                     xorb %dl, %bl
  1233: 498b4060                 movq 0x60(%r8), %rax
  1237: 428a1418                 movb (%rax, %r11), %dl
  1241: 32da                     xorb %dl, %bl
  1243: 498b8090000000           movq 0x90(%r8), %rax
  1250: 428a1418                 movb (%rax, %r11), %dl
  1254: 32da                     xorb %dl, %bl
  1256: 498b80c0000000           movq 0xc0(%r8), %rax
  1263: 428a1418                 movb (%rax, %r11), %dl
  1267: 32da                     xorb %dl, %bl
  1269: 498b80f0000000           movq 0xf0(%r8), %rax
  1276: 428a1418                 movb (%rax, %r11), %dl
  1280: 32da                     xorb %dl, %bl
  1282: 43881c1a                 movb %bl, (%r10, %r11)
  1286: 33db                     xorl %ebx, %ebx
  1288: 49ffc3                   incq %r11
  1291: 4c3bd9                   cmpq %rcx, %r11
  1294: 0f82aaffffff             jb 0x4be
  1300: 4d8b5108                 movq 8(%r9), %r10
  1304: 4c8b1c24                 movq (%rsp), %r11
  1308: 498b4008                 movq 8(%r8), %rax
  1312: 428a1418                 movb (%rax, %r11), %dl
  1316: 8ada                     movb %dl, %bl
  1318: 498b4038                 movq 0x38(%r8), %rax
  1322: 428a1418                 movb (%rax, %r11), %dl
  1326: 32da                     xorb %dl, %bl
  1328: 498b4068                 movq 0x68(%r8), %rax
  1332: 428a1418                 movb (%rax, %r11), %dl
  1336: 32da                     xorb %dl, %bl
  1338: 498b8098000000           movq 0x98(%r8), %rax
  1345: 428a1418                 movb (%rax, %r11), %dl
  1349: 32da                     xorb %dl, %bl
  1351: 498b80c8000000           movq 0xc8(%r8), %rax
  1358: 428a1418                 movb (%rax, %r11), %dl
  1362: 32da                     xorb %dl, %bl
  1364: 498b80f8000000           movq 0xf8(%r8), %rax
  1371: 428a1418                 movb (%rax, %r11), %dl
  1375: 32da                     xorb %dl, %bl
  1377: 43881c1a                 movb %bl, (%r10, %r11)
  1381: 33db                     xorl %ebx, %ebx
  1383: 49ffc3                   incq %r11
  1386: 4c3bd9                   cmpq %rcx, %r11
  1389: 0f82a9ffffff             jb 0x51c
  1395: 4d8b5110                 movq 0x10(%r9), %r10
  1399: 4c8b1c24                 movq (%rsp), %r11
  1403: 498b4010                 movq 0x10(%r8), %rax
  1407: 428a1418                 movb (%rax, %r11), %dl
  1411: 8ada                     movb %dl, %bl
  1413: 498b4040                 movq 0x40(%r8), %rax
  1417: 428a1418                 movb (%rax, %r11), %dl
  1421: 32da                     xorb %dl, %bl
  1423: 498b4070                 movq 0x70(%r8), %rax
  1427: 428a1418                 movb (%rax, %r11), %dl
  1431: 32da                     xorb %dl, %bl
  1433: 498b80a0000000           movq 0xa0(%r8), %rax
  1440: 428a1418                 movb (%rax, %r11), %dl
  1444: 32da                     xorb %dl, %bl
  1446: 498b80d0000000           movq 0xd0(%r8), %rax
  1453: 428a1418                 movb (%rax, %r11), %dl
  1457: 32da                     xorb %dl, %bl
  1459: 498b8000010000           movq 0x100(%r8), %rax
  1466: 428a1418                 movb (%rax, %r11), %dl
  1470: 32da                     xorb %dl, %bl
  1472: 43881c1a                 movb %bl, (%r10, %r11)
  1476: 33db                     xorl %ebx, %ebx
  1478: 49ffc3                   incq %r11
  1481: 4c3bd9                   cmpq %rcx, %r11
  1484: 0f82a9ffffff             jb 0x57b
  1490: 4d8b5118                 movq 0x18(%r9), %r10
  1494: 4c8b1c24                 movq (%rsp), %r11
  1498: 498b4018                 movq 0x18(%r8), %rax
  1502: 428a1418                 movb (%rax, %r11), %dl
  1506: 8ada                     movb %dl, %bl
  1508: 498b4048                 movq 0x48(%r8), %rax
  1512: 428a1418                 movb (%rax, %r11), %dl
  1516: 32da                     xorb %dl, %bl
  1518: 498b4078                 movq 0x78(%r8), %rax
  1522: 428a1418                 movb (%rax, %r11), %dl
  1526: 32da                     xorb %dl, %bl
  1528: 498b80a8000000           movq 0xa8(%r8), %rax
  1535: 428a1418                 movb (%rax, %r11), %dl
  1539: 32da                     xorb %dl, %bl
  1541: 498b80d8000000           movq 0xd8(%r8), %rax
  1548: 428a1418                 movb (%rax, %r11), %dl
  1552: 32da                     xorb %dl, %bl
  1554: 498b8008010000           movq 0x108(%r8), %rax
  1561: 428a1418                 movb (%rax, %r11), %dl
  1565: 32da                     xorb %dl, %bl
  1567: 43881c1a                 movb %bl, (%r10, %r11)
  1571: 33db                     xorl %ebx, %ebx
  1573: 49ffc3                   incq %r11
  1576: 4c3bd9                   cmpq %rcx, %r11
  1579: 0f82a9ffffff             jb 0x5da
  1585: 4d8b5120                 movq 0x20(%r9), %r10
  1589: 4c8b1c24                 movq (%rsp), %r11
  1593: 498b4020                 movq 0x20(%r8), %rax
  1597: 428a1418                 movb (%rax, %r11), %dl
  1601: 8ada                     movb %dl, %bl
  1603: 498b4050                 movq 0x50(%r8), %rax
  1607: 428a1418                 movb (%rax, %r11), %dl
  1611: 32da                     xorb %dl, %bl
  1613: 498b8080000000           movq 0x80(%r8), %rax
  1620: 428a1418                 movb (%rax, %r11), %dl
  1624: 32da                     xorb %dl, %bl
  1626: 498b80b0000000           movq 0xb0(%r8), %rax
  1633: 428a1418                 movb (%rax, %r11), %dl
  1637: 32da                     xorb %dl, %bl
  1639: 498b80e0000000           movq 0xe0(%r8), %rax
  1646: 428a1418                 movb (%rax, %r11), %dl
  1650: 32da                     xorb %dl, %bl
  1652: 498b8010010000           movq 0x110(%r8), %rax
  1659: 428a1418                 movb (%rax, %r11), %dl
  1663: 32da                     xorb %dl, %bl
  1665: 43881c1a                 movb %bl, (%r10, %r11)
  1669: 33db                     xorl %ebx, %ebx
  1671: 49ffc3                   incq %r11
  1674: 4c3bd9                   cmpq %rcx, %r11
  1677: 0f82a6ffffff             jb 0x639
  1683: 4d8b5128                 movq 0x28(%r9), %r10
  1687: 4c8b1c24                 movq (%rsp), %r11
  1691: 498b4028                 movq 0x28(%r8), %rax
  1695: 428a1418                 movb (%rax, %r11), %dl
  1699: 8ada                     movb %dl, %bl
  1701: 498b4058                 movq 0x58(%r8), %rax
  1705: 428a1418                 movb (%rax, %r11), %dl
  1709: 32da                     xorb %dl, %bl
  1711: 498b8088000000           movq 0x88(%r8), %rax
  1718: 428a1418                 movb (%rax, %r11), %dl
  1722: 32da                     xorb %dl, %bl
  1724: 498b80b8000000           movq 0xb8(%r8), %rax
  1731: 428a1418                 movb (%rax, %r11), %dl
  1735: 32da                     xorb %dl, %bl
  1737: 498b80e8000000           movq 0xe8(%r8), %rax
  1744: 428a1418                 movb (%rax, %r11), %dl
  1748: 32da                     xorb %dl, %bl
  1750: 498b8018010000           movq 0x118(%r8), %rax
  1757: 428a1418                 movb (%rax, %r11), %dl
  1761: 32da                     xorb %dl, %bl
  1763: 43881c1a                 movb %bl, (%r10, %r11)
  1767: 33db                     xorl %ebx, %ebx
  1769: 49ffc3                   incq %r11
  1772: 4c3bd9                   cmpq %rcx, %r11
  1775: 0f82a6ffffff             jb 0x69b
  1781: 4d8b5130                 movq 0x30(%r9), %r10
  1785: 4c8b1c24                 movq (%rsp), %r11
  1789: 498b00                   movq (%r8), %rax
  1792: 428a1418                 movb (%rax, %r11), %dl
  1796: 8ada                     movb %dl, %bl
  1798: 498b4008                 movq 8(%r8), %rax
  1802: 428a1418                 movb (%rax, %r11), %dl
  1806: 32da                     xorb %dl, %bl
  1808: 498b4038                 movq 0x38(%r8), %rax
  1812: 428a1418                 movb (%rax, %r11), %dl
  1816: 32da                     xorb %dl, %bl
  1818: 498b4068                 movq 0x68(%r8), %rax
  1822: 428a1418                 movb (%rax, %r11), %dl
  1826: 32da                     xorb %dl, %bl
  1828: 498b8088000000           movq 0x88(%r8), %rax
  1835: 428a1418                 movb (%rax, %r11), %dl
  1839: 32da                     xorb %dl, %bl
  1841: 498b8098000000           movq 0x98(%r8), %rax
  1848: 428a1418                 movb (%rax, %r11), %dl
  1852: 32da                     xorb %dl, %bl
  1854: 498b80b0000000           movq 0xb0(%r8), %rax
  1861: 428a1418                 movb (%rax, %r11), %dl
  1865: 32da                     xorb %dl, %bl
  1867: 498b80c8000000           movq 0xc8(%r8), %rax
  1874: 428a1418                 movb (%rax, %r11), %dl
  1878: 32da                     xorb %dl, %bl
  1880: 498b80d8000000           movq 0xd8(%r8), %rax
  1887: 428a1418                 movb (%rax, %r11), %dl
  1891: 32da                     xorb %dl, %bl
  1893: 498b80f8000000           movq 0xf8(%r8), %rax
  1900: 428a1418                 movb (%rax, %r11), %dl
  1904: 32da                     xorb %dl, %bl
  1906: 498b8000010000           movq 0x100(%r8), %rax
  1913: 428a1418                 movb (%rax, %r11), %dl
  1917: 32da                     xorb %dl, %bl
  1919: 43881c1a                 movb %bl, (%r10, %r11)
  1923: 33db                     xorl %ebx, %ebx
  1925: 49ffc3                   incq %r11
  1928: 4c3bd9                   cmpq %rcx, %r11
  1931: 0f826cffffff             jb 0x6fd
  1937: 4d8b5138                 movq 0x38(%r9), %r10
  1941: 4c8b1c24                 movq (%rsp), %r11
  1945: 498b4008                 movq 8(%r8), %rax
  1949: 428a1418                 movb (%rax, %r11), %dl
  1953: 8ada                     movb %dl, %bl
  1955: 498b4010                 movq 0x10(%r8), %rax
  1959: 428a1418                 movb (%rax, %r11), %dl
  1963: 32da                     xorb %dl, %bl
  1965: 498b4030                 movq 0x30(%r8), %rax
  1969: 428a1418                 movb (%rax, %r11), %dl
  1973: 32da                     xorb %dl, %bl
  1975: 498b4040                 movq 0x40(%r8), %rax
  1979: 428a1418                 movb (%rax, %r11), %dl
  1983: 32da                     xorb %dl, %bl
  1985: 498b4070                 movq 0x70(%r8), %rax
  1989: 428a1418                 movb (%rax, %r11), %dl
  1993: 32da                     xorb %dl, %bl
  1995: 498b80a0000000           movq 0xa0(%r8), %rax
  2002: 428a1418                 movb (%rax, %r11), %dl
  2006: 32da                     xorb %dl, %bl
  2008: 498b80b8000000           movq 0xb8(%r8), %rax
  2015: 428a1418                 movb (%rax, %r11), %dl
  2019: 32da                     xorb %dl, %bl
  2021: 498b80d0000000           movq 0xd0(%r8), %rax
  2028: 428a1418                 movb (%rax, %r11), %dl
  2032: 32da                     xorb %dl, %bl
  2034: 498b80e0000000           movq 0xe0(%r8), %rax
  2041: 428a1418                 movb (%rax, %r11), %dl
  2045: 32da                     xorb %dl, %bl
  2047: 498b8000010000           movq 0x100(%r8), %rax
  2054: 428a1418                 movb (%rax, %r11), %dl
  2058: 32da                     xorb %dl, %bl
  2060: 498b8008010000           movq 0x108(%r8), %rax
  2067: 428a1418                 movb (%rax, %r11), %dl
  2071: 32da                     xorb %dl, %bl
  2073: 43881c1a                 movb %bl, (%r10, %r11)
  2077: 33db                     xorl %ebx, %ebx
  2079: 49ffc3                   incq %r11
  2082: 4c3bd9                   cmpq %rcx, %r11
  2085: 0f826effffff             jb 0x799
  2091: 4d8b5140                 movq 0x40(%r9), %r10
  2095: 4c8b1c24                 movq (%rsp), %r11
  2099: 498b4010                 movq 0x10(%r8), %rax
  2103: 428a1418                 movb (%rax, %r11), %dl
  2107: 8ada                     movb %dl, %bl
  2109: 498b4018                 movq 0x18(%r8), %rax
  2113: 428a1418                 movb (%rax, %r11), %dl
  2117: 32da                     xorb %dl, %bl
  2119: 498b4038                 movq 0x38(%r8), %rax
  2123: 428a1418                 movb (%rax, %r11), %dl
  2127: 32da                     xorb %dl, %bl
  2129: 498b4048                 movq 0x48(%r8), %rax
  2133: 428a1418                 movb (%rax, %r11), %dl
  2137: 32da                     xorb %dl, %bl
  2139: 498b4060                 movq 0x60(%r8), %rax
  2143: 428a1418                 movb (%rax, %r11), %dl
  2147: 32da                     xorb %dl, %bl
  2149: 498b4078                 movq 0x78(%r8), %rax
  2153: 428a1418                 movb (%rax, %r11), %dl
  2157: 32da                     xorb %dl, %bl
  2159: 498b80a8000000           movq 0xa8(%r8), %rax
  2166: 428a1418                 movb (%rax, %r11), %dl
  2170: 32da                     xorb %dl, %bl
  2172: 498b80d8000000           movq 0xd8(%r8), %rax
  2179: 428a1418                 movb (%rax, %r11), %dl
  2183: 32da                     xorb %dl, %bl
  2185: 498b80e8000000           movq 0xe8(%r8), %rax
  2192: 428a1418                 movb (%rax, %r11), %dl
  2196: 32da                     xorb %dl, %bl
  2198: 498b8008010000           movq 0x108(%r8), %rax
  2205: 428a1418                 movb (%rax, %r11), %dl
  2209: 32da                     xorb %dl, %bl
  2211: 498b8010010000           movq 0x110(%r8), %rax
  2218: 428a1418                 movb (%rax, %r11), %dl
  2222: 32da                     xorb %dl, %bl
  2224: 43881c1a                 movb %bl, (%r10, %r11)
  2228: 33db                     xorl %ebx, %ebx
  2230: 49ffc3                   incq %r11
  2233: 4c3bd9                   cmpq %rcx, %r11
  2236: 0f8271ffffff             jb 0x833
  2242: 4d8b5148                 movq 0x48(%r9), %r10
  2246: 4c8b1c24                 movq (%rsp), %r11
  2250: 498b4018                 movq 0x18(%r8), %rax
  2254: 428a1418                 movb (%rax, %r11), %dl
  2258: 8ada                     movb %dl, %bl
  2260: 498b4020                 movq 0x20(%r8), %rax
  2264: 428a1418                 movb (%rax, %r11), %dl
  2268: 32da                     xorb %dl, %bl
  2270: 498b4040                 movq 0x40(%r8), %rax
  2274: 428a1418                 movb (%rax, %r11), %dl
  2278: 32da                     xorb %dl, %bl
  2280: 498b4050                 movq 0x50(%r8), %rax
  2284: 428a1418                 movb (%rax, %r11), %dl
  2288: 32da                     xorb %dl, %bl
  2290: 498b4068                 movq 0x68(%r8), %rax
  2294: 428a1418                 movb (%rax, %r11), %dl
  2298: 32da                     xorb %dl, %bl
  2300: 498b8080000000           movq 0x80(%r8), %rax
  2307: 428a1418                 movb (%rax, %r11), %dl
  2311: 32da                     xorb %dl, %bl
  2313: 498b8090000000           movq 0x90(%r8), %rax
  2320: 428a1418                 movb (%rax, %r11), %dl
  2324: 32da                     xorb %dl, %bl
  2326: 498b80b0000000           movq 0xb0(%r8), %rax
  2333: 428a1418                 movb (%rax, %r11), %dl
  2337: 32da                     xorb %dl, %bl
  2339: 498b80e0000000           movq 0xe0(%r8), %rax
  2346: 428a1418                 movb (%rax, %r11), %dl
  2350: 32da                     xorb %dl, %bl
  2352: 498b8010010000           movq 0x110(%r8), %rax
  2359: 428a1418                 movb (%rax, %r11), %dl
  2363: 32da                     xorb %dl, %bl
  2365: 498b8018010000           movq 0x118(%r8), %rax
  2372: 428a1418                 movb (%rax, %r11), %dl
  2376: 32da                     xorb %dl, %bl
  2378: 43881c1a                 movb %bl, (%r10, %r11)
  2382: 33db                     xorl %ebx, %ebx
  2384: 49ffc3                   incq %r11
  2387: 4c3bd9                   cmpq %rcx, %r11
  2390: 0f826effffff             jb 0x8ca
  2396: 4d8b5150                 movq 0x50(%r9), %r10
  2400: 4c8b1c24                 movq (%rsp), %r11
  2404: 498b4020                 movq 0x20(%r8), %rax
  2408: 428a1418                 movb (%rax, %r11), %dl
  2412: 8ada                     movb %dl, %bl
  2414: 498b4028                 movq 0x28(%r8), %rax
  2418: 428a1418                 movb (%rax, %r11), %dl
  2422: 32da                     xorb %dl, %bl
  2424: 498b4048                 movq 0x48(%r8), %rax
  2428: 428a1418                 movb (%rax, %r11), %dl
  2432: 32da                     xorb %dl, %bl
  2434: 498b4058                 movq 0x58(%r8), %rax
  2438: 428a1418                 movb (%rax, %r11), %dl
  2442: 32da                     xorb %dl, %bl
  2444: 498b4070                 movq 0x70(%r8), %rax
  2448: 428a1418                 movb (%rax, %r11), %dl
  2452: 32da                     xorb %dl, %bl
  2454: 498b8088000000           movq 0x88(%r8), %rax
  2461: 428a1418                 movb (%rax, %r11), %dl
  2465: 32da                     xorb %dl, %bl
  2467: 498b8098000000           movq 0x98(%r8), %rax
  2474: 428a1418                 movb (%rax, %r11), %dl
  2478: 32da                     xorb %dl, %bl
  2480: 498b80b8000000           movq 0xb8(%r8), %rax
  2487: 428a1418                 movb (%rax, %r11), %dl
  2491: 32da                     xorb %dl, %bl
  2493: 498b80c0000000           movq 0xc0(%r8), %rax
  2500: 428a1418                 movb (%rax, %r11), %dl
  2504: 32da                     xorb %dl, %bl
  2506: 498b80e8000000           movq 0xe8(%r8), %rax
  2513: 428a1418                 movb (%rax, %r11), %dl
  2517: 32da                     xorb %dl, %bl
  2519: 498b8018010000           movq 0x118(%r8), %rax
  2526: 428a1418                 movb (%rax, %r11), %dl
  2530: 32da                     xorb %dl, %bl
  2532: 43881c1a                 movb %bl, (%r10, %r11)
  2536: 33db                     xorl %ebx, %ebx
  2538: 49ffc3                   incq %r11
  2541: 4c3bd9                   cmpq %rcx, %r11
  2544: 0f826effffff             jb 0x964
  2550: 4d8b5158                 movq 0x58(%r9), %r10
  2554: 4c8b1c24                 movq (%rsp), %r11
  2558: 498b4028                 movq 0x28(%r8), %rax
  2562: 428a1418                 movb (%rax, %r11), %dl
  2566: 8ada                     movb %dl, %bl
  2568: 498b4050                 movq 0x50(%r8), %rax
  2572: 428a1418                 movb (%rax, %r11), %dl
  2576: 32da                     xorb %dl, %bl
  2578: 498b4078                 movq 0x78(%r8), %rax
  2582: 428a1418                 movb (%rax, %r11), %dl
  2586: 32da                     xorb %dl, %bl
  2588: 498b80a0000000           movq 0xa0(%r8), %rax
  2595: 428a1418                 movb (%rax, %r11), %dl
  2599: 32da                     xorb %dl, %bl
  2601: 498b80c8000000           movq 0xc8(%r8), %rax
  2608: 428a1418                 movb (%rax, %r11), %dl
  2612: 32da                     xorb %dl, %bl
  2614: 498b80f0000000           movq 0xf0(%r8), %rax
  2621: 428a1418                 movb (%rax, %r11), %dl
  2625: 32da                     xorb %dl, %bl
  2627: 43881c1a                 movb %bl, (%r10, %r11)
  2631: 33db                     xorl %ebx, %ebx
  2633: 49ffc3                   incq %r11
  2636: 4c3bd9                   cmpq %rcx, %r11
  2639: 0f82a9ffffff             jb 0x9fe
  2645: 4883c408                 addq $8, %rsp
  2649: 415e                     popq %r14
  2651: 415d                     popq %r13
  2653: 415c                     popq %r12
  2655: 5b                       popq %rbx
  2656: c5f877                   vzeroupper 
  2659: c3                       retq 
  2660: 0f0b                     ud2 
  2662: 0f0b                     ud2 
  2664: 0f0b                     ud2 
  2666: 0f0b                     ud2 
  2668: 0f0b                     ud2 
  2670: 0f0b                     ud2 
  2672: 0f0b                     ud2 
  2674: 0f0b                     ud2 
  2676: 0f0b                     ud2 
  2678: 0f0b                     ud2 
  2680: 0f0b                     ud2 
  2682: 0f0b                     ud2 
  2684: 0f0b                     ud2 
  2686: 0f0b                     ud2 
