package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    var v1 float64 = (0.1 * 0.0)
    v2 := make(map[uint32]uint32)
    var v3 bool = false
    var v4 float32 = (0.1 + 3.14159)
    var v5 int8 = -42
    var v6 string = "hello"
    v7 := []int{-1, 4294967295, 128}
    for fi8 := 0; fi8 < 7; fi8++ {
        _ = len(v7)
        var v9 int32 = 42
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 float32 = (v4 + 0.1)
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: v3}
    var v13 int32 = 0
    v14 := GenS1{F0: -764, F1: (float64(((v4 - v4) * v4)) == ((1.5 / float64(v4)) + v1))}
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS1{F0: 482, F1: false}
        _ = fi15
        _ = v16
    }
    v6 = v6
    var v17 int8 = v5
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
