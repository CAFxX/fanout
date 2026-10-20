package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int64 = -1
    var v2 uint64 = 42
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v5 uint16 = 2
    v6 := make(map[int]int)
    v1 = (v1 ^ v1)
    v2 = v2
    var v7 bool = (!false)
    v2 = (v2 + 1)
    if ((float32(0.0) <= (((1.5 - 0.0) * 0.0) + ((1e10 - (1.5 - -2.25)) / float32(1.5)))) || (!false)) {
        v1 += (v1 & v1)
        var v8 bool = ((v2 * (v2 ^ (v2 ^ (v2 * v2)))) != v2)
        v9 := GenS1{F0: 14, F1: 1, F2: uint32(v2)}
        _ = v8
        _ = v9
    } else {
        v1 = v1
    }
    fmt.Println(v1, v1)
    v10 := []int{}
    switch 32767 {
    case 32767:
        var v11 int32 = 0
        _ = v11
    }
    v5 -= (v5 + v5)
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int32 = 2147483647
        v14 := GenS2{F0: 880}
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        v1 = (v1 / (((-(v1))) | 1))
        v5 += v5
        _ = fi15
    }
    v16 := v4[9223372036854775807]
    var v17 uint16 = (v5 | (v5 & v5))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
