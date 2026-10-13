package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 int32 = 0
    var v2 uint8 = 0
    v3 := make(map[uint32]int)
    var v4 int16 = -681
    var v5 bool = false
    v6 := []int{2147483647}
    if (false || true) {
        if len(v6) > 0 {
            v6[0] = 1
        }
        var v7 uint32 = 2
        _ = v7
    }
    v2 = v2
    var v8 int16 = v4
    v9 := GenS1{F0: ((1 >= 0) || v5)}
    fmt.Println(v2, v5)
    if v5 {
        var v10 float64 = 0.0
        if false {
            v11 := GenS2{F0: 255, F1: 32767}
            v12 := GenS1{F0: false}
            v8 = v4
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: 710, F1: 971}
        _ = v10
        _ = v13
    } else {
        var v14 int64 = 173
        _ = v14
    }
    var v15 float32 = 1e10
    v8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
