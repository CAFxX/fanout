; config_t1_56 target=avx512-gfni twin=notail topo=rdp(7)/enc op=encode U=1
; code_size=1280 twins=["notail"] uops_per_byte=2.75
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
    45: 0f8298040000             jb 0x4cb
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
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
  1227: 415e                     popq %r14
  1229: 415d                     popq %r13
  1231: 415c                     popq %r12
  1233: 5b                       popq %rbx
  1234: c5f877                   vzeroupper 
  1237: c3                       retq 
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
