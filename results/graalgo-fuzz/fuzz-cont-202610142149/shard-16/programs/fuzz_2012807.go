package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint8 }

func main() {
    var v1 int = 422
    var v2 int16 = 1
    var v3 uint8 = 221
    var v4 float64 = ((1e10 + 0.0) / (float64(float32(0.1)) / (1.5 + 0.0)))
    var v5 int8 = 127
    v3 = (v3 ^ 96)
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            if true {
                v8 := GenS1{F0: v3, F1: fi7, F2: 2}
                _ = v8
            } else {
                v9 := GenS2{F0: 2, F1: (v3 ^ v3), F2: v3}
                _ = v9
            }
            _ = fi7
        }
        v1 += (v1 + v1)
        _ = fi6
    }
    v10 := []int{}
    var v11 int = v1
    v12 := GenS1{F0: (v3 - (v3 ^ 128)), F1: 212, F2: 810}
    v13 := GenS1{F0: (v3 % (((v3 << ((42) & 63))) | 1)), F1: 32768, F2: 1}
    _ = len(v10)
    v14 := GenS1{F0: 2, F1: v11, F2: 42}
    var v15 int8 = (v5 | v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
