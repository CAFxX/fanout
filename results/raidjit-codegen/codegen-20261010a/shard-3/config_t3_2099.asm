; config_t3_2099 target=avx512-gfni twin=baseline topo=lrc_azure()/enc op=encode U=4
; code_size=4160 twins=[] uops_per_byte=1.5234375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d27100000           leaq 0x1027(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e800010000           subq $0x100, %rax
    57: 0f82520a0000             jb 0xa91
    63: 90                       nop 
    64: 62a17d486f041e           vmovdqa32 (%rsi, %r11), %zmm16
    71: 62b17d486fc0             vmovdqa32 %zmm16, %zmm0
    77: 62b17d486fe0             vmovdqa32 %zmm16, %zmm4
    83: 62317d486fc0             vmovdqa32 %zmm16, %zmm8
    89: 62a17d486f441e01         vmovdqa32 0x40(%rsi, %r11), %zmm16
    97: 62b17d486fc8             vmovdqa32 %zmm16, %zmm1
   103: 62b17d486fe8             vmovdqa32 %zmm16, %zmm5
   109: 62317d486fc8             vmovdqa32 %zmm16, %zmm9
   115: 62a17d486f441e02         vmovdqa32 0x80(%rsi, %r11), %zmm16
   123: 62b17d486fd0             vmovdqa32 %zmm16, %zmm2
   129: 62b17d486ff0             vmovdqa32 %zmm16, %zmm6
   135: 62317d486fd0             vmovdqa32 %zmm16, %zmm10
   141: 62a17d486f441e03         vmovdqa32 0xc0(%rsi, %r11), %zmm16
   149: 62b17d486fd8             vmovdqa32 %zmm16, %zmm3
   155: 62b17d486ff8             vmovdqa32 %zmm16, %zmm7
   161: 62317d486fd8             vmovdqa32 %zmm16, %zmm11
   167: 62a17d486f041f           vmovdqa32 (%rdi, %r11), %zmm16
   174: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   180: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   187: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
   193: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   200: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   207: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
   213: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
   219: 62a17d486f441f01         vmovdqa32 0x40(%rdi, %r11), %zmm16
   227: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   233: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   240: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
   246: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   253: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   260: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
   266: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
   272: 62a17d486f441f02         vmovdqa32 0x80(%rdi, %r11), %zmm16
   280: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   286: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   293: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
   299: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   306: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   313: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
   319: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
   325: 62a17d486f441f03         vmovdqa32 0xc0(%rdi, %r11), %zmm16
   333: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   339: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   346: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
   352: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   359: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   366: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
   372: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
   378: 62817d486f041c           vmovdqa32 (%r12, %r11), %zmm16
   385: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   392: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
   398: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   405: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   412: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
   418: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
   424: 62817d486f441c01         vmovdqa32 0x40(%r12, %r11), %zmm16
   432: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   439: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   446: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
   452: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   459: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   466: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
   472: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
   478: 62817d486f441c02         vmovdqa32 0x80(%r12, %r11), %zmm16
   486: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   493: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   500: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
   506: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   513: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   520: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
   526: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
   532: 62817d486f441c03         vmovdqa32 0xc0(%r12, %r11), %zmm16
   540: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   547: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   554: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
   560: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   567: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   574: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
   580: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
   586: 62817d486f441d00         vmovdqa32 (%r13, %r11), %zmm16
   594: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   601: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   608: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
   614: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
   621: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   628: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
   634: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
   640: 62817d486f441d01         vmovdqa32 0x40(%r13, %r11), %zmm16
   648: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   655: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   662: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
   668: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
   675: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   682: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
   688: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
   694: 62817d486f441d02         vmovdqa32 0x80(%r13, %r11), %zmm16
   702: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   709: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   716: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
   722: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
   729: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   736: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
   742: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
   748: 62817d486f441d03         vmovdqa32 0xc0(%r13, %r11), %zmm16
   756: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   763: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   770: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
   776: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
   783: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   790: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
   796: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
   802: 62817d486f041e           vmovdqa32 (%r14, %r11), %zmm16
   809: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   816: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   823: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
   829: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
   836: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   843: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
   849: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
   855: 62817d486f441e01         vmovdqa32 0x40(%r14, %r11), %zmm16
   863: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   870: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   877: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
   883: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
   890: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   897: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
   903: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
   909: 62817d486f441e02         vmovdqa32 0x80(%r14, %r11), %zmm16
   917: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   924: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   931: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
   937: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
   944: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   951: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
   957: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
   963: 62817d486f441e03         vmovdqa32 0xc0(%r14, %r11), %zmm16
   971: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   978: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   985: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
   991: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
   998: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1005: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  1011: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
  1017: 4d8b5028                 movq 0x28(%r8), %r10
  1021: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  1028: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1035: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1042: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
  1048: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1055: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1062: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  1068: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
  1074: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  1082: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1089: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1096: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
  1102: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1109: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1116: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  1122: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
  1128: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  1136: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1143: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1150: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
  1156: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1163: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1170: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  1176: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
  1182: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  1190: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1197: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1204: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
  1210: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1217: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1224: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  1230: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
  1236: 4d8b5030                 movq 0x30(%r8), %r10
  1240: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  1247: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1254: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1261: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
  1267: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1274: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1281: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  1287: 62317d486fe0             vmovdqa32 %zmm16, %zmm12
  1293: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  1301: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1308: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1315: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
  1321: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1328: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1335: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  1341: 62317d486fe8             vmovdqa32 %zmm16, %zmm13
  1347: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  1355: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1362: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1369: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
  1375: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1382: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1389: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  1395: 62317d486ff0             vmovdqa32 %zmm16, %zmm14
  1401: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  1409: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1416: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1423: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
  1429: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1436: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1443: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  1449: 62317d486ff8             vmovdqa32 %zmm16, %zmm15
  1455: 4d8b5038                 movq 0x38(%r8), %r10
  1459: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  1466: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1473: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1480: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
  1486: 62c17d486f570a           vmovdqa32 0x280(%r15), %zmm18
  1493: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1500: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  1506: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
  1512: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  1520: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1527: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1534: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
  1540: 62c17d486f570a           vmovdqa32 0x280(%r15), %zmm18
  1547: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1554: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  1560: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
  1566: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  1574: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1581: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1588: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
  1594: 62c17d486f570a           vmovdqa32 0x280(%r15), %zmm18
  1601: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1608: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  1614: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
  1620: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  1628: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1635: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1642: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
  1648: 62c17d486f570a           vmovdqa32 0x280(%r15), %zmm18
  1655: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1662: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  1668: 62310548eff8             vpxord %zmm16, %zmm15, %zmm15
  1674: 4d8b5040                 movq 0x40(%r8), %r10
  1678: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  1685: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1692: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1699: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
  1705: 62c17d486f570b           vmovdqa32 0x2c0(%r15), %zmm18
  1712: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1719: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  1725: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
  1731: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  1739: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1746: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1753: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
  1759: 62c17d486f570b           vmovdqa32 0x2c0(%r15), %zmm18
  1766: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1773: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  1779: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
  1785: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  1793: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1800: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1807: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
  1813: 62c17d486f570b           vmovdqa32 0x2c0(%r15), %zmm18
  1820: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1827: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  1833: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
  1839: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  1847: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1854: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1861: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
  1867: 62c17d486f570b           vmovdqa32 0x2c0(%r15), %zmm18
  1874: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1881: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  1887: 62310548eff8             vpxord %zmm16, %zmm15, %zmm15
  1893: 4d8b5048                 movq 0x48(%r8), %r10
  1897: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  1904: 62c17d486f570c           vmovdqa32 0x300(%r15), %zmm18
  1911: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1918: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
  1924: 62c17d486f570d           vmovdqa32 0x340(%r15), %zmm18
  1931: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1938: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  1944: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
  1950: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  1958: 62c17d486f570c           vmovdqa32 0x300(%r15), %zmm18
  1965: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1972: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
  1978: 62c17d486f570d           vmovdqa32 0x340(%r15), %zmm18
  1985: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1992: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  1998: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
  2004: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  2012: 62c17d486f570c           vmovdqa32 0x300(%r15), %zmm18
  2019: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2026: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
  2032: 62c17d486f570d           vmovdqa32 0x340(%r15), %zmm18
  2039: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2046: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  2052: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
  2058: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  2066: 62c17d486f570c           vmovdqa32 0x300(%r15), %zmm18
  2073: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2080: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
  2086: 62c17d486f570d           vmovdqa32 0x340(%r15), %zmm18
  2093: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2100: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  2106: 62310548eff8             vpxord %zmm16, %zmm15, %zmm15
  2112: 4d8b5050                 movq 0x50(%r8), %r10
  2116: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  2123: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  2130: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2137: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
  2143: 62c17d486f570e           vmovdqa32 0x380(%r15), %zmm18
  2150: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2157: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  2163: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
  2169: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  2177: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  2184: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2191: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
  2197: 62c17d486f570e           vmovdqa32 0x380(%r15), %zmm18
  2204: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2211: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  2217: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
  2223: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  2231: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  2238: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2245: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
  2251: 62c17d486f570e           vmovdqa32 0x380(%r15), %zmm18
  2258: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2265: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  2271: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
  2277: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  2285: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  2292: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2299: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
  2305: 62c17d486f570e           vmovdqa32 0x380(%r15), %zmm18
  2312: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2319: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  2325: 62310548eff8             vpxord %zmm16, %zmm15, %zmm15
  2331: 4d8b5058                 movq 0x58(%r8), %r10
  2335: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  2342: 62c17d486f570f           vmovdqa32 0x3c0(%r15), %zmm18
  2349: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2356: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
  2362: 62c17d486f5710           vmovdqa32 0x400(%r15), %zmm18
  2369: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2376: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  2382: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
  2388: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  2396: 62c17d486f570f           vmovdqa32 0x3c0(%r15), %zmm18
  2403: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2410: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
  2416: 62c17d486f5710           vmovdqa32 0x400(%r15), %zmm18
  2423: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2430: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  2436: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
  2442: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  2450: 62c17d486f570f           vmovdqa32 0x3c0(%r15), %zmm18
  2457: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2464: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
  2470: 62c17d486f5710           vmovdqa32 0x400(%r15), %zmm18
  2477: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2484: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  2490: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
  2496: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  2504: 62c17d486f570f           vmovdqa32 0x3c0(%r15), %zmm18
  2511: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2518: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
  2524: 62c17d486f5710           vmovdqa32 0x400(%r15), %zmm18
  2531: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2538: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  2544: 62310548eff8             vpxord %zmm16, %zmm15, %zmm15
  2550: 4d8b11                   movq (%r9), %r10
  2553: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  2560: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  2568: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
  2576: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
  2584: 4d8b5108                 movq 8(%r9), %r10
  2588: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
  2595: 62917d487f6c1a01         vmovdqa32 %zmm5, 0x40(%r10, %r11)
  2603: 62917d487f741a02         vmovdqa32 %zmm6, 0x80(%r10, %r11)
  2611: 62917d487f7c1a03         vmovdqa32 %zmm7, 0xc0(%r10, %r11)
  2619: 4d8b5110                 movq 0x10(%r9), %r10
  2623: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  2630: 62117d487f4c1a01         vmovdqa32 %zmm9, 0x40(%r10, %r11)
  2638: 62117d487f541a02         vmovdqa32 %zmm10, 0x80(%r10, %r11)
  2646: 62117d487f5c1a03         vmovdqa32 %zmm11, 0xc0(%r10, %r11)
  2654: 4d8b5118                 movq 0x18(%r9), %r10
  2658: 62117d487f241a           vmovdqa32 %zmm12, (%r10, %r11)
  2665: 62117d487f6c1a01         vmovdqa32 %zmm13, 0x40(%r10, %r11)
  2673: 62117d487f741a02         vmovdqa32 %zmm14, 0x80(%r10, %r11)
  2681: 62117d487f7c1a03         vmovdqa32 %zmm15, 0xc0(%r10, %r11)
  2689: 4981c300010000           addq $0x100, %r11
  2696: 4c3bd8                   cmpq %rax, %r11
  2699: 0f86aff5ffff             jbe 0x40
  2705: 4c3bd9                   cmpq %rcx, %r11
  2708: 0f8371050000             jae 0x100b
  2714: 4153                     pushq %r11
  2716: 498dbf40040000           leaq 0x440(%r15), %rdi
  2723: 498db740050000           leaq 0x540(%r15), %rsi
  2730: 4d8b11                   movq (%r9), %r10
  2733: 4c8b1c24                 movq (%rsp), %r11
  2737: 33db                     xorl %ebx, %ebx
  2739: 498b00                   movq (%r8), %rax
  2742: 420fb61418               movzbl (%rax, %r11), %edx
  2747: 32da                     xorb %dl, %bl
  2749: 498b4008                 movq 8(%r8), %rax
  2753: 420fb61418               movzbl (%rax, %r11), %edx
  2758: 4885d2                   testq %rdx, %rdx
  2761: 0f841f000000             je 0xaee
  2767: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2771: 83c019                   addl $0x19, %eax
  2774: 81f8ff000000             cmpl $0xff, %eax
  2780: 0f8206000000             jb 0xae8
  2786: 81c001ffffff             addl $0xffffff01, %eax
  2792: 0fb60406                 movzbl (%rsi, %rax), %eax
  2796: 32d8                     xorb %al, %bl
  2798: 498b4010                 movq 0x10(%r8), %rax
  2802: 420fb61418               movzbl (%rax, %r11), %edx
  2807: 4885d2                   testq %rdx, %rdx
  2810: 0f841f000000             je 0xb1f
  2816: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2820: 83c032                   addl $0x32, %eax
  2823: 81f8ff000000             cmpl $0xff, %eax
  2829: 0f8206000000             jb 0xb19
  2835: 81c001ffffff             addl $0xffffff01, %eax
  2841: 0fb60406                 movzbl (%rsi, %rax), %eax
  2845: 32d8                     xorb %al, %bl
  2847: 498b4018                 movq 0x18(%r8), %rax
  2851: 420fb61418               movzbl (%rax, %r11), %edx
  2856: 4885d2                   testq %rdx, %rdx
  2859: 0f841f000000             je 0xb50
  2865: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2869: 83c04b                   addl $0x4b, %eax
  2872: 81f8ff000000             cmpl $0xff, %eax
  2878: 0f8206000000             jb 0xb4a
  2884: 81c001ffffff             addl $0xffffff01, %eax
  2890: 0fb60406                 movzbl (%rsi, %rax), %eax
  2894: 32d8                     xorb %al, %bl
  2896: 498b4020                 movq 0x20(%r8), %rax
  2900: 420fb61418               movzbl (%rax, %r11), %edx
  2905: 4885d2                   testq %rdx, %rdx
  2908: 0f841f000000             je 0xb81
  2914: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2918: 83c064                   addl $0x64, %eax
  2921: 81f8ff000000             cmpl $0xff, %eax
  2927: 0f8206000000             jb 0xb7b
  2933: 81c001ffffff             addl $0xffffff01, %eax
  2939: 0fb60406                 movzbl (%rsi, %rax), %eax
  2943: 32d8                     xorb %al, %bl
  2945: 498b4028                 movq 0x28(%r8), %rax
  2949: 420fb61418               movzbl (%rax, %r11), %edx
  2954: 4885d2                   testq %rdx, %rdx
  2957: 0f841f000000             je 0xbb2
  2963: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2967: 83c07d                   addl $0x7d, %eax
  2970: 81f8ff000000             cmpl $0xff, %eax
  2976: 0f8206000000             jb 0xbac
  2982: 81c001ffffff             addl $0xffffff01, %eax
  2988: 0fb60406                 movzbl (%rsi, %rax), %eax
  2992: 32d8                     xorb %al, %bl
  2994: 498b4030                 movq 0x30(%r8), %rax
  2998: 420fb61418               movzbl (%rax, %r11), %edx
  3003: 4885d2                   testq %rdx, %rdx
  3006: 0f8422000000             je 0xbe6
  3012: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3016: 81c096000000             addl $0x96, %eax
  3022: 81f8ff000000             cmpl $0xff, %eax
  3028: 0f8206000000             jb 0xbe0
  3034: 81c001ffffff             addl $0xffffff01, %eax
  3040: 0fb60406                 movzbl (%rsi, %rax), %eax
  3044: 32d8                     xorb %al, %bl
  3046: 498b4038                 movq 0x38(%r8), %rax
  3050: 420fb61418               movzbl (%rax, %r11), %edx
  3055: 4885d2                   testq %rdx, %rdx
  3058: 0f8422000000             je 0xc1a
  3064: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3068: 81c0af000000             addl $0xaf, %eax
  3074: 81f8ff000000             cmpl $0xff, %eax
  3080: 0f8206000000             jb 0xc14
  3086: 81c001ffffff             addl $0xffffff01, %eax
  3092: 0fb60406                 movzbl (%rsi, %rax), %eax
  3096: 32d8                     xorb %al, %bl
  3098: 498b4040                 movq 0x40(%r8), %rax
  3102: 420fb61418               movzbl (%rax, %r11), %edx
  3107: 4885d2                   testq %rdx, %rdx
  3110: 0f8422000000             je 0xc4e
  3116: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3120: 81c0c8000000             addl $0xc8, %eax
  3126: 81f8ff000000             cmpl $0xff, %eax
  3132: 0f8206000000             jb 0xc48
  3138: 81c001ffffff             addl $0xffffff01, %eax
  3144: 0fb60406                 movzbl (%rsi, %rax), %eax
  3148: 32d8                     xorb %al, %bl
  3150: 498b4048                 movq 0x48(%r8), %rax
  3154: 420fb61418               movzbl (%rax, %r11), %edx
  3159: 4885d2                   testq %rdx, %rdx
  3162: 0f8422000000             je 0xc82
  3168: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3172: 81c0e1000000             addl $0xe1, %eax
  3178: 81f8ff000000             cmpl $0xff, %eax
  3184: 0f8206000000             jb 0xc7c
  3190: 81c001ffffff             addl $0xffffff01, %eax
  3196: 0fb60406                 movzbl (%rsi, %rax), %eax
  3200: 32d8                     xorb %al, %bl
  3202: 498b4050                 movq 0x50(%r8), %rax
  3206: 420fb61418               movzbl (%rax, %r11), %edx
  3211: 4885d2                   testq %rdx, %rdx
  3214: 0f8422000000             je 0xcb6
  3220: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3224: 81c0fa000000             addl $0xfa, %eax
  3230: 81f8ff000000             cmpl $0xff, %eax
  3236: 0f8206000000             jb 0xcb0
  3242: 81c001ffffff             addl $0xffffff01, %eax
  3248: 0fb60406                 movzbl (%rsi, %rax), %eax
  3252: 32d8                     xorb %al, %bl
  3254: 498b4058                 movq 0x58(%r8), %rax
  3258: 420fb61418               movzbl (%rax, %r11), %edx
  3263: 4885d2                   testq %rdx, %rdx
  3266: 0f841f000000             je 0xce7
  3272: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3276: 83c014                   addl $0x14, %eax
  3279: 81f8ff000000             cmpl $0xff, %eax
  3285: 0f8206000000             jb 0xce1
  3291: 81c001ffffff             addl $0xffffff01, %eax
  3297: 0fb60406                 movzbl (%rsi, %rax), %eax
  3301: 32d8                     xorb %al, %bl
  3303: 43881c1a                 movb %bl, (%r10, %r11)
  3307: 33db                     xorl %ebx, %ebx
  3309: 49ffc3                   incq %r11
  3312: 4c3bd9                   cmpq %rcx, %r11
  3315: 0f82bafdffff             jb 0xab3
  3321: 4d8b5108                 movq 8(%r9), %r10
  3325: 4c8b1c24                 movq (%rsp), %r11
  3329: 33db                     xorl %ebx, %ebx
  3331: 498b00                   movq (%r8), %rax
  3334: 420fb61418               movzbl (%rax, %r11), %edx
  3339: 32da                     xorb %dl, %bl
  3341: 498b4008                 movq 8(%r8), %rax
  3345: 420fb61418               movzbl (%rax, %r11), %edx
  3350: 4885d2                   testq %rdx, %rdx
  3353: 0f841f000000             je 0xd3e
  3359: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3363: 83c032                   addl $0x32, %eax
  3366: 81f8ff000000             cmpl $0xff, %eax
  3372: 0f8206000000             jb 0xd38
  3378: 81c001ffffff             addl $0xffffff01, %eax
  3384: 0fb60406                 movzbl (%rsi, %rax), %eax
  3388: 32d8                     xorb %al, %bl
  3390: 498b4010                 movq 0x10(%r8), %rax
  3394: 420fb61418               movzbl (%rax, %r11), %edx
  3399: 4885d2                   testq %rdx, %rdx
  3402: 0f841f000000             je 0xd6f
  3408: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3412: 83c064                   addl $0x64, %eax
  3415: 81f8ff000000             cmpl $0xff, %eax
  3421: 0f8206000000             jb 0xd69
  3427: 81c001ffffff             addl $0xffffff01, %eax
  3433: 0fb60406                 movzbl (%rsi, %rax), %eax
  3437: 32d8                     xorb %al, %bl
  3439: 498b4018                 movq 0x18(%r8), %rax
  3443: 420fb61418               movzbl (%rax, %r11), %edx
  3448: 4885d2                   testq %rdx, %rdx
  3451: 0f8422000000             je 0xda3
  3457: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3461: 81c096000000             addl $0x96, %eax
  3467: 81f8ff000000             cmpl $0xff, %eax
  3473: 0f8206000000             jb 0xd9d
  3479: 81c001ffffff             addl $0xffffff01, %eax
  3485: 0fb60406                 movzbl (%rsi, %rax), %eax
  3489: 32d8                     xorb %al, %bl
  3491: 498b4020                 movq 0x20(%r8), %rax
  3495: 420fb61418               movzbl (%rax, %r11), %edx
  3500: 4885d2                   testq %rdx, %rdx
  3503: 0f8422000000             je 0xdd7
  3509: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3513: 81c0c8000000             addl $0xc8, %eax
  3519: 81f8ff000000             cmpl $0xff, %eax
  3525: 0f8206000000             jb 0xdd1
  3531: 81c001ffffff             addl $0xffffff01, %eax
  3537: 0fb60406                 movzbl (%rsi, %rax), %eax
  3541: 32d8                     xorb %al, %bl
  3543: 498b4028                 movq 0x28(%r8), %rax
  3547: 420fb61418               movzbl (%rax, %r11), %edx
  3552: 4885d2                   testq %rdx, %rdx
  3555: 0f8422000000             je 0xe0b
  3561: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3565: 81c0fa000000             addl $0xfa, %eax
  3571: 81f8ff000000             cmpl $0xff, %eax
  3577: 0f8206000000             jb 0xe05
  3583: 81c001ffffff             addl $0xffffff01, %eax
  3589: 0fb60406                 movzbl (%rsi, %rax), %eax
  3593: 32d8                     xorb %al, %bl
  3595: 498b4030                 movq 0x30(%r8), %rax
  3599: 420fb61418               movzbl (%rax, %r11), %edx
  3604: 4885d2                   testq %rdx, %rdx
  3607: 0f841f000000             je 0xe3c
  3613: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3617: 83c02d                   addl $0x2d, %eax
  3620: 81f8ff000000             cmpl $0xff, %eax
  3626: 0f8206000000             jb 0xe36
  3632: 81c001ffffff             addl $0xffffff01, %eax
  3638: 0fb60406                 movzbl (%rsi, %rax), %eax
  3642: 32d8                     xorb %al, %bl
  3644: 498b4038                 movq 0x38(%r8), %rax
  3648: 420fb61418               movzbl (%rax, %r11), %edx
  3653: 4885d2                   testq %rdx, %rdx
  3656: 0f841f000000             je 0xe6d
  3662: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3666: 83c05f                   addl $0x5f, %eax
  3669: 81f8ff000000             cmpl $0xff, %eax
  3675: 0f8206000000             jb 0xe67
  3681: 81c001ffffff             addl $0xffffff01, %eax
  3687: 0fb60406                 movzbl (%rsi, %rax), %eax
  3691: 32d8                     xorb %al, %bl
  3693: 498b4040                 movq 0x40(%r8), %rax
  3697: 420fb61418               movzbl (%rax, %r11), %edx
  3702: 4885d2                   testq %rdx, %rdx
  3705: 0f8422000000             je 0xea1
  3711: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3715: 81c091000000             addl $0x91, %eax
  3721: 81f8ff000000             cmpl $0xff, %eax
  3727: 0f8206000000             jb 0xe9b
  3733: 81c001ffffff             addl $0xffffff01, %eax
  3739: 0fb60406                 movzbl (%rsi, %rax), %eax
  3743: 32d8                     xorb %al, %bl
  3745: 498b4048                 movq 0x48(%r8), %rax
  3749: 420fb61418               movzbl (%rax, %r11), %edx
  3754: 4885d2                   testq %rdx, %rdx
  3757: 0f8422000000             je 0xed5
  3763: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3767: 81c0c3000000             addl $0xc3, %eax
  3773: 81f8ff000000             cmpl $0xff, %eax
  3779: 0f8206000000             jb 0xecf
  3785: 81c001ffffff             addl $0xffffff01, %eax
  3791: 0fb60406                 movzbl (%rsi, %rax), %eax
  3795: 32d8                     xorb %al, %bl
  3797: 498b4050                 movq 0x50(%r8), %rax
  3801: 420fb61418               movzbl (%rax, %r11), %edx
  3806: 4885d2                   testq %rdx, %rdx
  3809: 0f8422000000             je 0xf09
  3815: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3819: 81c0f5000000             addl $0xf5, %eax
  3825: 81f8ff000000             cmpl $0xff, %eax
  3831: 0f8206000000             jb 0xf03
  3837: 81c001ffffff             addl $0xffffff01, %eax
  3843: 0fb60406                 movzbl (%rsi, %rax), %eax
  3847: 32d8                     xorb %al, %bl
  3849: 498b4058                 movq 0x58(%r8), %rax
  3853: 420fb61418               movzbl (%rax, %r11), %edx
  3858: 4885d2                   testq %rdx, %rdx
  3861: 0f841f000000             je 0xf3a
  3867: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3871: 83c028                   addl $0x28, %eax
  3874: 81f8ff000000             cmpl $0xff, %eax
  3880: 0f8206000000             jb 0xf34
  3886: 81c001ffffff             addl $0xffffff01, %eax
  3892: 0fb60406                 movzbl (%rsi, %rax), %eax
  3896: 32d8                     xorb %al, %bl
  3898: 43881c1a                 movb %bl, (%r10, %r11)
  3902: 33db                     xorl %ebx, %ebx
  3904: 49ffc3                   incq %r11
  3907: 4c3bd9                   cmpq %rcx, %r11
  3910: 0f82b7fdffff             jb 0xd03
  3916: 4d8b5110                 movq 0x10(%r9), %r10
  3920: 4c8b1c24                 movq (%rsp), %r11
  3924: 33db                     xorl %ebx, %ebx
  3926: 498b00                   movq (%r8), %rax
  3929: 420fb61418               movzbl (%rax, %r11), %edx
  3934: 32da                     xorb %dl, %bl
  3936: 498b4008                 movq 8(%r8), %rax
  3940: 420fb61418               movzbl (%rax, %r11), %edx
  3945: 32da                     xorb %dl, %bl
  3947: 498b4010                 movq 0x10(%r8), %rax
  3951: 420fb61418               movzbl (%rax, %r11), %edx
  3956: 32da                     xorb %dl, %bl
  3958: 498b4018                 movq 0x18(%r8), %rax
  3962: 420fb61418               movzbl (%rax, %r11), %edx
  3967: 32da                     xorb %dl, %bl
  3969: 498b4020                 movq 0x20(%r8), %rax
  3973: 420fb61418               movzbl (%rax, %r11), %edx
  3978: 32da                     xorb %dl, %bl
  3980: 498b4028                 movq 0x28(%r8), %rax
  3984: 420fb61418               movzbl (%rax, %r11), %edx
  3989: 32da                     xorb %dl, %bl
  3991: 43881c1a                 movb %bl, (%r10, %r11)
  3995: 33db                     xorl %ebx, %ebx
  3997: 49ffc3                   incq %r11
  4000: 4c3bd9                   cmpq %rcx, %r11
  4003: 0f82adffffff             jb 0xf56
  4009: 4d8b5118                 movq 0x18(%r9), %r10
  4013: 4c8b1c24                 movq (%rsp), %r11
  4017: 33db                     xorl %ebx, %ebx
  4019: 498b4030                 movq 0x30(%r8), %rax
  4023: 420fb61418               movzbl (%rax, %r11), %edx
  4028: 32da                     xorb %dl, %bl
  4030: 498b4038                 movq 0x38(%r8), %rax
  4034: 420fb61418               movzbl (%rax, %r11), %edx
  4039: 32da                     xorb %dl, %bl
  4041: 498b4040                 movq 0x40(%r8), %rax
  4045: 420fb61418               movzbl (%rax, %r11), %edx
  4050: 32da                     xorb %dl, %bl
  4052: 498b4048                 movq 0x48(%r8), %rax
  4056: 420fb61418               movzbl (%rax, %r11), %edx
  4061: 32da                     xorb %dl, %bl
  4063: 498b4050                 movq 0x50(%r8), %rax
  4067: 420fb61418               movzbl (%rax, %r11), %edx
  4072: 32da                     xorb %dl, %bl
  4074: 498b4058                 movq 0x58(%r8), %rax
  4078: 420fb61418               movzbl (%rax, %r11), %edx
  4083: 32da                     xorb %dl, %bl
  4085: 43881c1a                 movb %bl, (%r10, %r11)
  4089: 33db                     xorl %ebx, %ebx
  4091: 49ffc3                   incq %r11
  4094: 4c3bd9                   cmpq %rcx, %r11
  4097: 0f82acffffff             jb 0xfb3
  4103: 4883c408                 addq $8, %rsp
  4107: 415e                     popq %r14
  4109: 415d                     popq %r13
  4111: 415c                     popq %r12
  4113: 415f                     popq %r15
  4115: 5b                       popq %rbx
  4116: c5f877                   vzeroupper 
  4119: c3                       retq 
  4120: 0f0b                     ud2 
  4122: 0f0b                     ud2 
  4124: 0f0b                     ud2 
  4126: 0f0b                     ud2 
  4128: 0f0b                     ud2 
  4130: 0f0b                     ud2 
  4132: 0f0b                     ud2 
  4134: 0f0b                     ud2 
  4136: 0f0b                     ud2 
  4138: 0f0b                     ud2 
  4140: 0f0b                     ud2 
  4142: 0f0b                     ud2 
  4144: 0f0b                     ud2 
  4146: 0f0b                     ud2 
  4148: 0f0b                     ud2 
  4150: 0f0b                     ud2 
  4152: 0f0b                     ud2 
  4154: 0f0b                     ud2 
  4156: 0f0b                     ud2 
  4158: 0f0b                     ud2 
