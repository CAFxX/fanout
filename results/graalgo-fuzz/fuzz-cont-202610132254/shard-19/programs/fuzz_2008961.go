package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int32; F1 int }
func genFn0(p0 uint8, p1 uint16, p2 string) int16 {
    p2 = p2
    switch -1 {
    case -353:
        if false {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 int64 = -9223372036854775808
                for fi8 := 0; fi8 < 12; fi8++ {
                    var v9 int = 0
                    v10 := GenS1{F0: p0, F1: (p1 & (p1 ^ p1))}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
            var v11 float64 = 3.14159
            _ = v11
        } else {
            var v12 uint8 = (p0 << ((255) & 63))
            _ = v12
        }
    default:
        p0 = 183
    }
    if (316 <= 128) {
        var v13 uint32 = 42
        _ = v13
    }
    v14 := GenS1{F0: 255, F1: (p1 << ((239) & 63))}
    v15 := GenS1{F0: (p0 / ((255) | 1)), F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 32767
}
func genFn1(p0 int64, p1 int8) uint32 {
    if (1 >= -810) {
        if (!true) {
            var v16 float32 = float32(float64((1e10 + 1e10)))
            fmt.Println(p1, v16)
            _ = v16
        } else {
            p0 = -1
        }
        v17 := GenS1{F0: 30, F1: 65535}
        v18 := GenS2{F0: 385, F1: 42}
        _ = v17
        _ = v18
    } else {
        p1 -= (p1 & p1)
        fmt.Println(p1, p1)
    }
    var v19 int64 = 9223372036854775807
    if (!(!true)) {
        fmt.Println(v19, p0)
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = v19
    return 2147483648
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{127, 1}
    v3 := make(map[uint32]bool)
    var v4 uint8 = 255
    var v5 int8 = 0
    v5 = v5
    v4 -= v4
    for fi20 := 0; fi20 < 2; fi20++ {
        v4 += (v4 << ((1) & 63))
        var v21 int16 = -981
        _ = fi20
        _ = v21
    }
    var v22 uint8 = (v4 ^ v4)
    v5 += (v5 * (^(v5)))
    v5 += v5
    var v23 uint8 = (^(v22))
    var v24 int32 = -796
    if len(v2) > 0 {
        v2[0] = 445
    }
    var v25 float64 = ((0.0 * (1.5 - 1.5)) + -2.25)
    _ = len(v1)
    fmt.Println(v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
