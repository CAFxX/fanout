package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]bool)
    var v4 int32 = 1
    var v5 int = 42
    fmt.Println(v4, v4, v4)
    var v6 float32 = 0.0
    v4 -= 1
    v2 += (v2 + v2)
    v7 := GenS1{F0: 0}
    fmt.Println(v6)
    v6 -= (((v6 - v6) + (0.0 - v6)) / v6)
    var v8 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
