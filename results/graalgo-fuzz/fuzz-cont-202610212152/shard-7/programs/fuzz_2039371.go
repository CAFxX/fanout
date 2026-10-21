package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int16; F2 int8 }
func genFn0(p0 uint64, p1 int64) int {
    fmt.Println(p1, p0)
    p0 += (p0 * p0)
    p0 = (^(p0))
    p0 -= p0
    switch -948 {
    case 127:
        p1 = p1
    case 32767:
        p1 += p1
    default:
        v6 := GenS1{F0: (p0 & p0)}
        _ = v6
    }
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 bool = (!true)
    var v2 uint64 = 588
    var v3 string = (("a" + "foo bar") + "foo bar")
    v4 := make(map[int]uint32)
    var v5 int = 169
    var v7 uint16 = 218
    fmt.Println(v5, v3)
    var v8 uint32 = 987
    v9 := v4[(v5 % (((v5 ^ (v5 * v5))) | 1))]
    _ = len(v4)
    var v10 int = -294
    var v11 int64 = -679
    if true {
        v7 = v7
        var v12 string = (v3 + v3)
        _ = v12
    } else {
        var v13 uint32 = (^(v8))
        v4[(v10 - (v5 >> ((1) & 63)))] = (v13 << ((1) & 63))
        v14 := GenS2{F0: 42, F1: 42, F2: -128}
        _ = v13
        _ = v14
    }
    v3 = v3
    v5 -= v5
    var v15 uint8 = 0
    v1 = (((1.5 + ((float32(3.14159) * 0.0) / -2.25)) < (1e10 * (float32((float64(0.0) - (0.1 - 0.0))) * float32((0.1 + -2.25))))) && (-2.25 > 0.0))
    fmt.Println(v7)
    fmt.Println(v3, v10, v10)
    fmt.Println(v10, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
