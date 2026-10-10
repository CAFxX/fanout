; config_t1_728 target=sse twin=baseline topo=rs(6,3)/rec{0} op=recover U=1
; code_size=640 twins=[] uops_per_byte=0.4453125
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
    41: 4881e880000000           subq $0x80, %rax
    48: 0f82a3010000             jb 0x1d9
    54: 66660f1f840000000000     nopw (%rax, %rax)
    64: 66420f6f041e             movdqa (%rsi, %r11), %xmm0
    70: 66420f6f4c1e10           movdqa 0x10(%rsi, %r11), %xmm1
    77: 66420f6f541e20           movdqa 0x20(%rsi, %r11), %xmm2
    84: 66420f6f5c1e30           movdqa 0x30(%rsi, %r11), %xmm3
    91: 66420f6f641e40           movdqa 0x40(%rsi, %r11), %xmm4
    98: 66420f6f6c1e50           movdqa 0x50(%rsi, %r11), %xmm5
   105: 66420f6f741e60           movdqa 0x60(%rsi, %r11), %xmm6
   112: 66420f6f7c1e70           movdqa 0x70(%rsi, %r11), %xmm7
   119: 66420fef041f             pxor (%rdi, %r11), %xmm0
   125: 66420fef4c1f10           pxor 0x10(%rdi, %r11), %xmm1
   132: 66420fef541f20           pxor 0x20(%rdi, %r11), %xmm2
   139: 66420fef5c1f30           pxor 0x30(%rdi, %r11), %xmm3
   146: 66420fef641f40           pxor 0x40(%rdi, %r11), %xmm4
   153: 66420fef6c1f50           pxor 0x50(%rdi, %r11), %xmm5
   160: 66420fef741f60           pxor 0x60(%rdi, %r11), %xmm6
   167: 66420fef7c1f70           pxor 0x70(%rdi, %r11), %xmm7
   174: 66430fef041c             pxor (%r12, %r11), %xmm0
   180: 66430fef4c1c10           pxor 0x10(%r12, %r11), %xmm1
   187: 66430fef541c20           pxor 0x20(%r12, %r11), %xmm2
   194: 66430fef5c1c30           pxor 0x30(%r12, %r11), %xmm3
   201: 66430fef641c40           pxor 0x40(%r12, %r11), %xmm4
   208: 66430fef6c1c50           pxor 0x50(%r12, %r11), %xmm5
   215: 66430fef741c60           pxor 0x60(%r12, %r11), %xmm6
   222: 66430fef7c1c70           pxor 0x70(%r12, %r11), %xmm7
   229: 66430fef441d00           pxor (%r13, %r11), %xmm0
   236: 66430fef4c1d10           pxor 0x10(%r13, %r11), %xmm1
   243: 66430fef541d20           pxor 0x20(%r13, %r11), %xmm2
   250: 66430fef5c1d30           pxor 0x30(%r13, %r11), %xmm3
   257: 66430fef641d40           pxor 0x40(%r13, %r11), %xmm4
   264: 66430fef6c1d50           pxor 0x50(%r13, %r11), %xmm5
   271: 66430fef741d60           pxor 0x60(%r13, %r11), %xmm6
   278: 66430fef7c1d70           pxor 0x70(%r13, %r11), %xmm7
   285: 66430fef041e             pxor (%r14, %r11), %xmm0
   291: 66430fef4c1e10           pxor 0x10(%r14, %r11), %xmm1
   298: 66430fef541e20           pxor 0x20(%r14, %r11), %xmm2
   305: 66430fef5c1e30           pxor 0x30(%r14, %r11), %xmm3
   312: 66430fef641e40           pxor 0x40(%r14, %r11), %xmm4
   319: 66430fef6c1e50           pxor 0x50(%r14, %r11), %xmm5
   326: 66430fef741e60           pxor 0x60(%r14, %r11), %xmm6
   333: 66430fef7c1e70           pxor 0x70(%r14, %r11), %xmm7
   340: 4d8b5028                 movq 0x28(%r8), %r10
   344: 66430fef041a             pxor (%r10, %r11), %xmm0
   350: 66430fef4c1a10           pxor 0x10(%r10, %r11), %xmm1
   357: 66430fef541a20           pxor 0x20(%r10, %r11), %xmm2
   364: 66430fef5c1a30           pxor 0x30(%r10, %r11), %xmm3
   371: 66430fef641a40           pxor 0x40(%r10, %r11), %xmm4
   378: 66430fef6c1a50           pxor 0x50(%r10, %r11), %xmm5
   385: 66430fef741a60           pxor 0x60(%r10, %r11), %xmm6
   392: 66430fef7c1a70           pxor 0x70(%r10, %r11), %xmm7
   399: 4d8b11                   movq (%r9), %r10
   402: 66430f7f041a             movdqa %xmm0, (%r10, %r11)
   408: 66430f7f4c1a10           movdqa %xmm1, 0x10(%r10, %r11)
   415: 66430f7f541a20           movdqa %xmm2, 0x20(%r10, %r11)
   422: 66430f7f5c1a30           movdqa %xmm3, 0x30(%r10, %r11)
   429: 66430f7f641a40           movdqa %xmm4, 0x40(%r10, %r11)
   436: 66430f7f6c1a50           movdqa %xmm5, 0x50(%r10, %r11)
   443: 66430f7f741a60           movdqa %xmm6, 0x60(%r10, %r11)
   450: 66430f7f7c1a70           movdqa %xmm7, 0x70(%r10, %r11)
   457: 4981c380000000           addq $0x80, %r11
   464: 4c3bd8                   cmpq %rax, %r11
   467: 0f8667feffff             jbe 0x40
   473: 4c3bd9                   cmpq %rcx, %r11
   476: 0f835a000000             jae 0x23c
   482: 4153                     pushq %r11
   484: 4d8b11                   movq (%r9), %r10
   487: 4c8b1c24                 movq (%rsp), %r11
   491: 498b00                   movq (%r8), %rax
   494: 428a1418                 movb (%rax, %r11), %dl
   498: 8ada                     movb %dl, %bl
   500: 498b4008                 movq 8(%r8), %rax
   504: 428a1418                 movb (%rax, %r11), %dl
   508: 32da                     xorb %dl, %bl
   510: 498b4010                 movq 0x10(%r8), %rax
   514: 428a1418                 movb (%rax, %r11), %dl
   518: 32da                     xorb %dl, %bl
   520: 498b4018                 movq 0x18(%r8), %rax
   524: 428a1418                 movb (%rax, %r11), %dl
   528: 32da                     xorb %dl, %bl
   530: 498b4020                 movq 0x20(%r8), %rax
   534: 428a1418                 movb (%rax, %r11), %dl
   538: 32da                     xorb %dl, %bl
   540: 498b4028                 movq 0x28(%r8), %rax
   544: 428a1418                 movb (%rax, %r11), %dl
   548: 32da                     xorb %dl, %bl
   550: 43881c1a                 movb %bl, (%r10, %r11)
   554: 33db                     xorl %ebx, %ebx
   556: 49ffc3                   incq %r11
   559: 4c3bd9                   cmpq %rcx, %r11
   562: 0f82b3ffffff             jb 0x1eb
   568: 4883c408                 addq $8, %rsp
   572: 415e                     popq %r14
   574: 415d                     popq %r13
   576: 415c                     popq %r12
   578: 5b                       popq %rbx
   579: c3                       retq 
   580: 0f0b                     ud2 
   582: 0f0b                     ud2 
   584: 0f0b                     ud2 
   586: 0f0b                     ud2 
   588: 0f0b                     ud2 
   590: 0f0b                     ud2 
   592: 0f0b                     ud2 
   594: 0f0b                     ud2 
   596: 0f0b                     ud2 
   598: 0f0b                     ud2 
   600: 0f0b                     ud2 
   602: 0f0b                     ud2 
   604: 0f0b                     ud2 
   606: 0f0b                     ud2 
   608: 0f0b                     ud2 
   610: 0f0b                     ud2 
   612: 0f0b                     ud2 
   614: 0f0b                     ud2 
   616: 0f0b                     ud2 
   618: 0f0b                     ud2 
   620: 0f0b                     ud2 
   622: 0f0b                     ud2 
   624: 0f0b                     ud2 
   626: 0f0b                     ud2 
   628: 0f0b                     ud2 
   630: 0f0b                     ud2 
   632: 0f0b                     ud2 
   634: 0f0b                     ud2 
   636: 0f0b                     ud2 
   638: 0f0b                     ud2 
