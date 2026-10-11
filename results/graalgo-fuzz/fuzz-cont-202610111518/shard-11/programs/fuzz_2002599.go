package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int; F1 uint64 }

func main() {
    v1 := []int{}
    var v2 uint64 = 42
    var v3 float32 = (0.0 / 1.5)
    var v4 int32 = 0
    v5 := make(map[uint32]int)
    var v6 int32 = v4
    var v7 float32 = v3
    fmt.Println(v2)
    v8 := GenS1{F0: 32767, F1: 42}
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS2{F0: (fi9 >> ((327) & 63)), F1: (v2 + 42)}
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: -1, F1: v2}
    v6 -= -2147483648
    v3 = (3.14159 - v7)
    v5[513] = 42
    v6 = v4
    var v12 int32 = (v6 ^ v6)
    var v13 uint32 = 65535
    v3 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
