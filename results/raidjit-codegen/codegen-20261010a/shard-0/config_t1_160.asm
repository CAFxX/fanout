; config_t1_160 target=avx512-gfni twin=notail topo=rs(6,3)/rec{0} op=recover U=1
; code_size=192 twins=["notail"] uops_per_byte=0.125
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
   114: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   121: 4983c340                 addq $0x40, %r11
   125: 4c3bd8                   cmpq %rax, %r11
   128: 0f86baffffff             jbe 0x40
   134: 415e                     popq %r14
   136: 415d                     popq %r13
   138: 415c                     popq %r12
   140: 5b                       popq %rbx
   141: c5f877                   vzeroupper 
   144: c3                       retq 
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
