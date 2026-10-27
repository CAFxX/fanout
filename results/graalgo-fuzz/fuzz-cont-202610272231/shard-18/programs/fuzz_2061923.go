package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 int64 }

func main() {
    v1 := []string{}
    v2 := make(map[int]int)
    v3 := []uint32{4294967295, 1, 0}
    var v4 uint16 = 1
    var v5 int8 = 0
    fmt.Println(v4, v5)
    v3 = append(v3, 1)
    var v6 uint64 = 898
    fmt.Println(v4)
    var v7 int64 = 1
    v8 := GenS1{F0: 65535, F1: 97, F2: 1}
    v7 = -14
    v9 := GenS2{F0: 100, F1: (v5 * 113), F2: 0}
    v4 -= (v4 / ((uint16(v6)) | 1))
    fmt.Println(v4, v7)
    var v10 string = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
