package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int16 }

func main() {
    v1 := []string{}
    var v2 float64 = (-2.25 + (1e10 / 0.1))
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 4294967295
    var v5 float32 = 3.14159
    fmt.Println(v5, v4, v2)
    if false {
        fmt.Println(v4)
    }
    fmt.Println(v2, v5, v2)
    v2 = 3.14159
    delete(v3, 0)
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int32 = -786
        v7 = (v7 % ((-1) | 1))
        v4 = 532
        _ = fi6
        _ = v7
    }
    switch -9223372036854775808 {
    case 1:
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int16 = 42
            _ = fi8
            _ = v9
        }
    case -100:
        var v10 int16 = 0
        _ = v10
    case 312:
        v11 := GenS1{F0: v4, F1: 42, F2: 0}
        _ = v11
    }
    var v12 bool = true
    switch 32768 {
    case 4294967295:
        v5 -= 0.0
    case 128:
        v2 = (-2.25 + v2)
    case 2147483647:
        v12 = v12
    default:
        fmt.Println(v4, v12, v2)
    }
    var v13 int32 = 2147483647
    var v14 int = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
