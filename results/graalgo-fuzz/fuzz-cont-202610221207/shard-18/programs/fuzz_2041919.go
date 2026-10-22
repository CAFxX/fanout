package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 2
    v3 := make(map[uint32]int)
    var v4 uint32 = 0
    var v5 uint64 = 0
    v6 := make(map[string]int)
    v7 := GenS1{F0: 42, F1: 2147483648}
    delete(v6, "\n\t")
    v2 -= (v2 << ((18446744073709551615) & 63))
    var v8 int64 = 42
    v2 -= v2
    v9 := GenS1{F0: 0, F1: v4}
    v2 -= (v2 ^ v2)
    var v10 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
