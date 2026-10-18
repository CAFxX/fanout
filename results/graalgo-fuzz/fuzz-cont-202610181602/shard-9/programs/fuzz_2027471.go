package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    v3 := []uint32{}
    var v4 uint8 = 128
    v2 = ((782 <= 32767) && (!true))
    fmt.Println(v4, v2)
    v5 := GenS1{F0: 334, F1: 135}
    _ = len(v3)
    if false {
        v2 = (0.1 < ((1.5 + 3.14159) - (3.14159 - (-2.25 - 1e10))))
    }
    v1[-303] = 191
    v6 := GenS2{F0: 1}
    for fi7 := 0; fi7 < 17; fi7++ {
        v2 = v2
        fmt.Println(fi7, v4)
        if (-66 != 1) {
            var v8 uint64 = 874
            v4 = (^(v4))
            _ = v8
        } else {
            fmt.Println(v2)
            var v9 float64 = -2.25
            var v10 uint64 = 301
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    switch "hello" {
    case "":
        var v11 uint16 = 261
        _ = v11
    }
    var v12 int32 = 1
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
