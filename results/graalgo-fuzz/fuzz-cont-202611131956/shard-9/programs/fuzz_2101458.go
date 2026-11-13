package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    v1 := []uint32{65535}
    v2 := []int{}
    var v3 float32 = (3.14159 * (0.1 / 1.5))
    var v4 uint8 = 2
    var v5 int32 = -123
    v6 := make(map[uint32]uint32)
    v3 = float32((0.0 + 0.1))
    fmt.Println(v4, v3, v3)
    v7 := GenS1{F0: -2147483648, F1: "\n\t", F2: -166}
    v8 := GenS1{F0: 38, F1: "", F2: -171}
    v9 := GenS2{F0: 2, F1: 42}
    v3 = -2.25
    v5 = -126
    if len(v1) > 0 {
        v1[0] = 65535
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
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
