package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int32 }
func genFn0() int {
    switch 32767 {
    case 4294967295:
        var v7 int32 = -793
        _ = v7
    case 0:
        v8 := []int{778, 4294967295, -146}
        _ = v8
    default:
        v9 := []int{32767}
        _ = v9
    }
    var v10 uint32 = 2147483648
    var v11 uint64 = 2
    _ = v10
    _ = v11
    return 9223372036854775807
}
func genFn1(p0 uint64, p1 uint64, p2 int32) int {
    v12 := GenS1{F0: true, F1: 2147483647, F2: 0}
    if (!true) {
        p0 = 0
        var v13 float64 = ((1.5 - float64((-2.25 * 0.0))) * 0.1)
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 9223372036854775807
}

func main() {
    v1 := []uint32{4294967295}
    var v2 float32 = (((-2.25 * 0.0) / (1.5 + 3.14159)) * (1e10 * (-2.25 * -2.25)))
    v3 := make(map[int]uint32)
    var v4 int32 = -1
    var v5 uint = 4294967295
    var v6 uint64 = 1
    for fi14 := 0; fi14 < 3; fi14++ {
        fmt.Println(v2)
        _ = fi14
    }
    v6 = uint64(v5)
    switch -827 {
    case -2147483648:
        var v15 int8 = 42
        _ = v15
    default:
        var v16 uint32 = 2147483647
        _ = v16
    }
    v17 := GenS1{F0: (false || false), F1: 2147483647, F2: 724}
    if false {
        v5 += v5
        v2 = -2.25
        var v18 bool = true
        _ = v18
    } else {
        var v19 uint16 = 65535
        fmt.Println(v2, v5, v6)
        _ = v19
    }
    v20 := v3[1]
    fmt.Println(v5, v4)
    fmt.Println(v20)
    for fi21 := 0; fi21 < 4; fi21++ {
        if (42 != 1) {
            fmt.Println(fi21, v4, v6)
        } else {
            v22 := GenS2{F0: 42, F1: v20, F2: (v4 % ((v4) | 1))}
            _ = v22
        }
        var v23 int32 = (^(v4))
        _ = fi21
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 int8 = 1
        _ = fi24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
