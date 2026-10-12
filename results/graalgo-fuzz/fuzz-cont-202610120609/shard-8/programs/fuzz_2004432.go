package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    var v3 float64 = 1.5
    var v4 uint16 = 0
    v5 := make(map[string]bool)
    v3 += v3
    v4 -= v4
    v6 := v5["hello"]
    v7 := GenS1{F0: 0, F1: 42}
    var v8 int8 = 1
    var v9 uint16 = v4
    var v10 uint64 = 2
    var v11 uint = 2
    var v12 int32 = 1
    var v13 uint16 = v9
    var v14 int32 = 0
    v6 = ((1.5 * (v3 + 0.0)) < ((0.1 - (((((v3 * (float64(-2.25) - (float64(-2.25) + v3))) * 0.0) / v3) + v3) * (1e10 / (v3 + v3)))) + v3))
    var v15 int16 = 465
    v10 += (v10 >> (((v11 + v11)) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
