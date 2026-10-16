package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint, p1 int8) int16 {
    v5 := GenS1{F0: "\n\t", F1: (765 <= -432)}
    v6 := GenS1{F0: "", F1: ((-2.25 * -2.25) <= 3.14159)}
    var v7 uint32 = 665
    p1 -= -128
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint16, p1 int16) uint64 {
    var v8 float32 = 1e10
    p0 += p0
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: -105}
        v11 := []int{(fi9 / ((fi9) | 1))}
        v11 = append(v11, fi9)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]int)
    var v3 int16 = 1000
    v4 := make(map[int]uint32)
    var v12 string = "a"
    var v13 uint = 255
    if true {
        var v14 int32 = -1
        var v15 int32 = (v14 ^ v14)
        _ = v14
        _ = v15
    }
    fmt.Println(v13, v3)
    v12 = "a"
    v1 -= (v13 ^ (^(v1)))
    v13 = (^(v13))
    if true {
        var v16 int16 = genFn0(v1, -128)
        fmt.Println(v12, v16)
        fmt.Println(v12, v12)
        _ = v16
    }
    var v17 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
