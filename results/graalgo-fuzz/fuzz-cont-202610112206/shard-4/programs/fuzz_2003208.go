package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 uint32 = 4294967295
    if false {
        if (4294967295 >= 2) {
            v5 := GenS1{F0: -32768, F1: -100}
            _ = v5
        }
        v6 := []int{-2147483648, 2147483647, 42}
        var v7 int = 32767
        _ = v6
        _ = v7
    }
    v3 = v3
    v8 := []int{9223372036854775807, 0, -1}
    _ = len(v8)
    v9 := GenS1{F0: -1, F1: -45}
    v4 += (v4 ^ 743)
    var v10 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
