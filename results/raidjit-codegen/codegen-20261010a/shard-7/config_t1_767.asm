; config_t1_767 target=sse twin=baseline topo=rs(6,3)/rec{0,1} op=recover U=8
; code_size=3072 twins=[] uops_per_byte=5.890625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de70b0000           leaq 0xbe7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66450f6f37               movdqa (%r15), %xmm14
    52: 488bc1                   movq %rcx, %rax
    55: 4883e840                 subq $0x40, %rax
    59: 0f82f0080000             jb 0x931
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66460f6f041e             movdqa (%rsi, %r11), %xmm8
   102: 66450f6fc8               movdqa %xmm8, %xmm9
   107: 66450fdbce               pand %xmm14, %xmm9
   112: 66450f6fd0               movdqa %xmm8, %xmm10
   117: 66410f71d204             psrlw $4, %xmm10
   123: 66450fdbd6               pand %xmm14, %xmm10
   128: 66410f6fc0               movdqa %xmm8, %xmm0
   133: 66450f6f5f40             movdqa 0x40(%r15), %xmm11
   139: 66450f3800d9             pshufb %xmm9, %xmm11
   145: 66450f6fa780000000       movdqa 0x80(%r15), %xmm12
   154: 66450f3800e2             pshufb %xmm10, %xmm12
   160: 66410f6fe3               movdqa %xmm11, %xmm4
   165: 66410fefe4               pxor %xmm12, %xmm4
   170: 66460f6f441e10           movdqa 0x10(%rsi, %r11), %xmm8
   177: 66450f6fc8               movdqa %xmm8, %xmm9
   182: 66450fdbce               pand %xmm14, %xmm9
   187: 66450f6fd0               movdqa %xmm8, %xmm10
   192: 66410f71d204             psrlw $4, %xmm10
   198: 66450fdbd6               pand %xmm14, %xmm10
   203: 66410f6fc8               movdqa %xmm8, %xmm1
   208: 66450f6f5f40             movdqa 0x40(%r15), %xmm11
   214: 66450f3800d9             pshufb %xmm9, %xmm11
   220: 66450f6fa780000000       movdqa 0x80(%r15), %xmm12
   229: 66450f3800e2             pshufb %xmm10, %xmm12
   235: 66410f6feb               movdqa %xmm11, %xmm5
   240: 66410fefec               pxor %xmm12, %xmm5
   245: 66460f6f441e20           movdqa 0x20(%rsi, %r11), %xmm8
   252: 66450f6fc8               movdqa %xmm8, %xmm9
   257: 66450fdbce               pand %xmm14, %xmm9
   262: 66450f6fd0               movdqa %xmm8, %xmm10
   267: 66410f71d204             psrlw $4, %xmm10
   273: 66450fdbd6               pand %xmm14, %xmm10
   278: 66410f6fd0               movdqa %xmm8, %xmm2
   283: 66450f6f5f40             movdqa 0x40(%r15), %xmm11
   289: 66450f3800d9             pshufb %xmm9, %xmm11
   295: 66450f6fa780000000       movdqa 0x80(%r15), %xmm12
   304: 66450f3800e2             pshufb %xmm10, %xmm12
   310: 66410f6ff3               movdqa %xmm11, %xmm6
   315: 66410feff4               pxor %xmm12, %xmm6
   320: 66460f6f441e30           movdqa 0x30(%rsi, %r11), %xmm8
   327: 66450f6fc8               movdqa %xmm8, %xmm9
   332: 66450fdbce               pand %xmm14, %xmm9
   337: 66450f6fd0               movdqa %xmm8, %xmm10
   342: 66410f71d204             psrlw $4, %xmm10
   348: 66450fdbd6               pand %xmm14, %xmm10
   353: 66410f6fd8               movdqa %xmm8, %xmm3
   358: 66450f6f5f40             movdqa 0x40(%r15), %xmm11
   364: 66450f3800d9             pshufb %xmm9, %xmm11
   370: 66450f6fa780000000       movdqa 0x80(%r15), %xmm12
   379: 66450f3800e2             pshufb %xmm10, %xmm12
   385: 66410f6ffb               movdqa %xmm11, %xmm7
   390: 66410feffc               pxor %xmm12, %xmm7
   395: 66460f6f041f             movdqa (%rdi, %r11), %xmm8
   401: 66450f6fc8               movdqa %xmm8, %xmm9
   406: 66450fdbce               pand %xmm14, %xmm9
   411: 66450f6fd0               movdqa %xmm8, %xmm10
   416: 66410f71d204             psrlw $4, %xmm10
   422: 66450fdbd6               pand %xmm14, %xmm10
   427: 66410fefc0               pxor %xmm8, %xmm0
   432: 66450f6f9fc0000000       movdqa 0xc0(%r15), %xmm11
   441: 66450f3800d9             pshufb %xmm9, %xmm11
   447: 66450f6fa700010000       movdqa 0x100(%r15), %xmm12
   456: 66450f3800e2             pshufb %xmm10, %xmm12
   462: 66410fefe3               pxor %xmm11, %xmm4
   467: 66410fefe4               pxor %xmm12, %xmm4
   472: 66460f6f441f10           movdqa 0x10(%rdi, %r11), %xmm8
   479: 66450f6fc8               movdqa %xmm8, %xmm9
   484: 66450fdbce               pand %xmm14, %xmm9
   489: 66450f6fd0               movdqa %xmm8, %xmm10
   494: 66410f71d204             psrlw $4, %xmm10
   500: 66450fdbd6               pand %xmm14, %xmm10
   505: 66410fefc8               pxor %xmm8, %xmm1
   510: 66450f6f9fc0000000       movdqa 0xc0(%r15), %xmm11
   519: 66450f3800d9             pshufb %xmm9, %xmm11
   525: 66450f6fa700010000       movdqa 0x100(%r15), %xmm12
   534: 66450f3800e2             pshufb %xmm10, %xmm12
   540: 66410fefeb               pxor %xmm11, %xmm5
   545: 66410fefec               pxor %xmm12, %xmm5
   550: 66460f6f441f20           movdqa 0x20(%rdi, %r11), %xmm8
   557: 66450f6fc8               movdqa %xmm8, %xmm9
   562: 66450fdbce               pand %xmm14, %xmm9
   567: 66450f6fd0               movdqa %xmm8, %xmm10
   572: 66410f71d204             psrlw $4, %xmm10
   578: 66450fdbd6               pand %xmm14, %xmm10
   583: 66410fefd0               pxor %xmm8, %xmm2
   588: 66450f6f9fc0000000       movdqa 0xc0(%r15), %xmm11
   597: 66450f3800d9             pshufb %xmm9, %xmm11
   603: 66450f6fa700010000       movdqa 0x100(%r15), %xmm12
   612: 66450f3800e2             pshufb %xmm10, %xmm12
   618: 66410feff3               pxor %xmm11, %xmm6
   623: 66410feff4               pxor %xmm12, %xmm6
   628: 66460f6f441f30           movdqa 0x30(%rdi, %r11), %xmm8
   635: 66450f6fc8               movdqa %xmm8, %xmm9
   640: 66450fdbce               pand %xmm14, %xmm9
   645: 66450f6fd0               movdqa %xmm8, %xmm10
   650: 66410f71d204             psrlw $4, %xmm10
   656: 66450fdbd6               pand %xmm14, %xmm10
   661: 66410fefd8               pxor %xmm8, %xmm3
   666: 66450f6f9fc0000000       movdqa 0xc0(%r15), %xmm11
   675: 66450f3800d9             pshufb %xmm9, %xmm11
   681: 66450f6fa700010000       movdqa 0x100(%r15), %xmm12
   690: 66450f3800e2             pshufb %xmm10, %xmm12
   696: 66410feffb               pxor %xmm11, %xmm7
   701: 66410feffc               pxor %xmm12, %xmm7
   706: 66470f6f041c             movdqa (%r12, %r11), %xmm8
   712: 66450f6fc8               movdqa %xmm8, %xmm9
   717: 66450fdbce               pand %xmm14, %xmm9
   722: 66450f6fd0               movdqa %xmm8, %xmm10
   727: 66410f71d204             psrlw $4, %xmm10
   733: 66450fdbd6               pand %xmm14, %xmm10
   738: 66410fefc0               pxor %xmm8, %xmm0
   743: 66450f6f9f40010000       movdqa 0x140(%r15), %xmm11
   752: 66450f3800d9             pshufb %xmm9, %xmm11
   758: 66450f6fa780010000       movdqa 0x180(%r15), %xmm12
   767: 66450f3800e2             pshufb %xmm10, %xmm12
   773: 66410fefe3               pxor %xmm11, %xmm4
   778: 66410fefe4               pxor %xmm12, %xmm4
   783: 66470f6f441c10           movdqa 0x10(%r12, %r11), %xmm8
   790: 66450f6fc8               movdqa %xmm8, %xmm9
   795: 66450fdbce               pand %xmm14, %xmm9
   800: 66450f6fd0               movdqa %xmm8, %xmm10
   805: 66410f71d204             psrlw $4, %xmm10
   811: 66450fdbd6               pand %xmm14, %xmm10
   816: 66410fefc8               pxor %xmm8, %xmm1
   821: 66450f6f9f40010000       movdqa 0x140(%r15), %xmm11
   830: 66450f3800d9             pshufb %xmm9, %xmm11
   836: 66450f6fa780010000       movdqa 0x180(%r15), %xmm12
   845: 66450f3800e2             pshufb %xmm10, %xmm12
   851: 66410fefeb               pxor %xmm11, %xmm5
   856: 66410fefec               pxor %xmm12, %xmm5
   861: 66470f6f441c20           movdqa 0x20(%r12, %r11), %xmm8
   868: 66450f6fc8               movdqa %xmm8, %xmm9
   873: 66450fdbce               pand %xmm14, %xmm9
   878: 66450f6fd0               movdqa %xmm8, %xmm10
   883: 66410f71d204             psrlw $4, %xmm10
   889: 66450fdbd6               pand %xmm14, %xmm10
   894: 66410fefd0               pxor %xmm8, %xmm2
   899: 66450f6f9f40010000       movdqa 0x140(%r15), %xmm11
   908: 66450f3800d9             pshufb %xmm9, %xmm11
   914: 66450f6fa780010000       movdqa 0x180(%r15), %xmm12
   923: 66450f3800e2             pshufb %xmm10, %xmm12
   929: 66410feff3               pxor %xmm11, %xmm6
   934: 66410feff4               pxor %xmm12, %xmm6
   939: 66470f6f441c30           movdqa 0x30(%r12, %r11), %xmm8
   946: 66450f6fc8               movdqa %xmm8, %xmm9
   951: 66450fdbce               pand %xmm14, %xmm9
   956: 66450f6fd0               movdqa %xmm8, %xmm10
   961: 66410f71d204             psrlw $4, %xmm10
   967: 66450fdbd6               pand %xmm14, %xmm10
   972: 66410fefd8               pxor %xmm8, %xmm3
   977: 66450f6f9f40010000       movdqa 0x140(%r15), %xmm11
   986: 66450f3800d9             pshufb %xmm9, %xmm11
   992: 66450f6fa780010000       movdqa 0x180(%r15), %xmm12
  1001: 66450f3800e2             pshufb %xmm10, %xmm12
  1007: 66410feffb               pxor %xmm11, %xmm7
  1012: 66410feffc               pxor %xmm12, %xmm7
  1017: 66470f6f441d00           movdqa (%r13, %r11), %xmm8
  1024: 66410fefc0               pxor %xmm8, %xmm0
  1029: 66410fefe0               pxor %xmm8, %xmm4
  1034: 66470f6f441d10           movdqa 0x10(%r13, %r11), %xmm8
  1041: 66410fefc8               pxor %xmm8, %xmm1
  1046: 66410fefe8               pxor %xmm8, %xmm5
  1051: 66470f6f441d20           movdqa 0x20(%r13, %r11), %xmm8
  1058: 66410fefd0               pxor %xmm8, %xmm2
  1063: 66410feff0               pxor %xmm8, %xmm6
  1068: 66470f6f441d30           movdqa 0x30(%r13, %r11), %xmm8
  1075: 66410fefd8               pxor %xmm8, %xmm3
  1080: 66410feff8               pxor %xmm8, %xmm7
  1085: 66470f6f041e             movdqa (%r14, %r11), %xmm8
  1091: 66410fefc0               pxor %xmm8, %xmm0
  1096: 66470f6f441e10           movdqa 0x10(%r14, %r11), %xmm8
  1103: 66410fefc8               pxor %xmm8, %xmm1
  1108: 66470f6f441e20           movdqa 0x20(%r14, %r11), %xmm8
  1115: 66410fefd0               pxor %xmm8, %xmm2
  1120: 66470f6f441e30           movdqa 0x30(%r14, %r11), %xmm8
  1127: 66410fefd8               pxor %xmm8, %xmm3
  1132: 4d8b5028                 movq 0x28(%r8), %r10
  1136: 66470f6f041a             movdqa (%r10, %r11), %xmm8
  1142: 66410fefe0               pxor %xmm8, %xmm4
  1147: 66470f6f441a10           movdqa 0x10(%r10, %r11), %xmm8
  1154: 66410fefe8               pxor %xmm8, %xmm5
  1159: 66470f6f441a20           movdqa 0x20(%r10, %r11), %xmm8
  1166: 66410feff0               pxor %xmm8, %xmm6
  1171: 66470f6f441a30           movdqa 0x30(%r10, %r11), %xmm8
  1178: 66410feff8               pxor %xmm8, %xmm7
  1183: 66440f6fc8               movdqa %xmm0, %xmm9
  1188: 66450fdbce               pand %xmm14, %xmm9
  1193: 66440f6fd0               movdqa %xmm0, %xmm10
  1198: 66410f71d204             psrlw $4, %xmm10
  1204: 66450fdbd6               pand %xmm14, %xmm10
  1209: 66450f6f9fc0010000       movdqa 0x1c0(%r15), %xmm11
  1218: 66450f3800d9             pshufb %xmm9, %xmm11
  1224: 66450f6fa700020000       movdqa 0x200(%r15), %xmm12
  1233: 66450f3800e2             pshufb %xmm10, %xmm12
  1239: 66450f6feb               movdqa %xmm11, %xmm13
  1244: 66450fefec               pxor %xmm12, %xmm13
  1249: 66440fefec               pxor %xmm4, %xmm13
  1254: 66440f6fc8               movdqa %xmm0, %xmm9
  1259: 66450fdbce               pand %xmm14, %xmm9
  1264: 66440f6fd0               movdqa %xmm0, %xmm10
  1269: 66410f71d204             psrlw $4, %xmm10
  1275: 66450fdbd6               pand %xmm14, %xmm10
  1280: 66450f6f9f40020000       movdqa 0x240(%r15), %xmm11
  1289: 66450f3800d9             pshufb %xmm9, %xmm11
  1295: 66450f6fa780020000       movdqa 0x280(%r15), %xmm12
  1304: 66450f3800e2             pshufb %xmm10, %xmm12
  1310: 66450f6fc3               movdqa %xmm11, %xmm8
  1315: 66450fefc4               pxor %xmm12, %xmm8
  1320: 66440fefc4               pxor %xmm4, %xmm8
  1325: 66450f6fcd               movdqa %xmm13, %xmm9
  1330: 66450fdbce               pand %xmm14, %xmm9
  1335: 66450f6fd5               movdqa %xmm13, %xmm10
  1340: 66410f71d204             psrlw $4, %xmm10
  1346: 66450fdbd6               pand %xmm14, %xmm10
  1351: 66450f6f9fc0020000       movdqa 0x2c0(%r15), %xmm11
  1360: 66450f3800d9             pshufb %xmm9, %xmm11
  1366: 66450f6fa700030000       movdqa 0x300(%r15), %xmm12
  1375: 66450f3800e2             pshufb %xmm10, %xmm12
  1381: 66410f6fe3               movdqa %xmm11, %xmm4
  1386: 66410fefe4               pxor %xmm12, %xmm4
  1391: 66450f6fc8               movdqa %xmm8, %xmm9
  1396: 66450fdbce               pand %xmm14, %xmm9
  1401: 66450f6fd0               movdqa %xmm8, %xmm10
  1406: 66410f71d204             psrlw $4, %xmm10
  1412: 66450fdbd6               pand %xmm14, %xmm10
  1417: 66450f6f9fc0020000       movdqa 0x2c0(%r15), %xmm11
  1426: 66450f3800d9             pshufb %xmm9, %xmm11
  1432: 66450f6fa700030000       movdqa 0x300(%r15), %xmm12
  1441: 66450f3800e2             pshufb %xmm10, %xmm12
  1447: 66410f6fc3               movdqa %xmm11, %xmm0
  1452: 66410fefc4               pxor %xmm12, %xmm0
  1457: 66440f6fc9               movdqa %xmm1, %xmm9
  1462: 66450fdbce               pand %xmm14, %xmm9
  1467: 66440f6fd1               movdqa %xmm1, %xmm10
  1472: 66410f71d204             psrlw $4, %xmm10
  1478: 66450fdbd6               pand %xmm14, %xmm10
  1483: 66450f6f9fc0010000       movdqa 0x1c0(%r15), %xmm11
  1492: 66450f3800d9             pshufb %xmm9, %xmm11
  1498: 66450f6fa700020000       movdqa 0x200(%r15), %xmm12
  1507: 66450f3800e2             pshufb %xmm10, %xmm12
  1513: 66450f6feb               movdqa %xmm11, %xmm13
  1518: 66450fefec               pxor %xmm12, %xmm13
  1523: 66440fefed               pxor %xmm5, %xmm13
  1528: 66440f6fc9               movdqa %xmm1, %xmm9
  1533: 66450fdbce               pand %xmm14, %xmm9
  1538: 66440f6fd1               movdqa %xmm1, %xmm10
  1543: 66410f71d204             psrlw $4, %xmm10
  1549: 66450fdbd6               pand %xmm14, %xmm10
  1554: 66450f6f9f40020000       movdqa 0x240(%r15), %xmm11
  1563: 66450f3800d9             pshufb %xmm9, %xmm11
  1569: 66450f6fa780020000       movdqa 0x280(%r15), %xmm12
  1578: 66450f3800e2             pshufb %xmm10, %xmm12
  1584: 66450f6fc3               movdqa %xmm11, %xmm8
  1589: 66450fefc4               pxor %xmm12, %xmm8
  1594: 66440fefc5               pxor %xmm5, %xmm8
  1599: 66450f6fcd               movdqa %xmm13, %xmm9
  1604: 66450fdbce               pand %xmm14, %xmm9
  1609: 66450f6fd5               movdqa %xmm13, %xmm10
  1614: 66410f71d204             psrlw $4, %xmm10
  1620: 66450fdbd6               pand %xmm14, %xmm10
  1625: 66450f6f9fc0020000       movdqa 0x2c0(%r15), %xmm11
  1634: 66450f3800d9             pshufb %xmm9, %xmm11
  1640: 66450f6fa700030000       movdqa 0x300(%r15), %xmm12
  1649: 66450f3800e2             pshufb %xmm10, %xmm12
  1655: 66410f6feb               movdqa %xmm11, %xmm5
  1660: 66410fefec               pxor %xmm12, %xmm5
  1665: 66450f6fc8               movdqa %xmm8, %xmm9
  1670: 66450fdbce               pand %xmm14, %xmm9
  1675: 66450f6fd0               movdqa %xmm8, %xmm10
  1680: 66410f71d204             psrlw $4, %xmm10
  1686: 66450fdbd6               pand %xmm14, %xmm10
  1691: 66450f6f9fc0020000       movdqa 0x2c0(%r15), %xmm11
  1700: 66450f3800d9             pshufb %xmm9, %xmm11
  1706: 66450f6fa700030000       movdqa 0x300(%r15), %xmm12
  1715: 66450f3800e2             pshufb %xmm10, %xmm12
  1721: 66410f6fcb               movdqa %xmm11, %xmm1
  1726: 66410fefcc               pxor %xmm12, %xmm1
  1731: 66440f6fca               movdqa %xmm2, %xmm9
  1736: 66450fdbce               pand %xmm14, %xmm9
  1741: 66440f6fd2               movdqa %xmm2, %xmm10
  1746: 66410f71d204             psrlw $4, %xmm10
  1752: 66450fdbd6               pand %xmm14, %xmm10
  1757: 66450f6f9fc0010000       movdqa 0x1c0(%r15), %xmm11
  1766: 66450f3800d9             pshufb %xmm9, %xmm11
  1772: 66450f6fa700020000       movdqa 0x200(%r15), %xmm12
  1781: 66450f3800e2             pshufb %xmm10, %xmm12
  1787: 66450f6feb               movdqa %xmm11, %xmm13
  1792: 66450fefec               pxor %xmm12, %xmm13
  1797: 66440fefee               pxor %xmm6, %xmm13
  1802: 66440f6fca               movdqa %xmm2, %xmm9
  1807: 66450fdbce               pand %xmm14, %xmm9
  1812: 66440f6fd2               movdqa %xmm2, %xmm10
  1817: 66410f71d204             psrlw $4, %xmm10
  1823: 66450fdbd6               pand %xmm14, %xmm10
  1828: 66450f6f9f40020000       movdqa 0x240(%r15), %xmm11
  1837: 66450f3800d9             pshufb %xmm9, %xmm11
  1843: 66450f6fa780020000       movdqa 0x280(%r15), %xmm12
  1852: 66450f3800e2             pshufb %xmm10, %xmm12
  1858: 66450f6fc3               movdqa %xmm11, %xmm8
  1863: 66450fefc4               pxor %xmm12, %xmm8
  1868: 66440fefc6               pxor %xmm6, %xmm8
  1873: 66450f6fcd               movdqa %xmm13, %xmm9
  1878: 66450fdbce               pand %xmm14, %xmm9
  1883: 66450f6fd5               movdqa %xmm13, %xmm10
  1888: 66410f71d204             psrlw $4, %xmm10
  1894: 66450fdbd6               pand %xmm14, %xmm10
  1899: 66450f6f9fc0020000       movdqa 0x2c0(%r15), %xmm11
  1908: 66450f3800d9             pshufb %xmm9, %xmm11
  1914: 66450f6fa700030000       movdqa 0x300(%r15), %xmm12
  1923: 66450f3800e2             pshufb %xmm10, %xmm12
  1929: 66410f6ff3               movdqa %xmm11, %xmm6
  1934: 66410feff4               pxor %xmm12, %xmm6
  1939: 66450f6fc8               movdqa %xmm8, %xmm9
  1944: 66450fdbce               pand %xmm14, %xmm9
  1949: 66450f6fd0               movdqa %xmm8, %xmm10
  1954: 66410f71d204             psrlw $4, %xmm10
  1960: 66450fdbd6               pand %xmm14, %xmm10
  1965: 66450f6f9fc0020000       movdqa 0x2c0(%r15), %xmm11
  1974: 66450f3800d9             pshufb %xmm9, %xmm11
  1980: 66450f6fa700030000       movdqa 0x300(%r15), %xmm12
  1989: 66450f3800e2             pshufb %xmm10, %xmm12
  1995: 66410f6fd3               movdqa %xmm11, %xmm2
  2000: 66410fefd4               pxor %xmm12, %xmm2
  2005: 66440f6fcb               movdqa %xmm3, %xmm9
  2010: 66450fdbce               pand %xmm14, %xmm9
  2015: 66440f6fd3               movdqa %xmm3, %xmm10
  2020: 66410f71d204             psrlw $4, %xmm10
  2026: 66450fdbd6               pand %xmm14, %xmm10
  2031: 66450f6f9fc0010000       movdqa 0x1c0(%r15), %xmm11
  2040: 66450f3800d9             pshufb %xmm9, %xmm11
  2046: 66450f6fa700020000       movdqa 0x200(%r15), %xmm12
  2055: 66450f3800e2             pshufb %xmm10, %xmm12
  2061: 66450f6feb               movdqa %xmm11, %xmm13
  2066: 66450fefec               pxor %xmm12, %xmm13
  2071: 66440fefef               pxor %xmm7, %xmm13
  2076: 66440f6fcb               movdqa %xmm3, %xmm9
  2081: 66450fdbce               pand %xmm14, %xmm9
  2086: 66440f6fd3               movdqa %xmm3, %xmm10
  2091: 66410f71d204             psrlw $4, %xmm10
  2097: 66450fdbd6               pand %xmm14, %xmm10
  2102: 66450f6f9f40020000       movdqa 0x240(%r15), %xmm11
  2111: 66450f3800d9             pshufb %xmm9, %xmm11
  2117: 66450f6fa780020000       movdqa 0x280(%r15), %xmm12
  2126: 66450f3800e2             pshufb %xmm10, %xmm12
  2132: 66450f6fc3               movdqa %xmm11, %xmm8
  2137: 66450fefc4               pxor %xmm12, %xmm8
  2142: 66440fefc7               pxor %xmm7, %xmm8
  2147: 66450f6fcd               movdqa %xmm13, %xmm9
  2152: 66450fdbce               pand %xmm14, %xmm9
  2157: 66450f6fd5               movdqa %xmm13, %xmm10
  2162: 66410f71d204             psrlw $4, %xmm10
  2168: 66450fdbd6               pand %xmm14, %xmm10
  2173: 66450f6f9fc0020000       movdqa 0x2c0(%r15), %xmm11
  2182: 66450f3800d9             pshufb %xmm9, %xmm11
  2188: 66450f6fa700030000       movdqa 0x300(%r15), %xmm12
  2197: 66450f3800e2             pshufb %xmm10, %xmm12
  2203: 66410f6ffb               movdqa %xmm11, %xmm7
  2208: 66410feffc               pxor %xmm12, %xmm7
  2213: 66450f6fc8               movdqa %xmm8, %xmm9
  2218: 66450fdbce               pand %xmm14, %xmm9
  2223: 66450f6fd0               movdqa %xmm8, %xmm10
  2228: 66410f71d204             psrlw $4, %xmm10
  2234: 66450fdbd6               pand %xmm14, %xmm10
  2239: 66450f6f9fc0020000       movdqa 0x2c0(%r15), %xmm11
  2248: 66450f3800d9             pshufb %xmm9, %xmm11
  2254: 66450f6fa700030000       movdqa 0x300(%r15), %xmm12
  2263: 66450f3800e2             pshufb %xmm10, %xmm12
  2269: 66410f6fdb               movdqa %xmm11, %xmm3
  2274: 66410fefdc               pxor %xmm12, %xmm3
  2279: 4d8b11                   movq (%r9), %r10
  2282: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
  2288: 66430f7f4c1a10           movdqa %xmm1, 0x10(%r10, %r11)
  2295: 66430f7f541a20           movdqa %xmm2, 0x20(%r10, %r11)
  2302: 66430f7f5c1a30           movdqa %xmm3, 0x30(%r10, %r11)
  2309: 4d8b5108                 movq 8(%r9), %r10
  2313: 66430f7f241a             movdqa %xmm4, (%r10, %r11)
  2319: 66430f7f6c1a10           movdqa %xmm5, 0x10(%r10, %r11)
  2326: 66430f7f741a20           movdqa %xmm6, 0x20(%r10, %r11)
  2333: 66430f7f7c1a30           movdqa %xmm7, 0x30(%r10, %r11)
  2340: 4983c340                 addq $0x40, %r11
  2344: 4c3bd8                   cmpq %rax, %r11
  2347: 0f862ff7ffff             jbe 0x60
  2353: 4c3bd9                   cmpq %rcx, %r11
  2356: 0f83a1020000             jae 0xbdb
  2362: 4153                     pushq %r11
  2364: 498dbf40030000           leaq 0x340(%r15), %rdi
  2371: 498db740040000           leaq 0x440(%r15), %rsi
  2378: 4d8b11                   movq (%r9), %r10
  2381: 4c8b1c24                 movq (%rsp), %r11
  2385: 33db                     xorl %ebx, %ebx
  2387: 498b00                   movq (%r8), %rax
  2390: 420fb61418               movzbl (%rax, %r11), %edx
  2395: 4885d2                   testq %rdx, %rdx
  2398: 0f8422000000             je 0x986
  2404: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2408: 81c0c1000000             addl $0xc1, %eax
  2414: 81f8ff000000             cmpl $0xff, %eax
  2420: 0f8206000000             jb 0x980
  2426: 81c001ffffff             addl $0xffffff01, %eax
  2432: 0fb60406                 movzbl (%rsi, %rax), %eax
  2436: 32d8                     xorb %al, %bl
  2438: 498b4008                 movq 8(%r8), %rax
  2442: 420fb61418               movzbl (%rax, %r11), %edx
  2447: 4885d2                   testq %rdx, %rdx
  2450: 0f841f000000             je 0x9b7
  2456: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2460: 83c003                   addl $3, %eax
  2463: 81f8ff000000             cmpl $0xff, %eax
  2469: 0f8206000000             jb 0x9b1
  2475: 81c001ffffff             addl $0xffffff01, %eax
  2481: 0fb60406                 movzbl (%rsi, %rax), %eax
  2485: 32d8                     xorb %al, %bl
  2487: 498b4010                 movq 0x10(%r8), %rax
  2491: 420fb61418               movzbl (%rax, %r11), %edx
  2496: 4885d2                   testq %rdx, %rdx
  2499: 0f8422000000             je 0x9eb
  2505: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2509: 81c0c6000000             addl $0xc6, %eax
  2515: 81f8ff000000             cmpl $0xff, %eax
  2521: 0f8206000000             jb 0x9e5
  2527: 81c001ffffff             addl $0xffffff01, %eax
  2533: 0fb60406                 movzbl (%rsi, %rax), %eax
  2537: 32d8                     xorb %al, %bl
  2539: 498b4018                 movq 0x18(%r8), %rax
  2543: 420fb61418               movzbl (%rax, %r11), %edx
  2548: 4885d2                   testq %rdx, %rdx
  2551: 0f841f000000             je 0xa1c
  2557: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2561: 83c001                   addl $1, %eax
  2564: 81f8ff000000             cmpl $0xff, %eax
  2570: 0f8206000000             jb 0xa16
  2576: 81c001ffffff             addl $0xffffff01, %eax
  2582: 0fb60406                 movzbl (%rsi, %rax), %eax
  2586: 32d8                     xorb %al, %bl
  2588: 498b4020                 movq 0x20(%r8), %rax
  2592: 420fb61418               movzbl (%rax, %r11), %edx
  2597: 4885d2                   testq %rdx, %rdx
  2600: 0f8422000000             je 0xa50
  2606: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2610: 81c0fe000000             addl $0xfe, %eax
  2616: 81f8ff000000             cmpl $0xff, %eax
  2622: 0f8206000000             jb 0xa4a
  2628: 81c001ffffff             addl $0xffffff01, %eax
  2634: 0fb60406                 movzbl (%rsi, %rax), %eax
  2638: 32d8                     xorb %al, %bl
  2640: 498b4028                 movq 0x28(%r8), %rax
  2644: 420fb61418               movzbl (%rax, %r11), %edx
  2649: 4885d2                   testq %rdx, %rdx
  2652: 0f841f000000             je 0xa81
  2658: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2662: 83c031                   addl $0x31, %eax
  2665: 81f8ff000000             cmpl $0xff, %eax
  2671: 0f8206000000             jb 0xa7b
  2677: 81c001ffffff             addl $0xffffff01, %eax
  2683: 0fb60406                 movzbl (%rsi, %rax), %eax
  2687: 32d8                     xorb %al, %bl
  2689: 43881c1a                 movb %bl, (%r10, %r11)
  2693: 33db                     xorl %ebx, %ebx
  2695: 49ffc3                   incq %r11
  2698: 4c3bd9                   cmpq %rcx, %r11
  2701: 0f82c0feffff             jb 0x953
  2707: 4d8b5108                 movq 8(%r9), %r10
  2711: 4c8b1c24                 movq (%rsp), %r11
  2715: 33db                     xorl %ebx, %ebx
  2717: 498b00                   movq (%r8), %rax
  2720: 420fb61418               movzbl (%rax, %r11), %edx
  2725: 4885d2                   testq %rdx, %rdx
  2728: 0f841f000000             je 0xacd
  2734: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2738: 83c01c                   addl $0x1c, %eax
  2741: 81f8ff000000             cmpl $0xff, %eax
  2747: 0f8206000000             jb 0xac7
  2753: 81c001ffffff             addl $0xffffff01, %eax
  2759: 0fb60406                 movzbl (%rsi, %rax), %eax
  2763: 32d8                     xorb %al, %bl
  2765: 498b4008                 movq 8(%r8), %rax
  2769: 420fb61418               movzbl (%rax, %r11), %edx
  2774: 4885d2                   testq %rdx, %rdx
  2777: 0f8422000000             je 0xb01
  2783: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2787: 81c0df000000             addl $0xdf, %eax
  2793: 81f8ff000000             cmpl $0xff, %eax
  2799: 0f8206000000             jb 0xafb
  2805: 81c001ffffff             addl $0xffffff01, %eax
  2811: 0fb60406                 movzbl (%rsi, %rax), %eax
  2815: 32d8                     xorb %al, %bl
  2817: 498b4010                 movq 0x10(%r8), %rax
  2821: 420fb61418               movzbl (%rax, %r11), %edx
  2826: 4885d2                   testq %rdx, %rdx
  2829: 0f841f000000             je 0xb32
  2835: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2839: 83c01a                   addl $0x1a, %eax
  2842: 81f8ff000000             cmpl $0xff, %eax
  2848: 0f8206000000             jb 0xb2c
  2854: 81c001ffffff             addl $0xffffff01, %eax
  2860: 0fb60406                 movzbl (%rsi, %rax), %eax
  2864: 32d8                     xorb %al, %bl
  2866: 498b4018                 movq 0x18(%r8), %rax
  2870: 420fb61418               movzbl (%rax, %r11), %edx
  2875: 4885d2                   testq %rdx, %rdx
  2878: 0f841f000000             je 0xb63
  2884: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2888: 83c019                   addl $0x19, %eax
  2891: 81f8ff000000             cmpl $0xff, %eax
  2897: 0f8206000000             jb 0xb5d
  2903: 81c001ffffff             addl $0xffffff01, %eax
  2909: 0fb60406                 movzbl (%rsi, %rax), %eax
  2913: 32d8                     xorb %al, %bl
  2915: 498b4020                 movq 0x20(%r8), %rax
  2919: 420fb61418               movzbl (%rax, %r11), %edx
  2924: 4885d2                   testq %rdx, %rdx
  2927: 0f841f000000             je 0xb94
  2933: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2937: 83c018                   addl $0x18, %eax
  2940: 81f8ff000000             cmpl $0xff, %eax
  2946: 0f8206000000             jb 0xb8e
  2952: 81c001ffffff             addl $0xffffff01, %eax
  2958: 0fb60406                 movzbl (%rsi, %rax), %eax
  2962: 32d8                     xorb %al, %bl
  2964: 498b4028                 movq 0x28(%r8), %rax
  2968: 420fb61418               movzbl (%rax, %r11), %edx
  2973: 4885d2                   testq %rdx, %rdx
  2976: 0f841f000000             je 0xbc5
  2982: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2986: 83c031                   addl $0x31, %eax
  2989: 81f8ff000000             cmpl $0xff, %eax
  2995: 0f8206000000             jb 0xbbf
  3001: 81c001ffffff             addl $0xffffff01, %eax
  3007: 0fb60406                 movzbl (%rsi, %rax), %eax
  3011: 32d8                     xorb %al, %bl
  3013: 43881c1a                 movb %bl, (%r10, %r11)
  3017: 33db                     xorl %ebx, %ebx
  3019: 49ffc3                   incq %r11
  3022: 4c3bd9                   cmpq %rcx, %r11
  3025: 0f82c6feffff             jb 0xa9d
  3031: 4883c408                 addq $8, %rsp
  3035: 415e                     popq %r14
  3037: 415d                     popq %r13
  3039: 415c                     popq %r12
  3041: 415f                     popq %r15
  3043: 5b                       popq %rbx
  3044: c3                       retq 
  3045: 0f0b                     ud2 
  3047: 0f0b                     ud2 
  3049: 0f0b                     ud2 
  3051: 0f0b                     ud2 
  3053: 0f0b                     ud2 
  3055: 0f0b                     ud2 
  3057: 0f0b                     ud2 
  3059: 0f0b                     ud2 
  3061: 0f0b                     ud2 
  3063: 0f0b                     ud2 
  3065: 0f0b                     ud2 
  3067: 0f0b                     ud2 
  3069: 0f0b                     ud2 
  3071: cc                       int3 
