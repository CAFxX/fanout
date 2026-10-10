; config_t3_2057 target=avx512-gfni twin=baseline topo=raid0(4)/enc op=encode U=1
; code_size=64 twins=[] uops_per_byte=0
     0: 53                       pushq %rbx
     1: 4c8bc7                   movq %rdi, %r8
     4: 4c8bce                   movq %rsi, %r9
     7: 488bca                   movq %rdx, %rcx
    10: 5b                       popq %rbx
    11: c5f877                   vzeroupper 
    14: c3                       retq 
    15: 0f0b                     ud2 
    17: 0f0b                     ud2 
    19: 0f0b                     ud2 
    21: 0f0b                     ud2 
    23: 0f0b                     ud2 
    25: 0f0b                     ud2 
    27: 0f0b                     ud2 
    29: 0f0b                     ud2 
    31: 0f0b                     ud2 
    33: 0f0b                     ud2 
    35: 0f0b                     ud2 
    37: 0f0b                     ud2 
    39: 0f0b                     ud2 
    41: 0f0b                     ud2 
    43: 0f0b                     ud2 
    45: 0f0b                     ud2 
    47: 0f0b                     ud2 
    49: 0f0b                     ud2 
    51: 0f0b                     ud2 
    53: 0f0b                     ud2 
    55: 0f0b                     ud2 
    57: 0f0b                     ud2 
    59: 0f0b                     ud2 
    61: 0f0b                     ud2 
    63: cc                       int3 
