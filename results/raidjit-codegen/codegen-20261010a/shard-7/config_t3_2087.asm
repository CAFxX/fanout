; config_t3_2087 target=rvv twin=baseline topo=raid0(4)/enc op=encode U=1
; code_size=32 twins=[] uops_per_byte=0
     0: 63140600                 bnez a2, 8
     4: 6f00c000                 j 0x10
     8: 970f0000                 auipc t6, 0
    12: 938f8f01                 addi t6, t6, 0x18
    16: 67800000                 ret 
    20: 0000                     unimp 
    22: 0000                     unimp 
    24: 0000                     unimp 
    26: 0000                     unimp 
    28: 0000                     unimp 
    30: 0000                     unimp 
