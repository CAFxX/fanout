; config_t1_650 target=sse twin=baseline topo=rs(6,3)/enc op=encode U=1
; code_size=3008 twins=[] uops_per_byte=7.645833333333333
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da70b0000           leaq 0xba7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66450f6f3f               movdqa (%r15), %xmm15
    52: 488bc1                   movq %rcx, %rax
    55: 4883e830                 subq $0x30, %rax
    59: 0f82b8080000             jb 0x8f9
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66460f6f0c1e             movdqa (%rsi, %r11), %xmm9
   102: 66450f6fd1               movdqa %xmm9, %xmm10
   107: 66450fdbd7               pand %xmm15, %xmm10
   112: 66450f6fd9               movdqa %xmm9, %xmm11
   117: 66410f71d304             psrlw $4, %xmm11
   123: 66450fdbdf               pand %xmm15, %xmm11
   128: 66410f6fc1               movdqa %xmm9, %xmm0
   133: 66450f6f7740             movdqa 0x40(%r15), %xmm14
   139: 66450f6fe6               movdqa %xmm14, %xmm12
   144: 66450f3800e2             pshufb %xmm10, %xmm12
   150: 66450f6fb780000000       movdqa 0x80(%r15), %xmm14
   159: 66450f6fee               movdqa %xmm14, %xmm13
   164: 66450f3800eb             pshufb %xmm11, %xmm13
   170: 66410f6fdc               movdqa %xmm12, %xmm3
   175: 66410fefdd               pxor %xmm13, %xmm3
   180: 66450f6fb7c0000000       movdqa 0xc0(%r15), %xmm14
   189: 66450f6fe6               movdqa %xmm14, %xmm12
   194: 66450f3800e2             pshufb %xmm10, %xmm12
   200: 66450f6fb700010000       movdqa 0x100(%r15), %xmm14
   209: 66450f6fee               movdqa %xmm14, %xmm13
   214: 66450f3800eb             pshufb %xmm11, %xmm13
   220: 66410f6ff4               movdqa %xmm12, %xmm6
   225: 66410feff5               pxor %xmm13, %xmm6
   230: 66460f6f4c1e10           movdqa 0x10(%rsi, %r11), %xmm9
   237: 66450f6fd1               movdqa %xmm9, %xmm10
   242: 66450fdbd7               pand %xmm15, %xmm10
   247: 66450f6fd9               movdqa %xmm9, %xmm11
   252: 66410f71d304             psrlw $4, %xmm11
   258: 66450fdbdf               pand %xmm15, %xmm11
   263: 66410f6fc9               movdqa %xmm9, %xmm1
   268: 66450f6f7740             movdqa 0x40(%r15), %xmm14
   274: 66450f6fe6               movdqa %xmm14, %xmm12
   279: 66450f3800e2             pshufb %xmm10, %xmm12
   285: 66450f6fb780000000       movdqa 0x80(%r15), %xmm14
   294: 66450f6fee               movdqa %xmm14, %xmm13
   299: 66450f3800eb             pshufb %xmm11, %xmm13
   305: 66410f6fe4               movdqa %xmm12, %xmm4
   310: 66410fefe5               pxor %xmm13, %xmm4
   315: 66450f6fb7c0000000       movdqa 0xc0(%r15), %xmm14
   324: 66450f6fe6               movdqa %xmm14, %xmm12
   329: 66450f3800e2             pshufb %xmm10, %xmm12
   335: 66450f6fb700010000       movdqa 0x100(%r15), %xmm14
   344: 66450f6fee               movdqa %xmm14, %xmm13
   349: 66450f3800eb             pshufb %xmm11, %xmm13
   355: 66410f6ffc               movdqa %xmm12, %xmm7
   360: 66410feffd               pxor %xmm13, %xmm7
   365: 66460f6f4c1e20           movdqa 0x20(%rsi, %r11), %xmm9
   372: 66450f6fd1               movdqa %xmm9, %xmm10
   377: 66450fdbd7               pand %xmm15, %xmm10
   382: 66450f6fd9               movdqa %xmm9, %xmm11
   387: 66410f71d304             psrlw $4, %xmm11
   393: 66450fdbdf               pand %xmm15, %xmm11
   398: 66410f6fd1               movdqa %xmm9, %xmm2
   403: 66450f6f7740             movdqa 0x40(%r15), %xmm14
   409: 66450f6fe6               movdqa %xmm14, %xmm12
   414: 66450f3800e2             pshufb %xmm10, %xmm12
   420: 66450f6fb780000000       movdqa 0x80(%r15), %xmm14
   429: 66450f6fee               movdqa %xmm14, %xmm13
   434: 66450f3800eb             pshufb %xmm11, %xmm13
   440: 66410f6fec               movdqa %xmm12, %xmm5
   445: 66410fefed               pxor %xmm13, %xmm5
   450: 66450f6fb7c0000000       movdqa 0xc0(%r15), %xmm14
   459: 66450f6fe6               movdqa %xmm14, %xmm12
   464: 66450f3800e2             pshufb %xmm10, %xmm12
   470: 66450f6fb700010000       movdqa 0x100(%r15), %xmm14
   479: 66450f6fee               movdqa %xmm14, %xmm13
   484: 66450f3800eb             pshufb %xmm11, %xmm13
   490: 66450f6fc4               movdqa %xmm12, %xmm8
   495: 66450fefc5               pxor %xmm13, %xmm8
   500: 66460f6f0c1f             movdqa (%rdi, %r11), %xmm9
   506: 66450f6fd1               movdqa %xmm9, %xmm10
   511: 66450fdbd7               pand %xmm15, %xmm10
   516: 66450f6fd9               movdqa %xmm9, %xmm11
   521: 66410f71d304             psrlw $4, %xmm11
   527: 66450fdbdf               pand %xmm15, %xmm11
   532: 66410fefc1               pxor %xmm9, %xmm0
   537: 66450f6fb7c0000000       movdqa 0xc0(%r15), %xmm14
   546: 66450f6fe6               movdqa %xmm14, %xmm12
   551: 66450f3800e2             pshufb %xmm10, %xmm12
   557: 66450f6fb700010000       movdqa 0x100(%r15), %xmm14
   566: 66450f6fee               movdqa %xmm14, %xmm13
   571: 66450f3800eb             pshufb %xmm11, %xmm13
   577: 66410fefdc               pxor %xmm12, %xmm3
   582: 66410fefdd               pxor %xmm13, %xmm3
   587: 66450f6fb740010000       movdqa 0x140(%r15), %xmm14
   596: 66450f6fe6               movdqa %xmm14, %xmm12
   601: 66450f3800e2             pshufb %xmm10, %xmm12
   607: 66450f6fb780010000       movdqa 0x180(%r15), %xmm14
   616: 66450f6fee               movdqa %xmm14, %xmm13
   621: 66450f3800eb             pshufb %xmm11, %xmm13
   627: 66410feff4               pxor %xmm12, %xmm6
   632: 66410feff5               pxor %xmm13, %xmm6
   637: 66460f6f4c1f10           movdqa 0x10(%rdi, %r11), %xmm9
   644: 66450f6fd1               movdqa %xmm9, %xmm10
   649: 66450fdbd7               pand %xmm15, %xmm10
   654: 66450f6fd9               movdqa %xmm9, %xmm11
   659: 66410f71d304             psrlw $4, %xmm11
   665: 66450fdbdf               pand %xmm15, %xmm11
   670: 66410fefc9               pxor %xmm9, %xmm1
   675: 66450f6fb7c0000000       movdqa 0xc0(%r15), %xmm14
   684: 66450f6fe6               movdqa %xmm14, %xmm12
   689: 66450f3800e2             pshufb %xmm10, %xmm12
   695: 66450f6fb700010000       movdqa 0x100(%r15), %xmm14
   704: 66450f6fee               movdqa %xmm14, %xmm13
   709: 66450f3800eb             pshufb %xmm11, %xmm13
   715: 66410fefe4               pxor %xmm12, %xmm4
   720: 66410fefe5               pxor %xmm13, %xmm4
   725: 66450f6fb740010000       movdqa 0x140(%r15), %xmm14
   734: 66450f6fe6               movdqa %xmm14, %xmm12
   739: 66450f3800e2             pshufb %xmm10, %xmm12
   745: 66450f6fb780010000       movdqa 0x180(%r15), %xmm14
   754: 66450f6fee               movdqa %xmm14, %xmm13
   759: 66450f3800eb             pshufb %xmm11, %xmm13
   765: 66410feffc               pxor %xmm12, %xmm7
   770: 66410feffd               pxor %xmm13, %xmm7
   775: 66460f6f4c1f20           movdqa 0x20(%rdi, %r11), %xmm9
   782: 66450f6fd1               movdqa %xmm9, %xmm10
   787: 66450fdbd7               pand %xmm15, %xmm10
   792: 66450f6fd9               movdqa %xmm9, %xmm11
   797: 66410f71d304             psrlw $4, %xmm11
   803: 66450fdbdf               pand %xmm15, %xmm11
   808: 66410fefd1               pxor %xmm9, %xmm2
   813: 66450f6fb7c0000000       movdqa 0xc0(%r15), %xmm14
   822: 66450f6fe6               movdqa %xmm14, %xmm12
   827: 66450f3800e2             pshufb %xmm10, %xmm12
   833: 66450f6fb700010000       movdqa 0x100(%r15), %xmm14
   842: 66450f6fee               movdqa %xmm14, %xmm13
   847: 66450f3800eb             pshufb %xmm11, %xmm13
   853: 66410fefec               pxor %xmm12, %xmm5
   858: 66410fefed               pxor %xmm13, %xmm5
   863: 66450f6fb740010000       movdqa 0x140(%r15), %xmm14
   872: 66450f6fe6               movdqa %xmm14, %xmm12
   877: 66450f3800e2             pshufb %xmm10, %xmm12
   883: 66450f6fb780010000       movdqa 0x180(%r15), %xmm14
   892: 66450f6fee               movdqa %xmm14, %xmm13
   897: 66450f3800eb             pshufb %xmm11, %xmm13
   903: 66450fefc4               pxor %xmm12, %xmm8
   908: 66450fefc5               pxor %xmm13, %xmm8
   913: 66470f6f0c1c             movdqa (%r12, %r11), %xmm9
   919: 66450f6fd1               movdqa %xmm9, %xmm10
   924: 66450fdbd7               pand %xmm15, %xmm10
   929: 66450f6fd9               movdqa %xmm9, %xmm11
   934: 66410f71d304             psrlw $4, %xmm11
   940: 66450fdbdf               pand %xmm15, %xmm11
   945: 66410fefc1               pxor %xmm9, %xmm0
   950: 66450f6fb740010000       movdqa 0x140(%r15), %xmm14
   959: 66450f6fe6               movdqa %xmm14, %xmm12
   964: 66450f3800e2             pshufb %xmm10, %xmm12
   970: 66450f6fb780010000       movdqa 0x180(%r15), %xmm14
   979: 66450f6fee               movdqa %xmm14, %xmm13
   984: 66450f3800eb             pshufb %xmm11, %xmm13
   990: 66410fefdc               pxor %xmm12, %xmm3
   995: 66410fefdd               pxor %xmm13, %xmm3
  1000: 66450f6fb7c0010000       movdqa 0x1c0(%r15), %xmm14
  1009: 66450f6fe6               movdqa %xmm14, %xmm12
  1014: 66450f3800e2             pshufb %xmm10, %xmm12
  1020: 66450f6fb700020000       movdqa 0x200(%r15), %xmm14
  1029: 66450f6fee               movdqa %xmm14, %xmm13
  1034: 66450f3800eb             pshufb %xmm11, %xmm13
  1040: 66410feff4               pxor %xmm12, %xmm6
  1045: 66410feff5               pxor %xmm13, %xmm6
  1050: 66470f6f4c1c10           movdqa 0x10(%r12, %r11), %xmm9
  1057: 66450f6fd1               movdqa %xmm9, %xmm10
  1062: 66450fdbd7               pand %xmm15, %xmm10
  1067: 66450f6fd9               movdqa %xmm9, %xmm11
  1072: 66410f71d304             psrlw $4, %xmm11
  1078: 66450fdbdf               pand %xmm15, %xmm11
  1083: 66410fefc9               pxor %xmm9, %xmm1
  1088: 66450f6fb740010000       movdqa 0x140(%r15), %xmm14
  1097: 66450f6fe6               movdqa %xmm14, %xmm12
  1102: 66450f3800e2             pshufb %xmm10, %xmm12
  1108: 66450f6fb780010000       movdqa 0x180(%r15), %xmm14
  1117: 66450f6fee               movdqa %xmm14, %xmm13
  1122: 66450f3800eb             pshufb %xmm11, %xmm13
  1128: 66410fefe4               pxor %xmm12, %xmm4
  1133: 66410fefe5               pxor %xmm13, %xmm4
  1138: 66450f6fb7c0010000       movdqa 0x1c0(%r15), %xmm14
  1147: 66450f6fe6               movdqa %xmm14, %xmm12
  1152: 66450f3800e2             pshufb %xmm10, %xmm12
  1158: 66450f6fb700020000       movdqa 0x200(%r15), %xmm14
  1167: 66450f6fee               movdqa %xmm14, %xmm13
  1172: 66450f3800eb             pshufb %xmm11, %xmm13
  1178: 66410feffc               pxor %xmm12, %xmm7
  1183: 66410feffd               pxor %xmm13, %xmm7
  1188: 66470f6f4c1c20           movdqa 0x20(%r12, %r11), %xmm9
  1195: 66450f6fd1               movdqa %xmm9, %xmm10
  1200: 66450fdbd7               pand %xmm15, %xmm10
  1205: 66450f6fd9               movdqa %xmm9, %xmm11
  1210: 66410f71d304             psrlw $4, %xmm11
  1216: 66450fdbdf               pand %xmm15, %xmm11
  1221: 66410fefd1               pxor %xmm9, %xmm2
  1226: 66450f6fb740010000       movdqa 0x140(%r15), %xmm14
  1235: 66450f6fe6               movdqa %xmm14, %xmm12
  1240: 66450f3800e2             pshufb %xmm10, %xmm12
  1246: 66450f6fb780010000       movdqa 0x180(%r15), %xmm14
  1255: 66450f6fee               movdqa %xmm14, %xmm13
  1260: 66450f3800eb             pshufb %xmm11, %xmm13
  1266: 66410fefec               pxor %xmm12, %xmm5
  1271: 66410fefed               pxor %xmm13, %xmm5
  1276: 66450f6fb7c0010000       movdqa 0x1c0(%r15), %xmm14
  1285: 66450f6fe6               movdqa %xmm14, %xmm12
  1290: 66450f3800e2             pshufb %xmm10, %xmm12
  1296: 66450f6fb700020000       movdqa 0x200(%r15), %xmm14
  1305: 66450f6fee               movdqa %xmm14, %xmm13
  1310: 66450f3800eb             pshufb %xmm11, %xmm13
  1316: 66450fefc4               pxor %xmm12, %xmm8
  1321: 66450fefc5               pxor %xmm13, %xmm8
  1326: 66470f6f4c1d00           movdqa (%r13, %r11), %xmm9
  1333: 66450f6fd1               movdqa %xmm9, %xmm10
  1338: 66450fdbd7               pand %xmm15, %xmm10
  1343: 66450f6fd9               movdqa %xmm9, %xmm11
  1348: 66410f71d304             psrlw $4, %xmm11
  1354: 66450fdbdf               pand %xmm15, %xmm11
  1359: 66410fefc1               pxor %xmm9, %xmm0
  1364: 66450f6fb740020000       movdqa 0x240(%r15), %xmm14
  1373: 66450f6fe6               movdqa %xmm14, %xmm12
  1378: 66450f3800e2             pshufb %xmm10, %xmm12
  1384: 66450f6fb780020000       movdqa 0x280(%r15), %xmm14
  1393: 66450f6fee               movdqa %xmm14, %xmm13
  1398: 66450f3800eb             pshufb %xmm11, %xmm13
  1404: 66410fefdc               pxor %xmm12, %xmm3
  1409: 66410fefdd               pxor %xmm13, %xmm3
  1414: 66450f6f7740             movdqa 0x40(%r15), %xmm14
  1420: 66450f6fe6               movdqa %xmm14, %xmm12
  1425: 66450f3800e2             pshufb %xmm10, %xmm12
  1431: 66450f6fb780000000       movdqa 0x80(%r15), %xmm14
  1440: 66450f6fee               movdqa %xmm14, %xmm13
  1445: 66450f3800eb             pshufb %xmm11, %xmm13
  1451: 66410feff4               pxor %xmm12, %xmm6
  1456: 66410feff5               pxor %xmm13, %xmm6
  1461: 66470f6f4c1d10           movdqa 0x10(%r13, %r11), %xmm9
  1468: 66450f6fd1               movdqa %xmm9, %xmm10
  1473: 66450fdbd7               pand %xmm15, %xmm10
  1478: 66450f6fd9               movdqa %xmm9, %xmm11
  1483: 66410f71d304             psrlw $4, %xmm11
  1489: 66450fdbdf               pand %xmm15, %xmm11
  1494: 66410fefc9               pxor %xmm9, %xmm1
  1499: 66450f6fb740020000       movdqa 0x240(%r15), %xmm14
  1508: 66450f6fe6               movdqa %xmm14, %xmm12
  1513: 66450f3800e2             pshufb %xmm10, %xmm12
  1519: 66450f6fb780020000       movdqa 0x280(%r15), %xmm14
  1528: 66450f6fee               movdqa %xmm14, %xmm13
  1533: 66450f3800eb             pshufb %xmm11, %xmm13
  1539: 66410fefe4               pxor %xmm12, %xmm4
  1544: 66410fefe5               pxor %xmm13, %xmm4
  1549: 66450f6f7740             movdqa 0x40(%r15), %xmm14
  1555: 66450f6fe6               movdqa %xmm14, %xmm12
  1560: 66450f3800e2             pshufb %xmm10, %xmm12
  1566: 66450f6fb780000000       movdqa 0x80(%r15), %xmm14
  1575: 66450f6fee               movdqa %xmm14, %xmm13
  1580: 66450f3800eb             pshufb %xmm11, %xmm13
  1586: 66410feffc               pxor %xmm12, %xmm7
  1591: 66410feffd               pxor %xmm13, %xmm7
  1596: 66470f6f4c1d20           movdqa 0x20(%r13, %r11), %xmm9
  1603: 66450f6fd1               movdqa %xmm9, %xmm10
  1608: 66450fdbd7               pand %xmm15, %xmm10
  1613: 66450f6fd9               movdqa %xmm9, %xmm11
  1618: 66410f71d304             psrlw $4, %xmm11
  1624: 66450fdbdf               pand %xmm15, %xmm11
  1629: 66410fefd1               pxor %xmm9, %xmm2
  1634: 66450f6fb740020000       movdqa 0x240(%r15), %xmm14
  1643: 66450f6fe6               movdqa %xmm14, %xmm12
  1648: 66450f3800e2             pshufb %xmm10, %xmm12
  1654: 66450f6fb780020000       movdqa 0x280(%r15), %xmm14
  1663: 66450f6fee               movdqa %xmm14, %xmm13
  1668: 66450f3800eb             pshufb %xmm11, %xmm13
  1674: 66410fefec               pxor %xmm12, %xmm5
  1679: 66410fefed               pxor %xmm13, %xmm5
  1684: 66450f6f7740             movdqa 0x40(%r15), %xmm14
  1690: 66450f6fe6               movdqa %xmm14, %xmm12
  1695: 66450f3800e2             pshufb %xmm10, %xmm12
  1701: 66450f6fb780000000       movdqa 0x80(%r15), %xmm14
  1710: 66450f6fee               movdqa %xmm14, %xmm13
  1715: 66450f3800eb             pshufb %xmm11, %xmm13
  1721: 66450fefc4               pxor %xmm12, %xmm8
  1726: 66450fefc5               pxor %xmm13, %xmm8
  1731: 66470f6f0c1e             movdqa (%r14, %r11), %xmm9
  1737: 66450f6fd1               movdqa %xmm9, %xmm10
  1742: 66450fdbd7               pand %xmm15, %xmm10
  1747: 66450f6fd9               movdqa %xmm9, %xmm11
  1752: 66410f71d304             psrlw $4, %xmm11
  1758: 66450fdbdf               pand %xmm15, %xmm11
  1763: 66410fefc1               pxor %xmm9, %xmm0
  1768: 66450f6fb7c0010000       movdqa 0x1c0(%r15), %xmm14
  1777: 66450f6fe6               movdqa %xmm14, %xmm12
  1782: 66450f3800e2             pshufb %xmm10, %xmm12
  1788: 66450f6fb700020000       movdqa 0x200(%r15), %xmm14
  1797: 66450f6fee               movdqa %xmm14, %xmm13
  1802: 66450f3800eb             pshufb %xmm11, %xmm13
  1808: 66410fefdc               pxor %xmm12, %xmm3
  1813: 66410fefdd               pxor %xmm13, %xmm3
  1818: 66450f6fb740020000       movdqa 0x240(%r15), %xmm14
  1827: 66450f6fe6               movdqa %xmm14, %xmm12
  1832: 66450f3800e2             pshufb %xmm10, %xmm12
  1838: 66450f6fb780020000       movdqa 0x280(%r15), %xmm14
  1847: 66450f6fee               movdqa %xmm14, %xmm13
  1852: 66450f3800eb             pshufb %xmm11, %xmm13
  1858: 66410feff4               pxor %xmm12, %xmm6
  1863: 66410feff5               pxor %xmm13, %xmm6
  1868: 66470f6f4c1e10           movdqa 0x10(%r14, %r11), %xmm9
  1875: 66450f6fd1               movdqa %xmm9, %xmm10
  1880: 66450fdbd7               pand %xmm15, %xmm10
  1885: 66450f6fd9               movdqa %xmm9, %xmm11
  1890: 66410f71d304             psrlw $4, %xmm11
  1896: 66450fdbdf               pand %xmm15, %xmm11
  1901: 66410fefc9               pxor %xmm9, %xmm1
  1906: 66450f6fb7c0010000       movdqa 0x1c0(%r15), %xmm14
  1915: 66450f6fe6               movdqa %xmm14, %xmm12
  1920: 66450f3800e2             pshufb %xmm10, %xmm12
  1926: 66450f6fb700020000       movdqa 0x200(%r15), %xmm14
  1935: 66450f6fee               movdqa %xmm14, %xmm13
  1940: 66450f3800eb             pshufb %xmm11, %xmm13
  1946: 66410fefe4               pxor %xmm12, %xmm4
  1951: 66410fefe5               pxor %xmm13, %xmm4
  1956: 66450f6fb740020000       movdqa 0x240(%r15), %xmm14
  1965: 66450f6fe6               movdqa %xmm14, %xmm12
  1970: 66450f3800e2             pshufb %xmm10, %xmm12
  1976: 66450f6fb780020000       movdqa 0x280(%r15), %xmm14
  1985: 66450f6fee               movdqa %xmm14, %xmm13
  1990: 66450f3800eb             pshufb %xmm11, %xmm13
  1996: 66410feffc               pxor %xmm12, %xmm7
  2001: 66410feffd               pxor %xmm13, %xmm7
  2006: 66470f6f4c1e20           movdqa 0x20(%r14, %r11), %xmm9
  2013: 66450f6fd1               movdqa %xmm9, %xmm10
  2018: 66450fdbd7               pand %xmm15, %xmm10
  2023: 66450f6fd9               movdqa %xmm9, %xmm11
  2028: 66410f71d304             psrlw $4, %xmm11
  2034: 66450fdbdf               pand %xmm15, %xmm11
  2039: 66410fefd1               pxor %xmm9, %xmm2
  2044: 66450f6fb7c0010000       movdqa 0x1c0(%r15), %xmm14
  2053: 66450f6fe6               movdqa %xmm14, %xmm12
  2058: 66450f3800e2             pshufb %xmm10, %xmm12
  2064: 66450f6fb700020000       movdqa 0x200(%r15), %xmm14
  2073: 66450f6fee               movdqa %xmm14, %xmm13
  2078: 66450f3800eb             pshufb %xmm11, %xmm13
  2084: 66410fefec               pxor %xmm12, %xmm5
  2089: 66410fefed               pxor %xmm13, %xmm5
  2094: 66450f6fb740020000       movdqa 0x240(%r15), %xmm14
  2103: 66450f6fe6               movdqa %xmm14, %xmm12
  2108: 66450f3800e2             pshufb %xmm10, %xmm12
  2114: 66450f6fb780020000       movdqa 0x280(%r15), %xmm14
  2123: 66450f6fee               movdqa %xmm14, %xmm13
  2128: 66450f3800eb             pshufb %xmm11, %xmm13
  2134: 66450fefc4               pxor %xmm12, %xmm8
  2139: 66450fefc5               pxor %xmm13, %xmm8
  2144: 4d8b5028                 movq 0x28(%r8), %r10
  2148: 66470f6f0c1a             movdqa (%r10, %r11), %xmm9
  2154: 66410fefc1               pxor %xmm9, %xmm0
  2159: 66410fefd9               pxor %xmm9, %xmm3
  2164: 66410feff1               pxor %xmm9, %xmm6
  2169: 66470f6f4c1a10           movdqa 0x10(%r10, %r11), %xmm9
  2176: 66410fefc9               pxor %xmm9, %xmm1
  2181: 66410fefe1               pxor %xmm9, %xmm4
  2186: 66410feff9               pxor %xmm9, %xmm7
  2191: 66470f6f4c1a20           movdqa 0x20(%r10, %r11), %xmm9
  2198: 66410fefd1               pxor %xmm9, %xmm2
  2203: 66410fefe9               pxor %xmm9, %xmm5
  2208: 66450fefc1               pxor %xmm9, %xmm8
  2213: 4d8b11                   movq (%r9), %r10
  2216: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
  2222: 66430f7f4c1a10           movdqa %xmm1, 0x10(%r10, %r11)
  2229: 66430f7f541a20           movdqa %xmm2, 0x20(%r10, %r11)
  2236: 4d8b5108                 movq 8(%r9), %r10
  2240: 66430f7f1c1a             movdqa %xmm3, (%r10, %r11)
  2246: 66430f7f641a10           movdqa %xmm4, 0x10(%r10, %r11)
  2253: 66430f7f6c1a20           movdqa %xmm5, 0x20(%r10, %r11)
  2260: 4d8b5110                 movq 0x10(%r9), %r10
  2264: 66430f7f341a             movdqa %xmm6, (%r10, %r11)
  2270: 66430f7f7c1a10           movdqa %xmm7, 0x10(%r10, %r11)
  2277: 66470f7f441a20           movdqa %xmm8, 0x20(%r10, %r11)
  2284: 4983c330                 addq $0x30, %r11
  2288: 4c3bd8                   cmpq %rax, %r11
  2291: 0f8667f7ffff             jbe 0x60
  2297: 4c3bd9                   cmpq %rcx, %r11
  2300: 0f83b2020000             jae 0xbb4
  2306: 4153                     pushq %r11
  2308: 498dbfc0020000           leaq 0x2c0(%r15), %rdi
  2315: 498db7c0030000           leaq 0x3c0(%r15), %rsi
  2322: 4d8b11                   movq (%r9), %r10
  2325: 4c8b1c24                 movq (%rsp), %r11
  2329: 33db                     xorl %ebx, %ebx
  2331: 498b00                   movq (%r8), %rax
  2334: 420fb61418               movzbl (%rax, %r11), %edx
  2339: 32da                     xorb %dl, %bl
  2341: 498b4008                 movq 8(%r8), %rax
  2345: 420fb61418               movzbl (%rax, %r11), %edx
  2350: 32da                     xorb %dl, %bl
  2352: 498b4010                 movq 0x10(%r8), %rax
  2356: 420fb61418               movzbl (%rax, %r11), %edx
  2361: 32da                     xorb %dl, %bl
  2363: 498b4018                 movq 0x18(%r8), %rax
  2367: 420fb61418               movzbl (%rax, %r11), %edx
  2372: 32da                     xorb %dl, %bl
  2374: 498b4020                 movq 0x20(%r8), %rax
  2378: 420fb61418               movzbl (%rax, %r11), %edx
  2383: 32da                     xorb %dl, %bl
  2385: 498b4028                 movq 0x28(%r8), %rax
  2389: 420fb61418               movzbl (%rax, %r11), %edx
  2394: 32da                     xorb %dl, %bl
  2396: 43881c1a                 movb %bl, (%r10, %r11)
  2400: 33db                     xorl %ebx, %ebx
  2402: 49ffc3                   incq %r11
  2405: 4c3bd9                   cmpq %rcx, %r11
  2408: 0f82adffffff             jb 0x91b
  2414: 4d8b5108                 movq 8(%r9), %r10
  2418: 4c8b1c24                 movq (%rsp), %r11
  2422: 33db                     xorl %ebx, %ebx
  2424: 498b00                   movq (%r8), %rax
  2427: 420fb61418               movzbl (%rax, %r11), %edx
  2432: 4885d2                   testq %rdx, %rdx
  2435: 0f8422000000             je 0x9ab
  2441: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2445: 81c0e6000000             addl $0xe6, %eax
  2451: 81f8ff000000             cmpl $0xff, %eax
  2457: 0f8206000000             jb 0x9a5
  2463: 81c001ffffff             addl $0xffffff01, %eax
  2469: 0fb60406                 movzbl (%rsi, %rax), %eax
  2473: 32d8                     xorb %al, %bl
  2475: 498b4008                 movq 8(%r8), %rax
  2479: 420fb61418               movzbl (%rax, %r11), %edx
  2484: 4885d2                   testq %rdx, %rdx
  2487: 0f8422000000             je 0x9df
  2493: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2497: 81c0cd000000             addl $0xcd, %eax
  2503: 81f8ff000000             cmpl $0xff, %eax
  2509: 0f8206000000             jb 0x9d9
  2515: 81c001ffffff             addl $0xffffff01, %eax
  2521: 0fb60406                 movzbl (%rsi, %rax), %eax
  2525: 32d8                     xorb %al, %bl
  2527: 498b4010                 movq 0x10(%r8), %rax
  2531: 420fb61418               movzbl (%rax, %r11), %edx
  2536: 4885d2                   testq %rdx, %rdx
  2539: 0f841f000000             je 0xa10
  2545: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2549: 83c04b                   addl $0x4b, %eax
  2552: 81f8ff000000             cmpl $0xff, %eax
  2558: 0f8206000000             jb 0xa0a
  2564: 81c001ffffff             addl $0xffffff01, %eax
  2570: 0fb60406                 movzbl (%rsi, %rax), %eax
  2574: 32d8                     xorb %al, %bl
  2576: 498b4018                 movq 0x18(%r8), %rax
  2580: 420fb61418               movzbl (%rax, %r11), %edx
  2585: 4885d2                   testq %rdx, %rdx
  2588: 0f841f000000             je 0xa41
  2594: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2598: 83c032                   addl $0x32, %eax
  2601: 81f8ff000000             cmpl $0xff, %eax
  2607: 0f8206000000             jb 0xa3b
  2613: 81c001ffffff             addl $0xffffff01, %eax
  2619: 0fb60406                 movzbl (%rsi, %rax), %eax
  2623: 32d8                     xorb %al, %bl
  2625: 498b4020                 movq 0x20(%r8), %rax
  2629: 420fb61418               movzbl (%rax, %r11), %edx
  2634: 4885d2                   testq %rdx, %rdx
  2637: 0f841f000000             je 0xa72
  2643: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2647: 83c019                   addl $0x19, %eax
  2650: 81f8ff000000             cmpl $0xff, %eax
  2656: 0f8206000000             jb 0xa6c
  2662: 81c001ffffff             addl $0xffffff01, %eax
  2668: 0fb60406                 movzbl (%rsi, %rax), %eax
  2672: 32d8                     xorb %al, %bl
  2674: 498b4028                 movq 0x28(%r8), %rax
  2678: 420fb61418               movzbl (%rax, %r11), %edx
  2683: 32da                     xorb %dl, %bl
  2685: 43881c1a                 movb %bl, (%r10, %r11)
  2689: 33db                     xorl %ebx, %ebx
  2691: 49ffc3                   incq %r11
  2694: 4c3bd9                   cmpq %rcx, %r11
  2697: 0f82e9feffff             jb 0x978
  2703: 4d8b5110                 movq 0x10(%r9), %r10
  2707: 4c8b1c24                 movq (%rsp), %r11
  2711: 33db                     xorl %ebx, %ebx
  2713: 498b00                   movq (%r8), %rax
  2716: 420fb61418               movzbl (%rax, %r11), %edx
  2721: 4885d2                   testq %rdx, %rdx
  2724: 0f8422000000             je 0xacc
  2730: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2734: 81c0cd000000             addl $0xcd, %eax
  2740: 81f8ff000000             cmpl $0xff, %eax
  2746: 0f8206000000             jb 0xac6
  2752: 81c001ffffff             addl $0xffffff01, %eax
  2758: 0fb60406                 movzbl (%rsi, %rax), %eax
  2762: 32d8                     xorb %al, %bl
  2764: 498b4008                 movq 8(%r8), %rax
  2768: 420fb61418               movzbl (%rax, %r11), %edx
  2773: 4885d2                   testq %rdx, %rdx
  2776: 0f841f000000             je 0xafd
  2782: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2786: 83c04b                   addl $0x4b, %eax
  2789: 81f8ff000000             cmpl $0xff, %eax
  2795: 0f8206000000             jb 0xaf7
  2801: 81c001ffffff             addl $0xffffff01, %eax
  2807: 0fb60406                 movzbl (%rsi, %rax), %eax
  2811: 32d8                     xorb %al, %bl
  2813: 498b4010                 movq 0x10(%r8), %rax
  2817: 420fb61418               movzbl (%rax, %r11), %edx
  2822: 4885d2                   testq %rdx, %rdx
  2825: 0f841f000000             je 0xb2e
  2831: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2835: 83c019                   addl $0x19, %eax
  2838: 81f8ff000000             cmpl $0xff, %eax
  2844: 0f8206000000             jb 0xb28
  2850: 81c001ffffff             addl $0xffffff01, %eax
  2856: 0fb60406                 movzbl (%rsi, %rax), %eax
  2860: 32d8                     xorb %al, %bl
  2862: 498b4018                 movq 0x18(%r8), %rax
  2866: 420fb61418               movzbl (%rax, %r11), %edx
  2871: 4885d2                   testq %rdx, %rdx
  2874: 0f8422000000             je 0xb62
  2880: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2884: 81c0e6000000             addl $0xe6, %eax
  2890: 81f8ff000000             cmpl $0xff, %eax
  2896: 0f8206000000             jb 0xb5c
  2902: 81c001ffffff             addl $0xffffff01, %eax
  2908: 0fb60406                 movzbl (%rsi, %rax), %eax
  2912: 32d8                     xorb %al, %bl
  2914: 498b4020                 movq 0x20(%r8), %rax
  2918: 420fb61418               movzbl (%rax, %r11), %edx
  2923: 4885d2                   testq %rdx, %rdx
  2926: 0f841f000000             je 0xb93
  2932: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2936: 83c032                   addl $0x32, %eax
  2939: 81f8ff000000             cmpl $0xff, %eax
  2945: 0f8206000000             jb 0xb8d
  2951: 81c001ffffff             addl $0xffffff01, %eax
  2957: 0fb60406                 movzbl (%rsi, %rax), %eax
  2961: 32d8                     xorb %al, %bl
  2963: 498b4028                 movq 0x28(%r8), %rax
  2967: 420fb61418               movzbl (%rax, %r11), %edx
  2972: 32da                     xorb %dl, %bl
  2974: 43881c1a                 movb %bl, (%r10, %r11)
  2978: 33db                     xorl %ebx, %ebx
  2980: 49ffc3                   incq %r11
  2983: 4c3bd9                   cmpq %rcx, %r11
  2986: 0f82e9feffff             jb 0xa99
  2992: 4883c408                 addq $8, %rsp
  2996: 415e                     popq %r14
  2998: 415d                     popq %r13
  3000: 415c                     popq %r12
  3002: 415f                     popq %r15
  3004: 5b                       popq %rbx
  3005: c3                       retq 
  3006: 0f0b                     ud2 
