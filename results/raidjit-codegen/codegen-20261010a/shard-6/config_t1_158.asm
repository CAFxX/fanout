; config_t1_158 target=avx512-gfni twin=ternlog topo=rs(6,3)/rec{0} op=recover U=1
; code_size=256 twins=["ternlog"] uops_per_byte=0.140625
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
    45: 0f825c000000             jb 0x8f
    51: 66666666660f1f840000000000 nopw (%rax, %rax)
    64: 62f17d48efc0             vpxord %zmm0, %zmm0, %zmm0
    70: 62b17d486f0c1e           vmovdqa32 (%rsi, %r11), %zmm1
    77: 62b3754825041f96         vpternlogd $0x96, (%rdi, %r11), %zmm1, %zmm0
    85: 62917d486f0c1c           vmovdqa32 (%r12, %r11), %zmm1
    92: 6293754825441d0096       vpternlogd $0x96, (%r13, %r11), %zmm1, %zmm0
   101: 498b5028                 movq 0x28(%r8), %rdx
   105: 62917d486f0c1e           vmovdqa32 (%r14, %r11), %zmm1
   112: 62b3754825041a96         vpternlogd $0x96, (%rdx, %r11), %zmm1, %zmm0
   120: 4d8b11                   movq (%r9), %r10
   123: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   130: 4983c340                 addq $0x40, %r11
   134: 4c3bd8                   cmpq %rax, %r11
   137: 0f86b1ffffff             jbe 0x40
   143: 4c3bd9                   cmpq %rcx, %r11
   146: 0f835a000000             jae 0xf2
   152: 4153                     pushq %r11
   154: 4d8b11                   movq (%r9), %r10
   157: 4c8b1c24                 movq (%rsp), %r11
   161: 498b00                   movq (%r8), %rax
   164: 428a1418                 movb (%rax, %r11), %dl
   168: 8ada                     movb %dl, %bl
   170: 498b4008                 movq 8(%r8), %rax
   174: 428a1418                 movb (%rax, %r11), %dl
   178: 32da                     xorb %dl, %bl
   180: 498b4010                 movq 0x10(%r8), %rax
   184: 428a1418                 movb (%rax, %r11), %dl
   188: 32da                     xorb %dl, %bl
   190: 498b4018                 movq 0x18(%r8), %rax
   194: 428a1418                 movb (%rax, %r11), %dl
   198: 32da                     xorb %dl, %bl
   200: 498b4020                 movq 0x20(%r8), %rax
   204: 428a1418                 movb (%rax, %r11), %dl
   208: 32da                     xorb %dl, %bl
   210: 498b4028                 movq 0x28(%r8), %rax
   214: 428a1418                 movb (%rax, %r11), %dl
   218: 32da                     xorb %dl, %bl
   220: 43881c1a                 movb %bl, (%r10, %r11)
   224: 33db                     xorl %ebx, %ebx
   226: 49ffc3                   incq %r11
   229: 4c3bd9                   cmpq %rcx, %r11
   232: 0f82b3ffffff             jb 0xa1
   238: 4883c408                 addq $8, %rsp
   242: 415e                     popq %r14
   244: 415d                     popq %r13
   246: 415c                     popq %r12
   248: 5b                       popq %rbx
   249: c5f877                   vzeroupper 
   252: c3                       retq 
   253: 0f0b                     ud2 
   255: cc                       int3 
