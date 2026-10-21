package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    v1 -= 3.14159
    var v5 uint = 658
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 string = (("a" + ("" + "a")) + "foo bar")
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 492, F1: 100}
    var v9 int16 = 1000
    v10 := GenS1{F0: 42, F1: 127}
    v9 = v9
    v1 = v1
    v1 += 0.0
    v11 := GenS1{F0: 875, F1: 1}
    v12 := []int{535, 2147483647, 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
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
