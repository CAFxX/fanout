package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 uint32 = 2147483647
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]bool)
    v4 = (v4 >> ((255) & 63))
    fmt.Println(v4)
    if (true || (!(224 < 119))) {
        v7 := GenS1{F0: false, F1: -646}
        var v8 int8 = -128
        _ = v7
        _ = v8
    } else {
        var v9 uint8 = 58
        _ = v9
    }
    v10 := GenS2{F0: true, F1: -119}
    v11 := []int{-2147483648, 9223372036854775807, -1}
    v4 += v4
    var v12 uint = 383
    v4 = (v4 - v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
