package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 bool; F2 bool }
func genFn0(p0 uint16, p1 int, p2 int64) int64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        switch 127 {
        case 32768:
            p2 -= -1
        case -2147483648:
            p1 += -109
        case 9223372036854775807:
            v8 := GenS1{F0: 2}
            _ = v8
        default:
            p0 = (p0 % (((p0 | (p0 + p0))) | 1))
        }
        _ = fi7
    }
    var v9 string = "\n\t"
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p2
}
func genFn1(p0 uint16, p1 bool) int32 {
    fmt.Println(p1)
    p1 = true
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 bool = (65535 <= 0)
    var v4 int8 = 0
    var v5 float64 = float64(-2.25)
    v6 := make(map[int]bool)
    var v10 uint16 = 2
    switch -2147483648 {
    case 32768:
        v11 := GenS1{F0: 42}
        _ = v11
    case 9223372036854775807:
        fmt.Println(v10)
    case 1:
        var v12 uint32 = 2
        _ = v12
    }
    v5 = v5
    v13 := GenS2{F0: 18446744073709551615, F1: (!(2 <= 1)), F2: (false && v3)}
    _ = len(v2)
    var v14 float64 = ((-2.25 - 3.14159) * 0.0)
    v15 := []int{-2147483648, 2147483647, -733, -564}
    v4 = v4
    var v16 bool = ((v3 || ((v4 ^ v4) >= -128)) || v3)
    v17 := GenS2{F0: 1, F1: (18446744073709551615 >= 65535), F2: (572 > 449)}
    var v18 uint16 = (v10 ^ (v10 >> ((2) & 63)))
    v19 := GenS2{F0: 42, F1: v3, F2: (true && v3)}
    var v20 int = 128
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 uint32 = 2
        fmt.Println(v10, v16)
        fmt.Println(v5, v20)
        _ = fi21
        _ = v22
    }
    v5 = (-2.25 / v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
