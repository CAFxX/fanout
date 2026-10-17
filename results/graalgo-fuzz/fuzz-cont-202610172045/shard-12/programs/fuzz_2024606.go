package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    v5 := make(map[string]uint32)
    v6 := []int{}
    var v7 uint64 = 18446744073709551615
    fmt.Println(v1)
    if false {
        v6 = append(v6, 207)
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint16 = 42
        v7 += (v7 >> ((280) & 63))
        v10 := GenS1{F0: -871, F1: (fi8 & fi8), F2: 850}
        _ = fi8
        _ = v10
        _ = v9
    }
    v6 = append(v6, -34)
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    v6 = append(v6, -9223372036854775808)
    fmt.Println(v1, v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
