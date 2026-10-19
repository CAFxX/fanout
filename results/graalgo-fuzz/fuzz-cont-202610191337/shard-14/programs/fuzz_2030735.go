package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }

func main() {
    v1 := []uint32{0}
    var v2 float64 = 3.14159
    v3 := make(map[int]bool)
    var v4 int8 = -1
    fmt.Println(v2, v2)
    v5 := GenS1{F0: -898, F1: 1}
    delete(v3, 4294967295)
    v4 += -54
    var v6 uint16 = 2
    v3[-1] = (0 >= 42)
    v6 -= (v6 % ((437) | 1))
    v4 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
