package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    var v1 int32 = 1
    var v2 float64 = float64((3.14159 - 1e10))
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    v2 += -2.25
    v2 = (v2 + v2)
    var v5 float64 = (((v2 + (v2 * v2)) / v2) + (((v2 - (v2 + (1.5 / 3.14159))) / (v2 - (0.0 + v2))) - 3.14159))
    var v6 uint16 = 42
    var v7 int8 = -1
    v4[""] = -2147483648
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint = 2
        var v10 uint32 = 4294967295
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: -943, F1: 1}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
