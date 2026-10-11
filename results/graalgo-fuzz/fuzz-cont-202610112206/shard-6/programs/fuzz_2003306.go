package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]bool)
    var v3 float32 = 1.5
    v4 := make(map[uint32]int)
    var v5 int32 = -1
    for fi6 := 0; fi6 < 11; fi6++ {
        v3 -= (v3 * (v3 / (1.5 + v3)))
        _ = fi6
    }
    var v7 int = 2147483647
    v8 := GenS1{F0: 1, F1: 42}
    v9 := []int{(v7 - v7), (^(v7))}
    var v10 uint16 = 0
    fmt.Println(v5)
    v5 += v5
    var v11 uint32 = 4294967295
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 float32 = -2.25
        _ = fi12
        _ = v13
    }
    v5 -= -1
    v1 += (v1 >> ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
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
