package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }

func main() {
    var v1 bool = true
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 uint8 = 243
    var v5 uint8 = v4
    v1 = v1
    var v6 int32 = 1
    v7 := []int{32768, -868, -9223372036854775808}
    var v8 float32 = 1e10
    if (!(v1 && v1)) {
        if (1 > 760) {
            var v9 float64 = ((0.1 * 3.14159) - (1.5 * (1.5 / 0.1)))
            _ = v9
        }
        v6 = (v6 % (((-(v6))) | 1))
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 bool = v1
            _ = fi10
            _ = v11
        }
    }
    var v12 int8 = 22
    v4 += v5
    v8 += (((v8 / v8) / -2.25) - (v8 * v8))
    v13 := GenS1{F0: 2, F1: 565, F2: -390}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
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
}
