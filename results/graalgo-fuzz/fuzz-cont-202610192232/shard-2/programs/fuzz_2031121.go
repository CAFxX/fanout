package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }

func main() {
    var v1 uint16 = 0
    v2 := []int{4294967295, 2147483647, 0, 42}
    v3 := make(map[int]uint32)
    v4 := []uint32{}
    var v5 uint32 = 986
    v6 := GenS1{F0: 42, F1: v1}
    var v7 string = "foo bar"
    v7 = "foo bar"
    var v8 bool = true
    v7 = v7
    v1 = 255
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
