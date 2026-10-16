package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint8 = 56
    var v2 int64 = -643
    var v3 int8 = 1
    v4 := make(map[int]uint32)
    var v5 uint64 = 264
    var v6 uint8 = v1
    var v7 bool = false
    v7 = (!v7)
    fmt.Println(v7, v7)
    v8 := []int{444}
    v7 = false
    v6 = (v6 >> ((414) & 63))
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v3, v6)
        _ = fi9
    }
    v10 := GenS1{F0: 42, F1: (v1 << ((42) & 63))}
    fmt.Println(v5, v5)
    _ = len(v8)
    v11 := GenS2{F0: 128}
    var v12 uint64 = (v5 >> ((177) & 63))
    v4[-141] = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
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
    _ = v8
}
