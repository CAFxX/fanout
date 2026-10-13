package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 int16 = 1000
    var v4 float32 = (3.14159 + 3.14159)
    if (123 >= 25) {
        var v5 int32 = -1
        v1[4294967295] = 0
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: 255, F1: v3, F2: 0}
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    var v8 int8 = -1
    if true {
        if false {
            v9 := v2[-290]
            v8 = v8
            _ = v9
        } else {
            if (((-520 <= 0) && true) && false) {
                v10 := GenS2{F0: 32767, F1: 570}
                switch 128 {
                case 403:
                    fmt.Println(v4, v8)
                }
                v8 += v8
                _ = v10
            } else {
                v11 := GenS3{F0: 2}
                v4 += (((v4 / 1e10) * (((v4 / -2.25) / v4) - v4)) / v4)
                _ = v11
            }
            if (!(((v3 * v3) == 1) && false)) {
                var v12 int = -953
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v4 = (-2.25 * float32(0.1))
        v4 += v4
        _ = fi13
    }
    var v14 uint8 = 40
    v15 := GenS2{F0: 554, F1: 32}
    v16 := []int{-9223372036854775808, -2147483648, 32767, -89}
    fmt.Println(v4)
    switch 32767 {
    case 9223372036854775807:
        var v17 uint16 = 739
        _ = v17
    }
    v18 := GenS4{F0: -2147483648}
    v4 = (v4 + 1e10)
    v3 = 929
    if false {
        v8 = v8
    } else {
        _ = len(v16)
        var v19 uint8 = v14
        var v20 float64 = 1.5
        _ = v19
        _ = v20
    }
    v16 = append(v16, -1)
    if (393 < 2) {
        fmt.Println(v4, v4, v3)
        v21 := v1[32767]
        _ = v21
    } else {
        var v22 uint16 = 42
        _ = v22
    }
    switch "a" {
    case "foo bar":
        for fi23 := 0; fi23 < 12; fi23++ {
            if (!true) {
                _ = len(v2)
                v8 = 89
            } else {
                switch fi23 {
                case 42:
                    v8 -= (v8 & (-(v8)))
                default:
                    v24 := v2[fi23]
                    _ = v24
                }
            }
            _ = fi23
        }
    case "hello":
        for fi25 := 0; fi25 < 2; fi25++ {
            for fi26 := 0; fi26 < 20; fi26++ {
                v14 += v14
                _ = fi26
            }
            _ = fi25
        }
    }
    v27 := GenS3{F0: 76}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v8
}
