package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint32 = 2
    var v2 int32 = 1
    var v3 int16 = 550
    var v4 uint64 = 4294967295
    v1 = (v1 >> ((518) & 63))
    fmt.Println(v4, v1)
    var v5 uint32 = (v1 + (v1 * v1))
    v1 = v5
    v6 := GenS1{F0: 0, F1: (0 >= 2)}
    v7 := []int{-515, 127}
    v4 += v4
    v8 := GenS2{F0: -215}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
