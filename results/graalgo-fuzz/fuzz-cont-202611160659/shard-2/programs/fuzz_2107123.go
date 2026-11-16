package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int8; F1 int32; F2 uint8 }

func main() {
    var v1 uint16 = 570
    var v2 bool = (!(50 > 185))
    var v3 uint64 = 279
    var v4 float64 = 0.0
    v3 += v3
    fmt.Println(v3, v4)
    v5 := GenS1{F0: v1, F1: 922}
    if v2 {
        v4 = 0.1
    }
    switch -448 {
    case 32768:
        v3 = v3
    case 958:
        v6 := GenS2{F0: -80, F1: 42, F2: 1}
        _ = v6
    }
    v7 := []int{170, 173, -9223372036854775808, 443}
    v8 := GenS2{F0: 1, F1: 915, F2: 95}
    fmt.Println(v4, v2)
    v3 += v3
    if v2 {
        v9 := GenS1{F0: v1, F1: (v1 >> ((2) & 63))}
        var v10 int = -292
        _ = v10
        _ = v9
    }
    v2 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
