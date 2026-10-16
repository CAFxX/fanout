package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS4 struct { F0 int; F1 int16; F2 bool }
type GenS5 struct { F0 uint32; F1 int8; F2 int8 }
type GenS6 struct { F0 uint16; F1 bool }

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    v4 := make(map[string]bool)
    var v5 int8 = 0
    var v6 int64 = 675
    v1 += (v1 >> ((18446744073709551615) & 63))
    var v7 uint = 595
    switch 9223372036854775807 {
    case 2147483647:
        if (true || (970 >= 405)) {
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := GenS1{F0: v5, F1: "foo bar"}
                delete(v4, "")
                _ = fi8
                _ = v9
            }
            v5 -= (v5 - (v5 & v5))
        } else {
            v10 := GenS2{F0: (v6 % ((v6) | 1))}
            v11 := GenS3{F0: 1, F1: 4294967295, F2: 40}
            _ = v10
            _ = v11
        }
    case -565:
        switch 1 {
        case 42:
            v1 = -1
        case -1:
            v12 := GenS4{F0: 0, F1: v1, F2: (int32(v5) < 0)}
            _ = v12
        case 2147483647:
            v7 += (v7 ^ v7)
        default:
            v13 := GenS4{F0: 9223372036854775807, F1: -1, F2: (32768 > 2147483647)}
            _ = v13
        }
    case 4294967295:
        var v14 uint8 = 2
        _ = v14
    }
    v15 := GenS5{F0: uint32(v7), F1: (-(v5)), F2: 0}
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS6{F0: 439, F1: (!false)}
        v5 -= (-(v5))
        v18 := GenS5{F0: 4294967295, F1: v5, F2: v5}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: (-(v6))}
    fmt.Println(v7, v1)
    v7 += v7
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 float32 = float32(0.1)
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
