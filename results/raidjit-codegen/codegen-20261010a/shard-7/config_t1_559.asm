; config_t1_559 target=avx2 twin=baseline topo=raid6(8)/upd{0,1} op=parity_update U=8
; code_size=1088 twins=[] uops_per_byte=0.875
     0: 53                       pushq %rbx
     1: 4157                     pushq %r15
     3: 4c8bc7                   movq %rdi, %r8
     6: 4c8bce                   movq %rsi, %r9
     9: 488bca                   movq %rdx, %rcx
    12: 4c8d3d2d040000           leaq 0x42d(%rip), %r15
    19: 4d33db                   xorq %r11, %r11
    22: c4417d6f37               vmovdqa (%r15), %ymm14
    27: 488bc1                   movq %rcx, %rax
    30: 4881e880000000           subq $0x80, %rax
    37: 0f82e7020000             jb 0x312
    43: 666666666666660f1f840000000000 nopw (%rax, %rax)
    58: 660f1f440000             nopw (%rax, %rax)
    64: 4d8b5020                 movq 0x20(%r8), %r10
    68: c4817d6f041a             vmovdqa (%r10, %r11), %ymm0
    74: c4817d6f4c1a20           vmovdqa 0x20(%r10, %r11), %ymm1
    81: c4817d6f541a40           vmovdqa 0x40(%r10, %r11), %ymm2
    88: c4817d6f5c1a60           vmovdqa 0x60(%r10, %r11), %ymm3
    95: 4d8b5028                 movq 0x28(%r8), %r10
    99: c4817d6f241a             vmovdqa (%r10, %r11), %ymm4
   105: c4817d6f6c1a20           vmovdqa 0x20(%r10, %r11), %ymm5
   112: c4817d6f741a40           vmovdqa 0x40(%r10, %r11), %ymm6
   119: c4817d6f7c1a60           vmovdqa 0x60(%r10, %r11), %ymm7
   126: 4d8b10                   movq (%r8), %r10
   129: 498b5010                 movq 0x10(%r8), %rdx
   133: c4017d6f041a             vmovdqa (%r10, %r11), %ymm8
   139: c4213def041a             vpxor (%rdx, %r11), %ymm8, %ymm8
   145: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   150: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   156: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   161: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   166: c4417d6f6f40             vmovdqa 0x40(%r15), %ymm13
   172: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   177: c4417d6faf80000000       vmovdqa 0x80(%r15), %ymm13
   186: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   191: c4c15defe3               vpxor %ymm11, %ymm4, %ymm4
   196: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   201: c4017d6f441a20           vmovdqa 0x20(%r10, %r11), %ymm8
   208: c4213def441a20           vpxor 0x20(%rdx, %r11), %ymm8, %ymm8
   215: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   220: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   226: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   231: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   236: c4417d6f6f40             vmovdqa 0x40(%r15), %ymm13
   242: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   247: c4417d6faf80000000       vmovdqa 0x80(%r15), %ymm13
   256: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   261: c4c155efeb               vpxor %ymm11, %ymm5, %ymm5
   266: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   271: c4017d6f441a40           vmovdqa 0x40(%r10, %r11), %ymm8
   278: c4213def441a40           vpxor 0x40(%rdx, %r11), %ymm8, %ymm8
   285: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   290: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   296: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   301: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   306: c4417d6f6f40             vmovdqa 0x40(%r15), %ymm13
   312: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   317: c4417d6faf80000000       vmovdqa 0x80(%r15), %ymm13
   326: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   331: c4c14deff3               vpxor %ymm11, %ymm6, %ymm6
   336: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   341: c4017d6f441a60           vmovdqa 0x60(%r10, %r11), %ymm8
   348: c4213def441a60           vpxor 0x60(%rdx, %r11), %ymm8, %ymm8
   355: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   360: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   366: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   371: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
   376: c4417d6f6f40             vmovdqa 0x40(%r15), %ymm13
   382: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   387: c4417d6faf80000000       vmovdqa 0x80(%r15), %ymm13
   396: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   401: c4c145effb               vpxor %ymm11, %ymm7, %ymm7
   406: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   411: 4d8b5008                 movq 8(%r8), %r10
   415: 498b5018                 movq 0x18(%r8), %rdx
   419: c4017d6f041a             vmovdqa (%r10, %r11), %ymm8
   425: c4213def041a             vpxor (%rdx, %r11), %ymm8, %ymm8
   431: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   436: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   442: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   447: c4c17defc0               vpxor %ymm8, %ymm0, %ymm0
   452: c4417d6fafc0000000       vmovdqa 0xc0(%r15), %ymm13
   461: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   466: c4417d6faf00010000       vmovdqa 0x100(%r15), %ymm13
   475: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   480: c4c15defe3               vpxor %ymm11, %ymm4, %ymm4
   485: c4c15defe4               vpxor %ymm12, %ymm4, %ymm4
   490: c4017d6f441a20           vmovdqa 0x20(%r10, %r11), %ymm8
   497: c4213def441a20           vpxor 0x20(%rdx, %r11), %ymm8, %ymm8
   504: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   509: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   515: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   520: c4c175efc8               vpxor %ymm8, %ymm1, %ymm1
   525: c4417d6fafc0000000       vmovdqa 0xc0(%r15), %ymm13
   534: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   539: c4417d6faf00010000       vmovdqa 0x100(%r15), %ymm13
   548: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   553: c4c155efeb               vpxor %ymm11, %ymm5, %ymm5
   558: c4c155efec               vpxor %ymm12, %ymm5, %ymm5
   563: c4017d6f441a40           vmovdqa 0x40(%r10, %r11), %ymm8
   570: c4213def441a40           vpxor 0x40(%rdx, %r11), %ymm8, %ymm8
   577: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   582: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   588: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   593: c4c16defd0               vpxor %ymm8, %ymm2, %ymm2
   598: c4417d6fafc0000000       vmovdqa 0xc0(%r15), %ymm13
   607: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   612: c4417d6faf00010000       vmovdqa 0x100(%r15), %ymm13
   621: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   626: c4c14deff3               vpxor %ymm11, %ymm6, %ymm6
   631: c4c14deff4               vpxor %ymm12, %ymm6, %ymm6
   636: c4017d6f441a60           vmovdqa 0x60(%r10, %r11), %ymm8
   643: c4213def441a60           vpxor 0x60(%rdx, %r11), %ymm8, %ymm8
   650: c4413ddbce               vpand %ymm14, %ymm8, %ymm9
   655: c4c12d71d004             vpsrlw $4, %ymm8, %ymm10
   661: c4412ddbd6               vpand %ymm14, %ymm10, %ymm10
   666: c4c165efd8               vpxor %ymm8, %ymm3, %ymm3
   671: c4417d6fafc0000000       vmovdqa 0xc0(%r15), %ymm13
   680: c4421500d9               vpshufb %ymm9, %ymm13, %ymm11
   685: c4417d6faf00010000       vmovdqa 0x100(%r15), %ymm13
   694: c4421500e2               vpshufb %ymm10, %ymm13, %ymm12
   699: c4c145effb               vpxor %ymm11, %ymm7, %ymm7
   704: c4c145effc               vpxor %ymm12, %ymm7, %ymm7
   709: 4d8b11                   movq (%r9), %r10
   712: c4817d7f041a             vmovdqa %ymm0, (%r10, %r11)
   718: c4817d7f4c1a20           vmovdqa %ymm1, 0x20(%r10, %r11)
   725: c4817d7f541a40           vmovdqa %ymm2, 0x40(%r10, %r11)
   732: c4817d7f5c1a60           vmovdqa %ymm3, 0x60(%r10, %r11)
   739: 4d8b5108                 movq 8(%r9), %r10
   743: c4817d7f241a             vmovdqa %ymm4, (%r10, %r11)
   749: c4817d7f6c1a20           vmovdqa %ymm5, 0x20(%r10, %r11)
   756: c4817d7f741a40           vmovdqa %ymm6, 0x40(%r10, %r11)
   763: c4817d7f7c1a60           vmovdqa %ymm7, 0x60(%r10, %r11)
   770: 4981c380000000           addq $0x80, %r11
   777: 4c3bd8                   cmpq %rax, %r11
   780: 0f862efdffff             jbe 0x40
   786: 4c3bd9                   cmpq %rcx, %r11
   789: 0f83fd000000             jae 0x418
   795: 4153                     pushq %r11
   797: 498dbf40010000           leaq 0x140(%r15), %rdi
   804: 498db740020000           leaq 0x240(%r15), %rsi
   811: 4d8b11                   movq (%r9), %r10
   814: 4c8b1c24                 movq (%rsp), %r11
   818: 498b4020                 movq 0x20(%r8), %rax
   822: 420fb61c18               movzbl (%rax, %r11), %ebx
   827: 498b00                   movq (%r8), %rax
   830: 420fb61418               movzbl (%rax, %r11), %edx
   835: 498b4010                 movq 0x10(%r8), %rax
   839: 420fb60418               movzbl (%rax, %r11), %eax
   844: 32d0                     xorb %al, %dl
   846: 32da                     xorb %dl, %bl
   848: 498b4008                 movq 8(%r8), %rax
   852: 420fb61418               movzbl (%rax, %r11), %edx
   857: 498b4018                 movq 0x18(%r8), %rax
   861: 420fb60418               movzbl (%rax, %r11), %eax
   866: 32d0                     xorb %al, %dl
   868: 32da                     xorb %dl, %bl
   870: 43881c1a                 movb %bl, (%r10, %r11)
   874: 49ffc3                   incq %r11
   877: 4c3bd9                   cmpq %rcx, %r11
   880: 0f82bcffffff             jb 0x332
   886: 4d8b5108                 movq 8(%r9), %r10
   890: 4c8b1c24                 movq (%rsp), %r11
   894: 498b4028                 movq 0x28(%r8), %rax
   898: 420fb61c18               movzbl (%rax, %r11), %ebx
   903: 498b00                   movq (%r8), %rax
   906: 420fb61418               movzbl (%rax, %r11), %edx
   911: 498b4010                 movq 0x10(%r8), %rax
   915: 420fb60418               movzbl (%rax, %r11), %eax
   920: 32d0                     xorb %al, %dl
   922: 4885d2                   testq %rdx, %rdx
   925: 0f8422000000             je 0x3c5
   931: 0fb60417                 movzbl (%rdi, %rdx), %eax
   935: 81c0e6000000             addl $0xe6, %eax
   941: 81f8ff000000             cmpl $0xff, %eax
   947: 0f8206000000             jb 0x3bf
   953: 81c001ffffff             addl $0xffffff01, %eax
   959: 0fb60406                 movzbl (%rsi, %rax), %eax
   963: 32d8                     xorb %al, %bl
   965: 498b4008                 movq 8(%r8), %rax
   969: 420fb61418               movzbl (%rax, %r11), %edx
   974: 498b4018                 movq 0x18(%r8), %rax
   978: 420fb60418               movzbl (%rax, %r11), %eax
   983: 32d0                     xorb %al, %dl
   985: 4885d2                   testq %rdx, %rdx
   988: 0f8422000000             je 0x404
   994: 0fb60417                 movzbl (%rdi, %rdx), %eax
   998: 81c0cd000000             addl $0xcd, %eax
  1004: 81f8ff000000             cmpl $0xff, %eax
  1010: 0f8206000000             jb 0x3fe
  1016: 81c001ffffff             addl $0xffffff01, %eax
  1022: 0fb60406                 movzbl (%rsi, %rax), %eax
  1026: 32d8                     xorb %al, %bl
  1028: 43881c1a                 movb %bl, (%r10, %r11)
  1032: 49ffc3                   incq %r11
  1035: 4c3bd9                   cmpq %rcx, %r11
  1038: 0f826affffff             jb 0x37e
  1044: 4883c408                 addq $8, %rsp
  1048: 415f                     popq %r15
  1050: 5b                       popq %rbx
  1051: c5f877                   vzeroupper 
  1054: c3                       retq 
  1055: 0f0b                     ud2 
  1057: 0f0b                     ud2 
  1059: 0f0b                     ud2 
  1061: 0f0b                     ud2 
  1063: 0f0b                     ud2 
  1065: 0f0b                     ud2 
  1067: 0f0b                     ud2 
  1069: 0f0b                     ud2 
  1071: 0f0b                     ud2 
  1073: 0f0b                     ud2 
  1075: 0f0b                     ud2 
  1077: 0f0b                     ud2 
  1079: 0f0b                     ud2 
  1081: 0f0b                     ud2 
  1083: 0f0b                     ud2 
  1085: 0f0b                     ud2 
  1087: cc                       int3 
