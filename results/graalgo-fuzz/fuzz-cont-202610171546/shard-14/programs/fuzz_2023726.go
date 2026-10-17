package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint = 42
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    fmt.Println(v1)
    v4 := []int{42, 1, 632}
    v5 := GenS1{F0: (^(v1))}
    var v6 uint = v1
    var v7 float64 = (((((float64(((0.0 - 1.5) / 1.5)) * (float64((-2.25 + (0.0 + (0.1 - (-2.25 - float32(3.14159)))))) - (float64(3.14159) + 0.0))) / 1.0) / 0.1) * float64(0.0)) / 1.5)
    v1 -= v1
    v1 = 18446744073709551615
    var v8 float32 = 1.5
    for fi9 := 0; fi9 < 2; fi9++ {
        v1 += v1
        _ = fi9
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
