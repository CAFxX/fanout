package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int8; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int8) uint {
    v6 := GenS1{F0: 498, F1: -1}
    v7 := []int{455, 175, -519}
    _ = p0
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -89
    var v3 int = 358
    v4 := []int{604}
    v5 := make(map[uint32]int)
    var v8 int16 = v2
    v3 = (v3 - (v3 - v3))
    _ = len(v1)
    var v9 string = ""
    v10 := GenS2{F0: -118, F1: -322, F2: 65535}
    v8 -= v2
    v11 := GenS3{F0: 571, F1: 2147483647}
    fmt.Println(v3, v2)
    if false {
        v3 -= (v3 % (((v3 - v3)) | 1))
        delete(v5, 2)
    }
    var v12 int64 = 321
    fmt.Println(v12, v8, v8)
    var v13 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
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
    _ = v8
    _ = v9
}
