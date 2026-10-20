package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint; F1 int8; F2 uint32 }

func main() {
    var v1 uint = 42
    var v2 uint64 = 51
    v3 := []uint32{}
    if false {
        if true {
            var v4 float32 = (((3.14159 * 1e10) - (1e10 / 0.1)) + 0.1)
            v1 -= (^(v1))
            _ = v4
        }
        fmt.Println(v1, v2)
        v5 := GenS1{F0: 828, F1: -2147483648}
        _ = v5
    }
    var v6 uint8 = 127
    v7 := GenS2{F0: (^(v1)), F1: 0, F2: uint32(v2)}
    for fi8 := 0; fi8 < 8; fi8++ {
        switch fi8 {
        case -763:
            var v9 uint64 = 18446744073709551615
            _ = v9
        default:
            v6 += (v6 * 127)
        }
        _ = fi8
    }
    _ = len(v3)
    if (32767 <= -32768) {
        v10 := GenS1{F0: -621, F1: 0}
        fmt.Println(v2)
        fmt.Println(v2, v1)
        _ = v10
    }
    if len(v3) > 0 {
        v3[0] = uint32(v1)
    }
    if false {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint32 = 2
            _ = fi11
            _ = v12
        }
        var v13 int64 = -1
        _ = v13
    } else {
        v14 := GenS2{F0: 255, F1: -42, F2: 2}
        _ = v14
    }
    var v15 float64 = 1e10
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(fi16, v15)
        v17 := GenS1{F0: 9223372036854775807, F1: 42}
        _ = fi16
        _ = v17
    }
    var v18 int = -182
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint32 = 0
        _ = len(v3)
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
