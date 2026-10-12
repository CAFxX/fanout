package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
func genFn0() uint16 {
    if true {
        var v6 float32 = (((((1e10 + 0.1) + float32(-2.25)) + 3.14159) / 1.0) / -2.25)
        var v7 int64 = -136
        v7 = -9223372036854775808
        _ = v6
        _ = v7
    }
    if false {
        var v8 uint32 = 2147483647
        var v9 uint16 = 430
        var v10 int = 32767
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 float64 = 1e10
    var v12 uint64 = 42
    _ = v11
    _ = v12
    return 922
}
func genFn1() int32 {
    v13 := GenS1{F0: -993, F1: (true && (1.5 < (float64(1.5) + (3.14159 + (0.0 + float64((0.1 * 1.5)))))))}
    v14 := []int{-296, -706}
    var v15 uint = 81
    v15 -= (v15 | (^(v15)))
    _ = v13
    _ = v14
    _ = v15
    return -1
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{1, 50}
    var v3 uint32 = 0
    v4 := make(map[int]uint32)
    v5 := []uint32{v3, v3, v3}
    delete(v4, -552)
    v3 -= 4294967295
    var v16 int32 = 42
    var v17 uint64 = 42
    v17 = 18446744073709551615
    fmt.Println(v1, v3, v17)
    v18 := v4[4294967295]
    var v19 bool = (false && false)
    v16 -= 547
    var v20 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
