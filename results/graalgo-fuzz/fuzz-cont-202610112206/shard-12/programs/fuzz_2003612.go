package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint }

func main() {
    var v1 float32 = 1.5
    var v2 float32 = 1e10
    v3 := []int{128, 32767}
    v4 := []string{}
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 15; fi6++ {
        v5[1] = 4294967295
        _ = fi6
    }
    var v7 bool = (-546 == 0)
    v1 = 0.0
    v7 = (54 == -7)
    var v8 uint64 = 2
    v2 -= v2
    if (1.5 == ((0.1 / float64(v1)) / ((3.14159 + -2.25) / -2.25))) {
        var v9 uint64 = 508
        _ = v9
    } else {
        var v10 uint = 4294967295
        var v11 int64 = 42
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: "a", F2: 1}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
