; config_t1_468 target=avx2 twin=baseline topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=2624 twins=[] uops_per_byte=2.3515625
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
    62: 0f822c070000             jb 0x770
    68: 666666666666660f1f840000000000 nopw (%rax, %rax)
    83: 66666666660f1f840000000000 nopw (%rax, %rax)
    96: c4217d6f041e             vmovdqa (%rsi, %r11), %ymm8
   102: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   107: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   113: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   118: c4c17d6fc0               vmovdqa %ymm8, %ymm0
   123: c4417d6f5f40             vmovdqa 0x40(%r15), %ymm11
   129: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   134: c4417d6fa780000000       vmovdqa 0x80(%r15), %ymm12
   143: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   148: c4c125efe4               vpxor %ymm12, %ymm11, %ymm4
   153: c4217d6f441e20           vmovdqa 0x20(%rsi, %r11), %ymm8
   160: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   165: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   171: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   176: c4c17d6fc8               vmovdqa %ymm8, %ymm1
   181: c4417d6f5f40             vmovdqa 0x40(%r15), %ymm11
   187: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   192: c4417d6fa780000000       vmovdqa 0x80(%r15), %ymm12
   201: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   206: c4c125efec               vpxor %ymm12, %ymm11, %ymm5
   211: c4217d6f441e40           vmovdqa 0x40(%rsi, %r11), %ymm8
   218: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   223: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   229: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   234: c4c17d6fd0               vmovdqa %ymm8, %ymm2
   239: c4417d6f5f40             vmovdqa 0x40(%r15), %ymm11
   245: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   250: c4417d6fa780000000       vmovdqa 0x80(%r15), %ymm12
   259: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   264: c4c125eff4               vpxor %ymm12, %ymm11, %ymm6
   269: c4217d6f441e60           vmovdqa 0x60(%rsi, %r11), %ymm8
   276: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   281: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   287: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   292: c4c17d6fd8               vmovdqa %ymm8, %ymm3
   297: c4417d6f5f40             vmovdqa 0x40(%r15), %ymm11
   303: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   308: c4417d6fa780000000       vmovdqa 0x80(%r15), %ymm12
   317: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   322: c4c125effc               vpxor %ymm12, %ymm11, %ymm7
   327: c4217d6f041f             vmovdqa (%rdi, %r11), %ymm8
   333: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   338: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   344: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   349: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   354: c4417d6f9fc0000000       vmovdqa 0xc0(%r15), %ymm11
   363: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   368: c4417d6fa700010000       vmovdqa 0x100(%r15), %ymm12
   377: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   382: c4c15defe3               vpxor %ymm11, %ymm4, %ymm4
   387: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   392: c4217d6f441f20           vmovdqa 0x20(%rdi, %r11), %ymm8
   399: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   404: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   410: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   415: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   420: c4417d6f9fc0000000       vmovdqa 0xc0(%r15), %ymm11
   429: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   434: c4417d6fa700010000       vmovdqa 0x100(%r15), %ymm12
   443: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   448: c4c155efeb               vpxor %ymm11, %ymm5, %ymm5
   453: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   458: c4217d6f441f40           vmovdqa 0x40(%rdi, %r11), %ymm8
   465: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   470: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   476: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   481: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   486: c4417d6f9fc0000000       vmovdqa 0xc0(%r15), %ymm11
   495: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   500: c4417d6fa700010000       vmovdqa 0x100(%r15), %ymm12
   509: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   514: c4c14deff3               vpxor %ymm11, %ymm6, %ymm6
   519: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   524: c4217d6f441f60           vmovdqa 0x60(%rdi, %r11), %ymm8
   531: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   536: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   542: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   547: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
   552: c4417d6f9fc0000000       vmovdqa 0xc0(%r15), %ymm11
   561: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   566: c4417d6fa700010000       vmovdqa 0x100(%r15), %ymm12
   575: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   580: c4c145effb               vpxor %ymm11, %ymm7, %ymm7
   585: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   590: c4017d6f041c             vmovdqa (%r12, %r11), %ymm8
   596: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   601: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   607: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   612: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   617: c4417d6f9f40010000       vmovdqa 0x140(%r15), %ymm11
   626: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   631: c4417d6fa780010000       vmovdqa 0x180(%r15), %ymm12
   640: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   645: c4c15defe3               vpxor %ymm11, %ymm4, %ymm4
   650: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   655: c4017d6f441c20           vmovdqa 0x20(%r12, %r11), %ymm8
   662: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   667: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   673: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   678: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   683: c4417d6f9f40010000       vmovdqa 0x140(%r15), %ymm11
   692: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   697: c4417d6fa780010000       vmovdqa 0x180(%r15), %ymm12
   706: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   711: c4c155efeb               vpxor %ymm11, %ymm5, %ymm5
   716: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   721: c4017d6f441c40           vmovdqa 0x40(%r12, %r11), %ymm8
   728: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   733: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   739: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   744: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   749: c4417d6f9f40010000       vmovdqa 0x140(%r15), %ymm11
   758: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   763: c4417d6fa780010000       vmovdqa 0x180(%r15), %ymm12
   772: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   777: c4c14deff3               vpxor %ymm11, %ymm6, %ymm6
   782: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   787: c4017d6f441c60           vmovdqa 0x60(%r12, %r11), %ymm8
   794: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   799: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   805: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   810: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
   815: c4417d6f9f40010000       vmovdqa 0x140(%r15), %ymm11
   824: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
   829: c4417d6fa780010000       vmovdqa 0x180(%r15), %ymm12
   838: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
   843: c4c145effb               vpxor %ymm11, %ymm7, %ymm7
   848: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   853: c4017d6f441d00           vmovdqa (%r13, %r11), %ymm8
   860: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   865: c4c15defe0               vpxor %ymm8, %ymm4, %ymm4
   870: c4017d6f441d20           vmovdqa 0x20(%r13, %r11), %ymm8
   877: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   882: c4c155efe8               vpxor %ymm8, %ymm5, %ymm5
   887: c4017d6f441d40           vmovdqa 0x40(%r13, %r11), %ymm8
   894: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   899: c4c14deff0               vpxor %ymm8, %ymm6, %ymm6
   904: c4017d6f441d60           vmovdqa 0x60(%r13, %r11), %ymm8
   911: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
   916: c4c145eff8               vpxor %ymm8, %ymm7, %ymm7
   921: c4017d6f041e             vmovdqa (%r14, %r11), %ymm8
   927: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   932: c4017d6f441e20           vmovdqa 0x20(%r14, %r11), %ymm8
   939: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   944: c4017d6f441e40           vmovdqa 0x40(%r14, %r11), %ymm8
   951: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   956: c4017d6f441e60           vmovdqa 0x60(%r14, %r11), %ymm8
   963: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
   968: 4d8b5028                 movq 0x28(%r8), %r10
   972: c4017d6f041a             vmovdqa (%r10, %r11), %ymm8
   978: c4c15defe0               vpxor %ymm8, %ymm4, %ymm4
   983: c4017d6f441a20           vmovdqa 0x20(%r10, %r11), %ymm8
   990: c4c155efe8               vpxor %ymm8, %ymm5, %ymm5
   995: c4017d6f441a40           vmovdqa 0x40(%r10, %r11), %ymm8
  1002: c4c14deff0               vpxor %ymm8, %ymm6, %ymm6
  1007: c4017d6f441a60           vmovdqa 0x60(%r10, %r11), %ymm8
  1014: c4c145eff8               vpxor %ymm8, %ymm7, %ymm7
  1019: c4417ddbce               vpand %ymm14, %ymm0, %ymm9
  1024: c5ad71d004               vpsrlw $4, %ymm0, %ymm10
  1029: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1034: c4417d6f9fc0010000       vmovdqa 0x1c0(%r15), %ymm11
  1043: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1048: c4417d6fa700020000       vmovdqa 0x200(%r15), %ymm12
  1057: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1062: c44125efec               vpxor %ymm12, %ymm11, %ymm13
  1067: c515efec                 vpxor %ymm4, %ymm13, %ymm13
  1071: c4417ddbce               vpand %ymm14, %ymm0, %ymm9
  1076: c5ad71d004               vpsrlw $4, %ymm0, %ymm10
  1081: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1086: c4417d6f9f40020000       vmovdqa 0x240(%r15), %ymm11
  1095: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1100: c4417d6fa780020000       vmovdqa 0x280(%r15), %ymm12
  1109: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1114: c44125efc4               vpxor %ymm12, %ymm11, %ymm8
  1119: c53defc4                 vpxor %ymm4, %ymm8, %ymm8
  1123: c44115dbce               vpand %ymm14, %ymm13, %ymm9
  1128: c4c12d71d504             vpsrlw $4, %ymm13, %ymm10
  1134: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1139: c4417d6f9fc0020000       vmovdqa 0x2c0(%r15), %ymm11
  1148: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1153: c4417d6fa700030000       vmovdqa 0x300(%r15), %ymm12
  1162: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1167: c4c125efe4               vpxor %ymm12, %ymm11, %ymm4
  1172: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1177: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1183: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1188: c4417d6f9fc0020000       vmovdqa 0x2c0(%r15), %ymm11
  1197: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1202: c4417d6fa700030000       vmovdqa 0x300(%r15), %ymm12
  1211: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1216: c4c125efc4               vpxor %ymm12, %ymm11, %ymm0
  1221: c44175dbce               vpand %ymm14, %ymm1, %ymm9
  1226: c5ad71d104               vpsrlw $4, %ymm1, %ymm10
  1231: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1236: c4417d6f9fc0010000       vmovdqa 0x1c0(%r15), %ymm11
  1245: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1250: c4417d6fa700020000       vmovdqa 0x200(%r15), %ymm12
  1259: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1264: c44125efec               vpxor %ymm12, %ymm11, %ymm13
  1269: c515efed                 vpxor %ymm5, %ymm13, %ymm13
  1273: c44175dbce               vpand %ymm14, %ymm1, %ymm9
  1278: c5ad71d104               vpsrlw $4, %ymm1, %ymm10
  1283: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1288: c4417d6f9f40020000       vmovdqa 0x240(%r15), %ymm11
  1297: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1302: c4417d6fa780020000       vmovdqa 0x280(%r15), %ymm12
  1311: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1316: c44125efc4               vpxor %ymm12, %ymm11, %ymm8
  1321: c53defc5                 vpxor %ymm5, %ymm8, %ymm8
  1325: c44115dbce               vpand %ymm14, %ymm13, %ymm9
  1330: c4c12d71d504             vpsrlw $4, %ymm13, %ymm10
  1336: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1341: c4417d6f9fc0020000       vmovdqa 0x2c0(%r15), %ymm11
  1350: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1355: c4417d6fa700030000       vmovdqa 0x300(%r15), %ymm12
  1364: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1369: c4c125efec               vpxor %ymm12, %ymm11, %ymm5
  1374: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1379: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1385: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1390: c4417d6f9fc0020000       vmovdqa 0x2c0(%r15), %ymm11
  1399: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1404: c4417d6fa700030000       vmovdqa 0x300(%r15), %ymm12
  1413: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1418: c4c125efcc               vpxor %ymm12, %ymm11, %ymm1
  1423: c4416ddbce               vpand %ymm14, %ymm2, %ymm9
  1428: c5ad71d204               vpsrlw $4, %ymm2, %ymm10
  1433: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1438: c4417d6f9fc0010000       vmovdqa 0x1c0(%r15), %ymm11
  1447: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1452: c4417d6fa700020000       vmovdqa 0x200(%r15), %ymm12
  1461: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1466: c44125efec               vpxor %ymm12, %ymm11, %ymm13
  1471: c515efee                 vpxor %ymm6, %ymm13, %ymm13
  1475: c4416ddbce               vpand %ymm14, %ymm2, %ymm9
  1480: c5ad71d204               vpsrlw $4, %ymm2, %ymm10
  1485: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1490: c4417d6f9f40020000       vmovdqa 0x240(%r15), %ymm11
  1499: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1504: c4417d6fa780020000       vmovdqa 0x280(%r15), %ymm12
  1513: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1518: c44125efc4               vpxor %ymm12, %ymm11, %ymm8
  1523: c53defc6                 vpxor %ymm6, %ymm8, %ymm8
  1527: c44115dbce               vpand %ymm14, %ymm13, %ymm9
  1532: c4c12d71d504             vpsrlw $4, %ymm13, %ymm10
  1538: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1543: c4417d6f9fc0020000       vmovdqa 0x2c0(%r15), %ymm11
  1552: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1557: c4417d6fa700030000       vmovdqa 0x300(%r15), %ymm12
  1566: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1571: c4c125eff4               vpxor %ymm12, %ymm11, %ymm6
  1576: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1581: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1587: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1592: c4417d6f9fc0020000       vmovdqa 0x2c0(%r15), %ymm11
  1601: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1606: c4417d6fa700030000       vmovdqa 0x300(%r15), %ymm12
  1615: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1620: c4c125efd4               vpxor %ymm12, %ymm11, %ymm2
  1625: c44165dbce               vpand %ymm14, %ymm3, %ymm9
  1630: c5ad71d304               vpsrlw $4, %ymm3, %ymm10
  1635: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1640: c4417d6f9fc0010000       vmovdqa 0x1c0(%r15), %ymm11
  1649: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1654: c4417d6fa700020000       vmovdqa 0x200(%r15), %ymm12
  1663: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1668: c44125efec               vpxor %ymm12, %ymm11, %ymm13
  1673: c515efef                 vpxor %ymm7, %ymm13, %ymm13
  1677: c44165dbce               vpand %ymm14, %ymm3, %ymm9
  1682: c5ad71d304               vpsrlw $4, %ymm3, %ymm10
  1687: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1692: c4417d6f9f40020000       vmovdqa 0x240(%r15), %ymm11
  1701: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1706: c4417d6fa780020000       vmovdqa 0x280(%r15), %ymm12
  1715: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1720: c44125efc4               vpxor %ymm12, %ymm11, %ymm8
  1725: c53defc7                 vpxor %ymm7, %ymm8, %ymm8
  1729: c44115dbce               vpand %ymm14, %ymm13, %ymm9
  1734: c4c12d71d504             vpsrlw $4, %ymm13, %ymm10
  1740: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1745: c4417d6f9fc0020000       vmovdqa 0x2c0(%r15), %ymm11
  1754: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1759: c4417d6fa700030000       vmovdqa 0x300(%r15), %ymm12
  1768: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1773: c4c125effc               vpxor %ymm12, %ymm11, %ymm7
  1778: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
  1783: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
  1789: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
  1794: c4417d6f9fc0020000       vmovdqa 0x2c0(%r15), %ymm11
  1803: c4422500d9               vpshufb %ymm9, %ymm11, %ymm11
  1808: c4417d6fa700030000       vmovdqa 0x300(%r15), %ymm12
  1817: c4421d00e2               vpshufb %ymm10, %ymm12, %ymm12
  1822: c4c125efdc               vpxor %ymm12, %ymm11, %ymm3
  1827: 4d8b11                   movq (%r9), %r10
  1830: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
  1836: c4817d7f4c1a20           vmovdqa %ymm1, 0x20(%r10, %r11)
  1843: c4817d7f541a40           vmovdqa %ymm2, 0x40(%r10, %r11)
  1850: c4817d7f5c1a60           vmovdqa %ymm3, 0x60(%r10, %r11)
  1857: 4d8b5108                 movq 8(%r9), %r10
  1861: c4817d7f241a             vmovdqa %ymm4, (%r10, %r11)
  1867: c4817d7f6c1a20           vmovdqa %ymm5, 0x20(%r10, %r11)
  1874: c4817d7f741a40           vmovdqa %ymm6, 0x40(%r10, %r11)
  1881: c4817d7f7c1a60           vmovdqa %ymm7, 0x60(%r10, %r11)
  1888: 4981c380000000           addq $0x80, %r11
  1895: 4c3bd8                   cmpq %rax, %r11
  1898: 0f86f0f8ffff             jbe 0x60
  1904: 4c3bd9                   cmpq %rcx, %r11
  1907: 0f83a1020000             jae 0xa1a
  1913: 4153                     pushq %r11
  1915: 498dbf40030000           leaq 0x340(%r15), %rdi
  1922: 498db740040000           leaq 0x440(%r15), %rsi
  1929: 4d8b11                   movq (%r9), %r10
  1932: 4c8b1c24                 movq (%rsp), %r11
  1936: 33db                     xorl %ebx, %ebx
  1938: 498b00                   movq (%r8), %rax
  1941: 420fb61418               movzbl (%rax, %r11), %edx
  1946: 4885d2                   testq %rdx, %rdx
  1949: 0f8422000000             je 0x7c5
  1955: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1959: 81c0c1000000             addl $0xc1, %eax
  1965: 81f8ff000000             cmpl $0xff, %eax
  1971: 0f8206000000             jb 0x7bf
  1977: 81c001ffffff             addl $0xffffff01, %eax
  1983: 0fb60406                 movzbl (%rsi, %rax), %eax
  1987: 32d8                     xorb %al, %bl
  1989: 498b4008                 movq 8(%r8), %rax
  1993: 420fb61418               movzbl (%rax, %r11), %edx
  1998: 4885d2                   testq %rdx, %rdx
  2001: 0f841f000000             je 0x7f6
  2007: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2011: 83c003                   addl $3, %eax
  2014: 81f8ff000000             cmpl $0xff, %eax
  2020: 0f8206000000             jb 0x7f0
  2026: 81c001ffffff             addl $0xffffff01, %eax
  2032: 0fb60406                 movzbl (%rsi, %rax), %eax
  2036: 32d8                     xorb %al, %bl
  2038: 498b4010                 movq 0x10(%r8), %rax
  2042: 420fb61418               movzbl (%rax, %r11), %edx
  2047: 4885d2                   testq %rdx, %rdx
  2050: 0f8422000000             je 0x82a
  2056: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2060: 81c0c6000000             addl $0xc6, %eax
  2066: 81f8ff000000             cmpl $0xff, %eax
  2072: 0f8206000000             jb 0x824
  2078: 81c001ffffff             addl $0xffffff01, %eax
  2084: 0fb60406                 movzbl (%rsi, %rax), %eax
  2088: 32d8                     xorb %al, %bl
  2090: 498b4018                 movq 0x18(%r8), %rax
  2094: 420fb61418               movzbl (%rax, %r11), %edx
  2099: 4885d2                   testq %rdx, %rdx
  2102: 0f841f000000             je 0x85b
  2108: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2112: 83c001                   addl $1, %eax
  2115: 81f8ff000000             cmpl $0xff, %eax
  2121: 0f8206000000             jb 0x855
  2127: 81c001ffffff             addl $0xffffff01, %eax
  2133: 0fb60406                 movzbl (%rsi, %rax), %eax
  2137: 32d8                     xorb %al, %bl
  2139: 498b4020                 movq 0x20(%r8), %rax
  2143: 420fb61418               movzbl (%rax, %r11), %edx
  2148: 4885d2                   testq %rdx, %rdx
  2151: 0f8422000000             je 0x88f
  2157: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2161: 81c0fe000000             addl $0xfe, %eax
  2167: 81f8ff000000             cmpl $0xff, %eax
  2173: 0f8206000000             jb 0x889
  2179: 81c001ffffff             addl $0xffffff01, %eax
  2185: 0fb60406                 movzbl (%rsi, %rax), %eax
  2189: 32d8                     xorb %al, %bl
  2191: 498b4028                 movq 0x28(%r8), %rax
  2195: 420fb61418               movzbl (%rax, %r11), %edx
  2200: 4885d2                   testq %rdx, %rdx
  2203: 0f841f000000             je 0x8c0
  2209: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2213: 83c031                   addl $0x31, %eax
  2216: 81f8ff000000             cmpl $0xff, %eax
  2222: 0f8206000000             jb 0x8ba
  2228: 81c001ffffff             addl $0xffffff01, %eax
  2234: 0fb60406                 movzbl (%rsi, %rax), %eax
  2238: 32d8                     xorb %al, %bl
  2240: 43881c1a                 movb %bl, (%r10, %r11)
  2244: 33db                     xorl %ebx, %ebx
  2246: 49ffc3                   incq %r11
  2249: 4c3bd9                   cmpq %rcx, %r11
  2252: 0f82c0feffff             jb 0x792
  2258: 4d8b5108                 movq 8(%r9), %r10
  2262: 4c8b1c24                 movq (%rsp), %r11
  2266: 33db                     xorl %ebx, %ebx
  2268: 498b00                   movq (%r8), %rax
  2271: 420fb61418               movzbl (%rax, %r11), %edx
  2276: 4885d2                   testq %rdx, %rdx
  2279: 0f841f000000             je 0x90c
  2285: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2289: 83c01c                   addl $0x1c, %eax
  2292: 81f8ff000000             cmpl $0xff, %eax
  2298: 0f8206000000             jb 0x906
  2304: 81c001ffffff             addl $0xffffff01, %eax
  2310: 0fb60406                 movzbl (%rsi, %rax), %eax
  2314: 32d8                     xorb %al, %bl
  2316: 498b4008                 movq 8(%r8), %rax
  2320: 420fb61418               movzbl (%rax, %r11), %edx
  2325: 4885d2                   testq %rdx, %rdx
  2328: 0f8422000000             je 0x940
  2334: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2338: 81c0df000000             addl $0xdf, %eax
  2344: 81f8ff000000             cmpl $0xff, %eax
  2350: 0f8206000000             jb 0x93a
  2356: 81c001ffffff             addl $0xffffff01, %eax
  2362: 0fb60406                 movzbl (%rsi, %rax), %eax
  2366: 32d8                     xorb %al, %bl
  2368: 498b4010                 movq 0x10(%r8), %rax
  2372: 420fb61418               movzbl (%rax, %r11), %edx
  2377: 4885d2                   testq %rdx, %rdx
  2380: 0f841f000000             je 0x971
  2386: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2390: 83c01a                   addl $0x1a, %eax
  2393: 81f8ff000000             cmpl $0xff, %eax
  2399: 0f8206000000             jb 0x96b
  2405: 81c001ffffff             addl $0xffffff01, %eax
  2411: 0fb60406                 movzbl (%rsi, %rax), %eax
  2415: 32d8                     xorb %al, %bl
  2417: 498b4018                 movq 0x18(%r8), %rax
  2421: 420fb61418               movzbl (%rax, %r11), %edx
  2426: 4885d2                   testq %rdx, %rdx
  2429: 0f841f000000             je 0x9a2
  2435: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2439: 83c019                   addl $0x19, %eax
  2442: 81f8ff000000             cmpl $0xff, %eax
  2448: 0f8206000000             jb 0x99c
  2454: 81c001ffffff             addl $0xffffff01, %eax
  2460: 0fb60406                 movzbl (%rsi, %rax), %eax
  2464: 32d8                     xorb %al, %bl
  2466: 498b4020                 movq 0x20(%r8), %rax
  2470: 420fb61418               movzbl (%rax, %r11), %edx
  2475: 4885d2                   testq %rdx, %rdx
  2478: 0f841f000000             je 0x9d3
  2484: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2488: 83c018                   addl $0x18, %eax
  2491: 81f8ff000000             cmpl $0xff, %eax
  2497: 0f8206000000             jb 0x9cd
  2503: 81c001ffffff             addl $0xffffff01, %eax
  2509: 0fb60406                 movzbl (%rsi, %rax), %eax
  2513: 32d8                     xorb %al, %bl
  2515: 498b4028                 movq 0x28(%r8), %rax
  2519: 420fb61418               movzbl (%rax, %r11), %edx
  2524: 4885d2                   testq %rdx, %rdx
  2527: 0f841f000000             je 0xa04
  2533: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2537: 83c031                   addl $0x31, %eax
  2540: 81f8ff000000             cmpl $0xff, %eax
  2546: 0f8206000000             jb 0x9fe
  2552: 81c001ffffff             addl $0xffffff01, %eax
  2558: 0fb60406                 movzbl (%rsi, %rax), %eax
  2562: 32d8                     xorb %al, %bl
  2564: 43881c1a                 movb %bl, (%r10, %r11)
  2568: 33db                     xorl %ebx, %ebx
  2570: 49ffc3                   incq %r11
  2573: 4c3bd9                   cmpq %rcx, %r11
  2576: 0f82c6feffff             jb 0x8dc
  2582: 4883c408                 addq $8, %rsp
  2586: 415e                     popq %r14
  2588: 415d                     popq %r13
  2590: 415c                     popq %r12
  2592: 415f                     popq %r15
  2594: 5b                       popq %rbx
  2595: c5f877                   vzeroupper 
  2598: c3                       retq 
  2599: 0f0b                     ud2 
  2601: 0f0b                     ud2 
  2603: 0f0b                     ud2 
  2605: 0f0b                     ud2 
  2607: 0f0b                     ud2 
  2609: 0f0b                     ud2 
  2611: 0f0b                     ud2 
  2613: 0f0b                     ud2 
  2615: 0f0b                     ud2 
  2617: 0f0b                     ud2 
  2619: 0f0b                     ud2 
  2621: 0f0b                     ud2 
  2623: cc                       int3 
