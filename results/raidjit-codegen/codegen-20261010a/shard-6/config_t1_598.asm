; config_t1_598 target=sse twin=baseline topo=raid6(8)/enc op=encode U=1
; code_size=3264 twins=[] uops_per_byte=6.921875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da70c0000           leaq 0xca7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66450f6f37               movdqa (%r15), %xmm14
    52: 488bc1                   movq %rcx, %rax
    55: 4883e840                 subq $0x40, %rax
    59: 0f82460a0000             jb 0xa87
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
   133: 66450f6f6f40             movdqa 0x40(%r15), %xmm13
   139: 66450f6fdd               movdqa %xmm13, %xmm11
   144: 66450f3800d9             pshufb %xmm9, %xmm11
   150: 66450f6faf80000000       movdqa 0x80(%r15), %xmm13
   159: 66450f6fe5               movdqa %xmm13, %xmm12
   164: 66450f3800e2             pshufb %xmm10, %xmm12
   170: 66410f6fe3               movdqa %xmm11, %xmm4
   175: 66410fefe4               pxor %xmm12, %xmm4
   180: 66460f6f441e10           movdqa 0x10(%rsi, %r11), %xmm8
   187: 66450f6fc8               movdqa %xmm8, %xmm9
   192: 66450fdbce               pand %xmm14, %xmm9
   197: 66450f6fd0               movdqa %xmm8, %xmm10
   202: 66410f71d204             psrlw $4, %xmm10
   208: 66450fdbd6               pand %xmm14, %xmm10
   213: 66410f6fc8               movdqa %xmm8, %xmm1
   218: 66450f6f6f40             movdqa 0x40(%r15), %xmm13
   224: 66450f6fdd               movdqa %xmm13, %xmm11
   229: 66450f3800d9             pshufb %xmm9, %xmm11
   235: 66450f6faf80000000       movdqa 0x80(%r15), %xmm13
   244: 66450f6fe5               movdqa %xmm13, %xmm12
   249: 66450f3800e2             pshufb %xmm10, %xmm12
   255: 66410f6feb               movdqa %xmm11, %xmm5
   260: 66410fefec               pxor %xmm12, %xmm5
   265: 66460f6f441e20           movdqa 0x20(%rsi, %r11), %xmm8
   272: 66450f6fc8               movdqa %xmm8, %xmm9
   277: 66450fdbce               pand %xmm14, %xmm9
   282: 66450f6fd0               movdqa %xmm8, %xmm10
   287: 66410f71d204             psrlw $4, %xmm10
   293: 66450fdbd6               pand %xmm14, %xmm10
   298: 66410f6fd0               movdqa %xmm8, %xmm2
   303: 66450f6f6f40             movdqa 0x40(%r15), %xmm13
   309: 66450f6fdd               movdqa %xmm13, %xmm11
   314: 66450f3800d9             pshufb %xmm9, %xmm11
   320: 66450f6faf80000000       movdqa 0x80(%r15), %xmm13
   329: 66450f6fe5               movdqa %xmm13, %xmm12
   334: 66450f3800e2             pshufb %xmm10, %xmm12
   340: 66410f6ff3               movdqa %xmm11, %xmm6
   345: 66410feff4               pxor %xmm12, %xmm6
   350: 66460f6f441e30           movdqa 0x30(%rsi, %r11), %xmm8
   357: 66450f6fc8               movdqa %xmm8, %xmm9
   362: 66450fdbce               pand %xmm14, %xmm9
   367: 66450f6fd0               movdqa %xmm8, %xmm10
   372: 66410f71d204             psrlw $4, %xmm10
   378: 66450fdbd6               pand %xmm14, %xmm10
   383: 66410f6fd8               movdqa %xmm8, %xmm3
   388: 66450f6f6f40             movdqa 0x40(%r15), %xmm13
   394: 66450f6fdd               movdqa %xmm13, %xmm11
   399: 66450f3800d9             pshufb %xmm9, %xmm11
   405: 66450f6faf80000000       movdqa 0x80(%r15), %xmm13
   414: 66450f6fe5               movdqa %xmm13, %xmm12
   419: 66450f3800e2             pshufb %xmm10, %xmm12
   425: 66410f6ffb               movdqa %xmm11, %xmm7
   430: 66410feffc               pxor %xmm12, %xmm7
   435: 66460f6f041f             movdqa (%rdi, %r11), %xmm8
   441: 66450f6fc8               movdqa %xmm8, %xmm9
   446: 66450fdbce               pand %xmm14, %xmm9
   451: 66450f6fd0               movdqa %xmm8, %xmm10
   456: 66410f71d204             psrlw $4, %xmm10
   462: 66450fdbd6               pand %xmm14, %xmm10
   467: 66410fefc0               pxor %xmm8, %xmm0
   472: 66450f6fafc0000000       movdqa 0xc0(%r15), %xmm13
   481: 66450f6fdd               movdqa %xmm13, %xmm11
   486: 66450f3800d9             pshufb %xmm9, %xmm11
   492: 66450f6faf00010000       movdqa 0x100(%r15), %xmm13
   501: 66450f6fe5               movdqa %xmm13, %xmm12
   506: 66450f3800e2             pshufb %xmm10, %xmm12
   512: 66410fefe3               pxor %xmm11, %xmm4
   517: 66410fefe4               pxor %xmm12, %xmm4
   522: 66460f6f441f10           movdqa 0x10(%rdi, %r11), %xmm8
   529: 66450f6fc8               movdqa %xmm8, %xmm9
   534: 66450fdbce               pand %xmm14, %xmm9
   539: 66450f6fd0               movdqa %xmm8, %xmm10
   544: 66410f71d204             psrlw $4, %xmm10
   550: 66450fdbd6               pand %xmm14, %xmm10
   555: 66410fefc8               pxor %xmm8, %xmm1
   560: 66450f6fafc0000000       movdqa 0xc0(%r15), %xmm13
   569: 66450f6fdd               movdqa %xmm13, %xmm11
   574: 66450f3800d9             pshufb %xmm9, %xmm11
   580: 66450f6faf00010000       movdqa 0x100(%r15), %xmm13
   589: 66450f6fe5               movdqa %xmm13, %xmm12
   594: 66450f3800e2             pshufb %xmm10, %xmm12
   600: 66410fefeb               pxor %xmm11, %xmm5
   605: 66410fefec               pxor %xmm12, %xmm5
   610: 66460f6f441f20           movdqa 0x20(%rdi, %r11), %xmm8
   617: 66450f6fc8               movdqa %xmm8, %xmm9
   622: 66450fdbce               pand %xmm14, %xmm9
   627: 66450f6fd0               movdqa %xmm8, %xmm10
   632: 66410f71d204             psrlw $4, %xmm10
   638: 66450fdbd6               pand %xmm14, %xmm10
   643: 66410fefd0               pxor %xmm8, %xmm2
   648: 66450f6fafc0000000       movdqa 0xc0(%r15), %xmm13
   657: 66450f6fdd               movdqa %xmm13, %xmm11
   662: 66450f3800d9             pshufb %xmm9, %xmm11
   668: 66450f6faf00010000       movdqa 0x100(%r15), %xmm13
   677: 66450f6fe5               movdqa %xmm13, %xmm12
   682: 66450f3800e2             pshufb %xmm10, %xmm12
   688: 66410feff3               pxor %xmm11, %xmm6
   693: 66410feff4               pxor %xmm12, %xmm6
   698: 66460f6f441f30           movdqa 0x30(%rdi, %r11), %xmm8
   705: 66450f6fc8               movdqa %xmm8, %xmm9
   710: 66450fdbce               pand %xmm14, %xmm9
   715: 66450f6fd0               movdqa %xmm8, %xmm10
   720: 66410f71d204             psrlw $4, %xmm10
   726: 66450fdbd6               pand %xmm14, %xmm10
   731: 66410fefd8               pxor %xmm8, %xmm3
   736: 66450f6fafc0000000       movdqa 0xc0(%r15), %xmm13
   745: 66450f6fdd               movdqa %xmm13, %xmm11
   750: 66450f3800d9             pshufb %xmm9, %xmm11
   756: 66450f6faf00010000       movdqa 0x100(%r15), %xmm13
   765: 66450f6fe5               movdqa %xmm13, %xmm12
   770: 66450f3800e2             pshufb %xmm10, %xmm12
   776: 66410feffb               pxor %xmm11, %xmm7
   781: 66410feffc               pxor %xmm12, %xmm7
   786: 66470f6f041c             movdqa (%r12, %r11), %xmm8
   792: 66450f6fc8               movdqa %xmm8, %xmm9
   797: 66450fdbce               pand %xmm14, %xmm9
   802: 66450f6fd0               movdqa %xmm8, %xmm10
   807: 66410f71d204             psrlw $4, %xmm10
   813: 66450fdbd6               pand %xmm14, %xmm10
   818: 66410fefc0               pxor %xmm8, %xmm0
   823: 66450f6faf40010000       movdqa 0x140(%r15), %xmm13
   832: 66450f6fdd               movdqa %xmm13, %xmm11
   837: 66450f3800d9             pshufb %xmm9, %xmm11
   843: 66450f6faf80010000       movdqa 0x180(%r15), %xmm13
   852: 66450f6fe5               movdqa %xmm13, %xmm12
   857: 66450f3800e2             pshufb %xmm10, %xmm12
   863: 66410fefe3               pxor %xmm11, %xmm4
   868: 66410fefe4               pxor %xmm12, %xmm4
   873: 66470f6f441c10           movdqa 0x10(%r12, %r11), %xmm8
   880: 66450f6fc8               movdqa %xmm8, %xmm9
   885: 66450fdbce               pand %xmm14, %xmm9
   890: 66450f6fd0               movdqa %xmm8, %xmm10
   895: 66410f71d204             psrlw $4, %xmm10
   901: 66450fdbd6               pand %xmm14, %xmm10
   906: 66410fefc8               pxor %xmm8, %xmm1
   911: 66450f6faf40010000       movdqa 0x140(%r15), %xmm13
   920: 66450f6fdd               movdqa %xmm13, %xmm11
   925: 66450f3800d9             pshufb %xmm9, %xmm11
   931: 66450f6faf80010000       movdqa 0x180(%r15), %xmm13
   940: 66450f6fe5               movdqa %xmm13, %xmm12
   945: 66450f3800e2             pshufb %xmm10, %xmm12
   951: 66410fefeb               pxor %xmm11, %xmm5
   956: 66410fefec               pxor %xmm12, %xmm5
   961: 66470f6f441c20           movdqa 0x20(%r12, %r11), %xmm8
   968: 66450f6fc8               movdqa %xmm8, %xmm9
   973: 66450fdbce               pand %xmm14, %xmm9
   978: 66450f6fd0               movdqa %xmm8, %xmm10
   983: 66410f71d204             psrlw $4, %xmm10
   989: 66450fdbd6               pand %xmm14, %xmm10
   994: 66410fefd0               pxor %xmm8, %xmm2
   999: 66450f6faf40010000       movdqa 0x140(%r15), %xmm13
  1008: 66450f6fdd               movdqa %xmm13, %xmm11
  1013: 66450f3800d9             pshufb %xmm9, %xmm11
  1019: 66450f6faf80010000       movdqa 0x180(%r15), %xmm13
  1028: 66450f6fe5               movdqa %xmm13, %xmm12
  1033: 66450f3800e2             pshufb %xmm10, %xmm12
  1039: 66410feff3               pxor %xmm11, %xmm6
  1044: 66410feff4               pxor %xmm12, %xmm6
  1049: 66470f6f441c30           movdqa 0x30(%r12, %r11), %xmm8
  1056: 66450f6fc8               movdqa %xmm8, %xmm9
  1061: 66450fdbce               pand %xmm14, %xmm9
  1066: 66450f6fd0               movdqa %xmm8, %xmm10
  1071: 66410f71d204             psrlw $4, %xmm10
  1077: 66450fdbd6               pand %xmm14, %xmm10
  1082: 66410fefd8               pxor %xmm8, %xmm3
  1087: 66450f6faf40010000       movdqa 0x140(%r15), %xmm13
  1096: 66450f6fdd               movdqa %xmm13, %xmm11
  1101: 66450f3800d9             pshufb %xmm9, %xmm11
  1107: 66450f6faf80010000       movdqa 0x180(%r15), %xmm13
  1116: 66450f6fe5               movdqa %xmm13, %xmm12
  1121: 66450f3800e2             pshufb %xmm10, %xmm12
  1127: 66410feffb               pxor %xmm11, %xmm7
  1132: 66410feffc               pxor %xmm12, %xmm7
  1137: 66470f6f441d00           movdqa (%r13, %r11), %xmm8
  1144: 66450f6fc8               movdqa %xmm8, %xmm9
  1149: 66450fdbce               pand %xmm14, %xmm9
  1154: 66450f6fd0               movdqa %xmm8, %xmm10
  1159: 66410f71d204             psrlw $4, %xmm10
  1165: 66450fdbd6               pand %xmm14, %xmm10
  1170: 66410fefc0               pxor %xmm8, %xmm0
  1175: 66450f6fafc0010000       movdqa 0x1c0(%r15), %xmm13
  1184: 66450f6fdd               movdqa %xmm13, %xmm11
  1189: 66450f3800d9             pshufb %xmm9, %xmm11
  1195: 66450f6faf00020000       movdqa 0x200(%r15), %xmm13
  1204: 66450f6fe5               movdqa %xmm13, %xmm12
  1209: 66450f3800e2             pshufb %xmm10, %xmm12
  1215: 66410fefe3               pxor %xmm11, %xmm4
  1220: 66410fefe4               pxor %xmm12, %xmm4
  1225: 66470f6f441d10           movdqa 0x10(%r13, %r11), %xmm8
  1232: 66450f6fc8               movdqa %xmm8, %xmm9
  1237: 66450fdbce               pand %xmm14, %xmm9
  1242: 66450f6fd0               movdqa %xmm8, %xmm10
  1247: 66410f71d204             psrlw $4, %xmm10
  1253: 66450fdbd6               pand %xmm14, %xmm10
  1258: 66410fefc8               pxor %xmm8, %xmm1
  1263: 66450f6fafc0010000       movdqa 0x1c0(%r15), %xmm13
  1272: 66450f6fdd               movdqa %xmm13, %xmm11
  1277: 66450f3800d9             pshufb %xmm9, %xmm11
  1283: 66450f6faf00020000       movdqa 0x200(%r15), %xmm13
  1292: 66450f6fe5               movdqa %xmm13, %xmm12
  1297: 66450f3800e2             pshufb %xmm10, %xmm12
  1303: 66410fefeb               pxor %xmm11, %xmm5
  1308: 66410fefec               pxor %xmm12, %xmm5
  1313: 66470f6f441d20           movdqa 0x20(%r13, %r11), %xmm8
  1320: 66450f6fc8               movdqa %xmm8, %xmm9
  1325: 66450fdbce               pand %xmm14, %xmm9
  1330: 66450f6fd0               movdqa %xmm8, %xmm10
  1335: 66410f71d204             psrlw $4, %xmm10
  1341: 66450fdbd6               pand %xmm14, %xmm10
  1346: 66410fefd0               pxor %xmm8, %xmm2
  1351: 66450f6fafc0010000       movdqa 0x1c0(%r15), %xmm13
  1360: 66450f6fdd               movdqa %xmm13, %xmm11
  1365: 66450f3800d9             pshufb %xmm9, %xmm11
  1371: 66450f6faf00020000       movdqa 0x200(%r15), %xmm13
  1380: 66450f6fe5               movdqa %xmm13, %xmm12
  1385: 66450f3800e2             pshufb %xmm10, %xmm12
  1391: 66410feff3               pxor %xmm11, %xmm6
  1396: 66410feff4               pxor %xmm12, %xmm6
  1401: 66470f6f441d30           movdqa 0x30(%r13, %r11), %xmm8
  1408: 66450f6fc8               movdqa %xmm8, %xmm9
  1413: 66450fdbce               pand %xmm14, %xmm9
  1418: 66450f6fd0               movdqa %xmm8, %xmm10
  1423: 66410f71d204             psrlw $4, %xmm10
  1429: 66450fdbd6               pand %xmm14, %xmm10
  1434: 66410fefd8               pxor %xmm8, %xmm3
  1439: 66450f6fafc0010000       movdqa 0x1c0(%r15), %xmm13
  1448: 66450f6fdd               movdqa %xmm13, %xmm11
  1453: 66450f3800d9             pshufb %xmm9, %xmm11
  1459: 66450f6faf00020000       movdqa 0x200(%r15), %xmm13
  1468: 66450f6fe5               movdqa %xmm13, %xmm12
  1473: 66450f3800e2             pshufb %xmm10, %xmm12
  1479: 66410feffb               pxor %xmm11, %xmm7
  1484: 66410feffc               pxor %xmm12, %xmm7
  1489: 66470f6f041e             movdqa (%r14, %r11), %xmm8
  1495: 66450f6fc8               movdqa %xmm8, %xmm9
  1500: 66450fdbce               pand %xmm14, %xmm9
  1505: 66450f6fd0               movdqa %xmm8, %xmm10
  1510: 66410f71d204             psrlw $4, %xmm10
  1516: 66450fdbd6               pand %xmm14, %xmm10
  1521: 66410fefc0               pxor %xmm8, %xmm0
  1526: 66450f6faf40020000       movdqa 0x240(%r15), %xmm13
  1535: 66450f6fdd               movdqa %xmm13, %xmm11
  1540: 66450f3800d9             pshufb %xmm9, %xmm11
  1546: 66450f6faf80020000       movdqa 0x280(%r15), %xmm13
  1555: 66450f6fe5               movdqa %xmm13, %xmm12
  1560: 66450f3800e2             pshufb %xmm10, %xmm12
  1566: 66410fefe3               pxor %xmm11, %xmm4
  1571: 66410fefe4               pxor %xmm12, %xmm4
  1576: 66470f6f441e10           movdqa 0x10(%r14, %r11), %xmm8
  1583: 66450f6fc8               movdqa %xmm8, %xmm9
  1588: 66450fdbce               pand %xmm14, %xmm9
  1593: 66450f6fd0               movdqa %xmm8, %xmm10
  1598: 66410f71d204             psrlw $4, %xmm10
  1604: 66450fdbd6               pand %xmm14, %xmm10
  1609: 66410fefc8               pxor %xmm8, %xmm1
  1614: 66450f6faf40020000       movdqa 0x240(%r15), %xmm13
  1623: 66450f6fdd               movdqa %xmm13, %xmm11
  1628: 66450f3800d9             pshufb %xmm9, %xmm11
  1634: 66450f6faf80020000       movdqa 0x280(%r15), %xmm13
  1643: 66450f6fe5               movdqa %xmm13, %xmm12
  1648: 66450f3800e2             pshufb %xmm10, %xmm12
  1654: 66410fefeb               pxor %xmm11, %xmm5
  1659: 66410fefec               pxor %xmm12, %xmm5
  1664: 66470f6f441e20           movdqa 0x20(%r14, %r11), %xmm8
  1671: 66450f6fc8               movdqa %xmm8, %xmm9
  1676: 66450fdbce               pand %xmm14, %xmm9
  1681: 66450f6fd0               movdqa %xmm8, %xmm10
  1686: 66410f71d204             psrlw $4, %xmm10
  1692: 66450fdbd6               pand %xmm14, %xmm10
  1697: 66410fefd0               pxor %xmm8, %xmm2
  1702: 66450f6faf40020000       movdqa 0x240(%r15), %xmm13
  1711: 66450f6fdd               movdqa %xmm13, %xmm11
  1716: 66450f3800d9             pshufb %xmm9, %xmm11
  1722: 66450f6faf80020000       movdqa 0x280(%r15), %xmm13
  1731: 66450f6fe5               movdqa %xmm13, %xmm12
  1736: 66450f3800e2             pshufb %xmm10, %xmm12
  1742: 66410feff3               pxor %xmm11, %xmm6
  1747: 66410feff4               pxor %xmm12, %xmm6
  1752: 66470f6f441e30           movdqa 0x30(%r14, %r11), %xmm8
  1759: 66450f6fc8               movdqa %xmm8, %xmm9
  1764: 66450fdbce               pand %xmm14, %xmm9
  1769: 66450f6fd0               movdqa %xmm8, %xmm10
  1774: 66410f71d204             psrlw $4, %xmm10
  1780: 66450fdbd6               pand %xmm14, %xmm10
  1785: 66410fefd8               pxor %xmm8, %xmm3
  1790: 66450f6faf40020000       movdqa 0x240(%r15), %xmm13
  1799: 66450f6fdd               movdqa %xmm13, %xmm11
  1804: 66450f3800d9             pshufb %xmm9, %xmm11
  1810: 66450f6faf80020000       movdqa 0x280(%r15), %xmm13
  1819: 66450f6fe5               movdqa %xmm13, %xmm12
  1824: 66450f3800e2             pshufb %xmm10, %xmm12
  1830: 66410feffb               pxor %xmm11, %xmm7
  1835: 66410feffc               pxor %xmm12, %xmm7
  1840: 4d8b5028                 movq 0x28(%r8), %r10
  1844: 66470f6f041a             movdqa (%r10, %r11), %xmm8
  1850: 66450f6fc8               movdqa %xmm8, %xmm9
  1855: 66450fdbce               pand %xmm14, %xmm9
  1860: 66450f6fd0               movdqa %xmm8, %xmm10
  1865: 66410f71d204             psrlw $4, %xmm10
  1871: 66450fdbd6               pand %xmm14, %xmm10
  1876: 66410fefc0               pxor %xmm8, %xmm0
  1881: 66450f6fafc0020000       movdqa 0x2c0(%r15), %xmm13
  1890: 66450f6fdd               movdqa %xmm13, %xmm11
  1895: 66450f3800d9             pshufb %xmm9, %xmm11
  1901: 66450f6faf00030000       movdqa 0x300(%r15), %xmm13
  1910: 66450f6fe5               movdqa %xmm13, %xmm12
  1915: 66450f3800e2             pshufb %xmm10, %xmm12
  1921: 66410fefe3               pxor %xmm11, %xmm4
  1926: 66410fefe4               pxor %xmm12, %xmm4
  1931: 66470f6f441a10           movdqa 0x10(%r10, %r11), %xmm8
  1938: 66450f6fc8               movdqa %xmm8, %xmm9
  1943: 66450fdbce               pand %xmm14, %xmm9
  1948: 66450f6fd0               movdqa %xmm8, %xmm10
  1953: 66410f71d204             psrlw $4, %xmm10
  1959: 66450fdbd6               pand %xmm14, %xmm10
  1964: 66410fefc8               pxor %xmm8, %xmm1
  1969: 66450f6fafc0020000       movdqa 0x2c0(%r15), %xmm13
  1978: 66450f6fdd               movdqa %xmm13, %xmm11
  1983: 66450f3800d9             pshufb %xmm9, %xmm11
  1989: 66450f6faf00030000       movdqa 0x300(%r15), %xmm13
  1998: 66450f6fe5               movdqa %xmm13, %xmm12
  2003: 66450f3800e2             pshufb %xmm10, %xmm12
  2009: 66410fefeb               pxor %xmm11, %xmm5
  2014: 66410fefec               pxor %xmm12, %xmm5
  2019: 66470f6f441a20           movdqa 0x20(%r10, %r11), %xmm8
  2026: 66450f6fc8               movdqa %xmm8, %xmm9
  2031: 66450fdbce               pand %xmm14, %xmm9
  2036: 66450f6fd0               movdqa %xmm8, %xmm10
  2041: 66410f71d204             psrlw $4, %xmm10
  2047: 66450fdbd6               pand %xmm14, %xmm10
  2052: 66410fefd0               pxor %xmm8, %xmm2
  2057: 66450f6fafc0020000       movdqa 0x2c0(%r15), %xmm13
  2066: 66450f6fdd               movdqa %xmm13, %xmm11
  2071: 66450f3800d9             pshufb %xmm9, %xmm11
  2077: 66450f6faf00030000       movdqa 0x300(%r15), %xmm13
  2086: 66450f6fe5               movdqa %xmm13, %xmm12
  2091: 66450f3800e2             pshufb %xmm10, %xmm12
  2097: 66410feff3               pxor %xmm11, %xmm6
  2102: 66410feff4               pxor %xmm12, %xmm6
  2107: 66470f6f441a30           movdqa 0x30(%r10, %r11), %xmm8
  2114: 66450f6fc8               movdqa %xmm8, %xmm9
  2119: 66450fdbce               pand %xmm14, %xmm9
  2124: 66450f6fd0               movdqa %xmm8, %xmm10
  2129: 66410f71d204             psrlw $4, %xmm10
  2135: 66450fdbd6               pand %xmm14, %xmm10
  2140: 66410fefd8               pxor %xmm8, %xmm3
  2145: 66450f6fafc0020000       movdqa 0x2c0(%r15), %xmm13
  2154: 66450f6fdd               movdqa %xmm13, %xmm11
  2159: 66450f3800d9             pshufb %xmm9, %xmm11
  2165: 66450f6faf00030000       movdqa 0x300(%r15), %xmm13
  2174: 66450f6fe5               movdqa %xmm13, %xmm12
  2179: 66450f3800e2             pshufb %xmm10, %xmm12
  2185: 66410feffb               pxor %xmm11, %xmm7
  2190: 66410feffc               pxor %xmm12, %xmm7
  2195: 4d8b5030                 movq 0x30(%r8), %r10
  2199: 66470f6f041a             movdqa (%r10, %r11), %xmm8
  2205: 66450f6fc8               movdqa %xmm8, %xmm9
  2210: 66450fdbce               pand %xmm14, %xmm9
  2215: 66450f6fd0               movdqa %xmm8, %xmm10
  2220: 66410f71d204             psrlw $4, %xmm10
  2226: 66450fdbd6               pand %xmm14, %xmm10
  2231: 66410fefc0               pxor %xmm8, %xmm0
  2236: 66450f6faf40030000       movdqa 0x340(%r15), %xmm13
  2245: 66450f6fdd               movdqa %xmm13, %xmm11
  2250: 66450f3800d9             pshufb %xmm9, %xmm11
  2256: 66450f6faf80030000       movdqa 0x380(%r15), %xmm13
  2265: 66450f6fe5               movdqa %xmm13, %xmm12
  2270: 66450f3800e2             pshufb %xmm10, %xmm12
  2276: 66410fefe3               pxor %xmm11, %xmm4
  2281: 66410fefe4               pxor %xmm12, %xmm4
  2286: 66470f6f441a10           movdqa 0x10(%r10, %r11), %xmm8
  2293: 66450f6fc8               movdqa %xmm8, %xmm9
  2298: 66450fdbce               pand %xmm14, %xmm9
  2303: 66450f6fd0               movdqa %xmm8, %xmm10
  2308: 66410f71d204             psrlw $4, %xmm10
  2314: 66450fdbd6               pand %xmm14, %xmm10
  2319: 66410fefc8               pxor %xmm8, %xmm1
  2324: 66450f6faf40030000       movdqa 0x340(%r15), %xmm13
  2333: 66450f6fdd               movdqa %xmm13, %xmm11
  2338: 66450f3800d9             pshufb %xmm9, %xmm11
  2344: 66450f6faf80030000       movdqa 0x380(%r15), %xmm13
  2353: 66450f6fe5               movdqa %xmm13, %xmm12
  2358: 66450f3800e2             pshufb %xmm10, %xmm12
  2364: 66410fefeb               pxor %xmm11, %xmm5
  2369: 66410fefec               pxor %xmm12, %xmm5
  2374: 66470f6f441a20           movdqa 0x20(%r10, %r11), %xmm8
  2381: 66450f6fc8               movdqa %xmm8, %xmm9
  2386: 66450fdbce               pand %xmm14, %xmm9
  2391: 66450f6fd0               movdqa %xmm8, %xmm10
  2396: 66410f71d204             psrlw $4, %xmm10
  2402: 66450fdbd6               pand %xmm14, %xmm10
  2407: 66410fefd0               pxor %xmm8, %xmm2
  2412: 66450f6faf40030000       movdqa 0x340(%r15), %xmm13
  2421: 66450f6fdd               movdqa %xmm13, %xmm11
  2426: 66450f3800d9             pshufb %xmm9, %xmm11
  2432: 66450f6faf80030000       movdqa 0x380(%r15), %xmm13
  2441: 66450f6fe5               movdqa %xmm13, %xmm12
  2446: 66450f3800e2             pshufb %xmm10, %xmm12
  2452: 66410feff3               pxor %xmm11, %xmm6
  2457: 66410feff4               pxor %xmm12, %xmm6
  2462: 66470f6f441a30           movdqa 0x30(%r10, %r11), %xmm8
  2469: 66450f6fc8               movdqa %xmm8, %xmm9
  2474: 66450fdbce               pand %xmm14, %xmm9
  2479: 66450f6fd0               movdqa %xmm8, %xmm10
  2484: 66410f71d204             psrlw $4, %xmm10
  2490: 66450fdbd6               pand %xmm14, %xmm10
  2495: 66410fefd8               pxor %xmm8, %xmm3
  2500: 66450f6faf40030000       movdqa 0x340(%r15), %xmm13
  2509: 66450f6fdd               movdqa %xmm13, %xmm11
  2514: 66450f3800d9             pshufb %xmm9, %xmm11
  2520: 66450f6faf80030000       movdqa 0x380(%r15), %xmm13
  2529: 66450f6fe5               movdqa %xmm13, %xmm12
  2534: 66450f3800e2             pshufb %xmm10, %xmm12
  2540: 66410feffb               pxor %xmm11, %xmm7
  2545: 66410feffc               pxor %xmm12, %xmm7
  2550: 4d8b5038                 movq 0x38(%r8), %r10
  2554: 66470f6f041a             movdqa (%r10, %r11), %xmm8
  2560: 66410fefc0               pxor %xmm8, %xmm0
  2565: 66410fefe0               pxor %xmm8, %xmm4
  2570: 66470f6f441a10           movdqa 0x10(%r10, %r11), %xmm8
  2577: 66410fefc8               pxor %xmm8, %xmm1
  2582: 66410fefe8               pxor %xmm8, %xmm5
  2587: 66470f6f441a20           movdqa 0x20(%r10, %r11), %xmm8
  2594: 66410fefd0               pxor %xmm8, %xmm2
  2599: 66410feff0               pxor %xmm8, %xmm6
  2604: 66470f6f441a30           movdqa 0x30(%r10, %r11), %xmm8
  2611: 66410fefd8               pxor %xmm8, %xmm3
  2616: 66410feff8               pxor %xmm8, %xmm7
  2621: 4d8b11                   movq (%r9), %r10
  2624: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
  2630: 66430f7f4c1a10           movdqa %xmm1, 0x10(%r10, %r11)
  2637: 66430f7f541a20           movdqa %xmm2, 0x20(%r10, %r11)
  2644: 66430f7f5c1a30           movdqa %xmm3, 0x30(%r10, %r11)
  2651: 4d8b5108                 movq 8(%r9), %r10
  2655: 66430f7f241a             movdqa %xmm4, (%r10, %r11)
  2661: 66430f7f6c1a10           movdqa %xmm5, 0x10(%r10, %r11)
  2668: 66430f7f741a20           movdqa %xmm6, 0x20(%r10, %r11)
  2675: 66430f7f7c1a30           movdqa %xmm7, 0x30(%r10, %r11)
  2682: 4983c340                 addq $0x40, %r11
  2686: 4c3bd8                   cmpq %rax, %r11
  2689: 0f86d9f5ffff             jbe 0x60
  2695: 4c3bd9                   cmpq %rcx, %r11
  2698: 0f830f020000             jae 0xc9f
  2704: 4153                     pushq %r11
  2706: 498dbfc0030000           leaq 0x3c0(%r15), %rdi
  2713: 498db7c0040000           leaq 0x4c0(%r15), %rsi
  2720: 4d8b11                   movq (%r9), %r10
  2723: 4c8b1c24                 movq (%rsp), %r11
  2727: 33db                     xorl %ebx, %ebx
  2729: 498b00                   movq (%r8), %rax
  2732: 420fb61418               movzbl (%rax, %r11), %edx
  2737: 32da                     xorb %dl, %bl
  2739: 498b4008                 movq 8(%r8), %rax
  2743: 420fb61418               movzbl (%rax, %r11), %edx
  2748: 32da                     xorb %dl, %bl
  2750: 498b4010                 movq 0x10(%r8), %rax
  2754: 420fb61418               movzbl (%rax, %r11), %edx
  2759: 32da                     xorb %dl, %bl
  2761: 498b4018                 movq 0x18(%r8), %rax
  2765: 420fb61418               movzbl (%rax, %r11), %edx
  2770: 32da                     xorb %dl, %bl
  2772: 498b4020                 movq 0x20(%r8), %rax
  2776: 420fb61418               movzbl (%rax, %r11), %edx
  2781: 32da                     xorb %dl, %bl
  2783: 498b4028                 movq 0x28(%r8), %rax
  2787: 420fb61418               movzbl (%rax, %r11), %edx
  2792: 32da                     xorb %dl, %bl
  2794: 498b4030                 movq 0x30(%r8), %rax
  2798: 420fb61418               movzbl (%rax, %r11), %edx
  2803: 32da                     xorb %dl, %bl
  2805: 498b4038                 movq 0x38(%r8), %rax
  2809: 420fb61418               movzbl (%rax, %r11), %edx
  2814: 32da                     xorb %dl, %bl
  2816: 43881c1a                 movb %bl, (%r10, %r11)
  2820: 33db                     xorl %ebx, %ebx
  2822: 49ffc3                   incq %r11
  2825: 4c3bd9                   cmpq %rcx, %r11
  2828: 0f8297ffffff             jb 0xaa9
  2834: 4d8b5108                 movq 8(%r9), %r10
  2838: 4c8b1c24                 movq (%rsp), %r11
  2842: 33db                     xorl %ebx, %ebx
  2844: 498b00                   movq (%r8), %rax
  2847: 420fb61418               movzbl (%rax, %r11), %edx
  2852: 4885d2                   testq %rdx, %rdx
  2855: 0f8422000000             je 0xb4f
  2861: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2865: 81c0e6000000             addl $0xe6, %eax
  2871: 81f8ff000000             cmpl $0xff, %eax
  2877: 0f8206000000             jb 0xb49
  2883: 81c001ffffff             addl $0xffffff01, %eax
  2889: 0fb60406                 movzbl (%rsi, %rax), %eax
  2893: 32d8                     xorb %al, %bl
  2895: 498b4008                 movq 8(%r8), %rax
  2899: 420fb61418               movzbl (%rax, %r11), %edx
  2904: 4885d2                   testq %rdx, %rdx
  2907: 0f8422000000             je 0xb83
  2913: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2917: 81c0cd000000             addl $0xcd, %eax
  2923: 81f8ff000000             cmpl $0xff, %eax
  2929: 0f8206000000             jb 0xb7d
  2935: 81c001ffffff             addl $0xffffff01, %eax
  2941: 0fb60406                 movzbl (%rsi, %rax), %eax
  2945: 32d8                     xorb %al, %bl
  2947: 498b4010                 movq 0x10(%r8), %rax
  2951: 420fb61418               movzbl (%rax, %r11), %edx
  2956: 4885d2                   testq %rdx, %rdx
  2959: 0f8422000000             je 0xbb7
  2965: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2969: 81c0e7000000             addl $0xe7, %eax
  2975: 81f8ff000000             cmpl $0xff, %eax
  2981: 0f8206000000             jb 0xbb1
  2987: 81c001ffffff             addl $0xffffff01, %eax
  2993: 0fb60406                 movzbl (%rsi, %rax), %eax
  2997: 32d8                     xorb %al, %bl
  2999: 498b4018                 movq 0x18(%r8), %rax
  3003: 420fb61418               movzbl (%rax, %r11), %edx
  3008: 4885d2                   testq %rdx, %rdx
  3011: 0f8422000000             je 0xbeb
  3017: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3021: 81c0ce000000             addl $0xce, %eax
  3027: 81f8ff000000             cmpl $0xff, %eax
  3033: 0f8206000000             jb 0xbe5
  3039: 81c001ffffff             addl $0xffffff01, %eax
  3045: 0fb60406                 movzbl (%rsi, %rax), %eax
  3049: 32d8                     xorb %al, %bl
  3051: 498b4020                 movq 0x20(%r8), %rax
  3055: 420fb61418               movzbl (%rax, %r11), %edx
  3060: 4885d2                   testq %rdx, %rdx
  3063: 0f841f000000             je 0xc1c
  3069: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3073: 83c04b                   addl $0x4b, %eax
  3076: 81f8ff000000             cmpl $0xff, %eax
  3082: 0f8206000000             jb 0xc16
  3088: 81c001ffffff             addl $0xffffff01, %eax
  3094: 0fb60406                 movzbl (%rsi, %rax), %eax
  3098: 32d8                     xorb %al, %bl
  3100: 498b4028                 movq 0x28(%r8), %rax
  3104: 420fb61418               movzbl (%rax, %r11), %edx
  3109: 4885d2                   testq %rdx, %rdx
  3112: 0f841f000000             je 0xc4d
  3118: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3122: 83c032                   addl $0x32, %eax
  3125: 81f8ff000000             cmpl $0xff, %eax
  3131: 0f8206000000             jb 0xc47
  3137: 81c001ffffff             addl $0xffffff01, %eax
  3143: 0fb60406                 movzbl (%rsi, %rax), %eax
  3147: 32d8                     xorb %al, %bl
  3149: 498b4030                 movq 0x30(%r8), %rax
  3153: 420fb61418               movzbl (%rax, %r11), %edx
  3158: 4885d2                   testq %rdx, %rdx
  3161: 0f841f000000             je 0xc7e
  3167: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3171: 83c019                   addl $0x19, %eax
  3174: 81f8ff000000             cmpl $0xff, %eax
  3180: 0f8206000000             jb 0xc78
  3186: 81c001ffffff             addl $0xffffff01, %eax
  3192: 0fb60406                 movzbl (%rsi, %rax), %eax
  3196: 32d8                     xorb %al, %bl
  3198: 498b4038                 movq 0x38(%r8), %rax
  3202: 420fb61418               movzbl (%rax, %r11), %edx
  3207: 32da                     xorb %dl, %bl
  3209: 43881c1a                 movb %bl, (%r10, %r11)
  3213: 33db                     xorl %ebx, %ebx
  3215: 49ffc3                   incq %r11
  3218: 4c3bd9                   cmpq %rcx, %r11
  3221: 0f8281feffff             jb 0xb1c
  3227: 4883c408                 addq $8, %rsp
  3231: 415e                     popq %r14
  3233: 415d                     popq %r13
  3235: 415c                     popq %r12
  3237: 415f                     popq %r15
  3239: 5b                       popq %rbx
  3240: c3                       retq 
  3241: 0f0b                     ud2 
  3243: 0f0b                     ud2 
  3245: 0f0b                     ud2 
  3247: 0f0b                     ud2 
  3249: 0f0b                     ud2 
  3251: 0f0b                     ud2 
  3253: 0f0b                     ud2 
  3255: 0f0b                     ud2 
  3257: 0f0b                     ud2 
  3259: 0f0b                     ud2 
  3261: 0f0b                     ud2 
  3263: cc                       int3 
