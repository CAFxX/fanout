package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 uint = 0
    var v5 uint16 = 2
    var v6 int = -969
    v5 -= 255
    v6 -= (v6 + 4294967295)
    v4 = 42
    v6 = (v6 | v6)
    var v7 int32 = 2147483647
    v8 := GenS1{F0: 18446744073709551615, F1: -1}
    var v9 float64 = (0.0 - 1e10)
    v10 := []int{-9223372036854775808}
    v9 -= v9
    v9 = (((v9 - 1e10) + 1.5) + float64((0.1 * -2.25)))
    v9 = ((v9 * float64((3.14159 * 0.0))) + (v9 * v9))
    v4 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
