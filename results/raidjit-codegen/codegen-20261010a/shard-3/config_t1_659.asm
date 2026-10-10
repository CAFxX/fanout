; config_t1_659 target=sse twin=xorbit topo=rs(6,3)/enc op=encode U=1
; code_size=3776 twins=["xorbit"] uops_per_byte=41.75
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da70e0000           leaq 0xea7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66410f6f37               movdqa (%r15), %xmm6
    52: 488bc1                   movq %rcx, %rax
    55: 4883e810                 subq $0x10, %rax
    59: 0f82ab0b0000             jb 0xbec
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66420f6f1c1e             movdqa (%rsi, %r11), %xmm3
   102: 660f6fe3                 movdqa %xmm3, %xmm4
   106: 660f71d400               psrlw $0, %xmm4
   111: 660fdbe6                 pand %xmm6, %xmm4
   115: 66420f6f1c1f             movdqa (%rdi, %r11), %xmm3
   121: 660f6feb                 movdqa %xmm3, %xmm5
   125: 660f71d500               psrlw $0, %xmm5
   130: 660fdbee                 pand %xmm6, %xmm5
   134: 660fefe5                 pxor %xmm5, %xmm4
   138: 660f6fc4                 movdqa %xmm4, %xmm0
   142: 660f6fcc                 movdqa %xmm4, %xmm1
   146: 660f6fec                 movdqa %xmm4, %xmm5
   150: 660f71f503               psllw $3, %xmm5
   155: 660fefcd                 pxor %xmm5, %xmm1
   159: 660f6fd5                 movdqa %xmm5, %xmm2
   163: 660f6fec                 movdqa %xmm4, %xmm5
   167: 660f71f507               psllw $7, %xmm5
   172: 660fefcd                 pxor %xmm5, %xmm1
   176: 660f6fe3                 movdqa %xmm3, %xmm4
   180: 660f71d405               psrlw $5, %xmm4
   185: 660fdbe6                 pand %xmm6, %xmm4
   189: 66430f6f1c1c             movdqa (%r12, %r11), %xmm3
   195: 660f6feb                 movdqa %xmm3, %xmm5
   199: 660f71d507               psrlw $7, %xmm5
   204: 660fdbee                 pand %xmm6, %xmm5
   208: 660fefe5                 pxor %xmm5, %xmm4
   212: 660fefd4                 pxor %xmm4, %xmm2
   216: 660f6fec                 movdqa %xmm4, %xmm5
   220: 660f71f501               psllw $1, %xmm5
   225: 660fefd5                 pxor %xmm5, %xmm2
   229: 660f6fec                 movdqa %xmm4, %xmm5
   233: 660f71f503               psllw $3, %xmm5
   238: 660fefcd                 pxor %xmm5, %xmm1
   242: 660fefd5                 pxor %xmm5, %xmm2
   246: 660f6fec                 movdqa %xmm4, %xmm5
   250: 660f71f504               psllw $4, %xmm5
   255: 660fefd5                 pxor %xmm5, %xmm2
   259: 66420f6f1c1f             movdqa (%rdi, %r11), %xmm3
   265: 660f6fe3                 movdqa %xmm3, %xmm4
   269: 660f71d406               psrlw $6, %xmm4
   274: 660fdbe6                 pand %xmm6, %xmm4
   278: 66430f6f1c1e             movdqa (%r14, %r11), %xmm3
   284: 660f6feb                 movdqa %xmm3, %xmm5
   288: 660f71d507               psrlw $7, %xmm5
   293: 660fdbee                 pand %xmm6, %xmm5
   297: 660fefe5                 pxor %xmm5, %xmm4
   301: 660f6fec                 movdqa %xmm4, %xmm5
   305: 660f71f501               psllw $1, %xmm5
   310: 660fefd5                 pxor %xmm5, %xmm2
   314: 660f6fec                 movdqa %xmm4, %xmm5
   318: 660f71f502               psllw $2, %xmm5
   323: 660fefd5                 pxor %xmm5, %xmm2
   327: 660f6fec                 movdqa %xmm4, %xmm5
   331: 660f71f504               psllw $4, %xmm5
   336: 660fefcd                 pxor %xmm5, %xmm1
   340: 660fefd5                 pxor %xmm5, %xmm2
   344: 660f6fec                 movdqa %xmm4, %xmm5
   348: 660f71f505               psllw $5, %xmm5
   353: 660fefd5                 pxor %xmm5, %xmm2
   357: 66420f6f1c1e             movdqa (%rsi, %r11), %xmm3
   363: 660f6fe3                 movdqa %xmm3, %xmm4
   367: 660f71d401               psrlw $1, %xmm4
   372: 660fdbe6                 pand %xmm6, %xmm4
   376: 66430f6f5c1d00           movdqa (%r13, %r11), %xmm3
   383: 660f6feb                 movdqa %xmm3, %xmm5
   387: 660f71d500               psrlw $0, %xmm5
   392: 660fdbee                 pand %xmm6, %xmm5
   396: 660fefe5                 pxor %xmm5, %xmm4
   400: 660fefd4                 pxor %xmm4, %xmm2
   404: 660f6fec                 movdqa %xmm4, %xmm5
   408: 660f71f502               psllw $2, %xmm5
   413: 660fefd5                 pxor %xmm5, %xmm2
   417: 660f6fec                 movdqa %xmm4, %xmm5
   421: 660f71f503               psllw $3, %xmm5
   426: 660fefd5                 pxor %xmm5, %xmm2
   430: 660f6fec                 movdqa %xmm4, %xmm5
   434: 660f71f507               psllw $7, %xmm5
   439: 660fefd5                 pxor %xmm5, %xmm2
   443: 66430f6f1c1c             movdqa (%r12, %r11), %xmm3
   449: 660f6fe3                 movdqa %xmm3, %xmm4
   453: 660f71d405               psrlw $5, %xmm4
   458: 660fdbe6                 pand %xmm6, %xmm4
   462: 66430f6f5c1d00           movdqa (%r13, %r11), %xmm3
   469: 660f6feb                 movdqa %xmm3, %xmm5
   473: 660f71d506               psrlw $6, %xmm5
   478: 660fdbee                 pand %xmm6, %xmm5
   482: 660fefe5                 pxor %xmm5, %xmm4
   486: 660fefcc                 pxor %xmm4, %xmm1
   490: 660f6fec                 movdqa %xmm4, %xmm5
   494: 660f71f501               psllw $1, %xmm5
   499: 660fefcd                 pxor %xmm5, %xmm1
   503: 660f6fec                 movdqa %xmm4, %xmm5
   507: 660f71f503               psllw $3, %xmm5
   512: 660fefcd                 pxor %xmm5, %xmm1
   516: 660f6fec                 movdqa %xmm4, %xmm5
   520: 660f71f504               psllw $4, %xmm5
   525: 660fefcd                 pxor %xmm5, %xmm1
   529: 66430f6f1c1c             movdqa (%r12, %r11), %xmm3
   535: 660f6fe3                 movdqa %xmm3, %xmm4
   539: 660f71d406               psrlw $6, %xmm4
   544: 660fdbe6                 pand %xmm6, %xmm4
   548: 66430f6f5c1d00           movdqa (%r13, %r11), %xmm3
   555: 660f6feb                 movdqa %xmm3, %xmm5
   559: 660f71d507               psrlw $7, %xmm5
   564: 660fdbee                 pand %xmm6, %xmm5
   568: 660fefe5                 pxor %xmm5, %xmm4
   572: 660f6fec                 movdqa %xmm4, %xmm5
   576: 660f71f501               psllw $1, %xmm5
   581: 660fefcd                 pxor %xmm5, %xmm1
   585: 660f6fec                 movdqa %xmm4, %xmm5
   589: 660f71f502               psllw $2, %xmm5
   594: 660fefcd                 pxor %xmm5, %xmm1
   598: 660f6fec                 movdqa %xmm4, %xmm5
   602: 660f71f504               psllw $4, %xmm5
   607: 660fefcd                 pxor %xmm5, %xmm1
   611: 660f6fec                 movdqa %xmm4, %xmm5
   615: 660f71f505               psllw $5, %xmm5
   620: 660fefcd                 pxor %xmm5, %xmm1
   624: 66420f6f1c1e             movdqa (%rsi, %r11), %xmm3
   630: 660f6fe3                 movdqa %xmm3, %xmm4
   634: 660f71d400               psrlw $0, %xmm4
   639: 660fdbe6                 pand %xmm6, %xmm4
   643: 660f6feb                 movdqa %xmm3, %xmm5
   647: 660f71d503               psrlw $3, %xmm5
   652: 660fdbee                 pand %xmm6, %xmm5
   656: 660fefe5                 pxor %xmm5, %xmm4
   660: 660f6fec                 movdqa %xmm4, %xmm5
   664: 660f71f501               psllw $1, %xmm5
   669: 660fefd5                 pxor %xmm5, %xmm2
   673: 660f6fec                 movdqa %xmm4, %xmm5
   677: 660f71f502               psllw $2, %xmm5
   682: 660fefcd                 pxor %xmm5, %xmm1
   686: 660f6fe3                 movdqa %xmm3, %xmm4
   690: 660f71d401               psrlw $1, %xmm4
   695: 660fdbe6                 pand %xmm6, %xmm4
   699: 66420f6f1c1f             movdqa (%rdi, %r11), %xmm3
   705: 660f6feb                 movdqa %xmm3, %xmm5
   709: 660f71d501               psrlw $1, %xmm5
   714: 660fdbee                 pand %xmm6, %xmm5
   718: 660fefe5                 pxor %xmm5, %xmm4
   722: 660fefcc                 pxor %xmm4, %xmm1
   726: 660f6fec                 movdqa %xmm4, %xmm5
   730: 660f71f501               psllw $1, %xmm5
   735: 660fefc5                 pxor %xmm5, %xmm0
   739: 66420f6f1c1e             movdqa (%rsi, %r11), %xmm3
   745: 660f6fe3                 movdqa %xmm3, %xmm4
   749: 660f71d406               psrlw $6, %xmm4
   754: 660fdbe6                 pand %xmm6, %xmm4
   758: 66430f6f1c1e             movdqa (%r14, %r11), %xmm3
   764: 660f6feb                 movdqa %xmm3, %xmm5
   768: 660f71d506               psrlw $6, %xmm5
   773: 660fdbee                 pand %xmm6, %xmm5
   777: 660fefe5                 pxor %xmm5, %xmm4
   781: 660f6fec                 movdqa %xmm4, %xmm5
   785: 660f71f504               psllw $4, %xmm5
   790: 660fefd5                 pxor %xmm5, %xmm2
   794: 660f6fec                 movdqa %xmm4, %xmm5
   798: 660f71f506               psllw $6, %xmm5
   803: 660fefc5                 pxor %xmm5, %xmm0
   807: 66420f6f1c1e             movdqa (%rsi, %r11), %xmm3
   813: 660f6fe3                 movdqa %xmm3, %xmm4
   817: 660f71d407               psrlw $7, %xmm4
   822: 660fdbe6                 pand %xmm6, %xmm4
   826: 66420f6f1c1f             movdqa (%rdi, %r11), %xmm3
   832: 660f6feb                 movdqa %xmm3, %xmm5
   836: 660f71d507               psrlw $7, %xmm5
   841: 660fdbee                 pand %xmm6, %xmm5
   845: 660fefe5                 pxor %xmm5, %xmm4
   849: 660f6fec                 movdqa %xmm4, %xmm5
   853: 660f71f505               psllw $5, %xmm5
   858: 660fefd5                 pxor %xmm5, %xmm2
   862: 660f6fec                 movdqa %xmm4, %xmm5
   866: 660f71f507               psllw $7, %xmm5
   871: 660fefc5                 pxor %xmm5, %xmm0
   875: 660f6fe3                 movdqa %xmm3, %xmm4
   879: 660f71d401               psrlw $1, %xmm4
   884: 660fdbe6                 pand %xmm6, %xmm4
   888: 66430f6f5c1d00           movdqa (%r13, %r11), %xmm3
   895: 660f6feb                 movdqa %xmm3, %xmm5
   899: 660f71d505               psrlw $5, %xmm5
   904: 660fdbee                 pand %xmm6, %xmm5
   908: 660fefe5                 pxor %xmm5, %xmm4
   912: 660f6fec                 movdqa %xmm4, %xmm5
   916: 660f71f504               psllw $4, %xmm5
   921: 660fefd5                 pxor %xmm5, %xmm2
   925: 660f6fec                 movdqa %xmm4, %xmm5
   929: 660f71f507               psllw $7, %xmm5
   934: 660fefcd                 pxor %xmm5, %xmm1
   938: 66420f6f1c1f             movdqa (%rdi, %r11), %xmm3
   944: 660f6fe3                 movdqa %xmm3, %xmm4
   948: 660f71d407               psrlw $7, %xmm4
   953: 660fdbe6                 pand %xmm6, %xmm4
   957: 66430f6f1c1c             movdqa (%r12, %r11), %xmm3
   963: 660f6feb                 movdqa %xmm3, %xmm5
   967: 660f71d502               psrlw $2, %xmm5
   972: 660fdbee                 pand %xmm6, %xmm5
   976: 660fefe5                 pxor %xmm5, %xmm4
   980: 660f6fec                 movdqa %xmm4, %xmm5
   984: 660f71f503               psllw $3, %xmm5
   989: 660fefd5                 pxor %xmm5, %xmm2
   993: 660f6fec                 movdqa %xmm4, %xmm5
   997: 660f71f505               psllw $5, %xmm5
  1002: 660fefcd                 pxor %xmm5, %xmm1
  1006: 66420f6f1c1e             movdqa (%rsi, %r11), %xmm3
  1012: 660f6fe3                 movdqa %xmm3, %xmm4
  1016: 660f71d400               psrlw $0, %xmm4
  1021: 660fdbe6                 pand %xmm6, %xmm4
  1025: 660fefd4                 pxor %xmm4, %xmm2
  1029: 660f6fec                 movdqa %xmm4, %xmm5
  1033: 660f71f506               psllw $6, %xmm5
  1038: 660fefd5                 pxor %xmm5, %xmm2
  1042: 660f6fec                 movdqa %xmm4, %xmm5
  1046: 660f71f507               psllw $7, %xmm5
  1051: 660fefd5                 pxor %xmm5, %xmm2
  1055: 660f6fe3                 movdqa %xmm3, %xmm4
  1059: 660f71d402               psrlw $2, %xmm4
  1064: 660fdbe6                 pand %xmm6, %xmm4
  1068: 660fefd4                 pxor %xmm4, %xmm2
  1072: 660f6fec                 movdqa %xmm4, %xmm5
  1076: 660f71f501               psllw $1, %xmm5
  1081: 660fefcd                 pxor %xmm5, %xmm1
  1085: 660f6fec                 movdqa %xmm4, %xmm5
  1089: 660f71f502               psllw $2, %xmm5
  1094: 660fefc5                 pxor %xmm5, %xmm0
  1098: 660f6fe3                 movdqa %xmm3, %xmm4
  1102: 660f71d403               psrlw $3, %xmm4
  1107: 660fdbe6                 pand %xmm6, %xmm4
  1111: 660f6fec                 movdqa %xmm4, %xmm5
  1115: 660f71f503               psllw $3, %xmm5
  1120: 660fefc5                 pxor %xmm5, %xmm0
  1124: 660f6fe3                 movdqa %xmm3, %xmm4
  1128: 660f71d404               psrlw $4, %xmm4
  1133: 660fdbe6                 pand %xmm6, %xmm4
  1137: 660f6fec                 movdqa %xmm4, %xmm5
  1141: 660f71f502               psllw $2, %xmm5
  1146: 660fefd5                 pxor %xmm5, %xmm2
  1150: 660f6fec                 movdqa %xmm4, %xmm5
  1154: 660f71f503               psllw $3, %xmm5
  1159: 660fefcd                 pxor %xmm5, %xmm1
  1163: 660f6fec                 movdqa %xmm4, %xmm5
  1167: 660f71f504               psllw $4, %xmm5
  1172: 660fefc5                 pxor %xmm5, %xmm0
  1176: 660f6fe3                 movdqa %xmm3, %xmm4
  1180: 660f71d405               psrlw $5, %xmm4
  1185: 660fdbe6                 pand %xmm6, %xmm4
  1189: 660f6fec                 movdqa %xmm4, %xmm5
  1193: 660f71f503               psllw $3, %xmm5
  1198: 660fefd5                 pxor %xmm5, %xmm2
  1202: 660f6fec                 movdqa %xmm4, %xmm5
  1206: 660f71f504               psllw $4, %xmm5
  1211: 660fefcd                 pxor %xmm5, %xmm1
  1215: 660f6fec                 movdqa %xmm4, %xmm5
  1219: 660f71f505               psllw $5, %xmm5
  1224: 660fefc5                 pxor %xmm5, %xmm0
  1228: 660f6fe3                 movdqa %xmm3, %xmm4
  1232: 660f71d406               psrlw $6, %xmm4
  1237: 660fdbe6                 pand %xmm6, %xmm4
  1241: 660f6fec                 movdqa %xmm4, %xmm5
  1245: 660f71f505               psllw $5, %xmm5
  1250: 660fefcd                 pxor %xmm5, %xmm1
  1254: 660f6fe3                 movdqa %xmm3, %xmm4
  1258: 660f71d407               psrlw $7, %xmm4
  1263: 660fdbe6                 pand %xmm6, %xmm4
  1267: 660f6fec                 movdqa %xmm4, %xmm5
  1271: 660f71f506               psllw $6, %xmm5
  1276: 660fefcd                 pxor %xmm5, %xmm1
  1280: 66420f6f1c1f             movdqa (%rdi, %r11), %xmm3
  1286: 660f6fe3                 movdqa %xmm3, %xmm4
  1290: 660f71d400               psrlw $0, %xmm4
  1295: 660fdbe6                 pand %xmm6, %xmm4
  1299: 660f6fec                 movdqa %xmm4, %xmm5
  1303: 660f71f501               psllw $1, %xmm5
  1308: 660fefcd                 pxor %xmm5, %xmm1
  1312: 660f6fec                 movdqa %xmm4, %xmm5
  1316: 660f71f506               psllw $6, %xmm5
  1321: 660fefcd                 pxor %xmm5, %xmm1
  1325: 660f6fe3                 movdqa %xmm3, %xmm4
  1329: 660f71d401               psrlw $1, %xmm4
  1334: 660fdbe6                 pand %xmm6, %xmm4
  1338: 660f6fec                 movdqa %xmm4, %xmm5
  1342: 660f71f502               psllw $2, %xmm5
  1347: 660fefcd                 pxor %xmm5, %xmm1
  1351: 660f6fec                 movdqa %xmm4, %xmm5
  1355: 660f71f503               psllw $3, %xmm5
  1360: 660fefcd                 pxor %xmm5, %xmm1
  1364: 660f6fe3                 movdqa %xmm3, %xmm4
  1368: 660f71d402               psrlw $2, %xmm4
  1373: 660fdbe6                 pand %xmm6, %xmm4
  1377: 660fefcc                 pxor %xmm4, %xmm1
  1381: 660f6fec                 movdqa %xmm4, %xmm5
  1385: 660f71f502               psllw $2, %xmm5
  1390: 660fefc5                 pxor %xmm5, %xmm0
  1394: 660f6fec                 movdqa %xmm4, %xmm5
  1398: 660f71f505               psllw $5, %xmm5
  1403: 660fefd5                 pxor %xmm5, %xmm2
  1407: 660f6fe3                 movdqa %xmm3, %xmm4
  1411: 660f71d403               psrlw $3, %xmm4
  1416: 660fdbe6                 pand %xmm6, %xmm4
  1420: 660f6fec                 movdqa %xmm4, %xmm5
  1424: 660f71f501               psllw $1, %xmm5
  1429: 660fefcd                 pxor %xmm5, %xmm1
  1433: 660f6fec                 movdqa %xmm4, %xmm5
  1437: 660f71f503               psllw $3, %xmm5
  1442: 660fefc5                 pxor %xmm5, %xmm0
  1446: 660f6fec                 movdqa %xmm4, %xmm5
  1450: 660f71f506               psllw $6, %xmm5
  1455: 660fefd5                 pxor %xmm5, %xmm2
  1459: 660f6fe3                 movdqa %xmm3, %xmm4
  1463: 660f71d404               psrlw $4, %xmm4
  1468: 660fdbe6                 pand %xmm6, %xmm4
  1472: 660f6fec                 movdqa %xmm4, %xmm5
  1476: 660f71f502               psllw $2, %xmm5
  1481: 660fefcd                 pxor %xmm5, %xmm1
  1485: 660f6fec                 movdqa %xmm4, %xmm5
  1489: 660f71f504               psllw $4, %xmm5
  1494: 660fefc5                 pxor %xmm5, %xmm0
  1498: 660f6fec                 movdqa %xmm4, %xmm5
  1502: 660f71f507               psllw $7, %xmm5
  1507: 660fefd5                 pxor %xmm5, %xmm2
  1511: 660f6fe3                 movdqa %xmm3, %xmm4
  1515: 660f71d405               psrlw $5, %xmm4
  1520: 660fdbe6                 pand %xmm6, %xmm4
  1524: 660f6fec                 movdqa %xmm4, %xmm5
  1528: 660f71f505               psllw $5, %xmm5
  1533: 660fefc5                 pxor %xmm5, %xmm0
  1537: 660f6fe3                 movdqa %xmm3, %xmm4
  1541: 660f71d406               psrlw $6, %xmm4
  1546: 660fdbe6                 pand %xmm6, %xmm4
  1550: 660f6fec                 movdqa %xmm4, %xmm5
  1554: 660f71f506               psllw $6, %xmm5
  1559: 660fefc5                 pxor %xmm5, %xmm0
  1563: 660f6fe3                 movdqa %xmm3, %xmm4
  1567: 660f71d407               psrlw $7, %xmm4
  1572: 660fdbe6                 pand %xmm6, %xmm4
  1576: 660f6fec                 movdqa %xmm4, %xmm5
  1580: 660f71f502               psllw $2, %xmm5
  1585: 660fefd5                 pxor %xmm5, %xmm2
  1589: 660f6fec                 movdqa %xmm4, %xmm5
  1593: 660f71f506               psllw $6, %xmm5
  1598: 660fefd5                 pxor %xmm5, %xmm2
  1602: 66430f6f1c1c             movdqa (%r12, %r11), %xmm3
  1608: 660f6fe3                 movdqa %xmm3, %xmm4
  1612: 660f71d400               psrlw $0, %xmm4
  1617: 660fdbe6                 pand %xmm6, %xmm4
  1621: 660fefc4                 pxor %xmm4, %xmm0
  1625: 660f6fec                 movdqa %xmm4, %xmm5
  1629: 660f71f501               psllw $1, %xmm5
  1634: 660fefd5                 pxor %xmm5, %xmm2
  1638: 660f6fec                 movdqa %xmm4, %xmm5
  1642: 660f71f503               psllw $3, %xmm5
  1647: 660fefcd                 pxor %xmm5, %xmm1
  1651: 660f6fe3                 movdqa %xmm3, %xmm4
  1655: 660f71d401               psrlw $1, %xmm4
  1660: 660fdbe6                 pand %xmm6, %xmm4
  1664: 660f6fec                 movdqa %xmm4, %xmm5
  1668: 660f71f501               psllw $1, %xmm5
  1673: 660fefc5                 pxor %xmm5, %xmm0
  1677: 660f6fec                 movdqa %xmm4, %xmm5
  1681: 660f71f502               psllw $2, %xmm5
  1686: 660fefd5                 pxor %xmm5, %xmm2
  1690: 660f6fec                 movdqa %xmm4, %xmm5
  1694: 660f71f504               psllw $4, %xmm5
  1699: 660fefcd                 pxor %xmm5, %xmm1
  1703: 660f6fe3                 movdqa %xmm3, %xmm4
  1707: 660f71d402               psrlw $2, %xmm4
  1712: 660fdbe6                 pand %xmm6, %xmm4
  1716: 660f6fec                 movdqa %xmm4, %xmm5
  1720: 660f71f502               psllw $2, %xmm5
  1725: 660fefc5                 pxor %xmm5, %xmm0
  1729: 660f6fe3                 movdqa %xmm3, %xmm4
  1733: 660f71d403               psrlw $3, %xmm4
  1738: 660fdbe6                 pand %xmm6, %xmm4
  1742: 660f6fec                 movdqa %xmm4, %xmm5
  1746: 660f71f503               psllw $3, %xmm5
  1751: 660fefc5                 pxor %xmm5, %xmm0
  1755: 660f6fec                 movdqa %xmm4, %xmm5
  1759: 660f71f504               psllw $4, %xmm5
  1764: 660fefd5                 pxor %xmm5, %xmm2
  1768: 660f6fec                 movdqa %xmm4, %xmm5
  1772: 660f71f506               psllw $6, %xmm5
  1777: 660fefcd                 pxor %xmm5, %xmm1
  1781: 660f6fe3                 movdqa %xmm3, %xmm4
  1785: 660f71d404               psrlw $4, %xmm4
  1790: 660fdbe6                 pand %xmm6, %xmm4
  1794: 660f6fec                 movdqa %xmm4, %xmm5
  1798: 660f71f504               psllw $4, %xmm5
  1803: 660fefc5                 pxor %xmm5, %xmm0
  1807: 660f6fec                 movdqa %xmm4, %xmm5
  1811: 660f71f505               psllw $5, %xmm5
  1816: 660fefd5                 pxor %xmm5, %xmm2
  1820: 660f6fec                 movdqa %xmm4, %xmm5
  1824: 660f71f507               psllw $7, %xmm5
  1829: 660fefcd                 pxor %xmm5, %xmm1
  1833: 660f6fe3                 movdqa %xmm3, %xmm4
  1837: 660f71d405               psrlw $5, %xmm4
  1842: 660fdbe6                 pand %xmm6, %xmm4
  1846: 660f6fec                 movdqa %xmm4, %xmm5
  1850: 660f71f505               psllw $5, %xmm5
  1855: 660fefc5                 pxor %xmm5, %xmm0
  1859: 660f6fec                 movdqa %xmm4, %xmm5
  1863: 660f71f506               psllw $6, %xmm5
  1868: 660fefd5                 pxor %xmm5, %xmm2
  1872: 660f6fe3                 movdqa %xmm3, %xmm4
  1876: 660f71d406               psrlw $6, %xmm4
  1881: 660fdbe6                 pand %xmm6, %xmm4
  1885: 660f6fec                 movdqa %xmm4, %xmm5
  1889: 660f71f506               psllw $6, %xmm5
  1894: 660fefc5                 pxor %xmm5, %xmm0
  1898: 660f6fec                 movdqa %xmm4, %xmm5
  1902: 660f71f507               psllw $7, %xmm5
  1907: 660fefd5                 pxor %xmm5, %xmm2
  1911: 660f6fe3                 movdqa %xmm3, %xmm4
  1915: 660f71d407               psrlw $7, %xmm4
  1920: 660fdbe6                 pand %xmm6, %xmm4
  1924: 660f6fec                 movdqa %xmm4, %xmm5
  1928: 660f71f502               psllw $2, %xmm5
  1933: 660fefcd                 pxor %xmm5, %xmm1
  1937: 660f6fec                 movdqa %xmm4, %xmm5
  1941: 660f71f505               psllw $5, %xmm5
  1946: 660fefcd                 pxor %xmm5, %xmm1
  1950: 660f6fec                 movdqa %xmm4, %xmm5
  1954: 660f71f506               psllw $6, %xmm5
  1959: 660fefcd                 pxor %xmm5, %xmm1
  1963: 660f6fec                 movdqa %xmm4, %xmm5
  1967: 660f71f507               psllw $7, %xmm5
  1972: 660fefc5                 pxor %xmm5, %xmm0
  1976: 66430f6f5c1d00           movdqa (%r13, %r11), %xmm3
  1983: 660f6fe3                 movdqa %xmm3, %xmm4
  1987: 660f71d400               psrlw $0, %xmm4
  1992: 660fdbe6                 pand %xmm6, %xmm4
  1996: 660fefc4                 pxor %xmm4, %xmm0
  2000: 660f6fec                 movdqa %xmm4, %xmm5
  2004: 660f71f502               psllw $2, %xmm5
  2009: 660fefcd                 pxor %xmm5, %xmm1
  2013: 660f6fe3                 movdqa %xmm3, %xmm4
  2017: 660f71d401               psrlw $1, %xmm4
  2022: 660fdbe6                 pand %xmm6, %xmm4
  2026: 660fefd4                 pxor %xmm4, %xmm2
  2030: 660f6fec                 movdqa %xmm4, %xmm5
  2034: 660f71f501               psllw $1, %xmm5
  2039: 660fefc5                 pxor %xmm5, %xmm0
  2043: 660f6fec                 movdqa %xmm4, %xmm5
  2047: 660f71f503               psllw $3, %xmm5
  2052: 660fefcd                 pxor %xmm5, %xmm1
  2056: 660f6fe3                 movdqa %xmm3, %xmm4
  2060: 660f71d402               psrlw $2, %xmm4
  2065: 660fdbe6                 pand %xmm6, %xmm4
  2069: 660f6fec                 movdqa %xmm4, %xmm5
  2073: 660f71f501               psllw $1, %xmm5
  2078: 660fefd5                 pxor %xmm5, %xmm2
  2082: 660f6fec                 movdqa %xmm4, %xmm5
  2086: 660f71f502               psllw $2, %xmm5
  2091: 660fefc5                 pxor %xmm5, %xmm0
  2095: 660f6fec                 movdqa %xmm4, %xmm5
  2099: 660f71f504               psllw $4, %xmm5
  2104: 660fefcd                 pxor %xmm5, %xmm1
  2108: 660f6fe3                 movdqa %xmm3, %xmm4
  2112: 660f71d403               psrlw $3, %xmm4
  2117: 660fdbe6                 pand %xmm6, %xmm4
  2121: 660f6fec                 movdqa %xmm4, %xmm5
  2125: 660f71f502               psllw $2, %xmm5
  2130: 660fefd5                 pxor %xmm5, %xmm2
  2134: 660f6fec                 movdqa %xmm4, %xmm5
  2138: 660f71f503               psllw $3, %xmm5
  2143: 660fefc5                 pxor %xmm5, %xmm0
  2147: 660f6fec                 movdqa %xmm4, %xmm5
  2151: 660f71f505               psllw $5, %xmm5
  2156: 660fefcd                 pxor %xmm5, %xmm1
  2160: 660f6fe3                 movdqa %xmm3, %xmm4
  2164: 660f71d404               psrlw $4, %xmm4
  2169: 660fdbe6                 pand %xmm6, %xmm4
  2173: 660f6fec                 movdqa %xmm4, %xmm5
  2177: 660f71f503               psllw $3, %xmm5
  2182: 660fefd5                 pxor %xmm5, %xmm2
  2186: 660f6fec                 movdqa %xmm4, %xmm5
  2190: 660f71f504               psllw $4, %xmm5
  2195: 660fefc5                 pxor %xmm5, %xmm0
  2199: 660f6fec                 movdqa %xmm4, %xmm5
  2203: 660f71f506               psllw $6, %xmm5
  2208: 660fefcd                 pxor %xmm5, %xmm1
  2212: 660f6fe3                 movdqa %xmm3, %xmm4
  2216: 660f71d405               psrlw $5, %xmm4
  2221: 660fdbe6                 pand %xmm6, %xmm4
  2225: 660f6fec                 movdqa %xmm4, %xmm5
  2229: 660f71f505               psllw $5, %xmm5
  2234: 660fefc5                 pxor %xmm5, %xmm0
  2238: 660f6fe3                 movdqa %xmm3, %xmm4
  2242: 660f71d406               psrlw $6, %xmm4
  2247: 660fdbe6                 pand %xmm6, %xmm4
  2251: 660f6fec                 movdqa %xmm4, %xmm5
  2255: 660f71f505               psllw $5, %xmm5
  2260: 660fefd5                 pxor %xmm5, %xmm2
  2264: 660f6fec                 movdqa %xmm4, %xmm5
  2268: 660f71f506               psllw $6, %xmm5
  2273: 660fefc5                 pxor %xmm5, %xmm0
  2277: 660f6fe3                 movdqa %xmm3, %xmm4
  2281: 660f71d407               psrlw $7, %xmm4
  2286: 660fdbe6                 pand %xmm6, %xmm4
  2290: 660f6fec                 movdqa %xmm4, %xmm5
  2294: 660f71f506               psllw $6, %xmm5
  2299: 660fefd5                 pxor %xmm5, %xmm2
  2303: 660f6fec                 movdqa %xmm4, %xmm5
  2307: 660f71f507               psllw $7, %xmm5
  2312: 660fefc5                 pxor %xmm5, %xmm0
  2316: 66430f6f1c1e             movdqa (%r14, %r11), %xmm3
  2322: 660f6fe3                 movdqa %xmm3, %xmm4
  2326: 660f71d400               psrlw $0, %xmm4
  2331: 660fdbe6                 pand %xmm6, %xmm4
  2335: 660fefc4                 pxor %xmm4, %xmm0
  2339: 660f6fec                 movdqa %xmm4, %xmm5
  2343: 660f71f501               psllw $1, %xmm5
  2348: 660fefcd                 pxor %xmm5, %xmm1
  2352: 660f6fec                 movdqa %xmm4, %xmm5
  2356: 660f71f502               psllw $2, %xmm5
  2361: 660fefd5                 pxor %xmm5, %xmm2
  2365: 660f6fe3                 movdqa %xmm3, %xmm4
  2369: 660f71d401               psrlw $1, %xmm4
  2374: 660fdbe6                 pand %xmm6, %xmm4
  2378: 660f6fec                 movdqa %xmm4, %xmm5
  2382: 660f71f501               psllw $1, %xmm5
  2387: 660fefc5                 pxor %xmm5, %xmm0
  2391: 660f6fec                 movdqa %xmm4, %xmm5
  2395: 660f71f502               psllw $2, %xmm5
  2400: 660fefcd                 pxor %xmm5, %xmm1
  2404: 660f6fec                 movdqa %xmm4, %xmm5
  2408: 660f71f503               psllw $3, %xmm5
  2413: 660fefd5                 pxor %xmm5, %xmm2
  2417: 660f6fe3                 movdqa %xmm3, %xmm4
  2421: 660f71d402               psrlw $2, %xmm4
  2426: 660fdbe6                 pand %xmm6, %xmm4
  2430: 660f6fec                 movdqa %xmm4, %xmm5
  2434: 660f71f502               psllw $2, %xmm5
  2439: 660fefc5                 pxor %xmm5, %xmm0
  2443: 660f6fec                 movdqa %xmm4, %xmm5
  2447: 660f71f503               psllw $3, %xmm5
  2452: 660fefcd                 pxor %xmm5, %xmm1
  2456: 660f6fec                 movdqa %xmm4, %xmm5
  2460: 660f71f504               psllw $4, %xmm5
  2465: 660fefd5                 pxor %xmm5, %xmm2
  2469: 660f6fe3                 movdqa %xmm3, %xmm4
  2473: 660f71d403               psrlw $3, %xmm4
  2478: 660fdbe6                 pand %xmm6, %xmm4
  2482: 660f6fec                 movdqa %xmm4, %xmm5
  2486: 660f71f503               psllw $3, %xmm5
  2491: 660fefc5                 pxor %xmm5, %xmm0
  2495: 660f6fec                 movdqa %xmm4, %xmm5
  2499: 660f71f504               psllw $4, %xmm5
  2504: 660fefcd                 pxor %xmm5, %xmm1
  2508: 660f6fec                 movdqa %xmm4, %xmm5
  2512: 660f71f505               psllw $5, %xmm5
  2517: 660fefd5                 pxor %xmm5, %xmm2
  2521: 660f6fe3                 movdqa %xmm3, %xmm4
  2525: 660f71d404               psrlw $4, %xmm4
  2530: 660fdbe6                 pand %xmm6, %xmm4
  2534: 660f6fec                 movdqa %xmm4, %xmm5
  2538: 660f71f504               psllw $4, %xmm5
  2543: 660fefc5                 pxor %xmm5, %xmm0
  2547: 660f6fec                 movdqa %xmm4, %xmm5
  2551: 660f71f505               psllw $5, %xmm5
  2556: 660fefcd                 pxor %xmm5, %xmm1
  2560: 660f6fec                 movdqa %xmm4, %xmm5
  2564: 660f71f506               psllw $6, %xmm5
  2569: 660fefd5                 pxor %xmm5, %xmm2
  2573: 660f6fe3                 movdqa %xmm3, %xmm4
  2577: 660f71d405               psrlw $5, %xmm4
  2582: 660fdbe6                 pand %xmm6, %xmm4
  2586: 660f6fec                 movdqa %xmm4, %xmm5
  2590: 660f71f505               psllw $5, %xmm5
  2595: 660fefc5                 pxor %xmm5, %xmm0
  2599: 660f6fec                 movdqa %xmm4, %xmm5
  2603: 660f71f506               psllw $6, %xmm5
  2608: 660fefcd                 pxor %xmm5, %xmm1
  2612: 660f6fec                 movdqa %xmm4, %xmm5
  2616: 660f71f507               psllw $7, %xmm5
  2621: 660fefd5                 pxor %xmm5, %xmm2
  2625: 660f6fe3                 movdqa %xmm3, %xmm4
  2629: 660f71d406               psrlw $6, %xmm4
  2634: 660fdbe6                 pand %xmm6, %xmm4
  2638: 660fefd4                 pxor %xmm4, %xmm2
  2642: 660f6fec                 movdqa %xmm4, %xmm5
  2646: 660f71f501               psllw $1, %xmm5
  2651: 660fefd5                 pxor %xmm5, %xmm2
  2655: 660f6fec                 movdqa %xmm4, %xmm5
  2659: 660f71f503               psllw $3, %xmm5
  2664: 660fefd5                 pxor %xmm5, %xmm2
  2668: 660f6fec                 movdqa %xmm4, %xmm5
  2672: 660f71f507               psllw $7, %xmm5
  2677: 660fefcd                 pxor %xmm5, %xmm1
  2681: 660f6fe3                 movdqa %xmm3, %xmm4
  2685: 660f71d407               psrlw $7, %xmm4
  2690: 660fdbe6                 pand %xmm6, %xmm4
  2694: 660fefcc                 pxor %xmm4, %xmm1
  2698: 660f6fec                 movdqa %xmm4, %xmm5
  2702: 660f71f501               psllw $1, %xmm5
  2707: 660fefcd                 pxor %xmm5, %xmm1
  2711: 660f6fec                 movdqa %xmm4, %xmm5
  2715: 660f71f503               psllw $3, %xmm5
  2720: 660fefcd                 pxor %xmm5, %xmm1
  2724: 660f6fec                 movdqa %xmm4, %xmm5
  2728: 660f71f507               psllw $7, %xmm5
  2733: 660fefc5                 pxor %xmm5, %xmm0
  2737: 4d8b5028                 movq 0x28(%r8), %r10
  2741: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
  2747: 660f6fe3                 movdqa %xmm3, %xmm4
  2751: 660f71d400               psrlw $0, %xmm4
  2756: 660fdbe6                 pand %xmm6, %xmm4
  2760: 660fefc4                 pxor %xmm4, %xmm0
  2764: 660fefcc                 pxor %xmm4, %xmm1
  2768: 660fefd4                 pxor %xmm4, %xmm2
  2772: 660f6fe3                 movdqa %xmm3, %xmm4
  2776: 660f71d401               psrlw $1, %xmm4
  2781: 660fdbe6                 pand %xmm6, %xmm4
  2785: 660f6fec                 movdqa %xmm4, %xmm5
  2789: 660f71f501               psllw $1, %xmm5
  2794: 660fefc5                 pxor %xmm5, %xmm0
  2798: 660fefcd                 pxor %xmm5, %xmm1
  2802: 660fefd5                 pxor %xmm5, %xmm2
  2806: 660f6fe3                 movdqa %xmm3, %xmm4
  2810: 660f71d402               psrlw $2, %xmm4
  2815: 660fdbe6                 pand %xmm6, %xmm4
  2819: 660f6fec                 movdqa %xmm4, %xmm5
  2823: 660f71f502               psllw $2, %xmm5
  2828: 660fefc5                 pxor %xmm5, %xmm0
  2832: 660fefcd                 pxor %xmm5, %xmm1
  2836: 660fefd5                 pxor %xmm5, %xmm2
  2840: 660f6fe3                 movdqa %xmm3, %xmm4
  2844: 660f71d403               psrlw $3, %xmm4
  2849: 660fdbe6                 pand %xmm6, %xmm4
  2853: 660f6fec                 movdqa %xmm4, %xmm5
  2857: 660f71f503               psllw $3, %xmm5
  2862: 660fefc5                 pxor %xmm5, %xmm0
  2866: 660fefcd                 pxor %xmm5, %xmm1
  2870: 660fefd5                 pxor %xmm5, %xmm2
  2874: 660f6fe3                 movdqa %xmm3, %xmm4
  2878: 660f71d404               psrlw $4, %xmm4
  2883: 660fdbe6                 pand %xmm6, %xmm4
  2887: 660f6fec                 movdqa %xmm4, %xmm5
  2891: 660f71f504               psllw $4, %xmm5
  2896: 660fefc5                 pxor %xmm5, %xmm0
  2900: 660fefcd                 pxor %xmm5, %xmm1
  2904: 660fefd5                 pxor %xmm5, %xmm2
  2908: 660f6fe3                 movdqa %xmm3, %xmm4
  2912: 660f71d405               psrlw $5, %xmm4
  2917: 660fdbe6                 pand %xmm6, %xmm4
  2921: 660f6fec                 movdqa %xmm4, %xmm5
  2925: 660f71f505               psllw $5, %xmm5
  2930: 660fefc5                 pxor %xmm5, %xmm0
  2934: 660fefcd                 pxor %xmm5, %xmm1
  2938: 660fefd5                 pxor %xmm5, %xmm2
  2942: 660f6fe3                 movdqa %xmm3, %xmm4
  2946: 660f71d406               psrlw $6, %xmm4
  2951: 660fdbe6                 pand %xmm6, %xmm4
  2955: 660f6fec                 movdqa %xmm4, %xmm5
  2959: 660f71f506               psllw $6, %xmm5
  2964: 660fefc5                 pxor %xmm5, %xmm0
  2968: 660fefcd                 pxor %xmm5, %xmm1
  2972: 660fefd5                 pxor %xmm5, %xmm2
  2976: 660f6fe3                 movdqa %xmm3, %xmm4
  2980: 660f71d407               psrlw $7, %xmm4
  2985: 660fdbe6                 pand %xmm6, %xmm4
  2989: 660f6fec                 movdqa %xmm4, %xmm5
  2993: 660f71f507               psllw $7, %xmm5
  2998: 660fefc5                 pxor %xmm5, %xmm0
  3002: 660fefcd                 pxor %xmm5, %xmm1
  3006: 660fefd5                 pxor %xmm5, %xmm2
  3010: 4d8b11                   movq (%r9), %r10
  3013: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
  3019: 4d8b5108                 movq 8(%r9), %r10
  3023: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
  3029: 4d8b5110                 movq 0x10(%r9), %r10
  3033: 66430f7f141a             movdqa %xmm2, (%r10, %r11)
  3039: 4983c310                 addq $0x10, %r11
  3043: 4c3bd8                   cmpq %rax, %r11
  3046: 0f8674f4ffff             jbe 0x60
  3052: 4c3bd9                   cmpq %rcx, %r11
  3055: 0f83af020000             jae 0xea4
  3061: 4153                     pushq %r11
  3063: 498d7f40                 leaq 0x40(%r15), %rdi
  3067: 498db740010000           leaq 0x140(%r15), %rsi
  3074: 4d8b11                   movq (%r9), %r10
  3077: 4c8b1c24                 movq (%rsp), %r11
  3081: 33db                     xorl %ebx, %ebx
  3083: 498b00                   movq (%r8), %rax
  3086: 420fb61418               movzbl (%rax, %r11), %edx
  3091: 32da                     xorb %dl, %bl
  3093: 498b4008                 movq 8(%r8), %rax
  3097: 420fb61418               movzbl (%rax, %r11), %edx
  3102: 32da                     xorb %dl, %bl
  3104: 498b4010                 movq 0x10(%r8), %rax
  3108: 420fb61418               movzbl (%rax, %r11), %edx
  3113: 32da                     xorb %dl, %bl
  3115: 498b4018                 movq 0x18(%r8), %rax
  3119: 420fb61418               movzbl (%rax, %r11), %edx
  3124: 32da                     xorb %dl, %bl
  3126: 498b4020                 movq 0x20(%r8), %rax
  3130: 420fb61418               movzbl (%rax, %r11), %edx
  3135: 32da                     xorb %dl, %bl
  3137: 498b4028                 movq 0x28(%r8), %rax
  3141: 420fb61418               movzbl (%rax, %r11), %edx
  3146: 32da                     xorb %dl, %bl
  3148: 43881c1a                 movb %bl, (%r10, %r11)
  3152: 33db                     xorl %ebx, %ebx
  3154: 49ffc3                   incq %r11
  3157: 4c3bd9                   cmpq %rcx, %r11
  3160: 0f82adffffff             jb 0xc0b
  3166: 4d8b5108                 movq 8(%r9), %r10
  3170: 4c8b1c24                 movq (%rsp), %r11
  3174: 33db                     xorl %ebx, %ebx
  3176: 498b00                   movq (%r8), %rax
  3179: 420fb61418               movzbl (%rax, %r11), %edx
  3184: 4885d2                   testq %rdx, %rdx
  3187: 0f8422000000             je 0xc9b
  3193: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3197: 81c0e6000000             addl $0xe6, %eax
  3203: 81f8ff000000             cmpl $0xff, %eax
  3209: 0f8206000000             jb 0xc95
  3215: 81c001ffffff             addl $0xffffff01, %eax
  3221: 0fb60406                 movzbl (%rsi, %rax), %eax
  3225: 32d8                     xorb %al, %bl
  3227: 498b4008                 movq 8(%r8), %rax
  3231: 420fb61418               movzbl (%rax, %r11), %edx
  3236: 4885d2                   testq %rdx, %rdx
  3239: 0f8422000000             je 0xccf
  3245: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3249: 81c0cd000000             addl $0xcd, %eax
  3255: 81f8ff000000             cmpl $0xff, %eax
  3261: 0f8206000000             jb 0xcc9
  3267: 81c001ffffff             addl $0xffffff01, %eax
  3273: 0fb60406                 movzbl (%rsi, %rax), %eax
  3277: 32d8                     xorb %al, %bl
  3279: 498b4010                 movq 0x10(%r8), %rax
  3283: 420fb61418               movzbl (%rax, %r11), %edx
  3288: 4885d2                   testq %rdx, %rdx
  3291: 0f841f000000             je 0xd00
  3297: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3301: 83c04b                   addl $0x4b, %eax
  3304: 81f8ff000000             cmpl $0xff, %eax
  3310: 0f8206000000             jb 0xcfa
  3316: 81c001ffffff             addl $0xffffff01, %eax
  3322: 0fb60406                 movzbl (%rsi, %rax), %eax
  3326: 32d8                     xorb %al, %bl
  3328: 498b4018                 movq 0x18(%r8), %rax
  3332: 420fb61418               movzbl (%rax, %r11), %edx
  3337: 4885d2                   testq %rdx, %rdx
  3340: 0f841f000000             je 0xd31
  3346: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3350: 83c032                   addl $0x32, %eax
  3353: 81f8ff000000             cmpl $0xff, %eax
  3359: 0f8206000000             jb 0xd2b
  3365: 81c001ffffff             addl $0xffffff01, %eax
  3371: 0fb60406                 movzbl (%rsi, %rax), %eax
  3375: 32d8                     xorb %al, %bl
  3377: 498b4020                 movq 0x20(%r8), %rax
  3381: 420fb61418               movzbl (%rax, %r11), %edx
  3386: 4885d2                   testq %rdx, %rdx
  3389: 0f841f000000             je 0xd62
  3395: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3399: 83c019                   addl $0x19, %eax
  3402: 81f8ff000000             cmpl $0xff, %eax
  3408: 0f8206000000             jb 0xd5c
  3414: 81c001ffffff             addl $0xffffff01, %eax
  3420: 0fb60406                 movzbl (%rsi, %rax), %eax
  3424: 32d8                     xorb %al, %bl
  3426: 498b4028                 movq 0x28(%r8), %rax
  3430: 420fb61418               movzbl (%rax, %r11), %edx
  3435: 32da                     xorb %dl, %bl
  3437: 43881c1a                 movb %bl, (%r10, %r11)
  3441: 33db                     xorl %ebx, %ebx
  3443: 49ffc3                   incq %r11
  3446: 4c3bd9                   cmpq %rcx, %r11
  3449: 0f82e9feffff             jb 0xc68
  3455: 4d8b5110                 movq 0x10(%r9), %r10
  3459: 4c8b1c24                 movq (%rsp), %r11
  3463: 33db                     xorl %ebx, %ebx
  3465: 498b00                   movq (%r8), %rax
  3468: 420fb61418               movzbl (%rax, %r11), %edx
  3473: 4885d2                   testq %rdx, %rdx
  3476: 0f8422000000             je 0xdbc
  3482: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3486: 81c0cd000000             addl $0xcd, %eax
  3492: 81f8ff000000             cmpl $0xff, %eax
  3498: 0f8206000000             jb 0xdb6
  3504: 81c001ffffff             addl $0xffffff01, %eax
  3510: 0fb60406                 movzbl (%rsi, %rax), %eax
  3514: 32d8                     xorb %al, %bl
  3516: 498b4008                 movq 8(%r8), %rax
  3520: 420fb61418               movzbl (%rax, %r11), %edx
  3525: 4885d2                   testq %rdx, %rdx
  3528: 0f841f000000             je 0xded
  3534: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3538: 83c04b                   addl $0x4b, %eax
  3541: 81f8ff000000             cmpl $0xff, %eax
  3547: 0f8206000000             jb 0xde7
  3553: 81c001ffffff             addl $0xffffff01, %eax
  3559: 0fb60406                 movzbl (%rsi, %rax), %eax
  3563: 32d8                     xorb %al, %bl
  3565: 498b4010                 movq 0x10(%r8), %rax
  3569: 420fb61418               movzbl (%rax, %r11), %edx
  3574: 4885d2                   testq %rdx, %rdx
  3577: 0f841f000000             je 0xe1e
  3583: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3587: 83c019                   addl $0x19, %eax
  3590: 81f8ff000000             cmpl $0xff, %eax
  3596: 0f8206000000             jb 0xe18
  3602: 81c001ffffff             addl $0xffffff01, %eax
  3608: 0fb60406                 movzbl (%rsi, %rax), %eax
  3612: 32d8                     xorb %al, %bl
  3614: 498b4018                 movq 0x18(%r8), %rax
  3618: 420fb61418               movzbl (%rax, %r11), %edx
  3623: 4885d2                   testq %rdx, %rdx
  3626: 0f8422000000             je 0xe52
  3632: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3636: 81c0e6000000             addl $0xe6, %eax
  3642: 81f8ff000000             cmpl $0xff, %eax
  3648: 0f8206000000             jb 0xe4c
  3654: 81c001ffffff             addl $0xffffff01, %eax
  3660: 0fb60406                 movzbl (%rsi, %rax), %eax
  3664: 32d8                     xorb %al, %bl
  3666: 498b4020                 movq 0x20(%r8), %rax
  3670: 420fb61418               movzbl (%rax, %r11), %edx
  3675: 4885d2                   testq %rdx, %rdx
  3678: 0f841f000000             je 0xe83
  3684: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3688: 83c032                   addl $0x32, %eax
  3691: 81f8ff000000             cmpl $0xff, %eax
  3697: 0f8206000000             jb 0xe7d
  3703: 81c001ffffff             addl $0xffffff01, %eax
  3709: 0fb60406                 movzbl (%rsi, %rax), %eax
  3713: 32d8                     xorb %al, %bl
  3715: 498b4028                 movq 0x28(%r8), %rax
  3719: 420fb61418               movzbl (%rax, %r11), %edx
  3724: 32da                     xorb %dl, %bl
  3726: 43881c1a                 movb %bl, (%r10, %r11)
  3730: 33db                     xorl %ebx, %ebx
  3732: 49ffc3                   incq %r11
  3735: 4c3bd9                   cmpq %rcx, %r11
  3738: 0f82e9feffff             jb 0xd89
  3744: 4883c408                 addq $8, %rsp
  3748: 415e                     popq %r14
  3750: 415d                     popq %r13
  3752: 415c                     popq %r12
  3754: 415f                     popq %r15
  3756: 5b                       popq %rbx
  3757: c3                       retq 
  3758: 0f0b                     ud2 
  3760: 0f0b                     ud2 
  3762: 0f0b                     ud2 
  3764: 0f0b                     ud2 
  3766: 0f0b                     ud2 
  3768: 0f0b                     ud2 
  3770: 0f0b                     ud2 
  3772: 0f0b                     ud2 
  3774: 0f0b                     ud2 
