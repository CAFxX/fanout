; config_t1_729 target=sse twin=nt topo=rs(6,3)/rec{0} op=recover U=1
; code_size=256 twins=["nt"] uops_per_byte=0.5
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
   108: 66430fe7041a             movntdq %xmm0, (%r10, %r11)
   114: 4983c310                 addq $0x10, %r11
   118: 4c3bd8                   cmpq %rax, %r11
   121: 0f86c1ffffff             jbe 0x40
   127: 4c3bd9                   cmpq %rcx, %r11
   130: 0f835a000000             jae 0xe2
   136: 4153                     pushq %r11
   138: 4d8b11                   movq (%r9), %r10
   141: 4c8b1c24                 movq (%rsp), %r11
   145: 498b00                   movq (%r8), %rax
   148: 428a1418                 movb (%rax, %r11), %dl
   152: 8ada                     movb %dl, %bl
   154: 498b4008                 movq 8(%r8), %rax
   158: 428a1418                 movb (%rax, %r11), %dl
   162: 32da                     xorb %dl, %bl
   164: 498b4010                 movq 0x10(%r8), %rax
   168: 428a1418                 movb (%rax, %r11), %dl
   172: 32da                     xorb %dl, %bl
   174: 498b4018                 movq 0x18(%r8), %rax
   178: 428a1418                 movb (%rax, %r11), %dl
   182: 32da                     xorb %dl, %bl
   184: 498b4020                 movq 0x20(%r8), %rax
   188: 428a1418                 movb (%rax, %r11), %dl
   192: 32da                     xorb %dl, %bl
   194: 498b4028                 movq 0x28(%r8), %rax
   198: 428a1418                 movb (%rax, %r11), %dl
   202: 32da                     xorb %dl, %bl
   204: 43881c1a                 movb %bl, (%r10, %r11)
   208: 33db                     xorl %ebx, %ebx
   210: 49ffc3                   incq %r11
   213: 4c3bd9                   cmpq %rcx, %r11
   216: 0f82b3ffffff             jb 0x91
   222: 4883c408                 addq $8, %rsp
   226: 415e                     popq %r14
   228: 415d                     popq %r13
   230: 415c                     popq %r12
   232: 5b                       popq %rbx
   233: c3                       retq 
   234: 0f0b                     ud2 
   236: 0f0b                     ud2 
   238: 0f0b                     ud2 
   240: 0f0b                     ud2 
   242: 0f0b                     ud2 
   244: 0f0b                     ud2 
   246: 0f0b                     ud2 
   248: 0f0b                     ud2 
   250: 0f0b                     ud2 
   252: 0f0b                     ud2 
   254: 0f0b                     ud2 
