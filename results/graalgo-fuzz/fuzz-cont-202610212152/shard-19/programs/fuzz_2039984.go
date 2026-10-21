package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0(p0 int16, p1 int) uint64 {
    switch (p1 % ((p1) | 1)) {
    case 923:
        v7 := GenS1{F0: 397, F1: 100}
        _ = v7
    case 790:
        v8 := GenS1{F0: -40, F1: 0}
        _ = v8
    case 2147483647:
        fmt.Println(p0, p0)
    default:
        var v9 float64 = (1e10 - ((1e10 * (3.14159 * (0.0 / 0.1))) - 0.1))
        _ = v9
    }
    var v10 uint16 = 32767
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return 754
}

func main() {
    var v1 int64 = -628
    v2 := make(map[string]string)
    v3 := []uint32{}
    var v4 uint = 511
    v5 := make(map[uint32]string)
    v6 := make(map[int]int)
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v13 int16 = 0
    var v14 uint32 = 2147483648
    v14 = uint32(v4)
    var v15 float64 = 3.14159
    v15 = (v15 / 0.1)
    var v16 bool = true
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v16)
    _ = len(v5)
    v17 := GenS1{F0: 1, F1: 115}
    v15 -= v15
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
