package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 int = 2147483647
    v5 := v2[(v4 >> ((0) & 63))]
    v5 = v5
    switch (v4 + v4) {
    case -9223372036854775808:
        var v6 int8 = 0
        _ = v6
    default:
        var v7 uint8 = 42
        _ = v7
    }
    v4 -= 32768
    v4 = 32768
    fmt.Println(v5, v5)
    v8 := GenS1{F0: 1, F1: 13}
    var v9 uint16 = 42
    _ = len(v3)
    var v10 int64 = 81
    if (1 != 227) {
        v5 += (v5 - v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
