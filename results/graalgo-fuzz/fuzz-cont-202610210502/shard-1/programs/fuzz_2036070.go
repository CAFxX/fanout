package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]string)
    var v3 string = v1
    v4 := make(map[int]string)
    var v5 uint16 = 1
    var v6 float64 = (0.1 * 0.0)
    var v7 uint64 = 4294967295
    _ = len(v2)
    var v8 float64 = (v6 - v6)
    switch 42 {
    case -2147483648:
        fmt.Println(v6, v5)
    case 32767:
        v8 -= v6
    case 32768:
        _ = len(v2)
    default:
        v9 := GenS1{F0: 100, F1: 65535}
        _ = v9
    }
    var v10 uint16 = 42
    fmt.Println(v10)
    delete(v4, 32767)
    var v11 int = 2147483647
    v12 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
