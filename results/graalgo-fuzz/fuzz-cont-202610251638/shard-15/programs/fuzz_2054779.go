package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int32; F1 uint; F2 int8 }

func main() {
    var v1 uint = 572
    v2 := make(map[int]bool)
    var v3 float32 = (3.14159 + -2.25)
    var v4 string = ("hello" + "hello")
    v5 := GenS1{F0: 362, F1: v1}
    v3 -= v3
    v1 -= v1
    v6 := GenS2{F0: -1, F1: (v1 << (((v1 | 106)) & 63)), F2: -128}
    v3 = (-2.25 / -2.25)
    v3 += 3.14159
    var v7 uint64 = 0
    v8 := GenS1{F0: 42, F1: v1}
    var v9 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
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
