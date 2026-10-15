package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 float64 = 1e10
    var v5 uint8 = 0
    v3 = true
    v6 := GenS1{F0: 42, F1: 255}
    var v7 uint16 = 42
    var v8 uint = 320
    var v9 float32 = float32(v4)
    fmt.Println(v9)
    v3 = v3
    v9 += (v9 + -2.25)
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
