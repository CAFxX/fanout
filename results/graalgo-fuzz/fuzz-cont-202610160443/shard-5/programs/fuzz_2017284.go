package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 float64 = float64(-2.25)
    var v5 int = 32767
    v6 := []int{128, (-(v5)), (v5 / ((v5) | 1)), v5}
    v4 += v4
    var v7 int64 = 9223372036854775807
    fmt.Println(v4, v7)
    if (false || (211 < 1)) {
        var v8 bool = false
        v7 -= v7
        v9 := GenS1{F0: 1, F1: -1}
        _ = v8
        _ = v9
    } else {
        v6 = append(v6, (v5 - 0))
    }
    v7 -= v7
    v4 -= (((-2.25 + -2.25) - v4) * float64(-2.25))
    var v10 string = ("foo bar" + "hello")
    if true {
        v1[2147483648] = 4294967295
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
