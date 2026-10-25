package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    v1 := []uint32{42}
    var v2 float64 = -2.25
    var v3 uint = 4294967295
    v4 := make(map[int]uint32)
    var v5 uint = v3
    var v6 int8 = 42
    v7 := GenS1{F0: "a", F1: true}
    var v8 uint = 0
    if (((0.1 * (0.0 - ((1.5 * float64((1e10 / 1.0))) + v2))) - v2) > (v2 - ((((1.5 * float64(0.0)) * float64(float32((v2 / v2)))) + v2) * v2))) {
        var v9 uint = (^(v3))
        _ = v9
    }
    v8 += v8
    v3 = 4294967295
    v10 := GenS1{F0: "\n\t", F1: false}
    v2 = v2
    switch 32768 {
    case -2147483648:
        if len(v1) > 0 {
            v1[0] = 1
        }
    case 719:
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := v4[(fi11 * (fi11 * fi11))]
            _ = fi11
            _ = v12
        }
    default:
        v13 := GenS2{F0: -9223372036854775808, F1: 411}
        _ = v13
    }
    if (-1 >= 1) {
        for fi14 := 0; fi14 < 8; fi14++ {
            v8 -= 1
            _ = fi14
        }
    }
    var v15 int32 = 2147483647
    if (!((v15 + (v15 >> ((73) & 63))) == (v15 * v15))) {
        if (v3 < 4294967295) {
            var v16 int64 = 9223372036854775807
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 int64 = (-(v16))
                fmt.Println(v6, v2, v15)
                v6 = v6
                _ = fi17
                _ = v18
            }
            _ = v16
        }
    } else {
        v1 = append(v1, 4294967295)
    }
    _ = len(v4)
    v15 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
