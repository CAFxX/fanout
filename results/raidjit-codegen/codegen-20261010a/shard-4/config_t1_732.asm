; config_t1_732 target=sse twin=notail topo=rs(6,3)/rec{0} op=recover U=1
; code_size=192 twins=["notail"] uops_per_byte=0.5
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
    27: 4d8b6818                 movq 0x18(%r8), %r13
    31: 4d8b7020                 movq 0x20(%r8), %r14
    35: 4d33db                   xorq %r11, %r11
    38: 488bc1                   movq %rcx, %rax
    41: 4883e810                 subq $0x10, %rax
    45: 0f824c000000             jb 0x7f
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 66420f6f041e             movdqa (%rsi, %r11), %xmm0
    70: 66420fef041f             pxor (%rdi, %r11), %xmm0
    76: 66430fef041c             pxor (%r12, %r11), %xmm0
    82: 66430fef441d00           pxor (%r13, %r11), %xmm0
    89: 66430fef041e             pxor (%r14, %r11), %xmm0
    95: 4d8b5028                 movq 0x28(%r8), %r10
    99: 66430fef041a             pxor (%r10, %r11), %xmm0
   105: 4d8b11                   movq (%r9), %r10
   108: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   114: 4983c310                 addq $0x10, %r11
   118: 4c3bd8                   cmpq %rax, %r11
   121: 0f86c1ffffff             jbe 0x40
   127: 415e                     popq %r14
   129: 415d                     popq %r13
   131: 415c                     popq %r12
   133: 5b                       popq %rbx
   134: c3                       retq 
   135: 0f0b                     ud2 
   137: 0f0b                     ud2 
   139: 0f0b                     ud2 
   141: 0f0b                     ud2 
   143: 0f0b                     ud2 
   145: 0f0b                     ud2 
   147: 0f0b                     ud2 
   149: 0f0b                     ud2 
   151: 0f0b                     ud2 
   153: 0f0b                     ud2 
   155: 0f0b                     ud2 
   157: 0f0b                     ud2 
   159: 0f0b                     ud2 
   161: 0f0b                     ud2 
   163: 0f0b                     ud2 
   165: 0f0b                     ud2 
   167: 0f0b                     ud2 
   169: 0f0b                     ud2 
   171: 0f0b                     ud2 
   173: 0f0b                     ud2 
   175: 0f0b                     ud2 
   177: 0f0b                     ud2 
   179: 0f0b                     ud2 
   181: 0f0b                     ud2 
   183: 0f0b                     ud2 
   185: 0f0b                     ud2 
   187: 0f0b                     ud2 
   189: 0f0b                     ud2 
   191: cc                       int3 
