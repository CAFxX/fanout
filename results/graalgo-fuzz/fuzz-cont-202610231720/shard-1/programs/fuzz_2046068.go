package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 float64 = (0.0 / 0.1)
    var v4 int32 = 1
    var v5 int64 = -692
    var v6 uint8 = 42
    var v7 int64 = -993
    v4 += 1
    v7 = 42
    fmt.Println(v7)
    var v8 int64 = v5
    v9 := []int{}
    fmt.Println(v3)
    v10 := GenS1{F0: 42, F1: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
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
