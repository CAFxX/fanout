; config_t1_810 target=sse twin=notail topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=320 twins=["notail"] uops_per_byte=1.9375
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3d2d010000           leaq 0x12d(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: 66450f6f0f               movdqa (%r15), %xmm9
    27: 488bc1                   movq %rcx, %rax
    30: 4883e810                 subq $0x10, %rax
    34: 0f82eb000000             jb 0x113
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5010                 movq 0x10(%r8), %r10
    68: 66430f6f041a             movdqa (%r10, %r11), %xmm0
    74: 4d8b5018                 movq 0x18(%r8), %r10
    78: 66430f6f0c1a             movdqa (%r10, %r11), %xmm1
    84: 4d8b5020                 movq 0x20(%r8), %r10
    88: 66430f6f141a             movdqa (%r10, %r11), %xmm2
    94: 4d8b10                   movq (%r8), %r10
    97: 498b5008                 movq 8(%r8), %rdx
   101: 66430f6f1c1a             movdqa (%r10, %r11), %xmm3
   107: 66420f6f341a             movdqa (%rdx, %r11), %xmm6
   113: 660fefde                 pxor %xmm6, %xmm3
   117: 660f6fe3                 movdqa %xmm3, %xmm4
   121: 66410fdbe1               pand %xmm9, %xmm4
   126: 660f6feb                 movdqa %xmm3, %xmm5
   130: 660f71d504               psrlw $4, %xmm5
   135: 66410fdbe9               pand %xmm9, %xmm5
   140: 660fefc3                 pxor %xmm3, %xmm0
   144: 66450f6f4740             movdqa 0x40(%r15), %xmm8
   150: 66410f6ff0               movdqa %xmm8, %xmm6
   155: 660f3800f4               pshufb %xmm4, %xmm6
   160: 66450f6f8780000000       movdqa 0x80(%r15), %xmm8
   169: 66410f6ff8               movdqa %xmm8, %xmm7
   174: 660f3800fd               pshufb %xmm5, %xmm7
   179: 660fefce                 pxor %xmm6, %xmm1
   183: 660fefcf                 pxor %xmm7, %xmm1
   187: 66450f6f87c0000000       movdqa 0xc0(%r15), %xmm8
   196: 66410f6ff0               movdqa %xmm8, %xmm6
   201: 660f3800f4               pshufb %xmm4, %xmm6
   206: 66450f6f8700010000       movdqa 0x100(%r15), %xmm8
   215: 66410f6ff8               movdqa %xmm8, %xmm7
   220: 660f3800fd               pshufb %xmm5, %xmm7
   225: 660fefd6                 pxor %xmm6, %xmm2
   229: 660fefd7                 pxor %xmm7, %xmm2
   233: 4d8b11                   movq (%r9), %r10
   236: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   242: 4d8b5108                 movq 8(%r9), %r10
   246: 66430f7f0c1a             movdqa %xmm1, (%r10, %r11)
   252: 4d8b5110                 movq 0x10(%r9), %r10
   256: 66430f7f141a             movdqa %xmm2, (%r10, %r11)
   262: 4983c310                 addq $0x10, %r11
   266: 4c3bd8                   cmpq %rax, %r11
   269: 0f862dffffff             jbe 0x40
   275: 415f                     popq %r15
   277: 5b                       popq %rbx
   278: c3                       retq 
   279: 0f0b                     ud2 
   281: 0f0b                     ud2 
   283: 0f0b                     ud2 
   285: 0f0b                     ud2 
   287: 0f0b                     ud2 
   289: 0f0b                     ud2 
   291: 0f0b                     ud2 
   293: 0f0b                     ud2 
   295: 0f0b                     ud2 
   297: 0f0b                     ud2 
   299: 0f0b                     ud2 
   301: 0f0b                     ud2 
   303: 0f0b                     ud2 
   305: 0f0b                     ud2 
   307: 0f0b                     ud2 
   309: 0f0b                     ud2 
   311: 0f0b                     ud2 
   313: 0f0b                     ud2 
   315: 0f0b                     ud2 
   317: 0f0b                     ud2 
   319: cc                       int3 
