; config_t1_446 target=avx2 twin=notail topo=rs(6,3)/rec{0} op=recover U=1
; code_size=192 twins=["notail"] uops_per_byte=0.25
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
    41: 4883e820                 subq $0x20, %rax
    45: 0f824c000000             jb 0x7f
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: c4a17d6f041e             vmovdqa (%rsi, %r11), %ymm0
    70: c4a17def041f             vpxor (%rdi, %r11), %ymm0, %ymm0
    76: c4817def041c             vpxor (%r12, %r11), %ymm0, %ymm0
    82: c4817def441d00           vpxor (%r13, %r11), %ymm0, %ymm0
    89: c4817def041e             vpxor (%r14, %r11), %ymm0, %ymm0
    95: 4d8b5028                 movq 0x28(%r8), %r10
    99: c4817def041a             vpxor (%r10, %r11), %ymm0, %ymm0
   105: 4d8b11                   movq (%r9), %r10
   108: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   114: 4983c320                 addq $0x20, %r11
   118: 4c3bd8                   cmpq %rax, %r11
   121: 0f86c1ffffff             jbe 0x40
   127: 415e                     popq %r14
   129: 415d                     popq %r13
   131: 415c                     popq %r12
   133: 5b                       popq %rbx
   134: c5f877                   vzeroupper 
   137: c3                       retq 
   138: 0f0b                     ud2 
   140: 0f0b                     ud2 
   142: 0f0b                     ud2 
   144: 0f0b                     ud2 
   146: 0f0b                     ud2 
   148: 0f0b                     ud2 
   150: 0f0b                     ud2 
   152: 0f0b                     ud2 
   154: 0f0b                     ud2 
   156: 0f0b                     ud2 
   158: 0f0b                     ud2 
   160: 0f0b                     ud2 
   162: 0f0b                     ud2 
   164: 0f0b                     ud2 
   166: 0f0b                     ud2 
   168: 0f0b                     ud2 
   170: 0f0b                     ud2 
   172: 0f0b                     ud2 
   174: 0f0b                     ud2 
   176: 0f0b                     ud2 
   178: 0f0b                     ud2 
   180: 0f0b                     ud2 
   182: 0f0b                     ud2 
   184: 0f0b                     ud2 
   186: 0f0b                     ud2 
   188: 0f0b                     ud2 
   190: 0f0b                     ud2 
