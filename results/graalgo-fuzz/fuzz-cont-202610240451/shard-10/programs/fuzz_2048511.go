package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int16; F1 uint16; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 int16 = 0
    var v4 int8 = -124
    var v5 bool = true
    v4 += (v4 * (v4 + 1))
    v6 := GenS1{F0: 42, F1: v5}
    v3 = (-(v3))
    var v7 float64 = float64(0.0)
    v3 += 0
    fmt.Println(v3)
    _ = len(v2)
    v5 = (v5 && ((v5 && v5) && ((1.5 > 3.14159) && v5)))
    v4 += v4
    v4 -= (v4 % ((v4) | 1))
    v2[87] = 1
    v8 := GenS2{F0: v3, F1: 0, F2: 42}
    v9 := []int{0, 9223372036854775807, -9223372036854775808}
    switch 809 {
    case 0:
        for fi10 := 0; fi10 < 11; fi10++ {
            v5 = (!(v7 == 0.1))
            var v11 uint8 = 3
            _ = fi10
            _ = v11
        }
    }
    var v12 uint32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
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
