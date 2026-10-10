; config_t1_43 target=avx512-gfni twin=notail topo=raid6(8)/enc op=encode U=8
; code_size=1984 twins=["notail"] uops_per_byte=0.53515625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da7070000           leaq 0x7a7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e800020000           subq $0x200, %rax
    57: 0f8253070000             jb 0x792
    63: 90                       nop 
    64: 62a17d486f041e           vmovdqa32 (%rsi, %r11), %zmm16
    71: 62b17d486fc0             vmovdqa32 %zmm16, %zmm0
    77: 62c17d486f17             vmovdqa32 (%r15), %zmm18
    83: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
    90: 62317d486fc1             vmovdqa32 %zmm17, %zmm8
    96: 62a17d486f441e01         vmovdqa32 0x40(%rsi, %r11), %zmm16
   104: 62b17d486fc8             vmovdqa32 %zmm16, %zmm1
   110: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   117: 62317d486fc9             vmovdqa32 %zmm17, %zmm9
   123: 62a17d486f441e02         vmovdqa32 0x80(%rsi, %r11), %zmm16
   131: 62b17d486fd0             vmovdqa32 %zmm16, %zmm2
   137: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   144: 62317d486fd1             vmovdqa32 %zmm17, %zmm10
   150: 62a17d486f441e03         vmovdqa32 0xc0(%rsi, %r11), %zmm16
   158: 62b17d486fd8             vmovdqa32 %zmm16, %zmm3
   164: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   171: 62317d486fd9             vmovdqa32 %zmm17, %zmm11
   177: 62a17d486f441e04         vmovdqa32 0x100(%rsi, %r11), %zmm16
   185: 62b17d486fe0             vmovdqa32 %zmm16, %zmm4
   191: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   198: 62317d486fe1             vmovdqa32 %zmm17, %zmm12
   204: 62a17d486f441e05         vmovdqa32 0x140(%rsi, %r11), %zmm16
   212: 62b17d486fe8             vmovdqa32 %zmm16, %zmm5
   218: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   225: 62317d486fe9             vmovdqa32 %zmm17, %zmm13
   231: 62a17d486f441e06         vmovdqa32 0x180(%rsi, %r11), %zmm16
   239: 62b17d486ff0             vmovdqa32 %zmm16, %zmm6
   245: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   252: 62317d486ff1             vmovdqa32 %zmm17, %zmm14
   258: 62a17d486f441e07         vmovdqa32 0x1c0(%rsi, %r11), %zmm16
   266: 62b17d486ff8             vmovdqa32 %zmm16, %zmm7
   272: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   279: 62317d486ff9             vmovdqa32 %zmm17, %zmm15
   285: 62a17d486f041f           vmovdqa32 (%rdi, %r11), %zmm16
   292: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
   298: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   305: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   312: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
   318: 62a17d486f441f01         vmovdqa32 0x40(%rdi, %r11), %zmm16
   326: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
   332: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   339: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
   345: 62a17d486f441f02         vmovdqa32 0x80(%rdi, %r11), %zmm16
   353: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
   359: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   366: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
   372: 62a17d486f441f03         vmovdqa32 0xc0(%rdi, %r11), %zmm16
   380: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
   386: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   393: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
   399: 62a17d486f441f04         vmovdqa32 0x100(%rdi, %r11), %zmm16
   407: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
   413: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   420: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
   426: 62a17d486f441f05         vmovdqa32 0x140(%rdi, %r11), %zmm16
   434: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
   440: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   447: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
   453: 62a17d486f441f06         vmovdqa32 0x180(%rdi, %r11), %zmm16
   461: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
   467: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   474: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
   480: 62a17d486f441f07         vmovdqa32 0x1c0(%rdi, %r11), %zmm16
   488: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
   494: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   501: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
   507: 62817d486f041c           vmovdqa32 (%r12, %r11), %zmm16
   514: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
   520: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   527: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   534: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
   540: 62817d486f441c01         vmovdqa32 0x40(%r12, %r11), %zmm16
   548: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
   554: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   561: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
   567: 62817d486f441c02         vmovdqa32 0x80(%r12, %r11), %zmm16
   575: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
   581: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   588: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
   594: 62817d486f441c03         vmovdqa32 0xc0(%r12, %r11), %zmm16
   602: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
   608: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   615: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
   621: 62817d486f441c04         vmovdqa32 0x100(%r12, %r11), %zmm16
   629: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
   635: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   642: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
   648: 62817d486f441c05         vmovdqa32 0x140(%r12, %r11), %zmm16
   656: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
   662: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   669: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
   675: 62817d486f441c06         vmovdqa32 0x180(%r12, %r11), %zmm16
   683: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
   689: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   696: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
   702: 62817d486f441c07         vmovdqa32 0x1c0(%r12, %r11), %zmm16
   710: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
   716: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   723: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
   729: 62817d486f441d00         vmovdqa32 (%r13, %r11), %zmm16
   737: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
   743: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   750: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   757: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
   763: 62817d486f441d01         vmovdqa32 0x40(%r13, %r11), %zmm16
   771: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
   777: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   784: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
   790: 62817d486f441d02         vmovdqa32 0x80(%r13, %r11), %zmm16
   798: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
   804: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   811: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
   817: 62817d486f441d03         vmovdqa32 0xc0(%r13, %r11), %zmm16
   825: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
   831: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   838: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
   844: 62817d486f441d04         vmovdqa32 0x100(%r13, %r11), %zmm16
   852: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
   858: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   865: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
   871: 62817d486f441d05         vmovdqa32 0x140(%r13, %r11), %zmm16
   879: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
   885: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   892: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
   898: 62817d486f441d06         vmovdqa32 0x180(%r13, %r11), %zmm16
   906: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
   912: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   919: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
   925: 62817d486f441d07         vmovdqa32 0x1c0(%r13, %r11), %zmm16
   933: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
   939: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   946: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
   952: 62817d486f041e           vmovdqa32 (%r14, %r11), %zmm16
   959: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
   965: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
   972: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   979: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
   985: 62817d486f441e01         vmovdqa32 0x40(%r14, %r11), %zmm16
   993: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
   999: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1006: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
  1012: 62817d486f441e02         vmovdqa32 0x80(%r14, %r11), %zmm16
  1020: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
  1026: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1033: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
  1039: 62817d486f441e03         vmovdqa32 0xc0(%r14, %r11), %zmm16
  1047: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
  1053: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1060: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
  1066: 62817d486f441e04         vmovdqa32 0x100(%r14, %r11), %zmm16
  1074: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
  1080: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1087: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
  1093: 62817d486f441e05         vmovdqa32 0x140(%r14, %r11), %zmm16
  1101: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
  1107: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1114: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
  1120: 62817d486f441e06         vmovdqa32 0x180(%r14, %r11), %zmm16
  1128: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
  1134: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1141: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
  1147: 62817d486f441e07         vmovdqa32 0x1c0(%r14, %r11), %zmm16
  1155: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
  1161: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1168: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
  1174: 4d8b5028                 movq 0x28(%r8), %r10
  1178: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  1185: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
  1191: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1198: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1205: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
  1211: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  1219: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
  1225: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1232: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
  1238: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  1246: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
  1252: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1259: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
  1265: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  1273: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
  1279: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1286: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
  1292: 62817d486f441a04         vmovdqa32 0x100(%r10, %r11), %zmm16
  1300: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
  1306: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1313: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
  1319: 62817d486f441a05         vmovdqa32 0x140(%r10, %r11), %zmm16
  1327: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
  1333: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1340: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
  1346: 62817d486f441a06         vmovdqa32 0x180(%r10, %r11), %zmm16
  1354: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
  1360: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1367: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
  1373: 62817d486f441a07         vmovdqa32 0x1c0(%r10, %r11), %zmm16
  1381: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
  1387: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1394: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
  1400: 4d8b5030                 movq 0x30(%r8), %r10
  1404: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  1411: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
  1417: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1424: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1431: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
  1437: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  1445: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
  1451: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1458: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
  1464: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  1472: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
  1478: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1485: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
  1491: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  1499: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
  1505: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1512: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
  1518: 62817d486f441a04         vmovdqa32 0x100(%r10, %r11), %zmm16
  1526: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
  1532: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1539: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
  1545: 62817d486f441a05         vmovdqa32 0x140(%r10, %r11), %zmm16
  1553: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
  1559: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1566: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
  1572: 62817d486f441a06         vmovdqa32 0x180(%r10, %r11), %zmm16
  1580: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
  1586: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1593: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
  1599: 62817d486f441a07         vmovdqa32 0x1c0(%r10, %r11), %zmm16
  1607: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
  1613: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1620: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
  1626: 4d8b5038                 movq 0x38(%r8), %r10
  1630: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  1637: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
  1643: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
  1649: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  1657: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
  1663: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
  1669: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  1677: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
  1683: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
  1689: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  1697: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
  1703: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
  1709: 62817d486f441a04         vmovdqa32 0x100(%r10, %r11), %zmm16
  1717: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
  1723: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
  1729: 62817d486f441a05         vmovdqa32 0x140(%r10, %r11), %zmm16
  1737: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
  1743: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
  1749: 62817d486f441a06         vmovdqa32 0x180(%r10, %r11), %zmm16
  1757: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
  1763: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
  1769: 62817d486f441a07         vmovdqa32 0x1c0(%r10, %r11), %zmm16
  1777: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
  1783: 62310548eff8             vpxord %zmm16, %zmm15, %zmm15
  1789: 4d8b11                   movq (%r9), %r10
  1792: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  1799: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  1807: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
  1815: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
  1823: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
  1831: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
  1839: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
  1847: 62917d487f7c1a07         vmovdqa32 %zmm7, 0x1c0(%r10, %r11)
  1855: 4d8b5108                 movq 8(%r9), %r10
  1859: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  1866: 62117d487f4c1a01         vmovdqa32 %zmm9, 0x40(%r10, %r11)
  1874: 62117d487f541a02         vmovdqa32 %zmm10, 0x80(%r10, %r11)
  1882: 62117d487f5c1a03         vmovdqa32 %zmm11, 0xc0(%r10, %r11)
  1890: 62117d487f641a04         vmovdqa32 %zmm12, 0x100(%r10, %r11)
  1898: 62117d487f6c1a05         vmovdqa32 %zmm13, 0x140(%r10, %r11)
  1906: 62117d487f741a06         vmovdqa32 %zmm14, 0x180(%r10, %r11)
  1914: 62117d487f7c1a07         vmovdqa32 %zmm15, 0x1c0(%r10, %r11)
  1922: 4981c300020000           addq $0x200, %r11
  1929: 4c3bd8                   cmpq %rax, %r11
  1932: 0f86aef8ffff             jbe 0x40
  1938: 415e                     popq %r14
  1940: 415d                     popq %r13
  1942: 415c                     popq %r12
  1944: 415f                     popq %r15
  1946: 5b                       popq %rbx
  1947: c5f877                   vzeroupper 
  1950: c3                       retq 
  1951: 0f0b                     ud2 
  1953: 0f0b                     ud2 
  1955: 0f0b                     ud2 
  1957: 0f0b                     ud2 
  1959: 0f0b                     ud2 
  1961: 0f0b                     ud2 
  1963: 0f0b                     ud2 
  1965: 0f0b                     ud2 
  1967: 0f0b                     ud2 
  1969: 0f0b                     ud2 
  1971: 0f0b                     ud2 
  1973: 0f0b                     ud2 
  1975: 0f0b                     ud2 
  1977: 0f0b                     ud2 
  1979: 0f0b                     ud2 
  1981: 0f0b                     ud2 
  1983: cc                       int3 
