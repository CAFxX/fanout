package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "a"
    var v3 float64 = 0.0
    v1[v2] = 0
    var v4 uint32 = 42
    var v5 uint = 920
    v3 += ((float64(0.1) / -2.25) * (0.1 + (((0.0 * v3) / v3) + v3)))
    v5 = (v5 % ((910) | 1))
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: -900, F1: v2, F2: (v5 % ((v5) | 1))}
        var v8 int8 = -1
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int64 = -361
    var v10 uint32 = uint32(v5)
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int64 = 1
        var v13 int32 = -333
        v5 = 42
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{-2147483648}
    switch v2 {
    case "a":
        v15 := GenS1{F0: (v9 >> (((v5 - v5)) & 63)), F1: ("foo bar" + v2), F2: 2}
        _ = v15
    case "":
        v9 += (-(v9))
    default:
        var v16 uint64 = 4294967295
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
