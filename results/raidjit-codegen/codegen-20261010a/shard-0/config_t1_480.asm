; config_t1_480 target=avx2 twin=rec_cf topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=1280 twins=[] uops_per_byte=2.375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de7040000           leaq 0x4e7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: c4417d6f07               vmovdqa (%r15), %ymm8
    52: 488bc1                   movq %rcx, %rax
    55: 4883e820                 subq $0x20, %rax
    59: 0f82da010000             jb 0x21b
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   102: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   107: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   112: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   117: c5fd6fc2                 vmovdqa %ymm2, %ymm0
   121: c4c17d6f6f40             vmovdqa 0x40(%r15), %ymm5
   127: c4e25500eb               vpshufb %ymm3, %ymm5, %ymm5
   132: c4c17d6fb780000000       vmovdqa 0x80(%r15), %ymm6
   141: c4e24d00f4               vpshufb %ymm4, %ymm6, %ymm6
   146: c5d5efce                 vpxor %ymm6, %ymm5, %ymm1
   150: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   156: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   161: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   166: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   171: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   175: c4c17d6fafc0000000       vmovdqa 0xc0(%r15), %ymm5
   184: c4e25500eb               vpshufb %ymm3, %ymm5, %ymm5
   189: c4c17d6fb700010000       vmovdqa 0x100(%r15), %ymm6
   198: c4e24d00f4               vpshufb %ymm4, %ymm6, %ymm6
   203: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   207: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   211: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   217: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   222: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   227: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   232: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   236: c4c17d6faf40010000       vmovdqa 0x140(%r15), %ymm5
   245: c4e25500eb               vpshufb %ymm3, %ymm5, %ymm5
   250: c4c17d6fb780010000       vmovdqa 0x180(%r15), %ymm6
   259: c4e24d00f4               vpshufb %ymm4, %ymm6, %ymm6
   264: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   268: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   272: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   279: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   283: c5f5efca                 vpxor %ymm2, %ymm1, %ymm1
   287: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   293: c5fdefc2                 vpxor %ymm2, %ymm0, %ymm0
   297: 4d8b5028                 movq 0x28(%r8), %r10
   301: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   307: c5f5efca                 vpxor %ymm2, %ymm1, %ymm1
   311: c4c17ddbd8               vpand %ymm8, %ymm0, %ymm3
   316: c5dd71d004               vpsrlw $4, %ymm0, %ymm4
   321: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   326: c4c17d6fafc0010000       vmovdqa 0x1c0(%r15), %ymm5
   335: c4e25500eb               vpshufb %ymm3, %ymm5, %ymm5
   340: c4c17d6fb700020000       vmovdqa 0x200(%r15), %ymm6
   349: c4e24d00f4               vpshufb %ymm4, %ymm6, %ymm6
   354: c5d5effe                 vpxor %ymm6, %ymm5, %ymm7
   358: c5c5eff9                 vpxor %ymm1, %ymm7, %ymm7
   362: c4c17ddbd8               vpand %ymm8, %ymm0, %ymm3
   367: c5dd71d004               vpsrlw $4, %ymm0, %ymm4
   372: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   377: c4c17d6faf40020000       vmovdqa 0x240(%r15), %ymm5
   386: c4e25500eb               vpshufb %ymm3, %ymm5, %ymm5
   391: c4c17d6fb780020000       vmovdqa 0x280(%r15), %ymm6
   400: c4e24d00f4               vpshufb %ymm4, %ymm6, %ymm6
   405: c5d5efd6                 vpxor %ymm6, %ymm5, %ymm2
   409: c5edefd1                 vpxor %ymm1, %ymm2, %ymm2
   413: c4c145dbd8               vpand %ymm8, %ymm7, %ymm3
   418: c5dd71d704               vpsrlw $4, %ymm7, %ymm4
   423: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   428: c4c17d6fafc0020000       vmovdqa 0x2c0(%r15), %ymm5
   437: c4e25500eb               vpshufb %ymm3, %ymm5, %ymm5
   442: c4c17d6fb700030000       vmovdqa 0x300(%r15), %ymm6
   451: c4e24d00f4               vpshufb %ymm4, %ymm6, %ymm6
   456: c5d5efce                 vpxor %ymm6, %ymm5, %ymm1
   460: c4c16ddbd8               vpand %ymm8, %ymm2, %ymm3
   465: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   470: c4c15ddbe0               vpand %ymm8, %ymm4, %ymm4
   475: c4c17d6fafc0020000       vmovdqa 0x2c0(%r15), %ymm5
   484: c4e25500eb               vpshufb %ymm3, %ymm5, %ymm5
   489: c4c17d6fb700030000       vmovdqa 0x300(%r15), %ymm6
   498: c4e24d00f4               vpshufb %ymm4, %ymm6, %ymm6
   503: c5d5efc6                 vpxor %ymm6, %ymm5, %ymm0
   507: 4d8b11                   movq (%r9), %r10
   510: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   516: 4d8b5108                 movq 8(%r9), %r10
   520: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
   526: 4983c320                 addq $0x20, %r11
   530: 4c3bd8                   cmpq %rax, %r11
   533: 0f8645feffff             jbe 0x60
   539: 4c3bd9                   cmpq %rcx, %r11
   542: 0f83a1020000             jae 0x4c5
   548: 4153                     pushq %r11
   550: 498dbf40030000           leaq 0x340(%r15), %rdi
   557: 498db740040000           leaq 0x440(%r15), %rsi
   564: 4d8b11                   movq (%r9), %r10
   567: 4c8b1c24                 movq (%rsp), %r11
   571: 33db                     xorl %ebx, %ebx
   573: 498b00                   movq (%r8), %rax
   576: 420fb61418               movzbl (%rax, %r11), %edx
   581: 4885d2                   testq %rdx, %rdx
   584: 0f8422000000             je 0x270
   590: 0fb60417                 movzbl (%rdi, %rdx), %eax
   594: 81c0c1000000             addl $0xc1, %eax
   600: 81f8ff000000             cmpl $0xff, %eax
   606: 0f8206000000             jb 0x26a
   612: 81c001ffffff             addl $0xffffff01, %eax
   618: 0fb60406                 movzbl (%rsi, %rax), %eax
   622: 32d8                     xorb %al, %bl
   624: 498b4008                 movq 8(%r8), %rax
   628: 420fb61418               movzbl (%rax, %r11), %edx
   633: 4885d2                   testq %rdx, %rdx
   636: 0f841f000000             je 0x2a1
   642: 0fb60417                 movzbl (%rdi, %rdx), %eax
   646: 83c003                   addl $3, %eax
   649: 81f8ff000000             cmpl $0xff, %eax
   655: 0f8206000000             jb 0x29b
   661: 81c001ffffff             addl $0xffffff01, %eax
   667: 0fb60406                 movzbl (%rsi, %rax), %eax
   671: 32d8                     xorb %al, %bl
   673: 498b4010                 movq 0x10(%r8), %rax
   677: 420fb61418               movzbl (%rax, %r11), %edx
   682: 4885d2                   testq %rdx, %rdx
   685: 0f8422000000             je 0x2d5
   691: 0fb60417                 movzbl (%rdi, %rdx), %eax
   695: 81c0c6000000             addl $0xc6, %eax
   701: 81f8ff000000             cmpl $0xff, %eax
   707: 0f8206000000             jb 0x2cf
   713: 81c001ffffff             addl $0xffffff01, %eax
   719: 0fb60406                 movzbl (%rsi, %rax), %eax
   723: 32d8                     xorb %al, %bl
   725: 498b4018                 movq 0x18(%r8), %rax
   729: 420fb61418               movzbl (%rax, %r11), %edx
   734: 4885d2                   testq %rdx, %rdx
   737: 0f841f000000             je 0x306
   743: 0fb60417                 movzbl (%rdi, %rdx), %eax
   747: 83c001                   addl $1, %eax
   750: 81f8ff000000             cmpl $0xff, %eax
   756: 0f8206000000             jb 0x300
   762: 81c001ffffff             addl $0xffffff01, %eax
   768: 0fb60406                 movzbl (%rsi, %rax), %eax
   772: 32d8                     xorb %al, %bl
   774: 498b4020                 movq 0x20(%r8), %rax
   778: 420fb61418               movzbl (%rax, %r11), %edx
   783: 4885d2                   testq %rdx, %rdx
   786: 0f8422000000             je 0x33a
   792: 0fb60417                 movzbl (%rdi, %rdx), %eax
   796: 81c0fe000000             addl $0xfe, %eax
   802: 81f8ff000000             cmpl $0xff, %eax
   808: 0f8206000000             jb 0x334
   814: 81c001ffffff             addl $0xffffff01, %eax
   820: 0fb60406                 movzbl (%rsi, %rax), %eax
   824: 32d8                     xorb %al, %bl
   826: 498b4028                 movq 0x28(%r8), %rax
   830: 420fb61418               movzbl (%rax, %r11), %edx
   835: 4885d2                   testq %rdx, %rdx
   838: 0f841f000000             je 0x36b
   844: 0fb60417                 movzbl (%rdi, %rdx), %eax
   848: 83c031                   addl $0x31, %eax
   851: 81f8ff000000             cmpl $0xff, %eax
   857: 0f8206000000             jb 0x365
   863: 81c001ffffff             addl $0xffffff01, %eax
   869: 0fb60406                 movzbl (%rsi, %rax), %eax
   873: 32d8                     xorb %al, %bl
   875: 43881c1a                 movb %bl, (%r10, %r11)
   879: 33db                     xorl %ebx, %ebx
   881: 49ffc3                   incq %r11
   884: 4c3bd9                   cmpq %rcx, %r11
   887: 0f82c0feffff             jb 0x23d
   893: 4d8b5108                 movq 8(%r9), %r10
   897: 4c8b1c24                 movq (%rsp), %r11
   901: 33db                     xorl %ebx, %ebx
   903: 498b00                   movq (%r8), %rax
   906: 420fb61418               movzbl (%rax, %r11), %edx
   911: 4885d2                   testq %rdx, %rdx
   914: 0f841f000000             je 0x3b7
   920: 0fb60417                 movzbl (%rdi, %rdx), %eax
   924: 83c01c                   addl $0x1c, %eax
   927: 81f8ff000000             cmpl $0xff, %eax
   933: 0f8206000000             jb 0x3b1
   939: 81c001ffffff             addl $0xffffff01, %eax
   945: 0fb60406                 movzbl (%rsi, %rax), %eax
   949: 32d8                     xorb %al, %bl
   951: 498b4008                 movq 8(%r8), %rax
   955: 420fb61418               movzbl (%rax, %r11), %edx
   960: 4885d2                   testq %rdx, %rdx
   963: 0f8422000000             je 0x3eb
   969: 0fb60417                 movzbl (%rdi, %rdx), %eax
   973: 81c0df000000             addl $0xdf, %eax
   979: 81f8ff000000             cmpl $0xff, %eax
   985: 0f8206000000             jb 0x3e5
   991: 81c001ffffff             addl $0xffffff01, %eax
   997: 0fb60406                 movzbl (%rsi, %rax), %eax
  1001: 32d8                     xorb %al, %bl
  1003: 498b4010                 movq 0x10(%r8), %rax
  1007: 420fb61418               movzbl (%rax, %r11), %edx
  1012: 4885d2                   testq %rdx, %rdx
  1015: 0f841f000000             je 0x41c
  1021: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1025: 83c01a                   addl $0x1a, %eax
  1028: 81f8ff000000             cmpl $0xff, %eax
  1034: 0f8206000000             jb 0x416
  1040: 81c001ffffff             addl $0xffffff01, %eax
  1046: 0fb60406                 movzbl (%rsi, %rax), %eax
  1050: 32d8                     xorb %al, %bl
  1052: 498b4018                 movq 0x18(%r8), %rax
  1056: 420fb61418               movzbl (%rax, %r11), %edx
  1061: 4885d2                   testq %rdx, %rdx
  1064: 0f841f000000             je 0x44d
  1070: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1074: 83c019                   addl $0x19, %eax
  1077: 81f8ff000000             cmpl $0xff, %eax
  1083: 0f8206000000             jb 0x447
  1089: 81c001ffffff             addl $0xffffff01, %eax
  1095: 0fb60406                 movzbl (%rsi, %rax), %eax
  1099: 32d8                     xorb %al, %bl
  1101: 498b4020                 movq 0x20(%r8), %rax
  1105: 420fb61418               movzbl (%rax, %r11), %edx
  1110: 4885d2                   testq %rdx, %rdx
  1113: 0f841f000000             je 0x47e
  1119: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1123: 83c018                   addl $0x18, %eax
  1126: 81f8ff000000             cmpl $0xff, %eax
  1132: 0f8206000000             jb 0x478
  1138: 81c001ffffff             addl $0xffffff01, %eax
  1144: 0fb60406                 movzbl (%rsi, %rax), %eax
  1148: 32d8                     xorb %al, %bl
  1150: 498b4028                 movq 0x28(%r8), %rax
  1154: 420fb61418               movzbl (%rax, %r11), %edx
  1159: 4885d2                   testq %rdx, %rdx
  1162: 0f841f000000             je 0x4af
  1168: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1172: 83c031                   addl $0x31, %eax
  1175: 81f8ff000000             cmpl $0xff, %eax
  1181: 0f8206000000             jb 0x4a9
  1187: 81c001ffffff             addl $0xffffff01, %eax
  1193: 0fb60406                 movzbl (%rsi, %rax), %eax
  1197: 32d8                     xorb %al, %bl
  1199: 43881c1a                 movb %bl, (%r10, %r11)
  1203: 33db                     xorl %ebx, %ebx
  1205: 49ffc3                   incq %r11
  1208: 4c3bd9                   cmpq %rcx, %r11
  1211: 0f82c6feffff             jb 0x387
  1217: 4883c408                 addq $8, %rsp
  1221: 415e                     popq %r14
  1223: 415d                     popq %r13
  1225: 415c                     popq %r12
  1227: 415f                     popq %r15
  1229: 5b                       popq %rbx
  1230: c5f877                   vzeroupper 
  1233: c3                       retq 
  1234: 0f0b                     ud2 
  1236: 0f0b                     ud2 
  1238: 0f0b                     ud2 
  1240: 0f0b                     ud2 
  1242: 0f0b                     ud2 
  1244: 0f0b                     ud2 
  1246: 0f0b                     ud2 
  1248: 0f0b                     ud2 
  1250: 0f0b                     ud2 
  1252: 0f0b                     ud2 
  1254: 0f0b                     ud2 
  1256: 0f0b                     ud2 
  1258: 0f0b                     ud2 
  1260: 0f0b                     ud2 
  1262: 0f0b                     ud2 
  1264: 0f0b                     ud2 
  1266: 0f0b                     ud2 
  1268: 0f0b                     ud2 
  1270: 0f0b                     ud2 
  1272: 0f0b                     ud2 
  1274: 0f0b                     ud2 
  1276: 0f0b                     ud2 
  1278: 0f0b                     ud2 
