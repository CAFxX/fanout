package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int; F2 int32 }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v5 int16 = 32767
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: 120}
    v8 := v4[42]
    v1 -= (0.1 + v1)
    var v9 string = "\n\t"
    v5 = -1
    v10 := []int{}
    if true {
        for fi11 := 0; fi11 < 9; fi11++ {
            if len(v10) > 0 {
                v10[0] = (fi11 % ((fi11) | 1))
            }
            _ = fi11
        }
        v8 -= v8
    }
    v12 := GenS2{F0: -85, F1: -2147483648, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
