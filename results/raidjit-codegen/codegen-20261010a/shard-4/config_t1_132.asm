; config_t1_132 target=avx512-gfni twin=ternlog topo=liberation(8)/enc op=encode U=1
; code_size=5184 twins=["ternlog"] uops_per_byte=4.796875
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
    41: 4883e840                 subq $0x40, %rax
    45: 0f82b8080000             jb 0x8eb
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
    70: 62f17548efc9             vpxord %zmm1, %zmm1, %zmm1
    76: 62f16d48efd2             vpxord %zmm2, %zmm2, %zmm2
    82: 62f16548efdb             vpxord %zmm3, %zmm3, %zmm3
    88: 62f15d48efe4             vpxord %zmm4, %zmm4, %zmm4
    94: 62f15548efed             vpxord %zmm5, %zmm5, %zmm5
   100: 62f14d48eff6             vpxord %zmm6, %zmm6, %zmm6
   106: 62f14548efff             vpxord %zmm7, %zmm7, %zmm7
   112: 62513d48efc0             vpxord %zmm8, %zmm8, %zmm8
   118: 62513548efc9             vpxord %zmm9, %zmm9, %zmm9
   124: 62512d48efd2             vpxord %zmm10, %zmm10, %zmm10
   130: 62512548efdb             vpxord %zmm11, %zmm11, %zmm11
   136: 62511d48efe4             vpxord %zmm12, %zmm12, %zmm12
   142: 62511548efed             vpxord %zmm13, %zmm13, %zmm13
   148: 62510d48eff6             vpxord %zmm14, %zmm14, %zmm14
   154: 62510548efff             vpxord %zmm15, %zmm15, %zmm15
   160: 62a17d40efc0             vpxord %zmm16, %zmm16, %zmm16
   166: 62a17540efc9             vpxord %zmm17, %zmm17, %zmm17
   172: 62a16d40efd2             vpxord %zmm18, %zmm18, %zmm18
   178: 62a16540efdb             vpxord %zmm19, %zmm19, %zmm19
   184: 62a15d40efe4             vpxord %zmm20, %zmm20, %zmm20
   190: 62a15540efed             vpxord %zmm21, %zmm21, %zmm21
   196: 4d8b5028                 movq 0x28(%r8), %r10
   200: 498b9080000000           movq 0x80(%r8), %rdx
   207: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   214: 62b34d40252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm5
   222: 62310548ef3c1a           vpxord (%rdx, %r11), %zmm15, %zmm15
   229: 62a34d4025041a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm16
   237: 4d8b5058                 movq 0x58(%r8), %r10
   241: 498b90b0000000           movq 0xb0(%r8), %rdx
   248: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   255: 62b34d4025041a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm0
   263: 62a15d40ef241a           vpxord (%rdx, %r11), %zmm20, %zmm20
   270: 62a34d40252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm21
   278: 4d8b90e0000000           movq 0xe0(%r8), %r10
   285: 498b9038010000           movq 0x138(%r8), %rdx
   292: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   299: 62b34d4025341a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm6
   307: 62310d48ef341a           vpxord (%rdx, %r11), %zmm14, %zmm14
   314: 62334d40253c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm15
   322: 4d8b9010010000           movq 0x110(%r8), %r10
   329: 498b9068010000           movq 0x168(%r8), %rdx
   336: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   343: 62b34d40250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm1
   351: 62a16540ef1c1a           vpxord (%rdx, %r11), %zmm19, %zmm19
   358: 62a34d4025241a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm20
   366: 4d8b9098010000           movq 0x198(%r8), %r10
   373: 498b90f0010000           movq 0x1f0(%r8), %rdx
   380: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   387: 62b34d40253c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm7
   395: 62311548ef2c1a           vpxord (%rdx, %r11), %zmm13, %zmm13
   402: 62334d4025341a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm14
   410: 4d8b90c8010000           movq 0x1c8(%r8), %r10
   417: 498b9020020000           movq 0x220(%r8), %rdx
   424: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   431: 62b34d4025141a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm2
   439: 62a16d40ef141a           vpxord (%rdx, %r11), %zmm18, %zmm18
   446: 62a34d40251c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm19
   454: 4d8b9050020000           movq 0x250(%r8), %r10
   461: 498b90a8020000           movq 0x2a8(%r8), %rdx
   468: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   475: 62334d4025041a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm8
   483: 62311d48ef241a           vpxord (%rdx, %r11), %zmm12, %zmm12
   490: 62334d40252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm13
   498: 498b5060                 movq 0x60(%r8), %rdx
   502: 62a17d486f341e           vmovdqa32 (%rsi, %r11), %zmm22
   509: 62b17d48efc6             vpxord %zmm22, %zmm0, %zmm0
   515: 62b17548ef0c1a           vpxord (%rdx, %r11), %zmm1, %zmm1
   522: 62334d40251c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm11
   530: 498b5068                 movq 0x68(%r8), %rdx
   534: 62a17d486f341f           vmovdqa32 (%rdi, %r11), %zmm22
   541: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
   547: 62b16d48ef141a           vpxord (%rdx, %r11), %zmm2, %zmm2
   554: 62334d4025241a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm12
   562: 498b5070                 movq 0x70(%r8), %rdx
   566: 62817d486f341c           vmovdqa32 (%r12, %r11), %zmm22
   573: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
   579: 62b16548ef1c1a           vpxord (%rdx, %r11), %zmm3, %zmm3
   586: 62334d40252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm13
   594: 498b5078                 movq 0x78(%r8), %rdx
   598: 62817d486f741d00         vmovdqa32 (%r13, %r11), %zmm22
   606: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
   612: 62b15d48ef241a           vpxord (%rdx, %r11), %zmm4, %zmm4
   619: 62334d4025341a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm14
   627: 498b90d0000000           movq 0xd0(%r8), %rdx
   634: 62817d486f341e           vmovdqa32 (%r14, %r11), %zmm22
   641: 62b34d4025241a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm4
   649: 62311548ef2c1a           vpxord (%rdx, %r11), %zmm13, %zmm13
   656: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
   662: 4d8b5030                 movq 0x30(%r8), %r10
   666: 498b9088000000           movq 0x88(%r8), %rdx
   673: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   680: 62b34d4025341a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm6
   688: 62a17d40ef041a           vpxord (%rdx, %r11), %zmm16, %zmm16
   695: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
   701: 4d8b5038                 movq 0x38(%r8), %r10
   705: 498b9090000000           movq 0x90(%r8), %rdx
   712: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   719: 62b34d40253c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm7
   727: 62a17540ef0c1a           vpxord (%rdx, %r11), %zmm17, %zmm17
   734: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
   740: 4d8b5040                 movq 0x40(%r8), %r10
   744: 498b9098000000           movq 0x98(%r8), %rdx
   751: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   758: 62334d4025041a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm8
   766: 62a16d40ef141a           vpxord (%rdx, %r11), %zmm18, %zmm18
   773: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
   779: 4d8b5048                 movq 0x48(%r8), %r10
   783: 498b90a0000000           movq 0xa0(%r8), %rdx
   790: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   797: 62334d40250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm9
   805: 62a16540ef1c1a           vpxord (%rdx, %r11), %zmm19, %zmm19
   812: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
   818: 4d8b5050                 movq 0x50(%r8), %r10
   822: 498b90a8000000           movq 0xa8(%r8), %rdx
   829: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   836: 62334d4025141a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm10
   844: 62a15d40ef241a           vpxord (%rdx, %r11), %zmm20, %zmm20
   851: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
   857: 4d8b90b8000000           movq 0xb8(%r8), %r10
   864: 498b9018010000           movq 0x118(%r8), %rdx
   871: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   878: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
   884: 62b16d48ef141a           vpxord (%rdx, %r11), %zmm2, %zmm2
   891: 62a34d40252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm21
   899: 4d8b90c0000000           movq 0xc0(%r8), %r10
   906: 498b9020010000           movq 0x120(%r8), %rdx
   913: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   920: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
   926: 62b16548ef1c1a           vpxord (%rdx, %r11), %zmm3, %zmm3
   933: 62334d40251c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm11
   941: 4d8b90c8000000           movq 0xc8(%r8), %r10
   948: 498b9028010000           movq 0x128(%r8), %rdx
   955: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
   962: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
   968: 62b15d48ef241a           vpxord (%rdx, %r11), %zmm4, %zmm4
   975: 62334d4025241a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm12
   983: 4d8b90d8000000           movq 0xd8(%r8), %r10
   990: 498b9030010000           movq 0x130(%r8), %rdx
   997: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1004: 62b34d40252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm5
  1012: 62311548ef2c1a           vpxord (%rdx, %r11), %zmm13, %zmm13
  1019: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  1025: 4d8b90e8000000           movq 0xe8(%r8), %r10
  1032: 498b9040010000           movq 0x140(%r8), %rdx
  1039: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1046: 62b34d40253c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm7
  1054: 62310548ef3c1a           vpxord (%rdx, %r11), %zmm15, %zmm15
  1061: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1067: 4d8b90f0000000           movq 0xf0(%r8), %r10
  1074: 498b9048010000           movq 0x148(%r8), %rdx
  1081: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1088: 62334d4025041a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm8
  1096: 62a17d40ef041a           vpxord (%rdx, %r11), %zmm16, %zmm16
  1103: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1109: 4d8b90f8000000           movq 0xf8(%r8), %r10
  1116: 498b9050010000           movq 0x150(%r8), %rdx
  1123: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1130: 62334d40250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm9
  1138: 62a17540ef0c1a           vpxord (%rdx, %r11), %zmm17, %zmm17
  1145: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1151: 4d8b9000010000           movq 0x100(%r8), %r10
  1158: 498b9008010000           movq 0x108(%r8), %rdx
  1165: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1172: 62b17d48ef041a           vpxord (%rdx, %r11), %zmm0, %zmm0
  1179: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1185: 62a34d40251c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm19
  1193: 4d8b9058010000           movq 0x158(%r8), %r10
  1200: 498b9060010000           movq 0x160(%r8), %rdx
  1207: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1214: 62b17d48ef041a           vpxord (%rdx, %r11), %zmm0, %zmm0
  1221: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1227: 62a34d4025141a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm18
  1235: 4d8b9070010000           movq 0x170(%r8), %r10
  1242: 498b90d0010000           movq 0x1d0(%r8), %rdx
  1249: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1256: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  1262: 62b16548ef1c1a           vpxord (%rdx, %r11), %zmm3, %zmm3
  1269: 62a34d4025241a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm20
  1277: 4d8b9078010000           movq 0x178(%r8), %r10
  1284: 498b90d8010000           movq 0x1d8(%r8), %rdx
  1291: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1298: 62b16548efde             vpxord %zmm22, %zmm3, %zmm3
  1304: 62b15d48ef241a           vpxord (%rdx, %r11), %zmm4, %zmm4
  1311: 62a34d40252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm21
  1319: 4d8b9080010000           movq 0x180(%r8), %r10
  1326: 498b90e0010000           movq 0x1e0(%r8), %rdx
  1333: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1340: 62b15d48efe6             vpxord %zmm22, %zmm4, %zmm4
  1346: 62b15548ef2c1a           vpxord (%rdx, %r11), %zmm5, %zmm5
  1353: 62334d40251c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm11
  1361: 4d8b9088010000           movq 0x188(%r8), %r10
  1368: 498b90e8010000           movq 0x1e8(%r8), %rdx
  1375: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1382: 62b15548efee             vpxord %zmm22, %zmm5, %zmm5
  1388: 62b14d48ef341a           vpxord (%rdx, %r11), %zmm6, %zmm6
  1395: 62334d4025241a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm12
  1403: 4d8b9090010000           movq 0x190(%r8), %r10
  1410: 498b9040020000           movq 0x240(%r8), %rdx
  1417: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1424: 62b34d4025341a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm6
  1432: 62312548ef1c1a           vpxord (%rdx, %r11), %zmm11, %zmm11
  1439: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  1445: 4d8b90a0010000           movq 0x1a0(%r8), %r10
  1452: 498b90f8010000           movq 0x1f8(%r8), %rdx
  1459: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1466: 62334d4025041a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm8
  1474: 62310d48ef341a           vpxord (%rdx, %r11), %zmm14, %zmm14
  1481: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  1487: 4d8b90a8010000           movq 0x1a8(%r8), %r10
  1494: 498b9000020000           movq 0x200(%r8), %rdx
  1501: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1508: 62334d40250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm9
  1516: 62310548ef3c1a           vpxord (%rdx, %r11), %zmm15, %zmm15
  1523: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1529: 4d8b90b0010000           movq 0x1b0(%r8), %r10
  1536: 498b90b8010000           movq 0x1b8(%r8), %rdx
  1543: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1550: 62b17d48ef041a           vpxord (%rdx, %r11), %zmm0, %zmm0
  1557: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1563: 62a34d40250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm17
  1571: 4d8b90c0010000           movq 0x1c0(%r8), %r10
  1578: 498b9018020000           movq 0x218(%r8), %rdx
  1585: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1592: 62b34d40250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm1
  1600: 62a17540ef0c1a           vpxord (%rdx, %r11), %zmm17, %zmm17
  1607: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1613: 4d8b9008020000           movq 0x208(%r8), %r10
  1620: 498b9010020000           movq 0x210(%r8), %rdx
  1627: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1634: 62b17d48ef041a           vpxord (%rdx, %r11), %zmm0, %zmm0
  1641: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1647: 62a34d4025041a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm16
  1655: 4d8b9028020000           movq 0x228(%r8), %r10
  1662: 498b9080020000           movq 0x280(%r8), %rdx
  1669: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1676: 62b34d40251c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm3
  1684: 62a16d40ef141a           vpxord (%rdx, %r11), %zmm18, %zmm18
  1691: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1697: 4d8b9030020000           movq 0x230(%r8), %r10
  1704: 498b9088020000           movq 0x288(%r8), %rdx
  1711: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1718: 62b34d4025241a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm4
  1726: 62a16540ef1c1a           vpxord (%rdx, %r11), %zmm19, %zmm19
  1733: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1739: 4d8b9038020000           movq 0x238(%r8), %r10
  1746: 498b9090020000           movq 0x290(%r8), %rdx
  1753: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1760: 62b34d40252c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm5
  1768: 62a15d40ef241a           vpxord (%rdx, %r11), %zmm20, %zmm20
  1775: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  1781: 4d8b9048020000           movq 0x248(%r8), %r10
  1788: 498b90a0020000           movq 0x2a0(%r8), %rdx
  1795: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1802: 62b34d40253c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm7
  1810: 62312548ef1c1a           vpxord (%rdx, %r11), %zmm11, %zmm11
  1817: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1823: 4d8b9058020000           movq 0x258(%r8), %r10
  1830: 498b90b0020000           movq 0x2b0(%r8), %rdx
  1837: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1844: 62334d40250c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm9
  1852: 62311548ef2c1a           vpxord (%rdx, %r11), %zmm13, %zmm13
  1859: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  1865: 4d8b9060020000           movq 0x260(%r8), %r10
  1872: 498b9068020000           movq 0x268(%r8), %rdx
  1879: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1886: 62b17d48ef041a           vpxord (%rdx, %r11), %zmm0, %zmm0
  1893: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1899: 62334d40253c1a96         vpternlogd $0x96, (%rdx, %r11), %zmm22, %zmm15
  1907: 4d8b9070020000           movq 0x270(%r8), %r10
  1914: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1921: 62b17548efce             vpxord %zmm22, %zmm1, %zmm1
  1927: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1933: 4d8b9078020000           movq 0x278(%r8), %r10
  1940: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1947: 62b16d48efd6             vpxord %zmm22, %zmm2, %zmm2
  1953: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1959: 4d8b9098020000           movq 0x298(%r8), %r10
  1966: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1973: 62b14d48eff6             vpxord %zmm22, %zmm6, %zmm6
  1979: 62a15540efee             vpxord %zmm22, %zmm21, %zmm21
  1985: 4d8b90b8020000           movq 0x2b8(%r8), %r10
  1992: 62817d486f341a           vmovdqa32 (%r10, %r11), %zmm22
  1999: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  2005: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  2011: 4d8b11                   movq (%r9), %r10
  2014: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  2021: 4d8b5108                 movq 8(%r9), %r10
  2025: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
  2032: 4d8b5110                 movq 0x10(%r9), %r10
  2036: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
  2043: 4d8b5118                 movq 0x18(%r9), %r10
  2047: 62917d487f1c1a           vmovdqa32 %zmm3, (%r10, %r11)
  2054: 4d8b5120                 movq 0x20(%r9), %r10
  2058: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
  2065: 4d8b5128                 movq 0x28(%r9), %r10
  2069: 62917d487f2c1a           vmovdqa32 %zmm5, (%r10, %r11)
  2076: 4d8b5130                 movq 0x30(%r9), %r10
  2080: 62917d487f341a           vmovdqa32 %zmm6, (%r10, %r11)
  2087: 4d8b5138                 movq 0x38(%r9), %r10
  2091: 62917d487f3c1a           vmovdqa32 %zmm7, (%r10, %r11)
  2098: 4d8b5140                 movq 0x40(%r9), %r10
  2102: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  2109: 4d8b5148                 movq 0x48(%r9), %r10
  2113: 62117d487f0c1a           vmovdqa32 %zmm9, (%r10, %r11)
  2120: 4d8b5150                 movq 0x50(%r9), %r10
  2124: 62117d487f141a           vmovdqa32 %zmm10, (%r10, %r11)
  2131: 4d8b5158                 movq 0x58(%r9), %r10
  2135: 62117d487f1c1a           vmovdqa32 %zmm11, (%r10, %r11)
  2142: 4d8b5160                 movq 0x60(%r9), %r10
  2146: 62117d487f241a           vmovdqa32 %zmm12, (%r10, %r11)
  2153: 4d8b5168                 movq 0x68(%r9), %r10
  2157: 62117d487f2c1a           vmovdqa32 %zmm13, (%r10, %r11)
  2164: 4d8b5170                 movq 0x70(%r9), %r10
  2168: 62117d487f341a           vmovdqa32 %zmm14, (%r10, %r11)
  2175: 4d8b5178                 movq 0x78(%r9), %r10
  2179: 62117d487f3c1a           vmovdqa32 %zmm15, (%r10, %r11)
  2186: 4d8b9180000000           movq 0x80(%r9), %r10
  2193: 62817d487f041a           vmovdqa32 %zmm16, (%r10, %r11)
  2200: 4d8b9188000000           movq 0x88(%r9), %r10
  2207: 62817d487f0c1a           vmovdqa32 %zmm17, (%r10, %r11)
  2214: 4d8b9190000000           movq 0x90(%r9), %r10
  2221: 62817d487f141a           vmovdqa32 %zmm18, (%r10, %r11)
  2228: 4d8b9198000000           movq 0x98(%r9), %r10
  2235: 62817d487f1c1a           vmovdqa32 %zmm19, (%r10, %r11)
  2242: 4d8b91a0000000           movq 0xa0(%r9), %r10
  2249: 62817d487f241a           vmovdqa32 %zmm20, (%r10, %r11)
  2256: 4d8b91a8000000           movq 0xa8(%r9), %r10
  2263: 62817d487f2c1a           vmovdqa32 %zmm21, (%r10, %r11)
  2270: 4983c340                 addq $0x40, %r11
  2274: 4c3bd8                   cmpq %rax, %r11
  2277: 0f8655f7ffff             jbe 0x40
  2283: 4c3bd9                   cmpq %rcx, %r11
  2286: 0f833c0b0000             jae 0x1430
  2292: 4153                     pushq %r11
  2294: 4d8b11                   movq (%r9), %r10
  2297: 4c8b1c24                 movq (%rsp), %r11
  2301: 498b00                   movq (%r8), %rax
  2304: 428a1418                 movb (%rax, %r11), %dl
  2308: 8ada                     movb %dl, %bl
  2310: 498b4058                 movq 0x58(%r8), %rax
  2314: 428a1418                 movb (%rax, %r11), %dl
  2318: 32da                     xorb %dl, %bl
  2320: 498b80b0000000           movq 0xb0(%r8), %rax
  2327: 428a1418                 movb (%rax, %r11), %dl
  2331: 32da                     xorb %dl, %bl
  2333: 498b8008010000           movq 0x108(%r8), %rax
  2340: 428a1418                 movb (%rax, %r11), %dl
  2344: 32da                     xorb %dl, %bl
  2346: 498b8060010000           movq 0x160(%r8), %rax
  2353: 428a1418                 movb (%rax, %r11), %dl
  2357: 32da                     xorb %dl, %bl
  2359: 498b80b8010000           movq 0x1b8(%r8), %rax
  2366: 428a1418                 movb (%rax, %r11), %dl
  2370: 32da                     xorb %dl, %bl
  2372: 498b8010020000           movq 0x210(%r8), %rax
  2379: 428a1418                 movb (%rax, %r11), %dl
  2383: 32da                     xorb %dl, %bl
  2385: 498b8068020000           movq 0x268(%r8), %rax
  2392: 428a1418                 movb (%rax, %r11), %dl
  2396: 32da                     xorb %dl, %bl
  2398: 43881c1a                 movb %bl, (%r10, %r11)
  2402: 33db                     xorl %ebx, %ebx
  2404: 49ffc3                   incq %r11
  2407: 4c3bd9                   cmpq %rcx, %r11
  2410: 0f828dffffff             jb 0x8fd
  2416: 4d8b5108                 movq 8(%r9), %r10
  2420: 4c8b1c24                 movq (%rsp), %r11
  2424: 498b4008                 movq 8(%r8), %rax
  2428: 428a1418                 movb (%rax, %r11), %dl
  2432: 8ada                     movb %dl, %bl
  2434: 498b4060                 movq 0x60(%r8), %rax
  2438: 428a1418                 movb (%rax, %r11), %dl
  2442: 32da                     xorb %dl, %bl
  2444: 498b80b8000000           movq 0xb8(%r8), %rax
  2451: 428a1418                 movb (%rax, %r11), %dl
  2455: 32da                     xorb %dl, %bl
  2457: 498b8010010000           movq 0x110(%r8), %rax
  2464: 428a1418                 movb (%rax, %r11), %dl
  2468: 32da                     xorb %dl, %bl
  2470: 498b8068010000           movq 0x168(%r8), %rax
  2477: 428a1418                 movb (%rax, %r11), %dl
  2481: 32da                     xorb %dl, %bl
  2483: 498b80c0010000           movq 0x1c0(%r8), %rax
  2490: 428a1418                 movb (%rax, %r11), %dl
  2494: 32da                     xorb %dl, %bl
  2496: 498b8018020000           movq 0x218(%r8), %rax
  2503: 428a1418                 movb (%rax, %r11), %dl
  2507: 32da                     xorb %dl, %bl
  2509: 498b8070020000           movq 0x270(%r8), %rax
  2516: 428a1418                 movb (%rax, %r11), %dl
  2520: 32da                     xorb %dl, %bl
  2522: 43881c1a                 movb %bl, (%r10, %r11)
  2526: 33db                     xorl %ebx, %ebx
  2528: 49ffc3                   incq %r11
  2531: 4c3bd9                   cmpq %rcx, %r11
  2534: 0f828cffffff             jb 0x978
  2540: 4d8b5110                 movq 0x10(%r9), %r10
  2544: 4c8b1c24                 movq (%rsp), %r11
  2548: 498b4010                 movq 0x10(%r8), %rax
  2552: 428a1418                 movb (%rax, %r11), %dl
  2556: 8ada                     movb %dl, %bl
  2558: 498b4068                 movq 0x68(%r8), %rax
  2562: 428a1418                 movb (%rax, %r11), %dl
  2566: 32da                     xorb %dl, %bl
  2568: 498b80c0000000           movq 0xc0(%r8), %rax
  2575: 428a1418                 movb (%rax, %r11), %dl
  2579: 32da                     xorb %dl, %bl
  2581: 498b8018010000           movq 0x118(%r8), %rax
  2588: 428a1418                 movb (%rax, %r11), %dl
  2592: 32da                     xorb %dl, %bl
  2594: 498b8070010000           movq 0x170(%r8), %rax
  2601: 428a1418                 movb (%rax, %r11), %dl
  2605: 32da                     xorb %dl, %bl
  2607: 498b80c8010000           movq 0x1c8(%r8), %rax
  2614: 428a1418                 movb (%rax, %r11), %dl
  2618: 32da                     xorb %dl, %bl
  2620: 498b8020020000           movq 0x220(%r8), %rax
  2627: 428a1418                 movb (%rax, %r11), %dl
  2631: 32da                     xorb %dl, %bl
  2633: 498b8078020000           movq 0x278(%r8), %rax
  2640: 428a1418                 movb (%rax, %r11), %dl
  2644: 32da                     xorb %dl, %bl
  2646: 43881c1a                 movb %bl, (%r10, %r11)
  2650: 33db                     xorl %ebx, %ebx
  2652: 49ffc3                   incq %r11
  2655: 4c3bd9                   cmpq %rcx, %r11
  2658: 0f828cffffff             jb 0x9f4
  2664: 4d8b5118                 movq 0x18(%r9), %r10
  2668: 4c8b1c24                 movq (%rsp), %r11
  2672: 498b4018                 movq 0x18(%r8), %rax
  2676: 428a1418                 movb (%rax, %r11), %dl
  2680: 8ada                     movb %dl, %bl
  2682: 498b4070                 movq 0x70(%r8), %rax
  2686: 428a1418                 movb (%rax, %r11), %dl
  2690: 32da                     xorb %dl, %bl
  2692: 498b80c8000000           movq 0xc8(%r8), %rax
  2699: 428a1418                 movb (%rax, %r11), %dl
  2703: 32da                     xorb %dl, %bl
  2705: 498b8020010000           movq 0x120(%r8), %rax
  2712: 428a1418                 movb (%rax, %r11), %dl
  2716: 32da                     xorb %dl, %bl
  2718: 498b8078010000           movq 0x178(%r8), %rax
  2725: 428a1418                 movb (%rax, %r11), %dl
  2729: 32da                     xorb %dl, %bl
  2731: 498b80d0010000           movq 0x1d0(%r8), %rax
  2738: 428a1418                 movb (%rax, %r11), %dl
  2742: 32da                     xorb %dl, %bl
  2744: 498b8028020000           movq 0x228(%r8), %rax
  2751: 428a1418                 movb (%rax, %r11), %dl
  2755: 32da                     xorb %dl, %bl
  2757: 498b8080020000           movq 0x280(%r8), %rax
  2764: 428a1418                 movb (%rax, %r11), %dl
  2768: 32da                     xorb %dl, %bl
  2770: 43881c1a                 movb %bl, (%r10, %r11)
  2774: 33db                     xorl %ebx, %ebx
  2776: 49ffc3                   incq %r11
  2779: 4c3bd9                   cmpq %rcx, %r11
  2782: 0f828cffffff             jb 0xa70
  2788: 4d8b5120                 movq 0x20(%r9), %r10
  2792: 4c8b1c24                 movq (%rsp), %r11
  2796: 498b4020                 movq 0x20(%r8), %rax
  2800: 428a1418                 movb (%rax, %r11), %dl
  2804: 8ada                     movb %dl, %bl
  2806: 498b4078                 movq 0x78(%r8), %rax
  2810: 428a1418                 movb (%rax, %r11), %dl
  2814: 32da                     xorb %dl, %bl
  2816: 498b80d0000000           movq 0xd0(%r8), %rax
  2823: 428a1418                 movb (%rax, %r11), %dl
  2827: 32da                     xorb %dl, %bl
  2829: 498b8028010000           movq 0x128(%r8), %rax
  2836: 428a1418                 movb (%rax, %r11), %dl
  2840: 32da                     xorb %dl, %bl
  2842: 498b8080010000           movq 0x180(%r8), %rax
  2849: 428a1418                 movb (%rax, %r11), %dl
  2853: 32da                     xorb %dl, %bl
  2855: 498b80d8010000           movq 0x1d8(%r8), %rax
  2862: 428a1418                 movb (%rax, %r11), %dl
  2866: 32da                     xorb %dl, %bl
  2868: 498b8030020000           movq 0x230(%r8), %rax
  2875: 428a1418                 movb (%rax, %r11), %dl
  2879: 32da                     xorb %dl, %bl
  2881: 498b8088020000           movq 0x288(%r8), %rax
  2888: 428a1418                 movb (%rax, %r11), %dl
  2892: 32da                     xorb %dl, %bl
  2894: 43881c1a                 movb %bl, (%r10, %r11)
  2898: 33db                     xorl %ebx, %ebx
  2900: 49ffc3                   incq %r11
  2903: 4c3bd9                   cmpq %rcx, %r11
  2906: 0f828cffffff             jb 0xaec
  2912: 4d8b5128                 movq 0x28(%r9), %r10
  2916: 4c8b1c24                 movq (%rsp), %r11
  2920: 498b4028                 movq 0x28(%r8), %rax
  2924: 428a1418                 movb (%rax, %r11), %dl
  2928: 8ada                     movb %dl, %bl
  2930: 498b8080000000           movq 0x80(%r8), %rax
  2937: 428a1418                 movb (%rax, %r11), %dl
  2941: 32da                     xorb %dl, %bl
  2943: 498b80d8000000           movq 0xd8(%r8), %rax
  2950: 428a1418                 movb (%rax, %r11), %dl
  2954: 32da                     xorb %dl, %bl
  2956: 498b8030010000           movq 0x130(%r8), %rax
  2963: 428a1418                 movb (%rax, %r11), %dl
  2967: 32da                     xorb %dl, %bl
  2969: 498b8088010000           movq 0x188(%r8), %rax
  2976: 428a1418                 movb (%rax, %r11), %dl
  2980: 32da                     xorb %dl, %bl
  2982: 498b80e0010000           movq 0x1e0(%r8), %rax
  2989: 428a1418                 movb (%rax, %r11), %dl
  2993: 32da                     xorb %dl, %bl
  2995: 498b8038020000           movq 0x238(%r8), %rax
  3002: 428a1418                 movb (%rax, %r11), %dl
  3006: 32da                     xorb %dl, %bl
  3008: 498b8090020000           movq 0x290(%r8), %rax
  3015: 428a1418                 movb (%rax, %r11), %dl
  3019: 32da                     xorb %dl, %bl
  3021: 43881c1a                 movb %bl, (%r10, %r11)
  3025: 33db                     xorl %ebx, %ebx
  3027: 49ffc3                   incq %r11
  3030: 4c3bd9                   cmpq %rcx, %r11
  3033: 0f8289ffffff             jb 0xb68
  3039: 4d8b5130                 movq 0x30(%r9), %r10
  3043: 4c8b1c24                 movq (%rsp), %r11
  3047: 498b4030                 movq 0x30(%r8), %rax
  3051: 428a1418                 movb (%rax, %r11), %dl
  3055: 8ada                     movb %dl, %bl
  3057: 498b8088000000           movq 0x88(%r8), %rax
  3064: 428a1418                 movb (%rax, %r11), %dl
  3068: 32da                     xorb %dl, %bl
  3070: 498b80e0000000           movq 0xe0(%r8), %rax
  3077: 428a1418                 movb (%rax, %r11), %dl
  3081: 32da                     xorb %dl, %bl
  3083: 498b8038010000           movq 0x138(%r8), %rax
  3090: 428a1418                 movb (%rax, %r11), %dl
  3094: 32da                     xorb %dl, %bl
  3096: 498b8090010000           movq 0x190(%r8), %rax
  3103: 428a1418                 movb (%rax, %r11), %dl
  3107: 32da                     xorb %dl, %bl
  3109: 498b80e8010000           movq 0x1e8(%r8), %rax
  3116: 428a1418                 movb (%rax, %r11), %dl
  3120: 32da                     xorb %dl, %bl
  3122: 498b8040020000           movq 0x240(%r8), %rax
  3129: 428a1418                 movb (%rax, %r11), %dl
  3133: 32da                     xorb %dl, %bl
  3135: 498b8098020000           movq 0x298(%r8), %rax
  3142: 428a1418                 movb (%rax, %r11), %dl
  3146: 32da                     xorb %dl, %bl
  3148: 43881c1a                 movb %bl, (%r10, %r11)
  3152: 33db                     xorl %ebx, %ebx
  3154: 49ffc3                   incq %r11
  3157: 4c3bd9                   cmpq %rcx, %r11
  3160: 0f8289ffffff             jb 0xbe7
  3166: 4d8b5138                 movq 0x38(%r9), %r10
  3170: 4c8b1c24                 movq (%rsp), %r11
  3174: 498b4038                 movq 0x38(%r8), %rax
  3178: 428a1418                 movb (%rax, %r11), %dl
  3182: 8ada                     movb %dl, %bl
  3184: 498b8090000000           movq 0x90(%r8), %rax
  3191: 428a1418                 movb (%rax, %r11), %dl
  3195: 32da                     xorb %dl, %bl
  3197: 498b80e8000000           movq 0xe8(%r8), %rax
  3204: 428a1418                 movb (%rax, %r11), %dl
  3208: 32da                     xorb %dl, %bl
  3210: 498b8040010000           movq 0x140(%r8), %rax
  3217: 428a1418                 movb (%rax, %r11), %dl
  3221: 32da                     xorb %dl, %bl
  3223: 498b8098010000           movq 0x198(%r8), %rax
  3230: 428a1418                 movb (%rax, %r11), %dl
  3234: 32da                     xorb %dl, %bl
  3236: 498b80f0010000           movq 0x1f0(%r8), %rax
  3243: 428a1418                 movb (%rax, %r11), %dl
  3247: 32da                     xorb %dl, %bl
  3249: 498b8048020000           movq 0x248(%r8), %rax
  3256: 428a1418                 movb (%rax, %r11), %dl
  3260: 32da                     xorb %dl, %bl
  3262: 498b80a0020000           movq 0x2a0(%r8), %rax
  3269: 428a1418                 movb (%rax, %r11), %dl
  3273: 32da                     xorb %dl, %bl
  3275: 43881c1a                 movb %bl, (%r10, %r11)
  3279: 33db                     xorl %ebx, %ebx
  3281: 49ffc3                   incq %r11
  3284: 4c3bd9                   cmpq %rcx, %r11
  3287: 0f8289ffffff             jb 0xc66
  3293: 4d8b5140                 movq 0x40(%r9), %r10
  3297: 4c8b1c24                 movq (%rsp), %r11
  3301: 498b4040                 movq 0x40(%r8), %rax
  3305: 428a1418                 movb (%rax, %r11), %dl
  3309: 8ada                     movb %dl, %bl
  3311: 498b8098000000           movq 0x98(%r8), %rax
  3318: 428a1418                 movb (%rax, %r11), %dl
  3322: 32da                     xorb %dl, %bl
  3324: 498b80f0000000           movq 0xf0(%r8), %rax
  3331: 428a1418                 movb (%rax, %r11), %dl
  3335: 32da                     xorb %dl, %bl
  3337: 498b8048010000           movq 0x148(%r8), %rax
  3344: 428a1418                 movb (%rax, %r11), %dl
  3348: 32da                     xorb %dl, %bl
  3350: 498b80a0010000           movq 0x1a0(%r8), %rax
  3357: 428a1418                 movb (%rax, %r11), %dl
  3361: 32da                     xorb %dl, %bl
  3363: 498b80f8010000           movq 0x1f8(%r8), %rax
  3370: 428a1418                 movb (%rax, %r11), %dl
  3374: 32da                     xorb %dl, %bl
  3376: 498b8050020000           movq 0x250(%r8), %rax
  3383: 428a1418                 movb (%rax, %r11), %dl
  3387: 32da                     xorb %dl, %bl
  3389: 498b80a8020000           movq 0x2a8(%r8), %rax
  3396: 428a1418                 movb (%rax, %r11), %dl
  3400: 32da                     xorb %dl, %bl
  3402: 43881c1a                 movb %bl, (%r10, %r11)
  3406: 33db                     xorl %ebx, %ebx
  3408: 49ffc3                   incq %r11
  3411: 4c3bd9                   cmpq %rcx, %r11
  3414: 0f8289ffffff             jb 0xce5
  3420: 4d8b5148                 movq 0x48(%r9), %r10
  3424: 4c8b1c24                 movq (%rsp), %r11
  3428: 498b4048                 movq 0x48(%r8), %rax
  3432: 428a1418                 movb (%rax, %r11), %dl
  3436: 8ada                     movb %dl, %bl
  3438: 498b80a0000000           movq 0xa0(%r8), %rax
  3445: 428a1418                 movb (%rax, %r11), %dl
  3449: 32da                     xorb %dl, %bl
  3451: 498b80f8000000           movq 0xf8(%r8), %rax
  3458: 428a1418                 movb (%rax, %r11), %dl
  3462: 32da                     xorb %dl, %bl
  3464: 498b8050010000           movq 0x150(%r8), %rax
  3471: 428a1418                 movb (%rax, %r11), %dl
  3475: 32da                     xorb %dl, %bl
  3477: 498b80a8010000           movq 0x1a8(%r8), %rax
  3484: 428a1418                 movb (%rax, %r11), %dl
  3488: 32da                     xorb %dl, %bl
  3490: 498b8000020000           movq 0x200(%r8), %rax
  3497: 428a1418                 movb (%rax, %r11), %dl
  3501: 32da                     xorb %dl, %bl
  3503: 498b8058020000           movq 0x258(%r8), %rax
  3510: 428a1418                 movb (%rax, %r11), %dl
  3514: 32da                     xorb %dl, %bl
  3516: 498b80b0020000           movq 0x2b0(%r8), %rax
  3523: 428a1418                 movb (%rax, %r11), %dl
  3527: 32da                     xorb %dl, %bl
  3529: 43881c1a                 movb %bl, (%r10, %r11)
  3533: 33db                     xorl %ebx, %ebx
  3535: 49ffc3                   incq %r11
  3538: 4c3bd9                   cmpq %rcx, %r11
  3541: 0f8289ffffff             jb 0xd64
  3547: 4d8b5150                 movq 0x50(%r9), %r10
  3551: 4c8b1c24                 movq (%rsp), %r11
  3555: 498b4050                 movq 0x50(%r8), %rax
  3559: 428a1418                 movb (%rax, %r11), %dl
  3563: 8ada                     movb %dl, %bl
  3565: 498b80a8000000           movq 0xa8(%r8), %rax
  3572: 428a1418                 movb (%rax, %r11), %dl
  3576: 32da                     xorb %dl, %bl
  3578: 498b8000010000           movq 0x100(%r8), %rax
  3585: 428a1418                 movb (%rax, %r11), %dl
  3589: 32da                     xorb %dl, %bl
  3591: 498b8058010000           movq 0x158(%r8), %rax
  3598: 428a1418                 movb (%rax, %r11), %dl
  3602: 32da                     xorb %dl, %bl
  3604: 498b80b0010000           movq 0x1b0(%r8), %rax
  3611: 428a1418                 movb (%rax, %r11), %dl
  3615: 32da                     xorb %dl, %bl
  3617: 498b8008020000           movq 0x208(%r8), %rax
  3624: 428a1418                 movb (%rax, %r11), %dl
  3628: 32da                     xorb %dl, %bl
  3630: 498b8060020000           movq 0x260(%r8), %rax
  3637: 428a1418                 movb (%rax, %r11), %dl
  3641: 32da                     xorb %dl, %bl
  3643: 498b80b8020000           movq 0x2b8(%r8), %rax
  3650: 428a1418                 movb (%rax, %r11), %dl
  3654: 32da                     xorb %dl, %bl
  3656: 43881c1a                 movb %bl, (%r10, %r11)
  3660: 33db                     xorl %ebx, %ebx
  3662: 49ffc3                   incq %r11
  3665: 4c3bd9                   cmpq %rcx, %r11
  3668: 0f8289ffffff             jb 0xde3
  3674: 4d8b5158                 movq 0x58(%r9), %r10
  3678: 4c8b1c24                 movq (%rsp), %r11
  3682: 498b00                   movq (%r8), %rax
  3685: 428a1418                 movb (%rax, %r11), %dl
  3689: 8ada                     movb %dl, %bl
  3691: 498b4060                 movq 0x60(%r8), %rax
  3695: 428a1418                 movb (%rax, %r11), %dl
  3699: 32da                     xorb %dl, %bl
  3701: 498b80c0000000           movq 0xc0(%r8), %rax
  3708: 428a1418                 movb (%rax, %r11), %dl
  3712: 32da                     xorb %dl, %bl
  3714: 498b8020010000           movq 0x120(%r8), %rax
  3721: 428a1418                 movb (%rax, %r11), %dl
  3725: 32da                     xorb %dl, %bl
  3727: 498b8080010000           movq 0x180(%r8), %rax
  3734: 428a1418                 movb (%rax, %r11), %dl
  3738: 32da                     xorb %dl, %bl
  3740: 498b80e0010000           movq 0x1e0(%r8), %rax
  3747: 428a1418                 movb (%rax, %r11), %dl
  3751: 32da                     xorb %dl, %bl
  3753: 498b8040020000           movq 0x240(%r8), %rax
  3760: 428a1418                 movb (%rax, %r11), %dl
  3764: 32da                     xorb %dl, %bl
  3766: 498b80a0020000           movq 0x2a0(%r8), %rax
  3773: 428a1418                 movb (%rax, %r11), %dl
  3777: 32da                     xorb %dl, %bl
  3779: 43881c1a                 movb %bl, (%r10, %r11)
  3783: 33db                     xorl %ebx, %ebx
  3785: 49ffc3                   incq %r11
  3788: 4c3bd9                   cmpq %rcx, %r11
  3791: 0f828dffffff             jb 0xe62
  3797: 4d8b5160                 movq 0x60(%r9), %r10
  3801: 4c8b1c24                 movq (%rsp), %r11
  3805: 498b4008                 movq 8(%r8), %rax
  3809: 428a1418                 movb (%rax, %r11), %dl
  3813: 8ada                     movb %dl, %bl
  3815: 498b4068                 movq 0x68(%r8), %rax
  3819: 428a1418                 movb (%rax, %r11), %dl
  3823: 32da                     xorb %dl, %bl
  3825: 498b80c8000000           movq 0xc8(%r8), %rax
  3832: 428a1418                 movb (%rax, %r11), %dl
  3836: 32da                     xorb %dl, %bl
  3838: 498b8028010000           movq 0x128(%r8), %rax
  3845: 428a1418                 movb (%rax, %r11), %dl
  3849: 32da                     xorb %dl, %bl
  3851: 498b8088010000           movq 0x188(%r8), %rax
  3858: 428a1418                 movb (%rax, %r11), %dl
  3862: 32da                     xorb %dl, %bl
  3864: 498b80e8010000           movq 0x1e8(%r8), %rax
  3871: 428a1418                 movb (%rax, %r11), %dl
  3875: 32da                     xorb %dl, %bl
  3877: 498b8048020000           movq 0x248(%r8), %rax
  3884: 428a1418                 movb (%rax, %r11), %dl
  3888: 32da                     xorb %dl, %bl
  3890: 498b80a8020000           movq 0x2a8(%r8), %rax
  3897: 428a1418                 movb (%rax, %r11), %dl
  3901: 32da                     xorb %dl, %bl
  3903: 43881c1a                 movb %bl, (%r10, %r11)
  3907: 33db                     xorl %ebx, %ebx
  3909: 49ffc3                   incq %r11
  3912: 4c3bd9                   cmpq %rcx, %r11
  3915: 0f828cffffff             jb 0xedd
  3921: 4d8b5168                 movq 0x68(%r9), %r10
  3925: 4c8b1c24                 movq (%rsp), %r11
  3929: 498b4010                 movq 0x10(%r8), %rax
  3933: 428a1418                 movb (%rax, %r11), %dl
  3937: 8ada                     movb %dl, %bl
  3939: 498b4070                 movq 0x70(%r8), %rax
  3943: 428a1418                 movb (%rax, %r11), %dl
  3947: 32da                     xorb %dl, %bl
  3949: 498b80d0000000           movq 0xd0(%r8), %rax
  3956: 428a1418                 movb (%rax, %r11), %dl
  3960: 32da                     xorb %dl, %bl
  3962: 498b8030010000           movq 0x130(%r8), %rax
  3969: 428a1418                 movb (%rax, %r11), %dl
  3973: 32da                     xorb %dl, %bl
  3975: 498b8090010000           movq 0x190(%r8), %rax
  3982: 428a1418                 movb (%rax, %r11), %dl
  3986: 32da                     xorb %dl, %bl
  3988: 498b80f0010000           movq 0x1f0(%r8), %rax
  3995: 428a1418                 movb (%rax, %r11), %dl
  3999: 32da                     xorb %dl, %bl
  4001: 498b8050020000           movq 0x250(%r8), %rax
  4008: 428a1418                 movb (%rax, %r11), %dl
  4012: 32da                     xorb %dl, %bl
  4014: 498b80a8020000           movq 0x2a8(%r8), %rax
  4021: 428a1418                 movb (%rax, %r11), %dl
  4025: 32da                     xorb %dl, %bl
  4027: 498b80b0020000           movq 0x2b0(%r8), %rax
  4034: 428a1418                 movb (%rax, %r11), %dl
  4038: 32da                     xorb %dl, %bl
  4040: 43881c1a                 movb %bl, (%r10, %r11)
  4044: 33db                     xorl %ebx, %ebx
  4046: 49ffc3                   incq %r11
  4049: 4c3bd9                   cmpq %rcx, %r11
  4052: 0f827fffffff             jb 0xf59
  4058: 4d8b5170                 movq 0x70(%r9), %r10
  4062: 4c8b1c24                 movq (%rsp), %r11
  4066: 498b4018                 movq 0x18(%r8), %rax
  4070: 428a1418                 movb (%rax, %r11), %dl
  4074: 8ada                     movb %dl, %bl
  4076: 498b4078                 movq 0x78(%r8), %rax
  4080: 428a1418                 movb (%rax, %r11), %dl
  4084: 32da                     xorb %dl, %bl
  4086: 498b80d8000000           movq 0xd8(%r8), %rax
  4093: 428a1418                 movb (%rax, %r11), %dl
  4097: 32da                     xorb %dl, %bl
  4099: 498b8038010000           movq 0x138(%r8), %rax
  4106: 428a1418                 movb (%rax, %r11), %dl
  4110: 32da                     xorb %dl, %bl
  4112: 498b8098010000           movq 0x198(%r8), %rax
  4119: 428a1418                 movb (%rax, %r11), %dl
  4123: 32da                     xorb %dl, %bl
  4125: 498b80f0010000           movq 0x1f0(%r8), %rax
  4132: 428a1418                 movb (%rax, %r11), %dl
  4136: 32da                     xorb %dl, %bl
  4138: 498b80f8010000           movq 0x1f8(%r8), %rax
  4145: 428a1418                 movb (%rax, %r11), %dl
  4149: 32da                     xorb %dl, %bl
  4151: 498b8058020000           movq 0x258(%r8), %rax
  4158: 428a1418                 movb (%rax, %r11), %dl
  4162: 32da                     xorb %dl, %bl
  4164: 498b80b8020000           movq 0x2b8(%r8), %rax
  4171: 428a1418                 movb (%rax, %r11), %dl
  4175: 32da                     xorb %dl, %bl
  4177: 43881c1a                 movb %bl, (%r10, %r11)
  4181: 33db                     xorl %ebx, %ebx
  4183: 49ffc3                   incq %r11
  4186: 4c3bd9                   cmpq %rcx, %r11
  4189: 0f827fffffff             jb 0xfe2
  4195: 4d8b5178                 movq 0x78(%r9), %r10
  4199: 4c8b1c24                 movq (%rsp), %r11
  4203: 498b4020                 movq 0x20(%r8), %rax
  4207: 428a1418                 movb (%rax, %r11), %dl
  4211: 8ada                     movb %dl, %bl
  4213: 498b8080000000           movq 0x80(%r8), %rax
  4220: 428a1418                 movb (%rax, %r11), %dl
  4224: 32da                     xorb %dl, %bl
  4226: 498b80e0000000           movq 0xe0(%r8), %rax
  4233: 428a1418                 movb (%rax, %r11), %dl
  4237: 32da                     xorb %dl, %bl
  4239: 498b8038010000           movq 0x138(%r8), %rax
  4246: 428a1418                 movb (%rax, %r11), %dl
  4250: 32da                     xorb %dl, %bl
  4252: 498b8040010000           movq 0x140(%r8), %rax
  4259: 428a1418                 movb (%rax, %r11), %dl
  4263: 32da                     xorb %dl, %bl
  4265: 498b80a0010000           movq 0x1a0(%r8), %rax
  4272: 428a1418                 movb (%rax, %r11), %dl
  4276: 32da                     xorb %dl, %bl
  4278: 498b8000020000           movq 0x200(%r8), %rax
  4285: 428a1418                 movb (%rax, %r11), %dl
  4289: 32da                     xorb %dl, %bl
  4291: 498b8060020000           movq 0x260(%r8), %rax
  4298: 428a1418                 movb (%rax, %r11), %dl
  4302: 32da                     xorb %dl, %bl
  4304: 498b8068020000           movq 0x268(%r8), %rax
  4311: 428a1418                 movb (%rax, %r11), %dl
  4315: 32da                     xorb %dl, %bl
  4317: 43881c1a                 movb %bl, (%r10, %r11)
  4321: 33db                     xorl %ebx, %ebx
  4323: 49ffc3                   incq %r11
  4326: 4c3bd9                   cmpq %rcx, %r11
  4329: 0f827cffffff             jb 0x106b
  4335: 4d8b9180000000           movq 0x80(%r9), %r10
  4342: 4c8b1c24                 movq (%rsp), %r11
  4346: 498b4028                 movq 0x28(%r8), %rax
  4350: 428a1418                 movb (%rax, %r11), %dl
  4354: 8ada                     movb %dl, %bl
  4356: 498b8080000000           movq 0x80(%r8), %rax
  4363: 428a1418                 movb (%rax, %r11), %dl
  4367: 32da                     xorb %dl, %bl
  4369: 498b8088000000           movq 0x88(%r8), %rax
  4376: 428a1418                 movb (%rax, %r11), %dl
  4380: 32da                     xorb %dl, %bl
  4382: 498b80e8000000           movq 0xe8(%r8), %rax
  4389: 428a1418                 movb (%rax, %r11), %dl
  4393: 32da                     xorb %dl, %bl
  4395: 498b8048010000           movq 0x148(%r8), %rax
  4402: 428a1418                 movb (%rax, %r11), %dl
  4406: 32da                     xorb %dl, %bl
  4408: 498b80a8010000           movq 0x1a8(%r8), %rax
  4415: 428a1418                 movb (%rax, %r11), %dl
  4419: 32da                     xorb %dl, %bl
  4421: 498b8008020000           movq 0x208(%r8), %rax
  4428: 428a1418                 movb (%rax, %r11), %dl
  4432: 32da                     xorb %dl, %bl
  4434: 498b8010020000           movq 0x210(%r8), %rax
  4441: 428a1418                 movb (%rax, %r11), %dl
  4445: 32da                     xorb %dl, %bl
  4447: 498b8070020000           movq 0x270(%r8), %rax
  4454: 428a1418                 movb (%rax, %r11), %dl
  4458: 32da                     xorb %dl, %bl
  4460: 43881c1a                 movb %bl, (%r10, %r11)
  4464: 33db                     xorl %ebx, %ebx
  4466: 49ffc3                   incq %r11
  4469: 4c3bd9                   cmpq %rcx, %r11
  4472: 0f827cffffff             jb 0x10fa
  4478: 4d8b9188000000           movq 0x88(%r9), %r10
  4485: 4c8b1c24                 movq (%rsp), %r11
  4489: 498b4030                 movq 0x30(%r8), %rax
  4493: 428a1418                 movb (%rax, %r11), %dl
  4497: 8ada                     movb %dl, %bl
  4499: 498b8090000000           movq 0x90(%r8), %rax
  4506: 428a1418                 movb (%rax, %r11), %dl
  4510: 32da                     xorb %dl, %bl
  4512: 498b80f0000000           movq 0xf0(%r8), %rax
  4519: 428a1418                 movb (%rax, %r11), %dl
  4523: 32da                     xorb %dl, %bl
  4525: 498b8050010000           movq 0x150(%r8), %rax
  4532: 428a1418                 movb (%rax, %r11), %dl
  4536: 32da                     xorb %dl, %bl
  4538: 498b80b0010000           movq 0x1b0(%r8), %rax
  4545: 428a1418                 movb (%rax, %r11), %dl
  4549: 32da                     xorb %dl, %bl
  4551: 498b80b8010000           movq 0x1b8(%r8), %rax
  4558: 428a1418                 movb (%rax, %r11), %dl
  4562: 32da                     xorb %dl, %bl
  4564: 498b8018020000           movq 0x218(%r8), %rax
  4571: 428a1418                 movb (%rax, %r11), %dl
  4575: 32da                     xorb %dl, %bl
  4577: 498b8078020000           movq 0x278(%r8), %rax
  4584: 428a1418                 movb (%rax, %r11), %dl
  4588: 32da                     xorb %dl, %bl
  4590: 43881c1a                 movb %bl, (%r10, %r11)
  4594: 33db                     xorl %ebx, %ebx
  4596: 49ffc3                   incq %r11
  4599: 4c3bd9                   cmpq %rcx, %r11
  4602: 0f8289ffffff             jb 0x1189
  4608: 4d8b9190000000           movq 0x90(%r9), %r10
  4615: 4c8b1c24                 movq (%rsp), %r11
  4619: 498b4038                 movq 0x38(%r8), %rax
  4623: 428a1418                 movb (%rax, %r11), %dl
  4627: 8ada                     movb %dl, %bl
  4629: 498b8098000000           movq 0x98(%r8), %rax
  4636: 428a1418                 movb (%rax, %r11), %dl
  4640: 32da                     xorb %dl, %bl
  4642: 498b80f8000000           movq 0xf8(%r8), %rax
  4649: 428a1418                 movb (%rax, %r11), %dl
  4653: 32da                     xorb %dl, %bl
  4655: 498b8058010000           movq 0x158(%r8), %rax
  4662: 428a1418                 movb (%rax, %r11), %dl
  4666: 32da                     xorb %dl, %bl
  4668: 498b8060010000           movq 0x160(%r8), %rax
  4675: 428a1418                 movb (%rax, %r11), %dl
  4679: 32da                     xorb %dl, %bl
  4681: 498b80c0010000           movq 0x1c0(%r8), %rax
  4688: 428a1418                 movb (%rax, %r11), %dl
  4692: 32da                     xorb %dl, %bl
  4694: 498b8020020000           movq 0x220(%r8), %rax
  4701: 428a1418                 movb (%rax, %r11), %dl
  4705: 32da                     xorb %dl, %bl
  4707: 498b8080020000           movq 0x280(%r8), %rax
  4714: 428a1418                 movb (%rax, %r11), %dl
  4718: 32da                     xorb %dl, %bl
  4720: 43881c1a                 movb %bl, (%r10, %r11)
  4724: 33db                     xorl %ebx, %ebx
  4726: 49ffc3                   incq %r11
  4729: 4c3bd9                   cmpq %rcx, %r11
  4732: 0f8289ffffff             jb 0x120b
  4738: 4d8b9198000000           movq 0x98(%r9), %r10
  4745: 4c8b1c24                 movq (%rsp), %r11
  4749: 498b4040                 movq 0x40(%r8), %rax
  4753: 428a1418                 movb (%rax, %r11), %dl
  4757: 8ada                     movb %dl, %bl
  4759: 498b80a0000000           movq 0xa0(%r8), %rax
  4766: 428a1418                 movb (%rax, %r11), %dl
  4770: 32da                     xorb %dl, %bl
  4772: 498b8000010000           movq 0x100(%r8), %rax
  4779: 428a1418                 movb (%rax, %r11), %dl
  4783: 32da                     xorb %dl, %bl
  4785: 498b8008010000           movq 0x108(%r8), %rax
  4792: 428a1418                 movb (%rax, %r11), %dl
  4796: 32da                     xorb %dl, %bl
  4798: 498b8068010000           movq 0x168(%r8), %rax
  4805: 428a1418                 movb (%rax, %r11), %dl
  4809: 32da                     xorb %dl, %bl
  4811: 498b80c8010000           movq 0x1c8(%r8), %rax
  4818: 428a1418                 movb (%rax, %r11), %dl
  4822: 32da                     xorb %dl, %bl
  4824: 498b8020020000           movq 0x220(%r8), %rax
  4831: 428a1418                 movb (%rax, %r11), %dl
  4835: 32da                     xorb %dl, %bl
  4837: 498b8028020000           movq 0x228(%r8), %rax
  4844: 428a1418                 movb (%rax, %r11), %dl
  4848: 32da                     xorb %dl, %bl
  4850: 498b8088020000           movq 0x288(%r8), %rax
  4857: 428a1418                 movb (%rax, %r11), %dl
  4861: 32da                     xorb %dl, %bl
  4863: 43881c1a                 movb %bl, (%r10, %r11)
  4867: 33db                     xorl %ebx, %ebx
  4869: 49ffc3                   incq %r11
  4872: 4c3bd9                   cmpq %rcx, %r11
  4875: 0f827cffffff             jb 0x128d
  4881: 4d8b91a0000000           movq 0xa0(%r9), %r10
  4888: 4c8b1c24                 movq (%rsp), %r11
  4892: 498b4048                 movq 0x48(%r8), %rax
  4896: 428a1418                 movb (%rax, %r11), %dl
  4900: 8ada                     movb %dl, %bl
  4902: 498b80a8000000           movq 0xa8(%r8), %rax
  4909: 428a1418                 movb (%rax, %r11), %dl
  4913: 32da                     xorb %dl, %bl
  4915: 498b80b0000000           movq 0xb0(%r8), %rax
  4922: 428a1418                 movb (%rax, %r11), %dl
  4926: 32da                     xorb %dl, %bl
  4928: 498b8010010000           movq 0x110(%r8), %rax
  4935: 428a1418                 movb (%rax, %r11), %dl
  4939: 32da                     xorb %dl, %bl
  4941: 498b8068010000           movq 0x168(%r8), %rax
  4948: 428a1418                 movb (%rax, %r11), %dl
  4952: 32da                     xorb %dl, %bl
  4954: 498b8070010000           movq 0x170(%r8), %rax
  4961: 428a1418                 movb (%rax, %r11), %dl
  4965: 32da                     xorb %dl, %bl
  4967: 498b80d0010000           movq 0x1d0(%r8), %rax
  4974: 428a1418                 movb (%rax, %r11), %dl
  4978: 32da                     xorb %dl, %bl
  4980: 498b8030020000           movq 0x230(%r8), %rax
  4987: 428a1418                 movb (%rax, %r11), %dl
  4991: 32da                     xorb %dl, %bl
  4993: 498b8090020000           movq 0x290(%r8), %rax
  5000: 428a1418                 movb (%rax, %r11), %dl
  5004: 32da                     xorb %dl, %bl
  5006: 43881c1a                 movb %bl, (%r10, %r11)
  5010: 33db                     xorl %ebx, %ebx
  5012: 49ffc3                   incq %r11
  5015: 4c3bd9                   cmpq %rcx, %r11
  5018: 0f827cffffff             jb 0x131c
  5024: 4d8b91a8000000           movq 0xa8(%r9), %r10
  5031: 4c8b1c24                 movq (%rsp), %r11
  5035: 498b4050                 movq 0x50(%r8), %rax
  5039: 428a1418                 movb (%rax, %r11), %dl
  5043: 8ada                     movb %dl, %bl
  5045: 498b4058                 movq 0x58(%r8), %rax
  5049: 428a1418                 movb (%rax, %r11), %dl
  5053: 32da                     xorb %dl, %bl
  5055: 498b80b0000000           movq 0xb0(%r8), %rax
  5062: 428a1418                 movb (%rax, %r11), %dl
  5066: 32da                     xorb %dl, %bl
  5068: 498b80b8000000           movq 0xb8(%r8), %rax
  5075: 428a1418                 movb (%rax, %r11), %dl
  5079: 32da                     xorb %dl, %bl
  5081: 498b8018010000           movq 0x118(%r8), %rax
  5088: 428a1418                 movb (%rax, %r11), %dl
  5092: 32da                     xorb %dl, %bl
  5094: 498b8078010000           movq 0x178(%r8), %rax
  5101: 428a1418                 movb (%rax, %r11), %dl
  5105: 32da                     xorb %dl, %bl
  5107: 498b80d8010000           movq 0x1d8(%r8), %rax
  5114: 428a1418                 movb (%rax, %r11), %dl
  5118: 32da                     xorb %dl, %bl
  5120: 498b8038020000           movq 0x238(%r8), %rax
  5127: 428a1418                 movb (%rax, %r11), %dl
  5131: 32da                     xorb %dl, %bl
  5133: 498b8098020000           movq 0x298(%r8), %rax
  5140: 428a1418                 movb (%rax, %r11), %dl
  5144: 32da                     xorb %dl, %bl
  5146: 43881c1a                 movb %bl, (%r10, %r11)
  5150: 33db                     xorl %ebx, %ebx
  5152: 49ffc3                   incq %r11
  5155: 4c3bd9                   cmpq %rcx, %r11
  5158: 0f827fffffff             jb 0x13ab
  5164: 4883c408                 addq $8, %rsp
  5168: 415e                     popq %r14
  5170: 415d                     popq %r13
  5172: 415c                     popq %r12
  5174: 5b                       popq %rbx
  5175: c5f877                   vzeroupper 
  5178: c3                       retq 
  5179: 0f0b                     ud2 
  5181: 0f0b                     ud2 
  5183: cc                       int3 
