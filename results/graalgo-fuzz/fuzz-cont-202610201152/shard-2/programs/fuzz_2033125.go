package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint8 }
func genFn0(p0 string) uint16 {
    v5 := []int{}
    p0 = p0
    p0 = ""
    v6 := GenS1{F0: 211, F1: 1000, F2: 1}
    _ = p0
    _ = v5
    _ = v6
    return 775
}
func genFn1(p0 int64, p1 uint64, p2 uint) int {
    if false {
        var v7 uint32 = 587
        _ = v7
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 32768
}

func main() {
    v1 := []uint32{}
    var v2 int8 = 100
    var v3 int = 4294967295
    v4 := make(map[uint32]string)
    var v8 int8 = v2
    delete(v4, 60)
    var v9 uint8 = 2
    fmt.Println(v8, v3, v9)
    fmt.Println(v2, v2)
    var v10 uint64 = 595
    fmt.Println(v10)
    v10 = uint64(v9)
    var v11 int64 = -1
    v11 = 42
    var v12 uint32 = 2
    v8 = (v2 + (v8 % ((v8) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
