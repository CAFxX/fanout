package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -893
    var v3 float32 = 3.14159
    v4 := make(map[string]int)
    var v5 float32 = (3.14159 + v3)
    var v6 bool = (!false)
    var v7 int8 = 1
    v6 = ((float32(0.0) / 3.14159) == float32(-2.25))
    v8 := []int{9223372036854775807, 1, 274}
    v9 := GenS1{F0: v7, F1: 83, F2: 848}
    v3 = 1.5
    v2 = (v2 * v2)
    v3 = 0.0
    v6 = ((65535 < 0) || (1000 > 1000))
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
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
