package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -1
    var v3 float64 = 3.14159
    var v4 uint32 = 873
    v5 := make(map[int]bool)
    v3 -= v3
    v3 = v3
    v2 = (v2 << ((2) & 63))
    v6 := GenS1{F0: 1, F1: 4294967295}
    v7 := []int{}
    var v8 uint64 = 69
    var v9 uint = 801
    fmt.Println(v3, v4)
    v3 = ((v3 * float64((0.0 * 0.1))) * v3)
    v3 -= -2.25
    switch -1 {
    case 42:
        v3 -= v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
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
