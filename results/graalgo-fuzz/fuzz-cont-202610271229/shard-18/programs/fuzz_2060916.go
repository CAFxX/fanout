package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int }
func genFn0(p0 int32, p1 uint64, p2 uint16) uint16 {
    fmt.Println(p2)
    fmt.Println(p0)
    var v5 int8 = -128
    if (255 != 65535) {
        var v6 int32 = (p0 << ((0) & 63))
        var v7 float32 = (-2.25 - 3.14159)
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p2 & p2)
}

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{}
    var v3 bool = (!false)
    var v4 int64 = 9223372036854775807
    v8 := GenS1{F0: int16(v1), F1: 565, F2: -265}
    var v9 int32 = v1
    var v10 uint8 = 177
    v1 = (-(v9))
    var v11 int32 = 1
    v9 = -1
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint16 = 255
        v3 = v3
        _ = fi12
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 20; fi14++ {
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 float64 = 1.5
                var v17 int16 = 0
                var v18 uint8 = (v10 * v10)
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi14
        }
    }
    switch 32767 {
    case -477:
        v3 = v3
    case 42:
        v19 := GenS1{F0: 42, F1: 42, F2: 32767}
        _ = v19
    case -9223372036854775808:
        var v20 float32 = 1e10
        _ = v20
    }
    if v3 {
        fmt.Println(v11, v3, v4)
        var v21 uint16 = 2
        _ = v21
    } else {
        var v22 int64 = (-(v4))
        v23 := GenS1{F0: -775, F1: 96, F2: -691}
        v9 -= (v1 * 2147483647)
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 int16 = -471
        var v26 int8 = 42
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
