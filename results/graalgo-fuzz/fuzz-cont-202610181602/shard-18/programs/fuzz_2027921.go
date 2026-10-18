package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    v1 := []int{1}
    var v2 int32 = 2147483647
    v3 := make(map[int]uint32)
    var v4 uint64 = 0
    var v5 int = 318
    var v6 int32 = v2
    v7 := GenS1{F0: -234, F1: v5}
    v8 := GenS1{F0: -872, F1: (^(v5))}
    var v9 float32 = (1e10 / -2.25)
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 float32 = (1.5 * v9)
        _ = fi10
        _ = v11
    }
    var v12 int8 = 29
    delete(v3, v5)
    var v13 uint64 = (v4 * v4)
    var v14 int = (v5 | (v5 << ((255) & 63)))
    fmt.Println(v14, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
