; config_t1_841 target=sse twin=xorbit topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=1664 twins=["xorbit"] uops_per_byte=16.25
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3d6d060000           leaq 0x66d(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: 66410f6f2f               movdqa (%r15), %xmm5
    27: 488bc1                   movq %rcx, %rax
    30: 4883e810                 subq $0x10, %rax
    34: 0f8247050000             jb 0x56f
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5020                 movq 0x20(%r8), %r10
    68: 66430f6f041a             movdqa (%r10, %r11), %xmm0
    74: 4d8b5028                 movq 0x28(%r8), %r10
    78: 66430f6f0c1a             movdqa (%r10, %r11), %xmm1
    84: 4d8b5008                 movq 8(%r8), %r10
    88: 66430f6f141a             movdqa (%r10, %r11), %xmm2
    94: 660f6fda                 movdqa %xmm2, %xmm3
    98: 660f71d300               psrlw $0, %xmm3
   103: 660fdbdd                 pand %xmm5, %xmm3
   107: 4d8b5018                 movq 0x18(%r8), %r10
   111: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   117: 660f6fe2                 movdqa %xmm2, %xmm4
   121: 660f71d400               psrlw $0, %xmm4
   126: 660fdbe5                 pand %xmm5, %xmm4
   130: 660fefdc                 pxor %xmm4, %xmm3
   134: 660fefc3                 pxor %xmm3, %xmm0
   138: 660fefcb                 pxor %xmm3, %xmm1
   142: 660f6fe3                 movdqa %xmm3, %xmm4
   146: 660f71f401               psllw $1, %xmm4
   151: 660fefcc                 pxor %xmm4, %xmm1
   155: 660f6fe3                 movdqa %xmm3, %xmm4
   159: 660f71f403               psllw $3, %xmm4
   164: 660fefcc                 pxor %xmm4, %xmm1
   168: 660f6fe3                 movdqa %xmm3, %xmm4
   172: 660f71f406               psllw $6, %xmm4
   177: 660fefcc                 pxor %xmm4, %xmm1
   181: 660f6fe3                 movdqa %xmm3, %xmm4
   185: 660f71f407               psllw $7, %xmm4
   190: 660fefcc                 pxor %xmm4, %xmm1
   194: 4d8b10                   movq (%r8), %r10
   197: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   203: 660f6fda                 movdqa %xmm2, %xmm3
   207: 660f71d300               psrlw $0, %xmm3
   212: 660fdbdd                 pand %xmm5, %xmm3
   216: 4d8b5010                 movq 0x10(%r8), %r10
   220: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   226: 660f6fe2                 movdqa %xmm2, %xmm4
   230: 660f71d400               psrlw $0, %xmm4
   235: 660fdbe5                 pand %xmm5, %xmm4
   239: 660fefdc                 pxor %xmm4, %xmm3
   243: 660fefc3                 pxor %xmm3, %xmm0
   247: 660fefcb                 pxor %xmm3, %xmm1
   251: 660f6fe3                 movdqa %xmm3, %xmm4
   255: 660f71f402               psllw $2, %xmm4
   260: 660fefcc                 pxor %xmm4, %xmm1
   264: 660f6fe3                 movdqa %xmm3, %xmm4
   268: 660f71f403               psllw $3, %xmm4
   273: 660fefcc                 pxor %xmm4, %xmm1
   277: 660f6fe3                 movdqa %xmm3, %xmm4
   281: 660f71f407               psllw $7, %xmm4
   286: 660fefcc                 pxor %xmm4, %xmm1
   290: 4d8b5008                 movq 8(%r8), %r10
   294: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   300: 660f6fda                 movdqa %xmm2, %xmm3
   304: 660f71d301               psrlw $1, %xmm3
   309: 660fdbdd                 pand %xmm5, %xmm3
   313: 4d8b5018                 movq 0x18(%r8), %r10
   317: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   323: 660f6fe2                 movdqa %xmm2, %xmm4
   327: 660f71d401               psrlw $1, %xmm4
   332: 660fdbe5                 pand %xmm5, %xmm4
   336: 660fefdc                 pxor %xmm4, %xmm3
   340: 660fefcb                 pxor %xmm3, %xmm1
   344: 660f6fe3                 movdqa %xmm3, %xmm4
   348: 660f71f401               psllw $1, %xmm4
   353: 660fefc4                 pxor %xmm4, %xmm0
   357: 660f6fe3                 movdqa %xmm3, %xmm4
   361: 660f71f402               psllw $2, %xmm4
   366: 660fefcc                 pxor %xmm4, %xmm1
   370: 660f6fe3                 movdqa %xmm3, %xmm4
   374: 660f71f403               psllw $3, %xmm4
   379: 660fefcc                 pxor %xmm4, %xmm1
   383: 660f6fe3                 movdqa %xmm3, %xmm4
   387: 660f71f407               psllw $7, %xmm4
   392: 660fefcc                 pxor %xmm4, %xmm1
   396: 4d8b10                   movq (%r8), %r10
   399: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   405: 660f6fda                 movdqa %xmm2, %xmm3
   409: 660f71d301               psrlw $1, %xmm3
   414: 660fdbdd                 pand %xmm5, %xmm3
   418: 4d8b5010                 movq 0x10(%r8), %r10
   422: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   428: 660f6fe2                 movdqa %xmm2, %xmm4
   432: 660f71d401               psrlw $1, %xmm4
   437: 660fdbe5                 pand %xmm5, %xmm4
   441: 660fefdc                 pxor %xmm4, %xmm3
   445: 660fefcb                 pxor %xmm3, %xmm1
   449: 660f6fe3                 movdqa %xmm3, %xmm4
   453: 660f71f401               psllw $1, %xmm4
   458: 660fefc4                 pxor %xmm4, %xmm0
   462: 4d8b10                   movq (%r8), %r10
   465: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   471: 660f6fda                 movdqa %xmm2, %xmm3
   475: 660f71d302               psrlw $2, %xmm3
   480: 660fdbdd                 pand %xmm5, %xmm3
   484: 4d8b5010                 movq 0x10(%r8), %r10
   488: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   494: 660f6fe2                 movdqa %xmm2, %xmm4
   498: 660f71d402               psrlw $2, %xmm4
   503: 660fdbe5                 pand %xmm5, %xmm4
   507: 660fefdc                 pxor %xmm4, %xmm3
   511: 660f6fe3                 movdqa %xmm3, %xmm4
   515: 660f71f401               psllw $1, %xmm4
   520: 660fefcc                 pxor %xmm4, %xmm1
   524: 660f6fe3                 movdqa %xmm3, %xmm4
   528: 660f71f402               psllw $2, %xmm4
   533: 660fefc4                 pxor %xmm4, %xmm0
   537: 4d8b10                   movq (%r8), %r10
   540: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   546: 660f6fda                 movdqa %xmm2, %xmm3
   550: 660f71d303               psrlw $3, %xmm3
   555: 660fdbdd                 pand %xmm5, %xmm3
   559: 4d8b5010                 movq 0x10(%r8), %r10
   563: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   569: 660f6fe2                 movdqa %xmm2, %xmm4
   573: 660f71d403               psrlw $3, %xmm4
   578: 660fdbe5                 pand %xmm5, %xmm4
   582: 660fefdc                 pxor %xmm4, %xmm3
   586: 660f6fe3                 movdqa %xmm3, %xmm4
   590: 660f71f402               psllw $2, %xmm4
   595: 660fefcc                 pxor %xmm4, %xmm1
   599: 660f6fe3                 movdqa %xmm3, %xmm4
   603: 660f71f403               psllw $3, %xmm4
   608: 660fefc4                 pxor %xmm4, %xmm0
   612: 4d8b10                   movq (%r8), %r10
   615: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   621: 660f6fda                 movdqa %xmm2, %xmm3
   625: 660f71d304               psrlw $4, %xmm3
   630: 660fdbdd                 pand %xmm5, %xmm3
   634: 4d8b5010                 movq 0x10(%r8), %r10
   638: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   644: 660f6fe2                 movdqa %xmm2, %xmm4
   648: 660f71d404               psrlw $4, %xmm4
   653: 660fdbe5                 pand %xmm5, %xmm4
   657: 660fefdc                 pxor %xmm4, %xmm3
   661: 660f6fe3                 movdqa %xmm3, %xmm4
   665: 660f71f403               psllw $3, %xmm4
   670: 660fefcc                 pxor %xmm4, %xmm1
   674: 660f6fe3                 movdqa %xmm3, %xmm4
   678: 660f71f404               psllw $4, %xmm4
   683: 660fefc4                 pxor %xmm4, %xmm0
   687: 4d8b10                   movq (%r8), %r10
   690: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   696: 660f6fda                 movdqa %xmm2, %xmm3
   700: 660f71d305               psrlw $5, %xmm3
   705: 660fdbdd                 pand %xmm5, %xmm3
   709: 4d8b5010                 movq 0x10(%r8), %r10
   713: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   719: 660f6fe2                 movdqa %xmm2, %xmm4
   723: 660f71d405               psrlw $5, %xmm4
   728: 660fdbe5                 pand %xmm5, %xmm4
   732: 660fefdc                 pxor %xmm4, %xmm3
   736: 660f6fe3                 movdqa %xmm3, %xmm4
   740: 660f71f404               psllw $4, %xmm4
   745: 660fefcc                 pxor %xmm4, %xmm1
   749: 660f6fe3                 movdqa %xmm3, %xmm4
   753: 660f71f405               psllw $5, %xmm4
   758: 660fefc4                 pxor %xmm4, %xmm0
   762: 4d8b10                   movq (%r8), %r10
   765: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   771: 660f6fda                 movdqa %xmm2, %xmm3
   775: 660f71d306               psrlw $6, %xmm3
   780: 660fdbdd                 pand %xmm5, %xmm3
   784: 4d8b5010                 movq 0x10(%r8), %r10
   788: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   794: 660f6fe2                 movdqa %xmm2, %xmm4
   798: 660f71d406               psrlw $6, %xmm4
   803: 660fdbe5                 pand %xmm5, %xmm4
   807: 660fefdc                 pxor %xmm4, %xmm3
   811: 660f6fe3                 movdqa %xmm3, %xmm4
   815: 660f71f405               psllw $5, %xmm4
   820: 660fefcc                 pxor %xmm4, %xmm1
   824: 660f6fe3                 movdqa %xmm3, %xmm4
   828: 660f71f406               psllw $6, %xmm4
   833: 660fefc4                 pxor %xmm4, %xmm0
   837: 4d8b10                   movq (%r8), %r10
   840: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   846: 660f6fda                 movdqa %xmm2, %xmm3
   850: 660f71d307               psrlw $7, %xmm3
   855: 660fdbdd                 pand %xmm5, %xmm3
   859: 4d8b5010                 movq 0x10(%r8), %r10
   863: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   869: 660f6fe2                 movdqa %xmm2, %xmm4
   873: 660f71d407               psrlw $7, %xmm4
   878: 660fdbe5                 pand %xmm5, %xmm4
   882: 660fefdc                 pxor %xmm4, %xmm3
   886: 660f6fe3                 movdqa %xmm3, %xmm4
   890: 660f71f406               psllw $6, %xmm4
   895: 660fefcc                 pxor %xmm4, %xmm1
   899: 660f6fe3                 movdqa %xmm3, %xmm4
   903: 660f71f407               psllw $7, %xmm4
   908: 660fefc4                 pxor %xmm4, %xmm0
   912: 4d8b5008                 movq 8(%r8), %r10
   916: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   922: 660f6fda                 movdqa %xmm2, %xmm3
   926: 660f71d302               psrlw $2, %xmm3
   931: 660fdbdd                 pand %xmm5, %xmm3
   935: 4d8b5018                 movq 0x18(%r8), %r10
   939: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   945: 660f6fe2                 movdqa %xmm2, %xmm4
   949: 660f71d402               psrlw $2, %xmm4
   954: 660fdbe5                 pand %xmm5, %xmm4
   958: 660fefdc                 pxor %xmm4, %xmm3
   962: 660fefcb                 pxor %xmm3, %xmm1
   966: 660f6fe3                 movdqa %xmm3, %xmm4
   970: 660f71f402               psllw $2, %xmm4
   975: 660fefc4                 pxor %xmm4, %xmm0
   979: 4d8b5008                 movq 8(%r8), %r10
   983: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   989: 660f6fda                 movdqa %xmm2, %xmm3
   993: 660f71d303               psrlw $3, %xmm3
   998: 660fdbdd                 pand %xmm5, %xmm3
  1002: 4d8b5018                 movq 0x18(%r8), %r10
  1006: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1012: 660f6fe2                 movdqa %xmm2, %xmm4
  1016: 660f71d403               psrlw $3, %xmm4
  1021: 660fdbe5                 pand %xmm5, %xmm4
  1025: 660fefdc                 pxor %xmm4, %xmm3
  1029: 660f6fe3                 movdqa %xmm3, %xmm4
  1033: 660f71f401               psllw $1, %xmm4
  1038: 660fefcc                 pxor %xmm4, %xmm1
  1042: 660f6fe3                 movdqa %xmm3, %xmm4
  1046: 660f71f403               psllw $3, %xmm4
  1051: 660fefc4                 pxor %xmm4, %xmm0
  1055: 4d8b5008                 movq 8(%r8), %r10
  1059: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1065: 660f6fda                 movdqa %xmm2, %xmm3
  1069: 660f71d304               psrlw $4, %xmm3
  1074: 660fdbdd                 pand %xmm5, %xmm3
  1078: 4d8b5018                 movq 0x18(%r8), %r10
  1082: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1088: 660f6fe2                 movdqa %xmm2, %xmm4
  1092: 660f71d404               psrlw $4, %xmm4
  1097: 660fdbe5                 pand %xmm5, %xmm4
  1101: 660fefdc                 pxor %xmm4, %xmm3
  1105: 660f6fe3                 movdqa %xmm3, %xmm4
  1109: 660f71f402               psllw $2, %xmm4
  1114: 660fefcc                 pxor %xmm4, %xmm1
  1118: 660f6fe3                 movdqa %xmm3, %xmm4
  1122: 660f71f404               psllw $4, %xmm4
  1127: 660fefc4                 pxor %xmm4, %xmm0
  1131: 4d8b5008                 movq 8(%r8), %r10
  1135: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1141: 660f6fda                 movdqa %xmm2, %xmm3
  1145: 660f71d305               psrlw $5, %xmm3
  1150: 660fdbdd                 pand %xmm5, %xmm3
  1154: 4d8b5018                 movq 0x18(%r8), %r10
  1158: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1164: 660f6fe2                 movdqa %xmm2, %xmm4
  1168: 660f71d405               psrlw $5, %xmm4
  1173: 660fdbe5                 pand %xmm5, %xmm4
  1177: 660fefdc                 pxor %xmm4, %xmm3
  1181: 660f6fe3                 movdqa %xmm3, %xmm4
  1185: 660f71f403               psllw $3, %xmm4
  1190: 660fefcc                 pxor %xmm4, %xmm1
  1194: 660f6fe3                 movdqa %xmm3, %xmm4
  1198: 660f71f405               psllw $5, %xmm4
  1203: 660fefc4                 pxor %xmm4, %xmm0
  1207: 4d8b5008                 movq 8(%r8), %r10
  1211: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1217: 660f6fda                 movdqa %xmm2, %xmm3
  1221: 660f71d306               psrlw $6, %xmm3
  1226: 660fdbdd                 pand %xmm5, %xmm3
  1230: 4d8b5018                 movq 0x18(%r8), %r10
  1234: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1240: 660f6fe2                 movdqa %xmm2, %xmm4
  1244: 660f71d406               psrlw $6, %xmm4
  1249: 660fdbe5                 pand %xmm5, %xmm4
  1253: 660fefdc                 pxor %xmm4, %xmm3
  1257: 660f6fe3                 movdqa %xmm3, %xmm4
  1261: 660f71f404               psllw $4, %xmm4
  1266: 660fefcc                 pxor %xmm4, %xmm1
  1270: 660f6fe3                 movdqa %xmm3, %xmm4
  1274: 660f71f406               psllw $6, %xmm4
  1279: 660fefc4                 pxor %xmm4, %xmm0
  1283: 4d8b5008                 movq 8(%r8), %r10
  1287: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1293: 660f6fda                 movdqa %xmm2, %xmm3
  1297: 660f71d307               psrlw $7, %xmm3
  1302: 660fdbdd                 pand %xmm5, %xmm3
  1306: 4d8b5018                 movq 0x18(%r8), %r10
  1310: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1316: 660f6fe2                 movdqa %xmm2, %xmm4
  1320: 660f71d407               psrlw $7, %xmm4
  1325: 660fdbe5                 pand %xmm5, %xmm4
  1329: 660fefdc                 pxor %xmm4, %xmm3
  1333: 660f6fe3                 movdqa %xmm3, %xmm4
  1337: 660f71f405               psllw $5, %xmm4
  1342: 660fefcc                 pxor %xmm4, %xmm1
  1346: 660f6fe3                 movdqa %xmm3, %xmm4
  1350: 660f71f407               psllw $7, %xmm4
  1355: 660fefc4                 pxor %xmm4, %xmm0
  1359: 4d8b11                   movq (%r9), %r10
  1362: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
  1368: 4d8b5108                 movq 8(%r9), %r10
  1372: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
  1378: 4983c310                 addq $0x10, %r11
  1382: 4c3bd8                   cmpq %rax, %r11
  1385: 0f86d1faffff             jbe 0x40
  1391: 4c3bd9                   cmpq %rcx, %r11
  1394: 0f83fa000000             jae 0x672
  1400: 4153                     pushq %r11
  1402: 498d7f40                 leaq 0x40(%r15), %rdi
  1406: 498db740010000           leaq 0x140(%r15), %rsi
  1413: 4d8b11                   movq (%r9), %r10
  1416: 4c8b1c24                 movq (%rsp), %r11
  1420: 498b4020                 movq 0x20(%r8), %rax
  1424: 420fb61c18               movzbl (%rax, %r11), %ebx
  1429: 498b00                   movq (%r8), %rax
  1432: 420fb61418               movzbl (%rax, %r11), %edx
  1437: 498b4010                 movq 0x10(%r8), %rax
  1441: 420fb60418               movzbl (%rax, %r11), %eax
  1446: 32d0                     xorb %al, %dl
  1448: 32da                     xorb %dl, %bl
  1450: 498b4008                 movq 8(%r8), %rax
  1454: 420fb61418               movzbl (%rax, %r11), %edx
  1459: 498b4018                 movq 0x18(%r8), %rax
  1463: 420fb60418               movzbl (%rax, %r11), %eax
  1468: 32d0                     xorb %al, %dl
  1470: 32da                     xorb %dl, %bl
  1472: 43881c1a                 movb %bl, (%r10, %r11)
  1476: 49ffc3                   incq %r11
  1479: 4c3bd9                   cmpq %rcx, %r11
  1482: 0f82bcffffff             jb 0x58c
  1488: 4d8b5108                 movq 8(%r9), %r10
  1492: 4c8b1c24                 movq (%rsp), %r11
  1496: 498b4028                 movq 0x28(%r8), %rax
  1500: 420fb61c18               movzbl (%rax, %r11), %ebx
  1505: 498b00                   movq (%r8), %rax
  1508: 420fb61418               movzbl (%rax, %r11), %edx
  1513: 498b4010                 movq 0x10(%r8), %rax
  1517: 420fb60418               movzbl (%rax, %r11), %eax
  1522: 32d0                     xorb %al, %dl
  1524: 4885d2                   testq %rdx, %rdx
  1527: 0f8422000000             je 0x61f
  1533: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1537: 81c0e6000000             addl $0xe6, %eax
  1543: 81f8ff000000             cmpl $0xff, %eax
  1549: 0f8206000000             jb 0x619
  1555: 81c001ffffff             addl $0xffffff01, %eax
  1561: 0fb60406                 movzbl (%rsi, %rax), %eax
  1565: 32d8                     xorb %al, %bl
  1567: 498b4008                 movq 8(%r8), %rax
  1571: 420fb61418               movzbl (%rax, %r11), %edx
  1576: 498b4018                 movq 0x18(%r8), %rax
  1580: 420fb60418               movzbl (%rax, %r11), %eax
  1585: 32d0                     xorb %al, %dl
  1587: 4885d2                   testq %rdx, %rdx
  1590: 0f8422000000             je 0x65e
  1596: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1600: 81c0cd000000             addl $0xcd, %eax
  1606: 81f8ff000000             cmpl $0xff, %eax
  1612: 0f8206000000             jb 0x658
  1618: 81c001ffffff             addl $0xffffff01, %eax
  1624: 0fb60406                 movzbl (%rsi, %rax), %eax
  1628: 32d8                     xorb %al, %bl
  1630: 43881c1a                 movb %bl, (%r10, %r11)
  1634: 49ffc3                   incq %r11
  1637: 4c3bd9                   cmpq %rcx, %r11
  1640: 0f826affffff             jb 0x5d8
  1646: 4883c408                 addq $8, %rsp
  1650: 415f                     popq %r15
  1652: 5b                       popq %rbx
  1653: c3                       retq 
  1654: 0f0b                     ud2 
  1656: 0f0b                     ud2 
  1658: 0f0b                     ud2 
  1660: 0f0b                     ud2 
  1662: 0f0b                     ud2 
