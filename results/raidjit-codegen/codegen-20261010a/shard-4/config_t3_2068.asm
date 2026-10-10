; config_t3_2068 target=avx2 twin=baseline topo=rs(3,1)/enc op=encode U=8
; code_size=448 twins=[] uops_per_byte=0.125
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
    33: 4881e800010000           subq $0x100, %rax
    40: 0f8231010000             jb 0x15f
    46: 666666666666660f1f840000000000 nopw (%rax, %rax)
    61: 0f1f00                   nopl (%rax)
    64: c4a17d6f041e             vmovdqa (%rsi, %r11), %ymm0
    70: c4a17d6f4c1e20           vmovdqa 0x20(%rsi, %r11), %ymm1
    77: c4a17d6f541e40           vmovdqa 0x40(%rsi, %r11), %ymm2
    84: c4a17d6f5c1e60           vmovdqa 0x60(%rsi, %r11), %ymm3
    91: c4a17d6fa41e80000000     vmovdqa 0x80(%rsi, %r11), %ymm4
   101: c4a17d6fac1ea0000000     vmovdqa 0xa0(%rsi, %r11), %ymm5
   111: c4a17d6fb41ec0000000     vmovdqa 0xc0(%rsi, %r11), %ymm6
   121: c4a17d6fbc1ee0000000     vmovdqa 0xe0(%rsi, %r11), %ymm7
   131: c4a17def041f             vpxor (%rdi, %r11), %ymm0, %ymm0
   137: c4a175ef4c1f20           vpxor 0x20(%rdi, %r11), %ymm1, %ymm1
   144: c4a16def541f40           vpxor 0x40(%rdi, %r11), %ymm2, %ymm2
   151: c4a165ef5c1f60           vpxor 0x60(%rdi, %r11), %ymm3, %ymm3
   158: c4a15defa41f80000000     vpxor 0x80(%rdi, %r11), %ymm4, %ymm4
   168: c4a155efac1fa0000000     vpxor 0xa0(%rdi, %r11), %ymm5, %ymm5
   178: c4a14defb41fc0000000     vpxor 0xc0(%rdi, %r11), %ymm6, %ymm6
   188: c4a145efbc1fe0000000     vpxor 0xe0(%rdi, %r11), %ymm7, %ymm7
   198: c4817def041c             vpxor (%r12, %r11), %ymm0, %ymm0
   204: c48175ef4c1c20           vpxor 0x20(%r12, %r11), %ymm1, %ymm1
   211: c4816def541c40           vpxor 0x40(%r12, %r11), %ymm2, %ymm2
   218: c48165ef5c1c60           vpxor 0x60(%r12, %r11), %ymm3, %ymm3
   225: c4815defa41c80000000     vpxor 0x80(%r12, %r11), %ymm4, %ymm4
   235: c48155efac1ca0000000     vpxor 0xa0(%r12, %r11), %ymm5, %ymm5
   245: c4814defb41cc0000000     vpxor 0xc0(%r12, %r11), %ymm6, %ymm6
   255: c48145efbc1ce0000000     vpxor 0xe0(%r12, %r11), %ymm7, %ymm7
   265: 4d8b11                   movq (%r9), %r10
   268: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   274: c4817d7f4c1a20           vmovdqa %ymm1, 0x20(%r10, %r11)
   281: c4817d7f541a40           vmovdqa %ymm2, 0x40(%r10, %r11)
   288: c4817d7f5c1a60           vmovdqa %ymm3, 0x60(%r10, %r11)
   295: c4817d7fa41a80000000     vmovdqa %ymm4, 0x80(%r10, %r11)
   305: c4817d7fac1aa0000000     vmovdqa %ymm5, 0xa0(%r10, %r11)
   315: c4817d7fb41ac0000000     vmovdqa %ymm6, 0xc0(%r10, %r11)
   325: c4817d7fbc1ae0000000     vmovdqa %ymm7, 0xe0(%r10, %r11)
   335: 4981c300010000           addq $0x100, %r11
   342: 4c3bd8                   cmpq %rax, %r11
   345: 0f86e1feffff             jbe 0x40
   351: 4c3bd9                   cmpq %rcx, %r11
   354: 0f833c000000             jae 0x1a4
   360: 4153                     pushq %r11
   362: 4d8b11                   movq (%r9), %r10
   365: 4c8b1c24                 movq (%rsp), %r11
   369: 498b00                   movq (%r8), %rax
   372: 428a1418                 movb (%rax, %r11), %dl
   376: 8ada                     movb %dl, %bl
   378: 498b4008                 movq 8(%r8), %rax
   382: 428a1418                 movb (%rax, %r11), %dl
   386: 32da                     xorb %dl, %bl
   388: 498b4010                 movq 0x10(%r8), %rax
   392: 428a1418                 movb (%rax, %r11), %dl
   396: 32da                     xorb %dl, %bl
   398: 43881c1a                 movb %bl, (%r10, %r11)
   402: 33db                     xorl %ebx, %ebx
   404: 49ffc3                   incq %r11
   407: 4c3bd9                   cmpq %rcx, %r11
   410: 0f82d1ffffff             jb 0x171
   416: 4883c408                 addq $8, %rsp
   420: 415e                     popq %r14
   422: 415d                     popq %r13
   424: 415c                     popq %r12
   426: 5b                       popq %rbx
   427: c5f877                   vzeroupper 
   430: c3                       retq 
   431: 0f0b                     ud2 
   433: 0f0b                     ud2 
   435: 0f0b                     ud2 
   437: 0f0b                     ud2 
   439: 0f0b                     ud2 
   441: 0f0b                     ud2 
   443: 0f0b                     ud2 
   445: 0f0b                     ud2 
   447: cc                       int3 
