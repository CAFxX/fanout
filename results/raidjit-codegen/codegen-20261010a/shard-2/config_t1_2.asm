; config_t1_2 target=avx512-gfni twin=ternlog topo=raid5(4)/enc op=encode U=1
; code_size=256 twins=["ternlog"] uops_per_byte=0.09375
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
    41: 0f824d000000             jb 0x7c
    47: 666666666666660f1f840000000000 nopw (%rax, %rax)
    62: 6690                     nop 
    64: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
    70: 62b17d486f0c1e           vmovdqa32 (%rsi, %r11), %zmm1
    77: 62b3754825041f96         vpternlogd $0x96, (%rdi, %r11), %zmm1, %zmm0
    85: 62917d486f0c1c           vmovdqa32 (%r12, %r11), %zmm1
    92: 6293754825441d0096       vpternlogd $0x96, (%r13, %r11), %zmm1, %zmm0
   101: 4d8b11                   movq (%r9), %r10
   104: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   111: 4983c340                 addq $0x40, %r11
   115: 4c3bd8                   cmpq %rax, %r11
   118: 0f86c4ffffff             jbe 0x40
   124: 4c3bd9                   cmpq %rcx, %r11
   127: 0f8346000000             jae 0xcb
   133: 4153                     pushq %r11
   135: 4d8b11                   movq (%r9), %r10
   138: 4c8b1c24                 movq (%rsp), %r11
   142: 498b00                   movq (%r8), %rax
   145: 428a1418                 movb (%rax, %r11), %dl
   149: 8ada                     movb %dl, %bl
   151: 498b4008                 movq 8(%r8), %rax
   155: 428a1418                 movb (%rax, %r11), %dl
   159: 32da                     xorb %dl, %bl
   161: 498b4010                 movq 0x10(%r8), %rax
   165: 428a1418                 movb (%rax, %r11), %dl
   169: 32da                     xorb %dl, %bl
   171: 498b4018                 movq 0x18(%r8), %rax
   175: 428a1418                 movb (%rax, %r11), %dl
   179: 32da                     xorb %dl, %bl
   181: 43881c1a                 movb %bl, (%r10, %r11)
   185: 33db                     xorl %ebx, %ebx
   187: 49ffc3                   incq %r11
   190: 4c3bd9                   cmpq %rcx, %r11
   193: 0f82c7ffffff             jb 0x8e
   199: 4883c408                 addq $8, %rsp
   203: 415e                     popq %r14
   205: 415d                     popq %r13
   207: 415c                     popq %r12
   209: 5b                       popq %rbx
   210: c5f877                   vzeroupper 
   213: c3                       retq 
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
