package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 int32 = -335
    v3 += v3
    v4 := []int{}
    v3 = v3
    if true {
        var v5 uint = 37
        v3 -= -1
        _ = v5
    }
    v6 := GenS1{F0: 42, F1: -698}
    v3 -= 42
    var v7 int16 = -32768
    delete(v2, 9223372036854775807)
    v3 -= v3
    var v8 uint8 = 127
    var v9 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
