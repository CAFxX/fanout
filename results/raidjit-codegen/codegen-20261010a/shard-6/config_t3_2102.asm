; config_t3_2102 target=avx512-gfni twin=baseline topo=lrc(8,2,2)/rec{0} op=recover U=4
; code_size=1280 twins=[] uops_per_byte=0.4140625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de7040000           leaq 0x4e7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4881e800010000           subq $0x100, %rax
    57: 0f82ec020000             jb 0x32b
    63: 90                       nop 
    64: 62b17d486f241e           vmovdqa32 (%rsi, %r11), %zmm4
    71: 62d17d486f37             vmovdqa32 (%r15), %zmm6
    77: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
    84: 62f17d486fc5             vmovdqa32 %zmm5, %zmm0
    90: 62b17d486f641e01         vmovdqa32 0x40(%rsi, %r11), %zmm4
    98: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   105: 62f17d486fcd             vmovdqa32 %zmm5, %zmm1
   111: 62b17d486f641e02         vmovdqa32 0x80(%rsi, %r11), %zmm4
   119: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   126: 62f17d486fd5             vmovdqa32 %zmm5, %zmm2
   132: 62b17d486f641e03         vmovdqa32 0xc0(%rsi, %r11), %zmm4
   140: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   147: 62f17d486fdd             vmovdqa32 %zmm5, %zmm3
   153: 62b17d486f241f           vmovdqa32 (%rdi, %r11), %zmm4
   160: 62d17d486f7701           vmovdqa32 0x40(%r15), %zmm6
   167: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   174: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   180: 62b17d486f641f01         vmovdqa32 0x40(%rdi, %r11), %zmm4
   188: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   195: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   201: 62b17d486f641f02         vmovdqa32 0x80(%rdi, %r11), %zmm4
   209: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   216: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   222: 62b17d486f641f03         vmovdqa32 0xc0(%rdi, %r11), %zmm4
   230: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   237: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   243: 62917d486f241c           vmovdqa32 (%r12, %r11), %zmm4
   250: 62d17d486f7702           vmovdqa32 0x80(%r15), %zmm6
   257: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   264: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   270: 62917d486f641c01         vmovdqa32 0x40(%r12, %r11), %zmm4
   278: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   285: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   291: 62917d486f641c02         vmovdqa32 0x80(%r12, %r11), %zmm4
   299: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   306: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   312: 62917d486f641c03         vmovdqa32 0xc0(%r12, %r11), %zmm4
   320: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   327: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   333: 62917d486f641d00         vmovdqa32 (%r13, %r11), %zmm4
   341: 62d17d486f7703           vmovdqa32 0xc0(%r15), %zmm6
   348: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   355: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   361: 62917d486f641d01         vmovdqa32 0x40(%r13, %r11), %zmm4
   369: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   376: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   382: 62917d486f641d02         vmovdqa32 0x80(%r13, %r11), %zmm4
   390: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   397: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   403: 62917d486f641d03         vmovdqa32 0xc0(%r13, %r11), %zmm4
   411: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   418: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   424: 62917d486f241e           vmovdqa32 (%r14, %r11), %zmm4
   431: 62d17d486f7704           vmovdqa32 0x100(%r15), %zmm6
   438: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   445: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   451: 62917d486f641e01         vmovdqa32 0x40(%r14, %r11), %zmm4
   459: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   466: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   472: 62917d486f641e02         vmovdqa32 0x80(%r14, %r11), %zmm4
   480: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   487: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   493: 62917d486f641e03         vmovdqa32 0xc0(%r14, %r11), %zmm4
   501: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   508: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   514: 4d8b5028                 movq 0x28(%r8), %r10
   518: 62917d486f241a           vmovdqa32 (%r10, %r11), %zmm4
   525: 62d17d486f7705           vmovdqa32 0x140(%r15), %zmm6
   532: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   539: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   545: 62917d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm4
   553: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   560: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   566: 62917d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm4
   574: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   581: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   587: 62917d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm4
   595: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   602: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   608: 4d8b5030                 movq 0x30(%r8), %r10
   612: 62917d486f241a           vmovdqa32 (%r10, %r11), %zmm4
   619: 62d17d486f7706           vmovdqa32 0x180(%r15), %zmm6
   626: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   633: 62f17d48efc5             vpxord %zmm5, %zmm0, %zmm0
   639: 62917d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm4
   647: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   654: 62f17548efcd             vpxord %zmm5, %zmm1, %zmm1
   660: 62917d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm4
   668: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   675: 62f16d48efd5             vpxord %zmm5, %zmm2, %zmm2
   681: 62917d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm4
   689: 62f3dd48ceee00           vgf2p8affineqb $0, %zmm6, %zmm4, %zmm5
   696: 62f16548efdd             vpxord %zmm5, %zmm3, %zmm3
   702: 4d8b5038                 movq 0x38(%r8), %r10
   706: 62917d486f241a           vmovdqa32 (%r10, %r11), %zmm4
   713: 62f17d48efc4             vpxord %zmm4, %zmm0, %zmm0
   719: 62917d486f641a01         vmovdqa32 0x40(%r10, %r11), %zmm4
   727: 62f17548efcc             vpxord %zmm4, %zmm1, %zmm1
   733: 62917d486f641a02         vmovdqa32 0x80(%r10, %r11), %zmm4
   741: 62f16d48efd4             vpxord %zmm4, %zmm2, %zmm2
   747: 62917d486f641a03         vmovdqa32 0xc0(%r10, %r11), %zmm4
   755: 62f16548efdc             vpxord %zmm4, %zmm3, %zmm3
   761: 4d8b11                   movq (%r9), %r10
   764: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   771: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
   779: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
   787: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
   795: 4981c300010000           addq $0x100, %r11
   802: 4c3bd8                   cmpq %rax, %r11
   805: 0f8615fdffff             jbe 0x40
   811: 4c3bd9                   cmpq %rcx, %r11
   814: 0f8396010000             jae 0x4ca
   820: 4153                     pushq %r11
   822: 498dbfc0010000           leaq 0x1c0(%r15), %rdi
   829: 498db7c0020000           leaq 0x2c0(%r15), %rsi
   836: 4d8b11                   movq (%r9), %r10
   839: 4c8b1c24                 movq (%rsp), %r11
   843: 33db                     xorl %ebx, %ebx
   845: 498b00                   movq (%r8), %rax
   848: 420fb61418               movzbl (%rax, %r11), %edx
   853: 4885d2                   testq %rdx, %rdx
   856: 0f841f000000             je 0x37d
   862: 0fb60417                 movzbl (%rdi, %rdx), %eax
   866: 83c019                   addl $0x19, %eax
   869: 81f8ff000000             cmpl $0xff, %eax
   875: 0f8206000000             jb 0x377
   881: 81c001ffffff             addl $0xffffff01, %eax
   887: 0fb60406                 movzbl (%rsi, %rax), %eax
   891: 32d8                     xorb %al, %bl
   893: 498b4008                 movq 8(%r8), %rax
   897: 420fb61418               movzbl (%rax, %r11), %edx
   902: 4885d2                   testq %rdx, %rdx
   905: 0f841f000000             je 0x3ae
   911: 0fb60417                 movzbl (%rdi, %rdx), %eax
   915: 83c032                   addl $0x32, %eax
   918: 81f8ff000000             cmpl $0xff, %eax
   924: 0f8206000000             jb 0x3a8
   930: 81c001ffffff             addl $0xffffff01, %eax
   936: 0fb60406                 movzbl (%rsi, %rax), %eax
   940: 32d8                     xorb %al, %bl
   942: 498b4010                 movq 0x10(%r8), %rax
   946: 420fb61418               movzbl (%rax, %r11), %edx
   951: 4885d2                   testq %rdx, %rdx
   954: 0f841f000000             je 0x3df
   960: 0fb60417                 movzbl (%rdi, %rdx), %eax
   964: 83c04b                   addl $0x4b, %eax
   967: 81f8ff000000             cmpl $0xff, %eax
   973: 0f8206000000             jb 0x3d9
   979: 81c001ffffff             addl $0xffffff01, %eax
   985: 0fb60406                 movzbl (%rsi, %rax), %eax
   989: 32d8                     xorb %al, %bl
   991: 498b4018                 movq 0x18(%r8), %rax
   995: 420fb61418               movzbl (%rax, %r11), %edx
  1000: 4885d2                   testq %rdx, %rdx
  1003: 0f841f000000             je 0x410
  1009: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1013: 83c064                   addl $0x64, %eax
  1016: 81f8ff000000             cmpl $0xff, %eax
  1022: 0f8206000000             jb 0x40a
  1028: 81c001ffffff             addl $0xffffff01, %eax
  1034: 0fb60406                 movzbl (%rsi, %rax), %eax
  1038: 32d8                     xorb %al, %bl
  1040: 498b4020                 movq 0x20(%r8), %rax
  1044: 420fb61418               movzbl (%rax, %r11), %edx
  1049: 4885d2                   testq %rdx, %rdx
  1052: 0f841f000000             je 0x441
  1058: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1062: 83c07d                   addl $0x7d, %eax
  1065: 81f8ff000000             cmpl $0xff, %eax
  1071: 0f8206000000             jb 0x43b
  1077: 81c001ffffff             addl $0xffffff01, %eax
  1083: 0fb60406                 movzbl (%rsi, %rax), %eax
  1087: 32d8                     xorb %al, %bl
  1089: 498b4028                 movq 0x28(%r8), %rax
  1093: 420fb61418               movzbl (%rax, %r11), %edx
  1098: 4885d2                   testq %rdx, %rdx
  1101: 0f8422000000             je 0x475
  1107: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1111: 81c096000000             addl $0x96, %eax
  1117: 81f8ff000000             cmpl $0xff, %eax
  1123: 0f8206000000             jb 0x46f
  1129: 81c001ffffff             addl $0xffffff01, %eax
  1135: 0fb60406                 movzbl (%rsi, %rax), %eax
  1139: 32d8                     xorb %al, %bl
  1141: 498b4030                 movq 0x30(%r8), %rax
  1145: 420fb61418               movzbl (%rax, %r11), %edx
  1150: 4885d2                   testq %rdx, %rdx
  1153: 0f8422000000             je 0x4a9
  1159: 0fb60417                 movzbl (%rdi, %rdx), %eax
  1163: 81c0af000000             addl $0xaf, %eax
  1169: 81f8ff000000             cmpl $0xff, %eax
  1175: 0f8206000000             jb 0x4a3
  1181: 81c001ffffff             addl $0xffffff01, %eax
  1187: 0fb60406                 movzbl (%rsi, %rax), %eax
  1191: 32d8                     xorb %al, %bl
  1193: 498b4038                 movq 0x38(%r8), %rax
  1197: 420fb61418               movzbl (%rax, %r11), %edx
  1202: 32da                     xorb %dl, %bl
  1204: 43881c1a                 movb %bl, (%r10, %r11)
  1208: 33db                     xorl %ebx, %ebx
  1210: 49ffc3                   incq %r11
  1213: 4c3bd9                   cmpq %rcx, %r11
  1216: 0f8287feffff             jb 0x34d
  1222: 4883c408                 addq $8, %rsp
  1226: 415e                     popq %r14
  1228: 415d                     popq %r13
  1230: 415c                     popq %r12
  1232: 415f                     popq %r15
  1234: 5b                       popq %rbx
  1235: c5f877                   vzeroupper 
  1238: c3                       retq 
  1239: 0f0b                     ud2 
  1241: 0f0b                     ud2 
  1243: 0f0b                     ud2 
  1245: 0f0b                     ud2 
  1247: 0f0b                     ud2 
  1249: 0f0b                     ud2 
  1251: 0f0b                     ud2 
  1253: 0f0b                     ud2 
  1255: 0f0b                     ud2 
  1257: 0f0b                     ud2 
  1259: 0f0b                     ud2 
  1261: 0f0b                     ud2 
  1263: 0f0b                     ud2 
  1265: 0f0b                     ud2 
  1267: 0f0b                     ud2 
  1269: 0f0b                     ud2 
  1271: 0f0b                     ud2 
  1273: 0f0b                     ud2 
  1275: 0f0b                     ud2 
  1277: 0f0b                     ud2 
  1279: cc                       int3 
