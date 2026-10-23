package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = (!false)
    var v3 float32 = 1e10
    var v4 uint64 = 0
    var v5 uint = 18446744073709551615
    v6 := v1[2147483647]
    var v7 bool = v2
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int16 = 0
        _ = fi8
        _ = v9
    }
    v7 = (!(v2 || false))
    v4 -= (v4 % ((v4) | 1))
    for fi10 := 0; fi10 < 15; fi10++ {
        v3 = float32(1.5)
        _ = fi10
    }
    fmt.Println(v2, v2, v7)
    if v7 {
        var v11 int = -2147483648
        _ = v11
    } else {
        v12 := GenS1{F0: 2, F1: 277, F2: 2}
        fmt.Println(v6, v7)
        _ = v12
    }
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
