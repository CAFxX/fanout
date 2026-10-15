package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int64 }

func main() {
    var v1 float64 = (1.5 / (0.1 * (0.1 + ((1e10 + (1.5 / (3.14159 + (3.14159 + (3.14159 * ((1.5 / 1e10) - (0.1 + (3.14159 - float64(((1.5 / float32(0.1)) + ((float32(((float64(1e10) / 0.1) - 0.1)) * 0.0) / 1e10))))))))))) * -2.25))))
    var v2 uint32 = 2
    var v3 uint8 = 42
    v4 := []int{-9223372036854775808}
    var v5 uint8 = (v3 % (((v3 * (^(v3)))) | 1))
    var v6 uint64 = 924
    v3 = 128
    fmt.Println(v1, v2)
    var v7 uint = 2
    _ = len(v4)
    var v8 float32 = 0.1
    v5 += v5
    v9 := GenS1{F0: 0, F1: 42, F2: 647}
    var v10 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
