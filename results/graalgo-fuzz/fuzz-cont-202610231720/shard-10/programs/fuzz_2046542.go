package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }
type GenS2 struct { F0 uint64; F1 string }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 32767
    var v3 string = "a"
    var v4 uint16 = v2
    var v5 uint8 = 42
    var v6 float64 = (1e10 - (0.0 / 1.5))
    var v7 uint8 = (v5 ^ v5)
    v8 := GenS1{F0: 0, F1: 73, F2: -1}
    var v9 uint32 = 2
    var v10 int16 = 1
    var v11 uint32 = v9
    if false {
        v12 := []int{0}
        _ = v12
    }
    var v13 uint64 = 887
    var v14 int8 = -122
    var v15 int16 = (-(v10))
    v16 := GenS2{F0: uint64(v2), F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
