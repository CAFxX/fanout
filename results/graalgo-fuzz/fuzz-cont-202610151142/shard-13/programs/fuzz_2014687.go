package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int16; F2 int32 }

func main() {
    var v1 int32 = 0
    var v2 int32 = (v1 + v1)
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v5 int8 = -1
    for fi6 := 0; fi6 < 13; fi6++ {
        v4[-2147483648] = 98
        _ = fi6
    }
    v1 = (v2 << ((18446744073709551615) & 63))
    if (2 >= 327) {
        if (!true) {
            var v7 uint16 = 65535
            _ = v7
        } else {
            var v8 int32 = (v1 - v1)
            v9 := GenS1{F0: 127}
            delete(v3, 32767)
            _ = v8
            _ = v9
        }
    } else {
        var v10 uint8 = 2
        _ = v10
    }
    v5 += (-(v5))
    fmt.Println(v1, v2, v5)
    var v11 uint8 = 1
    if (9223372036854775807 >= 1) {
        v12 := GenS1{F0: (v11 + 0)}
        var v13 int8 = (v5 & 100)
        var v14 float64 = (1.5 / (0.1 / 0.1))
        _ = v12
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            fmt.Println(v2, v1)
            var v16 uint = 255
            var v17 int16 = 32767
            _ = fi15
            _ = v16
            _ = v17
        }
        v1 -= -379
        v5 += (^(v5))
    }
    v18 := GenS1{F0: v11}
    v2 += (-(v1))
    fmt.Println(v5, v1)
    v19 := []int{-2147483648, 32767}
    v11 = v11
    v2 -= 0
    _ = len(v19)
    v3[258] = 1
    v4[-905] = 65535
    switch -9223372036854775808 {
    case -1:
        for fi20 := 0; fi20 < 10; fi20++ {
            v11 = (v11 ^ (v11 ^ (v11 / ((v11) | 1))))
            v21 := v4[fi20]
            v22 := GenS2{F0: 65535, F1: -887, F2: (v2 & v1)}
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
