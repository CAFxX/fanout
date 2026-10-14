package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint8 }

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]uint32)
    v3 := []uint32{402, 542}
    _ = len(v2)
    fmt.Println(v1)
    var v4 string = ("foo bar" + "hello")
    fmt.Println(v4, v1)
    v5 := GenS1{F0: -1, F1: (!true), F2: 127}
    v6 := GenS1{F0: 42, F1: (820 >= 0), F2: 42}
    var v7 int32 = -2147483648
    var v8 float32 = (-2.25 - 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
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
