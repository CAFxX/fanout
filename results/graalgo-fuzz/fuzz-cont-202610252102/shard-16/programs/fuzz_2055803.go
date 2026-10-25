package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int32; F1 int; F2 uint64 }

func main() {
    var v1 uint64 = 973
    var v2 uint8 = 81
    var v3 int64 = 1
    v4 := []string{}
    var v5 int64 = v3
    v6 := []int{-268, -854}
    var v7 uint64 = v1
    v8 := GenS1{F0: -128, F1: 255, F2: v3}
    fmt.Println(v7, v5, v5)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint8 = (^(v2))
        var v11 uint64 = 2
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v1, v7)
    v2 -= (v2 - (v2 & 255))
    fmt.Println(v7, v3, v7)
    var v12 int64 = (-(v5))
    v7 = v1
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS2{F0: 0, F1: 2147483647, F2: uint64(v2)}
        v12 += (v12 ^ (-(v12)))
        _ = fi13
        _ = v14
    }
    switch ("" + "") {
    case "hello":
        var v15 int16 = -32768
        _ = v15
    case "":
        v16 := GenS2{F0: 0, F1: 32768, F2: (v7 % (((^(v7))) | 1))}
        _ = v16
    default:
        var v17 int16 = -582
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
