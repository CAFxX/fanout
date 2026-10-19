package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    var v5 float64 = (0.0 / 1.5)
    fmt.Println(v3, v5)
    var v6 uint8 = 104
    v7 := []int{167, 0}
    v8 := GenS1{F0: v3, F1: 75}
    var v9 uint32 = 42
    if (!(1 >= 2)) {
        if (!v3) {
            fmt.Println(v5, v3)
        }
        var v10 int = -48
        fmt.Println(v10, v10, v5)
        _ = v10
    } else {
        if true {
            fmt.Println(v9)
            v5 = v5
            v5 -= (0.0 / -2.25)
        } else {
            if (-14 >= 42) {
                var v11 int64 = 1
                fmt.Println(v11, v6, v3)
                v7 = append(v7, 43)
                _ = v11
            }
        }
        v3 = (v3 || v3)
    }
    v12 := GenS2{F0: -1, F1: -128}
    if true {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v13 uint64 = 42
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
