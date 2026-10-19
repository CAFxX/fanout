package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint16; F2 int64 }
func genFn0(p0 uint32) uint32 {
    p0 = (p0 & (p0 % (((^(p0))) | 1)))
    var v4 uint32 = (p0 * p0)
    _ = p0
    _ = v4
    return v4
}

func main() {
    var v1 float32 = 1e10
    var v2 uint32 = 998
    v3 := make(map[int]uint32)
    v2 = (v2 + (v2 * (v2 ^ (v2 >> ((821) & 63)))))
    if true {
        if (128 == 128) {
            v5 := []int{21, 32767}
            v1 = v1
            _ = v5
        }
    } else {
        v6 := GenS1{F0: 42}
        _ = v6
    }
    if (-2147483648 >= -2147483648) {
        switch -613 {
        case 1:
            var v7 string = "foo bar"
            _ = v7
        case -404:
            var v8 int32 = 1
            _ = v8
        case -989:
            v1 += v1
        }
        v9 := GenS2{F0: 0, F1: int32(v2), F2: 81}
        var v10 uint = 18446744073709551615
        _ = v10
        _ = v9
    }
    fmt.Println(v1, v2)
    v2 -= v2
    fmt.Println(v2, v2)
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(fi11, v2, v1)
        _ = fi11
    }
    v12 := []int{}
    var v13 int16 = -1
    var v14 float64 = 1e10
    switch 32768 {
    case -1:
        v1 = v1
    default:
        if (65535 == 42) {
            var v15 int16 = 1
            _ = v15
        } else {
            v1 -= 0.1
            var v16 float64 = v14
            v17 := GenS3{F0: -1, F1: 2, F2: -758}
            _ = v16
            _ = v17
        }
    }
    var v18 int8 = 0
    _ = len(v12)
    var v19 float64 = float64(1.5)
    switch 879 {
    case 32767:
        v12 = append(v12, -2147483648)
    case 127:
        fmt.Println(v18)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
