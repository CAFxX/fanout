package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    v1 := []int{42}
    var v2 uint = 32
    v3 := []uint32{2147483648, 2147483648, 521}
    v4 := []uint32{2147483648, 65535}
    v2 -= v2
    v2 = v2
    var v5 float64 = 3.14159
    v2 += v2
    var v6 uint32 = 0
    switch -351 {
    case 4294967295:
        _ = len(v1)
    case -9223372036854775808:
        var v7 float32 = 0.0
        _ = v7
    case 0:
        var v8 float32 = -2.25
        _ = v8
    }
    v9 := GenS1{F0: (v6 + (^(v6))), F1: false}
    v10 := GenS1{F0: v6, F1: true}
    var v11 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
