package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }

func main() {
    v1 := []string{}
    v2 := []int{-238, -95, 127, 92}
    v3 := make(map[uint32]bool)
    var v4 int64 = 42
    var v5 uint16 = 65535
    v6 := make(map[int]bool)
    if false {
        var v7 uint16 = (v5 ^ v5)
        var v8 uint32 = 2
        var v9 float64 = 0.0
        _ = v7
        _ = v8
        _ = v9
    } else {
        v4 = (v4 << ((uint(v5)) & 63))
        fmt.Println(v4, v5)
    }
    v4 += -451
    var v10 int8 = -128
    v11 := GenS1{F0: 32768, F1: 1, F2: 33}
    for fi12 := 0; fi12 < 10; fi12++ {
        if true {
            fmt.Println(v10)
            v5 -= v5
            var v13 uint16 = v5
            _ = v13
        }
        fmt.Println(fi12)
        _ = fi12
    }
    var v14 int = -9223372036854775808
    v14 -= (v14 ^ v14)
    var v15 float64 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
