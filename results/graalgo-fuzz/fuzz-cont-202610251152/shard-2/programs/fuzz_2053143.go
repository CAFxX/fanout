package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }

func main() {
    v1 := []string{}
    var v2 uint16 = 529
    v3 := []uint32{65535, 155, 42}
    v2 -= 65535
    v4 := GenS1{F0: 1, F1: 817}
    var v5 float32 = (1.5 + 1.5)
    v6 := GenS1{F0: 0, F1: 0}
    var v7 int = 4294967295
    if true {
        v5 += v5
        var v8 int64 = 0
        fmt.Println(v2)
        _ = v8
    }
    var v9 int = -1
    fmt.Println(v2, v2, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
