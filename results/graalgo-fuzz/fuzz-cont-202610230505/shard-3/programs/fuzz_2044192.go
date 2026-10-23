package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int; F1 uint16 }

func main() {
    var v1 uint32 = 0
    var v2 uint32 = 2
    v3 := []uint32{v1, 1}
    var v4 int = 1
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]uint32)
    var v7 bool = true
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int = -9223372036854775808
        _ = fi8
        _ = v9
    }
    delete(v5, v2)
    var v10 float32 = (0.0 * 1e10)
    v11 := GenS1{F0: v1, F1: v1, F2: v1}
    fmt.Println(v10, v7, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (65535 < 42) {
        var v12 int8 = 100
        _ = v12
    } else {
        v13 := GenS2{F0: v4, F1: 255}
        _ = v13
    }
    delete(v5, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
