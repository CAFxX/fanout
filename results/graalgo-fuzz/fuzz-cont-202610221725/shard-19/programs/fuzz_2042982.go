package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }

func main() {
    var v1 int16 = 0
    var v2 uint16 = 1
    v3 := make(map[uint32]uint32)
    v4 := []uint32{204, 2, 549, 656}
    var v5 int = 127
    switch (v5 | (-(v5))) {
    case 127:
        fmt.Println(v1)
    case 4294967295:
        v1 += (-(v1))
    case 863:
        v5 = int(v1)
    default:
        if (!false) {
            v6 := GenS1{F0: 2, F1: 128}
            if (504 >= 0) {
                fmt.Println(v5, v2, v1)
            } else {
                var v7 uint32 = 206
                _ = v7
            }
            _ = v6
        }
    }
    v8 := GenS2{F0: 59, F1: -1, F2: (v2 + (v2 ^ v2))}
    for fi9 := 0; fi9 < 15; fi9++ {
        v2 += v2
        v4 = append(v4, 0)
        _ = fi9
    }
    fmt.Println(v1, v5)
    v1 -= (v1 ^ 32767)
    var v10 float64 = ((-2.25 * ((1e10 - 1.5) * 1.5)) / 0.1)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v1, v2)
    case "":
        var v11 int32 = 2147483647
        _ = v11
    case "hello":
        v12 := GenS2{F0: -128, F1: 100, F2: (v2 * v2)}
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
