; config_t3_2066 target=avx2 twin=baseline topo=raid1()/enc op=encode U=8
; code_size=256 twins=[] uops_per_byte=0.06640625
     0: 53                       pushq %rbx
     1: 4c8bc7                   movq %rdi, %r8
     4: 4c8bce                   movq %rsi, %r9
     7: 488bca                   movq %rdx, %rcx
    10: 4d33db                   xorq %r11, %r11
    13: 488bc1                   movq %rcx, %rax
    16: 4881e800010000           subq $0x100, %rax
    23: 0f829f000000             jb 0xbc
    29: 0f1f00                   nopl (%rax)
    32: 4d8b10                   movq (%r8), %r10
    35: c4817d6f041a             vmovdqa (%r10, %r11), %ymm0
    41: c4817d6f4c1a20           vmovdqa 0x20(%r10, %r11), %ymm1
    48: c4817d6f541a40           vmovdqa 0x40(%r10, %r11), %ymm2
    55: c4817d6f5c1a60           vmovdqa 0x60(%r10, %r11), %ymm3
    62: c4817d6fa41a80000000     vmovdqa 0x80(%r10, %r11), %ymm4
    72: c4817d6fac1aa0000000     vmovdqa 0xa0(%r10, %r11), %ymm5
    82: c4817d6fb41ac0000000     vmovdqa 0xc0(%r10, %r11), %ymm6
    92: c4817d6fbc1ae0000000     vmovdqa 0xe0(%r10, %r11), %ymm7
   102: 4d8b11                   movq (%r9), %r10
   105: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   111: c4817d7f4c1a20           vmovdqa %ymm1, 0x20(%r10, %r11)
   118: c4817d7f541a40           vmovdqa %ymm2, 0x40(%r10, %r11)
   125: c4817d7f5c1a60           vmovdqa %ymm3, 0x60(%r10, %r11)
   132: c4817d7fa41a80000000     vmovdqa %ymm4, 0x80(%r10, %r11)
   142: c4817d7fac1aa0000000     vmovdqa %ymm5, 0xa0(%r10, %r11)
   152: c4817d7fb41ac0000000     vmovdqa %ymm6, 0xc0(%r10, %r11)
   162: c4817d7fbc1ae0000000     vmovdqa %ymm7, 0xe0(%r10, %r11)
   172: 4981c300010000           addq $0x100, %r11
   179: 4c3bd8                   cmpq %rax, %r11
   182: 0f8664ffffff             jbe 0x20
   188: 4c3bd9                   cmpq %rcx, %r11
   191: 0f8328000000             jae 0xed
   197: 4153                     pushq %r11
   199: 4d8b11                   movq (%r9), %r10
   202: 4c8b1c24                 movq (%rsp), %r11
   206: 498b00                   movq (%r8), %rax
   209: 428a1418                 movb (%rax, %r11), %dl
   213: 8ada                     movb %dl, %bl
   215: 43881c1a                 movb %bl, (%r10, %r11)
   219: 33db                     xorl %ebx, %ebx
   221: 49ffc3                   incq %r11
   224: 4c3bd9                   cmpq %rcx, %r11
   227: 0f82e5ffffff             jb 0xce
   233: 4883c408                 addq $8, %rsp
   237: 5b                       popq %rbx
   238: c5f877                   vzeroupper 
   241: c3                       retq 
   242: 0f0b                     ud2 
   244: 0f0b                     ud2 
   246: 0f0b                     ud2 
   248: 0f0b                     ud2 
   250: 0f0b                     ud2 
   252: 0f0b                     ud2 
   254: 0f0b                     ud2 
