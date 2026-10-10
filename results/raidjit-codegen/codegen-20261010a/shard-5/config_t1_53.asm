; config_t1_53 target=avx512-gfni twin=nt topo=rdp(7)/enc op=encode U=1
; code_size=2752 twins=["nt"] uops_per_byte=2.75
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
    45: 0f8298040000             jb 0x4cb
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 62317d486f241e           vmovdqa32 (%rsi, %r11), %zmm12
    71: 62d17d486fc4             vmovdqa32 %zmm12, %zmm0
    77: 62d17d486ff4             vmovdqa32 %zmm12, %zmm6
    83: 62317d486f241f           vmovdqa32 (%rdi, %r11), %zmm12
    90: 62d17d486fcc             vmovdqa32 %zmm12, %zmm1
    96: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   102: 62d17d486ffc             vmovdqa32 %zmm12, %zmm7
   108: 62117d486f241c           vmovdqa32 (%r12, %r11), %zmm12
   115: 62d17d486fd4             vmovdqa32 %zmm12, %zmm2
   121: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   127: 62517d486fc4             vmovdqa32 %zmm12, %zmm8
   133: 62117d486f641d00         vmovdqa32 (%r13, %r11), %zmm12
   141: 62d17d486fdc             vmovdqa32 %zmm12, %zmm3
   147: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   153: 62517d486fcc             vmovdqa32 %zmm12, %zmm9
   159: 62117d486f241e           vmovdqa32 (%r14, %r11), %zmm12
   166: 62d17d486fe4             vmovdqa32 %zmm12, %zmm4
   172: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   178: 62517d486fd4             vmovdqa32 %zmm12, %zmm10
   184: 4d8b5028                 movq 0x28(%r8), %r10
   188: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   195: 62d17d486fec             vmovdqa32 %zmm12, %zmm5
   201: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   207: 62517d486fdc             vmovdqa32 %zmm12, %zmm11
   213: 4d8b5030                 movq 0x30(%r8), %r10
   217: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   224: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   230: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   236: 4d8b5038                 movq 0x38(%r8), %r10
   240: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   247: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   253: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   259: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   265: 4d8b5040                 movq 0x40(%r8), %r10
   269: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   276: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   282: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   288: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   294: 4d8b5048                 movq 0x48(%r8), %r10
   298: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   305: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   311: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   317: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   323: 4d8b5050                 movq 0x50(%r8), %r10
   327: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   334: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
   340: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   346: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   352: 4d8b5058                 movq 0x58(%r8), %r10
   356: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   363: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
   369: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   375: 4d8b5060                 movq 0x60(%r8), %r10
   379: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   386: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   392: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   398: 4d8b5068                 movq 0x68(%r8), %r10
   402: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   409: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   415: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   421: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   427: 4d8b5070                 movq 0x70(%r8), %r10
   431: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   438: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   444: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   450: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   456: 4d8b5078                 movq 0x78(%r8), %r10
   460: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   467: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   473: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   479: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   485: 4d8b9080000000           movq 0x80(%r8), %r10
   492: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   499: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
   505: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   511: 4d8b9088000000           movq 0x88(%r8), %r10
   518: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   525: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
   531: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   537: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   543: 4d8b9090000000           movq 0x90(%r8), %r10
   550: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   557: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   563: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   569: 4d8b9098000000           movq 0x98(%r8), %r10
   576: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   583: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   589: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   595: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   601: 4d8b90a0000000           movq 0xa0(%r8), %r10
   608: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   615: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   621: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   627: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   633: 4d8b90a8000000           movq 0xa8(%r8), %r10
   640: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   647: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   653: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   659: 4d8b90b0000000           movq 0xb0(%r8), %r10
   666: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   673: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
   679: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   685: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   691: 4d8b90b8000000           movq 0xb8(%r8), %r10
   698: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   705: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
   711: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   717: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   723: 4d8b90c0000000           movq 0xc0(%r8), %r10
   730: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   737: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   743: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   749: 4d8b90c8000000           movq 0xc8(%r8), %r10
   756: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   763: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   769: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   775: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   781: 4d8b90d0000000           movq 0xd0(%r8), %r10
   788: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   795: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   801: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   807: 4d8b90d8000000           movq 0xd8(%r8), %r10
   814: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   821: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   827: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   833: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   839: 4d8b90e0000000           movq 0xe0(%r8), %r10
   846: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   853: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
   859: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   865: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   871: 4d8b90e8000000           movq 0xe8(%r8), %r10
   878: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   885: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
   891: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   897: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   903: 4d8b90f0000000           movq 0xf0(%r8), %r10
   910: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   917: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   923: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   929: 4d8b90f8000000           movq 0xf8(%r8), %r10
   936: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   943: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   949: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   955: 4d8b9000010000           movq 0x100(%r8), %r10
   962: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   969: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   975: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   981: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   987: 4d8b9008010000           movq 0x108(%r8), %r10
   994: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
  1001: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
  1007: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  1013: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1019: 4d8b9010010000           movq 0x110(%r8), %r10
  1026: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
  1033: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
  1039: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1045: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1051: 4d8b9018010000           movq 0x118(%r8), %r10
  1058: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
  1065: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
  1071: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1077: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1083: 4d8b11                   movq (%r9), %r10
  1086: 62917d48e7041a           vmovntdq %zmm0, (%r10, %r11)
  1093: 4d8b5108                 movq 8(%r9), %r10
  1097: 62917d48e70c1a           vmovntdq %zmm1, (%r10, %r11)
  1104: 4d8b5110                 movq 0x10(%r9), %r10
  1108: 62917d48e7141a           vmovntdq %zmm2, (%r10, %r11)
  1115: 4d8b5118                 movq 0x18(%r9), %r10
  1119: 62917d48e71c1a           vmovntdq %zmm3, (%r10, %r11)
  1126: 4d8b5120                 movq 0x20(%r9), %r10
  1130: 62917d48e7241a           vmovntdq %zmm4, (%r10, %r11)
  1137: 4d8b5128                 movq 0x28(%r9), %r10
  1141: 62917d48e72c1a           vmovntdq %zmm5, (%r10, %r11)
  1148: 4d8b5130                 movq 0x30(%r9), %r10
  1152: 62917d48e7341a           vmovntdq %zmm6, (%r10, %r11)
  1159: 4d8b5138                 movq 0x38(%r9), %r10
  1163: 62917d48e73c1a           vmovntdq %zmm7, (%r10, %r11)
  1170: 4d8b5140                 movq 0x40(%r9), %r10
  1174: 62117d48e7041a           vmovntdq %zmm8, (%r10, %r11)
  1181: 4d8b5148                 movq 0x48(%r9), %r10
  1185: 62117d48e70c1a           vmovntdq %zmm9, (%r10, %r11)
  1192: 4d8b5150                 movq 0x50(%r9), %r10
  1196: 62117d48e7141a           vmovntdq %zmm10, (%r10, %r11)
  1203: 4d8b5158                 movq 0x58(%r9), %r10
  1207: 62117d48e71c1a           vmovntdq %zmm11, (%r10, %r11)
  1214: 4983c340                 addq $0x40, %r11
  1218: 4c3bd8                   cmpq %rax, %r11
  1221: 0f8675fbffff             jbe 0x40
  1227: 4c3bd9                   cmpq %rcx, %r11
  1230: 0f83a4050000             jae 0xa78
  1236: 4153                     pushq %r11
  1238: 4d8b11                   movq (%r9), %r10
  1241: 4c8b1c24                 movq (%rsp), %r11
  1245: 498b00                   movq (%r8), %rax
  1248: 428a1418                 movb (%rax, %r11), %dl
  1252: 8ada                     movb %dl, %bl
  1254: 498b4030                 movq 0x30(%r8), %rax
  1258: 428a1418                 movb (%rax, %r11), %dl
  1262: 32da                     xorb %dl, %bl
  1264: 498b4060                 movq 0x60(%r8), %rax
  1268: 428a1418                 movb (%rax, %r11), %dl
  1272: 32da                     xorb %dl, %bl
  1274: 498b8090000000           movq 0x90(%r8), %rax
  1281: 428a1418                 movb (%rax, %r11), %dl
  1285: 32da                     xorb %dl, %bl
  1287: 498b80c0000000           movq 0xc0(%r8), %rax
  1294: 428a1418                 movb (%rax, %r11), %dl
  1298: 32da                     xorb %dl, %bl
  1300: 498b80f0000000           movq 0xf0(%r8), %rax
  1307: 428a1418                 movb (%rax, %r11), %dl
  1311: 32da                     xorb %dl, %bl
  1313: 43881c1a                 movb %bl, (%r10, %r11)
  1317: 33db                     xorl %ebx, %ebx
  1319: 49ffc3                   incq %r11
  1322: 4c3bd9                   cmpq %rcx, %r11
  1325: 0f82aaffffff             jb 0x4dd
  1331: 4d8b5108                 movq 8(%r9), %r10
  1335: 4c8b1c24                 movq (%rsp), %r11
  1339: 498b4008                 movq 8(%r8), %rax
  1343: 428a1418                 movb (%rax, %r11), %dl
  1347: 8ada                     movb %dl, %bl
  1349: 498b4038                 movq 0x38(%r8), %rax
  1353: 428a1418                 movb (%rax, %r11), %dl
  1357: 32da                     xorb %dl, %bl
  1359: 498b4068                 movq 0x68(%r8), %rax
  1363: 428a1418                 movb (%rax, %r11), %dl
  1367: 32da                     xorb %dl, %bl
  1369: 498b8098000000           movq 0x98(%r8), %rax
  1376: 428a1418                 movb (%rax, %r11), %dl
  1380: 32da                     xorb %dl, %bl
  1382: 498b80c8000000           movq 0xc8(%r8), %rax
  1389: 428a1418                 movb (%rax, %r11), %dl
  1393: 32da                     xorb %dl, %bl
  1395: 498b80f8000000           movq 0xf8(%r8), %rax
  1402: 428a1418                 movb (%rax, %r11), %dl
  1406: 32da                     xorb %dl, %bl
  1408: 43881c1a                 movb %bl, (%r10, %r11)
  1412: 33db                     xorl %ebx, %ebx
  1414: 49ffc3                   incq %r11
  1417: 4c3bd9                   cmpq %rcx, %r11
  1420: 0f82a9ffffff             jb 0x53b
  1426: 4d8b5110                 movq 0x10(%r9), %r10
  1430: 4c8b1c24                 movq (%rsp), %r11
  1434: 498b4010                 movq 0x10(%r8), %rax
  1438: 428a1418                 movb (%rax, %r11), %dl
  1442: 8ada                     movb %dl, %bl
  1444: 498b4040                 movq 0x40(%r8), %rax
  1448: 428a1418                 movb (%rax, %r11), %dl
  1452: 32da                     xorb %dl, %bl
  1454: 498b4070                 movq 0x70(%r8), %rax
  1458: 428a1418                 movb (%rax, %r11), %dl
  1462: 32da                     xorb %dl, %bl
  1464: 498b80a0000000           movq 0xa0(%r8), %rax
  1471: 428a1418                 movb (%rax, %r11), %dl
  1475: 32da                     xorb %dl, %bl
  1477: 498b80d0000000           movq 0xd0(%r8), %rax
  1484: 428a1418                 movb (%rax, %r11), %dl
  1488: 32da                     xorb %dl, %bl
  1490: 498b8000010000           movq 0x100(%r8), %rax
  1497: 428a1418                 movb (%rax, %r11), %dl
  1501: 32da                     xorb %dl, %bl
  1503: 43881c1a                 movb %bl, (%r10, %r11)
  1507: 33db                     xorl %ebx, %ebx
  1509: 49ffc3                   incq %r11
  1512: 4c3bd9                   cmpq %rcx, %r11
  1515: 0f82a9ffffff             jb 0x59a
  1521: 4d8b5118                 movq 0x18(%r9), %r10
  1525: 4c8b1c24                 movq (%rsp), %r11
  1529: 498b4018                 movq 0x18(%r8), %rax
  1533: 428a1418                 movb (%rax, %r11), %dl
  1537: 8ada                     movb %dl, %bl
  1539: 498b4048                 movq 0x48(%r8), %rax
  1543: 428a1418                 movb (%rax, %r11), %dl
  1547: 32da                     xorb %dl, %bl
  1549: 498b4078                 movq 0x78(%r8), %rax
  1553: 428a1418                 movb (%rax, %r11), %dl
  1557: 32da                     xorb %dl, %bl
  1559: 498b80a8000000           movq 0xa8(%r8), %rax
  1566: 428a1418                 movb (%rax, %r11), %dl
  1570: 32da                     xorb %dl, %bl
  1572: 498b80d8000000           movq 0xd8(%r8), %rax
  1579: 428a1418                 movb (%rax, %r11), %dl
  1583: 32da                     xorb %dl, %bl
  1585: 498b8008010000           movq 0x108(%r8), %rax
  1592: 428a1418                 movb (%rax, %r11), %dl
  1596: 32da                     xorb %dl, %bl
  1598: 43881c1a                 movb %bl, (%r10, %r11)
  1602: 33db                     xorl %ebx, %ebx
  1604: 49ffc3                   incq %r11
  1607: 4c3bd9                   cmpq %rcx, %r11
  1610: 0f82a9ffffff             jb 0x5f9
  1616: 4d8b5120                 movq 0x20(%r9), %r10
  1620: 4c8b1c24                 movq (%rsp), %r11
  1624: 498b4020                 movq 0x20(%r8), %rax
  1628: 428a1418                 movb (%rax, %r11), %dl
  1632: 8ada                     movb %dl, %bl
  1634: 498b4050                 movq 0x50(%r8), %rax
  1638: 428a1418                 movb (%rax, %r11), %dl
  1642: 32da                     xorb %dl, %bl
  1644: 498b8080000000           movq 0x80(%r8), %rax
  1651: 428a1418                 movb (%rax, %r11), %dl
  1655: 32da                     xorb %dl, %bl
  1657: 498b80b0000000           movq 0xb0(%r8), %rax
  1664: 428a1418                 movb (%rax, %r11), %dl
  1668: 32da                     xorb %dl, %bl
  1670: 498b80e0000000           movq 0xe0(%r8), %rax
  1677: 428a1418                 movb (%rax, %r11), %dl
  1681: 32da                     xorb %dl, %bl
  1683: 498b8010010000           movq 0x110(%r8), %rax
  1690: 428a1418                 movb (%rax, %r11), %dl
  1694: 32da                     xorb %dl, %bl
  1696: 43881c1a                 movb %bl, (%r10, %r11)
  1700: 33db                     xorl %ebx, %ebx
  1702: 49ffc3                   incq %r11
  1705: 4c3bd9                   cmpq %rcx, %r11
  1708: 0f82a6ffffff             jb 0x658
  1714: 4d8b5128                 movq 0x28(%r9), %r10
  1718: 4c8b1c24                 movq (%rsp), %r11
  1722: 498b4028                 movq 0x28(%r8), %rax
  1726: 428a1418                 movb (%rax, %r11), %dl
  1730: 8ada                     movb %dl, %bl
  1732: 498b4058                 movq 0x58(%r8), %rax
  1736: 428a1418                 movb (%rax, %r11), %dl
  1740: 32da                     xorb %dl, %bl
  1742: 498b8088000000           movq 0x88(%r8), %rax
  1749: 428a1418                 movb (%rax, %r11), %dl
  1753: 32da                     xorb %dl, %bl
  1755: 498b80b8000000           movq 0xb8(%r8), %rax
  1762: 428a1418                 movb (%rax, %r11), %dl
  1766: 32da                     xorb %dl, %bl
  1768: 498b80e8000000           movq 0xe8(%r8), %rax
  1775: 428a1418                 movb (%rax, %r11), %dl
  1779: 32da                     xorb %dl, %bl
  1781: 498b8018010000           movq 0x118(%r8), %rax
  1788: 428a1418                 movb (%rax, %r11), %dl
  1792: 32da                     xorb %dl, %bl
  1794: 43881c1a                 movb %bl, (%r10, %r11)
  1798: 33db                     xorl %ebx, %ebx
  1800: 49ffc3                   incq %r11
  1803: 4c3bd9                   cmpq %rcx, %r11
  1806: 0f82a6ffffff             jb 0x6ba
  1812: 4d8b5130                 movq 0x30(%r9), %r10
  1816: 4c8b1c24                 movq (%rsp), %r11
  1820: 498b00                   movq (%r8), %rax
  1823: 428a1418                 movb (%rax, %r11), %dl
  1827: 8ada                     movb %dl, %bl
  1829: 498b4008                 movq 8(%r8), %rax
  1833: 428a1418                 movb (%rax, %r11), %dl
  1837: 32da                     xorb %dl, %bl
  1839: 498b4038                 movq 0x38(%r8), %rax
  1843: 428a1418                 movb (%rax, %r11), %dl
  1847: 32da                     xorb %dl, %bl
  1849: 498b4068                 movq 0x68(%r8), %rax
  1853: 428a1418                 movb (%rax, %r11), %dl
  1857: 32da                     xorb %dl, %bl
  1859: 498b8088000000           movq 0x88(%r8), %rax
  1866: 428a1418                 movb (%rax, %r11), %dl
  1870: 32da                     xorb %dl, %bl
  1872: 498b8098000000           movq 0x98(%r8), %rax
  1879: 428a1418                 movb (%rax, %r11), %dl
  1883: 32da                     xorb %dl, %bl
  1885: 498b80b0000000           movq 0xb0(%r8), %rax
  1892: 428a1418                 movb (%rax, %r11), %dl
  1896: 32da                     xorb %dl, %bl
  1898: 498b80c8000000           movq 0xc8(%r8), %rax
  1905: 428a1418                 movb (%rax, %r11), %dl
  1909: 32da                     xorb %dl, %bl
  1911: 498b80d8000000           movq 0xd8(%r8), %rax
  1918: 428a1418                 movb (%rax, %r11), %dl
  1922: 32da                     xorb %dl, %bl
  1924: 498b80f8000000           movq 0xf8(%r8), %rax
  1931: 428a1418                 movb (%rax, %r11), %dl
  1935: 32da                     xorb %dl, %bl
  1937: 498b8000010000           movq 0x100(%r8), %rax
  1944: 428a1418                 movb (%rax, %r11), %dl
  1948: 32da                     xorb %dl, %bl
  1950: 43881c1a                 movb %bl, (%r10, %r11)
  1954: 33db                     xorl %ebx, %ebx
  1956: 49ffc3                   incq %r11
  1959: 4c3bd9                   cmpq %rcx, %r11
  1962: 0f826cffffff             jb 0x71c
  1968: 4d8b5138                 movq 0x38(%r9), %r10
  1972: 4c8b1c24                 movq (%rsp), %r11
  1976: 498b4008                 movq 8(%r8), %rax
  1980: 428a1418                 movb (%rax, %r11), %dl
  1984: 8ada                     movb %dl, %bl
  1986: 498b4010                 movq 0x10(%r8), %rax
  1990: 428a1418                 movb (%rax, %r11), %dl
  1994: 32da                     xorb %dl, %bl
  1996: 498b4030                 movq 0x30(%r8), %rax
  2000: 428a1418                 movb (%rax, %r11), %dl
  2004: 32da                     xorb %dl, %bl
  2006: 498b4040                 movq 0x40(%r8), %rax
  2010: 428a1418                 movb (%rax, %r11), %dl
  2014: 32da                     xorb %dl, %bl
  2016: 498b4070                 movq 0x70(%r8), %rax
  2020: 428a1418                 movb (%rax, %r11), %dl
  2024: 32da                     xorb %dl, %bl
  2026: 498b80a0000000           movq 0xa0(%r8), %rax
  2033: 428a1418                 movb (%rax, %r11), %dl
  2037: 32da                     xorb %dl, %bl
  2039: 498b80b8000000           movq 0xb8(%r8), %rax
  2046: 428a1418                 movb (%rax, %r11), %dl
  2050: 32da                     xorb %dl, %bl
  2052: 498b80d0000000           movq 0xd0(%r8), %rax
  2059: 428a1418                 movb (%rax, %r11), %dl
  2063: 32da                     xorb %dl, %bl
  2065: 498b80e0000000           movq 0xe0(%r8), %rax
  2072: 428a1418                 movb (%rax, %r11), %dl
  2076: 32da                     xorb %dl, %bl
  2078: 498b8000010000           movq 0x100(%r8), %rax
  2085: 428a1418                 movb (%rax, %r11), %dl
  2089: 32da                     xorb %dl, %bl
  2091: 498b8008010000           movq 0x108(%r8), %rax
  2098: 428a1418                 movb (%rax, %r11), %dl
  2102: 32da                     xorb %dl, %bl
  2104: 43881c1a                 movb %bl, (%r10, %r11)
  2108: 33db                     xorl %ebx, %ebx
  2110: 49ffc3                   incq %r11
  2113: 4c3bd9                   cmpq %rcx, %r11
  2116: 0f826effffff             jb 0x7b8
  2122: 4d8b5140                 movq 0x40(%r9), %r10
  2126: 4c8b1c24                 movq (%rsp), %r11
  2130: 498b4010                 movq 0x10(%r8), %rax
  2134: 428a1418                 movb (%rax, %r11), %dl
  2138: 8ada                     movb %dl, %bl
  2140: 498b4018                 movq 0x18(%r8), %rax
  2144: 428a1418                 movb (%rax, %r11), %dl
  2148: 32da                     xorb %dl, %bl
  2150: 498b4038                 movq 0x38(%r8), %rax
  2154: 428a1418                 movb (%rax, %r11), %dl
  2158: 32da                     xorb %dl, %bl
  2160: 498b4048                 movq 0x48(%r8), %rax
  2164: 428a1418                 movb (%rax, %r11), %dl
  2168: 32da                     xorb %dl, %bl
  2170: 498b4060                 movq 0x60(%r8), %rax
  2174: 428a1418                 movb (%rax, %r11), %dl
  2178: 32da                     xorb %dl, %bl
  2180: 498b4078                 movq 0x78(%r8), %rax
  2184: 428a1418                 movb (%rax, %r11), %dl
  2188: 32da                     xorb %dl, %bl
  2190: 498b80a8000000           movq 0xa8(%r8), %rax
  2197: 428a1418                 movb (%rax, %r11), %dl
  2201: 32da                     xorb %dl, %bl
  2203: 498b80d8000000           movq 0xd8(%r8), %rax
  2210: 428a1418                 movb (%rax, %r11), %dl
  2214: 32da                     xorb %dl, %bl
  2216: 498b80e8000000           movq 0xe8(%r8), %rax
  2223: 428a1418                 movb (%rax, %r11), %dl
  2227: 32da                     xorb %dl, %bl
  2229: 498b8008010000           movq 0x108(%r8), %rax
  2236: 428a1418                 movb (%rax, %r11), %dl
  2240: 32da                     xorb %dl, %bl
  2242: 498b8010010000           movq 0x110(%r8), %rax
  2249: 428a1418                 movb (%rax, %r11), %dl
  2253: 32da                     xorb %dl, %bl
  2255: 43881c1a                 movb %bl, (%r10, %r11)
  2259: 33db                     xorl %ebx, %ebx
  2261: 49ffc3                   incq %r11
  2264: 4c3bd9                   cmpq %rcx, %r11
  2267: 0f8271ffffff             jb 0x852
  2273: 4d8b5148                 movq 0x48(%r9), %r10
  2277: 4c8b1c24                 movq (%rsp), %r11
  2281: 498b4018                 movq 0x18(%r8), %rax
  2285: 428a1418                 movb (%rax, %r11), %dl
  2289: 8ada                     movb %dl, %bl
  2291: 498b4020                 movq 0x20(%r8), %rax
  2295: 428a1418                 movb (%rax, %r11), %dl
  2299: 32da                     xorb %dl, %bl
  2301: 498b4040                 movq 0x40(%r8), %rax
  2305: 428a1418                 movb (%rax, %r11), %dl
  2309: 32da                     xorb %dl, %bl
  2311: 498b4050                 movq 0x50(%r8), %rax
  2315: 428a1418                 movb (%rax, %r11), %dl
  2319: 32da                     xorb %dl, %bl
  2321: 498b4068                 movq 0x68(%r8), %rax
  2325: 428a1418                 movb (%rax, %r11), %dl
  2329: 32da                     xorb %dl, %bl
  2331: 498b8080000000           movq 0x80(%r8), %rax
  2338: 428a1418                 movb (%rax, %r11), %dl
  2342: 32da                     xorb %dl, %bl
  2344: 498b8090000000           movq 0x90(%r8), %rax
  2351: 428a1418                 movb (%rax, %r11), %dl
  2355: 32da                     xorb %dl, %bl
  2357: 498b80b0000000           movq 0xb0(%r8), %rax
  2364: 428a1418                 movb (%rax, %r11), %dl
  2368: 32da                     xorb %dl, %bl
  2370: 498b80e0000000           movq 0xe0(%r8), %rax
  2377: 428a1418                 movb (%rax, %r11), %dl
  2381: 32da                     xorb %dl, %bl
  2383: 498b8010010000           movq 0x110(%r8), %rax
  2390: 428a1418                 movb (%rax, %r11), %dl
  2394: 32da                     xorb %dl, %bl
  2396: 498b8018010000           movq 0x118(%r8), %rax
  2403: 428a1418                 movb (%rax, %r11), %dl
  2407: 32da                     xorb %dl, %bl
  2409: 43881c1a                 movb %bl, (%r10, %r11)
  2413: 33db                     xorl %ebx, %ebx
  2415: 49ffc3                   incq %r11
  2418: 4c3bd9                   cmpq %rcx, %r11
  2421: 0f826effffff             jb 0x8e9
  2427: 4d8b5150                 movq 0x50(%r9), %r10
  2431: 4c8b1c24                 movq (%rsp), %r11
  2435: 498b4020                 movq 0x20(%r8), %rax
  2439: 428a1418                 movb (%rax, %r11), %dl
  2443: 8ada                     movb %dl, %bl
  2445: 498b4028                 movq 0x28(%r8), %rax
  2449: 428a1418                 movb (%rax, %r11), %dl
  2453: 32da                     xorb %dl, %bl
  2455: 498b4048                 movq 0x48(%r8), %rax
  2459: 428a1418                 movb (%rax, %r11), %dl
  2463: 32da                     xorb %dl, %bl
  2465: 498b4058                 movq 0x58(%r8), %rax
  2469: 428a1418                 movb (%rax, %r11), %dl
  2473: 32da                     xorb %dl, %bl
  2475: 498b4070                 movq 0x70(%r8), %rax
  2479: 428a1418                 movb (%rax, %r11), %dl
  2483: 32da                     xorb %dl, %bl
  2485: 498b8088000000           movq 0x88(%r8), %rax
  2492: 428a1418                 movb (%rax, %r11), %dl
  2496: 32da                     xorb %dl, %bl
  2498: 498b8098000000           movq 0x98(%r8), %rax
  2505: 428a1418                 movb (%rax, %r11), %dl
  2509: 32da                     xorb %dl, %bl
  2511: 498b80b8000000           movq 0xb8(%r8), %rax
  2518: 428a1418                 movb (%rax, %r11), %dl
  2522: 32da                     xorb %dl, %bl
  2524: 498b80c0000000           movq 0xc0(%r8), %rax
  2531: 428a1418                 movb (%rax, %r11), %dl
  2535: 32da                     xorb %dl, %bl
  2537: 498b80e8000000           movq 0xe8(%r8), %rax
  2544: 428a1418                 movb (%rax, %r11), %dl
  2548: 32da                     xorb %dl, %bl
  2550: 498b8018010000           movq 0x118(%r8), %rax
  2557: 428a1418                 movb (%rax, %r11), %dl
  2561: 32da                     xorb %dl, %bl
  2563: 43881c1a                 movb %bl, (%r10, %r11)
  2567: 33db                     xorl %ebx, %ebx
  2569: 49ffc3                   incq %r11
  2572: 4c3bd9                   cmpq %rcx, %r11
  2575: 0f826effffff             jb 0x983
  2581: 4d8b5158                 movq 0x58(%r9), %r10
  2585: 4c8b1c24                 movq (%rsp), %r11
  2589: 498b4028                 movq 0x28(%r8), %rax
  2593: 428a1418                 movb (%rax, %r11), %dl
  2597: 8ada                     movb %dl, %bl
  2599: 498b4050                 movq 0x50(%r8), %rax
  2603: 428a1418                 movb (%rax, %r11), %dl
  2607: 32da                     xorb %dl, %bl
  2609: 498b4078                 movq 0x78(%r8), %rax
  2613: 428a1418                 movb (%rax, %r11), %dl
  2617: 32da                     xorb %dl, %bl
  2619: 498b80a0000000           movq 0xa0(%r8), %rax
  2626: 428a1418                 movb (%rax, %r11), %dl
  2630: 32da                     xorb %dl, %bl
  2632: 498b80c8000000           movq 0xc8(%r8), %rax
  2639: 428a1418                 movb (%rax, %r11), %dl
  2643: 32da                     xorb %dl, %bl
  2645: 498b80f0000000           movq 0xf0(%r8), %rax
  2652: 428a1418                 movb (%rax, %r11), %dl
  2656: 32da                     xorb %dl, %bl
  2658: 43881c1a                 movb %bl, (%r10, %r11)
  2662: 33db                     xorl %ebx, %ebx
  2664: 49ffc3                   incq %r11
  2667: 4c3bd9                   cmpq %rcx, %r11
  2670: 0f82a9ffffff             jb 0xa1d
  2676: 4883c408                 addq $8, %rsp
  2680: 415e                     popq %r14
  2682: 415d                     popq %r13
  2684: 415c                     popq %r12
  2686: 5b                       popq %rbx
  2687: c5f877                   vzeroupper 
  2690: c3                       retq 
  2691: 0f0b                     ud2 
  2693: 0f0b                     ud2 
  2695: 0f0b                     ud2 
  2697: 0f0b                     ud2 
  2699: 0f0b                     ud2 
  2701: 0f0b                     ud2 
  2703: 0f0b                     ud2 
  2705: 0f0b                     ud2 
  2707: 0f0b                     ud2 
  2709: 0f0b                     ud2 
  2711: 0f0b                     ud2 
  2713: 0f0b                     ud2 
  2715: 0f0b                     ud2 
  2717: 0f0b                     ud2 
  2719: 0f0b                     ud2 
  2721: 0f0b                     ud2 
  2723: 0f0b                     ud2 
  2725: 0f0b                     ud2 
  2727: 0f0b                     ud2 
  2729: 0f0b                     ud2 
  2731: 0f0b                     ud2 
  2733: 0f0b                     ud2 
  2735: 0f0b                     ud2 
  2737: 0f0b                     ud2 
  2739: 0f0b                     ud2 
  2741: 0f0b                     ud2 
  2743: 0f0b                     ud2 
  2745: 0f0b                     ud2 
  2747: 0f0b                     ud2 
  2749: 0f0b                     ud2 
  2751: cc                       int3 
