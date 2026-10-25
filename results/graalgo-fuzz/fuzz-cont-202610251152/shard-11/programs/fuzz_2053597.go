package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int32 }

func main() {
    var v1 int16 = 42
    v2 := []uint32{}
    var v3 int32 = 676
    v4 := make(map[string]uint32)
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 float32 = -2.25
        v6 -= v6
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 184
    var v8 int64 = 9223372036854775807
    var v9 float64 = 0.1
    if (!false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v1 += (v1 - (v1 >> ((42) & 63)))
    v9 -= (float64(0.1) * (0.1 / v9))
    v8 -= -9223372036854775808
    v10 := GenS1{F0: 52, F1: 42, F2: 42}
    fmt.Println(v9, v7, v7)
    var v11 uint16 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
