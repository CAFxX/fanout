package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 1.5
    v3 := make(map[int]bool)
    v2 += 0.0
    _ = len(v1)
    v2 = v2
    if true {
        v4 := v3[-159]
        v5 := []int{9223372036854775807}
        _ = v4
        _ = v5
    }
    fmt.Println(v2)
    fmt.Println(v2)
    v6 := GenS1{F0: -797}
    if (42 < 2147483647) {
        fmt.Println(v2)
        v7 := GenS1{F0: 756}
        _ = v7
    }
    v2 = ((0.1 + (v2 / ((1e10 - ((0.1 * v2) / 1e10)) / 3.14159))) - float64(0.1))
    delete(v1, 1)
    v8 := GenS1{F0: -395}
    var v9 uint8 = 127
    v10 := []int{-1, -793, 42}
    fmt.Println(v2, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v8
    _ = v9
}
