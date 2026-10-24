package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 int = 32768
    v2 := make(map[int]string)
    v3 := []int{v1, 42, v1}
    v4 := make(map[uint32]bool)
    v1 -= -9223372036854775808
    var v5 float32 = (0.0 - (float32(3.14159) / float32(1e10)))
    if true {
        var v6 float32 = float32(1.5)
        _ = v6
    } else {
        fmt.Println(v1)
    }
    v1 += 4294967295
    var v7 int64 = -289
    v8 := GenS1{F0: false, F1: v1}
    v7 = v7
    v1 = (v1 >> ((923) & 63))
    var v9 float64 = float64(v5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v9 = v9
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
