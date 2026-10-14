package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint) int64 {
    p0 = 4294967295
    var v7 uint16 = 42
    v7 = (v7 ^ uint16(p0))
    _ = p0
    _ = v7
    return 201
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 0
    var v3 int = -9223372036854775808
    var v4 float64 = (3.14159 * ((float64(float32(3.14159)) * 1e10) / 1.0))
    var v5 uint64 = 994
    var v6 int16 = -612
    v2 = 783
    switch v3 {
    case 127:
        _ = len(v1)
    }
    switch 1 {
    case 182:
        var v8 uint = 1
        _ = v8
    case 128:
        var v9 int32 = int32(v6)
        _ = v9
    case -1:
        v2 += (v2 << ((160) & 63))
    default:
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := GenS1{F0: (v5 % ((v5) | 1)), F1: 999}
            _ = fi10
            _ = v11
        }
    }
    if true {
        var v12 int = v3
        var v13 uint64 = (v5 ^ 2)
        _ = v12
        _ = v13
    }
    if ((!true) || (42 <= 1)) {
        for fi14 := 0; fi14 < 12; fi14++ {
            v5 -= v5
            v1[2147483648] = "a"
            v15 := GenS1{F0: v5, F1: -860}
            _ = fi14
            _ = v15
        }
    } else {
        v4 = float64((-2.25 * (0.0 + (((0.1 - 0.1) - ((float32(((3.14159 - (v4 + v4)) - (0.0 - 0.1))) * 1e10) / 0.1)) / 0.1))))
        fmt.Println(v4, v5, v4)
        var v16 int16 = int16(v3)
        _ = v16
    }
    v17 := GenS1{F0: (^(v5)), F1: genFn0(637)}
    var v18 float32 = float32((0.0 + (v4 * 1.5)))
    _ = len(v1)
    var v19 uint64 = (v5 | v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
