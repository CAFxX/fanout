; config_t3_2060 target=avx512-gfni twin=baseline topo=raid1()/enc op=encode U=8
; code_size=256 twins=[] uops_per_byte=0.033203125
     0: 53                       pushq %rbx
     1: 4c8bc7                   movq %rdi, %r8
     4: 4c8bce                   movq %rsi, %r9
     7: 488bca                   movq %rdx, %rcx
    10: 4d33db                   xorq %r11, %r11
    13: 488bc1                   movq %rcx, %rax
    16: 4881e800020000           subq $0x200, %rax
    23: 0f8297000000             jb 0xb4
    29: 0f1f00                   nopl (%rax)
    32: 4d8b10                   movq (%r8), %r10
    35: 62917d486f041a           vmovdqa32 (%r10, %r11), %zmm0
    42: 62917d486f4c1a01         vmovdqa32 0x40(%r10, %r11), %zmm1
    50: 62917d486f541a02         vmovdqa32 0x80(%r10, %r11), %zmm2
    58: 62917d486f5c1a03         vmovdqa32 0xc0(%r10, %r11), %zmm3
    66: 62917d486f641a04         vmovdqa32 0x100(%r10, %r11), %zmm4
    74: 62917d486f6c1a05         vmovdqa32 0x140(%r10, %r11), %zmm5
    82: 62917d486f741a06         vmovdqa32 0x180(%r10, %r11), %zmm6
    90: 62917d486f7c1a07         vmovdqa32 0x1c0(%r10, %r11), %zmm7
    98: 4d8b11                   movq (%r9), %r10
   101: 62917d487f041a           vmovdqa32 %zmm0, (%r10, %r11)
   108: 62917d487f4c1a01         vmovdqa32 %zmm1, 0x40(%r10, %r11)
   116: 62917d487f541a02         vmovdqa32 %zmm2, 0x80(%r10, %r11)
   124: 62917d487f5c1a03         vmovdqa32 %zmm3, 0xc0(%r10, %r11)
   132: 62917d487f641a04         vmovdqa32 %zmm4, 0x100(%r10, %r11)
   140: 62917d487f6c1a05         vmovdqa32 %zmm5, 0x140(%r10, %r11)
   148: 62917d487f741a06         vmovdqa32 %zmm6, 0x180(%r10, %r11)
   156: 62917d487f7c1a07         vmovdqa32 %zmm7, 0x1c0(%r10, %r11)
   164: 4981c300020000           addq $0x200, %r11
   171: 4c3bd8                   cmpq %rax, %r11
   174: 0f866cffffff             jbe 0x20
   180: 4c3bd9                   cmpq %rcx, %r11
   183: 0f8328000000             jae 0xe5
   189: 4153                     pushq %r11
   191: 4d8b11                   movq (%r9), %r10
   194: 4c8b1c24                 movq (%rsp), %r11
   198: 498b00                   movq (%r8), %rax
   201: 428a1418                 movb (%rax, %r11), %dl
   205: 8ada                     movb %dl, %bl
   207: 43881c1a                 movb %bl, (%r10, %r11)
   211: 33db                     xorl %ebx, %ebx
   213: 49ffc3                   incq %r11
   216: 4c3bd9                   cmpq %rcx, %r11
   219: 0f82e5ffffff             jb 0xc6
   225: 4883c408                 addq $8, %rsp
   229: 5b                       popq %rbx
   230: c5f877                   vzeroupper 
   233: c3                       retq 
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
