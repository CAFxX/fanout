package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
type GenS2 struct { F0 int }
func genFn0(p0 int64, p1 int8, p2 uint32) uint16 {
    var v7 int8 = (p1 + (p1 - p1))
    p1 += (p1 - p1)
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: 2, F1: 22, F2: 42}
        var v10 int8 = 42
        v11 := GenS1{F0: p2, F1: -1, F2: v10}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    p0 = -9223372036854775808
    p1 -= 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1() uint64 {
    var v12 int8 = -1
    fmt.Println(v12)
    _ = v12
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 bool = (((0.0 / ((((3.14159 - (((0.1 / 1.5) / 0.1) + 3.14159)) + (1.5 + 1e10)) - -2.25) / -2.25)) / (0.1 + ((1e10 - 1.5) + (0.1 - 1.5)))) == (1.5 + -2.25))
    var v5 uint64 = 4294967295
    var v6 uint64 = (v5 ^ (v5 << ((1) & 63)))
    var v13 uint8 = 255
    _ = len(v3)
    switch 1 {
    case 127:
        switch -2147483648 {
        case -2147483648:
            v14 := GenS2{F0: 1}
            _ = v14
        default:
            fmt.Println(v5)
        }
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 float64 = ((1e10 - 1e10) - float64(0.0))
        var v17 string = ""
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 bool = (float32((1.5 * (((1.5 * (3.14159 + -2.25)) * float64(0.1)) / (0.1 - 0.0)))) > 0.0)
    v19 := GenS2{F0: -9223372036854775808}
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 uint64 = (v5 + 18446744073709551615)
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v5 = v6
        _ = fi22
    }
    v6 += (v5 ^ (^(v5)))
    v23 := []int{-9223372036854775808, -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
