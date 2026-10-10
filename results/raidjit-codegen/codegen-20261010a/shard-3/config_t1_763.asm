; config_t1_763 target=sse twin=xorbit topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=4224 twins=["xorbit"] uops_per_byte=47.8125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d67100000           leaq 0x1067(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66410f6f2f               movdqa (%r15), %xmm5
    52: 488bc1                   movq %rcx, %rax
    55: 4883e810                 subq $0x10, %rax
    59: 0f82620d0000             jb 0xda3
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   102: 660f6fda                 movdqa %xmm2, %xmm3
   106: 660f71d304               psrlw $4, %xmm3
   111: 660fdbdd                 pand %xmm5, %xmm3
   115: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
   121: 660f6fe2                 movdqa %xmm2, %xmm4
   125: 660f71d405               psrlw $5, %xmm4
   130: 660fdbe5                 pand %xmm5, %xmm4
   134: 660fefdc                 pxor %xmm4, %xmm3
   138: 660f6fc3                 movdqa %xmm3, %xmm0
   142: 660f6fcb                 movdqa %xmm3, %xmm1
   146: 660f6fe3                 movdqa %xmm3, %xmm4
   150: 660f71f401               psllw $1, %xmm4
   155: 660fefc4                 pxor %xmm4, %xmm0
   159: 660fefcc                 pxor %xmm4, %xmm1
   163: 660f6fe3                 movdqa %xmm3, %xmm4
   167: 660f71f403               psllw $3, %xmm4
   172: 660fefc4                 pxor %xmm4, %xmm0
   176: 660fefcc                 pxor %xmm4, %xmm1
   180: 660f6fe3                 movdqa %xmm3, %xmm4
   184: 660f71f404               psllw $4, %xmm4
   189: 660fefcc                 pxor %xmm4, %xmm1
   193: 660f6fe3                 movdqa %xmm3, %xmm4
   197: 660f71f405               psllw $5, %xmm4
   202: 660fefc4                 pxor %xmm4, %xmm0
   206: 660f6fe3                 movdqa %xmm3, %xmm4
   210: 660f71f406               psllw $6, %xmm4
   215: 660fefc4                 pxor %xmm4, %xmm0
   219: 660fefcc                 pxor %xmm4, %xmm1
   223: 660f6fe3                 movdqa %xmm3, %xmm4
   227: 660f71f407               psllw $7, %xmm4
   232: 660fefc4                 pxor %xmm4, %xmm0
   236: 660fefcc                 pxor %xmm4, %xmm1
   240: 66430f6f141e             movdqa (%r14, %r11), %xmm2
   246: 660f6fda                 movdqa %xmm2, %xmm3
   250: 660f71d300               psrlw $0, %xmm3
   255: 660fdbdd                 pand %xmm5, %xmm3
   259: 660f6fe2                 movdqa %xmm2, %xmm4
   263: 660f71d401               psrlw $1, %xmm4
   268: 660fdbe5                 pand %xmm5, %xmm4
   272: 660fefdc                 pxor %xmm4, %xmm3
   276: 660fefcb                 pxor %xmm3, %xmm1
   280: 660f6fe3                 movdqa %xmm3, %xmm4
   284: 660f71f401               psllw $1, %xmm4
   289: 660fefc4                 pxor %xmm4, %xmm0
   293: 660f6fe3                 movdqa %xmm3, %xmm4
   297: 660f71f402               psllw $2, %xmm4
   302: 660fefc4                 pxor %xmm4, %xmm0
   306: 660fefcc                 pxor %xmm4, %xmm1
   310: 660f6fe3                 movdqa %xmm3, %xmm4
   314: 660f71f404               psllw $4, %xmm4
   319: 660fefc4                 pxor %xmm4, %xmm0
   323: 660fefcc                 pxor %xmm4, %xmm1
   327: 660f6fe3                 movdqa %xmm3, %xmm4
   331: 660f71f405               psllw $5, %xmm4
   336: 660fefc4                 pxor %xmm4, %xmm0
   340: 660fefcc                 pxor %xmm4, %xmm1
   344: 660f6fe3                 movdqa %xmm3, %xmm4
   348: 660f71f406               psllw $6, %xmm4
   353: 660fefc4                 pxor %xmm4, %xmm0
   357: 660fefcc                 pxor %xmm4, %xmm1
   361: 660f6fe3                 movdqa %xmm3, %xmm4
   365: 660f71f407               psllw $7, %xmm4
   370: 660fefc4                 pxor %xmm4, %xmm0
   374: 660fefcc                 pxor %xmm4, %xmm1
   378: 660f6fda                 movdqa %xmm2, %xmm3
   382: 660f71d302               psrlw $2, %xmm3
   387: 660fdbdd                 pand %xmm5, %xmm3
   391: 4d8b5028                 movq 0x28(%r8), %r10
   395: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   401: 660f6fe2                 movdqa %xmm2, %xmm4
   405: 660f71d400               psrlw $0, %xmm4
   410: 660fdbe5                 pand %xmm5, %xmm4
   414: 660fefdc                 pxor %xmm4, %xmm3
   418: 660fefc3                 pxor %xmm3, %xmm0
   422: 660fefcb                 pxor %xmm3, %xmm1
   426: 660f6fe3                 movdqa %xmm3, %xmm4
   430: 660f71f402               psllw $2, %xmm4
   435: 660fefc4                 pxor %xmm4, %xmm0
   439: 660f6fe3                 movdqa %xmm3, %xmm4
   443: 660f71f404               psllw $4, %xmm4
   448: 660fefc4                 pxor %xmm4, %xmm0
   452: 660fefcc                 pxor %xmm4, %xmm1
   456: 660f6fe3                 movdqa %xmm3, %xmm4
   460: 660f71f405               psllw $5, %xmm4
   465: 660fefc4                 pxor %xmm4, %xmm0
   469: 660fefcc                 pxor %xmm4, %xmm1
   473: 660f6fe3                 movdqa %xmm3, %xmm4
   477: 660f71f406               psllw $6, %xmm4
   482: 660fefc4                 pxor %xmm4, %xmm0
   486: 660fefcc                 pxor %xmm4, %xmm1
   490: 660f6fe3                 movdqa %xmm3, %xmm4
   494: 660f71f407               psllw $7, %xmm4
   499: 660fefc4                 pxor %xmm4, %xmm0
   503: 660fefcc                 pxor %xmm4, %xmm1
   507: 66430f6f141e             movdqa (%r14, %r11), %xmm2
   513: 660f6fda                 movdqa %xmm2, %xmm3
   517: 660f71d303               psrlw $3, %xmm3
   522: 660fdbdd                 pand %xmm5, %xmm3
   526: 4d8b5028                 movq 0x28(%r8), %r10
   530: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   536: 660f6fe2                 movdqa %xmm2, %xmm4
   540: 660f71d402               psrlw $2, %xmm4
   545: 660fdbe5                 pand %xmm5, %xmm4
   549: 660fefdc                 pxor %xmm4, %xmm3
   553: 660fefc3                 pxor %xmm3, %xmm0
   557: 660fefcb                 pxor %xmm3, %xmm1
   561: 660f6fe3                 movdqa %xmm3, %xmm4
   565: 660f71f403               psllw $3, %xmm4
   570: 660fefcc                 pxor %xmm4, %xmm1
   574: 660f6fe3                 movdqa %xmm3, %xmm4
   578: 660f71f404               psllw $4, %xmm4
   583: 660fefc4                 pxor %xmm4, %xmm0
   587: 660fefcc                 pxor %xmm4, %xmm1
   591: 660f6fe3                 movdqa %xmm3, %xmm4
   595: 660f71f405               psllw $5, %xmm4
   600: 660fefc4                 pxor %xmm4, %xmm0
   604: 660fefcc                 pxor %xmm4, %xmm1
   608: 660f6fe3                 movdqa %xmm3, %xmm4
   612: 660f71f406               psllw $6, %xmm4
   617: 660fefc4                 pxor %xmm4, %xmm0
   621: 660fefcc                 pxor %xmm4, %xmm1
   625: 660f6fe3                 movdqa %xmm3, %xmm4
   629: 660f71f407               psllw $7, %xmm4
   634: 660fefc4                 pxor %xmm4, %xmm0
   638: 660fefcc                 pxor %xmm4, %xmm1
   642: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   648: 660f6fda                 movdqa %xmm2, %xmm3
   652: 660f71d305               psrlw $5, %xmm3
   657: 660fdbdd                 pand %xmm5, %xmm3
   661: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
   667: 660f6fe2                 movdqa %xmm2, %xmm4
   671: 660f71d406               psrlw $6, %xmm4
   676: 660fdbe5                 pand %xmm5, %xmm4
   680: 660fefdc                 pxor %xmm4, %xmm3
   684: 660fefc3                 pxor %xmm3, %xmm0
   688: 660fefcb                 pxor %xmm3, %xmm1
   692: 660f6fe3                 movdqa %xmm3, %xmm4
   696: 660f71f402               psllw $2, %xmm4
   701: 660fefc4                 pxor %xmm4, %xmm0
   705: 660fefcc                 pxor %xmm4, %xmm1
   709: 660f6fe3                 movdqa %xmm3, %xmm4
   713: 660f71f403               psllw $3, %xmm4
   718: 660fefc4                 pxor %xmm4, %xmm0
   722: 660fefcc                 pxor %xmm4, %xmm1
   726: 660f6fe3                 movdqa %xmm3, %xmm4
   730: 660f71f405               psllw $5, %xmm4
   735: 660fefcc                 pxor %xmm4, %xmm1
   739: 660f6fe3                 movdqa %xmm3, %xmm4
   743: 660f71f406               psllw $6, %xmm4
   748: 660fefc4                 pxor %xmm4, %xmm0
   752: 660f6fe3                 movdqa %xmm3, %xmm4
   756: 660f71f407               psllw $7, %xmm4
   761: 660fefc4                 pxor %xmm4, %xmm0
   765: 660fefcc                 pxor %xmm4, %xmm1
   769: 66430f6f141e             movdqa (%r14, %r11), %xmm2
   775: 660f6fda                 movdqa %xmm2, %xmm3
   779: 660f71d304               psrlw $4, %xmm3
   784: 660fdbdd                 pand %xmm5, %xmm3
   788: 4d8b5028                 movq 0x28(%r8), %r10
   792: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   798: 660f6fe2                 movdqa %xmm2, %xmm4
   802: 660f71d403               psrlw $3, %xmm4
   807: 660fdbe5                 pand %xmm5, %xmm4
   811: 660fefdc                 pxor %xmm4, %xmm3
   815: 660fefc3                 pxor %xmm3, %xmm0
   819: 660fefcb                 pxor %xmm3, %xmm1
   823: 660f6fe3                 movdqa %xmm3, %xmm4
   827: 660f71f403               psllw $3, %xmm4
   832: 660fefc4                 pxor %xmm4, %xmm0
   836: 660fefcc                 pxor %xmm4, %xmm1
   840: 660f6fe3                 movdqa %xmm3, %xmm4
   844: 660f71f405               psllw $5, %xmm4
   849: 660fefc4                 pxor %xmm4, %xmm0
   853: 660fefcc                 pxor %xmm4, %xmm1
   857: 660f6fe3                 movdqa %xmm3, %xmm4
   861: 660f71f406               psllw $6, %xmm4
   866: 660fefc4                 pxor %xmm4, %xmm0
   870: 660fefcc                 pxor %xmm4, %xmm1
   874: 660f6fe3                 movdqa %xmm3, %xmm4
   878: 660f71f407               psllw $7, %xmm4
   883: 660fefc4                 pxor %xmm4, %xmm0
   887: 660fefcc                 pxor %xmm4, %xmm1
   891: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   897: 660f6fda                 movdqa %xmm2, %xmm3
   901: 660f71d306               psrlw $6, %xmm3
   906: 660fdbdd                 pand %xmm5, %xmm3
   910: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
   917: 660f6fe2                 movdqa %xmm2, %xmm4
   921: 660f71d407               psrlw $7, %xmm4
   926: 660fdbe5                 pand %xmm5, %xmm4
   930: 660fefdc                 pxor %xmm4, %xmm3
   934: 660fefc3                 pxor %xmm3, %xmm0
   938: 660fefcb                 pxor %xmm3, %xmm1
   942: 660f6fe3                 movdqa %xmm3, %xmm4
   946: 660f71f401               psllw $1, %xmm4
   951: 660fefc4                 pxor %xmm4, %xmm0
   955: 660fefcc                 pxor %xmm4, %xmm1
   959: 660f6fe3                 movdqa %xmm3, %xmm4
   963: 660f71f403               psllw $3, %xmm4
   968: 660fefc4                 pxor %xmm4, %xmm0
   972: 660fefcc                 pxor %xmm4, %xmm1
   976: 660f6fe3                 movdqa %xmm3, %xmm4
   980: 660f71f404               psllw $4, %xmm4
   985: 660fefc4                 pxor %xmm4, %xmm0
   989: 660fefcc                 pxor %xmm4, %xmm1
   993: 660f6fe3                 movdqa %xmm3, %xmm4
   997: 660f71f407               psllw $7, %xmm4
  1002: 660fefc4                 pxor %xmm4, %xmm0
  1006: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
  1012: 660f6fda                 movdqa %xmm2, %xmm3
  1016: 660f71d303               psrlw $3, %xmm3
  1021: 660fdbdd                 pand %xmm5, %xmm3
  1025: 4d8b5028                 movq 0x28(%r8), %r10
  1029: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1035: 660f6fe2                 movdqa %xmm2, %xmm4
  1039: 660f71d401               psrlw $1, %xmm4
  1044: 660fdbe5                 pand %xmm5, %xmm4
  1048: 660fefdc                 pxor %xmm4, %xmm3
  1052: 660f6fe3                 movdqa %xmm3, %xmm4
  1056: 660f71f404               psllw $4, %xmm4
  1061: 660fefc4                 pxor %xmm4, %xmm0
  1065: 660fefcc                 pxor %xmm4, %xmm1
  1069: 660f6fe3                 movdqa %xmm3, %xmm4
  1073: 660f71f405               psllw $5, %xmm4
  1078: 660fefc4                 pxor %xmm4, %xmm0
  1082: 660fefcc                 pxor %xmm4, %xmm1
  1086: 660f6fe3                 movdqa %xmm3, %xmm4
  1090: 660f71f406               psllw $6, %xmm4
  1095: 660fefc4                 pxor %xmm4, %xmm0
  1099: 660fefcc                 pxor %xmm4, %xmm1
  1103: 660f6fe3                 movdqa %xmm3, %xmm4
  1107: 660f71f407               psllw $7, %xmm4
  1112: 660fefc4                 pxor %xmm4, %xmm0
  1116: 660fefcc                 pxor %xmm4, %xmm1
  1120: 66430f6f141e             movdqa (%r14, %r11), %xmm2
  1126: 660f6fda                 movdqa %xmm2, %xmm3
  1130: 660f71d305               psrlw $5, %xmm3
  1135: 660fdbdd                 pand %xmm5, %xmm3
  1139: 4d8b5028                 movq 0x28(%r8), %r10
  1143: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1149: 660f6fe2                 movdqa %xmm2, %xmm4
  1153: 660f71d404               psrlw $4, %xmm4
  1158: 660fdbe5                 pand %xmm5, %xmm4
  1162: 660fefdc                 pxor %xmm4, %xmm3
  1166: 660fefc3                 pxor %xmm3, %xmm0
  1170: 660fefcb                 pxor %xmm3, %xmm1
  1174: 660f6fe3                 movdqa %xmm3, %xmm4
  1178: 660f71f403               psllw $3, %xmm4
  1183: 660fefc4                 pxor %xmm4, %xmm0
  1187: 660fefcc                 pxor %xmm4, %xmm1
  1191: 660f6fe3                 movdqa %xmm3, %xmm4
  1195: 660f71f406               psllw $6, %xmm4
  1200: 660fefc4                 pxor %xmm4, %xmm0
  1204: 660fefcc                 pxor %xmm4, %xmm1
  1208: 660f6fe3                 movdqa %xmm3, %xmm4
  1212: 660f71f407               psllw $7, %xmm4
  1217: 660fefc4                 pxor %xmm4, %xmm0
  1221: 660fefcc                 pxor %xmm4, %xmm1
  1225: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
  1231: 660f6fda                 movdqa %xmm2, %xmm3
  1235: 660f71d300               psrlw $0, %xmm3
  1240: 660fdbdd                 pand %xmm5, %xmm3
  1244: 660f6fe2                 movdqa %xmm2, %xmm4
  1248: 660f71d401               psrlw $1, %xmm4
  1253: 660fdbe5                 pand %xmm5, %xmm4
  1257: 660fefdc                 pxor %xmm4, %xmm3
  1261: 660f6fe3                 movdqa %xmm3, %xmm4
  1265: 660f71f401               psllw $1, %xmm4
  1270: 660fefc4                 pxor %xmm4, %xmm0
  1274: 660f6fe3                 movdqa %xmm3, %xmm4
  1278: 660f71f402               psllw $2, %xmm4
  1283: 660fefc4                 pxor %xmm4, %xmm0
  1287: 660fefcc                 pxor %xmm4, %xmm1
  1291: 660f6fe3                 movdqa %xmm3, %xmm4
  1295: 660f71f403               psllw $3, %xmm4
  1300: 660fefc4                 pxor %xmm4, %xmm0
  1304: 660fefcc                 pxor %xmm4, %xmm1
  1308: 660f6fe3                 movdqa %xmm3, %xmm4
  1312: 660f71f404               psllw $4, %xmm4
  1317: 660fefc4                 pxor %xmm4, %xmm0
  1321: 660fefcc                 pxor %xmm4, %xmm1
  1325: 660f6fda                 movdqa %xmm2, %xmm3
  1329: 660f71d302               psrlw $2, %xmm3
  1334: 660fdbdd                 pand %xmm5, %xmm3
  1338: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  1344: 660f6fe2                 movdqa %xmm2, %xmm4
  1348: 660f71d402               psrlw $2, %xmm4
  1353: 660fdbe5                 pand %xmm5, %xmm4
  1357: 660fefdc                 pxor %xmm4, %xmm3
  1361: 660f6fe3                 movdqa %xmm3, %xmm4
  1365: 660f71f402               psllw $2, %xmm4
  1370: 660fefc4                 pxor %xmm4, %xmm0
  1374: 660f6fe3                 movdqa %xmm3, %xmm4
  1378: 660f71f403               psllw $3, %xmm4
  1383: 660fefc4                 pxor %xmm4, %xmm0
  1387: 660fefcc                 pxor %xmm4, %xmm1
  1391: 660f6fe3                 movdqa %xmm3, %xmm4
  1395: 660f71f404               psllw $4, %xmm4
  1400: 660fefc4                 pxor %xmm4, %xmm0
  1404: 660fefcc                 pxor %xmm4, %xmm1
  1408: 660f6fe3                 movdqa %xmm3, %xmm4
  1412: 660f71f405               psllw $5, %xmm4
  1417: 660fefc4                 pxor %xmm4, %xmm0
  1421: 660fefcc                 pxor %xmm4, %xmm1
  1425: 66430f6f141e             movdqa (%r14, %r11), %xmm2
  1431: 660f6fda                 movdqa %xmm2, %xmm3
  1435: 660f71d306               psrlw $6, %xmm3
  1440: 660fdbdd                 pand %xmm5, %xmm3
  1444: 4d8b5028                 movq 0x28(%r8), %r10
  1448: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  1454: 660f6fe2                 movdqa %xmm2, %xmm4
  1458: 660f71d405               psrlw $5, %xmm4
  1463: 660fdbe5                 pand %xmm5, %xmm4
  1467: 660fefdc                 pxor %xmm4, %xmm3
  1471: 660fefc3                 pxor %xmm3, %xmm0
  1475: 660fefcb                 pxor %xmm3, %xmm1
  1479: 660f6fe3                 movdqa %xmm3, %xmm4
  1483: 660f71f403               psllw $3, %xmm4
  1488: 660fefc4                 pxor %xmm4, %xmm0
  1492: 660fefcc                 pxor %xmm4, %xmm1
  1496: 660f6fe3                 movdqa %xmm3, %xmm4
  1500: 660f71f407               psllw $7, %xmm4
  1505: 660fefc4                 pxor %xmm4, %xmm0
  1509: 660fefcc                 pxor %xmm4, %xmm1
  1513: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  1519: 660f6fda                 movdqa %xmm2, %xmm3
  1523: 660f71d300               psrlw $0, %xmm3
  1528: 660fdbdd                 pand %xmm5, %xmm3
  1532: 660f6fe2                 movdqa %xmm2, %xmm4
  1536: 660f71d401               psrlw $1, %xmm4
  1541: 660fdbe5                 pand %xmm5, %xmm4
  1545: 660fefdc                 pxor %xmm4, %xmm3
  1549: 660f6fe3                 movdqa %xmm3, %xmm4
  1553: 660f71f401               psllw $1, %xmm4
  1558: 660fefc4                 pxor %xmm4, %xmm0
  1562: 660f6fe3                 movdqa %xmm3, %xmm4
  1566: 660f71f402               psllw $2, %xmm4
  1571: 660fefc4                 pxor %xmm4, %xmm0
  1575: 660fefcc                 pxor %xmm4, %xmm1
  1579: 660f6fe3                 movdqa %xmm3, %xmm4
  1583: 660f71f403               psllw $3, %xmm4
  1588: 660fefc4                 pxor %xmm4, %xmm0
  1592: 660fefcc                 pxor %xmm4, %xmm1
  1596: 660f6fda                 movdqa %xmm2, %xmm3
  1600: 660f71d303               psrlw $3, %xmm3
  1605: 660fdbdd                 pand %xmm5, %xmm3
  1609: 660f6fe2                 movdqa %xmm2, %xmm4
  1613: 660f71d404               psrlw $4, %xmm4
  1618: 660fdbe5                 pand %xmm5, %xmm4
  1622: 660fefdc                 pxor %xmm4, %xmm3
  1626: 660f6fe3                 movdqa %xmm3, %xmm4
  1630: 660f71f404               psllw $4, %xmm4
  1635: 660fefc4                 pxor %xmm4, %xmm0
  1639: 660f6fe3                 movdqa %xmm3, %xmm4
  1643: 660f71f405               psllw $5, %xmm4
  1648: 660fefc4                 pxor %xmm4, %xmm0
  1652: 660fefcc                 pxor %xmm4, %xmm1
  1656: 660f6fe3                 movdqa %xmm3, %xmm4
  1660: 660f71f406               psllw $6, %xmm4
  1665: 660fefc4                 pxor %xmm4, %xmm0
  1669: 660fefcc                 pxor %xmm4, %xmm1
  1673: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  1679: 660f6fda                 movdqa %xmm2, %xmm3
  1683: 660f71d307               psrlw $7, %xmm3
  1688: 660fdbdd                 pand %xmm5, %xmm3
  1692: 66430f6f141e             movdqa (%r14, %r11), %xmm2
  1698: 660f6fe2                 movdqa %xmm2, %xmm4
  1702: 660f71d407               psrlw $7, %xmm4
  1707: 660fdbe5                 pand %xmm5, %xmm4
  1711: 660fefdc                 pxor %xmm4, %xmm3
  1715: 660fefc3                 pxor %xmm3, %xmm0
  1719: 660fefcb                 pxor %xmm3, %xmm1
  1723: 660f6fe3                 movdqa %xmm3, %xmm4
  1727: 660f71f403               psllw $3, %xmm4
  1732: 660fefc4                 pxor %xmm4, %xmm0
  1736: 660fefcc                 pxor %xmm4, %xmm1
  1740: 660f6fe3                 movdqa %xmm3, %xmm4
  1744: 660f71f407               psllw $7, %xmm4
  1749: 660fefc4                 pxor %xmm4, %xmm0
  1753: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
  1759: 660f6fda                 movdqa %xmm2, %xmm3
  1763: 660f71d306               psrlw $6, %xmm3
  1768: 660fdbdd                 pand %xmm5, %xmm3
  1772: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  1778: 660f6fe2                 movdqa %xmm2, %xmm4
  1782: 660f71d407               psrlw $7, %xmm4
  1787: 660fdbe5                 pand %xmm5, %xmm4
  1791: 660fefdc                 pxor %xmm4, %xmm3
  1795: 660fefc3                 pxor %xmm3, %xmm0
  1799: 660fefcb                 pxor %xmm3, %xmm1
  1803: 660f6fe3                 movdqa %xmm3, %xmm4
  1807: 660f71f406               psllw $6, %xmm4
  1812: 660fefcc                 pxor %xmm4, %xmm1
  1816: 660f6fe3                 movdqa %xmm3, %xmm4
  1820: 660f71f407               psllw $7, %xmm4
  1825: 660fefc4                 pxor %xmm4, %xmm0
  1829: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
  1835: 660f6fda                 movdqa %xmm2, %xmm3
  1839: 660f71d307               psrlw $7, %xmm3
  1844: 660fdbdd                 pand %xmm5, %xmm3
  1848: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  1854: 660f6fe2                 movdqa %xmm2, %xmm4
  1858: 660f71d402               psrlw $2, %xmm4
  1863: 660fdbe5                 pand %xmm5, %xmm4
  1867: 660fefdc                 pxor %xmm4, %xmm3
  1871: 660f6fe3                 movdqa %xmm3, %xmm4
  1875: 660f71f403               psllw $3, %xmm4
  1880: 660fefc4                 pxor %xmm4, %xmm0
  1884: 660fefcc                 pxor %xmm4, %xmm1
  1888: 660f6fe3                 movdqa %xmm3, %xmm4
  1892: 660f71f404               psllw $4, %xmm4
  1897: 660fefc4                 pxor %xmm4, %xmm0
  1901: 660fefcc                 pxor %xmm4, %xmm1
  1905: 660f6fda                 movdqa %xmm2, %xmm3
  1909: 660f71d300               psrlw $0, %xmm3
  1914: 660fdbdd                 pand %xmm5, %xmm3
  1918: 660f6fe2                 movdqa %xmm2, %xmm4
  1922: 660f71d401               psrlw $1, %xmm4
  1927: 660fdbe5                 pand %xmm5, %xmm4
  1931: 660fefdc                 pxor %xmm4, %xmm3
  1935: 660f6fe3                 movdqa %xmm3, %xmm4
  1939: 660f71f401               psllw $1, %xmm4
  1944: 660fefc4                 pxor %xmm4, %xmm0
  1948: 660f6fe3                 movdqa %xmm3, %xmm4
  1952: 660f71f402               psllw $2, %xmm4
  1957: 660fefc4                 pxor %xmm4, %xmm0
  1961: 660fefcc                 pxor %xmm4, %xmm1
  1965: 660f6fda                 movdqa %xmm2, %xmm3
  1969: 660f71d303               psrlw $3, %xmm3
  1974: 660fdbdd                 pand %xmm5, %xmm3
  1978: 660f6fe2                 movdqa %xmm2, %xmm4
  1982: 660f71d404               psrlw $4, %xmm4
  1987: 660fdbe5                 pand %xmm5, %xmm4
  1991: 660fefdc                 pxor %xmm4, %xmm3
  1995: 660f6fe3                 movdqa %xmm3, %xmm4
  1999: 660f71f404               psllw $4, %xmm4
  2004: 660fefc4                 pxor %xmm4, %xmm0
  2008: 660f6fe3                 movdqa %xmm3, %xmm4
  2012: 660f71f405               psllw $5, %xmm4
  2017: 660fefc4                 pxor %xmm4, %xmm0
  2021: 660fefcc                 pxor %xmm4, %xmm1
  2025: 660f6fda                 movdqa %xmm2, %xmm3
  2029: 660f71d305               psrlw $5, %xmm3
  2034: 660fdbdd                 pand %xmm5, %xmm3
  2038: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  2045: 660f6fe2                 movdqa %xmm2, %xmm4
  2049: 660f71d405               psrlw $5, %xmm4
  2054: 660fdbe5                 pand %xmm5, %xmm4
  2058: 660fefdc                 pxor %xmm4, %xmm3
  2062: 660f6fe3                 movdqa %xmm3, %xmm4
  2066: 660f71f405               psllw $5, %xmm4
  2071: 660fefc4                 pxor %xmm4, %xmm0
  2075: 660f6fe3                 movdqa %xmm3, %xmm4
  2079: 660f71f406               psllw $6, %xmm4
  2084: 660fefc4                 pxor %xmm4, %xmm0
  2088: 660fefcc                 pxor %xmm4, %xmm1
  2092: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
  2098: 660f6fda                 movdqa %xmm2, %xmm3
  2102: 660f71d300               psrlw $0, %xmm3
  2107: 660fdbdd                 pand %xmm5, %xmm3
  2111: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  2117: 660f6fe2                 movdqa %xmm2, %xmm4
  2121: 660f71d400               psrlw $0, %xmm4
  2126: 660fdbe5                 pand %xmm5, %xmm4
  2130: 660fefdc                 pxor %xmm4, %xmm3
  2134: 660fefc3                 pxor %xmm3, %xmm0
  2138: 660f6fe3                 movdqa %xmm3, %xmm4
  2142: 660f71f401               psllw $1, %xmm4
  2147: 660fefcc                 pxor %xmm4, %xmm1
  2151: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
  2157: 660f6fda                 movdqa %xmm2, %xmm3
  2161: 660f71d301               psrlw $1, %xmm3
  2166: 660fdbdd                 pand %xmm5, %xmm3
  2170: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  2176: 660f6fe2                 movdqa %xmm2, %xmm4
  2180: 660f71d407               psrlw $7, %xmm4
  2185: 660fdbe5                 pand %xmm5, %xmm4
  2189: 660fefdc                 pxor %xmm4, %xmm3
  2193: 660f6fe3                 movdqa %xmm3, %xmm4
  2197: 660f71f405               psllw $5, %xmm4
  2202: 660fefc4                 pxor %xmm4, %xmm0
  2206: 660fefcc                 pxor %xmm4, %xmm1
  2210: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
  2216: 660f6fda                 movdqa %xmm2, %xmm3
  2220: 660f71d302               psrlw $2, %xmm3
  2225: 660fdbdd                 pand %xmm5, %xmm3
  2229: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  2235: 660f6fe2                 movdqa %xmm2, %xmm4
  2239: 660f71d404               psrlw $4, %xmm4
  2244: 660fdbe5                 pand %xmm5, %xmm4
  2248: 660fefdc                 pxor %xmm4, %xmm3
  2252: 660f6fe3                 movdqa %xmm3, %xmm4
  2256: 660f71f406               psllw $6, %xmm4
  2261: 660fefc4                 pxor %xmm4, %xmm0
  2265: 660fefcc                 pxor %xmm4, %xmm1
  2269: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
  2275: 660f6fda                 movdqa %xmm2, %xmm3
  2279: 660f71d307               psrlw $7, %xmm3
  2284: 660fdbdd                 pand %xmm5, %xmm3
  2288: 4d8b5028                 movq 0x28(%r8), %r10
  2292: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  2298: 660f6fe2                 movdqa %xmm2, %xmm4
  2302: 660f71d401               psrlw $1, %xmm4
  2307: 660fdbe5                 pand %xmm5, %xmm4
  2311: 660fefdc                 pxor %xmm4, %xmm3
  2315: 660fefc3                 pxor %xmm3, %xmm0
  2319: 660fefcb                 pxor %xmm3, %xmm1
  2323: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  2329: 660f6fda                 movdqa %xmm2, %xmm3
  2333: 660f71d301               psrlw $1, %xmm3
  2338: 660fdbdd                 pand %xmm5, %xmm3
  2342: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  2349: 660f6fe2                 movdqa %xmm2, %xmm4
  2353: 660f71d403               psrlw $3, %xmm4
  2358: 660fdbe5                 pand %xmm5, %xmm4
  2362: 660fefdc                 pxor %xmm4, %xmm3
  2366: 660f6fe3                 movdqa %xmm3, %xmm4
  2370: 660f71f404               psllw $4, %xmm4
  2375: 660fefc4                 pxor %xmm4, %xmm0
  2379: 660fefcc                 pxor %xmm4, %xmm1
  2383: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  2389: 660f6fda                 movdqa %xmm2, %xmm3
  2393: 660f71d303               psrlw $3, %xmm3
  2398: 660fdbdd                 pand %xmm5, %xmm3
  2402: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  2408: 660f6fe2                 movdqa %xmm2, %xmm4
  2412: 660f71d403               psrlw $3, %xmm4
  2417: 660fdbe5                 pand %xmm5, %xmm4
  2421: 660fefdc                 pxor %xmm4, %xmm3
  2425: 660f6fe3                 movdqa %xmm3, %xmm4
  2429: 660f71f403               psllw $3, %xmm4
  2434: 660fefc4                 pxor %xmm4, %xmm0
  2438: 660f6fe3                 movdqa %xmm3, %xmm4
  2442: 660f71f404               psllw $4, %xmm4
  2447: 660fefcc                 pxor %xmm4, %xmm1
  2451: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  2457: 660f6fda                 movdqa %xmm2, %xmm3
  2461: 660f71d304               psrlw $4, %xmm3
  2466: 660fdbdd                 pand %xmm5, %xmm3
  2470: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  2476: 660f6fe2                 movdqa %xmm2, %xmm4
  2480: 660f71d405               psrlw $5, %xmm4
  2485: 660fdbe5                 pand %xmm5, %xmm4
  2489: 660fefdc                 pxor %xmm4, %xmm3
  2493: 660f6fe3                 movdqa %xmm3, %xmm4
  2497: 660f71f407               psllw $7, %xmm4
  2502: 660fefc4                 pxor %xmm4, %xmm0
  2506: 660fefcc                 pxor %xmm4, %xmm1
  2510: 660f6fda                 movdqa %xmm2, %xmm3
  2514: 660f71d300               psrlw $0, %xmm3
  2519: 660fdbdd                 pand %xmm5, %xmm3
  2523: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  2530: 660f6fe2                 movdqa %xmm2, %xmm4
  2534: 660f71d400               psrlw $0, %xmm4
  2539: 660fdbe5                 pand %xmm5, %xmm4
  2543: 660fefdc                 pxor %xmm4, %xmm3
  2547: 660fefc3                 pxor %xmm3, %xmm0
  2551: 660f6fe3                 movdqa %xmm3, %xmm4
  2555: 660f71f401               psllw $1, %xmm4
  2560: 660fefcc                 pxor %xmm4, %xmm1
  2564: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  2570: 660f6fda                 movdqa %xmm2, %xmm3
  2574: 660f71d301               psrlw $1, %xmm3
  2579: 660fdbdd                 pand %xmm5, %xmm3
  2583: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  2590: 660f6fe2                 movdqa %xmm2, %xmm4
  2594: 660f71d402               psrlw $2, %xmm4
  2599: 660fdbe5                 pand %xmm5, %xmm4
  2603: 660fefdc                 pxor %xmm4, %xmm3
  2607: 660f6fe3                 movdqa %xmm3, %xmm4
  2611: 660f71f403               psllw $3, %xmm4
  2616: 660fefc4                 pxor %xmm4, %xmm0
  2620: 660fefcc                 pxor %xmm4, %xmm1
  2624: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  2630: 660f6fda                 movdqa %xmm2, %xmm3
  2634: 660f71d306               psrlw $6, %xmm3
  2639: 660fdbdd                 pand %xmm5, %xmm3
  2643: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  2650: 660f6fe2                 movdqa %xmm2, %xmm4
  2654: 660f71d406               psrlw $6, %xmm4
  2659: 660fdbe5                 pand %xmm5, %xmm4
  2663: 660fefdc                 pxor %xmm4, %xmm3
  2667: 660f6fe3                 movdqa %xmm3, %xmm4
  2671: 660f71f406               psllw $6, %xmm4
  2676: 660fefc4                 pxor %xmm4, %xmm0
  2680: 660f6fe3                 movdqa %xmm3, %xmm4
  2684: 660f71f407               psllw $7, %xmm4
  2689: 660fefcc                 pxor %xmm4, %xmm1
  2693: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  2699: 660f6fda                 movdqa %xmm2, %xmm3
  2703: 660f71d307               psrlw $7, %xmm3
  2708: 660fdbdd                 pand %xmm5, %xmm3
  2712: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  2719: 660f6fe2                 movdqa %xmm2, %xmm4
  2723: 660f71d401               psrlw $1, %xmm4
  2728: 660fdbe5                 pand %xmm5, %xmm4
  2732: 660fefdc                 pxor %xmm4, %xmm3
  2736: 660f6fe3                 movdqa %xmm3, %xmm4
  2740: 660f71f402               psllw $2, %xmm4
  2745: 660fefc4                 pxor %xmm4, %xmm0
  2749: 660fefcc                 pxor %xmm4, %xmm1
  2753: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
  2759: 660f6fda                 movdqa %xmm2, %xmm3
  2763: 660f71d303               psrlw $3, %xmm3
  2768: 660fdbdd                 pand %xmm5, %xmm3
  2772: 660f6fe3                 movdqa %xmm3, %xmm4
  2776: 660f71f403               psllw $3, %xmm4
  2781: 660fefc4                 pxor %xmm4, %xmm0
  2785: 660f6fda                 movdqa %xmm2, %xmm3
  2789: 660f71d304               psrlw $4, %xmm3
  2794: 660fdbdd                 pand %xmm5, %xmm3
  2798: 660f6fe3                 movdqa %xmm3, %xmm4
  2802: 660f71f405               psllw $5, %xmm4
  2807: 660fefcc                 pxor %xmm4, %xmm1
  2811: 660f6fda                 movdqa %xmm2, %xmm3
  2815: 660f71d305               psrlw $5, %xmm3
  2820: 660fdbdd                 pand %xmm5, %xmm3
  2824: 660f6fe3                 movdqa %xmm3, %xmm4
  2828: 660f71f406               psllw $6, %xmm4
  2833: 660fefcc                 pxor %xmm4, %xmm1
  2837: 660f6fda                 movdqa %xmm2, %xmm3
  2841: 660f71d306               psrlw $6, %xmm3
  2846: 660fdbdd                 pand %xmm5, %xmm3
  2850: 660f6fe3                 movdqa %xmm3, %xmm4
  2854: 660f71f407               psllw $7, %xmm4
  2859: 660fefcc                 pxor %xmm4, %xmm1
  2863: 660f6fda                 movdqa %xmm2, %xmm3
  2867: 660f71d307               psrlw $7, %xmm3
  2872: 660fdbdd                 pand %xmm5, %xmm3
  2876: 660f6fe3                 movdqa %xmm3, %xmm4
  2880: 660f71f407               psllw $7, %xmm4
  2885: 660fefcc                 pxor %xmm4, %xmm1
  2889: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
  2895: 660f6fda                 movdqa %xmm2, %xmm3
  2899: 660f71d305               psrlw $5, %xmm3
  2904: 660fdbdd                 pand %xmm5, %xmm3
  2908: 660f6fe3                 movdqa %xmm3, %xmm4
  2912: 660f71f404               psllw $4, %xmm4
  2917: 660fefc4                 pxor %xmm4, %xmm0
  2921: 660f6fda                 movdqa %xmm2, %xmm3
  2925: 660f71d306               psrlw $6, %xmm3
  2930: 660fdbdd                 pand %xmm5, %xmm3
  2934: 660f6fe3                 movdqa %xmm3, %xmm4
  2938: 660f71f405               psllw $5, %xmm4
  2943: 660fefc4                 pxor %xmm4, %xmm0
  2947: 660f6fda                 movdqa %xmm2, %xmm3
  2951: 660f71d307               psrlw $7, %xmm3
  2956: 660fdbdd                 pand %xmm5, %xmm3
  2960: 660f6fe3                 movdqa %xmm3, %xmm4
  2964: 660f71f406               psllw $6, %xmm4
  2969: 660fefc4                 pxor %xmm4, %xmm0
  2973: 66430f6f141c             movdqa (%r12, %r11), %xmm2
  2979: 660f6fda                 movdqa %xmm2, %xmm3
  2983: 660f71d302               psrlw $2, %xmm3
  2988: 660fdbdd                 pand %xmm5, %xmm3
  2992: 660f6fe3                 movdqa %xmm3, %xmm4
  2996: 660f71f402               psllw $2, %xmm4
  3001: 660fefc4                 pxor %xmm4, %xmm0
  3005: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
  3012: 660f6fda                 movdqa %xmm2, %xmm3
  3016: 660f71d300               psrlw $0, %xmm3
  3021: 660fdbdd                 pand %xmm5, %xmm3
  3025: 660f6fe3                 movdqa %xmm3, %xmm4
  3029: 660f71f401               psllw $1, %xmm4
  3034: 660fefc4                 pxor %xmm4, %xmm0
  3038: 660f6fda                 movdqa %xmm2, %xmm3
  3042: 660f71d301               psrlw $1, %xmm3
  3047: 660fdbdd                 pand %xmm5, %xmm3
  3051: 660f6fe3                 movdqa %xmm3, %xmm4
  3055: 660f71f401               psllw $1, %xmm4
  3060: 660fefc4                 pxor %xmm4, %xmm0
  3064: 660f6fda                 movdqa %xmm2, %xmm3
  3068: 660f71d302               psrlw $2, %xmm3
  3073: 660fdbdd                 pand %xmm5, %xmm3
  3077: 660f6fe3                 movdqa %xmm3, %xmm4
  3081: 660f71f402               psllw $2, %xmm4
  3086: 660fefc4                 pxor %xmm4, %xmm0
  3090: 660f6fda                 movdqa %xmm2, %xmm3
  3094: 660f71d303               psrlw $3, %xmm3
  3099: 660fdbdd                 pand %xmm5, %xmm3
  3103: 660f6fe3                 movdqa %xmm3, %xmm4
  3107: 660f71f403               psllw $3, %xmm4
  3112: 660fefc4                 pxor %xmm4, %xmm0
  3116: 660f6fda                 movdqa %xmm2, %xmm3
  3120: 660f71d304               psrlw $4, %xmm3
  3125: 660fdbdd                 pand %xmm5, %xmm3
  3129: 660f6fe3                 movdqa %xmm3, %xmm4
  3133: 660f71f404               psllw $4, %xmm4
  3138: 660fefc4                 pxor %xmm4, %xmm0
  3142: 660f6fe3                 movdqa %xmm3, %xmm4
  3146: 660f71f405               psllw $5, %xmm4
  3151: 660fefc4                 pxor %xmm4, %xmm0
  3155: 660fefcc                 pxor %xmm4, %xmm1
  3159: 660f6fda                 movdqa %xmm2, %xmm3
  3163: 660f71d306               psrlw $6, %xmm3
  3168: 660fdbdd                 pand %xmm5, %xmm3
  3172: 660f6fe3                 movdqa %xmm3, %xmm4
  3176: 660f71f407               psllw $7, %xmm4
  3181: 660fefc4                 pxor %xmm4, %xmm0
  3185: 66430f6f141e             movdqa (%r14, %r11), %xmm2
  3191: 660f6fda                 movdqa %xmm2, %xmm3
  3195: 660f71d300               psrlw $0, %xmm3
  3200: 660fdbdd                 pand %xmm5, %xmm3
  3204: 660f6fe3                 movdqa %xmm3, %xmm4
  3208: 660f71f401               psllw $1, %xmm4
  3213: 660fefcc                 pxor %xmm4, %xmm1
  3217: 660f6fda                 movdqa %xmm2, %xmm3
  3221: 660f71d301               psrlw $1, %xmm3
  3226: 660fdbdd                 pand %xmm5, %xmm3
  3230: 660fefc3                 pxor %xmm3, %xmm0
  3234: 660f6fda                 movdqa %xmm2, %xmm3
  3238: 660f71d304               psrlw $4, %xmm3
  3243: 660fdbdd                 pand %xmm5, %xmm3
  3247: 660f6fe3                 movdqa %xmm3, %xmm4
  3251: 660f71f404               psllw $4, %xmm4
  3256: 660fefc4                 pxor %xmm4, %xmm0
  3260: 660f6fda                 movdqa %xmm2, %xmm3
  3264: 660f71d305               psrlw $5, %xmm3
  3269: 660fdbdd                 pand %xmm5, %xmm3
  3273: 660f6fe3                 movdqa %xmm3, %xmm4
  3277: 660f71f405               psllw $5, %xmm4
  3282: 660fefc4                 pxor %xmm4, %xmm0
  3286: 660f6fda                 movdqa %xmm2, %xmm3
  3290: 660f71d306               psrlw $6, %xmm3
  3295: 660fdbdd                 pand %xmm5, %xmm3
  3299: 660f6fe3                 movdqa %xmm3, %xmm4
  3303: 660f71f406               psllw $6, %xmm4
  3308: 660fefc4                 pxor %xmm4, %xmm0
  3312: 4d8b5028                 movq 0x28(%r8), %r10
  3316: 66430f6f141a             movdqa (%r10, %r11), %xmm2
  3322: 660f6fda                 movdqa %xmm2, %xmm3
  3326: 660f71d300               psrlw $0, %xmm3
  3331: 660fdbdd                 pand %xmm5, %xmm3
  3335: 660f6fe3                 movdqa %xmm3, %xmm4
  3339: 660f71f402               psllw $2, %xmm4
  3344: 660fefcc                 pxor %xmm4, %xmm1
  3348: 660f6fda                 movdqa %xmm2, %xmm3
  3352: 660f71d302               psrlw $2, %xmm3
  3357: 660fdbdd                 pand %xmm5, %xmm3
  3361: 660f6fe3                 movdqa %xmm3, %xmm4
  3365: 660f71f403               psllw $3, %xmm4
  3370: 660fefc4                 pxor %xmm4, %xmm0
  3374: 660f6fda                 movdqa %xmm2, %xmm3
  3378: 660f71d306               psrlw $6, %xmm3
  3383: 660fdbdd                 pand %xmm5, %xmm3
  3387: 660fefc3                 pxor %xmm3, %xmm0
  3391: 660fefcb                 pxor %xmm3, %xmm1
  3395: 660f6fe3                 movdqa %xmm3, %xmm4
  3399: 660f71f403               psllw $3, %xmm4
  3404: 660fefc4                 pxor %xmm4, %xmm0
  3408: 660fefcc                 pxor %xmm4, %xmm1
  3412: 660f6fda                 movdqa %xmm2, %xmm3
  3416: 660f71d307               psrlw $7, %xmm3
  3421: 660fdbdd                 pand %xmm5, %xmm3
  3425: 660f6fe3                 movdqa %xmm3, %xmm4
  3429: 660f71f401               psllw $1, %xmm4
  3434: 660fefc4                 pxor %xmm4, %xmm0
  3438: 660fefcc                 pxor %xmm4, %xmm1
  3442: 660f6fe3                 movdqa %xmm3, %xmm4
  3446: 660f71f404               psllw $4, %xmm4
  3451: 660fefc4                 pxor %xmm4, %xmm0
  3455: 660fefcc                 pxor %xmm4, %xmm1
  3459: 4d8b11                   movq (%r9), %r10
  3462: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
  3468: 4d8b5108                 movq 8(%r9), %r10
  3472: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
  3478: 4983c310                 addq $0x10, %r11
  3482: 4c3bd8                   cmpq %rax, %r11
  3485: 0f86bdf2ffff             jbe 0x60
  3491: 4c3bd9                   cmpq %rcx, %r11
  3494: 0f839e020000             jae 0x104a
  3500: 4153                     pushq %r11
  3502: 498d7f40                 leaq 0x40(%r15), %rdi
  3506: 498db740010000           leaq 0x140(%r15), %rsi
  3513: 4d8b11                   movq (%r9), %r10
  3516: 4c8b1c24                 movq (%rsp), %r11
  3520: 33db                     xorl %ebx, %ebx
  3522: 498b00                   movq (%r8), %rax
  3525: 420fb61418               movzbl (%rax, %r11), %edx
  3530: 4885d2                   testq %rdx, %rdx
  3533: 0f8422000000             je 0xdf5
  3539: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3543: 81c0c1000000             addl $0xc1, %eax
  3549: 81f8ff000000             cmpl $0xff, %eax
  3555: 0f8206000000             jb 0xdef
  3561: 81c001ffffff             addl $0xffffff01, %eax
  3567: 0fb60406                 movzbl (%rsi, %rax), %eax
  3571: 32d8                     xorb %al, %bl
  3573: 498b4008                 movq 8(%r8), %rax
  3577: 420fb61418               movzbl (%rax, %r11), %edx
  3582: 4885d2                   testq %rdx, %rdx
  3585: 0f841f000000             je 0xe26
  3591: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3595: 83c003                   addl $3, %eax
  3598: 81f8ff000000             cmpl $0xff, %eax
  3604: 0f8206000000             jb 0xe20
  3610: 81c001ffffff             addl $0xffffff01, %eax
  3616: 0fb60406                 movzbl (%rsi, %rax), %eax
  3620: 32d8                     xorb %al, %bl
  3622: 498b4010                 movq 0x10(%r8), %rax
  3626: 420fb61418               movzbl (%rax, %r11), %edx
  3631: 4885d2                   testq %rdx, %rdx
  3634: 0f8422000000             je 0xe5a
  3640: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3644: 81c0c6000000             addl $0xc6, %eax
  3650: 81f8ff000000             cmpl $0xff, %eax
  3656: 0f8206000000             jb 0xe54
  3662: 81c001ffffff             addl $0xffffff01, %eax
  3668: 0fb60406                 movzbl (%rsi, %rax), %eax
  3672: 32d8                     xorb %al, %bl
  3674: 498b4018                 movq 0x18(%r8), %rax
  3678: 420fb61418               movzbl (%rax, %r11), %edx
  3683: 4885d2                   testq %rdx, %rdx
  3686: 0f841f000000             je 0xe8b
  3692: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3696: 83c001                   addl $1, %eax
  3699: 81f8ff000000             cmpl $0xff, %eax
  3705: 0f8206000000             jb 0xe85
  3711: 81c001ffffff             addl $0xffffff01, %eax
  3717: 0fb60406                 movzbl (%rsi, %rax), %eax
  3721: 32d8                     xorb %al, %bl
  3723: 498b4020                 movq 0x20(%r8), %rax
  3727: 420fb61418               movzbl (%rax, %r11), %edx
  3732: 4885d2                   testq %rdx, %rdx
  3735: 0f8422000000             je 0xebf
  3741: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3745: 81c0fe000000             addl $0xfe, %eax
  3751: 81f8ff000000             cmpl $0xff, %eax
  3757: 0f8206000000             jb 0xeb9
  3763: 81c001ffffff             addl $0xffffff01, %eax
  3769: 0fb60406                 movzbl (%rsi, %rax), %eax
  3773: 32d8                     xorb %al, %bl
  3775: 498b4028                 movq 0x28(%r8), %rax
  3779: 420fb61418               movzbl (%rax, %r11), %edx
  3784: 4885d2                   testq %rdx, %rdx
  3787: 0f841f000000             je 0xef0
  3793: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3797: 83c031                   addl $0x31, %eax
  3800: 81f8ff000000             cmpl $0xff, %eax
  3806: 0f8206000000             jb 0xeea
  3812: 81c001ffffff             addl $0xffffff01, %eax
  3818: 0fb60406                 movzbl (%rsi, %rax), %eax
  3822: 32d8                     xorb %al, %bl
  3824: 43881c1a                 movb %bl, (%r10, %r11)
  3828: 33db                     xorl %ebx, %ebx
  3830: 49ffc3                   incq %r11
  3833: 4c3bd9                   cmpq %rcx, %r11
  3836: 0f82c0feffff             jb 0xdc2
  3842: 4d8b5108                 movq 8(%r9), %r10
  3846: 4c8b1c24                 movq (%rsp), %r11
  3850: 33db                     xorl %ebx, %ebx
  3852: 498b00                   movq (%r8), %rax
  3855: 420fb61418               movzbl (%rax, %r11), %edx
  3860: 4885d2                   testq %rdx, %rdx
  3863: 0f841f000000             je 0xf3c
  3869: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3873: 83c01c                   addl $0x1c, %eax
  3876: 81f8ff000000             cmpl $0xff, %eax
  3882: 0f8206000000             jb 0xf36
  3888: 81c001ffffff             addl $0xffffff01, %eax
  3894: 0fb60406                 movzbl (%rsi, %rax), %eax
  3898: 32d8                     xorb %al, %bl
  3900: 498b4008                 movq 8(%r8), %rax
  3904: 420fb61418               movzbl (%rax, %r11), %edx
  3909: 4885d2                   testq %rdx, %rdx
  3912: 0f8422000000             je 0xf70
  3918: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3922: 81c0df000000             addl $0xdf, %eax
  3928: 81f8ff000000             cmpl $0xff, %eax
  3934: 0f8206000000             jb 0xf6a
  3940: 81c001ffffff             addl $0xffffff01, %eax
  3946: 0fb60406                 movzbl (%rsi, %rax), %eax
  3950: 32d8                     xorb %al, %bl
  3952: 498b4010                 movq 0x10(%r8), %rax
  3956: 420fb61418               movzbl (%rax, %r11), %edx
  3961: 4885d2                   testq %rdx, %rdx
  3964: 0f841f000000             je 0xfa1
  3970: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3974: 83c01a                   addl $0x1a, %eax
  3977: 81f8ff000000             cmpl $0xff, %eax
  3983: 0f8206000000             jb 0xf9b
  3989: 81c001ffffff             addl $0xffffff01, %eax
  3995: 0fb60406                 movzbl (%rsi, %rax), %eax
  3999: 32d8                     xorb %al, %bl
  4001: 498b4018                 movq 0x18(%r8), %rax
  4005: 420fb61418               movzbl (%rax, %r11), %edx
  4010: 4885d2                   testq %rdx, %rdx
  4013: 0f841f000000             je 0xfd2
  4019: 0fb60417                 movzbl (%rdi, %rdx), %eax
  4023: 83c019                   addl $0x19, %eax
  4026: 81f8ff000000             cmpl $0xff, %eax
  4032: 0f8206000000             jb 0xfcc
  4038: 81c001ffffff             addl $0xffffff01, %eax
  4044: 0fb60406                 movzbl (%rsi, %rax), %eax
  4048: 32d8                     xorb %al, %bl
  4050: 498b4020                 movq 0x20(%r8), %rax
  4054: 420fb61418               movzbl (%rax, %r11), %edx
  4059: 4885d2                   testq %rdx, %rdx
  4062: 0f841f000000             je 0x1003
  4068: 0fb60417                 movzbl (%rdi, %rdx), %eax
  4072: 83c018                   addl $0x18, %eax
  4075: 81f8ff000000             cmpl $0xff, %eax
  4081: 0f8206000000             jb 0xffd
  4087: 81c001ffffff             addl $0xffffff01, %eax
  4093: 0fb60406                 movzbl (%rsi, %rax), %eax
  4097: 32d8                     xorb %al, %bl
  4099: 498b4028                 movq 0x28(%r8), %rax
  4103: 420fb61418               movzbl (%rax, %r11), %edx
  4108: 4885d2                   testq %rdx, %rdx
  4111: 0f841f000000             je 0x1034
  4117: 0fb60417                 movzbl (%rdi, %rdx), %eax
  4121: 83c031                   addl $0x31, %eax
  4124: 81f8ff000000             cmpl $0xff, %eax
  4130: 0f8206000000             jb 0x102e
  4136: 81c001ffffff             addl $0xffffff01, %eax
  4142: 0fb60406                 movzbl (%rsi, %rax), %eax
  4146: 32d8                     xorb %al, %bl
  4148: 43881c1a                 movb %bl, (%r10, %r11)
  4152: 33db                     xorl %ebx, %ebx
  4154: 49ffc3                   incq %r11
  4157: 4c3bd9                   cmpq %rcx, %r11
  4160: 0f82c6feffff             jb 0xf0c
  4166: 4883c408                 addq $8, %rsp
  4170: 415e                     popq %r14
  4172: 415d                     popq %r13
  4174: 415c                     popq %r12
  4176: 415f                     popq %r15
  4178: 5b                       popq %rbx
  4179: c3                       retq 
  4180: 0f0b                     ud2 
  4182: 0f0b                     ud2 
  4184: 0f0b                     ud2 
  4186: 0f0b                     ud2 
  4188: 0f0b                     ud2 
  4190: 0f0b                     ud2 
  4192: 0f0b                     ud2 
  4194: 0f0b                     ud2 
  4196: 0f0b                     ud2 
  4198: 0f0b                     ud2 
  4200: 0f0b                     ud2 
  4202: 0f0b                     ud2 
  4204: 0f0b                     ud2 
  4206: 0f0b                     ud2 
  4208: 0f0b                     ud2 
  4210: 0f0b                     ud2 
  4212: 0f0b                     ud2 
  4214: 0f0b                     ud2 
  4216: 0f0b                     ud2 
  4218: 0f0b                     ud2 
  4220: 0f0b                     ud2 
  4222: 0f0b                     ud2 
