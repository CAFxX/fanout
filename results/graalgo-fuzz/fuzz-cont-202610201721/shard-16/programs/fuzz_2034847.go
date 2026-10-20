package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }

func main() {
    var v1 uint32 = 2147483647
    var v2 float64 = 1e10
    var v3 int16 = -1
    var v4 uint32 = v1
    var v5 int32 = int32(v4)
    var v6 float64 = v2
    v7 := GenS1{F0: 0, F1: (false && true), F2: (91 < 42)}
    v4 -= v1
    var v8 uint16 = 255
    if (0 != 9223372036854775807) {
        v5 = v5
    } else {
        v9 := []int{-1, -658, -991, -2147483648}
        v3 = (v3 | v3)
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(fi10, v1)
        _ = fi10
    }
    v3 = v3
    v11 := GenS1{F0: 128, F1: (-2.25 < v6), F2: (-1 >= 0)}
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint16 = v8
        v14 := []int{32767}
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := []int{302}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
