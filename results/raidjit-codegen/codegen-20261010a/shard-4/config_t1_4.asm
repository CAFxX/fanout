; config_t1_4 target=avx512-gfni twin=notail topo=raid5(4)/enc op=encode U=1
; code_size=192 twins=["notail"] uops_per_byte=0.078125
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
    96: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   103: 4983c340                 addq $0x40, %r11
   107: 4c3bd8                   cmpq %rax, %r11
   110: 0f86ccffffff             jbe 0x40
   116: 415e                     popq %r14
   118: 415d                     popq %r13
   120: 415c                     popq %r12
   122: 5b                       popq %rbx
   123: c5f877                   vzeroupper 
   126: c3                       retq 
   127: 0f0b                     ud2 
   129: 0f0b                     ud2 
   131: 0f0b                     ud2 
   133: 0f0b                     ud2 
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
