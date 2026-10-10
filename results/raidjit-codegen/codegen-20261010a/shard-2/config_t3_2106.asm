; config_t3_2106 target=avx2 twin=baseline topo=lrc(8,2,2)/rec{0} op=recover U=4
; code_size=2304 twins=[] uops_per_byte=2.2734375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de7080000           leaq 0x8e7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: c4417d6f17               vmovdqa (%r15), %ymm10
    52: 488bc1                   movq %rcx, %rax
    55: 4881e880000000           subq $0x80, %rax
    62: 0f82d3060000             jb 0x717
    68: 666666666666660f1f840000000000 nopw (%rax, %rax)
    83: 66666666660f1f840000000000 nopw (%rax, %rax)
    96: c4a17d6f241e             vmovdqa (%rsi, %r11), %ymm4
   102: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   107: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   112: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   117: c4417d6f4f40             vmovdqa 0x40(%r15), %ymm9
   123: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   128: c4417d6f8f80000000       vmovdqa 0x80(%r15), %ymm9
   137: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   142: c4c145efc0               vpxor %ymm8, %ymm7, %ymm0
   147: c4a17d6f641e20           vmovdqa 0x20(%rsi, %r11), %ymm4
   154: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   159: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   164: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   169: c4417d6f4f40             vmovdqa 0x40(%r15), %ymm9
   175: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   180: c4417d6f8f80000000       vmovdqa 0x80(%r15), %ymm9
   189: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   194: c4c145efc8               vpxor %ymm8, %ymm7, %ymm1
   199: c4a17d6f641e40           vmovdqa 0x40(%rsi, %r11), %ymm4
   206: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   211: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   216: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   221: c4417d6f4f40             vmovdqa 0x40(%r15), %ymm9
   227: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   232: c4417d6f8f80000000       vmovdqa 0x80(%r15), %ymm9
   241: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   246: c4c145efd0               vpxor %ymm8, %ymm7, %ymm2
   251: c4a17d6f641e60           vmovdqa 0x60(%rsi, %r11), %ymm4
   258: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   263: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   268: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   273: c4417d6f4f40             vmovdqa 0x40(%r15), %ymm9
   279: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   284: c4417d6f8f80000000       vmovdqa 0x80(%r15), %ymm9
   293: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   298: c4c145efd8               vpxor %ymm8, %ymm7, %ymm3
   303: c4a17d6f241f             vmovdqa (%rdi, %r11), %ymm4
   309: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   314: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   319: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   324: c4417d6f8fc0000000       vmovdqa 0xc0(%r15), %ymm9
   333: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   338: c4417d6f8f00010000       vmovdqa 0x100(%r15), %ymm9
   347: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   352: c5fdefc7                 vpxor %ymm7, %ymm0, %ymm0
   356: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   361: c4a17d6f641f20           vmovdqa 0x20(%rdi, %r11), %ymm4
   368: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   373: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   378: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   383: c4417d6f8fc0000000       vmovdqa 0xc0(%r15), %ymm9
   392: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   397: c4417d6f8f00010000       vmovdqa 0x100(%r15), %ymm9
   406: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   411: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   415: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   420: c4a17d6f641f40           vmovdqa 0x40(%rdi, %r11), %ymm4
   427: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   432: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   437: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   442: c4417d6f8fc0000000       vmovdqa 0xc0(%r15), %ymm9
   451: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   456: c4417d6f8f00010000       vmovdqa 0x100(%r15), %ymm9
   465: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   470: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   474: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   479: c4a17d6f641f60           vmovdqa 0x60(%rdi, %r11), %ymm4
   486: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   491: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   496: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   501: c4417d6f8fc0000000       vmovdqa 0xc0(%r15), %ymm9
   510: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   515: c4417d6f8f00010000       vmovdqa 0x100(%r15), %ymm9
   524: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   529: c5e5efdf                 vpxor %ymm7, %ymm3, %ymm3
   533: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
   538: c4817d6f241c             vmovdqa (%r12, %r11), %ymm4
   544: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   549: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   554: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   559: c4417d6f8f40010000       vmovdqa 0x140(%r15), %ymm9
   568: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   573: c4417d6f8f80010000       vmovdqa 0x180(%r15), %ymm9
   582: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   587: c5fdefc7                 vpxor %ymm7, %ymm0, %ymm0
   591: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   596: c4817d6f641c20           vmovdqa 0x20(%r12, %r11), %ymm4
   603: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   608: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   613: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   618: c4417d6f8f40010000       vmovdqa 0x140(%r15), %ymm9
   627: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   632: c4417d6f8f80010000       vmovdqa 0x180(%r15), %ymm9
   641: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   646: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   650: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   655: c4817d6f641c40           vmovdqa 0x40(%r12, %r11), %ymm4
   662: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   667: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   672: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   677: c4417d6f8f40010000       vmovdqa 0x140(%r15), %ymm9
   686: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   691: c4417d6f8f80010000       vmovdqa 0x180(%r15), %ymm9
   700: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   705: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   709: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   714: c4817d6f641c60           vmovdqa 0x60(%r12, %r11), %ymm4
   721: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   726: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   731: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   736: c4417d6f8f40010000       vmovdqa 0x140(%r15), %ymm9
   745: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   750: c4417d6f8f80010000       vmovdqa 0x180(%r15), %ymm9
   759: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   764: c5e5efdf                 vpxor %ymm7, %ymm3, %ymm3
   768: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
   773: c4817d6f641d00           vmovdqa (%r13, %r11), %ymm4
   780: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   785: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   790: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   795: c4417d6f8fc0010000       vmovdqa 0x1c0(%r15), %ymm9
   804: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   809: c4417d6f8f00020000       vmovdqa 0x200(%r15), %ymm9
   818: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   823: c5fdefc7                 vpxor %ymm7, %ymm0, %ymm0
   827: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   832: c4817d6f641d20           vmovdqa 0x20(%r13, %r11), %ymm4
   839: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   844: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   849: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   854: c4417d6f8fc0010000       vmovdqa 0x1c0(%r15), %ymm9
   863: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   868: c4417d6f8f00020000       vmovdqa 0x200(%r15), %ymm9
   877: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   882: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   886: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   891: c4817d6f641d40           vmovdqa 0x40(%r13, %r11), %ymm4
   898: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   903: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   908: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   913: c4417d6f8fc0010000       vmovdqa 0x1c0(%r15), %ymm9
   922: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   927: c4417d6f8f00020000       vmovdqa 0x200(%r15), %ymm9
   936: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
   941: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   945: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   950: c4817d6f641d60           vmovdqa 0x60(%r13, %r11), %ymm4
   957: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
   962: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
   967: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
   972: c4417d6f8fc0010000       vmovdqa 0x1c0(%r15), %ymm9
   981: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
   986: c4417d6f8f00020000       vmovdqa 0x200(%r15), %ymm9
   995: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1000: c5e5efdf                 vpxor %ymm7, %ymm3, %ymm3
  1004: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  1009: c4817d6f241e             vmovdqa (%r14, %r11), %ymm4
  1015: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1020: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1025: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1030: c4417d6f8f40020000       vmovdqa 0x240(%r15), %ymm9
  1039: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1044: c4417d6f8f80020000       vmovdqa 0x280(%r15), %ymm9
  1053: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1058: c5fdefc7                 vpxor %ymm7, %ymm0, %ymm0
  1062: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  1067: c4817d6f641e20           vmovdqa 0x20(%r14, %r11), %ymm4
  1074: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1079: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1084: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1089: c4417d6f8f40020000       vmovdqa 0x240(%r15), %ymm9
  1098: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1103: c4417d6f8f80020000       vmovdqa 0x280(%r15), %ymm9
  1112: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1117: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
  1121: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  1126: c4817d6f641e40           vmovdqa 0x40(%r14, %r11), %ymm4
  1133: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1138: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1143: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1148: c4417d6f8f40020000       vmovdqa 0x240(%r15), %ymm9
  1157: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1162: c4417d6f8f80020000       vmovdqa 0x280(%r15), %ymm9
  1171: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1176: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
  1180: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  1185: c4817d6f641e60           vmovdqa 0x60(%r14, %r11), %ymm4
  1192: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1197: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1202: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1207: c4417d6f8f40020000       vmovdqa 0x240(%r15), %ymm9
  1216: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1221: c4417d6f8f80020000       vmovdqa 0x280(%r15), %ymm9
  1230: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1235: c5e5efdf                 vpxor %ymm7, %ymm3, %ymm3
  1239: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  1244: 4d8b5028                 movq 0x28(%r8), %r10
  1248: c4817d6f241a             vmovdqa (%r10, %r11), %ymm4
  1254: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1259: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1264: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1269: c4417d6f8fc0020000       vmovdqa 0x2c0(%r15), %ymm9
  1278: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1283: c4417d6f8f00030000       vmovdqa 0x300(%r15), %ymm9
  1292: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1297: c5fdefc7                 vpxor %ymm7, %ymm0, %ymm0
  1301: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  1306: c4817d6f641a20           vmovdqa 0x20(%r10, %r11), %ymm4
  1313: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1318: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1323: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1328: c4417d6f8fc0020000       vmovdqa 0x2c0(%r15), %ymm9
  1337: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1342: c4417d6f8f00030000       vmovdqa 0x300(%r15), %ymm9
  1351: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1356: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
  1360: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  1365: c4817d6f641a40           vmovdqa 0x40(%r10, %r11), %ymm4
  1372: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1377: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1382: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1387: c4417d6f8fc0020000       vmovdqa 0x2c0(%r15), %ymm9
  1396: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1401: c4417d6f8f00030000       vmovdqa 0x300(%r15), %ymm9
  1410: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1415: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
  1419: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  1424: c4817d6f641a60           vmovdqa 0x60(%r10, %r11), %ymm4
  1431: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1436: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1441: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1446: c4417d6f8fc0020000       vmovdqa 0x2c0(%r15), %ymm9
  1455: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1460: c4417d6f8f00030000       vmovdqa 0x300(%r15), %ymm9
  1469: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1474: c5e5efdf                 vpxor %ymm7, %ymm3, %ymm3
  1478: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  1483: 4d8b5030                 movq 0x30(%r8), %r10
  1487: c4817d6f241a             vmovdqa (%r10, %r11), %ymm4
  1493: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1498: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1503: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1508: c4417d6f8f40030000       vmovdqa 0x340(%r15), %ymm9
  1517: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1522: c4417d6f8f80030000       vmovdqa 0x380(%r15), %ymm9
  1531: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1536: c5fdefc7                 vpxor %ymm7, %ymm0, %ymm0
  1540: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  1545: c4817d6f641a20           vmovdqa 0x20(%r10, %r11), %ymm4
  1552: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1557: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1562: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1567: c4417d6f8f40030000       vmovdqa 0x340(%r15), %ymm9
  1576: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1581: c4417d6f8f80030000       vmovdqa 0x380(%r15), %ymm9
  1590: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1595: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
  1599: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  1604: c4817d6f641a40           vmovdqa 0x40(%r10, %r11), %ymm4
  1611: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1616: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1621: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1626: c4417d6f8f40030000       vmovdqa 0x340(%r15), %ymm9
  1635: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1640: c4417d6f8f80030000       vmovdqa 0x380(%r15), %ymm9
  1649: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1654: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
  1658: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  1663: c4817d6f641a60           vmovdqa 0x60(%r10, %r11), %ymm4
  1670: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1675: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1680: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1685: c4417d6f8f40030000       vmovdqa 0x340(%r15), %ymm9
  1694: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1699: c4417d6f8f80030000       vmovdqa 0x380(%r15), %ymm9
  1708: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1713: c5e5efdf                 vpxor %ymm7, %ymm3, %ymm3
  1717: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  1722: 4d8b5038                 movq 0x38(%r8), %r10
  1726: c4817d6f241a             vmovdqa (%r10, %r11), %ymm4
  1732: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  1736: c4817d6f641a20           vmovdqa 0x20(%r10, %r11), %ymm4
  1743: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  1747: c4817d6f641a40           vmovdqa 0x40(%r10, %r11), %ymm4
  1754: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
  1758: c4817d6f641a60           vmovdqa 0x60(%r10, %r11), %ymm4
  1765: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  1769: 4d8b11                   movq (%r9), %r10
  1772: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
  1778: c4817d7f4c1a20           vmovdqa %ymm1, 0x20(%r10, %r11)
  1785: c4817d7f541a40           vmovdqa %ymm2, 0x40(%r10, %r11)
  1792: c4817d7f5c1a60           vmovdqa %ymm3, 0x60(%r10, %r11)
  1799: 4981c380000000           addq $0x80, %r11
  1806: 4c3bd8                   cmpq %rax, %r11
  1809: 0f8649f9ffff             jbe 0x60
  1815: 4c3bd9                   cmpq %rcx, %r11
  1818: 0f8396010000             jae 0x8b6
  1824: 4153                     pushq %r11
  1826: 498dbfc0030000           leaq 0x3c0(%r15), %rdi
  1833: 498db7c0040000           leaq 0x4c0(%r15), %rsi
  1840: 4d8b11                   movq (%r9), %r10
  1843: 4c8b1c24                 movq (%rsp), %r11
  1847: 33db                     xorl %ebx, %ebx
  1849: 498b00                   movq (%r8), %rax
  1852: 420fb61418               movzbl (%rax, %r11), %edx
  1857: 4885d2                   testq %rdx, %rdx
  1860: 0f841f000000             je 0x769
  1866: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1870: 83c019                   addl $0x19, %eax
  1873: 81f8ff000000             cmpl $0xff, %eax
  1879: 0f8206000000             jb 0x763
  1885: 81c001ffffff             addl $0xffffff01, %eax
  1891: 0fb60406                 movzbl (%rsi, %rax), %eax
  1895: 32d8                     xorb %al, %bl
  1897: 498b4008                 movq 8(%r8), %rax
  1901: 420fb61418               movzbl (%rax, %r11), %edx
  1906: 4885d2                   testq %rdx, %rdx
  1909: 0f841f000000             je 0x79a
  1915: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1919: 83c032                   addl $0x32, %eax
  1922: 81f8ff000000             cmpl $0xff, %eax
  1928: 0f8206000000             jb 0x794
  1934: 81c001ffffff             addl $0xffffff01, %eax
  1940: 0fb60406                 movzbl (%rsi, %rax), %eax
  1944: 32d8                     xorb %al, %bl
  1946: 498b4010                 movq 0x10(%r8), %rax
  1950: 420fb61418               movzbl (%rax, %r11), %edx
  1955: 4885d2                   testq %rdx, %rdx
  1958: 0f841f000000             je 0x7cb
  1964: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1968: 83c04b                   addl $0x4b, %eax
  1971: 81f8ff000000             cmpl $0xff, %eax
  1977: 0f8206000000             jb 0x7c5
  1983: 81c001ffffff             addl $0xffffff01, %eax
  1989: 0fb60406                 movzbl (%rsi, %rax), %eax
  1993: 32d8                     xorb %al, %bl
  1995: 498b4018                 movq 0x18(%r8), %rax
  1999: 420fb61418               movzbl (%rax, %r11), %edx
  2004: 4885d2                   testq %rdx, %rdx
  2007: 0f841f000000             je 0x7fc
  2013: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2017: 83c064                   addl $0x64, %eax
  2020: 81f8ff000000             cmpl $0xff, %eax
  2026: 0f8206000000             jb 0x7f6
  2032: 81c001ffffff             addl $0xffffff01, %eax
  2038: 0fb60406                 movzbl (%rsi, %rax), %eax
  2042: 32d8                     xorb %al, %bl
  2044: 498b4020                 movq 0x20(%r8), %rax
  2048: 420fb61418               movzbl (%rax, %r11), %edx
  2053: 4885d2                   testq %rdx, %rdx
  2056: 0f841f000000             je 0x82d
  2062: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2066: 83c07d                   addl $0x7d, %eax
  2069: 81f8ff000000             cmpl $0xff, %eax
  2075: 0f8206000000             jb 0x827
  2081: 81c001ffffff             addl $0xffffff01, %eax
  2087: 0fb60406                 movzbl (%rsi, %rax), %eax
  2091: 32d8                     xorb %al, %bl
  2093: 498b4028                 movq 0x28(%r8), %rax
  2097: 420fb61418               movzbl (%rax, %r11), %edx
  2102: 4885d2                   testq %rdx, %rdx
  2105: 0f8422000000             je 0x861
  2111: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2115: 81c096000000             addl $0x96, %eax
  2121: 81f8ff000000             cmpl $0xff, %eax
  2127: 0f8206000000             jb 0x85b
  2133: 81c001ffffff             addl $0xffffff01, %eax
  2139: 0fb60406                 movzbl (%rsi, %rax), %eax
  2143: 32d8                     xorb %al, %bl
  2145: 498b4030                 movq 0x30(%r8), %rax
  2149: 420fb61418               movzbl (%rax, %r11), %edx
  2154: 4885d2                   testq %rdx, %rdx
  2157: 0f8422000000             je 0x895
  2163: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2167: 81c0af000000             addl $0xaf, %eax
  2173: 81f8ff000000             cmpl $0xff, %eax
  2179: 0f8206000000             jb 0x88f
  2185: 81c001ffffff             addl $0xffffff01, %eax
  2191: 0fb60406                 movzbl (%rsi, %rax), %eax
  2195: 32d8                     xorb %al, %bl
  2197: 498b4038                 movq 0x38(%r8), %rax
  2201: 420fb61418               movzbl (%rax, %r11), %edx
  2206: 32da                     xorb %dl, %bl
  2208: 43881c1a                 movb %bl, (%r10, %r11)
  2212: 33db                     xorl %ebx, %ebx
  2214: 49ffc3                   incq %r11
  2217: 4c3bd9                   cmpq %rcx, %r11
  2220: 0f8287feffff             jb 0x739
  2226: 4883c408                 addq $8, %rsp
  2230: 415e                     popq %r14
  2232: 415d                     popq %r13
  2234: 415c                     popq %r12
  2236: 415f                     popq %r15
  2238: 5b                       popq %rbx
  2239: c5f877                   vzeroupper 
  2242: c3                       retq 
  2243: 0f0b                     ud2 
  2245: 0f0b                     ud2 
  2247: 0f0b                     ud2 
  2249: 0f0b                     ud2 
  2251: 0f0b                     ud2 
  2253: 0f0b                     ud2 
  2255: 0f0b                     ud2 
  2257: 0f0b                     ud2 
  2259: 0f0b                     ud2 
  2261: 0f0b                     ud2 
  2263: 0f0b                     ud2 
  2265: 0f0b                     ud2 
  2267: 0f0b                     ud2 
  2269: 0f0b                     ud2 
  2271: 0f0b                     ud2 
  2273: 0f0b                     ud2 
  2275: 0f0b                     ud2 
  2277: 0f0b                     ud2 
  2279: 0f0b                     ud2 
  2281: 0f0b                     ud2 
  2283: 0f0b                     ud2 
  2285: 0f0b                     ud2 
  2287: 0f0b                     ud2 
  2289: 0f0b                     ud2 
  2291: 0f0b                     ud2 
  2293: 0f0b                     ud2 
  2295: 0f0b                     ud2 
  2297: 0f0b                     ud2 
  2299: 0f0b                     ud2 
  2301: 0f0b                     ud2 
  2303: cc                       int3 
