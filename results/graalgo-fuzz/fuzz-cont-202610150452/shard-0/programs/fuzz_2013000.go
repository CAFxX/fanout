package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    var v1 uint32 = 4294967295
    var v2 int32 = -1
    var v3 bool = false
    var v4 int8 = 35
    var v5 float64 = ((-2.25 / (1.5 - (0.0 + 1e10))) / 1e10)
    v6 := []int{}
    var v7 int8 = v4
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    v2 += (v2 % (((^(v2))) | 1))
    var v8 int64 = 9223372036854775807
    v2 -= (v2 / (((-(v2))) | 1))
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: 928, F1: (fi9 - fi9)}
        v1 = (v1 << ((61) & 63))
        v4 += v7
        _ = fi9
        _ = v10
    }
    var v11 int64 = (v8 / (((^(v8))) | 1))
    var v12 int64 = (v11 * 42)
    if v3 {
        if v3 {
            fmt.Println(v4, v3, v1)
            var v13 uint16 = 1
            v2 = (v2 + v2)
            _ = v13
        } else {
            v5 += (v5 / v5)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
