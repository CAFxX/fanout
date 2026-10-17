package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint32 = 0
    var v2 int32 = 0
    v3 := make(map[string]bool)
    var v4 uint64 = 203
    v1 = (v1 << ((18446744073709551615) & 63))
    v5 := GenS1{F0: (v2 ^ -1), F1: 621, F2: 1}
    var v6 int32 = v2
    v7 := GenS2{F0: true}
    var v8 float64 = float64((1.5 * 1e10))
    var v9 uint16 = 32767
    var v10 uint16 = (v9 % (((v9 - v9)) | 1))
    v2 -= (-(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
