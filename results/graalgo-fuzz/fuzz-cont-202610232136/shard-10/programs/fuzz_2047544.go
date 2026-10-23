package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int; F2 bool }
type GenS4 struct { F0 uint16; F1 int16; F2 int16 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := GenS1{F0: 0, F1: (-(fi4))}
        v6 := GenS2{F0: -1}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS3{F0: -2147483648, F1: 280, F2: ((0.0 - -2.25) == (0.0 + ((0.0 + 1e10) + 0.1)))}
    var v8 uint64 = 18446744073709551615
    fmt.Println(v8)
    _ = v7
    _ = v8
    return -107
}

func main() {
    v1 := []int{-2147483648, 94, 964, -2147483648}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    var v9 int8 = 116
    switch 814 {
    case 42:
        fmt.Println(v9)
    case -401:
        v9 = (-(v9))
    case 117:
        fmt.Println(v9)
    }
    var v10 uint8 = 42
    var v11 int64 = 753
    if (!false) {
        if ((v9 + (v9 & (v9 ^ v9))) == v9) {
            v12 := GenS3{F0: 2147483647, F1: 276, F2: ((v10 ^ v10) > (v10 << ((255) & 63)))}
            _ = v12
        } else {
            fmt.Println(v11, v9, v9)
        }
    } else {
        v11 = v11
    }
    v11 -= 42
    switch 637 {
    case 2147483647:
        for fi13 := 0; fi13 < 9; fi13++ {
            v10 = v10
            _ = fi13
        }
    case 4294967295:
        var v14 uint = 719
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS4{F0: 42, F1: 32767, F2: -62}
        _ = fi15
        _ = v16
    }
    v11 -= v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
