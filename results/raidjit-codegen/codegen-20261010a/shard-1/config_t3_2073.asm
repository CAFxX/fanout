; config_t3_2073 target=sse twin=baseline topo=rs(3,1)/enc op=encode U=1
; code_size=192 twins=[] uops_per_byte=0.25
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
    33: 4883e810                 subq $0x10, %rax
    37: 0f823d000000             jb 0x68
    43: 666666666666660f1f840000000000 nopw (%rax, %rax)
    58: 660f1f440000             nopw (%rax, %rax)
    64: 66420f6f041e             movdqa (%rsi, %r11), %xmm0
    70: 66420fef041f             pxor (%rdi, %r11), %xmm0
    76: 66430fef041c             pxor (%r12, %r11), %xmm0
    82: 4d8b11                   movq (%r9), %r10
    85: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
    91: 4983c310                 addq $0x10, %r11
    95: 4c3bd8                   cmpq %rax, %r11
    98: 0f86d8ffffff             jbe 0x40
   104: 4c3bd9                   cmpq %rcx, %r11
   107: 0f833c000000             jae 0xad
   113: 4153                     pushq %r11
   115: 4d8b11                   movq (%r9), %r10
   118: 4c8b1c24                 movq (%rsp), %r11
   122: 498b00                   movq (%r8), %rax
   125: 428a1418                 movb (%rax, %r11), %dl
   129: 8ada                     movb %dl, %bl
   131: 498b4008                 movq 8(%r8), %rax
   135: 428a1418                 movb (%rax, %r11), %dl
   139: 32da                     xorb %dl, %bl
   141: 498b4010                 movq 0x10(%r8), %rax
   145: 428a1418                 movb (%rax, %r11), %dl
   149: 32da                     xorb %dl, %bl
   151: 43881c1a                 movb %bl, (%r10, %r11)
   155: 33db                     xorl %ebx, %ebx
   157: 49ffc3                   incq %r11
   160: 4c3bd9                   cmpq %rcx, %r11
   163: 0f82d1ffffff             jb 0x7a
   169: 4883c408                 addq $8, %rsp
   173: 415e                     popq %r14
   175: 415d                     popq %r13
   177: 415c                     popq %r12
   179: 5b                       popq %rbx
   180: c3                       retq 
   181: 0f0b                     ud2 
   183: 0f0b                     ud2 
   185: 0f0b                     ud2 
   187: 0f0b                     ud2 
   189: 0f0b                     ud2 
   191: cc                       int3 
