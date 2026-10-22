package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint16; F1 int32; F2 bool }
type GenS3 struct { F0 int64; F1 int16; F2 uint16 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 string; F1 uint64; F2 uint32 }

func main() {
    var v1 int32 = 591
    v2 := make(map[int]string)
    var v3 string = "a"
    var v4 int8 = 100
    var v5 uint = 291
    if true {
        v3 = v3
    } else {
        v3 = (v3 + v3)
        v6 := []int{-998}
        _ = v6
    }
    v7 := []int{-207}
    _ = len(v7)
    v8 := GenS1{F0: -32768, F1: 32767}
    v9 := GenS2{F0: 848, F1: (-(v1)), F2: true}
    _ = len(v2)
    var v10 uint64 = 65
    v10 = 837
    var v11 uint32 = 0
    v4 += v4
    v12 := GenS3{F0: -186, F1: -32768, F2: 42}
    v13 := GenS4{F0: -1, F1: v3}
    v14 := GenS5{F0: v3, F1: v10, F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
