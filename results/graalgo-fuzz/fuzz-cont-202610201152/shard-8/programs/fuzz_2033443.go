package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int8 }
type GenS2 struct { F0 int64 }
func genFn0() int8 {
    v4 := GenS1{F0: 2, F1: 0, F2: -128}
    var v5 uint64 = 0
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := []uint32{4294967295, 280, 42}
    var v2 int16 = 1000
    var v3 float64 = float64((1.5 + float32(-2.25)))
    v3 += float64(0.1)
    var v6 int8 = 42
    _ = len(v1)
    fmt.Println(v3, v3)
    fmt.Println(v2)
    v7 := GenS2{F0: 561}
    var v8 float32 = 0.0
    v8 = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
