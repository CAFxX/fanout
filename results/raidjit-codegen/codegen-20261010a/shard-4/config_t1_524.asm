; config_t1_524 target=avx2 twin=notail topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=256 twins=["notail"] uops_per_byte=0.75
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3ded000000           leaq 0xed(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: c4417d6f0f               vmovdqa (%r15), %ymm9
    27: 488bc1                   movq %rcx, %rax
    30: 4883e820                 subq $0x20, %rax
    34: 0f82cb000000             jb 0xf3
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5010                 movq 0x10(%r8), %r10
    68: c4817d6f041a             vmovdqa (%r10, %r11), %ymm0
    74: 4d8b5018                 movq 0x18(%r8), %r10
    78: c4817d6f0c1a             vmovdqa (%r10, %r11), %ymm1
    84: 4d8b5020                 movq 0x20(%r8), %r10
    88: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
    94: 4d8b10                   movq (%r8), %r10
    97: 498b5008                 movq 8(%r8), %rdx
   101: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   107: c4a165ef1c1a             vpxor (%rdx, %r11), %ymm3, %ymm3
   113: c4c165dbe1               vpand %ymm9, %ymm3, %ymm4
   118: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   123: c4c155dbe9               vpand %ymm9, %ymm5, %ymm5
   128: c5fdefc3                 vpxor %ymm3, %ymm0, %ymm0
   132: c4417d6f4740             vmovdqa 0x40(%r15), %ymm8
   138: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   143: c4417d6f8780000000       vmovdqa 0x80(%r15), %ymm8
   152: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   157: c5f5efce                 vpxor %ymm6, %ymm1, %ymm1
   161: c5f5efcf                 vpxor %ymm7, %ymm1, %ymm1
   165: c4417d6f87c0000000       vmovdqa 0xc0(%r15), %ymm8
   174: c4e23d00f4               vpshufb %ymm4, %ymm8, %ymm6
   179: c4417d6f8700010000       vmovdqa 0x100(%r15), %ymm8
   188: c4e23d00fd               vpshufb %ymm5, %ymm8, %ymm7
   193: c5edefd6                 vpxor %ymm6, %ymm2, %ymm2
   197: c5edefd7                 vpxor %ymm7, %ymm2, %ymm2
   201: 4d8b11                   movq (%r9), %r10
   204: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   210: 4d8b5108                 movq 8(%r9), %r10
   214: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
   220: 4d8b5110                 movq 0x10(%r9), %r10
   224: c4817d7f141a             vmovdqa %ymm2, (%r10, %r11)
   230: 4983c320                 addq $0x20, %r11
   234: 4c3bd8                   cmpq %rax, %r11
   237: 0f864dffffff             jbe 0x40
   243: 415f                     popq %r15
   245: 5b                       popq %rbx
   246: c5f877                   vzeroupper 
   249: c3                       retq 
   250: 0f0b                     ud2 
   252: 0f0b                     ud2 
   254: 0f0b                     ud2 
