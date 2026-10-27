package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 float64 = float64(float32((0.0 / float64((1.5 * -2.25)))))
    v2 := make(map[int]uint32)
    var v3 int = 776
    var v4 uint = 1
    v1 = v1
    v5 := GenS1{F0: v4}
    var v6 float32 = float32(v1)
    var v7 int32 = -268
    v7 = 2147483647
    var v8 int32 = (^(v7))
    v9 := []int{(v3 ^ v3), (v3 & v3), v3, -9223372036854775808}
    v2[v3] = 65535
    var v10 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
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
