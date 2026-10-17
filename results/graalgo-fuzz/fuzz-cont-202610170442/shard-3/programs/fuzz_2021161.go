package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }

func main() {
    var v1 bool = (176 < 2147483647)
    v2 := []int{-2147483648}
    v3 := make(map[int]uint32)
    var v4 int64 = -9223372036854775808
    v5 := make(map[int]string)
    var v6 uint64 = 42
    var v7 uint16 = 255
    fmt.Println(v1)
    v8 := GenS1{F0: uint64(v7), F1: 0}
    v9 := GenS1{F0: (v6 + 396), F1: 259}
    var v10 float64 = 0.1
    var v11 int = 0
    var v12 int = int(v6)
    var v13 float64 = v10
    delete(v3, (v12 * (^(v12))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
