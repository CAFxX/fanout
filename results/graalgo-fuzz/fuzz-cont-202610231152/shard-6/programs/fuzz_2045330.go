package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0(p0 uint32, p1 int32, p2 bool) int32 {
    var v5 int = -9223372036854775808
    v6 := []int{}
    var v7 int8 = -13
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 int16 = -1
    fmt.Println(v4)
    if true {
        delete(v1, -603)
        delete(v2, 1)
        var v8 uint8 = 42
        _ = v8
    } else {
        fmt.Println(v4)
    }
    v9 := []int{}
    delete(v2, 0)
    v4 += v4
    if false {
        v10 := GenS1{F0: 42, F1: -9223372036854775808, F2: 127}
        _ = v10
    } else {
        v4 += v4
        v11 := GenS2{F0: (genFn0(65535, 1, false) > 0), F1: 42}
        var v12 float64 = 3.14159
        _ = v11
        _ = v12
    }
    v13 := GenS3{F0: 1, F1: 42}
    var v14 bool = false
    fmt.Println(v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
