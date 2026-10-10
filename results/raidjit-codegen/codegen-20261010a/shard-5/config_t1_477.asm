; config_t1_477 target=avx2 twin=xorbit topo=rs(6,3)/rec{0,1} op=recover U=1
; code_size=3456 twins=["xorbit"] uops_per_byte=17.875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d670d0000           leaq 0xd67(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: c4c17d6f2f               vmovdqa (%r15), %ymm5
    52: 488bc1                   movq %rcx, %rax
    55: 4883e820                 subq $0x20, %rax
    59: 0f825e0a0000             jb 0xa9f
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   102: c5e571d204               vpsrlw $4, %ymm2, %ymm3
   107: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   111: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   117: c5dd71d205               vpsrlw $5, %ymm2, %ymm4
   122: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   126: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   130: c5fd6fc3                 vmovdqa %ymm3, %ymm0
   134: c5fd6fcb                 vmovdqa %ymm3, %ymm1
   138: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   143: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   147: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   151: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   156: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   160: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   164: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   169: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   173: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   178: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   182: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   187: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   191: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   195: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   200: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   204: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   208: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   214: c5e571d200               vpsrlw $0, %ymm2, %ymm3
   219: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   223: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
   228: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   232: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   236: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   240: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   245: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   249: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   254: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   258: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   262: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   267: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   271: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   275: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   280: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   284: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   288: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   293: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   297: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   301: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   306: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   310: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   314: c5e571d202               vpsrlw $2, %ymm2, %ymm3
   319: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   323: 4d8b5028                 movq 0x28(%r8), %r10
   327: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   333: c5dd71d200               vpsrlw $0, %ymm2, %ymm4
   338: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   342: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   346: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   350: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   354: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   359: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   363: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   368: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   372: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   376: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   381: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   385: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   389: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   394: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   398: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   402: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   407: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   411: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   415: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   421: c5e571d203               vpsrlw $3, %ymm2, %ymm3
   426: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   430: 4d8b5028                 movq 0x28(%r8), %r10
   434: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   440: c5dd71d202               vpsrlw $2, %ymm2, %ymm4
   445: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   449: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   453: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   457: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   461: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   466: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   470: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   475: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   479: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   483: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   488: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   492: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   496: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   501: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   505: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   509: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   514: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   518: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   522: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   528: c5e571d205               vpsrlw $5, %ymm2, %ymm3
   533: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   537: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   543: c5dd71d206               vpsrlw $6, %ymm2, %ymm4
   548: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   552: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   556: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   560: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   564: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   569: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   573: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   577: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   582: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   586: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   590: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   595: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   599: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   604: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   608: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   613: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   617: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   621: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   627: c5e571d204               vpsrlw $4, %ymm2, %ymm3
   632: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   636: 4d8b5028                 movq 0x28(%r8), %r10
   640: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   646: c5dd71d203               vpsrlw $3, %ymm2, %ymm4
   651: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   655: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   659: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   663: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   667: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   672: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   676: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   680: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   685: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   689: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   693: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   698: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   702: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   706: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   711: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   715: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   719: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   725: c5e571d206               vpsrlw $6, %ymm2, %ymm3
   730: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   734: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   741: c5dd71d207               vpsrlw $7, %ymm2, %ymm4
   746: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   750: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   754: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   758: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   762: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   767: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   771: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   775: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   780: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   784: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   788: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   793: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   797: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   801: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   806: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   810: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   816: c5e571d203               vpsrlw $3, %ymm2, %ymm3
   821: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   825: 4d8b5028                 movq 0x28(%r8), %r10
   829: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   835: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
   840: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   844: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   848: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   853: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   857: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   861: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   866: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   870: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   874: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   879: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   883: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   887: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   892: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   896: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   900: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   906: c5e571d205               vpsrlw $5, %ymm2, %ymm3
   911: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   915: 4d8b5028                 movq 0x28(%r8), %r10
   919: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   925: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   930: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   934: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   938: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   942: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   946: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   951: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   955: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   959: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   964: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   968: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   972: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   977: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   981: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   985: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   991: c5e571d200               vpsrlw $0, %ymm2, %ymm3
   996: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1000: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
  1005: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1009: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1013: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  1018: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1022: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1027: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1031: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1035: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1040: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1044: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1048: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1053: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1057: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1061: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  1066: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1070: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
  1076: c5dd71d202               vpsrlw $2, %ymm2, %ymm4
  1081: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1085: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1089: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1094: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1098: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1103: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1107: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1111: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1116: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1120: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1124: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1129: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1133: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1137: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
  1143: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  1148: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1152: 4d8b5028                 movq 0x28(%r8), %r10
  1156: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
  1162: c5dd71d205               vpsrlw $5, %ymm2, %ymm4
  1167: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1171: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1175: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  1179: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
  1183: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1188: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1192: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1196: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1201: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1205: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1209: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
  1215: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  1220: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1224: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
  1229: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1233: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1237: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  1242: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1246: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1251: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1255: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1259: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1264: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1268: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1272: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  1277: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1281: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
  1286: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1290: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1294: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1299: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1303: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1308: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1312: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1316: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1321: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1325: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1329: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  1335: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  1340: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1344: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
  1350: c5dd71d207               vpsrlw $7, %ymm2, %ymm4
  1355: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1359: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1363: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  1367: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
  1371: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1376: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1380: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1384: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1389: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1393: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
  1399: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  1404: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1408: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
  1414: c5dd71d207               vpsrlw $7, %ymm2, %ymm4
  1419: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1423: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1427: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  1431: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
  1435: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1440: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1444: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1449: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1453: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
  1459: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  1464: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1468: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  1474: c5dd71d202               vpsrlw $2, %ymm2, %ymm4
  1479: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1483: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1487: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1492: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1496: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1500: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1505: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1509: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1513: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  1518: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1522: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
  1527: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1531: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1535: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  1540: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1544: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1549: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1553: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1557: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  1562: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1566: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
  1571: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1575: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1579: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1584: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1588: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1593: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1597: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1601: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  1606: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1610: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  1617: c5dd71d205               vpsrlw $5, %ymm2, %ymm4
  1622: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1626: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1630: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1635: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1639: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1644: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1648: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1652: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
  1658: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  1663: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1667: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
  1673: c5dd71d200               vpsrlw $0, %ymm2, %ymm4
  1678: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1682: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1686: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  1690: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  1695: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1699: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
  1705: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  1710: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1714: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  1720: c5dd71d207               vpsrlw $7, %ymm2, %ymm4
  1725: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1729: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1733: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1738: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1742: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1746: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
  1752: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  1757: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1761: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  1767: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
  1772: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1776: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1780: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1785: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1789: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1793: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
  1799: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  1804: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1808: 4d8b5028                 movq 0x28(%r8), %r10
  1812: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
  1818: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
  1823: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1827: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1831: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  1835: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
  1839: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
  1845: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  1850: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1854: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  1861: c5dd71d203               vpsrlw $3, %ymm2, %ymm4
  1866: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1870: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1874: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1879: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1883: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1887: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
  1893: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  1898: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1902: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  1908: c5dd71d203               vpsrlw $3, %ymm2, %ymm4
  1913: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1917: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1921: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1926: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1930: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1935: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1939: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
  1945: c5e571d204               vpsrlw $4, %ymm2, %ymm3
  1950: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1954: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  1960: c5dd71d205               vpsrlw $5, %ymm2, %ymm4
  1965: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1969: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1973: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1978: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1982: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1986: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  1991: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1995: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  2002: c5dd71d200               vpsrlw $0, %ymm2, %ymm4
  2007: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  2011: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  2015: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  2019: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  2024: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2028: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  2034: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  2039: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2043: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  2050: c5dd71d202               vpsrlw $2, %ymm2, %ymm4
  2055: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  2059: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  2063: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  2068: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2072: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2076: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  2082: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  2087: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2091: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  2098: c5dd71d206               vpsrlw $6, %ymm2, %ymm4
  2103: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  2107: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  2111: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  2116: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2120: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  2125: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2129: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  2135: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  2140: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2144: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  2151: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
  2156: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  2160: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  2164: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  2169: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2173: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2177: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
  2183: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  2188: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2192: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  2197: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2201: c5e571d204               vpsrlw $4, %ymm2, %ymm3
  2206: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2210: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  2215: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2219: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  2224: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2228: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  2233: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2237: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  2242: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2246: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  2251: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2255: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  2260: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2264: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  2269: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2273: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
  2279: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  2284: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2288: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  2293: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2297: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  2302: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2306: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  2311: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2315: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  2320: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2324: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  2329: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2333: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  2339: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  2344: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2348: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  2353: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2357: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  2364: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  2369: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2373: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  2378: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2382: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  2387: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2391: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  2396: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2400: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  2405: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2409: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  2414: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2418: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  2423: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2427: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  2432: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2436: c5e571d204               vpsrlw $4, %ymm2, %ymm3
  2441: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2445: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  2450: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2454: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  2459: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2463: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2467: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  2472: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2476: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  2481: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2485: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
  2491: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  2496: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2500: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  2505: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2509: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  2514: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2518: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  2522: c5e571d204               vpsrlw $4, %ymm2, %ymm3
  2527: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2531: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  2536: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2540: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  2545: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2549: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  2554: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2558: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  2563: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2567: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  2572: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2576: 4d8b5028                 movq 0x28(%r8), %r10
  2580: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
  2586: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  2591: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2595: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  2600: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2604: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  2609: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2613: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  2618: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2622: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  2627: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2631: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  2635: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
  2639: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  2644: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2648: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2652: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  2657: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2661: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  2666: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2670: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2674: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  2679: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2683: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2687: 4d8b11                   movq (%r9), %r10
  2690: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
  2696: 4d8b5108                 movq 8(%r9), %r10
  2700: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
  2706: 4983c320                 addq $0x20, %r11
  2710: 4c3bd8                   cmpq %rax, %r11
  2713: 0f86c1f5ffff             jbe 0x60
  2719: 4c3bd9                   cmpq %rcx, %r11
  2722: 0f839e020000             jae 0xd46
  2728: 4153                     pushq %r11
  2730: 498d7f40                 leaq 0x40(%r15), %rdi
  2734: 498db740010000           leaq 0x140(%r15), %rsi
  2741: 4d8b11                   movq (%r9), %r10
  2744: 4c8b1c24                 movq (%rsp), %r11
  2748: 33db                     xorl %ebx, %ebx
  2750: 498b00                   movq (%r8), %rax
  2753: 420fb61418               movzbl (%rax, %r11), %edx
  2758: 4885d2                   testq %rdx, %rdx
  2761: 0f8422000000             je 0xaf1
  2767: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2771: 81c0c1000000             addl $0xc1, %eax
  2777: 81f8ff000000             cmpl $0xff, %eax
  2783: 0f8206000000             jb 0xaeb
  2789: 81c001ffffff             addl $0xffffff01, %eax
  2795: 0fb60406                 movzbl (%rsi, %rax), %eax
  2799: 32d8                     xorb %al, %bl
  2801: 498b4008                 movq 8(%r8), %rax
  2805: 420fb61418               movzbl (%rax, %r11), %edx
  2810: 4885d2                   testq %rdx, %rdx
  2813: 0f841f000000             je 0xb22
  2819: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2823: 83c003                   addl $3, %eax
  2826: 81f8ff000000             cmpl $0xff, %eax
  2832: 0f8206000000             jb 0xb1c
  2838: 81c001ffffff             addl $0xffffff01, %eax
  2844: 0fb60406                 movzbl (%rsi, %rax), %eax
  2848: 32d8                     xorb %al, %bl
  2850: 498b4010                 movq 0x10(%r8), %rax
  2854: 420fb61418               movzbl (%rax, %r11), %edx
  2859: 4885d2                   testq %rdx, %rdx
  2862: 0f8422000000             je 0xb56
  2868: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2872: 81c0c6000000             addl $0xc6, %eax
  2878: 81f8ff000000             cmpl $0xff, %eax
  2884: 0f8206000000             jb 0xb50
  2890: 81c001ffffff             addl $0xffffff01, %eax
  2896: 0fb60406                 movzbl (%rsi, %rax), %eax
  2900: 32d8                     xorb %al, %bl
  2902: 498b4018                 movq 0x18(%r8), %rax
  2906: 420fb61418               movzbl (%rax, %r11), %edx
  2911: 4885d2                   testq %rdx, %rdx
  2914: 0f841f000000             je 0xb87
  2920: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2924: 83c001                   addl $1, %eax
  2927: 81f8ff000000             cmpl $0xff, %eax
  2933: 0f8206000000             jb 0xb81
  2939: 81c001ffffff             addl $0xffffff01, %eax
  2945: 0fb60406                 movzbl (%rsi, %rax), %eax
  2949: 32d8                     xorb %al, %bl
  2951: 498b4020                 movq 0x20(%r8), %rax
  2955: 420fb61418               movzbl (%rax, %r11), %edx
  2960: 4885d2                   testq %rdx, %rdx
  2963: 0f8422000000             je 0xbbb
  2969: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2973: 81c0fe000000             addl $0xfe, %eax
  2979: 81f8ff000000             cmpl $0xff, %eax
  2985: 0f8206000000             jb 0xbb5
  2991: 81c001ffffff             addl $0xffffff01, %eax
  2997: 0fb60406                 movzbl (%rsi, %rax), %eax
  3001: 32d8                     xorb %al, %bl
  3003: 498b4028                 movq 0x28(%r8), %rax
  3007: 420fb61418               movzbl (%rax, %r11), %edx
  3012: 4885d2                   testq %rdx, %rdx
  3015: 0f841f000000             je 0xbec
  3021: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3025: 83c031                   addl $0x31, %eax
  3028: 81f8ff000000             cmpl $0xff, %eax
  3034: 0f8206000000             jb 0xbe6
  3040: 81c001ffffff             addl $0xffffff01, %eax
  3046: 0fb60406                 movzbl (%rsi, %rax), %eax
  3050: 32d8                     xorb %al, %bl
  3052: 43881c1a                 movb %bl, (%r10, %r11)
  3056: 33db                     xorl %ebx, %ebx
  3058: 49ffc3                   incq %r11
  3061: 4c3bd9                   cmpq %rcx, %r11
  3064: 0f82c0feffff             jb 0xabe
  3070: 4d8b5108                 movq 8(%r9), %r10
  3074: 4c8b1c24                 movq (%rsp), %r11
  3078: 33db                     xorl %ebx, %ebx
  3080: 498b00                   movq (%r8), %rax
  3083: 420fb61418               movzbl (%rax, %r11), %edx
  3088: 4885d2                   testq %rdx, %rdx
  3091: 0f841f000000             je 0xc38
  3097: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3101: 83c01c                   addl $0x1c, %eax
  3104: 81f8ff000000             cmpl $0xff, %eax
  3110: 0f8206000000             jb 0xc32
  3116: 81c001ffffff             addl $0xffffff01, %eax
  3122: 0fb60406                 movzbl (%rsi, %rax), %eax
  3126: 32d8                     xorb %al, %bl
  3128: 498b4008                 movq 8(%r8), %rax
  3132: 420fb61418               movzbl (%rax, %r11), %edx
  3137: 4885d2                   testq %rdx, %rdx
  3140: 0f8422000000             je 0xc6c
  3146: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3150: 81c0df000000             addl $0xdf, %eax
  3156: 81f8ff000000             cmpl $0xff, %eax
  3162: 0f8206000000             jb 0xc66
  3168: 81c001ffffff             addl $0xffffff01, %eax
  3174: 0fb60406                 movzbl (%rsi, %rax), %eax
  3178: 32d8                     xorb %al, %bl
  3180: 498b4010                 movq 0x10(%r8), %rax
  3184: 420fb61418               movzbl (%rax, %r11), %edx
  3189: 4885d2                   testq %rdx, %rdx
  3192: 0f841f000000             je 0xc9d
  3198: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3202: 83c01a                   addl $0x1a, %eax
  3205: 81f8ff000000             cmpl $0xff, %eax
  3211: 0f8206000000             jb 0xc97
  3217: 81c001ffffff             addl $0xffffff01, %eax
  3223: 0fb60406                 movzbl (%rsi, %rax), %eax
  3227: 32d8                     xorb %al, %bl
  3229: 498b4018                 movq 0x18(%r8), %rax
  3233: 420fb61418               movzbl (%rax, %r11), %edx
  3238: 4885d2                   testq %rdx, %rdx
  3241: 0f841f000000             je 0xcce
  3247: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3251: 83c019                   addl $0x19, %eax
  3254: 81f8ff000000             cmpl $0xff, %eax
  3260: 0f8206000000             jb 0xcc8
  3266: 81c001ffffff             addl $0xffffff01, %eax
  3272: 0fb60406                 movzbl (%rsi, %rax), %eax
  3276: 32d8                     xorb %al, %bl
  3278: 498b4020                 movq 0x20(%r8), %rax
  3282: 420fb61418               movzbl (%rax, %r11), %edx
  3287: 4885d2                   testq %rdx, %rdx
  3290: 0f841f000000             je 0xcff
  3296: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3300: 83c018                   addl $0x18, %eax
  3303: 81f8ff000000             cmpl $0xff, %eax
  3309: 0f8206000000             jb 0xcf9
  3315: 81c001ffffff             addl $0xffffff01, %eax
  3321: 0fb60406                 movzbl (%rsi, %rax), %eax
  3325: 32d8                     xorb %al, %bl
  3327: 498b4028                 movq 0x28(%r8), %rax
  3331: 420fb61418               movzbl (%rax, %r11), %edx
  3336: 4885d2                   testq %rdx, %rdx
  3339: 0f841f000000             je 0xd30
  3345: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3349: 83c031                   addl $0x31, %eax
  3352: 81f8ff000000             cmpl $0xff, %eax
  3358: 0f8206000000             jb 0xd2a
  3364: 81c001ffffff             addl $0xffffff01, %eax
  3370: 0fb60406                 movzbl (%rsi, %rax), %eax
  3374: 32d8                     xorb %al, %bl
  3376: 43881c1a                 movb %bl, (%r10, %r11)
  3380: 33db                     xorl %ebx, %ebx
  3382: 49ffc3                   incq %r11
  3385: 4c3bd9                   cmpq %rcx, %r11
  3388: 0f82c6feffff             jb 0xc08
  3394: 4883c408                 addq $8, %rsp
  3398: 415e                     popq %r14
  3400: 415d                     popq %r13
  3402: 415c                     popq %r12
  3404: 415f                     popq %r15
  3406: 5b                       popq %rbx
  3407: c5f877                   vzeroupper 
  3410: c3                       retq 
  3411: 0f0b                     ud2 
  3413: 0f0b                     ud2 
  3415: 0f0b                     ud2 
  3417: 0f0b                     ud2 
  3419: 0f0b                     ud2 
  3421: 0f0b                     ud2 
  3423: 0f0b                     ud2 
  3425: 0f0b                     ud2 
  3427: 0f0b                     ud2 
  3429: 0f0b                     ud2 
  3431: 0f0b                     ud2 
  3433: 0f0b                     ud2 
  3435: 0f0b                     ud2 
  3437: 0f0b                     ud2 
  3439: 0f0b                     ud2 
  3441: 0f0b                     ud2 
  3443: 0f0b                     ud2 
  3445: 0f0b                     ud2 
  3447: 0f0b                     ud2 
  3449: 0f0b                     ud2 
  3451: 0f0b                     ud2 
  3453: 0f0b                     ud2 
  3455: cc                       int3 
