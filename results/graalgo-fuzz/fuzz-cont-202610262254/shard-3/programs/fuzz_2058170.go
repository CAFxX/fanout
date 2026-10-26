package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    var v3 float64 = 0.1
    v4 := GenS1{F0: ((!false) || true), F1: 339}
    v5 := GenS2{F0: 2}
    if (false || ((-1 == -141) || true)) {
        var v6 int8 = -1
        _ = v6
    } else {
        v3 = v3
        v3 = v3
    }
    if (0 > 0) {
        v3 -= 3.14159
    }
    _ = len(v2)
    var v7 int8 = 127
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(v7, v3, fi8)
        _ = fi8
    }
    v9 := []int{2147483647}
    v3 = float64(((0.0 * 1.5) * -2.25))
    v10 := GenS3{F0: true}
    v9 = append(v9, -1)
    _ = len(v9)
    v11 := GenS4{F0: 2147483647}
    fmt.Println(v3)
    fmt.Println(v3, v7)
    v3 += v3
    var v12 int = -2147483648
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
