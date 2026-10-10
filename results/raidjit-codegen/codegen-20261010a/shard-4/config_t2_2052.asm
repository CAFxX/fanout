; config_t2_2052 target=avx512-gfni twin=ternlog topo=rs(6,3)/enc op=encode U=3
; code_size=1792 twins=["ternlog"] uops_per_byte=0.7291666666666666
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de7060000           leaq 0x6e7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e8c0000000           subq $0xc0, %rax
    57: 0f82c4030000             jb 0x403
    63: 90                       nop 
    64: 62317d486f0c1e           vmovdqa32 (%rsi, %r11), %zmm9
    71: 62d17d486fc1             vmovdqa32 %zmm9, %zmm0
    77: 62517d486f1f             vmovdqa32 (%r15), %zmm11
    83: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
    90: 62d17d486fda             vmovdqa32 %zmm10, %zmm3
    96: 62517d486f5f01           vmovdqa32 0x40(%r15), %zmm11
   103: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   110: 62d17d486ff2             vmovdqa32 %zmm10, %zmm6
   116: 62317d486f4c1e01         vmovdqa32 0x40(%rsi, %r11), %zmm9
   124: 62d17d486fc9             vmovdqa32 %zmm9, %zmm1
   130: 62517d486f1f             vmovdqa32 (%r15), %zmm11
   136: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   143: 62d17d486fe2             vmovdqa32 %zmm10, %zmm4
   149: 62517d486f5f01           vmovdqa32 0x40(%r15), %zmm11
   156: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   163: 62d17d486ffa             vmovdqa32 %zmm10, %zmm7
   169: 62317d486f4c1e02         vmovdqa32 0x80(%rsi, %r11), %zmm9
   177: 62d17d486fd1             vmovdqa32 %zmm9, %zmm2
   183: 62517d486f1f             vmovdqa32 (%r15), %zmm11
   189: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   196: 62d17d486fea             vmovdqa32 %zmm10, %zmm5
   202: 62517d486f5f01           vmovdqa32 0x40(%r15), %zmm11
   209: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   216: 62517d486fc2             vmovdqa32 %zmm10, %zmm8
   222: 62317d486f0c1f           vmovdqa32 (%rdi, %r11), %zmm9
   229: 62d17d48efc1             vpxord %zmm9, %zmm0, %zmm0
   235: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   242: 62d16548efda             vpxord %zmm10, %zmm3, %zmm3
   248: 62517d486f5f02           vmovdqa32 0x80(%r15), %zmm11
   255: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   262: 62d14d48eff2             vpxord %zmm10, %zmm6, %zmm6
   268: 62317d486f4c1f01         vmovdqa32 0x40(%rdi, %r11), %zmm9
   276: 62d17548efc9             vpxord %zmm9, %zmm1, %zmm1
   282: 62517d486f5f01           vmovdqa32 0x40(%r15), %zmm11
   289: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   296: 62d15d48efe2             vpxord %zmm10, %zmm4, %zmm4
   302: 62517d486f5f02           vmovdqa32 0x80(%r15), %zmm11
   309: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   316: 62d14548effa             vpxord %zmm10, %zmm7, %zmm7
   322: 62317d486f4c1f02         vmovdqa32 0x80(%rdi, %r11), %zmm9
   330: 62d16d48efd1             vpxord %zmm9, %zmm2, %zmm2
   336: 62517d486f5f01           vmovdqa32 0x40(%r15), %zmm11
   343: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   350: 62d15548efea             vpxord %zmm10, %zmm5, %zmm5
   356: 62517d486f5f02           vmovdqa32 0x80(%r15), %zmm11
   363: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   370: 62513d48efc2             vpxord %zmm10, %zmm8, %zmm8
   376: 62117d486f0c1c           vmovdqa32 (%r12, %r11), %zmm9
   383: 62d17d48efc1             vpxord %zmm9, %zmm0, %zmm0
   389: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   396: 62d16548efda             vpxord %zmm10, %zmm3, %zmm3
   402: 62517d486f5f03           vmovdqa32 0xc0(%r15), %zmm11
   409: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   416: 62d14d48eff2             vpxord %zmm10, %zmm6, %zmm6
   422: 62117d486f4c1c01         vmovdqa32 0x40(%r12, %r11), %zmm9
   430: 62d17548efc9             vpxord %zmm9, %zmm1, %zmm1
   436: 62517d486f5f02           vmovdqa32 0x80(%r15), %zmm11
   443: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   450: 62d15d48efe2             vpxord %zmm10, %zmm4, %zmm4
   456: 62517d486f5f03           vmovdqa32 0xc0(%r15), %zmm11
   463: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   470: 62d14548effa             vpxord %zmm10, %zmm7, %zmm7
   476: 62117d486f4c1c02         vmovdqa32 0x80(%r12, %r11), %zmm9
   484: 62d16d48efd1             vpxord %zmm9, %zmm2, %zmm2
   490: 62517d486f5f02           vmovdqa32 0x80(%r15), %zmm11
   497: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   504: 62d15548efea             vpxord %zmm10, %zmm5, %zmm5
   510: 62517d486f5f03           vmovdqa32 0xc0(%r15), %zmm11
   517: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   524: 62513d48efc2             vpxord %zmm10, %zmm8, %zmm8
   530: 62117d486f4c1d00         vmovdqa32 (%r13, %r11), %zmm9
   538: 62d17d48efc1             vpxord %zmm9, %zmm0, %zmm0
   544: 62517d486f5f04           vmovdqa32 0x100(%r15), %zmm11
   551: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   558: 62d16548efda             vpxord %zmm10, %zmm3, %zmm3
   564: 62517d486f1f             vmovdqa32 (%r15), %zmm11
   570: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   577: 62d14d48eff2             vpxord %zmm10, %zmm6, %zmm6
   583: 62117d486f4c1d01         vmovdqa32 0x40(%r13, %r11), %zmm9
   591: 62d17548efc9             vpxord %zmm9, %zmm1, %zmm1
   597: 62517d486f5f04           vmovdqa32 0x100(%r15), %zmm11
   604: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   611: 62d15d48efe2             vpxord %zmm10, %zmm4, %zmm4
   617: 62517d486f1f             vmovdqa32 (%r15), %zmm11
   623: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   630: 62d14548effa             vpxord %zmm10, %zmm7, %zmm7
   636: 62117d486f4c1d02         vmovdqa32 0x80(%r13, %r11), %zmm9
   644: 62d16d48efd1             vpxord %zmm9, %zmm2, %zmm2
   650: 62517d486f5f04           vmovdqa32 0x100(%r15), %zmm11
   657: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   664: 62d15548efea             vpxord %zmm10, %zmm5, %zmm5
   670: 62517d486f1f             vmovdqa32 (%r15), %zmm11
   676: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   683: 62513d48efc2             vpxord %zmm10, %zmm8, %zmm8
   689: 62117d486f0c1e           vmovdqa32 (%r14, %r11), %zmm9
   696: 62d17d48efc1             vpxord %zmm9, %zmm0, %zmm0
   702: 62517d486f5f03           vmovdqa32 0xc0(%r15), %zmm11
   709: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   716: 62d16548efda             vpxord %zmm10, %zmm3, %zmm3
   722: 62517d486f5f04           vmovdqa32 0x100(%r15), %zmm11
   729: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   736: 62d14d48eff2             vpxord %zmm10, %zmm6, %zmm6
   742: 62117d486f4c1e01         vmovdqa32 0x40(%r14, %r11), %zmm9
   750: 62d17548efc9             vpxord %zmm9, %zmm1, %zmm1
   756: 62517d486f5f03           vmovdqa32 0xc0(%r15), %zmm11
   763: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   770: 62d15d48efe2             vpxord %zmm10, %zmm4, %zmm4
   776: 62517d486f5f04           vmovdqa32 0x100(%r15), %zmm11
   783: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   790: 62d14548effa             vpxord %zmm10, %zmm7, %zmm7
   796: 62117d486f4c1e02         vmovdqa32 0x80(%r14, %r11), %zmm9
   804: 62d16d48efd1             vpxord %zmm9, %zmm2, %zmm2
   810: 62517d486f5f03           vmovdqa32 0xc0(%r15), %zmm11
   817: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   824: 62d15548efea             vpxord %zmm10, %zmm5, %zmm5
   830: 62517d486f5f04           vmovdqa32 0x100(%r15), %zmm11
   837: 6253b548ced300           vgf2p8affineqb $0, %zmm11, %zmm9, %zmm10
   844: 62513d48efc2             vpxord %zmm10, %zmm8, %zmm8
   850: 4d8b5028                 movq 0x28(%r8), %r10
   854: 62117d486f0c1a           vmovdqa32 (%r10, %r11), %zmm9
   861: 62d17d48efc1             vpxord %zmm9, %zmm0, %zmm0
   867: 62d16548efd9             vpxord %zmm9, %zmm3, %zmm3
   873: 62d14d48eff1             vpxord %zmm9, %zmm6, %zmm6
   879: 62117d486f4c1a01         vmovdqa32 0x40(%r10, %r11), %zmm9
   887: 62d17548efc9             vpxord %zmm9, %zmm1, %zmm1
   893: 62d15d48efe1             vpxord %zmm9, %zmm4, %zmm4
   899: 62d14548eff9             vpxord %zmm9, %zmm7, %zmm7
   905: 62117d486f4c1a02         vmovdqa32 0x80(%r10, %r11), %zmm9
   913: 62d16d48efd1             vpxord %zmm9, %zmm2, %zmm2
   919: 62d15548efe9             vpxord %zmm9, %zmm5, %zmm5
   925: 62513d48efc1             vpxord %zmm9, %zmm8, %zmm8
   931: 4d8b11                   movq (%r9), %r10
   934: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   941: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
   949: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
   957: 4d8b5108                 movq 8(%r9), %r10
   961: 62917d487f1c1a           vmovdqa32 %zmm3, (%r10, %r11)
   968: 62917d487f641a01         vmovdqa32 %zmm4, 0x40(%r10, %r11)
   976: 62917d487f6c1a02         vmovdqa32 %zmm5, 0x80(%r10, %r11)
   984: 4d8b5110                 movq 0x10(%r9), %r10
   988: 62917d487f341a           vmovdqa32 %zmm6, (%r10, %r11)
   995: 62917d487f7c1a01         vmovdqa32 %zmm7, 0x40(%r10, %r11)
  1003: 62117d487f441a02         vmovdqa32 %zmm8, 0x80(%r10, %r11)
  1011: 4981c3c0000000           addq $0xc0, %r11
  1018: 4c3bd8                   cmpq %rax, %r11
  1021: 0f863dfcffff             jbe 0x40
  1027: 4c3bd9                   cmpq %rcx, %r11
  1030: 0f83b2020000             jae 0x6be
  1036: 4153                     pushq %r11
  1038: 498dbf40010000           leaq 0x140(%r15), %rdi
  1045: 498db740020000           leaq 0x240(%r15), %rsi
  1052: 4d8b11                   movq (%r9), %r10
  1055: 4c8b1c24                 movq (%rsp), %r11
  1059: 33db                     xorl %ebx, %ebx
  1061: 498b00                   movq (%r8), %rax
  1064: 420fb61418               movzbl (%rax, %r11), %edx
  1069: 32da                     xorb %dl, %bl
  1071: 498b4008                 movq 8(%r8), %rax
  1075: 420fb61418               movzbl (%rax, %r11), %edx
  1080: 32da                     xorb %dl, %bl
  1082: 498b4010                 movq 0x10(%r8), %rax
  1086: 420fb61418               movzbl (%rax, %r11), %edx
  1091: 32da                     xorb %dl, %bl
  1093: 498b4018                 movq 0x18(%r8), %rax
  1097: 420fb61418               movzbl (%rax, %r11), %edx
  1102: 32da                     xorb %dl, %bl
  1104: 498b4020                 movq 0x20(%r8), %rax
  1108: 420fb61418               movzbl (%rax, %r11), %edx
  1113: 32da                     xorb %dl, %bl
  1115: 498b4028                 movq 0x28(%r8), %rax
  1119: 420fb61418               movzbl (%rax, %r11), %edx
  1124: 32da                     xorb %dl, %bl
  1126: 43881c1a                 movb %bl, (%r10, %r11)
  1130: 33db                     xorl %ebx, %ebx
  1132: 49ffc3                   incq %r11
  1135: 4c3bd9                   cmpq %rcx, %r11
  1138: 0f82adffffff             jb 0x425
  1144: 4d8b5108                 movq 8(%r9), %r10
  1148: 4c8b1c24                 movq (%rsp), %r11
  1152: 33db                     xorl %ebx, %ebx
  1154: 498b00                   movq (%r8), %rax
  1157: 420fb61418               movzbl (%rax, %r11), %edx
  1162: 4885d2                   testq %rdx, %rdx
  1165: 0f8422000000             je 0x4b5
  1171: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1175: 81c0e6000000             addl $0xe6, %eax
  1181: 81f8ff000000             cmpl $0xff, %eax
  1187: 0f8206000000             jb 0x4af
  1193: 81c001ffffff             addl $0xffffff01, %eax
  1199: 0fb60406                 movzbl (%rsi, %rax), %eax
  1203: 32d8                     xorb %al, %bl
  1205: 498b4008                 movq 8(%r8), %rax
  1209: 420fb61418               movzbl (%rax, %r11), %edx
  1214: 4885d2                   testq %rdx, %rdx
  1217: 0f8422000000             je 0x4e9
  1223: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1227: 81c0cd000000             addl $0xcd, %eax
  1233: 81f8ff000000             cmpl $0xff, %eax
  1239: 0f8206000000             jb 0x4e3
  1245: 81c001ffffff             addl $0xffffff01, %eax
  1251: 0fb60406                 movzbl (%rsi, %rax), %eax
  1255: 32d8                     xorb %al, %bl
  1257: 498b4010                 movq 0x10(%r8), %rax
  1261: 420fb61418               movzbl (%rax, %r11), %edx
  1266: 4885d2                   testq %rdx, %rdx
  1269: 0f841f000000             je 0x51a
  1275: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1279: 83c04b                   addl $0x4b, %eax
  1282: 81f8ff000000             cmpl $0xff, %eax
  1288: 0f8206000000             jb 0x514
  1294: 81c001ffffff             addl $0xffffff01, %eax
  1300: 0fb60406                 movzbl (%rsi, %rax), %eax
  1304: 32d8                     xorb %al, %bl
  1306: 498b4018                 movq 0x18(%r8), %rax
  1310: 420fb61418               movzbl (%rax, %r11), %edx
  1315: 4885d2                   testq %rdx, %rdx
  1318: 0f841f000000             je 0x54b
  1324: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1328: 83c032                   addl $0x32, %eax
  1331: 81f8ff000000             cmpl $0xff, %eax
  1337: 0f8206000000             jb 0x545
  1343: 81c001ffffff             addl $0xffffff01, %eax
  1349: 0fb60406                 movzbl (%rsi, %rax), %eax
  1353: 32d8                     xorb %al, %bl
  1355: 498b4020                 movq 0x20(%r8), %rax
  1359: 420fb61418               movzbl (%rax, %r11), %edx
  1364: 4885d2                   testq %rdx, %rdx
  1367: 0f841f000000             je 0x57c
  1373: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1377: 83c019                   addl $0x19, %eax
  1380: 81f8ff000000             cmpl $0xff, %eax
  1386: 0f8206000000             jb 0x576
  1392: 81c001ffffff             addl $0xffffff01, %eax
  1398: 0fb60406                 movzbl (%rsi, %rax), %eax
  1402: 32d8                     xorb %al, %bl
  1404: 498b4028                 movq 0x28(%r8), %rax
  1408: 420fb61418               movzbl (%rax, %r11), %edx
  1413: 32da                     xorb %dl, %bl
  1415: 43881c1a                 movb %bl, (%r10, %r11)
  1419: 33db                     xorl %ebx, %ebx
  1421: 49ffc3                   incq %r11
  1424: 4c3bd9                   cmpq %rcx, %r11
  1427: 0f82e9feffff             jb 0x482
  1433: 4d8b5110                 movq 0x10(%r9), %r10
  1437: 4c8b1c24                 movq (%rsp), %r11
  1441: 33db                     xorl %ebx, %ebx
  1443: 498b00                   movq (%r8), %rax
  1446: 420fb61418               movzbl (%rax, %r11), %edx
  1451: 4885d2                   testq %rdx, %rdx
  1454: 0f8422000000             je 0x5d6
  1460: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1464: 81c0cd000000             addl $0xcd, %eax
  1470: 81f8ff000000             cmpl $0xff, %eax
  1476: 0f8206000000             jb 0x5d0
  1482: 81c001ffffff             addl $0xffffff01, %eax
  1488: 0fb60406                 movzbl (%rsi, %rax), %eax
  1492: 32d8                     xorb %al, %bl
  1494: 498b4008                 movq 8(%r8), %rax
  1498: 420fb61418               movzbl (%rax, %r11), %edx
  1503: 4885d2                   testq %rdx, %rdx
  1506: 0f841f000000             je 0x607
  1512: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1516: 83c04b                   addl $0x4b, %eax
  1519: 81f8ff000000             cmpl $0xff, %eax
  1525: 0f8206000000             jb 0x601
  1531: 81c001ffffff             addl $0xffffff01, %eax
  1537: 0fb60406                 movzbl (%rsi, %rax), %eax
  1541: 32d8                     xorb %al, %bl
  1543: 498b4010                 movq 0x10(%r8), %rax
  1547: 420fb61418               movzbl (%rax, %r11), %edx
  1552: 4885d2                   testq %rdx, %rdx
  1555: 0f841f000000             je 0x638
  1561: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1565: 83c019                   addl $0x19, %eax
  1568: 81f8ff000000             cmpl $0xff, %eax
  1574: 0f8206000000             jb 0x632
  1580: 81c001ffffff             addl $0xffffff01, %eax
  1586: 0fb60406                 movzbl (%rsi, %rax), %eax
  1590: 32d8                     xorb %al, %bl
  1592: 498b4018                 movq 0x18(%r8), %rax
  1596: 420fb61418               movzbl (%rax, %r11), %edx
  1601: 4885d2                   testq %rdx, %rdx
  1604: 0f8422000000             je 0x66c
  1610: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1614: 81c0e6000000             addl $0xe6, %eax
  1620: 81f8ff000000             cmpl $0xff, %eax
  1626: 0f8206000000             jb 0x666
  1632: 81c001ffffff             addl $0xffffff01, %eax
  1638: 0fb60406                 movzbl (%rsi, %rax), %eax
  1642: 32d8                     xorb %al, %bl
  1644: 498b4020                 movq 0x20(%r8), %rax
  1648: 420fb61418               movzbl (%rax, %r11), %edx
  1653: 4885d2                   testq %rdx, %rdx
  1656: 0f841f000000             je 0x69d
  1662: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1666: 83c032                   addl $0x32, %eax
  1669: 81f8ff000000             cmpl $0xff, %eax
  1675: 0f8206000000             jb 0x697
  1681: 81c001ffffff             addl $0xffffff01, %eax
  1687: 0fb60406                 movzbl (%rsi, %rax), %eax
  1691: 32d8                     xorb %al, %bl
  1693: 498b4028                 movq 0x28(%r8), %rax
  1697: 420fb61418               movzbl (%rax, %r11), %edx
  1702: 32da                     xorb %dl, %bl
  1704: 43881c1a                 movb %bl, (%r10, %r11)
  1708: 33db                     xorl %ebx, %ebx
  1710: 49ffc3                   incq %r11
  1713: 4c3bd9                   cmpq %rcx, %r11
  1716: 0f82e9feffff             jb 0x5a3
  1722: 4883c408                 addq $8, %rsp
  1726: 415e                     popq %r14
  1728: 415d                     popq %r13
  1730: 415c                     popq %r12
  1732: 415f                     popq %r15
  1734: 5b                       popq %rbx
  1735: c5f877                   vzeroupper 
  1738: c3                       retq 
  1739: 0f0b                     ud2 
  1741: 0f0b                     ud2 
  1743: 0f0b                     ud2 
  1745: 0f0b                     ud2 
  1747: 0f0b                     ud2 
  1749: 0f0b                     ud2 
  1751: 0f0b                     ud2 
  1753: 0f0b                     ud2 
  1755: 0f0b                     ud2 
  1757: 0f0b                     ud2 
  1759: 0f0b                     ud2 
  1761: 0f0b                     ud2 
  1763: 0f0b                     ud2 
  1765: 0f0b                     ud2 
  1767: 0f0b                     ud2 
  1769: 0f0b                     ud2 
  1771: 0f0b                     ud2 
  1773: 0f0b                     ud2 
  1775: 0f0b                     ud2 
  1777: 0f0b                     ud2 
  1779: 0f0b                     ud2 
  1781: 0f0b                     ud2 
  1783: 0f0b                     ud2 
  1785: 0f0b                     ud2 
  1787: 0f0b                     ud2 
  1789: 0f0b                     ud2 
  1791: cc                       int3 
