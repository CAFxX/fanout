; config_t1_456 target=avx2 twin=nt topo=rs(6,3)/rec{0} op=recover U=8
; code_size=704 twins=["nt"] uops_per_byte=0.22265625
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
    41: 4881e800010000           subq $0x100, %rax
    48: 0f82f7010000             jb 0x22d
    54: 66660f1f840000000000     nopw (%rax, %rax)
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
   265: c4817def441d00           vpxor (%r13, %r11), %ymm0, %ymm0
   272: c48175ef4c1d20           vpxor 0x20(%r13, %r11), %ymm1, %ymm1
   279: c4816def541d40           vpxor 0x40(%r13, %r11), %ymm2, %ymm2
   286: c48165ef5c1d60           vpxor 0x60(%r13, %r11), %ymm3, %ymm3
   293: c4815defa41d80000000     vpxor 0x80(%r13, %r11), %ymm4, %ymm4
   303: c48155efac1da0000000     vpxor 0xa0(%r13, %r11), %ymm5, %ymm5
   313: c4814defb41dc0000000     vpxor 0xc0(%r13, %r11), %ymm6, %ymm6
   323: c48145efbc1de0000000     vpxor 0xe0(%r13, %r11), %ymm7, %ymm7
   333: c4817def041e             vpxor (%r14, %r11), %ymm0, %ymm0
   339: c48175ef4c1e20           vpxor 0x20(%r14, %r11), %ymm1, %ymm1
   346: c4816def541e40           vpxor 0x40(%r14, %r11), %ymm2, %ymm2
   353: c48165ef5c1e60           vpxor 0x60(%r14, %r11), %ymm3, %ymm3
   360: c4815defa41e80000000     vpxor 0x80(%r14, %r11), %ymm4, %ymm4
   370: c48155efac1ea0000000     vpxor 0xa0(%r14, %r11), %ymm5, %ymm5
   380: c4814defb41ec0000000     vpxor 0xc0(%r14, %r11), %ymm6, %ymm6
   390: c48145efbc1ee0000000     vpxor 0xe0(%r14, %r11), %ymm7, %ymm7
   400: 4d8b5028                 movq 0x28(%r8), %r10
   404: c4817def041a             vpxor (%r10, %r11), %ymm0, %ymm0
   410: c48175ef4c1a20           vpxor 0x20(%r10, %r11), %ymm1, %ymm1
   417: c4816def541a40           vpxor 0x40(%r10, %r11), %ymm2, %ymm2
   424: c48165ef5c1a60           vpxor 0x60(%r10, %r11), %ymm3, %ymm3
   431: c4815defa41a80000000     vpxor 0x80(%r10, %r11), %ymm4, %ymm4
   441: c48155efac1aa0000000     vpxor 0xa0(%r10, %r11), %ymm5, %ymm5
   451: c4814defb41ac0000000     vpxor 0xc0(%r10, %r11), %ymm6, %ymm6
   461: c48145efbc1ae0000000     vpxor 0xe0(%r10, %r11), %ymm7, %ymm7
   471: 4d8b11                   movq (%r9), %r10
   474: c4817de7041a             vmovntdq %ymm0, (%r10, %r11)
   480: c4817de74c1a20           vmovntdq %ymm1, 0x20(%r10, %r11)
   487: c4817de7541a40           vmovntdq %ymm2, 0x40(%r10, %r11)
   494: c4817de75c1a60           vmovntdq %ymm3, 0x60(%r10, %r11)
   501: c4817de7a41a80000000     vmovntdq %ymm4, 0x80(%r10, %r11)
   511: c4817de7ac1aa0000000     vmovntdq %ymm5, 0xa0(%r10, %r11)
   521: c4817de7b41ac0000000     vmovntdq %ymm6, 0xc0(%r10, %r11)
   531: c4817de7bc1ae0000000     vmovntdq %ymm7, 0xe0(%r10, %r11)
   541: 4981c300010000           addq $0x100, %r11
   548: 4c3bd8                   cmpq %rax, %r11
   551: 0f8613feffff             jbe 0x40
   557: 4c3bd9                   cmpq %rcx, %r11
   560: 0f835a000000             jae 0x290
   566: 4153                     pushq %r11
   568: 4d8b11                   movq (%r9), %r10
   571: 4c8b1c24                 movq (%rsp), %r11
   575: 498b00                   movq (%r8), %rax
   578: 428a1418                 movb (%rax, %r11), %dl
   582: 8ada                     movb %dl, %bl
   584: 498b4008                 movq 8(%r8), %rax
   588: 428a1418                 movb (%rax, %r11), %dl
   592: 32da                     xorb %dl, %bl
   594: 498b4010                 movq 0x10(%r8), %rax
   598: 428a1418                 movb (%rax, %r11), %dl
   602: 32da                     xorb %dl, %bl
   604: 498b4018                 movq 0x18(%r8), %rax
   608: 428a1418                 movb (%rax, %r11), %dl
   612: 32da                     xorb %dl, %bl
   614: 498b4020                 movq 0x20(%r8), %rax
   618: 428a1418                 movb (%rax, %r11), %dl
   622: 32da                     xorb %dl, %bl
   624: 498b4028                 movq 0x28(%r8), %rax
   628: 428a1418                 movb (%rax, %r11), %dl
   632: 32da                     xorb %dl, %bl
   634: 43881c1a                 movb %bl, (%r10, %r11)
   638: 33db                     xorl %ebx, %ebx
   640: 49ffc3                   incq %r11
   643: 4c3bd9                   cmpq %rcx, %r11
   646: 0f82b3ffffff             jb 0x23f
   652: 4883c408                 addq $8, %rsp
   656: 415e                     popq %r14
   658: 415d                     popq %r13
   660: 415c                     popq %r12
   662: 5b                       popq %rbx
   663: c5f877                   vzeroupper 
   666: c3                       retq 
   667: 0f0b                     ud2 
   669: 0f0b                     ud2 
   671: 0f0b                     ud2 
   673: 0f0b                     ud2 
   675: 0f0b                     ud2 
   677: 0f0b                     ud2 
   679: 0f0b                     ud2 
   681: 0f0b                     ud2 
   683: 0f0b                     ud2 
   685: 0f0b                     ud2 
   687: 0f0b                     ud2 
   689: 0f0b                     ud2 
   691: 0f0b                     ud2 
   693: 0f0b                     ud2 
   695: 0f0b                     ud2 
   697: 0f0b                     ud2 
   699: 0f0b                     ud2 
   701: 0f0b                     ud2 
   703: cc                       int3 
