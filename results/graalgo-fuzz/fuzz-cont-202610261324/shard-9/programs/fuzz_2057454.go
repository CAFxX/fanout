package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := []uint32{2147483647, 2}
    var v2 int16 = 65
    var v3 float64 = 0.0
    v4 := make(map[uint32]bool)
    var v5 int16 = (v2 ^ (v2 * v2))
    if (!false) {
        v6 := GenS1{F0: 0, F1: ((0.1 > (float32(v3) + 0.1)) || false), F2: 35}
        fmt.Println(v2, v2)
        v5 -= v2
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v5 = v5
        v5 += (v2 % ((19) | 1))
        v5 -= (v2 ^ v2)
        _ = fi7
    }
    v8 := GenS2{F0: 4294967295, F1: 2}
    var v9 int8 = -1
    var v10 int = 1
    var v11 uint = 2
    var v12 int = (v10 + (v10 - -1))
    v3 = 1.5
    switch "foo bar" {
    case "hello":
        var v13 int8 = (v9 * v9)
        _ = v13
    case "":
        var v14 bool = (v10 > -2147483648)
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 uint8 = 127
        fmt.Println(v12, v10)
        _ = fi15
        _ = v16
    }
    v12 = (^(v12))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
