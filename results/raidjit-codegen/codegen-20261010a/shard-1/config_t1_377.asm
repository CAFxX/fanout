; config_t1_377 target=avx2 twin=baseline topo=rs(6,3)/enc op=encode U=8
; code_size=2496 twins=[] uops_per_byte=2.8229166666666665
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da7090000           leaq 0x9a7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: c4417d6f3f               vmovdqa (%r15), %ymm15
    52: 488bc1                   movq %rcx, %rax
    55: 4883e860                 subq $0x60, %rax
    59: 0f829c060000             jb 0x6dd
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4217d6f0c1e             vmovdqa (%rsi, %r11), %ymm9
   102: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   107: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   113: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   118: c4c17d6fc1               vmovdqa %ymm9, %ymm0
   123: c4417d6f7740             vmovdqa 0x40(%r15), %ymm14
   129: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   134: c4417d6fb780000000       vmovdqa 0x80(%r15), %ymm14
   143: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   148: c4c11defdd               vpxor %ymm13, %ymm12, %ymm3
   153: c4417d6fb7c0000000       vmovdqa 0xc0(%r15), %ymm14
   162: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   167: c4417d6fb700010000       vmovdqa 0x100(%r15), %ymm14
   176: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   181: c4c11deff5               vpxor %ymm13, %ymm12, %ymm6
   186: c4217d6f4c1e20           vmovdqa 0x20(%rsi, %r11), %ymm9
   193: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   198: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   204: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   209: c4c17d6fc9               vmovdqa %ymm9, %ymm1
   214: c4417d6f7740             vmovdqa 0x40(%r15), %ymm14
   220: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   225: c4417d6fb780000000       vmovdqa 0x80(%r15), %ymm14
   234: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   239: c4c11defe5               vpxor %ymm13, %ymm12, %ymm4
   244: c4417d6fb7c0000000       vmovdqa 0xc0(%r15), %ymm14
   253: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   258: c4417d6fb700010000       vmovdqa 0x100(%r15), %ymm14
   267: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   272: c4c11deffd               vpxor %ymm13, %ymm12, %ymm7
   277: c4217d6f4c1e40           vmovdqa 0x40(%rsi, %r11), %ymm9
   284: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   289: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   295: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   300: c4c17d6fd1               vmovdqa %ymm9, %ymm2
   305: c4417d6f7740             vmovdqa 0x40(%r15), %ymm14
   311: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   316: c4417d6fb780000000       vmovdqa 0x80(%r15), %ymm14
   325: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   330: c4c11defed               vpxor %ymm13, %ymm12, %ymm5
   335: c4417d6fb7c0000000       vmovdqa 0xc0(%r15), %ymm14
   344: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   349: c4417d6fb700010000       vmovdqa 0x100(%r15), %ymm14
   358: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   363: c4411defc5               vpxor %ymm13, %ymm12, %ymm8
   368: c4217d6f0c1f             vmovdqa (%rdi, %r11), %ymm9
   374: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   379: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   385: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   390: c4c17defc1               vpxor %ymm9, %ymm0, %ymm0
   395: c4417d6fb7c0000000       vmovdqa 0xc0(%r15), %ymm14
   404: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   409: c4417d6fb700010000       vmovdqa 0x100(%r15), %ymm14
   418: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   423: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   428: c4c165efdd               vpxor %ymm13, %ymm3, %ymm3
   433: c4417d6fb740010000       vmovdqa 0x140(%r15), %ymm14
   442: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   447: c4417d6fb780010000       vmovdqa 0x180(%r15), %ymm14
   456: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   461: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   466: c4c14deff5               vpxor %ymm13, %ymm6, %ymm6
   471: c4217d6f4c1f20           vmovdqa 0x20(%rdi, %r11), %ymm9
   478: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   483: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   489: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   494: c4c175efc9               vpxor %ymm9, %ymm1, %ymm1
   499: c4417d6fb7c0000000       vmovdqa 0xc0(%r15), %ymm14
   508: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   513: c4417d6fb700010000       vmovdqa 0x100(%r15), %ymm14
   522: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   527: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   532: c4c15defe5               vpxor %ymm13, %ymm4, %ymm4
   537: c4417d6fb740010000       vmovdqa 0x140(%r15), %ymm14
   546: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   551: c4417d6fb780010000       vmovdqa 0x180(%r15), %ymm14
   560: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   565: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   570: c4c145effd               vpxor %ymm13, %ymm7, %ymm7
   575: c4217d6f4c1f40           vmovdqa 0x40(%rdi, %r11), %ymm9
   582: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   587: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   593: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   598: c4c16defd1               vpxor %ymm9, %ymm2, %ymm2
   603: c4417d6fb7c0000000       vmovdqa 0xc0(%r15), %ymm14
   612: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   617: c4417d6fb700010000       vmovdqa 0x100(%r15), %ymm14
   626: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   631: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   636: c4c155efed               vpxor %ymm13, %ymm5, %ymm5
   641: c4417d6fb740010000       vmovdqa 0x140(%r15), %ymm14
   650: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   655: c4417d6fb780010000       vmovdqa 0x180(%r15), %ymm14
   664: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   669: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   674: c4413defc5               vpxor %ymm13, %ymm8, %ymm8
   679: c4017d6f0c1c             vmovdqa (%r12, %r11), %ymm9
   685: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   690: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   696: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   701: c4c17defc1               vpxor %ymm9, %ymm0, %ymm0
   706: c4417d6fb740010000       vmovdqa 0x140(%r15), %ymm14
   715: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   720: c4417d6fb780010000       vmovdqa 0x180(%r15), %ymm14
   729: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   734: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   739: c4c165efdd               vpxor %ymm13, %ymm3, %ymm3
   744: c4417d6fb7c0010000       vmovdqa 0x1c0(%r15), %ymm14
   753: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   758: c4417d6fb700020000       vmovdqa 0x200(%r15), %ymm14
   767: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   772: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   777: c4c14deff5               vpxor %ymm13, %ymm6, %ymm6
   782: c4017d6f4c1c20           vmovdqa 0x20(%r12, %r11), %ymm9
   789: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   794: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   800: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   805: c4c175efc9               vpxor %ymm9, %ymm1, %ymm1
   810: c4417d6fb740010000       vmovdqa 0x140(%r15), %ymm14
   819: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   824: c4417d6fb780010000       vmovdqa 0x180(%r15), %ymm14
   833: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   838: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   843: c4c15defe5               vpxor %ymm13, %ymm4, %ymm4
   848: c4417d6fb7c0010000       vmovdqa 0x1c0(%r15), %ymm14
   857: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   862: c4417d6fb700020000       vmovdqa 0x200(%r15), %ymm14
   871: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   876: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   881: c4c145effd               vpxor %ymm13, %ymm7, %ymm7
   886: c4017d6f4c1c40           vmovdqa 0x40(%r12, %r11), %ymm9
   893: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
   898: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
   904: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
   909: c4c16defd1               vpxor %ymm9, %ymm2, %ymm2
   914: c4417d6fb740010000       vmovdqa 0x140(%r15), %ymm14
   923: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   928: c4417d6fb780010000       vmovdqa 0x180(%r15), %ymm14
   937: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   942: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   947: c4c155efed               vpxor %ymm13, %ymm5, %ymm5
   952: c4417d6fb7c0010000       vmovdqa 0x1c0(%r15), %ymm14
   961: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
   966: c4417d6fb700020000       vmovdqa 0x200(%r15), %ymm14
   975: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
   980: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   985: c4413defc5               vpxor %ymm13, %ymm8, %ymm8
   990: c4017d6f4c1d00           vmovdqa (%r13, %r11), %ymm9
   997: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
  1002: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
  1008: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
  1013: c4c17defc1               vpxor %ymm9, %ymm0, %ymm0
  1018: c4417d6fb740020000       vmovdqa 0x240(%r15), %ymm14
  1027: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1032: c4417d6fb780020000       vmovdqa 0x280(%r15), %ymm14
  1041: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1046: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
  1051: c4c165efdd               vpxor %ymm13, %ymm3, %ymm3
  1056: c4417d6f7740             vmovdqa 0x40(%r15), %ymm14
  1062: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1067: c4417d6fb780000000       vmovdqa 0x80(%r15), %ymm14
  1076: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1081: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
  1086: c4c14deff5               vpxor %ymm13, %ymm6, %ymm6
  1091: c4017d6f4c1d20           vmovdqa 0x20(%r13, %r11), %ymm9
  1098: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
  1103: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
  1109: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
  1114: c4c175efc9               vpxor %ymm9, %ymm1, %ymm1
  1119: c4417d6fb740020000       vmovdqa 0x240(%r15), %ymm14
  1128: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1133: c4417d6fb780020000       vmovdqa 0x280(%r15), %ymm14
  1142: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1147: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
  1152: c4c15defe5               vpxor %ymm13, %ymm4, %ymm4
  1157: c4417d6f7740             vmovdqa 0x40(%r15), %ymm14
  1163: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1168: c4417d6fb780000000       vmovdqa 0x80(%r15), %ymm14
  1177: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1182: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
  1187: c4c145effd               vpxor %ymm13, %ymm7, %ymm7
  1192: c4017d6f4c1d40           vmovdqa 0x40(%r13, %r11), %ymm9
  1199: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
  1204: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
  1210: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
  1215: c4c16defd1               vpxor %ymm9, %ymm2, %ymm2
  1220: c4417d6fb740020000       vmovdqa 0x240(%r15), %ymm14
  1229: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1234: c4417d6fb780020000       vmovdqa 0x280(%r15), %ymm14
  1243: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1248: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
  1253: c4c155efed               vpxor %ymm13, %ymm5, %ymm5
  1258: c4417d6f7740             vmovdqa 0x40(%r15), %ymm14
  1264: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1269: c4417d6fb780000000       vmovdqa 0x80(%r15), %ymm14
  1278: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1283: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
  1288: c4413defc5               vpxor %ymm13, %ymm8, %ymm8
  1293: c4017d6f0c1e             vmovdqa (%r14, %r11), %ymm9
  1299: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
  1304: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
  1310: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
  1315: c4c17defc1               vpxor %ymm9, %ymm0, %ymm0
  1320: c4417d6fb7c0010000       vmovdqa 0x1c0(%r15), %ymm14
  1329: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1334: c4417d6fb700020000       vmovdqa 0x200(%r15), %ymm14
  1343: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1348: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
  1353: c4c165efdd               vpxor %ymm13, %ymm3, %ymm3
  1358: c4417d6fb740020000       vmovdqa 0x240(%r15), %ymm14
  1367: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1372: c4417d6fb780020000       vmovdqa 0x280(%r15), %ymm14
  1381: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1386: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
  1391: c4c14deff5               vpxor %ymm13, %ymm6, %ymm6
  1396: c4017d6f4c1e20           vmovdqa 0x20(%r14, %r11), %ymm9
  1403: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
  1408: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
  1414: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
  1419: c4c175efc9               vpxor %ymm9, %ymm1, %ymm1
  1424: c4417d6fb7c0010000       vmovdqa 0x1c0(%r15), %ymm14
  1433: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1438: c4417d6fb700020000       vmovdqa 0x200(%r15), %ymm14
  1447: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1452: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
  1457: c4c15defe5               vpxor %ymm13, %ymm4, %ymm4
  1462: c4417d6fb740020000       vmovdqa 0x240(%r15), %ymm14
  1471: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1476: c4417d6fb780020000       vmovdqa 0x280(%r15), %ymm14
  1485: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1490: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
  1495: c4c145effd               vpxor %ymm13, %ymm7, %ymm7
  1500: c4017d6f4c1e40           vmovdqa 0x40(%r14, %r11), %ymm9
  1507: c44135dbd7               vpand %ymm15, %ymm9, %ymm10
  1512: c4c12571d104             vpsrlw $4, %ymm9, %ymm11
  1518: c44125dbdf               vpand %ymm15, %ymm11, %ymm11
  1523: c4c16defd1               vpxor %ymm9, %ymm2, %ymm2
  1528: c4417d6fb7c0010000       vmovdqa 0x1c0(%r15), %ymm14
  1537: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1542: c4417d6fb700020000       vmovdqa 0x200(%r15), %ymm14
  1551: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1556: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
  1561: c4c155efed               vpxor %ymm13, %ymm5, %ymm5
  1566: c4417d6fb740020000       vmovdqa 0x240(%r15), %ymm14
  1575: c4420d00e2               vpshufb %ymm10, %ymm14, %ymm12
  1580: c4417d6fb780020000       vmovdqa 0x280(%r15), %ymm14
  1589: c4420d00eb               vpshufb %ymm11, %ymm14, %ymm13
  1594: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
  1599: c4413defc5               vpxor %ymm13, %ymm8, %ymm8
  1604: 4d8b5028                 movq 0x28(%r8), %r10
  1608: c4017d6f0c1a             vmovdqa (%r10, %r11), %ymm9
  1614: c4c17defc1               vpxor %ymm9, %ymm0, %ymm0
  1619: c4c165efd9               vpxor %ymm9, %ymm3, %ymm3
  1624: c4c14deff1               vpxor %ymm9, %ymm6, %ymm6
  1629: c4017d6f4c1a20           vmovdqa 0x20(%r10, %r11), %ymm9
  1636: c4c175efc9               vpxor %ymm9, %ymm1, %ymm1
  1641: c4c15defe1               vpxor %ymm9, %ymm4, %ymm4
  1646: c4c145eff9               vpxor %ymm9, %ymm7, %ymm7
  1651: c4017d6f4c1a40           vmovdqa 0x40(%r10, %r11), %ymm9
  1658: c4c16defd1               vpxor %ymm9, %ymm2, %ymm2
  1663: c4c155efe9               vpxor %ymm9, %ymm5, %ymm5
  1668: c4413defc1               vpxor %ymm9, %ymm8, %ymm8
  1673: 4d8b11                   movq (%r9), %r10
  1676: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
  1682: c4817d7f4c1a20           vmovdqa %ymm1, 0x20(%r10, %r11)
  1689: c4817d7f541a40           vmovdqa %ymm2, 0x40(%r10, %r11)
  1696: 4d8b5108                 movq 8(%r9), %r10
  1700: c4817d7f1c1a             vmovdqa %ymm3, (%r10, %r11)
  1706: c4817d7f641a20           vmovdqa %ymm4, 0x20(%r10, %r11)
  1713: c4817d7f6c1a40           vmovdqa %ymm5, 0x40(%r10, %r11)
  1720: 4d8b5110                 movq 0x10(%r9), %r10
  1724: c4817d7f341a             vmovdqa %ymm6, (%r10, %r11)
  1730: c4817d7f7c1a20           vmovdqa %ymm7, 0x20(%r10, %r11)
  1737: c4017d7f441a40           vmovdqa %ymm8, 0x40(%r10, %r11)
  1744: 4983c360                 addq $0x60, %r11
  1748: 4c3bd8                   cmpq %rax, %r11
  1751: 0f8683f9ffff             jbe 0x60
  1757: 4c3bd9                   cmpq %rcx, %r11
  1760: 0f83b2020000             jae 0x998
  1766: 4153                     pushq %r11
  1768: 498dbfc0020000           leaq 0x2c0(%r15), %rdi
  1775: 498db7c0030000           leaq 0x3c0(%r15), %rsi
  1782: 4d8b11                   movq (%r9), %r10
  1785: 4c8b1c24                 movq (%rsp), %r11
  1789: 33db                     xorl %ebx, %ebx
  1791: 498b00                   movq (%r8), %rax
  1794: 420fb61418               movzbl (%rax, %r11), %edx
  1799: 32da                     xorb %dl, %bl
  1801: 498b4008                 movq 8(%r8), %rax
  1805: 420fb61418               movzbl (%rax, %r11), %edx
  1810: 32da                     xorb %dl, %bl
  1812: 498b4010                 movq 0x10(%r8), %rax
  1816: 420fb61418               movzbl (%rax, %r11), %edx
  1821: 32da                     xorb %dl, %bl
  1823: 498b4018                 movq 0x18(%r8), %rax
  1827: 420fb61418               movzbl (%rax, %r11), %edx
  1832: 32da                     xorb %dl, %bl
  1834: 498b4020                 movq 0x20(%r8), %rax
  1838: 420fb61418               movzbl (%rax, %r11), %edx
  1843: 32da                     xorb %dl, %bl
  1845: 498b4028                 movq 0x28(%r8), %rax
  1849: 420fb61418               movzbl (%rax, %r11), %edx
  1854: 32da                     xorb %dl, %bl
  1856: 43881c1a                 movb %bl, (%r10, %r11)
  1860: 33db                     xorl %ebx, %ebx
  1862: 49ffc3                   incq %r11
  1865: 4c3bd9                   cmpq %rcx, %r11
  1868: 0f82adffffff             jb 0x6ff
  1874: 4d8b5108                 movq 8(%r9), %r10
  1878: 4c8b1c24                 movq (%rsp), %r11
  1882: 33db                     xorl %ebx, %ebx
  1884: 498b00                   movq (%r8), %rax
  1887: 420fb61418               movzbl (%rax, %r11), %edx
  1892: 4885d2                   testq %rdx, %rdx
  1895: 0f8422000000             je 0x78f
  1901: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1905: 81c0e6000000             addl $0xe6, %eax
  1911: 81f8ff000000             cmpl $0xff, %eax
  1917: 0f8206000000             jb 0x789
  1923: 81c001ffffff             addl $0xffffff01, %eax
  1929: 0fb60406                 movzbl (%rsi, %rax), %eax
  1933: 32d8                     xorb %al, %bl
  1935: 498b4008                 movq 8(%r8), %rax
  1939: 420fb61418               movzbl (%rax, %r11), %edx
  1944: 4885d2                   testq %rdx, %rdx
  1947: 0f8422000000             je 0x7c3
  1953: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1957: 81c0cd000000             addl $0xcd, %eax
  1963: 81f8ff000000             cmpl $0xff, %eax
  1969: 0f8206000000             jb 0x7bd
  1975: 81c001ffffff             addl $0xffffff01, %eax
  1981: 0fb60406                 movzbl (%rsi, %rax), %eax
  1985: 32d8                     xorb %al, %bl
  1987: 498b4010                 movq 0x10(%r8), %rax
  1991: 420fb61418               movzbl (%rax, %r11), %edx
  1996: 4885d2                   testq %rdx, %rdx
  1999: 0f841f000000             je 0x7f4
  2005: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2009: 83c04b                   addl $0x4b, %eax
  2012: 81f8ff000000             cmpl $0xff, %eax
  2018: 0f8206000000             jb 0x7ee
  2024: 81c001ffffff             addl $0xffffff01, %eax
  2030: 0fb60406                 movzbl (%rsi, %rax), %eax
  2034: 32d8                     xorb %al, %bl
  2036: 498b4018                 movq 0x18(%r8), %rax
  2040: 420fb61418               movzbl (%rax, %r11), %edx
  2045: 4885d2                   testq %rdx, %rdx
  2048: 0f841f000000             je 0x825
  2054: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2058: 83c032                   addl $0x32, %eax
  2061: 81f8ff000000             cmpl $0xff, %eax
  2067: 0f8206000000             jb 0x81f
  2073: 81c001ffffff             addl $0xffffff01, %eax
  2079: 0fb60406                 movzbl (%rsi, %rax), %eax
  2083: 32d8                     xorb %al, %bl
  2085: 498b4020                 movq 0x20(%r8), %rax
  2089: 420fb61418               movzbl (%rax, %r11), %edx
  2094: 4885d2                   testq %rdx, %rdx
  2097: 0f841f000000             je 0x856
  2103: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2107: 83c019                   addl $0x19, %eax
  2110: 81f8ff000000             cmpl $0xff, %eax
  2116: 0f8206000000             jb 0x850
  2122: 81c001ffffff             addl $0xffffff01, %eax
  2128: 0fb60406                 movzbl (%rsi, %rax), %eax
  2132: 32d8                     xorb %al, %bl
  2134: 498b4028                 movq 0x28(%r8), %rax
  2138: 420fb61418               movzbl (%rax, %r11), %edx
  2143: 32da                     xorb %dl, %bl
  2145: 43881c1a                 movb %bl, (%r10, %r11)
  2149: 33db                     xorl %ebx, %ebx
  2151: 49ffc3                   incq %r11
  2154: 4c3bd9                   cmpq %rcx, %r11
  2157: 0f82e9feffff             jb 0x75c
  2163: 4d8b5110                 movq 0x10(%r9), %r10
  2167: 4c8b1c24                 movq (%rsp), %r11
  2171: 33db                     xorl %ebx, %ebx
  2173: 498b00                   movq (%r8), %rax
  2176: 420fb61418               movzbl (%rax, %r11), %edx
  2181: 4885d2                   testq %rdx, %rdx
  2184: 0f8422000000             je 0x8b0
  2190: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2194: 81c0cd000000             addl $0xcd, %eax
  2200: 81f8ff000000             cmpl $0xff, %eax
  2206: 0f8206000000             jb 0x8aa
  2212: 81c001ffffff             addl $0xffffff01, %eax
  2218: 0fb60406                 movzbl (%rsi, %rax), %eax
  2222: 32d8                     xorb %al, %bl
  2224: 498b4008                 movq 8(%r8), %rax
  2228: 420fb61418               movzbl (%rax, %r11), %edx
  2233: 4885d2                   testq %rdx, %rdx
  2236: 0f841f000000             je 0x8e1
  2242: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2246: 83c04b                   addl $0x4b, %eax
  2249: 81f8ff000000             cmpl $0xff, %eax
  2255: 0f8206000000             jb 0x8db
  2261: 81c001ffffff             addl $0xffffff01, %eax
  2267: 0fb60406                 movzbl (%rsi, %rax), %eax
  2271: 32d8                     xorb %al, %bl
  2273: 498b4010                 movq 0x10(%r8), %rax
  2277: 420fb61418               movzbl (%rax, %r11), %edx
  2282: 4885d2                   testq %rdx, %rdx
  2285: 0f841f000000             je 0x912
  2291: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2295: 83c019                   addl $0x19, %eax
  2298: 81f8ff000000             cmpl $0xff, %eax
  2304: 0f8206000000             jb 0x90c
  2310: 81c001ffffff             addl $0xffffff01, %eax
  2316: 0fb60406                 movzbl (%rsi, %rax), %eax
  2320: 32d8                     xorb %al, %bl
  2322: 498b4018                 movq 0x18(%r8), %rax
  2326: 420fb61418               movzbl (%rax, %r11), %edx
  2331: 4885d2                   testq %rdx, %rdx
  2334: 0f8422000000             je 0x946
  2340: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2344: 81c0e6000000             addl $0xe6, %eax
  2350: 81f8ff000000             cmpl $0xff, %eax
  2356: 0f8206000000             jb 0x940
  2362: 81c001ffffff             addl $0xffffff01, %eax
  2368: 0fb60406                 movzbl (%rsi, %rax), %eax
  2372: 32d8                     xorb %al, %bl
  2374: 498b4020                 movq 0x20(%r8), %rax
  2378: 420fb61418               movzbl (%rax, %r11), %edx
  2383: 4885d2                   testq %rdx, %rdx
  2386: 0f841f000000             je 0x977
  2392: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2396: 83c032                   addl $0x32, %eax
  2399: 81f8ff000000             cmpl $0xff, %eax
  2405: 0f8206000000             jb 0x971
  2411: 81c001ffffff             addl $0xffffff01, %eax
  2417: 0fb60406                 movzbl (%rsi, %rax), %eax
  2421: 32d8                     xorb %al, %bl
  2423: 498b4028                 movq 0x28(%r8), %rax
  2427: 420fb61418               movzbl (%rax, %r11), %edx
  2432: 32da                     xorb %dl, %bl
  2434: 43881c1a                 movb %bl, (%r10, %r11)
  2438: 33db                     xorl %ebx, %ebx
  2440: 49ffc3                   incq %r11
  2443: 4c3bd9                   cmpq %rcx, %r11
  2446: 0f82e9feffff             jb 0x87d
  2452: 4883c408                 addq $8, %rsp
  2456: 415e                     popq %r14
  2458: 415d                     popq %r13
  2460: 415c                     popq %r12
  2462: 415f                     popq %r15
  2464: 5b                       popq %rbx
  2465: c5f877                   vzeroupper 
  2468: c3                       retq 
  2469: 0f0b                     ud2 
  2471: 0f0b                     ud2 
  2473: 0f0b                     ud2 
  2475: 0f0b                     ud2 
  2477: 0f0b                     ud2 
  2479: 0f0b                     ud2 
  2481: 0f0b                     ud2 
  2483: 0f0b                     ud2 
  2485: 0f0b                     ud2 
  2487: 0f0b                     ud2 
  2489: 0f0b                     ud2 
  2491: 0f0b                     ud2 
  2493: 0f0b                     ud2 
  2495: cc                       int3 
