package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int16 = 1
    var v2 float64 = (0.0 + (1e10 * (-2.25 * 1e10)))
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    v5 := make(map[int]uint32)
    var v6 float64 = (v2 * v2)
    fmt.Println(v6, v2)
    var v7 uint16 = 865
    v1 = v1
    v8 := v4[("hello" + "hello")]
    var v9 bool = true
    var v10 int8 = 1
    var v11 int = -625
    v2 += (v2 / ((1.5 * v2) - 1.5))
    v12 := v4["a"]
    v3[v12] = ((^(v12)) == 712)
    var v13 uint64 = 1
    v5[0] = v8
    v14 := GenS1{F0: v9}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
