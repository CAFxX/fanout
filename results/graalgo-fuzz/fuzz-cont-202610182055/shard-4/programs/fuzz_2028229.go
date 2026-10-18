package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 int8, p1 int) int64 {
    p1 = (p1 - p1)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return -9223372036854775808
}
func genFn1(p0 int8) uint {
    p0 = (p0 | p0)
    fmt.Println(p0)
    _ = p0
    return 132
}

func main() {
    var v1 int8 = 0
    v2 := []int{}
    var v3 int32 = 0
    var v4 float64 = 1.5
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: "foo bar", F1: -688, F2: (v3 == v3)}
        var v9 uint64 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    v3 = v3
    var v10 float64 = (v4 - (0.1 + float64(0.1)))
    v4 = 3.14159
    var v11 int16 = 1
    fmt.Println(v11, v4, v10)
    v3 = v3
    v12 := GenS2{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
