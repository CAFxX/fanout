package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }

func main() {
    var v1 float32 = float32(1.5)
    v2 := make(map[uint32]uint32)
    var v3 string = "foo bar"
    v4 := make(map[int]bool)
    var v5 int16 = 624
    v5 += (-(v5))
    v5 = (v5 ^ (v5 << ((369) & 63)))
    v6 := GenS1{F0: 590, F1: 958, F2: -1}
    v5 -= (^(v5))
    var v7 float32 = float32((1e10 + 0.1))
    if true {
        var v8 uint64 = 18446744073709551615
        _ = v8
    }
    var v9 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
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
