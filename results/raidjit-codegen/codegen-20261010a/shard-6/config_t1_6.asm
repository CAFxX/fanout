; config_t1_6 target=avx512-gfni twin=maskedtail topo=raid5(4)/enc op=encode U=1
; code_size=256 twins=["maskedtail"] uops_per_byte=0.078125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4154                     pushq %r12
     5: 4155                     pushq %r13
     7: 4156                     pushq %r14
     9: 4c8bc7                   movq %rdi, %r8
    12: 4c8bce                   movq %rsi, %r9
    15: 488bca                   movq %rdx, %rcx
    18: 4c8d3de7000000           leaq 0xe7(%rip), %r15
    25: 498b30                   movq (%r8), %rsi
    28: 498b7808                 movq 8(%r8), %rdi
    32: 4d8b6010                 movq 0x10(%r8), %r12
    36: 4d8b6818                 movq 0x18(%r8), %r13
    40: 4d33db                   xorq %r11, %r11
    43: 488bc1                   movq %rcx, %rax
    46: 4883e840                 subq $0x40, %rax
    50: 0f823c000000             jb 0x74
    56: 0f1f840000000000         nopl (%rax, %rax)
    64: 62b17d486f041e           vmovdqa32 (%rsi, %r11), %zmm0
    71: 62b17d48ef041f           vpxord (%rdi, %r11), %zmm0, %zmm0
    78: 62917d48ef041c           vpxord (%r12, %r11), %zmm0, %zmm0
    85: 62917d48ef441d00         vpxord (%r13, %r11), %zmm0, %zmm0
    93: 4d8b11                   movq (%r9), %r10
    96: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   103: 4983c340                 addq $0x40, %r11
   107: 4c3bd8                   cmpq %rax, %r11
   110: 0f86ccffffff             jbe 0x40
   116: 4c3bd9                   cmpq %rcx, %r11
   119: 0f8371000000             jae 0xee
   125: 4c8bd1                   movq %rcx, %r10
   128: 4d2bd3                   subq %r11, %r10
   131: 498bd2                   movq %r10, %rdx
   134: 4883fa40                 cmpq $0x40, %rdx
   138: 0f8607000000             jbe 0x97
   144: 48c7c240000000           movq $0x40, %rdx
   151: c4c1f8900cd7             kmovq (%r15, %rdx, 8), %k1
   157: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
   163: 62b17fc96f0c1e           vmovdqu8 (%rsi, %r11), %zmm1 {%k1} {z}
   170: 62f17d48efc1             vpxord %zmm1, %zmm0, %zmm0
   176: 62b17fc96f0c1f           vmovdqu8 (%rdi, %r11), %zmm1 {%k1} {z}
   183: 62f17d48efc1             vpxord %zmm1, %zmm0, %zmm0
   189: 62917fc96f0c1c           vmovdqu8 (%r12, %r11), %zmm1 {%k1} {z}
   196: 62f17d48efc1             vpxord %zmm1, %zmm0, %zmm0
   202: 62917fc96f4c1d00         vmovdqu8 (%r13, %r11), %zmm1 {%k1} {z}
   210: 62f17d48efc1             vpxord %zmm1, %zmm0, %zmm0
   216: 498b01                   movq (%r9), %rax
   219: 62b17f497f0418           vmovdqu8 %zmm0, (%rax, %r11) {%k1}
   226: 4c03da                   addq %rdx, %r11
   229: 4c2bd2                   subq %rdx, %r10
   232: 0f8595ffffff             jne 0x83
   238: 415e                     popq %r14
   240: 415d                     popq %r13
   242: 415c                     popq %r12
   244: 415f                     popq %r15
   246: 5b                       popq %rbx
   247: c5f877                   vzeroupper 
   250: c3                       retq 
   251: 0f0b                     ud2 
   253: 0f0b                     ud2 
   255: cc                       int3 
