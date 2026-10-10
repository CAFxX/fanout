; config_t3_2059 target=avx512-gfni twin=baseline topo=raid1()/enc op=encode U=1
; code_size=128 twins=[] uops_per_byte=0.046875
     0: 53                       pushq %rbx
     1: 4c8bc7                   movq %rdi, %r8
     4: 4c8bce                   movq %rsi, %r9
     7: 488bca                   movq %rdx, %rcx
    10: 4d33db                   xorq %r11, %r11
    13: 488bc1                   movq %rcx, %rax
    16: 4883e840                 subq $0x40, %rax
    20: 0f8227000000             jb 0x41
    26: 660f1f440000             nopw (%rax, %rax)
    32: 4d8b10                   movq (%r8), %r10
    35: 62917d486f041a           vmovdqa32 (%r10, %r11), %zmm0
    42: 4d8b11                   movq (%r9), %r10
    45: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
    52: 4983c340                 addq $0x40, %r11
    56: 4c3bd8                   cmpq %rax, %r11
    59: 0f86dfffffff             jbe 0x20
    65: 4c3bd9                   cmpq %rcx, %r11
    68: 0f8328000000             jae 0x72
    74: 4153                     pushq %r11
    76: 4d8b11                   movq (%r9), %r10
    79: 4c8b1c24                 movq (%rsp), %r11
    83: 498b00                   movq (%r8), %rax
    86: 428a1418                 movb (%rax, %r11), %dl
    90: 8ada                     movb %dl, %bl
    92: 43881c1a                 movb %bl, (%r10, %r11)
    96: 33db                     xorl %ebx, %ebx
    98: 49ffc3                   incq %r11
   101: 4c3bd9                   cmpq %rcx, %r11
   104: 0f82e5ffffff             jb 0x53
   110: 4883c408                 addq $8, %rsp
   114: 5b                       popq %rbx
   115: c5f877                   vzeroupper 
   118: c3                       retq 
   119: 0f0b                     ud2 
   121: 0f0b                     ud2 
   123: 0f0b                     ud2 
   125: 0f0b                     ud2 
   127: cc                       int3 
