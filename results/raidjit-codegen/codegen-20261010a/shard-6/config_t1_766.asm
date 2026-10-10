; config_t1_766 target=sse twin=rec_cf topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1344 twins=[] uops_per_byte=5.9375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d27050000           leaq 0x527(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66450f6f07               movdqa (%r15), %xmm8
    52: 488bc1                   movq %rcx, %rax
    55: 4883e810                 subq $0x10, %rax
    59: 0f8226020000             jb 0x267
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66420f6f141e             movdqa (%rsi, %r11), %xmm2
   102: 660f6fda                 movdqa %xmm2, %xmm3
   106: 66410fdbd8               pand %xmm8, %xmm3
   111: 660f6fe2                 movdqa %xmm2, %xmm4
   115: 660f71d404               psrlw $4, %xmm4
   120: 66410fdbe0               pand %xmm8, %xmm4
   125: 660f6fc2                 movdqa %xmm2, %xmm0
   129: 66410f6f6f40             movdqa 0x40(%r15), %xmm5
   135: 660f3800eb               pshufb %xmm3, %xmm5
   140: 66410f6fb780000000       movdqa 0x80(%r15), %xmm6
   149: 660f3800f4               pshufb %xmm4, %xmm6
   154: 660f6fcd                 movdqa %xmm5, %xmm1
   158: 660fefce                 pxor %xmm6, %xmm1
   162: 66420f6f141f             movdqa (%rdi, %r11), %xmm2
   168: 660f6fda                 movdqa %xmm2, %xmm3
   172: 66410fdbd8               pand %xmm8, %xmm3
   177: 660f6fe2                 movdqa %xmm2, %xmm4
   181: 660f71d404               psrlw $4, %xmm4
   186: 66410fdbe0               pand %xmm8, %xmm4
   191: 660fefc2                 pxor %xmm2, %xmm0
   195: 66410f6fafc0000000       movdqa 0xc0(%r15), %xmm5
   204: 660f3800eb               pshufb %xmm3, %xmm5
   209: 66410f6fb700010000       movdqa 0x100(%r15), %xmm6
   218: 660f3800f4               pshufb %xmm4, %xmm6
   223: 660fefcd                 pxor %xmm5, %xmm1
   227: 660fefce                 pxor %xmm6, %xmm1
   231: 66430f6f141c             movdqa (%r12, %r11), %xmm2
   237: 660f6fda                 movdqa %xmm2, %xmm3
   241: 66410fdbd8               pand %xmm8, %xmm3
   246: 660f6fe2                 movdqa %xmm2, %xmm4
   250: 660f71d404               psrlw $4, %xmm4
   255: 66410fdbe0               pand %xmm8, %xmm4
   260: 660fefc2                 pxor %xmm2, %xmm0
   264: 66410f6faf40010000       movdqa 0x140(%r15), %xmm5
   273: 660f3800eb               pshufb %xmm3, %xmm5
   278: 66410f6fb780010000       movdqa 0x180(%r15), %xmm6
   287: 660f3800f4               pshufb %xmm4, %xmm6
   292: 660fefcd                 pxor %xmm5, %xmm1
   296: 660fefce                 pxor %xmm6, %xmm1
   300: 66430f6f541d00           movdqa (%r13, %r11), %xmm2
   307: 660fefc2                 pxor %xmm2, %xmm0
   311: 660fefca                 pxor %xmm2, %xmm1
   315: 66430f6f141e             movdqa (%r14, %r11), %xmm2
   321: 660fefc2                 pxor %xmm2, %xmm0
   325: 4d8b5028                 movq 0x28(%r8), %r10
   329: 66430f6f141a             movdqa (%r10, %r11), %xmm2
   335: 660fefca                 pxor %xmm2, %xmm1
   339: 660f6fd8                 movdqa %xmm0, %xmm3
   343: 66410fdbd8               pand %xmm8, %xmm3
   348: 660f6fe0                 movdqa %xmm0, %xmm4
   352: 660f71d404               psrlw $4, %xmm4
   357: 66410fdbe0               pand %xmm8, %xmm4
   362: 66410f6fafc0010000       movdqa 0x1c0(%r15), %xmm5
   371: 660f3800eb               pshufb %xmm3, %xmm5
   376: 66410f6fb700020000       movdqa 0x200(%r15), %xmm6
   385: 660f3800f4               pshufb %xmm4, %xmm6
   390: 660f6ffd                 movdqa %xmm5, %xmm7
   394: 660feffe                 pxor %xmm6, %xmm7
   398: 660feff9                 pxor %xmm1, %xmm7
   402: 660f6fd8                 movdqa %xmm0, %xmm3
   406: 66410fdbd8               pand %xmm8, %xmm3
   411: 660f6fe0                 movdqa %xmm0, %xmm4
   415: 660f71d404               psrlw $4, %xmm4
   420: 66410fdbe0               pand %xmm8, %xmm4
   425: 66410f6faf40020000       movdqa 0x240(%r15), %xmm5
   434: 660f3800eb               pshufb %xmm3, %xmm5
   439: 66410f6fb780020000       movdqa 0x280(%r15), %xmm6
   448: 660f3800f4               pshufb %xmm4, %xmm6
   453: 660f6fd5                 movdqa %xmm5, %xmm2
   457: 660fefd6                 pxor %xmm6, %xmm2
   461: 660fefd1                 pxor %xmm1, %xmm2
   465: 660f6fdf                 movdqa %xmm7, %xmm3
   469: 66410fdbd8               pand %xmm8, %xmm3
   474: 660f6fe7                 movdqa %xmm7, %xmm4
   478: 660f71d404               psrlw $4, %xmm4
   483: 66410fdbe0               pand %xmm8, %xmm4
   488: 66410f6fafc0020000       movdqa 0x2c0(%r15), %xmm5
   497: 660f3800eb               pshufb %xmm3, %xmm5
   502: 66410f6fb700030000       movdqa 0x300(%r15), %xmm6
   511: 660f3800f4               pshufb %xmm4, %xmm6
   516: 660f6fcd                 movdqa %xmm5, %xmm1
   520: 660fefce                 pxor %xmm6, %xmm1
   524: 660f6fda                 movdqa %xmm2, %xmm3
   528: 66410fdbd8               pand %xmm8, %xmm3
   533: 660f6fe2                 movdqa %xmm2, %xmm4
   537: 660f71d404               psrlw $4, %xmm4
   542: 66410fdbe0               pand %xmm8, %xmm4
   547: 66410f6fafc0020000       movdqa 0x2c0(%r15), %xmm5
   556: 660f3800eb               pshufb %xmm3, %xmm5
   561: 66410f6fb700030000       movdqa 0x300(%r15), %xmm6
   570: 660f3800f4               pshufb %xmm4, %xmm6
   575: 660f6fc5                 movdqa %xmm5, %xmm0
   579: 660fefc6                 pxor %xmm6, %xmm0
   583: 4d8b11                   movq (%r9), %r10
   586: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   592: 4d8b5108                 movq 8(%r9), %r10
   596: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
   602: 4983c310                 addq $0x10, %r11
   606: 4c3bd8                   cmpq %rax, %r11
   609: 0f86f9fdffff             jbe 0x60
   615: 4c3bd9                   cmpq %rcx, %r11
   618: 0f83a1020000             jae 0x511
   624: 4153                     pushq %r11
   626: 498dbf40030000           leaq 0x340(%r15), %rdi
   633: 498db740040000           leaq 0x440(%r15), %rsi
   640: 4d8b11                   movq (%r9), %r10
   643: 4c8b1c24                 movq (%rsp), %r11
   647: 33db                     xorl %ebx, %ebx
   649: 498b00                   movq (%r8), %rax
   652: 420fb61418               movzbl (%rax, %r11), %edx
   657: 4885d2                   testq %rdx, %rdx
   660: 0f8422000000             je 0x2bc
   666: 0fb60417                 movzbl (%rdi, %rdx), %eax
   670: 81c0c1000000             addl $0xc1, %eax
   676: 81f8ff000000             cmpl $0xff, %eax
   682: 0f8206000000             jb 0x2b6
   688: 81c001ffffff             addl $0xffffff01, %eax
   694: 0fb60406                 movzbl (%rsi, %rax), %eax
   698: 32d8                     xorb %al, %bl
   700: 498b4008                 movq 8(%r8), %rax
   704: 420fb61418               movzbl (%rax, %r11), %edx
   709: 4885d2                   testq %rdx, %rdx
   712: 0f841f000000             je 0x2ed
   718: 0fb60417                 movzbl (%rdi, %rdx), %eax
   722: 83c003                   addl $3, %eax
   725: 81f8ff000000             cmpl $0xff, %eax
   731: 0f8206000000             jb 0x2e7
   737: 81c001ffffff             addl $0xffffff01, %eax
   743: 0fb60406                 movzbl (%rsi, %rax), %eax
   747: 32d8                     xorb %al, %bl
   749: 498b4010                 movq 0x10(%r8), %rax
   753: 420fb61418               movzbl (%rax, %r11), %edx
   758: 4885d2                   testq %rdx, %rdx
   761: 0f8422000000             je 0x321
   767: 0fb60417                 movzbl (%rdi, %rdx), %eax
   771: 81c0c6000000             addl $0xc6, %eax
   777: 81f8ff000000             cmpl $0xff, %eax
   783: 0f8206000000             jb 0x31b
   789: 81c001ffffff             addl $0xffffff01, %eax
   795: 0fb60406                 movzbl (%rsi, %rax), %eax
   799: 32d8                     xorb %al, %bl
   801: 498b4018                 movq 0x18(%r8), %rax
   805: 420fb61418               movzbl (%rax, %r11), %edx
   810: 4885d2                   testq %rdx, %rdx
   813: 0f841f000000             je 0x352
   819: 0fb60417                 movzbl (%rdi, %rdx), %eax
   823: 83c001                   addl $1, %eax
   826: 81f8ff000000             cmpl $0xff, %eax
   832: 0f8206000000             jb 0x34c
   838: 81c001ffffff             addl $0xffffff01, %eax
   844: 0fb60406                 movzbl (%rsi, %rax), %eax
   848: 32d8                     xorb %al, %bl
   850: 498b4020                 movq 0x20(%r8), %rax
   854: 420fb61418               movzbl (%rax, %r11), %edx
   859: 4885d2                   testq %rdx, %rdx
   862: 0f8422000000             je 0x386
   868: 0fb60417                 movzbl (%rdi, %rdx), %eax
   872: 81c0fe000000             addl $0xfe, %eax
   878: 81f8ff000000             cmpl $0xff, %eax
   884: 0f8206000000             jb 0x380
   890: 81c001ffffff             addl $0xffffff01, %eax
   896: 0fb60406                 movzbl (%rsi, %rax), %eax
   900: 32d8                     xorb %al, %bl
   902: 498b4028                 movq 0x28(%r8), %rax
   906: 420fb61418               movzbl (%rax, %r11), %edx
   911: 4885d2                   testq %rdx, %rdx
   914: 0f841f000000             je 0x3b7
   920: 0fb60417                 movzbl (%rdi, %rdx), %eax
   924: 83c031                   addl $0x31, %eax
   927: 81f8ff000000             cmpl $0xff, %eax
   933: 0f8206000000             jb 0x3b1
   939: 81c001ffffff             addl $0xffffff01, %eax
   945: 0fb60406                 movzbl (%rsi, %rax), %eax
   949: 32d8                     xorb %al, %bl
   951: 43881c1a                 movb %bl, (%r10, %r11)
   955: 33db                     xorl %ebx, %ebx
   957: 49ffc3                   incq %r11
   960: 4c3bd9                   cmpq %rcx, %r11
   963: 0f82c0feffff             jb 0x289
   969: 4d8b5108                 movq 8(%r9), %r10
   973: 4c8b1c24                 movq (%rsp), %r11
   977: 33db                     xorl %ebx, %ebx
   979: 498b00                   movq (%r8), %rax
   982: 420fb61418               movzbl (%rax, %r11), %edx
   987: 4885d2                   testq %rdx, %rdx
   990: 0f841f000000             je 0x403
   996: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1000: 83c01c                   addl $0x1c, %eax
  1003: 81f8ff000000             cmpl $0xff, %eax
  1009: 0f8206000000             jb 0x3fd
  1015: 81c001ffffff             addl $0xffffff01, %eax
  1021: 0fb60406                 movzbl (%rsi, %rax), %eax
  1025: 32d8                     xorb %al, %bl
  1027: 498b4008                 movq 8(%r8), %rax
  1031: 420fb61418               movzbl (%rax, %r11), %edx
  1036: 4885d2                   testq %rdx, %rdx
  1039: 0f8422000000             je 0x437
  1045: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1049: 81c0df000000             addl $0xdf, %eax
  1055: 81f8ff000000             cmpl $0xff, %eax
  1061: 0f8206000000             jb 0x431
  1067: 81c001ffffff             addl $0xffffff01, %eax
  1073: 0fb60406                 movzbl (%rsi, %rax), %eax
  1077: 32d8                     xorb %al, %bl
  1079: 498b4010                 movq 0x10(%r8), %rax
  1083: 420fb61418               movzbl (%rax, %r11), %edx
  1088: 4885d2                   testq %rdx, %rdx
  1091: 0f841f000000             je 0x468
  1097: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1101: 83c01a                   addl $0x1a, %eax
  1104: 81f8ff000000             cmpl $0xff, %eax
  1110: 0f8206000000             jb 0x462
  1116: 81c001ffffff             addl $0xffffff01, %eax
  1122: 0fb60406                 movzbl (%rsi, %rax), %eax
  1126: 32d8                     xorb %al, %bl
  1128: 498b4018                 movq 0x18(%r8), %rax
  1132: 420fb61418               movzbl (%rax, %r11), %edx
  1137: 4885d2                   testq %rdx, %rdx
  1140: 0f841f000000             je 0x499
  1146: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1150: 83c019                   addl $0x19, %eax
  1153: 81f8ff000000             cmpl $0xff, %eax
  1159: 0f8206000000             jb 0x493
  1165: 81c001ffffff             addl $0xffffff01, %eax
  1171: 0fb60406                 movzbl (%rsi, %rax), %eax
  1175: 32d8                     xorb %al, %bl
  1177: 498b4020                 movq 0x20(%r8), %rax
  1181: 420fb61418               movzbl (%rax, %r11), %edx
  1186: 4885d2                   testq %rdx, %rdx
  1189: 0f841f000000             je 0x4ca
  1195: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1199: 83c018                   addl $0x18, %eax
  1202: 81f8ff000000             cmpl $0xff, %eax
  1208: 0f8206000000             jb 0x4c4
  1214: 81c001ffffff             addl $0xffffff01, %eax
  1220: 0fb60406                 movzbl (%rsi, %rax), %eax
  1224: 32d8                     xorb %al, %bl
  1226: 498b4028                 movq 0x28(%r8), %rax
  1230: 420fb61418               movzbl (%rax, %r11), %edx
  1235: 4885d2                   testq %rdx, %rdx
  1238: 0f841f000000             je 0x4fb
  1244: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1248: 83c031                   addl $0x31, %eax
  1251: 81f8ff000000             cmpl $0xff, %eax
  1257: 0f8206000000             jb 0x4f5
  1263: 81c001ffffff             addl $0xffffff01, %eax
  1269: 0fb60406                 movzbl (%rsi, %rax), %eax
  1273: 32d8                     xorb %al, %bl
  1275: 43881c1a                 movb %bl, (%r10, %r11)
  1279: 33db                     xorl %ebx, %ebx
  1281: 49ffc3                   incq %r11
  1284: 4c3bd9                   cmpq %rcx, %r11
  1287: 0f82c6feffff             jb 0x3d3
  1293: 4883c408                 addq $8, %rsp
  1297: 415e                     popq %r14
  1299: 415d                     popq %r13
  1301: 415c                     popq %r12
  1303: 415f                     popq %r15
  1305: 5b                       popq %rbx
  1306: c3                       retq 
  1307: 0f0b                     ud2 
  1309: 0f0b                     ud2 
  1311: 0f0b                     ud2 
  1313: 0f0b                     ud2 
  1315: 0f0b                     ud2 
  1317: 0f0b                     ud2 
  1319: 0f0b                     ud2 
  1321: 0f0b                     ud2 
  1323: 0f0b                     ud2 
  1325: 0f0b                     ud2 
  1327: 0f0b                     ud2 
  1329: 0f0b                     ud2 
  1331: 0f0b                     ud2 
  1333: 0f0b                     ud2 
  1335: 0f0b                     ud2 
  1337: 0f0b                     ud2 
  1339: 0f0b                     ud2 
  1341: 0f0b                     ud2 
  1343: cc                       int3 
