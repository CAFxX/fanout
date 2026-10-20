package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }

func main() {
    var v1 uint32 = 4294967295
    v2 := []uint32{(v1 / ((v1) | 1)), 1, (v1 | (v1 * (^(v1)))), v1}
    v3 := make(map[string]string)
    var v4 float64 = 0.0
    var v5 float64 = v4
    v6 := make(map[int]uint32)
    var v7 int8 = -128
    v5 += v4
    fmt.Println(v7, v1)
    var v8 uint8 = 127
    var v9 uint32 = (v1 ^ v1)
    var v10 uint32 = v9
    if (!(true && false)) {
        fmt.Println(v4)
        if (!((!(65535 <= 457)) && (true || (false || false)))) {
            v11 := GenS1{F0: 42, F1: 0, F2: 1}
            v10 -= 2
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 uint8 = v8
        var v14 int64 = 1
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := v6[-359]
    var v16 int16 = 1000
    var v17 uint64 = 0
    v7 += (-(v7))
    v18 := GenS1{F0: v7, F1: 417, F2: 65535}
    var v19 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
