package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    v3 := []uint32{2}
    var v4 uint8 = 128
    var v5 uint8 = (v4 & 213)
    var v6 float32 = (0.0 / (-2.25 / 1.5))
    v7 := GenS1{F0: 4294967295, F1: 42}
    v6 = v6
    v8 := GenS2{F0: 4294967295}
    v9 := GenS1{F0: 42, F1: 128}
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: -58, F1: (fi10 << ((777) & 63))}
        var v12 uint8 = (v5 / ((v4) | 1))
        fmt.Println(v1, v4)
        _ = fi10
        _ = v11
        _ = v12
    }
    v6 = (float32((0.0 + 3.14159)) * (1.5 * 0.0))
    v6 -= ((v6 / 1e10) + 0.1)
    var v13 string = (v1 + v1)
    v2[1] = 0
    if false {
        v13 = "\n\t"
        if false {
            v14 := v2[2147483647]
            v15 := GenS2{F0: -380}
            _ = v14
            _ = v15
        }
        var v16 uint64 = 18446744073709551615
        _ = v16
    }
    fmt.Println(v5, v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
