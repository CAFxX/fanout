; config_t1_1 target=avx512-gfni twin=nt topo=raid5(4)/enc op=encode U=1
; code_size=256 twins=["nt"] uops_per_byte=0.078125
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
    31: 4d33db                   xorq %r11, %r11
    34: 488bc1                   movq %rcx, %rax
    37: 4883e840                 subq $0x40, %rax
    41: 0f8245000000             jb 0x74
    47: 666666666666660f1f840000000000 nopw (%rax, %rax)
    62: 6690                     nop 
    64: 62b17d486f041e           vmovdqa32 (%rsi, %r11), %zmm0
    71: 62b17d48ef041f           vpxord (%rdi, %r11), %zmm0, %zmm0
    78: 62917d48ef041c           vpxord (%r12, %r11), %zmm0, %zmm0
    85: 62917d48ef441d00         vpxord (%r13, %r11), %zmm0, %zmm0
    93: 4d8b11                   movq (%r9), %r10
    96: 62917d48e7041a           vmovntdq %zmm0, (%r10, %r11)
   103: 4983c340                 addq $0x40, %r11
   107: 4c3bd8                   cmpq %rax, %r11
   110: 0f86ccffffff             jbe 0x40
   116: 4c3bd9                   cmpq %rcx, %r11
   119: 0f8346000000             jae 0xc3
   125: 4153                     pushq %r11
   127: 4d8b11                   movq (%r9), %r10
   130: 4c8b1c24                 movq (%rsp), %r11
   134: 498b00                   movq (%r8), %rax
   137: 428a1418                 movb (%rax, %r11), %dl
   141: 8ada                     movb %dl, %bl
   143: 498b4008                 movq 8(%r8), %rax
   147: 428a1418                 movb (%rax, %r11), %dl
   151: 32da                     xorb %dl, %bl
   153: 498b4010                 movq 0x10(%r8), %rax
   157: 428a1418                 movb (%rax, %r11), %dl
   161: 32da                     xorb %dl, %bl
   163: 498b4018                 movq 0x18(%r8), %rax
   167: 428a1418                 movb (%rax, %r11), %dl
   171: 32da                     xorb %dl, %bl
   173: 43881c1a                 movb %bl, (%r10, %r11)
   177: 33db                     xorl %ebx, %ebx
   179: 49ffc3                   incq %r11
   182: 4c3bd9                   cmpq %rcx, %r11
   185: 0f82c7ffffff             jb 0x86
   191: 4883c408                 addq $8, %rsp
   195: 415e                     popq %r14
   197: 415d                     popq %r13
   199: 415c                     popq %r12
   201: 5b                       popq %rbx
   202: c5f877                   vzeroupper 
   205: c3                       retq 
   206: 0f0b                     ud2 
   208: 0f0b                     ud2 
   210: 0f0b                     ud2 
   212: 0f0b                     ud2 
   214: 0f0b                     ud2 
   216: 0f0b                     ud2 
   218: 0f0b                     ud2 
   220: 0f0b                     ud2 
   222: 0f0b                     ud2 
   224: 0f0b                     ud2 
   226: 0f0b                     ud2 
   228: 0f0b                     ud2 
   230: 0f0b                     ud2 
   232: 0f0b                     ud2 
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
