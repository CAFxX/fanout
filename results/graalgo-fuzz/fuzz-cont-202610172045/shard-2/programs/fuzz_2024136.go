package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
func genFn0() int {
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 uint64 = 4294967295
    v7 := GenS1{F0: 795, F1: 128}
    fmt.Println(v6)
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    v4 := make(map[uint32]bool)
    v8 := GenS1{F0: -83, F1: 128}
    v9 := []int{}
    var v10 float32 = 1.5
    v10 += float32(0.1)
    v11 := v3[-1]
    fmt.Println(v10, v10)
    v10 += float32(((float64(((v10 * v10) - v10)) - 0.0) - 3.14159))
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
