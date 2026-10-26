package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint64; F1 uint16; F2 int16 }
func genFn0(p0 int32, p1 bool, p2 uint64) uint8 {
    p2 -= 18446744073709551615
    switch 0 {
    case 0:
        var v4 int = 193
        _ = v4
    case 127:
        v5 := GenS1{F0: 42, F1: 737}
        _ = v5
    case -2147483648:
        var v6 bool = (-1 <= 42)
        _ = v6
    }
    v7 := GenS1{F0: 9223372036854775807, F1: -946}
    for fi8 := 0; fi8 < 8; fi8++ {
        p0 = (p0 + p0)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 128
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]int)
    var v3 uint = 65535
    v3 = (v3 + v3)
    v3 -= 255
    var v9 int = -58
    var v10 uint = v3
    var v11 int16 = 858
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS1{F0: 9223372036854775807, F1: (fi12 + 42)}
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: 2, F1: 1, F2: (v11 * v11)}
    var v15 uint8 = 127
    switch (^(v9)) {
    case 753:
        fmt.Println(v9, v11)
    case 128:
        fmt.Println(v3, v3, v3)
    default:
        for fi16 := 0; fi16 < 12; fi16++ {
            if (false || false) {
                v15 += v15
                switch 1 {
                case 506:
                    var v17 uint32 = 65535
                    _ = v17
                case 962:
                    var v18 uint32 = 0
                    _ = v18
                default:
                    v10 = v3
                }
            }
            _ = fi16
        }
    }
    if (65535 != 816) {
        var v19 int64 = 1
        _ = v19
    } else {
        var v20 int32 = -230
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        v9 += (v9 | v9)
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
