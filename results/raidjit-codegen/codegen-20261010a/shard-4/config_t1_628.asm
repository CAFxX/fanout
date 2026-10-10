; config_t1_628 target=sse twin=notail topo=rdp(7)/enc op=encode U=1
; code_size=1152 twins=["notail"] uops_per_byte=11
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
    41: 4883e810                 subq $0x10, %rax
    45: 0f8207040000             jb 0x43a
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 66460f6f241e             movdqa (%rsi, %r11), %xmm12
    70: 66410f6fc4               movdqa %xmm12, %xmm0
    75: 66410f6ff4               movdqa %xmm12, %xmm6
    80: 66460f6f241f             movdqa (%rdi, %r11), %xmm12
    86: 66410f6fcc               movdqa %xmm12, %xmm1
    91: 66410feff4               pxor %xmm12, %xmm6
    96: 66410f6ffc               movdqa %xmm12, %xmm7
   101: 66470f6f241c             movdqa (%r12, %r11), %xmm12
   107: 66410f6fd4               movdqa %xmm12, %xmm2
   112: 66410feffc               pxor %xmm12, %xmm7
   117: 66450f6fc4               movdqa %xmm12, %xmm8
   122: 66470f6f641d00           movdqa (%r13, %r11), %xmm12
   129: 66410f6fdc               movdqa %xmm12, %xmm3
   134: 66450fefc4               pxor %xmm12, %xmm8
   139: 66450f6fcc               movdqa %xmm12, %xmm9
   144: 66470f6f241e             movdqa (%r14, %r11), %xmm12
   150: 66410f6fe4               movdqa %xmm12, %xmm4
   155: 66450fefcc               pxor %xmm12, %xmm9
   160: 66450f6fd4               movdqa %xmm12, %xmm10
   165: 4d8b5028                 movq 0x28(%r8), %r10
   169: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   175: 66410f6fec               movdqa %xmm12, %xmm5
   180: 66450fefd4               pxor %xmm12, %xmm10
   185: 66450f6fdc               movdqa %xmm12, %xmm11
   190: 4d8b5030                 movq 0x30(%r8), %r10
   194: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   200: 66410fefc4               pxor %xmm12, %xmm0
   205: 66410feffc               pxor %xmm12, %xmm7
   210: 4d8b5038                 movq 0x38(%r8), %r10
   214: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   220: 66410fefcc               pxor %xmm12, %xmm1
   225: 66410feff4               pxor %xmm12, %xmm6
   230: 66450fefc4               pxor %xmm12, %xmm8
   235: 4d8b5040                 movq 0x40(%r8), %r10
   239: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   245: 66410fefd4               pxor %xmm12, %xmm2
   250: 66410feffc               pxor %xmm12, %xmm7
   255: 66450fefcc               pxor %xmm12, %xmm9
   260: 4d8b5048                 movq 0x48(%r8), %r10
   264: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   270: 66410fefdc               pxor %xmm12, %xmm3
   275: 66450fefc4               pxor %xmm12, %xmm8
   280: 66450fefd4               pxor %xmm12, %xmm10
   285: 4d8b5050                 movq 0x50(%r8), %r10
   289: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   295: 66410fefe4               pxor %xmm12, %xmm4
   300: 66450fefcc               pxor %xmm12, %xmm9
   305: 66450fefdc               pxor %xmm12, %xmm11
   310: 4d8b5058                 movq 0x58(%r8), %r10
   314: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   320: 66410fefec               pxor %xmm12, %xmm5
   325: 66450fefd4               pxor %xmm12, %xmm10
   330: 4d8b5060                 movq 0x60(%r8), %r10
   334: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   340: 66410fefc4               pxor %xmm12, %xmm0
   345: 66450fefc4               pxor %xmm12, %xmm8
   350: 4d8b5068                 movq 0x68(%r8), %r10
   354: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   360: 66410fefcc               pxor %xmm12, %xmm1
   365: 66410feff4               pxor %xmm12, %xmm6
   370: 66450fefcc               pxor %xmm12, %xmm9
   375: 4d8b5070                 movq 0x70(%r8), %r10
   379: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   385: 66410fefd4               pxor %xmm12, %xmm2
   390: 66410feffc               pxor %xmm12, %xmm7
   395: 66450fefd4               pxor %xmm12, %xmm10
   400: 4d8b5078                 movq 0x78(%r8), %r10
   404: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   410: 66410fefdc               pxor %xmm12, %xmm3
   415: 66450fefc4               pxor %xmm12, %xmm8
   420: 66450fefdc               pxor %xmm12, %xmm11
   425: 4d8b9080000000           movq 0x80(%r8), %r10
   432: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   438: 66410fefe4               pxor %xmm12, %xmm4
   443: 66450fefcc               pxor %xmm12, %xmm9
   448: 4d8b9088000000           movq 0x88(%r8), %r10
   455: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   461: 66410fefec               pxor %xmm12, %xmm5
   466: 66410feff4               pxor %xmm12, %xmm6
   471: 66450fefd4               pxor %xmm12, %xmm10
   476: 4d8b9090000000           movq 0x90(%r8), %r10
   483: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   489: 66410fefc4               pxor %xmm12, %xmm0
   494: 66450fefcc               pxor %xmm12, %xmm9
   499: 4d8b9098000000           movq 0x98(%r8), %r10
   506: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   512: 66410fefcc               pxor %xmm12, %xmm1
   517: 66410feff4               pxor %xmm12, %xmm6
   522: 66450fefd4               pxor %xmm12, %xmm10
   527: 4d8b90a0000000           movq 0xa0(%r8), %r10
   534: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   540: 66410fefd4               pxor %xmm12, %xmm2
   545: 66410feffc               pxor %xmm12, %xmm7
   550: 66450fefdc               pxor %xmm12, %xmm11
   555: 4d8b90a8000000           movq 0xa8(%r8), %r10
   562: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   568: 66410fefdc               pxor %xmm12, %xmm3
   573: 66450fefc4               pxor %xmm12, %xmm8
   578: 4d8b90b0000000           movq 0xb0(%r8), %r10
   585: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   591: 66410fefe4               pxor %xmm12, %xmm4
   596: 66410feff4               pxor %xmm12, %xmm6
   601: 66450fefcc               pxor %xmm12, %xmm9
   606: 4d8b90b8000000           movq 0xb8(%r8), %r10
   613: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   619: 66410fefec               pxor %xmm12, %xmm5
   624: 66410feffc               pxor %xmm12, %xmm7
   629: 66450fefd4               pxor %xmm12, %xmm10
   634: 4d8b90c0000000           movq 0xc0(%r8), %r10
   641: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   647: 66410fefc4               pxor %xmm12, %xmm0
   652: 66450fefd4               pxor %xmm12, %xmm10
   657: 4d8b90c8000000           movq 0xc8(%r8), %r10
   664: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   670: 66410fefcc               pxor %xmm12, %xmm1
   675: 66410feff4               pxor %xmm12, %xmm6
   680: 66450fefdc               pxor %xmm12, %xmm11
   685: 4d8b90d0000000           movq 0xd0(%r8), %r10
   692: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   698: 66410fefd4               pxor %xmm12, %xmm2
   703: 66410feffc               pxor %xmm12, %xmm7
   708: 4d8b90d8000000           movq 0xd8(%r8), %r10
   715: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   721: 66410fefdc               pxor %xmm12, %xmm3
   726: 66410feff4               pxor %xmm12, %xmm6
   731: 66450fefc4               pxor %xmm12, %xmm8
   736: 4d8b90e0000000           movq 0xe0(%r8), %r10
   743: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   749: 66410fefe4               pxor %xmm12, %xmm4
   754: 66410feffc               pxor %xmm12, %xmm7
   759: 66450fefcc               pxor %xmm12, %xmm9
   764: 4d8b90e8000000           movq 0xe8(%r8), %r10
   771: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   777: 66410fefec               pxor %xmm12, %xmm5
   782: 66450fefc4               pxor %xmm12, %xmm8
   787: 66450fefd4               pxor %xmm12, %xmm10
   792: 4d8b90f0000000           movq 0xf0(%r8), %r10
   799: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   805: 66410fefc4               pxor %xmm12, %xmm0
   810: 66450fefdc               pxor %xmm12, %xmm11
   815: 4d8b90f8000000           movq 0xf8(%r8), %r10
   822: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   828: 66410fefcc               pxor %xmm12, %xmm1
   833: 66410feff4               pxor %xmm12, %xmm6
   838: 4d8b9000010000           movq 0x100(%r8), %r10
   845: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   851: 66410fefd4               pxor %xmm12, %xmm2
   856: 66410feff4               pxor %xmm12, %xmm6
   861: 66410feffc               pxor %xmm12, %xmm7
   866: 4d8b9008010000           movq 0x108(%r8), %r10
   873: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   879: 66410fefdc               pxor %xmm12, %xmm3
   884: 66410feffc               pxor %xmm12, %xmm7
   889: 66450fefc4               pxor %xmm12, %xmm8
   894: 4d8b9010010000           movq 0x110(%r8), %r10
   901: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   907: 66410fefe4               pxor %xmm12, %xmm4
   912: 66450fefc4               pxor %xmm12, %xmm8
   917: 66450fefcc               pxor %xmm12, %xmm9
   922: 4d8b9018010000           movq 0x118(%r8), %r10
   929: 66470f6f241a             movdqa (%r10, %r11), %xmm12
   935: 66410fefec               pxor %xmm12, %xmm5
   940: 66450fefcc               pxor %xmm12, %xmm9
   945: 66450fefd4               pxor %xmm12, %xmm10
   950: 4d8b11                   movq (%r9), %r10
   953: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   959: 4d8b5108                 movq 8(%r9), %r10
   963: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
   969: 4d8b5110                 movq 0x10(%r9), %r10
   973: 66430f7f141a             movdqa %xmm2, (%r10, %r11)
   979: 4d8b5118                 movq 0x18(%r9), %r10
   983: 66430f7f1c1a             movdqa %xmm3, (%r10, %r11)
   989: 4d8b5120                 movq 0x20(%r9), %r10
   993: 66430f7f241a             movdqa %xmm4, (%r10, %r11)
   999: 4d8b5128                 movq 0x28(%r9), %r10
  1003: 66430f7f2c1a             movdqa %xmm5, (%r10, %r11)
  1009: 4d8b5130                 movq 0x30(%r9), %r10
  1013: 66430f7f341a             movdqa %xmm6, (%r10, %r11)
  1019: 4d8b5138                 movq 0x38(%r9), %r10
  1023: 66430f7f3c1a             movdqa %xmm7, (%r10, %r11)
  1029: 4d8b5140                 movq 0x40(%r9), %r10
  1033: 66470f7f041a             movdqa %xmm8, (%r10, %r11)
  1039: 4d8b5148                 movq 0x48(%r9), %r10
  1043: 66470f7f0c1a             movdqa %xmm9, (%r10, %r11)
  1049: 4d8b5150                 movq 0x50(%r9), %r10
  1053: 66470f7f141a             movdqa %xmm10, (%r10, %r11)
  1059: 4d8b5158                 movq 0x58(%r9), %r10
  1063: 66470f7f1c1a             movdqa %xmm11, (%r10, %r11)
  1069: 4983c310                 addq $0x10, %r11
  1073: 4c3bd8                   cmpq %rax, %r11
  1076: 0f8606fcffff             jbe 0x40
  1082: 415e                     popq %r14
  1084: 415d                     popq %r13
  1086: 415c                     popq %r12
  1088: 5b                       popq %rbx
  1089: c3                       retq 
  1090: 0f0b                     ud2 
  1092: 0f0b                     ud2 
  1094: 0f0b                     ud2 
  1096: 0f0b                     ud2 
  1098: 0f0b                     ud2 
  1100: 0f0b                     ud2 
  1102: 0f0b                     ud2 
  1104: 0f0b                     ud2 
  1106: 0f0b                     ud2 
  1108: 0f0b                     ud2 
  1110: 0f0b                     ud2 
  1112: 0f0b                     ud2 
  1114: 0f0b                     ud2 
  1116: 0f0b                     ud2 
  1118: 0f0b                     ud2 
  1120: 0f0b                     ud2 
  1122: 0f0b                     ud2 
  1124: 0f0b                     ud2 
  1126: 0f0b                     ud2 
  1128: 0f0b                     ud2 
  1130: 0f0b                     ud2 
  1132: 0f0b                     ud2 
  1134: 0f0b                     ud2 
  1136: 0f0b                     ud2 
  1138: 0f0b                     ud2 
  1140: 0f0b                     ud2 
  1142: 0f0b                     ud2 
  1144: 0f0b                     ud2 
  1146: 0f0b                     ud2 
  1148: 0f0b                     ud2 
  1150: 0f0b                     ud2 
