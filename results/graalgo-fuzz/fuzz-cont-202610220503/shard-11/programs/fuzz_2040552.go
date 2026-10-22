package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 bool }
type GenS2 struct { F0 int; F1 int; F2 uint8 }

func main() {
    var v1 float32 = 3.14159
    var v2 int8 = 127
    v3 := make(map[uint32]int)
    var v4 uint8 = 0
    var v5 int32 = -151
    var v6 int64 = -446
    v5 = v5
    v1 += 0.0
    v7 := GenS1{F0: false, F1: 1, F2: false}
    _ = len(v3)
    var v8 uint64 = 1
    var v9 uint64 = uint64(v4)
    v8 += 493
    _ = len(v3)
    var v10 int = 128
    fmt.Println(v5)
    v11 := GenS2{F0: (v10 - v10), F1: (v10 * (-(v10))), F2: v4}
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 float64 = (1e10 - (0.1 - 1.5))
        var v14 float64 = (3.14159 / v13)
        var v15 float64 = (v14 * v14)
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := []int{(v10 >> ((255) & 63)), -748, (v10 / ((2147483647) | 1)), int(v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
