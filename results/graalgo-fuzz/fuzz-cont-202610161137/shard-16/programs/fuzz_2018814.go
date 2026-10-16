package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint }
func genFn0() int16 {
    var v7 int16 = 42
    switch -1 {
    case -1:
        v7 = (v7 + v7)
    }
    v8 := []int{32768, -2147483648, -225, 1}
    var v9 uint16 = 65535
    fmt.Println(v9, v9)
    _ = v7
    _ = v8
    _ = v9
    return v7
}
func genFn1(p0 int16, p1 int8, p2 string) bool {
    p2 = p2
    if (42 >= 9223372036854775807) {
        fmt.Println(p1, p2, p2)
        if true {
            p0 = (p0 % (((p0 & p0)) | 1))
        } else {
            p0 -= (p0 >> ((255) & 63))
        }
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1
    v3 := []uint32{0, 1, 1}
    v4 := []uint32{2147483648}
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]bool)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 float32 = float32(-2.25)
        _ = fi10
        _ = v11
    }
    fmt.Println(v2)
    var v12 bool = (-1 >= 0)
    fmt.Println(v2, v12)
    switch "" {
    case "\n\t":
        v13 := GenS1{F0: 85, F1: -116}
        _ = v13
    }
    var v14 int32 = 1
    var v15 uint64 = 1
    v15 += v15
    if true {
        v16 := GenS2{F0: 127, F1: 2}
        fmt.Println(v12, v14, v12)
        v12 = (0.0 > (((0.0 * (0.1 + (float32(0.0) * 1e10))) / -2.25) * (1.5 * 0.0)))
        _ = v16
    }
    fmt.Println(v2, v14, v15)
    fmt.Println(v12, v2, v2)
    var v17 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
