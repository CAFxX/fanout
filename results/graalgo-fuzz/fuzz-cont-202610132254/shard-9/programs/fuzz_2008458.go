package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }

func main() {
    v1 := []uint32{4294967295, 95, 2147483648}
    var v2 bool = true
    v3 := []uint32{65535, 2147483648, 65535, 42}
    var v4 uint = 448
    v5 := make(map[uint32]uint32)
    var v6 int8 = 115
    v7 := GenS1{F0: -9223372036854775808, F1: 222, F2: true}
    var v8 uint16 = 32767
    v9 := GenS1{F0: 1, F1: -812, F2: true}
    v6 -= v6
    if (v2 && v2) {
        v4 = v4
    }
    var v10 uint64 = uint64(v4)
    var v11 uint32 = 4294967295
    v4 = (v4 * v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
