; config_t1_529 target=avx2 twin=xorbit topo=rs(6,3)/upd{0} op=parity_update U=1
; code_size=1024 twins=["xorbit"] uops_per_byte=3.8125
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3ded030000           leaq 0x3ed(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: c4c17d6f37               vmovdqa (%r15), %ymm6
    27: 488bc1                   movq %rcx, %rax
    30: 4883e820                 subq $0x20, %rax
    34: 0f82b2020000             jb 0x2da
    40: 666666666666660f1f840000000000 nopw (%rax, %rax)
    55: 660f1f840000000000       nopw (%rax, %rax)
    64: 4d8b5010                 movq 0x10(%r8), %r10
    68: c4817d6f041a             vmovdqa (%r10, %r11), %ymm0
    74: 4d8b5018                 movq 0x18(%r8), %r10
    78: c4817d6f0c1a             vmovdqa (%r10, %r11), %ymm1
    84: 4d8b5020                 movq 0x20(%r8), %r10
    88: c4817d6f141a             vmovdqa (%r10, %r11), %ymm2
    94: 4d8b10                   movq (%r8), %r10
    97: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   103: c5dd71d300               vpsrlw $0, %ymm3, %ymm4
   108: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   112: 4d8b5008                 movq 8(%r8), %r10
   116: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   122: c5d571d300               vpsrlw $0, %ymm3, %ymm5
   127: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   131: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   135: c5fdefc4                 vpxor %ymm4, %ymm0, %ymm0
   139: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   143: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
   147: c5d571f401               vpsllw $1, %ymm4, %ymm5
   152: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   156: c5d571f402               vpsllw $2, %ymm4, %ymm5
   161: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   165: c5d571f403               vpsllw $3, %ymm4, %ymm5
   170: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   174: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   178: c5d571f406               vpsllw $6, %ymm4, %ymm5
   183: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   187: c5d571f407               vpsllw $7, %ymm4, %ymm5
   192: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   196: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   200: 4d8b10                   movq (%r8), %r10
   203: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   209: c5dd71d301               vpsrlw $1, %ymm3, %ymm4
   214: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   218: 4d8b5008                 movq 8(%r8), %r10
   222: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   228: c5d571d301               vpsrlw $1, %ymm3, %ymm5
   233: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   237: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   241: c5f5efcc                 vpxor %ymm4, %ymm1, %ymm1
   245: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
   249: c5d571f401               vpsllw $1, %ymm4, %ymm5
   254: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   258: c5d571f402               vpsllw $2, %ymm4, %ymm5
   263: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   267: c5d571f403               vpsllw $3, %ymm4, %ymm5
   272: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   276: c5d571f407               vpsllw $7, %ymm4, %ymm5
   281: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   285: 4d8b10                   movq (%r8), %r10
   288: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   294: c5dd71d302               vpsrlw $2, %ymm3, %ymm4
   299: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   303: 4d8b5008                 movq 8(%r8), %r10
   307: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   313: c5d571d302               vpsrlw $2, %ymm3, %ymm5
   318: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   322: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   326: c5edefd4                 vpxor %ymm4, %ymm2, %ymm2
   330: c5d571f401               vpsllw $1, %ymm4, %ymm5
   335: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   339: c5d571f402               vpsllw $2, %ymm4, %ymm5
   344: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   348: 4d8b10                   movq (%r8), %r10
   351: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   357: c5dd71d303               vpsrlw $3, %ymm3, %ymm4
   362: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   366: 4d8b5008                 movq 8(%r8), %r10
   370: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   376: c5d571d303               vpsrlw $3, %ymm3, %ymm5
   381: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   385: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   389: c5d571f401               vpsllw $1, %ymm4, %ymm5
   394: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   398: c5d571f402               vpsllw $2, %ymm4, %ymm5
   403: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   407: c5d571f403               vpsllw $3, %ymm4, %ymm5
   412: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   416: 4d8b10                   movq (%r8), %r10
   419: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   425: c5dd71d304               vpsrlw $4, %ymm3, %ymm4
   430: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   434: 4d8b5008                 movq 8(%r8), %r10
   438: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   444: c5d571d304               vpsrlw $4, %ymm3, %ymm5
   449: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   453: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   457: c5d571f402               vpsllw $2, %ymm4, %ymm5
   462: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   466: c5d571f403               vpsllw $3, %ymm4, %ymm5
   471: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   475: c5d571f404               vpsllw $4, %ymm4, %ymm5
   480: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   484: 4d8b10                   movq (%r8), %r10
   487: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   493: c5dd71d305               vpsrlw $5, %ymm3, %ymm4
   498: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   502: 4d8b5008                 movq 8(%r8), %r10
   506: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   512: c5d571d305               vpsrlw $5, %ymm3, %ymm5
   517: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   521: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   525: c5d571f403               vpsllw $3, %ymm4, %ymm5
   530: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   534: c5d571f404               vpsllw $4, %ymm4, %ymm5
   539: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   543: c5d571f405               vpsllw $5, %ymm4, %ymm5
   548: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   552: 4d8b10                   movq (%r8), %r10
   555: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   561: c5dd71d306               vpsrlw $6, %ymm3, %ymm4
   566: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   570: 4d8b5008                 movq 8(%r8), %r10
   574: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   580: c5d571d306               vpsrlw $6, %ymm3, %ymm5
   585: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   589: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   593: c5d571f404               vpsllw $4, %ymm4, %ymm5
   598: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   602: c5d571f405               vpsllw $5, %ymm4, %ymm5
   607: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   611: c5d571f406               vpsllw $6, %ymm4, %ymm5
   616: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   620: 4d8b10                   movq (%r8), %r10
   623: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   629: c5dd71d307               vpsrlw $7, %ymm3, %ymm4
   634: c5dddbe6                 vpand %ymm6, %ymm4, %ymm4
   638: 4d8b5008                 movq 8(%r8), %r10
   642: c4817d6f1c1a             vmovdqa (%r10, %r11), %ymm3
   648: c5d571d307               vpsrlw $7, %ymm3, %ymm5
   653: c5d5dbee                 vpand %ymm6, %ymm5, %ymm5
   657: c5ddefe5                 vpxor %ymm5, %ymm4, %ymm4
   661: c5d571f405               vpsllw $5, %ymm4, %ymm5
   666: c5edefd5                 vpxor %ymm5, %ymm2, %ymm2
   670: c5d571f406               vpsllw $6, %ymm4, %ymm5
   675: c5f5efcd                 vpxor %ymm5, %ymm1, %ymm1
   679: c5d571f407               vpsllw $7, %ymm4, %ymm5
   684: c5fdefc5                 vpxor %ymm5, %ymm0, %ymm0
   688: 4d8b11                   movq (%r9), %r10
   691: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   697: 4d8b5108                 movq 8(%r9), %r10
   701: c4817d7f0c1a             vmovdqa %ymm1, (%r10, %r11)
   707: 4d8b5110                 movq 0x10(%r9), %r10
   711: c4817d7f141a             vmovdqa %ymm2, (%r10, %r11)
   717: 4983c320                 addq $0x20, %r11
   721: 4c3bd8                   cmpq %rax, %r11
   724: 0f8666fdffff             jbe 0x40
   730: 4c3bd9                   cmpq %rcx, %r11
   733: 0f8304010000             jae 0x3e7
   739: 4153                     pushq %r11
   741: 498d7f40                 leaq 0x40(%r15), %rdi
   745: 498db740010000           leaq 0x140(%r15), %rsi
   752: 4d8b11                   movq (%r9), %r10
   755: 4c8b1c24                 movq (%rsp), %r11
   759: 498b4010                 movq 0x10(%r8), %rax
   763: 420fb61c18               movzbl (%rax, %r11), %ebx
   768: 498b00                   movq (%r8), %rax
   771: 420fb61418               movzbl (%rax, %r11), %edx
   776: 498b4008                 movq 8(%r8), %rax
   780: 420fb60418               movzbl (%rax, %r11), %eax
   785: 32d0                     xorb %al, %dl
   787: 32da                     xorb %dl, %bl
   789: 43881c1a                 movb %bl, (%r10, %r11)
   793: 49ffc3                   incq %r11
   796: 4c3bd9                   cmpq %rcx, %r11
   799: 0f82d2ffffff             jb 0x2f7
   805: 4d8b5108                 movq 8(%r9), %r10
   809: 4c8b1c24                 movq (%rsp), %r11
   813: 498b4018                 movq 0x18(%r8), %rax
   817: 420fb61c18               movzbl (%rax, %r11), %ebx
   822: 498b00                   movq (%r8), %rax
   825: 420fb61418               movzbl (%rax, %r11), %edx
   830: 498b4008                 movq 8(%r8), %rax
   834: 420fb60418               movzbl (%rax, %r11), %eax
   839: 32d0                     xorb %al, %dl
   841: 4885d2                   testq %rdx, %rdx
   844: 0f8422000000             je 0x374
   850: 0fb60417                 movzbl (%rdi, %rdx), %eax
   854: 81c0e6000000             addl $0xe6, %eax
   860: 81f8ff000000             cmpl $0xff, %eax
   866: 0f8206000000             jb 0x36e
   872: 81c001ffffff             addl $0xffffff01, %eax
   878: 0fb60406                 movzbl (%rsi, %rax), %eax
   882: 32d8                     xorb %al, %bl
   884: 43881c1a                 movb %bl, (%r10, %r11)
   888: 49ffc3                   incq %r11
   891: 4c3bd9                   cmpq %rcx, %r11
   894: 0f82a9ffffff             jb 0x32d
   900: 4d8b5110                 movq 0x10(%r9), %r10
   904: 4c8b1c24                 movq (%rsp), %r11
   908: 498b4020                 movq 0x20(%r8), %rax
   912: 420fb61c18               movzbl (%rax, %r11), %ebx
   917: 498b00                   movq (%r8), %rax
   920: 420fb61418               movzbl (%rax, %r11), %edx
   925: 498b4008                 movq 8(%r8), %rax
   929: 420fb60418               movzbl (%rax, %r11), %eax
   934: 32d0                     xorb %al, %dl
   936: 4885d2                   testq %rdx, %rdx
   939: 0f8422000000             je 0x3d3
   945: 0fb60417                 movzbl (%rdi, %rdx), %eax
   949: 81c0cd000000             addl $0xcd, %eax
   955: 81f8ff000000             cmpl $0xff, %eax
   961: 0f8206000000             jb 0x3cd
   967: 81c001ffffff             addl $0xffffff01, %eax
   973: 0fb60406                 movzbl (%rsi, %rax), %eax
   977: 32d8                     xorb %al, %bl
   979: 43881c1a                 movb %bl, (%r10, %r11)
   983: 49ffc3                   incq %r11
   986: 4c3bd9                   cmpq %rcx, %r11
   989: 0f82a9ffffff             jb 0x38c
   995: 4883c408                 addq $8, %rsp
   999: 415f                     popq %r15
  1001: 5b                       popq %rbx
  1002: c5f877                   vzeroupper 
  1005: c3                       retq 
  1006: 0f0b                     ud2 
  1008: 0f0b                     ud2 
  1010: 0f0b                     ud2 
  1012: 0f0b                     ud2 
  1014: 0f0b                     ud2 
  1016: 0f0b                     ud2 
  1018: 0f0b                     ud2 
  1020: 0f0b                     ud2 
  1022: 0f0b                     ud2 
