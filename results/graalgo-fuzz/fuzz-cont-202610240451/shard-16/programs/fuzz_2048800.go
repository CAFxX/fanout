package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    var v1 int64 = 0
    v2 := []string{"", "\n\t", "\n\t"}
    var v3 int64 = -1
    var v4 int8 = -78
    v5 := GenS1{F0: -1}
    var v6 uint16 = 888
    v7 := GenS2{F0: 0, F1: 42}
    fmt.Println(v4)
    for fi8 := 0; fi8 < 15; fi8++ {
        v1 = (v1 % ((v1) | 1))
        _ = fi8
    }
    var v9 uint = 1
    if ((3.14159 + ((0.0 + (float32(0.1) / (1.5 / 1.0))) + float32(-2.25))) > 1e10) {
        var v10 uint16 = v6
        var v11 uint = uint(v6)
        var v12 string = "a"
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        if (false || true) {
            fmt.Println(fi13, fi13, v9)
            v9 += v9
        }
        var v14 int = fi13
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
