package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int16 = 32767
    v2 := []uint32{}
    v3 := make(map[int]int)
    v4 := make(map[uint32]int)
    v5 := []uint32{0, 0, 181}
    v6 := []uint32{2147483648, 2147483648, 766}
    v7 := GenS1{F0: 1, F1: 2147483647}
    var v8 uint64 = 111
    v4[4294967295] = -886
    if len(v2) > 0 {
        v2[0] = 0
    }
    var v9 float64 = (3.14159 / 1e10)
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int16 = v1
        fmt.Println(v8)
        _ = fi10
        _ = v11
    }
    v8 = 42
    var v12 int = 4294967295
    v13 := GenS2{F0: v1}
    v8 = (v8 | 18446744073709551615)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
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
