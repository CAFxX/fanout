package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]int)
    var v3 int8 = -72
    var v4 uint = 114
    var v5 uint = (v4 << ((v4) & 63))
    v6 := v2[4294967295]
    var v7 float64 = 3.14159
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float64 = 3.14159
        v10 := GenS1{F0: -2147483648, F1: 4294967295}
        v2[fi8] = 203
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 = ((v7 + (float64(0.1) + 0.1)) + 3.14159)
    v11 := v2[-1]
    v11 = v11
    var v12 float64 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
