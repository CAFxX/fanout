package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 42
    var v3 int32 = -1
    fmt.Println(v3, v2)
    v4 := GenS1{F0: 815, F1: 164}
    var v5 float64 = (0.0 * float64((1.5 / 1.0)))
    v2 += v2
    var v6 bool = (float64((float32((v5 * v5)) / 1.5)) == (v5 / 3.14159))
    v3 += (v2 * v3)
    var v7 int32 = (v2 >> ((255) & 63))
    var v8 float32 = 0.0
    v1[42] = 127
    v6 = (-9223372036854775808 <= 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
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
}
