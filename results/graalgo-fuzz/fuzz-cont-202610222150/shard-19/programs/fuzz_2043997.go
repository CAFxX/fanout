package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }

func main() {
    v1 := []uint32{2, 42, 2147483647, 87}
    v2 := []int{-1, 114, 461}
    v3 := make(map[string]uint32)
    var v4 float32 = 0.0
    var v5 int32 = 1
    var v6 string = "foo bar"
    fmt.Println(v4, v5, v4)
    v4 = ((0.0 / (3.14159 / float32(3.14159))) - (v4 - v4))
    v4 -= 0.0
    v7 := GenS1{F0: 509, F1: 100, F2: 127}
    fmt.Println(v5, v5)
    v6 = v6
    fmt.Println(v4, v6)
    v8 := GenS1{F0: 1000, F1: -1, F2: 42}
    v9 := v3[v6]
    v9 = v9
    v6 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
