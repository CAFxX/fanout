package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int32 = 0
    v2 := make(map[int]uint32)
    var v3 int32 = v1
    var v4 uint = 1
    var v5 float64 = (-2.25 + (0.1 + (1e10 * 1e10)))
    fmt.Println(v1)
    v6 := []int{127}
    var v7 int32 = 187
    var v8 float64 = v5
    v4 -= v4
    if true {
        v4 += (^(v4))
        var v9 uint64 = 2
        _ = v9
    } else {
        v10 := GenS1{F0: 2}
        _ = v10
    }
    var v11 uint = 18446744073709551615
    v12 := GenS1{F0: 224}
    var v13 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
