package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]uint32)
    v3 := []int{-465, 32768, 127, 4294967295}
    var v4 int8 = 47
    v5 := make(map[string]bool)
    var v6 int = 1
    v7 := GenS1{F0: (v6 >> ((4294967295) & 63)), F1: v4, F2: 4294967295}
    fmt.Println(v4, v4, v1)
    v6 += (v6 / (((v6 - 2147483647)) | 1))
    v6 = 42
    var v8 uint8 = v1
    v4 += (v4 / ((22) | 1))
    var v9 float64 = 0.1
    v8 -= v8
    v10 := GenS2{F0: 2, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
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
