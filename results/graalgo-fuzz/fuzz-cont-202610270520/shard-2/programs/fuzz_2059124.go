package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16; F1 int16; F2 int64 }
type GenS3 struct { F0 int; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -1
    v3 := make(map[uint32]uint32)
    v4 := []string{}
    var v5 int64 = 9223372036854775807
    v6 := GenS1{F0: 80, F1: 0}
    var v7 uint32 = 2147483648
    v8 := GenS2{F0: 32767, F1: 0, F2: (^(v5))}
    var v9 bool = true
    v10 := GenS3{F0: 32767, F1: 42}
    fmt.Println(v7, v9, v7)
    v5 += v5
    v5 = v5
    var v11 uint8 = 210
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
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
