package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 string; F1 int32 }
func genFn0(p0 int32, p1 uint32, p2 int16) int16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: -1, F1: 812}
        p0 += 335
        v9 := []int{(-(fi7)), fi7, -2147483648}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (p2 & (p2 - p2))
}
func genFn1(p0 uint8, p1 uint32) int8 {
    v11 := []int{-949, -106}
    if (!((1.5 <= 3.14159) || true)) {
        v11 = append(v11, 1)
        p1 = p1
    }
    var v12 int32 = 818
    v13 := GenS2{F0: "\n\t", F1: v12}
    var v14 uint8 = (p0 >> ((0) & 63))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 126
}

func main() {
    var v1 int = -465
    var v2 int16 = 32767
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    var v5 int = -808
    v6 := make(map[uint32]string)
    v2 = (v2 | -32768)
    v2 = (-(v2))
    switch v5 {
    case 127:
        switch "\n\t" {
        case "\n\t":
            v2 += genFn0(350, 376, (v2 ^ v2))
        }
    default:
        v5 -= (v5 + (v5 % (((v1 % ((v5) | 1))) | 1)))
    }
    if (true && true) {
        var v15 int64 = 0
        _ = v15
    }
    if (!false) {
        var v16 int16 = genFn0(2147483647, 2147483648, (-(v2)))
        _ = v16
    } else {
        fmt.Println(v2, v5)
        var v17 uint16 = 32767
        _ = v17
    }
    v18 := []int{(v1 | v1), (v5 + (-(v1))), v5, v5}
    fmt.Println(v1, v5)
    _ = len(v18)
    if (-2.25 >= ((1e10 / -2.25) * 1e10)) {
        for fi19 := 0; fi19 < 7; fi19++ {
            fmt.Println(v5, v5, fi19)
            var v20 uint = 65535
            var v21 int16 = 32767
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        _ = len(v4)
        var v22 int16 = genFn0(1, 65535, (v2 >> ((0) & 63)))
        _ = v22
    }
    var v23 float32 = (1.5 - 0.0)
    fmt.Println(v1, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
