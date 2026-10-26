package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]bool)
    v3 := []int{32767, 127}
    v4 := []int{4294967295, 396, 127, 127}
    v5 := []string{}
    v6 := []string{}
    var v7 float32 = 0.1
    v8 := GenS1{F0: -2147483648}
    var v9 uint8 = 255
    v7 = ((float32(((-2.25 / v1) + (0.0 * (-2.25 * float64((-2.25 * v7)))))) * 0.0) + (v7 * v7))
    v7 = v7
    var v10 float64 = (v1 * v1)
    delete(v2, 42)
    v9 -= (v9 + 1)
    v3 = append(v3, 32767)
    fmt.Println(v7, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
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
