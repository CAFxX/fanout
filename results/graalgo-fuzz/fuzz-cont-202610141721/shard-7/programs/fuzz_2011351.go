package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int }
type GenS2 struct { F0 uint8; F1 string }

func main() {
    var v1 uint64 = 0
    var v2 int = 2147483647
    v3 := make(map[uint32]bool)
    v4 := []uint32{1, 1}
    var v5 uint64 = v1
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(fi6, v2)
        _ = fi6
    }
    v7 := GenS1{F0: 22, F1: 558, F2: -2147483648}
    var v8 int16 = 32767
    var v9 int32 = 0
    v10 := GenS2{F0: 128, F1: "foo bar"}
    if true {
        v8 -= (v8 & (v8 << ((270) & 63)))
        v9 = v9
    }
    var v11 float32 = ((0.0 + 1.5) + (1e10 * (0.1 - 3.14159)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
