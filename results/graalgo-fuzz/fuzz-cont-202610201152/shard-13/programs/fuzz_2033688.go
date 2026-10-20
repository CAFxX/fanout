package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }

func main() {
    var v1 uint8 = 2
    var v2 uint32 = 2
    var v3 uint64 = 0
    v4 := make(map[int]uint32)
    v1 = v1
    var v5 uint16 = 65535
    if (!false) {
        v6 := GenS1{F0: 403, F1: -1}
        v7 := GenS1{F0: (v3 & uint64(v1)), F1: -22}
        v8 := v4[9223372036854775807]
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int64 = 1
    v10 := []int{0, 42, 32768}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v9 -= (v9 & -9223372036854775808)
    if (9223372036854775807 <= v9) {
        v5 = (v5 / (((^(v5))) | 1))
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v2 = (v2 | 499)
        _ = fi11
    }
    v12 := GenS1{F0: v3, F1: 10}
    var v13 int16 = -32768
    v13 = (v13 & v13)
    var v14 int16 = (^(v13))
    var v15 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
