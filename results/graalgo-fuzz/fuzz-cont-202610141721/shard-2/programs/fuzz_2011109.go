package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 65535
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 uint32 = 2
    var v6 int16 = -379
    var v7 uint32 = v5
    v7 = v7
    fmt.Println(v6, v5, v2)
    v5 = (v7 % ((v5) | 1))
    var v8 int8 = -92
    var v9 int8 = v8
    fmt.Println(v5, v2, v9)
    v10 := GenS1{F0: (!true)}
    v6 -= (v6 - (v6 ^ v6))
    var v11 uint = 255
    fmt.Println(v6, v2)
    if (1 < 42) {
        v12 := GenS2{F0: v7}
        v13 := []int{}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: true}
    var v15 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
