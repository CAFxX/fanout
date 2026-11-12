package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }

func main() {
    var v1 uint = 2
    var v2 int32 = 161
    v3 := make(map[int]int)
    var v4 int16 = 0
    v5 := GenS1{F0: 1, F1: -1}
    var v6 uint64 = 42
    v7 := []int{9223372036854775807, 1, 170, 91}
    v6 += (v6 + 2)
    fmt.Println(v6, v2, v2)
    var v8 uint = (v1 + (v1 * v1))
    fmt.Println(v8, v1)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
