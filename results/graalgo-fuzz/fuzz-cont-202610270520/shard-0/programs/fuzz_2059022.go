package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint8 = 127
    var v2 int = -133
    var v3 int8 = 1
    var v4 float32 = 1e10
    v5 := make(map[uint32]uint32)
    var v6 bool = ((v3 | 95) <= (^(v3)))
    v7 := []int{4294967295, 32768}
    fmt.Println(v6, v3, v1)
    v4 -= -2.25
    var v8 float32 = v4
    for fi9 := 0; fi9 < 19; fi9++ {
        switch "" {
        case "":
            v10 := GenS1{F0: -1, F1: 930}
            _ = v10
        }
        _ = fi9
    }
    v8 -= float32(((0.1 * 1e10) * float64(v8)))
    if true {
        var v11 uint8 = v1
        fmt.Println(v8)
        _ = v11
    } else {
        var v12 float64 = (0.1 / (-2.25 / (0.1 + (float64(0.0) / 0.1))))
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
