package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0() uint32 {
    v7 := GenS1{F0: -1, F1: true}
    var v8 uint8 = 128
    _ = v7
    _ = v8
    return 783
}
func genFn1(p0 uint32, p1 int64, p2 uint) bool {
    p1 = -209
    p1 -= int64(p2)
    switch 128 {
    case 604:
        p2 += p2
    default:
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 float64 = float64(1.5)
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS2{F0: 1, F1: 0}
    if (42 > -104) {
        p0 -= (p0 | p0)
        p1 += (p1 >> (((p2 + p2)) & 63))
    } else {
        var v12 int16 = 1
        fmt.Println(p1, p1)
        var v13 int16 = v12
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (false && true)
}

func main() {
    var v1 float64 = (1.5 / (1e10 - 0.1))
    v2 := []int{-9223372036854775808, -27, -210}
    var v3 uint16 = 74
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]uint32)
    var v6 bool = ((!(false && (209 <= 428))) && false)
    fmt.Println(v6, v1, v3)
    var v14 float32 = 0.1
    fmt.Println(v3, v14)
    if (false && (32767 == 1000)) {
        if ((v3 / ((v3) | 1)) == v3) {
            var v15 int64 = -90
            switch "a" {
            case "\n\t":
                var v16 int16 = 1
                _ = v16
            default:
                var v17 int64 = v15
                _ = v17
            }
            v1 = float64(v14)
            _ = v15
        }
        v3 = 255
    } else {
        switch -9223372036854775808 {
        case 0:
            var v18 float32 = float32(float64(-2.25))
            _ = v18
        case -783:
            var v19 uint64 = 4294967295
            _ = v19
        case -539:
            if v6 {
                v6 = true
            }
        default:
            v3 -= v3
        }
        var v20 uint8 = 223
        _ = v20
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v21 int32 = 42
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint32 = uint32(v3)
        fmt.Println(v6, v14, v3)
        _ = fi22
        _ = v23
    }
    var v24 int8 = 100
    v25 := GenS1{F0: 193, F1: (true || v6)}
    v26 := GenS2{F0: 18446744073709551615, F1: v21}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
