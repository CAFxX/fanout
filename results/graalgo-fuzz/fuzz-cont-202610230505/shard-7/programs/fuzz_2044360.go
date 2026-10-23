package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int8; F1 string }

func main() {
    v1 := make(map[int]string)
    v2 := []int{0, -2147483648, 128, 127}
    var v3 uint32 = 4294967295
    v4 := GenS1{F0: -2147483648, F1: uint(v3)}
    var v5 int8 = 127
    v3 = v3
    fmt.Println(v5, v3)
    var v6 int64 = -1
    var v7 uint64 = 1
    var v8 int16 = 895
    v3 -= v3
    v9 := GenS2{F0: 86, F1: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
