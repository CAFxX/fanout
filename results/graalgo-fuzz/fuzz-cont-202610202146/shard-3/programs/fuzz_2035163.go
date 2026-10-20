package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint32; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint64 }
func genFn0() uint {
    var v4 int16 = -341
    v4 = (v4 | (^(v4)))
    fmt.Println(v4)
    v4 += v4
    v5 := GenS1{F0: 9223372036854775807}
    _ = v4
    _ = v5
    return 517
}
func genFn1(p0 int) string {
    v6 := GenS2{F0: -9223372036854775808, F1: 462, F2: 42}
    v7 := []int{(p0 % ((p0) | 1)), 1, p0, -1}
    v8 := GenS3{F0: (2 <= 46)}
    switch p0 {
    case 0:
        switch p0 {
        case 32767:
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 uint64 = 1
                _ = fi9
                _ = v10
            }
        default:
            fmt.Println(p0)
        }
    case -547:
        v11 := GenS4{F0: 1, F1: 2, F2: 42}
        _ = v11
    case -2147483648:
        p0 -= (-(p0))
    default:
        p0 -= (^(p0))
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    var v1 uint = 32
    v2 := make(map[int]string)
    var v3 uint16 = 0
    v1 += v1
    v3 += 0
    var v12 uint16 = v3
    v12 = v12
    var v13 uint = v1
    switch 582 {
    case -2147483648:
        v1 = 258
    case 32768:
        for fi14 := 0; fi14 < 20; fi14++ {
            for fi15 := 0; fi15 < 20; fi15++ {
                v16 := []int{(fi15 + -698), fi15, (fi15 >> (((v13 + v13)) & 63))}
                _ = fi15
                _ = v16
            }
            v12 = (v3 + (v12 - v12))
            var v17 uint16 = v12
            _ = fi14
            _ = v17
        }
    default:
        var v18 int8 = 13
        _ = v18
    }
    v3 -= v12
    v19 := []int{}
    delete(v2, 42)
    var v20 int64 = -810
    var v21 int8 = 51
    fmt.Println(v3, v13, v21)
    if (0 > 0) {
        var v22 string = "hello"
        for fi23 := 0; fi23 < 6; fi23++ {
            v24 := v2[int(v20)]
            v20 = v20
            _ = fi23
            _ = v24
        }
        v22 = v22
        _ = v22
    }
    var v25 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
}
