package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8; F2 int16 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]uint32)
    var v3 float64 = float64((0.0 * (-2.25 * (float32((float64(1.5) + 3.14159)) - (0.0 + 3.14159)))))
    var v4 int16 = 913
    v5 := []uint32{2147483647, 65535, 65535}
    v1 = v4
    delete(v2, 4294967295)
    var v6 float64 = 0.1
    if (!(v3 < v6)) {
        fmt.Println(v6, v4, v3)
        v4 = v1
        if (!false) {
            v3 += (v3 / v6)
        } else {
            v7 := GenS1{F0: true}
            _ = v7
        }
    }
    var v8 uint8 = 128
    var v9 float32 = 0.1
    v9 = v9
    v10 := GenS2{F0: 4294967295, F1: (v8 | 166), F2: (v1 & v4)}
    v4 += v4
    v4 = (-(v1))
    switch 303 {
    case -9223372036854775808:
        switch 1 {
        case -9223372036854775808:
            var v11 int32 = 0
            _ = v11
        case 4294967295:
            v4 -= v1
        }
    default:
        v6 -= (0.0 / -2.25)
    }
    v9 = (v9 + v9)
    var v12 int = 394
    delete(v2, 2147483648)
    v4 += v4
    v5 = append(v5, uint32(v8))
    var v13 float64 = 0.0
    for fi14 := 0; fi14 < 8; fi14++ {
        v5 = append(v5, 1)
        fmt.Println(v8)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
