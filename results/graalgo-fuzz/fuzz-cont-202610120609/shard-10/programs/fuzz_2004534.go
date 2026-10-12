package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }

func main() {
    var v1 int8 = 42
    var v2 int32 = -39
    var v3 float32 = 0.1
    v4 := make(map[uint32]string)
    var v5 uint64 = 2
    v6 := GenS1{F0: 42, F1: v2, F2: 4294967295}
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := v4[4294967295]
        _ = len(v4)
        _ = fi7
        _ = v8
    }
    fmt.Println(v1, v5)
    v2 = (-(v2))
    var v9 float64 = float64(v3)
    v10 := GenS1{F0: 75, F1: (v2 << ((18446744073709551615) & 63)), F2: 65535}
    var v11 uint64 = 18446744073709551615
    v12 := []int{128}
    var v13 float64 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
