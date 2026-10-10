; config_t3_2110 target=sse twin=baseline topo=lrc(8,2,2)/rec{0} op=recover U=4
; code_size=2816 twins=[] uops_per_byte=6.359375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de70a0000           leaq 0xae7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66450f6f17               movdqa (%r15), %xmm10
    52: 488bc1                   movq %rcx, %rax
    55: 4883e840                 subq $0x40, %rax
    59: 0f82f7080000             jb 0x938
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
  2280: 660fefc4                 pxor %xmm4, %xmm0
  2284: 66430f6f641a10           movdqa 0x10(%r10, %r11), %xmm4
  2291: 660fefcc                 pxor %xmm4, %xmm1
  2295: 66430f6f641a20           movdqa 0x20(%r10, %r11), %xmm4
  2302: 660fefd4                 pxor %xmm4, %xmm2
  2306: 66430f6f641a30           movdqa 0x30(%r10, %r11), %xmm4
  2313: 660fefdc                 pxor %xmm4, %xmm3
  2317: 4d8b11                   movq (%r9), %r10
  2320: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
  2326: 66430f7f4c1a10           movdqa %xmm1, 0x10(%r10, %r11)
  2333: 66430f7f541a20           movdqa %xmm2, 0x20(%r10, %r11)
  2340: 66430f7f5c1a30           movdqa %xmm3, 0x30(%r10, %r11)
  2347: 4983c340                 addq $0x40, %r11
  2351: 4c3bd8                   cmpq %rax, %r11
  2354: 0f8628f7ffff             jbe 0x60
  2360: 4c3bd9                   cmpq %rcx, %r11
  2363: 0f8396010000             jae 0xad7
  2369: 4153                     pushq %r11
  2371: 498dbfc0030000           leaq 0x3c0(%r15), %rdi
  2378: 498db7c0040000           leaq 0x4c0(%r15), %rsi
  2385: 4d8b11                   movq (%r9), %r10
  2388: 4c8b1c24                 movq (%rsp), %r11
  2392: 33db                     xorl %ebx, %ebx
  2394: 498b00                   movq (%r8), %rax
  2397: 420fb61418               movzbl (%rax, %r11), %edx
  2402: 4885d2                   testq %rdx, %rdx
  2405: 0f841f000000             je 0x98a
  2411: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2415: 83c019                   addl $0x19, %eax
  2418: 81f8ff000000             cmpl $0xff, %eax
  2424: 0f8206000000             jb 0x984
  2430: 81c001ffffff             addl $0xffffff01, %eax
  2436: 0fb60406                 movzbl (%rsi, %rax), %eax
  2440: 32d8                     xorb %al, %bl
  2442: 498b4008                 movq 8(%r8), %rax
  2446: 420fb61418               movzbl (%rax, %r11), %edx
  2451: 4885d2                   testq %rdx, %rdx
  2454: 0f841f000000             je 0x9bb
  2460: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2464: 83c032                   addl $0x32, %eax
  2467: 81f8ff000000             cmpl $0xff, %eax
  2473: 0f8206000000             jb 0x9b5
  2479: 81c001ffffff             addl $0xffffff01, %eax
  2485: 0fb60406                 movzbl (%rsi, %rax), %eax
  2489: 32d8                     xorb %al, %bl
  2491: 498b4010                 movq 0x10(%r8), %rax
  2495: 420fb61418               movzbl (%rax, %r11), %edx
  2500: 4885d2                   testq %rdx, %rdx
  2503: 0f841f000000             je 0x9ec
  2509: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2513: 83c04b                   addl $0x4b, %eax
  2516: 81f8ff000000             cmpl $0xff, %eax
  2522: 0f8206000000             jb 0x9e6
  2528: 81c001ffffff             addl $0xffffff01, %eax
  2534: 0fb60406                 movzbl (%rsi, %rax), %eax
  2538: 32d8                     xorb %al, %bl
  2540: 498b4018                 movq 0x18(%r8), %rax
  2544: 420fb61418               movzbl (%rax, %r11), %edx
  2549: 4885d2                   testq %rdx, %rdx
  2552: 0f841f000000             je 0xa1d
  2558: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2562: 83c064                   addl $0x64, %eax
  2565: 81f8ff000000             cmpl $0xff, %eax
  2571: 0f8206000000             jb 0xa17
  2577: 81c001ffffff             addl $0xffffff01, %eax
  2583: 0fb60406                 movzbl (%rsi, %rax), %eax
  2587: 32d8                     xorb %al, %bl
  2589: 498b4020                 movq 0x20(%r8), %rax
  2593: 420fb61418               movzbl (%rax, %r11), %edx
  2598: 4885d2                   testq %rdx, %rdx
  2601: 0f841f000000             je 0xa4e
  2607: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2611: 83c07d                   addl $0x7d, %eax
  2614: 81f8ff000000             cmpl $0xff, %eax
  2620: 0f8206000000             jb 0xa48
  2626: 81c001ffffff             addl $0xffffff01, %eax
  2632: 0fb60406                 movzbl (%rsi, %rax), %eax
  2636: 32d8                     xorb %al, %bl
  2638: 498b4028                 movq 0x28(%r8), %rax
  2642: 420fb61418               movzbl (%rax, %r11), %edx
  2647: 4885d2                   testq %rdx, %rdx
  2650: 0f8422000000             je 0xa82
  2656: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2660: 81c096000000             addl $0x96, %eax
  2666: 81f8ff000000             cmpl $0xff, %eax
  2672: 0f8206000000             jb 0xa7c
  2678: 81c001ffffff             addl $0xffffff01, %eax
  2684: 0fb60406                 movzbl (%rsi, %rax), %eax
  2688: 32d8                     xorb %al, %bl
  2690: 498b4030                 movq 0x30(%r8), %rax
  2694: 420fb61418               movzbl (%rax, %r11), %edx
  2699: 4885d2                   testq %rdx, %rdx
  2702: 0f8422000000             je 0xab6
  2708: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2712: 81c0af000000             addl $0xaf, %eax
  2718: 81f8ff000000             cmpl $0xff, %eax
  2724: 0f8206000000             jb 0xab0
  2730: 81c001ffffff             addl $0xffffff01, %eax
  2736: 0fb60406                 movzbl (%rsi, %rax), %eax
  2740: 32d8                     xorb %al, %bl
  2742: 498b4038                 movq 0x38(%r8), %rax
  2746: 420fb61418               movzbl (%rax, %r11), %edx
  2751: 32da                     xorb %dl, %bl
  2753: 43881c1a                 movb %bl, (%r10, %r11)
  2757: 33db                     xorl %ebx, %ebx
  2759: 49ffc3                   incq %r11
  2762: 4c3bd9                   cmpq %rcx, %r11
  2765: 0f8287feffff             jb 0x95a
  2771: 4883c408                 addq $8, %rsp
  2775: 415e                     popq %r14
  2777: 415d                     popq %r13
  2779: 415c                     popq %r12
  2781: 415f                     popq %r15
  2783: 5b                       popq %rbx
  2784: c3                       retq 
  2785: 0f0b                     ud2 
  2787: 0f0b                     ud2 
  2789: 0f0b                     ud2 
  2791: 0f0b                     ud2 
  2793: 0f0b                     ud2 
  2795: 0f0b                     ud2 
  2797: 0f0b                     ud2 
  2799: 0f0b                     ud2 
  2801: 0f0b                     ud2 
  2803: 0f0b                     ud2 
  2805: 0f0b                     ud2 
  2807: 0f0b                     ud2 
  2809: 0f0b                     ud2 
  2811: 0f0b                     ud2 
  2813: 0f0b                     ud2 
  2815: cc                       int3 
