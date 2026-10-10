; config_t3_2104 target=avx2 twin=baseline topo=lrc_azure()/rec{0} op=recover U=4
; code_size=3456 twins=[] uops_per_byte=3.5546875
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
    47: c4417d6f17               vmovdqa (%r15), %ymm10
    52: 488bc1                   movq %rcx, %rax
    55: 4881e880000000           subq $0x80, %rax
    62: 0f828f0a0000             jb 0xad3
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
  1732: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1737: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1742: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1747: c4417d6f8fc0030000       vmovdqa 0x3c0(%r15), %ymm9
  1756: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1761: c4417d6f8f00040000       vmovdqa 0x400(%r15), %ymm9
  1770: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1775: c5fdefc7                 vpxor %ymm7, %ymm0, %ymm0
  1779: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  1784: c4817d6f641a20           vmovdqa 0x20(%r10, %r11), %ymm4
  1791: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1796: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1801: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1806: c4417d6f8fc0030000       vmovdqa 0x3c0(%r15), %ymm9
  1815: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1820: c4417d6f8f00040000       vmovdqa 0x400(%r15), %ymm9
  1829: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1834: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
  1838: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  1843: c4817d6f641a40           vmovdqa 0x40(%r10, %r11), %ymm4
  1850: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1855: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1860: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1865: c4417d6f8fc0030000       vmovdqa 0x3c0(%r15), %ymm9
  1874: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1879: c4417d6f8f00040000       vmovdqa 0x400(%r15), %ymm9
  1888: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1893: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
  1897: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  1902: c4817d6f641a60           vmovdqa 0x60(%r10, %r11), %ymm4
  1909: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1914: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1919: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1924: c4417d6f8fc0030000       vmovdqa 0x3c0(%r15), %ymm9
  1933: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  1938: c4417d6f8f00040000       vmovdqa 0x400(%r15), %ymm9
  1947: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  1952: c5e5efdf                 vpxor %ymm7, %ymm3, %ymm3
  1956: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  1961: 4d8b5040                 movq 0x40(%r8), %r10
  1965: c4817d6f241a             vmovdqa (%r10, %r11), %ymm4
  1971: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  1976: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  1981: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  1986: c4417d6f8f40040000       vmovdqa 0x440(%r15), %ymm9
  1995: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2000: c4417d6f8f80040000       vmovdqa 0x480(%r15), %ymm9
  2009: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2014: c5fdefc7                 vpxor %ymm7, %ymm0, %ymm0
  2018: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  2023: c4817d6f641a20           vmovdqa 0x20(%r10, %r11), %ymm4
  2030: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2035: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2040: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2045: c4417d6f8f40040000       vmovdqa 0x440(%r15), %ymm9
  2054: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2059: c4417d6f8f80040000       vmovdqa 0x480(%r15), %ymm9
  2068: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2073: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
  2077: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  2082: c4817d6f641a40           vmovdqa 0x40(%r10, %r11), %ymm4
  2089: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2094: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2099: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2104: c4417d6f8f40040000       vmovdqa 0x440(%r15), %ymm9
  2113: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2118: c4417d6f8f80040000       vmovdqa 0x480(%r15), %ymm9
  2127: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2132: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
  2136: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  2141: c4817d6f641a60           vmovdqa 0x60(%r10, %r11), %ymm4
  2148: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2153: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2158: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2163: c4417d6f8f40040000       vmovdqa 0x440(%r15), %ymm9
  2172: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2177: c4417d6f8f80040000       vmovdqa 0x480(%r15), %ymm9
  2186: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2191: c5e5efdf                 vpxor %ymm7, %ymm3, %ymm3
  2195: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  2200: 4d8b5048                 movq 0x48(%r8), %r10
  2204: c4817d6f241a             vmovdqa (%r10, %r11), %ymm4
  2210: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2215: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2220: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2225: c4417d6f8fc0040000       vmovdqa 0x4c0(%r15), %ymm9
  2234: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2239: c4417d6f8f00050000       vmovdqa 0x500(%r15), %ymm9
  2248: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2253: c5fdefc7                 vpxor %ymm7, %ymm0, %ymm0
  2257: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  2262: c4817d6f641a20           vmovdqa 0x20(%r10, %r11), %ymm4
  2269: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2274: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2279: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2284: c4417d6f8fc0040000       vmovdqa 0x4c0(%r15), %ymm9
  2293: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2298: c4417d6f8f00050000       vmovdqa 0x500(%r15), %ymm9
  2307: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2312: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
  2316: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  2321: c4817d6f641a40           vmovdqa 0x40(%r10, %r11), %ymm4
  2328: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2333: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2338: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2343: c4417d6f8fc0040000       vmovdqa 0x4c0(%r15), %ymm9
  2352: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2357: c4417d6f8f00050000       vmovdqa 0x500(%r15), %ymm9
  2366: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2371: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
  2375: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  2380: c4817d6f641a60           vmovdqa 0x60(%r10, %r11), %ymm4
  2387: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2392: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2397: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2402: c4417d6f8fc0040000       vmovdqa 0x4c0(%r15), %ymm9
  2411: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2416: c4417d6f8f00050000       vmovdqa 0x500(%r15), %ymm9
  2425: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2430: c5e5efdf                 vpxor %ymm7, %ymm3, %ymm3
  2434: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  2439: 4d8b5050                 movq 0x50(%r8), %r10
  2443: c4817d6f241a             vmovdqa (%r10, %r11), %ymm4
  2449: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2454: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2459: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2464: c4417d6f8f40050000       vmovdqa 0x540(%r15), %ymm9
  2473: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2478: c4417d6f8f80050000       vmovdqa 0x580(%r15), %ymm9
  2487: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2492: c5fdefc7                 vpxor %ymm7, %ymm0, %ymm0
  2496: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
  2501: c4817d6f641a20           vmovdqa 0x20(%r10, %r11), %ymm4
  2508: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2513: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2518: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2523: c4417d6f8f40050000       vmovdqa 0x540(%r15), %ymm9
  2532: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2537: c4417d6f8f80050000       vmovdqa 0x580(%r15), %ymm9
  2546: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2551: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
  2555: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
  2560: c4817d6f641a40           vmovdqa 0x40(%r10, %r11), %ymm4
  2567: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2572: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2577: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2582: c4417d6f8f40050000       vmovdqa 0x540(%r15), %ymm9
  2591: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2596: c4417d6f8f80050000       vmovdqa 0x580(%r15), %ymm9
  2605: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2610: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
  2614: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
  2619: c4817d6f641a60           vmovdqa 0x60(%r10, %r11), %ymm4
  2626: c4c15ddbea               vpand %ymm10, %ymm4, %ymm5
  2631: c5cd71d404               vpsrlw $4, %ymm4, %ymm6
  2636: c4c14ddbf2               vpand %ymm10, %ymm6, %ymm6
  2641: c4417d6f8f40050000       vmovdqa 0x540(%r15), %ymm9
  2650: c4e23500fd               vpshufb %ymm5, %ymm9, %ymm7
  2655: c4417d6f8f80050000       vmovdqa 0x580(%r15), %ymm9
  2664: c4623500c6               vpshufb %ymm6, %ymm9, %ymm8
  2669: c5e5efdf                 vpxor %ymm7, %ymm3, %ymm3
  2673: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
  2678: 4d8b5058                 movq 0x58(%r8), %r10
  2682: c4817d6f241a             vmovdqa (%r10, %r11), %ymm4
  2688: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
  2692: c4817d6f641a20           vmovdqa 0x20(%r10, %r11), %ymm4
  2699: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
  2703: c4817d6f641a40           vmovdqa 0x40(%r10, %r11), %ymm4
  2710: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
  2714: c4817d6f641a60           vmovdqa 0x60(%r10, %r11), %ymm4
  2721: c5e5efdc                 vpxor %ymm4, %ymm3, %ymm3
  2725: 4d8b11                   movq (%r9), %r10
  2728: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
  2734: c4817d7f4c1a20           vmovdqa %ymm1, 0x20(%r10, %r11)
  2741: c4817d7f541a40           vmovdqa %ymm2, 0x40(%r10, %r11)
  2748: c4817d7f5c1a60           vmovdqa %ymm3, 0x60(%r10, %r11)
  2755: 4981c380000000           addq $0x80, %r11
  2762: 4c3bd8                   cmpq %rax, %r11
  2765: 0f868df5ffff             jbe 0x60
  2771: 4c3bd9                   cmpq %rcx, %r11
  2774: 0f8363020000             jae 0xd3f
  2780: 4153                     pushq %r11
  2782: 498dbfc0050000           leaq 0x5c0(%r15), %rdi
  2789: 498db7c0060000           leaq 0x6c0(%r15), %rsi
  2796: 4d8b11                   movq (%r9), %r10
  2799: 4c8b1c24                 movq (%rsp), %r11
  2803: 33db                     xorl %ebx, %ebx
  2805: 498b00                   movq (%r8), %rax
  2808: 420fb61418               movzbl (%rax, %r11), %edx
  2813: 4885d2                   testq %rdx, %rdx
  2816: 0f841f000000             je 0xb25
  2822: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2826: 83c019                   addl $0x19, %eax
  2829: 81f8ff000000             cmpl $0xff, %eax
  2835: 0f8206000000             jb 0xb1f
  2841: 81c001ffffff             addl $0xffffff01, %eax
  2847: 0fb60406                 movzbl (%rsi, %rax), %eax
  2851: 32d8                     xorb %al, %bl
  2853: 498b4008                 movq 8(%r8), %rax
  2857: 420fb61418               movzbl (%rax, %r11), %edx
  2862: 4885d2                   testq %rdx, %rdx
  2865: 0f841f000000             je 0xb56
  2871: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2875: 83c032                   addl $0x32, %eax
  2878: 81f8ff000000             cmpl $0xff, %eax
  2884: 0f8206000000             jb 0xb50
  2890: 81c001ffffff             addl $0xffffff01, %eax
  2896: 0fb60406                 movzbl (%rsi, %rax), %eax
  2900: 32d8                     xorb %al, %bl
  2902: 498b4010                 movq 0x10(%r8), %rax
  2906: 420fb61418               movzbl (%rax, %r11), %edx
  2911: 4885d2                   testq %rdx, %rdx
  2914: 0f841f000000             je 0xb87
  2920: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2924: 83c04b                   addl $0x4b, %eax
  2927: 81f8ff000000             cmpl $0xff, %eax
  2933: 0f8206000000             jb 0xb81
  2939: 81c001ffffff             addl $0xffffff01, %eax
  2945: 0fb60406                 movzbl (%rsi, %rax), %eax
  2949: 32d8                     xorb %al, %bl
  2951: 498b4018                 movq 0x18(%r8), %rax
  2955: 420fb61418               movzbl (%rax, %r11), %edx
  2960: 4885d2                   testq %rdx, %rdx
  2963: 0f841f000000             je 0xbb8
  2969: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2973: 83c064                   addl $0x64, %eax
  2976: 81f8ff000000             cmpl $0xff, %eax
  2982: 0f8206000000             jb 0xbb2
  2988: 81c001ffffff             addl $0xffffff01, %eax
  2994: 0fb60406                 movzbl (%rsi, %rax), %eax
  2998: 32d8                     xorb %al, %bl
  3000: 498b4020                 movq 0x20(%r8), %rax
  3004: 420fb61418               movzbl (%rax, %r11), %edx
  3009: 4885d2                   testq %rdx, %rdx
  3012: 0f841f000000             je 0xbe9
  3018: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3022: 83c07d                   addl $0x7d, %eax
  3025: 81f8ff000000             cmpl $0xff, %eax
  3031: 0f8206000000             jb 0xbe3
  3037: 81c001ffffff             addl $0xffffff01, %eax
  3043: 0fb60406                 movzbl (%rsi, %rax), %eax
  3047: 32d8                     xorb %al, %bl
  3049: 498b4028                 movq 0x28(%r8), %rax
  3053: 420fb61418               movzbl (%rax, %r11), %edx
  3058: 4885d2                   testq %rdx, %rdx
  3061: 0f8422000000             je 0xc1d
  3067: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3071: 81c096000000             addl $0x96, %eax
  3077: 81f8ff000000             cmpl $0xff, %eax
  3083: 0f8206000000             jb 0xc17
  3089: 81c001ffffff             addl $0xffffff01, %eax
  3095: 0fb60406                 movzbl (%rsi, %rax), %eax
  3099: 32d8                     xorb %al, %bl
  3101: 498b4030                 movq 0x30(%r8), %rax
  3105: 420fb61418               movzbl (%rax, %r11), %edx
  3110: 4885d2                   testq %rdx, %rdx
  3113: 0f8422000000             je 0xc51
  3119: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3123: 81c0af000000             addl $0xaf, %eax
  3129: 81f8ff000000             cmpl $0xff, %eax
  3135: 0f8206000000             jb 0xc4b
  3141: 81c001ffffff             addl $0xffffff01, %eax
  3147: 0fb60406                 movzbl (%rsi, %rax), %eax
  3151: 32d8                     xorb %al, %bl
  3153: 498b4038                 movq 0x38(%r8), %rax
  3157: 420fb61418               movzbl (%rax, %r11), %edx
  3162: 4885d2                   testq %rdx, %rdx
  3165: 0f8422000000             je 0xc85
  3171: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3175: 81c0c8000000             addl $0xc8, %eax
  3181: 81f8ff000000             cmpl $0xff, %eax
  3187: 0f8206000000             jb 0xc7f
  3193: 81c001ffffff             addl $0xffffff01, %eax
  3199: 0fb60406                 movzbl (%rsi, %rax), %eax
  3203: 32d8                     xorb %al, %bl
  3205: 498b4040                 movq 0x40(%r8), %rax
  3209: 420fb61418               movzbl (%rax, %r11), %edx
  3214: 4885d2                   testq %rdx, %rdx
  3217: 0f8422000000             je 0xcb9
  3223: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3227: 81c0e1000000             addl $0xe1, %eax
  3233: 81f8ff000000             cmpl $0xff, %eax
  3239: 0f8206000000             jb 0xcb3
  3245: 81c001ffffff             addl $0xffffff01, %eax
  3251: 0fb60406                 movzbl (%rsi, %rax), %eax
  3255: 32d8                     xorb %al, %bl
  3257: 498b4048                 movq 0x48(%r8), %rax
  3261: 420fb61418               movzbl (%rax, %r11), %edx
  3266: 4885d2                   testq %rdx, %rdx
  3269: 0f8422000000             je 0xced
  3275: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3279: 81c0fa000000             addl $0xfa, %eax
  3285: 81f8ff000000             cmpl $0xff, %eax
  3291: 0f8206000000             jb 0xce7
  3297: 81c001ffffff             addl $0xffffff01, %eax
  3303: 0fb60406                 movzbl (%rsi, %rax), %eax
  3307: 32d8                     xorb %al, %bl
  3309: 498b4050                 movq 0x50(%r8), %rax
  3313: 420fb61418               movzbl (%rax, %r11), %edx
  3318: 4885d2                   testq %rdx, %rdx
  3321: 0f841f000000             je 0xd1e
  3327: 0fb60417                 movzbl (%rdi, %rdx), %eax
  3331: 83c014                   addl $0x14, %eax
  3334: 81f8ff000000             cmpl $0xff, %eax
  3340: 0f8206000000             jb 0xd18
  3346: 81c001ffffff             addl $0xffffff01, %eax
  3352: 0fb60406                 movzbl (%rsi, %rax), %eax
  3356: 32d8                     xorb %al, %bl
  3358: 498b4058                 movq 0x58(%r8), %rax
  3362: 420fb61418               movzbl (%rax, %r11), %edx
  3367: 32da                     xorb %dl, %bl
  3369: 43881c1a                 movb %bl, (%r10, %r11)
  3373: 33db                     xorl %ebx, %ebx
  3375: 49ffc3                   incq %r11
  3378: 4c3bd9                   cmpq %rcx, %r11
  3381: 0f82bafdffff             jb 0xaf5
  3387: 4883c408                 addq $8, %rsp
  3391: 415e                     popq %r14
  3393: 415d                     popq %r13
  3395: 415c                     popq %r12
  3397: 415f                     popq %r15
  3399: 5b                       popq %rbx
  3400: c5f877                   vzeroupper 
  3403: c3                       retq 
  3404: 0f0b                     ud2 
  3406: 0f0b                     ud2 
  3408: 0f0b                     ud2 
  3410: 0f0b                     ud2 
  3412: 0f0b                     ud2 
  3414: 0f0b                     ud2 
  3416: 0f0b                     ud2 
  3418: 0f0b                     ud2 
  3420: 0f0b                     ud2 
  3422: 0f0b                     ud2 
  3424: 0f0b                     ud2 
  3426: 0f0b                     ud2 
  3428: 0f0b                     ud2 
  3430: 0f0b                     ud2 
  3432: 0f0b                     ud2 
  3434: 0f0b                     ud2 
  3436: 0f0b                     ud2 
  3438: 0f0b                     ud2 
  3440: 0f0b                     ud2 
  3442: 0f0b                     ud2 
  3444: 0f0b                     ud2 
  3446: 0f0b                     ud2 
  3448: 0f0b                     ud2 
  3450: 0f0b                     ud2 
  3452: 0f0b                     ud2 
  3454: 0f0b                     ud2 
