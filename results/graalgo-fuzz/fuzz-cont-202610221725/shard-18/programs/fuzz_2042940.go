package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{4294967295, 2147483647, 65535}
    var v3 int16 = 1000
    var v4 uint16 = 42
    v5 := make(map[uint32]uint32)
    var v6 int = 912
    v7 := GenS1{F0: -128, F1: -1, F2: 1}
    v4 = (v4 * 1)
    var v8 uint = 398
    fmt.Println(v8)
    if (!((((!true) && false) || (((-2.25 + float32(1.5)) > 0.0) || true)) || false)) {
        var v9 float32 = (0.1 * 0.0)
        for fi10 := 0; fi10 < 14; fi10++ {
            switch fi10 {
            case -2147483648:
                var v11 uint64 = 18446744073709551615
                _ = v11
            default:
                v9 = float32((float64((v9 - v9)) + ((1.5 * -2.25) - (3.14159 - 1e10))))
            }
            v3 = v3
            _ = fi10
        }
        _ = v9
    } else {
        var v12 uint = 42
        _ = v12
    }
    var v13 uint = 279
    fmt.Println(v8)
    if (true && false) {
        var v14 float32 = 1e10
        var v15 int8 = -128
        _ = v14
        _ = v15
    } else {
        _ = len(v2)
        var v16 uint64 = 18446744073709551615
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
