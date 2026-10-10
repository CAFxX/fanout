; config_t3_2101 target=avx512-gfni twin=baseline topo=lrc(8,2,2)/enc op=encode U=4
; code_size=2944 twins=[] uops_per_byte=1.0390625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d670b0000           leaq 0xb67(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e800010000           subq $0x100, %rax
    57: 0f822c070000             jb 0x76b
    63: 90                       nop 
    64: 62217d486f041e           vmovdqa32 (%rsi, %r11), %zmm24
    71: 62917d486fc0             vmovdqa32 %zmm24, %zmm0
    77: 62917d486fe0             vmovdqa32 %zmm24, %zmm4
    83: 62117d486fc0             vmovdqa32 %zmm24, %zmm8
    89: 62217d486f441e01         vmovdqa32 0x40(%rsi, %r11), %zmm24
    97: 62917d486fc8             vmovdqa32 %zmm24, %zmm1
   103: 62917d486fe8             vmovdqa32 %zmm24, %zmm5
   109: 62117d486fc8             vmovdqa32 %zmm24, %zmm9
   115: 62217d486f441e02         vmovdqa32 0x80(%rsi, %r11), %zmm24
   123: 62917d486fd0             vmovdqa32 %zmm24, %zmm2
   129: 62917d486ff0             vmovdqa32 %zmm24, %zmm6
   135: 62117d486fd0             vmovdqa32 %zmm24, %zmm10
   141: 62217d486f441e03         vmovdqa32 0xc0(%rsi, %r11), %zmm24
   149: 62917d486fd8             vmovdqa32 %zmm24, %zmm3
   155: 62917d486ff8             vmovdqa32 %zmm24, %zmm7
   161: 62117d486fd8             vmovdqa32 %zmm24, %zmm11
   167: 62217d486f041f           vmovdqa32 (%rdi, %r11), %zmm24
   174: 62417d486f17             vmovdqa32 (%r15), %zmm26
   180: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   187: 62917d48efc1             vpxord %zmm25, %zmm0, %zmm0
   193: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   200: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   207: 62915d48efe1             vpxord %zmm25, %zmm4, %zmm4
   213: 62113d48efc0             vpxord %zmm24, %zmm8, %zmm8
   219: 62217d486f441f01         vmovdqa32 0x40(%rdi, %r11), %zmm24
   227: 62417d486f17             vmovdqa32 (%r15), %zmm26
   233: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   240: 62917548efc9             vpxord %zmm25, %zmm1, %zmm1
   246: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   253: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   260: 62915548efe9             vpxord %zmm25, %zmm5, %zmm5
   266: 62113548efc8             vpxord %zmm24, %zmm9, %zmm9
   272: 62217d486f441f02         vmovdqa32 0x80(%rdi, %r11), %zmm24
   280: 62417d486f17             vmovdqa32 (%r15), %zmm26
   286: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   293: 62916d48efd1             vpxord %zmm25, %zmm2, %zmm2
   299: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   306: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   313: 62914d48eff1             vpxord %zmm25, %zmm6, %zmm6
   319: 62112d48efd0             vpxord %zmm24, %zmm10, %zmm10
   325: 62217d486f441f03         vmovdqa32 0xc0(%rdi, %r11), %zmm24
   333: 62417d486f17             vmovdqa32 (%r15), %zmm26
   339: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   346: 62916548efd9             vpxord %zmm25, %zmm3, %zmm3
   352: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   359: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   366: 62914548eff9             vpxord %zmm25, %zmm7, %zmm7
   372: 62112548efd8             vpxord %zmm24, %zmm11, %zmm11
   378: 62017d486f041c           vmovdqa32 (%r12, %r11), %zmm24
   385: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   392: 62917d48efc1             vpxord %zmm25, %zmm0, %zmm0
   398: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   405: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   412: 62915d48efe1             vpxord %zmm25, %zmm4, %zmm4
   418: 62117d486fe0             vmovdqa32 %zmm24, %zmm12
   424: 62017d486f441c01         vmovdqa32 0x40(%r12, %r11), %zmm24
   432: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   439: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   446: 62917548efc9             vpxord %zmm25, %zmm1, %zmm1
   452: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   459: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   466: 62915548efe9             vpxord %zmm25, %zmm5, %zmm5
   472: 62117d486fe8             vmovdqa32 %zmm24, %zmm13
   478: 62017d486f441c02         vmovdqa32 0x80(%r12, %r11), %zmm24
   486: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   493: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   500: 62916d48efd1             vpxord %zmm25, %zmm2, %zmm2
   506: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   513: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   520: 62914d48eff1             vpxord %zmm25, %zmm6, %zmm6
   526: 62117d486ff0             vmovdqa32 %zmm24, %zmm14
   532: 62017d486f441c03         vmovdqa32 0xc0(%r12, %r11), %zmm24
   540: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   547: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   554: 62916548efd9             vpxord %zmm25, %zmm3, %zmm3
   560: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   567: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   574: 62914548eff9             vpxord %zmm25, %zmm7, %zmm7
   580: 62117d486ff8             vmovdqa32 %zmm24, %zmm15
   586: 62017d486f441d00         vmovdqa32 (%r13, %r11), %zmm24
   594: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
   601: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   608: 62917d48efc1             vpxord %zmm25, %zmm0, %zmm0
   614: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
   621: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   628: 62915d48efe1             vpxord %zmm25, %zmm4, %zmm4
   634: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
   640: 62017d486f441d01         vmovdqa32 0x40(%r13, %r11), %zmm24
   648: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
   655: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   662: 62917548efc9             vpxord %zmm25, %zmm1, %zmm1
   668: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
   675: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   682: 62915548efe9             vpxord %zmm25, %zmm5, %zmm5
   688: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
   694: 62017d486f441d02         vmovdqa32 0x80(%r13, %r11), %zmm24
   702: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
   709: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   716: 62916d48efd1             vpxord %zmm25, %zmm2, %zmm2
   722: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
   729: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   736: 62914d48eff1             vpxord %zmm25, %zmm6, %zmm6
   742: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
   748: 62017d486f441d03         vmovdqa32 0xc0(%r13, %r11), %zmm24
   756: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
   763: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   770: 62916548efd9             vpxord %zmm25, %zmm3, %zmm3
   776: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
   783: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   790: 62914548eff9             vpxord %zmm25, %zmm7, %zmm7
   796: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
   802: 62017d486f041e           vmovdqa32 (%r14, %r11), %zmm24
   809: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   816: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   823: 62917d48efc1             vpxord %zmm25, %zmm0, %zmm0
   829: 62417d486f5705           vmovdqa32 0x140(%r15), %zmm26
   836: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   843: 62915d48efe1             vpxord %zmm25, %zmm4, %zmm4
   849: 62817d486fc0             vmovdqa32 %zmm24, %zmm16
   855: 62017d486f441e01         vmovdqa32 0x40(%r14, %r11), %zmm24
   863: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   870: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   877: 62917548efc9             vpxord %zmm25, %zmm1, %zmm1
   883: 62417d486f5705           vmovdqa32 0x140(%r15), %zmm26
   890: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   897: 62915548efe9             vpxord %zmm25, %zmm5, %zmm5
   903: 62817d486fc8             vmovdqa32 %zmm24, %zmm17
   909: 62017d486f441e02         vmovdqa32 0x80(%r14, %r11), %zmm24
   917: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   924: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   931: 62916d48efd1             vpxord %zmm25, %zmm2, %zmm2
   937: 62417d486f5705           vmovdqa32 0x140(%r15), %zmm26
   944: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   951: 62914d48eff1             vpxord %zmm25, %zmm6, %zmm6
   957: 62817d486fd0             vmovdqa32 %zmm24, %zmm18
   963: 62017d486f441e03         vmovdqa32 0xc0(%r14, %r11), %zmm24
   971: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   978: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   985: 62916548efd9             vpxord %zmm25, %zmm3, %zmm3
   991: 62417d486f5705           vmovdqa32 0x140(%r15), %zmm26
   998: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1005: 62914548eff9             vpxord %zmm25, %zmm7, %zmm7
  1011: 62817d486fd8             vmovdqa32 %zmm24, %zmm19
  1017: 4d8b5028                 movq 0x28(%r8), %r10
  1021: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1028: 62417d486f5706           vmovdqa32 0x180(%r15), %zmm26
  1035: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1042: 62917d48efc1             vpxord %zmm25, %zmm0, %zmm0
  1048: 62417d486f5707           vmovdqa32 0x1c0(%r15), %zmm26
  1055: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1062: 62915d48efe1             vpxord %zmm25, %zmm4, %zmm4
  1068: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
  1074: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1082: 62417d486f5706           vmovdqa32 0x180(%r15), %zmm26
  1089: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1096: 62917548efc9             vpxord %zmm25, %zmm1, %zmm1
  1102: 62417d486f5707           vmovdqa32 0x1c0(%r15), %zmm26
  1109: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1116: 62915548efe9             vpxord %zmm25, %zmm5, %zmm5
  1122: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
  1128: 62017d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm24
  1136: 62417d486f5706           vmovdqa32 0x180(%r15), %zmm26
  1143: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1150: 62916d48efd1             vpxord %zmm25, %zmm2, %zmm2
  1156: 62417d486f5707           vmovdqa32 0x1c0(%r15), %zmm26
  1163: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1170: 62914d48eff1             vpxord %zmm25, %zmm6, %zmm6
  1176: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
  1182: 62017d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm24
  1190: 62417d486f5706           vmovdqa32 0x180(%r15), %zmm26
  1197: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1204: 62916548efd9             vpxord %zmm25, %zmm3, %zmm3
  1210: 62417d486f5707           vmovdqa32 0x1c0(%r15), %zmm26
  1217: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1224: 62914548eff9             vpxord %zmm25, %zmm7, %zmm7
  1230: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
  1236: 4d8b5030                 movq 0x30(%r8), %r10
  1240: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1247: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1254: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1261: 62917d48efc1             vpxord %zmm25, %zmm0, %zmm0
  1267: 62417d486f5708           vmovdqa32 0x200(%r15), %zmm26
  1274: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1281: 62915d48efe1             vpxord %zmm25, %zmm4, %zmm4
  1287: 62817d486fe0             vmovdqa32 %zmm24, %zmm20
  1293: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1301: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1308: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1315: 62917548efc9             vpxord %zmm25, %zmm1, %zmm1
  1321: 62417d486f5708           vmovdqa32 0x200(%r15), %zmm26
  1328: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1335: 62915548efe9             vpxord %zmm25, %zmm5, %zmm5
  1341: 62817d486fe8             vmovdqa32 %zmm24, %zmm21
  1347: 62017d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm24
  1355: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1362: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1369: 62916d48efd1             vpxord %zmm25, %zmm2, %zmm2
  1375: 62417d486f5708           vmovdqa32 0x200(%r15), %zmm26
  1382: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1389: 62914d48eff1             vpxord %zmm25, %zmm6, %zmm6
  1395: 62817d486ff0             vmovdqa32 %zmm24, %zmm22
  1401: 62017d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm24
  1409: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1416: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1423: 62916548efd9             vpxord %zmm25, %zmm3, %zmm3
  1429: 62417d486f5708           vmovdqa32 0x200(%r15), %zmm26
  1436: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1443: 62914548eff9             vpxord %zmm25, %zmm7, %zmm7
  1449: 62817d486ff8             vmovdqa32 %zmm24, %zmm23
  1455: 4d8b5038                 movq 0x38(%r8), %r10
  1459: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  1466: 62417d486f5709           vmovdqa32 0x240(%r15), %zmm26
  1473: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1480: 62917d48efc1             vpxord %zmm25, %zmm0, %zmm0
  1486: 62417d486f570a           vmovdqa32 0x280(%r15), %zmm26
  1493: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1500: 62915d48efe1             vpxord %zmm25, %zmm4, %zmm4
  1506: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
  1512: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  1520: 62417d486f5709           vmovdqa32 0x240(%r15), %zmm26
  1527: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1534: 62917548efc9             vpxord %zmm25, %zmm1, %zmm1
  1540: 62417d486f570a           vmovdqa32 0x280(%r15), %zmm26
  1547: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1554: 62915548efe9             vpxord %zmm25, %zmm5, %zmm5
  1560: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
  1566: 62017d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm24
  1574: 62417d486f5709           vmovdqa32 0x240(%r15), %zmm26
  1581: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1588: 62916d48efd1             vpxord %zmm25, %zmm2, %zmm2
  1594: 62417d486f570a           vmovdqa32 0x280(%r15), %zmm26
  1601: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1608: 62914d48eff1             vpxord %zmm25, %zmm6, %zmm6
  1614: 62814d40eff0             vpxord %zmm24, %zmm22, %zmm22
  1620: 62017d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm24
  1628: 62417d486f5709           vmovdqa32 0x240(%r15), %zmm26
  1635: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1642: 62916548efd9             vpxord %zmm25, %zmm3, %zmm3
  1648: 62417d486f570a           vmovdqa32 0x280(%r15), %zmm26
  1655: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1662: 62914548eff9             vpxord %zmm25, %zmm7, %zmm7
  1668: 62814540eff8             vpxord %zmm24, %zmm23, %zmm23
  1674: 4d8b11                   movq (%r9), %r10
  1677: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  1684: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  1692: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
  1700: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
  1708: 4d8b5108                 movq 8(%r9), %r10
  1712: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
  1719: 62917d487f6c1a01         vmovdqa32 %zmm5, 0x40(%r10, %r11)
  1727: 62917d487f741a02         vmovdqa32 %zmm6, 0x80(%r10, %r11)
  1735: 62917d487f7c1a03         vmovdqa32 %zmm7, 0xc0(%r10, %r11)
  1743: 4d8b5110                 movq 0x10(%r9), %r10
  1747: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  1754: 62117d487f4c1a01         vmovdqa32 %zmm9, 0x40(%r10, %r11)
  1762: 62117d487f541a02         vmovdqa32 %zmm10, 0x80(%r10, %r11)
  1770: 62117d487f5c1a03         vmovdqa32 %zmm11, 0xc0(%r10, %r11)
  1778: 4d8b5118                 movq 0x18(%r9), %r10
  1782: 62117d487f241a           vmovdqa32 %zmm12, (%r10, %r11)
  1789: 62117d487f6c1a01         vmovdqa32 %zmm13, 0x40(%r10, %r11)
  1797: 62117d487f741a02         vmovdqa32 %zmm14, 0x80(%r10, %r11)
  1805: 62117d487f7c1a03         vmovdqa32 %zmm15, 0xc0(%r10, %r11)
  1813: 4d8b5120                 movq 0x20(%r9), %r10
  1817: 62817d487f041a           vmovdqa32 %zmm16, (%r10, %r11)
  1824: 62817d487f4c1a01         vmovdqa32 %zmm17, 0x40(%r10, %r11)
  1832: 62817d487f541a02         vmovdqa32 %zmm18, 0x80(%r10, %r11)
  1840: 62817d487f5c1a03         vmovdqa32 %zmm19, 0xc0(%r10, %r11)
  1848: 4d8b5128                 movq 0x28(%r9), %r10
  1852: 62817d487f241a           vmovdqa32 %zmm20, (%r10, %r11)
  1859: 62817d487f6c1a01         vmovdqa32 %zmm21, 0x40(%r10, %r11)
  1867: 62817d487f741a02         vmovdqa32 %zmm22, 0x80(%r10, %r11)
  1875: 62817d487f7c1a03         vmovdqa32 %zmm23, 0xc0(%r10, %r11)
  1883: 4981c300010000           addq $0x100, %r11
  1890: 4c3bd8                   cmpq %rax, %r11
  1893: 0f86d5f8ffff             jbe 0x40
  1899: 4c3bd9                   cmpq %rcx, %r11
  1902: 0f83e3030000             jae 0xb57
  1908: 4153                     pushq %r11
  1910: 498dbfc0020000           leaq 0x2c0(%r15), %rdi
  1917: 498db7c0030000           leaq 0x3c0(%r15), %rsi
  1924: 4d8b11                   movq (%r9), %r10
  1927: 4c8b1c24                 movq (%rsp), %r11
  1931: 33db                     xorl %ebx, %ebx
  1933: 498b00                   movq (%r8), %rax
  1936: 420fb61418               movzbl (%rax, %r11), %edx
  1941: 32da                     xorb %dl, %bl
  1943: 498b4008                 movq 8(%r8), %rax
  1947: 420fb61418               movzbl (%rax, %r11), %edx
  1952: 4885d2                   testq %rdx, %rdx
  1955: 0f841f000000             je 0x7c8
  1961: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1965: 83c019                   addl $0x19, %eax
  1968: 81f8ff000000             cmpl $0xff, %eax
  1974: 0f8206000000             jb 0x7c2
  1980: 81c001ffffff             addl $0xffffff01, %eax
  1986: 0fb60406                 movzbl (%rsi, %rax), %eax
  1990: 32d8                     xorb %al, %bl
  1992: 498b4010                 movq 0x10(%r8), %rax
  1996: 420fb61418               movzbl (%rax, %r11), %edx
  2001: 4885d2                   testq %rdx, %rdx
  2004: 0f841f000000             je 0x7f9
  2010: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2014: 83c032                   addl $0x32, %eax
  2017: 81f8ff000000             cmpl $0xff, %eax
  2023: 0f8206000000             jb 0x7f3
  2029: 81c001ffffff             addl $0xffffff01, %eax
  2035: 0fb60406                 movzbl (%rsi, %rax), %eax
  2039: 32d8                     xorb %al, %bl
  2041: 498b4018                 movq 0x18(%r8), %rax
  2045: 420fb61418               movzbl (%rax, %r11), %edx
  2050: 4885d2                   testq %rdx, %rdx
  2053: 0f841f000000             je 0x82a
  2059: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2063: 83c04b                   addl $0x4b, %eax
  2066: 81f8ff000000             cmpl $0xff, %eax
  2072: 0f8206000000             jb 0x824
  2078: 81c001ffffff             addl $0xffffff01, %eax
  2084: 0fb60406                 movzbl (%rsi, %rax), %eax
  2088: 32d8                     xorb %al, %bl
  2090: 498b4020                 movq 0x20(%r8), %rax
  2094: 420fb61418               movzbl (%rax, %r11), %edx
  2099: 4885d2                   testq %rdx, %rdx
  2102: 0f841f000000             je 0x85b
  2108: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2112: 83c064                   addl $0x64, %eax
  2115: 81f8ff000000             cmpl $0xff, %eax
  2121: 0f8206000000             jb 0x855
  2127: 81c001ffffff             addl $0xffffff01, %eax
  2133: 0fb60406                 movzbl (%rsi, %rax), %eax
  2137: 32d8                     xorb %al, %bl
  2139: 498b4028                 movq 0x28(%r8), %rax
  2143: 420fb61418               movzbl (%rax, %r11), %edx
  2148: 4885d2                   testq %rdx, %rdx
  2151: 0f841f000000             je 0x88c
  2157: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2161: 83c07d                   addl $0x7d, %eax
  2164: 81f8ff000000             cmpl $0xff, %eax
  2170: 0f8206000000             jb 0x886
  2176: 81c001ffffff             addl $0xffffff01, %eax
  2182: 0fb60406                 movzbl (%rsi, %rax), %eax
  2186: 32d8                     xorb %al, %bl
  2188: 498b4030                 movq 0x30(%r8), %rax
  2192: 420fb61418               movzbl (%rax, %r11), %edx
  2197: 4885d2                   testq %rdx, %rdx
  2200: 0f8422000000             je 0x8c0
  2206: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2210: 81c096000000             addl $0x96, %eax
  2216: 81f8ff000000             cmpl $0xff, %eax
  2222: 0f8206000000             jb 0x8ba
  2228: 81c001ffffff             addl $0xffffff01, %eax
  2234: 0fb60406                 movzbl (%rsi, %rax), %eax
  2238: 32d8                     xorb %al, %bl
  2240: 498b4038                 movq 0x38(%r8), %rax
  2244: 420fb61418               movzbl (%rax, %r11), %edx
  2249: 4885d2                   testq %rdx, %rdx
  2252: 0f8422000000             je 0x8f4
  2258: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2262: 81c0af000000             addl $0xaf, %eax
  2268: 81f8ff000000             cmpl $0xff, %eax
  2274: 0f8206000000             jb 0x8ee
  2280: 81c001ffffff             addl $0xffffff01, %eax
  2286: 0fb60406                 movzbl (%rsi, %rax), %eax
  2290: 32d8                     xorb %al, %bl
  2292: 43881c1a                 movb %bl, (%r10, %r11)
  2296: 33db                     xorl %ebx, %ebx
  2298: 49ffc3                   incq %r11
  2301: 4c3bd9                   cmpq %rcx, %r11
  2304: 0f8287feffff             jb 0x78d
  2310: 4d8b5108                 movq 8(%r9), %r10
  2314: 4c8b1c24                 movq (%rsp), %r11
  2318: 33db                     xorl %ebx, %ebx
  2320: 498b00                   movq (%r8), %rax
  2323: 420fb61418               movzbl (%rax, %r11), %edx
  2328: 32da                     xorb %dl, %bl
  2330: 498b4008                 movq 8(%r8), %rax
  2334: 420fb61418               movzbl (%rax, %r11), %edx
  2339: 4885d2                   testq %rdx, %rdx
  2342: 0f841f000000             je 0x94b
  2348: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2352: 83c032                   addl $0x32, %eax
  2355: 81f8ff000000             cmpl $0xff, %eax
  2361: 0f8206000000             jb 0x945
  2367: 81c001ffffff             addl $0xffffff01, %eax
  2373: 0fb60406                 movzbl (%rsi, %rax), %eax
  2377: 32d8                     xorb %al, %bl
  2379: 498b4010                 movq 0x10(%r8), %rax
  2383: 420fb61418               movzbl (%rax, %r11), %edx
  2388: 4885d2                   testq %rdx, %rdx
  2391: 0f841f000000             je 0x97c
  2397: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2401: 83c064                   addl $0x64, %eax
  2404: 81f8ff000000             cmpl $0xff, %eax
  2410: 0f8206000000             jb 0x976
  2416: 81c001ffffff             addl $0xffffff01, %eax
  2422: 0fb60406                 movzbl (%rsi, %rax), %eax
  2426: 32d8                     xorb %al, %bl
  2428: 498b4018                 movq 0x18(%r8), %rax
  2432: 420fb61418               movzbl (%rax, %r11), %edx
  2437: 4885d2                   testq %rdx, %rdx
  2440: 0f8422000000             je 0x9b0
  2446: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2450: 81c096000000             addl $0x96, %eax
  2456: 81f8ff000000             cmpl $0xff, %eax
  2462: 0f8206000000             jb 0x9aa
  2468: 81c001ffffff             addl $0xffffff01, %eax
  2474: 0fb60406                 movzbl (%rsi, %rax), %eax
  2478: 32d8                     xorb %al, %bl
  2480: 498b4020                 movq 0x20(%r8), %rax
  2484: 420fb61418               movzbl (%rax, %r11), %edx
  2489: 4885d2                   testq %rdx, %rdx
  2492: 0f8422000000             je 0x9e4
  2498: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2502: 81c0c8000000             addl $0xc8, %eax
  2508: 81f8ff000000             cmpl $0xff, %eax
  2514: 0f8206000000             jb 0x9de
  2520: 81c001ffffff             addl $0xffffff01, %eax
  2526: 0fb60406                 movzbl (%rsi, %rax), %eax
  2530: 32d8                     xorb %al, %bl
  2532: 498b4028                 movq 0x28(%r8), %rax
  2536: 420fb61418               movzbl (%rax, %r11), %edx
  2541: 4885d2                   testq %rdx, %rdx
  2544: 0f8422000000             je 0xa18
  2550: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2554: 81c0fa000000             addl $0xfa, %eax
  2560: 81f8ff000000             cmpl $0xff, %eax
  2566: 0f8206000000             jb 0xa12
  2572: 81c001ffffff             addl $0xffffff01, %eax
  2578: 0fb60406                 movzbl (%rsi, %rax), %eax
  2582: 32d8                     xorb %al, %bl
  2584: 498b4030                 movq 0x30(%r8), %rax
  2588: 420fb61418               movzbl (%rax, %r11), %edx
  2593: 4885d2                   testq %rdx, %rdx
  2596: 0f841f000000             je 0xa49
  2602: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2606: 83c02d                   addl $0x2d, %eax
  2609: 81f8ff000000             cmpl $0xff, %eax
  2615: 0f8206000000             jb 0xa43
  2621: 81c001ffffff             addl $0xffffff01, %eax
  2627: 0fb60406                 movzbl (%rsi, %rax), %eax
  2631: 32d8                     xorb %al, %bl
  2633: 498b4038                 movq 0x38(%r8), %rax
  2637: 420fb61418               movzbl (%rax, %r11), %edx
  2642: 4885d2                   testq %rdx, %rdx
  2645: 0f841f000000             je 0xa7a
  2651: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2655: 83c05f                   addl $0x5f, %eax
  2658: 81f8ff000000             cmpl $0xff, %eax
  2664: 0f8206000000             jb 0xa74
  2670: 81c001ffffff             addl $0xffffff01, %eax
  2676: 0fb60406                 movzbl (%rsi, %rax), %eax
  2680: 32d8                     xorb %al, %bl
  2682: 43881c1a                 movb %bl, (%r10, %r11)
  2686: 33db                     xorl %ebx, %ebx
  2688: 49ffc3                   incq %r11
  2691: 4c3bd9                   cmpq %rcx, %r11
  2694: 0f8284feffff             jb 0x910
  2700: 4d8b5110                 movq 0x10(%r9), %r10
  2704: 4c8b1c24                 movq (%rsp), %r11
  2708: 33db                     xorl %ebx, %ebx
  2710: 498b00                   movq (%r8), %rax
  2713: 420fb61418               movzbl (%rax, %r11), %edx
  2718: 32da                     xorb %dl, %bl
  2720: 498b4008                 movq 8(%r8), %rax
  2724: 420fb61418               movzbl (%rax, %r11), %edx
  2729: 32da                     xorb %dl, %bl
  2731: 43881c1a                 movb %bl, (%r10, %r11)
  2735: 33db                     xorl %ebx, %ebx
  2737: 49ffc3                   incq %r11
  2740: 4c3bd9                   cmpq %rcx, %r11
  2743: 0f82d9ffffff             jb 0xa96
  2749: 4d8b5118                 movq 0x18(%r9), %r10
  2753: 4c8b1c24                 movq (%rsp), %r11
  2757: 33db                     xorl %ebx, %ebx
  2759: 498b4010                 movq 0x10(%r8), %rax
  2763: 420fb61418               movzbl (%rax, %r11), %edx
  2768: 32da                     xorb %dl, %bl
  2770: 498b4018                 movq 0x18(%r8), %rax
  2774: 420fb61418               movzbl (%rax, %r11), %edx
  2779: 32da                     xorb %dl, %bl
  2781: 43881c1a                 movb %bl, (%r10, %r11)
  2785: 33db                     xorl %ebx, %ebx
  2787: 49ffc3                   incq %r11
  2790: 4c3bd9                   cmpq %rcx, %r11
  2793: 0f82d8ffffff             jb 0xac7
  2799: 4d8b5120                 movq 0x20(%r9), %r10
  2803: 4c8b1c24                 movq (%rsp), %r11
  2807: 33db                     xorl %ebx, %ebx
  2809: 498b4020                 movq 0x20(%r8), %rax
  2813: 420fb61418               movzbl (%rax, %r11), %edx
  2818: 32da                     xorb %dl, %bl
  2820: 498b4028                 movq 0x28(%r8), %rax
  2824: 420fb61418               movzbl (%rax, %r11), %edx
  2829: 32da                     xorb %dl, %bl
  2831: 43881c1a                 movb %bl, (%r10, %r11)
  2835: 33db                     xorl %ebx, %ebx
  2837: 49ffc3                   incq %r11
  2840: 4c3bd9                   cmpq %rcx, %r11
  2843: 0f82d8ffffff             jb 0xaf9
  2849: 4d8b5128                 movq 0x28(%r9), %r10
  2853: 4c8b1c24                 movq (%rsp), %r11
  2857: 33db                     xorl %ebx, %ebx
  2859: 498b4030                 movq 0x30(%r8), %rax
  2863: 420fb61418               movzbl (%rax, %r11), %edx
  2868: 32da                     xorb %dl, %bl
  2870: 498b4038                 movq 0x38(%r8), %rax
  2874: 420fb61418               movzbl (%rax, %r11), %edx
  2879: 32da                     xorb %dl, %bl
  2881: 43881c1a                 movb %bl, (%r10, %r11)
  2885: 33db                     xorl %ebx, %ebx
  2887: 49ffc3                   incq %r11
  2890: 4c3bd9                   cmpq %rcx, %r11
  2893: 0f82d8ffffff             jb 0xb2b
  2899: 4883c408                 addq $8, %rsp
  2903: 415e                     popq %r14
  2905: 415d                     popq %r13
  2907: 415c                     popq %r12
  2909: 415f                     popq %r15
  2911: 5b                       popq %rbx
  2912: c5f877                   vzeroupper 
  2915: c3                       retq 
  2916: 0f0b                     ud2 
  2918: 0f0b                     ud2 
  2920: 0f0b                     ud2 
  2922: 0f0b                     ud2 
  2924: 0f0b                     ud2 
  2926: 0f0b                     ud2 
  2928: 0f0b                     ud2 
  2930: 0f0b                     ud2 
  2932: 0f0b                     ud2 
  2934: 0f0b                     ud2 
  2936: 0f0b                     ud2 
  2938: 0f0b                     ud2 
  2940: 0f0b                     ud2 
  2942: 0f0b                     ud2 
