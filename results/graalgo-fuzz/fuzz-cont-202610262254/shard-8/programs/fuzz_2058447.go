package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint16; F2 uint16 }

func main() {
    var v1 string = "hello"
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    v5 := []uint32{0}
    fmt.Println(v1)
    var v6 uint16 = 1
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(v6)
        fmt.Println(v6, v1, v6)
        _ = fi8
    }
    var v9 uint8 = 134
    v10 := GenS1{F0: 65535}
    if len(v5) > 0 {
        v5[0] = 1
    }
    v11 := v4[2147483647]
    v12 := GenS2{F0: v1, F1: v6, F2: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
