package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    var v1 int64 = 1
    v2 := []uint32{218}
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    var v5 uint = 42
    v5 = v5
    var v6 uint8 = 0
    var v7 uint = v5
    var v8 int16 = -32768
    v9 := GenS1{F0: 0, F1: (v1 >> (((^(v7))) & 63))}
    switch 128 {
    case -54:
        fmt.Println(v6)
    case 367:
        v10 := GenS2{F0: -1, F1: true}
        _ = v10
    case 127:
        if (!(v5 <= uint(v6))) {
            var v11 int8 = -117
            var v12 uint32 = 0
            _ = v11
            _ = v12
        }
    }
    if (int32(v1) != -2147483648) {
        var v13 uint = (v5 + (v7 ^ (^(v5))))
        v1 += v1
        var v14 string = ""
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 862, F1: (!true)}
    v16 := GenS1{F0: 0, F1: (v1 << ((v5) & 63))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
