package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
func genFn0(p0 uint, p1 uint64) int8 {
    p1 -= (p1 / (((^(p1))) | 1))
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(p0, p0, fi7)
        var v8 uint = 18446744073709551615
        _ = fi7
        _ = v8
    }
    switch 127 {
    case 32768:
        p1 = p1
    case -678:
        if (((0.1 + float64(1.5)) + (0.0 * float64((1.5 / 1.0)))) < -2.25) {
            v9 := []int{-150, 0, 127}
            _ = v9
        } else {
            v10 := GenS1{F0: 85, F1: 566}
            fmt.Println(p1, p0)
            _ = v10
        }
    case -9223372036854775808:
        var v11 uint = p0
        _ = v11
    }
    _ = p0
    _ = p1
    return 107
}
func genFn1(p0 uint, p1 int16, p2 uint16) bool {
    var v12 uint16 = 0
    v13 := GenS1{F0: -124, F1: int32(p0)}
    switch "a" {
    case "":
        p1 = (p1 | p1)
    case "foo bar":
        p2 = (v12 * v12)
    default:
        var v14 int64 = -9223372036854775808
        _ = v14
    }
    p1 += p1
    var v15 int8 = -22
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v15
    return false
}

func main() {
    var v1 uint64 = 42
    var v2 uint8 = uint8(v1)
    var v3 int16 = -32768
    var v4 uint64 = 136
    var v5 uint64 = v4
    var v6 float32 = (-2.25 + -2.25)
    var v16 int16 = v3
    for fi17 := 0; fi17 < 2; fi17++ {
        v2 = (v2 & (v2 - v2))
        _ = fi17
    }
    v4 += v1
    v6 += v6
    v18 := GenS1{F0: -77, F1: 1}
    fmt.Println(v4)
    v19 := GenS1{F0: genFn0(18446744073709551615, v1), F1: -418}
    fmt.Println(v1, v1)
    if true {
        fmt.Println(v5, v6, v16)
    }
    v2 = v2
    v3 = v3
    v20 := []int{42}
    var v21 bool = false
    v20 = append(v20, -588)
    v20 = append(v20, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
