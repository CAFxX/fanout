package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []uint32{1, 2147483648, 0}
    var v2 uint8 = 190
    v3 := make(map[int]int)
    v2 = (v2 + v2)
    fmt.Println(v2)
    v4 := GenS1{F0: -1}
    v5 := GenS1{F0: 456}
    if (-32768 > 1) {
        v6 := GenS1{F0: -117}
        delete(v3, -2147483648)
        v2 = (v2 / ((232) | 1))
        _ = v6
    }
    _ = len(v3)
    fmt.Println(v2)
    var v7 int8 = 1
    v7 -= (v7 % ((0) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
