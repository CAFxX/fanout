package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int64 }
func genFn0(p0 int8, p1 int8, p2 bool) bool {
    v5 := []int{}
    p2 = true
    if p2 {
        p1 -= p1
    }
    var v6 int16 = 820
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ((float32((1e10 - 0.0)) >= 0.1) || (4294967295 != 731))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 0.0
    v3 := make(map[string]uint32)
    var v4 int8 = 53
    v7 := GenS1{F0: 9223372036854775807, F1: 1, F2: 911}
    fmt.Println(v2)
    v4 -= 120
    v4 -= v4
    var v8 float32 = 1e10
    var v9 int32 = 42
    switch 2147483647 {
    case -749:
        var v10 uint = 0
        _ = v10
    case 1:
        _ = len(v3)
    }
    v2 -= 3.14159
    var v11 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
