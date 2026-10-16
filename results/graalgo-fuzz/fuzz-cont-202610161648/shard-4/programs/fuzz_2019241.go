package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 float32 = (0.0 * 0.0)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    v4 := []int{127}
    var v5 uint8 = 255
    v5 -= v5
    fmt.Println(v5)
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int64 = 42
        v5 = v5
        var v8 float32 = v1
        _ = fi6
        _ = v7
        _ = v8
    }
    v1 = 3.14159
    v9 := GenS1{F0: (42 < -527)}
    v1 -= (float32(0.0) / (v1 / ((1e10 - 0.1) / (0.1 - (((v1 - v1) + 0.0) - v1)))))
    switch 128 {
    case 834:
        var v10 bool = true
        _ = v10
    }
    var v11 uint64 = 205
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
