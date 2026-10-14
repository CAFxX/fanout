package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint }

func main() {
    var v1 uint32 = 507
    var v2 float64 = (0.0 - -2.25)
    v3 := make(map[int]uint32)
    v4 := []string{}
    var v5 int = -543
    v6 := make(map[int]uint32)
    fmt.Println(v1, v1)
    var v7 uint32 = v1
    var v8 string = ""
    fmt.Println(v7, v7, v2)
    v6[42] = 0
    v9 := GenS1{F0: 690}
    fmt.Println(v7, v2, v8)
    v10 := GenS2{F0: 246, F1: -128, F2: 4294967295}
    fmt.Println(v8, v5)
    v11 := GenS1{F0: 1}
    var v12 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
