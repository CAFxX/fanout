package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 0
    var v3 int16 = -215
    var v4 int = 32768
    var v5 int = (^(v4))
    v2 = (-(v2))
    var v6 uint32 = 42
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint16 = 65535
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            v6 -= v6
            v5 += v4
            v11 := []int{(^(fi10))}
            v12 := GenS1{F0: v6, F1: 0, F2: (^(v3))}
            _ = fi10
            _ = v11
            _ = v12
        }
        v2 += (v2 / (((v2 ^ 42)) | 1))
        var v13 float32 = 0.0
        _ = fi9
        _ = v13
    }
    v3 = (v3 + (v3 ^ 1))
    v14 := v1[817]
    if true {
        var v15 uint32 = 853
        _ = v15
    }
    v16 := GenS2{F0: (v6 | (v6 >> ((18446744073709551615) & 63))), F1: (-(v2)), F2: 4294967295}
    var v17 int16 = (v3 ^ v3)
    fmt.Println(v6)
    fmt.Println(v6, v2)
    v17 += v3
    fmt.Println(v5, v6, v4)
    v18 := GenS1{F0: 1, F1: -9223372036854775808, F2: 0}
    if (((((0.0 / -2.25) / 3.14159) / 0.1) == float64((-2.25 / -2.25))) && false) {
        var v19 uint8 = 13
        _ = v19
    }
    var v20 int32 = 185
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
