; config_t1_654 target=sse twin=notail topo=rs(6,3)/enc op=encode U=1
; code_size=832 twins=["notail"] uops_per_byte=7.6875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d27030000           leaq 0x327(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 66450f6f0f               movdqa (%r15), %xmm9
    52: 488bc1                   movq %rcx, %rax
    55: 4883e810                 subq $0x10, %rax
    59: 0f82cb020000             jb 0x30c
    65: 666666666666660f1f840000000000 nopw (%rax, %rax)
    80: 666666666666660f1f840000000000 nopw (%rax, %rax)
    95: 90                       nop 
    96: 66420f6f1c1e             movdqa (%rsi, %r11), %xmm3
   102: 660f6fe3                 movdqa %xmm3, %xmm4
   106: 66410fdbe1               pand %xmm9, %xmm4
   111: 660f6feb                 movdqa %xmm3, %xmm5
   115: 660f71d504               psrlw $4, %xmm5
   120: 66410fdbe9               pand %xmm9, %xmm5
   125: 660f6fc3                 movdqa %xmm3, %xmm0
   129: 66450f6f4740             movdqa 0x40(%r15), %xmm8
   135: 66410f6ff0               movdqa %xmm8, %xmm6
   140: 660f3800f4               pshufb %xmm4, %xmm6
   145: 66450f6f8780000000       movdqa 0x80(%r15), %xmm8
   154: 66410f6ff8               movdqa %xmm8, %xmm7
   159: 660f3800fd               pshufb %xmm5, %xmm7
   164: 660f6fce                 movdqa %xmm6, %xmm1
   168: 660fefcf                 pxor %xmm7, %xmm1
   172: 66450f6f87c0000000       movdqa 0xc0(%r15), %xmm8
   181: 66410f6ff0               movdqa %xmm8, %xmm6
   186: 660f3800f4               pshufb %xmm4, %xmm6
   191: 66450f6f8700010000       movdqa 0x100(%r15), %xmm8
   200: 66410f6ff8               movdqa %xmm8, %xmm7
   205: 660f3800fd               pshufb %xmm5, %xmm7
   210: 660f6fd6                 movdqa %xmm6, %xmm2
   214: 660fefd7                 pxor %xmm7, %xmm2
   218: 66420f6f1c1f             movdqa (%rdi, %r11), %xmm3
   224: 660f6fe3                 movdqa %xmm3, %xmm4
   228: 66410fdbe1               pand %xmm9, %xmm4
   233: 660f6feb                 movdqa %xmm3, %xmm5
   237: 660f71d504               psrlw $4, %xmm5
   242: 66410fdbe9               pand %xmm9, %xmm5
   247: 660fefc3                 pxor %xmm3, %xmm0
   251: 66450f6f87c0000000       movdqa 0xc0(%r15), %xmm8
   260: 66410f6ff0               movdqa %xmm8, %xmm6
   265: 660f3800f4               pshufb %xmm4, %xmm6
   270: 66450f6f8700010000       movdqa 0x100(%r15), %xmm8
   279: 66410f6ff8               movdqa %xmm8, %xmm7
   284: 660f3800fd               pshufb %xmm5, %xmm7
   289: 660fefce                 pxor %xmm6, %xmm1
   293: 660fefcf                 pxor %xmm7, %xmm1
   297: 66450f6f8740010000       movdqa 0x140(%r15), %xmm8
   306: 66410f6ff0               movdqa %xmm8, %xmm6
   311: 660f3800f4               pshufb %xmm4, %xmm6
   316: 66450f6f8780010000       movdqa 0x180(%r15), %xmm8
   325: 66410f6ff8               movdqa %xmm8, %xmm7
   330: 660f3800fd               pshufb %xmm5, %xmm7
   335: 660fefd6                 pxor %xmm6, %xmm2
   339: 660fefd7                 pxor %xmm7, %xmm2
   343: 66430f6f1c1c             movdqa (%r12, %r11), %xmm3
   349: 660f6fe3                 movdqa %xmm3, %xmm4
   353: 66410fdbe1               pand %xmm9, %xmm4
   358: 660f6feb                 movdqa %xmm3, %xmm5
   362: 660f71d504               psrlw $4, %xmm5
   367: 66410fdbe9               pand %xmm9, %xmm5
   372: 660fefc3                 pxor %xmm3, %xmm0
   376: 66450f6f8740010000       movdqa 0x140(%r15), %xmm8
   385: 66410f6ff0               movdqa %xmm8, %xmm6
   390: 660f3800f4               pshufb %xmm4, %xmm6
   395: 66450f6f8780010000       movdqa 0x180(%r15), %xmm8
   404: 66410f6ff8               movdqa %xmm8, %xmm7
   409: 660f3800fd               pshufb %xmm5, %xmm7
   414: 660fefce                 pxor %xmm6, %xmm1
   418: 660fefcf                 pxor %xmm7, %xmm1
   422: 66450f6f87c0010000       movdqa 0x1c0(%r15), %xmm8
   431: 66410f6ff0               movdqa %xmm8, %xmm6
   436: 660f3800f4               pshufb %xmm4, %xmm6
   441: 66450f6f8700020000       movdqa 0x200(%r15), %xmm8
   450: 66410f6ff8               movdqa %xmm8, %xmm7
   455: 660f3800fd               pshufb %xmm5, %xmm7
   460: 660fefd6                 pxor %xmm6, %xmm2
   464: 660fefd7                 pxor %xmm7, %xmm2
   468: 66430f6f5c1d00           movdqa (%r13, %r11), %xmm3
   475: 660f6fe3                 movdqa %xmm3, %xmm4
   479: 66410fdbe1               pand %xmm9, %xmm4
   484: 660f6feb                 movdqa %xmm3, %xmm5
   488: 660f71d504               psrlw $4, %xmm5
   493: 66410fdbe9               pand %xmm9, %xmm5
   498: 660fefc3                 pxor %xmm3, %xmm0
   502: 66450f6f8740020000       movdqa 0x240(%r15), %xmm8
   511: 66410f6ff0               movdqa %xmm8, %xmm6
   516: 660f3800f4               pshufb %xmm4, %xmm6
   521: 66450f6f8780020000       movdqa 0x280(%r15), %xmm8
   530: 66410f6ff8               movdqa %xmm8, %xmm7
   535: 660f3800fd               pshufb %xmm5, %xmm7
   540: 660fefce                 pxor %xmm6, %xmm1
   544: 660fefcf                 pxor %xmm7, %xmm1
   548: 66450f6f4740             movdqa 0x40(%r15), %xmm8
   554: 66410f6ff0               movdqa %xmm8, %xmm6
   559: 660f3800f4               pshufb %xmm4, %xmm6
   564: 66450f6f8780000000       movdqa 0x80(%r15), %xmm8
   573: 66410f6ff8               movdqa %xmm8, %xmm7
   578: 660f3800fd               pshufb %xmm5, %xmm7
   583: 660fefd6                 pxor %xmm6, %xmm2
   587: 660fefd7                 pxor %xmm7, %xmm2
   591: 66430f6f1c1e             movdqa (%r14, %r11), %xmm3
   597: 660f6fe3                 movdqa %xmm3, %xmm4
   601: 66410fdbe1               pand %xmm9, %xmm4
   606: 660f6feb                 movdqa %xmm3, %xmm5
   610: 660f71d504               psrlw $4, %xmm5
   615: 66410fdbe9               pand %xmm9, %xmm5
   620: 660fefc3                 pxor %xmm3, %xmm0
   624: 66450f6f87c0010000       movdqa 0x1c0(%r15), %xmm8
   633: 66410f6ff0               movdqa %xmm8, %xmm6
   638: 660f3800f4               pshufb %xmm4, %xmm6
   643: 66450f6f8700020000       movdqa 0x200(%r15), %xmm8
   652: 66410f6ff8               movdqa %xmm8, %xmm7
   657: 660f3800fd               pshufb %xmm5, %xmm7
   662: 660fefce                 pxor %xmm6, %xmm1
   666: 660fefcf                 pxor %xmm7, %xmm1
   670: 66450f6f8740020000       movdqa 0x240(%r15), %xmm8
   679: 66410f6ff0               movdqa %xmm8, %xmm6
   684: 660f3800f4               pshufb %xmm4, %xmm6
   689: 66450f6f8780020000       movdqa 0x280(%r15), %xmm8
   698: 66410f6ff8               movdqa %xmm8, %xmm7
   703: 660f3800fd               pshufb %xmm5, %xmm7
   708: 660fefd6                 pxor %xmm6, %xmm2
   712: 660fefd7                 pxor %xmm7, %xmm2
   716: 4d8b5028                 movq 0x28(%r8), %r10
   720: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   726: 660fefc3                 pxor %xmm3, %xmm0
   730: 660fefcb                 pxor %xmm3, %xmm1
   734: 660fefd3                 pxor %xmm3, %xmm2
   738: 4d8b11                   movq (%r9), %r10
   741: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   747: 4d8b5108                 movq 8(%r9), %r10
   751: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
   757: 4d8b5110                 movq 0x10(%r9), %r10
   761: 66430f7f141a             movdqa %xmm2, (%r10, %r11)
   767: 4983c310                 addq $0x10, %r11
   771: 4c3bd8                   cmpq %rax, %r11
   774: 0f8654fdffff             jbe 0x60
   780: 415e                     popq %r14
   782: 415d                     popq %r13
   784: 415c                     popq %r12
   786: 415f                     popq %r15
   788: 5b                       popq %rbx
   789: c3                       retq 
   790: 0f0b                     ud2 
   792: 0f0b                     ud2 
   794: 0f0b                     ud2 
   796: 0f0b                     ud2 
   798: 0f0b                     ud2 
   800: 0f0b                     ud2 
   802: 0f0b                     ud2 
   804: 0f0b                     ud2 
   806: 0f0b                     ud2 
   808: 0f0b                     ud2 
   810: 0f0b                     ud2 
   812: 0f0b                     ud2 
   814: 0f0b                     ud2 
   816: 0f0b                     ud2 
   818: 0f0b                     ud2 
   820: 0f0b                     ud2 
   822: 0f0b                     ud2 
   824: 0f0b                     ud2 
   826: 0f0b                     ud2 
   828: 0f0b                     ud2 
   830: 0f0b                     ud2 
