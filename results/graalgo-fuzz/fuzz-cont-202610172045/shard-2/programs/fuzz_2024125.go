package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 1
    var v3 uint32 = 842
    var v4 float32 = 1.5
    v5 := make(map[uint32]bool)
    v6 := []uint32{v3, (v3 & (v3 / (((v3 & (v3 ^ v3))) | 1))), v3}
    v7 := GenS1{F0: v3, F1: 1}
    if true {
        var v8 int8 = -128
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: (^(v3)), F1: fi9}
        _ = fi9
        _ = v10
    }
    v2 = v2
    if ((v2 | v2) != (v2 << ((65535) & 63))) {
        var v11 bool = (18446744073709551615 != 731)
        v12 := GenS1{F0: v3, F1: 416}
        _ = v11
        _ = v12
    }
    var v13 float64 = 1.5
    var v14 bool = true
    for fi15 := 0; fi15 < 1; fi15++ {
        _ = len(v1)
        v13 = (float64(v4) / ((v13 * 1e10) + 0.0))
        v6 = append(v6, (v3 | v3))
        _ = fi15
    }
    v4 -= ((v4 - v4) / v4)
    v3 += v3
    v14 = (!v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
