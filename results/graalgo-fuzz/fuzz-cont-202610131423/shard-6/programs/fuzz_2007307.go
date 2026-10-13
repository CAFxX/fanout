package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 uint32 }

func main() {
    var v1 uint = 65535
    v2 := []uint32{uint32(v1)}
    var v3 string = "a"
    v4 := make(map[uint32]uint32)
    switch -1 {
    case -630:
        v3 = "a"
    default:
        if (58 > 255) {
            switch 9223372036854775807 {
            case 42:
                v5 := v4[65535]
                _ = v5
            }
            var v6 uint64 = 567
            delete(v4, 550)
            _ = v6
        }
    }
    var v7 float32 = (0.0 + 1e10)
    switch 127 {
    case 127:
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(v3)
            _ = fi8
        }
    case 4294967295:
        v7 = v7
    default:
        if false {
            v1 += (v1 << ((255) & 63))
            for fi9 := 0; fi9 < 15; fi9++ {
                fmt.Println(v3)
                _ = fi9
            }
        } else {
            v7 -= v7
        }
    }
    _ = len(v2)
    v10 := GenS1{F0: 65535, F1: 27}
    if (!false) {
        v11 := GenS2{F0: 9223372036854775807, F1: -412, F2: 2}
        var v12 uint64 = 946
        v13 := GenS1{F0: 65535, F1: 952}
        _ = v11
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 3; fi14++ {
            if len(v2) > 0 {
                v2[0] = 169
            }
            _ = fi14
        }
    }
    fmt.Println(v7)
    v15 := GenS2{F0: 0, F1: -596, F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
