; config_t1_40 target=avx512-gfni twin=nt topo=raid6(8)/enc op=encode U=8
; code_size=2496 twins=["nt"] uops_per_byte=0.53515625
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
  1792: 62917d48e7041a           vmovntdq %zmm0, (%r10, %r11)
  1799: 62917d48e74c1a01         vmovntdq %zmm1, 0x40(%r10, %r11)
  1807: 62917d48e7541a02         vmovntdq %zmm2, 0x80(%r10, %r11)
  1815: 62917d48e75c1a03         vmovntdq %zmm3, 0xc0(%r10, %r11)
  1823: 62917d48e7641a04         vmovntdq %zmm4, 0x100(%r10, %r11)
  1831: 62917d48e76c1a05         vmovntdq %zmm5, 0x140(%r10, %r11)
  1839: 62917d48e7741a06         vmovntdq %zmm6, 0x180(%r10, %r11)
  1847: 62917d48e77c1a07         vmovntdq %zmm7, 0x1c0(%r10, %r11)
  1855: 4d8b5108                 movq 8(%r9), %r10
  1859: 62117d48e7041a           vmovntdq %zmm8, (%r10, %r11)
  1866: 62117d48e74c1a01         vmovntdq %zmm9, 0x40(%r10, %r11)
  1874: 62117d48e7541a02         vmovntdq %zmm10, 0x80(%r10, %r11)
  1882: 62117d48e75c1a03         vmovntdq %zmm11, 0xc0(%r10, %r11)
  1890: 62117d48e7641a04         vmovntdq %zmm12, 0x100(%r10, %r11)
  1898: 62117d48e76c1a05         vmovntdq %zmm13, 0x140(%r10, %r11)
  1906: 62117d48e7741a06         vmovntdq %zmm14, 0x180(%r10, %r11)
  1914: 62117d48e77c1a07         vmovntdq %zmm15, 0x1c0(%r10, %r11)
  1922: 4981c300020000           addq $0x200, %r11
  1929: 4c3bd8                   cmpq %rax, %r11
  1932: 0f86aef8ffff             jbe 0x40
  1938: 4c3bd9                   cmpq %rcx, %r11
  1941: 0f830f020000             jae 0x9aa
  1947: 4153                     pushq %r11
  1949: 498dbfc0010000           leaq 0x1c0(%r15), %rdi
  1956: 498db7c0020000           leaq 0x2c0(%r15), %rsi
  1963: 4d8b11                   movq (%r9), %r10
  1966: 4c8b1c24                 movq (%rsp), %r11
  1970: 33db                     xorl %ebx, %ebx
  1972: 498b00                   movq (%r8), %rax
  1975: 420fb61418               movzbl (%rax, %r11), %edx
  1980: 32da                     xorb %dl, %bl
  1982: 498b4008                 movq 8(%r8), %rax
  1986: 420fb61418               movzbl (%rax, %r11), %edx
  1991: 32da                     xorb %dl, %bl
  1993: 498b4010                 movq 0x10(%r8), %rax
  1997: 420fb61418               movzbl (%rax, %r11), %edx
  2002: 32da                     xorb %dl, %bl
  2004: 498b4018                 movq 0x18(%r8), %rax
  2008: 420fb61418               movzbl (%rax, %r11), %edx
  2013: 32da                     xorb %dl, %bl
  2015: 498b4020                 movq 0x20(%r8), %rax
  2019: 420fb61418               movzbl (%rax, %r11), %edx
  2024: 32da                     xorb %dl, %bl
  2026: 498b4028                 movq 0x28(%r8), %rax
  2030: 420fb61418               movzbl (%rax, %r11), %edx
  2035: 32da                     xorb %dl, %bl
  2037: 498b4030                 movq 0x30(%r8), %rax
  2041: 420fb61418               movzbl (%rax, %r11), %edx
  2046: 32da                     xorb %dl, %bl
  2048: 498b4038                 movq 0x38(%r8), %rax
  2052: 420fb61418               movzbl (%rax, %r11), %edx
  2057: 32da                     xorb %dl, %bl
  2059: 43881c1a                 movb %bl, (%r10, %r11)
  2063: 33db                     xorl %ebx, %ebx
  2065: 49ffc3                   incq %r11
  2068: 4c3bd9                   cmpq %rcx, %r11
  2071: 0f8297ffffff             jb 0x7b4
  2077: 4d8b5108                 movq 8(%r9), %r10
  2081: 4c8b1c24                 movq (%rsp), %r11
  2085: 33db                     xorl %ebx, %ebx
  2087: 498b00                   movq (%r8), %rax
  2090: 420fb61418               movzbl (%rax, %r11), %edx
  2095: 4885d2                   testq %rdx, %rdx
  2098: 0f8422000000             je 0x85a
  2104: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2108: 81c0e6000000             addl $0xe6, %eax
  2114: 81f8ff000000             cmpl $0xff, %eax
  2120: 0f8206000000             jb 0x854
  2126: 81c001ffffff             addl $0xffffff01, %eax
  2132: 0fb60406                 movzbl (%rsi, %rax), %eax
  2136: 32d8                     xorb %al, %bl
  2138: 498b4008                 movq 8(%r8), %rax
  2142: 420fb61418               movzbl (%rax, %r11), %edx
  2147: 4885d2                   testq %rdx, %rdx
  2150: 0f8422000000             je 0x88e
  2156: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2160: 81c0cd000000             addl $0xcd, %eax
  2166: 81f8ff000000             cmpl $0xff, %eax
  2172: 0f8206000000             jb 0x888
  2178: 81c001ffffff             addl $0xffffff01, %eax
  2184: 0fb60406                 movzbl (%rsi, %rax), %eax
  2188: 32d8                     xorb %al, %bl
  2190: 498b4010                 movq 0x10(%r8), %rax
  2194: 420fb61418               movzbl (%rax, %r11), %edx
  2199: 4885d2                   testq %rdx, %rdx
  2202: 0f8422000000             je 0x8c2
  2208: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2212: 81c0e7000000             addl $0xe7, %eax
  2218: 81f8ff000000             cmpl $0xff, %eax
  2224: 0f8206000000             jb 0x8bc
  2230: 81c001ffffff             addl $0xffffff01, %eax
  2236: 0fb60406                 movzbl (%rsi, %rax), %eax
  2240: 32d8                     xorb %al, %bl
  2242: 498b4018                 movq 0x18(%r8), %rax
  2246: 420fb61418               movzbl (%rax, %r11), %edx
  2251: 4885d2                   testq %rdx, %rdx
  2254: 0f8422000000             je 0x8f6
  2260: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2264: 81c0ce000000             addl $0xce, %eax
  2270: 81f8ff000000             cmpl $0xff, %eax
  2276: 0f8206000000             jb 0x8f0
  2282: 81c001ffffff             addl $0xffffff01, %eax
  2288: 0fb60406                 movzbl (%rsi, %rax), %eax
  2292: 32d8                     xorb %al, %bl
  2294: 498b4020                 movq 0x20(%r8), %rax
  2298: 420fb61418               movzbl (%rax, %r11), %edx
  2303: 4885d2                   testq %rdx, %rdx
  2306: 0f841f000000             je 0x927
  2312: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2316: 83c04b                   addl $0x4b, %eax
  2319: 81f8ff000000             cmpl $0xff, %eax
  2325: 0f8206000000             jb 0x921
  2331: 81c001ffffff             addl $0xffffff01, %eax
  2337: 0fb60406                 movzbl (%rsi, %rax), %eax
  2341: 32d8                     xorb %al, %bl
  2343: 498b4028                 movq 0x28(%r8), %rax
  2347: 420fb61418               movzbl (%rax, %r11), %edx
  2352: 4885d2                   testq %rdx, %rdx
  2355: 0f841f000000             je 0x958
  2361: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2365: 83c032                   addl $0x32, %eax
  2368: 81f8ff000000             cmpl $0xff, %eax
  2374: 0f8206000000             jb 0x952
  2380: 81c001ffffff             addl $0xffffff01, %eax
  2386: 0fb60406                 movzbl (%rsi, %rax), %eax
  2390: 32d8                     xorb %al, %bl
  2392: 498b4030                 movq 0x30(%r8), %rax
  2396: 420fb61418               movzbl (%rax, %r11), %edx
  2401: 4885d2                   testq %rdx, %rdx
  2404: 0f841f000000             je 0x989
  2410: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2414: 83c019                   addl $0x19, %eax
  2417: 81f8ff000000             cmpl $0xff, %eax
  2423: 0f8206000000             jb 0x983
  2429: 81c001ffffff             addl $0xffffff01, %eax
  2435: 0fb60406                 movzbl (%rsi, %rax), %eax
  2439: 32d8                     xorb %al, %bl
  2441: 498b4038                 movq 0x38(%r8), %rax
  2445: 420fb61418               movzbl (%rax, %r11), %edx
  2450: 32da                     xorb %dl, %bl
  2452: 43881c1a                 movb %bl, (%r10, %r11)
  2456: 33db                     xorl %ebx, %ebx
  2458: 49ffc3                   incq %r11
  2461: 4c3bd9                   cmpq %rcx, %r11
  2464: 0f8281feffff             jb 0x827
  2470: 4883c408                 addq $8, %rsp
  2474: 415e                     popq %r14
  2476: 415d                     popq %r13
  2478: 415c                     popq %r12
  2480: 415f                     popq %r15
  2482: 5b                       popq %rbx
  2483: c5f877                   vzeroupper 
  2486: c3                       retq 
  2487: 0f0b                     ud2 
  2489: 0f0b                     ud2 
  2491: 0f0b                     ud2 
  2493: 0f0b                     ud2 
  2495: cc                       int3 
