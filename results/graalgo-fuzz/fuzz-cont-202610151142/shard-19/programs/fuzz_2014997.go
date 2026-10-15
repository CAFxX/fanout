package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint32) int64 {
    p0 += (p0 >> ((1) & 63))
    var v6 int64 = 399
    fmt.Println(v6)
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 uint16, p1 uint32) int {
    p1 -= p1
    var v7 float64 = float64((float32((3.14159 / ((-2.25 * 0.1) * ((-2.25 / (float64(((((1.5 - ((-2.25 - 1.5) + float32(-2.25))) + 0.1) - (0.0 / (float32(1.5) + 0.0))) / -2.25)) * 3.14159)) - 1e10)))) + 0.0))
    _ = p0
    _ = p1
    _ = v7
    return -562
}

func main() {
    var v1 int64 = 42
    var v2 float64 = -2.25
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 32767
    var v5 uint16 = v4
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int8 = 100
        _ = fi8
        _ = v9
    }
    var v10 uint64 = 42
    fmt.Println(v5, v5)
    var v11 int16 = 32767
    var v12 int8 = 9
    var v13 int32 = 0
    v3[1] = 318
    _ = len(v3)
    v11 += (v11 + (v11 & (v11 * v11)))
    for fi14 := 0; fi14 < 6; fi14++ {
        v13 -= (v13 & (^(v13)))
        _ = fi14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 bool = (-2.25 >= float32(1e10))
        if (v16 && ((v16 && false) || (!v16))) {
            var v17 float32 = (1e10 * 0.0)
            v18 := GenS1{F0: (v1 >> ((1) & 63))}
            _ = v17
            _ = v18
        }
        if false {
            if v16 {
                v11 = (v11 & v11)
                var v19 int = genFn1((^(v4)), 4294967295)
                fmt.Println(v2)
                _ = v19
            }
            var v20 int = (fi15 ^ fi15)
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    v2 -= v2
    v21 := v3[281]
    var v22 int16 = v11
    var v23 float32 = 1.5
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v24 int32 = int32(v21)
        _ = v24
    default:
        if ((v12 - v12) >= -1) {
            v25 := []int{}
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
