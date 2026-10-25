package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint64 = 2
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 int64 = -9223372036854775808
        _ = fi4
        _ = v5
    }
    v3 -= v3
    v3 = 422
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(v3, v3)
        _ = fi6
    }
    var v7 uint64 = (v3 - (v3 * v3))
    v8 := []int{32767}
    v3 -= v7
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 int32 = 42
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 65535
    switch "a" {
    case "\n\t":
        v12 := GenS1{F0: 65535, F1: v11}
        _ = v12
    case "hello":
        _ = len(v8)
    case "":
        v11 = 513
    default:
        var v13 float64 = 0.0
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS1{F0: 1, F1: v11}
        _ = fi14
        _ = v15
    }
    _ = len(v1)
    if false {
        for fi16 := 0; fi16 < 16; fi16++ {
            v3 += v3
            var v17 uint64 = v7
            _ = fi16
            _ = v17
        }
        var v18 int = 32767
        v19 := GenS1{F0: 1, F1: (v11 % (((v11 & (v11 - v11))) | 1))}
        _ = v18
        _ = v19
    } else {
        fmt.Println(v7, v7)
        var v20 uint8 = 2
        if true {
            v21 := GenS1{F0: 42, F1: uint16(v3)}
            _ = v21
        }
        _ = v20
    }
    fmt.Println(v11)
    var v22 uint32 = 2147483647
    v23 := GenS2{F0: 1, F1: 128, F2: 2}
    var v24 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v8
}
