package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int32 }

func main() {
    var v1 float32 = 1e10
    var v2 int16 = 1000
    v3 := []uint32{2, 843}
    v4 := make(map[string]string)
    var v5 bool = false
    v4["hello"] = ("" + "a")
    v2 = (v2 % ((v2) | 1))
    fmt.Println(v5, v1, v2)
    var v6 int16 = (v2 / (((v2 ^ v2)) | 1))
    var v7 float64 = (1.5 * (1e10 - 0.0))
    for fi8 := 0; fi8 < 18; fi8++ {
        v5 = (!(508 != 42))
        _ = fi8
    }
    v2 -= (-(v6))
    var v9 uint64 = 703
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 bool = v5
        var v12 int8 = 100
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int = (fi13 * fi13)
            v15 := GenS1{F0: 2, F1: 21, F2: 280}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    if v5 {
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
