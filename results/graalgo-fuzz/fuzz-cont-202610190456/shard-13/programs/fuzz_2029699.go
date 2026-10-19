package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 0
    v3 := make(map[uint32]string)
    if true {
        var v4 uint8 = 42
        for fi5 := 0; fi5 < 3; fi5++ {
            v4 -= (v4 >> ((0) & 63))
            v6 := GenS1{F0: 0, F1: 2147483647}
            _ = fi5
            _ = v6
        }
        v2 = (v2 & (v2 % ((v2) | 1)))
        _ = v4
    }
    _ = len(v1)
    v2 += v2
    v2 = (v2 * v2)
    _ = len(v3)
    v7 := GenS1{F0: -128, F1: -703}
    fmt.Println(v2)
    var v8 int32 = 0
    var v9 float32 = 3.14159
    v9 += ((float32(((1e10 * 3.14159) * 1.5)) * v9) - (3.14159 - 0.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
