package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int64 = -975
    var v2 float64 = 0.1
    v3 := make(map[int]uint32)
    var v4 uint16 = 982
    var v5 int = 128
    if true {
        var v6 int8 = -128
        _ = v6
    } else {
        var v7 float64 = v2
        _ = v7
    }
    var v8 int64 = (v1 * 605)
    var v9 int = (v5 % (((v5 >> ((644) & 63))) | 1))
    v10 := GenS1{F0: 635}
    v3[v9] = 2147483647
    v11 := GenS2{F0: (v1 ^ 1), F1: v4}
    v2 += 1e10
    v1 = int64(v4)
    if ((128 >= 2) || false) {
        v9 += (v5 % ((v5) | 1))
        v1 = 9223372036854775807
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(v5)
        _ = fi12
    }
    var v13 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
