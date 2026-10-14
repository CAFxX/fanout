package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint32 }

func main() {
    v1 := []uint32{65535, 0, 1, 0}
    var v2 uint16 = 888
    v3 := make(map[int]uint32)
    var v4 uint64 = 42
    var v5 int = -1
    var v6 float64 = (3.14159 / 1e10)
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = 65535
    }
    v7 := GenS1{F0: (v5 >> ((2) & 63)), F1: 1, F2: 481}
    var v8 float64 = ((v6 - (v6 - v6)) / v6)
    fmt.Println(v4, v8)
    var v9 int32 = 0
    v6 = ((0.0 / v8) / 3.14159)
    _ = len(v1)
    var v10 int8 = 22
    for fi11 := 0; fi11 < 2; fi11++ {
        switch fi11 {
        case -734:
            _ = len(v3)
        default:
            v2 = (^(v2))
        }
        delete(v3, v5)
        _ = fi11
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
