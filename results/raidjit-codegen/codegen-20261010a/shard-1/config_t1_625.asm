; config_t1_625 target=sse twin=nt topo=rdp(7)/enc op=encode U=1
; code_size=2560 twins=["nt"] uops_per_byte=11
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
    41: 4883e810                 subq $0x10, %rax
    45: 0f8207040000             jb 0x43a
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 66460f6f241e             movdqa (%rsi, %r11), %xmm12
    70: 66410f6fc4               movdqa %xmm12, %xmm0
    75: 66410f6ff4               movdqa %xmm12, %xmm6
    80: 66460f6f241f             movdqa (%rdi, %r11), %xmm12
    86: 66410f6fcc               movdqa %xmm12, %xmm1
    91: 66410feff4               pxor %xmm12, %xmm6
    96: 66410f6ffc               movdqa %xmm12, %xmm7
   101: 66470f6f241c             movdqa (%r12, %r11), %xmm12
   107: 66410f6fd4               movdqa %xmm12, %xmm2
   112: 66410feffc               pxor %xmm12, %xmm7
   117: 66450f6fc4               movdqa %xmm12, %xmm8
   122: 66470f6f641d00           movdqa (%r13, %r11), %xmm12
   129: 66410f6fdc               movdqa %xmm12, %xmm3
   134: 66450fefc4               pxor %xmm12, %xmm8
   139: 66450f6fcc               movdqa %xmm12, %xmm9
   144: 66470f6f241e             movdqa (%r14, %r11), %xmm12
   150: 66410f6fe4               movdqa %xmm12, %xmm4
   155: 66450fefcc               pxor %xmm12, %xmm9
   160: 66450f6fd4               movdqa %xmm12, %xmm10
   165: 4d8b5028                 movq 0x28(%r8), %r10
   169: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   175: 66410f6fec               movdqa %xmm12, %xmm5
   180: 66450fefd4               pxor %xmm12, %xmm10
   185: 66450f6fdc               movdqa %xmm12, %xmm11
   190: 4d8b5030                 movq 0x30(%r8), %r10
   194: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   200: 66410fefc4               pxor %xmm12, %xmm0
   205: 66410feffc               pxor %xmm12, %xmm7
   210: 4d8b5038                 movq 0x38(%r8), %r10
   214: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   220: 66410fefcc               pxor %xmm12, %xmm1
   225: 66410feff4               pxor %xmm12, %xmm6
   230: 66450fefc4               pxor %xmm12, %xmm8
   235: 4d8b5040                 movq 0x40(%r8), %r10
   239: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   245: 66410fefd4               pxor %xmm12, %xmm2
   250: 66410feffc               pxor %xmm12, %xmm7
   255: 66450fefcc               pxor %xmm12, %xmm9
   260: 4d8b5048                 movq 0x48(%r8), %r10
   264: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   270: 66410fefdc               pxor %xmm12, %xmm3
   275: 66450fefc4               pxor %xmm12, %xmm8
   280: 66450fefd4               pxor %xmm12, %xmm10
   285: 4d8b5050                 movq 0x50(%r8), %r10
   289: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   295: 66410fefe4               pxor %xmm12, %xmm4
   300: 66450fefcc               pxor %xmm12, %xmm9
   305: 66450fefdc               pxor %xmm12, %xmm11
   310: 4d8b5058                 movq 0x58(%r8), %r10
   314: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   320: 66410fefec               pxor %xmm12, %xmm5
   325: 66450fefd4               pxor %xmm12, %xmm10
   330: 4d8b5060                 movq 0x60(%r8), %r10
   334: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   340: 66410fefc4               pxor %xmm12, %xmm0
   345: 66450fefc4               pxor %xmm12, %xmm8
   350: 4d8b5068                 movq 0x68(%r8), %r10
   354: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   360: 66410fefcc               pxor %xmm12, %xmm1
   365: 66410feff4               pxor %xmm12, %xmm6
   370: 66450fefcc               pxor %xmm12, %xmm9
   375: 4d8b5070                 movq 0x70(%r8), %r10
   379: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   385: 66410fefd4               pxor %xmm12, %xmm2
   390: 66410feffc               pxor %xmm12, %xmm7
   395: 66450fefd4               pxor %xmm12, %xmm10
   400: 4d8b5078                 movq 0x78(%r8), %r10
   404: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   410: 66410fefdc               pxor %xmm12, %xmm3
   415: 66450fefc4               pxor %xmm12, %xmm8
   420: 66450fefdc               pxor %xmm12, %xmm11
   425: 4d8b9080000000           movq 0x80(%r8), %r10
   432: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   438: 66410fefe4               pxor %xmm12, %xmm4
   443: 66450fefcc               pxor %xmm12, %xmm9
   448: 4d8b9088000000           movq 0x88(%r8), %r10
   455: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   461: 66410fefec               pxor %xmm12, %xmm5
   466: 66410feff4               pxor %xmm12, %xmm6
   471: 66450fefd4               pxor %xmm12, %xmm10
   476: 4d8b9090000000           movq 0x90(%r8), %r10
   483: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   489: 66410fefc4               pxor %xmm12, %xmm0
   494: 66450fefcc               pxor %xmm12, %xmm9
   499: 4d8b9098000000           movq 0x98(%r8), %r10
   506: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   512: 66410fefcc               pxor %xmm12, %xmm1
   517: 66410feff4               pxor %xmm12, %xmm6
   522: 66450fefd4               pxor %xmm12, %xmm10
   527: 4d8b90a0000000           movq 0xa0(%r8), %r10
   534: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   540: 66410fefd4               pxor %xmm12, %xmm2
   545: 66410feffc               pxor %xmm12, %xmm7
   550: 66450fefdc               pxor %xmm12, %xmm11
   555: 4d8b90a8000000           movq 0xa8(%r8), %r10
   562: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   568: 66410fefdc               pxor %xmm12, %xmm3
   573: 66450fefc4               pxor %xmm12, %xmm8
   578: 4d8b90b0000000           movq 0xb0(%r8), %r10
   585: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   591: 66410fefe4               pxor %xmm12, %xmm4
   596: 66410feff4               pxor %xmm12, %xmm6
   601: 66450fefcc               pxor %xmm12, %xmm9
   606: 4d8b90b8000000           movq 0xb8(%r8), %r10
   613: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   619: 66410fefec               pxor %xmm12, %xmm5
   624: 66410feffc               pxor %xmm12, %xmm7
   629: 66450fefd4               pxor %xmm12, %xmm10
   634: 4d8b90c0000000           movq 0xc0(%r8), %r10
   641: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   647: 66410fefc4               pxor %xmm12, %xmm0
   652: 66450fefd4               pxor %xmm12, %xmm10
   657: 4d8b90c8000000           movq 0xc8(%r8), %r10
   664: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   670: 66410fefcc               pxor %xmm12, %xmm1
   675: 66410feff4               pxor %xmm12, %xmm6
   680: 66450fefdc               pxor %xmm12, %xmm11
   685: 4d8b90d0000000           movq 0xd0(%r8), %r10
   692: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   698: 66410fefd4               pxor %xmm12, %xmm2
   703: 66410feffc               pxor %xmm12, %xmm7
   708: 4d8b90d8000000           movq 0xd8(%r8), %r10
   715: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   721: 66410fefdc               pxor %xmm12, %xmm3
   726: 66410feff4               pxor %xmm12, %xmm6
   731: 66450fefc4               pxor %xmm12, %xmm8
   736: 4d8b90e0000000           movq 0xe0(%r8), %r10
   743: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   749: 66410fefe4               pxor %xmm12, %xmm4
   754: 66410feffc               pxor %xmm12, %xmm7
   759: 66450fefcc               pxor %xmm12, %xmm9
   764: 4d8b90e8000000           movq 0xe8(%r8), %r10
   771: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   777: 66410fefec               pxor %xmm12, %xmm5
   782: 66450fefc4               pxor %xmm12, %xmm8
   787: 66450fefd4               pxor %xmm12, %xmm10
   792: 4d8b90f0000000           movq 0xf0(%r8), %r10
   799: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   805: 66410fefc4               pxor %xmm12, %xmm0
   810: 66450fefdc               pxor %xmm12, %xmm11
   815: 4d8b90f8000000           movq 0xf8(%r8), %r10
   822: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   828: 66410fefcc               pxor %xmm12, %xmm1
   833: 66410feff4               pxor %xmm12, %xmm6
   838: 4d8b9000010000           movq 0x100(%r8), %r10
   845: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   851: 66410fefd4               pxor %xmm12, %xmm2
   856: 66410feff4               pxor %xmm12, %xmm6
   861: 66410feffc               pxor %xmm12, %xmm7
   866: 4d8b9008010000           movq 0x108(%r8), %r10
   873: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   879: 66410fefdc               pxor %xmm12, %xmm3
   884: 66410feffc               pxor %xmm12, %xmm7
   889: 66450fefc4               pxor %xmm12, %xmm8
   894: 4d8b9010010000           movq 0x110(%r8), %r10
   901: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   907: 66410fefe4               pxor %xmm12, %xmm4
   912: 66450fefc4               pxor %xmm12, %xmm8
   917: 66450fefcc               pxor %xmm12, %xmm9
   922: 4d8b9018010000           movq 0x118(%r8), %r10
   929: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   935: 66410fefec               pxor %xmm12, %xmm5
   940: 66450fefcc               pxor %xmm12, %xmm9
   945: 66450fefd4               pxor %xmm12, %xmm10
   950: 4d8b11                   movq (%r9), %r10
   953: 66430fe7041a             movntdq %xmm0, (%r10, %r11)
   959: 4d8b5108                 movq 8(%r9), %r10
   963: 66430fe70c1a             movntdq %xmm1, (%r10, %r11)
   969: 4d8b5110                 movq 0x10(%r9), %r10
   973: 66430fe7141a             movntdq %xmm2, (%r10, %r11)
   979: 4d8b5118                 movq 0x18(%r9), %r10
   983: 66430fe71c1a             movntdq %xmm3, (%r10, %r11)
   989: 4d8b5120                 movq 0x20(%r9), %r10
   993: 66430fe7241a             movntdq %xmm4, (%r10, %r11)
   999: 4d8b5128                 movq 0x28(%r9), %r10
  1003: 66430fe72c1a             movntdq %xmm5, (%r10, %r11)
  1009: 4d8b5130                 movq 0x30(%r9), %r10
  1013: 66430fe7341a             movntdq %xmm6, (%r10, %r11)
  1019: 4d8b5138                 movq 0x38(%r9), %r10
  1023: 66430fe73c1a             movntdq %xmm7, (%r10, %r11)
  1029: 4d8b5140                 movq 0x40(%r9), %r10
  1033: 66470fe7041a             movntdq %xmm8, (%r10, %r11)
  1039: 4d8b5148                 movq 0x48(%r9), %r10
  1043: 66470fe70c1a             movntdq %xmm9, (%r10, %r11)
  1049: 4d8b5150                 movq 0x50(%r9), %r10
  1053: 66470fe7141a             movntdq %xmm10, (%r10, %r11)
  1059: 4d8b5158                 movq 0x58(%r9), %r10
  1063: 66470fe71c1a             movntdq %xmm11, (%r10, %r11)
  1069: 4983c310                 addq $0x10, %r11
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
  2542: c3                       retq 
  2543: 0f0b                     ud2 
  2545: 0f0b                     ud2 
  2547: 0f0b                     ud2 
  2549: 0f0b                     ud2 
  2551: 0f0b                     ud2 
  2553: 0f0b                     ud2 
  2555: 0f0b                     ud2 
  2557: 0f0b                     ud2 
  2559: cc                       int3 
