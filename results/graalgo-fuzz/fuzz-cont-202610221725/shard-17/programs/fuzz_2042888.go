package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int16 }

func main() {
    var v1 int8 = 100
    var v2 float32 = 1.5
    v3 := make(map[string]uint32)
    var v4 uint32 = 2147483647
    var v5 int = 4294967295
    v5 = v5
    if false {
        var v6 uint16 = 0
        for fi7 := 0; fi7 < 5; fi7++ {
            v3[("a" + "")] = (^(v4))
            _ = fi7
        }
        var v8 uint8 = 0
        _ = v6
        _ = v8
    }
    v9 := GenS1{F0: false, F1: -105, F2: 1}
    v10 := GenS1{F0: (2147483647 <= 270), F1: 127, F2: 1000}
    v2 += float32(((1.5 * (-2.25 * (0.1 + (((-2.25 * float64(v2)) * (1.5 / 1e10)) + 1e10)))) + 0.0))
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 uint64 = 2
        var v13 int32 = 1
        _ = fi11
        _ = v12
        _ = v13
    }
    switch v5 {
    case 4294967295:
        v14 := []int{(v5 * v5), (v5 % (((v5 << ((65535) & 63))) | 1)), (-(v5))}
        _ = v14
    case -880:
        v5 = v5
    default:
        v4 = v4
    }
    v5 = v5
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 float32 = (v2 - (v2 / ((v2 + 0.0) - (v2 * v2))))
        _ = fi15
        _ = v16
    }
    var v17 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
