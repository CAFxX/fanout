package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int64 }
func genFn0() bool {
    v7 := GenS1{F0: (0 == 42), F1: 18446744073709551615, F2: 813}
    var v8 uint = 18446744073709551615
    fmt.Println(v8)
    var v9 bool = (-32768 <= 32767)
    _ = v7
    _ = v8
    _ = v9
    return v9
}
func genFn1(p0 string, p1 int32, p2 bool) int64 {
    if (!(-520 < 9223372036854775807)) {
        var v10 float64 = 0.0
        _ = v10
    }
    p2 = (32767 != 1)
    switch -461 {
    case 1:
        var v11 float64 = 1e10
        _ = v11
    case 42:
        fmt.Println(p2, p2, p2)
    default:
        var v12 uint32 = 2147483647
        _ = v12
    }
    p0 = p0
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return -328
}

func main() {
    var v1 uint16 = 266
    v2 := []uint32{4294967295, 4294967295}
    v3 := []int{127, -871, 32767, -522}
    v4 := make(map[int]int)
    var v5 uint8 = 128
    v6 := []int{}
    var v13 int64 = 325
    if (v5 == (v5 / ((v5) | 1))) {
        v14 := GenS1{F0: true, F1: 0, F2: 65535}
        v13 -= genFn1("a", 635, true)
        _ = v14
    }
    var v15 float64 = (float64(1e10) / 1e10)
    var v16 int32 = 1
    var v17 uint16 = (v1 * (v1 >> ((2) & 63)))
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 uint16 = v17
        for fi20 := 0; fi20 < 4; fi20++ {
            fmt.Println(v16, v15, v17)
            _ = fi20
        }
        fmt.Println(v17)
        _ = fi18
        _ = v19
    }
    v21 := GenS2{F0: 1, F1: (v13 / ((42) | 1))}
    for fi22 := 0; fi22 < 15; fi22++ {
        v1 = uint16(v5)
        _ = fi22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 uint32 = 4294967295
        v15 = (v15 / v15)
        var v25 int = fi23
        _ = fi23
        _ = v24
        _ = v25
    }
    v13 += genFn1("\n\t", (v16 ^ (v16 * v16)), false)
    var v26 bool = true
    v4[-69] = 23
    var v27 uint8 = v5
    v15 = float64(1.5)
    v28 := v4[127]
    v13 = (v13 + v13)
    v28 += -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
