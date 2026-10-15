package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 float32 = float32(-2.25)
    v2 := make(map[int]string)
    v3 := []uint32{2147483647, 860, 2147483648}
    if false {
        v1 = (v1 + v1)
        v1 = (v1 + float32(1.5))
    }
    var v4 int32 = 781
    fmt.Println(v1, v1)
    fmt.Println(v4)
    v4 -= v4
    v1 = ((v1 - v1) - (float32(1.5) / v1))
    v4 = v4
    v5 := GenS1{F0: -128}
    v6 := GenS2{F0: 2147483647, F1: v4}
    v7 := GenS3{F0: 904}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
