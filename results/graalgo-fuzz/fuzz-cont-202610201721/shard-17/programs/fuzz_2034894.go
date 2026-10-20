package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 42
    var v3 uint32 = 426
    v4 := make(map[uint32]bool)
    var v5 uint = 2
    var v6 int16 = 42
    v1[42] = (v3 - v3)
    var v7 int32 = v2
    var v8 bool = (1 <= 467)
    var v9 int = 1
    v10 := GenS1{F0: (-(v6)), F1: 100}
    if (1e10 <= 1e10) {
        var v11 uint32 = 2
        _ = v11
    }
    v2 -= -927
    fmt.Println(v7, v8, v2)
    _ = len(v4)
    var v12 bool = ((float32(float64((0.1 + (0.0 * (((-2.25 * (0.1 + ((3.14159 + 1.5) - 1e10))) + 1.5) - 1e10))))) + 0.1) <= 0.1)
    v13 := GenS1{F0: (v6 / (((v6 | v6)) | 1)), F1: -128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
