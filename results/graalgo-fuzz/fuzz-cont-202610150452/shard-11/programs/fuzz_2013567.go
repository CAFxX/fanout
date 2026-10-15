package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint8; F2 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := []uint32{2147483647, 2147483647, 2147483647}
    var v4 int8 = 100
    var v5 float32 = 1.5
    v5 = 0.1
    v5 -= v5
    v5 = v5
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(v4, v4, v5)
        fmt.Println(v4)
        _ = fi6
    }
    var v7 float64 = (1.5 + ((float64(v5) / -2.25) + 0.0))
    v8 := GenS1{F0: 255}
    v7 -= 0.0
    v7 = 0.0
    v9 := GenS2{F0: 0, F1: 174, F2: 0}
    var v10 int = 32768
    v11 := GenS2{F0: 669, F1: 1, F2: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
