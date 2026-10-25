package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }

func main() {
    var v1 uint16 = 0
    var v2 uint16 = v1
    v3 := make(map[uint32]int)
    if (42 <= 18446744073709551615) {
        var v4 uint16 = v2
        var v5 int64 = -554
        var v6 float64 = -2.25
        _ = v4
        _ = v5
        _ = v6
    }
    var v7 int16 = 321
    var v8 uint32 = 2
    var v9 bool = true
    v10 := GenS1{F0: 42, F1: -9223372036854775808}
    if true {
        var v11 int64 = -1
        _ = v11
    }
    v9 = v9
    switch 32767 {
    case -2147483648:
        var v12 uint8 = 42
        _ = v12
    default:
        if (1 <= 0) {
            v13 := []int{629}
            var v14 uint64 = 2
            _ = v13
            _ = v14
        }
    }
    var v15 float64 = 0.1
    var v16 uint8 = 127
    v2 += (v2 | (^(v1)))
    v9 = v9
    v9 = (((1e10 + 1.5) + (1.5 + (0.0 * float32(0.1)))) > 1e10)
    fmt.Println(v8, v8)
    var v17 int8 = 0
    v18 := v3[v8]
    v19 := []int{(v18 >> ((255) & 63)), (^(v18))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
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
    _ = v7
    _ = v8
    _ = v9
}
