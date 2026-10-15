package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int64 }

func main() {
    v1 := []int{32768, -9223372036854775808, -2147483648, -9223372036854775808}
    v2 := make(map[uint32]int)
    v3 := []int{}
    var v4 float32 = float32(0.1)
    var v5 float64 = float64(v4)
    v4 += v4
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint8 = 128
        fmt.Println(v4)
        var v8 uint32 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        if false {
            v10 := GenS1{F0: 532, F1: 0, F2: 355}
            var v11 int32 = -263
            _ = v10
            _ = v11
        } else {
            var v12 int = (fi9 ^ fi9)
            _ = v12
        }
        v5 = 3.14159
        var v13 uint32 = 2147483647
        _ = fi9
        _ = v13
    }
    if (!false) {
        v4 -= v4
        v14 := v2[2147483647]
        v14 = 509
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        switch (fi15 * 2147483647) {
        case -2147483648:
            if false {
                v16 := GenS1{F0: 1, F1: 136, F2: -1}
                _ = v16
            }
        case 127:
            var v17 int = (^(fi15))
            _ = v17
        default:
            delete(v2, 837)
        }
        var v18 float32 = 1e10
        _ = fi15
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        fmt.Println(v5)
        _ = fi19
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v20 := GenS1{F0: 18446744073709551615, F1: 53, F2: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
