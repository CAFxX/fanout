package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int; F2 uint16 }

func main() {
    v1 := []int{4294967295, -9223372036854775808, -1, -817}
    v2 := make(map[uint32]bool)
    var v3 uint16 = 105
    var v4 int8 = -128
    var v5 int16 = -32768
    v6 := make(map[int]uint32)
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v7 string = "\n\t"
    v7 = v7
    v8 := GenS1{F0: 42, F1: v4, F2: 614}
    var v9 uint64 = 18446744073709551615
    for fi10 := 0; fi10 < 12; fi10++ {
        v9 += v9
        _ = fi10
    }
    v4 += 42
    v11 := GenS2{F0: 1, F1: 680, F2: (v3 | 42)}
    v12 := GenS2{F0: -1, F1: 42, F2: 494}
    v13 := GenS1{F0: 42, F1: (v4 << ((87) & 63)), F2: 543}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
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
    _ = v9
}
