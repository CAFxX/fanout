; config_t1_136 target=avx512-gfni twin=maskedtail topo=liberation(8)/enc op=encode U=1
; code_size=5376 twins=["maskedtail"] uops_per_byte=5.875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de7140000           leaq 0x14e7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4883e840                 subq $0x40, %rax
    54: 0f82eb090000             jb 0xa27
    60: 0f1f4000                 nopl (%rax)
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
  2602: 0f838a0a0000             jae 0x14ba
  2608: 4c8bd1                   movq %rcx, %r10
  2611: 4d2bd3                   subq %r11, %r10
  2614: 498bd2                   movq %r10, %rdx
  2617: 4883fa40                 cmpq $0x40, %rdx
  2621: 0f8607000000             jbe 0xa4a
  2627: 48c7c240000000           movq $0x40, %rdx
  2634: c4c1f8900cd7             kmovq (%r15, %rdx, 8), %k1
  2640: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
  2646: 62f17548efc9             vpxord %zmm1, %zmm1, %zmm1
  2652: 62f16d48efd2             vpxord %zmm2, %zmm2, %zmm2
  2658: 62f16548efdb             vpxord %zmm3, %zmm3, %zmm3
  2664: 62f15d48efe4             vpxord %zmm4, %zmm4, %zmm4
  2670: 62f15548efed             vpxord %zmm5, %zmm5, %zmm5
  2676: 62f14d48eff6             vpxord %zmm6, %zmm6, %zmm6
  2682: 62f14548efff             vpxord %zmm7, %zmm7, %zmm7
  2688: 62513d48efc0             vpxord %zmm8, %zmm8, %zmm8
  2694: 62513548efc9             vpxord %zmm9, %zmm9, %zmm9
  2700: 62512d48efd2             vpxord %zmm10, %zmm10, %zmm10
  2706: 62512548efdb             vpxord %zmm11, %zmm11, %zmm11
  2712: 62511d48efe4             vpxord %zmm12, %zmm12, %zmm12
  2718: 62511548efed             vpxord %zmm13, %zmm13, %zmm13
  2724: 62510d48eff6             vpxord %zmm14, %zmm14, %zmm14
  2730: 62510548efff             vpxord %zmm15, %zmm15, %zmm15
  2736: 62a17d40efc0             vpxord %zmm16, %zmm16, %zmm16
  2742: 62a17540efc9             vpxord %zmm17, %zmm17, %zmm17
  2748: 62a16d40efd2             vpxord %zmm18, %zmm18, %zmm18
  2754: 62a16540efdb             vpxord %zmm19, %zmm19, %zmm19
  2760: 62a15d40efe4             vpxord %zmm20, %zmm20, %zmm20
  2766: 62a15540efed             vpxord %zmm21, %zmm21, %zmm21
  2772: 62a17fc96f341e           vmovdqu8 (%rsi, %r11), %zmm22 {%k1} {z}
  2779: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  2785: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  2791: 62a17fc96f341f           vmovdqu8 (%rdi, %r11), %zmm22 {%k1} {z}
  2798: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  2804: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  2810: 62817fc96f341c           vmovdqu8 (%r12, %r11), %zmm22 {%k1} {z}
  2817: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  2823: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  2829: 62817fc96f741d00         vmovdqu8 (%r13, %r11), %zmm22 {%k1} {z}
  2837: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  2843: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  2849: 62817fc96f341e           vmovdqu8 (%r14, %r11), %zmm22 {%k1} {z}
  2856: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  2862: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  2868: 498b4028                 movq 0x28(%r8), %rax
  2872: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  2879: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  2885: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  2891: 498b4030                 movq 0x30(%r8), %rax
  2895: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  2902: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  2908: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  2914: 498b4038                 movq 0x38(%r8), %rax
  2918: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  2925: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  2931: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  2937: 498b4040                 movq 0x40(%r8), %rax
  2941: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  2948: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  2954: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  2960: 498b4048                 movq 0x48(%r8), %rax
  2964: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  2971: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  2977: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  2983: 498b4050                 movq 0x50(%r8), %rax
  2987: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  2994: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  3000: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  3006: 498b4058                 movq 0x58(%r8), %rax
  3010: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3017: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  3023: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  3029: 498b4060                 movq 0x60(%r8), %rax
  3033: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3040: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  3046: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  3052: 498b4068                 movq 0x68(%r8), %rax
  3056: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3063: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  3069: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  3075: 498b4070                 movq 0x70(%r8), %rax
  3079: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3086: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  3092: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  3098: 498b4078                 movq 0x78(%r8), %rax
  3102: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3109: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  3115: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  3121: 498b8080000000           movq 0x80(%r8), %rax
  3128: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3135: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  3141: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  3147: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  3153: 498b8088000000           movq 0x88(%r8), %rax
  3160: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3167: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  3173: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  3179: 498b8090000000           movq 0x90(%r8), %rax
  3186: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3193: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  3199: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  3205: 498b8098000000           movq 0x98(%r8), %rax
  3212: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3219: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  3225: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  3231: 498b80a0000000           movq 0xa0(%r8), %rax
  3238: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3245: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  3251: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  3257: 498b80a8000000           movq 0xa8(%r8), %rax
  3264: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3271: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  3277: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  3283: 498b80b0000000           movq 0xb0(%r8), %rax
  3290: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3297: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  3303: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  3309: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  3315: 498b80b8000000           movq 0xb8(%r8), %rax
  3322: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3329: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  3335: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  3341: 498b80c0000000           movq 0xc0(%r8), %rax
  3348: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3355: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  3361: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  3367: 498b80c8000000           movq 0xc8(%r8), %rax
  3374: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3381: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  3387: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  3393: 498b80d0000000           movq 0xd0(%r8), %rax
  3400: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3407: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  3413: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  3419: 498b80d8000000           movq 0xd8(%r8), %rax
  3426: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3433: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  3439: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  3445: 498b80e0000000           movq 0xe0(%r8), %rax
  3452: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3459: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  3465: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  3471: 498b80e8000000           movq 0xe8(%r8), %rax
  3478: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3485: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  3491: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  3497: 498b80f0000000           movq 0xf0(%r8), %rax
  3504: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3511: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  3517: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  3523: 498b80f8000000           movq 0xf8(%r8), %rax
  3530: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3537: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  3543: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  3549: 498b8000010000           movq 0x100(%r8), %rax
  3556: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3563: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  3569: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  3575: 498b8008010000           movq 0x108(%r8), %rax
  3582: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3589: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  3595: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  3601: 498b8010010000           movq 0x110(%r8), %rax
  3608: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3615: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  3621: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  3627: 498b8018010000           movq 0x118(%r8), %rax
  3634: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3641: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  3647: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  3653: 498b8020010000           movq 0x120(%r8), %rax
  3660: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3667: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  3673: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  3679: 498b8028010000           movq 0x128(%r8), %rax
  3686: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3693: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  3699: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  3705: 498b8030010000           movq 0x130(%r8), %rax
  3712: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3719: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  3725: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  3731: 498b8038010000           movq 0x138(%r8), %rax
  3738: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3745: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  3751: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  3757: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  3763: 498b8040010000           movq 0x140(%r8), %rax
  3770: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3777: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  3783: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  3789: 498b8048010000           movq 0x148(%r8), %rax
  3796: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3803: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  3809: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  3815: 498b8050010000           movq 0x150(%r8), %rax
  3822: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3829: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  3835: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  3841: 498b8058010000           movq 0x158(%r8), %rax
  3848: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3855: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  3861: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  3867: 498b8060010000           movq 0x160(%r8), %rax
  3874: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3881: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  3887: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  3893: 498b8068010000           movq 0x168(%r8), %rax
  3900: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3907: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  3913: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  3919: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  3925: 498b8070010000           movq 0x170(%r8), %rax
  3932: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3939: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  3945: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  3951: 498b8078010000           movq 0x178(%r8), %rax
  3958: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3965: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  3971: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  3977: 498b8080010000           movq 0x180(%r8), %rax
  3984: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  3991: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  3997: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  4003: 498b8088010000           movq 0x188(%r8), %rax
  4010: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4017: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  4023: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  4029: 498b8090010000           movq 0x190(%r8), %rax
  4036: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4043: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  4049: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  4055: 498b8098010000           movq 0x198(%r8), %rax
  4062: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4069: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  4075: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  4081: 498b80a0010000           movq 0x1a0(%r8), %rax
  4088: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4095: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  4101: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  4107: 498b80a8010000           movq 0x1a8(%r8), %rax
  4114: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4121: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  4127: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  4133: 498b80b0010000           movq 0x1b0(%r8), %rax
  4140: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4147: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  4153: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  4159: 498b80b8010000           movq 0x1b8(%r8), %rax
  4166: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4173: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  4179: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  4185: 498b80c0010000           movq 0x1c0(%r8), %rax
  4192: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4199: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  4205: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  4211: 498b80c8010000           movq 0x1c8(%r8), %rax
  4218: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4225: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  4231: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  4237: 498b80d0010000           movq 0x1d0(%r8), %rax
  4244: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4251: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  4257: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  4263: 498b80d8010000           movq 0x1d8(%r8), %rax
  4270: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4277: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  4283: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  4289: 498b80e0010000           movq 0x1e0(%r8), %rax
  4296: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4303: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  4309: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  4315: 498b80e8010000           movq 0x1e8(%r8), %rax
  4322: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4329: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  4335: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  4341: 498b80f0010000           movq 0x1f0(%r8), %rax
  4348: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4355: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  4361: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  4367: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  4373: 498b80f8010000           movq 0x1f8(%r8), %rax
  4380: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4387: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  4393: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  4399: 498b8000020000           movq 0x200(%r8), %rax
  4406: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4413: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  4419: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  4425: 498b8008020000           movq 0x208(%r8), %rax
  4432: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4439: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  4445: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  4451: 498b8010020000           movq 0x210(%r8), %rax
  4458: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4465: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  4471: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  4477: 498b8018020000           movq 0x218(%r8), %rax
  4484: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4491: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  4497: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  4503: 498b8020020000           movq 0x220(%r8), %rax
  4510: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4517: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  4523: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  4529: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  4535: 498b8028020000           movq 0x228(%r8), %rax
  4542: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4549: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  4555: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  4561: 498b8030020000           movq 0x230(%r8), %rax
  4568: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4575: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  4581: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  4587: 498b8038020000           movq 0x238(%r8), %rax
  4594: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4601: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  4607: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  4613: 498b8040020000           movq 0x240(%r8), %rax
  4620: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4627: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  4633: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  4639: 498b8048020000           movq 0x248(%r8), %rax
  4646: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4653: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  4659: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  4665: 498b8050020000           movq 0x250(%r8), %rax
  4672: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4679: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  4685: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  4691: 498b8058020000           movq 0x258(%r8), %rax
  4698: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4705: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  4711: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  4717: 498b8060020000           movq 0x260(%r8), %rax
  4724: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4731: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  4737: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  4743: 498b8068020000           movq 0x268(%r8), %rax
  4750: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4757: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
  4763: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  4769: 498b8070020000           movq 0x270(%r8), %rax
  4776: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4783: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  4789: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  4795: 498b8078020000           movq 0x278(%r8), %rax
  4802: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4809: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  4815: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  4821: 498b8080020000           movq 0x280(%r8), %rax
  4828: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4835: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  4841: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  4847: 498b8088020000           movq 0x288(%r8), %rax
  4854: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4861: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  4867: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  4873: 498b8090020000           movq 0x290(%r8), %rax
  4880: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4887: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  4893: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  4899: 498b8098020000           movq 0x298(%r8), %rax
  4906: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4913: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  4919: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  4925: 498b80a0020000           movq 0x2a0(%r8), %rax
  4932: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4939: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  4945: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  4951: 498b80a8020000           movq 0x2a8(%r8), %rax
  4958: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4965: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  4971: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  4977: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  4983: 498b80b0020000           movq 0x2b0(%r8), %rax
  4990: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  4997: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  5003: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  5009: 498b80b8020000           movq 0x2b8(%r8), %rax
  5016: 62a17fc96f3418           vmovdqu8 (%rax, %r11), %zmm22 {%k1} {z}
  5023: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  5029: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  5035: 498b01                   movq (%r9), %rax
  5038: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
  5045: 498b4108                 movq 8(%r9), %rax
  5049: 62b17f497f0c18           vmovdqu8 %zmm1, (%rax, %r11) {%k1}
  5056: 498b4110                 movq 0x10(%r9), %rax
  5060: 62b17f497f1418           vmovdqu8 %zmm2, (%rax, %r11) {%k1}
  5067: 498b4118                 movq 0x18(%r9), %rax
  5071: 62b17f497f1c18           vmovdqu8 %zmm3, (%rax, %r11) {%k1}
  5078: 498b4120                 movq 0x20(%r9), %rax
  5082: 62b17f497f2418           vmovdqu8 %zmm4, (%rax, %r11) {%k1}
  5089: 498b4128                 movq 0x28(%r9), %rax
  5093: 62b17f497f2c18           vmovdqu8 %zmm5, (%rax, %r11) {%k1}
  5100: 498b4130                 movq 0x30(%r9), %rax
  5104: 62b17f497f3418           vmovdqu8 %zmm6, (%rax, %r11) {%k1}
  5111: 498b4138                 movq 0x38(%r9), %rax
  5115: 62b17f497f3c18           vmovdqu8 %zmm7, (%rax, %r11) {%k1}
  5122: 498b4140                 movq 0x40(%r9), %rax
  5126: 62317f497f0418           vmovdqu8 %zmm8, (%rax, %r11) {%k1}
  5133: 498b4148                 movq 0x48(%r9), %rax
  5137: 62317f497f0c18           vmovdqu8 %zmm9, (%rax, %r11) {%k1}
  5144: 498b4150                 movq 0x50(%r9), %rax
  5148: 62317f497f1418           vmovdqu8 %zmm10, (%rax, %r11) {%k1}
  5155: 498b4158                 movq 0x58(%r9), %rax
  5159: 62317f497f1c18           vmovdqu8 %zmm11, (%rax, %r11) {%k1}
  5166: 498b4160                 movq 0x60(%r9), %rax
  5170: 62317f497f2418           vmovdqu8 %zmm12, (%rax, %r11) {%k1}
  5177: 498b4168                 movq 0x68(%r9), %rax
  5181: 62317f497f2c18           vmovdqu8 %zmm13, (%rax, %r11) {%k1}
  5188: 498b4170                 movq 0x70(%r9), %rax
  5192: 62317f497f3418           vmovdqu8 %zmm14, (%rax, %r11) {%k1}
  5199: 498b4178                 movq 0x78(%r9), %rax
  5203: 62317f497f3c18           vmovdqu8 %zmm15, (%rax, %r11) {%k1}
  5210: 498b8180000000           movq 0x80(%r9), %rax
  5217: 62a17f497f0418           vmovdqu8 %zmm16, (%rax, %r11) {%k1}
  5224: 498b8188000000           movq 0x88(%r9), %rax
  5231: 62a17f497f0c18           vmovdqu8 %zmm17, (%rax, %r11) {%k1}
  5238: 498b8190000000           movq 0x90(%r9), %rax
  5245: 62a17f497f1418           vmovdqu8 %zmm18, (%rax, %r11) {%k1}
  5252: 498b8198000000           movq 0x98(%r9), %rax
  5259: 62a17f497f1c18           vmovdqu8 %zmm19, (%rax, %r11) {%k1}
  5266: 498b81a0000000           movq 0xa0(%r9), %rax
  5273: 62a17f497f2418           vmovdqu8 %zmm20, (%rax, %r11) {%k1}
  5280: 498b81a8000000           movq 0xa8(%r9), %rax
  5287: 62a17f497f2c18           vmovdqu8 %zmm21, (%rax, %r11) {%k1}
  5294: 4c03da                   addq %rdx, %r11
  5297: 4c2bd2                   subq %rdx, %r10
  5300: 0f857cf5ffff             jne 0xa36
  5306: 415e                     popq %r14
  5308: 415d                     popq %r13
  5310: 415c                     popq %r12
  5312: 415f                     popq %r15
  5314: 5b                       popq %rbx
  5315: c5f877                   vzeroupper 
  5318: c3                       retq 
  5319: 0f0b                     ud2 
  5321: 0f0b                     ud2 
  5323: 0f0b                     ud2 
  5325: 0f0b                     ud2 
  5327: 0f0b                     ud2 
  5329: 0f0b                     ud2 
  5331: 0f0b                     ud2 
  5333: 0f0b                     ud2 
  5335: 0f0b                     ud2 
  5337: 0f0b                     ud2 
  5339: 0f0b                     ud2 
  5341: 0f0b                     ud2 
  5343: 0f0b                     ud2 
  5345: 0f0b                     ud2 
  5347: 0f0b                     ud2 
  5349: 0f0b                     ud2 
  5351: 0f0b                     ud2 
  5353: 0f0b                     ud2 
  5355: 0f0b                     ud2 
  5357: 0f0b                     ud2 
  5359: 0f0b                     ud2 
  5361: 0f0b                     ud2 
  5363: 0f0b                     ud2 
  5365: 0f0b                     ud2 
  5367: 0f0b                     ud2 
  5369: 0f0b                     ud2 
  5371: 0f0b                     ud2 
  5373: 0f0b                     ud2 
  5375: cc                       int3 
