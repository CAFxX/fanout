package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int64 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    var v5 int64 = 42
    var v6 int32 = -1
    fmt.Println(v6, v6)
    v5 = (v5 ^ (v5 << ((1) & 63)))
    fmt.Println(v5, v6)
    v7 := GenS1{F0: false, F1: -913, F2: (v5 ^ v5)}
    fmt.Println(v6, v5)
    v6 = v6
    fmt.Println(v6, v5)
    var v8 int = 32768
    v6 = (v6 * (v6 ^ v6))
    var v9 bool = (v8 > (v8 >> ((0) & 63)))
    fmt.Println(v5)
    v2[858] = v8
    _ = len(v4)
    var v10 int32 = v6
    v10 -= v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
