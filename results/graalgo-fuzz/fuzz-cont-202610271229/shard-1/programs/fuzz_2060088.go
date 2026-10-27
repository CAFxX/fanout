package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 1
    v3 := make(map[string]uint32)
    var v4 float32 = 0.1
    if (2 > 18446744073709551615) {
        var v5 float32 = float32((float64(v4) * -2.25))
        v6 := GenS1{F0: v2, F1: 2, F2: 538}
        _ = v5
        _ = v6
    }
    var v7 uint = 1
    fmt.Println(v2)
    v8 := v3["\n\t"]
    fmt.Println(v7)
    v9 := GenS2{F0: 230}
    fmt.Println(v4, v7, v7)
    var v10 int64 = -9223372036854775808
    v10 = 0
    v10 = -482
    if false {
        if false {
            v10 += v10
        }
        v2 = (v2 & (v2 | 233))
    } else {
        v7 += (v7 ^ 590)
        v2 -= (v2 * uint8(v7))
    }
    v11 := []int{}
    var v12 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
