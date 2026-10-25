package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 255
    var v3 float32 = 0.1
    v4 := make(map[uint32]uint32)
    v5 := []uint32{}
    var v6 int64 = 1
    v2 = (v2 | v2)
    for fi7 := 0; fi7 < 7; fi7++ {
        v2 = v2
        _ = fi7
    }
    fmt.Println(v6)
    fmt.Println(v2, v2, v2)
    switch 2147483647 {
    case -286:
        v3 += v3
    case -142:
        fmt.Println(v2)
    }
    var v8 float64 = (1e10 + 3.14159)
    fmt.Println(v3)
    v2 += 2
    _ = len(v5)
    v3 += (float32(1e10) * (v3 * float32(1e10)))
    fmt.Println(v8, v6)
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: 0, F1: -123, F2: v2}
        _ = fi9
        _ = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
