; config_t1_143 target=avx512-gfni twin=baseline topo=liberation(8)/enc op=encode U=8
; code_size=5504 twins=[] uops_per_byte=5.875
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
    45: 0f82f4090000             jb 0xa27
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 62a17d486f341e           vmovdqa32 (%rsi, %r11), %zmm22
    71: 62b17d486fc6             vmovdqa32 %zmm22, %zmm0
    77: 62317d486fde             vmovdqa32 %zmm22, %zmm11
    83: 62a17d486f341f           vmovdqa32 (%rdi, %r11), %zmm22
    90: 62b17d486fce             vmovdqa32 %zmm22, %zmm1
    96: 62317d486fe6             vmovdqa32 %zmm22, %zmm12
   102: 62817d486f341c           vmovdqa32 (%r12, %r11), %zmm22
   109: 62b17d486fd6             vmovdqa32 %zmm22, %zmm2
   115: 62317d486fee             vmovdqa32 %zmm22, %zmm13
   121: 62817d486f741d00         vmovdqa32 (%r13, %r11), %zmm22
   129: 62b17d486fde             vmovdqa32 %zmm22, %zmm3
   135: 62317d486ff6             vmovdqa32 %zmm22, %zmm14
   141: 62817d486f341e           vmovdqa32 (%r14, %r11), %zmm22
   148: 62b17d486fe6             vmovdqa32 %zmm22, %zmm4
   154: 62317d486ffe             vmovdqa32 %zmm22, %zmm15
   160: 4d8b5028                 movq 0x28(%r8), %r10
   164: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   171: 62b17d486fee             vmovdqa32 %zmm22, %zmm5
   177: 62a17d486fc6             vmovdqa32 %zmm22, %zmm16
   183: 4d8b5030                 movq 0x30(%r8), %r10
   187: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   194: 62b17d486ff6             vmovdqa32 %zmm22, %zmm6
   200: 62a17d486fce             vmovdqa32 %zmm22, %zmm17
   206: 4d8b5038                 movq 0x38(%r8), %r10
   210: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   217: 62b17d486ffe             vmovdqa32 %zmm22, %zmm7
   223: 62a17d486fd6             vmovdqa32 %zmm22, %zmm18
   229: 4d8b5040                 movq 0x40(%r8), %r10
   233: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   240: 62317d486fc6             vmovdqa32 %zmm22, %zmm8
   246: 62a17d486fde             vmovdqa32 %zmm22, %zmm19
   252: 4d8b5048                 movq 0x48(%r8), %r10
   256: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   263: 62317d486fce             vmovdqa32 %zmm22, %zmm9
   269: 62a17d486fe6             vmovdqa32 %zmm22, %zmm20
   275: 4d8b5050                 movq 0x50(%r8), %r10
   279: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   286: 62317d486fd6             vmovdqa32 %zmm22, %zmm10
   292: 62a17d486fee             vmovdqa32 %zmm22, %zmm21
   298: 4d8b5058                 movq 0x58(%r8), %r10
   302: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   309: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
   315: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
   321: 4d8b5060                 movq 0x60(%r8), %r10
   325: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   332: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
   338: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
   344: 4d8b5068                 movq 0x68(%r8), %r10
   348: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   355: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
   361: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
   367: 4d8b5070                 movq 0x70(%r8), %r10
   371: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   378: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
   384: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
   390: 4d8b5078                 movq 0x78(%r8), %r10
   394: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   401: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
   407: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
   413: 4d8b9080000000           movq 0x80(%r8), %r10
   420: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   427: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
   433: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
   439: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
   445: 4d8b9088000000           movq 0x88(%r8), %r10
   452: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   459: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
   465: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
   471: 4d8b9090000000           movq 0x90(%r8), %r10
   478: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   485: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
   491: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
   497: 4d8b9098000000           movq 0x98(%r8), %r10
   504: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   511: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
   517: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
   523: 4d8b90a0000000           movq 0xa0(%r8), %r10
   530: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   537: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
   543: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
   549: 4d8b90a8000000           movq 0xa8(%r8), %r10
   556: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   563: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
   569: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
   575: 4d8b90b0000000           movq 0xb0(%r8), %r10
   582: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   589: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
   595: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
   601: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
   607: 4d8b90b8000000           movq 0xb8(%r8), %r10
   614: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   621: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
   627: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
   633: 4d8b90c0000000           movq 0xc0(%r8), %r10
   640: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   647: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
   653: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
   659: 4d8b90c8000000           movq 0xc8(%r8), %r10
   666: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   673: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
   679: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
   685: 4d8b90d0000000           movq 0xd0(%r8), %r10
   692: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   699: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
   705: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
   711: 4d8b90d8000000           movq 0xd8(%r8), %r10
   718: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   725: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
   731: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
   737: 4d8b90e0000000           movq 0xe0(%r8), %r10
   744: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   751: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
   757: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
   763: 4d8b90e8000000           movq 0xe8(%r8), %r10
   770: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   777: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
   783: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
   789: 4d8b90f0000000           movq 0xf0(%r8), %r10
   796: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   803: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
   809: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
   815: 4d8b90f8000000           movq 0xf8(%r8), %r10
   822: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   829: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
   835: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
   841: 4d8b9000010000           movq 0x100(%r8), %r10
   848: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   855: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
   861: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
   867: 4d8b9008010000           movq 0x108(%r8), %r10
   874: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   881: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
   887: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
   893: 4d8b9010010000           movq 0x110(%r8), %r10
   900: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   907: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
   913: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
   919: 4d8b9018010000           movq 0x118(%r8), %r10
   926: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   933: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
   939: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
   945: 4d8b9020010000           movq 0x120(%r8), %r10
   952: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   959: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
   965: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
   971: 4d8b9028010000           movq 0x128(%r8), %r10
   978: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   985: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
   991: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
   997: 4d8b9030010000           movq 0x130(%r8), %r10
  1004: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1011: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  1017: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  1023: 4d8b9038010000           movq 0x138(%r8), %r10
  1030: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1037: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  1043: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  1049: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  1055: 4d8b9040010000           movq 0x140(%r8), %r10
  1062: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1069: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  1075: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  1081: 4d8b9048010000           movq 0x148(%r8), %r10
  1088: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1095: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  1101: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1107: 4d8b9050010000           movq 0x150(%r8), %r10
  1114: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1121: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  1127: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1133: 4d8b9058010000           movq 0x158(%r8), %r10
  1140: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1147: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1153: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1159: 4d8b9060010000           movq 0x160(%r8), %r10
  1166: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1173: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  1179: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1185: 4d8b9068010000           movq 0x168(%r8), %r10
  1192: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1199: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  1205: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1211: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1217: 4d8b9070010000           movq 0x170(%r8), %r10
  1224: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1231: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  1237: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1243: 4d8b9078010000           movq 0x178(%r8), %r10
  1250: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1257: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  1263: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  1269: 4d8b9080010000           movq 0x180(%r8), %r10
  1276: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1283: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  1289: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  1295: 4d8b9088010000           movq 0x188(%r8), %r10
  1302: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1309: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  1315: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1321: 4d8b9090010000           movq 0x190(%r8), %r10
  1328: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1335: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  1341: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  1347: 4d8b9098010000           movq 0x198(%r8), %r10
  1354: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1361: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  1367: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  1373: 4d8b90a0010000           movq 0x1a0(%r8), %r10
  1380: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1387: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  1393: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  1399: 4d8b90a8010000           movq 0x1a8(%r8), %r10
  1406: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1413: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  1419: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1425: 4d8b90b0010000           movq 0x1b0(%r8), %r10
  1432: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1439: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1445: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1451: 4d8b90b8010000           movq 0x1b8(%r8), %r10
  1458: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1465: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  1471: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1477: 4d8b90c0010000           movq 0x1c0(%r8), %r10
  1484: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1491: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  1497: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1503: 4d8b90c8010000           movq 0x1c8(%r8), %r10
  1510: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1517: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  1523: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1529: 4d8b90d0010000           movq 0x1d0(%r8), %r10
  1536: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1543: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  1549: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1555: 4d8b90d8010000           movq 0x1d8(%r8), %r10
  1562: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1569: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  1575: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  1581: 4d8b90e0010000           movq 0x1e0(%r8), %r10
  1588: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1595: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  1601: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  1607: 4d8b90e8010000           movq 0x1e8(%r8), %r10
  1614: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1621: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  1627: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1633: 4d8b90f0010000           movq 0x1f0(%r8), %r10
  1640: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1647: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  1653: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  1659: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  1665: 4d8b90f8010000           movq 0x1f8(%r8), %r10
  1672: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1679: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  1685: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  1691: 4d8b9000020000           movq 0x200(%r8), %r10
  1698: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1705: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  1711: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  1717: 4d8b9008020000           movq 0x208(%r8), %r10
  1724: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1731: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1737: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1743: 4d8b9010020000           movq 0x210(%r8), %r10
  1750: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1757: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  1763: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1769: 4d8b9018020000           movq 0x218(%r8), %r10
  1776: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1783: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  1789: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1795: 4d8b9020020000           movq 0x220(%r8), %r10
  1802: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1809: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  1815: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1821: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1827: 4d8b9028020000           movq 0x228(%r8), %r10
  1834: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1841: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  1847: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1853: 4d8b9030020000           movq 0x230(%r8), %r10
  1860: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1867: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  1873: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1879: 4d8b9038020000           movq 0x238(%r8), %r10
  1886: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1893: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  1899: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  1905: 4d8b9040020000           movq 0x240(%r8), %r10
  1912: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1919: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  1925: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  1931: 4d8b9048020000           movq 0x248(%r8), %r10
  1938: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1945: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  1951: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1957: 4d8b9050020000           movq 0x250(%r8), %r10
  1964: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1971: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  1977: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  1983: 4d8b9058020000           movq 0x258(%r8), %r10
  1990: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1997: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  2003: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  2009: 4d8b9060020000           movq 0x260(%r8), %r10
  2016: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2023: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  2029: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  2035: 4d8b9068020000           movq 0x268(%r8), %r10
  2042: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2049: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  2055: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  2061: 4d8b9070020000           movq 0x270(%r8), %r10
  2068: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2075: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  2081: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  2087: 4d8b9078020000           movq 0x278(%r8), %r10
  2094: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2101: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  2107: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  2113: 4d8b9080020000           movq 0x280(%r8), %r10
  2120: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2127: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  2133: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  2139: 4d8b9088020000           movq 0x288(%r8), %r10
  2146: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2153: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  2159: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  2165: 4d8b9090020000           movq 0x290(%r8), %r10
  2172: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2179: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  2185: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  2191: 4d8b9098020000           movq 0x298(%r8), %r10
  2198: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2205: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  2211: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  2217: 4d8b90a0020000           movq 0x2a0(%r8), %r10
  2224: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2231: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  2237: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  2243: 4d8b90a8020000           movq 0x2a8(%r8), %r10
  2250: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2257: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  2263: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  2269: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  2275: 4d8b90b0020000           movq 0x2b0(%r8), %r10
  2282: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2289: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  2295: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  2301: 4d8b90b8020000           movq 0x2b8(%r8), %r10
  2308: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  2315: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  2321: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  2327: 4d8b11                   movq (%r9), %r10
  2330: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  2337: 4d8b5108                 movq 8(%r9), %r10
  2341: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
  2348: 4d8b5110                 movq 0x10(%r9), %r10
  2352: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
  2359: 4d8b5118                 movq 0x18(%r9), %r10
  2363: 62917d487f1c1a           vmovdqa32 %zmm3, (%r10, %r11)
  2370: 4d8b5120                 movq 0x20(%r9), %r10
  2374: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
  2381: 4d8b5128                 movq 0x28(%r9), %r10
  2385: 62917d487f2c1a           vmovdqa32 %zmm5, (%r10, %r11)
  2392: 4d8b5130                 movq 0x30(%r9), %r10
  2396: 62917d487f341a           vmovdqa32 %zmm6, (%r10, %r11)
  2403: 4d8b5138                 movq 0x38(%r9), %r10
  2407: 62917d487f3c1a           vmovdqa32 %zmm7, (%r10, %r11)
  2414: 4d8b5140                 movq 0x40(%r9), %r10
  2418: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  2425: 4d8b5148                 movq 0x48(%r9), %r10
  2429: 62117d487f0c1a           vmovdqa32 %zmm9, (%r10, %r11)
  2436: 4d8b5150                 movq 0x50(%r9), %r10
  2440: 62117d487f141a           vmovdqa32 %zmm10, (%r10, %r11)
  2447: 4d8b5158                 movq 0x58(%r9), %r10
  2451: 62117d487f1c1a           vmovdqa32 %zmm11, (%r10, %r11)
  2458: 4d8b5160                 movq 0x60(%r9), %r10
  2462: 62117d487f241a           vmovdqa32 %zmm12, (%r10, %r11)
  2469: 4d8b5168                 movq 0x68(%r9), %r10
  2473: 62117d487f2c1a           vmovdqa32 %zmm13, (%r10, %r11)
  2480: 4d8b5170                 movq 0x70(%r9), %r10
  2484: 62117d487f341a           vmovdqa32 %zmm14, (%r10, %r11)
  2491: 4d8b5178                 movq 0x78(%r9), %r10
  2495: 62117d487f3c1a           vmovdqa32 %zmm15, (%r10, %r11)
  2502: 4d8b9180000000           movq 0x80(%r9), %r10
  2509: 62817d487f041a           vmovdqa32 %zmm16, (%r10, %r11)
  2516: 4d8b9188000000           movq 0x88(%r9), %r10
  2523: 62817d487f0c1a           vmovdqa32 %zmm17, (%r10, %r11)
  2530: 4d8b9190000000           movq 0x90(%r9), %r10
  2537: 62817d487f141a           vmovdqa32 %zmm18, (%r10, %r11)
  2544: 4d8b9198000000           movq 0x98(%r9), %r10
  2551: 62817d487f1c1a           vmovdqa32 %zmm19, (%r10, %r11)
  2558: 4d8b91a0000000           movq 0xa0(%r9), %r10
  2565: 62817d487f241a           vmovdqa32 %zmm20, (%r10, %r11)
  2572: 4d8b91a8000000           movq 0xa8(%r9), %r10
  2579: 62817d487f2c1a           vmovdqa32 %zmm21, (%r10, %r11)
  2586: 4983c340                 addq $0x40, %r11
  2590: 4c3bd8                   cmpq %rax, %r11
  2593: 0f8619f6ffff             jbe 0x40
  2599: 4c3bd9                   cmpq %rcx, %r11
  2602: 0f833c0b0000             jae 0x156c
  2608: 4153                     pushq %r11
  2610: 4d8b11                   movq (%r9), %r10
  2613: 4c8b1c24                 movq (%rsp), %r11
  2617: 498b00                   movq (%r8), %rax
  2620: 428a1418                 movb (%rax, %r11), %dl
  2624: 8ada                     movb %dl, %bl
  2626: 498b4058                 movq 0x58(%r8), %rax
  2630: 428a1418                 movb (%rax, %r11), %dl
  2634: 32da                     xorb %dl, %bl
  2636: 498b80b0000000           movq 0xb0(%r8), %rax
  2643: 428a1418                 movb (%rax, %r11), %dl
  2647: 32da                     xorb %dl, %bl
  2649: 498b8008010000           movq 0x108(%r8), %rax
  2656: 428a1418                 movb (%rax, %r11), %dl
  2660: 32da                     xorb %dl, %bl
  2662: 498b8060010000           movq 0x160(%r8), %rax
  2669: 428a1418                 movb (%rax, %r11), %dl
  2673: 32da                     xorb %dl, %bl
  2675: 498b80b8010000           movq 0x1b8(%r8), %rax
  2682: 428a1418                 movb (%rax, %r11), %dl
  2686: 32da                     xorb %dl, %bl
  2688: 498b8010020000           movq 0x210(%r8), %rax
  2695: 428a1418                 movb (%rax, %r11), %dl
  2699: 32da                     xorb %dl, %bl
  2701: 498b8068020000           movq 0x268(%r8), %rax
  2708: 428a1418                 movb (%rax, %r11), %dl
  2712: 32da                     xorb %dl, %bl
  2714: 43881c1a                 movb %bl, (%r10, %r11)
  2718: 33db                     xorl %ebx, %ebx
  2720: 49ffc3                   incq %r11
  2723: 4c3bd9                   cmpq %rcx, %r11
  2726: 0f828dffffff             jb 0xa39
  2732: 4d8b5108                 movq 8(%r9), %r10
  2736: 4c8b1c24                 movq (%rsp), %r11
  2740: 498b4008                 movq 8(%r8), %rax
  2744: 428a1418                 movb (%rax, %r11), %dl
  2748: 8ada                     movb %dl, %bl
  2750: 498b4060                 movq 0x60(%r8), %rax
  2754: 428a1418                 movb (%rax, %r11), %dl
  2758: 32da                     xorb %dl, %bl
  2760: 498b80b8000000           movq 0xb8(%r8), %rax
  2767: 428a1418                 movb (%rax, %r11), %dl
  2771: 32da                     xorb %dl, %bl
  2773: 498b8010010000           movq 0x110(%r8), %rax
  2780: 428a1418                 movb (%rax, %r11), %dl
  2784: 32da                     xorb %dl, %bl
  2786: 498b8068010000           movq 0x168(%r8), %rax
  2793: 428a1418                 movb (%rax, %r11), %dl
  2797: 32da                     xorb %dl, %bl
  2799: 498b80c0010000           movq 0x1c0(%r8), %rax
  2806: 428a1418                 movb (%rax, %r11), %dl
  2810: 32da                     xorb %dl, %bl
  2812: 498b8018020000           movq 0x218(%r8), %rax
  2819: 428a1418                 movb (%rax, %r11), %dl
  2823: 32da                     xorb %dl, %bl
  2825: 498b8070020000           movq 0x270(%r8), %rax
  2832: 428a1418                 movb (%rax, %r11), %dl
  2836: 32da                     xorb %dl, %bl
  2838: 43881c1a                 movb %bl, (%r10, %r11)
  2842: 33db                     xorl %ebx, %ebx
  2844: 49ffc3                   incq %r11
  2847: 4c3bd9                   cmpq %rcx, %r11
  2850: 0f828cffffff             jb 0xab4
  2856: 4d8b5110                 movq 0x10(%r9), %r10
  2860: 4c8b1c24                 movq (%rsp), %r11
  2864: 498b4010                 movq 0x10(%r8), %rax
  2868: 428a1418                 movb (%rax, %r11), %dl
  2872: 8ada                     movb %dl, %bl
  2874: 498b4068                 movq 0x68(%r8), %rax
  2878: 428a1418                 movb (%rax, %r11), %dl
  2882: 32da                     xorb %dl, %bl
  2884: 498b80c0000000           movq 0xc0(%r8), %rax
  2891: 428a1418                 movb (%rax, %r11), %dl
  2895: 32da                     xorb %dl, %bl
  2897: 498b8018010000           movq 0x118(%r8), %rax
  2904: 428a1418                 movb (%rax, %r11), %dl
  2908: 32da                     xorb %dl, %bl
  2910: 498b8070010000           movq 0x170(%r8), %rax
  2917: 428a1418                 movb (%rax, %r11), %dl
  2921: 32da                     xorb %dl, %bl
  2923: 498b80c8010000           movq 0x1c8(%r8), %rax
  2930: 428a1418                 movb (%rax, %r11), %dl
  2934: 32da                     xorb %dl, %bl
  2936: 498b8020020000           movq 0x220(%r8), %rax
  2943: 428a1418                 movb (%rax, %r11), %dl
  2947: 32da                     xorb %dl, %bl
  2949: 498b8078020000           movq 0x278(%r8), %rax
  2956: 428a1418                 movb (%rax, %r11), %dl
  2960: 32da                     xorb %dl, %bl
  2962: 43881c1a                 movb %bl, (%r10, %r11)
  2966: 33db                     xorl %ebx, %ebx
  2968: 49ffc3                   incq %r11
  2971: 4c3bd9                   cmpq %rcx, %r11
  2974: 0f828cffffff             jb 0xb30
  2980: 4d8b5118                 movq 0x18(%r9), %r10
  2984: 4c8b1c24                 movq (%rsp), %r11
  2988: 498b4018                 movq 0x18(%r8), %rax
  2992: 428a1418                 movb (%rax, %r11), %dl
  2996: 8ada                     movb %dl, %bl
  2998: 498b4070                 movq 0x70(%r8), %rax
  3002: 428a1418                 movb (%rax, %r11), %dl
  3006: 32da                     xorb %dl, %bl
  3008: 498b80c8000000           movq 0xc8(%r8), %rax
  3015: 428a1418                 movb (%rax, %r11), %dl
  3019: 32da                     xorb %dl, %bl
  3021: 498b8020010000           movq 0x120(%r8), %rax
  3028: 428a1418                 movb (%rax, %r11), %dl
  3032: 32da                     xorb %dl, %bl
  3034: 498b8078010000           movq 0x178(%r8), %rax
  3041: 428a1418                 movb (%rax, %r11), %dl
  3045: 32da                     xorb %dl, %bl
  3047: 498b80d0010000           movq 0x1d0(%r8), %rax
  3054: 428a1418                 movb (%rax, %r11), %dl
  3058: 32da                     xorb %dl, %bl
  3060: 498b8028020000           movq 0x228(%r8), %rax
  3067: 428a1418                 movb (%rax, %r11), %dl
  3071: 32da                     xorb %dl, %bl
  3073: 498b8080020000           movq 0x280(%r8), %rax
  3080: 428a1418                 movb (%rax, %r11), %dl
  3084: 32da                     xorb %dl, %bl
  3086: 43881c1a                 movb %bl, (%r10, %r11)
  3090: 33db                     xorl %ebx, %ebx
  3092: 49ffc3                   incq %r11
  3095: 4c3bd9                   cmpq %rcx, %r11
  3098: 0f828cffffff             jb 0xbac
  3104: 4d8b5120                 movq 0x20(%r9), %r10
  3108: 4c8b1c24                 movq (%rsp), %r11
  3112: 498b4020                 movq 0x20(%r8), %rax
  3116: 428a1418                 movb (%rax, %r11), %dl
  3120: 8ada                     movb %dl, %bl
  3122: 498b4078                 movq 0x78(%r8), %rax
  3126: 428a1418                 movb (%rax, %r11), %dl
  3130: 32da                     xorb %dl, %bl
  3132: 498b80d0000000           movq 0xd0(%r8), %rax
  3139: 428a1418                 movb (%rax, %r11), %dl
  3143: 32da                     xorb %dl, %bl
  3145: 498b8028010000           movq 0x128(%r8), %rax
  3152: 428a1418                 movb (%rax, %r11), %dl
  3156: 32da                     xorb %dl, %bl
  3158: 498b8080010000           movq 0x180(%r8), %rax
  3165: 428a1418                 movb (%rax, %r11), %dl
  3169: 32da                     xorb %dl, %bl
  3171: 498b80d8010000           movq 0x1d8(%r8), %rax
  3178: 428a1418                 movb (%rax, %r11), %dl
  3182: 32da                     xorb %dl, %bl
  3184: 498b8030020000           movq 0x230(%r8), %rax
  3191: 428a1418                 movb (%rax, %r11), %dl
  3195: 32da                     xorb %dl, %bl
  3197: 498b8088020000           movq 0x288(%r8), %rax
  3204: 428a1418                 movb (%rax, %r11), %dl
  3208: 32da                     xorb %dl, %bl
  3210: 43881c1a                 movb %bl, (%r10, %r11)
  3214: 33db                     xorl %ebx, %ebx
  3216: 49ffc3                   incq %r11
  3219: 4c3bd9                   cmpq %rcx, %r11
  3222: 0f828cffffff             jb 0xc28
  3228: 4d8b5128                 movq 0x28(%r9), %r10
  3232: 4c8b1c24                 movq (%rsp), %r11
  3236: 498b4028                 movq 0x28(%r8), %rax
  3240: 428a1418                 movb (%rax, %r11), %dl
  3244: 8ada                     movb %dl, %bl
  3246: 498b8080000000           movq 0x80(%r8), %rax
  3253: 428a1418                 movb (%rax, %r11), %dl
  3257: 32da                     xorb %dl, %bl
  3259: 498b80d8000000           movq 0xd8(%r8), %rax
  3266: 428a1418                 movb (%rax, %r11), %dl
  3270: 32da                     xorb %dl, %bl
  3272: 498b8030010000           movq 0x130(%r8), %rax
  3279: 428a1418                 movb (%rax, %r11), %dl
  3283: 32da                     xorb %dl, %bl
  3285: 498b8088010000           movq 0x188(%r8), %rax
  3292: 428a1418                 movb (%rax, %r11), %dl
  3296: 32da                     xorb %dl, %bl
  3298: 498b80e0010000           movq 0x1e0(%r8), %rax
  3305: 428a1418                 movb (%rax, %r11), %dl
  3309: 32da                     xorb %dl, %bl
  3311: 498b8038020000           movq 0x238(%r8), %rax
  3318: 428a1418                 movb (%rax, %r11), %dl
  3322: 32da                     xorb %dl, %bl
  3324: 498b8090020000           movq 0x290(%r8), %rax
  3331: 428a1418                 movb (%rax, %r11), %dl
  3335: 32da                     xorb %dl, %bl
  3337: 43881c1a                 movb %bl, (%r10, %r11)
  3341: 33db                     xorl %ebx, %ebx
  3343: 49ffc3                   incq %r11
  3346: 4c3bd9                   cmpq %rcx, %r11
  3349: 0f8289ffffff             jb 0xca4
  3355: 4d8b5130                 movq 0x30(%r9), %r10
  3359: 4c8b1c24                 movq (%rsp), %r11
  3363: 498b4030                 movq 0x30(%r8), %rax
  3367: 428a1418                 movb (%rax, %r11), %dl
  3371: 8ada                     movb %dl, %bl
  3373: 498b8088000000           movq 0x88(%r8), %rax
  3380: 428a1418                 movb (%rax, %r11), %dl
  3384: 32da                     xorb %dl, %bl
  3386: 498b80e0000000           movq 0xe0(%r8), %rax
  3393: 428a1418                 movb (%rax, %r11), %dl
  3397: 32da                     xorb %dl, %bl
  3399: 498b8038010000           movq 0x138(%r8), %rax
  3406: 428a1418                 movb (%rax, %r11), %dl
  3410: 32da                     xorb %dl, %bl
  3412: 498b8090010000           movq 0x190(%r8), %rax
  3419: 428a1418                 movb (%rax, %r11), %dl
  3423: 32da                     xorb %dl, %bl
  3425: 498b80e8010000           movq 0x1e8(%r8), %rax
  3432: 428a1418                 movb (%rax, %r11), %dl
  3436: 32da                     xorb %dl, %bl
  3438: 498b8040020000           movq 0x240(%r8), %rax
  3445: 428a1418                 movb (%rax, %r11), %dl
  3449: 32da                     xorb %dl, %bl
  3451: 498b8098020000           movq 0x298(%r8), %rax
  3458: 428a1418                 movb (%rax, %r11), %dl
  3462: 32da                     xorb %dl, %bl
  3464: 43881c1a                 movb %bl, (%r10, %r11)
  3468: 33db                     xorl %ebx, %ebx
  3470: 49ffc3                   incq %r11
  3473: 4c3bd9                   cmpq %rcx, %r11
  3476: 0f8289ffffff             jb 0xd23
  3482: 4d8b5138                 movq 0x38(%r9), %r10
  3486: 4c8b1c24                 movq (%rsp), %r11
  3490: 498b4038                 movq 0x38(%r8), %rax
  3494: 428a1418                 movb (%rax, %r11), %dl
  3498: 8ada                     movb %dl, %bl
  3500: 498b8090000000           movq 0x90(%r8), %rax
  3507: 428a1418                 movb (%rax, %r11), %dl
  3511: 32da                     xorb %dl, %bl
  3513: 498b80e8000000           movq 0xe8(%r8), %rax
  3520: 428a1418                 movb (%rax, %r11), %dl
  3524: 32da                     xorb %dl, %bl
  3526: 498b8040010000           movq 0x140(%r8), %rax
  3533: 428a1418                 movb (%rax, %r11), %dl
  3537: 32da                     xorb %dl, %bl
  3539: 498b8098010000           movq 0x198(%r8), %rax
  3546: 428a1418                 movb (%rax, %r11), %dl
  3550: 32da                     xorb %dl, %bl
  3552: 498b80f0010000           movq 0x1f0(%r8), %rax
  3559: 428a1418                 movb (%rax, %r11), %dl
  3563: 32da                     xorb %dl, %bl
  3565: 498b8048020000           movq 0x248(%r8), %rax
  3572: 428a1418                 movb (%rax, %r11), %dl
  3576: 32da                     xorb %dl, %bl
  3578: 498b80a0020000           movq 0x2a0(%r8), %rax
  3585: 428a1418                 movb (%rax, %r11), %dl
  3589: 32da                     xorb %dl, %bl
  3591: 43881c1a                 movb %bl, (%r10, %r11)
  3595: 33db                     xorl %ebx, %ebx
  3597: 49ffc3                   incq %r11
  3600: 4c3bd9                   cmpq %rcx, %r11
  3603: 0f8289ffffff             jb 0xda2
  3609: 4d8b5140                 movq 0x40(%r9), %r10
  3613: 4c8b1c24                 movq (%rsp), %r11
  3617: 498b4040                 movq 0x40(%r8), %rax
  3621: 428a1418                 movb (%rax, %r11), %dl
  3625: 8ada                     movb %dl, %bl
  3627: 498b8098000000           movq 0x98(%r8), %rax
  3634: 428a1418                 movb (%rax, %r11), %dl
  3638: 32da                     xorb %dl, %bl
  3640: 498b80f0000000           movq 0xf0(%r8), %rax
  3647: 428a1418                 movb (%rax, %r11), %dl
  3651: 32da                     xorb %dl, %bl
  3653: 498b8048010000           movq 0x148(%r8), %rax
  3660: 428a1418                 movb (%rax, %r11), %dl
  3664: 32da                     xorb %dl, %bl
  3666: 498b80a0010000           movq 0x1a0(%r8), %rax
  3673: 428a1418                 movb (%rax, %r11), %dl
  3677: 32da                     xorb %dl, %bl
  3679: 498b80f8010000           movq 0x1f8(%r8), %rax
  3686: 428a1418                 movb (%rax, %r11), %dl
  3690: 32da                     xorb %dl, %bl
  3692: 498b8050020000           movq 0x250(%r8), %rax
  3699: 428a1418                 movb (%rax, %r11), %dl
  3703: 32da                     xorb %dl, %bl
  3705: 498b80a8020000           movq 0x2a8(%r8), %rax
  3712: 428a1418                 movb (%rax, %r11), %dl
  3716: 32da                     xorb %dl, %bl
  3718: 43881c1a                 movb %bl, (%r10, %r11)
  3722: 33db                     xorl %ebx, %ebx
  3724: 49ffc3                   incq %r11
  3727: 4c3bd9                   cmpq %rcx, %r11
  3730: 0f8289ffffff             jb 0xe21
  3736: 4d8b5148                 movq 0x48(%r9), %r10
  3740: 4c8b1c24                 movq (%rsp), %r11
  3744: 498b4048                 movq 0x48(%r8), %rax
  3748: 428a1418                 movb (%rax, %r11), %dl
  3752: 8ada                     movb %dl, %bl
  3754: 498b80a0000000           movq 0xa0(%r8), %rax
  3761: 428a1418                 movb (%rax, %r11), %dl
  3765: 32da                     xorb %dl, %bl
  3767: 498b80f8000000           movq 0xf8(%r8), %rax
  3774: 428a1418                 movb (%rax, %r11), %dl
  3778: 32da                     xorb %dl, %bl
  3780: 498b8050010000           movq 0x150(%r8), %rax
  3787: 428a1418                 movb (%rax, %r11), %dl
  3791: 32da                     xorb %dl, %bl
  3793: 498b80a8010000           movq 0x1a8(%r8), %rax
  3800: 428a1418                 movb (%rax, %r11), %dl
  3804: 32da                     xorb %dl, %bl
  3806: 498b8000020000           movq 0x200(%r8), %rax
  3813: 428a1418                 movb (%rax, %r11), %dl
  3817: 32da                     xorb %dl, %bl
  3819: 498b8058020000           movq 0x258(%r8), %rax
  3826: 428a1418                 movb (%rax, %r11), %dl
  3830: 32da                     xorb %dl, %bl
  3832: 498b80b0020000           movq 0x2b0(%r8), %rax
  3839: 428a1418                 movb (%rax, %r11), %dl
  3843: 32da                     xorb %dl, %bl
  3845: 43881c1a                 movb %bl, (%r10, %r11)
  3849: 33db                     xorl %ebx, %ebx
  3851: 49ffc3                   incq %r11
  3854: 4c3bd9                   cmpq %rcx, %r11
  3857: 0f8289ffffff             jb 0xea0
  3863: 4d8b5150                 movq 0x50(%r9), %r10
  3867: 4c8b1c24                 movq (%rsp), %r11
  3871: 498b4050                 movq 0x50(%r8), %rax
  3875: 428a1418                 movb (%rax, %r11), %dl
  3879: 8ada                     movb %dl, %bl
  3881: 498b80a8000000           movq 0xa8(%r8), %rax
  3888: 428a1418                 movb (%rax, %r11), %dl
  3892: 32da                     xorb %dl, %bl
  3894: 498b8000010000           movq 0x100(%r8), %rax
  3901: 428a1418                 movb (%rax, %r11), %dl
  3905: 32da                     xorb %dl, %bl
  3907: 498b8058010000           movq 0x158(%r8), %rax
  3914: 428a1418                 movb (%rax, %r11), %dl
  3918: 32da                     xorb %dl, %bl
  3920: 498b80b0010000           movq 0x1b0(%r8), %rax
  3927: 428a1418                 movb (%rax, %r11), %dl
  3931: 32da                     xorb %dl, %bl
  3933: 498b8008020000           movq 0x208(%r8), %rax
  3940: 428a1418                 movb (%rax, %r11), %dl
  3944: 32da                     xorb %dl, %bl
  3946: 498b8060020000           movq 0x260(%r8), %rax
  3953: 428a1418                 movb (%rax, %r11), %dl
  3957: 32da                     xorb %dl, %bl
  3959: 498b80b8020000           movq 0x2b8(%r8), %rax
  3966: 428a1418                 movb (%rax, %r11), %dl
  3970: 32da                     xorb %dl, %bl
  3972: 43881c1a                 movb %bl, (%r10, %r11)
  3976: 33db                     xorl %ebx, %ebx
  3978: 49ffc3                   incq %r11
  3981: 4c3bd9                   cmpq %rcx, %r11
  3984: 0f8289ffffff             jb 0xf1f
  3990: 4d8b5158                 movq 0x58(%r9), %r10
  3994: 4c8b1c24                 movq (%rsp), %r11
  3998: 498b00                   movq (%r8), %rax
  4001: 428a1418                 movb (%rax, %r11), %dl
  4005: 8ada                     movb %dl, %bl
  4007: 498b4060                 movq 0x60(%r8), %rax
  4011: 428a1418                 movb (%rax, %r11), %dl
  4015: 32da                     xorb %dl, %bl
  4017: 498b80c0000000           movq 0xc0(%r8), %rax
  4024: 428a1418                 movb (%rax, %r11), %dl
  4028: 32da                     xorb %dl, %bl
  4030: 498b8020010000           movq 0x120(%r8), %rax
  4037: 428a1418                 movb (%rax, %r11), %dl
  4041: 32da                     xorb %dl, %bl
  4043: 498b8080010000           movq 0x180(%r8), %rax
  4050: 428a1418                 movb (%rax, %r11), %dl
  4054: 32da                     xorb %dl, %bl
  4056: 498b80e0010000           movq 0x1e0(%r8), %rax
  4063: 428a1418                 movb (%rax, %r11), %dl
  4067: 32da                     xorb %dl, %bl
  4069: 498b8040020000           movq 0x240(%r8), %rax
  4076: 428a1418                 movb (%rax, %r11), %dl
  4080: 32da                     xorb %dl, %bl
  4082: 498b80a0020000           movq 0x2a0(%r8), %rax
  4089: 428a1418                 movb (%rax, %r11), %dl
  4093: 32da                     xorb %dl, %bl
  4095: 43881c1a                 movb %bl, (%r10, %r11)
  4099: 33db                     xorl %ebx, %ebx
  4101: 49ffc3                   incq %r11
  4104: 4c3bd9                   cmpq %rcx, %r11
  4107: 0f828dffffff             jb 0xf9e
  4113: 4d8b5160                 movq 0x60(%r9), %r10
  4117: 4c8b1c24                 movq (%rsp), %r11
  4121: 498b4008                 movq 8(%r8), %rax
  4125: 428a1418                 movb (%rax, %r11), %dl
  4129: 8ada                     movb %dl, %bl
  4131: 498b4068                 movq 0x68(%r8), %rax
  4135: 428a1418                 movb (%rax, %r11), %dl
  4139: 32da                     xorb %dl, %bl
  4141: 498b80c8000000           movq 0xc8(%r8), %rax
  4148: 428a1418                 movb (%rax, %r11), %dl
  4152: 32da                     xorb %dl, %bl
  4154: 498b8028010000           movq 0x128(%r8), %rax
  4161: 428a1418                 movb (%rax, %r11), %dl
  4165: 32da                     xorb %dl, %bl
  4167: 498b8088010000           movq 0x188(%r8), %rax
  4174: 428a1418                 movb (%rax, %r11), %dl
  4178: 32da                     xorb %dl, %bl
  4180: 498b80e8010000           movq 0x1e8(%r8), %rax
  4187: 428a1418                 movb (%rax, %r11), %dl
  4191: 32da                     xorb %dl, %bl
  4193: 498b8048020000           movq 0x248(%r8), %rax
  4200: 428a1418                 movb (%rax, %r11), %dl
  4204: 32da                     xorb %dl, %bl
  4206: 498b80a8020000           movq 0x2a8(%r8), %rax
  4213: 428a1418                 movb (%rax, %r11), %dl
  4217: 32da                     xorb %dl, %bl
  4219: 43881c1a                 movb %bl, (%r10, %r11)
  4223: 33db                     xorl %ebx, %ebx
  4225: 49ffc3                   incq %r11
  4228: 4c3bd9                   cmpq %rcx, %r11
  4231: 0f828cffffff             jb 0x1019
  4237: 4d8b5168                 movq 0x68(%r9), %r10
  4241: 4c8b1c24                 movq (%rsp), %r11
  4245: 498b4010                 movq 0x10(%r8), %rax
  4249: 428a1418                 movb (%rax, %r11), %dl
  4253: 8ada                     movb %dl, %bl
  4255: 498b4070                 movq 0x70(%r8), %rax
  4259: 428a1418                 movb (%rax, %r11), %dl
  4263: 32da                     xorb %dl, %bl
  4265: 498b80d0000000           movq 0xd0(%r8), %rax
  4272: 428a1418                 movb (%rax, %r11), %dl
  4276: 32da                     xorb %dl, %bl
  4278: 498b8030010000           movq 0x130(%r8), %rax
  4285: 428a1418                 movb (%rax, %r11), %dl
  4289: 32da                     xorb %dl, %bl
  4291: 498b8090010000           movq 0x190(%r8), %rax
  4298: 428a1418                 movb (%rax, %r11), %dl
  4302: 32da                     xorb %dl, %bl
  4304: 498b80f0010000           movq 0x1f0(%r8), %rax
  4311: 428a1418                 movb (%rax, %r11), %dl
  4315: 32da                     xorb %dl, %bl
  4317: 498b8050020000           movq 0x250(%r8), %rax
  4324: 428a1418                 movb (%rax, %r11), %dl
  4328: 32da                     xorb %dl, %bl
  4330: 498b80a8020000           movq 0x2a8(%r8), %rax
  4337: 428a1418                 movb (%rax, %r11), %dl
  4341: 32da                     xorb %dl, %bl
  4343: 498b80b0020000           movq 0x2b0(%r8), %rax
  4350: 428a1418                 movb (%rax, %r11), %dl
  4354: 32da                     xorb %dl, %bl
  4356: 43881c1a                 movb %bl, (%r10, %r11)
  4360: 33db                     xorl %ebx, %ebx
  4362: 49ffc3                   incq %r11
  4365: 4c3bd9                   cmpq %rcx, %r11
  4368: 0f827fffffff             jb 0x1095
  4374: 4d8b5170                 movq 0x70(%r9), %r10
  4378: 4c8b1c24                 movq (%rsp), %r11
  4382: 498b4018                 movq 0x18(%r8), %rax
  4386: 428a1418                 movb (%rax, %r11), %dl
  4390: 8ada                     movb %dl, %bl
  4392: 498b4078                 movq 0x78(%r8), %rax
  4396: 428a1418                 movb (%rax, %r11), %dl
  4400: 32da                     xorb %dl, %bl
  4402: 498b80d8000000           movq 0xd8(%r8), %rax
  4409: 428a1418                 movb (%rax, %r11), %dl
  4413: 32da                     xorb %dl, %bl
  4415: 498b8038010000           movq 0x138(%r8), %rax
  4422: 428a1418                 movb (%rax, %r11), %dl
  4426: 32da                     xorb %dl, %bl
  4428: 498b8098010000           movq 0x198(%r8), %rax
  4435: 428a1418                 movb (%rax, %r11), %dl
  4439: 32da                     xorb %dl, %bl
  4441: 498b80f0010000           movq 0x1f0(%r8), %rax
  4448: 428a1418                 movb (%rax, %r11), %dl
  4452: 32da                     xorb %dl, %bl
  4454: 498b80f8010000           movq 0x1f8(%r8), %rax
  4461: 428a1418                 movb (%rax, %r11), %dl
  4465: 32da                     xorb %dl, %bl
  4467: 498b8058020000           movq 0x258(%r8), %rax
  4474: 428a1418                 movb (%rax, %r11), %dl
  4478: 32da                     xorb %dl, %bl
  4480: 498b80b8020000           movq 0x2b8(%r8), %rax
  4487: 428a1418                 movb (%rax, %r11), %dl
  4491: 32da                     xorb %dl, %bl
  4493: 43881c1a                 movb %bl, (%r10, %r11)
  4497: 33db                     xorl %ebx, %ebx
  4499: 49ffc3                   incq %r11
  4502: 4c3bd9                   cmpq %rcx, %r11
  4505: 0f827fffffff             jb 0x111e
  4511: 4d8b5178                 movq 0x78(%r9), %r10
  4515: 4c8b1c24                 movq (%rsp), %r11
  4519: 498b4020                 movq 0x20(%r8), %rax
  4523: 428a1418                 movb (%rax, %r11), %dl
  4527: 8ada                     movb %dl, %bl
  4529: 498b8080000000           movq 0x80(%r8), %rax
  4536: 428a1418                 movb (%rax, %r11), %dl
  4540: 32da                     xorb %dl, %bl
  4542: 498b80e0000000           movq 0xe0(%r8), %rax
  4549: 428a1418                 movb (%rax, %r11), %dl
  4553: 32da                     xorb %dl, %bl
  4555: 498b8038010000           movq 0x138(%r8), %rax
  4562: 428a1418                 movb (%rax, %r11), %dl
  4566: 32da                     xorb %dl, %bl
  4568: 498b8040010000           movq 0x140(%r8), %rax
  4575: 428a1418                 movb (%rax, %r11), %dl
  4579: 32da                     xorb %dl, %bl
  4581: 498b80a0010000           movq 0x1a0(%r8), %rax
  4588: 428a1418                 movb (%rax, %r11), %dl
  4592: 32da                     xorb %dl, %bl
  4594: 498b8000020000           movq 0x200(%r8), %rax
  4601: 428a1418                 movb (%rax, %r11), %dl
  4605: 32da                     xorb %dl, %bl
  4607: 498b8060020000           movq 0x260(%r8), %rax
  4614: 428a1418                 movb (%rax, %r11), %dl
  4618: 32da                     xorb %dl, %bl
  4620: 498b8068020000           movq 0x268(%r8), %rax
  4627: 428a1418                 movb (%rax, %r11), %dl
  4631: 32da                     xorb %dl, %bl
  4633: 43881c1a                 movb %bl, (%r10, %r11)
  4637: 33db                     xorl %ebx, %ebx
  4639: 49ffc3                   incq %r11
  4642: 4c3bd9                   cmpq %rcx, %r11
  4645: 0f827cffffff             jb 0x11a7
  4651: 4d8b9180000000           movq 0x80(%r9), %r10
  4658: 4c8b1c24                 movq (%rsp), %r11
  4662: 498b4028                 movq 0x28(%r8), %rax
  4666: 428a1418                 movb (%rax, %r11), %dl
  4670: 8ada                     movb %dl, %bl
  4672: 498b8080000000           movq 0x80(%r8), %rax
  4679: 428a1418                 movb (%rax, %r11), %dl
  4683: 32da                     xorb %dl, %bl
  4685: 498b8088000000           movq 0x88(%r8), %rax
  4692: 428a1418                 movb (%rax, %r11), %dl
  4696: 32da                     xorb %dl, %bl
  4698: 498b80e8000000           movq 0xe8(%r8), %rax
  4705: 428a1418                 movb (%rax, %r11), %dl
  4709: 32da                     xorb %dl, %bl
  4711: 498b8048010000           movq 0x148(%r8), %rax
  4718: 428a1418                 movb (%rax, %r11), %dl
  4722: 32da                     xorb %dl, %bl
  4724: 498b80a8010000           movq 0x1a8(%r8), %rax
  4731: 428a1418                 movb (%rax, %r11), %dl
  4735: 32da                     xorb %dl, %bl
  4737: 498b8008020000           movq 0x208(%r8), %rax
  4744: 428a1418                 movb (%rax, %r11), %dl
  4748: 32da                     xorb %dl, %bl
  4750: 498b8010020000           movq 0x210(%r8), %rax
  4757: 428a1418                 movb (%rax, %r11), %dl
  4761: 32da                     xorb %dl, %bl
  4763: 498b8070020000           movq 0x270(%r8), %rax
  4770: 428a1418                 movb (%rax, %r11), %dl
  4774: 32da                     xorb %dl, %bl
  4776: 43881c1a                 movb %bl, (%r10, %r11)
  4780: 33db                     xorl %ebx, %ebx
  4782: 49ffc3                   incq %r11
  4785: 4c3bd9                   cmpq %rcx, %r11
  4788: 0f827cffffff             jb 0x1236
  4794: 4d8b9188000000           movq 0x88(%r9), %r10
  4801: 4c8b1c24                 movq (%rsp), %r11
  4805: 498b4030                 movq 0x30(%r8), %rax
  4809: 428a1418                 movb (%rax, %r11), %dl
  4813: 8ada                     movb %dl, %bl
  4815: 498b8090000000           movq 0x90(%r8), %rax
  4822: 428a1418                 movb (%rax, %r11), %dl
  4826: 32da                     xorb %dl, %bl
  4828: 498b80f0000000           movq 0xf0(%r8), %rax
  4835: 428a1418                 movb (%rax, %r11), %dl
  4839: 32da                     xorb %dl, %bl
  4841: 498b8050010000           movq 0x150(%r8), %rax
  4848: 428a1418                 movb (%rax, %r11), %dl
  4852: 32da                     xorb %dl, %bl
  4854: 498b80b0010000           movq 0x1b0(%r8), %rax
  4861: 428a1418                 movb (%rax, %r11), %dl
  4865: 32da                     xorb %dl, %bl
  4867: 498b80b8010000           movq 0x1b8(%r8), %rax
  4874: 428a1418                 movb (%rax, %r11), %dl
  4878: 32da                     xorb %dl, %bl
  4880: 498b8018020000           movq 0x218(%r8), %rax
  4887: 428a1418                 movb (%rax, %r11), %dl
  4891: 32da                     xorb %dl, %bl
  4893: 498b8078020000           movq 0x278(%r8), %rax
  4900: 428a1418                 movb (%rax, %r11), %dl
  4904: 32da                     xorb %dl, %bl
  4906: 43881c1a                 movb %bl, (%r10, %r11)
  4910: 33db                     xorl %ebx, %ebx
  4912: 49ffc3                   incq %r11
  4915: 4c3bd9                   cmpq %rcx, %r11
  4918: 0f8289ffffff             jb 0x12c5
  4924: 4d8b9190000000           movq 0x90(%r9), %r10
  4931: 4c8b1c24                 movq (%rsp), %r11
  4935: 498b4038                 movq 0x38(%r8), %rax
  4939: 428a1418                 movb (%rax, %r11), %dl
  4943: 8ada                     movb %dl, %bl
  4945: 498b8098000000           movq 0x98(%r8), %rax
  4952: 428a1418                 movb (%rax, %r11), %dl
  4956: 32da                     xorb %dl, %bl
  4958: 498b80f8000000           movq 0xf8(%r8), %rax
  4965: 428a1418                 movb (%rax, %r11), %dl
  4969: 32da                     xorb %dl, %bl
  4971: 498b8058010000           movq 0x158(%r8), %rax
  4978: 428a1418                 movb (%rax, %r11), %dl
  4982: 32da                     xorb %dl, %bl
  4984: 498b8060010000           movq 0x160(%r8), %rax
  4991: 428a1418                 movb (%rax, %r11), %dl
  4995: 32da                     xorb %dl, %bl
  4997: 498b80c0010000           movq 0x1c0(%r8), %rax
  5004: 428a1418                 movb (%rax, %r11), %dl
  5008: 32da                     xorb %dl, %bl
  5010: 498b8020020000           movq 0x220(%r8), %rax
  5017: 428a1418                 movb (%rax, %r11), %dl
  5021: 32da                     xorb %dl, %bl
  5023: 498b8080020000           movq 0x280(%r8), %rax
  5030: 428a1418                 movb (%rax, %r11), %dl
  5034: 32da                     xorb %dl, %bl
  5036: 43881c1a                 movb %bl, (%r10, %r11)
  5040: 33db                     xorl %ebx, %ebx
  5042: 49ffc3                   incq %r11
  5045: 4c3bd9                   cmpq %rcx, %r11
  5048: 0f8289ffffff             jb 0x1347
  5054: 4d8b9198000000           movq 0x98(%r9), %r10
  5061: 4c8b1c24                 movq (%rsp), %r11
  5065: 498b4040                 movq 0x40(%r8), %rax
  5069: 428a1418                 movb (%rax, %r11), %dl
  5073: 8ada                     movb %dl, %bl
  5075: 498b80a0000000           movq 0xa0(%r8), %rax
  5082: 428a1418                 movb (%rax, %r11), %dl
  5086: 32da                     xorb %dl, %bl
  5088: 498b8000010000           movq 0x100(%r8), %rax
  5095: 428a1418                 movb (%rax, %r11), %dl
  5099: 32da                     xorb %dl, %bl
  5101: 498b8008010000           movq 0x108(%r8), %rax
  5108: 428a1418                 movb (%rax, %r11), %dl
  5112: 32da                     xorb %dl, %bl
  5114: 498b8068010000           movq 0x168(%r8), %rax
  5121: 428a1418                 movb (%rax, %r11), %dl
  5125: 32da                     xorb %dl, %bl
  5127: 498b80c8010000           movq 0x1c8(%r8), %rax
  5134: 428a1418                 movb (%rax, %r11), %dl
  5138: 32da                     xorb %dl, %bl
  5140: 498b8020020000           movq 0x220(%r8), %rax
  5147: 428a1418                 movb (%rax, %r11), %dl
  5151: 32da                     xorb %dl, %bl
  5153: 498b8028020000           movq 0x228(%r8), %rax
  5160: 428a1418                 movb (%rax, %r11), %dl
  5164: 32da                     xorb %dl, %bl
  5166: 498b8088020000           movq 0x288(%r8), %rax
  5173: 428a1418                 movb (%rax, %r11), %dl
  5177: 32da                     xorb %dl, %bl
  5179: 43881c1a                 movb %bl, (%r10, %r11)
  5183: 33db                     xorl %ebx, %ebx
  5185: 49ffc3                   incq %r11
  5188: 4c3bd9                   cmpq %rcx, %r11
  5191: 0f827cffffff             jb 0x13c9
  5197: 4d8b91a0000000           movq 0xa0(%r9), %r10
  5204: 4c8b1c24                 movq (%rsp), %r11
  5208: 498b4048                 movq 0x48(%r8), %rax
  5212: 428a1418                 movb (%rax, %r11), %dl
  5216: 8ada                     movb %dl, %bl
  5218: 498b80a8000000           movq 0xa8(%r8), %rax
  5225: 428a1418                 movb (%rax, %r11), %dl
  5229: 32da                     xorb %dl, %bl
  5231: 498b80b0000000           movq 0xb0(%r8), %rax
  5238: 428a1418                 movb (%rax, %r11), %dl
  5242: 32da                     xorb %dl, %bl
  5244: 498b8010010000           movq 0x110(%r8), %rax
  5251: 428a1418                 movb (%rax, %r11), %dl
  5255: 32da                     xorb %dl, %bl
  5257: 498b8068010000           movq 0x168(%r8), %rax
  5264: 428a1418                 movb (%rax, %r11), %dl
  5268: 32da                     xorb %dl, %bl
  5270: 498b8070010000           movq 0x170(%r8), %rax
  5277: 428a1418                 movb (%rax, %r11), %dl
  5281: 32da                     xorb %dl, %bl
  5283: 498b80d0010000           movq 0x1d0(%r8), %rax
  5290: 428a1418                 movb (%rax, %r11), %dl
  5294: 32da                     xorb %dl, %bl
  5296: 498b8030020000           movq 0x230(%r8), %rax
  5303: 428a1418                 movb (%rax, %r11), %dl
  5307: 32da                     xorb %dl, %bl
  5309: 498b8090020000           movq 0x290(%r8), %rax
  5316: 428a1418                 movb (%rax, %r11), %dl
  5320: 32da                     xorb %dl, %bl
  5322: 43881c1a                 movb %bl, (%r10, %r11)
  5326: 33db                     xorl %ebx, %ebx
  5328: 49ffc3                   incq %r11
  5331: 4c3bd9                   cmpq %rcx, %r11
  5334: 0f827cffffff             jb 0x1458
  5340: 4d8b91a8000000           movq 0xa8(%r9), %r10
  5347: 4c8b1c24                 movq (%rsp), %r11
  5351: 498b4050                 movq 0x50(%r8), %rax
  5355: 428a1418                 movb (%rax, %r11), %dl
  5359: 8ada                     movb %dl, %bl
  5361: 498b4058                 movq 0x58(%r8), %rax
  5365: 428a1418                 movb (%rax, %r11), %dl
  5369: 32da                     xorb %dl, %bl
  5371: 498b80b0000000           movq 0xb0(%r8), %rax
  5378: 428a1418                 movb (%rax, %r11), %dl
  5382: 32da                     xorb %dl, %bl
  5384: 498b80b8000000           movq 0xb8(%r8), %rax
  5391: 428a1418                 movb (%rax, %r11), %dl
  5395: 32da                     xorb %dl, %bl
  5397: 498b8018010000           movq 0x118(%r8), %rax
  5404: 428a1418                 movb (%rax, %r11), %dl
  5408: 32da                     xorb %dl, %bl
  5410: 498b8078010000           movq 0x178(%r8), %rax
  5417: 428a1418                 movb (%rax, %r11), %dl
  5421: 32da                     xorb %dl, %bl
  5423: 498b80d8010000           movq 0x1d8(%r8), %rax
  5430: 428a1418                 movb (%rax, %r11), %dl
  5434: 32da                     xorb %dl, %bl
  5436: 498b8038020000           movq 0x238(%r8), %rax
  5443: 428a1418                 movb (%rax, %r11), %dl
  5447: 32da                     xorb %dl, %bl
  5449: 498b8098020000           movq 0x298(%r8), %rax
  5456: 428a1418                 movb (%rax, %r11), %dl
  5460: 32da                     xorb %dl, %bl
  5462: 43881c1a                 movb %bl, (%r10, %r11)
  5466: 33db                     xorl %ebx, %ebx
  5468: 49ffc3                   incq %r11
  5471: 4c3bd9                   cmpq %rcx, %r11
  5474: 0f827fffffff             jb 0x14e7
  5480: 4883c408                 addq $8, %rsp
  5484: 415e                     popq %r14
  5486: 415d                     popq %r13
  5488: 415c                     popq %r12
  5490: 5b                       popq %rbx
  5491: c5f877                   vzeroupper 
  5494: c3                       retq 
  5495: 0f0b                     ud2 
  5497: 0f0b                     ud2 
  5499: 0f0b                     ud2 
  5501: 0f0b                     ud2 
  5503: cc                       int3 
