package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[int]uint32)
    v3 := []int{0}
    fmt.Println(v1)
    v1 = ((v1 - ((-2.25 - v1) + v1)) - (1.5 + (1e10 + v1)))
    v4 := v2[127]
    v4 = (v4 & v4)
    var v5 int = -2147483648
    v2[(v5 * (v5 << ((2) & 63)))] = 387
    v6 := GenS1{F0: 2147483647}
    v4 -= v4
    var v7 int64 = 42
    v3 = append(v3, -798)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
