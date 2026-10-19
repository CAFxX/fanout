package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 bool = true
    v2 := make(map[int]int)
    var v3 int32 = -566
    v4 := []int{-195, 127}
    var v5 uint16 = 2
    v6 := []int{}
    if (false && v1) {
        v1 = true
    }
    var v7 uint16 = v5
    v7 = (^(v5))
    if v1 {
        v8 := GenS1{F0: v3, F1: int8(v3), F2: (v3 % ((v3) | 1))}
        v9 := GenS2{F0: 65535}
        _ = v8
        _ = v9
    }
    if len(v6) > 0 {
        v6[0] = 145
    }
    fmt.Println(v1, v7, v1)
    v2[264] = -538
    switch -1 {
    case -95:
        if v1 {
            fmt.Println(v1, v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
