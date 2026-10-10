; config_t1_54 target=avx512-gfni twin=ternlog topo=rdp(7)/enc op=encode U=1
; code_size=2560 twins=["ternlog"] uops_per_byte=2.140625
     0: 53                       pushq %rbx
     1: 4154                     pushq %r12
     3: 4155                     pushq %r13
     5: 4156                     pushq %r14
     7: 4c8bc7                   movq %rdi, %r8
    10: 4c8bce                   movq %rsi, %r9
    13: 488bca                   movq %rdx, %rcx
    16: 498b30                   movq (%r8), %rsi
    19: 498b7808                 movq 8(%r8), %rdi
    23: 4d8b6010                 movq 0x10(%r8), %r12
    27: 4d8b6818                 movq 0x18(%r8), %r13
    31: 4d8b7020                 movq 0x20(%r8), %r14
    35: 4d33db                   xorq %r11, %r11
    38: 488bc1                   movq %rcx, %rax
    41: 4883e840                 subq $0x40, %rax
    45: 0f82ed030000             jb 0x420
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
    70: 62f17548efc9             vpxord %zmm1, %zmm1, %zmm1
    76: 62f16d48efd2             vpxord %zmm2, %zmm2, %zmm2
    82: 62f16548efdb             vpxord %zmm3, %zmm3, %zmm3
    88: 62f15d48efe4             vpxord %zmm4, %zmm4, %zmm4
    94: 62f15548efed             vpxord %zmm5, %zmm5, %zmm5
   100: 62f14d48eff6             vpxord %zmm6, %zmm6, %zmm6
   106: 62f14548efff             vpxord %zmm7, %zmm7, %zmm7
   112: 62513d48efc0             vpxord %zmm8, %zmm8, %zmm8
   118: 62513548efc9             vpxord %zmm9, %zmm9, %zmm9
   124: 62512d48efd2             vpxord %zmm10, %zmm10, %zmm10
   130: 62512548efdb             vpxord %zmm11, %zmm11, %zmm11
   136: 498b5038                 movq 0x38(%r8), %rdx
   140: 62317d486f241f           vmovdqa32 (%rdi, %r11), %zmm12
   147: 62b31d48250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm1
   155: 62b31d4825341a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm6
   163: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   169: 62313d48ef041a           vpxord (%rdx, %r11), %zmm8, %zmm8
   176: 498b5040                 movq 0x40(%r8), %rdx
   180: 62117d486f241c           vmovdqa32 (%r12, %r11), %zmm12
   187: 62b31d4825141a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm2
   195: 62b31d48253c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm7
   203: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   209: 62313548ef0c1a           vpxord (%rdx, %r11), %zmm9, %zmm9
   216: 498b5048                 movq 0x48(%r8), %rdx
   220: 62117d486f641d00         vmovdqa32 (%r13, %r11), %zmm12
   228: 62b31d48251c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm3
   236: 62331d4825041a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm8
   244: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   250: 62312d48ef141a           vpxord (%rdx, %r11), %zmm10, %zmm10
   257: 498b5050                 movq 0x50(%r8), %rdx
   261: 62117d486f241e           vmovdqa32 (%r14, %r11), %zmm12
   268: 62b31d4825241a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm4
   276: 62331d48250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm9
   284: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   290: 62312548ef1c1a           vpxord (%rdx, %r11), %zmm11, %zmm11
   297: 4d8b5028                 movq 0x28(%r8), %r10
   301: 498b5058                 movq 0x58(%r8), %rdx
   305: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   312: 62b31d48252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm5
   320: 62331d4825141a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm10
   328: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   334: 4d8b5068                 movq 0x68(%r8), %r10
   338: 498b9098000000           movq 0x98(%r8), %rdx
   345: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   352: 62b31d48250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm1
   360: 62b31d4825341a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm6
   368: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   374: 62312d48ef141a           vpxord (%rdx, %r11), %zmm10, %zmm10
   381: 4d8b5070                 movq 0x70(%r8), %r10
   385: 498b90a0000000           movq 0xa0(%r8), %rdx
   392: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   399: 62b31d4825141a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm2
   407: 62b31d48253c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm7
   415: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   421: 62312548ef1c1a           vpxord (%rdx, %r11), %zmm11, %zmm11
   428: 4d8b5078                 movq 0x78(%r8), %r10
   432: 498b90a8000000           movq 0xa8(%r8), %rdx
   439: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   446: 62b31d48251c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm3
   454: 62331d4825041a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm8
   462: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   468: 4d8b9080000000           movq 0x80(%r8), %r10
   475: 498b90b0000000           movq 0xb0(%r8), %rdx
   482: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   489: 62b31d4825241a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm4
   497: 62b14d48ef341a           vpxord (%rdx, %r11), %zmm6, %zmm6
   504: 62331d48250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm9
   512: 4d8b9088000000           movq 0x88(%r8), %r10
   519: 498b90b8000000           movq 0xb8(%r8), %rdx
   526: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   533: 62b31d48252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm5
   541: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   547: 62b14548ef3c1a           vpxord (%rdx, %r11), %zmm7, %zmm7
   554: 62331d4825141a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm10
   562: 4d8b90c8000000           movq 0xc8(%r8), %r10
   569: 498b90f8000000           movq 0xf8(%r8), %rdx
   576: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   583: 62b31d48250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm1
   591: 62b31d4825341a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm6
   599: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   605: 4d8b90d0000000           movq 0xd0(%r8), %r10
   612: 498b9000010000           movq 0x100(%r8), %rdx
   619: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   626: 62b31d4825141a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm2
   634: 62b14d48ef341a           vpxord (%rdx, %r11), %zmm6, %zmm6
   641: 62b31d48253c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm7
   649: 4d8b90d8000000           movq 0xd8(%r8), %r10
   656: 498b9008010000           movq 0x108(%r8), %rdx
   663: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   670: 62b31d48251c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm3
   678: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   684: 62b14548ef3c1a           vpxord (%rdx, %r11), %zmm7, %zmm7
   691: 62331d4825041a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm8
   699: 4d8b90e0000000           movq 0xe0(%r8), %r10
   706: 498b9010010000           movq 0x110(%r8), %rdx
   713: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   720: 62b31d4825241a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm4
   728: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   734: 62313d48ef041a           vpxord (%rdx, %r11), %zmm8, %zmm8
   741: 62331d48250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm9
   749: 4d8b90e8000000           movq 0xe8(%r8), %r10
   756: 498b9018010000           movq 0x118(%r8), %rdx
   763: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   770: 62b31d48252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm5
   778: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   784: 62313548ef0c1a           vpxord (%rdx, %r11), %zmm9, %zmm9
   791: 62331d4825141a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm10
   799: 498b5030                 movq 0x30(%r8), %rdx
   803: 62317d486f241e           vmovdqa32 (%rsi, %r11), %zmm12
   810: 62b31d4825041a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm0
   818: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   824: 62b14548ef3c1a           vpxord (%rdx, %r11), %zmm7, %zmm7
   831: 4d8b5060                 movq 0x60(%r8), %r10
   835: 498b9090000000           movq 0x90(%r8), %rdx
   842: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   849: 62b31d4825041a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm0
   857: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   863: 62313548ef0c1a           vpxord (%rdx, %r11), %zmm9, %zmm9
   870: 4d8b90c0000000           movq 0xc0(%r8), %r10
   877: 498b90f0000000           movq 0xf0(%r8), %rdx
   884: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   891: 62b31d4825041a96         vpternlogd $0x96, (%rdx, %r11), %zmm12, %zmm0
   899: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   905: 62312548ef1c1a           vpxord (%rdx, %r11), %zmm11, %zmm11
   912: 4d8b11                   movq (%r9), %r10
   915: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   922: 4d8b5108                 movq 8(%r9), %r10
   926: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
   933: 4d8b5110                 movq 0x10(%r9), %r10
   937: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
   944: 4d8b5118                 movq 0x18(%r9), %r10
   948: 62917d487f1c1a           vmovdqa32 %zmm3, (%r10, %r11)
   955: 4d8b5120                 movq 0x20(%r9), %r10
   959: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
   966: 4d8b5128                 movq 0x28(%r9), %r10
   970: 62917d487f2c1a           vmovdqa32 %zmm5, (%r10, %r11)
   977: 4d8b5130                 movq 0x30(%r9), %r10
   981: 62917d487f341a           vmovdqa32 %zmm6, (%r10, %r11)
   988: 4d8b5138                 movq 0x38(%r9), %r10
   992: 62917d487f3c1a           vmovdqa32 %zmm7, (%r10, %r11)
   999: 4d8b5140                 movq 0x40(%r9), %r10
  1003: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  1010: 4d8b5148                 movq 0x48(%r9), %r10
  1014: 62117d487f0c1a           vmovdqa32 %zmm9, (%r10, %r11)
  1021: 4d8b5150                 movq 0x50(%r9), %r10
  1025: 62117d487f141a           vmovdqa32 %zmm10, (%r10, %r11)
  1032: 4d8b5158                 movq 0x58(%r9), %r10
  1036: 62117d487f1c1a           vmovdqa32 %zmm11, (%r10, %r11)
  1043: 4983c340                 addq $0x40, %r11
  1047: 4c3bd8                   cmpq %rax, %r11
  1050: 0f8620fcffff             jbe 0x40
  1056: 4c3bd9                   cmpq %rcx, %r11
  1059: 0f83a4050000             jae 0x9cd
  1065: 4153                     pushq %r11
  1067: 4d8b11                   movq (%r9), %r10
  1070: 4c8b1c24                 movq (%rsp), %r11
  1074: 498b00                   movq (%r8), %rax
  1077: 428a1418                 movb (%rax, %r11), %dl
  1081: 8ada                     movb %dl, %bl
  1083: 498b4030                 movq 0x30(%r8), %rax
  1087: 428a1418                 movb (%rax, %r11), %dl
  1091: 32da                     xorb %dl, %bl
  1093: 498b4060                 movq 0x60(%r8), %rax
  1097: 428a1418                 movb (%rax, %r11), %dl
  1101: 32da                     xorb %dl, %bl
  1103: 498b8090000000           movq 0x90(%r8), %rax
  1110: 428a1418                 movb (%rax, %r11), %dl
  1114: 32da                     xorb %dl, %bl
  1116: 498b80c0000000           movq 0xc0(%r8), %rax
  1123: 428a1418                 movb (%rax, %r11), %dl
  1127: 32da                     xorb %dl, %bl
  1129: 498b80f0000000           movq 0xf0(%r8), %rax
  1136: 428a1418                 movb (%rax, %r11), %dl
  1140: 32da                     xorb %dl, %bl
  1142: 43881c1a                 movb %bl, (%r10, %r11)
  1146: 33db                     xorl %ebx, %ebx
  1148: 49ffc3                   incq %r11
  1151: 4c3bd9                   cmpq %rcx, %r11
  1154: 0f82aaffffff             jb 0x432
  1160: 4d8b5108                 movq 8(%r9), %r10
  1164: 4c8b1c24                 movq (%rsp), %r11
  1168: 498b4008                 movq 8(%r8), %rax
  1172: 428a1418                 movb (%rax, %r11), %dl
  1176: 8ada                     movb %dl, %bl
  1178: 498b4038                 movq 0x38(%r8), %rax
  1182: 428a1418                 movb (%rax, %r11), %dl
  1186: 32da                     xorb %dl, %bl
  1188: 498b4068                 movq 0x68(%r8), %rax
  1192: 428a1418                 movb (%rax, %r11), %dl
  1196: 32da                     xorb %dl, %bl
  1198: 498b8098000000           movq 0x98(%r8), %rax
  1205: 428a1418                 movb (%rax, %r11), %dl
  1209: 32da                     xorb %dl, %bl
  1211: 498b80c8000000           movq 0xc8(%r8), %rax
  1218: 428a1418                 movb (%rax, %r11), %dl
  1222: 32da                     xorb %dl, %bl
  1224: 498b80f8000000           movq 0xf8(%r8), %rax
  1231: 428a1418                 movb (%rax, %r11), %dl
  1235: 32da                     xorb %dl, %bl
  1237: 43881c1a                 movb %bl, (%r10, %r11)
  1241: 33db                     xorl %ebx, %ebx
  1243: 49ffc3                   incq %r11
  1246: 4c3bd9                   cmpq %rcx, %r11
  1249: 0f82a9ffffff             jb 0x490
  1255: 4d8b5110                 movq 0x10(%r9), %r10
  1259: 4c8b1c24                 movq (%rsp), %r11
  1263: 498b4010                 movq 0x10(%r8), %rax
  1267: 428a1418                 movb (%rax, %r11), %dl
  1271: 8ada                     movb %dl, %bl
  1273: 498b4040                 movq 0x40(%r8), %rax
  1277: 428a1418                 movb (%rax, %r11), %dl
  1281: 32da                     xorb %dl, %bl
  1283: 498b4070                 movq 0x70(%r8), %rax
  1287: 428a1418                 movb (%rax, %r11), %dl
  1291: 32da                     xorb %dl, %bl
  1293: 498b80a0000000           movq 0xa0(%r8), %rax
  1300: 428a1418                 movb (%rax, %r11), %dl
  1304: 32da                     xorb %dl, %bl
  1306: 498b80d0000000           movq 0xd0(%r8), %rax
  1313: 428a1418                 movb (%rax, %r11), %dl
  1317: 32da                     xorb %dl, %bl
  1319: 498b8000010000           movq 0x100(%r8), %rax
  1326: 428a1418                 movb (%rax, %r11), %dl
  1330: 32da                     xorb %dl, %bl
  1332: 43881c1a                 movb %bl, (%r10, %r11)
  1336: 33db                     xorl %ebx, %ebx
  1338: 49ffc3                   incq %r11
  1341: 4c3bd9                   cmpq %rcx, %r11
  1344: 0f82a9ffffff             jb 0x4ef
  1350: 4d8b5118                 movq 0x18(%r9), %r10
  1354: 4c8b1c24                 movq (%rsp), %r11
  1358: 498b4018                 movq 0x18(%r8), %rax
  1362: 428a1418                 movb (%rax, %r11), %dl
  1366: 8ada                     movb %dl, %bl
  1368: 498b4048                 movq 0x48(%r8), %rax
  1372: 428a1418                 movb (%rax, %r11), %dl
  1376: 32da                     xorb %dl, %bl
  1378: 498b4078                 movq 0x78(%r8), %rax
  1382: 428a1418                 movb (%rax, %r11), %dl
  1386: 32da                     xorb %dl, %bl
  1388: 498b80a8000000           movq 0xa8(%r8), %rax
  1395: 428a1418                 movb (%rax, %r11), %dl
  1399: 32da                     xorb %dl, %bl
  1401: 498b80d8000000           movq 0xd8(%r8), %rax
  1408: 428a1418                 movb (%rax, %r11), %dl
  1412: 32da                     xorb %dl, %bl
  1414: 498b8008010000           movq 0x108(%r8), %rax
  1421: 428a1418                 movb (%rax, %r11), %dl
  1425: 32da                     xorb %dl, %bl
  1427: 43881c1a                 movb %bl, (%r10, %r11)
  1431: 33db                     xorl %ebx, %ebx
  1433: 49ffc3                   incq %r11
  1436: 4c3bd9                   cmpq %rcx, %r11
  1439: 0f82a9ffffff             jb 0x54e
  1445: 4d8b5120                 movq 0x20(%r9), %r10
  1449: 4c8b1c24                 movq (%rsp), %r11
  1453: 498b4020                 movq 0x20(%r8), %rax
  1457: 428a1418                 movb (%rax, %r11), %dl
  1461: 8ada                     movb %dl, %bl
  1463: 498b4050                 movq 0x50(%r8), %rax
  1467: 428a1418                 movb (%rax, %r11), %dl
  1471: 32da                     xorb %dl, %bl
  1473: 498b8080000000           movq 0x80(%r8), %rax
  1480: 428a1418                 movb (%rax, %r11), %dl
  1484: 32da                     xorb %dl, %bl
  1486: 498b80b0000000           movq 0xb0(%r8), %rax
  1493: 428a1418                 movb (%rax, %r11), %dl
  1497: 32da                     xorb %dl, %bl
  1499: 498b80e0000000           movq 0xe0(%r8), %rax
  1506: 428a1418                 movb (%rax, %r11), %dl
  1510: 32da                     xorb %dl, %bl
  1512: 498b8010010000           movq 0x110(%r8), %rax
  1519: 428a1418                 movb (%rax, %r11), %dl
  1523: 32da                     xorb %dl, %bl
  1525: 43881c1a                 movb %bl, (%r10, %r11)
  1529: 33db                     xorl %ebx, %ebx
  1531: 49ffc3                   incq %r11
  1534: 4c3bd9                   cmpq %rcx, %r11
  1537: 0f82a6ffffff             jb 0x5ad
  1543: 4d8b5128                 movq 0x28(%r9), %r10
  1547: 4c8b1c24                 movq (%rsp), %r11
  1551: 498b4028                 movq 0x28(%r8), %rax
  1555: 428a1418                 movb (%rax, %r11), %dl
  1559: 8ada                     movb %dl, %bl
  1561: 498b4058                 movq 0x58(%r8), %rax
  1565: 428a1418                 movb (%rax, %r11), %dl
  1569: 32da                     xorb %dl, %bl
  1571: 498b8088000000           movq 0x88(%r8), %rax
  1578: 428a1418                 movb (%rax, %r11), %dl
  1582: 32da                     xorb %dl, %bl
  1584: 498b80b8000000           movq 0xb8(%r8), %rax
  1591: 428a1418                 movb (%rax, %r11), %dl
  1595: 32da                     xorb %dl, %bl
  1597: 498b80e8000000           movq 0xe8(%r8), %rax
  1604: 428a1418                 movb (%rax, %r11), %dl
  1608: 32da                     xorb %dl, %bl
  1610: 498b8018010000           movq 0x118(%r8), %rax
  1617: 428a1418                 movb (%rax, %r11), %dl
  1621: 32da                     xorb %dl, %bl
  1623: 43881c1a                 movb %bl, (%r10, %r11)
  1627: 33db                     xorl %ebx, %ebx
  1629: 49ffc3                   incq %r11
  1632: 4c3bd9                   cmpq %rcx, %r11
  1635: 0f82a6ffffff             jb 0x60f
  1641: 4d8b5130                 movq 0x30(%r9), %r10
  1645: 4c8b1c24                 movq (%rsp), %r11
  1649: 498b00                   movq (%r8), %rax
  1652: 428a1418                 movb (%rax, %r11), %dl
  1656: 8ada                     movb %dl, %bl
  1658: 498b4008                 movq 8(%r8), %rax
  1662: 428a1418                 movb (%rax, %r11), %dl
  1666: 32da                     xorb %dl, %bl
  1668: 498b4038                 movq 0x38(%r8), %rax
  1672: 428a1418                 movb (%rax, %r11), %dl
  1676: 32da                     xorb %dl, %bl
  1678: 498b4068                 movq 0x68(%r8), %rax
  1682: 428a1418                 movb (%rax, %r11), %dl
  1686: 32da                     xorb %dl, %bl
  1688: 498b8088000000           movq 0x88(%r8), %rax
  1695: 428a1418                 movb (%rax, %r11), %dl
  1699: 32da                     xorb %dl, %bl
  1701: 498b8098000000           movq 0x98(%r8), %rax
  1708: 428a1418                 movb (%rax, %r11), %dl
  1712: 32da                     xorb %dl, %bl
  1714: 498b80b0000000           movq 0xb0(%r8), %rax
  1721: 428a1418                 movb (%rax, %r11), %dl
  1725: 32da                     xorb %dl, %bl
  1727: 498b80c8000000           movq 0xc8(%r8), %rax
  1734: 428a1418                 movb (%rax, %r11), %dl
  1738: 32da                     xorb %dl, %bl
  1740: 498b80d8000000           movq 0xd8(%r8), %rax
  1747: 428a1418                 movb (%rax, %r11), %dl
  1751: 32da                     xorb %dl, %bl
  1753: 498b80f8000000           movq 0xf8(%r8), %rax
  1760: 428a1418                 movb (%rax, %r11), %dl
  1764: 32da                     xorb %dl, %bl
  1766: 498b8000010000           movq 0x100(%r8), %rax
  1773: 428a1418                 movb (%rax, %r11), %dl
  1777: 32da                     xorb %dl, %bl
  1779: 43881c1a                 movb %bl, (%r10, %r11)
  1783: 33db                     xorl %ebx, %ebx
  1785: 49ffc3                   incq %r11
  1788: 4c3bd9                   cmpq %rcx, %r11
  1791: 0f826cffffff             jb 0x671
  1797: 4d8b5138                 movq 0x38(%r9), %r10
  1801: 4c8b1c24                 movq (%rsp), %r11
  1805: 498b4008                 movq 8(%r8), %rax
  1809: 428a1418                 movb (%rax, %r11), %dl
  1813: 8ada                     movb %dl, %bl
  1815: 498b4010                 movq 0x10(%r8), %rax
  1819: 428a1418                 movb (%rax, %r11), %dl
  1823: 32da                     xorb %dl, %bl
  1825: 498b4030                 movq 0x30(%r8), %rax
  1829: 428a1418                 movb (%rax, %r11), %dl
  1833: 32da                     xorb %dl, %bl
  1835: 498b4040                 movq 0x40(%r8), %rax
  1839: 428a1418                 movb (%rax, %r11), %dl
  1843: 32da                     xorb %dl, %bl
  1845: 498b4070                 movq 0x70(%r8), %rax
  1849: 428a1418                 movb (%rax, %r11), %dl
  1853: 32da                     xorb %dl, %bl
  1855: 498b80a0000000           movq 0xa0(%r8), %rax
  1862: 428a1418                 movb (%rax, %r11), %dl
  1866: 32da                     xorb %dl, %bl
  1868: 498b80b8000000           movq 0xb8(%r8), %rax
  1875: 428a1418                 movb (%rax, %r11), %dl
  1879: 32da                     xorb %dl, %bl
  1881: 498b80d0000000           movq 0xd0(%r8), %rax
  1888: 428a1418                 movb (%rax, %r11), %dl
  1892: 32da                     xorb %dl, %bl
  1894: 498b80e0000000           movq 0xe0(%r8), %rax
  1901: 428a1418                 movb (%rax, %r11), %dl
  1905: 32da                     xorb %dl, %bl
  1907: 498b8000010000           movq 0x100(%r8), %rax
  1914: 428a1418                 movb (%rax, %r11), %dl
  1918: 32da                     xorb %dl, %bl
  1920: 498b8008010000           movq 0x108(%r8), %rax
  1927: 428a1418                 movb (%rax, %r11), %dl
  1931: 32da                     xorb %dl, %bl
  1933: 43881c1a                 movb %bl, (%r10, %r11)
  1937: 33db                     xorl %ebx, %ebx
  1939: 49ffc3                   incq %r11
  1942: 4c3bd9                   cmpq %rcx, %r11
  1945: 0f826effffff             jb 0x70d
  1951: 4d8b5140                 movq 0x40(%r9), %r10
  1955: 4c8b1c24                 movq (%rsp), %r11
  1959: 498b4010                 movq 0x10(%r8), %rax
  1963: 428a1418                 movb (%rax, %r11), %dl
  1967: 8ada                     movb %dl, %bl
  1969: 498b4018                 movq 0x18(%r8), %rax
  1973: 428a1418                 movb (%rax, %r11), %dl
  1977: 32da                     xorb %dl, %bl
  1979: 498b4038                 movq 0x38(%r8), %rax
  1983: 428a1418                 movb (%rax, %r11), %dl
  1987: 32da                     xorb %dl, %bl
  1989: 498b4048                 movq 0x48(%r8), %rax
  1993: 428a1418                 movb (%rax, %r11), %dl
  1997: 32da                     xorb %dl, %bl
  1999: 498b4060                 movq 0x60(%r8), %rax
  2003: 428a1418                 movb (%rax, %r11), %dl
  2007: 32da                     xorb %dl, %bl
  2009: 498b4078                 movq 0x78(%r8), %rax
  2013: 428a1418                 movb (%rax, %r11), %dl
  2017: 32da                     xorb %dl, %bl
  2019: 498b80a8000000           movq 0xa8(%r8), %rax
  2026: 428a1418                 movb (%rax, %r11), %dl
  2030: 32da                     xorb %dl, %bl
  2032: 498b80d8000000           movq 0xd8(%r8), %rax
  2039: 428a1418                 movb (%rax, %r11), %dl
  2043: 32da                     xorb %dl, %bl
  2045: 498b80e8000000           movq 0xe8(%r8), %rax
  2052: 428a1418                 movb (%rax, %r11), %dl
  2056: 32da                     xorb %dl, %bl
  2058: 498b8008010000           movq 0x108(%r8), %rax
  2065: 428a1418                 movb (%rax, %r11), %dl
  2069: 32da                     xorb %dl, %bl
  2071: 498b8010010000           movq 0x110(%r8), %rax
  2078: 428a1418                 movb (%rax, %r11), %dl
  2082: 32da                     xorb %dl, %bl
  2084: 43881c1a                 movb %bl, (%r10, %r11)
  2088: 33db                     xorl %ebx, %ebx
  2090: 49ffc3                   incq %r11
  2093: 4c3bd9                   cmpq %rcx, %r11
  2096: 0f8271ffffff             jb 0x7a7
  2102: 4d8b5148                 movq 0x48(%r9), %r10
  2106: 4c8b1c24                 movq (%rsp), %r11
  2110: 498b4018                 movq 0x18(%r8), %rax
  2114: 428a1418                 movb (%rax, %r11), %dl
  2118: 8ada                     movb %dl, %bl
  2120: 498b4020                 movq 0x20(%r8), %rax
  2124: 428a1418                 movb (%rax, %r11), %dl
  2128: 32da                     xorb %dl, %bl
  2130: 498b4040                 movq 0x40(%r8), %rax
  2134: 428a1418                 movb (%rax, %r11), %dl
  2138: 32da                     xorb %dl, %bl
  2140: 498b4050                 movq 0x50(%r8), %rax
  2144: 428a1418                 movb (%rax, %r11), %dl
  2148: 32da                     xorb %dl, %bl
  2150: 498b4068                 movq 0x68(%r8), %rax
  2154: 428a1418                 movb (%rax, %r11), %dl
  2158: 32da                     xorb %dl, %bl
  2160: 498b8080000000           movq 0x80(%r8), %rax
  2167: 428a1418                 movb (%rax, %r11), %dl
  2171: 32da                     xorb %dl, %bl
  2173: 498b8090000000           movq 0x90(%r8), %rax
  2180: 428a1418                 movb (%rax, %r11), %dl
  2184: 32da                     xorb %dl, %bl
  2186: 498b80b0000000           movq 0xb0(%r8), %rax
  2193: 428a1418                 movb (%rax, %r11), %dl
  2197: 32da                     xorb %dl, %bl
  2199: 498b80e0000000           movq 0xe0(%r8), %rax
  2206: 428a1418                 movb (%rax, %r11), %dl
  2210: 32da                     xorb %dl, %bl
  2212: 498b8010010000           movq 0x110(%r8), %rax
  2219: 428a1418                 movb (%rax, %r11), %dl
  2223: 32da                     xorb %dl, %bl
  2225: 498b8018010000           movq 0x118(%r8), %rax
  2232: 428a1418                 movb (%rax, %r11), %dl
  2236: 32da                     xorb %dl, %bl
  2238: 43881c1a                 movb %bl, (%r10, %r11)
  2242: 33db                     xorl %ebx, %ebx
  2244: 49ffc3                   incq %r11
  2247: 4c3bd9                   cmpq %rcx, %r11
  2250: 0f826effffff             jb 0x83e
  2256: 4d8b5150                 movq 0x50(%r9), %r10
  2260: 4c8b1c24                 movq (%rsp), %r11
  2264: 498b4020                 movq 0x20(%r8), %rax
  2268: 428a1418                 movb (%rax, %r11), %dl
  2272: 8ada                     movb %dl, %bl
  2274: 498b4028                 movq 0x28(%r8), %rax
  2278: 428a1418                 movb (%rax, %r11), %dl
  2282: 32da                     xorb %dl, %bl
  2284: 498b4048                 movq 0x48(%r8), %rax
  2288: 428a1418                 movb (%rax, %r11), %dl
  2292: 32da                     xorb %dl, %bl
  2294: 498b4058                 movq 0x58(%r8), %rax
  2298: 428a1418                 movb (%rax, %r11), %dl
  2302: 32da                     xorb %dl, %bl
  2304: 498b4070                 movq 0x70(%r8), %rax
  2308: 428a1418                 movb (%rax, %r11), %dl
  2312: 32da                     xorb %dl, %bl
  2314: 498b8088000000           movq 0x88(%r8), %rax
  2321: 428a1418                 movb (%rax, %r11), %dl
  2325: 32da                     xorb %dl, %bl
  2327: 498b8098000000           movq 0x98(%r8), %rax
  2334: 428a1418                 movb (%rax, %r11), %dl
  2338: 32da                     xorb %dl, %bl
  2340: 498b80b8000000           movq 0xb8(%r8), %rax
  2347: 428a1418                 movb (%rax, %r11), %dl
  2351: 32da                     xorb %dl, %bl
  2353: 498b80c0000000           movq 0xc0(%r8), %rax
  2360: 428a1418                 movb (%rax, %r11), %dl
  2364: 32da                     xorb %dl, %bl
  2366: 498b80e8000000           movq 0xe8(%r8), %rax
  2373: 428a1418                 movb (%rax, %r11), %dl
  2377: 32da                     xorb %dl, %bl
  2379: 498b8018010000           movq 0x118(%r8), %rax
  2386: 428a1418                 movb (%rax, %r11), %dl
  2390: 32da                     xorb %dl, %bl
  2392: 43881c1a                 movb %bl, (%r10, %r11)
  2396: 33db                     xorl %ebx, %ebx
  2398: 49ffc3                   incq %r11
  2401: 4c3bd9                   cmpq %rcx, %r11
  2404: 0f826effffff             jb 0x8d8
  2410: 4d8b5158                 movq 0x58(%r9), %r10
  2414: 4c8b1c24                 movq (%rsp), %r11
  2418: 498b4028                 movq 0x28(%r8), %rax
  2422: 428a1418                 movb (%rax, %r11), %dl
  2426: 8ada                     movb %dl, %bl
  2428: 498b4050                 movq 0x50(%r8), %rax
  2432: 428a1418                 movb (%rax, %r11), %dl
  2436: 32da                     xorb %dl, %bl
  2438: 498b4078                 movq 0x78(%r8), %rax
  2442: 428a1418                 movb (%rax, %r11), %dl
  2446: 32da                     xorb %dl, %bl
  2448: 498b80a0000000           movq 0xa0(%r8), %rax
  2455: 428a1418                 movb (%rax, %r11), %dl
  2459: 32da                     xorb %dl, %bl
  2461: 498b80c8000000           movq 0xc8(%r8), %rax
  2468: 428a1418                 movb (%rax, %r11), %dl
  2472: 32da                     xorb %dl, %bl
  2474: 498b80f0000000           movq 0xf0(%r8), %rax
  2481: 428a1418                 movb (%rax, %r11), %dl
  2485: 32da                     xorb %dl, %bl
  2487: 43881c1a                 movb %bl, (%r10, %r11)
  2491: 33db                     xorl %ebx, %ebx
  2493: 49ffc3                   incq %r11
  2496: 4c3bd9                   cmpq %rcx, %r11
  2499: 0f82a9ffffff             jb 0x972
  2505: 4883c408                 addq $8, %rsp
  2509: 415e                     popq %r14
  2511: 415d                     popq %r13
  2513: 415c                     popq %r12
  2515: 5b                       popq %rbx
  2516: c5f877                   vzeroupper 
  2519: c3                       retq 
  2520: 0f0b                     ud2 
  2522: 0f0b                     ud2 
  2524: 0f0b                     ud2 
  2526: 0f0b                     ud2 
  2528: 0f0b                     ud2 
  2530: 0f0b                     ud2 
  2532: 0f0b                     ud2 
  2534: 0f0b                     ud2 
  2536: 0f0b                     ud2 
  2538: 0f0b                     ud2 
  2540: 0f0b                     ud2 
  2542: 0f0b                     ud2 
  2544: 0f0b                     ud2 
  2546: 0f0b                     ud2 
  2548: 0f0b                     ud2 
  2550: 0f0b                     ud2 
  2552: 0f0b                     ud2 
  2554: 0f0b                     ud2 
  2556: 0f0b                     ud2 
  2558: 0f0b                     ud2 
