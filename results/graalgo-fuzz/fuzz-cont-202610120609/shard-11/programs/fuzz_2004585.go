package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 729
    v3 := []uint32{390, 2147483648, 0}
    var v4 uint16 = 75
    var v5 uint32 = (v2 & (v2 - (v2 & v2)))
    v6 := GenS1{F0: -2147483648}
    var v7 float64 = 1.5
    v2 += v5
    v2 = uint32(v4)
    fmt.Println(v2, v7, v2)
    v8 := v1[-678]
    v9 := GenS1{F0: -1}
    v2 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
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
