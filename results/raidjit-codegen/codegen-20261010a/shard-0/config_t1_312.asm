; config_t1_312 target=avx2 twin=baseline topo=raid6(8)/enc op=encode U=1
; code_size=2624 twins=[] uops_per_byte=2.5546875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d270a0000           leaq 0xa27(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: c4417d6f37               vmovdqa (%r15), %ymm14
    52: 488bc1                   movq %rcx, %rax
    55: 4881e880000000           subq $0x80, %rax
    62: 0f82ca070000             jb 0x80e
    68: 666666666666660f1f840000000000 nopw (%rax, %rax)
    83: 66666666660f1f840000000000 nopw (%rax, %rax)
    96: c4217d6f041e             vmovdqa (%rsi, %r11), %ymm8
   102: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   107: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   113: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   118: c4c17d6fc0               vmovdqa %ymm8, %ymm0
   123: c4417d6f6f40             vmovdqa 0x40(%r15), %ymm13
   129: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   134: c4417d6faf80000000       vmovdqa 0x80(%r15), %ymm13
   143: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   148: c4c125efe4               vpxor %ymm12, %ymm11, %ymm4
   153: c4217d6f441e20           vmovdqa 0x20(%rsi, %r11), %ymm8
   160: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   165: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   171: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   176: c4c17d6fc8               vmovdqa %ymm8, %ymm1
   181: c4417d6f6f40             vmovdqa 0x40(%r15), %ymm13
   187: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   192: c4417d6faf80000000       vmovdqa 0x80(%r15), %ymm13
   201: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   206: c4c125efec               vpxor %ymm12, %ymm11, %ymm5
   211: c4217d6f441e40           vmovdqa 0x40(%rsi, %r11), %ymm8
   218: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   223: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   229: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   234: c4c17d6fd0               vmovdqa %ymm8, %ymm2
   239: c4417d6f6f40             vmovdqa 0x40(%r15), %ymm13
   245: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   250: c4417d6faf80000000       vmovdqa 0x80(%r15), %ymm13
   259: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   264: c4c125eff4               vpxor %ymm12, %ymm11, %ymm6
   269: c4217d6f441e60           vmovdqa 0x60(%rsi, %r11), %ymm8
   276: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   281: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   287: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   292: c4c17d6fd8               vmovdqa %ymm8, %ymm3
   297: c4417d6f6f40             vmovdqa 0x40(%r15), %ymm13
   303: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   308: c4417d6faf80000000       vmovdqa 0x80(%r15), %ymm13
   317: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   322: c4c125effc               vpxor %ymm12, %ymm11, %ymm7
   327: c4217d6f041f             vmovdqa (%rdi, %r11), %ymm8
   333: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   338: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   344: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   349: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   354: c4417d6fafc0000000       vmovdqa 0xc0(%r15), %ymm13
   363: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   368: c4417d6faf00010000       vmovdqa 0x100(%r15), %ymm13
   377: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   382: c4c15defe3               vpxor %ymm11, %ymm4, %ymm4
   387: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   392: c4217d6f441f20           vmovdqa 0x20(%rdi, %r11), %ymm8
   399: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   404: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   410: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   415: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   420: c4417d6fafc0000000       vmovdqa 0xc0(%r15), %ymm13
   429: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   434: c4417d6faf00010000       vmovdqa 0x100(%r15), %ymm13
   443: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   448: c4c155efeb               vpxor %ymm11, %ymm5, %ymm5
   453: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   458: c4217d6f441f40           vmovdqa 0x40(%rdi, %r11), %ymm8
   465: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   470: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   476: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   481: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   486: c4417d6fafc0000000       vmovdqa 0xc0(%r15), %ymm13
   495: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   500: c4417d6faf00010000       vmovdqa 0x100(%r15), %ymm13
   509: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   514: c4c14deff3               vpxor %ymm11, %ymm6, %ymm6
   519: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   524: c4217d6f441f60           vmovdqa 0x60(%rdi, %r11), %ymm8
   531: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   536: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   542: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   547: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
   552: c4417d6fafc0000000       vmovdqa 0xc0(%r15), %ymm13
   561: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   566: c4417d6faf00010000       vmovdqa 0x100(%r15), %ymm13
   575: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   580: c4c145effb               vpxor %ymm11, %ymm7, %ymm7
   585: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   590: c4017d6f041c             vmovdqa (%r12, %r11), %ymm8
   596: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   601: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   607: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   612: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   617: c4417d6faf40010000       vmovdqa 0x140(%r15), %ymm13
   626: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   631: c4417d6faf80010000       vmovdqa 0x180(%r15), %ymm13
   640: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   645: c4c15defe3               vpxor %ymm11, %ymm4, %ymm4
   650: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   655: c4017d6f441c20           vmovdqa 0x20(%r12, %r11), %ymm8
   662: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   667: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   673: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   678: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   683: c4417d6faf40010000       vmovdqa 0x140(%r15), %ymm13
   692: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   697: c4417d6faf80010000       vmovdqa 0x180(%r15), %ymm13
   706: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   711: c4c155efeb               vpxor %ymm11, %ymm5, %ymm5
   716: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   721: c4017d6f441c40           vmovdqa 0x40(%r12, %r11), %ymm8
   728: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   733: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   739: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   744: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   749: c4417d6faf40010000       vmovdqa 0x140(%r15), %ymm13
   758: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   763: c4417d6faf80010000       vmovdqa 0x180(%r15), %ymm13
   772: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   777: c4c14deff3               vpxor %ymm11, %ymm6, %ymm6
   782: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   787: c4017d6f441c60           vmovdqa 0x60(%r12, %r11), %ymm8
   794: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   799: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   805: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   810: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
   815: c4417d6faf40010000       vmovdqa 0x140(%r15), %ymm13
   824: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   829: c4417d6faf80010000       vmovdqa 0x180(%r15), %ymm13
   838: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   843: c4c145effb               vpxor %ymm11, %ymm7, %ymm7
   848: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   853: c4017d6f441d00           vmovdqa (%r13, %r11), %ymm8
   860: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   865: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   871: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   876: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   881: c4417d6fafc0010000       vmovdqa 0x1c0(%r15), %ymm13
   890: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   895: c4417d6faf00020000       vmovdqa 0x200(%r15), %ymm13
   904: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   909: c4c15defe3               vpxor %ymm11, %ymm4, %ymm4
   914: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   919: c4017d6f441d20           vmovdqa 0x20(%r13, %r11), %ymm8
   926: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   931: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   937: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   942: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   947: c4417d6fafc0010000       vmovdqa 0x1c0(%r15), %ymm13
   956: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   961: c4417d6faf00020000       vmovdqa 0x200(%r15), %ymm13
   970: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   975: c4c155efeb               vpxor %ymm11, %ymm5, %ymm5
   980: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   985: c4017d6f441d40           vmovdqa 0x40(%r13, %r11), %ymm8
   992: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   997: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1003: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1008: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  1013: c4417d6fafc0010000       vmovdqa 0x1c0(%r15), %ymm13
  1022: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1027: c4417d6faf00020000       vmovdqa 0x200(%r15), %ymm13
  1036: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1041: c4c14deff3               vpxor %ymm11, %ymm6, %ymm6
  1046: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
  1051: c4017d6f441d60           vmovdqa 0x60(%r13, %r11), %ymm8
  1058: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1063: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1069: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1074: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  1079: c4417d6fafc0010000       vmovdqa 0x1c0(%r15), %ymm13
  1088: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1093: c4417d6faf00020000       vmovdqa 0x200(%r15), %ymm13
  1102: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1107: c4c145effb               vpxor %ymm11, %ymm7, %ymm7
  1112: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
  1117: c4017d6f041e             vmovdqa (%r14, %r11), %ymm8
  1123: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1128: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1134: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1139: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  1144: c4417d6faf40020000       vmovdqa 0x240(%r15), %ymm13
  1153: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1158: c4417d6faf80020000       vmovdqa 0x280(%r15), %ymm13
  1167: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1172: c4c15defe3               vpxor %ymm11, %ymm4, %ymm4
  1177: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
  1182: c4017d6f441e20           vmovdqa 0x20(%r14, %r11), %ymm8
  1189: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1194: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1200: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1205: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  1210: c4417d6faf40020000       vmovdqa 0x240(%r15), %ymm13
  1219: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1224: c4417d6faf80020000       vmovdqa 0x280(%r15), %ymm13
  1233: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1238: c4c155efeb               vpxor %ymm11, %ymm5, %ymm5
  1243: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
  1248: c4017d6f441e40           vmovdqa 0x40(%r14, %r11), %ymm8
  1255: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1260: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1266: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1271: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  1276: c4417d6faf40020000       vmovdqa 0x240(%r15), %ymm13
  1285: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1290: c4417d6faf80020000       vmovdqa 0x280(%r15), %ymm13
  1299: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1304: c4c14deff3               vpxor %ymm11, %ymm6, %ymm6
  1309: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
  1314: c4017d6f441e60           vmovdqa 0x60(%r14, %r11), %ymm8
  1321: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1326: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1332: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1337: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  1342: c4417d6faf40020000       vmovdqa 0x240(%r15), %ymm13
  1351: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1356: c4417d6faf80020000       vmovdqa 0x280(%r15), %ymm13
  1365: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1370: c4c145effb               vpxor %ymm11, %ymm7, %ymm7
  1375: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
  1380: 4d8b5028                 movq 0x28(%r8), %r10
  1384: c4017d6f041a             vmovdqa (%r10, %r11), %ymm8
  1390: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1395: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1401: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1406: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  1411: c4417d6fafc0020000       vmovdqa 0x2c0(%r15), %ymm13
  1420: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1425: c4417d6faf00030000       vmovdqa 0x300(%r15), %ymm13
  1434: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1439: c4c15defe3               vpxor %ymm11, %ymm4, %ymm4
  1444: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
  1449: c4017d6f441a20           vmovdqa 0x20(%r10, %r11), %ymm8
  1456: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1461: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1467: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1472: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  1477: c4417d6fafc0020000       vmovdqa 0x2c0(%r15), %ymm13
  1486: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1491: c4417d6faf00030000       vmovdqa 0x300(%r15), %ymm13
  1500: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1505: c4c155efeb               vpxor %ymm11, %ymm5, %ymm5
  1510: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
  1515: c4017d6f441a40           vmovdqa 0x40(%r10, %r11), %ymm8
  1522: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1527: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1533: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1538: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  1543: c4417d6fafc0020000       vmovdqa 0x2c0(%r15), %ymm13
  1552: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1557: c4417d6faf00030000       vmovdqa 0x300(%r15), %ymm13
  1566: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1571: c4c14deff3               vpxor %ymm11, %ymm6, %ymm6
  1576: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
  1581: c4017d6f441a60           vmovdqa 0x60(%r10, %r11), %ymm8
  1588: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1593: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1599: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1604: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  1609: c4417d6fafc0020000       vmovdqa 0x2c0(%r15), %ymm13
  1618: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1623: c4417d6faf00030000       vmovdqa 0x300(%r15), %ymm13
  1632: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1637: c4c145effb               vpxor %ymm11, %ymm7, %ymm7
  1642: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
  1647: 4d8b5030                 movq 0x30(%r8), %r10
  1651: c4017d6f041a             vmovdqa (%r10, %r11), %ymm8
  1657: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1662: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1668: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1673: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  1678: c4417d6faf40030000       vmovdqa 0x340(%r15), %ymm13
  1687: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1692: c4417d6faf80030000       vmovdqa 0x380(%r15), %ymm13
  1701: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1706: c4c15defe3               vpxor %ymm11, %ymm4, %ymm4
  1711: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
  1716: c4017d6f441a20           vmovdqa 0x20(%r10, %r11), %ymm8
  1723: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1728: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1734: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1739: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  1744: c4417d6faf40030000       vmovdqa 0x340(%r15), %ymm13
  1753: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1758: c4417d6faf80030000       vmovdqa 0x380(%r15), %ymm13
  1767: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1772: c4c155efeb               vpxor %ymm11, %ymm5, %ymm5
  1777: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
  1782: c4017d6f441a40           vmovdqa 0x40(%r10, %r11), %ymm8
  1789: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1794: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1800: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1805: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  1810: c4417d6faf40030000       vmovdqa 0x340(%r15), %ymm13
  1819: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1824: c4417d6faf80030000       vmovdqa 0x380(%r15), %ymm13
  1833: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1838: c4c14deff3               vpxor %ymm11, %ymm6, %ymm6
  1843: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
  1848: c4017d6f441a60           vmovdqa 0x60(%r10, %r11), %ymm8
  1855: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1860: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1866: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1871: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  1876: c4417d6faf40030000       vmovdqa 0x340(%r15), %ymm13
  1885: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
  1890: c4417d6faf80030000       vmovdqa 0x380(%r15), %ymm13
  1899: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
  1904: c4c145effb               vpxor %ymm11, %ymm7, %ymm7
  1909: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
  1914: 4d8b5038                 movq 0x38(%r8), %r10
  1918: c4017d6f041a             vmovdqa (%r10, %r11), %ymm8
  1924: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  1929: c4c15defe0               vpxor %ymm8, %ymm4, %ymm4
  1934: c4017d6f441a20           vmovdqa 0x20(%r10, %r11), %ymm8
  1941: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  1946: c4c155efe8               vpxor %ymm8, %ymm5, %ymm5
  1951: c4017d6f441a40           vmovdqa 0x40(%r10, %r11), %ymm8
  1958: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  1963: c4c14deff0               vpxor %ymm8, %ymm6, %ymm6
  1968: c4017d6f441a60           vmovdqa 0x60(%r10, %r11), %ymm8
  1975: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  1980: c4c145eff8               vpxor %ymm8, %ymm7, %ymm7
  1985: 4d8b11                   movq (%r9), %r10
  1988: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
  1994: c4817d7f4c1a20           vmovdqa %ymm1, 0x20(%r10, %r11)
  2001: c4817d7f541a40           vmovdqa %ymm2, 0x40(%r10, %r11)
  2008: c4817d7f5c1a60           vmovdqa %ymm3, 0x60(%r10, %r11)
  2015: 4d8b5108                 movq 8(%r9), %r10
  2019: c4817d7f241a             vmovdqa %ymm4, (%r10, %r11)
  2025: c4817d7f6c1a20           vmovdqa %ymm5, 0x20(%r10, %r11)
  2032: c4817d7f741a40           vmovdqa %ymm6, 0x40(%r10, %r11)
  2039: c4817d7f7c1a60           vmovdqa %ymm7, 0x60(%r10, %r11)
  2046: 4981c380000000           addq $0x80, %r11
  2053: 4c3bd8                   cmpq %rax, %r11
  2056: 0f8652f8ffff             jbe 0x60
  2062: 4c3bd9                   cmpq %rcx, %r11
  2065: 0f830f020000             jae 0xa26
  2071: 4153                     pushq %r11
  2073: 498dbfc0030000           leaq 0x3c0(%r15), %rdi
  2080: 498db7c0040000           leaq 0x4c0(%r15), %rsi
  2087: 4d8b11                   movq (%r9), %r10
  2090: 4c8b1c24                 movq (%rsp), %r11
  2094: 33db                     xorl %ebx, %ebx
  2096: 498b00                   movq (%r8), %rax
  2099: 420fb61418               movzbl (%rax, %r11), %edx
  2104: 32da                     xorb %dl, %bl
  2106: 498b4008                 movq 8(%r8), %rax
  2110: 420fb61418               movzbl (%rax, %r11), %edx
  2115: 32da                     xorb %dl, %bl
  2117: 498b4010                 movq 0x10(%r8), %rax
  2121: 420fb61418               movzbl (%rax, %r11), %edx
  2126: 32da                     xorb %dl, %bl
  2128: 498b4018                 movq 0x18(%r8), %rax
  2132: 420fb61418               movzbl (%rax, %r11), %edx
  2137: 32da                     xorb %dl, %bl
  2139: 498b4020                 movq 0x20(%r8), %rax
  2143: 420fb61418               movzbl (%rax, %r11), %edx
  2148: 32da                     xorb %dl, %bl
  2150: 498b4028                 movq 0x28(%r8), %rax
  2154: 420fb61418               movzbl (%rax, %r11), %edx
  2159: 32da                     xorb %dl, %bl
  2161: 498b4030                 movq 0x30(%r8), %rax
  2165: 420fb61418               movzbl (%rax, %r11), %edx
  2170: 32da                     xorb %dl, %bl
  2172: 498b4038                 movq 0x38(%r8), %rax
  2176: 420fb61418               movzbl (%rax, %r11), %edx
  2181: 32da                     xorb %dl, %bl
  2183: 43881c1a                 movb %bl, (%r10, %r11)
  2187: 33db                     xorl %ebx, %ebx
  2189: 49ffc3                   incq %r11
  2192: 4c3bd9                   cmpq %rcx, %r11
  2195: 0f8297ffffff             jb 0x830
  2201: 4d8b5108                 movq 8(%r9), %r10
  2205: 4c8b1c24                 movq (%rsp), %r11
  2209: 33db                     xorl %ebx, %ebx
  2211: 498b00                   movq (%r8), %rax
  2214: 420fb61418               movzbl (%rax, %r11), %edx
  2219: 4885d2                   testq %rdx, %rdx
  2222: 0f8422000000             je 0x8d6
  2228: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2232: 81c0e6000000             addl $0xe6, %eax
  2238: 81f8ff000000             cmpl $0xff, %eax
  2244: 0f8206000000             jb 0x8d0
  2250: 81c001ffffff             addl $0xffffff01, %eax
  2256: 0fb60406                 movzbl (%rsi, %rax), %eax
  2260: 32d8                     xorb %al, %bl
  2262: 498b4008                 movq 8(%r8), %rax
  2266: 420fb61418               movzbl (%rax, %r11), %edx
  2271: 4885d2                   testq %rdx, %rdx
  2274: 0f8422000000             je 0x90a
  2280: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2284: 81c0cd000000             addl $0xcd, %eax
  2290: 81f8ff000000             cmpl $0xff, %eax
  2296: 0f8206000000             jb 0x904
  2302: 81c001ffffff             addl $0xffffff01, %eax
  2308: 0fb60406                 movzbl (%rsi, %rax), %eax
  2312: 32d8                     xorb %al, %bl
  2314: 498b4010                 movq 0x10(%r8), %rax
  2318: 420fb61418               movzbl (%rax, %r11), %edx
  2323: 4885d2                   testq %rdx, %rdx
  2326: 0f8422000000             je 0x93e
  2332: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2336: 81c0e7000000             addl $0xe7, %eax
  2342: 81f8ff000000             cmpl $0xff, %eax
  2348: 0f8206000000             jb 0x938
  2354: 81c001ffffff             addl $0xffffff01, %eax
  2360: 0fb60406                 movzbl (%rsi, %rax), %eax
  2364: 32d8                     xorb %al, %bl
  2366: 498b4018                 movq 0x18(%r8), %rax
  2370: 420fb61418               movzbl (%rax, %r11), %edx
  2375: 4885d2                   testq %rdx, %rdx
  2378: 0f8422000000             je 0x972
  2384: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2388: 81c0ce000000             addl $0xce, %eax
  2394: 81f8ff000000             cmpl $0xff, %eax
  2400: 0f8206000000             jb 0x96c
  2406: 81c001ffffff             addl $0xffffff01, %eax
  2412: 0fb60406                 movzbl (%rsi, %rax), %eax
  2416: 32d8                     xorb %al, %bl
  2418: 498b4020                 movq 0x20(%r8), %rax
  2422: 420fb61418               movzbl (%rax, %r11), %edx
  2427: 4885d2                   testq %rdx, %rdx
  2430: 0f841f000000             je 0x9a3
  2436: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2440: 83c04b                   addl $0x4b, %eax
  2443: 81f8ff000000             cmpl $0xff, %eax
  2449: 0f8206000000             jb 0x99d
  2455: 81c001ffffff             addl $0xffffff01, %eax
  2461: 0fb60406                 movzbl (%rsi, %rax), %eax
  2465: 32d8                     xorb %al, %bl
  2467: 498b4028                 movq 0x28(%r8), %rax
  2471: 420fb61418               movzbl (%rax, %r11), %edx
  2476: 4885d2                   testq %rdx, %rdx
  2479: 0f841f000000             je 0x9d4
  2485: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2489: 83c032                   addl $0x32, %eax
  2492: 81f8ff000000             cmpl $0xff, %eax
  2498: 0f8206000000             jb 0x9ce
  2504: 81c001ffffff             addl $0xffffff01, %eax
  2510: 0fb60406                 movzbl (%rsi, %rax), %eax
  2514: 32d8                     xorb %al, %bl
  2516: 498b4030                 movq 0x30(%r8), %rax
  2520: 420fb61418               movzbl (%rax, %r11), %edx
  2525: 4885d2                   testq %rdx, %rdx
  2528: 0f841f000000             je 0xa05
  2534: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2538: 83c019                   addl $0x19, %eax
  2541: 81f8ff000000             cmpl $0xff, %eax
  2547: 0f8206000000             jb 0x9ff
  2553: 81c001ffffff             addl $0xffffff01, %eax
  2559: 0fb60406                 movzbl (%rsi, %rax), %eax
  2563: 32d8                     xorb %al, %bl
  2565: 498b4038                 movq 0x38(%r8), %rax
  2569: 420fb61418               movzbl (%rax, %r11), %edx
  2574: 32da                     xorb %dl, %bl
  2576: 43881c1a                 movb %bl, (%r10, %r11)
  2580: 33db                     xorl %ebx, %ebx
  2582: 49ffc3                   incq %r11
  2585: 4c3bd9                   cmpq %rcx, %r11
  2588: 0f8281feffff             jb 0x8a3
  2594: 4883c408                 addq $8, %rsp
  2598: 415e                     popq %r14
  2600: 415d                     popq %r13
  2602: 415c                     popq %r12
  2604: 415f                     popq %r15
  2606: 5b                       popq %rbx
  2607: c5f877                   vzeroupper 
  2610: c3                       retq 
  2611: 0f0b                     ud2 
  2613: 0f0b                     ud2 
  2615: 0f0b                     ud2 
  2617: 0f0b                     ud2 
  2619: 0f0b                     ud2 
  2621: 0f0b                     ud2 
  2623: cc                       int3 
