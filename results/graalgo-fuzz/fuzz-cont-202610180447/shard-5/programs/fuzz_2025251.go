package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 2
    var v3 float64 = float64(0.1)
    var v4 float64 = 0.0
    v5 := []uint32{}
    v6 := []string{}
    var v7 int64 = 9223372036854775807
    if (v7 >= (v7 | -194)) {
        var v8 int8 = 42
        _ = v8
    }
    var v9 int16 = 681
    v4 = (float64(1e10) / (v3 + v4))
    var v10 int32 = -746
    var v11 uint16 = 1
    v12 := GenS1{F0: "hello", F1: uint8(v2), F2: 0}
    if false {
        fmt.Println(v11, v10)
    }
    if true {
        var v13 int8 = 54
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
