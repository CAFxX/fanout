package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    var v1 float64 = 1e10
    var v2 int8 = 42
    var v3 int8 = v2
    var v4 int = 0
    v2 -= v2
    v5 := GenS1{F0: -1, F1: 32767}
    var v6 float64 = (v1 + v1)
    v4 = v4
    var v7 uint64 = 2
    v2 += (v2 % (((v3 - (v3 << ((4294967295) & 63)))) | 1))
    v8 := []int{v4, (-(v4)), v4}
    if ((!true) && (true || true)) {
        fmt.Println(v2)
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
