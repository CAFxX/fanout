package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int64; F1 uint64; F2 bool }

func main() {
    v1 := []string{}
    var v2 float32 = float32((0.1 - 0.0))
    var v3 uint = 235
    var v4 float64 = ((0.0 + 0.1) + 1e10)
    v5 := make(map[int]string)
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v3)
        var v8 uint64 = 684
        v9 := GenS1{F0: 128, F1: v3}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 42
    var v11 uint16 = 32767
    fmt.Println(v4, v2)
    v10 = 0
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v2 = v2
    v12 := GenS2{F0: -9223372036854775808, F1: (v10 % ((v10) | 1)), F2: true}
    fmt.Println(v3, v10, v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
