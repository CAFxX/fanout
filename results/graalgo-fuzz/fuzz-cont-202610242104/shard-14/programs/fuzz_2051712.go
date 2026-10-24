package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }

func main() {
    var v1 uint64 = 2
    var v2 uint = 42
    v3 := make(map[uint32]uint32)
    var v4 int8 = 1
    var v5 uint32 = 4294967295
    fmt.Println(v2)
    var v6 int8 = (^(v4))
    v7 := []int{-1}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint64 = 659
        _ = fi8
        _ = v9
    }
    v7 = append(v7, 0)
    v4 -= v6
    fmt.Println(v5, v6)
    _ = len(v7)
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int64 = 323
        var v12 string = "hello"
        _ = fi10
        _ = v11
        _ = v12
    }
    if len(v7) > 0 {
        v7[0] = -665
    }
    var v13 int = 2147483647
    for fi14 := 0; fi14 < 11; fi14++ {
        v2 = v2
        var v15 uint32 = v5
        v3[v15] = v15
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: v13, F1: uint(v1), F2: 2}
    fmt.Println(v4)
    fmt.Println(v1, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
