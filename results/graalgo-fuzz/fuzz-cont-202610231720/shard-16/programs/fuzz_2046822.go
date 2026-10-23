package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
func genFn0(p0 int, p1 uint32, p2 bool) int8 {
    v5 := GenS1{F0: 158, F1: -742}
    var v6 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 float32 = 1e10
    var v2 int8 = 127
    var v3 uint64 = 0
    v4 := []uint32{2, 2147483648, 1}
    v7 := GenS1{F0: 703, F1: 32767}
    v3 += v3
    var v8 int = -876
    var v9 uint64 = 42
    v1 = 0.1
    var v10 int8 = (v2 << ((65535) & 63))
    v2 -= (v2 ^ v2)
    v11 := GenS1{F0: 0, F1: 32767}
    fmt.Println(v10)
    var v12 float64 = 1.5
    var v13 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
