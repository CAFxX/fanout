; config_t1_95 target=avx512-gfni twin=notail topo=rs(6,3)/enc op=encode U=8
; code_size=2688 twins=["notail"] uops_per_byte=0.732421875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d670a0000           leaq 0xa67(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e800020000           subq $0x200, %rax
    57: 0f82fa090000             jb 0xa39
    63: 90                       nop 
    64: 62217d486f041e           vmovdqa32 (%rsi, %r11), %zmm24
    71: 62917d486fc0             vmovdqa32 %zmm24, %zmm0
    77: 62417d486f17             vmovdqa32 (%r15), %zmm26
    83: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
    90: 62117d486fc1             vmovdqa32 %zmm25, %zmm8
    96: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   103: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   110: 62817d486fc1             vmovdqa32 %zmm25, %zmm16
   116: 62217d486f441e01         vmovdqa32 0x40(%rsi, %r11), %zmm24
   124: 62917d486fc8             vmovdqa32 %zmm24, %zmm1
   130: 62417d486f17             vmovdqa32 (%r15), %zmm26
   136: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   143: 62117d486fc9             vmovdqa32 %zmm25, %zmm9
   149: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   156: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   163: 62817d486fc9             vmovdqa32 %zmm25, %zmm17
   169: 62217d486f441e02         vmovdqa32 0x80(%rsi, %r11), %zmm24
   177: 62917d486fd0             vmovdqa32 %zmm24, %zmm2
   183: 62417d486f17             vmovdqa32 (%r15), %zmm26
   189: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   196: 62117d486fd1             vmovdqa32 %zmm25, %zmm10
   202: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   209: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   216: 62817d486fd1             vmovdqa32 %zmm25, %zmm18
   222: 62217d486f441e03         vmovdqa32 0xc0(%rsi, %r11), %zmm24
   230: 62917d486fd8             vmovdqa32 %zmm24, %zmm3
   236: 62417d486f17             vmovdqa32 (%r15), %zmm26
   242: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   249: 62117d486fd9             vmovdqa32 %zmm25, %zmm11
   255: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   262: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   269: 62817d486fd9             vmovdqa32 %zmm25, %zmm19
   275: 62217d486f441e04         vmovdqa32 0x100(%rsi, %r11), %zmm24
   283: 62917d486fe0             vmovdqa32 %zmm24, %zmm4
   289: 62417d486f17             vmovdqa32 (%r15), %zmm26
   295: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   302: 62117d486fe1             vmovdqa32 %zmm25, %zmm12
   308: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   315: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   322: 62817d486fe1             vmovdqa32 %zmm25, %zmm20
   328: 62217d486f441e05         vmovdqa32 0x140(%rsi, %r11), %zmm24
   336: 62917d486fe8             vmovdqa32 %zmm24, %zmm5
   342: 62417d486f17             vmovdqa32 (%r15), %zmm26
   348: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   355: 62117d486fe9             vmovdqa32 %zmm25, %zmm13
   361: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   368: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   375: 62817d486fe9             vmovdqa32 %zmm25, %zmm21
   381: 62217d486f441e06         vmovdqa32 0x180(%rsi, %r11), %zmm24
   389: 62917d486ff0             vmovdqa32 %zmm24, %zmm6
   395: 62417d486f17             vmovdqa32 (%r15), %zmm26
   401: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   408: 62117d486ff1             vmovdqa32 %zmm25, %zmm14
   414: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   421: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   428: 62817d486ff1             vmovdqa32 %zmm25, %zmm22
   434: 62217d486f441e07         vmovdqa32 0x1c0(%rsi, %r11), %zmm24
   442: 62917d486ff8             vmovdqa32 %zmm24, %zmm7
   448: 62417d486f17             vmovdqa32 (%r15), %zmm26
   454: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   461: 62117d486ff9             vmovdqa32 %zmm25, %zmm15
   467: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   474: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   481: 62817d486ff9             vmovdqa32 %zmm25, %zmm23
   487: 62217d486f041f           vmovdqa32 (%rdi, %r11), %zmm24
   494: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
   500: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   507: 62113d48efc1             vpxord %zmm25, %zmm8, %zmm8
   513: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   520: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   527: 62817d40efc1             vpxord %zmm25, %zmm16, %zmm16
   533: 62217d486f441f01         vmovdqa32 0x40(%rdi, %r11), %zmm24
   541: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
   547: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   554: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   561: 62113548efc9             vpxord %zmm25, %zmm9, %zmm9
   567: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   574: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   581: 62817540efc9             vpxord %zmm25, %zmm17, %zmm17
   587: 62217d486f441f02         vmovdqa32 0x80(%rdi, %r11), %zmm24
   595: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
   601: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   608: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   615: 62112d48efd1             vpxord %zmm25, %zmm10, %zmm10
   621: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   628: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   635: 62816d40efd1             vpxord %zmm25, %zmm18, %zmm18
   641: 62217d486f441f03         vmovdqa32 0xc0(%rdi, %r11), %zmm24
   649: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
   655: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   662: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   669: 62112548efd9             vpxord %zmm25, %zmm11, %zmm11
   675: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   682: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   689: 62816540efd9             vpxord %zmm25, %zmm19, %zmm19
   695: 62217d486f441f04         vmovdqa32 0x100(%rdi, %r11), %zmm24
   703: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
   709: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   716: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   723: 62111d48efe1             vpxord %zmm25, %zmm12, %zmm12
   729: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   736: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   743: 62815d40efe1             vpxord %zmm25, %zmm20, %zmm20
   749: 62217d486f441f05         vmovdqa32 0x140(%rdi, %r11), %zmm24
   757: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
   763: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   770: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   777: 62111548efe9             vpxord %zmm25, %zmm13, %zmm13
   783: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   790: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   797: 62815540efe9             vpxord %zmm25, %zmm21, %zmm21
   803: 62217d486f441f06         vmovdqa32 0x180(%rdi, %r11), %zmm24
   811: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
   817: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   824: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   831: 62110d48eff1             vpxord %zmm25, %zmm14, %zmm14
   837: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   844: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   851: 62814d40eff1             vpxord %zmm25, %zmm22, %zmm22
   857: 62217d486f441f07         vmovdqa32 0x1c0(%rdi, %r11), %zmm24
   865: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
   871: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   878: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   885: 62110548eff9             vpxord %zmm25, %zmm15, %zmm15
   891: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   898: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   905: 62814540eff9             vpxord %zmm25, %zmm23, %zmm23
   911: 62017d486f041c           vmovdqa32 (%r12, %r11), %zmm24
   918: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
   924: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   931: 62113d48efc1             vpxord %zmm25, %zmm8, %zmm8
   937: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
   944: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   951: 62817d40efc1             vpxord %zmm25, %zmm16, %zmm16
   957: 62017d486f441c01         vmovdqa32 0x40(%r12, %r11), %zmm24
   965: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
   971: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
   978: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   985: 62113548efc9             vpxord %zmm25, %zmm9, %zmm9
   991: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
   998: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1005: 62817540efc9             vpxord %zmm25, %zmm17, %zmm17
  1011: 62017d486f441c02         vmovdqa32 0x80(%r12, %r11), %zmm24
  1019: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
  1025: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
  1032: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1039: 62112d48efd1             vpxord %zmm25, %zmm10, %zmm10
  1045: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1052: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1059: 62816d40efd1             vpxord %zmm25, %zmm18, %zmm18
  1065: 62017d486f441c03         vmovdqa32 0xc0(%r12, %r11), %zmm24
  1073: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
  1079: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
  1086: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1093: 62112548efd9             vpxord %zmm25, %zmm11, %zmm11
  1099: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1106: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1113: 62816540efd9             vpxord %zmm25, %zmm19, %zmm19
  1119: 62017d486f441c04         vmovdqa32 0x100(%r12, %r11), %zmm24
  1127: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
  1133: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
  1140: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1147: 62111d48efe1             vpxord %zmm25, %zmm12, %zmm12
  1153: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1160: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1167: 62815d40efe1             vpxord %zmm25, %zmm20, %zmm20
  1173: 62017d486f441c05         vmovdqa32 0x140(%r12, %r11), %zmm24
  1181: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
  1187: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
  1194: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1201: 62111548efe9             vpxord %zmm25, %zmm13, %zmm13
  1207: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1214: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1221: 62815540efe9             vpxord %zmm25, %zmm21, %zmm21
  1227: 62017d486f441c06         vmovdqa32 0x180(%r12, %r11), %zmm24
  1235: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
  1241: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
  1248: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1255: 62110d48eff1             vpxord %zmm25, %zmm14, %zmm14
  1261: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1268: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1275: 62814d40eff1             vpxord %zmm25, %zmm22, %zmm22
  1281: 62017d486f441c07         vmovdqa32 0x1c0(%r12, %r11), %zmm24
  1289: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
  1295: 62417d486f5702           vmovdqa32 0x80(%r15), %zmm26
  1302: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1309: 62110548eff9             vpxord %zmm25, %zmm15, %zmm15
  1315: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1322: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1329: 62814540eff9             vpxord %zmm25, %zmm23, %zmm23
  1335: 62017d486f441d00         vmovdqa32 (%r13, %r11), %zmm24
  1343: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
  1349: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1356: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1363: 62113d48efc1             vpxord %zmm25, %zmm8, %zmm8
  1369: 62417d486f17             vmovdqa32 (%r15), %zmm26
  1375: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1382: 62817d40efc1             vpxord %zmm25, %zmm16, %zmm16
  1388: 62017d486f441d01         vmovdqa32 0x40(%r13, %r11), %zmm24
  1396: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
  1402: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1409: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1416: 62113548efc9             vpxord %zmm25, %zmm9, %zmm9
  1422: 62417d486f17             vmovdqa32 (%r15), %zmm26
  1428: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1435: 62817540efc9             vpxord %zmm25, %zmm17, %zmm17
  1441: 62017d486f441d02         vmovdqa32 0x80(%r13, %r11), %zmm24
  1449: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
  1455: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1462: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1469: 62112d48efd1             vpxord %zmm25, %zmm10, %zmm10
  1475: 62417d486f17             vmovdqa32 (%r15), %zmm26
  1481: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1488: 62816d40efd1             vpxord %zmm25, %zmm18, %zmm18
  1494: 62017d486f441d03         vmovdqa32 0xc0(%r13, %r11), %zmm24
  1502: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
  1508: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1515: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1522: 62112548efd9             vpxord %zmm25, %zmm11, %zmm11
  1528: 62417d486f17             vmovdqa32 (%r15), %zmm26
  1534: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1541: 62816540efd9             vpxord %zmm25, %zmm19, %zmm19
  1547: 62017d486f441d04         vmovdqa32 0x100(%r13, %r11), %zmm24
  1555: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
  1561: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1568: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1575: 62111d48efe1             vpxord %zmm25, %zmm12, %zmm12
  1581: 62417d486f17             vmovdqa32 (%r15), %zmm26
  1587: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1594: 62815d40efe1             vpxord %zmm25, %zmm20, %zmm20
  1600: 62017d486f441d05         vmovdqa32 0x140(%r13, %r11), %zmm24
  1608: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
  1614: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1621: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1628: 62111548efe9             vpxord %zmm25, %zmm13, %zmm13
  1634: 62417d486f17             vmovdqa32 (%r15), %zmm26
  1640: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1647: 62815540efe9             vpxord %zmm25, %zmm21, %zmm21
  1653: 62017d486f441d06         vmovdqa32 0x180(%r13, %r11), %zmm24
  1661: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
  1667: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1674: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1681: 62110d48eff1             vpxord %zmm25, %zmm14, %zmm14
  1687: 62417d486f17             vmovdqa32 (%r15), %zmm26
  1693: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1700: 62814d40eff1             vpxord %zmm25, %zmm22, %zmm22
  1706: 62017d486f441d07         vmovdqa32 0x1c0(%r13, %r11), %zmm24
  1714: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
  1720: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1727: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1734: 62110548eff9             vpxord %zmm25, %zmm15, %zmm15
  1740: 62417d486f17             vmovdqa32 (%r15), %zmm26
  1746: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1753: 62814540eff9             vpxord %zmm25, %zmm23, %zmm23
  1759: 62017d486f041e           vmovdqa32 (%r14, %r11), %zmm24
  1766: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
  1772: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1779: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1786: 62113d48efc1             vpxord %zmm25, %zmm8, %zmm8
  1792: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1799: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1806: 62817d40efc1             vpxord %zmm25, %zmm16, %zmm16
  1812: 62017d486f441e01         vmovdqa32 0x40(%r14, %r11), %zmm24
  1820: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
  1826: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1833: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1840: 62113548efc9             vpxord %zmm25, %zmm9, %zmm9
  1846: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1853: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1860: 62817540efc9             vpxord %zmm25, %zmm17, %zmm17
  1866: 62017d486f441e02         vmovdqa32 0x80(%r14, %r11), %zmm24
  1874: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
  1880: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1887: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1894: 62112d48efd1             vpxord %zmm25, %zmm10, %zmm10
  1900: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1907: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1914: 62816d40efd1             vpxord %zmm25, %zmm18, %zmm18
  1920: 62017d486f441e03         vmovdqa32 0xc0(%r14, %r11), %zmm24
  1928: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
  1934: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1941: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1948: 62112548efd9             vpxord %zmm25, %zmm11, %zmm11
  1954: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  1961: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  1968: 62816540efd9             vpxord %zmm25, %zmm19, %zmm19
  1974: 62017d486f441e04         vmovdqa32 0x100(%r14, %r11), %zmm24
  1982: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
  1988: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  1995: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  2002: 62111d48efe1             vpxord %zmm25, %zmm12, %zmm12
  2008: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  2015: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  2022: 62815d40efe1             vpxord %zmm25, %zmm20, %zmm20
  2028: 62017d486f441e05         vmovdqa32 0x140(%r14, %r11), %zmm24
  2036: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
  2042: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  2049: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  2056: 62111548efe9             vpxord %zmm25, %zmm13, %zmm13
  2062: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  2069: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  2076: 62815540efe9             vpxord %zmm25, %zmm21, %zmm21
  2082: 62017d486f441e06         vmovdqa32 0x180(%r14, %r11), %zmm24
  2090: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
  2096: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  2103: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  2110: 62110d48eff1             vpxord %zmm25, %zmm14, %zmm14
  2116: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  2123: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  2130: 62814d40eff1             vpxord %zmm25, %zmm22, %zmm22
  2136: 62017d486f441e07         vmovdqa32 0x1c0(%r14, %r11), %zmm24
  2144: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
  2150: 62417d486f5703           vmovdqa32 0xc0(%r15), %zmm26
  2157: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  2164: 62110548eff9             vpxord %zmm25, %zmm15, %zmm15
  2170: 62417d486f5704           vmovdqa32 0x100(%r15), %zmm26
  2177: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
  2184: 62814540eff9             vpxord %zmm25, %zmm23, %zmm23
  2190: 4d8b5028                 movq 0x28(%r8), %r10
  2194: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
  2201: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
  2207: 62113d48efc0             vpxord %zmm24, %zmm8, %zmm8
  2213: 62817d40efc0             vpxord %zmm24, %zmm16, %zmm16
  2219: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
  2227: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
  2233: 62113548efc8             vpxord %zmm24, %zmm9, %zmm9
  2239: 62817540efc8             vpxord %zmm24, %zmm17, %zmm17
  2245: 62017d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm24
  2253: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
  2259: 62112d48efd0             vpxord %zmm24, %zmm10, %zmm10
  2265: 62816d40efd0             vpxord %zmm24, %zmm18, %zmm18
  2271: 62017d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm24
  2279: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
  2285: 62112548efd8             vpxord %zmm24, %zmm11, %zmm11
  2291: 62816540efd8             vpxord %zmm24, %zmm19, %zmm19
  2297: 62017d486f441a04         vmovdqa32 0x100(%r10, %r11), %zmm24
  2305: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
  2311: 62111d48efe0             vpxord %zmm24, %zmm12, %zmm12
  2317: 62815d40efe0             vpxord %zmm24, %zmm20, %zmm20
  2323: 62017d486f441a05         vmovdqa32 0x140(%r10, %r11), %zmm24
  2331: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
  2337: 62111548efe8             vpxord %zmm24, %zmm13, %zmm13
  2343: 62815540efe8             vpxord %zmm24, %zmm21, %zmm21
  2349: 62017d486f441a06         vmovdqa32 0x180(%r10, %r11), %zmm24
  2357: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
  2363: 62110d48eff0             vpxord %zmm24, %zmm14, %zmm14
  2369: 62814d40eff0             vpxord %zmm24, %zmm22, %zmm22
  2375: 62017d486f441a07         vmovdqa32 0x1c0(%r10, %r11), %zmm24
  2383: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
  2389: 62110548eff8             vpxord %zmm24, %zmm15, %zmm15
  2395: 62814540eff8             vpxord %zmm24, %zmm23, %zmm23
  2401: 4d8b11                   movq (%r9), %r10
  2404: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
  2411: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
  2419: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
  2427: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
  2435: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
  2443: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
  2451: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
  2459: 62917d487f7c1a07         vmovdqa32 %zmm7, 0x1c0(%r10, %r11)
  2467: 4d8b5108                 movq 8(%r9), %r10
  2471: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
  2478: 62117d487f4c1a01         vmovdqa32 %zmm9, 0x40(%r10, %r11)
  2486: 62117d487f541a02         vmovdqa32 %zmm10, 0x80(%r10, %r11)
  2494: 62117d487f5c1a03         vmovdqa32 %zmm11, 0xc0(%r10, %r11)
  2502: 62117d487f641a04         vmovdqa32 %zmm12, 0x100(%r10, %r11)
  2510: 62117d487f6c1a05         vmovdqa32 %zmm13, 0x140(%r10, %r11)
  2518: 62117d487f741a06         vmovdqa32 %zmm14, 0x180(%r10, %r11)
  2526: 62117d487f7c1a07         vmovdqa32 %zmm15, 0x1c0(%r10, %r11)
  2534: 4d8b5110                 movq 0x10(%r9), %r10
  2538: 62817d487f041a           vmovdqa32 %zmm16, (%r10, %r11)
  2545: 62817d487f4c1a01         vmovdqa32 %zmm17, 0x40(%r10, %r11)
  2553: 62817d487f541a02         vmovdqa32 %zmm18, 0x80(%r10, %r11)
  2561: 62817d487f5c1a03         vmovdqa32 %zmm19, 0xc0(%r10, %r11)
  2569: 62817d487f641a04         vmovdqa32 %zmm20, 0x100(%r10, %r11)
  2577: 62817d487f6c1a05         vmovdqa32 %zmm21, 0x140(%r10, %r11)
  2585: 62817d487f741a06         vmovdqa32 %zmm22, 0x180(%r10, %r11)
  2593: 62817d487f7c1a07         vmovdqa32 %zmm23, 0x1c0(%r10, %r11)
  2601: 4981c300020000           addq $0x200, %r11
  2608: 4c3bd8                   cmpq %rax, %r11
  2611: 0f8607f6ffff             jbe 0x40
  2617: 415e                     popq %r14
  2619: 415d                     popq %r13
  2621: 415c                     popq %r12
  2623: 415f                     popq %r15
  2625: 5b                       popq %rbx
  2626: c5f877                   vzeroupper 
  2629: c3                       retq 
  2630: 0f0b                     ud2 
  2632: 0f0b                     ud2 
  2634: 0f0b                     ud2 
  2636: 0f0b                     ud2 
  2638: 0f0b                     ud2 
  2640: 0f0b                     ud2 
  2642: 0f0b                     ud2 
  2644: 0f0b                     ud2 
  2646: 0f0b                     ud2 
  2648: 0f0b                     ud2 
  2650: 0f0b                     ud2 
  2652: 0f0b                     ud2 
  2654: 0f0b                     ud2 
  2656: 0f0b                     ud2 
  2658: 0f0b                     ud2 
  2660: 0f0b                     ud2 
  2662: 0f0b                     ud2 
  2664: 0f0b                     ud2 
  2666: 0f0b                     ud2 
  2668: 0f0b                     ud2 
  2670: 0f0b                     ud2 
  2672: 0f0b                     ud2 
  2674: 0f0b                     ud2 
  2676: 0f0b                     ud2 
  2678: 0f0b                     ud2 
  2680: 0f0b                     ud2 
  2682: 0f0b                     ud2 
  2684: 0f0b                     ud2 
  2686: 0f0b                     ud2 
