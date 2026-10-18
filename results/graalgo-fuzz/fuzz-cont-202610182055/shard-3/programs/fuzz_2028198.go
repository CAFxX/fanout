package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 bool; F1 bool }

func main() {
    var v1 float32 = ((1e10 / (0.0 - 1.5)) * 0.1)
    var v2 float64 = -2.25
    var v3 uint64 = 161
    v4 := []uint32{1, 2147483647, 2}
    v5 := GenS1{F0: 849, F1: 32767}
    v6 := GenS1{F0: 1, F1: 623}
    v1 += ((((v1 / (v1 / (v1 * v1))) + ((v1 + v1) / v1)) * ((3.14159 / v1) - ((3.14159 + 0.0) + 0.0))) / v1)
    v1 = (((v1 + (((v1 / (v1 * v1)) + (float32(v2) / v1)) - v1)) - (0.1 * (v1 / v1))) / float32(0.1))
    v4 = append(v4, 2147483647)
    fmt.Println(v1)
    v3 = v3
    var v7 uint32 = 0
    v4 = append(v4, (v7 - v7))
    var v8 float32 = ((v1 / v1) / v1)
    var v9 uint8 = 42
    v10 := GenS2{F0: 4294967295, F1: (v3 + (^(v3)))}
    v9 -= (v9 << ((0) & 63))
    v11 := GenS2{F0: 711, F1: (v3 & v3)}
    v9 -= (v9 - uint8(v3))
    v12 := GenS3{F0: true, F1: true}
    v1 += float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
