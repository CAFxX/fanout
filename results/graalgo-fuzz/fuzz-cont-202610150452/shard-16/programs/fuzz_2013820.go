package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint32; F1 int8 }

func main() {
    var v1 bool = (!true)
    var v2 uint = 18446744073709551615
    var v3 int = 32768
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    v6 := make(map[int]bool)
    fmt.Println(v3)
    v7 := []int{2147483647, (v3 - (v3 % ((v3) | 1))), 2147483647, 9223372036854775807}
    if v1 {
        _ = len(v7)
        var v8 int16 = -32768
        _ = v8
    }
    var v9 int8 = 42
    v10 := GenS1{F0: 32767, F1: (292 == 566)}
    v11 := GenS2{F0: 2147483647, F1: -59}
    var v12 uint = v2
    v13 := GenS1{F0: int16(v2), F1: v1}
    var v14 float64 = ((1.5 / 1.0) - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
