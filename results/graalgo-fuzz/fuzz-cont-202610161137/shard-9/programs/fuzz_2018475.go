package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 0.0
    var v3 int64 = -1
    var v4 uint = 42
    var v5 int64 = (v3 - 1)
    v6 := []int{}
    v1[-9223372036854775808] = (v4 == 404)
    v7 := GenS1{F0: v3, F1: 144}
    var v8 uint32 = 0
    v2 -= 1.5
    var v9 uint8 = 106
    v9 -= 127
    if (0 < v9) {
        var v10 uint16 = 0
        _ = v10
    }
    v8 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
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
