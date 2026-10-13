package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint32, p1 bool, p2 uint) uint {
    var v7 int64 = -9223372036854775808
    var v8 uint32 = 340
    p2 -= 61
    var v9 float64 = (0.1 + ((0.1 * 3.14159) * 0.1))
    switch 128 {
    case -2147483648:
        switch 9223372036854775807 {
        case 1:
            v8 += v8
        default:
            var v10 int32 = 2147483647
            _ = v10
        }
    case -9223372036854775808:
        var v11 int16 = 1000
        _ = v11
    case 4294967295:
        fmt.Println(p0, p1)
    default:
        for fi12 := 0; fi12 < 17; fi12++ {
            v9 = (float64(3.14159) + v9)
            _ = fi12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int8 = 100
    var v2 int32 = 0
    var v3 float64 = 0.1
    var v4 int32 = v2
    var v5 string = "\n\t"
    v6 := make(map[int]uint32)
    if (("" + v5) < v5) {
        var v13 string = v5
        _ = v13
    } else {
        v3 -= (1.5 / (1e10 + (0.0 * v3)))
        fmt.Println(v1, v4)
    }
    v3 -= 0.1
    fmt.Println(v4, v1)
    switch 9223372036854775807 {
    case 4294967295:
        fmt.Println(v5)
    case 0:
        v3 = v3
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v3 -= ((0.1 / 0.1) + v3)
        v15 := GenS1{F0: int16(v2)}
        switch (-(fi14)) {
        case -419:
            var v16 int32 = (-(v4))
            _ = v16
        case -380:
            for fi17 := 0; fi17 < 1; fi17++ {
                delete(v6, (fi17 - (fi14 * fi14)))
                _ = fi17
            }
        default:
            var v18 int = (fi14 / (((fi14 ^ 732)) | 1))
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        v1 = v1
        var v20 uint64 = 268
        fmt.Println(v3)
        _ = fi19
        _ = v20
    }
    v21 := v6[657]
    var v22 bool = true
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 int = (fi24 % ((fi24) | 1))
            v2 += (v2 * 2147483647)
            var v26 float32 = (0.1 + float32(v3))
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 int64 = 0
        var v28 int16 = -32768
        _ = fi23
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
