package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 string }
func genFn0(p0 uint, p1 uint16, p2 bool) int64 {
    p1 = 65535
    p0 = (^(p0))
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    var v3 uint64 = 2
    v4 := []int{-805, -877, 42}
    var v5 uint = 896
    v6 := GenS1{F0: 0, F1: v5}
    v2[""] = 322
    var v7 float64 = 3.14159
    var v8 float32 = float32(1e10)
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 0, F1: v5}
        v7 = (3.14159 / v7)
        var v11 int64 = 9223372036854775807
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint32 = 42
    v13 := GenS2{F0: 195, F1: 42}
    var v14 int8 = 1
    v15 := GenS1{F0: -363, F1: v5}
    v16 := GenS3{F0: "hello"}
    fmt.Println(v3, v3)
    var v17 int = -976
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
