package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (3.14159 - 0.1)
    var v3 uint = 42
    v4 := make(map[int]int)
    var v5 int = 302
    v6 := GenS1{F0: 1, F1: 4294967295, F2: 0}
    var v7 uint = v3
    var v8 bool = ((^(v7)) >= 537)
    var v9 uint32 = 2147483648
    fmt.Println(v5, v9)
    v7 = 2
    var v10 int8 = -128
    v9 = v9
    var v11 uint32 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
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
