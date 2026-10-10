; config_t1_196 target=avx512-gfni twin=nt topo=rs(6,3)/rec{0,1} op=recover U=8
; code_size=3136 twins=["nt"] uops_per_byte=0.66015625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d270c0000           leaq 0xc27(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e800020000           subq $0x200, %rax
    57: 0f8224090000             jb 0x963
    63: 90                       nop 
    64: 62a17d486f041e           vmovdqa32 (%rsi, %r11), %zmm16
    71: 62c17d486f17             vmovdqa32 (%r15), %zmm18
    77: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
    84: 62b17d486fc1             vmovdqa32 %zmm17, %zmm0
    90: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
    97: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   104: 62317d486fc1             vmovdqa32 %zmm17, %zmm8
   110: 62a17d486f441e01         vmovdqa32 0x40(%rsi, %r11), %zmm16
   118: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   124: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   131: 62b17d486fc9             vmovdqa32 %zmm17, %zmm1
   137: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   144: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   151: 62317d486fc9             vmovdqa32 %zmm17, %zmm9
   157: 62a17d486f441e02         vmovdqa32 0x80(%rsi, %r11), %zmm16
   165: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   171: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   178: 62b17d486fd1             vmovdqa32 %zmm17, %zmm2
   184: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   191: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   198: 62317d486fd1             vmovdqa32 %zmm17, %zmm10
   204: 62a17d486f441e03         vmovdqa32 0xc0(%rsi, %r11), %zmm16
   212: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   218: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   225: 62b17d486fd9             vmovdqa32 %zmm17, %zmm3
   231: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   238: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   245: 62317d486fd9             vmovdqa32 %zmm17, %zmm11
   251: 62a17d486f441e04         vmovdqa32 0x100(%rsi, %r11), %zmm16
   259: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   265: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   272: 62b17d486fe1             vmovdqa32 %zmm17, %zmm4
   278: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   285: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   292: 62317d486fe1             vmovdqa32 %zmm17, %zmm12
   298: 62a17d486f441e05         vmovdqa32 0x140(%rsi, %r11), %zmm16
   306: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   312: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   319: 62b17d486fe9             vmovdqa32 %zmm17, %zmm5
   325: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   332: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   339: 62317d486fe9             vmovdqa32 %zmm17, %zmm13
   345: 62a17d486f441e06         vmovdqa32 0x180(%rsi, %r11), %zmm16
   353: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   359: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   366: 62b17d486ff1             vmovdqa32 %zmm17, %zmm6
   372: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   379: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   386: 62317d486ff1             vmovdqa32 %zmm17, %zmm14
   392: 62a17d486f441e07         vmovdqa32 0x1c0(%rsi, %r11), %zmm16
   400: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   406: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   413: 62b17d486ff9             vmovdqa32 %zmm17, %zmm7
   419: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   426: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   433: 62317d486ff9             vmovdqa32 %zmm17, %zmm15
   439: 62a17d486f041f           vmovdqa32 (%rdi, %r11), %zmm16
   446: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   453: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   460: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
   466: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   473: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   480: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
   486: 62a17d486f441f01         vmovdqa32 0x40(%rdi, %r11), %zmm16
   494: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   501: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   508: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
   514: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   521: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   528: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
   534: 62a17d486f441f02         vmovdqa32 0x80(%rdi, %r11), %zmm16
   542: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   549: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   556: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
   562: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   569: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   576: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
   582: 62a17d486f441f03         vmovdqa32 0xc0(%rdi, %r11), %zmm16
   590: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   597: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   604: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
   610: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   617: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   624: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
   630: 62a17d486f441f04         vmovdqa32 0x100(%rdi, %r11), %zmm16
   638: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   645: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   652: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
   658: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   665: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   672: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
   678: 62a17d486f441f05         vmovdqa32 0x140(%rdi, %r11), %zmm16
   686: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   693: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   700: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
   706: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   713: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   720: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
   726: 62a17d486f441f06         vmovdqa32 0x180(%rdi, %r11), %zmm16
   734: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   741: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   748: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
   754: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   761: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   768: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
   774: 62a17d486f441f07         vmovdqa32 0x1c0(%rdi, %r11), %zmm16
   782: 62c17d486f5702           vmovdqa32 0x80(%r15), %zmm18
   789: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   796: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
   802: 62c17d486f5703           vmovdqa32 0xc0(%r15), %zmm18
   809: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   816: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
   822: 62817d486f041c           vmovdqa32 (%r12, %r11), %zmm16
   829: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
   836: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   843: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
   849: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
   856: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   863: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
   869: 62817d486f441c01         vmovdqa32 0x40(%r12, %r11), %zmm16
   877: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
   884: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   891: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
   897: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
   904: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   911: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
   917: 62817d486f441c02         vmovdqa32 0x80(%r12, %r11), %zmm16
   925: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
   932: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   939: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
   945: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
   952: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   959: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
   965: 62817d486f441c03         vmovdqa32 0xc0(%r12, %r11), %zmm16
   973: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
   980: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   987: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
   993: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1000: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1007: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
  1013: 62817d486f441c04         vmovdqa32 0x100(%r12, %r11), %zmm16
  1021: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1028: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1035: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  1041: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1048: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1055: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
  1061: 62817d486f441c05         vmovdqa32 0x140(%r12, %r11), %zmm16
  1069: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1076: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1083: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  1089: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1096: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1103: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
  1109: 62817d486f441c06         vmovdqa32 0x180(%r12, %r11), %zmm16
  1117: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1124: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1131: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  1137: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1144: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1151: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
  1157: 62817d486f441c07         vmovdqa32 0x1c0(%r12, %r11), %zmm16
  1165: 62c17d486f5704           vmovdqa32 0x100(%r15), %zmm18
  1172: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1179: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  1185: 62c17d486f5705           vmovdqa32 0x140(%r15), %zmm18
  1192: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1199: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
  1205: 62817d486f441d00         vmovdqa32 (%r13, %r11), %zmm16
  1213: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1220: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1227: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
  1233: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1240: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1247: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
  1253: 62817d486f441d01         vmovdqa32 0x40(%r13, %r11), %zmm16
  1261: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1268: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1275: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
  1281: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1288: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1295: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
  1301: 62817d486f441d02         vmovdqa32 0x80(%r13, %r11), %zmm16
  1309: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1316: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1323: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
  1329: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1336: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1343: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
  1349: 62817d486f441d03         vmovdqa32 0xc0(%r13, %r11), %zmm16
  1357: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1364: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1371: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
  1377: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1384: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1391: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
  1397: 62817d486f441d04         vmovdqa32 0x100(%r13, %r11), %zmm16
  1405: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1412: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1419: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  1425: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1432: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1439: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
  1445: 62817d486f441d05         vmovdqa32 0x140(%r13, %r11), %zmm16
  1453: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1460: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1467: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  1473: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1480: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1487: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
  1493: 62817d486f441d06         vmovdqa32 0x180(%r13, %r11), %zmm16
  1501: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1508: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1515: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  1521: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1528: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1535: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
  1541: 62817d486f441d07         vmovdqa32 0x1c0(%r13, %r11), %zmm16
  1549: 62c17d486f5706           vmovdqa32 0x180(%r15), %zmm18
  1556: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1563: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  1569: 62c17d486f5707           vmovdqa32 0x1c0(%r15), %zmm18
  1576: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1583: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
  1589: 62817d486f041e           vmovdqa32 (%r14, %r11), %zmm16
  1596: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1603: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1610: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
  1616: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1623: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1630: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
  1636: 62817d486f441e01         vmovdqa32 0x40(%r14, %r11), %zmm16
  1644: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1651: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1658: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
  1664: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1671: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1678: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
  1684: 62817d486f441e02         vmovdqa32 0x80(%r14, %r11), %zmm16
  1692: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1699: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1706: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
  1712: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1719: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1726: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
  1732: 62817d486f441e03         vmovdqa32 0xc0(%r14, %r11), %zmm16
  1740: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1747: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1754: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
  1760: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1767: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1774: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
  1780: 62817d486f441e04         vmovdqa32 0x100(%r14, %r11), %zmm16
  1788: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1795: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1802: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  1808: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1815: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1822: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
  1828: 62817d486f441e05         vmovdqa32 0x140(%r14, %r11), %zmm16
  1836: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1843: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1850: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  1856: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1863: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1870: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
  1876: 62817d486f441e06         vmovdqa32 0x180(%r14, %r11), %zmm16
  1884: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1891: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1898: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  1904: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1911: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1918: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
  1924: 62817d486f441e07         vmovdqa32 0x1c0(%r14, %r11), %zmm16
  1932: 62c17d486f5708           vmovdqa32 0x200(%r15), %zmm18
  1939: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1946: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  1952: 62c17d486f5709           vmovdqa32 0x240(%r15), %zmm18
  1959: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1966: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
  1972: 4d8b5028                 movq 0x28(%r8), %r10
  1976: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
  1983: 62c17d486f570a           vmovdqa32 0x280(%r15), %zmm18
  1990: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1997: 62b17d48efc1             vpxord %zmm17, %zmm0, %zmm0
  2003: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2010: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
  2016: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
  2024: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2031: 62b17548efc9             vpxord %zmm17, %zmm1, %zmm1
  2037: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2044: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
  2050: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
  2058: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2065: 62b16d48efd1             vpxord %zmm17, %zmm2, %zmm2
  2071: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2078: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
  2084: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
  2092: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2099: 62b16548efd9             vpxord %zmm17, %zmm3, %zmm3
  2105: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2112: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
  2118: 62817d486f441a04         vmovdqa32 0x100(%r10, %r11), %zmm16
  2126: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2133: 62b15d48efe1             vpxord %zmm17, %zmm4, %zmm4
  2139: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2146: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
  2152: 62817d486f441a05         vmovdqa32 0x140(%r10, %r11), %zmm16
  2160: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2167: 62b15548efe9             vpxord %zmm17, %zmm5, %zmm5
  2173: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2180: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
  2186: 62817d486f441a06         vmovdqa32 0x180(%r10, %r11), %zmm16
  2194: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2201: 62b14d48eff1             vpxord %zmm17, %zmm6, %zmm6
  2207: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2214: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
  2220: 62817d486f441a07         vmovdqa32 0x1c0(%r10, %r11), %zmm16
  2228: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2235: 62b14548eff9             vpxord %zmm17, %zmm7, %zmm7
  2241: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  2248: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
  2254: 4d8b11                   movq (%r9), %r10
  2257: 62917d48e7041a           vmovntdq %zmm0, (%r10, %r11)
  2264: 62917d48e74c1a01         vmovntdq %zmm1, 0x40(%r10, %r11)
  2272: 62917d48e7541a02         vmovntdq %zmm2, 0x80(%r10, %r11)
  2280: 62917d48e75c1a03         vmovntdq %zmm3, 0xc0(%r10, %r11)
  2288: 62917d48e7641a04         vmovntdq %zmm4, 0x100(%r10, %r11)
  2296: 62917d48e76c1a05         vmovntdq %zmm5, 0x140(%r10, %r11)
  2304: 62917d48e7741a06         vmovntdq %zmm6, 0x180(%r10, %r11)
  2312: 62917d48e77c1a07         vmovntdq %zmm7, 0x1c0(%r10, %r11)
  2320: 4d8b5108                 movq 8(%r9), %r10
  2324: 62117d48e7041a           vmovntdq %zmm8, (%r10, %r11)
  2331: 62117d48e74c1a01         vmovntdq %zmm9, 0x40(%r10, %r11)
  2339: 62117d48e7541a02         vmovntdq %zmm10, 0x80(%r10, %r11)
  2347: 62117d48e75c1a03         vmovntdq %zmm11, 0xc0(%r10, %r11)
  2355: 62117d48e7641a04         vmovntdq %zmm12, 0x100(%r10, %r11)
  2363: 62117d48e76c1a05         vmovntdq %zmm13, 0x140(%r10, %r11)
  2371: 62117d48e7741a06         vmovntdq %zmm14, 0x180(%r10, %r11)
  2379: 62117d48e77c1a07         vmovntdq %zmm15, 0x1c0(%r10, %r11)
  2387: 4981c300020000           addq $0x200, %r11
  2394: 4c3bd8                   cmpq %rax, %r11
  2397: 0f86ddf6ffff             jbe 0x40
  2403: 4c3bd9                   cmpq %rcx, %r11
  2406: 0f83a1020000             jae 0xc0d
  2412: 4153                     pushq %r11
  2414: 498dbfc0020000           leaq 0x2c0(%r15), %rdi
  2421: 498db7c0030000           leaq 0x3c0(%r15), %rsi
  2428: 4d8b11                   movq (%r9), %r10
  2431: 4c8b1c24                 movq (%rsp), %r11
  2435: 33db                     xorl %ebx, %ebx
  2437: 498b00                   movq (%r8), %rax
  2440: 420fb61418               movzbl (%rax, %r11), %edx
  2445: 4885d2                   testq %rdx, %rdx
  2448: 0f8422000000             je 0x9b8
  2454: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2458: 81c0c1000000             addl $0xc1, %eax
  2464: 81f8ff000000             cmpl $0xff, %eax
  2470: 0f8206000000             jb 0x9b2
  2476: 81c001ffffff             addl $0xffffff01, %eax
  2482: 0fb60406                 movzbl (%rsi, %rax), %eax
  2486: 32d8                     xorb %al, %bl
  2488: 498b4008                 movq 8(%r8), %rax
  2492: 420fb61418               movzbl (%rax, %r11), %edx
  2497: 4885d2                   testq %rdx, %rdx
  2500: 0f841f000000             je 0x9e9
  2506: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2510: 83c003                   addl $3, %eax
  2513: 81f8ff000000             cmpl $0xff, %eax
  2519: 0f8206000000             jb 0x9e3
  2525: 81c001ffffff             addl $0xffffff01, %eax
  2531: 0fb60406                 movzbl (%rsi, %rax), %eax
  2535: 32d8                     xorb %al, %bl
  2537: 498b4010                 movq 0x10(%r8), %rax
  2541: 420fb61418               movzbl (%rax, %r11), %edx
  2546: 4885d2                   testq %rdx, %rdx
  2549: 0f8422000000             je 0xa1d
  2555: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2559: 81c0c6000000             addl $0xc6, %eax
  2565: 81f8ff000000             cmpl $0xff, %eax
  2571: 0f8206000000             jb 0xa17
  2577: 81c001ffffff             addl $0xffffff01, %eax
  2583: 0fb60406                 movzbl (%rsi, %rax), %eax
  2587: 32d8                     xorb %al, %bl
  2589: 498b4018                 movq 0x18(%r8), %rax
  2593: 420fb61418               movzbl (%rax, %r11), %edx
  2598: 4885d2                   testq %rdx, %rdx
  2601: 0f841f000000             je 0xa4e
  2607: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2611: 83c001                   addl $1, %eax
  2614: 81f8ff000000             cmpl $0xff, %eax
  2620: 0f8206000000             jb 0xa48
  2626: 81c001ffffff             addl $0xffffff01, %eax
  2632: 0fb60406                 movzbl (%rsi, %rax), %eax
  2636: 32d8                     xorb %al, %bl
  2638: 498b4020                 movq 0x20(%r8), %rax
  2642: 420fb61418               movzbl (%rax, %r11), %edx
  2647: 4885d2                   testq %rdx, %rdx
  2650: 0f8422000000             je 0xa82
  2656: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2660: 81c0fe000000             addl $0xfe, %eax
  2666: 81f8ff000000             cmpl $0xff, %eax
  2672: 0f8206000000             jb 0xa7c
  2678: 81c001ffffff             addl $0xffffff01, %eax
  2684: 0fb60406                 movzbl (%rsi, %rax), %eax
  2688: 32d8                     xorb %al, %bl
  2690: 498b4028                 movq 0x28(%r8), %rax
  2694: 420fb61418               movzbl (%rax, %r11), %edx
  2699: 4885d2                   testq %rdx, %rdx
  2702: 0f841f000000             je 0xab3
  2708: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2712: 83c031                   addl $0x31, %eax
  2715: 81f8ff000000             cmpl $0xff, %eax
  2721: 0f8206000000             jb 0xaad
  2727: 81c001ffffff             addl $0xffffff01, %eax
  2733: 0fb60406                 movzbl (%rsi, %rax), %eax
  2737: 32d8                     xorb %al, %bl
  2739: 43881c1a                 movb %bl, (%r10, %r11)
  2743: 33db                     xorl %ebx, %ebx
  2745: 49ffc3                   incq %r11
  2748: 4c3bd9                   cmpq %rcx, %r11
  2751: 0f82c0feffff             jb 0x985
  2757: 4d8b5108                 movq 8(%r9), %r10
  2761: 4c8b1c24                 movq (%rsp), %r11
  2765: 33db                     xorl %ebx, %ebx
  2767: 498b00                   movq (%r8), %rax
  2770: 420fb61418               movzbl (%rax, %r11), %edx
  2775: 4885d2                   testq %rdx, %rdx
  2778: 0f841f000000             je 0xaff
  2784: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2788: 83c01c                   addl $0x1c, %eax
  2791: 81f8ff000000             cmpl $0xff, %eax
  2797: 0f8206000000             jb 0xaf9
  2803: 81c001ffffff             addl $0xffffff01, %eax
  2809: 0fb60406                 movzbl (%rsi, %rax), %eax
  2813: 32d8                     xorb %al, %bl
  2815: 498b4008                 movq 8(%r8), %rax
  2819: 420fb61418               movzbl (%rax, %r11), %edx
  2824: 4885d2                   testq %rdx, %rdx
  2827: 0f8422000000             je 0xb33
  2833: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2837: 81c0df000000             addl $0xdf, %eax
  2843: 81f8ff000000             cmpl $0xff, %eax
  2849: 0f8206000000             jb 0xb2d
  2855: 81c001ffffff             addl $0xffffff01, %eax
  2861: 0fb60406                 movzbl (%rsi, %rax), %eax
  2865: 32d8                     xorb %al, %bl
  2867: 498b4010                 movq 0x10(%r8), %rax
  2871: 420fb61418               movzbl (%rax, %r11), %edx
  2876: 4885d2                   testq %rdx, %rdx
  2879: 0f841f000000             je 0xb64
  2885: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2889: 83c01a                   addl $0x1a, %eax
  2892: 81f8ff000000             cmpl $0xff, %eax
  2898: 0f8206000000             jb 0xb5e
  2904: 81c001ffffff             addl $0xffffff01, %eax
  2910: 0fb60406                 movzbl (%rsi, %rax), %eax
  2914: 32d8                     xorb %al, %bl
  2916: 498b4018                 movq 0x18(%r8), %rax
  2920: 420fb61418               movzbl (%rax, %r11), %edx
  2925: 4885d2                   testq %rdx, %rdx
  2928: 0f841f000000             je 0xb95
  2934: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2938: 83c019                   addl $0x19, %eax
  2941: 81f8ff000000             cmpl $0xff, %eax
  2947: 0f8206000000             jb 0xb8f
  2953: 81c001ffffff             addl $0xffffff01, %eax
  2959: 0fb60406                 movzbl (%rsi, %rax), %eax
  2963: 32d8                     xorb %al, %bl
  2965: 498b4020                 movq 0x20(%r8), %rax
  2969: 420fb61418               movzbl (%rax, %r11), %edx
  2974: 4885d2                   testq %rdx, %rdx
  2977: 0f841f000000             je 0xbc6
  2983: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2987: 83c018                   addl $0x18, %eax
  2990: 81f8ff000000             cmpl $0xff, %eax
  2996: 0f8206000000             jb 0xbc0
  3002: 81c001ffffff             addl $0xffffff01, %eax
  3008: 0fb60406                 movzbl (%rsi, %rax), %eax
  3012: 32d8                     xorb %al, %bl
  3014: 498b4028                 movq 0x28(%r8), %rax
  3018: 420fb61418               movzbl (%rax, %r11), %edx
  3023: 4885d2                   testq %rdx, %rdx
  3026: 0f841f000000             je 0xbf7
  3032: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3036: 83c031                   addl $0x31, %eax
  3039: 81f8ff000000             cmpl $0xff, %eax
  3045: 0f8206000000             jb 0xbf1
  3051: 81c001ffffff             addl $0xffffff01, %eax
  3057: 0fb60406                 movzbl (%rsi, %rax), %eax
  3061: 32d8                     xorb %al, %bl
  3063: 43881c1a                 movb %bl, (%r10, %r11)
  3067: 33db                     xorl %ebx, %ebx
  3069: 49ffc3                   incq %r11
  3072: 4c3bd9                   cmpq %rcx, %r11
  3075: 0f82c6feffff             jb 0xacf
  3081: 4883c408                 addq $8, %rsp
  3085: 415e                     popq %r14
  3087: 415d                     popq %r13
  3089: 415c                     popq %r12
  3091: 415f                     popq %r15
  3093: 5b                       popq %rbx
  3094: c5f877                   vzeroupper 
  3097: c3                       retq 
  3098: 0f0b                     ud2 
  3100: 0f0b                     ud2 
  3102: 0f0b                     ud2 
  3104: 0f0b                     ud2 
  3106: 0f0b                     ud2 
  3108: 0f0b                     ud2 
  3110: 0f0b                     ud2 
  3112: 0f0b                     ud2 
  3114: 0f0b                     ud2 
  3116: 0f0b                     ud2 
  3118: 0f0b                     ud2 
  3120: 0f0b                     ud2 
  3122: 0f0b                     ud2 
  3124: 0f0b                     ud2 
  3126: 0f0b                     ud2 
  3128: 0f0b                     ud2 
  3130: 0f0b                     ud2 
  3132: 0f0b                     ud2 
  3134: 0f0b                     ud2 
