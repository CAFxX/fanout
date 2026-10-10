; config_t3_2074 target=sse twin=baseline topo=rs(3,1)/enc op=encode U=8
; code_size=384 twins=[] uops_per_byte=0.25
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
    33: 4881e880000000           subq $0x80, %rax
    40: 0f8201010000             jb 0x12f
    46: 666666666666660f1f840000000000 nopw (%rax, %rax)
    61: 0f1f00                   nopl (%rax)
    64: 66420f6f041e             movdqa (%rsi, %r11), %xmm0
    70: 66420f6f4c1e10           movdqa 0x10(%rsi, %r11), %xmm1
    77: 66420f6f541e20           movdqa 0x20(%rsi, %r11), %xmm2
    84: 66420f6f5c1e30           movdqa 0x30(%rsi, %r11), %xmm3
    91: 66420f6f641e40           movdqa 0x40(%rsi, %r11), %xmm4
    98: 66420f6f6c1e50           movdqa 0x50(%rsi, %r11), %xmm5
   105: 66420f6f741e60           movdqa 0x60(%rsi, %r11), %xmm6
   112: 66420f6f7c1e70           movdqa 0x70(%rsi, %r11), %xmm7
   119: 66420fef041f             pxor (%rdi, %r11), %xmm0
   125: 66420fef4c1f10           pxor 0x10(%rdi, %r11), %xmm1
   132: 66420fef541f20           pxor 0x20(%rdi, %r11), %xmm2
   139: 66420fef5c1f30           pxor 0x30(%rdi, %r11), %xmm3
   146: 66420fef641f40           pxor 0x40(%rdi, %r11), %xmm4
   153: 66420fef6c1f50           pxor 0x50(%rdi, %r11), %xmm5
   160: 66420fef741f60           pxor 0x60(%rdi, %r11), %xmm6
   167: 66420fef7c1f70           pxor 0x70(%rdi, %r11), %xmm7
   174: 66430fef041c             pxor (%r12, %r11), %xmm0
   180: 66430fef4c1c10           pxor 0x10(%r12, %r11), %xmm1
   187: 66430fef541c20           pxor 0x20(%r12, %r11), %xmm2
   194: 66430fef5c1c30           pxor 0x30(%r12, %r11), %xmm3
   201: 66430fef641c40           pxor 0x40(%r12, %r11), %xmm4
   208: 66430fef6c1c50           pxor 0x50(%r12, %r11), %xmm5
   215: 66430fef741c60           pxor 0x60(%r12, %r11), %xmm6
   222: 66430fef7c1c70           pxor 0x70(%r12, %r11), %xmm7
   229: 4d8b11                   movq (%r9), %r10
   232: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   238: 66430f7f4c1a10           movdqa %xmm1, 0x10(%r10, %r11)
   245: 66430f7f541a20           movdqa %xmm2, 0x20(%r10, %r11)
   252: 66430f7f5c1a30           movdqa %xmm3, 0x30(%r10, %r11)
   259: 66430f7f641a40           movdqa %xmm4, 0x40(%r10, %r11)
   266: 66430f7f6c1a50           movdqa %xmm5, 0x50(%r10, %r11)
   273: 66430f7f741a60           movdqa %xmm6, 0x60(%r10, %r11)
   280: 66430f7f7c1a70           movdqa %xmm7, 0x70(%r10, %r11)
   287: 4981c380000000           addq $0x80, %r11
   294: 4c3bd8                   cmpq %rax, %r11
   297: 0f8611ffffff             jbe 0x40
   303: 4c3bd9                   cmpq %rcx, %r11
   306: 0f833c000000             jae 0x174
   312: 4153                     pushq %r11
   314: 4d8b11                   movq (%r9), %r10
   317: 4c8b1c24                 movq (%rsp), %r11
   321: 498b00                   movq (%r8), %rax
   324: 428a1418                 movb (%rax, %r11), %dl
   328: 8ada                     movb %dl, %bl
   330: 498b4008                 movq 8(%r8), %rax
   334: 428a1418                 movb (%rax, %r11), %dl
   338: 32da                     xorb %dl, %bl
   340: 498b4010                 movq 0x10(%r8), %rax
   344: 428a1418                 movb (%rax, %r11), %dl
   348: 32da                     xorb %dl, %bl
   350: 43881c1a                 movb %bl, (%r10, %r11)
   354: 33db                     xorl %ebx, %ebx
   356: 49ffc3                   incq %r11
   359: 4c3bd9                   cmpq %rcx, %r11
   362: 0f82d1ffffff             jb 0x141
   368: 4883c408                 addq $8, %rsp
   372: 415e                     popq %r14
   374: 415d                     popq %r13
   376: 415c                     popq %r12
   378: 5b                       popq %rbx
   379: c3                       retq 
   380: 0f0b                     ud2 
   382: 0f0b                     ud2 
