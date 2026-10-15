package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
func genFn0(p0 uint32) uint16 {
    var v5 uint32 = (p0 / (((p0 | 42)) | 1))
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 931
}

func main() {
    var v1 float64 = 1e10
    var v2 int32 = -301
    var v3 int64 = -851
    var v4 int = 32768
    var v6 int32 = (^(v2))
    v1 += v1
    v7 := []int{}
    var v8 float32 = 3.14159
    v9 := GenS1{F0: -66, F1: 862, F2: 617}
    _ = len(v7)
    switch "" {
    case "a":
        var v10 uint = 51
        _ = v10
    case "":
        v3 = (v3 >> ((182) & 63))
    default:
        v1 = float64(float32((0.1 - v1)))
    }
    v11 := GenS1{F0: -83, F1: -1, F2: (^(v4))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
