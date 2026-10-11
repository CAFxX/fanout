package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    var v1 float64 = 0.1
    v2 := []uint32{65535, 1}
    v3 := []int{-541, 32767}
    var v4 uint32 = 42
    var v5 uint64 = 2
    var v6 uint = 141
    v5 -= 419
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        v4 = v4
        v7 := GenS1{F0: uint8(v6), F1: 32767}
        _ = v7
    }
    var v8 uint8 = 2
    v4 = (v4 - (v4 - 2))
    v1 -= v1
    var v9 float64 = 1e10
    var v10 int32 = -1
    v6 = (v6 ^ v6)
    v11 := GenS1{F0: v8, F1: -48}
    fmt.Println(v1, v5, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
