package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := []uint32{2147483647, 65535, 2147483647}
    v2 := make(map[int]uint32)
    v3 := []int{-1, 1}
    var v4 int = 9223372036854775807
    v5 := []string{}
    v4 += (v4 / (((v4 * v4)) | 1))
    var v6 int32 = 1
    v7 := GenS1{F0: "a", F1: (v4 * (-(v4)))}
    if true {
        v4 = v4
        v8 := GenS2{F0: 199, F1: 255}
        _ = v8
    }
    v6 += (-(v6))
    var v9 float32 = -2.25
    var v10 float32 = v9
    fmt.Println(v10, v10)
    fmt.Println(v4, v9, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
