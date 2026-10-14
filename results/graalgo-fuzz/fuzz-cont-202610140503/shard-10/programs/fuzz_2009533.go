package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
func genFn0(p0 int8, p1 uint8) int8 {
    v6 := []int{-2147483648, 0, 9223372036854775807}
    var v7 uint8 = (^(p1))
    var v8 float64 = (0.0 / 1.0)
    v9 := GenS1{F0: 838, F1: 2, F2: 0}
    if (18446744073709551615 >= uint(p1)) {
        var v10 uint8 = (p1 + v7)
        var v11 int = 9223372036854775807
        v12 := GenS1{F0: 912, F1: 0, F2: -1}
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (p0 % (((p0 % (((p0 ^ p0)) | 1))) | 1))
}

func main() {
    var v1 float64 = -2.25
    v2 := []int{153, -9223372036854775808, 2147483647, -834}
    var v3 float64 = (v1 * -2.25)
    var v4 int32 = 1
    var v5 uint64 = 705
    var v13 float64 = 1e10
    v14 := GenS1{F0: 2, F1: 2147483647, F2: -32768}
    if len(v2) > 0 {
        v2[0] = 372
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint16 = 1
        _ = fi15
        _ = v16
    }
    v4 = (v4 - (v4 % ((v4) | 1)))
    v2 = append(v2, 42)
    v3 -= v13
    for fi17 := 0; fi17 < 6; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            v13 = ((v1 - (0.0 + (1e10 - v3))) - v3)
            v3 = (((0.0 * 1.5) + (-2.25 + float64((3.14159 * 3.14159)))) - v13)
            _ = len(v2)
            _ = fi18
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi17
    }
    v4 = 81
    v19 := GenS1{F0: v5, F1: 2, F2: -1}
    var v20 float64 = ((v3 / v3) - v13)
    var v21 uint8 = 1
    var v22 int8 = -60
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := GenS1{F0: v5, F1: 580, F2: 1}
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
