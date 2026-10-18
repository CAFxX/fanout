package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 float64 = 3.14159
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: fi6, F1: -20, F2: "\n\t"}
        _ = fi6
        _ = v7
    }
    fmt.Println(v1)
    var v8 float64 = ((v5 / ((v5 - (v5 * v5)) + v5)) + (v5 - v5))
    v9 := GenS1{F0: 42, F1: -1, F2: ("\n\t" + "hello")}
    fmt.Println(v5)
    var v10 uint = 914
    v11 := GenS2{F0: 1, F1: (!true)}
    v12 := []int{-135}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
