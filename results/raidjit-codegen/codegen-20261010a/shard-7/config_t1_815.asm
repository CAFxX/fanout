; config_t1_815 target=sse twin=xorbit topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=1216 twins=["xorbit"] uops_per_byte=10.25
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3dad040000           leaq 0x4ad(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: 66410f6f37               movdqa (%r15), %xmm6
    27: 488bc1                   movq %rcx, %rax
    30: 4883e810                 subq $0x10, %rax
    34: 0f825a030000             jb 0x382
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5010                 movq 0x10(%r8), %r10
    68: 66430f6f041a             movdqa (%r10, %r11), %xmm0
    74: 4d8b5018                 movq 0x18(%r8), %r10
    78: 66430f6f0c1a             movdqa (%r10, %r11), %xmm1
    84: 4d8b5020                 movq 0x20(%r8), %r10
    88: 66430f6f141a             movdqa (%r10, %r11), %xmm2
    94: 4d8b10                   movq (%r8), %r10
    97: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   103: 660f6fe3                 movdqa %xmm3, %xmm4
   107: 660f71d400               psrlw $0, %xmm4
   112: 660fdbe6                 pand %xmm6, %xmm4
   116: 4d8b5008                 movq 8(%r8), %r10
   120: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   126: 660f6feb                 movdqa %xmm3, %xmm5
   130: 660f71d500               psrlw $0, %xmm5
   135: 660fdbee                 pand %xmm6, %xmm5
   139: 660fefe5                 pxor %xmm5, %xmm4
   143: 660fefc4                 pxor %xmm4, %xmm0
   147: 660fefcc                 pxor %xmm4, %xmm1
   151: 660fefd4                 pxor %xmm4, %xmm2
   155: 660f6fec                 movdqa %xmm4, %xmm5
   159: 660f71f501               psllw $1, %xmm5
   164: 660fefd5                 pxor %xmm5, %xmm2
   168: 660f6fec                 movdqa %xmm4, %xmm5
   172: 660f71f502               psllw $2, %xmm5
   177: 660fefcd                 pxor %xmm5, %xmm1
   181: 660f6fec                 movdqa %xmm4, %xmm5
   185: 660f71f503               psllw $3, %xmm5
   190: 660fefcd                 pxor %xmm5, %xmm1
   194: 660fefd5                 pxor %xmm5, %xmm2
   198: 660f6fec                 movdqa %xmm4, %xmm5
   202: 660f71f506               psllw $6, %xmm5
   207: 660fefd5                 pxor %xmm5, %xmm2
   211: 660f6fec                 movdqa %xmm4, %xmm5
   215: 660f71f507               psllw $7, %xmm5
   220: 660fefcd                 pxor %xmm5, %xmm1
   224: 660fefd5                 pxor %xmm5, %xmm2
   228: 4d8b10                   movq (%r8), %r10
   231: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   237: 660f6fe3                 movdqa %xmm3, %xmm4
   241: 660f71d401               psrlw $1, %xmm4
   246: 660fdbe6                 pand %xmm6, %xmm4
   250: 4d8b5008                 movq 8(%r8), %r10
   254: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   260: 660f6feb                 movdqa %xmm3, %xmm5
   264: 660f71d501               psrlw $1, %xmm5
   269: 660fdbee                 pand %xmm6, %xmm5
   273: 660fefe5                 pxor %xmm5, %xmm4
   277: 660fefcc                 pxor %xmm4, %xmm1
   281: 660fefd4                 pxor %xmm4, %xmm2
   285: 660f6fec                 movdqa %xmm4, %xmm5
   289: 660f71f501               psllw $1, %xmm5
   294: 660fefc5                 pxor %xmm5, %xmm0
   298: 660f6fec                 movdqa %xmm4, %xmm5
   302: 660f71f502               psllw $2, %xmm5
   307: 660fefd5                 pxor %xmm5, %xmm2
   311: 660f6fec                 movdqa %xmm4, %xmm5
   315: 660f71f503               psllw $3, %xmm5
   320: 660fefd5                 pxor %xmm5, %xmm2
   324: 660f6fec                 movdqa %xmm4, %xmm5
   328: 660f71f507               psllw $7, %xmm5
   333: 660fefd5                 pxor %xmm5, %xmm2
   337: 4d8b10                   movq (%r8), %r10
   340: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   346: 660f6fe3                 movdqa %xmm3, %xmm4
   350: 660f71d402               psrlw $2, %xmm4
   355: 660fdbe6                 pand %xmm6, %xmm4
   359: 4d8b5008                 movq 8(%r8), %r10
   363: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   369: 660f6feb                 movdqa %xmm3, %xmm5
   373: 660f71d502               psrlw $2, %xmm5
   378: 660fdbee                 pand %xmm6, %xmm5
   382: 660fefe5                 pxor %xmm5, %xmm4
   386: 660fefd4                 pxor %xmm4, %xmm2
   390: 660f6fec                 movdqa %xmm4, %xmm5
   394: 660f71f501               psllw $1, %xmm5
   399: 660fefcd                 pxor %xmm5, %xmm1
   403: 660f6fec                 movdqa %xmm4, %xmm5
   407: 660f71f502               psllw $2, %xmm5
   412: 660fefc5                 pxor %xmm5, %xmm0
   416: 4d8b10                   movq (%r8), %r10
   419: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   425: 660f6fe3                 movdqa %xmm3, %xmm4
   429: 660f71d403               psrlw $3, %xmm4
   434: 660fdbe6                 pand %xmm6, %xmm4
   438: 4d8b5008                 movq 8(%r8), %r10
   442: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   448: 660f6feb                 movdqa %xmm3, %xmm5
   452: 660f71d503               psrlw $3, %xmm5
   457: 660fdbee                 pand %xmm6, %xmm5
   461: 660fefe5                 pxor %xmm5, %xmm4
   465: 660f6fec                 movdqa %xmm4, %xmm5
   469: 660f71f501               psllw $1, %xmm5
   474: 660fefd5                 pxor %xmm5, %xmm2
   478: 660f6fec                 movdqa %xmm4, %xmm5
   482: 660f71f502               psllw $2, %xmm5
   487: 660fefcd                 pxor %xmm5, %xmm1
   491: 660f6fec                 movdqa %xmm4, %xmm5
   495: 660f71f503               psllw $3, %xmm5
   500: 660fefc5                 pxor %xmm5, %xmm0
   504: 4d8b10                   movq (%r8), %r10
   507: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   513: 660f6fe3                 movdqa %xmm3, %xmm4
   517: 660f71d404               psrlw $4, %xmm4
   522: 660fdbe6                 pand %xmm6, %xmm4
   526: 4d8b5008                 movq 8(%r8), %r10
   530: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   536: 660f6feb                 movdqa %xmm3, %xmm5
   540: 660f71d504               psrlw $4, %xmm5
   545: 660fdbee                 pand %xmm6, %xmm5
   549: 660fefe5                 pxor %xmm5, %xmm4
   553: 660f6fec                 movdqa %xmm4, %xmm5
   557: 660f71f502               psllw $2, %xmm5
   562: 660fefd5                 pxor %xmm5, %xmm2
   566: 660f6fec                 movdqa %xmm4, %xmm5
   570: 660f71f503               psllw $3, %xmm5
   575: 660fefcd                 pxor %xmm5, %xmm1
   579: 660f6fec                 movdqa %xmm4, %xmm5
   583: 660f71f504               psllw $4, %xmm5
   588: 660fefc5                 pxor %xmm5, %xmm0
   592: 4d8b10                   movq (%r8), %r10
   595: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   601: 660f6fe3                 movdqa %xmm3, %xmm4
   605: 660f71d405               psrlw $5, %xmm4
   610: 660fdbe6                 pand %xmm6, %xmm4
   614: 4d8b5008                 movq 8(%r8), %r10
   618: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   624: 660f6feb                 movdqa %xmm3, %xmm5
   628: 660f71d505               psrlw $5, %xmm5
   633: 660fdbee                 pand %xmm6, %xmm5
   637: 660fefe5                 pxor %xmm5, %xmm4
   641: 660f6fec                 movdqa %xmm4, %xmm5
   645: 660f71f503               psllw $3, %xmm5
   650: 660fefd5                 pxor %xmm5, %xmm2
   654: 660f6fec                 movdqa %xmm4, %xmm5
   658: 660f71f504               psllw $4, %xmm5
   663: 660fefcd                 pxor %xmm5, %xmm1
   667: 660f6fec                 movdqa %xmm4, %xmm5
   671: 660f71f505               psllw $5, %xmm5
   676: 660fefc5                 pxor %xmm5, %xmm0
   680: 4d8b10                   movq (%r8), %r10
   683: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   689: 660f6fe3                 movdqa %xmm3, %xmm4
   693: 660f71d406               psrlw $6, %xmm4
   698: 660fdbe6                 pand %xmm6, %xmm4
   702: 4d8b5008                 movq 8(%r8), %r10
   706: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   712: 660f6feb                 movdqa %xmm3, %xmm5
   716: 660f71d506               psrlw $6, %xmm5
   721: 660fdbee                 pand %xmm6, %xmm5
   725: 660fefe5                 pxor %xmm5, %xmm4
   729: 660f6fec                 movdqa %xmm4, %xmm5
   733: 660f71f504               psllw $4, %xmm5
   738: 660fefd5                 pxor %xmm5, %xmm2
   742: 660f6fec                 movdqa %xmm4, %xmm5
   746: 660f71f505               psllw $5, %xmm5
   751: 660fefcd                 pxor %xmm5, %xmm1
   755: 660f6fec                 movdqa %xmm4, %xmm5
   759: 660f71f506               psllw $6, %xmm5
   764: 660fefc5                 pxor %xmm5, %xmm0
   768: 4d8b10                   movq (%r8), %r10
   771: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   777: 660f6fe3                 movdqa %xmm3, %xmm4
   781: 660f71d407               psrlw $7, %xmm4
   786: 660fdbe6                 pand %xmm6, %xmm4
   790: 4d8b5008                 movq 8(%r8), %r10
   794: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   800: 660f6feb                 movdqa %xmm3, %xmm5
   804: 660f71d507               psrlw $7, %xmm5
   809: 660fdbee                 pand %xmm6, %xmm5
   813: 660fefe5                 pxor %xmm5, %xmm4
   817: 660f6fec                 movdqa %xmm4, %xmm5
   821: 660f71f505               psllw $5, %xmm5
   826: 660fefd5                 pxor %xmm5, %xmm2
   830: 660f6fec                 movdqa %xmm4, %xmm5
   834: 660f71f506               psllw $6, %xmm5
   839: 660fefcd                 pxor %xmm5, %xmm1
   843: 660f6fec                 movdqa %xmm4, %xmm5
   847: 660f71f507               psllw $7, %xmm5
   852: 660fefc5                 pxor %xmm5, %xmm0
   856: 4d8b11                   movq (%r9), %r10
   859: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   865: 4d8b5108                 movq 8(%r9), %r10
   869: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
   875: 4d8b5110                 movq 0x10(%r9), %r10
   879: 66430f7f141a             movdqa %xmm2, (%r10, %r11)
   885: 4983c310                 addq $0x10, %r11
   889: 4c3bd8                   cmpq %rax, %r11
   892: 0f86befcffff             jbe 0x40
   898: 4c3bd9                   cmpq %rcx, %r11
   901: 0f8304010000             jae 0x48f
   907: 4153                     pushq %r11
   909: 498d7f40                 leaq 0x40(%r15), %rdi
   913: 498db740010000           leaq 0x140(%r15), %rsi
   920: 4d8b11                   movq (%r9), %r10
   923: 4c8b1c24                 movq (%rsp), %r11
   927: 498b4010                 movq 0x10(%r8), %rax
   931: 420fb61c18               movzbl (%rax, %r11), %ebx
   936: 498b00                   movq (%r8), %rax
   939: 420fb61418               movzbl (%rax, %r11), %edx
   944: 498b4008                 movq 8(%r8), %rax
   948: 420fb60418               movzbl (%rax, %r11), %eax
   953: 32d0                     xorb %al, %dl
   955: 32da                     xorb %dl, %bl
   957: 43881c1a                 movb %bl, (%r10, %r11)
   961: 49ffc3                   incq %r11
   964: 4c3bd9                   cmpq %rcx, %r11
   967: 0f82d2ffffff             jb 0x39f
   973: 4d8b5108                 movq 8(%r9), %r10
   977: 4c8b1c24                 movq (%rsp), %r11
   981: 498b4018                 movq 0x18(%r8), %rax
   985: 420fb61c18               movzbl (%rax, %r11), %ebx
   990: 498b00                   movq (%r8), %rax
   993: 420fb61418               movzbl (%rax, %r11), %edx
   998: 498b4008                 movq 8(%r8), %rax
  1002: 420fb60418               movzbl (%rax, %r11), %eax
  1007: 32d0                     xorb %al, %dl
  1009: 4885d2                   testq %rdx, %rdx
  1012: 0f8422000000             je 0x41c
  1018: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1022: 81c0e6000000             addl $0xe6, %eax
  1028: 81f8ff000000             cmpl $0xff, %eax
  1034: 0f8206000000             jb 0x416
  1040: 81c001ffffff             addl $0xffffff01, %eax
  1046: 0fb60406                 movzbl (%rsi, %rax), %eax
  1050: 32d8                     xorb %al, %bl
  1052: 43881c1a                 movb %bl, (%r10, %r11)
  1056: 49ffc3                   incq %r11
  1059: 4c3bd9                   cmpq %rcx, %r11
  1062: 0f82a9ffffff             jb 0x3d5
  1068: 4d8b5110                 movq 0x10(%r9), %r10
  1072: 4c8b1c24                 movq (%rsp), %r11
  1076: 498b4020                 movq 0x20(%r8), %rax
  1080: 420fb61c18               movzbl (%rax, %r11), %ebx
  1085: 498b00                   movq (%r8), %rax
  1088: 420fb61418               movzbl (%rax, %r11), %edx
  1093: 498b4008                 movq 8(%r8), %rax
  1097: 420fb60418               movzbl (%rax, %r11), %eax
  1102: 32d0                     xorb %al, %dl
  1104: 4885d2                   testq %rdx, %rdx
  1107: 0f8422000000             je 0x47b
  1113: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1117: 81c0cd000000             addl $0xcd, %eax
  1123: 81f8ff000000             cmpl $0xff, %eax
  1129: 0f8206000000             jb 0x475
  1135: 81c001ffffff             addl $0xffffff01, %eax
  1141: 0fb60406                 movzbl (%rsi, %rax), %eax
  1145: 32d8                     xorb %al, %bl
  1147: 43881c1a                 movb %bl, (%r10, %r11)
  1151: 49ffc3                   incq %r11
  1154: 4c3bd9                   cmpq %rcx, %r11
  1157: 0f82a9ffffff             jb 0x434
  1163: 4883c408                 addq $8, %rsp
  1167: 415f                     popq %r15
  1169: 5b                       popq %rbx
  1170: c3                       retq 
  1171: 0f0b                     ud2 
  1173: 0f0b                     ud2 
  1175: 0f0b                     ud2 
  1177: 0f0b                     ud2 
  1179: 0f0b                     ud2 
  1181: 0f0b                     ud2 
  1183: 0f0b                     ud2 
  1185: 0f0b                     ud2 
  1187: 0f0b                     ud2 
  1189: 0f0b                     ud2 
  1191: 0f0b                     ud2 
  1193: 0f0b                     ud2 
  1195: 0f0b                     ud2 
  1197: 0f0b                     ud2 
  1199: 0f0b                     ud2 
  1201: 0f0b                     ud2 
  1203: 0f0b                     ud2 
  1205: 0f0b                     ud2 
  1207: 0f0b                     ud2 
  1209: 0f0b                     ud2 
  1211: 0f0b                     ud2 
  1213: 0f0b                     ud2 
  1215: cc                       int3 
