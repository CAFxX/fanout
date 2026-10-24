package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -2147483648
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 2
    var v7 uint16 = 255
    v2 += v2
    v8 := GenS1{F0: 4294967295}
    v9 := []int{(v2 ^ (v2 - (v2 % ((635) | 1)))), 4294967295, v2, (v2 % ((v2) | 1))}
    v10 := v4[v2]
    v7 = (v7 + (v6 & (v7 & v7)))
    _ = len(v3)
    fmt.Println(v10, v2, v7)
    v11 := GenS2{F0: 255, F1: 1}
    v12 := GenS2{F0: 42, F1: -56}
    v6 = 601
    var v13 bool = (!false)
    v2 -= (v2 % ((9223372036854775807) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
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
    _ = v9
}
