; config_t2_2055 target=avx512-gfni twin=ternlog topo=rs(6,3)/enc op=encode U=6
; code_size=2752 twins=["ternlog"] uops_per_byte=0.7317708333333334
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da70a0000           leaq 0xaa7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e880010000           subq $0x180, %rax
    57: 0f827e070000             jb 0x7bd
    63: 90                       nop 
    64: 62a17d486f141e           vmovdqa32 (%rsi, %r11), %zmm18
    71: 62b17d486fc2             vmovdqa32 %zmm18, %zmm0
    77: 62c17d486f27             vmovdqa32 (%r15), %zmm20
    83: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
    90: 62b17d486ff3             vmovdqa32 %zmm19, %zmm6
    96: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   103: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   110: 62317d486fe3             vmovdqa32 %zmm19, %zmm12
   116: 62a17d486f541e01         vmovdqa32 0x40(%rsi, %r11), %zmm18
   124: 62b17d486fca             vmovdqa32 %zmm18, %zmm1
   130: 62c17d486f27             vmovdqa32 (%r15), %zmm20
   136: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   143: 62b17d486ffb             vmovdqa32 %zmm19, %zmm7
   149: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   156: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   163: 62317d486feb             vmovdqa32 %zmm19, %zmm13
   169: 62a17d486f541e02         vmovdqa32 0x80(%rsi, %r11), %zmm18
   177: 62b17d486fd2             vmovdqa32 %zmm18, %zmm2
   183: 62c17d486f27             vmovdqa32 (%r15), %zmm20
   189: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   196: 62317d486fc3             vmovdqa32 %zmm19, %zmm8
   202: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   209: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   216: 62317d486ff3             vmovdqa32 %zmm19, %zmm14
   222: 62a17d486f541e03         vmovdqa32 0xc0(%rsi, %r11), %zmm18
   230: 62b17d486fda             vmovdqa32 %zmm18, %zmm3
   236: 62c17d486f27             vmovdqa32 (%r15), %zmm20
   242: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   249: 62317d486fcb             vmovdqa32 %zmm19, %zmm9
   255: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   262: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   269: 62317d486ffb             vmovdqa32 %zmm19, %zmm15
   275: 62a17d486f541e04         vmovdqa32 0x100(%rsi, %r11), %zmm18
   283: 62b17d486fe2             vmovdqa32 %zmm18, %zmm4
   289: 62c17d486f27             vmovdqa32 (%r15), %zmm20
   295: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   302: 62317d486fd3             vmovdqa32 %zmm19, %zmm10
   308: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   315: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   322: 62a17d486fc3             vmovdqa32 %zmm19, %zmm16
   328: 62a17d486f541e05         vmovdqa32 0x140(%rsi, %r11), %zmm18
   336: 62b17d486fea             vmovdqa32 %zmm18, %zmm5
   342: 62c17d486f27             vmovdqa32 (%r15), %zmm20
   348: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   355: 62317d486fdb             vmovdqa32 %zmm19, %zmm11
   361: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   368: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   375: 62a17d486fcb             vmovdqa32 %zmm19, %zmm17
   381: 62a17d486f141f           vmovdqa32 (%rdi, %r11), %zmm18
   388: 62b17d48efc2             vpxord %zmm18, %zmm0, %zmm0
   394: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   401: 62b14d48eff3             vpxord %zmm19, %zmm6, %zmm6
   407: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   414: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   421: 62311d48efe3             vpxord %zmm19, %zmm12, %zmm12
   427: 62a17d486f541f01         vmovdqa32 0x40(%rdi, %r11), %zmm18
   435: 62b17548efca             vpxord %zmm18, %zmm1, %zmm1
   441: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   448: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   455: 62b14548effb             vpxord %zmm19, %zmm7, %zmm7
   461: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   468: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   475: 62311548efeb             vpxord %zmm19, %zmm13, %zmm13
   481: 62a17d486f541f02         vmovdqa32 0x80(%rdi, %r11), %zmm18
   489: 62b16d48efd2             vpxord %zmm18, %zmm2, %zmm2
   495: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   502: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   509: 62313d48efc3             vpxord %zmm19, %zmm8, %zmm8
   515: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   522: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   529: 62310d48eff3             vpxord %zmm19, %zmm14, %zmm14
   535: 62a17d486f541f03         vmovdqa32 0xc0(%rdi, %r11), %zmm18
   543: 62b16548efda             vpxord %zmm18, %zmm3, %zmm3
   549: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   556: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   563: 62313548efcb             vpxord %zmm19, %zmm9, %zmm9
   569: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   576: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   583: 62310548effb             vpxord %zmm19, %zmm15, %zmm15
   589: 62a17d486f541f04         vmovdqa32 0x100(%rdi, %r11), %zmm18
   597: 62b15d48efe2             vpxord %zmm18, %zmm4, %zmm4
   603: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   610: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   617: 62312d48efd3             vpxord %zmm19, %zmm10, %zmm10
   623: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   630: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   637: 62a17d40efc3             vpxord %zmm19, %zmm16, %zmm16
   643: 62a17d486f541f05         vmovdqa32 0x140(%rdi, %r11), %zmm18
   651: 62b15548efea             vpxord %zmm18, %zmm5, %zmm5
   657: 62c17d486f6701           vmovdqa32 0x40(%r15), %zmm20
   664: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   671: 62312548efdb             vpxord %zmm19, %zmm11, %zmm11
   677: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   684: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   691: 62a17540efcb             vpxord %zmm19, %zmm17, %zmm17
   697: 62817d486f141c           vmovdqa32 (%r12, %r11), %zmm18
   704: 62b17d48efc2             vpxord %zmm18, %zmm0, %zmm0
   710: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   717: 62b14d48eff3             vpxord %zmm19, %zmm6, %zmm6
   723: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
   730: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   737: 62311d48efe3             vpxord %zmm19, %zmm12, %zmm12
   743: 62817d486f541c01         vmovdqa32 0x40(%r12, %r11), %zmm18
   751: 62b17548efca             vpxord %zmm18, %zmm1, %zmm1
   757: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   764: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   771: 62b14548effb             vpxord %zmm19, %zmm7, %zmm7
   777: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
   784: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   791: 62311548efeb             vpxord %zmm19, %zmm13, %zmm13
   797: 62817d486f541c02         vmovdqa32 0x80(%r12, %r11), %zmm18
   805: 62b16d48efd2             vpxord %zmm18, %zmm2, %zmm2
   811: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   818: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   825: 62313d48efc3             vpxord %zmm19, %zmm8, %zmm8
   831: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
   838: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   845: 62310d48eff3             vpxord %zmm19, %zmm14, %zmm14
   851: 62817d486f541c03         vmovdqa32 0xc0(%r12, %r11), %zmm18
   859: 62b16548efda             vpxord %zmm18, %zmm3, %zmm3
   865: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   872: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   879: 62313548efcb             vpxord %zmm19, %zmm9, %zmm9
   885: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
   892: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   899: 62310548effb             vpxord %zmm19, %zmm15, %zmm15
   905: 62817d486f541c04         vmovdqa32 0x100(%r12, %r11), %zmm18
   913: 62b15d48efe2             vpxord %zmm18, %zmm4, %zmm4
   919: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   926: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   933: 62312d48efd3             vpxord %zmm19, %zmm10, %zmm10
   939: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
   946: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   953: 62a17d40efc3             vpxord %zmm19, %zmm16, %zmm16
   959: 62817d486f541c05         vmovdqa32 0x140(%r12, %r11), %zmm18
   967: 62b15548efea             vpxord %zmm18, %zmm5, %zmm5
   973: 62c17d486f6702           vmovdqa32 0x80(%r15), %zmm20
   980: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
   987: 62312548efdb             vpxord %zmm19, %zmm11, %zmm11
   993: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
  1000: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1007: 62a17540efcb             vpxord %zmm19, %zmm17, %zmm17
  1013: 62817d486f541d00         vmovdqa32 (%r13, %r11), %zmm18
  1021: 62b17d48efc2             vpxord %zmm18, %zmm0, %zmm0
  1027: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1034: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1041: 62b14d48eff3             vpxord %zmm19, %zmm6, %zmm6
  1047: 62c17d486f27             vmovdqa32 (%r15), %zmm20
  1053: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1060: 62311d48efe3             vpxord %zmm19, %zmm12, %zmm12
  1066: 62817d486f541d01         vmovdqa32 0x40(%r13, %r11), %zmm18
  1074: 62b17548efca             vpxord %zmm18, %zmm1, %zmm1
  1080: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1087: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1094: 62b14548effb             vpxord %zmm19, %zmm7, %zmm7
  1100: 62c17d486f27             vmovdqa32 (%r15), %zmm20
  1106: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1113: 62311548efeb             vpxord %zmm19, %zmm13, %zmm13
  1119: 62817d486f541d02         vmovdqa32 0x80(%r13, %r11), %zmm18
  1127: 62b16d48efd2             vpxord %zmm18, %zmm2, %zmm2
  1133: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1140: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1147: 62313d48efc3             vpxord %zmm19, %zmm8, %zmm8
  1153: 62c17d486f27             vmovdqa32 (%r15), %zmm20
  1159: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1166: 62310d48eff3             vpxord %zmm19, %zmm14, %zmm14
  1172: 62817d486f541d03         vmovdqa32 0xc0(%r13, %r11), %zmm18
  1180: 62b16548efda             vpxord %zmm18, %zmm3, %zmm3
  1186: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1193: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1200: 62313548efcb             vpxord %zmm19, %zmm9, %zmm9
  1206: 62c17d486f27             vmovdqa32 (%r15), %zmm20
  1212: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1219: 62310548effb             vpxord %zmm19, %zmm15, %zmm15
  1225: 62817d486f541d04         vmovdqa32 0x100(%r13, %r11), %zmm18
  1233: 62b15d48efe2             vpxord %zmm18, %zmm4, %zmm4
  1239: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1246: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1253: 62312d48efd3             vpxord %zmm19, %zmm10, %zmm10
  1259: 62c17d486f27             vmovdqa32 (%r15), %zmm20
  1265: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1272: 62a17d40efc3             vpxord %zmm19, %zmm16, %zmm16
  1278: 62817d486f541d05         vmovdqa32 0x140(%r13, %r11), %zmm18
  1286: 62b15548efea             vpxord %zmm18, %zmm5, %zmm5
  1292: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1299: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1306: 62312548efdb             vpxord %zmm19, %zmm11, %zmm11
  1312: 62c17d486f27             vmovdqa32 (%r15), %zmm20
  1318: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1325: 62a17540efcb             vpxord %zmm19, %zmm17, %zmm17
  1331: 62817d486f141e           vmovdqa32 (%r14, %r11), %zmm18
  1338: 62b17d48efc2             vpxord %zmm18, %zmm0, %zmm0
  1344: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
  1351: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1358: 62b14d48eff3             vpxord %zmm19, %zmm6, %zmm6
  1364: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1371: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1378: 62311d48efe3             vpxord %zmm19, %zmm12, %zmm12
  1384: 62817d486f541e01         vmovdqa32 0x40(%r14, %r11), %zmm18
  1392: 62b17548efca             vpxord %zmm18, %zmm1, %zmm1
  1398: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
  1405: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1412: 62b14548effb             vpxord %zmm19, %zmm7, %zmm7
  1418: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1425: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1432: 62311548efeb             vpxord %zmm19, %zmm13, %zmm13
  1438: 62817d486f541e02         vmovdqa32 0x80(%r14, %r11), %zmm18
  1446: 62b16d48efd2             vpxord %zmm18, %zmm2, %zmm2
  1452: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
  1459: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1466: 62313d48efc3             vpxord %zmm19, %zmm8, %zmm8
  1472: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1479: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1486: 62310d48eff3             vpxord %zmm19, %zmm14, %zmm14
  1492: 62817d486f541e03         vmovdqa32 0xc0(%r14, %r11), %zmm18
  1500: 62b16548efda             vpxord %zmm18, %zmm3, %zmm3
  1506: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
  1513: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1520: 62313548efcb             vpxord %zmm19, %zmm9, %zmm9
  1526: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1533: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1540: 62310548effb             vpxord %zmm19, %zmm15, %zmm15
  1546: 62817d486f541e04         vmovdqa32 0x100(%r14, %r11), %zmm18
  1554: 62b15d48efe2             vpxord %zmm18, %zmm4, %zmm4
  1560: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
  1567: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1574: 62312d48efd3             vpxord %zmm19, %zmm10, %zmm10
  1580: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1587: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1594: 62a17d40efc3             vpxord %zmm19, %zmm16, %zmm16
  1600: 62817d486f541e05         vmovdqa32 0x140(%r14, %r11), %zmm18
  1608: 62b15548efea             vpxord %zmm18, %zmm5, %zmm5
  1614: 62c17d486f6703           vmovdqa32 0xc0(%r15), %zmm20
  1621: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1628: 62312548efdb             vpxord %zmm19, %zmm11, %zmm11
  1634: 62c17d486f6704           vmovdqa32 0x100(%r15), %zmm20
  1641: 62a3ed40cedc00           vgf2p8affineqb $0, %zmm20, %zmm18, %zmm19
  1648: 62a17540efcb             vpxord %zmm19, %zmm17, %zmm17
  1654: 4d8b5028                 movq 0x28(%r8), %r10
  1658: 62817d486f141a           vmovdqa32 (%r10, %r11), %zmm18
  1665: 62b17d48efc2             vpxord %zmm18, %zmm0, %zmm0
  1671: 62b14d48eff2             vpxord %zmm18, %zmm6, %zmm6
  1677: 62311d48efe2             vpxord %zmm18, %zmm12, %zmm12
  1683: 62817d486f541a01         vmovdqa32 0x40(%r10, %r11), %zmm18
  1691: 62b17548efca             vpxord %zmm18, %zmm1, %zmm1
  1697: 62b14548effa             vpxord %zmm18, %zmm7, %zmm7
  1703: 62311548efea             vpxord %zmm18, %zmm13, %zmm13
  1709: 62817d486f541a02         vmovdqa32 0x80(%r10, %r11), %zmm18
  1717: 62b16d48efd2             vpxord %zmm18, %zmm2, %zmm2
  1723: 62313d48efc2             vpxord %zmm18, %zmm8, %zmm8
  1729: 62310d48eff2             vpxord %zmm18, %zmm14, %zmm14
  1735: 62817d486f541a03         vmovdqa32 0xc0(%r10, %r11), %zmm18
  1743: 62b16548efda             vpxord %zmm18, %zmm3, %zmm3
  1749: 62313548efca             vpxord %zmm18, %zmm9, %zmm9
  1755: 62310548effa             vpxord %zmm18, %zmm15, %zmm15
  1761: 62817d486f541a04         vmovdqa32 0x100(%r10, %r11), %zmm18
  1769: 62b15d48efe2             vpxord %zmm18, %zmm4, %zmm4
  1775: 62312d48efd2             vpxord %zmm18, %zmm10, %zmm10
  1781: 62a17d40efc2             vpxord %zmm18, %zmm16, %zmm16
  1787: 62817d486f541a05         vmovdqa32 0x140(%r10, %r11), %zmm18
  1795: 62b15548efea             vpxord %zmm18, %zmm5, %zmm5
  1801: 62312548efda             vpxord %zmm18, %zmm11, %zmm11
  1807: 62a17540efca             vpxord %zmm18, %zmm17, %zmm17
  1813: 4d8b11                   movq (%r9), %r10
  1816: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  1823: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  1831: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
  1839: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
  1847: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
  1855: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
  1863: 4d8b5108                 movq 8(%r9), %r10
  1867: 62917d487f341a           vmovdqa32 %zmm6, (%r10, %r11)
  1874: 62917d487f7c1a01         vmovdqa32 %zmm7, 0x40(%r10, %r11)
  1882: 62117d487f441a02         vmovdqa32 %zmm8, 0x80(%r10, %r11)
  1890: 62117d487f4c1a03         vmovdqa32 %zmm9, 0xc0(%r10, %r11)
  1898: 62117d487f541a04         vmovdqa32 %zmm10, 0x100(%r10, %r11)
  1906: 62117d487f5c1a05         vmovdqa32 %zmm11, 0x140(%r10, %r11)
  1914: 4d8b5110                 movq 0x10(%r9), %r10
  1918: 62117d487f241a           vmovdqa32 %zmm12, (%r10, %r11)
  1925: 62117d487f6c1a01         vmovdqa32 %zmm13, 0x40(%r10, %r11)
  1933: 62117d487f741a02         vmovdqa32 %zmm14, 0x80(%r10, %r11)
  1941: 62117d487f7c1a03         vmovdqa32 %zmm15, 0xc0(%r10, %r11)
  1949: 62817d487f441a04         vmovdqa32 %zmm16, 0x100(%r10, %r11)
  1957: 62817d487f4c1a05         vmovdqa32 %zmm17, 0x140(%r10, %r11)
  1965: 4981c380010000           addq $0x180, %r11
  1972: 4c3bd8                   cmpq %rax, %r11
  1975: 0f8683f8ffff             jbe 0x40
  1981: 4c3bd9                   cmpq %rcx, %r11
  1984: 0f83b2020000             jae 0xa78
  1990: 4153                     pushq %r11
  1992: 498dbf40010000           leaq 0x140(%r15), %rdi
  1999: 498db740020000           leaq 0x240(%r15), %rsi
  2006: 4d8b11                   movq (%r9), %r10
  2009: 4c8b1c24                 movq (%rsp), %r11
  2013: 33db                     xorl %ebx, %ebx
  2015: 498b00                   movq (%r8), %rax
  2018: 420fb61418               movzbl (%rax, %r11), %edx
  2023: 32da                     xorb %dl, %bl
  2025: 498b4008                 movq 8(%r8), %rax
  2029: 420fb61418               movzbl (%rax, %r11), %edx
  2034: 32da                     xorb %dl, %bl
  2036: 498b4010                 movq 0x10(%r8), %rax
  2040: 420fb61418               movzbl (%rax, %r11), %edx
  2045: 32da                     xorb %dl, %bl
  2047: 498b4018                 movq 0x18(%r8), %rax
  2051: 420fb61418               movzbl (%rax, %r11), %edx
  2056: 32da                     xorb %dl, %bl
  2058: 498b4020                 movq 0x20(%r8), %rax
  2062: 420fb61418               movzbl (%rax, %r11), %edx
  2067: 32da                     xorb %dl, %bl
  2069: 498b4028                 movq 0x28(%r8), %rax
  2073: 420fb61418               movzbl (%rax, %r11), %edx
  2078: 32da                     xorb %dl, %bl
  2080: 43881c1a                 movb %bl, (%r10, %r11)
  2084: 33db                     xorl %ebx, %ebx
  2086: 49ffc3                   incq %r11
  2089: 4c3bd9                   cmpq %rcx, %r11
  2092: 0f82adffffff             jb 0x7df
  2098: 4d8b5108                 movq 8(%r9), %r10
  2102: 4c8b1c24                 movq (%rsp), %r11
  2106: 33db                     xorl %ebx, %ebx
  2108: 498b00                   movq (%r8), %rax
  2111: 420fb61418               movzbl (%rax, %r11), %edx
  2116: 4885d2                   testq %rdx, %rdx
  2119: 0f8422000000             je 0x86f
  2125: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2129: 81c0e6000000             addl $0xe6, %eax
  2135: 81f8ff000000             cmpl $0xff, %eax
  2141: 0f8206000000             jb 0x869
  2147: 81c001ffffff             addl $0xffffff01, %eax
  2153: 0fb60406                 movzbl (%rsi, %rax), %eax
  2157: 32d8                     xorb %al, %bl
  2159: 498b4008                 movq 8(%r8), %rax
  2163: 420fb61418               movzbl (%rax, %r11), %edx
  2168: 4885d2                   testq %rdx, %rdx
  2171: 0f8422000000             je 0x8a3
  2177: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2181: 81c0cd000000             addl $0xcd, %eax
  2187: 81f8ff000000             cmpl $0xff, %eax
  2193: 0f8206000000             jb 0x89d
  2199: 81c001ffffff             addl $0xffffff01, %eax
  2205: 0fb60406                 movzbl (%rsi, %rax), %eax
  2209: 32d8                     xorb %al, %bl
  2211: 498b4010                 movq 0x10(%r8), %rax
  2215: 420fb61418               movzbl (%rax, %r11), %edx
  2220: 4885d2                   testq %rdx, %rdx
  2223: 0f841f000000             je 0x8d4
  2229: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2233: 83c04b                   addl $0x4b, %eax
  2236: 81f8ff000000             cmpl $0xff, %eax
  2242: 0f8206000000             jb 0x8ce
  2248: 81c001ffffff             addl $0xffffff01, %eax
  2254: 0fb60406                 movzbl (%rsi, %rax), %eax
  2258: 32d8                     xorb %al, %bl
  2260: 498b4018                 movq 0x18(%r8), %rax
  2264: 420fb61418               movzbl (%rax, %r11), %edx
  2269: 4885d2                   testq %rdx, %rdx
  2272: 0f841f000000             je 0x905
  2278: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2282: 83c032                   addl $0x32, %eax
  2285: 81f8ff000000             cmpl $0xff, %eax
  2291: 0f8206000000             jb 0x8ff
  2297: 81c001ffffff             addl $0xffffff01, %eax
  2303: 0fb60406                 movzbl (%rsi, %rax), %eax
  2307: 32d8                     xorb %al, %bl
  2309: 498b4020                 movq 0x20(%r8), %rax
  2313: 420fb61418               movzbl (%rax, %r11), %edx
  2318: 4885d2                   testq %rdx, %rdx
  2321: 0f841f000000             je 0x936
  2327: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2331: 83c019                   addl $0x19, %eax
  2334: 81f8ff000000             cmpl $0xff, %eax
  2340: 0f8206000000             jb 0x930
  2346: 81c001ffffff             addl $0xffffff01, %eax
  2352: 0fb60406                 movzbl (%rsi, %rax), %eax
  2356: 32d8                     xorb %al, %bl
  2358: 498b4028                 movq 0x28(%r8), %rax
  2362: 420fb61418               movzbl (%rax, %r11), %edx
  2367: 32da                     xorb %dl, %bl
  2369: 43881c1a                 movb %bl, (%r10, %r11)
  2373: 33db                     xorl %ebx, %ebx
  2375: 49ffc3                   incq %r11
  2378: 4c3bd9                   cmpq %rcx, %r11
  2381: 0f82e9feffff             jb 0x83c
  2387: 4d8b5110                 movq 0x10(%r9), %r10
  2391: 4c8b1c24                 movq (%rsp), %r11
  2395: 33db                     xorl %ebx, %ebx
  2397: 498b00                   movq (%r8), %rax
  2400: 420fb61418               movzbl (%rax, %r11), %edx
  2405: 4885d2                   testq %rdx, %rdx
  2408: 0f8422000000             je 0x990
  2414: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2418: 81c0cd000000             addl $0xcd, %eax
  2424: 81f8ff000000             cmpl $0xff, %eax
  2430: 0f8206000000             jb 0x98a
  2436: 81c001ffffff             addl $0xffffff01, %eax
  2442: 0fb60406                 movzbl (%rsi, %rax), %eax
  2446: 32d8                     xorb %al, %bl
  2448: 498b4008                 movq 8(%r8), %rax
  2452: 420fb61418               movzbl (%rax, %r11), %edx
  2457: 4885d2                   testq %rdx, %rdx
  2460: 0f841f000000             je 0x9c1
  2466: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2470: 83c04b                   addl $0x4b, %eax
  2473: 81f8ff000000             cmpl $0xff, %eax
  2479: 0f8206000000             jb 0x9bb
  2485: 81c001ffffff             addl $0xffffff01, %eax
  2491: 0fb60406                 movzbl (%rsi, %rax), %eax
  2495: 32d8                     xorb %al, %bl
  2497: 498b4010                 movq 0x10(%r8), %rax
  2501: 420fb61418               movzbl (%rax, %r11), %edx
  2506: 4885d2                   testq %rdx, %rdx
  2509: 0f841f000000             je 0x9f2
  2515: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2519: 83c019                   addl $0x19, %eax
  2522: 81f8ff000000             cmpl $0xff, %eax
  2528: 0f8206000000             jb 0x9ec
  2534: 81c001ffffff             addl $0xffffff01, %eax
  2540: 0fb60406                 movzbl (%rsi, %rax), %eax
  2544: 32d8                     xorb %al, %bl
  2546: 498b4018                 movq 0x18(%r8), %rax
  2550: 420fb61418               movzbl (%rax, %r11), %edx
  2555: 4885d2                   testq %rdx, %rdx
  2558: 0f8422000000             je 0xa26
  2564: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2568: 81c0e6000000             addl $0xe6, %eax
  2574: 81f8ff000000             cmpl $0xff, %eax
  2580: 0f8206000000             jb 0xa20
  2586: 81c001ffffff             addl $0xffffff01, %eax
  2592: 0fb60406                 movzbl (%rsi, %rax), %eax
  2596: 32d8                     xorb %al, %bl
  2598: 498b4020                 movq 0x20(%r8), %rax
  2602: 420fb61418               movzbl (%rax, %r11), %edx
  2607: 4885d2                   testq %rdx, %rdx
  2610: 0f841f000000             je 0xa57
  2616: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2620: 83c032                   addl $0x32, %eax
  2623: 81f8ff000000             cmpl $0xff, %eax
  2629: 0f8206000000             jb 0xa51
  2635: 81c001ffffff             addl $0xffffff01, %eax
  2641: 0fb60406                 movzbl (%rsi, %rax), %eax
  2645: 32d8                     xorb %al, %bl
  2647: 498b4028                 movq 0x28(%r8), %rax
  2651: 420fb61418               movzbl (%rax, %r11), %edx
  2656: 32da                     xorb %dl, %bl
  2658: 43881c1a                 movb %bl, (%r10, %r11)
  2662: 33db                     xorl %ebx, %ebx
  2664: 49ffc3                   incq %r11
  2667: 4c3bd9                   cmpq %rcx, %r11
  2670: 0f82e9feffff             jb 0x95d
  2676: 4883c408                 addq $8, %rsp
  2680: 415e                     popq %r14
  2682: 415d                     popq %r13
  2684: 415c                     popq %r12
  2686: 415f                     popq %r15
  2688: 5b                       popq %rbx
  2689: c5f877                   vzeroupper 
  2692: c3                       retq 
  2693: 0f0b                     ud2 
  2695: 0f0b                     ud2 
  2697: 0f0b                     ud2 
  2699: 0f0b                     ud2 
  2701: 0f0b                     ud2 
  2703: 0f0b                     ud2 
  2705: 0f0b                     ud2 
  2707: 0f0b                     ud2 
  2709: 0f0b                     ud2 
  2711: 0f0b                     ud2 
  2713: 0f0b                     ud2 
  2715: 0f0b                     ud2 
  2717: 0f0b                     ud2 
  2719: 0f0b                     ud2 
  2721: 0f0b                     ud2 
  2723: 0f0b                     ud2 
  2725: 0f0b                     ud2 
  2727: 0f0b                     ud2 
  2729: 0f0b                     ud2 
  2731: 0f0b                     ud2 
  2733: 0f0b                     ud2 
  2735: 0f0b                     ud2 
  2737: 0f0b                     ud2 
  2739: 0f0b                     ud2 
  2741: 0f0b                     ud2 
  2743: 0f0b                     ud2 
  2745: 0f0b                     ud2 
  2747: 0f0b                     ud2 
  2749: 0f0b                     ud2 
  2751: cc                       int3 
