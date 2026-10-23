package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 int64 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]bool)
    var v3 uint64 = 1
    v4 := []string{"foo bar", "\n\t"}
    var v5 uint8 = 255
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 float32 = ((0.0 + (3.14159 + (0.0 * -2.25))) * 0.0)
        _ = fi6
        _ = v7
    }
    switch 42 {
    case -1:
        v8 := GenS1{F0: 1, F1: (v3 + (v3 + (v3 << ((1) & 63))))}
        _ = v8
    case 2147483647:
        fmt.Println(v1, v1)
    default:
        fmt.Println(v1, v1)
    }
    var v9 uint = 514
    if false {
        if ((1.5 > 1.5) && (0 >= 856)) {
            v3 -= v3
            var v10 uint8 = v5
            fmt.Println(v3, v1)
            _ = v10
        } else {
            v1 += v1
        }
        fmt.Println(v1, v5, v1)
    }
    var v11 uint64 = 782
    var v12 int32 = -2147483648
    v13 := GenS1{F0: (v12 % (((v12 ^ 2147483647)) | 1)), F1: (v3 & (v3 | v3))}
    var v14 uint16 = uint16(v3)
    var v15 bool = false
    for fi16 := 0; fi16 < 12; fi16++ {
        if false {
            var v17 int32 = 2147483647
            v18 := GenS2{F0: 42, F1: 42, F2: 514}
            _ = len(v2)
            _ = v17
            _ = v18
        }
        v9 -= (v9 + 140)
        _ = fi16
    }
    v9 = v9
    v14 = (v14 & (v14 >> (((v9 / ((798) | 1))) & 63)))
    v14 += v14
    v19 := GenS1{F0: (v12 * 1), F1: 2}
    if v15 {
        v11 -= (v11 - v3)
    }
    v20 := GenS1{F0: v12, F1: (v11 % ((1) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
