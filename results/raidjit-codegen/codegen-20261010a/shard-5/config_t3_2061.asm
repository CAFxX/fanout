; config_t3_2061 target=avx512-gfni twin=baseline topo=rs(3,1)/enc op=encode U=1
; code_size=192 twins=[] uops_per_byte=0.0625
     0: 53                       pushq %rbx
     1: 4154                     pushq %r12
     3: 4155                     pushq %r13
     5: 4156                     pushq %r14
     7: 4c8bc7                   movq %rdi, %r8
    10: 4c8bce                   movq %rsi, %r9
    13: 488bca                   movq %rdx, %rcx
    16: 498b30                   movq (%r8), %rsi
    19: 498b7808                 movq 8(%r8), %rdi
    23: 4d8b6010                 movq 0x10(%r8), %r12
    27: 4d33db                   xorq %r11, %r11
    30: 488bc1                   movq %rcx, %rax
    33: 4883e840                 subq $0x40, %rax
    37: 0f8241000000             jb 0x6c
    43: 666666666666660f1f840000000000 nopw (%rax, %rax)
    58: 660f1f440000             nopw (%rax, %rax)
    64: 62b17d486f041e           vmovdqa32 (%rsi, %r11), %zmm0
    71: 62b17d48ef041f           vpxord (%rdi, %r11), %zmm0, %zmm0
    78: 62917d48ef041c           vpxord (%r12, %r11), %zmm0, %zmm0
    85: 4d8b11                   movq (%r9), %r10
    88: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
    95: 4983c340                 addq $0x40, %r11
    99: 4c3bd8                   cmpq %rax, %r11
   102: 0f86d4ffffff             jbe 0x40
   108: 4c3bd9                   cmpq %rcx, %r11
   111: 0f833c000000             jae 0xb1
   117: 4153                     pushq %r11
   119: 4d8b11                   movq (%r9), %r10
   122: 4c8b1c24                 movq (%rsp), %r11
   126: 498b00                   movq (%r8), %rax
   129: 428a1418                 movb (%rax, %r11), %dl
   133: 8ada                     movb %dl, %bl
   135: 498b4008                 movq 8(%r8), %rax
   139: 428a1418                 movb (%rax, %r11), %dl
   143: 32da                     xorb %dl, %bl
   145: 498b4010                 movq 0x10(%r8), %rax
   149: 428a1418                 movb (%rax, %r11), %dl
   153: 32da                     xorb %dl, %bl
   155: 43881c1a                 movb %bl, (%r10, %r11)
   159: 33db                     xorl %ebx, %ebx
   161: 49ffc3                   incq %r11
   164: 4c3bd9                   cmpq %rcx, %r11
   167: 0f82d1ffffff             jb 0x7e
   173: 4883c408                 addq $8, %rsp
   177: 415e                     popq %r14
   179: 415d                     popq %r13
   181: 415c                     popq %r12
   183: 5b                       popq %rbx
   184: c5f877                   vzeroupper 
   187: c3                       retq 
   188: 0f0b                     ud2 
   190: 0f0b                     ud2 
