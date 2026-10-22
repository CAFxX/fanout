package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }

func main() {
    var v1 float32 = (0.1 / 0.1)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]bool)
    var v4 uint16 = 802
    v5 := make(map[int]bool)
    v6 := []int{4294967295}
    var v7 float64 = 0.1
    fmt.Println(v7, v1)
    var v8 uint = 42
    var v9 uint8 = 2
    v5[0] = true
    v10 := GenS1{F0: v8, F1: (v2 + 1), F2: -128}
    var v11 int = -9223372036854775808
    fmt.Println(v4, v2, v8)
    v11 -= (v11 % ((v11) | 1))
    v12 := GenS1{F0: (v8 >> (((v8 * v8)) & 63)), F1: (v2 + v2), F2: 127}
    var v13 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
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
