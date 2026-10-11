package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }
type GenS3 struct { F0 uint; F1 uint32; F2 uint8 }

func main() {
    var v1 uint = 0
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    v4 := make(map[uint32]bool)
    var v5 bool = true
    var v6 int8 = 42
    switch 990 {
    case 127:
        var v7 int8 = 1
        _ = v7
    case -795:
        var v8 bool = (972 < 4294967295)
        _ = v8
    case 9223372036854775807:
        switch 0 {
        case -9223372036854775808:
            if v5 {
                var v9 uint32 = 2147483648
                _ = v9
            } else {
                var v10 float32 = 0.0
                var v11 int32 = 2147483647
                fmt.Println(v10, v1, v6)
                _ = v10
                _ = v11
            }
        case 2147483647:
            v12 := []int{532, -512, 194, 235}
            _ = v12
        }
    default:
        v13 := []int{128, -1}
        _ = v13
    }
    v5 = (0 >= 127)
    v14 := GenS1{F0: 18446744073709551615, F1: 65535, F2: 114}
    if (494 == uint32(v1)) {
        v15 := []int{}
        _ = v15
    }
    if true {
        if v5 {
            v16 := []int{339, -2147483648, 32767, 0}
            fmt.Println(v5)
            _ = v16
        }
        fmt.Println(v6, v5)
    }
    v1 = 689
    v17 := GenS2{F0: (^(v1)), F1: 4294967295, F2: (v5 || (245 == 0))}
    var v18 uint8 = 1
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 uint16 = 42
        var v21 uint8 = (^(v18))
        switch fi19 {
        case 4294967295:
            v22 := []int{fi19, -941, fi19, (fi19 / ((-2147483648) | 1))}
            _ = v22
        case 2147483647:
            if v5 {
                v23 := GenS1{F0: 2, F1: (v20 >> ((v1) & 63)), F2: 42}
                _ = v23
            }
        default:
            if true {
                v6 = v6
                if ((true || true) || v5) {
                    fmt.Println(v18)
                    if ((3.14159 < (-2.25 - 3.14159)) || ((fi19 & (fi19 << ((v1) & 63))) >= fi19)) {
                        var v24 bool = v5
                        _ = v24
                    }
                    v25 := GenS2{F0: v1, F1: 339, F2: ((1 != 32767) && (1 < 65535))}
                    _ = v25
                }
            } else {
                v1 += 800
                v26 := GenS1{F0: 42, F1: uint16(v18), F2: 86}
                for fi27 := 0; fi27 < 9; fi27++ {
                    v28 := []int{(fi27 << (((v1 ^ uint(v21))) & 63))}
                    var v29 bool = v5
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            }
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        v31 := v4[2147483647]
        var v32 int16 = 32767
        v33 := GenS3{F0: 382, F1: 4294967295, F2: (^(v18))}
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
