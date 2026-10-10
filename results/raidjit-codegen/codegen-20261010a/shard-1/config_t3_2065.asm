; config_t3_2065 target=avx2 twin=baseline topo=raid1()/enc op=encode U=1
; code_size=128 twins=[] uops_per_byte=0.09375
     0: 53                       pushq %rbx
     1: 4c8bc7                   movq %rdi, %r8
     4: 4c8bce                   movq %rsi, %r9
     7: 488bca                   movq %rdx, %rcx
    10: 4d33db                   xorq %r11, %r11
    13: 488bc1                   movq %rcx, %rax
    16: 4883e820                 subq $0x20, %rax
    20: 0f8225000000             jb 0x3f
    26: 660f1f440000             nopw (%rax, %rax)
    32: 4d8b10                   movq (%r8), %r10
    35: c4817d6f041a             vmovdqa (%r10, %r11), %ymm0
    41: 4d8b11                   movq (%r9), %r10
    44: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
    50: 4983c320                 addq $0x20, %r11
    54: 4c3bd8                   cmpq %rax, %r11
    57: 0f86e1ffffff             jbe 0x20
    63: 4c3bd9                   cmpq %rcx, %r11
    66: 0f8328000000             jae 0x70
    72: 4153                     pushq %r11
    74: 4d8b11                   movq (%r9), %r10
    77: 4c8b1c24                 movq (%rsp), %r11
    81: 498b00                   movq (%r8), %rax
    84: 428a1418                 movb (%rax, %r11), %dl
    88: 8ada                     movb %dl, %bl
    90: 43881c1a                 movb %bl, (%r10, %r11)
    94: 33db                     xorl %ebx, %ebx
    96: 49ffc3                   incq %r11
    99: 4c3bd9                   cmpq %rcx, %r11
   102: 0f82e5ffffff             jb 0x51
   108: 4883c408                 addq $8, %rsp
   112: 5b                       popq %rbx
   113: c5f877                   vzeroupper 
   116: c3                       retq 
   117: 0f0b                     ud2 
   119: 0f0b                     ud2 
   121: 0f0b                     ud2 
   123: 0f0b                     ud2 
   125: 0f0b                     ud2 
   127: cc                       int3 
