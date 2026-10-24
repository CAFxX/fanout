package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    var v1 uint64 = 709
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    var v4 uint64 = v1
    var v5 uint = 0
    var v6 int64 = -9
    v7 := GenS1{F0: (!(!(((((-2.25 / 1e10) + 1.5) + 0.0) != float32(-2.25)) || true))), F1: 1, F2: 520}
    v8 := []int{}
    fmt.Println(v1, v5, v4)
    var v9 uint32 = 4294967295
    var v10 float32 = (0.0 / 1.0)
    var v11 int32 = 0
    v12 := GenS2{F0: -101, F1: 2}
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
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
