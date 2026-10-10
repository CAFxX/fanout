; config_t1_251 target=avx512-gfni twin=notail topo=rs(6,3)/upd{0} op=parity_update U=8
; code_size=1024 twins=["notail"] uops_per_byte=0.234375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3ded030000           leaq 0x3ed(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: 488bc1                   movq %rcx, %rax
    25: 4881e800020000           subq $0x200, %rax
    32: 0f82a8030000             jb 0x3ce
    38: 666666666666660f1f840000000000 nopw (%rax, %rax)
    53: 6666660f1f840000000000   nopw (%rax, %rax)
    64: 4d8b5010                 movq 0x10(%r8), %r10
    68: 62917d486f041a           vmovdqa32 (%r10, %r11), %zmm0
    75: 62917d486f4c1a01         vmovdqa32 0x40(%r10, %r11), %zmm1
    83: 62917d486f541a02         vmovdqa32 0x80(%r10, %r11), %zmm2
    91: 62917d486f5c1a03         vmovdqa32 0xc0(%r10, %r11), %zmm3
    99: 62917d486f641a04         vmovdqa32 0x100(%r10, %r11), %zmm4
   107: 62917d486f6c1a05         vmovdqa32 0x140(%r10, %r11), %zmm5
   115: 62917d486f741a06         vmovdqa32 0x180(%r10, %r11), %zmm6
   123: 62917d486f7c1a07         vmovdqa32 0x1c0(%r10, %r11), %zmm7
   131: 4d8b5018                 movq 0x18(%r8), %r10
   135: 62117d486f041a           vmovdqa32 (%r10, %r11), %zmm8
   142: 62117d486f4c1a01         vmovdqa32 0x40(%r10, %r11), %zmm9
   150: 62117d486f541a02         vmovdqa32 0x80(%r10, %r11), %zmm10
   158: 62117d486f5c1a03         vmovdqa32 0xc0(%r10, %r11), %zmm11
   166: 62117d486f641a04         vmovdqa32 0x100(%r10, %r11), %zmm12
   174: 62117d486f6c1a05         vmovdqa32 0x140(%r10, %r11), %zmm13
   182: 62117d486f741a06         vmovdqa32 0x180(%r10, %r11), %zmm14
   190: 62117d486f7c1a07         vmovdqa32 0x1c0(%r10, %r11), %zmm15
   198: 4d8b5020                 movq 0x20(%r8), %r10
   202: 62817d486f041a           vmovdqa32 (%r10, %r11), %zmm16
   209: 62817d486f4c1a01         vmovdqa32 0x40(%r10, %r11), %zmm17
   217: 62817d486f541a02         vmovdqa32 0x80(%r10, %r11), %zmm18
   225: 62817d486f5c1a03         vmovdqa32 0xc0(%r10, %r11), %zmm19
   233: 62817d486f641a04         vmovdqa32 0x100(%r10, %r11), %zmm20
   241: 62817d486f6c1a05         vmovdqa32 0x140(%r10, %r11), %zmm21
   249: 62817d486f741a06         vmovdqa32 0x180(%r10, %r11), %zmm22
   257: 62817d486f7c1a07         vmovdqa32 0x1c0(%r10, %r11), %zmm23
   265: 4d8b10                   movq (%r8), %r10
   268: 498b5008                 movq 8(%r8), %rdx
   272: 62017d486f041a           vmovdqa32 (%r10, %r11), %zmm24
   279: 62213d40ef041a           vpxord (%rdx, %r11), %zmm24, %zmm24
   286: 62917d48efc0             vpxord %zmm24, %zmm0, %zmm0
   292: 62417d486f17             vmovdqa32 (%r15), %zmm26
   298: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   305: 62113d48efc1             vpxord %zmm25, %zmm8, %zmm8
   311: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   318: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   325: 62817d40efc1             vpxord %zmm25, %zmm16, %zmm16
   331: 62017d486f441a01         vmovdqa32 0x40(%r10, %r11), %zmm24
   339: 62213d40ef441a01         vpxord 0x40(%rdx, %r11), %zmm24, %zmm24
   347: 62917548efc8             vpxord %zmm24, %zmm1, %zmm1
   353: 62417d486f17             vmovdqa32 (%r15), %zmm26
   359: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   366: 62113548efc9             vpxord %zmm25, %zmm9, %zmm9
   372: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   379: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   386: 62817540efc9             vpxord %zmm25, %zmm17, %zmm17
   392: 62017d486f441a02         vmovdqa32 0x80(%r10, %r11), %zmm24
   400: 62213d40ef441a02         vpxord 0x80(%rdx, %r11), %zmm24, %zmm24
   408: 62916d48efd0             vpxord %zmm24, %zmm2, %zmm2
   414: 62417d486f17             vmovdqa32 (%r15), %zmm26
   420: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   427: 62112d48efd1             vpxord %zmm25, %zmm10, %zmm10
   433: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   440: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   447: 62816d40efd1             vpxord %zmm25, %zmm18, %zmm18
   453: 62017d486f441a03         vmovdqa32 0xc0(%r10, %r11), %zmm24
   461: 62213d40ef441a03         vpxord 0xc0(%rdx, %r11), %zmm24, %zmm24
   469: 62916548efd8             vpxord %zmm24, %zmm3, %zmm3
   475: 62417d486f17             vmovdqa32 (%r15), %zmm26
   481: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   488: 62112548efd9             vpxord %zmm25, %zmm11, %zmm11
   494: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   501: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   508: 62816540efd9             vpxord %zmm25, %zmm19, %zmm19
   514: 62017d486f441a04         vmovdqa32 0x100(%r10, %r11), %zmm24
   522: 62213d40ef441a04         vpxord 0x100(%rdx, %r11), %zmm24, %zmm24
   530: 62915d48efe0             vpxord %zmm24, %zmm4, %zmm4
   536: 62417d486f17             vmovdqa32 (%r15), %zmm26
   542: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   549: 62111d48efe1             vpxord %zmm25, %zmm12, %zmm12
   555: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   562: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   569: 62815d40efe1             vpxord %zmm25, %zmm20, %zmm20
   575: 62017d486f441a05         vmovdqa32 0x140(%r10, %r11), %zmm24
   583: 62213d40ef441a05         vpxord 0x140(%rdx, %r11), %zmm24, %zmm24
   591: 62915548efe8             vpxord %zmm24, %zmm5, %zmm5
   597: 62417d486f17             vmovdqa32 (%r15), %zmm26
   603: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   610: 62111548efe9             vpxord %zmm25, %zmm13, %zmm13
   616: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   623: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   630: 62815540efe9             vpxord %zmm25, %zmm21, %zmm21
   636: 62017d486f441a06         vmovdqa32 0x180(%r10, %r11), %zmm24
   644: 62213d40ef441a06         vpxord 0x180(%rdx, %r11), %zmm24, %zmm24
   652: 62914d48eff0             vpxord %zmm24, %zmm6, %zmm6
   658: 62417d486f17             vmovdqa32 (%r15), %zmm26
   664: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   671: 62110d48eff1             vpxord %zmm25, %zmm14, %zmm14
   677: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   684: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   691: 62814d40eff1             vpxord %zmm25, %zmm22, %zmm22
   697: 62017d486f441a07         vmovdqa32 0x1c0(%r10, %r11), %zmm24
   705: 62213d40ef441a07         vpxord 0x1c0(%rdx, %r11), %zmm24, %zmm24
   713: 62914548eff8             vpxord %zmm24, %zmm7, %zmm7
   719: 62417d486f17             vmovdqa32 (%r15), %zmm26
   725: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   732: 62110548eff9             vpxord %zmm25, %zmm15, %zmm15
   738: 62417d486f5701           vmovdqa32 0x40(%r15), %zmm26
   745: 6203bd40ceca00           vgf2p8affineqb $0, %zmm26, %zmm24, %zmm25
   752: 62814540eff9             vpxord %zmm25, %zmm23, %zmm23
   758: 4d8b11                   movq (%r9), %r10
   761: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   768: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
   776: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
   784: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
   792: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
   800: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
   808: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
   816: 62917d487f7c1a07         vmovdqa32 %zmm7, 0x1c0(%r10, %r11)
   824: 4d8b5108                 movq 8(%r9), %r10
   828: 62117d487f041a           vmovdqa32 %zmm8, (%r10, %r11)
   835: 62117d487f4c1a01         vmovdqa32 %zmm9, 0x40(%r10, %r11)
   843: 62117d487f541a02         vmovdqa32 %zmm10, 0x80(%r10, %r11)
   851: 62117d487f5c1a03         vmovdqa32 %zmm11, 0xc0(%r10, %r11)
   859: 62117d487f641a04         vmovdqa32 %zmm12, 0x100(%r10, %r11)
   867: 62117d487f6c1a05         vmovdqa32 %zmm13, 0x140(%r10, %r11)
   875: 62117d487f741a06         vmovdqa32 %zmm14, 0x180(%r10, %r11)
   883: 62117d487f7c1a07         vmovdqa32 %zmm15, 0x1c0(%r10, %r11)
   891: 4d8b5110                 movq 0x10(%r9), %r10
   895: 62817d487f041a           vmovdqa32 %zmm16, (%r10, %r11)
   902: 62817d487f4c1a01         vmovdqa32 %zmm17, 0x40(%r10, %r11)
   910: 62817d487f541a02         vmovdqa32 %zmm18, 0x80(%r10, %r11)
   918: 62817d487f5c1a03         vmovdqa32 %zmm19, 0xc0(%r10, %r11)
   926: 62817d487f641a04         vmovdqa32 %zmm20, 0x100(%r10, %r11)
   934: 62817d487f6c1a05         vmovdqa32 %zmm21, 0x140(%r10, %r11)
   942: 62817d487f741a06         vmovdqa32 %zmm22, 0x180(%r10, %r11)
   950: 62817d487f7c1a07         vmovdqa32 %zmm23, 0x1c0(%r10, %r11)
   958: 4981c300020000           addq $0x200, %r11
   965: 4c3bd8                   cmpq %rax, %r11
   968: 0f8672fcffff             jbe 0x40
   974: 415f                     popq %r15
   976: 5b                       popq %rbx
   977: c5f877                   vzeroupper 
   980: c3                       retq 
   981: 0f0b                     ud2 
   983: 0f0b                     ud2 
   985: 0f0b                     ud2 
   987: 0f0b                     ud2 
   989: 0f0b                     ud2 
   991: 0f0b                     ud2 
   993: 0f0b                     ud2 
   995: 0f0b                     ud2 
   997: 0f0b                     ud2 
   999: 0f0b                     ud2 
  1001: 0f0b                     ud2 
  1003: 0f0b                     ud2 
  1005: 0f0b                     ud2 
  1007: 0f0b                     ud2 
  1009: 0f0b                     ud2 
  1011: 0f0b                     ud2 
  1013: 0f0b                     ud2 
  1015: 0f0b                     ud2 
  1017: 0f0b                     ud2 
  1019: 0f0b                     ud2 
  1021: 0f0b                     ud2 
  1023: cc                       int3 
