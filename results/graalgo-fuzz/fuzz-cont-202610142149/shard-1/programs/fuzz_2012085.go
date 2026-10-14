package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 int16 = 910
    var v2 uint16 = 42
    var v3 uint32 = 1
    var v4 uint64 = 951
    var v5 int = 32767
    var v6 uint32 = 2147483647
    if false {
        fmt.Println(v5, v4, v4)
    }
    v1 += (v1 & (v1 - (v1 >> ((18446744073709551615) & 63))))
    if false {
        var v7 int64 = 0
        v8 := GenS1{F0: 4294967295, F1: v4, F2: (v1 / ((v1) | 1))}
        v9 := GenS1{F0: v4, F1: v4, F2: int16(v3)}
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(v6, v2)
        v11 := GenS2{F0: 0, F1: uint8(v2), F2: 4294967295}
        var v12 float64 = 0.1
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS3{F0: v4, F1: (v6 - (v6 * v3))}
    v14 := []int{(v5 << ((583) & 63)), (^(v5)), -9223372036854775808, (-(v5))}
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 15; fi16++ {
        if len(v14) > 0 {
            v14[0] = (v5 / ((1) | 1))
        }
        v4 -= (v4 / (((^(v4))) | 1))
        v14 = append(v14, int(v1))
        _ = fi16
    }
    _ = len(v14)
    var v17 int8 = 1
    var v18 int = v5
    v19 := GenS2{F0: -2147483648, F1: 128, F2: 652}
    var v20 int8 = (v17 ^ (v17 % ((-34) | 1)))
    var v21 int = 127
    for fi22 := 0; fi22 < 7; fi22++ {
        fmt.Println(v17, v20)
        _ = len(v14)
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
