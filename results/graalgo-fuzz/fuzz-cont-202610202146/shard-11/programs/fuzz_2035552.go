package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int16 }
func genFn0(p0 uint) int16 {
    p0 = p0
    var v6 int = 127
    _ = p0
    _ = v6
    return -379
}
func genFn1(p0 int32) bool {
    if (!(255 == 134)) {
        v7 := []int{-2147483648, 4294967295, -268, 0}
        _ = v7
    }
    v8 := GenS1{F0: 128, F1: 127, F2: -1}
    _ = p0
    _ = v8
    return true
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 2
    var v3 int8 = 42
    var v4 int = 348
    var v5 int32 = 464
    v9 := GenS1{F0: v4, F1: 1, F2: genFn0(42)}
    var v10 uint = 255
    v3 -= (-(v3))
    v11 := GenS1{F0: 914, F1: 128, F2: 0}
    var v12 int64 = -1
    v13 := GenS1{F0: -39, F1: 2, F2: 254}
    v5 += (v5 | (v5 / ((-1) | 1)))
    fmt.Println(v12, v12, v12)
    switch int(v12) {
    case 232:
        v1 = append(v1, v2)
    case 709:
        _ = len(v1)
    case -594:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
