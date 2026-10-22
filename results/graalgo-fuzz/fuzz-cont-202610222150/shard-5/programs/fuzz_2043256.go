package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 32767
    var v3 float32 = 0.0
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 float64 = float64((v3 / 3.14159))
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 0
    v8 := []int{v2, (v2 - v2), (v2 & (v2 + (v2 | 769)))}
    v3 = (v3 / 1e10)
    v9 := v4[(-(v2))]
    var v10 float32 = (v3 * 1e10)
    v11 := GenS1{F0: (-(v2)), F1: 1}
    var v12 uint = 0
    var v13 uint8 = 255
    if false {
        _ = len(v1)
        fmt.Println(v10, v10, v7)
    } else {
        var v14 int8 = -56
        v15 := GenS1{F0: 2147483647, F1: -285}
        switch (^(v2)) {
        case 2147483647:
            var v16 uint16 = v7
            _ = v16
        case 60:
            v10 = (0.1 + float32((0.0 + (1.5 * ((((1e10 + 3.14159) * 1.5) + -2.25) / float64(v3))))))
        case 128:
            v17 := GenS2{F0: (v13 << ((4294967295) & 63))}
            _ = v17
        default:
            var v18 int8 = (v14 ^ -1)
            _ = v18
        }
        _ = v14
        _ = v15
    }
    var v19 uint32 = 368
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
