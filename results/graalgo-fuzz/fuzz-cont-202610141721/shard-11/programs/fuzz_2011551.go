package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = (^(v1))
    var v3 int8 = 1
    var v4 int8 = (v3 | (^(v3)))
    v3 -= (v3 ^ (v4 & v3))
    v5 := []int{-115}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 := GenS1{F0: 2, F1: (96 == 55), F2: 2147483648}
    var v7 uint32 = 42
    v4 -= (v4 << ((2) & 63))
    v8 := GenS1{F0: 18446744073709551615, F1: true, F2: uint32(v1)}
    v4 -= (v4 % (((v4 % ((v4) | 1))) | 1))
    v9 := GenS1{F0: 232, F1: false, F2: (v7 ^ 2147483647)}
    v1 = (v2 ^ v2)
    for fi10 := 0; fi10 < 19; fi10++ {
        _ = len(v5)
        var v11 bool = false
        v5 = append(v5, int(v7))
        _ = fi10
        _ = v11
    }
    if (-9223372036854775808 <= -1) {
        var v12 uint16 = 42
        fmt.Println(v2, v2)
        _ = v12
    }
    var v13 float32 = (((0.1 / 1.0) * float32(-2.25)) - 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
