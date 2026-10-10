; config_t1_157 target=avx512-gfni twin=nt topo=rs(6,3)/rec{0} op=recover U=1
; code_size=256 twins=["nt"] uops_per_byte=0.125
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
    41: 4883e840                 subq $0x40, %rax
    45: 0f8253000000             jb 0x86
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 62b17d486f041e           vmovdqa32 (%rsi, %r11), %zmm0
    71: 62b17d48ef041f           vpxord (%rdi, %r11), %zmm0, %zmm0
    78: 62917d48ef041c           vpxord (%r12, %r11), %zmm0, %zmm0
    85: 62917d48ef441d00         vpxord (%r13, %r11), %zmm0, %zmm0
    93: 62917d48ef041e           vpxord (%r14, %r11), %zmm0, %zmm0
   100: 4d8b5028                 movq 0x28(%r8), %r10
   104: 62917d48ef041a           vpxord (%r10, %r11), %zmm0, %zmm0
   111: 4d8b11                   movq (%r9), %r10
   114: 62917d48e7041a           vmovntdq %zmm0, (%r10, %r11)
   121: 4983c340                 addq $0x40, %r11
   125: 4c3bd8                   cmpq %rax, %r11
   128: 0f86baffffff             jbe 0x40
   134: 4c3bd9                   cmpq %rcx, %r11
   137: 0f835a000000             jae 0xe9
   143: 4153                     pushq %r11
   145: 4d8b11                   movq (%r9), %r10
   148: 4c8b1c24                 movq (%rsp), %r11
   152: 498b00                   movq (%r8), %rax
   155: 428a1418                 movb (%rax, %r11), %dl
   159: 8ada                     movb %dl, %bl
   161: 498b4008                 movq 8(%r8), %rax
   165: 428a1418                 movb (%rax, %r11), %dl
   169: 32da                     xorb %dl, %bl
   171: 498b4010                 movq 0x10(%r8), %rax
   175: 428a1418                 movb (%rax, %r11), %dl
   179: 32da                     xorb %dl, %bl
   181: 498b4018                 movq 0x18(%r8), %rax
   185: 428a1418                 movb (%rax, %r11), %dl
   189: 32da                     xorb %dl, %bl
   191: 498b4020                 movq 0x20(%r8), %rax
   195: 428a1418                 movb (%rax, %r11), %dl
   199: 32da                     xorb %dl, %bl
   201: 498b4028                 movq 0x28(%r8), %rax
   205: 428a1418                 movb (%rax, %r11), %dl
   209: 32da                     xorb %dl, %bl
   211: 43881c1a                 movb %bl, (%r10, %r11)
   215: 33db                     xorl %ebx, %ebx
   217: 49ffc3                   incq %r11
   220: 4c3bd9                   cmpq %rcx, %r11
   223: 0f82b3ffffff             jb 0x98
   229: 4883c408                 addq $8, %rsp
   233: 415e                     popq %r14
   235: 415d                     popq %r13
   237: 415c                     popq %r12
   239: 5b                       popq %rbx
   240: c5f877                   vzeroupper 
   243: c3                       retq 
   244: 0f0b                     ud2 
   246: 0f0b                     ud2 
   248: 0f0b                     ud2 
   250: 0f0b                     ud2 
   252: 0f0b                     ud2 
   254: 0f0b                     ud2 
