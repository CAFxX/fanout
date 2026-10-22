package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int8 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ""
    v3 := make(map[int]string)
    var v4 int64 = -1
    fmt.Println(v4, v2)
    fmt.Println(v2)
    delete(v1, 0)
    if (0.0 != (float64(1.5) - 0.0)) {
        var v5 uint16 = 2
        v2 = ""
        v6 := []int{-2147483648, 284, 128}
        _ = v5
        _ = v6
    }
    v4 = (v4 << ((2) & 63))
    v7 := []int{42, -613, -818, 151}
    _ = len(v7)
    switch 127 {
    case 32768:
        v8 := GenS1{F0: v2, F1: -75, F2: 1}
        _ = v8
    case 42:
        v9 := GenS2{F0: 42}
        _ = v9
    }
    v4 = 42
    var v10 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
