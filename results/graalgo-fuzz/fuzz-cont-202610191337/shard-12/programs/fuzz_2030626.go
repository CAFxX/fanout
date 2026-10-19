package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint32 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]bool)
    var v3 float32 = 1.5
    v4 := []int{9223372036854775807}
    fmt.Println(v3, v3)
    switch "\n\t" {
    case "\n\t":
        if (v1 < v1) {
            fmt.Println(v3, v1)
        }
    case "a":
        v2[1] = true
    }
    v3 = 0.0
    v5 := GenS1{F0: 627}
    var v6 int16 = -664
    v7 := GenS2{F0: 810, F1: -500, F2: 348}
    fmt.Println(v6)
    var v8 float32 = float32(((0.1 / (3.14159 - -2.25)) * 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
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
