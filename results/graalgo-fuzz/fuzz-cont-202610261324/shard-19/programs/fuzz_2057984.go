package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int32 }

func main() {
    var v1 uint16 = 2
    var v2 uint16 = v1
    v3 := []string{}
    v4 := make(map[uint32]int)
    var v5 int = 915
    var v6 uint64 = 4294967295
    v2 += (^(v2))
    v5 += (v5 * (v5 * (v5 << ((42) & 63))))
    var v7 uint32 = 2147483648
    var v8 int16 = -1
    v9 := GenS1{F0: 1, F1: 255, F2: -1}
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: -9223372036854775808, F1: 0, F2: -753}
        fmt.Println(v6, v2)
        v4[(v7 % ((4294967295) | 1))] = fi10
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 1, F1: 255, F2: 268}
    fmt.Println(v8, v8, v2)
    v13 := v4[(v7 % ((4294967295) | 1))]
    var v14 float64 = 0.1
    v3 = append(v3, "\n\t")
    v2 -= (v2 / ((1) | 1))
    v13 = (v13 + (v13 % ((-9223372036854775808) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
