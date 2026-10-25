package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0() uint8 {
    v7 := GenS1{F0: (!false), F1: -2147483648}
    var v8 uint16 = 41
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint8 = 165
    var v2 uint16 = 2
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    v5 := []uint32{}
    var v6 int64 = 9223372036854775807
    fmt.Println(v1, v1, v2)
    switch -831 {
    case 4294967295:
        fmt.Println(v1, v2)
    }
    v1 -= genFn0()
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int = fi9
        var v11 uint64 = 4294967295
        var v12 uint16 = (v2 % (((v2 | 150)) | 1))
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    delete(v4, 570)
    v2 -= 1
    if true {
        v2 = (v2 << ((0) & 63))
    } else {
        if ((-2.25 * float64((-2.25 / 0.1))) > (3.14159 * (1e10 / -2.25))) {
            v1 = (v1 << ((0) & 63))
            v6 = (v6 - v6)
        }
    }
    v1 = 132
    var v13 uint32 = 200
    var v14 uint16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
