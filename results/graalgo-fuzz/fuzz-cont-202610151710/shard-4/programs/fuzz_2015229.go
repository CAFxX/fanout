package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0() uint64 {
    v7 := GenS1{F0: 1, F1: 1, F2: 1}
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    v9 := []int{-693, -2147483648, 887}
    if false {
        var v10 uint32 = 65535
        _ = v10
    }
    _ = v7
    _ = v9
    return 2
}
func genFn1(p0 uint64, p1 int64, p2 uint16) int32 {
    p0 = 2
    if ((0.1 != 1.5) && (p0 < (p0 / (((p0 - (p0 * p0))) | 1)))) {
        var v11 float32 = ((0.1 + 0.0) * (float32(1e10) / 1e10))
        _ = v11
    } else {
        var v12 int16 = 861
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = 3.14159
    v3 := make(map[string]bool)
    var v4 uint64 = 4294967295
    v5 := []int{}
    var v6 uint = 59
    fmt.Println(v1, v2)
    switch 42 {
    case 0:
        for fi13 := 0; fi13 < 19; fi13++ {
            fmt.Println(v4, v6)
            _ = fi13
        }
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v1 += (v1 | v1)
    v14 := GenS1{F0: -9223372036854775808, F1: -388, F2: v4}
    v15 := GenS2{F0: 1000}
    v4 = v4
    v6 = v6
    v16 := GenS3{F0: 42}
    if false {
        v2 = float64(1e10)
        var v17 int8 = -80
        switch 127 {
        case -2147483648:
            var v18 int64 = -9223372036854775808
            _ = v18
        case -694:
            v19 := GenS4{F0: (v4 | v4), F1: -32768, F2: 2}
            _ = v19
        case 9223372036854775807:
            fmt.Println(v17, v4, v6)
        }
        _ = v17
    } else {
        v6 = (v6 / ((4294967295) | 1))
        v1 = (v1 << (((^(v6))) & 63))
        var v20 uint32 = 4294967295
        _ = v20
    }
    v21 := GenS3{F0: 9223372036854775807}
    v1 = 32767
    v6 += (^(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
