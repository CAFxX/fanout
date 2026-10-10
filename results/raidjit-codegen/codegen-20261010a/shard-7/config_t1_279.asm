; config_t1_279 target=avx512-gfni twin=maskedtail topo=raid6(8)/upd{0,1} op=parity_update U=8
; code_size=1152 twins=["maskedtail"] uops_per_byte=0.22265625
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3d6d040000           leaq 0x46d(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: 488bc1                   movq %rcx, %rax
    25: 4881e800020000           subq $0x200, %rax
    32: 0f827d030000             jb 0x3a3
    38: 666666666666660f1f840000000000 nopw (%rax, %rax)
    53: 6666660f1f840000000000   nopw (%rax, %rax)
    64: 4d8b5020                 movq 0x20(%r8), %r10
    68: 62917d486f041a           vmovdqa32 (%r10, %r11), %zmm0
    75: 62917d486f4c1a01         vmovdqa32 0x40(%r10, %r11), %zmm1
    83: 62917d486f541a02         vmovdqa32 0x80(%r10, %r11), %zmm2
    91: 62917d486f5c1a03         vmovdqa32 0xc0(%r10, %r11), %zmm3
    99: 62917d486f641a04         vmovdqa32 0x100(%r10, %r11), %zmm4
   107: 62917d486f6c1a05         vmovdqa32 0x140(%r10, %r11), %zmm5
   115: 62917d486f741a06         vmovdqa32 0x180(%r10, %r11), %zmm6
   123: 62917d486f7c1a07         vmovdqa32 0x1c0(%r10, %r11), %zmm7
   131: 4d8b5028                 movq 0x28(%r8), %r10
   135: 62117d486f041a           vmovdqa32 (%r10, %r11), %zmm8
   142: 62117d486f4c1a01         vmovdqa32 0x40(%r10, %r11), %zmm9
   150: 62117d486f541a02         vmovdqa32 0x80(%r10, %r11), %zmm10
   158: 62117d486f5c1a03         vmovdqa32 0xc0(%r10, %r11), %zmm11
   166: 62117d486f641a04         vmovdqa32 0x100(%r10, %r11), %zmm12
   174: 62117d486f6c1a05         vmovdqa32 0x140(%r10, %r11), %zmm13
   182: 62117d486f741a06         vmovdqa32 0x180(%r10, %r11), %zmm14
   190: 62117d486f7c1a07         vmovdqa32 0x1c0(%r10, %r11), %zmm15
   198: 4d8b10                   movq (%r8), %r10
   201: 498b5010                 movq 0x10(%r8), %rdx
   205: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
   212: 62a17d40ef041a           vpxord (%rdx, %r11), %zmm16, %zmm16
   219: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
   225: 62c17d486f17             vmovdqa32 (%r15), %zmm18
   231: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   238: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
   244: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
   252: 62a17d40ef441a01         vpxord 0x40(%rdx, %r11), %zmm16, %zmm16
   260: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
   266: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   273: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
   279: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
   287: 62a17d40ef441a02         vpxord 0x80(%rdx, %r11), %zmm16, %zmm16
   295: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
   301: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   308: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
   314: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
   322: 62a17d40ef441a03         vpxord 0xc0(%rdx, %r11), %zmm16, %zmm16
   330: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
   336: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   343: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
   349: 62817d486f441a04         vmovdqa32 0x100(%r10, %r11), %zmm16
   357: 62a17d40ef441a04         vpxord 0x100(%rdx, %r11), %zmm16, %zmm16
   365: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
   371: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   378: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
   384: 62817d486f441a05         vmovdqa32 0x140(%r10, %r11), %zmm16
   392: 62a17d40ef441a05         vpxord 0x140(%rdx, %r11), %zmm16, %zmm16
   400: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
   406: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   413: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
   419: 62817d486f441a06         vmovdqa32 0x180(%r10, %r11), %zmm16
   427: 62a17d40ef441a06         vpxord 0x180(%rdx, %r11), %zmm16, %zmm16
   435: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
   441: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   448: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
   454: 62817d486f441a07         vmovdqa32 0x1c0(%r10, %r11), %zmm16
   462: 62a17d40ef441a07         vpxord 0x1c0(%rdx, %r11), %zmm16, %zmm16
   470: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
   476: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   483: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
   489: 4d8b5008                 movq 8(%r8), %r10
   493: 498b5018                 movq 0x18(%r8), %rdx
   497: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
   504: 62a17d40ef041a           vpxord (%rdx, %r11), %zmm16, %zmm16
   511: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
   517: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
   524: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   531: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
   537: 62817d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm16
   545: 62a17d40ef441a01         vpxord 0x40(%rdx, %r11), %zmm16, %zmm16
   553: 62b17548efc8             vpxord %zmm16, %zmm1, %zmm1
   559: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   566: 62313548efc9             vpxord %zmm17, %zmm9, %zmm9
   572: 62817d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm16
   580: 62a17d40ef441a02         vpxord 0x80(%rdx, %r11), %zmm16, %zmm16
   588: 62b16d48efd0             vpxord %zmm16, %zmm2, %zmm2
   594: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   601: 62312d48efd1             vpxord %zmm17, %zmm10, %zmm10
   607: 62817d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm16
   615: 62a17d40ef441a03         vpxord 0xc0(%rdx, %r11), %zmm16, %zmm16
   623: 62b16548efd8             vpxord %zmm16, %zmm3, %zmm3
   629: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   636: 62312548efd9             vpxord %zmm17, %zmm11, %zmm11
   642: 62817d486f441a04         vmovdqa32 0x100(%r10, %r11), %zmm16
   650: 62a17d40ef441a04         vpxord 0x100(%rdx, %r11), %zmm16, %zmm16
   658: 62b15d48efe0             vpxord %zmm16, %zmm4, %zmm4
   664: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   671: 62311d48efe1             vpxord %zmm17, %zmm12, %zmm12
   677: 62817d486f441a05         vmovdqa32 0x140(%r10, %r11), %zmm16
   685: 62a17d40ef441a05         vpxord 0x140(%rdx, %r11), %zmm16, %zmm16
   693: 62b15548efe8             vpxord %zmm16, %zmm5, %zmm5
   699: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   706: 62311548efe9             vpxord %zmm17, %zmm13, %zmm13
   712: 62817d486f441a06         vmovdqa32 0x180(%r10, %r11), %zmm16
   720: 62a17d40ef441a06         vpxord 0x180(%rdx, %r11), %zmm16, %zmm16
   728: 62b14d48eff0             vpxord %zmm16, %zmm6, %zmm6
   734: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   741: 62310d48eff1             vpxord %zmm17, %zmm14, %zmm14
   747: 62817d486f441a07         vmovdqa32 0x1c0(%r10, %r11), %zmm16
   755: 62a17d40ef441a07         vpxord 0x1c0(%rdx, %r11), %zmm16, %zmm16
   763: 62b14548eff8             vpxord %zmm16, %zmm7, %zmm7
   769: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
   776: 62310548eff9             vpxord %zmm17, %zmm15, %zmm15
   782: 4d8b11                   movq (%r9), %r10
   785: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   792: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
   800: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
   808: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
   816: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
   824: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
   832: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
   840: 62917d487f7c1a07         vmovdqa32 %zmm7, 0x1c0(%r10, %r11)
   848: 4d8b5108                 movq 8(%r9), %r10
   852: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
   859: 62117d487f4c1a01         vmovdqa32 %zmm9, 0x40(%r10, %r11)
   867: 62117d487f541a02         vmovdqa32 %zmm10, 0x80(%r10, %r11)
   875: 62117d487f5c1a03         vmovdqa32 %zmm11, 0xc0(%r10, %r11)
   883: 62117d487f641a04         vmovdqa32 %zmm12, 0x100(%r10, %r11)
   891: 62117d487f6c1a05         vmovdqa32 %zmm13, 0x140(%r10, %r11)
   899: 62117d487f741a06         vmovdqa32 %zmm14, 0x180(%r10, %r11)
   907: 62117d487f7c1a07         vmovdqa32 %zmm15, 0x1c0(%r10, %r11)
   915: 4981c300020000           addq $0x200, %r11
   922: 4c3bd8                   cmpq %rax, %r11
   925: 0f869dfcffff             jbe 0x40
   931: 4c3bd9                   cmpq %rcx, %r11
   934: 0f83c5000000             jae 0x471
   940: 4c8bd1                   movq %rcx, %r10
   943: 4d2bd3                   subq %r11, %r10
   946: 498bd2                   movq %r10, %rdx
   949: 4883fa40                 cmpq $0x40, %rdx
   953: 0f8607000000             jbe 0x3c6
   959: 48c7c240000000           movq $0x40, %rdx
   966: c4c1f8908cd780000000     kmovq 0x80(%r15, %rdx, 8), %k1
   976: 498b4020                 movq 0x20(%r8), %rax
   980: 62b17fc96f0418           vmovdqu8 (%rax, %r11), %zmm0 {%k1} {z}
   987: 498b4028                 movq 0x28(%r8), %rax
   991: 62317fc96f0418           vmovdqu8 (%rax, %r11), %zmm8 {%k1} {z}
   998: 498b00                   movq (%r8), %rax
  1001: 62a17fc96f0418           vmovdqu8 (%rax, %r11), %zmm16 {%k1} {z}
  1008: 498b4010                 movq 0x10(%r8), %rax
  1012: 62a17fc96f0c18           vmovdqu8 (%rax, %r11), %zmm17 {%k1} {z}
  1019: 62a17d40efc1             vpxord %zmm17, %zmm16, %zmm16
  1025: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
  1031: 62c17d486f17             vmovdqa32 (%r15), %zmm18
  1037: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1044: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
  1050: 498b4008                 movq 8(%r8), %rax
  1054: 62a17fc96f0418           vmovdqu8 (%rax, %r11), %zmm16 {%k1} {z}
  1061: 498b4018                 movq 0x18(%r8), %rax
  1065: 62a17fc96f0c18           vmovdqu8 (%rax, %r11), %zmm17 {%k1} {z}
  1072: 62a17d40efc1             vpxord %zmm17, %zmm16, %zmm16
  1078: 62b17d48efc0             vpxord %zmm16, %zmm0, %zmm0
  1084: 62c17d486f5701           vmovdqa32 0x40(%r15), %zmm18
  1091: 62a3fd40ceca00           vgf2p8affineqb $0, %zmm18, %zmm16, %zmm17
  1098: 62313d48efc1             vpxord %zmm17, %zmm8, %zmm8
  1104: 498b01                   movq (%r9), %rax
  1107: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
  1114: 498b4108                 movq 8(%r9), %rax
  1118: 62317f497f0418           vmovdqu8 %zmm8, (%rax, %r11) {%k1}
  1125: 4c03da                   addq %rdx, %r11
  1128: 4c2bd2                   subq %rdx, %r10
  1131: 0f8541ffffff             jne 0x3b2
  1137: 415f                     popq %r15
  1139: 5b                       popq %rbx
  1140: c5f877                   vzeroupper 
  1143: c3                       retq 
  1144: 0f0b                     ud2 
  1146: 0f0b                     ud2 
  1148: 0f0b                     ud2 
  1150: 0f0b                     ud2 
