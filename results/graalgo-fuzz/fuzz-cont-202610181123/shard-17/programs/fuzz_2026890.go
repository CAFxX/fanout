package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint8, p1 bool, p2 uint64) int64 {
    var v7 uint = 361
    var v8 float64 = ((-2.25 - float64(-2.25)) + 0.0)
    v9 := []int{-161, int(p2), -1, -340}
    p1 = p1
    var v10 int8 = 101
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint8, p1 int32) bool {
    for fi11 := 0; fi11 < 7; fi11++ {
        p1 += p1
        fmt.Println(p1)
        _ = fi11
    }
    v12 := []int{-581, 256, 532, 127}
    var v13 uint32 = 463
    v12 = append(v12, 128)
    for fi14 := 0; fi14 < 13; fi14++ {
        p1 = p1
        for fi15 := 0; fi15 < 6; fi15++ {
            p1 = p1
            _ = len(v12)
            fmt.Println(fi15, fi15)
            v16 := GenS1{F0: p1, F1: 1, F2: uint(v13)}
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return (!false)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 1
    var v3 uint32 = 4294967295
    var v4 uint8 = 0
    var v5 int8 = 123
    v6 := []string{"hello", "\n\t", ""}
    var v17 int = 2147483647
    for fi18 := 0; fi18 < 10; fi18++ {
        fmt.Println(v5)
        var v19 uint64 = 0
        switch (fi18 - (fi18 << ((198) & 63))) {
        case -1:
            v20 := GenS2{F0: v3}
            _ = v20
        case -765:
            v21 := v1[(v3 & (v3 + (^(v3))))]
            _ = v21
        default:
            v6 = append(v6, "a")
        }
        _ = fi18
        _ = v19
    }
    v17 = int(v4)
    v22 := GenS3{F0: v4, F1: 42}
    if (793 < 1) {
        v4 -= (v4 | 42)
        fmt.Println(v3, v5)
    }
    var v23 uint8 = (v4 - (v4 | (v4 - (v4 ^ 57))))
    v24 := GenS2{F0: (v3 - (v3 * (v3 << ((0) & 63))))}
    fmt.Println(v3, v3)
    v17 = 800
    var v25 int64 = 0
    fmt.Println(v3, v25)
    v4 -= (v4 >> ((2) & 63))
    if true {
        fmt.Println(v23)
        v26 := v1[v3]
        _ = v26
    }
    _ = len(v1)
    var v27 uint8 = uint8(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
