package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint32 = 65535
    var v2 float64 = 3.14159
    v3 := make(map[int]string)
    fmt.Println(v2, v1)
    v2 -= ((v2 * v2) * float64(1e10))
    v4 := GenS1{F0: uint(v1), F1: 42}
    if true {
        v2 -= v2
    } else {
        var v5 uint64 = 2
        fmt.Println(v1, v1, v1)
        _ = v5
    }
    fmt.Println(v2)
    v2 = (((v2 + -2.25) / 1.0) / 1.0)
    v6 := GenS2{F0: (0 > 0)}
    v7 := GenS2{F0: ((0 > 1) || (!true))}
    fmt.Println(v1, v1)
    delete(v3, -2147483648)
    v1 = (v1 | (v1 - v1))
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
