package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 string; F1 int8 }

func main() {
    var v1 float32 = (1e10 + 1e10)
    var v2 bool = (true && (42 == 4294967295))
    var v3 int16 = -32768
    var v4 uint8 = 227
    fmt.Println(v3, v1, v3)
    fmt.Println(v1, v2)
    v5 := GenS1{F0: 85, F1: 0}
    var v6 uint = 1
    v7 := []int{42, -644, 9223372036854775807}
    v2 = v2
    v4 -= (^(v4))
    v8 := GenS1{F0: 1, F1: 232}
    fmt.Println(v3)
    var v9 float32 = v1
    fmt.Println(v2, v9, v4)
    v10 := GenS2{F0: "", F1: -49}
    var v11 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
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
