; config_t3_2108 target=sse twin=baseline topo=lrc_azure()/rec{0} op=recover U=4
; code_size=4288 twins=[] uops_per_byte=9.921875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da7100000           leaq 0x10a7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66450f6f17               movdqa (%r15), %xmm10
    52: 488bc1                   movq %rcx, %rax
    55: 4883e840                 subq $0x40, %rax
    59: 0f82e30d0000             jb 0xe24
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66420f6f241e             movdqa (%rsi, %r11), %xmm4
   102: 660f6fec                 movdqa %xmm4, %xmm5
   106: 66410fdbea               pand %xmm10, %xmm5
   111: 660f6ff4                 movdqa %xmm4, %xmm6
   115: 660f71d604               psrlw $4, %xmm6
   120: 66410fdbf2               pand %xmm10, %xmm6
   125: 66450f6f4f40             movdqa 0x40(%r15), %xmm9
   131: 66410f6ff9               movdqa %xmm9, %xmm7
   136: 660f3800fd               pshufb %xmm5, %xmm7
   141: 66450f6f8f80000000       movdqa 0x80(%r15), %xmm9
   150: 66450f6fc1               movdqa %xmm9, %xmm8
   155: 66440f3800c6             pshufb %xmm6, %xmm8
   161: 660f6fc7                 movdqa %xmm7, %xmm0
   165: 66410fefc0               pxor %xmm8, %xmm0
   170: 66420f6f641e10           movdqa 0x10(%rsi, %r11), %xmm4
   177: 660f6fec                 movdqa %xmm4, %xmm5
   181: 66410fdbea               pand %xmm10, %xmm5
   186: 660f6ff4                 movdqa %xmm4, %xmm6
   190: 660f71d604               psrlw $4, %xmm6
   195: 66410fdbf2               pand %xmm10, %xmm6
   200: 66450f6f4f40             movdqa 0x40(%r15), %xmm9
   206: 66410f6ff9               movdqa %xmm9, %xmm7
   211: 660f3800fd               pshufb %xmm5, %xmm7
   216: 66450f6f8f80000000       movdqa 0x80(%r15), %xmm9
   225: 66450f6fc1               movdqa %xmm9, %xmm8
   230: 66440f3800c6             pshufb %xmm6, %xmm8
   236: 660f6fcf                 movdqa %xmm7, %xmm1
   240: 66410fefc8               pxor %xmm8, %xmm1
   245: 66420f6f641e20           movdqa 0x20(%rsi, %r11), %xmm4
   252: 660f6fec                 movdqa %xmm4, %xmm5
   256: 66410fdbea               pand %xmm10, %xmm5
   261: 660f6ff4                 movdqa %xmm4, %xmm6
   265: 660f71d604               psrlw $4, %xmm6
   270: 66410fdbf2               pand %xmm10, %xmm6
   275: 66450f6f4f40             movdqa 0x40(%r15), %xmm9
   281: 66410f6ff9               movdqa %xmm9, %xmm7
   286: 660f3800fd               pshufb %xmm5, %xmm7
   291: 66450f6f8f80000000       movdqa 0x80(%r15), %xmm9
   300: 66450f6fc1               movdqa %xmm9, %xmm8
   305: 66440f3800c6             pshufb %xmm6, %xmm8
   311: 660f6fd7                 movdqa %xmm7, %xmm2
   315: 66410fefd0               pxor %xmm8, %xmm2
   320: 66420f6f641e30           movdqa 0x30(%rsi, %r11), %xmm4
   327: 660f6fec                 movdqa %xmm4, %xmm5
   331: 66410fdbea               pand %xmm10, %xmm5
   336: 660f6ff4                 movdqa %xmm4, %xmm6
   340: 660f71d604               psrlw $4, %xmm6
   345: 66410fdbf2               pand %xmm10, %xmm6
   350: 66450f6f4f40             movdqa 0x40(%r15), %xmm9
   356: 66410f6ff9               movdqa %xmm9, %xmm7
   361: 660f3800fd               pshufb %xmm5, %xmm7
   366: 66450f6f8f80000000       movdqa 0x80(%r15), %xmm9
   375: 66450f6fc1               movdqa %xmm9, %xmm8
   380: 66440f3800c6             pshufb %xmm6, %xmm8
   386: 660f6fdf                 movdqa %xmm7, %xmm3
   390: 66410fefd8               pxor %xmm8, %xmm3
   395: 66420f6f241f             movdqa (%rdi, %r11), %xmm4
   401: 660f6fec                 movdqa %xmm4, %xmm5
   405: 66410fdbea               pand %xmm10, %xmm5
   410: 660f6ff4                 movdqa %xmm4, %xmm6
   414: 660f71d604               psrlw $4, %xmm6
   419: 66410fdbf2               pand %xmm10, %xmm6
   424: 66450f6f8fc0000000       movdqa 0xc0(%r15), %xmm9
   433: 66410f6ff9               movdqa %xmm9, %xmm7
   438: 660f3800fd               pshufb %xmm5, %xmm7
   443: 66450f6f8f00010000       movdqa 0x100(%r15), %xmm9
   452: 66450f6fc1               movdqa %xmm9, %xmm8
   457: 66440f3800c6             pshufb %xmm6, %xmm8
   463: 660fefc7                 pxor %xmm7, %xmm0
   467: 66410fefc0               pxor %xmm8, %xmm0
   472: 66420f6f641f10           movdqa 0x10(%rdi, %r11), %xmm4
   479: 660f6fec                 movdqa %xmm4, %xmm5
   483: 66410fdbea               pand %xmm10, %xmm5
   488: 660f6ff4                 movdqa %xmm4, %xmm6
   492: 660f71d604               psrlw $4, %xmm6
   497: 66410fdbf2               pand %xmm10, %xmm6
   502: 66450f6f8fc0000000       movdqa 0xc0(%r15), %xmm9
   511: 66410f6ff9               movdqa %xmm9, %xmm7
   516: 660f3800fd               pshufb %xmm5, %xmm7
   521: 66450f6f8f00010000       movdqa 0x100(%r15), %xmm9
   530: 66450f6fc1               movdqa %xmm9, %xmm8
   535: 66440f3800c6             pshufb %xmm6, %xmm8
   541: 660fefcf                 pxor %xmm7, %xmm1
   545: 66410fefc8               pxor %xmm8, %xmm1
   550: 66420f6f641f20           movdqa 0x20(%rdi, %r11), %xmm4
   557: 660f6fec                 movdqa %xmm4, %xmm5
   561: 66410fdbea               pand %xmm10, %xmm5
   566: 660f6ff4                 movdqa %xmm4, %xmm6
   570: 660f71d604               psrlw $4, %xmm6
   575: 66410fdbf2               pand %xmm10, %xmm6
   580: 66450f6f8fc0000000       movdqa 0xc0(%r15), %xmm9
   589: 66410f6ff9               movdqa %xmm9, %xmm7
   594: 660f3800fd               pshufb %xmm5, %xmm7
   599: 66450f6f8f00010000       movdqa 0x100(%r15), %xmm9
   608: 66450f6fc1               movdqa %xmm9, %xmm8
   613: 66440f3800c6             pshufb %xmm6, %xmm8
   619: 660fefd7                 pxor %xmm7, %xmm2
   623: 66410fefd0               pxor %xmm8, %xmm2
   628: 66420f6f641f30           movdqa 0x30(%rdi, %r11), %xmm4
   635: 660f6fec                 movdqa %xmm4, %xmm5
   639: 66410fdbea               pand %xmm10, %xmm5
   644: 660f6ff4                 movdqa %xmm4, %xmm6
   648: 660f71d604               psrlw $4, %xmm6
   653: 66410fdbf2               pand %xmm10, %xmm6
   658: 66450f6f8fc0000000       movdqa 0xc0(%r15), %xmm9
   667: 66410f6ff9               movdqa %xmm9, %xmm7
   672: 660f3800fd               pshufb %xmm5, %xmm7
   677: 66450f6f8f00010000       movdqa 0x100(%r15), %xmm9
   686: 66450f6fc1               movdqa %xmm9, %xmm8
   691: 66440f3800c6             pshufb %xmm6, %xmm8
   697: 660fefdf                 pxor %xmm7, %xmm3
   701: 66410fefd8               pxor %xmm8, %xmm3
   706: 66430f6f241c             movdqa (%r12, %r11), %xmm4
   712: 660f6fec                 movdqa %xmm4, %xmm5
   716: 66410fdbea               pand %xmm10, %xmm5
   721: 660f6ff4                 movdqa %xmm4, %xmm6
   725: 660f71d604               psrlw $4, %xmm6
   730: 66410fdbf2               pand %xmm10, %xmm6
   735: 66450f6f8f40010000       movdqa 0x140(%r15), %xmm9
   744: 66410f6ff9               movdqa %xmm9, %xmm7
   749: 660f3800fd               pshufb %xmm5, %xmm7
   754: 66450f6f8f80010000       movdqa 0x180(%r15), %xmm9
   763: 66450f6fc1               movdqa %xmm9, %xmm8
   768: 66440f3800c6             pshufb %xmm6, %xmm8
   774: 660fefc7                 pxor %xmm7, %xmm0
   778: 66410fefc0               pxor %xmm8, %xmm0
   783: 66430f6f641c10           movdqa 0x10(%r12, %r11), %xmm4
   790: 660f6fec                 movdqa %xmm4, %xmm5
   794: 66410fdbea               pand %xmm10, %xmm5
   799: 660f6ff4                 movdqa %xmm4, %xmm6
   803: 660f71d604               psrlw $4, %xmm6
   808: 66410fdbf2               pand %xmm10, %xmm6
   813: 66450f6f8f40010000       movdqa 0x140(%r15), %xmm9
   822: 66410f6ff9               movdqa %xmm9, %xmm7
   827: 660f3800fd               pshufb %xmm5, %xmm7
   832: 66450f6f8f80010000       movdqa 0x180(%r15), %xmm9
   841: 66450f6fc1               movdqa %xmm9, %xmm8
   846: 66440f3800c6             pshufb %xmm6, %xmm8
   852: 660fefcf                 pxor %xmm7, %xmm1
   856: 66410fefc8               pxor %xmm8, %xmm1
   861: 66430f6f641c20           movdqa 0x20(%r12, %r11), %xmm4
   868: 660f6fec                 movdqa %xmm4, %xmm5
   872: 66410fdbea               pand %xmm10, %xmm5
   877: 660f6ff4                 movdqa %xmm4, %xmm6
   881: 660f71d604               psrlw $4, %xmm6
   886: 66410fdbf2               pand %xmm10, %xmm6
   891: 66450f6f8f40010000       movdqa 0x140(%r15), %xmm9
   900: 66410f6ff9               movdqa %xmm9, %xmm7
   905: 660f3800fd               pshufb %xmm5, %xmm7
   910: 66450f6f8f80010000       movdqa 0x180(%r15), %xmm9
   919: 66450f6fc1               movdqa %xmm9, %xmm8
   924: 66440f3800c6             pshufb %xmm6, %xmm8
   930: 660fefd7                 pxor %xmm7, %xmm2
   934: 66410fefd0               pxor %xmm8, %xmm2
   939: 66430f6f641c30           movdqa 0x30(%r12, %r11), %xmm4
   946: 660f6fec                 movdqa %xmm4, %xmm5
   950: 66410fdbea               pand %xmm10, %xmm5
   955: 660f6ff4                 movdqa %xmm4, %xmm6
   959: 660f71d604               psrlw $4, %xmm6
   964: 66410fdbf2               pand %xmm10, %xmm6
   969: 66450f6f8f40010000       movdqa 0x140(%r15), %xmm9
   978: 66410f6ff9               movdqa %xmm9, %xmm7
   983: 660f3800fd               pshufb %xmm5, %xmm7
   988: 66450f6f8f80010000       movdqa 0x180(%r15), %xmm9
   997: 66450f6fc1               movdqa %xmm9, %xmm8
  1002: 66440f3800c6             pshufb %xmm6, %xmm8
  1008: 660fefdf                 pxor %xmm7, %xmm3
  1012: 66410fefd8               pxor %xmm8, %xmm3
  1017: 66430f6f641d00           movdqa (%r13, %r11), %xmm4
  1024: 660f6fec                 movdqa %xmm4, %xmm5
  1028: 66410fdbea               pand %xmm10, %xmm5
  1033: 660f6ff4                 movdqa %xmm4, %xmm6
  1037: 660f71d604               psrlw $4, %xmm6
  1042: 66410fdbf2               pand %xmm10, %xmm6
  1047: 66450f6f8fc0010000       movdqa 0x1c0(%r15), %xmm9
  1056: 66410f6ff9               movdqa %xmm9, %xmm7
  1061: 660f3800fd               pshufb %xmm5, %xmm7
  1066: 66450f6f8f00020000       movdqa 0x200(%r15), %xmm9
  1075: 66450f6fc1               movdqa %xmm9, %xmm8
  1080: 66440f3800c6             pshufb %xmm6, %xmm8
  1086: 660fefc7                 pxor %xmm7, %xmm0
  1090: 66410fefc0               pxor %xmm8, %xmm0
  1095: 66430f6f641d10           movdqa 0x10(%r13, %r11), %xmm4
  1102: 660f6fec                 movdqa %xmm4, %xmm5
  1106: 66410fdbea               pand %xmm10, %xmm5
  1111: 660f6ff4                 movdqa %xmm4, %xmm6
  1115: 660f71d604               psrlw $4, %xmm6
  1120: 66410fdbf2               pand %xmm10, %xmm6
  1125: 66450f6f8fc0010000       movdqa 0x1c0(%r15), %xmm9
  1134: 66410f6ff9               movdqa %xmm9, %xmm7
  1139: 660f3800fd               pshufb %xmm5, %xmm7
  1144: 66450f6f8f00020000       movdqa 0x200(%r15), %xmm9
  1153: 66450f6fc1               movdqa %xmm9, %xmm8
  1158: 66440f3800c6             pshufb %xmm6, %xmm8
  1164: 660fefcf                 pxor %xmm7, %xmm1
  1168: 66410fefc8               pxor %xmm8, %xmm1
  1173: 66430f6f641d20           movdqa 0x20(%r13, %r11), %xmm4
  1180: 660f6fec                 movdqa %xmm4, %xmm5
  1184: 66410fdbea               pand %xmm10, %xmm5
  1189: 660f6ff4                 movdqa %xmm4, %xmm6
  1193: 660f71d604               psrlw $4, %xmm6
  1198: 66410fdbf2               pand %xmm10, %xmm6
  1203: 66450f6f8fc0010000       movdqa 0x1c0(%r15), %xmm9
  1212: 66410f6ff9               movdqa %xmm9, %xmm7
  1217: 660f3800fd               pshufb %xmm5, %xmm7
  1222: 66450f6f8f00020000       movdqa 0x200(%r15), %xmm9
  1231: 66450f6fc1               movdqa %xmm9, %xmm8
  1236: 66440f3800c6             pshufb %xmm6, %xmm8
  1242: 660fefd7                 pxor %xmm7, %xmm2
  1246: 66410fefd0               pxor %xmm8, %xmm2
  1251: 66430f6f641d30           movdqa 0x30(%r13, %r11), %xmm4
  1258: 660f6fec                 movdqa %xmm4, %xmm5
  1262: 66410fdbea               pand %xmm10, %xmm5
  1267: 660f6ff4                 movdqa %xmm4, %xmm6
  1271: 660f71d604               psrlw $4, %xmm6
  1276: 66410fdbf2               pand %xmm10, %xmm6
  1281: 66450f6f8fc0010000       movdqa 0x1c0(%r15), %xmm9
  1290: 66410f6ff9               movdqa %xmm9, %xmm7
  1295: 660f3800fd               pshufb %xmm5, %xmm7
  1300: 66450f6f8f00020000       movdqa 0x200(%r15), %xmm9
  1309: 66450f6fc1               movdqa %xmm9, %xmm8
  1314: 66440f3800c6             pshufb %xmm6, %xmm8
  1320: 660fefdf                 pxor %xmm7, %xmm3
  1324: 66410fefd8               pxor %xmm8, %xmm3
  1329: 66430f6f241e             movdqa (%r14, %r11), %xmm4
  1335: 660f6fec                 movdqa %xmm4, %xmm5
  1339: 66410fdbea               pand %xmm10, %xmm5
  1344: 660f6ff4                 movdqa %xmm4, %xmm6
  1348: 660f71d604               psrlw $4, %xmm6
  1353: 66410fdbf2               pand %xmm10, %xmm6
  1358: 66450f6f8f40020000       movdqa 0x240(%r15), %xmm9
  1367: 66410f6ff9               movdqa %xmm9, %xmm7
  1372: 660f3800fd               pshufb %xmm5, %xmm7
  1377: 66450f6f8f80020000       movdqa 0x280(%r15), %xmm9
  1386: 66450f6fc1               movdqa %xmm9, %xmm8
  1391: 66440f3800c6             pshufb %xmm6, %xmm8
  1397: 660fefc7                 pxor %xmm7, %xmm0
  1401: 66410fefc0               pxor %xmm8, %xmm0
  1406: 66430f6f641e10           movdqa 0x10(%r14, %r11), %xmm4
  1413: 660f6fec                 movdqa %xmm4, %xmm5
  1417: 66410fdbea               pand %xmm10, %xmm5
  1422: 660f6ff4                 movdqa %xmm4, %xmm6
  1426: 660f71d604               psrlw $4, %xmm6
  1431: 66410fdbf2               pand %xmm10, %xmm6
  1436: 66450f6f8f40020000       movdqa 0x240(%r15), %xmm9
  1445: 66410f6ff9               movdqa %xmm9, %xmm7
  1450: 660f3800fd               pshufb %xmm5, %xmm7
  1455: 66450f6f8f80020000       movdqa 0x280(%r15), %xmm9
  1464: 66450f6fc1               movdqa %xmm9, %xmm8
  1469: 66440f3800c6             pshufb %xmm6, %xmm8
  1475: 660fefcf                 pxor %xmm7, %xmm1
  1479: 66410fefc8               pxor %xmm8, %xmm1
  1484: 66430f6f641e20           movdqa 0x20(%r14, %r11), %xmm4
  1491: 660f6fec                 movdqa %xmm4, %xmm5
  1495: 66410fdbea               pand %xmm10, %xmm5
  1500: 660f6ff4                 movdqa %xmm4, %xmm6
  1504: 660f71d604               psrlw $4, %xmm6
  1509: 66410fdbf2               pand %xmm10, %xmm6
  1514: 66450f6f8f40020000       movdqa 0x240(%r15), %xmm9
  1523: 66410f6ff9               movdqa %xmm9, %xmm7
  1528: 660f3800fd               pshufb %xmm5, %xmm7
  1533: 66450f6f8f80020000       movdqa 0x280(%r15), %xmm9
  1542: 66450f6fc1               movdqa %xmm9, %xmm8
  1547: 66440f3800c6             pshufb %xmm6, %xmm8
  1553: 660fefd7                 pxor %xmm7, %xmm2
  1557: 66410fefd0               pxor %xmm8, %xmm2
  1562: 66430f6f641e30           movdqa 0x30(%r14, %r11), %xmm4
  1569: 660f6fec                 movdqa %xmm4, %xmm5
  1573: 66410fdbea               pand %xmm10, %xmm5
  1578: 660f6ff4                 movdqa %xmm4, %xmm6
  1582: 660f71d604               psrlw $4, %xmm6
  1587: 66410fdbf2               pand %xmm10, %xmm6
  1592: 66450f6f8f40020000       movdqa 0x240(%r15), %xmm9
  1601: 66410f6ff9               movdqa %xmm9, %xmm7
  1606: 660f3800fd               pshufb %xmm5, %xmm7
  1611: 66450f6f8f80020000       movdqa 0x280(%r15), %xmm9
  1620: 66450f6fc1               movdqa %xmm9, %xmm8
  1625: 66440f3800c6             pshufb %xmm6, %xmm8
  1631: 660fefdf                 pxor %xmm7, %xmm3
  1635: 66410fefd8               pxor %xmm8, %xmm3
  1640: 4d8b5028                 movq 0x28(%r8), %r10
  1644: 66430f6f241a             movdqa (%r10, %r11), %xmm4
  1650: 660f6fec                 movdqa %xmm4, %xmm5
  1654: 66410fdbea               pand %xmm10, %xmm5
  1659: 660f6ff4                 movdqa %xmm4, %xmm6
  1663: 660f71d604               psrlw $4, %xmm6
  1668: 66410fdbf2               pand %xmm10, %xmm6
  1673: 66450f6f8fc0020000       movdqa 0x2c0(%r15), %xmm9
  1682: 66410f6ff9               movdqa %xmm9, %xmm7
  1687: 660f3800fd               pshufb %xmm5, %xmm7
  1692: 66450f6f8f00030000       movdqa 0x300(%r15), %xmm9
  1701: 66450f6fc1               movdqa %xmm9, %xmm8
  1706: 66440f3800c6             pshufb %xmm6, %xmm8
  1712: 660fefc7                 pxor %xmm7, %xmm0
  1716: 66410fefc0               pxor %xmm8, %xmm0
  1721: 66430f6f641a10           movdqa 0x10(%r10, %r11), %xmm4
  1728: 660f6fec                 movdqa %xmm4, %xmm5
  1732: 66410fdbea               pand %xmm10, %xmm5
  1737: 660f6ff4                 movdqa %xmm4, %xmm6
  1741: 660f71d604               psrlw $4, %xmm6
  1746: 66410fdbf2               pand %xmm10, %xmm6
  1751: 66450f6f8fc0020000       movdqa 0x2c0(%r15), %xmm9
  1760: 66410f6ff9               movdqa %xmm9, %xmm7
  1765: 660f3800fd               pshufb %xmm5, %xmm7
  1770: 66450f6f8f00030000       movdqa 0x300(%r15), %xmm9
  1779: 66450f6fc1               movdqa %xmm9, %xmm8
  1784: 66440f3800c6             pshufb %xmm6, %xmm8
  1790: 660fefcf                 pxor %xmm7, %xmm1
  1794: 66410fefc8               pxor %xmm8, %xmm1
  1799: 66430f6f641a20           movdqa 0x20(%r10, %r11), %xmm4
  1806: 660f6fec                 movdqa %xmm4, %xmm5
  1810: 66410fdbea               pand %xmm10, %xmm5
  1815: 660f6ff4                 movdqa %xmm4, %xmm6
  1819: 660f71d604               psrlw $4, %xmm6
  1824: 66410fdbf2               pand %xmm10, %xmm6
  1829: 66450f6f8fc0020000       movdqa 0x2c0(%r15), %xmm9
  1838: 66410f6ff9               movdqa %xmm9, %xmm7
  1843: 660f3800fd               pshufb %xmm5, %xmm7
  1848: 66450f6f8f00030000       movdqa 0x300(%r15), %xmm9
  1857: 66450f6fc1               movdqa %xmm9, %xmm8
  1862: 66440f3800c6             pshufb %xmm6, %xmm8
  1868: 660fefd7                 pxor %xmm7, %xmm2
  1872: 66410fefd0               pxor %xmm8, %xmm2
  1877: 66430f6f641a30           movdqa 0x30(%r10, %r11), %xmm4
  1884: 660f6fec                 movdqa %xmm4, %xmm5
  1888: 66410fdbea               pand %xmm10, %xmm5
  1893: 660f6ff4                 movdqa %xmm4, %xmm6
  1897: 660f71d604               psrlw $4, %xmm6
  1902: 66410fdbf2               pand %xmm10, %xmm6
  1907: 66450f6f8fc0020000       movdqa 0x2c0(%r15), %xmm9
  1916: 66410f6ff9               movdqa %xmm9, %xmm7
  1921: 660f3800fd               pshufb %xmm5, %xmm7
  1926: 66450f6f8f00030000       movdqa 0x300(%r15), %xmm9
  1935: 66450f6fc1               movdqa %xmm9, %xmm8
  1940: 66440f3800c6             pshufb %xmm6, %xmm8
  1946: 660fefdf                 pxor %xmm7, %xmm3
  1950: 66410fefd8               pxor %xmm8, %xmm3
  1955: 4d8b5030                 movq 0x30(%r8), %r10
  1959: 66430f6f241a             movdqa (%r10, %r11), %xmm4
  1965: 660f6fec                 movdqa %xmm4, %xmm5
  1969: 66410fdbea               pand %xmm10, %xmm5
  1974: 660f6ff4                 movdqa %xmm4, %xmm6
  1978: 660f71d604               psrlw $4, %xmm6
  1983: 66410fdbf2               pand %xmm10, %xmm6
  1988: 66450f6f8f40030000       movdqa 0x340(%r15), %xmm9
  1997: 66410f6ff9               movdqa %xmm9, %xmm7
  2002: 660f3800fd               pshufb %xmm5, %xmm7
  2007: 66450f6f8f80030000       movdqa 0x380(%r15), %xmm9
  2016: 66450f6fc1               movdqa %xmm9, %xmm8
  2021: 66440f3800c6             pshufb %xmm6, %xmm8
  2027: 660fefc7                 pxor %xmm7, %xmm0
  2031: 66410fefc0               pxor %xmm8, %xmm0
  2036: 66430f6f641a10           movdqa 0x10(%r10, %r11), %xmm4
  2043: 660f6fec                 movdqa %xmm4, %xmm5
  2047: 66410fdbea               pand %xmm10, %xmm5
  2052: 660f6ff4                 movdqa %xmm4, %xmm6
  2056: 660f71d604               psrlw $4, %xmm6
  2061: 66410fdbf2               pand %xmm10, %xmm6
  2066: 66450f6f8f40030000       movdqa 0x340(%r15), %xmm9
  2075: 66410f6ff9               movdqa %xmm9, %xmm7
  2080: 660f3800fd               pshufb %xmm5, %xmm7
  2085: 66450f6f8f80030000       movdqa 0x380(%r15), %xmm9
  2094: 66450f6fc1               movdqa %xmm9, %xmm8
  2099: 66440f3800c6             pshufb %xmm6, %xmm8
  2105: 660fefcf                 pxor %xmm7, %xmm1
  2109: 66410fefc8               pxor %xmm8, %xmm1
  2114: 66430f6f641a20           movdqa 0x20(%r10, %r11), %xmm4
  2121: 660f6fec                 movdqa %xmm4, %xmm5
  2125: 66410fdbea               pand %xmm10, %xmm5
  2130: 660f6ff4                 movdqa %xmm4, %xmm6
  2134: 660f71d604               psrlw $4, %xmm6
  2139: 66410fdbf2               pand %xmm10, %xmm6
  2144: 66450f6f8f40030000       movdqa 0x340(%r15), %xmm9
  2153: 66410f6ff9               movdqa %xmm9, %xmm7
  2158: 660f3800fd               pshufb %xmm5, %xmm7
  2163: 66450f6f8f80030000       movdqa 0x380(%r15), %xmm9
  2172: 66450f6fc1               movdqa %xmm9, %xmm8
  2177: 66440f3800c6             pshufb %xmm6, %xmm8
  2183: 660fefd7                 pxor %xmm7, %xmm2
  2187: 66410fefd0               pxor %xmm8, %xmm2
  2192: 66430f6f641a30           movdqa 0x30(%r10, %r11), %xmm4
  2199: 660f6fec                 movdqa %xmm4, %xmm5
  2203: 66410fdbea               pand %xmm10, %xmm5
  2208: 660f6ff4                 movdqa %xmm4, %xmm6
  2212: 660f71d604               psrlw $4, %xmm6
  2217: 66410fdbf2               pand %xmm10, %xmm6
  2222: 66450f6f8f40030000       movdqa 0x340(%r15), %xmm9
  2231: 66410f6ff9               movdqa %xmm9, %xmm7
  2236: 660f3800fd               pshufb %xmm5, %xmm7
  2241: 66450f6f8f80030000       movdqa 0x380(%r15), %xmm9
  2250: 66450f6fc1               movdqa %xmm9, %xmm8
  2255: 66440f3800c6             pshufb %xmm6, %xmm8
  2261: 660fefdf                 pxor %xmm7, %xmm3
  2265: 66410fefd8               pxor %xmm8, %xmm3
  2270: 4d8b5038                 movq 0x38(%r8), %r10
  2274: 66430f6f241a             movdqa (%r10, %r11), %xmm4
  2280: 660f6fec                 movdqa %xmm4, %xmm5
  2284: 66410fdbea               pand %xmm10, %xmm5
  2289: 660f6ff4                 movdqa %xmm4, %xmm6
  2293: 660f71d604               psrlw $4, %xmm6
  2298: 66410fdbf2               pand %xmm10, %xmm6
  2303: 66450f6f8fc0030000       movdqa 0x3c0(%r15), %xmm9
  2312: 66410f6ff9               movdqa %xmm9, %xmm7
  2317: 660f3800fd               pshufb %xmm5, %xmm7
  2322: 66450f6f8f00040000       movdqa 0x400(%r15), %xmm9
  2331: 66450f6fc1               movdqa %xmm9, %xmm8
  2336: 66440f3800c6             pshufb %xmm6, %xmm8
  2342: 660fefc7                 pxor %xmm7, %xmm0
  2346: 66410fefc0               pxor %xmm8, %xmm0
  2351: 66430f6f641a10           movdqa 0x10(%r10, %r11), %xmm4
  2358: 660f6fec                 movdqa %xmm4, %xmm5
  2362: 66410fdbea               pand %xmm10, %xmm5
  2367: 660f6ff4                 movdqa %xmm4, %xmm6
  2371: 660f71d604               psrlw $4, %xmm6
  2376: 66410fdbf2               pand %xmm10, %xmm6
  2381: 66450f6f8fc0030000       movdqa 0x3c0(%r15), %xmm9
  2390: 66410f6ff9               movdqa %xmm9, %xmm7
  2395: 660f3800fd               pshufb %xmm5, %xmm7
  2400: 66450f6f8f00040000       movdqa 0x400(%r15), %xmm9
  2409: 66450f6fc1               movdqa %xmm9, %xmm8
  2414: 66440f3800c6             pshufb %xmm6, %xmm8
  2420: 660fefcf                 pxor %xmm7, %xmm1
  2424: 66410fefc8               pxor %xmm8, %xmm1
  2429: 66430f6f641a20           movdqa 0x20(%r10, %r11), %xmm4
  2436: 660f6fec                 movdqa %xmm4, %xmm5
  2440: 66410fdbea               pand %xmm10, %xmm5
  2445: 660f6ff4                 movdqa %xmm4, %xmm6
  2449: 660f71d604               psrlw $4, %xmm6
  2454: 66410fdbf2               pand %xmm10, %xmm6
  2459: 66450f6f8fc0030000       movdqa 0x3c0(%r15), %xmm9
  2468: 66410f6ff9               movdqa %xmm9, %xmm7
  2473: 660f3800fd               pshufb %xmm5, %xmm7
  2478: 66450f6f8f00040000       movdqa 0x400(%r15), %xmm9
  2487: 66450f6fc1               movdqa %xmm9, %xmm8
  2492: 66440f3800c6             pshufb %xmm6, %xmm8
  2498: 660fefd7                 pxor %xmm7, %xmm2
  2502: 66410fefd0               pxor %xmm8, %xmm2
  2507: 66430f6f641a30           movdqa 0x30(%r10, %r11), %xmm4
  2514: 660f6fec                 movdqa %xmm4, %xmm5
  2518: 66410fdbea               pand %xmm10, %xmm5
  2523: 660f6ff4                 movdqa %xmm4, %xmm6
  2527: 660f71d604               psrlw $4, %xmm6
  2532: 66410fdbf2               pand %xmm10, %xmm6
  2537: 66450f6f8fc0030000       movdqa 0x3c0(%r15), %xmm9
  2546: 66410f6ff9               movdqa %xmm9, %xmm7
  2551: 660f3800fd               pshufb %xmm5, %xmm7
  2556: 66450f6f8f00040000       movdqa 0x400(%r15), %xmm9
  2565: 66450f6fc1               movdqa %xmm9, %xmm8
  2570: 66440f3800c6             pshufb %xmm6, %xmm8
  2576: 660fefdf                 pxor %xmm7, %xmm3
  2580: 66410fefd8               pxor %xmm8, %xmm3
  2585: 4d8b5040                 movq 0x40(%r8), %r10
  2589: 66430f6f241a             movdqa (%r10, %r11), %xmm4
  2595: 660f6fec                 movdqa %xmm4, %xmm5
  2599: 66410fdbea               pand %xmm10, %xmm5
  2604: 660f6ff4                 movdqa %xmm4, %xmm6
  2608: 660f71d604               psrlw $4, %xmm6
  2613: 66410fdbf2               pand %xmm10, %xmm6
  2618: 66450f6f8f40040000       movdqa 0x440(%r15), %xmm9
  2627: 66410f6ff9               movdqa %xmm9, %xmm7
  2632: 660f3800fd               pshufb %xmm5, %xmm7
  2637: 66450f6f8f80040000       movdqa 0x480(%r15), %xmm9
  2646: 66450f6fc1               movdqa %xmm9, %xmm8
  2651: 66440f3800c6             pshufb %xmm6, %xmm8
  2657: 660fefc7                 pxor %xmm7, %xmm0
  2661: 66410fefc0               pxor %xmm8, %xmm0
  2666: 66430f6f641a10           movdqa 0x10(%r10, %r11), %xmm4
  2673: 660f6fec                 movdqa %xmm4, %xmm5
  2677: 66410fdbea               pand %xmm10, %xmm5
  2682: 660f6ff4                 movdqa %xmm4, %xmm6
  2686: 660f71d604               psrlw $4, %xmm6
  2691: 66410fdbf2               pand %xmm10, %xmm6
  2696: 66450f6f8f40040000       movdqa 0x440(%r15), %xmm9
  2705: 66410f6ff9               movdqa %xmm9, %xmm7
  2710: 660f3800fd               pshufb %xmm5, %xmm7
  2715: 66450f6f8f80040000       movdqa 0x480(%r15), %xmm9
  2724: 66450f6fc1               movdqa %xmm9, %xmm8
  2729: 66440f3800c6             pshufb %xmm6, %xmm8
  2735: 660fefcf                 pxor %xmm7, %xmm1
  2739: 66410fefc8               pxor %xmm8, %xmm1
  2744: 66430f6f641a20           movdqa 0x20(%r10, %r11), %xmm4
  2751: 660f6fec                 movdqa %xmm4, %xmm5
  2755: 66410fdbea               pand %xmm10, %xmm5
  2760: 660f6ff4                 movdqa %xmm4, %xmm6
  2764: 660f71d604               psrlw $4, %xmm6
  2769: 66410fdbf2               pand %xmm10, %xmm6
  2774: 66450f6f8f40040000       movdqa 0x440(%r15), %xmm9
  2783: 66410f6ff9               movdqa %xmm9, %xmm7
  2788: 660f3800fd               pshufb %xmm5, %xmm7
  2793: 66450f6f8f80040000       movdqa 0x480(%r15), %xmm9
  2802: 66450f6fc1               movdqa %xmm9, %xmm8
  2807: 66440f3800c6             pshufb %xmm6, %xmm8
  2813: 660fefd7                 pxor %xmm7, %xmm2
  2817: 66410fefd0               pxor %xmm8, %xmm2
  2822: 66430f6f641a30           movdqa 0x30(%r10, %r11), %xmm4
  2829: 660f6fec                 movdqa %xmm4, %xmm5
  2833: 66410fdbea               pand %xmm10, %xmm5
  2838: 660f6ff4                 movdqa %xmm4, %xmm6
  2842: 660f71d604               psrlw $4, %xmm6
  2847: 66410fdbf2               pand %xmm10, %xmm6
  2852: 66450f6f8f40040000       movdqa 0x440(%r15), %xmm9
  2861: 66410f6ff9               movdqa %xmm9, %xmm7
  2866: 660f3800fd               pshufb %xmm5, %xmm7
  2871: 66450f6f8f80040000       movdqa 0x480(%r15), %xmm9
  2880: 66450f6fc1               movdqa %xmm9, %xmm8
  2885: 66440f3800c6             pshufb %xmm6, %xmm8
  2891: 660fefdf                 pxor %xmm7, %xmm3
  2895: 66410fefd8               pxor %xmm8, %xmm3
  2900: 4d8b5048                 movq 0x48(%r8), %r10
  2904: 66430f6f241a             movdqa (%r10, %r11), %xmm4
  2910: 660f6fec                 movdqa %xmm4, %xmm5
  2914: 66410fdbea               pand %xmm10, %xmm5
  2919: 660f6ff4                 movdqa %xmm4, %xmm6
  2923: 660f71d604               psrlw $4, %xmm6
  2928: 66410fdbf2               pand %xmm10, %xmm6
  2933: 66450f6f8fc0040000       movdqa 0x4c0(%r15), %xmm9
  2942: 66410f6ff9               movdqa %xmm9, %xmm7
  2947: 660f3800fd               pshufb %xmm5, %xmm7
  2952: 66450f6f8f00050000       movdqa 0x500(%r15), %xmm9
  2961: 66450f6fc1               movdqa %xmm9, %xmm8
  2966: 66440f3800c6             pshufb %xmm6, %xmm8
  2972: 660fefc7                 pxor %xmm7, %xmm0
  2976: 66410fefc0               pxor %xmm8, %xmm0
  2981: 66430f6f641a10           movdqa 0x10(%r10, %r11), %xmm4
  2988: 660f6fec                 movdqa %xmm4, %xmm5
  2992: 66410fdbea               pand %xmm10, %xmm5
  2997: 660f6ff4                 movdqa %xmm4, %xmm6
  3001: 660f71d604               psrlw $4, %xmm6
  3006: 66410fdbf2               pand %xmm10, %xmm6
  3011: 66450f6f8fc0040000       movdqa 0x4c0(%r15), %xmm9
  3020: 66410f6ff9               movdqa %xmm9, %xmm7
  3025: 660f3800fd               pshufb %xmm5, %xmm7
  3030: 66450f6f8f00050000       movdqa 0x500(%r15), %xmm9
  3039: 66450f6fc1               movdqa %xmm9, %xmm8
  3044: 66440f3800c6             pshufb %xmm6, %xmm8
  3050: 660fefcf                 pxor %xmm7, %xmm1
  3054: 66410fefc8               pxor %xmm8, %xmm1
  3059: 66430f6f641a20           movdqa 0x20(%r10, %r11), %xmm4
  3066: 660f6fec                 movdqa %xmm4, %xmm5
  3070: 66410fdbea               pand %xmm10, %xmm5
  3075: 660f6ff4                 movdqa %xmm4, %xmm6
  3079: 660f71d604               psrlw $4, %xmm6
  3084: 66410fdbf2               pand %xmm10, %xmm6
  3089: 66450f6f8fc0040000       movdqa 0x4c0(%r15), %xmm9
  3098: 66410f6ff9               movdqa %xmm9, %xmm7
  3103: 660f3800fd               pshufb %xmm5, %xmm7
  3108: 66450f6f8f00050000       movdqa 0x500(%r15), %xmm9
  3117: 66450f6fc1               movdqa %xmm9, %xmm8
  3122: 66440f3800c6             pshufb %xmm6, %xmm8
  3128: 660fefd7                 pxor %xmm7, %xmm2
  3132: 66410fefd0               pxor %xmm8, %xmm2
  3137: 66430f6f641a30           movdqa 0x30(%r10, %r11), %xmm4
  3144: 660f6fec                 movdqa %xmm4, %xmm5
  3148: 66410fdbea               pand %xmm10, %xmm5
  3153: 660f6ff4                 movdqa %xmm4, %xmm6
  3157: 660f71d604               psrlw $4, %xmm6
  3162: 66410fdbf2               pand %xmm10, %xmm6
  3167: 66450f6f8fc0040000       movdqa 0x4c0(%r15), %xmm9
  3176: 66410f6ff9               movdqa %xmm9, %xmm7
  3181: 660f3800fd               pshufb %xmm5, %xmm7
  3186: 66450f6f8f00050000       movdqa 0x500(%r15), %xmm9
  3195: 66450f6fc1               movdqa %xmm9, %xmm8
  3200: 66440f3800c6             pshufb %xmm6, %xmm8
  3206: 660fefdf                 pxor %xmm7, %xmm3
  3210: 66410fefd8               pxor %xmm8, %xmm3
  3215: 4d8b5050                 movq 0x50(%r8), %r10
  3219: 66430f6f241a             movdqa (%r10, %r11), %xmm4
  3225: 660f6fec                 movdqa %xmm4, %xmm5
  3229: 66410fdbea               pand %xmm10, %xmm5
  3234: 660f6ff4                 movdqa %xmm4, %xmm6
  3238: 660f71d604               psrlw $4, %xmm6
  3243: 66410fdbf2               pand %xmm10, %xmm6
  3248: 66450f6f8f40050000       movdqa 0x540(%r15), %xmm9
  3257: 66410f6ff9               movdqa %xmm9, %xmm7
  3262: 660f3800fd               pshufb %xmm5, %xmm7
  3267: 66450f6f8f80050000       movdqa 0x580(%r15), %xmm9
  3276: 66450f6fc1               movdqa %xmm9, %xmm8
  3281: 66440f3800c6             pshufb %xmm6, %xmm8
  3287: 660fefc7                 pxor %xmm7, %xmm0
  3291: 66410fefc0               pxor %xmm8, %xmm0
  3296: 66430f6f641a10           movdqa 0x10(%r10, %r11), %xmm4
  3303: 660f6fec                 movdqa %xmm4, %xmm5
  3307: 66410fdbea               pand %xmm10, %xmm5
  3312: 660f6ff4                 movdqa %xmm4, %xmm6
  3316: 660f71d604               psrlw $4, %xmm6
  3321: 66410fdbf2               pand %xmm10, %xmm6
  3326: 66450f6f8f40050000       movdqa 0x540(%r15), %xmm9
  3335: 66410f6ff9               movdqa %xmm9, %xmm7
  3340: 660f3800fd               pshufb %xmm5, %xmm7
  3345: 66450f6f8f80050000       movdqa 0x580(%r15), %xmm9
  3354: 66450f6fc1               movdqa %xmm9, %xmm8
  3359: 66440f3800c6             pshufb %xmm6, %xmm8
  3365: 660fefcf                 pxor %xmm7, %xmm1
  3369: 66410fefc8               pxor %xmm8, %xmm1
  3374: 66430f6f641a20           movdqa 0x20(%r10, %r11), %xmm4
  3381: 660f6fec                 movdqa %xmm4, %xmm5
  3385: 66410fdbea               pand %xmm10, %xmm5
  3390: 660f6ff4                 movdqa %xmm4, %xmm6
  3394: 660f71d604               psrlw $4, %xmm6
  3399: 66410fdbf2               pand %xmm10, %xmm6
  3404: 66450f6f8f40050000       movdqa 0x540(%r15), %xmm9
  3413: 66410f6ff9               movdqa %xmm9, %xmm7
  3418: 660f3800fd               pshufb %xmm5, %xmm7
  3423: 66450f6f8f80050000       movdqa 0x580(%r15), %xmm9
  3432: 66450f6fc1               movdqa %xmm9, %xmm8
  3437: 66440f3800c6             pshufb %xmm6, %xmm8
  3443: 660fefd7                 pxor %xmm7, %xmm2
  3447: 66410fefd0               pxor %xmm8, %xmm2
  3452: 66430f6f641a30           movdqa 0x30(%r10, %r11), %xmm4
  3459: 660f6fec                 movdqa %xmm4, %xmm5
  3463: 66410fdbea               pand %xmm10, %xmm5
  3468: 660f6ff4                 movdqa %xmm4, %xmm6
  3472: 660f71d604               psrlw $4, %xmm6
  3477: 66410fdbf2               pand %xmm10, %xmm6
  3482: 66450f6f8f40050000       movdqa 0x540(%r15), %xmm9
  3491: 66410f6ff9               movdqa %xmm9, %xmm7
  3496: 660f3800fd               pshufb %xmm5, %xmm7
  3501: 66450f6f8f80050000       movdqa 0x580(%r15), %xmm9
  3510: 66450f6fc1               movdqa %xmm9, %xmm8
  3515: 66440f3800c6             pshufb %xmm6, %xmm8
  3521: 660fefdf                 pxor %xmm7, %xmm3
  3525: 66410fefd8               pxor %xmm8, %xmm3
  3530: 4d8b5058                 movq 0x58(%r8), %r10
  3534: 66430f6f241a             movdqa (%r10, %r11), %xmm4
  3540: 660fefc4                 pxor %xmm4, %xmm0
  3544: 66430f6f641a10           movdqa 0x10(%r10, %r11), %xmm4
  3551: 660fefcc                 pxor %xmm4, %xmm1
  3555: 66430f6f641a20           movdqa 0x20(%r10, %r11), %xmm4
  3562: 660fefd4                 pxor %xmm4, %xmm2
  3566: 66430f6f641a30           movdqa 0x30(%r10, %r11), %xmm4
  3573: 660fefdc                 pxor %xmm4, %xmm3
  3577: 4d8b11                   movq (%r9), %r10
  3580: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
  3586: 66430f7f4c1a10           movdqa %xmm1, 0x10(%r10, %r11)
  3593: 66430f7f541a20           movdqa %xmm2, 0x20(%r10, %r11)
  3600: 66430f7f5c1a30           movdqa %xmm3, 0x30(%r10, %r11)
  3607: 4983c340                 addq $0x40, %r11
  3611: 4c3bd8                   cmpq %rax, %r11
  3614: 0f863cf2ffff             jbe 0x60
  3620: 4c3bd9                   cmpq %rcx, %r11
  3623: 0f8363020000             jae 0x1090
  3629: 4153                     pushq %r11
  3631: 498dbfc0050000           leaq 0x5c0(%r15), %rdi
  3638: 498db7c0060000           leaq 0x6c0(%r15), %rsi
  3645: 4d8b11                   movq (%r9), %r10
  3648: 4c8b1c24                 movq (%rsp), %r11
  3652: 33db                     xorl %ebx, %ebx
  3654: 498b00                   movq (%r8), %rax
  3657: 420fb61418               movzbl (%rax, %r11), %edx
  3662: 4885d2                   testq %rdx, %rdx
  3665: 0f841f000000             je 0xe76
  3671: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3675: 83c019                   addl $0x19, %eax
  3678: 81f8ff000000             cmpl $0xff, %eax
  3684: 0f8206000000             jb 0xe70
  3690: 81c001ffffff             addl $0xffffff01, %eax
  3696: 0fb60406                 movzbl (%rsi, %rax), %eax
  3700: 32d8                     xorb %al, %bl
  3702: 498b4008                 movq 8(%r8), %rax
  3706: 420fb61418               movzbl (%rax, %r11), %edx
  3711: 4885d2                   testq %rdx, %rdx
  3714: 0f841f000000             je 0xea7
  3720: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3724: 83c032                   addl $0x32, %eax
  3727: 81f8ff000000             cmpl $0xff, %eax
  3733: 0f8206000000             jb 0xea1
  3739: 81c001ffffff             addl $0xffffff01, %eax
  3745: 0fb60406                 movzbl (%rsi, %rax), %eax
  3749: 32d8                     xorb %al, %bl
  3751: 498b4010                 movq 0x10(%r8), %rax
  3755: 420fb61418               movzbl (%rax, %r11), %edx
  3760: 4885d2                   testq %rdx, %rdx
  3763: 0f841f000000             je 0xed8
  3769: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3773: 83c04b                   addl $0x4b, %eax
  3776: 81f8ff000000             cmpl $0xff, %eax
  3782: 0f8206000000             jb 0xed2
  3788: 81c001ffffff             addl $0xffffff01, %eax
  3794: 0fb60406                 movzbl (%rsi, %rax), %eax
  3798: 32d8                     xorb %al, %bl
  3800: 498b4018                 movq 0x18(%r8), %rax
  3804: 420fb61418               movzbl (%rax, %r11), %edx
  3809: 4885d2                   testq %rdx, %rdx
  3812: 0f841f000000             je 0xf09
  3818: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3822: 83c064                   addl $0x64, %eax
  3825: 81f8ff000000             cmpl $0xff, %eax
  3831: 0f8206000000             jb 0xf03
  3837: 81c001ffffff             addl $0xffffff01, %eax
  3843: 0fb60406                 movzbl (%rsi, %rax), %eax
  3847: 32d8                     xorb %al, %bl
  3849: 498b4020                 movq 0x20(%r8), %rax
  3853: 420fb61418               movzbl (%rax, %r11), %edx
  3858: 4885d2                   testq %rdx, %rdx
  3861: 0f841f000000             je 0xf3a
  3867: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3871: 83c07d                   addl $0x7d, %eax
  3874: 81f8ff000000             cmpl $0xff, %eax
  3880: 0f8206000000             jb 0xf34
  3886: 81c001ffffff             addl $0xffffff01, %eax
  3892: 0fb60406                 movzbl (%rsi, %rax), %eax
  3896: 32d8                     xorb %al, %bl
  3898: 498b4028                 movq 0x28(%r8), %rax
  3902: 420fb61418               movzbl (%rax, %r11), %edx
  3907: 4885d2                   testq %rdx, %rdx
  3910: 0f8422000000             je 0xf6e
  3916: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3920: 81c096000000             addl $0x96, %eax
  3926: 81f8ff000000             cmpl $0xff, %eax
  3932: 0f8206000000             jb 0xf68
  3938: 81c001ffffff             addl $0xffffff01, %eax
  3944: 0fb60406                 movzbl (%rsi, %rax), %eax
  3948: 32d8                     xorb %al, %bl
  3950: 498b4030                 movq 0x30(%r8), %rax
  3954: 420fb61418               movzbl (%rax, %r11), %edx
  3959: 4885d2                   testq %rdx, %rdx
  3962: 0f8422000000             je 0xfa2
  3968: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3972: 81c0af000000             addl $0xaf, %eax
  3978: 81f8ff000000             cmpl $0xff, %eax
  3984: 0f8206000000             jb 0xf9c
  3990: 81c001ffffff             addl $0xffffff01, %eax
  3996: 0fb60406                 movzbl (%rsi, %rax), %eax
  4000: 32d8                     xorb %al, %bl
  4002: 498b4038                 movq 0x38(%r8), %rax
  4006: 420fb61418               movzbl (%rax, %r11), %edx
  4011: 4885d2                   testq %rdx, %rdx
  4014: 0f8422000000             je 0xfd6
  4020: 0fb60417                 movzbl (%rdi, %rdx), %eax
  4024: 81c0c8000000             addl $0xc8, %eax
  4030: 81f8ff000000             cmpl $0xff, %eax
  4036: 0f8206000000             jb 0xfd0
  4042: 81c001ffffff             addl $0xffffff01, %eax
  4048: 0fb60406                 movzbl (%rsi, %rax), %eax
  4052: 32d8                     xorb %al, %bl
  4054: 498b4040                 movq 0x40(%r8), %rax
  4058: 420fb61418               movzbl (%rax, %r11), %edx
  4063: 4885d2                   testq %rdx, %rdx
  4066: 0f8422000000             je 0x100a
  4072: 0fb60417                 movzbl (%rdi, %rdx), %eax
  4076: 81c0e1000000             addl $0xe1, %eax
  4082: 81f8ff000000             cmpl $0xff, %eax
  4088: 0f8206000000             jb 0x1004
  4094: 81c001ffffff             addl $0xffffff01, %eax
  4100: 0fb60406                 movzbl (%rsi, %rax), %eax
  4104: 32d8                     xorb %al, %bl
  4106: 498b4048                 movq 0x48(%r8), %rax
  4110: 420fb61418               movzbl (%rax, %r11), %edx
  4115: 4885d2                   testq %rdx, %rdx
  4118: 0f8422000000             je 0x103e
  4124: 0fb60417                 movzbl (%rdi, %rdx), %eax
  4128: 81c0fa000000             addl $0xfa, %eax
  4134: 81f8ff000000             cmpl $0xff, %eax
  4140: 0f8206000000             jb 0x1038
  4146: 81c001ffffff             addl $0xffffff01, %eax
  4152: 0fb60406                 movzbl (%rsi, %rax), %eax
  4156: 32d8                     xorb %al, %bl
  4158: 498b4050                 movq 0x50(%r8), %rax
  4162: 420fb61418               movzbl (%rax, %r11), %edx
  4167: 4885d2                   testq %rdx, %rdx
  4170: 0f841f000000             je 0x106f
  4176: 0fb60417                 movzbl (%rdi, %rdx), %eax
  4180: 83c014                   addl $0x14, %eax
  4183: 81f8ff000000             cmpl $0xff, %eax
  4189: 0f8206000000             jb 0x1069
  4195: 81c001ffffff             addl $0xffffff01, %eax
  4201: 0fb60406                 movzbl (%rsi, %rax), %eax
  4205: 32d8                     xorb %al, %bl
  4207: 498b4058                 movq 0x58(%r8), %rax
  4211: 420fb61418               movzbl (%rax, %r11), %edx
  4216: 32da                     xorb %dl, %bl
  4218: 43881c1a                 movb %bl, (%r10, %r11)
  4222: 33db                     xorl %ebx, %ebx
  4224: 49ffc3                   incq %r11
  4227: 4c3bd9                   cmpq %rcx, %r11
  4230: 0f82bafdffff             jb 0xe46
  4236: 4883c408                 addq $8, %rsp
  4240: 415e                     popq %r14
  4242: 415d                     popq %r13
  4244: 415c                     popq %r12
  4246: 415f                     popq %r15
  4248: 5b                       popq %rbx
  4249: c3                       retq 
  4250: 0f0b                     ud2 
  4252: 0f0b                     ud2 
  4254: 0f0b                     ud2 
  4256: 0f0b                     ud2 
  4258: 0f0b                     ud2 
  4260: 0f0b                     ud2 
  4262: 0f0b                     ud2 
  4264: 0f0b                     ud2 
  4266: 0f0b                     ud2 
  4268: 0f0b                     ud2 
  4270: 0f0b                     ud2 
  4272: 0f0b                     ud2 
  4274: 0f0b                     ud2 
  4276: 0f0b                     ud2 
  4278: 0f0b                     ud2 
  4280: 0f0b                     ud2 
  4282: 0f0b                     ud2 
  4284: 0f0b                     ud2 
  4286: 0f0b                     ud2 
