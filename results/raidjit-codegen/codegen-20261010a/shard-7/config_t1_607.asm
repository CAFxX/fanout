; config_t1_607 target=sse twin=xorbit topo=raid6(8)/enc op=encode U=1
; code_size=3712 twins=["xorbit"] uops_per_byte=42.0625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d670e0000           leaq 0xe67(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66410f6f2f               movdqa (%r15), %xmm5
    52: 488bc1                   movq %rcx, %rax
    55: 4883e810                 subq $0x10, %rax
    59: 0f82e10b0000             jb 0xc22
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   102: 660f6fda                 movdqa %xmm2, %xmm3
   106: 660f71d300               psrlw $0, %xmm3
   111: 660fdbdd                 pand %xmm5, %xmm3
   115: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
   121: 660f6fe2                 movdqa %xmm2, %xmm4
   125: 660f71d400               psrlw $0, %xmm4
   130: 660fdbe5                 pand %xmm5, %xmm4
   134: 660fefdc                 pxor %xmm4, %xmm3
   138: 660f6fc3                 movdqa %xmm3, %xmm0
   142: 660f6fcb                 movdqa %xmm3, %xmm1
   146: 660f6fe3                 movdqa %xmm3, %xmm4
   150: 660f71f403               psllw $3, %xmm4
   155: 660fefcc                 pxor %xmm4, %xmm1
   159: 660f6fe3                 movdqa %xmm3, %xmm4
   163: 660f71f407               psllw $7, %xmm4
   168: 660fefcc                 pxor %xmm4, %xmm1
   172: 660f6fda                 movdqa %xmm2, %xmm3
   176: 660f71d301               psrlw $1, %xmm3
   181: 660fdbdd                 pand %xmm5, %xmm3
   185: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
   192: 660f6fe2                 movdqa %xmm2, %xmm4
   196: 660f71d401               psrlw $1, %xmm4
   201: 660fdbe5                 pand %xmm5, %xmm4
   205: 660fefdc                 pxor %xmm4, %xmm3
   209: 660f6fe3                 movdqa %xmm3, %xmm4
   213: 660f71f401               psllw $1, %xmm4
   218: 660fefc4                 pxor %xmm4, %xmm0
   222: 660f6fe3                 movdqa %xmm3, %xmm4
   226: 660f71f402               psllw $2, %xmm4
   231: 660fefcc                 pxor %xmm4, %xmm1
   235: 660f6fe3                 movdqa %xmm3, %xmm4
   239: 660f71f403               psllw $3, %xmm4
   244: 660fefcc                 pxor %xmm4, %xmm1
   248: 660f6fe3                 movdqa %xmm3, %xmm4
   252: 660f71f407               psllw $7, %xmm4
   257: 660fefcc                 pxor %xmm4, %xmm1
   261: 66430f6f141e             movdqa (%r14, %r11), %xmm2
   267: 660f6fda                 movdqa %xmm2, %xmm3
   271: 660f71d305               psrlw $5, %xmm3
   276: 660fdbdd                 pand %xmm5, %xmm3
   280: 4d8b5028                 movq 0x28(%r8), %r10
   284: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   290: 660f6fe2                 movdqa %xmm2, %xmm4
   294: 660f71d406               psrlw $6, %xmm4
   299: 660fdbe5                 pand %xmm5, %xmm4
   303: 660fefdc                 pxor %xmm4, %xmm3
   307: 660fefcb                 pxor %xmm3, %xmm1
   311: 660f6fe3                 movdqa %xmm3, %xmm4
   315: 660f71f401               psllw $1, %xmm4
   320: 660fefcc                 pxor %xmm4, %xmm1
   324: 660f6fe3                 movdqa %xmm3, %xmm4
   328: 660f71f403               psllw $3, %xmm4
   333: 660fefcc                 pxor %xmm4, %xmm1
   337: 660f6fe3                 movdqa %xmm3, %xmm4
   341: 660f71f404               psllw $4, %xmm4
   346: 660fefcc                 pxor %xmm4, %xmm1
   350: 66430f6f141e             movdqa (%r14, %r11), %xmm2
   356: 660f6fda                 movdqa %xmm2, %xmm3
   360: 660f71d306               psrlw $6, %xmm3
   365: 660fdbdd                 pand %xmm5, %xmm3
   369: 4d8b5028                 movq 0x28(%r8), %r10
   373: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   379: 660f6fe2                 movdqa %xmm2, %xmm4
   383: 660f71d407               psrlw $7, %xmm4
   388: 660fdbe5                 pand %xmm5, %xmm4
   392: 660fefdc                 pxor %xmm4, %xmm3
   396: 660f6fe3                 movdqa %xmm3, %xmm4
   400: 660f71f401               psllw $1, %xmm4
   405: 660fefcc                 pxor %xmm4, %xmm1
   409: 660f6fe3                 movdqa %xmm3, %xmm4
   413: 660f71f402               psllw $2, %xmm4
   418: 660fefcc                 pxor %xmm4, %xmm1
   422: 660f6fe3                 movdqa %xmm3, %xmm4
   426: 660f71f404               psllw $4, %xmm4
   431: 660fefcc                 pxor %xmm4, %xmm1
   435: 660f6fe3                 movdqa %xmm3, %xmm4
   439: 660f71f405               psllw $5, %xmm4
   444: 660fefcc                 pxor %xmm4, %xmm1
   448: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
   455: 660f6fda                 movdqa %xmm2, %xmm3
   459: 660f71d307               psrlw $7, %xmm3
   464: 660fdbdd                 pand %xmm5, %xmm3
   468: 66430f6f141e             movdqa (%r14, %r11), %xmm2
   474: 660f6fe2                 movdqa %xmm2, %xmm4
   478: 660f71d407               psrlw $7, %xmm4
   483: 660fdbe5                 pand %xmm5, %xmm4
   487: 660fefdc                 pxor %xmm4, %xmm3
   491: 660f6fe3                 movdqa %xmm3, %xmm4
   495: 660f71f405               psllw $5, %xmm4
   500: 660fefcc                 pxor %xmm4, %xmm1
   504: 660f6fe3                 movdqa %xmm3, %xmm4
   508: 660f71f406               psllw $6, %xmm4
   513: 660fefcc                 pxor %xmm4, %xmm1
   517: 660f6fe3                 movdqa %xmm3, %xmm4
   521: 660f71f407               psllw $7, %xmm4
   526: 660fefc4                 pxor %xmm4, %xmm0
   530: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   536: 660f6fda                 movdqa %xmm2, %xmm3
   540: 660f71d301               psrlw $1, %xmm3
   545: 660fdbdd                 pand %xmm5, %xmm3
   549: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   555: 660f6fe2                 movdqa %xmm2, %xmm4
   559: 660f71d401               psrlw $1, %xmm4
   564: 660fdbe5                 pand %xmm5, %xmm4
   568: 660fefdc                 pxor %xmm4, %xmm3
   572: 660fefcb                 pxor %xmm3, %xmm1
   576: 660f6fe3                 movdqa %xmm3, %xmm4
   580: 660f71f401               psllw $1, %xmm4
   585: 660fefc4                 pxor %xmm4, %xmm0
   589: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   595: 660f6fda                 movdqa %xmm2, %xmm3
   599: 660f71d302               psrlw $2, %xmm3
   604: 660fdbdd                 pand %xmm5, %xmm3
   608: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   614: 660f6fe2                 movdqa %xmm2, %xmm4
   618: 660f71d402               psrlw $2, %xmm4
   623: 660fdbe5                 pand %xmm5, %xmm4
   627: 660fefdc                 pxor %xmm4, %xmm3
   631: 660f6fe3                 movdqa %xmm3, %xmm4
   635: 660f71f401               psllw $1, %xmm4
   640: 660fefcc                 pxor %xmm4, %xmm1
   644: 660f6fe3                 movdqa %xmm3, %xmm4
   648: 660f71f402               psllw $2, %xmm4
   653: 660fefc4                 pxor %xmm4, %xmm0
   657: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   663: 660f6fda                 movdqa %xmm2, %xmm3
   667: 660f71d303               psrlw $3, %xmm3
   672: 660fdbdd                 pand %xmm5, %xmm3
   676: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   682: 660f6fe2                 movdqa %xmm2, %xmm4
   686: 660f71d403               psrlw $3, %xmm4
   691: 660fdbe5                 pand %xmm5, %xmm4
   695: 660fefdc                 pxor %xmm4, %xmm3
   699: 660f6fe3                 movdqa %xmm3, %xmm4
   703: 660f71f402               psllw $2, %xmm4
   708: 660fefcc                 pxor %xmm4, %xmm1
   712: 660f6fe3                 movdqa %xmm3, %xmm4
   716: 660f71f403               psllw $3, %xmm4
   721: 660fefc4                 pxor %xmm4, %xmm0
   725: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   731: 660f6fda                 movdqa %xmm2, %xmm3
   735: 660f71d304               psrlw $4, %xmm3
   740: 660fdbdd                 pand %xmm5, %xmm3
   744: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   750: 660f6fe2                 movdqa %xmm2, %xmm4
   754: 660f71d404               psrlw $4, %xmm4
   759: 660fdbe5                 pand %xmm5, %xmm4
   763: 660fefdc                 pxor %xmm4, %xmm3
   767: 660f6fe3                 movdqa %xmm3, %xmm4
   771: 660f71f403               psllw $3, %xmm4
   776: 660fefcc                 pxor %xmm4, %xmm1
   780: 660f6fe3                 movdqa %xmm3, %xmm4
   784: 660f71f404               psllw $4, %xmm4
   789: 660fefc4                 pxor %xmm4, %xmm0
   793: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   799: 660f6fda                 movdqa %xmm2, %xmm3
   803: 660f71d305               psrlw $5, %xmm3
   808: 660fdbdd                 pand %xmm5, %xmm3
   812: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   818: 660f6fe2                 movdqa %xmm2, %xmm4
   822: 660f71d405               psrlw $5, %xmm4
   827: 660fdbe5                 pand %xmm5, %xmm4
   831: 660fefdc                 pxor %xmm4, %xmm3
   835: 660f6fe3                 movdqa %xmm3, %xmm4
   839: 660f71f404               psllw $4, %xmm4
   844: 660fefcc                 pxor %xmm4, %xmm1
   848: 660f6fe3                 movdqa %xmm3, %xmm4
   852: 660f71f405               psllw $5, %xmm4
   857: 660fefc4                 pxor %xmm4, %xmm0
   861: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   867: 660f6fda                 movdqa %xmm2, %xmm3
   871: 660f71d306               psrlw $6, %xmm3
   876: 660fdbdd                 pand %xmm5, %xmm3
   880: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   886: 660f6fe2                 movdqa %xmm2, %xmm4
   890: 660f71d406               psrlw $6, %xmm4
   895: 660fdbe5                 pand %xmm5, %xmm4
   899: 660fefdc                 pxor %xmm4, %xmm3
   903: 660f6fe3                 movdqa %xmm3, %xmm4
   907: 660f71f405               psllw $5, %xmm4
   912: 660fefcc                 pxor %xmm4, %xmm1
   916: 660f6fe3                 movdqa %xmm3, %xmm4
   920: 660f71f406               psllw $6, %xmm4
   925: 660fefc4                 pxor %xmm4, %xmm0
   929: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   935: 660f6fda                 movdqa %xmm2, %xmm3
   939: 660f71d307               psrlw $7, %xmm3
   944: 660fdbdd                 pand %xmm5, %xmm3
   948: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   954: 660f6fe2                 movdqa %xmm2, %xmm4
   958: 660f71d407               psrlw $7, %xmm4
   963: 660fdbe5                 pand %xmm5, %xmm4
   967: 660fefdc                 pxor %xmm4, %xmm3
   971: 660f6fe3                 movdqa %xmm3, %xmm4
   975: 660f71f406               psllw $6, %xmm4
   980: 660fefcc                 pxor %xmm4, %xmm1
   984: 660f6fe3                 movdqa %xmm3, %xmm4
   988: 660f71f407               psllw $7, %xmm4
   993: 660fefc4                 pxor %xmm4, %xmm0
   997: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  1003: 660f6fda                 movdqa %xmm2, %xmm3
  1007: 660f71d300               psrlw $0, %xmm3
  1012: 660fdbdd                 pand %xmm5, %xmm3
  1016: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  1023: 660f6fe2                 movdqa %xmm2, %xmm4
  1027: 660f71d400               psrlw $0, %xmm4
  1032: 660fdbe5                 pand %xmm5, %xmm4
  1036: 660fefdc                 pxor %xmm4, %xmm3
  1040: 660f6fe3                 movdqa %xmm3, %xmm4
  1044: 660f71f401               psllw $1, %xmm4
  1049: 660fefcc                 pxor %xmm4, %xmm1
  1053: 660f6fe3                 movdqa %xmm3, %xmm4
  1057: 660f71f406               psllw $6, %xmm4
  1062: 660fefcc                 pxor %xmm4, %xmm1
  1066: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  1072: 660f6fda                 movdqa %xmm2, %xmm3
  1076: 660f71d302               psrlw $2, %xmm3
  1081: 660fdbdd                 pand %xmm5, %xmm3
  1085: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  1092: 660f6fe2                 movdqa %xmm2, %xmm4
  1096: 660f71d402               psrlw $2, %xmm4
  1101: 660fdbe5                 pand %xmm5, %xmm4
  1105: 660fefdc                 pxor %xmm4, %xmm3
  1109: 660fefcb                 pxor %xmm3, %xmm1
  1113: 660f6fe3                 movdqa %xmm3, %xmm4
  1117: 660f71f402               psllw $2, %xmm4
  1122: 660fefc4                 pxor %xmm4, %xmm0
  1126: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  1132: 660f6fda                 movdqa %xmm2, %xmm3
  1136: 660f71d303               psrlw $3, %xmm3
  1141: 660fdbdd                 pand %xmm5, %xmm3
  1145: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  1152: 660f6fe2                 movdqa %xmm2, %xmm4
  1156: 660f71d403               psrlw $3, %xmm4
  1161: 660fdbe5                 pand %xmm5, %xmm4
  1165: 660fefdc                 pxor %xmm4, %xmm3
  1169: 660f6fe3                 movdqa %xmm3, %xmm4
  1173: 660f71f401               psllw $1, %xmm4
  1178: 660fefcc                 pxor %xmm4, %xmm1
  1182: 660f6fe3                 movdqa %xmm3, %xmm4
  1186: 660f71f403               psllw $3, %xmm4
  1191: 660fefc4                 pxor %xmm4, %xmm0
  1195: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  1201: 660f6fda                 movdqa %xmm2, %xmm3
  1205: 660f71d304               psrlw $4, %xmm3
  1210: 660fdbdd                 pand %xmm5, %xmm3
  1214: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  1221: 660f6fe2                 movdqa %xmm2, %xmm4
  1225: 660f71d404               psrlw $4, %xmm4
  1230: 660fdbe5                 pand %xmm5, %xmm4
  1234: 660fefdc                 pxor %xmm4, %xmm3
  1238: 660f6fe3                 movdqa %xmm3, %xmm4
  1242: 660f71f402               psllw $2, %xmm4
  1247: 660fefcc                 pxor %xmm4, %xmm1
  1251: 660f6fe3                 movdqa %xmm3, %xmm4
  1255: 660f71f404               psllw $4, %xmm4
  1260: 660fefc4                 pxor %xmm4, %xmm0
  1264: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  1270: 660f6fda                 movdqa %xmm2, %xmm3
  1274: 660f71d305               psrlw $5, %xmm3
  1279: 660fdbdd                 pand %xmm5, %xmm3
  1283: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  1290: 660f6fe2                 movdqa %xmm2, %xmm4
  1294: 660f71d405               psrlw $5, %xmm4
  1299: 660fdbe5                 pand %xmm5, %xmm4
  1303: 660fefdc                 pxor %xmm4, %xmm3
  1307: 660f6fe3                 movdqa %xmm3, %xmm4
  1311: 660f71f403               psllw $3, %xmm4
  1316: 660fefcc                 pxor %xmm4, %xmm1
  1320: 660f6fe3                 movdqa %xmm3, %xmm4
  1324: 660f71f405               psllw $5, %xmm4
  1329: 660fefc4                 pxor %xmm4, %xmm0
  1333: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  1339: 660f6fda                 movdqa %xmm2, %xmm3
  1343: 660f71d306               psrlw $6, %xmm3
  1348: 660fdbdd                 pand %xmm5, %xmm3
  1352: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  1359: 660f6fe2                 movdqa %xmm2, %xmm4
  1363: 660f71d406               psrlw $6, %xmm4
  1368: 660fdbe5                 pand %xmm5, %xmm4
  1372: 660fefdc                 pxor %xmm4, %xmm3
  1376: 660f6fe3                 movdqa %xmm3, %xmm4
  1380: 660f71f404               psllw $4, %xmm4
  1385: 660fefcc                 pxor %xmm4, %xmm1
  1389: 660f6fe3                 movdqa %xmm3, %xmm4
  1393: 660f71f406               psllw $6, %xmm4
  1398: 660fefc4                 pxor %xmm4, %xmm0
  1402: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  1408: 660f6fda                 movdqa %xmm2, %xmm3
  1412: 660f71d300               psrlw $0, %xmm3
  1417: 660fdbdd                 pand %xmm5, %xmm3
  1421: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  1428: 660f6fe2                 movdqa %xmm2, %xmm4
  1432: 660f71d400               psrlw $0, %xmm4
  1437: 660fdbe5                 pand %xmm5, %xmm4
  1441: 660fefdc                 pxor %xmm4, %xmm3
  1445: 660fefc3                 pxor %xmm3, %xmm0
  1449: 660f6fe3                 movdqa %xmm3, %xmm4
  1453: 660f71f402               psllw $2, %xmm4
  1458: 660fefcc                 pxor %xmm4, %xmm1
  1462: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
  1468: 660f6fda                 movdqa %xmm2, %xmm3
  1472: 660f71d300               psrlw $0, %xmm3
  1477: 660fdbdd                 pand %xmm5, %xmm3
  1481: 660f6fe3                 movdqa %xmm3, %xmm4
  1485: 660f71f402               psllw $2, %xmm4
  1490: 660fefcc                 pxor %xmm4, %xmm1
  1494: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  1500: 660f6fda                 movdqa %xmm2, %xmm3
  1504: 660f71d301               psrlw $1, %xmm3
  1509: 660fdbdd                 pand %xmm5, %xmm3
  1513: 660fefcb                 pxor %xmm3, %xmm1
  1517: 660f6fda                 movdqa %xmm2, %xmm3
  1521: 660f71d307               psrlw $7, %xmm3
  1526: 660fdbdd                 pand %xmm5, %xmm3
  1530: 660f6fe3                 movdqa %xmm3, %xmm4
  1534: 660f71f405               psllw $5, %xmm4
  1539: 660fefcc                 pxor %xmm4, %xmm1
  1543: 660f6fe3                 movdqa %xmm3, %xmm4
  1547: 660f71f407               psllw $7, %xmm4
  1552: 660fefc4                 pxor %xmm4, %xmm0
  1556: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  1562: 660f6fda                 movdqa %xmm2, %xmm3
  1566: 660f71d300               psrlw $0, %xmm3
  1571: 660fdbdd                 pand %xmm5, %xmm3
  1575: 660f6fe3                 movdqa %xmm3, %xmm4
  1579: 660f71f403               psllw $3, %xmm4
  1584: 660fefcc                 pxor %xmm4, %xmm1
  1588: 660f6fe3                 movdqa %xmm3, %xmm4
  1592: 660f71f407               psllw $7, %xmm4
  1597: 660fefcc                 pxor %xmm4, %xmm1
  1601: 660f6fda                 movdqa %xmm2, %xmm3
  1605: 660f71d301               psrlw $1, %xmm3
  1610: 660fdbdd                 pand %xmm5, %xmm3
  1614: 660f6fe3                 movdqa %xmm3, %xmm4
  1618: 660f71f401               psllw $1, %xmm4
  1623: 660fefcc                 pxor %xmm4, %xmm1
  1627: 660f6fda                 movdqa %xmm2, %xmm3
  1631: 660f71d302               psrlw $2, %xmm3
  1636: 660fdbdd                 pand %xmm5, %xmm3
  1640: 660f6fe3                 movdqa %xmm3, %xmm4
  1644: 660f71f402               psllw $2, %xmm4
  1649: 660fefcc                 pxor %xmm4, %xmm1
  1653: 660f6fda                 movdqa %xmm2, %xmm3
  1657: 660f71d303               psrlw $3, %xmm3
  1662: 660fdbdd                 pand %xmm5, %xmm3
  1666: 660f6fe3                 movdqa %xmm3, %xmm4
  1670: 660f71f403               psllw $3, %xmm4
  1675: 660fefcc                 pxor %xmm4, %xmm1
  1679: 660f6fda                 movdqa %xmm2, %xmm3
  1683: 660f71d304               psrlw $4, %xmm3
  1688: 660fdbdd                 pand %xmm5, %xmm3
  1692: 660f6fe3                 movdqa %xmm3, %xmm4
  1696: 660f71f404               psllw $4, %xmm4
  1701: 660fefcc                 pxor %xmm4, %xmm1
  1705: 660f6fda                 movdqa %xmm2, %xmm3
  1709: 660f71d305               psrlw $5, %xmm3
  1714: 660fdbdd                 pand %xmm5, %xmm3
  1718: 660f6fe3                 movdqa %xmm3, %xmm4
  1722: 660f71f405               psllw $5, %xmm4
  1727: 660fefcc                 pxor %xmm4, %xmm1
  1731: 660f6fda                 movdqa %xmm2, %xmm3
  1735: 660f71d306               psrlw $6, %xmm3
  1740: 660fdbdd                 pand %xmm5, %xmm3
  1744: 660f6fe3                 movdqa %xmm3, %xmm4
  1748: 660f71f406               psllw $6, %xmm4
  1753: 660fefcc                 pxor %xmm4, %xmm1
  1757: 660f6fda                 movdqa %xmm2, %xmm3
  1761: 660f71d307               psrlw $7, %xmm3
  1766: 660fdbdd                 pand %xmm5, %xmm3
  1770: 660f6fe3                 movdqa %xmm3, %xmm4
  1774: 660f71f407               psllw $7, %xmm4
  1779: 660fefcc                 pxor %xmm4, %xmm1
  1783: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  1790: 660f6fda                 movdqa %xmm2, %xmm3
  1794: 660f71d302               psrlw $2, %xmm3
  1799: 660fdbdd                 pand %xmm5, %xmm3
  1803: 660f6fe3                 movdqa %xmm3, %xmm4
  1807: 660f71f401               psllw $1, %xmm4
  1812: 660fefcc                 pxor %xmm4, %xmm1
  1816: 660f6fda                 movdqa %xmm2, %xmm3
  1820: 660f71d303               psrlw $3, %xmm3
  1825: 660fdbdd                 pand %xmm5, %xmm3
  1829: 660f6fe3                 movdqa %xmm3, %xmm4
  1833: 660f71f402               psllw $2, %xmm4
  1838: 660fefcc                 pxor %xmm4, %xmm1
  1842: 660f6fda                 movdqa %xmm2, %xmm3
  1846: 660f71d304               psrlw $4, %xmm3
  1851: 660fdbdd                 pand %xmm5, %xmm3
  1855: 660f6fe3                 movdqa %xmm3, %xmm4
  1859: 660f71f403               psllw $3, %xmm4
  1864: 660fefcc                 pxor %xmm4, %xmm1
  1868: 660f6fda                 movdqa %xmm2, %xmm3
  1872: 660f71d305               psrlw $5, %xmm3
  1877: 660fdbdd                 pand %xmm5, %xmm3
  1881: 660f6fe3                 movdqa %xmm3, %xmm4
  1885: 660f71f404               psllw $4, %xmm4
  1890: 660fefcc                 pxor %xmm4, %xmm1
  1894: 660f6fda                 movdqa %xmm2, %xmm3
  1898: 660f71d306               psrlw $6, %xmm3
  1903: 660fdbdd                 pand %xmm5, %xmm3
  1907: 660f6fe3                 movdqa %xmm3, %xmm4
  1911: 660f71f405               psllw $5, %xmm4
  1916: 660fefcc                 pxor %xmm4, %xmm1
  1920: 66430f6f141e             movdqa (%r14, %r11), %xmm2
  1926: 660f6fda                 movdqa %xmm2, %xmm3
  1930: 660f71d300               psrlw $0, %xmm3
  1935: 660fdbdd                 pand %xmm5, %xmm3
  1939: 660fefc3                 pxor %xmm3, %xmm0
  1943: 660f6fe3                 movdqa %xmm3, %xmm4
  1947: 660f71f403               psllw $3, %xmm4
  1952: 660fefcc                 pxor %xmm4, %xmm1
  1956: 660f6fda                 movdqa %xmm2, %xmm3
  1960: 660f71d301               psrlw $1, %xmm3
  1965: 660fdbdd                 pand %xmm5, %xmm3
  1969: 660f6fe3                 movdqa %xmm3, %xmm4
  1973: 660f71f401               psllw $1, %xmm4
  1978: 660fefc4                 pxor %xmm4, %xmm0
  1982: 660f6fe3                 movdqa %xmm3, %xmm4
  1986: 660f71f404               psllw $4, %xmm4
  1991: 660fefcc                 pxor %xmm4, %xmm1
  1995: 660f6fda                 movdqa %xmm2, %xmm3
  1999: 660f71d302               psrlw $2, %xmm3
  2004: 660fdbdd                 pand %xmm5, %xmm3
  2008: 660f6fe3                 movdqa %xmm3, %xmm4
  2012: 660f71f402               psllw $2, %xmm4
  2017: 660fefc4                 pxor %xmm4, %xmm0
  2021: 660f6fe3                 movdqa %xmm3, %xmm4
  2025: 660f71f405               psllw $5, %xmm4
  2030: 660fefcc                 pxor %xmm4, %xmm1
  2034: 660f6fda                 movdqa %xmm2, %xmm3
  2038: 660f71d303               psrlw $3, %xmm3
  2043: 660fdbdd                 pand %xmm5, %xmm3
  2047: 660f6fe3                 movdqa %xmm3, %xmm4
  2051: 660f71f403               psllw $3, %xmm4
  2056: 660fefc4                 pxor %xmm4, %xmm0
  2060: 660f6fe3                 movdqa %xmm3, %xmm4
  2064: 660f71f406               psllw $6, %xmm4
  2069: 660fefcc                 pxor %xmm4, %xmm1
  2073: 660f6fda                 movdqa %xmm2, %xmm3
  2077: 660f71d304               psrlw $4, %xmm3
  2082: 660fdbdd                 pand %xmm5, %xmm3
  2086: 660f6fe3                 movdqa %xmm3, %xmm4
  2090: 660f71f404               psllw $4, %xmm4
  2095: 660fefc4                 pxor %xmm4, %xmm0
  2099: 660f6fe3                 movdqa %xmm3, %xmm4
  2103: 660f71f407               psllw $7, %xmm4
  2108: 660fefcc                 pxor %xmm4, %xmm1
  2112: 660f6fda                 movdqa %xmm2, %xmm3
  2116: 660f71d305               psrlw $5, %xmm3
  2121: 660fdbdd                 pand %xmm5, %xmm3
  2125: 660f6fe3                 movdqa %xmm3, %xmm4
  2129: 660f71f405               psllw $5, %xmm4
  2134: 660fefc4                 pxor %xmm4, %xmm0
  2138: 660f6fda                 movdqa %xmm2, %xmm3
  2142: 660f71d306               psrlw $6, %xmm3
  2147: 660fdbdd                 pand %xmm5, %xmm3
  2151: 660f6fe3                 movdqa %xmm3, %xmm4
  2155: 660f71f406               psllw $6, %xmm4
  2160: 660fefc4                 pxor %xmm4, %xmm0
  2164: 660f6fda                 movdqa %xmm2, %xmm3
  2168: 660f71d307               psrlw $7, %xmm3
  2173: 660fdbdd                 pand %xmm5, %xmm3
  2177: 660f6fe3                 movdqa %xmm3, %xmm4
  2181: 660f71f402               psllw $2, %xmm4
  2186: 660fefcc                 pxor %xmm4, %xmm1
  2190: 660f6fe3                 movdqa %xmm3, %xmm4
  2194: 660f71f403               psllw $3, %xmm4
  2199: 660fefcc                 pxor %xmm4, %xmm1
  2203: 4d8b5028                 movq 0x28(%r8), %r10
  2207: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  2213: 660f6fda                 movdqa %xmm2, %xmm3
  2217: 660f71d300               psrlw $0, %xmm3
  2222: 660fdbdd                 pand %xmm5, %xmm3
  2226: 660fefc3                 pxor %xmm3, %xmm0
  2230: 660f6fe3                 movdqa %xmm3, %xmm4
  2234: 660f71f402               psllw $2, %xmm4
  2239: 660fefcc                 pxor %xmm4, %xmm1
  2243: 660f6fda                 movdqa %xmm2, %xmm3
  2247: 660f71d301               psrlw $1, %xmm3
  2252: 660fdbdd                 pand %xmm5, %xmm3
  2256: 660f6fe3                 movdqa %xmm3, %xmm4
  2260: 660f71f401               psllw $1, %xmm4
  2265: 660fefc4                 pxor %xmm4, %xmm0
  2269: 660f6fe3                 movdqa %xmm3, %xmm4
  2273: 660f71f403               psllw $3, %xmm4
  2278: 660fefcc                 pxor %xmm4, %xmm1
  2282: 660f6fda                 movdqa %xmm2, %xmm3
  2286: 660f71d302               psrlw $2, %xmm3
  2291: 660fdbdd                 pand %xmm5, %xmm3
  2295: 660f6fe3                 movdqa %xmm3, %xmm4
  2299: 660f71f402               psllw $2, %xmm4
  2304: 660fefc4                 pxor %xmm4, %xmm0
  2308: 660f6fe3                 movdqa %xmm3, %xmm4
  2312: 660f71f404               psllw $4, %xmm4
  2317: 660fefcc                 pxor %xmm4, %xmm1
  2321: 660f6fda                 movdqa %xmm2, %xmm3
  2325: 660f71d303               psrlw $3, %xmm3
  2330: 660fdbdd                 pand %xmm5, %xmm3
  2334: 660f6fe3                 movdqa %xmm3, %xmm4
  2338: 660f71f403               psllw $3, %xmm4
  2343: 660fefc4                 pxor %xmm4, %xmm0
  2347: 660f6fe3                 movdqa %xmm3, %xmm4
  2351: 660f71f405               psllw $5, %xmm4
  2356: 660fefcc                 pxor %xmm4, %xmm1
  2360: 660f6fda                 movdqa %xmm2, %xmm3
  2364: 660f71d304               psrlw $4, %xmm3
  2369: 660fdbdd                 pand %xmm5, %xmm3
  2373: 660f6fe3                 movdqa %xmm3, %xmm4
  2377: 660f71f404               psllw $4, %xmm4
  2382: 660fefc4                 pxor %xmm4, %xmm0
  2386: 660f6fe3                 movdqa %xmm3, %xmm4
  2390: 660f71f406               psllw $6, %xmm4
  2395: 660fefcc                 pxor %xmm4, %xmm1
  2399: 660f6fda                 movdqa %xmm2, %xmm3
  2403: 660f71d305               psrlw $5, %xmm3
  2408: 660fdbdd                 pand %xmm5, %xmm3
  2412: 660f6fe3                 movdqa %xmm3, %xmm4
  2416: 660f71f405               psllw $5, %xmm4
  2421: 660fefc4                 pxor %xmm4, %xmm0
  2425: 660f6fe3                 movdqa %xmm3, %xmm4
  2429: 660f71f407               psllw $7, %xmm4
  2434: 660fefcc                 pxor %xmm4, %xmm1
  2438: 660f6fda                 movdqa %xmm2, %xmm3
  2442: 660f71d306               psrlw $6, %xmm3
  2447: 660fdbdd                 pand %xmm5, %xmm3
  2451: 660f6fe3                 movdqa %xmm3, %xmm4
  2455: 660f71f406               psllw $6, %xmm4
  2460: 660fefc4                 pxor %xmm4, %xmm0
  2464: 660f6fda                 movdqa %xmm2, %xmm3
  2468: 660f71d307               psrlw $7, %xmm3
  2473: 660fdbdd                 pand %xmm5, %xmm3
  2477: 660f6fe3                 movdqa %xmm3, %xmm4
  2481: 660f71f407               psllw $7, %xmm4
  2486: 660fefc4                 pxor %xmm4, %xmm0
  2490: 4d8b5030                 movq 0x30(%r8), %r10
  2494: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  2500: 660f6fda                 movdqa %xmm2, %xmm3
  2504: 660f71d300               psrlw $0, %xmm3
  2509: 660fdbdd                 pand %xmm5, %xmm3
  2513: 660fefc3                 pxor %xmm3, %xmm0
  2517: 660f6fe3                 movdqa %xmm3, %xmm4
  2521: 660f71f401               psllw $1, %xmm4
  2526: 660fefcc                 pxor %xmm4, %xmm1
  2530: 660f6fda                 movdqa %xmm2, %xmm3
  2534: 660f71d301               psrlw $1, %xmm3
  2539: 660fdbdd                 pand %xmm5, %xmm3
  2543: 660f6fe3                 movdqa %xmm3, %xmm4
  2547: 660f71f401               psllw $1, %xmm4
  2552: 660fefc4                 pxor %xmm4, %xmm0
  2556: 660f6fe3                 movdqa %xmm3, %xmm4
  2560: 660f71f402               psllw $2, %xmm4
  2565: 660fefcc                 pxor %xmm4, %xmm1
  2569: 660f6fda                 movdqa %xmm2, %xmm3
  2573: 660f71d302               psrlw $2, %xmm3
  2578: 660fdbdd                 pand %xmm5, %xmm3
  2582: 660f6fe3                 movdqa %xmm3, %xmm4
  2586: 660f71f402               psllw $2, %xmm4
  2591: 660fefc4                 pxor %xmm4, %xmm0
  2595: 660f6fe3                 movdqa %xmm3, %xmm4
  2599: 660f71f403               psllw $3, %xmm4
  2604: 660fefcc                 pxor %xmm4, %xmm1
  2608: 660f6fda                 movdqa %xmm2, %xmm3
  2612: 660f71d303               psrlw $3, %xmm3
  2617: 660fdbdd                 pand %xmm5, %xmm3
  2621: 660f6fe3                 movdqa %xmm3, %xmm4
  2625: 660f71f403               psllw $3, %xmm4
  2630: 660fefc4                 pxor %xmm4, %xmm0
  2634: 660f6fe3                 movdqa %xmm3, %xmm4
  2638: 660f71f404               psllw $4, %xmm4
  2643: 660fefcc                 pxor %xmm4, %xmm1
  2647: 660f6fda                 movdqa %xmm2, %xmm3
  2651: 660f71d304               psrlw $4, %xmm3
  2656: 660fdbdd                 pand %xmm5, %xmm3
  2660: 660f6fe3                 movdqa %xmm3, %xmm4
  2664: 660f71f404               psllw $4, %xmm4
  2669: 660fefc4                 pxor %xmm4, %xmm0
  2673: 660f6fe3                 movdqa %xmm3, %xmm4
  2677: 660f71f405               psllw $5, %xmm4
  2682: 660fefcc                 pxor %xmm4, %xmm1
  2686: 660f6fda                 movdqa %xmm2, %xmm3
  2690: 660f71d305               psrlw $5, %xmm3
  2695: 660fdbdd                 pand %xmm5, %xmm3
  2699: 660f6fe3                 movdqa %xmm3, %xmm4
  2703: 660f71f405               psllw $5, %xmm4
  2708: 660fefc4                 pxor %xmm4, %xmm0
  2712: 660f6fe3                 movdqa %xmm3, %xmm4
  2716: 660f71f406               psllw $6, %xmm4
  2721: 660fefcc                 pxor %xmm4, %xmm1
  2725: 660f6fda                 movdqa %xmm2, %xmm3
  2729: 660f71d306               psrlw $6, %xmm3
  2734: 660fdbdd                 pand %xmm5, %xmm3
  2738: 660f6fe3                 movdqa %xmm3, %xmm4
  2742: 660f71f406               psllw $6, %xmm4
  2747: 660fefc4                 pxor %xmm4, %xmm0
  2751: 660f6fe3                 movdqa %xmm3, %xmm4
  2755: 660f71f407               psllw $7, %xmm4
  2760: 660fefcc                 pxor %xmm4, %xmm1
  2764: 660f6fda                 movdqa %xmm2, %xmm3
  2768: 660f71d307               psrlw $7, %xmm3
  2773: 660fdbdd                 pand %xmm5, %xmm3
  2777: 660fefcb                 pxor %xmm3, %xmm1
  2781: 660f6fe3                 movdqa %xmm3, %xmm4
  2785: 660f71f401               psllw $1, %xmm4
  2790: 660fefcc                 pxor %xmm4, %xmm1
  2794: 660f6fe3                 movdqa %xmm3, %xmm4
  2798: 660f71f403               psllw $3, %xmm4
  2803: 660fefcc                 pxor %xmm4, %xmm1
  2807: 660f6fe3                 movdqa %xmm3, %xmm4
  2811: 660f71f404               psllw $4, %xmm4
  2816: 660fefcc                 pxor %xmm4, %xmm1
  2820: 660f6fe3                 movdqa %xmm3, %xmm4
  2824: 660f71f407               psllw $7, %xmm4
  2829: 660fefc4                 pxor %xmm4, %xmm0
  2833: 4d8b5038                 movq 0x38(%r8), %r10
  2837: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  2843: 660f6fda                 movdqa %xmm2, %xmm3
  2847: 660f71d300               psrlw $0, %xmm3
  2852: 660fdbdd                 pand %xmm5, %xmm3
  2856: 660fefc3                 pxor %xmm3, %xmm0
  2860: 660fefcb                 pxor %xmm3, %xmm1
  2864: 660f6fda                 movdqa %xmm2, %xmm3
  2868: 660f71d301               psrlw $1, %xmm3
  2873: 660fdbdd                 pand %xmm5, %xmm3
  2877: 660f6fe3                 movdqa %xmm3, %xmm4
  2881: 660f71f401               psllw $1, %xmm4
  2886: 660fefc4                 pxor %xmm4, %xmm0
  2890: 660fefcc                 pxor %xmm4, %xmm1
  2894: 660f6fda                 movdqa %xmm2, %xmm3
  2898: 660f71d302               psrlw $2, %xmm3
  2903: 660fdbdd                 pand %xmm5, %xmm3
  2907: 660f6fe3                 movdqa %xmm3, %xmm4
  2911: 660f71f402               psllw $2, %xmm4
  2916: 660fefc4                 pxor %xmm4, %xmm0
  2920: 660fefcc                 pxor %xmm4, %xmm1
  2924: 660f6fda                 movdqa %xmm2, %xmm3
  2928: 660f71d303               psrlw $3, %xmm3
  2933: 660fdbdd                 pand %xmm5, %xmm3
  2937: 660f6fe3                 movdqa %xmm3, %xmm4
  2941: 660f71f403               psllw $3, %xmm4
  2946: 660fefc4                 pxor %xmm4, %xmm0
  2950: 660fefcc                 pxor %xmm4, %xmm1
  2954: 660f6fda                 movdqa %xmm2, %xmm3
  2958: 660f71d304               psrlw $4, %xmm3
  2963: 660fdbdd                 pand %xmm5, %xmm3
  2967: 660f6fe3                 movdqa %xmm3, %xmm4
  2971: 660f71f404               psllw $4, %xmm4
  2976: 660fefc4                 pxor %xmm4, %xmm0
  2980: 660fefcc                 pxor %xmm4, %xmm1
  2984: 660f6fda                 movdqa %xmm2, %xmm3
  2988: 660f71d305               psrlw $5, %xmm3
  2993: 660fdbdd                 pand %xmm5, %xmm3
  2997: 660f6fe3                 movdqa %xmm3, %xmm4
  3001: 660f71f405               psllw $5, %xmm4
  3006: 660fefc4                 pxor %xmm4, %xmm0
  3010: 660fefcc                 pxor %xmm4, %xmm1
  3014: 660f6fda                 movdqa %xmm2, %xmm3
  3018: 660f71d306               psrlw $6, %xmm3
  3023: 660fdbdd                 pand %xmm5, %xmm3
  3027: 660f6fe3                 movdqa %xmm3, %xmm4
  3031: 660f71f406               psllw $6, %xmm4
  3036: 660fefc4                 pxor %xmm4, %xmm0
  3040: 660fefcc                 pxor %xmm4, %xmm1
  3044: 660f6fda                 movdqa %xmm2, %xmm3
  3048: 660f71d307               psrlw $7, %xmm3
  3053: 660fdbdd                 pand %xmm5, %xmm3
  3057: 660f6fe3                 movdqa %xmm3, %xmm4
  3061: 660f71f407               psllw $7, %xmm4
  3066: 660fefc4                 pxor %xmm4, %xmm0
  3070: 660fefcc                 pxor %xmm4, %xmm1
  3074: 4d8b11                   movq (%r9), %r10
  3077: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
  3083: 4d8b5108                 movq 8(%r9), %r10
  3087: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
  3093: 4983c310                 addq $0x10, %r11
  3097: 4c3bd8                   cmpq %rax, %r11
  3100: 0f863ef4ffff             jbe 0x60
  3106: 4c3bd9                   cmpq %rcx, %r11
  3109: 0f830c020000             jae 0xe37
  3115: 4153                     pushq %r11
  3117: 498d7f40                 leaq 0x40(%r15), %rdi
  3121: 498db740010000           leaq 0x140(%r15), %rsi
  3128: 4d8b11                   movq (%r9), %r10
  3131: 4c8b1c24                 movq (%rsp), %r11
  3135: 33db                     xorl %ebx, %ebx
  3137: 498b00                   movq (%r8), %rax
  3140: 420fb61418               movzbl (%rax, %r11), %edx
  3145: 32da                     xorb %dl, %bl
  3147: 498b4008                 movq 8(%r8), %rax
  3151: 420fb61418               movzbl (%rax, %r11), %edx
  3156: 32da                     xorb %dl, %bl
  3158: 498b4010                 movq 0x10(%r8), %rax
  3162: 420fb61418               movzbl (%rax, %r11), %edx
  3167: 32da                     xorb %dl, %bl
  3169: 498b4018                 movq 0x18(%r8), %rax
  3173: 420fb61418               movzbl (%rax, %r11), %edx
  3178: 32da                     xorb %dl, %bl
  3180: 498b4020                 movq 0x20(%r8), %rax
  3184: 420fb61418               movzbl (%rax, %r11), %edx
  3189: 32da                     xorb %dl, %bl
  3191: 498b4028                 movq 0x28(%r8), %rax
  3195: 420fb61418               movzbl (%rax, %r11), %edx
  3200: 32da                     xorb %dl, %bl
  3202: 498b4030                 movq 0x30(%r8), %rax
  3206: 420fb61418               movzbl (%rax, %r11), %edx
  3211: 32da                     xorb %dl, %bl
  3213: 498b4038                 movq 0x38(%r8), %rax
  3217: 420fb61418               movzbl (%rax, %r11), %edx
  3222: 32da                     xorb %dl, %bl
  3224: 43881c1a                 movb %bl, (%r10, %r11)
  3228: 33db                     xorl %ebx, %ebx
  3230: 49ffc3                   incq %r11
  3233: 4c3bd9                   cmpq %rcx, %r11
  3236: 0f8297ffffff             jb 0xc41
  3242: 4d8b5108                 movq 8(%r9), %r10
  3246: 4c8b1c24                 movq (%rsp), %r11
  3250: 33db                     xorl %ebx, %ebx
  3252: 498b00                   movq (%r8), %rax
  3255: 420fb61418               movzbl (%rax, %r11), %edx
  3260: 4885d2                   testq %rdx, %rdx
  3263: 0f8422000000             je 0xce7
  3269: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3273: 81c0e6000000             addl $0xe6, %eax
  3279: 81f8ff000000             cmpl $0xff, %eax
  3285: 0f8206000000             jb 0xce1
  3291: 81c001ffffff             addl $0xffffff01, %eax
  3297: 0fb60406                 movzbl (%rsi, %rax), %eax
  3301: 32d8                     xorb %al, %bl
  3303: 498b4008                 movq 8(%r8), %rax
  3307: 420fb61418               movzbl (%rax, %r11), %edx
  3312: 4885d2                   testq %rdx, %rdx
  3315: 0f8422000000             je 0xd1b
  3321: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3325: 81c0cd000000             addl $0xcd, %eax
  3331: 81f8ff000000             cmpl $0xff, %eax
  3337: 0f8206000000             jb 0xd15
  3343: 81c001ffffff             addl $0xffffff01, %eax
  3349: 0fb60406                 movzbl (%rsi, %rax), %eax
  3353: 32d8                     xorb %al, %bl
  3355: 498b4010                 movq 0x10(%r8), %rax
  3359: 420fb61418               movzbl (%rax, %r11), %edx
  3364: 4885d2                   testq %rdx, %rdx
  3367: 0f8422000000             je 0xd4f
  3373: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3377: 81c0e7000000             addl $0xe7, %eax
  3383: 81f8ff000000             cmpl $0xff, %eax
  3389: 0f8206000000             jb 0xd49
  3395: 81c001ffffff             addl $0xffffff01, %eax
  3401: 0fb60406                 movzbl (%rsi, %rax), %eax
  3405: 32d8                     xorb %al, %bl
  3407: 498b4018                 movq 0x18(%r8), %rax
  3411: 420fb61418               movzbl (%rax, %r11), %edx
  3416: 4885d2                   testq %rdx, %rdx
  3419: 0f8422000000             je 0xd83
  3425: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3429: 81c0ce000000             addl $0xce, %eax
  3435: 81f8ff000000             cmpl $0xff, %eax
  3441: 0f8206000000             jb 0xd7d
  3447: 81c001ffffff             addl $0xffffff01, %eax
  3453: 0fb60406                 movzbl (%rsi, %rax), %eax
  3457: 32d8                     xorb %al, %bl
  3459: 498b4020                 movq 0x20(%r8), %rax
  3463: 420fb61418               movzbl (%rax, %r11), %edx
  3468: 4885d2                   testq %rdx, %rdx
  3471: 0f841f000000             je 0xdb4
  3477: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3481: 83c04b                   addl $0x4b, %eax
  3484: 81f8ff000000             cmpl $0xff, %eax
  3490: 0f8206000000             jb 0xdae
  3496: 81c001ffffff             addl $0xffffff01, %eax
  3502: 0fb60406                 movzbl (%rsi, %rax), %eax
  3506: 32d8                     xorb %al, %bl
  3508: 498b4028                 movq 0x28(%r8), %rax
  3512: 420fb61418               movzbl (%rax, %r11), %edx
  3517: 4885d2                   testq %rdx, %rdx
  3520: 0f841f000000             je 0xde5
  3526: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3530: 83c032                   addl $0x32, %eax
  3533: 81f8ff000000             cmpl $0xff, %eax
  3539: 0f8206000000             jb 0xddf
  3545: 81c001ffffff             addl $0xffffff01, %eax
  3551: 0fb60406                 movzbl (%rsi, %rax), %eax
  3555: 32d8                     xorb %al, %bl
  3557: 498b4030                 movq 0x30(%r8), %rax
  3561: 420fb61418               movzbl (%rax, %r11), %edx
  3566: 4885d2                   testq %rdx, %rdx
  3569: 0f841f000000             je 0xe16
  3575: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3579: 83c019                   addl $0x19, %eax
  3582: 81f8ff000000             cmpl $0xff, %eax
  3588: 0f8206000000             jb 0xe10
  3594: 81c001ffffff             addl $0xffffff01, %eax
  3600: 0fb60406                 movzbl (%rsi, %rax), %eax
  3604: 32d8                     xorb %al, %bl
  3606: 498b4038                 movq 0x38(%r8), %rax
  3610: 420fb61418               movzbl (%rax, %r11), %edx
  3615: 32da                     xorb %dl, %bl
  3617: 43881c1a                 movb %bl, (%r10, %r11)
  3621: 33db                     xorl %ebx, %ebx
  3623: 49ffc3                   incq %r11
  3626: 4c3bd9                   cmpq %rcx, %r11
  3629: 0f8281feffff             jb 0xcb4
  3635: 4883c408                 addq $8, %rsp
  3639: 415e                     popq %r14
  3641: 415d                     popq %r13
  3643: 415c                     popq %r12
  3645: 415f                     popq %r15
  3647: 5b                       popq %rbx
  3648: c3                       retq 
  3649: 0f0b                     ud2 
  3651: 0f0b                     ud2 
  3653: 0f0b                     ud2 
  3655: 0f0b                     ud2 
  3657: 0f0b                     ud2 
  3659: 0f0b                     ud2 
  3661: 0f0b                     ud2 
  3663: 0f0b                     ud2 
  3665: 0f0b                     ud2 
  3667: 0f0b                     ud2 
  3669: 0f0b                     ud2 
  3671: 0f0b                     ud2 
  3673: 0f0b                     ud2 
  3675: 0f0b                     ud2 
  3677: 0f0b                     ud2 
  3679: 0f0b                     ud2 
  3681: 0f0b                     ud2 
  3683: 0f0b                     ud2 
  3685: 0f0b                     ud2 
  3687: 0f0b                     ud2 
  3689: 0f0b                     ud2 
  3691: 0f0b                     ud2 
  3693: 0f0b                     ud2 
  3695: 0f0b                     ud2 
  3697: 0f0b                     ud2 
  3699: 0f0b                     ud2 
  3701: 0f0b                     ud2 
  3703: 0f0b                     ud2 
  3705: 0f0b                     ud2 
  3707: 0f0b                     ud2 
  3709: 0f0b                     ud2 
  3711: cc                       int3 
