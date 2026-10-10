; config_t1_138 target=avx512-gfni twin=xorsched topo=liberation(8)/enc op=encode U=1
; code_size=5504 twins=["xorsched"] uops_per_byte=5.765625
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
    45: 0f82d8090000             jb 0xa0b
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 4d8b5028                 movq 0x28(%r8), %r10
    68: 498b9080000000           movq 0x80(%r8), %rdx
    75: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
    82: 62a14d40ef341a           vpxord (%rdx, %r11), %zmm22, %zmm22
    89: 62b17d486fee             vmovdqa32 %zmm22, %zmm5
    95: 62a17d486fc6             vmovdqa32 %zmm22, %zmm16
   101: 4d8b5058                 movq 0x58(%r8), %r10
   105: 498b90b0000000           movq 0xb0(%r8), %rdx
   112: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   119: 62a14d40ef341a           vpxord (%rdx, %r11), %zmm22, %zmm22
   126: 62b17d486fc6             vmovdqa32 %zmm22, %zmm0
   132: 62a17d486fee             vmovdqa32 %zmm22, %zmm21
   138: 4d8b90e0000000           movq 0xe0(%r8), %r10
   145: 498b9038010000           movq 0x138(%r8), %rdx
   152: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   159: 62a14d40ef341a           vpxord (%rdx, %r11), %zmm22, %zmm22
   166: 62b17d486ff6             vmovdqa32 %zmm22, %zmm6
   172: 62317d486ffe             vmovdqa32 %zmm22, %zmm15
   178: 4d8b9010010000           movq 0x110(%r8), %r10
   185: 498b9068010000           movq 0x168(%r8), %rdx
   192: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   199: 62a14d40ef341a           vpxord (%rdx, %r11), %zmm22, %zmm22
   206: 62b17d486fce             vmovdqa32 %zmm22, %zmm1
   212: 62a17d486fe6             vmovdqa32 %zmm22, %zmm20
   218: 4d8b9098010000           movq 0x198(%r8), %r10
   225: 498b90f0010000           movq 0x1f0(%r8), %rdx
   232: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   239: 62a14d40ef341a           vpxord (%rdx, %r11), %zmm22, %zmm22
   246: 62b17d486ffe             vmovdqa32 %zmm22, %zmm7
   252: 62317d486ff6             vmovdqa32 %zmm22, %zmm14
   258: 4d8b90c8010000           movq 0x1c8(%r8), %r10
   265: 498b9020020000           movq 0x220(%r8), %rdx
   272: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   279: 62a14d40ef341a           vpxord (%rdx, %r11), %zmm22, %zmm22
   286: 62b17d486fd6             vmovdqa32 %zmm22, %zmm2
   292: 62a17d486fde             vmovdqa32 %zmm22, %zmm19
   298: 4d8b9050020000           movq 0x250(%r8), %r10
   305: 498b90a8020000           movq 0x2a8(%r8), %rdx
   312: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   319: 62a14d40ef341a           vpxord (%rdx, %r11), %zmm22, %zmm22
   326: 62317d486fc6             vmovdqa32 %zmm22, %zmm8
   332: 62317d486fee             vmovdqa32 %zmm22, %zmm13
   338: 62a17d486f341e           vmovdqa32 (%rsi, %r11), %zmm22
   345: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
   351: 62317d486fde             vmovdqa32 %zmm22, %zmm11
   357: 62a17d486f341f           vmovdqa32 (%rdi, %r11), %zmm22
   364: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
   370: 62317d486fe6             vmovdqa32 %zmm22, %zmm12
   376: 62817d486f341c           vmovdqa32 (%r12, %r11), %zmm22
   383: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
   389: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
   395: 62817d486f741d00         vmovdqa32 (%r13, %r11), %zmm22
   403: 62b17d486fde             vmovdqa32 %zmm22, %zmm3
   409: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
   415: 62817d486f341e           vmovdqa32 (%r14, %r11), %zmm22
   422: 62b17d486fe6             vmovdqa32 %zmm22, %zmm4
   428: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
   434: 4d8b5030                 movq 0x30(%r8), %r10
   438: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   445: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
   451: 62a17d486fce             vmovdqa32 %zmm22, %zmm17
   457: 4d8b5038                 movq 0x38(%r8), %r10
   461: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   468: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
   474: 62a17d486fd6             vmovdqa32 %zmm22, %zmm18
   480: 4d8b5040                 movq 0x40(%r8), %r10
   484: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   491: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
   497: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
   503: 4d8b5048                 movq 0x48(%r8), %r10
   507: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   514: 62317d486fce             vmovdqa32 %zmm22, %zmm9
   520: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
   526: 4d8b5050                 movq 0x50(%r8), %r10
   530: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   537: 62317d486fd6             vmovdqa32 %zmm22, %zmm10
   543: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
   549: 4d8b5060                 movq 0x60(%r8), %r10
   553: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   560: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
   566: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
   572: 4d8b5068                 movq 0x68(%r8), %r10
   576: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   583: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
   589: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
   595: 4d8b5070                 movq 0x70(%r8), %r10
   599: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   606: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
   612: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
   618: 4d8b5078                 movq 0x78(%r8), %r10
   622: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   629: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
   635: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
   641: 4d8b9080000000           movq 0x80(%r8), %r10
   648: 62110548ef3c1a           vpxord (%r10, %r11), %zmm15, %zmm15
   655: 4d8b9088000000           movq 0x88(%r8), %r10
   662: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   669: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
   675: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
   681: 4d8b9090000000           movq 0x90(%r8), %r10
   688: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   695: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
   701: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
   707: 4d8b9098000000           movq 0x98(%r8), %r10
   714: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   721: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
   727: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
   733: 4d8b90a0000000           movq 0xa0(%r8), %r10
   740: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   747: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
   753: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
   759: 4d8b90a8000000           movq 0xa8(%r8), %r10
   766: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   773: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
   779: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
   785: 4d8b90b0000000           movq 0xb0(%r8), %r10
   792: 62815d40ef241a           vpxord (%r10, %r11), %zmm20, %zmm20
   799: 4d8b90b8000000           movq 0xb8(%r8), %r10
   806: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   813: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
   819: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
   825: 4d8b90c0000000           movq 0xc0(%r8), %r10
   832: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   839: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
   845: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
   851: 4d8b90c8000000           movq 0xc8(%r8), %r10
   858: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   865: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
   871: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
   877: 4d8b90d0000000           movq 0xd0(%r8), %r10
   884: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   891: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
   897: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
   903: 4d8b90d8000000           movq 0xd8(%r8), %r10
   910: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   917: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
   923: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
   929: 4d8b90e8000000           movq 0xe8(%r8), %r10
   936: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   943: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
   949: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
   955: 4d8b90f0000000           movq 0xf0(%r8), %r10
   962: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   969: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
   975: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
   981: 4d8b90f8000000           movq 0xf8(%r8), %r10
   988: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   995: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  1001: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1007: 4d8b9000010000           movq 0x100(%r8), %r10
  1014: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1021: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1027: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1033: 4d8b9008010000           movq 0x108(%r8), %r10
  1040: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1047: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  1053: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1059: 4d8b9018010000           movq 0x118(%r8), %r10
  1066: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1073: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  1079: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  1085: 4d8b9020010000           movq 0x120(%r8), %r10
  1092: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1099: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  1105: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  1111: 4d8b9028010000           movq 0x128(%r8), %r10
  1118: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1125: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  1131: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1137: 4d8b9030010000           movq 0x130(%r8), %r10
  1144: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1151: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  1157: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  1163: 4d8b9038010000           movq 0x138(%r8), %r10
  1170: 62110d48ef341a           vpxord (%r10, %r11), %zmm14, %zmm14
  1177: 4d8b9040010000           movq 0x140(%r8), %r10
  1184: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1191: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  1197: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  1203: 4d8b9048010000           movq 0x148(%r8), %r10
  1210: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1217: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  1223: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1229: 4d8b9050010000           movq 0x150(%r8), %r10
  1236: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1243: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  1249: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1255: 4d8b9058010000           movq 0x158(%r8), %r10
  1262: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1269: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1275: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1281: 4d8b9060010000           movq 0x160(%r8), %r10
  1288: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1295: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  1301: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1307: 4d8b9068010000           movq 0x168(%r8), %r10
  1314: 62816540ef1c1a           vpxord (%r10, %r11), %zmm19, %zmm19
  1321: 4d8b9070010000           movq 0x170(%r8), %r10
  1328: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1335: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  1341: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1347: 4d8b9078010000           movq 0x178(%r8), %r10
  1354: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1361: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  1367: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  1373: 4d8b9080010000           movq 0x180(%r8), %r10
  1380: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1387: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  1393: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  1399: 4d8b9088010000           movq 0x188(%r8), %r10
  1406: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1413: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  1419: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1425: 4d8b9090010000           movq 0x190(%r8), %r10
  1432: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1439: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  1445: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  1451: 4d8b90a0010000           movq 0x1a0(%r8), %r10
  1458: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1465: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  1471: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  1477: 4d8b90a8010000           movq 0x1a8(%r8), %r10
  1484: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1491: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  1497: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1503: 4d8b90b0010000           movq 0x1b0(%r8), %r10
  1510: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1517: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1523: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1529: 4d8b90b8010000           movq 0x1b8(%r8), %r10
  1536: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1543: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  1549: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1555: 4d8b90c0010000           movq 0x1c0(%r8), %r10
  1562: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1569: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  1575: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1581: 4d8b90d0010000           movq 0x1d0(%r8), %r10
  1588: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1595: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  1601: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1607: 4d8b90d8010000           movq 0x1d8(%r8), %r10
  1614: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1621: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  1627: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  1633: 4d8b90e0010000           movq 0x1e0(%r8), %r10
  1640: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1647: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  1653: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  1659: 4d8b90e8010000           movq 0x1e8(%r8), %r10
  1666: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1673: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  1679: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1685: 4d8b90f0010000           movq 0x1f0(%r8), %r10
  1692: 62111548ef2c1a           vpxord (%r10, %r11), %zmm13, %zmm13
  1699: 4d8b90f8010000           movq 0x1f8(%r8), %r10
  1706: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1713: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  1719: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  1725: 4d8b9000020000           movq 0x200(%r8), %r10
  1732: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1739: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  1745: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  1751: 4d8b9008020000           movq 0x208(%r8), %r10
  1758: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1765: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1771: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1777: 4d8b9010020000           movq 0x210(%r8), %r10
  1784: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1791: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  1797: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1803: 4d8b9018020000           movq 0x218(%r8), %r10
  1810: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1817: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  1823: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1829: 4d8b9020020000           movq 0x220(%r8), %r10
  1836: 62816d40ef141a           vpxord (%r10, %r11), %zmm18, %zmm18
  1843: 4d8b9028020000           movq 0x228(%r8), %r10
  1850: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1857: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  1863: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1869: 4d8b9030020000           movq 0x230(%r8), %r10
  1876: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1883: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  1889: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1895: 4d8b9038020000           movq 0x238(%r8), %r10
  1902: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1909: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  1915: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  1921: 4d8b9040020000           movq 0x240(%r8), %r10
  1928: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1935: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  1941: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  1947: 4d8b9048020000           movq 0x248(%r8), %r10
  1954: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1961: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  1967: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1973: 4d8b9058020000           movq 0x258(%r8), %r10
  1980: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1987: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  1993: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  1999: 4d8b9060020000           movq 0x260(%r8), %r10
  2006: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2013: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  2019: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  2025: 4d8b9068020000           movq 0x268(%r8), %r10
  2032: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2039: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  2045: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  2051: 4d8b9070020000           movq 0x270(%r8), %r10
  2058: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2065: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  2071: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  2077: 4d8b9078020000           movq 0x278(%r8), %r10
  2084: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2091: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  2097: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  2103: 4d8b9080020000           movq 0x280(%r8), %r10
  2110: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2117: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  2123: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  2129: 4d8b9088020000           movq 0x288(%r8), %r10
  2136: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2143: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  2149: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  2155: 4d8b9090020000           movq 0x290(%r8), %r10
  2162: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2169: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  2175: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  2181: 4d8b9098020000           movq 0x298(%r8), %r10
  2188: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2195: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  2201: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  2207: 4d8b90a0020000           movq 0x2a0(%r8), %r10
  2214: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2221: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  2227: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  2233: 4d8b90a8020000           movq 0x2a8(%r8), %r10
  2240: 62111d48ef241a           vpxord (%r10, %r11), %zmm12, %zmm12
  2247: 4d8b90b0020000           movq 0x2b0(%r8), %r10
  2254: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2261: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  2267: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  2273: 4d8b90b8020000           movq 0x2b8(%r8), %r10
  2280: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2287: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  2293: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  2299: 4d8b11                   movq (%r9), %r10
  2302: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  2309: 4d8b5108                 movq 8(%r9), %r10
  2313: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
  2320: 4d8b5110                 movq 0x10(%r9), %r10
  2324: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
  2331: 4d8b5118                 movq 0x18(%r9), %r10
  2335: 62917d487f1c1a           vmovdqa32 %zmm3, (%r10, %r11)
  2342: 4d8b5120                 movq 0x20(%r9), %r10
  2346: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
  2353: 4d8b5128                 movq 0x28(%r9), %r10
  2357: 62917d487f2c1a           vmovdqa32 %zmm5, (%r10, %r11)
  2364: 4d8b5130                 movq 0x30(%r9), %r10
  2368: 62917d487f341a           vmovdqa32 %zmm6, (%r10, %r11)
  2375: 4d8b5138                 movq 0x38(%r9), %r10
  2379: 62917d487f3c1a           vmovdqa32 %zmm7, (%r10, %r11)
  2386: 4d8b5140                 movq 0x40(%r9), %r10
  2390: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  2397: 4d8b5148                 movq 0x48(%r9), %r10
  2401: 62117d487f0c1a           vmovdqa32 %zmm9, (%r10, %r11)
  2408: 4d8b5150                 movq 0x50(%r9), %r10
  2412: 62117d487f141a           vmovdqa32 %zmm10, (%r10, %r11)
  2419: 4d8b5158                 movq 0x58(%r9), %r10
  2423: 62117d487f1c1a           vmovdqa32 %zmm11, (%r10, %r11)
  2430: 4d8b5160                 movq 0x60(%r9), %r10
  2434: 62117d487f241a           vmovdqa32 %zmm12, (%r10, %r11)
  2441: 4d8b5168                 movq 0x68(%r9), %r10
  2445: 62117d487f2c1a           vmovdqa32 %zmm13, (%r10, %r11)
  2452: 4d8b5170                 movq 0x70(%r9), %r10
  2456: 62117d487f341a           vmovdqa32 %zmm14, (%r10, %r11)
  2463: 4d8b5178                 movq 0x78(%r9), %r10
  2467: 62117d487f3c1a           vmovdqa32 %zmm15, (%r10, %r11)
  2474: 4d8b9180000000           movq 0x80(%r9), %r10
  2481: 62817d487f041a           vmovdqa32 %zmm16, (%r10, %r11)
  2488: 4d8b9188000000           movq 0x88(%r9), %r10
  2495: 62817d487f0c1a           vmovdqa32 %zmm17, (%r10, %r11)
  2502: 4d8b9190000000           movq 0x90(%r9), %r10
  2509: 62817d487f141a           vmovdqa32 %zmm18, (%r10, %r11)
  2516: 4d8b9198000000           movq 0x98(%r9), %r10
  2523: 62817d487f1c1a           vmovdqa32 %zmm19, (%r10, %r11)
  2530: 4d8b91a0000000           movq 0xa0(%r9), %r10
  2537: 62817d487f241a           vmovdqa32 %zmm20, (%r10, %r11)
  2544: 4d8b91a8000000           movq 0xa8(%r9), %r10
  2551: 62817d487f2c1a           vmovdqa32 %zmm21, (%r10, %r11)
  2558: 4983c340                 addq $0x40, %r11
  2562: 4c3bd8                   cmpq %rax, %r11
  2565: 0f8635f6ffff             jbe 0x40
  2571: 4c3bd9                   cmpq %rcx, %r11
  2574: 0f833c0b0000             jae 0x1550
  2580: 4153                     pushq %r11
  2582: 4d8b11                   movq (%r9), %r10
  2585: 4c8b1c24                 movq (%rsp), %r11
  2589: 498b00                   movq (%r8), %rax
  2592: 428a1418                 movb (%rax, %r11), %dl
  2596: 8ada                     movb %dl, %bl
  2598: 498b4058                 movq 0x58(%r8), %rax
  2602: 428a1418                 movb (%rax, %r11), %dl
  2606: 32da                     xorb %dl, %bl
  2608: 498b80b0000000           movq 0xb0(%r8), %rax
  2615: 428a1418                 movb (%rax, %r11), %dl
  2619: 32da                     xorb %dl, %bl
  2621: 498b8008010000           movq 0x108(%r8), %rax
  2628: 428a1418                 movb (%rax, %r11), %dl
  2632: 32da                     xorb %dl, %bl
  2634: 498b8060010000           movq 0x160(%r8), %rax
  2641: 428a1418                 movb (%rax, %r11), %dl
  2645: 32da                     xorb %dl, %bl
  2647: 498b80b8010000           movq 0x1b8(%r8), %rax
  2654: 428a1418                 movb (%rax, %r11), %dl
  2658: 32da                     xorb %dl, %bl
  2660: 498b8010020000           movq 0x210(%r8), %rax
  2667: 428a1418                 movb (%rax, %r11), %dl
  2671: 32da                     xorb %dl, %bl
  2673: 498b8068020000           movq 0x268(%r8), %rax
  2680: 428a1418                 movb (%rax, %r11), %dl
  2684: 32da                     xorb %dl, %bl
  2686: 43881c1a                 movb %bl, (%r10, %r11)
  2690: 33db                     xorl %ebx, %ebx
  2692: 49ffc3                   incq %r11
  2695: 4c3bd9                   cmpq %rcx, %r11
  2698: 0f828dffffff             jb 0xa1d
  2704: 4d8b5108                 movq 8(%r9), %r10
  2708: 4c8b1c24                 movq (%rsp), %r11
  2712: 498b4008                 movq 8(%r8), %rax
  2716: 428a1418                 movb (%rax, %r11), %dl
  2720: 8ada                     movb %dl, %bl
  2722: 498b4060                 movq 0x60(%r8), %rax
  2726: 428a1418                 movb (%rax, %r11), %dl
  2730: 32da                     xorb %dl, %bl
  2732: 498b80b8000000           movq 0xb8(%r8), %rax
  2739: 428a1418                 movb (%rax, %r11), %dl
  2743: 32da                     xorb %dl, %bl
  2745: 498b8010010000           movq 0x110(%r8), %rax
  2752: 428a1418                 movb (%rax, %r11), %dl
  2756: 32da                     xorb %dl, %bl
  2758: 498b8068010000           movq 0x168(%r8), %rax
  2765: 428a1418                 movb (%rax, %r11), %dl
  2769: 32da                     xorb %dl, %bl
  2771: 498b80c0010000           movq 0x1c0(%r8), %rax
  2778: 428a1418                 movb (%rax, %r11), %dl
  2782: 32da                     xorb %dl, %bl
  2784: 498b8018020000           movq 0x218(%r8), %rax
  2791: 428a1418                 movb (%rax, %r11), %dl
  2795: 32da                     xorb %dl, %bl
  2797: 498b8070020000           movq 0x270(%r8), %rax
  2804: 428a1418                 movb (%rax, %r11), %dl
  2808: 32da                     xorb %dl, %bl
  2810: 43881c1a                 movb %bl, (%r10, %r11)
  2814: 33db                     xorl %ebx, %ebx
  2816: 49ffc3                   incq %r11
  2819: 4c3bd9                   cmpq %rcx, %r11
  2822: 0f828cffffff             jb 0xa98
  2828: 4d8b5110                 movq 0x10(%r9), %r10
  2832: 4c8b1c24                 movq (%rsp), %r11
  2836: 498b4010                 movq 0x10(%r8), %rax
  2840: 428a1418                 movb (%rax, %r11), %dl
  2844: 8ada                     movb %dl, %bl
  2846: 498b4068                 movq 0x68(%r8), %rax
  2850: 428a1418                 movb (%rax, %r11), %dl
  2854: 32da                     xorb %dl, %bl
  2856: 498b80c0000000           movq 0xc0(%r8), %rax
  2863: 428a1418                 movb (%rax, %r11), %dl
  2867: 32da                     xorb %dl, %bl
  2869: 498b8018010000           movq 0x118(%r8), %rax
  2876: 428a1418                 movb (%rax, %r11), %dl
  2880: 32da                     xorb %dl, %bl
  2882: 498b8070010000           movq 0x170(%r8), %rax
  2889: 428a1418                 movb (%rax, %r11), %dl
  2893: 32da                     xorb %dl, %bl
  2895: 498b80c8010000           movq 0x1c8(%r8), %rax
  2902: 428a1418                 movb (%rax, %r11), %dl
  2906: 32da                     xorb %dl, %bl
  2908: 498b8020020000           movq 0x220(%r8), %rax
  2915: 428a1418                 movb (%rax, %r11), %dl
  2919: 32da                     xorb %dl, %bl
  2921: 498b8078020000           movq 0x278(%r8), %rax
  2928: 428a1418                 movb (%rax, %r11), %dl
  2932: 32da                     xorb %dl, %bl
  2934: 43881c1a                 movb %bl, (%r10, %r11)
  2938: 33db                     xorl %ebx, %ebx
  2940: 49ffc3                   incq %r11
  2943: 4c3bd9                   cmpq %rcx, %r11
  2946: 0f828cffffff             jb 0xb14
  2952: 4d8b5118                 movq 0x18(%r9), %r10
  2956: 4c8b1c24                 movq (%rsp), %r11
  2960: 498b4018                 movq 0x18(%r8), %rax
  2964: 428a1418                 movb (%rax, %r11), %dl
  2968: 8ada                     movb %dl, %bl
  2970: 498b4070                 movq 0x70(%r8), %rax
  2974: 428a1418                 movb (%rax, %r11), %dl
  2978: 32da                     xorb %dl, %bl
  2980: 498b80c8000000           movq 0xc8(%r8), %rax
  2987: 428a1418                 movb (%rax, %r11), %dl
  2991: 32da                     xorb %dl, %bl
  2993: 498b8020010000           movq 0x120(%r8), %rax
  3000: 428a1418                 movb (%rax, %r11), %dl
  3004: 32da                     xorb %dl, %bl
  3006: 498b8078010000           movq 0x178(%r8), %rax
  3013: 428a1418                 movb (%rax, %r11), %dl
  3017: 32da                     xorb %dl, %bl
  3019: 498b80d0010000           movq 0x1d0(%r8), %rax
  3026: 428a1418                 movb (%rax, %r11), %dl
  3030: 32da                     xorb %dl, %bl
  3032: 498b8028020000           movq 0x228(%r8), %rax
  3039: 428a1418                 movb (%rax, %r11), %dl
  3043: 32da                     xorb %dl, %bl
  3045: 498b8080020000           movq 0x280(%r8), %rax
  3052: 428a1418                 movb (%rax, %r11), %dl
  3056: 32da                     xorb %dl, %bl
  3058: 43881c1a                 movb %bl, (%r10, %r11)
  3062: 33db                     xorl %ebx, %ebx
  3064: 49ffc3                   incq %r11
  3067: 4c3bd9                   cmpq %rcx, %r11
  3070: 0f828cffffff             jb 0xb90
  3076: 4d8b5120                 movq 0x20(%r9), %r10
  3080: 4c8b1c24                 movq (%rsp), %r11
  3084: 498b4020                 movq 0x20(%r8), %rax
  3088: 428a1418                 movb (%rax, %r11), %dl
  3092: 8ada                     movb %dl, %bl
  3094: 498b4078                 movq 0x78(%r8), %rax
  3098: 428a1418                 movb (%rax, %r11), %dl
  3102: 32da                     xorb %dl, %bl
  3104: 498b80d0000000           movq 0xd0(%r8), %rax
  3111: 428a1418                 movb (%rax, %r11), %dl
  3115: 32da                     xorb %dl, %bl
  3117: 498b8028010000           movq 0x128(%r8), %rax
  3124: 428a1418                 movb (%rax, %r11), %dl
  3128: 32da                     xorb %dl, %bl
  3130: 498b8080010000           movq 0x180(%r8), %rax
  3137: 428a1418                 movb (%rax, %r11), %dl
  3141: 32da                     xorb %dl, %bl
  3143: 498b80d8010000           movq 0x1d8(%r8), %rax
  3150: 428a1418                 movb (%rax, %r11), %dl
  3154: 32da                     xorb %dl, %bl
  3156: 498b8030020000           movq 0x230(%r8), %rax
  3163: 428a1418                 movb (%rax, %r11), %dl
  3167: 32da                     xorb %dl, %bl
  3169: 498b8088020000           movq 0x288(%r8), %rax
  3176: 428a1418                 movb (%rax, %r11), %dl
  3180: 32da                     xorb %dl, %bl
  3182: 43881c1a                 movb %bl, (%r10, %r11)
  3186: 33db                     xorl %ebx, %ebx
  3188: 49ffc3                   incq %r11
  3191: 4c3bd9                   cmpq %rcx, %r11
  3194: 0f828cffffff             jb 0xc0c
  3200: 4d8b5128                 movq 0x28(%r9), %r10
  3204: 4c8b1c24                 movq (%rsp), %r11
  3208: 498b4028                 movq 0x28(%r8), %rax
  3212: 428a1418                 movb (%rax, %r11), %dl
  3216: 8ada                     movb %dl, %bl
  3218: 498b8080000000           movq 0x80(%r8), %rax
  3225: 428a1418                 movb (%rax, %r11), %dl
  3229: 32da                     xorb %dl, %bl
  3231: 498b80d8000000           movq 0xd8(%r8), %rax
  3238: 428a1418                 movb (%rax, %r11), %dl
  3242: 32da                     xorb %dl, %bl
  3244: 498b8030010000           movq 0x130(%r8), %rax
  3251: 428a1418                 movb (%rax, %r11), %dl
  3255: 32da                     xorb %dl, %bl
  3257: 498b8088010000           movq 0x188(%r8), %rax
  3264: 428a1418                 movb (%rax, %r11), %dl
  3268: 32da                     xorb %dl, %bl
  3270: 498b80e0010000           movq 0x1e0(%r8), %rax
  3277: 428a1418                 movb (%rax, %r11), %dl
  3281: 32da                     xorb %dl, %bl
  3283: 498b8038020000           movq 0x238(%r8), %rax
  3290: 428a1418                 movb (%rax, %r11), %dl
  3294: 32da                     xorb %dl, %bl
  3296: 498b8090020000           movq 0x290(%r8), %rax
  3303: 428a1418                 movb (%rax, %r11), %dl
  3307: 32da                     xorb %dl, %bl
  3309: 43881c1a                 movb %bl, (%r10, %r11)
  3313: 33db                     xorl %ebx, %ebx
  3315: 49ffc3                   incq %r11
  3318: 4c3bd9                   cmpq %rcx, %r11
  3321: 0f8289ffffff             jb 0xc88
  3327: 4d8b5130                 movq 0x30(%r9), %r10
  3331: 4c8b1c24                 movq (%rsp), %r11
  3335: 498b4030                 movq 0x30(%r8), %rax
  3339: 428a1418                 movb (%rax, %r11), %dl
  3343: 8ada                     movb %dl, %bl
  3345: 498b8088000000           movq 0x88(%r8), %rax
  3352: 428a1418                 movb (%rax, %r11), %dl
  3356: 32da                     xorb %dl, %bl
  3358: 498b80e0000000           movq 0xe0(%r8), %rax
  3365: 428a1418                 movb (%rax, %r11), %dl
  3369: 32da                     xorb %dl, %bl
  3371: 498b8038010000           movq 0x138(%r8), %rax
  3378: 428a1418                 movb (%rax, %r11), %dl
  3382: 32da                     xorb %dl, %bl
  3384: 498b8090010000           movq 0x190(%r8), %rax
  3391: 428a1418                 movb (%rax, %r11), %dl
  3395: 32da                     xorb %dl, %bl
  3397: 498b80e8010000           movq 0x1e8(%r8), %rax
  3404: 428a1418                 movb (%rax, %r11), %dl
  3408: 32da                     xorb %dl, %bl
  3410: 498b8040020000           movq 0x240(%r8), %rax
  3417: 428a1418                 movb (%rax, %r11), %dl
  3421: 32da                     xorb %dl, %bl
  3423: 498b8098020000           movq 0x298(%r8), %rax
  3430: 428a1418                 movb (%rax, %r11), %dl
  3434: 32da                     xorb %dl, %bl
  3436: 43881c1a                 movb %bl, (%r10, %r11)
  3440: 33db                     xorl %ebx, %ebx
  3442: 49ffc3                   incq %r11
  3445: 4c3bd9                   cmpq %rcx, %r11
  3448: 0f8289ffffff             jb 0xd07
  3454: 4d8b5138                 movq 0x38(%r9), %r10
  3458: 4c8b1c24                 movq (%rsp), %r11
  3462: 498b4038                 movq 0x38(%r8), %rax
  3466: 428a1418                 movb (%rax, %r11), %dl
  3470: 8ada                     movb %dl, %bl
  3472: 498b8090000000           movq 0x90(%r8), %rax
  3479: 428a1418                 movb (%rax, %r11), %dl
  3483: 32da                     xorb %dl, %bl
  3485: 498b80e8000000           movq 0xe8(%r8), %rax
  3492: 428a1418                 movb (%rax, %r11), %dl
  3496: 32da                     xorb %dl, %bl
  3498: 498b8040010000           movq 0x140(%r8), %rax
  3505: 428a1418                 movb (%rax, %r11), %dl
  3509: 32da                     xorb %dl, %bl
  3511: 498b8098010000           movq 0x198(%r8), %rax
  3518: 428a1418                 movb (%rax, %r11), %dl
  3522: 32da                     xorb %dl, %bl
  3524: 498b80f0010000           movq 0x1f0(%r8), %rax
  3531: 428a1418                 movb (%rax, %r11), %dl
  3535: 32da                     xorb %dl, %bl
  3537: 498b8048020000           movq 0x248(%r8), %rax
  3544: 428a1418                 movb (%rax, %r11), %dl
  3548: 32da                     xorb %dl, %bl
  3550: 498b80a0020000           movq 0x2a0(%r8), %rax
  3557: 428a1418                 movb (%rax, %r11), %dl
  3561: 32da                     xorb %dl, %bl
  3563: 43881c1a                 movb %bl, (%r10, %r11)
  3567: 33db                     xorl %ebx, %ebx
  3569: 49ffc3                   incq %r11
  3572: 4c3bd9                   cmpq %rcx, %r11
  3575: 0f8289ffffff             jb 0xd86
  3581: 4d8b5140                 movq 0x40(%r9), %r10
  3585: 4c8b1c24                 movq (%rsp), %r11
  3589: 498b4040                 movq 0x40(%r8), %rax
  3593: 428a1418                 movb (%rax, %r11), %dl
  3597: 8ada                     movb %dl, %bl
  3599: 498b8098000000           movq 0x98(%r8), %rax
  3606: 428a1418                 movb (%rax, %r11), %dl
  3610: 32da                     xorb %dl, %bl
  3612: 498b80f0000000           movq 0xf0(%r8), %rax
  3619: 428a1418                 movb (%rax, %r11), %dl
  3623: 32da                     xorb %dl, %bl
  3625: 498b8048010000           movq 0x148(%r8), %rax
  3632: 428a1418                 movb (%rax, %r11), %dl
  3636: 32da                     xorb %dl, %bl
  3638: 498b80a0010000           movq 0x1a0(%r8), %rax
  3645: 428a1418                 movb (%rax, %r11), %dl
  3649: 32da                     xorb %dl, %bl
  3651: 498b80f8010000           movq 0x1f8(%r8), %rax
  3658: 428a1418                 movb (%rax, %r11), %dl
  3662: 32da                     xorb %dl, %bl
  3664: 498b8050020000           movq 0x250(%r8), %rax
  3671: 428a1418                 movb (%rax, %r11), %dl
  3675: 32da                     xorb %dl, %bl
  3677: 498b80a8020000           movq 0x2a8(%r8), %rax
  3684: 428a1418                 movb (%rax, %r11), %dl
  3688: 32da                     xorb %dl, %bl
  3690: 43881c1a                 movb %bl, (%r10, %r11)
  3694: 33db                     xorl %ebx, %ebx
  3696: 49ffc3                   incq %r11
  3699: 4c3bd9                   cmpq %rcx, %r11
  3702: 0f8289ffffff             jb 0xe05
  3708: 4d8b5148                 movq 0x48(%r9), %r10
  3712: 4c8b1c24                 movq (%rsp), %r11
  3716: 498b4048                 movq 0x48(%r8), %rax
  3720: 428a1418                 movb (%rax, %r11), %dl
  3724: 8ada                     movb %dl, %bl
  3726: 498b80a0000000           movq 0xa0(%r8), %rax
  3733: 428a1418                 movb (%rax, %r11), %dl
  3737: 32da                     xorb %dl, %bl
  3739: 498b80f8000000           movq 0xf8(%r8), %rax
  3746: 428a1418                 movb (%rax, %r11), %dl
  3750: 32da                     xorb %dl, %bl
  3752: 498b8050010000           movq 0x150(%r8), %rax
  3759: 428a1418                 movb (%rax, %r11), %dl
  3763: 32da                     xorb %dl, %bl
  3765: 498b80a8010000           movq 0x1a8(%r8), %rax
  3772: 428a1418                 movb (%rax, %r11), %dl
  3776: 32da                     xorb %dl, %bl
  3778: 498b8000020000           movq 0x200(%r8), %rax
  3785: 428a1418                 movb (%rax, %r11), %dl
  3789: 32da                     xorb %dl, %bl
  3791: 498b8058020000           movq 0x258(%r8), %rax
  3798: 428a1418                 movb (%rax, %r11), %dl
  3802: 32da                     xorb %dl, %bl
  3804: 498b80b0020000           movq 0x2b0(%r8), %rax
  3811: 428a1418                 movb (%rax, %r11), %dl
  3815: 32da                     xorb %dl, %bl
  3817: 43881c1a                 movb %bl, (%r10, %r11)
  3821: 33db                     xorl %ebx, %ebx
  3823: 49ffc3                   incq %r11
  3826: 4c3bd9                   cmpq %rcx, %r11
  3829: 0f8289ffffff             jb 0xe84
  3835: 4d8b5150                 movq 0x50(%r9), %r10
  3839: 4c8b1c24                 movq (%rsp), %r11
  3843: 498b4050                 movq 0x50(%r8), %rax
  3847: 428a1418                 movb (%rax, %r11), %dl
  3851: 8ada                     movb %dl, %bl
  3853: 498b80a8000000           movq 0xa8(%r8), %rax
  3860: 428a1418                 movb (%rax, %r11), %dl
  3864: 32da                     xorb %dl, %bl
  3866: 498b8000010000           movq 0x100(%r8), %rax
  3873: 428a1418                 movb (%rax, %r11), %dl
  3877: 32da                     xorb %dl, %bl
  3879: 498b8058010000           movq 0x158(%r8), %rax
  3886: 428a1418                 movb (%rax, %r11), %dl
  3890: 32da                     xorb %dl, %bl
  3892: 498b80b0010000           movq 0x1b0(%r8), %rax
  3899: 428a1418                 movb (%rax, %r11), %dl
  3903: 32da                     xorb %dl, %bl
  3905: 498b8008020000           movq 0x208(%r8), %rax
  3912: 428a1418                 movb (%rax, %r11), %dl
  3916: 32da                     xorb %dl, %bl
  3918: 498b8060020000           movq 0x260(%r8), %rax
  3925: 428a1418                 movb (%rax, %r11), %dl
  3929: 32da                     xorb %dl, %bl
  3931: 498b80b8020000           movq 0x2b8(%r8), %rax
  3938: 428a1418                 movb (%rax, %r11), %dl
  3942: 32da                     xorb %dl, %bl
  3944: 43881c1a                 movb %bl, (%r10, %r11)
  3948: 33db                     xorl %ebx, %ebx
  3950: 49ffc3                   incq %r11
  3953: 4c3bd9                   cmpq %rcx, %r11
  3956: 0f8289ffffff             jb 0xf03
  3962: 4d8b5158                 movq 0x58(%r9), %r10
  3966: 4c8b1c24                 movq (%rsp), %r11
  3970: 498b00                   movq (%r8), %rax
  3973: 428a1418                 movb (%rax, %r11), %dl
  3977: 8ada                     movb %dl, %bl
  3979: 498b4060                 movq 0x60(%r8), %rax
  3983: 428a1418                 movb (%rax, %r11), %dl
  3987: 32da                     xorb %dl, %bl
  3989: 498b80c0000000           movq 0xc0(%r8), %rax
  3996: 428a1418                 movb (%rax, %r11), %dl
  4000: 32da                     xorb %dl, %bl
  4002: 498b8020010000           movq 0x120(%r8), %rax
  4009: 428a1418                 movb (%rax, %r11), %dl
  4013: 32da                     xorb %dl, %bl
  4015: 498b8080010000           movq 0x180(%r8), %rax
  4022: 428a1418                 movb (%rax, %r11), %dl
  4026: 32da                     xorb %dl, %bl
  4028: 498b80e0010000           movq 0x1e0(%r8), %rax
  4035: 428a1418                 movb (%rax, %r11), %dl
  4039: 32da                     xorb %dl, %bl
  4041: 498b8040020000           movq 0x240(%r8), %rax
  4048: 428a1418                 movb (%rax, %r11), %dl
  4052: 32da                     xorb %dl, %bl
  4054: 498b80a0020000           movq 0x2a0(%r8), %rax
  4061: 428a1418                 movb (%rax, %r11), %dl
  4065: 32da                     xorb %dl, %bl
  4067: 43881c1a                 movb %bl, (%r10, %r11)
  4071: 33db                     xorl %ebx, %ebx
  4073: 49ffc3                   incq %r11
  4076: 4c3bd9                   cmpq %rcx, %r11
  4079: 0f828dffffff             jb 0xf82
  4085: 4d8b5160                 movq 0x60(%r9), %r10
  4089: 4c8b1c24                 movq (%rsp), %r11
  4093: 498b4008                 movq 8(%r8), %rax
  4097: 428a1418                 movb (%rax, %r11), %dl
  4101: 8ada                     movb %dl, %bl
  4103: 498b4068                 movq 0x68(%r8), %rax
  4107: 428a1418                 movb (%rax, %r11), %dl
  4111: 32da                     xorb %dl, %bl
  4113: 498b80c8000000           movq 0xc8(%r8), %rax
  4120: 428a1418                 movb (%rax, %r11), %dl
  4124: 32da                     xorb %dl, %bl
  4126: 498b8028010000           movq 0x128(%r8), %rax
  4133: 428a1418                 movb (%rax, %r11), %dl
  4137: 32da                     xorb %dl, %bl
  4139: 498b8088010000           movq 0x188(%r8), %rax
  4146: 428a1418                 movb (%rax, %r11), %dl
  4150: 32da                     xorb %dl, %bl
  4152: 498b80e8010000           movq 0x1e8(%r8), %rax
  4159: 428a1418                 movb (%rax, %r11), %dl
  4163: 32da                     xorb %dl, %bl
  4165: 498b8048020000           movq 0x248(%r8), %rax
  4172: 428a1418                 movb (%rax, %r11), %dl
  4176: 32da                     xorb %dl, %bl
  4178: 498b80a8020000           movq 0x2a8(%r8), %rax
  4185: 428a1418                 movb (%rax, %r11), %dl
  4189: 32da                     xorb %dl, %bl
  4191: 43881c1a                 movb %bl, (%r10, %r11)
  4195: 33db                     xorl %ebx, %ebx
  4197: 49ffc3                   incq %r11
  4200: 4c3bd9                   cmpq %rcx, %r11
  4203: 0f828cffffff             jb 0xffd
  4209: 4d8b5168                 movq 0x68(%r9), %r10
  4213: 4c8b1c24                 movq (%rsp), %r11
  4217: 498b4010                 movq 0x10(%r8), %rax
  4221: 428a1418                 movb (%rax, %r11), %dl
  4225: 8ada                     movb %dl, %bl
  4227: 498b4070                 movq 0x70(%r8), %rax
  4231: 428a1418                 movb (%rax, %r11), %dl
  4235: 32da                     xorb %dl, %bl
  4237: 498b80d0000000           movq 0xd0(%r8), %rax
  4244: 428a1418                 movb (%rax, %r11), %dl
  4248: 32da                     xorb %dl, %bl
  4250: 498b8030010000           movq 0x130(%r8), %rax
  4257: 428a1418                 movb (%rax, %r11), %dl
  4261: 32da                     xorb %dl, %bl
  4263: 498b8090010000           movq 0x190(%r8), %rax
  4270: 428a1418                 movb (%rax, %r11), %dl
  4274: 32da                     xorb %dl, %bl
  4276: 498b80f0010000           movq 0x1f0(%r8), %rax
  4283: 428a1418                 movb (%rax, %r11), %dl
  4287: 32da                     xorb %dl, %bl
  4289: 498b8050020000           movq 0x250(%r8), %rax
  4296: 428a1418                 movb (%rax, %r11), %dl
  4300: 32da                     xorb %dl, %bl
  4302: 498b80a8020000           movq 0x2a8(%r8), %rax
  4309: 428a1418                 movb (%rax, %r11), %dl
  4313: 32da                     xorb %dl, %bl
  4315: 498b80b0020000           movq 0x2b0(%r8), %rax
  4322: 428a1418                 movb (%rax, %r11), %dl
  4326: 32da                     xorb %dl, %bl
  4328: 43881c1a                 movb %bl, (%r10, %r11)
  4332: 33db                     xorl %ebx, %ebx
  4334: 49ffc3                   incq %r11
  4337: 4c3bd9                   cmpq %rcx, %r11
  4340: 0f827fffffff             jb 0x1079
  4346: 4d8b5170                 movq 0x70(%r9), %r10
  4350: 4c8b1c24                 movq (%rsp), %r11
  4354: 498b4018                 movq 0x18(%r8), %rax
  4358: 428a1418                 movb (%rax, %r11), %dl
  4362: 8ada                     movb %dl, %bl
  4364: 498b4078                 movq 0x78(%r8), %rax
  4368: 428a1418                 movb (%rax, %r11), %dl
  4372: 32da                     xorb %dl, %bl
  4374: 498b80d8000000           movq 0xd8(%r8), %rax
  4381: 428a1418                 movb (%rax, %r11), %dl
  4385: 32da                     xorb %dl, %bl
  4387: 498b8038010000           movq 0x138(%r8), %rax
  4394: 428a1418                 movb (%rax, %r11), %dl
  4398: 32da                     xorb %dl, %bl
  4400: 498b8098010000           movq 0x198(%r8), %rax
  4407: 428a1418                 movb (%rax, %r11), %dl
  4411: 32da                     xorb %dl, %bl
  4413: 498b80f0010000           movq 0x1f0(%r8), %rax
  4420: 428a1418                 movb (%rax, %r11), %dl
  4424: 32da                     xorb %dl, %bl
  4426: 498b80f8010000           movq 0x1f8(%r8), %rax
  4433: 428a1418                 movb (%rax, %r11), %dl
  4437: 32da                     xorb %dl, %bl
  4439: 498b8058020000           movq 0x258(%r8), %rax
  4446: 428a1418                 movb (%rax, %r11), %dl
  4450: 32da                     xorb %dl, %bl
  4452: 498b80b8020000           movq 0x2b8(%r8), %rax
  4459: 428a1418                 movb (%rax, %r11), %dl
  4463: 32da                     xorb %dl, %bl
  4465: 43881c1a                 movb %bl, (%r10, %r11)
  4469: 33db                     xorl %ebx, %ebx
  4471: 49ffc3                   incq %r11
  4474: 4c3bd9                   cmpq %rcx, %r11
  4477: 0f827fffffff             jb 0x1102
  4483: 4d8b5178                 movq 0x78(%r9), %r10
  4487: 4c8b1c24                 movq (%rsp), %r11
  4491: 498b4020                 movq 0x20(%r8), %rax
  4495: 428a1418                 movb (%rax, %r11), %dl
  4499: 8ada                     movb %dl, %bl
  4501: 498b8080000000           movq 0x80(%r8), %rax
  4508: 428a1418                 movb (%rax, %r11), %dl
  4512: 32da                     xorb %dl, %bl
  4514: 498b80e0000000           movq 0xe0(%r8), %rax
  4521: 428a1418                 movb (%rax, %r11), %dl
  4525: 32da                     xorb %dl, %bl
  4527: 498b8038010000           movq 0x138(%r8), %rax
  4534: 428a1418                 movb (%rax, %r11), %dl
  4538: 32da                     xorb %dl, %bl
  4540: 498b8040010000           movq 0x140(%r8), %rax
  4547: 428a1418                 movb (%rax, %r11), %dl
  4551: 32da                     xorb %dl, %bl
  4553: 498b80a0010000           movq 0x1a0(%r8), %rax
  4560: 428a1418                 movb (%rax, %r11), %dl
  4564: 32da                     xorb %dl, %bl
  4566: 498b8000020000           movq 0x200(%r8), %rax
  4573: 428a1418                 movb (%rax, %r11), %dl
  4577: 32da                     xorb %dl, %bl
  4579: 498b8060020000           movq 0x260(%r8), %rax
  4586: 428a1418                 movb (%rax, %r11), %dl
  4590: 32da                     xorb %dl, %bl
  4592: 498b8068020000           movq 0x268(%r8), %rax
  4599: 428a1418                 movb (%rax, %r11), %dl
  4603: 32da                     xorb %dl, %bl
  4605: 43881c1a                 movb %bl, (%r10, %r11)
  4609: 33db                     xorl %ebx, %ebx
  4611: 49ffc3                   incq %r11
  4614: 4c3bd9                   cmpq %rcx, %r11
  4617: 0f827cffffff             jb 0x118b
  4623: 4d8b9180000000           movq 0x80(%r9), %r10
  4630: 4c8b1c24                 movq (%rsp), %r11
  4634: 498b4028                 movq 0x28(%r8), %rax
  4638: 428a1418                 movb (%rax, %r11), %dl
  4642: 8ada                     movb %dl, %bl
  4644: 498b8080000000           movq 0x80(%r8), %rax
  4651: 428a1418                 movb (%rax, %r11), %dl
  4655: 32da                     xorb %dl, %bl
  4657: 498b8088000000           movq 0x88(%r8), %rax
  4664: 428a1418                 movb (%rax, %r11), %dl
  4668: 32da                     xorb %dl, %bl
  4670: 498b80e8000000           movq 0xe8(%r8), %rax
  4677: 428a1418                 movb (%rax, %r11), %dl
  4681: 32da                     xorb %dl, %bl
  4683: 498b8048010000           movq 0x148(%r8), %rax
  4690: 428a1418                 movb (%rax, %r11), %dl
  4694: 32da                     xorb %dl, %bl
  4696: 498b80a8010000           movq 0x1a8(%r8), %rax
  4703: 428a1418                 movb (%rax, %r11), %dl
  4707: 32da                     xorb %dl, %bl
  4709: 498b8008020000           movq 0x208(%r8), %rax
  4716: 428a1418                 movb (%rax, %r11), %dl
  4720: 32da                     xorb %dl, %bl
  4722: 498b8010020000           movq 0x210(%r8), %rax
  4729: 428a1418                 movb (%rax, %r11), %dl
  4733: 32da                     xorb %dl, %bl
  4735: 498b8070020000           movq 0x270(%r8), %rax
  4742: 428a1418                 movb (%rax, %r11), %dl
  4746: 32da                     xorb %dl, %bl
  4748: 43881c1a                 movb %bl, (%r10, %r11)
  4752: 33db                     xorl %ebx, %ebx
  4754: 49ffc3                   incq %r11
  4757: 4c3bd9                   cmpq %rcx, %r11
  4760: 0f827cffffff             jb 0x121a
  4766: 4d8b9188000000           movq 0x88(%r9), %r10
  4773: 4c8b1c24                 movq (%rsp), %r11
  4777: 498b4030                 movq 0x30(%r8), %rax
  4781: 428a1418                 movb (%rax, %r11), %dl
  4785: 8ada                     movb %dl, %bl
  4787: 498b8090000000           movq 0x90(%r8), %rax
  4794: 428a1418                 movb (%rax, %r11), %dl
  4798: 32da                     xorb %dl, %bl
  4800: 498b80f0000000           movq 0xf0(%r8), %rax
  4807: 428a1418                 movb (%rax, %r11), %dl
  4811: 32da                     xorb %dl, %bl
  4813: 498b8050010000           movq 0x150(%r8), %rax
  4820: 428a1418                 movb (%rax, %r11), %dl
  4824: 32da                     xorb %dl, %bl
  4826: 498b80b0010000           movq 0x1b0(%r8), %rax
  4833: 428a1418                 movb (%rax, %r11), %dl
  4837: 32da                     xorb %dl, %bl
  4839: 498b80b8010000           movq 0x1b8(%r8), %rax
  4846: 428a1418                 movb (%rax, %r11), %dl
  4850: 32da                     xorb %dl, %bl
  4852: 498b8018020000           movq 0x218(%r8), %rax
  4859: 428a1418                 movb (%rax, %r11), %dl
  4863: 32da                     xorb %dl, %bl
  4865: 498b8078020000           movq 0x278(%r8), %rax
  4872: 428a1418                 movb (%rax, %r11), %dl
  4876: 32da                     xorb %dl, %bl
  4878: 43881c1a                 movb %bl, (%r10, %r11)
  4882: 33db                     xorl %ebx, %ebx
  4884: 49ffc3                   incq %r11
  4887: 4c3bd9                   cmpq %rcx, %r11
  4890: 0f8289ffffff             jb 0x12a9
  4896: 4d8b9190000000           movq 0x90(%r9), %r10
  4903: 4c8b1c24                 movq (%rsp), %r11
  4907: 498b4038                 movq 0x38(%r8), %rax
  4911: 428a1418                 movb (%rax, %r11), %dl
  4915: 8ada                     movb %dl, %bl
  4917: 498b8098000000           movq 0x98(%r8), %rax
  4924: 428a1418                 movb (%rax, %r11), %dl
  4928: 32da                     xorb %dl, %bl
  4930: 498b80f8000000           movq 0xf8(%r8), %rax
  4937: 428a1418                 movb (%rax, %r11), %dl
  4941: 32da                     xorb %dl, %bl
  4943: 498b8058010000           movq 0x158(%r8), %rax
  4950: 428a1418                 movb (%rax, %r11), %dl
  4954: 32da                     xorb %dl, %bl
  4956: 498b8060010000           movq 0x160(%r8), %rax
  4963: 428a1418                 movb (%rax, %r11), %dl
  4967: 32da                     xorb %dl, %bl
  4969: 498b80c0010000           movq 0x1c0(%r8), %rax
  4976: 428a1418                 movb (%rax, %r11), %dl
  4980: 32da                     xorb %dl, %bl
  4982: 498b8020020000           movq 0x220(%r8), %rax
  4989: 428a1418                 movb (%rax, %r11), %dl
  4993: 32da                     xorb %dl, %bl
  4995: 498b8080020000           movq 0x280(%r8), %rax
  5002: 428a1418                 movb (%rax, %r11), %dl
  5006: 32da                     xorb %dl, %bl
  5008: 43881c1a                 movb %bl, (%r10, %r11)
  5012: 33db                     xorl %ebx, %ebx
  5014: 49ffc3                   incq %r11
  5017: 4c3bd9                   cmpq %rcx, %r11
  5020: 0f8289ffffff             jb 0x132b
  5026: 4d8b9198000000           movq 0x98(%r9), %r10
  5033: 4c8b1c24                 movq (%rsp), %r11
  5037: 498b4040                 movq 0x40(%r8), %rax
  5041: 428a1418                 movb (%rax, %r11), %dl
  5045: 8ada                     movb %dl, %bl
  5047: 498b80a0000000           movq 0xa0(%r8), %rax
  5054: 428a1418                 movb (%rax, %r11), %dl
  5058: 32da                     xorb %dl, %bl
  5060: 498b8000010000           movq 0x100(%r8), %rax
  5067: 428a1418                 movb (%rax, %r11), %dl
  5071: 32da                     xorb %dl, %bl
  5073: 498b8008010000           movq 0x108(%r8), %rax
  5080: 428a1418                 movb (%rax, %r11), %dl
  5084: 32da                     xorb %dl, %bl
  5086: 498b8068010000           movq 0x168(%r8), %rax
  5093: 428a1418                 movb (%rax, %r11), %dl
  5097: 32da                     xorb %dl, %bl
  5099: 498b80c8010000           movq 0x1c8(%r8), %rax
  5106: 428a1418                 movb (%rax, %r11), %dl
  5110: 32da                     xorb %dl, %bl
  5112: 498b8020020000           movq 0x220(%r8), %rax
  5119: 428a1418                 movb (%rax, %r11), %dl
  5123: 32da                     xorb %dl, %bl
  5125: 498b8028020000           movq 0x228(%r8), %rax
  5132: 428a1418                 movb (%rax, %r11), %dl
  5136: 32da                     xorb %dl, %bl
  5138: 498b8088020000           movq 0x288(%r8), %rax
  5145: 428a1418                 movb (%rax, %r11), %dl
  5149: 32da                     xorb %dl, %bl
  5151: 43881c1a                 movb %bl, (%r10, %r11)
  5155: 33db                     xorl %ebx, %ebx
  5157: 49ffc3                   incq %r11
  5160: 4c3bd9                   cmpq %rcx, %r11
  5163: 0f827cffffff             jb 0x13ad
  5169: 4d8b91a0000000           movq 0xa0(%r9), %r10
  5176: 4c8b1c24                 movq (%rsp), %r11
  5180: 498b4048                 movq 0x48(%r8), %rax
  5184: 428a1418                 movb (%rax, %r11), %dl
  5188: 8ada                     movb %dl, %bl
  5190: 498b80a8000000           movq 0xa8(%r8), %rax
  5197: 428a1418                 movb (%rax, %r11), %dl
  5201: 32da                     xorb %dl, %bl
  5203: 498b80b0000000           movq 0xb0(%r8), %rax
  5210: 428a1418                 movb (%rax, %r11), %dl
  5214: 32da                     xorb %dl, %bl
  5216: 498b8010010000           movq 0x110(%r8), %rax
  5223: 428a1418                 movb (%rax, %r11), %dl
  5227: 32da                     xorb %dl, %bl
  5229: 498b8068010000           movq 0x168(%r8), %rax
  5236: 428a1418                 movb (%rax, %r11), %dl
  5240: 32da                     xorb %dl, %bl
  5242: 498b8070010000           movq 0x170(%r8), %rax
  5249: 428a1418                 movb (%rax, %r11), %dl
  5253: 32da                     xorb %dl, %bl
  5255: 498b80d0010000           movq 0x1d0(%r8), %rax
  5262: 428a1418                 movb (%rax, %r11), %dl
  5266: 32da                     xorb %dl, %bl
  5268: 498b8030020000           movq 0x230(%r8), %rax
  5275: 428a1418                 movb (%rax, %r11), %dl
  5279: 32da                     xorb %dl, %bl
  5281: 498b8090020000           movq 0x290(%r8), %rax
  5288: 428a1418                 movb (%rax, %r11), %dl
  5292: 32da                     xorb %dl, %bl
  5294: 43881c1a                 movb %bl, (%r10, %r11)
  5298: 33db                     xorl %ebx, %ebx
  5300: 49ffc3                   incq %r11
  5303: 4c3bd9                   cmpq %rcx, %r11
  5306: 0f827cffffff             jb 0x143c
  5312: 4d8b91a8000000           movq 0xa8(%r9), %r10
  5319: 4c8b1c24                 movq (%rsp), %r11
  5323: 498b4050                 movq 0x50(%r8), %rax
  5327: 428a1418                 movb (%rax, %r11), %dl
  5331: 8ada                     movb %dl, %bl
  5333: 498b4058                 movq 0x58(%r8), %rax
  5337: 428a1418                 movb (%rax, %r11), %dl
  5341: 32da                     xorb %dl, %bl
  5343: 498b80b0000000           movq 0xb0(%r8), %rax
  5350: 428a1418                 movb (%rax, %r11), %dl
  5354: 32da                     xorb %dl, %bl
  5356: 498b80b8000000           movq 0xb8(%r8), %rax
  5363: 428a1418                 movb (%rax, %r11), %dl
  5367: 32da                     xorb %dl, %bl
  5369: 498b8018010000           movq 0x118(%r8), %rax
  5376: 428a1418                 movb (%rax, %r11), %dl
  5380: 32da                     xorb %dl, %bl
  5382: 498b8078010000           movq 0x178(%r8), %rax
  5389: 428a1418                 movb (%rax, %r11), %dl
  5393: 32da                     xorb %dl, %bl
  5395: 498b80d8010000           movq 0x1d8(%r8), %rax
  5402: 428a1418                 movb (%rax, %r11), %dl
  5406: 32da                     xorb %dl, %bl
  5408: 498b8038020000           movq 0x238(%r8), %rax
  5415: 428a1418                 movb (%rax, %r11), %dl
  5419: 32da                     xorb %dl, %bl
  5421: 498b8098020000           movq 0x298(%r8), %rax
  5428: 428a1418                 movb (%rax, %r11), %dl
  5432: 32da                     xorb %dl, %bl
  5434: 43881c1a                 movb %bl, (%r10, %r11)
  5438: 33db                     xorl %ebx, %ebx
  5440: 49ffc3                   incq %r11
  5443: 4c3bd9                   cmpq %rcx, %r11
  5446: 0f827fffffff             jb 0x14cb
  5452: 4883c408                 addq $8, %rsp
  5456: 415e                     popq %r14
  5458: 415d                     popq %r13
  5460: 415c                     popq %r12
  5462: 5b                       popq %rbx
  5463: c5f877                   vzeroupper 
  5466: c3                       retq 
  5467: 0f0b                     ud2 
  5469: 0f0b                     ud2 
  5471: 0f0b                     ud2 
  5473: 0f0b                     ud2 
  5475: 0f0b                     ud2 
  5477: 0f0b                     ud2 
  5479: 0f0b                     ud2 
  5481: 0f0b                     ud2 
  5483: 0f0b                     ud2 
  5485: 0f0b                     ud2 
  5487: 0f0b                     ud2 
  5489: 0f0b                     ud2 
  5491: 0f0b                     ud2 
  5493: 0f0b                     ud2 
  5495: 0f0b                     ud2 
  5497: 0f0b                     ud2 
  5499: 0f0b                     ud2 
  5501: 0f0b                     ud2 
  5503: cc                       int3 
