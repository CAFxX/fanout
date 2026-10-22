package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int16 = 20
    v2 := []int{-9223372036854775808}
    var v3 int64 = -1
    v4 := []uint32{0}
    var v5 uint = 65535
    v5 += v5
    v3 += v3
    v6 := GenS1{F0: 127}
    var v7 float32 = float32(0.1)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int8 = -128
        _ = fi8
        _ = v9
    }
    var v10 uint16 = 936
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v4)
    if (!false) {
        var v11 int64 = (v3 + v3)
        _ = v11
    } else {
        switch 473 {
        case 962:
            fmt.Println(v5, v3, v1)
        case -2147483648:
            v5 += 331
        default:
            v1 -= (v1 + (v1 * (v1 - (-(v1)))))
        }
        var v12 int16 = (v1 % (((v1 / ((646) | 1))) | 1))
        _ = v12
    }
    if (4294967295 == 2) {
        var v13 int64 = v3
        for fi14 := 0; fi14 < 11; fi14++ {
            fmt.Println(v1, v10)
            _ = fi14
        }
        _ = len(v2)
        _ = v13
    }
    v7 = float32(float64(1e10))
    var v15 uint32 = 4294967295
    var v16 uint32 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
