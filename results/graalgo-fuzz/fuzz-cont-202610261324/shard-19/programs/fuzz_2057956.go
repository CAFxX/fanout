package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint16, p2 int16) int8 {
    fmt.Println(p1)
    p1 = (^(p1))
    var v5 float32 = (0.1 + 3.14159)
    v6 := GenS1{F0: (p2 | p2), F1: p0, F2: -1}
    var v7 int64 = int64(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -29
}
func genFn1(p0 int, p1 int64, p2 int) uint8 {
    if true {
        fmt.Println(p0, p2)
        v8 := []int{(p2 << ((18446744073709551615) & 63)), (p0 >> ((0) & 63)), p2}
        fmt.Println(p0, p0)
        _ = v8
    }
    var v9 int16 = 1000
    if false {
        var v10 uint16 = 1
        _ = v10
    }
    p0 -= p0
    v9 = v9
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 255
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]bool)
    var v3 uint64 = 4294967295
    v4 := make(map[string]int)
    v3 -= 747
    fmt.Println(v3, v1)
    v11 := GenS1{F0: 1000, F1: 128, F2: -1}
    v12 := []int{32768, int(v1)}
    var v13 uint8 = 2
    v12 = append(v12, 127)
    v14 := GenS1{F0: 1000, F1: (v13 ^ v13), F2: 33}
    var v15 int32 = -468
    v16 := GenS2{F0: v13}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
