package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    var v1 float32 = -2.25
    var v2 int32 = 8
    v3 := make(map[string]uint32)
    var v4 int = 0
    v2 = v2
    var v5 float64 = 3.14159
    v3["a"] = 70
    v6 := v3[("a" + "hello")]
    var v7 int16 = 0
    v8 := GenS1{F0: 2, F1: 42}
    v9 := []int{(v4 / ((32767) | 1)), v4, (v4 >> ((uint(v6)) & 63))}
    _ = len(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
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
