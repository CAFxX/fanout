; config_t1_342 target=avx2 twin=notail topo=rdp(7)/enc op=encode U=1
; code_size=1152 twins=["notail"] uops_per_byte=5.5
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
    41: 4883e820                 subq $0x20, %rax
    45: 0f8207040000             jb 0x43a
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: c4217d6f241e             vmovdqa (%rsi, %r11), %ymm12
    70: c4c17d6fc4               vmovdqa %ymm12, %ymm0
    75: c4c17d6ff4               vmovdqa %ymm12, %ymm6
    80: c4217d6f241f             vmovdqa (%rdi, %r11), %ymm12
    86: c4c17d6fcc               vmovdqa %ymm12, %ymm1
    91: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
    96: c4c17d6ffc               vmovdqa %ymm12, %ymm7
   101: c4017d6f241c             vmovdqa (%r12, %r11), %ymm12
   107: c4c17d6fd4               vmovdqa %ymm12, %ymm2
   112: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   117: c4417d6fc4               vmovdqa %ymm12, %ymm8
   122: c4017d6f641d00           vmovdqa (%r13, %r11), %ymm12
   129: c4c17d6fdc               vmovdqa %ymm12, %ymm3
   134: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   139: c4417d6fcc               vmovdqa %ymm12, %ymm9
   144: c4017d6f241e             vmovdqa (%r14, %r11), %ymm12
   150: c4c17d6fe4               vmovdqa %ymm12, %ymm4
   155: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   160: c4417d6fd4               vmovdqa %ymm12, %ymm10
   165: 4d8b5028                 movq 0x28(%r8), %r10
   169: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   175: c4c17d6fec               vmovdqa %ymm12, %ymm5
   180: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   185: c4417d6fdc               vmovdqa %ymm12, %ymm11
   190: 4d8b5030                 movq 0x30(%r8), %r10
   194: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   200: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   205: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   210: 4d8b5038                 movq 0x38(%r8), %r10
   214: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   220: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   225: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   230: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   235: 4d8b5040                 movq 0x40(%r8), %r10
   239: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   245: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   250: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   255: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   260: 4d8b5048                 movq 0x48(%r8), %r10
   264: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   270: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   275: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   280: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   285: 4d8b5050                 movq 0x50(%r8), %r10
   289: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   295: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   300: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   305: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   310: 4d8b5058                 movq 0x58(%r8), %r10
   314: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   320: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   325: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   330: 4d8b5060                 movq 0x60(%r8), %r10
   334: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   340: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   345: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   350: 4d8b5068                 movq 0x68(%r8), %r10
   354: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   360: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   365: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   370: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   375: 4d8b5070                 movq 0x70(%r8), %r10
   379: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   385: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   390: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   395: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   400: 4d8b5078                 movq 0x78(%r8), %r10
   404: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   410: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   415: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   420: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   425: 4d8b9080000000           movq 0x80(%r8), %r10
   432: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   438: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   443: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   448: 4d8b9088000000           movq 0x88(%r8), %r10
   455: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   461: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   466: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   471: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   476: 4d8b9090000000           movq 0x90(%r8), %r10
   483: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   489: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   494: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   499: 4d8b9098000000           movq 0x98(%r8), %r10
   506: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   512: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   517: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   522: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   527: 4d8b90a0000000           movq 0xa0(%r8), %r10
   534: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   540: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   545: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   550: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   555: 4d8b90a8000000           movq 0xa8(%r8), %r10
   562: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   568: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   573: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   578: 4d8b90b0000000           movq 0xb0(%r8), %r10
   585: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   591: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   596: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   601: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   606: 4d8b90b8000000           movq 0xb8(%r8), %r10
   613: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   619: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   624: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   629: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   634: 4d8b90c0000000           movq 0xc0(%r8), %r10
   641: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   647: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   652: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   657: 4d8b90c8000000           movq 0xc8(%r8), %r10
   664: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   670: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   675: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   680: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   685: 4d8b90d0000000           movq 0xd0(%r8), %r10
   692: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   698: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   703: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   708: 4d8b90d8000000           movq 0xd8(%r8), %r10
   715: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   721: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   726: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   731: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   736: 4d8b90e0000000           movq 0xe0(%r8), %r10
   743: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   749: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   754: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   759: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   764: 4d8b90e8000000           movq 0xe8(%r8), %r10
   771: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   777: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   782: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   787: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   792: 4d8b90f0000000           movq 0xf0(%r8), %r10
   799: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   805: c4c17defc4               vpxor %ymm12, %ymm0, %ymm0
   810: c44125efdc               vpxor %ymm12, %ymm11, %ymm11
   815: 4d8b90f8000000           movq 0xf8(%r8), %r10
   822: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   828: c4c175efcc               vpxor %ymm12, %ymm1, %ymm1
   833: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   838: 4d8b9000010000           movq 0x100(%r8), %r10
   845: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   851: c4c16defd4               vpxor %ymm12, %ymm2, %ymm2
   856: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   861: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   866: 4d8b9008010000           movq 0x108(%r8), %r10
   873: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   879: c4c165efdc               vpxor %ymm12, %ymm3, %ymm3
   884: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   889: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   894: 4d8b9010010000           movq 0x110(%r8), %r10
   901: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   907: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   912: c4413defc4               vpxor %ymm12, %ymm8, %ymm8
   917: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   922: 4d8b9018010000           movq 0x118(%r8), %r10
   929: c4017d6f241a             vmovdqa (%r10, %r11), %ymm12
   935: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   940: c44135efcc               vpxor %ymm12, %ymm9, %ymm9
   945: c4412defd4               vpxor %ymm12, %ymm10, %ymm10
   950: 4d8b11                   movq (%r9), %r10
   953: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   959: 4d8b5108                 movq 8(%r9), %r10
   963: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
   969: 4d8b5110                 movq 0x10(%r9), %r10
   973: c4817d7f141a             vmovdqa %ymm2, (%r10, %r11)
   979: 4d8b5118                 movq 0x18(%r9), %r10
   983: c4817d7f1c1a             vmovdqa %ymm3, (%r10, %r11)
   989: 4d8b5120                 movq 0x20(%r9), %r10
   993: c4817d7f241a             vmovdqa %ymm4, (%r10, %r11)
   999: 4d8b5128                 movq 0x28(%r9), %r10
  1003: c4817d7f2c1a             vmovdqa %ymm5, (%r10, %r11)
  1009: 4d8b5130                 movq 0x30(%r9), %r10
  1013: c4817d7f341a             vmovdqa %ymm6, (%r10, %r11)
  1019: 4d8b5138                 movq 0x38(%r9), %r10
  1023: c4817d7f3c1a             vmovdqa %ymm7, (%r10, %r11)
  1029: 4d8b5140                 movq 0x40(%r9), %r10
  1033: c4017d7f041a             vmovdqa %ymm8, (%r10, %r11)
  1039: 4d8b5148                 movq 0x48(%r9), %r10
  1043: c4017d7f0c1a             vmovdqa %ymm9, (%r10, %r11)
  1049: 4d8b5150                 movq 0x50(%r9), %r10
  1053: c4017d7f141a             vmovdqa %ymm10, (%r10, %r11)
  1059: 4d8b5158                 movq 0x58(%r9), %r10
  1063: c4017d7f1c1a             vmovdqa %ymm11, (%r10, %r11)
  1069: 4983c320                 addq $0x20, %r11
  1073: 4c3bd8                   cmpq %rax, %r11
  1076: 0f8606fcffff             jbe 0x40
  1082: 415e                     popq %r14
  1084: 415d                     popq %r13
  1086: 415c                     popq %r12
  1088: 5b                       popq %rbx
  1089: c5f877                   vzeroupper 
  1092: c3                       retq 
  1093: 0f0b                     ud2 
  1095: 0f0b                     ud2 
  1097: 0f0b                     ud2 
  1099: 0f0b                     ud2 
  1101: 0f0b                     ud2 
  1103: 0f0b                     ud2 
  1105: 0f0b                     ud2 
  1107: 0f0b                     ud2 
  1109: 0f0b                     ud2 
  1111: 0f0b                     ud2 
  1113: 0f0b                     ud2 
  1115: 0f0b                     ud2 
  1117: 0f0b                     ud2 
  1119: 0f0b                     ud2 
  1121: 0f0b                     ud2 
  1123: 0f0b                     ud2 
  1125: 0f0b                     ud2 
  1127: 0f0b                     ud2 
  1129: 0f0b                     ud2 
  1131: 0f0b                     ud2 
  1133: 0f0b                     ud2 
  1135: 0f0b                     ud2 
  1137: 0f0b                     ud2 
  1139: 0f0b                     ud2 
  1141: 0f0b                     ud2 
  1143: 0f0b                     ud2 
  1145: 0f0b                     ud2 
  1147: 0f0b                     ud2 
  1149: 0f0b                     ud2 
  1151: cc                       int3 
