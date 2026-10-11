package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }

func main() {
    var v1 int8 = -128
    v2 := []int{}
    var v3 int8 = v1
    var v4 float32 = 0.1
    var v5 uint8 = 127
    var v6 int64 = -9223372036854775808
    var v7 bool = true
    v6 += v6
    v8 := GenS1{F0: (v1 ^ (v1 % (((v1 / (((v1 * v1)) | 1))) | 1)))}
    var v9 int16 = 178
    v10 := GenS2{F0: 255, F1: 0, F2: v1}
    v6 -= -9223372036854775808
    var v11 uint16 = 32767
    v1 -= (v1 - v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
