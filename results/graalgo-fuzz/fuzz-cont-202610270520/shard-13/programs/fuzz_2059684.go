package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 float64 = -2.25
    v3 := make(map[string]string)
    var v4 uint32 = v1
    v5 := make(map[uint32]uint32)
    var v6 uint = 1
    fmt.Println(v1, v2)
    v7 := []int{-715}
    v8 := GenS1{F0: 127, F1: 42, F2: -57}
    if false {
        v2 += 1e10
    }
    var v9 uint16 = 0
    v10 := GenS2{F0: 733, F1: 118}
    v1 += v1
    v11 := GenS3{F0: -2147483648, F1: (v9 % (((v9 << ((255) & 63))) | 1))}
    var v12 bool = (0 < -100)
    var v13 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
