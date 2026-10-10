; config_t3_2072 target=sse twin=baseline topo=raid1()/enc op=encode U=8
; code_size=256 twins=[] uops_per_byte=0.1328125
     0: 53                       pushq %rbx
     1: 4c8bc7                   movq %rdi, %r8
     4: 4c8bce                   movq %rsi, %r9
     7: 488bca                   movq %rdx, %rcx
    10: 4d33db                   xorq %r11, %r11
    13: 488bc1                   movq %rcx, %rax
    16: 4881e880000000           subq $0x80, %rax
    23: 0f8287000000             jb 0xa4
    29: 0f1f00                   nopl (%rax)
    32: 4d8b10                   movq (%r8), %r10
    35: 66430f6f041a             movdqa (%r10, %r11), %xmm0
    41: 66430f6f4c1a10           movdqa 0x10(%r10, %r11), %xmm1
    48: 66430f6f541a20           movdqa 0x20(%r10, %r11), %xmm2
    55: 66430f6f5c1a30           movdqa 0x30(%r10, %r11), %xmm3
    62: 66430f6f641a40           movdqa 0x40(%r10, %r11), %xmm4
    69: 66430f6f6c1a50           movdqa 0x50(%r10, %r11), %xmm5
    76: 66430f6f741a60           movdqa 0x60(%r10, %r11), %xmm6
    83: 66430f6f7c1a70           movdqa 0x70(%r10, %r11), %xmm7
    90: 4d8b11                   movq (%r9), %r10
    93: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
    99: 66430f7f4c1a10           movdqa %xmm1, 0x10(%r10, %r11)
   106: 66430f7f541a20           movdqa %xmm2, 0x20(%r10, %r11)
   113: 66430f7f5c1a30           movdqa %xmm3, 0x30(%r10, %r11)
   120: 66430f7f641a40           movdqa %xmm4, 0x40(%r10, %r11)
   127: 66430f7f6c1a50           movdqa %xmm5, 0x50(%r10, %r11)
   134: 66430f7f741a60           movdqa %xmm6, 0x60(%r10, %r11)
   141: 66430f7f7c1a70           movdqa %xmm7, 0x70(%r10, %r11)
   148: 4981c380000000           addq $0x80, %r11
   155: 4c3bd8                   cmpq %rax, %r11
   158: 0f867cffffff             jbe 0x20
   164: 4c3bd9                   cmpq %rcx, %r11
   167: 0f8328000000             jae 0xd5
   173: 4153                     pushq %r11
   175: 4d8b11                   movq (%r9), %r10
   178: 4c8b1c24                 movq (%rsp), %r11
   182: 498b00                   movq (%r8), %rax
   185: 428a1418                 movb (%rax, %r11), %dl
   189: 8ada                     movb %dl, %bl
   191: 43881c1a                 movb %bl, (%r10, %r11)
   195: 33db                     xorl %ebx, %ebx
   197: 49ffc3                   incq %r11
   200: 4c3bd9                   cmpq %rcx, %r11
   203: 0f82e5ffffff             jb 0xb6
   209: 4883c408                 addq $8, %rsp
   213: 5b                       popq %rbx
   214: c3                       retq 
   215: 0f0b                     ud2 
   217: 0f0b                     ud2 
   219: 0f0b                     ud2 
   221: 0f0b                     ud2 
   223: 0f0b                     ud2 
   225: 0f0b                     ud2 
   227: 0f0b                     ud2 
   229: 0f0b                     ud2 
   231: 0f0b                     ud2 
   233: 0f0b                     ud2 
   235: 0f0b                     ud2 
   237: 0f0b                     ud2 
   239: 0f0b                     ud2 
   241: 0f0b                     ud2 
   243: 0f0b                     ud2 
   245: 0f0b                     ud2 
   247: 0f0b                     ud2 
   249: 0f0b                     ud2 
   251: 0f0b                     ud2 
   253: 0f0b                     ud2 
   255: cc                       int3 
