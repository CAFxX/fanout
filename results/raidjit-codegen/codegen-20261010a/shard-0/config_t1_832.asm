; config_t1_832 target=sse twin=baseline topo=raid6(8)/upd{0,1} op=parity_update U=1
; code_size=1280 twins=[] uops_per_byte=2.375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3ded040000           leaq 0x4ed(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: 66450f6f37               movdqa (%r15), %xmm14
    27: 488bc1                   movq %rcx, %rax
    30: 4883e840                 subq $0x40, %rax
    34: 0f82bf030000             jb 0x3e7
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5020                 movq 0x20(%r8), %r10
    68: 66430f6f041a             movdqa (%r10, %r11), %xmm0
    74: 66430f6f4c1a10           movdqa 0x10(%r10, %r11), %xmm1
    81: 66430f6f541a20           movdqa 0x20(%r10, %r11), %xmm2
    88: 66430f6f5c1a30           movdqa 0x30(%r10, %r11), %xmm3
    95: 4d8b5028                 movq 0x28(%r8), %r10
    99: 66430f6f241a             movdqa (%r10, %r11), %xmm4
   105: 66430f6f6c1a10           movdqa 0x10(%r10, %r11), %xmm5
   112: 66430f6f741a20           movdqa 0x20(%r10, %r11), %xmm6
   119: 66430f6f7c1a30           movdqa 0x30(%r10, %r11), %xmm7
   126: 4d8b10                   movq (%r8), %r10
   129: 498b5010                 movq 0x10(%r8), %rdx
   133: 66470f6f041a             movdqa (%r10, %r11), %xmm8
   139: 66460f6f1c1a             movdqa (%rdx, %r11), %xmm11
   145: 66450fefc3               pxor %xmm11, %xmm8
   150: 66450f6fc8               movdqa %xmm8, %xmm9
   155: 66450fdbce               pand %xmm14, %xmm9
   160: 66450f6fd0               movdqa %xmm8, %xmm10
   165: 66410f71d204             psrlw $4, %xmm10
   171: 66450fdbd6               pand %xmm14, %xmm10
   176: 66410fefc0               pxor %xmm8, %xmm0
   181: 66450f6f6f40             movdqa 0x40(%r15), %xmm13
   187: 66450f6fdd               movdqa %xmm13, %xmm11
   192: 66450f3800d9             pshufb %xmm9, %xmm11
   198: 66450f6faf80000000       movdqa 0x80(%r15), %xmm13
   207: 66450f6fe5               movdqa %xmm13, %xmm12
   212: 66450f3800e2             pshufb %xmm10, %xmm12
   218: 66410fefe3               pxor %xmm11, %xmm4
   223: 66410fefe4               pxor %xmm12, %xmm4
   228: 66470f6f441a10           movdqa 0x10(%r10, %r11), %xmm8
   235: 66460f6f5c1a10           movdqa 0x10(%rdx, %r11), %xmm11
   242: 66450fefc3               pxor %xmm11, %xmm8
   247: 66450f6fc8               movdqa %xmm8, %xmm9
   252: 66450fdbce               pand %xmm14, %xmm9
   257: 66450f6fd0               movdqa %xmm8, %xmm10
   262: 66410f71d204             psrlw $4, %xmm10
   268: 66450fdbd6               pand %xmm14, %xmm10
   273: 66410fefc8               pxor %xmm8, %xmm1
   278: 66450f6f6f40             movdqa 0x40(%r15), %xmm13
   284: 66450f6fdd               movdqa %xmm13, %xmm11
   289: 66450f3800d9             pshufb %xmm9, %xmm11
   295: 66450f6faf80000000       movdqa 0x80(%r15), %xmm13
   304: 66450f6fe5               movdqa %xmm13, %xmm12
   309: 66450f3800e2             pshufb %xmm10, %xmm12
   315: 66410fefeb               pxor %xmm11, %xmm5
   320: 66410fefec               pxor %xmm12, %xmm5
   325: 66470f6f441a20           movdqa 0x20(%r10, %r11), %xmm8
   332: 66460f6f5c1a20           movdqa 0x20(%rdx, %r11), %xmm11
   339: 66450fefc3               pxor %xmm11, %xmm8
   344: 66450f6fc8               movdqa %xmm8, %xmm9
   349: 66450fdbce               pand %xmm14, %xmm9
   354: 66450f6fd0               movdqa %xmm8, %xmm10
   359: 66410f71d204             psrlw $4, %xmm10
   365: 66450fdbd6               pand %xmm14, %xmm10
   370: 66410fefd0               pxor %xmm8, %xmm2
   375: 66450f6f6f40             movdqa 0x40(%r15), %xmm13
   381: 66450f6fdd               movdqa %xmm13, %xmm11
   386: 66450f3800d9             pshufb %xmm9, %xmm11
   392: 66450f6faf80000000       movdqa 0x80(%r15), %xmm13
   401: 66450f6fe5               movdqa %xmm13, %xmm12
   406: 66450f3800e2             pshufb %xmm10, %xmm12
   412: 66410feff3               pxor %xmm11, %xmm6
   417: 66410feff4               pxor %xmm12, %xmm6
   422: 66470f6f441a30           movdqa 0x30(%r10, %r11), %xmm8
   429: 66460f6f5c1a30           movdqa 0x30(%rdx, %r11), %xmm11
   436: 66450fefc3               pxor %xmm11, %xmm8
   441: 66450f6fc8               movdqa %xmm8, %xmm9
   446: 66450fdbce               pand %xmm14, %xmm9
   451: 66450f6fd0               movdqa %xmm8, %xmm10
   456: 66410f71d204             psrlw $4, %xmm10
   462: 66450fdbd6               pand %xmm14, %xmm10
   467: 66410fefd8               pxor %xmm8, %xmm3
   472: 66450f6f6f40             movdqa 0x40(%r15), %xmm13
   478: 66450f6fdd               movdqa %xmm13, %xmm11
   483: 66450f3800d9             pshufb %xmm9, %xmm11
   489: 66450f6faf80000000       movdqa 0x80(%r15), %xmm13
   498: 66450f6fe5               movdqa %xmm13, %xmm12
   503: 66450f3800e2             pshufb %xmm10, %xmm12
   509: 66410feffb               pxor %xmm11, %xmm7
   514: 66410feffc               pxor %xmm12, %xmm7
   519: 4d8b5008                 movq 8(%r8), %r10
   523: 498b5018                 movq 0x18(%r8), %rdx
   527: 66470f6f041a             movdqa (%r10, %r11), %xmm8
   533: 66460f6f1c1a             movdqa (%rdx, %r11), %xmm11
   539: 66450fefc3               pxor %xmm11, %xmm8
   544: 66450f6fc8               movdqa %xmm8, %xmm9
   549: 66450fdbce               pand %xmm14, %xmm9
   554: 66450f6fd0               movdqa %xmm8, %xmm10
   559: 66410f71d204             psrlw $4, %xmm10
   565: 66450fdbd6               pand %xmm14, %xmm10
   570: 66410fefc0               pxor %xmm8, %xmm0
   575: 66450f6fafc0000000       movdqa 0xc0(%r15), %xmm13
   584: 66450f6fdd               movdqa %xmm13, %xmm11
   589: 66450f3800d9             pshufb %xmm9, %xmm11
   595: 66450f6faf00010000       movdqa 0x100(%r15), %xmm13
   604: 66450f6fe5               movdqa %xmm13, %xmm12
   609: 66450f3800e2             pshufb %xmm10, %xmm12
   615: 66410fefe3               pxor %xmm11, %xmm4
   620: 66410fefe4               pxor %xmm12, %xmm4
   625: 66470f6f441a10           movdqa 0x10(%r10, %r11), %xmm8
   632: 66460f6f5c1a10           movdqa 0x10(%rdx, %r11), %xmm11
   639: 66450fefc3               pxor %xmm11, %xmm8
   644: 66450f6fc8               movdqa %xmm8, %xmm9
   649: 66450fdbce               pand %xmm14, %xmm9
   654: 66450f6fd0               movdqa %xmm8, %xmm10
   659: 66410f71d204             psrlw $4, %xmm10
   665: 66450fdbd6               pand %xmm14, %xmm10
   670: 66410fefc8               pxor %xmm8, %xmm1
   675: 66450f6fafc0000000       movdqa 0xc0(%r15), %xmm13
   684: 66450f6fdd               movdqa %xmm13, %xmm11
   689: 66450f3800d9             pshufb %xmm9, %xmm11
   695: 66450f6faf00010000       movdqa 0x100(%r15), %xmm13
   704: 66450f6fe5               movdqa %xmm13, %xmm12
   709: 66450f3800e2             pshufb %xmm10, %xmm12
   715: 66410fefeb               pxor %xmm11, %xmm5
   720: 66410fefec               pxor %xmm12, %xmm5
   725: 66470f6f441a20           movdqa 0x20(%r10, %r11), %xmm8
   732: 66460f6f5c1a20           movdqa 0x20(%rdx, %r11), %xmm11
   739: 66450fefc3               pxor %xmm11, %xmm8
   744: 66450f6fc8               movdqa %xmm8, %xmm9
   749: 66450fdbce               pand %xmm14, %xmm9
   754: 66450f6fd0               movdqa %xmm8, %xmm10
   759: 66410f71d204             psrlw $4, %xmm10
   765: 66450fdbd6               pand %xmm14, %xmm10
   770: 66410fefd0               pxor %xmm8, %xmm2
   775: 66450f6fafc0000000       movdqa 0xc0(%r15), %xmm13
   784: 66450f6fdd               movdqa %xmm13, %xmm11
   789: 66450f3800d9             pshufb %xmm9, %xmm11
   795: 66450f6faf00010000       movdqa 0x100(%r15), %xmm13
   804: 66450f6fe5               movdqa %xmm13, %xmm12
   809: 66450f3800e2             pshufb %xmm10, %xmm12
   815: 66410feff3               pxor %xmm11, %xmm6
   820: 66410feff4               pxor %xmm12, %xmm6
   825: 66470f6f441a30           movdqa 0x30(%r10, %r11), %xmm8
   832: 66460f6f5c1a30           movdqa 0x30(%rdx, %r11), %xmm11
   839: 66450fefc3               pxor %xmm11, %xmm8
   844: 66450f6fc8               movdqa %xmm8, %xmm9
   849: 66450fdbce               pand %xmm14, %xmm9
   854: 66450f6fd0               movdqa %xmm8, %xmm10
   859: 66410f71d204             psrlw $4, %xmm10
   865: 66450fdbd6               pand %xmm14, %xmm10
   870: 66410fefd8               pxor %xmm8, %xmm3
   875: 66450f6fafc0000000       movdqa 0xc0(%r15), %xmm13
   884: 66450f6fdd               movdqa %xmm13, %xmm11
   889: 66450f3800d9             pshufb %xmm9, %xmm11
   895: 66450f6faf00010000       movdqa 0x100(%r15), %xmm13
   904: 66450f6fe5               movdqa %xmm13, %xmm12
   909: 66450f3800e2             pshufb %xmm10, %xmm12
   915: 66410feffb               pxor %xmm11, %xmm7
   920: 66410feffc               pxor %xmm12, %xmm7
   925: 4d8b11                   movq (%r9), %r10
   928: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   934: 66430f7f4c1a10           movdqa %xmm1, 0x10(%r10, %r11)
   941: 66430f7f541a20           movdqa %xmm2, 0x20(%r10, %r11)
   948: 66430f7f5c1a30           movdqa %xmm3, 0x30(%r10, %r11)
   955: 4d8b5108                 movq 8(%r9), %r10
   959: 66430f7f241a             movdqa %xmm4, (%r10, %r11)
   965: 66430f7f6c1a10           movdqa %xmm5, 0x10(%r10, %r11)
   972: 66430f7f741a20           movdqa %xmm6, 0x20(%r10, %r11)
   979: 66430f7f7c1a30           movdqa %xmm7, 0x30(%r10, %r11)
   986: 4983c340                 addq $0x40, %r11
   990: 4c3bd8                   cmpq %rax, %r11
   993: 0f8659fcffff             jbe 0x40
   999: 4c3bd9                   cmpq %rcx, %r11
  1002: 0f83fd000000             jae 0x4ed
  1008: 4153                     pushq %r11
  1010: 498dbf40010000           leaq 0x140(%r15), %rdi
  1017: 498db740020000           leaq 0x240(%r15), %rsi
  1024: 4d8b11                   movq (%r9), %r10
  1027: 4c8b1c24                 movq (%rsp), %r11
  1031: 498b4020                 movq 0x20(%r8), %rax
  1035: 420fb61c18               movzbl (%rax, %r11), %ebx
  1040: 498b00                   movq (%r8), %rax
  1043: 420fb61418               movzbl (%rax, %r11), %edx
  1048: 498b4010                 movq 0x10(%r8), %rax
  1052: 420fb60418               movzbl (%rax, %r11), %eax
  1057: 32d0                     xorb %al, %dl
  1059: 32da                     xorb %dl, %bl
  1061: 498b4008                 movq 8(%r8), %rax
  1065: 420fb61418               movzbl (%rax, %r11), %edx
  1070: 498b4018                 movq 0x18(%r8), %rax
  1074: 420fb60418               movzbl (%rax, %r11), %eax
  1079: 32d0                     xorb %al, %dl
  1081: 32da                     xorb %dl, %bl
  1083: 43881c1a                 movb %bl, (%r10, %r11)
  1087: 49ffc3                   incq %r11
  1090: 4c3bd9                   cmpq %rcx, %r11
  1093: 0f82bcffffff             jb 0x407
  1099: 4d8b5108                 movq 8(%r9), %r10
  1103: 4c8b1c24                 movq (%rsp), %r11
  1107: 498b4028                 movq 0x28(%r8), %rax
  1111: 420fb61c18               movzbl (%rax, %r11), %ebx
  1116: 498b00                   movq (%r8), %rax
  1119: 420fb61418               movzbl (%rax, %r11), %edx
  1124: 498b4010                 movq 0x10(%r8), %rax
  1128: 420fb60418               movzbl (%rax, %r11), %eax
  1133: 32d0                     xorb %al, %dl
  1135: 4885d2                   testq %rdx, %rdx
  1138: 0f8422000000             je 0x49a
  1144: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1148: 81c0e6000000             addl $0xe6, %eax
  1154: 81f8ff000000             cmpl $0xff, %eax
  1160: 0f8206000000             jb 0x494
  1166: 81c001ffffff             addl $0xffffff01, %eax
  1172: 0fb60406                 movzbl (%rsi, %rax), %eax
  1176: 32d8                     xorb %al, %bl
  1178: 498b4008                 movq 8(%r8), %rax
  1182: 420fb61418               movzbl (%rax, %r11), %edx
  1187: 498b4018                 movq 0x18(%r8), %rax
  1191: 420fb60418               movzbl (%rax, %r11), %eax
  1196: 32d0                     xorb %al, %dl
  1198: 4885d2                   testq %rdx, %rdx
  1201: 0f8422000000             je 0x4d9
  1207: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1211: 81c0cd000000             addl $0xcd, %eax
  1217: 81f8ff000000             cmpl $0xff, %eax
  1223: 0f8206000000             jb 0x4d3
  1229: 81c001ffffff             addl $0xffffff01, %eax
  1235: 0fb60406                 movzbl (%rsi, %rax), %eax
  1239: 32d8                     xorb %al, %bl
  1241: 43881c1a                 movb %bl, (%r10, %r11)
  1245: 49ffc3                   incq %r11
  1248: 4c3bd9                   cmpq %rcx, %r11
  1251: 0f826affffff             jb 0x453
  1257: 4883c408                 addq $8, %rsp
  1261: 415f                     popq %r15
  1263: 5b                       popq %rbx
  1264: c3                       retq 
  1265: 0f0b                     ud2 
  1267: 0f0b                     ud2 
  1269: 0f0b                     ud2 
  1271: 0f0b                     ud2 
  1273: 0f0b                     ud2 
  1275: 0f0b                     ud2 
  1277: 0f0b                     ud2 
  1279: cc                       int3 
