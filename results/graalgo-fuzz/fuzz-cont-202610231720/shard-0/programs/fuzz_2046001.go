package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }

func main() {
    var v1 string = ""
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 341
    var v5 uint16 = 32767
    fmt.Println(v5, v1, v4)
    v4 += (v4 / (((v4 / ((v4) | 1))) | 1))
    if true {
        v6 := GenS1{F0: 5, F1: 255, F2: 42}
        var v7 float64 = (1e10 - 1e10)
        fmt.Println(v4, v7, v1)
        _ = v6
        _ = v7
    } else {
        var v8 int8 = 1
        var v9 int16 = 0
        _ = v8
        _ = v9
    }
    var v10 int8 = -1
    switch "\n\t" {
    case "\n\t":
        v11 := GenS1{F0: int64(v4), F1: 255, F2: 42}
        _ = v11
    case "":
        v5 += v5
    case "hello":
        v12 := v2["\n\t"]
        _ = v12
    default:
        fmt.Println(v5)
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint = 42
        _ = fi13
        _ = v14
    }
    v10 -= v10
    v15 := []int{-2147483648, 9223372036854775807, -471, -2147483648}
    var v16 uint16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
