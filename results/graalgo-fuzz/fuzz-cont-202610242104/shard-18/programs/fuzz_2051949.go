package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }
type GenS2 struct { F0 string; F1 uint }

func main() {
    var v1 uint8 = 8
    var v2 int16 = 0
    var v3 bool = false
    var v4 int8 = 42
    v5 := make(map[string]string)
    v3 = v3
    v6 := []int{188, 0}
    v7 := GenS1{F0: 0, F1: 128, F2: 4294967295}
    if v3 {
        var v8 int8 = v4
        _ = v8
    } else {
        var v9 float64 = 0.1
        _ = v9
    }
    var v10 uint8 = 127
    v11 := GenS2{F0: ("" + "\n\t"), F1: 1}
    v3 = v3
    v10 += 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
