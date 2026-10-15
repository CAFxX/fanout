package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{2, 1}
    var v3 int16 = 1
    var v4 float32 = 1.5
    fmt.Println(v3, v4)
    v4 = float32(((-2.25 / 1.0) / float64(v4)))
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint = 2
        _ = fi5
        _ = v6
    }
    fmt.Println(v4)
    v4 = 1e10
    v3 += (v3 + (v3 ^ -1))
    var v7 uint64 = 0
    delete(v1, -309)
    fmt.Println(v4, v3)
    v8 := v1[1]
    _ = len(v2)
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 741, F1: true}
        var v11 int8 = -1
        if false {
            var v12 float64 = 0.0
            _ = v12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v8 = v8
    var v13 float32 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
