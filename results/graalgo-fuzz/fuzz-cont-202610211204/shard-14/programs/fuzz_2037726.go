package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -96
    var v3 int32 = -1
    v3 = 42
    var v4 uint32 = 842
    v2 = 123
    v3 += -2147483648
    v5 := GenS1{F0: false, F1: -464, F2: 65535}
    for fi6 := 0; fi6 < 12; fi6++ {
        v4 = (v4 & (v4 << ((2) & 63)))
        var v7 int32 = (-(v3))
        _ = fi6
        _ = v7
    }
    var v8 int8 = v2
    var v9 float64 = (-2.25 - 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
