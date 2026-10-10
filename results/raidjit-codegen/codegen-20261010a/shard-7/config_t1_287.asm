; config_t1_287 target=avx2 twin=nt topo=raid5(4)/enc op=encode U=1
; code_size=256 twins=["nt"] uops_per_byte=0.15625
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
    37: 4883e820                 subq $0x20, %rax
    41: 0f8240000000             jb 0x6f
    47: 666666666666660f1f840000000000 nopw (%rax, %rax)
    62: 6690                     nop 
    64: c4a17d6f041e             vmovdqa (%rsi, %r11), %ymm0
    70: c4a17def041f             vpxor (%rdi, %r11), %ymm0, %ymm0
    76: c4817def041c             vpxor (%r12, %r11), %ymm0, %ymm0
    82: c4817def441d00           vpxor (%r13, %r11), %ymm0, %ymm0
    89: 4d8b11                   movq (%r9), %r10
    92: c4817de7041a             vmovntdq %ymm0, (%r10, %r11)
    98: 4983c320                 addq $0x20, %r11
   102: 4c3bd8                   cmpq %rax, %r11
   105: 0f86d1ffffff             jbe 0x40
   111: 4c3bd9                   cmpq %rcx, %r11
   114: 0f8346000000             jae 0xbe
   120: 4153                     pushq %r11
   122: 4d8b11                   movq (%r9), %r10
   125: 4c8b1c24                 movq (%rsp), %r11
   129: 498b00                   movq (%r8), %rax
   132: 428a1418                 movb (%rax, %r11), %dl
   136: 8ada                     movb %dl, %bl
   138: 498b4008                 movq 8(%r8), %rax
   142: 428a1418                 movb (%rax, %r11), %dl
   146: 32da                     xorb %dl, %bl
   148: 498b4010                 movq 0x10(%r8), %rax
   152: 428a1418                 movb (%rax, %r11), %dl
   156: 32da                     xorb %dl, %bl
   158: 498b4018                 movq 0x18(%r8), %rax
   162: 428a1418                 movb (%rax, %r11), %dl
   166: 32da                     xorb %dl, %bl
   168: 43881c1a                 movb %bl, (%r10, %r11)
   172: 33db                     xorl %ebx, %ebx
   174: 49ffc3                   incq %r11
   177: 4c3bd9                   cmpq %rcx, %r11
   180: 0f82c7ffffff             jb 0x81
   186: 4883c408                 addq $8, %rsp
   190: 415e                     popq %r14
   192: 415d                     popq %r13
   194: 415c                     popq %r12
   196: 5b                       popq %rbx
   197: c5f877                   vzeroupper 
   200: c3                       retq 
   201: 0f0b                     ud2 
   203: 0f0b                     ud2 
   205: 0f0b                     ud2 
   207: 0f0b                     ud2 
   209: 0f0b                     ud2 
   211: 0f0b                     ud2 
   213: 0f0b                     ud2 
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
