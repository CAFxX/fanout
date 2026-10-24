package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    var v1 float64 = (float64(float32(1.5)) + -2.25)
    v2 := make(map[uint32]bool)
    var v3 float64 = v1
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 42
    switch 32767 {
    case -2147483648:
        v6 := GenS1{F0: 127, F1: 0, F2: -32768}
        _ = v6
    }
    v7 := GenS2{F0: 2147483647, F1: 9223372036854775807}
    var v8 int16 = 1
    v9 := []int{}
    var v10 uint64 = 1
    v11 := GenS1{F0: 0, F1: 903, F2: (v8 + 1)}
    fmt.Println(v10, v10, v10)
    v12 := GenS1{F0: 100, F1: v10, F2: (v8 | v8)}
    if (32767 != 134) {
        var v13 int32 = 524
        v1 += v1
        _ = v13
    } else {
        for fi14 := 0; fi14 < 11; fi14++ {
            v3 = -2.25
            _ = fi14
        }
        fmt.Println(v5, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
