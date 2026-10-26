package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8; F1 uint }

func main() {
    var v1 int16 = -1
    v2 := []uint32{}
    v3 := make(map[string]int)
    if ("" > "\n\t") {
        var v4 int8 = 0
        var v5 int64 = -9223372036854775808
        _ = v4
        _ = v5
    } else {
        v6 := GenS1{F0: ("a" + "a"), F1: 42, F2: -1}
        if (1 <= 65535) {
            v7 := GenS2{F0: 913, F1: (690 >= 0)}
            _ = v7
        }
        v8 := GenS2{F0: 42, F1: true}
        _ = v6
        _ = v8
    }
    v1 -= (v1 * v1)
    if (2147483647 > 0) {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint64 = 42
            v1 -= (v1 + (-(v1)))
            v11 := GenS3{F0: 127, F1: 0}
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint64 = 42
        _ = v12
    }
    if false {
        v1 += (v1 - 1)
        fmt.Println(v1)
        var v13 uint64 = 0
        _ = v13
    } else {
        fmt.Println(v1)
        v14 := GenS3{F0: 42, F1: 788}
        var v15 float32 = 0.0
        _ = v14
        _ = v15
    }
    var v16 int32 = 0
    switch 490 {
    case -175:
        if false {
            if len(v2) > 0 {
                v2[0] = 0
            }
            var v17 uint64 = 42
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v17
        }
    case -1:
        var v18 int8 = -1
        _ = v18
    case -11:
        if false {
            var v19 uint8 = 127
            _ = v19
        } else {
            if (!false) {
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 uint64 = 4294967295
                    _ = fi20
                    _ = v21
                }
            }
        }
    default:
        v1 = (v1 >> ((255) & 63))
    }
    v16 -= -1
    if true {
        fmt.Println(v16, v1)
    } else {
        v16 -= v16
        fmt.Println(v16, v1)
        fmt.Println(v1, v16)
    }
    var v22 int8 = -1
    var v23 float64 = 0.0
    v23 -= (v23 * -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
