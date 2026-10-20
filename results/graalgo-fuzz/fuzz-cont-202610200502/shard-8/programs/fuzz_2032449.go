package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0(p0 uint16, p1 uint32, p2 uint32) int {
    var v6 int = 0
    fmt.Println(p1, p2, p2)
    var v7 int32 = -1
    v8 := []int{}
    var v9 uint32 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -423
}

func main() {
    var v1 bool = (2 <= 140)
    v2 := make(map[uint32]uint32)
    var v3 float64 = float64(float32(-2.25))
    v4 := []uint32{}
    var v5 bool = (1.5 != ((-2.25 * 1.5) + ((3.14159 - 3.14159) * 0.0)))
    fmt.Println(v3)
    if (!true) {
        fmt.Println(v3, v5)
        var v10 uint32 = 2147483647
        _ = v10
    } else {
        v11 := GenS1{F0: 456, F1: 1}
        var v12 uint8 = 2
        var v13 string = ""
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 65535, F1: -396}
    v15 := GenS2{F0: 875, F1: -1}
    v4 = append(v4, 395)
    _ = len(v4)
    v16 := GenS2{F0: 1, F1: 127}
    if len(v4) > 0 {
        v4[0] = 125
    }
    v17 := GenS1{F0: 131, F1: 357}
    v18 := GenS3{F0: 0, F1: "\n\t"}
    fmt.Println(v5, v5)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
