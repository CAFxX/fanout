; config_t1_290 target=avx2 twin=notail topo=raid5(4)/enc op=encode U=1
; code_size=128 twins=["notail"] uops_per_byte=0.15625
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
    92: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
    98: 4983c320                 addq $0x20, %r11
   102: 4c3bd8                   cmpq %rax, %r11
   105: 0f86d1ffffff             jbe 0x40
   111: 415e                     popq %r14
   113: 415d                     popq %r13
   115: 415c                     popq %r12
   117: 5b                       popq %rbx
   118: c5f877                   vzeroupper 
   121: c3                       retq 
   122: 0f0b                     ud2 
   124: 0f0b                     ud2 
   126: 0f0b                     ud2 
