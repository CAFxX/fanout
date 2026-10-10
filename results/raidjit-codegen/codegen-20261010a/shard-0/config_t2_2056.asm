; config_t2_2056 target=avx512-gfni twin=ternlog topo=rs(6,3)/enc op=encode U=7
; code_size=3072 twins=["ternlog"] uops_per_byte=0.732142857142857
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de70b0000           leaq 0xbe7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e8c0010000           subq $0x1c0, %rax
    57: 0f82bc080000             jb 0x8fb
    63: 90                       nop 
    64: 62a17d486f2c1e           vmovdqa32 (%rsi, %r11), %zmm21
    71: 62b17d486fc5             vmovdqa32 %zmm21, %zmm0
    77: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
    83: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
    90: 62b17d486ffe             vmovdqa32 %zmm22, %zmm7
    96: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   103: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   110: 62317d486ff6             vmovdqa32 %zmm22, %zmm14
   116: 62a17d486f6c1e01         vmovdqa32 0x40(%rsi, %r11), %zmm21
   124: 62b17d486fcd             vmovdqa32 %zmm21, %zmm1
   130: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
   136: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   143: 62317d486fc6             vmovdqa32 %zmm22, %zmm8
   149: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   156: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   163: 62317d486ffe             vmovdqa32 %zmm22, %zmm15
   169: 62a17d486f6c1e02         vmovdqa32 0x80(%rsi, %r11), %zmm21
   177: 62b17d486fd5             vmovdqa32 %zmm21, %zmm2
   183: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
   189: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   196: 62317d486fce             vmovdqa32 %zmm22, %zmm9
   202: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   209: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   216: 62a17d486fc6             vmovdqa32 %zmm22, %zmm16
   222: 62a17d486f6c1e03         vmovdqa32 0xc0(%rsi, %r11), %zmm21
   230: 62b17d486fdd             vmovdqa32 %zmm21, %zmm3
   236: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
   242: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   249: 62317d486fd6             vmovdqa32 %zmm22, %zmm10
   255: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   262: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   269: 62a17d486fce             vmovdqa32 %zmm22, %zmm17
   275: 62a17d486f6c1e04         vmovdqa32 0x100(%rsi, %r11), %zmm21
   283: 62b17d486fe5             vmovdqa32 %zmm21, %zmm4
   289: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
   295: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   302: 62317d486fde             vmovdqa32 %zmm22, %zmm11
   308: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   315: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   322: 62a17d486fd6             vmovdqa32 %zmm22, %zmm18
   328: 62a17d486f6c1e05         vmovdqa32 0x140(%rsi, %r11), %zmm21
   336: 62b17d486fed             vmovdqa32 %zmm21, %zmm5
   342: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
   348: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   355: 62317d486fe6             vmovdqa32 %zmm22, %zmm12
   361: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   368: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   375: 62a17d486fde             vmovdqa32 %zmm22, %zmm19
   381: 62a17d486f6c1e06         vmovdqa32 0x180(%rsi, %r11), %zmm21
   389: 62b17d486ff5             vmovdqa32 %zmm21, %zmm6
   395: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
   401: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   408: 62317d486fee             vmovdqa32 %zmm22, %zmm13
   414: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   421: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   428: 62a17d486fe6             vmovdqa32 %zmm22, %zmm20
   434: 62a17d486f2c1f           vmovdqa32 (%rdi, %r11), %zmm21
   441: 62b17d48efc5             vpxord %zmm21, %zmm0, %zmm0
   447: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   454: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
   460: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
   467: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   474: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
   480: 62a17d486f6c1f01         vmovdqa32 0x40(%rdi, %r11), %zmm21
   488: 62b17548efcd             vpxord %zmm21, %zmm1, %zmm1
   494: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   501: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   508: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
   514: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
   521: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   528: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
   534: 62a17d486f6c1f02         vmovdqa32 0x80(%rdi, %r11), %zmm21
   542: 62b16d48efd5             vpxord %zmm21, %zmm2, %zmm2
   548: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   555: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   562: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
   568: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
   575: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   582: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
   588: 62a17d486f6c1f03         vmovdqa32 0xc0(%rdi, %r11), %zmm21
   596: 62b16548efdd             vpxord %zmm21, %zmm3, %zmm3
   602: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   609: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   616: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
   622: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
   629: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   636: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
   642: 62a17d486f6c1f04         vmovdqa32 0x100(%rdi, %r11), %zmm21
   650: 62b15d48efe5             vpxord %zmm21, %zmm4, %zmm4
   656: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   663: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   670: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
   676: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
   683: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   690: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
   696: 62a17d486f6c1f05         vmovdqa32 0x140(%rdi, %r11), %zmm21
   704: 62b15548efed             vpxord %zmm21, %zmm5, %zmm5
   710: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   717: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   724: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
   730: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
   737: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   744: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
   750: 62a17d486f6c1f06         vmovdqa32 0x180(%rdi, %r11), %zmm21
   758: 62b14d48eff5             vpxord %zmm21, %zmm6, %zmm6
   764: 62c17d486f7f01           vmovdqa32 0x40(%r15), %zmm23
   771: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   778: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
   784: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
   791: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   798: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
   804: 62817d486f2c1c           vmovdqa32 (%r12, %r11), %zmm21
   811: 62b17d48efc5             vpxord %zmm21, %zmm0, %zmm0
   817: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   824: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
   830: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
   837: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   844: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
   850: 62817d486f6c1c01         vmovdqa32 0x40(%r12, %r11), %zmm21
   858: 62b17548efcd             vpxord %zmm21, %zmm1, %zmm1
   864: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
   871: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   878: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
   884: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
   891: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   898: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
   904: 62817d486f6c1c02         vmovdqa32 0x80(%r12, %r11), %zmm21
   912: 62b16d48efd5             vpxord %zmm21, %zmm2, %zmm2
   918: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
   925: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   932: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
   938: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
   945: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   952: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
   958: 62817d486f6c1c03         vmovdqa32 0xc0(%r12, %r11), %zmm21
   966: 62b16548efdd             vpxord %zmm21, %zmm3, %zmm3
   972: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
   979: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
   986: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
   992: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
   999: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1006: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1012: 62817d486f6c1c04         vmovdqa32 0x100(%r12, %r11), %zmm21
  1020: 62b15d48efe5             vpxord %zmm21, %zmm4, %zmm4
  1026: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
  1033: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1040: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  1046: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
  1053: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1060: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1066: 62817d486f6c1c05         vmovdqa32 0x140(%r12, %r11), %zmm21
  1074: 62b15548efed             vpxord %zmm21, %zmm5, %zmm5
  1080: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
  1087: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1094: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1100: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
  1107: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1114: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1120: 62817d486f6c1c06         vmovdqa32 0x180(%r12, %r11), %zmm21
  1128: 62b14d48eff5             vpxord %zmm21, %zmm6, %zmm6
  1134: 62c17d486f7f02           vmovdqa32 0x80(%r15), %zmm23
  1141: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1148: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  1154: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
  1161: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1168: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1174: 62817d486f6c1d00         vmovdqa32 (%r13, %r11), %zmm21
  1182: 62b17d48efc5             vpxord %zmm21, %zmm0, %zmm0
  1188: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1195: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1202: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  1208: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
  1214: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1221: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  1227: 62817d486f6c1d01         vmovdqa32 0x40(%r13, %r11), %zmm21
  1235: 62b17548efcd             vpxord %zmm21, %zmm1, %zmm1
  1241: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1248: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1255: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  1261: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
  1267: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1274: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  1280: 62817d486f6c1d02         vmovdqa32 0x80(%r13, %r11), %zmm21
  1288: 62b16d48efd5             vpxord %zmm21, %zmm2, %zmm2
  1294: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1301: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1308: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  1314: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
  1320: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1327: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1333: 62817d486f6c1d03         vmovdqa32 0xc0(%r13, %r11), %zmm21
  1341: 62b16548efdd             vpxord %zmm21, %zmm3, %zmm3
  1347: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1354: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1361: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1367: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
  1373: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1380: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1386: 62817d486f6c1d04         vmovdqa32 0x100(%r13, %r11), %zmm21
  1394: 62b15d48efe5             vpxord %zmm21, %zmm4, %zmm4
  1400: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1407: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1414: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  1420: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
  1426: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1433: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1439: 62817d486f6c1d05         vmovdqa32 0x140(%r13, %r11), %zmm21
  1447: 62b15548efed             vpxord %zmm21, %zmm5, %zmm5
  1453: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1460: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1467: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1473: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
  1479: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1486: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1492: 62817d486f6c1d06         vmovdqa32 0x180(%r13, %r11), %zmm21
  1500: 62b14d48eff5             vpxord %zmm21, %zmm6, %zmm6
  1506: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1513: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1520: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  1526: 62c17d486f3f             vmovdqa32 (%r15), %zmm23
  1532: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1539: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1545: 62817d486f2c1e           vmovdqa32 (%r14, %r11), %zmm21
  1552: 62b17d48efc5             vpxord %zmm21, %zmm0, %zmm0
  1558: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
  1565: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1572: 62b14548effe             vpxord %zmm22, %zmm7, %zmm7
  1578: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1585: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1592: 62310d48eff6             vpxord %zmm22, %zmm14, %zmm14
  1598: 62817d486f6c1e01         vmovdqa32 0x40(%r14, %r11), %zmm21
  1606: 62b17548efcd             vpxord %zmm21, %zmm1, %zmm1
  1612: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
  1619: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1626: 62313d48efc6             vpxord %zmm22, %zmm8, %zmm8
  1632: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1639: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1646: 62310548effe             vpxord %zmm22, %zmm15, %zmm15
  1652: 62817d486f6c1e02         vmovdqa32 0x80(%r14, %r11), %zmm21
  1660: 62b16d48efd5             vpxord %zmm21, %zmm2, %zmm2
  1666: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
  1673: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1680: 62313548efce             vpxord %zmm22, %zmm9, %zmm9
  1686: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1693: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1700: 62a17d40efc6             vpxord %zmm22, %zmm16, %zmm16
  1706: 62817d486f6c1e03         vmovdqa32 0xc0(%r14, %r11), %zmm21
  1714: 62b16548efdd             vpxord %zmm21, %zmm3, %zmm3
  1720: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
  1727: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1734: 62312d48efd6             vpxord %zmm22, %zmm10, %zmm10
  1740: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1747: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1754: 62a17540efce             vpxord %zmm22, %zmm17, %zmm17
  1760: 62817d486f6c1e04         vmovdqa32 0x100(%r14, %r11), %zmm21
  1768: 62b15d48efe5             vpxord %zmm21, %zmm4, %zmm4
  1774: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
  1781: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1788: 62312548efde             vpxord %zmm22, %zmm11, %zmm11
  1794: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1801: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1808: 62a16d40efd6             vpxord %zmm22, %zmm18, %zmm18
  1814: 62817d486f6c1e05         vmovdqa32 0x140(%r14, %r11), %zmm21
  1822: 62b15548efed             vpxord %zmm21, %zmm5, %zmm5
  1828: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
  1835: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1842: 62311d48efe6             vpxord %zmm22, %zmm12, %zmm12
  1848: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1855: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1862: 62a16540efde             vpxord %zmm22, %zmm19, %zmm19
  1868: 62817d486f6c1e06         vmovdqa32 0x180(%r14, %r11), %zmm21
  1876: 62b14d48eff5             vpxord %zmm21, %zmm6, %zmm6
  1882: 62c17d486f7f03           vmovdqa32 0xc0(%r15), %zmm23
  1889: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1896: 62311548efee             vpxord %zmm22, %zmm13, %zmm13
  1902: 62c17d486f7f04           vmovdqa32 0x100(%r15), %zmm23
  1909: 62a3d540cef700           vgf2p8affineqb $0, %zmm23, %zmm21, %zmm22
  1916: 62a15d40efe6             vpxord %zmm22, %zmm20, %zmm20
  1922: 4d8b5028                 movq 0x28(%r8), %r10
  1926: 62817d486f2c1a           vmovdqa32 (%r10, %r11), %zmm21
  1933: 62b17d48efc5             vpxord %zmm21, %zmm0, %zmm0
  1939: 62b14548effd             vpxord %zmm21, %zmm7, %zmm7
  1945: 62310d48eff5             vpxord %zmm21, %zmm14, %zmm14
  1951: 62817d486f6c1a01         vmovdqa32 0x40(%r10, %r11), %zmm21
  1959: 62b17548efcd             vpxord %zmm21, %zmm1, %zmm1
  1965: 62313d48efc5             vpxord %zmm21, %zmm8, %zmm8
  1971: 62310548effd             vpxord %zmm21, %zmm15, %zmm15
  1977: 62817d486f6c1a02         vmovdqa32 0x80(%r10, %r11), %zmm21
  1985: 62b16d48efd5             vpxord %zmm21, %zmm2, %zmm2
  1991: 62313548efcd             vpxord %zmm21, %zmm9, %zmm9
  1997: 62a17d40efc5             vpxord %zmm21, %zmm16, %zmm16
  2003: 62817d486f6c1a03         vmovdqa32 0xc0(%r10, %r11), %zmm21
  2011: 62b16548efdd             vpxord %zmm21, %zmm3, %zmm3
  2017: 62312d48efd5             vpxord %zmm21, %zmm10, %zmm10
  2023: 62a17540efcd             vpxord %zmm21, %zmm17, %zmm17
  2029: 62817d486f6c1a04         vmovdqa32 0x100(%r10, %r11), %zmm21
  2037: 62b15d48efe5             vpxord %zmm21, %zmm4, %zmm4
  2043: 62312548efdd             vpxord %zmm21, %zmm11, %zmm11
  2049: 62a16d40efd5             vpxord %zmm21, %zmm18, %zmm18
  2055: 62817d486f6c1a05         vmovdqa32 0x140(%r10, %r11), %zmm21
  2063: 62b15548efed             vpxord %zmm21, %zmm5, %zmm5
  2069: 62311d48efe5             vpxord %zmm21, %zmm12, %zmm12
  2075: 62a16540efdd             vpxord %zmm21, %zmm19, %zmm19
  2081: 62817d486f6c1a06         vmovdqa32 0x180(%r10, %r11), %zmm21
  2089: 62b14d48eff5             vpxord %zmm21, %zmm6, %zmm6
  2095: 62311548efed             vpxord %zmm21, %zmm13, %zmm13
  2101: 62a15d40efe5             vpxord %zmm21, %zmm20, %zmm20
  2107: 4d8b11                   movq (%r9), %r10
  2110: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  2117: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  2125: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
  2133: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
  2141: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
  2149: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
  2157: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
  2165: 4d8b5108                 movq 8(%r9), %r10
  2169: 62917d487f3c1a           vmovdqa32 %zmm7, (%r10, %r11)
  2176: 62117d487f441a01         vmovdqa32 %zmm8, 0x40(%r10, %r11)
  2184: 62117d487f4c1a02         vmovdqa32 %zmm9, 0x80(%r10, %r11)
  2192: 62117d487f541a03         vmovdqa32 %zmm10, 0xc0(%r10, %r11)
  2200: 62117d487f5c1a04         vmovdqa32 %zmm11, 0x100(%r10, %r11)
  2208: 62117d487f641a05         vmovdqa32 %zmm12, 0x140(%r10, %r11)
  2216: 62117d487f6c1a06         vmovdqa32 %zmm13, 0x180(%r10, %r11)
  2224: 4d8b5110                 movq 0x10(%r9), %r10
  2228: 62117d487f341a           vmovdqa32 %zmm14, (%r10, %r11)
  2235: 62117d487f7c1a01         vmovdqa32 %zmm15, 0x40(%r10, %r11)
  2243: 62817d487f441a02         vmovdqa32 %zmm16, 0x80(%r10, %r11)
  2251: 62817d487f4c1a03         vmovdqa32 %zmm17, 0xc0(%r10, %r11)
  2259: 62817d487f541a04         vmovdqa32 %zmm18, 0x100(%r10, %r11)
  2267: 62817d487f5c1a05         vmovdqa32 %zmm19, 0x140(%r10, %r11)
  2275: 62817d487f641a06         vmovdqa32 %zmm20, 0x180(%r10, %r11)
  2283: 4981c3c0010000           addq $0x1c0, %r11
  2290: 4c3bd8                   cmpq %rax, %r11
  2293: 0f8645f7ffff             jbe 0x40
  2299: 4c3bd9                   cmpq %rcx, %r11
  2302: 0f83b2020000             jae 0xbb6
  2308: 4153                     pushq %r11
  2310: 498dbf40010000           leaq 0x140(%r15), %rdi
  2317: 498db740020000           leaq 0x240(%r15), %rsi
  2324: 4d8b11                   movq (%r9), %r10
  2327: 4c8b1c24                 movq (%rsp), %r11
  2331: 33db                     xorl %ebx, %ebx
  2333: 498b00                   movq (%r8), %rax
  2336: 420fb61418               movzbl (%rax, %r11), %edx
  2341: 32da                     xorb %dl, %bl
  2343: 498b4008                 movq 8(%r8), %rax
  2347: 420fb61418               movzbl (%rax, %r11), %edx
  2352: 32da                     xorb %dl, %bl
  2354: 498b4010                 movq 0x10(%r8), %rax
  2358: 420fb61418               movzbl (%rax, %r11), %edx
  2363: 32da                     xorb %dl, %bl
  2365: 498b4018                 movq 0x18(%r8), %rax
  2369: 420fb61418               movzbl (%rax, %r11), %edx
  2374: 32da                     xorb %dl, %bl
  2376: 498b4020                 movq 0x20(%r8), %rax
  2380: 420fb61418               movzbl (%rax, %r11), %edx
  2385: 32da                     xorb %dl, %bl
  2387: 498b4028                 movq 0x28(%r8), %rax
  2391: 420fb61418               movzbl (%rax, %r11), %edx
  2396: 32da                     xorb %dl, %bl
  2398: 43881c1a                 movb %bl, (%r10, %r11)
  2402: 33db                     xorl %ebx, %ebx
  2404: 49ffc3                   incq %r11
  2407: 4c3bd9                   cmpq %rcx, %r11
  2410: 0f82adffffff             jb 0x91d
  2416: 4d8b5108                 movq 8(%r9), %r10
  2420: 4c8b1c24                 movq (%rsp), %r11
  2424: 33db                     xorl %ebx, %ebx
  2426: 498b00                   movq (%r8), %rax
  2429: 420fb61418               movzbl (%rax, %r11), %edx
  2434: 4885d2                   testq %rdx, %rdx
  2437: 0f8422000000             je 0x9ad
  2443: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2447: 81c0e6000000             addl $0xe6, %eax
  2453: 81f8ff000000             cmpl $0xff, %eax
  2459: 0f8206000000             jb 0x9a7
  2465: 81c001ffffff             addl $0xffffff01, %eax
  2471: 0fb60406                 movzbl (%rsi, %rax), %eax
  2475: 32d8                     xorb %al, %bl
  2477: 498b4008                 movq 8(%r8), %rax
  2481: 420fb61418               movzbl (%rax, %r11), %edx
  2486: 4885d2                   testq %rdx, %rdx
  2489: 0f8422000000             je 0x9e1
  2495: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2499: 81c0cd000000             addl $0xcd, %eax
  2505: 81f8ff000000             cmpl $0xff, %eax
  2511: 0f8206000000             jb 0x9db
  2517: 81c001ffffff             addl $0xffffff01, %eax
  2523: 0fb60406                 movzbl (%rsi, %rax), %eax
  2527: 32d8                     xorb %al, %bl
  2529: 498b4010                 movq 0x10(%r8), %rax
  2533: 420fb61418               movzbl (%rax, %r11), %edx
  2538: 4885d2                   testq %rdx, %rdx
  2541: 0f841f000000             je 0xa12
  2547: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2551: 83c04b                   addl $0x4b, %eax
  2554: 81f8ff000000             cmpl $0xff, %eax
  2560: 0f8206000000             jb 0xa0c
  2566: 81c001ffffff             addl $0xffffff01, %eax
  2572: 0fb60406                 movzbl (%rsi, %rax), %eax
  2576: 32d8                     xorb %al, %bl
  2578: 498b4018                 movq 0x18(%r8), %rax
  2582: 420fb61418               movzbl (%rax, %r11), %edx
  2587: 4885d2                   testq %rdx, %rdx
  2590: 0f841f000000             je 0xa43
  2596: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2600: 83c032                   addl $0x32, %eax
  2603: 81f8ff000000             cmpl $0xff, %eax
  2609: 0f8206000000             jb 0xa3d
  2615: 81c001ffffff             addl $0xffffff01, %eax
  2621: 0fb60406                 movzbl (%rsi, %rax), %eax
  2625: 32d8                     xorb %al, %bl
  2627: 498b4020                 movq 0x20(%r8), %rax
  2631: 420fb61418               movzbl (%rax, %r11), %edx
  2636: 4885d2                   testq %rdx, %rdx
  2639: 0f841f000000             je 0xa74
  2645: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2649: 83c019                   addl $0x19, %eax
  2652: 81f8ff000000             cmpl $0xff, %eax
  2658: 0f8206000000             jb 0xa6e
  2664: 81c001ffffff             addl $0xffffff01, %eax
  2670: 0fb60406                 movzbl (%rsi, %rax), %eax
  2674: 32d8                     xorb %al, %bl
  2676: 498b4028                 movq 0x28(%r8), %rax
  2680: 420fb61418               movzbl (%rax, %r11), %edx
  2685: 32da                     xorb %dl, %bl
  2687: 43881c1a                 movb %bl, (%r10, %r11)
  2691: 33db                     xorl %ebx, %ebx
  2693: 49ffc3                   incq %r11
  2696: 4c3bd9                   cmpq %rcx, %r11
  2699: 0f82e9feffff             jb 0x97a
  2705: 4d8b5110                 movq 0x10(%r9), %r10
  2709: 4c8b1c24                 movq (%rsp), %r11
  2713: 33db                     xorl %ebx, %ebx
  2715: 498b00                   movq (%r8), %rax
  2718: 420fb61418               movzbl (%rax, %r11), %edx
  2723: 4885d2                   testq %rdx, %rdx
  2726: 0f8422000000             je 0xace
  2732: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2736: 81c0cd000000             addl $0xcd, %eax
  2742: 81f8ff000000             cmpl $0xff, %eax
  2748: 0f8206000000             jb 0xac8
  2754: 81c001ffffff             addl $0xffffff01, %eax
  2760: 0fb60406                 movzbl (%rsi, %rax), %eax
  2764: 32d8                     xorb %al, %bl
  2766: 498b4008                 movq 8(%r8), %rax
  2770: 420fb61418               movzbl (%rax, %r11), %edx
  2775: 4885d2                   testq %rdx, %rdx
  2778: 0f841f000000             je 0xaff
  2784: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2788: 83c04b                   addl $0x4b, %eax
  2791: 81f8ff000000             cmpl $0xff, %eax
  2797: 0f8206000000             jb 0xaf9
  2803: 81c001ffffff             addl $0xffffff01, %eax
  2809: 0fb60406                 movzbl (%rsi, %rax), %eax
  2813: 32d8                     xorb %al, %bl
  2815: 498b4010                 movq 0x10(%r8), %rax
  2819: 420fb61418               movzbl (%rax, %r11), %edx
  2824: 4885d2                   testq %rdx, %rdx
  2827: 0f841f000000             je 0xb30
  2833: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2837: 83c019                   addl $0x19, %eax
  2840: 81f8ff000000             cmpl $0xff, %eax
  2846: 0f8206000000             jb 0xb2a
  2852: 81c001ffffff             addl $0xffffff01, %eax
  2858: 0fb60406                 movzbl (%rsi, %rax), %eax
  2862: 32d8                     xorb %al, %bl
  2864: 498b4018                 movq 0x18(%r8), %rax
  2868: 420fb61418               movzbl (%rax, %r11), %edx
  2873: 4885d2                   testq %rdx, %rdx
  2876: 0f8422000000             je 0xb64
  2882: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2886: 81c0e6000000             addl $0xe6, %eax
  2892: 81f8ff000000             cmpl $0xff, %eax
  2898: 0f8206000000             jb 0xb5e
  2904: 81c001ffffff             addl $0xffffff01, %eax
  2910: 0fb60406                 movzbl (%rsi, %rax), %eax
  2914: 32d8                     xorb %al, %bl
  2916: 498b4020                 movq 0x20(%r8), %rax
  2920: 420fb61418               movzbl (%rax, %r11), %edx
  2925: 4885d2                   testq %rdx, %rdx
  2928: 0f841f000000             je 0xb95
  2934: 0fb60417                 movzbl (%rdi, %rdx), %eax
  2938: 83c032                   addl $0x32, %eax
  2941: 81f8ff000000             cmpl $0xff, %eax
  2947: 0f8206000000             jb 0xb8f
  2953: 81c001ffffff             addl $0xffffff01, %eax
  2959: 0fb60406                 movzbl (%rsi, %rax), %eax
  2963: 32d8                     xorb %al, %bl
  2965: 498b4028                 movq 0x28(%r8), %rax
  2969: 420fb61418               movzbl (%rax, %r11), %edx
  2974: 32da                     xorb %dl, %bl
  2976: 43881c1a                 movb %bl, (%r10, %r11)
  2980: 33db                     xorl %ebx, %ebx
  2982: 49ffc3                   incq %r11
  2985: 4c3bd9                   cmpq %rcx, %r11
  2988: 0f82e9feffff             jb 0xa9b
  2994: 4883c408                 addq $8, %rsp
  2998: 415e                     popq %r14
  3000: 415d                     popq %r13
  3002: 415c                     popq %r12
  3004: 415f                     popq %r15
  3006: 5b                       popq %rbx
  3007: c5f877                   vzeroupper 
  3010: c3                       retq 
  3011: 0f0b                     ud2 
  3013: 0f0b                     ud2 
  3015: 0f0b                     ud2 
  3017: 0f0b                     ud2 
  3019: 0f0b                     ud2 
  3021: 0f0b                     ud2 
  3023: 0f0b                     ud2 
  3025: 0f0b                     ud2 
  3027: 0f0b                     ud2 
  3029: 0f0b                     ud2 
  3031: 0f0b                     ud2 
  3033: 0f0b                     ud2 
  3035: 0f0b                     ud2 
  3037: 0f0b                     ud2 
  3039: 0f0b                     ud2 
  3041: 0f0b                     ud2 
  3043: 0f0b                     ud2 
  3045: 0f0b                     ud2 
  3047: 0f0b                     ud2 
  3049: 0f0b                     ud2 
  3051: 0f0b                     ud2 
  3053: 0f0b                     ud2 
  3055: 0f0b                     ud2 
  3057: 0f0b                     ud2 
  3059: 0f0b                     ud2 
  3061: 0f0b                     ud2 
  3063: 0f0b                     ud2 
  3065: 0f0b                     ud2 
  3067: 0f0b                     ud2 
  3069: 0f0b                     ud2 
  3071: cc                       int3 
