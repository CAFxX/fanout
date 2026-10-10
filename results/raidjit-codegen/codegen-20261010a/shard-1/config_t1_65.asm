; config_t1_65 target=avx512-gfni twin=baseline topo=rdp(7)/enc op=encode U=8
; code_size=3712 twins=[] uops_per_byte=2.5078125
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
    41: 4881e880000000           subq $0x80, %rax
    48: 0f825e080000             jb 0x894
    54: 66660f1f840000000000     nopw (%rax, %rax)
    64: 62217d486f041e           vmovdqa32 (%rsi, %r11), %zmm24
    71: 62917d486fc0             vmovdqa32 %zmm24, %zmm0
    77: 62117d486fe0             vmovdqa32 %zmm24, %zmm12
    83: 62217d486f441e01         vmovdqa32 0x40(%rsi, %r11), %zmm24
    91: 62917d486fc8             vmovdqa32 %zmm24, %zmm1
    97: 62117d486fe8             vmovdqa32 %zmm24, %zmm13
   103: 62217d486f041f           vmovdqa32 (%rdi, %r11), %zmm24
   110: 62917d486fd0             vmovdqa32 %zmm24, %zmm2
   116: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
   122: 62117d486ff0             vmovdqa32 %zmm24, %zmm14
   128: 62217d486f441f01         vmovdqa32 0x40(%rdi, %r11), %zmm24
   136: 62917d486fd8             vmovdqa32 %zmm24, %zmm3
   142: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
   148: 62117d486ff8             vmovdqa32 %zmm24, %zmm15
   154: 62017d486f041c           vmovdqa32 (%r12, %r11), %zmm24
   161: 62917d486fe0             vmovdqa32 %zmm24, %zmm4
   167: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
   173: 62817d486fc0             vmovdqa32 %zmm24, %zmm16
   179: 62017d486f441c01         vmovdqa32 0x40(%r12, %r11), %zmm24
   187: 62917d486fe8             vmovdqa32 %zmm24, %zmm5
   193: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
   199: 62817d486fc8             vmovdqa32 %zmm24, %zmm17
   205: 62017d486f441d00         vmovdqa32 (%r13, %r11), %zmm24
   213: 62917d486ff0             vmovdqa32 %zmm24, %zmm6
   219: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
   225: 62817d486fd0             vmovdqa32 %zmm24, %zmm18
   231: 62017d486f441d01         vmovdqa32 0x40(%r13, %r11), %zmm24
   239: 62917d486ff8             vmovdqa32 %zmm24, %zmm7
   245: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
   251: 62817d486fd8             vmovdqa32 %zmm24, %zmm19
   257: 62017d486f041e           vmovdqa32 (%r14, %r11), %zmm24
   264: 62117d486fc0             vmovdqa32 %zmm24, %zmm8
   270: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
   276: 62817d486fe0             vmovdqa32 %zmm24, %zmm20
   282: 62017d486f441e01         vmovdqa32 0x40(%r14, %r11), %zmm24
   290: 62117d486fc8             vmovdqa32 %zmm24, %zmm9
   296: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
   302: 62817d486fe8             vmovdqa32 %zmm24, %zmm21
   308: 4d8b5028                 movq 0x28(%r8), %r10
   312: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   319: 62117d486fd0             vmovdqa32 %zmm24, %zmm10
   325: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
   331: 62817d486ff0             vmovdqa32 %zmm24, %zmm22
   337: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   345: 62117d486fd8             vmovdqa32 %zmm24, %zmm11
   351: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
   357: 62817d486ff8             vmovdqa32 %zmm24, %zmm23
   363: 4d8b5030                 movq 0x30(%r8), %r10
   367: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   374: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
   380: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
   386: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   394: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
   400: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
   406: 4d8b5038                 movq 0x38(%r8), %r10
   410: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   417: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
   423: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
   429: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
   435: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   443: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
   449: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
   455: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
   461: 4d8b5040                 movq 0x40(%r8), %r10
   465: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   472: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
   478: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
   484: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
   490: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   498: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
   504: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
   510: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
   516: 4d8b5048                 movq 0x48(%r8), %r10
   520: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   527: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
   533: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
   539: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
   545: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   553: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
   559: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
   565: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
   571: 4d8b5050                 movq 0x50(%r8), %r10
   575: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   582: 62113d48efc0             vpxord %zmm24, %zmm8, %zmm8
   588: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
   594: 62814d40eff0             vpxord %zmm24, %zmm22, %zmm22
   600: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   608: 62113548efc8             vpxord %zmm24, %zmm9, %zmm9
   614: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
   620: 62814540eff8             vpxord %zmm24, %zmm23, %zmm23
   626: 4d8b5058                 movq 0x58(%r8), %r10
   630: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   637: 62112d48efd0             vpxord %zmm24, %zmm10, %zmm10
   643: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
   649: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   657: 62112548efd8             vpxord %zmm24, %zmm11, %zmm11
   663: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
   669: 4d8b5060                 movq 0x60(%r8), %r10
   673: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   680: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
   686: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
   692: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   700: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
   706: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
   712: 4d8b5068                 movq 0x68(%r8), %r10
   716: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   723: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
   729: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
   735: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
   741: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   749: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
   755: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
   761: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
   767: 4d8b5070                 movq 0x70(%r8), %r10
   771: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   778: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
   784: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
   790: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
   796: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   804: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
   810: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
   816: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
   822: 4d8b5078                 movq 0x78(%r8), %r10
   826: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   833: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
   839: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
   845: 62814d40eff0             vpxord %zmm24, %zmm22, %zmm22
   851: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   859: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
   865: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
   871: 62814540eff8             vpxord %zmm24, %zmm23, %zmm23
   877: 4d8b9080000000           movq 0x80(%r8), %r10
   884: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   891: 62113d48efc0             vpxord %zmm24, %zmm8, %zmm8
   897: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
   903: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   911: 62113548efc8             vpxord %zmm24, %zmm9, %zmm9
   917: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
   923: 4d8b9088000000           movq 0x88(%r8), %r10
   930: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   937: 62112d48efd0             vpxord %zmm24, %zmm10, %zmm10
   943: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
   949: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
   955: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   963: 62112548efd8             vpxord %zmm24, %zmm11, %zmm11
   969: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
   975: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
   981: 4d8b9090000000           movq 0x90(%r8), %r10
   988: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   995: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
  1001: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
  1007: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1015: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
  1021: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
  1027: 4d8b9098000000           movq 0x98(%r8), %r10
  1034: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1041: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
  1047: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
  1053: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
  1059: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1067: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
  1073: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
  1079: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
  1085: 4d8b90a0000000           movq 0xa0(%r8), %r10
  1092: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1099: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
  1105: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
  1111: 62814d40eff0             vpxord %zmm24, %zmm22, %zmm22
  1117: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1125: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
  1131: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
  1137: 62814540eff8             vpxord %zmm24, %zmm23, %zmm23
  1143: 4d8b90a8000000           movq 0xa8(%r8), %r10
  1150: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1157: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
  1163: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
  1169: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1177: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
  1183: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
  1189: 4d8b90b0000000           movq 0xb0(%r8), %r10
  1196: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1203: 62113d48efc0             vpxord %zmm24, %zmm8, %zmm8
  1209: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
  1215: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
  1221: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1229: 62113548efc8             vpxord %zmm24, %zmm9, %zmm9
  1235: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
  1241: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
  1247: 4d8b90b8000000           movq 0xb8(%r8), %r10
  1254: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1261: 62112d48efd0             vpxord %zmm24, %zmm10, %zmm10
  1267: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
  1273: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
  1279: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1287: 62112548efd8             vpxord %zmm24, %zmm11, %zmm11
  1293: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
  1299: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
  1305: 4d8b90c0000000           movq 0xc0(%r8), %r10
  1312: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1319: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
  1325: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
  1331: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1339: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
  1345: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
  1351: 4d8b90c8000000           movq 0xc8(%r8), %r10
  1358: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1365: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
  1371: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
  1377: 62814d40eff0             vpxord %zmm24, %zmm22, %zmm22
  1383: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1391: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
  1397: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
  1403: 62814540eff8             vpxord %zmm24, %zmm23, %zmm23
  1409: 4d8b90d0000000           movq 0xd0(%r8), %r10
  1416: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1423: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
  1429: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
  1435: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1443: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
  1449: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
  1455: 4d8b90d8000000           movq 0xd8(%r8), %r10
  1462: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1469: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
  1475: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
  1481: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
  1487: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1495: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
  1501: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
  1507: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
  1513: 4d8b90e0000000           movq 0xe0(%r8), %r10
  1520: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1527: 62113d48efc0             vpxord %zmm24, %zmm8, %zmm8
  1533: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
  1539: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
  1545: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1553: 62113548efc8             vpxord %zmm24, %zmm9, %zmm9
  1559: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
  1565: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
  1571: 4d8b90e8000000           movq 0xe8(%r8), %r10
  1578: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1585: 62112d48efd0             vpxord %zmm24, %zmm10, %zmm10
  1591: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
  1597: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
  1603: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1611: 62112548efd8             vpxord %zmm24, %zmm11, %zmm11
  1617: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
  1623: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
  1629: 4d8b90f0000000           movq 0xf0(%r8), %r10
  1636: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1643: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
  1649: 62814d40eff0             vpxord %zmm24, %zmm22, %zmm22
  1655: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1663: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
  1669: 62814540eff8             vpxord %zmm24, %zmm23, %zmm23
  1675: 4d8b90f8000000           movq 0xf8(%r8), %r10
  1682: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1689: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
  1695: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
  1701: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1709: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
  1715: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
  1721: 4d8b9000010000           movq 0x100(%r8), %r10
  1728: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1735: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
  1741: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
  1747: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
  1753: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1761: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
  1767: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
  1773: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
  1779: 4d8b9008010000           movq 0x108(%r8), %r10
  1786: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1793: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
  1799: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
  1805: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
  1811: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1819: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
  1825: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
  1831: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
  1837: 4d8b9010010000           movq 0x110(%r8), %r10
  1844: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1851: 62113d48efc0             vpxord %zmm24, %zmm8, %zmm8
  1857: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
  1863: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
  1869: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1877: 62113548efc8             vpxord %zmm24, %zmm9, %zmm9
  1883: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
  1889: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
  1895: 4d8b9018010000           movq 0x118(%r8), %r10
  1902: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1909: 62112d48efd0             vpxord %zmm24, %zmm10, %zmm10
  1915: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
  1921: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
  1927: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1935: 62112548efd8             vpxord %zmm24, %zmm11, %zmm11
  1941: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
  1947: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
  1953: 4d8b11                   movq (%r9), %r10
  1956: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  1963: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  1971: 4d8b5108                 movq 8(%r9), %r10
  1975: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
  1982: 62917d487f5c1a01         vmovdqa32 %zmm3, 0x40(%r10, %r11)
  1990: 4d8b5110                 movq 0x10(%r9), %r10
  1994: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
  2001: 62917d487f6c1a01         vmovdqa32 %zmm5, 0x40(%r10, %r11)
  2009: 4d8b5118                 movq 0x18(%r9), %r10
  2013: 62917d487f341a           vmovdqa32 %zmm6, (%r10, %r11)
  2020: 62917d487f7c1a01         vmovdqa32 %zmm7, 0x40(%r10, %r11)
  2028: 4d8b5120                 movq 0x20(%r9), %r10
  2032: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  2039: 62117d487f4c1a01         vmovdqa32 %zmm9, 0x40(%r10, %r11)
  2047: 4d8b5128                 movq 0x28(%r9), %r10
  2051: 62117d487f141a           vmovdqa32 %zmm10, (%r10, %r11)
  2058: 62117d487f5c1a01         vmovdqa32 %zmm11, 0x40(%r10, %r11)
  2066: 4d8b5130                 movq 0x30(%r9), %r10
  2070: 62117d487f241a           vmovdqa32 %zmm12, (%r10, %r11)
  2077: 62117d487f6c1a01         vmovdqa32 %zmm13, 0x40(%r10, %r11)
  2085: 4d8b5138                 movq 0x38(%r9), %r10
  2089: 62117d487f341a           vmovdqa32 %zmm14, (%r10, %r11)
  2096: 62117d487f7c1a01         vmovdqa32 %zmm15, 0x40(%r10, %r11)
  2104: 4d8b5140                 movq 0x40(%r9), %r10
  2108: 62817d487f041a           vmovdqa32 %zmm16, (%r10, %r11)
  2115: 62817d487f4c1a01         vmovdqa32 %zmm17, 0x40(%r10, %r11)
  2123: 4d8b5148                 movq 0x48(%r9), %r10
  2127: 62817d487f141a           vmovdqa32 %zmm18, (%r10, %r11)
  2134: 62817d487f5c1a01         vmovdqa32 %zmm19, 0x40(%r10, %r11)
  2142: 4d8b5150                 movq 0x50(%r9), %r10
  2146: 62817d487f241a           vmovdqa32 %zmm20, (%r10, %r11)
  2153: 62817d487f6c1a01         vmovdqa32 %zmm21, 0x40(%r10, %r11)
  2161: 4d8b5158                 movq 0x58(%r9), %r10
  2165: 62817d487f341a           vmovdqa32 %zmm22, (%r10, %r11)
  2172: 62817d487f7c1a01         vmovdqa32 %zmm23, 0x40(%r10, %r11)
  2180: 4981c380000000           addq $0x80, %r11
  2187: 4c3bd8                   cmpq %rax, %r11
  2190: 0f86acf7ffff             jbe 0x40
  2196: 4c3bd9                   cmpq %rcx, %r11
  2199: 0f83a4050000             jae 0xe41
  2205: 4153                     pushq %r11
  2207: 4d8b11                   movq (%r9), %r10
  2210: 4c8b1c24                 movq (%rsp), %r11
  2214: 498b00                   movq (%r8), %rax
  2217: 428a1418                 movb (%rax, %r11), %dl
  2221: 8ada                     movb %dl, %bl
  2223: 498b4030                 movq 0x30(%r8), %rax
  2227: 428a1418                 movb (%rax, %r11), %dl
  2231: 32da                     xorb %dl, %bl
  2233: 498b4060                 movq 0x60(%r8), %rax
  2237: 428a1418                 movb (%rax, %r11), %dl
  2241: 32da                     xorb %dl, %bl
  2243: 498b8090000000           movq 0x90(%r8), %rax
  2250: 428a1418                 movb (%rax, %r11), %dl
  2254: 32da                     xorb %dl, %bl
  2256: 498b80c0000000           movq 0xc0(%r8), %rax
  2263: 428a1418                 movb (%rax, %r11), %dl
  2267: 32da                     xorb %dl, %bl
  2269: 498b80f0000000           movq 0xf0(%r8), %rax
  2276: 428a1418                 movb (%rax, %r11), %dl
  2280: 32da                     xorb %dl, %bl
  2282: 43881c1a                 movb %bl, (%r10, %r11)
  2286: 33db                     xorl %ebx, %ebx
  2288: 49ffc3                   incq %r11
  2291: 4c3bd9                   cmpq %rcx, %r11
  2294: 0f82aaffffff             jb 0x8a6
  2300: 4d8b5108                 movq 8(%r9), %r10
  2304: 4c8b1c24                 movq (%rsp), %r11
  2308: 498b4008                 movq 8(%r8), %rax
  2312: 428a1418                 movb (%rax, %r11), %dl
  2316: 8ada                     movb %dl, %bl
  2318: 498b4038                 movq 0x38(%r8), %rax
  2322: 428a1418                 movb (%rax, %r11), %dl
  2326: 32da                     xorb %dl, %bl
  2328: 498b4068                 movq 0x68(%r8), %rax
  2332: 428a1418                 movb (%rax, %r11), %dl
  2336: 32da                     xorb %dl, %bl
  2338: 498b8098000000           movq 0x98(%r8), %rax
  2345: 428a1418                 movb (%rax, %r11), %dl
  2349: 32da                     xorb %dl, %bl
  2351: 498b80c8000000           movq 0xc8(%r8), %rax
  2358: 428a1418                 movb (%rax, %r11), %dl
  2362: 32da                     xorb %dl, %bl
  2364: 498b80f8000000           movq 0xf8(%r8), %rax
  2371: 428a1418                 movb (%rax, %r11), %dl
  2375: 32da                     xorb %dl, %bl
  2377: 43881c1a                 movb %bl, (%r10, %r11)
  2381: 33db                     xorl %ebx, %ebx
  2383: 49ffc3                   incq %r11
  2386: 4c3bd9                   cmpq %rcx, %r11
  2389: 0f82a9ffffff             jb 0x904
  2395: 4d8b5110                 movq 0x10(%r9), %r10
  2399: 4c8b1c24                 movq (%rsp), %r11
  2403: 498b4010                 movq 0x10(%r8), %rax
  2407: 428a1418                 movb (%rax, %r11), %dl
  2411: 8ada                     movb %dl, %bl
  2413: 498b4040                 movq 0x40(%r8), %rax
  2417: 428a1418                 movb (%rax, %r11), %dl
  2421: 32da                     xorb %dl, %bl
  2423: 498b4070                 movq 0x70(%r8), %rax
  2427: 428a1418                 movb (%rax, %r11), %dl
  2431: 32da                     xorb %dl, %bl
  2433: 498b80a0000000           movq 0xa0(%r8), %rax
  2440: 428a1418                 movb (%rax, %r11), %dl
  2444: 32da                     xorb %dl, %bl
  2446: 498b80d0000000           movq 0xd0(%r8), %rax
  2453: 428a1418                 movb (%rax, %r11), %dl
  2457: 32da                     xorb %dl, %bl
  2459: 498b8000010000           movq 0x100(%r8), %rax
  2466: 428a1418                 movb (%rax, %r11), %dl
  2470: 32da                     xorb %dl, %bl
  2472: 43881c1a                 movb %bl, (%r10, %r11)
  2476: 33db                     xorl %ebx, %ebx
  2478: 49ffc3                   incq %r11
  2481: 4c3bd9                   cmpq %rcx, %r11
  2484: 0f82a9ffffff             jb 0x963
  2490: 4d8b5118                 movq 0x18(%r9), %r10
  2494: 4c8b1c24                 movq (%rsp), %r11
  2498: 498b4018                 movq 0x18(%r8), %rax
  2502: 428a1418                 movb (%rax, %r11), %dl
  2506: 8ada                     movb %dl, %bl
  2508: 498b4048                 movq 0x48(%r8), %rax
  2512: 428a1418                 movb (%rax, %r11), %dl
  2516: 32da                     xorb %dl, %bl
  2518: 498b4078                 movq 0x78(%r8), %rax
  2522: 428a1418                 movb (%rax, %r11), %dl
  2526: 32da                     xorb %dl, %bl
  2528: 498b80a8000000           movq 0xa8(%r8), %rax
  2535: 428a1418                 movb (%rax, %r11), %dl
  2539: 32da                     xorb %dl, %bl
  2541: 498b80d8000000           movq 0xd8(%r8), %rax
  2548: 428a1418                 movb (%rax, %r11), %dl
  2552: 32da                     xorb %dl, %bl
  2554: 498b8008010000           movq 0x108(%r8), %rax
  2561: 428a1418                 movb (%rax, %r11), %dl
  2565: 32da                     xorb %dl, %bl
  2567: 43881c1a                 movb %bl, (%r10, %r11)
  2571: 33db                     xorl %ebx, %ebx
  2573: 49ffc3                   incq %r11
  2576: 4c3bd9                   cmpq %rcx, %r11
  2579: 0f82a9ffffff             jb 0x9c2
  2585: 4d8b5120                 movq 0x20(%r9), %r10
  2589: 4c8b1c24                 movq (%rsp), %r11
  2593: 498b4020                 movq 0x20(%r8), %rax
  2597: 428a1418                 movb (%rax, %r11), %dl
  2601: 8ada                     movb %dl, %bl
  2603: 498b4050                 movq 0x50(%r8), %rax
  2607: 428a1418                 movb (%rax, %r11), %dl
  2611: 32da                     xorb %dl, %bl
  2613: 498b8080000000           movq 0x80(%r8), %rax
  2620: 428a1418                 movb (%rax, %r11), %dl
  2624: 32da                     xorb %dl, %bl
  2626: 498b80b0000000           movq 0xb0(%r8), %rax
  2633: 428a1418                 movb (%rax, %r11), %dl
  2637: 32da                     xorb %dl, %bl
  2639: 498b80e0000000           movq 0xe0(%r8), %rax
  2646: 428a1418                 movb (%rax, %r11), %dl
  2650: 32da                     xorb %dl, %bl
  2652: 498b8010010000           movq 0x110(%r8), %rax
  2659: 428a1418                 movb (%rax, %r11), %dl
  2663: 32da                     xorb %dl, %bl
  2665: 43881c1a                 movb %bl, (%r10, %r11)
  2669: 33db                     xorl %ebx, %ebx
  2671: 49ffc3                   incq %r11
  2674: 4c3bd9                   cmpq %rcx, %r11
  2677: 0f82a6ffffff             jb 0xa21
  2683: 4d8b5128                 movq 0x28(%r9), %r10
  2687: 4c8b1c24                 movq (%rsp), %r11
  2691: 498b4028                 movq 0x28(%r8), %rax
  2695: 428a1418                 movb (%rax, %r11), %dl
  2699: 8ada                     movb %dl, %bl
  2701: 498b4058                 movq 0x58(%r8), %rax
  2705: 428a1418                 movb (%rax, %r11), %dl
  2709: 32da                     xorb %dl, %bl
  2711: 498b8088000000           movq 0x88(%r8), %rax
  2718: 428a1418                 movb (%rax, %r11), %dl
  2722: 32da                     xorb %dl, %bl
  2724: 498b80b8000000           movq 0xb8(%r8), %rax
  2731: 428a1418                 movb (%rax, %r11), %dl
  2735: 32da                     xorb %dl, %bl
  2737: 498b80e8000000           movq 0xe8(%r8), %rax
  2744: 428a1418                 movb (%rax, %r11), %dl
  2748: 32da                     xorb %dl, %bl
  2750: 498b8018010000           movq 0x118(%r8), %rax
  2757: 428a1418                 movb (%rax, %r11), %dl
  2761: 32da                     xorb %dl, %bl
  2763: 43881c1a                 movb %bl, (%r10, %r11)
  2767: 33db                     xorl %ebx, %ebx
  2769: 49ffc3                   incq %r11
  2772: 4c3bd9                   cmpq %rcx, %r11
  2775: 0f82a6ffffff             jb 0xa83
  2781: 4d8b5130                 movq 0x30(%r9), %r10
  2785: 4c8b1c24                 movq (%rsp), %r11
  2789: 498b00                   movq (%r8), %rax
  2792: 428a1418                 movb (%rax, %r11), %dl
  2796: 8ada                     movb %dl, %bl
  2798: 498b4008                 movq 8(%r8), %rax
  2802: 428a1418                 movb (%rax, %r11), %dl
  2806: 32da                     xorb %dl, %bl
  2808: 498b4038                 movq 0x38(%r8), %rax
  2812: 428a1418                 movb (%rax, %r11), %dl
  2816: 32da                     xorb %dl, %bl
  2818: 498b4068                 movq 0x68(%r8), %rax
  2822: 428a1418                 movb (%rax, %r11), %dl
  2826: 32da                     xorb %dl, %bl
  2828: 498b8088000000           movq 0x88(%r8), %rax
  2835: 428a1418                 movb (%rax, %r11), %dl
  2839: 32da                     xorb %dl, %bl
  2841: 498b8098000000           movq 0x98(%r8), %rax
  2848: 428a1418                 movb (%rax, %r11), %dl
  2852: 32da                     xorb %dl, %bl
  2854: 498b80b0000000           movq 0xb0(%r8), %rax
  2861: 428a1418                 movb (%rax, %r11), %dl
  2865: 32da                     xorb %dl, %bl
  2867: 498b80c8000000           movq 0xc8(%r8), %rax
  2874: 428a1418                 movb (%rax, %r11), %dl
  2878: 32da                     xorb %dl, %bl
  2880: 498b80d8000000           movq 0xd8(%r8), %rax
  2887: 428a1418                 movb (%rax, %r11), %dl
  2891: 32da                     xorb %dl, %bl
  2893: 498b80f8000000           movq 0xf8(%r8), %rax
  2900: 428a1418                 movb (%rax, %r11), %dl
  2904: 32da                     xorb %dl, %bl
  2906: 498b8000010000           movq 0x100(%r8), %rax
  2913: 428a1418                 movb (%rax, %r11), %dl
  2917: 32da                     xorb %dl, %bl
  2919: 43881c1a                 movb %bl, (%r10, %r11)
  2923: 33db                     xorl %ebx, %ebx
  2925: 49ffc3                   incq %r11
  2928: 4c3bd9                   cmpq %rcx, %r11
  2931: 0f826cffffff             jb 0xae5
  2937: 4d8b5138                 movq 0x38(%r9), %r10
  2941: 4c8b1c24                 movq (%rsp), %r11
  2945: 498b4008                 movq 8(%r8), %rax
  2949: 428a1418                 movb (%rax, %r11), %dl
  2953: 8ada                     movb %dl, %bl
  2955: 498b4010                 movq 0x10(%r8), %rax
  2959: 428a1418                 movb (%rax, %r11), %dl
  2963: 32da                     xorb %dl, %bl
  2965: 498b4030                 movq 0x30(%r8), %rax
  2969: 428a1418                 movb (%rax, %r11), %dl
  2973: 32da                     xorb %dl, %bl
  2975: 498b4040                 movq 0x40(%r8), %rax
  2979: 428a1418                 movb (%rax, %r11), %dl
  2983: 32da                     xorb %dl, %bl
  2985: 498b4070                 movq 0x70(%r8), %rax
  2989: 428a1418                 movb (%rax, %r11), %dl
  2993: 32da                     xorb %dl, %bl
  2995: 498b80a0000000           movq 0xa0(%r8), %rax
  3002: 428a1418                 movb (%rax, %r11), %dl
  3006: 32da                     xorb %dl, %bl
  3008: 498b80b8000000           movq 0xb8(%r8), %rax
  3015: 428a1418                 movb (%rax, %r11), %dl
  3019: 32da                     xorb %dl, %bl
  3021: 498b80d0000000           movq 0xd0(%r8), %rax
  3028: 428a1418                 movb (%rax, %r11), %dl
  3032: 32da                     xorb %dl, %bl
  3034: 498b80e0000000           movq 0xe0(%r8), %rax
  3041: 428a1418                 movb (%rax, %r11), %dl
  3045: 32da                     xorb %dl, %bl
  3047: 498b8000010000           movq 0x100(%r8), %rax
  3054: 428a1418                 movb (%rax, %r11), %dl
  3058: 32da                     xorb %dl, %bl
  3060: 498b8008010000           movq 0x108(%r8), %rax
  3067: 428a1418                 movb (%rax, %r11), %dl
  3071: 32da                     xorb %dl, %bl
  3073: 43881c1a                 movb %bl, (%r10, %r11)
  3077: 33db                     xorl %ebx, %ebx
  3079: 49ffc3                   incq %r11
  3082: 4c3bd9                   cmpq %rcx, %r11
  3085: 0f826effffff             jb 0xb81
  3091: 4d8b5140                 movq 0x40(%r9), %r10
  3095: 4c8b1c24                 movq (%rsp), %r11
  3099: 498b4010                 movq 0x10(%r8), %rax
  3103: 428a1418                 movb (%rax, %r11), %dl
  3107: 8ada                     movb %dl, %bl
  3109: 498b4018                 movq 0x18(%r8), %rax
  3113: 428a1418                 movb (%rax, %r11), %dl
  3117: 32da                     xorb %dl, %bl
  3119: 498b4038                 movq 0x38(%r8), %rax
  3123: 428a1418                 movb (%rax, %r11), %dl
  3127: 32da                     xorb %dl, %bl
  3129: 498b4048                 movq 0x48(%r8), %rax
  3133: 428a1418                 movb (%rax, %r11), %dl
  3137: 32da                     xorb %dl, %bl
  3139: 498b4060                 movq 0x60(%r8), %rax
  3143: 428a1418                 movb (%rax, %r11), %dl
  3147: 32da                     xorb %dl, %bl
  3149: 498b4078                 movq 0x78(%r8), %rax
  3153: 428a1418                 movb (%rax, %r11), %dl
  3157: 32da                     xorb %dl, %bl
  3159: 498b80a8000000           movq 0xa8(%r8), %rax
  3166: 428a1418                 movb (%rax, %r11), %dl
  3170: 32da                     xorb %dl, %bl
  3172: 498b80d8000000           movq 0xd8(%r8), %rax
  3179: 428a1418                 movb (%rax, %r11), %dl
  3183: 32da                     xorb %dl, %bl
  3185: 498b80e8000000           movq 0xe8(%r8), %rax
  3192: 428a1418                 movb (%rax, %r11), %dl
  3196: 32da                     xorb %dl, %bl
  3198: 498b8008010000           movq 0x108(%r8), %rax
  3205: 428a1418                 movb (%rax, %r11), %dl
  3209: 32da                     xorb %dl, %bl
  3211: 498b8010010000           movq 0x110(%r8), %rax
  3218: 428a1418                 movb (%rax, %r11), %dl
  3222: 32da                     xorb %dl, %bl
  3224: 43881c1a                 movb %bl, (%r10, %r11)
  3228: 33db                     xorl %ebx, %ebx
  3230: 49ffc3                   incq %r11
  3233: 4c3bd9                   cmpq %rcx, %r11
  3236: 0f8271ffffff             jb 0xc1b
  3242: 4d8b5148                 movq 0x48(%r9), %r10
  3246: 4c8b1c24                 movq (%rsp), %r11
  3250: 498b4018                 movq 0x18(%r8), %rax
  3254: 428a1418                 movb (%rax, %r11), %dl
  3258: 8ada                     movb %dl, %bl
  3260: 498b4020                 movq 0x20(%r8), %rax
  3264: 428a1418                 movb (%rax, %r11), %dl
  3268: 32da                     xorb %dl, %bl
  3270: 498b4040                 movq 0x40(%r8), %rax
  3274: 428a1418                 movb (%rax, %r11), %dl
  3278: 32da                     xorb %dl, %bl
  3280: 498b4050                 movq 0x50(%r8), %rax
  3284: 428a1418                 movb (%rax, %r11), %dl
  3288: 32da                     xorb %dl, %bl
  3290: 498b4068                 movq 0x68(%r8), %rax
  3294: 428a1418                 movb (%rax, %r11), %dl
  3298: 32da                     xorb %dl, %bl
  3300: 498b8080000000           movq 0x80(%r8), %rax
  3307: 428a1418                 movb (%rax, %r11), %dl
  3311: 32da                     xorb %dl, %bl
  3313: 498b8090000000           movq 0x90(%r8), %rax
  3320: 428a1418                 movb (%rax, %r11), %dl
  3324: 32da                     xorb %dl, %bl
  3326: 498b80b0000000           movq 0xb0(%r8), %rax
  3333: 428a1418                 movb (%rax, %r11), %dl
  3337: 32da                     xorb %dl, %bl
  3339: 498b80e0000000           movq 0xe0(%r8), %rax
  3346: 428a1418                 movb (%rax, %r11), %dl
  3350: 32da                     xorb %dl, %bl
  3352: 498b8010010000           movq 0x110(%r8), %rax
  3359: 428a1418                 movb (%rax, %r11), %dl
  3363: 32da                     xorb %dl, %bl
  3365: 498b8018010000           movq 0x118(%r8), %rax
  3372: 428a1418                 movb (%rax, %r11), %dl
  3376: 32da                     xorb %dl, %bl
  3378: 43881c1a                 movb %bl, (%r10, %r11)
  3382: 33db                     xorl %ebx, %ebx
  3384: 49ffc3                   incq %r11
  3387: 4c3bd9                   cmpq %rcx, %r11
  3390: 0f826effffff             jb 0xcb2
  3396: 4d8b5150                 movq 0x50(%r9), %r10
  3400: 4c8b1c24                 movq (%rsp), %r11
  3404: 498b4020                 movq 0x20(%r8), %rax
  3408: 428a1418                 movb (%rax, %r11), %dl
  3412: 8ada                     movb %dl, %bl
  3414: 498b4028                 movq 0x28(%r8), %rax
  3418: 428a1418                 movb (%rax, %r11), %dl
  3422: 32da                     xorb %dl, %bl
  3424: 498b4048                 movq 0x48(%r8), %rax
  3428: 428a1418                 movb (%rax, %r11), %dl
  3432: 32da                     xorb %dl, %bl
  3434: 498b4058                 movq 0x58(%r8), %rax
  3438: 428a1418                 movb (%rax, %r11), %dl
  3442: 32da                     xorb %dl, %bl
  3444: 498b4070                 movq 0x70(%r8), %rax
  3448: 428a1418                 movb (%rax, %r11), %dl
  3452: 32da                     xorb %dl, %bl
  3454: 498b8088000000           movq 0x88(%r8), %rax
  3461: 428a1418                 movb (%rax, %r11), %dl
  3465: 32da                     xorb %dl, %bl
  3467: 498b8098000000           movq 0x98(%r8), %rax
  3474: 428a1418                 movb (%rax, %r11), %dl
  3478: 32da                     xorb %dl, %bl
  3480: 498b80b8000000           movq 0xb8(%r8), %rax
  3487: 428a1418                 movb (%rax, %r11), %dl
  3491: 32da                     xorb %dl, %bl
  3493: 498b80c0000000           movq 0xc0(%r8), %rax
  3500: 428a1418                 movb (%rax, %r11), %dl
  3504: 32da                     xorb %dl, %bl
  3506: 498b80e8000000           movq 0xe8(%r8), %rax
  3513: 428a1418                 movb (%rax, %r11), %dl
  3517: 32da                     xorb %dl, %bl
  3519: 498b8018010000           movq 0x118(%r8), %rax
  3526: 428a1418                 movb (%rax, %r11), %dl
  3530: 32da                     xorb %dl, %bl
  3532: 43881c1a                 movb %bl, (%r10, %r11)
  3536: 33db                     xorl %ebx, %ebx
  3538: 49ffc3                   incq %r11
  3541: 4c3bd9                   cmpq %rcx, %r11
  3544: 0f826effffff             jb 0xd4c
  3550: 4d8b5158                 movq 0x58(%r9), %r10
  3554: 4c8b1c24                 movq (%rsp), %r11
  3558: 498b4028                 movq 0x28(%r8), %rax
  3562: 428a1418                 movb (%rax, %r11), %dl
  3566: 8ada                     movb %dl, %bl
  3568: 498b4050                 movq 0x50(%r8), %rax
  3572: 428a1418                 movb (%rax, %r11), %dl
  3576: 32da                     xorb %dl, %bl
  3578: 498b4078                 movq 0x78(%r8), %rax
  3582: 428a1418                 movb (%rax, %r11), %dl
  3586: 32da                     xorb %dl, %bl
  3588: 498b80a0000000           movq 0xa0(%r8), %rax
  3595: 428a1418                 movb (%rax, %r11), %dl
  3599: 32da                     xorb %dl, %bl
  3601: 498b80c8000000           movq 0xc8(%r8), %rax
  3608: 428a1418                 movb (%rax, %r11), %dl
  3612: 32da                     xorb %dl, %bl
  3614: 498b80f0000000           movq 0xf0(%r8), %rax
  3621: 428a1418                 movb (%rax, %r11), %dl
  3625: 32da                     xorb %dl, %bl
  3627: 43881c1a                 movb %bl, (%r10, %r11)
  3631: 33db                     xorl %ebx, %ebx
  3633: 49ffc3                   incq %r11
  3636: 4c3bd9                   cmpq %rcx, %r11
  3639: 0f82a9ffffff             jb 0xde6
  3645: 4883c408                 addq $8, %rsp
  3649: 415e                     popq %r14
  3651: 415d                     popq %r13
  3653: 415c                     popq %r12
  3655: 5b                       popq %rbx
  3656: c5f877                   vzeroupper 
  3659: c3                       retq 
  3660: 0f0b                     ud2 
  3662: 0f0b                     ud2 
  3664: 0f0b                     ud2 
  3666: 0f0b                     ud2 
  3668: 0f0b                     ud2 
  3670: 0f0b                     ud2 
  3672: 0f0b                     ud2 
  3674: 0f0b                     ud2 
  3676: 0f0b                     ud2 
  3678: 0f0b                     ud2 
  3680: 0f0b                     ud2 
  3682: 0f0b                     ud2 
  3684: 0f0b                     ud2 
  3686: 0f0b                     ud2 
  3688: 0f0b                     ud2 
  3690: 0f0b                     ud2 
  3692: 0f0b                     ud2 
  3694: 0f0b                     ud2 
  3696: 0f0b                     ud2 
  3698: 0f0b                     ud2 
  3700: 0f0b                     ud2 
  3702: 0f0b                     ud2 
  3704: 0f0b                     ud2 
  3706: 0f0b                     ud2 
  3708: 0f0b                     ud2 
  3710: 0f0b                     ud2 
