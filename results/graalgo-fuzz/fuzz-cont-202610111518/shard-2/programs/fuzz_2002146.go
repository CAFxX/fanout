package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 bool }

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := []string{}
    var v5 float32 = 3.14159
    v6 := make(map[uint32]int)
    var v7 uint32 = 42
    v8 := GenS1{F0: 4294967295, F1: -184}
    v9 := GenS1{F0: 2, F1: 2147483647}
    fmt.Println(v5)
    v10 := GenS2{F0: "\n\t", F1: -1, F2: -73}
    v1 = v1
    v11 := GenS2{F0: "", F1: -100, F2: -1}
    var v12 int16 = -1
    fmt.Println(v7, v5)
    v13 := GenS1{F0: 582, F1: 42}
    var v14 uint = 18446744073709551615
    var v15 int64 = int64(v1)
    v16 := GenS3{F0: ((!true) || (((-2147483648 > 902) && true) || (v15 > v15)))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
