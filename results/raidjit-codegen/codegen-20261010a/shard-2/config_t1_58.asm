; config_t1_58 target=avx512-gfni twin=maskedtail topo=rdp(7)/enc op=encode U=1
; code_size=2560 twins=["maskedtail"] uops_per_byte=2.75
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de7090000           leaq 0x9e7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4883e840                 subq $0x40, %rax
    54: 0f828f040000             jb 0x4cb
    60: 0f1f4000                 nopl (%rax)
    64: 62317d486f241e           vmovdqa32 (%rsi, %r11), %zmm12
    71: 62d17d486fc4             vmovdqa32 %zmm12, %zmm0
    77: 62d17d486ff4             vmovdqa32 %zmm12, %zmm6
    83: 62317d486f241f           vmovdqa32 (%rdi, %r11), %zmm12
    90: 62d17d486fcc             vmovdqa32 %zmm12, %zmm1
    96: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   102: 62d17d486ffc             vmovdqa32 %zmm12, %zmm7
   108: 62117d486f241c           vmovdqa32 (%r12, %r11), %zmm12
   115: 62d17d486fd4             vmovdqa32 %zmm12, %zmm2
   121: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   127: 62517d486fc4             vmovdqa32 %zmm12, %zmm8
   133: 62117d486f641d00         vmovdqa32 (%r13, %r11), %zmm12
   141: 62d17d486fdc             vmovdqa32 %zmm12, %zmm3
   147: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   153: 62517d486fcc             vmovdqa32 %zmm12, %zmm9
   159: 62117d486f241e           vmovdqa32 (%r14, %r11), %zmm12
   166: 62d17d486fe4             vmovdqa32 %zmm12, %zmm4
   172: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   178: 62517d486fd4             vmovdqa32 %zmm12, %zmm10
   184: 4d8b5028                 movq 0x28(%r8), %r10
   188: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   195: 62d17d486fec             vmovdqa32 %zmm12, %zmm5
   201: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   207: 62517d486fdc             vmovdqa32 %zmm12, %zmm11
   213: 4d8b5030                 movq 0x30(%r8), %r10
   217: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   224: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   230: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   236: 4d8b5038                 movq 0x38(%r8), %r10
   240: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   247: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   253: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   259: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   265: 4d8b5040                 movq 0x40(%r8), %r10
   269: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   276: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   282: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   288: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   294: 4d8b5048                 movq 0x48(%r8), %r10
   298: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   305: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   311: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   317: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   323: 4d8b5050                 movq 0x50(%r8), %r10
   327: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   334: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
   340: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   346: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   352: 4d8b5058                 movq 0x58(%r8), %r10
   356: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   363: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
   369: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   375: 4d8b5060                 movq 0x60(%r8), %r10
   379: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   386: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   392: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   398: 4d8b5068                 movq 0x68(%r8), %r10
   402: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   409: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   415: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   421: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   427: 4d8b5070                 movq 0x70(%r8), %r10
   431: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   438: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   444: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   450: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   456: 4d8b5078                 movq 0x78(%r8), %r10
   460: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   467: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   473: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   479: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   485: 4d8b9080000000           movq 0x80(%r8), %r10
   492: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   499: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
   505: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   511: 4d8b9088000000           movq 0x88(%r8), %r10
   518: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   525: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
   531: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   537: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   543: 4d8b9090000000           movq 0x90(%r8), %r10
   550: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   557: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   563: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   569: 4d8b9098000000           movq 0x98(%r8), %r10
   576: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   583: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   589: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   595: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   601: 4d8b90a0000000           movq 0xa0(%r8), %r10
   608: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   615: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   621: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   627: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   633: 4d8b90a8000000           movq 0xa8(%r8), %r10
   640: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   647: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   653: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   659: 4d8b90b0000000           movq 0xb0(%r8), %r10
   666: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   673: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
   679: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   685: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   691: 4d8b90b8000000           movq 0xb8(%r8), %r10
   698: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   705: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
   711: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   717: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   723: 4d8b90c0000000           movq 0xc0(%r8), %r10
   730: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   737: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   743: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   749: 4d8b90c8000000           movq 0xc8(%r8), %r10
   756: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   763: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   769: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   775: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   781: 4d8b90d0000000           movq 0xd0(%r8), %r10
   788: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   795: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   801: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   807: 4d8b90d8000000           movq 0xd8(%r8), %r10
   814: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   821: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
   827: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   833: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   839: 4d8b90e0000000           movq 0xe0(%r8), %r10
   846: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   853: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
   859: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   865: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
   871: 4d8b90e8000000           movq 0xe8(%r8), %r10
   878: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   885: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
   891: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
   897: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
   903: 4d8b90f0000000           movq 0xf0(%r8), %r10
   910: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   917: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
   923: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
   929: 4d8b90f8000000           movq 0xf8(%r8), %r10
   936: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   943: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
   949: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   955: 4d8b9000010000           movq 0x100(%r8), %r10
   962: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
   969: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
   975: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
   981: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
   987: 4d8b9008010000           movq 0x108(%r8), %r10
   994: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
  1001: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
  1007: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  1013: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1019: 4d8b9010010000           movq 0x110(%r8), %r10
  1026: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
  1033: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
  1039: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1045: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1051: 4d8b9018010000           movq 0x118(%r8), %r10
  1058: 62117d486f241a           vmovdqa32 (%r10, %r11), %zmm12
  1065: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
  1071: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1077: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1083: 4d8b11                   movq (%r9), %r10
  1086: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  1093: 4d8b5108                 movq 8(%r9), %r10
  1097: 62917d487f0c1a           vmovdqa32 %zmm1, (%r10, %r11)
  1104: 4d8b5110                 movq 0x10(%r9), %r10
  1108: 62917d487f141a           vmovdqa32 %zmm2, (%r10, %r11)
  1115: 4d8b5118                 movq 0x18(%r9), %r10
  1119: 62917d487f1c1a           vmovdqa32 %zmm3, (%r10, %r11)
  1126: 4d8b5120                 movq 0x20(%r9), %r10
  1130: 62917d487f241a           vmovdqa32 %zmm4, (%r10, %r11)
  1137: 4d8b5128                 movq 0x28(%r9), %r10
  1141: 62917d487f2c1a           vmovdqa32 %zmm5, (%r10, %r11)
  1148: 4d8b5130                 movq 0x30(%r9), %r10
  1152: 62917d487f341a           vmovdqa32 %zmm6, (%r10, %r11)
  1159: 4d8b5138                 movq 0x38(%r9), %r10
  1163: 62917d487f3c1a           vmovdqa32 %zmm7, (%r10, %r11)
  1170: 4d8b5140                 movq 0x40(%r9), %r10
  1174: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  1181: 4d8b5148                 movq 0x48(%r9), %r10
  1185: 62117d487f0c1a           vmovdqa32 %zmm9, (%r10, %r11)
  1192: 4d8b5150                 movq 0x50(%r9), %r10
  1196: 62117d487f141a           vmovdqa32 %zmm10, (%r10, %r11)
  1203: 4d8b5158                 movq 0x58(%r9), %r10
  1207: 62117d487f1c1a           vmovdqa32 %zmm11, (%r10, %r11)
  1214: 4983c340                 addq $0x40, %r11
  1218: 4c3bd8                   cmpq %rax, %r11
  1221: 0f8675fbffff             jbe 0x40
  1227: 4c3bd9                   cmpq %rcx, %r11
  1230: 0f83f2040000             jae 0x9c6
  1236: 4c8bd1                   movq %rcx, %r10
  1239: 4d2bd3                   subq %r11, %r10
  1242: 498bd2                   movq %r10, %rdx
  1245: 4883fa40                 cmpq $0x40, %rdx
  1249: 0f8607000000             jbe 0x4ee
  1255: 48c7c240000000           movq $0x40, %rdx
  1262: c4c1f8900cd7             kmovq (%r15, %rdx, 8), %k1
  1268: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
  1274: 62f17548efc9             vpxord %zmm1, %zmm1, %zmm1
  1280: 62f16d48efd2             vpxord %zmm2, %zmm2, %zmm2
  1286: 62f16548efdb             vpxord %zmm3, %zmm3, %zmm3
  1292: 62f15d48efe4             vpxord %zmm4, %zmm4, %zmm4
  1298: 62f15548efed             vpxord %zmm5, %zmm5, %zmm5
  1304: 62f14d48eff6             vpxord %zmm6, %zmm6, %zmm6
  1310: 62f14548efff             vpxord %zmm7, %zmm7, %zmm7
  1316: 62513d48efc0             vpxord %zmm8, %zmm8, %zmm8
  1322: 62513548efc9             vpxord %zmm9, %zmm9, %zmm9
  1328: 62512d48efd2             vpxord %zmm10, %zmm10, %zmm10
  1334: 62512548efdb             vpxord %zmm11, %zmm11, %zmm11
  1340: 62317fc96f241e           vmovdqu8 (%rsi, %r11), %zmm12 {%k1} {z}
  1347: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
  1353: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  1359: 62317fc96f241f           vmovdqu8 (%rdi, %r11), %zmm12 {%k1} {z}
  1366: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
  1372: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  1378: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  1384: 62117fc96f241c           vmovdqu8 (%r12, %r11), %zmm12 {%k1} {z}
  1391: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
  1397: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  1403: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1409: 62117fc96f641d00         vmovdqu8 (%r13, %r11), %zmm12 {%k1} {z}
  1417: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
  1423: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1429: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1435: 62117fc96f241e           vmovdqu8 (%r14, %r11), %zmm12 {%k1} {z}
  1442: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
  1448: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1454: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1460: 498b4028                 movq 0x28(%r8), %rax
  1464: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1471: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
  1477: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1483: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
  1489: 498b4030                 movq 0x30(%r8), %rax
  1493: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1500: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
  1506: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  1512: 498b4038                 movq 0x38(%r8), %rax
  1516: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1523: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
  1529: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  1535: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1541: 498b4040                 movq 0x40(%r8), %rax
  1545: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1552: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
  1558: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  1564: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1570: 498b4048                 movq 0x48(%r8), %rax
  1574: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1581: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
  1587: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1593: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1599: 498b4050                 movq 0x50(%r8), %rax
  1603: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1610: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
  1616: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1622: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
  1628: 498b4058                 movq 0x58(%r8), %rax
  1632: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1639: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
  1645: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1651: 498b4060                 movq 0x60(%r8), %rax
  1655: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1662: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
  1668: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1674: 498b4068                 movq 0x68(%r8), %rax
  1678: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1685: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
  1691: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  1697: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1703: 498b4070                 movq 0x70(%r8), %rax
  1707: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1714: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
  1720: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  1726: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1732: 498b4078                 movq 0x78(%r8), %rax
  1736: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1743: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
  1749: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1755: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
  1761: 498b8080000000           movq 0x80(%r8), %rax
  1768: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1775: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
  1781: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1787: 498b8088000000           movq 0x88(%r8), %rax
  1794: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1801: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
  1807: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  1813: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1819: 498b8090000000           movq 0x90(%r8), %rax
  1826: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1833: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
  1839: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1845: 498b8098000000           movq 0x98(%r8), %rax
  1852: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1859: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
  1865: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  1871: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1877: 498b80a0000000           movq 0xa0(%r8), %rax
  1884: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1891: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
  1897: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  1903: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
  1909: 498b80a8000000           movq 0xa8(%r8), %rax
  1916: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1923: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
  1929: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  1935: 498b80b0000000           movq 0xb0(%r8), %rax
  1942: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1949: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
  1955: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  1961: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  1967: 498b80b8000000           movq 0xb8(%r8), %rax
  1974: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  1981: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
  1987: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  1993: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  1999: 498b80c0000000           movq 0xc0(%r8), %rax
  2006: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2013: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
  2019: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  2025: 498b80c8000000           movq 0xc8(%r8), %rax
  2032: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2039: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
  2045: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  2051: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
  2057: 498b80d0000000           movq 0xd0(%r8), %rax
  2064: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2071: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
  2077: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  2083: 498b80d8000000           movq 0xd8(%r8), %rax
  2090: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2097: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
  2103: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  2109: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  2115: 498b80e0000000           movq 0xe0(%r8), %rax
  2122: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2129: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
  2135: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  2141: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  2147: 498b80e8000000           movq 0xe8(%r8), %rax
  2154: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2161: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
  2167: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  2173: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  2179: 498b80f0000000           movq 0xf0(%r8), %rax
  2186: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2193: 62d17d48efc4             vpxord %zmm12, %zmm0, %zmm0
  2199: 62512548efdc             vpxord %zmm12, %zmm11, %zmm11
  2205: 498b80f8000000           movq 0xf8(%r8), %rax
  2212: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2219: 62d17548efcc             vpxord %zmm12, %zmm1, %zmm1
  2225: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  2231: 498b8000010000           movq 0x100(%r8), %rax
  2238: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2245: 62d16d48efd4             vpxord %zmm12, %zmm2, %zmm2
  2251: 62d14d48eff4             vpxord %zmm12, %zmm6, %zmm6
  2257: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  2263: 498b8008010000           movq 0x108(%r8), %rax
  2270: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2277: 62d16548efdc             vpxord %zmm12, %zmm3, %zmm3
  2283: 62d14548effc             vpxord %zmm12, %zmm7, %zmm7
  2289: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  2295: 498b8010010000           movq 0x110(%r8), %rax
  2302: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2309: 62d15d48efe4             vpxord %zmm12, %zmm4, %zmm4
  2315: 62513d48efc4             vpxord %zmm12, %zmm8, %zmm8
  2321: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  2327: 498b8018010000           movq 0x118(%r8), %rax
  2334: 62317fc96f2418           vmovdqu8 (%rax, %r11), %zmm12 {%k1} {z}
  2341: 62d15548efec             vpxord %zmm12, %zmm5, %zmm5
  2347: 62513548efcc             vpxord %zmm12, %zmm9, %zmm9
  2353: 62512d48efd4             vpxord %zmm12, %zmm10, %zmm10
  2359: 498b01                   movq (%r9), %rax
  2362: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
  2369: 498b4108                 movq 8(%r9), %rax
  2373: 62b17f497f0c18           vmovdqu8 %zmm1, (%rax, %r11) {%k1}
  2380: 498b4110                 movq 0x10(%r9), %rax
  2384: 62b17f497f1418           vmovdqu8 %zmm2, (%rax, %r11) {%k1}
  2391: 498b4118                 movq 0x18(%r9), %rax
  2395: 62b17f497f1c18           vmovdqu8 %zmm3, (%rax, %r11) {%k1}
  2402: 498b4120                 movq 0x20(%r9), %rax
  2406: 62b17f497f2418           vmovdqu8 %zmm4, (%rax, %r11) {%k1}
  2413: 498b4128                 movq 0x28(%r9), %rax
  2417: 62b17f497f2c18           vmovdqu8 %zmm5, (%rax, %r11) {%k1}
  2424: 498b4130                 movq 0x30(%r9), %rax
  2428: 62b17f497f3418           vmovdqu8 %zmm6, (%rax, %r11) {%k1}
  2435: 498b4138                 movq 0x38(%r9), %rax
  2439: 62b17f497f3c18           vmovdqu8 %zmm7, (%rax, %r11) {%k1}
  2446: 498b4140                 movq 0x40(%r9), %rax
  2450: 62317f497f0418           vmovdqu8 %zmm8, (%rax, %r11) {%k1}
  2457: 498b4148                 movq 0x48(%r9), %rax
  2461: 62317f497f0c18           vmovdqu8 %zmm9, (%rax, %r11) {%k1}
  2468: 498b4150                 movq 0x50(%r9), %rax
  2472: 62317f497f1418           vmovdqu8 %zmm10, (%rax, %r11) {%k1}
  2479: 498b4158                 movq 0x58(%r9), %rax
  2483: 62317f497f1c18           vmovdqu8 %zmm11, (%rax, %r11) {%k1}
  2490: 4c03da                   addq %rdx, %r11
  2493: 4c2bd2                   subq %rdx, %r10
  2496: 0f8514fbffff             jne 0x4da
  2502: 415e                     popq %r14
  2504: 415d                     popq %r13
  2506: 415c                     popq %r12
  2508: 415f                     popq %r15
  2510: 5b                       popq %rbx
  2511: c5f877                   vzeroupper 
  2514: c3                       retq 
  2515: 0f0b                     ud2 
  2517: 0f0b                     ud2 
  2519: 0f0b                     ud2 
  2521: 0f0b                     ud2 
  2523: 0f0b                     ud2 
  2525: 0f0b                     ud2 
  2527: 0f0b                     ud2 
  2529: 0f0b                     ud2 
  2531: 0f0b                     ud2 
  2533: 0f0b                     ud2 
  2535: 0f0b                     ud2 
  2537: 0f0b                     ud2 
  2539: 0f0b                     ud2 
  2541: 0f0b                     ud2 
  2543: 0f0b                     ud2 
  2545: 0f0b                     ud2 
  2547: 0f0b                     ud2 
  2549: 0f0b                     ud2 
  2551: 0f0b                     ud2 
  2553: 0f0b                     ud2 
  2555: 0f0b                     ud2 
  2557: 0f0b                     ud2 
  2559: cc                       int3 
