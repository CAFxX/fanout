package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint16 = 282
    v2 := make(map[string]int)
    var v3 bool = false
    var v4 int8 = -128
    v5 := GenS1{F0: 900}
    fmt.Println(v4)
    var v6 float32 = 0.0
    _ = len(v2)
    v7 := []int{}
    v8 := GenS2{F0: -1}
    v9 := GenS1{F0: 748}
    v6 += v6
    v1 -= (v1 + (v1 >> ((312) & 63)))
    v10 := v2["foo bar"]
    var v11 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
