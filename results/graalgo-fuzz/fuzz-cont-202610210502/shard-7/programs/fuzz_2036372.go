package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    var v1 float32 = -2.25
    var v2 bool = false
    var v3 float64 = (-2.25 - 1.5)
    var v4 float32 = 0.1
    v5 := GenS1{F0: false, F1: 42}
    v4 = (0.1 / 1.5)
    v6 := GenS1{F0: v2, F1: 4294967295}
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint16 = 1
        _ = fi7
        _ = v8
    }
    fmt.Println(v1, v4, v4)
    var v9 uint32 = 269
    fmt.Println(v2)
    v10 := []int{0, 0, -99, 1}
    var v11 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
