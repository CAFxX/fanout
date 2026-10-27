package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int32 }

func main() {
    var v1 int64 = 0
    var v2 int = 949
    var v3 uint16 = 2
    var v4 bool = (-1 <= -1)
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: 4294967295, F1: v4, F2: 0}
    v7 := GenS1{F0: 42, F1: true, F2: -635}
    v8 := GenS1{F0: 2147483648, F1: v4, F2: 2147483647}
    v2 += (v2 & 1)
    var v9 float32 = -2.25
    fmt.Println(v4, v2)
    fmt.Println(v2)
    v10 := GenS1{F0: 42, F1: (float32(1.5) == -2.25), F2: -1}
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
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
