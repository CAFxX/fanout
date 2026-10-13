package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]int)
    var v3 uint64 = 18446744073709551615
    var v4 int64 = 0
    v4 += -1
    v5 := []int{2147483647, 1, -14, 4294967295}
    v3 += (v3 / ((v3) | 1))
    var v6 int32 = 1
    var v7 uint64 = v3
    var v8 bool = true
    v9 := GenS1{F0: 638, F1: -620}
    var v10 uint32 = (v1 ^ v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
