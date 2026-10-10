; config_t2_2054 target=avx512-gfni twin=ternlog topo=rs(6,3)/enc op=encode U=5
; code_size=2432 twins=["ternlog"] uops_per_byte=0.73125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d67090000           leaq 0x967(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e840010000           subq $0x140, %rax
    57: 0f8240060000             jb 0x67f
    63: 90                       nop 
    64: 62317d486f3c1e           vmovdqa32 (%rsi, %r11), %zmm15
    71: 62d17d486fc7             vmovdqa32 %zmm15, %zmm0
    77: 62c17d486f0f             vmovdqa32 (%r15), %zmm17
    83: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
    90: 62b17d486fe8             vmovdqa32 %zmm16, %zmm5
    96: 62c17d486f4f01           vmovdqa32 0x40(%r15), %zmm17
   103: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   110: 62317d486fd0             vmovdqa32 %zmm16, %zmm10
   116: 62317d486f7c1e01         vmovdqa32 0x40(%rsi, %r11), %zmm15
   124: 62d17d486fcf             vmovdqa32 %zmm15, %zmm1
   130: 62c17d486f0f             vmovdqa32 (%r15), %zmm17
   136: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   143: 62b17d486ff0             vmovdqa32 %zmm16, %zmm6
   149: 62c17d486f4f01           vmovdqa32 0x40(%r15), %zmm17
   156: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   163: 62317d486fd8             vmovdqa32 %zmm16, %zmm11
   169: 62317d486f7c1e02         vmovdqa32 0x80(%rsi, %r11), %zmm15
   177: 62d17d486fd7             vmovdqa32 %zmm15, %zmm2
   183: 62c17d486f0f             vmovdqa32 (%r15), %zmm17
   189: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   196: 62b17d486ff8             vmovdqa32 %zmm16, %zmm7
   202: 62c17d486f4f01           vmovdqa32 0x40(%r15), %zmm17
   209: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   216: 62317d486fe0             vmovdqa32 %zmm16, %zmm12
   222: 62317d486f7c1e03         vmovdqa32 0xc0(%rsi, %r11), %zmm15
   230: 62d17d486fdf             vmovdqa32 %zmm15, %zmm3
   236: 62c17d486f0f             vmovdqa32 (%r15), %zmm17
   242: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   249: 62317d486fc0             vmovdqa32 %zmm16, %zmm8
   255: 62c17d486f4f01           vmovdqa32 0x40(%r15), %zmm17
   262: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   269: 62317d486fe8             vmovdqa32 %zmm16, %zmm13
   275: 62317d486f7c1e04         vmovdqa32 0x100(%rsi, %r11), %zmm15
   283: 62d17d486fe7             vmovdqa32 %zmm15, %zmm4
   289: 62c17d486f0f             vmovdqa32 (%r15), %zmm17
   295: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   302: 62317d486fc8             vmovdqa32 %zmm16, %zmm9
   308: 62c17d486f4f01           vmovdqa32 0x40(%r15), %zmm17
   315: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   322: 62317d486ff0             vmovdqa32 %zmm16, %zmm14
   328: 62317d486f3c1f           vmovdqa32 (%rdi, %r11), %zmm15
   335: 62d17d48efc7             vpxord %zmm15, %zmm0, %zmm0
   341: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   348: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
   354: 62c17d486f4f02           vmovdqa32 0x80(%r15), %zmm17
   361: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   368: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
   374: 62317d486f7c1f01         vmovdqa32 0x40(%rdi, %r11), %zmm15
   382: 62d17548efcf             vpxord %zmm15, %zmm1, %zmm1
   388: 62c17d486f4f01           vmovdqa32 0x40(%r15), %zmm17
   395: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   402: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
   408: 62c17d486f4f02           vmovdqa32 0x80(%r15), %zmm17
   415: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   422: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
   428: 62317d486f7c1f02         vmovdqa32 0x80(%rdi, %r11), %zmm15
   436: 62d16d48efd7             vpxord %zmm15, %zmm2, %zmm2
   442: 62c17d486f4f01           vmovdqa32 0x40(%r15), %zmm17
   449: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   456: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
   462: 62c17d486f4f02           vmovdqa32 0x80(%r15), %zmm17
   469: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   476: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
   482: 62317d486f7c1f03         vmovdqa32 0xc0(%rdi, %r11), %zmm15
   490: 62d16548efdf             vpxord %zmm15, %zmm3, %zmm3
   496: 62c17d486f4f01           vmovdqa32 0x40(%r15), %zmm17
   503: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   510: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
   516: 62c17d486f4f02           vmovdqa32 0x80(%r15), %zmm17
   523: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   530: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
   536: 62317d486f7c1f04         vmovdqa32 0x100(%rdi, %r11), %zmm15
   544: 62d15d48efe7             vpxord %zmm15, %zmm4, %zmm4
   550: 62c17d486f4f01           vmovdqa32 0x40(%r15), %zmm17
   557: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   564: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
   570: 62c17d486f4f02           vmovdqa32 0x80(%r15), %zmm17
   577: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   584: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
   590: 62117d486f3c1c           vmovdqa32 (%r12, %r11), %zmm15
   597: 62d17d48efc7             vpxord %zmm15, %zmm0, %zmm0
   603: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   610: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
   616: 62c17d486f4f03           vmovdqa32 0xc0(%r15), %zmm17
   623: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   630: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
   636: 62117d486f7c1c01         vmovdqa32 0x40(%r12, %r11), %zmm15
   644: 62d17548efcf             vpxord %zmm15, %zmm1, %zmm1
   650: 62c17d486f4f02           vmovdqa32 0x80(%r15), %zmm17
   657: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   664: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
   670: 62c17d486f4f03           vmovdqa32 0xc0(%r15), %zmm17
   677: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   684: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
   690: 62117d486f7c1c02         vmovdqa32 0x80(%r12, %r11), %zmm15
   698: 62d16d48efd7             vpxord %zmm15, %zmm2, %zmm2
   704: 62c17d486f4f02           vmovdqa32 0x80(%r15), %zmm17
   711: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   718: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
   724: 62c17d486f4f03           vmovdqa32 0xc0(%r15), %zmm17
   731: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   738: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
   744: 62117d486f7c1c03         vmovdqa32 0xc0(%r12, %r11), %zmm15
   752: 62d16548efdf             vpxord %zmm15, %zmm3, %zmm3
   758: 62c17d486f4f02           vmovdqa32 0x80(%r15), %zmm17
   765: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   772: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
   778: 62c17d486f4f03           vmovdqa32 0xc0(%r15), %zmm17
   785: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   792: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
   798: 62117d486f7c1c04         vmovdqa32 0x100(%r12, %r11), %zmm15
   806: 62d15d48efe7             vpxord %zmm15, %zmm4, %zmm4
   812: 62c17d486f4f02           vmovdqa32 0x80(%r15), %zmm17
   819: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   826: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
   832: 62c17d486f4f03           vmovdqa32 0xc0(%r15), %zmm17
   839: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   846: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
   852: 62117d486f7c1d00         vmovdqa32 (%r13, %r11), %zmm15
   860: 62d17d48efc7             vpxord %zmm15, %zmm0, %zmm0
   866: 62c17d486f4f04           vmovdqa32 0x100(%r15), %zmm17
   873: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   880: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
   886: 62c17d486f0f             vmovdqa32 (%r15), %zmm17
   892: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   899: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
   905: 62117d486f7c1d01         vmovdqa32 0x40(%r13, %r11), %zmm15
   913: 62d17548efcf             vpxord %zmm15, %zmm1, %zmm1
   919: 62c17d486f4f04           vmovdqa32 0x100(%r15), %zmm17
   926: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   933: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
   939: 62c17d486f0f             vmovdqa32 (%r15), %zmm17
   945: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   952: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
   958: 62117d486f7c1d02         vmovdqa32 0x80(%r13, %r11), %zmm15
   966: 62d16d48efd7             vpxord %zmm15, %zmm2, %zmm2
   972: 62c17d486f4f04           vmovdqa32 0x100(%r15), %zmm17
   979: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
   986: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
   992: 62c17d486f0f             vmovdqa32 (%r15), %zmm17
   998: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1005: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
  1011: 62117d486f7c1d03         vmovdqa32 0xc0(%r13, %r11), %zmm15
  1019: 62d16548efdf             vpxord %zmm15, %zmm3, %zmm3
  1025: 62c17d486f4f04           vmovdqa32 0x100(%r15), %zmm17
  1032: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1039: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
  1045: 62c17d486f0f             vmovdqa32 (%r15), %zmm17
  1051: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1058: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
  1064: 62117d486f7c1d04         vmovdqa32 0x100(%r13, %r11), %zmm15
  1072: 62d15d48efe7             vpxord %zmm15, %zmm4, %zmm4
  1078: 62c17d486f4f04           vmovdqa32 0x100(%r15), %zmm17
  1085: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1092: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
  1098: 62c17d486f0f             vmovdqa32 (%r15), %zmm17
  1104: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1111: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
  1117: 62117d486f3c1e           vmovdqa32 (%r14, %r11), %zmm15
  1124: 62d17d48efc7             vpxord %zmm15, %zmm0, %zmm0
  1130: 62c17d486f4f03           vmovdqa32 0xc0(%r15), %zmm17
  1137: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1144: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
  1150: 62c17d486f4f04           vmovdqa32 0x100(%r15), %zmm17
  1157: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1164: 62312d48efd0             vpxord %zmm16, %zmm10, %zmm10
  1170: 62117d486f7c1e01         vmovdqa32 0x40(%r14, %r11), %zmm15
  1178: 62d17548efcf             vpxord %zmm15, %zmm1, %zmm1
  1184: 62c17d486f4f03           vmovdqa32 0xc0(%r15), %zmm17
  1191: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1198: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
  1204: 62c17d486f4f04           vmovdqa32 0x100(%r15), %zmm17
  1211: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1218: 62312548efd8             vpxord %zmm16, %zmm11, %zmm11
  1224: 62117d486f7c1e02         vmovdqa32 0x80(%r14, %r11), %zmm15
  1232: 62d16d48efd7             vpxord %zmm15, %zmm2, %zmm2
  1238: 62c17d486f4f03           vmovdqa32 0xc0(%r15), %zmm17
  1245: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1252: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
  1258: 62c17d486f4f04           vmovdqa32 0x100(%r15), %zmm17
  1265: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1272: 62311d48efe0             vpxord %zmm16, %zmm12, %zmm12
  1278: 62117d486f7c1e03         vmovdqa32 0xc0(%r14, %r11), %zmm15
  1286: 62d16548efdf             vpxord %zmm15, %zmm3, %zmm3
  1292: 62c17d486f4f03           vmovdqa32 0xc0(%r15), %zmm17
  1299: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1306: 62313d48efc0             vpxord %zmm16, %zmm8, %zmm8
  1312: 62c17d486f4f04           vmovdqa32 0x100(%r15), %zmm17
  1319: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1326: 62311548efe8             vpxord %zmm16, %zmm13, %zmm13
  1332: 62117d486f7c1e04         vmovdqa32 0x100(%r14, %r11), %zmm15
  1340: 62d15d48efe7             vpxord %zmm15, %zmm4, %zmm4
  1346: 62c17d486f4f03           vmovdqa32 0xc0(%r15), %zmm17
  1353: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1360: 62313548efc8             vpxord %zmm16, %zmm9, %zmm9
  1366: 62c17d486f4f04           vmovdqa32 0x100(%r15), %zmm17
  1373: 62a38548cec100           vgf2p8affineqb $0, %zmm17, %zmm15, %zmm16
  1380: 62310d48eff0             vpxord %zmm16, %zmm14, %zmm14
  1386: 4d8b5028                 movq 0x28(%r8), %r10
  1390: 62117d486f3c1a           vmovdqa32 (%r10, %r11), %zmm15
  1397: 62d17d48efc7             vpxord %zmm15, %zmm0, %zmm0
  1403: 62d15548efef             vpxord %zmm15, %zmm5, %zmm5
  1409: 62512d48efd7             vpxord %zmm15, %zmm10, %zmm10
  1415: 62117d486f7c1a01         vmovdqa32 0x40(%r10, %r11), %zmm15
  1423: 62d17548efcf             vpxord %zmm15, %zmm1, %zmm1
  1429: 62d14d48eff7             vpxord %zmm15, %zmm6, %zmm6
  1435: 62512548efdf             vpxord %zmm15, %zmm11, %zmm11
  1441: 62117d486f7c1a02         vmovdqa32 0x80(%r10, %r11), %zmm15
  1449: 62d16d48efd7             vpxord %zmm15, %zmm2, %zmm2
  1455: 62d14548efff             vpxord %zmm15, %zmm7, %zmm7
  1461: 62511d48efe7             vpxord %zmm15, %zmm12, %zmm12
  1467: 62117d486f7c1a03         vmovdqa32 0xc0(%r10, %r11), %zmm15
  1475: 62d16548efdf             vpxord %zmm15, %zmm3, %zmm3
  1481: 62513d48efc7             vpxord %zmm15, %zmm8, %zmm8
  1487: 62511548efef             vpxord %zmm15, %zmm13, %zmm13
  1493: 62117d486f7c1a04         vmovdqa32 0x100(%r10, %r11), %zmm15
  1501: 62d15d48efe7             vpxord %zmm15, %zmm4, %zmm4
  1507: 62513548efcf             vpxord %zmm15, %zmm9, %zmm9
  1513: 62510d48eff7             vpxord %zmm15, %zmm14, %zmm14
  1519: 4d8b11                   movq (%r9), %r10
  1522: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  1529: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  1537: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
  1545: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
  1553: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
  1561: 4d8b5108                 movq 8(%r9), %r10
  1565: 62917d487f2c1a           vmovdqa32 %zmm5, (%r10, %r11)
  1572: 62917d487f741a01         vmovdqa32 %zmm6, 0x40(%r10, %r11)
  1580: 62917d487f7c1a02         vmovdqa32 %zmm7, 0x80(%r10, %r11)
  1588: 62117d487f441a03         vmovdqa32 %zmm8, 0xc0(%r10, %r11)
  1596: 62117d487f4c1a04         vmovdqa32 %zmm9, 0x100(%r10, %r11)
  1604: 4d8b5110                 movq 0x10(%r9), %r10
  1608: 62117d487f141a           vmovdqa32 %zmm10, (%r10, %r11)
  1615: 62117d487f5c1a01         vmovdqa32 %zmm11, 0x40(%r10, %r11)
  1623: 62117d487f641a02         vmovdqa32 %zmm12, 0x80(%r10, %r11)
  1631: 62117d487f6c1a03         vmovdqa32 %zmm13, 0xc0(%r10, %r11)
  1639: 62117d487f741a04         vmovdqa32 %zmm14, 0x100(%r10, %r11)
  1647: 4981c340010000           addq $0x140, %r11
  1654: 4c3bd8                   cmpq %rax, %r11
  1657: 0f86c1f9ffff             jbe 0x40
  1663: 4c3bd9                   cmpq %rcx, %r11
  1666: 0f83b2020000             jae 0x93a
  1672: 4153                     pushq %r11
  1674: 498dbf40010000           leaq 0x140(%r15), %rdi
  1681: 498db740020000           leaq 0x240(%r15), %rsi
  1688: 4d8b11                   movq (%r9), %r10
  1691: 4c8b1c24                 movq (%rsp), %r11
  1695: 33db                     xorl %ebx, %ebx
  1697: 498b00                   movq (%r8), %rax
  1700: 420fb61418               movzbl (%rax, %r11), %edx
  1705: 32da                     xorb %dl, %bl
  1707: 498b4008                 movq 8(%r8), %rax
  1711: 420fb61418               movzbl (%rax, %r11), %edx
  1716: 32da                     xorb %dl, %bl
  1718: 498b4010                 movq 0x10(%r8), %rax
  1722: 420fb61418               movzbl (%rax, %r11), %edx
  1727: 32da                     xorb %dl, %bl
  1729: 498b4018                 movq 0x18(%r8), %rax
  1733: 420fb61418               movzbl (%rax, %r11), %edx
  1738: 32da                     xorb %dl, %bl
  1740: 498b4020                 movq 0x20(%r8), %rax
  1744: 420fb61418               movzbl (%rax, %r11), %edx
  1749: 32da                     xorb %dl, %bl
  1751: 498b4028                 movq 0x28(%r8), %rax
  1755: 420fb61418               movzbl (%rax, %r11), %edx
  1760: 32da                     xorb %dl, %bl
  1762: 43881c1a                 movb %bl, (%r10, %r11)
  1766: 33db                     xorl %ebx, %ebx
  1768: 49ffc3                   incq %r11
  1771: 4c3bd9                   cmpq %rcx, %r11
  1774: 0f82adffffff             jb 0x6a1
  1780: 4d8b5108                 movq 8(%r9), %r10
  1784: 4c8b1c24                 movq (%rsp), %r11
  1788: 33db                     xorl %ebx, %ebx
  1790: 498b00                   movq (%r8), %rax
  1793: 420fb61418               movzbl (%rax, %r11), %edx
  1798: 4885d2                   testq %rdx, %rdx
  1801: 0f8422000000             je 0x731
  1807: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1811: 81c0e6000000             addl $0xe6, %eax
  1817: 81f8ff000000             cmpl $0xff, %eax
  1823: 0f8206000000             jb 0x72b
  1829: 81c001ffffff             addl $0xffffff01, %eax
  1835: 0fb60406                 movzbl (%rsi, %rax), %eax
  1839: 32d8                     xorb %al, %bl
  1841: 498b4008                 movq 8(%r8), %rax
  1845: 420fb61418               movzbl (%rax, %r11), %edx
  1850: 4885d2                   testq %rdx, %rdx
  1853: 0f8422000000             je 0x765
  1859: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1863: 81c0cd000000             addl $0xcd, %eax
  1869: 81f8ff000000             cmpl $0xff, %eax
  1875: 0f8206000000             jb 0x75f
  1881: 81c001ffffff             addl $0xffffff01, %eax
  1887: 0fb60406                 movzbl (%rsi, %rax), %eax
  1891: 32d8                     xorb %al, %bl
  1893: 498b4010                 movq 0x10(%r8), %rax
  1897: 420fb61418               movzbl (%rax, %r11), %edx
  1902: 4885d2                   testq %rdx, %rdx
  1905: 0f841f000000             je 0x796
  1911: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1915: 83c04b                   addl $0x4b, %eax
  1918: 81f8ff000000             cmpl $0xff, %eax
  1924: 0f8206000000             jb 0x790
  1930: 81c001ffffff             addl $0xffffff01, %eax
  1936: 0fb60406                 movzbl (%rsi, %rax), %eax
  1940: 32d8                     xorb %al, %bl
  1942: 498b4018                 movq 0x18(%r8), %rax
  1946: 420fb61418               movzbl (%rax, %r11), %edx
  1951: 4885d2                   testq %rdx, %rdx
  1954: 0f841f000000             je 0x7c7
  1960: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1964: 83c032                   addl $0x32, %eax
  1967: 81f8ff000000             cmpl $0xff, %eax
  1973: 0f8206000000             jb 0x7c1
  1979: 81c001ffffff             addl $0xffffff01, %eax
  1985: 0fb60406                 movzbl (%rsi, %rax), %eax
  1989: 32d8                     xorb %al, %bl
  1991: 498b4020                 movq 0x20(%r8), %rax
  1995: 420fb61418               movzbl (%rax, %r11), %edx
  2000: 4885d2                   testq %rdx, %rdx
  2003: 0f841f000000             je 0x7f8
  2009: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2013: 83c019                   addl $0x19, %eax
  2016: 81f8ff000000             cmpl $0xff, %eax
  2022: 0f8206000000             jb 0x7f2
  2028: 81c001ffffff             addl $0xffffff01, %eax
  2034: 0fb60406                 movzbl (%rsi, %rax), %eax
  2038: 32d8                     xorb %al, %bl
  2040: 498b4028                 movq 0x28(%r8), %rax
  2044: 420fb61418               movzbl (%rax, %r11), %edx
  2049: 32da                     xorb %dl, %bl
  2051: 43881c1a                 movb %bl, (%r10, %r11)
  2055: 33db                     xorl %ebx, %ebx
  2057: 49ffc3                   incq %r11
  2060: 4c3bd9                   cmpq %rcx, %r11
  2063: 0f82e9feffff             jb 0x6fe
  2069: 4d8b5110                 movq 0x10(%r9), %r10
  2073: 4c8b1c24                 movq (%rsp), %r11
  2077: 33db                     xorl %ebx, %ebx
  2079: 498b00                   movq (%r8), %rax
  2082: 420fb61418               movzbl (%rax, %r11), %edx
  2087: 4885d2                   testq %rdx, %rdx
  2090: 0f8422000000             je 0x852
  2096: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2100: 81c0cd000000             addl $0xcd, %eax
  2106: 81f8ff000000             cmpl $0xff, %eax
  2112: 0f8206000000             jb 0x84c
  2118: 81c001ffffff             addl $0xffffff01, %eax
  2124: 0fb60406                 movzbl (%rsi, %rax), %eax
  2128: 32d8                     xorb %al, %bl
  2130: 498b4008                 movq 8(%r8), %rax
  2134: 420fb61418               movzbl (%rax, %r11), %edx
  2139: 4885d2                   testq %rdx, %rdx
  2142: 0f841f000000             je 0x883
  2148: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2152: 83c04b                   addl $0x4b, %eax
  2155: 81f8ff000000             cmpl $0xff, %eax
  2161: 0f8206000000             jb 0x87d
  2167: 81c001ffffff             addl $0xffffff01, %eax
  2173: 0fb60406                 movzbl (%rsi, %rax), %eax
  2177: 32d8                     xorb %al, %bl
  2179: 498b4010                 movq 0x10(%r8), %rax
  2183: 420fb61418               movzbl (%rax, %r11), %edx
  2188: 4885d2                   testq %rdx, %rdx
  2191: 0f841f000000             je 0x8b4
  2197: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2201: 83c019                   addl $0x19, %eax
  2204: 81f8ff000000             cmpl $0xff, %eax
  2210: 0f8206000000             jb 0x8ae
  2216: 81c001ffffff             addl $0xffffff01, %eax
  2222: 0fb60406                 movzbl (%rsi, %rax), %eax
  2226: 32d8                     xorb %al, %bl
  2228: 498b4018                 movq 0x18(%r8), %rax
  2232: 420fb61418               movzbl (%rax, %r11), %edx
  2237: 4885d2                   testq %rdx, %rdx
  2240: 0f8422000000             je 0x8e8
  2246: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2250: 81c0e6000000             addl $0xe6, %eax
  2256: 81f8ff000000             cmpl $0xff, %eax
  2262: 0f8206000000             jb 0x8e2
  2268: 81c001ffffff             addl $0xffffff01, %eax
  2274: 0fb60406                 movzbl (%rsi, %rax), %eax
  2278: 32d8                     xorb %al, %bl
  2280: 498b4020                 movq 0x20(%r8), %rax
  2284: 420fb61418               movzbl (%rax, %r11), %edx
  2289: 4885d2                   testq %rdx, %rdx
  2292: 0f841f000000             je 0x919
  2298: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2302: 83c032                   addl $0x32, %eax
  2305: 81f8ff000000             cmpl $0xff, %eax
  2311: 0f8206000000             jb 0x913
  2317: 81c001ffffff             addl $0xffffff01, %eax
  2323: 0fb60406                 movzbl (%rsi, %rax), %eax
  2327: 32d8                     xorb %al, %bl
  2329: 498b4028                 movq 0x28(%r8), %rax
  2333: 420fb61418               movzbl (%rax, %r11), %edx
  2338: 32da                     xorb %dl, %bl
  2340: 43881c1a                 movb %bl, (%r10, %r11)
  2344: 33db                     xorl %ebx, %ebx
  2346: 49ffc3                   incq %r11
  2349: 4c3bd9                   cmpq %rcx, %r11
  2352: 0f82e9feffff             jb 0x81f
  2358: 4883c408                 addq $8, %rsp
  2362: 415e                     popq %r14
  2364: 415d                     popq %r13
  2366: 415c                     popq %r12
  2368: 415f                     popq %r15
  2370: 5b                       popq %rbx
  2371: c5f877                   vzeroupper 
  2374: c3                       retq 
  2375: 0f0b                     ud2 
  2377: 0f0b                     ud2 
  2379: 0f0b                     ud2 
  2381: 0f0b                     ud2 
  2383: 0f0b                     ud2 
  2385: 0f0b                     ud2 
  2387: 0f0b                     ud2 
  2389: 0f0b                     ud2 
  2391: 0f0b                     ud2 
  2393: 0f0b                     ud2 
  2395: 0f0b                     ud2 
  2397: 0f0b                     ud2 
  2399: 0f0b                     ud2 
  2401: 0f0b                     ud2 
  2403: 0f0b                     ud2 
  2405: 0f0b                     ud2 
  2407: 0f0b                     ud2 
  2409: 0f0b                     ud2 
  2411: 0f0b                     ud2 
  2413: 0f0b                     ud2 
  2415: 0f0b                     ud2 
  2417: 0f0b                     ud2 
  2419: 0f0b                     ud2 
  2421: 0f0b                     ud2 
  2423: 0f0b                     ud2 
  2425: 0f0b                     ud2 
  2427: 0f0b                     ud2 
  2429: 0f0b                     ud2 
  2431: cc                       int3 
