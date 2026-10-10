; config_t1_321 target=avx2 twin=xorbit topo=raid6(8)/enc op=encode U=1
; code_size=2880 twins=["xorbit"] uops_per_byte=14.96875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d270b0000           leaq 0xb27(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: c4c17d6f2f               vmovdqa (%r15), %ymm5
    52: 488bc1                   movq %rcx, %rax
    55: 4883e820                 subq $0x20, %rax
    59: 0f82d9080000             jb 0x91a
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   102: c5e571d200               vpsrlw $0, %ymm2, %ymm3
   107: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   111: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   117: c5dd71d200               vpsrlw $0, %ymm2, %ymm4
   122: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   126: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   130: c5fd6fc3                 vmovdqa %ymm3, %ymm0
   134: c5fd6fcb                 vmovdqa %ymm3, %ymm1
   138: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   143: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   147: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   152: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   156: c5e571d201               vpsrlw $1, %ymm2, %ymm3
   161: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   165: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   172: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
   177: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   181: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   185: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   190: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   194: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   199: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   203: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   208: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   212: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   217: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   221: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   227: c5e571d205               vpsrlw $5, %ymm2, %ymm3
   232: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   236: 4d8b5028                 movq 0x28(%r8), %r10
   240: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   246: c5dd71d206               vpsrlw $6, %ymm2, %ymm4
   251: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   255: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   259: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   263: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   268: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   272: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   277: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   281: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   286: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   290: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   296: c5e571d206               vpsrlw $6, %ymm2, %ymm3
   301: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   305: 4d8b5028                 movq 0x28(%r8), %r10
   309: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
   315: c5dd71d207               vpsrlw $7, %ymm2, %ymm4
   320: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   324: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   328: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   333: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   337: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   342: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   346: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   351: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   355: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   360: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   364: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   371: c5e571d207               vpsrlw $7, %ymm2, %ymm3
   376: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   380: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
   386: c5dd71d207               vpsrlw $7, %ymm2, %ymm4
   391: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   395: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   399: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   404: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   408: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   413: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   417: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   422: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   426: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   432: c5e571d201               vpsrlw $1, %ymm2, %ymm3
   437: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   441: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   447: c5dd71d201               vpsrlw $1, %ymm2, %ymm4
   452: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   456: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   460: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   464: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   469: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   473: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   479: c5e571d202               vpsrlw $2, %ymm2, %ymm3
   484: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   488: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   494: c5dd71d202               vpsrlw $2, %ymm2, %ymm4
   499: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   503: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   507: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   512: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   516: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   521: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   525: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   531: c5e571d203               vpsrlw $3, %ymm2, %ymm3
   536: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   540: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   546: c5dd71d203               vpsrlw $3, %ymm2, %ymm4
   551: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   555: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   559: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   564: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   568: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   573: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   577: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   583: c5e571d204               vpsrlw $4, %ymm2, %ymm3
   588: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   592: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   598: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   603: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   607: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   611: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   616: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   620: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   625: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   629: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   635: c5e571d205               vpsrlw $5, %ymm2, %ymm3
   640: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   644: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   650: c5dd71d205               vpsrlw $5, %ymm2, %ymm4
   655: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   659: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   663: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   668: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   672: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   677: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   681: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   687: c5e571d206               vpsrlw $6, %ymm2, %ymm3
   692: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   696: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   702: c5dd71d206               vpsrlw $6, %ymm2, %ymm4
   707: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   711: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   715: c5dd71f305               vpsllw $5, %ymm3, %ymm4
   720: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   724: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   729: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   733: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
   739: c5e571d207               vpsrlw $7, %ymm2, %ymm3
   744: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   748: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
   754: c5dd71d207               vpsrlw $7, %ymm2, %ymm4
   759: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   763: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   767: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   772: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   776: c5dd71f307               vpsllw $7, %ymm3, %ymm4
   781: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   785: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   791: c5e571d200               vpsrlw $0, %ymm2, %ymm3
   796: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   800: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   807: c5dd71d200               vpsrlw $0, %ymm2, %ymm4
   812: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   816: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   820: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   825: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   829: c5dd71f306               vpsllw $6, %ymm3, %ymm4
   834: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   838: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   844: c5e571d202               vpsrlw $2, %ymm2, %ymm3
   849: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   853: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   860: c5dd71d202               vpsrlw $2, %ymm2, %ymm4
   865: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   869: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   873: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
   877: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   882: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   886: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   892: c5e571d203               vpsrlw $3, %ymm2, %ymm3
   897: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   901: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   908: c5dd71d203               vpsrlw $3, %ymm2, %ymm4
   913: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   917: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   921: c5dd71f301               vpsllw $1, %ymm3, %ymm4
   926: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   930: c5dd71f303               vpsllw $3, %ymm3, %ymm4
   935: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   939: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   945: c5e571d204               vpsrlw $4, %ymm2, %ymm3
   950: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
   954: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
   961: c5dd71d204               vpsrlw $4, %ymm2, %ymm4
   966: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
   970: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
   974: c5dd71f302               vpsllw $2, %ymm3, %ymm4
   979: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   983: c5dd71f304               vpsllw $4, %ymm3, %ymm4
   988: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   992: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
   998: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  1003: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1007: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  1014: c5dd71d205               vpsrlw $5, %ymm2, %ymm4
  1019: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1023: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1027: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1032: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1036: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1041: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1045: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
  1051: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  1056: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1060: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  1067: c5dd71d206               vpsrlw $6, %ymm2, %ymm4
  1072: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1076: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1080: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1085: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1089: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1094: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1098: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  1104: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  1109: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1113: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  1120: c5dd71d200               vpsrlw $0, %ymm2, %ymm4
  1125: c5dddbe5                 vpand %ymm5, %ymm4, %ymm4
  1129: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1133: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  1137: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1142: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1146: c4a17d6f141e             vmovdqa (%rsi, %r11), %ymm2
  1152: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  1157: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1161: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1166: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1170: c4a17d6f141f             vmovdqa (%rdi, %r11), %ymm2
  1176: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  1181: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1185: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
  1189: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  1194: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1198: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1203: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1207: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1212: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1216: c4817d6f141c             vmovdqa (%r12, %r11), %ymm2
  1222: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  1227: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1231: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1236: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1240: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1245: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1249: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  1254: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1258: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  1263: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1267: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  1272: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1276: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1281: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1285: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  1290: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1294: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1299: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1303: c5e571d204               vpsrlw $4, %ymm2, %ymm3
  1308: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1312: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1317: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1321: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  1326: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1330: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1335: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1339: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  1344: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1348: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1353: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1357: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  1362: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1366: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1371: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1375: c4817d6f541d00           vmovdqa (%r13, %r11), %ymm2
  1382: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  1387: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1391: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  1396: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1400: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  1405: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1409: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1414: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1418: c5e571d204               vpsrlw $4, %ymm2, %ymm3
  1423: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1427: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1432: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1436: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  1441: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1445: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1450: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1454: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  1459: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1463: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1468: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1472: c4817d6f141e             vmovdqa (%r14, %r11), %ymm2
  1478: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  1483: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1487: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  1491: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1496: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1500: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  1505: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1509: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  1514: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1518: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1523: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1527: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  1532: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1536: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1541: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1545: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1550: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1554: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  1559: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1563: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1568: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1572: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1577: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1581: c5e571d204               vpsrlw $4, %ymm2, %ymm3
  1586: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1590: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1595: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1599: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1604: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1608: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  1613: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1617: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1622: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1626: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  1631: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1635: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1640: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1644: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  1649: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1653: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1658: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1662: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1667: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1671: 4d8b5028                 movq 0x28(%r8), %r10
  1675: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
  1681: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  1686: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1690: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  1694: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1699: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1703: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  1708: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1712: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  1717: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1721: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1726: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1730: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  1735: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1739: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1744: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1748: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1753: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1757: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  1762: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1766: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1771: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1775: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1780: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1784: c5e571d204               vpsrlw $4, %ymm2, %ymm3
  1789: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1793: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1798: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1802: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1807: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1811: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  1816: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1820: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  1825: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1829: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1834: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1838: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  1843: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1847: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  1852: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1856: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  1861: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1865: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  1870: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1874: 4d8b5030                 movq 0x30(%r8), %r10
  1878: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
  1884: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  1889: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1893: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  1897: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  1902: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1906: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  1911: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1915: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  1920: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1924: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1929: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1933: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  1938: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1942: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  1947: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1951: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1956: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1960: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  1965: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1969: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  1974: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1978: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  1983: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1987: c5e571d204               vpsrlw $4, %ymm2, %ymm3
  1992: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  1996: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  2001: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2005: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  2010: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2014: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  2019: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2023: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  2028: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2032: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  2037: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2041: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  2046: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2050: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  2055: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2059: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  2064: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2068: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  2073: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2077: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
  2081: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  2086: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2090: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  2095: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2099: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  2104: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2108: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  2113: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2117: 4d8b5038                 movq 0x38(%r8), %r10
  2121: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
  2127: c5e571d200               vpsrlw $0, %ymm2, %ymm3
  2132: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2136: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
  2140: c5f5efcb                 vpxor %ymm3, %ymm1, %ymm1
  2144: c5e571d201               vpsrlw $1, %ymm2, %ymm3
  2149: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2153: c5dd71f301               vpsllw $1, %ymm3, %ymm4
  2158: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2162: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2166: c5e571d202               vpsrlw $2, %ymm2, %ymm3
  2171: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2175: c5dd71f302               vpsllw $2, %ymm3, %ymm4
  2180: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2184: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2188: c5e571d203               vpsrlw $3, %ymm2, %ymm3
  2193: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2197: c5dd71f303               vpsllw $3, %ymm3, %ymm4
  2202: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2206: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2210: c5e571d204               vpsrlw $4, %ymm2, %ymm3
  2215: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2219: c5dd71f304               vpsllw $4, %ymm3, %ymm4
  2224: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2228: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2232: c5e571d205               vpsrlw $5, %ymm2, %ymm3
  2237: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2241: c5dd71f305               vpsllw $5, %ymm3, %ymm4
  2246: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2250: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2254: c5e571d206               vpsrlw $6, %ymm2, %ymm3
  2259: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2263: c5dd71f306               vpsllw $6, %ymm3, %ymm4
  2268: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2272: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2276: c5e571d207               vpsrlw $7, %ymm2, %ymm3
  2281: c5e5dbdd                 vpand %ymm5, %ymm3, %ymm3
  2285: c5dd71f307               vpsllw $7, %ymm3, %ymm4
  2290: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2294: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2298: 4d8b11                   movq (%r9), %r10
  2301: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
  2307: 4d8b5108                 movq 8(%r9), %r10
  2311: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
  2317: 4983c320                 addq $0x20, %r11
  2321: 4c3bd8                   cmpq %rax, %r11
  2324: 0f8646f7ffff             jbe 0x60
  2330: 4c3bd9                   cmpq %rcx, %r11
  2333: 0f830c020000             jae 0xb2f
  2339: 4153                     pushq %r11
  2341: 498d7f40                 leaq 0x40(%r15), %rdi
  2345: 498db740010000           leaq 0x140(%r15), %rsi
  2352: 4d8b11                   movq (%r9), %r10
  2355: 4c8b1c24                 movq (%rsp), %r11
  2359: 33db                     xorl %ebx, %ebx
  2361: 498b00                   movq (%r8), %rax
  2364: 420fb61418               movzbl (%rax, %r11), %edx
  2369: 32da                     xorb %dl, %bl
  2371: 498b4008                 movq 8(%r8), %rax
  2375: 420fb61418               movzbl (%rax, %r11), %edx
  2380: 32da                     xorb %dl, %bl
  2382: 498b4010                 movq 0x10(%r8), %rax
  2386: 420fb61418               movzbl (%rax, %r11), %edx
  2391: 32da                     xorb %dl, %bl
  2393: 498b4018                 movq 0x18(%r8), %rax
  2397: 420fb61418               movzbl (%rax, %r11), %edx
  2402: 32da                     xorb %dl, %bl
  2404: 498b4020                 movq 0x20(%r8), %rax
  2408: 420fb61418               movzbl (%rax, %r11), %edx
  2413: 32da                     xorb %dl, %bl
  2415: 498b4028                 movq 0x28(%r8), %rax
  2419: 420fb61418               movzbl (%rax, %r11), %edx
  2424: 32da                     xorb %dl, %bl
  2426: 498b4030                 movq 0x30(%r8), %rax
  2430: 420fb61418               movzbl (%rax, %r11), %edx
  2435: 32da                     xorb %dl, %bl
  2437: 498b4038                 movq 0x38(%r8), %rax
  2441: 420fb61418               movzbl (%rax, %r11), %edx
  2446: 32da                     xorb %dl, %bl
  2448: 43881c1a                 movb %bl, (%r10, %r11)
  2452: 33db                     xorl %ebx, %ebx
  2454: 49ffc3                   incq %r11
  2457: 4c3bd9                   cmpq %rcx, %r11
  2460: 0f8297ffffff             jb 0x939
  2466: 4d8b5108                 movq 8(%r9), %r10
  2470: 4c8b1c24                 movq (%rsp), %r11
  2474: 33db                     xorl %ebx, %ebx
  2476: 498b00                   movq (%r8), %rax
  2479: 420fb61418               movzbl (%rax, %r11), %edx
  2484: 4885d2                   testq %rdx, %rdx
  2487: 0f8422000000             je 0x9df
  2493: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2497: 81c0e6000000             addl $0xe6, %eax
  2503: 81f8ff000000             cmpl $0xff, %eax
  2509: 0f8206000000             jb 0x9d9
  2515: 81c001ffffff             addl $0xffffff01, %eax
  2521: 0fb60406                 movzbl (%rsi, %rax), %eax
  2525: 32d8                     xorb %al, %bl
  2527: 498b4008                 movq 8(%r8), %rax
  2531: 420fb61418               movzbl (%rax, %r11), %edx
  2536: 4885d2                   testq %rdx, %rdx
  2539: 0f8422000000             je 0xa13
  2545: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2549: 81c0cd000000             addl $0xcd, %eax
  2555: 81f8ff000000             cmpl $0xff, %eax
  2561: 0f8206000000             jb 0xa0d
  2567: 81c001ffffff             addl $0xffffff01, %eax
  2573: 0fb60406                 movzbl (%rsi, %rax), %eax
  2577: 32d8                     xorb %al, %bl
  2579: 498b4010                 movq 0x10(%r8), %rax
  2583: 420fb61418               movzbl (%rax, %r11), %edx
  2588: 4885d2                   testq %rdx, %rdx
  2591: 0f8422000000             je 0xa47
  2597: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2601: 81c0e7000000             addl $0xe7, %eax
  2607: 81f8ff000000             cmpl $0xff, %eax
  2613: 0f8206000000             jb 0xa41
  2619: 81c001ffffff             addl $0xffffff01, %eax
  2625: 0fb60406                 movzbl (%rsi, %rax), %eax
  2629: 32d8                     xorb %al, %bl
  2631: 498b4018                 movq 0x18(%r8), %rax
  2635: 420fb61418               movzbl (%rax, %r11), %edx
  2640: 4885d2                   testq %rdx, %rdx
  2643: 0f8422000000             je 0xa7b
  2649: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2653: 81c0ce000000             addl $0xce, %eax
  2659: 81f8ff000000             cmpl $0xff, %eax
  2665: 0f8206000000             jb 0xa75
  2671: 81c001ffffff             addl $0xffffff01, %eax
  2677: 0fb60406                 movzbl (%rsi, %rax), %eax
  2681: 32d8                     xorb %al, %bl
  2683: 498b4020                 movq 0x20(%r8), %rax
  2687: 420fb61418               movzbl (%rax, %r11), %edx
  2692: 4885d2                   testq %rdx, %rdx
  2695: 0f841f000000             je 0xaac
  2701: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2705: 83c04b                   addl $0x4b, %eax
  2708: 81f8ff000000             cmpl $0xff, %eax
  2714: 0f8206000000             jb 0xaa6
  2720: 81c001ffffff             addl $0xffffff01, %eax
  2726: 0fb60406                 movzbl (%rsi, %rax), %eax
  2730: 32d8                     xorb %al, %bl
  2732: 498b4028                 movq 0x28(%r8), %rax
  2736: 420fb61418               movzbl (%rax, %r11), %edx
  2741: 4885d2                   testq %rdx, %rdx
  2744: 0f841f000000             je 0xadd
  2750: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2754: 83c032                   addl $0x32, %eax
  2757: 81f8ff000000             cmpl $0xff, %eax
  2763: 0f8206000000             jb 0xad7
  2769: 81c001ffffff             addl $0xffffff01, %eax
  2775: 0fb60406                 movzbl (%rsi, %rax), %eax
  2779: 32d8                     xorb %al, %bl
  2781: 498b4030                 movq 0x30(%r8), %rax
  2785: 420fb61418               movzbl (%rax, %r11), %edx
  2790: 4885d2                   testq %rdx, %rdx
  2793: 0f841f000000             je 0xb0e
  2799: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2803: 83c019                   addl $0x19, %eax
  2806: 81f8ff000000             cmpl $0xff, %eax
  2812: 0f8206000000             jb 0xb08
  2818: 81c001ffffff             addl $0xffffff01, %eax
  2824: 0fb60406                 movzbl (%rsi, %rax), %eax
  2828: 32d8                     xorb %al, %bl
  2830: 498b4038                 movq 0x38(%r8), %rax
  2834: 420fb61418               movzbl (%rax, %r11), %edx
  2839: 32da                     xorb %dl, %bl
  2841: 43881c1a                 movb %bl, (%r10, %r11)
  2845: 33db                     xorl %ebx, %ebx
  2847: 49ffc3                   incq %r11
  2850: 4c3bd9                   cmpq %rcx, %r11
  2853: 0f8281feffff             jb 0x9ac
  2859: 4883c408                 addq $8, %rsp
  2863: 415e                     popq %r14
  2865: 415d                     popq %r13
  2867: 415c                     popq %r12
  2869: 415f                     popq %r15
  2871: 5b                       popq %rbx
  2872: c5f877                   vzeroupper 
  2875: c3                       retq 
  2876: 0f0b                     ud2 
  2878: 0f0b                     ud2 
