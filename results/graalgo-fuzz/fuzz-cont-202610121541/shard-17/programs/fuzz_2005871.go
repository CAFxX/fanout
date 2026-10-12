package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    var v1 int = -2147483648
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    var v5 int16 = 42
    v6 := GenS1{F0: 2, F1: 0}
    v1 = 4294967295
    fmt.Println(v1)
    v7 := GenS1{F0: 1, F1: 2147483647}
    var v8 int16 = v5
    v9 := []int{(v1 << ((1) & 63)), v1}
    var v10 uint32 = 0
    if true {
        if len(v9) > 0 {
            v9[0] = (v1 | v1)
        }
        v11 := GenS1{F0: uint8(v10), F1: 497}
        _ = v11
    } else {
        _ = len(v4)
        _ = len(v2)
        v9 = append(v9, v1)
    }
    var v12 uint8 = 1
    var v13 uint16 = 0
    fmt.Println(v8, v8, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
