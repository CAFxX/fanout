package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint32 }

func main() {
    v1 := []int{-552, 452}
    var v2 float64 = ((1.5 / float64(((0.1 * 1.5) * (0.0 + float32((float64((1.5 - 0.0)) + ((0.1 * (1.5 / 1.0)) / 1.0))))))) / 1.0)
    v3 := make(map[int]string)
    fmt.Println(v2)
    if (2 == 830) {
        v2 += v2
    }
    v2 = (((0.0 / 1.5) + float64(((0.0 - 0.0) * 0.0))) + v2)
    var v4 int8 = -1
    var v5 int8 = (v4 / ((100) | 1))
    var v6 bool = false
    var v7 uint = 736
    v8 := GenS1{F0: 42, F1: -1, F2: 2147483647}
    v5 = (v5 & (-(v5)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
