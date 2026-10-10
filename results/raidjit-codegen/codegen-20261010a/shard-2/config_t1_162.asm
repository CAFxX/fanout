; config_t1_162 target=avx512-gfni twin=maskedtail topo=rs(6,3)/rec{0} op=recover U=1
; code_size=320 twins=["maskedtail"] uops_per_byte=0.125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3d27010000           leaq 0x127(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d8b7020                 movq 0x20(%r8), %r14
    44: 4d33db                   xorq %r11, %r11
    47: 488bc1                   movq %rcx, %rax
    50: 4883e840                 subq $0x40, %rax
    54: 0f824a000000             jb 0x86
    60: 0f1f4000                 nopl (%rax)
    64: 62b17d486f041e           vmovdqa32 (%rsi, %r11), %zmm0
    71: 62b17d48ef041f           vpxord (%rdi, %r11), %zmm0, %zmm0
    78: 62917d48ef041c           vpxord (%r12, %r11), %zmm0, %zmm0
    85: 62917d48ef441d00         vpxord (%r13, %r11), %zmm0, %zmm0
    93: 62917d48ef041e           vpxord (%r14, %r11), %zmm0, %zmm0
   100: 4d8b5028                 movq 0x28(%r8), %r10
   104: 62917d48ef041a           vpxord (%r10, %r11), %zmm0, %zmm0
   111: 4d8b11                   movq (%r9), %r10
   114: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   121: 4983c340                 addq $0x40, %r11
   125: 4c3bd8                   cmpq %rax, %r11
   128: 0f86baffffff             jbe 0x40
   134: 4c3bd9                   cmpq %rcx, %r11
   137: 0f838f000000             jae 0x11e
   143: 4c8bd1                   movq %rcx, %r10
   146: 4d2bd3                   subq %r11, %r10
   149: 498bd2                   movq %r10, %rdx
   152: 4883fa40                 cmpq $0x40, %rdx
   156: 0f8607000000             jbe 0xa9
   162: 48c7c240000000           movq $0x40, %rdx
   169: c4c1f8900cd7             kmovq (%r15, %rdx, 8), %k1
   175: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
   181: 62b17fc96f0c1e           vmovdqu8 (%rsi, %r11), %zmm1 {%k1} {z}
   188: 62f17d48efc1             vpxord %zmm1, %zmm0, %zmm0
   194: 62b17fc96f0c1f           vmovdqu8 (%rdi, %r11), %zmm1 {%k1} {z}
   201: 62f17d48efc1             vpxord %zmm1, %zmm0, %zmm0
   207: 62917fc96f0c1c           vmovdqu8 (%r12, %r11), %zmm1 {%k1} {z}
   214: 62f17d48efc1             vpxord %zmm1, %zmm0, %zmm0
   220: 62917fc96f4c1d00         vmovdqu8 (%r13, %r11), %zmm1 {%k1} {z}
   228: 62f17d48efc1             vpxord %zmm1, %zmm0, %zmm0
   234: 62917fc96f0c1e           vmovdqu8 (%r14, %r11), %zmm1 {%k1} {z}
   241: 62f17d48efc1             vpxord %zmm1, %zmm0, %zmm0
   247: 498b4028                 movq 0x28(%r8), %rax
   251: 62b17fc96f0c18           vmovdqu8 (%rax, %r11), %zmm1 {%k1} {z}
   258: 62f17d48efc1             vpxord %zmm1, %zmm0, %zmm0
   264: 498b01                   movq (%r9), %rax
   267: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
   274: 4c03da                   addq %rdx, %r11
   277: 4c2bd2                   subq %rdx, %r10
   280: 0f8577ffffff             jne 0x95
   286: 415e                     popq %r14
   288: 415d                     popq %r13
   290: 415c                     popq %r12
   292: 415f                     popq %r15
   294: 5b                       popq %rbx
   295: c5f877                   vzeroupper 
   298: c3                       retq 
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
