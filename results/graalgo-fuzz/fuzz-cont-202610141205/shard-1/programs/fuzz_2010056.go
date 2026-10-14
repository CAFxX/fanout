package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int64 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    var v4 int32 = 429
    v5 := GenS1{F0: 0, F1: 1, F2: 42}
    switch 42 {
    case -2147483648:
        v6 := []int{}
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := v3[fi8]
            v1 = (v1 + (^(v9)))
            v10 := []int{(^(fi8)), (fi8 * (fi7 - fi8)), fi8, -2147483648}
            _ = fi8
            _ = v10
            _ = v9
        }
        v4 += (v4 ^ v4)
        _ = fi7
    }
    v11 := GenS1{F0: 1, F1: 282, F2: 42}
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 float32 = (1e10 * -2.25)
        v14 := []int{128, (fi12 - fi12), (fi12 * fi12), 581}
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 32768, F1: 1, F2: 1}
    v4 = -2147483648
    var v16 int32 = -2147483648
    var v17 uint16 = 2
    v1 = v1
    v18 := []int{32768}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
