; config_t2_2051 target=avx512-gfni twin=ternlog topo=rs(6,3)/enc op=encode U=2
; code_size=1472 twins=["ternlog"] uops_per_byte=0.7265625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da7050000           leaq 0x5a7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e880000000           subq $0x80, %rax
    57: 0f8286020000             jb 0x2c5
    63: 90                       nop 
    64: 62b17d486f341e           vmovdqa32 (%rsi, %r11), %zmm6
    71: 62f17d486fc6             vmovdqa32 %zmm6, %zmm0
    77: 62517d486f07             vmovdqa32 (%r15), %zmm8
    83: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
    90: 62f17d486fd7             vmovdqa32 %zmm7, %zmm2
    96: 62517d486f4701           vmovdqa32 0x40(%r15), %zmm8
   103: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   110: 62f17d486fe7             vmovdqa32 %zmm7, %zmm4
   116: 62b17d486f741e01         vmovdqa32 0x40(%rsi, %r11), %zmm6
   124: 62f17d486fce             vmovdqa32 %zmm6, %zmm1
   130: 62517d486f07             vmovdqa32 (%r15), %zmm8
   136: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   143: 62f17d486fdf             vmovdqa32 %zmm7, %zmm3
   149: 62517d486f4701           vmovdqa32 0x40(%r15), %zmm8
   156: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   163: 62f17d486fef             vmovdqa32 %zmm7, %zmm5
   169: 62b17d486f341f           vmovdqa32 (%rdi, %r11), %zmm6
   176: 62f17d48efc6             vpxord %zmm6, %zmm0, %zmm0
   182: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   189: 62f16d48efd7             vpxord %zmm7, %zmm2, %zmm2
   195: 62517d486f4702           vmovdqa32 0x80(%r15), %zmm8
   202: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   209: 62f15d48efe7             vpxord %zmm7, %zmm4, %zmm4
   215: 62b17d486f741f01         vmovdqa32 0x40(%rdi, %r11), %zmm6
   223: 62f17548efce             vpxord %zmm6, %zmm1, %zmm1
   229: 62517d486f4701           vmovdqa32 0x40(%r15), %zmm8
   236: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   243: 62f16548efdf             vpxord %zmm7, %zmm3, %zmm3
   249: 62517d486f4702           vmovdqa32 0x80(%r15), %zmm8
   256: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   263: 62f15548efef             vpxord %zmm7, %zmm5, %zmm5
   269: 62917d486f341c           vmovdqa32 (%r12, %r11), %zmm6
   276: 62f17d48efc6             vpxord %zmm6, %zmm0, %zmm0
   282: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   289: 62f16d48efd7             vpxord %zmm7, %zmm2, %zmm2
   295: 62517d486f4703           vmovdqa32 0xc0(%r15), %zmm8
   302: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   309: 62f15d48efe7             vpxord %zmm7, %zmm4, %zmm4
   315: 62917d486f741c01         vmovdqa32 0x40(%r12, %r11), %zmm6
   323: 62f17548efce             vpxord %zmm6, %zmm1, %zmm1
   329: 62517d486f4702           vmovdqa32 0x80(%r15), %zmm8
   336: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   343: 62f16548efdf             vpxord %zmm7, %zmm3, %zmm3
   349: 62517d486f4703           vmovdqa32 0xc0(%r15), %zmm8
   356: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   363: 62f15548efef             vpxord %zmm7, %zmm5, %zmm5
   369: 62917d486f741d00         vmovdqa32 (%r13, %r11), %zmm6
   377: 62f17d48efc6             vpxord %zmm6, %zmm0, %zmm0
   383: 62517d486f4704           vmovdqa32 0x100(%r15), %zmm8
   390: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   397: 62f16d48efd7             vpxord %zmm7, %zmm2, %zmm2
   403: 62517d486f07             vmovdqa32 (%r15), %zmm8
   409: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   416: 62f15d48efe7             vpxord %zmm7, %zmm4, %zmm4
   422: 62917d486f741d01         vmovdqa32 0x40(%r13, %r11), %zmm6
   430: 62f17548efce             vpxord %zmm6, %zmm1, %zmm1
   436: 62517d486f4704           vmovdqa32 0x100(%r15), %zmm8
   443: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   450: 62f16548efdf             vpxord %zmm7, %zmm3, %zmm3
   456: 62517d486f07             vmovdqa32 (%r15), %zmm8
   462: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   469: 62f15548efef             vpxord %zmm7, %zmm5, %zmm5
   475: 62917d486f341e           vmovdqa32 (%r14, %r11), %zmm6
   482: 62f17d48efc6             vpxord %zmm6, %zmm0, %zmm0
   488: 62517d486f4703           vmovdqa32 0xc0(%r15), %zmm8
   495: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   502: 62f16d48efd7             vpxord %zmm7, %zmm2, %zmm2
   508: 62517d486f4704           vmovdqa32 0x100(%r15), %zmm8
   515: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   522: 62f15d48efe7             vpxord %zmm7, %zmm4, %zmm4
   528: 62917d486f741e01         vmovdqa32 0x40(%r14, %r11), %zmm6
   536: 62f17548efce             vpxord %zmm6, %zmm1, %zmm1
   542: 62517d486f4703           vmovdqa32 0xc0(%r15), %zmm8
   549: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   556: 62f16548efdf             vpxord %zmm7, %zmm3, %zmm3
   562: 62517d486f4704           vmovdqa32 0x100(%r15), %zmm8
   569: 62d3cd48cef800           vgf2p8affineqb $0, %zmm8, %zmm6, %zmm7
   576: 62f15548efef             vpxord %zmm7, %zmm5, %zmm5
   582: 4d8b5028                 movq 0x28(%r8), %r10
   586: 62917d486f341a           vmovdqa32 (%r10, %r11), %zmm6
   593: 62f17d48efc6             vpxord %zmm6, %zmm0, %zmm0
   599: 62f16d48efd6             vpxord %zmm6, %zmm2, %zmm2
   605: 62f15d48efe6             vpxord %zmm6, %zmm4, %zmm4
   611: 62917d486f741a01         vmovdqa32 0x40(%r10, %r11), %zmm6
   619: 62f17548efce             vpxord %zmm6, %zmm1, %zmm1
   625: 62f16548efde             vpxord %zmm6, %zmm3, %zmm3
   631: 62f15548efee             vpxord %zmm6, %zmm5, %zmm5
   637: 4d8b11                   movq (%r9), %r10
   640: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   647: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
   655: 4d8b5108                 movq 8(%r9), %r10
   659: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
   666: 62917d487f5c1a01         vmovdqa32 %zmm3, 0x40(%r10, %r11)
   674: 4d8b5110                 movq 0x10(%r9), %r10
   678: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
   685: 62917d487f6c1a01         vmovdqa32 %zmm5, 0x40(%r10, %r11)
   693: 4981c380000000           addq $0x80, %r11
   700: 4c3bd8                   cmpq %rax, %r11
   703: 0f867bfdffff             jbe 0x40
   709: 4c3bd9                   cmpq %rcx, %r11
   712: 0f83b2020000             jae 0x580
   718: 4153                     pushq %r11
   720: 498dbf40010000           leaq 0x140(%r15), %rdi
   727: 498db740020000           leaq 0x240(%r15), %rsi
   734: 4d8b11                   movq (%r9), %r10
   737: 4c8b1c24                 movq (%rsp), %r11
   741: 33db                     xorl %ebx, %ebx
   743: 498b00                   movq (%r8), %rax
   746: 420fb61418               movzbl (%rax, %r11), %edx
   751: 32da                     xorb %dl, %bl
   753: 498b4008                 movq 8(%r8), %rax
   757: 420fb61418               movzbl (%rax, %r11), %edx
   762: 32da                     xorb %dl, %bl
   764: 498b4010                 movq 0x10(%r8), %rax
   768: 420fb61418               movzbl (%rax, %r11), %edx
   773: 32da                     xorb %dl, %bl
   775: 498b4018                 movq 0x18(%r8), %rax
   779: 420fb61418               movzbl (%rax, %r11), %edx
   784: 32da                     xorb %dl, %bl
   786: 498b4020                 movq 0x20(%r8), %rax
   790: 420fb61418               movzbl (%rax, %r11), %edx
   795: 32da                     xorb %dl, %bl
   797: 498b4028                 movq 0x28(%r8), %rax
   801: 420fb61418               movzbl (%rax, %r11), %edx
   806: 32da                     xorb %dl, %bl
   808: 43881c1a                 movb %bl, (%r10, %r11)
   812: 33db                     xorl %ebx, %ebx
   814: 49ffc3                   incq %r11
   817: 4c3bd9                   cmpq %rcx, %r11
   820: 0f82adffffff             jb 0x2e7
   826: 4d8b5108                 movq 8(%r9), %r10
   830: 4c8b1c24                 movq (%rsp), %r11
   834: 33db                     xorl %ebx, %ebx
   836: 498b00                   movq (%r8), %rax
   839: 420fb61418               movzbl (%rax, %r11), %edx
   844: 4885d2                   testq %rdx, %rdx
   847: 0f8422000000             je 0x377
   853: 0fb60417                 movzbl (%rdi, %rdx), %eax
   857: 81c0e6000000             addl $0xe6, %eax
   863: 81f8ff000000             cmpl $0xff, %eax
   869: 0f8206000000             jb 0x371
   875: 81c001ffffff             addl $0xffffff01, %eax
   881: 0fb60406                 movzbl (%rsi, %rax), %eax
   885: 32d8                     xorb %al, %bl
   887: 498b4008                 movq 8(%r8), %rax
   891: 420fb61418               movzbl (%rax, %r11), %edx
   896: 4885d2                   testq %rdx, %rdx
   899: 0f8422000000             je 0x3ab
   905: 0fb60417                 movzbl (%rdi, %rdx), %eax
   909: 81c0cd000000             addl $0xcd, %eax
   915: 81f8ff000000             cmpl $0xff, %eax
   921: 0f8206000000             jb 0x3a5
   927: 81c001ffffff             addl $0xffffff01, %eax
   933: 0fb60406                 movzbl (%rsi, %rax), %eax
   937: 32d8                     xorb %al, %bl
   939: 498b4010                 movq 0x10(%r8), %rax
   943: 420fb61418               movzbl (%rax, %r11), %edx
   948: 4885d2                   testq %rdx, %rdx
   951: 0f841f000000             je 0x3dc
   957: 0fb60417                 movzbl (%rdi, %rdx), %eax
   961: 83c04b                   addl $0x4b, %eax
   964: 81f8ff000000             cmpl $0xff, %eax
   970: 0f8206000000             jb 0x3d6
   976: 81c001ffffff             addl $0xffffff01, %eax
   982: 0fb60406                 movzbl (%rsi, %rax), %eax
   986: 32d8                     xorb %al, %bl
   988: 498b4018                 movq 0x18(%r8), %rax
   992: 420fb61418               movzbl (%rax, %r11), %edx
   997: 4885d2                   testq %rdx, %rdx
  1000: 0f841f000000             je 0x40d
  1006: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1010: 83c032                   addl $0x32, %eax
  1013: 81f8ff000000             cmpl $0xff, %eax
  1019: 0f8206000000             jb 0x407
  1025: 81c001ffffff             addl $0xffffff01, %eax
  1031: 0fb60406                 movzbl (%rsi, %rax), %eax
  1035: 32d8                     xorb %al, %bl
  1037: 498b4020                 movq 0x20(%r8), %rax
  1041: 420fb61418               movzbl (%rax, %r11), %edx
  1046: 4885d2                   testq %rdx, %rdx
  1049: 0f841f000000             je 0x43e
  1055: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1059: 83c019                   addl $0x19, %eax
  1062: 81f8ff000000             cmpl $0xff, %eax
  1068: 0f8206000000             jb 0x438
  1074: 81c001ffffff             addl $0xffffff01, %eax
  1080: 0fb60406                 movzbl (%rsi, %rax), %eax
  1084: 32d8                     xorb %al, %bl
  1086: 498b4028                 movq 0x28(%r8), %rax
  1090: 420fb61418               movzbl (%rax, %r11), %edx
  1095: 32da                     xorb %dl, %bl
  1097: 43881c1a                 movb %bl, (%r10, %r11)
  1101: 33db                     xorl %ebx, %ebx
  1103: 49ffc3                   incq %r11
  1106: 4c3bd9                   cmpq %rcx, %r11
  1109: 0f82e9feffff             jb 0x344
  1115: 4d8b5110                 movq 0x10(%r9), %r10
  1119: 4c8b1c24                 movq (%rsp), %r11
  1123: 33db                     xorl %ebx, %ebx
  1125: 498b00                   movq (%r8), %rax
  1128: 420fb61418               movzbl (%rax, %r11), %edx
  1133: 4885d2                   testq %rdx, %rdx
  1136: 0f8422000000             je 0x498
  1142: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1146: 81c0cd000000             addl $0xcd, %eax
  1152: 81f8ff000000             cmpl $0xff, %eax
  1158: 0f8206000000             jb 0x492
  1164: 81c001ffffff             addl $0xffffff01, %eax
  1170: 0fb60406                 movzbl (%rsi, %rax), %eax
  1174: 32d8                     xorb %al, %bl
  1176: 498b4008                 movq 8(%r8), %rax
  1180: 420fb61418               movzbl (%rax, %r11), %edx
  1185: 4885d2                   testq %rdx, %rdx
  1188: 0f841f000000             je 0x4c9
  1194: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1198: 83c04b                   addl $0x4b, %eax
  1201: 81f8ff000000             cmpl $0xff, %eax
  1207: 0f8206000000             jb 0x4c3
  1213: 81c001ffffff             addl $0xffffff01, %eax
  1219: 0fb60406                 movzbl (%rsi, %rax), %eax
  1223: 32d8                     xorb %al, %bl
  1225: 498b4010                 movq 0x10(%r8), %rax
  1229: 420fb61418               movzbl (%rax, %r11), %edx
  1234: 4885d2                   testq %rdx, %rdx
  1237: 0f841f000000             je 0x4fa
  1243: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1247: 83c019                   addl $0x19, %eax
  1250: 81f8ff000000             cmpl $0xff, %eax
  1256: 0f8206000000             jb 0x4f4
  1262: 81c001ffffff             addl $0xffffff01, %eax
  1268: 0fb60406                 movzbl (%rsi, %rax), %eax
  1272: 32d8                     xorb %al, %bl
  1274: 498b4018                 movq 0x18(%r8), %rax
  1278: 420fb61418               movzbl (%rax, %r11), %edx
  1283: 4885d2                   testq %rdx, %rdx
  1286: 0f8422000000             je 0x52e
  1292: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1296: 81c0e6000000             addl $0xe6, %eax
  1302: 81f8ff000000             cmpl $0xff, %eax
  1308: 0f8206000000             jb 0x528
  1314: 81c001ffffff             addl $0xffffff01, %eax
  1320: 0fb60406                 movzbl (%rsi, %rax), %eax
  1324: 32d8                     xorb %al, %bl
  1326: 498b4020                 movq 0x20(%r8), %rax
  1330: 420fb61418               movzbl (%rax, %r11), %edx
  1335: 4885d2                   testq %rdx, %rdx
  1338: 0f841f000000             je 0x55f
  1344: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1348: 83c032                   addl $0x32, %eax
  1351: 81f8ff000000             cmpl $0xff, %eax
  1357: 0f8206000000             jb 0x559
  1363: 81c001ffffff             addl $0xffffff01, %eax
  1369: 0fb60406                 movzbl (%rsi, %rax), %eax
  1373: 32d8                     xorb %al, %bl
  1375: 498b4028                 movq 0x28(%r8), %rax
  1379: 420fb61418               movzbl (%rax, %r11), %edx
  1384: 32da                     xorb %dl, %bl
  1386: 43881c1a                 movb %bl, (%r10, %r11)
  1390: 33db                     xorl %ebx, %ebx
  1392: 49ffc3                   incq %r11
  1395: 4c3bd9                   cmpq %rcx, %r11
  1398: 0f82e9feffff             jb 0x465
  1404: 4883c408                 addq $8, %rsp
  1408: 415e                     popq %r14
  1410: 415d                     popq %r13
  1412: 415c                     popq %r12
  1414: 415f                     popq %r15
  1416: 5b                       popq %rbx
  1417: c5f877                   vzeroupper 
  1420: c3                       retq 
  1421: 0f0b                     ud2 
  1423: 0f0b                     ud2 
  1425: 0f0b                     ud2 
  1427: 0f0b                     ud2 
  1429: 0f0b                     ud2 
  1431: 0f0b                     ud2 
  1433: 0f0b                     ud2 
  1435: 0f0b                     ud2 
  1437: 0f0b                     ud2 
  1439: 0f0b                     ud2 
  1441: 0f0b                     ud2 
  1443: 0f0b                     ud2 
  1445: 0f0b                     ud2 
  1447: 0f0b                     ud2 
  1449: 0f0b                     ud2 
  1451: 0f0b                     ud2 
  1453: 0f0b                     ud2 
  1455: 0f0b                     ud2 
  1457: 0f0b                     ud2 
  1459: 0f0b                     ud2 
  1461: 0f0b                     ud2 
  1463: 0f0b                     ud2 
  1465: 0f0b                     ud2 
  1467: 0f0b                     ud2 
  1469: 0f0b                     ud2 
  1471: cc                       int3 
