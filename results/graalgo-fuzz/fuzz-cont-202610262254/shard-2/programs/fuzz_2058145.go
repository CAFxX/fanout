package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }

func main() {
    var v1 float32 = ((-2.25 + 1e10) - (float32(0.1) + 1.5))
    var v2 uint8 = 2
    var v3 int64 = 1
    if ((false || (1 != 1)) || true) {
        v1 += v1
    }
    var v4 uint16 = 958
    v1 += (0.0 + 1e10)
    v1 = (v1 / -2.25)
    v4 = (^(v4))
    v4 += 105
    v5 := []int{128, 32767, 127}
    v5 = append(v5, 42)
    var v6 uint16 = v4
    var v7 int8 = 0
    v8 := GenS1{F0: v7, F1: (v7 ^ (v7 / ((42) | 1)))}
    fmt.Println(v1, v4)
    var v9 bool = (1 >= 1000)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
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
