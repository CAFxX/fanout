package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 128
    v3 := []int{-2147483648, 585}
    var v4 int64 = 875
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: v4, F1: false}
        var v8 int = (-(fi6))
        v2 -= 255
        _ = fi6
        _ = v7
        _ = v8
    }
    v3 = append(v3, 1)
    if ((1e10 <= 0.1) || true) {
        var v9 int16 = 202
        v10 := GenS1{F0: v4, F1: false}
        fmt.Println(v2, v4, v4)
        _ = v10
        _ = v9
    }
    v4 = v4
    var v11 uint64 = 0
    fmt.Println(v2, v4, v2)
    var v12 float32 = -2.25
    var v13 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
