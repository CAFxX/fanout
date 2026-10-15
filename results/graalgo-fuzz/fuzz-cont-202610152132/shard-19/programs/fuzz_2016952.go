package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int; F2 int64 }
func genFn0(p0 int16, p1 int, p2 int) int32 {
    if false {
        switch p2 {
        case -522:
            var v7 int16 = p0
            _ = v7
        case -336:
            v8 := []int{}
            _ = v8
        case -674:
            var v9 float32 = 0.1
            _ = v9
        }
    } else {
        fmt.Println(p1, p2, p0)
    }
    p2 -= p2
    var v10 float64 = (3.14159 * float64((0.0 * 0.1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2147483647
}

func main() {
    var v1 int64 = 567
    var v2 uint8 = 236
    var v3 int8 = 1
    var v4 int16 = 0
    v5 := make(map[int]int)
    var v6 uint64 = 42
    v11 := GenS1{F0: -32768, F1: (v2 - v2), F2: v3}
    if (false || (false && false)) {
        switch "foo bar" {
        case "\n\t":
            v12 := GenS1{F0: v4, F1: v2, F2: v3}
            _ = v12
        }
    } else {
        if false {
            for fi13 := 0; fi13 < 3; fi13++ {
                v6 += v6
                _ = fi13
            }
            v1 = (v1 >> ((607) & 63))
        }
    }
    v14 := GenS1{F0: (v4 * v4), F1: (v2 & (v2 / ((241) | 1))), F2: v3}
    fmt.Println(v4, v1, v1)
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 float32 = 3.14159
        switch fi15 {
        case 128:
            v17 := v5[(^(fi15))]
            _ = v17
        default:
            v18 := []int{fi15}
            _ = v18
        }
        v19 := []int{(-(fi15))}
        _ = fi15
        _ = v16
        _ = v19
    }
    fmt.Println(v4, v2)
    v1 = -348
    fmt.Println(v1, v3, v2)
    v4 = (v4 - 32767)
    var v20 float64 = ((0.0 * 1e10) - (1e10 * 0.0))
    var v21 uint8 = (v2 / ((v2) | 1))
    v2 -= (v2 >> ((255) & 63))
    for fi22 := 0; fi22 < 16; fi22++ {
        v4 += 32767
        v23 := GenS2{F0: v3, F1: v6}
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: 504, F1: 127, F2: (v1 << ((65535) & 63))}
    var v25 int16 = v4
    v6 += (v6 >> ((4294967295) & 63))
    v26 := []int{32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
