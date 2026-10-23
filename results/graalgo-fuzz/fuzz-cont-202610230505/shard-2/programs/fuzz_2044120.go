package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int8 }

func main() {
    v1 := []uint32{65535, 4294967295}
    var v2 uint64 = 4294967295
    var v3 int32 = 2147483647
    var v4 float64 = (0.0 / 1.0)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v5 bool = true
    switch 0 {
    case -1:
        v6 := GenS1{F0: 128}
        _ = v6
    case 369:
        var v7 uint8 = 71
        _ = v7
    case -873:
        var v8 int = -9223372036854775808
        _ = v8
    default:
        var v9 float32 = 1.5
        _ = v9
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v2 -= (^(v2))
    var v10 int32 = 1
    var v11 uint16 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "" {
    case "\n\t":
        v12 := GenS2{F0: v5, F1: 127}
        _ = v12
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint = 1
            _ = fi13
            _ = v14
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
