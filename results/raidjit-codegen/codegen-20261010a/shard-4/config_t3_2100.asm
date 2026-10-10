; config_t3_2100 target=avx512-gfni twin=baseline topo=lrc_azure()/rec{0} op=recover U=4
; code_size=1856 twins=[] uops_per_byte=0.6328125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d27070000           leaq 0x727(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e800010000           subq $0x100, %rax
    57: 0f8264040000             jb 0x4a3
    63: 90                       nop 
    64: 62b17d486f241e           vmovdqa32 (%rsi, %r11), %zmm4
    71: 62d17d486f37             vmovdqa32 (%r15), %zmm6
    77: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
    84: 62f17d486fc5             vmovdqa32 %zmm5, %zmm0
    90: 62b17d486f641e01         vmovdqa32 0x40(%rsi, %r11), %zmm4
    98: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   105: 62f17d486fcd             vmovdqa32 %zmm5, %zmm1
   111: 62b17d486f641e02         vmovdqa32 0x80(%rsi, %r11), %zmm4
   119: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   126: 62f17d486fd5             vmovdqa32 %zmm5, %zmm2
   132: 62b17d486f641e03         vmovdqa32 0xc0(%rsi, %r11), %zmm4
   140: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   147: 62f17d486fdd             vmovdqa32 %zmm5, %zmm3
   153: 62b17d486f241f           vmovdqa32 (%rdi, %r11), %zmm4
   160: 62d17d486f7701           vmovdqa32 0x40(%r15), %zmm6
   167: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   174: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   180: 62b17d486f641f01         vmovdqa32 0x40(%rdi, %r11), %zmm4
   188: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   195: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   201: 62b17d486f641f02         vmovdqa32 0x80(%rdi, %r11), %zmm4
   209: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   216: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   222: 62b17d486f641f03         vmovdqa32 0xc0(%rdi, %r11), %zmm4
   230: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   237: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   243: 62917d486f241c           vmovdqa32 (%r12, %r11), %zmm4
   250: 62d17d486f7702           vmovdqa32 0x80(%r15), %zmm6
   257: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   264: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   270: 62917d486f641c01         vmovdqa32 0x40(%r12, %r11), %zmm4
   278: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   285: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   291: 62917d486f641c02         vmovdqa32 0x80(%r12, %r11), %zmm4
   299: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   306: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   312: 62917d486f641c03         vmovdqa32 0xc0(%r12, %r11), %zmm4
   320: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   327: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   333: 62917d486f641d00         vmovdqa32 (%r13, %r11), %zmm4
   341: 62d17d486f7703           vmovdqa32 0xc0(%r15), %zmm6
   348: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   355: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   361: 62917d486f641d01         vmovdqa32 0x40(%r13, %r11), %zmm4
   369: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   376: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   382: 62917d486f641d02         vmovdqa32 0x80(%r13, %r11), %zmm4
   390: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   397: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   403: 62917d486f641d03         vmovdqa32 0xc0(%r13, %r11), %zmm4
   411: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   418: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   424: 62917d486f241e           vmovdqa32 (%r14, %r11), %zmm4
   431: 62d17d486f7704           vmovdqa32 0x100(%r15), %zmm6
   438: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   445: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   451: 62917d486f641e01         vmovdqa32 0x40(%r14, %r11), %zmm4
   459: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   466: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   472: 62917d486f641e02         vmovdqa32 0x80(%r14, %r11), %zmm4
   480: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   487: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   493: 62917d486f641e03         vmovdqa32 0xc0(%r14, %r11), %zmm4
   501: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   508: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   514: 4d8b5028                 movq 0x28(%r8), %r10
   518: 62917d486f241a           vmovdqa32 (%r10, %r11), %zmm4
   525: 62d17d486f7705           vmovdqa32 0x140(%r15), %zmm6
   532: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   539: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   545: 62917d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm4
   553: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   560: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   566: 62917d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm4
   574: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   581: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   587: 62917d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm4
   595: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   602: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   608: 4d8b5030                 movq 0x30(%r8), %r10
   612: 62917d486f241a           vmovdqa32 (%r10, %r11), %zmm4
   619: 62d17d486f7706           vmovdqa32 0x180(%r15), %zmm6
   626: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   633: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   639: 62917d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm4
   647: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   654: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   660: 62917d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm4
   668: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   675: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   681: 62917d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm4
   689: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   696: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   702: 4d8b5038                 movq 0x38(%r8), %r10
   706: 62917d486f241a           vmovdqa32 (%r10, %r11), %zmm4
   713: 62d17d486f7707           vmovdqa32 0x1c0(%r15), %zmm6
   720: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   727: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   733: 62917d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm4
   741: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   748: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   754: 62917d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm4
   762: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   769: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   775: 62917d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm4
   783: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   790: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   796: 4d8b5040                 movq 0x40(%r8), %r10
   800: 62917d486f241a           vmovdqa32 (%r10, %r11), %zmm4
   807: 62d17d486f7708           vmovdqa32 0x200(%r15), %zmm6
   814: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   821: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   827: 62917d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm4
   835: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   842: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   848: 62917d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm4
   856: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   863: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   869: 62917d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm4
   877: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   884: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   890: 4d8b5048                 movq 0x48(%r8), %r10
   894: 62917d486f241a           vmovdqa32 (%r10, %r11), %zmm4
   901: 62d17d486f7709           vmovdqa32 0x240(%r15), %zmm6
   908: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   915: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   921: 62917d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm4
   929: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   936: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   942: 62917d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm4
   950: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   957: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   963: 62917d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm4
   971: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   978: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   984: 4d8b5050                 movq 0x50(%r8), %r10
   988: 62917d486f241a           vmovdqa32 (%r10, %r11), %zmm4
   995: 62d17d486f770a           vmovdqa32 0x280(%r15), %zmm6
  1002: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
  1009: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
  1015: 62917d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm4
  1023: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
  1030: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
  1036: 62917d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm4
  1044: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
  1051: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
  1057: 62917d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm4
  1065: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
  1072: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
  1078: 4d8b5058                 movq 0x58(%r8), %r10
  1082: 62917d486f241a           vmovdqa32 (%r10, %r11), %zmm4
  1089: 62f17d48efc4             vpxord %zmm4, %zmm0, %zmm0
  1095: 62917d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm4
  1103: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
  1109: 62917d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm4
  1117: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
  1123: 62917d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm4
  1131: 62f16548efdc             vpxord %zmm4, %zmm3, %zmm3
  1137: 4d8b11                   movq (%r9), %r10
  1140: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  1147: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  1155: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
  1163: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
  1171: 4981c300010000           addq $0x100, %r11
  1178: 4c3bd8                   cmpq %rax, %r11
  1181: 0f869dfbffff             jbe 0x40
  1187: 4c3bd9                   cmpq %rcx, %r11
  1190: 0f8363020000             jae 0x70f
  1196: 4153                     pushq %r11
  1198: 498dbfc0020000           leaq 0x2c0(%r15), %rdi
  1205: 498db7c0030000           leaq 0x3c0(%r15), %rsi
  1212: 4d8b11                   movq (%r9), %r10
  1215: 4c8b1c24                 movq (%rsp), %r11
  1219: 33db                     xorl %ebx, %ebx
  1221: 498b00                   movq (%r8), %rax
  1224: 420fb61418               movzbl (%rax, %r11), %edx
  1229: 4885d2                   testq %rdx, %rdx
  1232: 0f841f000000             je 0x4f5
  1238: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1242: 83c019                   addl $0x19, %eax
  1245: 81f8ff000000             cmpl $0xff, %eax
  1251: 0f8206000000             jb 0x4ef
  1257: 81c001ffffff             addl $0xffffff01, %eax
  1263: 0fb60406                 movzbl (%rsi, %rax), %eax
  1267: 32d8                     xorb %al, %bl
  1269: 498b4008                 movq 8(%r8), %rax
  1273: 420fb61418               movzbl (%rax, %r11), %edx
  1278: 4885d2                   testq %rdx, %rdx
  1281: 0f841f000000             je 0x526
  1287: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1291: 83c032                   addl $0x32, %eax
  1294: 81f8ff000000             cmpl $0xff, %eax
  1300: 0f8206000000             jb 0x520
  1306: 81c001ffffff             addl $0xffffff01, %eax
  1312: 0fb60406                 movzbl (%rsi, %rax), %eax
  1316: 32d8                     xorb %al, %bl
  1318: 498b4010                 movq 0x10(%r8), %rax
  1322: 420fb61418               movzbl (%rax, %r11), %edx
  1327: 4885d2                   testq %rdx, %rdx
  1330: 0f841f000000             je 0x557
  1336: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1340: 83c04b                   addl $0x4b, %eax
  1343: 81f8ff000000             cmpl $0xff, %eax
  1349: 0f8206000000             jb 0x551
  1355: 81c001ffffff             addl $0xffffff01, %eax
  1361: 0fb60406                 movzbl (%rsi, %rax), %eax
  1365: 32d8                     xorb %al, %bl
  1367: 498b4018                 movq 0x18(%r8), %rax
  1371: 420fb61418               movzbl (%rax, %r11), %edx
  1376: 4885d2                   testq %rdx, %rdx
  1379: 0f841f000000             je 0x588
  1385: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1389: 83c064                   addl $0x64, %eax
  1392: 81f8ff000000             cmpl $0xff, %eax
  1398: 0f8206000000             jb 0x582
  1404: 81c001ffffff             addl $0xffffff01, %eax
  1410: 0fb60406                 movzbl (%rsi, %rax), %eax
  1414: 32d8                     xorb %al, %bl
  1416: 498b4020                 movq 0x20(%r8), %rax
  1420: 420fb61418               movzbl (%rax, %r11), %edx
  1425: 4885d2                   testq %rdx, %rdx
  1428: 0f841f000000             je 0x5b9
  1434: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1438: 83c07d                   addl $0x7d, %eax
  1441: 81f8ff000000             cmpl $0xff, %eax
  1447: 0f8206000000             jb 0x5b3
  1453: 81c001ffffff             addl $0xffffff01, %eax
  1459: 0fb60406                 movzbl (%rsi, %rax), %eax
  1463: 32d8                     xorb %al, %bl
  1465: 498b4028                 movq 0x28(%r8), %rax
  1469: 420fb61418               movzbl (%rax, %r11), %edx
  1474: 4885d2                   testq %rdx, %rdx
  1477: 0f8422000000             je 0x5ed
  1483: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1487: 81c096000000             addl $0x96, %eax
  1493: 81f8ff000000             cmpl $0xff, %eax
  1499: 0f8206000000             jb 0x5e7
  1505: 81c001ffffff             addl $0xffffff01, %eax
  1511: 0fb60406                 movzbl (%rsi, %rax), %eax
  1515: 32d8                     xorb %al, %bl
  1517: 498b4030                 movq 0x30(%r8), %rax
  1521: 420fb61418               movzbl (%rax, %r11), %edx
  1526: 4885d2                   testq %rdx, %rdx
  1529: 0f8422000000             je 0x621
  1535: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1539: 81c0af000000             addl $0xaf, %eax
  1545: 81f8ff000000             cmpl $0xff, %eax
  1551: 0f8206000000             jb 0x61b
  1557: 81c001ffffff             addl $0xffffff01, %eax
  1563: 0fb60406                 movzbl (%rsi, %rax), %eax
  1567: 32d8                     xorb %al, %bl
  1569: 498b4038                 movq 0x38(%r8), %rax
  1573: 420fb61418               movzbl (%rax, %r11), %edx
  1578: 4885d2                   testq %rdx, %rdx
  1581: 0f8422000000             je 0x655
  1587: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1591: 81c0c8000000             addl $0xc8, %eax
  1597: 81f8ff000000             cmpl $0xff, %eax
  1603: 0f8206000000             jb 0x64f
  1609: 81c001ffffff             addl $0xffffff01, %eax
  1615: 0fb60406                 movzbl (%rsi, %rax), %eax
  1619: 32d8                     xorb %al, %bl
  1621: 498b4040                 movq 0x40(%r8), %rax
  1625: 420fb61418               movzbl (%rax, %r11), %edx
  1630: 4885d2                   testq %rdx, %rdx
  1633: 0f8422000000             je 0x689
  1639: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1643: 81c0e1000000             addl $0xe1, %eax
  1649: 81f8ff000000             cmpl $0xff, %eax
  1655: 0f8206000000             jb 0x683
  1661: 81c001ffffff             addl $0xffffff01, %eax
  1667: 0fb60406                 movzbl (%rsi, %rax), %eax
  1671: 32d8                     xorb %al, %bl
  1673: 498b4048                 movq 0x48(%r8), %rax
  1677: 420fb61418               movzbl (%rax, %r11), %edx
  1682: 4885d2                   testq %rdx, %rdx
  1685: 0f8422000000             je 0x6bd
  1691: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1695: 81c0fa000000             addl $0xfa, %eax
  1701: 81f8ff000000             cmpl $0xff, %eax
  1707: 0f8206000000             jb 0x6b7
  1713: 81c001ffffff             addl $0xffffff01, %eax
  1719: 0fb60406                 movzbl (%rsi, %rax), %eax
  1723: 32d8                     xorb %al, %bl
  1725: 498b4050                 movq 0x50(%r8), %rax
  1729: 420fb61418               movzbl (%rax, %r11), %edx
  1734: 4885d2                   testq %rdx, %rdx
  1737: 0f841f000000             je 0x6ee
  1743: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1747: 83c014                   addl $0x14, %eax
  1750: 81f8ff000000             cmpl $0xff, %eax
  1756: 0f8206000000             jb 0x6e8
  1762: 81c001ffffff             addl $0xffffff01, %eax
  1768: 0fb60406                 movzbl (%rsi, %rax), %eax
  1772: 32d8                     xorb %al, %bl
  1774: 498b4058                 movq 0x58(%r8), %rax
  1778: 420fb61418               movzbl (%rax, %r11), %edx
  1783: 32da                     xorb %dl, %bl
  1785: 43881c1a                 movb %bl, (%r10, %r11)
  1789: 33db                     xorl %ebx, %ebx
  1791: 49ffc3                   incq %r11
  1794: 4c3bd9                   cmpq %rcx, %r11
  1797: 0f82bafdffff             jb 0x4c5
  1803: 4883c408                 addq $8, %rsp
  1807: 415e                     popq %r14
  1809: 415d                     popq %r13
  1811: 415c                     popq %r12
  1813: 415f                     popq %r15
  1815: 5b                       popq %rbx
  1816: c5f877                   vzeroupper 
  1819: c3                       retq 
  1820: 0f0b                     ud2 
  1822: 0f0b                     ud2 
  1824: 0f0b                     ud2 
  1826: 0f0b                     ud2 
  1828: 0f0b                     ud2 
  1830: 0f0b                     ud2 
  1832: 0f0b                     ud2 
  1834: 0f0b                     ud2 
  1836: 0f0b                     ud2 
  1838: 0f0b                     ud2 
  1840: 0f0b                     ud2 
  1842: 0f0b                     ud2 
  1844: 0f0b                     ud2 
  1846: 0f0b                     ud2 
  1848: 0f0b                     ud2 
  1850: 0f0b                     ud2 
  1852: 0f0b                     ud2 
  1854: 0f0b                     ud2 
