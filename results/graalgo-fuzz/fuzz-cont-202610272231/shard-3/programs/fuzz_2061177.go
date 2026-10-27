package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint }
func genFn0(p0 int32, p1 uint32) uint16 {
    p0 = p0
    if (439 <= 748) {
        fmt.Println(p1, p0)
    }
    v6 := []int{-448, 217, -887, -2147483648}
    _ = p0
    _ = p1
    _ = v6
    return 255
}

func main() {
    var v1 uint16 = 255
    var v2 float64 = (1e10 - (3.14159 * 1.5))
    var v3 uint16 = v1
    var v4 uint16 = (v3 * v1)
    var v5 int = -328
    v7 := []int{128, 2147483647, -7, v5}
    var v8 int8 = -128
    v9 := GenS1{F0: "", F1: -294}
    v3 = genFn0(0, 2147483647)
    v10 := GenS1{F0: "", F1: 0}
    var v11 uint32 = 65535
    if false {
        v8 -= -4
        v7 = append(v7, -1)
    }
    v12 := GenS2{F0: v8, F1: (v4 + v1), F2: 42}
    fmt.Println(v11, v5, v1)
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 bool = false
        v15 := GenS3{F0: ((v11 >> ((0) & 63)) < 2147483648)}
        var v16 int8 = -127
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := GenS4{F0: v5}
    fmt.Println(v11)
    v8 -= (-(v8))
    if true {
        var v18 uint64 = 471
        v11 -= (v11 >> ((2) & 63))
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 int16 = 912
        _ = fi19
        _ = v20
    }
    var v21 int = (v5 * (v5 % ((v5) | 1)))
    v22 := GenS5{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
