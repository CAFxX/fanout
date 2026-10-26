package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    var v1 bool = (true && false)
    var v2 int64 = 9223372036854775807
    var v3 float32 = (1.5 - 1.5)
    v4 := []int{4294967295}
    v5 := make(map[int]bool)
    v6 := GenS1{F0: v2}
    var v7 float64 = (3.14159 / 1e10)
    var v8 uint = 1
    var v9 float32 = ((1e10 - ((v3 - v3) / 1.0)) / 0.1)
    var v10 int = 4294967295
    v11 := GenS2{F0: v2}
    v12 := GenS2{F0: (v2 - v2)}
    v13 := GenS3{F0: (v8 | (v8 << (((v8 << ((2) & 63))) & 63))), F1: -379}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
