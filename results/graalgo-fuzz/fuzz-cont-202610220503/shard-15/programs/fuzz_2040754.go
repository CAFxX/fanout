package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 uint16, p1 bool, p2 uint) int32 {
    var v5 float32 = 0.0
    v6 := []int{128}
    var v7 int8 = -128
    p2 = (p2 * (p2 >> (((p2 % ((430) | 1))) & 63)))
    for fi8 := 0; fi8 < 3; fi8++ {
        if (-1 < 32767) {
            v7 = v7
        } else {
            var v9 int8 = (v7 << ((18446744073709551615) & 63))
            p0 += (^(p0))
            var v10 int64 = 9223372036854775807
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: (p2 & 0)}
        if ((!((^(p2)) <= 42)) || true) {
            if p1 {
                var v12 int = (fi8 & (-(fi8)))
                _ = v12
            }
        } else {
            p0 += p0
            _ = len(v6)
            p1 = p1
        }
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 793
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]uint32)
    var v3 uint64 = 2
    var v4 uint = uint(v3)
    var v13 int8 = 127
    v14 := GenS2{F0: -739, F1: -899}
    v1 -= (v1 << ((v4) & 63))
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := []int{fi15}
        _ = fi15
        _ = v16
    }
    fmt.Println(v13)
    var v17 uint64 = v3
    var v18 uint = (^(v4))
    fmt.Println(v4)
    v19 := []int{-239, 32768}
    switch 42 {
    case 32767:
        switch 4294967295 {
        case 127:
            v13 = v13
        }
    case 0:
        switch 128 {
        case -2147483648:
            var v20 uint32 = 807
            _ = v20
        }
    default:
        var v21 int = 32768
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
