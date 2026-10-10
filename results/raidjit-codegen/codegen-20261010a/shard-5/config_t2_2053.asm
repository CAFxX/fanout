; config_t2_2053 target=avx512-gfni twin=ternlog topo=rs(6,3)/enc op=encode U=4
; code_size=2112 twins=["ternlog"] uops_per_byte=0.73046875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d27080000           leaq 0x827(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e800010000           subq $0x100, %rax
    57: 0f8202050000             jb 0x541
    63: 90                       nop 
    64: 62317d486f241e           vmovdqa32 (%rsi, %r11), %zmm12
    71: 62d17d486fc4             vmovdqa32 %zmm12, %zmm0
    77: 62517d486f37             vmovdqa32 (%r15), %zmm14
    83: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
    90: 62d17d486fe5             vmovdqa32 %zmm13, %zmm4
    96: 62517d486f7701           vmovdqa32 0x40(%r15), %zmm14
   103: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   110: 62517d486fc5             vmovdqa32 %zmm13, %zmm8
   116: 62317d486f641e01         vmovdqa32 0x40(%rsi, %r11), %zmm12
   124: 62d17d486fcc             vmovdqa32 %zmm12, %zmm1
   130: 62517d486f37             vmovdqa32 (%r15), %zmm14
   136: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   143: 62d17d486fed             vmovdqa32 %zmm13, %zmm5
   149: 62517d486f7701           vmovdqa32 0x40(%r15), %zmm14
   156: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   163: 62517d486fcd             vmovdqa32 %zmm13, %zmm9
   169: 62317d486f641e02         vmovdqa32 0x80(%rsi, %r11), %zmm12
   177: 62d17d486fd4             vmovdqa32 %zmm12, %zmm2
   183: 62517d486f37             vmovdqa32 (%r15), %zmm14
   189: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   196: 62d17d486ff5             vmovdqa32 %zmm13, %zmm6
   202: 62517d486f7701           vmovdqa32 0x40(%r15), %zmm14
   209: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   216: 62517d486fd5             vmovdqa32 %zmm13, %zmm10
   222: 62317d486f641e03         vmovdqa32 0xc0(%rsi, %r11), %zmm12
   230: 62d17d486fdc             vmovdqa32 %zmm12, %zmm3
   236: 62517d486f37             vmovdqa32 (%r15), %zmm14
   242: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   249: 62d17d486ffd             vmovdqa32 %zmm13, %zmm7
   255: 62517d486f7701           vmovdqa32 0x40(%r15), %zmm14
   262: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   269: 62517d486fdd             vmovdqa32 %zmm13, %zmm11
   275: 62317d486f241f           vmovdqa32 (%rdi, %r11), %zmm12
   282: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   288: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   295: 62d15d48efe5             vpxord %zmm13, %zmm4, %zmm4
   301: 62517d486f7702           vmovdqa32 0x80(%r15), %zmm14
   308: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   315: 62513d48efc5             vpxord %zmm13, %zmm8, %zmm8
   321: 62317d486f641f01         vmovdqa32 0x40(%rdi, %r11), %zmm12
   329: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   335: 62517d486f7701           vmovdqa32 0x40(%r15), %zmm14
   342: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   349: 62d15548efed             vpxord %zmm13, %zmm5, %zmm5
   355: 62517d486f7702           vmovdqa32 0x80(%r15), %zmm14
   362: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   369: 62513548efcd             vpxord %zmm13, %zmm9, %zmm9
   375: 62317d486f641f02         vmovdqa32 0x80(%rdi, %r11), %zmm12
   383: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   389: 62517d486f7701           vmovdqa32 0x40(%r15), %zmm14
   396: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   403: 62d14d48eff5             vpxord %zmm13, %zmm6, %zmm6
   409: 62517d486f7702           vmovdqa32 0x80(%r15), %zmm14
   416: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   423: 62512d48efd5             vpxord %zmm13, %zmm10, %zmm10
   429: 62317d486f641f03         vmovdqa32 0xc0(%rdi, %r11), %zmm12
   437: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   443: 62517d486f7701           vmovdqa32 0x40(%r15), %zmm14
   450: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   457: 62d14548effd             vpxord %zmm13, %zmm7, %zmm7
   463: 62517d486f7702           vmovdqa32 0x80(%r15), %zmm14
   470: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   477: 62512548efdd             vpxord %zmm13, %zmm11, %zmm11
   483: 62117d486f241c           vmovdqa32 (%r12, %r11), %zmm12
   490: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   496: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   503: 62d15d48efe5             vpxord %zmm13, %zmm4, %zmm4
   509: 62517d486f7703           vmovdqa32 0xc0(%r15), %zmm14
   516: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   523: 62513d48efc5             vpxord %zmm13, %zmm8, %zmm8
   529: 62117d486f641c01         vmovdqa32 0x40(%r12, %r11), %zmm12
   537: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   543: 62517d486f7702           vmovdqa32 0x80(%r15), %zmm14
   550: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   557: 62d15548efed             vpxord %zmm13, %zmm5, %zmm5
   563: 62517d486f7703           vmovdqa32 0xc0(%r15), %zmm14
   570: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   577: 62513548efcd             vpxord %zmm13, %zmm9, %zmm9
   583: 62117d486f641c02         vmovdqa32 0x80(%r12, %r11), %zmm12
   591: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   597: 62517d486f7702           vmovdqa32 0x80(%r15), %zmm14
   604: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   611: 62d14d48eff5             vpxord %zmm13, %zmm6, %zmm6
   617: 62517d486f7703           vmovdqa32 0xc0(%r15), %zmm14
   624: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   631: 62512d48efd5             vpxord %zmm13, %zmm10, %zmm10
   637: 62117d486f641c03         vmovdqa32 0xc0(%r12, %r11), %zmm12
   645: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   651: 62517d486f7702           vmovdqa32 0x80(%r15), %zmm14
   658: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   665: 62d14548effd             vpxord %zmm13, %zmm7, %zmm7
   671: 62517d486f7703           vmovdqa32 0xc0(%r15), %zmm14
   678: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   685: 62512548efdd             vpxord %zmm13, %zmm11, %zmm11
   691: 62117d486f641d00         vmovdqa32 (%r13, %r11), %zmm12
   699: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   705: 62517d486f7704           vmovdqa32 0x100(%r15), %zmm14
   712: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   719: 62d15d48efe5             vpxord %zmm13, %zmm4, %zmm4
   725: 62517d486f37             vmovdqa32 (%r15), %zmm14
   731: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   738: 62513d48efc5             vpxord %zmm13, %zmm8, %zmm8
   744: 62117d486f641d01         vmovdqa32 0x40(%r13, %r11), %zmm12
   752: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   758: 62517d486f7704           vmovdqa32 0x100(%r15), %zmm14
   765: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   772: 62d15548efed             vpxord %zmm13, %zmm5, %zmm5
   778: 62517d486f37             vmovdqa32 (%r15), %zmm14
   784: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   791: 62513548efcd             vpxord %zmm13, %zmm9, %zmm9
   797: 62117d486f641d02         vmovdqa32 0x80(%r13, %r11), %zmm12
   805: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   811: 62517d486f7704           vmovdqa32 0x100(%r15), %zmm14
   818: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   825: 62d14d48eff5             vpxord %zmm13, %zmm6, %zmm6
   831: 62517d486f37             vmovdqa32 (%r15), %zmm14
   837: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   844: 62512d48efd5             vpxord %zmm13, %zmm10, %zmm10
   850: 62117d486f641d03         vmovdqa32 0xc0(%r13, %r11), %zmm12
   858: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   864: 62517d486f7704           vmovdqa32 0x100(%r15), %zmm14
   871: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   878: 62d14548effd             vpxord %zmm13, %zmm7, %zmm7
   884: 62517d486f37             vmovdqa32 (%r15), %zmm14
   890: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   897: 62512548efdd             vpxord %zmm13, %zmm11, %zmm11
   903: 62117d486f241e           vmovdqa32 (%r14, %r11), %zmm12
   910: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   916: 62517d486f7703           vmovdqa32 0xc0(%r15), %zmm14
   923: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   930: 62d15d48efe5             vpxord %zmm13, %zmm4, %zmm4
   936: 62517d486f7704           vmovdqa32 0x100(%r15), %zmm14
   943: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   950: 62513d48efc5             vpxord %zmm13, %zmm8, %zmm8
   956: 62117d486f641e01         vmovdqa32 0x40(%r14, %r11), %zmm12
   964: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   970: 62517d486f7703           vmovdqa32 0xc0(%r15), %zmm14
   977: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
   984: 62d15548efed             vpxord %zmm13, %zmm5, %zmm5
   990: 62517d486f7704           vmovdqa32 0x100(%r15), %zmm14
   997: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
  1004: 62513548efcd             vpxord %zmm13, %zmm9, %zmm9
  1010: 62117d486f641e02         vmovdqa32 0x80(%r14, %r11), %zmm12
  1018: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
  1024: 62517d486f7703           vmovdqa32 0xc0(%r15), %zmm14
  1031: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
  1038: 62d14d48eff5             vpxord %zmm13, %zmm6, %zmm6
  1044: 62517d486f7704           vmovdqa32 0x100(%r15), %zmm14
  1051: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
  1058: 62512d48efd5             vpxord %zmm13, %zmm10, %zmm10
  1064: 62117d486f641e03         vmovdqa32 0xc0(%r14, %r11), %zmm12
  1072: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
  1078: 62517d486f7703           vmovdqa32 0xc0(%r15), %zmm14
  1085: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
  1092: 62d14548effd             vpxord %zmm13, %zmm7, %zmm7
  1098: 62517d486f7704           vmovdqa32 0x100(%r15), %zmm14
  1105: 62539d48ceee00           vgf2p8affineqb $0, %zmm14, %zmm12, %zmm13
  1112: 62512548efdd             vpxord %zmm13, %zmm11, %zmm11
  1118: 4d8b5028                 movq 0x28(%r8), %r10
  1122: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
  1129: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
  1135: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
  1141: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1147: 62117d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm12
  1155: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
  1161: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
  1167: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1173: 62117d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm12
  1181: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
  1187: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  1193: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1199: 62117d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm12
  1207: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
  1213: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  1219: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
  1225: 4d8b11                   movq (%r9), %r10
  1228: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  1235: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  1243: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
  1251: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
  1259: 4d8b5108                 movq 8(%r9), %r10
  1263: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
  1270: 62917d487f6c1a01         vmovdqa32 %zmm5, 0x40(%r10, %r11)
  1278: 62917d487f741a02         vmovdqa32 %zmm6, 0x80(%r10, %r11)
  1286: 62917d487f7c1a03         vmovdqa32 %zmm7, 0xc0(%r10, %r11)
  1294: 4d8b5110                 movq 0x10(%r9), %r10
  1298: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  1305: 62117d487f4c1a01         vmovdqa32 %zmm9, 0x40(%r10, %r11)
  1313: 62117d487f541a02         vmovdqa32 %zmm10, 0x80(%r10, %r11)
  1321: 62117d487f5c1a03         vmovdqa32 %zmm11, 0xc0(%r10, %r11)
  1329: 4981c300010000           addq $0x100, %r11
  1336: 4c3bd8                   cmpq %rax, %r11
  1339: 0f86fffaffff             jbe 0x40
  1345: 4c3bd9                   cmpq %rcx, %r11
  1348: 0f83b2020000             jae 0x7fc
  1354: 4153                     pushq %r11
  1356: 498dbf40010000           leaq 0x140(%r15), %rdi
  1363: 498db740020000           leaq 0x240(%r15), %rsi
  1370: 4d8b11                   movq (%r9), %r10
  1373: 4c8b1c24                 movq (%rsp), %r11
  1377: 33db                     xorl %ebx, %ebx
  1379: 498b00                   movq (%r8), %rax
  1382: 420fb61418               movzbl (%rax, %r11), %edx
  1387: 32da                     xorb %dl, %bl
  1389: 498b4008                 movq 8(%r8), %rax
  1393: 420fb61418               movzbl (%rax, %r11), %edx
  1398: 32da                     xorb %dl, %bl
  1400: 498b4010                 movq 0x10(%r8), %rax
  1404: 420fb61418               movzbl (%rax, %r11), %edx
  1409: 32da                     xorb %dl, %bl
  1411: 498b4018                 movq 0x18(%r8), %rax
  1415: 420fb61418               movzbl (%rax, %r11), %edx
  1420: 32da                     xorb %dl, %bl
  1422: 498b4020                 movq 0x20(%r8), %rax
  1426: 420fb61418               movzbl (%rax, %r11), %edx
  1431: 32da                     xorb %dl, %bl
  1433: 498b4028                 movq 0x28(%r8), %rax
  1437: 420fb61418               movzbl (%rax, %r11), %edx
  1442: 32da                     xorb %dl, %bl
  1444: 43881c1a                 movb %bl, (%r10, %r11)
  1448: 33db                     xorl %ebx, %ebx
  1450: 49ffc3                   incq %r11
  1453: 4c3bd9                   cmpq %rcx, %r11
  1456: 0f82adffffff             jb 0x563
  1462: 4d8b5108                 movq 8(%r9), %r10
  1466: 4c8b1c24                 movq (%rsp), %r11
  1470: 33db                     xorl %ebx, %ebx
  1472: 498b00                   movq (%r8), %rax
  1475: 420fb61418               movzbl (%rax, %r11), %edx
  1480: 4885d2                   testq %rdx, %rdx
  1483: 0f8422000000             je 0x5f3
  1489: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1493: 81c0e6000000             addl $0xe6, %eax
  1499: 81f8ff000000             cmpl $0xff, %eax
  1505: 0f8206000000             jb 0x5ed
  1511: 81c001ffffff             addl $0xffffff01, %eax
  1517: 0fb60406                 movzbl (%rsi, %rax), %eax
  1521: 32d8                     xorb %al, %bl
  1523: 498b4008                 movq 8(%r8), %rax
  1527: 420fb61418               movzbl (%rax, %r11), %edx
  1532: 4885d2                   testq %rdx, %rdx
  1535: 0f8422000000             je 0x627
  1541: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1545: 81c0cd000000             addl $0xcd, %eax
  1551: 81f8ff000000             cmpl $0xff, %eax
  1557: 0f8206000000             jb 0x621
  1563: 81c001ffffff             addl $0xffffff01, %eax
  1569: 0fb60406                 movzbl (%rsi, %rax), %eax
  1573: 32d8                     xorb %al, %bl
  1575: 498b4010                 movq 0x10(%r8), %rax
  1579: 420fb61418               movzbl (%rax, %r11), %edx
  1584: 4885d2                   testq %rdx, %rdx
  1587: 0f841f000000             je 0x658
  1593: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1597: 83c04b                   addl $0x4b, %eax
  1600: 81f8ff000000             cmpl $0xff, %eax
  1606: 0f8206000000             jb 0x652
  1612: 81c001ffffff             addl $0xffffff01, %eax
  1618: 0fb60406                 movzbl (%rsi, %rax), %eax
  1622: 32d8                     xorb %al, %bl
  1624: 498b4018                 movq 0x18(%r8), %rax
  1628: 420fb61418               movzbl (%rax, %r11), %edx
  1633: 4885d2                   testq %rdx, %rdx
  1636: 0f841f000000             je 0x689
  1642: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1646: 83c032                   addl $0x32, %eax
  1649: 81f8ff000000             cmpl $0xff, %eax
  1655: 0f8206000000             jb 0x683
  1661: 81c001ffffff             addl $0xffffff01, %eax
  1667: 0fb60406                 movzbl (%rsi, %rax), %eax
  1671: 32d8                     xorb %al, %bl
  1673: 498b4020                 movq 0x20(%r8), %rax
  1677: 420fb61418               movzbl (%rax, %r11), %edx
  1682: 4885d2                   testq %rdx, %rdx
  1685: 0f841f000000             je 0x6ba
  1691: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1695: 83c019                   addl $0x19, %eax
  1698: 81f8ff000000             cmpl $0xff, %eax
  1704: 0f8206000000             jb 0x6b4
  1710: 81c001ffffff             addl $0xffffff01, %eax
  1716: 0fb60406                 movzbl (%rsi, %rax), %eax
  1720: 32d8                     xorb %al, %bl
  1722: 498b4028                 movq 0x28(%r8), %rax
  1726: 420fb61418               movzbl (%rax, %r11), %edx
  1731: 32da                     xorb %dl, %bl
  1733: 43881c1a                 movb %bl, (%r10, %r11)
  1737: 33db                     xorl %ebx, %ebx
  1739: 49ffc3                   incq %r11
  1742: 4c3bd9                   cmpq %rcx, %r11
  1745: 0f82e9feffff             jb 0x5c0
  1751: 4d8b5110                 movq 0x10(%r9), %r10
  1755: 4c8b1c24                 movq (%rsp), %r11
  1759: 33db                     xorl %ebx, %ebx
  1761: 498b00                   movq (%r8), %rax
  1764: 420fb61418               movzbl (%rax, %r11), %edx
  1769: 4885d2                   testq %rdx, %rdx
  1772: 0f8422000000             je 0x714
  1778: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1782: 81c0cd000000             addl $0xcd, %eax
  1788: 81f8ff000000             cmpl $0xff, %eax
  1794: 0f8206000000             jb 0x70e
  1800: 81c001ffffff             addl $0xffffff01, %eax
  1806: 0fb60406                 movzbl (%rsi, %rax), %eax
  1810: 32d8                     xorb %al, %bl
  1812: 498b4008                 movq 8(%r8), %rax
  1816: 420fb61418               movzbl (%rax, %r11), %edx
  1821: 4885d2                   testq %rdx, %rdx
  1824: 0f841f000000             je 0x745
  1830: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1834: 83c04b                   addl $0x4b, %eax
  1837: 81f8ff000000             cmpl $0xff, %eax
  1843: 0f8206000000             jb 0x73f
  1849: 81c001ffffff             addl $0xffffff01, %eax
  1855: 0fb60406                 movzbl (%rsi, %rax), %eax
  1859: 32d8                     xorb %al, %bl
  1861: 498b4010                 movq 0x10(%r8), %rax
  1865: 420fb61418               movzbl (%rax, %r11), %edx
  1870: 4885d2                   testq %rdx, %rdx
  1873: 0f841f000000             je 0x776
  1879: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1883: 83c019                   addl $0x19, %eax
  1886: 81f8ff000000             cmpl $0xff, %eax
  1892: 0f8206000000             jb 0x770
  1898: 81c001ffffff             addl $0xffffff01, %eax
  1904: 0fb60406                 movzbl (%rsi, %rax), %eax
  1908: 32d8                     xorb %al, %bl
  1910: 498b4018                 movq 0x18(%r8), %rax
  1914: 420fb61418               movzbl (%rax, %r11), %edx
  1919: 4885d2                   testq %rdx, %rdx
  1922: 0f8422000000             je 0x7aa
  1928: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1932: 81c0e6000000             addl $0xe6, %eax
  1938: 81f8ff000000             cmpl $0xff, %eax
  1944: 0f8206000000             jb 0x7a4
  1950: 81c001ffffff             addl $0xffffff01, %eax
  1956: 0fb60406                 movzbl (%rsi, %rax), %eax
  1960: 32d8                     xorb %al, %bl
  1962: 498b4020                 movq 0x20(%r8), %rax
  1966: 420fb61418               movzbl (%rax, %r11), %edx
  1971: 4885d2                   testq %rdx, %rdx
  1974: 0f841f000000             je 0x7db
  1980: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1984: 83c032                   addl $0x32, %eax
  1987: 81f8ff000000             cmpl $0xff, %eax
  1993: 0f8206000000             jb 0x7d5
  1999: 81c001ffffff             addl $0xffffff01, %eax
  2005: 0fb60406                 movzbl (%rsi, %rax), %eax
  2009: 32d8                     xorb %al, %bl
  2011: 498b4028                 movq 0x28(%r8), %rax
  2015: 420fb61418               movzbl (%rax, %r11), %edx
  2020: 32da                     xorb %dl, %bl
  2022: 43881c1a                 movb %bl, (%r10, %r11)
  2026: 33db                     xorl %ebx, %ebx
  2028: 49ffc3                   incq %r11
  2031: 4c3bd9                   cmpq %rcx, %r11
  2034: 0f82e9feffff             jb 0x6e1
  2040: 4883c408                 addq $8, %rsp
  2044: 415e                     popq %r14
  2046: 415d                     popq %r13
  2048: 415c                     popq %r12
  2050: 415f                     popq %r15
  2052: 5b                       popq %rbx
  2053: c5f877                   vzeroupper 
  2056: c3                       retq 
  2057: 0f0b                     ud2 
  2059: 0f0b                     ud2 
  2061: 0f0b                     ud2 
  2063: 0f0b                     ud2 
  2065: 0f0b                     ud2 
  2067: 0f0b                     ud2 
  2069: 0f0b                     ud2 
  2071: 0f0b                     ud2 
  2073: 0f0b                     ud2 
  2075: 0f0b                     ud2 
  2077: 0f0b                     ud2 
  2079: 0f0b                     ud2 
  2081: 0f0b                     ud2 
  2083: 0f0b                     ud2 
  2085: 0f0b                     ud2 
  2087: 0f0b                     ud2 
  2089: 0f0b                     ud2 
  2091: 0f0b                     ud2 
  2093: 0f0b                     ud2 
  2095: 0f0b                     ud2 
  2097: 0f0b                     ud2 
  2099: 0f0b                     ud2 
  2101: 0f0b                     ud2 
  2103: 0f0b                     ud2 
  2105: 0f0b                     ud2 
  2107: 0f0b                     ud2 
  2109: 0f0b                     ud2 
  2111: cc                       int3 
