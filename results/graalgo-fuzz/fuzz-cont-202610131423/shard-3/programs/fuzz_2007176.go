package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int32; F2 uint64 }
func genFn0(p0 bool) int64 {
    fmt.Println(p0)
    var v6 string = ""
    _ = p0
    _ = v6
    return 0
}
func genFn1() uint8 {
    v7 := []int{0, 0, 376, 42}
    switch -1 {
    case -556:
        v8 := GenS1{F0: 128}
        _ = v8
    default:
        var v9 bool = (-793 < 1)
        _ = v9
    }
    _ = v7
    return 127
}

func main() {
    var v1 int64 = -1
    var v2 int64 = 0
    v3 := make(map[uint32]uint32)
    var v4 int32 = 2147483647
    v5 := make(map[uint32]bool)
    v4 += -196
    if (502 <= int16(v4)) {
        v10 := GenS2{F0: 18446744073709551615, F1: (v4 << ((255) & 63)), F2: 972}
        _ = v10
    } else {
        v1 -= v2
        var v11 uint8 = 1
        _ = v11
    }
    var v12 bool = (32767 <= 32767)
    if ((!(-1 >= v4)) || (1.5 == -2.25)) {
        fmt.Println(v12)
    } else {
        v4 += v4
        var v13 float32 = ((3.14159 + float32(((((0.1 + 1e10) * 0.0) - (-2.25 * float64(-2.25))) * 0.0))) * 0.1)
        switch 128 {
        case -1:
            for fi14 := 0; fi14 < 11; fi14++ {
                v4 = (v4 / (((^(v4))) | 1))
                _ = fi14
            }
        default:
            var v15 uint64 = 18446744073709551615
            _ = v15
        }
        _ = v13
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        if (v2 <= (v2 | genFn0(true))) {
            v17 := GenS2{F0: 516, F1: -1, F2: 181}
            var v18 uint8 = 127
            var v19 uint8 = (v18 >> ((uint(v18)) & 63))
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    var v20 uint = 42
    var v21 uint8 = 64
    var v22 int32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
