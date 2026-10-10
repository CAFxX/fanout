; config_t1_555 target=avx2 twin=xorbit topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=1408 twins=["xorbit"] uops_per_byte=6.03125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3d6d050000           leaq 0x56d(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: c4c17d6f2f               vmovdqa (%r15), %ymm5
    27: 488bc1                   movq %rcx, %rax
    30: 4883e820                 subq $0x20, %rax
    34: 0f823b040000             jb 0x463
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5020                 movq 0x20(%r8), %r10
    68: c4817d6f041a             vmovdqa (%r10, %r11), %ymm0
    74: 4d8b5028                 movq 0x28(%r8), %r10
    78: c4817d6f0c1a             vmovdqa (%r10, %r11), %ymm1
    84: 4d8b5008                 movq 8(%r8), %r10
    88: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
    94: c5e571d200               vpsrlw $0, %ymm2, %ymm3
    99: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   103: 4d8b5018                 movq 0x18(%r8), %r10
   107: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   113: c5dd71d200               vpsrlw $0, %ymm2, %ymm4
   118: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   122: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   126: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   130: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   134: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   139: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   143: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   148: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   152: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   157: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   161: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   166: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   170: 4d8b10                   movq (%r8), %r10
   173: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   179: c5e571d200               vpsrlw $0, %ymm2, %ymm3
   184: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   188: 4d8b5010                 movq 0x10(%r8), %r10
   192: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   198: c5dd71d200               vpsrlw $0, %ymm2, %ymm4
   203: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   207: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   211: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   215: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   219: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   224: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   228: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   233: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   237: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   242: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   246: 4d8b5008                 movq 8(%r8), %r10
   250: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   256: c5e571d201               vpsrlw $1, %ymm2, %ymm3
   261: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   265: 4d8b5018                 movq 0x18(%r8), %r10
   269: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   275: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
   280: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   284: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   288: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   292: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   297: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   301: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   306: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   310: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   315: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   319: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   324: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   328: 4d8b10                   movq (%r8), %r10
   331: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   337: c5e571d201               vpsrlw $1, %ymm2, %ymm3
   342: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   346: 4d8b5010                 movq 0x10(%r8), %r10
   350: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   356: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
   361: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   365: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   369: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   373: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   378: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   382: 4d8b10                   movq (%r8), %r10
   385: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   391: c5e571d202               vpsrlw $2, %ymm2, %ymm3
   396: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   400: 4d8b5010                 movq 0x10(%r8), %r10
   404: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   410: c5dd71d202               vpsrlw $2, %ymm2, %ymm4
   415: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   419: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   423: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   428: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   432: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   437: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   441: 4d8b10                   movq (%r8), %r10
   444: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   450: c5e571d203               vpsrlw $3, %ymm2, %ymm3
   455: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   459: 4d8b5010                 movq 0x10(%r8), %r10
   463: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   469: c5dd71d203               vpsrlw $3, %ymm2, %ymm4
   474: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   478: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   482: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   487: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   491: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   496: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   500: 4d8b10                   movq (%r8), %r10
   503: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   509: c5e571d204               vpsrlw $4, %ymm2, %ymm3
   514: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   518: 4d8b5010                 movq 0x10(%r8), %r10
   522: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   528: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   533: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   537: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   541: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   546: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   550: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   555: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   559: 4d8b10                   movq (%r8), %r10
   562: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   568: c5e571d205               vpsrlw $5, %ymm2, %ymm3
   573: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   577: 4d8b5010                 movq 0x10(%r8), %r10
   581: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   587: c5dd71d205               vpsrlw $5, %ymm2, %ymm4
   592: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   596: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   600: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   605: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   609: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   614: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   618: 4d8b10                   movq (%r8), %r10
   621: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   627: c5e571d206               vpsrlw $6, %ymm2, %ymm3
   632: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   636: 4d8b5010                 movq 0x10(%r8), %r10
   640: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   646: c5dd71d206               vpsrlw $6, %ymm2, %ymm4
   651: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   655: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   659: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   664: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   668: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   673: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   677: 4d8b10                   movq (%r8), %r10
   680: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   686: c5e571d207               vpsrlw $7, %ymm2, %ymm3
   691: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   695: 4d8b5010                 movq 0x10(%r8), %r10
   699: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   705: c5dd71d207               vpsrlw $7, %ymm2, %ymm4
   710: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   714: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   718: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   723: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   727: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   732: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   736: 4d8b5008                 movq 8(%r8), %r10
   740: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   746: c5e571d202               vpsrlw $2, %ymm2, %ymm3
   751: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   755: 4d8b5018                 movq 0x18(%r8), %r10
   759: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   765: c5dd71d202               vpsrlw $2, %ymm2, %ymm4
   770: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   774: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   778: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   782: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   787: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   791: 4d8b5008                 movq 8(%r8), %r10
   795: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   801: c5e571d203               vpsrlw $3, %ymm2, %ymm3
   806: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   810: 4d8b5018                 movq 0x18(%r8), %r10
   814: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   820: c5dd71d203               vpsrlw $3, %ymm2, %ymm4
   825: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   829: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   833: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   838: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   842: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   847: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   851: 4d8b5008                 movq 8(%r8), %r10
   855: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   861: c5e571d204               vpsrlw $4, %ymm2, %ymm3
   866: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   870: 4d8b5018                 movq 0x18(%r8), %r10
   874: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   880: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   885: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   889: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   893: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   898: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   902: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   907: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   911: 4d8b5008                 movq 8(%r8), %r10
   915: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   921: c5e571d205               vpsrlw $5, %ymm2, %ymm3
   926: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   930: 4d8b5018                 movq 0x18(%r8), %r10
   934: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   940: c5dd71d205               vpsrlw $5, %ymm2, %ymm4
   945: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   949: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   953: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   958: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   962: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   967: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   971: 4d8b5008                 movq 8(%r8), %r10
   975: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   981: c5e571d206               vpsrlw $6, %ymm2, %ymm3
   986: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   990: 4d8b5018                 movq 0x18(%r8), %r10
   994: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
  1000: c5dd71d206               vpsrlw $6, %ymm2, %ymm4
  1005: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1009: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1013: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1018: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1022: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1027: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1031: 4d8b5008                 movq 8(%r8), %r10
  1035: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
  1041: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  1046: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1050: 4d8b5018                 movq 0x18(%r8), %r10
  1054: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
  1060: c5dd71d207               vpsrlw $7, %ymm2, %ymm4
  1065: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1069: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1073: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1078: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1082: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1087: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1091: 4d8b11                   movq (%r9), %r10
  1094: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
  1100: 4d8b5108                 movq 8(%r9), %r10
  1104: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
  1110: 4983c320                 addq $0x20, %r11
  1114: 4c3bd8                   cmpq %rax, %r11
  1117: 0f86ddfbffff             jbe 0x40
  1123: 4c3bd9                   cmpq %rcx, %r11
  1126: 0f83fa000000             jae 0x566
  1132: 4153                     pushq %r11
  1134: 498d7f40                 leaq 0x40(%r15), %rdi
  1138: 498db740010000           leaq 0x140(%r15), %rsi
  1145: 4d8b11                   movq (%r9), %r10
  1148: 4c8b1c24                 movq (%rsp), %r11
  1152: 498b4020                 movq 0x20(%r8), %rax
  1156: 420fb61c18               movzbl (%rax, %r11), %ebx
  1161: 498b00                   movq (%r8), %rax
  1164: 420fb61418               movzbl (%rax, %r11), %edx
  1169: 498b4010                 movq 0x10(%r8), %rax
  1173: 420fb60418               movzbl (%rax, %r11), %eax
  1178: 32d0                     xorb %al, %dl
  1180: 32da                     xorb %dl, %bl
  1182: 498b4008                 movq 8(%r8), %rax
  1186: 420fb61418               movzbl (%rax, %r11), %edx
  1191: 498b4018                 movq 0x18(%r8), %rax
  1195: 420fb60418               movzbl (%rax, %r11), %eax
  1200: 32d0                     xorb %al, %dl
  1202: 32da                     xorb %dl, %bl
  1204: 43881c1a                 movb %bl, (%r10, %r11)
  1208: 49ffc3                   incq %r11
  1211: 4c3bd9                   cmpq %rcx, %r11
  1214: 0f82bcffffff             jb 0x480
  1220: 4d8b5108                 movq 8(%r9), %r10
  1224: 4c8b1c24                 movq (%rsp), %r11
  1228: 498b4028                 movq 0x28(%r8), %rax
  1232: 420fb61c18               movzbl (%rax, %r11), %ebx
  1237: 498b00                   movq (%r8), %rax
  1240: 420fb61418               movzbl (%rax, %r11), %edx
  1245: 498b4010                 movq 0x10(%r8), %rax
  1249: 420fb60418               movzbl (%rax, %r11), %eax
  1254: 32d0                     xorb %al, %dl
  1256: 4885d2                   testq %rdx, %rdx
  1259: 0f8422000000             je 0x513
  1265: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1269: 81c0e6000000             addl $0xe6, %eax
  1275: 81f8ff000000             cmpl $0xff, %eax
  1281: 0f8206000000             jb 0x50d
  1287: 81c001ffffff             addl $0xffffff01, %eax
  1293: 0fb60406                 movzbl (%rsi, %rax), %eax
  1297: 32d8                     xorb %al, %bl
  1299: 498b4008                 movq 8(%r8), %rax
  1303: 420fb61418               movzbl (%rax, %r11), %edx
  1308: 498b4018                 movq 0x18(%r8), %rax
  1312: 420fb60418               movzbl (%rax, %r11), %eax
  1317: 32d0                     xorb %al, %dl
  1319: 4885d2                   testq %rdx, %rdx
  1322: 0f8422000000             je 0x552
  1328: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1332: 81c0cd000000             addl $0xcd, %eax
  1338: 81f8ff000000             cmpl $0xff, %eax
  1344: 0f8206000000             jb 0x54c
  1350: 81c001ffffff             addl $0xffffff01, %eax
  1356: 0fb60406                 movzbl (%rsi, %rax), %eax
  1360: 32d8                     xorb %al, %bl
  1362: 43881c1a                 movb %bl, (%r10, %r11)
  1366: 49ffc3                   incq %r11
  1369: 4c3bd9                   cmpq %rcx, %r11
  1372: 0f826affffff             jb 0x4cc
  1378: 4883c408                 addq $8, %rsp
  1382: 415f                     popq %r15
  1384: 5b                       popq %rbx
  1385: c5f877                   vzeroupper 
  1388: c3                       retq 
  1389: 0f0b                     ud2 
  1391: 0f0b                     ud2 
  1393: 0f0b                     ud2 
  1395: 0f0b                     ud2 
  1397: 0f0b                     ud2 
  1399: 0f0b                     ud2 
  1401: 0f0b                     ud2 
  1403: 0f0b                     ud2 
  1405: 0f0b                     ud2 
  1407: cc                       int3 
