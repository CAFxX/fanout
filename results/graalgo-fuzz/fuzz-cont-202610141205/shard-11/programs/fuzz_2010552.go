package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }

func main() {
    var v1 int16 = 1000
    var v2 float64 = 0.1
    var v3 int64 = 9223372036854775807
    var v4 uint16 = 849
    var v5 int32 = -806
    fmt.Println(v4, v5, v2)
    v1 = 0
    var v6 int32 = v5
    v7 := GenS1{F0: v1, F1: (v3 << ((2) & 63))}
    v2 = 0.1
    v8 := []int{1, 32768, -88}
    var v9 int = 0
    var v10 int16 = 32767
    _ = len(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
