package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float64 = (0.0 - 1e10)
    var v2 int8 = -128
    var v3 float64 = (v1 + 1.5)
    var v4 uint8 = 1
    fmt.Println(v2, v4)
    fmt.Println(v4, v1)
    var v5 float32 = float32(((v1 - v3) / float64((((1.5 / -2.25) / (1.5 / 0.1)) / (float32(1.5) + 1e10)))))
    var v6 string = ""
    var v7 float32 = (0.0 - 1e10)
    v8 := GenS1{F0: 945}
    v9 := []int{9223372036854775807, 0, 2147483647}
    v9 = append(v9, 32768)
    v5 = (v5 / v5)
    v9 = append(v9, -1)
    var v10 uint32 = 173
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
