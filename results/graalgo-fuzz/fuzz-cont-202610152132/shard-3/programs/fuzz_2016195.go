package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]int)
    var v3 uint32 = 1
    var v4 float64 = (-2.25 / 1e10)
    v5 := make(map[uint32]uint32)
    var v6 int8 = 1
    var v7 uint = v1
    v4 = v4
    fmt.Println(v7)
    fmt.Println(v6, v3, v4)
    v8 := GenS1{F0: -2147483648}
    var v9 float64 = v4
    if (!(-42 < -312)) {
        v2[""] = 0
        fmt.Println(v3, v7, v3)
        v9 = 0.0
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        v9 += v9
        v9 = 0.0
        _ = fi10
    }
    v11 := GenS2{F0: 2147483647, F1: (v3 * v3)}
    var v12 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
