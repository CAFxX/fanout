package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint16 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]uint32)
    var v3 bool = false
    v4 := []int{-713, 675, -2147483648, 2147483647}
    v5 := GenS1{F0: 1}
    v6 := GenS1{F0: -662}
    v1 += 0.0
    v7 := GenS1{F0: 9223372036854775807}
    var v8 int32 = 2147483647
    _ = len(v4)
    fmt.Println(v8, v1, v8)
    v9 := GenS2{F0: -1, F1: 0, F2: 42}
    v3 = v3
    fmt.Println(v8, v3, v1)
    v1 = 0.0
    v10 := GenS2{F0: 1, F1: 598, F2: 206}
    var v11 float64 = v1
    _ = len(v4)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
