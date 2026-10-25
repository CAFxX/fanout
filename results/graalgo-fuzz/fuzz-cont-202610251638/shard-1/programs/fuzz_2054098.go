package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }

func main() {
    var v1 float64 = 1e10
    var v2 int8 = 89
    var v3 uint32 = 351
    v2 -= v2
    fmt.Println(v2)
    v3 -= (v3 % ((4294967295) | 1))
    v4 := GenS1{F0: v2, F1: 0}
    switch 128 {
    case 127:
        if (v2 < (v2 - (v2 | v2))) {
            if (-9223372036854775808 <= -118) {
                v5 := []int{127, 2147483647}
                _ = v5
            }
            fmt.Println(v3)
        } else {
            v6 := []int{-1, 112, -404}
            v7 := GenS1{F0: (v2 - (v2 ^ 100)), F1: -801}
            _ = v6
            _ = v7
        }
    }
    var v8 bool = (-9223372036854775808 <= 0)
    var v9 uint8 = 0
    v8 = true
    v1 = v1
    v10 := GenS1{F0: (v2 >> ((uint(v3)) & 63)), F1: 517}
    v8 = v8
    var v11 uint16 = 255
    v12 := []int{1}
    v13 := GenS1{F0: v2, F1: 9223372036854775807}
    fmt.Println(v9, v3)
    v3 = v3
    v14 := GenS1{F0: 0, F1: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
