package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint64 }

func main() {
    v1 := []int{9223372036854775807, -1}
    var v2 string = "hello"
    var v3 int64 = 1
    var v4 uint64 = 0
    v4 += 1
    var v5 string = v2
    v6 := GenS1{F0: v5, F1: v3, F2: v4}
    v2 = v5
    v5 = v2
    v3 = 42
    if false {
        v7 := GenS1{F0: "\n\t", F1: -1, F2: 2}
        _ = v7
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
