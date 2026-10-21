package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }

func main() {
    var v1 float64 = (0.1 - 1.5)
    var v2 int16 = 1000
    v3 := make(map[string]string)
    var v4 int8 = -128
    var v5 uint = 42
    var v6 int16 = -32768
    var v7 uint = v5
    v8 := []int{-1, 2147483647}
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint8 = 0
        _ = fi9
        _ = v10
    }
    v4 = (v4 + (-(v4)))
    var v11 uint16 = 2
    var v12 uint32 = uint32(v7)
    var v13 int32 = 1
    v14 := GenS1{F0: 42, F1: v5}
    v12 = v12
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
