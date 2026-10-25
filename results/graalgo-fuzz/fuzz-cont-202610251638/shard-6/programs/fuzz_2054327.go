package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint8; F2 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 1
    var v3 uint32 = 2147483647
    var v4 uint8 = 42
    var v5 uint8 = (^(v4))
    var v6 int = -333
    var v7 float64 = (1e10 + 0.1)
    v1["hello"] = 42
    v8 := GenS1{F0: (!true)}
    v9 := GenS2{F0: v6, F1: v5, F2: 42}
    var v10 uint = 144
    v11 := GenS2{F0: v6, F1: 127, F2: 42}
    for fi12 := 0; fi12 < 20; fi12++ {
        v10 = v2
        _ = fi12
    }
    var v13 int64 = -9223372036854775808
    v5 = (v5 ^ (v5 >> (((v2 * v10)) & 63)))
    v14 := []int{2147483647, (-(v6)), 501}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
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
