package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint64 }

func main() {
    var v1 int32 = 1
    var v2 int32 = (v1 * (v1 % ((v1) | 1)))
    var v3 float64 = (float64(-2.25) + (-2.25 * 1.5))
    v4 := []uint32{2147483648, 4294967295, 1, 2147483648}
    var v5 uint32 = 182
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint64 = 4294967295
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 2
    switch -1 {
    case 9223372036854775807:
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 float32 = 1e10
            _ = fi9
            _ = v10
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v2)
    if true {
        v2 = (-(v1))
        v11 := GenS1{F0: -944, F1: false, F2: 178}
        _ = v11
    }
    var v12 uint16 = 2
    if true {
        _ = len(v4)
        if len(v4) > 0 {
            v4[0] = (^(v5))
        }
    } else {
        v3 -= v3
        var v13 int8 = 100
        _ = v13
    }
    v14 := GenS1{F0: 127, F1: true, F2: 42}
    v5 = (v5 + (v5 >> ((2) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
