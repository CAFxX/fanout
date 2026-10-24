package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (3.14159 + 3.14159)
    var v3 uint8 = 51
    v4 := make(map[string]int)
    var v5 int8 = 42
    v6 := GenS1{F0: 42, F1: 255}
    v7 := GenS1{F0: 979, F1: 160}
    v3 -= (v3 / ((v3) | 1))
    var v8 int = int(v5)
    var v9 uint64 = 878
    _ = len(v1)
    v10 := []int{}
    v11 := GenS1{F0: 704, F1: 274}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
