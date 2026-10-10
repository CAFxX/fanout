; config_t1_373 target=avx2 twin=xorbit topo=rs(6,3)/enc op=encode U=1
; code_size=3008 twins=["xorbit"] uops_per_byte=14.71875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3da70b0000           leaq 0xba7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: c4c17d6f37               vmovdqa (%r15), %ymm6
    52: 488bc1                   movq %rcx, %rax
    55: 4883e820                 subq $0x20, %rax
    59: 0f8297080000             jb 0x8d8
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: c4a17d6f1c1e             vmovdqa (%rsi, %r11), %ymm3
   102: c5dd71d300               vpsrlw $0, %ymm3, %ymm4
   107: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   111: c4a17d6f1c1f             vmovdqa (%rdi, %r11), %ymm3
   117: c5d571d300               vpsrlw $0, %ymm3, %ymm5
   122: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   126: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   130: c5fd6fc4                 vmovdqa %ymm4, %ymm0
   134: c5fd6fcc                 vmovdqa %ymm4, %ymm1
   138: c5d571f403               vpsllw $3, %ymm4, %ymm5
   143: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   147: c5fd6fd5                 vmovdqa %ymm5, %ymm2
   151: c5d571f407               vpsllw $7, %ymm4, %ymm5
   156: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   160: c5dd71d305               vpsrlw $5, %ymm3, %ymm4
   165: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   169: c4817d6f1c1c             vmovdqa (%r12, %r11), %ymm3
   175: c5d571d307               vpsrlw $7, %ymm3, %ymm5
   180: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   184: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   188: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
   192: c5d571f401               vpsllw $1, %ymm4, %ymm5
   197: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   201: c5d571f403               vpsllw $3, %ymm4, %ymm5
   206: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   210: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   214: c5d571f404               vpsllw $4, %ymm4, %ymm5
   219: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   223: c4a17d6f1c1f             vmovdqa (%rdi, %r11), %ymm3
   229: c5dd71d306               vpsrlw $6, %ymm3, %ymm4
   234: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   238: c4817d6f1c1e             vmovdqa (%r14, %r11), %ymm3
   244: c5d571d307               vpsrlw $7, %ymm3, %ymm5
   249: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   253: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   257: c5d571f401               vpsllw $1, %ymm4, %ymm5
   262: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   266: c5d571f402               vpsllw $2, %ymm4, %ymm5
   271: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   275: c5d571f404               vpsllw $4, %ymm4, %ymm5
   280: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   284: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   288: c5d571f405               vpsllw $5, %ymm4, %ymm5
   293: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   297: c4a17d6f1c1e             vmovdqa (%rsi, %r11), %ymm3
   303: c5dd71d301               vpsrlw $1, %ymm3, %ymm4
   308: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   312: c4817d6f5c1d00           vmovdqa (%r13, %r11), %ymm3
   319: c5d571d300               vpsrlw $0, %ymm3, %ymm5
   324: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   328: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   332: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
   336: c5d571f402               vpsllw $2, %ymm4, %ymm5
   341: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   345: c5d571f403               vpsllw $3, %ymm4, %ymm5
   350: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   354: c5d571f407               vpsllw $7, %ymm4, %ymm5
   359: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   363: c4817d6f1c1c             vmovdqa (%r12, %r11), %ymm3
   369: c5dd71d305               vpsrlw $5, %ymm3, %ymm4
   374: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   378: c4817d6f5c1d00           vmovdqa (%r13, %r11), %ymm3
   385: c5d571d306               vpsrlw $6, %ymm3, %ymm5
   390: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   394: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   398: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   402: c5d571f401               vpsllw $1, %ymm4, %ymm5
   407: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   411: c5d571f403               vpsllw $3, %ymm4, %ymm5
   416: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   420: c5d571f404               vpsllw $4, %ymm4, %ymm5
   425: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   429: c4817d6f1c1c             vmovdqa (%r12, %r11), %ymm3
   435: c5dd71d306               vpsrlw $6, %ymm3, %ymm4
   440: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   444: c4817d6f5c1d00           vmovdqa (%r13, %r11), %ymm3
   451: c5d571d307               vpsrlw $7, %ymm3, %ymm5
   456: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   460: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   464: c5d571f401               vpsllw $1, %ymm4, %ymm5
   469: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   473: c5d571f402               vpsllw $2, %ymm4, %ymm5
   478: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   482: c5d571f404               vpsllw $4, %ymm4, %ymm5
   487: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   491: c5d571f405               vpsllw $5, %ymm4, %ymm5
   496: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   500: c4a17d6f1c1e             vmovdqa (%rsi, %r11), %ymm3
   506: c5dd71d300               vpsrlw $0, %ymm3, %ymm4
   511: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   515: c5d571d303               vpsrlw $3, %ymm3, %ymm5
   520: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   524: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   528: c5d571f401               vpsllw $1, %ymm4, %ymm5
   533: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   537: c5d571f402               vpsllw $2, %ymm4, %ymm5
   542: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   546: c5dd71d301               vpsrlw $1, %ymm3, %ymm4
   551: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   555: c4a17d6f1c1f             vmovdqa (%rdi, %r11), %ymm3
   561: c5d571d301               vpsrlw $1, %ymm3, %ymm5
   566: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   570: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   574: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   578: c5d571f401               vpsllw $1, %ymm4, %ymm5
   583: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   587: c4a17d6f1c1e             vmovdqa (%rsi, %r11), %ymm3
   593: c5dd71d306               vpsrlw $6, %ymm3, %ymm4
   598: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   602: c4817d6f1c1e             vmovdqa (%r14, %r11), %ymm3
   608: c5d571d306               vpsrlw $6, %ymm3, %ymm5
   613: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   617: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   621: c5d571f404               vpsllw $4, %ymm4, %ymm5
   626: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   630: c5d571f406               vpsllw $6, %ymm4, %ymm5
   635: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   639: c4a17d6f1c1e             vmovdqa (%rsi, %r11), %ymm3
   645: c5dd71d307               vpsrlw $7, %ymm3, %ymm4
   650: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   654: c4a17d6f1c1f             vmovdqa (%rdi, %r11), %ymm3
   660: c5d571d307               vpsrlw $7, %ymm3, %ymm5
   665: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   669: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   673: c5d571f405               vpsllw $5, %ymm4, %ymm5
   678: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   682: c5d571f407               vpsllw $7, %ymm4, %ymm5
   687: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   691: c5dd71d301               vpsrlw $1, %ymm3, %ymm4
   696: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   700: c4817d6f5c1d00           vmovdqa (%r13, %r11), %ymm3
   707: c5d571d305               vpsrlw $5, %ymm3, %ymm5
   712: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   716: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   720: c5d571f404               vpsllw $4, %ymm4, %ymm5
   725: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   729: c5d571f407               vpsllw $7, %ymm4, %ymm5
   734: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   738: c4a17d6f1c1f             vmovdqa (%rdi, %r11), %ymm3
   744: c5dd71d307               vpsrlw $7, %ymm3, %ymm4
   749: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   753: c4817d6f1c1c             vmovdqa (%r12, %r11), %ymm3
   759: c5d571d302               vpsrlw $2, %ymm3, %ymm5
   764: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   768: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   772: c5d571f403               vpsllw $3, %ymm4, %ymm5
   777: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   781: c5d571f405               vpsllw $5, %ymm4, %ymm5
   786: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   790: c4a17d6f1c1e             vmovdqa (%rsi, %r11), %ymm3
   796: c5dd71d300               vpsrlw $0, %ymm3, %ymm4
   801: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   805: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
   809: c5d571f406               vpsllw $6, %ymm4, %ymm5
   814: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   818: c5d571f407               vpsllw $7, %ymm4, %ymm5
   823: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   827: c5dd71d302               vpsrlw $2, %ymm3, %ymm4
   832: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   836: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
   840: c5d571f401               vpsllw $1, %ymm4, %ymm5
   845: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   849: c5d571f402               vpsllw $2, %ymm4, %ymm5
   854: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   858: c5dd71d303               vpsrlw $3, %ymm3, %ymm4
   863: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   867: c5d571f403               vpsllw $3, %ymm4, %ymm5
   872: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   876: c5dd71d304               vpsrlw $4, %ymm3, %ymm4
   881: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   885: c5d571f402               vpsllw $2, %ymm4, %ymm5
   890: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   894: c5d571f403               vpsllw $3, %ymm4, %ymm5
   899: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   903: c5d571f404               vpsllw $4, %ymm4, %ymm5
   908: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   912: c5dd71d305               vpsrlw $5, %ymm3, %ymm4
   917: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   921: c5d571f403               vpsllw $3, %ymm4, %ymm5
   926: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   930: c5d571f404               vpsllw $4, %ymm4, %ymm5
   935: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   939: c5d571f405               vpsllw $5, %ymm4, %ymm5
   944: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   948: c5dd71d306               vpsrlw $6, %ymm3, %ymm4
   953: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   957: c5d571f405               vpsllw $5, %ymm4, %ymm5
   962: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   966: c5dd71d307               vpsrlw $7, %ymm3, %ymm4
   971: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   975: c5d571f406               vpsllw $6, %ymm4, %ymm5
   980: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   984: c4a17d6f1c1f             vmovdqa (%rdi, %r11), %ymm3
   990: c5dd71d300               vpsrlw $0, %ymm3, %ymm4
   995: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   999: c5d571f401               vpsllw $1, %ymm4, %ymm5
  1004: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1008: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1013: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1017: c5dd71d301               vpsrlw $1, %ymm3, %ymm4
  1022: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1026: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1031: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1035: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1040: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1044: c5dd71d302               vpsrlw $2, %ymm3, %ymm4
  1049: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1053: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1057: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1062: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1066: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1071: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1075: c5dd71d303               vpsrlw $3, %ymm3, %ymm4
  1080: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1084: c5d571f401               vpsllw $1, %ymm4, %ymm5
  1089: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1093: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1098: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1102: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1107: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1111: c5dd71d304               vpsrlw $4, %ymm3, %ymm4
  1116: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1120: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1125: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1129: c5d571f404               vpsllw $4, %ymm4, %ymm5
  1134: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1138: c5d571f407               vpsllw $7, %ymm4, %ymm5
  1143: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1147: c5dd71d305               vpsrlw $5, %ymm3, %ymm4
  1152: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1156: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1161: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1165: c5dd71d306               vpsrlw $6, %ymm3, %ymm4
  1170: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1174: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1179: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1183: c5dd71d307               vpsrlw $7, %ymm3, %ymm4
  1188: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1192: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1197: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1201: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1206: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1210: c4817d6f1c1c             vmovdqa (%r12, %r11), %ymm3
  1216: c5dd71d300               vpsrlw $0, %ymm3, %ymm4
  1221: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1225: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1229: c5d571f401               vpsllw $1, %ymm4, %ymm5
  1234: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1238: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1243: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1247: c5dd71d301               vpsrlw $1, %ymm3, %ymm4
  1252: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1256: c5d571f401               vpsllw $1, %ymm4, %ymm5
  1261: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1265: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1270: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1274: c5d571f404               vpsllw $4, %ymm4, %ymm5
  1279: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1283: c5dd71d302               vpsrlw $2, %ymm3, %ymm4
  1288: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1292: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1297: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1301: c5dd71d303               vpsrlw $3, %ymm3, %ymm4
  1306: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1310: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1315: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1319: c5d571f404               vpsllw $4, %ymm4, %ymm5
  1324: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1328: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1333: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1337: c5dd71d304               vpsrlw $4, %ymm3, %ymm4
  1342: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1346: c5d571f404               vpsllw $4, %ymm4, %ymm5
  1351: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1355: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1360: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1364: c5d571f407               vpsllw $7, %ymm4, %ymm5
  1369: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1373: c5dd71d305               vpsrlw $5, %ymm3, %ymm4
  1378: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1382: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1387: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1391: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1396: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1400: c5dd71d306               vpsrlw $6, %ymm3, %ymm4
  1405: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1409: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1414: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1418: c5d571f407               vpsllw $7, %ymm4, %ymm5
  1423: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1427: c5dd71d307               vpsrlw $7, %ymm3, %ymm4
  1432: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1436: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1441: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1445: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1450: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1454: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1459: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1463: c5d571f407               vpsllw $7, %ymm4, %ymm5
  1468: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1472: c4817d6f5c1d00           vmovdqa (%r13, %r11), %ymm3
  1479: c5dd71d300               vpsrlw $0, %ymm3, %ymm4
  1484: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1488: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1492: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1497: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1501: c5dd71d301               vpsrlw $1, %ymm3, %ymm4
  1506: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1510: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
  1514: c5d571f401               vpsllw $1, %ymm4, %ymm5
  1519: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1523: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1528: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1532: c5dd71d302               vpsrlw $2, %ymm3, %ymm4
  1537: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1541: c5d571f401               vpsllw $1, %ymm4, %ymm5
  1546: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1550: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1555: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1559: c5d571f404               vpsllw $4, %ymm4, %ymm5
  1564: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1568: c5dd71d303               vpsrlw $3, %ymm3, %ymm4
  1573: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1577: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1582: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1586: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1591: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1595: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1600: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1604: c5dd71d304               vpsrlw $4, %ymm3, %ymm4
  1609: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1613: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1618: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1622: c5d571f404               vpsllw $4, %ymm4, %ymm5
  1627: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1631: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1636: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1640: c5dd71d305               vpsrlw $5, %ymm3, %ymm4
  1645: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1649: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1654: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1658: c5dd71d306               vpsrlw $6, %ymm3, %ymm4
  1663: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1667: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1672: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1676: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1681: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1685: c5dd71d307               vpsrlw $7, %ymm3, %ymm4
  1690: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1694: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1699: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1703: c5d571f407               vpsllw $7, %ymm4, %ymm5
  1708: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1712: c4817d6f1c1e             vmovdqa (%r14, %r11), %ymm3
  1718: c5dd71d300               vpsrlw $0, %ymm3, %ymm4
  1723: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1727: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1731: c5d571f401               vpsllw $1, %ymm4, %ymm5
  1736: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1740: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1745: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1749: c5dd71d301               vpsrlw $1, %ymm3, %ymm4
  1754: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1758: c5d571f401               vpsllw $1, %ymm4, %ymm5
  1763: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1767: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1772: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1776: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1781: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1785: c5dd71d302               vpsrlw $2, %ymm3, %ymm4
  1790: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1794: c5d571f402               vpsllw $2, %ymm4, %ymm5
  1799: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1803: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1808: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1812: c5d571f404               vpsllw $4, %ymm4, %ymm5
  1817: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1821: c5dd71d303               vpsrlw $3, %ymm3, %ymm4
  1826: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1830: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1835: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1839: c5d571f404               vpsllw $4, %ymm4, %ymm5
  1844: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1848: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1853: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1857: c5dd71d304               vpsrlw $4, %ymm3, %ymm4
  1862: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1866: c5d571f404               vpsllw $4, %ymm4, %ymm5
  1871: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1875: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1880: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1884: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1889: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1893: c5dd71d305               vpsrlw $5, %ymm3, %ymm4
  1898: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1902: c5d571f405               vpsllw $5, %ymm4, %ymm5
  1907: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  1911: c5d571f406               vpsllw $6, %ymm4, %ymm5
  1916: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1920: c5d571f407               vpsllw $7, %ymm4, %ymm5
  1925: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1929: c5dd71d306               vpsrlw $6, %ymm3, %ymm4
  1934: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1938: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
  1942: c5d571f401               vpsllw $1, %ymm4, %ymm5
  1947: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1951: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1956: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  1960: c5d571f407               vpsllw $7, %ymm4, %ymm5
  1965: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1969: c5dd71d307               vpsrlw $7, %ymm3, %ymm4
  1974: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  1978: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1982: c5d571f401               vpsllw $1, %ymm4, %ymm5
  1987: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  1991: c5d571f403               vpsllw $3, %ymm4, %ymm5
  1996: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  2000: c5d571f407               vpsllw $7, %ymm4, %ymm5
  2005: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  2009: 4d8b5028                 movq 0x28(%r8), %r10
  2013: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
  2019: c5dd71d300               vpsrlw $0, %ymm3, %ymm4
  2024: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  2028: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2032: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2036: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
  2040: c5dd71d301               vpsrlw $1, %ymm3, %ymm4
  2045: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  2049: c5d571f401               vpsllw $1, %ymm4, %ymm5
  2054: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  2058: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  2062: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  2066: c5dd71d302               vpsrlw $2, %ymm3, %ymm4
  2071: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  2075: c5d571f402               vpsllw $2, %ymm4, %ymm5
  2080: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  2084: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  2088: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  2092: c5dd71d303               vpsrlw $3, %ymm3, %ymm4
  2097: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  2101: c5d571f403               vpsllw $3, %ymm4, %ymm5
  2106: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  2110: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  2114: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  2118: c5dd71d304               vpsrlw $4, %ymm3, %ymm4
  2123: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  2127: c5d571f404               vpsllw $4, %ymm4, %ymm5
  2132: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  2136: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  2140: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  2144: c5dd71d305               vpsrlw $5, %ymm3, %ymm4
  2149: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  2153: c5d571f405               vpsllw $5, %ymm4, %ymm5
  2158: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  2162: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  2166: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  2170: c5dd71d306               vpsrlw $6, %ymm3, %ymm4
  2175: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  2179: c5d571f406               vpsllw $6, %ymm4, %ymm5
  2184: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  2188: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  2192: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  2196: c5dd71d307               vpsrlw $7, %ymm3, %ymm4
  2201: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
  2205: c5d571f407               vpsllw $7, %ymm4, %ymm5
  2210: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
  2214: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
  2218: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
  2222: 4d8b11                   movq (%r9), %r10
  2225: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
  2231: 4d8b5108                 movq 8(%r9), %r10
  2235: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
  2241: 4d8b5110                 movq 0x10(%r9), %r10
  2245: c4817d7f141a             vmovdqa %ymm2, (%r10, %r11)
  2251: 4983c320                 addq $0x20, %r11
  2255: 4c3bd8                   cmpq %rax, %r11
  2258: 0f8688f7ffff             jbe 0x60
  2264: 4c3bd9                   cmpq %rcx, %r11
  2267: 0f83af020000             jae 0xb90
  2273: 4153                     pushq %r11
  2275: 498d7f40                 leaq 0x40(%r15), %rdi
  2279: 498db740010000           leaq 0x140(%r15), %rsi
  2286: 4d8b11                   movq (%r9), %r10
  2289: 4c8b1c24                 movq (%rsp), %r11
  2293: 33db                     xorl %ebx, %ebx
  2295: 498b00                   movq (%r8), %rax
  2298: 420fb61418               movzbl (%rax, %r11), %edx
  2303: 32da                     xorb %dl, %bl
  2305: 498b4008                 movq 8(%r8), %rax
  2309: 420fb61418               movzbl (%rax, %r11), %edx
  2314: 32da                     xorb %dl, %bl
  2316: 498b4010                 movq 0x10(%r8), %rax
  2320: 420fb61418               movzbl (%rax, %r11), %edx
  2325: 32da                     xorb %dl, %bl
  2327: 498b4018                 movq 0x18(%r8), %rax
  2331: 420fb61418               movzbl (%rax, %r11), %edx
  2336: 32da                     xorb %dl, %bl
  2338: 498b4020                 movq 0x20(%r8), %rax
  2342: 420fb61418               movzbl (%rax, %r11), %edx
  2347: 32da                     xorb %dl, %bl
  2349: 498b4028                 movq 0x28(%r8), %rax
  2353: 420fb61418               movzbl (%rax, %r11), %edx
  2358: 32da                     xorb %dl, %bl
  2360: 43881c1a                 movb %bl, (%r10, %r11)
  2364: 33db                     xorl %ebx, %ebx
  2366: 49ffc3                   incq %r11
  2369: 4c3bd9                   cmpq %rcx, %r11
  2372: 0f82adffffff             jb 0x8f7
  2378: 4d8b5108                 movq 8(%r9), %r10
  2382: 4c8b1c24                 movq (%rsp), %r11
  2386: 33db                     xorl %ebx, %ebx
  2388: 498b00                   movq (%r8), %rax
  2391: 420fb61418               movzbl (%rax, %r11), %edx
  2396: 4885d2                   testq %rdx, %rdx
  2399: 0f8422000000             je 0x987
  2405: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2409: 81c0e6000000             addl $0xe6, %eax
  2415: 81f8ff000000             cmpl $0xff, %eax
  2421: 0f8206000000             jb 0x981
  2427: 81c001ffffff             addl $0xffffff01, %eax
  2433: 0fb60406                 movzbl (%rsi, %rax), %eax
  2437: 32d8                     xorb %al, %bl
  2439: 498b4008                 movq 8(%r8), %rax
  2443: 420fb61418               movzbl (%rax, %r11), %edx
  2448: 4885d2                   testq %rdx, %rdx
  2451: 0f8422000000             je 0x9bb
  2457: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2461: 81c0cd000000             addl $0xcd, %eax
  2467: 81f8ff000000             cmpl $0xff, %eax
  2473: 0f8206000000             jb 0x9b5
  2479: 81c001ffffff             addl $0xffffff01, %eax
  2485: 0fb60406                 movzbl (%rsi, %rax), %eax
  2489: 32d8                     xorb %al, %bl
  2491: 498b4010                 movq 0x10(%r8), %rax
  2495: 420fb61418               movzbl (%rax, %r11), %edx
  2500: 4885d2                   testq %rdx, %rdx
  2503: 0f841f000000             je 0x9ec
  2509: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2513: 83c04b                   addl $0x4b, %eax
  2516: 81f8ff000000             cmpl $0xff, %eax
  2522: 0f8206000000             jb 0x9e6
  2528: 81c001ffffff             addl $0xffffff01, %eax
  2534: 0fb60406                 movzbl (%rsi, %rax), %eax
  2538: 32d8                     xorb %al, %bl
  2540: 498b4018                 movq 0x18(%r8), %rax
  2544: 420fb61418               movzbl (%rax, %r11), %edx
  2549: 4885d2                   testq %rdx, %rdx
  2552: 0f841f000000             je 0xa1d
  2558: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2562: 83c032                   addl $0x32, %eax
  2565: 81f8ff000000             cmpl $0xff, %eax
  2571: 0f8206000000             jb 0xa17
  2577: 81c001ffffff             addl $0xffffff01, %eax
  2583: 0fb60406                 movzbl (%rsi, %rax), %eax
  2587: 32d8                     xorb %al, %bl
  2589: 498b4020                 movq 0x20(%r8), %rax
  2593: 420fb61418               movzbl (%rax, %r11), %edx
  2598: 4885d2                   testq %rdx, %rdx
  2601: 0f841f000000             je 0xa4e
  2607: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2611: 83c019                   addl $0x19, %eax
  2614: 81f8ff000000             cmpl $0xff, %eax
  2620: 0f8206000000             jb 0xa48
  2626: 81c001ffffff             addl $0xffffff01, %eax
  2632: 0fb60406                 movzbl (%rsi, %rax), %eax
  2636: 32d8                     xorb %al, %bl
  2638: 498b4028                 movq 0x28(%r8), %rax
  2642: 420fb61418               movzbl (%rax, %r11), %edx
  2647: 32da                     xorb %dl, %bl
  2649: 43881c1a                 movb %bl, (%r10, %r11)
  2653: 33db                     xorl %ebx, %ebx
  2655: 49ffc3                   incq %r11
  2658: 4c3bd9                   cmpq %rcx, %r11
  2661: 0f82e9feffff             jb 0x954
  2667: 4d8b5110                 movq 0x10(%r9), %r10
  2671: 4c8b1c24                 movq (%rsp), %r11
  2675: 33db                     xorl %ebx, %ebx
  2677: 498b00                   movq (%r8), %rax
  2680: 420fb61418               movzbl (%rax, %r11), %edx
  2685: 4885d2                   testq %rdx, %rdx
  2688: 0f8422000000             je 0xaa8
  2694: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2698: 81c0cd000000             addl $0xcd, %eax
  2704: 81f8ff000000             cmpl $0xff, %eax
  2710: 0f8206000000             jb 0xaa2
  2716: 81c001ffffff             addl $0xffffff01, %eax
  2722: 0fb60406                 movzbl (%rsi, %rax), %eax
  2726: 32d8                     xorb %al, %bl
  2728: 498b4008                 movq 8(%r8), %rax
  2732: 420fb61418               movzbl (%rax, %r11), %edx
  2737: 4885d2                   testq %rdx, %rdx
  2740: 0f841f000000             je 0xad9
  2746: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2750: 83c04b                   addl $0x4b, %eax
  2753: 81f8ff000000             cmpl $0xff, %eax
  2759: 0f8206000000             jb 0xad3
  2765: 81c001ffffff             addl $0xffffff01, %eax
  2771: 0fb60406                 movzbl (%rsi, %rax), %eax
  2775: 32d8                     xorb %al, %bl
  2777: 498b4010                 movq 0x10(%r8), %rax
  2781: 420fb61418               movzbl (%rax, %r11), %edx
  2786: 4885d2                   testq %rdx, %rdx
  2789: 0f841f000000             je 0xb0a
  2795: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2799: 83c019                   addl $0x19, %eax
  2802: 81f8ff000000             cmpl $0xff, %eax
  2808: 0f8206000000             jb 0xb04
  2814: 81c001ffffff             addl $0xffffff01, %eax
  2820: 0fb60406                 movzbl (%rsi, %rax), %eax
  2824: 32d8                     xorb %al, %bl
  2826: 498b4018                 movq 0x18(%r8), %rax
  2830: 420fb61418               movzbl (%rax, %r11), %edx
  2835: 4885d2                   testq %rdx, %rdx
  2838: 0f8422000000             je 0xb3e
  2844: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2848: 81c0e6000000             addl $0xe6, %eax
  2854: 81f8ff000000             cmpl $0xff, %eax
  2860: 0f8206000000             jb 0xb38
  2866: 81c001ffffff             addl $0xffffff01, %eax
  2872: 0fb60406                 movzbl (%rsi, %rax), %eax
  2876: 32d8                     xorb %al, %bl
  2878: 498b4020                 movq 0x20(%r8), %rax
  2882: 420fb61418               movzbl (%rax, %r11), %edx
  2887: 4885d2                   testq %rdx, %rdx
  2890: 0f841f000000             je 0xb6f
  2896: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2900: 83c032                   addl $0x32, %eax
  2903: 81f8ff000000             cmpl $0xff, %eax
  2909: 0f8206000000             jb 0xb69
  2915: 81c001ffffff             addl $0xffffff01, %eax
  2921: 0fb60406                 movzbl (%rsi, %rax), %eax
  2925: 32d8                     xorb %al, %bl
  2927: 498b4028                 movq 0x28(%r8), %rax
  2931: 420fb61418               movzbl (%rax, %r11), %edx
  2936: 32da                     xorb %dl, %bl
  2938: 43881c1a                 movb %bl, (%r10, %r11)
  2942: 33db                     xorl %ebx, %ebx
  2944: 49ffc3                   incq %r11
  2947: 4c3bd9                   cmpq %rcx, %r11
  2950: 0f82e9feffff             jb 0xa75
  2956: 4883c408                 addq $8, %rsp
  2960: 415e                     popq %r14
  2962: 415d                     popq %r13
  2964: 415c                     popq %r12
  2966: 415f                     popq %r15
  2968: 5b                       popq %rbx
  2969: c5f877                   vzeroupper 
  2972: c3                       retq 
  2973: 0f0b                     ud2 
  2975: 0f0b                     ud2 
  2977: 0f0b                     ud2 
  2979: 0f0b                     ud2 
  2981: 0f0b                     ud2 
  2983: 0f0b                     ud2 
  2985: 0f0b                     ud2 
  2987: 0f0b                     ud2 
  2989: 0f0b                     ud2 
  2991: 0f0b                     ud2 
  2993: 0f0b                     ud2 
  2995: 0f0b                     ud2 
  2997: 0f0b                     ud2 
  2999: 0f0b                     ud2 
  3001: 0f0b                     ud2 
  3003: 0f0b                     ud2 
  3005: 0f0b                     ud2 
  3007: cc                       int3 
