package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }

func main() {
    var v1 int8 = 100
    var v2 uint64 = 0
    var v3 uint64 = 42
    var v4 int32 = -2147483648
    v5 := GenS1{F0: (v1 - v1), F1: 4294967295}
    v6 := GenS1{F0: v1, F1: 47}
    var v7 int32 = v4
    var v8 uint32 = 42
    for fi9 := 0; fi9 < 9; fi9++ {
        fmt.Println(v4)
        _ = fi9
    }
    v10 := []int{127, 929}
    fmt.Println(v8)
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
