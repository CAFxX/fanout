package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint }
func genFn0(p0 uint16, p1 int) uint32 {
    if true {
        var v5 uint32 = 221
        p0 += (p0 * (^(p0)))
        p0 = 42
        _ = v5
    }
    var v6 uint8 = 42
    switch (("\n\t" + "hello") + "a") {
    case "foo bar":
        var v7 int64 = -394
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 4; fi8++ {
            v6 -= (v6 >> ((1) & 63))
            var v9 float32 = (-2.25 * 0.1)
            var v10 float64 = (0.1 + 1e10)
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        p0 = (p0 * (p0 ^ 169))
    }
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    var v1 float32 = -2.25
    var v2 float64 = 0.1
    v3 := []uint32{2147483648, 2147483647}
    var v4 uint = 18446744073709551615
    if false {
        var v11 bool = true
        v4 += (v4 + v4)
        v1 = (v1 - v1)
        _ = v11
    } else {
        v12 := GenS1{F0: 703, F1: v4, F2: 18446744073709551615}
        _ = v12
    }
    v2 = float64((v1 - 3.14159))
    var v13 uint32 = 2147483648
    for fi14 := 0; fi14 < 6; fi14++ {
        switch "\n\t" {
        case "foo bar":
            v3 = append(v3, v13)
        case "hello":
            v1 = v1
        default:
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := GenS2{F0: 288, F1: -118, F2: (^(v4))}
                v13 += v13
                var v17 uint = (v4 / ((591) | 1))
                var v18 uint8 = 48
                v1 = (0.1 + v1)
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        _ = fi14
    }
    if false {
        v3 = append(v3, v13)
        fmt.Println(v4)
        for fi19 := 0; fi19 < 3; fi19++ {
            for fi20 := 0; fi20 < 10; fi20++ {
                v2 = ((v2 / 1.5) * v2)
                var v21 int64 = 0
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 bool = (!true)
                _ = fi22
                _ = v23
            }
            _ = fi19
        }
    } else {
        fmt.Println(v2, v1, v1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 uint = 255
        fmt.Println(v25)
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 int32 = 0
        var v28 uint64 = 411
        fmt.Println(v27)
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int8 = 100
    v2 -= ((float64(-2.25) + ((v2 + (3.14159 * v2)) * (float64(v1) + v2))) * float64(v1))
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
