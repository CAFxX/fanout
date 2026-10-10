; config_t1_207 target=avx512-gfni twin=rec_cf topo=rs(6,3)/rec{0,1} op=recover U=8
; code_size=2496 twins=[] uops_per_byte=0.4765625
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
    57: 0f8299060000             jb 0x6d8
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
   743: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
   749: 62817d486f441d01         vmovdqa32 0x40(%r13, %r11), %zmm16
   757: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
   763: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
   769: 62817d486f441d02         vmovdqa32 0x80(%r13, %r11), %zmm16
   777: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
   783: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
   789: 62817d486f441d03         vmovdqa32 0xc0(%r13, %r11), %zmm16
   797: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
   803: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
   809: 62817d486f441d04         vmovdqa32 0x100(%r13, %r11), %zmm16
   817: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
   823: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
   829: 62817d486f441d05         vmovdqa32 0x140(%r13, %r11), %zmm16
   837: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
   843: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
   849: 62817d486f441d06         vmovdqa32 0x180(%r13, %r11), %zmm16
   857: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
   863: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
   869: 62817d486f441d07         vmovdqa32 0x1c0(%r13, %r11), %zmm16
   877: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
   883: 62310548eff8             vpxord %zmm16, %zmm15, %zmm15
   889: 62817d486f041e           vmovdqa32 (%r14, %r11), %zmm16
   896: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
   902: 62817d486f441e01         vmovdqa32 0x40(%r14, %r11), %zmm16
   910: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
   916: 62817d486f441e02         vmovdqa32 0x80(%r14, %r11), %zmm16
   924: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
   930: 62817d486f441e03         vmovdqa32 0xc0(%r14, %r11), %zmm16
   938: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
   944: 62817d486f441e04         vmovdqa32 0x100(%r14, %r11), %zmm16
   952: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
   958: 62817d486f441e05         vmovdqa32 0x140(%r14, %r11), %zmm16
   966: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
   972: 62817d486f441e06         vmovdqa32 0x180(%r14, %r11), %zmm16
   980: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
   986: 62817d486f441e07         vmovdqa32 0x1c0(%r14, %r11), %zmm16
   994: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
  1000: 4d8b5028                 movq 0x28(%r8), %r10
  1004: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  1011: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
  1017: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  1025: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
  1031: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  1039: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
  1045: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  1053: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
  1059: 62817d486f441a04         vmovdqa32 0x100(%r10, %r11), %zmm16
  1067: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
  1073: 62817d486f441a05         vmovdqa32 0x140(%r10, %r11), %zmm16
  1081: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
  1087: 62817d486f441a06         vmovdqa32 0x180(%r10, %r11), %zmm16
  1095: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
  1101: 62817d486f441a07         vmovdqa32 0x1c0(%r10, %r11), %zmm16
  1109: 62310548eff8             vpxord %zmm16, %zmm15, %zmm15
  1115: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
  1122: 62a3fd48ceca00           vgf2p8affineqb $0, %zmm18, %zmm0, %zmm17
  1129: 62c17540efc8             vpxord %zmm8, %zmm17, %zmm17
  1135: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1142: 62a3fd48cec200           vgf2p8affineqb $0, %zmm18, %zmm0, %zmm16
  1149: 62c17d40efc0             vpxord %zmm8, %zmm16, %zmm16
  1155: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1162: 6233f540cec200           vgf2p8affineqb $0, %zmm18, %zmm17, %zmm8
  1169: 62b3fd40cec200           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm0
  1176: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
  1183: 62a3f548ceca00           vgf2p8affineqb $0, %zmm18, %zmm1, %zmm17
  1190: 62c17540efc9             vpxord %zmm9, %zmm17, %zmm17
  1196: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1203: 62a3f548cec200           vgf2p8affineqb $0, %zmm18, %zmm1, %zmm16
  1210: 62c17d40efc1             vpxord %zmm9, %zmm16, %zmm16
  1216: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1223: 6233f540ceca00           vgf2p8affineqb $0, %zmm18, %zmm17, %zmm9
  1230: 62b3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm1
  1237: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
  1244: 62a3ed48ceca00           vgf2p8affineqb $0, %zmm18, %zmm2, %zmm17
  1251: 62c17540efca             vpxord %zmm10, %zmm17, %zmm17
  1257: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1264: 62a3ed48cec200           vgf2p8affineqb $0, %zmm18, %zmm2, %zmm16
  1271: 62c17d40efc2             vpxord %zmm10, %zmm16, %zmm16
  1277: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1284: 6233f540ced200           vgf2p8affineqb $0, %zmm18, %zmm17, %zmm10
  1291: 62b3fd40ced200           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm2
  1298: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
  1305: 62a3e548ceca00           vgf2p8affineqb $0, %zmm18, %zmm3, %zmm17
  1312: 62c17540efcb             vpxord %zmm11, %zmm17, %zmm17
  1318: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1325: 62a3e548cec200           vgf2p8affineqb $0, %zmm18, %zmm3, %zmm16
  1332: 62c17d40efc3             vpxord %zmm11, %zmm16, %zmm16
  1338: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1345: 6233f540ceda00           vgf2p8affineqb $0, %zmm18, %zmm17, %zmm11
  1352: 62b3fd40ceda00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm3
  1359: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
  1366: 62a3dd48ceca00           vgf2p8affineqb $0, %zmm18, %zmm4, %zmm17
  1373: 62c17540efcc             vpxord %zmm12, %zmm17, %zmm17
  1379: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1386: 62a3dd48cec200           vgf2p8affineqb $0, %zmm18, %zmm4, %zmm16
  1393: 62c17d40efc4             vpxord %zmm12, %zmm16, %zmm16
  1399: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1406: 6233f540cee200           vgf2p8affineqb $0, %zmm18, %zmm17, %zmm12
  1413: 62b3fd40cee200           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm4
  1420: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
  1427: 62a3d548ceca00           vgf2p8affineqb $0, %zmm18, %zmm5, %zmm17
  1434: 62c17540efcd             vpxord %zmm13, %zmm17, %zmm17
  1440: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1447: 62a3d548cec200           vgf2p8affineqb $0, %zmm18, %zmm5, %zmm16
  1454: 62c17d40efc5             vpxord %zmm13, %zmm16, %zmm16
  1460: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1467: 6233f540ceea00           vgf2p8affineqb $0, %zmm18, %zmm17, %zmm13
  1474: 62b3fd40ceea00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm5
  1481: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
  1488: 62a3cd48ceca00           vgf2p8affineqb $0, %zmm18, %zmm6, %zmm17
  1495: 62c17540efce             vpxord %zmm14, %zmm17, %zmm17
  1501: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1508: 62a3cd48cec200           vgf2p8affineqb $0, %zmm18, %zmm6, %zmm16
  1515: 62c17d40efc6             vpxord %zmm14, %zmm16, %zmm16
  1521: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1528: 6233f540cef200           vgf2p8affineqb $0, %zmm18, %zmm17, %zmm14
  1535: 62b3fd40cef200           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm6
  1542: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
  1549: 62a3c548ceca00           vgf2p8affineqb $0, %zmm18, %zmm7, %zmm17
  1556: 62c17540efcf             vpxord %zmm15, %zmm17, %zmm17
  1562: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1569: 62a3c548cec200           vgf2p8affineqb $0, %zmm18, %zmm7, %zmm16
  1576: 62c17d40efc7             vpxord %zmm15, %zmm16, %zmm16
  1582: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1589: 6233f540cefa00           vgf2p8affineqb $0, %zmm18, %zmm17, %zmm15
  1596: 62b3fd40cefa00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm7
  1603: 4d8b11                   movq (%r9), %r10
  1606: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  1613: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  1621: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
  1629: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
  1637: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
  1645: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
  1653: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
  1661: 62917d487f7c1a07         vmovdqa32 %zmm7, 0x1c0(%r10, %r11)
  1669: 4d8b5108                 movq 8(%r9), %r10
  1673: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  1680: 62117d487f4c1a01         vmovdqa32 %zmm9, 0x40(%r10, %r11)
  1688: 62117d487f541a02         vmovdqa32 %zmm10, 0x80(%r10, %r11)
  1696: 62117d487f5c1a03         vmovdqa32 %zmm11, 0xc0(%r10, %r11)
  1704: 62117d487f641a04         vmovdqa32 %zmm12, 0x100(%r10, %r11)
  1712: 62117d487f6c1a05         vmovdqa32 %zmm13, 0x140(%r10, %r11)
  1720: 62117d487f741a06         vmovdqa32 %zmm14, 0x180(%r10, %r11)
  1728: 62117d487f7c1a07         vmovdqa32 %zmm15, 0x1c0(%r10, %r11)
  1736: 4981c300020000           addq $0x200, %r11
  1743: 4c3bd8                   cmpq %rax, %r11
  1746: 0f8668f9ffff             jbe 0x40
  1752: 4c3bd9                   cmpq %rcx, %r11
  1755: 0f83a1020000             jae 0x982
  1761: 4153                     pushq %r11
  1763: 498dbf80010000           leaq 0x180(%r15), %rdi
  1770: 498db780020000           leaq 0x280(%r15), %rsi
  1777: 4d8b11                   movq (%r9), %r10
  1780: 4c8b1c24                 movq (%rsp), %r11
  1784: 33db                     xorl %ebx, %ebx
  1786: 498b00                   movq (%r8), %rax
  1789: 420fb61418               movzbl (%rax, %r11), %edx
  1794: 4885d2                   testq %rdx, %rdx
  1797: 0f8422000000             je 0x72d
  1803: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1807: 81c0c1000000             addl $0xc1, %eax
  1813: 81f8ff000000             cmpl $0xff, %eax
  1819: 0f8206000000             jb 0x727
  1825: 81c001ffffff             addl $0xffffff01, %eax
  1831: 0fb60406                 movzbl (%rsi, %rax), %eax
  1835: 32d8                     xorb %al, %bl
  1837: 498b4008                 movq 8(%r8), %rax
  1841: 420fb61418               movzbl (%rax, %r11), %edx
  1846: 4885d2                   testq %rdx, %rdx
  1849: 0f841f000000             je 0x75e
  1855: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1859: 83c003                   addl $3, %eax
  1862: 81f8ff000000             cmpl $0xff, %eax
  1868: 0f8206000000             jb 0x758
  1874: 81c001ffffff             addl $0xffffff01, %eax
  1880: 0fb60406                 movzbl (%rsi, %rax), %eax
  1884: 32d8                     xorb %al, %bl
  1886: 498b4010                 movq 0x10(%r8), %rax
  1890: 420fb61418               movzbl (%rax, %r11), %edx
  1895: 4885d2                   testq %rdx, %rdx
  1898: 0f8422000000             je 0x792
  1904: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1908: 81c0c6000000             addl $0xc6, %eax
  1914: 81f8ff000000             cmpl $0xff, %eax
  1920: 0f8206000000             jb 0x78c
  1926: 81c001ffffff             addl $0xffffff01, %eax
  1932: 0fb60406                 movzbl (%rsi, %rax), %eax
  1936: 32d8                     xorb %al, %bl
  1938: 498b4018                 movq 0x18(%r8), %rax
  1942: 420fb61418               movzbl (%rax, %r11), %edx
  1947: 4885d2                   testq %rdx, %rdx
  1950: 0f841f000000             je 0x7c3
  1956: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1960: 83c001                   addl $1, %eax
  1963: 81f8ff000000             cmpl $0xff, %eax
  1969: 0f8206000000             jb 0x7bd
  1975: 81c001ffffff             addl $0xffffff01, %eax
  1981: 0fb60406                 movzbl (%rsi, %rax), %eax
  1985: 32d8                     xorb %al, %bl
  1987: 498b4020                 movq 0x20(%r8), %rax
  1991: 420fb61418               movzbl (%rax, %r11), %edx
  1996: 4885d2                   testq %rdx, %rdx
  1999: 0f8422000000             je 0x7f7
  2005: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2009: 81c0fe000000             addl $0xfe, %eax
  2015: 81f8ff000000             cmpl $0xff, %eax
  2021: 0f8206000000             jb 0x7f1
  2027: 81c001ffffff             addl $0xffffff01, %eax
  2033: 0fb60406                 movzbl (%rsi, %rax), %eax
  2037: 32d8                     xorb %al, %bl
  2039: 498b4028                 movq 0x28(%r8), %rax
  2043: 420fb61418               movzbl (%rax, %r11), %edx
  2048: 4885d2                   testq %rdx, %rdx
  2051: 0f841f000000             je 0x828
  2057: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2061: 83c031                   addl $0x31, %eax
  2064: 81f8ff000000             cmpl $0xff, %eax
  2070: 0f8206000000             jb 0x822
  2076: 81c001ffffff             addl $0xffffff01, %eax
  2082: 0fb60406                 movzbl (%rsi, %rax), %eax
  2086: 32d8                     xorb %al, %bl
  2088: 43881c1a                 movb %bl, (%r10, %r11)
  2092: 33db                     xorl %ebx, %ebx
  2094: 49ffc3                   incq %r11
  2097: 4c3bd9                   cmpq %rcx, %r11
  2100: 0f82c0feffff             jb 0x6fa
  2106: 4d8b5108                 movq 8(%r9), %r10
  2110: 4c8b1c24                 movq (%rsp), %r11
  2114: 33db                     xorl %ebx, %ebx
  2116: 498b00                   movq (%r8), %rax
  2119: 420fb61418               movzbl (%rax, %r11), %edx
  2124: 4885d2                   testq %rdx, %rdx
  2127: 0f841f000000             je 0x874
  2133: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2137: 83c01c                   addl $0x1c, %eax
  2140: 81f8ff000000             cmpl $0xff, %eax
  2146: 0f8206000000             jb 0x86e
  2152: 81c001ffffff             addl $0xffffff01, %eax
  2158: 0fb60406                 movzbl (%rsi, %rax), %eax
  2162: 32d8                     xorb %al, %bl
  2164: 498b4008                 movq 8(%r8), %rax
  2168: 420fb61418               movzbl (%rax, %r11), %edx
  2173: 4885d2                   testq %rdx, %rdx
  2176: 0f8422000000             je 0x8a8
  2182: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2186: 81c0df000000             addl $0xdf, %eax
  2192: 81f8ff000000             cmpl $0xff, %eax
  2198: 0f8206000000             jb 0x8a2
  2204: 81c001ffffff             addl $0xffffff01, %eax
  2210: 0fb60406                 movzbl (%rsi, %rax), %eax
  2214: 32d8                     xorb %al, %bl
  2216: 498b4010                 movq 0x10(%r8), %rax
  2220: 420fb61418               movzbl (%rax, %r11), %edx
  2225: 4885d2                   testq %rdx, %rdx
  2228: 0f841f000000             je 0x8d9
  2234: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2238: 83c01a                   addl $0x1a, %eax
  2241: 81f8ff000000             cmpl $0xff, %eax
  2247: 0f8206000000             jb 0x8d3
  2253: 81c001ffffff             addl $0xffffff01, %eax
  2259: 0fb60406                 movzbl (%rsi, %rax), %eax
  2263: 32d8                     xorb %al, %bl
  2265: 498b4018                 movq 0x18(%r8), %rax
  2269: 420fb61418               movzbl (%rax, %r11), %edx
  2274: 4885d2                   testq %rdx, %rdx
  2277: 0f841f000000             je 0x90a
  2283: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2287: 83c019                   addl $0x19, %eax
  2290: 81f8ff000000             cmpl $0xff, %eax
  2296: 0f8206000000             jb 0x904
  2302: 81c001ffffff             addl $0xffffff01, %eax
  2308: 0fb60406                 movzbl (%rsi, %rax), %eax
  2312: 32d8                     xorb %al, %bl
  2314: 498b4020                 movq 0x20(%r8), %rax
  2318: 420fb61418               movzbl (%rax, %r11), %edx
  2323: 4885d2                   testq %rdx, %rdx
  2326: 0f841f000000             je 0x93b
  2332: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2336: 83c018                   addl $0x18, %eax
  2339: 81f8ff000000             cmpl $0xff, %eax
  2345: 0f8206000000             jb 0x935
  2351: 81c001ffffff             addl $0xffffff01, %eax
  2357: 0fb60406                 movzbl (%rsi, %rax), %eax
  2361: 32d8                     xorb %al, %bl
  2363: 498b4028                 movq 0x28(%r8), %rax
  2367: 420fb61418               movzbl (%rax, %r11), %edx
  2372: 4885d2                   testq %rdx, %rdx
  2375: 0f841f000000             je 0x96c
  2381: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2385: 83c031                   addl $0x31, %eax
  2388: 81f8ff000000             cmpl $0xff, %eax
  2394: 0f8206000000             jb 0x966
  2400: 81c001ffffff             addl $0xffffff01, %eax
  2406: 0fb60406                 movzbl (%rsi, %rax), %eax
  2410: 32d8                     xorb %al, %bl
  2412: 43881c1a                 movb %bl, (%r10, %r11)
  2416: 33db                     xorl %ebx, %ebx
  2418: 49ffc3                   incq %r11
  2421: 4c3bd9                   cmpq %rcx, %r11
  2424: 0f82c6feffff             jb 0x844
  2430: 4883c408                 addq $8, %rsp
  2434: 415e                     popq %r14
  2436: 415d                     popq %r13
  2438: 415c                     popq %r12
  2440: 415f                     popq %r15
  2442: 5b                       popq %rbx
  2443: c5f877                   vzeroupper 
  2446: c3                       retq 
  2447: 0f0b                     ud2 
  2449: 0f0b                     ud2 
  2451: 0f0b                     ud2 
  2453: 0f0b                     ud2 
  2455: 0f0b                     ud2 
  2457: 0f0b                     ud2 
  2459: 0f0b                     ud2 
  2461: 0f0b                     ud2 
  2463: 0f0b                     ud2 
  2465: 0f0b                     ud2 
  2467: 0f0b                     ud2 
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
